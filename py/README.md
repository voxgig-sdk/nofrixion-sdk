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

### 3. Load a cardcustomertoken

CardCustomerToken is nested under customer_email_address, so provide the `customer_email_address`.
`load()` returns the bare record (a `dict`) and raises on error.

```python
try:
    cardcustomertoken = client.CardCustomerToken().load({"customer_email_address": "example_customer_email_address"})
    print(cardcustomertoken)
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
| `MerchantDirectDebitMandate` | `(data) -> MerchantDirectDebitMandateEntity` | Create a MerchantDirectDebitMandate entity instance. |
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
| `PayoutKeyset` | `(data) -> PayoutKeysetEntity` | Create a PayoutKeyset entity instance. |
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

#### MerchantDirectDebitMandate

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

#### PayoutKeyset

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
| `account_balance` | `list` |  |
| `account_id` | `str` |  |
| `account_identification` | `list` |  |
| `account_name` | `str` |  |
| `account_supplier_name` | `str` |  |
| `account_type` | `str` |  |
| `available_balance` | `float` |  |
| `available_balance_minor_unit` | `int` |  |
| `balance` | `float` |  |
| `balance_minor_unit` | `int` |  |
| `bank_name` | `str` |  |
| `consent_id` | `str` |  |
| `consolidated_account_information` | `dict` |  |
| `created_by` | `dict` |  |
| `created_by_display_name` | `str` |  |
| `currency` | `str` |  |
| `default_payment_rail` | `str` |  |
| `description` | `str` |  |
| `detail` | `str` |  |
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
| `nickname` | `str` |  |
| `physical_account_id` | `str` |  |
| `role_i_d` | `list` |  |
| `rule` | `list` |  |
| `submitted_payouts_balance` | `float` |  |
| `submitted_payouts_balance_minor_unit` | `int` |  |
| `summary` | `str` |  |
| `supplier_physical_account_id` | `str` |  |
| `supplier_sepa_instant_status` | `str` |  |
| `to_date` | `str` |  |
| `type` | `str` |  |
| `usage_type` | `str` |  |
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


### Batch

Create an instance: `batch = client.Batch()`

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
batch = client.Batch().load({"id": "batch_id"})
```

#### Example: Create

```python
batch = client.Batch().create({
})
```


### Beneficiary

Create an instance: `beneficiary = client.Beneficiary()`

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
| `approval_callback_url` | `str` |  |
| `authentication_method` | `list` |  |
| `authorisation` | `list` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `beneficiary` | `list` |  |
| `beneficiary_event` | `list` |  |
| `can_authorise` | `bool` |  |
| `can_update` | `bool` |  |
| `created_by` | `dict` |  |
| `created_by_email_address` | `str` |  |
| `currency` | `str` |  |
| `destination` | `dict` |  |
| `failed_beneficiary` | `dict` |  |
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

#### Example: List

```python
beneficiarys = client.Beneficiary().list()
```

#### Example: Create

```python
beneficiary = client.Beneficiary().create({
})
```


### BeneficiaryGroup

Create an instance: `beneficiary_group = client.BeneficiaryGroup()`

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
beneficiary_groups = client.BeneficiaryGroup().list()
```


### Card

Create an instance: `card = client.Card()`

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
card = client.Card().create({
    "paymentrequest_id": "example_paymentrequest_id",  # str
})
```


### CardCustomerToken

Create an instance: `card_customer_token = client.CardCustomerToken()`

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
card_customer_token = client.CardCustomerToken().load({"customer_email_address": "customer_email_address"})
```

#### Example: List

```python
card_customer_tokens = client.CardCustomerToken().list()
```


### CardPayment

Create an instance: `card_payment = client.CardPayment()`

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
card_payment = client.CardPayment().create({
    "paymentrequest_id": "example_paymentrequest_id",  # str
})
```


### CardPublicKey

Create an instance: `card_public_key = client.CardPublicKey()`

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
card_public_key = client.CardPublicKey().load({"paymentrequest_id": "paymentrequest_id"})
```


### Consent

Create an instance: `consent = client.Consent()`

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
consent = client.Consent().load({"id": "consent_id"})
```

#### Example: List

```python
consents = client.Consent().list()
```

#### Example: Create

```python
consent = client.Consent().create({
})
```


### Currency

Create an instance: `currency = client.Currency()`

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
currencys = client.Currency().list()
```


### DirectDebitBatchSubmit

Create an instance: `direct_debit_batch_submit = client.DirectDebitBatchSubmit()`

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
direct_debit_batch_submit = client.DirectDebitBatchSubmit().create({
})
```


### FxRate

Create an instance: `fx_rate = client.FxRate()`

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
fx_rate = client.FxRate().load({"destination": "destination", "source": "source", "valid_for_minute": 1})
```

#### Example: List

```python
fx_rates = client.FxRate().list()
```


### IPayment

Create an instance: `i_payment = client.IPayment()`

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
i_payment = client.IPayment().create({
})
```


### Mandate

Create an instance: `mandate = client.Mandate()`

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
mandate = client.Mandate().load({"id": "mandate_id"})
```

#### Example: Create

```python
mandate = client.Mandate().create({
    "address_line1": "example_address_line1",  # str
    "city": "example_city",  # str
    "country_code": "example_country_code",  # str
    "email_address": "example_email_address",  # str
    "first_name": "example_first_name",  # str
    "last_name": "example_last_name",  # str
    "postal_code": "example_postal_code",  # str
})
```


### Merchant

Create an instance: `merchant = client.Merchant()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
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
| `reason` | `str` |  |
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
merchant = client.Merchant().load({"id": "merchant_id"})
```

#### Example: List

```python
merchants = client.Merchant().list()
```


### MerchantAuthorisationSetting

Create an instance: `merchant_authorisation_setting = client.MerchantAuthorisationSetting()`

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
merchant_authorisation_settings = client.MerchantAuthorisationSetting().list()
```


### MerchantDirectDebitMandate

Create an instance: `merchant_direct_debit_mandate = client.MerchantDirectDebitMandate()`

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
merchant_direct_debit_mandates = client.MerchantDirectDebitMandate().list()
```


### MerchantPayByBankSetting

Create an instance: `merchant_pay_by_bank_setting = client.MerchantPayByBankSetting()`

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
merchant_pay_by_bank_settings = client.MerchantPayByBankSetting().list()
```


### MerchantPaymentRequestTemplate

Create an instance: `merchant_payment_request_template = client.MerchantPaymentRequestTemplate()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
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

#### Example: Load

```python
merchant_payment_request_template = client.MerchantPaymentRequestTemplate().load({"id": "merchant_payment_request_template_id", "paymentrequest_id": "paymentrequest_id"})
```

#### Example: List

```python
merchant_payment_request_templates = client.MerchantPaymentRequestTemplate().list()
```


### MerchantToken

Create an instance: `merchant_token = client.MerchantToken()`

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
merchant_token = client.MerchantToken().load({"id": "merchant_token_id"})
```

#### Example: List

```python
merchant_tokens = client.MerchantToken().list()
```

#### Example: Create

```python
merchant_token = client.MerchantToken().create({
    "nonce": "example_nonce",  # str
})
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


### NoFrixionVersion

Create an instance: `no_frixion_version = client.NoFrixionVersion()`

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
no_frixion_version = client.NoFrixionVersion().load()
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


### Payment

Create an instance: `payment = client.Payment()`

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
payment = client.Payment().load({"id": "payment_id"})
```

#### Example: Create

```python
payment = client.Payment().create({
    "created_by_user": {},  # dict
})
```


### PaymentAccount

Create an instance: `payment_account = client.PaymentAccount()`

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
payment_accounts = client.PaymentAccount().list()
```


### PaymentAccountMinimal

Create an instance: `payment_account_minimal = client.PaymentAccountMinimal()`

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
payment_account_minimals = client.PaymentAccountMinimal().list()
```


### PaymentInitiation

Create an instance: `payment_initiation = client.PaymentInitiation()`

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
payment_initiation = client.PaymentInitiation().create({
    "paymentrequest_id": "example_paymentrequest_id",  # str
})
```


### PaymentRequest

Create an instance: `payment_request = client.PaymentRequest()`

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
| `do_simulate_settlement_failure` | `bool` |  |
| `due_date` | `str` |  |
| `error_description` | `str` |  |
| `event` | `list` |  |
| `failed_payment_request` | `dict` |  |
| `failure_callback_url` | `str` |  |
| `field_display_setting` | `list` |  |
| `formatted_amount` | `str` |  |
| `hosted_pay_checkout_url` | `str` |  |
| `id` | `str` |  |
| `ignore_address_verification` | `bool` |  |
| `inserted` | `str` |  |
| `inserted_sortable` | `str` |  |
| `institution` | `str` |  |
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
| `payment_initiation_id` | `str` |  |
| `payment_method` | `list` |  |
| `payment_processor` | `str` |  |
| `payment_request` | `list` |  |
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

#### Example: Load

```python
payment_request = client.PaymentRequest().load()
```

#### Example: List

```python
payment_requests = client.PaymentRequest().list()
```

#### Example: Create

```python
payment_request = client.PaymentRequest().create({
})
```


### PaymentRequestEvent

Create an instance: `payment_request_event = client.PaymentRequestEvent()`

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
payment_request_events = client.PaymentRequestEvent().list()
```


### PaymentRequestMetric

Create an instance: `payment_request_metric = client.PaymentRequestMetric()`

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
payment_request_metric = client.PaymentRequestMetric().load()
```


### PaymentRequestMinimal

Create an instance: `payment_request_minimal = client.PaymentRequestMinimal()`

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
payment_request_minimals = client.PaymentRequestMinimal().list()
```


### PaymentRequestResult

Create an instance: `payment_request_result = client.PaymentRequestResult()`

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
payment_request_results = client.PaymentRequestResult().list()
```


### Payout

Create an instance: `payout = client.Payout()`

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
| `failed_payout` | `dict` |  |
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
| `payout` | `list` |  |
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

#### Example: List

```python
payouts = client.Payout().list()
```

#### Example: Create

```python
payout = client.Payout().create({
})
```


### PayoutKeyset

Create an instance: `payout_keyset = client.PayoutKeyset()`

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
payout_keysets = client.PayoutKeyset().list()
```


### PayoutMetric

Create an instance: `payout_metric = client.PayoutMetric()`

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
payout_metric = client.PayoutMetric().load()
```


### Payrun

Create an instance: `payrun = client.Payrun()`

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
| `note` | `str` |  |
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
payrun = client.Payrun().load({"id": "payrun_id"})
```

#### Example: List

```python
payruns = client.Payrun().list()
```

#### Example: Create

```python
payrun = client.Payrun().create({
    "id": "example_id",  # str
})
```


### Report

Create an instance: `report = client.Report()`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |


### ReportResult

Create an instance: `report_result = client.ReportResult()`

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
report_result = client.ReportResult().load({"id": 1, "report_id": "report_id"})
```


### Role

Create an instance: `role = client.Role()`

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
role = client.Role().create({
    "merchant_id": "example_merchant_id",  # str
})
```


### Rule

Create an instance: `rule = client.Rule()`

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
rule = client.Rule().load({"id": "rule_id"})
```

#### Example: List

```python
rules = client.Rule().list()
```

#### Example: Create

```python
rule = client.Rule().create({
    "created_by": {},  # dict
    "nonce": "example_nonce",  # str
})
```


### RuleEvent

Create an instance: `rule_event = client.RuleEvent()`

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
rule_events = client.RuleEvent().list()
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
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `str` |  |
| `account_name` | `str` |  |
| `account_sequence_number` | `int` |  |
| `address_detail` | `dict` |  |
| `amount` | `float` |  |
| `amount_minor_unit` | `int` |  |
| `balance` | `float` |  |
| `balance_minor_unit` | `int` |  |
| `booking_date_time` | `str` |  |
| `charge_detail` | `dict` |  |
| `content` | `list` |  |
| `counterparty` | `dict` |  |
| `counterparty_summary` | `str` |  |
| `currency` | `str` |  |
| `currency_exchange` | `dict` |  |
| `date` | `str` |  |
| `description` | `str` |  |
| `enrichment` | `dict` |  |
| `fx_amount` | `float` |  |
| `fx_currency` | `str` |  |
| `fx_rate` | `float` |  |
| `gross_amount` | `dict` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `iso_bank_transaction_code` | `dict` |  |
| `merchant` | `dict` |  |
| `merchant_id` | `str` |  |
| `page_number` | `int` |  |
| `page_size` | `int` |  |
| `payee_detail` | `dict` |  |
| `payer_detail` | `dict` |  |
| `payment_request_custom_field` | `dict` |  |
| `payment_request_id` | `str` |  |
| `payout_id` | `str` |  |
| `proprietary_bank_transaction_code` | `dict` |  |
| `raw_reference` | `str` |  |
| `reference` | `str` |  |
| `rule_id` | `str` |  |
| `statement_reference` | `list` |  |
| `status` | `str` |  |
| `supplementary_data` | `Any` |  |
| `tag` | `list` |  |
| `their_reference` | `str` |  |
| `total_page` | `int` |  |
| `total_size` | `int` |  |
| `transaction_amount` | `dict` |  |
| `transaction_date` | `str` |  |
| `transaction_information` | `list` |  |
| `transaction_mutability` | `str` |  |
| `type` | `str` |  |
| `value_date_time` | `str` |  |
| `virtual_iban` | `str` |  |
| `your_reference` | `str` |  |

#### Example: Load

```python
transaction = client.Transaction().load({"id": "transaction_id"})
```

#### Example: List

```python
transactions = client.Transaction().list()
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
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_status` | `dict` |  |
| `failed_user_invite` | `dict` |  |
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
| `user_invite` | `list` |  |

#### Example: Load

```python
user_invite = client.UserInvite().load({"id": "user_invite_id"})
```

#### Example: List

```python
user_invites = client.UserInvite().list()
```

#### Example: Create

```python
user_invite = client.UserInvite().create({
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
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
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
webhook = client.Webhook().load({"id": "webhook_id"})
```

#### Example: List

```python
webhooks = client.Webhook().list()
```

#### Example: Create

```python
webhook = client.Webhook().create({
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
