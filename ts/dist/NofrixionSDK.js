"use strict";
// Nofrixion Ts SDK
Object.defineProperty(exports, "__esModule", { value: true });
exports.SDK = exports.NofrixionSDK = exports.NofrixionEntityBase = exports.BaseFeature = exports.config = exports.stdutil = void 0;
const AccountEntity_1 = require("./entity/AccountEntity");
const BatchEntity_1 = require("./entity/BatchEntity");
const BeneficiaryEntity_1 = require("./entity/BeneficiaryEntity");
const BeneficiaryGroupEntity_1 = require("./entity/BeneficiaryGroupEntity");
const CardEntity_1 = require("./entity/CardEntity");
const CardCustomerTokenEntity_1 = require("./entity/CardCustomerTokenEntity");
const CardPaymentEntity_1 = require("./entity/CardPaymentEntity");
const CardPublicKeyEntity_1 = require("./entity/CardPublicKeyEntity");
const ConsentEntity_1 = require("./entity/ConsentEntity");
const CurrencyEntity_1 = require("./entity/CurrencyEntity");
const DirectDebitBatchSubmitEntity_1 = require("./entity/DirectDebitBatchSubmitEntity");
const FxRateEntity_1 = require("./entity/FxRateEntity");
const IPaymentEntity_1 = require("./entity/IPaymentEntity");
const MandateEntity_1 = require("./entity/MandateEntity");
const MerchantEntity_1 = require("./entity/MerchantEntity");
const MerchantAuthorisationSettingEntity_1 = require("./entity/MerchantAuthorisationSettingEntity");
const MerchantDirectDebitMandatePageEntity_1 = require("./entity/MerchantDirectDebitMandatePageEntity");
const MerchantPayByBankSettingEntity_1 = require("./entity/MerchantPayByBankSettingEntity");
const MerchantPaymentRequestTemplateEntity_1 = require("./entity/MerchantPaymentRequestTemplateEntity");
const MerchantTokenEntity_1 = require("./entity/MerchantTokenEntity");
const MetadataEntity_1 = require("./entity/MetadataEntity");
const NoFrixionVersionEntity_1 = require("./entity/NoFrixionVersionEntity");
const OpenBankingEntity_1 = require("./entity/OpenBankingEntity");
const PayeeverificationEntity_1 = require("./entity/PayeeverificationEntity");
const PaymentEntity_1 = require("./entity/PaymentEntity");
const PaymentAccountEntity_1 = require("./entity/PaymentAccountEntity");
const PaymentAccountMinimalEntity_1 = require("./entity/PaymentAccountMinimalEntity");
const PaymentInitiationEntity_1 = require("./entity/PaymentInitiationEntity");
const PaymentRequestEntity_1 = require("./entity/PaymentRequestEntity");
const PaymentRequestEventEntity_1 = require("./entity/PaymentRequestEventEntity");
const PaymentRequestMetricEntity_1 = require("./entity/PaymentRequestMetricEntity");
const PaymentRequestMinimalEntity_1 = require("./entity/PaymentRequestMinimalEntity");
const PaymentRequestResultEntity_1 = require("./entity/PaymentRequestResultEntity");
const PayoutEntity_1 = require("./entity/PayoutEntity");
const PayoutKeysetPageEntity_1 = require("./entity/PayoutKeysetPageEntity");
const PayoutMetricEntity_1 = require("./entity/PayoutMetricEntity");
const PayrunEntity_1 = require("./entity/PayrunEntity");
const ReportEntity_1 = require("./entity/ReportEntity");
const ReportResultEntity_1 = require("./entity/ReportResultEntity");
const RoleEntity_1 = require("./entity/RoleEntity");
const RuleEntity_1 = require("./entity/RuleEntity");
const RuleEventEntity_1 = require("./entity/RuleEventEntity");
const TagEntity_1 = require("./entity/TagEntity");
const TokenEntity_1 = require("./entity/TokenEntity");
const TransactionEntity_1 = require("./entity/TransactionEntity");
const UserEntity_1 = require("./entity/UserEntity");
const UserInviteEntity_1 = require("./entity/UserInviteEntity");
const VirtualEntity_1 = require("./entity/VirtualEntity");
const WebhookEntity_1 = require("./entity/WebhookEntity");
const node_util_1 = require("node:util");
const Config_1 = require("./Config");
Object.defineProperty(exports, "config", { enumerable: true, get: function () { return Config_1.config; } });
const NofrixionEntityBase_1 = require("./NofrixionEntityBase");
Object.defineProperty(exports, "NofrixionEntityBase", { enumerable: true, get: function () { return NofrixionEntityBase_1.NofrixionEntityBase; } });
const Utility_1 = require("./utility/Utility");
const BaseFeature_1 = require("./feature/base/BaseFeature");
Object.defineProperty(exports, "BaseFeature", { enumerable: true, get: function () { return BaseFeature_1.BaseFeature; } });
const stdutil = new Utility_1.Utility();
exports.stdutil = stdutil;
class NofrixionSDK {
    _mode = 'live';
    _options;
    _utility = new Utility_1.Utility();
    _features;
    _rootctx;
    constructor(options) {
        this._rootctx = this._utility.makeContext({
            client: this,
            utility: this._utility,
            config: Config_1.config,
            options,
            shared: new WeakMap()
        });
        this._options = this._utility.makeOptions(this._rootctx);
        const struct = this._utility.struct;
        const getpath = struct.getpath;
        if (true === getpath(this._options.feature, 'test.active')) {
            this._mode = 'test';
        }
        this._rootctx.options = this._options;
        this._features = [];
        const featureAdd = this._utility.featureAdd;
        const featureInit = this._utility.featureInit;
        // Add features in the resolved order (makeOptions puts an explicit
        // array order first, else defaults to test-first). Ordering matters:
        // the `test` feature installs the base mock transport and the transport
        // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is current,
        // so `test` must be added before them to sit at the base of the chain.
        const extend = this._options.extend || [];
        const featureorder = getpath(this._options, '__derived__.featureorder') || [];
        for (const fname of featureorder) {
            const fopts = this._options.feature[fname] || {};
            if (fopts.active) {
                // An active name with no generated class is legal when an
                // extend-supplied instance carries that name (station's adopt
                // path): the instance is added below, positioned by its own
                // __after__ entry, so skip it here rather than fail construction.
                if (!this._rootctx.config.hasFeature(fname) &&
                    extend.some((f) => fname === f.name)) {
                    continue;
                }
                featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname));
            }
        }
        for (let f of extend) {
            featureAdd(this._rootctx, f);
        }
        for (let f of this._features) {
            featureInit(this._rootctx, f);
        }
        const featureHook = this._utility.featureHook;
        featureHook(this._rootctx, 'PostConstruct');
    }
    options() {
        return this._utility.struct.clone(this._options);
    }
    utility() {
        return this._utility.struct.clone(this._utility);
    }
    async prepare(fetchargs) {
        const utility = this._utility;
        const struct = utility.struct;
        const clone = struct.clone;
        const { makeContext, makeFetchDef, prepareHeaders, prepareAuth, } = utility;
        fetchargs = fetchargs || {};
        let ctx = makeContext({
            opname: 'prepare',
            ctrl: fetchargs.ctrl || {},
        }, this._rootctx);
        const options = this._options;
        // Build spec directly from SDK options + user-provided fetch args.
        const spec = {
            base: options.base,
            prefix: options.prefix,
            suffix: options.suffix,
            path: fetchargs.path || '',
            method: fetchargs.method || 'GET',
            params: fetchargs.params || {},
            query: fetchargs.query || {},
            headers: prepareHeaders(ctx),
            body: fetchargs.body,
            step: 'start',
        };
        ctx.spec = spec;
        // Merge user-provided headers over SDK defaults.
        if (fetchargs.headers) {
            const uheaders = fetchargs.headers;
            for (let key in uheaders) {
                spec.headers[key] = uheaders[key];
            }
        }
        // Apply SDK auth (apikey, auth prefix, etc.)
        const authResult = prepareAuth(ctx);
        if (authResult instanceof Error) {
            return authResult;
        }
        return makeFetchDef(ctx);
    }
    // Raw endpoint access is operator-controllable, like every entity op.
    // Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
    // either one reaches the same endpoint.
    async direct(fetchargs) {
        if (!this._options.allow.op.includes('direct')) {
            return {
                ok: false,
                err: new Error('NofrixionSDK: direct: operation not allowed by' +
                    ' SDK option allow.op value: "' + this._options.allow.op + '"'),
            };
        }
        return this._rawRequest(fetchargs);
    }
    // Ungated request path shared by direct() and graphql(), each of which
    // checks its own allow.op token first. Private, rather than a flag on
    // fetchargs: a caller-supplied marker would let anyone opt straight back
    // out of the gate by passing it.
    async _rawRequest(fetchargs) {
        const utility = this._utility;
        const fetcher = utility.fetcher;
        const makeContext = utility.makeContext;
        const fetchdef = await this.prepare(fetchargs);
        if (fetchdef instanceof Error) {
            return fetchdef;
        }
        let ctx = makeContext({
            opname: 'direct',
            ctrl: (fetchargs || {}).ctrl || {},
        }, this._rootctx);
        try {
            const fetched = await fetcher(ctx, fetchdef.url, fetchdef);
            if (null == fetched) {
                return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') };
            }
            else if (fetched instanceof Error) {
                return { ok: false, err: fetched };
            }
            const status = fetched.status;
            // No body responses (204 No Content, 304 Not Modified) and explicit
            // zero content-length must skip JSON parsing — fetched.json() would
            // throw `Unexpected end of JSON input` on an empty body.
            const headers = fetched.headers;
            const contentLength = headers && 'function' === typeof headers.get
                ? headers.get('content-length')
                : (headers || {})['content-length'];
            const noBody = 204 === status || 304 === status || '0' === String(contentLength);
            let json = undefined;
            if (!noBody) {
                try {
                    json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json;
                }
                catch (parseErr) {
                    // Body wasn't valid JSON — surface the raw response rather than
                    // throwing. data stays undefined; callers can inspect status/headers.
                    json = undefined;
                }
            }
            return {
                ok: status >= 200 && status < 300,
                status,
                headers: fetched.headers,
                data: json,
            };
        }
        catch (err) {
            return { ok: false, err };
        }
    }
    // Raw GraphQL access: the pressure valve that makes the generated
    // surface's deliberate omissions (per-call selection sets, typed filter
    // builders, batching, subscriptions) livable — the whole schema stays
    // reachable.
    //
    // Thin wrapper over the same prepare/fetch path `direct` uses, with the
    // one thing raw `direct` cannot do for GraphQL: a GraphQL failure rides
    // HTTP 200 as a top-level `errors` array, so status alone would report a
    // failed query as ok.
    //
    // NOTE: like `direct`, this bypasses the feature pipeline — no retry,
    // ratelimit or paging features apply.
    async graphql(query, variables, ctrl) {
        const options = this._options;
        if (!options.allow.op.includes('graphql')) {
            return {
                ok: false,
                err: new Error('NofrixionSDK: graphql: operation not allowed by' +
                    ' SDK option allow.op value: "' + options.allow.op + '"'),
            };
        }
        const res = await this._rawRequest({
            method: 'POST',
            headers: { 'content-type': 'application/json' },
            body: { query, variables: variables || {} },
            ctrl,
        });
        if (res instanceof Error) {
            return res;
        }
        // Errors are read BEFORE any status check: a GraphQL parse or validation
        // failure comes back as HTTP 400 carrying the standard { errors: [...] }
        // body, and the raw path represents a non-2xx as { ok: false } with no
        // err — so returning early on status would discard the server's own
        // diagnostics, which are the only useful part of that response.
        const errors = null == res.data ? undefined : res.data.errors;
        if (null != errors && Array.isArray(errors) && 0 < errors.length) {
            const first = errors[0] || {};
            const err = new Error('NofrixionSDK: graphql: ' +
                (first.message || 'graphql error'));
            err.graphql = errors;
            return { ok: false, status: res.status, headers: res.headers, err, data: res.data };
        }
        return res;
    }
    // Entity access: `client.Account().list()` / `client.Account().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Account(entopts) {
        const self = this;
        return new AccountEntity_1.AccountEntity(self, entopts);
    }
    // Entity access: `client.Batch().list()` / `client.Batch().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Batch(entopts) {
        const self = this;
        return new BatchEntity_1.BatchEntity(self, entopts);
    }
    // Entity access: `client.Beneficiary().list()` / `client.Beneficiary().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Beneficiary(entopts) {
        const self = this;
        return new BeneficiaryEntity_1.BeneficiaryEntity(self, entopts);
    }
    // Entity access: `client.BeneficiaryGroup().list()` / `client.BeneficiaryGroup().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    BeneficiaryGroup(entopts) {
        const self = this;
        return new BeneficiaryGroupEntity_1.BeneficiaryGroupEntity(self, entopts);
    }
    // Entity access: `client.Card().list()` / `client.Card().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Card(entopts) {
        const self = this;
        return new CardEntity_1.CardEntity(self, entopts);
    }
    // Entity access: `client.CardCustomerToken().list()` / `client.CardCustomerToken().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    CardCustomerToken(entopts) {
        const self = this;
        return new CardCustomerTokenEntity_1.CardCustomerTokenEntity(self, entopts);
    }
    // Entity access: `client.CardPayment().list()` / `client.CardPayment().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    CardPayment(entopts) {
        const self = this;
        return new CardPaymentEntity_1.CardPaymentEntity(self, entopts);
    }
    // Entity access: `client.CardPublicKey().list()` / `client.CardPublicKey().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    CardPublicKey(entopts) {
        const self = this;
        return new CardPublicKeyEntity_1.CardPublicKeyEntity(self, entopts);
    }
    // Entity access: `client.Consent().list()` / `client.Consent().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Consent(entopts) {
        const self = this;
        return new ConsentEntity_1.ConsentEntity(self, entopts);
    }
    // Entity access: `client.Currency().list()` / `client.Currency().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Currency(entopts) {
        const self = this;
        return new CurrencyEntity_1.CurrencyEntity(self, entopts);
    }
    // Entity access: `client.DirectDebitBatchSubmit().list()` / `client.DirectDebitBatchSubmit().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    DirectDebitBatchSubmit(entopts) {
        const self = this;
        return new DirectDebitBatchSubmitEntity_1.DirectDebitBatchSubmitEntity(self, entopts);
    }
    // Entity access: `client.FxRate().list()` / `client.FxRate().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    FxRate(entopts) {
        const self = this;
        return new FxRateEntity_1.FxRateEntity(self, entopts);
    }
    // Entity access: `client.IPayment().list()` / `client.IPayment().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    IPayment(entopts) {
        const self = this;
        return new IPaymentEntity_1.IPaymentEntity(self, entopts);
    }
    // Entity access: `client.Mandate().list()` / `client.Mandate().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Mandate(entopts) {
        const self = this;
        return new MandateEntity_1.MandateEntity(self, entopts);
    }
    // Entity access: `client.Merchant().list()` / `client.Merchant().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Merchant(entopts) {
        const self = this;
        return new MerchantEntity_1.MerchantEntity(self, entopts);
    }
    // Entity access: `client.MerchantAuthorisationSetting().list()` / `client.MerchantAuthorisationSetting().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    MerchantAuthorisationSetting(entopts) {
        const self = this;
        return new MerchantAuthorisationSettingEntity_1.MerchantAuthorisationSettingEntity(self, entopts);
    }
    // Entity access: `client.MerchantDirectDebitMandatePage().list()` / `client.MerchantDirectDebitMandatePage().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    MerchantDirectDebitMandatePage(entopts) {
        const self = this;
        return new MerchantDirectDebitMandatePageEntity_1.MerchantDirectDebitMandatePageEntity(self, entopts);
    }
    // Entity access: `client.MerchantPayByBankSetting().list()` / `client.MerchantPayByBankSetting().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    MerchantPayByBankSetting(entopts) {
        const self = this;
        return new MerchantPayByBankSettingEntity_1.MerchantPayByBankSettingEntity(self, entopts);
    }
    // Entity access: `client.MerchantPaymentRequestTemplate().list()` / `client.MerchantPaymentRequestTemplate().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    MerchantPaymentRequestTemplate(entopts) {
        const self = this;
        return new MerchantPaymentRequestTemplateEntity_1.MerchantPaymentRequestTemplateEntity(self, entopts);
    }
    // Entity access: `client.MerchantToken().list()` / `client.MerchantToken().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    MerchantToken(entopts) {
        const self = this;
        return new MerchantTokenEntity_1.MerchantTokenEntity(self, entopts);
    }
    // Entity access: `client.Metadata().list()` / `client.Metadata().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Metadata(entopts) {
        const self = this;
        return new MetadataEntity_1.MetadataEntity(self, entopts);
    }
    // Entity access: `client.NoFrixionVersion().list()` / `client.NoFrixionVersion().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    NoFrixionVersion(entopts) {
        const self = this;
        return new NoFrixionVersionEntity_1.NoFrixionVersionEntity(self, entopts);
    }
    // Entity access: `client.OpenBanking().list()` / `client.OpenBanking().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    OpenBanking(entopts) {
        const self = this;
        return new OpenBankingEntity_1.OpenBankingEntity(self, entopts);
    }
    // Entity access: `client.Payeeverification().list()` / `client.Payeeverification().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Payeeverification(entopts) {
        const self = this;
        return new PayeeverificationEntity_1.PayeeverificationEntity(self, entopts);
    }
    // Entity access: `client.Payment().list()` / `client.Payment().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Payment(entopts) {
        const self = this;
        return new PaymentEntity_1.PaymentEntity(self, entopts);
    }
    // Entity access: `client.PaymentAccount().list()` / `client.PaymentAccount().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PaymentAccount(entopts) {
        const self = this;
        return new PaymentAccountEntity_1.PaymentAccountEntity(self, entopts);
    }
    // Entity access: `client.PaymentAccountMinimal().list()` / `client.PaymentAccountMinimal().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PaymentAccountMinimal(entopts) {
        const self = this;
        return new PaymentAccountMinimalEntity_1.PaymentAccountMinimalEntity(self, entopts);
    }
    // Entity access: `client.PaymentInitiation().list()` / `client.PaymentInitiation().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PaymentInitiation(entopts) {
        const self = this;
        return new PaymentInitiationEntity_1.PaymentInitiationEntity(self, entopts);
    }
    // Entity access: `client.PaymentRequest().list()` / `client.PaymentRequest().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PaymentRequest(entopts) {
        const self = this;
        return new PaymentRequestEntity_1.PaymentRequestEntity(self, entopts);
    }
    // Entity access: `client.PaymentRequestEvent().list()` / `client.PaymentRequestEvent().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PaymentRequestEvent(entopts) {
        const self = this;
        return new PaymentRequestEventEntity_1.PaymentRequestEventEntity(self, entopts);
    }
    // Entity access: `client.PaymentRequestMetric().list()` / `client.PaymentRequestMetric().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PaymentRequestMetric(entopts) {
        const self = this;
        return new PaymentRequestMetricEntity_1.PaymentRequestMetricEntity(self, entopts);
    }
    // Entity access: `client.PaymentRequestMinimal().list()` / `client.PaymentRequestMinimal().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PaymentRequestMinimal(entopts) {
        const self = this;
        return new PaymentRequestMinimalEntity_1.PaymentRequestMinimalEntity(self, entopts);
    }
    // Entity access: `client.PaymentRequestResult().list()` / `client.PaymentRequestResult().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PaymentRequestResult(entopts) {
        const self = this;
        return new PaymentRequestResultEntity_1.PaymentRequestResultEntity(self, entopts);
    }
    // Entity access: `client.Payout().list()` / `client.Payout().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Payout(entopts) {
        const self = this;
        return new PayoutEntity_1.PayoutEntity(self, entopts);
    }
    // Entity access: `client.PayoutKeysetPage().list()` / `client.PayoutKeysetPage().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PayoutKeysetPage(entopts) {
        const self = this;
        return new PayoutKeysetPageEntity_1.PayoutKeysetPageEntity(self, entopts);
    }
    // Entity access: `client.PayoutMetric().list()` / `client.PayoutMetric().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PayoutMetric(entopts) {
        const self = this;
        return new PayoutMetricEntity_1.PayoutMetricEntity(self, entopts);
    }
    // Entity access: `client.Payrun().list()` / `client.Payrun().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Payrun(entopts) {
        const self = this;
        return new PayrunEntity_1.PayrunEntity(self, entopts);
    }
    // Entity access: `client.Report().list()` / `client.Report().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Report(entopts) {
        const self = this;
        return new ReportEntity_1.ReportEntity(self, entopts);
    }
    // Entity access: `client.ReportResult().list()` / `client.ReportResult().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ReportResult(entopts) {
        const self = this;
        return new ReportResultEntity_1.ReportResultEntity(self, entopts);
    }
    // Entity access: `client.Role().list()` / `client.Role().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Role(entopts) {
        const self = this;
        return new RoleEntity_1.RoleEntity(self, entopts);
    }
    // Entity access: `client.Rule().list()` / `client.Rule().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Rule(entopts) {
        const self = this;
        return new RuleEntity_1.RuleEntity(self, entopts);
    }
    // Entity access: `client.RuleEvent().list()` / `client.RuleEvent().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    RuleEvent(entopts) {
        const self = this;
        return new RuleEventEntity_1.RuleEventEntity(self, entopts);
    }
    // Entity access: `client.Tag().list()` / `client.Tag().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Tag(entopts) {
        const self = this;
        return new TagEntity_1.TagEntity(self, entopts);
    }
    // Entity access: `client.Token().list()` / `client.Token().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Token(entopts) {
        const self = this;
        return new TokenEntity_1.TokenEntity(self, entopts);
    }
    // Entity access: `client.Transaction().list()` / `client.Transaction().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Transaction(entopts) {
        const self = this;
        return new TransactionEntity_1.TransactionEntity(self, entopts);
    }
    // Entity access: `client.User().list()` / `client.User().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    User(entopts) {
        const self = this;
        return new UserEntity_1.UserEntity(self, entopts);
    }
    // Entity access: `client.UserInvite().list()` / `client.UserInvite().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    UserInvite(entopts) {
        const self = this;
        return new UserInviteEntity_1.UserInviteEntity(self, entopts);
    }
    // Entity access: `client.Virtual().list()` / `client.Virtual().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Virtual(entopts) {
        const self = this;
        return new VirtualEntity_1.VirtualEntity(self, entopts);
    }
    // Entity access: `client.Webhook().list()` / `client.Webhook().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Webhook(entopts) {
        const self = this;
        return new WebhookEntity_1.WebhookEntity(self, entopts);
    }
    static test(testoptsarg, sdkoptsarg) {
        const struct = stdutil.struct;
        const setpath = struct.setpath;
        const getdef = struct.getdef;
        const clone = struct.clone;
        const setprop = struct.setprop;
        const sdkopts = getdef(clone(sdkoptsarg), {});
        const testopts = getdef(clone(testoptsarg), {});
        setprop(testopts, 'active', true);
        setpath(sdkopts, 'feature.test', testopts);
        const testsdk = new NofrixionSDK(sdkopts);
        testsdk._mode = 'test';
        return testsdk;
    }
    tester(testopts, sdkopts) {
        return NofrixionSDK.test(testopts, sdkopts);
    }
    toJSON() {
        return { name: 'Nofrixion' };
    }
    toString() {
        return 'Nofrixion ' + this._utility.struct.jsonify(this.toJSON());
    }
    [node_util_1.inspect.custom]() {
        return this.toString();
    }
}
exports.NofrixionSDK = NofrixionSDK;
const SDK = NofrixionSDK;
exports.SDK = SDK;
//# sourceMappingURL=NofrixionSDK.js.map