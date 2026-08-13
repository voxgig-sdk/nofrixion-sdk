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
        echo $item["id"] . " " . $item["accountBalances"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 3. Load a cardcustomertoken

CardCustomerToken is nested under customer_email_address, so provide the `customer_email_address`.

```php
try {
    // load() returns the ENTITY — call data_get() for the CardCustomerToken record (throws on error).
    $cardcustomertoken = $client->CardCustomerToken()->load(["customer_email_address" => "example_customer_email_address"]);
    print_r($cardcustomertoken);
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 4. Create, update, and remove

```php
// create() returns the ENTITY — call data_get() for the created Account record.
$created = $client->Account()->create(["createdBy" => [], "identifier" => []]);

// Update — index the record via data_get() ($created->data_get()["id"]).
$client->Account()->update(["id" => $created->data_get()["id"], "account_id" => "example_account_id", "amount" => 1]);

// Remove
$client->Account()->remove(["id" => $created->data_get()["id"]]);
```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $consents = $client->Consent()->list();
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
    "entity" => ["consent" => ["test01" => ["id" => "test01"]]],
]);

// Entity ops return the ENTITY (throws on error);
// call data_get() for the mock record.
$consent = $client->Consent()->list();
print_r($consent);
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
| `MerchantDirectDebitMandatePage` | `($data): MerchantDirectDebitMandatePageEntity` | Create a MerchantDirectDebitMandatePage entity instance. |
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
| `Payout` | `($data): PayoutEntity` | Create a Payout entity instance. |
| `PayoutKeysetPage` | `($data): PayoutKeysetPageEntity` | Create a PayoutKeysetPage entity instance. |
| `PayoutMetric` | `($data): PayoutMetricEntity` | Create a PayoutMetric entity instance. |
| `Payrun` | `($data): PayrunEntity` | Create a Payrun entity instance. |
| `Report` | `($data): ReportEntity` | Create a Report entity instance. |
| `ReportResult` | `($data): ReportResultEntity` | Create a ReportResult entity instance. |
| `Role` | `($data): RoleEntity` | Create a Role entity instance. |
| `Rule` | `($data): RuleEntity` | Create a Rule entity instance. |
| `RuleEvent` | `($data): RuleEventEntity` | Create a RuleEvent entity instance. |
| `Tag` | `($data): TagEntity` | Create a Tag entity instance. |
| `Token` | `($data): TokenEntity` | Create a Token entity instance. |
| `Transaction` | `($data): TransactionEntity` | Create a Transaction entity instance. |
| `User` | `($data): UserEntity` | Create an User entity instance. |
| `UserInvite` | `($data): UserInviteEntity` | Create an UserInvite entity instance. |
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

Entity operations return the ENTITY (call data_get() for the record) (an `array` for single-entity
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
| `accountBalances` |  |
| `accountID` |  |
| `accountIdentifications` |  |
| `accountName` |  |
| `accountNames` |  |
| `accountSupplierName` |  |
| `accountType` |  |
| `availableBalance` |  |
| `availableBalanceMinorUnits` |  |
| `balance` |  |
| `balanceMinorUnits` |  |
| `bankName` |  |
| `consentID` |  |
| `consolidatedAccountInformation` |  |
| `createdBy` |  |
| `createdByDisplayName` |  |
| `currency` |  |
| `defaultPaymentRail` |  |
| `description` |  |
| `details` |  |
| `displayName` |  |
| `expiryDate` |  |
| `externalAccountIcon` |  |
| `format` |  |
| `fromDate` |  |
| `id` |  |
| `identifier` |  |
| `inserted` |  |
| `isArchived` |  |
| `isConnectedAccount` |  |
| `isDefault` |  |
| `isTrustAccount` |  |
| `isVirtual` |  |
| `lastTransaction` |  |
| `lastUpdated` |  |
| `merchantID` |  |
| `merchantName` |  |
| `nickname` |  |
| `physicalAccountID` |  |
| `roleIDs` |  |
| `rules` |  |
| `submittedPayoutsBalance` |  |
| `submittedPayoutsBalanceMinorUnits` |  |
| `summary` |  |
| `supplierPhysicalAccountID` |  |
| `supplierSepaInstantStatus` |  |
| `toDate` |  |
| `type` |  |
| `usageType` |  |
| `xeroBankFeedConnectionStatus` |  |
| `xeroBankFeedLastSyncedAt` |  |
| `xeroBankFeedSyncLastFailedAt` |  |
| `xeroBankFeedSyncLastFailureReason` |  |
| `xeroBankFeedSyncStatus` |  |
| `xeroUnsynchronisedTransactionsCount` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/accounts/{accountID}/{currency}`

#### Batch

| Field | Description |
| --- | --- |
| `approveUrl` |  |
| `id` |  |
| `payouts` |  |

Operations: Create, Load.

API path: `/api/v1/payouts/batch`

#### Beneficiary

| Field | Description |
| --- | --- |
| `approvalCallbackUrl` |  |
| `authenticationMethods` |  |
| `authorisations` |  |
| `authorisersCompletedCount` |  |
| `authorisersRequiredCount` |  |
| `beneficiaries` |  |
| `beneficiaryEvents` |  |
| `canAuthorise` |  |
| `canUpdate` |  |
| `createdBy` |  |
| `createdByEmailAddress` |  |
| `currency` |  |
| `destination` |  |
| `failedBeneficiaries` |  |
| `hasCurrentUserAuthorised` |  |
| `id` |  |
| `inserted` |  |
| `isEnabled` |  |
| `lastAuthorised` |  |
| `lastUpdated` |  |
| `merchantID` |  |
| `name` |  |
| `nonce` |  |
| `sourceAccountIDs` |  |
| `sourceAccounts` |  |
| `theirReference` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/beneficiaries/authorise/{id}`

#### BeneficiaryGroup

| Field | Description |
| --- | --- |
| `groupMembers` |  |
| `groupName` |  |
| `id` |  |
| `inserted` |  |
| `lastUpdated` |  |
| `merchantID` |  |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/beneficiarygroups`

#### Card

| Field | Description |
| --- | --- |
| `authorizedAmount` |  |
| `currencyCode` |  |
| `isPayerAuthenticationRequired` |  |
| `isSoftDecline` |  |
| `payerAuthenticationAccessToken` |  |
| `payerAuthenticationMerchantData` |  |
| `payerAuthenticationUrl` |  |
| `payerAuthenticationWindowHeight` |  |
| `payerAuthenticationWindowWidth` |  |
| `paymentRequestCallbackUrl` |  |
| `paymentRequestID` |  |
| `requestID` |  |
| `responseCode` |  |
| `responseType` |  |
| `status` |  |
| `threeDSRedirectUrl` |  |
| `transactionID` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/{id}/card`

#### CardCustomerToken

| Field | Description |
| --- | --- |
| `cardType` |  |
| `customerEmailAddress` |  |
| `expiryMonth` |  |
| `expiryYear` |  |
| `id` |  |
| `inserted` |  |
| `lastFourDigits` |  |
| `lastUpdated` |  |
| `maskedCardNumber` |  |
| `merchantID` |  |
| `paymentRequestID` |  |

Operations: List, Load, Remove.

API path: `/api/v1/paymentrequests/card/customertokens/{merchantID}/{customerEmailAddress}`

#### CardPayment

| Field | Description |
| --- | --- |
| `authorizedAmount` |  |
| `currencyCode` |  |
| `isPayerAuthenticationRequired` |  |
| `isSoftDecline` |  |
| `payerAuthenticationAccessToken` |  |
| `payerAuthenticationMerchantData` |  |
| `payerAuthenticationUrl` |  |
| `payerAuthenticationWindowHeight` |  |
| `payerAuthenticationWindowWidth` |  |
| `paymentRequestCallbackUrl` |  |
| `paymentRequestID` |  |
| `requestID` |  |
| `responseCode` |  |
| `responseType` |  |
| `status` |  |
| `threeDSRedirectUrl` |  |
| `transactionID` |  |

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
| `authorisationUrl` |  |
| `callbackUrl` |  |
| `consentID` |  |
| `emailAddress` |  |
| `expiryDate` |  |
| `failureCallbackUrl` |  |
| `id` |  |
| `inserted` |  |
| `institutionID` |  |
| `isConnectedAccounts` |  |
| `isEnabled` |  |
| `merchantID` |  |
| `provider` |  |
| `successWebHookUrl` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/openbanking/consents`

#### Currency

| Field | Description |
| --- | --- |
| `code` |  |
| `decimals` |  |
| `isFiat` |  |
| `iso4217AlphaCode` |  |
| `iso4217NumericCode` |  |
| `symbol` |  |

Operations: List.

API path: `/api/v1/currencies`

#### DirectDebitBatchSubmit

| Field | Description |
| --- | --- |
| `failedSubmissions` |  |
| `successfulSubmissions` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/directdebit/batchsubmit`

#### FxRate

| Field | Description |
| --- | --- |
| `destinationCurrency` |  |
| `exchangeRate` |  |
| `expiryTime` |  |
| `quoteID` |  |
| `sourceCurrency` |  |

Operations: List, Load.

API path: `/api/v1/payouts/fxallheldrates/{source}/{destination}`

#### IPayment

| Field | Description |
| --- | --- |
| `paymentRequestID` |  |
| `responseType` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/payondemand`

#### Mandate

| Field | Description |
| --- | --- |
| `accountNumber` |  |
| `addressLine1` |  |
| `addressLine2` |  |
| `approvedAt` |  |
| `city` |  |
| `countryCode` |  |
| `currency` |  |
| `customerAccountNumber` |  |
| `customerCity` |  |
| `customerCountryCode` |  |
| `customerCountryName` |  |
| `customerEmailAddress` |  |
| `customerFirstName` |  |
| `customerIban` |  |
| `customerLastName` |  |
| `customerSortCode` |  |
| `emailAddress` |  |
| `firstName` |  |
| `iban` |  |
| `id` |  |
| `inserted` |  |
| `isRecurring` |  |
| `lastName` |  |
| `lastUpdated` |  |
| `merchantID` |  |
| `postalCode` |  |
| `reference` |  |
| `sortCode` |  |
| `status` |  |
| `supplierBankAccountID` |  |
| `supplierCustomerID` |  |
| `supplierMandateID` |  |
| `supplierName` |  |
| `supplierStatus` |  |

Operations: Create, Load.

API path: `/api/v1/mandates`

#### Merchant

| Field | Description |
| --- | --- |
| `accountCurrencies` |  |
| `canHaveTrustAccounts` |  |
| `cardPaymentProcessor` |  |
| `companyID` |  |
| `displayQrOnHostedPay` |  |
| `hostedPayVersion` |  |
| `id` |  |
| `inserted` |  |
| `isBlocked` |  |
| `isExited` |  |
| `isSuspended` |  |
| `jurisdiction` |  |
| `logoUrlPng` |  |
| `logoUrlSvg` |  |
| `merchantCategoryCode` |  |
| `name` |  |
| `notes` |  |
| `parentMerchant` |  |
| `paymentAccountLimit` |  |
| `paymentAccounts` |  |
| `reason` |  |
| `shortName` |  |
| `supportedPaymentMethodsList` |  |
| `suspensionReason` |  |
| `tags` |  |
| `timeZoneId` |  |
| `tradingName` |  |
| `webHookLimit` |  |
| `yourRoleName` |  |

Operations: List, Load, Remove, Update.

API path: `/api/v1/merchants/{merchantID}/childmerchants`

#### MerchantAuthorisationSetting

| Field | Description |
| --- | --- |
| `amountLower` |  |
| `amountUpper` |  |
| `authorisationType` |  |
| `beneficiariesOnly` |  |
| `id` |  |
| `inserted` |  |
| `lastEditorCantAuthorise` |  |
| `lastUpdated` |  |
| `merchantID` |  |
| `numberOfAuthorisers` |  |
| `roleSettings` |  |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/authorisationsettings`

#### MerchantDirectDebitMandatePage

| Field | Description |
| --- | --- |
| `approvedAt` |  |
| `currency` |  |
| `customerAccountNumber` |  |
| `customerCity` |  |
| `customerCountryCode` |  |
| `customerCountryName` |  |
| `customerEmailAddress` |  |
| `customerFirstName` |  |
| `customerIban` |  |
| `customerLastName` |  |
| `customerSortCode` |  |
| `id` |  |
| `inserted` |  |
| `isRecurring` |  |
| `lastUpdated` |  |
| `merchantID` |  |
| `reference` |  |
| `status` |  |
| `supplierBankAccountID` |  |
| `supplierCustomerID` |  |
| `supplierMandateID` |  |
| `supplierName` |  |
| `supplierStatus` |  |

Operations: List.

API path: `/api/v1/mandates`

#### MerchantPayByBankSetting

| Field | Description |
| --- | --- |
| `bankCountryCodes` |  |
| `bankID` |  |
| `bankName` |  |
| `businessInstitutionID` |  |
| `currency` |  |
| `logo` |  |
| `message` |  |
| `messageImageUrl` |  |
| `order` |  |
| `personalInstitutionID` |  |
| `processor` |  |
| `warningHeading` |  |
| `warningMessage` |  |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/banksettings`

#### MerchantPaymentRequestTemplate

| Field | Description |
| --- | --- |
| `bankPaymentOptions` |  |
| `cardPaymentAddressOptions` |  |
| `cardPaymentCaptureOptions` |  |
| `customFields` |  |
| `defaultFields` |  |
| `description` |  |
| `id` |  |
| `inserted` |  |
| `lastUpdated` |  |
| `merchantID` |  |
| `name` |  |
| `notificationOptions` |  |
| `paymentMethods` |  |
| `paymentTerms` |  |
| `priorityBankOptions` |  |
| `template` |  |

Operations: List, Load, Remove, Update.

API path: `/api/v1/paymentrequests/{merchantID}/templates`

#### MerchantToken

| Field | Description |
| --- | --- |
| `authenticationMethods` |  |
| `authorisations` |  |
| `authorisersCompletedCount` |  |
| `authorisersRequiredCount` |  |
| `canAuthorise` |  |
| `description` |  |
| `expiresAt` |  |
| `hasCurrentUserAuthorised` |  |
| `hmacAlgorithm` |  |
| `id` |  |
| `inserted` |  |
| `ipAddressWhitelist` |  |
| `isArchived` |  |
| `isEnabled` |  |
| `lastAuthorised` |  |
| `lastUpdated` |  |
| `merchantID` |  |
| `nonce` |  |
| `permissionTypes` |  |
| `requestSignatureVersion` |  |
| `sharedSecretAlgorithm` |  |
| `sharedSecretBase64` |  |
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
| `buildVersion` |  |
| `majorVersion` |  |
| `minorVersion` |  |
| `releaseName` |  |

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
| `accountName` |  |
| `accountNumber` |  |
| `iban` |  |
| `payeeVerifiedAccountName` |  |
| `result` |  |
| `secondaryIdentification` |  |
| `sortCode` |  |

Operations: Create.

API path: `/api/v1/openbanking/payeeverification`

#### Payment

| Field | Description |
| --- | --- |
| `addresses` |  |
| `amount` |  |
| `amountPending` |  |
| `amountReceived` |  |
| `amountRefunded` |  |
| `autoSendReceipt` |  |
| `baseOriginUrl` |  |
| `callbackUrl` |  |
| `cardAuthorizeOnly` |  |
| `cardCreateToken` |  |
| `cardCreateTokenMode` |  |
| `cardIgnoreCVN` |  |
| `cardNoPayerAuthentication` |  |
| `cardProcessorMerchantID` |  |
| `cardStripePaymentIntentID` |  |
| `cardStripePaymentIntentSecret` |  |
| `cardTransmitRawDetails` |  |
| `createdByUser` |  |
| `currency` |  |
| `customFields` |  |
| `customerEmailAddress` |  |
| `customerID` |  |
| `customerName` |  |
| `description` |  |
| `destinationAccount` |  |
| `directDebitPayment` |  |
| `dueDate` |  |
| `events` |  |
| `failureCallbackUrl` |  |
| `fieldDisplaySettings` |  |
| `formattedAmount` |  |
| `hostedPayCheckoutUrl` |  |
| `id` |  |
| `ignoreAddressVerification` |  |
| `inserted` |  |
| `insertedSortable` |  |
| `isArchived` |  |
| `jwk` |  |
| `lastUpdated` |  |
| `lightningInvoice` |  |
| `lightningInvoiceExpiresAt` |  |
| `merchantDirectDebitMandateID` |  |
| `merchantID` |  |
| `merchantTokenDescription` |  |
| `notificationEmailAddresses` |  |
| `notificationRoleIDs` |  |
| `orderID` |  |
| `partialPaymentMethod` |  |
| `partialPaymentSteps` |  |
| `paymentAttempts` |  |
| `paymentMethods` |  |
| `paymentProcessor` |  |
| `payrunID` |  |
| `pispAccountID` |  |
| `priorityBankID` |  |
| `result` |  |
| `sandboxSettleDelayInSeconds` |  |
| `shippingAddress` |  |
| `shippingAddressCity` |  |
| `shippingAddressCountryCode` |  |
| `shippingAddressCounty` |  |
| `shippingAddressLine1` |  |
| `shippingAddressLine2` |  |
| `shippingAddressPostCode` |  |
| `shippingEmail` |  |
| `shippingFirstName` |  |
| `shippingLastName` |  |
| `shippingPhone` |  |
| `status` |  |
| `successWebHookUrl` |  |
| `tagIds` |  |
| `tags` |  |
| `title` |  |
| `tokenisedCards` |  |
| `transactions` |  |
| `useHostedPaymentPage` |  |

Operations: Create, Load, Update.

API path: `/api/v1/paymentrequests`

#### PaymentAccount

| Field | Description |
| --- | --- |
| `accountName` |  |
| `accountSupplierName` |  |
| `availableBalance` |  |
| `availableBalanceMinorUnits` |  |
| `balance` |  |
| `balanceMinorUnits` |  |
| `bankName` |  |
| `consentID` |  |
| `createdBy` |  |
| `createdByDisplayName` |  |
| `currency` |  |
| `defaultPaymentRail` |  |
| `displayName` |  |
| `expiryDate` |  |
| `externalAccountIcon` |  |
| `id` |  |
| `identifier` |  |
| `inserted` |  |
| `isArchived` |  |
| `isConnectedAccount` |  |
| `isDefault` |  |
| `isTrustAccount` |  |
| `isVirtual` |  |
| `lastTransaction` |  |
| `lastUpdated` |  |
| `merchantID` |  |
| `merchantName` |  |
| `physicalAccountID` |  |
| `rules` |  |
| `submittedPayoutsBalance` |  |
| `submittedPayoutsBalanceMinorUnits` |  |
| `summary` |  |
| `supplierSepaInstantStatus` |  |
| `xeroBankFeedConnectionStatus` |  |
| `xeroBankFeedLastSyncedAt` |  |
| `xeroBankFeedSyncLastFailedAt` |  |
| `xeroBankFeedSyncLastFailureReason` |  |
| `xeroBankFeedSyncStatus` |  |
| `xeroUnsynchronisedTransactionsCount` |  |

Operations: List.

API path: `/api/v1/accounts/paged`

#### PaymentAccountMinimal

| Field | Description |
| --- | --- |
| `accountName` |  |
| `availableBalance` |  |
| `balance` |  |
| `balanceMinorUnits` |  |
| `currency` |  |
| `id` |  |
| `identifier` |  |
| `isArchived` |  |
| `isConnectedAccount` |  |
| `merchantID` |  |
| `submittedPayoutsBalance` |  |

Operations: List.

API path: `/api/v1/accounts/minimal`

#### PaymentInitiation

| Field | Description |
| --- | --- |
| `paymentInitiationID` |  |
| `paymentRequestCallbackUrl` |  |
| `paymentRequestID` |  |
| `redirectUrl` |  |
| `responseType` |  |
| `specificErrorMessage` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/{id}/pisp`

#### PaymentRequest

| Field | Description |
| --- | --- |
| `addresses` |  |
| `amount` |  |
| `amountPending` |  |
| `amountReceived` |  |
| `amountRefunded` |  |
| `autoSendReceipt` |  |
| `baseOriginUrl` |  |
| `callbackUrl` |  |
| `cardAuthorizeOnly` |  |
| `cardCreateToken` |  |
| `cardCreateTokenMode` |  |
| `cardIgnoreCVN` |  |
| `cardProcessorMerchantID` |  |
| `cardStripePaymentIntentID` |  |
| `cardStripePaymentIntentSecret` |  |
| `createdByUser` |  |
| `currency` |  |
| `customFields` |  |
| `customerEmailAddress` |  |
| `customerID` |  |
| `customerName` |  |
| `description` |  |
| `destinationAccount` |  |
| `directDebitPayment` |  |
| `doSimulateSettlementFailure` |  |
| `dueDate` |  |
| `errorDescription` |  |
| `events` |  |
| `failedPaymentRequests` |  |
| `failureCallbackUrl` |  |
| `fieldDisplaySettings` |  |
| `formattedAmount` |  |
| `hostedPayCheckoutUrl` |  |
| `id` |  |
| `ignoreAddressVerification` |  |
| `inserted` |  |
| `insertedSortable` |  |
| `institution` |  |
| `isArchived` |  |
| `jwk` |  |
| `lastUpdated` |  |
| `lightningInvoice` |  |
| `lightningInvoiceExpiresAt` |  |
| `merchantDirectDebitMandateID` |  |
| `merchantID` |  |
| `merchantTokenDescription` |  |
| `notificationEmailAddresses` |  |
| `notificationRoleIDs` |  |
| `orderID` |  |
| `partialPaymentMethod` |  |
| `partialPaymentSteps` |  |
| `paymentAttempts` |  |
| `paymentInitiationID` |  |
| `paymentMethods` |  |
| `paymentProcessor` |  |
| `paymentRequests` |  |
| `payrunID` |  |
| `pispAccountID` |  |
| `priorityBankID` |  |
| `result` |  |
| `sandboxSettleDelayInSeconds` |  |
| `shippingAddress` |  |
| `status` |  |
| `successWebHookUrl` |  |
| `tags` |  |
| `title` |  |
| `tokenisedCards` |  |
| `transactions` |  |
| `useHostedPaymentPage` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/paymentrequests/{id}/directdebit`

#### PaymentRequestEvent

| Field | Description |
| --- | --- |
| `amount` |  |
| `applePayTransactionID` |  |
| `cardAuthorizationResponseID` |  |
| `cardExpiryMonth` |  |
| `cardExpiryYear` |  |
| `cardIssuer` |  |
| `cardIssuerCountry` |  |
| `cardLastFourDigits` |  |
| `cardRequestID` |  |
| `cardScheme` |  |
| `cardTokenCustomerID` |  |
| `cardTransactionID` |  |
| `currency` |  |
| `directDebitPaymentID` |  |
| `directDebitPaymentReference` |  |
| `drirectDebitMandateID` |  |
| `errorMessage` |  |
| `errorReason` |  |
| `eventType` |  |
| `id` |  |
| `inserted` |  |
| `lightningInvoice` |  |
| `lightningRHash` |  |
| `originUrl` |  |
| `paymentMethodType` |  |
| `paymentProcessorName` |  |
| `paymentRequestID` |  |
| `pispBankStatus` |  |
| `pispPaymentInitiationID` |  |
| `pispPaymentInstitutionName` |  |
| `pispPaymentServiceProviderID` |  |
| `pispRedirectUrl` |  |
| `reconciledTransactionID` |  |
| `refundPayoutID` |  |
| `status` |  |
| `walletName` |  |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/events`

#### PaymentRequestMetric

| Field | Description |
| --- | --- |

Operations: Load.

API path: `/api/v1/paymentrequests/metrics`

#### PaymentRequestMinimal

| Field | Description |
| --- | --- |
| `amount` |  |
| `amountPending` |  |
| `amountReceived` |  |
| `amountRefunded` |  |
| `callbackUrl` |  |
| `cardStripePaymentIntentSecret` |  |
| `countryCode` |  |
| `currency` |  |
| `customFieldsToDisplay` |  |
| `description` |  |
| `dueDate` |  |
| `fieldDisplaySettings` |  |
| `googlePayMerchantID` |  |
| `id` |  |
| `jwk` |  |
| `merchantID` |  |
| `merchantLogoUrlPng` |  |
| `merchantLogoUrlSvg` |  |
| `merchantName` |  |
| `merchantShortName` |  |
| `partialPaymentMethod` |  |
| `paymentAttempts` |  |
| `paymentMethodsList` |  |
| `paymentProcessor` |  |
| `paymentProcessorKey` |  |
| `pispError` |  |
| `priorityBankID` |  |
| `status` |  |
| `stripeAccountID` |  |
| `title` |  |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/minimal`

#### PaymentRequestResult

| Field | Description |
| --- | --- |
| `amount` |  |
| `amountPending` |  |
| `amountReceived` |  |
| `amountRefunded` |  |
| `currency` |  |
| `customerID` |  |
| `paymentRequestID` |  |
| `payments` |  |
| `pispAuthorizations` |  |
| `requestedAmount` |  |
| `result` |  |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/result`

#### Payout

| Field | Description |
| --- | --- |
| `accountID` |  |
| `allowIncomplete` |  |
| `amount` |  |
| `amountMinorUnits` |  |
| `approvePayoutUrl` |  |
| `approverID` |  |
| `authenticationMethods` |  |
| `authorisations` |  |
| `authorisersCompletedCount` |  |
| `authorisersRequiredCount` |  |
| `batchPayoutID` |  |
| `beneficiary` |  |
| `beneficiaryID` |  |
| `canAuthorise` |  |
| `canProcess` |  |
| `canUpdate` |  |
| `chargeBearer` |  |
| `createdBy` |  |
| `createdByEmailAddress` |  |
| `currency` |  |
| `currentUserID` |  |
| `description` |  |
| `destination` |  |
| `documents` |  |
| `events` |  |
| `failedPayouts` |  |
| `formattedAmount` |  |
| `formattedFxDestinationAmount` |  |
| `formattedSchedule` |  |
| `formattedScheduleDayOnly` |  |
| `formattedSourceAccountAvailableBalance` |  |
| `fxDestinationAmount` |  |
| `fxDestinationAmountMinorUnits` |  |
| `fxDestinationCurrency` |  |
| `fxQuoteExpiresAt` |  |
| `fxQuoteID` |  |
| `fxRate` |  |
| `fxUseDestinationAmount` |  |
| `hasCurrentUserAuthorised` |  |
| `id` |  |
| `inserted` |  |
| `invoiceID` |  |
| `isArchived` |  |
| `isFailed` |  |
| `isSettled` |  |
| `isSubmitted` |  |
| `lastUpdated` |  |
| `merchantID` |  |
| `merchantTokenDescription` |  |
| `nonce` |  |
| `paymentProcessor` |  |
| `paymentRail` |  |
| `payouts` |  |
| `payrunID` |  |
| `payrunName` |  |
| `reason` |  |
| `rule` |  |
| `scheduleDate` |  |
| `scheduled` |  |
| `sourceAccountAvailableBalance` |  |
| `sourceAccountAvailableBalanceMinorUnits` |  |
| `sourceAccountBic` |  |
| `sourceAccountCurrency` |  |
| `sourceAccountIban` |  |
| `sourceAccountIdentifier` |  |
| `sourceAccountName` |  |
| `sourceAccountNumber` |  |
| `sourceAccountSortcode` |  |
| `status` |  |
| `tagIds` |  |
| `tags` |  |
| `theirReference` |  |
| `topupPayrunID` |  |
| `transactedAmount` |  |
| `transactedFxAmount` |  |
| `transactedFxRate` |  |
| `type` |  |
| `userID` |  |
| `yourReference` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/payouts/batch/submit/{id}`

#### PayoutKeysetPage

| Field | Description |
| --- | --- |
| `accountID` |  |
| `amount` |  |
| `amountMinorUnits` |  |
| `approvePayoutUrl` |  |
| `approverID` |  |
| `authenticationMethods` |  |
| `authorisations` |  |
| `authorisersCompletedCount` |  |
| `authorisersRequiredCount` |  |
| `batchPayoutID` |  |
| `beneficiary` |  |
| `canAuthorise` |  |
| `canProcess` |  |
| `canUpdate` |  |
| `chargeBearer` |  |
| `createdBy` |  |
| `createdByEmailAddress` |  |
| `currency` |  |
| `currentUserID` |  |
| `description` |  |
| `destination` |  |
| `documents` |  |
| `events` |  |
| `formattedAmount` |  |
| `formattedFxDestinationAmount` |  |
| `formattedSchedule` |  |
| `formattedScheduleDayOnly` |  |
| `formattedSourceAccountAvailableBalance` |  |
| `fxDestinationAmount` |  |
| `fxDestinationAmountMinorUnits` |  |
| `fxDestinationCurrency` |  |
| `fxQuoteExpiresAt` |  |
| `fxQuoteID` |  |
| `fxRate` |  |
| `fxUseDestinationAmount` |  |
| `hasCurrentUserAuthorised` |  |
| `id` |  |
| `inserted` |  |
| `invoiceID` |  |
| `isArchived` |  |
| `isFailed` |  |
| `isSettled` |  |
| `isSubmitted` |  |
| `lastUpdated` |  |
| `merchantID` |  |
| `merchantTokenDescription` |  |
| `nonce` |  |
| `paymentProcessor` |  |
| `paymentRail` |  |
| `payrunID` |  |
| `payrunName` |  |
| `rule` |  |
| `scheduleDate` |  |
| `scheduled` |  |
| `sourceAccountAvailableBalance` |  |
| `sourceAccountAvailableBalanceMinorUnits` |  |
| `sourceAccountBic` |  |
| `sourceAccountCurrency` |  |
| `sourceAccountIban` |  |
| `sourceAccountIdentifier` |  |
| `sourceAccountName` |  |
| `sourceAccountNumber` |  |
| `sourceAccountSortcode` |  |
| `status` |  |
| `tags` |  |
| `theirReference` |  |
| `topupPayrunID` |  |
| `transactedAmount` |  |
| `transactedFxAmount` |  |
| `transactedFxRate` |  |
| `type` |  |
| `userID` |  |
| `yourReference` |  |

Operations: List.

API path: `/api/v1/accounts/{accountID}/payouts/failed`

#### PayoutMetric

| Field | Description |
| --- | --- |

Operations: Load.

API path: `/api/v1/payouts/metrics`

#### Payrun

| Field | Description |
| --- | --- |
| `authorisationDate` |  |
| `authorisations` |  |
| `authorisersCompletedCount` |  |
| `authorisersRequiredCount` |  |
| `batchPayoutID` |  |
| `canAuthorise` |  |
| `canDelete` |  |
| `canEdit` |  |
| `events` |  |
| `hasCurrentUserAuthorised` |  |
| `id` |  |
| `inserted` |  |
| `invoices` |  |
| `invoicesMinimal` |  |
| `isArchived` |  |
| `lastUpdated` |  |
| `lastUpdatedBy` |  |
| `merchantID` |  |
| `name` |  |
| `nonce` |  |
| `notes` |  |
| `payments` |  |
| `payouts` |  |
| `payoutsCount` |  |
| `reason` |  |
| `scheduleDate` |  |
| `scheduledDate` |  |
| `sourceAccounts` |  |
| `status` |  |
| `totalEur` |  |
| `totalGbp` |  |
| `totalUsd` |  |

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
| `contentType` |  |
| `contents` |  |
| `lastCompletedAt` |  |
| `merchantID` |  |
| `reportName` |  |
| `reportType` |  |
| `statementNumber` |  |

Operations: Load.

API path: `/api/v1/reports/{id}/result/{statementNumber}`

#### Role

| Field | Description |
| --- | --- |
| `failedRoles` |  |
| `roles` |  |

Operations: Create.

API path: `/api/v1/merchants/{merchantID}/roles/batchcreate`

#### Rule

| Field | Description |
| --- | --- |
| `account` |  |
| `accountID` |  |
| `approveUrl` |  |
| `approverID` |  |
| `authenticationMethods` |  |
| `authorisations` |  |
| `authorisersCompletedCount` |  |
| `authorisersRequiredCount` |  |
| `canAuthorise` |  |
| `createdBy` |  |
| `description` |  |
| `endAt` |  |
| `hasCurrentUserAuthorised` |  |
| `id` |  |
| `inserted` |  |
| `isDisabled` |  |
| `lastExecutedAt` |  |
| `lastRunAtTransactionDate` |  |
| `lastUpdated` |  |
| `merchantID` |  |
| `name` |  |
| `nonce` |  |
| `onApprovedWebHookUrl` |  |
| `onExecutionErrorWebHookUrl` |  |
| `onExecutionSuccessWebHookUrl` |  |
| `startAt` |  |
| `status` |  |
| `sweepAction` |  |
| `timeZoneId` |  |
| `triggerCronExpression` |  |
| `triggerOnPayIn` |  |
| `userID` |  |
| `webHookSecret` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/rules`

#### RuleEvent

| Field | Description |
| --- | --- |
| `errorMessage` |  |
| `id` |  |
| `inserted` |  |
| `isAuthoriseToEnable` |  |
| `message` |  |
| `rawResponse` |  |
| `ruleEventType` |  |
| `ruleID` |  |
| `user` |  |

Operations: List.

API path: `/api/v1/rules/{id}/events`

#### Tag

| Field | Description |
| --- | --- |
| `colourHex` |  |
| `description` |  |
| `id` |  |
| `merchantID` |  |
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
| `accountID` |  |
| `accountName` |  |
| `accountSequenceNumber` |  |
| `addressDetails` |  |
| `amount` |  |
| `amountMinorUnits` |  |
| `balance` |  |
| `balanceMinorUnits` |  |
| `bookingDateTime` |  |
| `chargeDetails` |  |
| `content` |  |
| `counterparty` |  |
| `counterpartySummary` |  |
| `currency` |  |
| `currencyExchange` |  |
| `date` |  |
| `description` |  |
| `enrichment` |  |
| `fxAmount` |  |
| `fxCurrency` |  |
| `fxRate` |  |
| `grossAmount` |  |
| `id` |  |
| `inserted` |  |
| `isoBankTransactionCode` |  |
| `merchant` |  |
| `merchantID` |  |
| `pageNumber` |  |
| `pageSize` |  |
| `payeeDetails` |  |
| `payerDetails` |  |
| `paymentRequestCustomFields` |  |
| `paymentRequestID` |  |
| `payoutID` |  |
| `proprietaryBankTransactionCode` |  |
| `rawReference` |  |
| `reference` |  |
| `ruleID` |  |
| `statementReferences` |  |
| `status` |  |
| `supplementaryData` |  |
| `tags` |  |
| `theirReference` |  |
| `totalPages` |  |
| `totalSize` |  |
| `transactionAmount` |  |
| `transactionDate` |  |
| `transactionInformation` |  |
| `transactionMutability` |  |
| `type` |  |
| `valueDateTime` |  |
| `virtualIBAN` |  |
| `yourReference` |  |

Operations: Create, List, Load, Remove.

API path: `/api/v1/transactions/{id}/tags`

#### User

| Field | Description |
| --- | --- |
| `clientSessionTimeouts` |  |
| `emailAddress` |  |
| `firstName` |  |
| `id` |  |
| `lastName` |  |
| `passkeyAdded` |  |
| `permissions` |  |
| `profile` |  |
| `rolesWithScope` |  |
| `twoFactorEnabled` |  |
| `userInviteID` |  |

Operations: List, Update.

API path: `/api/v1/user/{merchantID}/userspaged`

#### UserInvite

| Field | Description |
| --- | --- |
| `authorisationStatus` |  |
| `failedUserInvites` |  |
| `id` |  |
| `initialRoleID` |  |
| `inviteeEmailAddress` |  |
| `inviteeFirstName` |  |
| `inviteeLastName` |  |
| `inviterEmailAddress` |  |
| `inviterFirstName` |  |
| `inviterLastName` |  |
| `isAuthorised` |  |
| `isInviteeRegistered` |  |
| `lastInvited` |  |
| `merchantID` |  |
| `merchantName` |  |
| `message` |  |
| `registrationUrl` |  |
| `sendInviteEmail` |  |
| `status` |  |
| `user` |  |
| `userID` |  |
| `userInvites` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/userinvites/authorise/{id}`

#### Virtual

| Field | Description |
| --- | --- |
| `accountName` |  |
| `accountSupplierName` |  |
| `availableBalance` |  |
| `availableBalanceMinorUnits` |  |
| `balance` |  |
| `balanceMinorUnits` |  |
| `bankName` |  |
| `consentID` |  |
| `createdBy` |  |
| `createdByDisplayName` |  |
| `currency` |  |
| `defaultPaymentRail` |  |
| `displayName` |  |
| `expiryDate` |  |
| `externalAccountIcon` |  |
| `id` |  |
| `identifier` |  |
| `inserted` |  |
| `isArchived` |  |
| `isConnectedAccount` |  |
| `isDefault` |  |
| `isTrustAccount` |  |
| `isVirtual` |  |
| `lastTransaction` |  |
| `lastUpdated` |  |
| `merchantID` |  |
| `merchantName` |  |
| `name` |  |
| `physicalAccountID` |  |
| `rules` |  |
| `submittedPayoutsBalance` |  |
| `submittedPayoutsBalanceMinorUnits` |  |
| `summary` |  |
| `supplierSepaInstantStatus` |  |
| `xeroBankFeedConnectionStatus` |  |
| `xeroBankFeedLastSyncedAt` |  |
| `xeroBankFeedSyncLastFailedAt` |  |
| `xeroBankFeedSyncLastFailureReason` |  |
| `xeroBankFeedSyncStatus` |  |
| `xeroUnsynchronisedTransactionsCount` |  |

Operations: Create, Update.

API path: `/api/v1/accounts/{accountID}/virtual`

#### Webhook

| Field | Description |
| --- | --- |
| `destinationUrl` |  |
| `emailAddress` |  |
| `failedNotificationEmailAddress` |  |
| `id` |  |
| `isActive` |  |
| `merchantID` |  |
| `notificationMethod` |  |
| `resourceTypes` |  |
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
| `accountBalances` | `array` |  |
| `accountID` | `string` |  |
| `accountIdentifications` | `array` |  |
| `accountName` | `string` |  |
| `accountNames` | `array` |  |
| `accountSupplierName` | `string` |  |
| `accountType` | `string` |  |
| `availableBalance` | `float` |  |
| `availableBalanceMinorUnits` | `int` |  |
| `balance` | `float` |  |
| `balanceMinorUnits` | `int` |  |
| `bankName` | `string` |  |
| `consentID` | `string` |  |
| `consolidatedAccountInformation` | `array` |  |
| `createdBy` | `array` |  |
| `createdByDisplayName` | `string` |  |
| `currency` | `string` |  |
| `defaultPaymentRail` | `string` |  |
| `description` | `string` |  |
| `details` | `string` |  |
| `displayName` | `string` |  |
| `expiryDate` | `string` |  |
| `externalAccountIcon` | `string` |  |
| `format` | `string` |  |
| `fromDate` | `string` |  |
| `id` | `string` |  |
| `identifier` | `array` |  |
| `inserted` | `string` |  |
| `isArchived` | `bool` |  |
| `isConnectedAccount` | `bool` |  |
| `isDefault` | `bool` |  |
| `isTrustAccount` | `bool` |  |
| `isVirtual` | `bool` |  |
| `lastTransaction` | `array` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantName` | `string` |  |
| `nickname` | `string` |  |
| `physicalAccountID` | `string` |  |
| `roleIDs` | `array` |  |
| `rules` | `array` |  |
| `submittedPayoutsBalance` | `float` |  |
| `submittedPayoutsBalanceMinorUnits` | `int` |  |
| `summary` | `string` |  |
| `supplierPhysicalAccountID` | `string` |  |
| `supplierSepaInstantStatus` | `string` |  |
| `toDate` | `string` |  |
| `type` | `string` |  |
| `usageType` | `string` |  |
| `xeroBankFeedConnectionStatus` | `string` |  |
| `xeroBankFeedLastSyncedAt` | `string` |  |
| `xeroBankFeedSyncLastFailedAt` | `string` |  |
| `xeroBankFeedSyncLastFailureReason` | `string` |  |
| `xeroBankFeedSyncStatus` | `string` |  |
| `xeroUnsynchronisedTransactionsCount` | `int` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Account record (throws on error).
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
    "createdBy" => null, // array
    "identifier" => null, // array
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
| `approveUrl` | `string` |  |
| `id` | `string` |  |
| `payouts` | `array` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Batch record (throws on error).
$batch = $client->Batch()->load(["id" => "batch_id"]);
```

#### Example: Create

```php
$batch = $client->Batch()->create([
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
| `approvalCallbackUrl` | `string` |  |
| `authenticationMethods` | `array` |  |
| `authorisations` | `array` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `beneficiaries` | `array` |  |
| `beneficiaryEvents` | `array` |  |
| `canAuthorise` | `bool` |  |
| `canUpdate` | `bool` |  |
| `createdBy` | `array` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` |  |
| `destination` | `array` |  |
| `failedBeneficiaries` | `array` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isEnabled` | `bool` |  |
| `lastAuthorised` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `sourceAccountIDs` | `array` |  |
| `sourceAccounts` | `array` |  |
| `theirReference` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Beneficiary record (throws on error).
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
    "createdBy" => null, // array
    "currency" => null, // string
    "name" => null, // string
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
| `groupMembers` | `array` |  |
| `groupName` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |

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
| `authorizedAmount` | `string` |  |
| `currencyCode` | `string` |  |
| `isPayerAuthenticationRequired` | `bool` |  |
| `isSoftDecline` | `bool` |  |
| `payerAuthenticationAccessToken` | `string` |  |
| `payerAuthenticationMerchantData` | `string` |  |
| `payerAuthenticationUrl` | `string` |  |
| `payerAuthenticationWindowHeight` | `int` |  |
| `payerAuthenticationWindowWidth` | `int` |  |
| `paymentRequestCallbackUrl` | `string` |  |
| `paymentRequestID` | `string` |  |
| `requestID` | `string` |  |
| `responseCode` | `string` |  |
| `responseType` | `string` |  |
| `status` | `string` |  |
| `threeDSRedirectUrl` | `string` |  |
| `transactionID` | `string` |  |

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
| `cardType` | `string` |  |
| `customerEmailAddress` | `string` |  |
| `expiryMonth` | `string` |  |
| `expiryYear` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lastFourDigits` | `string` |  |
| `lastUpdated` | `string` |  |
| `maskedCardNumber` | `string` |  |
| `merchantID` | `string` |  |
| `paymentRequestID` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the CardCustomerToken record (throws on error).
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
| `authorizedAmount` | `string` |  |
| `currencyCode` | `string` |  |
| `isPayerAuthenticationRequired` | `bool` |  |
| `isSoftDecline` | `bool` |  |
| `payerAuthenticationAccessToken` | `string` |  |
| `payerAuthenticationMerchantData` | `string` |  |
| `payerAuthenticationUrl` | `string` |  |
| `payerAuthenticationWindowHeight` | `int` |  |
| `payerAuthenticationWindowWidth` | `int` |  |
| `paymentRequestCallbackUrl` | `string` |  |
| `paymentRequestID` | `string` |  |
| `requestID` | `string` |  |
| `responseCode` | `string` |  |
| `responseType` | `string` |  |
| `status` | `string` |  |
| `threeDSRedirectUrl` | `string` |  |
| `transactionID` | `string` |  |

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
// load() returns the ENTITY — call data_get() for the CardPublicKey record (throws on error).
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
| `authorisationUrl` | `string` |  |
| `callbackUrl` | `string` |  |
| `consentID` | `string` |  |
| `emailAddress` | `string` |  |
| `expiryDate` | `string` |  |
| `failureCallbackUrl` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `institutionID` | `string` |  |
| `isConnectedAccounts` | `bool` |  |
| `isEnabled` | `bool` |  |
| `merchantID` | `string` |  |
| `provider` | `string` |  |
| `successWebHookUrl` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Consent record (throws on error).
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
| `decimals` | `int` |  |
| `isFiat` | `bool` |  |
| `iso4217AlphaCode` | `string` |  |
| `iso4217NumericCode` | `string` |  |
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
| `failedSubmissions` | `array` |  |
| `successfulSubmissions` | `array` |  |

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
| `destinationCurrency` | `string` |  |
| `exchangeRate` | `float` |  |
| `expiryTime` | `string` |  |
| `quoteID` | `string` |  |
| `sourceCurrency` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the FxRate record (throws on error).
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
| `paymentRequestID` | `string` |  |
| `responseType` | `string` |  |

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
| `accountNumber` | `string` |  |
| `addressLine1` | `string` |  |
| `addressLine2` | `string` |  |
| `approvedAt` | `string` |  |
| `city` | `string` |  |
| `countryCode` | `string` |  |
| `currency` | `string` |  |
| `customerAccountNumber` | `string` |  |
| `customerCity` | `string` |  |
| `customerCountryCode` | `string` |  |
| `customerCountryName` | `string` |  |
| `customerEmailAddress` | `string` |  |
| `customerFirstName` | `string` |  |
| `customerIban` | `string` |  |
| `customerLastName` | `string` |  |
| `customerSortCode` | `string` |  |
| `emailAddress` | `string` |  |
| `firstName` | `string` |  |
| `iban` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isRecurring` | `bool` |  |
| `lastName` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `postalCode` | `string` |  |
| `reference` | `string` |  |
| `sortCode` | `string` |  |
| `status` | `string` |  |
| `supplierBankAccountID` | `string` |  |
| `supplierCustomerID` | `string` |  |
| `supplierMandateID` | `string` |  |
| `supplierName` | `string` |  |
| `supplierStatus` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Mandate record (throws on error).
$mandate = $client->Mandate()->load(["id" => "mandate_id"]);
```

#### Example: Create

```php
$mandate = $client->Mandate()->create([
    "addressLine1" => null, // string
    "city" => null, // string
    "countryCode" => null, // string
    "emailAddress" => null, // string
    "firstName" => null, // string
    "lastName" => null, // string
    "postalCode" => null, // string
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
| `accountCurrencies` | `array` |  |
| `canHaveTrustAccounts` | `bool` |  |
| `cardPaymentProcessor` | `string` |  |
| `companyID` | `string` |  |
| `displayQrOnHostedPay` | `bool` |  |
| `hostedPayVersion` | `int` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isBlocked` | `bool` |  |
| `isExited` | `bool` |  |
| `isSuspended` | `bool` |  |
| `jurisdiction` | `string` |  |
| `logoUrlPng` | `string` |  |
| `logoUrlSvg` | `string` |  |
| `merchantCategoryCode` | `string` |  |
| `name` | `string` |  |
| `notes` | `string` |  |
| `parentMerchant` | `array` |  |
| `paymentAccountLimit` | `int` |  |
| `paymentAccounts` | `array` |  |
| `reason` | `string` |  |
| `shortName` | `string` |  |
| `supportedPaymentMethodsList` | `array` |  |
| `suspensionReason` | `string` |  |
| `tags` | `array` |  |
| `timeZoneId` | `string` |  |
| `tradingName` | `string` |  |
| `webHookLimit` | `int` |  |
| `yourRoleName` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Merchant record (throws on error).
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
| `amountLower` | `float` |  |
| `amountUpper` | `float` |  |
| `authorisationType` | `string` |  |
| `beneficiariesOnly` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lastEditorCantAuthorise` | `bool` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `numberOfAuthorisers` | `int` |  |
| `roleSettings` | `array` |  |

#### Example: List

```php
// list() returns an array of MerchantAuthorisationSetting records (throws on error).
$merchant_authorisation_settings = $client->MerchantAuthorisationSetting()->list();
```


### MerchantDirectDebitMandatePage

Create an instance: `$merchant_direct_debit_mandate_page = $client->MerchantDirectDebitMandatePage();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approvedAt` | `string` |  |
| `currency` | `string` |  |
| `customerAccountNumber` | `string` |  |
| `customerCity` | `string` |  |
| `customerCountryCode` | `string` |  |
| `customerCountryName` | `string` |  |
| `customerEmailAddress` | `string` |  |
| `customerFirstName` | `string` |  |
| `customerIban` | `string` |  |
| `customerLastName` | `string` |  |
| `customerSortCode` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isRecurring` | `bool` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `reference` | `string` |  |
| `status` | `string` |  |
| `supplierBankAccountID` | `string` |  |
| `supplierCustomerID` | `string` |  |
| `supplierMandateID` | `string` |  |
| `supplierName` | `string` |  |
| `supplierStatus` | `string` |  |

#### Example: List

```php
// list() returns an array of MerchantDirectDebitMandatePage records (throws on error).
$merchant_direct_debit_mandate_pages = $client->MerchantDirectDebitMandatePage()->list();
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
| `bankCountryCodes` | `array` |  |
| `bankID` | `string` |  |
| `bankName` | `string` |  |
| `businessInstitutionID` | `string` |  |
| `currency` | `string` |  |
| `logo` | `string` |  |
| `message` | `string` |  |
| `messageImageUrl` | `string` |  |
| `order` | `int` |  |
| `personalInstitutionID` | `string` |  |
| `processor` | `string` |  |
| `warningHeading` | `string` |  |
| `warningMessage` | `string` |  |

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
| `bankPaymentOptions` | `array` |  |
| `cardPaymentAddressOptions` | `array` |  |
| `cardPaymentCaptureOptions` | `array` |  |
| `customFields` | `array` |  |
| `defaultFields` | `array` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `name` | `string` |  |
| `notificationOptions` | `array` |  |
| `paymentMethods` | `array` |  |
| `paymentTerms` | `array` |  |
| `priorityBankOptions` | `array` |  |
| `template` | `array` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the MerchantPaymentRequestTemplate record (throws on error).
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
| `authenticationMethods` | `array` |  |
| `authorisations` | `array` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `canAuthorise` | `bool` |  |
| `description` | `string` |  |
| `expiresAt` | `string` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `hmacAlgorithm` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `ipAddressWhitelist` | `string` |  |
| `isArchived` | `bool` |  |
| `isEnabled` | `bool` |  |
| `lastAuthorised` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `nonce` | `string` |  |
| `permissionTypes` | `array` |  |
| `requestSignatureVersion` | `int` |  |
| `sharedSecretAlgorithm` | `string` |  |
| `sharedSecretBase64` | `string` |  |
| `token` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the MerchantToken record (throws on error).
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
// load() returns the ENTITY — call data_get() for the Metadata record (throws on error).
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
| `buildVersion` | `int` |  |
| `majorVersion` | `int` |  |
| `minorVersion` | `int` |  |
| `releaseName` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the NoFrixionVersion record (throws on error).
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
| `accountName` | `string` |  |
| `accountNumber` | `string` |  |
| `iban` | `string` |  |
| `payeeVerifiedAccountName` | `string` |  |
| `result` | `string` |  |
| `secondaryIdentification` | `string` |  |
| `sortCode` | `string` |  |

#### Example: Create

```php
$payeeverification = $client->Payeeverification()->create([
    "accountName" => null, // string
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
| `addresses` | `array` |  |
| `amount` | `float` |  |
| `amountPending` | `float` |  |
| `amountReceived` | `float` |  |
| `amountRefunded` | `float` |  |
| `autoSendReceipt` | `bool` |  |
| `baseOriginUrl` | `string` |  |
| `callbackUrl` | `string` |  |
| `cardAuthorizeOnly` | `bool` |  |
| `cardCreateToken` | `bool` |  |
| `cardCreateTokenMode` | `string` |  |
| `cardIgnoreCVN` | `bool` |  |
| `cardNoPayerAuthentication` | `bool` |  |
| `cardProcessorMerchantID` | `string` |  |
| `cardStripePaymentIntentID` | `string` |  |
| `cardStripePaymentIntentSecret` | `string` |  |
| `cardTransmitRawDetails` | `bool` |  |
| `createdByUser` | `array` |  |
| `currency` | `string` |  |
| `customFields` | `array` |  |
| `customerEmailAddress` | `string` |  |
| `customerID` | `string` |  |
| `customerName` | `string` |  |
| `description` | `string` |  |
| `destinationAccount` | `array` |  |
| `directDebitPayment` | `array` |  |
| `dueDate` | `string` |  |
| `events` | `array` |  |
| `failureCallbackUrl` | `string` |  |
| `fieldDisplaySettings` | `array` |  |
| `formattedAmount` | `string` |  |
| `hostedPayCheckoutUrl` | `string` |  |
| `id` | `string` |  |
| `ignoreAddressVerification` | `bool` |  |
| `inserted` | `string` |  |
| `insertedSortable` | `string` |  |
| `isArchived` | `bool` |  |
| `jwk` | `string` |  |
| `lastUpdated` | `string` |  |
| `lightningInvoice` | `string` |  |
| `lightningInvoiceExpiresAt` | `string` |  |
| `merchantDirectDebitMandateID` | `string` |  |
| `merchantID` | `string` |  |
| `merchantTokenDescription` | `string` |  |
| `notificationEmailAddresses` | `string` |  |
| `notificationRoleIDs` | `array` |  |
| `orderID` | `string` |  |
| `partialPaymentMethod` | `string` |  |
| `partialPaymentSteps` | `string` |  |
| `paymentAttempts` | `array` |  |
| `paymentMethods` | `array` |  |
| `paymentProcessor` | `string` |  |
| `payrunID` | `string` |  |
| `pispAccountID` | `string` |  |
| `priorityBankID` | `string` |  |
| `result` | `array` |  |
| `sandboxSettleDelayInSeconds` | `int` |  |
| `shippingAddress` | `array` |  |
| `shippingAddressCity` | `string` |  |
| `shippingAddressCountryCode` | `string` |  |
| `shippingAddressCounty` | `string` |  |
| `shippingAddressLine1` | `string` |  |
| `shippingAddressLine2` | `string` |  |
| `shippingAddressPostCode` | `string` |  |
| `shippingEmail` | `string` |  |
| `shippingFirstName` | `string` |  |
| `shippingLastName` | `string` |  |
| `shippingPhone` | `string` |  |
| `status` | `string` |  |
| `successWebHookUrl` | `string` |  |
| `tagIds` | `array` |  |
| `tags` | `array` |  |
| `title` | `string` |  |
| `tokenisedCards` | `array` |  |
| `transactions` | `array` |  |
| `useHostedPaymentPage` | `bool` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Payment record (throws on error).
$payment = $client->Payment()->load(["id" => "payment_id"]);
```

#### Example: Create

```php
$payment = $client->Payment()->create([
    "createdByUser" => null, // array
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
| `accountName` | `string` |  |
| `accountSupplierName` | `string` |  |
| `availableBalance` | `float` |  |
| `availableBalanceMinorUnits` | `int` |  |
| `balance` | `float` |  |
| `balanceMinorUnits` | `int` |  |
| `bankName` | `string` |  |
| `consentID` | `string` |  |
| `createdBy` | `array` |  |
| `createdByDisplayName` | `string` |  |
| `currency` | `string` |  |
| `defaultPaymentRail` | `string` |  |
| `displayName` | `string` |  |
| `expiryDate` | `string` |  |
| `externalAccountIcon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `array` |  |
| `inserted` | `string` |  |
| `isArchived` | `bool` |  |
| `isConnectedAccount` | `bool` |  |
| `isDefault` | `bool` |  |
| `isTrustAccount` | `bool` |  |
| `isVirtual` | `bool` |  |
| `lastTransaction` | `array` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantName` | `string` |  |
| `physicalAccountID` | `string` |  |
| `rules` | `array` |  |
| `submittedPayoutsBalance` | `float` |  |
| `submittedPayoutsBalanceMinorUnits` | `int` |  |
| `summary` | `string` |  |
| `supplierSepaInstantStatus` | `string` |  |
| `xeroBankFeedConnectionStatus` | `string` |  |
| `xeroBankFeedLastSyncedAt` | `string` |  |
| `xeroBankFeedSyncLastFailedAt` | `string` |  |
| `xeroBankFeedSyncLastFailureReason` | `string` |  |
| `xeroBankFeedSyncStatus` | `string` |  |
| `xeroUnsynchronisedTransactionsCount` | `int` |  |

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
| `accountName` | `string` |  |
| `availableBalance` | `float` |  |
| `balance` | `float` |  |
| `balanceMinorUnits` | `int` |  |
| `currency` | `string` |  |
| `id` | `string` |  |
| `identifier` | `array` |  |
| `isArchived` | `bool` |  |
| `isConnectedAccount` | `bool` |  |
| `merchantID` | `string` |  |
| `submittedPayoutsBalance` | `float` |  |

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
| `paymentInitiationID` | `string` |  |
| `paymentRequestCallbackUrl` | `string` |  |
| `paymentRequestID` | `string` |  |
| `redirectUrl` | `string` |  |
| `responseType` | `string` |  |
| `specificErrorMessage` | `string` |  |

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
| `addresses` | `array` |  |
| `amount` | `float` |  |
| `amountPending` | `float` |  |
| `amountReceived` | `float` |  |
| `amountRefunded` | `float` |  |
| `autoSendReceipt` | `bool` |  |
| `baseOriginUrl` | `string` |  |
| `callbackUrl` | `string` |  |
| `cardAuthorizeOnly` | `bool` |  |
| `cardCreateToken` | `bool` |  |
| `cardCreateTokenMode` | `string` |  |
| `cardIgnoreCVN` | `bool` |  |
| `cardProcessorMerchantID` | `string` |  |
| `cardStripePaymentIntentID` | `string` |  |
| `cardStripePaymentIntentSecret` | `string` |  |
| `createdByUser` | `array` |  |
| `currency` | `string` |  |
| `customFields` | `array` |  |
| `customerEmailAddress` | `string` |  |
| `customerID` | `string` |  |
| `customerName` | `string` |  |
| `description` | `string` |  |
| `destinationAccount` | `array` |  |
| `directDebitPayment` | `array` |  |
| `doSimulateSettlementFailure` | `bool` |  |
| `dueDate` | `string` |  |
| `errorDescription` | `string` |  |
| `events` | `array` |  |
| `failedPaymentRequests` | `array` |  |
| `failureCallbackUrl` | `string` |  |
| `fieldDisplaySettings` | `array` |  |
| `formattedAmount` | `string` |  |
| `hostedPayCheckoutUrl` | `string` |  |
| `id` | `string` |  |
| `ignoreAddressVerification` | `bool` |  |
| `inserted` | `string` |  |
| `insertedSortable` | `string` |  |
| `institution` | `string` |  |
| `isArchived` | `bool` |  |
| `jwk` | `string` |  |
| `lastUpdated` | `string` |  |
| `lightningInvoice` | `string` |  |
| `lightningInvoiceExpiresAt` | `string` |  |
| `merchantDirectDebitMandateID` | `string` |  |
| `merchantID` | `string` |  |
| `merchantTokenDescription` | `string` |  |
| `notificationEmailAddresses` | `string` |  |
| `notificationRoleIDs` | `array` |  |
| `orderID` | `string` |  |
| `partialPaymentMethod` | `string` |  |
| `partialPaymentSteps` | `string` |  |
| `paymentAttempts` | `array` |  |
| `paymentInitiationID` | `string` |  |
| `paymentMethods` | `array` |  |
| `paymentProcessor` | `string` |  |
| `paymentRequests` | `array` |  |
| `payrunID` | `string` |  |
| `pispAccountID` | `string` |  |
| `priorityBankID` | `string` |  |
| `result` | `array` |  |
| `sandboxSettleDelayInSeconds` | `int` |  |
| `shippingAddress` | `array` |  |
| `status` | `string` |  |
| `successWebHookUrl` | `string` |  |
| `tags` | `array` |  |
| `title` | `string` |  |
| `tokenisedCards` | `array` |  |
| `transactions` | `array` |  |
| `useHostedPaymentPage` | `bool` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the PaymentRequest record (throws on error).
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
    "createdByUser" => null, // array
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
| `applePayTransactionID` | `string` |  |
| `cardAuthorizationResponseID` | `string` |  |
| `cardExpiryMonth` | `int` |  |
| `cardExpiryYear` | `int` |  |
| `cardIssuer` | `string` |  |
| `cardIssuerCountry` | `string` |  |
| `cardLastFourDigits` | `string` |  |
| `cardRequestID` | `string` |  |
| `cardScheme` | `string` |  |
| `cardTokenCustomerID` | `string` |  |
| `cardTransactionID` | `string` |  |
| `currency` | `string` |  |
| `directDebitPaymentID` | `string` |  |
| `directDebitPaymentReference` | `string` |  |
| `drirectDebitMandateID` | `string` |  |
| `errorMessage` | `string` |  |
| `errorReason` | `string` |  |
| `eventType` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lightningInvoice` | `string` |  |
| `lightningRHash` | `string` |  |
| `originUrl` | `string` |  |
| `paymentMethodType` | `string` |  |
| `paymentProcessorName` | `string` |  |
| `paymentRequestID` | `string` |  |
| `pispBankStatus` | `string` |  |
| `pispPaymentInitiationID` | `string` |  |
| `pispPaymentInstitutionName` | `string` |  |
| `pispPaymentServiceProviderID` | `string` |  |
| `pispRedirectUrl` | `string` |  |
| `reconciledTransactionID` | `string` |  |
| `refundPayoutID` | `string` |  |
| `status` | `string` |  |
| `walletName` | `string` |  |

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

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the PaymentRequestMetric record (throws on error).
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
| `amountPending` | `float` |  |
| `amountReceived` | `float` |  |
| `amountRefunded` | `float` |  |
| `callbackUrl` | `string` |  |
| `cardStripePaymentIntentSecret` | `string` |  |
| `countryCode` | `string` |  |
| `currency` | `string` |  |
| `customFieldsToDisplay` | `array` |  |
| `description` | `string` |  |
| `dueDate` | `string` |  |
| `fieldDisplaySettings` | `array` |  |
| `googlePayMerchantID` | `string` |  |
| `id` | `string` |  |
| `jwk` | `string` |  |
| `merchantID` | `string` |  |
| `merchantLogoUrlPng` | `string` |  |
| `merchantLogoUrlSvg` | `string` |  |
| `merchantName` | `string` |  |
| `merchantShortName` | `string` |  |
| `partialPaymentMethod` | `string` |  |
| `paymentAttempts` | `array` |  |
| `paymentMethodsList` | `array` |  |
| `paymentProcessor` | `string` |  |
| `paymentProcessorKey` | `string` |  |
| `pispError` | `string` |  |
| `priorityBankID` | `string` |  |
| `status` | `string` |  |
| `stripeAccountID` | `string` |  |
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
| `amountPending` | `float` |  |
| `amountReceived` | `float` |  |
| `amountRefunded` | `float` |  |
| `currency` | `string` |  |
| `customerID` | `string` |  |
| `paymentRequestID` | `string` |  |
| `payments` | `array` |  |
| `pispAuthorizations` | `array` |  |
| `requestedAmount` | `float` |  |
| `result` | `string` |  |

#### Example: List

```php
// list() returns an array of PaymentRequestResult records (throws on error).
$payment_request_results = $client->PaymentRequestResult()->list();
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
| `accountID` | `string` |  |
| `allowIncomplete` | `bool` |  |
| `amount` | `float` |  |
| `amountMinorUnits` | `int` |  |
| `approvePayoutUrl` | `string` |  |
| `approverID` | `string` |  |
| `authenticationMethods` | `array` |  |
| `authorisations` | `array` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `batchPayoutID` | `string` |  |
| `beneficiary` | `array` |  |
| `beneficiaryID` | `string` |  |
| `canAuthorise` | `bool` |  |
| `canProcess` | `bool` |  |
| `canUpdate` | `bool` |  |
| `chargeBearer` | `string` |  |
| `createdBy` | `string` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` |  |
| `currentUserID` | `string` |  |
| `description` | `string` |  |
| `destination` | `array` |  |
| `documents` | `array` |  |
| `events` | `array` |  |
| `failedPayouts` | `array` |  |
| `formattedAmount` | `string` |  |
| `formattedFxDestinationAmount` | `string` |  |
| `formattedSchedule` | `string` |  |
| `formattedScheduleDayOnly` | `string` |  |
| `formattedSourceAccountAvailableBalance` | `string` |  |
| `fxDestinationAmount` | `float` |  |
| `fxDestinationAmountMinorUnits` | `int` |  |
| `fxDestinationCurrency` | `string` |  |
| `fxQuoteExpiresAt` | `string` |  |
| `fxQuoteID` | `string` |  |
| `fxRate` | `float` |  |
| `fxUseDestinationAmount` | `bool` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoiceID` | `string` |  |
| `isArchived` | `bool` |  |
| `isFailed` | `bool` |  |
| `isSettled` | `bool` |  |
| `isSubmitted` | `bool` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantTokenDescription` | `string` |  |
| `nonce` | `string` |  |
| `paymentProcessor` | `string` |  |
| `paymentRail` | `string` |  |
| `payouts` | `array` |  |
| `payrunID` | `string` |  |
| `payrunName` | `string` |  |
| `reason` | `string` |  |
| `rule` | `array` |  |
| `scheduleDate` | `string` |  |
| `scheduled` | `bool` |  |
| `sourceAccountAvailableBalance` | `float` |  |
| `sourceAccountAvailableBalanceMinorUnits` | `int` |  |
| `sourceAccountBic` | `string` |  |
| `sourceAccountCurrency` | `string` |  |
| `sourceAccountIban` | `string` |  |
| `sourceAccountIdentifier` | `array` |  |
| `sourceAccountName` | `string` |  |
| `sourceAccountNumber` | `string` |  |
| `sourceAccountSortcode` | `string` |  |
| `status` | `string` |  |
| `tagIds` | `array` |  |
| `tags` | `array` |  |
| `theirReference` | `string` |  |
| `topupPayrunID` | `string` |  |
| `transactedAmount` | `float` |  |
| `transactedFxAmount` | `float` |  |
| `transactedFxRate` | `float` |  |
| `type` | `string` |  |
| `userID` | `string` |  |
| `yourReference` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Payout record (throws on error).
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
    "beneficiary" => null, // array
    "sourceAccountIdentifier" => null, // array
]);
```


### PayoutKeysetPage

Create an instance: `$payout_keyset_page = $client->PayoutKeysetPage();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountID` | `string` |  |
| `amount` | `float` |  |
| `amountMinorUnits` | `int` |  |
| `approvePayoutUrl` | `string` |  |
| `approverID` | `string` |  |
| `authenticationMethods` | `array` |  |
| `authorisations` | `array` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `batchPayoutID` | `string` |  |
| `beneficiary` | `array` |  |
| `canAuthorise` | `bool` |  |
| `canProcess` | `bool` |  |
| `canUpdate` | `bool` |  |
| `chargeBearer` | `string` |  |
| `createdBy` | `string` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` |  |
| `currentUserID` | `string` |  |
| `description` | `string` |  |
| `destination` | `array` |  |
| `documents` | `array` |  |
| `events` | `array` |  |
| `formattedAmount` | `string` |  |
| `formattedFxDestinationAmount` | `string` |  |
| `formattedSchedule` | `string` |  |
| `formattedScheduleDayOnly` | `string` |  |
| `formattedSourceAccountAvailableBalance` | `string` |  |
| `fxDestinationAmount` | `float` |  |
| `fxDestinationAmountMinorUnits` | `int` |  |
| `fxDestinationCurrency` | `string` |  |
| `fxQuoteExpiresAt` | `string` |  |
| `fxQuoteID` | `string` |  |
| `fxRate` | `float` |  |
| `fxUseDestinationAmount` | `bool` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoiceID` | `string` |  |
| `isArchived` | `bool` |  |
| `isFailed` | `bool` |  |
| `isSettled` | `bool` |  |
| `isSubmitted` | `bool` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantTokenDescription` | `string` |  |
| `nonce` | `string` |  |
| `paymentProcessor` | `string` |  |
| `paymentRail` | `string` |  |
| `payrunID` | `string` |  |
| `payrunName` | `string` |  |
| `rule` | `array` |  |
| `scheduleDate` | `string` |  |
| `scheduled` | `bool` |  |
| `sourceAccountAvailableBalance` | `float` |  |
| `sourceAccountAvailableBalanceMinorUnits` | `int` |  |
| `sourceAccountBic` | `string` |  |
| `sourceAccountCurrency` | `string` |  |
| `sourceAccountIban` | `string` |  |
| `sourceAccountIdentifier` | `array` |  |
| `sourceAccountName` | `string` |  |
| `sourceAccountNumber` | `string` |  |
| `sourceAccountSortcode` | `string` |  |
| `status` | `string` |  |
| `tags` | `array` |  |
| `theirReference` | `string` |  |
| `topupPayrunID` | `string` |  |
| `transactedAmount` | `float` |  |
| `transactedFxAmount` | `float` |  |
| `transactedFxRate` | `float` |  |
| `type` | `string` |  |
| `userID` | `string` |  |
| `yourReference` | `string` |  |

#### Example: List

```php
// list() returns an array of PayoutKeysetPage records (throws on error).
$payout_keyset_pages = $client->PayoutKeysetPage()->list();
```


### PayoutMetric

Create an instance: `$payout_metric = $client->PayoutMetric();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the PayoutMetric record (throws on error).
$payout_metric = $client->PayoutMetric()->load();
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
| `authorisationDate` | `string` |  |
| `authorisations` | `array` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `batchPayoutID` | `string` |  |
| `canAuthorise` | `bool` |  |
| `canDelete` | `bool` |  |
| `canEdit` | `bool` |  |
| `events` | `array` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoices` | `array` |  |
| `invoicesMinimal` | `array` |  |
| `isArchived` | `bool` |  |
| `lastUpdated` | `string` |  |
| `lastUpdatedBy` | `array` |  |
| `merchantID` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `notes` | `string` |  |
| `payments` | `array` |  |
| `payouts` | `array` |  |
| `payoutsCount` | `int` |  |
| `reason` | `string` |  |
| `scheduleDate` | `string` |  |
| `scheduledDate` | `string` |  |
| `sourceAccounts` | `array` |  |
| `status` | `string` |  |
| `totalEur` | `float` |  |
| `totalGbp` | `float` |  |
| `totalUsd` | `float` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Payrun record (throws on error).
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
    "lastUpdatedBy" => null, // array
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
| `contentType` | `string` |  |
| `contents` | `string` |  |
| `lastCompletedAt` | `string` |  |
| `merchantID` | `string` |  |
| `reportName` | `string` |  |
| `reportType` | `string` |  |
| `statementNumber` | `int` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the ReportResult record (throws on error).
$report_result = $client->ReportResult()->load(["id" => 1, "report_id" => "report_id"]);
```


### Role

Create an instance: `$role = $client->Role();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failedRoles` | `array` |  |
| `roles` | `array` |  |

#### Example: Create

```php
$role = $client->Role()->create([
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
| `accountID` | `string` |  |
| `approveUrl` | `string` |  |
| `approverID` | `string` |  |
| `authenticationMethods` | `array` |  |
| `authorisations` | `array` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `canAuthorise` | `bool` |  |
| `createdBy` | `array` |  |
| `description` | `string` |  |
| `endAt` | `string` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isDisabled` | `bool` |  |
| `lastExecutedAt` | `string` |  |
| `lastRunAtTransactionDate` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `onApprovedWebHookUrl` | `string` |  |
| `onExecutionErrorWebHookUrl` | `string` |  |
| `onExecutionSuccessWebHookUrl` | `string` |  |
| `startAt` | `string` |  |
| `status` | `string` |  |
| `sweepAction` | `array` |  |
| `timeZoneId` | `string` |  |
| `triggerCronExpression` | `string` |  |
| `triggerOnPayIn` | `bool` |  |
| `userID` | `string` |  |
| `webHookSecret` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Rule record (throws on error).
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
    "createdBy" => null, // array
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
| `errorMessage` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isAuthoriseToEnable` | `bool` |  |
| `message` | `string` |  |
| `rawResponse` | `string` |  |
| `ruleEventType` | `string` |  |
| `ruleID` | `string` |  |
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
| `colourHex` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `merchantID` | `string` |  |
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
    "merchantID" => null, // string
    "name" => null, // string
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
| `accountID` | `string` |  |
| `accountName` | `string` |  |
| `accountSequenceNumber` | `int` |  |
| `addressDetails` | `array` |  |
| `amount` | `float` |  |
| `amountMinorUnits` | `int` |  |
| `balance` | `float` |  |
| `balanceMinorUnits` | `int` |  |
| `bookingDateTime` | `string` |  |
| `chargeDetails` | `array` |  |
| `content` | `array` |  |
| `counterparty` | `array` |  |
| `counterpartySummary` | `string` |  |
| `currency` | `string` |  |
| `currencyExchange` | `array` |  |
| `date` | `string` |  |
| `description` | `string` |  |
| `enrichment` | `array` |  |
| `fxAmount` | `float` |  |
| `fxCurrency` | `string` |  |
| `fxRate` | `float` |  |
| `grossAmount` | `array` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isoBankTransactionCode` | `array` |  |
| `merchant` | `array` |  |
| `merchantID` | `string` |  |
| `pageNumber` | `int` |  |
| `pageSize` | `int` |  |
| `payeeDetails` | `array` |  |
| `payerDetails` | `array` |  |
| `paymentRequestCustomFields` | `array` |  |
| `paymentRequestID` | `string` |  |
| `payoutID` | `string` |  |
| `proprietaryBankTransactionCode` | `array` |  |
| `rawReference` | `string` |  |
| `reference` | `string` |  |
| `ruleID` | `string` |  |
| `statementReferences` | `array` |  |
| `status` | `string` |  |
| `supplementaryData` | `mixed` |  |
| `tags` | `array` |  |
| `theirReference` | `string` |  |
| `totalPages` | `int` |  |
| `totalSize` | `int` |  |
| `transactionAmount` | `array` |  |
| `transactionDate` | `string` |  |
| `transactionInformation` | `array` |  |
| `transactionMutability` | `string` |  |
| `type` | `string` |  |
| `valueDateTime` | `string` |  |
| `virtualIBAN` | `string` |  |
| `yourReference` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Transaction record (throws on error).
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
    "grossAmount" => null, // array
    "payeeDetails" => null, // array
    "payerDetails" => null, // array
    "transactionAmount" => null, // array
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
| `clientSessionTimeouts` | `array` |  |
| `emailAddress` | `string` |  |
| `firstName` | `string` |  |
| `id` | `string` |  |
| `lastName` | `string` |  |
| `passkeyAdded` | `bool` |  |
| `permissions` | `array` |  |
| `profile` | `string` |  |
| `rolesWithScope` | `array` |  |
| `twoFactorEnabled` | `bool` |  |
| `userInviteID` | `string` |  |

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
| `authorisationStatus` | `array` |  |
| `failedUserInvites` | `array` |  |
| `id` | `string` |  |
| `initialRoleID` | `string` |  |
| `inviteeEmailAddress` | `string` |  |
| `inviteeFirstName` | `string` |  |
| `inviteeLastName` | `string` |  |
| `inviterEmailAddress` | `string` |  |
| `inviterFirstName` | `string` |  |
| `inviterLastName` | `string` |  |
| `isAuthorised` | `bool` |  |
| `isInviteeRegistered` | `bool` |  |
| `lastInvited` | `string` |  |
| `merchantID` | `string` |  |
| `merchantName` | `string` |  |
| `message` | `string` |  |
| `registrationUrl` | `string` |  |
| `sendInviteEmail` | `bool` |  |
| `status` | `string` |  |
| `user` | `array` |  |
| `userID` | `string` |  |
| `userInvites` | `array` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the UserInvite record (throws on error).
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
    "user" => null, // array
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
| `accountName` | `string` |  |
| `accountSupplierName` | `string` |  |
| `availableBalance` | `float` |  |
| `availableBalanceMinorUnits` | `int` |  |
| `balance` | `float` |  |
| `balanceMinorUnits` | `int` |  |
| `bankName` | `string` |  |
| `consentID` | `string` |  |
| `createdBy` | `array` |  |
| `createdByDisplayName` | `string` |  |
| `currency` | `string` |  |
| `defaultPaymentRail` | `string` |  |
| `displayName` | `string` |  |
| `expiryDate` | `string` |  |
| `externalAccountIcon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `array` |  |
| `inserted` | `string` |  |
| `isArchived` | `bool` |  |
| `isConnectedAccount` | `bool` |  |
| `isDefault` | `bool` |  |
| `isTrustAccount` | `bool` |  |
| `isVirtual` | `bool` |  |
| `lastTransaction` | `array` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantName` | `string` |  |
| `name` | `string` |  |
| `physicalAccountID` | `string` |  |
| `rules` | `array` |  |
| `submittedPayoutsBalance` | `float` |  |
| `submittedPayoutsBalanceMinorUnits` | `int` |  |
| `summary` | `string` |  |
| `supplierSepaInstantStatus` | `string` |  |
| `xeroBankFeedConnectionStatus` | `string` |  |
| `xeroBankFeedLastSyncedAt` | `string` |  |
| `xeroBankFeedSyncLastFailedAt` | `string` |  |
| `xeroBankFeedSyncLastFailureReason` | `string` |  |
| `xeroBankFeedSyncStatus` | `string` |  |
| `xeroUnsynchronisedTransactionsCount` | `int` |  |

#### Example: Create

```php
$virtual = $client->Virtual()->create([
    "account_id" => null, // string
    "createdBy" => null, // array
    "identifier" => null, // array
    "name" => null, // string
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
| `destinationUrl` | `string` |  |
| `emailAddress` | `string` |  |
| `failedNotificationEmailAddress` | `string` |  |
| `id` | `string` |  |
| `isActive` | `bool` |  |
| `merchantID` | `string` |  |
| `notificationMethod` | `string` |  |
| `resourceTypes` | `array` |  |
| `retry` | `bool` |  |
| `secret` | `string` |  |
| `version` | `int` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Webhook record (throws on error).
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
$consent = $client->Consent();
$consent->list();

// $consent->data_get() now returns the consent data from the last list
// $consent->match_get() returns the last match criteria
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
