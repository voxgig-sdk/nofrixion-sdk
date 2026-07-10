// Nofrixion Ts SDK

import { AccountEntity } from './entity/AccountEntity'
import { BatchEntity } from './entity/BatchEntity'
import { BeneficiariesCreateEntity } from './entity/BeneficiariesCreateEntity'
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
import { MerchantDirectDebitMandateEntity } from './entity/MerchantDirectDebitMandateEntity'
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
import { PaymentRequestsCreateEntity } from './entity/PaymentRequestsCreateEntity'
import { PayoutEntity } from './entity/PayoutEntity'
import { PayoutKeysetEntity } from './entity/PayoutKeysetEntity'
import { PayoutMetricEntity } from './entity/PayoutMetricEntity'
import { PayoutsCreateEntity } from './entity/PayoutsCreateEntity'
import { PayrunEntity } from './entity/PayrunEntity'
import { ReportEntity } from './entity/ReportEntity'
import { ReportResultEntity } from './entity/ReportResultEntity'
import { RolesCreateEntity } from './entity/RolesCreateEntity'
import { RuleEntity } from './entity/RuleEntity'
import { RuleEventEntity } from './entity/RuleEventEntity'
import { TagEntity } from './entity/TagEntity'
import { TokenEntity } from './entity/TokenEntity'
import { TransactionEntity } from './entity/TransactionEntity'
import { UserEntity } from './entity/UserEntity'
import { UserInviteEntity } from './entity/UserInviteEntity'
import { UserInvitesCreateEntity } from './entity/UserInvitesCreateEntity'
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
    const items = struct.items

    if (true === getpath(this._options.feature, 'test.active')) {
      this._mode = 'test'
    }

    this._rootctx.options = this._options

    this._features = []

    const featureAdd = this._utility.featureAdd
    const featureInit = this._utility.featureInit

    items(this._options.feature, (fitem: [string, any]) => {
      const fname = fitem[0]
      const fopts = fitem[1]
      if (fopts.active) {
        featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname))
      }
    })

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


  async direct(fetchargs?: any) {
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



  // Entity access: `client.Account().list()` / `client.Account().load({ id })`.
  Account(data?: any) {
    const self = this
    return new AccountEntity(self,data)
  }


  // Entity access: `client.Batch().list()` / `client.Batch().load({ id })`.
  Batch(data?: any) {
    const self = this
    return new BatchEntity(self,data)
  }


  // Entity access: `client.BeneficiariesCreate().list()` / `client.BeneficiariesCreate().load({ id })`.
  BeneficiariesCreate(data?: any) {
    const self = this
    return new BeneficiariesCreateEntity(self,data)
  }


  // Entity access: `client.Beneficiary().list()` / `client.Beneficiary().load({ id })`.
  Beneficiary(data?: any) {
    const self = this
    return new BeneficiaryEntity(self,data)
  }


  // Entity access: `client.BeneficiaryGroup().list()` / `client.BeneficiaryGroup().load({ id })`.
  BeneficiaryGroup(data?: any) {
    const self = this
    return new BeneficiaryGroupEntity(self,data)
  }


  // Entity access: `client.Card().list()` / `client.Card().load({ id })`.
  Card(data?: any) {
    const self = this
    return new CardEntity(self,data)
  }


  // Entity access: `client.CardCustomerToken().list()` / `client.CardCustomerToken().load({ id })`.
  CardCustomerToken(data?: any) {
    const self = this
    return new CardCustomerTokenEntity(self,data)
  }


  // Entity access: `client.CardPayment().list()` / `client.CardPayment().load({ id })`.
  CardPayment(data?: any) {
    const self = this
    return new CardPaymentEntity(self,data)
  }


  // Entity access: `client.CardPublicKey().list()` / `client.CardPublicKey().load({ id })`.
  CardPublicKey(data?: any) {
    const self = this
    return new CardPublicKeyEntity(self,data)
  }


  // Entity access: `client.Consent().list()` / `client.Consent().load({ id })`.
  Consent(data?: any) {
    const self = this
    return new ConsentEntity(self,data)
  }


  // Entity access: `client.Currency().list()` / `client.Currency().load({ id })`.
  Currency(data?: any) {
    const self = this
    return new CurrencyEntity(self,data)
  }


  // Entity access: `client.DirectDebitBatchSubmit().list()` / `client.DirectDebitBatchSubmit().load({ id })`.
  DirectDebitBatchSubmit(data?: any) {
    const self = this
    return new DirectDebitBatchSubmitEntity(self,data)
  }


  // Entity access: `client.FxRate().list()` / `client.FxRate().load({ id })`.
  FxRate(data?: any) {
    const self = this
    return new FxRateEntity(self,data)
  }


  // Entity access: `client.IPayment().list()` / `client.IPayment().load({ id })`.
  IPayment(data?: any) {
    const self = this
    return new IPaymentEntity(self,data)
  }


  // Entity access: `client.Mandate().list()` / `client.Mandate().load({ id })`.
  Mandate(data?: any) {
    const self = this
    return new MandateEntity(self,data)
  }


  // Entity access: `client.Merchant().list()` / `client.Merchant().load({ id })`.
  Merchant(data?: any) {
    const self = this
    return new MerchantEntity(self,data)
  }


  // Entity access: `client.MerchantAuthorisationSetting().list()` / `client.MerchantAuthorisationSetting().load({ id })`.
  MerchantAuthorisationSetting(data?: any) {
    const self = this
    return new MerchantAuthorisationSettingEntity(self,data)
  }


  // Entity access: `client.MerchantDirectDebitMandate().list()` / `client.MerchantDirectDebitMandate().load({ id })`.
  MerchantDirectDebitMandate(data?: any) {
    const self = this
    return new MerchantDirectDebitMandateEntity(self,data)
  }


  // Entity access: `client.MerchantPayByBankSetting().list()` / `client.MerchantPayByBankSetting().load({ id })`.
  MerchantPayByBankSetting(data?: any) {
    const self = this
    return new MerchantPayByBankSettingEntity(self,data)
  }


  // Entity access: `client.MerchantPaymentRequestTemplate().list()` / `client.MerchantPaymentRequestTemplate().load({ id })`.
  MerchantPaymentRequestTemplate(data?: any) {
    const self = this
    return new MerchantPaymentRequestTemplateEntity(self,data)
  }


  // Entity access: `client.MerchantToken().list()` / `client.MerchantToken().load({ id })`.
  MerchantToken(data?: any) {
    const self = this
    return new MerchantTokenEntity(self,data)
  }


  // Entity access: `client.Metadata().list()` / `client.Metadata().load({ id })`.
  Metadata(data?: any) {
    const self = this
    return new MetadataEntity(self,data)
  }


  // Entity access: `client.NoFrixionVersion().list()` / `client.NoFrixionVersion().load({ id })`.
  NoFrixionVersion(data?: any) {
    const self = this
    return new NoFrixionVersionEntity(self,data)
  }


  // Entity access: `client.OpenBanking().list()` / `client.OpenBanking().load({ id })`.
  OpenBanking(data?: any) {
    const self = this
    return new OpenBankingEntity(self,data)
  }


  // Entity access: `client.Payeeverification().list()` / `client.Payeeverification().load({ id })`.
  Payeeverification(data?: any) {
    const self = this
    return new PayeeverificationEntity(self,data)
  }


  // Entity access: `client.Payment().list()` / `client.Payment().load({ id })`.
  Payment(data?: any) {
    const self = this
    return new PaymentEntity(self,data)
  }


  // Entity access: `client.PaymentAccount().list()` / `client.PaymentAccount().load({ id })`.
  PaymentAccount(data?: any) {
    const self = this
    return new PaymentAccountEntity(self,data)
  }


  // Entity access: `client.PaymentAccountMinimal().list()` / `client.PaymentAccountMinimal().load({ id })`.
  PaymentAccountMinimal(data?: any) {
    const self = this
    return new PaymentAccountMinimalEntity(self,data)
  }


  // Entity access: `client.PaymentInitiation().list()` / `client.PaymentInitiation().load({ id })`.
  PaymentInitiation(data?: any) {
    const self = this
    return new PaymentInitiationEntity(self,data)
  }


  // Entity access: `client.PaymentRequest().list()` / `client.PaymentRequest().load({ id })`.
  PaymentRequest(data?: any) {
    const self = this
    return new PaymentRequestEntity(self,data)
  }


  // Entity access: `client.PaymentRequestEvent().list()` / `client.PaymentRequestEvent().load({ id })`.
  PaymentRequestEvent(data?: any) {
    const self = this
    return new PaymentRequestEventEntity(self,data)
  }


  // Entity access: `client.PaymentRequestMetric().list()` / `client.PaymentRequestMetric().load({ id })`.
  PaymentRequestMetric(data?: any) {
    const self = this
    return new PaymentRequestMetricEntity(self,data)
  }


  // Entity access: `client.PaymentRequestMinimal().list()` / `client.PaymentRequestMinimal().load({ id })`.
  PaymentRequestMinimal(data?: any) {
    const self = this
    return new PaymentRequestMinimalEntity(self,data)
  }


  // Entity access: `client.PaymentRequestResult().list()` / `client.PaymentRequestResult().load({ id })`.
  PaymentRequestResult(data?: any) {
    const self = this
    return new PaymentRequestResultEntity(self,data)
  }


  // Entity access: `client.PaymentRequestsCreate().list()` / `client.PaymentRequestsCreate().load({ id })`.
  PaymentRequestsCreate(data?: any) {
    const self = this
    return new PaymentRequestsCreateEntity(self,data)
  }


  // Entity access: `client.Payout().list()` / `client.Payout().load({ id })`.
  Payout(data?: any) {
    const self = this
    return new PayoutEntity(self,data)
  }


  // Entity access: `client.PayoutKeyset().list()` / `client.PayoutKeyset().load({ id })`.
  PayoutKeyset(data?: any) {
    const self = this
    return new PayoutKeysetEntity(self,data)
  }


  // Entity access: `client.PayoutMetric().list()` / `client.PayoutMetric().load({ id })`.
  PayoutMetric(data?: any) {
    const self = this
    return new PayoutMetricEntity(self,data)
  }


  // Entity access: `client.PayoutsCreate().list()` / `client.PayoutsCreate().load({ id })`.
  PayoutsCreate(data?: any) {
    const self = this
    return new PayoutsCreateEntity(self,data)
  }


  // Entity access: `client.Payrun().list()` / `client.Payrun().load({ id })`.
  Payrun(data?: any) {
    const self = this
    return new PayrunEntity(self,data)
  }


  // Entity access: `client.Report().list()` / `client.Report().load({ id })`.
  Report(data?: any) {
    const self = this
    return new ReportEntity(self,data)
  }


  // Entity access: `client.ReportResult().list()` / `client.ReportResult().load({ id })`.
  ReportResult(data?: any) {
    const self = this
    return new ReportResultEntity(self,data)
  }


  // Entity access: `client.RolesCreate().list()` / `client.RolesCreate().load({ id })`.
  RolesCreate(data?: any) {
    const self = this
    return new RolesCreateEntity(self,data)
  }


  // Entity access: `client.Rule().list()` / `client.Rule().load({ id })`.
  Rule(data?: any) {
    const self = this
    return new RuleEntity(self,data)
  }


  // Entity access: `client.RuleEvent().list()` / `client.RuleEvent().load({ id })`.
  RuleEvent(data?: any) {
    const self = this
    return new RuleEventEntity(self,data)
  }


  // Entity access: `client.Tag().list()` / `client.Tag().load({ id })`.
  Tag(data?: any) {
    const self = this
    return new TagEntity(self,data)
  }


  // Entity access: `client.Token().list()` / `client.Token().load({ id })`.
  Token(data?: any) {
    const self = this
    return new TokenEntity(self,data)
  }


  // Entity access: `client.Transaction().list()` / `client.Transaction().load({ id })`.
  Transaction(data?: any) {
    const self = this
    return new TransactionEntity(self,data)
  }


  // Entity access: `client.User().list()` / `client.User().load({ id })`.
  User(data?: any) {
    const self = this
    return new UserEntity(self,data)
  }


  // Entity access: `client.UserInvite().list()` / `client.UserInvite().load({ id })`.
  UserInvite(data?: any) {
    const self = this
    return new UserInviteEntity(self,data)
  }


  // Entity access: `client.UserInvitesCreate().list()` / `client.UserInvitesCreate().load({ id })`.
  UserInvitesCreate(data?: any) {
    const self = this
    return new UserInvitesCreateEntity(self,data)
  }


  // Entity access: `client.Virtual().list()` / `client.Virtual().load({ id })`.
  Virtual(data?: any) {
    const self = this
    return new VirtualEntity(self,data)
  }


  // Entity access: `client.Webhook().list()` / `client.Webhook().load({ id })`.
  Webhook(data?: any) {
    const self = this
    return new WebhookEntity(self,data)
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

  BaseFeature,
  NofrixionEntityBase,

  NofrixionSDK,
  SDK,
}


