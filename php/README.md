# Nofrixion PHP SDK



The PHP SDK for the Nofrixion API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->Account()` — with named operations (`list`/`load`/`create`/`update`/`remove`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/nofrixion-sdk/releases](https://github.com/voxgig-sdk/nofrixion-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'nofrixion_sdk.php';

$client = new NofrixionSDK([
    "apikey" => getenv("NOFRIXION_APIKEY"),
]);
```

### 2. List account records

```php
try {
    // list() returns an array of Account records — iterate directly.
    $accounts = $client->Account()->list();
    foreach ($accounts as $item) {
        echo $item["id"] . " " . $item["account_balance"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 3. Load a cardcustomertoken

CardCustomerToken is nested under customer_email_address, so provide the `customer_email_address`.

```php
try {
    // load() returns the bare CardCustomerToken record (throws on error).
    $cardcustomertoken = $client->CardCustomerToken()->load(["customer_email_address" => "example_customer_email_address"]);
    print_r($cardcustomertoken);
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 4. Create, update, and remove

```php
// create() returns the bare created Account record.
$created = $client->Account()->create(["account_id" => "example_account_id", "currency" => "example_currency"]);

// Update — index the bare record directly ($created["id"]).
$client->Account()->update(["id" => $created["id"], "account_id" => "example_account_id", "amount" => 1]);

// Remove
$client->Account()->remove(["id" => $created["id"]]);
```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $accounts = $client->Account()->list();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```php
$client = NofrixionSDK::test([
    "entity" => ["account" => ["test01" => ["id" => "test01"]]],
]);

// Entity ops return the bare mock record (throws on error).
$account = $client->Account()->list();
print_r($account);
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new NofrixionSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
NOFRIXION_TEST_LIVE=TRUE
NOFRIXION_APIKEY=<your-key>
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### NofrixionSDK

```php
require_once 'nofrixion_sdk.php';
$client = new NofrixionSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = NofrixionSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### NofrixionSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `Account` | `($data): AccountEntity` | Create an Account entity instance. |
| `Batch` | `($data): BatchEntity` | Create a Batch entity instance. |
| `BeneficiariesCreate` | `($data): BeneficiariesCreateEntity` | Create a BeneficiariesCreate entity instance. |
| `Beneficiary` | `($data): BeneficiaryEntity` | Create a Beneficiary entity instance. |
| `BeneficiaryGroup` | `($data): BeneficiaryGroupEntity` | Create a BeneficiaryGroup entity instance. |
| `Card` | `($data): CardEntity` | Create a Card entity instance. |
| `CardCustomerToken` | `($data): CardCustomerTokenEntity` | Create a CardCustomerToken entity instance. |
| `CardPayment` | `($data): CardPaymentEntity` | Create a CardPayment entity instance. |
| `CardPublicKey` | `($data): CardPublicKeyEntity` | Create a CardPublicKey entity instance. |
| `Consent` | `($data): ConsentEntity` | Create a Consent entity instance. |
| `Currency` | `($data): CurrencyEntity` | Create a Currency entity instance. |
| `DirectDebitBatchSubmit` | `($data): DirectDebitBatchSubmitEntity` | Create a DirectDebitBatchSubmit entity instance. |
| `FxRate` | `($data): FxRateEntity` | Create a FxRate entity instance. |
| `IPayment` | `($data): IPaymentEntity` | Create an IPayment entity instance. |
| `Mandate` | `($data): MandateEntity` | Create a Mandate entity instance. |
| `Merchant` | `($data): MerchantEntity` | Create a Merchant entity instance. |
| `MerchantAuthorisationSetting` | `($data): MerchantAuthorisationSettingEntity` | Create a MerchantAuthorisationSetting entity instance. |
| `MerchantDirectDebitMandate` | `($data): MerchantDirectDebitMandateEntity` | Create a MerchantDirectDebitMandate entity instance. |
| `MerchantPayByBankSetting` | `($data): MerchantPayByBankSettingEntity` | Create a MerchantPayByBankSetting entity instance. |
| `MerchantPaymentRequestTemplate` | `($data): MerchantPaymentRequestTemplateEntity` | Create a MerchantPaymentRequestTemplate entity instance. |
| `MerchantToken` | `($data): MerchantTokenEntity` | Create a MerchantToken entity instance. |
| `Metadata` | `($data): MetadataEntity` | Create a Metadata entity instance. |
| `NoFrixionVersion` | `($data): NoFrixionVersionEntity` | Create a NoFrixionVersion entity instance. |
| `OpenBanking` | `($data): OpenBankingEntity` | Create an OpenBanking entity instance. |
| `Payeeverification` | `($data): PayeeverificationEntity` | Create a Payeeverification entity instance. |
| `Payment` | `($data): PaymentEntity` | Create a Payment entity instance. |
| `PaymentAccount` | `($data): PaymentAccountEntity` | Create a PaymentAccount entity instance. |
| `PaymentAccountMinimal` | `($data): PaymentAccountMinimalEntity` | Create a PaymentAccountMinimal entity instance. |
| `PaymentInitiation` | `($data): PaymentInitiationEntity` | Create a PaymentInitiation entity instance. |
| `PaymentRequest` | `($data): PaymentRequestEntity` | Create a PaymentRequest entity instance. |
| `PaymentRequestEvent` | `($data): PaymentRequestEventEntity` | Create a PaymentRequestEvent entity instance. |
| `PaymentRequestMetric` | `($data): PaymentRequestMetricEntity` | Create a PaymentRequestMetric entity instance. |
| `PaymentRequestMinimal` | `($data): PaymentRequestMinimalEntity` | Create a PaymentRequestMinimal entity instance. |
| `PaymentRequestResult` | `($data): PaymentRequestResultEntity` | Create a PaymentRequestResult entity instance. |
| `PaymentRequestsCreate` | `($data): PaymentRequestsCreateEntity` | Create a PaymentRequestsCreate entity instance. |
| `Payout` | `($data): PayoutEntity` | Create a Payout entity instance. |
| `PayoutKeyset` | `($data): PayoutKeysetEntity` | Create a PayoutKeyset entity instance. |
| `PayoutMetric` | `($data): PayoutMetricEntity` | Create a PayoutMetric entity instance. |
| `PayoutsCreate` | `($data): PayoutsCreateEntity` | Create a PayoutsCreate entity instance. |
| `Payrun` | `($data): PayrunEntity` | Create a Payrun entity instance. |
| `Report` | `($data): ReportEntity` | Create a Report entity instance. |
| `ReportResult` | `($data): ReportResultEntity` | Create a ReportResult entity instance. |
| `RolesCreate` | `($data): RolesCreateEntity` | Create a RolesCreate entity instance. |
| `Rule` | `($data): RuleEntity` | Create a Rule entity instance. |
| `RuleEvent` | `($data): RuleEventEntity` | Create a RuleEvent entity instance. |
| `Tag` | `($data): TagEntity` | Create a Tag entity instance. |
| `Token` | `($data): TokenEntity` | Create a Token entity instance. |
| `Transaction` | `($data): TransactionEntity` | Create a Transaction entity instance. |
| `User` | `($data): UserEntity` | Create an User entity instance. |
| `UserInvite` | `($data): UserInviteEntity` | Create an UserInvite entity instance. |
| `UserInvitesCreate` | `($data): UserInvitesCreateEntity` | Create an UserInvitesCreate entity instance. |
| `Virtual` | `($data): VirtualEntity` | Create a Virtual entity instance. |
| `Webhook` | `($data): WebhookEntity` | Create a Webhook entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `create` | `($reqdata, $ctrl): array` | Create a new entity. |
| `update` | `($reqdata, $ctrl): array` | Update an existing entity. |
| `remove` | `($reqmatch, $ctrl): array` | Remove an entity. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the bare result data (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

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

#### BeneficiariesCreate

| Field | Description |
| --- | --- |
| `beneficiary` |  |
| `failed_beneficiary` |  |

Operations: Create.

API path: `/api/v1/beneficiaries/batchcreate`

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

#### PaymentRequestsCreate

| Field | Description |
| --- | --- |
| `failed_payment_request` |  |
| `payment_request` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/batchcreate`

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

#### PayoutsCreate

| Field | Description |
| --- | --- |
| `failed_payout` |  |
| `payout` |  |

Operations: Create.

API path: `/api/v1/payouts/batchcreate`

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

#### RolesCreate

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

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/userinvites/authorise/{id}`

#### UserInvitesCreate

| Field | Description |
| --- | --- |
| `failed_user_invite` |  |
| `user_invite` |  |

Operations: Create.

API path: `/api/v1/userinvites/batchcreate`

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

Create an instance: `$account = $client->Account();`

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
| `account_balance` | `array` |  |
| `account_id` | `string` |  |
| `account_identification` | `array` |  |
| `account_name` | `string` |  |
| `account_supplier_name` | `string` |  |
| `account_type` | `string` |  |
| `available_balance` | `float` |  |
| `available_balance_minor_unit` | `int` |  |
| `balance` | `float` |  |
| `balance_minor_unit` | `int` |  |
| `bank_name` | `string` |  |
| `consent_id` | `string` |  |
| `consolidated_account_information` | `array` |  |
| `created_by` | `array` |  |
| `created_by_display_name` | `string` |  |
| `currency` | `string` |  |
| `default_payment_rail` | `string` |  |
| `description` | `string` |  |
| `detail` | `string` |  |
| `display_name` | `string` |  |
| `expiry_date` | `string` |  |
| `external_account_icon` | `string` |  |
| `format` | `string` |  |
| `from_date` | `string` |  |
| `id` | `string` |  |
| `identifier` | `array` |  |
| `inserted` | `string` |  |
| `is_archived` | `bool` |  |
| `is_connected_account` | `bool` |  |
| `is_default` | `bool` |  |
| `is_trust_account` | `bool` |  |
| `is_virtual` | `bool` |  |
| `last_transaction` | `array` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `nickname` | `string` |  |
| `physical_account_id` | `string` |  |
| `role_i_d` | `array` |  |
| `rule` | `array` |  |
| `submitted_payouts_balance` | `float` |  |
| `submitted_payouts_balance_minor_unit` | `int` |  |
| `summary` | `string` |  |
| `supplier_physical_account_id` | `string` |  |
| `supplier_sepa_instant_status` | `string` |  |
| `to_date` | `string` |  |
| `type` | `string` |  |
| `usage_type` | `string` |  |
| `xero_bank_feed_connection_status` | `string` |  |
| `xero_bank_feed_last_synced_at` | `string` |  |
| `xero_bank_feed_sync_last_failed_at` | `string` |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` |  |
| `xero_bank_feed_sync_status` | `string` |  |
| `xero_unsynchronised_transactions_count` | `int` |  |

#### Example: Load

```php
// load() returns the bare Account record (throws on error).
$account = $client->Account()->load(["id" => "account_id"]);
```

#### Example: List

```php
// list() returns an array of Account records (throws on error).
$accounts = $client->Account()->list();
```

#### Example: Create

```php
$account = $client->Account()->create([
]);
```


### Batch

Create an instance: `$batch = $client->Batch();`

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
| `payout` | `array` |  |

#### Example: Load

```php
// load() returns the bare Batch record (throws on error).
$batch = $client->Batch()->load(["id" => "batch_id"]);
```

#### Example: Create

```php
$batch = $client->Batch()->create([
]);
```


### BeneficiariesCreate

Create an instance: `$beneficiaries_create = $client->BeneficiariesCreate();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `beneficiary` | `array` |  |
| `failed_beneficiary` | `array` |  |

#### Example: Create

```php
$beneficiaries_create = $client->BeneficiariesCreate()->create([
]);
```


### Beneficiary

Create an instance: `$beneficiary = $client->Beneficiary();`

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
| `approval_callback_url` | `string` |  |
| `authentication_method` | `array` |  |
| `authorisation` | `array` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `beneficiary_event` | `array` |  |
| `can_authorise` | `bool` |  |
| `can_update` | `bool` |  |
| `created_by` | `array` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `destination` | `array` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_enabled` | `bool` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `source_account` | `array` |  |
| `source_account_i_d` | `array` |  |
| `their_reference` | `string` |  |

#### Example: Load

```php
// load() returns the bare Beneficiary record (throws on error).
$beneficiary = $client->Beneficiary()->load(["id" => "beneficiary_id"]);
```

#### Example: List

```php
// list() returns an array of Beneficiary records (throws on error).
$beneficiarys = $client->Beneficiary()->list();
```

#### Example: Create

```php
$beneficiary = $client->Beneficiary()->create([
]);
```


### BeneficiaryGroup

Create an instance: `$beneficiary_group = $client->BeneficiaryGroup();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `group_member` | `array` |  |
| `group_name` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |

#### Example: List

```php
// list() returns an array of BeneficiaryGroup records (throws on error).
$beneficiary_groups = $client->BeneficiaryGroup()->list();
```


### Card

Create an instance: `$card = $client->Card();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorized_amount` | `string` |  |
| `currency_code` | `string` |  |
| `is_payer_authentication_required` | `bool` |  |
| `is_soft_decline` | `bool` |  |
| `payer_authentication_access_token` | `string` |  |
| `payer_authentication_merchant_data` | `string` |  |
| `payer_authentication_url` | `string` |  |
| `payer_authentication_window_height` | `int` |  |
| `payer_authentication_window_width` | `int` |  |
| `payment_request_callback_url` | `string` |  |
| `payment_request_id` | `string` |  |
| `request_id` | `string` |  |
| `response_code` | `string` |  |
| `response_type` | `string` |  |
| `status` | `string` |  |
| `three_ds_redirect_url` | `string` |  |
| `transaction_id` | `string` |  |

#### Example: Create

```php
$card = $client->Card()->create([
    "paymentrequest_id" => null, // string
]);
```


### CardCustomerToken

Create an instance: `$card_customer_token = $client->CardCustomerToken();`

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

```php
// load() returns the bare CardCustomerToken record (throws on error).
$card_customer_token = $client->CardCustomerToken()->load(["customer_email_address" => "customer_email_address"]);
```

#### Example: List

```php
// list() returns an array of CardCustomerToken records (throws on error).
$card_customer_tokens = $client->CardCustomerToken()->list();
```


### CardPayment

Create an instance: `$card_payment = $client->CardPayment();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorized_amount` | `string` |  |
| `currency_code` | `string` |  |
| `is_payer_authentication_required` | `bool` |  |
| `is_soft_decline` | `bool` |  |
| `payer_authentication_access_token` | `string` |  |
| `payer_authentication_merchant_data` | `string` |  |
| `payer_authentication_url` | `string` |  |
| `payer_authentication_window_height` | `int` |  |
| `payer_authentication_window_width` | `int` |  |
| `payment_request_callback_url` | `string` |  |
| `payment_request_id` | `string` |  |
| `request_id` | `string` |  |
| `response_code` | `string` |  |
| `response_type` | `string` |  |
| `status` | `string` |  |
| `three_ds_redirect_url` | `string` |  |
| `transaction_id` | `string` |  |

#### Example: Create

```php
$card_payment = $client->CardPayment()->create([
    "paymentrequest_id" => null, // string
]);
```


### CardPublicKey

Create an instance: `$card_public_key = $client->CardPublicKey();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `jwt` | `string` |  |

#### Example: Load

```php
// load() returns the bare CardPublicKey record (throws on error).
$card_public_key = $client->CardPublicKey()->load(["paymentrequest_id" => "paymentrequest_id"]);
```


### Consent

Create an instance: `$consent = $client->Consent();`

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
| `is_connected_account` | `bool` |  |
| `is_enabled` | `bool` |  |
| `merchant_id` | `string` |  |
| `provider` | `string` |  |
| `success_web_hook_url` | `string` |  |

#### Example: Load

```php
// load() returns the bare Consent record (throws on error).
$consent = $client->Consent()->load(["id" => "consent_id"]);
```

#### Example: List

```php
// list() returns an array of Consent records (throws on error).
$consents = $client->Consent()->list();
```

#### Example: Create

```php
$consent = $client->Consent()->create([
]);
```


### Currency

Create an instance: `$currency = $client->Currency();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `code` | `string` |  |
| `decimal` | `int` |  |
| `is_fiat` | `bool` |  |
| `iso4217_alpha_code` | `string` |  |
| `iso4217_numeric_code` | `string` |  |
| `symbol` | `string` |  |

#### Example: List

```php
// list() returns an array of Currency records (throws on error).
$currencys = $client->Currency()->list();
```


### DirectDebitBatchSubmit

Create an instance: `$direct_debit_batch_submit = $client->DirectDebitBatchSubmit();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_submission` | `array` |  |
| `successful_submission` | `array` |  |

#### Example: Create

```php
$direct_debit_batch_submit = $client->DirectDebitBatchSubmit()->create([
]);
```


### FxRate

Create an instance: `$fx_rate = $client->FxRate();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `destination_currency` | `string` |  |
| `exchange_rate` | `float` |  |
| `expiry_time` | `string` |  |
| `quote_id` | `string` |  |
| `source_currency` | `string` |  |

#### Example: Load

```php
// load() returns the bare FxRate record (throws on error).
$fx_rate = $client->FxRate()->load(["destination" => "destination", "source" => "source", "valid_for_minute" => 1]);
```

#### Example: List

```php
// list() returns an array of FxRate records (throws on error).
$fx_rates = $client->FxRate()->list();
```


### IPayment

Create an instance: `$i_payment = $client->IPayment();`

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

```php
$i_payment = $client->IPayment()->create([
]);
```


### Mandate

Create an instance: `$mandate = $client->Mandate();`

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
| `is_recurring` | `bool` |  |
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

```php
// load() returns the bare Mandate record (throws on error).
$mandate = $client->Mandate()->load(["id" => "mandate_id"]);
```

#### Example: Create

```php
$mandate = $client->Mandate()->create([
    "address_line1" => null, // string
    "city" => null, // string
    "country_code" => null, // string
    "email_address" => null, // string
    "first_name" => null, // string
    "last_name" => null, // string
    "postal_code" => null, // string
]);
```


### Merchant

Create an instance: `$merchant = $client->Merchant();`

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
| `account_currency` | `array` |  |
| `can_have_trust_account` | `bool` |  |
| `card_payment_processor` | `string` |  |
| `company_id` | `string` |  |
| `display_qr_on_hosted_pay` | `bool` |  |
| `hosted_pay_version` | `int` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_blocked` | `bool` |  |
| `is_exited` | `bool` |  |
| `is_suspended` | `bool` |  |
| `jurisdiction` | `string` |  |
| `logo_url_png` | `string` |  |
| `logo_url_svg` | `string` |  |
| `merchant_category_code` | `string` |  |
| `name` | `string` |  |
| `note` | `string` |  |
| `parent_merchant` | `array` |  |
| `payment_account` | `array` |  |
| `payment_account_limit` | `int` |  |
| `reason` | `string` |  |
| `short_name` | `string` |  |
| `supported_payment_methods_list` | `array` |  |
| `suspension_reason` | `string` |  |
| `tag` | `array` |  |
| `time_zone_id` | `string` |  |
| `trading_name` | `string` |  |
| `web_hook_limit` | `int` |  |
| `your_role_name` | `string` |  |

#### Example: Load

```php
// load() returns the bare Merchant record (throws on error).
$merchant = $client->Merchant()->load(["id" => "merchant_id"]);
```

#### Example: List

```php
// list() returns an array of Merchant records (throws on error).
$merchants = $client->Merchant()->list();
```


### MerchantAuthorisationSetting

Create an instance: `$merchant_authorisation_setting = $client->MerchantAuthorisationSetting();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount_lower` | `float` |  |
| `amount_upper` | `float` |  |
| `authorisation_type` | `string` |  |
| `beneficiaries_only` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_editor_cant_authorise` | `bool` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `number_of_authoriser` | `int` |  |
| `role_setting` | `array` |  |

#### Example: List

```php
// list() returns an array of MerchantAuthorisationSetting records (throws on error).
$merchant_authorisation_settings = $client->MerchantAuthorisationSetting()->list();
```


### MerchantDirectDebitMandate

Create an instance: `$merchant_direct_debit_mandate = $client->MerchantDirectDebitMandate();`

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
| `is_recurring` | `bool` |  |
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

```php
// list() returns an array of MerchantDirectDebitMandate records (throws on error).
$merchant_direct_debit_mandates = $client->MerchantDirectDebitMandate()->list();
```


### MerchantPayByBankSetting

Create an instance: `$merchant_pay_by_bank_setting = $client->MerchantPayByBankSetting();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bank_country_code` | `array` |  |
| `bank_id` | `string` |  |
| `bank_name` | `string` |  |
| `business_institution_id` | `string` |  |
| `currency` | `string` |  |
| `logo` | `string` |  |
| `message` | `string` |  |
| `message_image_url` | `string` |  |
| `order` | `int` |  |
| `personal_institution_id` | `string` |  |
| `processor` | `string` |  |
| `warning_heading` | `string` |  |
| `warning_message` | `string` |  |

#### Example: List

```php
// list() returns an array of MerchantPayByBankSetting records (throws on error).
$merchant_pay_by_bank_settings = $client->MerchantPayByBankSetting()->list();
```


### MerchantPaymentRequestTemplate

Create an instance: `$merchant_payment_request_template = $client->MerchantPaymentRequestTemplate();`

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
| `description` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `template` | `array` |  |

#### Example: Load

```php
// load() returns the bare MerchantPaymentRequestTemplate record (throws on error).
$merchant_payment_request_template = $client->MerchantPaymentRequestTemplate()->load(["id" => "merchant_payment_request_template_id", "paymentrequest_id" => "paymentrequest_id"]);
```

#### Example: List

```php
// list() returns an array of MerchantPaymentRequestTemplate records (throws on error).
$merchant_payment_request_templates = $client->MerchantPaymentRequestTemplate()->list();
```


### MerchantToken

Create an instance: `$merchant_token = $client->MerchantToken();`

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
| `authentication_method` | `array` |  |
| `authorisation` | `array` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `can_authorise` | `bool` |  |
| `description` | `string` |  |
| `expires_at` | `string` |  |
| `has_current_user_authorised` | `bool` |  |
| `hmac_algorithm` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `ip_address_whitelist` | `string` |  |
| `is_archived` | `bool` |  |
| `is_enabled` | `bool` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `nonce` | `string` |  |
| `permission_type` | `array` |  |
| `request_signature_version` | `int` |  |
| `shared_secret_algorithm` | `string` |  |
| `shared_secret_base64` | `string` |  |
| `token` | `string` |  |

#### Example: Load

```php
// load() returns the bare MerchantToken record (throws on error).
$merchant_token = $client->MerchantToken()->load(["id" => "merchant_token_id"]);
```

#### Example: List

```php
// list() returns an array of MerchantToken records (throws on error).
$merchant_tokens = $client->MerchantToken()->list();
```

#### Example: Create

```php
$merchant_token = $client->MerchantToken()->create([
    "nonce" => null, // string
]);
```


### Metadata

Create an instance: `$metadata = $client->Metadata();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```php
// load() returns the bare Metadata record (throws on error).
$metadata = $client->Metadata()->load();
```


### NoFrixionVersion

Create an instance: `$no_frixion_version = $client->NoFrixionVersion();`

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
| `release_name` | `string` |  |

#### Example: Load

```php
// load() returns the bare NoFrixionVersion record (throws on error).
$no_frixion_version = $client->NoFrixionVersion()->load();
```


### OpenBanking

Create an instance: `$open_banking = $client->OpenBanking();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```php
$open_banking = $client->OpenBanking()->create([
    "account_id" => null, // string
]);
```


### Payeeverification

Create an instance: `$payeeverification = $client->Payeeverification();`

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

```php
$payeeverification = $client->Payeeverification()->create([
    "account_name" => null, // string
    "iban" => null, // string
]);
```


### Payment

Create an instance: `$payment = $client->Payment();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `array` |  |
| `amount` | `float` |  |
| `amount_pending` | `float` |  |
| `amount_received` | `float` |  |
| `amount_refunded` | `float` |  |
| `auto_send_receipt` | `bool` |  |
| `base_origin_url` | `string` |  |
| `callback_url` | `string` |  |
| `card_authorize_only` | `bool` |  |
| `card_create_token` | `bool` |  |
| `card_create_token_mode` | `string` |  |
| `card_ignore_cvn` | `bool` |  |
| `card_no_payer_authentication` | `bool` |  |
| `card_processor_merchant_id` | `string` |  |
| `card_stripe_payment_intent_id` | `string` |  |
| `card_stripe_payment_intent_secret` | `string` |  |
| `card_transmit_raw_detail` | `bool` |  |
| `created_by_user` | `array` |  |
| `currency` | `string` |  |
| `custom_field` | `array` |  |
| `customer_email_address` | `string` |  |
| `customer_id` | `string` |  |
| `customer_name` | `string` |  |
| `description` | `string` |  |
| `destination_account` | `array` |  |
| `direct_debit_payment` | `array` |  |
| `due_date` | `string` |  |
| `event` | `array` |  |
| `failure_callback_url` | `string` |  |
| `field_display_setting` | `array` |  |
| `formatted_amount` | `string` |  |
| `hosted_pay_checkout_url` | `string` |  |
| `id` | `string` |  |
| `ignore_address_verification` | `bool` |  |
| `inserted` | `string` |  |
| `inserted_sortable` | `string` |  |
| `is_archived` | `bool` |  |
| `jwk` | `string` |  |
| `last_updated` | `string` |  |
| `lightning_invoice` | `string` |  |
| `lightning_invoice_expires_at` | `string` |  |
| `merchant_direct_debit_mandate_id` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `notification_email_address` | `string` |  |
| `notification_role_i_d` | `array` |  |
| `order_id` | `string` |  |
| `partial_payment_method` | `string` |  |
| `partial_payment_step` | `string` |  |
| `payment_attempt` | `array` |  |
| `payment_method` | `array` |  |
| `payment_processor` | `string` |  |
| `payrun_id` | `string` |  |
| `pisp_account_id` | `string` |  |
| `priority_bank_id` | `string` |  |
| `result` | `array` |  |
| `sandbox_settle_delay_in_second` | `int` |  |
| `shipping_address` | `array` |  |
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
| `tag` | `array` |  |
| `tag_id` | `array` |  |
| `title` | `string` |  |
| `tokenised_card` | `array` |  |
| `transaction` | `array` |  |
| `use_hosted_payment_page` | `bool` |  |

#### Example: Load

```php
// load() returns the bare Payment record (throws on error).
$payment = $client->Payment()->load(["id" => "payment_id"]);
```

#### Example: Create

```php
$payment = $client->Payment()->create([
    "created_by_user" => null, // array
]);
```


### PaymentAccount

Create an instance: `$payment_account = $client->PaymentAccount();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `string` |  |
| `account_supplier_name` | `string` |  |
| `available_balance` | `float` |  |
| `available_balance_minor_unit` | `int` |  |
| `balance` | `float` |  |
| `balance_minor_unit` | `int` |  |
| `bank_name` | `string` |  |
| `consent_id` | `string` |  |
| `created_by` | `array` |  |
| `created_by_display_name` | `string` |  |
| `currency` | `string` |  |
| `default_payment_rail` | `string` |  |
| `display_name` | `string` |  |
| `expiry_date` | `string` |  |
| `external_account_icon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `array` |  |
| `inserted` | `string` |  |
| `is_archived` | `bool` |  |
| `is_connected_account` | `bool` |  |
| `is_default` | `bool` |  |
| `is_trust_account` | `bool` |  |
| `is_virtual` | `bool` |  |
| `last_transaction` | `array` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `physical_account_id` | `string` |  |
| `rule` | `array` |  |
| `submitted_payouts_balance` | `float` |  |
| `submitted_payouts_balance_minor_unit` | `int` |  |
| `summary` | `string` |  |
| `supplier_sepa_instant_status` | `string` |  |
| `xero_bank_feed_connection_status` | `string` |  |
| `xero_bank_feed_last_synced_at` | `string` |  |
| `xero_bank_feed_sync_last_failed_at` | `string` |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` |  |
| `xero_bank_feed_sync_status` | `string` |  |
| `xero_unsynchronised_transactions_count` | `int` |  |

#### Example: List

```php
// list() returns an array of PaymentAccount records (throws on error).
$payment_accounts = $client->PaymentAccount()->list();
```


### PaymentAccountMinimal

Create an instance: `$payment_account_minimal = $client->PaymentAccountMinimal();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `string` |  |
| `available_balance` | `float` |  |
| `balance` | `float` |  |
| `balance_minor_unit` | `int` |  |
| `currency` | `string` |  |
| `id` | `string` |  |
| `identifier` | `array` |  |
| `is_archived` | `bool` |  |
| `is_connected_account` | `bool` |  |
| `merchant_id` | `string` |  |
| `submitted_payouts_balance` | `float` |  |

#### Example: List

```php
// list() returns an array of PaymentAccountMinimal records (throws on error).
$payment_account_minimals = $client->PaymentAccountMinimal()->list();
```


### PaymentInitiation

Create an instance: `$payment_initiation = $client->PaymentInitiation();`

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

```php
$payment_initiation = $client->PaymentInitiation()->create([
    "paymentrequest_id" => null, // string
]);
```


### PaymentRequest

Create an instance: `$payment_request = $client->PaymentRequest();`

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
| `address` | `array` |  |
| `amount` | `float` |  |
| `amount_pending` | `float` |  |
| `amount_received` | `float` |  |
| `amount_refunded` | `float` |  |
| `auto_send_receipt` | `bool` |  |
| `base_origin_url` | `string` |  |
| `callback_url` | `string` |  |
| `card_authorize_only` | `bool` |  |
| `card_create_token` | `bool` |  |
| `card_create_token_mode` | `string` |  |
| `card_ignore_cvn` | `bool` |  |
| `card_processor_merchant_id` | `string` |  |
| `card_stripe_payment_intent_id` | `string` |  |
| `card_stripe_payment_intent_secret` | `string` |  |
| `created_by_user` | `array` |  |
| `currency` | `string` |  |
| `custom_field` | `array` |  |
| `customer_email_address` | `string` |  |
| `customer_id` | `string` |  |
| `customer_name` | `string` |  |
| `description` | `string` |  |
| `destination_account` | `array` |  |
| `direct_debit_payment` | `array` |  |
| `do_simulate_settlement_failure` | `bool` |  |
| `due_date` | `string` |  |
| `error_description` | `string` |  |
| `event` | `array` |  |
| `failure_callback_url` | `string` |  |
| `field_display_setting` | `array` |  |
| `formatted_amount` | `string` |  |
| `hosted_pay_checkout_url` | `string` |  |
| `id` | `string` |  |
| `ignore_address_verification` | `bool` |  |
| `inserted` | `string` |  |
| `inserted_sortable` | `string` |  |
| `institution` | `string` |  |
| `is_archived` | `bool` |  |
| `jwk` | `string` |  |
| `last_updated` | `string` |  |
| `lightning_invoice` | `string` |  |
| `lightning_invoice_expires_at` | `string` |  |
| `merchant_direct_debit_mandate_id` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `notification_email_address` | `string` |  |
| `notification_role_i_d` | `array` |  |
| `order_id` | `string` |  |
| `partial_payment_method` | `string` |  |
| `partial_payment_step` | `string` |  |
| `payment_attempt` | `array` |  |
| `payment_initiation_id` | `string` |  |
| `payment_method` | `array` |  |
| `payment_processor` | `string` |  |
| `payrun_id` | `string` |  |
| `pisp_account_id` | `string` |  |
| `priority_bank_id` | `string` |  |
| `result` | `array` |  |
| `sandbox_settle_delay_in_second` | `int` |  |
| `shipping_address` | `array` |  |
| `status` | `string` |  |
| `success_web_hook_url` | `string` |  |
| `tag` | `array` |  |
| `title` | `string` |  |
| `tokenised_card` | `array` |  |
| `transaction` | `array` |  |
| `use_hosted_payment_page` | `bool` |  |

#### Example: Load

```php
// load() returns the bare PaymentRequest record (throws on error).
$payment_request = $client->PaymentRequest()->load();
```

#### Example: List

```php
// list() returns an array of PaymentRequest records (throws on error).
$payment_requests = $client->PaymentRequest()->list();
```

#### Example: Create

```php
$payment_request = $client->PaymentRequest()->create([
    "paymentrequest_id" => null, // string
]);
```


### PaymentRequestEvent

Create an instance: `$payment_request_event = $client->PaymentRequestEvent();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float` |  |
| `apple_pay_transaction_id` | `string` |  |
| `card_authorization_response_id` | `string` |  |
| `card_expiry_month` | `int` |  |
| `card_expiry_year` | `int` |  |
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

```php
// list() returns an array of PaymentRequestEvent records (throws on error).
$payment_request_events = $client->PaymentRequestEvent()->list();
```


### PaymentRequestMetric

Create an instance: `$payment_request_metric = $client->PaymentRequestMetric();`

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
| `total_amounts_by_currency` | `array` |  |
| `unpaid` | `int` |  |

#### Example: Load

```php
// load() returns the bare PaymentRequestMetric record (throws on error).
$payment_request_metric = $client->PaymentRequestMetric()->load();
```


### PaymentRequestMinimal

Create an instance: `$payment_request_minimal = $client->PaymentRequestMinimal();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float` |  |
| `amount_pending` | `float` |  |
| `amount_received` | `float` |  |
| `amount_refunded` | `float` |  |
| `callback_url` | `string` |  |
| `card_stripe_payment_intent_secret` | `string` |  |
| `country_code` | `string` |  |
| `currency` | `string` |  |
| `custom_fields_to_display` | `array` |  |
| `description` | `string` |  |
| `due_date` | `string` |  |
| `field_display_setting` | `array` |  |
| `google_pay_merchant_id` | `string` |  |
| `id` | `string` |  |
| `jwk` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_logo_url_png` | `string` |  |
| `merchant_logo_url_svg` | `string` |  |
| `merchant_name` | `string` |  |
| `merchant_short_name` | `string` |  |
| `partial_payment_method` | `string` |  |
| `payment_attempt` | `array` |  |
| `payment_methods_list` | `array` |  |
| `payment_processor` | `string` |  |
| `payment_processor_key` | `string` |  |
| `pisp_error` | `string` |  |
| `priority_bank_id` | `string` |  |
| `status` | `string` |  |
| `stripe_account_id` | `string` |  |
| `title` | `string` |  |

#### Example: List

```php
// list() returns an array of PaymentRequestMinimal records (throws on error).
$payment_request_minimals = $client->PaymentRequestMinimal()->list();
```


### PaymentRequestResult

Create an instance: `$payment_request_result = $client->PaymentRequestResult();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float` |  |
| `amount_pending` | `float` |  |
| `amount_received` | `float` |  |
| `amount_refunded` | `float` |  |
| `currency` | `string` |  |
| `customer_id` | `string` |  |
| `payment` | `array` |  |
| `payment_request_id` | `string` |  |
| `pisp_authorization` | `array` |  |
| `requested_amount` | `float` |  |
| `result` | `string` |  |

#### Example: List

```php
// list() returns an array of PaymentRequestResult records (throws on error).
$payment_request_results = $client->PaymentRequestResult()->list();
```


### PaymentRequestsCreate

Create an instance: `$payment_requests_create = $client->PaymentRequestsCreate();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_payment_request` | `array` |  |
| `payment_request` | `array` |  |

#### Example: Create

```php
$payment_requests_create = $client->PaymentRequestsCreate()->create([
]);
```


### Payout

Create an instance: `$payout = $client->Payout();`

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
| `allow_incomplete` | `bool` |  |
| `amount` | `float` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `array` |  |
| `authorisation` | `array` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `array` |  |
| `beneficiary_id` | `string` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `array` |  |
| `document` | `array` |  |
| `event` | `array` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `float` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `float` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `reason` | `string` |  |
| `rule` | `array` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `array` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `array` |  |
| `tag_id` | `array` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `float` |  |
| `transacted_fx_amount` | `float` |  |
| `transacted_fx_rate` | `float` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |

#### Example: Load

```php
// load() returns the bare Payout record (throws on error).
$payout = $client->Payout()->load(["id" => "payout_id"]);
```

#### Example: List

```php
// list() returns an array of Payout records (throws on error).
$payouts = $client->Payout()->list();
```

#### Example: Create

```php
$payout = $client->Payout()->create([
]);
```


### PayoutKeyset

Create an instance: `$payout_keyset = $client->PayoutKeyset();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `amount` | `float` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `array` |  |
| `authorisation` | `array` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `array` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `array` |  |
| `document` | `array` |  |
| `event` | `array` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `float` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `float` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `rule` | `array` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `array` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `array` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `float` |  |
| `transacted_fx_amount` | `float` |  |
| `transacted_fx_rate` | `float` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |

#### Example: List

```php
// list() returns an array of PayoutKeyset records (throws on error).
$payout_keysets = $client->PayoutKeyset()->list();
```


### PayoutMetric

Create an instance: `$payout_metric = $client->PayoutMetric();`

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
| `total_amounts_by_currency` | `array` |  |

#### Example: Load

```php
// load() returns the bare PayoutMetric record (throws on error).
$payout_metric = $client->PayoutMetric()->load();
```


### PayoutsCreate

Create an instance: `$payouts_create = $client->PayoutsCreate();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_payout` | `array` |  |
| `payout` | `array` |  |

#### Example: Create

```php
$payouts_create = $client->PayoutsCreate()->create([
]);
```


### Payrun

Create an instance: `$payrun = $client->Payrun();`

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
| `authorisation` | `array` |  |
| `authorisation_date` | `string` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `string` |  |
| `can_authorise` | `bool` |  |
| `can_delete` | `bool` |  |
| `can_edit` | `bool` |  |
| `event` | `array` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice` | `array` |  |
| `invoices_minimal` | `array` |  |
| `is_archived` | `bool` |  |
| `last_updated` | `string` |  |
| `last_updated_by` | `array` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `note` | `string` |  |
| `payment` | `array` |  |
| `payout` | `array` |  |
| `payouts_count` | `int` |  |
| `reason` | `string` |  |
| `schedule_date` | `string` |  |
| `scheduled_date` | `string` |  |
| `source_account` | `array` |  |
| `status` | `string` |  |
| `total_eur` | `float` |  |
| `total_gbp` | `float` |  |
| `total_usd` | `float` |  |

#### Example: Load

```php
// load() returns the bare Payrun record (throws on error).
$payrun = $client->Payrun()->load(["id" => "payrun_id"]);
```

#### Example: List

```php
// list() returns an array of Payrun records (throws on error).
$payruns = $client->Payrun()->list();
```

#### Example: Create

```php
$payrun = $client->Payrun()->create([
    "id" => null, // string
]);
```


### Report

Create an instance: `$report = $client->Report();`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |


### ReportResult

Create an instance: `$report_result = $client->ReportResult();`

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
| `statement_number` | `int` |  |

#### Example: Load

```php
// load() returns the bare ReportResult record (throws on error).
$report_result = $client->ReportResult()->load(["id" => 1, "report_id" => "report_id"]);
```


### RolesCreate

Create an instance: `$roles_create = $client->RolesCreate();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_role` | `array` |  |
| `role` | `array` |  |

#### Example: Create

```php
$roles_create = $client->RolesCreate()->create([
    "merchant_id" => null, // string
]);
```


### Rule

Create an instance: `$rule = $client->Rule();`

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
| `account` | `array` |  |
| `account_id` | `string` |  |
| `approve_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `array` |  |
| `authorisation` | `array` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `can_authorise` | `bool` |  |
| `created_by` | `array` |  |
| `description` | `string` |  |
| `end_at` | `string` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_disabled` | `bool` |  |
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
| `sweep_action` | `array` |  |
| `time_zone_id` | `string` |  |
| `trigger_cron_expression` | `string` |  |
| `trigger_on_pay_in` | `bool` |  |
| `user_id` | `string` |  |
| `web_hook_secret` | `string` |  |

#### Example: Load

```php
// load() returns the bare Rule record (throws on error).
$rule = $client->Rule()->load(["id" => "rule_id"]);
```

#### Example: List

```php
// list() returns an array of Rule records (throws on error).
$rules = $client->Rule()->list();
```

#### Example: Create

```php
$rule = $client->Rule()->create([
    "created_by" => null, // array
    "nonce" => null, // string
]);
```


### RuleEvent

Create an instance: `$rule_event = $client->RuleEvent();`

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
| `is_authorise_to_enable` | `bool` |  |
| `message` | `string` |  |
| `raw_response` | `string` |  |
| `rule_event_type` | `string` |  |
| `rule_id` | `string` |  |
| `user` | `array` |  |

#### Example: List

```php
// list() returns an array of RuleEvent records (throws on error).
$rule_events = $client->RuleEvent()->list();
```


### Tag

Create an instance: `$tag = $client->Tag();`

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

```php
// list() returns an array of Tag records (throws on error).
$tags = $client->Tag()->list();
```

#### Example: Create

```php
$tag = $client->Tag()->create([
    "merchant_id" => null, // string
]);
```


### Token

Create an instance: `$token = $client->Token();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```php
$token = $client->Token()->create([
    "id" => null, // string
]);
```


### Transaction

Create an instance: `$transaction = $client->Transaction();`

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
| `account_id` | `string` |  |
| `account_name` | `string` |  |
| `account_sequence_number` | `int` |  |
| `address_detail` | `array` |  |
| `amount` | `float` |  |
| `amount_minor_unit` | `int` |  |
| `balance` | `float` |  |
| `balance_minor_unit` | `int` |  |
| `booking_date_time` | `string` |  |
| `charge_detail` | `array` |  |
| `content` | `array` |  |
| `counterparty` | `array` |  |
| `counterparty_summary` | `string` |  |
| `currency` | `string` |  |
| `currency_exchange` | `array` |  |
| `date` | `string` |  |
| `description` | `string` |  |
| `enrichment` | `array` |  |
| `fx_amount` | `float` |  |
| `fx_currency` | `string` |  |
| `fx_rate` | `float` |  |
| `gross_amount` | `array` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `iso_bank_transaction_code` | `array` |  |
| `merchant` | `array` |  |
| `merchant_id` | `string` |  |
| `page_number` | `int` |  |
| `page_size` | `int` |  |
| `payee_detail` | `array` |  |
| `payer_detail` | `array` |  |
| `payment_request_custom_field` | `array` |  |
| `payment_request_id` | `string` |  |
| `payout_id` | `string` |  |
| `proprietary_bank_transaction_code` | `array` |  |
| `raw_reference` | `string` |  |
| `reference` | `string` |  |
| `rule_id` | `string` |  |
| `statement_reference` | `array` |  |
| `status` | `string` |  |
| `supplementary_data` | `mixed` |  |
| `tag` | `array` |  |
| `their_reference` | `string` |  |
| `total_page` | `int` |  |
| `total_size` | `int` |  |
| `transaction_amount` | `array` |  |
| `transaction_date` | `string` |  |
| `transaction_information` | `array` |  |
| `transaction_mutability` | `string` |  |
| `type` | `string` |  |
| `value_date_time` | `string` |  |
| `virtual_iban` | `string` |  |
| `your_reference` | `string` |  |

#### Example: Load

```php
// load() returns the bare Transaction record (throws on error).
$transaction = $client->Transaction()->load(["id" => "transaction_id"]);
```

#### Example: List

```php
// list() returns an array of Transaction records (throws on error).
$transactions = $client->Transaction()->list();
```

#### Example: Create

```php
$transaction = $client->Transaction()->create([
    "id" => null, // string
]);
```


### User

Create an instance: `$user = $client->User();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client_session_timeout` | `array` |  |
| `email_address` | `string` |  |
| `first_name` | `string` |  |
| `id` | `string` |  |
| `last_name` | `string` |  |
| `passkey_added` | `bool` |  |
| `permission` | `array` |  |
| `profile` | `string` |  |
| `roles_with_scope` | `array` |  |
| `two_factor_enabled` | `bool` |  |
| `user_invite_id` | `string` |  |

#### Example: List

```php
// list() returns an array of User records (throws on error).
$users = $client->User()->list();
```


### UserInvite

Create an instance: `$user_invite = $client->UserInvite();`

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
| `authorisation_status` | `array` |  |
| `id` | `string` |  |
| `initial_role_id` | `string` |  |
| `invitee_email_address` | `string` |  |
| `invitee_first_name` | `string` |  |
| `invitee_last_name` | `string` |  |
| `inviter_email_address` | `string` |  |
| `inviter_first_name` | `string` |  |
| `inviter_last_name` | `string` |  |
| `is_authorised` | `bool` |  |
| `is_invitee_registered` | `bool` |  |
| `last_invited` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `message` | `string` |  |
| `registration_url` | `string` |  |
| `send_invite_email` | `bool` |  |
| `status` | `string` |  |
| `user` | `array` |  |
| `user_id` | `string` |  |

#### Example: Load

```php
// load() returns the bare UserInvite record (throws on error).
$user_invite = $client->UserInvite()->load(["id" => "user_invite_id"]);
```

#### Example: List

```php
// list() returns an array of UserInvite records (throws on error).
$user_invites = $client->UserInvite()->list();
```

#### Example: Create

```php
$user_invite = $client->UserInvite()->create([
]);
```


### UserInvitesCreate

Create an instance: `$user_invites_create = $client->UserInvitesCreate();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_user_invite` | `array` |  |
| `user_invite` | `array` |  |

#### Example: Create

```php
$user_invites_create = $client->UserInvitesCreate()->create([
]);
```


### Virtual

Create an instance: `$virtual = $client->Virtual();`

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
| `available_balance` | `float` |  |
| `available_balance_minor_unit` | `int` |  |
| `balance` | `float` |  |
| `balance_minor_unit` | `int` |  |
| `bank_name` | `string` |  |
| `consent_id` | `string` |  |
| `created_by` | `array` |  |
| `created_by_display_name` | `string` |  |
| `currency` | `string` |  |
| `default_payment_rail` | `string` |  |
| `display_name` | `string` |  |
| `expiry_date` | `string` |  |
| `external_account_icon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `array` |  |
| `inserted` | `string` |  |
| `is_archived` | `bool` |  |
| `is_connected_account` | `bool` |  |
| `is_default` | `bool` |  |
| `is_trust_account` | `bool` |  |
| `is_virtual` | `bool` |  |
| `last_transaction` | `array` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `name` | `string` |  |
| `physical_account_id` | `string` |  |
| `rule` | `array` |  |
| `submitted_payouts_balance` | `float` |  |
| `submitted_payouts_balance_minor_unit` | `int` |  |
| `summary` | `string` |  |
| `supplier_sepa_instant_status` | `string` |  |
| `xero_bank_feed_connection_status` | `string` |  |
| `xero_bank_feed_last_synced_at` | `string` |  |
| `xero_bank_feed_sync_last_failed_at` | `string` |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` |  |
| `xero_bank_feed_sync_status` | `string` |  |
| `xero_unsynchronised_transactions_count` | `int` |  |

#### Example: Create

```php
$virtual = $client->Virtual()->create([
    "account_id" => null, // string
]);
```


### Webhook

Create an instance: `$webhook = $client->Webhook();`

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
| `destination_url` | `string` |  |
| `email_address` | `string` |  |
| `failed_notification_email_address` | `string` |  |
| `id` | `string` |  |
| `is_active` | `bool` |  |
| `merchant_id` | `string` |  |
| `notification_method` | `string` |  |
| `resource_type` | `array` |  |
| `retry` | `bool` |  |
| `secret` | `string` |  |
| `version` | `int` |  |

#### Example: Load

```php
// load() returns the bare Webhook record (throws on error).
$webhook = $client->Webhook()->load(["id" => "webhook_id"]);
```

#### Example: List

```php
// list() returns an array of Webhook records (throws on error).
$webhooks = $client->Webhook()->list();
```

#### Example: Create

```php
$webhook = $client->Webhook()->create([
]);
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

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── nofrixion_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`nofrixion_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```php
$account = $client->Account();
$account->list();

// $account->data_get() now returns the account data from the last list
// $account->match_get() returns the last match criteria
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
