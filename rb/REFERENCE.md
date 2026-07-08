# Nofrixion Ruby SDK Reference

Complete API reference for the Nofrixion Ruby SDK.


## NofrixionSDK

### Constructor

```ruby
require_relative 'Nofrixion_sdk'

client = NofrixionSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["apikey"]` | `String` | API key for authentication. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NofrixionSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = NofrixionSDK.test
```


### Instance Methods

#### `Account(data = nil)`

Create a new `Account` entity instance. Pass `nil` for no initial data.

#### `Beneficiary(data = nil)`

Create a new `Beneficiary` entity instance. Pass `nil` for no initial data.

#### `Cancel(data = nil)`

Create a new `Cancel` entity instance. Pass `nil` for no initial data.

#### `Disable(data = nil)`

Create a new `Disable` entity instance. Pass `nil` for no initial data.

#### `Enable(data = nil)`

Create a new `Enable` entity instance. Pass `nil` for no initial data.

#### `Merchant(data = nil)`

Create a new `Merchant` entity instance. Pass `nil` for no initial data.

#### `Metadata(data = nil)`

Create a new `Metadata` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(data = nil)`

Create a new `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingPaymentRequestPage(data = nil)`

Create a new `NoFrixionBizBizModelsPagingPaymentRequestPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingPayoutPage(data = nil)`

Create a new `NoFrixionBizBizModelsPagingPayoutPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingPayrunPage(data = nil)`

Create a new `NoFrixionBizBizModelsPagingPayrunPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingRuleEventsPage(data = nil)`

Create a new `NoFrixionBizBizModelsPagingRuleEventsPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPagingRulesPage(data = nil)`

Create a new `NoFrixionBizBizModelsPagingRulesPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPaymentsCardPayment(data = nil)`

Create a new `NoFrixionBizBizModelsPaymentsCardPayment` entity instance. Pass `nil` for no initial data.

#### `NoFrixionBizBizModelsPaymentsCardPublicKey(data = nil)`

Create a new `NoFrixionBizBizModelsPaymentsCardPublicKey` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(data = nil)`

Create a new `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(data = nil)`

Create a new `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(data = nil)`

Create a new `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(data = nil)`

Create a new `NoFrixionMoneyMoovApiFeaturesUserInvitesCreate` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(data = nil)`

Create a new `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsBatchPayout(data = nil)`

Create a new `NoFrixionMoneyMoovModelsBatchPayout` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsBeneficiaryGroupPage(data = nil)`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryGroupPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsBeneficiaryPage(data = nil)`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsCardCustomerToken(data = nil)`

Create a new `NoFrixionMoneyMoovModelsCardCustomerToken` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(data = nil)`

Create a new `NoFrixionMoneyMoovModelsCurrencyCurrencyInfo` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(data = nil)`

Create a new `NoFrixionMoneyMoovModelsDirectDebitBatchSubmit` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsFxRate(data = nil)`

Create a new `NoFrixionMoneyMoovModelsFxRate` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsIPayment(data = nil)`

Create a new `NoFrixionMoneyMoovModelsIPayment` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMandatesMandate(data = nil)`

Create a new `NoFrixionMoneyMoovModelsMandatesMandate` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchant(data = nil)`

Create a new `NoFrixionMoneyMoovModelsMerchant` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantPage(data = nil)`

Create a new `NoFrixionMoneyMoovModelsMerchantPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantPayByBankSetting(data = nil)`

Create a new `NoFrixionMoneyMoovModelsMerchantPayByBankSetting` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantToken(data = nil)`

Create a new `NoFrixionMoneyMoovModelsMerchantToken` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantTokenPage(data = nil)`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsNoFrixionVersion(data = nil)`

Create a new `NoFrixionMoneyMoovModelsNoFrixionVersion` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsOpenBankingAccount(data = nil)`

Create a new `NoFrixionMoneyMoovModelsOpenBankingAccount` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsOpenBankingConsent(data = nil)`

Create a new `NoFrixionMoneyMoovModelsOpenBankingConsent` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsOpenBankingTransaction(data = nil)`

Create a new `NoFrixionMoneyMoovModelsOpenBankingTransaction` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPayment(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPayment` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountMinimalPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentAccountPage(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentInitiation(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPaymentInitiation` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestEvent(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestEvent` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestMetric(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMetric` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestMinimal(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMinimal` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestResult(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestResult` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPayoutKeysetPage(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPayoutKeysetPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPayoutMetric(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPayoutMetric` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPayoutsPayoutsCreate` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsPayrun(data = nil)`

Create a new `NoFrixionMoneyMoovModelsPayrun` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsReportResult(data = nil)`

Create a new `NoFrixionMoneyMoovModelsReportResult` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsRule(data = nil)`

Create a new `NoFrixionMoneyMoovModelsRule` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsTransaction(data = nil)`

Create a new `NoFrixionMoneyMoovModelsTransaction` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsTransactionPage(data = nil)`

Create a new `NoFrixionMoneyMoovModelsTransactionPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsUserInvite(data = nil)`

Create a new `NoFrixionMoneyMoovModelsUserInvite` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsUserInvitePage(data = nil)`

Create a new `NoFrixionMoneyMoovModelsUserInvitePage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsUserPage(data = nil)`

Create a new `NoFrixionMoneyMoovModelsUserPage` entity instance. Pass `nil` for no initial data.

#### `NoFrixionMoneyMoovModelsWebhook(data = nil)`

Create a new `NoFrixionMoneyMoovModelsWebhook` entity instance. Pass `nil` for no initial data.

#### `OpenBanking(data = nil)`

Create a new `OpenBanking` entity instance. Pass `nil` for no initial data.

#### `Payeeverification(data = nil)`

Create a new `Payeeverification` entity instance. Pass `nil` for no initial data.

#### `PaymentRequest(data = nil)`

Create a new `PaymentRequest` entity instance. Pass `nil` for no initial data.

#### `Payout(data = nil)`

Create a new `Payout` entity instance. Pass `nil` for no initial data.

#### `Payrun(data = nil)`

Create a new `Payrun` entity instance. Pass `nil` for no initial data.

#### `Reject(data = nil)`

Create a new `Reject` entity instance. Pass `nil` for no initial data.

#### `Report(data = nil)`

Create a new `Report` entity instance. Pass `nil` for no initial data.

#### `Rule(data = nil)`

Create a new `Rule` entity instance. Pass `nil` for no initial data.

#### `Send(data = nil)`

Create a new `Send` entity instance. Pass `nil` for no initial data.

#### `Sendbeneficiary(data = nil)`

Create a new `Sendbeneficiary` entity instance. Pass `nil` for no initial data.

#### `Tag(data = nil)`

Create a new `Tag` entity instance. Pass `nil` for no initial data.

#### `Token(data = nil)`

Create a new `Token` entity instance. Pass `nil` for no initial data.

#### `Transaction(data = nil)`

Create a new `Transaction` entity instance. Pass `nil` for no initial data.

#### `User(data = nil)`

Create a new `User` entity instance. Pass `nil` for no initial data.

#### `UserInvite(data = nil)`

Create a new `UserInvite` entity instance. Pass `nil` for no initial data.

#### `Virtual(data = nil)`

Create a new `Virtual` entity instance. Pass `nil` for no initial data.

#### `Webhook(data = nil)`

Create a new `Webhook` entity instance. Pass `nil` for no initial data.

#### `Whoami(data = nil)`

Create a new `Whoami` entity instance. Pass `nil` for no initial data.

#### `Whoamitrustedapp(data = nil)`

Create a new `Whoamitrustedapp` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## AccountEntity

```ruby
account = client.Account
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `String` | No |  |
| `account_name` | `String` | No |  |
| `account_supplier_name` | `String` | No |  |
| `account_type` | `String` | No |  |
| `available_balance` | `Float` | No |  |
| `available_balance_minor_unit` | `Integer` | No |  |
| `balance` | `Float` | No |  |
| `balance_minor_unit` | `Integer` | No |  |
| `bank_name` | `String` | No |  |
| `consent_id` | `String` | No |  |
| `created_by` | `Hash` | Yes |  |
| `created_by_display_name` | `String` | No |  |
| `currency` | `String` | No |  |
| `default_payment_rail` | `String` | No |  |
| `display_name` | `String` | No |  |
| `expiry_date` | `String` | No |  |
| `external_account_icon` | `String` | No |  |
| `format` | `String` | No |  |
| `from_date` | `String` | No |  |
| `id` | `String` | No |  |
| `identifier` | `Hash` | Yes |  |
| `inserted` | `String` | No |  |
| `is_archived` | `Boolean` | No |  |
| `is_connected_account` | `Boolean` | No |  |
| `is_default` | `Boolean` | No |  |
| `is_trust_account` | `Boolean` | No |  |
| `is_virtual` | `Boolean` | No |  |
| `last_transaction` | `Hash` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `merchant_name` | `String` | No |  |
| `physical_account_id` | `String` | No |  |
| `role_i_d` | `Array` | No |  |
| `rule` | `Array` | No |  |
| `submitted_payouts_balance` | `Float` | No |  |
| `submitted_payouts_balance_minor_unit` | `Integer` | No |  |
| `summary` | `String` | No |  |
| `supplier_physical_account_id` | `String` | No |  |
| `supplier_sepa_instant_status` | `String` | No |  |
| `to_date` | `String` | No |  |
| `xero_bank_feed_connection_status` | `String` | No |  |
| `xero_bank_feed_last_synced_at` | `String` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `String` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `String` | No |  |
| `xero_bank_feed_sync_status` | `String` | No |  |
| `xero_unsynchronised_transactions_count` | `Integer` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Account.create({
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Account.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Account.load({ "id" => "account_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Account.remove({ "id" => "account_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Account.update({
  "id" => "account_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `AccountEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## BeneficiaryEntity

```ruby
beneficiary = client.Beneficiary
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `String` | No |  |
| `authentication_method` | `Array` | No |  |
| `authorisation` | `Array` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `beneficiary_event` | `Array` | No |  |
| `can_authorise` | `Boolean` | No |  |
| `can_update` | `Boolean` | No |  |
| `created_by` | `Hash` | Yes |  |
| `created_by_email_address` | `String` | No |  |
| `currency` | `String` | Yes |  |
| `destination` | `Hash` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `is_enabled` | `Boolean` | No |  |
| `last_authorised` | `String` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `name` | `String` | Yes |  |
| `nonce` | `String` | No |  |
| `source_account` | `Array` | No |  |
| `source_account_i_d` | `Array` | No |  |
| `their_reference` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Beneficiary.create({
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Beneficiary.load({ "id" => "beneficiary_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Beneficiary.remove({ "id" => "beneficiary_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Beneficiary.update({
  "id" => "beneficiary_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `BeneficiaryEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CancelEntity

```ruby
cancel = client.Cancel
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `String` | No |  |
| `amount` | `Float` | No |  |
| `amount_minor_unit` | `Integer` | No |  |
| `approve_payout_url` | `String` | No |  |
| `approver_id` | `String` | No |  |
| `authentication_method` | `Array` | No |  |
| `authorisation` | `Array` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `batch_payout_id` | `String` | No |  |
| `beneficiary` | `Hash` | Yes |  |
| `can_authorise` | `Boolean` | No |  |
| `can_process` | `Boolean` | No |  |
| `can_update` | `Boolean` | No |  |
| `charge_bearer` | `String` | No |  |
| `created_by` | `String` | No |  |
| `created_by_email_address` | `String` | No |  |
| `currency` | `String` | No |  |
| `current_user_id` | `String` | No |  |
| `description` | `String` | No |  |
| `destination` | `Hash` | No |  |
| `document` | `Array` | No |  |
| `event` | `Array` | No |  |
| `formatted_amount` | `String` | No |  |
| `formatted_fx_destination_amount` | `String` | No |  |
| `formatted_schedule` | `String` | No |  |
| `formatted_schedule_day_only` | `String` | No |  |
| `formatted_source_account_available_balance` | `String` | No |  |
| `fx_destination_amount` | `Float` | No |  |
| `fx_destination_amount_minor_unit` | `Integer` | No |  |
| `fx_destination_currency` | `String` | No |  |
| `fx_quote_expires_at` | `String` | No |  |
| `fx_quote_id` | `String` | No |  |
| `fx_rate` | `Float` | No |  |
| `fx_use_destination_amount` | `Boolean` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `invoice_id` | `String` | No |  |
| `is_archived` | `Boolean` | No |  |
| `is_failed` | `Boolean` | No |  |
| `is_settled` | `Boolean` | No |  |
| `is_submitted` | `Boolean` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `merchant_token_description` | `String` | No |  |
| `nonce` | `String` | No |  |
| `payment_processor` | `String` | No |  |
| `payment_rail` | `String` | No |  |
| `payrun_id` | `String` | No |  |
| `payrun_name` | `String` | No |  |
| `rule` | `Hash` | No |  |
| `schedule_date` | `String` | No |  |
| `scheduled` | `Boolean` | No |  |
| `source_account_available_balance` | `Float` | No |  |
| `source_account_available_balance_minor_unit` | `Integer` | No |  |
| `source_account_bic` | `String` | No |  |
| `source_account_currency` | `String` | No |  |
| `source_account_iban` | `String` | No |  |
| `source_account_identifier` | `Hash` | Yes |  |
| `source_account_name` | `String` | No |  |
| `source_account_number` | `String` | No |  |
| `source_account_sortcode` | `String` | No |  |
| `status` | `String` | No |  |
| `tag` | `Array` | No |  |
| `their_reference` | `String` | No |  |
| `topup_payrun_id` | `String` | No |  |
| `transacted_amount` | `Float` | No |  |
| `transacted_fx_amount` | `Float` | No |  |
| `transacted_fx_rate` | `Float` | No |  |
| `type` | `String` | No |  |
| `user_id` | `String` | No |  |
| `your_reference` | `String` | No |  |

### Operations

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Cancel.update({
  "id" => "id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `CancelEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## DisableEntity

```ruby
disable = client.Disable
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `String` | No |  |
| `authentication_method` | `Array` | No |  |
| `authorisation` | `Array` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `beneficiary_event` | `Array` | No |  |
| `can_authorise` | `Boolean` | No |  |
| `can_update` | `Boolean` | No |  |
| `created_by` | `Hash` | Yes |  |
| `created_by_email_address` | `String` | No |  |
| `currency` | `String` | Yes |  |
| `destination` | `Hash` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `is_enabled` | `Boolean` | No |  |
| `last_authorised` | `String` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `name` | `String` | Yes |  |
| `nonce` | `String` | No |  |
| `source_account` | `Array` | No |  |
| `their_reference` | `String` | No |  |

### Operations

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Disable.update({
  "id" => "id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `DisableEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## EnableEntity

```ruby
enable = client.Enable
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `String` | No |  |
| `authentication_method` | `Array` | No |  |
| `authorisation` | `Array` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `beneficiary_event` | `Array` | No |  |
| `can_authorise` | `Boolean` | No |  |
| `can_update` | `Boolean` | No |  |
| `created_by` | `Hash` | Yes |  |
| `created_by_email_address` | `String` | No |  |
| `currency` | `String` | Yes |  |
| `destination` | `Hash` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `is_enabled` | `Boolean` | No |  |
| `last_authorised` | `String` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `name` | `String` | Yes |  |
| `nonce` | `String` | No |  |
| `source_account` | `Array` | No |  |
| `their_reference` | `String` | No |  |

### Operations

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Enable.update({
  "id" => "id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `EnableEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## MerchantEntity

```ruby
merchant = client.Merchant
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `reason` | `String` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Merchant.load({ "merchant_id" => "merchant_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Merchant.remove()
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Merchant.update({
  "id" => "id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `MerchantEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## MetadataEntity

```ruby
metadata = client.Metadata
```

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Metadata.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `MetadataEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity

```ruby
no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page = client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approved_at` | `String` | No |  |
| `currency` | `String` | No |  |
| `customer_account_number` | `String` | No |  |
| `customer_city` | `String` | No |  |
| `customer_country_code` | `String` | No |  |
| `customer_country_name` | `String` | No |  |
| `customer_email_address` | `String` | No |  |
| `customer_first_name` | `String` | No |  |
| `customer_iban` | `String` | No |  |
| `customer_last_name` | `String` | No |  |
| `customer_sort_code` | `String` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `is_recurring` | `Boolean` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `reference` | `String` | No |  |
| `status` | `String` | No |  |
| `supplier_bank_account_id` | `String` | No |  |
| `supplier_customer_id` | `String` | No |  |
| `supplier_mandate_id` | `String` | No |  |
| `supplier_name` | `String` | No |  |
| `supplier_status` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionBizBizModelsPagingPaymentRequestPageEntity

```ruby
no_frixion_biz_biz_models_paging_payment_request_page = client.NoFrixionBizBizModelsPagingPaymentRequestPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `Array` | No |  |
| `amount` | `Float` | No |  |
| `amount_pending` | `Float` | No |  |
| `amount_received` | `Float` | No |  |
| `amount_refunded` | `Float` | No |  |
| `auto_send_receipt` | `Boolean` | No |  |
| `base_origin_url` | `String` | No |  |
| `callback_url` | `String` | No |  |
| `card_authorize_only` | `Boolean` | No |  |
| `card_create_token` | `Boolean` | No |  |
| `card_create_token_mode` | `String` | No |  |
| `card_ignore_cvn` | `Boolean` | No |  |
| `card_processor_merchant_id` | `String` | No |  |
| `card_stripe_payment_intent_id` | `String` | No |  |
| `card_stripe_payment_intent_secret` | `String` | No |  |
| `created_by_user` | `Hash` | Yes |  |
| `currency` | `String` | No |  |
| `custom_field` | `Array` | No |  |
| `customer_email_address` | `String` | No |  |
| `customer_id` | `String` | No |  |
| `customer_name` | `String` | No |  |
| `description` | `String` | No |  |
| `destination_account` | `Hash` | No |  |
| `direct_debit_payment` | `Hash` | No |  |
| `due_date` | `String` | No |  |
| `event` | `Array` | No |  |
| `failure_callback_url` | `String` | No |  |
| `field_display_setting` | `Array` | No |  |
| `formatted_amount` | `String` | No |  |
| `hosted_pay_checkout_url` | `String` | No |  |
| `id` | `String` | No |  |
| `ignore_address_verification` | `Boolean` | No |  |
| `inserted` | `String` | No |  |
| `inserted_sortable` | `String` | No |  |
| `is_archived` | `Boolean` | No |  |
| `jwk` | `String` | No |  |
| `last_updated` | `String` | No |  |
| `lightning_invoice` | `String` | No |  |
| `lightning_invoice_expires_at` | `String` | No |  |
| `merchant_direct_debit_mandate_id` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `merchant_token_description` | `String` | No |  |
| `notification_email_address` | `String` | No |  |
| `notification_role_i_d` | `Array` | No |  |
| `order_id` | `String` | No |  |
| `partial_payment_method` | `String` | No |  |
| `partial_payment_step` | `String` | No |  |
| `payment_attempt` | `Array` | No |  |
| `payment_method` | `Array` | No |  |
| `payment_processor` | `String` | No |  |
| `payrun_id` | `String` | No |  |
| `pisp_account_id` | `String` | No |  |
| `priority_bank_id` | `String` | No |  |
| `result` | `Hash` | No |  |
| `sandbox_settle_delay_in_second` | `Integer` | No |  |
| `shipping_address` | `Hash` | No |  |
| `status` | `String` | No |  |
| `success_web_hook_url` | `String` | No |  |
| `tag` | `Array` | No |  |
| `title` | `String` | No |  |
| `tokenised_card` | `Array` | No |  |
| `transaction` | `Array` | No |  |
| `use_hosted_payment_page` | `Boolean` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionBizBizModelsPagingPaymentRequestPage.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionBizBizModelsPagingPaymentRequestPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionBizBizModelsPagingPayoutPageEntity

```ruby
no_frixion_biz_biz_models_paging_payout_page = client.NoFrixionBizBizModelsPagingPayoutPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `String` | No |  |
| `amount` | `Float` | No |  |
| `amount_minor_unit` | `Integer` | No |  |
| `approve_payout_url` | `String` | No |  |
| `approver_id` | `String` | No |  |
| `authentication_method` | `Array` | No |  |
| `authorisation` | `Array` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `batch_payout_id` | `String` | No |  |
| `beneficiary` | `Hash` | Yes |  |
| `can_authorise` | `Boolean` | No |  |
| `can_process` | `Boolean` | No |  |
| `can_update` | `Boolean` | No |  |
| `charge_bearer` | `String` | No |  |
| `created_by` | `String` | No |  |
| `created_by_email_address` | `String` | No |  |
| `currency` | `String` | No |  |
| `current_user_id` | `String` | No |  |
| `description` | `String` | No |  |
| `destination` | `Hash` | No |  |
| `document` | `Array` | No |  |
| `event` | `Array` | No |  |
| `formatted_amount` | `String` | No |  |
| `formatted_fx_destination_amount` | `String` | No |  |
| `formatted_schedule` | `String` | No |  |
| `formatted_schedule_day_only` | `String` | No |  |
| `formatted_source_account_available_balance` | `String` | No |  |
| `fx_destination_amount` | `Float` | No |  |
| `fx_destination_amount_minor_unit` | `Integer` | No |  |
| `fx_destination_currency` | `String` | No |  |
| `fx_quote_expires_at` | `String` | No |  |
| `fx_quote_id` | `String` | No |  |
| `fx_rate` | `Float` | No |  |
| `fx_use_destination_amount` | `Boolean` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `invoice_id` | `String` | No |  |
| `is_archived` | `Boolean` | No |  |
| `is_failed` | `Boolean` | No |  |
| `is_settled` | `Boolean` | No |  |
| `is_submitted` | `Boolean` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `merchant_token_description` | `String` | No |  |
| `nonce` | `String` | No |  |
| `payment_processor` | `String` | No |  |
| `payment_rail` | `String` | No |  |
| `payrun_id` | `String` | No |  |
| `payrun_name` | `String` | No |  |
| `rule` | `Hash` | No |  |
| `schedule_date` | `String` | No |  |
| `scheduled` | `Boolean` | No |  |
| `source_account_available_balance` | `Float` | No |  |
| `source_account_available_balance_minor_unit` | `Integer` | No |  |
| `source_account_bic` | `String` | No |  |
| `source_account_currency` | `String` | No |  |
| `source_account_iban` | `String` | No |  |
| `source_account_identifier` | `Hash` | Yes |  |
| `source_account_name` | `String` | No |  |
| `source_account_number` | `String` | No |  |
| `source_account_sortcode` | `String` | No |  |
| `status` | `String` | No |  |
| `tag` | `Array` | No |  |
| `their_reference` | `String` | No |  |
| `topup_payrun_id` | `String` | No |  |
| `transacted_amount` | `Float` | No |  |
| `transacted_fx_amount` | `Float` | No |  |
| `transacted_fx_rate` | `Float` | No |  |
| `type` | `String` | No |  |
| `user_id` | `String` | No |  |
| `your_reference` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionBizBizModelsPagingPayoutPage.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionBizBizModelsPagingPayoutPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionBizBizModelsPagingPayrunPageEntity

```ruby
no_frixion_biz_biz_models_paging_payrun_page = client.NoFrixionBizBizModelsPagingPayrunPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation` | `Array` | No |  |
| `authorisation_date` | `String` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `batch_payout_id` | `String` | No |  |
| `can_authorise` | `Boolean` | No |  |
| `can_delete` | `Boolean` | No |  |
| `can_edit` | `Boolean` | No |  |
| `event` | `Array` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `invoice` | `Array` | No |  |
| `invoices_minimal` | `Array` | No |  |
| `is_archived` | `Boolean` | No |  |
| `last_updated` | `String` | No |  |
| `last_updated_by` | `Hash` | Yes |  |
| `merchant_id` | `String` | No |  |
| `name` | `String` | No |  |
| `nonce` | `String` | No |  |
| `payment` | `Array` | No |  |
| `payout` | `Array` | No |  |
| `payouts_count` | `Integer` | No |  |
| `schedule_date` | `String` | No |  |
| `source_account` | `Array` | No |  |
| `status` | `String` | No |  |
| `total_eur` | `Float` | No |  |
| `total_gbp` | `Float` | No |  |
| `total_usd` | `Float` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionBizBizModelsPagingPayrunPage.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionBizBizModelsPagingPayrunPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionBizBizModelsPagingRuleEventsPageEntity

```ruby
no_frixion_biz_biz_models_paging_rule_events_page = client.NoFrixionBizBizModelsPagingRuleEventsPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_message` | `String` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `is_authorise_to_enable` | `Boolean` | No |  |
| `message` | `String` | No |  |
| `raw_response` | `String` | No |  |
| `rule_event_type` | `String` | No |  |
| `rule_id` | `String` | No |  |
| `user` | `Hash` | Yes |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionBizBizModelsPagingRuleEventsPage.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionBizBizModelsPagingRuleEventsPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionBizBizModelsPagingRulesPageEntity

```ruby
no_frixion_biz_biz_models_paging_rules_page = client.NoFrixionBizBizModelsPagingRulesPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account` | `Hash` | No |  |
| `account_id` | `String` | No |  |
| `approve_url` | `String` | No |  |
| `approver_id` | `String` | No |  |
| `authentication_method` | `Array` | No |  |
| `authorisation` | `Array` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `can_authorise` | `Boolean` | No |  |
| `created_by` | `Hash` | Yes |  |
| `description` | `String` | No |  |
| `end_at` | `String` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `is_disabled` | `Boolean` | No |  |
| `last_executed_at` | `String` | No |  |
| `last_run_at_transaction_date` | `String` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `name` | `String` | No |  |
| `nonce` | `String` | Yes |  |
| `on_approved_web_hook_url` | `String` | No |  |
| `on_execution_error_web_hook_url` | `String` | No |  |
| `on_execution_success_web_hook_url` | `String` | No |  |
| `start_at` | `String` | No |  |
| `status` | `String` | No |  |
| `sweep_action` | `Hash` | No |  |
| `time_zone_id` | `String` | No |  |
| `trigger_cron_expression` | `String` | No |  |
| `trigger_on_pay_in` | `Boolean` | No |  |
| `user_id` | `String` | No |  |
| `web_hook_secret` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionBizBizModelsPagingRulesPage.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionBizBizModelsPagingRulesPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionBizBizModelsPaymentsCardPaymentEntity

```ruby
no_frixion_biz_biz_models_payments_card_payment = client.NoFrixionBizBizModelsPaymentsCardPayment
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorized_amount` | `String` | No |  |
| `currency_code` | `String` | No |  |
| `is_payer_authentication_required` | `Boolean` | No |  |
| `is_soft_decline` | `Boolean` | No |  |
| `payer_authentication_access_token` | `String` | No |  |
| `payer_authentication_merchant_data` | `String` | No |  |
| `payer_authentication_url` | `String` | No |  |
| `payer_authentication_window_height` | `Integer` | No |  |
| `payer_authentication_window_width` | `Integer` | No |  |
| `payment_request_callback_url` | `String` | No |  |
| `payment_request_id` | `String` | No |  |
| `request_id` | `String` | No |  |
| `response_code` | `String` | No |  |
| `response_type` | `String` | No |  |
| `status` | `String` | No |  |
| `three_ds_redirect_url` | `String` | No |  |
| `transaction_id` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionBizBizModelsPaymentsCardPayment.create({
  "paymentrequest_id" => "example_paymentrequest_id", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionBizBizModelsPaymentsCardPaymentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionBizBizModelsPaymentsCardPublicKeyEntity

```ruby
no_frixion_biz_biz_models_payments_card_public_key = client.NoFrixionBizBizModelsPaymentsCardPublicKey
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jwt` | `String` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionBizBizModelsPaymentsCardPublicKey.load({ "paymentrequest_id" => "paymentrequest_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionBizBizModelsPaymentsCardPublicKeyEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity

```ruby
no_frixion_money_moov_api_features_beneficiaries_beneficiaries = client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `beneficiary` | `Array` | No |  |
| `failed_beneficiary` | `Hash` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries.create({
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity

```ruby
no_frixion_money_moov_api_features_payment_requests_payment = client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_payment_request` | `Hash` | No |  |
| `payment_request` | `Array` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment.create({
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity

```ruby
no_frixion_money_moov_api_features_permissions_roles_create = client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_role` | `Hash` | No |  |
| `role` | `Array` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate.create({
  "merchant_id" => "example_merchant_id", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity

```ruby
no_frixion_money_moov_api_features_user_invites_create = client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_user_invite` | `Hash` | No |  |
| `user_invite` | `Array` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate.create({
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity

```ruby
no_frixion_money_moov_models_authorisation_settings_merchant = client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount_lower` | `Float` | No |  |
| `amount_upper` | `Float` | No |  |
| `authorisation_type` | `String` | No |  |
| `beneficiaries_only` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `last_editor_cant_authorise` | `Boolean` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `number_of_authoriser` | `Integer` | No |  |
| `role_setting` | `Array` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsBatchPayoutEntity

```ruby
no_frixion_money_moov_models_batch_payout = client.NoFrixionMoneyMoovModelsBatchPayout
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approve_url` | `String` | No |  |
| `id` | `String` | No |  |
| `payout` | `Array` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsBatchPayout.create({
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsBatchPayout.load({ "id" => "no_frixion_money_moov_models_batch_payout_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsBatchPayoutEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity

```ruby
no_frixion_money_moov_models_beneficiary_group_page = client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `group_member` | `Array` | No |  |
| `group_name` | `String` | Yes |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | Yes |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsBeneficiaryPageEntity

```ruby
no_frixion_money_moov_models_beneficiary_page = client.NoFrixionMoneyMoovModelsBeneficiaryPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `String` | No |  |
| `authentication_method` | `Array` | No |  |
| `authorisation` | `Array` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `beneficiary_event` | `Array` | No |  |
| `can_authorise` | `Boolean` | No |  |
| `can_update` | `Boolean` | No |  |
| `created_by` | `Hash` | Yes |  |
| `created_by_email_address` | `String` | No |  |
| `currency` | `String` | Yes |  |
| `destination` | `Hash` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `is_enabled` | `Boolean` | No |  |
| `last_authorised` | `String` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `name` | `String` | Yes |  |
| `nonce` | `String` | No |  |
| `source_account` | `Array` | No |  |
| `their_reference` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsBeneficiaryPage.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsCardCustomerTokenEntity

```ruby
no_frixion_money_moov_models_card_customer_token = client.NoFrixionMoneyMoovModelsCardCustomerToken
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `card_type` | `String` | No |  |
| `customer_email_address` | `String` | No |  |
| `expiry_month` | `String` | No |  |
| `expiry_year` | `String` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `last_four_digit` | `String` | No |  |
| `last_updated` | `String` | No |  |
| `masked_card_number` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `payment_request_id` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsCardCustomerToken.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsCardCustomerToken.load({ "customer_email_address" => "customer_email_address" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsCardCustomerToken.remove()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsCardCustomerTokenEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity

```ruby
no_frixion_money_moov_models_currency_currency_info = client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `code` | `String` | No |  |
| `decimal` | `Integer` | No |  |
| `is_fiat` | `Boolean` | No |  |
| `iso4217_alpha_code` | `String` | No |  |
| `iso4217_numeric_code` | `String` | No |  |
| `symbol` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity

```ruby
no_frixion_money_moov_models_direct_debit_batch_submit = client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_submission` | `Hash` | No |  |
| `successful_submission` | `Array` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit.create({
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsFxRateEntity

```ruby
no_frixion_money_moov_models_fx_rate = client.NoFrixionMoneyMoovModelsFxRate
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destination_currency` | `String` | No |  |
| `exchange_rate` | `Float` | No |  |
| `expiry_time` | `String` | No |  |
| `quote_id` | `String` | No |  |
| `source_currency` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsFxRate.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsFxRate.load({ "destination" => "destination", "source" => "source", "valid_for_minute" => 1 })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsFxRateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsIPaymentEntity

```ruby
no_frixion_money_moov_models_i_payment = client.NoFrixionMoneyMoovModelsIPayment
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `payment_request_id` | `String` | No |  |
| `response_type` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsIPayment.create({
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsIPaymentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMandatesMandateEntity

```ruby
no_frixion_money_moov_models_mandates_mandate = client.NoFrixionMoneyMoovModelsMandatesMandate
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_number` | `String` | No |  |
| `address_line1` | `String` | Yes |  |
| `address_line2` | `String` | No |  |
| `approved_at` | `String` | No |  |
| `city` | `String` | Yes |  |
| `country_code` | `String` | Yes |  |
| `currency` | `String` | No |  |
| `customer_account_number` | `String` | No |  |
| `customer_city` | `String` | No |  |
| `customer_country_code` | `String` | No |  |
| `customer_country_name` | `String` | No |  |
| `customer_email_address` | `String` | No |  |
| `customer_first_name` | `String` | No |  |
| `customer_iban` | `String` | No |  |
| `customer_last_name` | `String` | No |  |
| `customer_sort_code` | `String` | No |  |
| `email_address` | `String` | Yes |  |
| `first_name` | `String` | Yes |  |
| `iban` | `String` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `is_recurring` | `Boolean` | No |  |
| `last_name` | `String` | Yes |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `postal_code` | `String` | Yes |  |
| `reference` | `String` | No |  |
| `sort_code` | `String` | No |  |
| `status` | `String` | No |  |
| `supplier_bank_account_id` | `String` | No |  |
| `supplier_customer_id` | `String` | No |  |
| `supplier_mandate_id` | `String` | No |  |
| `supplier_name` | `String` | No |  |
| `supplier_status` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsMandatesMandate.create({
  "address_line1" => "example_address_line1", # String
  "city" => "example_city", # String
  "country_code" => "example_country_code", # String
  "email_address" => "example_email_address", # String
  "first_name" => "example_first_name", # String
  "last_name" => "example_last_name", # String
  "postal_code" => "example_postal_code", # String
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsMandatesMandate.load({ "id" => "no_frixion_money_moov_models_mandates_mandate_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsMandatesMandateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantEntity

```ruby
no_frixion_money_moov_models_merchant = client.NoFrixionMoneyMoovModelsMerchant
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_currency` | `Array` | No |  |
| `can_have_trust_account` | `Boolean` | No |  |
| `card_payment_processor` | `String` | No |  |
| `company_id` | `String` | No |  |
| `display_qr_on_hosted_pay` | `Boolean` | No |  |
| `hosted_pay_version` | `Integer` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `is_blocked` | `Boolean` | No |  |
| `is_exited` | `Boolean` | No |  |
| `is_suspended` | `Boolean` | No |  |
| `jurisdiction` | `String` | No |  |
| `logo_url_png` | `String` | No |  |
| `logo_url_svg` | `String` | No |  |
| `merchant_category_code` | `String` | No |  |
| `name` | `String` | No |  |
| `note` | `String` | No |  |
| `parent_merchant` | `Hash` | No |  |
| `payment_account` | `Array` | No |  |
| `payment_account_limit` | `Integer` | No |  |
| `short_name` | `String` | No |  |
| `supported_payment_methods_list` | `Array` | No |  |
| `suspension_reason` | `String` | No |  |
| `tag` | `Array` | No |  |
| `time_zone_id` | `String` | No |  |
| `trading_name` | `String` | No |  |
| `web_hook_limit` | `Integer` | No |  |
| `your_role_name` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsMerchant.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsMerchant.load({ "id" => "no_frixion_money_moov_models_merchant_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsMerchant.update({
  "id" => "no_frixion_money_moov_models_merchant_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsMerchantEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantPageEntity

```ruby
no_frixion_money_moov_models_merchant_page = client.NoFrixionMoneyMoovModelsMerchantPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_currency` | `Array` | No |  |
| `can_have_trust_account` | `Boolean` | No |  |
| `card_payment_processor` | `String` | No |  |
| `company_id` | `String` | No |  |
| `display_qr_on_hosted_pay` | `Boolean` | No |  |
| `hosted_pay_version` | `Integer` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `is_blocked` | `Boolean` | No |  |
| `is_exited` | `Boolean` | No |  |
| `is_suspended` | `Boolean` | No |  |
| `jurisdiction` | `String` | No |  |
| `logo_url_png` | `String` | No |  |
| `logo_url_svg` | `String` | No |  |
| `merchant_category_code` | `String` | No |  |
| `name` | `String` | No |  |
| `note` | `String` | No |  |
| `parent_merchant` | `Hash` | No |  |
| `payment_account` | `Array` | No |  |
| `payment_account_limit` | `Integer` | No |  |
| `short_name` | `String` | No |  |
| `supported_payment_methods_list` | `Array` | No |  |
| `suspension_reason` | `String` | No |  |
| `tag` | `Array` | No |  |
| `time_zone_id` | `String` | No |  |
| `trading_name` | `String` | No |  |
| `web_hook_limit` | `Integer` | No |  |
| `your_role_name` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsMerchantPage.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsMerchantPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity

```ruby
no_frixion_money_moov_models_merchant_pay_by_bank_setting = client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bank_country_code` | `Array` | No |  |
| `bank_id` | `String` | No |  |
| `bank_name` | `String` | No |  |
| `business_institution_id` | `String` | No |  |
| `currency` | `String` | No |  |
| `logo` | `String` | No |  |
| `message` | `String` | No |  |
| `message_image_url` | `String` | No |  |
| `order` | `Integer` | No |  |
| `personal_institution_id` | `String` | No |  |
| `processor` | `String` | No |  |
| `warning_heading` | `String` | No |  |
| `warning_message` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantTokenEntity

```ruby
no_frixion_money_moov_models_merchant_token = client.NoFrixionMoneyMoovModelsMerchantToken
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authentication_method` | `Array` | No |  |
| `authorisation` | `Array` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `can_authorise` | `Boolean` | No |  |
| `description` | `String` | No |  |
| `expires_at` | `String` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `hmac_algorithm` | `String` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `ip_address_whitelist` | `String` | No |  |
| `is_archived` | `Boolean` | No |  |
| `is_enabled` | `Boolean` | No |  |
| `last_authorised` | `String` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `nonce` | `String` | Yes |  |
| `permission_type` | `Array` | No |  |
| `request_signature_version` | `Integer` | No |  |
| `shared_secret_algorithm` | `String` | No |  |
| `shared_secret_base64` | `String` | No |  |
| `token` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsMerchantToken.create({
  "nonce" => "example_nonce", # String
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsMerchantToken.load({ "id" => "no_frixion_money_moov_models_merchant_token_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsMerchantToken.update({
  "id" => "no_frixion_money_moov_models_merchant_token_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantTokenPageEntity

```ruby
no_frixion_money_moov_models_merchant_token_page = client.NoFrixionMoneyMoovModelsMerchantTokenPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authentication_method` | `Array` | No |  |
| `authorisation` | `Array` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `can_authorise` | `Boolean` | No |  |
| `description` | `String` | No |  |
| `expires_at` | `String` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `ip_address_whitelist` | `String` | No |  |
| `is_archived` | `Boolean` | No |  |
| `is_enabled` | `Boolean` | No |  |
| `last_authorised` | `String` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `nonce` | `String` | Yes |  |
| `permission_type` | `Array` | No |  |
| `request_signature_version` | `Integer` | No |  |
| `shared_secret_algorithm` | `String` | No |  |
| `shared_secret_base64` | `String` | No |  |
| `token` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsMerchantTokenPage.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsNoFrixionVersionEntity

```ruby
no_frixion_money_moov_models_no_frixion_version = client.NoFrixionMoneyMoovModelsNoFrixionVersion
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `build_version` | `Integer` | No |  |
| `major_version` | `Integer` | No |  |
| `minor_version` | `Integer` | No |  |
| `release_name` | `String` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsNoFrixionVersion.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsNoFrixionVersionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsOpenBankingAccountEntity

```ruby
no_frixion_money_moov_models_open_banking_account = client.NoFrixionMoneyMoovModelsOpenBankingAccount
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_balance` | `Array` | No |  |
| `account_identification` | `Array` | No |  |
| `account_name` | `Array` | No |  |
| `account_type` | `String` | No |  |
| `balance` | `Float` | No |  |
| `consolidated_account_information` | `Hash` | No |  |
| `currency` | `String` | No |  |
| `description` | `String` | No |  |
| `detail` | `String` | No |  |
| `id` | `String` | No |  |
| `nickname` | `String` | No |  |
| `type` | `String` | No |  |
| `usage_type` | `String` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsOpenBankingAccount.load({ "id" => "no_frixion_money_moov_models_open_banking_account_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsOpenBankingAccountEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsOpenBankingConsentEntity

```ruby
no_frixion_money_moov_models_open_banking_consent = client.NoFrixionMoneyMoovModelsOpenBankingConsent
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_url` | `String` | No |  |
| `callback_url` | `String` | No |  |
| `consent_id` | `String` | No |  |
| `email_address` | `String` | No |  |
| `expiry_date` | `String` | No |  |
| `failure_callback_url` | `String` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `institution_id` | `String` | No |  |
| `is_connected_account` | `Boolean` | No |  |
| `is_enabled` | `Boolean` | No |  |
| `merchant_id` | `String` | No |  |
| `provider` | `String` | No |  |
| `success_web_hook_url` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsOpenBankingConsent.create({
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsOpenBankingConsent.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsOpenBankingConsent.load({ "id" => "no_frixion_money_moov_models_open_banking_consent_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsOpenBankingConsent.remove({ "id" => "no_frixion_money_moov_models_open_banking_consent_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsOpenBankingConsent.update({
  "id" => "no_frixion_money_moov_models_open_banking_consent_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsOpenBankingConsentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsOpenBankingTransactionEntity

```ruby
no_frixion_money_moov_models_open_banking_transaction = client.NoFrixionMoneyMoovModelsOpenBankingTransaction
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address_detail` | `Hash` | No |  |
| `amount` | `Float` | No |  |
| `balance` | `Hash` | No |  |
| `booking_date_time` | `String` | No |  |
| `charge_detail` | `Hash` | No |  |
| `currency` | `String` | No |  |
| `currency_exchange` | `Hash` | No |  |
| `date` | `String` | No |  |
| `description` | `String` | No |  |
| `enrichment` | `Hash` | No |  |
| `gross_amount` | `Hash` | Yes |  |
| `id` | `String` | No |  |
| `iso_bank_transaction_code` | `Hash` | No |  |
| `merchant` | `Hash` | No |  |
| `payee_detail` | `Hash` | Yes |  |
| `payer_detail` | `Hash` | Yes |  |
| `proprietary_bank_transaction_code` | `Hash` | No |  |
| `reference` | `String` | No |  |
| `statement_reference` | `Array` | No |  |
| `status` | `String` | No |  |
| `supplementary_data` | `Object` | No |  |
| `transaction_amount` | `Hash` | Yes |  |
| `transaction_information` | `Array` | No |  |
| `transaction_mutability` | `String` | No |  |
| `value_date_time` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsOpenBankingTransaction.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsOpenBankingTransactionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentEntity

```ruby
no_frixion_money_moov_models_payment = client.NoFrixionMoneyMoovModelsPayment
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `Array` | No |  |
| `amount` | `Float` | No |  |
| `amount_pending` | `Float` | No |  |
| `amount_received` | `Float` | No |  |
| `amount_refunded` | `Float` | No |  |
| `auto_send_receipt` | `Boolean` | No |  |
| `base_origin_url` | `String` | No |  |
| `callback_url` | `String` | No |  |
| `card_authorize_only` | `Boolean` | No |  |
| `card_create_token` | `Boolean` | No |  |
| `card_create_token_mode` | `String` | No |  |
| `card_ignore_cvn` | `Boolean` | No |  |
| `card_no_payer_authentication` | `Boolean` | No |  |
| `card_processor_merchant_id` | `String` | No |  |
| `card_stripe_payment_intent_id` | `String` | No |  |
| `card_stripe_payment_intent_secret` | `String` | No |  |
| `card_transmit_raw_detail` | `Boolean` | No |  |
| `created_by_user` | `Hash` | Yes |  |
| `currency` | `String` | No |  |
| `custom_field` | `Array` | No |  |
| `customer_email_address` | `String` | No |  |
| `customer_id` | `String` | No |  |
| `customer_name` | `String` | No |  |
| `description` | `String` | No |  |
| `destination_account` | `Hash` | No |  |
| `direct_debit_payment` | `Hash` | No |  |
| `due_date` | `String` | No |  |
| `event` | `Array` | No |  |
| `failure_callback_url` | `String` | No |  |
| `field_display_setting` | `Array` | No |  |
| `formatted_amount` | `String` | No |  |
| `hosted_pay_checkout_url` | `String` | No |  |
| `id` | `String` | No |  |
| `ignore_address_verification` | `Boolean` | No |  |
| `inserted` | `String` | No |  |
| `inserted_sortable` | `String` | No |  |
| `is_archived` | `Boolean` | No |  |
| `jwk` | `String` | No |  |
| `last_updated` | `String` | No |  |
| `lightning_invoice` | `String` | No |  |
| `lightning_invoice_expires_at` | `String` | No |  |
| `merchant_direct_debit_mandate_id` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `merchant_token_description` | `String` | No |  |
| `notification_email_address` | `String` | No |  |
| `notification_role_i_d` | `Array` | No |  |
| `order_id` | `String` | No |  |
| `partial_payment_method` | `String` | No |  |
| `partial_payment_step` | `String` | No |  |
| `payment_attempt` | `Array` | No |  |
| `payment_method` | `Array` | No |  |
| `payment_processor` | `String` | No |  |
| `payrun_id` | `String` | No |  |
| `pisp_account_id` | `String` | No |  |
| `priority_bank_id` | `String` | No |  |
| `result` | `Hash` | No |  |
| `sandbox_settle_delay_in_second` | `Integer` | No |  |
| `shipping_address` | `Hash` | No |  |
| `shipping_address_city` | `String` | No |  |
| `shipping_address_country_code` | `String` | No |  |
| `shipping_address_county` | `String` | No |  |
| `shipping_address_line1` | `String` | No |  |
| `shipping_address_line2` | `String` | No |  |
| `shipping_address_post_code` | `String` | No |  |
| `shipping_email` | `String` | No |  |
| `shipping_first_name` | `String` | No |  |
| `shipping_last_name` | `String` | No |  |
| `shipping_phone` | `String` | No |  |
| `status` | `String` | No |  |
| `success_web_hook_url` | `String` | No |  |
| `tag` | `Array` | No |  |
| `tag_id` | `Array` | No |  |
| `title` | `String` | No |  |
| `tokenised_card` | `Array` | No |  |
| `transaction` | `Array` | No |  |
| `use_hosted_payment_page` | `Boolean` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsPayment.create({
  "created_by_user" => {}, # Hash
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsPayment.load({ "id" => "no_frixion_money_moov_models_payment_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsPayment.update({
  "id" => "no_frixion_money_moov_models_payment_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity

```ruby
no_frixion_money_moov_models_payment_account_minimal_page = client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_name` | `String` | No |  |
| `available_balance` | `Float` | No |  |
| `balance` | `Float` | No |  |
| `balance_minor_unit` | `Integer` | No |  |
| `currency` | `String` | No |  |
| `id` | `String` | No |  |
| `identifier` | `Hash` | Yes |  |
| `is_archived` | `Boolean` | No |  |
| `is_connected_account` | `Boolean` | No |  |
| `merchant_id` | `String` | No |  |
| `submitted_payouts_balance` | `Float` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentAccountPageEntity

```ruby
no_frixion_money_moov_models_payment_account_page = client.NoFrixionMoneyMoovModelsPaymentAccountPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_name` | `String` | No |  |
| `account_supplier_name` | `String` | No |  |
| `available_balance` | `Float` | No |  |
| `available_balance_minor_unit` | `Integer` | No |  |
| `balance` | `Float` | No |  |
| `balance_minor_unit` | `Integer` | No |  |
| `bank_name` | `String` | No |  |
| `consent_id` | `String` | No |  |
| `created_by` | `Hash` | Yes |  |
| `created_by_display_name` | `String` | No |  |
| `currency` | `String` | No |  |
| `default_payment_rail` | `String` | No |  |
| `display_name` | `String` | No |  |
| `expiry_date` | `String` | No |  |
| `external_account_icon` | `String` | No |  |
| `id` | `String` | No |  |
| `identifier` | `Hash` | Yes |  |
| `inserted` | `String` | No |  |
| `is_archived` | `Boolean` | No |  |
| `is_connected_account` | `Boolean` | No |  |
| `is_default` | `Boolean` | No |  |
| `is_trust_account` | `Boolean` | No |  |
| `is_virtual` | `Boolean` | No |  |
| `last_transaction` | `Hash` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `merchant_name` | `String` | No |  |
| `physical_account_id` | `String` | No |  |
| `rule` | `Array` | No |  |
| `submitted_payouts_balance` | `Float` | No |  |
| `submitted_payouts_balance_minor_unit` | `Integer` | No |  |
| `summary` | `String` | No |  |
| `supplier_sepa_instant_status` | `String` | No |  |
| `xero_bank_feed_connection_status` | `String` | No |  |
| `xero_bank_feed_last_synced_at` | `String` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `String` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `String` | No |  |
| `xero_bank_feed_sync_status` | `String` | No |  |
| `xero_unsynchronised_transactions_count` | `Integer` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsPaymentAccountPage.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentInitiationEntity

```ruby
no_frixion_money_moov_models_payment_initiation = client.NoFrixionMoneyMoovModelsPaymentInitiation
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `payment_initiation_id` | `String` | No |  |
| `payment_request_callback_url` | `String` | No |  |
| `payment_request_id` | `String` | No |  |
| `redirect_url` | `String` | No |  |
| `response_type` | `String` | No |  |
| `specific_error_message` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsPaymentInitiation.create({
  "paymentrequest_id" => "example_paymentrequest_id", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentInitiationEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestEventEntity

```ruby
no_frixion_money_moov_models_payment_request_event = client.NoFrixionMoneyMoovModelsPaymentRequestEvent
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `Float` | Yes |  |
| `apple_pay_transaction_id` | `String` | No |  |
| `card_authorization_response_id` | `String` | No |  |
| `card_expiry_month` | `Integer` | No |  |
| `card_expiry_year` | `Integer` | No |  |
| `card_issuer` | `String` | No |  |
| `card_issuer_country` | `String` | No |  |
| `card_last_four_digit` | `String` | No |  |
| `card_request_id` | `String` | No |  |
| `card_scheme` | `String` | No |  |
| `card_token_customer_id` | `String` | No |  |
| `card_transaction_id` | `String` | No |  |
| `currency` | `String` | No |  |
| `direct_debit_payment_id` | `String` | No |  |
| `direct_debit_payment_reference` | `String` | No |  |
| `drirect_debit_mandate_id` | `String` | No |  |
| `error_message` | `String` | No |  |
| `error_reason` | `String` | No |  |
| `event_type` | `String` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `lightning_invoice` | `String` | No |  |
| `lightning_r_hash` | `String` | No |  |
| `origin_url` | `String` | No |  |
| `payment_method_type` | `String` | No |  |
| `payment_processor_name` | `String` | No |  |
| `payment_request_id` | `String` | No |  |
| `pisp_bank_status` | `String` | No |  |
| `pisp_payment_initiation_id` | `String` | No |  |
| `pisp_payment_institution_name` | `String` | No |  |
| `pisp_payment_service_provider_id` | `String` | No |  |
| `pisp_redirect_url` | `String` | No |  |
| `reconciled_transaction_id` | `String` | No |  |
| `refund_payout_id` | `String` | No |  |
| `status` | `String` | No |  |
| `wallet_name` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsPaymentRequestEvent.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestEventEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestMetricEntity

```ruby
no_frixion_money_moov_models_payment_request_metric = client.NoFrixionMoneyMoovModelsPaymentRequestMetric
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `all` | `Integer` | No |  |
| `authorized` | `Integer` | No |  |
| `paid` | `Integer` | No |  |
| `partially_paid` | `Integer` | No |  |
| `total_amounts_by_currency` | `Hash` | No |  |
| `unpaid` | `Integer` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsPaymentRequestMetric.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMetricEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity

```ruby
no_frixion_money_moov_models_payment_request_minimal = client.NoFrixionMoneyMoovModelsPaymentRequestMinimal
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `Float` | No |  |
| `amount_pending` | `Float` | No |  |
| `amount_received` | `Float` | No |  |
| `amount_refunded` | `Float` | No |  |
| `callback_url` | `String` | No |  |
| `card_stripe_payment_intent_secret` | `String` | No |  |
| `country_code` | `String` | No |  |
| `currency` | `String` | No |  |
| `custom_fields_to_display` | `Array` | No |  |
| `description` | `String` | No |  |
| `due_date` | `String` | No |  |
| `field_display_setting` | `Array` | No |  |
| `google_pay_merchant_id` | `String` | No |  |
| `id` | `String` | No |  |
| `jwk` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `merchant_logo_url_png` | `String` | No |  |
| `merchant_logo_url_svg` | `String` | No |  |
| `merchant_name` | `String` | No |  |
| `merchant_short_name` | `String` | No |  |
| `partial_payment_method` | `String` | No |  |
| `payment_attempt` | `Array` | No |  |
| `payment_methods_list` | `Array` | No |  |
| `payment_processor` | `String` | No |  |
| `payment_processor_key` | `String` | No |  |
| `pisp_error` | `String` | No |  |
| `priority_bank_id` | `String` | No |  |
| `status` | `String` | No |  |
| `stripe_account_id` | `String` | No |  |
| `title` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsPaymentRequestMinimal.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestResultEntity

```ruby
no_frixion_money_moov_models_payment_request_result = client.NoFrixionMoneyMoovModelsPaymentRequestResult
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `Float` | No |  |
| `amount_pending` | `Float` | No |  |
| `amount_received` | `Float` | No |  |
| `amount_refunded` | `Float` | No |  |
| `currency` | `String` | No |  |
| `customer_id` | `String` | No |  |
| `payment` | `Array` | No |  |
| `payment_request_id` | `String` | No |  |
| `pisp_authorization` | `Array` | No |  |
| `requested_amount` | `Float` | No |  |
| `result` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsPaymentRequestResult.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestResultEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity

```ruby
no_frixion_money_moov_models_payment_requests_merchant_payment = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `description` | `String` | Yes |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `name` | `String` | Yes |  |
| `template` | `Hash` | Yes |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity

```ruby
no_frixion_money_moov_models_payment_requests_merchant_payment2 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `description` | `String` | Yes |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `name` | `String` | Yes |  |
| `template` | `Hash` | Yes |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2.load({ "paymentrequest_id" => "paymentrequest_id", "template_id" => "template_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity

```ruby
no_frixion_money_moov_models_payment_requests_merchant_payment3 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `description` | `String` | Yes |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `name` | `String` | Yes |  |
| `template` | `Hash` | Yes |  |

### Operations

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3.update({
  "paymentrequest_id" => "paymentrequest_id",
  "template_id" => "template_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity

```ruby
no_frixion_money_moov_models_payment_requests_merchant_payment4 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4
```

### Operations

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4.remove({ "paymentrequest_id" => "paymentrequest_id", "template_id" => "template_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayoutKeysetPageEntity

```ruby
no_frixion_money_moov_models_payout_keyset_page = client.NoFrixionMoneyMoovModelsPayoutKeysetPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `String` | No |  |
| `amount` | `Float` | No |  |
| `amount_minor_unit` | `Integer` | No |  |
| `approve_payout_url` | `String` | No |  |
| `approver_id` | `String` | No |  |
| `authentication_method` | `Array` | No |  |
| `authorisation` | `Array` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `batch_payout_id` | `String` | No |  |
| `beneficiary` | `Hash` | Yes |  |
| `can_authorise` | `Boolean` | No |  |
| `can_process` | `Boolean` | No |  |
| `can_update` | `Boolean` | No |  |
| `charge_bearer` | `String` | No |  |
| `created_by` | `String` | No |  |
| `created_by_email_address` | `String` | No |  |
| `currency` | `String` | No |  |
| `current_user_id` | `String` | No |  |
| `description` | `String` | No |  |
| `destination` | `Hash` | No |  |
| `document` | `Array` | No |  |
| `event` | `Array` | No |  |
| `formatted_amount` | `String` | No |  |
| `formatted_fx_destination_amount` | `String` | No |  |
| `formatted_schedule` | `String` | No |  |
| `formatted_schedule_day_only` | `String` | No |  |
| `formatted_source_account_available_balance` | `String` | No |  |
| `fx_destination_amount` | `Float` | No |  |
| `fx_destination_amount_minor_unit` | `Integer` | No |  |
| `fx_destination_currency` | `String` | No |  |
| `fx_quote_expires_at` | `String` | No |  |
| `fx_quote_id` | `String` | No |  |
| `fx_rate` | `Float` | No |  |
| `fx_use_destination_amount` | `Boolean` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `invoice_id` | `String` | No |  |
| `is_archived` | `Boolean` | No |  |
| `is_failed` | `Boolean` | No |  |
| `is_settled` | `Boolean` | No |  |
| `is_submitted` | `Boolean` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `merchant_token_description` | `String` | No |  |
| `nonce` | `String` | No |  |
| `payment_processor` | `String` | No |  |
| `payment_rail` | `String` | No |  |
| `payrun_id` | `String` | No |  |
| `payrun_name` | `String` | No |  |
| `rule` | `Hash` | No |  |
| `schedule_date` | `String` | No |  |
| `scheduled` | `Boolean` | No |  |
| `source_account_available_balance` | `Float` | No |  |
| `source_account_available_balance_minor_unit` | `Integer` | No |  |
| `source_account_bic` | `String` | No |  |
| `source_account_currency` | `String` | No |  |
| `source_account_iban` | `String` | No |  |
| `source_account_identifier` | `Hash` | Yes |  |
| `source_account_name` | `String` | No |  |
| `source_account_number` | `String` | No |  |
| `source_account_sortcode` | `String` | No |  |
| `status` | `String` | No |  |
| `tag` | `Array` | No |  |
| `their_reference` | `String` | No |  |
| `topup_payrun_id` | `String` | No |  |
| `transacted_amount` | `Float` | No |  |
| `transacted_fx_amount` | `Float` | No |  |
| `transacted_fx_rate` | `Float` | No |  |
| `type` | `String` | No |  |
| `user_id` | `String` | No |  |
| `your_reference` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsPayoutKeysetPage.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPayoutKeysetPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayoutMetricEntity

```ruby
no_frixion_money_moov_models_payout_metric = client.NoFrixionMoneyMoovModelsPayoutMetric
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `all` | `Float` | No |  |
| `failed` | `Float` | No |  |
| `in_progress` | `Float` | No |  |
| `paid` | `Float` | No |  |
| `pending_approval` | `Float` | No |  |
| `scheduled` | `Float` | No |  |
| `total_amounts_by_currency` | `Hash` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsPayoutMetric.load()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPayoutMetricEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity

```ruby
no_frixion_money_moov_models_payouts_payouts_create = client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_payout` | `Hash` | No |  |
| `payout` | `Array` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate.create({
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayrunEntity

```ruby
no_frixion_money_moov_models_payrun = client.NoFrixionMoneyMoovModelsPayrun
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation` | `Array` | No |  |
| `authorisation_date` | `String` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `batch_payout_id` | `String` | No |  |
| `can_authorise` | `Boolean` | No |  |
| `can_delete` | `Boolean` | No |  |
| `can_edit` | `Boolean` | No |  |
| `event` | `Array` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `invoice` | `Array` | No |  |
| `invoices_minimal` | `Array` | No |  |
| `is_archived` | `Boolean` | No |  |
| `last_updated` | `String` | No |  |
| `last_updated_by` | `Hash` | Yes |  |
| `merchant_id` | `String` | No |  |
| `name` | `String` | No |  |
| `nonce` | `String` | No |  |
| `payment` | `Array` | No |  |
| `payout` | `Array` | No |  |
| `payouts_count` | `Integer` | No |  |
| `reason` | `String` | No |  |
| `schedule_date` | `String` | No |  |
| `scheduled_date` | `String` | No |  |
| `source_account` | `Array` | No |  |
| `status` | `String` | No |  |
| `total_eur` | `Float` | No |  |
| `total_gbp` | `Float` | No |  |
| `total_usd` | `Float` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsPayrun.create({
  "id" => "example_id", # String
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsPayrun.load({ "id" => "no_frixion_money_moov_models_payrun_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsPayrun.update({
  "id" => "no_frixion_money_moov_models_payrun_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsPayrunEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsReportResultEntity

```ruby
no_frixion_money_moov_models_report_result = client.NoFrixionMoneyMoovModelsReportResult
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content` | `String` | No |  |
| `content_type` | `String` | No |  |
| `last_completed_at` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `report_name` | `String` | No |  |
| `report_type` | `String` | No |  |
| `statement_number` | `Integer` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsReportResult.load({ "id" => 1, "report_id" => "report_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsReportResultEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsRuleEntity

```ruby
no_frixion_money_moov_models_rule = client.NoFrixionMoneyMoovModelsRule
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account` | `Hash` | No |  |
| `account_id` | `String` | No |  |
| `approve_url` | `String` | No |  |
| `approver_id` | `String` | No |  |
| `authentication_method` | `Array` | No |  |
| `authorisation` | `Array` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `can_authorise` | `Boolean` | No |  |
| `created_by` | `Hash` | Yes |  |
| `description` | `String` | No |  |
| `end_at` | `String` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `is_disabled` | `Boolean` | No |  |
| `last_executed_at` | `String` | No |  |
| `last_run_at_transaction_date` | `String` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `name` | `String` | No |  |
| `nonce` | `String` | Yes |  |
| `on_approved_web_hook_url` | `String` | No |  |
| `on_execution_error_web_hook_url` | `String` | No |  |
| `on_execution_success_web_hook_url` | `String` | No |  |
| `start_at` | `String` | No |  |
| `status` | `String` | No |  |
| `sweep_action` | `Hash` | No |  |
| `time_zone_id` | `String` | No |  |
| `trigger_cron_expression` | `String` | No |  |
| `trigger_on_pay_in` | `Boolean` | No |  |
| `user_id` | `String` | No |  |
| `web_hook_secret` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsRule.create({
  "created_by" => {}, # Hash
  "nonce" => "example_nonce", # String
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsRule.load({ "id" => "no_frixion_money_moov_models_rule_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsRule.update({
  "id" => "no_frixion_money_moov_models_rule_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsRuleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsTransactionEntity

```ruby
no_frixion_money_moov_models_transaction = client.NoFrixionMoneyMoovModelsTransaction
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `String` | No |  |
| `account_name` | `String` | No |  |
| `account_sequence_number` | `Integer` | No |  |
| `amount` | `Float` | No |  |
| `amount_minor_unit` | `Integer` | No |  |
| `balance` | `Float` | No |  |
| `balance_minor_unit` | `Integer` | No |  |
| `counterparty` | `Hash` | No |  |
| `counterparty_summary` | `String` | No |  |
| `currency` | `String` | No |  |
| `description` | `String` | No |  |
| `fx_amount` | `Float` | No |  |
| `fx_currency` | `String` | No |  |
| `fx_rate` | `Float` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `payment_request_custom_field` | `Hash` | No |  |
| `payment_request_id` | `String` | No |  |
| `payout_id` | `String` | No |  |
| `raw_reference` | `String` | No |  |
| `rule_id` | `String` | No |  |
| `tag` | `Array` | No |  |
| `their_reference` | `String` | No |  |
| `transaction_date` | `String` | No |  |
| `type` | `String` | No |  |
| `virtual_iban` | `String` | No |  |
| `your_reference` | `String` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsTransaction.load({ "id" => "no_frixion_money_moov_models_transaction_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsTransactionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsTransactionPageEntity

```ruby
no_frixion_money_moov_models_transaction_page = client.NoFrixionMoneyMoovModelsTransactionPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `String` | No |  |
| `account_name` | `String` | No |  |
| `account_sequence_number` | `Integer` | No |  |
| `amount` | `Float` | No |  |
| `amount_minor_unit` | `Integer` | No |  |
| `balance` | `Float` | No |  |
| `balance_minor_unit` | `Integer` | No |  |
| `content` | `Array` | No |  |
| `counterparty` | `Hash` | No |  |
| `counterparty_summary` | `String` | No |  |
| `currency` | `String` | No |  |
| `description` | `String` | No |  |
| `fx_amount` | `Float` | No |  |
| `fx_currency` | `String` | No |  |
| `fx_rate` | `Float` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `page_number` | `Integer` | No |  |
| `page_size` | `Integer` | No |  |
| `payment_request_custom_field` | `Hash` | No |  |
| `payment_request_id` | `String` | No |  |
| `payout_id` | `String` | No |  |
| `raw_reference` | `String` | No |  |
| `rule_id` | `String` | No |  |
| `tag` | `Array` | No |  |
| `their_reference` | `String` | No |  |
| `total_page` | `Integer` | No |  |
| `total_size` | `Integer` | No |  |
| `transaction_date` | `String` | No |  |
| `type` | `String` | No |  |
| `virtual_iban` | `String` | No |  |
| `your_reference` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsTransactionPage.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsTransactionPage.load({ "account_id" => "account_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsTransactionPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsUserInviteEntity

```ruby
no_frixion_money_moov_models_user_invite = client.NoFrixionMoneyMoovModelsUserInvite
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_status` | `Hash` | No |  |
| `id` | `String` | No |  |
| `initial_role_id` | `String` | No |  |
| `invitee_email_address` | `String` | No |  |
| `invitee_first_name` | `String` | No |  |
| `invitee_last_name` | `String` | No |  |
| `inviter_email_address` | `String` | No |  |
| `inviter_first_name` | `String` | No |  |
| `inviter_last_name` | `String` | No |  |
| `is_authorised` | `Boolean` | No |  |
| `is_invitee_registered` | `Boolean` | No |  |
| `last_invited` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `merchant_name` | `String` | No |  |
| `message` | `String` | No |  |
| `registration_url` | `String` | No |  |
| `send_invite_email` | `Boolean` | No |  |
| `status` | `String` | No |  |
| `user` | `Hash` | Yes |  |
| `user_id` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsUserInvite.create({
  "user" => {}, # Hash
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsUserInvite.load({ "id" => "no_frixion_money_moov_models_user_invite_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsUserInviteEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsUserInvitePageEntity

```ruby
no_frixion_money_moov_models_user_invite_page = client.NoFrixionMoneyMoovModelsUserInvitePage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_status` | `Hash` | No |  |
| `id` | `String` | No |  |
| `initial_role_id` | `String` | No |  |
| `invitee_email_address` | `String` | No |  |
| `invitee_first_name` | `String` | No |  |
| `invitee_last_name` | `String` | No |  |
| `inviter_email_address` | `String` | No |  |
| `inviter_first_name` | `String` | No |  |
| `inviter_last_name` | `String` | No |  |
| `is_authorised` | `Boolean` | No |  |
| `is_invitee_registered` | `Boolean` | No |  |
| `last_invited` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `merchant_name` | `String` | No |  |
| `message` | `String` | No |  |
| `registration_url` | `String` | No |  |
| `status` | `String` | No |  |
| `user` | `Hash` | Yes |  |
| `user_id` | `String` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsUserInvitePage.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsUserInvitePageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsUserPageEntity

```ruby
no_frixion_money_moov_models_user_page = client.NoFrixionMoneyMoovModelsUserPage
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `Array` | No |  |
| `email_address` | `String` | Yes |  |
| `first_name` | `String` | Yes |  |
| `id` | `String` | No |  |
| `last_name` | `String` | Yes |  |
| `passkey_added` | `Boolean` | No |  |
| `permission` | `Hash` | No |  |
| `roles_with_scope` | `Array` | No |  |
| `two_factor_enabled` | `Boolean` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsUserPage.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsUserPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## NoFrixionMoneyMoovModelsWebhookEntity

```ruby
no_frixion_money_moov_models_webhook = client.NoFrixionMoneyMoovModelsWebhook
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destination_url` | `String` | No |  |
| `email_address` | `String` | No |  |
| `failed_notification_email_address` | `String` | No |  |
| `id` | `String` | No |  |
| `is_active` | `Boolean` | No |  |
| `merchant_id` | `String` | No |  |
| `notification_method` | `String` | No |  |
| `resource_type` | `Array` | No |  |
| `retry` | `Boolean` | No |  |
| `secret` | `String` | No |  |
| `version` | `Integer` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsWebhook.create({
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.NoFrixionMoneyMoovModelsWebhook.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsWebhook.load({ "id" => "no_frixion_money_moov_models_webhook_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.NoFrixionMoneyMoovModelsWebhook.update({
  "id" => "no_frixion_money_moov_models_webhook_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `NoFrixionMoneyMoovModelsWebhookEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## OpenBankingEntity

```ruby
open_banking = client.OpenBanking
```

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.OpenBanking.create({
  "account_id" => "example_account_id", # String
})
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.OpenBanking.remove()
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `OpenBankingEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PayeeverificationEntity

```ruby
payeeverification = client.Payeeverification
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_name` | `String` | Yes |  |
| `account_number` | `String` | No |  |
| `iban` | `String` | Yes |  |
| `payee_verified_account_name` | `String` | No |  |
| `result` | `String` | No |  |
| `secondary_identification` | `String` | No |  |
| `sort_code` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Payeeverification.create({
  "account_name" => "example_account_name", # String
  "iban" => "example_iban", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PayeeverificationEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PaymentRequestEntity

```ruby
payment_request = client.PaymentRequest
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `Float` | No |  |
| `do_simulate_settlement_failure` | `Boolean` | No |  |
| `error_description` | `String` | No |  |
| `institution` | `String` | No |  |
| `payment_initiation_id` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PaymentRequest.create({
  "paymentrequest_id" => "example_paymentrequest_id", # String
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.PaymentRequest.load()
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.PaymentRequest.remove({ "id" => "id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.PaymentRequest.update({
  "paymentrequest_id" => "paymentrequest_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PaymentRequestEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PayoutEntity

```ruby
payout = client.Payout
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `String` | No |  |
| `allow_incomplete` | `Boolean` | No |  |
| `amount` | `Float` | No |  |
| `amount_minor_unit` | `Integer` | No |  |
| `approve_payout_url` | `String` | No |  |
| `approver_id` | `String` | No |  |
| `authentication_method` | `Array` | No |  |
| `authorisation` | `Array` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `batch_payout_id` | `String` | No |  |
| `beneficiary` | `Hash` | Yes |  |
| `beneficiary_id` | `String` | No |  |
| `can_authorise` | `Boolean` | No |  |
| `can_process` | `Boolean` | No |  |
| `can_update` | `Boolean` | No |  |
| `charge_bearer` | `String` | No |  |
| `created_by` | `String` | No |  |
| `created_by_email_address` | `String` | No |  |
| `currency` | `String` | No |  |
| `current_user_id` | `String` | No |  |
| `description` | `String` | No |  |
| `destination` | `Hash` | No |  |
| `document` | `Array` | No |  |
| `event` | `Array` | No |  |
| `formatted_amount` | `String` | No |  |
| `formatted_fx_destination_amount` | `String` | No |  |
| `formatted_schedule` | `String` | No |  |
| `formatted_schedule_day_only` | `String` | No |  |
| `formatted_source_account_available_balance` | `String` | No |  |
| `fx_destination_amount` | `Float` | No |  |
| `fx_destination_amount_minor_unit` | `Integer` | No |  |
| `fx_destination_currency` | `String` | No |  |
| `fx_quote_expires_at` | `String` | No |  |
| `fx_quote_id` | `String` | No |  |
| `fx_rate` | `Float` | No |  |
| `fx_use_destination_amount` | `Boolean` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `invoice_id` | `String` | No |  |
| `is_archived` | `Boolean` | No |  |
| `is_failed` | `Boolean` | No |  |
| `is_settled` | `Boolean` | No |  |
| `is_submitted` | `Boolean` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `merchant_token_description` | `String` | No |  |
| `nonce` | `String` | No |  |
| `payment_processor` | `String` | No |  |
| `payment_rail` | `String` | No |  |
| `payrun_id` | `String` | No |  |
| `payrun_name` | `String` | No |  |
| `rule` | `Hash` | No |  |
| `schedule_date` | `String` | No |  |
| `scheduled` | `Boolean` | No |  |
| `source_account_available_balance` | `Float` | No |  |
| `source_account_available_balance_minor_unit` | `Integer` | No |  |
| `source_account_bic` | `String` | No |  |
| `source_account_currency` | `String` | No |  |
| `source_account_iban` | `String` | No |  |
| `source_account_identifier` | `Hash` | Yes |  |
| `source_account_name` | `String` | No |  |
| `source_account_number` | `String` | No |  |
| `source_account_sortcode` | `String` | No |  |
| `status` | `String` | No |  |
| `tag` | `Array` | No |  |
| `tag_id` | `Array` | No |  |
| `their_reference` | `String` | No |  |
| `topup_payrun_id` | `String` | No |  |
| `transacted_amount` | `Float` | No |  |
| `transacted_fx_amount` | `Float` | No |  |
| `transacted_fx_rate` | `Float` | No |  |
| `type` | `String` | No |  |
| `user_id` | `String` | No |  |
| `your_reference` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Payout.create({
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Payout.load({ "id" => "payout_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Payout.remove({ "id" => "payout_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Payout.update({
  "id" => "payout_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PayoutEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PayrunEntity

```ruby
payrun = client.Payrun
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `String` | No |  |
| `note` | `String` | No |  |
| `scheduled_date` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Payrun.create({
  "id" => "example_id", # String
})
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Payrun.remove({ "id" => "id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Payrun.update({
  "id" => "id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PayrunEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## RejectEntity

```ruby
reject = client.Reject
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `String` | No |  |
| `amount` | `Float` | No |  |
| `amount_minor_unit` | `Integer` | No |  |
| `approve_payout_url` | `String` | No |  |
| `approver_id` | `String` | No |  |
| `authentication_method` | `Array` | No |  |
| `authorisation` | `Array` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `batch_payout_id` | `String` | No |  |
| `beneficiary` | `Hash` | Yes |  |
| `can_authorise` | `Boolean` | No |  |
| `can_process` | `Boolean` | No |  |
| `can_update` | `Boolean` | No |  |
| `charge_bearer` | `String` | No |  |
| `created_by` | `String` | No |  |
| `created_by_email_address` | `String` | No |  |
| `currency` | `String` | No |  |
| `current_user_id` | `String` | No |  |
| `description` | `String` | No |  |
| `destination` | `Hash` | No |  |
| `document` | `Array` | No |  |
| `event` | `Array` | No |  |
| `formatted_amount` | `String` | No |  |
| `formatted_fx_destination_amount` | `String` | No |  |
| `formatted_schedule` | `String` | No |  |
| `formatted_schedule_day_only` | `String` | No |  |
| `formatted_source_account_available_balance` | `String` | No |  |
| `fx_destination_amount` | `Float` | No |  |
| `fx_destination_amount_minor_unit` | `Integer` | No |  |
| `fx_destination_currency` | `String` | No |  |
| `fx_quote_expires_at` | `String` | No |  |
| `fx_quote_id` | `String` | No |  |
| `fx_rate` | `Float` | No |  |
| `fx_use_destination_amount` | `Boolean` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `invoice_id` | `String` | No |  |
| `is_archived` | `Boolean` | No |  |
| `is_failed` | `Boolean` | No |  |
| `is_settled` | `Boolean` | No |  |
| `is_submitted` | `Boolean` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `merchant_token_description` | `String` | No |  |
| `nonce` | `String` | No |  |
| `payment_processor` | `String` | No |  |
| `payment_rail` | `String` | No |  |
| `payrun_id` | `String` | No |  |
| `payrun_name` | `String` | No |  |
| `reason` | `String` | No |  |
| `rule` | `Hash` | No |  |
| `schedule_date` | `String` | No |  |
| `scheduled` | `Boolean` | No |  |
| `source_account_available_balance` | `Float` | No |  |
| `source_account_available_balance_minor_unit` | `Integer` | No |  |
| `source_account_bic` | `String` | No |  |
| `source_account_currency` | `String` | No |  |
| `source_account_iban` | `String` | No |  |
| `source_account_identifier` | `Hash` | Yes |  |
| `source_account_name` | `String` | No |  |
| `source_account_number` | `String` | No |  |
| `source_account_sortcode` | `String` | No |  |
| `status` | `String` | No |  |
| `tag` | `Array` | No |  |
| `their_reference` | `String` | No |  |
| `topup_payrun_id` | `String` | No |  |
| `transacted_amount` | `Float` | No |  |
| `transacted_fx_amount` | `Float` | No |  |
| `transacted_fx_rate` | `Float` | No |  |
| `type` | `String` | No |  |
| `user_id` | `String` | No |  |
| `your_reference` | `String` | No |  |

### Operations

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Reject.update({
  "id" => "id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `RejectEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ReportEntity

```ruby
report = client.Report
```

### Operations

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Report.update({
  "id" => "id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ReportEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## RuleEntity

```ruby
rule = client.Rule
```

### Operations

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Rule.remove({ "id" => "id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Rule.update({
  "id" => "id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `RuleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SendEntity

```ruby
send = client.Send
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `String` | No |  |
| `allow_incomplete` | `Boolean` | No |  |
| `amount` | `Float` | No |  |
| `amount_minor_unit` | `Integer` | No |  |
| `approve_payout_url` | `String` | No |  |
| `approver_id` | `String` | No |  |
| `authentication_method` | `Array` | No |  |
| `authorisation` | `Array` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `batch_payout_id` | `String` | No |  |
| `beneficiary` | `Hash` | Yes |  |
| `beneficiary_id` | `String` | No |  |
| `can_authorise` | `Boolean` | No |  |
| `can_process` | `Boolean` | No |  |
| `can_update` | `Boolean` | No |  |
| `charge_bearer` | `String` | No |  |
| `created_by` | `String` | No |  |
| `created_by_email_address` | `String` | No |  |
| `currency` | `String` | No |  |
| `current_user_id` | `String` | No |  |
| `description` | `String` | No |  |
| `destination` | `Hash` | No |  |
| `document` | `Array` | No |  |
| `event` | `Array` | No |  |
| `formatted_amount` | `String` | No |  |
| `formatted_fx_destination_amount` | `String` | No |  |
| `formatted_schedule` | `String` | No |  |
| `formatted_schedule_day_only` | `String` | No |  |
| `formatted_source_account_available_balance` | `String` | No |  |
| `fx_destination_amount` | `Float` | No |  |
| `fx_destination_amount_minor_unit` | `Integer` | No |  |
| `fx_destination_currency` | `String` | No |  |
| `fx_quote_expires_at` | `String` | No |  |
| `fx_quote_id` | `String` | No |  |
| `fx_rate` | `Float` | No |  |
| `fx_use_destination_amount` | `Boolean` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `invoice_id` | `String` | No |  |
| `is_archived` | `Boolean` | No |  |
| `is_failed` | `Boolean` | No |  |
| `is_settled` | `Boolean` | No |  |
| `is_submitted` | `Boolean` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `merchant_token_description` | `String` | No |  |
| `nonce` | `String` | No |  |
| `payment_processor` | `String` | No |  |
| `payment_rail` | `String` | No |  |
| `payrun_id` | `String` | No |  |
| `payrun_name` | `String` | No |  |
| `rule` | `Hash` | No |  |
| `schedule_date` | `String` | No |  |
| `scheduled` | `Boolean` | No |  |
| `source_account_available_balance` | `Float` | No |  |
| `source_account_available_balance_minor_unit` | `Integer` | No |  |
| `source_account_bic` | `String` | No |  |
| `source_account_currency` | `String` | No |  |
| `source_account_iban` | `String` | No |  |
| `source_account_identifier` | `Hash` | Yes |  |
| `source_account_name` | `String` | No |  |
| `source_account_number` | `String` | No |  |
| `source_account_sortcode` | `String` | No |  |
| `status` | `String` | No |  |
| `tag` | `Array` | No |  |
| `tag_id` | `Array` | No |  |
| `their_reference` | `String` | No |  |
| `topup_payrun_id` | `String` | No |  |
| `transacted_amount` | `Float` | No |  |
| `transacted_fx_amount` | `Float` | No |  |
| `transacted_fx_rate` | `Float` | No |  |
| `type` | `String` | No |  |
| `user_id` | `String` | No |  |
| `your_reference` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Send.create({
  "beneficiary" => {}, # Hash
  "source_account_identifier" => {}, # Hash
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SendEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SendbeneficiaryEntity

```ruby
sendbeneficiary = client.Sendbeneficiary
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `String` | No |  |
| `allow_incomplete` | `Boolean` | No |  |
| `amount` | `Float` | No |  |
| `amount_minor_unit` | `Integer` | No |  |
| `approve_payout_url` | `String` | No |  |
| `approver_id` | `String` | No |  |
| `authentication_method` | `Array` | No |  |
| `authorisation` | `Array` | No |  |
| `authorisers_completed_count` | `Integer` | No |  |
| `authorisers_required_count` | `Integer` | No |  |
| `batch_payout_id` | `String` | No |  |
| `beneficiary` | `Hash` | Yes |  |
| `beneficiary_id` | `String` | No |  |
| `can_authorise` | `Boolean` | No |  |
| `can_process` | `Boolean` | No |  |
| `can_update` | `Boolean` | No |  |
| `charge_bearer` | `String` | No |  |
| `created_by` | `String` | No |  |
| `created_by_email_address` | `String` | No |  |
| `currency` | `String` | No |  |
| `current_user_id` | `String` | No |  |
| `description` | `String` | No |  |
| `destination` | `Hash` | No |  |
| `document` | `Array` | No |  |
| `event` | `Array` | No |  |
| `formatted_amount` | `String` | No |  |
| `formatted_fx_destination_amount` | `String` | No |  |
| `formatted_schedule` | `String` | No |  |
| `formatted_schedule_day_only` | `String` | No |  |
| `formatted_source_account_available_balance` | `String` | No |  |
| `fx_destination_amount` | `Float` | No |  |
| `fx_destination_amount_minor_unit` | `Integer` | No |  |
| `fx_destination_currency` | `String` | No |  |
| `fx_quote_expires_at` | `String` | No |  |
| `fx_quote_id` | `String` | No |  |
| `fx_rate` | `Float` | No |  |
| `fx_use_destination_amount` | `Boolean` | No |  |
| `has_current_user_authorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `invoice_id` | `String` | No |  |
| `is_archived` | `Boolean` | No |  |
| `is_failed` | `Boolean` | No |  |
| `is_settled` | `Boolean` | No |  |
| `is_submitted` | `Boolean` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `merchant_token_description` | `String` | No |  |
| `nonce` | `String` | No |  |
| `payment_processor` | `String` | No |  |
| `payment_rail` | `String` | No |  |
| `payrun_id` | `String` | No |  |
| `payrun_name` | `String` | No |  |
| `rule` | `Hash` | No |  |
| `schedule_date` | `String` | No |  |
| `scheduled` | `Boolean` | No |  |
| `source_account_available_balance` | `Float` | No |  |
| `source_account_available_balance_minor_unit` | `Integer` | No |  |
| `source_account_bic` | `String` | No |  |
| `source_account_currency` | `String` | No |  |
| `source_account_iban` | `String` | No |  |
| `source_account_identifier` | `Hash` | Yes |  |
| `source_account_name` | `String` | No |  |
| `source_account_number` | `String` | No |  |
| `source_account_sortcode` | `String` | No |  |
| `status` | `String` | No |  |
| `tag` | `Array` | No |  |
| `tag_id` | `Array` | No |  |
| `their_reference` | `String` | No |  |
| `topup_payrun_id` | `String` | No |  |
| `transacted_amount` | `Float` | No |  |
| `transacted_fx_amount` | `Float` | No |  |
| `transacted_fx_rate` | `Float` | No |  |
| `type` | `String` | No |  |
| `user_id` | `String` | No |  |
| `your_reference` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Sendbeneficiary.create({
  "beneficiary" => {}, # Hash
  "source_account_identifier" => {}, # Hash
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SendbeneficiaryEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## TagEntity

```ruby
tag = client.Tag
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `colour_hex` | `String` | No |  |
| `description` | `String` | No |  |
| `id` | `String` | No |  |
| `merchant_id` | `String` | Yes |  |
| `name` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Tag.create({
  "merchant_id" => "example_merchant_id", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Tag.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `TagEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## TokenEntity

```ruby
token = client.Token
```

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Token.create({
  "id" => "example_id", # String
})
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Token.remove({ "id" => "id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `TokenEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## TransactionEntity

```ruby
transaction = client.Transaction
```

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Transaction.create({
  "id" => "example_id", # String
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Transaction.load({ "sequence_number" => 1, "transaction_id" => "transaction_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Transaction.remove({ "id" => "id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `TransactionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## UserEntity

```ruby
user = client.User
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `Array` | No |  |
| `email_address` | `String` | Yes |  |
| `first_name` | `String` | Yes |  |
| `id` | `String` | No |  |
| `last_name` | `String` | Yes |  |
| `passkey_added` | `Boolean` | No |  |
| `permission` | `Hash` | No |  |
| `profile` | `String` | No |  |
| `roles_with_scope` | `Array` | No |  |
| `two_factor_enabled` | `Boolean` | No |  |
| `user_invite_id` | `String` | No |  |

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

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.User.list
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.User.update({
  "id" => "id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `UserEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## UserInviteEntity

```ruby
user_invite = client.UserInvite
```

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.UserInvite.create({
  "id" => "example_id", # String
})
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.UserInvite.remove({ "id" => "id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.UserInvite.update({
  "id" => "id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `UserInviteEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## VirtualEntity

```ruby
virtual = client.Virtual
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_name` | `String` | No |  |
| `account_supplier_name` | `String` | No |  |
| `available_balance` | `Float` | No |  |
| `available_balance_minor_unit` | `Integer` | No |  |
| `balance` | `Float` | No |  |
| `balance_minor_unit` | `Integer` | No |  |
| `bank_name` | `String` | No |  |
| `consent_id` | `String` | No |  |
| `created_by` | `Hash` | Yes |  |
| `created_by_display_name` | `String` | No |  |
| `currency` | `String` | No |  |
| `default_payment_rail` | `String` | No |  |
| `display_name` | `String` | No |  |
| `expiry_date` | `String` | No |  |
| `external_account_icon` | `String` | No |  |
| `id` | `String` | No |  |
| `identifier` | `Hash` | Yes |  |
| `inserted` | `String` | No |  |
| `is_archived` | `Boolean` | No |  |
| `is_connected_account` | `Boolean` | No |  |
| `is_default` | `Boolean` | No |  |
| `is_trust_account` | `Boolean` | No |  |
| `is_virtual` | `Boolean` | No |  |
| `last_transaction` | `Hash` | No |  |
| `last_updated` | `String` | No |  |
| `merchant_id` | `String` | No |  |
| `merchant_name` | `String` | No |  |
| `name` | `String` | Yes |  |
| `physical_account_id` | `String` | No |  |
| `rule` | `Array` | No |  |
| `submitted_payouts_balance` | `Float` | No |  |
| `submitted_payouts_balance_minor_unit` | `Integer` | No |  |
| `summary` | `String` | No |  |
| `supplier_sepa_instant_status` | `String` | No |  |
| `xero_bank_feed_connection_status` | `String` | No |  |
| `xero_bank_feed_last_synced_at` | `String` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `String` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `String` | No |  |
| `xero_bank_feed_sync_status` | `String` | No |  |
| `xero_unsynchronised_transactions_count` | `Integer` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Virtual.create({
  "account_id" => "example_account_id", # String
})
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Virtual.update({
  "account_id" => "account_id",
  "id" => "id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `VirtualEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## WebhookEntity

```ruby
webhook = client.Webhook
```

### Operations

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Webhook.remove({ "id" => "id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `WebhookEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## WhoamiEntity

```ruby
whoami = client.Whoami
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `Array` | No |  |
| `email_address` | `String` | Yes |  |
| `first_name` | `String` | Yes |  |
| `id` | `String` | No |  |
| `last_name` | `String` | Yes |  |
| `passkey_added` | `Boolean` | No |  |
| `permission` | `Hash` | No |  |
| `roles_with_scope` | `Array` | No |  |
| `two_factor_enabled` | `Boolean` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Whoami.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `WhoamiEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## WhoamitrustedappEntity

```ruby
whoamitrustedapp = client.Whoamitrustedapp
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `Array` | No |  |
| `email_address` | `String` | Yes |  |
| `first_name` | `String` | Yes |  |
| `id` | `String` | No |  |
| `last_name` | `String` | Yes |  |
| `passkey_added` | `Boolean` | No |  |
| `permission` | `Hash` | No |  |
| `roles_with_scope` | `Array` | No |  |
| `two_factor_enabled` | `Boolean` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Whoamitrustedapp.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `WhoamitrustedappEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = NofrixionSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```

