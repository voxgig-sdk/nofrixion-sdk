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
| `accountBalances` | `Array` | No |  |
| `accountID` | `String` | No |  |
| `accountIdentifications` | `Array` | No |  |
| `accountName` | `String` | No |  |
| `accountNames` | `Array` | No |  |
| `accountSupplierName` | `String` | No |  |
| `accountType` | `String` | No |  |
| `availableBalance` | `Float` | No |  |
| `availableBalanceMinorUnits` | `Integer` | No |  |
| `balance` | `Float` | No |  |
| `balanceMinorUnits` | `Integer` | No |  |
| `bankName` | `String` | No |  |
| `consentID` | `String` | No |  |
| `consolidatedAccountInformation` | `Hash` | No |  |
| `createdBy` | `Hash` | Yes |  |
| `createdByDisplayName` | `String` | No |  |
| `currency` | `String` | No |  |
| `defaultPaymentRail` | `String` | No |  |
| `description` | `String` | No |  |
| `details` | `String` | No |  |
| `displayName` | `String` | No |  |
| `expiryDate` | `String` | No |  |
| `externalAccountIcon` | `String` | No |  |
| `format` | `String` | No |  |
| `fromDate` | `String` | No |  |
| `id` | `String` | No |  |
| `identifier` | `Hash` | Yes |  |
| `inserted` | `String` | No |  |
| `isArchived` | `Boolean` | No |  |
| `isConnectedAccount` | `Boolean` | No |  |
| `isDefault` | `Boolean` | No |  |
| `isTrustAccount` | `Boolean` | No |  |
| `isVirtual` | `Boolean` | No |  |
| `lastTransaction` | `Hash` | No |  |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `merchantName` | `String` | No |  |
| `nickname` | `String` | No |  |
| `physicalAccountID` | `String` | No |  |
| `roleIDs` | `Array` | No |  |
| `rules` | `Array` | No |  |
| `submittedPayoutsBalance` | `Float` | No |  |
| `submittedPayoutsBalanceMinorUnits` | `Integer` | No |  |
| `summary` | `String` | No |  |
| `supplierPhysicalAccountID` | `String` | No |  |
| `supplierSepaInstantStatus` | `String` | No |  |
| `toDate` | `String` | No |  |
| `type` | `String` | No |  |
| `usageType` | `String` | No |  |
| `xeroBankFeedConnectionStatus` | `String` | No |  |
| `xeroBankFeedLastSyncedAt` | `String` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `String` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `String` | No |  |
| `xeroBankFeedSyncStatus` | `String` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `Integer` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Account.create({
  "createdBy" => {}, # Hash
  "identifier" => {}, # Hash
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
| `approveUrl` | `String` | No |  |
| `id` | `String` | No |  |
| `payouts` | `Array` | No |  |

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
| `approvalCallbackUrl` | `String` | No |  |
| `authenticationMethods` | `Array` | No |  |
| `authorisations` | `Array` | No |  |
| `authorisersCompletedCount` | `Integer` | No |  |
| `authorisersRequiredCount` | `Integer` | No |  |
| `beneficiaries` | `Array` | No |  |
| `beneficiaryEvents` | `Array` | No |  |
| `canAuthorise` | `Boolean` | No |  |
| `canUpdate` | `Boolean` | No |  |
| `createdBy` | `Hash` | Yes |  |
| `createdByEmailAddress` | `String` | No |  |
| `currency` | `String` | Yes |  |
| `destination` | `Hash` | No |  |
| `failedBeneficiaries` | `Hash` | No |  |
| `hasCurrentUserAuthorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `isEnabled` | `Boolean` | No |  |
| `lastAuthorised` | `String` | No |  |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `name` | `String` | Yes |  |
| `nonce` | `String` | No |  |
| `sourceAccountIDs` | `Array` | No |  |
| `sourceAccounts` | `Array` | No |  |
| `theirReference` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Beneficiary.create({
  "createdBy" => {}, # Hash
  "currency" => "example_currency", # String
  "name" => "example_name", # String
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
| `groupMembers` | `Array` | No |  |
| `groupName` | `String` | Yes |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | Yes |  |

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
| `authorizedAmount` | `String` | No |  |
| `currencyCode` | `String` | No |  |
| `isPayerAuthenticationRequired` | `Boolean` | No |  |
| `isSoftDecline` | `Boolean` | No |  |
| `payerAuthenticationAccessToken` | `String` | No |  |
| `payerAuthenticationMerchantData` | `String` | No |  |
| `payerAuthenticationUrl` | `String` | No |  |
| `payerAuthenticationWindowHeight` | `Integer` | No |  |
| `payerAuthenticationWindowWidth` | `Integer` | No |  |
| `paymentRequestCallbackUrl` | `String` | No |  |
| `paymentRequestID` | `String` | No |  |
| `requestID` | `String` | No |  |
| `responseCode` | `String` | No |  |
| `responseType` | `String` | No |  |
| `status` | `String` | No |  |
| `threeDSRedirectUrl` | `String` | No |  |
| `transactionID` | `String` | No |  |

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
| `cardType` | `String` | No |  |
| `customerEmailAddress` | `String` | No |  |
| `expiryMonth` | `String` | No |  |
| `expiryYear` | `String` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `lastFourDigits` | `String` | No |  |
| `lastUpdated` | `String` | No |  |
| `maskedCardNumber` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `paymentRequestID` | `String` | No |  |

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
| `authorizedAmount` | `String` | No |  |
| `currencyCode` | `String` | No |  |
| `isPayerAuthenticationRequired` | `Boolean` | No |  |
| `isSoftDecline` | `Boolean` | No |  |
| `payerAuthenticationAccessToken` | `String` | No |  |
| `payerAuthenticationMerchantData` | `String` | No |  |
| `payerAuthenticationUrl` | `String` | No |  |
| `payerAuthenticationWindowHeight` | `Integer` | No |  |
| `payerAuthenticationWindowWidth` | `Integer` | No |  |
| `paymentRequestCallbackUrl` | `String` | No |  |
| `paymentRequestID` | `String` | No |  |
| `requestID` | `String` | No |  |
| `responseCode` | `String` | No |  |
| `responseType` | `String` | No |  |
| `status` | `String` | No |  |
| `threeDSRedirectUrl` | `String` | No |  |
| `transactionID` | `String` | No |  |

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
| `authorisationUrl` | `String` | No |  |
| `callbackUrl` | `String` | No |  |
| `consentID` | `String` | No |  |
| `emailAddress` | `String` | No |  |
| `expiryDate` | `String` | No |  |
| `failureCallbackUrl` | `String` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `institutionID` | `String` | No |  |
| `isConnectedAccounts` | `Boolean` | No |  |
| `isEnabled` | `Boolean` | No |  |
| `merchantID` | `String` | No |  |
| `provider` | `String` | No |  |
| `successWebHookUrl` | `String` | No |  |

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
| `decimals` | `Integer` | No |  |
| `isFiat` | `Boolean` | No |  |
| `iso4217AlphaCode` | `String` | No |  |
| `iso4217NumericCode` | `String` | No |  |
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
| `failedSubmissions` | `Hash` | No |  |
| `successfulSubmissions` | `Array` | No |  |

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
| `destinationCurrency` | `String` | No |  |
| `exchangeRate` | `Float` | No |  |
| `expiryTime` | `String` | No |  |
| `quoteID` | `String` | No |  |
| `sourceCurrency` | `String` | No |  |

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
| `paymentRequestID` | `String` | No |  |
| `responseType` | `String` | No |  |

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
| `accountNumber` | `String` | No |  |
| `addressLine1` | `String` | Yes |  |
| `addressLine2` | `String` | No |  |
| `approvedAt` | `String` | No |  |
| `city` | `String` | Yes |  |
| `countryCode` | `String` | Yes |  |
| `currency` | `String` | No |  |
| `customerAccountNumber` | `String` | No |  |
| `customerCity` | `String` | No |  |
| `customerCountryCode` | `String` | No |  |
| `customerCountryName` | `String` | No |  |
| `customerEmailAddress` | `String` | No |  |
| `customerFirstName` | `String` | No |  |
| `customerIban` | `String` | No |  |
| `customerLastName` | `String` | No |  |
| `customerSortCode` | `String` | No |  |
| `emailAddress` | `String` | Yes |  |
| `firstName` | `String` | Yes |  |
| `iban` | `String` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `isRecurring` | `Boolean` | No |  |
| `lastName` | `String` | Yes |  |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `postalCode` | `String` | Yes |  |
| `reference` | `String` | No |  |
| `sortCode` | `String` | No |  |
| `status` | `String` | No |  |
| `supplierBankAccountID` | `String` | No |  |
| `supplierCustomerID` | `String` | No |  |
| `supplierMandateID` | `String` | No |  |
| `supplierName` | `String` | No |  |
| `supplierStatus` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Mandate.create({
  "addressLine1" => "example_addressLine1", # String
  "city" => "example_city", # String
  "countryCode" => "example_countryCode", # String
  "emailAddress" => "example_emailAddress", # String
  "firstName" => "example_firstName", # String
  "lastName" => "example_lastName", # String
  "postalCode" => "example_postalCode", # String
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
| `accountCurrencies` | `Array` | No |  |
| `canHaveTrustAccounts` | `Boolean` | No |  |
| `cardPaymentProcessor` | `String` | No |  |
| `companyID` | `String` | No |  |
| `displayQrOnHostedPay` | `Boolean` | No |  |
| `hostedPayVersion` | `Integer` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `isBlocked` | `Boolean` | No |  |
| `isExited` | `Boolean` | No |  |
| `isSuspended` | `Boolean` | No |  |
| `jurisdiction` | `String` | No |  |
| `logoUrlPng` | `String` | No |  |
| `logoUrlSvg` | `String` | No |  |
| `merchantCategoryCode` | `String` | No |  |
| `name` | `String` | No |  |
| `notes` | `String` | No |  |
| `parentMerchant` | `Hash` | No |  |
| `paymentAccountLimit` | `Integer` | No |  |
| `paymentAccounts` | `Array` | No |  |
| `reason` | `String` | No |  |
| `shortName` | `String` | No |  |
| `supportedPaymentMethodsList` | `Array` | No |  |
| `suspensionReason` | `String` | No |  |
| `tags` | `Array` | No |  |
| `timeZoneId` | `String` | No |  |
| `tradingName` | `String` | No |  |
| `webHookLimit` | `Integer` | No |  |
| `yourRoleName` | `String` | No |  |

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
| `amountLower` | `Float` | No |  |
| `amountUpper` | `Float` | No |  |
| `authorisationType` | `String` | No |  |
| `beneficiariesOnly` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `lastEditorCantAuthorise` | `Boolean` | No |  |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `numberOfAuthorisers` | `Integer` | No |  |
| `roleSettings` | `Array` | No |  |

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
| `approvedAt` | `String` | No |  |
| `currency` | `String` | No |  |
| `customerAccountNumber` | `String` | No |  |
| `customerCity` | `String` | No |  |
| `customerCountryCode` | `String` | No |  |
| `customerCountryName` | `String` | No |  |
| `customerEmailAddress` | `String` | No |  |
| `customerFirstName` | `String` | No |  |
| `customerIban` | `String` | No |  |
| `customerLastName` | `String` | No |  |
| `customerSortCode` | `String` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `isRecurring` | `Boolean` | No |  |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `reference` | `String` | No |  |
| `status` | `String` | No |  |
| `supplierBankAccountID` | `String` | No |  |
| `supplierCustomerID` | `String` | No |  |
| `supplierMandateID` | `String` | No |  |
| `supplierName` | `String` | No |  |
| `supplierStatus` | `String` | No |  |

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
| `bankCountryCodes` | `Array` | No |  |
| `bankID` | `String` | No |  |
| `bankName` | `String` | No |  |
| `businessInstitutionID` | `String` | No |  |
| `currency` | `String` | No |  |
| `logo` | `String` | No |  |
| `message` | `String` | No |  |
| `messageImageUrl` | `String` | No |  |
| `order` | `Integer` | No |  |
| `personalInstitutionID` | `String` | No |  |
| `processor` | `String` | No |  |
| `warningHeading` | `String` | No |  |
| `warningMessage` | `String` | No |  |

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
| `bankPaymentOptions` | `Hash` | No |  |
| `cardPaymentAddressOptions` | `Hash` | No |  |
| `cardPaymentCaptureOptions` | `Hash` | No |  |
| `customFields` | `Array` | No |  |
| `defaultFields` | `Array` | No |  |
| `description` | `String` | Yes |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `name` | `String` | Yes |  |
| `notificationOptions` | `Hash` | No |  |
| `paymentMethods` | `Hash` | No |  |
| `paymentTerms` | `Hash` | No |  |
| `priorityBankOptions` | `Hash` | No |  |
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
| `authenticationMethods` | `Array` | No |  |
| `authorisations` | `Array` | No |  |
| `authorisersCompletedCount` | `Integer` | No |  |
| `authorisersRequiredCount` | `Integer` | No |  |
| `canAuthorise` | `Boolean` | No |  |
| `description` | `String` | No |  |
| `expiresAt` | `String` | No |  |
| `hasCurrentUserAuthorised` | `Boolean` | No |  |
| `hmacAlgorithm` | `String` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `ipAddressWhitelist` | `String` | No |  |
| `isArchived` | `Boolean` | No |  |
| `isEnabled` | `Boolean` | No |  |
| `lastAuthorised` | `String` | No |  |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `nonce` | `String` | Yes |  |
| `permissionTypes` | `Array` | No |  |
| `requestSignatureVersion` | `Integer` | No |  |
| `sharedSecretAlgorithm` | `String` | No |  |
| `sharedSecretBase64` | `String` | No |  |
| `token` | `String` | No |  |

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
| `buildVersion` | `Integer` | No |  |
| `majorVersion` | `Integer` | No |  |
| `minorVersion` | `Integer` | No |  |
| `releaseName` | `String` | No |  |

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
| `accountName` | `String` | Yes |  |
| `accountNumber` | `String` | No |  |
| `iban` | `String` | Yes |  |
| `payeeVerifiedAccountName` | `String` | No |  |
| `result` | `String` | No |  |
| `secondaryIdentification` | `String` | No |  |
| `sortCode` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Payeeverification.create({
  "accountName" => "example_accountName", # String
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
| `addresses` | `Array` | No |  |
| `amount` | `Float` | No |  |
| `amountPending` | `Float` | No |  |
| `amountReceived` | `Float` | No |  |
| `amountRefunded` | `Float` | No |  |
| `autoSendReceipt` | `Boolean` | No |  |
| `baseOriginUrl` | `String` | No |  |
| `callbackUrl` | `String` | No |  |
| `cardAuthorizeOnly` | `Boolean` | No |  |
| `cardCreateToken` | `Boolean` | No |  |
| `cardCreateTokenMode` | `String` | No |  |
| `cardIgnoreCVN` | `Boolean` | No |  |
| `cardNoPayerAuthentication` | `Boolean` | No |  |
| `cardProcessorMerchantID` | `String` | No |  |
| `cardStripePaymentIntentID` | `String` | No |  |
| `cardStripePaymentIntentSecret` | `String` | No |  |
| `cardTransmitRawDetails` | `Boolean` | No |  |
| `createdByUser` | `Hash` | Yes |  |
| `currency` | `String` | No |  |
| `customFields` | `Array` | No |  |
| `customerEmailAddress` | `String` | No |  |
| `customerID` | `String` | No |  |
| `customerName` | `String` | No |  |
| `description` | `String` | No |  |
| `destinationAccount` | `Hash` | No |  |
| `directDebitPayment` | `Hash` | No |  |
| `dueDate` | `String` | No |  |
| `events` | `Array` | No |  |
| `failureCallbackUrl` | `String` | No |  |
| `fieldDisplaySettings` | `Array` | No |  |
| `formattedAmount` | `String` | No |  |
| `hostedPayCheckoutUrl` | `String` | No |  |
| `id` | `String` | No |  |
| `ignoreAddressVerification` | `Boolean` | No |  |
| `inserted` | `String` | No |  |
| `insertedSortable` | `String` | No |  |
| `isArchived` | `Boolean` | No |  |
| `jwk` | `String` | No |  |
| `lastUpdated` | `String` | No |  |
| `lightningInvoice` | `String` | No |  |
| `lightningInvoiceExpiresAt` | `String` | No |  |
| `merchantDirectDebitMandateID` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `merchantTokenDescription` | `String` | No |  |
| `notificationEmailAddresses` | `String` | No |  |
| `notificationRoleIDs` | `Array` | No |  |
| `orderID` | `String` | No |  |
| `partialPaymentMethod` | `String` | No |  |
| `partialPaymentSteps` | `String` | No |  |
| `paymentAttempts` | `Array` | No |  |
| `paymentMethods` | `Array` | No |  |
| `paymentProcessor` | `String` | No |  |
| `payrunID` | `String` | No |  |
| `pispAccountID` | `String` | No |  |
| `priorityBankID` | `String` | No |  |
| `result` | `Hash` | No |  |
| `sandboxSettleDelayInSeconds` | `Integer` | No |  |
| `shippingAddress` | `Hash` | No |  |
| `shippingAddressCity` | `String` | No |  |
| `shippingAddressCountryCode` | `String` | No |  |
| `shippingAddressCounty` | `String` | No |  |
| `shippingAddressLine1` | `String` | No |  |
| `shippingAddressLine2` | `String` | No |  |
| `shippingAddressPostCode` | `String` | No |  |
| `shippingEmail` | `String` | No |  |
| `shippingFirstName` | `String` | No |  |
| `shippingLastName` | `String` | No |  |
| `shippingPhone` | `String` | No |  |
| `status` | `String` | No |  |
| `successWebHookUrl` | `String` | No |  |
| `tagIds` | `Array` | No |  |
| `tags` | `Array` | No |  |
| `title` | `String` | No |  |
| `tokenisedCards` | `Array` | No |  |
| `transactions` | `Array` | No |  |
| `useHostedPaymentPage` | `Boolean` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Payment.create({
  "createdByUser" => {}, # Hash
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
| `accountName` | `String` | No |  |
| `accountSupplierName` | `String` | No |  |
| `availableBalance` | `Float` | No |  |
| `availableBalanceMinorUnits` | `Integer` | No |  |
| `balance` | `Float` | No |  |
| `balanceMinorUnits` | `Integer` | No |  |
| `bankName` | `String` | No |  |
| `consentID` | `String` | No |  |
| `createdBy` | `Hash` | Yes |  |
| `createdByDisplayName` | `String` | No |  |
| `currency` | `String` | No |  |
| `defaultPaymentRail` | `String` | No |  |
| `displayName` | `String` | No |  |
| `expiryDate` | `String` | No |  |
| `externalAccountIcon` | `String` | No |  |
| `id` | `String` | No |  |
| `identifier` | `Hash` | Yes |  |
| `inserted` | `String` | No |  |
| `isArchived` | `Boolean` | No |  |
| `isConnectedAccount` | `Boolean` | No |  |
| `isDefault` | `Boolean` | No |  |
| `isTrustAccount` | `Boolean` | No |  |
| `isVirtual` | `Boolean` | No |  |
| `lastTransaction` | `Hash` | No |  |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `merchantName` | `String` | No |  |
| `physicalAccountID` | `String` | No |  |
| `rules` | `Array` | No |  |
| `submittedPayoutsBalance` | `Float` | No |  |
| `submittedPayoutsBalanceMinorUnits` | `Integer` | No |  |
| `summary` | `String` | No |  |
| `supplierSepaInstantStatus` | `String` | No |  |
| `xeroBankFeedConnectionStatus` | `String` | No |  |
| `xeroBankFeedLastSyncedAt` | `String` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `String` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `String` | No |  |
| `xeroBankFeedSyncStatus` | `String` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `Integer` | No |  |

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
| `accountName` | `String` | No |  |
| `availableBalance` | `Float` | No |  |
| `balance` | `Float` | No |  |
| `balanceMinorUnits` | `Integer` | No |  |
| `currency` | `String` | No |  |
| `id` | `String` | No |  |
| `identifier` | `Hash` | Yes |  |
| `isArchived` | `Boolean` | No |  |
| `isConnectedAccount` | `Boolean` | No |  |
| `merchantID` | `String` | No |  |
| `submittedPayoutsBalance` | `Float` | No |  |

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
| `paymentInitiationID` | `String` | No |  |
| `paymentRequestCallbackUrl` | `String` | No |  |
| `paymentRequestID` | `String` | No |  |
| `redirectUrl` | `String` | No |  |
| `responseType` | `String` | No |  |
| `specificErrorMessage` | `String` | No |  |

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
| `addresses` | `Array` | No |  |
| `amount` | `Float` | No |  |
| `amountPending` | `Float` | No |  |
| `amountReceived` | `Float` | No |  |
| `amountRefunded` | `Float` | No |  |
| `autoSendReceipt` | `Boolean` | No |  |
| `baseOriginUrl` | `String` | No |  |
| `callbackUrl` | `String` | No |  |
| `cardAuthorizeOnly` | `Boolean` | No |  |
| `cardCreateToken` | `Boolean` | No |  |
| `cardCreateTokenMode` | `String` | No |  |
| `cardIgnoreCVN` | `Boolean` | No |  |
| `cardProcessorMerchantID` | `String` | No |  |
| `cardStripePaymentIntentID` | `String` | No |  |
| `cardStripePaymentIntentSecret` | `String` | No |  |
| `createdByUser` | `Hash` | Yes |  |
| `currency` | `String` | No |  |
| `customFields` | `Array` | No |  |
| `customerEmailAddress` | `String` | No |  |
| `customerID` | `String` | No |  |
| `customerName` | `String` | No |  |
| `description` | `String` | No |  |
| `destinationAccount` | `Hash` | No |  |
| `directDebitPayment` | `Hash` | No |  |
| `doSimulateSettlementFailure` | `Boolean` | No |  |
| `dueDate` | `String` | No |  |
| `errorDescription` | `String` | No |  |
| `events` | `Array` | No |  |
| `failedPaymentRequests` | `Hash` | No |  |
| `failureCallbackUrl` | `String` | No |  |
| `fieldDisplaySettings` | `Array` | No |  |
| `formattedAmount` | `String` | No |  |
| `hostedPayCheckoutUrl` | `String` | No |  |
| `id` | `String` | No |  |
| `ignoreAddressVerification` | `Boolean` | No |  |
| `inserted` | `String` | No |  |
| `insertedSortable` | `String` | No |  |
| `institution` | `String` | No |  |
| `isArchived` | `Boolean` | No |  |
| `jwk` | `String` | No |  |
| `lastUpdated` | `String` | No |  |
| `lightningInvoice` | `String` | No |  |
| `lightningInvoiceExpiresAt` | `String` | No |  |
| `merchantDirectDebitMandateID` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `merchantTokenDescription` | `String` | No |  |
| `notificationEmailAddresses` | `String` | No |  |
| `notificationRoleIDs` | `Array` | No |  |
| `orderID` | `String` | No |  |
| `partialPaymentMethod` | `String` | No |  |
| `partialPaymentSteps` | `String` | No |  |
| `paymentAttempts` | `Array` | No |  |
| `paymentInitiationID` | `String` | No |  |
| `paymentMethods` | `Array` | No |  |
| `paymentProcessor` | `String` | No |  |
| `paymentRequests` | `Array` | No |  |
| `payrunID` | `String` | No |  |
| `pispAccountID` | `String` | No |  |
| `priorityBankID` | `String` | No |  |
| `result` | `Hash` | No |  |
| `sandboxSettleDelayInSeconds` | `Integer` | No |  |
| `shippingAddress` | `Hash` | No |  |
| `status` | `String` | No |  |
| `successWebHookUrl` | `String` | No |  |
| `tags` | `Array` | No |  |
| `title` | `String` | No |  |
| `tokenisedCards` | `Array` | No |  |
| `transactions` | `Array` | No |  |
| `useHostedPaymentPage` | `Boolean` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PaymentRequest.create({
  "createdByUser" => {}, # Hash
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
| `applePayTransactionID` | `String` | No |  |
| `cardAuthorizationResponseID` | `String` | No |  |
| `cardExpiryMonth` | `Integer` | No |  |
| `cardExpiryYear` | `Integer` | No |  |
| `cardIssuer` | `String` | No |  |
| `cardIssuerCountry` | `String` | No |  |
| `cardLastFourDigits` | `String` | No |  |
| `cardRequestID` | `String` | No |  |
| `cardScheme` | `String` | No |  |
| `cardTokenCustomerID` | `String` | No |  |
| `cardTransactionID` | `String` | No |  |
| `currency` | `String` | No |  |
| `directDebitPaymentID` | `String` | No |  |
| `directDebitPaymentReference` | `String` | No |  |
| `drirectDebitMandateID` | `String` | No |  |
| `errorMessage` | `String` | No |  |
| `errorReason` | `String` | No |  |
| `eventType` | `String` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `lightningInvoice` | `String` | No |  |
| `lightningRHash` | `String` | No |  |
| `originUrl` | `String` | No |  |
| `paymentMethodType` | `String` | No |  |
| `paymentProcessorName` | `String` | No |  |
| `paymentRequestID` | `String` | No |  |
| `pispBankStatus` | `String` | No |  |
| `pispPaymentInitiationID` | `String` | No |  |
| `pispPaymentInstitutionName` | `String` | No |  |
| `pispPaymentServiceProviderID` | `String` | No |  |
| `pispRedirectUrl` | `String` | No |  |
| `reconciledTransactionID` | `String` | No |  |
| `refundPayoutID` | `String` | No |  |
| `status` | `String` | No |  |
| `walletName` | `String` | No |  |

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
| `amountPending` | `Float` | No |  |
| `amountReceived` | `Float` | No |  |
| `amountRefunded` | `Float` | No |  |
| `callbackUrl` | `String` | No |  |
| `cardStripePaymentIntentSecret` | `String` | No |  |
| `countryCode` | `String` | No |  |
| `currency` | `String` | No |  |
| `customFieldsToDisplay` | `Array` | No |  |
| `description` | `String` | No |  |
| `dueDate` | `String` | No |  |
| `fieldDisplaySettings` | `Array` | No |  |
| `googlePayMerchantID` | `String` | No |  |
| `id` | `String` | No |  |
| `jwk` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `merchantLogoUrlPng` | `String` | No |  |
| `merchantLogoUrlSvg` | `String` | No |  |
| `merchantName` | `String` | No |  |
| `merchantShortName` | `String` | No |  |
| `partialPaymentMethod` | `String` | No |  |
| `paymentAttempts` | `Array` | No |  |
| `paymentMethodsList` | `Array` | No |  |
| `paymentProcessor` | `String` | No |  |
| `paymentProcessorKey` | `String` | No |  |
| `pispError` | `String` | No |  |
| `priorityBankID` | `String` | No |  |
| `status` | `String` | No |  |
| `stripeAccountID` | `String` | No |  |
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
| `amountPending` | `Float` | No |  |
| `amountReceived` | `Float` | No |  |
| `amountRefunded` | `Float` | No |  |
| `currency` | `String` | No |  |
| `customerID` | `String` | No |  |
| `paymentRequestID` | `String` | No |  |
| `payments` | `Array` | No |  |
| `pispAuthorizations` | `Array` | No |  |
| `requestedAmount` | `Float` | No |  |
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
| `accountID` | `String` | No |  |
| `allowIncomplete` | `Boolean` | No |  |
| `amount` | `Float` | No |  |
| `amountMinorUnits` | `Integer` | No |  |
| `approvePayoutUrl` | `String` | No |  |
| `approverID` | `String` | No |  |
| `authenticationMethods` | `Array` | No |  |
| `authorisations` | `Array` | No |  |
| `authorisersCompletedCount` | `Integer` | No |  |
| `authorisersRequiredCount` | `Integer` | No |  |
| `batchPayoutID` | `String` | No |  |
| `beneficiary` | `Hash` | Yes |  |
| `beneficiaryID` | `String` | No |  |
| `canAuthorise` | `Boolean` | No |  |
| `canProcess` | `Boolean` | No |  |
| `canUpdate` | `Boolean` | No |  |
| `chargeBearer` | `String` | No |  |
| `createdBy` | `String` | No |  |
| `createdByEmailAddress` | `String` | No |  |
| `currency` | `String` | No |  |
| `currentUserID` | `String` | No |  |
| `description` | `String` | No |  |
| `destination` | `Hash` | No |  |
| `documents` | `Array` | No |  |
| `events` | `Array` | No |  |
| `failedPayouts` | `Hash` | No |  |
| `formattedAmount` | `String` | No |  |
| `formattedFxDestinationAmount` | `String` | No |  |
| `formattedSchedule` | `String` | No |  |
| `formattedScheduleDayOnly` | `String` | No |  |
| `formattedSourceAccountAvailableBalance` | `String` | No |  |
| `fxDestinationAmount` | `Float` | No |  |
| `fxDestinationAmountMinorUnits` | `Integer` | No |  |
| `fxDestinationCurrency` | `String` | No |  |
| `fxQuoteExpiresAt` | `String` | No |  |
| `fxQuoteID` | `String` | No |  |
| `fxRate` | `Float` | No |  |
| `fxUseDestinationAmount` | `Boolean` | No |  |
| `hasCurrentUserAuthorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `invoiceID` | `String` | No |  |
| `isArchived` | `Boolean` | No |  |
| `isFailed` | `Boolean` | No |  |
| `isSettled` | `Boolean` | No |  |
| `isSubmitted` | `Boolean` | No |  |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `merchantTokenDescription` | `String` | No |  |
| `nonce` | `String` | No |  |
| `paymentProcessor` | `String` | No |  |
| `paymentRail` | `String` | No |  |
| `payouts` | `Array` | No |  |
| `payrunID` | `String` | No |  |
| `payrunName` | `String` | No |  |
| `reason` | `String` | No |  |
| `rule` | `Hash` | No |  |
| `scheduleDate` | `String` | No |  |
| `scheduled` | `Boolean` | No |  |
| `sourceAccountAvailableBalance` | `Float` | No |  |
| `sourceAccountAvailableBalanceMinorUnits` | `Integer` | No |  |
| `sourceAccountBic` | `String` | No |  |
| `sourceAccountCurrency` | `String` | No |  |
| `sourceAccountIban` | `String` | No |  |
| `sourceAccountIdentifier` | `Hash` | Yes |  |
| `sourceAccountName` | `String` | No |  |
| `sourceAccountNumber` | `String` | No |  |
| `sourceAccountSortcode` | `String` | No |  |
| `status` | `String` | No |  |
| `tagIds` | `Array` | No |  |
| `tags` | `Array` | No |  |
| `theirReference` | `String` | No |  |
| `topupPayrunID` | `String` | No |  |
| `transactedAmount` | `Float` | No |  |
| `transactedFxAmount` | `Float` | No |  |
| `transactedFxRate` | `Float` | No |  |
| `type` | `String` | No |  |
| `userID` | `String` | No |  |
| `yourReference` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Payout.create({
  "beneficiary" => {}, # Hash
  "sourceAccountIdentifier" => {}, # Hash
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
| `accountID` | `String` | No |  |
| `amount` | `Float` | No |  |
| `amountMinorUnits` | `Integer` | No |  |
| `approvePayoutUrl` | `String` | No |  |
| `approverID` | `String` | No |  |
| `authenticationMethods` | `Array` | No |  |
| `authorisations` | `Array` | No |  |
| `authorisersCompletedCount` | `Integer` | No |  |
| `authorisersRequiredCount` | `Integer` | No |  |
| `batchPayoutID` | `String` | No |  |
| `beneficiary` | `Hash` | Yes |  |
| `canAuthorise` | `Boolean` | No |  |
| `canProcess` | `Boolean` | No |  |
| `canUpdate` | `Boolean` | No |  |
| `chargeBearer` | `String` | No |  |
| `createdBy` | `String` | No |  |
| `createdByEmailAddress` | `String` | No |  |
| `currency` | `String` | No |  |
| `currentUserID` | `String` | No |  |
| `description` | `String` | No |  |
| `destination` | `Hash` | No |  |
| `documents` | `Array` | No |  |
| `events` | `Array` | No |  |
| `formattedAmount` | `String` | No |  |
| `formattedFxDestinationAmount` | `String` | No |  |
| `formattedSchedule` | `String` | No |  |
| `formattedScheduleDayOnly` | `String` | No |  |
| `formattedSourceAccountAvailableBalance` | `String` | No |  |
| `fxDestinationAmount` | `Float` | No |  |
| `fxDestinationAmountMinorUnits` | `Integer` | No |  |
| `fxDestinationCurrency` | `String` | No |  |
| `fxQuoteExpiresAt` | `String` | No |  |
| `fxQuoteID` | `String` | No |  |
| `fxRate` | `Float` | No |  |
| `fxUseDestinationAmount` | `Boolean` | No |  |
| `hasCurrentUserAuthorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `invoiceID` | `String` | No |  |
| `isArchived` | `Boolean` | No |  |
| `isFailed` | `Boolean` | No |  |
| `isSettled` | `Boolean` | No |  |
| `isSubmitted` | `Boolean` | No |  |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `merchantTokenDescription` | `String` | No |  |
| `nonce` | `String` | No |  |
| `paymentProcessor` | `String` | No |  |
| `paymentRail` | `String` | No |  |
| `payrunID` | `String` | No |  |
| `payrunName` | `String` | No |  |
| `rule` | `Hash` | No |  |
| `scheduleDate` | `String` | No |  |
| `scheduled` | `Boolean` | No |  |
| `sourceAccountAvailableBalance` | `Float` | No |  |
| `sourceAccountAvailableBalanceMinorUnits` | `Integer` | No |  |
| `sourceAccountBic` | `String` | No |  |
| `sourceAccountCurrency` | `String` | No |  |
| `sourceAccountIban` | `String` | No |  |
| `sourceAccountIdentifier` | `Hash` | Yes |  |
| `sourceAccountName` | `String` | No |  |
| `sourceAccountNumber` | `String` | No |  |
| `sourceAccountSortcode` | `String` | No |  |
| `status` | `String` | No |  |
| `tags` | `Array` | No |  |
| `theirReference` | `String` | No |  |
| `topupPayrunID` | `String` | No |  |
| `transactedAmount` | `Float` | No |  |
| `transactedFxAmount` | `Float` | No |  |
| `transactedFxRate` | `Float` | No |  |
| `type` | `String` | No |  |
| `userID` | `String` | No |  |
| `yourReference` | `String` | No |  |

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
| `authorisationDate` | `String` | No |  |
| `authorisations` | `Array` | No |  |
| `authorisersCompletedCount` | `Integer` | No |  |
| `authorisersRequiredCount` | `Integer` | No |  |
| `batchPayoutID` | `String` | No |  |
| `canAuthorise` | `Boolean` | No |  |
| `canDelete` | `Boolean` | No |  |
| `canEdit` | `Boolean` | No |  |
| `events` | `Array` | No |  |
| `hasCurrentUserAuthorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `invoices` | `Array` | No |  |
| `invoicesMinimal` | `Array` | No |  |
| `isArchived` | `Boolean` | No |  |
| `lastUpdated` | `String` | No |  |
| `lastUpdatedBy` | `Hash` | Yes |  |
| `merchantID` | `String` | No |  |
| `name` | `String` | No |  |
| `nonce` | `String` | No |  |
| `notes` | `String` | No |  |
| `payments` | `Array` | No |  |
| `payouts` | `Array` | No |  |
| `payoutsCount` | `Integer` | No |  |
| `reason` | `String` | No |  |
| `scheduleDate` | `String` | No |  |
| `scheduledDate` | `String` | No |  |
| `sourceAccounts` | `Array` | No |  |
| `status` | `String` | No |  |
| `totalEur` | `Float` | No |  |
| `totalGbp` | `Float` | No |  |
| `totalUsd` | `Float` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Payrun.create({
  "id" => "example_id", # String
  "lastUpdatedBy" => {}, # Hash
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
| `contentType` | `String` | No |  |
| `contents` | `String` | No |  |
| `lastCompletedAt` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `reportName` | `String` | No |  |
| `reportType` | `String` | No |  |
| `statementNumber` | `Integer` | No |  |

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
| `failedRoles` | `Hash` | No |  |
| `roles` | `Array` | No |  |

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
| `accountID` | `String` | No |  |
| `approveUrl` | `String` | No |  |
| `approverID` | `String` | No |  |
| `authenticationMethods` | `Array` | No |  |
| `authorisations` | `Array` | No |  |
| `authorisersCompletedCount` | `Integer` | No |  |
| `authorisersRequiredCount` | `Integer` | No |  |
| `canAuthorise` | `Boolean` | No |  |
| `createdBy` | `Hash` | Yes |  |
| `description` | `String` | No |  |
| `endAt` | `String` | No |  |
| `hasCurrentUserAuthorised` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `isDisabled` | `Boolean` | No |  |
| `lastExecutedAt` | `String` | No |  |
| `lastRunAtTransactionDate` | `String` | No |  |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `name` | `String` | No |  |
| `nonce` | `String` | Yes |  |
| `onApprovedWebHookUrl` | `String` | No |  |
| `onExecutionErrorWebHookUrl` | `String` | No |  |
| `onExecutionSuccessWebHookUrl` | `String` | No |  |
| `startAt` | `String` | No |  |
| `status` | `String` | No |  |
| `sweepAction` | `Hash` | No |  |
| `timeZoneId` | `String` | No |  |
| `triggerCronExpression` | `String` | No |  |
| `triggerOnPayIn` | `Boolean` | No |  |
| `userID` | `String` | No |  |
| `webHookSecret` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Rule.create({
  "createdBy" => {}, # Hash
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
| `errorMessage` | `String` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `isAuthoriseToEnable` | `Boolean` | No |  |
| `message` | `String` | No |  |
| `rawResponse` | `String` | No |  |
| `ruleEventType` | `String` | No |  |
| `ruleID` | `String` | No |  |
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
| `colourHex` | `String` | No |  |
| `description` | `String` | No |  |
| `id` | `String` | No |  |
| `merchantID` | `String` | Yes |  |
| `name` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Tag.create({
  "merchant_id" => "example_merchant_id", # String
  "merchantID" => "example_merchantID", # String
  "name" => "example_name", # String
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
| `accountID` | `String` | No |  |
| `accountName` | `String` | No |  |
| `accountSequenceNumber` | `Integer` | No |  |
| `addressDetails` | `Hash` | No |  |
| `amount` | `Float` | No |  |
| `amountMinorUnits` | `Integer` | No |  |
| `balance` | `Float` | No |  |
| `balanceMinorUnits` | `Integer` | No |  |
| `bookingDateTime` | `String` | No |  |
| `chargeDetails` | `Hash` | No |  |
| `content` | `Array` | No |  |
| `counterparty` | `Hash` | No |  |
| `counterpartySummary` | `String` | No |  |
| `currency` | `String` | No |  |
| `currencyExchange` | `Hash` | No |  |
| `date` | `String` | No |  |
| `description` | `String` | No |  |
| `enrichment` | `Hash` | No |  |
| `fxAmount` | `Float` | No |  |
| `fxCurrency` | `String` | No |  |
| `fxRate` | `Float` | No |  |
| `grossAmount` | `Hash` | Yes |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `isoBankTransactionCode` | `Hash` | No |  |
| `merchant` | `Hash` | No |  |
| `merchantID` | `String` | No |  |
| `pageNumber` | `Integer` | No |  |
| `pageSize` | `Integer` | No |  |
| `payeeDetails` | `Hash` | Yes |  |
| `payerDetails` | `Hash` | Yes |  |
| `paymentRequestCustomFields` | `Hash` | No |  |
| `paymentRequestID` | `String` | No |  |
| `payoutID` | `String` | No |  |
| `proprietaryBankTransactionCode` | `Hash` | No |  |
| `rawReference` | `String` | No |  |
| `reference` | `String` | No |  |
| `ruleID` | `String` | No |  |
| `statementReferences` | `Array` | No |  |
| `status` | `String` | No |  |
| `supplementaryData` | `Object` | No |  |
| `tags` | `Array` | No |  |
| `theirReference` | `String` | No |  |
| `totalPages` | `Integer` | No |  |
| `totalSize` | `Integer` | No |  |
| `transactionAmount` | `Hash` | Yes |  |
| `transactionDate` | `String` | No |  |
| `transactionInformation` | `Array` | No |  |
| `transactionMutability` | `String` | No |  |
| `type` | `String` | No |  |
| `valueDateTime` | `String` | No |  |
| `virtualIBAN` | `String` | No |  |
| `yourReference` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Transaction.create({
  "id" => "example_id", # String
  "grossAmount" => {}, # Hash
  "payeeDetails" => {}, # Hash
  "payerDetails" => {}, # Hash
  "transactionAmount" => {}, # Hash
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
| `clientSessionTimeouts` | `Array` | No |  |
| `emailAddress` | `String` | Yes |  |
| `firstName` | `String` | Yes |  |
| `id` | `String` | No |  |
| `lastName` | `String` | Yes |  |
| `passkeyAdded` | `Boolean` | No |  |
| `permissions` | `Hash` | No |  |
| `profile` | `String` | No |  |
| `rolesWithScope` | `Array` | No |  |
| `twoFactorEnabled` | `Boolean` | No |  |
| `userInviteID` | `String` | No |  |

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
| `authorisationStatus` | `Hash` | No |  |
| `failedUserInvites` | `Hash` | No |  |
| `id` | `String` | No |  |
| `initialRoleID` | `String` | No |  |
| `inviteeEmailAddress` | `String` | No |  |
| `inviteeFirstName` | `String` | No |  |
| `inviteeLastName` | `String` | No |  |
| `inviterEmailAddress` | `String` | No |  |
| `inviterFirstName` | `String` | No |  |
| `inviterLastName` | `String` | No |  |
| `isAuthorised` | `Boolean` | No |  |
| `isInviteeRegistered` | `Boolean` | No |  |
| `lastInvited` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `merchantName` | `String` | No |  |
| `message` | `String` | No |  |
| `registrationUrl` | `String` | No |  |
| `sendInviteEmail` | `Boolean` | No |  |
| `status` | `String` | No |  |
| `user` | `Hash` | Yes |  |
| `userID` | `String` | No |  |
| `userInvites` | `Array` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.UserInvite.create({
  "user" => {}, # Hash
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
| `accountName` | `String` | No |  |
| `accountSupplierName` | `String` | No |  |
| `availableBalance` | `Float` | No |  |
| `availableBalanceMinorUnits` | `Integer` | No |  |
| `balance` | `Float` | No |  |
| `balanceMinorUnits` | `Integer` | No |  |
| `bankName` | `String` | No |  |
| `consentID` | `String` | No |  |
| `createdBy` | `Hash` | Yes |  |
| `createdByDisplayName` | `String` | No |  |
| `currency` | `String` | No |  |
| `defaultPaymentRail` | `String` | No |  |
| `displayName` | `String` | No |  |
| `expiryDate` | `String` | No |  |
| `externalAccountIcon` | `String` | No |  |
| `id` | `String` | No |  |
| `identifier` | `Hash` | Yes |  |
| `inserted` | `String` | No |  |
| `isArchived` | `Boolean` | No |  |
| `isConnectedAccount` | `Boolean` | No |  |
| `isDefault` | `Boolean` | No |  |
| `isTrustAccount` | `Boolean` | No |  |
| `isVirtual` | `Boolean` | No |  |
| `lastTransaction` | `Hash` | No |  |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No |  |
| `merchantName` | `String` | No |  |
| `name` | `String` | Yes |  |
| `physicalAccountID` | `String` | No |  |
| `rules` | `Array` | No |  |
| `submittedPayoutsBalance` | `Float` | No |  |
| `submittedPayoutsBalanceMinorUnits` | `Integer` | No |  |
| `summary` | `String` | No |  |
| `supplierSepaInstantStatus` | `String` | No |  |
| `xeroBankFeedConnectionStatus` | `String` | No |  |
| `xeroBankFeedLastSyncedAt` | `String` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `String` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `String` | No |  |
| `xeroBankFeedSyncStatus` | `String` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `Integer` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Virtual.create({
  "account_id" => "example_account_id", # String
  "createdBy" => {}, # Hash
  "identifier" => {}, # Hash
  "name" => "example_name", # String
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
| `destinationUrl` | `String` | No |  |
| `emailAddress` | `String` | No |  |
| `failedNotificationEmailAddress` | `String` | No |  |
| `id` | `String` | No |  |
| `isActive` | `Boolean` | No |  |
| `merchantID` | `String` | No |  |
| `notificationMethod` | `String` | No |  |
| `resourceTypes` | `Array` | No |  |
| `retry` | `Boolean` | No |  |
| `secret` | `String` | No |  |
| `version` | `Integer` | No |  |

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

