# Nofrixion SDK

from utility.voxgig_struct import voxgig_struct as vs
from core.utility_type import NofrixionUtility
from core.spec import NofrixionSpec
from core import helpers

# Load utility registration (populates Utility._registrar)
from utility import register

# Load features
from feature.base_feature import NofrixionBaseFeature
from features import _make_feature


class NofrixionSDK:

    def __init__(self, options=None):
        self.mode = "live"
        self.features = []
        self.options = None

        utility = NofrixionUtility()
        self._utility = utility

        from config import make_config
        config = make_config()

        self._rootctx = utility.make_context({
            "client": self,
            "utility": utility,
            "config": config,
            "options": options if options is not None else {},
            "shared": {},
        }, None)

        self.options = utility.make_options(self._rootctx)

        if vs.getpath(self.options, "feature.test.active") is True:
            self.mode = "test"

        self._rootctx.options = self.options

        # Add features from config.
        feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
        if feature_opts is not None:
            feature_items = vs.items(feature_opts)
            if feature_items is not None:
                for item in feature_items:
                    fname = item[0]
                    fopts = helpers.to_map(item[1])
                    if fopts is not None and fopts.get("active") is True:
                        utility.feature_add(self._rootctx, _make_feature(fname))

        # Add extension features.
        extend = vs.getprop(self.options, "extend")
        if isinstance(extend, list):
            for f in extend:
                if isinstance(f, dict) or (hasattr(f, "get_name") and callable(f.get_name)):
                    utility.feature_add(self._rootctx, f)

        # Initialize features.
        for f in self.features:
            utility.feature_init(self._rootctx, f)

        utility.feature_hook(self._rootctx, "PostConstruct")

        # #BuildFeatures

    def options_map(self):
        out = vs.clone(self.options)
        if isinstance(out, dict):
            return out
        return {}

    def get_utility(self):
        return NofrixionUtility.copy(self._utility)

    def get_root_ctx(self):
        return self._rootctx

    def prepare(self, fetchargs=None):
        utility = self._utility

        if fetchargs is None:
            fetchargs = {}

        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "prepare",
            "ctrl": ctrl,
        }, self._rootctx)

        options = self.options

        path = vs.getprop(fetchargs, "path") or ""
        if not isinstance(path, str):
            path = ""

        method = vs.getprop(fetchargs, "method") or "GET"
        if not isinstance(method, str):
            method = "GET"

        params = helpers.to_map(vs.getprop(fetchargs, "params"))
        if params is None:
            params = {}
        query = helpers.to_map(vs.getprop(fetchargs, "query"))
        if query is None:
            query = {}

        headers = utility.prepare_headers(ctx)

        base = vs.getprop(options, "base") or ""
        if not isinstance(base, str):
            base = ""
        prefix = vs.getprop(options, "prefix") or ""
        if not isinstance(prefix, str):
            prefix = ""
        suffix = vs.getprop(options, "suffix") or ""
        if not isinstance(suffix, str):
            suffix = ""

        ctx.spec = NofrixionSpec({
            "base": base,
            "prefix": prefix,
            "suffix": suffix,
            "path": path,
            "method": method,
            "params": params,
            "query": query,
            "headers": headers,
            "body": vs.getprop(fetchargs, "body"),
            "step": "start",
        })

        # Merge user-provided headers.
        uh = vs.getprop(fetchargs, "headers")
        if isinstance(uh, dict):
            for k, v in uh.items():
                ctx.spec.headers[k] = v

        _, err = utility.prepare_auth(ctx)
        if err is not None:
            raise err

        fetchdef, err = utility.make_fetch_def(ctx)
        if err is not None:
            raise err

        return fetchdef

    def direct(self, fetchargs=None):
        utility = self._utility

        try:
            fetchdef = self.prepare(fetchargs)
        except Exception as err:
            # direct() is the raw-HTTP escape hatch: it never raises, it
            # returns a result object callers branch on via result["ok"].
            return {"ok": False, "err": err}

        if fetchargs is None:
            fetchargs = {}
        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "direct",
            "ctrl": ctrl,
        }, self._rootctx)

        url = fetchdef.get("url", "")
        fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

        if fetch_err is not None:
            return {"ok": False, "err": fetch_err}

        if fetched is None:
            return {
                "ok": False,
                "err": ctx.make_error("direct_no_response", "response: undefined"),
            }

        if isinstance(fetched, dict):
            status = helpers.to_int(vs.getprop(fetched, "status"))
            headers = vs.getprop(fetched, "headers") or {}

            # No-body responses (204, 304) and explicit zero content-length
            # must skip JSON parsing — calling json() on an empty body raises.
            content_length = None
            if isinstance(headers, dict):
                content_length = headers.get("content-length")
            no_body = status in (204, 304) or str(content_length) == "0"

            json_data = None
            if not no_body:
                jf = vs.getprop(fetched, "json")
                if callable(jf):
                    try:
                        json_data = jf()
                    except Exception:
                        # Non-JSON body (e.g. text/plain, text/html). Surface
                        # status + headers but leave data as None.
                        json_data = None

            return {
                "ok": status >= 200 and status < 300,
                "status": status,
                "headers": headers,
                "data": json_data,
            }

        return {
            "ok": False,
            "err": ctx.make_error("direct_invalid", "invalid response type"),
        }


    def Account(self, data=None) -> "AccountEntity":
        """Entity factory: client.Account().list() / client.Account().load({"id": ...})."""
        from entity.account_entity import AccountEntity
        return AccountEntity(self, data)


    def Beneficiary(self, data=None) -> "BeneficiaryEntity":
        """Entity factory: client.Beneficiary().list() / client.Beneficiary().load({"id": ...})."""
        from entity.beneficiary_entity import BeneficiaryEntity
        return BeneficiaryEntity(self, data)


    def Cancel(self, data=None) -> "CancelEntity":
        """Entity factory: client.Cancel().list() / client.Cancel().load({"id": ...})."""
        from entity.cancel_entity import CancelEntity
        return CancelEntity(self, data)


    def Disable(self, data=None) -> "DisableEntity":
        """Entity factory: client.Disable().list() / client.Disable().load({"id": ...})."""
        from entity.disable_entity import DisableEntity
        return DisableEntity(self, data)


    def Enable(self, data=None) -> "EnableEntity":
        """Entity factory: client.Enable().list() / client.Enable().load({"id": ...})."""
        from entity.enable_entity import EnableEntity
        return EnableEntity(self, data)


    def Merchant(self, data=None) -> "MerchantEntity":
        """Entity factory: client.Merchant().list() / client.Merchant().load({"id": ...})."""
        from entity.merchant_entity import MerchantEntity
        return MerchantEntity(self, data)


    def Metadata(self, data=None) -> "MetadataEntity":
        """Entity factory: client.Metadata().list() / client.Metadata().load({"id": ...})."""
        from entity.metadata_entity import MetadataEntity
        return MetadataEntity(self, data)


    def NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(self, data=None) -> "NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity":
        """Entity factory: client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage().list() / client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage().load({"id": ...})."""
        from entity.no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page_entity import NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity
        return NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity(self, data)


    def NoFrixionBizBizModelsPagingPaymentRequestPage(self, data=None) -> "NoFrixionBizBizModelsPagingPaymentRequestPageEntity":
        """Entity factory: client.NoFrixionBizBizModelsPagingPaymentRequestPage().list() / client.NoFrixionBizBizModelsPagingPaymentRequestPage().load({"id": ...})."""
        from entity.no_frixion_biz_biz_models_paging_payment_request_page_entity import NoFrixionBizBizModelsPagingPaymentRequestPageEntity
        return NoFrixionBizBizModelsPagingPaymentRequestPageEntity(self, data)


    def NoFrixionBizBizModelsPagingPayoutPage(self, data=None) -> "NoFrixionBizBizModelsPagingPayoutPageEntity":
        """Entity factory: client.NoFrixionBizBizModelsPagingPayoutPage().list() / client.NoFrixionBizBizModelsPagingPayoutPage().load({"id": ...})."""
        from entity.no_frixion_biz_biz_models_paging_payout_page_entity import NoFrixionBizBizModelsPagingPayoutPageEntity
        return NoFrixionBizBizModelsPagingPayoutPageEntity(self, data)


    def NoFrixionBizBizModelsPagingPayrunPage(self, data=None) -> "NoFrixionBizBizModelsPagingPayrunPageEntity":
        """Entity factory: client.NoFrixionBizBizModelsPagingPayrunPage().list() / client.NoFrixionBizBizModelsPagingPayrunPage().load({"id": ...})."""
        from entity.no_frixion_biz_biz_models_paging_payrun_page_entity import NoFrixionBizBizModelsPagingPayrunPageEntity
        return NoFrixionBizBizModelsPagingPayrunPageEntity(self, data)


    def NoFrixionBizBizModelsPagingRuleEventsPage(self, data=None) -> "NoFrixionBizBizModelsPagingRuleEventsPageEntity":
        """Entity factory: client.NoFrixionBizBizModelsPagingRuleEventsPage().list() / client.NoFrixionBizBizModelsPagingRuleEventsPage().load({"id": ...})."""
        from entity.no_frixion_biz_biz_models_paging_rule_events_page_entity import NoFrixionBizBizModelsPagingRuleEventsPageEntity
        return NoFrixionBizBizModelsPagingRuleEventsPageEntity(self, data)


    def NoFrixionBizBizModelsPagingRulesPage(self, data=None) -> "NoFrixionBizBizModelsPagingRulesPageEntity":
        """Entity factory: client.NoFrixionBizBizModelsPagingRulesPage().list() / client.NoFrixionBizBizModelsPagingRulesPage().load({"id": ...})."""
        from entity.no_frixion_biz_biz_models_paging_rules_page_entity import NoFrixionBizBizModelsPagingRulesPageEntity
        return NoFrixionBizBizModelsPagingRulesPageEntity(self, data)


    def NoFrixionBizBizModelsPaymentsCardPayment(self, data=None) -> "NoFrixionBizBizModelsPaymentsCardPaymentEntity":
        """Entity factory: client.NoFrixionBizBizModelsPaymentsCardPayment().list() / client.NoFrixionBizBizModelsPaymentsCardPayment().load({"id": ...})."""
        from entity.no_frixion_biz_biz_models_payments_card_payment_entity import NoFrixionBizBizModelsPaymentsCardPaymentEntity
        return NoFrixionBizBizModelsPaymentsCardPaymentEntity(self, data)


    def NoFrixionBizBizModelsPaymentsCardPublicKey(self, data=None) -> "NoFrixionBizBizModelsPaymentsCardPublicKeyEntity":
        """Entity factory: client.NoFrixionBizBizModelsPaymentsCardPublicKey().list() / client.NoFrixionBizBizModelsPaymentsCardPublicKey().load({"id": ...})."""
        from entity.no_frixion_biz_biz_models_payments_card_public_key_entity import NoFrixionBizBizModelsPaymentsCardPublicKeyEntity
        return NoFrixionBizBizModelsPaymentsCardPublicKeyEntity(self, data)


    def NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(self, data=None) -> "NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity":
        """Entity factory: client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries().list() / client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries().load({"id": ...})."""
        from entity.no_frixion_money_moov_api_features_beneficiaries_beneficiaries_entity import NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity
        return NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity(self, data)


    def NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(self, data=None) -> "NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity":
        """Entity factory: client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment().list() / client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment().load({"id": ...})."""
        from entity.no_frixion_money_moov_api_features_payment_requests_payment_entity import NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity
        return NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity(self, data)


    def NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(self, data=None) -> "NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity":
        """Entity factory: client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate().list() / client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate().load({"id": ...})."""
        from entity.no_frixion_money_moov_api_features_permissions_roles_create_entity import NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity
        return NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity(self, data)


    def NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(self, data=None) -> "NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity":
        """Entity factory: client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate().list() / client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate().load({"id": ...})."""
        from entity.no_frixion_money_moov_api_features_user_invites_create_entity import NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity
        return NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity(self, data)


    def NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(self, data=None) -> "NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant().list() / client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_authorisation_settings_merchant_entity import NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity
        return NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity(self, data)


    def NoFrixionMoneyMoovModelsBatchPayout(self, data=None) -> "NoFrixionMoneyMoovModelsBatchPayoutEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsBatchPayout().list() / client.NoFrixionMoneyMoovModelsBatchPayout().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_batch_payout_entity import NoFrixionMoneyMoovModelsBatchPayoutEntity
        return NoFrixionMoneyMoovModelsBatchPayoutEntity(self, data)


    def NoFrixionMoneyMoovModelsBeneficiaryGroupPage(self, data=None) -> "NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage().list() / client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_beneficiary_group_page_entity import NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity
        return NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity(self, data)


    def NoFrixionMoneyMoovModelsBeneficiaryPage(self, data=None) -> "NoFrixionMoneyMoovModelsBeneficiaryPageEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsBeneficiaryPage().list() / client.NoFrixionMoneyMoovModelsBeneficiaryPage().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_beneficiary_page_entity import NoFrixionMoneyMoovModelsBeneficiaryPageEntity
        return NoFrixionMoneyMoovModelsBeneficiaryPageEntity(self, data)


    def NoFrixionMoneyMoovModelsCardCustomerToken(self, data=None) -> "NoFrixionMoneyMoovModelsCardCustomerTokenEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsCardCustomerToken().list() / client.NoFrixionMoneyMoovModelsCardCustomerToken().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_card_customer_token_entity import NoFrixionMoneyMoovModelsCardCustomerTokenEntity
        return NoFrixionMoneyMoovModelsCardCustomerTokenEntity(self, data)


    def NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(self, data=None) -> "NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo().list() / client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_currency_currency_info_entity import NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity
        return NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity(self, data)


    def NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(self, data=None) -> "NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit().list() / client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_direct_debit_batch_submit_entity import NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity
        return NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity(self, data)


    def NoFrixionMoneyMoovModelsFxRate(self, data=None) -> "NoFrixionMoneyMoovModelsFxRateEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsFxRate().list() / client.NoFrixionMoneyMoovModelsFxRate().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_fx_rate_entity import NoFrixionMoneyMoovModelsFxRateEntity
        return NoFrixionMoneyMoovModelsFxRateEntity(self, data)


    def NoFrixionMoneyMoovModelsIPayment(self, data=None) -> "NoFrixionMoneyMoovModelsIPaymentEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsIPayment().list() / client.NoFrixionMoneyMoovModelsIPayment().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_i_payment_entity import NoFrixionMoneyMoovModelsIPaymentEntity
        return NoFrixionMoneyMoovModelsIPaymentEntity(self, data)


    def NoFrixionMoneyMoovModelsMandatesMandate(self, data=None) -> "NoFrixionMoneyMoovModelsMandatesMandateEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsMandatesMandate().list() / client.NoFrixionMoneyMoovModelsMandatesMandate().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_mandates_mandate_entity import NoFrixionMoneyMoovModelsMandatesMandateEntity
        return NoFrixionMoneyMoovModelsMandatesMandateEntity(self, data)


    def NoFrixionMoneyMoovModelsMerchant(self, data=None) -> "NoFrixionMoneyMoovModelsMerchantEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsMerchant().list() / client.NoFrixionMoneyMoovModelsMerchant().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_merchant_entity import NoFrixionMoneyMoovModelsMerchantEntity
        return NoFrixionMoneyMoovModelsMerchantEntity(self, data)


    def NoFrixionMoneyMoovModelsMerchantPage(self, data=None) -> "NoFrixionMoneyMoovModelsMerchantPageEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsMerchantPage().list() / client.NoFrixionMoneyMoovModelsMerchantPage().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_merchant_page_entity import NoFrixionMoneyMoovModelsMerchantPageEntity
        return NoFrixionMoneyMoovModelsMerchantPageEntity(self, data)


    def NoFrixionMoneyMoovModelsMerchantPayByBankSetting(self, data=None) -> "NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting().list() / client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_merchant_pay_by_bank_setting_entity import NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity
        return NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity(self, data)


    def NoFrixionMoneyMoovModelsMerchantToken(self, data=None) -> "NoFrixionMoneyMoovModelsMerchantTokenEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsMerchantToken().list() / client.NoFrixionMoneyMoovModelsMerchantToken().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_merchant_token_entity import NoFrixionMoneyMoovModelsMerchantTokenEntity
        return NoFrixionMoneyMoovModelsMerchantTokenEntity(self, data)


    def NoFrixionMoneyMoovModelsMerchantTokenPage(self, data=None) -> "NoFrixionMoneyMoovModelsMerchantTokenPageEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsMerchantTokenPage().list() / client.NoFrixionMoneyMoovModelsMerchantTokenPage().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_merchant_token_page_entity import NoFrixionMoneyMoovModelsMerchantTokenPageEntity
        return NoFrixionMoneyMoovModelsMerchantTokenPageEntity(self, data)


    def NoFrixionMoneyMoovModelsNoFrixionVersion(self, data=None) -> "NoFrixionMoneyMoovModelsNoFrixionVersionEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsNoFrixionVersion().list() / client.NoFrixionMoneyMoovModelsNoFrixionVersion().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_no_frixion_version_entity import NoFrixionMoneyMoovModelsNoFrixionVersionEntity
        return NoFrixionMoneyMoovModelsNoFrixionVersionEntity(self, data)


    def NoFrixionMoneyMoovModelsOpenBankingAccount(self, data=None) -> "NoFrixionMoneyMoovModelsOpenBankingAccountEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsOpenBankingAccount().list() / client.NoFrixionMoneyMoovModelsOpenBankingAccount().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_open_banking_account_entity import NoFrixionMoneyMoovModelsOpenBankingAccountEntity
        return NoFrixionMoneyMoovModelsOpenBankingAccountEntity(self, data)


    def NoFrixionMoneyMoovModelsOpenBankingConsent(self, data=None) -> "NoFrixionMoneyMoovModelsOpenBankingConsentEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsOpenBankingConsent().list() / client.NoFrixionMoneyMoovModelsOpenBankingConsent().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_open_banking_consent_entity import NoFrixionMoneyMoovModelsOpenBankingConsentEntity
        return NoFrixionMoneyMoovModelsOpenBankingConsentEntity(self, data)


    def NoFrixionMoneyMoovModelsOpenBankingTransaction(self, data=None) -> "NoFrixionMoneyMoovModelsOpenBankingTransactionEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsOpenBankingTransaction().list() / client.NoFrixionMoneyMoovModelsOpenBankingTransaction().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_open_banking_transaction_entity import NoFrixionMoneyMoovModelsOpenBankingTransactionEntity
        return NoFrixionMoneyMoovModelsOpenBankingTransactionEntity(self, data)


    def NoFrixionMoneyMoovModelsPayment(self, data=None) -> "NoFrixionMoneyMoovModelsPaymentEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPayment().list() / client.NoFrixionMoneyMoovModelsPayment().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payment_entity import NoFrixionMoneyMoovModelsPaymentEntity
        return NoFrixionMoneyMoovModelsPaymentEntity(self, data)


    def NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(self, data=None) -> "NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage().list() / client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payment_account_minimal_page_entity import NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity
        return NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity(self, data)


    def NoFrixionMoneyMoovModelsPaymentAccountPage(self, data=None) -> "NoFrixionMoneyMoovModelsPaymentAccountPageEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPaymentAccountPage().list() / client.NoFrixionMoneyMoovModelsPaymentAccountPage().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payment_account_page_entity import NoFrixionMoneyMoovModelsPaymentAccountPageEntity
        return NoFrixionMoneyMoovModelsPaymentAccountPageEntity(self, data)


    def NoFrixionMoneyMoovModelsPaymentInitiation(self, data=None) -> "NoFrixionMoneyMoovModelsPaymentInitiationEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPaymentInitiation().list() / client.NoFrixionMoneyMoovModelsPaymentInitiation().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payment_initiation_entity import NoFrixionMoneyMoovModelsPaymentInitiationEntity
        return NoFrixionMoneyMoovModelsPaymentInitiationEntity(self, data)


    def NoFrixionMoneyMoovModelsPaymentRequestEvent(self, data=None) -> "NoFrixionMoneyMoovModelsPaymentRequestEventEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPaymentRequestEvent().list() / client.NoFrixionMoneyMoovModelsPaymentRequestEvent().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payment_request_event_entity import NoFrixionMoneyMoovModelsPaymentRequestEventEntity
        return NoFrixionMoneyMoovModelsPaymentRequestEventEntity(self, data)


    def NoFrixionMoneyMoovModelsPaymentRequestMetric(self, data=None) -> "NoFrixionMoneyMoovModelsPaymentRequestMetricEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPaymentRequestMetric().list() / client.NoFrixionMoneyMoovModelsPaymentRequestMetric().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payment_request_metric_entity import NoFrixionMoneyMoovModelsPaymentRequestMetricEntity
        return NoFrixionMoneyMoovModelsPaymentRequestMetricEntity(self, data)


    def NoFrixionMoneyMoovModelsPaymentRequestMinimal(self, data=None) -> "NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPaymentRequestMinimal().list() / client.NoFrixionMoneyMoovModelsPaymentRequestMinimal().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payment_request_minimal_entity import NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity
        return NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity(self, data)


    def NoFrixionMoneyMoovModelsPaymentRequestResult(self, data=None) -> "NoFrixionMoneyMoovModelsPaymentRequestResultEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPaymentRequestResult().list() / client.NoFrixionMoneyMoovModelsPaymentRequestResult().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payment_request_result_entity import NoFrixionMoneyMoovModelsPaymentRequestResultEntity
        return NoFrixionMoneyMoovModelsPaymentRequestResultEntity(self, data)


    def NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(self, data=None) -> "NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment().list() / client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payment_requests_merchant_payment_entity import NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity
        return NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity(self, data)


    def NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(self, data=None) -> "NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2().list() / client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payment_requests_merchant_payment2_entity import NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity
        return NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity(self, data)


    def NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(self, data=None) -> "NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3().list() / client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payment_requests_merchant_payment3_entity import NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity
        return NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity(self, data)


    def NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(self, data=None) -> "NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4().list() / client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payment_requests_merchant_payment4_entity import NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity
        return NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity(self, data)


    def NoFrixionMoneyMoovModelsPayoutKeysetPage(self, data=None) -> "NoFrixionMoneyMoovModelsPayoutKeysetPageEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPayoutKeysetPage().list() / client.NoFrixionMoneyMoovModelsPayoutKeysetPage().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payout_keyset_page_entity import NoFrixionMoneyMoovModelsPayoutKeysetPageEntity
        return NoFrixionMoneyMoovModelsPayoutKeysetPageEntity(self, data)


    def NoFrixionMoneyMoovModelsPayoutMetric(self, data=None) -> "NoFrixionMoneyMoovModelsPayoutMetricEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPayoutMetric().list() / client.NoFrixionMoneyMoovModelsPayoutMetric().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payout_metric_entity import NoFrixionMoneyMoovModelsPayoutMetricEntity
        return NoFrixionMoneyMoovModelsPayoutMetricEntity(self, data)


    def NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(self, data=None) -> "NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate().list() / client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payouts_payouts_create_entity import NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity
        return NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity(self, data)


    def NoFrixionMoneyMoovModelsPayrun(self, data=None) -> "NoFrixionMoneyMoovModelsPayrunEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsPayrun().list() / client.NoFrixionMoneyMoovModelsPayrun().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_payrun_entity import NoFrixionMoneyMoovModelsPayrunEntity
        return NoFrixionMoneyMoovModelsPayrunEntity(self, data)


    def NoFrixionMoneyMoovModelsReportResult(self, data=None) -> "NoFrixionMoneyMoovModelsReportResultEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsReportResult().list() / client.NoFrixionMoneyMoovModelsReportResult().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_report_result_entity import NoFrixionMoneyMoovModelsReportResultEntity
        return NoFrixionMoneyMoovModelsReportResultEntity(self, data)


    def NoFrixionMoneyMoovModelsRule(self, data=None) -> "NoFrixionMoneyMoovModelsRuleEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsRule().list() / client.NoFrixionMoneyMoovModelsRule().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_rule_entity import NoFrixionMoneyMoovModelsRuleEntity
        return NoFrixionMoneyMoovModelsRuleEntity(self, data)


    def NoFrixionMoneyMoovModelsTransaction(self, data=None) -> "NoFrixionMoneyMoovModelsTransactionEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsTransaction().list() / client.NoFrixionMoneyMoovModelsTransaction().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_transaction_entity import NoFrixionMoneyMoovModelsTransactionEntity
        return NoFrixionMoneyMoovModelsTransactionEntity(self, data)


    def NoFrixionMoneyMoovModelsTransactionPage(self, data=None) -> "NoFrixionMoneyMoovModelsTransactionPageEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsTransactionPage().list() / client.NoFrixionMoneyMoovModelsTransactionPage().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_transaction_page_entity import NoFrixionMoneyMoovModelsTransactionPageEntity
        return NoFrixionMoneyMoovModelsTransactionPageEntity(self, data)


    def NoFrixionMoneyMoovModelsUserInvite(self, data=None) -> "NoFrixionMoneyMoovModelsUserInviteEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsUserInvite().list() / client.NoFrixionMoneyMoovModelsUserInvite().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_user_invite_entity import NoFrixionMoneyMoovModelsUserInviteEntity
        return NoFrixionMoneyMoovModelsUserInviteEntity(self, data)


    def NoFrixionMoneyMoovModelsUserInvitePage(self, data=None) -> "NoFrixionMoneyMoovModelsUserInvitePageEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsUserInvitePage().list() / client.NoFrixionMoneyMoovModelsUserInvitePage().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_user_invite_page_entity import NoFrixionMoneyMoovModelsUserInvitePageEntity
        return NoFrixionMoneyMoovModelsUserInvitePageEntity(self, data)


    def NoFrixionMoneyMoovModelsUserPage(self, data=None) -> "NoFrixionMoneyMoovModelsUserPageEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsUserPage().list() / client.NoFrixionMoneyMoovModelsUserPage().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_user_page_entity import NoFrixionMoneyMoovModelsUserPageEntity
        return NoFrixionMoneyMoovModelsUserPageEntity(self, data)


    def NoFrixionMoneyMoovModelsWebhook(self, data=None) -> "NoFrixionMoneyMoovModelsWebhookEntity":
        """Entity factory: client.NoFrixionMoneyMoovModelsWebhook().list() / client.NoFrixionMoneyMoovModelsWebhook().load({"id": ...})."""
        from entity.no_frixion_money_moov_models_webhook_entity import NoFrixionMoneyMoovModelsWebhookEntity
        return NoFrixionMoneyMoovModelsWebhookEntity(self, data)


    def OpenBanking(self, data=None) -> "OpenBankingEntity":
        """Entity factory: client.OpenBanking().list() / client.OpenBanking().load({"id": ...})."""
        from entity.open_banking_entity import OpenBankingEntity
        return OpenBankingEntity(self, data)


    def Payeeverification(self, data=None) -> "PayeeverificationEntity":
        """Entity factory: client.Payeeverification().list() / client.Payeeverification().load({"id": ...})."""
        from entity.payeeverification_entity import PayeeverificationEntity
        return PayeeverificationEntity(self, data)


    def PaymentRequest(self, data=None) -> "PaymentRequestEntity":
        """Entity factory: client.PaymentRequest().list() / client.PaymentRequest().load({"id": ...})."""
        from entity.payment_request_entity import PaymentRequestEntity
        return PaymentRequestEntity(self, data)


    def Payout(self, data=None) -> "PayoutEntity":
        """Entity factory: client.Payout().list() / client.Payout().load({"id": ...})."""
        from entity.payout_entity import PayoutEntity
        return PayoutEntity(self, data)


    def Payrun(self, data=None) -> "PayrunEntity":
        """Entity factory: client.Payrun().list() / client.Payrun().load({"id": ...})."""
        from entity.payrun_entity import PayrunEntity
        return PayrunEntity(self, data)


    def Reject(self, data=None) -> "RejectEntity":
        """Entity factory: client.Reject().list() / client.Reject().load({"id": ...})."""
        from entity.reject_entity import RejectEntity
        return RejectEntity(self, data)


    def Report(self, data=None) -> "ReportEntity":
        """Entity factory: client.Report().list() / client.Report().load({"id": ...})."""
        from entity.report_entity import ReportEntity
        return ReportEntity(self, data)


    def Rule(self, data=None) -> "RuleEntity":
        """Entity factory: client.Rule().list() / client.Rule().load({"id": ...})."""
        from entity.rule_entity import RuleEntity
        return RuleEntity(self, data)


    def Send(self, data=None) -> "SendEntity":
        """Entity factory: client.Send().list() / client.Send().load({"id": ...})."""
        from entity.send_entity import SendEntity
        return SendEntity(self, data)


    def Sendbeneficiary(self, data=None) -> "SendbeneficiaryEntity":
        """Entity factory: client.Sendbeneficiary().list() / client.Sendbeneficiary().load({"id": ...})."""
        from entity.sendbeneficiary_entity import SendbeneficiaryEntity
        return SendbeneficiaryEntity(self, data)


    def Tag(self, data=None) -> "TagEntity":
        """Entity factory: client.Tag().list() / client.Tag().load({"id": ...})."""
        from entity.tag_entity import TagEntity
        return TagEntity(self, data)


    def Token(self, data=None) -> "TokenEntity":
        """Entity factory: client.Token().list() / client.Token().load({"id": ...})."""
        from entity.token_entity import TokenEntity
        return TokenEntity(self, data)


    def Transaction(self, data=None) -> "TransactionEntity":
        """Entity factory: client.Transaction().list() / client.Transaction().load({"id": ...})."""
        from entity.transaction_entity import TransactionEntity
        return TransactionEntity(self, data)


    def User(self, data=None) -> "UserEntity":
        """Entity factory: client.User().list() / client.User().load({"id": ...})."""
        from entity.user_entity import UserEntity
        return UserEntity(self, data)


    def UserInvite(self, data=None) -> "UserInviteEntity":
        """Entity factory: client.UserInvite().list() / client.UserInvite().load({"id": ...})."""
        from entity.user_invite_entity import UserInviteEntity
        return UserInviteEntity(self, data)


    def Virtual(self, data=None) -> "VirtualEntity":
        """Entity factory: client.Virtual().list() / client.Virtual().load({"id": ...})."""
        from entity.virtual_entity import VirtualEntity
        return VirtualEntity(self, data)


    def Webhook(self, data=None) -> "WebhookEntity":
        """Entity factory: client.Webhook().list() / client.Webhook().load({"id": ...})."""
        from entity.webhook_entity import WebhookEntity
        return WebhookEntity(self, data)


    def Whoami(self, data=None) -> "WhoamiEntity":
        """Entity factory: client.Whoami().list() / client.Whoami().load({"id": ...})."""
        from entity.whoami_entity import WhoamiEntity
        return WhoamiEntity(self, data)


    def Whoamitrustedapp(self, data=None) -> "WhoamitrustedappEntity":
        """Entity factory: client.Whoamitrustedapp().list() / client.Whoamitrustedapp().load({"id": ...})."""
        from entity.whoamitrustedapp_entity import WhoamitrustedappEntity
        return WhoamitrustedappEntity(self, data)



    @classmethod
    def test(cls, testopts=None, sdkopts=None) -> "NofrixionSDK":
        if sdkopts is None:
            sdkopts = {}
        sdkopts = vs.clone(sdkopts)
        if not isinstance(sdkopts, dict):
            sdkopts = {}

        if testopts is None:
            testopts = {}
        testopts = vs.clone(testopts)
        if not isinstance(testopts, dict):
            testopts = {}
        testopts["active"] = True

        vs.setpath(sdkopts, "feature.test", testopts)

        sdk = cls(sdkopts)
        sdk.mode = "test"

        return sdk


from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from entity.account_entity import AccountEntity
    from entity.beneficiary_entity import BeneficiaryEntity
    from entity.cancel_entity import CancelEntity
    from entity.disable_entity import DisableEntity
    from entity.enable_entity import EnableEntity
    from entity.merchant_entity import MerchantEntity
    from entity.metadata_entity import MetadataEntity
    from entity.no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page_entity import NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity
    from entity.no_frixion_biz_biz_models_paging_payment_request_page_entity import NoFrixionBizBizModelsPagingPaymentRequestPageEntity
    from entity.no_frixion_biz_biz_models_paging_payout_page_entity import NoFrixionBizBizModelsPagingPayoutPageEntity
    from entity.no_frixion_biz_biz_models_paging_payrun_page_entity import NoFrixionBizBizModelsPagingPayrunPageEntity
    from entity.no_frixion_biz_biz_models_paging_rule_events_page_entity import NoFrixionBizBizModelsPagingRuleEventsPageEntity
    from entity.no_frixion_biz_biz_models_paging_rules_page_entity import NoFrixionBizBizModelsPagingRulesPageEntity
    from entity.no_frixion_biz_biz_models_payments_card_payment_entity import NoFrixionBizBizModelsPaymentsCardPaymentEntity
    from entity.no_frixion_biz_biz_models_payments_card_public_key_entity import NoFrixionBizBizModelsPaymentsCardPublicKeyEntity
    from entity.no_frixion_money_moov_api_features_beneficiaries_beneficiaries_entity import NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity
    from entity.no_frixion_money_moov_api_features_payment_requests_payment_entity import NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity
    from entity.no_frixion_money_moov_api_features_permissions_roles_create_entity import NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity
    from entity.no_frixion_money_moov_api_features_user_invites_create_entity import NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity
    from entity.no_frixion_money_moov_models_authorisation_settings_merchant_entity import NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity
    from entity.no_frixion_money_moov_models_batch_payout_entity import NoFrixionMoneyMoovModelsBatchPayoutEntity
    from entity.no_frixion_money_moov_models_beneficiary_group_page_entity import NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity
    from entity.no_frixion_money_moov_models_beneficiary_page_entity import NoFrixionMoneyMoovModelsBeneficiaryPageEntity
    from entity.no_frixion_money_moov_models_card_customer_token_entity import NoFrixionMoneyMoovModelsCardCustomerTokenEntity
    from entity.no_frixion_money_moov_models_currency_currency_info_entity import NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity
    from entity.no_frixion_money_moov_models_direct_debit_batch_submit_entity import NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity
    from entity.no_frixion_money_moov_models_fx_rate_entity import NoFrixionMoneyMoovModelsFxRateEntity
    from entity.no_frixion_money_moov_models_i_payment_entity import NoFrixionMoneyMoovModelsIPaymentEntity
    from entity.no_frixion_money_moov_models_mandates_mandate_entity import NoFrixionMoneyMoovModelsMandatesMandateEntity
    from entity.no_frixion_money_moov_models_merchant_entity import NoFrixionMoneyMoovModelsMerchantEntity
    from entity.no_frixion_money_moov_models_merchant_page_entity import NoFrixionMoneyMoovModelsMerchantPageEntity
    from entity.no_frixion_money_moov_models_merchant_pay_by_bank_setting_entity import NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity
    from entity.no_frixion_money_moov_models_merchant_token_entity import NoFrixionMoneyMoovModelsMerchantTokenEntity
    from entity.no_frixion_money_moov_models_merchant_token_page_entity import NoFrixionMoneyMoovModelsMerchantTokenPageEntity
    from entity.no_frixion_money_moov_models_no_frixion_version_entity import NoFrixionMoneyMoovModelsNoFrixionVersionEntity
    from entity.no_frixion_money_moov_models_open_banking_account_entity import NoFrixionMoneyMoovModelsOpenBankingAccountEntity
    from entity.no_frixion_money_moov_models_open_banking_consent_entity import NoFrixionMoneyMoovModelsOpenBankingConsentEntity
    from entity.no_frixion_money_moov_models_open_banking_transaction_entity import NoFrixionMoneyMoovModelsOpenBankingTransactionEntity
    from entity.no_frixion_money_moov_models_payment_entity import NoFrixionMoneyMoovModelsPaymentEntity
    from entity.no_frixion_money_moov_models_payment_account_minimal_page_entity import NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity
    from entity.no_frixion_money_moov_models_payment_account_page_entity import NoFrixionMoneyMoovModelsPaymentAccountPageEntity
    from entity.no_frixion_money_moov_models_payment_initiation_entity import NoFrixionMoneyMoovModelsPaymentInitiationEntity
    from entity.no_frixion_money_moov_models_payment_request_event_entity import NoFrixionMoneyMoovModelsPaymentRequestEventEntity
    from entity.no_frixion_money_moov_models_payment_request_metric_entity import NoFrixionMoneyMoovModelsPaymentRequestMetricEntity
    from entity.no_frixion_money_moov_models_payment_request_minimal_entity import NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity
    from entity.no_frixion_money_moov_models_payment_request_result_entity import NoFrixionMoneyMoovModelsPaymentRequestResultEntity
    from entity.no_frixion_money_moov_models_payment_requests_merchant_payment_entity import NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity
    from entity.no_frixion_money_moov_models_payment_requests_merchant_payment2_entity import NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity
    from entity.no_frixion_money_moov_models_payment_requests_merchant_payment3_entity import NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity
    from entity.no_frixion_money_moov_models_payment_requests_merchant_payment4_entity import NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity
    from entity.no_frixion_money_moov_models_payout_keyset_page_entity import NoFrixionMoneyMoovModelsPayoutKeysetPageEntity
    from entity.no_frixion_money_moov_models_payout_metric_entity import NoFrixionMoneyMoovModelsPayoutMetricEntity
    from entity.no_frixion_money_moov_models_payouts_payouts_create_entity import NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity
    from entity.no_frixion_money_moov_models_payrun_entity import NoFrixionMoneyMoovModelsPayrunEntity
    from entity.no_frixion_money_moov_models_report_result_entity import NoFrixionMoneyMoovModelsReportResultEntity
    from entity.no_frixion_money_moov_models_rule_entity import NoFrixionMoneyMoovModelsRuleEntity
    from entity.no_frixion_money_moov_models_transaction_entity import NoFrixionMoneyMoovModelsTransactionEntity
    from entity.no_frixion_money_moov_models_transaction_page_entity import NoFrixionMoneyMoovModelsTransactionPageEntity
    from entity.no_frixion_money_moov_models_user_invite_entity import NoFrixionMoneyMoovModelsUserInviteEntity
    from entity.no_frixion_money_moov_models_user_invite_page_entity import NoFrixionMoneyMoovModelsUserInvitePageEntity
    from entity.no_frixion_money_moov_models_user_page_entity import NoFrixionMoneyMoovModelsUserPageEntity
    from entity.no_frixion_money_moov_models_webhook_entity import NoFrixionMoneyMoovModelsWebhookEntity
    from entity.open_banking_entity import OpenBankingEntity
    from entity.payeeverification_entity import PayeeverificationEntity
    from entity.payment_request_entity import PaymentRequestEntity
    from entity.payout_entity import PayoutEntity
    from entity.payrun_entity import PayrunEntity
    from entity.reject_entity import RejectEntity
    from entity.report_entity import ReportEntity
    from entity.rule_entity import RuleEntity
    from entity.send_entity import SendEntity
    from entity.sendbeneficiary_entity import SendbeneficiaryEntity
    from entity.tag_entity import TagEntity
    from entity.token_entity import TokenEntity
    from entity.transaction_entity import TransactionEntity
    from entity.user_entity import UserEntity
    from entity.user_invite_entity import UserInviteEntity
    from entity.virtual_entity import VirtualEntity
    from entity.webhook_entity import WebhookEntity
    from entity.whoami_entity import WhoamiEntity
    from entity.whoamitrustedapp_entity import WhoamitrustedappEntity
