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


    def Batch(self, data=None) -> "BatchEntity":
        """Entity factory: client.Batch().list() / client.Batch().load({"id": ...})."""
        from entity.batch_entity import BatchEntity
        return BatchEntity(self, data)


    def BeneficiariesCreate(self, data=None) -> "BeneficiariesCreateEntity":
        """Entity factory: client.BeneficiariesCreate().list() / client.BeneficiariesCreate().load({"id": ...})."""
        from entity.beneficiaries_create_entity import BeneficiariesCreateEntity
        return BeneficiariesCreateEntity(self, data)


    def Beneficiary(self, data=None) -> "BeneficiaryEntity":
        """Entity factory: client.Beneficiary().list() / client.Beneficiary().load({"id": ...})."""
        from entity.beneficiary_entity import BeneficiaryEntity
        return BeneficiaryEntity(self, data)


    def BeneficiaryGroup(self, data=None) -> "BeneficiaryGroupEntity":
        """Entity factory: client.BeneficiaryGroup().list() / client.BeneficiaryGroup().load({"id": ...})."""
        from entity.beneficiary_group_entity import BeneficiaryGroupEntity
        return BeneficiaryGroupEntity(self, data)


    def Card(self, data=None) -> "CardEntity":
        """Entity factory: client.Card().list() / client.Card().load({"id": ...})."""
        from entity.card_entity import CardEntity
        return CardEntity(self, data)


    def CardCustomerToken(self, data=None) -> "CardCustomerTokenEntity":
        """Entity factory: client.CardCustomerToken().list() / client.CardCustomerToken().load({"id": ...})."""
        from entity.card_customer_token_entity import CardCustomerTokenEntity
        return CardCustomerTokenEntity(self, data)


    def CardPayment(self, data=None) -> "CardPaymentEntity":
        """Entity factory: client.CardPayment().list() / client.CardPayment().load({"id": ...})."""
        from entity.card_payment_entity import CardPaymentEntity
        return CardPaymentEntity(self, data)


    def CardPublicKey(self, data=None) -> "CardPublicKeyEntity":
        """Entity factory: client.CardPublicKey().list() / client.CardPublicKey().load({"id": ...})."""
        from entity.card_public_key_entity import CardPublicKeyEntity
        return CardPublicKeyEntity(self, data)


    def Consent(self, data=None) -> "ConsentEntity":
        """Entity factory: client.Consent().list() / client.Consent().load({"id": ...})."""
        from entity.consent_entity import ConsentEntity
        return ConsentEntity(self, data)


    def Currency(self, data=None) -> "CurrencyEntity":
        """Entity factory: client.Currency().list() / client.Currency().load({"id": ...})."""
        from entity.currency_entity import CurrencyEntity
        return CurrencyEntity(self, data)


    def DirectDebitBatchSubmit(self, data=None) -> "DirectDebitBatchSubmitEntity":
        """Entity factory: client.DirectDebitBatchSubmit().list() / client.DirectDebitBatchSubmit().load({"id": ...})."""
        from entity.direct_debit_batch_submit_entity import DirectDebitBatchSubmitEntity
        return DirectDebitBatchSubmitEntity(self, data)


    def FxRate(self, data=None) -> "FxRateEntity":
        """Entity factory: client.FxRate().list() / client.FxRate().load({"id": ...})."""
        from entity.fx_rate_entity import FxRateEntity
        return FxRateEntity(self, data)


    def IPayment(self, data=None) -> "IPaymentEntity":
        """Entity factory: client.IPayment().list() / client.IPayment().load({"id": ...})."""
        from entity.i_payment_entity import IPaymentEntity
        return IPaymentEntity(self, data)


    def Mandate(self, data=None) -> "MandateEntity":
        """Entity factory: client.Mandate().list() / client.Mandate().load({"id": ...})."""
        from entity.mandate_entity import MandateEntity
        return MandateEntity(self, data)


    def Merchant(self, data=None) -> "MerchantEntity":
        """Entity factory: client.Merchant().list() / client.Merchant().load({"id": ...})."""
        from entity.merchant_entity import MerchantEntity
        return MerchantEntity(self, data)


    def MerchantAuthorisationSetting(self, data=None) -> "MerchantAuthorisationSettingEntity":
        """Entity factory: client.MerchantAuthorisationSetting().list() / client.MerchantAuthorisationSetting().load({"id": ...})."""
        from entity.merchant_authorisation_setting_entity import MerchantAuthorisationSettingEntity
        return MerchantAuthorisationSettingEntity(self, data)


    def MerchantDirectDebitMandate(self, data=None) -> "MerchantDirectDebitMandateEntity":
        """Entity factory: client.MerchantDirectDebitMandate().list() / client.MerchantDirectDebitMandate().load({"id": ...})."""
        from entity.merchant_direct_debit_mandate_entity import MerchantDirectDebitMandateEntity
        return MerchantDirectDebitMandateEntity(self, data)


    def MerchantPayByBankSetting(self, data=None) -> "MerchantPayByBankSettingEntity":
        """Entity factory: client.MerchantPayByBankSetting().list() / client.MerchantPayByBankSetting().load({"id": ...})."""
        from entity.merchant_pay_by_bank_setting_entity import MerchantPayByBankSettingEntity
        return MerchantPayByBankSettingEntity(self, data)


    def MerchantPaymentRequestTemplate(self, data=None) -> "MerchantPaymentRequestTemplateEntity":
        """Entity factory: client.MerchantPaymentRequestTemplate().list() / client.MerchantPaymentRequestTemplate().load({"id": ...})."""
        from entity.merchant_payment_request_template_entity import MerchantPaymentRequestTemplateEntity
        return MerchantPaymentRequestTemplateEntity(self, data)


    def MerchantToken(self, data=None) -> "MerchantTokenEntity":
        """Entity factory: client.MerchantToken().list() / client.MerchantToken().load({"id": ...})."""
        from entity.merchant_token_entity import MerchantTokenEntity
        return MerchantTokenEntity(self, data)


    def Metadata(self, data=None) -> "MetadataEntity":
        """Entity factory: client.Metadata().list() / client.Metadata().load({"id": ...})."""
        from entity.metadata_entity import MetadataEntity
        return MetadataEntity(self, data)


    def NoFrixionVersion(self, data=None) -> "NoFrixionVersionEntity":
        """Entity factory: client.NoFrixionVersion().list() / client.NoFrixionVersion().load({"id": ...})."""
        from entity.no_frixion_version_entity import NoFrixionVersionEntity
        return NoFrixionVersionEntity(self, data)


    def OpenBanking(self, data=None) -> "OpenBankingEntity":
        """Entity factory: client.OpenBanking().list() / client.OpenBanking().load({"id": ...})."""
        from entity.open_banking_entity import OpenBankingEntity
        return OpenBankingEntity(self, data)


    def Payeeverification(self, data=None) -> "PayeeverificationEntity":
        """Entity factory: client.Payeeverification().list() / client.Payeeverification().load({"id": ...})."""
        from entity.payeeverification_entity import PayeeverificationEntity
        return PayeeverificationEntity(self, data)


    def Payment(self, data=None) -> "PaymentEntity":
        """Entity factory: client.Payment().list() / client.Payment().load({"id": ...})."""
        from entity.payment_entity import PaymentEntity
        return PaymentEntity(self, data)


    def PaymentAccount(self, data=None) -> "PaymentAccountEntity":
        """Entity factory: client.PaymentAccount().list() / client.PaymentAccount().load({"id": ...})."""
        from entity.payment_account_entity import PaymentAccountEntity
        return PaymentAccountEntity(self, data)


    def PaymentAccountMinimal(self, data=None) -> "PaymentAccountMinimalEntity":
        """Entity factory: client.PaymentAccountMinimal().list() / client.PaymentAccountMinimal().load({"id": ...})."""
        from entity.payment_account_minimal_entity import PaymentAccountMinimalEntity
        return PaymentAccountMinimalEntity(self, data)


    def PaymentInitiation(self, data=None) -> "PaymentInitiationEntity":
        """Entity factory: client.PaymentInitiation().list() / client.PaymentInitiation().load({"id": ...})."""
        from entity.payment_initiation_entity import PaymentInitiationEntity
        return PaymentInitiationEntity(self, data)


    def PaymentRequest(self, data=None) -> "PaymentRequestEntity":
        """Entity factory: client.PaymentRequest().list() / client.PaymentRequest().load({"id": ...})."""
        from entity.payment_request_entity import PaymentRequestEntity
        return PaymentRequestEntity(self, data)


    def PaymentRequestEvent(self, data=None) -> "PaymentRequestEventEntity":
        """Entity factory: client.PaymentRequestEvent().list() / client.PaymentRequestEvent().load({"id": ...})."""
        from entity.payment_request_event_entity import PaymentRequestEventEntity
        return PaymentRequestEventEntity(self, data)


    def PaymentRequestMetric(self, data=None) -> "PaymentRequestMetricEntity":
        """Entity factory: client.PaymentRequestMetric().list() / client.PaymentRequestMetric().load({"id": ...})."""
        from entity.payment_request_metric_entity import PaymentRequestMetricEntity
        return PaymentRequestMetricEntity(self, data)


    def PaymentRequestMinimal(self, data=None) -> "PaymentRequestMinimalEntity":
        """Entity factory: client.PaymentRequestMinimal().list() / client.PaymentRequestMinimal().load({"id": ...})."""
        from entity.payment_request_minimal_entity import PaymentRequestMinimalEntity
        return PaymentRequestMinimalEntity(self, data)


    def PaymentRequestResult(self, data=None) -> "PaymentRequestResultEntity":
        """Entity factory: client.PaymentRequestResult().list() / client.PaymentRequestResult().load({"id": ...})."""
        from entity.payment_request_result_entity import PaymentRequestResultEntity
        return PaymentRequestResultEntity(self, data)


    def PaymentRequestsCreate(self, data=None) -> "PaymentRequestsCreateEntity":
        """Entity factory: client.PaymentRequestsCreate().list() / client.PaymentRequestsCreate().load({"id": ...})."""
        from entity.payment_requests_create_entity import PaymentRequestsCreateEntity
        return PaymentRequestsCreateEntity(self, data)


    def Payout(self, data=None) -> "PayoutEntity":
        """Entity factory: client.Payout().list() / client.Payout().load({"id": ...})."""
        from entity.payout_entity import PayoutEntity
        return PayoutEntity(self, data)


    def PayoutKeyset(self, data=None) -> "PayoutKeysetEntity":
        """Entity factory: client.PayoutKeyset().list() / client.PayoutKeyset().load({"id": ...})."""
        from entity.payout_keyset_entity import PayoutKeysetEntity
        return PayoutKeysetEntity(self, data)


    def PayoutMetric(self, data=None) -> "PayoutMetricEntity":
        """Entity factory: client.PayoutMetric().list() / client.PayoutMetric().load({"id": ...})."""
        from entity.payout_metric_entity import PayoutMetricEntity
        return PayoutMetricEntity(self, data)


    def PayoutsCreate(self, data=None) -> "PayoutsCreateEntity":
        """Entity factory: client.PayoutsCreate().list() / client.PayoutsCreate().load({"id": ...})."""
        from entity.payouts_create_entity import PayoutsCreateEntity
        return PayoutsCreateEntity(self, data)


    def Payrun(self, data=None) -> "PayrunEntity":
        """Entity factory: client.Payrun().list() / client.Payrun().load({"id": ...})."""
        from entity.payrun_entity import PayrunEntity
        return PayrunEntity(self, data)


    def Report(self, data=None) -> "ReportEntity":
        """Entity factory: client.Report().list() / client.Report().load({"id": ...})."""
        from entity.report_entity import ReportEntity
        return ReportEntity(self, data)


    def ReportResult(self, data=None) -> "ReportResultEntity":
        """Entity factory: client.ReportResult().list() / client.ReportResult().load({"id": ...})."""
        from entity.report_result_entity import ReportResultEntity
        return ReportResultEntity(self, data)


    def RolesCreate(self, data=None) -> "RolesCreateEntity":
        """Entity factory: client.RolesCreate().list() / client.RolesCreate().load({"id": ...})."""
        from entity.roles_create_entity import RolesCreateEntity
        return RolesCreateEntity(self, data)


    def Rule(self, data=None) -> "RuleEntity":
        """Entity factory: client.Rule().list() / client.Rule().load({"id": ...})."""
        from entity.rule_entity import RuleEntity
        return RuleEntity(self, data)


    def RuleEvent(self, data=None) -> "RuleEventEntity":
        """Entity factory: client.RuleEvent().list() / client.RuleEvent().load({"id": ...})."""
        from entity.rule_event_entity import RuleEventEntity
        return RuleEventEntity(self, data)


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


    def UserInvitesCreate(self, data=None) -> "UserInvitesCreateEntity":
        """Entity factory: client.UserInvitesCreate().list() / client.UserInvitesCreate().load({"id": ...})."""
        from entity.user_invites_create_entity import UserInvitesCreateEntity
        return UserInvitesCreateEntity(self, data)


    def Virtual(self, data=None) -> "VirtualEntity":
        """Entity factory: client.Virtual().list() / client.Virtual().load({"id": ...})."""
        from entity.virtual_entity import VirtualEntity
        return VirtualEntity(self, data)


    def Webhook(self, data=None) -> "WebhookEntity":
        """Entity factory: client.Webhook().list() / client.Webhook().load({"id": ...})."""
        from entity.webhook_entity import WebhookEntity
        return WebhookEntity(self, data)



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
    from entity.batch_entity import BatchEntity
    from entity.beneficiaries_create_entity import BeneficiariesCreateEntity
    from entity.beneficiary_entity import BeneficiaryEntity
    from entity.beneficiary_group_entity import BeneficiaryGroupEntity
    from entity.card_entity import CardEntity
    from entity.card_customer_token_entity import CardCustomerTokenEntity
    from entity.card_payment_entity import CardPaymentEntity
    from entity.card_public_key_entity import CardPublicKeyEntity
    from entity.consent_entity import ConsentEntity
    from entity.currency_entity import CurrencyEntity
    from entity.direct_debit_batch_submit_entity import DirectDebitBatchSubmitEntity
    from entity.fx_rate_entity import FxRateEntity
    from entity.i_payment_entity import IPaymentEntity
    from entity.mandate_entity import MandateEntity
    from entity.merchant_entity import MerchantEntity
    from entity.merchant_authorisation_setting_entity import MerchantAuthorisationSettingEntity
    from entity.merchant_direct_debit_mandate_entity import MerchantDirectDebitMandateEntity
    from entity.merchant_pay_by_bank_setting_entity import MerchantPayByBankSettingEntity
    from entity.merchant_payment_request_template_entity import MerchantPaymentRequestTemplateEntity
    from entity.merchant_token_entity import MerchantTokenEntity
    from entity.metadata_entity import MetadataEntity
    from entity.no_frixion_version_entity import NoFrixionVersionEntity
    from entity.open_banking_entity import OpenBankingEntity
    from entity.payeeverification_entity import PayeeverificationEntity
    from entity.payment_entity import PaymentEntity
    from entity.payment_account_entity import PaymentAccountEntity
    from entity.payment_account_minimal_entity import PaymentAccountMinimalEntity
    from entity.payment_initiation_entity import PaymentInitiationEntity
    from entity.payment_request_entity import PaymentRequestEntity
    from entity.payment_request_event_entity import PaymentRequestEventEntity
    from entity.payment_request_metric_entity import PaymentRequestMetricEntity
    from entity.payment_request_minimal_entity import PaymentRequestMinimalEntity
    from entity.payment_request_result_entity import PaymentRequestResultEntity
    from entity.payment_requests_create_entity import PaymentRequestsCreateEntity
    from entity.payout_entity import PayoutEntity
    from entity.payout_keyset_entity import PayoutKeysetEntity
    from entity.payout_metric_entity import PayoutMetricEntity
    from entity.payouts_create_entity import PayoutsCreateEntity
    from entity.payrun_entity import PayrunEntity
    from entity.report_entity import ReportEntity
    from entity.report_result_entity import ReportResultEntity
    from entity.roles_create_entity import RolesCreateEntity
    from entity.rule_entity import RuleEntity
    from entity.rule_event_entity import RuleEventEntity
    from entity.tag_entity import TagEntity
    from entity.token_entity import TokenEntity
    from entity.transaction_entity import TransactionEntity
    from entity.user_entity import UserEntity
    from entity.user_invite_entity import UserInviteEntity
    from entity.user_invites_create_entity import UserInvitesCreateEntity
    from entity.virtual_entity import VirtualEntity
    from entity.webhook_entity import WebhookEntity
