# Nofrixion Lua SDK



The Lua SDK for the Nofrixion API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:Account()` — each with the same small set of operations (`list`, `load`, `create`, `update`, `remove`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/nofrixion-sdk/releases)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("nofrixion_sdk")

local client = sdk.new({
  apikey = os.getenv("NOFRIXION_APIKEY"),
})
```

### 2. List account records

Entity operations return `(value, err)`. For `list`, `value` is the
array of records itself — iterate it directly (there is no wrapper).

```lua
local accounts, err = client:Account():list()
if err then error(err) end

for _, item in ipairs(accounts) do
  print(item["id"], item["account_id"])
end
```

### 3. Load a merchant

Merchant is nested under merchant, so provide the `merchant_id`.

```lua
local merchant, err = client:Merchant():load({ merchant_id = "example_merchant_id" })
if err then error(err) end
print(merchant)
```

### 4. Create, update, and remove

```lua
-- Create
local created, err = client:Account():create({ account_id = "example_account_id", currency = "example_currency" })
if err then error(err) end

-- Update
client:Account():update({ id = created["id"], account_id = "example_account_id", amount = 1 })

-- Remove
client:Account():remove({ id = created["id"] })
```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local accounts, err = client:Account():list()
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:Account():list()
-- result is the returned data; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
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
cd lua && busted test/
```


## Reference

### NofrixionSDK

```lua
local sdk = require("nofrixion_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### NofrixionSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
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
| `load` | `(reqmatch, ctrl) -> any, err` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> any, err` | List entities matching the criteria. |
| `create` | `(reqdata, ctrl) -> any, err` | Create a new entity. |
| `update` | `(reqdata, ctrl) -> any, err` | Update an existing entity. |
| `remove` | `(reqmatch, ctrl) -> any, err` | Remove an entity. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the operation's
data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `load` / `create` / `update` / `remove` | the entity record (a `table`) |
| `list` | an array (`table`) of entity records |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local account, err = client:Account():load({ id = "example_id" })
    if err then error(err) end
    -- account is the loaded record

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

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

Create an instance: `local account = client:Account(nil)`

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
| `account_id` | `string` |  |
| `account_name` | `string` |  |
| `account_supplier_name` | `string` |  |
| `account_type` | `string` |  |
| `available_balance` | `number` |  |
| `available_balance_minor_unit` | `number` |  |
| `balance` | `number` |  |
| `balance_minor_unit` | `number` |  |
| `bank_name` | `string` |  |
| `consent_id` | `string` |  |
| `created_by` | `table` |  |
| `created_by_display_name` | `string` |  |
| `currency` | `string` |  |
| `default_payment_rail` | `string` |  |
| `display_name` | `string` |  |
| `expiry_date` | `string` |  |
| `external_account_icon` | `string` |  |
| `format` | `string` |  |
| `from_date` | `string` |  |
| `id` | `string` |  |
| `identifier` | `table` |  |
| `inserted` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_connected_account` | `boolean` |  |
| `is_default` | `boolean` |  |
| `is_trust_account` | `boolean` |  |
| `is_virtual` | `boolean` |  |
| `last_transaction` | `table` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `physical_account_id` | `string` |  |
| `role_i_d` | `table` |  |
| `rule` | `table` |  |
| `submitted_payouts_balance` | `number` |  |
| `submitted_payouts_balance_minor_unit` | `number` |  |
| `summary` | `string` |  |
| `supplier_physical_account_id` | `string` |  |
| `supplier_sepa_instant_status` | `string` |  |
| `to_date` | `string` |  |
| `xero_bank_feed_connection_status` | `string` |  |
| `xero_bank_feed_last_synced_at` | `string` |  |
| `xero_bank_feed_sync_last_failed_at` | `string` |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` |  |
| `xero_bank_feed_sync_status` | `string` |  |
| `xero_unsynchronised_transactions_count` | `number` |  |

#### Example: Load

```lua
local account, err = client:Account():load({ id = "account_id" })
```

#### Example: List

```lua
local accounts, err = client:Account():list()
```

#### Example: Create

```lua
local account, err = client:Account():create({
})
```


### Beneficiary

Create an instance: `local beneficiary = client:Beneficiary(nil)`

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
| `approval_callback_url` | `string` |  |
| `authentication_method` | `table` |  |
| `authorisation` | `table` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `beneficiary_event` | `table` |  |
| `can_authorise` | `boolean` |  |
| `can_update` | `boolean` |  |
| `created_by` | `table` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `destination` | `table` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_enabled` | `boolean` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `source_account` | `table` |  |
| `source_account_i_d` | `table` |  |
| `their_reference` | `string` |  |

#### Example: Load

```lua
local beneficiary, err = client:Beneficiary():load({ id = "beneficiary_id" })
```

#### Example: Create

```lua
local beneficiary, err = client:Beneficiary():create({
})
```


### Cancel

Create an instance: `local cancel = client:Cancel(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `amount` | `number` |  |
| `amount_minor_unit` | `number` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `table` |  |
| `authorisation` | `table` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `table` |  |
| `can_authorise` | `boolean` |  |
| `can_process` | `boolean` |  |
| `can_update` | `boolean` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `table` |  |
| `document` | `table` |  |
| `event` | `table` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `number` |  |
| `fx_destination_amount_minor_unit` | `number` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `number` |  |
| `fx_use_destination_amount` | `boolean` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_failed` | `boolean` |  |
| `is_settled` | `boolean` |  |
| `is_submitted` | `boolean` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `rule` | `table` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `boolean` |  |
| `source_account_available_balance` | `number` |  |
| `source_account_available_balance_minor_unit` | `number` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `table` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `table` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `number` |  |
| `transacted_fx_amount` | `number` |  |
| `transacted_fx_rate` | `number` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |


### Disable

Create an instance: `local disable = client:Disable(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_callback_url` | `string` |  |
| `authentication_method` | `table` |  |
| `authorisation` | `table` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `beneficiary_event` | `table` |  |
| `can_authorise` | `boolean` |  |
| `can_update` | `boolean` |  |
| `created_by` | `table` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `destination` | `table` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_enabled` | `boolean` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `source_account` | `table` |  |
| `their_reference` | `string` |  |


### Enable

Create an instance: `local enable = client:Enable(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_callback_url` | `string` |  |
| `authentication_method` | `table` |  |
| `authorisation` | `table` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `beneficiary_event` | `table` |  |
| `can_authorise` | `boolean` |  |
| `can_update` | `boolean` |  |
| `created_by` | `table` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `destination` | `table` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_enabled` | `boolean` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `source_account` | `table` |  |
| `their_reference` | `string` |  |


### Merchant

Create an instance: `local merchant = client:Merchant(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `reason` | `string` |  |

#### Example: Load

```lua
local merchant, err = client:Merchant():load({ merchant_id = "merchant_id" })
```


### Metadata

Create an instance: `local metadata = client:Metadata(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```lua
local metadata, err = client:Metadata():load()
```


### NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage

Create an instance: `local no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page = client:NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approved_at` | `string` |  |
| `currency` | `string` |  |
| `customer_account_number` | `string` |  |
| `customer_city` | `string` |  |
| `customer_country_code` | `string` |  |
| `customer_country_name` | `string` |  |
| `customer_email_address` | `string` |  |
| `customer_first_name` | `string` |  |
| `customer_iban` | `string` |  |
| `customer_last_name` | `string` |  |
| `customer_sort_code` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_recurring` | `boolean` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `reference` | `string` |  |
| `status` | `string` |  |
| `supplier_bank_account_id` | `string` |  |
| `supplier_customer_id` | `string` |  |
| `supplier_mandate_id` | `string` |  |
| `supplier_name` | `string` |  |
| `supplier_status` | `string` |  |

#### Example: List

```lua
local no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_pages, err = client:NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage():list()
```


### NoFrixionBizBizModelsPagingPaymentRequestPage

Create an instance: `local no_frixion_biz_biz_models_paging_payment_request_page = client:NoFrixionBizBizModelsPagingPaymentRequestPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `table` |  |
| `amount` | `number` |  |
| `amount_pending` | `number` |  |
| `amount_received` | `number` |  |
| `amount_refunded` | `number` |  |
| `auto_send_receipt` | `boolean` |  |
| `base_origin_url` | `string` |  |
| `callback_url` | `string` |  |
| `card_authorize_only` | `boolean` |  |
| `card_create_token` | `boolean` |  |
| `card_create_token_mode` | `string` |  |
| `card_ignore_cvn` | `boolean` |  |
| `card_processor_merchant_id` | `string` |  |
| `card_stripe_payment_intent_id` | `string` |  |
| `card_stripe_payment_intent_secret` | `string` |  |
| `created_by_user` | `table` |  |
| `currency` | `string` |  |
| `custom_field` | `table` |  |
| `customer_email_address` | `string` |  |
| `customer_id` | `string` |  |
| `customer_name` | `string` |  |
| `description` | `string` |  |
| `destination_account` | `table` |  |
| `direct_debit_payment` | `table` |  |
| `due_date` | `string` |  |
| `event` | `table` |  |
| `failure_callback_url` | `string` |  |
| `field_display_setting` | `table` |  |
| `formatted_amount` | `string` |  |
| `hosted_pay_checkout_url` | `string` |  |
| `id` | `string` |  |
| `ignore_address_verification` | `boolean` |  |
| `inserted` | `string` |  |
| `inserted_sortable` | `string` |  |
| `is_archived` | `boolean` |  |
| `jwk` | `string` |  |
| `last_updated` | `string` |  |
| `lightning_invoice` | `string` |  |
| `lightning_invoice_expires_at` | `string` |  |
| `merchant_direct_debit_mandate_id` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `notification_email_address` | `string` |  |
| `notification_role_i_d` | `table` |  |
| `order_id` | `string` |  |
| `partial_payment_method` | `string` |  |
| `partial_payment_step` | `string` |  |
| `payment_attempt` | `table` |  |
| `payment_method` | `table` |  |
| `payment_processor` | `string` |  |
| `payrun_id` | `string` |  |
| `pisp_account_id` | `string` |  |
| `priority_bank_id` | `string` |  |
| `result` | `table` |  |
| `sandbox_settle_delay_in_second` | `number` |  |
| `shipping_address` | `table` |  |
| `status` | `string` |  |
| `success_web_hook_url` | `string` |  |
| `tag` | `table` |  |
| `title` | `string` |  |
| `tokenised_card` | `table` |  |
| `transaction` | `table` |  |
| `use_hosted_payment_page` | `boolean` |  |

#### Example: List

```lua
local no_frixion_biz_biz_models_paging_payment_request_pages, err = client:NoFrixionBizBizModelsPagingPaymentRequestPage():list()
```


### NoFrixionBizBizModelsPagingPayoutPage

Create an instance: `local no_frixion_biz_biz_models_paging_payout_page = client:NoFrixionBizBizModelsPagingPayoutPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `amount` | `number` |  |
| `amount_minor_unit` | `number` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `table` |  |
| `authorisation` | `table` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `table` |  |
| `can_authorise` | `boolean` |  |
| `can_process` | `boolean` |  |
| `can_update` | `boolean` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `table` |  |
| `document` | `table` |  |
| `event` | `table` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `number` |  |
| `fx_destination_amount_minor_unit` | `number` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `number` |  |
| `fx_use_destination_amount` | `boolean` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_failed` | `boolean` |  |
| `is_settled` | `boolean` |  |
| `is_submitted` | `boolean` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `rule` | `table` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `boolean` |  |
| `source_account_available_balance` | `number` |  |
| `source_account_available_balance_minor_unit` | `number` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `table` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `table` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `number` |  |
| `transacted_fx_amount` | `number` |  |
| `transacted_fx_rate` | `number` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |

#### Example: List

```lua
local no_frixion_biz_biz_models_paging_payout_pages, err = client:NoFrixionBizBizModelsPagingPayoutPage():list()
```


### NoFrixionBizBizModelsPagingPayrunPage

Create an instance: `local no_frixion_biz_biz_models_paging_payrun_page = client:NoFrixionBizBizModelsPagingPayrunPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation` | `table` |  |
| `authorisation_date` | `string` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `batch_payout_id` | `string` |  |
| `can_authorise` | `boolean` |  |
| `can_delete` | `boolean` |  |
| `can_edit` | `boolean` |  |
| `event` | `table` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice` | `table` |  |
| `invoices_minimal` | `table` |  |
| `is_archived` | `boolean` |  |
| `last_updated` | `string` |  |
| `last_updated_by` | `table` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `payment` | `table` |  |
| `payout` | `table` |  |
| `payouts_count` | `number` |  |
| `schedule_date` | `string` |  |
| `source_account` | `table` |  |
| `status` | `string` |  |
| `total_eur` | `number` |  |
| `total_gbp` | `number` |  |
| `total_usd` | `number` |  |

#### Example: List

```lua
local no_frixion_biz_biz_models_paging_payrun_pages, err = client:NoFrixionBizBizModelsPagingPayrunPage():list()
```


### NoFrixionBizBizModelsPagingRuleEventsPage

Create an instance: `local no_frixion_biz_biz_models_paging_rule_events_page = client:NoFrixionBizBizModelsPagingRuleEventsPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_message` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_authorise_to_enable` | `boolean` |  |
| `message` | `string` |  |
| `raw_response` | `string` |  |
| `rule_event_type` | `string` |  |
| `rule_id` | `string` |  |
| `user` | `table` |  |

#### Example: List

```lua
local no_frixion_biz_biz_models_paging_rule_events_pages, err = client:NoFrixionBizBizModelsPagingRuleEventsPage():list()
```


### NoFrixionBizBizModelsPagingRulesPage

Create an instance: `local no_frixion_biz_biz_models_paging_rules_page = client:NoFrixionBizBizModelsPagingRulesPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account` | `table` |  |
| `account_id` | `string` |  |
| `approve_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `table` |  |
| `authorisation` | `table` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `can_authorise` | `boolean` |  |
| `created_by` | `table` |  |
| `description` | `string` |  |
| `end_at` | `string` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_disabled` | `boolean` |  |
| `last_executed_at` | `string` |  |
| `last_run_at_transaction_date` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `on_approved_web_hook_url` | `string` |  |
| `on_execution_error_web_hook_url` | `string` |  |
| `on_execution_success_web_hook_url` | `string` |  |
| `start_at` | `string` |  |
| `status` | `string` |  |
| `sweep_action` | `table` |  |
| `time_zone_id` | `string` |  |
| `trigger_cron_expression` | `string` |  |
| `trigger_on_pay_in` | `boolean` |  |
| `user_id` | `string` |  |
| `web_hook_secret` | `string` |  |

#### Example: List

```lua
local no_frixion_biz_biz_models_paging_rules_pages, err = client:NoFrixionBizBizModelsPagingRulesPage():list()
```


### NoFrixionBizBizModelsPaymentsCardPayment

Create an instance: `local no_frixion_biz_biz_models_payments_card_payment = client:NoFrixionBizBizModelsPaymentsCardPayment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorized_amount` | `string` |  |
| `currency_code` | `string` |  |
| `is_payer_authentication_required` | `boolean` |  |
| `is_soft_decline` | `boolean` |  |
| `payer_authentication_access_token` | `string` |  |
| `payer_authentication_merchant_data` | `string` |  |
| `payer_authentication_url` | `string` |  |
| `payer_authentication_window_height` | `number` |  |
| `payer_authentication_window_width` | `number` |  |
| `payment_request_callback_url` | `string` |  |
| `payment_request_id` | `string` |  |
| `request_id` | `string` |  |
| `response_code` | `string` |  |
| `response_type` | `string` |  |
| `status` | `string` |  |
| `three_ds_redirect_url` | `string` |  |
| `transaction_id` | `string` |  |

#### Example: Create

```lua
local no_frixion_biz_biz_models_payments_card_payment, err = client:NoFrixionBizBizModelsPaymentsCardPayment():create({
  paymentrequest_id = "example_paymentrequest_id", -- string
})
```


### NoFrixionBizBizModelsPaymentsCardPublicKey

Create an instance: `local no_frixion_biz_biz_models_payments_card_public_key = client:NoFrixionBizBizModelsPaymentsCardPublicKey(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `jwt` | `string` |  |

#### Example: Load

```lua
local no_frixion_biz_biz_models_payments_card_public_key, err = client:NoFrixionBizBizModelsPaymentsCardPublicKey():load({ paymentrequest_id = "paymentrequest_id" })
```


### NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries

Create an instance: `local no_frixion_money_moov_api_features_beneficiaries_beneficiaries = client:NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `beneficiary` | `table` |  |
| `failed_beneficiary` | `table` |  |

#### Example: Create

```lua
local no_frixion_money_moov_api_features_beneficiaries_beneficiaries, err = client:NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries():create({
})
```


### NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment

Create an instance: `local no_frixion_money_moov_api_features_payment_requests_payment = client:NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_payment_request` | `table` |  |
| `payment_request` | `table` |  |

#### Example: Create

```lua
local no_frixion_money_moov_api_features_payment_requests_payment, err = client:NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment():create({
})
```


### NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate

Create an instance: `local no_frixion_money_moov_api_features_permissions_roles_create = client:NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_role` | `table` |  |
| `role` | `table` |  |

#### Example: Create

```lua
local no_frixion_money_moov_api_features_permissions_roles_create, err = client:NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate():create({
  merchant_id = "example_merchant_id", -- string
})
```


### NoFrixionMoneyMoovApiFeaturesUserInvitesCreate

Create an instance: `local no_frixion_money_moov_api_features_user_invites_create = client:NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_user_invite` | `table` |  |
| `user_invite` | `table` |  |

#### Example: Create

```lua
local no_frixion_money_moov_api_features_user_invites_create, err = client:NoFrixionMoneyMoovApiFeaturesUserInvitesCreate():create({
})
```


### NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant

Create an instance: `local no_frixion_money_moov_models_authorisation_settings_merchant = client:NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount_lower` | `number` |  |
| `amount_upper` | `number` |  |
| `authorisation_type` | `string` |  |
| `beneficiaries_only` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_editor_cant_authorise` | `boolean` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `number_of_authoriser` | `number` |  |
| `role_setting` | `table` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_authorisation_settings_merchants, err = client:NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant():list()
```


### NoFrixionMoneyMoovModelsBatchPayout

Create an instance: `local no_frixion_money_moov_models_batch_payout = client:NoFrixionMoneyMoovModelsBatchPayout(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approve_url` | `string` |  |
| `id` | `string` |  |
| `payout` | `table` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_batch_payout, err = client:NoFrixionMoneyMoovModelsBatchPayout():load({ id = "no_frixion_money_moov_models_batch_payout_id" })
```

#### Example: Create

```lua
local no_frixion_money_moov_models_batch_payout, err = client:NoFrixionMoneyMoovModelsBatchPayout():create({
})
```


### NoFrixionMoneyMoovModelsBeneficiaryGroupPage

Create an instance: `local no_frixion_money_moov_models_beneficiary_group_page = client:NoFrixionMoneyMoovModelsBeneficiaryGroupPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `group_member` | `table` |  |
| `group_name` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_beneficiary_group_pages, err = client:NoFrixionMoneyMoovModelsBeneficiaryGroupPage():list()
```


### NoFrixionMoneyMoovModelsBeneficiaryPage

Create an instance: `local no_frixion_money_moov_models_beneficiary_page = client:NoFrixionMoneyMoovModelsBeneficiaryPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_callback_url` | `string` |  |
| `authentication_method` | `table` |  |
| `authorisation` | `table` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `beneficiary_event` | `table` |  |
| `can_authorise` | `boolean` |  |
| `can_update` | `boolean` |  |
| `created_by` | `table` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `destination` | `table` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_enabled` | `boolean` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `source_account` | `table` |  |
| `their_reference` | `string` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_beneficiary_pages, err = client:NoFrixionMoneyMoovModelsBeneficiaryPage():list()
```


### NoFrixionMoneyMoovModelsCardCustomerToken

Create an instance: `local no_frixion_money_moov_models_card_customer_token = client:NoFrixionMoneyMoovModelsCardCustomerToken(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `card_type` | `string` |  |
| `customer_email_address` | `string` |  |
| `expiry_month` | `string` |  |
| `expiry_year` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_four_digit` | `string` |  |
| `last_updated` | `string` |  |
| `masked_card_number` | `string` |  |
| `merchant_id` | `string` |  |
| `payment_request_id` | `string` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_card_customer_token, err = client:NoFrixionMoneyMoovModelsCardCustomerToken():load({ customer_email_address = "customer_email_address" })
```

#### Example: List

```lua
local no_frixion_money_moov_models_card_customer_tokens, err = client:NoFrixionMoneyMoovModelsCardCustomerToken():list()
```


### NoFrixionMoneyMoovModelsCurrencyCurrencyInfo

Create an instance: `local no_frixion_money_moov_models_currency_currency_info = client:NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `code` | `string` |  |
| `decimal` | `number` |  |
| `is_fiat` | `boolean` |  |
| `iso4217_alpha_code` | `string` |  |
| `iso4217_numeric_code` | `string` |  |
| `symbol` | `string` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_currency_currency_infos, err = client:NoFrixionMoneyMoovModelsCurrencyCurrencyInfo():list()
```


### NoFrixionMoneyMoovModelsDirectDebitBatchSubmit

Create an instance: `local no_frixion_money_moov_models_direct_debit_batch_submit = client:NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_submission` | `table` |  |
| `successful_submission` | `table` |  |

#### Example: Create

```lua
local no_frixion_money_moov_models_direct_debit_batch_submit, err = client:NoFrixionMoneyMoovModelsDirectDebitBatchSubmit():create({
})
```


### NoFrixionMoneyMoovModelsFxRate

Create an instance: `local no_frixion_money_moov_models_fx_rate = client:NoFrixionMoneyMoovModelsFxRate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `destination_currency` | `string` |  |
| `exchange_rate` | `number` |  |
| `expiry_time` | `string` |  |
| `quote_id` | `string` |  |
| `source_currency` | `string` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_fx_rate, err = client:NoFrixionMoneyMoovModelsFxRate():load({ destination = "destination", source = "source", valid_for_minute = 1 })
```

#### Example: List

```lua
local no_frixion_money_moov_models_fx_rates, err = client:NoFrixionMoneyMoovModelsFxRate():list()
```


### NoFrixionMoneyMoovModelsIPayment

Create an instance: `local no_frixion_money_moov_models_i_payment = client:NoFrixionMoneyMoovModelsIPayment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `payment_request_id` | `string` |  |
| `response_type` | `string` |  |

#### Example: Create

```lua
local no_frixion_money_moov_models_i_payment, err = client:NoFrixionMoneyMoovModelsIPayment():create({
})
```


### NoFrixionMoneyMoovModelsMandatesMandate

Create an instance: `local no_frixion_money_moov_models_mandates_mandate = client:NoFrixionMoneyMoovModelsMandatesMandate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_number` | `string` |  |
| `address_line1` | `string` |  |
| `address_line2` | `string` |  |
| `approved_at` | `string` |  |
| `city` | `string` |  |
| `country_code` | `string` |  |
| `currency` | `string` |  |
| `customer_account_number` | `string` |  |
| `customer_city` | `string` |  |
| `customer_country_code` | `string` |  |
| `customer_country_name` | `string` |  |
| `customer_email_address` | `string` |  |
| `customer_first_name` | `string` |  |
| `customer_iban` | `string` |  |
| `customer_last_name` | `string` |  |
| `customer_sort_code` | `string` |  |
| `email_address` | `string` |  |
| `first_name` | `string` |  |
| `iban` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_recurring` | `boolean` |  |
| `last_name` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `postal_code` | `string` |  |
| `reference` | `string` |  |
| `sort_code` | `string` |  |
| `status` | `string` |  |
| `supplier_bank_account_id` | `string` |  |
| `supplier_customer_id` | `string` |  |
| `supplier_mandate_id` | `string` |  |
| `supplier_name` | `string` |  |
| `supplier_status` | `string` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_mandates_mandate, err = client:NoFrixionMoneyMoovModelsMandatesMandate():load({ id = "no_frixion_money_moov_models_mandates_mandate_id" })
```

#### Example: Create

```lua
local no_frixion_money_moov_models_mandates_mandate, err = client:NoFrixionMoneyMoovModelsMandatesMandate():create({
  address_line1 = "example_address_line1", -- string
  city = "example_city", -- string
  country_code = "example_country_code", -- string
  email_address = "example_email_address", -- string
  first_name = "example_first_name", -- string
  last_name = "example_last_name", -- string
  postal_code = "example_postal_code", -- string
})
```


### NoFrixionMoneyMoovModelsMerchant

Create an instance: `local no_frixion_money_moov_models_merchant = client:NoFrixionMoneyMoovModelsMerchant(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_currency` | `table` |  |
| `can_have_trust_account` | `boolean` |  |
| `card_payment_processor` | `string` |  |
| `company_id` | `string` |  |
| `display_qr_on_hosted_pay` | `boolean` |  |
| `hosted_pay_version` | `number` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_blocked` | `boolean` |  |
| `is_exited` | `boolean` |  |
| `is_suspended` | `boolean` |  |
| `jurisdiction` | `string` |  |
| `logo_url_png` | `string` |  |
| `logo_url_svg` | `string` |  |
| `merchant_category_code` | `string` |  |
| `name` | `string` |  |
| `note` | `string` |  |
| `parent_merchant` | `table` |  |
| `payment_account` | `table` |  |
| `payment_account_limit` | `number` |  |
| `short_name` | `string` |  |
| `supported_payment_methods_list` | `table` |  |
| `suspension_reason` | `string` |  |
| `tag` | `table` |  |
| `time_zone_id` | `string` |  |
| `trading_name` | `string` |  |
| `web_hook_limit` | `number` |  |
| `your_role_name` | `string` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_merchant, err = client:NoFrixionMoneyMoovModelsMerchant():load({ id = "no_frixion_money_moov_models_merchant_id" })
```

#### Example: List

```lua
local no_frixion_money_moov_models_merchants, err = client:NoFrixionMoneyMoovModelsMerchant():list()
```


### NoFrixionMoneyMoovModelsMerchantPage

Create an instance: `local no_frixion_money_moov_models_merchant_page = client:NoFrixionMoneyMoovModelsMerchantPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_currency` | `table` |  |
| `can_have_trust_account` | `boolean` |  |
| `card_payment_processor` | `string` |  |
| `company_id` | `string` |  |
| `display_qr_on_hosted_pay` | `boolean` |  |
| `hosted_pay_version` | `number` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_blocked` | `boolean` |  |
| `is_exited` | `boolean` |  |
| `is_suspended` | `boolean` |  |
| `jurisdiction` | `string` |  |
| `logo_url_png` | `string` |  |
| `logo_url_svg` | `string` |  |
| `merchant_category_code` | `string` |  |
| `name` | `string` |  |
| `note` | `string` |  |
| `parent_merchant` | `table` |  |
| `payment_account` | `table` |  |
| `payment_account_limit` | `number` |  |
| `short_name` | `string` |  |
| `supported_payment_methods_list` | `table` |  |
| `suspension_reason` | `string` |  |
| `tag` | `table` |  |
| `time_zone_id` | `string` |  |
| `trading_name` | `string` |  |
| `web_hook_limit` | `number` |  |
| `your_role_name` | `string` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_merchant_pages, err = client:NoFrixionMoneyMoovModelsMerchantPage():list()
```


### NoFrixionMoneyMoovModelsMerchantPayByBankSetting

Create an instance: `local no_frixion_money_moov_models_merchant_pay_by_bank_setting = client:NoFrixionMoneyMoovModelsMerchantPayByBankSetting(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bank_country_code` | `table` |  |
| `bank_id` | `string` |  |
| `bank_name` | `string` |  |
| `business_institution_id` | `string` |  |
| `currency` | `string` |  |
| `logo` | `string` |  |
| `message` | `string` |  |
| `message_image_url` | `string` |  |
| `order` | `number` |  |
| `personal_institution_id` | `string` |  |
| `processor` | `string` |  |
| `warning_heading` | `string` |  |
| `warning_message` | `string` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_merchant_pay_by_bank_settings, err = client:NoFrixionMoneyMoovModelsMerchantPayByBankSetting():list()
```


### NoFrixionMoneyMoovModelsMerchantToken

Create an instance: `local no_frixion_money_moov_models_merchant_token = client:NoFrixionMoneyMoovModelsMerchantToken(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authentication_method` | `table` |  |
| `authorisation` | `table` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `can_authorise` | `boolean` |  |
| `description` | `string` |  |
| `expires_at` | `string` |  |
| `has_current_user_authorised` | `boolean` |  |
| `hmac_algorithm` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `ip_address_whitelist` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_enabled` | `boolean` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `nonce` | `string` |  |
| `permission_type` | `table` |  |
| `request_signature_version` | `number` |  |
| `shared_secret_algorithm` | `string` |  |
| `shared_secret_base64` | `string` |  |
| `token` | `string` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_merchant_token, err = client:NoFrixionMoneyMoovModelsMerchantToken():load({ id = "no_frixion_money_moov_models_merchant_token_id" })
```

#### Example: Create

```lua
local no_frixion_money_moov_models_merchant_token, err = client:NoFrixionMoneyMoovModelsMerchantToken():create({
  nonce = "example_nonce", -- string
})
```


### NoFrixionMoneyMoovModelsMerchantTokenPage

Create an instance: `local no_frixion_money_moov_models_merchant_token_page = client:NoFrixionMoneyMoovModelsMerchantTokenPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authentication_method` | `table` |  |
| `authorisation` | `table` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `can_authorise` | `boolean` |  |
| `description` | `string` |  |
| `expires_at` | `string` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `ip_address_whitelist` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_enabled` | `boolean` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `nonce` | `string` |  |
| `permission_type` | `table` |  |
| `request_signature_version` | `number` |  |
| `shared_secret_algorithm` | `string` |  |
| `shared_secret_base64` | `string` |  |
| `token` | `string` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_merchant_token_pages, err = client:NoFrixionMoneyMoovModelsMerchantTokenPage():list()
```


### NoFrixionMoneyMoovModelsNoFrixionVersion

Create an instance: `local no_frixion_money_moov_models_no_frixion_version = client:NoFrixionMoneyMoovModelsNoFrixionVersion(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `build_version` | `number` |  |
| `major_version` | `number` |  |
| `minor_version` | `number` |  |
| `release_name` | `string` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_no_frixion_version, err = client:NoFrixionMoneyMoovModelsNoFrixionVersion():load()
```


### NoFrixionMoneyMoovModelsOpenBankingAccount

Create an instance: `local no_frixion_money_moov_models_open_banking_account = client:NoFrixionMoneyMoovModelsOpenBankingAccount(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_balance` | `table` |  |
| `account_identification` | `table` |  |
| `account_name` | `table` |  |
| `account_type` | `string` |  |
| `balance` | `number` |  |
| `consolidated_account_information` | `table` |  |
| `currency` | `string` |  |
| `description` | `string` |  |
| `detail` | `string` |  |
| `id` | `string` |  |
| `nickname` | `string` |  |
| `type` | `string` |  |
| `usage_type` | `string` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_open_banking_account, err = client:NoFrixionMoneyMoovModelsOpenBankingAccount():load({ id = "no_frixion_money_moov_models_open_banking_account_id" })
```


### NoFrixionMoneyMoovModelsOpenBankingConsent

Create an instance: `local no_frixion_money_moov_models_open_banking_consent = client:NoFrixionMoneyMoovModelsOpenBankingConsent(nil)`

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
| `authorisation_url` | `string` |  |
| `callback_url` | `string` |  |
| `consent_id` | `string` |  |
| `email_address` | `string` |  |
| `expiry_date` | `string` |  |
| `failure_callback_url` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `institution_id` | `string` |  |
| `is_connected_account` | `boolean` |  |
| `is_enabled` | `boolean` |  |
| `merchant_id` | `string` |  |
| `provider` | `string` |  |
| `success_web_hook_url` | `string` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_open_banking_consent, err = client:NoFrixionMoneyMoovModelsOpenBankingConsent():load({ id = "no_frixion_money_moov_models_open_banking_consent_id" })
```

#### Example: List

```lua
local no_frixion_money_moov_models_open_banking_consents, err = client:NoFrixionMoneyMoovModelsOpenBankingConsent():list()
```

#### Example: Create

```lua
local no_frixion_money_moov_models_open_banking_consent, err = client:NoFrixionMoneyMoovModelsOpenBankingConsent():create({
})
```


### NoFrixionMoneyMoovModelsOpenBankingTransaction

Create an instance: `local no_frixion_money_moov_models_open_banking_transaction = client:NoFrixionMoneyMoovModelsOpenBankingTransaction(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address_detail` | `table` |  |
| `amount` | `number` |  |
| `balance` | `table` |  |
| `booking_date_time` | `string` |  |
| `charge_detail` | `table` |  |
| `currency` | `string` |  |
| `currency_exchange` | `table` |  |
| `date` | `string` |  |
| `description` | `string` |  |
| `enrichment` | `table` |  |
| `gross_amount` | `table` |  |
| `id` | `string` |  |
| `iso_bank_transaction_code` | `table` |  |
| `merchant` | `table` |  |
| `payee_detail` | `table` |  |
| `payer_detail` | `table` |  |
| `proprietary_bank_transaction_code` | `table` |  |
| `reference` | `string` |  |
| `statement_reference` | `table` |  |
| `status` | `string` |  |
| `supplementary_data` | `any` |  |
| `transaction_amount` | `table` |  |
| `transaction_information` | `table` |  |
| `transaction_mutability` | `string` |  |
| `value_date_time` | `string` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_open_banking_transactions, err = client:NoFrixionMoneyMoovModelsOpenBankingTransaction():list()
```


### NoFrixionMoneyMoovModelsPayment

Create an instance: `local no_frixion_money_moov_models_payment = client:NoFrixionMoneyMoovModelsPayment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `table` |  |
| `amount` | `number` |  |
| `amount_pending` | `number` |  |
| `amount_received` | `number` |  |
| `amount_refunded` | `number` |  |
| `auto_send_receipt` | `boolean` |  |
| `base_origin_url` | `string` |  |
| `callback_url` | `string` |  |
| `card_authorize_only` | `boolean` |  |
| `card_create_token` | `boolean` |  |
| `card_create_token_mode` | `string` |  |
| `card_ignore_cvn` | `boolean` |  |
| `card_no_payer_authentication` | `boolean` |  |
| `card_processor_merchant_id` | `string` |  |
| `card_stripe_payment_intent_id` | `string` |  |
| `card_stripe_payment_intent_secret` | `string` |  |
| `card_transmit_raw_detail` | `boolean` |  |
| `created_by_user` | `table` |  |
| `currency` | `string` |  |
| `custom_field` | `table` |  |
| `customer_email_address` | `string` |  |
| `customer_id` | `string` |  |
| `customer_name` | `string` |  |
| `description` | `string` |  |
| `destination_account` | `table` |  |
| `direct_debit_payment` | `table` |  |
| `due_date` | `string` |  |
| `event` | `table` |  |
| `failure_callback_url` | `string` |  |
| `field_display_setting` | `table` |  |
| `formatted_amount` | `string` |  |
| `hosted_pay_checkout_url` | `string` |  |
| `id` | `string` |  |
| `ignore_address_verification` | `boolean` |  |
| `inserted` | `string` |  |
| `inserted_sortable` | `string` |  |
| `is_archived` | `boolean` |  |
| `jwk` | `string` |  |
| `last_updated` | `string` |  |
| `lightning_invoice` | `string` |  |
| `lightning_invoice_expires_at` | `string` |  |
| `merchant_direct_debit_mandate_id` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `notification_email_address` | `string` |  |
| `notification_role_i_d` | `table` |  |
| `order_id` | `string` |  |
| `partial_payment_method` | `string` |  |
| `partial_payment_step` | `string` |  |
| `payment_attempt` | `table` |  |
| `payment_method` | `table` |  |
| `payment_processor` | `string` |  |
| `payrun_id` | `string` |  |
| `pisp_account_id` | `string` |  |
| `priority_bank_id` | `string` |  |
| `result` | `table` |  |
| `sandbox_settle_delay_in_second` | `number` |  |
| `shipping_address` | `table` |  |
| `shipping_address_city` | `string` |  |
| `shipping_address_country_code` | `string` |  |
| `shipping_address_county` | `string` |  |
| `shipping_address_line1` | `string` |  |
| `shipping_address_line2` | `string` |  |
| `shipping_address_post_code` | `string` |  |
| `shipping_email` | `string` |  |
| `shipping_first_name` | `string` |  |
| `shipping_last_name` | `string` |  |
| `shipping_phone` | `string` |  |
| `status` | `string` |  |
| `success_web_hook_url` | `string` |  |
| `tag` | `table` |  |
| `tag_id` | `table` |  |
| `title` | `string` |  |
| `tokenised_card` | `table` |  |
| `transaction` | `table` |  |
| `use_hosted_payment_page` | `boolean` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_payment, err = client:NoFrixionMoneyMoovModelsPayment():load({ id = "no_frixion_money_moov_models_payment_id" })
```

#### Example: Create

```lua
local no_frixion_money_moov_models_payment, err = client:NoFrixionMoneyMoovModelsPayment():create({
  created_by_user = {}, -- table
})
```


### NoFrixionMoneyMoovModelsPaymentAccountMinimalPage

Create an instance: `local no_frixion_money_moov_models_payment_account_minimal_page = client:NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `string` |  |
| `available_balance` | `number` |  |
| `balance` | `number` |  |
| `balance_minor_unit` | `number` |  |
| `currency` | `string` |  |
| `id` | `string` |  |
| `identifier` | `table` |  |
| `is_archived` | `boolean` |  |
| `is_connected_account` | `boolean` |  |
| `merchant_id` | `string` |  |
| `submitted_payouts_balance` | `number` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_payment_account_minimal_pages, err = client:NoFrixionMoneyMoovModelsPaymentAccountMinimalPage():list()
```


### NoFrixionMoneyMoovModelsPaymentAccountPage

Create an instance: `local no_frixion_money_moov_models_payment_account_page = client:NoFrixionMoneyMoovModelsPaymentAccountPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `string` |  |
| `account_supplier_name` | `string` |  |
| `available_balance` | `number` |  |
| `available_balance_minor_unit` | `number` |  |
| `balance` | `number` |  |
| `balance_minor_unit` | `number` |  |
| `bank_name` | `string` |  |
| `consent_id` | `string` |  |
| `created_by` | `table` |  |
| `created_by_display_name` | `string` |  |
| `currency` | `string` |  |
| `default_payment_rail` | `string` |  |
| `display_name` | `string` |  |
| `expiry_date` | `string` |  |
| `external_account_icon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `table` |  |
| `inserted` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_connected_account` | `boolean` |  |
| `is_default` | `boolean` |  |
| `is_trust_account` | `boolean` |  |
| `is_virtual` | `boolean` |  |
| `last_transaction` | `table` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `physical_account_id` | `string` |  |
| `rule` | `table` |  |
| `submitted_payouts_balance` | `number` |  |
| `submitted_payouts_balance_minor_unit` | `number` |  |
| `summary` | `string` |  |
| `supplier_sepa_instant_status` | `string` |  |
| `xero_bank_feed_connection_status` | `string` |  |
| `xero_bank_feed_last_synced_at` | `string` |  |
| `xero_bank_feed_sync_last_failed_at` | `string` |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` |  |
| `xero_bank_feed_sync_status` | `string` |  |
| `xero_unsynchronised_transactions_count` | `number` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_payment_account_pages, err = client:NoFrixionMoneyMoovModelsPaymentAccountPage():list()
```


### NoFrixionMoneyMoovModelsPaymentInitiation

Create an instance: `local no_frixion_money_moov_models_payment_initiation = client:NoFrixionMoneyMoovModelsPaymentInitiation(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `payment_initiation_id` | `string` |  |
| `payment_request_callback_url` | `string` |  |
| `payment_request_id` | `string` |  |
| `redirect_url` | `string` |  |
| `response_type` | `string` |  |
| `specific_error_message` | `string` |  |

#### Example: Create

```lua
local no_frixion_money_moov_models_payment_initiation, err = client:NoFrixionMoneyMoovModelsPaymentInitiation():create({
  paymentrequest_id = "example_paymentrequest_id", -- string
})
```


### NoFrixionMoneyMoovModelsPaymentRequestEvent

Create an instance: `local no_frixion_money_moov_models_payment_request_event = client:NoFrixionMoneyMoovModelsPaymentRequestEvent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `number` |  |
| `apple_pay_transaction_id` | `string` |  |
| `card_authorization_response_id` | `string` |  |
| `card_expiry_month` | `number` |  |
| `card_expiry_year` | `number` |  |
| `card_issuer` | `string` |  |
| `card_issuer_country` | `string` |  |
| `card_last_four_digit` | `string` |  |
| `card_request_id` | `string` |  |
| `card_scheme` | `string` |  |
| `card_token_customer_id` | `string` |  |
| `card_transaction_id` | `string` |  |
| `currency` | `string` |  |
| `direct_debit_payment_id` | `string` |  |
| `direct_debit_payment_reference` | `string` |  |
| `drirect_debit_mandate_id` | `string` |  |
| `error_message` | `string` |  |
| `error_reason` | `string` |  |
| `event_type` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lightning_invoice` | `string` |  |
| `lightning_r_hash` | `string` |  |
| `origin_url` | `string` |  |
| `payment_method_type` | `string` |  |
| `payment_processor_name` | `string` |  |
| `payment_request_id` | `string` |  |
| `pisp_bank_status` | `string` |  |
| `pisp_payment_initiation_id` | `string` |  |
| `pisp_payment_institution_name` | `string` |  |
| `pisp_payment_service_provider_id` | `string` |  |
| `pisp_redirect_url` | `string` |  |
| `reconciled_transaction_id` | `string` |  |
| `refund_payout_id` | `string` |  |
| `status` | `string` |  |
| `wallet_name` | `string` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_payment_request_events, err = client:NoFrixionMoneyMoovModelsPaymentRequestEvent():list()
```


### NoFrixionMoneyMoovModelsPaymentRequestMetric

Create an instance: `local no_frixion_money_moov_models_payment_request_metric = client:NoFrixionMoneyMoovModelsPaymentRequestMetric(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `all` | `number` |  |
| `authorized` | `number` |  |
| `paid` | `number` |  |
| `partially_paid` | `number` |  |
| `total_amounts_by_currency` | `table` |  |
| `unpaid` | `number` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_payment_request_metric, err = client:NoFrixionMoneyMoovModelsPaymentRequestMetric():load()
```


### NoFrixionMoneyMoovModelsPaymentRequestMinimal

Create an instance: `local no_frixion_money_moov_models_payment_request_minimal = client:NoFrixionMoneyMoovModelsPaymentRequestMinimal(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `number` |  |
| `amount_pending` | `number` |  |
| `amount_received` | `number` |  |
| `amount_refunded` | `number` |  |
| `callback_url` | `string` |  |
| `card_stripe_payment_intent_secret` | `string` |  |
| `country_code` | `string` |  |
| `currency` | `string` |  |
| `custom_fields_to_display` | `table` |  |
| `description` | `string` |  |
| `due_date` | `string` |  |
| `field_display_setting` | `table` |  |
| `google_pay_merchant_id` | `string` |  |
| `id` | `string` |  |
| `jwk` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_logo_url_png` | `string` |  |
| `merchant_logo_url_svg` | `string` |  |
| `merchant_name` | `string` |  |
| `merchant_short_name` | `string` |  |
| `partial_payment_method` | `string` |  |
| `payment_attempt` | `table` |  |
| `payment_methods_list` | `table` |  |
| `payment_processor` | `string` |  |
| `payment_processor_key` | `string` |  |
| `pisp_error` | `string` |  |
| `priority_bank_id` | `string` |  |
| `status` | `string` |  |
| `stripe_account_id` | `string` |  |
| `title` | `string` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_payment_request_minimals, err = client:NoFrixionMoneyMoovModelsPaymentRequestMinimal():list()
```


### NoFrixionMoneyMoovModelsPaymentRequestResult

Create an instance: `local no_frixion_money_moov_models_payment_request_result = client:NoFrixionMoneyMoovModelsPaymentRequestResult(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `number` |  |
| `amount_pending` | `number` |  |
| `amount_received` | `number` |  |
| `amount_refunded` | `number` |  |
| `currency` | `string` |  |
| `customer_id` | `string` |  |
| `payment` | `table` |  |
| `payment_request_id` | `string` |  |
| `pisp_authorization` | `table` |  |
| `requested_amount` | `number` |  |
| `result` | `string` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_payment_request_results, err = client:NoFrixionMoneyMoovModelsPaymentRequestResult():list()
```


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment

Create an instance: `local no_frixion_money_moov_models_payment_requests_merchant_payment = client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `description` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `template` | `table` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_payment_requests_merchant_payments, err = client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment():list()
```


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2

Create an instance: `local no_frixion_money_moov_models_payment_requests_merchant_payment2 = client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `description` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `template` | `table` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_payment_requests_merchant_payment2, err = client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2():load({ paymentrequest_id = "paymentrequest_id", template_id = "template_id" })
```


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3

Create an instance: `local no_frixion_money_moov_models_payment_requests_merchant_payment3 = client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `description` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `template` | `table` |  |


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4

Create an instance: `local no_frixion_money_moov_models_payment_requests_merchant_payment4 = client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |


### NoFrixionMoneyMoovModelsPayoutKeysetPage

Create an instance: `local no_frixion_money_moov_models_payout_keyset_page = client:NoFrixionMoneyMoovModelsPayoutKeysetPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `amount` | `number` |  |
| `amount_minor_unit` | `number` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `table` |  |
| `authorisation` | `table` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `table` |  |
| `can_authorise` | `boolean` |  |
| `can_process` | `boolean` |  |
| `can_update` | `boolean` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `table` |  |
| `document` | `table` |  |
| `event` | `table` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `number` |  |
| `fx_destination_amount_minor_unit` | `number` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `number` |  |
| `fx_use_destination_amount` | `boolean` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_failed` | `boolean` |  |
| `is_settled` | `boolean` |  |
| `is_submitted` | `boolean` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `rule` | `table` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `boolean` |  |
| `source_account_available_balance` | `number` |  |
| `source_account_available_balance_minor_unit` | `number` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `table` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `table` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `number` |  |
| `transacted_fx_amount` | `number` |  |
| `transacted_fx_rate` | `number` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_payout_keyset_pages, err = client:NoFrixionMoneyMoovModelsPayoutKeysetPage():list()
```


### NoFrixionMoneyMoovModelsPayoutMetric

Create an instance: `local no_frixion_money_moov_models_payout_metric = client:NoFrixionMoneyMoovModelsPayoutMetric(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `all` | `number` |  |
| `failed` | `number` |  |
| `in_progress` | `number` |  |
| `paid` | `number` |  |
| `pending_approval` | `number` |  |
| `scheduled` | `number` |  |
| `total_amounts_by_currency` | `table` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_payout_metric, err = client:NoFrixionMoneyMoovModelsPayoutMetric():load()
```


### NoFrixionMoneyMoovModelsPayoutsPayoutsCreate

Create an instance: `local no_frixion_money_moov_models_payouts_payouts_create = client:NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_payout` | `table` |  |
| `payout` | `table` |  |

#### Example: Create

```lua
local no_frixion_money_moov_models_payouts_payouts_create, err = client:NoFrixionMoneyMoovModelsPayoutsPayoutsCreate():create({
})
```


### NoFrixionMoneyMoovModelsPayrun

Create an instance: `local no_frixion_money_moov_models_payrun = client:NoFrixionMoneyMoovModelsPayrun(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation` | `table` |  |
| `authorisation_date` | `string` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `batch_payout_id` | `string` |  |
| `can_authorise` | `boolean` |  |
| `can_delete` | `boolean` |  |
| `can_edit` | `boolean` |  |
| `event` | `table` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice` | `table` |  |
| `invoices_minimal` | `table` |  |
| `is_archived` | `boolean` |  |
| `last_updated` | `string` |  |
| `last_updated_by` | `table` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `payment` | `table` |  |
| `payout` | `table` |  |
| `payouts_count` | `number` |  |
| `reason` | `string` |  |
| `schedule_date` | `string` |  |
| `scheduled_date` | `string` |  |
| `source_account` | `table` |  |
| `status` | `string` |  |
| `total_eur` | `number` |  |
| `total_gbp` | `number` |  |
| `total_usd` | `number` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_payrun, err = client:NoFrixionMoneyMoovModelsPayrun():load({ id = "no_frixion_money_moov_models_payrun_id" })
```

#### Example: Create

```lua
local no_frixion_money_moov_models_payrun, err = client:NoFrixionMoneyMoovModelsPayrun():create({
  id = "example_id", -- string
})
```


### NoFrixionMoneyMoovModelsReportResult

Create an instance: `local no_frixion_money_moov_models_report_result = client:NoFrixionMoneyMoovModelsReportResult(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `content` | `string` |  |
| `content_type` | `string` |  |
| `last_completed_at` | `string` |  |
| `merchant_id` | `string` |  |
| `report_name` | `string` |  |
| `report_type` | `string` |  |
| `statement_number` | `number` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_report_result, err = client:NoFrixionMoneyMoovModelsReportResult():load({ id = 1, report_id = "report_id" })
```


### NoFrixionMoneyMoovModelsRule

Create an instance: `local no_frixion_money_moov_models_rule = client:NoFrixionMoneyMoovModelsRule(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account` | `table` |  |
| `account_id` | `string` |  |
| `approve_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `table` |  |
| `authorisation` | `table` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `can_authorise` | `boolean` |  |
| `created_by` | `table` |  |
| `description` | `string` |  |
| `end_at` | `string` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_disabled` | `boolean` |  |
| `last_executed_at` | `string` |  |
| `last_run_at_transaction_date` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `on_approved_web_hook_url` | `string` |  |
| `on_execution_error_web_hook_url` | `string` |  |
| `on_execution_success_web_hook_url` | `string` |  |
| `start_at` | `string` |  |
| `status` | `string` |  |
| `sweep_action` | `table` |  |
| `time_zone_id` | `string` |  |
| `trigger_cron_expression` | `string` |  |
| `trigger_on_pay_in` | `boolean` |  |
| `user_id` | `string` |  |
| `web_hook_secret` | `string` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_rule, err = client:NoFrixionMoneyMoovModelsRule():load({ id = "no_frixion_money_moov_models_rule_id" })
```

#### Example: Create

```lua
local no_frixion_money_moov_models_rule, err = client:NoFrixionMoneyMoovModelsRule():create({
  created_by = {}, -- table
  nonce = "example_nonce", -- string
})
```


### NoFrixionMoneyMoovModelsTransaction

Create an instance: `local no_frixion_money_moov_models_transaction = client:NoFrixionMoneyMoovModelsTransaction(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `account_name` | `string` |  |
| `account_sequence_number` | `number` |  |
| `amount` | `number` |  |
| `amount_minor_unit` | `number` |  |
| `balance` | `number` |  |
| `balance_minor_unit` | `number` |  |
| `counterparty` | `table` |  |
| `counterparty_summary` | `string` |  |
| `currency` | `string` |  |
| `description` | `string` |  |
| `fx_amount` | `number` |  |
| `fx_currency` | `string` |  |
| `fx_rate` | `number` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `merchant_id` | `string` |  |
| `payment_request_custom_field` | `table` |  |
| `payment_request_id` | `string` |  |
| `payout_id` | `string` |  |
| `raw_reference` | `string` |  |
| `rule_id` | `string` |  |
| `tag` | `table` |  |
| `their_reference` | `string` |  |
| `transaction_date` | `string` |  |
| `type` | `string` |  |
| `virtual_iban` | `string` |  |
| `your_reference` | `string` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_transaction, err = client:NoFrixionMoneyMoovModelsTransaction():load({ id = "no_frixion_money_moov_models_transaction_id" })
```


### NoFrixionMoneyMoovModelsTransactionPage

Create an instance: `local no_frixion_money_moov_models_transaction_page = client:NoFrixionMoneyMoovModelsTransactionPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `account_name` | `string` |  |
| `account_sequence_number` | `number` |  |
| `amount` | `number` |  |
| `amount_minor_unit` | `number` |  |
| `balance` | `number` |  |
| `balance_minor_unit` | `number` |  |
| `content` | `table` |  |
| `counterparty` | `table` |  |
| `counterparty_summary` | `string` |  |
| `currency` | `string` |  |
| `description` | `string` |  |
| `fx_amount` | `number` |  |
| `fx_currency` | `string` |  |
| `fx_rate` | `number` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `merchant_id` | `string` |  |
| `page_number` | `number` |  |
| `page_size` | `number` |  |
| `payment_request_custom_field` | `table` |  |
| `payment_request_id` | `string` |  |
| `payout_id` | `string` |  |
| `raw_reference` | `string` |  |
| `rule_id` | `string` |  |
| `tag` | `table` |  |
| `their_reference` | `string` |  |
| `total_page` | `number` |  |
| `total_size` | `number` |  |
| `transaction_date` | `string` |  |
| `type` | `string` |  |
| `virtual_iban` | `string` |  |
| `your_reference` | `string` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_transaction_page, err = client:NoFrixionMoneyMoovModelsTransactionPage():load({ account_id = "account_id" })
```

#### Example: List

```lua
local no_frixion_money_moov_models_transaction_pages, err = client:NoFrixionMoneyMoovModelsTransactionPage():list()
```


### NoFrixionMoneyMoovModelsUserInvite

Create an instance: `local no_frixion_money_moov_models_user_invite = client:NoFrixionMoneyMoovModelsUserInvite(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_status` | `table` |  |
| `id` | `string` |  |
| `initial_role_id` | `string` |  |
| `invitee_email_address` | `string` |  |
| `invitee_first_name` | `string` |  |
| `invitee_last_name` | `string` |  |
| `inviter_email_address` | `string` |  |
| `inviter_first_name` | `string` |  |
| `inviter_last_name` | `string` |  |
| `is_authorised` | `boolean` |  |
| `is_invitee_registered` | `boolean` |  |
| `last_invited` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `message` | `string` |  |
| `registration_url` | `string` |  |
| `send_invite_email` | `boolean` |  |
| `status` | `string` |  |
| `user` | `table` |  |
| `user_id` | `string` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_user_invite, err = client:NoFrixionMoneyMoovModelsUserInvite():load({ id = "no_frixion_money_moov_models_user_invite_id" })
```

#### Example: Create

```lua
local no_frixion_money_moov_models_user_invite, err = client:NoFrixionMoneyMoovModelsUserInvite():create({
  user = {}, -- table
})
```


### NoFrixionMoneyMoovModelsUserInvitePage

Create an instance: `local no_frixion_money_moov_models_user_invite_page = client:NoFrixionMoneyMoovModelsUserInvitePage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_status` | `table` |  |
| `id` | `string` |  |
| `initial_role_id` | `string` |  |
| `invitee_email_address` | `string` |  |
| `invitee_first_name` | `string` |  |
| `invitee_last_name` | `string` |  |
| `inviter_email_address` | `string` |  |
| `inviter_first_name` | `string` |  |
| `inviter_last_name` | `string` |  |
| `is_authorised` | `boolean` |  |
| `is_invitee_registered` | `boolean` |  |
| `last_invited` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `message` | `string` |  |
| `registration_url` | `string` |  |
| `status` | `string` |  |
| `user` | `table` |  |
| `user_id` | `string` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_user_invite_pages, err = client:NoFrixionMoneyMoovModelsUserInvitePage():list()
```


### NoFrixionMoneyMoovModelsUserPage

Create an instance: `local no_frixion_money_moov_models_user_page = client:NoFrixionMoneyMoovModelsUserPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client_session_timeout` | `table` |  |
| `email_address` | `string` |  |
| `first_name` | `string` |  |
| `id` | `string` |  |
| `last_name` | `string` |  |
| `passkey_added` | `boolean` |  |
| `permission` | `table` |  |
| `roles_with_scope` | `table` |  |
| `two_factor_enabled` | `boolean` |  |

#### Example: List

```lua
local no_frixion_money_moov_models_user_pages, err = client:NoFrixionMoneyMoovModelsUserPage():list()
```


### NoFrixionMoneyMoovModelsWebhook

Create an instance: `local no_frixion_money_moov_models_webhook = client:NoFrixionMoneyMoovModelsWebhook(nil)`

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
| `destination_url` | `string` |  |
| `email_address` | `string` |  |
| `failed_notification_email_address` | `string` |  |
| `id` | `string` |  |
| `is_active` | `boolean` |  |
| `merchant_id` | `string` |  |
| `notification_method` | `string` |  |
| `resource_type` | `table` |  |
| `retry` | `boolean` |  |
| `secret` | `string` |  |
| `version` | `number` |  |

#### Example: Load

```lua
local no_frixion_money_moov_models_webhook, err = client:NoFrixionMoneyMoovModelsWebhook():load({ id = "no_frixion_money_moov_models_webhook_id" })
```

#### Example: List

```lua
local no_frixion_money_moov_models_webhooks, err = client:NoFrixionMoneyMoovModelsWebhook():list()
```

#### Example: Create

```lua
local no_frixion_money_moov_models_webhook, err = client:NoFrixionMoneyMoovModelsWebhook():create({
})
```


### OpenBanking

Create an instance: `local open_banking = client:OpenBanking(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```lua
local open_banking, err = client:OpenBanking():create({
  account_id = "example_account_id", -- string
})
```


### Payeeverification

Create an instance: `local payeeverification = client:Payeeverification(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `string` |  |
| `account_number` | `string` |  |
| `iban` | `string` |  |
| `payee_verified_account_name` | `string` |  |
| `result` | `string` |  |
| `secondary_identification` | `string` |  |
| `sort_code` | `string` |  |

#### Example: Create

```lua
local payeeverification, err = client:Payeeverification():create({
  account_name = "example_account_name", -- string
  iban = "example_iban", -- string
})
```


### PaymentRequest

Create an instance: `local payment_request = client:PaymentRequest(nil)`

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
| `amount` | `number` |  |
| `do_simulate_settlement_failure` | `boolean` |  |
| `error_description` | `string` |  |
| `institution` | `string` |  |
| `payment_initiation_id` | `string` |  |

#### Example: Load

```lua
local payment_request, err = client:PaymentRequest():load()
```

#### Example: Create

```lua
local payment_request, err = client:PaymentRequest():create({
  paymentrequest_id = "example_paymentrequest_id", -- string
})
```


### Payout

Create an instance: `local payout = client:Payout(nil)`

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
| `account_id` | `string` |  |
| `allow_incomplete` | `boolean` |  |
| `amount` | `number` |  |
| `amount_minor_unit` | `number` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `table` |  |
| `authorisation` | `table` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `table` |  |
| `beneficiary_id` | `string` |  |
| `can_authorise` | `boolean` |  |
| `can_process` | `boolean` |  |
| `can_update` | `boolean` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `table` |  |
| `document` | `table` |  |
| `event` | `table` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `number` |  |
| `fx_destination_amount_minor_unit` | `number` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `number` |  |
| `fx_use_destination_amount` | `boolean` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_failed` | `boolean` |  |
| `is_settled` | `boolean` |  |
| `is_submitted` | `boolean` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `rule` | `table` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `boolean` |  |
| `source_account_available_balance` | `number` |  |
| `source_account_available_balance_minor_unit` | `number` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `table` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `table` |  |
| `tag_id` | `table` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `number` |  |
| `transacted_fx_amount` | `number` |  |
| `transacted_fx_rate` | `number` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |

#### Example: Load

```lua
local payout, err = client:Payout():load({ id = "payout_id" })
```

#### Example: Create

```lua
local payout, err = client:Payout():create({
})
```


### Payrun

Create an instance: `local payrun = client:Payrun(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |
| `note` | `string` |  |
| `scheduled_date` | `string` |  |

#### Example: Create

```lua
local payrun, err = client:Payrun():create({
  id = "example_id", -- string
})
```


### Reject

Create an instance: `local reject = client:Reject(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `amount` | `number` |  |
| `amount_minor_unit` | `number` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `table` |  |
| `authorisation` | `table` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `table` |  |
| `can_authorise` | `boolean` |  |
| `can_process` | `boolean` |  |
| `can_update` | `boolean` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `table` |  |
| `document` | `table` |  |
| `event` | `table` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `number` |  |
| `fx_destination_amount_minor_unit` | `number` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `number` |  |
| `fx_use_destination_amount` | `boolean` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_failed` | `boolean` |  |
| `is_settled` | `boolean` |  |
| `is_submitted` | `boolean` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `reason` | `string` |  |
| `rule` | `table` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `boolean` |  |
| `source_account_available_balance` | `number` |  |
| `source_account_available_balance_minor_unit` | `number` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `table` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `table` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `number` |  |
| `transacted_fx_amount` | `number` |  |
| `transacted_fx_rate` | `number` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |


### Report

Create an instance: `local report = client:Report(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |


### Rule

Create an instance: `local rule = client:Rule(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |


### Send

Create an instance: `local send = client:Send(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `allow_incomplete` | `boolean` |  |
| `amount` | `number` |  |
| `amount_minor_unit` | `number` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `table` |  |
| `authorisation` | `table` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `table` |  |
| `beneficiary_id` | `string` |  |
| `can_authorise` | `boolean` |  |
| `can_process` | `boolean` |  |
| `can_update` | `boolean` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `table` |  |
| `document` | `table` |  |
| `event` | `table` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `number` |  |
| `fx_destination_amount_minor_unit` | `number` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `number` |  |
| `fx_use_destination_amount` | `boolean` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_failed` | `boolean` |  |
| `is_settled` | `boolean` |  |
| `is_submitted` | `boolean` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `rule` | `table` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `boolean` |  |
| `source_account_available_balance` | `number` |  |
| `source_account_available_balance_minor_unit` | `number` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `table` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `table` |  |
| `tag_id` | `table` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `number` |  |
| `transacted_fx_amount` | `number` |  |
| `transacted_fx_rate` | `number` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |

#### Example: Create

```lua
local send, err = client:Send():create({
  beneficiary = {}, -- table
  source_account_identifier = {}, -- table
})
```


### Sendbeneficiary

Create an instance: `local sendbeneficiary = client:Sendbeneficiary(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `allow_incomplete` | `boolean` |  |
| `amount` | `number` |  |
| `amount_minor_unit` | `number` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `table` |  |
| `authorisation` | `table` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `table` |  |
| `beneficiary_id` | `string` |  |
| `can_authorise` | `boolean` |  |
| `can_process` | `boolean` |  |
| `can_update` | `boolean` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `table` |  |
| `document` | `table` |  |
| `event` | `table` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `number` |  |
| `fx_destination_amount_minor_unit` | `number` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `number` |  |
| `fx_use_destination_amount` | `boolean` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_failed` | `boolean` |  |
| `is_settled` | `boolean` |  |
| `is_submitted` | `boolean` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `rule` | `table` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `boolean` |  |
| `source_account_available_balance` | `number` |  |
| `source_account_available_balance_minor_unit` | `number` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `table` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `table` |  |
| `tag_id` | `table` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `number` |  |
| `transacted_fx_amount` | `number` |  |
| `transacted_fx_rate` | `number` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |

#### Example: Create

```lua
local sendbeneficiary, err = client:Sendbeneficiary():create({
  beneficiary = {}, -- table
  source_account_identifier = {}, -- table
})
```


### Tag

Create an instance: `local tag = client:Tag(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `colour_hex` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |

#### Example: List

```lua
local tags, err = client:Tag():list()
```

#### Example: Create

```lua
local tag, err = client:Tag():create({
  merchant_id = "example_merchant_id", -- string
})
```


### Token

Create an instance: `local token = client:Token(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```lua
local token, err = client:Token():create({
  id = "example_id", -- string
})
```


### Transaction

Create an instance: `local transaction = client:Transaction(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Example: Load

```lua
local transaction, err = client:Transaction():load({ sequence_number = 1, transaction_id = "transaction_id" })
```

#### Example: Create

```lua
local transaction, err = client:Transaction():create({
  id = "example_id", -- string
})
```


### User

Create an instance: `local user = client:User(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client_session_timeout` | `table` |  |
| `email_address` | `string` |  |
| `first_name` | `string` |  |
| `id` | `string` |  |
| `last_name` | `string` |  |
| `passkey_added` | `boolean` |  |
| `permission` | `table` |  |
| `profile` | `string` |  |
| `roles_with_scope` | `table` |  |
| `two_factor_enabled` | `boolean` |  |
| `user_invite_id` | `string` |  |

#### Example: List

```lua
local users, err = client:User():list()
```


### UserInvite

Create an instance: `local user_invite = client:UserInvite(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Example: Create

```lua
local user_invite, err = client:UserInvite():create({
  id = "example_id", -- string
})
```


### Virtual

Create an instance: `local virtual = client:Virtual(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `string` |  |
| `account_supplier_name` | `string` |  |
| `available_balance` | `number` |  |
| `available_balance_minor_unit` | `number` |  |
| `balance` | `number` |  |
| `balance_minor_unit` | `number` |  |
| `bank_name` | `string` |  |
| `consent_id` | `string` |  |
| `created_by` | `table` |  |
| `created_by_display_name` | `string` |  |
| `currency` | `string` |  |
| `default_payment_rail` | `string` |  |
| `display_name` | `string` |  |
| `expiry_date` | `string` |  |
| `external_account_icon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `table` |  |
| `inserted` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_connected_account` | `boolean` |  |
| `is_default` | `boolean` |  |
| `is_trust_account` | `boolean` |  |
| `is_virtual` | `boolean` |  |
| `last_transaction` | `table` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `name` | `string` |  |
| `physical_account_id` | `string` |  |
| `rule` | `table` |  |
| `submitted_payouts_balance` | `number` |  |
| `submitted_payouts_balance_minor_unit` | `number` |  |
| `summary` | `string` |  |
| `supplier_sepa_instant_status` | `string` |  |
| `xero_bank_feed_connection_status` | `string` |  |
| `xero_bank_feed_last_synced_at` | `string` |  |
| `xero_bank_feed_sync_last_failed_at` | `string` |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` |  |
| `xero_bank_feed_sync_status` | `string` |  |
| `xero_unsynchronised_transactions_count` | `number` |  |

#### Example: Create

```lua
local virtual, err = client:Virtual():create({
  account_id = "example_account_id", -- string
})
```


### Webhook

Create an instance: `local webhook = client:Webhook(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |


### Whoami

Create an instance: `local whoami = client:Whoami(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client_session_timeout` | `table` |  |
| `email_address` | `string` |  |
| `first_name` | `string` |  |
| `id` | `string` |  |
| `last_name` | `string` |  |
| `passkey_added` | `boolean` |  |
| `permission` | `table` |  |
| `roles_with_scope` | `table` |  |
| `two_factor_enabled` | `boolean` |  |

#### Example: List

```lua
local whoamis, err = client:Whoami():list()
```


### Whoamitrustedapp

Create an instance: `local whoamitrustedapp = client:Whoamitrustedapp(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client_session_timeout` | `table` |  |
| `email_address` | `string` |  |
| `first_name` | `string` |  |
| `id` | `string` |  |
| `last_name` | `string` |  |
| `passkey_added` | `boolean` |  |
| `permission` | `table` |  |
| `roles_with_scope` | `table` |  |
| `two_factor_enabled` | `boolean` |  |

#### Example: List

```lua
local whoamitrustedapps, err = client:Whoamitrustedapp():list()
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

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── nofrixion_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`nofrixion_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```lua
local account = client:Account()
account:list()

-- account:data_get() now returns the account data from the last list
-- account:match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
