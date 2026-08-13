// Nofrixion Ts SDK

import { AccountEntity } from './entity/AccountEntity'
import { BatchEntity } from './entity/BatchEntity'
import { BeneficiaryEntity } from './entity/BeneficiaryEntity'
import { BeneficiaryGroupEntity } from './entity/BeneficiaryGroupEntity'
import { CardEntity } from './entity/CardEntity'
import { CardCustomerTokenEntity } from './entity/CardCustomerTokenEntity'
import { CardPaymentEntity } from './entity/CardPaymentEntity'
import { CardPublicKeyEntity } from './entity/CardPublicKeyEntity'
import { ConsentEntity } from './entity/ConsentEntity'
import { CurrencyEntity } from './entity/CurrencyEntity'
import { DirectDebitBatchSubmitEntity } from './entity/DirectDebitBatchSubmitEntity'
import { FxRateEntity } from './entity/FxRateEntity'
import { IPaymentEntity } from './entity/IPaymentEntity'
import { MandateEntity } from './entity/MandateEntity'
import { MerchantEntity } from './entity/MerchantEntity'
import { MerchantAuthorisationSettingEntity } from './entity/MerchantAuthorisationSettingEntity'
import { MerchantDirectDebitMandatePageEntity } from './entity/MerchantDirectDebitMandatePageEntity'
import { MerchantPayByBankSettingEntity } from './entity/MerchantPayByBankSettingEntity'
import { MerchantPaymentRequestTemplateEntity } from './entity/MerchantPaymentRequestTemplateEntity'
import { MerchantTokenEntity } from './entity/MerchantTokenEntity'
import { MetadataEntity } from './entity/MetadataEntity'
import { NoFrixionVersionEntity } from './entity/NoFrixionVersionEntity'
import { OpenBankingEntity } from './entity/OpenBankingEntity'
import { PayeeverificationEntity } from './entity/PayeeverificationEntity'
import { PaymentEntity } from './entity/PaymentEntity'
import { PaymentAccountEntity } from './entity/PaymentAccountEntity'
import { PaymentAccountMinimalEntity } from './entity/PaymentAccountMinimalEntity'
import { PaymentInitiationEntity } from './entity/PaymentInitiationEntity'
import { PaymentRequestEntity } from './entity/PaymentRequestEntity'
import { PaymentRequestEventEntity } from './entity/PaymentRequestEventEntity'
import { PaymentRequestMetricEntity } from './entity/PaymentRequestMetricEntity'
import { PaymentRequestMinimalEntity } from './entity/PaymentRequestMinimalEntity'
import { PaymentRequestResultEntity } from './entity/PaymentRequestResultEntity'
import { PayoutEntity } from './entity/PayoutEntity'
import { PayoutKeysetPageEntity } from './entity/PayoutKeysetPageEntity'
import { PayoutMetricEntity } from './entity/PayoutMetricEntity'
import { PayrunEntity } from './entity/PayrunEntity'
import { ReportEntity } from './entity/ReportEntity'
import { ReportResultEntity } from './entity/ReportResultEntity'
import { RoleEntity } from './entity/RoleEntity'
import { RuleEntity } from './entity/RuleEntity'
import { RuleEventEntity } from './entity/RuleEventEntity'
import { TagEntity } from './entity/TagEntity'
import { TokenEntity } from './entity/TokenEntity'
import { TransactionEntity } from './entity/TransactionEntity'
import { UserEntity } from './entity/UserEntity'
import { UserInviteEntity } from './entity/UserInviteEntity'
import { VirtualEntity } from './entity/VirtualEntity'
import { WebhookEntity } from './entity/WebhookEntity'

export type * from './NofrixionTypes'


import { inspect } from 'node:util'

import type { Context, Feature } from './types'

import { config } from './Config'
import { NofrixionEntityBase } from './NofrixionEntityBase'
import { Utility } from './utility/Utility'


import { BaseFeature } from './feature/base/BaseFeature'


const stdutil = new Utility()


class NofrixionSDK {
  _mode: string = 'live'
  _options: any
  _utility = new Utility()
  _features: Feature[]
  _rootctx: Context

  constructor(options?: any) {

    this._rootctx = this._utility.makeContext({
      client: this,
      utility: this._utility,
      config,
      options,
      shared: new WeakMap()
    })

    this._options = this._utility.makeOptions(this._rootctx)

    const struct = this._utility.struct
    const getpath = struct.getpath

    if (true === getpath(this._options.feature, 'test.active')) {
      this._mode = 'test'
    }

    this._rootctx.options = this._options

    this._features = []

    const featureAdd = this._utility.featureAdd
    const featureInit = this._utility.featureInit

    // Add features in the resolved order (makeOptions puts an explicit
    // array order first, else defaults to test-first). Ordering matters:
    // the `test` feature installs the base mock transport and the transport
    // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is current,
    // so `test` must be added before them to sit at the base of the chain.
    const featureorder = getpath(this._options, '__derived__.featureorder') || []
    for (const fname of featureorder) {
      const fopts = this._options.feature[fname] || {}
      if (fopts.active) {
        featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname))
      }
    }

    if (null != this._options.extend) {
      for (let f of this._options.extend) {
        featureAdd(this._rootctx, f)
      }
    }

    for (let f of this._features) {
      featureInit(this._rootctx, f)
    }

    const featureHook = this._utility.featureHook
    featureHook(this._rootctx, 'PostConstruct')
  }


  options() {
    return this._utility.struct.clone(this._options)
  }


  utility() {
    return this._utility.struct.clone(this._utility)
  }


  async prepare(fetchargs?: any) {
    const utility = this._utility
    const struct = utility.struct
    const clone = struct.clone

    const {
      makeContext,
      makeFetchDef,
      prepareHeaders,
      prepareAuth,
    } = utility

    fetchargs = fetchargs || {}

    let ctx: Context = makeContext({
      opname: 'prepare',
      ctrl: fetchargs.ctrl || {},
    }, this._rootctx)

    const options = this._options

    // Build spec directly from SDK options + user-provided fetch args.
    const spec: any = {
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
    }

    ctx.spec = spec

    // Merge user-provided headers over SDK defaults.
    if (fetchargs.headers) {
      const uheaders = fetchargs.headers
      for (let key in uheaders) {
        spec.headers[key] = uheaders[key]
      }
    }

    // Apply SDK auth (apikey, auth prefix, etc.)
    const authResult = prepareAuth(ctx)
    if (authResult instanceof Error) {
      return authResult
    }

    return makeFetchDef(ctx)
  }


  // Raw endpoint access is operator-controllable, like every entity op.
  // Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
  // either one reaches the same endpoint.
  async direct(fetchargs?: any) {
    if (!this._options.allow.op.includes('direct')) {
      return {
        ok: false,
        err: new Error('NofrixionSDK: direct: operation not allowed by' +
          ' SDK option allow.op value: "' + this._options.allow.op + '"'),
      }
    }

    return this._rawRequest(fetchargs)
  }


  // Ungated request path shared by direct() and graphql(), each of which
  // checks its own allow.op token first. Private, rather than a flag on
  // fetchargs: a caller-supplied marker would let anyone opt straight back
  // out of the gate by passing it.
  async _rawRequest(fetchargs?: any) {
    const utility = this._utility

    const fetcher = utility.fetcher
    const makeContext = utility.makeContext

    const fetchdef = await this.prepare(fetchargs)
    if (fetchdef instanceof Error) {
      return fetchdef
    }

    let ctx: Context = makeContext({
      opname: 'direct',
      ctrl: (fetchargs || {}).ctrl || {},
    }, this._rootctx)

    try {
      const fetched = await fetcher(ctx, fetchdef.url, fetchdef)

      if (null == fetched) {
        return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') }
      }
      else if (fetched instanceof Error) {
        return { ok: false, err: fetched }
      }

      const status = fetched.status

      // No body responses (204 No Content, 304 Not Modified) and explicit
      // zero content-length must skip JSON parsing — fetched.json() would
      // throw `Unexpected end of JSON input` on an empty body.
      const headers = fetched.headers
      const contentLength = headers && 'function' === typeof headers.get
        ? headers.get('content-length')
        : (headers || {})['content-length']
      const noBody = 204 === status || 304 === status || '0' === String(contentLength)

      let json: any = undefined
      if (!noBody) {
        try {
          json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json
        }
        catch (parseErr) {
          // Body wasn't valid JSON — surface the raw response rather than
          // throwing. data stays undefined; callers can inspect status/headers.
          json = undefined
        }
      }

      return {
        ok: status >= 200 && status < 300,
        status,
        headers: fetched.headers,
        data: json,
      }
    }
    catch (err: any) {
      return { ok: false, err }
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
  async graphql(query: string, variables?: any, ctrl?: any) {
    const options = this._options

    if (!options.allow.op.includes('graphql')) {
      return {
        ok: false,
        err: new Error('NofrixionSDK: graphql: operation not allowed by' +
          ' SDK option allow.op value: "' + options.allow.op + '"'),
      }
    }

    const res: any = await this._rawRequest({
      method: 'POST',
      headers: { 'content-type': 'application/json' },
      body: { query, variables: variables || {} },
      ctrl,
    })

    if (res instanceof Error) {
      return res
    }

    // Errors are read BEFORE any status check: a GraphQL parse or validation
    // failure comes back as HTTP 400 carrying the standard { errors: [...] }
    // body, and the raw path represents a non-2xx as { ok: false } with no
    // err — so returning early on status would discard the server's own
    // diagnostics, which are the only useful part of that response.
    const errors = null == res.data ? undefined : res.data.errors

    if (null != errors && Array.isArray(errors) && 0 < errors.length) {
      const first = errors[0] || {}
      const err: any = new Error('NofrixionSDK: graphql: ' +
        (first.message || 'graphql error'))
      err.graphql = errors
      return { ok: false, status: res.status, headers: res.headers, err, data: res.data }
    }

    return res
  }



  // Entity access: `client.Account().list()` / `client.Account().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Account(entopts?: Record<string, any>) {
    const self = this
    return new AccountEntity(self, entopts)
  }


  // Entity access: `client.Batch().list()` / `client.Batch().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Batch(entopts?: Record<string, any>) {
    const self = this
    return new BatchEntity(self, entopts)
  }


  // Entity access: `client.Beneficiary().list()` / `client.Beneficiary().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Beneficiary(entopts?: Record<string, any>) {
    const self = this
    return new BeneficiaryEntity(self, entopts)
  }


  // Entity access: `client.BeneficiaryGroup().list()` / `client.BeneficiaryGroup().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  BeneficiaryGroup(entopts?: Record<string, any>) {
    const self = this
    return new BeneficiaryGroupEntity(self, entopts)
  }


  // Entity access: `client.Card().list()` / `client.Card().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Card(entopts?: Record<string, any>) {
    const self = this
    return new CardEntity(self, entopts)
  }


  // Entity access: `client.CardCustomerToken().list()` / `client.CardCustomerToken().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CardCustomerToken(entopts?: Record<string, any>) {
    const self = this
    return new CardCustomerTokenEntity(self, entopts)
  }


  // Entity access: `client.CardPayment().list()` / `client.CardPayment().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CardPayment(entopts?: Record<string, any>) {
    const self = this
    return new CardPaymentEntity(self, entopts)
  }


  // Entity access: `client.CardPublicKey().list()` / `client.CardPublicKey().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  CardPublicKey(entopts?: Record<string, any>) {
    const self = this
    return new CardPublicKeyEntity(self, entopts)
  }


  // Entity access: `client.Consent().list()` / `client.Consent().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Consent(entopts?: Record<string, any>) {
    const self = this
    return new ConsentEntity(self, entopts)
  }


  // Entity access: `client.Currency().list()` / `client.Currency().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Currency(entopts?: Record<string, any>) {
    const self = this
    return new CurrencyEntity(self, entopts)
  }


  // Entity access: `client.DirectDebitBatchSubmit().list()` / `client.DirectDebitBatchSubmit().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  DirectDebitBatchSubmit(entopts?: Record<string, any>) {
    const self = this
    return new DirectDebitBatchSubmitEntity(self, entopts)
  }


  // Entity access: `client.FxRate().list()` / `client.FxRate().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  FxRate(entopts?: Record<string, any>) {
    const self = this
    return new FxRateEntity(self, entopts)
  }


  // Entity access: `client.IPayment().list()` / `client.IPayment().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  IPayment(entopts?: Record<string, any>) {
    const self = this
    return new IPaymentEntity(self, entopts)
  }


  // Entity access: `client.Mandate().list()` / `client.Mandate().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Mandate(entopts?: Record<string, any>) {
    const self = this
    return new MandateEntity(self, entopts)
  }


  // Entity access: `client.Merchant().list()` / `client.Merchant().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Merchant(entopts?: Record<string, any>) {
    const self = this
    return new MerchantEntity(self, entopts)
  }


  // Entity access: `client.MerchantAuthorisationSetting().list()` / `client.MerchantAuthorisationSetting().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  MerchantAuthorisationSetting(entopts?: Record<string, any>) {
    const self = this
    return new MerchantAuthorisationSettingEntity(self, entopts)
  }


  // Entity access: `client.MerchantDirectDebitMandatePage().list()` / `client.MerchantDirectDebitMandatePage().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  MerchantDirectDebitMandatePage(entopts?: Record<string, any>) {
    const self = this
    return new MerchantDirectDebitMandatePageEntity(self, entopts)
  }


  // Entity access: `client.MerchantPayByBankSetting().list()` / `client.MerchantPayByBankSetting().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  MerchantPayByBankSetting(entopts?: Record<string, any>) {
    const self = this
    return new MerchantPayByBankSettingEntity(self, entopts)
  }


  // Entity access: `client.MerchantPaymentRequestTemplate().list()` / `client.MerchantPaymentRequestTemplate().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  MerchantPaymentRequestTemplate(entopts?: Record<string, any>) {
    const self = this
    return new MerchantPaymentRequestTemplateEntity(self, entopts)
  }


  // Entity access: `client.MerchantToken().list()` / `client.MerchantToken().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  MerchantToken(entopts?: Record<string, any>) {
    const self = this
    return new MerchantTokenEntity(self, entopts)
  }


  // Entity access: `client.Metadata().list()` / `client.Metadata().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Metadata(entopts?: Record<string, any>) {
    const self = this
    return new MetadataEntity(self, entopts)
  }


  // Entity access: `client.NoFrixionVersion().list()` / `client.NoFrixionVersion().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  NoFrixionVersion(entopts?: Record<string, any>) {
    const self = this
    return new NoFrixionVersionEntity(self, entopts)
  }


  // Entity access: `client.OpenBanking().list()` / `client.OpenBanking().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  OpenBanking(entopts?: Record<string, any>) {
    const self = this
    return new OpenBankingEntity(self, entopts)
  }


  // Entity access: `client.Payeeverification().list()` / `client.Payeeverification().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Payeeverification(entopts?: Record<string, any>) {
    const self = this
    return new PayeeverificationEntity(self, entopts)
  }


  // Entity access: `client.Payment().list()` / `client.Payment().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Payment(entopts?: Record<string, any>) {
    const self = this
    return new PaymentEntity(self, entopts)
  }


  // Entity access: `client.PaymentAccount().list()` / `client.PaymentAccount().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PaymentAccount(entopts?: Record<string, any>) {
    const self = this
    return new PaymentAccountEntity(self, entopts)
  }


  // Entity access: `client.PaymentAccountMinimal().list()` / `client.PaymentAccountMinimal().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PaymentAccountMinimal(entopts?: Record<string, any>) {
    const self = this
    return new PaymentAccountMinimalEntity(self, entopts)
  }


  // Entity access: `client.PaymentInitiation().list()` / `client.PaymentInitiation().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PaymentInitiation(entopts?: Record<string, any>) {
    const self = this
    return new PaymentInitiationEntity(self, entopts)
  }


  // Entity access: `client.PaymentRequest().list()` / `client.PaymentRequest().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PaymentRequest(entopts?: Record<string, any>) {
    const self = this
    return new PaymentRequestEntity(self, entopts)
  }


  // Entity access: `client.PaymentRequestEvent().list()` / `client.PaymentRequestEvent().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PaymentRequestEvent(entopts?: Record<string, any>) {
    const self = this
    return new PaymentRequestEventEntity(self, entopts)
  }


  // Entity access: `client.PaymentRequestMetric().list()` / `client.PaymentRequestMetric().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PaymentRequestMetric(entopts?: Record<string, any>) {
    const self = this
    return new PaymentRequestMetricEntity(self, entopts)
  }


  // Entity access: `client.PaymentRequestMinimal().list()` / `client.PaymentRequestMinimal().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PaymentRequestMinimal(entopts?: Record<string, any>) {
    const self = this
    return new PaymentRequestMinimalEntity(self, entopts)
  }


  // Entity access: `client.PaymentRequestResult().list()` / `client.PaymentRequestResult().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PaymentRequestResult(entopts?: Record<string, any>) {
    const self = this
    return new PaymentRequestResultEntity(self, entopts)
  }


  // Entity access: `client.Payout().list()` / `client.Payout().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Payout(entopts?: Record<string, any>) {
    const self = this
    return new PayoutEntity(self, entopts)
  }


  // Entity access: `client.PayoutKeysetPage().list()` / `client.PayoutKeysetPage().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PayoutKeysetPage(entopts?: Record<string, any>) {
    const self = this
    return new PayoutKeysetPageEntity(self, entopts)
  }


  // Entity access: `client.PayoutMetric().list()` / `client.PayoutMetric().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PayoutMetric(entopts?: Record<string, any>) {
    const self = this
    return new PayoutMetricEntity(self, entopts)
  }


  // Entity access: `client.Payrun().list()` / `client.Payrun().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Payrun(entopts?: Record<string, any>) {
    const self = this
    return new PayrunEntity(self, entopts)
  }


  // Entity access: `client.Report().list()` / `client.Report().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Report(entopts?: Record<string, any>) {
    const self = this
    return new ReportEntity(self, entopts)
  }


  // Entity access: `client.ReportResult().list()` / `client.ReportResult().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ReportResult(entopts?: Record<string, any>) {
    const self = this
    return new ReportResultEntity(self, entopts)
  }


  // Entity access: `client.Role().list()` / `client.Role().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Role(entopts?: Record<string, any>) {
    const self = this
    return new RoleEntity(self, entopts)
  }


  // Entity access: `client.Rule().list()` / `client.Rule().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Rule(entopts?: Record<string, any>) {
    const self = this
    return new RuleEntity(self, entopts)
  }


  // Entity access: `client.RuleEvent().list()` / `client.RuleEvent().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  RuleEvent(entopts?: Record<string, any>) {
    const self = this
    return new RuleEventEntity(self, entopts)
  }


  // Entity access: `client.Tag().list()` / `client.Tag().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Tag(entopts?: Record<string, any>) {
    const self = this
    return new TagEntity(self, entopts)
  }


  // Entity access: `client.Token().list()` / `client.Token().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Token(entopts?: Record<string, any>) {
    const self = this
    return new TokenEntity(self, entopts)
  }


  // Entity access: `client.Transaction().list()` / `client.Transaction().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Transaction(entopts?: Record<string, any>) {
    const self = this
    return new TransactionEntity(self, entopts)
  }


  // Entity access: `client.User().list()` / `client.User().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  User(entopts?: Record<string, any>) {
    const self = this
    return new UserEntity(self, entopts)
  }


  // Entity access: `client.UserInvite().list()` / `client.UserInvite().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  UserInvite(entopts?: Record<string, any>) {
    const self = this
    return new UserInviteEntity(self, entopts)
  }


  // Entity access: `client.Virtual().list()` / `client.Virtual().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Virtual(entopts?: Record<string, any>) {
    const self = this
    return new VirtualEntity(self, entopts)
  }


  // Entity access: `client.Webhook().list()` / `client.Webhook().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Webhook(entopts?: Record<string, any>) {
    const self = this
    return new WebhookEntity(self, entopts)
  }




  static test(testoptsarg?: any, sdkoptsarg?: any) {
    const struct = stdutil.struct
    const setpath = struct.setpath
    const getdef = struct.getdef
    const clone = struct.clone
    const setprop = struct.setprop

    const sdkopts = getdef(clone(sdkoptsarg), {})
    const testopts = getdef(clone(testoptsarg), {})
    setprop(testopts, 'active', true)
    setpath(sdkopts, 'feature.test', testopts)

    const testsdk = new NofrixionSDK(sdkopts)
    testsdk._mode = 'test'

    return testsdk
  }


  tester(testopts?: any, sdkopts?: any) {
    return NofrixionSDK.test(testopts, sdkopts)
  }


  toJSON() {
    return { name: 'Nofrixion' }
  }

  toString() {
    return 'Nofrixion ' + this._utility.struct.jsonify(this.toJSON())
  }

  [inspect.custom]() {
    return this.toString()
  }

}




const SDK = NofrixionSDK


export {
  stdutil,
  config,

  BaseFeature,
  NofrixionEntityBase,

  NofrixionSDK,
  SDK,
}


