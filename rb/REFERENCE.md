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

#### `Batch(data = nil)`

Create a new `Batch` entity instance. Pass `nil` for no initial data.

#### `Beneficiary(data = nil)`

Create a new `Beneficiary` entity instance. Pass `nil` for no initial data.

#### `BeneficiaryGroup(data = nil)`

Create a new `BeneficiaryGroup` entity instance. Pass `nil` for no initial data.

#### `Card(data = nil)`

Create a new `Card` entity instance. Pass `nil` for no initial data.

#### `CardCustomerToken(data = nil)`

Create a new `CardCustomerToken` entity instance. Pass `nil` for no initial data.

#### `CardPayment(data = nil)`

Create a new `CardPayment` entity instance. Pass `nil` for no initial data.

#### `CardPublicKey(data = nil)`

Create a new `CardPublicKey` entity instance. Pass `nil` for no initial data.

#### `Consent(data = nil)`

Create a new `Consent` entity instance. Pass `nil` for no initial data.

#### `Currency(data = nil)`

Create a new `Currency` entity instance. Pass `nil` for no initial data.

#### `DirectDebitBatchSubmit(data = nil)`

Create a new `DirectDebitBatchSubmit` entity instance. Pass `nil` for no initial data.

#### `FxRate(data = nil)`

Create a new `FxRate` entity instance. Pass `nil` for no initial data.

#### `IPayment(data = nil)`

Create a new `IPayment` entity instance. Pass `nil` for no initial data.

#### `Mandate(data = nil)`

Create a new `Mandate` entity instance. Pass `nil` for no initial data.

#### `Merchant(data = nil)`

Create a new `Merchant` entity instance. Pass `nil` for no initial data.

#### `MerchantAuthorisationSetting(data = nil)`

Create a new `MerchantAuthorisationSetting` entity instance. Pass `nil` for no initial data.

#### `MerchantDirectDebitMandatePage(data = nil)`

Create a new `MerchantDirectDebitMandatePage` entity instance. Pass `nil` for no initial data.

#### `MerchantPayByBankSetting(data = nil)`

Create a new `MerchantPayByBankSetting` entity instance. Pass `nil` for no initial data.

#### `MerchantPaymentRequestTemplate(data = nil)`

Create a new `MerchantPaymentRequestTemplate` entity instance. Pass `nil` for no initial data.

#### `MerchantToken(data = nil)`

Create a new `MerchantToken` entity instance. Pass `nil` for no initial data.

#### `Metadata(data = nil)`

Create a new `Metadata` entity instance. Pass `nil` for no initial data.

#### `NoFrixionVersion(data = nil)`

Create a new `NoFrixionVersion` entity instance. Pass `nil` for no initial data.

#### `OpenBanking(data = nil)`

Create a new `OpenBanking` entity instance. Pass `nil` for no initial data.

#### `Payeeverification(data = nil)`

Create a new `Payeeverification` entity instance. Pass `nil` for no initial data.

#### `Payment(data = nil)`

Create a new `Payment` entity instance. Pass `nil` for no initial data.

#### `PaymentAccount(data = nil)`

Create a new `PaymentAccount` entity instance. Pass `nil` for no initial data.

#### `PaymentAccountMinimal(data = nil)`

Create a new `PaymentAccountMinimal` entity instance. Pass `nil` for no initial data.

#### `PaymentInitiation(data = nil)`

Create a new `PaymentInitiation` entity instance. Pass `nil` for no initial data.

#### `PaymentRequest(data = nil)`

Create a new `PaymentRequest` entity instance. Pass `nil` for no initial data.

#### `PaymentRequestEvent(data = nil)`

Create a new `PaymentRequestEvent` entity instance. Pass `nil` for no initial data.

#### `PaymentRequestMetric(data = nil)`

Create a new `PaymentRequestMetric` entity instance. Pass `nil` for no initial data.

#### `PaymentRequestMinimal(data = nil)`

Create a new `PaymentRequestMinimal` entity instance. Pass `nil` for no initial data.

#### `PaymentRequestResult(data = nil)`

Create a new `PaymentRequestResult` entity instance. Pass `nil` for no initial data.

#### `Payout(data = nil)`

Create a new `Payout` entity instance. Pass `nil` for no initial data.

#### `PayoutKeysetPage(data = nil)`

Create a new `PayoutKeysetPage` entity instance. Pass `nil` for no initial data.

#### `PayoutMetric(data = nil)`

Create a new `PayoutMetric` entity instance. Pass `nil` for no initial data.

#### `Payrun(data = nil)`

Create a new `Payrun` entity instance. Pass `nil` for no initial data.

#### `Report(data = nil)`

Create a new `Report` entity instance. Pass `nil` for no initial data.

#### `ReportResult(data = nil)`

Create a new `ReportResult` entity instance. Pass `nil` for no initial data.

#### `Role(data = nil)`

Create a new `Role` entity instance. Pass `nil` for no initial data.

#### `Rule(data = nil)`

Create a new `Rule` entity instance. Pass `nil` for no initial data.

#### `RuleEvent(data = nil)`

Create a new `RuleEvent` entity instance. Pass `nil` for no initial data.

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
| `account_balance` | `Array` | No |  |
| `account_id` | `String` | No |  |
| `account_identification` | `Array` | No |  |
| `account_name` | `String` | No |  |
| `account_supplier_name` | `String` | No |  |
| `account_type` | `String` | No |  |
| `available_balance` | `Float` | No |  |
| `available_balance_minor_unit` | `Integer` | No |  |
| `balance` | `Float` | No |  |
| `balance_minor_unit` | `Integer` | No |  |
| `bank_name` | `String` | No |  |
| `consent_id` | `String` | No |  |
| `consolidated_account_information` | `Hash` | No |  |
| `created_by` | `Hash` | Yes |  |
| `created_by_display_name` | `String` | No |  |
| `currency` | `String` | No |  |
| `default_payment_rail` | `String` | No |  |
| `description` | `String` | No |  |
| `detail` | `String` | No |  |
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
| `nickname` | `String` | No |  |
| `physical_account_id` | `String` | No |  |
| `role_i_d` | `Array` | No |  |
| `rule` | `Array` | No |  |
| `submitted_payouts_balance` | `Float` | No |  |
| `submitted_payouts_balance_minor_unit` | `Integer` | No |  |
| `summary` | `String` | No |  |
| `supplier_physical_account_id` | `String` | No |  |
| `supplier_sepa_instant_status` | `String` | No |  |
| `to_date` | `String` | No |  |
| `type` | `String` | No |  |
| `usage_type` | `String` | No |  |
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

## BatchEntity

```ruby
batch = client.Batch
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
result = client.Batch.create({
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Batch.load({ "id" => "batch_id" })
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

Create a new `BatchEntity` instance with the same client and
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
| `beneficiary` | `Array` | No |  |
| `beneficiary_event` | `Array` | No |  |
| `can_authorise` | `Boolean` | No |  |
| `can_update` | `Boolean` | No |  |
| `created_by` | `Hash` | Yes |  |
| `created_by_email_address` | `String` | No |  |
| `currency` | `String` | Yes |  |
| `destination` | `Hash` | No |  |
| `failed_beneficiary` | `Hash` | No |  |
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

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `approval_callback_url` | - | - | - | - | - |
| `authentication_method` | - | - | - | - | - |
| `authorisation` | - | - | - | - | - |
| `authorisers_completed_count` | - | - | - | - | - |
| `authorisers_required_count` | - | - | - | - | - |
| `beneficiary` | - | - | - | - | - |
| `beneficiary_event` | - | - | - | - | - |
| `can_authorise` | - | - | - | - | - |
| `can_update` | - | - | - | - | - |
| `created_by` | - | - | - | - | - |
| `created_by_email_address` | - | - | - | - | - |
| `currency` | - | - | - | Yes | - |
| `destination` | - | - | Yes | - | - |
| `failed_beneficiary` | - | - | - | - | - |
| `has_current_user_authorised` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `inserted` | - | - | - | - | - |
| `is_enabled` | - | - | - | - | - |
| `last_authorised` | - | - | - | - | - |
| `last_updated` | - | - | - | - | - |
| `merchant_id` | - | - | Yes | - | - |
| `name` | - | - | - | Yes | - |
| `nonce` | - | - | - | - | - |
| `source_account` | - | - | - | - | - |
| `source_account_i_d` | - | - | - | - | - |
| `their_reference` | - | - | - | - | - |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Beneficiary.create({
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Beneficiary.list
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

## BeneficiaryGroupEntity

```ruby
beneficiary_group = client.BeneficiaryGroup
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
results = client.BeneficiaryGroup.list
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

Create a new `BeneficiaryGroupEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CardEntity

```ruby
card = client.Card
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
result = client.Card.create({
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

Create a new `CardEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CardCustomerTokenEntity

```ruby
card_customer_token = client.CardCustomerToken
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
results = client.CardCustomerToken.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.CardCustomerToken.load({ "customer_email_address" => "customer_email_address" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.CardCustomerToken.remove()
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

Create a new `CardCustomerTokenEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CardPaymentEntity

```ruby
card_payment = client.CardPayment
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
result = client.CardPayment.create({
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

Create a new `CardPaymentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CardPublicKeyEntity

```ruby
card_public_key = client.CardPublicKey
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jwt` | `String` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.CardPublicKey.load({ "paymentrequest_id" => "paymentrequest_id" })
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

Create a new `CardPublicKeyEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ConsentEntity

```ruby
consent = client.Consent
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
result = client.Consent.create({
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Consent.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Consent.load({ "id" => "consent_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Consent.remove({ "id" => "consent_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Consent.update({
  "id" => "consent_id",
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

Create a new `ConsentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## CurrencyEntity

```ruby
currency = client.Currency
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
results = client.Currency.list
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

Create a new `CurrencyEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## DirectDebitBatchSubmitEntity

```ruby
direct_debit_batch_submit = client.DirectDebitBatchSubmit
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
result = client.DirectDebitBatchSubmit.create({
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

Create a new `DirectDebitBatchSubmitEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## FxRateEntity

```ruby
fx_rate = client.FxRate
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
results = client.FxRate.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.FxRate.load({ "destination" => "destination", "source" => "source", "valid_for_minute" => 1 })
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

Create a new `FxRateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## IPaymentEntity

```ruby
i_payment = client.IPayment
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
result = client.IPayment.create({
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

Create a new `IPaymentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## MandateEntity

```ruby
mandate = client.Mandate
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
result = client.Mandate.create({
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
result = client.Mandate.load({ "id" => "mandate_id" })
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

Create a new `MandateEntity` instance with the same client and
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
| `reason` | `String` | No |  |
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
results = client.Merchant.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Merchant.load({ "id" => "merchant_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Merchant.remove({ "id" => "merchant_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Merchant.update({
  "id" => "merchant_id",
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

## MerchantAuthorisationSettingEntity

```ruby
merchant_authorisation_setting = client.MerchantAuthorisationSetting
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
results = client.MerchantAuthorisationSetting.list
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

Create a new `MerchantAuthorisationSettingEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## MerchantDirectDebitMandatePageEntity

```ruby
merchant_direct_debit_mandate_page = client.MerchantDirectDebitMandatePage
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
results = client.MerchantDirectDebitMandatePage.list
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

Create a new `MerchantDirectDebitMandatePageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## MerchantPayByBankSettingEntity

```ruby
merchant_pay_by_bank_setting = client.MerchantPayByBankSetting
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
results = client.MerchantPayByBankSetting.list
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

Create a new `MerchantPayByBankSettingEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## MerchantPaymentRequestTemplateEntity

```ruby
merchant_payment_request_template = client.MerchantPaymentRequestTemplate
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
results = client.MerchantPaymentRequestTemplate.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.MerchantPaymentRequestTemplate.load({ "id" => "merchant_payment_request_template_id", "paymentrequest_id" => "paymentrequest_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.MerchantPaymentRequestTemplate.remove({ "id" => "merchant_payment_request_template_id", "paymentrequest_id" => "paymentrequest_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.MerchantPaymentRequestTemplate.update({
  "id" => "merchant_payment_request_template_id",
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

Create a new `MerchantPaymentRequestTemplateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## MerchantTokenEntity

```ruby
merchant_token = client.MerchantToken
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

| Field | load | list | create | update |
| --- | --- | --- | --- | --- |
| `authentication_method` | - | - | - | - |
| `authorisation` | - | - | - | - |
| `authorisers_completed_count` | - | - | - | - |
| `authorisers_required_count` | - | - | - | - |
| `can_authorise` | - | - | - | - |
| `description` | - | - | Yes | - |
| `expires_at` | - | - | - | - |
| `has_current_user_authorised` | - | - | - | - |
| `hmac_algorithm` | - | - | - | - |
| `id` | - | - | - | - |
| `inserted` | - | - | - | - |
| `ip_address_whitelist` | - | - | - | - |
| `is_archived` | - | - | - | - |
| `is_enabled` | - | - | - | - |
| `last_authorised` | - | - | - | - |
| `last_updated` | - | - | - | - |
| `merchant_id` | - | - | Yes | - |
| `nonce` | - | - | - | - |
| `permission_type` | - | - | - | - |
| `request_signature_version` | - | - | - | - |
| `shared_secret_algorithm` | - | - | - | - |
| `shared_secret_base64` | - | - | - | - |
| `token` | - | - | - | - |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.MerchantToken.create({
  "nonce" => "example_nonce", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.MerchantToken.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.MerchantToken.load({ "id" => "merchant_token_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.MerchantToken.update({
  "id" => "merchant_token_id",
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

Create a new `MerchantTokenEntity` instance with the same client and
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

## NoFrixionVersionEntity

```ruby
no_frixion_version = client.NoFrixionVersion
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
result = client.NoFrixionVersion.load()
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

Create a new `NoFrixionVersionEntity` instance with the same client and
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

## PaymentEntity

```ruby
payment = client.Payment
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
result = client.Payment.create({
  "created_by_user" => {}, # Hash
})
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Payment.load({ "id" => "payment_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Payment.update({
  "id" => "payment_id",
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

Create a new `PaymentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PaymentAccountEntity

```ruby
payment_account = client.PaymentAccount
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
results = client.PaymentAccount.list
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

Create a new `PaymentAccountEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PaymentAccountMinimalEntity

```ruby
payment_account_minimal = client.PaymentAccountMinimal
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
results = client.PaymentAccountMinimal.list
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

Create a new `PaymentAccountMinimalEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PaymentInitiationEntity

```ruby
payment_initiation = client.PaymentInitiation
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
result = client.PaymentInitiation.create({
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

Create a new `PaymentInitiationEntity` instance with the same client and
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
| `do_simulate_settlement_failure` | `Boolean` | No |  |
| `due_date` | `String` | No |  |
| `error_description` | `String` | No |  |
| `event` | `Array` | No |  |
| `failed_payment_request` | `Hash` | No |  |
| `failure_callback_url` | `String` | No |  |
| `field_display_setting` | `Array` | No |  |
| `formatted_amount` | `String` | No |  |
| `hosted_pay_checkout_url` | `String` | No |  |
| `id` | `String` | No |  |
| `ignore_address_verification` | `Boolean` | No |  |
| `inserted` | `String` | No |  |
| `inserted_sortable` | `String` | No |  |
| `institution` | `String` | No |  |
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
| `payment_initiation_id` | `String` | No |  |
| `payment_method` | `Array` | No |  |
| `payment_processor` | `String` | No |  |
| `payment_request` | `Array` | No |  |
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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PaymentRequest.create({
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.PaymentRequest.list
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

## PaymentRequestEventEntity

```ruby
payment_request_event = client.PaymentRequestEvent
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
results = client.PaymentRequestEvent.list
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

Create a new `PaymentRequestEventEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PaymentRequestMetricEntity

```ruby
payment_request_metric = client.PaymentRequestMetric
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
result = client.PaymentRequestMetric.load()
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

Create a new `PaymentRequestMetricEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PaymentRequestMinimalEntity

```ruby
payment_request_minimal = client.PaymentRequestMinimal
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
results = client.PaymentRequestMinimal.list
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

Create a new `PaymentRequestMinimalEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PaymentRequestResultEntity

```ruby
payment_request_result = client.PaymentRequestResult
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
results = client.PaymentRequestResult.list
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

Create a new `PaymentRequestResultEntity` instance with the same client and
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
| `failed_payout` | `Hash` | No |  |
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
| `payout` | `Array` | No |  |
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

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `account_id` | - | - | Yes | - | - |
| `allow_incomplete` | - | - | - | - | - |
| `amount` | - | - | - | - | - |
| `amount_minor_unit` | - | - | - | - | - |
| `approve_payout_url` | - | - | - | - | - |
| `approver_id` | - | - | - | - | - |
| `authentication_method` | - | - | - | - | - |
| `authorisation` | - | - | - | - | - |
| `authorisers_completed_count` | - | - | - | - | - |
| `authorisers_required_count` | - | - | - | - | - |
| `batch_payout_id` | - | - | - | - | - |
| `beneficiary` | - | - | - | - | - |
| `beneficiary_id` | - | - | - | - | - |
| `can_authorise` | - | - | - | - | - |
| `can_process` | - | - | - | - | - |
| `can_update` | - | - | - | - | - |
| `charge_bearer` | - | - | - | - | - |
| `created_by` | - | - | - | - | - |
| `created_by_email_address` | - | - | - | - | - |
| `currency` | - | - | Yes | - | - |
| `current_user_id` | - | - | - | - | - |
| `description` | - | - | - | - | - |
| `destination` | - | - | - | - | - |
| `document` | - | - | - | - | - |
| `event` | - | - | - | - | - |
| `failed_payout` | - | - | - | - | - |
| `formatted_amount` | - | - | - | - | - |
| `formatted_fx_destination_amount` | - | - | - | - | - |
| `formatted_schedule` | - | - | - | - | - |
| `formatted_schedule_day_only` | - | - | - | - | - |
| `formatted_source_account_available_balance` | - | - | - | - | - |
| `fx_destination_amount` | - | - | - | - | - |
| `fx_destination_amount_minor_unit` | - | - | - | - | - |
| `fx_destination_currency` | - | - | - | - | - |
| `fx_quote_expires_at` | - | - | - | - | - |
| `fx_quote_id` | - | - | - | - | - |
| `fx_rate` | - | - | - | - | - |
| `fx_use_destination_amount` | - | - | - | - | - |
| `has_current_user_authorised` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `inserted` | - | - | - | - | - |
| `invoice_id` | - | - | - | - | - |
| `is_archived` | - | - | - | - | - |
| `is_failed` | - | - | - | - | - |
| `is_settled` | - | - | - | - | - |
| `is_submitted` | - | - | - | - | - |
| `last_updated` | - | - | - | - | - |
| `merchant_id` | - | - | - | - | - |
| `merchant_token_description` | - | - | - | - | - |
| `nonce` | - | - | - | - | - |
| `payment_processor` | - | - | - | - | - |
| `payment_rail` | - | - | - | - | - |
| `payout` | - | - | - | - | - |
| `payrun_id` | - | - | - | - | - |
| `payrun_name` | - | - | - | - | - |
| `reason` | - | - | - | - | - |
| `rule` | - | - | - | - | - |
| `schedule_date` | - | - | - | - | - |
| `scheduled` | - | - | - | - | - |
| `source_account_available_balance` | - | - | - | - | - |
| `source_account_available_balance_minor_unit` | - | - | - | - | - |
| `source_account_bic` | - | - | - | - | - |
| `source_account_currency` | - | - | - | - | - |
| `source_account_iban` | - | - | - | - | - |
| `source_account_identifier` | - | - | - | - | - |
| `source_account_name` | - | - | - | - | - |
| `source_account_number` | - | - | - | - | - |
| `source_account_sortcode` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `tag` | - | - | - | - | - |
| `tag_id` | - | - | - | - | - |
| `their_reference` | - | - | - | - | - |
| `topup_payrun_id` | - | - | - | - | - |
| `transacted_amount` | - | - | - | - | - |
| `transacted_fx_amount` | - | - | - | - | - |
| `transacted_fx_rate` | - | - | - | - | - |
| `type` | - | - | Yes | - | - |
| `user_id` | - | - | - | - | - |
| `your_reference` | - | - | - | - | - |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Payout.create({
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Payout.list
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

## PayoutKeysetPageEntity

```ruby
payout_keyset_page = client.PayoutKeysetPage
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
results = client.PayoutKeysetPage.list
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

Create a new `PayoutKeysetPageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PayoutMetricEntity

```ruby
payout_metric = client.PayoutMetric
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
result = client.PayoutMetric.load()
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

Create a new `PayoutMetricEntity` instance with the same client and
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
| `note` | `String` | No |  |
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
result = client.Payrun.create({
  "id" => "example_id", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Payrun.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Payrun.load({ "id" => "payrun_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Payrun.remove({ "id" => "payrun_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Payrun.update({
  "id" => "payrun_id",
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

## ReportResultEntity

```ruby
report_result = client.ReportResult
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
result = client.ReportResult.load({ "id" => 1, "report_id" => "report_id" })
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

Create a new `ReportResultEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## RoleEntity

```ruby
role = client.Role
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
result = client.Role.create({
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

Create a new `RoleEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## RuleEntity

```ruby
rule = client.Rule
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

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `account` | - | - | - | - | - |
| `account_id` | - | - | - | - | - |
| `approve_url` | - | - | - | - | - |
| `approver_id` | - | - | - | - | - |
| `authentication_method` | - | - | - | - | - |
| `authorisation` | - | - | - | - | - |
| `authorisers_completed_count` | - | - | - | - | - |
| `authorisers_required_count` | - | - | - | - | - |
| `can_authorise` | - | - | - | - | - |
| `created_by` | - | - | - | - | - |
| `description` | - | - | - | - | - |
| `end_at` | - | - | - | - | - |
| `has_current_user_authorised` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `inserted` | - | - | - | - | - |
| `is_disabled` | - | - | - | - | - |
| `last_executed_at` | - | - | - | - | - |
| `last_run_at_transaction_date` | - | - | - | - | - |
| `last_updated` | - | - | - | - | - |
| `merchant_id` | - | - | - | - | - |
| `name` | - | - | Yes | - | - |
| `nonce` | - | - | - | - | - |
| `on_approved_web_hook_url` | - | - | - | - | - |
| `on_execution_error_web_hook_url` | - | - | - | - | - |
| `on_execution_success_web_hook_url` | - | - | - | - | - |
| `start_at` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `sweep_action` | - | - | Yes | - | - |
| `time_zone_id` | - | - | - | - | - |
| `trigger_cron_expression` | - | - | - | - | - |
| `trigger_on_pay_in` | - | - | - | - | - |
| `user_id` | - | - | - | - | - |
| `web_hook_secret` | - | - | - | - | - |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Rule.create({
  "created_by" => {}, # Hash
  "nonce" => "example_nonce", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Rule.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Rule.load({ "id" => "rule_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Rule.remove({ "id" => "rule_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Rule.update({
  "id" => "rule_id",
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

## RuleEventEntity

```ruby
rule_event = client.RuleEvent
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
results = client.RuleEvent.list
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

Create a new `RuleEventEntity` instance with the same client and
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

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `String` | No |  |
| `account_name` | `String` | No |  |
| `account_sequence_number` | `Integer` | No |  |
| `address_detail` | `Hash` | No |  |
| `amount` | `Float` | No |  |
| `amount_minor_unit` | `Integer` | No |  |
| `balance` | `Float` | No |  |
| `balance_minor_unit` | `Integer` | No |  |
| `booking_date_time` | `String` | No |  |
| `charge_detail` | `Hash` | No |  |
| `content` | `Array` | No |  |
| `counterparty` | `Hash` | No |  |
| `counterparty_summary` | `String` | No |  |
| `currency` | `String` | No |  |
| `currency_exchange` | `Hash` | No |  |
| `date` | `String` | No |  |
| `description` | `String` | No |  |
| `enrichment` | `Hash` | No |  |
| `fx_amount` | `Float` | No |  |
| `fx_currency` | `String` | No |  |
| `fx_rate` | `Float` | No |  |
| `gross_amount` | `Hash` | Yes |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `iso_bank_transaction_code` | `Hash` | No |  |
| `merchant` | `Hash` | No |  |
| `merchant_id` | `String` | No |  |
| `page_number` | `Integer` | No |  |
| `page_size` | `Integer` | No |  |
| `payee_detail` | `Hash` | Yes |  |
| `payer_detail` | `Hash` | Yes |  |
| `payment_request_custom_field` | `Hash` | No |  |
| `payment_request_id` | `String` | No |  |
| `payout_id` | `String` | No |  |
| `proprietary_bank_transaction_code` | `Hash` | No |  |
| `raw_reference` | `String` | No |  |
| `reference` | `String` | No |  |
| `rule_id` | `String` | No |  |
| `statement_reference` | `Array` | No |  |
| `status` | `String` | No |  |
| `supplementary_data` | `Object` | No |  |
| `tag` | `Array` | No |  |
| `their_reference` | `String` | No |  |
| `total_page` | `Integer` | No |  |
| `total_size` | `Integer` | No |  |
| `transaction_amount` | `Hash` | Yes |  |
| `transaction_date` | `String` | No |  |
| `transaction_information` | `Array` | No |  |
| `transaction_mutability` | `String` | No |  |
| `type` | `String` | No |  |
| `value_date_time` | `String` | No |  |
| `virtual_iban` | `String` | No |  |
| `your_reference` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Transaction.create({
  "id" => "example_id", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Transaction.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Transaction.load({ "id" => "transaction_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Transaction.remove({ "id" => "transaction_id" })
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

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_status` | `Hash` | No |  |
| `failed_user_invite` | `Hash` | No |  |
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
| `user_invite` | `Array` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `authorisation_status` | - | - | - | - | - |
| `failed_user_invite` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `initial_role_id` | - | - | - | - | - |
| `invitee_email_address` | - | - | Yes | - | - |
| `invitee_first_name` | - | - | - | - | - |
| `invitee_last_name` | - | - | - | - | - |
| `inviter_email_address` | - | - | - | - | - |
| `inviter_first_name` | - | - | - | - | - |
| `inviter_last_name` | - | - | - | - | - |
| `is_authorised` | - | - | - | - | - |
| `is_invitee_registered` | - | - | - | - | - |
| `last_invited` | - | - | - | - | - |
| `merchant_id` | - | - | - | - | - |
| `merchant_name` | - | - | - | - | - |
| `message` | - | - | - | - | - |
| `registration_url` | - | - | - | - | - |
| `send_invite_email` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `user` | - | - | - | - | - |
| `user_id` | - | - | - | - | - |
| `user_invite` | - | - | - | - | - |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.UserInvite.create({
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.UserInvite.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.UserInvite.load({ "id" => "user_invite_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.UserInvite.remove({ "id" => "user_invite_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.UserInvite.update({
  "id" => "user_invite_id",
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

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `destination_url` | - | - | - | - | - |
| `email_address` | - | - | - | - | - |
| `failed_notification_email_address` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `is_active` | - | - | - | - | - |
| `merchant_id` | - | - | Yes | Yes | - |
| `notification_method` | - | - | Yes | Yes | - |
| `resource_type` | - | - | - | - | - |
| `retry` | - | - | - | - | - |
| `secret` | - | - | - | - | - |
| `version` | - | - | - | - | - |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Webhook.create({
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Webhook.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Webhook.load({ "id" => "webhook_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Webhook.remove({ "id" => "webhook_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Webhook.update({
  "id" => "webhook_id",
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

Create a new `WebhookEntity` instance with the same client and
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

