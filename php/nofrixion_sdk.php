<?php
declare(strict_types=1);

// Nofrixion SDK

require_once __DIR__ . '/utility/struct/Struct.php';
require_once __DIR__ . '/core/UtilityType.php';
require_once __DIR__ . '/core/Spec.php';
require_once __DIR__ . '/core/Helpers.php';

// Load utility registration
require_once __DIR__ . '/utility/Register.php';

// Load config and features
require_once __DIR__ . '/config.php';
require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/features.php';

use Voxgig\Struct\Struct;

class NofrixionSDK
{
    public string $mode;
    public array $features;
    public ?array $options;

    private $_utility;
    private $_rootctx;

    public function __construct(array $options = [])
    {
        $this->mode = "live";
        $this->features = [];
        $this->options = null;

        $utility = new NofrixionUtility();
        $this->_utility = $utility;

        $config = NofrixionConfig::make_config();

        $this->_rootctx = ($utility->make_context)([
            "client" => $this,
            "utility" => $utility,
            "config" => $config,
            "options" => $options ?? [],
            "shared" => [],
        ], null);

        $this->options = ($utility->make_options)($this->_rootctx);

        if (Struct::getpath($this->options, "feature.test.active") === true) {
            $this->mode = "test";
        }

        $this->_rootctx->options = $this->options;

        // Add features from config.
        $feature_opts = NofrixionHelpers::to_map(Struct::getprop($this->options, "feature"));
        if ($feature_opts) {
            $items = Struct::items($feature_opts);
            if ($items) {
                foreach ($items as $item) {
                    $fname = $item[0];
                    $fopts = NofrixionHelpers::to_map($item[1]);
                    if ($fopts && isset($fopts["active"]) && $fopts["active"] === true) {
                        ($utility->feature_add)($this->_rootctx, NofrixionFeatures::make_feature($fname));
                    }
                }
            }
        }

        // Add extension features.
        $extend_val = Struct::getprop($this->options, "extend");
        if (is_array($extend_val)) {
            foreach ($extend_val as $f) {
                if (is_object($f) && method_exists($f, 'get_name')) {
                    ($utility->feature_add)($this->_rootctx, $f);
                }
            }
        }

        // Initialize features.
        foreach ($this->features as $f) {
            ($utility->feature_init)($this->_rootctx, $f);
        }

        ($utility->feature_hook)($this->_rootctx, "PostConstruct");
    }

    public function options_map(): array
    {
        $out = Struct::clone($this->options);
        return is_array($out) ? $out : [];
    }

    public function get_utility()
    {
        return NofrixionUtility::copy($this->_utility);
    }

    public function get_root_ctx()
    {
        return $this->_rootctx;
    }

    public function prepare(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;
        $fetchargs = $fetchargs ?? [];

        $ctrl = NofrixionHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "prepare",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $opts = $this->options;
        $path = Struct::getprop($fetchargs, "path") ?? "";
        $path = is_string($path) ? $path : "";
        $method_val = Struct::getprop($fetchargs, "method") ?? "GET";
        $method_val = is_string($method_val) ? $method_val : "GET";
        $params = NofrixionHelpers::to_map(Struct::getprop($fetchargs, "params")) ?? [];
        $query = NofrixionHelpers::to_map(Struct::getprop($fetchargs, "query")) ?? [];
        $headers = ($utility->prepare_headers)($ctx);

        $base = Struct::getprop($opts, "base") ?? "";
        $base = is_string($base) ? $base : "";
        $prefix = Struct::getprop($opts, "prefix") ?? "";
        $prefix = is_string($prefix) ? $prefix : "";
        $suffix = Struct::getprop($opts, "suffix") ?? "";
        $suffix = is_string($suffix) ? $suffix : "";

        $ctx->spec = new NofrixionSpec([
            "base" => $base, "prefix" => $prefix, "suffix" => $suffix,
            "path" => $path, "method" => $method_val,
            "params" => $params, "query" => $query, "headers" => $headers,
            "body" => Struct::getprop($fetchargs, "body"),
            "step" => "start",
        ]);

        // Merge user-provided headers.
        $uh = Struct::getprop($fetchargs, "headers");
        if (is_array($uh)) {
            foreach ($uh as $k => $v) {
                $ctx->spec->headers[$k] = $v;
            }
        }

        [$_, $err] = ($utility->prepare_auth)($ctx);
        if ($err) {
            return ($utility->make_error)($ctx, $err);
        }

        [$fetchdef, $fd_err] = ($utility->make_fetch_def)($ctx);
        if ($fd_err) {
            return ($utility->make_error)($ctx, $fd_err);
        }
        return $fetchdef;
    }

    public function direct(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;

        // direct() is the raw-HTTP escape hatch: it never throws, it returns
        // an {ok, err, ...} dict. prepare() now raises on error, so catch it
        // and surface the failure through the dict instead.
        try {
            $fetchdef = $this->prepare($fetchargs);
        } catch (\Throwable $err) {
            return ["ok" => false, "err" => $err];
        }

        $fetchargs = $fetchargs ?? [];
        $ctrl = NofrixionHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "direct",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $url = $fetchdef["url"] ?? "";
        [$fetched, $fetch_err] = ($utility->fetcher)($ctx, $url, $fetchdef);

        if ($fetch_err) {
            return ["ok" => false, "err" => $fetch_err];
        }

        if ($fetched === null) {
            return [
                "ok" => false,
                "err" => $ctx->make_error("direct_no_response", "response: undefined"),
            ];
        }

        if (is_array($fetched)) {
            $status = NofrixionHelpers::to_int(Struct::getprop($fetched, "status"));
            $headers = Struct::getprop($fetched, "headers") ?? [];

            // No-body responses (204, 304) and explicit zero content-length
            // must skip JSON parsing — calling json() on an empty body errors.
            $content_length = is_array($headers) ? ($headers["content-length"] ?? null) : null;
            $no_body = $status === 204 || $status === 304 || (string)$content_length === "0";

            $json_data = null;
            if (!$no_body) {
                $jf = Struct::getprop($fetched, "json");
                if (is_callable($jf)) {
                    try {
                        $json_data = $jf();
                    } catch (\Throwable $e) {
                        // Non-JSON body — leave data null but keep status/ok.
                        $json_data = null;
                    }
                }
            }

            return [
                "ok" => $status >= 200 && $status < 300,
                "status" => $status,
                "headers" => Struct::getprop($fetched, "headers"),
                "data" => $json_data,
            ];
        }

        return [
            "ok" => false,
            "err" => $ctx->make_error("direct_invalid", "invalid response type"),
        ];
    }


    private $_account = null;

    // Canonical facade: $client->Account()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->account()
    // resolves here too.
    public function Account($data = null)
    {
        require_once __DIR__ . '/entity/account_entity.php';
        if ($data === null) {
            if ($this->_account === null) {
                $this->_account = new AccountEntity($this, null);
            }
            return $this->_account;
        }
        return new AccountEntity($this, $data);
    }


    private $_beneficiary = null;

    // Canonical facade: $client->Beneficiary()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->beneficiary()
    // resolves here too.
    public function Beneficiary($data = null)
    {
        require_once __DIR__ . '/entity/beneficiary_entity.php';
        if ($data === null) {
            if ($this->_beneficiary === null) {
                $this->_beneficiary = new BeneficiaryEntity($this, null);
            }
            return $this->_beneficiary;
        }
        return new BeneficiaryEntity($this, $data);
    }


    private $_cancel = null;

    // Canonical facade: $client->Cancel()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->cancel()
    // resolves here too.
    public function Cancel($data = null)
    {
        require_once __DIR__ . '/entity/cancel_entity.php';
        if ($data === null) {
            if ($this->_cancel === null) {
                $this->_cancel = new CancelEntity($this, null);
            }
            return $this->_cancel;
        }
        return new CancelEntity($this, $data);
    }


    private $_disable = null;

    // Canonical facade: $client->Disable()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->disable()
    // resolves here too.
    public function Disable($data = null)
    {
        require_once __DIR__ . '/entity/disable_entity.php';
        if ($data === null) {
            if ($this->_disable === null) {
                $this->_disable = new DisableEntity($this, null);
            }
            return $this->_disable;
        }
        return new DisableEntity($this, $data);
    }


    private $_enable = null;

    // Canonical facade: $client->Enable()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->enable()
    // resolves here too.
    public function Enable($data = null)
    {
        require_once __DIR__ . '/entity/enable_entity.php';
        if ($data === null) {
            if ($this->_enable === null) {
                $this->_enable = new EnableEntity($this, null);
            }
            return $this->_enable;
        }
        return new EnableEntity($this, $data);
    }


    private $_merchant = null;

    // Canonical facade: $client->Merchant()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->merchant()
    // resolves here too.
    public function Merchant($data = null)
    {
        require_once __DIR__ . '/entity/merchant_entity.php';
        if ($data === null) {
            if ($this->_merchant === null) {
                $this->_merchant = new MerchantEntity($this, null);
            }
            return $this->_merchant;
        }
        return new MerchantEntity($this, $data);
    }


    private $_metadata = null;

    // Canonical facade: $client->Metadata()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->metadata()
    // resolves here too.
    public function Metadata($data = null)
    {
        require_once __DIR__ . '/entity/metadata_entity.php';
        if ($data === null) {
            if ($this->_metadata === null) {
                $this->_metadata = new MetadataEntity($this, null);
            }
            return $this->_metadata;
        }
        return new MetadataEntity($this, $data);
    }


    private $_no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page = null;

    // Canonical facade: $client->NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page()
    // resolves here too.
    public function NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page === null) {
                $this->_no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page = new NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity($this, null);
            }
            return $this->_no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page;
        }
        return new NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity($this, $data);
    }


    private $_no_frixion_biz_biz_models_paging_payment_request_page = null;

    // Canonical facade: $client->NoFrixionBizBizModelsPagingPaymentRequestPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_biz_biz_models_paging_payment_request_page()
    // resolves here too.
    public function NoFrixionBizBizModelsPagingPaymentRequestPage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_biz_biz_models_paging_payment_request_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_biz_biz_models_paging_payment_request_page === null) {
                $this->_no_frixion_biz_biz_models_paging_payment_request_page = new NoFrixionBizBizModelsPagingPaymentRequestPageEntity($this, null);
            }
            return $this->_no_frixion_biz_biz_models_paging_payment_request_page;
        }
        return new NoFrixionBizBizModelsPagingPaymentRequestPageEntity($this, $data);
    }


    private $_no_frixion_biz_biz_models_paging_payout_page = null;

    // Canonical facade: $client->NoFrixionBizBizModelsPagingPayoutPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_biz_biz_models_paging_payout_page()
    // resolves here too.
    public function NoFrixionBizBizModelsPagingPayoutPage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_biz_biz_models_paging_payout_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_biz_biz_models_paging_payout_page === null) {
                $this->_no_frixion_biz_biz_models_paging_payout_page = new NoFrixionBizBizModelsPagingPayoutPageEntity($this, null);
            }
            return $this->_no_frixion_biz_biz_models_paging_payout_page;
        }
        return new NoFrixionBizBizModelsPagingPayoutPageEntity($this, $data);
    }


    private $_no_frixion_biz_biz_models_paging_payrun_page = null;

    // Canonical facade: $client->NoFrixionBizBizModelsPagingPayrunPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_biz_biz_models_paging_payrun_page()
    // resolves here too.
    public function NoFrixionBizBizModelsPagingPayrunPage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_biz_biz_models_paging_payrun_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_biz_biz_models_paging_payrun_page === null) {
                $this->_no_frixion_biz_biz_models_paging_payrun_page = new NoFrixionBizBizModelsPagingPayrunPageEntity($this, null);
            }
            return $this->_no_frixion_biz_biz_models_paging_payrun_page;
        }
        return new NoFrixionBizBizModelsPagingPayrunPageEntity($this, $data);
    }


    private $_no_frixion_biz_biz_models_paging_rule_events_page = null;

    // Canonical facade: $client->NoFrixionBizBizModelsPagingRuleEventsPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_biz_biz_models_paging_rule_events_page()
    // resolves here too.
    public function NoFrixionBizBizModelsPagingRuleEventsPage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_biz_biz_models_paging_rule_events_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_biz_biz_models_paging_rule_events_page === null) {
                $this->_no_frixion_biz_biz_models_paging_rule_events_page = new NoFrixionBizBizModelsPagingRuleEventsPageEntity($this, null);
            }
            return $this->_no_frixion_biz_biz_models_paging_rule_events_page;
        }
        return new NoFrixionBizBizModelsPagingRuleEventsPageEntity($this, $data);
    }


    private $_no_frixion_biz_biz_models_paging_rules_page = null;

    // Canonical facade: $client->NoFrixionBizBizModelsPagingRulesPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_biz_biz_models_paging_rules_page()
    // resolves here too.
    public function NoFrixionBizBizModelsPagingRulesPage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_biz_biz_models_paging_rules_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_biz_biz_models_paging_rules_page === null) {
                $this->_no_frixion_biz_biz_models_paging_rules_page = new NoFrixionBizBizModelsPagingRulesPageEntity($this, null);
            }
            return $this->_no_frixion_biz_biz_models_paging_rules_page;
        }
        return new NoFrixionBizBizModelsPagingRulesPageEntity($this, $data);
    }


    private $_no_frixion_biz_biz_models_payments_card_payment = null;

    // Canonical facade: $client->NoFrixionBizBizModelsPaymentsCardPayment()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_biz_biz_models_payments_card_payment()
    // resolves here too.
    public function NoFrixionBizBizModelsPaymentsCardPayment($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_biz_biz_models_payments_card_payment_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_biz_biz_models_payments_card_payment === null) {
                $this->_no_frixion_biz_biz_models_payments_card_payment = new NoFrixionBizBizModelsPaymentsCardPaymentEntity($this, null);
            }
            return $this->_no_frixion_biz_biz_models_payments_card_payment;
        }
        return new NoFrixionBizBizModelsPaymentsCardPaymentEntity($this, $data);
    }


    private $_no_frixion_biz_biz_models_payments_card_public_key = null;

    // Canonical facade: $client->NoFrixionBizBizModelsPaymentsCardPublicKey()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_biz_biz_models_payments_card_public_key()
    // resolves here too.
    public function NoFrixionBizBizModelsPaymentsCardPublicKey($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_biz_biz_models_payments_card_public_key_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_biz_biz_models_payments_card_public_key === null) {
                $this->_no_frixion_biz_biz_models_payments_card_public_key = new NoFrixionBizBizModelsPaymentsCardPublicKeyEntity($this, null);
            }
            return $this->_no_frixion_biz_biz_models_payments_card_public_key;
        }
        return new NoFrixionBizBizModelsPaymentsCardPublicKeyEntity($this, $data);
    }


    private $_no_frixion_money_moov_api_features_beneficiaries_beneficiaries = null;

    // Canonical facade: $client->NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_api_features_beneficiaries_beneficiaries()
    // resolves here too.
    public function NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_api_features_beneficiaries_beneficiaries_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_api_features_beneficiaries_beneficiaries === null) {
                $this->_no_frixion_money_moov_api_features_beneficiaries_beneficiaries = new NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity($this, null);
            }
            return $this->_no_frixion_money_moov_api_features_beneficiaries_beneficiaries;
        }
        return new NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity($this, $data);
    }


    private $_no_frixion_money_moov_api_features_payment_requests_payment = null;

    // Canonical facade: $client->NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_api_features_payment_requests_payment()
    // resolves here too.
    public function NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_api_features_payment_requests_payment_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_api_features_payment_requests_payment === null) {
                $this->_no_frixion_money_moov_api_features_payment_requests_payment = new NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity($this, null);
            }
            return $this->_no_frixion_money_moov_api_features_payment_requests_payment;
        }
        return new NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity($this, $data);
    }


    private $_no_frixion_money_moov_api_features_permissions_roles_create = null;

    // Canonical facade: $client->NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_api_features_permissions_roles_create()
    // resolves here too.
    public function NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_api_features_permissions_roles_create_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_api_features_permissions_roles_create === null) {
                $this->_no_frixion_money_moov_api_features_permissions_roles_create = new NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity($this, null);
            }
            return $this->_no_frixion_money_moov_api_features_permissions_roles_create;
        }
        return new NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity($this, $data);
    }


    private $_no_frixion_money_moov_api_features_user_invites_create = null;

    // Canonical facade: $client->NoFrixionMoneyMoovApiFeaturesUserInvitesCreate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_api_features_user_invites_create()
    // resolves here too.
    public function NoFrixionMoneyMoovApiFeaturesUserInvitesCreate($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_api_features_user_invites_create_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_api_features_user_invites_create === null) {
                $this->_no_frixion_money_moov_api_features_user_invites_create = new NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity($this, null);
            }
            return $this->_no_frixion_money_moov_api_features_user_invites_create;
        }
        return new NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_authorisation_settings_merchant = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_authorisation_settings_merchant()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_authorisation_settings_merchant_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_authorisation_settings_merchant === null) {
                $this->_no_frixion_money_moov_models_authorisation_settings_merchant = new NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_authorisation_settings_merchant;
        }
        return new NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_batch_payout = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsBatchPayout()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_batch_payout()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsBatchPayout($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_batch_payout_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_batch_payout === null) {
                $this->_no_frixion_money_moov_models_batch_payout = new NoFrixionMoneyMoovModelsBatchPayoutEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_batch_payout;
        }
        return new NoFrixionMoneyMoovModelsBatchPayoutEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_beneficiary_group_page = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsBeneficiaryGroupPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_beneficiary_group_page()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsBeneficiaryGroupPage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_beneficiary_group_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_beneficiary_group_page === null) {
                $this->_no_frixion_money_moov_models_beneficiary_group_page = new NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_beneficiary_group_page;
        }
        return new NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_beneficiary_page = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsBeneficiaryPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_beneficiary_page()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsBeneficiaryPage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_beneficiary_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_beneficiary_page === null) {
                $this->_no_frixion_money_moov_models_beneficiary_page = new NoFrixionMoneyMoovModelsBeneficiaryPageEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_beneficiary_page;
        }
        return new NoFrixionMoneyMoovModelsBeneficiaryPageEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_card_customer_token = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsCardCustomerToken()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_card_customer_token()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsCardCustomerToken($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_card_customer_token_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_card_customer_token === null) {
                $this->_no_frixion_money_moov_models_card_customer_token = new NoFrixionMoneyMoovModelsCardCustomerTokenEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_card_customer_token;
        }
        return new NoFrixionMoneyMoovModelsCardCustomerTokenEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_currency_currency_info = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsCurrencyCurrencyInfo()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_currency_currency_info()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsCurrencyCurrencyInfo($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_currency_currency_info_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_currency_currency_info === null) {
                $this->_no_frixion_money_moov_models_currency_currency_info = new NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_currency_currency_info;
        }
        return new NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_direct_debit_batch_submit = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsDirectDebitBatchSubmit()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_direct_debit_batch_submit()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsDirectDebitBatchSubmit($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_direct_debit_batch_submit_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_direct_debit_batch_submit === null) {
                $this->_no_frixion_money_moov_models_direct_debit_batch_submit = new NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_direct_debit_batch_submit;
        }
        return new NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_fx_rate = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsFxRate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_fx_rate()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsFxRate($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_fx_rate_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_fx_rate === null) {
                $this->_no_frixion_money_moov_models_fx_rate = new NoFrixionMoneyMoovModelsFxRateEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_fx_rate;
        }
        return new NoFrixionMoneyMoovModelsFxRateEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_i_payment = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsIPayment()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_i_payment()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsIPayment($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_i_payment_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_i_payment === null) {
                $this->_no_frixion_money_moov_models_i_payment = new NoFrixionMoneyMoovModelsIPaymentEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_i_payment;
        }
        return new NoFrixionMoneyMoovModelsIPaymentEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_mandates_mandate = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsMandatesMandate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_mandates_mandate()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsMandatesMandate($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_mandates_mandate_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_mandates_mandate === null) {
                $this->_no_frixion_money_moov_models_mandates_mandate = new NoFrixionMoneyMoovModelsMandatesMandateEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_mandates_mandate;
        }
        return new NoFrixionMoneyMoovModelsMandatesMandateEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_merchant = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsMerchant()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_merchant()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsMerchant($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_merchant_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_merchant === null) {
                $this->_no_frixion_money_moov_models_merchant = new NoFrixionMoneyMoovModelsMerchantEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_merchant;
        }
        return new NoFrixionMoneyMoovModelsMerchantEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_merchant_page = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsMerchantPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_merchant_page()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsMerchantPage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_merchant_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_merchant_page === null) {
                $this->_no_frixion_money_moov_models_merchant_page = new NoFrixionMoneyMoovModelsMerchantPageEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_merchant_page;
        }
        return new NoFrixionMoneyMoovModelsMerchantPageEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_merchant_pay_by_bank_setting = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsMerchantPayByBankSetting()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_merchant_pay_by_bank_setting()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsMerchantPayByBankSetting($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_merchant_pay_by_bank_setting_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_merchant_pay_by_bank_setting === null) {
                $this->_no_frixion_money_moov_models_merchant_pay_by_bank_setting = new NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_merchant_pay_by_bank_setting;
        }
        return new NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_merchant_token = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsMerchantToken()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_merchant_token()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsMerchantToken($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_merchant_token_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_merchant_token === null) {
                $this->_no_frixion_money_moov_models_merchant_token = new NoFrixionMoneyMoovModelsMerchantTokenEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_merchant_token;
        }
        return new NoFrixionMoneyMoovModelsMerchantTokenEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_merchant_token_page = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsMerchantTokenPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_merchant_token_page()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsMerchantTokenPage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_merchant_token_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_merchant_token_page === null) {
                $this->_no_frixion_money_moov_models_merchant_token_page = new NoFrixionMoneyMoovModelsMerchantTokenPageEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_merchant_token_page;
        }
        return new NoFrixionMoneyMoovModelsMerchantTokenPageEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_no_frixion_version = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsNoFrixionVersion()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_no_frixion_version()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsNoFrixionVersion($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_no_frixion_version_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_no_frixion_version === null) {
                $this->_no_frixion_money_moov_models_no_frixion_version = new NoFrixionMoneyMoovModelsNoFrixionVersionEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_no_frixion_version;
        }
        return new NoFrixionMoneyMoovModelsNoFrixionVersionEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_open_banking_account = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsOpenBankingAccount()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_open_banking_account()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsOpenBankingAccount($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_open_banking_account_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_open_banking_account === null) {
                $this->_no_frixion_money_moov_models_open_banking_account = new NoFrixionMoneyMoovModelsOpenBankingAccountEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_open_banking_account;
        }
        return new NoFrixionMoneyMoovModelsOpenBankingAccountEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_open_banking_consent = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsOpenBankingConsent()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_open_banking_consent()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsOpenBankingConsent($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_open_banking_consent_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_open_banking_consent === null) {
                $this->_no_frixion_money_moov_models_open_banking_consent = new NoFrixionMoneyMoovModelsOpenBankingConsentEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_open_banking_consent;
        }
        return new NoFrixionMoneyMoovModelsOpenBankingConsentEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_open_banking_transaction = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsOpenBankingTransaction()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_open_banking_transaction()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsOpenBankingTransaction($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_open_banking_transaction_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_open_banking_transaction === null) {
                $this->_no_frixion_money_moov_models_open_banking_transaction = new NoFrixionMoneyMoovModelsOpenBankingTransactionEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_open_banking_transaction;
        }
        return new NoFrixionMoneyMoovModelsOpenBankingTransactionEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_payment = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPayment()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payment()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPayment($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payment_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payment === null) {
                $this->_no_frixion_money_moov_models_payment = new NoFrixionMoneyMoovModelsPaymentEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payment;
        }
        return new NoFrixionMoneyMoovModelsPaymentEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_payment_account_minimal_page = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPaymentAccountMinimalPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payment_account_minimal_page()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPaymentAccountMinimalPage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payment_account_minimal_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payment_account_minimal_page === null) {
                $this->_no_frixion_money_moov_models_payment_account_minimal_page = new NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payment_account_minimal_page;
        }
        return new NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_payment_account_page = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPaymentAccountPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payment_account_page()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPaymentAccountPage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payment_account_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payment_account_page === null) {
                $this->_no_frixion_money_moov_models_payment_account_page = new NoFrixionMoneyMoovModelsPaymentAccountPageEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payment_account_page;
        }
        return new NoFrixionMoneyMoovModelsPaymentAccountPageEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_payment_initiation = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPaymentInitiation()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payment_initiation()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPaymentInitiation($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payment_initiation_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payment_initiation === null) {
                $this->_no_frixion_money_moov_models_payment_initiation = new NoFrixionMoneyMoovModelsPaymentInitiationEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payment_initiation;
        }
        return new NoFrixionMoneyMoovModelsPaymentInitiationEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_payment_request_event = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPaymentRequestEvent()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payment_request_event()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPaymentRequestEvent($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payment_request_event_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payment_request_event === null) {
                $this->_no_frixion_money_moov_models_payment_request_event = new NoFrixionMoneyMoovModelsPaymentRequestEventEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payment_request_event;
        }
        return new NoFrixionMoneyMoovModelsPaymentRequestEventEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_payment_request_metric = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPaymentRequestMetric()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payment_request_metric()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPaymentRequestMetric($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payment_request_metric_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payment_request_metric === null) {
                $this->_no_frixion_money_moov_models_payment_request_metric = new NoFrixionMoneyMoovModelsPaymentRequestMetricEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payment_request_metric;
        }
        return new NoFrixionMoneyMoovModelsPaymentRequestMetricEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_payment_request_minimal = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPaymentRequestMinimal()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payment_request_minimal()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPaymentRequestMinimal($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payment_request_minimal_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payment_request_minimal === null) {
                $this->_no_frixion_money_moov_models_payment_request_minimal = new NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payment_request_minimal;
        }
        return new NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_payment_request_result = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPaymentRequestResult()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payment_request_result()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPaymentRequestResult($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payment_request_result_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payment_request_result === null) {
                $this->_no_frixion_money_moov_models_payment_request_result = new NoFrixionMoneyMoovModelsPaymentRequestResultEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payment_request_result;
        }
        return new NoFrixionMoneyMoovModelsPaymentRequestResultEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_payment_requests_merchant_payment = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payment_requests_merchant_payment()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payment_requests_merchant_payment_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payment_requests_merchant_payment === null) {
                $this->_no_frixion_money_moov_models_payment_requests_merchant_payment = new NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payment_requests_merchant_payment;
        }
        return new NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_payment_requests_merchant_payment2 = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payment_requests_merchant_payment2()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payment_requests_merchant_payment2_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payment_requests_merchant_payment2 === null) {
                $this->_no_frixion_money_moov_models_payment_requests_merchant_payment2 = new NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payment_requests_merchant_payment2;
        }
        return new NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity($this, $data);
    }


    private $_no_frixion_money_moov_models_payment_requests_merchant_payment3 = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payment_requests_merchant_payment3()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payment_requests_merchant_payment3_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payment_requests_merchant_payment3 === null) {
                $this->_no_frixion_money_moov_models_payment_requests_merchant_payment3 = new NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payment_requests_merchant_payment3;
        }
        return new NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity($this, $data);
    }


    private $_no_frixion_money_moov_models_payment_requests_merchant_payment4 = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payment_requests_merchant_payment4()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payment_requests_merchant_payment4_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payment_requests_merchant_payment4 === null) {
                $this->_no_frixion_money_moov_models_payment_requests_merchant_payment4 = new NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payment_requests_merchant_payment4;
        }
        return new NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity($this, $data);
    }


    private $_no_frixion_money_moov_models_payout_keyset_page = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPayoutKeysetPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payout_keyset_page()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPayoutKeysetPage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payout_keyset_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payout_keyset_page === null) {
                $this->_no_frixion_money_moov_models_payout_keyset_page = new NoFrixionMoneyMoovModelsPayoutKeysetPageEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payout_keyset_page;
        }
        return new NoFrixionMoneyMoovModelsPayoutKeysetPageEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_payout_metric = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPayoutMetric()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payout_metric()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPayoutMetric($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payout_metric_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payout_metric === null) {
                $this->_no_frixion_money_moov_models_payout_metric = new NoFrixionMoneyMoovModelsPayoutMetricEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payout_metric;
        }
        return new NoFrixionMoneyMoovModelsPayoutMetricEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_payouts_payouts_create = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPayoutsPayoutsCreate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payouts_payouts_create()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPayoutsPayoutsCreate($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payouts_payouts_create_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payouts_payouts_create === null) {
                $this->_no_frixion_money_moov_models_payouts_payouts_create = new NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payouts_payouts_create;
        }
        return new NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_payrun = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsPayrun()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_payrun()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsPayrun($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_payrun_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_payrun === null) {
                $this->_no_frixion_money_moov_models_payrun = new NoFrixionMoneyMoovModelsPayrunEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_payrun;
        }
        return new NoFrixionMoneyMoovModelsPayrunEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_report_result = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsReportResult()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_report_result()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsReportResult($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_report_result_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_report_result === null) {
                $this->_no_frixion_money_moov_models_report_result = new NoFrixionMoneyMoovModelsReportResultEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_report_result;
        }
        return new NoFrixionMoneyMoovModelsReportResultEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_rule = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsRule()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_rule()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsRule($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_rule_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_rule === null) {
                $this->_no_frixion_money_moov_models_rule = new NoFrixionMoneyMoovModelsRuleEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_rule;
        }
        return new NoFrixionMoneyMoovModelsRuleEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_transaction = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsTransaction()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_transaction()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsTransaction($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_transaction_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_transaction === null) {
                $this->_no_frixion_money_moov_models_transaction = new NoFrixionMoneyMoovModelsTransactionEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_transaction;
        }
        return new NoFrixionMoneyMoovModelsTransactionEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_transaction_page = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsTransactionPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_transaction_page()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsTransactionPage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_transaction_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_transaction_page === null) {
                $this->_no_frixion_money_moov_models_transaction_page = new NoFrixionMoneyMoovModelsTransactionPageEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_transaction_page;
        }
        return new NoFrixionMoneyMoovModelsTransactionPageEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_user_invite = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsUserInvite()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_user_invite()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsUserInvite($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_user_invite_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_user_invite === null) {
                $this->_no_frixion_money_moov_models_user_invite = new NoFrixionMoneyMoovModelsUserInviteEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_user_invite;
        }
        return new NoFrixionMoneyMoovModelsUserInviteEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_user_invite_page = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsUserInvitePage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_user_invite_page()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsUserInvitePage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_user_invite_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_user_invite_page === null) {
                $this->_no_frixion_money_moov_models_user_invite_page = new NoFrixionMoneyMoovModelsUserInvitePageEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_user_invite_page;
        }
        return new NoFrixionMoneyMoovModelsUserInvitePageEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_user_page = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsUserPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_user_page()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsUserPage($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_user_page_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_user_page === null) {
                $this->_no_frixion_money_moov_models_user_page = new NoFrixionMoneyMoovModelsUserPageEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_user_page;
        }
        return new NoFrixionMoneyMoovModelsUserPageEntity($this, $data);
    }


    private $_no_frixion_money_moov_models_webhook = null;

    // Canonical facade: $client->NoFrixionMoneyMoovModelsWebhook()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_money_moov_models_webhook()
    // resolves here too.
    public function NoFrixionMoneyMoovModelsWebhook($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_money_moov_models_webhook_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_money_moov_models_webhook === null) {
                $this->_no_frixion_money_moov_models_webhook = new NoFrixionMoneyMoovModelsWebhookEntity($this, null);
            }
            return $this->_no_frixion_money_moov_models_webhook;
        }
        return new NoFrixionMoneyMoovModelsWebhookEntity($this, $data);
    }


    private $_open_banking = null;

    // Canonical facade: $client->OpenBanking()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->open_banking()
    // resolves here too.
    public function OpenBanking($data = null)
    {
        require_once __DIR__ . '/entity/open_banking_entity.php';
        if ($data === null) {
            if ($this->_open_banking === null) {
                $this->_open_banking = new OpenBankingEntity($this, null);
            }
            return $this->_open_banking;
        }
        return new OpenBankingEntity($this, $data);
    }


    private $_payeeverification = null;

    // Canonical facade: $client->Payeeverification()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->payeeverification()
    // resolves here too.
    public function Payeeverification($data = null)
    {
        require_once __DIR__ . '/entity/payeeverification_entity.php';
        if ($data === null) {
            if ($this->_payeeverification === null) {
                $this->_payeeverification = new PayeeverificationEntity($this, null);
            }
            return $this->_payeeverification;
        }
        return new PayeeverificationEntity($this, $data);
    }


    private $_payment_request = null;

    // Canonical facade: $client->PaymentRequest()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->payment_request()
    // resolves here too.
    public function PaymentRequest($data = null)
    {
        require_once __DIR__ . '/entity/payment_request_entity.php';
        if ($data === null) {
            if ($this->_payment_request === null) {
                $this->_payment_request = new PaymentRequestEntity($this, null);
            }
            return $this->_payment_request;
        }
        return new PaymentRequestEntity($this, $data);
    }


    private $_payout = null;

    // Canonical facade: $client->Payout()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->payout()
    // resolves here too.
    public function Payout($data = null)
    {
        require_once __DIR__ . '/entity/payout_entity.php';
        if ($data === null) {
            if ($this->_payout === null) {
                $this->_payout = new PayoutEntity($this, null);
            }
            return $this->_payout;
        }
        return new PayoutEntity($this, $data);
    }


    private $_payrun = null;

    // Canonical facade: $client->Payrun()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->payrun()
    // resolves here too.
    public function Payrun($data = null)
    {
        require_once __DIR__ . '/entity/payrun_entity.php';
        if ($data === null) {
            if ($this->_payrun === null) {
                $this->_payrun = new PayrunEntity($this, null);
            }
            return $this->_payrun;
        }
        return new PayrunEntity($this, $data);
    }


    private $_reject = null;

    // Canonical facade: $client->Reject()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->reject()
    // resolves here too.
    public function Reject($data = null)
    {
        require_once __DIR__ . '/entity/reject_entity.php';
        if ($data === null) {
            if ($this->_reject === null) {
                $this->_reject = new RejectEntity($this, null);
            }
            return $this->_reject;
        }
        return new RejectEntity($this, $data);
    }


    private $_report = null;

    // Canonical facade: $client->Report()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->report()
    // resolves here too.
    public function Report($data = null)
    {
        require_once __DIR__ . '/entity/report_entity.php';
        if ($data === null) {
            if ($this->_report === null) {
                $this->_report = new ReportEntity($this, null);
            }
            return $this->_report;
        }
        return new ReportEntity($this, $data);
    }


    private $_rule = null;

    // Canonical facade: $client->Rule()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->rule()
    // resolves here too.
    public function Rule($data = null)
    {
        require_once __DIR__ . '/entity/rule_entity.php';
        if ($data === null) {
            if ($this->_rule === null) {
                $this->_rule = new RuleEntity($this, null);
            }
            return $this->_rule;
        }
        return new RuleEntity($this, $data);
    }


    private $_send = null;

    // Canonical facade: $client->Send()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->send()
    // resolves here too.
    public function Send($data = null)
    {
        require_once __DIR__ . '/entity/send_entity.php';
        if ($data === null) {
            if ($this->_send === null) {
                $this->_send = new SendEntity($this, null);
            }
            return $this->_send;
        }
        return new SendEntity($this, $data);
    }


    private $_sendbeneficiary = null;

    // Canonical facade: $client->Sendbeneficiary()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->sendbeneficiary()
    // resolves here too.
    public function Sendbeneficiary($data = null)
    {
        require_once __DIR__ . '/entity/sendbeneficiary_entity.php';
        if ($data === null) {
            if ($this->_sendbeneficiary === null) {
                $this->_sendbeneficiary = new SendbeneficiaryEntity($this, null);
            }
            return $this->_sendbeneficiary;
        }
        return new SendbeneficiaryEntity($this, $data);
    }


    private $_tag = null;

    // Canonical facade: $client->Tag()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->tag()
    // resolves here too.
    public function Tag($data = null)
    {
        require_once __DIR__ . '/entity/tag_entity.php';
        if ($data === null) {
            if ($this->_tag === null) {
                $this->_tag = new TagEntity($this, null);
            }
            return $this->_tag;
        }
        return new TagEntity($this, $data);
    }


    private $_token = null;

    // Canonical facade: $client->Token()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->token()
    // resolves here too.
    public function Token($data = null)
    {
        require_once __DIR__ . '/entity/token_entity.php';
        if ($data === null) {
            if ($this->_token === null) {
                $this->_token = new TokenEntity($this, null);
            }
            return $this->_token;
        }
        return new TokenEntity($this, $data);
    }


    private $_transaction = null;

    // Canonical facade: $client->Transaction()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->transaction()
    // resolves here too.
    public function Transaction($data = null)
    {
        require_once __DIR__ . '/entity/transaction_entity.php';
        if ($data === null) {
            if ($this->_transaction === null) {
                $this->_transaction = new TransactionEntity($this, null);
            }
            return $this->_transaction;
        }
        return new TransactionEntity($this, $data);
    }


    private $_user = null;

    // Canonical facade: $client->User()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->user()
    // resolves here too.
    public function User($data = null)
    {
        require_once __DIR__ . '/entity/user_entity.php';
        if ($data === null) {
            if ($this->_user === null) {
                $this->_user = new UserEntity($this, null);
            }
            return $this->_user;
        }
        return new UserEntity($this, $data);
    }


    private $_user_invite = null;

    // Canonical facade: $client->UserInvite()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->user_invite()
    // resolves here too.
    public function UserInvite($data = null)
    {
        require_once __DIR__ . '/entity/user_invite_entity.php';
        if ($data === null) {
            if ($this->_user_invite === null) {
                $this->_user_invite = new UserInviteEntity($this, null);
            }
            return $this->_user_invite;
        }
        return new UserInviteEntity($this, $data);
    }


    private $_virtual = null;

    // Canonical facade: $client->Virtual()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->virtual()
    // resolves here too.
    public function Virtual($data = null)
    {
        require_once __DIR__ . '/entity/virtual_entity.php';
        if ($data === null) {
            if ($this->_virtual === null) {
                $this->_virtual = new VirtualEntity($this, null);
            }
            return $this->_virtual;
        }
        return new VirtualEntity($this, $data);
    }


    private $_webhook = null;

    // Canonical facade: $client->Webhook()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->webhook()
    // resolves here too.
    public function Webhook($data = null)
    {
        require_once __DIR__ . '/entity/webhook_entity.php';
        if ($data === null) {
            if ($this->_webhook === null) {
                $this->_webhook = new WebhookEntity($this, null);
            }
            return $this->_webhook;
        }
        return new WebhookEntity($this, $data);
    }


    private $_whoami = null;

    // Canonical facade: $client->Whoami()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->whoami()
    // resolves here too.
    public function Whoami($data = null)
    {
        require_once __DIR__ . '/entity/whoami_entity.php';
        if ($data === null) {
            if ($this->_whoami === null) {
                $this->_whoami = new WhoamiEntity($this, null);
            }
            return $this->_whoami;
        }
        return new WhoamiEntity($this, $data);
    }


    private $_whoamitrustedapp = null;

    // Canonical facade: $client->Whoamitrustedapp()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->whoamitrustedapp()
    // resolves here too.
    public function Whoamitrustedapp($data = null)
    {
        require_once __DIR__ . '/entity/whoamitrustedapp_entity.php';
        if ($data === null) {
            if ($this->_whoamitrustedapp === null) {
                $this->_whoamitrustedapp = new WhoamitrustedappEntity($this, null);
            }
            return $this->_whoamitrustedapp;
        }
        return new WhoamitrustedappEntity($this, $data);
    }



    public static function test(?array $testopts = null, ?array $sdkopts = null): self
    {
        $sdkopts = $sdkopts ?? [];
        $sdkopts = Struct::clone($sdkopts);
        $sdkopts = is_array($sdkopts) ? $sdkopts : [];

        $testopts = $testopts ?? [];
        $testopts = Struct::clone($testopts);
        $testopts = is_array($testopts) ? $testopts : [];
        $testopts["active"] = true;

        if (!isset($sdkopts["feature"])) {
            $sdkopts["feature"] = [];
        }
        $sdkopts["feature"]["test"] = $testopts;

        $sdk = new NofrixionSDK($sdkopts);
        $sdk->mode = "test";
        return $sdk;
    }
}
