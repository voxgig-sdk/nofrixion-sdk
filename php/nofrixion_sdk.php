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

// Features record diagnostic state on the client as dynamic properties
// (_retry, _cache, _metrics, ...); allow them explicitly (PHP 8.2+
// deprecates implicit dynamic properties).
#[\AllowDynamicProperties]
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

        // Add features in the resolved order (make_options puts an explicit
        // list order first, else defaults to test-first). Ordering matters: the
        // `test` feature installs the base mock transport and the transport
        // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is
        // current, so `test` must be added before them to sit at the base.
        $feature_opts = NofrixionHelpers::to_map(Struct::getprop($this->options, "feature"));
        if ($feature_opts) {
            $featureorder = Struct::getpath($this->options, "__derived__.featureorder");
            if (is_array($featureorder)) {
                foreach ($featureorder as $fname) {
                    $fopts = NofrixionHelpers::to_map($feature_opts[$fname] ?? null);
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


    private $_batch = null;

    // Canonical facade: $client->Batch()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->batch()
    // resolves here too.
    public function Batch($data = null)
    {
        require_once __DIR__ . '/entity/batch_entity.php';
        if ($data === null) {
            if ($this->_batch === null) {
                $this->_batch = new BatchEntity($this, null);
            }
            return $this->_batch;
        }
        return new BatchEntity($this, $data);
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


    private $_beneficiary_group = null;

    // Canonical facade: $client->BeneficiaryGroup()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->beneficiary_group()
    // resolves here too.
    public function BeneficiaryGroup($data = null)
    {
        require_once __DIR__ . '/entity/beneficiary_group_entity.php';
        if ($data === null) {
            if ($this->_beneficiary_group === null) {
                $this->_beneficiary_group = new BeneficiaryGroupEntity($this, null);
            }
            return $this->_beneficiary_group;
        }
        return new BeneficiaryGroupEntity($this, $data);
    }


    private $_card = null;

    // Canonical facade: $client->Card()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->card()
    // resolves here too.
    public function Card($data = null)
    {
        require_once __DIR__ . '/entity/card_entity.php';
        if ($data === null) {
            if ($this->_card === null) {
                $this->_card = new CardEntity($this, null);
            }
            return $this->_card;
        }
        return new CardEntity($this, $data);
    }


    private $_card_customer_token = null;

    // Canonical facade: $client->CardCustomerToken()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->card_customer_token()
    // resolves here too.
    public function CardCustomerToken($data = null)
    {
        require_once __DIR__ . '/entity/card_customer_token_entity.php';
        if ($data === null) {
            if ($this->_card_customer_token === null) {
                $this->_card_customer_token = new CardCustomerTokenEntity($this, null);
            }
            return $this->_card_customer_token;
        }
        return new CardCustomerTokenEntity($this, $data);
    }


    private $_card_payment = null;

    // Canonical facade: $client->CardPayment()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->card_payment()
    // resolves here too.
    public function CardPayment($data = null)
    {
        require_once __DIR__ . '/entity/card_payment_entity.php';
        if ($data === null) {
            if ($this->_card_payment === null) {
                $this->_card_payment = new CardPaymentEntity($this, null);
            }
            return $this->_card_payment;
        }
        return new CardPaymentEntity($this, $data);
    }


    private $_card_public_key = null;

    // Canonical facade: $client->CardPublicKey()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->card_public_key()
    // resolves here too.
    public function CardPublicKey($data = null)
    {
        require_once __DIR__ . '/entity/card_public_key_entity.php';
        if ($data === null) {
            if ($this->_card_public_key === null) {
                $this->_card_public_key = new CardPublicKeyEntity($this, null);
            }
            return $this->_card_public_key;
        }
        return new CardPublicKeyEntity($this, $data);
    }


    private $_consent = null;

    // Canonical facade: $client->Consent()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->consent()
    // resolves here too.
    public function Consent($data = null)
    {
        require_once __DIR__ . '/entity/consent_entity.php';
        if ($data === null) {
            if ($this->_consent === null) {
                $this->_consent = new ConsentEntity($this, null);
            }
            return $this->_consent;
        }
        return new ConsentEntity($this, $data);
    }


    private $_currency = null;

    // Canonical facade: $client->Currency()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->currency()
    // resolves here too.
    public function Currency($data = null)
    {
        require_once __DIR__ . '/entity/currency_entity.php';
        if ($data === null) {
            if ($this->_currency === null) {
                $this->_currency = new CurrencyEntity($this, null);
            }
            return $this->_currency;
        }
        return new CurrencyEntity($this, $data);
    }


    private $_direct_debit_batch_submit = null;

    // Canonical facade: $client->DirectDebitBatchSubmit()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->direct_debit_batch_submit()
    // resolves here too.
    public function DirectDebitBatchSubmit($data = null)
    {
        require_once __DIR__ . '/entity/direct_debit_batch_submit_entity.php';
        if ($data === null) {
            if ($this->_direct_debit_batch_submit === null) {
                $this->_direct_debit_batch_submit = new DirectDebitBatchSubmitEntity($this, null);
            }
            return $this->_direct_debit_batch_submit;
        }
        return new DirectDebitBatchSubmitEntity($this, $data);
    }


    private $_fx_rate = null;

    // Canonical facade: $client->FxRate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->fx_rate()
    // resolves here too.
    public function FxRate($data = null)
    {
        require_once __DIR__ . '/entity/fx_rate_entity.php';
        if ($data === null) {
            if ($this->_fx_rate === null) {
                $this->_fx_rate = new FxRateEntity($this, null);
            }
            return $this->_fx_rate;
        }
        return new FxRateEntity($this, $data);
    }


    private $_i_payment = null;

    // Canonical facade: $client->IPayment()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->i_payment()
    // resolves here too.
    public function IPayment($data = null)
    {
        require_once __DIR__ . '/entity/i_payment_entity.php';
        if ($data === null) {
            if ($this->_i_payment === null) {
                $this->_i_payment = new IPaymentEntity($this, null);
            }
            return $this->_i_payment;
        }
        return new IPaymentEntity($this, $data);
    }


    private $_mandate = null;

    // Canonical facade: $client->Mandate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->mandate()
    // resolves here too.
    public function Mandate($data = null)
    {
        require_once __DIR__ . '/entity/mandate_entity.php';
        if ($data === null) {
            if ($this->_mandate === null) {
                $this->_mandate = new MandateEntity($this, null);
            }
            return $this->_mandate;
        }
        return new MandateEntity($this, $data);
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


    private $_merchant_authorisation_setting = null;

    // Canonical facade: $client->MerchantAuthorisationSetting()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->merchant_authorisation_setting()
    // resolves here too.
    public function MerchantAuthorisationSetting($data = null)
    {
        require_once __DIR__ . '/entity/merchant_authorisation_setting_entity.php';
        if ($data === null) {
            if ($this->_merchant_authorisation_setting === null) {
                $this->_merchant_authorisation_setting = new MerchantAuthorisationSettingEntity($this, null);
            }
            return $this->_merchant_authorisation_setting;
        }
        return new MerchantAuthorisationSettingEntity($this, $data);
    }


    private $_merchant_direct_debit_mandate_page = null;

    // Canonical facade: $client->MerchantDirectDebitMandatePage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->merchant_direct_debit_mandate_page()
    // resolves here too.
    public function MerchantDirectDebitMandatePage($data = null)
    {
        require_once __DIR__ . '/entity/merchant_direct_debit_mandate_page_entity.php';
        if ($data === null) {
            if ($this->_merchant_direct_debit_mandate_page === null) {
                $this->_merchant_direct_debit_mandate_page = new MerchantDirectDebitMandatePageEntity($this, null);
            }
            return $this->_merchant_direct_debit_mandate_page;
        }
        return new MerchantDirectDebitMandatePageEntity($this, $data);
    }


    private $_merchant_pay_by_bank_setting = null;

    // Canonical facade: $client->MerchantPayByBankSetting()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->merchant_pay_by_bank_setting()
    // resolves here too.
    public function MerchantPayByBankSetting($data = null)
    {
        require_once __DIR__ . '/entity/merchant_pay_by_bank_setting_entity.php';
        if ($data === null) {
            if ($this->_merchant_pay_by_bank_setting === null) {
                $this->_merchant_pay_by_bank_setting = new MerchantPayByBankSettingEntity($this, null);
            }
            return $this->_merchant_pay_by_bank_setting;
        }
        return new MerchantPayByBankSettingEntity($this, $data);
    }


    private $_merchant_payment_request_template = null;

    // Canonical facade: $client->MerchantPaymentRequestTemplate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->merchant_payment_request_template()
    // resolves here too.
    public function MerchantPaymentRequestTemplate($data = null)
    {
        require_once __DIR__ . '/entity/merchant_payment_request_template_entity.php';
        if ($data === null) {
            if ($this->_merchant_payment_request_template === null) {
                $this->_merchant_payment_request_template = new MerchantPaymentRequestTemplateEntity($this, null);
            }
            return $this->_merchant_payment_request_template;
        }
        return new MerchantPaymentRequestTemplateEntity($this, $data);
    }


    private $_merchant_token = null;

    // Canonical facade: $client->MerchantToken()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->merchant_token()
    // resolves here too.
    public function MerchantToken($data = null)
    {
        require_once __DIR__ . '/entity/merchant_token_entity.php';
        if ($data === null) {
            if ($this->_merchant_token === null) {
                $this->_merchant_token = new MerchantTokenEntity($this, null);
            }
            return $this->_merchant_token;
        }
        return new MerchantTokenEntity($this, $data);
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


    private $_no_frixion_version = null;

    // Canonical facade: $client->NoFrixionVersion()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->no_frixion_version()
    // resolves here too.
    public function NoFrixionVersion($data = null)
    {
        require_once __DIR__ . '/entity/no_frixion_version_entity.php';
        if ($data === null) {
            if ($this->_no_frixion_version === null) {
                $this->_no_frixion_version = new NoFrixionVersionEntity($this, null);
            }
            return $this->_no_frixion_version;
        }
        return new NoFrixionVersionEntity($this, $data);
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


    private $_payment = null;

    // Canonical facade: $client->Payment()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->payment()
    // resolves here too.
    public function Payment($data = null)
    {
        require_once __DIR__ . '/entity/payment_entity.php';
        if ($data === null) {
            if ($this->_payment === null) {
                $this->_payment = new PaymentEntity($this, null);
            }
            return $this->_payment;
        }
        return new PaymentEntity($this, $data);
    }


    private $_payment_account = null;

    // Canonical facade: $client->PaymentAccount()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->payment_account()
    // resolves here too.
    public function PaymentAccount($data = null)
    {
        require_once __DIR__ . '/entity/payment_account_entity.php';
        if ($data === null) {
            if ($this->_payment_account === null) {
                $this->_payment_account = new PaymentAccountEntity($this, null);
            }
            return $this->_payment_account;
        }
        return new PaymentAccountEntity($this, $data);
    }


    private $_payment_account_minimal = null;

    // Canonical facade: $client->PaymentAccountMinimal()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->payment_account_minimal()
    // resolves here too.
    public function PaymentAccountMinimal($data = null)
    {
        require_once __DIR__ . '/entity/payment_account_minimal_entity.php';
        if ($data === null) {
            if ($this->_payment_account_minimal === null) {
                $this->_payment_account_minimal = new PaymentAccountMinimalEntity($this, null);
            }
            return $this->_payment_account_minimal;
        }
        return new PaymentAccountMinimalEntity($this, $data);
    }


    private $_payment_initiation = null;

    // Canonical facade: $client->PaymentInitiation()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->payment_initiation()
    // resolves here too.
    public function PaymentInitiation($data = null)
    {
        require_once __DIR__ . '/entity/payment_initiation_entity.php';
        if ($data === null) {
            if ($this->_payment_initiation === null) {
                $this->_payment_initiation = new PaymentInitiationEntity($this, null);
            }
            return $this->_payment_initiation;
        }
        return new PaymentInitiationEntity($this, $data);
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


    private $_payment_request_event = null;

    // Canonical facade: $client->PaymentRequestEvent()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->payment_request_event()
    // resolves here too.
    public function PaymentRequestEvent($data = null)
    {
        require_once __DIR__ . '/entity/payment_request_event_entity.php';
        if ($data === null) {
            if ($this->_payment_request_event === null) {
                $this->_payment_request_event = new PaymentRequestEventEntity($this, null);
            }
            return $this->_payment_request_event;
        }
        return new PaymentRequestEventEntity($this, $data);
    }


    private $_payment_request_metric = null;

    // Canonical facade: $client->PaymentRequestMetric()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->payment_request_metric()
    // resolves here too.
    public function PaymentRequestMetric($data = null)
    {
        require_once __DIR__ . '/entity/payment_request_metric_entity.php';
        if ($data === null) {
            if ($this->_payment_request_metric === null) {
                $this->_payment_request_metric = new PaymentRequestMetricEntity($this, null);
            }
            return $this->_payment_request_metric;
        }
        return new PaymentRequestMetricEntity($this, $data);
    }


    private $_payment_request_minimal = null;

    // Canonical facade: $client->PaymentRequestMinimal()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->payment_request_minimal()
    // resolves here too.
    public function PaymentRequestMinimal($data = null)
    {
        require_once __DIR__ . '/entity/payment_request_minimal_entity.php';
        if ($data === null) {
            if ($this->_payment_request_minimal === null) {
                $this->_payment_request_minimal = new PaymentRequestMinimalEntity($this, null);
            }
            return $this->_payment_request_minimal;
        }
        return new PaymentRequestMinimalEntity($this, $data);
    }


    private $_payment_request_result = null;

    // Canonical facade: $client->PaymentRequestResult()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->payment_request_result()
    // resolves here too.
    public function PaymentRequestResult($data = null)
    {
        require_once __DIR__ . '/entity/payment_request_result_entity.php';
        if ($data === null) {
            if ($this->_payment_request_result === null) {
                $this->_payment_request_result = new PaymentRequestResultEntity($this, null);
            }
            return $this->_payment_request_result;
        }
        return new PaymentRequestResultEntity($this, $data);
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


    private $_payout_keyset_page = null;

    // Canonical facade: $client->PayoutKeysetPage()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->payout_keyset_page()
    // resolves here too.
    public function PayoutKeysetPage($data = null)
    {
        require_once __DIR__ . '/entity/payout_keyset_page_entity.php';
        if ($data === null) {
            if ($this->_payout_keyset_page === null) {
                $this->_payout_keyset_page = new PayoutKeysetPageEntity($this, null);
            }
            return $this->_payout_keyset_page;
        }
        return new PayoutKeysetPageEntity($this, $data);
    }


    private $_payout_metric = null;

    // Canonical facade: $client->PayoutMetric()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->payout_metric()
    // resolves here too.
    public function PayoutMetric($data = null)
    {
        require_once __DIR__ . '/entity/payout_metric_entity.php';
        if ($data === null) {
            if ($this->_payout_metric === null) {
                $this->_payout_metric = new PayoutMetricEntity($this, null);
            }
            return $this->_payout_metric;
        }
        return new PayoutMetricEntity($this, $data);
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


    private $_report_result = null;

    // Canonical facade: $client->ReportResult()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->report_result()
    // resolves here too.
    public function ReportResult($data = null)
    {
        require_once __DIR__ . '/entity/report_result_entity.php';
        if ($data === null) {
            if ($this->_report_result === null) {
                $this->_report_result = new ReportResultEntity($this, null);
            }
            return $this->_report_result;
        }
        return new ReportResultEntity($this, $data);
    }


    private $_role = null;

    // Canonical facade: $client->Role()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->role()
    // resolves here too.
    public function Role($data = null)
    {
        require_once __DIR__ . '/entity/role_entity.php';
        if ($data === null) {
            if ($this->_role === null) {
                $this->_role = new RoleEntity($this, null);
            }
            return $this->_role;
        }
        return new RoleEntity($this, $data);
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


    private $_rule_event = null;

    // Canonical facade: $client->RuleEvent()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->rule_event()
    // resolves here too.
    public function RuleEvent($data = null)
    {
        require_once __DIR__ . '/entity/rule_event_entity.php';
        if ($data === null) {
            if ($this->_rule_event === null) {
                $this->_rule_event = new RuleEventEntity($this, null);
            }
            return $this->_rule_event;
        }
        return new RuleEventEntity($this, $data);
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
