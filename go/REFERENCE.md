# Nofrixion Golang SDK Reference

Complete API reference for the Nofrixion Golang SDK.


## NofrixionSDK

### Constructor

```go
func NewNofrixionSDK(options map[string]any) *NofrixionSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *NofrixionSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *NofrixionSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `Account(data map[string]any) NofrixionEntity`

Create a new `Account` entity instance. Pass `nil` for no initial data.

#### `Beneficiary(data map[string]any) NofrixionEntity`

Create a new `Beneficiary` entity instance. Pass `nil` for no initial data.

#### `Cancel(data map[string]any) NofrixionEntity`

Create a new `Cancel` entity instance. Pass `nil` for no initial data.

#### `Disable(data map[string]any) NofrixionEntity`

Create a new `Disable` entity instance. Pass `nil` for no initial data.

#### `Enable(data map[string]any) NofrixionEntity`

Create a new `Enable` entity instance. Pass `nil` for no initial data.

#### `Merchant(data map[string]any) NofrixionEntity`

Create a new `Merchant` entity instance. Pass `nil` for no initial data.

#### `Metadata(data map[string]any) NofrixionEntity`

Create a new `Metadata` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingPaymentRequestPage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionBizBizModelsPagingPaymentRequestPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingPayoutPage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionBizBizModelsPagingPayoutPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingPayrunPage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionBizBizModelsPagingPayrunPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingRuleEventsPage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionBizBizModelsPagingRuleEventsPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingRulesPage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionBizBizModelsPagingRulesPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPaymentsCardPayment(data map[string]any) NofrixionEntity`

Create a new `NoFrixionBizBizModelsPaymentsCardPayment` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPaymentsCardPublicKey(data map[string]any) NofrixionEntity`

Create a new `NoFrixionBizBizModelsPaymentsCardPublicKey` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovApiFeaturesUserInvitesCreate` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsBatchPayout(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsBatchPayout` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsBeneficiaryGroupPage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryGroupPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsBeneficiaryPage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsCardCustomerToken(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsCardCustomerToken` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsCurrencyCurrencyInfo` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsDirectDebitBatchSubmit` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsFxRate(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsFxRate` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsIPayment(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsIPayment` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMandatesMandate(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsMandatesMandate` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchant(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsMerchant` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantPage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsMerchantPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantPayByBankSetting(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsMerchantPayByBankSetting` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantToken(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsMerchantToken` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantTokenPage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsNoFrixionVersion(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsNoFrixionVersion` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsOpenBankingAccount(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsOpenBankingAccount` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsOpenBankingConsent(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsOpenBankingConsent` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsOpenBankingTransaction(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsOpenBankingTransaction` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPayment(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPayment` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountMinimalPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentAccountPage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentInitiation(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentInitiation` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestEvent(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestEvent` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestMetric(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMetric` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestMinimal(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMinimal` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestResult(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestResult` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPayoutKeysetPage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPayoutKeysetPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPayoutMetric(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPayoutMetric` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPayoutsPayoutsCreate` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPayrun(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsPayrun` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsReportResult(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsReportResult` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsRule(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsRule` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsTransaction(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsTransaction` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsTransactionPage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsTransactionPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsUserInvite(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsUserInvite` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsUserInvitePage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsUserInvitePage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsUserPage(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsUserPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsWebhook(data map[string]any) NofrixionEntity`

Create a new `NoFrixionMoneyMoovModelsWebhook` entity instance. Pass `nil` for no initial data.

#### `OpenBanking(data map[string]any) NofrixionEntity`

Create a new `OpenBanking` entity instance. Pass `nil` for no initial data.

#### `Payeeverification(data map[string]any) NofrixionEntity`

Create a new `Payeeverification` entity instance. Pass `nil` for no initial data.

#### `PaymentRequest(data map[string]any) NofrixionEntity`

Create a new `PaymentRequest` entity instance. Pass `nil` for no initial data.

#### `Payout(data map[string]any) NofrixionEntity`

Create a new `Payout` entity instance. Pass `nil` for no initial data.

#### `Payrun(data map[string]any) NofrixionEntity`

Create a new `Payrun` entity instance. Pass `nil` for no initial data.

#### `Reject(data map[string]any) NofrixionEntity`

Create a new `Reject` entity instance. Pass `nil` for no initial data.

#### `Report(data map[string]any) NofrixionEntity`

Create a new `Report` entity instance. Pass `nil` for no initial data.

#### `Rule(data map[string]any) NofrixionEntity`

Create a new `Rule` entity instance. Pass `nil` for no initial data.

#### `Send(data map[string]any) NofrixionEntity`

Create a new `Send` entity instance. Pass `nil` for no initial data.

#### `Sendbeneficiary(data map[string]any) NofrixionEntity`

Create a new `Sendbeneficiary` entity instance. Pass `nil` for no initial data.

#### `Tag(data map[string]any) NofrixionEntity`

Create a new `Tag` entity instance. Pass `nil` for no initial data.

#### `Token(data map[string]any) NofrixionEntity`

Create a new `Token` entity instance. Pass `nil` for no initial data.

#### `Transaction(data map[string]any) NofrixionEntity`

Create a new `Transaction` entity instance. Pass `nil` for no initial data.

#### `User(data map[string]any) NofrixionEntity`

Create a new `User` entity instance. Pass `nil` for no initial data.

#### `UserInvite(data map[string]any) NofrixionEntity`

Create a new `UserInvite` entity instance. Pass `nil` for no initial data.

#### `Virtual(data map[string]any) NofrixionEntity`

Create a new `Virtual` entity instance. Pass `nil` for no initial data.

#### `Webhook(data map[string]any) NofrixionEntity`

Create a new `Webhook` entity instance. Pass `nil` for no initial data.

#### `Whoami(data map[string]any) NofrixionEntity`

Create a new `Whoami` entity instance. Pass `nil` for no initial data.

#### `Whoamitrustedapp(data map[string]any) NofrixionEntity`

Create a new `Whoamitrustedapp` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## AccountEntity

```go
account := client.Account(nil)
fmt.Println(account.GetName()) // "account"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `account_name` | `string` | No |  |
| `account_supplier_name` | `string` | No |  |
| `account_type` | `string` | No |  |
| `available_balance` | `float64` | No |  |
| `available_balance_minor_unit` | `int` | No |  |
| `balance` | `float64` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `bank_name` | `string` | No |  |
| `consent_id` | `string` | No |  |
| `created_by` | `map[string]any` | Yes |  |
| `created_by_display_name` | `string` | No |  |
| `currency` | `string` | No |  |
| `default_payment_rail` | `string` | No |  |
| `display_name` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `external_account_icon` | `string` | No |  |
| `format` | `string` | No |  |
| `from_date` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `map[string]any` | Yes |  |
| `inserted` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_connected_account` | `bool` | No |  |
| `is_default` | `bool` | No |  |
| `is_trust_account` | `bool` | No |  |
| `is_virtual` | `bool` | No |  |
| `last_transaction` | `map[string]any` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `physical_account_id` | `string` | No |  |
| `role_i_d` | `[]any` | No |  |
| `rule` | `[]any` | No |  |
| `submitted_payouts_balance` | `float64` | No |  |
| `submitted_payouts_balance_minor_unit` | `int` | No |  |
| `summary` | `string` | No |  |
| `supplier_physical_account_id` | `string` | No |  |
| `supplier_sepa_instant_status` | `string` | No |  |
| `to_date` | `string` | No |  |
| `xero_bank_feed_connection_status` | `string` | No |  |
| `xero_bank_feed_last_synced_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` | No |  |
| `xero_bank_feed_sync_status` | `string` | No |  |
| `xero_unsynchronised_transactions_count` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Account(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Account(nil).Load(map[string]any{"id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Account(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Account(nil).Update(map[string]any{
    "id": "account_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Account(nil).Remove(map[string]any{"id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AccountEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## BeneficiaryEntity

```go
beneficiary := client.Beneficiary(nil)
fmt.Println(beneficiary.GetName()) // "beneficiary"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `[]any` | No |  |
| `authorisation` | `[]any` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `beneficiary_event` | `[]any` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `created_by` | `map[string]any` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `map[string]any` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `bool` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `[]any` | No |  |
| `source_account_i_d` | `[]any` | No |  |
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Beneficiary(nil).Load(map[string]any{"id": "beneficiary_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Beneficiary(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Beneficiary(nil).Update(map[string]any{
    "id": "beneficiary_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Beneficiary(nil).Remove(map[string]any{"id": "beneficiary_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `BeneficiaryEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CancelEntity

```go
cancel := client.Cancel(nil)
fmt.Println(cancel.GetName()) // "cancel"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `float64` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `[]any` | No |  |
| `authorisation` | `[]any` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `map[string]any` | Yes |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `map[string]any` | No |  |
| `document` | `[]any` | No |  |
| `event` | `[]any` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `float64` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `float64` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `map[string]any` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float64` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `map[string]any` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `[]any` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `float64` | No |  |
| `transacted_fx_amount` | `float64` | No |  |
| `transacted_fx_rate` | `float64` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Cancel(nil).Update(map[string]any{
    "id": "id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CancelEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## DisableEntity

```go
disable := client.Disable(nil)
fmt.Println(disable.GetName()) // "disable"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `[]any` | No |  |
| `authorisation` | `[]any` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `beneficiary_event` | `[]any` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `created_by` | `map[string]any` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `map[string]any` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `bool` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `[]any` | No |  |
| `their_reference` | `string` | No |  |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Disable(nil).Update(map[string]any{
    "id": "id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `DisableEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## EnableEntity

```go
enable := client.Enable(nil)
fmt.Println(enable.GetName()) // "enable"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `[]any` | No |  |
| `authorisation` | `[]any` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `beneficiary_event` | `[]any` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `created_by` | `map[string]any` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `map[string]any` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `bool` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `[]any` | No |  |
| `their_reference` | `string` | No |  |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Enable(nil).Update(map[string]any{
    "id": "id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `EnableEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MerchantEntity

```go
merchant := client.Merchant(nil)
fmt.Println(merchant.GetName()) // "merchant"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `reason` | `string` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Merchant(nil).Load(map[string]any{"merchant_id": "merchant_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Merchant(nil).Update(map[string]any{
    "id": "id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Merchant(nil).Remove(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MerchantEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MetadataEntity

```go
metadata := client.Metadata(nil)
fmt.Println(metadata.GetName()) // "metadata"
```

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Metadata(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MetadataEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity

```go
noFrixionBizBizModelsPagingMerchantDirectDebitMandatePage := client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(nil)
fmt.Println(noFrixionBizBizModelsPagingMerchantDirectDebitMandatePage.GetName()) // "no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page"
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
| `is_recurring` | `bool` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingPaymentRequestPageEntity

```go
noFrixionBizBizModelsPagingPaymentRequestPage := client.NoFrixionBizBizModelsPagingPaymentRequestPage(nil)
fmt.Println(noFrixionBizBizModelsPagingPaymentRequestPage.GetName()) // "no_frixion_biz_biz_models_paging_payment_request_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `[]any` | No |  |
| `amount` | `float64` | No |  |
| `amount_pending` | `float64` | No |  |
| `amount_received` | `float64` | No |  |
| `amount_refunded` | `float64` | No |  |
| `auto_send_receipt` | `bool` | No |  |
| `base_origin_url` | `string` | No |  |
| `callback_url` | `string` | No |  |
| `card_authorize_only` | `bool` | No |  |
| `card_create_token` | `bool` | No |  |
| `card_create_token_mode` | `string` | No |  |
| `card_ignore_cvn` | `bool` | No |  |
| `card_processor_merchant_id` | `string` | No |  |
| `card_stripe_payment_intent_id` | `string` | No |  |
| `card_stripe_payment_intent_secret` | `string` | No |  |
| `created_by_user` | `map[string]any` | Yes |  |
| `currency` | `string` | No |  |
| `custom_field` | `[]any` | No |  |
| `customer_email_address` | `string` | No |  |
| `customer_id` | `string` | No |  |
| `customer_name` | `string` | No |  |
| `description` | `string` | No |  |
| `destination_account` | `map[string]any` | No |  |
| `direct_debit_payment` | `map[string]any` | No |  |
| `due_date` | `string` | No |  |
| `event` | `[]any` | No |  |
| `failure_callback_url` | `string` | No |  |
| `field_display_setting` | `[]any` | No |  |
| `formatted_amount` | `string` | No |  |
| `hosted_pay_checkout_url` | `string` | No |  |
| `id` | `string` | No |  |
| `ignore_address_verification` | `bool` | No |  |
| `inserted` | `string` | No |  |
| `inserted_sortable` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `jwk` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `lightning_invoice` | `string` | No |  |
| `lightning_invoice_expires_at` | `string` | No |  |
| `merchant_direct_debit_mandate_id` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `notification_email_address` | `string` | No |  |
| `notification_role_i_d` | `[]any` | No |  |
| `order_id` | `string` | No |  |
| `partial_payment_method` | `string` | No |  |
| `partial_payment_step` | `string` | No |  |
| `payment_attempt` | `[]any` | No |  |
| `payment_method` | `[]any` | No |  |
| `payment_processor` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `pisp_account_id` | `string` | No |  |
| `priority_bank_id` | `string` | No |  |
| `result` | `map[string]any` | No |  |
| `sandbox_settle_delay_in_second` | `int` | No |  |
| `shipping_address` | `map[string]any` | No |  |
| `status` | `string` | No |  |
| `success_web_hook_url` | `string` | No |  |
| `tag` | `[]any` | No |  |
| `title` | `string` | No |  |
| `tokenised_card` | `[]any` | No |  |
| `transaction` | `[]any` | No |  |
| `use_hosted_payment_page` | `bool` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionBizBizModelsPagingPaymentRequestPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionBizBizModelsPagingPaymentRequestPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingPayoutPageEntity

```go
noFrixionBizBizModelsPagingPayoutPage := client.NoFrixionBizBizModelsPagingPayoutPage(nil)
fmt.Println(noFrixionBizBizModelsPagingPayoutPage.GetName()) // "no_frixion_biz_biz_models_paging_payout_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `float64` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `[]any` | No |  |
| `authorisation` | `[]any` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `map[string]any` | Yes |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `map[string]any` | No |  |
| `document` | `[]any` | No |  |
| `event` | `[]any` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `float64` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `float64` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `map[string]any` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float64` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `map[string]any` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `[]any` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `float64` | No |  |
| `transacted_fx_amount` | `float64` | No |  |
| `transacted_fx_rate` | `float64` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionBizBizModelsPagingPayoutPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionBizBizModelsPagingPayoutPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingPayrunPageEntity

```go
noFrixionBizBizModelsPagingPayrunPage := client.NoFrixionBizBizModelsPagingPayrunPage(nil)
fmt.Println(noFrixionBizBizModelsPagingPayrunPage.GetName()) // "no_frixion_biz_biz_models_paging_payrun_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation` | `[]any` | No |  |
| `authorisation_date` | `string` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_delete` | `bool` | No |  |
| `can_edit` | `bool` | No |  |
| `event` | `[]any` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice` | `[]any` | No |  |
| `invoices_minimal` | `[]any` | No |  |
| `is_archived` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `last_updated_by` | `map[string]any` | Yes |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment` | `[]any` | No |  |
| `payout` | `[]any` | No |  |
| `payouts_count` | `int` | No |  |
| `schedule_date` | `string` | No |  |
| `source_account` | `[]any` | No |  |
| `status` | `string` | No |  |
| `total_eur` | `float64` | No |  |
| `total_gbp` | `float64` | No |  |
| `total_usd` | `float64` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionBizBizModelsPagingPayrunPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionBizBizModelsPagingPayrunPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingRuleEventsPageEntity

```go
noFrixionBizBizModelsPagingRuleEventsPage := client.NoFrixionBizBizModelsPagingRuleEventsPage(nil)
fmt.Println(noFrixionBizBizModelsPagingRuleEventsPage.GetName()) // "no_frixion_biz_biz_models_paging_rule_events_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_message` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_authorise_to_enable` | `bool` | No |  |
| `message` | `string` | No |  |
| `raw_response` | `string` | No |  |
| `rule_event_type` | `string` | No |  |
| `rule_id` | `string` | No |  |
| `user` | `map[string]any` | Yes |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionBizBizModelsPagingRuleEventsPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionBizBizModelsPagingRuleEventsPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingRulesPageEntity

```go
noFrixionBizBizModelsPagingRulesPage := client.NoFrixionBizBizModelsPagingRulesPage(nil)
fmt.Println(noFrixionBizBizModelsPagingRulesPage.GetName()) // "no_frixion_biz_biz_models_paging_rules_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account` | `map[string]any` | No |  |
| `account_id` | `string` | No |  |
| `approve_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `[]any` | No |  |
| `authorisation` | `[]any` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `can_authorise` | `bool` | No |  |
| `created_by` | `map[string]any` | Yes |  |
| `description` | `string` | No |  |
| `end_at` | `string` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_disabled` | `bool` | No |  |
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
| `sweep_action` | `map[string]any` | No |  |
| `time_zone_id` | `string` | No |  |
| `trigger_cron_expression` | `string` | No |  |
| `trigger_on_pay_in` | `bool` | No |  |
| `user_id` | `string` | No |  |
| `web_hook_secret` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionBizBizModelsPagingRulesPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionBizBizModelsPagingRulesPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionBizBizModelsPaymentsCardPaymentEntity

```go
noFrixionBizBizModelsPaymentsCardPayment := client.NoFrixionBizBizModelsPaymentsCardPayment(nil)
fmt.Println(noFrixionBizBizModelsPaymentsCardPayment.GetName()) // "no_frixion_biz_biz_models_payments_card_payment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorized_amount` | `string` | No |  |
| `currency_code` | `string` | No |  |
| `is_payer_authentication_required` | `bool` | No |  |
| `is_soft_decline` | `bool` | No |  |
| `payer_authentication_access_token` | `string` | No |  |
| `payer_authentication_merchant_data` | `string` | No |  |
| `payer_authentication_url` | `string` | No |  |
| `payer_authentication_window_height` | `int` | No |  |
| `payer_authentication_window_width` | `int` | No |  |
| `payment_request_callback_url` | `string` | No |  |
| `payment_request_id` | `string` | No |  |
| `request_id` | `string` | No |  |
| `response_code` | `string` | No |  |
| `response_type` | `string` | No |  |
| `status` | `string` | No |  |
| `three_ds_redirect_url` | `string` | No |  |
| `transaction_id` | `string` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionBizBizModelsPaymentsCardPayment(nil).Create(map[string]any{
    "paymentrequest_id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionBizBizModelsPaymentsCardPaymentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionBizBizModelsPaymentsCardPublicKeyEntity

```go
noFrixionBizBizModelsPaymentsCardPublicKey := client.NoFrixionBizBizModelsPaymentsCardPublicKey(nil)
fmt.Println(noFrixionBizBizModelsPaymentsCardPublicKey.GetName()) // "no_frixion_biz_biz_models_payments_card_public_key"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jwt` | `string` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionBizBizModelsPaymentsCardPublicKey(nil).Load(map[string]any{"paymentrequest_id": "paymentrequest_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionBizBizModelsPaymentsCardPublicKeyEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity

```go
noFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries := client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(nil)
fmt.Println(noFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries.GetName()) // "no_frixion_money_moov_api_features_beneficiaries_beneficiaries"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `beneficiary` | `[]any` | No |  |
| `failed_beneficiary` | `map[string]any` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity

```go
noFrixionMoneyMoovApiFeaturesPaymentRequestsPayment := client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(nil)
fmt.Println(noFrixionMoneyMoovApiFeaturesPaymentRequestsPayment.GetName()) // "no_frixion_money_moov_api_features_payment_requests_payment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_payment_request` | `map[string]any` | No |  |
| `payment_request` | `[]any` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity

```go
noFrixionMoneyMoovApiFeaturesPermissionsRolesCreate := client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(nil)
fmt.Println(noFrixionMoneyMoovApiFeaturesPermissionsRolesCreate.GetName()) // "no_frixion_money_moov_api_features_permissions_roles_create"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_role` | `map[string]any` | No |  |
| `role` | `[]any` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(nil).Create(map[string]any{
    "merchant_id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity

```go
noFrixionMoneyMoovApiFeaturesUserInvitesCreate := client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(nil)
fmt.Println(noFrixionMoneyMoovApiFeaturesUserInvitesCreate.GetName()) // "no_frixion_money_moov_api_features_user_invites_create"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_user_invite` | `map[string]any` | No |  |
| `user_invite` | `[]any` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity

```go
noFrixionMoneyMoovModelsAuthorisationSettingsMerchant := client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(nil)
fmt.Println(noFrixionMoneyMoovModelsAuthorisationSettingsMerchant.GetName()) // "no_frixion_money_moov_models_authorisation_settings_merchant"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount_lower` | `float64` | No |  |
| `amount_upper` | `float64` | No |  |
| `authorisation_type` | `string` | No |  |
| `beneficiaries_only` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `last_editor_cant_authorise` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `number_of_authoriser` | `int` | No |  |
| `role_setting` | `[]any` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsBatchPayoutEntity

```go
noFrixionMoneyMoovModelsBatchPayout := client.NoFrixionMoneyMoovModelsBatchPayout(nil)
fmt.Println(noFrixionMoneyMoovModelsBatchPayout.GetName()) // "no_frixion_money_moov_models_batch_payout"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approve_url` | `string` | No |  |
| `id` | `string` | No |  |
| `payout` | `[]any` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsBatchPayout(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_batch_payout_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovModelsBatchPayout(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsBatchPayoutEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity

```go
noFrixionMoneyMoovModelsBeneficiaryGroupPage := client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage(nil)
fmt.Println(noFrixionMoneyMoovModelsBeneficiaryGroupPage.GetName()) // "no_frixion_money_moov_models_beneficiary_group_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `group_member` | `[]any` | No |  |
| `group_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | Yes |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsBeneficiaryPageEntity

```go
noFrixionMoneyMoovModelsBeneficiaryPage := client.NoFrixionMoneyMoovModelsBeneficiaryPage(nil)
fmt.Println(noFrixionMoneyMoovModelsBeneficiaryPage.GetName()) // "no_frixion_money_moov_models_beneficiary_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `[]any` | No |  |
| `authorisation` | `[]any` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `beneficiary_event` | `[]any` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `created_by` | `map[string]any` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `map[string]any` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `bool` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `[]any` | No |  |
| `their_reference` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsBeneficiaryPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsCardCustomerTokenEntity

```go
noFrixionMoneyMoovModelsCardCustomerToken := client.NoFrixionMoneyMoovModelsCardCustomerToken(nil)
fmt.Println(noFrixionMoneyMoovModelsCardCustomerToken.GetName()) // "no_frixion_money_moov_models_card_customer_token"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsCardCustomerToken(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsCardCustomerToken(nil).Load(map[string]any{"customer_email_address": "customer_email_address"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsCardCustomerToken(nil).Remove(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsCardCustomerTokenEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity

```go
noFrixionMoneyMoovModelsCurrencyCurrencyInfo := client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(nil)
fmt.Println(noFrixionMoneyMoovModelsCurrencyCurrencyInfo.GetName()) // "no_frixion_money_moov_models_currency_currency_info"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `code` | `string` | No |  |
| `decimal` | `int` | No |  |
| `is_fiat` | `bool` | No |  |
| `iso4217_alpha_code` | `string` | No |  |
| `iso4217_numeric_code` | `string` | No |  |
| `symbol` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity

```go
noFrixionMoneyMoovModelsDirectDebitBatchSubmit := client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(nil)
fmt.Println(noFrixionMoneyMoovModelsDirectDebitBatchSubmit.GetName()) // "no_frixion_money_moov_models_direct_debit_batch_submit"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_submission` | `map[string]any` | No |  |
| `successful_submission` | `[]any` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsFxRateEntity

```go
noFrixionMoneyMoovModelsFxRate := client.NoFrixionMoneyMoovModelsFxRate(nil)
fmt.Println(noFrixionMoneyMoovModelsFxRate.GetName()) // "no_frixion_money_moov_models_fx_rate"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destination_currency` | `string` | No |  |
| `exchange_rate` | `float64` | No |  |
| `expiry_time` | `string` | No |  |
| `quote_id` | `string` | No |  |
| `source_currency` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsFxRate(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsFxRate(nil).Load(map[string]any{"destination": "destination", "source": "source", "valid_for_minute": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsFxRateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsIPaymentEntity

```go
noFrixionMoneyMoovModelsIPayment := client.NoFrixionMoneyMoovModelsIPayment(nil)
fmt.Println(noFrixionMoneyMoovModelsIPayment.GetName()) // "no_frixion_money_moov_models_i_payment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `payment_request_id` | `string` | No |  |
| `response_type` | `string` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovModelsIPayment(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsIPaymentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMandatesMandateEntity

```go
noFrixionMoneyMoovModelsMandatesMandate := client.NoFrixionMoneyMoovModelsMandatesMandate(nil)
fmt.Println(noFrixionMoneyMoovModelsMandatesMandate.GetName()) // "no_frixion_money_moov_models_mandates_mandate"
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
| `is_recurring` | `bool` | No |  |
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsMandatesMandate(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_mandates_mandate_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovModelsMandatesMandate(nil).Create(map[string]any{
    "address_line1": /* string */,
    "city": /* string */,
    "country_code": /* string */,
    "email_address": /* string */,
    "first_name": /* string */,
    "last_name": /* string */,
    "postal_code": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsMandatesMandateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantEntity

```go
noFrixionMoneyMoovModelsMerchant := client.NoFrixionMoneyMoovModelsMerchant(nil)
fmt.Println(noFrixionMoneyMoovModelsMerchant.GetName()) // "no_frixion_money_moov_models_merchant"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_currency` | `[]any` | No |  |
| `can_have_trust_account` | `bool` | No |  |
| `card_payment_processor` | `string` | No |  |
| `company_id` | `string` | No |  |
| `display_qr_on_hosted_pay` | `bool` | No |  |
| `hosted_pay_version` | `int` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_blocked` | `bool` | No |  |
| `is_exited` | `bool` | No |  |
| `is_suspended` | `bool` | No |  |
| `jurisdiction` | `string` | No |  |
| `logo_url_png` | `string` | No |  |
| `logo_url_svg` | `string` | No |  |
| `merchant_category_code` | `string` | No |  |
| `name` | `string` | No |  |
| `note` | `string` | No |  |
| `parent_merchant` | `map[string]any` | No |  |
| `payment_account` | `[]any` | No |  |
| `payment_account_limit` | `int` | No |  |
| `short_name` | `string` | No |  |
| `supported_payment_methods_list` | `[]any` | No |  |
| `suspension_reason` | `string` | No |  |
| `tag` | `[]any` | No |  |
| `time_zone_id` | `string` | No |  |
| `trading_name` | `string` | No |  |
| `web_hook_limit` | `int` | No |  |
| `your_role_name` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsMerchant(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsMerchant(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_merchant_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.NoFrixionMoneyMoovModelsMerchant(nil).Update(map[string]any{
    "id": "no_frixion_money_moov_models_merchant_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsMerchantEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantPageEntity

```go
noFrixionMoneyMoovModelsMerchantPage := client.NoFrixionMoneyMoovModelsMerchantPage(nil)
fmt.Println(noFrixionMoneyMoovModelsMerchantPage.GetName()) // "no_frixion_money_moov_models_merchant_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_currency` | `[]any` | No |  |
| `can_have_trust_account` | `bool` | No |  |
| `card_payment_processor` | `string` | No |  |
| `company_id` | `string` | No |  |
| `display_qr_on_hosted_pay` | `bool` | No |  |
| `hosted_pay_version` | `int` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_blocked` | `bool` | No |  |
| `is_exited` | `bool` | No |  |
| `is_suspended` | `bool` | No |  |
| `jurisdiction` | `string` | No |  |
| `logo_url_png` | `string` | No |  |
| `logo_url_svg` | `string` | No |  |
| `merchant_category_code` | `string` | No |  |
| `name` | `string` | No |  |
| `note` | `string` | No |  |
| `parent_merchant` | `map[string]any` | No |  |
| `payment_account` | `[]any` | No |  |
| `payment_account_limit` | `int` | No |  |
| `short_name` | `string` | No |  |
| `supported_payment_methods_list` | `[]any` | No |  |
| `suspension_reason` | `string` | No |  |
| `tag` | `[]any` | No |  |
| `time_zone_id` | `string` | No |  |
| `trading_name` | `string` | No |  |
| `web_hook_limit` | `int` | No |  |
| `your_role_name` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsMerchantPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsMerchantPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity

```go
noFrixionMoneyMoovModelsMerchantPayByBankSetting := client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting(nil)
fmt.Println(noFrixionMoneyMoovModelsMerchantPayByBankSetting.GetName()) // "no_frixion_money_moov_models_merchant_pay_by_bank_setting"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bank_country_code` | `[]any` | No |  |
| `bank_id` | `string` | No |  |
| `bank_name` | `string` | No |  |
| `business_institution_id` | `string` | No |  |
| `currency` | `string` | No |  |
| `logo` | `string` | No |  |
| `message` | `string` | No |  |
| `message_image_url` | `string` | No |  |
| `order` | `int` | No |  |
| `personal_institution_id` | `string` | No |  |
| `processor` | `string` | No |  |
| `warning_heading` | `string` | No |  |
| `warning_message` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantTokenEntity

```go
noFrixionMoneyMoovModelsMerchantToken := client.NoFrixionMoneyMoovModelsMerchantToken(nil)
fmt.Println(noFrixionMoneyMoovModelsMerchantToken.GetName()) // "no_frixion_money_moov_models_merchant_token"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authentication_method` | `[]any` | No |  |
| `authorisation` | `[]any` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `can_authorise` | `bool` | No |  |
| `description` | `string` | No |  |
| `expires_at` | `string` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `hmac_algorithm` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `ip_address_whitelist` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_enabled` | `bool` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `nonce` | `string` | Yes |  |
| `permission_type` | `[]any` | No |  |
| `request_signature_version` | `int` | No |  |
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsMerchantToken(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_merchant_token_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovModelsMerchantToken(nil).Create(map[string]any{
    "nonce": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.NoFrixionMoneyMoovModelsMerchantToken(nil).Update(map[string]any{
    "id": "no_frixion_money_moov_models_merchant_token_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantTokenPageEntity

```go
noFrixionMoneyMoovModelsMerchantTokenPage := client.NoFrixionMoneyMoovModelsMerchantTokenPage(nil)
fmt.Println(noFrixionMoneyMoovModelsMerchantTokenPage.GetName()) // "no_frixion_money_moov_models_merchant_token_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authentication_method` | `[]any` | No |  |
| `authorisation` | `[]any` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `can_authorise` | `bool` | No |  |
| `description` | `string` | No |  |
| `expires_at` | `string` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `ip_address_whitelist` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_enabled` | `bool` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `nonce` | `string` | Yes |  |
| `permission_type` | `[]any` | No |  |
| `request_signature_version` | `int` | No |  |
| `shared_secret_algorithm` | `string` | No |  |
| `shared_secret_base64` | `string` | No |  |
| `token` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsMerchantTokenPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsNoFrixionVersionEntity

```go
noFrixionMoneyMoovModelsNoFrixionVersion := client.NoFrixionMoneyMoovModelsNoFrixionVersion(nil)
fmt.Println(noFrixionMoneyMoovModelsNoFrixionVersion.GetName()) // "no_frixion_money_moov_models_no_frixion_version"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `build_version` | `int` | No |  |
| `major_version` | `int` | No |  |
| `minor_version` | `int` | No |  |
| `release_name` | `string` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsNoFrixionVersion(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsNoFrixionVersionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsOpenBankingAccountEntity

```go
noFrixionMoneyMoovModelsOpenBankingAccount := client.NoFrixionMoneyMoovModelsOpenBankingAccount(nil)
fmt.Println(noFrixionMoneyMoovModelsOpenBankingAccount.GetName()) // "no_frixion_money_moov_models_open_banking_account"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_balance` | `[]any` | No |  |
| `account_identification` | `[]any` | No |  |
| `account_name` | `[]any` | No |  |
| `account_type` | `string` | No |  |
| `balance` | `float64` | No |  |
| `consolidated_account_information` | `map[string]any` | No |  |
| `currency` | `string` | No |  |
| `description` | `string` | No |  |
| `detail` | `string` | No |  |
| `id` | `string` | No |  |
| `nickname` | `string` | No |  |
| `type` | `string` | No |  |
| `usage_type` | `string` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsOpenBankingAccount(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_open_banking_account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsOpenBankingAccountEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsOpenBankingConsentEntity

```go
noFrixionMoneyMoovModelsOpenBankingConsent := client.NoFrixionMoneyMoovModelsOpenBankingConsent(nil)
fmt.Println(noFrixionMoneyMoovModelsOpenBankingConsent.GetName()) // "no_frixion_money_moov_models_open_banking_consent"
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
| `is_connected_account` | `bool` | No |  |
| `is_enabled` | `bool` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsOpenBankingConsent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsOpenBankingConsent(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_open_banking_consent_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovModelsOpenBankingConsent(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.NoFrixionMoneyMoovModelsOpenBankingConsent(nil).Update(map[string]any{
    "id": "no_frixion_money_moov_models_open_banking_consent_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsOpenBankingConsent(nil).Remove(map[string]any{"id": "no_frixion_money_moov_models_open_banking_consent_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsOpenBankingConsentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsOpenBankingTransactionEntity

```go
noFrixionMoneyMoovModelsOpenBankingTransaction := client.NoFrixionMoneyMoovModelsOpenBankingTransaction(nil)
fmt.Println(noFrixionMoneyMoovModelsOpenBankingTransaction.GetName()) // "no_frixion_money_moov_models_open_banking_transaction"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address_detail` | `map[string]any` | No |  |
| `amount` | `float64` | No |  |
| `balance` | `map[string]any` | No |  |
| `booking_date_time` | `string` | No |  |
| `charge_detail` | `map[string]any` | No |  |
| `currency` | `string` | No |  |
| `currency_exchange` | `map[string]any` | No |  |
| `date` | `string` | No |  |
| `description` | `string` | No |  |
| `enrichment` | `map[string]any` | No |  |
| `gross_amount` | `map[string]any` | Yes |  |
| `id` | `string` | No |  |
| `iso_bank_transaction_code` | `map[string]any` | No |  |
| `merchant` | `map[string]any` | No |  |
| `payee_detail` | `map[string]any` | Yes |  |
| `payer_detail` | `map[string]any` | Yes |  |
| `proprietary_bank_transaction_code` | `map[string]any` | No |  |
| `reference` | `string` | No |  |
| `statement_reference` | `[]any` | No |  |
| `status` | `string` | No |  |
| `supplementary_data` | `any` | No |  |
| `transaction_amount` | `map[string]any` | Yes |  |
| `transaction_information` | `[]any` | No |  |
| `transaction_mutability` | `string` | No |  |
| `value_date_time` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsOpenBankingTransaction(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsOpenBankingTransactionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentEntity

```go
noFrixionMoneyMoovModelsPayment := client.NoFrixionMoneyMoovModelsPayment(nil)
fmt.Println(noFrixionMoneyMoovModelsPayment.GetName()) // "no_frixion_money_moov_models_payment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `[]any` | No |  |
| `amount` | `float64` | No |  |
| `amount_pending` | `float64` | No |  |
| `amount_received` | `float64` | No |  |
| `amount_refunded` | `float64` | No |  |
| `auto_send_receipt` | `bool` | No |  |
| `base_origin_url` | `string` | No |  |
| `callback_url` | `string` | No |  |
| `card_authorize_only` | `bool` | No |  |
| `card_create_token` | `bool` | No |  |
| `card_create_token_mode` | `string` | No |  |
| `card_ignore_cvn` | `bool` | No |  |
| `card_no_payer_authentication` | `bool` | No |  |
| `card_processor_merchant_id` | `string` | No |  |
| `card_stripe_payment_intent_id` | `string` | No |  |
| `card_stripe_payment_intent_secret` | `string` | No |  |
| `card_transmit_raw_detail` | `bool` | No |  |
| `created_by_user` | `map[string]any` | Yes |  |
| `currency` | `string` | No |  |
| `custom_field` | `[]any` | No |  |
| `customer_email_address` | `string` | No |  |
| `customer_id` | `string` | No |  |
| `customer_name` | `string` | No |  |
| `description` | `string` | No |  |
| `destination_account` | `map[string]any` | No |  |
| `direct_debit_payment` | `map[string]any` | No |  |
| `due_date` | `string` | No |  |
| `event` | `[]any` | No |  |
| `failure_callback_url` | `string` | No |  |
| `field_display_setting` | `[]any` | No |  |
| `formatted_amount` | `string` | No |  |
| `hosted_pay_checkout_url` | `string` | No |  |
| `id` | `string` | No |  |
| `ignore_address_verification` | `bool` | No |  |
| `inserted` | `string` | No |  |
| `inserted_sortable` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `jwk` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `lightning_invoice` | `string` | No |  |
| `lightning_invoice_expires_at` | `string` | No |  |
| `merchant_direct_debit_mandate_id` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `notification_email_address` | `string` | No |  |
| `notification_role_i_d` | `[]any` | No |  |
| `order_id` | `string` | No |  |
| `partial_payment_method` | `string` | No |  |
| `partial_payment_step` | `string` | No |  |
| `payment_attempt` | `[]any` | No |  |
| `payment_method` | `[]any` | No |  |
| `payment_processor` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `pisp_account_id` | `string` | No |  |
| `priority_bank_id` | `string` | No |  |
| `result` | `map[string]any` | No |  |
| `sandbox_settle_delay_in_second` | `int` | No |  |
| `shipping_address` | `map[string]any` | No |  |
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
| `tag` | `[]any` | No |  |
| `tag_id` | `[]any` | No |  |
| `title` | `string` | No |  |
| `tokenised_card` | `[]any` | No |  |
| `transaction` | `[]any` | No |  |
| `use_hosted_payment_page` | `bool` | No |  |

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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsPayment(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_payment_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovModelsPayment(nil).Create(map[string]any{
    "created_by_user": /* map[string]any */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.NoFrixionMoneyMoovModelsPayment(nil).Update(map[string]any{
    "id": "no_frixion_money_moov_models_payment_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity

```go
noFrixionMoneyMoovModelsPaymentAccountMinimalPage := client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(nil)
fmt.Println(noFrixionMoneyMoovModelsPaymentAccountMinimalPage.GetName()) // "no_frixion_money_moov_models_payment_account_minimal_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_name` | `string` | No |  |
| `available_balance` | `float64` | No |  |
| `balance` | `float64` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `currency` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `map[string]any` | Yes |  |
| `is_archived` | `bool` | No |  |
| `is_connected_account` | `bool` | No |  |
| `merchant_id` | `string` | No |  |
| `submitted_payouts_balance` | `float64` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentAccountPageEntity

```go
noFrixionMoneyMoovModelsPaymentAccountPage := client.NoFrixionMoneyMoovModelsPaymentAccountPage(nil)
fmt.Println(noFrixionMoneyMoovModelsPaymentAccountPage.GetName()) // "no_frixion_money_moov_models_payment_account_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_name` | `string` | No |  |
| `account_supplier_name` | `string` | No |  |
| `available_balance` | `float64` | No |  |
| `available_balance_minor_unit` | `int` | No |  |
| `balance` | `float64` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `bank_name` | `string` | No |  |
| `consent_id` | `string` | No |  |
| `created_by` | `map[string]any` | Yes |  |
| `created_by_display_name` | `string` | No |  |
| `currency` | `string` | No |  |
| `default_payment_rail` | `string` | No |  |
| `display_name` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `external_account_icon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `map[string]any` | Yes |  |
| `inserted` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_connected_account` | `bool` | No |  |
| `is_default` | `bool` | No |  |
| `is_trust_account` | `bool` | No |  |
| `is_virtual` | `bool` | No |  |
| `last_transaction` | `map[string]any` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `physical_account_id` | `string` | No |  |
| `rule` | `[]any` | No |  |
| `submitted_payouts_balance` | `float64` | No |  |
| `submitted_payouts_balance_minor_unit` | `int` | No |  |
| `summary` | `string` | No |  |
| `supplier_sepa_instant_status` | `string` | No |  |
| `xero_bank_feed_connection_status` | `string` | No |  |
| `xero_bank_feed_last_synced_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` | No |  |
| `xero_bank_feed_sync_status` | `string` | No |  |
| `xero_unsynchronised_transactions_count` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsPaymentAccountPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentInitiationEntity

```go
noFrixionMoneyMoovModelsPaymentInitiation := client.NoFrixionMoneyMoovModelsPaymentInitiation(nil)
fmt.Println(noFrixionMoneyMoovModelsPaymentInitiation.GetName()) // "no_frixion_money_moov_models_payment_initiation"
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

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovModelsPaymentInitiation(nil).Create(map[string]any{
    "paymentrequest_id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentInitiationEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestEventEntity

```go
noFrixionMoneyMoovModelsPaymentRequestEvent := client.NoFrixionMoneyMoovModelsPaymentRequestEvent(nil)
fmt.Println(noFrixionMoneyMoovModelsPaymentRequestEvent.GetName()) // "no_frixion_money_moov_models_payment_request_event"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float64` | Yes |  |
| `apple_pay_transaction_id` | `string` | No |  |
| `card_authorization_response_id` | `string` | No |  |
| `card_expiry_month` | `int` | No |  |
| `card_expiry_year` | `int` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsPaymentRequestEvent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestEventEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestMetricEntity

```go
noFrixionMoneyMoovModelsPaymentRequestMetric := client.NoFrixionMoneyMoovModelsPaymentRequestMetric(nil)
fmt.Println(noFrixionMoneyMoovModelsPaymentRequestMetric.GetName()) // "no_frixion_money_moov_models_payment_request_metric"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `all` | `int` | No |  |
| `authorized` | `int` | No |  |
| `paid` | `int` | No |  |
| `partially_paid` | `int` | No |  |
| `total_amounts_by_currency` | `map[string]any` | No |  |
| `unpaid` | `int` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsPaymentRequestMetric(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMetricEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity

```go
noFrixionMoneyMoovModelsPaymentRequestMinimal := client.NoFrixionMoneyMoovModelsPaymentRequestMinimal(nil)
fmt.Println(noFrixionMoneyMoovModelsPaymentRequestMinimal.GetName()) // "no_frixion_money_moov_models_payment_request_minimal"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float64` | No |  |
| `amount_pending` | `float64` | No |  |
| `amount_received` | `float64` | No |  |
| `amount_refunded` | `float64` | No |  |
| `callback_url` | `string` | No |  |
| `card_stripe_payment_intent_secret` | `string` | No |  |
| `country_code` | `string` | No |  |
| `currency` | `string` | No |  |
| `custom_fields_to_display` | `[]any` | No |  |
| `description` | `string` | No |  |
| `due_date` | `string` | No |  |
| `field_display_setting` | `[]any` | No |  |
| `google_pay_merchant_id` | `string` | No |  |
| `id` | `string` | No |  |
| `jwk` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_logo_url_png` | `string` | No |  |
| `merchant_logo_url_svg` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `merchant_short_name` | `string` | No |  |
| `partial_payment_method` | `string` | No |  |
| `payment_attempt` | `[]any` | No |  |
| `payment_methods_list` | `[]any` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_processor_key` | `string` | No |  |
| `pisp_error` | `string` | No |  |
| `priority_bank_id` | `string` | No |  |
| `status` | `string` | No |  |
| `stripe_account_id` | `string` | No |  |
| `title` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsPaymentRequestMinimal(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestResultEntity

```go
noFrixionMoneyMoovModelsPaymentRequestResult := client.NoFrixionMoneyMoovModelsPaymentRequestResult(nil)
fmt.Println(noFrixionMoneyMoovModelsPaymentRequestResult.GetName()) // "no_frixion_money_moov_models_payment_request_result"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float64` | No |  |
| `amount_pending` | `float64` | No |  |
| `amount_received` | `float64` | No |  |
| `amount_refunded` | `float64` | No |  |
| `currency` | `string` | No |  |
| `customer_id` | `string` | No |  |
| `payment` | `[]any` | No |  |
| `payment_request_id` | `string` | No |  |
| `pisp_authorization` | `[]any` | No |  |
| `requested_amount` | `float64` | No |  |
| `result` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsPaymentRequestResult(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestResultEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity

```go
noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(nil)
fmt.Println(noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment.GetName()) // "no_frixion_money_moov_models_payment_requests_merchant_payment"
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
| `template` | `map[string]any` | Yes |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity

```go
noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2 := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(nil)
fmt.Println(noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2.GetName()) // "no_frixion_money_moov_models_payment_requests_merchant_payment2"
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
| `template` | `map[string]any` | Yes |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(nil).Load(map[string]any{"paymentrequest_id": "paymentrequest_id", "template_id": "template_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity

```go
noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3 := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(nil)
fmt.Println(noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3.GetName()) // "no_frixion_money_moov_models_payment_requests_merchant_payment3"
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
| `template` | `map[string]any` | Yes |  |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(nil).Update(map[string]any{
    "paymentrequest_id": "paymentrequest_id",
    "template_id": "template_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity

```go
noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4 := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(nil)
fmt.Println(noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4.GetName()) // "no_frixion_money_moov_models_payment_requests_merchant_payment4"
```

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(nil).Remove(map[string]any{"paymentrequest_id": "paymentrequest_id", "template_id": "template_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayoutKeysetPageEntity

```go
noFrixionMoneyMoovModelsPayoutKeysetPage := client.NoFrixionMoneyMoovModelsPayoutKeysetPage(nil)
fmt.Println(noFrixionMoneyMoovModelsPayoutKeysetPage.GetName()) // "no_frixion_money_moov_models_payout_keyset_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `float64` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `[]any` | No |  |
| `authorisation` | `[]any` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `map[string]any` | Yes |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `map[string]any` | No |  |
| `document` | `[]any` | No |  |
| `event` | `[]any` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `float64` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `float64` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `map[string]any` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float64` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `map[string]any` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `[]any` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `float64` | No |  |
| `transacted_fx_amount` | `float64` | No |  |
| `transacted_fx_rate` | `float64` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsPayoutKeysetPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPayoutKeysetPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayoutMetricEntity

```go
noFrixionMoneyMoovModelsPayoutMetric := client.NoFrixionMoneyMoovModelsPayoutMetric(nil)
fmt.Println(noFrixionMoneyMoovModelsPayoutMetric.GetName()) // "no_frixion_money_moov_models_payout_metric"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `all` | `float64` | No |  |
| `failed` | `float64` | No |  |
| `in_progress` | `float64` | No |  |
| `paid` | `float64` | No |  |
| `pending_approval` | `float64` | No |  |
| `scheduled` | `float64` | No |  |
| `total_amounts_by_currency` | `map[string]any` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsPayoutMetric(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPayoutMetricEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity

```go
noFrixionMoneyMoovModelsPayoutsPayoutsCreate := client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(nil)
fmt.Println(noFrixionMoneyMoovModelsPayoutsPayoutsCreate.GetName()) // "no_frixion_money_moov_models_payouts_payouts_create"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_payout` | `map[string]any` | No |  |
| `payout` | `[]any` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayrunEntity

```go
noFrixionMoneyMoovModelsPayrun := client.NoFrixionMoneyMoovModelsPayrun(nil)
fmt.Println(noFrixionMoneyMoovModelsPayrun.GetName()) // "no_frixion_money_moov_models_payrun"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation` | `[]any` | No |  |
| `authorisation_date` | `string` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_delete` | `bool` | No |  |
| `can_edit` | `bool` | No |  |
| `event` | `[]any` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice` | `[]any` | No |  |
| `invoices_minimal` | `[]any` | No |  |
| `is_archived` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `last_updated_by` | `map[string]any` | Yes |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment` | `[]any` | No |  |
| `payout` | `[]any` | No |  |
| `payouts_count` | `int` | No |  |
| `reason` | `string` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled_date` | `string` | No |  |
| `source_account` | `[]any` | No |  |
| `status` | `string` | No |  |
| `total_eur` | `float64` | No |  |
| `total_gbp` | `float64` | No |  |
| `total_usd` | `float64` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsPayrun(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_payrun_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovModelsPayrun(nil).Create(map[string]any{
    "id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.NoFrixionMoneyMoovModelsPayrun(nil).Update(map[string]any{
    "id": "no_frixion_money_moov_models_payrun_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsPayrunEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsReportResultEntity

```go
noFrixionMoneyMoovModelsReportResult := client.NoFrixionMoneyMoovModelsReportResult(nil)
fmt.Println(noFrixionMoneyMoovModelsReportResult.GetName()) // "no_frixion_money_moov_models_report_result"
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
| `statement_number` | `int` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsReportResult(nil).Load(map[string]any{"id": 1, "report_id": "report_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsReportResultEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsRuleEntity

```go
noFrixionMoneyMoovModelsRule := client.NoFrixionMoneyMoovModelsRule(nil)
fmt.Println(noFrixionMoneyMoovModelsRule.GetName()) // "no_frixion_money_moov_models_rule"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account` | `map[string]any` | No |  |
| `account_id` | `string` | No |  |
| `approve_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `[]any` | No |  |
| `authorisation` | `[]any` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `can_authorise` | `bool` | No |  |
| `created_by` | `map[string]any` | Yes |  |
| `description` | `string` | No |  |
| `end_at` | `string` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_disabled` | `bool` | No |  |
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
| `sweep_action` | `map[string]any` | No |  |
| `time_zone_id` | `string` | No |  |
| `trigger_cron_expression` | `string` | No |  |
| `trigger_on_pay_in` | `bool` | No |  |
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsRule(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_rule_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovModelsRule(nil).Create(map[string]any{
    "created_by": /* map[string]any */,
    "nonce": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.NoFrixionMoneyMoovModelsRule(nil).Update(map[string]any{
    "id": "no_frixion_money_moov_models_rule_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsRuleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsTransactionEntity

```go
noFrixionMoneyMoovModelsTransaction := client.NoFrixionMoneyMoovModelsTransaction(nil)
fmt.Println(noFrixionMoneyMoovModelsTransaction.GetName()) // "no_frixion_money_moov_models_transaction"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `account_name` | `string` | No |  |
| `account_sequence_number` | `int` | No |  |
| `amount` | `float64` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `balance` | `float64` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `counterparty` | `map[string]any` | No |  |
| `counterparty_summary` | `string` | No |  |
| `currency` | `string` | No |  |
| `description` | `string` | No |  |
| `fx_amount` | `float64` | No |  |
| `fx_currency` | `string` | No |  |
| `fx_rate` | `float64` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `payment_request_custom_field` | `map[string]any` | No |  |
| `payment_request_id` | `string` | No |  |
| `payout_id` | `string` | No |  |
| `raw_reference` | `string` | No |  |
| `rule_id` | `string` | No |  |
| `tag` | `[]any` | No |  |
| `their_reference` | `string` | No |  |
| `transaction_date` | `string` | No |  |
| `type` | `string` | No |  |
| `virtual_iban` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsTransaction(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_transaction_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsTransactionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsTransactionPageEntity

```go
noFrixionMoneyMoovModelsTransactionPage := client.NoFrixionMoneyMoovModelsTransactionPage(nil)
fmt.Println(noFrixionMoneyMoovModelsTransactionPage.GetName()) // "no_frixion_money_moov_models_transaction_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `account_name` | `string` | No |  |
| `account_sequence_number` | `int` | No |  |
| `amount` | `float64` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `balance` | `float64` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `content` | `[]any` | No |  |
| `counterparty` | `map[string]any` | No |  |
| `counterparty_summary` | `string` | No |  |
| `currency` | `string` | No |  |
| `description` | `string` | No |  |
| `fx_amount` | `float64` | No |  |
| `fx_currency` | `string` | No |  |
| `fx_rate` | `float64` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `page_number` | `int` | No |  |
| `page_size` | `int` | No |  |
| `payment_request_custom_field` | `map[string]any` | No |  |
| `payment_request_id` | `string` | No |  |
| `payout_id` | `string` | No |  |
| `raw_reference` | `string` | No |  |
| `rule_id` | `string` | No |  |
| `tag` | `[]any` | No |  |
| `their_reference` | `string` | No |  |
| `total_page` | `int` | No |  |
| `total_size` | `int` | No |  |
| `transaction_date` | `string` | No |  |
| `type` | `string` | No |  |
| `virtual_iban` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsTransactionPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsTransactionPage(nil).Load(map[string]any{"account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsTransactionPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsUserInviteEntity

```go
noFrixionMoneyMoovModelsUserInvite := client.NoFrixionMoneyMoovModelsUserInvite(nil)
fmt.Println(noFrixionMoneyMoovModelsUserInvite.GetName()) // "no_frixion_money_moov_models_user_invite"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_status` | `map[string]any` | No |  |
| `id` | `string` | No |  |
| `initial_role_id` | `string` | No |  |
| `invitee_email_address` | `string` | No |  |
| `invitee_first_name` | `string` | No |  |
| `invitee_last_name` | `string` | No |  |
| `inviter_email_address` | `string` | No |  |
| `inviter_first_name` | `string` | No |  |
| `inviter_last_name` | `string` | No |  |
| `is_authorised` | `bool` | No |  |
| `is_invitee_registered` | `bool` | No |  |
| `last_invited` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `message` | `string` | No |  |
| `registration_url` | `string` | No |  |
| `send_invite_email` | `bool` | No |  |
| `status` | `string` | No |  |
| `user` | `map[string]any` | Yes |  |
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsUserInvite(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_user_invite_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovModelsUserInvite(nil).Create(map[string]any{
    "user": /* map[string]any */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsUserInviteEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsUserInvitePageEntity

```go
noFrixionMoneyMoovModelsUserInvitePage := client.NoFrixionMoneyMoovModelsUserInvitePage(nil)
fmt.Println(noFrixionMoneyMoovModelsUserInvitePage.GetName()) // "no_frixion_money_moov_models_user_invite_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_status` | `map[string]any` | No |  |
| `id` | `string` | No |  |
| `initial_role_id` | `string` | No |  |
| `invitee_email_address` | `string` | No |  |
| `invitee_first_name` | `string` | No |  |
| `invitee_last_name` | `string` | No |  |
| `inviter_email_address` | `string` | No |  |
| `inviter_first_name` | `string` | No |  |
| `inviter_last_name` | `string` | No |  |
| `is_authorised` | `bool` | No |  |
| `is_invitee_registered` | `bool` | No |  |
| `last_invited` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `message` | `string` | No |  |
| `registration_url` | `string` | No |  |
| `status` | `string` | No |  |
| `user` | `map[string]any` | Yes |  |
| `user_id` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsUserInvitePage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsUserInvitePageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsUserPageEntity

```go
noFrixionMoneyMoovModelsUserPage := client.NoFrixionMoneyMoovModelsUserPage(nil)
fmt.Println(noFrixionMoneyMoovModelsUserPage.GetName()) // "no_frixion_money_moov_models_user_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `[]any` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `bool` | No |  |
| `permission` | `map[string]any` | No |  |
| `roles_with_scope` | `[]any` | No |  |
| `two_factor_enabled` | `bool` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsUserPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsUserPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsWebhookEntity

```go
noFrixionMoneyMoovModelsWebhook := client.NoFrixionMoneyMoovModelsWebhook(nil)
fmt.Println(noFrixionMoneyMoovModelsWebhook.GetName()) // "no_frixion_money_moov_models_webhook"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destination_url` | `string` | No |  |
| `email_address` | `string` | No |  |
| `failed_notification_email_address` | `string` | No |  |
| `id` | `string` | No |  |
| `is_active` | `bool` | No |  |
| `merchant_id` | `string` | No |  |
| `notification_method` | `string` | No |  |
| `resource_type` | `[]any` | No |  |
| `retry` | `bool` | No |  |
| `secret` | `string` | No |  |
| `version` | `int` | No |  |

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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.NoFrixionMoneyMoovModelsWebhook(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionMoneyMoovModelsWebhook(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_webhook_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.NoFrixionMoneyMoovModelsWebhook(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.NoFrixionMoneyMoovModelsWebhook(nil).Update(map[string]any{
    "id": "no_frixion_money_moov_models_webhook_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionMoneyMoovModelsWebhookEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## OpenBankingEntity

```go
openBanking := client.OpenBanking(nil)
fmt.Println(openBanking.GetName()) // "open_banking"
```

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.OpenBanking(nil).Create(map[string]any{
    "account_id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.OpenBanking(nil).Remove(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `OpenBankingEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PayeeverificationEntity

```go
payeeverification := client.Payeeverification(nil)
fmt.Println(payeeverification.GetName()) // "payeeverification"
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

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Payeeverification(nil).Create(map[string]any{
    "account_name": /* string */,
    "iban": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PayeeverificationEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PaymentRequestEntity

```go
paymentRequest := client.PaymentRequest(nil)
fmt.Println(paymentRequest.GetName()) // "payment_request"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float64` | No |  |
| `do_simulate_settlement_failure` | `bool` | No |  |
| `error_description` | `string` | No |  |
| `institution` | `string` | No |  |
| `payment_initiation_id` | `string` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.PaymentRequest(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PaymentRequest(nil).Create(map[string]any{
    "paymentrequest_id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.PaymentRequest(nil).Update(map[string]any{
    "paymentrequest_id": "paymentrequest_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.PaymentRequest(nil).Remove(map[string]any{"id": "id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PaymentRequestEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PayoutEntity

```go
payout := client.Payout(nil)
fmt.Println(payout.GetName()) // "payout"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `allow_incomplete` | `bool` | No |  |
| `amount` | `float64` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `[]any` | No |  |
| `authorisation` | `[]any` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `map[string]any` | Yes |  |
| `beneficiary_id` | `string` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `map[string]any` | No |  |
| `document` | `[]any` | No |  |
| `event` | `[]any` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `float64` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `float64` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `map[string]any` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float64` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `map[string]any` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `[]any` | No |  |
| `tag_id` | `[]any` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `float64` | No |  |
| `transacted_fx_amount` | `float64` | No |  |
| `transacted_fx_rate` | `float64` | No |  |
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Payout(nil).Load(map[string]any{"id": "payout_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Payout(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Payout(nil).Update(map[string]any{
    "id": "payout_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Payout(nil).Remove(map[string]any{"id": "payout_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PayoutEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PayrunEntity

```go
payrun := client.Payrun(nil)
fmt.Println(payrun.GetName()) // "payrun"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |
| `note` | `string` | No |  |
| `scheduled_date` | `string` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Payrun(nil).Create(map[string]any{
    "id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Payrun(nil).Update(map[string]any{
    "id": "id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Payrun(nil).Remove(map[string]any{"id": "id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PayrunEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## RejectEntity

```go
reject := client.Reject(nil)
fmt.Println(reject.GetName()) // "reject"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `float64` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `[]any` | No |  |
| `authorisation` | `[]any` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `map[string]any` | Yes |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `map[string]any` | No |  |
| `document` | `[]any` | No |  |
| `event` | `[]any` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `float64` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `float64` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `reason` | `string` | No |  |
| `rule` | `map[string]any` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float64` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `map[string]any` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `[]any` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `float64` | No |  |
| `transacted_fx_amount` | `float64` | No |  |
| `transacted_fx_rate` | `float64` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Reject(nil).Update(map[string]any{
    "id": "id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `RejectEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ReportEntity

```go
report := client.Report(nil)
fmt.Println(report.GetName()) // "report"
```

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Report(nil).Update(map[string]any{
    "id": "id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ReportEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## RuleEntity

```go
rule := client.Rule(nil)
fmt.Println(rule.GetName()) // "rule"
```

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Rule(nil).Update(map[string]any{
    "id": "id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Rule(nil).Remove(map[string]any{"id": "id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `RuleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SendEntity

```go
send := client.Send(nil)
fmt.Println(send.GetName()) // "send"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `allow_incomplete` | `bool` | No |  |
| `amount` | `float64` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `[]any` | No |  |
| `authorisation` | `[]any` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `map[string]any` | Yes |  |
| `beneficiary_id` | `string` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `map[string]any` | No |  |
| `document` | `[]any` | No |  |
| `event` | `[]any` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `float64` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `float64` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `map[string]any` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float64` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `map[string]any` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `[]any` | No |  |
| `tag_id` | `[]any` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `float64` | No |  |
| `transacted_fx_amount` | `float64` | No |  |
| `transacted_fx_rate` | `float64` | No |  |
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

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Send(nil).Create(map[string]any{
    "beneficiary": /* map[string]any */,
    "source_account_identifier": /* map[string]any */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SendEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SendbeneficiaryEntity

```go
sendbeneficiary := client.Sendbeneficiary(nil)
fmt.Println(sendbeneficiary.GetName()) // "sendbeneficiary"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `allow_incomplete` | `bool` | No |  |
| `amount` | `float64` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `[]any` | No |  |
| `authorisation` | `[]any` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `map[string]any` | Yes |  |
| `beneficiary_id` | `string` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `map[string]any` | No |  |
| `document` | `[]any` | No |  |
| `event` | `[]any` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `float64` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `float64` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `map[string]any` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float64` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `map[string]any` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `[]any` | No |  |
| `tag_id` | `[]any` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `float64` | No |  |
| `transacted_fx_amount` | `float64` | No |  |
| `transacted_fx_rate` | `float64` | No |  |
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

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Sendbeneficiary(nil).Create(map[string]any{
    "beneficiary": /* map[string]any */,
    "source_account_identifier": /* map[string]any */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SendbeneficiaryEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TagEntity

```go
tag := client.Tag(nil)
fmt.Println(tag.GetName()) // "tag"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Tag(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Tag(nil).Create(map[string]any{
    "merchant_id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TagEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TokenEntity

```go
token := client.Token(nil)
fmt.Println(token.GetName()) // "token"
```

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Token(nil).Create(map[string]any{
    "id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Token(nil).Remove(map[string]any{"id": "id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TokenEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TransactionEntity

```go
transaction := client.Transaction(nil)
fmt.Println(transaction.GetName()) // "transaction"
```

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Transaction(nil).Load(map[string]any{"sequence_number": 1, "transaction_id": "transaction_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Transaction(nil).Create(map[string]any{
    "id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Transaction(nil).Remove(map[string]any{"id": "id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TransactionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## UserEntity

```go
user := client.User(nil)
fmt.Println(user.GetName()) // "user"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `[]any` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `bool` | No |  |
| `permission` | `map[string]any` | No |  |
| `profile` | `string` | No |  |
| `roles_with_scope` | `[]any` | No |  |
| `two_factor_enabled` | `bool` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.User(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.User(nil).Update(map[string]any{
    "id": "id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `UserEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## UserInviteEntity

```go
userInvite := client.UserInvite(nil)
fmt.Println(userInvite.GetName()) // "user_invite"
```

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.UserInvite(nil).Create(map[string]any{
    "id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.UserInvite(nil).Update(map[string]any{
    "id": "id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.UserInvite(nil).Remove(map[string]any{"id": "id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `UserInviteEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## VirtualEntity

```go
virtual := client.Virtual(nil)
fmt.Println(virtual.GetName()) // "virtual"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_name` | `string` | No |  |
| `account_supplier_name` | `string` | No |  |
| `available_balance` | `float64` | No |  |
| `available_balance_minor_unit` | `int` | No |  |
| `balance` | `float64` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `bank_name` | `string` | No |  |
| `consent_id` | `string` | No |  |
| `created_by` | `map[string]any` | Yes |  |
| `created_by_display_name` | `string` | No |  |
| `currency` | `string` | No |  |
| `default_payment_rail` | `string` | No |  |
| `display_name` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `external_account_icon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `map[string]any` | Yes |  |
| `inserted` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_connected_account` | `bool` | No |  |
| `is_default` | `bool` | No |  |
| `is_trust_account` | `bool` | No |  |
| `is_virtual` | `bool` | No |  |
| `last_transaction` | `map[string]any` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `name` | `string` | Yes |  |
| `physical_account_id` | `string` | No |  |
| `rule` | `[]any` | No |  |
| `submitted_payouts_balance` | `float64` | No |  |
| `submitted_payouts_balance_minor_unit` | `int` | No |  |
| `summary` | `string` | No |  |
| `supplier_sepa_instant_status` | `string` | No |  |
| `xero_bank_feed_connection_status` | `string` | No |  |
| `xero_bank_feed_last_synced_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` | No |  |
| `xero_bank_feed_sync_status` | `string` | No |  |
| `xero_unsynchronised_transactions_count` | `int` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Virtual(nil).Create(map[string]any{
    "account_id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Virtual(nil).Update(map[string]any{
    "account_id": "account_id",
    "id": "id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `VirtualEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## WebhookEntity

```go
webhook := client.Webhook(nil)
fmt.Println(webhook.GetName()) // "webhook"
```

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Webhook(nil).Remove(map[string]any{"id": "id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `WebhookEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## WhoamiEntity

```go
whoami := client.Whoami(nil)
fmt.Println(whoami.GetName()) // "whoami"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `[]any` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `bool` | No |  |
| `permission` | `map[string]any` | No |  |
| `roles_with_scope` | `[]any` | No |  |
| `two_factor_enabled` | `bool` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Whoami(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `WhoamiEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## WhoamitrustedappEntity

```go
whoamitrustedapp := client.Whoamitrustedapp(nil)
fmt.Println(whoamitrustedapp.GetName()) // "whoamitrustedapp"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `[]any` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `bool` | No |  |
| `permission` | `map[string]any` | No |  |
| `roles_with_scope` | `[]any` | No |  |
| `two_factor_enabled` | `bool` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Whoamitrustedapp(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `WhoamitrustedappEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewNofrixionSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

