# Nofrixion Lua SDK Reference

Complete API reference for the Nofrixion Lua SDK.


## NofrixionSDK

### Constructor

```lua
local sdk = require("nofrixion_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `Account(data)`

Create a new `Account` entity instance. Pass `nil` for no initial data.

#### `Batch(data)`

Create a new `Batch` entity instance. Pass `nil` for no initial data.

#### `Beneficiary(data)`

Create a new `Beneficiary` entity instance. Pass `nil` for no initial data.

#### `BeneficiaryGroup(data)`

Create a new `BeneficiaryGroup` entity instance. Pass `nil` for no initial data.

#### `Card(data)`

Create a new `Card` entity instance. Pass `nil` for no initial data.

#### `CardCustomerToken(data)`

Create a new `CardCustomerToken` entity instance. Pass `nil` for no initial data.

#### `CardPayment(data)`

Create a new `CardPayment` entity instance. Pass `nil` for no initial data.

#### `CardPublicKey(data)`

Create a new `CardPublicKey` entity instance. Pass `nil` for no initial data.

#### `Consent(data)`

Create a new `Consent` entity instance. Pass `nil` for no initial data.

#### `Currency(data)`

Create a new `Currency` entity instance. Pass `nil` for no initial data.

#### `DirectDebitBatchSubmit(data)`

Create a new `DirectDebitBatchSubmit` entity instance. Pass `nil` for no initial data.

#### `FxRate(data)`

Create a new `FxRate` entity instance. Pass `nil` for no initial data.

#### `IPayment(data)`

Create a new `IPayment` entity instance. Pass `nil` for no initial data.

#### `Mandate(data)`

Create a new `Mandate` entity instance. Pass `nil` for no initial data.

#### `Merchant(data)`

Create a new `Merchant` entity instance. Pass `nil` for no initial data.

#### `MerchantAuthorisationSetting(data)`

Create a new `MerchantAuthorisationSetting` entity instance. Pass `nil` for no initial data.

#### `MerchantDirectDebitMandatePage(data)`

Create a new `MerchantDirectDebitMandatePage` entity instance. Pass `nil` for no initial data.

#### `MerchantPayByBankSetting(data)`

Create a new `MerchantPayByBankSetting` entity instance. Pass `nil` for no initial data.

#### `MerchantPaymentRequestTemplate(data)`

Create a new `MerchantPaymentRequestTemplate` entity instance. Pass `nil` for no initial data.

#### `MerchantToken(data)`

Create a new `MerchantToken` entity instance. Pass `nil` for no initial data.

#### `Metadata(data)`

Create a new `Metadata` entity instance. Pass `nil` for no initial data.

#### `NoFrixionVersion(data)`

Create a new `NoFrixionVersion` entity instance. Pass `nil` for no initial data.

#### `OpenBanking(data)`

Create a new `OpenBanking` entity instance. Pass `nil` for no initial data.

#### `Payeeverification(data)`

Create a new `Payeeverification` entity instance. Pass `nil` for no initial data.

#### `Payment(data)`

Create a new `Payment` entity instance. Pass `nil` for no initial data.

#### `PaymentAccount(data)`

Create a new `PaymentAccount` entity instance. Pass `nil` for no initial data.

#### `PaymentAccountMinimal(data)`

Create a new `PaymentAccountMinimal` entity instance. Pass `nil` for no initial data.

#### `PaymentInitiation(data)`

Create a new `PaymentInitiation` entity instance. Pass `nil` for no initial data.

#### `PaymentRequest(data)`

Create a new `PaymentRequest` entity instance. Pass `nil` for no initial data.

#### `PaymentRequestEvent(data)`

Create a new `PaymentRequestEvent` entity instance. Pass `nil` for no initial data.

#### `PaymentRequestMetric(data)`

Create a new `PaymentRequestMetric` entity instance. Pass `nil` for no initial data.

#### `PaymentRequestMinimal(data)`

Create a new `PaymentRequestMinimal` entity instance. Pass `nil` for no initial data.

#### `PaymentRequestResult(data)`

Create a new `PaymentRequestResult` entity instance. Pass `nil` for no initial data.

#### `Payout(data)`

Create a new `Payout` entity instance. Pass `nil` for no initial data.

#### `PayoutKeysetPage(data)`

Create a new `PayoutKeysetPage` entity instance. Pass `nil` for no initial data.

#### `PayoutMetric(data)`

Create a new `PayoutMetric` entity instance. Pass `nil` for no initial data.

#### `Payrun(data)`

Create a new `Payrun` entity instance. Pass `nil` for no initial data.

#### `Report(data)`

Create a new `Report` entity instance. Pass `nil` for no initial data.

#### `ReportResult(data)`

Create a new `ReportResult` entity instance. Pass `nil` for no initial data.

#### `Role(data)`

Create a new `Role` entity instance. Pass `nil` for no initial data.

#### `Rule(data)`

Create a new `Rule` entity instance. Pass `nil` for no initial data.

#### `RuleEvent(data)`

Create a new `RuleEvent` entity instance. Pass `nil` for no initial data.

#### `Tag(data)`

Create a new `Tag` entity instance. Pass `nil` for no initial data.

#### `Token(data)`

Create a new `Token` entity instance. Pass `nil` for no initial data.

#### `Transaction(data)`

Create a new `Transaction` entity instance. Pass `nil` for no initial data.

#### `User(data)`

Create a new `User` entity instance. Pass `nil` for no initial data.

#### `UserInvite(data)`

Create a new `UserInvite` entity instance. Pass `nil` for no initial data.

#### `Virtual(data)`

Create a new `Virtual` entity instance. Pass `nil` for no initial data.

#### `Webhook(data)`

Create a new `Webhook` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## AccountEntity

```lua
local account = client:Account(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountBalances` | `table` | No |  |
| `accountID` | `string` | No |  |
| `accountIdentifications` | `table` | No |  |
| `accountName` | `string` | No |  |
| `accountNames` | `table` | No |  |
| `accountSupplierName` | `string` | No |  |
| `accountType` | `string` | No |  |
| `availableBalance` | `number` | No |  |
| `availableBalanceMinorUnits` | `number` | No |  |
| `balance` | `number` | No |  |
| `balanceMinorUnits` | `number` | No |  |
| `bankName` | `string` | No |  |
| `consentID` | `string` | No |  |
| `consolidatedAccountInformation` | `table` | No |  |
| `createdBy` | `table` | Yes |  |
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
| `identifier` | `table` | Yes |  |
| `inserted` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `isConnectedAccount` | `boolean` | No |  |
| `isDefault` | `boolean` | No |  |
| `isTrustAccount` | `boolean` | No |  |
| `isVirtual` | `boolean` | No |  |
| `lastTransaction` | `table` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `nickname` | `string` | No |  |
| `physicalAccountID` | `string` | No |  |
| `roleIDs` | `table` | No |  |
| `rules` | `table` | No |  |
| `submittedPayoutsBalance` | `number` | No |  |
| `submittedPayoutsBalanceMinorUnits` | `number` | No |  |
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
| `xeroUnsynchronisedTransactionsCount` | `number` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Account():create({
  createdBy = --[[ table ]],
  identifier = --[[ table ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Account():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Account():load({ id = "account_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Account():remove({ id = "account_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Account():update({
  id = "account_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AccountEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## BatchEntity

```lua
local batch = client:Batch(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approveUrl` | `string` | No |  |
| `id` | `string` | No |  |
| `payouts` | `table` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Batch():create({
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Batch():load({ id = "batch_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BatchEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## BeneficiaryEntity

```lua
local beneficiary = client:Beneficiary(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approvalCallbackUrl` | `string` | No |  |
| `authenticationMethods` | `table` | No |  |
| `authorisations` | `table` | No |  |
| `authorisersCompletedCount` | `number` | No |  |
| `authorisersRequiredCount` | `number` | No |  |
| `beneficiaries` | `table` | No |  |
| `beneficiaryEvents` | `table` | No |  |
| `canAuthorise` | `boolean` | No |  |
| `canUpdate` | `boolean` | No |  |
| `createdBy` | `table` | Yes |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `table` | No |  |
| `failedBeneficiaries` | `table` | No |  |
| `hasCurrentUserAuthorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isEnabled` | `boolean` | No |  |
| `lastAuthorised` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `sourceAccountIDs` | `table` | No |  |
| `sourceAccounts` | `table` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Beneficiary():create({
  createdBy = --[[ table ]],
  currency = --[[ string ]],
  name = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Beneficiary():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Beneficiary():load({ id = "beneficiary_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Beneficiary():remove({ id = "beneficiary_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Beneficiary():update({
  id = "beneficiary_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BeneficiaryEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## BeneficiaryGroupEntity

```lua
local beneficiary_group = client:BeneficiaryGroup(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `groupMembers` | `table` | No |  |
| `groupName` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | Yes |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:BeneficiaryGroup():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BeneficiaryGroupEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CardEntity

```lua
local card = client:Card(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorizedAmount` | `string` | No |  |
| `currencyCode` | `string` | No |  |
| `isPayerAuthenticationRequired` | `boolean` | No |  |
| `isSoftDecline` | `boolean` | No |  |
| `payerAuthenticationAccessToken` | `string` | No |  |
| `payerAuthenticationMerchantData` | `string` | No |  |
| `payerAuthenticationUrl` | `string` | No |  |
| `payerAuthenticationWindowHeight` | `number` | No |  |
| `payerAuthenticationWindowWidth` | `number` | No |  |
| `paymentRequestCallbackUrl` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |
| `requestID` | `string` | No |  |
| `responseCode` | `string` | No |  |
| `responseType` | `string` | No |  |
| `status` | `string` | No |  |
| `threeDSRedirectUrl` | `string` | No |  |
| `transactionID` | `string` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Card():create({
  paymentrequest_id = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CardEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CardCustomerTokenEntity

```lua
local card_customer_token = client:CardCustomerToken(nil)
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:CardCustomerToken():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:CardCustomerToken():load({ customer_email_address = "customer_email_address" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:CardCustomerToken():remove()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CardCustomerTokenEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CardPaymentEntity

```lua
local card_payment = client:CardPayment(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorizedAmount` | `string` | No |  |
| `currencyCode` | `string` | No |  |
| `isPayerAuthenticationRequired` | `boolean` | No |  |
| `isSoftDecline` | `boolean` | No |  |
| `payerAuthenticationAccessToken` | `string` | No |  |
| `payerAuthenticationMerchantData` | `string` | No |  |
| `payerAuthenticationUrl` | `string` | No |  |
| `payerAuthenticationWindowHeight` | `number` | No |  |
| `payerAuthenticationWindowWidth` | `number` | No |  |
| `paymentRequestCallbackUrl` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |
| `requestID` | `string` | No |  |
| `responseCode` | `string` | No |  |
| `responseType` | `string` | No |  |
| `status` | `string` | No |  |
| `threeDSRedirectUrl` | `string` | No |  |
| `transactionID` | `string` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:CardPayment():create({
  paymentrequest_id = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CardPaymentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CardPublicKeyEntity

```lua
local card_public_key = client:CardPublicKey(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jwt` | `string` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:CardPublicKey():load({ paymentrequest_id = "paymentrequest_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CardPublicKeyEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ConsentEntity

```lua
local consent = client:Consent(nil)
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
| `isConnectedAccounts` | `boolean` | No |  |
| `isEnabled` | `boolean` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Consent():create({
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Consent():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Consent():load({ id = "consent_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Consent():remove({ id = "consent_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Consent():update({
  id = "consent_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ConsentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## CurrencyEntity

```lua
local currency = client:Currency(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `code` | `string` | No |  |
| `decimals` | `number` | No |  |
| `isFiat` | `boolean` | No |  |
| `iso4217AlphaCode` | `string` | No |  |
| `iso4217NumericCode` | `string` | No |  |
| `symbol` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Currency():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CurrencyEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## DirectDebitBatchSubmitEntity

```lua
local direct_debit_batch_submit = client:DirectDebitBatchSubmit(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failedSubmissions` | `table` | No |  |
| `successfulSubmissions` | `table` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:DirectDebitBatchSubmit():create({
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DirectDebitBatchSubmitEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## FxRateEntity

```lua
local fx_rate = client:FxRate(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destinationCurrency` | `string` | No |  |
| `exchangeRate` | `number` | No |  |
| `expiryTime` | `string` | No |  |
| `quoteID` | `string` | No |  |
| `sourceCurrency` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:FxRate():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:FxRate():load({ destination = "destination", source = "source", valid_for_minute = 1 })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FxRateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## IPaymentEntity

```lua
local i_payment = client:IPayment(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `paymentRequestID` | `string` | No |  |
| `responseType` | `string` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:IPayment():create({
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `IPaymentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MandateEntity

```lua
local mandate = client:Mandate(nil)
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
| `isRecurring` | `boolean` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Mandate():create({
  addressLine1 = --[[ string ]],
  city = --[[ string ]],
  countryCode = --[[ string ]],
  emailAddress = --[[ string ]],
  firstName = --[[ string ]],
  lastName = --[[ string ]],
  postalCode = --[[ string ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Mandate():load({ id = "mandate_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MandateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MerchantEntity

```lua
local merchant = client:Merchant(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountCurrencies` | `table` | No |  |
| `canHaveTrustAccounts` | `boolean` | No |  |
| `cardPaymentProcessor` | `string` | No |  |
| `companyID` | `string` | No |  |
| `displayQrOnHostedPay` | `boolean` | No |  |
| `hostedPayVersion` | `number` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isBlocked` | `boolean` | No |  |
| `isExited` | `boolean` | No |  |
| `isSuspended` | `boolean` | No |  |
| `jurisdiction` | `string` | No |  |
| `logoUrlPng` | `string` | No |  |
| `logoUrlSvg` | `string` | No |  |
| `merchantCategoryCode` | `string` | No |  |
| `name` | `string` | No |  |
| `notes` | `string` | No |  |
| `parentMerchant` | `table` | No |  |
| `paymentAccountLimit` | `number` | No |  |
| `paymentAccounts` | `table` | No |  |
| `reason` | `string` | No |  |
| `shortName` | `string` | No |  |
| `supportedPaymentMethodsList` | `table` | No |  |
| `suspensionReason` | `string` | No |  |
| `tags` | `table` | No |  |
| `timeZoneId` | `string` | No |  |
| `tradingName` | `string` | No |  |
| `webHookLimit` | `number` | No |  |
| `yourRoleName` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Merchant():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Merchant():load({ id = "merchant_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Merchant():remove({ id = "merchant_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Merchant():update({
  id = "merchant_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MerchantEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MerchantAuthorisationSettingEntity

```lua
local merchant_authorisation_setting = client:MerchantAuthorisationSetting(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amountLower` | `number` | No |  |
| `amountUpper` | `number` | No |  |
| `authorisationType` | `string` | No |  |
| `beneficiariesOnly` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lastEditorCantAuthorise` | `boolean` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `numberOfAuthorisers` | `number` | No |  |
| `roleSettings` | `table` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:MerchantAuthorisationSetting():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MerchantAuthorisationSettingEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MerchantDirectDebitMandatePageEntity

```lua
local merchant_direct_debit_mandate_page = client:MerchantDirectDebitMandatePage(nil)
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
| `isRecurring` | `boolean` | No |  |
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:MerchantDirectDebitMandatePage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MerchantDirectDebitMandatePageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MerchantPayByBankSettingEntity

```lua
local merchant_pay_by_bank_setting = client:MerchantPayByBankSetting(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bankCountryCodes` | `table` | No |  |
| `bankID` | `string` | No |  |
| `bankName` | `string` | No |  |
| `businessInstitutionID` | `string` | No |  |
| `currency` | `string` | No |  |
| `logo` | `string` | No |  |
| `message` | `string` | No |  |
| `messageImageUrl` | `string` | No |  |
| `order` | `number` | No |  |
| `personalInstitutionID` | `string` | No |  |
| `processor` | `string` | No |  |
| `warningHeading` | `string` | No |  |
| `warningMessage` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:MerchantPayByBankSetting():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MerchantPayByBankSettingEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MerchantPaymentRequestTemplateEntity

```lua
local merchant_payment_request_template = client:MerchantPaymentRequestTemplate(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bankPaymentOptions` | `table` | No |  |
| `cardPaymentAddressOptions` | `table` | No |  |
| `cardPaymentCaptureOptions` | `table` | No |  |
| `customFields` | `table` | No |  |
| `defaultFields` | `table` | No |  |
| `description` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `name` | `string` | Yes |  |
| `notificationOptions` | `table` | No |  |
| `paymentMethods` | `table` | No |  |
| `paymentTerms` | `table` | No |  |
| `priorityBankOptions` | `table` | No |  |
| `template` | `table` | Yes |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:MerchantPaymentRequestTemplate():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:MerchantPaymentRequestTemplate():load({ id = "merchant_payment_request_template_id", paymentrequest_id = "paymentrequest_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:MerchantPaymentRequestTemplate():remove({ id = "merchant_payment_request_template_id", paymentrequest_id = "paymentrequest_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:MerchantPaymentRequestTemplate():update({
  id = "merchant_payment_request_template_id",
  paymentrequest_id = "paymentrequest_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MerchantPaymentRequestTemplateEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MerchantTokenEntity

```lua
local merchant_token = client:MerchantToken(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authenticationMethods` | `table` | No |  |
| `authorisations` | `table` | No |  |
| `authorisersCompletedCount` | `number` | No |  |
| `authorisersRequiredCount` | `number` | No |  |
| `canAuthorise` | `boolean` | No |  |
| `description` | `string` | No |  |
| `expiresAt` | `string` | No |  |
| `hasCurrentUserAuthorised` | `boolean` | No |  |
| `hmacAlgorithm` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `ipAddressWhitelist` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `isEnabled` | `boolean` | No |  |
| `lastAuthorised` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `nonce` | `string` | Yes |  |
| `permissionTypes` | `table` | No |  |
| `requestSignatureVersion` | `number` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:MerchantToken():create({
  nonce = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:MerchantToken():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:MerchantToken():load({ id = "merchant_token_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:MerchantToken():update({
  id = "merchant_token_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MerchantTokenEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## MetadataEntity

```lua
local metadata = client:Metadata(nil)
```

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Metadata():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MetadataEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## NoFrixionVersionEntity

```lua
local no_frixion_version = client:NoFrixionVersion(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `buildVersion` | `number` | No |  |
| `majorVersion` | `number` | No |  |
| `minorVersion` | `number` | No |  |
| `releaseName` | `string` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:NoFrixionVersion():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionVersionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## OpenBankingEntity

```lua
local open_banking = client:OpenBanking(nil)
```

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:OpenBanking():create({
  account_id = --[[ string ]],
})
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:OpenBanking():remove()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OpenBankingEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PayeeverificationEntity

```lua
local payeeverification = client:Payeeverification(nil)
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Payeeverification():create({
  accountName = --[[ string ]],
  iban = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PayeeverificationEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PaymentEntity

```lua
local payment = client:Payment(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresses` | `table` | No |  |
| `amount` | `number` | No |  |
| `amountPending` | `number` | No |  |
| `amountReceived` | `number` | No |  |
| `amountRefunded` | `number` | No |  |
| `autoSendReceipt` | `boolean` | No |  |
| `baseOriginUrl` | `string` | No |  |
| `callbackUrl` | `string` | No |  |
| `cardAuthorizeOnly` | `boolean` | No |  |
| `cardCreateToken` | `boolean` | No |  |
| `cardCreateTokenMode` | `string` | No |  |
| `cardIgnoreCVN` | `boolean` | No |  |
| `cardNoPayerAuthentication` | `boolean` | No |  |
| `cardProcessorMerchantID` | `string` | No |  |
| `cardStripePaymentIntentID` | `string` | No |  |
| `cardStripePaymentIntentSecret` | `string` | No |  |
| `cardTransmitRawDetails` | `boolean` | No |  |
| `createdByUser` | `table` | Yes |  |
| `currency` | `string` | No |  |
| `customFields` | `table` | No |  |
| `customerEmailAddress` | `string` | No |  |
| `customerID` | `string` | No |  |
| `customerName` | `string` | No |  |
| `description` | `string` | No |  |
| `destinationAccount` | `table` | No |  |
| `directDebitPayment` | `table` | No |  |
| `dueDate` | `string` | No |  |
| `events` | `table` | No |  |
| `failureCallbackUrl` | `string` | No |  |
| `fieldDisplaySettings` | `table` | No |  |
| `formattedAmount` | `string` | No |  |
| `hostedPayCheckoutUrl` | `string` | No |  |
| `id` | `string` | No |  |
| `ignoreAddressVerification` | `boolean` | No |  |
| `inserted` | `string` | No |  |
| `insertedSortable` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `jwk` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `lightningInvoice` | `string` | No |  |
| `lightningInvoiceExpiresAt` | `string` | No |  |
| `merchantDirectDebitMandateID` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantTokenDescription` | `string` | No |  |
| `notificationEmailAddresses` | `string` | No |  |
| `notificationRoleIDs` | `table` | No |  |
| `orderID` | `string` | No |  |
| `partialPaymentMethod` | `string` | No |  |
| `partialPaymentSteps` | `string` | No |  |
| `paymentAttempts` | `table` | No |  |
| `paymentMethods` | `table` | No |  |
| `paymentProcessor` | `string` | No |  |
| `payrunID` | `string` | No |  |
| `pispAccountID` | `string` | No |  |
| `priorityBankID` | `string` | No |  |
| `result` | `table` | No |  |
| `sandboxSettleDelayInSeconds` | `number` | No |  |
| `shippingAddress` | `table` | No |  |
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
| `tagIds` | `table` | No |  |
| `tags` | `table` | No |  |
| `title` | `string` | No |  |
| `tokenisedCards` | `table` | No |  |
| `transactions` | `table` | No |  |
| `useHostedPaymentPage` | `boolean` | No |  |

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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Payment():create({
  createdByUser = --[[ table ]],
})
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Payment():load({ id = "payment_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Payment():update({
  id = "payment_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PaymentAccountEntity

```lua
local payment_account = client:PaymentAccount(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `string` | No |  |
| `accountSupplierName` | `string` | No |  |
| `availableBalance` | `number` | No |  |
| `availableBalanceMinorUnits` | `number` | No |  |
| `balance` | `number` | No |  |
| `balanceMinorUnits` | `number` | No |  |
| `bankName` | `string` | No |  |
| `consentID` | `string` | No |  |
| `createdBy` | `table` | Yes |  |
| `createdByDisplayName` | `string` | No |  |
| `currency` | `string` | No |  |
| `defaultPaymentRail` | `string` | No |  |
| `displayName` | `string` | No |  |
| `expiryDate` | `string` | No |  |
| `externalAccountIcon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `table` | Yes |  |
| `inserted` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `isConnectedAccount` | `boolean` | No |  |
| `isDefault` | `boolean` | No |  |
| `isTrustAccount` | `boolean` | No |  |
| `isVirtual` | `boolean` | No |  |
| `lastTransaction` | `table` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `physicalAccountID` | `string` | No |  |
| `rules` | `table` | No |  |
| `submittedPayoutsBalance` | `number` | No |  |
| `submittedPayoutsBalanceMinorUnits` | `number` | No |  |
| `summary` | `string` | No |  |
| `supplierSepaInstantStatus` | `string` | No |  |
| `xeroBankFeedConnectionStatus` | `string` | No |  |
| `xeroBankFeedLastSyncedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `string` | No |  |
| `xeroBankFeedSyncStatus` | `string` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `number` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:PaymentAccount():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentAccountEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PaymentAccountMinimalEntity

```lua
local payment_account_minimal = client:PaymentAccountMinimal(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `string` | No |  |
| `availableBalance` | `number` | No |  |
| `balance` | `number` | No |  |
| `balanceMinorUnits` | `number` | No |  |
| `currency` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `table` | Yes |  |
| `isArchived` | `boolean` | No |  |
| `isConnectedAccount` | `boolean` | No |  |
| `merchantID` | `string` | No |  |
| `submittedPayoutsBalance` | `number` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:PaymentAccountMinimal():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentAccountMinimalEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PaymentInitiationEntity

```lua
local payment_initiation = client:PaymentInitiation(nil)
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PaymentInitiation():create({
  paymentrequest_id = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentInitiationEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PaymentRequestEntity

```lua
local payment_request = client:PaymentRequest(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresses` | `table` | No |  |
| `amount` | `number` | No |  |
| `amountPending` | `number` | No |  |
| `amountReceived` | `number` | No |  |
| `amountRefunded` | `number` | No |  |
| `autoSendReceipt` | `boolean` | No |  |
| `baseOriginUrl` | `string` | No |  |
| `callbackUrl` | `string` | No |  |
| `cardAuthorizeOnly` | `boolean` | No |  |
| `cardCreateToken` | `boolean` | No |  |
| `cardCreateTokenMode` | `string` | No |  |
| `cardIgnoreCVN` | `boolean` | No |  |
| `cardProcessorMerchantID` | `string` | No |  |
| `cardStripePaymentIntentID` | `string` | No |  |
| `cardStripePaymentIntentSecret` | `string` | No |  |
| `createdByUser` | `table` | Yes |  |
| `currency` | `string` | No |  |
| `customFields` | `table` | No |  |
| `customerEmailAddress` | `string` | No |  |
| `customerID` | `string` | No |  |
| `customerName` | `string` | No |  |
| `description` | `string` | No |  |
| `destinationAccount` | `table` | No |  |
| `directDebitPayment` | `table` | No |  |
| `doSimulateSettlementFailure` | `boolean` | No |  |
| `dueDate` | `string` | No |  |
| `errorDescription` | `string` | No |  |
| `events` | `table` | No |  |
| `failedPaymentRequests` | `table` | No |  |
| `failureCallbackUrl` | `string` | No |  |
| `fieldDisplaySettings` | `table` | No |  |
| `formattedAmount` | `string` | No |  |
| `hostedPayCheckoutUrl` | `string` | No |  |
| `id` | `string` | No |  |
| `ignoreAddressVerification` | `boolean` | No |  |
| `inserted` | `string` | No |  |
| `insertedSortable` | `string` | No |  |
| `institution` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `jwk` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `lightningInvoice` | `string` | No |  |
| `lightningInvoiceExpiresAt` | `string` | No |  |
| `merchantDirectDebitMandateID` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantTokenDescription` | `string` | No |  |
| `notificationEmailAddresses` | `string` | No |  |
| `notificationRoleIDs` | `table` | No |  |
| `orderID` | `string` | No |  |
| `partialPaymentMethod` | `string` | No |  |
| `partialPaymentSteps` | `string` | No |  |
| `paymentAttempts` | `table` | No |  |
| `paymentInitiationID` | `string` | No |  |
| `paymentMethods` | `table` | No |  |
| `paymentProcessor` | `string` | No |  |
| `paymentRequests` | `table` | No |  |
| `payrunID` | `string` | No |  |
| `pispAccountID` | `string` | No |  |
| `priorityBankID` | `string` | No |  |
| `result` | `table` | No |  |
| `sandboxSettleDelayInSeconds` | `number` | No |  |
| `shippingAddress` | `table` | No |  |
| `status` | `string` | No |  |
| `successWebHookUrl` | `string` | No |  |
| `tags` | `table` | No |  |
| `title` | `string` | No |  |
| `tokenisedCards` | `table` | No |  |
| `transactions` | `table` | No |  |
| `useHostedPaymentPage` | `boolean` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:PaymentRequest():create({
  createdByUser = --[[ table ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:PaymentRequest():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:PaymentRequest():load()
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:PaymentRequest():remove({ id = "id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:PaymentRequest():update({
  paymentrequest_id = "paymentrequest_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentRequestEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PaymentRequestEventEntity

```lua
local payment_request_event = client:PaymentRequestEvent(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `number` | Yes |  |
| `applePayTransactionID` | `string` | No |  |
| `cardAuthorizationResponseID` | `string` | No |  |
| `cardExpiryMonth` | `number` | No |  |
| `cardExpiryYear` | `number` | No |  |
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:PaymentRequestEvent():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentRequestEventEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PaymentRequestMetricEntity

```lua
local payment_request_metric = client:PaymentRequestMetric(nil)
```

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:PaymentRequestMetric():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentRequestMetricEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PaymentRequestMinimalEntity

```lua
local payment_request_minimal = client:PaymentRequestMinimal(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `number` | No |  |
| `amountPending` | `number` | No |  |
| `amountReceived` | `number` | No |  |
| `amountRefunded` | `number` | No |  |
| `callbackUrl` | `string` | No |  |
| `cardStripePaymentIntentSecret` | `string` | No |  |
| `countryCode` | `string` | No |  |
| `currency` | `string` | No |  |
| `customFieldsToDisplay` | `table` | No |  |
| `description` | `string` | No |  |
| `dueDate` | `string` | No |  |
| `fieldDisplaySettings` | `table` | No |  |
| `googlePayMerchantID` | `string` | No |  |
| `id` | `string` | No |  |
| `jwk` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantLogoUrlPng` | `string` | No |  |
| `merchantLogoUrlSvg` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `merchantShortName` | `string` | No |  |
| `partialPaymentMethod` | `string` | No |  |
| `paymentAttempts` | `table` | No |  |
| `paymentMethodsList` | `table` | No |  |
| `paymentProcessor` | `string` | No |  |
| `paymentProcessorKey` | `string` | No |  |
| `pispError` | `string` | No |  |
| `priorityBankID` | `string` | No |  |
| `status` | `string` | No |  |
| `stripeAccountID` | `string` | No |  |
| `title` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:PaymentRequestMinimal():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentRequestMinimalEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PaymentRequestResultEntity

```lua
local payment_request_result = client:PaymentRequestResult(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `number` | No |  |
| `amountPending` | `number` | No |  |
| `amountReceived` | `number` | No |  |
| `amountRefunded` | `number` | No |  |
| `currency` | `string` | No |  |
| `customerID` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |
| `payments` | `table` | No |  |
| `pispAuthorizations` | `table` | No |  |
| `requestedAmount` | `number` | No |  |
| `result` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:PaymentRequestResult():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentRequestResultEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PayoutEntity

```lua
local payout = client:Payout(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountID` | `string` | No |  |
| `allowIncomplete` | `boolean` | No |  |
| `amount` | `number` | No |  |
| `amountMinorUnits` | `number` | No |  |
| `approvePayoutUrl` | `string` | No |  |
| `approverID` | `string` | No |  |
| `authenticationMethods` | `table` | No |  |
| `authorisations` | `table` | No |  |
| `authorisersCompletedCount` | `number` | No |  |
| `authorisersRequiredCount` | `number` | No |  |
| `batchPayoutID` | `string` | No |  |
| `beneficiary` | `table` | Yes |  |
| `beneficiaryID` | `string` | No |  |
| `canAuthorise` | `boolean` | No |  |
| `canProcess` | `boolean` | No |  |
| `canUpdate` | `boolean` | No |  |
| `chargeBearer` | `string` | No |  |
| `createdBy` | `string` | No |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | No |  |
| `currentUserID` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `table` | No |  |
| `documents` | `table` | No |  |
| `events` | `table` | No |  |
| `failedPayouts` | `table` | No |  |
| `formattedAmount` | `string` | No |  |
| `formattedFxDestinationAmount` | `string` | No |  |
| `formattedSchedule` | `string` | No |  |
| `formattedScheduleDayOnly` | `string` | No |  |
| `formattedSourceAccountAvailableBalance` | `string` | No |  |
| `fxDestinationAmount` | `number` | No |  |
| `fxDestinationAmountMinorUnits` | `number` | No |  |
| `fxDestinationCurrency` | `string` | No |  |
| `fxQuoteExpiresAt` | `string` | No |  |
| `fxQuoteID` | `string` | No |  |
| `fxRate` | `number` | No |  |
| `fxUseDestinationAmount` | `boolean` | No |  |
| `hasCurrentUserAuthorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoiceID` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `isFailed` | `boolean` | No |  |
| `isSettled` | `boolean` | No |  |
| `isSubmitted` | `boolean` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantTokenDescription` | `string` | No |  |
| `nonce` | `string` | No |  |
| `paymentProcessor` | `string` | No |  |
| `paymentRail` | `string` | No |  |
| `payouts` | `table` | No |  |
| `payrunID` | `string` | No |  |
| `payrunName` | `string` | No |  |
| `reason` | `string` | No |  |
| `rule` | `table` | No |  |
| `scheduleDate` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `sourceAccountAvailableBalance` | `number` | No |  |
| `sourceAccountAvailableBalanceMinorUnits` | `number` | No |  |
| `sourceAccountBic` | `string` | No |  |
| `sourceAccountCurrency` | `string` | No |  |
| `sourceAccountIban` | `string` | No |  |
| `sourceAccountIdentifier` | `table` | Yes |  |
| `sourceAccountName` | `string` | No |  |
| `sourceAccountNumber` | `string` | No |  |
| `sourceAccountSortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tagIds` | `table` | No |  |
| `tags` | `table` | No |  |
| `theirReference` | `string` | No |  |
| `topupPayrunID` | `string` | No |  |
| `transactedAmount` | `number` | No |  |
| `transactedFxAmount` | `number` | No |  |
| `transactedFxRate` | `number` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Payout():create({
  beneficiary = --[[ table ]],
  sourceAccountIdentifier = --[[ table ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Payout():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Payout():load({ id = "payout_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Payout():remove({ id = "payout_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Payout():update({
  id = "payout_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PayoutEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PayoutKeysetPageEntity

```lua
local payout_keyset_page = client:PayoutKeysetPage(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountID` | `string` | No |  |
| `amount` | `number` | No |  |
| `amountMinorUnits` | `number` | No |  |
| `approvePayoutUrl` | `string` | No |  |
| `approverID` | `string` | No |  |
| `authenticationMethods` | `table` | No |  |
| `authorisations` | `table` | No |  |
| `authorisersCompletedCount` | `number` | No |  |
| `authorisersRequiredCount` | `number` | No |  |
| `batchPayoutID` | `string` | No |  |
| `beneficiary` | `table` | Yes |  |
| `canAuthorise` | `boolean` | No |  |
| `canProcess` | `boolean` | No |  |
| `canUpdate` | `boolean` | No |  |
| `chargeBearer` | `string` | No |  |
| `createdBy` | `string` | No |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | No |  |
| `currentUserID` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `table` | No |  |
| `documents` | `table` | No |  |
| `events` | `table` | No |  |
| `formattedAmount` | `string` | No |  |
| `formattedFxDestinationAmount` | `string` | No |  |
| `formattedSchedule` | `string` | No |  |
| `formattedScheduleDayOnly` | `string` | No |  |
| `formattedSourceAccountAvailableBalance` | `string` | No |  |
| `fxDestinationAmount` | `number` | No |  |
| `fxDestinationAmountMinorUnits` | `number` | No |  |
| `fxDestinationCurrency` | `string` | No |  |
| `fxQuoteExpiresAt` | `string` | No |  |
| `fxQuoteID` | `string` | No |  |
| `fxRate` | `number` | No |  |
| `fxUseDestinationAmount` | `boolean` | No |  |
| `hasCurrentUserAuthorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoiceID` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `isFailed` | `boolean` | No |  |
| `isSettled` | `boolean` | No |  |
| `isSubmitted` | `boolean` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantTokenDescription` | `string` | No |  |
| `nonce` | `string` | No |  |
| `paymentProcessor` | `string` | No |  |
| `paymentRail` | `string` | No |  |
| `payrunID` | `string` | No |  |
| `payrunName` | `string` | No |  |
| `rule` | `table` | No |  |
| `scheduleDate` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `sourceAccountAvailableBalance` | `number` | No |  |
| `sourceAccountAvailableBalanceMinorUnits` | `number` | No |  |
| `sourceAccountBic` | `string` | No |  |
| `sourceAccountCurrency` | `string` | No |  |
| `sourceAccountIban` | `string` | No |  |
| `sourceAccountIdentifier` | `table` | Yes |  |
| `sourceAccountName` | `string` | No |  |
| `sourceAccountNumber` | `string` | No |  |
| `sourceAccountSortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tags` | `table` | No |  |
| `theirReference` | `string` | No |  |
| `topupPayrunID` | `string` | No |  |
| `transactedAmount` | `number` | No |  |
| `transactedFxAmount` | `number` | No |  |
| `transactedFxRate` | `number` | No |  |
| `type` | `string` | No |  |
| `userID` | `string` | No |  |
| `yourReference` | `string` | No |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:PayoutKeysetPage():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PayoutKeysetPageEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PayoutMetricEntity

```lua
local payout_metric = client:PayoutMetric(nil)
```

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:PayoutMetric():load()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PayoutMetricEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PayrunEntity

```lua
local payrun = client:Payrun(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisationDate` | `string` | No |  |
| `authorisations` | `table` | No |  |
| `authorisersCompletedCount` | `number` | No |  |
| `authorisersRequiredCount` | `number` | No |  |
| `batchPayoutID` | `string` | No |  |
| `canAuthorise` | `boolean` | No |  |
| `canDelete` | `boolean` | No |  |
| `canEdit` | `boolean` | No |  |
| `events` | `table` | No |  |
| `hasCurrentUserAuthorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoices` | `table` | No |  |
| `invoicesMinimal` | `table` | No |  |
| `isArchived` | `boolean` | No |  |
| `lastUpdated` | `string` | No |  |
| `lastUpdatedBy` | `table` | Yes |  |
| `merchantID` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | No |  |
| `notes` | `string` | No |  |
| `payments` | `table` | No |  |
| `payouts` | `table` | No |  |
| `payoutsCount` | `number` | No |  |
| `reason` | `string` | No |  |
| `scheduleDate` | `string` | No |  |
| `scheduledDate` | `string` | No |  |
| `sourceAccounts` | `table` | No |  |
| `status` | `string` | No |  |
| `totalEur` | `number` | No |  |
| `totalGbp` | `number` | No |  |
| `totalUsd` | `number` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Payrun():create({
  id = --[[ string ]],
  lastUpdatedBy = --[[ table ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Payrun():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Payrun():load({ id = "payrun_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Payrun():remove({ id = "payrun_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Payrun():update({
  id = "payrun_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PayrunEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ReportEntity

```lua
local report = client:Report(nil)
```

### Operations

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Report():update({
  id = "id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ReportEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## ReportResultEntity

```lua
local report_result = client:ReportResult(nil)
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
| `statementNumber` | `number` | No |  |

### Operations

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:ReportResult():load({ id = 1, report_id = "report_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ReportResultEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## RoleEntity

```lua
local role = client:Role(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failedRoles` | `table` | No |  |
| `roles` | `table` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Role():create({
  merchant_id = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RoleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## RuleEntity

```lua
local rule = client:Rule(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account` | `table` | No |  |
| `accountID` | `string` | No |  |
| `approveUrl` | `string` | No |  |
| `approverID` | `string` | No |  |
| `authenticationMethods` | `table` | No |  |
| `authorisations` | `table` | No |  |
| `authorisersCompletedCount` | `number` | No |  |
| `authorisersRequiredCount` | `number` | No |  |
| `canAuthorise` | `boolean` | No |  |
| `createdBy` | `table` | Yes |  |
| `description` | `string` | No |  |
| `endAt` | `string` | No |  |
| `hasCurrentUserAuthorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isDisabled` | `boolean` | No |  |
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
| `sweepAction` | `table` | No |  |
| `timeZoneId` | `string` | No |  |
| `triggerCronExpression` | `string` | No |  |
| `triggerOnPayIn` | `boolean` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Rule():create({
  createdBy = --[[ table ]],
  nonce = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Rule():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Rule():load({ id = "rule_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Rule():remove({ id = "rule_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Rule():update({
  id = "rule_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RuleEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## RuleEventEntity

```lua
local rule_event = client:RuleEvent(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errorMessage` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isAuthoriseToEnable` | `boolean` | No |  |
| `message` | `string` | No |  |
| `rawResponse` | `string` | No |  |
| `ruleEventType` | `string` | No |  |
| `ruleID` | `string` | No |  |
| `user` | `table` | Yes |  |

### Operations

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:RuleEvent():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RuleEventEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## TagEntity

```lua
local tag = client:Tag(nil)
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Tag():create({
  merchant_id = --[[ string ]],
  merchantID = --[[ string ]],
  name = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Tag():list()
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TagEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## TokenEntity

```lua
local token = client:Token(nil)
```

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Token():create({
  id = --[[ string ]],
})
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Token():remove({ id = "id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TokenEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## TransactionEntity

```lua
local transaction = client:Transaction(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountID` | `string` | No |  |
| `accountName` | `string` | No |  |
| `accountSequenceNumber` | `number` | No |  |
| `addressDetails` | `table` | No |  |
| `amount` | `number` | No |  |
| `amountMinorUnits` | `number` | No |  |
| `balance` | `number` | No |  |
| `balanceMinorUnits` | `number` | No |  |
| `bookingDateTime` | `string` | No |  |
| `chargeDetails` | `table` | No |  |
| `content` | `table` | No |  |
| `counterparty` | `table` | No |  |
| `counterpartySummary` | `string` | No |  |
| `currency` | `string` | No |  |
| `currencyExchange` | `table` | No |  |
| `date` | `string` | No |  |
| `description` | `string` | No |  |
| `enrichment` | `table` | No |  |
| `fxAmount` | `number` | No |  |
| `fxCurrency` | `string` | No |  |
| `fxRate` | `number` | No |  |
| `grossAmount` | `table` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isoBankTransactionCode` | `table` | No |  |
| `merchant` | `table` | No |  |
| `merchantID` | `string` | No |  |
| `pageNumber` | `number` | No |  |
| `pageSize` | `number` | No |  |
| `payeeDetails` | `table` | Yes |  |
| `payerDetails` | `table` | Yes |  |
| `paymentRequestCustomFields` | `table` | No |  |
| `paymentRequestID` | `string` | No |  |
| `payoutID` | `string` | No |  |
| `proprietaryBankTransactionCode` | `table` | No |  |
| `rawReference` | `string` | No |  |
| `reference` | `string` | No |  |
| `ruleID` | `string` | No |  |
| `statementReferences` | `table` | No |  |
| `status` | `string` | No |  |
| `supplementaryData` | `any` | No |  |
| `tags` | `table` | No |  |
| `theirReference` | `string` | No |  |
| `totalPages` | `number` | No |  |
| `totalSize` | `number` | No |  |
| `transactionAmount` | `table` | Yes |  |
| `transactionDate` | `string` | No |  |
| `transactionInformation` | `table` | No |  |
| `transactionMutability` | `string` | No |  |
| `type` | `string` | No |  |
| `valueDateTime` | `string` | No |  |
| `virtualIBAN` | `string` | No |  |
| `yourReference` | `string` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Transaction():create({
  id = --[[ string ]],
  grossAmount = --[[ table ]],
  payeeDetails = --[[ table ]],
  payerDetails = --[[ table ]],
  transactionAmount = --[[ table ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Transaction():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Transaction():load({ id = "transaction_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Transaction():remove({ id = "transaction_id" })
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TransactionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## UserEntity

```lua
local user = client:User(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `clientSessionTimeouts` | `table` | No |  |
| `emailAddress` | `string` | Yes |  |
| `firstName` | `string` | Yes |  |
| `id` | `string` | No |  |
| `lastName` | `string` | Yes |  |
| `passkeyAdded` | `boolean` | No |  |
| `permissions` | `table` | No |  |
| `profile` | `string` | No |  |
| `rolesWithScope` | `table` | No |  |
| `twoFactorEnabled` | `boolean` | No |  |
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

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:User():list()
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:User():update({
  id = "id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UserEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## UserInviteEntity

```lua
local user_invite = client:UserInvite(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisationStatus` | `table` | No |  |
| `failedUserInvites` | `table` | No |  |
| `id` | `string` | No |  |
| `initialRoleID` | `string` | No |  |
| `inviteeEmailAddress` | `string` | No |  |
| `inviteeFirstName` | `string` | No |  |
| `inviteeLastName` | `string` | No |  |
| `inviterEmailAddress` | `string` | No |  |
| `inviterFirstName` | `string` | No |  |
| `inviterLastName` | `string` | No |  |
| `isAuthorised` | `boolean` | No |  |
| `isInviteeRegistered` | `boolean` | No |  |
| `lastInvited` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `message` | `string` | No |  |
| `registrationUrl` | `string` | No |  |
| `sendInviteEmail` | `boolean` | No |  |
| `status` | `string` | No |  |
| `user` | `table` | Yes |  |
| `userID` | `string` | No |  |
| `userInvites` | `table` | No |  |

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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:UserInvite():create({
  user = --[[ table ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:UserInvite():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:UserInvite():load({ id = "user_invite_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:UserInvite():remove({ id = "user_invite_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:UserInvite():update({
  id = "user_invite_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UserInviteEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## VirtualEntity

```lua
local virtual = client:Virtual(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `string` | No |  |
| `accountSupplierName` | `string` | No |  |
| `availableBalance` | `number` | No |  |
| `availableBalanceMinorUnits` | `number` | No |  |
| `balance` | `number` | No |  |
| `balanceMinorUnits` | `number` | No |  |
| `bankName` | `string` | No |  |
| `consentID` | `string` | No |  |
| `createdBy` | `table` | Yes |  |
| `createdByDisplayName` | `string` | No |  |
| `currency` | `string` | No |  |
| `defaultPaymentRail` | `string` | No |  |
| `displayName` | `string` | No |  |
| `expiryDate` | `string` | No |  |
| `externalAccountIcon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `table` | Yes |  |
| `inserted` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `isConnectedAccount` | `boolean` | No |  |
| `isDefault` | `boolean` | No |  |
| `isTrustAccount` | `boolean` | No |  |
| `isVirtual` | `boolean` | No |  |
| `lastTransaction` | `table` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `name` | `string` | Yes |  |
| `physicalAccountID` | `string` | No |  |
| `rules` | `table` | No |  |
| `submittedPayoutsBalance` | `number` | No |  |
| `submittedPayoutsBalanceMinorUnits` | `number` | No |  |
| `summary` | `string` | No |  |
| `supplierSepaInstantStatus` | `string` | No |  |
| `xeroBankFeedConnectionStatus` | `string` | No |  |
| `xeroBankFeedLastSyncedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `string` | No |  |
| `xeroBankFeedSyncStatus` | `string` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `number` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Virtual():create({
  account_id = --[[ string ]],
  createdBy = --[[ table ]],
  identifier = --[[ table ]],
  name = --[[ string ]],
})
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Virtual():update({
  account_id = "account_id",
  id = "id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VirtualEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## WebhookEntity

```lua
local webhook = client:Webhook(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destinationUrl` | `string` | No |  |
| `emailAddress` | `string` | No |  |
| `failedNotificationEmailAddress` | `string` | No |  |
| `id` | `string` | No |  |
| `isActive` | `boolean` | No |  |
| `merchantID` | `string` | No |  |
| `notificationMethod` | `string` | No |  |
| `resourceTypes` | `table` | No |  |
| `retry` | `boolean` | No |  |
| `secret` | `string` | No |  |
| `version` | `number` | No |  |

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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Webhook():create({
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Webhook():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Webhook():load({ id = "webhook_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Webhook():remove({ id = "webhook_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Webhook():update({
  id = "webhook_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WebhookEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
  },
})
```

