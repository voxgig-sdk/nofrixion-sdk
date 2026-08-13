# Nofrixion Golang SDK Reference

Complete API reference for the Nofrixion Golang SDK.


## NofrixionSDK

### Constructor

```go
func NewNofrixionSDK(options map[string]any) *NofrixionSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *NofrixionSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *NofrixionSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `Account(data map[string]any) NofrixionEntity`

Create a new `Account` entity instance. Pass `nil` for no initial data.

#### `Batch(data map[string]any) NofrixionEntity`

Create a new `Batch` entity instance. Pass `nil` for no initial data.

#### `Beneficiary(data map[string]any) NofrixionEntity`

Create a new `Beneficiary` entity instance. Pass `nil` for no initial data.

#### `BeneficiaryGroup(data map[string]any) NofrixionEntity`

Create a new `BeneficiaryGroup` entity instance. Pass `nil` for no initial data.

#### `Card(data map[string]any) NofrixionEntity`

Create a new `Card` entity instance. Pass `nil` for no initial data.

#### `CardCustomerToken(data map[string]any) NofrixionEntity`

Create a new `CardCustomerToken` entity instance. Pass `nil` for no initial data.

#### `CardPayment(data map[string]any) NofrixionEntity`

Create a new `CardPayment` entity instance. Pass `nil` for no initial data.

#### `CardPublicKey(data map[string]any) NofrixionEntity`

Create a new `CardPublicKey` entity instance. Pass `nil` for no initial data.

#### `Consent(data map[string]any) NofrixionEntity`

Create a new `Consent` entity instance. Pass `nil` for no initial data.

#### `Currency(data map[string]any) NofrixionEntity`

Create a new `Currency` entity instance. Pass `nil` for no initial data.

#### `DirectDebitBatchSubmit(data map[string]any) NofrixionEntity`

Create a new `DirectDebitBatchSubmit` entity instance. Pass `nil` for no initial data.

#### `FxRate(data map[string]any) NofrixionEntity`

Create a new `FxRate` entity instance. Pass `nil` for no initial data.

#### `IPayment(data map[string]any) NofrixionEntity`

Create a new `IPayment` entity instance. Pass `nil` for no initial data.

#### `Mandate(data map[string]any) NofrixionEntity`

Create a new `Mandate` entity instance. Pass `nil` for no initial data.

#### `Merchant(data map[string]any) NofrixionEntity`

Create a new `Merchant` entity instance. Pass `nil` for no initial data.

#### `MerchantAuthorisationSetting(data map[string]any) NofrixionEntity`

Create a new `MerchantAuthorisationSetting` entity instance. Pass `nil` for no initial data.

#### `MerchantDirectDebitMandatePage(data map[string]any) NofrixionEntity`

Create a new `MerchantDirectDebitMandatePage` entity instance. Pass `nil` for no initial data.

#### `MerchantPayByBankSetting(data map[string]any) NofrixionEntity`

Create a new `MerchantPayByBankSetting` entity instance. Pass `nil` for no initial data.

#### `MerchantPaymentRequestTemplate(data map[string]any) NofrixionEntity`

Create a new `MerchantPaymentRequestTemplate` entity instance. Pass `nil` for no initial data.

#### `MerchantToken(data map[string]any) NofrixionEntity`

Create a new `MerchantToken` entity instance. Pass `nil` for no initial data.

#### `Metadata(data map[string]any) NofrixionEntity`

Create a new `Metadata` entity instance. Pass `nil` for no initial data.

#### `NoFrixionVersion(data map[string]any) NofrixionEntity`

Create a new `NoFrixionVersion` entity instance. Pass `nil` for no initial data.

#### `OpenBanking(data map[string]any) NofrixionEntity`

Create a new `OpenBanking` entity instance. Pass `nil` for no initial data.

#### `Payeeverification(data map[string]any) NofrixionEntity`

Create a new `Payeeverification` entity instance. Pass `nil` for no initial data.

#### `Payment(data map[string]any) NofrixionEntity`

Create a new `Payment` entity instance. Pass `nil` for no initial data.

#### `PaymentAccount(data map[string]any) NofrixionEntity`

Create a new `PaymentAccount` entity instance. Pass `nil` for no initial data.

#### `PaymentAccountMinimal(data map[string]any) NofrixionEntity`

Create a new `PaymentAccountMinimal` entity instance. Pass `nil` for no initial data.

#### `PaymentInitiation(data map[string]any) NofrixionEntity`

Create a new `PaymentInitiation` entity instance. Pass `nil` for no initial data.

#### `PaymentRequest(data map[string]any) NofrixionEntity`

Create a new `PaymentRequest` entity instance. Pass `nil` for no initial data.

#### `PaymentRequestEvent(data map[string]any) NofrixionEntity`

Create a new `PaymentRequestEvent` entity instance. Pass `nil` for no initial data.

#### `PaymentRequestMetric(data map[string]any) NofrixionEntity`

Create a new `PaymentRequestMetric` entity instance. Pass `nil` for no initial data.

#### `PaymentRequestMinimal(data map[string]any) NofrixionEntity`

Create a new `PaymentRequestMinimal` entity instance. Pass `nil` for no initial data.

#### `PaymentRequestResult(data map[string]any) NofrixionEntity`

Create a new `PaymentRequestResult` entity instance. Pass `nil` for no initial data.

#### `Payout(data map[string]any) NofrixionEntity`

Create a new `Payout` entity instance. Pass `nil` for no initial data.

#### `PayoutKeysetPage(data map[string]any) NofrixionEntity`

Create a new `PayoutKeysetPage` entity instance. Pass `nil` for no initial data.

#### `PayoutMetric(data map[string]any) NofrixionEntity`

Create a new `PayoutMetric` entity instance. Pass `nil` for no initial data.

#### `Payrun(data map[string]any) NofrixionEntity`

Create a new `Payrun` entity instance. Pass `nil` for no initial data.

#### `Report(data map[string]any) NofrixionEntity`

Create a new `Report` entity instance. Pass `nil` for no initial data.

#### `ReportResult(data map[string]any) NofrixionEntity`

Create a new `ReportResult` entity instance. Pass `nil` for no initial data.

#### `Role(data map[string]any) NofrixionEntity`

Create a new `Role` entity instance. Pass `nil` for no initial data.

#### `Rule(data map[string]any) NofrixionEntity`

Create a new `Rule` entity instance. Pass `nil` for no initial data.

#### `RuleEvent(data map[string]any) NofrixionEntity`

Create a new `RuleEvent` entity instance. Pass `nil` for no initial data.

#### `Tag(data map[string]any) NofrixionEntity`

Create a new `Tag` entity instance. Pass `nil` for no initial data.

#### `Token(data map[string]any) NofrixionEntity`

Create a new `Token` entity instance. Pass `nil` for no initial data.

#### `Transaction(data map[string]any) NofrixionEntity`

Create a new `Transaction` entity instance. Pass `nil` for no initial data.

#### `User(data map[string]any) NofrixionEntity`

Create a new `User` entity instance. Pass `nil` for no initial data.

#### `UserInvite(data map[string]any) NofrixionEntity`

Create a new `UserInvite` entity instance. Pass `nil` for no initial data.

#### `Virtual(data map[string]any) NofrixionEntity`

Create a new `Virtual` entity instance. Pass `nil` for no initial data.

#### `Webhook(data map[string]any) NofrixionEntity`

Create a new `Webhook` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## AccountEntity

```go
account := client.Account(nil)
fmt.Println(account.GetName()) // "account"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountBalances` | `[]any` | No |  |
| `accountID` | `string` | No |  |
| `accountIdentifications` | `[]any` | No |  |
| `accountName` | `string` | No |  |
| `accountNames` | `[]any` | No |  |
| `accountSupplierName` | `string` | No |  |
| `accountType` | `string` | No |  |
| `availableBalance` | `float64` | No |  |
| `availableBalanceMinorUnits` | `int` | No |  |
| `balance` | `float64` | No |  |
| `balanceMinorUnits` | `int` | No |  |
| `bankName` | `string` | No |  |
| `consentID` | `string` | No |  |
| `consolidatedAccountInformation` | `map[string]any` | No |  |
| `createdBy` | `map[string]any` | Yes |  |
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
| `identifier` | `map[string]any` | Yes |  |
| `inserted` | `string` | No |  |
| `isArchived` | `bool` | No |  |
| `isConnectedAccount` | `bool` | No |  |
| `isDefault` | `bool` | No |  |
| `isTrustAccount` | `bool` | No |  |
| `isVirtual` | `bool` | No |  |
| `lastTransaction` | `map[string]any` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `nickname` | `string` | No |  |
| `physicalAccountID` | `string` | No |  |
| `roleIDs` | `[]any` | No |  |
| `rules` | `[]any` | No |  |
| `submittedPayoutsBalance` | `float64` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Account(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Account(nil).Load(map[string]any{"id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Account(nil).Create(map[string]any{
    "createdBy": map[string]any{},
    "identifier": map[string]any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Account(nil).Update(map[string]any{
    "id": "account_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Account(nil).Remove(map[string]any{"id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `AccountEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## BatchEntity

```go
batch := client.Batch(nil)
fmt.Println(batch.GetName()) // "batch"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approveUrl` | `string` | No |  |
| `id` | `string` | No |  |
| `payouts` | `[]any` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Batch(nil).Load(map[string]any{"id": "batch_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Batch(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `BatchEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## BeneficiaryEntity

```go
beneficiary := client.Beneficiary(nil)
fmt.Println(beneficiary.GetName()) // "beneficiary"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approvalCallbackUrl` | `string` | No |  |
| `authenticationMethods` | `[]any` | No |  |
| `authorisations` | `[]any` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `beneficiaries` | `[]any` | No |  |
| `beneficiaryEvents` | `[]any` | No |  |
| `canAuthorise` | `bool` | No |  |
| `canUpdate` | `bool` | No |  |
| `createdBy` | `map[string]any` | Yes |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `map[string]any` | No |  |
| `failedBeneficiaries` | `map[string]any` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isEnabled` | `bool` | No |  |
| `lastAuthorised` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `sourceAccountIDs` | `[]any` | No |  |
| `sourceAccounts` | `[]any` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Beneficiary(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Beneficiary(nil).Load(map[string]any{"id": "beneficiary_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Beneficiary(nil).Create(map[string]any{
    "createdBy": map[string]any{},
    "currency": "example_currency",
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Beneficiary(nil).Update(map[string]any{
    "id": "beneficiary_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Beneficiary(nil).Remove(map[string]any{"id": "beneficiary_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `BeneficiaryEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## BeneficiaryGroupEntity

```go
beneficiaryGroup := client.BeneficiaryGroup(nil)
fmt.Println(beneficiaryGroup.GetName()) // "beneficiary_group"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `groupMembers` | `[]any` | No |  |
| `groupName` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | Yes |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.BeneficiaryGroup(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `BeneficiaryGroupEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CardEntity

```go
card := client.Card(nil)
fmt.Println(card.GetName()) // "card"
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

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Card(nil).Create(map[string]any{
    "paymentrequest_id": "example_paymentrequest_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CardEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CardCustomerTokenEntity

```go
cardCustomerToken := client.CardCustomerToken(nil)
fmt.Println(cardCustomerToken.GetName()) // "card_customer_token"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.CardCustomerToken(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.CardCustomerToken(nil).Load(map[string]any{"customer_email_address": "customer_email_address"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.CardCustomerToken(nil).Remove(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CardCustomerTokenEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CardPaymentEntity

```go
cardPayment := client.CardPayment(nil)
fmt.Println(cardPayment.GetName()) // "card_payment"
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

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.CardPayment(nil).Create(map[string]any{
    "paymentrequest_id": "example_paymentrequest_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CardPaymentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CardPublicKeyEntity

```go
cardPublicKey := client.CardPublicKey(nil)
fmt.Println(cardPublicKey.GetName()) // "card_public_key"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jwt` | `string` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.CardPublicKey(nil).Load(map[string]any{"paymentrequest_id": "paymentrequest_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CardPublicKeyEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ConsentEntity

```go
consent := client.Consent(nil)
fmt.Println(consent.GetName()) // "consent"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Consent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Consent(nil).Load(map[string]any{"id": "consent_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Consent(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Consent(nil).Update(map[string]any{
    "id": "consent_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Consent(nil).Remove(map[string]any{"id": "consent_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ConsentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## CurrencyEntity

```go
currency := client.Currency(nil)
fmt.Println(currency.GetName()) // "currency"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Currency(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `CurrencyEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## DirectDebitBatchSubmitEntity

```go
directDebitBatchSubmit := client.DirectDebitBatchSubmit(nil)
fmt.Println(directDebitBatchSubmit.GetName()) // "direct_debit_batch_submit"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failedSubmissions` | `map[string]any` | No |  |
| `successfulSubmissions` | `[]any` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.DirectDebitBatchSubmit(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `DirectDebitBatchSubmitEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## FxRateEntity

```go
fxRate := client.FxRate(nil)
fmt.Println(fxRate.GetName()) // "fx_rate"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destinationCurrency` | `string` | No |  |
| `exchangeRate` | `float64` | No |  |
| `expiryTime` | `string` | No |  |
| `quoteID` | `string` | No |  |
| `sourceCurrency` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.FxRate(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.FxRate(nil).Load(map[string]any{"destination": "destination", "source": "source", "valid_for_minute": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `FxRateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## IPaymentEntity

```go
iPayment := client.IPayment(nil)
fmt.Println(iPayment.GetName()) // "i_payment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `paymentRequestID` | `string` | No |  |
| `responseType` | `string` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.IPayment(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `IPaymentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MandateEntity

```go
mandate := client.Mandate(nil)
fmt.Println(mandate.GetName()) // "mandate"
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Mandate(nil).Load(map[string]any{"id": "mandate_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Mandate(nil).Create(map[string]any{
    "addressLine1": "example_addressLine1",
    "city": "example_city",
    "countryCode": "example_countryCode",
    "emailAddress": "example_emailAddress",
    "firstName": "example_firstName",
    "lastName": "example_lastName",
    "postalCode": "example_postalCode",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MandateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MerchantEntity

```go
merchant := client.Merchant(nil)
fmt.Println(merchant.GetName()) // "merchant"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountCurrencies` | `[]any` | No |  |
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
| `parentMerchant` | `map[string]any` | No |  |
| `paymentAccountLimit` | `int` | No |  |
| `paymentAccounts` | `[]any` | No |  |
| `reason` | `string` | No |  |
| `shortName` | `string` | No |  |
| `supportedPaymentMethodsList` | `[]any` | No |  |
| `suspensionReason` | `string` | No |  |
| `tags` | `[]any` | No |  |
| `timeZoneId` | `string` | No |  |
| `tradingName` | `string` | No |  |
| `webHookLimit` | `int` | No |  |
| `yourRoleName` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Merchant(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Merchant(nil).Load(map[string]any{"id": "merchant_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Merchant(nil).Update(map[string]any{
    "id": "merchant_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Merchant(nil).Remove(map[string]any{"id": "merchant_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MerchantEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MerchantAuthorisationSettingEntity

```go
merchantAuthorisationSetting := client.MerchantAuthorisationSetting(nil)
fmt.Println(merchantAuthorisationSetting.GetName()) // "merchant_authorisation_setting"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amountLower` | `float64` | No |  |
| `amountUpper` | `float64` | No |  |
| `authorisationType` | `string` | No |  |
| `beneficiariesOnly` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lastEditorCantAuthorise` | `bool` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `numberOfAuthorisers` | `int` | No |  |
| `roleSettings` | `[]any` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.MerchantAuthorisationSetting(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MerchantAuthorisationSettingEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MerchantDirectDebitMandatePageEntity

```go
merchantDirectDebitMandatePage := client.MerchantDirectDebitMandatePage(nil)
fmt.Println(merchantDirectDebitMandatePage.GetName()) // "merchant_direct_debit_mandate_page"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.MerchantDirectDebitMandatePage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MerchantDirectDebitMandatePageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MerchantPayByBankSettingEntity

```go
merchantPayByBankSetting := client.MerchantPayByBankSetting(nil)
fmt.Println(merchantPayByBankSetting.GetName()) // "merchant_pay_by_bank_setting"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bankCountryCodes` | `[]any` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.MerchantPayByBankSetting(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MerchantPayByBankSettingEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MerchantPaymentRequestTemplateEntity

```go
merchantPaymentRequestTemplate := client.MerchantPaymentRequestTemplate(nil)
fmt.Println(merchantPaymentRequestTemplate.GetName()) // "merchant_payment_request_template"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bankPaymentOptions` | `map[string]any` | No |  |
| `cardPaymentAddressOptions` | `map[string]any` | No |  |
| `cardPaymentCaptureOptions` | `map[string]any` | No |  |
| `customFields` | `[]any` | No |  |
| `defaultFields` | `[]any` | No |  |
| `description` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `name` | `string` | Yes |  |
| `notificationOptions` | `map[string]any` | No |  |
| `paymentMethods` | `map[string]any` | No |  |
| `paymentTerms` | `map[string]any` | No |  |
| `priorityBankOptions` | `map[string]any` | No |  |
| `template` | `map[string]any` | Yes |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.MerchantPaymentRequestTemplate(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.MerchantPaymentRequestTemplate(nil).Load(map[string]any{"id": "merchant_payment_request_template_id", "paymentrequest_id": "paymentrequest_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.MerchantPaymentRequestTemplate(nil).Update(map[string]any{
    "id": "merchant_payment_request_template_id",
    "paymentrequest_id": "paymentrequest_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.MerchantPaymentRequestTemplate(nil).Remove(map[string]any{"id": "merchant_payment_request_template_id", "paymentrequest_id": "paymentrequest_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MerchantPaymentRequestTemplateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MerchantTokenEntity

```go
merchantToken := client.MerchantToken(nil)
fmt.Println(merchantToken.GetName()) // "merchant_token"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authenticationMethods` | `[]any` | No |  |
| `authorisations` | `[]any` | No |  |
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
| `permissionTypes` | `[]any` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.MerchantToken(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.MerchantToken(nil).Load(map[string]any{"id": "merchant_token_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.MerchantToken(nil).Create(map[string]any{
    "nonce": "example_nonce",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.MerchantToken(nil).Update(map[string]any{
    "id": "merchant_token_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MerchantTokenEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MetadataEntity

```go
metadata := client.Metadata(nil)
fmt.Println(metadata.GetName()) // "metadata"
```

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Metadata(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MetadataEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## NoFrixionVersionEntity

```go
noFrixionVersion := client.NoFrixionVersion(nil)
fmt.Println(noFrixionVersion.GetName()) // "no_frixion_version"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `buildVersion` | `int` | No |  |
| `majorVersion` | `int` | No |  |
| `minorVersion` | `int` | No |  |
| `releaseName` | `string` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.NoFrixionVersion(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `NoFrixionVersionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## OpenBankingEntity

```go
openBanking := client.OpenBanking(nil)
fmt.Println(openBanking.GetName()) // "open_banking"
```

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.OpenBanking(nil).Create(map[string]any{
    "account_id": "example_account_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.OpenBanking(nil).Remove(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `OpenBankingEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PayeeverificationEntity

```go
payeeverification := client.Payeeverification(nil)
fmt.Println(payeeverification.GetName()) // "payeeverification"
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

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Payeeverification(nil).Create(map[string]any{
    "accountName": "example_accountName",
    "iban": "example_iban",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PayeeverificationEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PaymentEntity

```go
payment := client.Payment(nil)
fmt.Println(payment.GetName()) // "payment"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresses` | `[]any` | No |  |
| `amount` | `float64` | No |  |
| `amountPending` | `float64` | No |  |
| `amountReceived` | `float64` | No |  |
| `amountRefunded` | `float64` | No |  |
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
| `createdByUser` | `map[string]any` | Yes |  |
| `currency` | `string` | No |  |
| `customFields` | `[]any` | No |  |
| `customerEmailAddress` | `string` | No |  |
| `customerID` | `string` | No |  |
| `customerName` | `string` | No |  |
| `description` | `string` | No |  |
| `destinationAccount` | `map[string]any` | No |  |
| `directDebitPayment` | `map[string]any` | No |  |
| `dueDate` | `string` | No |  |
| `events` | `[]any` | No |  |
| `failureCallbackUrl` | `string` | No |  |
| `fieldDisplaySettings` | `[]any` | No |  |
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
| `notificationRoleIDs` | `[]any` | No |  |
| `orderID` | `string` | No |  |
| `partialPaymentMethod` | `string` | No |  |
| `partialPaymentSteps` | `string` | No |  |
| `paymentAttempts` | `[]any` | No |  |
| `paymentMethods` | `[]any` | No |  |
| `paymentProcessor` | `string` | No |  |
| `payrunID` | `string` | No |  |
| `pispAccountID` | `string` | No |  |
| `priorityBankID` | `string` | No |  |
| `result` | `map[string]any` | No |  |
| `sandboxSettleDelayInSeconds` | `int` | No |  |
| `shippingAddress` | `map[string]any` | No |  |
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
| `tagIds` | `[]any` | No |  |
| `tags` | `[]any` | No |  |
| `title` | `string` | No |  |
| `tokenisedCards` | `[]any` | No |  |
| `transactions` | `[]any` | No |  |
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Payment(nil).Load(map[string]any{"id": "payment_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Payment(nil).Create(map[string]any{
    "createdByUser": map[string]any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Payment(nil).Update(map[string]any{
    "id": "payment_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PaymentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PaymentAccountEntity

```go
paymentAccount := client.PaymentAccount(nil)
fmt.Println(paymentAccount.GetName()) // "payment_account"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `string` | No |  |
| `accountSupplierName` | `string` | No |  |
| `availableBalance` | `float64` | No |  |
| `availableBalanceMinorUnits` | `int` | No |  |
| `balance` | `float64` | No |  |
| `balanceMinorUnits` | `int` | No |  |
| `bankName` | `string` | No |  |
| `consentID` | `string` | No |  |
| `createdBy` | `map[string]any` | Yes |  |
| `createdByDisplayName` | `string` | No |  |
| `currency` | `string` | No |  |
| `defaultPaymentRail` | `string` | No |  |
| `displayName` | `string` | No |  |
| `expiryDate` | `string` | No |  |
| `externalAccountIcon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `map[string]any` | Yes |  |
| `inserted` | `string` | No |  |
| `isArchived` | `bool` | No |  |
| `isConnectedAccount` | `bool` | No |  |
| `isDefault` | `bool` | No |  |
| `isTrustAccount` | `bool` | No |  |
| `isVirtual` | `bool` | No |  |
| `lastTransaction` | `map[string]any` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `physicalAccountID` | `string` | No |  |
| `rules` | `[]any` | No |  |
| `submittedPayoutsBalance` | `float64` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PaymentAccount(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PaymentAccountEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PaymentAccountMinimalEntity

```go
paymentAccountMinimal := client.PaymentAccountMinimal(nil)
fmt.Println(paymentAccountMinimal.GetName()) // "payment_account_minimal"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `string` | No |  |
| `availableBalance` | `float64` | No |  |
| `balance` | `float64` | No |  |
| `balanceMinorUnits` | `int` | No |  |
| `currency` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `map[string]any` | Yes |  |
| `isArchived` | `bool` | No |  |
| `isConnectedAccount` | `bool` | No |  |
| `merchantID` | `string` | No |  |
| `submittedPayoutsBalance` | `float64` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PaymentAccountMinimal(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PaymentAccountMinimalEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PaymentInitiationEntity

```go
paymentInitiation := client.PaymentInitiation(nil)
fmt.Println(paymentInitiation.GetName()) // "payment_initiation"
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

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PaymentInitiation(nil).Create(map[string]any{
    "paymentrequest_id": "example_paymentrequest_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PaymentInitiationEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PaymentRequestEntity

```go
paymentRequest := client.PaymentRequest(nil)
fmt.Println(paymentRequest.GetName()) // "payment_request"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresses` | `[]any` | No |  |
| `amount` | `float64` | No |  |
| `amountPending` | `float64` | No |  |
| `amountReceived` | `float64` | No |  |
| `amountRefunded` | `float64` | No |  |
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
| `createdByUser` | `map[string]any` | Yes |  |
| `currency` | `string` | No |  |
| `customFields` | `[]any` | No |  |
| `customerEmailAddress` | `string` | No |  |
| `customerID` | `string` | No |  |
| `customerName` | `string` | No |  |
| `description` | `string` | No |  |
| `destinationAccount` | `map[string]any` | No |  |
| `directDebitPayment` | `map[string]any` | No |  |
| `doSimulateSettlementFailure` | `bool` | No |  |
| `dueDate` | `string` | No |  |
| `errorDescription` | `string` | No |  |
| `events` | `[]any` | No |  |
| `failedPaymentRequests` | `map[string]any` | No |  |
| `failureCallbackUrl` | `string` | No |  |
| `fieldDisplaySettings` | `[]any` | No |  |
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
| `notificationRoleIDs` | `[]any` | No |  |
| `orderID` | `string` | No |  |
| `partialPaymentMethod` | `string` | No |  |
| `partialPaymentSteps` | `string` | No |  |
| `paymentAttempts` | `[]any` | No |  |
| `paymentInitiationID` | `string` | No |  |
| `paymentMethods` | `[]any` | No |  |
| `paymentProcessor` | `string` | No |  |
| `paymentRequests` | `[]any` | No |  |
| `payrunID` | `string` | No |  |
| `pispAccountID` | `string` | No |  |
| `priorityBankID` | `string` | No |  |
| `result` | `map[string]any` | No |  |
| `sandboxSettleDelayInSeconds` | `int` | No |  |
| `shippingAddress` | `map[string]any` | No |  |
| `status` | `string` | No |  |
| `successWebHookUrl` | `string` | No |  |
| `tags` | `[]any` | No |  |
| `title` | `string` | No |  |
| `tokenisedCards` | `[]any` | No |  |
| `transactions` | `[]any` | No |  |
| `useHostedPaymentPage` | `bool` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PaymentRequest(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.PaymentRequest(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PaymentRequest(nil).Create(map[string]any{
    "createdByUser": map[string]any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.PaymentRequest(nil).Update(map[string]any{
    "paymentrequest_id": "paymentrequest_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.PaymentRequest(nil).Remove(map[string]any{"id": "id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PaymentRequestEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PaymentRequestEventEntity

```go
paymentRequestEvent := client.PaymentRequestEvent(nil)
fmt.Println(paymentRequestEvent.GetName()) // "payment_request_event"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float64` | Yes |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PaymentRequestEvent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PaymentRequestEventEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PaymentRequestMetricEntity

```go
paymentRequestMetric := client.PaymentRequestMetric(nil)
fmt.Println(paymentRequestMetric.GetName()) // "payment_request_metric"
```

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.PaymentRequestMetric(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PaymentRequestMetricEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PaymentRequestMinimalEntity

```go
paymentRequestMinimal := client.PaymentRequestMinimal(nil)
fmt.Println(paymentRequestMinimal.GetName()) // "payment_request_minimal"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float64` | No |  |
| `amountPending` | `float64` | No |  |
| `amountReceived` | `float64` | No |  |
| `amountRefunded` | `float64` | No |  |
| `callbackUrl` | `string` | No |  |
| `cardStripePaymentIntentSecret` | `string` | No |  |
| `countryCode` | `string` | No |  |
| `currency` | `string` | No |  |
| `customFieldsToDisplay` | `[]any` | No |  |
| `description` | `string` | No |  |
| `dueDate` | `string` | No |  |
| `fieldDisplaySettings` | `[]any` | No |  |
| `googlePayMerchantID` | `string` | No |  |
| `id` | `string` | No |  |
| `jwk` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantLogoUrlPng` | `string` | No |  |
| `merchantLogoUrlSvg` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `merchantShortName` | `string` | No |  |
| `partialPaymentMethod` | `string` | No |  |
| `paymentAttempts` | `[]any` | No |  |
| `paymentMethodsList` | `[]any` | No |  |
| `paymentProcessor` | `string` | No |  |
| `paymentProcessorKey` | `string` | No |  |
| `pispError` | `string` | No |  |
| `priorityBankID` | `string` | No |  |
| `status` | `string` | No |  |
| `stripeAccountID` | `string` | No |  |
| `title` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PaymentRequestMinimal(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PaymentRequestMinimalEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PaymentRequestResultEntity

```go
paymentRequestResult := client.PaymentRequestResult(nil)
fmt.Println(paymentRequestResult.GetName()) // "payment_request_result"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float64` | No |  |
| `amountPending` | `float64` | No |  |
| `amountReceived` | `float64` | No |  |
| `amountRefunded` | `float64` | No |  |
| `currency` | `string` | No |  |
| `customerID` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |
| `payments` | `[]any` | No |  |
| `pispAuthorizations` | `[]any` | No |  |
| `requestedAmount` | `float64` | No |  |
| `result` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PaymentRequestResult(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PaymentRequestResultEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PayoutEntity

```go
payout := client.Payout(nil)
fmt.Println(payout.GetName()) // "payout"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountID` | `string` | No |  |
| `allowIncomplete` | `bool` | No |  |
| `amount` | `float64` | No |  |
| `amountMinorUnits` | `int` | No |  |
| `approvePayoutUrl` | `string` | No |  |
| `approverID` | `string` | No |  |
| `authenticationMethods` | `[]any` | No |  |
| `authorisations` | `[]any` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `batchPayoutID` | `string` | No |  |
| `beneficiary` | `map[string]any` | Yes |  |
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
| `destination` | `map[string]any` | No |  |
| `documents` | `[]any` | No |  |
| `events` | `[]any` | No |  |
| `failedPayouts` | `map[string]any` | No |  |
| `formattedAmount` | `string` | No |  |
| `formattedFxDestinationAmount` | `string` | No |  |
| `formattedSchedule` | `string` | No |  |
| `formattedScheduleDayOnly` | `string` | No |  |
| `formattedSourceAccountAvailableBalance` | `string` | No |  |
| `fxDestinationAmount` | `float64` | No |  |
| `fxDestinationAmountMinorUnits` | `int` | No |  |
| `fxDestinationCurrency` | `string` | No |  |
| `fxQuoteExpiresAt` | `string` | No |  |
| `fxQuoteID` | `string` | No |  |
| `fxRate` | `float64` | No |  |
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
| `payouts` | `[]any` | No |  |
| `payrunID` | `string` | No |  |
| `payrunName` | `string` | No |  |
| `reason` | `string` | No |  |
| `rule` | `map[string]any` | No |  |
| `scheduleDate` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `sourceAccountAvailableBalance` | `float64` | No |  |
| `sourceAccountAvailableBalanceMinorUnits` | `int` | No |  |
| `sourceAccountBic` | `string` | No |  |
| `sourceAccountCurrency` | `string` | No |  |
| `sourceAccountIban` | `string` | No |  |
| `sourceAccountIdentifier` | `map[string]any` | Yes |  |
| `sourceAccountName` | `string` | No |  |
| `sourceAccountNumber` | `string` | No |  |
| `sourceAccountSortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tagIds` | `[]any` | No |  |
| `tags` | `[]any` | No |  |
| `theirReference` | `string` | No |  |
| `topupPayrunID` | `string` | No |  |
| `transactedAmount` | `float64` | No |  |
| `transactedFxAmount` | `float64` | No |  |
| `transactedFxRate` | `float64` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Payout(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Payout(nil).Load(map[string]any{"id": "payout_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Payout(nil).Create(map[string]any{
    "beneficiary": map[string]any{},
    "sourceAccountIdentifier": map[string]any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Payout(nil).Update(map[string]any{
    "id": "payout_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Payout(nil).Remove(map[string]any{"id": "payout_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PayoutEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PayoutKeysetPageEntity

```go
payoutKeysetPage := client.PayoutKeysetPage(nil)
fmt.Println(payoutKeysetPage.GetName()) // "payout_keyset_page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountID` | `string` | No |  |
| `amount` | `float64` | No |  |
| `amountMinorUnits` | `int` | No |  |
| `approvePayoutUrl` | `string` | No |  |
| `approverID` | `string` | No |  |
| `authenticationMethods` | `[]any` | No |  |
| `authorisations` | `[]any` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `batchPayoutID` | `string` | No |  |
| `beneficiary` | `map[string]any` | Yes |  |
| `canAuthorise` | `bool` | No |  |
| `canProcess` | `bool` | No |  |
| `canUpdate` | `bool` | No |  |
| `chargeBearer` | `string` | No |  |
| `createdBy` | `string` | No |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | No |  |
| `currentUserID` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `map[string]any` | No |  |
| `documents` | `[]any` | No |  |
| `events` | `[]any` | No |  |
| `formattedAmount` | `string` | No |  |
| `formattedFxDestinationAmount` | `string` | No |  |
| `formattedSchedule` | `string` | No |  |
| `formattedScheduleDayOnly` | `string` | No |  |
| `formattedSourceAccountAvailableBalance` | `string` | No |  |
| `fxDestinationAmount` | `float64` | No |  |
| `fxDestinationAmountMinorUnits` | `int` | No |  |
| `fxDestinationCurrency` | `string` | No |  |
| `fxQuoteExpiresAt` | `string` | No |  |
| `fxQuoteID` | `string` | No |  |
| `fxRate` | `float64` | No |  |
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
| `rule` | `map[string]any` | No |  |
| `scheduleDate` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `sourceAccountAvailableBalance` | `float64` | No |  |
| `sourceAccountAvailableBalanceMinorUnits` | `int` | No |  |
| `sourceAccountBic` | `string` | No |  |
| `sourceAccountCurrency` | `string` | No |  |
| `sourceAccountIban` | `string` | No |  |
| `sourceAccountIdentifier` | `map[string]any` | Yes |  |
| `sourceAccountName` | `string` | No |  |
| `sourceAccountNumber` | `string` | No |  |
| `sourceAccountSortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tags` | `[]any` | No |  |
| `theirReference` | `string` | No |  |
| `topupPayrunID` | `string` | No |  |
| `transactedAmount` | `float64` | No |  |
| `transactedFxAmount` | `float64` | No |  |
| `transactedFxRate` | `float64` | No |  |
| `type` | `string` | No |  |
| `userID` | `string` | No |  |
| `yourReference` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PayoutKeysetPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PayoutKeysetPageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PayoutMetricEntity

```go
payoutMetric := client.PayoutMetric(nil)
fmt.Println(payoutMetric.GetName()) // "payout_metric"
```

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.PayoutMetric(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PayoutMetricEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PayrunEntity

```go
payrun := client.Payrun(nil)
fmt.Println(payrun.GetName()) // "payrun"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisationDate` | `string` | No |  |
| `authorisations` | `[]any` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `batchPayoutID` | `string` | No |  |
| `canAuthorise` | `bool` | No |  |
| `canDelete` | `bool` | No |  |
| `canEdit` | `bool` | No |  |
| `events` | `[]any` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoices` | `[]any` | No |  |
| `invoicesMinimal` | `[]any` | No |  |
| `isArchived` | `bool` | No |  |
| `lastUpdated` | `string` | No |  |
| `lastUpdatedBy` | `map[string]any` | Yes |  |
| `merchantID` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | No |  |
| `notes` | `string` | No |  |
| `payments` | `[]any` | No |  |
| `payouts` | `[]any` | No |  |
| `payoutsCount` | `int` | No |  |
| `reason` | `string` | No |  |
| `scheduleDate` | `string` | No |  |
| `scheduledDate` | `string` | No |  |
| `sourceAccounts` | `[]any` | No |  |
| `status` | `string` | No |  |
| `totalEur` | `float64` | No |  |
| `totalGbp` | `float64` | No |  |
| `totalUsd` | `float64` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Payrun(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Payrun(nil).Load(map[string]any{"id": "payrun_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Payrun(nil).Create(map[string]any{
    "id": "example_id",
    "lastUpdatedBy": map[string]any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Payrun(nil).Update(map[string]any{
    "id": "payrun_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Payrun(nil).Remove(map[string]any{"id": "payrun_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PayrunEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ReportEntity

```go
report := client.Report(nil)
fmt.Println(report.GetName()) // "report"
```

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Report(nil).Update(map[string]any{
    "id": "id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ReportEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ReportResultEntity

```go
reportResult := client.ReportResult(nil)
fmt.Println(reportResult.GetName()) // "report_result"
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.ReportResult(nil).Load(map[string]any{"id": 1, "report_id": "report_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ReportResultEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## RoleEntity

```go
role := client.Role(nil)
fmt.Println(role.GetName()) // "role"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failedRoles` | `map[string]any` | No |  |
| `roles` | `[]any` | No |  |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Role(nil).Create(map[string]any{
    "merchant_id": "example_merchant_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `RoleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## RuleEntity

```go
rule := client.Rule(nil)
fmt.Println(rule.GetName()) // "rule"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account` | `map[string]any` | No |  |
| `accountID` | `string` | No |  |
| `approveUrl` | `string` | No |  |
| `approverID` | `string` | No |  |
| `authenticationMethods` | `[]any` | No |  |
| `authorisations` | `[]any` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `canAuthorise` | `bool` | No |  |
| `createdBy` | `map[string]any` | Yes |  |
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
| `sweepAction` | `map[string]any` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Rule(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Rule(nil).Load(map[string]any{"id": "rule_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Rule(nil).Create(map[string]any{
    "createdBy": map[string]any{},
    "nonce": "example_nonce",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Rule(nil).Update(map[string]any{
    "id": "rule_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Rule(nil).Remove(map[string]any{"id": "rule_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `RuleEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## RuleEventEntity

```go
ruleEvent := client.RuleEvent(nil)
fmt.Println(ruleEvent.GetName()) // "rule_event"
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
| `user` | `map[string]any` | Yes |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.RuleEvent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `RuleEventEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TagEntity

```go
tag := client.Tag(nil)
fmt.Println(tag.GetName()) // "tag"
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Tag(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Tag(nil).Create(map[string]any{
    "merchant_id": "example_merchant_id",
    "merchantID": "example_merchantID",
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TagEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TokenEntity

```go
token := client.Token(nil)
fmt.Println(token.GetName()) // "token"
```

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Token(nil).Create(map[string]any{
    "id": "example_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Token(nil).Remove(map[string]any{"id": "id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TokenEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TransactionEntity

```go
transaction := client.Transaction(nil)
fmt.Println(transaction.GetName()) // "transaction"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountID` | `string` | No |  |
| `accountName` | `string` | No |  |
| `accountSequenceNumber` | `int` | No |  |
| `addressDetails` | `map[string]any` | No |  |
| `amount` | `float64` | No |  |
| `amountMinorUnits` | `int` | No |  |
| `balance` | `float64` | No |  |
| `balanceMinorUnits` | `int` | No |  |
| `bookingDateTime` | `string` | No |  |
| `chargeDetails` | `map[string]any` | No |  |
| `content` | `[]any` | No |  |
| `counterparty` | `map[string]any` | No |  |
| `counterpartySummary` | `string` | No |  |
| `currency` | `string` | No |  |
| `currencyExchange` | `map[string]any` | No |  |
| `date` | `string` | No |  |
| `description` | `string` | No |  |
| `enrichment` | `map[string]any` | No |  |
| `fxAmount` | `float64` | No |  |
| `fxCurrency` | `string` | No |  |
| `fxRate` | `float64` | No |  |
| `grossAmount` | `map[string]any` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isoBankTransactionCode` | `map[string]any` | No |  |
| `merchant` | `map[string]any` | No |  |
| `merchantID` | `string` | No |  |
| `pageNumber` | `int` | No |  |
| `pageSize` | `int` | No |  |
| `payeeDetails` | `map[string]any` | Yes |  |
| `payerDetails` | `map[string]any` | Yes |  |
| `paymentRequestCustomFields` | `map[string]any` | No |  |
| `paymentRequestID` | `string` | No |  |
| `payoutID` | `string` | No |  |
| `proprietaryBankTransactionCode` | `map[string]any` | No |  |
| `rawReference` | `string` | No |  |
| `reference` | `string` | No |  |
| `ruleID` | `string` | No |  |
| `statementReferences` | `[]any` | No |  |
| `status` | `string` | No |  |
| `supplementaryData` | `any` | No |  |
| `tags` | `[]any` | No |  |
| `theirReference` | `string` | No |  |
| `totalPages` | `int` | No |  |
| `totalSize` | `int` | No |  |
| `transactionAmount` | `map[string]any` | Yes |  |
| `transactionDate` | `string` | No |  |
| `transactionInformation` | `[]any` | No |  |
| `transactionMutability` | `string` | No |  |
| `type` | `string` | No |  |
| `valueDateTime` | `string` | No |  |
| `virtualIBAN` | `string` | No |  |
| `yourReference` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Transaction(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Transaction(nil).Load(map[string]any{"id": "transaction_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Transaction(nil).Create(map[string]any{
    "id": "example_id",
    "grossAmount": map[string]any{},
    "payeeDetails": map[string]any{},
    "payerDetails": map[string]any{},
    "transactionAmount": map[string]any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Transaction(nil).Remove(map[string]any{"id": "transaction_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TransactionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## UserEntity

```go
user := client.User(nil)
fmt.Println(user.GetName()) // "user"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `clientSessionTimeouts` | `[]any` | No |  |
| `emailAddress` | `string` | Yes |  |
| `firstName` | `string` | Yes |  |
| `id` | `string` | No |  |
| `lastName` | `string` | Yes |  |
| `passkeyAdded` | `bool` | No |  |
| `permissions` | `map[string]any` | No |  |
| `profile` | `string` | No |  |
| `rolesWithScope` | `[]any` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.User(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.User(nil).Update(map[string]any{
    "id": "id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `UserEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## UserInviteEntity

```go
userInvite := client.UserInvite(nil)
fmt.Println(userInvite.GetName()) // "user_invite"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisationStatus` | `map[string]any` | No |  |
| `failedUserInvites` | `map[string]any` | No |  |
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
| `user` | `map[string]any` | Yes |  |
| `userID` | `string` | No |  |
| `userInvites` | `[]any` | No |  |

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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.UserInvite(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.UserInvite(nil).Load(map[string]any{"id": "user_invite_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.UserInvite(nil).Create(map[string]any{
    "user": map[string]any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.UserInvite(nil).Update(map[string]any{
    "id": "user_invite_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.UserInvite(nil).Remove(map[string]any{"id": "user_invite_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `UserInviteEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## VirtualEntity

```go
virtual := client.Virtual(nil)
fmt.Println(virtual.GetName()) // "virtual"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `string` | No |  |
| `accountSupplierName` | `string` | No |  |
| `availableBalance` | `float64` | No |  |
| `availableBalanceMinorUnits` | `int` | No |  |
| `balance` | `float64` | No |  |
| `balanceMinorUnits` | `int` | No |  |
| `bankName` | `string` | No |  |
| `consentID` | `string` | No |  |
| `createdBy` | `map[string]any` | Yes |  |
| `createdByDisplayName` | `string` | No |  |
| `currency` | `string` | No |  |
| `defaultPaymentRail` | `string` | No |  |
| `displayName` | `string` | No |  |
| `expiryDate` | `string` | No |  |
| `externalAccountIcon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `map[string]any` | Yes |  |
| `inserted` | `string` | No |  |
| `isArchived` | `bool` | No |  |
| `isConnectedAccount` | `bool` | No |  |
| `isDefault` | `bool` | No |  |
| `isTrustAccount` | `bool` | No |  |
| `isVirtual` | `bool` | No |  |
| `lastTransaction` | `map[string]any` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `name` | `string` | Yes |  |
| `physicalAccountID` | `string` | No |  |
| `rules` | `[]any` | No |  |
| `submittedPayoutsBalance` | `float64` | No |  |
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

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Virtual(nil).Create(map[string]any{
    "account_id": "example_account_id",
    "createdBy": map[string]any{},
    "identifier": map[string]any{},
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Virtual(nil).Update(map[string]any{
    "account_id": "account_id",
    "id": "id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `VirtualEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## WebhookEntity

```go
webhook := client.Webhook(nil)
fmt.Println(webhook.GetName()) // "webhook"
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
| `resourceTypes` | `[]any` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Webhook(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Webhook(nil).Load(map[string]any{"id": "webhook_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Webhook(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Webhook(nil).Update(map[string]any{
    "id": "webhook_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Webhook(nil).Remove(map[string]any{"id": "webhook_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `WebhookEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewNofrixionSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

