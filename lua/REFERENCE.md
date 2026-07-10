# Nofrixion Lua SDK Reference

Complete API reference for the Nofrixion Lua SDK.


## NofrixionSDK

### Constructor

```lua
local sdk = require("nofrixion_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `Account(data)`

Create a new `Account` entity instance. Pass `nil` for no initial data.

#### `Beneficiary(data)`

Create a new `Beneficiary` entity instance. Pass `nil` for no initial data.

#### `Cancel(data)`

Create a new `Cancel` entity instance. Pass `nil` for no initial data.

#### `Disable(data)`

Create a new `Disable` entity instance. Pass `nil` for no initial data.

#### `Enable(data)`

Create a new `Enable` entity instance. Pass `nil` for no initial data.

#### `Merchant(data)`

Create a new `Merchant` entity instance. Pass `nil` for no initial data.

#### `Metadata(data)`

Create a new `Metadata` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(data)`

Create a new `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingPaymentRequestPage(data)`

Create a new `NoFrixionBizBizModelsPagingPaymentRequestPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingPayoutPage(data)`

Create a new `NoFrixionBizBizModelsPagingPayoutPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingPayrunPage(data)`

Create a new `NoFrixionBizBizModelsPagingPayrunPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingRuleEventsPage(data)`

Create a new `NoFrixionBizBizModelsPagingRuleEventsPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingRulesPage(data)`

Create a new `NoFrixionBizBizModelsPagingRulesPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPaymentsCardPayment(data)`

Create a new `NoFrixionBizBizModelsPaymentsCardPayment` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPaymentsCardPublicKey(data)`

Create a new `NoFrixionBizBizModelsPaymentsCardPublicKey` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(data)`

Create a new `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(data)`

Create a new `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(data)`

Create a new `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(data)`

Create a new `NoFrixionMoneyMoovApiFeaturesUserInvitesCreate` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(data)`

Create a new `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsBatchPayout(data)`

Create a new `NoFrixionMoneyMoovModelsBatchPayout` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsBeneficiaryGroupPage(data)`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryGroupPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsBeneficiaryPage(data)`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsCardCustomerToken(data)`

Create a new `NoFrixionMoneyMoovModelsCardCustomerToken` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(data)`

Create a new `NoFrixionMoneyMoovModelsCurrencyCurrencyInfo` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(data)`

Create a new `NoFrixionMoneyMoovModelsDirectDebitBatchSubmit` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsFxRate(data)`

Create a new `NoFrixionMoneyMoovModelsFxRate` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsIPayment(data)`

Create a new `NoFrixionMoneyMoovModelsIPayment` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMandatesMandate(data)`

Create a new `NoFrixionMoneyMoovModelsMandatesMandate` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchant(data)`

Create a new `NoFrixionMoneyMoovModelsMerchant` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantPage(data)`

Create a new `NoFrixionMoneyMoovModelsMerchantPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantPayByBankSetting(data)`

Create a new `NoFrixionMoneyMoovModelsMerchantPayByBankSetting` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantToken(data)`

Create a new `NoFrixionMoneyMoovModelsMerchantToken` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantTokenPage(data)`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsNoFrixionVersion(data)`

Create a new `NoFrixionMoneyMoovModelsNoFrixionVersion` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsOpenBankingAccount(data)`

Create a new `NoFrixionMoneyMoovModelsOpenBankingAccount` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsOpenBankingConsent(data)`

Create a new `NoFrixionMoneyMoovModelsOpenBankingConsent` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsOpenBankingTransaction(data)`

Create a new `NoFrixionMoneyMoovModelsOpenBankingTransaction` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPayment(data)`

Create a new `NoFrixionMoneyMoovModelsPayment` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(data)`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountMinimalPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentAccountPage(data)`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentInitiation(data)`

Create a new `NoFrixionMoneyMoovModelsPaymentInitiation` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestEvent(data)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestEvent` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestMetric(data)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMetric` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestMinimal(data)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMinimal` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestResult(data)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestResult` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(data)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(data)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(data)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(data)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPayoutKeysetPage(data)`

Create a new `NoFrixionMoneyMoovModelsPayoutKeysetPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPayoutMetric(data)`

Create a new `NoFrixionMoneyMoovModelsPayoutMetric` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(data)`

Create a new `NoFrixionMoneyMoovModelsPayoutsPayoutsCreate` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPayrun(data)`

Create a new `NoFrixionMoneyMoovModelsPayrun` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsReportResult(data)`

Create a new `NoFrixionMoneyMoovModelsReportResult` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsRule(data)`

Create a new `NoFrixionMoneyMoovModelsRule` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsTransaction(data)`

Create a new `NoFrixionMoneyMoovModelsTransaction` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsTransactionPage(data)`

Create a new `NoFrixionMoneyMoovModelsTransactionPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsUserInvite(data)`

Create a new `NoFrixionMoneyMoovModelsUserInvite` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsUserInvitePage(data)`

Create a new `NoFrixionMoneyMoovModelsUserInvitePage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsUserPage(data)`

Create a new `NoFrixionMoneyMoovModelsUserPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsWebhook(data)`

Create a new `NoFrixionMoneyMoovModelsWebhook` entity instance. Pass `nil` for no initial data.

#### `OpenBanking(data)`

Create a new `OpenBanking` entity instance. Pass `nil` for no initial data.

#### `Payeeverification(data)`

Create a new `Payeeverification` entity instance. Pass `nil` for no initial data.

#### `PaymentRequest(data)`

Create a new `PaymentRequest` entity instance. Pass `nil` for no initial data.

#### `Payout(data)`

Create a new `Payout` entity instance. Pass `nil` for no initial data.

#### `Payrun(data)`

Create a new `Payrun` entity instance. Pass `nil` for no initial data.

#### `Reject(data)`

Create a new `Reject` entity instance. Pass `nil` for no initial data.

#### `Report(data)`

Create a new `Report` entity instance. Pass `nil` for no initial data.

#### `Rule(data)`

Create a new `Rule` entity instance. Pass `nil` for no initial data.

#### `Send(data)`

Create a new `Send` entity instance. Pass `nil` for no initial data.

#### `Sendbeneficiary(data)`

Create a new `Sendbeneficiary` entity instance. Pass `nil` for no initial data.

#### `Tag(data)`

Create a new `Tag` entity instance. Pass `nil` for no initial data.

#### `Token(data)`

Create a new `Token` entity instance. Pass `nil` for no initial data.

#### `Transaction(data)`

Create a new `Transaction` entity instance. Pass `nil` for no initial data.

#### `User(data)`

Create a new `User` entity instance. Pass `nil` for no initial data.

#### `UserInvite(data)`

Create a new `UserInvite` entity instance. Pass `nil` for no initial data.

#### `Virtual(data)`

Create a new `Virtual` entity instance. Pass `nil` for no initial data.

#### `Webhook(data)`

Create a new `Webhook` entity instance. Pass `nil` for no initial data.

#### `Whoami(data)`

Create a new `Whoami` entity instance. Pass `nil` for no initial data.

#### `Whoamitrustedapp(data)`

Create a new `Whoamitrustedapp` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## AccountEntity

```lua
local account = client:Account(nil)
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
| `created_by` | `table` | Yes |  |
| `created_by_display_name` | `string` | No |  |
| `currency` | `string` | No |  |
| `default_payment_rail` | `string` | No |  |
| `display_name` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `external_account_icon` | `string` | No |  |
| `format` | `string` | No |  |
| `from_date` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `table` | Yes |  |
| `inserted` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_connected_account` | `boolean` | No |  |
| `is_default` | `boolean` | No |  |
| `is_trust_account` | `boolean` | No |  |
| `is_virtual` | `boolean` | No |  |
| `last_transaction` | `table` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `physical_account_id` | `string` | No |  |
| `role_i_d` | `table` | No |  |
| `rule` | `table` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Account():create({
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Account():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Account():load({ id = "account_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Account():remove({ id = "account_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Account():update({
  id = "account_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AccountEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## BeneficiaryEntity

```lua
local beneficiary = client:Beneficiary(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `table` | No |  |
| `authorisation` | `table` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `beneficiary_event` | `table` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `created_by` | `table` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `table` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `boolean` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `table` | No |  |
| `source_account_i_d` | `table` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Beneficiary():create({
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Beneficiary():load({ id = "beneficiary_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Beneficiary():remove({ id = "beneficiary_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Beneficiary():update({
  id = "beneficiary_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BeneficiaryEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CancelEntity

```lua
local cancel = client:Cancel(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `table` | No |  |
| `authorisation` | `table` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `table` | Yes |  |
| `can_authorise` | `boolean` | No |  |
| `can_process` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `table` | No |  |
| `document` | `table` | No |  |
| `event` | `table` | No |  |
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
| `rule` | `table` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `table` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `table` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `number` | No |  |
| `transacted_fx_amount` | `number` | No |  |
| `transacted_fx_rate` | `number` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Cancel():update({
  id = "id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CancelEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## DisableEntity

```lua
local disable = client:Disable(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `table` | No |  |
| `authorisation` | `table` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `beneficiary_event` | `table` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `created_by` | `table` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `table` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `boolean` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `table` | No |  |
| `their_reference` | `string` | No |  |

### Operations

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Disable():update({
  id = "id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DisableEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## EnableEntity

```lua
local enable = client:Enable(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `table` | No |  |
| `authorisation` | `table` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `beneficiary_event` | `table` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `created_by` | `table` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `table` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `boolean` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `table` | No |  |
| `their_reference` | `string` | No |  |

### Operations

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Enable():update({
  id = "id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `EnableEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MerchantEntity

```lua
local merchant = client:Merchant(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `reason` | `string` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Merchant():load({ merchant_id = "merchant_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Merchant():remove()
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Merchant():update({
  id = "id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MerchantEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MetadataEntity

```lua
local metadata = client:Metadata(nil)
```

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Metadata():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MetadataEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity

```lua
local no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page = client:NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingPaymentRequestPageEntity

```lua
local no_frixion_biz_biz_models_paging_payment_request_page = client:NoFrixionBizBizModelsPagingPaymentRequestPage(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `table` | No |  |
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
| `created_by_user` | `table` | Yes |  |
| `currency` | `string` | No |  |
| `custom_field` | `table` | No |  |
| `customer_email_address` | `string` | No |  |
| `customer_id` | `string` | No |  |
| `customer_name` | `string` | No |  |
| `description` | `string` | No |  |
| `destination_account` | `table` | No |  |
| `direct_debit_payment` | `table` | No |  |
| `due_date` | `string` | No |  |
| `event` | `table` | No |  |
| `failure_callback_url` | `string` | No |  |
| `field_display_setting` | `table` | No |  |
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
| `notification_role_i_d` | `table` | No |  |
| `order_id` | `string` | No |  |
| `partial_payment_method` | `string` | No |  |
| `partial_payment_step` | `string` | No |  |
| `payment_attempt` | `table` | No |  |
| `payment_method` | `table` | No |  |
| `payment_processor` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `pisp_account_id` | `string` | No |  |
| `priority_bank_id` | `string` | No |  |
| `result` | `table` | No |  |
| `sandbox_settle_delay_in_second` | `number` | No |  |
| `shipping_address` | `table` | No |  |
| `status` | `string` | No |  |
| `success_web_hook_url` | `string` | No |  |
| `tag` | `table` | No |  |
| `title` | `string` | No |  |
| `tokenised_card` | `table` | No |  |
| `transaction` | `table` | No |  |
| `use_hosted_payment_page` | `boolean` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionBizBizModelsPagingPaymentRequestPage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionBizBizModelsPagingPaymentRequestPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingPayoutPageEntity

```lua
local no_frixion_biz_biz_models_paging_payout_page = client:NoFrixionBizBizModelsPagingPayoutPage(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `table` | No |  |
| `authorisation` | `table` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `table` | Yes |  |
| `can_authorise` | `boolean` | No |  |
| `can_process` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `table` | No |  |
| `document` | `table` | No |  |
| `event` | `table` | No |  |
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
| `rule` | `table` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `table` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `table` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `number` | No |  |
| `transacted_fx_amount` | `number` | No |  |
| `transacted_fx_rate` | `number` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionBizBizModelsPagingPayoutPage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionBizBizModelsPagingPayoutPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingPayrunPageEntity

```lua
local no_frixion_biz_biz_models_paging_payrun_page = client:NoFrixionBizBizModelsPagingPayrunPage(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation` | `table` | No |  |
| `authorisation_date` | `string` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_delete` | `boolean` | No |  |
| `can_edit` | `boolean` | No |  |
| `event` | `table` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice` | `table` | No |  |
| `invoices_minimal` | `table` | No |  |
| `is_archived` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `last_updated_by` | `table` | Yes |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment` | `table` | No |  |
| `payout` | `table` | No |  |
| `payouts_count` | `number` | No |  |
| `schedule_date` | `string` | No |  |
| `source_account` | `table` | No |  |
| `status` | `string` | No |  |
| `total_eur` | `number` | No |  |
| `total_gbp` | `number` | No |  |
| `total_usd` | `number` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionBizBizModelsPagingPayrunPage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionBizBizModelsPagingPayrunPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingRuleEventsPageEntity

```lua
local no_frixion_biz_biz_models_paging_rule_events_page = client:NoFrixionBizBizModelsPagingRuleEventsPage(nil)
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
| `user` | `table` | Yes |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionBizBizModelsPagingRuleEventsPage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionBizBizModelsPagingRuleEventsPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingRulesPageEntity

```lua
local no_frixion_biz_biz_models_paging_rules_page = client:NoFrixionBizBizModelsPagingRulesPage(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account` | `table` | No |  |
| `account_id` | `string` | No |  |
| `approve_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `table` | No |  |
| `authorisation` | `table` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `can_authorise` | `boolean` | No |  |
| `created_by` | `table` | Yes |  |
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
| `sweep_action` | `table` | No |  |
| `time_zone_id` | `string` | No |  |
| `trigger_cron_expression` | `string` | No |  |
| `trigger_on_pay_in` | `boolean` | No |  |
| `user_id` | `string` | No |  |
| `web_hook_secret` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionBizBizModelsPagingRulesPage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionBizBizModelsPagingRulesPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionBizBizModelsPaymentsCardPaymentEntity

```lua
local no_frixion_biz_biz_models_payments_card_payment = client:NoFrixionBizBizModelsPaymentsCardPayment(nil)
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionBizBizModelsPaymentsCardPayment():create({
  paymentrequest_id = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionBizBizModelsPaymentsCardPaymentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionBizBizModelsPaymentsCardPublicKeyEntity

```lua
local no_frixion_biz_biz_models_payments_card_public_key = client:NoFrixionBizBizModelsPaymentsCardPublicKey(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jwt` | `string` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionBizBizModelsPaymentsCardPublicKey():load({ paymentrequest_id = "paymentrequest_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionBizBizModelsPaymentsCardPublicKeyEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity

```lua
local no_frixion_money_moov_api_features_beneficiaries_beneficiaries = client:NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `beneficiary` | `table` | No |  |
| `failed_beneficiary` | `table` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries():create({
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity

```lua
local no_frixion_money_moov_api_features_payment_requests_payment = client:NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_payment_request` | `table` | No |  |
| `payment_request` | `table` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment():create({
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity

```lua
local no_frixion_money_moov_api_features_permissions_roles_create = client:NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_role` | `table` | No |  |
| `role` | `table` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate():create({
  merchant_id = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity

```lua
local no_frixion_money_moov_api_features_user_invites_create = client:NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_user_invite` | `table` | No |  |
| `user_invite` | `table` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovApiFeaturesUserInvitesCreate():create({
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity

```lua
local no_frixion_money_moov_models_authorisation_settings_merchant = client:NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(nil)
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
| `role_setting` | `table` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsBatchPayoutEntity

```lua
local no_frixion_money_moov_models_batch_payout = client:NoFrixionMoneyMoovModelsBatchPayout(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approve_url` | `string` | No |  |
| `id` | `string` | No |  |
| `payout` | `table` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovModelsBatchPayout():create({
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsBatchPayout():load({ id = "no_frixion_money_moov_models_batch_payout_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsBatchPayoutEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity

```lua
local no_frixion_money_moov_models_beneficiary_group_page = client:NoFrixionMoneyMoovModelsBeneficiaryGroupPage(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `group_member` | `table` | No |  |
| `group_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | Yes |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsBeneficiaryGroupPage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsBeneficiaryPageEntity

```lua
local no_frixion_money_moov_models_beneficiary_page = client:NoFrixionMoneyMoovModelsBeneficiaryPage(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `table` | No |  |
| `authorisation` | `table` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `beneficiary_event` | `table` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `created_by` | `table` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `table` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `boolean` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `table` | No |  |
| `their_reference` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsBeneficiaryPage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsCardCustomerTokenEntity

```lua
local no_frixion_money_moov_models_card_customer_token = client:NoFrixionMoneyMoovModelsCardCustomerToken(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsCardCustomerToken():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsCardCustomerToken():load({ customer_email_address = "customer_email_address" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsCardCustomerToken():remove()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsCardCustomerTokenEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity

```lua
local no_frixion_money_moov_models_currency_currency_info = client:NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsCurrencyCurrencyInfo():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity

```lua
local no_frixion_money_moov_models_direct_debit_batch_submit = client:NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_submission` | `table` | No |  |
| `successful_submission` | `table` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovModelsDirectDebitBatchSubmit():create({
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsFxRateEntity

```lua
local no_frixion_money_moov_models_fx_rate = client:NoFrixionMoneyMoovModelsFxRate(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsFxRate():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsFxRate():load({ destination = "destination", source = "source", valid_for_minute = 1 })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsFxRateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsIPaymentEntity

```lua
local no_frixion_money_moov_models_i_payment = client:NoFrixionMoneyMoovModelsIPayment(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `payment_request_id` | `string` | No |  |
| `response_type` | `string` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovModelsIPayment():create({
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsIPaymentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMandatesMandateEntity

```lua
local no_frixion_money_moov_models_mandates_mandate = client:NoFrixionMoneyMoovModelsMandatesMandate(nil)
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovModelsMandatesMandate():create({
  address_line1 = --[[ string ]],
  city = --[[ string ]],
  country_code = --[[ string ]],
  email_address = --[[ string ]],
  first_name = --[[ string ]],
  last_name = --[[ string ]],
  postal_code = --[[ string ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsMandatesMandate():load({ id = "no_frixion_money_moov_models_mandates_mandate_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsMandatesMandateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantEntity

```lua
local no_frixion_money_moov_models_merchant = client:NoFrixionMoneyMoovModelsMerchant(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_currency` | `table` | No |  |
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
| `parent_merchant` | `table` | No |  |
| `payment_account` | `table` | No |  |
| `payment_account_limit` | `number` | No |  |
| `short_name` | `string` | No |  |
| `supported_payment_methods_list` | `table` | No |  |
| `suspension_reason` | `string` | No |  |
| `tag` | `table` | No |  |
| `time_zone_id` | `string` | No |  |
| `trading_name` | `string` | No |  |
| `web_hook_limit` | `number` | No |  |
| `your_role_name` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsMerchant():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsMerchant():load({ id = "no_frixion_money_moov_models_merchant_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:NoFrixionMoneyMoovModelsMerchant():update({
  id = "no_frixion_money_moov_models_merchant_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsMerchantEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantPageEntity

```lua
local no_frixion_money_moov_models_merchant_page = client:NoFrixionMoneyMoovModelsMerchantPage(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_currency` | `table` | No |  |
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
| `parent_merchant` | `table` | No |  |
| `payment_account` | `table` | No |  |
| `payment_account_limit` | `number` | No |  |
| `short_name` | `string` | No |  |
| `supported_payment_methods_list` | `table` | No |  |
| `suspension_reason` | `string` | No |  |
| `tag` | `table` | No |  |
| `time_zone_id` | `string` | No |  |
| `trading_name` | `string` | No |  |
| `web_hook_limit` | `number` | No |  |
| `your_role_name` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsMerchantPage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsMerchantPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity

```lua
local no_frixion_money_moov_models_merchant_pay_by_bank_setting = client:NoFrixionMoneyMoovModelsMerchantPayByBankSetting(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bank_country_code` | `table` | No |  |
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsMerchantPayByBankSetting():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantTokenEntity

```lua
local no_frixion_money_moov_models_merchant_token = client:NoFrixionMoneyMoovModelsMerchantToken(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authentication_method` | `table` | No |  |
| `authorisation` | `table` | No |  |
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
| `permission_type` | `table` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovModelsMerchantToken():create({
  nonce = --[[ string ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsMerchantToken():load({ id = "no_frixion_money_moov_models_merchant_token_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:NoFrixionMoneyMoovModelsMerchantToken():update({
  id = "no_frixion_money_moov_models_merchant_token_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantTokenPageEntity

```lua
local no_frixion_money_moov_models_merchant_token_page = client:NoFrixionMoneyMoovModelsMerchantTokenPage(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authentication_method` | `table` | No |  |
| `authorisation` | `table` | No |  |
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
| `permission_type` | `table` | No |  |
| `request_signature_version` | `number` | No |  |
| `shared_secret_algorithm` | `string` | No |  |
| `shared_secret_base64` | `string` | No |  |
| `token` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsMerchantTokenPage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsNoFrixionVersionEntity

```lua
local no_frixion_money_moov_models_no_frixion_version = client:NoFrixionMoneyMoovModelsNoFrixionVersion(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `build_version` | `number` | No |  |
| `major_version` | `number` | No |  |
| `minor_version` | `number` | No |  |
| `release_name` | `string` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsNoFrixionVersion():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsNoFrixionVersionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsOpenBankingAccountEntity

```lua
local no_frixion_money_moov_models_open_banking_account = client:NoFrixionMoneyMoovModelsOpenBankingAccount(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_balance` | `table` | No |  |
| `account_identification` | `table` | No |  |
| `account_name` | `table` | No |  |
| `account_type` | `string` | No |  |
| `balance` | `number` | No |  |
| `consolidated_account_information` | `table` | No |  |
| `currency` | `string` | No |  |
| `description` | `string` | No |  |
| `detail` | `string` | No |  |
| `id` | `string` | No |  |
| `nickname` | `string` | No |  |
| `type` | `string` | No |  |
| `usage_type` | `string` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsOpenBankingAccount():load({ id = "no_frixion_money_moov_models_open_banking_account_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsOpenBankingAccountEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsOpenBankingConsentEntity

```lua
local no_frixion_money_moov_models_open_banking_consent = client:NoFrixionMoneyMoovModelsOpenBankingConsent(nil)
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovModelsOpenBankingConsent():create({
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsOpenBankingConsent():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsOpenBankingConsent():load({ id = "no_frixion_money_moov_models_open_banking_consent_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsOpenBankingConsent():remove({ id = "no_frixion_money_moov_models_open_banking_consent_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:NoFrixionMoneyMoovModelsOpenBankingConsent():update({
  id = "no_frixion_money_moov_models_open_banking_consent_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsOpenBankingConsentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsOpenBankingTransactionEntity

```lua
local no_frixion_money_moov_models_open_banking_transaction = client:NoFrixionMoneyMoovModelsOpenBankingTransaction(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address_detail` | `table` | No |  |
| `amount` | `number` | No |  |
| `balance` | `table` | No |  |
| `booking_date_time` | `string` | No |  |
| `charge_detail` | `table` | No |  |
| `currency` | `string` | No |  |
| `currency_exchange` | `table` | No |  |
| `date` | `string` | No |  |
| `description` | `string` | No |  |
| `enrichment` | `table` | No |  |
| `gross_amount` | `table` | Yes |  |
| `id` | `string` | No |  |
| `iso_bank_transaction_code` | `table` | No |  |
| `merchant` | `table` | No |  |
| `payee_detail` | `table` | Yes |  |
| `payer_detail` | `table` | Yes |  |
| `proprietary_bank_transaction_code` | `table` | No |  |
| `reference` | `string` | No |  |
| `statement_reference` | `table` | No |  |
| `status` | `string` | No |  |
| `supplementary_data` | `any` | No |  |
| `transaction_amount` | `table` | Yes |  |
| `transaction_information` | `table` | No |  |
| `transaction_mutability` | `string` | No |  |
| `value_date_time` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsOpenBankingTransaction():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsOpenBankingTransactionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentEntity

```lua
local no_frixion_money_moov_models_payment = client:NoFrixionMoneyMoovModelsPayment(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `table` | No |  |
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
| `created_by_user` | `table` | Yes |  |
| `currency` | `string` | No |  |
| `custom_field` | `table` | No |  |
| `customer_email_address` | `string` | No |  |
| `customer_id` | `string` | No |  |
| `customer_name` | `string` | No |  |
| `description` | `string` | No |  |
| `destination_account` | `table` | No |  |
| `direct_debit_payment` | `table` | No |  |
| `due_date` | `string` | No |  |
| `event` | `table` | No |  |
| `failure_callback_url` | `string` | No |  |
| `field_display_setting` | `table` | No |  |
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
| `notification_role_i_d` | `table` | No |  |
| `order_id` | `string` | No |  |
| `partial_payment_method` | `string` | No |  |
| `partial_payment_step` | `string` | No |  |
| `payment_attempt` | `table` | No |  |
| `payment_method` | `table` | No |  |
| `payment_processor` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `pisp_account_id` | `string` | No |  |
| `priority_bank_id` | `string` | No |  |
| `result` | `table` | No |  |
| `sandbox_settle_delay_in_second` | `number` | No |  |
| `shipping_address` | `table` | No |  |
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
| `tag` | `table` | No |  |
| `tag_id` | `table` | No |  |
| `title` | `string` | No |  |
| `tokenised_card` | `table` | No |  |
| `transaction` | `table` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovModelsPayment():create({
  created_by_user = --[[ table ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsPayment():load({ id = "no_frixion_money_moov_models_payment_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:NoFrixionMoneyMoovModelsPayment():update({
  id = "no_frixion_money_moov_models_payment_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity

```lua
local no_frixion_money_moov_models_payment_account_minimal_page = client:NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(nil)
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
| `identifier` | `table` | Yes |  |
| `is_archived` | `boolean` | No |  |
| `is_connected_account` | `boolean` | No |  |
| `merchant_id` | `string` | No |  |
| `submitted_payouts_balance` | `number` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsPaymentAccountMinimalPage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentAccountPageEntity

```lua
local no_frixion_money_moov_models_payment_account_page = client:NoFrixionMoneyMoovModelsPaymentAccountPage(nil)
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
| `created_by` | `table` | Yes |  |
| `created_by_display_name` | `string` | No |  |
| `currency` | `string` | No |  |
| `default_payment_rail` | `string` | No |  |
| `display_name` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `external_account_icon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `table` | Yes |  |
| `inserted` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_connected_account` | `boolean` | No |  |
| `is_default` | `boolean` | No |  |
| `is_trust_account` | `boolean` | No |  |
| `is_virtual` | `boolean` | No |  |
| `last_transaction` | `table` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `physical_account_id` | `string` | No |  |
| `rule` | `table` | No |  |
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsPaymentAccountPage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentInitiationEntity

```lua
local no_frixion_money_moov_models_payment_initiation = client:NoFrixionMoneyMoovModelsPaymentInitiation(nil)
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovModelsPaymentInitiation():create({
  paymentrequest_id = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentInitiationEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestEventEntity

```lua
local no_frixion_money_moov_models_payment_request_event = client:NoFrixionMoneyMoovModelsPaymentRequestEvent(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsPaymentRequestEvent():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestEventEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestMetricEntity

```lua
local no_frixion_money_moov_models_payment_request_metric = client:NoFrixionMoneyMoovModelsPaymentRequestMetric(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `all` | `number` | No |  |
| `authorized` | `number` | No |  |
| `paid` | `number` | No |  |
| `partially_paid` | `number` | No |  |
| `total_amounts_by_currency` | `table` | No |  |
| `unpaid` | `number` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsPaymentRequestMetric():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMetricEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity

```lua
local no_frixion_money_moov_models_payment_request_minimal = client:NoFrixionMoneyMoovModelsPaymentRequestMinimal(nil)
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
| `custom_fields_to_display` | `table` | No |  |
| `description` | `string` | No |  |
| `due_date` | `string` | No |  |
| `field_display_setting` | `table` | No |  |
| `google_pay_merchant_id` | `string` | No |  |
| `id` | `string` | No |  |
| `jwk` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_logo_url_png` | `string` | No |  |
| `merchant_logo_url_svg` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `merchant_short_name` | `string` | No |  |
| `partial_payment_method` | `string` | No |  |
| `payment_attempt` | `table` | No |  |
| `payment_methods_list` | `table` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_processor_key` | `string` | No |  |
| `pisp_error` | `string` | No |  |
| `priority_bank_id` | `string` | No |  |
| `status` | `string` | No |  |
| `stripe_account_id` | `string` | No |  |
| `title` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsPaymentRequestMinimal():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestResultEntity

```lua
local no_frixion_money_moov_models_payment_request_result = client:NoFrixionMoneyMoovModelsPaymentRequestResult(nil)
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
| `payment` | `table` | No |  |
| `payment_request_id` | `string` | No |  |
| `pisp_authorization` | `table` | No |  |
| `requested_amount` | `number` | No |  |
| `result` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsPaymentRequestResult():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestResultEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity

```lua
local no_frixion_money_moov_models_payment_requests_merchant_payment = client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(nil)
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
| `template` | `table` | Yes |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity

```lua
local no_frixion_money_moov_models_payment_requests_merchant_payment2 = client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(nil)
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
| `template` | `table` | Yes |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2():load({ paymentrequest_id = "paymentrequest_id", template_id = "template_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity

```lua
local no_frixion_money_moov_models_payment_requests_merchant_payment3 = client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(nil)
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
| `template` | `table` | Yes |  |

### Operations

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3():update({
  paymentrequest_id = "paymentrequest_id",
  template_id = "template_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity

```lua
local no_frixion_money_moov_models_payment_requests_merchant_payment4 = client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(nil)
```

### Operations

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4():remove({ paymentrequest_id = "paymentrequest_id", template_id = "template_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayoutKeysetPageEntity

```lua
local no_frixion_money_moov_models_payout_keyset_page = client:NoFrixionMoneyMoovModelsPayoutKeysetPage(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `table` | No |  |
| `authorisation` | `table` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `table` | Yes |  |
| `can_authorise` | `boolean` | No |  |
| `can_process` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `table` | No |  |
| `document` | `table` | No |  |
| `event` | `table` | No |  |
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
| `rule` | `table` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `table` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `table` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `number` | No |  |
| `transacted_fx_amount` | `number` | No |  |
| `transacted_fx_rate` | `number` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsPayoutKeysetPage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPayoutKeysetPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayoutMetricEntity

```lua
local no_frixion_money_moov_models_payout_metric = client:NoFrixionMoneyMoovModelsPayoutMetric(nil)
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
| `total_amounts_by_currency` | `table` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsPayoutMetric():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPayoutMetricEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity

```lua
local no_frixion_money_moov_models_payouts_payouts_create = client:NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_payout` | `table` | No |  |
| `payout` | `table` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovModelsPayoutsPayoutsCreate():create({
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayrunEntity

```lua
local no_frixion_money_moov_models_payrun = client:NoFrixionMoneyMoovModelsPayrun(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation` | `table` | No |  |
| `authorisation_date` | `string` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_delete` | `boolean` | No |  |
| `can_edit` | `boolean` | No |  |
| `event` | `table` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice` | `table` | No |  |
| `invoices_minimal` | `table` | No |  |
| `is_archived` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `last_updated_by` | `table` | Yes |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment` | `table` | No |  |
| `payout` | `table` | No |  |
| `payouts_count` | `number` | No |  |
| `reason` | `string` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled_date` | `string` | No |  |
| `source_account` | `table` | No |  |
| `status` | `string` | No |  |
| `total_eur` | `number` | No |  |
| `total_gbp` | `number` | No |  |
| `total_usd` | `number` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovModelsPayrun():create({
  id = --[[ string ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsPayrun():load({ id = "no_frixion_money_moov_models_payrun_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:NoFrixionMoneyMoovModelsPayrun():update({
  id = "no_frixion_money_moov_models_payrun_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsPayrunEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsReportResultEntity

```lua
local no_frixion_money_moov_models_report_result = client:NoFrixionMoneyMoovModelsReportResult(nil)
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

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsReportResult():load({ id = 1, report_id = "report_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsReportResultEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsRuleEntity

```lua
local no_frixion_money_moov_models_rule = client:NoFrixionMoneyMoovModelsRule(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account` | `table` | No |  |
| `account_id` | `string` | No |  |
| `approve_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `table` | No |  |
| `authorisation` | `table` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `can_authorise` | `boolean` | No |  |
| `created_by` | `table` | Yes |  |
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
| `sweep_action` | `table` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovModelsRule():create({
  created_by = --[[ table ]],
  nonce = --[[ string ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsRule():load({ id = "no_frixion_money_moov_models_rule_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:NoFrixionMoneyMoovModelsRule():update({
  id = "no_frixion_money_moov_models_rule_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsRuleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsTransactionEntity

```lua
local no_frixion_money_moov_models_transaction = client:NoFrixionMoneyMoovModelsTransaction(nil)
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
| `counterparty` | `table` | No |  |
| `counterparty_summary` | `string` | No |  |
| `currency` | `string` | No |  |
| `description` | `string` | No |  |
| `fx_amount` | `number` | No |  |
| `fx_currency` | `string` | No |  |
| `fx_rate` | `number` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `payment_request_custom_field` | `table` | No |  |
| `payment_request_id` | `string` | No |  |
| `payout_id` | `string` | No |  |
| `raw_reference` | `string` | No |  |
| `rule_id` | `string` | No |  |
| `tag` | `table` | No |  |
| `their_reference` | `string` | No |  |
| `transaction_date` | `string` | No |  |
| `type` | `string` | No |  |
| `virtual_iban` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsTransaction():load({ id = "no_frixion_money_moov_models_transaction_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsTransactionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsTransactionPageEntity

```lua
local no_frixion_money_moov_models_transaction_page = client:NoFrixionMoneyMoovModelsTransactionPage(nil)
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
| `content` | `table` | No |  |
| `counterparty` | `table` | No |  |
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
| `payment_request_custom_field` | `table` | No |  |
| `payment_request_id` | `string` | No |  |
| `payout_id` | `string` | No |  |
| `raw_reference` | `string` | No |  |
| `rule_id` | `string` | No |  |
| `tag` | `table` | No |  |
| `their_reference` | `string` | No |  |
| `total_page` | `number` | No |  |
| `total_size` | `number` | No |  |
| `transaction_date` | `string` | No |  |
| `type` | `string` | No |  |
| `virtual_iban` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsTransactionPage():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsTransactionPage():load({ account_id = "account_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsTransactionPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsUserInviteEntity

```lua
local no_frixion_money_moov_models_user_invite = client:NoFrixionMoneyMoovModelsUserInvite(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_status` | `table` | No |  |
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
| `user` | `table` | Yes |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovModelsUserInvite():create({
  user = --[[ table ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsUserInvite():load({ id = "no_frixion_money_moov_models_user_invite_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsUserInviteEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsUserInvitePageEntity

```lua
local no_frixion_money_moov_models_user_invite_page = client:NoFrixionMoneyMoovModelsUserInvitePage(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_status` | `table` | No |  |
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
| `user` | `table` | Yes |  |
| `user_id` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsUserInvitePage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsUserInvitePageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsUserPageEntity

```lua
local no_frixion_money_moov_models_user_page = client:NoFrixionMoneyMoovModelsUserPage(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `table` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `boolean` | No |  |
| `permission` | `table` | No |  |
| `roles_with_scope` | `table` | No |  |
| `two_factor_enabled` | `boolean` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsUserPage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsUserPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsWebhookEntity

```lua
local no_frixion_money_moov_models_webhook = client:NoFrixionMoneyMoovModelsWebhook(nil)
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
| `resource_type` | `table` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:NoFrixionMoneyMoovModelsWebhook():create({
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:NoFrixionMoneyMoovModelsWebhook():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionMoneyMoovModelsWebhook():load({ id = "no_frixion_money_moov_models_webhook_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:NoFrixionMoneyMoovModelsWebhook():update({
  id = "no_frixion_money_moov_models_webhook_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionMoneyMoovModelsWebhookEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## OpenBankingEntity

```lua
local open_banking = client:OpenBanking(nil)
```

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:OpenBanking():create({
  account_id = --[[ string ]],
})
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:OpenBanking():remove()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OpenBankingEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PayeeverificationEntity

```lua
local payeeverification = client:Payeeverification(nil)
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Payeeverification():create({
  account_name = --[[ string ]],
  iban = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PayeeverificationEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PaymentRequestEntity

```lua
local payment_request = client:PaymentRequest(nil)
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PaymentRequest():create({
  paymentrequest_id = --[[ string ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:PaymentRequest():load()
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:PaymentRequest():remove({ id = "id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:PaymentRequest():update({
  paymentrequest_id = "paymentrequest_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentRequestEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PayoutEntity

```lua
local payout = client:Payout(nil)
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
| `authentication_method` | `table` | No |  |
| `authorisation` | `table` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `table` | Yes |  |
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
| `destination` | `table` | No |  |
| `document` | `table` | No |  |
| `event` | `table` | No |  |
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
| `rule` | `table` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `table` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `table` | No |  |
| `tag_id` | `table` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Payout():create({
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Payout():load({ id = "payout_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Payout():remove({ id = "payout_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Payout():update({
  id = "payout_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PayoutEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PayrunEntity

```lua
local payrun = client:Payrun(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |
| `note` | `string` | No |  |
| `scheduled_date` | `string` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Payrun():create({
  id = --[[ string ]],
})
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Payrun():remove({ id = "id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Payrun():update({
  id = "id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PayrunEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## RejectEntity

```lua
local reject = client:Reject(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `table` | No |  |
| `authorisation` | `table` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `table` | Yes |  |
| `can_authorise` | `boolean` | No |  |
| `can_process` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `table` | No |  |
| `document` | `table` | No |  |
| `event` | `table` | No |  |
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
| `rule` | `table` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `table` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `table` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `number` | No |  |
| `transacted_fx_amount` | `number` | No |  |
| `transacted_fx_rate` | `number` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Reject():update({
  id = "id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RejectEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ReportEntity

```lua
local report = client:Report(nil)
```

### Operations

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Report():update({
  id = "id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ReportEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## RuleEntity

```lua
local rule = client:Rule(nil)
```

### Operations

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Rule():remove({ id = "id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Rule():update({
  id = "id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RuleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SendEntity

```lua
local send = client:Send(nil)
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
| `authentication_method` | `table` | No |  |
| `authorisation` | `table` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `table` | Yes |  |
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
| `destination` | `table` | No |  |
| `document` | `table` | No |  |
| `event` | `table` | No |  |
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
| `rule` | `table` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `table` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `table` | No |  |
| `tag_id` | `table` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Send():create({
  beneficiary = --[[ table ]],
  source_account_identifier = --[[ table ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SendEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SendbeneficiaryEntity

```lua
local sendbeneficiary = client:Sendbeneficiary(nil)
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
| `authentication_method` | `table` | No |  |
| `authorisation` | `table` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `table` | Yes |  |
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
| `destination` | `table` | No |  |
| `document` | `table` | No |  |
| `event` | `table` | No |  |
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
| `rule` | `table` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `table` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `table` | No |  |
| `tag_id` | `table` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Sendbeneficiary():create({
  beneficiary = --[[ table ]],
  source_account_identifier = --[[ table ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SendbeneficiaryEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## TagEntity

```lua
local tag = client:Tag(nil)
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Tag():create({
  merchant_id = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Tag():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TagEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## TokenEntity

```lua
local token = client:Token(nil)
```

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Token():create({
  id = --[[ string ]],
})
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Token():remove({ id = "id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TokenEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## TransactionEntity

```lua
local transaction = client:Transaction(nil)
```

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Transaction():create({
  id = --[[ string ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Transaction():load({ sequence_number = 1, transaction_id = "transaction_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Transaction():remove({ id = "id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TransactionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## UserEntity

```lua
local user = client:User(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `table` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `boolean` | No |  |
| `permission` | `table` | No |  |
| `profile` | `string` | No |  |
| `roles_with_scope` | `table` | No |  |
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:User():list()
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:User():update({
  id = "id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UserEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## UserInviteEntity

```lua
local user_invite = client:UserInvite(nil)
```

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:UserInvite():create({
  id = --[[ string ]],
})
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:UserInvite():remove({ id = "id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:UserInvite():update({
  id = "id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UserInviteEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## VirtualEntity

```lua
local virtual = client:Virtual(nil)
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
| `created_by` | `table` | Yes |  |
| `created_by_display_name` | `string` | No |  |
| `currency` | `string` | No |  |
| `default_payment_rail` | `string` | No |  |
| `display_name` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `external_account_icon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `table` | Yes |  |
| `inserted` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_connected_account` | `boolean` | No |  |
| `is_default` | `boolean` | No |  |
| `is_trust_account` | `boolean` | No |  |
| `is_virtual` | `boolean` | No |  |
| `last_transaction` | `table` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `name` | `string` | Yes |  |
| `physical_account_id` | `string` | No |  |
| `rule` | `table` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Virtual():create({
  account_id = --[[ string ]],
})
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Virtual():update({
  account_id = "account_id",
  id = "id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VirtualEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## WebhookEntity

```lua
local webhook = client:Webhook(nil)
```

### Operations

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Webhook():remove({ id = "id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WebhookEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## WhoamiEntity

```lua
local whoami = client:Whoami(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `table` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `boolean` | No |  |
| `permission` | `table` | No |  |
| `roles_with_scope` | `table` | No |  |
| `two_factor_enabled` | `boolean` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Whoami():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WhoamiEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## WhoamitrustedappEntity

```lua
local whoamitrustedapp = client:Whoamitrustedapp(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `table` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `boolean` | No |  |
| `permission` | `table` | No |  |
| `roles_with_scope` | `table` | No |  |
| `two_factor_enabled` | `boolean` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Whoamitrustedapp():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WhoamitrustedappEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
  },
})
```

