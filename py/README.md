# Nofrixion Python SDK



The Python SDK for the Nofrixion API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Account()` — each
carrying a small, uniform set of operations (`list`, `load`, `create`, `update`, `remove`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/nofrixion-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from nofrixion_sdk import NofrixionSDK

client = NofrixionSDK({
    "apikey": os.environ.get("NOFRIXION_APIKEY"),
})
```

### 2. List account records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    accounts = client.Account().list()
    for account in accounts:
        print(account)
except Exception as err:
    print(f"list failed: {err}")
```

### 3. Load a merchant

Merchant is nested under merchant, so provide the `merchant_id`.
`load()` returns the bare record (a `dict`) and raises on error.

```python
try:
    merchant = client.Merchant().load({"merchant_id": "example_merchant_id"})
    print(merchant)
except Exception as err:
    print(f"load failed: {err}")
```

### 4. Create, update, and remove

```python
# Create — returns the bare created record (a dict)
created = client.Account().create({"account_id": "example_account_id", "currency": "example_currency"})

# Update — the created record's id is a plain dict key
client.Account().update({"id": created["id"], "account_id": "example_account_id", "amount": 1})

# Remove
client.Account().remove({"id": created["id"]})
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    accounts = client.Account().list()
    print(accounts)
except Exception as err:
    print(f"list failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = NofrixionSDK.test()

# Entity ops return the bare record and raise on error.
account = client.Account().list()
# account contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = NofrixionSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### NofrixionSDK

```python
from nofrixion_sdk import NofrixionSDK

client = NofrixionSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = NofrixionSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### NofrixionSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
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
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the bare result data (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

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

Create an instance: `account = client.Account()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `str` |  |
| `account_name` | `str` |  |
| `account_supplier_name` | `str` |  |
| `account_type` | `str` |  |
| `available_balance` | `float` |  |
| `available_balance_minor_unit` | `int` |  |
| `balance` | `float` |  |
| `balance_minor_unit` | `int` |  |
| `bank_name` | `str` |  |
| `consent_id` | `str` |  |
| `created_by` | `dict` |  |
| `created_by_display_name` | `str` |  |
| `currency` | `str` |  |
| `default_payment_rail` | `str` |  |
| `display_name` | `str` |  |
| `expiry_date` | `str` |  |
| `external_account_icon` | `str` |  |
| `format` | `str` |  |
| `from_date` | `str` |  |
| `id` | `str` |  |
| `identifier` | `dict` |  |
| `inserted` | `str` |  |
| `is_archived` | `bool` |  |
| `is_connected_account` | `bool` |  |
| `is_default` | `bool` |  |
| `is_trust_account` | `bool` |  |
| `is_virtual` | `bool` |  |
| `last_transaction` | `dict` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `merchant_name` | `str` |  |
| `physical_account_id` | `str` |  |
| `role_i_d` | `list` |  |
| `rule` | `list` |  |
| `submitted_payouts_balance` | `float` |  |
| `submitted_payouts_balance_minor_unit` | `int` |  |
| `summary` | `str` |  |
| `supplier_physical_account_id` | `str` |  |
| `supplier_sepa_instant_status` | `str` |  |
| `to_date` | `str` |  |
| `xero_bank_feed_connection_status` | `str` |  |
| `xero_bank_feed_last_synced_at` | `str` |  |
| `xero_bank_feed_sync_last_failed_at` | `str` |  |
| `xero_bank_feed_sync_last_failure_reason` | `str` |  |
| `xero_bank_feed_sync_status` | `str` |  |
| `xero_unsynchronised_transactions_count` | `int` |  |

#### Example: Load

```python
account = client.Account().load({"id": "account_id"})
```

#### Example: List

```python
accounts = client.Account().list()
```

#### Example: Create

```python
account = client.Account().create({
})
```


### Beneficiary

Create an instance: `beneficiary = client.Beneficiary()`

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
| `approval_callback_url` | `str` |  |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `beneficiary_event` | `list` |  |
| `can_authorise` | `bool` |  |
| `can_update` | `bool` |  |
| `created_by` | `dict` |  |
| `created_by_email_address` | `str` |  |
| `currency` | `str` |  |
| `destination` | `dict` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `is_enabled` | `bool` |  |
| `last_authorised` | `str` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `name` | `str` |  |
| `nonce` | `str` |  |
| `source_account` | `list` |  |
| `source_account_i_d` | `list` |  |
| `their_reference` | `str` |  |

#### Example: Load

```python
beneficiary = client.Beneficiary().load({"id": "beneficiary_id"})
```

#### Example: Create

```python
beneficiary = client.Beneficiary().create({
})
```


### Cancel

Create an instance: `cancel = client.Cancel()`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `str` |  |
| `amount` | `float` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `str` |  |
| `approver_id` | `str` |  |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `str` |  |
| `beneficiary` | `dict` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `str` |  |
| `created_by` | `str` |  |
| `created_by_email_address` | `str` |  |
| `currency` | `str` |  |
| `current_user_id` | `str` |  |
| `description` | `str` |  |
| `destination` | `dict` |  |
| `document` | `list` |  |
| `event` | `list` |  |
| `formatted_amount` | `str` |  |
| `formatted_fx_destination_amount` | `str` |  |
| `formatted_schedule` | `str` |  |
| `formatted_schedule_day_only` | `str` |  |
| `formatted_source_account_available_balance` | `str` |  |
| `fx_destination_amount` | `float` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `str` |  |
| `fx_quote_expires_at` | `str` |  |
| `fx_quote_id` | `str` |  |
| `fx_rate` | `float` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `invoice_id` | `str` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `merchant_token_description` | `str` |  |
| `nonce` | `str` |  |
| `payment_processor` | `str` |  |
| `payment_rail` | `str` |  |
| `payrun_id` | `str` |  |
| `payrun_name` | `str` |  |
| `rule` | `dict` |  |
| `schedule_date` | `str` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `str` |  |
| `source_account_currency` | `str` |  |
| `source_account_iban` | `str` |  |
| `source_account_identifier` | `dict` |  |
| `source_account_name` | `str` |  |
| `source_account_number` | `str` |  |
| `source_account_sortcode` | `str` |  |
| `status` | `str` |  |
| `tag` | `list` |  |
| `their_reference` | `str` |  |
| `topup_payrun_id` | `str` |  |
| `transacted_amount` | `float` |  |
| `transacted_fx_amount` | `float` |  |
| `transacted_fx_rate` | `float` |  |
| `type` | `str` |  |
| `user_id` | `str` |  |
| `your_reference` | `str` |  |


### Disable

Create an instance: `disable = client.Disable()`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_callback_url` | `str` |  |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `beneficiary_event` | `list` |  |
| `can_authorise` | `bool` |  |
| `can_update` | `bool` |  |
| `created_by` | `dict` |  |
| `created_by_email_address` | `str` |  |
| `currency` | `str` |  |
| `destination` | `dict` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `is_enabled` | `bool` |  |
| `last_authorised` | `str` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `name` | `str` |  |
| `nonce` | `str` |  |
| `source_account` | `list` |  |
| `their_reference` | `str` |  |


### Enable

Create an instance: `enable = client.Enable()`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_callback_url` | `str` |  |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `beneficiary_event` | `list` |  |
| `can_authorise` | `bool` |  |
| `can_update` | `bool` |  |
| `created_by` | `dict` |  |
| `created_by_email_address` | `str` |  |
| `currency` | `str` |  |
| `destination` | `dict` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `is_enabled` | `bool` |  |
| `last_authorised` | `str` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `name` | `str` |  |
| `nonce` | `str` |  |
| `source_account` | `list` |  |
| `their_reference` | `str` |  |


### Merchant

Create an instance: `merchant = client.Merchant()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `reason` | `str` |  |

#### Example: Load

```python
merchant = client.Merchant().load({"merchant_id": "merchant_id"})
```


### Metadata

Create an instance: `metadata = client.Metadata()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```python
metadata = client.Metadata().load()
```


### NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage

Create an instance: `no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page = client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approved_at` | `str` |  |
| `currency` | `str` |  |
| `customer_account_number` | `str` |  |
| `customer_city` | `str` |  |
| `customer_country_code` | `str` |  |
| `customer_country_name` | `str` |  |
| `customer_email_address` | `str` |  |
| `customer_first_name` | `str` |  |
| `customer_iban` | `str` |  |
| `customer_last_name` | `str` |  |
| `customer_sort_code` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `is_recurring` | `bool` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `reference` | `str` |  |
| `status` | `str` |  |
| `supplier_bank_account_id` | `str` |  |
| `supplier_customer_id` | `str` |  |
| `supplier_mandate_id` | `str` |  |
| `supplier_name` | `str` |  |
| `supplier_status` | `str` |  |

#### Example: List

```python
no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_pages = client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage().list()
```


### NoFrixionBizBizModelsPagingPaymentRequestPage

Create an instance: `no_frixion_biz_biz_models_paging_payment_request_page = client.NoFrixionBizBizModelsPagingPaymentRequestPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `list` |  |
| `amount` | `float` |  |
| `amount_pending` | `float` |  |
| `amount_received` | `float` |  |
| `amount_refunded` | `float` |  |
| `auto_send_receipt` | `bool` |  |
| `base_origin_url` | `str` |  |
| `callback_url` | `str` |  |
| `card_authorize_only` | `bool` |  |
| `card_create_token` | `bool` |  |
| `card_create_token_mode` | `str` |  |
| `card_ignore_cvn` | `bool` |  |
| `card_processor_merchant_id` | `str` |  |
| `card_stripe_payment_intent_id` | `str` |  |
| `card_stripe_payment_intent_secret` | `str` |  |
| `created_by_user` | `dict` |  |
| `currency` | `str` |  |
| `custom_field` | `list` |  |
| `customer_email_address` | `str` |  |
| `customer_id` | `str` |  |
| `customer_name` | `str` |  |
| `description` | `str` |  |
| `destination_account` | `dict` |  |
| `direct_debit_payment` | `dict` |  |
| `due_date` | `str` |  |
| `event` | `list` |  |
| `failure_callback_url` | `str` |  |
| `field_display_setting` | `list` |  |
| `formatted_amount` | `str` |  |
| `hosted_pay_checkout_url` | `str` |  |
| `id` | `str` |  |
| `ignore_address_verification` | `bool` |  |
| `inserted` | `str` |  |
| `inserted_sortable` | `str` |  |
| `is_archived` | `bool` |  |
| `jwk` | `str` |  |
| `last_updated` | `str` |  |
| `lightning_invoice` | `str` |  |
| `lightning_invoice_expires_at` | `str` |  |
| `merchant_direct_debit_mandate_id` | `str` |  |
| `merchant_id` | `str` |  |
| `merchant_token_description` | `str` |  |
| `notification_email_address` | `str` |  |
| `notification_role_i_d` | `list` |  |
| `order_id` | `str` |  |
| `partial_payment_method` | `str` |  |
| `partial_payment_step` | `str` |  |
| `payment_attempt` | `list` |  |
| `payment_method` | `list` |  |
| `payment_processor` | `str` |  |
| `payrun_id` | `str` |  |
| `pisp_account_id` | `str` |  |
| `priority_bank_id` | `str` |  |
| `result` | `dict` |  |
| `sandbox_settle_delay_in_second` | `int` |  |
| `shipping_address` | `dict` |  |
| `status` | `str` |  |
| `success_web_hook_url` | `str` |  |
| `tag` | `list` |  |
| `title` | `str` |  |
| `tokenised_card` | `list` |  |
| `transaction` | `list` |  |
| `use_hosted_payment_page` | `bool` |  |

#### Example: List

```python
no_frixion_biz_biz_models_paging_payment_request_pages = client.NoFrixionBizBizModelsPagingPaymentRequestPage().list()
```


### NoFrixionBizBizModelsPagingPayoutPage

Create an instance: `no_frixion_biz_biz_models_paging_payout_page = client.NoFrixionBizBizModelsPagingPayoutPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `str` |  |
| `amount` | `float` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `str` |  |
| `approver_id` | `str` |  |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `str` |  |
| `beneficiary` | `dict` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `str` |  |
| `created_by` | `str` |  |
| `created_by_email_address` | `str` |  |
| `currency` | `str` |  |
| `current_user_id` | `str` |  |
| `description` | `str` |  |
| `destination` | `dict` |  |
| `document` | `list` |  |
| `event` | `list` |  |
| `formatted_amount` | `str` |  |
| `formatted_fx_destination_amount` | `str` |  |
| `formatted_schedule` | `str` |  |
| `formatted_schedule_day_only` | `str` |  |
| `formatted_source_account_available_balance` | `str` |  |
| `fx_destination_amount` | `float` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `str` |  |
| `fx_quote_expires_at` | `str` |  |
| `fx_quote_id` | `str` |  |
| `fx_rate` | `float` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `invoice_id` | `str` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `merchant_token_description` | `str` |  |
| `nonce` | `str` |  |
| `payment_processor` | `str` |  |
| `payment_rail` | `str` |  |
| `payrun_id` | `str` |  |
| `payrun_name` | `str` |  |
| `rule` | `dict` |  |
| `schedule_date` | `str` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `str` |  |
| `source_account_currency` | `str` |  |
| `source_account_iban` | `str` |  |
| `source_account_identifier` | `dict` |  |
| `source_account_name` | `str` |  |
| `source_account_number` | `str` |  |
| `source_account_sortcode` | `str` |  |
| `status` | `str` |  |
| `tag` | `list` |  |
| `their_reference` | `str` |  |
| `topup_payrun_id` | `str` |  |
| `transacted_amount` | `float` |  |
| `transacted_fx_amount` | `float` |  |
| `transacted_fx_rate` | `float` |  |
| `type` | `str` |  |
| `user_id` | `str` |  |
| `your_reference` | `str` |  |

#### Example: List

```python
no_frixion_biz_biz_models_paging_payout_pages = client.NoFrixionBizBizModelsPagingPayoutPage().list()
```


### NoFrixionBizBizModelsPagingPayrunPage

Create an instance: `no_frixion_biz_biz_models_paging_payrun_page = client.NoFrixionBizBizModelsPagingPayrunPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation` | `list` |  |
| `authorisation_date` | `str` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `str` |  |
| `can_authorise` | `bool` |  |
| `can_delete` | `bool` |  |
| `can_edit` | `bool` |  |
| `event` | `list` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `invoice` | `list` |  |
| `invoices_minimal` | `list` |  |
| `is_archived` | `bool` |  |
| `last_updated` | `str` |  |
| `last_updated_by` | `dict` |  |
| `merchant_id` | `str` |  |
| `name` | `str` |  |
| `nonce` | `str` |  |
| `payment` | `list` |  |
| `payout` | `list` |  |
| `payouts_count` | `int` |  |
| `schedule_date` | `str` |  |
| `source_account` | `list` |  |
| `status` | `str` |  |
| `total_eur` | `float` |  |
| `total_gbp` | `float` |  |
| `total_usd` | `float` |  |

#### Example: List

```python
no_frixion_biz_biz_models_paging_payrun_pages = client.NoFrixionBizBizModelsPagingPayrunPage().list()
```


### NoFrixionBizBizModelsPagingRuleEventsPage

Create an instance: `no_frixion_biz_biz_models_paging_rule_events_page = client.NoFrixionBizBizModelsPagingRuleEventsPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_message` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `is_authorise_to_enable` | `bool` |  |
| `message` | `str` |  |
| `raw_response` | `str` |  |
| `rule_event_type` | `str` |  |
| `rule_id` | `str` |  |
| `user` | `dict` |  |

#### Example: List

```python
no_frixion_biz_biz_models_paging_rule_events_pages = client.NoFrixionBizBizModelsPagingRuleEventsPage().list()
```


### NoFrixionBizBizModelsPagingRulesPage

Create an instance: `no_frixion_biz_biz_models_paging_rules_page = client.NoFrixionBizBizModelsPagingRulesPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account` | `dict` |  |
| `account_id` | `str` |  |
| `approve_url` | `str` |  |
| `approver_id` | `str` |  |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `can_authorise` | `bool` |  |
| `created_by` | `dict` |  |
| `description` | `str` |  |
| `end_at` | `str` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `is_disabled` | `bool` |  |
| `last_executed_at` | `str` |  |
| `last_run_at_transaction_date` | `str` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `name` | `str` |  |
| `nonce` | `str` |  |
| `on_approved_web_hook_url` | `str` |  |
| `on_execution_error_web_hook_url` | `str` |  |
| `on_execution_success_web_hook_url` | `str` |  |
| `start_at` | `str` |  |
| `status` | `str` |  |
| `sweep_action` | `dict` |  |
| `time_zone_id` | `str` |  |
| `trigger_cron_expression` | `str` |  |
| `trigger_on_pay_in` | `bool` |  |
| `user_id` | `str` |  |
| `web_hook_secret` | `str` |  |

#### Example: List

```python
no_frixion_biz_biz_models_paging_rules_pages = client.NoFrixionBizBizModelsPagingRulesPage().list()
```


### NoFrixionBizBizModelsPaymentsCardPayment

Create an instance: `no_frixion_biz_biz_models_payments_card_payment = client.NoFrixionBizBizModelsPaymentsCardPayment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorized_amount` | `str` |  |
| `currency_code` | `str` |  |
| `is_payer_authentication_required` | `bool` |  |
| `is_soft_decline` | `bool` |  |
| `payer_authentication_access_token` | `str` |  |
| `payer_authentication_merchant_data` | `str` |  |
| `payer_authentication_url` | `str` |  |
| `payer_authentication_window_height` | `int` |  |
| `payer_authentication_window_width` | `int` |  |
| `payment_request_callback_url` | `str` |  |
| `payment_request_id` | `str` |  |
| `request_id` | `str` |  |
| `response_code` | `str` |  |
| `response_type` | `str` |  |
| `status` | `str` |  |
| `three_ds_redirect_url` | `str` |  |
| `transaction_id` | `str` |  |

#### Example: Create

```python
no_frixion_biz_biz_models_payments_card_payment = client.NoFrixionBizBizModelsPaymentsCardPayment().create({
    "paymentrequest_id": "example_paymentrequest_id",  # str
})
```


### NoFrixionBizBizModelsPaymentsCardPublicKey

Create an instance: `no_frixion_biz_biz_models_payments_card_public_key = client.NoFrixionBizBizModelsPaymentsCardPublicKey()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `jwt` | `str` |  |

#### Example: Load

```python
no_frixion_biz_biz_models_payments_card_public_key = client.NoFrixionBizBizModelsPaymentsCardPublicKey().load({"paymentrequest_id": "paymentrequest_id"})
```


### NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries

Create an instance: `no_frixion_money_moov_api_features_beneficiaries_beneficiaries = client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `beneficiary` | `list` |  |
| `failed_beneficiary` | `dict` |  |

#### Example: Create

```python
no_frixion_money_moov_api_features_beneficiaries_beneficiaries = client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries().create({
})
```


### NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment

Create an instance: `no_frixion_money_moov_api_features_payment_requests_payment = client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_payment_request` | `dict` |  |
| `payment_request` | `list` |  |

#### Example: Create

```python
no_frixion_money_moov_api_features_payment_requests_payment = client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment().create({
})
```


### NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate

Create an instance: `no_frixion_money_moov_api_features_permissions_roles_create = client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_role` | `dict` |  |
| `role` | `list` |  |

#### Example: Create

```python
no_frixion_money_moov_api_features_permissions_roles_create = client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate().create({
    "merchant_id": "example_merchant_id",  # str
})
```


### NoFrixionMoneyMoovApiFeaturesUserInvitesCreate

Create an instance: `no_frixion_money_moov_api_features_user_invites_create = client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_user_invite` | `dict` |  |
| `user_invite` | `list` |  |

#### Example: Create

```python
no_frixion_money_moov_api_features_user_invites_create = client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate().create({
})
```


### NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant

Create an instance: `no_frixion_money_moov_models_authorisation_settings_merchant = client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount_lower` | `float` |  |
| `amount_upper` | `float` |  |
| `authorisation_type` | `str` |  |
| `beneficiaries_only` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `last_editor_cant_authorise` | `bool` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `number_of_authoriser` | `int` |  |
| `role_setting` | `list` |  |

#### Example: List

```python
no_frixion_money_moov_models_authorisation_settings_merchants = client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant().list()
```


### NoFrixionMoneyMoovModelsBatchPayout

Create an instance: `no_frixion_money_moov_models_batch_payout = client.NoFrixionMoneyMoovModelsBatchPayout()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approve_url` | `str` |  |
| `id` | `str` |  |
| `payout` | `list` |  |

#### Example: Load

```python
no_frixion_money_moov_models_batch_payout = client.NoFrixionMoneyMoovModelsBatchPayout().load({"id": "no_frixion_money_moov_models_batch_payout_id"})
```

#### Example: Create

```python
no_frixion_money_moov_models_batch_payout = client.NoFrixionMoneyMoovModelsBatchPayout().create({
})
```


### NoFrixionMoneyMoovModelsBeneficiaryGroupPage

Create an instance: `no_frixion_money_moov_models_beneficiary_group_page = client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `group_member` | `list` |  |
| `group_name` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |

#### Example: List

```python
no_frixion_money_moov_models_beneficiary_group_pages = client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage().list()
```


### NoFrixionMoneyMoovModelsBeneficiaryPage

Create an instance: `no_frixion_money_moov_models_beneficiary_page = client.NoFrixionMoneyMoovModelsBeneficiaryPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_callback_url` | `str` |  |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `beneficiary_event` | `list` |  |
| `can_authorise` | `bool` |  |
| `can_update` | `bool` |  |
| `created_by` | `dict` |  |
| `created_by_email_address` | `str` |  |
| `currency` | `str` |  |
| `destination` | `dict` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `is_enabled` | `bool` |  |
| `last_authorised` | `str` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `name` | `str` |  |
| `nonce` | `str` |  |
| `source_account` | `list` |  |
| `their_reference` | `str` |  |

#### Example: List

```python
no_frixion_money_moov_models_beneficiary_pages = client.NoFrixionMoneyMoovModelsBeneficiaryPage().list()
```


### NoFrixionMoneyMoovModelsCardCustomerToken

Create an instance: `no_frixion_money_moov_models_card_customer_token = client.NoFrixionMoneyMoovModelsCardCustomerToken()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `card_type` | `str` |  |
| `customer_email_address` | `str` |  |
| `expiry_month` | `str` |  |
| `expiry_year` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `last_four_digit` | `str` |  |
| `last_updated` | `str` |  |
| `masked_card_number` | `str` |  |
| `merchant_id` | `str` |  |
| `payment_request_id` | `str` |  |

#### Example: Load

```python
no_frixion_money_moov_models_card_customer_token = client.NoFrixionMoneyMoovModelsCardCustomerToken().load({"customer_email_address": "customer_email_address"})
```

#### Example: List

```python
no_frixion_money_moov_models_card_customer_tokens = client.NoFrixionMoneyMoovModelsCardCustomerToken().list()
```


### NoFrixionMoneyMoovModelsCurrencyCurrencyInfo

Create an instance: `no_frixion_money_moov_models_currency_currency_info = client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `code` | `str` |  |
| `decimal` | `int` |  |
| `is_fiat` | `bool` |  |
| `iso4217_alpha_code` | `str` |  |
| `iso4217_numeric_code` | `str` |  |
| `symbol` | `str` |  |

#### Example: List

```python
no_frixion_money_moov_models_currency_currency_infos = client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo().list()
```


### NoFrixionMoneyMoovModelsDirectDebitBatchSubmit

Create an instance: `no_frixion_money_moov_models_direct_debit_batch_submit = client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_submission` | `dict` |  |
| `successful_submission` | `list` |  |

#### Example: Create

```python
no_frixion_money_moov_models_direct_debit_batch_submit = client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit().create({
})
```


### NoFrixionMoneyMoovModelsFxRate

Create an instance: `no_frixion_money_moov_models_fx_rate = client.NoFrixionMoneyMoovModelsFxRate()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `destination_currency` | `str` |  |
| `exchange_rate` | `float` |  |
| `expiry_time` | `str` |  |
| `quote_id` | `str` |  |
| `source_currency` | `str` |  |

#### Example: Load

```python
no_frixion_money_moov_models_fx_rate = client.NoFrixionMoneyMoovModelsFxRate().load({"destination": "destination", "source": "source", "valid_for_minute": 1})
```

#### Example: List

```python
no_frixion_money_moov_models_fx_rates = client.NoFrixionMoneyMoovModelsFxRate().list()
```


### NoFrixionMoneyMoovModelsIPayment

Create an instance: `no_frixion_money_moov_models_i_payment = client.NoFrixionMoneyMoovModelsIPayment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `payment_request_id` | `str` |  |
| `response_type` | `str` |  |

#### Example: Create

```python
no_frixion_money_moov_models_i_payment = client.NoFrixionMoneyMoovModelsIPayment().create({
})
```


### NoFrixionMoneyMoovModelsMandatesMandate

Create an instance: `no_frixion_money_moov_models_mandates_mandate = client.NoFrixionMoneyMoovModelsMandatesMandate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_number` | `str` |  |
| `address_line1` | `str` |  |
| `address_line2` | `str` |  |
| `approved_at` | `str` |  |
| `city` | `str` |  |
| `country_code` | `str` |  |
| `currency` | `str` |  |
| `customer_account_number` | `str` |  |
| `customer_city` | `str` |  |
| `customer_country_code` | `str` |  |
| `customer_country_name` | `str` |  |
| `customer_email_address` | `str` |  |
| `customer_first_name` | `str` |  |
| `customer_iban` | `str` |  |
| `customer_last_name` | `str` |  |
| `customer_sort_code` | `str` |  |
| `email_address` | `str` |  |
| `first_name` | `str` |  |
| `iban` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `is_recurring` | `bool` |  |
| `last_name` | `str` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `postal_code` | `str` |  |
| `reference` | `str` |  |
| `sort_code` | `str` |  |
| `status` | `str` |  |
| `supplier_bank_account_id` | `str` |  |
| `supplier_customer_id` | `str` |  |
| `supplier_mandate_id` | `str` |  |
| `supplier_name` | `str` |  |
| `supplier_status` | `str` |  |

#### Example: Load

```python
no_frixion_money_moov_models_mandates_mandate = client.NoFrixionMoneyMoovModelsMandatesMandate().load({"id": "no_frixion_money_moov_models_mandates_mandate_id"})
```

#### Example: Create

```python
no_frixion_money_moov_models_mandates_mandate = client.NoFrixionMoneyMoovModelsMandatesMandate().create({
    "address_line1": "example_address_line1",  # str
    "city": "example_city",  # str
    "country_code": "example_country_code",  # str
    "email_address": "example_email_address",  # str
    "first_name": "example_first_name",  # str
    "last_name": "example_last_name",  # str
    "postal_code": "example_postal_code",  # str
})
```


### NoFrixionMoneyMoovModelsMerchant

Create an instance: `no_frixion_money_moov_models_merchant = client.NoFrixionMoneyMoovModelsMerchant()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_currency` | `list` |  |
| `can_have_trust_account` | `bool` |  |
| `card_payment_processor` | `str` |  |
| `company_id` | `str` |  |
| `display_qr_on_hosted_pay` | `bool` |  |
| `hosted_pay_version` | `int` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `is_blocked` | `bool` |  |
| `is_exited` | `bool` |  |
| `is_suspended` | `bool` |  |
| `jurisdiction` | `str` |  |
| `logo_url_png` | `str` |  |
| `logo_url_svg` | `str` |  |
| `merchant_category_code` | `str` |  |
| `name` | `str` |  |
| `note` | `str` |  |
| `parent_merchant` | `dict` |  |
| `payment_account` | `list` |  |
| `payment_account_limit` | `int` |  |
| `short_name` | `str` |  |
| `supported_payment_methods_list` | `list` |  |
| `suspension_reason` | `str` |  |
| `tag` | `list` |  |
| `time_zone_id` | `str` |  |
| `trading_name` | `str` |  |
| `web_hook_limit` | `int` |  |
| `your_role_name` | `str` |  |

#### Example: Load

```python
no_frixion_money_moov_models_merchant = client.NoFrixionMoneyMoovModelsMerchant().load({"id": "no_frixion_money_moov_models_merchant_id"})
```

#### Example: List

```python
no_frixion_money_moov_models_merchants = client.NoFrixionMoneyMoovModelsMerchant().list()
```


### NoFrixionMoneyMoovModelsMerchantPage

Create an instance: `no_frixion_money_moov_models_merchant_page = client.NoFrixionMoneyMoovModelsMerchantPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_currency` | `list` |  |
| `can_have_trust_account` | `bool` |  |
| `card_payment_processor` | `str` |  |
| `company_id` | `str` |  |
| `display_qr_on_hosted_pay` | `bool` |  |
| `hosted_pay_version` | `int` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `is_blocked` | `bool` |  |
| `is_exited` | `bool` |  |
| `is_suspended` | `bool` |  |
| `jurisdiction` | `str` |  |
| `logo_url_png` | `str` |  |
| `logo_url_svg` | `str` |  |
| `merchant_category_code` | `str` |  |
| `name` | `str` |  |
| `note` | `str` |  |
| `parent_merchant` | `dict` |  |
| `payment_account` | `list` |  |
| `payment_account_limit` | `int` |  |
| `short_name` | `str` |  |
| `supported_payment_methods_list` | `list` |  |
| `suspension_reason` | `str` |  |
| `tag` | `list` |  |
| `time_zone_id` | `str` |  |
| `trading_name` | `str` |  |
| `web_hook_limit` | `int` |  |
| `your_role_name` | `str` |  |

#### Example: List

```python
no_frixion_money_moov_models_merchant_pages = client.NoFrixionMoneyMoovModelsMerchantPage().list()
```


### NoFrixionMoneyMoovModelsMerchantPayByBankSetting

Create an instance: `no_frixion_money_moov_models_merchant_pay_by_bank_setting = client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bank_country_code` | `list` |  |
| `bank_id` | `str` |  |
| `bank_name` | `str` |  |
| `business_institution_id` | `str` |  |
| `currency` | `str` |  |
| `logo` | `str` |  |
| `message` | `str` |  |
| `message_image_url` | `str` |  |
| `order` | `int` |  |
| `personal_institution_id` | `str` |  |
| `processor` | `str` |  |
| `warning_heading` | `str` |  |
| `warning_message` | `str` |  |

#### Example: List

```python
no_frixion_money_moov_models_merchant_pay_by_bank_settings = client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting().list()
```


### NoFrixionMoneyMoovModelsMerchantToken

Create an instance: `no_frixion_money_moov_models_merchant_token = client.NoFrixionMoneyMoovModelsMerchantToken()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `can_authorise` | `bool` |  |
| `description` | `str` |  |
| `expires_at` | `str` |  |
| `has_current_user_authorised` | `bool` |  |
| `hmac_algorithm` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `ip_address_whitelist` | `str` |  |
| `is_archived` | `bool` |  |
| `is_enabled` | `bool` |  |
| `last_authorised` | `str` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `nonce` | `str` |  |
| `permission_type` | `list` |  |
| `request_signature_version` | `int` |  |
| `shared_secret_algorithm` | `str` |  |
| `shared_secret_base64` | `str` |  |
| `token` | `str` |  |

#### Example: Load

```python
no_frixion_money_moov_models_merchant_token = client.NoFrixionMoneyMoovModelsMerchantToken().load({"id": "no_frixion_money_moov_models_merchant_token_id"})
```

#### Example: Create

```python
no_frixion_money_moov_models_merchant_token = client.NoFrixionMoneyMoovModelsMerchantToken().create({
    "nonce": "example_nonce",  # str
})
```


### NoFrixionMoneyMoovModelsMerchantTokenPage

Create an instance: `no_frixion_money_moov_models_merchant_token_page = client.NoFrixionMoneyMoovModelsMerchantTokenPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `can_authorise` | `bool` |  |
| `description` | `str` |  |
| `expires_at` | `str` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `ip_address_whitelist` | `str` |  |
| `is_archived` | `bool` |  |
| `is_enabled` | `bool` |  |
| `last_authorised` | `str` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `nonce` | `str` |  |
| `permission_type` | `list` |  |
| `request_signature_version` | `int` |  |
| `shared_secret_algorithm` | `str` |  |
| `shared_secret_base64` | `str` |  |
| `token` | `str` |  |

#### Example: List

```python
no_frixion_money_moov_models_merchant_token_pages = client.NoFrixionMoneyMoovModelsMerchantTokenPage().list()
```


### NoFrixionMoneyMoovModelsNoFrixionVersion

Create an instance: `no_frixion_money_moov_models_no_frixion_version = client.NoFrixionMoneyMoovModelsNoFrixionVersion()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `build_version` | `int` |  |
| `major_version` | `int` |  |
| `minor_version` | `int` |  |
| `release_name` | `str` |  |

#### Example: Load

```python
no_frixion_money_moov_models_no_frixion_version = client.NoFrixionMoneyMoovModelsNoFrixionVersion().load()
```


### NoFrixionMoneyMoovModelsOpenBankingAccount

Create an instance: `no_frixion_money_moov_models_open_banking_account = client.NoFrixionMoneyMoovModelsOpenBankingAccount()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_balance` | `list` |  |
| `account_identification` | `list` |  |
| `account_name` | `list` |  |
| `account_type` | `str` |  |
| `balance` | `float` |  |
| `consolidated_account_information` | `dict` |  |
| `currency` | `str` |  |
| `description` | `str` |  |
| `detail` | `str` |  |
| `id` | `str` |  |
| `nickname` | `str` |  |
| `type` | `str` |  |
| `usage_type` | `str` |  |

#### Example: Load

```python
no_frixion_money_moov_models_open_banking_account = client.NoFrixionMoneyMoovModelsOpenBankingAccount().load({"id": "no_frixion_money_moov_models_open_banking_account_id"})
```


### NoFrixionMoneyMoovModelsOpenBankingConsent

Create an instance: `no_frixion_money_moov_models_open_banking_consent = client.NoFrixionMoneyMoovModelsOpenBankingConsent()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_url` | `str` |  |
| `callback_url` | `str` |  |
| `consent_id` | `str` |  |
| `email_address` | `str` |  |
| `expiry_date` | `str` |  |
| `failure_callback_url` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `institution_id` | `str` |  |
| `is_connected_account` | `bool` |  |
| `is_enabled` | `bool` |  |
| `merchant_id` | `str` |  |
| `provider` | `str` |  |
| `success_web_hook_url` | `str` |  |

#### Example: Load

```python
no_frixion_money_moov_models_open_banking_consent = client.NoFrixionMoneyMoovModelsOpenBankingConsent().load({"id": "no_frixion_money_moov_models_open_banking_consent_id"})
```

#### Example: List

```python
no_frixion_money_moov_models_open_banking_consents = client.NoFrixionMoneyMoovModelsOpenBankingConsent().list()
```

#### Example: Create

```python
no_frixion_money_moov_models_open_banking_consent = client.NoFrixionMoneyMoovModelsOpenBankingConsent().create({
})
```


### NoFrixionMoneyMoovModelsOpenBankingTransaction

Create an instance: `no_frixion_money_moov_models_open_banking_transaction = client.NoFrixionMoneyMoovModelsOpenBankingTransaction()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address_detail` | `dict` |  |
| `amount` | `float` |  |
| `balance` | `dict` |  |
| `booking_date_time` | `str` |  |
| `charge_detail` | `dict` |  |
| `currency` | `str` |  |
| `currency_exchange` | `dict` |  |
| `date` | `str` |  |
| `description` | `str` |  |
| `enrichment` | `dict` |  |
| `gross_amount` | `dict` |  |
| `id` | `str` |  |
| `iso_bank_transaction_code` | `dict` |  |
| `merchant` | `dict` |  |
| `payee_detail` | `dict` |  |
| `payer_detail` | `dict` |  |
| `proprietary_bank_transaction_code` | `dict` |  |
| `reference` | `str` |  |
| `statement_reference` | `list` |  |
| `status` | `str` |  |
| `supplementary_data` | `Any` |  |
| `transaction_amount` | `dict` |  |
| `transaction_information` | `list` |  |
| `transaction_mutability` | `str` |  |
| `value_date_time` | `str` |  |

#### Example: List

```python
no_frixion_money_moov_models_open_banking_transactions = client.NoFrixionMoneyMoovModelsOpenBankingTransaction().list()
```


### NoFrixionMoneyMoovModelsPayment

Create an instance: `no_frixion_money_moov_models_payment = client.NoFrixionMoneyMoovModelsPayment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `list` |  |
| `amount` | `float` |  |
| `amount_pending` | `float` |  |
| `amount_received` | `float` |  |
| `amount_refunded` | `float` |  |
| `auto_send_receipt` | `bool` |  |
| `base_origin_url` | `str` |  |
| `callback_url` | `str` |  |
| `card_authorize_only` | `bool` |  |
| `card_create_token` | `bool` |  |
| `card_create_token_mode` | `str` |  |
| `card_ignore_cvn` | `bool` |  |
| `card_no_payer_authentication` | `bool` |  |
| `card_processor_merchant_id` | `str` |  |
| `card_stripe_payment_intent_id` | `str` |  |
| `card_stripe_payment_intent_secret` | `str` |  |
| `card_transmit_raw_detail` | `bool` |  |
| `created_by_user` | `dict` |  |
| `currency` | `str` |  |
| `custom_field` | `list` |  |
| `customer_email_address` | `str` |  |
| `customer_id` | `str` |  |
| `customer_name` | `str` |  |
| `description` | `str` |  |
| `destination_account` | `dict` |  |
| `direct_debit_payment` | `dict` |  |
| `due_date` | `str` |  |
| `event` | `list` |  |
| `failure_callback_url` | `str` |  |
| `field_display_setting` | `list` |  |
| `formatted_amount` | `str` |  |
| `hosted_pay_checkout_url` | `str` |  |
| `id` | `str` |  |
| `ignore_address_verification` | `bool` |  |
| `inserted` | `str` |  |
| `inserted_sortable` | `str` |  |
| `is_archived` | `bool` |  |
| `jwk` | `str` |  |
| `last_updated` | `str` |  |
| `lightning_invoice` | `str` |  |
| `lightning_invoice_expires_at` | `str` |  |
| `merchant_direct_debit_mandate_id` | `str` |  |
| `merchant_id` | `str` |  |
| `merchant_token_description` | `str` |  |
| `notification_email_address` | `str` |  |
| `notification_role_i_d` | `list` |  |
| `order_id` | `str` |  |
| `partial_payment_method` | `str` |  |
| `partial_payment_step` | `str` |  |
| `payment_attempt` | `list` |  |
| `payment_method` | `list` |  |
| `payment_processor` | `str` |  |
| `payrun_id` | `str` |  |
| `pisp_account_id` | `str` |  |
| `priority_bank_id` | `str` |  |
| `result` | `dict` |  |
| `sandbox_settle_delay_in_second` | `int` |  |
| `shipping_address` | `dict` |  |
| `shipping_address_city` | `str` |  |
| `shipping_address_country_code` | `str` |  |
| `shipping_address_county` | `str` |  |
| `shipping_address_line1` | `str` |  |
| `shipping_address_line2` | `str` |  |
| `shipping_address_post_code` | `str` |  |
| `shipping_email` | `str` |  |
| `shipping_first_name` | `str` |  |
| `shipping_last_name` | `str` |  |
| `shipping_phone` | `str` |  |
| `status` | `str` |  |
| `success_web_hook_url` | `str` |  |
| `tag` | `list` |  |
| `tag_id` | `list` |  |
| `title` | `str` |  |
| `tokenised_card` | `list` |  |
| `transaction` | `list` |  |
| `use_hosted_payment_page` | `bool` |  |

#### Example: Load

```python
no_frixion_money_moov_models_payment = client.NoFrixionMoneyMoovModelsPayment().load({"id": "no_frixion_money_moov_models_payment_id"})
```

#### Example: Create

```python
no_frixion_money_moov_models_payment = client.NoFrixionMoneyMoovModelsPayment().create({
    "created_by_user": {},  # dict
})
```


### NoFrixionMoneyMoovModelsPaymentAccountMinimalPage

Create an instance: `no_frixion_money_moov_models_payment_account_minimal_page = client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `str` |  |
| `available_balance` | `float` |  |
| `balance` | `float` |  |
| `balance_minor_unit` | `int` |  |
| `currency` | `str` |  |
| `id` | `str` |  |
| `identifier` | `dict` |  |
| `is_archived` | `bool` |  |
| `is_connected_account` | `bool` |  |
| `merchant_id` | `str` |  |
| `submitted_payouts_balance` | `float` |  |

#### Example: List

```python
no_frixion_money_moov_models_payment_account_minimal_pages = client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage().list()
```


### NoFrixionMoneyMoovModelsPaymentAccountPage

Create an instance: `no_frixion_money_moov_models_payment_account_page = client.NoFrixionMoneyMoovModelsPaymentAccountPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `str` |  |
| `account_supplier_name` | `str` |  |
| `available_balance` | `float` |  |
| `available_balance_minor_unit` | `int` |  |
| `balance` | `float` |  |
| `balance_minor_unit` | `int` |  |
| `bank_name` | `str` |  |
| `consent_id` | `str` |  |
| `created_by` | `dict` |  |
| `created_by_display_name` | `str` |  |
| `currency` | `str` |  |
| `default_payment_rail` | `str` |  |
| `display_name` | `str` |  |
| `expiry_date` | `str` |  |
| `external_account_icon` | `str` |  |
| `id` | `str` |  |
| `identifier` | `dict` |  |
| `inserted` | `str` |  |
| `is_archived` | `bool` |  |
| `is_connected_account` | `bool` |  |
| `is_default` | `bool` |  |
| `is_trust_account` | `bool` |  |
| `is_virtual` | `bool` |  |
| `last_transaction` | `dict` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `merchant_name` | `str` |  |
| `physical_account_id` | `str` |  |
| `rule` | `list` |  |
| `submitted_payouts_balance` | `float` |  |
| `submitted_payouts_balance_minor_unit` | `int` |  |
| `summary` | `str` |  |
| `supplier_sepa_instant_status` | `str` |  |
| `xero_bank_feed_connection_status` | `str` |  |
| `xero_bank_feed_last_synced_at` | `str` |  |
| `xero_bank_feed_sync_last_failed_at` | `str` |  |
| `xero_bank_feed_sync_last_failure_reason` | `str` |  |
| `xero_bank_feed_sync_status` | `str` |  |
| `xero_unsynchronised_transactions_count` | `int` |  |

#### Example: List

```python
no_frixion_money_moov_models_payment_account_pages = client.NoFrixionMoneyMoovModelsPaymentAccountPage().list()
```


### NoFrixionMoneyMoovModelsPaymentInitiation

Create an instance: `no_frixion_money_moov_models_payment_initiation = client.NoFrixionMoneyMoovModelsPaymentInitiation()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `payment_initiation_id` | `str` |  |
| `payment_request_callback_url` | `str` |  |
| `payment_request_id` | `str` |  |
| `redirect_url` | `str` |  |
| `response_type` | `str` |  |
| `specific_error_message` | `str` |  |

#### Example: Create

```python
no_frixion_money_moov_models_payment_initiation = client.NoFrixionMoneyMoovModelsPaymentInitiation().create({
    "paymentrequest_id": "example_paymentrequest_id",  # str
})
```


### NoFrixionMoneyMoovModelsPaymentRequestEvent

Create an instance: `no_frixion_money_moov_models_payment_request_event = client.NoFrixionMoneyMoovModelsPaymentRequestEvent()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float` |  |
| `apple_pay_transaction_id` | `str` |  |
| `card_authorization_response_id` | `str` |  |
| `card_expiry_month` | `int` |  |
| `card_expiry_year` | `int` |  |
| `card_issuer` | `str` |  |
| `card_issuer_country` | `str` |  |
| `card_last_four_digit` | `str` |  |
| `card_request_id` | `str` |  |
| `card_scheme` | `str` |  |
| `card_token_customer_id` | `str` |  |
| `card_transaction_id` | `str` |  |
| `currency` | `str` |  |
| `direct_debit_payment_id` | `str` |  |
| `direct_debit_payment_reference` | `str` |  |
| `drirect_debit_mandate_id` | `str` |  |
| `error_message` | `str` |  |
| `error_reason` | `str` |  |
| `event_type` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `lightning_invoice` | `str` |  |
| `lightning_r_hash` | `str` |  |
| `origin_url` | `str` |  |
| `payment_method_type` | `str` |  |
| `payment_processor_name` | `str` |  |
| `payment_request_id` | `str` |  |
| `pisp_bank_status` | `str` |  |
| `pisp_payment_initiation_id` | `str` |  |
| `pisp_payment_institution_name` | `str` |  |
| `pisp_payment_service_provider_id` | `str` |  |
| `pisp_redirect_url` | `str` |  |
| `reconciled_transaction_id` | `str` |  |
| `refund_payout_id` | `str` |  |
| `status` | `str` |  |
| `wallet_name` | `str` |  |

#### Example: List

```python
no_frixion_money_moov_models_payment_request_events = client.NoFrixionMoneyMoovModelsPaymentRequestEvent().list()
```


### NoFrixionMoneyMoovModelsPaymentRequestMetric

Create an instance: `no_frixion_money_moov_models_payment_request_metric = client.NoFrixionMoneyMoovModelsPaymentRequestMetric()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `all` | `int` |  |
| `authorized` | `int` |  |
| `paid` | `int` |  |
| `partially_paid` | `int` |  |
| `total_amounts_by_currency` | `dict` |  |
| `unpaid` | `int` |  |

#### Example: Load

```python
no_frixion_money_moov_models_payment_request_metric = client.NoFrixionMoneyMoovModelsPaymentRequestMetric().load()
```


### NoFrixionMoneyMoovModelsPaymentRequestMinimal

Create an instance: `no_frixion_money_moov_models_payment_request_minimal = client.NoFrixionMoneyMoovModelsPaymentRequestMinimal()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float` |  |
| `amount_pending` | `float` |  |
| `amount_received` | `float` |  |
| `amount_refunded` | `float` |  |
| `callback_url` | `str` |  |
| `card_stripe_payment_intent_secret` | `str` |  |
| `country_code` | `str` |  |
| `currency` | `str` |  |
| `custom_fields_to_display` | `list` |  |
| `description` | `str` |  |
| `due_date` | `str` |  |
| `field_display_setting` | `list` |  |
| `google_pay_merchant_id` | `str` |  |
| `id` | `str` |  |
| `jwk` | `str` |  |
| `merchant_id` | `str` |  |
| `merchant_logo_url_png` | `str` |  |
| `merchant_logo_url_svg` | `str` |  |
| `merchant_name` | `str` |  |
| `merchant_short_name` | `str` |  |
| `partial_payment_method` | `str` |  |
| `payment_attempt` | `list` |  |
| `payment_methods_list` | `list` |  |
| `payment_processor` | `str` |  |
| `payment_processor_key` | `str` |  |
| `pisp_error` | `str` |  |
| `priority_bank_id` | `str` |  |
| `status` | `str` |  |
| `stripe_account_id` | `str` |  |
| `title` | `str` |  |

#### Example: List

```python
no_frixion_money_moov_models_payment_request_minimals = client.NoFrixionMoneyMoovModelsPaymentRequestMinimal().list()
```


### NoFrixionMoneyMoovModelsPaymentRequestResult

Create an instance: `no_frixion_money_moov_models_payment_request_result = client.NoFrixionMoneyMoovModelsPaymentRequestResult()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float` |  |
| `amount_pending` | `float` |  |
| `amount_received` | `float` |  |
| `amount_refunded` | `float` |  |
| `currency` | `str` |  |
| `customer_id` | `str` |  |
| `payment` | `list` |  |
| `payment_request_id` | `str` |  |
| `pisp_authorization` | `list` |  |
| `requested_amount` | `float` |  |
| `result` | `str` |  |

#### Example: List

```python
no_frixion_money_moov_models_payment_request_results = client.NoFrixionMoneyMoovModelsPaymentRequestResult().list()
```


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment

Create an instance: `no_frixion_money_moov_models_payment_requests_merchant_payment = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `description` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `name` | `str` |  |
| `template` | `dict` |  |

#### Example: List

```python
no_frixion_money_moov_models_payment_requests_merchant_payments = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment().list()
```


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2

Create an instance: `no_frixion_money_moov_models_payment_requests_merchant_payment2 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `description` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `name` | `str` |  |
| `template` | `dict` |  |

#### Example: Load

```python
no_frixion_money_moov_models_payment_requests_merchant_payment2 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2().load({"paymentrequest_id": "paymentrequest_id", "template_id": "template_id"})
```


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3

Create an instance: `no_frixion_money_moov_models_payment_requests_merchant_payment3 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3()`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `description` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `name` | `str` |  |
| `template` | `dict` |  |


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4

Create an instance: `no_frixion_money_moov_models_payment_requests_merchant_payment4 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4()`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |


### NoFrixionMoneyMoovModelsPayoutKeysetPage

Create an instance: `no_frixion_money_moov_models_payout_keyset_page = client.NoFrixionMoneyMoovModelsPayoutKeysetPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `str` |  |
| `amount` | `float` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `str` |  |
| `approver_id` | `str` |  |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `str` |  |
| `beneficiary` | `dict` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `str` |  |
| `created_by` | `str` |  |
| `created_by_email_address` | `str` |  |
| `currency` | `str` |  |
| `current_user_id` | `str` |  |
| `description` | `str` |  |
| `destination` | `dict` |  |
| `document` | `list` |  |
| `event` | `list` |  |
| `formatted_amount` | `str` |  |
| `formatted_fx_destination_amount` | `str` |  |
| `formatted_schedule` | `str` |  |
| `formatted_schedule_day_only` | `str` |  |
| `formatted_source_account_available_balance` | `str` |  |
| `fx_destination_amount` | `float` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `str` |  |
| `fx_quote_expires_at` | `str` |  |
| `fx_quote_id` | `str` |  |
| `fx_rate` | `float` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `invoice_id` | `str` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `merchant_token_description` | `str` |  |
| `nonce` | `str` |  |
| `payment_processor` | `str` |  |
| `payment_rail` | `str` |  |
| `payrun_id` | `str` |  |
| `payrun_name` | `str` |  |
| `rule` | `dict` |  |
| `schedule_date` | `str` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `str` |  |
| `source_account_currency` | `str` |  |
| `source_account_iban` | `str` |  |
| `source_account_identifier` | `dict` |  |
| `source_account_name` | `str` |  |
| `source_account_number` | `str` |  |
| `source_account_sortcode` | `str` |  |
| `status` | `str` |  |
| `tag` | `list` |  |
| `their_reference` | `str` |  |
| `topup_payrun_id` | `str` |  |
| `transacted_amount` | `float` |  |
| `transacted_fx_amount` | `float` |  |
| `transacted_fx_rate` | `float` |  |
| `type` | `str` |  |
| `user_id` | `str` |  |
| `your_reference` | `str` |  |

#### Example: List

```python
no_frixion_money_moov_models_payout_keyset_pages = client.NoFrixionMoneyMoovModelsPayoutKeysetPage().list()
```


### NoFrixionMoneyMoovModelsPayoutMetric

Create an instance: `no_frixion_money_moov_models_payout_metric = client.NoFrixionMoneyMoovModelsPayoutMetric()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `all` | `float` |  |
| `failed` | `float` |  |
| `in_progress` | `float` |  |
| `paid` | `float` |  |
| `pending_approval` | `float` |  |
| `scheduled` | `float` |  |
| `total_amounts_by_currency` | `dict` |  |

#### Example: Load

```python
no_frixion_money_moov_models_payout_metric = client.NoFrixionMoneyMoovModelsPayoutMetric().load()
```


### NoFrixionMoneyMoovModelsPayoutsPayoutsCreate

Create an instance: `no_frixion_money_moov_models_payouts_payouts_create = client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_payout` | `dict` |  |
| `payout` | `list` |  |

#### Example: Create

```python
no_frixion_money_moov_models_payouts_payouts_create = client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate().create({
})
```


### NoFrixionMoneyMoovModelsPayrun

Create an instance: `no_frixion_money_moov_models_payrun = client.NoFrixionMoneyMoovModelsPayrun()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation` | `list` |  |
| `authorisation_date` | `str` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `str` |  |
| `can_authorise` | `bool` |  |
| `can_delete` | `bool` |  |
| `can_edit` | `bool` |  |
| `event` | `list` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `invoice` | `list` |  |
| `invoices_minimal` | `list` |  |
| `is_archived` | `bool` |  |
| `last_updated` | `str` |  |
| `last_updated_by` | `dict` |  |
| `merchant_id` | `str` |  |
| `name` | `str` |  |
| `nonce` | `str` |  |
| `payment` | `list` |  |
| `payout` | `list` |  |
| `payouts_count` | `int` |  |
| `reason` | `str` |  |
| `schedule_date` | `str` |  |
| `scheduled_date` | `str` |  |
| `source_account` | `list` |  |
| `status` | `str` |  |
| `total_eur` | `float` |  |
| `total_gbp` | `float` |  |
| `total_usd` | `float` |  |

#### Example: Load

```python
no_frixion_money_moov_models_payrun = client.NoFrixionMoneyMoovModelsPayrun().load({"id": "no_frixion_money_moov_models_payrun_id"})
```

#### Example: Create

```python
no_frixion_money_moov_models_payrun = client.NoFrixionMoneyMoovModelsPayrun().create({
    "id": "example_id",  # str
})
```


### NoFrixionMoneyMoovModelsReportResult

Create an instance: `no_frixion_money_moov_models_report_result = client.NoFrixionMoneyMoovModelsReportResult()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `content` | `str` |  |
| `content_type` | `str` |  |
| `last_completed_at` | `str` |  |
| `merchant_id` | `str` |  |
| `report_name` | `str` |  |
| `report_type` | `str` |  |
| `statement_number` | `int` |  |

#### Example: Load

```python
no_frixion_money_moov_models_report_result = client.NoFrixionMoneyMoovModelsReportResult().load({"id": 1, "report_id": "report_id"})
```


### NoFrixionMoneyMoovModelsRule

Create an instance: `no_frixion_money_moov_models_rule = client.NoFrixionMoneyMoovModelsRule()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account` | `dict` |  |
| `account_id` | `str` |  |
| `approve_url` | `str` |  |
| `approver_id` | `str` |  |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `can_authorise` | `bool` |  |
| `created_by` | `dict` |  |
| `description` | `str` |  |
| `end_at` | `str` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `is_disabled` | `bool` |  |
| `last_executed_at` | `str` |  |
| `last_run_at_transaction_date` | `str` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `name` | `str` |  |
| `nonce` | `str` |  |
| `on_approved_web_hook_url` | `str` |  |
| `on_execution_error_web_hook_url` | `str` |  |
| `on_execution_success_web_hook_url` | `str` |  |
| `start_at` | `str` |  |
| `status` | `str` |  |
| `sweep_action` | `dict` |  |
| `time_zone_id` | `str` |  |
| `trigger_cron_expression` | `str` |  |
| `trigger_on_pay_in` | `bool` |  |
| `user_id` | `str` |  |
| `web_hook_secret` | `str` |  |

#### Example: Load

```python
no_frixion_money_moov_models_rule = client.NoFrixionMoneyMoovModelsRule().load({"id": "no_frixion_money_moov_models_rule_id"})
```

#### Example: Create

```python
no_frixion_money_moov_models_rule = client.NoFrixionMoneyMoovModelsRule().create({
    "created_by": {},  # dict
    "nonce": "example_nonce",  # str
})
```


### NoFrixionMoneyMoovModelsTransaction

Create an instance: `no_frixion_money_moov_models_transaction = client.NoFrixionMoneyMoovModelsTransaction()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `str` |  |
| `account_name` | `str` |  |
| `account_sequence_number` | `int` |  |
| `amount` | `float` |  |
| `amount_minor_unit` | `int` |  |
| `balance` | `float` |  |
| `balance_minor_unit` | `int` |  |
| `counterparty` | `dict` |  |
| `counterparty_summary` | `str` |  |
| `currency` | `str` |  |
| `description` | `str` |  |
| `fx_amount` | `float` |  |
| `fx_currency` | `str` |  |
| `fx_rate` | `float` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `merchant_id` | `str` |  |
| `payment_request_custom_field` | `dict` |  |
| `payment_request_id` | `str` |  |
| `payout_id` | `str` |  |
| `raw_reference` | `str` |  |
| `rule_id` | `str` |  |
| `tag` | `list` |  |
| `their_reference` | `str` |  |
| `transaction_date` | `str` |  |
| `type` | `str` |  |
| `virtual_iban` | `str` |  |
| `your_reference` | `str` |  |

#### Example: Load

```python
no_frixion_money_moov_models_transaction = client.NoFrixionMoneyMoovModelsTransaction().load({"id": "no_frixion_money_moov_models_transaction_id"})
```


### NoFrixionMoneyMoovModelsTransactionPage

Create an instance: `no_frixion_money_moov_models_transaction_page = client.NoFrixionMoneyMoovModelsTransactionPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `str` |  |
| `account_name` | `str` |  |
| `account_sequence_number` | `int` |  |
| `amount` | `float` |  |
| `amount_minor_unit` | `int` |  |
| `balance` | `float` |  |
| `balance_minor_unit` | `int` |  |
| `content` | `list` |  |
| `counterparty` | `dict` |  |
| `counterparty_summary` | `str` |  |
| `currency` | `str` |  |
| `description` | `str` |  |
| `fx_amount` | `float` |  |
| `fx_currency` | `str` |  |
| `fx_rate` | `float` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `merchant_id` | `str` |  |
| `page_number` | `int` |  |
| `page_size` | `int` |  |
| `payment_request_custom_field` | `dict` |  |
| `payment_request_id` | `str` |  |
| `payout_id` | `str` |  |
| `raw_reference` | `str` |  |
| `rule_id` | `str` |  |
| `tag` | `list` |  |
| `their_reference` | `str` |  |
| `total_page` | `int` |  |
| `total_size` | `int` |  |
| `transaction_date` | `str` |  |
| `type` | `str` |  |
| `virtual_iban` | `str` |  |
| `your_reference` | `str` |  |

#### Example: Load

```python
no_frixion_money_moov_models_transaction_page = client.NoFrixionMoneyMoovModelsTransactionPage().load({"account_id": "account_id"})
```

#### Example: List

```python
no_frixion_money_moov_models_transaction_pages = client.NoFrixionMoneyMoovModelsTransactionPage().list()
```


### NoFrixionMoneyMoovModelsUserInvite

Create an instance: `no_frixion_money_moov_models_user_invite = client.NoFrixionMoneyMoovModelsUserInvite()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_status` | `dict` |  |
| `id` | `str` |  |
| `initial_role_id` | `str` |  |
| `invitee_email_address` | `str` |  |
| `invitee_first_name` | `str` |  |
| `invitee_last_name` | `str` |  |
| `inviter_email_address` | `str` |  |
| `inviter_first_name` | `str` |  |
| `inviter_last_name` | `str` |  |
| `is_authorised` | `bool` |  |
| `is_invitee_registered` | `bool` |  |
| `last_invited` | `str` |  |
| `merchant_id` | `str` |  |
| `merchant_name` | `str` |  |
| `message` | `str` |  |
| `registration_url` | `str` |  |
| `send_invite_email` | `bool` |  |
| `status` | `str` |  |
| `user` | `dict` |  |
| `user_id` | `str` |  |

#### Example: Load

```python
no_frixion_money_moov_models_user_invite = client.NoFrixionMoneyMoovModelsUserInvite().load({"id": "no_frixion_money_moov_models_user_invite_id"})
```

#### Example: Create

```python
no_frixion_money_moov_models_user_invite = client.NoFrixionMoneyMoovModelsUserInvite().create({
    "user": {},  # dict
})
```


### NoFrixionMoneyMoovModelsUserInvitePage

Create an instance: `no_frixion_money_moov_models_user_invite_page = client.NoFrixionMoneyMoovModelsUserInvitePage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_status` | `dict` |  |
| `id` | `str` |  |
| `initial_role_id` | `str` |  |
| `invitee_email_address` | `str` |  |
| `invitee_first_name` | `str` |  |
| `invitee_last_name` | `str` |  |
| `inviter_email_address` | `str` |  |
| `inviter_first_name` | `str` |  |
| `inviter_last_name` | `str` |  |
| `is_authorised` | `bool` |  |
| `is_invitee_registered` | `bool` |  |
| `last_invited` | `str` |  |
| `merchant_id` | `str` |  |
| `merchant_name` | `str` |  |
| `message` | `str` |  |
| `registration_url` | `str` |  |
| `status` | `str` |  |
| `user` | `dict` |  |
| `user_id` | `str` |  |

#### Example: List

```python
no_frixion_money_moov_models_user_invite_pages = client.NoFrixionMoneyMoovModelsUserInvitePage().list()
```


### NoFrixionMoneyMoovModelsUserPage

Create an instance: `no_frixion_money_moov_models_user_page = client.NoFrixionMoneyMoovModelsUserPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client_session_timeout` | `list` |  |
| `email_address` | `str` |  |
| `first_name` | `str` |  |
| `id` | `str` |  |
| `last_name` | `str` |  |
| `passkey_added` | `bool` |  |
| `permission` | `dict` |  |
| `roles_with_scope` | `list` |  |
| `two_factor_enabled` | `bool` |  |

#### Example: List

```python
no_frixion_money_moov_models_user_pages = client.NoFrixionMoneyMoovModelsUserPage().list()
```


### NoFrixionMoneyMoovModelsWebhook

Create an instance: `no_frixion_money_moov_models_webhook = client.NoFrixionMoneyMoovModelsWebhook()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `destination_url` | `str` |  |
| `email_address` | `str` |  |
| `failed_notification_email_address` | `str` |  |
| `id` | `str` |  |
| `is_active` | `bool` |  |
| `merchant_id` | `str` |  |
| `notification_method` | `str` |  |
| `resource_type` | `list` |  |
| `retry` | `bool` |  |
| `secret` | `str` |  |
| `version` | `int` |  |

#### Example: Load

```python
no_frixion_money_moov_models_webhook = client.NoFrixionMoneyMoovModelsWebhook().load({"id": "no_frixion_money_moov_models_webhook_id"})
```

#### Example: List

```python
no_frixion_money_moov_models_webhooks = client.NoFrixionMoneyMoovModelsWebhook().list()
```

#### Example: Create

```python
no_frixion_money_moov_models_webhook = client.NoFrixionMoneyMoovModelsWebhook().create({
})
```


### OpenBanking

Create an instance: `open_banking = client.OpenBanking()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```python
open_banking = client.OpenBanking().create({
    "account_id": "example_account_id",  # str
})
```


### Payeeverification

Create an instance: `payeeverification = client.Payeeverification()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `str` |  |
| `account_number` | `str` |  |
| `iban` | `str` |  |
| `payee_verified_account_name` | `str` |  |
| `result` | `str` |  |
| `secondary_identification` | `str` |  |
| `sort_code` | `str` |  |

#### Example: Create

```python
payeeverification = client.Payeeverification().create({
    "account_name": "example_account_name",  # str
    "iban": "example_iban",  # str
})
```


### PaymentRequest

Create an instance: `payment_request = client.PaymentRequest()`

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
| `amount` | `float` |  |
| `do_simulate_settlement_failure` | `bool` |  |
| `error_description` | `str` |  |
| `institution` | `str` |  |
| `payment_initiation_id` | `str` |  |

#### Example: Load

```python
payment_request = client.PaymentRequest().load()
```

#### Example: Create

```python
payment_request = client.PaymentRequest().create({
    "paymentrequest_id": "example_paymentrequest_id",  # str
})
```


### Payout

Create an instance: `payout = client.Payout()`

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
| `account_id` | `str` |  |
| `allow_incomplete` | `bool` |  |
| `amount` | `float` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `str` |  |
| `approver_id` | `str` |  |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `str` |  |
| `beneficiary` | `dict` |  |
| `beneficiary_id` | `str` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `str` |  |
| `created_by` | `str` |  |
| `created_by_email_address` | `str` |  |
| `currency` | `str` |  |
| `current_user_id` | `str` |  |
| `description` | `str` |  |
| `destination` | `dict` |  |
| `document` | `list` |  |
| `event` | `list` |  |
| `formatted_amount` | `str` |  |
| `formatted_fx_destination_amount` | `str` |  |
| `formatted_schedule` | `str` |  |
| `formatted_schedule_day_only` | `str` |  |
| `formatted_source_account_available_balance` | `str` |  |
| `fx_destination_amount` | `float` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `str` |  |
| `fx_quote_expires_at` | `str` |  |
| `fx_quote_id` | `str` |  |
| `fx_rate` | `float` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `invoice_id` | `str` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `merchant_token_description` | `str` |  |
| `nonce` | `str` |  |
| `payment_processor` | `str` |  |
| `payment_rail` | `str` |  |
| `payrun_id` | `str` |  |
| `payrun_name` | `str` |  |
| `rule` | `dict` |  |
| `schedule_date` | `str` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `str` |  |
| `source_account_currency` | `str` |  |
| `source_account_iban` | `str` |  |
| `source_account_identifier` | `dict` |  |
| `source_account_name` | `str` |  |
| `source_account_number` | `str` |  |
| `source_account_sortcode` | `str` |  |
| `status` | `str` |  |
| `tag` | `list` |  |
| `tag_id` | `list` |  |
| `their_reference` | `str` |  |
| `topup_payrun_id` | `str` |  |
| `transacted_amount` | `float` |  |
| `transacted_fx_amount` | `float` |  |
| `transacted_fx_rate` | `float` |  |
| `type` | `str` |  |
| `user_id` | `str` |  |
| `your_reference` | `str` |  |

#### Example: Load

```python
payout = client.Payout().load({"id": "payout_id"})
```

#### Example: Create

```python
payout = client.Payout().create({
})
```


### Payrun

Create an instance: `payrun = client.Payrun()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `str` |  |
| `note` | `str` |  |
| `scheduled_date` | `str` |  |

#### Example: Create

```python
payrun = client.Payrun().create({
    "id": "example_id",  # str
})
```


### Reject

Create an instance: `reject = client.Reject()`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `str` |  |
| `amount` | `float` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `str` |  |
| `approver_id` | `str` |  |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `str` |  |
| `beneficiary` | `dict` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `str` |  |
| `created_by` | `str` |  |
| `created_by_email_address` | `str` |  |
| `currency` | `str` |  |
| `current_user_id` | `str` |  |
| `description` | `str` |  |
| `destination` | `dict` |  |
| `document` | `list` |  |
| `event` | `list` |  |
| `formatted_amount` | `str` |  |
| `formatted_fx_destination_amount` | `str` |  |
| `formatted_schedule` | `str` |  |
| `formatted_schedule_day_only` | `str` |  |
| `formatted_source_account_available_balance` | `str` |  |
| `fx_destination_amount` | `float` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `str` |  |
| `fx_quote_expires_at` | `str` |  |
| `fx_quote_id` | `str` |  |
| `fx_rate` | `float` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `invoice_id` | `str` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `merchant_token_description` | `str` |  |
| `nonce` | `str` |  |
| `payment_processor` | `str` |  |
| `payment_rail` | `str` |  |
| `payrun_id` | `str` |  |
| `payrun_name` | `str` |  |
| `reason` | `str` |  |
| `rule` | `dict` |  |
| `schedule_date` | `str` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `str` |  |
| `source_account_currency` | `str` |  |
| `source_account_iban` | `str` |  |
| `source_account_identifier` | `dict` |  |
| `source_account_name` | `str` |  |
| `source_account_number` | `str` |  |
| `source_account_sortcode` | `str` |  |
| `status` | `str` |  |
| `tag` | `list` |  |
| `their_reference` | `str` |  |
| `topup_payrun_id` | `str` |  |
| `transacted_amount` | `float` |  |
| `transacted_fx_amount` | `float` |  |
| `transacted_fx_rate` | `float` |  |
| `type` | `str` |  |
| `user_id` | `str` |  |
| `your_reference` | `str` |  |


### Report

Create an instance: `report = client.Report()`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |


### Rule

Create an instance: `rule = client.Rule()`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |


### Send

Create an instance: `send = client.Send()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `str` |  |
| `allow_incomplete` | `bool` |  |
| `amount` | `float` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `str` |  |
| `approver_id` | `str` |  |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `str` |  |
| `beneficiary` | `dict` |  |
| `beneficiary_id` | `str` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `str` |  |
| `created_by` | `str` |  |
| `created_by_email_address` | `str` |  |
| `currency` | `str` |  |
| `current_user_id` | `str` |  |
| `description` | `str` |  |
| `destination` | `dict` |  |
| `document` | `list` |  |
| `event` | `list` |  |
| `formatted_amount` | `str` |  |
| `formatted_fx_destination_amount` | `str` |  |
| `formatted_schedule` | `str` |  |
| `formatted_schedule_day_only` | `str` |  |
| `formatted_source_account_available_balance` | `str` |  |
| `fx_destination_amount` | `float` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `str` |  |
| `fx_quote_expires_at` | `str` |  |
| `fx_quote_id` | `str` |  |
| `fx_rate` | `float` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `invoice_id` | `str` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `merchant_token_description` | `str` |  |
| `nonce` | `str` |  |
| `payment_processor` | `str` |  |
| `payment_rail` | `str` |  |
| `payrun_id` | `str` |  |
| `payrun_name` | `str` |  |
| `rule` | `dict` |  |
| `schedule_date` | `str` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `str` |  |
| `source_account_currency` | `str` |  |
| `source_account_iban` | `str` |  |
| `source_account_identifier` | `dict` |  |
| `source_account_name` | `str` |  |
| `source_account_number` | `str` |  |
| `source_account_sortcode` | `str` |  |
| `status` | `str` |  |
| `tag` | `list` |  |
| `tag_id` | `list` |  |
| `their_reference` | `str` |  |
| `topup_payrun_id` | `str` |  |
| `transacted_amount` | `float` |  |
| `transacted_fx_amount` | `float` |  |
| `transacted_fx_rate` | `float` |  |
| `type` | `str` |  |
| `user_id` | `str` |  |
| `your_reference` | `str` |  |

#### Example: Create

```python
send = client.Send().create({
    "beneficiary": {},  # dict
    "source_account_identifier": {},  # dict
})
```


### Sendbeneficiary

Create an instance: `sendbeneficiary = client.Sendbeneficiary()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `str` |  |
| `allow_incomplete` | `bool` |  |
| `amount` | `float` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `str` |  |
| `approver_id` | `str` |  |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `str` |  |
| `beneficiary` | `dict` |  |
| `beneficiary_id` | `str` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `str` |  |
| `created_by` | `str` |  |
| `created_by_email_address` | `str` |  |
| `currency` | `str` |  |
| `current_user_id` | `str` |  |
| `description` | `str` |  |
| `destination` | `dict` |  |
| `document` | `list` |  |
| `event` | `list` |  |
| `formatted_amount` | `str` |  |
| `formatted_fx_destination_amount` | `str` |  |
| `formatted_schedule` | `str` |  |
| `formatted_schedule_day_only` | `str` |  |
| `formatted_source_account_available_balance` | `str` |  |
| `fx_destination_amount` | `float` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `str` |  |
| `fx_quote_expires_at` | `str` |  |
| `fx_quote_id` | `str` |  |
| `fx_rate` | `float` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `invoice_id` | `str` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `merchant_token_description` | `str` |  |
| `nonce` | `str` |  |
| `payment_processor` | `str` |  |
| `payment_rail` | `str` |  |
| `payrun_id` | `str` |  |
| `payrun_name` | `str` |  |
| `rule` | `dict` |  |
| `schedule_date` | `str` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `str` |  |
| `source_account_currency` | `str` |  |
| `source_account_iban` | `str` |  |
| `source_account_identifier` | `dict` |  |
| `source_account_name` | `str` |  |
| `source_account_number` | `str` |  |
| `source_account_sortcode` | `str` |  |
| `status` | `str` |  |
| `tag` | `list` |  |
| `tag_id` | `list` |  |
| `their_reference` | `str` |  |
| `topup_payrun_id` | `str` |  |
| `transacted_amount` | `float` |  |
| `transacted_fx_amount` | `float` |  |
| `transacted_fx_rate` | `float` |  |
| `type` | `str` |  |
| `user_id` | `str` |  |
| `your_reference` | `str` |  |

#### Example: Create

```python
sendbeneficiary = client.Sendbeneficiary().create({
    "beneficiary": {},  # dict
    "source_account_identifier": {},  # dict
})
```


### Tag

Create an instance: `tag = client.Tag()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `colour_hex` | `str` |  |
| `description` | `str` |  |
| `id` | `str` |  |
| `merchant_id` | `str` |  |
| `name` | `str` |  |

#### Example: List

```python
tags = client.Tag().list()
```

#### Example: Create

```python
tag = client.Tag().create({
    "merchant_id": "example_merchant_id",  # str
})
```


### Token

Create an instance: `token = client.Token()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```python
token = client.Token().create({
    "id": "example_id",  # str
})
```


### Transaction

Create an instance: `transaction = client.Transaction()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Example: Load

```python
transaction = client.Transaction().load({"sequence_number": 1, "transaction_id": "transaction_id"})
```

#### Example: Create

```python
transaction = client.Transaction().create({
    "id": "example_id",  # str
})
```


### User

Create an instance: `user = client.User()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client_session_timeout` | `list` |  |
| `email_address` | `str` |  |
| `first_name` | `str` |  |
| `id` | `str` |  |
| `last_name` | `str` |  |
| `passkey_added` | `bool` |  |
| `permission` | `dict` |  |
| `profile` | `str` |  |
| `roles_with_scope` | `list` |  |
| `two_factor_enabled` | `bool` |  |
| `user_invite_id` | `str` |  |

#### Example: List

```python
users = client.User().list()
```


### UserInvite

Create an instance: `user_invite = client.UserInvite()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Example: Create

```python
user_invite = client.UserInvite().create({
    "id": "example_id",  # str
})
```


### Virtual

Create an instance: `virtual = client.Virtual()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `str` |  |
| `account_supplier_name` | `str` |  |
| `available_balance` | `float` |  |
| `available_balance_minor_unit` | `int` |  |
| `balance` | `float` |  |
| `balance_minor_unit` | `int` |  |
| `bank_name` | `str` |  |
| `consent_id` | `str` |  |
| `created_by` | `dict` |  |
| `created_by_display_name` | `str` |  |
| `currency` | `str` |  |
| `default_payment_rail` | `str` |  |
| `display_name` | `str` |  |
| `expiry_date` | `str` |  |
| `external_account_icon` | `str` |  |
| `id` | `str` |  |
| `identifier` | `dict` |  |
| `inserted` | `str` |  |
| `is_archived` | `bool` |  |
| `is_connected_account` | `bool` |  |
| `is_default` | `bool` |  |
| `is_trust_account` | `bool` |  |
| `is_virtual` | `bool` |  |
| `last_transaction` | `dict` |  |
| `last_updated` | `str` |  |
| `merchant_id` | `str` |  |
| `merchant_name` | `str` |  |
| `name` | `str` |  |
| `physical_account_id` | `str` |  |
| `rule` | `list` |  |
| `submitted_payouts_balance` | `float` |  |
| `submitted_payouts_balance_minor_unit` | `int` |  |
| `summary` | `str` |  |
| `supplier_sepa_instant_status` | `str` |  |
| `xero_bank_feed_connection_status` | `str` |  |
| `xero_bank_feed_last_synced_at` | `str` |  |
| `xero_bank_feed_sync_last_failed_at` | `str` |  |
| `xero_bank_feed_sync_last_failure_reason` | `str` |  |
| `xero_bank_feed_sync_status` | `str` |  |
| `xero_unsynchronised_transactions_count` | `int` |  |

#### Example: Create

```python
virtual = client.Virtual().create({
    "account_id": "example_account_id",  # str
})
```


### Webhook

Create an instance: `webhook = client.Webhook()`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |


### Whoami

Create an instance: `whoami = client.Whoami()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client_session_timeout` | `list` |  |
| `email_address` | `str` |  |
| `first_name` | `str` |  |
| `id` | `str` |  |
| `last_name` | `str` |  |
| `passkey_added` | `bool` |  |
| `permission` | `dict` |  |
| `roles_with_scope` | `list` |  |
| `two_factor_enabled` | `bool` |  |

#### Example: List

```python
whoamis = client.Whoami().list()
```


### Whoamitrustedapp

Create an instance: `whoamitrustedapp = client.Whoamitrustedapp()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client_session_timeout` | `list` |  |
| `email_address` | `str` |  |
| `first_name` | `str` |  |
| `id` | `str` |  |
| `last_name` | `str` |  |
| `passkey_added` | `bool` |  |
| `permission` | `dict` |  |
| `roles_with_scope` | `list` |  |
| `two_factor_enabled` | `bool` |  |

#### Example: List

```python
whoamitrustedapps = client.Whoamitrustedapp().list()
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

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── nofrixion_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`nofrixion_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```python
account = client.Account()
account.list()

# account.data_get() now returns the account data from the last list
# account.match_get() returns the last match criteria
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
