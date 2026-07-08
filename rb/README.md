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
    puts "#{item["id"]} #{item["account_id"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```

### 3. Load a merchant

Merchant is nested under merchant, so provide the `merchant_id`.

```ruby
begin
  # load returns the bare Merchant record (raises on error).
  merchant = client.Merchant.load({ "merchant_id" => "example_merchant_id" })
  puts merchant
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
| `Beneficiary` | `(data) -> BeneficiaryEntity` | Create a Beneficiary entity instance. |
| `Cancel` | `(data) -> CancelEntity` | Create a Cancel entity instance. |
| `Disable` | `(data) -> DisableEntity` | Create a Disable entity instance. |
| `Enable` | `(data) -> EnableEntity` | Create an Enable entity instance. |
| `Merchant` | `(data) -> MerchantEntity` | Create a Merchant entity instance. |
| `Metadata` | `(data) -> MetadataEntity` | Create a Metadata entity instance. |
| `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage` | `(data) -> NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity` | Create a NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage entity instance. |
| `NoFrixionBizBizModelsPagingPaymentRequestPage` | `(data) -> NoFrixionBizBizModelsPagingPaymentRequestPageEntity` | Create a NoFrixionBizBizModelsPagingPaymentRequestPage entity instance. |
| `NoFrixionBizBizModelsPagingPayoutPage` | `(data) -> NoFrixionBizBizModelsPagingPayoutPageEntity` | Create a NoFrixionBizBizModelsPagingPayoutPage entity instance. |
| `NoFrixionBizBizModelsPagingPayrunPage` | `(data) -> NoFrixionBizBizModelsPagingPayrunPageEntity` | Create a NoFrixionBizBizModelsPagingPayrunPage entity instance. |
| `NoFrixionBizBizModelsPagingRuleEventsPage` | `(data) -> NoFrixionBizBizModelsPagingRuleEventsPageEntity` | Create a NoFrixionBizBizModelsPagingRuleEventsPage entity instance. |
| `NoFrixionBizBizModelsPagingRulesPage` | `(data) -> NoFrixionBizBizModelsPagingRulesPageEntity` | Create a NoFrixionBizBizModelsPagingRulesPage entity instance. |
| `NoFrixionBizBizModelsPaymentsCardPayment` | `(data) -> NoFrixionBizBizModelsPaymentsCardPaymentEntity` | Create a NoFrixionBizBizModelsPaymentsCardPayment entity instance. |
| `NoFrixionBizBizModelsPaymentsCardPublicKey` | `(data) -> NoFrixionBizBizModelsPaymentsCardPublicKeyEntity` | Create a NoFrixionBizBizModelsPaymentsCardPublicKey entity instance. |
| `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries` | `(data) -> NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity` | Create a NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries entity instance. |
| `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment` | `(data) -> NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity` | Create a NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment entity instance. |
| `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate` | `(data) -> NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity` | Create a NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate entity instance. |
| `NoFrixionMoneyMoovApiFeaturesUserInvitesCreate` | `(data) -> NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity` | Create a NoFrixionMoneyMoovApiFeaturesUserInvitesCreate entity instance. |
| `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant` | `(data) -> NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity` | Create a NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant entity instance. |
| `NoFrixionMoneyMoovModelsBatchPayout` | `(data) -> NoFrixionMoneyMoovModelsBatchPayoutEntity` | Create a NoFrixionMoneyMoovModelsBatchPayout entity instance. |
| `NoFrixionMoneyMoovModelsBeneficiaryGroupPage` | `(data) -> NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity` | Create a NoFrixionMoneyMoovModelsBeneficiaryGroupPage entity instance. |
| `NoFrixionMoneyMoovModelsBeneficiaryPage` | `(data) -> NoFrixionMoneyMoovModelsBeneficiaryPageEntity` | Create a NoFrixionMoneyMoovModelsBeneficiaryPage entity instance. |
| `NoFrixionMoneyMoovModelsCardCustomerToken` | `(data) -> NoFrixionMoneyMoovModelsCardCustomerTokenEntity` | Create a NoFrixionMoneyMoovModelsCardCustomerToken entity instance. |
| `NoFrixionMoneyMoovModelsCurrencyCurrencyInfo` | `(data) -> NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity` | Create a NoFrixionMoneyMoovModelsCurrencyCurrencyInfo entity instance. |
| `NoFrixionMoneyMoovModelsDirectDebitBatchSubmit` | `(data) -> NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity` | Create a NoFrixionMoneyMoovModelsDirectDebitBatchSubmit entity instance. |
| `NoFrixionMoneyMoovModelsFxRate` | `(data) -> NoFrixionMoneyMoovModelsFxRateEntity` | Create a NoFrixionMoneyMoovModelsFxRate entity instance. |
| `NoFrixionMoneyMoovModelsIPayment` | `(data) -> NoFrixionMoneyMoovModelsIPaymentEntity` | Create a NoFrixionMoneyMoovModelsIPayment entity instance. |
| `NoFrixionMoneyMoovModelsMandatesMandate` | `(data) -> NoFrixionMoneyMoovModelsMandatesMandateEntity` | Create a NoFrixionMoneyMoovModelsMandatesMandate entity instance. |
| `NoFrixionMoneyMoovModelsMerchant` | `(data) -> NoFrixionMoneyMoovModelsMerchantEntity` | Create a NoFrixionMoneyMoovModelsMerchant entity instance. |
| `NoFrixionMoneyMoovModelsMerchantPage` | `(data) -> NoFrixionMoneyMoovModelsMerchantPageEntity` | Create a NoFrixionMoneyMoovModelsMerchantPage entity instance. |
| `NoFrixionMoneyMoovModelsMerchantPayByBankSetting` | `(data) -> NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity` | Create a NoFrixionMoneyMoovModelsMerchantPayByBankSetting entity instance. |
| `NoFrixionMoneyMoovModelsMerchantToken` | `(data) -> NoFrixionMoneyMoovModelsMerchantTokenEntity` | Create a NoFrixionMoneyMoovModelsMerchantToken entity instance. |
| `NoFrixionMoneyMoovModelsMerchantTokenPage` | `(data) -> NoFrixionMoneyMoovModelsMerchantTokenPageEntity` | Create a NoFrixionMoneyMoovModelsMerchantTokenPage entity instance. |
| `NoFrixionMoneyMoovModelsNoFrixionVersion` | `(data) -> NoFrixionMoneyMoovModelsNoFrixionVersionEntity` | Create a NoFrixionMoneyMoovModelsNoFrixionVersion entity instance. |
| `NoFrixionMoneyMoovModelsOpenBankingAccount` | `(data) -> NoFrixionMoneyMoovModelsOpenBankingAccountEntity` | Create a NoFrixionMoneyMoovModelsOpenBankingAccount entity instance. |
| `NoFrixionMoneyMoovModelsOpenBankingConsent` | `(data) -> NoFrixionMoneyMoovModelsOpenBankingConsentEntity` | Create a NoFrixionMoneyMoovModelsOpenBankingConsent entity instance. |
| `NoFrixionMoneyMoovModelsOpenBankingTransaction` | `(data) -> NoFrixionMoneyMoovModelsOpenBankingTransactionEntity` | Create a NoFrixionMoneyMoovModelsOpenBankingTransaction entity instance. |
| `NoFrixionMoneyMoovModelsPayment` | `(data) -> NoFrixionMoneyMoovModelsPaymentEntity` | Create a NoFrixionMoneyMoovModelsPayment entity instance. |
| `NoFrixionMoneyMoovModelsPaymentAccountMinimalPage` | `(data) -> NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity` | Create a NoFrixionMoneyMoovModelsPaymentAccountMinimalPage entity instance. |
| `NoFrixionMoneyMoovModelsPaymentAccountPage` | `(data) -> NoFrixionMoneyMoovModelsPaymentAccountPageEntity` | Create a NoFrixionMoneyMoovModelsPaymentAccountPage entity instance. |
| `NoFrixionMoneyMoovModelsPaymentInitiation` | `(data) -> NoFrixionMoneyMoovModelsPaymentInitiationEntity` | Create a NoFrixionMoneyMoovModelsPaymentInitiation entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestEvent` | `(data) -> NoFrixionMoneyMoovModelsPaymentRequestEventEntity` | Create a NoFrixionMoneyMoovModelsPaymentRequestEvent entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestMetric` | `(data) -> NoFrixionMoneyMoovModelsPaymentRequestMetricEntity` | Create a NoFrixionMoneyMoovModelsPaymentRequestMetric entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestMinimal` | `(data) -> NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity` | Create a NoFrixionMoneyMoovModelsPaymentRequestMinimal entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestResult` | `(data) -> NoFrixionMoneyMoovModelsPaymentRequestResultEntity` | Create a NoFrixionMoneyMoovModelsPaymentRequestResult entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment` | `(data) -> NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity` | Create a NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2` | `(data) -> NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity` | Create a NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2 entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3` | `(data) -> NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity` | Create a NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3 entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4` | `(data) -> NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity` | Create a NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4 entity instance. |
| `NoFrixionMoneyMoovModelsPayoutKeysetPage` | `(data) -> NoFrixionMoneyMoovModelsPayoutKeysetPageEntity` | Create a NoFrixionMoneyMoovModelsPayoutKeysetPage entity instance. |
| `NoFrixionMoneyMoovModelsPayoutMetric` | `(data) -> NoFrixionMoneyMoovModelsPayoutMetricEntity` | Create a NoFrixionMoneyMoovModelsPayoutMetric entity instance. |
| `NoFrixionMoneyMoovModelsPayoutsPayoutsCreate` | `(data) -> NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity` | Create a NoFrixionMoneyMoovModelsPayoutsPayoutsCreate entity instance. |
| `NoFrixionMoneyMoovModelsPayrun` | `(data) -> NoFrixionMoneyMoovModelsPayrunEntity` | Create a NoFrixionMoneyMoovModelsPayrun entity instance. |
| `NoFrixionMoneyMoovModelsReportResult` | `(data) -> NoFrixionMoneyMoovModelsReportResultEntity` | Create a NoFrixionMoneyMoovModelsReportResult entity instance. |
| `NoFrixionMoneyMoovModelsRule` | `(data) -> NoFrixionMoneyMoovModelsRuleEntity` | Create a NoFrixionMoneyMoovModelsRule entity instance. |
| `NoFrixionMoneyMoovModelsTransaction` | `(data) -> NoFrixionMoneyMoovModelsTransactionEntity` | Create a NoFrixionMoneyMoovModelsTransaction entity instance. |
| `NoFrixionMoneyMoovModelsTransactionPage` | `(data) -> NoFrixionMoneyMoovModelsTransactionPageEntity` | Create a NoFrixionMoneyMoovModelsTransactionPage entity instance. |
| `NoFrixionMoneyMoovModelsUserInvite` | `(data) -> NoFrixionMoneyMoovModelsUserInviteEntity` | Create a NoFrixionMoneyMoovModelsUserInvite entity instance. |
| `NoFrixionMoneyMoovModelsUserInvitePage` | `(data) -> NoFrixionMoneyMoovModelsUserInvitePageEntity` | Create a NoFrixionMoneyMoovModelsUserInvitePage entity instance. |
| `NoFrixionMoneyMoovModelsUserPage` | `(data) -> NoFrixionMoneyMoovModelsUserPageEntity` | Create a NoFrixionMoneyMoovModelsUserPage entity instance. |
| `NoFrixionMoneyMoovModelsWebhook` | `(data) -> NoFrixionMoneyMoovModelsWebhookEntity` | Create a NoFrixionMoneyMoovModelsWebhook entity instance. |
| `OpenBanking` | `(data) -> OpenBankingEntity` | Create an OpenBanking entity instance. |
| `Payeeverification` | `(data) -> PayeeverificationEntity` | Create a Payeeverification entity instance. |
| `PaymentRequest` | `(data) -> PaymentRequestEntity` | Create a PaymentRequest entity instance. |
| `Payout` | `(data) -> PayoutEntity` | Create a Payout entity instance. |
| `Payrun` | `(data) -> PayrunEntity` | Create a Payrun entity instance. |
| `Reject` | `(data) -> RejectEntity` | Create a Reject entity instance. |
| `Report` | `(data) -> ReportEntity` | Create a Report entity instance. |
| `Rule` | `(data) -> RuleEntity` | Create a Rule entity instance. |
| `Send` | `(data) -> SendEntity` | Create a Send entity instance. |
| `Sendbeneficiary` | `(data) -> SendbeneficiaryEntity` | Create a Sendbeneficiary entity instance. |
| `Tag` | `(data) -> TagEntity` | Create a Tag entity instance. |
| `Token` | `(data) -> TokenEntity` | Create a Token entity instance. |
| `Transaction` | `(data) -> TransactionEntity` | Create a Transaction entity instance. |
| `User` | `(data) -> UserEntity` | Create an User entity instance. |
| `UserInvite` | `(data) -> UserInviteEntity` | Create an UserInvite entity instance. |
| `Virtual` | `(data) -> VirtualEntity` | Create a Virtual entity instance. |
| `Webhook` | `(data) -> WebhookEntity` | Create a Webhook entity instance. |
| `Whoami` | `(data) -> WhoamiEntity` | Create a Whoami entity instance. |
| `Whoamitrustedapp` | `(data) -> WhoamitrustedappEntity` | Create a Whoamitrustedapp entity instance. |

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
| `account_id` |  |
| `account_name` |  |
| `account_supplier_name` |  |
| `account_type` |  |
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
| `physical_account_id` |  |
| `role_i_d` |  |
| `rule` |  |
| `submitted_payouts_balance` |  |
| `submitted_payouts_balance_minor_unit` |  |
| `summary` |  |
| `supplier_physical_account_id` |  |
| `supplier_sepa_instant_status` |  |
| `to_date` |  |
| `xero_bank_feed_connection_status` |  |
| `xero_bank_feed_last_synced_at` |  |
| `xero_bank_feed_sync_last_failed_at` |  |
| `xero_bank_feed_sync_last_failure_reason` |  |
| `xero_bank_feed_sync_status` |  |
| `xero_unsynchronised_transactions_count` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/accounts/{accountID}/{currency}`

#### Beneficiary

| Field | Description |
| --- | --- |
| `approval_callback_url` |  |
| `authentication_method` |  |
| `authorisation` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `beneficiary_event` |  |
| `can_authorise` |  |
| `can_update` |  |
| `created_by` |  |
| `created_by_email_address` |  |
| `currency` |  |
| `destination` |  |
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

Operations: Create, Load, Remove, Update.

API path: `/api/v1/beneficiaries/authorise/{id}`

#### Cancel

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

Operations: Update.

API path: `/api/v1/payouts/cancel/{id}`

#### Disable

| Field | Description |
| --- | --- |
| `approval_callback_url` |  |
| `authentication_method` |  |
| `authorisation` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `beneficiary_event` |  |
| `can_authorise` |  |
| `can_update` |  |
| `created_by` |  |
| `created_by_email_address` |  |
| `currency` |  |
| `destination` |  |
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
| `their_reference` |  |

Operations: Update.

API path: `/api/v1/beneficiaries/disable/{id}`

#### Enable

| Field | Description |
| --- | --- |
| `approval_callback_url` |  |
| `authentication_method` |  |
| `authorisation` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `beneficiary_event` |  |
| `can_authorise` |  |
| `can_update` |  |
| `created_by` |  |
| `created_by_email_address` |  |
| `currency` |  |
| `destination` |  |
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
| `their_reference` |  |

Operations: Update.

API path: `/api/v1/beneficiaries/enable/{id}`

#### Merchant

| Field | Description |
| --- | --- |
| `reason` |  |

Operations: Load, Remove, Update.

API path: `/api/v1/merchants/{merchantID}/payouts/export`

#### Metadata

| Field | Description |
| --- | --- |

Operations: Load.

API path: `/api/v1/metadata/problemnotification`

#### NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage

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

#### NoFrixionBizBizModelsPagingPaymentRequestPage

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
| `status` |  |
| `success_web_hook_url` |  |
| `tag` |  |
| `title` |  |
| `tokenised_card` |  |
| `transaction` |  |
| `use_hosted_payment_page` |  |

Operations: List.

API path: `/api/v1/paymentrequests`

#### NoFrixionBizBizModelsPagingPayoutPage

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

API path: `/api/v1/payouts`

#### NoFrixionBizBizModelsPagingPayrunPage

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
| `payment` |  |
| `payout` |  |
| `payouts_count` |  |
| `schedule_date` |  |
| `source_account` |  |
| `status` |  |
| `total_eur` |  |
| `total_gbp` |  |
| `total_usd` |  |

Operations: List.

API path: `/api/v1/payruns`

#### NoFrixionBizBizModelsPagingRuleEventsPage

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

#### NoFrixionBizBizModelsPagingRulesPage

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

Operations: List.

API path: `/api/v1/rules`

#### NoFrixionBizBizModelsPaymentsCardPayment

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

#### NoFrixionBizBizModelsPaymentsCardPublicKey

| Field | Description |
| --- | --- |
| `jwt` |  |

Operations: Load.

API path: `/api/v1/paymentrequests/{id}/card/publickey`

#### NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries

| Field | Description |
| --- | --- |
| `beneficiary` |  |
| `failed_beneficiary` |  |

Operations: Create.

API path: `/api/v1/beneficiaries/batchcreate`

#### NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment

| Field | Description |
| --- | --- |
| `failed_payment_request` |  |
| `payment_request` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/batchcreate`

#### NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate

| Field | Description |
| --- | --- |
| `failed_role` |  |
| `role` |  |

Operations: Create.

API path: `/api/v1/merchants/{merchantID}/roles/batchcreate`

#### NoFrixionMoneyMoovApiFeaturesUserInvitesCreate

| Field | Description |
| --- | --- |
| `failed_user_invite` |  |
| `user_invite` |  |

Operations: Create.

API path: `/api/v1/userinvites/batchcreate`

#### NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant

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

#### NoFrixionMoneyMoovModelsBatchPayout

| Field | Description |
| --- | --- |
| `approve_url` |  |
| `id` |  |
| `payout` |  |

Operations: Create, Load.

API path: `/api/v1/payouts/batch`

#### NoFrixionMoneyMoovModelsBeneficiaryGroupPage

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

#### NoFrixionMoneyMoovModelsBeneficiaryPage

| Field | Description |
| --- | --- |
| `approval_callback_url` |  |
| `authentication_method` |  |
| `authorisation` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `beneficiary_event` |  |
| `can_authorise` |  |
| `can_update` |  |
| `created_by` |  |
| `created_by_email_address` |  |
| `currency` |  |
| `destination` |  |
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
| `their_reference` |  |

Operations: List.

API path: `/api/v1/beneficiaries`

#### NoFrixionMoneyMoovModelsCardCustomerToken

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

#### NoFrixionMoneyMoovModelsCurrencyCurrencyInfo

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

#### NoFrixionMoneyMoovModelsDirectDebitBatchSubmit

| Field | Description |
| --- | --- |
| `failed_submission` |  |
| `successful_submission` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/directdebit/batchsubmit`

#### NoFrixionMoneyMoovModelsFxRate

| Field | Description |
| --- | --- |
| `destination_currency` |  |
| `exchange_rate` |  |
| `expiry_time` |  |
| `quote_id` |  |
| `source_currency` |  |

Operations: List, Load.

API path: `/api/v1/payouts/fxallheldrates/{source}/{destination}`

#### NoFrixionMoneyMoovModelsIPayment

| Field | Description |
| --- | --- |
| `payment_request_id` |  |
| `response_type` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/payondemand`

#### NoFrixionMoneyMoovModelsMandatesMandate

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

#### NoFrixionMoneyMoovModelsMerchant

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
| `short_name` |  |
| `supported_payment_methods_list` |  |
| `suspension_reason` |  |
| `tag` |  |
| `time_zone_id` |  |
| `trading_name` |  |
| `web_hook_limit` |  |
| `your_role_name` |  |

Operations: List, Load, Update.

API path: `/api/v1/merchants`

#### NoFrixionMoneyMoovModelsMerchantPage

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
| `short_name` |  |
| `supported_payment_methods_list` |  |
| `suspension_reason` |  |
| `tag` |  |
| `time_zone_id` |  |
| `trading_name` |  |
| `web_hook_limit` |  |
| `your_role_name` |  |

Operations: List.

API path: `/api/v1/merchants/paged`

#### NoFrixionMoneyMoovModelsMerchantPayByBankSetting

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

#### NoFrixionMoneyMoovModelsMerchantToken

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

Operations: Create, Load, Update.

API path: `/api/v1/tokens`

#### NoFrixionMoneyMoovModelsMerchantTokenPage

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

Operations: List.

API path: `/api/v1/merchants/{merchantID}/tokens`

#### NoFrixionMoneyMoovModelsNoFrixionVersion

| Field | Description |
| --- | --- |
| `build_version` |  |
| `major_version` |  |
| `minor_version` |  |
| `release_name` |  |

Operations: Load.

API path: `/api/v1/metadata/version`

#### NoFrixionMoneyMoovModelsOpenBankingAccount

| Field | Description |
| --- | --- |
| `account_balance` |  |
| `account_identification` |  |
| `account_name` |  |
| `account_type` |  |
| `balance` |  |
| `consolidated_account_information` |  |
| `currency` |  |
| `description` |  |
| `detail` |  |
| `id` |  |
| `nickname` |  |
| `type` |  |
| `usage_type` |  |

Operations: Load.

API path: `/api/v1/openbanking/accounts/{id}`

#### NoFrixionMoneyMoovModelsOpenBankingConsent

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

#### NoFrixionMoneyMoovModelsOpenBankingTransaction

| Field | Description |
| --- | --- |
| `address_detail` |  |
| `amount` |  |
| `balance` |  |
| `booking_date_time` |  |
| `charge_detail` |  |
| `currency` |  |
| `currency_exchange` |  |
| `date` |  |
| `description` |  |
| `enrichment` |  |
| `gross_amount` |  |
| `id` |  |
| `iso_bank_transaction_code` |  |
| `merchant` |  |
| `payee_detail` |  |
| `payer_detail` |  |
| `proprietary_bank_transaction_code` |  |
| `reference` |  |
| `statement_reference` |  |
| `status` |  |
| `supplementary_data` |  |
| `transaction_amount` |  |
| `transaction_information` |  |
| `transaction_mutability` |  |
| `value_date_time` |  |

Operations: List.

API path: `/api/v1/openbanking/transactions/{id}/{accountID}`

#### NoFrixionMoneyMoovModelsPayment

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

#### NoFrixionMoneyMoovModelsPaymentAccountMinimalPage

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

#### NoFrixionMoneyMoovModelsPaymentAccountPage

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

#### NoFrixionMoneyMoovModelsPaymentInitiation

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

#### NoFrixionMoneyMoovModelsPaymentRequestEvent

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

#### NoFrixionMoneyMoovModelsPaymentRequestMetric

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

#### NoFrixionMoneyMoovModelsPaymentRequestMinimal

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

#### NoFrixionMoneyMoovModelsPaymentRequestResult

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

#### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment

| Field | Description |
| --- | --- |
| `description` |  |
| `id` |  |
| `inserted` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `name` |  |
| `template` |  |

Operations: List.

API path: `/api/v1/paymentrequests/{merchantID}/templates`

#### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2

| Field | Description |
| --- | --- |
| `description` |  |
| `id` |  |
| `inserted` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `name` |  |
| `template` |  |

Operations: Load.

API path: `/api/v1/paymentrequests/{merchantID}/templates/{templateID}`

#### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3

| Field | Description |
| --- | --- |
| `description` |  |
| `id` |  |
| `inserted` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `name` |  |
| `template` |  |

Operations: Update.

API path: `/api/v1/paymentrequests/{merchantID}/templates/{templateID}`

#### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4

| Field | Description |
| --- | --- |

Operations: Remove.

API path: `/api/v1/paymentrequests/{merchantID}/templates/{templateID}`

#### NoFrixionMoneyMoovModelsPayoutKeysetPage

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

#### NoFrixionMoneyMoovModelsPayoutMetric

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

#### NoFrixionMoneyMoovModelsPayoutsPayoutsCreate

| Field | Description |
| --- | --- |
| `failed_payout` |  |
| `payout` |  |

Operations: Create.

API path: `/api/v1/payouts/batchcreate`

#### NoFrixionMoneyMoovModelsPayrun

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

Operations: Create, Load, Update.

API path: `/api/v1/payruns/{merchantID}`

#### NoFrixionMoneyMoovModelsReportResult

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

#### NoFrixionMoneyMoovModelsRule

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

Operations: Create, Load, Update.

API path: `/api/v1/rules`

#### NoFrixionMoneyMoovModelsTransaction

| Field | Description |
| --- | --- |
| `account_id` |  |
| `account_name` |  |
| `account_sequence_number` |  |
| `amount` |  |
| `amount_minor_unit` |  |
| `balance` |  |
| `balance_minor_unit` |  |
| `counterparty` |  |
| `counterparty_summary` |  |
| `currency` |  |
| `description` |  |
| `fx_amount` |  |
| `fx_currency` |  |
| `fx_rate` |  |
| `id` |  |
| `inserted` |  |
| `merchant_id` |  |
| `payment_request_custom_field` |  |
| `payment_request_id` |  |
| `payout_id` |  |
| `raw_reference` |  |
| `rule_id` |  |
| `tag` |  |
| `their_reference` |  |
| `transaction_date` |  |
| `type` |  |
| `virtual_iban` |  |
| `your_reference` |  |

Operations: Load.

API path: `/api/v1/accounts/{accountID}/transactions/{id}`

#### NoFrixionMoneyMoovModelsTransactionPage

| Field | Description |
| --- | --- |
| `account_id` |  |
| `account_name` |  |
| `account_sequence_number` |  |
| `amount` |  |
| `amount_minor_unit` |  |
| `balance` |  |
| `balance_minor_unit` |  |
| `content` |  |
| `counterparty` |  |
| `counterparty_summary` |  |
| `currency` |  |
| `description` |  |
| `fx_amount` |  |
| `fx_currency` |  |
| `fx_rate` |  |
| `id` |  |
| `inserted` |  |
| `merchant_id` |  |
| `page_number` |  |
| `page_size` |  |
| `payment_request_custom_field` |  |
| `payment_request_id` |  |
| `payout_id` |  |
| `raw_reference` |  |
| `rule_id` |  |
| `tag` |  |
| `their_reference` |  |
| `total_page` |  |
| `total_size` |  |
| `transaction_date` |  |
| `type` |  |
| `virtual_iban` |  |
| `your_reference` |  |

Operations: List, Load.

API path: `/api/v1/accounts/{accountID}/transactions`

#### NoFrixionMoneyMoovModelsUserInvite

| Field | Description |
| --- | --- |
| `authorisation_status` |  |
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

Operations: Create, Load.

API path: `/api/v1/userinvites`

#### NoFrixionMoneyMoovModelsUserInvitePage

| Field | Description |
| --- | --- |
| `authorisation_status` |  |
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
| `status` |  |
| `user` |  |
| `user_id` |  |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/userinvitespaged`

#### NoFrixionMoneyMoovModelsUserPage

| Field | Description |
| --- | --- |
| `client_session_timeout` |  |
| `email_address` |  |
| `first_name` |  |
| `id` |  |
| `last_name` |  |
| `passkey_added` |  |
| `permission` |  |
| `roles_with_scope` |  |
| `two_factor_enabled` |  |

Operations: List.

API path: `/api/v1/user/{merchantID}/userspaged`

#### NoFrixionMoneyMoovModelsWebhook

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

Operations: Create, List, Load, Update.

API path: `/api/v1/webhooks`

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

#### PaymentRequest

| Field | Description |
| --- | --- |
| `amount` |  |
| `do_simulate_settlement_failure` |  |
| `error_description` |  |
| `institution` |  |
| `payment_initiation_id` |  |

Operations: Create, Load, Remove, Update.

API path: `/api/v1/paymentrequests/{id}/directdebit`

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
| `tag_id` |  |
| `their_reference` |  |
| `topup_payrun_id` |  |
| `transacted_amount` |  |
| `transacted_fx_amount` |  |
| `transacted_fx_rate` |  |
| `type` |  |
| `user_id` |  |
| `your_reference` |  |

Operations: Create, Load, Remove, Update.

API path: `/api/v1/payouts/batch/submit/{id}`

#### Payrun

| Field | Description |
| --- | --- |
| `id` |  |
| `note` |  |
| `scheduled_date` |  |

Operations: Create, Remove, Update.

API path: `/api/v1/payruns/{id}/request-authorisation`

#### Reject

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
| `their_reference` |  |
| `topup_payrun_id` |  |
| `transacted_amount` |  |
| `transacted_fx_amount` |  |
| `transacted_fx_rate` |  |
| `type` |  |
| `user_id` |  |
| `your_reference` |  |

Operations: Update.

API path: `/api/v1/payouts/reject/{id}`

#### Report

| Field | Description |
| --- | --- |

Operations: Update.

API path: `/api/v1/reports/{id}/initiate`

#### Rule

| Field | Description |
| --- | --- |

Operations: Remove, Update.

API path: `/api/v1/rules/{id}`

#### Send

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
| `tag_id` |  |
| `their_reference` |  |
| `topup_payrun_id` |  |
| `transacted_amount` |  |
| `transacted_fx_amount` |  |
| `transacted_fx_rate` |  |
| `type` |  |
| `user_id` |  |
| `your_reference` |  |

Operations: Create.

API path: `/api/v1/payouts/send`

#### Sendbeneficiary

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
| `tag_id` |  |
| `their_reference` |  |
| `topup_payrun_id` |  |
| `transacted_amount` |  |
| `transacted_fx_amount` |  |
| `transacted_fx_rate` |  |
| `type` |  |
| `user_id` |  |
| `your_reference` |  |

Operations: Create.

API path: `/api/v1/payouts/sendbeneficiary`

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

Operations: Create, Load, Remove.

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

API path: `/api/v1/merchants/{merchantID}/users`

#### UserInvite

| Field | Description |
| --- | --- |

Operations: Create, Remove, Update.

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

Operations: Remove.

API path: `/api/v1/webhooks/{id}`

#### Whoami

| Field | Description |
| --- | --- |
| `client_session_timeout` |  |
| `email_address` |  |
| `first_name` |  |
| `id` |  |
| `last_name` |  |
| `passkey_added` |  |
| `permission` |  |
| `roles_with_scope` |  |
| `two_factor_enabled` |  |

Operations: List.

API path: `/api/v1/metadata/whoami`

#### Whoamitrustedapp

| Field | Description |
| --- | --- |
| `client_session_timeout` |  |
| `email_address` |  |
| `first_name` |  |
| `id` |  |
| `last_name` |  |
| `passkey_added` |  |
| `permission` |  |
| `roles_with_scope` |  |
| `two_factor_enabled` |  |

Operations: List.

API path: `/api/v1/metadata/whoamitrustedapp`



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
| `account_id` | `String` |  |
| `account_name` | `String` |  |
| `account_supplier_name` | `String` |  |
| `account_type` | `String` |  |
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
| `physical_account_id` | `String` |  |
| `role_i_d` | `Array` |  |
| `rule` | `Array` |  |
| `submitted_payouts_balance` | `Float` |  |
| `submitted_payouts_balance_minor_unit` | `Integer` |  |
| `summary` | `String` |  |
| `supplier_physical_account_id` | `String` |  |
| `supplier_sepa_instant_status` | `String` |  |
| `to_date` | `String` |  |
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


### Beneficiary

Create an instance: `beneficiary = client.Beneficiary`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
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
| `beneficiary_event` | `Array` |  |
| `can_authorise` | `Boolean` |  |
| `can_update` | `Boolean` |  |
| `created_by` | `Hash` |  |
| `created_by_email_address` | `String` |  |
| `currency` | `String` |  |
| `destination` | `Hash` |  |
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

#### Example: Create

```ruby
beneficiary = client.Beneficiary.create({
})
```


### Cancel

Create an instance: `cancel = client.Cancel`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

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


### Disable

Create an instance: `disable = client.Disable`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_callback_url` | `String` |  |
| `authentication_method` | `Array` |  |
| `authorisation` | `Array` |  |
| `authorisers_completed_count` | `Integer` |  |
| `authorisers_required_count` | `Integer` |  |
| `beneficiary_event` | `Array` |  |
| `can_authorise` | `Boolean` |  |
| `can_update` | `Boolean` |  |
| `created_by` | `Hash` |  |
| `created_by_email_address` | `String` |  |
| `currency` | `String` |  |
| `destination` | `Hash` |  |
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
| `their_reference` | `String` |  |


### Enable

Create an instance: `enable = client.Enable`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_callback_url` | `String` |  |
| `authentication_method` | `Array` |  |
| `authorisation` | `Array` |  |
| `authorisers_completed_count` | `Integer` |  |
| `authorisers_required_count` | `Integer` |  |
| `beneficiary_event` | `Array` |  |
| `can_authorise` | `Boolean` |  |
| `can_update` | `Boolean` |  |
| `created_by` | `Hash` |  |
| `created_by_email_address` | `String` |  |
| `currency` | `String` |  |
| `destination` | `Hash` |  |
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
| `their_reference` | `String` |  |


### Merchant

Create an instance: `merchant = client.Merchant`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `reason` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Merchant record (raises on error).
merchant = client.Merchant.load({ "merchant_id" => "merchant_id" })
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


### NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage

Create an instance: `no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page = client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage`

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
# list returns an Array of NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage records (raises on error).
no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_pages = client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage.list
```


### NoFrixionBizBizModelsPagingPaymentRequestPage

Create an instance: `no_frixion_biz_biz_models_paging_payment_request_page = client.NoFrixionBizBizModelsPagingPaymentRequestPage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

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
| `status` | `String` |  |
| `success_web_hook_url` | `String` |  |
| `tag` | `Array` |  |
| `title` | `String` |  |
| `tokenised_card` | `Array` |  |
| `transaction` | `Array` |  |
| `use_hosted_payment_page` | `Boolean` |  |

#### Example: List

```ruby
# list returns an Array of NoFrixionBizBizModelsPagingPaymentRequestPage records (raises on error).
no_frixion_biz_biz_models_paging_payment_request_pages = client.NoFrixionBizBizModelsPagingPaymentRequestPage.list
```


### NoFrixionBizBizModelsPagingPayoutPage

Create an instance: `no_frixion_biz_biz_models_paging_payout_page = client.NoFrixionBizBizModelsPagingPayoutPage`

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
# list returns an Array of NoFrixionBizBizModelsPagingPayoutPage records (raises on error).
no_frixion_biz_biz_models_paging_payout_pages = client.NoFrixionBizBizModelsPagingPayoutPage.list
```


### NoFrixionBizBizModelsPagingPayrunPage

Create an instance: `no_frixion_biz_biz_models_paging_payrun_page = client.NoFrixionBizBizModelsPagingPayrunPage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

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
| `payment` | `Array` |  |
| `payout` | `Array` |  |
| `payouts_count` | `Integer` |  |
| `schedule_date` | `String` |  |
| `source_account` | `Array` |  |
| `status` | `String` |  |
| `total_eur` | `Float` |  |
| `total_gbp` | `Float` |  |
| `total_usd` | `Float` |  |

#### Example: List

```ruby
# list returns an Array of NoFrixionBizBizModelsPagingPayrunPage records (raises on error).
no_frixion_biz_biz_models_paging_payrun_pages = client.NoFrixionBizBizModelsPagingPayrunPage.list
```


### NoFrixionBizBizModelsPagingRuleEventsPage

Create an instance: `no_frixion_biz_biz_models_paging_rule_events_page = client.NoFrixionBizBizModelsPagingRuleEventsPage`

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
# list returns an Array of NoFrixionBizBizModelsPagingRuleEventsPage records (raises on error).
no_frixion_biz_biz_models_paging_rule_events_pages = client.NoFrixionBizBizModelsPagingRuleEventsPage.list
```


### NoFrixionBizBizModelsPagingRulesPage

Create an instance: `no_frixion_biz_biz_models_paging_rules_page = client.NoFrixionBizBizModelsPagingRulesPage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

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

#### Example: List

```ruby
# list returns an Array of NoFrixionBizBizModelsPagingRulesPage records (raises on error).
no_frixion_biz_biz_models_paging_rules_pages = client.NoFrixionBizBizModelsPagingRulesPage.list
```


### NoFrixionBizBizModelsPaymentsCardPayment

Create an instance: `no_frixion_biz_biz_models_payments_card_payment = client.NoFrixionBizBizModelsPaymentsCardPayment`

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
no_frixion_biz_biz_models_payments_card_payment = client.NoFrixionBizBizModelsPaymentsCardPayment.create({
  "paymentrequest_id" => "example_paymentrequest_id", # String
})
```


### NoFrixionBizBizModelsPaymentsCardPublicKey

Create an instance: `no_frixion_biz_biz_models_payments_card_public_key = client.NoFrixionBizBizModelsPaymentsCardPublicKey`

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
# load returns the bare NoFrixionBizBizModelsPaymentsCardPublicKey record (raises on error).
no_frixion_biz_biz_models_payments_card_public_key = client.NoFrixionBizBizModelsPaymentsCardPublicKey.load({ "paymentrequest_id" => "paymentrequest_id" })
```


### NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries

Create an instance: `no_frixion_money_moov_api_features_beneficiaries_beneficiaries = client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `beneficiary` | `Array` |  |
| `failed_beneficiary` | `Hash` |  |

#### Example: Create

```ruby
no_frixion_money_moov_api_features_beneficiaries_beneficiaries = client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries.create({
})
```


### NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment

Create an instance: `no_frixion_money_moov_api_features_payment_requests_payment = client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_payment_request` | `Hash` |  |
| `payment_request` | `Array` |  |

#### Example: Create

```ruby
no_frixion_money_moov_api_features_payment_requests_payment = client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment.create({
})
```


### NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate

Create an instance: `no_frixion_money_moov_api_features_permissions_roles_create = client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate`

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
no_frixion_money_moov_api_features_permissions_roles_create = client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate.create({
  "merchant_id" => "example_merchant_id", # String
})
```


### NoFrixionMoneyMoovApiFeaturesUserInvitesCreate

Create an instance: `no_frixion_money_moov_api_features_user_invites_create = client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_user_invite` | `Hash` |  |
| `user_invite` | `Array` |  |

#### Example: Create

```ruby
no_frixion_money_moov_api_features_user_invites_create = client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate.create({
})
```


### NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant

Create an instance: `no_frixion_money_moov_models_authorisation_settings_merchant = client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant`

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
# list returns an Array of NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant records (raises on error).
no_frixion_money_moov_models_authorisation_settings_merchants = client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant.list
```


### NoFrixionMoneyMoovModelsBatchPayout

Create an instance: `no_frixion_money_moov_models_batch_payout = client.NoFrixionMoneyMoovModelsBatchPayout`

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
# load returns the bare NoFrixionMoneyMoovModelsBatchPayout record (raises on error).
no_frixion_money_moov_models_batch_payout = client.NoFrixionMoneyMoovModelsBatchPayout.load({ "id" => "no_frixion_money_moov_models_batch_payout_id" })
```

#### Example: Create

```ruby
no_frixion_money_moov_models_batch_payout = client.NoFrixionMoneyMoovModelsBatchPayout.create({
})
```


### NoFrixionMoneyMoovModelsBeneficiaryGroupPage

Create an instance: `no_frixion_money_moov_models_beneficiary_group_page = client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage`

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
# list returns an Array of NoFrixionMoneyMoovModelsBeneficiaryGroupPage records (raises on error).
no_frixion_money_moov_models_beneficiary_group_pages = client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage.list
```


### NoFrixionMoneyMoovModelsBeneficiaryPage

Create an instance: `no_frixion_money_moov_models_beneficiary_page = client.NoFrixionMoneyMoovModelsBeneficiaryPage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_callback_url` | `String` |  |
| `authentication_method` | `Array` |  |
| `authorisation` | `Array` |  |
| `authorisers_completed_count` | `Integer` |  |
| `authorisers_required_count` | `Integer` |  |
| `beneficiary_event` | `Array` |  |
| `can_authorise` | `Boolean` |  |
| `can_update` | `Boolean` |  |
| `created_by` | `Hash` |  |
| `created_by_email_address` | `String` |  |
| `currency` | `String` |  |
| `destination` | `Hash` |  |
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
| `their_reference` | `String` |  |

#### Example: List

```ruby
# list returns an Array of NoFrixionMoneyMoovModelsBeneficiaryPage records (raises on error).
no_frixion_money_moov_models_beneficiary_pages = client.NoFrixionMoneyMoovModelsBeneficiaryPage.list
```


### NoFrixionMoneyMoovModelsCardCustomerToken

Create an instance: `no_frixion_money_moov_models_card_customer_token = client.NoFrixionMoneyMoovModelsCardCustomerToken`

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
# load returns the bare NoFrixionMoneyMoovModelsCardCustomerToken record (raises on error).
no_frixion_money_moov_models_card_customer_token = client.NoFrixionMoneyMoovModelsCardCustomerToken.load({ "customer_email_address" => "customer_email_address" })
```

#### Example: List

```ruby
# list returns an Array of NoFrixionMoneyMoovModelsCardCustomerToken records (raises on error).
no_frixion_money_moov_models_card_customer_tokens = client.NoFrixionMoneyMoovModelsCardCustomerToken.list
```


### NoFrixionMoneyMoovModelsCurrencyCurrencyInfo

Create an instance: `no_frixion_money_moov_models_currency_currency_info = client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo`

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
# list returns an Array of NoFrixionMoneyMoovModelsCurrencyCurrencyInfo records (raises on error).
no_frixion_money_moov_models_currency_currency_infos = client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo.list
```


### NoFrixionMoneyMoovModelsDirectDebitBatchSubmit

Create an instance: `no_frixion_money_moov_models_direct_debit_batch_submit = client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit`

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
no_frixion_money_moov_models_direct_debit_batch_submit = client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit.create({
})
```


### NoFrixionMoneyMoovModelsFxRate

Create an instance: `no_frixion_money_moov_models_fx_rate = client.NoFrixionMoneyMoovModelsFxRate`

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
# load returns the bare NoFrixionMoneyMoovModelsFxRate record (raises on error).
no_frixion_money_moov_models_fx_rate = client.NoFrixionMoneyMoovModelsFxRate.load({ "destination" => "destination", "source" => "source", "valid_for_minute" => 1 })
```

#### Example: List

```ruby
# list returns an Array of NoFrixionMoneyMoovModelsFxRate records (raises on error).
no_frixion_money_moov_models_fx_rates = client.NoFrixionMoneyMoovModelsFxRate.list
```


### NoFrixionMoneyMoovModelsIPayment

Create an instance: `no_frixion_money_moov_models_i_payment = client.NoFrixionMoneyMoovModelsIPayment`

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
no_frixion_money_moov_models_i_payment = client.NoFrixionMoneyMoovModelsIPayment.create({
})
```


### NoFrixionMoneyMoovModelsMandatesMandate

Create an instance: `no_frixion_money_moov_models_mandates_mandate = client.NoFrixionMoneyMoovModelsMandatesMandate`

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
# load returns the bare NoFrixionMoneyMoovModelsMandatesMandate record (raises on error).
no_frixion_money_moov_models_mandates_mandate = client.NoFrixionMoneyMoovModelsMandatesMandate.load({ "id" => "no_frixion_money_moov_models_mandates_mandate_id" })
```

#### Example: Create

```ruby
no_frixion_money_moov_models_mandates_mandate = client.NoFrixionMoneyMoovModelsMandatesMandate.create({
  "address_line1" => "example_address_line1", # String
  "city" => "example_city", # String
  "country_code" => "example_country_code", # String
  "email_address" => "example_email_address", # String
  "first_name" => "example_first_name", # String
  "last_name" => "example_last_name", # String
  "postal_code" => "example_postal_code", # String
})
```


### NoFrixionMoneyMoovModelsMerchant

Create an instance: `no_frixion_money_moov_models_merchant = client.NoFrixionMoneyMoovModelsMerchant`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
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
# load returns the bare NoFrixionMoneyMoovModelsMerchant record (raises on error).
no_frixion_money_moov_models_merchant = client.NoFrixionMoneyMoovModelsMerchant.load({ "id" => "no_frixion_money_moov_models_merchant_id" })
```

#### Example: List

```ruby
# list returns an Array of NoFrixionMoneyMoovModelsMerchant records (raises on error).
no_frixion_money_moov_models_merchants = client.NoFrixionMoneyMoovModelsMerchant.list
```


### NoFrixionMoneyMoovModelsMerchantPage

Create an instance: `no_frixion_money_moov_models_merchant_page = client.NoFrixionMoneyMoovModelsMerchantPage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

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
| `short_name` | `String` |  |
| `supported_payment_methods_list` | `Array` |  |
| `suspension_reason` | `String` |  |
| `tag` | `Array` |  |
| `time_zone_id` | `String` |  |
| `trading_name` | `String` |  |
| `web_hook_limit` | `Integer` |  |
| `your_role_name` | `String` |  |

#### Example: List

```ruby
# list returns an Array of NoFrixionMoneyMoovModelsMerchantPage records (raises on error).
no_frixion_money_moov_models_merchant_pages = client.NoFrixionMoneyMoovModelsMerchantPage.list
```


### NoFrixionMoneyMoovModelsMerchantPayByBankSetting

Create an instance: `no_frixion_money_moov_models_merchant_pay_by_bank_setting = client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting`

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
# list returns an Array of NoFrixionMoneyMoovModelsMerchantPayByBankSetting records (raises on error).
no_frixion_money_moov_models_merchant_pay_by_bank_settings = client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting.list
```


### NoFrixionMoneyMoovModelsMerchantToken

Create an instance: `no_frixion_money_moov_models_merchant_token = client.NoFrixionMoneyMoovModelsMerchantToken`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
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
# load returns the bare NoFrixionMoneyMoovModelsMerchantToken record (raises on error).
no_frixion_money_moov_models_merchant_token = client.NoFrixionMoneyMoovModelsMerchantToken.load({ "id" => "no_frixion_money_moov_models_merchant_token_id" })
```

#### Example: Create

```ruby
no_frixion_money_moov_models_merchant_token = client.NoFrixionMoneyMoovModelsMerchantToken.create({
  "nonce" => "example_nonce", # String
})
```


### NoFrixionMoneyMoovModelsMerchantTokenPage

Create an instance: `no_frixion_money_moov_models_merchant_token_page = client.NoFrixionMoneyMoovModelsMerchantTokenPage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

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

#### Example: List

```ruby
# list returns an Array of NoFrixionMoneyMoovModelsMerchantTokenPage records (raises on error).
no_frixion_money_moov_models_merchant_token_pages = client.NoFrixionMoneyMoovModelsMerchantTokenPage.list
```


### NoFrixionMoneyMoovModelsNoFrixionVersion

Create an instance: `no_frixion_money_moov_models_no_frixion_version = client.NoFrixionMoneyMoovModelsNoFrixionVersion`

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
# load returns the bare NoFrixionMoneyMoovModelsNoFrixionVersion record (raises on error).
no_frixion_money_moov_models_no_frixion_version = client.NoFrixionMoneyMoovModelsNoFrixionVersion.load()
```


### NoFrixionMoneyMoovModelsOpenBankingAccount

Create an instance: `no_frixion_money_moov_models_open_banking_account = client.NoFrixionMoneyMoovModelsOpenBankingAccount`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_balance` | `Array` |  |
| `account_identification` | `Array` |  |
| `account_name` | `Array` |  |
| `account_type` | `String` |  |
| `balance` | `Float` |  |
| `consolidated_account_information` | `Hash` |  |
| `currency` | `String` |  |
| `description` | `String` |  |
| `detail` | `String` |  |
| `id` | `String` |  |
| `nickname` | `String` |  |
| `type` | `String` |  |
| `usage_type` | `String` |  |

#### Example: Load

```ruby
# load returns the bare NoFrixionMoneyMoovModelsOpenBankingAccount record (raises on error).
no_frixion_money_moov_models_open_banking_account = client.NoFrixionMoneyMoovModelsOpenBankingAccount.load({ "id" => "no_frixion_money_moov_models_open_banking_account_id" })
```


### NoFrixionMoneyMoovModelsOpenBankingConsent

Create an instance: `no_frixion_money_moov_models_open_banking_consent = client.NoFrixionMoneyMoovModelsOpenBankingConsent`

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
# load returns the bare NoFrixionMoneyMoovModelsOpenBankingConsent record (raises on error).
no_frixion_money_moov_models_open_banking_consent = client.NoFrixionMoneyMoovModelsOpenBankingConsent.load({ "id" => "no_frixion_money_moov_models_open_banking_consent_id" })
```

#### Example: List

```ruby
# list returns an Array of NoFrixionMoneyMoovModelsOpenBankingConsent records (raises on error).
no_frixion_money_moov_models_open_banking_consents = client.NoFrixionMoneyMoovModelsOpenBankingConsent.list
```

#### Example: Create

```ruby
no_frixion_money_moov_models_open_banking_consent = client.NoFrixionMoneyMoovModelsOpenBankingConsent.create({
})
```


### NoFrixionMoneyMoovModelsOpenBankingTransaction

Create an instance: `no_frixion_money_moov_models_open_banking_transaction = client.NoFrixionMoneyMoovModelsOpenBankingTransaction`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address_detail` | `Hash` |  |
| `amount` | `Float` |  |
| `balance` | `Hash` |  |
| `booking_date_time` | `String` |  |
| `charge_detail` | `Hash` |  |
| `currency` | `String` |  |
| `currency_exchange` | `Hash` |  |
| `date` | `String` |  |
| `description` | `String` |  |
| `enrichment` | `Hash` |  |
| `gross_amount` | `Hash` |  |
| `id` | `String` |  |
| `iso_bank_transaction_code` | `Hash` |  |
| `merchant` | `Hash` |  |
| `payee_detail` | `Hash` |  |
| `payer_detail` | `Hash` |  |
| `proprietary_bank_transaction_code` | `Hash` |  |
| `reference` | `String` |  |
| `statement_reference` | `Array` |  |
| `status` | `String` |  |
| `supplementary_data` | `Object` |  |
| `transaction_amount` | `Hash` |  |
| `transaction_information` | `Array` |  |
| `transaction_mutability` | `String` |  |
| `value_date_time` | `String` |  |

#### Example: List

```ruby
# list returns an Array of NoFrixionMoneyMoovModelsOpenBankingTransaction records (raises on error).
no_frixion_money_moov_models_open_banking_transactions = client.NoFrixionMoneyMoovModelsOpenBankingTransaction.list
```


### NoFrixionMoneyMoovModelsPayment

Create an instance: `no_frixion_money_moov_models_payment = client.NoFrixionMoneyMoovModelsPayment`

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
# load returns the bare NoFrixionMoneyMoovModelsPayment record (raises on error).
no_frixion_money_moov_models_payment = client.NoFrixionMoneyMoovModelsPayment.load({ "id" => "no_frixion_money_moov_models_payment_id" })
```

#### Example: Create

```ruby
no_frixion_money_moov_models_payment = client.NoFrixionMoneyMoovModelsPayment.create({
  "created_by_user" => {}, # Hash
})
```


### NoFrixionMoneyMoovModelsPaymentAccountMinimalPage

Create an instance: `no_frixion_money_moov_models_payment_account_minimal_page = client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage`

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
# list returns an Array of NoFrixionMoneyMoovModelsPaymentAccountMinimalPage records (raises on error).
no_frixion_money_moov_models_payment_account_minimal_pages = client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage.list
```


### NoFrixionMoneyMoovModelsPaymentAccountPage

Create an instance: `no_frixion_money_moov_models_payment_account_page = client.NoFrixionMoneyMoovModelsPaymentAccountPage`

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
# list returns an Array of NoFrixionMoneyMoovModelsPaymentAccountPage records (raises on error).
no_frixion_money_moov_models_payment_account_pages = client.NoFrixionMoneyMoovModelsPaymentAccountPage.list
```


### NoFrixionMoneyMoovModelsPaymentInitiation

Create an instance: `no_frixion_money_moov_models_payment_initiation = client.NoFrixionMoneyMoovModelsPaymentInitiation`

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
no_frixion_money_moov_models_payment_initiation = client.NoFrixionMoneyMoovModelsPaymentInitiation.create({
  "paymentrequest_id" => "example_paymentrequest_id", # String
})
```


### NoFrixionMoneyMoovModelsPaymentRequestEvent

Create an instance: `no_frixion_money_moov_models_payment_request_event = client.NoFrixionMoneyMoovModelsPaymentRequestEvent`

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
# list returns an Array of NoFrixionMoneyMoovModelsPaymentRequestEvent records (raises on error).
no_frixion_money_moov_models_payment_request_events = client.NoFrixionMoneyMoovModelsPaymentRequestEvent.list
```


### NoFrixionMoneyMoovModelsPaymentRequestMetric

Create an instance: `no_frixion_money_moov_models_payment_request_metric = client.NoFrixionMoneyMoovModelsPaymentRequestMetric`

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
# load returns the bare NoFrixionMoneyMoovModelsPaymentRequestMetric record (raises on error).
no_frixion_money_moov_models_payment_request_metric = client.NoFrixionMoneyMoovModelsPaymentRequestMetric.load()
```


### NoFrixionMoneyMoovModelsPaymentRequestMinimal

Create an instance: `no_frixion_money_moov_models_payment_request_minimal = client.NoFrixionMoneyMoovModelsPaymentRequestMinimal`

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
# list returns an Array of NoFrixionMoneyMoovModelsPaymentRequestMinimal records (raises on error).
no_frixion_money_moov_models_payment_request_minimals = client.NoFrixionMoneyMoovModelsPaymentRequestMinimal.list
```


### NoFrixionMoneyMoovModelsPaymentRequestResult

Create an instance: `no_frixion_money_moov_models_payment_request_result = client.NoFrixionMoneyMoovModelsPaymentRequestResult`

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
# list returns an Array of NoFrixionMoneyMoovModelsPaymentRequestResult records (raises on error).
no_frixion_money_moov_models_payment_request_results = client.NoFrixionMoneyMoovModelsPaymentRequestResult.list
```


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment

Create an instance: `no_frixion_money_moov_models_payment_requests_merchant_payment = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

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

#### Example: List

```ruby
# list returns an Array of NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment records (raises on error).
no_frixion_money_moov_models_payment_requests_merchant_payments = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment.list
```


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2

Create an instance: `no_frixion_money_moov_models_payment_requests_merchant_payment2 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

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
# load returns the bare NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2 record (raises on error).
no_frixion_money_moov_models_payment_requests_merchant_payment2 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2.load({ "paymentrequest_id" => "paymentrequest_id", "template_id" => "template_id" })
```


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3

Create an instance: `no_frixion_money_moov_models_payment_requests_merchant_payment3 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3`

#### Operations

| Method | Description |
| --- | --- |
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


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4

Create an instance: `no_frixion_money_moov_models_payment_requests_merchant_payment4 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |


### NoFrixionMoneyMoovModelsPayoutKeysetPage

Create an instance: `no_frixion_money_moov_models_payout_keyset_page = client.NoFrixionMoneyMoovModelsPayoutKeysetPage`

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
# list returns an Array of NoFrixionMoneyMoovModelsPayoutKeysetPage records (raises on error).
no_frixion_money_moov_models_payout_keyset_pages = client.NoFrixionMoneyMoovModelsPayoutKeysetPage.list
```


### NoFrixionMoneyMoovModelsPayoutMetric

Create an instance: `no_frixion_money_moov_models_payout_metric = client.NoFrixionMoneyMoovModelsPayoutMetric`

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
# load returns the bare NoFrixionMoneyMoovModelsPayoutMetric record (raises on error).
no_frixion_money_moov_models_payout_metric = client.NoFrixionMoneyMoovModelsPayoutMetric.load()
```


### NoFrixionMoneyMoovModelsPayoutsPayoutsCreate

Create an instance: `no_frixion_money_moov_models_payouts_payouts_create = client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_payout` | `Hash` |  |
| `payout` | `Array` |  |

#### Example: Create

```ruby
no_frixion_money_moov_models_payouts_payouts_create = client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate.create({
})
```


### NoFrixionMoneyMoovModelsPayrun

Create an instance: `no_frixion_money_moov_models_payrun = client.NoFrixionMoneyMoovModelsPayrun`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
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
# load returns the bare NoFrixionMoneyMoovModelsPayrun record (raises on error).
no_frixion_money_moov_models_payrun = client.NoFrixionMoneyMoovModelsPayrun.load({ "id" => "no_frixion_money_moov_models_payrun_id" })
```

#### Example: Create

```ruby
no_frixion_money_moov_models_payrun = client.NoFrixionMoneyMoovModelsPayrun.create({
  "id" => "example_id", # String
})
```


### NoFrixionMoneyMoovModelsReportResult

Create an instance: `no_frixion_money_moov_models_report_result = client.NoFrixionMoneyMoovModelsReportResult`

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
# load returns the bare NoFrixionMoneyMoovModelsReportResult record (raises on error).
no_frixion_money_moov_models_report_result = client.NoFrixionMoneyMoovModelsReportResult.load({ "id" => 1, "report_id" => "report_id" })
```


### NoFrixionMoneyMoovModelsRule

Create an instance: `no_frixion_money_moov_models_rule = client.NoFrixionMoneyMoovModelsRule`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
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
# load returns the bare NoFrixionMoneyMoovModelsRule record (raises on error).
no_frixion_money_moov_models_rule = client.NoFrixionMoneyMoovModelsRule.load({ "id" => "no_frixion_money_moov_models_rule_id" })
```

#### Example: Create

```ruby
no_frixion_money_moov_models_rule = client.NoFrixionMoneyMoovModelsRule.create({
  "created_by" => {}, # Hash
  "nonce" => "example_nonce", # String
})
```


### NoFrixionMoneyMoovModelsTransaction

Create an instance: `no_frixion_money_moov_models_transaction = client.NoFrixionMoneyMoovModelsTransaction`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `String` |  |
| `account_name` | `String` |  |
| `account_sequence_number` | `Integer` |  |
| `amount` | `Float` |  |
| `amount_minor_unit` | `Integer` |  |
| `balance` | `Float` |  |
| `balance_minor_unit` | `Integer` |  |
| `counterparty` | `Hash` |  |
| `counterparty_summary` | `String` |  |
| `currency` | `String` |  |
| `description` | `String` |  |
| `fx_amount` | `Float` |  |
| `fx_currency` | `String` |  |
| `fx_rate` | `Float` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `merchant_id` | `String` |  |
| `payment_request_custom_field` | `Hash` |  |
| `payment_request_id` | `String` |  |
| `payout_id` | `String` |  |
| `raw_reference` | `String` |  |
| `rule_id` | `String` |  |
| `tag` | `Array` |  |
| `their_reference` | `String` |  |
| `transaction_date` | `String` |  |
| `type` | `String` |  |
| `virtual_iban` | `String` |  |
| `your_reference` | `String` |  |

#### Example: Load

```ruby
# load returns the bare NoFrixionMoneyMoovModelsTransaction record (raises on error).
no_frixion_money_moov_models_transaction = client.NoFrixionMoneyMoovModelsTransaction.load({ "id" => "no_frixion_money_moov_models_transaction_id" })
```


### NoFrixionMoneyMoovModelsTransactionPage

Create an instance: `no_frixion_money_moov_models_transaction_page = client.NoFrixionMoneyMoovModelsTransactionPage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `String` |  |
| `account_name` | `String` |  |
| `account_sequence_number` | `Integer` |  |
| `amount` | `Float` |  |
| `amount_minor_unit` | `Integer` |  |
| `balance` | `Float` |  |
| `balance_minor_unit` | `Integer` |  |
| `content` | `Array` |  |
| `counterparty` | `Hash` |  |
| `counterparty_summary` | `String` |  |
| `currency` | `String` |  |
| `description` | `String` |  |
| `fx_amount` | `Float` |  |
| `fx_currency` | `String` |  |
| `fx_rate` | `Float` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `merchant_id` | `String` |  |
| `page_number` | `Integer` |  |
| `page_size` | `Integer` |  |
| `payment_request_custom_field` | `Hash` |  |
| `payment_request_id` | `String` |  |
| `payout_id` | `String` |  |
| `raw_reference` | `String` |  |
| `rule_id` | `String` |  |
| `tag` | `Array` |  |
| `their_reference` | `String` |  |
| `total_page` | `Integer` |  |
| `total_size` | `Integer` |  |
| `transaction_date` | `String` |  |
| `type` | `String` |  |
| `virtual_iban` | `String` |  |
| `your_reference` | `String` |  |

#### Example: Load

```ruby
# load returns the bare NoFrixionMoneyMoovModelsTransactionPage record (raises on error).
no_frixion_money_moov_models_transaction_page = client.NoFrixionMoneyMoovModelsTransactionPage.load({ "account_id" => "account_id" })
```

#### Example: List

```ruby
# list returns an Array of NoFrixionMoneyMoovModelsTransactionPage records (raises on error).
no_frixion_money_moov_models_transaction_pages = client.NoFrixionMoneyMoovModelsTransactionPage.list
```


### NoFrixionMoneyMoovModelsUserInvite

Create an instance: `no_frixion_money_moov_models_user_invite = client.NoFrixionMoneyMoovModelsUserInvite`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_status` | `Hash` |  |
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

#### Example: Load

```ruby
# load returns the bare NoFrixionMoneyMoovModelsUserInvite record (raises on error).
no_frixion_money_moov_models_user_invite = client.NoFrixionMoneyMoovModelsUserInvite.load({ "id" => "no_frixion_money_moov_models_user_invite_id" })
```

#### Example: Create

```ruby
no_frixion_money_moov_models_user_invite = client.NoFrixionMoneyMoovModelsUserInvite.create({
  "user" => {}, # Hash
})
```


### NoFrixionMoneyMoovModelsUserInvitePage

Create an instance: `no_frixion_money_moov_models_user_invite_page = client.NoFrixionMoneyMoovModelsUserInvitePage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_status` | `Hash` |  |
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
| `status` | `String` |  |
| `user` | `Hash` |  |
| `user_id` | `String` |  |

#### Example: List

```ruby
# list returns an Array of NoFrixionMoneyMoovModelsUserInvitePage records (raises on error).
no_frixion_money_moov_models_user_invite_pages = client.NoFrixionMoneyMoovModelsUserInvitePage.list
```


### NoFrixionMoneyMoovModelsUserPage

Create an instance: `no_frixion_money_moov_models_user_page = client.NoFrixionMoneyMoovModelsUserPage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

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
| `roles_with_scope` | `Array` |  |
| `two_factor_enabled` | `Boolean` |  |

#### Example: List

```ruby
# list returns an Array of NoFrixionMoneyMoovModelsUserPage records (raises on error).
no_frixion_money_moov_models_user_pages = client.NoFrixionMoneyMoovModelsUserPage.list
```


### NoFrixionMoneyMoovModelsWebhook

Create an instance: `no_frixion_money_moov_models_webhook = client.NoFrixionMoneyMoovModelsWebhook`

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
# load returns the bare NoFrixionMoneyMoovModelsWebhook record (raises on error).
no_frixion_money_moov_models_webhook = client.NoFrixionMoneyMoovModelsWebhook.load({ "id" => "no_frixion_money_moov_models_webhook_id" })
```

#### Example: List

```ruby
# list returns an Array of NoFrixionMoneyMoovModelsWebhook records (raises on error).
no_frixion_money_moov_models_webhooks = client.NoFrixionMoneyMoovModelsWebhook.list
```

#### Example: Create

```ruby
no_frixion_money_moov_models_webhook = client.NoFrixionMoneyMoovModelsWebhook.create({
})
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


### PaymentRequest

Create an instance: `payment_request = client.PaymentRequest`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `Float` |  |
| `do_simulate_settlement_failure` | `Boolean` |  |
| `error_description` | `String` |  |
| `institution` | `String` |  |
| `payment_initiation_id` | `String` |  |

#### Example: Load

```ruby
# load returns the bare PaymentRequest record (raises on error).
payment_request = client.PaymentRequest.load()
```

#### Example: Create

```ruby
payment_request = client.PaymentRequest.create({
  "paymentrequest_id" => "example_paymentrequest_id", # String
})
```


### Payout

Create an instance: `payout = client.Payout`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
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

#### Example: Create

```ruby
payout = client.Payout.create({
})
```


### Payrun

Create an instance: `payrun = client.Payrun`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `String` |  |
| `note` | `String` |  |
| `scheduled_date` | `String` |  |

#### Example: Create

```ruby
payrun = client.Payrun.create({
  "id" => "example_id", # String
})
```


### Reject

Create an instance: `reject = client.Reject`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

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
| `their_reference` | `String` |  |
| `topup_payrun_id` | `String` |  |
| `transacted_amount` | `Float` |  |
| `transacted_fx_amount` | `Float` |  |
| `transacted_fx_rate` | `Float` |  |
| `type` | `String` |  |
| `user_id` | `String` |  |
| `your_reference` | `String` |  |


### Report

Create an instance: `report = client.Report`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |


### Rule

Create an instance: `rule = client.Rule`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |


### Send

Create an instance: `send = client.Send`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

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
| `tag_id` | `Array` |  |
| `their_reference` | `String` |  |
| `topup_payrun_id` | `String` |  |
| `transacted_amount` | `Float` |  |
| `transacted_fx_amount` | `Float` |  |
| `transacted_fx_rate` | `Float` |  |
| `type` | `String` |  |
| `user_id` | `String` |  |
| `your_reference` | `String` |  |

#### Example: Create

```ruby
send = client.Send.create({
  "beneficiary" => {}, # Hash
  "source_account_identifier" => {}, # Hash
})
```


### Sendbeneficiary

Create an instance: `sendbeneficiary = client.Sendbeneficiary`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

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
| `tag_id` | `Array` |  |
| `their_reference` | `String` |  |
| `topup_payrun_id` | `String` |  |
| `transacted_amount` | `Float` |  |
| `transacted_fx_amount` | `Float` |  |
| `transacted_fx_rate` | `Float` |  |
| `type` | `String` |  |
| `user_id` | `String` |  |
| `your_reference` | `String` |  |

#### Example: Create

```ruby
sendbeneficiary = client.Sendbeneficiary.create({
  "beneficiary" => {}, # Hash
  "source_account_identifier" => {}, # Hash
})
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
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Example: Load

```ruby
# load returns the bare Transaction record (raises on error).
transaction = client.Transaction.load({ "sequence_number" => 1, "transaction_id" => "transaction_id" })
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
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Example: Create

```ruby
user_invite = client.UserInvite.create({
  "id" => "example_id", # String
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
| `remove(match)` | Remove the matching entity. |


### Whoami

Create an instance: `whoami = client.Whoami`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

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
| `roles_with_scope` | `Array` |  |
| `two_factor_enabled` | `Boolean` |  |

#### Example: List

```ruby
# list returns an Array of Whoami records (raises on error).
whoamis = client.Whoami.list
```


### Whoamitrustedapp

Create an instance: `whoamitrustedapp = client.Whoamitrustedapp`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

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
| `roles_with_scope` | `Array` |  |
| `two_factor_enabled` | `Boolean` |  |

#### Example: List

```ruby
# list returns an Array of Whoamitrustedapp records (raises on error).
whoamitrustedapps = client.Whoamitrustedapp.list
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
