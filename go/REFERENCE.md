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
| `accountBalances` | `[]any` | No | The various balances for the account. |
| `accountID` | `string` | No | ID of the account. |
| `accountIdentifications` | `[]any` | No | The canoncial identifiers for the account. |
| `accountName` | `string` | No | Name for the account |
| `accountNames` | `[]any` | No | Optional account names set by the account holder. |
| `accountSupplierName` | `string` | No | The payment account supplier name. |
| `accountType` | `string` | No | The type of account e.g. |
| `availableBalance` | `float64` | No | The current available balance of the account. |
| `availableBalanceMinorUnits` | `int` | No | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `float64` | No | Balance of the account. |
| `balanceMinorUnits` | `int` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `string` | No | The bank name for external accounts |
| `consentID` | `string` | No | The ID of the consent used to connect the external account. |
| `consolidatedAccountInformation` | `map[string]any` | No | Summary information regarding account balances of the overall account provided by the bank. |
| `createdBy` | `map[string]any` | Yes |  |
| `createdByDisplayName` | `string` | No | Either the name of the user, merchant token or api key that created the account |
| `currency` | `string` | No | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `string` | No | Indicates the default payment rail for this account. |
| `description` | `string` | No | Product name as defined by the financial institution for this account. |
| `details` | `string` | No | Supplementary specifications that might be provided by the Bank. |
| `displayName` | `string` | No | Gets a unique display name for the payment account. |
| `expiryDate` | `string` | No | The date that the external account will expire |
| `externalAccountIcon` | `string` | No | The Icon for external accounts |
| `format` | `string` | No | File format to save the statement as. |
| `fromDate` | `string` | No | Minimum transaction date for the statement. |
| `id` | `string` | No | Unique id for the account. |
| `identifier` | `map[string]any` | Yes |  |
| `inserted` | `string` | No | Timestamp when the account was created. |
| `isArchived` | `bool` | No | Indicates whether the account is archived. |
| `isConnectedAccount` | `bool` | No | Indicates if the payment account is an externally connected account. |
| `isDefault` | `bool` | No | Is the default account |
| `isTrustAccount` | `bool` | No | Indicates if the payment account is a trust account. |
| `isVirtual` | `bool` | No | True if the account is a virtual account. |
| `lastTransaction` | `map[string]any` | No |  |
| `lastUpdated` | `string` | No | Timestamp when the account was last updated. |
| `merchantID` | `string` | No | The ID of the merchant that owns the account. |
| `merchantName` | `string` | No | The name of the merchant that owns the account. |
| `nickname` | `string` | No | Nickname of the account that was provided by the account owner. |
| `physicalAccountID` | `string` | No | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `roleIDs` | `[]any` | No | Optional list of role IDs that will get access to the payment account when created. |
| `rules` | `[]any` | No | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `float64` | No | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `int` | No | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `string` | No | Gets a summary of the payments account's most important properties. |
| `supplierPhysicalAccountID` | `string` | No | For internal use only. |
| `supplierSepaInstantStatus` | `string` | No | Indicates the status of the SEPA Instant payment rail for this account. |
| `toDate` | `string` | No | Maximum transaction date for the statement. |
| `type` | `string` | No | Specifies the type of account e.g. |
| `usageType` | `string` | No |  |
| `xeroBankFeedConnectionStatus` | `string` | No | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `string` | No |  |
| `xeroBankFeedSyncStatus` | `string` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | No | Indicates the number of unsynchronised transactions with Xero |

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
    "account_id": "example_account_id",
    "currency": "example_currency",
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
| `approveUrl` | `string` | No | This field is used when returning a batch payout record to a client. |
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
| `authenticationMethods` | `[]any` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `[]any` | No | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `int` | No | The number of distinct authorisers that have authorised the beneficiary. |
| `authorisersRequiredCount` | `int` | No | The number of authorisers required for this beneficiary. |
| `beneficiaries` | `[]any` | No |  |
| `beneficiaryEvents` | `[]any` | No |  |
| `canAuthorise` | `bool` | No | True if the beneficiary can be authorised by the user who loaded it. |
| `canUpdate` | `bool` | No | True if the beneficiary can be updated by the user who loaded it. |
| `createdBy` | `map[string]any` | Yes |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | Yes | Gets or Sets the currency. |
| `destination` | `map[string]any` | No |  |
| `failedBeneficiaries` | `map[string]any` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isEnabled` | `bool` | No |  |
| `lastAuthorised` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No | Gets or Sets the merchant id. |
| `name` | `string` | Yes | The descriptive name for the beneficiary. |
| `nonce` | `string` | No |  |
| `sourceAccountIDs` | `[]any` | No | ID of the accounts which are authorised to act as a source for the beneficiary. |
| `sourceAccounts` | `[]any` | No |  |
| `theirReference` | `string` | No | The reference that will be used by default as TheirReference when creating payouts to this beneficiary if no TheirReference is specified for the payout. |

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
    "id": "example_id",
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
| `groupMembers` | `[]any` | No | The existing group members. |
| `groupName` | `string` | Yes | The descriptive name for the beneficiary group. |
| `id` | `string` | No |  |
| `inserted` | `string` | No | Timestamp indicating when the group was created. |
| `lastUpdated` | `string` | No | Timestamp indicating when the group was last updated. |
| `merchantID` | `string` | Yes | Gets or Sets the merchant id. |

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
| `isPayerAuthenticationRequired` | `bool` | No | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `bool` | No | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `string` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `string` | No | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `string` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `int` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `int` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `string` | No | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `string` | No |  |
| `requestID` | `string` | No |  |
| `responseCode` | `string` | No |  |
| `responseType` | `string` | No |  |
| `status` | `string` | No |  |
| `threeDSRedirectUrl` | `string` | No | Checkout.com require a redirect for 3DS authentication. |
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
| `cardType` | `string` | No | The type of the tokenised card, e.g. |
| `customerEmailAddress` | `string` | No | When creating a tokenised card the payer's email address must be supplied. |
| `expiryMonth` | `string` | No |  |
| `expiryYear` | `string` | No |  |
| `id` | `string` | No | The unique ID of the card token that has been stored for the customer. |
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
result, err := client.CardCustomerToken(nil).Remove(map[string]any{"id": "id"}, nil)
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
| `isPayerAuthenticationRequired` | `bool` | No | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `bool` | No | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `string` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `string` | No | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `string` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `int` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `int` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `string` | No | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `string` | No |  |
| `requestID` | `string` | No |  |
| `responseCode` | `string` | No |  |
| `responseType` | `string` | No |  |
| `status` | `string` | No |  |
| `threeDSRedirectUrl` | `string` | No | Checkout.com require a redirect for 3DS authentication. |
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
| `authorisationUrl` | `string` | No | The URL the authorising user needs to be redirected to in order to get the open banking consent token. |
| `callbackUrl` | `string` | No | Optional callback URL that the end user performing the open banking authorisation will be redirected to on completion. |
| `consentID` | `string` | No | The ID of the open banking consent. |
| `emailAddress` | `string` | No | The email address that identifies the end user that will be authorising the open banking consent request. |
| `expiryDate` | `string` | No |  |
| `failureCallbackUrl` | `string` | No | Optional callback URL for open banking consent authorisation failure. |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `institutionID` | `string` | No | The institution ID the open banking consent is being requested for. |
| `isConnectedAccounts` | `bool` | No | Optional setting. |
| `isEnabled` | `bool` | No |  |
| `merchantID` | `string` | No | The ID of the merchant the consent token is being created to be used with. |
| `provider` | `string` | No | Lists the supported card and PIS processors. |
| `successWebHookUrl` | `string` | No | A web hook URL to send an HTTP request to when the open banking consent is successfuly authorised. |

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
| `failedSubmissions` | `map[string]any` | No | Dictionary of failed submissions, keyed by the index (1-based) in the original request. |
| `successfulSubmissions` | `[]any` | No | List of successfully submitted direct debit payments. |

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
| `exchangeRate` | `float64` | No | The price at which the transaction will buy the source currency using the destination currency. |
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
| `accountNumber` | `string` | No | Account number of the customer's bank account in case of GBP account. |
| `addressLine1` | `string` | Yes | First line of the customer's address. |
| `addressLine2` | `string` | No | Second line of the customer's address. |
| `approvedAt` | `string` | No | Date at which the supplier approved this mandate. |
| `city` | `string` | Yes | Customer's city. |
| `countryCode` | `string` | Yes | 2-character country code of the customer's bank account. |
| `currency` | `string` | No | Currency of this mandate. |
| `customerAccountNumber` | `string` | No | Customer's account number in case of GBP account. |
| `customerCity` | `string` | No | Customer's city of residence. |
| `customerCountryCode` | `string` | No | Customer's country of residence code. |
| `customerCountryName` | `string` | No | Customer's country of residence. |
| `customerEmailAddress` | `string` | No | Customer's email address. |
| `customerFirstName` | `string` | No | Customer's first name. |
| `customerIban` | `string` | No | Customer's IBAN in case of EUR account. |
| `customerLastName` | `string` | No | Customer's last name. |
| `customerSortCode` | `string` | No | Customer's sort code in case of GBP account. |
| `emailAddress` | `string` | Yes | Customer's email address. |
| `firstName` | `string` | Yes | Customer's first name. |
| `iban` | `string` | No | IBAN of the customer's bank account in case of EUR account. |
| `id` | `string` | No | Internal ID of the mandate. |
| `inserted` | `string` | No | The timestamp this mandate was created at. |
| `isRecurring` | `bool` | No | Whether this mandate is single-use or recurring. |
| `lastName` | `string` | Yes | Customer's last name. |
| `lastUpdated` | `string` | No | The timestamp this mandate was last updated at. |
| `merchantID` | `string` | No | Internal ID of this mandate's merchant. |
| `postalCode` | `string` | Yes | Customer's postal code. |
| `reference` | `string` | No | Reference assigned to this mandate. |
| `sortCode` | `string` | No | Sort code of the customer's bank account in case of GBP account. |
| `status` | `string` | No | General status of this mandate. |
| `supplierBankAccountID` | `string` | No | ID that the supplier assigned to this mandate's bank account. |
| `supplierCustomerID` | `string` | No | ID that the supplier assigned to this mandate's customer. |
| `supplierMandateID` | `string` | No | ID that the supplier assigned to this mandate. |
| `supplierName` | `string` | No | Name of the supplier used to create this mandate. |
| `supplierStatus` | `string` | No | Last status that the supplier reported for this mandate. |

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
| `accountCurrencies` | `[]any` | No | The list of currencies that the merchant has accounts for. |
| `canHaveTrustAccounts` | `bool` | No | Trust accounts are a special type of account that allow the account name to be trusted for use in statements and verification of payee checks. |
| `cardPaymentProcessor` | `string` | No | Name of the card payment processor. |
| `companyID` | `string` | No | The Company ID recorded in the Compliance system. |
| `displayQrOnHostedPay` | `bool` | No | Indicates if a QR Code containing the payment link should be displayed on the hosted payment page. |
| `hostedPayVersion` | `int` | No | The version of the hosted payment page to use with the merchant. |
| `id` | `string` | No | Unique ID for the merchant. |
| `inserted` | `string` | No | Timestamp the merchant was added to MoneyMoov. |
| `isBlocked` | `bool` | No | The merchant is blocked from making payments (payouts). |
| `isExited` | `bool` | No | The merchant has formally terminated their relationship and is no longer a customer. |
| `isSuspended` | `bool` | No | The merchant has temporarily suspended their own account. |
| `jurisdiction` | `string` | No | The jurisdiction the merchant entity is incorporated or established in. |
| `logoUrlPng` | `string` | No | The CDN URL of the merchant's logo in PNG format. |
| `logoUrlSvg` | `string` | No | The CDN URL of the merchant's logo in SVG format. |
| `merchantCategoryCode` | `string` | No | The industry code that represents the merchant's primary trading activity. |
| `name` | `string` | No | The registered business name of the merchant. |
| `notes` | `string` | No | The notes field is an optional free text field that can be used to store any additional information about the merchant. |
| `parentMerchant` | `map[string]any` | No |  |
| `paymentAccountLimit` | `int` | No | The maximum number of payment accounts that can be created for the Merchant. |
| `paymentAccounts` | `[]any` | No |  |
| `reason` | `string` | No | The reason for the suspension. |
| `shortName` | `string` | No | A URL friendly shortish name for the merchant. |
| `supportedPaymentMethodsList` | `[]any` | No | The payment methods that are configured and supported for this merchant. |
| `suspensionReason` | `string` | No | The reason for the suspension, provided by the merchant. |
| `tags` | `[]any` | No | An optional list of descriptive tags that can be used on merchant entities such as payment requests. |
| `timeZoneId` | `string` | No | The IANA (Internet Assigned Numbers Authority) time zone identifier of the merchant. |
| `tradingName` | `string` | No | An optional trading name. |
| `webHookLimit` | `int` | No | The maximum number of web hooks that can be created for the Merchant. |
| `yourRoleName` | `string` | No | The name of the role for the identity that loaded the merchant record. |

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
result, err := client.Merchant(nil).Remove(map[string]any{"id": "merchant_id", "user_id": "user_id"}, nil)
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
| `approvedAt` | `string` | No | Date at which the supplier approved this mandate. |
| `currency` | `string` | No | Currency of this mandate. |
| `customerAccountNumber` | `string` | No | Customer's account number in case of GBP account. |
| `customerCity` | `string` | No | Customer's city of residence. |
| `customerCountryCode` | `string` | No | Customer's country of residence code. |
| `customerCountryName` | `string` | No | Customer's country of residence. |
| `customerEmailAddress` | `string` | No | Customer's email address. |
| `customerFirstName` | `string` | No | Customer's first name. |
| `customerIban` | `string` | No | Customer's IBAN in case of EUR account. |
| `customerLastName` | `string` | No | Customer's last name. |
| `customerSortCode` | `string` | No | Customer's sort code in case of GBP account. |
| `id` | `string` | No | Internal ID of the mandate. |
| `inserted` | `string` | No | The timestamp this mandate was created at. |
| `isRecurring` | `bool` | No | Whether this mandate is single-use or recurring. |
| `lastUpdated` | `string` | No | The timestamp this mandate was last updated at. |
| `merchantID` | `string` | No | Internal ID of this mandate's merchant. |
| `reference` | `string` | No | Reference assigned to this mandate. |
| `status` | `string` | No | General status of this mandate. |
| `supplierBankAccountID` | `string` | No | ID that the supplier assigned to this mandate's bank account. |
| `supplierCustomerID` | `string` | No | ID that the supplier assigned to this mandate's customer. |
| `supplierMandateID` | `string` | No | ID that the supplier assigned to this mandate. |
| `supplierName` | `string` | No | Name of the supplier used to create this mandate. |
| `supplierStatus` | `string` | No | Last status that the supplier reported for this mandate. |

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
| `bankCountryCodes` | `[]any` | No | The list of country codes representing the banks the country supports. |
| `bankID` | `string` | No | ID of the bank to be configured for the merchant. |
| `bankName` | `string` | No | Name of the Bank/Institution. |
| `businessInstitutionID` | `string` | No | ID that the processor uses to identify the bank (business accounts). |
| `currency` | `string` | No | Currency supported by the bank. |
| `logo` | `string` | No | URL of the bank's logo. |
| `message` | `string` | No | Message relating to specific bank. |
| `messageImageUrl` | `string` | No | Optional image URL to be displayed with the message. |
| `order` | `int` | No | Order in which this setting will appear in the UI. |
| `personalInstitutionID` | `string` | No | ID that the processor uses to identify the bank (personal accounts). |
| `processor` | `string` | No | Name of the bank payment processor. |
| `warningHeading` | `string` | No | The heading for a warning message related to the bank institution to be displayed to the user. |
| `warningMessage` | `string` | No | The warning message related to the bank institution to be displayed to the user. |

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
| `customFields` | `[]any` | No | A list of custom fields that can be included in the payment request template. |
| `defaultFields` | `[]any` | No | A list of default fields that are included in the payment request template. |
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
| `authenticationMethods` | `[]any` | No | A list of authentication types allowed to authorise the merchant token. |
| `authorisations` | `[]any` | No | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `int` | No | The number of distinct authorisers that have authorised the merchant token. |
| `authorisersRequiredCount` | `int` | No | The number of authorisers required for this merchant token. |
| `canAuthorise` | `bool` | No | True if the merchant token can be authorised by the user who loaded it. |
| `description` | `string` | No | Token description |
| `expiresAt` | `string` | No | Optional. |
| `hasCurrentUserAuthorised` | `bool` | No | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `hmacAlgorithm` | `string` | No | Optional shared secret algorithm to use for HMAC authentication. |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `ipAddressWhitelist` | `string` | No | Optional. |
| `isArchived` | `bool` | No | Indicates whether the merchant token is archived. |
| `isEnabled` | `bool` | No | If set to false the merchant token will not be accepted to authorise a request. |
| `lastAuthorised` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No | The merchant id to add to the token |
| `nonce` | `string` | Yes |  |
| `permissionTypes` | `[]any` | No | The permissions that the merchant token supports. |
| `requestSignatureVersion` | `int` | No | Represent the version of the overall merchant token. |
| `sharedSecretAlgorithm` | `string` | No | Optional shared secret algorithm to use for HMAC authentication. |
| `sharedSecretBase64` | `string` | No | The base 64 encoded shared secret that is used for request authentication with an HMAC. |
| `token` | `string` | No | The JWT merchant token. |

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
result, err := client.OpenBanking(nil).Remove(map[string]any{"account_id": "account_id"}, nil)
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
| `accountName` | `string` | Yes | The name of the account to verify |
| `accountNumber` | `string` | No | The account number of the account to verify (for CoP checks) |
| `iban` | `string` | Yes | The IBAN of the account to verify (for VoP checks) |
| `payeeVerifiedAccountName` | `string` | No | The verified account name of the payee, if available (in case of a close match) |
| `result` | `string` | No | The result of the payee verification |
| `secondaryIdentification` | `string` | No | Optional secondary identifier for the account to verify. |
| `sortCode` | `string` | No | The sort code of the account to verify (for CoP checks) |

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
| `amount` | `float64` | No | The amount of money to request. |
| `amountPending` | `float64` | No | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `float64` | No | Total amount received for this payment request. |
| `amountRefunded` | `float64` | No | Total amount refunded for this payment request. |
| `autoSendReceipt` | `bool` | No | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `string` | No | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `string` | No | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `bool` | No | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `bool` | No | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `string` | No | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `bool` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardNoPayerAuthentication` | `bool` | No | If set to true for card payments no attempt will be made to use payer authentication (3-D Secure and equivalent). |
| `cardProcessorMerchantID` | `string` | No | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `string` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `string` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `cardTransmitRawDetails` | `bool` | No | If set to true for card payments the sensitive card number and card verification number will be transmitted directly rather than being tokenised. |
| `createdByUser` | `map[string]any` | Yes |  |
| `currency` | `string` | No | The currency of the request. |
| `customFields` | `[]any` | No | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `string` | No | Optional email address for the customer. |
| `customerID` | `string` | No | An optional customer identifier for the payment request. |
| `customerName` | `string` | No |  |
| `description` | `string` | No | An optional description for the payment request. |
| `destinationAccount` | `map[string]any` | No |  |
| `directDebitPayment` | `map[string]any` | No | Contains information about a Direct Debit payment attempt for a payment request. |
| `dueDate` | `string` | No | The due date for the payment request. |
| `events` | `[]any` | No |  |
| `failureCallbackUrl` | `string` | No | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `[]any` | No | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `string` | No |  |
| `hostedPayCheckoutUrl` | `string` | No | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `string` | No |  |
| `ignoreAddressVerification` | `bool` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `string` | No | The timestamp the payment request was created at. |
| `insertedSortable` | `string` | No | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `isArchived` | `bool` | No | Indicates whether the payment request is archived. |
| `jwk` | `string` | No | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `string` | No | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `string` | No | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `string` | No | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `string` | No | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `string` | No | The ID of the merchant to create the payment request for. |
| `merchantTokenDescription` | `string` | No | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `string` | No |  |
| `notificationRoleIDs` | `[]any` | No | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `string` | No | An optional order ID for the payment request. |
| `partialPaymentMethod` | `string` | No | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `string` | No | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `[]any` | No | The payment attempts made against this payment request. |
| `paymentMethods` | `[]any` | No | The payment methods that the payment request supports. |
| `paymentProcessor` | `string` | No | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `payrunID` | `string` | No | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `string` | No | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `string` | No | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `map[string]any` | No |  |
| `sandboxSettleDelayInSeconds` | `int` | No | Sandbox only. |
| `shippingAddress` | `map[string]any` | No |  |
| `shippingAddressCity` | `string` | No | Optionally the city of the customer's shipping address. |
| `shippingAddressCountryCode` | `string` | No | Optionally the country code of the customer's shipping address. |
| `shippingAddressCounty` | `string` | No | Optionally the state or county of the customer's shipping address. |
| `shippingAddressLine1` | `string` | No | Optionally the first line of the customer's shipping address. |
| `shippingAddressLine2` | `string` | No | Optionally the second line of the customer's shipping address. |
| `shippingAddressPostCode` | `string` | No | Optionally the post code of the customer's shipping address. |
| `shippingEmail` | `string` | No | Optionally the shipping email address for the customer. |
| `shippingFirstName` | `string` | No | Optionally the first name of the customer's shipping address. |
| `shippingLastName` | `string` | No | Optionally the last name of the customer's shipping address. |
| `shippingPhone` | `string` | No | Optionally the shipping phone number for the customer. |
| `status` | `string` | No | The current status of the payment request. |
| `successWebHookUrl` | `string` | No | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tagIds` | `[]any` | No | An optional list of tag ids to add to the payment request |
| `tags` | `[]any` | No | An optional list of descriptive tags attached to the payment request. |
| `title` | `string` | No | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `[]any` | No |  |
| `transactions` | `[]any` | No |  |
| `useHostedPaymentPage` | `bool` | No | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

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
| `accountName` | `string` | No | Name for the account |
| `accountSupplierName` | `string` | No | The payment account supplier name. |
| `availableBalance` | `float64` | No | The current available balance of the account. |
| `availableBalanceMinorUnits` | `int` | No | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `float64` | No | Balance of the account. |
| `balanceMinorUnits` | `int` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `string` | No | The bank name for external accounts |
| `consentID` | `string` | No | The ID of the consent used to connect the external account. |
| `createdBy` | `map[string]any` | Yes |  |
| `createdByDisplayName` | `string` | No | Either the name of the user, merchant token or api key that created the account |
| `currency` | `string` | No | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `string` | No | Indicates the default payment rail for this account. |
| `displayName` | `string` | No | Gets a unique display name for the payment account. |
| `expiryDate` | `string` | No | The date that the external account will expire |
| `externalAccountIcon` | `string` | No | The Icon for external accounts |
| `id` | `string` | No | Unique id for the account. |
| `identifier` | `map[string]any` | Yes |  |
| `inserted` | `string` | No | Timestamp when the account was created. |
| `isArchived` | `bool` | No | Indicates whether the account is archived. |
| `isConnectedAccount` | `bool` | No | Indicates if the payment account is an externally connected account. |
| `isDefault` | `bool` | No | Is the default account |
| `isTrustAccount` | `bool` | No | Indicates if the payment account is a trust account. |
| `isVirtual` | `bool` | No | True if the account is a virtual account. |
| `lastTransaction` | `map[string]any` | No |  |
| `lastUpdated` | `string` | No | Timestamp when the account was last updated. |
| `merchantID` | `string` | No | The ID of the merchant that owns the account. |
| `merchantName` | `string` | No | The name of the merchant that owns the account. |
| `physicalAccountID` | `string` | No | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `[]any` | No | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `float64` | No | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `int` | No | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `string` | No | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `string` | No | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `string` | No | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `string` | No |  |
| `xeroBankFeedSyncStatus` | `string` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | No | Indicates the number of unsynchronised transactions with Xero |

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
| `accountName` | `string` | No | Name for the account |
| `availableBalance` | `float64` | No | The current available balance of the account. |
| `balance` | `float64` | No | Balance of the account. |
| `balanceMinorUnits` | `int` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `currency` | `string` | No | Currency of the account in ISO 4217 format |
| `id` | `string` | No | Unique id for the account. |
| `identifier` | `map[string]any` | Yes |  |
| `isArchived` | `bool` | No | Is the account archived |
| `isConnectedAccount` | `bool` | No | Indicates if the payment account is an externally connected account. |
| `merchantID` | `string` | No | The ID of the merchant that owns the account. |
| `submittedPayoutsBalance` | `float64` | No | Total of the payouts that have been submitted for processing. |

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
| `paymentInitiationID` | `string` | No | The unique identifier of the payment initiation request. |
| `paymentRequestCallbackUrl` | `string` | No | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `string` | No |  |
| `redirectUrl` | `string` | No | A redirect URL for the user to authorise the payment initiation request at the ASPSP |
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
| `amount` | `float64` | No | The amount of money to request. |
| `amountPending` | `float64` | No | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `float64` | No | Total amount received for this payment request. |
| `amountRefunded` | `float64` | No | Total amount refunded for this payment request. |
| `autoSendReceipt` | `bool` | No | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `string` | No | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `string` | No | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `bool` | No | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `bool` | No | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `string` | No | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `bool` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardProcessorMerchantID` | `string` | No | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `string` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `string` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `createdByUser` | `map[string]any` | Yes |  |
| `currency` | `string` | No | The currency of the request. |
| `customFields` | `[]any` | No | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `string` | No | Optional email address for the customer. |
| `customerID` | `string` | No | An optional customer identifier for the payment request. |
| `customerName` | `string` | No |  |
| `description` | `string` | No | An optional description for the payment request. |
| `destinationAccount` | `map[string]any` | No |  |
| `directDebitPayment` | `map[string]any` | No | Contains information about a Direct Debit payment attempt for a payment request. |
| `doSimulateSettlementFailure` | `bool` | No |  |
| `dueDate` | `string` | No | The due date for the payment request. |
| `errorDescription` | `string` | No |  |
| `events` | `[]any` | No |  |
| `failedPaymentRequests` | `map[string]any` | No |  |
| `failureCallbackUrl` | `string` | No | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `[]any` | No | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `string` | No |  |
| `hostedPayCheckoutUrl` | `string` | No | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `string` | No |  |
| `ignoreAddressVerification` | `bool` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `string` | No | The timestamp the payment request was created at. |
| `insertedSortable` | `string` | No | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `institution` | `string` | No |  |
| `isArchived` | `bool` | No | Indicates whether the payment request is archived. |
| `jwk` | `string` | No | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `string` | No | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `string` | No | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `string` | No | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `string` | No | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `string` | No |  |
| `merchantTokenDescription` | `string` | No | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `string` | No |  |
| `notificationRoleIDs` | `[]any` | No | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `string` | No | An optional order ID for the payment request. |
| `partialPaymentMethod` | `string` | No | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `string` | No | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `[]any` | No | The payment attempts made against this payment request. |
| `paymentInitiationID` | `string` | No |  |
| `paymentMethods` | `[]any` | No | The payment methods that the payment request supports. |
| `paymentProcessor` | `string` | No | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `paymentRequests` | `[]any` | No |  |
| `payrunID` | `string` | No | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `string` | No | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `string` | No | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `map[string]any` | No |  |
| `sandboxSettleDelayInSeconds` | `int` | No | Sandbox only. |
| `shippingAddress` | `map[string]any` | No |  |
| `status` | `string` | No | The current status of the payment request. |
| `successWebHookUrl` | `string` | No | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tags` | `[]any` | No | An optional list of descriptive tags attached to the payment request. |
| `title` | `string` | No | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `[]any` | No |  |
| `transactions` | `[]any` | No |  |
| `useHostedPaymentPage` | `bool` | No | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

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
result, err := client.PaymentRequest(nil).Load(map[string]any{"id": "payment_request_id"}, nil)
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
    "id": "payment_request_id",
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
result, err := client.PaymentRequest(nil).Remove(map[string]any{"id": "payment_request_id"}, nil)
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
| `applePayTransactionID` | `string` | No | Transaction ID received in Apple pay token. |
| `cardAuthorizationResponseID` | `string` | No | For a successful card authorization this field will hold the response ID. |
| `cardExpiryMonth` | `int` | No | For card payment events this field holds the payer's card expiry month. |
| `cardExpiryYear` | `int` | No | For card payment events this field holds the payer's card expiry year. |
| `cardIssuer` | `string` | No | For card payment events this field holds the payer's card issuer. |
| `cardIssuerCountry` | `string` | No | For card payment events this field holds the payer's card issuer country of origin. |
| `cardLastFourDigits` | `string` | No | For card payment events this field holds the payer's card last four digits. |
| `cardRequestID` | `string` | No |  |
| `cardScheme` | `string` | No | For card payment events this field holds the scheme of the payer's card, e.g. |
| `cardTokenCustomerID` | `string` | No | If the option to create a reusable token for card payments was set this field contains the token the merchant can store to use for repeat payments. |
| `cardTransactionID` | `string` | No |  |
| `currency` | `string` | No |  |
| `directDebitPaymentID` | `string` | No | Payment ID issued by the Direct Debit supplier. |
| `directDebitPaymentReference` | `string` | No | Reference string issued by the Direct Debit supplier. |
| `drirectDebitMandateID` | `string` | No | The ID of the mandate that was used wehn requesting payment. |
| `errorMessage` | `string` | No |  |
| `errorReason` | `string` | No |  |
| `eventType` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lightningInvoice` | `string` | No | For Bitcoin Lightning payments this field holds the invoice presented to the payer. |
| `lightningRHash` | `string` | No | For Bitcoin Lightning payments the hash of the invoice presented to the payer. |
| `originUrl` | `string` | No | Optional field that can be set by payment methods, such as pay by bank, that may want to redirect back to the URL that initiated the attempt in the case of a failure condition. |
| `paymentMethodType` | `string` | No | The type of payment method the event relates to, e.g. |
| `paymentProcessorName` | `string` | No | If the event was for a card payment this is the name of the card processor, e.g. |
| `paymentRequestID` | `string` | No |  |
| `pispBankStatus` | `string` | No | For payment initiation attempts some providers (e.g. |
| `pispPaymentInitiationID` | `string` | No | For a payment initiation this is the ID returned by the service provider initiating the payment for us. |
| `pispPaymentInstitutionName` | `string` | No | For a payment initiation this is the name of the financial institution that is used to initiate and authorise the payment. |
| `pispPaymentServiceProviderID` | `string` | No | For a payment initiation this is the service provider ID selected by the payer, typically the ID for the bank or similar financial institution. |
| `pispRedirectUrl` | `string` | No | For a payment initiation this is the redirect URL returned by the service provider initiating the payment for us. |
| `reconciledTransactionID` | `string` | No | For settlement events (only relevant for non-card payments) this is the payin transaction that the payment request event was reconciled with. |
| `refundPayoutID` | `string` | No | ID of the Payout that was created for refund. |
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
| `amount` | `float64` | No | The amount of money to request. |
| `amountPending` | `float64` | No | The amount of money that was authorised but has not arrived in the account yet. |
| `amountReceived` | `float64` | No | The amount of money that has been received for this payment request. |
| `amountRefunded` | `float64` | No | The amount of money that has been refunded for this payment request. |
| `callbackUrl` | `string` | No |  |
| `cardStripePaymentIntentSecret` | `string` | No |  |
| `countryCode` | `string` | No | The country code associated with the payment. |
| `currency` | `string` | No | The currency of the request. |
| `customFieldsToDisplay` | `[]any` | No | Custom fields to display to the customer. |
| `description` | `string` | No | An optional description for the payment request. |
| `dueDate` | `string` | No | The due date of the payment request. |
| `fieldDisplaySettings` | `[]any` | No |  |
| `googlePayMerchantID` | `string` | No | Merchant ID from Google Pay |
| `id` | `string` | No |  |
| `jwk` | `string` | No | The jwk containing the public key |
| `merchantID` | `string` | No |  |
| `merchantLogoUrlPng` | `string` | No |  |
| `merchantLogoUrlSvg` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `merchantShortName` | `string` | No |  |
| `partialPaymentMethod` | `string` | No |  |
| `paymentAttempts` | `[]any` | No | The payment attempts for this payment request. |
| `paymentMethodsList` | `[]any` | No | The payment methods that the payment request supports. |
| `paymentProcessor` | `string` | No | The card processor |
| `paymentProcessorKey` | `string` | No | The card processors public key |
| `pispError` | `string` | No | This is the error returned from the bank which is recorded in payment request events. |
| `priorityBankID` | `string` | No |  |
| `status` | `string` | No | The status of the payment request. |
| `stripeAccountID` | `string` | No | Account ID of connected customers in Stripe |
| `title` | `string` | No | The title of the payment request. |

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
| `amount` | `float64` | No | The authorised payment amount. |
| `amountPending` | `float64` | No |  |
| `amountReceived` | `float64` | No |  |
| `amountRefunded` | `float64` | No |  |
| `currency` | `string` | No | The authorised payment currency. |
| `customerID` | `string` | No | The customer id |
| `paymentRequestID` | `string` | No | The ID of the payment request the result is for. |
| `payments` | `[]any` | No | The list of payment attempts that have been received for the payment request. |
| `pispAuthorizations` | `[]any` | No |  |
| `requestedAmount` | `float64` | No | The full original payment amount requested. |
| `result` | `string` | No | The result of the payment attempt. |

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
| `accountID` | `string` | No | Gets or Sets Account Id of sending account |
| `allowIncomplete` | `bool` | No | If set to true the payout will get created even if the business validation rules fail. |
| `amount` | `float64` | No | Gets or Sets payout amount |
| `amountMinorUnits` | `int` | No | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `string` | No | This field is used when returning an payout record to a client. |
| `approverID` | `string` | No | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `[]any` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `[]any` | No | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `int` | No | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `int` | No | The number of authorisers required for this payout. |
| `batchPayoutID` | `string` | No | The ID of the batch the payout is associated with. |
| `beneficiary` | `map[string]any` | Yes |  |
| `beneficiaryID` | `string` | No | Optional. |
| `canAuthorise` | `bool` | No | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `bool` | No | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `bool` | No | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `string` | No | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `string` | No |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | No | Gets or Sets Currency of payout request |
| `currentUserID` | `string` | No | The ID of the user that requested access to the PayOut record. |
| `description` | `string` | No | Gets or Sets description of payout request |
| `destination` | `map[string]any` | No |  |
| `documents` | `[]any` | No | Documents associated with the payout. |
| `events` | `[]any` | No | The activity associated with the payout. |
| `failedPayouts` | `map[string]any` | No |  |
| `formattedAmount` | `string` | No | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `string` | No | FX destination currency and amount formatted string. |
| `formattedSchedule` | `string` | No |  |
| `formattedScheduleDayOnly` | `string` | No |  |
| `formattedSourceAccountAvailableBalance` | `string` | No | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `float64` | No | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `int` | No | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `string` | No | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `string` | No | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `string` | No | Optional. |
| `fxRate` | `float64` | No | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `bool` | No | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `bool` | No | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `string` | No | The ID for the payout. |
| `inserted` | `string` | No |  |
| `invoiceID` | `string` | No | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `bool` | No | Indicates whether the payout is archived. |
| `isFailed` | `bool` | No | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `bool` | No | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `bool` | No | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `string` | No |  |
| `nonce` | `string` | No |  |
| `paymentProcessor` | `string` | No | The usptream payment processor for the payout. |
| `paymentRail` | `string` | No | Optional field to indicate the payment rail to use for the payout. |
| `payouts` | `[]any` | No |  |
| `payrunID` | `string` | No | The ID of the payrun that this payout is associated with. |
| `payrunName` | `string` | No | The name of the payrun that this payout is associated with. |
| `reason` | `string` | No |  |
| `rule` | `map[string]any` | No |  |
| `scheduleDate` | `string` | No | The date the payout should be submitted. |
| `scheduled` | `bool` | No | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `float64` | No | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `int` | No | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `string` | No | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `string` | No | The currency of the source account. |
| `sourceAccountIban` | `string` | No | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `map[string]any` | Yes |  |
| `sourceAccountName` | `string` | No | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `string` | No | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `string` | No | The sort code of the account the payout is being made from. |
| `status` | `string` | No | Gets or Sets the status of payout request |
| `tagIds` | `[]any` | No | An optional list of tag ids to add to the payout. |
| `tags` | `[]any` | No | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `string` | No | Gets or Sets destination reference ID |
| `topupPayrunID` | `string` | No | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `float64` | No | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `float64` | No | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `float64` | No | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `string` | No | Gets or Sets payout type |
| `userID` | `string` | No | Gets or Sets User ID of who created the payout request |
| `yourReference` | `string` | No | Gets or Sets your reference ID |

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
    "id": "example_id",
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
| `accountID` | `string` | No | Gets or Sets Account Id of sending account |
| `amount` | `float64` | No | Gets or Sets payout amount |
| `amountMinorUnits` | `int` | No | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `string` | No | This field is used when returning an payout record to a client. |
| `approverID` | `string` | No | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `[]any` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `[]any` | No | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `int` | No | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `int` | No | The number of authorisers required for this payout. |
| `batchPayoutID` | `string` | No | The ID of the batch the payout is associated with. |
| `beneficiary` | `map[string]any` | Yes |  |
| `canAuthorise` | `bool` | No | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `bool` | No | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `bool` | No | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `string` | No | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `string` | No |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | No | Gets or Sets Currency of payout request |
| `currentUserID` | `string` | No | The ID of the user that requested access to the PayOut record. |
| `description` | `string` | No | Gets or Sets description of payout request |
| `destination` | `map[string]any` | No |  |
| `documents` | `[]any` | No | Documents associated with the payout. |
| `events` | `[]any` | No | The activity associated with the payout. |
| `formattedAmount` | `string` | No | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `string` | No | FX destination currency and amount formatted string. |
| `formattedSchedule` | `string` | No |  |
| `formattedScheduleDayOnly` | `string` | No |  |
| `formattedSourceAccountAvailableBalance` | `string` | No | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `float64` | No | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `int` | No | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `string` | No | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `string` | No | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `string` | No | Optional. |
| `fxRate` | `float64` | No | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `bool` | No | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `bool` | No | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `string` | No | The ID for the payout. |
| `inserted` | `string` | No |  |
| `invoiceID` | `string` | No | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `bool` | No | Indicates whether the payout is archived. |
| `isFailed` | `bool` | No | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `bool` | No | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `bool` | No | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `string` | No |  |
| `nonce` | `string` | No |  |
| `paymentProcessor` | `string` | No | The usptream payment processor for the payout. |
| `paymentRail` | `string` | No | Optional field to indicate the payment rail to use for the payout. |
| `payrunID` | `string` | No | The ID of the payrun that this payout is associated with. |
| `payrunName` | `string` | No | The name of the payrun that this payout is associated with. |
| `rule` | `map[string]any` | No |  |
| `scheduleDate` | `string` | No | The date the payout should be submitted. |
| `scheduled` | `bool` | No | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `float64` | No | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `int` | No | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `string` | No | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `string` | No | The currency of the source account. |
| `sourceAccountIban` | `string` | No | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `map[string]any` | Yes |  |
| `sourceAccountName` | `string` | No | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `string` | No | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `string` | No | The sort code of the account the payout is being made from. |
| `status` | `string` | No | Gets or Sets the status of payout request |
| `tags` | `[]any` | No | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `string` | No | Gets or Sets destination reference ID |
| `topupPayrunID` | `string` | No | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `float64` | No | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `float64` | No | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `float64` | No | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `string` | No | Gets or Sets payout type |
| `userID` | `string` | No | Gets or Sets User ID of who created the payout request |
| `yourReference` | `string` | No | Gets or Sets your reference ID |

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
| `authorisations` | `[]any` | No | A list of the users who have successfully authorised the latest version of the payrun and when. |
| `authorisersCompletedCount` | `int` | No | The number of distinct authorisers that have authorised the payrun. |
| `authorisersRequiredCount` | `int` | No | The number of authorisers required for this payrun. |
| `batchPayoutID` | `string` | No |  |
| `canAuthorise` | `bool` | No | True if the payrun can be authorised by the user who loaded it. |
| `canDelete` | `bool` | No |  |
| `canEdit` | `bool` | No |  |
| `events` | `[]any` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No | True if the payrun was loaded for a user and that user has already authorised the latest version of the payrun. |
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
| `accountID` | `string` | No | The ID of the account the rule will apply to. |
| `approveUrl` | `string` | No | If set this property holds the URL an approver needs to visit in order to complete a strong authentication check in order to approve the rule. |
| `approverID` | `string` | No |  |
| `authenticationMethods` | `[]any` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `[]any` | No | A list of the users who have successfully authorised the latest version of the rule and when. |
| `authorisersCompletedCount` | `int` | No | The number of distinct authorisers that have authorised the rule. |
| `authorisersRequiredCount` | `int` | No | The number of authorisers required for this rule. |
| `canAuthorise` | `bool` | No | True if the rule can be authorised by the user who loaded it. |
| `createdBy` | `map[string]any` | Yes |  |
| `description` | `string` | No | Arbitrary description for the rule. |
| `endAt` | `string` | No | Optional end time for rule executions. |
| `hasCurrentUserAuthorised` | `bool` | No | True if the current user has authorised. |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isDisabled` | `bool` | No | If set to true the rule will be disabled from executing. |
| `lastExecutedAt` | `string` | No |  |
| `lastRunAtTransactionDate` | `string` | No | The most recent transaction date when the rule was last run. |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No | The ID of the merchant that owns the account. |
| `name` | `string` | No | A name to succinctly describe the rule. |
| `nonce` | `string` | Yes |  |
| `onApprovedWebHookUrl` | `string` | No | Optional URL to receive an HTTP request with the rule details when the rule status changes to approved. |
| `onExecutionErrorWebHookUrl` | `string` | No | Optional URL to receive an HTTP request when a rule execution attempt fails. |
| `onExecutionSuccessWebHookUrl` | `string` | No | Optional URL to receive an HTTP request when a rule execution attempt succeeds. |
| `startAt` | `string` | No | Optional start time for rule executions. |
| `status` | `string` | No |  |
| `sweepAction` | `map[string]any` | No |  |
| `timeZoneId` | `string` | No | If the rule should be executed on a recurring schedule this is the timezone that the CRON expression should be evaluated in. |
| `triggerCronExpression` | `string` | No | If the rule should be executed on a recurring schedule this is the expression that sets the schedule. |
| `triggerOnPayIn` | `bool` | No | Set to true if the rule execution should be triggered when the account receives a pay in (credit). |
| `userID` | `string` | No |  |
| `webHookSecret` | `string` | No | If set this secret will be used to sign Web Hook requests. |

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
| `accountID` | `string` | No | The ID of the account the transaction belongs to. |
| `accountName` | `string` | No | The name of the account the transaction belongs to. |
| `accountSequenceNumber` | `int` | No | The sequence number of transaction on a per account basis. |
| `addressDetails` | `map[string]any` | No |  |
| `amount` | `float64` | No | Amount of the transaction. |
| `amountMinorUnits` | `int` | No | Amount of the transaction expressed in the currency’s minor units (e.g. |
| `balance` | `float64` | No | Balance left on the account after the transaction. |
| `balanceMinorUnits` | `int` | No | Balance on the account expressed in the currency’s minor units (e.g. |
| `bookingDateTime` | `string` | No |  |
| `chargeDetails` | `map[string]any` | No |  |
| `content` | `[]any` | No |  |
| `counterparty` | `map[string]any` | No |  |
| `counterpartySummary` | `string` | No | For pay in (credit) transactions this will contain a descriptive string with the most important fields about the counterparty. |
| `currency` | `string` | No | Currency of transaction. |
| `currencyExchange` | `map[string]any` | No | Provides details on the currency exchange. |
| `date` | `string` | No |  |
| `description` | `string` | No | Description of the transaction. |
| `enrichment` | `map[string]any` | No |  |
| `fxAmount` | `float64` | No | For an FX payout this is the amound in the FX currency. |
| `fxCurrency` | `string` | No | For an FX payout this is the currency that was received or that was instructed. |
| `fxRate` | `float64` | No | For an FX payout this is the exchange rate between the transaction currency and the FX currency. |
| `grossAmount` | `map[string]any` | Yes |  |
| `id` | `string` | No | Unique ID for the transaction. |
| `inserted` | `string` | No | Date when the transaction was inserted into the ledger. |
| `isoBankTransactionCode` | `map[string]any` | No |  |
| `merchant` | `map[string]any` | No |  |
| `merchantID` | `string` | No | The ID of the merchant that owns the account. |
| `pageNumber` | `int` | No | Current page number. |
| `pageSize` | `int` | No | Page size |
| `payeeDetails` | `map[string]any` | Yes | The Payee object contains details of the beneficiary, person or business. |
| `payerDetails` | `map[string]any` | Yes |  |
| `paymentRequestCustomFields` | `map[string]any` | No | The custom fields that were attached to the payment request that resulted in this transaction. |
| `paymentRequestID` | `string` | No | For Pay by Bank and Direct Debit transactions this will contain the ID of the payment request. |
| `payoutID` | `string` | No | ID of the payout that resulted in the transaction. |
| `proprietaryBankTransactionCode` | `map[string]any` | No |  |
| `rawReference` | `string` | No | The raw payment reference details as received from the payment processor. |
| `reference` | `string` | No |  |
| `ruleID` | `string` | No | ID of the rule that resulted in the transaction. |
| `statementReferences` | `[]any` | No |  |
| `status` | `string` | No |  |
| `supplementaryData` | `any` | No |  |
| `tags` | `[]any` | No | An optional list of descriptive tags attached to the transaction. |
| `theirReference` | `string` | No | For a pay out the reference that the payer attached for the receiving party. |
| `totalPages` | `int` | No | Total pages |
| `totalSize` | `int` | No | Total count |
| `transactionAmount` | `map[string]any` | Yes |  |
| `transactionDate` | `string` | No | Date when the transaction occurred. |
| `transactionInformation` | `[]any` | No |  |
| `transactionMutability` | `string` | No |  |
| `type` | `string` | No | Type of the transaction. |
| `valueDateTime` | `string` | No |  |
| `virtualIBAN` | `string` | No | If set it indicates the payin was to a virtual IBAN. |
| `yourReference` | `string` | No | For a pay in the reference the sending party attached. |

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
| `clientSessionTimeouts` | `[]any` | No | The number of seconds a session for this user should last before expiring. |
| `emailAddress` | `string` | Yes |  |
| `firstName` | `string` | Yes |  |
| `id` | `string` | No |  |
| `lastName` | `string` | Yes |  |
| `passkeyAdded` | `bool` | No |  |
| `permissions` | `map[string]any` | No |  |
| `profile` | `string` | No |  |
| `rolesWithScope` | `[]any` | No |  |
| `twoFactorEnabled` | `bool` | No |  |
| `userInviteID` | `string` | No | Optional ID of the invite that is being accepted so the user can be assigned a role on a new merchant. |

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
| `initialRoleID` | `string` | No | The role ID to automatically assign to the merchant’s very first user. |
| `inviteeEmailAddress` | `string` | No | Email address of the user being invited. |
| `inviteeFirstName` | `string` | No | First Name of the user being invited. |
| `inviteeLastName` | `string` | No | Last Name of the user being invited. |
| `inviterEmailAddress` | `string` | No |  |
| `inviterFirstName` | `string` | No |  |
| `inviterLastName` | `string` | No |  |
| `isAuthorised` | `bool` | No | Will be set to true once the invite has met the authorisation requirements. |
| `isInviteeRegistered` | `bool` | No | If true, indicates the invitee's email address corresponds to an existing MoneyMoov user. |
| `lastInvited` | `string` | No |  |
| `merchantID` | `string` | No | ID of the merchant the user is being invited to. |
| `merchantName` | `string` | No |  |
| `message` | `string` | No |  |
| `registrationUrl` | `string` | No |  |
| `sendInviteEmail` | `bool` | No | If set to true an email will be sent to the invitee with instructions on how to accept the invite. |
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
    "id": "example_id",
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
| `accountName` | `string` | No | Name for the account |
| `accountSupplierName` | `string` | No | The payment account supplier name. |
| `availableBalance` | `float64` | No | The current available balance of the account. |
| `availableBalanceMinorUnits` | `int` | No | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `float64` | No | Balance of the account. |
| `balanceMinorUnits` | `int` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `string` | No | The bank name for external accounts |
| `consentID` | `string` | No | The ID of the consent used to connect the external account. |
| `createdBy` | `map[string]any` | Yes |  |
| `createdByDisplayName` | `string` | No | Either the name of the user, merchant token or api key that created the account |
| `currency` | `string` | No | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `string` | No | Indicates the default payment rail for this account. |
| `displayName` | `string` | No | Gets a unique display name for the payment account. |
| `expiryDate` | `string` | No | The date that the external account will expire |
| `externalAccountIcon` | `string` | No | The Icon for external accounts |
| `id` | `string` | No | Unique id for the account. |
| `identifier` | `map[string]any` | Yes |  |
| `inserted` | `string` | No | Timestamp when the account was created. |
| `isArchived` | `bool` | No | Indicates whether the account is archived. |
| `isConnectedAccount` | `bool` | No | Indicates if the payment account is an externally connected account. |
| `isDefault` | `bool` | No | Is the default account |
| `isTrustAccount` | `bool` | No | Indicates if the payment account is a trust account. |
| `isVirtual` | `bool` | No | True if the account is a virtual account. |
| `lastTransaction` | `map[string]any` | No |  |
| `lastUpdated` | `string` | No | Timestamp when the account was last updated. |
| `merchantID` | `string` | No | The ID of the merchant that owns the account. |
| `merchantName` | `string` | No | The name of the merchant that owns the account. |
| `name` | `string` | Yes | The name of the virtual account. |
| `physicalAccountID` | `string` | No | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `[]any` | No | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `float64` | No | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `int` | No | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `string` | No | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `string` | No | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `string` | No | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `string` | No |  |
| `xeroBankFeedSyncStatus` | `string` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | No | Indicates the number of unsynchronised transactions with Xero |

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
| `destinationUrl` | `string` | No | The destination URL for the webhook. |
| `emailAddress` | `string` | No | The recipient email address(es) for notifications. |
| `failedNotificationEmailAddress` | `string` | No | The email address to which notifications about failed webhook deliveries will be sent. |
| `id` | `string` | No |  |
| `isActive` | `bool` | No |  |
| `merchantID` | `string` | No | The ID of the merchant that the webhook is for. |
| `notificationMethod` | `string` | No | The type of notification that will be sent. |
| `resourceTypes` | `[]any` | No | The resource types that the webhook will be generated for. |
| `retry` | `bool` | No |  |
| `secret` | `string` | No | The secret key required to authenticate webhook notifications. |
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

