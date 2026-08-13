# Nofrixion PHP SDK Reference

Complete API reference for the Nofrixion PHP SDK.


## NofrixionSDK

### Constructor

```php
require_once __DIR__ . '/nofrixion_sdk.php';

$client = new NofrixionSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NofrixionSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = NofrixionSDK::test();
```


### Instance Methods

#### `Account($data = null)`

Create a new `AccountEntity` instance. Pass `null` for no initial data.

#### `Batch($data = null)`

Create a new `BatchEntity` instance. Pass `null` for no initial data.

#### `Beneficiary($data = null)`

Create a new `BeneficiaryEntity` instance. Pass `null` for no initial data.

#### `BeneficiaryGroup($data = null)`

Create a new `BeneficiaryGroupEntity` instance. Pass `null` for no initial data.

#### `Card($data = null)`

Create a new `CardEntity` instance. Pass `null` for no initial data.

#### `CardCustomerToken($data = null)`

Create a new `CardCustomerTokenEntity` instance. Pass `null` for no initial data.

#### `CardPayment($data = null)`

Create a new `CardPaymentEntity` instance. Pass `null` for no initial data.

#### `CardPublicKey($data = null)`

Create a new `CardPublicKeyEntity` instance. Pass `null` for no initial data.

#### `Consent($data = null)`

Create a new `ConsentEntity` instance. Pass `null` for no initial data.

#### `Currency($data = null)`

Create a new `CurrencyEntity` instance. Pass `null` for no initial data.

#### `DirectDebitBatchSubmit($data = null)`

Create a new `DirectDebitBatchSubmitEntity` instance. Pass `null` for no initial data.

#### `FxRate($data = null)`

Create a new `FxRateEntity` instance. Pass `null` for no initial data.

#### `IPayment($data = null)`

Create a new `IPaymentEntity` instance. Pass `null` for no initial data.

#### `Mandate($data = null)`

Create a new `MandateEntity` instance. Pass `null` for no initial data.

#### `Merchant($data = null)`

Create a new `MerchantEntity` instance. Pass `null` for no initial data.

#### `MerchantAuthorisationSetting($data = null)`

Create a new `MerchantAuthorisationSettingEntity` instance. Pass `null` for no initial data.

#### `MerchantDirectDebitMandatePage($data = null)`

Create a new `MerchantDirectDebitMandatePageEntity` instance. Pass `null` for no initial data.

#### `MerchantPayByBankSetting($data = null)`

Create a new `MerchantPayByBankSettingEntity` instance. Pass `null` for no initial data.

#### `MerchantPaymentRequestTemplate($data = null)`

Create a new `MerchantPaymentRequestTemplateEntity` instance. Pass `null` for no initial data.

#### `MerchantToken($data = null)`

Create a new `MerchantTokenEntity` instance. Pass `null` for no initial data.

#### `Metadata($data = null)`

Create a new `MetadataEntity` instance. Pass `null` for no initial data.

#### `NoFrixionVersion($data = null)`

Create a new `NoFrixionVersionEntity` instance. Pass `null` for no initial data.

#### `OpenBanking($data = null)`

Create a new `OpenBankingEntity` instance. Pass `null` for no initial data.

#### `Payeeverification($data = null)`

Create a new `PayeeverificationEntity` instance. Pass `null` for no initial data.

#### `Payment($data = null)`

Create a new `PaymentEntity` instance. Pass `null` for no initial data.

#### `PaymentAccount($data = null)`

Create a new `PaymentAccountEntity` instance. Pass `null` for no initial data.

#### `PaymentAccountMinimal($data = null)`

Create a new `PaymentAccountMinimalEntity` instance. Pass `null` for no initial data.

#### `PaymentInitiation($data = null)`

Create a new `PaymentInitiationEntity` instance. Pass `null` for no initial data.

#### `PaymentRequest($data = null)`

Create a new `PaymentRequestEntity` instance. Pass `null` for no initial data.

#### `PaymentRequestEvent($data = null)`

Create a new `PaymentRequestEventEntity` instance. Pass `null` for no initial data.

#### `PaymentRequestMetric($data = null)`

Create a new `PaymentRequestMetricEntity` instance. Pass `null` for no initial data.

#### `PaymentRequestMinimal($data = null)`

Create a new `PaymentRequestMinimalEntity` instance. Pass `null` for no initial data.

#### `PaymentRequestResult($data = null)`

Create a new `PaymentRequestResultEntity` instance. Pass `null` for no initial data.

#### `Payout($data = null)`

Create a new `PayoutEntity` instance. Pass `null` for no initial data.

#### `PayoutKeysetPage($data = null)`

Create a new `PayoutKeysetPageEntity` instance. Pass `null` for no initial data.

#### `PayoutMetric($data = null)`

Create a new `PayoutMetricEntity` instance. Pass `null` for no initial data.

#### `Payrun($data = null)`

Create a new `PayrunEntity` instance. Pass `null` for no initial data.

#### `Report($data = null)`

Create a new `ReportEntity` instance. Pass `null` for no initial data.

#### `ReportResult($data = null)`

Create a new `ReportResultEntity` instance. Pass `null` for no initial data.

#### `Role($data = null)`

Create a new `RoleEntity` instance. Pass `null` for no initial data.

#### `Rule($data = null)`

Create a new `RuleEntity` instance. Pass `null` for no initial data.

#### `RuleEvent($data = null)`

Create a new `RuleEventEntity` instance. Pass `null` for no initial data.

#### `Tag($data = null)`

Create a new `TagEntity` instance. Pass `null` for no initial data.

#### `Token($data = null)`

Create a new `TokenEntity` instance. Pass `null` for no initial data.

#### `Transaction($data = null)`

Create a new `TransactionEntity` instance. Pass `null` for no initial data.

#### `User($data = null)`

Create a new `UserEntity` instance. Pass `null` for no initial data.

#### `UserInvite($data = null)`

Create a new `UserInviteEntity` instance. Pass `null` for no initial data.

#### `Virtual($data = null)`

Create a new `VirtualEntity` instance. Pass `null` for no initial data.

#### `Webhook($data = null)`

Create a new `WebhookEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): NofrixionUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## AccountEntity

```php
$account = $client->Account();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountBalances` | `array` | No |  |
| `accountID` | `string` | No |  |
| `accountIdentifications` | `array` | No |  |
| `accountName` | `string` | No |  |
| `accountNames` | `array` | No |  |
| `accountSupplierName` | `string` | No |  |
| `accountType` | `string` | No |  |
| `availableBalance` | `float` | No |  |
| `availableBalanceMinorUnits` | `int` | No |  |
| `balance` | `float` | No |  |
| `balanceMinorUnits` | `int` | No |  |
| `bankName` | `string` | No |  |
| `consentID` | `string` | No |  |
| `consolidatedAccountInformation` | `array` | No |  |
| `createdBy` | `array` | Yes |  |
| `createdByDisplayName` | `string` | No |  |
| `currency` | `string` | No |  |
| `defaultPaymentRail` | `string` | No |  |
| `description` | `string` | No |  |
| `details` | `string` | No |  |
| `displayName` | `string` | No |  |
| `expiryDate` | `string` | No |  |
| `externalAccountIcon` | `string` | No |  |
| `format` | `string` | No |  |
| `fromDate` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `array` | Yes |  |
| `inserted` | `string` | No |  |
| `isArchived` | `bool` | No |  |
| `isConnectedAccount` | `bool` | No |  |
| `isDefault` | `bool` | No |  |
| `isTrustAccount` | `bool` | No |  |
| `isVirtual` | `bool` | No |  |
| `lastTransaction` | `array` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `nickname` | `string` | No |  |
| `physicalAccountID` | `string` | No |  |
| `roleIDs` | `array` | No |  |
| `rules` | `array` | No |  |
| `submittedPayoutsBalance` | `float` | No |  |
| `submittedPayoutsBalanceMinorUnits` | `int` | No |  |
| `summary` | `string` | No |  |
| `supplierPhysicalAccountID` | `string` | No |  |
| `supplierSepaInstantStatus` | `string` | No |  |
| `toDate` | `string` | No |  |
| `type` | `string` | No |  |
| `usageType` | `string` | No |  |
| `xeroBankFeedConnectionStatus` | `string` | No |  |
| `xeroBankFeedLastSyncedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `string` | No |  |
| `xeroBankFeedSyncStatus` | `string` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Account()->create([
  "createdBy" => null, // array
  "identifier" => null, // array
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Account()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Account()->load(["id" => "account_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Account()->remove(["id" => "account_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Account()->update([
  "id" => "account_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): AccountEntity`

Create a new `AccountEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## BatchEntity

```php
$batch = $client->Batch();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approveUrl` | `string` | No |  |
| `id` | `string` | No |  |
| `payouts` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Batch()->create([
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Batch()->load(["id" => "batch_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): BatchEntity`

Create a new `BatchEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## BeneficiaryEntity

```php
$beneficiary = $client->Beneficiary();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approvalCallbackUrl` | `string` | No |  |
| `authenticationMethods` | `array` | No |  |
| `authorisations` | `array` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `beneficiaries` | `array` | No |  |
| `beneficiaryEvents` | `array` | No |  |
| `canAuthorise` | `bool` | No |  |
| `canUpdate` | `bool` | No |  |
| `createdBy` | `array` | Yes |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `array` | No |  |
| `failedBeneficiaries` | `array` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isEnabled` | `bool` | No |  |
| `lastAuthorised` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `sourceAccountIDs` | `array` | No |  |
| `sourceAccounts` | `array` | No |  |
| `theirReference` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `approvalCallbackUrl` | - | - | - | - | - |
| `authenticationMethods` | - | - | - | - | - |
| `authorisations` | - | - | - | - | - |
| `authorisersCompletedCount` | - | - | - | - | - |
| `authorisersRequiredCount` | - | - | - | - | - |
| `beneficiaries` | - | - | - | - | - |
| `beneficiaryEvents` | - | - | - | - | - |
| `canAuthorise` | - | - | - | - | - |
| `canUpdate` | - | - | - | - | - |
| `createdBy` | - | - | - | - | - |
| `createdByEmailAddress` | - | - | - | - | - |
| `currency` | - | - | - | Yes | - |
| `destination` | - | - | Yes | - | - |
| `failedBeneficiaries` | - | - | - | - | - |
| `hasCurrentUserAuthorised` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `inserted` | - | - | - | - | - |
| `isEnabled` | - | - | - | - | - |
| `lastAuthorised` | - | - | - | - | - |
| `lastUpdated` | - | - | - | - | - |
| `merchantID` | - | - | Yes | - | - |
| `name` | - | - | - | Yes | - |
| `nonce` | - | - | - | - | - |
| `sourceAccountIDs` | - | - | - | - | - |
| `sourceAccounts` | - | - | - | - | - |
| `theirReference` | - | - | - | - | - |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Beneficiary()->create([
  "createdBy" => null, // array
  "currency" => null, // string
  "name" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Beneficiary()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Beneficiary()->load(["id" => "beneficiary_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Beneficiary()->remove(["id" => "beneficiary_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Beneficiary()->update([
  "id" => "beneficiary_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): BeneficiaryEntity`

Create a new `BeneficiaryEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## BeneficiaryGroupEntity

```php
$beneficiary_group = $client->BeneficiaryGroup();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `groupMembers` | `array` | No |  |
| `groupName` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->BeneficiaryGroup()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): BeneficiaryGroupEntity`

Create a new `BeneficiaryGroupEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CardEntity

```php
$card = $client->Card();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorizedAmount` | `string` | No |  |
| `currencyCode` | `string` | No |  |
| `isPayerAuthenticationRequired` | `bool` | No |  |
| `isSoftDecline` | `bool` | No |  |
| `payerAuthenticationAccessToken` | `string` | No |  |
| `payerAuthenticationMerchantData` | `string` | No |  |
| `payerAuthenticationUrl` | `string` | No |  |
| `payerAuthenticationWindowHeight` | `int` | No |  |
| `payerAuthenticationWindowWidth` | `int` | No |  |
| `paymentRequestCallbackUrl` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |
| `requestID` | `string` | No |  |
| `responseCode` | `string` | No |  |
| `responseType` | `string` | No |  |
| `status` | `string` | No |  |
| `threeDSRedirectUrl` | `string` | No |  |
| `transactionID` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Card()->create([
  "paymentrequest_id" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CardEntity`

Create a new `CardEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CardCustomerTokenEntity

```php
$card_customer_token = $client->CardCustomerToken();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cardType` | `string` | No |  |
| `customerEmailAddress` | `string` | No |  |
| `expiryMonth` | `string` | No |  |
| `expiryYear` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lastFourDigits` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `maskedCardNumber` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->CardCustomerToken()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->CardCustomerToken()->load(["customer_email_address" => "customer_email_address"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->CardCustomerToken()->remove();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CardCustomerTokenEntity`

Create a new `CardCustomerTokenEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CardPaymentEntity

```php
$card_payment = $client->CardPayment();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorizedAmount` | `string` | No |  |
| `currencyCode` | `string` | No |  |
| `isPayerAuthenticationRequired` | `bool` | No |  |
| `isSoftDecline` | `bool` | No |  |
| `payerAuthenticationAccessToken` | `string` | No |  |
| `payerAuthenticationMerchantData` | `string` | No |  |
| `payerAuthenticationUrl` | `string` | No |  |
| `payerAuthenticationWindowHeight` | `int` | No |  |
| `payerAuthenticationWindowWidth` | `int` | No |  |
| `paymentRequestCallbackUrl` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |
| `requestID` | `string` | No |  |
| `responseCode` | `string` | No |  |
| `responseType` | `string` | No |  |
| `status` | `string` | No |  |
| `threeDSRedirectUrl` | `string` | No |  |
| `transactionID` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->CardPayment()->create([
  "paymentrequest_id" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CardPaymentEntity`

Create a new `CardPaymentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CardPublicKeyEntity

```php
$card_public_key = $client->CardPublicKey();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jwt` | `string` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->CardPublicKey()->load(["paymentrequest_id" => "paymentrequest_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CardPublicKeyEntity`

Create a new `CardPublicKeyEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ConsentEntity

```php
$consent = $client->Consent();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisationUrl` | `string` | No |  |
| `callbackUrl` | `string` | No |  |
| `consentID` | `string` | No |  |
| `emailAddress` | `string` | No |  |
| `expiryDate` | `string` | No |  |
| `failureCallbackUrl` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `institutionID` | `string` | No |  |
| `isConnectedAccounts` | `bool` | No |  |
| `isEnabled` | `bool` | No |  |
| `merchantID` | `string` | No |  |
| `provider` | `string` | No |  |
| `successWebHookUrl` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `authorisationUrl` | - | - | - | - | - |
| `callbackUrl` | - | - | - | - | - |
| `consentID` | - | - | - | - | - |
| `emailAddress` | - | - | - | - | - |
| `expiryDate` | - | - | - | - | - |
| `failureCallbackUrl` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `inserted` | - | - | - | - | - |
| `institutionID` | - | - | Yes | - | - |
| `isConnectedAccounts` | - | - | - | - | - |
| `isEnabled` | - | - | - | - | - |
| `merchantID` | - | - | Yes | - | - |
| `provider` | - | - | - | - | - |
| `successWebHookUrl` | - | - | - | - | - |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Consent()->create([
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Consent()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Consent()->load(["id" => "consent_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Consent()->remove(["id" => "consent_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Consent()->update([
  "id" => "consent_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ConsentEntity`

Create a new `ConsentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CurrencyEntity

```php
$currency = $client->Currency();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `code` | `string` | No |  |
| `decimals` | `int` | No |  |
| `isFiat` | `bool` | No |  |
| `iso4217AlphaCode` | `string` | No |  |
| `iso4217NumericCode` | `string` | No |  |
| `symbol` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Currency()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CurrencyEntity`

Create a new `CurrencyEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## DirectDebitBatchSubmitEntity

```php
$direct_debit_batch_submit = $client->DirectDebitBatchSubmit();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failedSubmissions` | `array` | No |  |
| `successfulSubmissions` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->DirectDebitBatchSubmit()->create([
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): DirectDebitBatchSubmitEntity`

Create a new `DirectDebitBatchSubmitEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## FxRateEntity

```php
$fx_rate = $client->FxRate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destinationCurrency` | `string` | No |  |
| `exchangeRate` | `float` | No |  |
| `expiryTime` | `string` | No |  |
| `quoteID` | `string` | No |  |
| `sourceCurrency` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->FxRate()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->FxRate()->load(["destination" => "destination", "source" => "source", "valid_for_minute" => 1]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): FxRateEntity`

Create a new `FxRateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## IPaymentEntity

```php
$i_payment = $client->IPayment();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `paymentRequestID` | `string` | No |  |
| `responseType` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->IPayment()->create([
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): IPaymentEntity`

Create a new `IPaymentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MandateEntity

```php
$mandate = $client->Mandate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountNumber` | `string` | No |  |
| `addressLine1` | `string` | Yes |  |
| `addressLine2` | `string` | No |  |
| `approvedAt` | `string` | No |  |
| `city` | `string` | Yes |  |
| `countryCode` | `string` | Yes |  |
| `currency` | `string` | No |  |
| `customerAccountNumber` | `string` | No |  |
| `customerCity` | `string` | No |  |
| `customerCountryCode` | `string` | No |  |
| `customerCountryName` | `string` | No |  |
| `customerEmailAddress` | `string` | No |  |
| `customerFirstName` | `string` | No |  |
| `customerIban` | `string` | No |  |
| `customerLastName` | `string` | No |  |
| `customerSortCode` | `string` | No |  |
| `emailAddress` | `string` | Yes |  |
| `firstName` | `string` | Yes |  |
| `iban` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isRecurring` | `bool` | No |  |
| `lastName` | `string` | Yes |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `postalCode` | `string` | Yes |  |
| `reference` | `string` | No |  |
| `sortCode` | `string` | No |  |
| `status` | `string` | No |  |
| `supplierBankAccountID` | `string` | No |  |
| `supplierCustomerID` | `string` | No |  |
| `supplierMandateID` | `string` | No |  |
| `supplierName` | `string` | No |  |
| `supplierStatus` | `string` | No |  |

### Field Usage by Operation

| Field | load | create |
| --- | --- | --- |
| `accountNumber` | - | - |
| `addressLine1` | - | - |
| `addressLine2` | - | - |
| `approvedAt` | - | - |
| `city` | - | - |
| `countryCode` | - | - |
| `currency` | - | Yes |
| `customerAccountNumber` | - | - |
| `customerCity` | - | - |
| `customerCountryCode` | - | - |
| `customerCountryName` | - | - |
| `customerEmailAddress` | - | - |
| `customerFirstName` | - | - |
| `customerIban` | - | - |
| `customerLastName` | - | - |
| `customerSortCode` | - | - |
| `emailAddress` | - | - |
| `firstName` | - | - |
| `iban` | - | - |
| `id` | - | - |
| `inserted` | - | - |
| `isRecurring` | - | - |
| `lastName` | - | - |
| `lastUpdated` | - | - |
| `merchantID` | - | Yes |
| `postalCode` | - | - |
| `reference` | - | - |
| `sortCode` | - | - |
| `status` | - | - |
| `supplierBankAccountID` | - | - |
| `supplierCustomerID` | - | - |
| `supplierMandateID` | - | - |
| `supplierName` | - | - |
| `supplierStatus` | - | - |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Mandate()->create([
  "addressLine1" => null, // string
  "city" => null, // string
  "countryCode" => null, // string
  "emailAddress" => null, // string
  "firstName" => null, // string
  "lastName" => null, // string
  "postalCode" => null, // string
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Mandate()->load(["id" => "mandate_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MandateEntity`

Create a new `MandateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MerchantEntity

```php
$merchant = $client->Merchant();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountCurrencies` | `array` | No |  |
| `canHaveTrustAccounts` | `bool` | No |  |
| `cardPaymentProcessor` | `string` | No |  |
| `companyID` | `string` | No |  |
| `displayQrOnHostedPay` | `bool` | No |  |
| `hostedPayVersion` | `int` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isBlocked` | `bool` | No |  |
| `isExited` | `bool` | No |  |
| `isSuspended` | `bool` | No |  |
| `jurisdiction` | `string` | No |  |
| `logoUrlPng` | `string` | No |  |
| `logoUrlSvg` | `string` | No |  |
| `merchantCategoryCode` | `string` | No |  |
| `name` | `string` | No |  |
| `notes` | `string` | No |  |
| `parentMerchant` | `array` | No |  |
| `paymentAccountLimit` | `int` | No |  |
| `paymentAccounts` | `array` | No |  |
| `reason` | `string` | No |  |
| `shortName` | `string` | No |  |
| `supportedPaymentMethodsList` | `array` | No |  |
| `suspensionReason` | `string` | No |  |
| `tags` | `array` | No |  |
| `timeZoneId` | `string` | No |  |
| `tradingName` | `string` | No |  |
| `webHookLimit` | `int` | No |  |
| `yourRoleName` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Merchant()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Merchant()->load(["id" => "merchant_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Merchant()->remove(["id" => "merchant_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Merchant()->update([
  "id" => "merchant_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MerchantEntity`

Create a new `MerchantEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MerchantAuthorisationSettingEntity

```php
$merchant_authorisation_setting = $client->MerchantAuthorisationSetting();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amountLower` | `float` | No |  |
| `amountUpper` | `float` | No |  |
| `authorisationType` | `string` | No |  |
| `beneficiariesOnly` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lastEditorCantAuthorise` | `bool` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `numberOfAuthorisers` | `int` | No |  |
| `roleSettings` | `array` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->MerchantAuthorisationSetting()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MerchantAuthorisationSettingEntity`

Create a new `MerchantAuthorisationSettingEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MerchantDirectDebitMandatePageEntity

```php
$merchant_direct_debit_mandate_page = $client->MerchantDirectDebitMandatePage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approvedAt` | `string` | No |  |
| `currency` | `string` | No |  |
| `customerAccountNumber` | `string` | No |  |
| `customerCity` | `string` | No |  |
| `customerCountryCode` | `string` | No |  |
| `customerCountryName` | `string` | No |  |
| `customerEmailAddress` | `string` | No |  |
| `customerFirstName` | `string` | No |  |
| `customerIban` | `string` | No |  |
| `customerLastName` | `string` | No |  |
| `customerSortCode` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isRecurring` | `bool` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `reference` | `string` | No |  |
| `status` | `string` | No |  |
| `supplierBankAccountID` | `string` | No |  |
| `supplierCustomerID` | `string` | No |  |
| `supplierMandateID` | `string` | No |  |
| `supplierName` | `string` | No |  |
| `supplierStatus` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->MerchantDirectDebitMandatePage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MerchantDirectDebitMandatePageEntity`

Create a new `MerchantDirectDebitMandatePageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MerchantPayByBankSettingEntity

```php
$merchant_pay_by_bank_setting = $client->MerchantPayByBankSetting();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bankCountryCodes` | `array` | No |  |
| `bankID` | `string` | No |  |
| `bankName` | `string` | No |  |
| `businessInstitutionID` | `string` | No |  |
| `currency` | `string` | No |  |
| `logo` | `string` | No |  |
| `message` | `string` | No |  |
| `messageImageUrl` | `string` | No |  |
| `order` | `int` | No |  |
| `personalInstitutionID` | `string` | No |  |
| `processor` | `string` | No |  |
| `warningHeading` | `string` | No |  |
| `warningMessage` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->MerchantPayByBankSetting()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MerchantPayByBankSettingEntity`

Create a new `MerchantPayByBankSettingEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MerchantPaymentRequestTemplateEntity

```php
$merchant_payment_request_template = $client->MerchantPaymentRequestTemplate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bankPaymentOptions` | `array` | No |  |
| `cardPaymentAddressOptions` | `array` | No |  |
| `cardPaymentCaptureOptions` | `array` | No |  |
| `customFields` | `array` | No |  |
| `defaultFields` | `array` | No |  |
| `description` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `name` | `string` | Yes |  |
| `notificationOptions` | `array` | No |  |
| `paymentMethods` | `array` | No |  |
| `paymentTerms` | `array` | No |  |
| `priorityBankOptions` | `array` | No |  |
| `template` | `array` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->MerchantPaymentRequestTemplate()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->MerchantPaymentRequestTemplate()->load(["id" => "merchant_payment_request_template_id", "paymentrequest_id" => "paymentrequest_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->MerchantPaymentRequestTemplate()->remove(["id" => "merchant_payment_request_template_id", "paymentrequest_id" => "paymentrequest_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->MerchantPaymentRequestTemplate()->update([
  "id" => "merchant_payment_request_template_id",
  "paymentrequest_id" => "paymentrequest_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MerchantPaymentRequestTemplateEntity`

Create a new `MerchantPaymentRequestTemplateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MerchantTokenEntity

```php
$merchant_token = $client->MerchantToken();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authenticationMethods` | `array` | No |  |
| `authorisations` | `array` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `canAuthorise` | `bool` | No |  |
| `description` | `string` | No |  |
| `expiresAt` | `string` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No |  |
| `hmacAlgorithm` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `ipAddressWhitelist` | `string` | No |  |
| `isArchived` | `bool` | No |  |
| `isEnabled` | `bool` | No |  |
| `lastAuthorised` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `nonce` | `string` | Yes |  |
| `permissionTypes` | `array` | No |  |
| `requestSignatureVersion` | `int` | No |  |
| `sharedSecretAlgorithm` | `string` | No |  |
| `sharedSecretBase64` | `string` | No |  |
| `token` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update |
| --- | --- | --- | --- | --- |
| `authenticationMethods` | - | - | - | - |
| `authorisations` | - | - | - | - |
| `authorisersCompletedCount` | - | - | - | - |
| `authorisersRequiredCount` | - | - | - | - |
| `canAuthorise` | - | - | - | - |
| `description` | - | - | Yes | - |
| `expiresAt` | - | - | - | - |
| `hasCurrentUserAuthorised` | - | - | - | - |
| `hmacAlgorithm` | - | - | - | - |
| `id` | - | - | - | - |
| `inserted` | - | - | - | - |
| `ipAddressWhitelist` | - | - | - | - |
| `isArchived` | - | - | - | - |
| `isEnabled` | - | - | - | - |
| `lastAuthorised` | - | - | - | - |
| `lastUpdated` | - | - | - | - |
| `merchantID` | - | - | Yes | - |
| `nonce` | - | - | - | - |
| `permissionTypes` | - | - | - | - |
| `requestSignatureVersion` | - | - | - | - |
| `sharedSecretAlgorithm` | - | - | - | - |
| `sharedSecretBase64` | - | - | - | - |
| `token` | - | - | - | - |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->MerchantToken()->create([
  "nonce" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->MerchantToken()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->MerchantToken()->load(["id" => "merchant_token_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->MerchantToken()->update([
  "id" => "merchant_token_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MerchantTokenEntity`

Create a new `MerchantTokenEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MetadataEntity

```php
$metadata = $client->Metadata();
```

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Metadata()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MetadataEntity`

Create a new `MetadataEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionVersionEntity

```php
$no_frixion_version = $client->NoFrixionVersion();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `buildVersion` | `int` | No |  |
| `majorVersion` | `int` | No |  |
| `minorVersion` | `int` | No |  |
| `releaseName` | `string` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionVersion()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionVersionEntity`

Create a new `NoFrixionVersionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## OpenBankingEntity

```php
$open_banking = $client->OpenBanking();
```

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->OpenBanking()->create([
  "account_id" => null, // string
]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->OpenBanking()->remove();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): OpenBankingEntity`

Create a new `OpenBankingEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PayeeverificationEntity

```php
$payeeverification = $client->Payeeverification();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `string` | Yes |  |
| `accountNumber` | `string` | No |  |
| `iban` | `string` | Yes |  |
| `payeeVerifiedAccountName` | `string` | No |  |
| `result` | `string` | No |  |
| `secondaryIdentification` | `string` | No |  |
| `sortCode` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Payeeverification()->create([
  "accountName" => null, // string
  "iban" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PayeeverificationEntity`

Create a new `PayeeverificationEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PaymentEntity

```php
$payment = $client->Payment();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresses` | `array` | No |  |
| `amount` | `float` | No |  |
| `amountPending` | `float` | No |  |
| `amountReceived` | `float` | No |  |
| `amountRefunded` | `float` | No |  |
| `autoSendReceipt` | `bool` | No |  |
| `baseOriginUrl` | `string` | No |  |
| `callbackUrl` | `string` | No |  |
| `cardAuthorizeOnly` | `bool` | No |  |
| `cardCreateToken` | `bool` | No |  |
| `cardCreateTokenMode` | `string` | No |  |
| `cardIgnoreCVN` | `bool` | No |  |
| `cardNoPayerAuthentication` | `bool` | No |  |
| `cardProcessorMerchantID` | `string` | No |  |
| `cardStripePaymentIntentID` | `string` | No |  |
| `cardStripePaymentIntentSecret` | `string` | No |  |
| `cardTransmitRawDetails` | `bool` | No |  |
| `createdByUser` | `array` | Yes |  |
| `currency` | `string` | No |  |
| `customFields` | `array` | No |  |
| `customerEmailAddress` | `string` | No |  |
| `customerID` | `string` | No |  |
| `customerName` | `string` | No |  |
| `description` | `string` | No |  |
| `destinationAccount` | `array` | No |  |
| `directDebitPayment` | `array` | No |  |
| `dueDate` | `string` | No |  |
| `events` | `array` | No |  |
| `failureCallbackUrl` | `string` | No |  |
| `fieldDisplaySettings` | `array` | No |  |
| `formattedAmount` | `string` | No |  |
| `hostedPayCheckoutUrl` | `string` | No |  |
| `id` | `string` | No |  |
| `ignoreAddressVerification` | `bool` | No |  |
| `inserted` | `string` | No |  |
| `insertedSortable` | `string` | No |  |
| `isArchived` | `bool` | No |  |
| `jwk` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `lightningInvoice` | `string` | No |  |
| `lightningInvoiceExpiresAt` | `string` | No |  |
| `merchantDirectDebitMandateID` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantTokenDescription` | `string` | No |  |
| `notificationEmailAddresses` | `string` | No |  |
| `notificationRoleIDs` | `array` | No |  |
| `orderID` | `string` | No |  |
| `partialPaymentMethod` | `string` | No |  |
| `partialPaymentSteps` | `string` | No |  |
| `paymentAttempts` | `array` | No |  |
| `paymentMethods` | `array` | No |  |
| `paymentProcessor` | `string` | No |  |
| `payrunID` | `string` | No |  |
| `pispAccountID` | `string` | No |  |
| `priorityBankID` | `string` | No |  |
| `result` | `array` | No |  |
| `sandboxSettleDelayInSeconds` | `int` | No |  |
| `shippingAddress` | `array` | No |  |
| `shippingAddressCity` | `string` | No |  |
| `shippingAddressCountryCode` | `string` | No |  |
| `shippingAddressCounty` | `string` | No |  |
| `shippingAddressLine1` | `string` | No |  |
| `shippingAddressLine2` | `string` | No |  |
| `shippingAddressPostCode` | `string` | No |  |
| `shippingEmail` | `string` | No |  |
| `shippingFirstName` | `string` | No |  |
| `shippingLastName` | `string` | No |  |
| `shippingPhone` | `string` | No |  |
| `status` | `string` | No |  |
| `successWebHookUrl` | `string` | No |  |
| `tagIds` | `array` | No |  |
| `tags` | `array` | No |  |
| `title` | `string` | No |  |
| `tokenisedCards` | `array` | No |  |
| `transactions` | `array` | No |  |
| `useHostedPaymentPage` | `bool` | No |  |

### Field Usage by Operation

| Field | load | create | update |
| --- | --- | --- | --- |
| `addresses` | - | - | - |
| `amount` | - | Yes | - |
| `amountPending` | - | - | - |
| `amountReceived` | - | - | - |
| `amountRefunded` | - | - | - |
| `autoSendReceipt` | - | - | - |
| `baseOriginUrl` | - | - | - |
| `callbackUrl` | - | - | - |
| `cardAuthorizeOnly` | - | - | - |
| `cardCreateToken` | - | - | - |
| `cardCreateTokenMode` | - | - | - |
| `cardIgnoreCVN` | - | - | - |
| `cardNoPayerAuthentication` | - | - | - |
| `cardProcessorMerchantID` | - | - | - |
| `cardStripePaymentIntentID` | - | - | - |
| `cardStripePaymentIntentSecret` | - | - | - |
| `cardTransmitRawDetails` | - | - | - |
| `createdByUser` | - | - | - |
| `currency` | - | - | - |
| `customFields` | - | - | - |
| `customerEmailAddress` | - | - | - |
| `customerID` | - | - | - |
| `customerName` | - | - | - |
| `description` | - | - | - |
| `destinationAccount` | - | - | - |
| `directDebitPayment` | - | - | - |
| `dueDate` | - | - | - |
| `events` | - | - | - |
| `failureCallbackUrl` | - | - | - |
| `fieldDisplaySettings` | - | - | - |
| `formattedAmount` | - | - | - |
| `hostedPayCheckoutUrl` | - | - | - |
| `id` | - | - | - |
| `ignoreAddressVerification` | - | - | - |
| `inserted` | - | - | - |
| `insertedSortable` | - | - | - |
| `isArchived` | - | - | - |
| `jwk` | - | - | - |
| `lastUpdated` | - | - | - |
| `lightningInvoice` | - | - | - |
| `lightningInvoiceExpiresAt` | - | - | - |
| `merchantDirectDebitMandateID` | - | - | - |
| `merchantID` | - | - | - |
| `merchantTokenDescription` | - | - | - |
| `notificationEmailAddresses` | - | - | - |
| `notificationRoleIDs` | - | - | - |
| `orderID` | - | - | - |
| `partialPaymentMethod` | - | - | - |
| `partialPaymentSteps` | - | - | - |
| `paymentAttempts` | - | - | - |
| `paymentMethods` | - | - | - |
| `paymentProcessor` | - | - | - |
| `payrunID` | - | - | - |
| `pispAccountID` | - | - | - |
| `priorityBankID` | - | - | - |
| `result` | - | - | - |
| `sandboxSettleDelayInSeconds` | - | - | - |
| `shippingAddress` | - | - | - |
| `shippingAddressCity` | - | - | - |
| `shippingAddressCountryCode` | - | - | - |
| `shippingAddressCounty` | - | - | - |
| `shippingAddressLine1` | - | - | - |
| `shippingAddressLine2` | - | - | - |
| `shippingAddressPostCode` | - | - | - |
| `shippingEmail` | - | - | - |
| `shippingFirstName` | - | - | - |
| `shippingLastName` | - | - | - |
| `shippingPhone` | - | - | - |
| `status` | - | - | - |
| `successWebHookUrl` | - | - | - |
| `tagIds` | - | - | - |
| `tags` | - | - | - |
| `title` | - | - | - |
| `tokenisedCards` | - | - | - |
| `transactions` | - | - | - |
| `useHostedPaymentPage` | - | - | - |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Payment()->create([
  "createdByUser" => null, // array
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Payment()->load(["id" => "payment_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Payment()->update([
  "id" => "payment_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PaymentEntity`

Create a new `PaymentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PaymentAccountEntity

```php
$payment_account = $client->PaymentAccount();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `string` | No |  |
| `accountSupplierName` | `string` | No |  |
| `availableBalance` | `float` | No |  |
| `availableBalanceMinorUnits` | `int` | No |  |
| `balance` | `float` | No |  |
| `balanceMinorUnits` | `int` | No |  |
| `bankName` | `string` | No |  |
| `consentID` | `string` | No |  |
| `createdBy` | `array` | Yes |  |
| `createdByDisplayName` | `string` | No |  |
| `currency` | `string` | No |  |
| `defaultPaymentRail` | `string` | No |  |
| `displayName` | `string` | No |  |
| `expiryDate` | `string` | No |  |
| `externalAccountIcon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `array` | Yes |  |
| `inserted` | `string` | No |  |
| `isArchived` | `bool` | No |  |
| `isConnectedAccount` | `bool` | No |  |
| `isDefault` | `bool` | No |  |
| `isTrustAccount` | `bool` | No |  |
| `isVirtual` | `bool` | No |  |
| `lastTransaction` | `array` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `physicalAccountID` | `string` | No |  |
| `rules` | `array` | No |  |
| `submittedPayoutsBalance` | `float` | No |  |
| `submittedPayoutsBalanceMinorUnits` | `int` | No |  |
| `summary` | `string` | No |  |
| `supplierSepaInstantStatus` | `string` | No |  |
| `xeroBankFeedConnectionStatus` | `string` | No |  |
| `xeroBankFeedLastSyncedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `string` | No |  |
| `xeroBankFeedSyncStatus` | `string` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->PaymentAccount()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PaymentAccountEntity`

Create a new `PaymentAccountEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PaymentAccountMinimalEntity

```php
$payment_account_minimal = $client->PaymentAccountMinimal();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `string` | No |  |
| `availableBalance` | `float` | No |  |
| `balance` | `float` | No |  |
| `balanceMinorUnits` | `int` | No |  |
| `currency` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `array` | Yes |  |
| `isArchived` | `bool` | No |  |
| `isConnectedAccount` | `bool` | No |  |
| `merchantID` | `string` | No |  |
| `submittedPayoutsBalance` | `float` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->PaymentAccountMinimal()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PaymentAccountMinimalEntity`

Create a new `PaymentAccountMinimalEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PaymentInitiationEntity

```php
$payment_initiation = $client->PaymentInitiation();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `paymentInitiationID` | `string` | No |  |
| `paymentRequestCallbackUrl` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |
| `redirectUrl` | `string` | No |  |
| `responseType` | `string` | No |  |
| `specificErrorMessage` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PaymentInitiation()->create([
  "paymentrequest_id" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PaymentInitiationEntity`

Create a new `PaymentInitiationEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PaymentRequestEntity

```php
$payment_request = $client->PaymentRequest();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresses` | `array` | No |  |
| `amount` | `float` | No |  |
| `amountPending` | `float` | No |  |
| `amountReceived` | `float` | No |  |
| `amountRefunded` | `float` | No |  |
| `autoSendReceipt` | `bool` | No |  |
| `baseOriginUrl` | `string` | No |  |
| `callbackUrl` | `string` | No |  |
| `cardAuthorizeOnly` | `bool` | No |  |
| `cardCreateToken` | `bool` | No |  |
| `cardCreateTokenMode` | `string` | No |  |
| `cardIgnoreCVN` | `bool` | No |  |
| `cardProcessorMerchantID` | `string` | No |  |
| `cardStripePaymentIntentID` | `string` | No |  |
| `cardStripePaymentIntentSecret` | `string` | No |  |
| `createdByUser` | `array` | Yes |  |
| `currency` | `string` | No |  |
| `customFields` | `array` | No |  |
| `customerEmailAddress` | `string` | No |  |
| `customerID` | `string` | No |  |
| `customerName` | `string` | No |  |
| `description` | `string` | No |  |
| `destinationAccount` | `array` | No |  |
| `directDebitPayment` | `array` | No |  |
| `doSimulateSettlementFailure` | `bool` | No |  |
| `dueDate` | `string` | No |  |
| `errorDescription` | `string` | No |  |
| `events` | `array` | No |  |
| `failedPaymentRequests` | `array` | No |  |
| `failureCallbackUrl` | `string` | No |  |
| `fieldDisplaySettings` | `array` | No |  |
| `formattedAmount` | `string` | No |  |
| `hostedPayCheckoutUrl` | `string` | No |  |
| `id` | `string` | No |  |
| `ignoreAddressVerification` | `bool` | No |  |
| `inserted` | `string` | No |  |
| `insertedSortable` | `string` | No |  |
| `institution` | `string` | No |  |
| `isArchived` | `bool` | No |  |
| `jwk` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `lightningInvoice` | `string` | No |  |
| `lightningInvoiceExpiresAt` | `string` | No |  |
| `merchantDirectDebitMandateID` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantTokenDescription` | `string` | No |  |
| `notificationEmailAddresses` | `string` | No |  |
| `notificationRoleIDs` | `array` | No |  |
| `orderID` | `string` | No |  |
| `partialPaymentMethod` | `string` | No |  |
| `partialPaymentSteps` | `string` | No |  |
| `paymentAttempts` | `array` | No |  |
| `paymentInitiationID` | `string` | No |  |
| `paymentMethods` | `array` | No |  |
| `paymentProcessor` | `string` | No |  |
| `paymentRequests` | `array` | No |  |
| `payrunID` | `string` | No |  |
| `pispAccountID` | `string` | No |  |
| `priorityBankID` | `string` | No |  |
| `result` | `array` | No |  |
| `sandboxSettleDelayInSeconds` | `int` | No |  |
| `shippingAddress` | `array` | No |  |
| `status` | `string` | No |  |
| `successWebHookUrl` | `string` | No |  |
| `tags` | `array` | No |  |
| `title` | `string` | No |  |
| `tokenisedCards` | `array` | No |  |
| `transactions` | `array` | No |  |
| `useHostedPaymentPage` | `bool` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PaymentRequest()->create([
  "createdByUser" => null, // array
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->PaymentRequest()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->PaymentRequest()->load();
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->PaymentRequest()->remove(["id" => "id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->PaymentRequest()->update([
  "paymentrequest_id" => "paymentrequest_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PaymentRequestEntity`

Create a new `PaymentRequestEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PaymentRequestEventEntity

```php
$payment_request_event = $client->PaymentRequestEvent();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float` | Yes |  |
| `applePayTransactionID` | `string` | No |  |
| `cardAuthorizationResponseID` | `string` | No |  |
| `cardExpiryMonth` | `int` | No |  |
| `cardExpiryYear` | `int` | No |  |
| `cardIssuer` | `string` | No |  |
| `cardIssuerCountry` | `string` | No |  |
| `cardLastFourDigits` | `string` | No |  |
| `cardRequestID` | `string` | No |  |
| `cardScheme` | `string` | No |  |
| `cardTokenCustomerID` | `string` | No |  |
| `cardTransactionID` | `string` | No |  |
| `currency` | `string` | No |  |
| `directDebitPaymentID` | `string` | No |  |
| `directDebitPaymentReference` | `string` | No |  |
| `drirectDebitMandateID` | `string` | No |  |
| `errorMessage` | `string` | No |  |
| `errorReason` | `string` | No |  |
| `eventType` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lightningInvoice` | `string` | No |  |
| `lightningRHash` | `string` | No |  |
| `originUrl` | `string` | No |  |
| `paymentMethodType` | `string` | No |  |
| `paymentProcessorName` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |
| `pispBankStatus` | `string` | No |  |
| `pispPaymentInitiationID` | `string` | No |  |
| `pispPaymentInstitutionName` | `string` | No |  |
| `pispPaymentServiceProviderID` | `string` | No |  |
| `pispRedirectUrl` | `string` | No |  |
| `reconciledTransactionID` | `string` | No |  |
| `refundPayoutID` | `string` | No |  |
| `status` | `string` | No |  |
| `walletName` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->PaymentRequestEvent()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PaymentRequestEventEntity`

Create a new `PaymentRequestEventEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PaymentRequestMetricEntity

```php
$payment_request_metric = $client->PaymentRequestMetric();
```

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->PaymentRequestMetric()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PaymentRequestMetricEntity`

Create a new `PaymentRequestMetricEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PaymentRequestMinimalEntity

```php
$payment_request_minimal = $client->PaymentRequestMinimal();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float` | No |  |
| `amountPending` | `float` | No |  |
| `amountReceived` | `float` | No |  |
| `amountRefunded` | `float` | No |  |
| `callbackUrl` | `string` | No |  |
| `cardStripePaymentIntentSecret` | `string` | No |  |
| `countryCode` | `string` | No |  |
| `currency` | `string` | No |  |
| `customFieldsToDisplay` | `array` | No |  |
| `description` | `string` | No |  |
| `dueDate` | `string` | No |  |
| `fieldDisplaySettings` | `array` | No |  |
| `googlePayMerchantID` | `string` | No |  |
| `id` | `string` | No |  |
| `jwk` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantLogoUrlPng` | `string` | No |  |
| `merchantLogoUrlSvg` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `merchantShortName` | `string` | No |  |
| `partialPaymentMethod` | `string` | No |  |
| `paymentAttempts` | `array` | No |  |
| `paymentMethodsList` | `array` | No |  |
| `paymentProcessor` | `string` | No |  |
| `paymentProcessorKey` | `string` | No |  |
| `pispError` | `string` | No |  |
| `priorityBankID` | `string` | No |  |
| `status` | `string` | No |  |
| `stripeAccountID` | `string` | No |  |
| `title` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->PaymentRequestMinimal()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PaymentRequestMinimalEntity`

Create a new `PaymentRequestMinimalEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PaymentRequestResultEntity

```php
$payment_request_result = $client->PaymentRequestResult();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float` | No |  |
| `amountPending` | `float` | No |  |
| `amountReceived` | `float` | No |  |
| `amountRefunded` | `float` | No |  |
| `currency` | `string` | No |  |
| `customerID` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |
| `payments` | `array` | No |  |
| `pispAuthorizations` | `array` | No |  |
| `requestedAmount` | `float` | No |  |
| `result` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->PaymentRequestResult()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PaymentRequestResultEntity`

Create a new `PaymentRequestResultEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PayoutEntity

```php
$payout = $client->Payout();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountID` | `string` | No |  |
| `allowIncomplete` | `bool` | No |  |
| `amount` | `float` | No |  |
| `amountMinorUnits` | `int` | No |  |
| `approvePayoutUrl` | `string` | No |  |
| `approverID` | `string` | No |  |
| `authenticationMethods` | `array` | No |  |
| `authorisations` | `array` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `batchPayoutID` | `string` | No |  |
| `beneficiary` | `array` | Yes |  |
| `beneficiaryID` | `string` | No |  |
| `canAuthorise` | `bool` | No |  |
| `canProcess` | `bool` | No |  |
| `canUpdate` | `bool` | No |  |
| `chargeBearer` | `string` | No |  |
| `createdBy` | `string` | No |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | No |  |
| `currentUserID` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `array` | No |  |
| `documents` | `array` | No |  |
| `events` | `array` | No |  |
| `failedPayouts` | `array` | No |  |
| `formattedAmount` | `string` | No |  |
| `formattedFxDestinationAmount` | `string` | No |  |
| `formattedSchedule` | `string` | No |  |
| `formattedScheduleDayOnly` | `string` | No |  |
| `formattedSourceAccountAvailableBalance` | `string` | No |  |
| `fxDestinationAmount` | `float` | No |  |
| `fxDestinationAmountMinorUnits` | `int` | No |  |
| `fxDestinationCurrency` | `string` | No |  |
| `fxQuoteExpiresAt` | `string` | No |  |
| `fxQuoteID` | `string` | No |  |
| `fxRate` | `float` | No |  |
| `fxUseDestinationAmount` | `bool` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoiceID` | `string` | No |  |
| `isArchived` | `bool` | No |  |
| `isFailed` | `bool` | No |  |
| `isSettled` | `bool` | No |  |
| `isSubmitted` | `bool` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantTokenDescription` | `string` | No |  |
| `nonce` | `string` | No |  |
| `paymentProcessor` | `string` | No |  |
| `paymentRail` | `string` | No |  |
| `payouts` | `array` | No |  |
| `payrunID` | `string` | No |  |
| `payrunName` | `string` | No |  |
| `reason` | `string` | No |  |
| `rule` | `array` | No |  |
| `scheduleDate` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `sourceAccountAvailableBalance` | `float` | No |  |
| `sourceAccountAvailableBalanceMinorUnits` | `int` | No |  |
| `sourceAccountBic` | `string` | No |  |
| `sourceAccountCurrency` | `string` | No |  |
| `sourceAccountIban` | `string` | No |  |
| `sourceAccountIdentifier` | `array` | Yes |  |
| `sourceAccountName` | `string` | No |  |
| `sourceAccountNumber` | `string` | No |  |
| `sourceAccountSortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tagIds` | `array` | No |  |
| `tags` | `array` | No |  |
| `theirReference` | `string` | No |  |
| `topupPayrunID` | `string` | No |  |
| `transactedAmount` | `float` | No |  |
| `transactedFxAmount` | `float` | No |  |
| `transactedFxRate` | `float` | No |  |
| `type` | `string` | No |  |
| `userID` | `string` | No |  |
| `yourReference` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `accountID` | - | - | Yes | - | - |
| `allowIncomplete` | - | - | - | - | - |
| `amount` | - | - | - | - | - |
| `amountMinorUnits` | - | - | - | - | - |
| `approvePayoutUrl` | - | - | - | - | - |
| `approverID` | - | - | - | - | - |
| `authenticationMethods` | - | - | - | - | - |
| `authorisations` | - | - | - | - | - |
| `authorisersCompletedCount` | - | - | - | - | - |
| `authorisersRequiredCount` | - | - | - | - | - |
| `batchPayoutID` | - | - | - | - | - |
| `beneficiary` | - | - | - | - | - |
| `beneficiaryID` | - | - | - | - | - |
| `canAuthorise` | - | - | - | - | - |
| `canProcess` | - | - | - | - | - |
| `canUpdate` | - | - | - | - | - |
| `chargeBearer` | - | - | - | - | - |
| `createdBy` | - | - | - | - | - |
| `createdByEmailAddress` | - | - | - | - | - |
| `currency` | - | - | Yes | - | - |
| `currentUserID` | - | - | - | - | - |
| `description` | - | - | - | - | - |
| `destination` | - | - | - | - | - |
| `documents` | - | - | - | - | - |
| `events` | - | - | - | - | - |
| `failedPayouts` | - | - | - | - | - |
| `formattedAmount` | - | - | - | - | - |
| `formattedFxDestinationAmount` | - | - | - | - | - |
| `formattedSchedule` | - | - | - | - | - |
| `formattedScheduleDayOnly` | - | - | - | - | - |
| `formattedSourceAccountAvailableBalance` | - | - | - | - | - |
| `fxDestinationAmount` | - | - | - | - | - |
| `fxDestinationAmountMinorUnits` | - | - | - | - | - |
| `fxDestinationCurrency` | - | - | - | - | - |
| `fxQuoteExpiresAt` | - | - | - | - | - |
| `fxQuoteID` | - | - | - | - | - |
| `fxRate` | - | - | - | - | - |
| `fxUseDestinationAmount` | - | - | - | - | - |
| `hasCurrentUserAuthorised` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `inserted` | - | - | - | - | - |
| `invoiceID` | - | - | - | - | - |
| `isArchived` | - | - | - | - | - |
| `isFailed` | - | - | - | - | - |
| `isSettled` | - | - | - | - | - |
| `isSubmitted` | - | - | - | - | - |
| `lastUpdated` | - | - | - | - | - |
| `merchantID` | - | - | - | - | - |
| `merchantTokenDescription` | - | - | - | - | - |
| `nonce` | - | - | - | - | - |
| `paymentProcessor` | - | - | - | - | - |
| `paymentRail` | - | - | - | - | - |
| `payouts` | - | - | - | - | - |
| `payrunID` | - | - | - | - | - |
| `payrunName` | - | - | - | - | - |
| `reason` | - | - | - | - | - |
| `rule` | - | - | - | - | - |
| `scheduleDate` | - | - | - | - | - |
| `scheduled` | - | - | - | - | - |
| `sourceAccountAvailableBalance` | - | - | - | - | - |
| `sourceAccountAvailableBalanceMinorUnits` | - | - | - | - | - |
| `sourceAccountBic` | - | - | - | - | - |
| `sourceAccountCurrency` | - | - | - | - | - |
| `sourceAccountIban` | - | - | - | - | - |
| `sourceAccountIdentifier` | - | - | - | - | - |
| `sourceAccountName` | - | - | - | - | - |
| `sourceAccountNumber` | - | - | - | - | - |
| `sourceAccountSortcode` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `tagIds` | - | - | - | - | - |
| `tags` | - | - | - | - | - |
| `theirReference` | - | - | - | - | - |
| `topupPayrunID` | - | - | - | - | - |
| `transactedAmount` | - | - | - | - | - |
| `transactedFxAmount` | - | - | - | - | - |
| `transactedFxRate` | - | - | - | - | - |
| `type` | - | - | Yes | - | - |
| `userID` | - | - | - | - | - |
| `yourReference` | - | - | - | - | - |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Payout()->create([
  "beneficiary" => null, // array
  "sourceAccountIdentifier" => null, // array
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Payout()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Payout()->load(["id" => "payout_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Payout()->remove(["id" => "payout_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Payout()->update([
  "id" => "payout_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PayoutEntity`

Create a new `PayoutEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PayoutKeysetPageEntity

```php
$payout_keyset_page = $client->PayoutKeysetPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountID` | `string` | No |  |
| `amount` | `float` | No |  |
| `amountMinorUnits` | `int` | No |  |
| `approvePayoutUrl` | `string` | No |  |
| `approverID` | `string` | No |  |
| `authenticationMethods` | `array` | No |  |
| `authorisations` | `array` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `batchPayoutID` | `string` | No |  |
| `beneficiary` | `array` | Yes |  |
| `canAuthorise` | `bool` | No |  |
| `canProcess` | `bool` | No |  |
| `canUpdate` | `bool` | No |  |
| `chargeBearer` | `string` | No |  |
| `createdBy` | `string` | No |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | No |  |
| `currentUserID` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `array` | No |  |
| `documents` | `array` | No |  |
| `events` | `array` | No |  |
| `formattedAmount` | `string` | No |  |
| `formattedFxDestinationAmount` | `string` | No |  |
| `formattedSchedule` | `string` | No |  |
| `formattedScheduleDayOnly` | `string` | No |  |
| `formattedSourceAccountAvailableBalance` | `string` | No |  |
| `fxDestinationAmount` | `float` | No |  |
| `fxDestinationAmountMinorUnits` | `int` | No |  |
| `fxDestinationCurrency` | `string` | No |  |
| `fxQuoteExpiresAt` | `string` | No |  |
| `fxQuoteID` | `string` | No |  |
| `fxRate` | `float` | No |  |
| `fxUseDestinationAmount` | `bool` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoiceID` | `string` | No |  |
| `isArchived` | `bool` | No |  |
| `isFailed` | `bool` | No |  |
| `isSettled` | `bool` | No |  |
| `isSubmitted` | `bool` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantTokenDescription` | `string` | No |  |
| `nonce` | `string` | No |  |
| `paymentProcessor` | `string` | No |  |
| `paymentRail` | `string` | No |  |
| `payrunID` | `string` | No |  |
| `payrunName` | `string` | No |  |
| `rule` | `array` | No |  |
| `scheduleDate` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `sourceAccountAvailableBalance` | `float` | No |  |
| `sourceAccountAvailableBalanceMinorUnits` | `int` | No |  |
| `sourceAccountBic` | `string` | No |  |
| `sourceAccountCurrency` | `string` | No |  |
| `sourceAccountIban` | `string` | No |  |
| `sourceAccountIdentifier` | `array` | Yes |  |
| `sourceAccountName` | `string` | No |  |
| `sourceAccountNumber` | `string` | No |  |
| `sourceAccountSortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tags` | `array` | No |  |
| `theirReference` | `string` | No |  |
| `topupPayrunID` | `string` | No |  |
| `transactedAmount` | `float` | No |  |
| `transactedFxAmount` | `float` | No |  |
| `transactedFxRate` | `float` | No |  |
| `type` | `string` | No |  |
| `userID` | `string` | No |  |
| `yourReference` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->PayoutKeysetPage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PayoutKeysetPageEntity`

Create a new `PayoutKeysetPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PayoutMetricEntity

```php
$payout_metric = $client->PayoutMetric();
```

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->PayoutMetric()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PayoutMetricEntity`

Create a new `PayoutMetricEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PayrunEntity

```php
$payrun = $client->Payrun();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisationDate` | `string` | No |  |
| `authorisations` | `array` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `batchPayoutID` | `string` | No |  |
| `canAuthorise` | `bool` | No |  |
| `canDelete` | `bool` | No |  |
| `canEdit` | `bool` | No |  |
| `events` | `array` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoices` | `array` | No |  |
| `invoicesMinimal` | `array` | No |  |
| `isArchived` | `bool` | No |  |
| `lastUpdated` | `string` | No |  |
| `lastUpdatedBy` | `array` | Yes |  |
| `merchantID` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | No |  |
| `notes` | `string` | No |  |
| `payments` | `array` | No |  |
| `payouts` | `array` | No |  |
| `payoutsCount` | `int` | No |  |
| `reason` | `string` | No |  |
| `scheduleDate` | `string` | No |  |
| `scheduledDate` | `string` | No |  |
| `sourceAccounts` | `array` | No |  |
| `status` | `string` | No |  |
| `totalEur` | `float` | No |  |
| `totalGbp` | `float` | No |  |
| `totalUsd` | `float` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Payrun()->create([
  "id" => null, // string
  "lastUpdatedBy" => null, // array
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Payrun()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Payrun()->load(["id" => "payrun_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Payrun()->remove(["id" => "payrun_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Payrun()->update([
  "id" => "payrun_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PayrunEntity`

Create a new `PayrunEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ReportEntity

```php
$report = $client->Report();
```

### Operations

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Report()->update([
  "id" => "id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ReportEntity`

Create a new `ReportEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ReportResultEntity

```php
$report_result = $client->ReportResult();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contentType` | `string` | No |  |
| `contents` | `string` | No |  |
| `lastCompletedAt` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `reportName` | `string` | No |  |
| `reportType` | `string` | No |  |
| `statementNumber` | `int` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->ReportResult()->load(["id" => 1, "report_id" => "report_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ReportResultEntity`

Create a new `ReportResultEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## RoleEntity

```php
$role = $client->Role();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failedRoles` | `array` | No |  |
| `roles` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Role()->create([
  "merchant_id" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): RoleEntity`

Create a new `RoleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## RuleEntity

```php
$rule = $client->Rule();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account` | `array` | No |  |
| `accountID` | `string` | No |  |
| `approveUrl` | `string` | No |  |
| `approverID` | `string` | No |  |
| `authenticationMethods` | `array` | No |  |
| `authorisations` | `array` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `canAuthorise` | `bool` | No |  |
| `createdBy` | `array` | Yes |  |
| `description` | `string` | No |  |
| `endAt` | `string` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isDisabled` | `bool` | No |  |
| `lastExecutedAt` | `string` | No |  |
| `lastRunAtTransactionDate` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | Yes |  |
| `onApprovedWebHookUrl` | `string` | No |  |
| `onExecutionErrorWebHookUrl` | `string` | No |  |
| `onExecutionSuccessWebHookUrl` | `string` | No |  |
| `startAt` | `string` | No |  |
| `status` | `string` | No |  |
| `sweepAction` | `array` | No |  |
| `timeZoneId` | `string` | No |  |
| `triggerCronExpression` | `string` | No |  |
| `triggerOnPayIn` | `bool` | No |  |
| `userID` | `string` | No |  |
| `webHookSecret` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `account` | - | - | - | - | - |
| `accountID` | - | - | - | - | - |
| `approveUrl` | - | - | - | - | - |
| `approverID` | - | - | - | - | - |
| `authenticationMethods` | - | - | - | - | - |
| `authorisations` | - | - | - | - | - |
| `authorisersCompletedCount` | - | - | - | - | - |
| `authorisersRequiredCount` | - | - | - | - | - |
| `canAuthorise` | - | - | - | - | - |
| `createdBy` | - | - | - | - | - |
| `description` | - | - | - | - | - |
| `endAt` | - | - | - | - | - |
| `hasCurrentUserAuthorised` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `inserted` | - | - | - | - | - |
| `isDisabled` | - | - | - | - | - |
| `lastExecutedAt` | - | - | - | - | - |
| `lastRunAtTransactionDate` | - | - | - | - | - |
| `lastUpdated` | - | - | - | - | - |
| `merchantID` | - | - | - | - | - |
| `name` | - | - | Yes | - | - |
| `nonce` | - | - | - | - | - |
| `onApprovedWebHookUrl` | - | - | - | - | - |
| `onExecutionErrorWebHookUrl` | - | - | - | - | - |
| `onExecutionSuccessWebHookUrl` | - | - | - | - | - |
| `startAt` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `sweepAction` | - | - | Yes | - | - |
| `timeZoneId` | - | - | - | - | - |
| `triggerCronExpression` | - | - | - | - | - |
| `triggerOnPayIn` | - | - | - | - | - |
| `userID` | - | - | - | - | - |
| `webHookSecret` | - | - | - | - | - |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Rule()->create([
  "createdBy" => null, // array
  "nonce" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Rule()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Rule()->load(["id" => "rule_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Rule()->remove(["id" => "rule_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Rule()->update([
  "id" => "rule_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): RuleEntity`

Create a new `RuleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## RuleEventEntity

```php
$rule_event = $client->RuleEvent();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errorMessage` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isAuthoriseToEnable` | `bool` | No |  |
| `message` | `string` | No |  |
| `rawResponse` | `string` | No |  |
| `ruleEventType` | `string` | No |  |
| `ruleID` | `string` | No |  |
| `user` | `array` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->RuleEvent()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): RuleEventEntity`

Create a new `RuleEventEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TagEntity

```php
$tag = $client->Tag();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `colourHex` | `string` | No |  |
| `description` | `string` | No |  |
| `id` | `string` | No |  |
| `merchantID` | `string` | Yes |  |
| `name` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Tag()->create([
  "merchant_id" => null, // string
  "merchantID" => null, // string
  "name" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Tag()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TagEntity`

Create a new `TagEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TokenEntity

```php
$token = $client->Token();
```

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Token()->create([
  "id" => null, // string
]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Token()->remove(["id" => "id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TokenEntity`

Create a new `TokenEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TransactionEntity

```php
$transaction = $client->Transaction();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountID` | `string` | No |  |
| `accountName` | `string` | No |  |
| `accountSequenceNumber` | `int` | No |  |
| `addressDetails` | `array` | No |  |
| `amount` | `float` | No |  |
| `amountMinorUnits` | `int` | No |  |
| `balance` | `float` | No |  |
| `balanceMinorUnits` | `int` | No |  |
| `bookingDateTime` | `string` | No |  |
| `chargeDetails` | `array` | No |  |
| `content` | `array` | No |  |
| `counterparty` | `array` | No |  |
| `counterpartySummary` | `string` | No |  |
| `currency` | `string` | No |  |
| `currencyExchange` | `array` | No |  |
| `date` | `string` | No |  |
| `description` | `string` | No |  |
| `enrichment` | `array` | No |  |
| `fxAmount` | `float` | No |  |
| `fxCurrency` | `string` | No |  |
| `fxRate` | `float` | No |  |
| `grossAmount` | `array` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isoBankTransactionCode` | `array` | No |  |
| `merchant` | `array` | No |  |
| `merchantID` | `string` | No |  |
| `pageNumber` | `int` | No |  |
| `pageSize` | `int` | No |  |
| `payeeDetails` | `array` | Yes |  |
| `payerDetails` | `array` | Yes |  |
| `paymentRequestCustomFields` | `array` | No |  |
| `paymentRequestID` | `string` | No |  |
| `payoutID` | `string` | No |  |
| `proprietaryBankTransactionCode` | `array` | No |  |
| `rawReference` | `string` | No |  |
| `reference` | `string` | No |  |
| `ruleID` | `string` | No |  |
| `statementReferences` | `array` | No |  |
| `status` | `string` | No |  |
| `supplementaryData` | `mixed` | No |  |
| `tags` | `array` | No |  |
| `theirReference` | `string` | No |  |
| `totalPages` | `int` | No |  |
| `totalSize` | `int` | No |  |
| `transactionAmount` | `array` | Yes |  |
| `transactionDate` | `string` | No |  |
| `transactionInformation` | `array` | No |  |
| `transactionMutability` | `string` | No |  |
| `type` | `string` | No |  |
| `valueDateTime` | `string` | No |  |
| `virtualIBAN` | `string` | No |  |
| `yourReference` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Transaction()->create([
  "id" => null, // string
  "grossAmount" => null, // array
  "payeeDetails" => null, // array
  "payerDetails" => null, // array
  "transactionAmount" => null, // array
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Transaction()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Transaction()->load(["id" => "transaction_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Transaction()->remove(["id" => "transaction_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TransactionEntity`

Create a new `TransactionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## UserEntity

```php
$user = $client->User();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `clientSessionTimeouts` | `array` | No |  |
| `emailAddress` | `string` | Yes |  |
| `firstName` | `string` | Yes |  |
| `id` | `string` | No |  |
| `lastName` | `string` | Yes |  |
| `passkeyAdded` | `bool` | No |  |
| `permissions` | `array` | No |  |
| `profile` | `string` | No |  |
| `rolesWithScope` | `array` | No |  |
| `twoFactorEnabled` | `bool` | No |  |
| `userInviteID` | `string` | No |  |

### Field Usage by Operation

| Field | list | update |
| --- | --- | --- |
| `clientSessionTimeouts` | - | - |
| `emailAddress` | - | Yes |
| `firstName` | - | Yes |
| `id` | - | - |
| `lastName` | - | Yes |
| `passkeyAdded` | - | - |
| `permissions` | - | - |
| `profile` | - | - |
| `rolesWithScope` | - | - |
| `twoFactorEnabled` | - | - |
| `userInviteID` | - | - |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->User()->list();
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->User()->update([
  "id" => "id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): UserEntity`

Create a new `UserEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## UserInviteEntity

```php
$user_invite = $client->UserInvite();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisationStatus` | `array` | No |  |
| `failedUserInvites` | `array` | No |  |
| `id` | `string` | No |  |
| `initialRoleID` | `string` | No |  |
| `inviteeEmailAddress` | `string` | No |  |
| `inviteeFirstName` | `string` | No |  |
| `inviteeLastName` | `string` | No |  |
| `inviterEmailAddress` | `string` | No |  |
| `inviterFirstName` | `string` | No |  |
| `inviterLastName` | `string` | No |  |
| `isAuthorised` | `bool` | No |  |
| `isInviteeRegistered` | `bool` | No |  |
| `lastInvited` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `message` | `string` | No |  |
| `registrationUrl` | `string` | No |  |
| `sendInviteEmail` | `bool` | No |  |
| `status` | `string` | No |  |
| `user` | `array` | Yes |  |
| `userID` | `string` | No |  |
| `userInvites` | `array` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `authorisationStatus` | - | - | - | - | - |
| `failedUserInvites` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `initialRoleID` | - | - | - | - | - |
| `inviteeEmailAddress` | - | - | Yes | - | - |
| `inviteeFirstName` | - | - | - | - | - |
| `inviteeLastName` | - | - | - | - | - |
| `inviterEmailAddress` | - | - | - | - | - |
| `inviterFirstName` | - | - | - | - | - |
| `inviterLastName` | - | - | - | - | - |
| `isAuthorised` | - | - | - | - | - |
| `isInviteeRegistered` | - | - | - | - | - |
| `lastInvited` | - | - | - | - | - |
| `merchantID` | - | - | - | - | - |
| `merchantName` | - | - | - | - | - |
| `message` | - | - | - | - | - |
| `registrationUrl` | - | - | - | - | - |
| `sendInviteEmail` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `user` | - | - | - | - | - |
| `userID` | - | - | - | - | - |
| `userInvites` | - | - | - | - | - |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->UserInvite()->create([
  "user" => null, // array
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->UserInvite()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->UserInvite()->load(["id" => "user_invite_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->UserInvite()->remove(["id" => "user_invite_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->UserInvite()->update([
  "id" => "user_invite_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): UserInviteEntity`

Create a new `UserInviteEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## VirtualEntity

```php
$virtual = $client->Virtual();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `string` | No |  |
| `accountSupplierName` | `string` | No |  |
| `availableBalance` | `float` | No |  |
| `availableBalanceMinorUnits` | `int` | No |  |
| `balance` | `float` | No |  |
| `balanceMinorUnits` | `int` | No |  |
| `bankName` | `string` | No |  |
| `consentID` | `string` | No |  |
| `createdBy` | `array` | Yes |  |
| `createdByDisplayName` | `string` | No |  |
| `currency` | `string` | No |  |
| `defaultPaymentRail` | `string` | No |  |
| `displayName` | `string` | No |  |
| `expiryDate` | `string` | No |  |
| `externalAccountIcon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `array` | Yes |  |
| `inserted` | `string` | No |  |
| `isArchived` | `bool` | No |  |
| `isConnectedAccount` | `bool` | No |  |
| `isDefault` | `bool` | No |  |
| `isTrustAccount` | `bool` | No |  |
| `isVirtual` | `bool` | No |  |
| `lastTransaction` | `array` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `name` | `string` | Yes |  |
| `physicalAccountID` | `string` | No |  |
| `rules` | `array` | No |  |
| `submittedPayoutsBalance` | `float` | No |  |
| `submittedPayoutsBalanceMinorUnits` | `int` | No |  |
| `summary` | `string` | No |  |
| `supplierSepaInstantStatus` | `string` | No |  |
| `xeroBankFeedConnectionStatus` | `string` | No |  |
| `xeroBankFeedLastSyncedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `string` | No |  |
| `xeroBankFeedSyncStatus` | `string` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Virtual()->create([
  "account_id" => null, // string
  "createdBy" => null, // array
  "identifier" => null, // array
  "name" => null, // string
]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Virtual()->update([
  "account_id" => "account_id",
  "id" => "id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): VirtualEntity`

Create a new `VirtualEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## WebhookEntity

```php
$webhook = $client->Webhook();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destinationUrl` | `string` | No |  |
| `emailAddress` | `string` | No |  |
| `failedNotificationEmailAddress` | `string` | No |  |
| `id` | `string` | No |  |
| `isActive` | `bool` | No |  |
| `merchantID` | `string` | No |  |
| `notificationMethod` | `string` | No |  |
| `resourceTypes` | `array` | No |  |
| `retry` | `bool` | No |  |
| `secret` | `string` | No |  |
| `version` | `int` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `destinationUrl` | - | - | - | - | - |
| `emailAddress` | - | - | - | - | - |
| `failedNotificationEmailAddress` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `isActive` | - | - | - | - | - |
| `merchantID` | - | - | Yes | Yes | - |
| `notificationMethod` | - | - | Yes | Yes | - |
| `resourceTypes` | - | - | - | - | - |
| `retry` | - | - | - | - | - |
| `secret` | - | - | - | - | - |
| `version` | - | - | - | - | - |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Webhook()->create([
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Webhook()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Webhook()->load(["id" => "webhook_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Webhook()->remove(["id" => "webhook_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Webhook()->update([
  "id" => "webhook_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): WebhookEntity`

Create a new `WebhookEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new NofrixionSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

