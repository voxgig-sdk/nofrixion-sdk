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
| `accountBalances` | `table` | No | The various balances for the account. |
| `accountID` | `string` | No | ID of the account. |
| `accountIdentifications` | `table` | No | The canoncial identifiers for the account. |
| `accountName` | `string` | No | Name for the account |
| `accountNames` | `table` | No | Optional account names set by the account holder. |
| `accountSupplierName` | `string` | No | The payment account supplier name. |
| `accountType` | `string` | No | The type of account e.g. |
| `availableBalance` | `number` | No | The current available balance of the account. |
| `availableBalanceMinorUnits` | `number` | No | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `number` | No | Balance of the account. |
| `balanceMinorUnits` | `number` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `string` | No | The bank name for external accounts |
| `consentID` | `string` | No | The ID of the consent used to connect the external account. |
| `consolidatedAccountInformation` | `table` | No | Summary information regarding account balances of the overall account provided by the bank. |
| `createdBy` | `table` | Yes |  |
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
| `identifier` | `table` | Yes |  |
| `inserted` | `string` | No | Timestamp when the account was created. |
| `isArchived` | `boolean` | No | Indicates whether the account is archived. |
| `isConnectedAccount` | `boolean` | No | Indicates if the payment account is an externally connected account. |
| `isDefault` | `boolean` | No | Is the default account |
| `isTrustAccount` | `boolean` | No | Indicates if the payment account is a trust account. |
| `isVirtual` | `boolean` | No | True if the account is a virtual account. |
| `lastTransaction` | `table` | No |  |
| `lastUpdated` | `string` | No | Timestamp when the account was last updated. |
| `merchantID` | `string` | No | The ID of the merchant that owns the account. |
| `merchantName` | `string` | No | The name of the merchant that owns the account. |
| `nickname` | `string` | No | Nickname of the account that was provided by the account owner. |
| `physicalAccountID` | `string` | No | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `roleIDs` | `table` | No | Optional list of role IDs that will get access to the payment account when created. |
| `rules` | `table` | No | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `number` | No | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `number` | No | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
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
| `xeroUnsynchronisedTransactionsCount` | `number` | No | Indicates the number of unsynchronised transactions with Xero |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Account():create({
  account_id = --[[ string ]],
  currency = --[[ string ]],
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
| `approveUrl` | `string` | No | This field is used when returning a batch payout record to a client. |
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
| `authenticationMethods` | `table` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `table` | No | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `number` | No | The number of distinct authorisers that have authorised the beneficiary. |
| `authorisersRequiredCount` | `number` | No | The number of authorisers required for this beneficiary. |
| `beneficiaries` | `table` | No |  |
| `beneficiaryEvents` | `table` | No |  |
| `canAuthorise` | `boolean` | No | True if the beneficiary can be authorised by the user who loaded it. |
| `canUpdate` | `boolean` | No | True if the beneficiary can be updated by the user who loaded it. |
| `createdBy` | `table` | Yes |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | Yes | Gets or Sets the currency. |
| `destination` | `table` | No |  |
| `failedBeneficiaries` | `table` | No |  |
| `hasCurrentUserAuthorised` | `boolean` | No | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isEnabled` | `boolean` | No |  |
| `lastAuthorised` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No | Gets or Sets the merchant id. |
| `name` | `string` | Yes | The descriptive name for the beneficiary. |
| `nonce` | `string` | No |  |
| `sourceAccountIDs` | `table` | No | ID of the accounts which are authorised to act as a source for the beneficiary. |
| `sourceAccounts` | `table` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Beneficiary():create({
  id = --[[ string ]],
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
| `groupMembers` | `table` | No | The existing group members. |
| `groupName` | `string` | Yes | The descriptive name for the beneficiary group. |
| `id` | `string` | No |  |
| `inserted` | `string` | No | Timestamp indicating when the group was created. |
| `lastUpdated` | `string` | No | Timestamp indicating when the group was last updated. |
| `merchantID` | `string` | Yes | Gets or Sets the merchant id. |

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
| `isPayerAuthenticationRequired` | `boolean` | No | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `boolean` | No | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `string` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `string` | No | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `string` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `number` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `number` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `string` | No | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `string` | No |  |
| `requestID` | `string` | No |  |
| `responseCode` | `string` | No |  |
| `responseType` | `string` | No |  |
| `status` | `string` | No |  |
| `threeDSRedirectUrl` | `string` | No | Checkout.com require a redirect for 3DS authentication. |
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
local result, err = client:CardCustomerToken():remove({ id = "id" })
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
| `isPayerAuthenticationRequired` | `boolean` | No | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `boolean` | No | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `string` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `string` | No | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `string` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `number` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `number` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `string` | No | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `string` | No |  |
| `requestID` | `string` | No |  |
| `responseCode` | `string` | No |  |
| `responseType` | `string` | No |  |
| `status` | `string` | No |  |
| `threeDSRedirectUrl` | `string` | No | Checkout.com require a redirect for 3DS authentication. |
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
| `authorisationUrl` | `string` | No | The URL the authorising user needs to be redirected to in order to get the open banking consent token. |
| `callbackUrl` | `string` | No | Optional callback URL that the end user performing the open banking authorisation will be redirected to on completion. |
| `consentID` | `string` | No | The ID of the open banking consent. |
| `emailAddress` | `string` | No | The email address that identifies the end user that will be authorising the open banking consent request. |
| `expiryDate` | `string` | No |  |
| `failureCallbackUrl` | `string` | No | Optional callback URL for open banking consent authorisation failure. |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `institutionID` | `string` | No | The institution ID the open banking consent is being requested for. |
| `isConnectedAccounts` | `boolean` | No | Optional setting. |
| `isEnabled` | `boolean` | No |  |
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
| `failedSubmissions` | `table` | No | Dictionary of failed submissions, keyed by the index (1-based) in the original request. |
| `successfulSubmissions` | `table` | No | List of successfully submitted direct debit payments. |

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
| `exchangeRate` | `number` | No | The price at which the transaction will buy the source currency using the destination currency. |
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
| `isRecurring` | `boolean` | No | Whether this mandate is single-use or recurring. |
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
| `accountCurrencies` | `table` | No | The list of currencies that the merchant has accounts for. |
| `canHaveTrustAccounts` | `boolean` | No | Trust accounts are a special type of account that allow the account name to be trusted for use in statements and verification of payee checks. |
| `cardPaymentProcessor` | `string` | No | Name of the card payment processor. |
| `companyID` | `string` | No | The Company ID recorded in the Compliance system. |
| `displayQrOnHostedPay` | `boolean` | No | Indicates if a QR Code containing the payment link should be displayed on the hosted payment page. |
| `hostedPayVersion` | `number` | No | The version of the hosted payment page to use with the merchant. |
| `id` | `string` | No | Unique ID for the merchant. |
| `inserted` | `string` | No | Timestamp the merchant was added to MoneyMoov. |
| `isBlocked` | `boolean` | No | The merchant is blocked from making payments (payouts). |
| `isExited` | `boolean` | No | The merchant has formally terminated their relationship and is no longer a customer. |
| `isSuspended` | `boolean` | No | The merchant has temporarily suspended their own account. |
| `jurisdiction` | `string` | No | The jurisdiction the merchant entity is incorporated or established in. |
| `logoUrlPng` | `string` | No | The CDN URL of the merchant's logo in PNG format. |
| `logoUrlSvg` | `string` | No | The CDN URL of the merchant's logo in SVG format. |
| `merchantCategoryCode` | `string` | No | The industry code that represents the merchant's primary trading activity. |
| `name` | `string` | No | The registered business name of the merchant. |
| `notes` | `string` | No | The notes field is an optional free text field that can be used to store any additional information about the merchant. |
| `parentMerchant` | `table` | No |  |
| `paymentAccountLimit` | `number` | No | The maximum number of payment accounts that can be created for the Merchant. |
| `paymentAccounts` | `table` | No |  |
| `reason` | `string` | No | The reason for the suspension. |
| `shortName` | `string` | No | A URL friendly shortish name for the merchant. |
| `supportedPaymentMethodsList` | `table` | No | The payment methods that are configured and supported for this merchant. |
| `suspensionReason` | `string` | No | The reason for the suspension, provided by the merchant. |
| `tags` | `table` | No | An optional list of descriptive tags that can be used on merchant entities such as payment requests. |
| `timeZoneId` | `string` | No | The IANA (Internet Assigned Numbers Authority) time zone identifier of the merchant. |
| `tradingName` | `string` | No | An optional trading name. |
| `webHookLimit` | `number` | No | The maximum number of web hooks that can be created for the Merchant. |
| `yourRoleName` | `string` | No | The name of the role for the identity that loaded the merchant record. |

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
local result, err = client:Merchant():remove({ id = "merchant_id", user_id = "user_id" })
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
| `isRecurring` | `boolean` | No | Whether this mandate is single-use or recurring. |
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
| `bankCountryCodes` | `table` | No | The list of country codes representing the banks the country supports. |
| `bankID` | `string` | No | ID of the bank to be configured for the merchant. |
| `bankName` | `string` | No | Name of the Bank/Institution. |
| `businessInstitutionID` | `string` | No | ID that the processor uses to identify the bank (business accounts). |
| `currency` | `string` | No | Currency supported by the bank. |
| `logo` | `string` | No | URL of the bank's logo. |
| `message` | `string` | No | Message relating to specific bank. |
| `messageImageUrl` | `string` | No | Optional image URL to be displayed with the message. |
| `order` | `number` | No | Order in which this setting will appear in the UI. |
| `personalInstitutionID` | `string` | No | ID that the processor uses to identify the bank (personal accounts). |
| `processor` | `string` | No | Name of the bank payment processor. |
| `warningHeading` | `string` | No | The heading for a warning message related to the bank institution to be displayed to the user. |
| `warningMessage` | `string` | No | The warning message related to the bank institution to be displayed to the user. |

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
| `customFields` | `table` | No | A list of custom fields that can be included in the payment request template. |
| `defaultFields` | `table` | No | A list of default fields that are included in the payment request template. |
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
| `authenticationMethods` | `table` | No | A list of authentication types allowed to authorise the merchant token. |
| `authorisations` | `table` | No | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `number` | No | The number of distinct authorisers that have authorised the merchant token. |
| `authorisersRequiredCount` | `number` | No | The number of authorisers required for this merchant token. |
| `canAuthorise` | `boolean` | No | True if the merchant token can be authorised by the user who loaded it. |
| `description` | `string` | No | Token description |
| `expiresAt` | `string` | No | Optional. |
| `hasCurrentUserAuthorised` | `boolean` | No | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `hmacAlgorithm` | `string` | No | Optional shared secret algorithm to use for HMAC authentication. |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `ipAddressWhitelist` | `string` | No | Optional. |
| `isArchived` | `boolean` | No | Indicates whether the merchant token is archived. |
| `isEnabled` | `boolean` | No | If set to false the merchant token will not be accepted to authorise a request. |
| `lastAuthorised` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No | The merchant id to add to the token |
| `nonce` | `string` | Yes |  |
| `permissionTypes` | `table` | No | The permissions that the merchant token supports. |
| `requestSignatureVersion` | `number` | No | Represent the version of the overall merchant token. |
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
local result, err = client:OpenBanking():remove({ account_id = "account_id" })
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
| `accountName` | `string` | Yes | The name of the account to verify |
| `accountNumber` | `string` | No | The account number of the account to verify (for CoP checks) |
| `iban` | `string` | Yes | The IBAN of the account to verify (for VoP checks) |
| `payeeVerifiedAccountName` | `string` | No | The verified account name of the payee, if available (in case of a close match) |
| `result` | `string` | No | The result of the payee verification |
| `secondaryIdentification` | `string` | No | Optional secondary identifier for the account to verify. |
| `sortCode` | `string` | No | The sort code of the account to verify (for CoP checks) |

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
| `amount` | `number` | No | The amount of money to request. |
| `amountPending` | `number` | No | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `number` | No | Total amount received for this payment request. |
| `amountRefunded` | `number` | No | Total amount refunded for this payment request. |
| `autoSendReceipt` | `boolean` | No | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `string` | No | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `string` | No | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `boolean` | No | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `boolean` | No | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `string` | No | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `boolean` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardNoPayerAuthentication` | `boolean` | No | If set to true for card payments no attempt will be made to use payer authentication (3-D Secure and equivalent). |
| `cardProcessorMerchantID` | `string` | No | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `string` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `string` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `cardTransmitRawDetails` | `boolean` | No | If set to true for card payments the sensitive card number and card verification number will be transmitted directly rather than being tokenised. |
| `createdByUser` | `table` | Yes |  |
| `currency` | `string` | No | The currency of the request. |
| `customFields` | `table` | No | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `string` | No | Optional email address for the customer. |
| `customerID` | `string` | No | An optional customer identifier for the payment request. |
| `customerName` | `string` | No |  |
| `description` | `string` | No | An optional description for the payment request. |
| `destinationAccount` | `table` | No |  |
| `directDebitPayment` | `table` | No | Contains information about a Direct Debit payment attempt for a payment request. |
| `dueDate` | `string` | No | The due date for the payment request. |
| `events` | `table` | No |  |
| `failureCallbackUrl` | `string` | No | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `table` | No | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `string` | No |  |
| `hostedPayCheckoutUrl` | `string` | No | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `string` | No |  |
| `ignoreAddressVerification` | `boolean` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `string` | No | The timestamp the payment request was created at. |
| `insertedSortable` | `string` | No | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `isArchived` | `boolean` | No | Indicates whether the payment request is archived. |
| `jwk` | `string` | No | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `string` | No | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `string` | No | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `string` | No | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `string` | No | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `string` | No | The ID of the merchant to create the payment request for. |
| `merchantTokenDescription` | `string` | No | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `string` | No |  |
| `notificationRoleIDs` | `table` | No | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `string` | No | An optional order ID for the payment request. |
| `partialPaymentMethod` | `string` | No | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `string` | No | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `table` | No | The payment attempts made against this payment request. |
| `paymentMethods` | `table` | No | The payment methods that the payment request supports. |
| `paymentProcessor` | `string` | No | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `payrunID` | `string` | No | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `string` | No | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `string` | No | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `table` | No |  |
| `sandboxSettleDelayInSeconds` | `number` | No | Sandbox only. |
| `shippingAddress` | `table` | No |  |
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
| `tagIds` | `table` | No | An optional list of tag ids to add to the payment request |
| `tags` | `table` | No | An optional list of descriptive tags attached to the payment request. |
| `title` | `string` | No | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `table` | No |  |
| `transactions` | `table` | No |  |
| `useHostedPaymentPage` | `boolean` | No | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

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
| `accountName` | `string` | No | Name for the account |
| `accountSupplierName` | `string` | No | The payment account supplier name. |
| `availableBalance` | `number` | No | The current available balance of the account. |
| `availableBalanceMinorUnits` | `number` | No | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `number` | No | Balance of the account. |
| `balanceMinorUnits` | `number` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `string` | No | The bank name for external accounts |
| `consentID` | `string` | No | The ID of the consent used to connect the external account. |
| `createdBy` | `table` | Yes |  |
| `createdByDisplayName` | `string` | No | Either the name of the user, merchant token or api key that created the account |
| `currency` | `string` | No | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `string` | No | Indicates the default payment rail for this account. |
| `displayName` | `string` | No | Gets a unique display name for the payment account. |
| `expiryDate` | `string` | No | The date that the external account will expire |
| `externalAccountIcon` | `string` | No | The Icon for external accounts |
| `id` | `string` | No | Unique id for the account. |
| `identifier` | `table` | Yes |  |
| `inserted` | `string` | No | Timestamp when the account was created. |
| `isArchived` | `boolean` | No | Indicates whether the account is archived. |
| `isConnectedAccount` | `boolean` | No | Indicates if the payment account is an externally connected account. |
| `isDefault` | `boolean` | No | Is the default account |
| `isTrustAccount` | `boolean` | No | Indicates if the payment account is a trust account. |
| `isVirtual` | `boolean` | No | True if the account is a virtual account. |
| `lastTransaction` | `table` | No |  |
| `lastUpdated` | `string` | No | Timestamp when the account was last updated. |
| `merchantID` | `string` | No | The ID of the merchant that owns the account. |
| `merchantName` | `string` | No | The name of the merchant that owns the account. |
| `physicalAccountID` | `string` | No | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `table` | No | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `number` | No | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `number` | No | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `string` | No | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `string` | No | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `string` | No | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `string` | No |  |
| `xeroBankFeedSyncStatus` | `string` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `number` | No | Indicates the number of unsynchronised transactions with Xero |

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
| `accountName` | `string` | No | Name for the account |
| `availableBalance` | `number` | No | The current available balance of the account. |
| `balance` | `number` | No | Balance of the account. |
| `balanceMinorUnits` | `number` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `currency` | `string` | No | Currency of the account in ISO 4217 format |
| `id` | `string` | No | Unique id for the account. |
| `identifier` | `table` | Yes |  |
| `isArchived` | `boolean` | No | Is the account archived |
| `isConnectedAccount` | `boolean` | No | Indicates if the payment account is an externally connected account. |
| `merchantID` | `string` | No | The ID of the merchant that owns the account. |
| `submittedPayoutsBalance` | `number` | No | Total of the payouts that have been submitted for processing. |

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
| `paymentInitiationID` | `string` | No | The unique identifier of the payment initiation request. |
| `paymentRequestCallbackUrl` | `string` | No | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `string` | No |  |
| `redirectUrl` | `string` | No | A redirect URL for the user to authorise the payment initiation request at the ASPSP |
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
| `amount` | `number` | No | The amount of money to request. |
| `amountPending` | `number` | No | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `number` | No | Total amount received for this payment request. |
| `amountRefunded` | `number` | No | Total amount refunded for this payment request. |
| `autoSendReceipt` | `boolean` | No | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `string` | No | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `string` | No | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `boolean` | No | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `boolean` | No | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `string` | No | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `boolean` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardProcessorMerchantID` | `string` | No | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `string` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `string` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `createdByUser` | `table` | Yes |  |
| `currency` | `string` | No | The currency of the request. |
| `customFields` | `table` | No | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `string` | No | Optional email address for the customer. |
| `customerID` | `string` | No | An optional customer identifier for the payment request. |
| `customerName` | `string` | No |  |
| `description` | `string` | No | An optional description for the payment request. |
| `destinationAccount` | `table` | No |  |
| `directDebitPayment` | `table` | No | Contains information about a Direct Debit payment attempt for a payment request. |
| `doSimulateSettlementFailure` | `boolean` | No |  |
| `dueDate` | `string` | No | The due date for the payment request. |
| `errorDescription` | `string` | No |  |
| `events` | `table` | No |  |
| `failedPaymentRequests` | `table` | No |  |
| `failureCallbackUrl` | `string` | No | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `table` | No | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `string` | No |  |
| `hostedPayCheckoutUrl` | `string` | No | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `string` | No |  |
| `ignoreAddressVerification` | `boolean` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `string` | No | The timestamp the payment request was created at. |
| `insertedSortable` | `string` | No | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `institution` | `string` | No |  |
| `isArchived` | `boolean` | No | Indicates whether the payment request is archived. |
| `jwk` | `string` | No | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `string` | No | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `string` | No | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `string` | No | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `string` | No | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `string` | No |  |
| `merchantTokenDescription` | `string` | No | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `string` | No |  |
| `notificationRoleIDs` | `table` | No | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `string` | No | An optional order ID for the payment request. |
| `partialPaymentMethod` | `string` | No | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `string` | No | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `table` | No | The payment attempts made against this payment request. |
| `paymentInitiationID` | `string` | No |  |
| `paymentMethods` | `table` | No | The payment methods that the payment request supports. |
| `paymentProcessor` | `string` | No | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `paymentRequests` | `table` | No |  |
| `payrunID` | `string` | No | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `string` | No | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `string` | No | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `table` | No |  |
| `sandboxSettleDelayInSeconds` | `number` | No | Sandbox only. |
| `shippingAddress` | `table` | No |  |
| `status` | `string` | No | The current status of the payment request. |
| `successWebHookUrl` | `string` | No | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tags` | `table` | No | An optional list of descriptive tags attached to the payment request. |
| `title` | `string` | No | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `table` | No |  |
| `transactions` | `table` | No |  |
| `useHostedPaymentPage` | `boolean` | No | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

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
local result, err = client:PaymentRequest():load({ id = "payment_request_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:PaymentRequest():remove({ id = "payment_request_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:PaymentRequest():update({
  id = "payment_request_id",
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
| `applePayTransactionID` | `string` | No | Transaction ID received in Apple pay token. |
| `cardAuthorizationResponseID` | `string` | No | For a successful card authorization this field will hold the response ID. |
| `cardExpiryMonth` | `number` | No | For card payment events this field holds the payer's card expiry month. |
| `cardExpiryYear` | `number` | No | For card payment events this field holds the payer's card expiry year. |
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
| `amount` | `number` | No | The amount of money to request. |
| `amountPending` | `number` | No | The amount of money that was authorised but has not arrived in the account yet. |
| `amountReceived` | `number` | No | The amount of money that has been received for this payment request. |
| `amountRefunded` | `number` | No | The amount of money that has been refunded for this payment request. |
| `callbackUrl` | `string` | No |  |
| `cardStripePaymentIntentSecret` | `string` | No |  |
| `countryCode` | `string` | No | The country code associated with the payment. |
| `currency` | `string` | No | The currency of the request. |
| `customFieldsToDisplay` | `table` | No | Custom fields to display to the customer. |
| `description` | `string` | No | An optional description for the payment request. |
| `dueDate` | `string` | No | The due date of the payment request. |
| `fieldDisplaySettings` | `table` | No |  |
| `googlePayMerchantID` | `string` | No | Merchant ID from Google Pay |
| `id` | `string` | No |  |
| `jwk` | `string` | No | The jwk containing the public key |
| `merchantID` | `string` | No |  |
| `merchantLogoUrlPng` | `string` | No |  |
| `merchantLogoUrlSvg` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `merchantShortName` | `string` | No |  |
| `partialPaymentMethod` | `string` | No |  |
| `paymentAttempts` | `table` | No | The payment attempts for this payment request. |
| `paymentMethodsList` | `table` | No | The payment methods that the payment request supports. |
| `paymentProcessor` | `string` | No | The card processor |
| `paymentProcessorKey` | `string` | No | The card processors public key |
| `pispError` | `string` | No | This is the error returned from the bank which is recorded in payment request events. |
| `priorityBankID` | `string` | No |  |
| `status` | `string` | No | The status of the payment request. |
| `stripeAccountID` | `string` | No | Account ID of connected customers in Stripe |
| `title` | `string` | No | The title of the payment request. |

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
| `amount` | `number` | No | The authorised payment amount. |
| `amountPending` | `number` | No |  |
| `amountReceived` | `number` | No |  |
| `amountRefunded` | `number` | No |  |
| `currency` | `string` | No | The authorised payment currency. |
| `customerID` | `string` | No | The customer id |
| `paymentRequestID` | `string` | No | The ID of the payment request the result is for. |
| `payments` | `table` | No | The list of payment attempts that have been received for the payment request. |
| `pispAuthorizations` | `table` | No |  |
| `requestedAmount` | `number` | No | The full original payment amount requested. |
| `result` | `string` | No | The result of the payment attempt. |

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
| `accountID` | `string` | No | Gets or Sets Account Id of sending account |
| `allowIncomplete` | `boolean` | No | If set to true the payout will get created even if the business validation rules fail. |
| `amount` | `number` | No | Gets or Sets payout amount |
| `amountMinorUnits` | `number` | No | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `string` | No | This field is used when returning an payout record to a client. |
| `approverID` | `string` | No | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `table` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `table` | No | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `number` | No | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `number` | No | The number of authorisers required for this payout. |
| `batchPayoutID` | `string` | No | The ID of the batch the payout is associated with. |
| `beneficiary` | `table` | Yes |  |
| `beneficiaryID` | `string` | No | Optional. |
| `canAuthorise` | `boolean` | No | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `boolean` | No | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `boolean` | No | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `string` | No | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `string` | No |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | No | Gets or Sets Currency of payout request |
| `currentUserID` | `string` | No | The ID of the user that requested access to the PayOut record. |
| `description` | `string` | No | Gets or Sets description of payout request |
| `destination` | `table` | No |  |
| `documents` | `table` | No | Documents associated with the payout. |
| `events` | `table` | No | The activity associated with the payout. |
| `failedPayouts` | `table` | No |  |
| `formattedAmount` | `string` | No | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `string` | No | FX destination currency and amount formatted string. |
| `formattedSchedule` | `string` | No |  |
| `formattedScheduleDayOnly` | `string` | No |  |
| `formattedSourceAccountAvailableBalance` | `string` | No | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `number` | No | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `number` | No | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `string` | No | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `string` | No | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `string` | No | Optional. |
| `fxRate` | `number` | No | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `boolean` | No | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `boolean` | No | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `string` | No | The ID for the payout. |
| `inserted` | `string` | No |  |
| `invoiceID` | `string` | No | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `boolean` | No | Indicates whether the payout is archived. |
| `isFailed` | `boolean` | No | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `boolean` | No | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `boolean` | No | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `string` | No |  |
| `nonce` | `string` | No |  |
| `paymentProcessor` | `string` | No | The usptream payment processor for the payout. |
| `paymentRail` | `string` | No | Optional field to indicate the payment rail to use for the payout. |
| `payouts` | `table` | No |  |
| `payrunID` | `string` | No | The ID of the payrun that this payout is associated with. |
| `payrunName` | `string` | No | The name of the payrun that this payout is associated with. |
| `reason` | `string` | No |  |
| `rule` | `table` | No |  |
| `scheduleDate` | `string` | No | The date the payout should be submitted. |
| `scheduled` | `boolean` | No | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `number` | No | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `number` | No | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `string` | No | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `string` | No | The currency of the source account. |
| `sourceAccountIban` | `string` | No | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `table` | Yes |  |
| `sourceAccountName` | `string` | No | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `string` | No | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `string` | No | The sort code of the account the payout is being made from. |
| `status` | `string` | No | Gets or Sets the status of payout request |
| `tagIds` | `table` | No | An optional list of tag ids to add to the payout. |
| `tags` | `table` | No | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `string` | No | Gets or Sets destination reference ID |
| `topupPayrunID` | `string` | No | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `number` | No | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `number` | No | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `number` | No | The actual FX rate applied during settlement, as recorded on the associated transaction. |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Payout():create({
  id = --[[ string ]],
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
| `accountID` | `string` | No | Gets or Sets Account Id of sending account |
| `amount` | `number` | No | Gets or Sets payout amount |
| `amountMinorUnits` | `number` | No | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `string` | No | This field is used when returning an payout record to a client. |
| `approverID` | `string` | No | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `table` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `table` | No | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `number` | No | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `number` | No | The number of authorisers required for this payout. |
| `batchPayoutID` | `string` | No | The ID of the batch the payout is associated with. |
| `beneficiary` | `table` | Yes |  |
| `canAuthorise` | `boolean` | No | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `boolean` | No | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `boolean` | No | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `string` | No | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `string` | No |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | No | Gets or Sets Currency of payout request |
| `currentUserID` | `string` | No | The ID of the user that requested access to the PayOut record. |
| `description` | `string` | No | Gets or Sets description of payout request |
| `destination` | `table` | No |  |
| `documents` | `table` | No | Documents associated with the payout. |
| `events` | `table` | No | The activity associated with the payout. |
| `formattedAmount` | `string` | No | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `string` | No | FX destination currency and amount formatted string. |
| `formattedSchedule` | `string` | No |  |
| `formattedScheduleDayOnly` | `string` | No |  |
| `formattedSourceAccountAvailableBalance` | `string` | No | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `number` | No | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `number` | No | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `string` | No | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `string` | No | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `string` | No | Optional. |
| `fxRate` | `number` | No | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `boolean` | No | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `boolean` | No | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `string` | No | The ID for the payout. |
| `inserted` | `string` | No |  |
| `invoiceID` | `string` | No | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `boolean` | No | Indicates whether the payout is archived. |
| `isFailed` | `boolean` | No | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `boolean` | No | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `boolean` | No | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `string` | No |  |
| `nonce` | `string` | No |  |
| `paymentProcessor` | `string` | No | The usptream payment processor for the payout. |
| `paymentRail` | `string` | No | Optional field to indicate the payment rail to use for the payout. |
| `payrunID` | `string` | No | The ID of the payrun that this payout is associated with. |
| `payrunName` | `string` | No | The name of the payrun that this payout is associated with. |
| `rule` | `table` | No |  |
| `scheduleDate` | `string` | No | The date the payout should be submitted. |
| `scheduled` | `boolean` | No | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `number` | No | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `number` | No | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `string` | No | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `string` | No | The currency of the source account. |
| `sourceAccountIban` | `string` | No | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `table` | Yes |  |
| `sourceAccountName` | `string` | No | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `string` | No | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `string` | No | The sort code of the account the payout is being made from. |
| `status` | `string` | No | Gets or Sets the status of payout request |
| `tags` | `table` | No | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `string` | No | Gets or Sets destination reference ID |
| `topupPayrunID` | `string` | No | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `number` | No | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `number` | No | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `number` | No | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `string` | No | Gets or Sets payout type |
| `userID` | `string` | No | Gets or Sets User ID of who created the payout request |
| `yourReference` | `string` | No | Gets or Sets your reference ID |

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
| `authorisations` | `table` | No | A list of the users who have successfully authorised the latest version of the payrun and when. |
| `authorisersCompletedCount` | `number` | No | The number of distinct authorisers that have authorised the payrun. |
| `authorisersRequiredCount` | `number` | No | The number of authorisers required for this payrun. |
| `batchPayoutID` | `string` | No |  |
| `canAuthorise` | `boolean` | No | True if the payrun can be authorised by the user who loaded it. |
| `canDelete` | `boolean` | No |  |
| `canEdit` | `boolean` | No |  |
| `events` | `table` | No |  |
| `hasCurrentUserAuthorised` | `boolean` | No | True if the payrun was loaded for a user and that user has already authorised the latest version of the payrun. |
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
| `accountID` | `string` | No | The ID of the account the rule will apply to. |
| `approveUrl` | `string` | No | If set this property holds the URL an approver needs to visit in order to complete a strong authentication check in order to approve the rule. |
| `approverID` | `string` | No |  |
| `authenticationMethods` | `table` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `table` | No | A list of the users who have successfully authorised the latest version of the rule and when. |
| `authorisersCompletedCount` | `number` | No | The number of distinct authorisers that have authorised the rule. |
| `authorisersRequiredCount` | `number` | No | The number of authorisers required for this rule. |
| `canAuthorise` | `boolean` | No | True if the rule can be authorised by the user who loaded it. |
| `createdBy` | `table` | Yes |  |
| `description` | `string` | No | Arbitrary description for the rule. |
| `endAt` | `string` | No | Optional end time for rule executions. |
| `hasCurrentUserAuthorised` | `boolean` | No | True if the current user has authorised. |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isDisabled` | `boolean` | No | If set to true the rule will be disabled from executing. |
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
| `sweepAction` | `table` | No |  |
| `timeZoneId` | `string` | No | If the rule should be executed on a recurring schedule this is the timezone that the CRON expression should be evaluated in. |
| `triggerCronExpression` | `string` | No | If the rule should be executed on a recurring schedule this is the expression that sets the schedule. |
| `triggerOnPayIn` | `boolean` | No | Set to true if the rule execution should be triggered when the account receives a pay in (credit). |
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
| `accountID` | `string` | No | The ID of the account the transaction belongs to. |
| `accountName` | `string` | No | The name of the account the transaction belongs to. |
| `accountSequenceNumber` | `number` | No | The sequence number of transaction on a per account basis. |
| `addressDetails` | `table` | No |  |
| `amount` | `number` | No | Amount of the transaction. |
| `amountMinorUnits` | `number` | No | Amount of the transaction expressed in the currency’s minor units (e.g. |
| `balance` | `number` | No | Balance left on the account after the transaction. |
| `balanceMinorUnits` | `number` | No | Balance on the account expressed in the currency’s minor units (e.g. |
| `bookingDateTime` | `string` | No |  |
| `chargeDetails` | `table` | No |  |
| `content` | `table` | No |  |
| `counterparty` | `table` | No |  |
| `counterpartySummary` | `string` | No | For pay in (credit) transactions this will contain a descriptive string with the most important fields about the counterparty. |
| `currency` | `string` | No | Currency of transaction. |
| `currencyExchange` | `table` | No | Provides details on the currency exchange. |
| `date` | `string` | No |  |
| `description` | `string` | No | Description of the transaction. |
| `enrichment` | `table` | No |  |
| `fxAmount` | `number` | No | For an FX payout this is the amound in the FX currency. |
| `fxCurrency` | `string` | No | For an FX payout this is the currency that was received or that was instructed. |
| `fxRate` | `number` | No | For an FX payout this is the exchange rate between the transaction currency and the FX currency. |
| `grossAmount` | `table` | Yes |  |
| `id` | `string` | No | Unique ID for the transaction. |
| `inserted` | `string` | No | Date when the transaction was inserted into the ledger. |
| `isoBankTransactionCode` | `table` | No |  |
| `merchant` | `table` | No |  |
| `merchantID` | `string` | No | The ID of the merchant that owns the account. |
| `pageNumber` | `number` | No | Current page number. |
| `pageSize` | `number` | No | Page size |
| `payeeDetails` | `table` | Yes | The Payee object contains details of the beneficiary, person or business. |
| `payerDetails` | `table` | Yes |  |
| `paymentRequestCustomFields` | `table` | No | The custom fields that were attached to the payment request that resulted in this transaction. |
| `paymentRequestID` | `string` | No | For Pay by Bank and Direct Debit transactions this will contain the ID of the payment request. |
| `payoutID` | `string` | No | ID of the payout that resulted in the transaction. |
| `proprietaryBankTransactionCode` | `table` | No |  |
| `rawReference` | `string` | No | The raw payment reference details as received from the payment processor. |
| `reference` | `string` | No |  |
| `ruleID` | `string` | No | ID of the rule that resulted in the transaction. |
| `statementReferences` | `table` | No |  |
| `status` | `string` | No |  |
| `supplementaryData` | `any` | No |  |
| `tags` | `table` | No | An optional list of descriptive tags attached to the transaction. |
| `theirReference` | `string` | No | For a pay out the reference that the payer attached for the receiving party. |
| `totalPages` | `number` | No | Total pages |
| `totalSize` | `number` | No | Total count |
| `transactionAmount` | `table` | Yes |  |
| `transactionDate` | `string` | No | Date when the transaction occurred. |
| `transactionInformation` | `table` | No |  |
| `transactionMutability` | `string` | No |  |
| `type` | `string` | No | Type of the transaction. |
| `valueDateTime` | `string` | No |  |
| `virtualIBAN` | `string` | No | If set it indicates the payin was to a virtual IBAN. |
| `yourReference` | `string` | No | For a pay in the reference the sending party attached. |

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
| `clientSessionTimeouts` | `table` | No | The number of seconds a session for this user should last before expiring. |
| `emailAddress` | `string` | Yes |  |
| `firstName` | `string` | Yes |  |
| `id` | `string` | No |  |
| `lastName` | `string` | Yes |  |
| `passkeyAdded` | `boolean` | No |  |
| `permissions` | `table` | No |  |
| `profile` | `string` | No |  |
| `rolesWithScope` | `table` | No |  |
| `twoFactorEnabled` | `boolean` | No |  |
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
| `initialRoleID` | `string` | No | The role ID to automatically assign to the merchant’s very first user. |
| `inviteeEmailAddress` | `string` | No | Email address of the user being invited. |
| `inviteeFirstName` | `string` | No | First Name of the user being invited. |
| `inviteeLastName` | `string` | No | Last Name of the user being invited. |
| `inviterEmailAddress` | `string` | No |  |
| `inviterFirstName` | `string` | No |  |
| `inviterLastName` | `string` | No |  |
| `isAuthorised` | `boolean` | No | Will be set to true once the invite has met the authorisation requirements. |
| `isInviteeRegistered` | `boolean` | No | If true, indicates the invitee's email address corresponds to an existing MoneyMoov user. |
| `lastInvited` | `string` | No |  |
| `merchantID` | `string` | No | ID of the merchant the user is being invited to. |
| `merchantName` | `string` | No |  |
| `message` | `string` | No |  |
| `registrationUrl` | `string` | No |  |
| `sendInviteEmail` | `boolean` | No | If set to true an email will be sent to the invitee with instructions on how to accept the invite. |
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
  id = --[[ string ]],
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
| `accountName` | `string` | No | Name for the account |
| `accountSupplierName` | `string` | No | The payment account supplier name. |
| `availableBalance` | `number` | No | The current available balance of the account. |
| `availableBalanceMinorUnits` | `number` | No | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `number` | No | Balance of the account. |
| `balanceMinorUnits` | `number` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `string` | No | The bank name for external accounts |
| `consentID` | `string` | No | The ID of the consent used to connect the external account. |
| `createdBy` | `table` | Yes |  |
| `createdByDisplayName` | `string` | No | Either the name of the user, merchant token or api key that created the account |
| `currency` | `string` | No | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `string` | No | Indicates the default payment rail for this account. |
| `displayName` | `string` | No | Gets a unique display name for the payment account. |
| `expiryDate` | `string` | No | The date that the external account will expire |
| `externalAccountIcon` | `string` | No | The Icon for external accounts |
| `id` | `string` | No | Unique id for the account. |
| `identifier` | `table` | Yes |  |
| `inserted` | `string` | No | Timestamp when the account was created. |
| `isArchived` | `boolean` | No | Indicates whether the account is archived. |
| `isConnectedAccount` | `boolean` | No | Indicates if the payment account is an externally connected account. |
| `isDefault` | `boolean` | No | Is the default account |
| `isTrustAccount` | `boolean` | No | Indicates if the payment account is a trust account. |
| `isVirtual` | `boolean` | No | True if the account is a virtual account. |
| `lastTransaction` | `table` | No |  |
| `lastUpdated` | `string` | No | Timestamp when the account was last updated. |
| `merchantID` | `string` | No | The ID of the merchant that owns the account. |
| `merchantName` | `string` | No | The name of the merchant that owns the account. |
| `name` | `string` | Yes | The name of the virtual account. |
| `physicalAccountID` | `string` | No | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `table` | No | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `number` | No | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `number` | No | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `string` | No | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `string` | No | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `string` | No | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `string` | No |  |
| `xeroBankFeedSyncStatus` | `string` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `number` | No | Indicates the number of unsynchronised transactions with Xero |

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
| `destinationUrl` | `string` | No | The destination URL for the webhook. |
| `emailAddress` | `string` | No | The recipient email address(es) for notifications. |
| `failedNotificationEmailAddress` | `string` | No | The email address to which notifications about failed webhook deliveries will be sent. |
| `id` | `string` | No |  |
| `isActive` | `boolean` | No |  |
| `merchantID` | `string` | No | The ID of the merchant that the webhook is for. |
| `notificationMethod` | `string` | No | The type of notification that will be sent. |
| `resourceTypes` | `table` | No | The resource types that the webhook will be generated for. |
| `retry` | `boolean` | No |  |
| `secret` | `string` | No | The secret key required to authenticate webhook notifications. |
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

