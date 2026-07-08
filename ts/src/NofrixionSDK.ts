// Nofrixion Ts SDK

import { AccountEntity } from './entity/AccountEntity'
import { BeneficiaryEntity } from './entity/BeneficiaryEntity'
import { CancelEntity } from './entity/CancelEntity'
import { DisableEntity } from './entity/DisableEntity'
import { EnableEntity } from './entity/EnableEntity'
import { MerchantEntity } from './entity/MerchantEntity'
import { MetadataEntity } from './entity/MetadataEntity'
import { NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity } from './entity/NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity'
import { NoFrixionBizBizModelsPagingPaymentRequestPageEntity } from './entity/NoFrixionBizBizModelsPagingPaymentRequestPageEntity'
import { NoFrixionBizBizModelsPagingPayoutPageEntity } from './entity/NoFrixionBizBizModelsPagingPayoutPageEntity'
import { NoFrixionBizBizModelsPagingPayrunPageEntity } from './entity/NoFrixionBizBizModelsPagingPayrunPageEntity'
import { NoFrixionBizBizModelsPagingRuleEventsPageEntity } from './entity/NoFrixionBizBizModelsPagingRuleEventsPageEntity'
import { NoFrixionBizBizModelsPagingRulesPageEntity } from './entity/NoFrixionBizBizModelsPagingRulesPageEntity'
import { NoFrixionBizBizModelsPaymentsCardPaymentEntity } from './entity/NoFrixionBizBizModelsPaymentsCardPaymentEntity'
import { NoFrixionBizBizModelsPaymentsCardPublicKeyEntity } from './entity/NoFrixionBizBizModelsPaymentsCardPublicKeyEntity'
import { NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity } from './entity/NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity'
import { NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity } from './entity/NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity'
import { NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity } from './entity/NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity'
import { NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity } from './entity/NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity'
import { NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity } from './entity/NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity'
import { NoFrixionMoneyMoovModelsBatchPayoutEntity } from './entity/NoFrixionMoneyMoovModelsBatchPayoutEntity'
import { NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity } from './entity/NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity'
import { NoFrixionMoneyMoovModelsBeneficiaryPageEntity } from './entity/NoFrixionMoneyMoovModelsBeneficiaryPageEntity'
import { NoFrixionMoneyMoovModelsCardCustomerTokenEntity } from './entity/NoFrixionMoneyMoovModelsCardCustomerTokenEntity'
import { NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity } from './entity/NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity'
import { NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity } from './entity/NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity'
import { NoFrixionMoneyMoovModelsFxRateEntity } from './entity/NoFrixionMoneyMoovModelsFxRateEntity'
import { NoFrixionMoneyMoovModelsIPaymentEntity } from './entity/NoFrixionMoneyMoovModelsIPaymentEntity'
import { NoFrixionMoneyMoovModelsMandatesMandateEntity } from './entity/NoFrixionMoneyMoovModelsMandatesMandateEntity'
import { NoFrixionMoneyMoovModelsMerchantEntity } from './entity/NoFrixionMoneyMoovModelsMerchantEntity'
import { NoFrixionMoneyMoovModelsMerchantPageEntity } from './entity/NoFrixionMoneyMoovModelsMerchantPageEntity'
import { NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity } from './entity/NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity'
import { NoFrixionMoneyMoovModelsMerchantTokenEntity } from './entity/NoFrixionMoneyMoovModelsMerchantTokenEntity'
import { NoFrixionMoneyMoovModelsMerchantTokenPageEntity } from './entity/NoFrixionMoneyMoovModelsMerchantTokenPageEntity'
import { NoFrixionMoneyMoovModelsNoFrixionVersionEntity } from './entity/NoFrixionMoneyMoovModelsNoFrixionVersionEntity'
import { NoFrixionMoneyMoovModelsOpenBankingAccountEntity } from './entity/NoFrixionMoneyMoovModelsOpenBankingAccountEntity'
import { NoFrixionMoneyMoovModelsOpenBankingConsentEntity } from './entity/NoFrixionMoneyMoovModelsOpenBankingConsentEntity'
import { NoFrixionMoneyMoovModelsOpenBankingTransactionEntity } from './entity/NoFrixionMoneyMoovModelsOpenBankingTransactionEntity'
import { NoFrixionMoneyMoovModelsPaymentEntity } from './entity/NoFrixionMoneyMoovModelsPaymentEntity'
import { NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity } from './entity/NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity'
import { NoFrixionMoneyMoovModelsPaymentAccountPageEntity } from './entity/NoFrixionMoneyMoovModelsPaymentAccountPageEntity'
import { NoFrixionMoneyMoovModelsPaymentInitiationEntity } from './entity/NoFrixionMoneyMoovModelsPaymentInitiationEntity'
import { NoFrixionMoneyMoovModelsPaymentRequestEventEntity } from './entity/NoFrixionMoneyMoovModelsPaymentRequestEventEntity'
import { NoFrixionMoneyMoovModelsPaymentRequestMetricEntity } from './entity/NoFrixionMoneyMoovModelsPaymentRequestMetricEntity'
import { NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity } from './entity/NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity'
import { NoFrixionMoneyMoovModelsPaymentRequestResultEntity } from './entity/NoFrixionMoneyMoovModelsPaymentRequestResultEntity'
import { NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity } from './entity/NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity'
import { NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity } from './entity/NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity'
import { NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity } from './entity/NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity'
import { NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity } from './entity/NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity'
import { NoFrixionMoneyMoovModelsPayoutKeysetPageEntity } from './entity/NoFrixionMoneyMoovModelsPayoutKeysetPageEntity'
import { NoFrixionMoneyMoovModelsPayoutMetricEntity } from './entity/NoFrixionMoneyMoovModelsPayoutMetricEntity'
import { NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity } from './entity/NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity'
import { NoFrixionMoneyMoovModelsPayrunEntity } from './entity/NoFrixionMoneyMoovModelsPayrunEntity'
import { NoFrixionMoneyMoovModelsReportResultEntity } from './entity/NoFrixionMoneyMoovModelsReportResultEntity'
import { NoFrixionMoneyMoovModelsRuleEntity } from './entity/NoFrixionMoneyMoovModelsRuleEntity'
import { NoFrixionMoneyMoovModelsTransactionEntity } from './entity/NoFrixionMoneyMoovModelsTransactionEntity'
import { NoFrixionMoneyMoovModelsTransactionPageEntity } from './entity/NoFrixionMoneyMoovModelsTransactionPageEntity'
import { NoFrixionMoneyMoovModelsUserInviteEntity } from './entity/NoFrixionMoneyMoovModelsUserInviteEntity'
import { NoFrixionMoneyMoovModelsUserInvitePageEntity } from './entity/NoFrixionMoneyMoovModelsUserInvitePageEntity'
import { NoFrixionMoneyMoovModelsUserPageEntity } from './entity/NoFrixionMoneyMoovModelsUserPageEntity'
import { NoFrixionMoneyMoovModelsWebhookEntity } from './entity/NoFrixionMoneyMoovModelsWebhookEntity'
import { OpenBankingEntity } from './entity/OpenBankingEntity'
import { PayeeverificationEntity } from './entity/PayeeverificationEntity'
import { PaymentRequestEntity } from './entity/PaymentRequestEntity'
import { PayoutEntity } from './entity/PayoutEntity'
import { PayrunEntity } from './entity/PayrunEntity'
import { RejectEntity } from './entity/RejectEntity'
import { ReportEntity } from './entity/ReportEntity'
import { RuleEntity } from './entity/RuleEntity'
import { SendEntity } from './entity/SendEntity'
import { SendbeneficiaryEntity } from './entity/SendbeneficiaryEntity'
import { TagEntity } from './entity/TagEntity'
import { TokenEntity } from './entity/TokenEntity'
import { TransactionEntity } from './entity/TransactionEntity'
import { UserEntity } from './entity/UserEntity'
import { UserInviteEntity } from './entity/UserInviteEntity'
import { VirtualEntity } from './entity/VirtualEntity'
import { WebhookEntity } from './entity/WebhookEntity'
import { WhoamiEntity } from './entity/WhoamiEntity'
import { WhoamitrustedappEntity } from './entity/WhoamitrustedappEntity'

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


  // Entity access: `client.Beneficiary().list()` / `client.Beneficiary().load({ id })`.
  Beneficiary(data?: any) {
    const self = this
    return new BeneficiaryEntity(self,data)
  }


  // Entity access: `client.Cancel().list()` / `client.Cancel().load({ id })`.
  Cancel(data?: any) {
    const self = this
    return new CancelEntity(self,data)
  }


  // Entity access: `client.Disable().list()` / `client.Disable().load({ id })`.
  Disable(data?: any) {
    const self = this
    return new DisableEntity(self,data)
  }


  // Entity access: `client.Enable().list()` / `client.Enable().load({ id })`.
  Enable(data?: any) {
    const self = this
    return new EnableEntity(self,data)
  }


  // Entity access: `client.Merchant().list()` / `client.Merchant().load({ id })`.
  Merchant(data?: any) {
    const self = this
    return new MerchantEntity(self,data)
  }


  // Entity access: `client.Metadata().list()` / `client.Metadata().load({ id })`.
  Metadata(data?: any) {
    const self = this
    return new MetadataEntity(self,data)
  }


  // Entity access: `client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage().list()` / `client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage().load({ id })`.
  NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(data?: any) {
    const self = this
    return new NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity(self,data)
  }


  // Entity access: `client.NoFrixionBizBizModelsPagingPaymentRequestPage().list()` / `client.NoFrixionBizBizModelsPagingPaymentRequestPage().load({ id })`.
  NoFrixionBizBizModelsPagingPaymentRequestPage(data?: any) {
    const self = this
    return new NoFrixionBizBizModelsPagingPaymentRequestPageEntity(self,data)
  }


  // Entity access: `client.NoFrixionBizBizModelsPagingPayoutPage().list()` / `client.NoFrixionBizBizModelsPagingPayoutPage().load({ id })`.
  NoFrixionBizBizModelsPagingPayoutPage(data?: any) {
    const self = this
    return new NoFrixionBizBizModelsPagingPayoutPageEntity(self,data)
  }


  // Entity access: `client.NoFrixionBizBizModelsPagingPayrunPage().list()` / `client.NoFrixionBizBizModelsPagingPayrunPage().load({ id })`.
  NoFrixionBizBizModelsPagingPayrunPage(data?: any) {
    const self = this
    return new NoFrixionBizBizModelsPagingPayrunPageEntity(self,data)
  }


  // Entity access: `client.NoFrixionBizBizModelsPagingRuleEventsPage().list()` / `client.NoFrixionBizBizModelsPagingRuleEventsPage().load({ id })`.
  NoFrixionBizBizModelsPagingRuleEventsPage(data?: any) {
    const self = this
    return new NoFrixionBizBizModelsPagingRuleEventsPageEntity(self,data)
  }


  // Entity access: `client.NoFrixionBizBizModelsPagingRulesPage().list()` / `client.NoFrixionBizBizModelsPagingRulesPage().load({ id })`.
  NoFrixionBizBizModelsPagingRulesPage(data?: any) {
    const self = this
    return new NoFrixionBizBizModelsPagingRulesPageEntity(self,data)
  }


  // Entity access: `client.NoFrixionBizBizModelsPaymentsCardPayment().list()` / `client.NoFrixionBizBizModelsPaymentsCardPayment().load({ id })`.
  NoFrixionBizBizModelsPaymentsCardPayment(data?: any) {
    const self = this
    return new NoFrixionBizBizModelsPaymentsCardPaymentEntity(self,data)
  }


  // Entity access: `client.NoFrixionBizBizModelsPaymentsCardPublicKey().list()` / `client.NoFrixionBizBizModelsPaymentsCardPublicKey().load({ id })`.
  NoFrixionBizBizModelsPaymentsCardPublicKey(data?: any) {
    const self = this
    return new NoFrixionBizBizModelsPaymentsCardPublicKeyEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries().list()` / `client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries().load({ id })`.
  NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment().list()` / `client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment().load({ id })`.
  NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate().list()` / `client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate().load({ id })`.
  NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate().list()` / `client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate().load({ id })`.
  NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant().list()` / `client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant().load({ id })`.
  NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsBatchPayout().list()` / `client.NoFrixionMoneyMoovModelsBatchPayout().load({ id })`.
  NoFrixionMoneyMoovModelsBatchPayout(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsBatchPayoutEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage().list()` / `client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage().load({ id })`.
  NoFrixionMoneyMoovModelsBeneficiaryGroupPage(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsBeneficiaryPage().list()` / `client.NoFrixionMoneyMoovModelsBeneficiaryPage().load({ id })`.
  NoFrixionMoneyMoovModelsBeneficiaryPage(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsBeneficiaryPageEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsCardCustomerToken().list()` / `client.NoFrixionMoneyMoovModelsCardCustomerToken().load({ id })`.
  NoFrixionMoneyMoovModelsCardCustomerToken(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsCardCustomerTokenEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo().list()` / `client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo().load({ id })`.
  NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit().list()` / `client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit().load({ id })`.
  NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsFxRate().list()` / `client.NoFrixionMoneyMoovModelsFxRate().load({ id })`.
  NoFrixionMoneyMoovModelsFxRate(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsFxRateEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsIPayment().list()` / `client.NoFrixionMoneyMoovModelsIPayment().load({ id })`.
  NoFrixionMoneyMoovModelsIPayment(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsIPaymentEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsMandatesMandate().list()` / `client.NoFrixionMoneyMoovModelsMandatesMandate().load({ id })`.
  NoFrixionMoneyMoovModelsMandatesMandate(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsMandatesMandateEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsMerchant().list()` / `client.NoFrixionMoneyMoovModelsMerchant().load({ id })`.
  NoFrixionMoneyMoovModelsMerchant(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsMerchantEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsMerchantPage().list()` / `client.NoFrixionMoneyMoovModelsMerchantPage().load({ id })`.
  NoFrixionMoneyMoovModelsMerchantPage(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsMerchantPageEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting().list()` / `client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting().load({ id })`.
  NoFrixionMoneyMoovModelsMerchantPayByBankSetting(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsMerchantToken().list()` / `client.NoFrixionMoneyMoovModelsMerchantToken().load({ id })`.
  NoFrixionMoneyMoovModelsMerchantToken(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsMerchantTokenEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsMerchantTokenPage().list()` / `client.NoFrixionMoneyMoovModelsMerchantTokenPage().load({ id })`.
  NoFrixionMoneyMoovModelsMerchantTokenPage(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsMerchantTokenPageEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsNoFrixionVersion().list()` / `client.NoFrixionMoneyMoovModelsNoFrixionVersion().load({ id })`.
  NoFrixionMoneyMoovModelsNoFrixionVersion(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsNoFrixionVersionEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsOpenBankingAccount().list()` / `client.NoFrixionMoneyMoovModelsOpenBankingAccount().load({ id })`.
  NoFrixionMoneyMoovModelsOpenBankingAccount(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsOpenBankingAccountEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsOpenBankingConsent().list()` / `client.NoFrixionMoneyMoovModelsOpenBankingConsent().load({ id })`.
  NoFrixionMoneyMoovModelsOpenBankingConsent(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsOpenBankingConsentEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsOpenBankingTransaction().list()` / `client.NoFrixionMoneyMoovModelsOpenBankingTransaction().load({ id })`.
  NoFrixionMoneyMoovModelsOpenBankingTransaction(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsOpenBankingTransactionEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPayment().list()` / `client.NoFrixionMoneyMoovModelsPayment().load({ id })`.
  NoFrixionMoneyMoovModelsPayment(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPaymentEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage().list()` / `client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage().load({ id })`.
  NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPaymentAccountPage().list()` / `client.NoFrixionMoneyMoovModelsPaymentAccountPage().load({ id })`.
  NoFrixionMoneyMoovModelsPaymentAccountPage(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPaymentAccountPageEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPaymentInitiation().list()` / `client.NoFrixionMoneyMoovModelsPaymentInitiation().load({ id })`.
  NoFrixionMoneyMoovModelsPaymentInitiation(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPaymentInitiationEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPaymentRequestEvent().list()` / `client.NoFrixionMoneyMoovModelsPaymentRequestEvent().load({ id })`.
  NoFrixionMoneyMoovModelsPaymentRequestEvent(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPaymentRequestEventEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPaymentRequestMetric().list()` / `client.NoFrixionMoneyMoovModelsPaymentRequestMetric().load({ id })`.
  NoFrixionMoneyMoovModelsPaymentRequestMetric(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPaymentRequestMetricEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPaymentRequestMinimal().list()` / `client.NoFrixionMoneyMoovModelsPaymentRequestMinimal().load({ id })`.
  NoFrixionMoneyMoovModelsPaymentRequestMinimal(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPaymentRequestResult().list()` / `client.NoFrixionMoneyMoovModelsPaymentRequestResult().load({ id })`.
  NoFrixionMoneyMoovModelsPaymentRequestResult(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPaymentRequestResultEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment().list()` / `client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment().load({ id })`.
  NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2().list()` / `client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2().load({ id })`.
  NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3().list()` / `client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3().load({ id })`.
  NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4().list()` / `client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4().load({ id })`.
  NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPayoutKeysetPage().list()` / `client.NoFrixionMoneyMoovModelsPayoutKeysetPage().load({ id })`.
  NoFrixionMoneyMoovModelsPayoutKeysetPage(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPayoutKeysetPageEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPayoutMetric().list()` / `client.NoFrixionMoneyMoovModelsPayoutMetric().load({ id })`.
  NoFrixionMoneyMoovModelsPayoutMetric(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPayoutMetricEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate().list()` / `client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate().load({ id })`.
  NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsPayrun().list()` / `client.NoFrixionMoneyMoovModelsPayrun().load({ id })`.
  NoFrixionMoneyMoovModelsPayrun(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsPayrunEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsReportResult().list()` / `client.NoFrixionMoneyMoovModelsReportResult().load({ id })`.
  NoFrixionMoneyMoovModelsReportResult(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsReportResultEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsRule().list()` / `client.NoFrixionMoneyMoovModelsRule().load({ id })`.
  NoFrixionMoneyMoovModelsRule(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsRuleEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsTransaction().list()` / `client.NoFrixionMoneyMoovModelsTransaction().load({ id })`.
  NoFrixionMoneyMoovModelsTransaction(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsTransactionEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsTransactionPage().list()` / `client.NoFrixionMoneyMoovModelsTransactionPage().load({ id })`.
  NoFrixionMoneyMoovModelsTransactionPage(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsTransactionPageEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsUserInvite().list()` / `client.NoFrixionMoneyMoovModelsUserInvite().load({ id })`.
  NoFrixionMoneyMoovModelsUserInvite(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsUserInviteEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsUserInvitePage().list()` / `client.NoFrixionMoneyMoovModelsUserInvitePage().load({ id })`.
  NoFrixionMoneyMoovModelsUserInvitePage(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsUserInvitePageEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsUserPage().list()` / `client.NoFrixionMoneyMoovModelsUserPage().load({ id })`.
  NoFrixionMoneyMoovModelsUserPage(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsUserPageEntity(self,data)
  }


  // Entity access: `client.NoFrixionMoneyMoovModelsWebhook().list()` / `client.NoFrixionMoneyMoovModelsWebhook().load({ id })`.
  NoFrixionMoneyMoovModelsWebhook(data?: any) {
    const self = this
    return new NoFrixionMoneyMoovModelsWebhookEntity(self,data)
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


  // Entity access: `client.PaymentRequest().list()` / `client.PaymentRequest().load({ id })`.
  PaymentRequest(data?: any) {
    const self = this
    return new PaymentRequestEntity(self,data)
  }


  // Entity access: `client.Payout().list()` / `client.Payout().load({ id })`.
  Payout(data?: any) {
    const self = this
    return new PayoutEntity(self,data)
  }


  // Entity access: `client.Payrun().list()` / `client.Payrun().load({ id })`.
  Payrun(data?: any) {
    const self = this
    return new PayrunEntity(self,data)
  }


  // Entity access: `client.Reject().list()` / `client.Reject().load({ id })`.
  Reject(data?: any) {
    const self = this
    return new RejectEntity(self,data)
  }


  // Entity access: `client.Report().list()` / `client.Report().load({ id })`.
  Report(data?: any) {
    const self = this
    return new ReportEntity(self,data)
  }


  // Entity access: `client.Rule().list()` / `client.Rule().load({ id })`.
  Rule(data?: any) {
    const self = this
    return new RuleEntity(self,data)
  }


  // Entity access: `client.Send().list()` / `client.Send().load({ id })`.
  Send(data?: any) {
    const self = this
    return new SendEntity(self,data)
  }


  // Entity access: `client.Sendbeneficiary().list()` / `client.Sendbeneficiary().load({ id })`.
  Sendbeneficiary(data?: any) {
    const self = this
    return new SendbeneficiaryEntity(self,data)
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


  // Entity access: `client.Whoami().list()` / `client.Whoami().load({ id })`.
  Whoami(data?: any) {
    const self = this
    return new WhoamiEntity(self,data)
  }


  // Entity access: `client.Whoamitrustedapp().list()` / `client.Whoamitrustedapp().load({ id })`.
  Whoamitrustedapp(data?: any) {
    const self = this
    return new WhoamitrustedappEntity(self,data)
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


