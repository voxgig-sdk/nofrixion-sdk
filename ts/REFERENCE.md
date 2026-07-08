# Nofrixion TypeScript SDK Reference

Complete API reference for the Nofrixion TypeScript SDK.


## NofrixionSDK

### Constructor

```ts
new NofrixionSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NofrixionSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = NofrixionSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `NofrixionSDK` instance in test mode.


### Instance Methods

#### `Account(data?: object)`

Create a new `Account` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AccountEntity` instance.

#### `Beneficiary(data?: object)`

Create a new `Beneficiary` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `BeneficiaryEntity` instance.

#### `Cancel(data?: object)`

Create a new `Cancel` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CancelEntity` instance.

#### `Disable(data?: object)`

Create a new `Disable` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DisableEntity` instance.

#### `Enable(data?: object)`

Create a new `Enable` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `EnableEntity` instance.

#### `Merchant(data?: object)`

Create a new `Merchant` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MerchantEntity` instance.

#### `Metadata(data?: object)`

Create a new `Metadata` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MetadataEntity` instance.

#### `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(data?: object)`

Create a new `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity` instance.

#### `NoFrixionBizBizModelsPagingPaymentRequestPage(data?: object)`

Create a new `NoFrixionBizBizModelsPagingPaymentRequestPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionBizBizModelsPagingPaymentRequestPageEntity` instance.

#### `NoFrixionBizBizModelsPagingPayoutPage(data?: object)`

Create a new `NoFrixionBizBizModelsPagingPayoutPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionBizBizModelsPagingPayoutPageEntity` instance.

#### `NoFrixionBizBizModelsPagingPayrunPage(data?: object)`

Create a new `NoFrixionBizBizModelsPagingPayrunPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionBizBizModelsPagingPayrunPageEntity` instance.

#### `NoFrixionBizBizModelsPagingRuleEventsPage(data?: object)`

Create a new `NoFrixionBizBizModelsPagingRuleEventsPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionBizBizModelsPagingRuleEventsPageEntity` instance.

#### `NoFrixionBizBizModelsPagingRulesPage(data?: object)`

Create a new `NoFrixionBizBizModelsPagingRulesPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionBizBizModelsPagingRulesPageEntity` instance.

#### `NoFrixionBizBizModelsPaymentsCardPayment(data?: object)`

Create a new `NoFrixionBizBizModelsPaymentsCardPayment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionBizBizModelsPaymentsCardPaymentEntity` instance.

#### `NoFrixionBizBizModelsPaymentsCardPublicKey(data?: object)`

Create a new `NoFrixionBizBizModelsPaymentsCardPublicKey` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionBizBizModelsPaymentsCardPublicKeyEntity` instance.

#### `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(data?: object)`

Create a new `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity` instance.

#### `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(data?: object)`

Create a new `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity` instance.

#### `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(data?: object)`

Create a new `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity` instance.

#### `NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(data?: object)`

Create a new `NoFrixionMoneyMoovApiFeaturesUserInvitesCreate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity` instance.

#### `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(data?: object)`

Create a new `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity` instance.

#### `NoFrixionMoneyMoovModelsBatchPayout(data?: object)`

Create a new `NoFrixionMoneyMoovModelsBatchPayout` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsBatchPayoutEntity` instance.

#### `NoFrixionMoneyMoovModelsBeneficiaryGroupPage(data?: object)`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryGroupPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity` instance.

#### `NoFrixionMoneyMoovModelsBeneficiaryPage(data?: object)`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsBeneficiaryPageEntity` instance.

#### `NoFrixionMoneyMoovModelsCardCustomerToken(data?: object)`

Create a new `NoFrixionMoneyMoovModelsCardCustomerToken` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsCardCustomerTokenEntity` instance.

#### `NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(data?: object)`

Create a new `NoFrixionMoneyMoovModelsCurrencyCurrencyInfo` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity` instance.

#### `NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(data?: object)`

Create a new `NoFrixionMoneyMoovModelsDirectDebitBatchSubmit` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity` instance.

#### `NoFrixionMoneyMoovModelsFxRate(data?: object)`

Create a new `NoFrixionMoneyMoovModelsFxRate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsFxRateEntity` instance.

#### `NoFrixionMoneyMoovModelsIPayment(data?: object)`

Create a new `NoFrixionMoneyMoovModelsIPayment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsIPaymentEntity` instance.

#### `NoFrixionMoneyMoovModelsMandatesMandate(data?: object)`

Create a new `NoFrixionMoneyMoovModelsMandatesMandate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsMandatesMandateEntity` instance.

#### `NoFrixionMoneyMoovModelsMerchant(data?: object)`

Create a new `NoFrixionMoneyMoovModelsMerchant` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsMerchantEntity` instance.

#### `NoFrixionMoneyMoovModelsMerchantPage(data?: object)`

Create a new `NoFrixionMoneyMoovModelsMerchantPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsMerchantPageEntity` instance.

#### `NoFrixionMoneyMoovModelsMerchantPayByBankSetting(data?: object)`

Create a new `NoFrixionMoneyMoovModelsMerchantPayByBankSetting` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity` instance.

#### `NoFrixionMoneyMoovModelsMerchantToken(data?: object)`

Create a new `NoFrixionMoneyMoovModelsMerchantToken` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsMerchantTokenEntity` instance.

#### `NoFrixionMoneyMoovModelsMerchantTokenPage(data?: object)`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsMerchantTokenPageEntity` instance.

#### `NoFrixionMoneyMoovModelsNoFrixionVersion(data?: object)`

Create a new `NoFrixionMoneyMoovModelsNoFrixionVersion` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsNoFrixionVersionEntity` instance.

#### `NoFrixionMoneyMoovModelsOpenBankingAccount(data?: object)`

Create a new `NoFrixionMoneyMoovModelsOpenBankingAccount` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsOpenBankingAccountEntity` instance.

#### `NoFrixionMoneyMoovModelsOpenBankingConsent(data?: object)`

Create a new `NoFrixionMoneyMoovModelsOpenBankingConsent` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsOpenBankingConsentEntity` instance.

#### `NoFrixionMoneyMoovModelsOpenBankingTransaction(data?: object)`

Create a new `NoFrixionMoneyMoovModelsOpenBankingTransaction` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsOpenBankingTransactionEntity` instance.

#### `NoFrixionMoneyMoovModelsPayment(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPayment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPaymentEntity` instance.

#### `NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountMinimalPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity` instance.

#### `NoFrixionMoneyMoovModelsPaymentAccountPage(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPaymentAccountPageEntity` instance.

#### `NoFrixionMoneyMoovModelsPaymentInitiation(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPaymentInitiation` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPaymentInitiationEntity` instance.

#### `NoFrixionMoneyMoovModelsPaymentRequestEvent(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestEvent` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPaymentRequestEventEntity` instance.

#### `NoFrixionMoneyMoovModelsPaymentRequestMetric(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMetric` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPaymentRequestMetricEntity` instance.

#### `NoFrixionMoneyMoovModelsPaymentRequestMinimal(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMinimal` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity` instance.

#### `NoFrixionMoneyMoovModelsPaymentRequestResult(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestResult` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPaymentRequestResultEntity` instance.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity` instance.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity` instance.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity` instance.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity` instance.

#### `NoFrixionMoneyMoovModelsPayoutKeysetPage(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPayoutKeysetPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPayoutKeysetPageEntity` instance.

#### `NoFrixionMoneyMoovModelsPayoutMetric(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPayoutMetric` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPayoutMetricEntity` instance.

#### `NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPayoutsPayoutsCreate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity` instance.

#### `NoFrixionMoneyMoovModelsPayrun(data?: object)`

Create a new `NoFrixionMoneyMoovModelsPayrun` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsPayrunEntity` instance.

#### `NoFrixionMoneyMoovModelsReportResult(data?: object)`

Create a new `NoFrixionMoneyMoovModelsReportResult` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsReportResultEntity` instance.

#### `NoFrixionMoneyMoovModelsRule(data?: object)`

Create a new `NoFrixionMoneyMoovModelsRule` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsRuleEntity` instance.

#### `NoFrixionMoneyMoovModelsTransaction(data?: object)`

Create a new `NoFrixionMoneyMoovModelsTransaction` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsTransactionEntity` instance.

#### `NoFrixionMoneyMoovModelsTransactionPage(data?: object)`

Create a new `NoFrixionMoneyMoovModelsTransactionPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsTransactionPageEntity` instance.

#### `NoFrixionMoneyMoovModelsUserInvite(data?: object)`

Create a new `NoFrixionMoneyMoovModelsUserInvite` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsUserInviteEntity` instance.

#### `NoFrixionMoneyMoovModelsUserInvitePage(data?: object)`

Create a new `NoFrixionMoneyMoovModelsUserInvitePage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsUserInvitePageEntity` instance.

#### `NoFrixionMoneyMoovModelsUserPage(data?: object)`

Create a new `NoFrixionMoneyMoovModelsUserPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsUserPageEntity` instance.

#### `NoFrixionMoneyMoovModelsWebhook(data?: object)`

Create a new `NoFrixionMoneyMoovModelsWebhook` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionMoneyMoovModelsWebhookEntity` instance.

#### `OpenBanking(data?: object)`

Create a new `OpenBanking` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `OpenBankingEntity` instance.

#### `Payeeverification(data?: object)`

Create a new `Payeeverification` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PayeeverificationEntity` instance.

#### `PaymentRequest(data?: object)`

Create a new `PaymentRequest` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PaymentRequestEntity` instance.

#### `Payout(data?: object)`

Create a new `Payout` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PayoutEntity` instance.

#### `Payrun(data?: object)`

Create a new `Payrun` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PayrunEntity` instance.

#### `Reject(data?: object)`

Create a new `Reject` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `RejectEntity` instance.

#### `Report(data?: object)`

Create a new `Report` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ReportEntity` instance.

#### `Rule(data?: object)`

Create a new `Rule` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `RuleEntity` instance.

#### `Send(data?: object)`

Create a new `Send` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SendEntity` instance.

#### `Sendbeneficiary(data?: object)`

Create a new `Sendbeneficiary` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SendbeneficiaryEntity` instance.

#### `Tag(data?: object)`

Create a new `Tag` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TagEntity` instance.

#### `Token(data?: object)`

Create a new `Token` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TokenEntity` instance.

#### `Transaction(data?: object)`

Create a new `Transaction` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TransactionEntity` instance.

#### `User(data?: object)`

Create a new `User` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `UserEntity` instance.

#### `UserInvite(data?: object)`

Create a new `UserInvite` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `UserInviteEntity` instance.

#### `Virtual(data?: object)`

Create a new `Virtual` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `VirtualEntity` instance.

#### `Webhook(data?: object)`

Create a new `Webhook` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `WebhookEntity` instance.

#### `Whoami(data?: object)`

Create a new `Whoami` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `WhoamiEntity` instance.

#### `Whoamitrustedapp(data?: object)`

Create a new `Whoamitrustedapp` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `WhoamitrustedappEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `NofrixionSDK.test()`.

**Returns:** `NofrixionSDK` instance in test mode.


---

## AccountEntity

```ts
const account = client.Account()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `account_name` | `string` | No |  |
| `account_supplier_name` | `string` | No |  |
| `account_type` | `string` | No |  |
| `available_balance` | `number` | No |  |
| `available_balance_minor_unit` | `number` | No |  |
| `balance` | `number` | No |  |
| `balance_minor_unit` | `number` | No |  |
| `bank_name` | `string` | No |  |
| `consent_id` | `string` | No |  |
| `created_by` | `Record<string, any>` | Yes |  |
| `created_by_display_name` | `string` | No |  |
| `currency` | `string` | No |  |
| `default_payment_rail` | `string` | No |  |
| `display_name` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `external_account_icon` | `string` | No |  |
| `format` | `string` | No |  |
| `from_date` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `Record<string, any>` | Yes |  |
| `inserted` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_connected_account` | `boolean` | No |  |
| `is_default` | `boolean` | No |  |
| `is_trust_account` | `boolean` | No |  |
| `is_virtual` | `boolean` | No |  |
| `last_transaction` | `Record<string, any>` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `physical_account_id` | `string` | No |  |
| `role_i_d` | `any[]` | No |  |
| `rule` | `any[]` | No |  |
| `submitted_payouts_balance` | `number` | No |  |
| `submitted_payouts_balance_minor_unit` | `number` | No |  |
| `summary` | `string` | No |  |
| `supplier_physical_account_id` | `string` | No |  |
| `supplier_sepa_instant_status` | `string` | No |  |
| `to_date` | `string` | No |  |
| `xero_bank_feed_connection_status` | `string` | No |  |
| `xero_bank_feed_last_synced_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` | No |  |
| `xero_bank_feed_sync_status` | `string` | No |  |
| `xero_unsynchronised_transactions_count` | `number` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Account().create({
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Account().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Account().load({ id: 'account_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Account().remove({ id: 'account_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Account().update({
  id: 'account_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AccountEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## BeneficiaryEntity

```ts
const beneficiary = client.Beneficiary()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `beneficiary_event` | `any[]` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `created_by` | `Record<string, any>` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `Record<string, any>` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `boolean` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `any[]` | No |  |
| `source_account_i_d` | `any[]` | No |  |
| `their_reference` | `string` | No |  |

### Field Usage by Operation

| Field | load | create | update | remove |
| --- | --- | --- | --- | --- |
| `approval_callback_url` | - | - | - | - |
| `authentication_method` | - | - | - | - |
| `authorisation` | - | - | - | - |
| `authorisers_completed_count` | - | - | - | - |
| `authorisers_required_count` | - | - | - | - |
| `beneficiary_event` | - | - | - | - |
| `can_authorise` | - | - | - | - |
| `can_update` | - | - | - | - |
| `created_by` | - | - | - | - |
| `created_by_email_address` | - | - | - | - |
| `currency` | - | - | Yes | - |
| `destination` | - | Yes | - | - |
| `has_current_user_authorised` | - | - | - | - |
| `id` | - | - | - | - |
| `inserted` | - | - | - | - |
| `is_enabled` | - | - | - | - |
| `last_authorised` | - | - | - | - |
| `last_updated` | - | - | - | - |
| `merchant_id` | - | Yes | - | - |
| `name` | - | - | Yes | - |
| `nonce` | - | - | - | - |
| `source_account` | - | - | - | - |
| `source_account_i_d` | - | - | - | - |
| `their_reference` | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Beneficiary().create({
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Beneficiary().load({ id: 'beneficiary_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Beneficiary().remove({ id: 'beneficiary_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Beneficiary().update({
  id: 'beneficiary_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `BeneficiaryEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CancelEntity

```ts
const cancel = client.Cancel()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `Record<string, any>` | Yes |  |
| `can_authorise` | `boolean` | No |  |
| `can_process` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `Record<string, any>` | No |  |
| `document` | `any[]` | No |  |
| `event` | `any[]` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `number` | No |  |
| `fx_destination_amount_minor_unit` | `number` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `number` | No |  |
| `fx_use_destination_amount` | `boolean` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_failed` | `boolean` | No |  |
| `is_settled` | `boolean` | No |  |
| `is_submitted` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `Record<string, any>` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `Record<string, any>` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `number` | No |  |
| `transacted_fx_amount` | `number` | No |  |
| `transacted_fx_rate` | `number` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Cancel().update({
  id: 'id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CancelEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DisableEntity

```ts
const disable = client.Disable()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `beneficiary_event` | `any[]` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `created_by` | `Record<string, any>` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `Record<string, any>` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `boolean` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `any[]` | No |  |
| `their_reference` | `string` | No |  |

### Operations

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Disable().update({
  id: 'id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DisableEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## EnableEntity

```ts
const enable = client.Enable()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `beneficiary_event` | `any[]` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `created_by` | `Record<string, any>` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `Record<string, any>` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `boolean` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `any[]` | No |  |
| `their_reference` | `string` | No |  |

### Operations

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Enable().update({
  id: 'id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `EnableEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MerchantEntity

```ts
const merchant = client.Merchant()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `reason` | `string` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Merchant().load({ merchant_id: 'merchant_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Merchant().remove()
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Merchant().update({
  id: 'id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MerchantEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MetadataEntity

```ts
const metadata = client.Metadata()
```

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Metadata().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MetadataEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity

```ts
const no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page = client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approved_at` | `string` | No |  |
| `currency` | `string` | No |  |
| `customer_account_number` | `string` | No |  |
| `customer_city` | `string` | No |  |
| `customer_country_code` | `string` | No |  |
| `customer_country_name` | `string` | No |  |
| `customer_email_address` | `string` | No |  |
| `customer_first_name` | `string` | No |  |
| `customer_iban` | `string` | No |  |
| `customer_last_name` | `string` | No |  |
| `customer_sort_code` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_recurring` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `reference` | `string` | No |  |
| `status` | `string` | No |  |
| `supplier_bank_account_id` | `string` | No |  |
| `supplier_customer_id` | `string` | No |  |
| `supplier_mandate_id` | `string` | No |  |
| `supplier_name` | `string` | No |  |
| `supplier_status` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionBizBizModelsPagingPaymentRequestPageEntity

```ts
const no_frixion_biz_biz_models_paging_payment_request_page = client.NoFrixionBizBizModelsPagingPaymentRequestPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `any[]` | No |  |
| `amount` | `number` | No |  |
| `amount_pending` | `number` | No |  |
| `amount_received` | `number` | No |  |
| `amount_refunded` | `number` | No |  |
| `auto_send_receipt` | `boolean` | No |  |
| `base_origin_url` | `string` | No |  |
| `callback_url` | `string` | No |  |
| `card_authorize_only` | `boolean` | No |  |
| `card_create_token` | `boolean` | No |  |
| `card_create_token_mode` | `string` | No |  |
| `card_ignore_cvn` | `boolean` | No |  |
| `card_processor_merchant_id` | `string` | No |  |
| `card_stripe_payment_intent_id` | `string` | No |  |
| `card_stripe_payment_intent_secret` | `string` | No |  |
| `created_by_user` | `Record<string, any>` | Yes |  |
| `currency` | `string` | No |  |
| `custom_field` | `any[]` | No |  |
| `customer_email_address` | `string` | No |  |
| `customer_id` | `string` | No |  |
| `customer_name` | `string` | No |  |
| `description` | `string` | No |  |
| `destination_account` | `Record<string, any>` | No |  |
| `direct_debit_payment` | `Record<string, any>` | No |  |
| `due_date` | `string` | No |  |
| `event` | `any[]` | No |  |
| `failure_callback_url` | `string` | No |  |
| `field_display_setting` | `any[]` | No |  |
| `formatted_amount` | `string` | No |  |
| `hosted_pay_checkout_url` | `string` | No |  |
| `id` | `string` | No |  |
| `ignore_address_verification` | `boolean` | No |  |
| `inserted` | `string` | No |  |
| `inserted_sortable` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `jwk` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `lightning_invoice` | `string` | No |  |
| `lightning_invoice_expires_at` | `string` | No |  |
| `merchant_direct_debit_mandate_id` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `notification_email_address` | `string` | No |  |
| `notification_role_i_d` | `any[]` | No |  |
| `order_id` | `string` | No |  |
| `partial_payment_method` | `string` | No |  |
| `partial_payment_step` | `string` | No |  |
| `payment_attempt` | `any[]` | No |  |
| `payment_method` | `any[]` | No |  |
| `payment_processor` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `pisp_account_id` | `string` | No |  |
| `priority_bank_id` | `string` | No |  |
| `result` | `Record<string, any>` | No |  |
| `sandbox_settle_delay_in_second` | `number` | No |  |
| `shipping_address` | `Record<string, any>` | No |  |
| `status` | `string` | No |  |
| `success_web_hook_url` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `title` | `string` | No |  |
| `tokenised_card` | `any[]` | No |  |
| `transaction` | `any[]` | No |  |
| `use_hosted_payment_page` | `boolean` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionBizBizModelsPagingPaymentRequestPage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionBizBizModelsPagingPaymentRequestPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionBizBizModelsPagingPayoutPageEntity

```ts
const no_frixion_biz_biz_models_paging_payout_page = client.NoFrixionBizBizModelsPagingPayoutPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `Record<string, any>` | Yes |  |
| `can_authorise` | `boolean` | No |  |
| `can_process` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `Record<string, any>` | No |  |
| `document` | `any[]` | No |  |
| `event` | `any[]` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `number` | No |  |
| `fx_destination_amount_minor_unit` | `number` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `number` | No |  |
| `fx_use_destination_amount` | `boolean` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_failed` | `boolean` | No |  |
| `is_settled` | `boolean` | No |  |
| `is_submitted` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `Record<string, any>` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `Record<string, any>` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `number` | No |  |
| `transacted_fx_amount` | `number` | No |  |
| `transacted_fx_rate` | `number` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionBizBizModelsPagingPayoutPage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionBizBizModelsPagingPayoutPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionBizBizModelsPagingPayrunPageEntity

```ts
const no_frixion_biz_biz_models_paging_payrun_page = client.NoFrixionBizBizModelsPagingPayrunPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation` | `any[]` | No |  |
| `authorisation_date` | `string` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_delete` | `boolean` | No |  |
| `can_edit` | `boolean` | No |  |
| `event` | `any[]` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice` | `any[]` | No |  |
| `invoices_minimal` | `any[]` | No |  |
| `is_archived` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `last_updated_by` | `Record<string, any>` | Yes |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment` | `any[]` | No |  |
| `payout` | `any[]` | No |  |
| `payouts_count` | `number` | No |  |
| `schedule_date` | `string` | No |  |
| `source_account` | `any[]` | No |  |
| `status` | `string` | No |  |
| `total_eur` | `number` | No |  |
| `total_gbp` | `number` | No |  |
| `total_usd` | `number` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionBizBizModelsPagingPayrunPage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionBizBizModelsPagingPayrunPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionBizBizModelsPagingRuleEventsPageEntity

```ts
const no_frixion_biz_biz_models_paging_rule_events_page = client.NoFrixionBizBizModelsPagingRuleEventsPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_message` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_authorise_to_enable` | `boolean` | No |  |
| `message` | `string` | No |  |
| `raw_response` | `string` | No |  |
| `rule_event_type` | `string` | No |  |
| `rule_id` | `string` | No |  |
| `user` | `Record<string, any>` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionBizBizModelsPagingRuleEventsPage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionBizBizModelsPagingRuleEventsPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionBizBizModelsPagingRulesPageEntity

```ts
const no_frixion_biz_biz_models_paging_rules_page = client.NoFrixionBizBizModelsPagingRulesPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account` | `Record<string, any>` | No |  |
| `account_id` | `string` | No |  |
| `approve_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `can_authorise` | `boolean` | No |  |
| `created_by` | `Record<string, any>` | Yes |  |
| `description` | `string` | No |  |
| `end_at` | `string` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_disabled` | `boolean` | No |  |
| `last_executed_at` | `string` | No |  |
| `last_run_at_transaction_date` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | Yes |  |
| `on_approved_web_hook_url` | `string` | No |  |
| `on_execution_error_web_hook_url` | `string` | No |  |
| `on_execution_success_web_hook_url` | `string` | No |  |
| `start_at` | `string` | No |  |
| `status` | `string` | No |  |
| `sweep_action` | `Record<string, any>` | No |  |
| `time_zone_id` | `string` | No |  |
| `trigger_cron_expression` | `string` | No |  |
| `trigger_on_pay_in` | `boolean` | No |  |
| `user_id` | `string` | No |  |
| `web_hook_secret` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionBizBizModelsPagingRulesPage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionBizBizModelsPagingRulesPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionBizBizModelsPaymentsCardPaymentEntity

```ts
const no_frixion_biz_biz_models_payments_card_payment = client.NoFrixionBizBizModelsPaymentsCardPayment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorized_amount` | `string` | No |  |
| `currency_code` | `string` | No |  |
| `is_payer_authentication_required` | `boolean` | No |  |
| `is_soft_decline` | `boolean` | No |  |
| `payer_authentication_access_token` | `string` | No |  |
| `payer_authentication_merchant_data` | `string` | No |  |
| `payer_authentication_url` | `string` | No |  |
| `payer_authentication_window_height` | `number` | No |  |
| `payer_authentication_window_width` | `number` | No |  |
| `payment_request_callback_url` | `string` | No |  |
| `payment_request_id` | `string` | No |  |
| `request_id` | `string` | No |  |
| `response_code` | `string` | No |  |
| `response_type` | `string` | No |  |
| `status` | `string` | No |  |
| `three_ds_redirect_url` | `string` | No |  |
| `transaction_id` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionBizBizModelsPaymentsCardPayment().create({
  paymentrequest_id: /* string */,
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionBizBizModelsPaymentsCardPaymentEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionBizBizModelsPaymentsCardPublicKeyEntity

```ts
const no_frixion_biz_biz_models_payments_card_public_key = client.NoFrixionBizBizModelsPaymentsCardPublicKey()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jwt` | `string` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionBizBizModelsPaymentsCardPublicKey().load({ paymentrequest_id: 'paymentrequest_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionBizBizModelsPaymentsCardPublicKeyEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity

```ts
const no_frixion_money_moov_api_features_beneficiaries_beneficiaries = client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `beneficiary` | `any[]` | No |  |
| `failed_beneficiary` | `Record<string, any>` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries().create({
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity

```ts
const no_frixion_money_moov_api_features_payment_requests_payment = client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_payment_request` | `Record<string, any>` | No |  |
| `payment_request` | `any[]` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment().create({
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity

```ts
const no_frixion_money_moov_api_features_permissions_roles_create = client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_role` | `Record<string, any>` | No |  |
| `role` | `any[]` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate().create({
  merchant_id: /* string */,
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity

```ts
const no_frixion_money_moov_api_features_user_invites_create = client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_user_invite` | `Record<string, any>` | No |  |
| `user_invite` | `any[]` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate().create({
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity

```ts
const no_frixion_money_moov_models_authorisation_settings_merchant = client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount_lower` | `number` | No |  |
| `amount_upper` | `number` | No |  |
| `authorisation_type` | `string` | No |  |
| `beneficiaries_only` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `last_editor_cant_authorise` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `number_of_authoriser` | `number` | No |  |
| `role_setting` | `any[]` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsBatchPayoutEntity

```ts
const no_frixion_money_moov_models_batch_payout = client.NoFrixionMoneyMoovModelsBatchPayout()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approve_url` | `string` | No |  |
| `id` | `string` | No |  |
| `payout` | `any[]` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovModelsBatchPayout().create({
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsBatchPayout().load({ id: 'no_frixion_money_moov_models_batch_payout_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsBatchPayoutEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity

```ts
const no_frixion_money_moov_models_beneficiary_group_page = client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `group_member` | `any[]` | No |  |
| `group_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsBeneficiaryPageEntity

```ts
const no_frixion_money_moov_models_beneficiary_page = client.NoFrixionMoneyMoovModelsBeneficiaryPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `beneficiary_event` | `any[]` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `created_by` | `Record<string, any>` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `Record<string, any>` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `boolean` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `any[]` | No |  |
| `their_reference` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsBeneficiaryPage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsCardCustomerTokenEntity

```ts
const no_frixion_money_moov_models_card_customer_token = client.NoFrixionMoneyMoovModelsCardCustomerToken()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `card_type` | `string` | No |  |
| `customer_email_address` | `string` | No |  |
| `expiry_month` | `string` | No |  |
| `expiry_year` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `last_four_digit` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `masked_card_number` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `payment_request_id` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsCardCustomerToken().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsCardCustomerToken().load({ customer_email_address: 'customer_email_address' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsCardCustomerToken().remove()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsCardCustomerTokenEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity

```ts
const no_frixion_money_moov_models_currency_currency_info = client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `code` | `string` | No |  |
| `decimal` | `number` | No |  |
| `is_fiat` | `boolean` | No |  |
| `iso4217_alpha_code` | `string` | No |  |
| `iso4217_numeric_code` | `string` | No |  |
| `symbol` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity

```ts
const no_frixion_money_moov_models_direct_debit_batch_submit = client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_submission` | `Record<string, any>` | No |  |
| `successful_submission` | `any[]` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit().create({
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsFxRateEntity

```ts
const no_frixion_money_moov_models_fx_rate = client.NoFrixionMoneyMoovModelsFxRate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destination_currency` | `string` | No |  |
| `exchange_rate` | `number` | No |  |
| `expiry_time` | `string` | No |  |
| `quote_id` | `string` | No |  |
| `source_currency` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsFxRate().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsFxRate().load({ destination: 'destination', source: 'source', valid_for_minute: 1 })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsFxRateEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsIPaymentEntity

```ts
const no_frixion_money_moov_models_i_payment = client.NoFrixionMoneyMoovModelsIPayment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `payment_request_id` | `string` | No |  |
| `response_type` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovModelsIPayment().create({
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsIPaymentEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsMandatesMandateEntity

```ts
const no_frixion_money_moov_models_mandates_mandate = client.NoFrixionMoneyMoovModelsMandatesMandate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_number` | `string` | No |  |
| `address_line1` | `string` | Yes |  |
| `address_line2` | `string` | No |  |
| `approved_at` | `string` | No |  |
| `city` | `string` | Yes |  |
| `country_code` | `string` | Yes |  |
| `currency` | `string` | No |  |
| `customer_account_number` | `string` | No |  |
| `customer_city` | `string` | No |  |
| `customer_country_code` | `string` | No |  |
| `customer_country_name` | `string` | No |  |
| `customer_email_address` | `string` | No |  |
| `customer_first_name` | `string` | No |  |
| `customer_iban` | `string` | No |  |
| `customer_last_name` | `string` | No |  |
| `customer_sort_code` | `string` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `iban` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_recurring` | `boolean` | No |  |
| `last_name` | `string` | Yes |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `postal_code` | `string` | Yes |  |
| `reference` | `string` | No |  |
| `sort_code` | `string` | No |  |
| `status` | `string` | No |  |
| `supplier_bank_account_id` | `string` | No |  |
| `supplier_customer_id` | `string` | No |  |
| `supplier_mandate_id` | `string` | No |  |
| `supplier_name` | `string` | No |  |
| `supplier_status` | `string` | No |  |

### Field Usage by Operation

| Field | load | create |
| --- | --- | --- |
| `account_number` | - | - |
| `address_line1` | - | - |
| `address_line2` | - | - |
| `approved_at` | - | - |
| `city` | - | - |
| `country_code` | - | - |
| `currency` | - | Yes |
| `customer_account_number` | - | - |
| `customer_city` | - | - |
| `customer_country_code` | - | - |
| `customer_country_name` | - | - |
| `customer_email_address` | - | - |
| `customer_first_name` | - | - |
| `customer_iban` | - | - |
| `customer_last_name` | - | - |
| `customer_sort_code` | - | - |
| `email_address` | - | - |
| `first_name` | - | - |
| `iban` | - | - |
| `id` | - | - |
| `inserted` | - | - |
| `is_recurring` | - | - |
| `last_name` | - | - |
| `last_updated` | - | - |
| `merchant_id` | - | Yes |
| `postal_code` | - | - |
| `reference` | - | - |
| `sort_code` | - | - |
| `status` | - | - |
| `supplier_bank_account_id` | - | - |
| `supplier_customer_id` | - | - |
| `supplier_mandate_id` | - | - |
| `supplier_name` | - | - |
| `supplier_status` | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovModelsMandatesMandate().create({
  address_line1: /* string */,
  city: /* string */,
  country_code: /* string */,
  email_address: /* string */,
  first_name: /* string */,
  last_name: /* string */,
  postal_code: /* string */,
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsMandatesMandate().load({ id: 'no_frixion_money_moov_models_mandates_mandate_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsMandatesMandateEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsMerchantEntity

```ts
const no_frixion_money_moov_models_merchant = client.NoFrixionMoneyMoovModelsMerchant()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_currency` | `any[]` | No |  |
| `can_have_trust_account` | `boolean` | No |  |
| `card_payment_processor` | `string` | No |  |
| `company_id` | `string` | No |  |
| `display_qr_on_hosted_pay` | `boolean` | No |  |
| `hosted_pay_version` | `number` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_blocked` | `boolean` | No |  |
| `is_exited` | `boolean` | No |  |
| `is_suspended` | `boolean` | No |  |
| `jurisdiction` | `string` | No |  |
| `logo_url_png` | `string` | No |  |
| `logo_url_svg` | `string` | No |  |
| `merchant_category_code` | `string` | No |  |
| `name` | `string` | No |  |
| `note` | `string` | No |  |
| `parent_merchant` | `Record<string, any>` | No |  |
| `payment_account` | `any[]` | No |  |
| `payment_account_limit` | `number` | No |  |
| `short_name` | `string` | No |  |
| `supported_payment_methods_list` | `any[]` | No |  |
| `suspension_reason` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `time_zone_id` | `string` | No |  |
| `trading_name` | `string` | No |  |
| `web_hook_limit` | `number` | No |  |
| `your_role_name` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsMerchant().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsMerchant().load({ id: 'no_frixion_money_moov_models_merchant_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.NoFrixionMoneyMoovModelsMerchant().update({
  id: 'no_frixion_money_moov_models_merchant_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsMerchantEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsMerchantPageEntity

```ts
const no_frixion_money_moov_models_merchant_page = client.NoFrixionMoneyMoovModelsMerchantPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_currency` | `any[]` | No |  |
| `can_have_trust_account` | `boolean` | No |  |
| `card_payment_processor` | `string` | No |  |
| `company_id` | `string` | No |  |
| `display_qr_on_hosted_pay` | `boolean` | No |  |
| `hosted_pay_version` | `number` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_blocked` | `boolean` | No |  |
| `is_exited` | `boolean` | No |  |
| `is_suspended` | `boolean` | No |  |
| `jurisdiction` | `string` | No |  |
| `logo_url_png` | `string` | No |  |
| `logo_url_svg` | `string` | No |  |
| `merchant_category_code` | `string` | No |  |
| `name` | `string` | No |  |
| `note` | `string` | No |  |
| `parent_merchant` | `Record<string, any>` | No |  |
| `payment_account` | `any[]` | No |  |
| `payment_account_limit` | `number` | No |  |
| `short_name` | `string` | No |  |
| `supported_payment_methods_list` | `any[]` | No |  |
| `suspension_reason` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `time_zone_id` | `string` | No |  |
| `trading_name` | `string` | No |  |
| `web_hook_limit` | `number` | No |  |
| `your_role_name` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsMerchantPage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsMerchantPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity

```ts
const no_frixion_money_moov_models_merchant_pay_by_bank_setting = client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bank_country_code` | `any[]` | No |  |
| `bank_id` | `string` | No |  |
| `bank_name` | `string` | No |  |
| `business_institution_id` | `string` | No |  |
| `currency` | `string` | No |  |
| `logo` | `string` | No |  |
| `message` | `string` | No |  |
| `message_image_url` | `string` | No |  |
| `order` | `number` | No |  |
| `personal_institution_id` | `string` | No |  |
| `processor` | `string` | No |  |
| `warning_heading` | `string` | No |  |
| `warning_message` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsMerchantTokenEntity

```ts
const no_frixion_money_moov_models_merchant_token = client.NoFrixionMoneyMoovModelsMerchantToken()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `can_authorise` | `boolean` | No |  |
| `description` | `string` | No |  |
| `expires_at` | `string` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `hmac_algorithm` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `ip_address_whitelist` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_enabled` | `boolean` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `nonce` | `string` | Yes |  |
| `permission_type` | `any[]` | No |  |
| `request_signature_version` | `number` | No |  |
| `shared_secret_algorithm` | `string` | No |  |
| `shared_secret_base64` | `string` | No |  |
| `token` | `string` | No |  |

### Field Usage by Operation

| Field | load | create | update |
| --- | --- | --- | --- |
| `authentication_method` | - | - | - |
| `authorisation` | - | - | - |
| `authorisers_completed_count` | - | - | - |
| `authorisers_required_count` | - | - | - |
| `can_authorise` | - | - | - |
| `description` | - | Yes | - |
| `expires_at` | - | - | - |
| `has_current_user_authorised` | - | - | - |
| `hmac_algorithm` | - | - | - |
| `id` | - | - | - |
| `inserted` | - | - | - |
| `ip_address_whitelist` | - | - | - |
| `is_archived` | - | - | - |
| `is_enabled` | - | - | - |
| `last_authorised` | - | - | - |
| `last_updated` | - | - | - |
| `merchant_id` | - | Yes | - |
| `nonce` | - | - | - |
| `permission_type` | - | - | - |
| `request_signature_version` | - | - | - |
| `shared_secret_algorithm` | - | - | - |
| `shared_secret_base64` | - | - | - |
| `token` | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovModelsMerchantToken().create({
  nonce: /* string */,
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsMerchantToken().load({ id: 'no_frixion_money_moov_models_merchant_token_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.NoFrixionMoneyMoovModelsMerchantToken().update({
  id: 'no_frixion_money_moov_models_merchant_token_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsMerchantTokenPageEntity

```ts
const no_frixion_money_moov_models_merchant_token_page = client.NoFrixionMoneyMoovModelsMerchantTokenPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `can_authorise` | `boolean` | No |  |
| `description` | `string` | No |  |
| `expires_at` | `string` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `ip_address_whitelist` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_enabled` | `boolean` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `nonce` | `string` | Yes |  |
| `permission_type` | `any[]` | No |  |
| `request_signature_version` | `number` | No |  |
| `shared_secret_algorithm` | `string` | No |  |
| `shared_secret_base64` | `string` | No |  |
| `token` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsMerchantTokenPage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsNoFrixionVersionEntity

```ts
const no_frixion_money_moov_models_no_frixion_version = client.NoFrixionMoneyMoovModelsNoFrixionVersion()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `build_version` | `number` | No |  |
| `major_version` | `number` | No |  |
| `minor_version` | `number` | No |  |
| `release_name` | `string` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsNoFrixionVersion().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsNoFrixionVersionEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsOpenBankingAccountEntity

```ts
const no_frixion_money_moov_models_open_banking_account = client.NoFrixionMoneyMoovModelsOpenBankingAccount()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_balance` | `any[]` | No |  |
| `account_identification` | `any[]` | No |  |
| `account_name` | `any[]` | No |  |
| `account_type` | `string` | No |  |
| `balance` | `number` | No |  |
| `consolidated_account_information` | `Record<string, any>` | No |  |
| `currency` | `string` | No |  |
| `description` | `string` | No |  |
| `detail` | `string` | No |  |
| `id` | `string` | No |  |
| `nickname` | `string` | No |  |
| `type` | `string` | No |  |
| `usage_type` | `string` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsOpenBankingAccount().load({ id: 'no_frixion_money_moov_models_open_banking_account_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsOpenBankingAccountEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsOpenBankingConsentEntity

```ts
const no_frixion_money_moov_models_open_banking_consent = client.NoFrixionMoneyMoovModelsOpenBankingConsent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_url` | `string` | No |  |
| `callback_url` | `string` | No |  |
| `consent_id` | `string` | No |  |
| `email_address` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `failure_callback_url` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `institution_id` | `string` | No |  |
| `is_connected_account` | `boolean` | No |  |
| `is_enabled` | `boolean` | No |  |
| `merchant_id` | `string` | No |  |
| `provider` | `string` | No |  |
| `success_web_hook_url` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `authorisation_url` | - | - | - | - | - |
| `callback_url` | - | - | - | - | - |
| `consent_id` | - | - | - | - | - |
| `email_address` | - | - | - | - | - |
| `expiry_date` | - | - | - | - | - |
| `failure_callback_url` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `inserted` | - | - | - | - | - |
| `institution_id` | - | - | Yes | - | - |
| `is_connected_account` | - | - | - | - | - |
| `is_enabled` | - | - | - | - | - |
| `merchant_id` | - | - | Yes | - | - |
| `provider` | - | - | - | - | - |
| `success_web_hook_url` | - | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovModelsOpenBankingConsent().create({
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsOpenBankingConsent().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsOpenBankingConsent().load({ id: 'no_frixion_money_moov_models_open_banking_consent_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsOpenBankingConsent().remove({ id: 'no_frixion_money_moov_models_open_banking_consent_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.NoFrixionMoneyMoovModelsOpenBankingConsent().update({
  id: 'no_frixion_money_moov_models_open_banking_consent_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsOpenBankingConsentEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsOpenBankingTransactionEntity

```ts
const no_frixion_money_moov_models_open_banking_transaction = client.NoFrixionMoneyMoovModelsOpenBankingTransaction()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address_detail` | `Record<string, any>` | No |  |
| `amount` | `number` | No |  |
| `balance` | `Record<string, any>` | No |  |
| `booking_date_time` | `string` | No |  |
| `charge_detail` | `Record<string, any>` | No |  |
| `currency` | `string` | No |  |
| `currency_exchange` | `Record<string, any>` | No |  |
| `date` | `string` | No |  |
| `description` | `string` | No |  |
| `enrichment` | `Record<string, any>` | No |  |
| `gross_amount` | `Record<string, any>` | Yes |  |
| `id` | `string` | No |  |
| `iso_bank_transaction_code` | `Record<string, any>` | No |  |
| `merchant` | `Record<string, any>` | No |  |
| `payee_detail` | `Record<string, any>` | Yes |  |
| `payer_detail` | `Record<string, any>` | Yes |  |
| `proprietary_bank_transaction_code` | `Record<string, any>` | No |  |
| `reference` | `string` | No |  |
| `statement_reference` | `any[]` | No |  |
| `status` | `string` | No |  |
| `supplementary_data` | `any` | No |  |
| `transaction_amount` | `Record<string, any>` | Yes |  |
| `transaction_information` | `any[]` | No |  |
| `transaction_mutability` | `string` | No |  |
| `value_date_time` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsOpenBankingTransaction().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsOpenBankingTransactionEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPaymentEntity

```ts
const no_frixion_money_moov_models_payment = client.NoFrixionMoneyMoovModelsPayment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `any[]` | No |  |
| `amount` | `number` | No |  |
| `amount_pending` | `number` | No |  |
| `amount_received` | `number` | No |  |
| `amount_refunded` | `number` | No |  |
| `auto_send_receipt` | `boolean` | No |  |
| `base_origin_url` | `string` | No |  |
| `callback_url` | `string` | No |  |
| `card_authorize_only` | `boolean` | No |  |
| `card_create_token` | `boolean` | No |  |
| `card_create_token_mode` | `string` | No |  |
| `card_ignore_cvn` | `boolean` | No |  |
| `card_no_payer_authentication` | `boolean` | No |  |
| `card_processor_merchant_id` | `string` | No |  |
| `card_stripe_payment_intent_id` | `string` | No |  |
| `card_stripe_payment_intent_secret` | `string` | No |  |
| `card_transmit_raw_detail` | `boolean` | No |  |
| `created_by_user` | `Record<string, any>` | Yes |  |
| `currency` | `string` | No |  |
| `custom_field` | `any[]` | No |  |
| `customer_email_address` | `string` | No |  |
| `customer_id` | `string` | No |  |
| `customer_name` | `string` | No |  |
| `description` | `string` | No |  |
| `destination_account` | `Record<string, any>` | No |  |
| `direct_debit_payment` | `Record<string, any>` | No |  |
| `due_date` | `string` | No |  |
| `event` | `any[]` | No |  |
| `failure_callback_url` | `string` | No |  |
| `field_display_setting` | `any[]` | No |  |
| `formatted_amount` | `string` | No |  |
| `hosted_pay_checkout_url` | `string` | No |  |
| `id` | `string` | No |  |
| `ignore_address_verification` | `boolean` | No |  |
| `inserted` | `string` | No |  |
| `inserted_sortable` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `jwk` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `lightning_invoice` | `string` | No |  |
| `lightning_invoice_expires_at` | `string` | No |  |
| `merchant_direct_debit_mandate_id` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `notification_email_address` | `string` | No |  |
| `notification_role_i_d` | `any[]` | No |  |
| `order_id` | `string` | No |  |
| `partial_payment_method` | `string` | No |  |
| `partial_payment_step` | `string` | No |  |
| `payment_attempt` | `any[]` | No |  |
| `payment_method` | `any[]` | No |  |
| `payment_processor` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `pisp_account_id` | `string` | No |  |
| `priority_bank_id` | `string` | No |  |
| `result` | `Record<string, any>` | No |  |
| `sandbox_settle_delay_in_second` | `number` | No |  |
| `shipping_address` | `Record<string, any>` | No |  |
| `shipping_address_city` | `string` | No |  |
| `shipping_address_country_code` | `string` | No |  |
| `shipping_address_county` | `string` | No |  |
| `shipping_address_line1` | `string` | No |  |
| `shipping_address_line2` | `string` | No |  |
| `shipping_address_post_code` | `string` | No |  |
| `shipping_email` | `string` | No |  |
| `shipping_first_name` | `string` | No |  |
| `shipping_last_name` | `string` | No |  |
| `shipping_phone` | `string` | No |  |
| `status` | `string` | No |  |
| `success_web_hook_url` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `tag_id` | `any[]` | No |  |
| `title` | `string` | No |  |
| `tokenised_card` | `any[]` | No |  |
| `transaction` | `any[]` | No |  |
| `use_hosted_payment_page` | `boolean` | No |  |

### Field Usage by Operation

| Field | load | create | update |
| --- | --- | --- | --- |
| `address` | - | - | - |
| `amount` | - | Yes | - |
| `amount_pending` | - | - | - |
| `amount_received` | - | - | - |
| `amount_refunded` | - | - | - |
| `auto_send_receipt` | - | - | - |
| `base_origin_url` | - | - | - |
| `callback_url` | - | - | - |
| `card_authorize_only` | - | - | - |
| `card_create_token` | - | - | - |
| `card_create_token_mode` | - | - | - |
| `card_ignore_cvn` | - | - | - |
| `card_no_payer_authentication` | - | - | - |
| `card_processor_merchant_id` | - | - | - |
| `card_stripe_payment_intent_id` | - | - | - |
| `card_stripe_payment_intent_secret` | - | - | - |
| `card_transmit_raw_detail` | - | - | - |
| `created_by_user` | - | - | - |
| `currency` | - | - | - |
| `custom_field` | - | - | - |
| `customer_email_address` | - | - | - |
| `customer_id` | - | - | - |
| `customer_name` | - | - | - |
| `description` | - | - | - |
| `destination_account` | - | - | - |
| `direct_debit_payment` | - | - | - |
| `due_date` | - | - | - |
| `event` | - | - | - |
| `failure_callback_url` | - | - | - |
| `field_display_setting` | - | - | - |
| `formatted_amount` | - | - | - |
| `hosted_pay_checkout_url` | - | - | - |
| `id` | - | - | - |
| `ignore_address_verification` | - | - | - |
| `inserted` | - | - | - |
| `inserted_sortable` | - | - | - |
| `is_archived` | - | - | - |
| `jwk` | - | - | - |
| `last_updated` | - | - | - |
| `lightning_invoice` | - | - | - |
| `lightning_invoice_expires_at` | - | - | - |
| `merchant_direct_debit_mandate_id` | - | - | - |
| `merchant_id` | - | - | - |
| `merchant_token_description` | - | - | - |
| `notification_email_address` | - | - | - |
| `notification_role_i_d` | - | - | - |
| `order_id` | - | - | - |
| `partial_payment_method` | - | - | - |
| `partial_payment_step` | - | - | - |
| `payment_attempt` | - | - | - |
| `payment_method` | - | - | - |
| `payment_processor` | - | - | - |
| `payrun_id` | - | - | - |
| `pisp_account_id` | - | - | - |
| `priority_bank_id` | - | - | - |
| `result` | - | - | - |
| `sandbox_settle_delay_in_second` | - | - | - |
| `shipping_address` | - | - | - |
| `shipping_address_city` | - | - | - |
| `shipping_address_country_code` | - | - | - |
| `shipping_address_county` | - | - | - |
| `shipping_address_line1` | - | - | - |
| `shipping_address_line2` | - | - | - |
| `shipping_address_post_code` | - | - | - |
| `shipping_email` | - | - | - |
| `shipping_first_name` | - | - | - |
| `shipping_last_name` | - | - | - |
| `shipping_phone` | - | - | - |
| `status` | - | - | - |
| `success_web_hook_url` | - | - | - |
| `tag` | - | - | - |
| `tag_id` | - | - | - |
| `title` | - | - | - |
| `tokenised_card` | - | - | - |
| `transaction` | - | - | - |
| `use_hosted_payment_page` | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovModelsPayment().create({
  created_by_user: /* Record<string, any> */,
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsPayment().load({ id: 'no_frixion_money_moov_models_payment_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.NoFrixionMoneyMoovModelsPayment().update({
  id: 'no_frixion_money_moov_models_payment_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPaymentEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity

```ts
const no_frixion_money_moov_models_payment_account_minimal_page = client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_name` | `string` | No |  |
| `available_balance` | `number` | No |  |
| `balance` | `number` | No |  |
| `balance_minor_unit` | `number` | No |  |
| `currency` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `Record<string, any>` | Yes |  |
| `is_archived` | `boolean` | No |  |
| `is_connected_account` | `boolean` | No |  |
| `merchant_id` | `string` | No |  |
| `submitted_payouts_balance` | `number` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPaymentAccountPageEntity

```ts
const no_frixion_money_moov_models_payment_account_page = client.NoFrixionMoneyMoovModelsPaymentAccountPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_name` | `string` | No |  |
| `account_supplier_name` | `string` | No |  |
| `available_balance` | `number` | No |  |
| `available_balance_minor_unit` | `number` | No |  |
| `balance` | `number` | No |  |
| `balance_minor_unit` | `number` | No |  |
| `bank_name` | `string` | No |  |
| `consent_id` | `string` | No |  |
| `created_by` | `Record<string, any>` | Yes |  |
| `created_by_display_name` | `string` | No |  |
| `currency` | `string` | No |  |
| `default_payment_rail` | `string` | No |  |
| `display_name` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `external_account_icon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `Record<string, any>` | Yes |  |
| `inserted` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_connected_account` | `boolean` | No |  |
| `is_default` | `boolean` | No |  |
| `is_trust_account` | `boolean` | No |  |
| `is_virtual` | `boolean` | No |  |
| `last_transaction` | `Record<string, any>` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `physical_account_id` | `string` | No |  |
| `rule` | `any[]` | No |  |
| `submitted_payouts_balance` | `number` | No |  |
| `submitted_payouts_balance_minor_unit` | `number` | No |  |
| `summary` | `string` | No |  |
| `supplier_sepa_instant_status` | `string` | No |  |
| `xero_bank_feed_connection_status` | `string` | No |  |
| `xero_bank_feed_last_synced_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` | No |  |
| `xero_bank_feed_sync_status` | `string` | No |  |
| `xero_unsynchronised_transactions_count` | `number` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsPaymentAccountPage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPaymentInitiationEntity

```ts
const no_frixion_money_moov_models_payment_initiation = client.NoFrixionMoneyMoovModelsPaymentInitiation()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `payment_initiation_id` | `string` | No |  |
| `payment_request_callback_url` | `string` | No |  |
| `payment_request_id` | `string` | No |  |
| `redirect_url` | `string` | No |  |
| `response_type` | `string` | No |  |
| `specific_error_message` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovModelsPaymentInitiation().create({
  paymentrequest_id: /* string */,
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPaymentInitiationEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPaymentRequestEventEntity

```ts
const no_frixion_money_moov_models_payment_request_event = client.NoFrixionMoneyMoovModelsPaymentRequestEvent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `number` | Yes |  |
| `apple_pay_transaction_id` | `string` | No |  |
| `card_authorization_response_id` | `string` | No |  |
| `card_expiry_month` | `number` | No |  |
| `card_expiry_year` | `number` | No |  |
| `card_issuer` | `string` | No |  |
| `card_issuer_country` | `string` | No |  |
| `card_last_four_digit` | `string` | No |  |
| `card_request_id` | `string` | No |  |
| `card_scheme` | `string` | No |  |
| `card_token_customer_id` | `string` | No |  |
| `card_transaction_id` | `string` | No |  |
| `currency` | `string` | No |  |
| `direct_debit_payment_id` | `string` | No |  |
| `direct_debit_payment_reference` | `string` | No |  |
| `drirect_debit_mandate_id` | `string` | No |  |
| `error_message` | `string` | No |  |
| `error_reason` | `string` | No |  |
| `event_type` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lightning_invoice` | `string` | No |  |
| `lightning_r_hash` | `string` | No |  |
| `origin_url` | `string` | No |  |
| `payment_method_type` | `string` | No |  |
| `payment_processor_name` | `string` | No |  |
| `payment_request_id` | `string` | No |  |
| `pisp_bank_status` | `string` | No |  |
| `pisp_payment_initiation_id` | `string` | No |  |
| `pisp_payment_institution_name` | `string` | No |  |
| `pisp_payment_service_provider_id` | `string` | No |  |
| `pisp_redirect_url` | `string` | No |  |
| `reconciled_transaction_id` | `string` | No |  |
| `refund_payout_id` | `string` | No |  |
| `status` | `string` | No |  |
| `wallet_name` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsPaymentRequestEvent().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestEventEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPaymentRequestMetricEntity

```ts
const no_frixion_money_moov_models_payment_request_metric = client.NoFrixionMoneyMoovModelsPaymentRequestMetric()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `all` | `number` | No |  |
| `authorized` | `number` | No |  |
| `paid` | `number` | No |  |
| `partially_paid` | `number` | No |  |
| `total_amounts_by_currency` | `Record<string, any>` | No |  |
| `unpaid` | `number` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsPaymentRequestMetric().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMetricEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity

```ts
const no_frixion_money_moov_models_payment_request_minimal = client.NoFrixionMoneyMoovModelsPaymentRequestMinimal()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `number` | No |  |
| `amount_pending` | `number` | No |  |
| `amount_received` | `number` | No |  |
| `amount_refunded` | `number` | No |  |
| `callback_url` | `string` | No |  |
| `card_stripe_payment_intent_secret` | `string` | No |  |
| `country_code` | `string` | No |  |
| `currency` | `string` | No |  |
| `custom_fields_to_display` | `any[]` | No |  |
| `description` | `string` | No |  |
| `due_date` | `string` | No |  |
| `field_display_setting` | `any[]` | No |  |
| `google_pay_merchant_id` | `string` | No |  |
| `id` | `string` | No |  |
| `jwk` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_logo_url_png` | `string` | No |  |
| `merchant_logo_url_svg` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `merchant_short_name` | `string` | No |  |
| `partial_payment_method` | `string` | No |  |
| `payment_attempt` | `any[]` | No |  |
| `payment_methods_list` | `any[]` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_processor_key` | `string` | No |  |
| `pisp_error` | `string` | No |  |
| `priority_bank_id` | `string` | No |  |
| `status` | `string` | No |  |
| `stripe_account_id` | `string` | No |  |
| `title` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsPaymentRequestMinimal().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPaymentRequestResultEntity

```ts
const no_frixion_money_moov_models_payment_request_result = client.NoFrixionMoneyMoovModelsPaymentRequestResult()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `number` | No |  |
| `amount_pending` | `number` | No |  |
| `amount_received` | `number` | No |  |
| `amount_refunded` | `number` | No |  |
| `currency` | `string` | No |  |
| `customer_id` | `string` | No |  |
| `payment` | `any[]` | No |  |
| `payment_request_id` | `string` | No |  |
| `pisp_authorization` | `any[]` | No |  |
| `requested_amount` | `number` | No |  |
| `result` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsPaymentRequestResult().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestResultEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity

```ts
const no_frixion_money_moov_models_payment_requests_merchant_payment = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `description` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `template` | `Record<string, any>` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity

```ts
const no_frixion_money_moov_models_payment_requests_merchant_payment2 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `description` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `template` | `Record<string, any>` | Yes |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2().load({ paymentrequest_id: 'paymentrequest_id', template_id: 'template_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity

```ts
const no_frixion_money_moov_models_payment_requests_merchant_payment3 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `description` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `template` | `Record<string, any>` | Yes |  |

### Operations

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3().update({
  paymentrequest_id: 'paymentrequest_id',
  template_id: 'template_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity

```ts
const no_frixion_money_moov_models_payment_requests_merchant_payment4 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4()
```

### Operations

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4().remove({ paymentrequest_id: 'paymentrequest_id', template_id: 'template_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPayoutKeysetPageEntity

```ts
const no_frixion_money_moov_models_payout_keyset_page = client.NoFrixionMoneyMoovModelsPayoutKeysetPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `Record<string, any>` | Yes |  |
| `can_authorise` | `boolean` | No |  |
| `can_process` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `Record<string, any>` | No |  |
| `document` | `any[]` | No |  |
| `event` | `any[]` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `number` | No |  |
| `fx_destination_amount_minor_unit` | `number` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `number` | No |  |
| `fx_use_destination_amount` | `boolean` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_failed` | `boolean` | No |  |
| `is_settled` | `boolean` | No |  |
| `is_submitted` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `Record<string, any>` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `Record<string, any>` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `number` | No |  |
| `transacted_fx_amount` | `number` | No |  |
| `transacted_fx_rate` | `number` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsPayoutKeysetPage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPayoutKeysetPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPayoutMetricEntity

```ts
const no_frixion_money_moov_models_payout_metric = client.NoFrixionMoneyMoovModelsPayoutMetric()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `all` | `number` | No |  |
| `failed` | `number` | No |  |
| `in_progress` | `number` | No |  |
| `paid` | `number` | No |  |
| `pending_approval` | `number` | No |  |
| `scheduled` | `number` | No |  |
| `total_amounts_by_currency` | `Record<string, any>` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsPayoutMetric().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPayoutMetricEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity

```ts
const no_frixion_money_moov_models_payouts_payouts_create = client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_payout` | `Record<string, any>` | No |  |
| `payout` | `any[]` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate().create({
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsPayrunEntity

```ts
const no_frixion_money_moov_models_payrun = client.NoFrixionMoneyMoovModelsPayrun()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation` | `any[]` | No |  |
| `authorisation_date` | `string` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_delete` | `boolean` | No |  |
| `can_edit` | `boolean` | No |  |
| `event` | `any[]` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice` | `any[]` | No |  |
| `invoices_minimal` | `any[]` | No |  |
| `is_archived` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `last_updated_by` | `Record<string, any>` | Yes |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment` | `any[]` | No |  |
| `payout` | `any[]` | No |  |
| `payouts_count` | `number` | No |  |
| `reason` | `string` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled_date` | `string` | No |  |
| `source_account` | `any[]` | No |  |
| `status` | `string` | No |  |
| `total_eur` | `number` | No |  |
| `total_gbp` | `number` | No |  |
| `total_usd` | `number` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovModelsPayrun().create({
  id: /* string */,
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsPayrun().load({ id: 'no_frixion_money_moov_models_payrun_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.NoFrixionMoneyMoovModelsPayrun().update({
  id: 'no_frixion_money_moov_models_payrun_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsPayrunEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsReportResultEntity

```ts
const no_frixion_money_moov_models_report_result = client.NoFrixionMoneyMoovModelsReportResult()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content` | `string` | No |  |
| `content_type` | `string` | No |  |
| `last_completed_at` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `report_name` | `string` | No |  |
| `report_type` | `string` | No |  |
| `statement_number` | `number` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsReportResult().load({ id: 1, report_id: 'report_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsReportResultEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsRuleEntity

```ts
const no_frixion_money_moov_models_rule = client.NoFrixionMoneyMoovModelsRule()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account` | `Record<string, any>` | No |  |
| `account_id` | `string` | No |  |
| `approve_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `can_authorise` | `boolean` | No |  |
| `created_by` | `Record<string, any>` | Yes |  |
| `description` | `string` | No |  |
| `end_at` | `string` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_disabled` | `boolean` | No |  |
| `last_executed_at` | `string` | No |  |
| `last_run_at_transaction_date` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | Yes |  |
| `on_approved_web_hook_url` | `string` | No |  |
| `on_execution_error_web_hook_url` | `string` | No |  |
| `on_execution_success_web_hook_url` | `string` | No |  |
| `start_at` | `string` | No |  |
| `status` | `string` | No |  |
| `sweep_action` | `Record<string, any>` | No |  |
| `time_zone_id` | `string` | No |  |
| `trigger_cron_expression` | `string` | No |  |
| `trigger_on_pay_in` | `boolean` | No |  |
| `user_id` | `string` | No |  |
| `web_hook_secret` | `string` | No |  |

### Field Usage by Operation

| Field | load | create | update |
| --- | --- | --- | --- |
| `account` | - | - | - |
| `account_id` | - | - | - |
| `approve_url` | - | - | - |
| `approver_id` | - | - | - |
| `authentication_method` | - | - | - |
| `authorisation` | - | - | - |
| `authorisers_completed_count` | - | - | - |
| `authorisers_required_count` | - | - | - |
| `can_authorise` | - | - | - |
| `created_by` | - | - | - |
| `description` | - | - | - |
| `end_at` | - | - | - |
| `has_current_user_authorised` | - | - | - |
| `id` | - | - | - |
| `inserted` | - | - | - |
| `is_disabled` | - | - | - |
| `last_executed_at` | - | - | - |
| `last_run_at_transaction_date` | - | - | - |
| `last_updated` | - | - | - |
| `merchant_id` | - | - | - |
| `name` | - | Yes | - |
| `nonce` | - | - | - |
| `on_approved_web_hook_url` | - | - | - |
| `on_execution_error_web_hook_url` | - | - | - |
| `on_execution_success_web_hook_url` | - | - | - |
| `start_at` | - | - | - |
| `status` | - | - | - |
| `sweep_action` | - | Yes | - |
| `time_zone_id` | - | - | - |
| `trigger_cron_expression` | - | - | - |
| `trigger_on_pay_in` | - | - | - |
| `user_id` | - | - | - |
| `web_hook_secret` | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovModelsRule().create({
  created_by: /* Record<string, any> */,
  nonce: /* string */,
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsRule().load({ id: 'no_frixion_money_moov_models_rule_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.NoFrixionMoneyMoovModelsRule().update({
  id: 'no_frixion_money_moov_models_rule_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsRuleEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsTransactionEntity

```ts
const no_frixion_money_moov_models_transaction = client.NoFrixionMoneyMoovModelsTransaction()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `account_name` | `string` | No |  |
| `account_sequence_number` | `number` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `balance` | `number` | No |  |
| `balance_minor_unit` | `number` | No |  |
| `counterparty` | `Record<string, any>` | No |  |
| `counterparty_summary` | `string` | No |  |
| `currency` | `string` | No |  |
| `description` | `string` | No |  |
| `fx_amount` | `number` | No |  |
| `fx_currency` | `string` | No |  |
| `fx_rate` | `number` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `payment_request_custom_field` | `Record<string, any>` | No |  |
| `payment_request_id` | `string` | No |  |
| `payout_id` | `string` | No |  |
| `raw_reference` | `string` | No |  |
| `rule_id` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `their_reference` | `string` | No |  |
| `transaction_date` | `string` | No |  |
| `type` | `string` | No |  |
| `virtual_iban` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsTransaction().load({ id: 'no_frixion_money_moov_models_transaction_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsTransactionEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsTransactionPageEntity

```ts
const no_frixion_money_moov_models_transaction_page = client.NoFrixionMoneyMoovModelsTransactionPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `account_name` | `string` | No |  |
| `account_sequence_number` | `number` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `balance` | `number` | No |  |
| `balance_minor_unit` | `number` | No |  |
| `content` | `any[]` | No |  |
| `counterparty` | `Record<string, any>` | No |  |
| `counterparty_summary` | `string` | No |  |
| `currency` | `string` | No |  |
| `description` | `string` | No |  |
| `fx_amount` | `number` | No |  |
| `fx_currency` | `string` | No |  |
| `fx_rate` | `number` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `page_number` | `number` | No |  |
| `page_size` | `number` | No |  |
| `payment_request_custom_field` | `Record<string, any>` | No |  |
| `payment_request_id` | `string` | No |  |
| `payout_id` | `string` | No |  |
| `raw_reference` | `string` | No |  |
| `rule_id` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `their_reference` | `string` | No |  |
| `total_page` | `number` | No |  |
| `total_size` | `number` | No |  |
| `transaction_date` | `string` | No |  |
| `type` | `string` | No |  |
| `virtual_iban` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsTransactionPage().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsTransactionPage().load({ account_id: 'account_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsTransactionPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsUserInviteEntity

```ts
const no_frixion_money_moov_models_user_invite = client.NoFrixionMoneyMoovModelsUserInvite()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_status` | `Record<string, any>` | No |  |
| `id` | `string` | No |  |
| `initial_role_id` | `string` | No |  |
| `invitee_email_address` | `string` | No |  |
| `invitee_first_name` | `string` | No |  |
| `invitee_last_name` | `string` | No |  |
| `inviter_email_address` | `string` | No |  |
| `inviter_first_name` | `string` | No |  |
| `inviter_last_name` | `string` | No |  |
| `is_authorised` | `boolean` | No |  |
| `is_invitee_registered` | `boolean` | No |  |
| `last_invited` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `message` | `string` | No |  |
| `registration_url` | `string` | No |  |
| `send_invite_email` | `boolean` | No |  |
| `status` | `string` | No |  |
| `user` | `Record<string, any>` | Yes |  |
| `user_id` | `string` | No |  |

### Field Usage by Operation

| Field | load | create |
| --- | --- | --- |
| `authorisation_status` | - | - |
| `id` | - | - |
| `initial_role_id` | - | - |
| `invitee_email_address` | - | Yes |
| `invitee_first_name` | - | - |
| `invitee_last_name` | - | - |
| `inviter_email_address` | - | - |
| `inviter_first_name` | - | - |
| `inviter_last_name` | - | - |
| `is_authorised` | - | - |
| `is_invitee_registered` | - | - |
| `last_invited` | - | - |
| `merchant_id` | - | - |
| `merchant_name` | - | - |
| `message` | - | - |
| `registration_url` | - | - |
| `send_invite_email` | - | - |
| `status` | - | - |
| `user` | - | - |
| `user_id` | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovModelsUserInvite().create({
  user: /* Record<string, any> */,
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsUserInvite().load({ id: 'no_frixion_money_moov_models_user_invite_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsUserInviteEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsUserInvitePageEntity

```ts
const no_frixion_money_moov_models_user_invite_page = client.NoFrixionMoneyMoovModelsUserInvitePage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_status` | `Record<string, any>` | No |  |
| `id` | `string` | No |  |
| `initial_role_id` | `string` | No |  |
| `invitee_email_address` | `string` | No |  |
| `invitee_first_name` | `string` | No |  |
| `invitee_last_name` | `string` | No |  |
| `inviter_email_address` | `string` | No |  |
| `inviter_first_name` | `string` | No |  |
| `inviter_last_name` | `string` | No |  |
| `is_authorised` | `boolean` | No |  |
| `is_invitee_registered` | `boolean` | No |  |
| `last_invited` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `message` | `string` | No |  |
| `registration_url` | `string` | No |  |
| `status` | `string` | No |  |
| `user` | `Record<string, any>` | Yes |  |
| `user_id` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsUserInvitePage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsUserInvitePageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsUserPageEntity

```ts
const no_frixion_money_moov_models_user_page = client.NoFrixionMoneyMoovModelsUserPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `any[]` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `boolean` | No |  |
| `permission` | `Record<string, any>` | No |  |
| `roles_with_scope` | `any[]` | No |  |
| `two_factor_enabled` | `boolean` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsUserPage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsUserPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionMoneyMoovModelsWebhookEntity

```ts
const no_frixion_money_moov_models_webhook = client.NoFrixionMoneyMoovModelsWebhook()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destination_url` | `string` | No |  |
| `email_address` | `string` | No |  |
| `failed_notification_email_address` | `string` | No |  |
| `id` | `string` | No |  |
| `is_active` | `boolean` | No |  |
| `merchant_id` | `string` | No |  |
| `notification_method` | `string` | No |  |
| `resource_type` | `any[]` | No |  |
| `retry` | `boolean` | No |  |
| `secret` | `string` | No |  |
| `version` | `number` | No |  |

### Field Usage by Operation

| Field | load | list | create | update |
| --- | --- | --- | --- | --- |
| `destination_url` | - | - | - | - |
| `email_address` | - | - | - | - |
| `failed_notification_email_address` | - | - | - | - |
| `id` | - | - | - | - |
| `is_active` | - | - | - | - |
| `merchant_id` | - | - | Yes | Yes |
| `notification_method` | - | - | Yes | Yes |
| `resource_type` | - | - | - | - |
| `retry` | - | - | - | - |
| `secret` | - | - | - | - |
| `version` | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.NoFrixionMoneyMoovModelsWebhook().create({
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.NoFrixionMoneyMoovModelsWebhook().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionMoneyMoovModelsWebhook().load({ id: 'no_frixion_money_moov_models_webhook_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.NoFrixionMoneyMoovModelsWebhook().update({
  id: 'no_frixion_money_moov_models_webhook_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionMoneyMoovModelsWebhookEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## OpenBankingEntity

```ts
const open_banking = client.OpenBanking()
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.OpenBanking().create({
  account_id: /* string */,
})
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.OpenBanking().remove()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `OpenBankingEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PayeeverificationEntity

```ts
const payeeverification = client.Payeeverification()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_name` | `string` | Yes |  |
| `account_number` | `string` | No |  |
| `iban` | `string` | Yes |  |
| `payee_verified_account_name` | `string` | No |  |
| `result` | `string` | No |  |
| `secondary_identification` | `string` | No |  |
| `sort_code` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Payeeverification().create({
  account_name: /* string */,
  iban: /* string */,
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PayeeverificationEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PaymentRequestEntity

```ts
const payment_request = client.PaymentRequest()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `number` | No |  |
| `do_simulate_settlement_failure` | `boolean` | No |  |
| `error_description` | `string` | No |  |
| `institution` | `string` | No |  |
| `payment_initiation_id` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PaymentRequest().create({
  paymentrequest_id: /* string */,
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.PaymentRequest().load()
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.PaymentRequest().remove({ id: 'id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.PaymentRequest().update({
  paymentrequest_id: 'paymentrequest_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PaymentRequestEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PayoutEntity

```ts
const payout = client.Payout()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `allow_incomplete` | `boolean` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `Record<string, any>` | Yes |  |
| `beneficiary_id` | `string` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_process` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `Record<string, any>` | No |  |
| `document` | `any[]` | No |  |
| `event` | `any[]` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `number` | No |  |
| `fx_destination_amount_minor_unit` | `number` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `number` | No |  |
| `fx_use_destination_amount` | `boolean` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_failed` | `boolean` | No |  |
| `is_settled` | `boolean` | No |  |
| `is_submitted` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `Record<string, any>` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `Record<string, any>` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `tag_id` | `any[]` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `number` | No |  |
| `transacted_fx_amount` | `number` | No |  |
| `transacted_fx_rate` | `number` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Field Usage by Operation

| Field | load | create | update | remove |
| --- | --- | --- | --- | --- |
| `account_id` | - | Yes | - | - |
| `allow_incomplete` | - | - | - | - |
| `amount` | - | - | - | - |
| `amount_minor_unit` | - | - | - | - |
| `approve_payout_url` | - | - | - | - |
| `approver_id` | - | - | - | - |
| `authentication_method` | - | - | - | - |
| `authorisation` | - | - | - | - |
| `authorisers_completed_count` | - | - | - | - |
| `authorisers_required_count` | - | - | - | - |
| `batch_payout_id` | - | - | - | - |
| `beneficiary` | - | - | - | - |
| `beneficiary_id` | - | - | - | - |
| `can_authorise` | - | - | - | - |
| `can_process` | - | - | - | - |
| `can_update` | - | - | - | - |
| `charge_bearer` | - | - | - | - |
| `created_by` | - | - | - | - |
| `created_by_email_address` | - | - | - | - |
| `currency` | - | Yes | - | - |
| `current_user_id` | - | - | - | - |
| `description` | - | - | - | - |
| `destination` | - | - | - | - |
| `document` | - | - | - | - |
| `event` | - | - | - | - |
| `formatted_amount` | - | - | - | - |
| `formatted_fx_destination_amount` | - | - | - | - |
| `formatted_schedule` | - | - | - | - |
| `formatted_schedule_day_only` | - | - | - | - |
| `formatted_source_account_available_balance` | - | - | - | - |
| `fx_destination_amount` | - | - | - | - |
| `fx_destination_amount_minor_unit` | - | - | - | - |
| `fx_destination_currency` | - | - | - | - |
| `fx_quote_expires_at` | - | - | - | - |
| `fx_quote_id` | - | - | - | - |
| `fx_rate` | - | - | - | - |
| `fx_use_destination_amount` | - | - | - | - |
| `has_current_user_authorised` | - | - | - | - |
| `id` | - | - | - | - |
| `inserted` | - | - | - | - |
| `invoice_id` | - | - | - | - |
| `is_archived` | - | - | - | - |
| `is_failed` | - | - | - | - |
| `is_settled` | - | - | - | - |
| `is_submitted` | - | - | - | - |
| `last_updated` | - | - | - | - |
| `merchant_id` | - | - | - | - |
| `merchant_token_description` | - | - | - | - |
| `nonce` | - | - | - | - |
| `payment_processor` | - | - | - | - |
| `payment_rail` | - | - | - | - |
| `payrun_id` | - | - | - | - |
| `payrun_name` | - | - | - | - |
| `rule` | - | - | - | - |
| `schedule_date` | - | - | - | - |
| `scheduled` | - | - | - | - |
| `source_account_available_balance` | - | - | - | - |
| `source_account_available_balance_minor_unit` | - | - | - | - |
| `source_account_bic` | - | - | - | - |
| `source_account_currency` | - | - | - | - |
| `source_account_iban` | - | - | - | - |
| `source_account_identifier` | - | - | - | - |
| `source_account_name` | - | - | - | - |
| `source_account_number` | - | - | - | - |
| `source_account_sortcode` | - | - | - | - |
| `status` | - | - | - | - |
| `tag` | - | - | - | - |
| `tag_id` | - | - | - | - |
| `their_reference` | - | - | - | - |
| `topup_payrun_id` | - | - | - | - |
| `transacted_amount` | - | - | - | - |
| `transacted_fx_amount` | - | - | - | - |
| `transacted_fx_rate` | - | - | - | - |
| `type` | - | Yes | - | - |
| `user_id` | - | - | - | - |
| `your_reference` | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Payout().create({
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Payout().load({ id: 'payout_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Payout().remove({ id: 'payout_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Payout().update({
  id: 'payout_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PayoutEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PayrunEntity

```ts
const payrun = client.Payrun()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |
| `note` | `string` | No |  |
| `scheduled_date` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Payrun().create({
  id: /* string */,
})
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Payrun().remove({ id: 'id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Payrun().update({
  id: 'id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PayrunEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## RejectEntity

```ts
const reject = client.Reject()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `Record<string, any>` | Yes |  |
| `can_authorise` | `boolean` | No |  |
| `can_process` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `Record<string, any>` | No |  |
| `document` | `any[]` | No |  |
| `event` | `any[]` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `number` | No |  |
| `fx_destination_amount_minor_unit` | `number` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `number` | No |  |
| `fx_use_destination_amount` | `boolean` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_failed` | `boolean` | No |  |
| `is_settled` | `boolean` | No |  |
| `is_submitted` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `reason` | `string` | No |  |
| `rule` | `Record<string, any>` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `Record<string, any>` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `number` | No |  |
| `transacted_fx_amount` | `number` | No |  |
| `transacted_fx_rate` | `number` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Reject().update({
  id: 'id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `RejectEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ReportEntity

```ts
const report = client.Report()
```

### Operations

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Report().update({
  id: 'id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ReportEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## RuleEntity

```ts
const rule = client.Rule()
```

### Operations

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Rule().remove({ id: 'id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Rule().update({
  id: 'id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `RuleEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SendEntity

```ts
const send = client.Send()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `allow_incomplete` | `boolean` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `Record<string, any>` | Yes |  |
| `beneficiary_id` | `string` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_process` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `Record<string, any>` | No |  |
| `document` | `any[]` | No |  |
| `event` | `any[]` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `number` | No |  |
| `fx_destination_amount_minor_unit` | `number` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `number` | No |  |
| `fx_use_destination_amount` | `boolean` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_failed` | `boolean` | No |  |
| `is_settled` | `boolean` | No |  |
| `is_submitted` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `Record<string, any>` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `Record<string, any>` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `tag_id` | `any[]` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `number` | No |  |
| `transacted_fx_amount` | `number` | No |  |
| `transacted_fx_rate` | `number` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Field Usage by Operation

| Field | create |
| --- | --- |
| `account_id` | Yes |
| `allow_incomplete` | - |
| `amount` | - |
| `amount_minor_unit` | - |
| `approve_payout_url` | - |
| `approver_id` | - |
| `authentication_method` | - |
| `authorisation` | - |
| `authorisers_completed_count` | - |
| `authorisers_required_count` | - |
| `batch_payout_id` | - |
| `beneficiary` | - |
| `beneficiary_id` | - |
| `can_authorise` | - |
| `can_process` | - |
| `can_update` | - |
| `charge_bearer` | - |
| `created_by` | - |
| `created_by_email_address` | - |
| `currency` | Yes |
| `current_user_id` | - |
| `description` | - |
| `destination` | - |
| `document` | - |
| `event` | - |
| `formatted_amount` | - |
| `formatted_fx_destination_amount` | - |
| `formatted_schedule` | - |
| `formatted_schedule_day_only` | - |
| `formatted_source_account_available_balance` | - |
| `fx_destination_amount` | - |
| `fx_destination_amount_minor_unit` | - |
| `fx_destination_currency` | - |
| `fx_quote_expires_at` | - |
| `fx_quote_id` | - |
| `fx_rate` | - |
| `fx_use_destination_amount` | - |
| `has_current_user_authorised` | - |
| `id` | - |
| `inserted` | - |
| `invoice_id` | - |
| `is_archived` | - |
| `is_failed` | - |
| `is_settled` | - |
| `is_submitted` | - |
| `last_updated` | - |
| `merchant_id` | - |
| `merchant_token_description` | - |
| `nonce` | - |
| `payment_processor` | - |
| `payment_rail` | - |
| `payrun_id` | - |
| `payrun_name` | - |
| `rule` | - |
| `schedule_date` | - |
| `scheduled` | - |
| `source_account_available_balance` | - |
| `source_account_available_balance_minor_unit` | - |
| `source_account_bic` | - |
| `source_account_currency` | - |
| `source_account_iban` | - |
| `source_account_identifier` | - |
| `source_account_name` | - |
| `source_account_number` | - |
| `source_account_sortcode` | - |
| `status` | - |
| `tag` | - |
| `tag_id` | - |
| `their_reference` | - |
| `topup_payrun_id` | - |
| `transacted_amount` | - |
| `transacted_fx_amount` | - |
| `transacted_fx_rate` | - |
| `type` | Yes |
| `user_id` | - |
| `your_reference` | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Send().create({
  beneficiary: /* Record<string, any> */,
  source_account_identifier: /* Record<string, any> */,
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SendEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SendbeneficiaryEntity

```ts
const sendbeneficiary = client.Sendbeneficiary()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `allow_incomplete` | `boolean` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `Record<string, any>` | Yes |  |
| `beneficiary_id` | `string` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_process` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `Record<string, any>` | No |  |
| `document` | `any[]` | No |  |
| `event` | `any[]` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `number` | No |  |
| `fx_destination_amount_minor_unit` | `number` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `number` | No |  |
| `fx_use_destination_amount` | `boolean` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_failed` | `boolean` | No |  |
| `is_settled` | `boolean` | No |  |
| `is_submitted` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `Record<string, any>` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `Record<string, any>` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `tag_id` | `any[]` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `number` | No |  |
| `transacted_fx_amount` | `number` | No |  |
| `transacted_fx_rate` | `number` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Field Usage by Operation

| Field | create |
| --- | --- |
| `account_id` | Yes |
| `allow_incomplete` | - |
| `amount` | - |
| `amount_minor_unit` | - |
| `approve_payout_url` | - |
| `approver_id` | - |
| `authentication_method` | - |
| `authorisation` | - |
| `authorisers_completed_count` | - |
| `authorisers_required_count` | - |
| `batch_payout_id` | - |
| `beneficiary` | - |
| `beneficiary_id` | - |
| `can_authorise` | - |
| `can_process` | - |
| `can_update` | - |
| `charge_bearer` | - |
| `created_by` | - |
| `created_by_email_address` | - |
| `currency` | Yes |
| `current_user_id` | - |
| `description` | - |
| `destination` | - |
| `document` | - |
| `event` | - |
| `formatted_amount` | - |
| `formatted_fx_destination_amount` | - |
| `formatted_schedule` | - |
| `formatted_schedule_day_only` | - |
| `formatted_source_account_available_balance` | - |
| `fx_destination_amount` | - |
| `fx_destination_amount_minor_unit` | - |
| `fx_destination_currency` | - |
| `fx_quote_expires_at` | - |
| `fx_quote_id` | - |
| `fx_rate` | - |
| `fx_use_destination_amount` | - |
| `has_current_user_authorised` | - |
| `id` | - |
| `inserted` | - |
| `invoice_id` | - |
| `is_archived` | - |
| `is_failed` | - |
| `is_settled` | - |
| `is_submitted` | - |
| `last_updated` | - |
| `merchant_id` | - |
| `merchant_token_description` | - |
| `nonce` | - |
| `payment_processor` | - |
| `payment_rail` | - |
| `payrun_id` | - |
| `payrun_name` | - |
| `rule` | - |
| `schedule_date` | - |
| `scheduled` | - |
| `source_account_available_balance` | - |
| `source_account_available_balance_minor_unit` | - |
| `source_account_bic` | - |
| `source_account_currency` | - |
| `source_account_iban` | - |
| `source_account_identifier` | - |
| `source_account_name` | - |
| `source_account_number` | - |
| `source_account_sortcode` | - |
| `status` | - |
| `tag` | - |
| `tag_id` | - |
| `their_reference` | - |
| `topup_payrun_id` | - |
| `transacted_amount` | - |
| `transacted_fx_amount` | - |
| `transacted_fx_rate` | - |
| `type` | Yes |
| `user_id` | - |
| `your_reference` | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Sendbeneficiary().create({
  beneficiary: /* Record<string, any> */,
  source_account_identifier: /* Record<string, any> */,
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SendbeneficiaryEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TagEntity

```ts
const tag = client.Tag()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `colour_hex` | `string` | No |  |
| `description` | `string` | No |  |
| `id` | `string` | No |  |
| `merchant_id` | `string` | Yes |  |
| `name` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Tag().create({
  merchant_id: /* string */,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Tag().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TagEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TokenEntity

```ts
const token = client.Token()
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Token().create({
  id: /* string */,
})
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Token().remove({ id: 'id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TokenEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TransactionEntity

```ts
const transaction = client.Transaction()
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Transaction().create({
  id: /* string */,
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Transaction().load({ sequence_number: 1, transaction_id: 'transaction_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Transaction().remove({ id: 'id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TransactionEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## UserEntity

```ts
const user = client.User()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `any[]` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `boolean` | No |  |
| `permission` | `Record<string, any>` | No |  |
| `profile` | `string` | No |  |
| `roles_with_scope` | `any[]` | No |  |
| `two_factor_enabled` | `boolean` | No |  |
| `user_invite_id` | `string` | No |  |

### Field Usage by Operation

| Field | list | update |
| --- | --- | --- |
| `client_session_timeout` | - | - |
| `email_address` | - | Yes |
| `first_name` | - | Yes |
| `id` | - | - |
| `last_name` | - | Yes |
| `passkey_added` | - | - |
| `permission` | - | - |
| `profile` | - | - |
| `roles_with_scope` | - | - |
| `two_factor_enabled` | - | - |
| `user_invite_id` | - | - |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.User().list()
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.User().update({
  id: 'id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `UserEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## UserInviteEntity

```ts
const user_invite = client.UserInvite()
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.UserInvite().create({
  id: /* string */,
})
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.UserInvite().remove({ id: 'id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.UserInvite().update({
  id: 'id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `UserInviteEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## VirtualEntity

```ts
const virtual = client.Virtual()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_name` | `string` | No |  |
| `account_supplier_name` | `string` | No |  |
| `available_balance` | `number` | No |  |
| `available_balance_minor_unit` | `number` | No |  |
| `balance` | `number` | No |  |
| `balance_minor_unit` | `number` | No |  |
| `bank_name` | `string` | No |  |
| `consent_id` | `string` | No |  |
| `created_by` | `Record<string, any>` | Yes |  |
| `created_by_display_name` | `string` | No |  |
| `currency` | `string` | No |  |
| `default_payment_rail` | `string` | No |  |
| `display_name` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `external_account_icon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `Record<string, any>` | Yes |  |
| `inserted` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_connected_account` | `boolean` | No |  |
| `is_default` | `boolean` | No |  |
| `is_trust_account` | `boolean` | No |  |
| `is_virtual` | `boolean` | No |  |
| `last_transaction` | `Record<string, any>` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `name` | `string` | Yes |  |
| `physical_account_id` | `string` | No |  |
| `rule` | `any[]` | No |  |
| `submitted_payouts_balance` | `number` | No |  |
| `submitted_payouts_balance_minor_unit` | `number` | No |  |
| `summary` | `string` | No |  |
| `supplier_sepa_instant_status` | `string` | No |  |
| `xero_bank_feed_connection_status` | `string` | No |  |
| `xero_bank_feed_last_synced_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` | No |  |
| `xero_bank_feed_sync_status` | `string` | No |  |
| `xero_unsynchronised_transactions_count` | `number` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Virtual().create({
  account_id: /* string */,
})
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Virtual().update({
  account_id: 'account_id',
  id: 'id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `VirtualEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## WebhookEntity

```ts
const webhook = client.Webhook()
```

### Operations

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Webhook().remove({ id: 'id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `WebhookEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## WhoamiEntity

```ts
const whoami = client.Whoami()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `any[]` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `boolean` | No |  |
| `permission` | `Record<string, any>` | No |  |
| `roles_with_scope` | `any[]` | No |  |
| `two_factor_enabled` | `boolean` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Whoami().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `WhoamiEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## WhoamitrustedappEntity

```ts
const whoamitrustedapp = client.Whoamitrustedapp()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `any[]` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `boolean` | No |  |
| `permission` | `Record<string, any>` | No |  |
| `roles_with_scope` | `any[]` | No |  |
| `two_factor_enabled` | `boolean` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Whoamitrustedapp().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `WhoamitrustedappEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new NofrixionSDK({
  feature: {
    test: { active: true },
  }
})
```

