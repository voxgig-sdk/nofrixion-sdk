# Nofrixion Ruby SDK



The Ruby SDK for the Nofrixion API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Account` — with named operations (`list`/`load`/`create`/`update`/`remove`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/nofrixion-sdk/releases](https://github.com/voxgig-sdk/nofrixion-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "Nofrixion_sdk"

client = NofrixionSDK.new({
  "apikey" => ENV["NOFRIXION_APIKEY"],
})
```

### 2. List account records

```ruby
begin
  # list returns an Array of Account records — iterate directly.
  accounts = client.Account.list
  accounts.each do |item|
    puts "#{item["id"]} #{item["account_balance"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```

### 3. Load a cardcustomertoken

CardCustomerToken is nested under customer_email_address, so provide the `customer_email_address`.

```ruby
begin
  # load returns the bare CardCustomerToken record (raises on error).
  cardcustomertoken = client.CardCustomerToken.load({ "customer_email_address" => "example_customer_email_address" })
  puts cardcustomertoken
rescue => err
  warn "load failed: #{err}"
end
```

### 4. Create, update, and remove

```ruby
# create returns the bare created Account record.
created = client.Account.create({ "account_id" => "example_account_id", "currency" => "example_currency" })

# Update — index the bare record directly (created["id"]).
client.Account.update({ "id" => created["id"], "account_id" => "example_account_id", "amount" => 1 })

# Remove
client.Account.remove({ "id" => created["id"] })
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  accounts = client.Account.list()
rescue => err
  warn "list failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```ruby
client = NofrixionSDK.test({
  "entity" => { "account" => { "test01" => { "id" => "test01" } } },
})

# Entity ops return the bare mock record (raises on error).
account = client.Account.list()
puts account
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = NofrixionSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
  },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
NOFRIXION_TEST_LIVE=TRUE
NOFRIXION_APIKEY=<your-key>
```

Then run:

```bash
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### NofrixionSDK

```ruby
require_relative "Nofrixion_sdk"
client = NofrixionSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `String` | API key for authentication. |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = NofrixionSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### NofrixionSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
| `Account` | `(data) -> AccountEntity` | Create an Account entity instance. |
| `Batch` | `(data) -> BatchEntity` | Create a Batch entity instance. |
| `Beneficiary` | `(data) -> BeneficiaryEntity` | Create a Beneficiary entity instance. |
| `BeneficiaryGroup` | `(data) -> BeneficiaryGroupEntity` | Create a BeneficiaryGroup entity instance. |
| `Card` | `(data) -> CardEntity` | Create a Card entity instance. |
| `CardCustomerToken` | `(data) -> CardCustomerTokenEntity` | Create a CardCustomerToken entity instance. |
| `CardPayment` | `(data) -> CardPaymentEntity` | Create a CardPayment entity instance. |
| `CardPublicKey` | `(data) -> CardPublicKeyEntity` | Create a CardPublicKey entity instance. |
| `Consent` | `(data) -> ConsentEntity` | Create a Consent entity instance. |
| `Currency` | `(data) -> CurrencyEntity` | Create a Currency entity instance. |
| `DirectDebitBatchSubmit` | `(data) -> DirectDebitBatchSubmitEntity` | Create a DirectDebitBatchSubmit entity instance. |
| `FxRate` | `(data) -> FxRateEntity` | Create a FxRate entity instance. |
| `IPayment` | `(data) -> IPaymentEntity` | Create an IPayment entity instance. |
| `Mandate` | `(data) -> MandateEntity` | Create a Mandate entity instance. |
| `Merchant` | `(data) -> MerchantEntity` | Create a Merchant entity instance. |
| `MerchantAuthorisationSetting` | `(data) -> MerchantAuthorisationSettingEntity` | Create a MerchantAuthorisationSetting entity instance. |
| `MerchantDirectDebitMandatePage` | `(data) -> MerchantDirectDebitMandatePageEntity` | Create a MerchantDirectDebitMandatePage entity instance. |
| `MerchantPayByBankSetting` | `(data) -> MerchantPayByBankSettingEntity` | Create a MerchantPayByBankSetting entity instance. |
| `MerchantPaymentRequestTemplate` | `(data) -> MerchantPaymentRequestTemplateEntity` | Create a MerchantPaymentRequestTemplate entity instance. |
| `MerchantToken` | `(data) -> MerchantTokenEntity` | Create a MerchantToken entity instance. |
| `Metadata` | `(data) -> MetadataEntity` | Create a Metadata entity instance. |
| `NoFrixionVersion` | `(data) -> NoFrixionVersionEntity` | Create a NoFrixionVersion entity instance. |
| `OpenBanking` | `(data) -> OpenBankingEntity` | Create an OpenBanking entity instance. |
| `Payeeverification` | `(data) -> PayeeverificationEntity` | Create a Payeeverification entity instance. |
| `Payment` | `(data) -> PaymentEntity` | Create a Payment entity instance. |
| `PaymentAccount` | `(data) -> PaymentAccountEntity` | Create a PaymentAccount entity instance. |
| `PaymentAccountMinimal` | `(data) -> PaymentAccountMinimalEntity` | Create a PaymentAccountMinimal entity instance. |
| `PaymentInitiation` | `(data) -> PaymentInitiationEntity` | Create a PaymentInitiation entity instance. |
| `PaymentRequest` | `(data) -> PaymentRequestEntity` | Create a PaymentRequest entity instance. |
| `PaymentRequestEvent` | `(data) -> PaymentRequestEventEntity` | Create a PaymentRequestEvent entity instance. |
| `PaymentRequestMetric` | `(data) -> PaymentRequestMetricEntity` | Create a PaymentRequestMetric entity instance. |
| `PaymentRequestMinimal` | `(data) -> PaymentRequestMinimalEntity` | Create a PaymentRequestMinimal entity instance. |
| `PaymentRequestResult` | `(data) -> PaymentRequestResultEntity` | Create a PaymentRequestResult entity instance. |
| `Payout` | `(data) -> PayoutEntity` | Create a Payout entity instance. |
| `PayoutKeysetPage` | `(data) -> PayoutKeysetPageEntity` | Create a PayoutKeysetPage entity instance. |
| `PayoutMetric` | `(data) -> PayoutMetricEntity` | Create a PayoutMetric entity instance. |
| `Payrun` | `(data) -> PayrunEntity` | Create a Payrun entity instance. |
| `Report` | `(data) -> ReportEntity` | Create a Report entity instance. |
| `ReportResult` | `(data) -> ReportResultEntity` | Create a ReportResult entity instance. |
| `Role` | `(data) -> RoleEntity` | Create a Role entity instance. |
| `Rule` | `(data) -> RuleEntity` | Create a Rule entity instance. |
| `RuleEvent` | `(data) -> RuleEventEntity` | Create a RuleEvent entity instance. |
| `Tag` | `(data) -> TagEntity` | Create a Tag entity instance. |
| `Token` | `(data) -> TokenEntity` | Create a Token entity instance. |
| `Transaction` | `(data) -> TransactionEntity` | Create a Transaction entity instance. |
| `User` | `(data) -> UserEntity` | Create an User entity instance. |
| `UserInvite` | `(data) -> UserInviteEntity` | Create an UserInvite entity instance. |
| `Virtual` | `(data) -> VirtualEntity` | Create a Virtual entity instance. |
| `Webhook` | `(data) -> WebhookEntity` | Create a Webhook entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch = nil, ctrl) -> Array` | List entities matching the criteria (call with no argument to list all). Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `NofrixionError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

### Entities

#### Account

| Field | Description |
| --- | --- |
| `account_balance` |  |
| `account_id` |  |
| `account_identification` |  |
| `account_name` |  |
| `account_supplier_name` |  |
| `account_type` |  |
| `available_balance` |  |
| `available_balance_minor_unit` |  |
| `balance` |  |
| `balance_minor_unit` |  |
| `bank_name` |  |
| `consent_id` |  |
| `consolidated_account_information` |  |
| `created_by` |  |
| `created_by_display_name` |  |
| `currency` |  |
| `default_payment_rail` |  |
| `description` |  |
| `detail` |  |
| `display_name` |  |
| `expiry_date` |  |
| `external_account_icon` |  |
| `format` |  |
| `from_date` |  |
| `id` |  |
| `identifier` |  |
| `inserted` |  |
| `is_archived` |  |
| `is_connected_account` |  |
| `is_default` |  |
| `is_trust_account` |  |
| `is_virtual` |  |
| `last_transaction` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `merchant_name` |  |
| `nickname` |  |
| `physical_account_id` |  |
| `role_i_d` |  |
| `rule` |  |
| `submitted_payouts_balance` |  |
| `submitted_payouts_balance_minor_unit` |  |
| `summary` |  |
| `supplier_physical_account_id` |  |
| `supplier_sepa_instant_status` |  |
| `to_date` |  |
| `type` |  |
| `usage_type` |  |
| `xero_bank_feed_connection_status` |  |
| `xero_bank_feed_last_synced_at` |  |
| `xero_bank_feed_sync_last_failed_at` |  |
| `xero_bank_feed_sync_last_failure_reason` |  |
| `xero_bank_feed_sync_status` |  |
| `xero_unsynchronised_transactions_count` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/accounts/{accountID}/{currency}`

#### Batch

| Field | Description |
| --- | --- |
| `approve_url` |  |
| `id` |  |
| `payout` |  |

Operations: Create, Load.

API path: `/api/v1/payouts/batch`

#### Beneficiary

| Field | Description |
| --- | --- |
| `approval_callback_url` |  |
| `authentication_method` |  |
| `authorisation` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `beneficiary` |  |
| `beneficiary_event` |  |
| `can_authorise` |  |
| `can_update` |  |
| `created_by` |  |
| `created_by_email_address` |  |
| `currency` |  |
| `destination` |  |
| `failed_beneficiary` |  |
| `has_current_user_authorised` |  |
| `id` |  |
| `inserted` |  |
| `is_enabled` |  |
| `last_authorised` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `name` |  |
| `nonce` |  |
| `source_account` |  |
| `source_account_i_d` |  |
| `their_reference` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/beneficiaries/authorise/{id}`

#### BeneficiaryGroup

| Field | Description |
| --- | --- |
| `group_member` |  |
| `group_name` |  |
| `id` |  |
| `inserted` |  |
| `last_updated` |  |
| `merchant_id` |  |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/beneficiarygroups`

#### Card

| Field | Description |
| --- | --- |
| `authorized_amount` |  |
| `currency_code` |  |
| `is_payer_authentication_required` |  |
| `is_soft_decline` |  |
| `payer_authentication_access_token` |  |
| `payer_authentication_merchant_data` |  |
| `payer_authentication_url` |  |
| `payer_authentication_window_height` |  |
| `payer_authentication_window_width` |  |
| `payment_request_callback_url` |  |
| `payment_request_id` |  |
| `request_id` |  |
| `response_code` |  |
| `response_type` |  |
| `status` |  |
| `three_ds_redirect_url` |  |
| `transaction_id` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/{id}/card`

#### CardCustomerToken

| Field | Description |
| --- | --- |
| `card_type` |  |
| `customer_email_address` |  |
| `expiry_month` |  |
| `expiry_year` |  |
| `id` |  |
| `inserted` |  |
| `last_four_digit` |  |
| `last_updated` |  |
| `masked_card_number` |  |
| `merchant_id` |  |
| `payment_request_id` |  |

Operations: List, Load, Remove.

API path: `/api/v1/paymentrequests/card/customertokens/{merchantID}/{customerEmailAddress}`

#### CardPayment

| Field | Description |
| --- | --- |
| `authorized_amount` |  |
| `currency_code` |  |
| `is_payer_authentication_required` |  |
| `is_soft_decline` |  |
| `payer_authentication_access_token` |  |
| `payer_authentication_merchant_data` |  |
| `payer_authentication_url` |  |
| `payer_authentication_window_height` |  |
| `payer_authentication_window_width` |  |
| `payment_request_callback_url` |  |
| `payment_request_id` |  |
| `request_id` |  |
| `response_code` |  |
| `response_type` |  |
| `status` |  |
| `three_ds_redirect_url` |  |
| `transaction_id` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/{id}/card/refund/{partialRefundAmount}`

#### CardPublicKey

| Field | Description |
| --- | --- |
| `jwt` |  |

Operations: Load.

API path: `/api/v1/paymentrequests/{id}/card/publickey`

#### Consent

| Field | Description |
| --- | --- |
| `authorisation_url` |  |
| `callback_url` |  |
| `consent_id` |  |
| `email_address` |  |
| `expiry_date` |  |
| `failure_callback_url` |  |
| `id` |  |
| `inserted` |  |
| `institution_id` |  |
| `is_connected_account` |  |
| `is_enabled` |  |
| `merchant_id` |  |
| `provider` |  |
| `success_web_hook_url` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/openbanking/consents`

#### Currency

| Field | Description |
| --- | --- |
| `code` |  |
| `decimal` |  |
| `is_fiat` |  |
| `iso4217_alpha_code` |  |
| `iso4217_numeric_code` |  |
| `symbol` |  |

Operations: List.

API path: `/api/v1/currencies`

#### DirectDebitBatchSubmit

| Field | Description |
| --- | --- |
| `failed_submission` |  |
| `successful_submission` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/directdebit/batchsubmit`

#### FxRate

| Field | Description |
| --- | --- |
| `destination_currency` |  |
| `exchange_rate` |  |
| `expiry_time` |  |
| `quote_id` |  |
| `source_currency` |  |

Operations: List, Load.

API path: `/api/v1/payouts/fxallheldrates/{source}/{destination}`

#### IPayment

| Field | Description |
| --- | --- |
| `payment_request_id` |  |
| `response_type` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/payondemand`

#### Mandate

| Field | Description |
| --- | --- |
| `account_number` |  |
| `address_line1` |  |
| `address_line2` |  |
| `approved_at` |  |
| `city` |  |
| `country_code` |  |
| `currency` |  |
| `customer_account_number` |  |
| `customer_city` |  |
| `customer_country_code` |  |
| `customer_country_name` |  |
| `customer_email_address` |  |
| `customer_first_name` |  |
| `customer_iban` |  |
| `customer_last_name` |  |
| `customer_sort_code` |  |
| `email_address` |  |
| `first_name` |  |
| `iban` |  |
| `id` |  |
| `inserted` |  |
| `is_recurring` |  |
| `last_name` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `postal_code` |  |
| `reference` |  |
| `sort_code` |  |
| `status` |  |
| `supplier_bank_account_id` |  |
| `supplier_customer_id` |  |
| `supplier_mandate_id` |  |
| `supplier_name` |  |
| `supplier_status` |  |

Operations: Create, Load.

API path: `/api/v1/mandates`

#### Merchant

| Field | Description |
| --- | --- |
| `account_currency` |  |
| `can_have_trust_account` |  |
| `card_payment_processor` |  |
| `company_id` |  |
| `display_qr_on_hosted_pay` |  |
| `hosted_pay_version` |  |
| `id` |  |
| `inserted` |  |
| `is_blocked` |  |
| `is_exited` |  |
| `is_suspended` |  |
| `jurisdiction` |  |
| `logo_url_png` |  |
| `logo_url_svg` |  |
| `merchant_category_code` |  |
| `name` |  |
| `note` |  |
| `parent_merchant` |  |
| `payment_account` |  |
| `payment_account_limit` |  |
| `reason` |  |
| `short_name` |  |
| `supported_payment_methods_list` |  |
| `suspension_reason` |  |
| `tag` |  |
| `time_zone_id` |  |
| `trading_name` |  |
| `web_hook_limit` |  |
| `your_role_name` |  |

Operations: List, Load, Remove, Update.

API path: `/api/v1/merchants/{merchantID}/childmerchants`

#### MerchantAuthorisationSetting

| Field | Description |
| --- | --- |
| `amount_lower` |  |
| `amount_upper` |  |
| `authorisation_type` |  |
| `beneficiaries_only` |  |
| `id` |  |
| `inserted` |  |
| `last_editor_cant_authorise` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `number_of_authoriser` |  |
| `role_setting` |  |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/authorisationsettings`

#### MerchantDirectDebitMandatePage

| Field | Description |
| --- | --- |
| `approved_at` |  |
| `currency` |  |
| `customer_account_number` |  |
| `customer_city` |  |
| `customer_country_code` |  |
| `customer_country_name` |  |
| `customer_email_address` |  |
| `customer_first_name` |  |
| `customer_iban` |  |
| `customer_last_name` |  |
| `customer_sort_code` |  |
| `id` |  |
| `inserted` |  |
| `is_recurring` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `reference` |  |
| `status` |  |
| `supplier_bank_account_id` |  |
| `supplier_customer_id` |  |
| `supplier_mandate_id` |  |
| `supplier_name` |  |
| `supplier_status` |  |

Operations: List.

API path: `/api/v1/mandates`

#### MerchantPayByBankSetting

| Field | Description |
| --- | --- |
| `bank_country_code` |  |
| `bank_id` |  |
| `bank_name` |  |
| `business_institution_id` |  |
| `currency` |  |
| `logo` |  |
| `message` |  |
| `message_image_url` |  |
| `order` |  |
| `personal_institution_id` |  |
| `processor` |  |
| `warning_heading` |  |
| `warning_message` |  |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/banksettings`

#### MerchantPaymentRequestTemplate

| Field | Description |
| --- | --- |
| `description` |  |
| `id` |  |
| `inserted` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `name` |  |
| `template` |  |

Operations: List, Load, Remove, Update.

API path: `/api/v1/paymentrequests/{merchantID}/templates`

#### MerchantToken

| Field | Description |
| --- | --- |
| `authentication_method` |  |
| `authorisation` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `can_authorise` |  |
| `description` |  |
| `expires_at` |  |
| `has_current_user_authorised` |  |
| `hmac_algorithm` |  |
| `id` |  |
| `inserted` |  |
| `ip_address_whitelist` |  |
| `is_archived` |  |
| `is_enabled` |  |
| `last_authorised` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `nonce` |  |
| `permission_type` |  |
| `request_signature_version` |  |
| `shared_secret_algorithm` |  |
| `shared_secret_base64` |  |
| `token` |  |

Operations: Create, List, Load, Update.

API path: `/api/v1/tokens`

#### Metadata

| Field | Description |
| --- | --- |

Operations: Load.

API path: `/api/v1/metadata/problemnotification`

#### NoFrixionVersion

| Field | Description |
| --- | --- |
| `build_version` |  |
| `major_version` |  |
| `minor_version` |  |
| `release_name` |  |

Operations: Load.

API path: `/api/v1/metadata/version`

#### OpenBanking

| Field | Description |
| --- | --- |

Operations: Create, Remove.

API path: `/api/v1/openbanking/account/{accountID}/synchronise`

#### Payeeverification

| Field | Description |
| --- | --- |
| `account_name` |  |
| `account_number` |  |
| `iban` |  |
| `payee_verified_account_name` |  |
| `result` |  |
| `secondary_identification` |  |
| `sort_code` |  |

Operations: Create.

API path: `/api/v1/openbanking/payeeverification`

#### Payment

| Field | Description |
| --- | --- |
| `address` |  |
| `amount` |  |
| `amount_pending` |  |
| `amount_received` |  |
| `amount_refunded` |  |
| `auto_send_receipt` |  |
| `base_origin_url` |  |
| `callback_url` |  |
| `card_authorize_only` |  |
| `card_create_token` |  |
| `card_create_token_mode` |  |
| `card_ignore_cvn` |  |
| `card_no_payer_authentication` |  |
| `card_processor_merchant_id` |  |
| `card_stripe_payment_intent_id` |  |
| `card_stripe_payment_intent_secret` |  |
| `card_transmit_raw_detail` |  |
| `created_by_user` |  |
| `currency` |  |
| `custom_field` |  |
| `customer_email_address` |  |
| `customer_id` |  |
| `customer_name` |  |
| `description` |  |
| `destination_account` |  |
| `direct_debit_payment` |  |
| `due_date` |  |
| `event` |  |
| `failure_callback_url` |  |
| `field_display_setting` |  |
| `formatted_amount` |  |
| `hosted_pay_checkout_url` |  |
| `id` |  |
| `ignore_address_verification` |  |
| `inserted` |  |
| `inserted_sortable` |  |
| `is_archived` |  |
| `jwk` |  |
| `last_updated` |  |
| `lightning_invoice` |  |
| `lightning_invoice_expires_at` |  |
| `merchant_direct_debit_mandate_id` |  |
| `merchant_id` |  |
| `merchant_token_description` |  |
| `notification_email_address` |  |
| `notification_role_i_d` |  |
| `order_id` |  |
| `partial_payment_method` |  |
| `partial_payment_step` |  |
| `payment_attempt` |  |
| `payment_method` |  |
| `payment_processor` |  |
| `payrun_id` |  |
| `pisp_account_id` |  |
| `priority_bank_id` |  |
| `result` |  |
| `sandbox_settle_delay_in_second` |  |
| `shipping_address` |  |
| `shipping_address_city` |  |
| `shipping_address_country_code` |  |
| `shipping_address_county` |  |
| `shipping_address_line1` |  |
| `shipping_address_line2` |  |
| `shipping_address_post_code` |  |
| `shipping_email` |  |
| `shipping_first_name` |  |
| `shipping_last_name` |  |
| `shipping_phone` |  |
| `status` |  |
| `success_web_hook_url` |  |
| `tag` |  |
| `tag_id` |  |
| `title` |  |
| `tokenised_card` |  |
| `transaction` |  |
| `use_hosted_payment_page` |  |

Operations: Create, Load, Update.

API path: `/api/v1/paymentrequests`

#### PaymentAccount

| Field | Description |
| --- | --- |
| `account_name` |  |
| `account_supplier_name` |  |
| `available_balance` |  |
| `available_balance_minor_unit` |  |
| `balance` |  |
| `balance_minor_unit` |  |
| `bank_name` |  |
| `consent_id` |  |
| `created_by` |  |
| `created_by_display_name` |  |
| `currency` |  |
| `default_payment_rail` |  |
| `display_name` |  |
| `expiry_date` |  |
| `external_account_icon` |  |
| `id` |  |
| `identifier` |  |
| `inserted` |  |
| `is_archived` |  |
| `is_connected_account` |  |
| `is_default` |  |
| `is_trust_account` |  |
| `is_virtual` |  |
| `last_transaction` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `merchant_name` |  |
| `physical_account_id` |  |
| `rule` |  |
| `submitted_payouts_balance` |  |
| `submitted_payouts_balance_minor_unit` |  |
| `summary` |  |
| `supplier_sepa_instant_status` |  |
| `xero_bank_feed_connection_status` |  |
| `xero_bank_feed_last_synced_at` |  |
| `xero_bank_feed_sync_last_failed_at` |  |
| `xero_bank_feed_sync_last_failure_reason` |  |
| `xero_bank_feed_sync_status` |  |
| `xero_unsynchronised_transactions_count` |  |

Operations: List.

API path: `/api/v1/accounts/paged`

#### PaymentAccountMinimal

| Field | Description |
| --- | --- |
| `account_name` |  |
| `available_balance` |  |
| `balance` |  |
| `balance_minor_unit` |  |
| `currency` |  |
| `id` |  |
| `identifier` |  |
| `is_archived` |  |
| `is_connected_account` |  |
| `merchant_id` |  |
| `submitted_payouts_balance` |  |

Operations: List.

API path: `/api/v1/accounts/minimal`

#### PaymentInitiation

| Field | Description |
| --- | --- |
| `payment_initiation_id` |  |
| `payment_request_callback_url` |  |
| `payment_request_id` |  |
| `redirect_url` |  |
| `response_type` |  |
| `specific_error_message` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/{id}/pisp`

#### PaymentRequest

| Field | Description |
| --- | --- |
| `address` |  |
| `amount` |  |
| `amount_pending` |  |
| `amount_received` |  |
| `amount_refunded` |  |
| `auto_send_receipt` |  |
| `base_origin_url` |  |
| `callback_url` |  |
| `card_authorize_only` |  |
| `card_create_token` |  |
| `card_create_token_mode` |  |
| `card_ignore_cvn` |  |
| `card_processor_merchant_id` |  |
| `card_stripe_payment_intent_id` |  |
| `card_stripe_payment_intent_secret` |  |
| `created_by_user` |  |
| `currency` |  |
| `custom_field` |  |
| `customer_email_address` |  |
| `customer_id` |  |
| `customer_name` |  |
| `description` |  |
| `destination_account` |  |
| `direct_debit_payment` |  |
| `do_simulate_settlement_failure` |  |
| `due_date` |  |
| `error_description` |  |
| `event` |  |
| `failed_payment_request` |  |
| `failure_callback_url` |  |
| `field_display_setting` |  |
| `formatted_amount` |  |
| `hosted_pay_checkout_url` |  |
| `id` |  |
| `ignore_address_verification` |  |
| `inserted` |  |
| `inserted_sortable` |  |
| `institution` |  |
| `is_archived` |  |
| `jwk` |  |
| `last_updated` |  |
| `lightning_invoice` |  |
| `lightning_invoice_expires_at` |  |
| `merchant_direct_debit_mandate_id` |  |
| `merchant_id` |  |
| `merchant_token_description` |  |
| `notification_email_address` |  |
| `notification_role_i_d` |  |
| `order_id` |  |
| `partial_payment_method` |  |
| `partial_payment_step` |  |
| `payment_attempt` |  |
| `payment_initiation_id` |  |
| `payment_method` |  |
| `payment_processor` |  |
| `payment_request` |  |
| `payrun_id` |  |
| `pisp_account_id` |  |
| `priority_bank_id` |  |
| `result` |  |
| `sandbox_settle_delay_in_second` |  |
| `shipping_address` |  |
| `status` |  |
| `success_web_hook_url` |  |
| `tag` |  |
| `title` |  |
| `tokenised_card` |  |
| `transaction` |  |
| `use_hosted_payment_page` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/paymentrequests/{id}/directdebit`

#### PaymentRequestEvent

| Field | Description |
| --- | --- |
| `amount` |  |
| `apple_pay_transaction_id` |  |
| `card_authorization_response_id` |  |
| `card_expiry_month` |  |
| `card_expiry_year` |  |
| `card_issuer` |  |
| `card_issuer_country` |  |
| `card_last_four_digit` |  |
| `card_request_id` |  |
| `card_scheme` |  |
| `card_token_customer_id` |  |
| `card_transaction_id` |  |
| `currency` |  |
| `direct_debit_payment_id` |  |
| `direct_debit_payment_reference` |  |
| `drirect_debit_mandate_id` |  |
| `error_message` |  |
| `error_reason` |  |
| `event_type` |  |
| `id` |  |
| `inserted` |  |
| `lightning_invoice` |  |
| `lightning_r_hash` |  |
| `origin_url` |  |
| `payment_method_type` |  |
| `payment_processor_name` |  |
| `payment_request_id` |  |
| `pisp_bank_status` |  |
| `pisp_payment_initiation_id` |  |
| `pisp_payment_institution_name` |  |
| `pisp_payment_service_provider_id` |  |
| `pisp_redirect_url` |  |
| `reconciled_transaction_id` |  |
| `refund_payout_id` |  |
| `status` |  |
| `wallet_name` |  |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/events`

#### PaymentRequestMetric

| Field | Description |
| --- | --- |
| `all` |  |
| `authorized` |  |
| `paid` |  |
| `partially_paid` |  |
| `total_amounts_by_currency` |  |
| `unpaid` |  |

Operations: Load.

API path: `/api/v1/paymentrequests/metrics`

#### PaymentRequestMinimal

| Field | Description |
| --- | --- |
| `amount` |  |
| `amount_pending` |  |
| `amount_received` |  |
| `amount_refunded` |  |
| `callback_url` |  |
| `card_stripe_payment_intent_secret` |  |
| `country_code` |  |
| `currency` |  |
| `custom_fields_to_display` |  |
| `description` |  |
| `due_date` |  |
| `field_display_setting` |  |
| `google_pay_merchant_id` |  |
| `id` |  |
| `jwk` |  |
| `merchant_id` |  |
| `merchant_logo_url_png` |  |
| `merchant_logo_url_svg` |  |
| `merchant_name` |  |
| `merchant_short_name` |  |
| `partial_payment_method` |  |
| `payment_attempt` |  |
| `payment_methods_list` |  |
| `payment_processor` |  |
| `payment_processor_key` |  |
| `pisp_error` |  |
| `priority_bank_id` |  |
| `status` |  |
| `stripe_account_id` |  |
| `title` |  |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/minimal`

#### PaymentRequestResult

| Field | Description |
| --- | --- |
| `amount` |  |
| `amount_pending` |  |
| `amount_received` |  |
| `amount_refunded` |  |
| `currency` |  |
| `customer_id` |  |
| `payment` |  |
| `payment_request_id` |  |
| `pisp_authorization` |  |
| `requested_amount` |  |
| `result` |  |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/result`

#### Payout

| Field | Description |
| --- | --- |
| `account_id` |  |
| `allow_incomplete` |  |
| `amount` |  |
| `amount_minor_unit` |  |
| `approve_payout_url` |  |
| `approver_id` |  |
| `authentication_method` |  |
| `authorisation` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `batch_payout_id` |  |
| `beneficiary` |  |
| `beneficiary_id` |  |
| `can_authorise` |  |
| `can_process` |  |
| `can_update` |  |
| `charge_bearer` |  |
| `created_by` |  |
| `created_by_email_address` |  |
| `currency` |  |
| `current_user_id` |  |
| `description` |  |
| `destination` |  |
| `document` |  |
| `event` |  |
| `failed_payout` |  |
| `formatted_amount` |  |
| `formatted_fx_destination_amount` |  |
| `formatted_schedule` |  |
| `formatted_schedule_day_only` |  |
| `formatted_source_account_available_balance` |  |
| `fx_destination_amount` |  |
| `fx_destination_amount_minor_unit` |  |
| `fx_destination_currency` |  |
| `fx_quote_expires_at` |  |
| `fx_quote_id` |  |
| `fx_rate` |  |
| `fx_use_destination_amount` |  |
| `has_current_user_authorised` |  |
| `id` |  |
| `inserted` |  |
| `invoice_id` |  |
| `is_archived` |  |
| `is_failed` |  |
| `is_settled` |  |
| `is_submitted` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `merchant_token_description` |  |
| `nonce` |  |
| `payment_processor` |  |
| `payment_rail` |  |
| `payout` |  |
| `payrun_id` |  |
| `payrun_name` |  |
| `reason` |  |
| `rule` |  |
| `schedule_date` |  |
| `scheduled` |  |
| `source_account_available_balance` |  |
| `source_account_available_balance_minor_unit` |  |
| `source_account_bic` |  |
| `source_account_currency` |  |
| `source_account_iban` |  |
| `source_account_identifier` |  |
| `source_account_name` |  |
| `source_account_number` |  |
| `source_account_sortcode` |  |
| `status` |  |
| `tag` |  |
| `tag_id` |  |
| `their_reference` |  |
| `topup_payrun_id` |  |
| `transacted_amount` |  |
| `transacted_fx_amount` |  |
| `transacted_fx_rate` |  |
| `type` |  |
| `user_id` |  |
| `your_reference` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/payouts/batch/submit/{id}`

#### PayoutKeysetPage

| Field | Description |
| --- | --- |
| `account_id` |  |
| `amount` |  |
| `amount_minor_unit` |  |
| `approve_payout_url` |  |
| `approver_id` |  |
| `authentication_method` |  |
| `authorisation` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `batch_payout_id` |  |
| `beneficiary` |  |
| `can_authorise` |  |
| `can_process` |  |
| `can_update` |  |
| `charge_bearer` |  |
| `created_by` |  |
| `created_by_email_address` |  |
| `currency` |  |
| `current_user_id` |  |
| `description` |  |
| `destination` |  |
| `document` |  |
| `event` |  |
| `formatted_amount` |  |
| `formatted_fx_destination_amount` |  |
| `formatted_schedule` |  |
| `formatted_schedule_day_only` |  |
| `formatted_source_account_available_balance` |  |
| `fx_destination_amount` |  |
| `fx_destination_amount_minor_unit` |  |
| `fx_destination_currency` |  |
| `fx_quote_expires_at` |  |
| `fx_quote_id` |  |
| `fx_rate` |  |
| `fx_use_destination_amount` |  |
| `has_current_user_authorised` |  |
| `id` |  |
| `inserted` |  |
| `invoice_id` |  |
| `is_archived` |  |
| `is_failed` |  |
| `is_settled` |  |
| `is_submitted` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `merchant_token_description` |  |
| `nonce` |  |
| `payment_processor` |  |
| `payment_rail` |  |
| `payrun_id` |  |
| `payrun_name` |  |
| `rule` |  |
| `schedule_date` |  |
| `scheduled` |  |
| `source_account_available_balance` |  |
| `source_account_available_balance_minor_unit` |  |
| `source_account_bic` |  |
| `source_account_currency` |  |
| `source_account_iban` |  |
| `source_account_identifier` |  |
| `source_account_name` |  |
| `source_account_number` |  |
| `source_account_sortcode` |  |
| `status` |  |
| `tag` |  |
| `their_reference` |  |
| `topup_payrun_id` |  |
| `transacted_amount` |  |
| `transacted_fx_amount` |  |
| `transacted_fx_rate` |  |
| `type` |  |
| `user_id` |  |
| `your_reference` |  |

Operations: List.

API path: `/api/v1/accounts/{accountID}/payouts/failed`

#### PayoutMetric

| Field | Description |
| --- | --- |
| `all` |  |
| `failed` |  |
| `in_progress` |  |
| `paid` |  |
| `pending_approval` |  |
| `scheduled` |  |
| `total_amounts_by_currency` |  |

Operations: Load.

API path: `/api/v1/payouts/metrics`

#### Payrun

| Field | Description |
| --- | --- |
| `authorisation` |  |
| `authorisation_date` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `batch_payout_id` |  |
| `can_authorise` |  |
| `can_delete` |  |
| `can_edit` |  |
| `event` |  |
| `has_current_user_authorised` |  |
| `id` |  |
| `inserted` |  |
| `invoice` |  |
| `invoices_minimal` |  |
| `is_archived` |  |
| `last_updated` |  |
| `last_updated_by` |  |
| `merchant_id` |  |
| `name` |  |
| `nonce` |  |
| `note` |  |
| `payment` |  |
| `payout` |  |
| `payouts_count` |  |
| `reason` |  |
| `schedule_date` |  |
| `scheduled_date` |  |
| `source_account` |  |
| `status` |  |
| `total_eur` |  |
| `total_gbp` |  |
| `total_usd` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/payruns/{id}/request-authorisation`

#### Report

| Field | Description |
| --- | --- |

Operations: Update.

API path: `/api/v1/reports/{id}/initiate`

#### ReportResult

| Field | Description |
| --- | --- |
| `content` |  |
| `content_type` |  |
| `last_completed_at` |  |
| `merchant_id` |  |
| `report_name` |  |
| `report_type` |  |
| `statement_number` |  |

Operations: Load.

API path: `/api/v1/reports/{id}/result/{statementNumber}`

#### Role

| Field | Description |
| --- | --- |
| `failed_role` |  |
| `role` |  |

Operations: Create.

API path: `/api/v1/merchants/{merchantID}/roles/batchcreate`

#### Rule

| Field | Description |
| --- | --- |
| `account` |  |
| `account_id` |  |
| `approve_url` |  |
| `approver_id` |  |
| `authentication_method` |  |
| `authorisation` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `can_authorise` |  |
| `created_by` |  |
| `description` |  |
| `end_at` |  |
| `has_current_user_authorised` |  |
| `id` |  |
| `inserted` |  |
| `is_disabled` |  |
| `last_executed_at` |  |
| `last_run_at_transaction_date` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `name` |  |
| `nonce` |  |
| `on_approved_web_hook_url` |  |
| `on_execution_error_web_hook_url` |  |
| `on_execution_success_web_hook_url` |  |
| `start_at` |  |
| `status` |  |
| `sweep_action` |  |
| `time_zone_id` |  |
| `trigger_cron_expression` |  |
| `trigger_on_pay_in` |  |
| `user_id` |  |
| `web_hook_secret` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/rules`

#### RuleEvent

| Field | Description |
| --- | --- |
| `error_message` |  |
| `id` |  |
| `inserted` |  |
| `is_authorise_to_enable` |  |
| `message` |  |
| `raw_response` |  |
| `rule_event_type` |  |
| `rule_id` |  |
| `user` |  |

Operations: List.

API path: `/api/v1/rules/{id}/events`

#### Tag

| Field | Description |
| --- | --- |
| `colour_hex` |  |
| `description` |  |
| `id` |  |
| `merchant_id` |  |
| `name` |  |

Operations: Create, List.

API path: `/api/v1/merchants/{merchantID}/tags`

#### Token

| Field | Description |
| --- | --- |

Operations: Create, Remove.

API path: `/api/v1/tokens/authorise/{id}`

#### Transaction

| Field | Description |
| --- | --- |
| `account_id` |  |
| `account_name` |  |
| `account_sequence_number` |  |
| `address_detail` |  |
| `amount` |  |
| `amount_minor_unit` |  |
| `balance` |  |
| `balance_minor_unit` |  |
| `booking_date_time` |  |
| `charge_detail` |  |
| `content` |  |
| `counterparty` |  |
| `counterparty_summary` |  |
| `currency` |  |
| `currency_exchange` |  |
| `date` |  |
| `description` |  |
| `enrichment` |  |
| `fx_amount` |  |
| `fx_currency` |  |
| `fx_rate` |  |
| `gross_amount` |  |
| `id` |  |
| `inserted` |  |
| `iso_bank_transaction_code` |  |
| `merchant` |  |
| `merchant_id` |  |
| `page_number` |  |
| `page_size` |  |
| `payee_detail` |  |
| `payer_detail` |  |
| `payment_request_custom_field` |  |
| `payment_request_id` |  |
| `payout_id` |  |
| `proprietary_bank_transaction_code` |  |
| `raw_reference` |  |
| `reference` |  |
| `rule_id` |  |
| `statement_reference` |  |
| `status` |  |
| `supplementary_data` |  |
| `tag` |  |
| `their_reference` |  |
| `total_page` |  |
| `total_size` |  |
| `transaction_amount` |  |
| `transaction_date` |  |
| `transaction_information` |  |
| `transaction_mutability` |  |
| `type` |  |
| `value_date_time` |  |
| `virtual_iban` |  |
| `your_reference` |  |

Operations: Create, List, Load, Remove.

API path: `/api/v1/transactions/{id}/tags`

#### User

| Field | Description |
| --- | --- |
| `client_session_timeout` |  |
| `email_address` |  |
| `first_name` |  |
| `id` |  |
| `last_name` |  |
| `passkey_added` |  |
| `permission` |  |
| `profile` |  |
| `roles_with_scope` |  |
| `two_factor_enabled` |  |
| `user_invite_id` |  |

Operations: List, Update.

API path: `/api/v1/user/{merchantID}/userspaged`

#### UserInvite

| Field | Description |
| --- | --- |
| `authorisation_status` |  |
| `failed_user_invite` |  |
| `id` |  |
| `initial_role_id` |  |
| `invitee_email_address` |  |
| `invitee_first_name` |  |
| `invitee_last_name` |  |
| `inviter_email_address` |  |
| `inviter_first_name` |  |
| `inviter_last_name` |  |
| `is_authorised` |  |
| `is_invitee_registered` |  |
| `last_invited` |  |
| `merchant_id` |  |
| `merchant_name` |  |
| `message` |  |
| `registration_url` |  |
| `send_invite_email` |  |
| `status` |  |
| `user` |  |
| `user_id` |  |
| `user_invite` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/userinvites/authorise/{id}`

#### Virtual

| Field | Description |
| --- | --- |
| `account_name` |  |
| `account_supplier_name` |  |
| `available_balance` |  |
| `available_balance_minor_unit` |  |
| `balance` |  |
| `balance_minor_unit` |  |
| `bank_name` |  |
| `consent_id` |  |
| `created_by` |  |
| `created_by_display_name` |  |
| `currency` |  |
| `default_payment_rail` |  |
| `display_name` |  |
| `expiry_date` |  |
| `external_account_icon` |  |
| `id` |  |
| `identifier` |  |
| `inserted` |  |
| `is_archived` |  |
| `is_connected_account` |  |
| `is_default` |  |
| `is_trust_account` |  |
| `is_virtual` |  |
| `last_transaction` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `merchant_name` |  |
| `name` |  |
| `physical_account_id` |  |
| `rule` |  |
| `submitted_payouts_balance` |  |
| `submitted_payouts_balance_minor_unit` |  |
| `summary` |  |
| `supplier_sepa_instant_status` |  |
| `xero_bank_feed_connection_status` |  |
| `xero_bank_feed_last_synced_at` |  |
| `xero_bank_feed_sync_last_failed_at` |  |
| `xero_bank_feed_sync_last_failure_reason` |  |
| `xero_bank_feed_sync_status` |  |
| `xero_unsynchronised_transactions_count` |  |

Operations: Create, Update.

API path: `/api/v1/accounts/{accountID}/virtual`

#### Webhook

| Field | Description |
| --- | --- |
| `destination_url` |  |
| `email_address` |  |
| `failed_notification_email_address` |  |
| `id` |  |
| `is_active` |  |
| `merchant_id` |  |
| `notification_method` |  |
| `resource_type` |  |
| `retry` |  |
| `secret` |  |
| `version` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/webhooks`



## Entities


### Account

Create an instance: `account = client.Account`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_balance` | `Array` |  |
| `account_id` | `String` |  |
| `account_identification` | `Array` |  |
| `account_name` | `String` |  |
| `account_supplier_name` | `String` |  |
| `account_type` | `String` |  |
| `available_balance` | `Float` |  |
| `available_balance_minor_unit` | `Integer` |  |
| `balance` | `Float` |  |
| `balance_minor_unit` | `Integer` |  |
| `bank_name` | `String` |  |
| `consent_id` | `String` |  |
| `consolidated_account_information` | `Hash` |  |
| `created_by` | `Hash` |  |
| `created_by_display_name` | `String` |  |
| `currency` | `String` |  |
| `default_payment_rail` | `String` |  |
| `description` | `String` |  |
| `detail` | `String` |  |
| `display_name` | `String` |  |
| `expiry_date` | `String` |  |
| `external_account_icon` | `String` |  |
| `format` | `String` |  |
| `from_date` | `String` |  |
| `id` | `String` |  |
| `identifier` | `Hash` |  |
| `inserted` | `String` |  |
| `is_archived` | `Boolean` |  |
| `is_connected_account` | `Boolean` |  |
| `is_default` | `Boolean` |  |
| `is_trust_account` | `Boolean` |  |
| `is_virtual` | `Boolean` |  |
| `last_transaction` | `Hash` |  |
| `last_updated` | `String` |  |
| `merchant_id` | `String` |  |
| `merchant_name` | `String` |  |
| `nickname` | `String` |  |
| `physical_account_id` | `String` |  |
| `role_i_d` | `Array` |  |
| `rule` | `Array` |  |
| `submitted_payouts_balance` | `Float` |  |
| `submitted_payouts_balance_minor_unit` | `Integer` |  |
| `summary` | `String` |  |
| `supplier_physical_account_id` | `String` |  |
| `supplier_sepa_instant_status` | `String` |  |
| `to_date` | `String` |  |
| `type` | `String` |  |
| `usage_type` | `String` |  |
| `xero_bank_feed_connection_status` | `String` |  |
| `xero_bank_feed_last_synced_at` | `String` |  |
| `xero_bank_feed_sync_last_failed_at` | `String` |  |
| `xero_bank_feed_sync_last_failure_reason` | `String` |  |
| `xero_bank_feed_sync_status` | `String` |  |
| `xero_unsynchronised_transactions_count` | `Integer` |  |

#### Example: Load

```ruby
# load returns the bare Account record (raises on error).
account = client.Account.load({ "id" => "account_id" })
```

#### Example: List

```ruby
# list returns an Array of Account records (raises on error).
accounts = client.Account.list
```

#### Example: Create

```ruby
account = client.Account.create({
})
```


### Batch

Create an instance: `batch = client.Batch`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approve_url` | `String` |  |
| `id` | `String` |  |
| `payout` | `Array` |  |

#### Example: Load

```ruby
# load returns the bare Batch record (raises on error).
batch = client.Batch.load({ "id" => "batch_id" })
```

#### Example: Create

```ruby
batch = client.Batch.create({
})
```


### Beneficiary

Create an instance: `beneficiary = client.Beneficiary`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_callback_url` | `String` |  |
| `authentication_method` | `Array` |  |
| `authorisation` | `Array` |  |
| `authorisers_completed_count` | `Integer` |  |
| `authorisers_required_count` | `Integer` |  |
| `beneficiary` | `Array` |  |
| `beneficiary_event` | `Array` |  |
| `can_authorise` | `Boolean` |  |
| `can_update` | `Boolean` |  |
| `created_by` | `Hash` |  |
| `created_by_email_address` | `String` |  |
| `currency` | `String` |  |
| `destination` | `Hash` |  |
| `failed_beneficiary` | `Hash` |  |
| `has_current_user_authorised` | `Boolean` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `is_enabled` | `Boolean` |  |
| `last_authorised` | `String` |  |
| `last_updated` | `String` |  |
| `merchant_id` | `String` |  |
| `name` | `String` |  |
| `nonce` | `String` |  |
| `source_account` | `Array` |  |
| `source_account_i_d` | `Array` |  |
| `their_reference` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Beneficiary record (raises on error).
beneficiary = client.Beneficiary.load({ "id" => "beneficiary_id" })
```

#### Example: List

```ruby
# list returns an Array of Beneficiary records (raises on error).
beneficiarys = client.Beneficiary.list
```

#### Example: Create

```ruby
beneficiary = client.Beneficiary.create({
})
```


### BeneficiaryGroup

Create an instance: `beneficiary_group = client.BeneficiaryGroup`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `group_member` | `Array` |  |
| `group_name` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `last_updated` | `String` |  |
| `merchant_id` | `String` |  |

#### Example: List

```ruby
# list returns an Array of BeneficiaryGroup records (raises on error).
beneficiary_groups = client.BeneficiaryGroup.list
```


### Card

Create an instance: `card = client.Card`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorized_amount` | `String` |  |
| `currency_code` | `String` |  |
| `is_payer_authentication_required` | `Boolean` |  |
| `is_soft_decline` | `Boolean` |  |
| `payer_authentication_access_token` | `String` |  |
| `payer_authentication_merchant_data` | `String` |  |
| `payer_authentication_url` | `String` |  |
| `payer_authentication_window_height` | `Integer` |  |
| `payer_authentication_window_width` | `Integer` |  |
| `payment_request_callback_url` | `String` |  |
| `payment_request_id` | `String` |  |
| `request_id` | `String` |  |
| `response_code` | `String` |  |
| `response_type` | `String` |  |
| `status` | `String` |  |
| `three_ds_redirect_url` | `String` |  |
| `transaction_id` | `String` |  |

#### Example: Create

```ruby
card = client.Card.create({
  "paymentrequest_id" => "example_paymentrequest_id", # String
})
```


### CardCustomerToken

Create an instance: `card_customer_token = client.CardCustomerToken`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `card_type` | `String` |  |
| `customer_email_address` | `String` |  |
| `expiry_month` | `String` |  |
| `expiry_year` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `last_four_digit` | `String` |  |
| `last_updated` | `String` |  |
| `masked_card_number` | `String` |  |
| `merchant_id` | `String` |  |
| `payment_request_id` | `String` |  |

#### Example: Load

```ruby
# load returns the bare CardCustomerToken record (raises on error).
card_customer_token = client.CardCustomerToken.load({ "customer_email_address" => "customer_email_address" })
```

#### Example: List

```ruby
# list returns an Array of CardCustomerToken records (raises on error).
card_customer_tokens = client.CardCustomerToken.list
```


### CardPayment

Create an instance: `card_payment = client.CardPayment`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorized_amount` | `String` |  |
| `currency_code` | `String` |  |
| `is_payer_authentication_required` | `Boolean` |  |
| `is_soft_decline` | `Boolean` |  |
| `payer_authentication_access_token` | `String` |  |
| `payer_authentication_merchant_data` | `String` |  |
| `payer_authentication_url` | `String` |  |
| `payer_authentication_window_height` | `Integer` |  |
| `payer_authentication_window_width` | `Integer` |  |
| `payment_request_callback_url` | `String` |  |
| `payment_request_id` | `String` |  |
| `request_id` | `String` |  |
| `response_code` | `String` |  |
| `response_type` | `String` |  |
| `status` | `String` |  |
| `three_ds_redirect_url` | `String` |  |
| `transaction_id` | `String` |  |

#### Example: Create

```ruby
card_payment = client.CardPayment.create({
  "paymentrequest_id" => "example_paymentrequest_id", # String
})
```


### CardPublicKey

Create an instance: `card_public_key = client.CardPublicKey`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `jwt` | `String` |  |

#### Example: Load

```ruby
# load returns the bare CardPublicKey record (raises on error).
card_public_key = client.CardPublicKey.load({ "paymentrequest_id" => "paymentrequest_id" })
```


### Consent

Create an instance: `consent = client.Consent`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_url` | `String` |  |
| `callback_url` | `String` |  |
| `consent_id` | `String` |  |
| `email_address` | `String` |  |
| `expiry_date` | `String` |  |
| `failure_callback_url` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `institution_id` | `String` |  |
| `is_connected_account` | `Boolean` |  |
| `is_enabled` | `Boolean` |  |
| `merchant_id` | `String` |  |
| `provider` | `String` |  |
| `success_web_hook_url` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Consent record (raises on error).
consent = client.Consent.load({ "id" => "consent_id" })
```

#### Example: List

```ruby
# list returns an Array of Consent records (raises on error).
consents = client.Consent.list
```

#### Example: Create

```ruby
consent = client.Consent.create({
})
```


### Currency

Create an instance: `currency = client.Currency`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `code` | `String` |  |
| `decimal` | `Integer` |  |
| `is_fiat` | `Boolean` |  |
| `iso4217_alpha_code` | `String` |  |
| `iso4217_numeric_code` | `String` |  |
| `symbol` | `String` |  |

#### Example: List

```ruby
# list returns an Array of Currency records (raises on error).
currencys = client.Currency.list
```


### DirectDebitBatchSubmit

Create an instance: `direct_debit_batch_submit = client.DirectDebitBatchSubmit`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_submission` | `Hash` |  |
| `successful_submission` | `Array` |  |

#### Example: Create

```ruby
direct_debit_batch_submit = client.DirectDebitBatchSubmit.create({
})
```


### FxRate

Create an instance: `fx_rate = client.FxRate`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `destination_currency` | `String` |  |
| `exchange_rate` | `Float` |  |
| `expiry_time` | `String` |  |
| `quote_id` | `String` |  |
| `source_currency` | `String` |  |

#### Example: Load

```ruby
# load returns the bare FxRate record (raises on error).
fx_rate = client.FxRate.load({ "destination" => "destination", "source" => "source", "valid_for_minute" => 1 })
```

#### Example: List

```ruby
# list returns an Array of FxRate records (raises on error).
fx_rates = client.FxRate.list
```


### IPayment

Create an instance: `i_payment = client.IPayment`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `payment_request_id` | `String` |  |
| `response_type` | `String` |  |

#### Example: Create

```ruby
i_payment = client.IPayment.create({
})
```


### Mandate

Create an instance: `mandate = client.Mandate`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_number` | `String` |  |
| `address_line1` | `String` |  |
| `address_line2` | `String` |  |
| `approved_at` | `String` |  |
| `city` | `String` |  |
| `country_code` | `String` |  |
| `currency` | `String` |  |
| `customer_account_number` | `String` |  |
| `customer_city` | `String` |  |
| `customer_country_code` | `String` |  |
| `customer_country_name` | `String` |  |
| `customer_email_address` | `String` |  |
| `customer_first_name` | `String` |  |
| `customer_iban` | `String` |  |
| `customer_last_name` | `String` |  |
| `customer_sort_code` | `String` |  |
| `email_address` | `String` |  |
| `first_name` | `String` |  |
| `iban` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `is_recurring` | `Boolean` |  |
| `last_name` | `String` |  |
| `last_updated` | `String` |  |
| `merchant_id` | `String` |  |
| `postal_code` | `String` |  |
| `reference` | `String` |  |
| `sort_code` | `String` |  |
| `status` | `String` |  |
| `supplier_bank_account_id` | `String` |  |
| `supplier_customer_id` | `String` |  |
| `supplier_mandate_id` | `String` |  |
| `supplier_name` | `String` |  |
| `supplier_status` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Mandate record (raises on error).
mandate = client.Mandate.load({ "id" => "mandate_id" })
```

#### Example: Create

```ruby
mandate = client.Mandate.create({
  "address_line1" => "example_address_line1", # String
  "city" => "example_city", # String
  "country_code" => "example_country_code", # String
  "email_address" => "example_email_address", # String
  "first_name" => "example_first_name", # String
  "last_name" => "example_last_name", # String
  "postal_code" => "example_postal_code", # String
})
```


### Merchant

Create an instance: `merchant = client.Merchant`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_currency` | `Array` |  |
| `can_have_trust_account` | `Boolean` |  |
| `card_payment_processor` | `String` |  |
| `company_id` | `String` |  |
| `display_qr_on_hosted_pay` | `Boolean` |  |
| `hosted_pay_version` | `Integer` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `is_blocked` | `Boolean` |  |
| `is_exited` | `Boolean` |  |
| `is_suspended` | `Boolean` |  |
| `jurisdiction` | `String` |  |
| `logo_url_png` | `String` |  |
| `logo_url_svg` | `String` |  |
| `merchant_category_code` | `String` |  |
| `name` | `String` |  |
| `note` | `String` |  |
| `parent_merchant` | `Hash` |  |
| `payment_account` | `Array` |  |
| `payment_account_limit` | `Integer` |  |
| `reason` | `String` |  |
| `short_name` | `String` |  |
| `supported_payment_methods_list` | `Array` |  |
| `suspension_reason` | `String` |  |
| `tag` | `Array` |  |
| `time_zone_id` | `String` |  |
| `trading_name` | `String` |  |
| `web_hook_limit` | `Integer` |  |
| `your_role_name` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Merchant record (raises on error).
merchant = client.Merchant.load({ "id" => "merchant_id" })
```

#### Example: List

```ruby
# list returns an Array of Merchant records (raises on error).
merchants = client.Merchant.list
```


### MerchantAuthorisationSetting

Create an instance: `merchant_authorisation_setting = client.MerchantAuthorisationSetting`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount_lower` | `Float` |  |
| `amount_upper` | `Float` |  |
| `authorisation_type` | `String` |  |
| `beneficiaries_only` | `Boolean` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `last_editor_cant_authorise` | `Boolean` |  |
| `last_updated` | `String` |  |
| `merchant_id` | `String` |  |
| `number_of_authoriser` | `Integer` |  |
| `role_setting` | `Array` |  |

#### Example: List

```ruby
# list returns an Array of MerchantAuthorisationSetting records (raises on error).
merchant_authorisation_settings = client.MerchantAuthorisationSetting.list
```


### MerchantDirectDebitMandatePage

Create an instance: `merchant_direct_debit_mandate_page = client.MerchantDirectDebitMandatePage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approved_at` | `String` |  |
| `currency` | `String` |  |
| `customer_account_number` | `String` |  |
| `customer_city` | `String` |  |
| `customer_country_code` | `String` |  |
| `customer_country_name` | `String` |  |
| `customer_email_address` | `String` |  |
| `customer_first_name` | `String` |  |
| `customer_iban` | `String` |  |
| `customer_last_name` | `String` |  |
| `customer_sort_code` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `is_recurring` | `Boolean` |  |
| `last_updated` | `String` |  |
| `merchant_id` | `String` |  |
| `reference` | `String` |  |
| `status` | `String` |  |
| `supplier_bank_account_id` | `String` |  |
| `supplier_customer_id` | `String` |  |
| `supplier_mandate_id` | `String` |  |
| `supplier_name` | `String` |  |
| `supplier_status` | `String` |  |

#### Example: List

```ruby
# list returns an Array of MerchantDirectDebitMandatePage records (raises on error).
merchant_direct_debit_mandate_pages = client.MerchantDirectDebitMandatePage.list
```


### MerchantPayByBankSetting

Create an instance: `merchant_pay_by_bank_setting = client.MerchantPayByBankSetting`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bank_country_code` | `Array` |  |
| `bank_id` | `String` |  |
| `bank_name` | `String` |  |
| `business_institution_id` | `String` |  |
| `currency` | `String` |  |
| `logo` | `String` |  |
| `message` | `String` |  |
| `message_image_url` | `String` |  |
| `order` | `Integer` |  |
| `personal_institution_id` | `String` |  |
| `processor` | `String` |  |
| `warning_heading` | `String` |  |
| `warning_message` | `String` |  |

#### Example: List

```ruby
# list returns an Array of MerchantPayByBankSetting records (raises on error).
merchant_pay_by_bank_settings = client.MerchantPayByBankSetting.list
```


### MerchantPaymentRequestTemplate

Create an instance: `merchant_payment_request_template = client.MerchantPaymentRequestTemplate`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `description` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `last_updated` | `String` |  |
| `merchant_id` | `String` |  |
| `name` | `String` |  |
| `template` | `Hash` |  |

#### Example: Load

```ruby
# load returns the bare MerchantPaymentRequestTemplate record (raises on error).
merchant_payment_request_template = client.MerchantPaymentRequestTemplate.load({ "id" => "merchant_payment_request_template_id", "paymentrequest_id" => "paymentrequest_id" })
```

#### Example: List

```ruby
# list returns an Array of MerchantPaymentRequestTemplate records (raises on error).
merchant_payment_request_templates = client.MerchantPaymentRequestTemplate.list
```


### MerchantToken

Create an instance: `merchant_token = client.MerchantToken`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authentication_method` | `Array` |  |
| `authorisation` | `Array` |  |
| `authorisers_completed_count` | `Integer` |  |
| `authorisers_required_count` | `Integer` |  |
| `can_authorise` | `Boolean` |  |
| `description` | `String` |  |
| `expires_at` | `String` |  |
| `has_current_user_authorised` | `Boolean` |  |
| `hmac_algorithm` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `ip_address_whitelist` | `String` |  |
| `is_archived` | `Boolean` |  |
| `is_enabled` | `Boolean` |  |
| `last_authorised` | `String` |  |
| `last_updated` | `String` |  |
| `merchant_id` | `String` |  |
| `nonce` | `String` |  |
| `permission_type` | `Array` |  |
| `request_signature_version` | `Integer` |  |
| `shared_secret_algorithm` | `String` |  |
| `shared_secret_base64` | `String` |  |
| `token` | `String` |  |

#### Example: Load

```ruby
# load returns the bare MerchantToken record (raises on error).
merchant_token = client.MerchantToken.load({ "id" => "merchant_token_id" })
```

#### Example: List

```ruby
# list returns an Array of MerchantToken records (raises on error).
merchant_tokens = client.MerchantToken.list
```

#### Example: Create

```ruby
merchant_token = client.MerchantToken.create({
  "nonce" => "example_nonce", # String
})
```


### Metadata

Create an instance: `metadata = client.Metadata`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```ruby
# load returns the bare Metadata record (raises on error).
metadata = client.Metadata.load()
```


### NoFrixionVersion

Create an instance: `no_frixion_version = client.NoFrixionVersion`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `build_version` | `Integer` |  |
| `major_version` | `Integer` |  |
| `minor_version` | `Integer` |  |
| `release_name` | `String` |  |

#### Example: Load

```ruby
# load returns the bare NoFrixionVersion record (raises on error).
no_frixion_version = client.NoFrixionVersion.load()
```


### OpenBanking

Create an instance: `open_banking = client.OpenBanking`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```ruby
open_banking = client.OpenBanking.create({
  "account_id" => "example_account_id", # String
})
```


### Payeeverification

Create an instance: `payeeverification = client.Payeeverification`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `String` |  |
| `account_number` | `String` |  |
| `iban` | `String` |  |
| `payee_verified_account_name` | `String` |  |
| `result` | `String` |  |
| `secondary_identification` | `String` |  |
| `sort_code` | `String` |  |

#### Example: Create

```ruby
payeeverification = client.Payeeverification.create({
  "account_name" => "example_account_name", # String
  "iban" => "example_iban", # String
})
```


### Payment

Create an instance: `payment = client.Payment`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `Array` |  |
| `amount` | `Float` |  |
| `amount_pending` | `Float` |  |
| `amount_received` | `Float` |  |
| `amount_refunded` | `Float` |  |
| `auto_send_receipt` | `Boolean` |  |
| `base_origin_url` | `String` |  |
| `callback_url` | `String` |  |
| `card_authorize_only` | `Boolean` |  |
| `card_create_token` | `Boolean` |  |
| `card_create_token_mode` | `String` |  |
| `card_ignore_cvn` | `Boolean` |  |
| `card_no_payer_authentication` | `Boolean` |  |
| `card_processor_merchant_id` | `String` |  |
| `card_stripe_payment_intent_id` | `String` |  |
| `card_stripe_payment_intent_secret` | `String` |  |
| `card_transmit_raw_detail` | `Boolean` |  |
| `created_by_user` | `Hash` |  |
| `currency` | `String` |  |
| `custom_field` | `Array` |  |
| `customer_email_address` | `String` |  |
| `customer_id` | `String` |  |
| `customer_name` | `String` |  |
| `description` | `String` |  |
| `destination_account` | `Hash` |  |
| `direct_debit_payment` | `Hash` |  |
| `due_date` | `String` |  |
| `event` | `Array` |  |
| `failure_callback_url` | `String` |  |
| `field_display_setting` | `Array` |  |
| `formatted_amount` | `String` |  |
| `hosted_pay_checkout_url` | `String` |  |
| `id` | `String` |  |
| `ignore_address_verification` | `Boolean` |  |
| `inserted` | `String` |  |
| `inserted_sortable` | `String` |  |
| `is_archived` | `Boolean` |  |
| `jwk` | `String` |  |
| `last_updated` | `String` |  |
| `lightning_invoice` | `String` |  |
| `lightning_invoice_expires_at` | `String` |  |
| `merchant_direct_debit_mandate_id` | `String` |  |
| `merchant_id` | `String` |  |
| `merchant_token_description` | `String` |  |
| `notification_email_address` | `String` |  |
| `notification_role_i_d` | `Array` |  |
| `order_id` | `String` |  |
| `partial_payment_method` | `String` |  |
| `partial_payment_step` | `String` |  |
| `payment_attempt` | `Array` |  |
| `payment_method` | `Array` |  |
| `payment_processor` | `String` |  |
| `payrun_id` | `String` |  |
| `pisp_account_id` | `String` |  |
| `priority_bank_id` | `String` |  |
| `result` | `Hash` |  |
| `sandbox_settle_delay_in_second` | `Integer` |  |
| `shipping_address` | `Hash` |  |
| `shipping_address_city` | `String` |  |
| `shipping_address_country_code` | `String` |  |
| `shipping_address_county` | `String` |  |
| `shipping_address_line1` | `String` |  |
| `shipping_address_line2` | `String` |  |
| `shipping_address_post_code` | `String` |  |
| `shipping_email` | `String` |  |
| `shipping_first_name` | `String` |  |
| `shipping_last_name` | `String` |  |
| `shipping_phone` | `String` |  |
| `status` | `String` |  |
| `success_web_hook_url` | `String` |  |
| `tag` | `Array` |  |
| `tag_id` | `Array` |  |
| `title` | `String` |  |
| `tokenised_card` | `Array` |  |
| `transaction` | `Array` |  |
| `use_hosted_payment_page` | `Boolean` |  |

#### Example: Load

```ruby
# load returns the bare Payment record (raises on error).
payment = client.Payment.load({ "id" => "payment_id" })
```

#### Example: Create

```ruby
payment = client.Payment.create({
  "created_by_user" => {}, # Hash
})
```


### PaymentAccount

Create an instance: `payment_account = client.PaymentAccount`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `String` |  |
| `account_supplier_name` | `String` |  |
| `available_balance` | `Float` |  |
| `available_balance_minor_unit` | `Integer` |  |
| `balance` | `Float` |  |
| `balance_minor_unit` | `Integer` |  |
| `bank_name` | `String` |  |
| `consent_id` | `String` |  |
| `created_by` | `Hash` |  |
| `created_by_display_name` | `String` |  |
| `currency` | `String` |  |
| `default_payment_rail` | `String` |  |
| `display_name` | `String` |  |
| `expiry_date` | `String` |  |
| `external_account_icon` | `String` |  |
| `id` | `String` |  |
| `identifier` | `Hash` |  |
| `inserted` | `String` |  |
| `is_archived` | `Boolean` |  |
| `is_connected_account` | `Boolean` |  |
| `is_default` | `Boolean` |  |
| `is_trust_account` | `Boolean` |  |
| `is_virtual` | `Boolean` |  |
| `last_transaction` | `Hash` |  |
| `last_updated` | `String` |  |
| `merchant_id` | `String` |  |
| `merchant_name` | `String` |  |
| `physical_account_id` | `String` |  |
| `rule` | `Array` |  |
| `submitted_payouts_balance` | `Float` |  |
| `submitted_payouts_balance_minor_unit` | `Integer` |  |
| `summary` | `String` |  |
| `supplier_sepa_instant_status` | `String` |  |
| `xero_bank_feed_connection_status` | `String` |  |
| `xero_bank_feed_last_synced_at` | `String` |  |
| `xero_bank_feed_sync_last_failed_at` | `String` |  |
| `xero_bank_feed_sync_last_failure_reason` | `String` |  |
| `xero_bank_feed_sync_status` | `String` |  |
| `xero_unsynchronised_transactions_count` | `Integer` |  |

#### Example: List

```ruby
# list returns an Array of PaymentAccount records (raises on error).
payment_accounts = client.PaymentAccount.list
```


### PaymentAccountMinimal

Create an instance: `payment_account_minimal = client.PaymentAccountMinimal`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `String` |  |
| `available_balance` | `Float` |  |
| `balance` | `Float` |  |
| `balance_minor_unit` | `Integer` |  |
| `currency` | `String` |  |
| `id` | `String` |  |
| `identifier` | `Hash` |  |
| `is_archived` | `Boolean` |  |
| `is_connected_account` | `Boolean` |  |
| `merchant_id` | `String` |  |
| `submitted_payouts_balance` | `Float` |  |

#### Example: List

```ruby
# list returns an Array of PaymentAccountMinimal records (raises on error).
payment_account_minimals = client.PaymentAccountMinimal.list
```


### PaymentInitiation

Create an instance: `payment_initiation = client.PaymentInitiation`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `payment_initiation_id` | `String` |  |
| `payment_request_callback_url` | `String` |  |
| `payment_request_id` | `String` |  |
| `redirect_url` | `String` |  |
| `response_type` | `String` |  |
| `specific_error_message` | `String` |  |

#### Example: Create

```ruby
payment_initiation = client.PaymentInitiation.create({
  "paymentrequest_id" => "example_paymentrequest_id", # String
})
```


### PaymentRequest

Create an instance: `payment_request = client.PaymentRequest`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `Array` |  |
| `amount` | `Float` |  |
| `amount_pending` | `Float` |  |
| `amount_received` | `Float` |  |
| `amount_refunded` | `Float` |  |
| `auto_send_receipt` | `Boolean` |  |
| `base_origin_url` | `String` |  |
| `callback_url` | `String` |  |
| `card_authorize_only` | `Boolean` |  |
| `card_create_token` | `Boolean` |  |
| `card_create_token_mode` | `String` |  |
| `card_ignore_cvn` | `Boolean` |  |
| `card_processor_merchant_id` | `String` |  |
| `card_stripe_payment_intent_id` | `String` |  |
| `card_stripe_payment_intent_secret` | `String` |  |
| `created_by_user` | `Hash` |  |
| `currency` | `String` |  |
| `custom_field` | `Array` |  |
| `customer_email_address` | `String` |  |
| `customer_id` | `String` |  |
| `customer_name` | `String` |  |
| `description` | `String` |  |
| `destination_account` | `Hash` |  |
| `direct_debit_payment` | `Hash` |  |
| `do_simulate_settlement_failure` | `Boolean` |  |
| `due_date` | `String` |  |
| `error_description` | `String` |  |
| `event` | `Array` |  |
| `failed_payment_request` | `Hash` |  |
| `failure_callback_url` | `String` |  |
| `field_display_setting` | `Array` |  |
| `formatted_amount` | `String` |  |
| `hosted_pay_checkout_url` | `String` |  |
| `id` | `String` |  |
| `ignore_address_verification` | `Boolean` |  |
| `inserted` | `String` |  |
| `inserted_sortable` | `String` |  |
| `institution` | `String` |  |
| `is_archived` | `Boolean` |  |
| `jwk` | `String` |  |
| `last_updated` | `String` |  |
| `lightning_invoice` | `String` |  |
| `lightning_invoice_expires_at` | `String` |  |
| `merchant_direct_debit_mandate_id` | `String` |  |
| `merchant_id` | `String` |  |
| `merchant_token_description` | `String` |  |
| `notification_email_address` | `String` |  |
| `notification_role_i_d` | `Array` |  |
| `order_id` | `String` |  |
| `partial_payment_method` | `String` |  |
| `partial_payment_step` | `String` |  |
| `payment_attempt` | `Array` |  |
| `payment_initiation_id` | `String` |  |
| `payment_method` | `Array` |  |
| `payment_processor` | `String` |  |
| `payment_request` | `Array` |  |
| `payrun_id` | `String` |  |
| `pisp_account_id` | `String` |  |
| `priority_bank_id` | `String` |  |
| `result` | `Hash` |  |
| `sandbox_settle_delay_in_second` | `Integer` |  |
| `shipping_address` | `Hash` |  |
| `status` | `String` |  |
| `success_web_hook_url` | `String` |  |
| `tag` | `Array` |  |
| `title` | `String` |  |
| `tokenised_card` | `Array` |  |
| `transaction` | `Array` |  |
| `use_hosted_payment_page` | `Boolean` |  |

#### Example: Load

```ruby
# load returns the bare PaymentRequest record (raises on error).
payment_request = client.PaymentRequest.load()
```

#### Example: List

```ruby
# list returns an Array of PaymentRequest records (raises on error).
payment_requests = client.PaymentRequest.list
```

#### Example: Create

```ruby
payment_request = client.PaymentRequest.create({
})
```


### PaymentRequestEvent

Create an instance: `payment_request_event = client.PaymentRequestEvent`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `Float` |  |
| `apple_pay_transaction_id` | `String` |  |
| `card_authorization_response_id` | `String` |  |
| `card_expiry_month` | `Integer` |  |
| `card_expiry_year` | `Integer` |  |
| `card_issuer` | `String` |  |
| `card_issuer_country` | `String` |  |
| `card_last_four_digit` | `String` |  |
| `card_request_id` | `String` |  |
| `card_scheme` | `String` |  |
| `card_token_customer_id` | `String` |  |
| `card_transaction_id` | `String` |  |
| `currency` | `String` |  |
| `direct_debit_payment_id` | `String` |  |
| `direct_debit_payment_reference` | `String` |  |
| `drirect_debit_mandate_id` | `String` |  |
| `error_message` | `String` |  |
| `error_reason` | `String` |  |
| `event_type` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `lightning_invoice` | `String` |  |
| `lightning_r_hash` | `String` |  |
| `origin_url` | `String` |  |
| `payment_method_type` | `String` |  |
| `payment_processor_name` | `String` |  |
| `payment_request_id` | `String` |  |
| `pisp_bank_status` | `String` |  |
| `pisp_payment_initiation_id` | `String` |  |
| `pisp_payment_institution_name` | `String` |  |
| `pisp_payment_service_provider_id` | `String` |  |
| `pisp_redirect_url` | `String` |  |
| `reconciled_transaction_id` | `String` |  |
| `refund_payout_id` | `String` |  |
| `status` | `String` |  |
| `wallet_name` | `String` |  |

#### Example: List

```ruby
# list returns an Array of PaymentRequestEvent records (raises on error).
payment_request_events = client.PaymentRequestEvent.list
```


### PaymentRequestMetric

Create an instance: `payment_request_metric = client.PaymentRequestMetric`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `all` | `Integer` |  |
| `authorized` | `Integer` |  |
| `paid` | `Integer` |  |
| `partially_paid` | `Integer` |  |
| `total_amounts_by_currency` | `Hash` |  |
| `unpaid` | `Integer` |  |

#### Example: Load

```ruby
# load returns the bare PaymentRequestMetric record (raises on error).
payment_request_metric = client.PaymentRequestMetric.load()
```


### PaymentRequestMinimal

Create an instance: `payment_request_minimal = client.PaymentRequestMinimal`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `Float` |  |
| `amount_pending` | `Float` |  |
| `amount_received` | `Float` |  |
| `amount_refunded` | `Float` |  |
| `callback_url` | `String` |  |
| `card_stripe_payment_intent_secret` | `String` |  |
| `country_code` | `String` |  |
| `currency` | `String` |  |
| `custom_fields_to_display` | `Array` |  |
| `description` | `String` |  |
| `due_date` | `String` |  |
| `field_display_setting` | `Array` |  |
| `google_pay_merchant_id` | `String` |  |
| `id` | `String` |  |
| `jwk` | `String` |  |
| `merchant_id` | `String` |  |
| `merchant_logo_url_png` | `String` |  |
| `merchant_logo_url_svg` | `String` |  |
| `merchant_name` | `String` |  |
| `merchant_short_name` | `String` |  |
| `partial_payment_method` | `String` |  |
| `payment_attempt` | `Array` |  |
| `payment_methods_list` | `Array` |  |
| `payment_processor` | `String` |  |
| `payment_processor_key` | `String` |  |
| `pisp_error` | `String` |  |
| `priority_bank_id` | `String` |  |
| `status` | `String` |  |
| `stripe_account_id` | `String` |  |
| `title` | `String` |  |

#### Example: List

```ruby
# list returns an Array of PaymentRequestMinimal records (raises on error).
payment_request_minimals = client.PaymentRequestMinimal.list
```


### PaymentRequestResult

Create an instance: `payment_request_result = client.PaymentRequestResult`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `Float` |  |
| `amount_pending` | `Float` |  |
| `amount_received` | `Float` |  |
| `amount_refunded` | `Float` |  |
| `currency` | `String` |  |
| `customer_id` | `String` |  |
| `payment` | `Array` |  |
| `payment_request_id` | `String` |  |
| `pisp_authorization` | `Array` |  |
| `requested_amount` | `Float` |  |
| `result` | `String` |  |

#### Example: List

```ruby
# list returns an Array of PaymentRequestResult records (raises on error).
payment_request_results = client.PaymentRequestResult.list
```


### Payout

Create an instance: `payout = client.Payout`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `String` |  |
| `allow_incomplete` | `Boolean` |  |
| `amount` | `Float` |  |
| `amount_minor_unit` | `Integer` |  |
| `approve_payout_url` | `String` |  |
| `approver_id` | `String` |  |
| `authentication_method` | `Array` |  |
| `authorisation` | `Array` |  |
| `authorisers_completed_count` | `Integer` |  |
| `authorisers_required_count` | `Integer` |  |
| `batch_payout_id` | `String` |  |
| `beneficiary` | `Hash` |  |
| `beneficiary_id` | `String` |  |
| `can_authorise` | `Boolean` |  |
| `can_process` | `Boolean` |  |
| `can_update` | `Boolean` |  |
| `charge_bearer` | `String` |  |
| `created_by` | `String` |  |
| `created_by_email_address` | `String` |  |
| `currency` | `String` |  |
| `current_user_id` | `String` |  |
| `description` | `String` |  |
| `destination` | `Hash` |  |
| `document` | `Array` |  |
| `event` | `Array` |  |
| `failed_payout` | `Hash` |  |
| `formatted_amount` | `String` |  |
| `formatted_fx_destination_amount` | `String` |  |
| `formatted_schedule` | `String` |  |
| `formatted_schedule_day_only` | `String` |  |
| `formatted_source_account_available_balance` | `String` |  |
| `fx_destination_amount` | `Float` |  |
| `fx_destination_amount_minor_unit` | `Integer` |  |
| `fx_destination_currency` | `String` |  |
| `fx_quote_expires_at` | `String` |  |
| `fx_quote_id` | `String` |  |
| `fx_rate` | `Float` |  |
| `fx_use_destination_amount` | `Boolean` |  |
| `has_current_user_authorised` | `Boolean` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `invoice_id` | `String` |  |
| `is_archived` | `Boolean` |  |
| `is_failed` | `Boolean` |  |
| `is_settled` | `Boolean` |  |
| `is_submitted` | `Boolean` |  |
| `last_updated` | `String` |  |
| `merchant_id` | `String` |  |
| `merchant_token_description` | `String` |  |
| `nonce` | `String` |  |
| `payment_processor` | `String` |  |
| `payment_rail` | `String` |  |
| `payout` | `Array` |  |
| `payrun_id` | `String` |  |
| `payrun_name` | `String` |  |
| `reason` | `String` |  |
| `rule` | `Hash` |  |
| `schedule_date` | `String` |  |
| `scheduled` | `Boolean` |  |
| `source_account_available_balance` | `Float` |  |
| `source_account_available_balance_minor_unit` | `Integer` |  |
| `source_account_bic` | `String` |  |
| `source_account_currency` | `String` |  |
| `source_account_iban` | `String` |  |
| `source_account_identifier` | `Hash` |  |
| `source_account_name` | `String` |  |
| `source_account_number` | `String` |  |
| `source_account_sortcode` | `String` |  |
| `status` | `String` |  |
| `tag` | `Array` |  |
| `tag_id` | `Array` |  |
| `their_reference` | `String` |  |
| `topup_payrun_id` | `String` |  |
| `transacted_amount` | `Float` |  |
| `transacted_fx_amount` | `Float` |  |
| `transacted_fx_rate` | `Float` |  |
| `type` | `String` |  |
| `user_id` | `String` |  |
| `your_reference` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Payout record (raises on error).
payout = client.Payout.load({ "id" => "payout_id" })
```

#### Example: List

```ruby
# list returns an Array of Payout records (raises on error).
payouts = client.Payout.list
```

#### Example: Create

```ruby
payout = client.Payout.create({
})
```


### PayoutKeysetPage

Create an instance: `payout_keyset_page = client.PayoutKeysetPage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `String` |  |
| `amount` | `Float` |  |
| `amount_minor_unit` | `Integer` |  |
| `approve_payout_url` | `String` |  |
| `approver_id` | `String` |  |
| `authentication_method` | `Array` |  |
| `authorisation` | `Array` |  |
| `authorisers_completed_count` | `Integer` |  |
| `authorisers_required_count` | `Integer` |  |
| `batch_payout_id` | `String` |  |
| `beneficiary` | `Hash` |  |
| `can_authorise` | `Boolean` |  |
| `can_process` | `Boolean` |  |
| `can_update` | `Boolean` |  |
| `charge_bearer` | `String` |  |
| `created_by` | `String` |  |
| `created_by_email_address` | `String` |  |
| `currency` | `String` |  |
| `current_user_id` | `String` |  |
| `description` | `String` |  |
| `destination` | `Hash` |  |
| `document` | `Array` |  |
| `event` | `Array` |  |
| `formatted_amount` | `String` |  |
| `formatted_fx_destination_amount` | `String` |  |
| `formatted_schedule` | `String` |  |
| `formatted_schedule_day_only` | `String` |  |
| `formatted_source_account_available_balance` | `String` |  |
| `fx_destination_amount` | `Float` |  |
| `fx_destination_amount_minor_unit` | `Integer` |  |
| `fx_destination_currency` | `String` |  |
| `fx_quote_expires_at` | `String` |  |
| `fx_quote_id` | `String` |  |
| `fx_rate` | `Float` |  |
| `fx_use_destination_amount` | `Boolean` |  |
| `has_current_user_authorised` | `Boolean` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `invoice_id` | `String` |  |
| `is_archived` | `Boolean` |  |
| `is_failed` | `Boolean` |  |
| `is_settled` | `Boolean` |  |
| `is_submitted` | `Boolean` |  |
| `last_updated` | `String` |  |
| `merchant_id` | `String` |  |
| `merchant_token_description` | `String` |  |
| `nonce` | `String` |  |
| `payment_processor` | `String` |  |
| `payment_rail` | `String` |  |
| `payrun_id` | `String` |  |
| `payrun_name` | `String` |  |
| `rule` | `Hash` |  |
| `schedule_date` | `String` |  |
| `scheduled` | `Boolean` |  |
| `source_account_available_balance` | `Float` |  |
| `source_account_available_balance_minor_unit` | `Integer` |  |
| `source_account_bic` | `String` |  |
| `source_account_currency` | `String` |  |
| `source_account_iban` | `String` |  |
| `source_account_identifier` | `Hash` |  |
| `source_account_name` | `String` |  |
| `source_account_number` | `String` |  |
| `source_account_sortcode` | `String` |  |
| `status` | `String` |  |
| `tag` | `Array` |  |
| `their_reference` | `String` |  |
| `topup_payrun_id` | `String` |  |
| `transacted_amount` | `Float` |  |
| `transacted_fx_amount` | `Float` |  |
| `transacted_fx_rate` | `Float` |  |
| `type` | `String` |  |
| `user_id` | `String` |  |
| `your_reference` | `String` |  |

#### Example: List

```ruby
# list returns an Array of PayoutKeysetPage records (raises on error).
payout_keyset_pages = client.PayoutKeysetPage.list
```


### PayoutMetric

Create an instance: `payout_metric = client.PayoutMetric`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `all` | `Float` |  |
| `failed` | `Float` |  |
| `in_progress` | `Float` |  |
| `paid` | `Float` |  |
| `pending_approval` | `Float` |  |
| `scheduled` | `Float` |  |
| `total_amounts_by_currency` | `Hash` |  |

#### Example: Load

```ruby
# load returns the bare PayoutMetric record (raises on error).
payout_metric = client.PayoutMetric.load()
```


### Payrun

Create an instance: `payrun = client.Payrun`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation` | `Array` |  |
| `authorisation_date` | `String` |  |
| `authorisers_completed_count` | `Integer` |  |
| `authorisers_required_count` | `Integer` |  |
| `batch_payout_id` | `String` |  |
| `can_authorise` | `Boolean` |  |
| `can_delete` | `Boolean` |  |
| `can_edit` | `Boolean` |  |
| `event` | `Array` |  |
| `has_current_user_authorised` | `Boolean` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `invoice` | `Array` |  |
| `invoices_minimal` | `Array` |  |
| `is_archived` | `Boolean` |  |
| `last_updated` | `String` |  |
| `last_updated_by` | `Hash` |  |
| `merchant_id` | `String` |  |
| `name` | `String` |  |
| `nonce` | `String` |  |
| `note` | `String` |  |
| `payment` | `Array` |  |
| `payout` | `Array` |  |
| `payouts_count` | `Integer` |  |
| `reason` | `String` |  |
| `schedule_date` | `String` |  |
| `scheduled_date` | `String` |  |
| `source_account` | `Array` |  |
| `status` | `String` |  |
| `total_eur` | `Float` |  |
| `total_gbp` | `Float` |  |
| `total_usd` | `Float` |  |

#### Example: Load

```ruby
# load returns the bare Payrun record (raises on error).
payrun = client.Payrun.load({ "id" => "payrun_id" })
```

#### Example: List

```ruby
# list returns an Array of Payrun records (raises on error).
payruns = client.Payrun.list
```

#### Example: Create

```ruby
payrun = client.Payrun.create({
  "id" => "example_id", # String
})
```


### Report

Create an instance: `report = client.Report`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |


### ReportResult

Create an instance: `report_result = client.ReportResult`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `content` | `String` |  |
| `content_type` | `String` |  |
| `last_completed_at` | `String` |  |
| `merchant_id` | `String` |  |
| `report_name` | `String` |  |
| `report_type` | `String` |  |
| `statement_number` | `Integer` |  |

#### Example: Load

```ruby
# load returns the bare ReportResult record (raises on error).
report_result = client.ReportResult.load({ "id" => 1, "report_id" => "report_id" })
```


### Role

Create an instance: `role = client.Role`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_role` | `Hash` |  |
| `role` | `Array` |  |

#### Example: Create

```ruby
role = client.Role.create({
  "merchant_id" => "example_merchant_id", # String
})
```


### Rule

Create an instance: `rule = client.Rule`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account` | `Hash` |  |
| `account_id` | `String` |  |
| `approve_url` | `String` |  |
| `approver_id` | `String` |  |
| `authentication_method` | `Array` |  |
| `authorisation` | `Array` |  |
| `authorisers_completed_count` | `Integer` |  |
| `authorisers_required_count` | `Integer` |  |
| `can_authorise` | `Boolean` |  |
| `created_by` | `Hash` |  |
| `description` | `String` |  |
| `end_at` | `String` |  |
| `has_current_user_authorised` | `Boolean` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `is_disabled` | `Boolean` |  |
| `last_executed_at` | `String` |  |
| `last_run_at_transaction_date` | `String` |  |
| `last_updated` | `String` |  |
| `merchant_id` | `String` |  |
| `name` | `String` |  |
| `nonce` | `String` |  |
| `on_approved_web_hook_url` | `String` |  |
| `on_execution_error_web_hook_url` | `String` |  |
| `on_execution_success_web_hook_url` | `String` |  |
| `start_at` | `String` |  |
| `status` | `String` |  |
| `sweep_action` | `Hash` |  |
| `time_zone_id` | `String` |  |
| `trigger_cron_expression` | `String` |  |
| `trigger_on_pay_in` | `Boolean` |  |
| `user_id` | `String` |  |
| `web_hook_secret` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Rule record (raises on error).
rule = client.Rule.load({ "id" => "rule_id" })
```

#### Example: List

```ruby
# list returns an Array of Rule records (raises on error).
rules = client.Rule.list
```

#### Example: Create

```ruby
rule = client.Rule.create({
  "created_by" => {}, # Hash
  "nonce" => "example_nonce", # String
})
```


### RuleEvent

Create an instance: `rule_event = client.RuleEvent`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_message` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `is_authorise_to_enable` | `Boolean` |  |
| `message` | `String` |  |
| `raw_response` | `String` |  |
| `rule_event_type` | `String` |  |
| `rule_id` | `String` |  |
| `user` | `Hash` |  |

#### Example: List

```ruby
# list returns an Array of RuleEvent records (raises on error).
rule_events = client.RuleEvent.list
```


### Tag

Create an instance: `tag = client.Tag`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `colour_hex` | `String` |  |
| `description` | `String` |  |
| `id` | `String` |  |
| `merchant_id` | `String` |  |
| `name` | `String` |  |

#### Example: List

```ruby
# list returns an Array of Tag records (raises on error).
tags = client.Tag.list
```

#### Example: Create

```ruby
tag = client.Tag.create({
  "merchant_id" => "example_merchant_id", # String
})
```


### Token

Create an instance: `token = client.Token`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```ruby
token = client.Token.create({
  "id" => "example_id", # String
})
```


### Transaction

Create an instance: `transaction = client.Transaction`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `String` |  |
| `account_name` | `String` |  |
| `account_sequence_number` | `Integer` |  |
| `address_detail` | `Hash` |  |
| `amount` | `Float` |  |
| `amount_minor_unit` | `Integer` |  |
| `balance` | `Float` |  |
| `balance_minor_unit` | `Integer` |  |
| `booking_date_time` | `String` |  |
| `charge_detail` | `Hash` |  |
| `content` | `Array` |  |
| `counterparty` | `Hash` |  |
| `counterparty_summary` | `String` |  |
| `currency` | `String` |  |
| `currency_exchange` | `Hash` |  |
| `date` | `String` |  |
| `description` | `String` |  |
| `enrichment` | `Hash` |  |
| `fx_amount` | `Float` |  |
| `fx_currency` | `String` |  |
| `fx_rate` | `Float` |  |
| `gross_amount` | `Hash` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `iso_bank_transaction_code` | `Hash` |  |
| `merchant` | `Hash` |  |
| `merchant_id` | `String` |  |
| `page_number` | `Integer` |  |
| `page_size` | `Integer` |  |
| `payee_detail` | `Hash` |  |
| `payer_detail` | `Hash` |  |
| `payment_request_custom_field` | `Hash` |  |
| `payment_request_id` | `String` |  |
| `payout_id` | `String` |  |
| `proprietary_bank_transaction_code` | `Hash` |  |
| `raw_reference` | `String` |  |
| `reference` | `String` |  |
| `rule_id` | `String` |  |
| `statement_reference` | `Array` |  |
| `status` | `String` |  |
| `supplementary_data` | `Object` |  |
| `tag` | `Array` |  |
| `their_reference` | `String` |  |
| `total_page` | `Integer` |  |
| `total_size` | `Integer` |  |
| `transaction_amount` | `Hash` |  |
| `transaction_date` | `String` |  |
| `transaction_information` | `Array` |  |
| `transaction_mutability` | `String` |  |
| `type` | `String` |  |
| `value_date_time` | `String` |  |
| `virtual_iban` | `String` |  |
| `your_reference` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Transaction record (raises on error).
transaction = client.Transaction.load({ "id" => "transaction_id" })
```

#### Example: List

```ruby
# list returns an Array of Transaction records (raises on error).
transactions = client.Transaction.list
```

#### Example: Create

```ruby
transaction = client.Transaction.create({
  "id" => "example_id", # String
})
```


### User

Create an instance: `user = client.User`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client_session_timeout` | `Array` |  |
| `email_address` | `String` |  |
| `first_name` | `String` |  |
| `id` | `String` |  |
| `last_name` | `String` |  |
| `passkey_added` | `Boolean` |  |
| `permission` | `Hash` |  |
| `profile` | `String` |  |
| `roles_with_scope` | `Array` |  |
| `two_factor_enabled` | `Boolean` |  |
| `user_invite_id` | `String` |  |

#### Example: List

```ruby
# list returns an Array of User records (raises on error).
users = client.User.list
```


### UserInvite

Create an instance: `user_invite = client.UserInvite`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_status` | `Hash` |  |
| `failed_user_invite` | `Hash` |  |
| `id` | `String` |  |
| `initial_role_id` | `String` |  |
| `invitee_email_address` | `String` |  |
| `invitee_first_name` | `String` |  |
| `invitee_last_name` | `String` |  |
| `inviter_email_address` | `String` |  |
| `inviter_first_name` | `String` |  |
| `inviter_last_name` | `String` |  |
| `is_authorised` | `Boolean` |  |
| `is_invitee_registered` | `Boolean` |  |
| `last_invited` | `String` |  |
| `merchant_id` | `String` |  |
| `merchant_name` | `String` |  |
| `message` | `String` |  |
| `registration_url` | `String` |  |
| `send_invite_email` | `Boolean` |  |
| `status` | `String` |  |
| `user` | `Hash` |  |
| `user_id` | `String` |  |
| `user_invite` | `Array` |  |

#### Example: Load

```ruby
# load returns the bare UserInvite record (raises on error).
user_invite = client.UserInvite.load({ "id" => "user_invite_id" })
```

#### Example: List

```ruby
# list returns an Array of UserInvite records (raises on error).
user_invites = client.UserInvite.list
```

#### Example: Create

```ruby
user_invite = client.UserInvite.create({
})
```


### Virtual

Create an instance: `virtual = client.Virtual`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `String` |  |
| `account_supplier_name` | `String` |  |
| `available_balance` | `Float` |  |
| `available_balance_minor_unit` | `Integer` |  |
| `balance` | `Float` |  |
| `balance_minor_unit` | `Integer` |  |
| `bank_name` | `String` |  |
| `consent_id` | `String` |  |
| `created_by` | `Hash` |  |
| `created_by_display_name` | `String` |  |
| `currency` | `String` |  |
| `default_payment_rail` | `String` |  |
| `display_name` | `String` |  |
| `expiry_date` | `String` |  |
| `external_account_icon` | `String` |  |
| `id` | `String` |  |
| `identifier` | `Hash` |  |
| `inserted` | `String` |  |
| `is_archived` | `Boolean` |  |
| `is_connected_account` | `Boolean` |  |
| `is_default` | `Boolean` |  |
| `is_trust_account` | `Boolean` |  |
| `is_virtual` | `Boolean` |  |
| `last_transaction` | `Hash` |  |
| `last_updated` | `String` |  |
| `merchant_id` | `String` |  |
| `merchant_name` | `String` |  |
| `name` | `String` |  |
| `physical_account_id` | `String` |  |
| `rule` | `Array` |  |
| `submitted_payouts_balance` | `Float` |  |
| `submitted_payouts_balance_minor_unit` | `Integer` |  |
| `summary` | `String` |  |
| `supplier_sepa_instant_status` | `String` |  |
| `xero_bank_feed_connection_status` | `String` |  |
| `xero_bank_feed_last_synced_at` | `String` |  |
| `xero_bank_feed_sync_last_failed_at` | `String` |  |
| `xero_bank_feed_sync_last_failure_reason` | `String` |  |
| `xero_bank_feed_sync_status` | `String` |  |
| `xero_unsynchronised_transactions_count` | `Integer` |  |

#### Example: Create

```ruby
virtual = client.Virtual.create({
  "account_id" => "example_account_id", # String
})
```


### Webhook

Create an instance: `webhook = client.Webhook`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `destination_url` | `String` |  |
| `email_address` | `String` |  |
| `failed_notification_email_address` | `String` |  |
| `id` | `String` |  |
| `is_active` | `Boolean` |  |
| `merchant_id` | `String` |  |
| `notification_method` | `String` |  |
| `resource_type` | `Array` |  |
| `retry` | `Boolean` |  |
| `secret` | `String` |  |
| `version` | `Integer` |  |

#### Example: Load

```ruby
# load returns the bare Webhook record (raises on error).
webhook = client.Webhook.load({ "id" => "webhook_id" })
```

#### Example: List

```ruby
# list returns an Array of Webhook records (raises on error).
webhooks = client.Webhook.list
```

#### Example: Create

```ruby
webhook = client.Webhook.create({
})
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── Nofrixion_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`Nofrixion_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```ruby
account = client.Account
account.list()

# account.data_get now returns the account data from the last list
# account.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
