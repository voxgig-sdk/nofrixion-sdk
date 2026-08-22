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
| `accountBalances` | `Array` | No | The various balances for the account. |
| `accountID` | `String` | No | ID of the account. |
| `accountIdentifications` | `Array` | No | The canoncial identifiers for the account. |
| `accountName` | `String` | No | Name for the account |
| `accountNames` | `Array` | No | Optional account names set by the account holder. |
| `accountSupplierName` | `String` | No | The payment account supplier name. |
| `accountType` | `String` | No | The type of account e.g. |
| `availableBalance` | `Float` | No | The current available balance of the account. |
| `availableBalanceMinorUnits` | `Integer` | No | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `Float` | No | Balance of the account. |
| `balanceMinorUnits` | `Integer` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `String` | No | The bank name for external accounts |
| `consentID` | `String` | No | The ID of the consent used to connect the external account. |
| `consolidatedAccountInformation` | `Hash` | No | Summary information regarding account balances of the overall account provided by the bank. |
| `createdBy` | `Hash` | Yes |  |
| `createdByDisplayName` | `String` | No | Either the name of the user, merchant token or api key that created the account |
| `currency` | `String` | No | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `String` | No | Indicates the default payment rail for this account. |
| `description` | `String` | No | Product name as defined by the financial institution for this account. |
| `details` | `String` | No | Supplementary specifications that might be provided by the Bank. |
| `displayName` | `String` | No | Gets a unique display name for the payment account. |
| `expiryDate` | `String` | No | The date that the external account will expire |
| `externalAccountIcon` | `String` | No | The Icon for external accounts |
| `format` | `String` | No | File format to save the statement as. |
| `fromDate` | `String` | No | Minimum transaction date for the statement. |
| `id` | `String` | No | Unique id for the account. |
| `identifier` | `Hash` | Yes |  |
| `inserted` | `String` | No | Timestamp when the account was created. |
| `isArchived` | `Boolean` | No | Indicates whether the account is archived. |
| `isConnectedAccount` | `Boolean` | No | Indicates if the payment account is an externally connected account. |
| `isDefault` | `Boolean` | No | Is the default account |
| `isTrustAccount` | `Boolean` | No | Indicates if the payment account is a trust account. |
| `isVirtual` | `Boolean` | No | True if the account is a virtual account. |
| `lastTransaction` | `Hash` | No |  |
| `lastUpdated` | `String` | No | Timestamp when the account was last updated. |
| `merchantID` | `String` | No | The ID of the merchant that owns the account. |
| `merchantName` | `String` | No | The name of the merchant that owns the account. |
| `nickname` | `String` | No | Nickname of the account that was provided by the account owner. |
| `physicalAccountID` | `String` | No | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `roleIDs` | `Array` | No | Optional list of role IDs that will get access to the payment account when created. |
| `rules` | `Array` | No | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `Float` | No | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `Integer` | No | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `String` | No | Gets a summary of the payments account's most important properties. |
| `supplierPhysicalAccountID` | `String` | No | For internal use only. |
| `supplierSepaInstantStatus` | `String` | No | Indicates the status of the SEPA Instant payment rail for this account. |
| `toDate` | `String` | No | Maximum transaction date for the statement. |
| `type` | `String` | No | Specifies the type of account e.g. |
| `usageType` | `String` | No |  |
| `xeroBankFeedConnectionStatus` | `String` | No | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `String` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `String` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `String` | No |  |
| `xeroBankFeedSyncStatus` | `String` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `Integer` | No | Indicates the number of unsynchronised transactions with Xero |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Account.create({
  "account_id" => "example_account_id", # String
  "currency" => "example_currency", # String
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
| `approveUrl` | `String` | No | This field is used when returning a batch payout record to a client. |
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
| `authenticationMethods` | `Array` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `Array` | No | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `Integer` | No | The number of distinct authorisers that have authorised the beneficiary. |
| `authorisersRequiredCount` | `Integer` | No | The number of authorisers required for this beneficiary. |
| `beneficiaries` | `Array` | No |  |
| `beneficiaryEvents` | `Array` | No |  |
| `canAuthorise` | `Boolean` | No | True if the beneficiary can be authorised by the user who loaded it. |
| `canUpdate` | `Boolean` | No | True if the beneficiary can be updated by the user who loaded it. |
| `createdBy` | `Hash` | Yes |  |
| `createdByEmailAddress` | `String` | No |  |
| `currency` | `String` | Yes | Gets or Sets the currency. |
| `destination` | `Hash` | No |  |
| `failedBeneficiaries` | `Hash` | No |  |
| `hasCurrentUserAuthorised` | `Boolean` | No | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `isEnabled` | `Boolean` | No |  |
| `lastAuthorised` | `String` | No |  |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No | Gets or Sets the merchant id. |
| `name` | `String` | Yes | The descriptive name for the beneficiary. |
| `nonce` | `String` | No |  |
| `sourceAccountIDs` | `Array` | No | ID of the accounts which are authorised to act as a source for the beneficiary. |
| `sourceAccounts` | `Array` | No |  |
| `theirReference` | `String` | No | The reference that will be used by default as TheirReference when creating payouts to this beneficiary if no TheirReference is specified for the payout. |

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
  "id" => "example_id", # String
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
| `groupMembers` | `Array` | No | The existing group members. |
| `groupName` | `String` | Yes | The descriptive name for the beneficiary group. |
| `id` | `String` | No |  |
| `inserted` | `String` | No | Timestamp indicating when the group was created. |
| `lastUpdated` | `String` | No | Timestamp indicating when the group was last updated. |
| `merchantID` | `String` | Yes | Gets or Sets the merchant id. |

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
| `isPayerAuthenticationRequired` | `Boolean` | No | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `Boolean` | No | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `String` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `String` | No | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `String` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `Integer` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `Integer` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `String` | No | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `String` | No |  |
| `requestID` | `String` | No |  |
| `responseCode` | `String` | No |  |
| `responseType` | `String` | No |  |
| `status` | `String` | No |  |
| `threeDSRedirectUrl` | `String` | No | Checkout.com require a redirect for 3DS authentication. |
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
| `cardType` | `String` | No | The type of the tokenised card, e.g. |
| `customerEmailAddress` | `String` | No | When creating a tokenised card the payer's email address must be supplied. |
| `expiryMonth` | `String` | No |  |
| `expiryYear` | `String` | No |  |
| `id` | `String` | No | The unique ID of the card token that has been stored for the customer. |
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
result = client.CardCustomerToken.remove({ "id" => "id" })
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
| `isPayerAuthenticationRequired` | `Boolean` | No | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `Boolean` | No | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `String` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `String` | No | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `String` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `Integer` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `Integer` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `String` | No | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `String` | No |  |
| `requestID` | `String` | No |  |
| `responseCode` | `String` | No |  |
| `responseType` | `String` | No |  |
| `status` | `String` | No |  |
| `threeDSRedirectUrl` | `String` | No | Checkout.com require a redirect for 3DS authentication. |
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
| `authorisationUrl` | `String` | No | The URL the authorising user needs to be redirected to in order to get the open banking consent token. |
| `callbackUrl` | `String` | No | Optional callback URL that the end user performing the open banking authorisation will be redirected to on completion. |
| `consentID` | `String` | No | The ID of the open banking consent. |
| `emailAddress` | `String` | No | The email address that identifies the end user that will be authorising the open banking consent request. |
| `expiryDate` | `String` | No |  |
| `failureCallbackUrl` | `String` | No | Optional callback URL for open banking consent authorisation failure. |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `institutionID` | `String` | No | The institution ID the open banking consent is being requested for. |
| `isConnectedAccounts` | `Boolean` | No | Optional setting. |
| `isEnabled` | `Boolean` | No |  |
| `merchantID` | `String` | No | The ID of the merchant the consent token is being created to be used with. |
| `provider` | `String` | No | Lists the supported card and PIS processors. |
| `successWebHookUrl` | `String` | No | A web hook URL to send an HTTP request to when the open banking consent is successfuly authorised. |

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
| `failedSubmissions` | `Hash` | No | Dictionary of failed submissions, keyed by the index (1-based) in the original request. |
| `successfulSubmissions` | `Array` | No | List of successfully submitted direct debit payments. |

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
| `exchangeRate` | `Float` | No | The price at which the transaction will buy the source currency using the destination currency. |
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
| `accountNumber` | `String` | No | Account number of the customer's bank account in case of GBP account. |
| `addressLine1` | `String` | Yes | First line of the customer's address. |
| `addressLine2` | `String` | No | Second line of the customer's address. |
| `approvedAt` | `String` | No | Date at which the supplier approved this mandate. |
| `city` | `String` | Yes | Customer's city. |
| `countryCode` | `String` | Yes | 2-character country code of the customer's bank account. |
| `currency` | `String` | No | Currency of this mandate. |
| `customerAccountNumber` | `String` | No | Customer's account number in case of GBP account. |
| `customerCity` | `String` | No | Customer's city of residence. |
| `customerCountryCode` | `String` | No | Customer's country of residence code. |
| `customerCountryName` | `String` | No | Customer's country of residence. |
| `customerEmailAddress` | `String` | No | Customer's email address. |
| `customerFirstName` | `String` | No | Customer's first name. |
| `customerIban` | `String` | No | Customer's IBAN in case of EUR account. |
| `customerLastName` | `String` | No | Customer's last name. |
| `customerSortCode` | `String` | No | Customer's sort code in case of GBP account. |
| `emailAddress` | `String` | Yes | Customer's email address. |
| `firstName` | `String` | Yes | Customer's first name. |
| `iban` | `String` | No | IBAN of the customer's bank account in case of EUR account. |
| `id` | `String` | No | Internal ID of the mandate. |
| `inserted` | `String` | No | The timestamp this mandate was created at. |
| `isRecurring` | `Boolean` | No | Whether this mandate is single-use or recurring. |
| `lastName` | `String` | Yes | Customer's last name. |
| `lastUpdated` | `String` | No | The timestamp this mandate was last updated at. |
| `merchantID` | `String` | No | Internal ID of this mandate's merchant. |
| `postalCode` | `String` | Yes | Customer's postal code. |
| `reference` | `String` | No | Reference assigned to this mandate. |
| `sortCode` | `String` | No | Sort code of the customer's bank account in case of GBP account. |
| `status` | `String` | No | General status of this mandate. |
| `supplierBankAccountID` | `String` | No | ID that the supplier assigned to this mandate's bank account. |
| `supplierCustomerID` | `String` | No | ID that the supplier assigned to this mandate's customer. |
| `supplierMandateID` | `String` | No | ID that the supplier assigned to this mandate. |
| `supplierName` | `String` | No | Name of the supplier used to create this mandate. |
| `supplierStatus` | `String` | No | Last status that the supplier reported for this mandate. |

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
| `accountCurrencies` | `Array` | No | The list of currencies that the merchant has accounts for. |
| `canHaveTrustAccounts` | `Boolean` | No | Trust accounts are a special type of account that allow the account name to be trusted for use in statements and verification of payee checks. |
| `cardPaymentProcessor` | `String` | No | Name of the card payment processor. |
| `companyID` | `String` | No | The Company ID recorded in the Compliance system. |
| `displayQrOnHostedPay` | `Boolean` | No | Indicates if a QR Code containing the payment link should be displayed on the hosted payment page. |
| `hostedPayVersion` | `Integer` | No | The version of the hosted payment page to use with the merchant. |
| `id` | `String` | No | Unique ID for the merchant. |
| `inserted` | `String` | No | Timestamp the merchant was added to MoneyMoov. |
| `isBlocked` | `Boolean` | No | The merchant is blocked from making payments (payouts). |
| `isExited` | `Boolean` | No | The merchant has formally terminated their relationship and is no longer a customer. |
| `isSuspended` | `Boolean` | No | The merchant has temporarily suspended their own account. |
| `jurisdiction` | `String` | No | The jurisdiction the merchant entity is incorporated or established in. |
| `logoUrlPng` | `String` | No | The CDN URL of the merchant's logo in PNG format. |
| `logoUrlSvg` | `String` | No | The CDN URL of the merchant's logo in SVG format. |
| `merchantCategoryCode` | `String` | No | The industry code that represents the merchant's primary trading activity. |
| `name` | `String` | No | The registered business name of the merchant. |
| `notes` | `String` | No | The notes field is an optional free text field that can be used to store any additional information about the merchant. |
| `parentMerchant` | `Hash` | No |  |
| `paymentAccountLimit` | `Integer` | No | The maximum number of payment accounts that can be created for the Merchant. |
| `paymentAccounts` | `Array` | No |  |
| `reason` | `String` | No | The reason for the suspension. |
| `shortName` | `String` | No | A URL friendly shortish name for the merchant. |
| `supportedPaymentMethodsList` | `Array` | No | The payment methods that are configured and supported for this merchant. |
| `suspensionReason` | `String` | No | The reason for the suspension, provided by the merchant. |
| `tags` | `Array` | No | An optional list of descriptive tags that can be used on merchant entities such as payment requests. |
| `timeZoneId` | `String` | No | The IANA (Internet Assigned Numbers Authority) time zone identifier of the merchant. |
| `tradingName` | `String` | No | An optional trading name. |
| `webHookLimit` | `Integer` | No | The maximum number of web hooks that can be created for the Merchant. |
| `yourRoleName` | `String` | No | The name of the role for the identity that loaded the merchant record. |

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
result = client.Merchant.remove({ "id" => "merchant_id", "user_id" => "user_id" })
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
| `approvedAt` | `String` | No | Date at which the supplier approved this mandate. |
| `currency` | `String` | No | Currency of this mandate. |
| `customerAccountNumber` | `String` | No | Customer's account number in case of GBP account. |
| `customerCity` | `String` | No | Customer's city of residence. |
| `customerCountryCode` | `String` | No | Customer's country of residence code. |
| `customerCountryName` | `String` | No | Customer's country of residence. |
| `customerEmailAddress` | `String` | No | Customer's email address. |
| `customerFirstName` | `String` | No | Customer's first name. |
| `customerIban` | `String` | No | Customer's IBAN in case of EUR account. |
| `customerLastName` | `String` | No | Customer's last name. |
| `customerSortCode` | `String` | No | Customer's sort code in case of GBP account. |
| `id` | `String` | No | Internal ID of the mandate. |
| `inserted` | `String` | No | The timestamp this mandate was created at. |
| `isRecurring` | `Boolean` | No | Whether this mandate is single-use or recurring. |
| `lastUpdated` | `String` | No | The timestamp this mandate was last updated at. |
| `merchantID` | `String` | No | Internal ID of this mandate's merchant. |
| `reference` | `String` | No | Reference assigned to this mandate. |
| `status` | `String` | No | General status of this mandate. |
| `supplierBankAccountID` | `String` | No | ID that the supplier assigned to this mandate's bank account. |
| `supplierCustomerID` | `String` | No | ID that the supplier assigned to this mandate's customer. |
| `supplierMandateID` | `String` | No | ID that the supplier assigned to this mandate. |
| `supplierName` | `String` | No | Name of the supplier used to create this mandate. |
| `supplierStatus` | `String` | No | Last status that the supplier reported for this mandate. |

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
| `bankCountryCodes` | `Array` | No | The list of country codes representing the banks the country supports. |
| `bankID` | `String` | No | ID of the bank to be configured for the merchant. |
| `bankName` | `String` | No | Name of the Bank/Institution. |
| `businessInstitutionID` | `String` | No | ID that the processor uses to identify the bank (business accounts). |
| `currency` | `String` | No | Currency supported by the bank. |
| `logo` | `String` | No | URL of the bank's logo. |
| `message` | `String` | No | Message relating to specific bank. |
| `messageImageUrl` | `String` | No | Optional image URL to be displayed with the message. |
| `order` | `Integer` | No | Order in which this setting will appear in the UI. |
| `personalInstitutionID` | `String` | No | ID that the processor uses to identify the bank (personal accounts). |
| `processor` | `String` | No | Name of the bank payment processor. |
| `warningHeading` | `String` | No | The heading for a warning message related to the bank institution to be displayed to the user. |
| `warningMessage` | `String` | No | The warning message related to the bank institution to be displayed to the user. |

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
| `customFields` | `Array` | No | A list of custom fields that can be included in the payment request template. |
| `defaultFields` | `Array` | No | A list of default fields that are included in the payment request template. |
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
| `authenticationMethods` | `Array` | No | A list of authentication types allowed to authorise the merchant token. |
| `authorisations` | `Array` | No | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `Integer` | No | The number of distinct authorisers that have authorised the merchant token. |
| `authorisersRequiredCount` | `Integer` | No | The number of authorisers required for this merchant token. |
| `canAuthorise` | `Boolean` | No | True if the merchant token can be authorised by the user who loaded it. |
| `description` | `String` | No | Token description |
| `expiresAt` | `String` | No | Optional. |
| `hasCurrentUserAuthorised` | `Boolean` | No | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `hmacAlgorithm` | `String` | No | Optional shared secret algorithm to use for HMAC authentication. |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `ipAddressWhitelist` | `String` | No | Optional. |
| `isArchived` | `Boolean` | No | Indicates whether the merchant token is archived. |
| `isEnabled` | `Boolean` | No | If set to false the merchant token will not be accepted to authorise a request. |
| `lastAuthorised` | `String` | No |  |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No | The merchant id to add to the token |
| `nonce` | `String` | Yes |  |
| `permissionTypes` | `Array` | No | The permissions that the merchant token supports. |
| `requestSignatureVersion` | `Integer` | No | Represent the version of the overall merchant token. |
| `sharedSecretAlgorithm` | `String` | No | Optional shared secret algorithm to use for HMAC authentication. |
| `sharedSecretBase64` | `String` | No | The base 64 encoded shared secret that is used for request authentication with an HMAC. |
| `token` | `String` | No | The JWT merchant token. |

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
result = client.OpenBanking.remove({ "account_id" => "account_id" })
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
| `accountName` | `String` | Yes | The name of the account to verify |
| `accountNumber` | `String` | No | The account number of the account to verify (for CoP checks) |
| `iban` | `String` | Yes | The IBAN of the account to verify (for VoP checks) |
| `payeeVerifiedAccountName` | `String` | No | The verified account name of the payee, if available (in case of a close match) |
| `result` | `String` | No | The result of the payee verification |
| `secondaryIdentification` | `String` | No | Optional secondary identifier for the account to verify. |
| `sortCode` | `String` | No | The sort code of the account to verify (for CoP checks) |

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
| `amount` | `Float` | No | The amount of money to request. |
| `amountPending` | `Float` | No | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `Float` | No | Total amount received for this payment request. |
| `amountRefunded` | `Float` | No | Total amount refunded for this payment request. |
| `autoSendReceipt` | `Boolean` | No | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `String` | No | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `String` | No | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `Boolean` | No | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `Boolean` | No | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `String` | No | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `Boolean` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardNoPayerAuthentication` | `Boolean` | No | If set to true for card payments no attempt will be made to use payer authentication (3-D Secure and equivalent). |
| `cardProcessorMerchantID` | `String` | No | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `String` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `String` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `cardTransmitRawDetails` | `Boolean` | No | If set to true for card payments the sensitive card number and card verification number will be transmitted directly rather than being tokenised. |
| `createdByUser` | `Hash` | Yes |  |
| `currency` | `String` | No | The currency of the request. |
| `customFields` | `Array` | No | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `String` | No | Optional email address for the customer. |
| `customerID` | `String` | No | An optional customer identifier for the payment request. |
| `customerName` | `String` | No |  |
| `description` | `String` | No | An optional description for the payment request. |
| `destinationAccount` | `Hash` | No |  |
| `directDebitPayment` | `Hash` | No | Contains information about a Direct Debit payment attempt for a payment request. |
| `dueDate` | `String` | No | The due date for the payment request. |
| `events` | `Array` | No |  |
| `failureCallbackUrl` | `String` | No | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `Array` | No | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `String` | No |  |
| `hostedPayCheckoutUrl` | `String` | No | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `String` | No |  |
| `ignoreAddressVerification` | `Boolean` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `String` | No | The timestamp the payment request was created at. |
| `insertedSortable` | `String` | No | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `isArchived` | `Boolean` | No | Indicates whether the payment request is archived. |
| `jwk` | `String` | No | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `String` | No | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `String` | No | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `String` | No | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `String` | No | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `String` | No | The ID of the merchant to create the payment request for. |
| `merchantTokenDescription` | `String` | No | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `String` | No |  |
| `notificationRoleIDs` | `Array` | No | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `String` | No | An optional order ID for the payment request. |
| `partialPaymentMethod` | `String` | No | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `String` | No | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `Array` | No | The payment attempts made against this payment request. |
| `paymentMethods` | `Array` | No | The payment methods that the payment request supports. |
| `paymentProcessor` | `String` | No | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `payrunID` | `String` | No | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `String` | No | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `String` | No | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `Hash` | No |  |
| `sandboxSettleDelayInSeconds` | `Integer` | No | Sandbox only. |
| `shippingAddress` | `Hash` | No |  |
| `shippingAddressCity` | `String` | No | Optionally the city of the customer's shipping address. |
| `shippingAddressCountryCode` | `String` | No | Optionally the country code of the customer's shipping address. |
| `shippingAddressCounty` | `String` | No | Optionally the state or county of the customer's shipping address. |
| `shippingAddressLine1` | `String` | No | Optionally the first line of the customer's shipping address. |
| `shippingAddressLine2` | `String` | No | Optionally the second line of the customer's shipping address. |
| `shippingAddressPostCode` | `String` | No | Optionally the post code of the customer's shipping address. |
| `shippingEmail` | `String` | No | Optionally the shipping email address for the customer. |
| `shippingFirstName` | `String` | No | Optionally the first name of the customer's shipping address. |
| `shippingLastName` | `String` | No | Optionally the last name of the customer's shipping address. |
| `shippingPhone` | `String` | No | Optionally the shipping phone number for the customer. |
| `status` | `String` | No | The current status of the payment request. |
| `successWebHookUrl` | `String` | No | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tagIds` | `Array` | No | An optional list of tag ids to add to the payment request |
| `tags` | `Array` | No | An optional list of descriptive tags attached to the payment request. |
| `title` | `String` | No | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `Array` | No |  |
| `transactions` | `Array` | No |  |
| `useHostedPaymentPage` | `Boolean` | No | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

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
| `accountName` | `String` | No | Name for the account |
| `accountSupplierName` | `String` | No | The payment account supplier name. |
| `availableBalance` | `Float` | No | The current available balance of the account. |
| `availableBalanceMinorUnits` | `Integer` | No | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `Float` | No | Balance of the account. |
| `balanceMinorUnits` | `Integer` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `String` | No | The bank name for external accounts |
| `consentID` | `String` | No | The ID of the consent used to connect the external account. |
| `createdBy` | `Hash` | Yes |  |
| `createdByDisplayName` | `String` | No | Either the name of the user, merchant token or api key that created the account |
| `currency` | `String` | No | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `String` | No | Indicates the default payment rail for this account. |
| `displayName` | `String` | No | Gets a unique display name for the payment account. |
| `expiryDate` | `String` | No | The date that the external account will expire |
| `externalAccountIcon` | `String` | No | The Icon for external accounts |
| `id` | `String` | No | Unique id for the account. |
| `identifier` | `Hash` | Yes |  |
| `inserted` | `String` | No | Timestamp when the account was created. |
| `isArchived` | `Boolean` | No | Indicates whether the account is archived. |
| `isConnectedAccount` | `Boolean` | No | Indicates if the payment account is an externally connected account. |
| `isDefault` | `Boolean` | No | Is the default account |
| `isTrustAccount` | `Boolean` | No | Indicates if the payment account is a trust account. |
| `isVirtual` | `Boolean` | No | True if the account is a virtual account. |
| `lastTransaction` | `Hash` | No |  |
| `lastUpdated` | `String` | No | Timestamp when the account was last updated. |
| `merchantID` | `String` | No | The ID of the merchant that owns the account. |
| `merchantName` | `String` | No | The name of the merchant that owns the account. |
| `physicalAccountID` | `String` | No | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `Array` | No | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `Float` | No | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `Integer` | No | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `String` | No | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `String` | No | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `String` | No | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `String` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `String` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `String` | No |  |
| `xeroBankFeedSyncStatus` | `String` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `Integer` | No | Indicates the number of unsynchronised transactions with Xero |

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
| `accountName` | `String` | No | Name for the account |
| `availableBalance` | `Float` | No | The current available balance of the account. |
| `balance` | `Float` | No | Balance of the account. |
| `balanceMinorUnits` | `Integer` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `currency` | `String` | No | Currency of the account in ISO 4217 format |
| `id` | `String` | No | Unique id for the account. |
| `identifier` | `Hash` | Yes |  |
| `isArchived` | `Boolean` | No | Is the account archived |
| `isConnectedAccount` | `Boolean` | No | Indicates if the payment account is an externally connected account. |
| `merchantID` | `String` | No | The ID of the merchant that owns the account. |
| `submittedPayoutsBalance` | `Float` | No | Total of the payouts that have been submitted for processing. |

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
| `paymentInitiationID` | `String` | No | The unique identifier of the payment initiation request. |
| `paymentRequestCallbackUrl` | `String` | No | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `String` | No |  |
| `redirectUrl` | `String` | No | A redirect URL for the user to authorise the payment initiation request at the ASPSP |
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
| `amount` | `Float` | No | The amount of money to request. |
| `amountPending` | `Float` | No | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `Float` | No | Total amount received for this payment request. |
| `amountRefunded` | `Float` | No | Total amount refunded for this payment request. |
| `autoSendReceipt` | `Boolean` | No | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `String` | No | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `String` | No | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `Boolean` | No | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `Boolean` | No | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `String` | No | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `Boolean` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardProcessorMerchantID` | `String` | No | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `String` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `String` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `createdByUser` | `Hash` | Yes |  |
| `currency` | `String` | No | The currency of the request. |
| `customFields` | `Array` | No | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `String` | No | Optional email address for the customer. |
| `customerID` | `String` | No | An optional customer identifier for the payment request. |
| `customerName` | `String` | No |  |
| `description` | `String` | No | An optional description for the payment request. |
| `destinationAccount` | `Hash` | No |  |
| `directDebitPayment` | `Hash` | No | Contains information about a Direct Debit payment attempt for a payment request. |
| `doSimulateSettlementFailure` | `Boolean` | No |  |
| `dueDate` | `String` | No | The due date for the payment request. |
| `errorDescription` | `String` | No |  |
| `events` | `Array` | No |  |
| `failedPaymentRequests` | `Hash` | No |  |
| `failureCallbackUrl` | `String` | No | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `Array` | No | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `String` | No |  |
| `hostedPayCheckoutUrl` | `String` | No | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `String` | No |  |
| `ignoreAddressVerification` | `Boolean` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `String` | No | The timestamp the payment request was created at. |
| `insertedSortable` | `String` | No | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `institution` | `String` | No |  |
| `isArchived` | `Boolean` | No | Indicates whether the payment request is archived. |
| `jwk` | `String` | No | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `String` | No | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `String` | No | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `String` | No | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `String` | No | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `String` | No |  |
| `merchantTokenDescription` | `String` | No | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `String` | No |  |
| `notificationRoleIDs` | `Array` | No | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `String` | No | An optional order ID for the payment request. |
| `partialPaymentMethod` | `String` | No | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `String` | No | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `Array` | No | The payment attempts made against this payment request. |
| `paymentInitiationID` | `String` | No |  |
| `paymentMethods` | `Array` | No | The payment methods that the payment request supports. |
| `paymentProcessor` | `String` | No | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `paymentRequests` | `Array` | No |  |
| `payrunID` | `String` | No | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `String` | No | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `String` | No | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `Hash` | No |  |
| `sandboxSettleDelayInSeconds` | `Integer` | No | Sandbox only. |
| `shippingAddress` | `Hash` | No |  |
| `status` | `String` | No | The current status of the payment request. |
| `successWebHookUrl` | `String` | No | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tags` | `Array` | No | An optional list of descriptive tags attached to the payment request. |
| `title` | `String` | No | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `Array` | No |  |
| `transactions` | `Array` | No |  |
| `useHostedPaymentPage` | `Boolean` | No | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

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
result = client.PaymentRequest.load({ "id" => "payment_request_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.PaymentRequest.remove({ "id" => "payment_request_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.PaymentRequest.update({
  "id" => "payment_request_id",
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
| `applePayTransactionID` | `String` | No | Transaction ID received in Apple pay token. |
| `cardAuthorizationResponseID` | `String` | No | For a successful card authorization this field will hold the response ID. |
| `cardExpiryMonth` | `Integer` | No | For card payment events this field holds the payer's card expiry month. |
| `cardExpiryYear` | `Integer` | No | For card payment events this field holds the payer's card expiry year. |
| `cardIssuer` | `String` | No | For card payment events this field holds the payer's card issuer. |
| `cardIssuerCountry` | `String` | No | For card payment events this field holds the payer's card issuer country of origin. |
| `cardLastFourDigits` | `String` | No | For card payment events this field holds the payer's card last four digits. |
| `cardRequestID` | `String` | No |  |
| `cardScheme` | `String` | No | For card payment events this field holds the scheme of the payer's card, e.g. |
| `cardTokenCustomerID` | `String` | No | If the option to create a reusable token for card payments was set this field contains the token the merchant can store to use for repeat payments. |
| `cardTransactionID` | `String` | No |  |
| `currency` | `String` | No |  |
| `directDebitPaymentID` | `String` | No | Payment ID issued by the Direct Debit supplier. |
| `directDebitPaymentReference` | `String` | No | Reference string issued by the Direct Debit supplier. |
| `drirectDebitMandateID` | `String` | No | The ID of the mandate that was used wehn requesting payment. |
| `errorMessage` | `String` | No |  |
| `errorReason` | `String` | No |  |
| `eventType` | `String` | No |  |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `lightningInvoice` | `String` | No | For Bitcoin Lightning payments this field holds the invoice presented to the payer. |
| `lightningRHash` | `String` | No | For Bitcoin Lightning payments the hash of the invoice presented to the payer. |
| `originUrl` | `String` | No | Optional field that can be set by payment methods, such as pay by bank, that may want to redirect back to the URL that initiated the attempt in the case of a failure condition. |
| `paymentMethodType` | `String` | No | The type of payment method the event relates to, e.g. |
| `paymentProcessorName` | `String` | No | If the event was for a card payment this is the name of the card processor, e.g. |
| `paymentRequestID` | `String` | No |  |
| `pispBankStatus` | `String` | No | For payment initiation attempts some providers (e.g. |
| `pispPaymentInitiationID` | `String` | No | For a payment initiation this is the ID returned by the service provider initiating the payment for us. |
| `pispPaymentInstitutionName` | `String` | No | For a payment initiation this is the name of the financial institution that is used to initiate and authorise the payment. |
| `pispPaymentServiceProviderID` | `String` | No | For a payment initiation this is the service provider ID selected by the payer, typically the ID for the bank or similar financial institution. |
| `pispRedirectUrl` | `String` | No | For a payment initiation this is the redirect URL returned by the service provider initiating the payment for us. |
| `reconciledTransactionID` | `String` | No | For settlement events (only relevant for non-card payments) this is the payin transaction that the payment request event was reconciled with. |
| `refundPayoutID` | `String` | No | ID of the Payout that was created for refund. |
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
| `amount` | `Float` | No | The amount of money to request. |
| `amountPending` | `Float` | No | The amount of money that was authorised but has not arrived in the account yet. |
| `amountReceived` | `Float` | No | The amount of money that has been received for this payment request. |
| `amountRefunded` | `Float` | No | The amount of money that has been refunded for this payment request. |
| `callbackUrl` | `String` | No |  |
| `cardStripePaymentIntentSecret` | `String` | No |  |
| `countryCode` | `String` | No | The country code associated with the payment. |
| `currency` | `String` | No | The currency of the request. |
| `customFieldsToDisplay` | `Array` | No | Custom fields to display to the customer. |
| `description` | `String` | No | An optional description for the payment request. |
| `dueDate` | `String` | No | The due date of the payment request. |
| `fieldDisplaySettings` | `Array` | No |  |
| `googlePayMerchantID` | `String` | No | Merchant ID from Google Pay |
| `id` | `String` | No |  |
| `jwk` | `String` | No | The jwk containing the public key |
| `merchantID` | `String` | No |  |
| `merchantLogoUrlPng` | `String` | No |  |
| `merchantLogoUrlSvg` | `String` | No |  |
| `merchantName` | `String` | No |  |
| `merchantShortName` | `String` | No |  |
| `partialPaymentMethod` | `String` | No |  |
| `paymentAttempts` | `Array` | No | The payment attempts for this payment request. |
| `paymentMethodsList` | `Array` | No | The payment methods that the payment request supports. |
| `paymentProcessor` | `String` | No | The card processor |
| `paymentProcessorKey` | `String` | No | The card processors public key |
| `pispError` | `String` | No | This is the error returned from the bank which is recorded in payment request events. |
| `priorityBankID` | `String` | No |  |
| `status` | `String` | No | The status of the payment request. |
| `stripeAccountID` | `String` | No | Account ID of connected customers in Stripe |
| `title` | `String` | No | The title of the payment request. |

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
| `amount` | `Float` | No | The authorised payment amount. |
| `amountPending` | `Float` | No |  |
| `amountReceived` | `Float` | No |  |
| `amountRefunded` | `Float` | No |  |
| `currency` | `String` | No | The authorised payment currency. |
| `customerID` | `String` | No | The customer id |
| `paymentRequestID` | `String` | No | The ID of the payment request the result is for. |
| `payments` | `Array` | No | The list of payment attempts that have been received for the payment request. |
| `pispAuthorizations` | `Array` | No |  |
| `requestedAmount` | `Float` | No | The full original payment amount requested. |
| `result` | `String` | No | The result of the payment attempt. |

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
| `accountID` | `String` | No | Gets or Sets Account Id of sending account |
| `allowIncomplete` | `Boolean` | No | If set to true the payout will get created even if the business validation rules fail. |
| `amount` | `Float` | No | Gets or Sets payout amount |
| `amountMinorUnits` | `Integer` | No | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `String` | No | This field is used when returning an payout record to a client. |
| `approverID` | `String` | No | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `Array` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `Array` | No | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `Integer` | No | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `Integer` | No | The number of authorisers required for this payout. |
| `batchPayoutID` | `String` | No | The ID of the batch the payout is associated with. |
| `beneficiary` | `Hash` | Yes |  |
| `beneficiaryID` | `String` | No | Optional. |
| `canAuthorise` | `Boolean` | No | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `Boolean` | No | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `Boolean` | No | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `String` | No | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `String` | No |  |
| `createdByEmailAddress` | `String` | No |  |
| `currency` | `String` | No | Gets or Sets Currency of payout request |
| `currentUserID` | `String` | No | The ID of the user that requested access to the PayOut record. |
| `description` | `String` | No | Gets or Sets description of payout request |
| `destination` | `Hash` | No |  |
| `documents` | `Array` | No | Documents associated with the payout. |
| `events` | `Array` | No | The activity associated with the payout. |
| `failedPayouts` | `Hash` | No |  |
| `formattedAmount` | `String` | No | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `String` | No | FX destination currency and amount formatted string. |
| `formattedSchedule` | `String` | No |  |
| `formattedScheduleDayOnly` | `String` | No |  |
| `formattedSourceAccountAvailableBalance` | `String` | No | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `Float` | No | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `Integer` | No | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `String` | No | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `String` | No | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `String` | No | Optional. |
| `fxRate` | `Float` | No | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `Boolean` | No | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `Boolean` | No | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `String` | No | The ID for the payout. |
| `inserted` | `String` | No |  |
| `invoiceID` | `String` | No | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `Boolean` | No | Indicates whether the payout is archived. |
| `isFailed` | `Boolean` | No | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `Boolean` | No | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `Boolean` | No | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `String` | No |  |
| `nonce` | `String` | No |  |
| `paymentProcessor` | `String` | No | The usptream payment processor for the payout. |
| `paymentRail` | `String` | No | Optional field to indicate the payment rail to use for the payout. |
| `payouts` | `Array` | No |  |
| `payrunID` | `String` | No | The ID of the payrun that this payout is associated with. |
| `payrunName` | `String` | No | The name of the payrun that this payout is associated with. |
| `reason` | `String` | No |  |
| `rule` | `Hash` | No |  |
| `scheduleDate` | `String` | No | The date the payout should be submitted. |
| `scheduled` | `Boolean` | No | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `Float` | No | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `Integer` | No | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `String` | No | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `String` | No | The currency of the source account. |
| `sourceAccountIban` | `String` | No | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `Hash` | Yes |  |
| `sourceAccountName` | `String` | No | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `String` | No | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `String` | No | The sort code of the account the payout is being made from. |
| `status` | `String` | No | Gets or Sets the status of payout request |
| `tagIds` | `Array` | No | An optional list of tag ids to add to the payout. |
| `tags` | `Array` | No | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `String` | No | Gets or Sets destination reference ID |
| `topupPayrunID` | `String` | No | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `Float` | No | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `Float` | No | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `Float` | No | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `String` | No | Gets or Sets payout type |
| `userID` | `String` | No | Gets or Sets User ID of who created the payout request |
| `yourReference` | `String` | No | Gets or Sets your reference ID |

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
  "id" => "example_id", # String
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
| `accountID` | `String` | No | Gets or Sets Account Id of sending account |
| `amount` | `Float` | No | Gets or Sets payout amount |
| `amountMinorUnits` | `Integer` | No | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `String` | No | This field is used when returning an payout record to a client. |
| `approverID` | `String` | No | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `Array` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `Array` | No | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `Integer` | No | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `Integer` | No | The number of authorisers required for this payout. |
| `batchPayoutID` | `String` | No | The ID of the batch the payout is associated with. |
| `beneficiary` | `Hash` | Yes |  |
| `canAuthorise` | `Boolean` | No | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `Boolean` | No | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `Boolean` | No | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `String` | No | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `String` | No |  |
| `createdByEmailAddress` | `String` | No |  |
| `currency` | `String` | No | Gets or Sets Currency of payout request |
| `currentUserID` | `String` | No | The ID of the user that requested access to the PayOut record. |
| `description` | `String` | No | Gets or Sets description of payout request |
| `destination` | `Hash` | No |  |
| `documents` | `Array` | No | Documents associated with the payout. |
| `events` | `Array` | No | The activity associated with the payout. |
| `formattedAmount` | `String` | No | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `String` | No | FX destination currency and amount formatted string. |
| `formattedSchedule` | `String` | No |  |
| `formattedScheduleDayOnly` | `String` | No |  |
| `formattedSourceAccountAvailableBalance` | `String` | No | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `Float` | No | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `Integer` | No | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `String` | No | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `String` | No | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `String` | No | Optional. |
| `fxRate` | `Float` | No | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `Boolean` | No | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `Boolean` | No | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `String` | No | The ID for the payout. |
| `inserted` | `String` | No |  |
| `invoiceID` | `String` | No | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `Boolean` | No | Indicates whether the payout is archived. |
| `isFailed` | `Boolean` | No | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `Boolean` | No | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `Boolean` | No | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `String` | No |  |
| `nonce` | `String` | No |  |
| `paymentProcessor` | `String` | No | The usptream payment processor for the payout. |
| `paymentRail` | `String` | No | Optional field to indicate the payment rail to use for the payout. |
| `payrunID` | `String` | No | The ID of the payrun that this payout is associated with. |
| `payrunName` | `String` | No | The name of the payrun that this payout is associated with. |
| `rule` | `Hash` | No |  |
| `scheduleDate` | `String` | No | The date the payout should be submitted. |
| `scheduled` | `Boolean` | No | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `Float` | No | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `Integer` | No | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `String` | No | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `String` | No | The currency of the source account. |
| `sourceAccountIban` | `String` | No | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `Hash` | Yes |  |
| `sourceAccountName` | `String` | No | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `String` | No | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `String` | No | The sort code of the account the payout is being made from. |
| `status` | `String` | No | Gets or Sets the status of payout request |
| `tags` | `Array` | No | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `String` | No | Gets or Sets destination reference ID |
| `topupPayrunID` | `String` | No | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `Float` | No | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `Float` | No | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `Float` | No | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `String` | No | Gets or Sets payout type |
| `userID` | `String` | No | Gets or Sets User ID of who created the payout request |
| `yourReference` | `String` | No | Gets or Sets your reference ID |

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
| `authorisations` | `Array` | No | A list of the users who have successfully authorised the latest version of the payrun and when. |
| `authorisersCompletedCount` | `Integer` | No | The number of distinct authorisers that have authorised the payrun. |
| `authorisersRequiredCount` | `Integer` | No | The number of authorisers required for this payrun. |
| `batchPayoutID` | `String` | No |  |
| `canAuthorise` | `Boolean` | No | True if the payrun can be authorised by the user who loaded it. |
| `canDelete` | `Boolean` | No |  |
| `canEdit` | `Boolean` | No |  |
| `events` | `Array` | No |  |
| `hasCurrentUserAuthorised` | `Boolean` | No | True if the payrun was loaded for a user and that user has already authorised the latest version of the payrun. |
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
| `accountID` | `String` | No | The ID of the account the rule will apply to. |
| `approveUrl` | `String` | No | If set this property holds the URL an approver needs to visit in order to complete a strong authentication check in order to approve the rule. |
| `approverID` | `String` | No |  |
| `authenticationMethods` | `Array` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `Array` | No | A list of the users who have successfully authorised the latest version of the rule and when. |
| `authorisersCompletedCount` | `Integer` | No | The number of distinct authorisers that have authorised the rule. |
| `authorisersRequiredCount` | `Integer` | No | The number of authorisers required for this rule. |
| `canAuthorise` | `Boolean` | No | True if the rule can be authorised by the user who loaded it. |
| `createdBy` | `Hash` | Yes |  |
| `description` | `String` | No | Arbitrary description for the rule. |
| `endAt` | `String` | No | Optional end time for rule executions. |
| `hasCurrentUserAuthorised` | `Boolean` | No | True if the current user has authorised. |
| `id` | `String` | No |  |
| `inserted` | `String` | No |  |
| `isDisabled` | `Boolean` | No | If set to true the rule will be disabled from executing. |
| `lastExecutedAt` | `String` | No |  |
| `lastRunAtTransactionDate` | `String` | No | The most recent transaction date when the rule was last run. |
| `lastUpdated` | `String` | No |  |
| `merchantID` | `String` | No | The ID of the merchant that owns the account. |
| `name` | `String` | No | A name to succinctly describe the rule. |
| `nonce` | `String` | Yes |  |
| `onApprovedWebHookUrl` | `String` | No | Optional URL to receive an HTTP request with the rule details when the rule status changes to approved. |
| `onExecutionErrorWebHookUrl` | `String` | No | Optional URL to receive an HTTP request when a rule execution attempt fails. |
| `onExecutionSuccessWebHookUrl` | `String` | No | Optional URL to receive an HTTP request when a rule execution attempt succeeds. |
| `startAt` | `String` | No | Optional start time for rule executions. |
| `status` | `String` | No |  |
| `sweepAction` | `Hash` | No |  |
| `timeZoneId` | `String` | No | If the rule should be executed on a recurring schedule this is the timezone that the CRON expression should be evaluated in. |
| `triggerCronExpression` | `String` | No | If the rule should be executed on a recurring schedule this is the expression that sets the schedule. |
| `triggerOnPayIn` | `Boolean` | No | Set to true if the rule execution should be triggered when the account receives a pay in (credit). |
| `userID` | `String` | No |  |
| `webHookSecret` | `String` | No | If set this secret will be used to sign Web Hook requests. |

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
| `accountID` | `String` | No | The ID of the account the transaction belongs to. |
| `accountName` | `String` | No | The name of the account the transaction belongs to. |
| `accountSequenceNumber` | `Integer` | No | The sequence number of transaction on a per account basis. |
| `addressDetails` | `Hash` | No |  |
| `amount` | `Float` | No | Amount of the transaction. |
| `amountMinorUnits` | `Integer` | No | Amount of the transaction expressed in the currency’s minor units (e.g. |
| `balance` | `Float` | No | Balance left on the account after the transaction. |
| `balanceMinorUnits` | `Integer` | No | Balance on the account expressed in the currency’s minor units (e.g. |
| `bookingDateTime` | `String` | No |  |
| `chargeDetails` | `Hash` | No |  |
| `content` | `Array` | No |  |
| `counterparty` | `Hash` | No |  |
| `counterpartySummary` | `String` | No | For pay in (credit) transactions this will contain a descriptive string with the most important fields about the counterparty. |
| `currency` | `String` | No | Currency of transaction. |
| `currencyExchange` | `Hash` | No | Provides details on the currency exchange. |
| `date` | `String` | No |  |
| `description` | `String` | No | Description of the transaction. |
| `enrichment` | `Hash` | No |  |
| `fxAmount` | `Float` | No | For an FX payout this is the amound in the FX currency. |
| `fxCurrency` | `String` | No | For an FX payout this is the currency that was received or that was instructed. |
| `fxRate` | `Float` | No | For an FX payout this is the exchange rate between the transaction currency and the FX currency. |
| `grossAmount` | `Hash` | Yes |  |
| `id` | `String` | No | Unique ID for the transaction. |
| `inserted` | `String` | No | Date when the transaction was inserted into the ledger. |
| `isoBankTransactionCode` | `Hash` | No |  |
| `merchant` | `Hash` | No |  |
| `merchantID` | `String` | No | The ID of the merchant that owns the account. |
| `pageNumber` | `Integer` | No | Current page number. |
| `pageSize` | `Integer` | No | Page size |
| `payeeDetails` | `Hash` | Yes | The Payee object contains details of the beneficiary, person or business. |
| `payerDetails` | `Hash` | Yes |  |
| `paymentRequestCustomFields` | `Hash` | No | The custom fields that were attached to the payment request that resulted in this transaction. |
| `paymentRequestID` | `String` | No | For Pay by Bank and Direct Debit transactions this will contain the ID of the payment request. |
| `payoutID` | `String` | No | ID of the payout that resulted in the transaction. |
| `proprietaryBankTransactionCode` | `Hash` | No |  |
| `rawReference` | `String` | No | The raw payment reference details as received from the payment processor. |
| `reference` | `String` | No |  |
| `ruleID` | `String` | No | ID of the rule that resulted in the transaction. |
| `statementReferences` | `Array` | No |  |
| `status` | `String` | No |  |
| `supplementaryData` | `Object` | No |  |
| `tags` | `Array` | No | An optional list of descriptive tags attached to the transaction. |
| `theirReference` | `String` | No | For a pay out the reference that the payer attached for the receiving party. |
| `totalPages` | `Integer` | No | Total pages |
| `totalSize` | `Integer` | No | Total count |
| `transactionAmount` | `Hash` | Yes |  |
| `transactionDate` | `String` | No | Date when the transaction occurred. |
| `transactionInformation` | `Array` | No |  |
| `transactionMutability` | `String` | No |  |
| `type` | `String` | No | Type of the transaction. |
| `valueDateTime` | `String` | No |  |
| `virtualIBAN` | `String` | No | If set it indicates the payin was to a virtual IBAN. |
| `yourReference` | `String` | No | For a pay in the reference the sending party attached. |

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
| `clientSessionTimeouts` | `Array` | No | The number of seconds a session for this user should last before expiring. |
| `emailAddress` | `String` | Yes |  |
| `firstName` | `String` | Yes |  |
| `id` | `String` | No |  |
| `lastName` | `String` | Yes |  |
| `passkeyAdded` | `Boolean` | No |  |
| `permissions` | `Hash` | No |  |
| `profile` | `String` | No |  |
| `rolesWithScope` | `Array` | No |  |
| `twoFactorEnabled` | `Boolean` | No |  |
| `userInviteID` | `String` | No | Optional ID of the invite that is being accepted so the user can be assigned a role on a new merchant. |

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
| `initialRoleID` | `String` | No | The role ID to automatically assign to the merchant’s very first user. |
| `inviteeEmailAddress` | `String` | No | Email address of the user being invited. |
| `inviteeFirstName` | `String` | No | First Name of the user being invited. |
| `inviteeLastName` | `String` | No | Last Name of the user being invited. |
| `inviterEmailAddress` | `String` | No |  |
| `inviterFirstName` | `String` | No |  |
| `inviterLastName` | `String` | No |  |
| `isAuthorised` | `Boolean` | No | Will be set to true once the invite has met the authorisation requirements. |
| `isInviteeRegistered` | `Boolean` | No | If true, indicates the invitee's email address corresponds to an existing MoneyMoov user. |
| `lastInvited` | `String` | No |  |
| `merchantID` | `String` | No | ID of the merchant the user is being invited to. |
| `merchantName` | `String` | No |  |
| `message` | `String` | No |  |
| `registrationUrl` | `String` | No |  |
| `sendInviteEmail` | `Boolean` | No | If set to true an email will be sent to the invitee with instructions on how to accept the invite. |
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
  "id" => "example_id", # String
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
| `accountName` | `String` | No | Name for the account |
| `accountSupplierName` | `String` | No | The payment account supplier name. |
| `availableBalance` | `Float` | No | The current available balance of the account. |
| `availableBalanceMinorUnits` | `Integer` | No | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `Float` | No | Balance of the account. |
| `balanceMinorUnits` | `Integer` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `String` | No | The bank name for external accounts |
| `consentID` | `String` | No | The ID of the consent used to connect the external account. |
| `createdBy` | `Hash` | Yes |  |
| `createdByDisplayName` | `String` | No | Either the name of the user, merchant token or api key that created the account |
| `currency` | `String` | No | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `String` | No | Indicates the default payment rail for this account. |
| `displayName` | `String` | No | Gets a unique display name for the payment account. |
| `expiryDate` | `String` | No | The date that the external account will expire |
| `externalAccountIcon` | `String` | No | The Icon for external accounts |
| `id` | `String` | No | Unique id for the account. |
| `identifier` | `Hash` | Yes |  |
| `inserted` | `String` | No | Timestamp when the account was created. |
| `isArchived` | `Boolean` | No | Indicates whether the account is archived. |
| `isConnectedAccount` | `Boolean` | No | Indicates if the payment account is an externally connected account. |
| `isDefault` | `Boolean` | No | Is the default account |
| `isTrustAccount` | `Boolean` | No | Indicates if the payment account is a trust account. |
| `isVirtual` | `Boolean` | No | True if the account is a virtual account. |
| `lastTransaction` | `Hash` | No |  |
| `lastUpdated` | `String` | No | Timestamp when the account was last updated. |
| `merchantID` | `String` | No | The ID of the merchant that owns the account. |
| `merchantName` | `String` | No | The name of the merchant that owns the account. |
| `name` | `String` | Yes | The name of the virtual account. |
| `physicalAccountID` | `String` | No | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `Array` | No | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `Float` | No | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `Integer` | No | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `String` | No | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `String` | No | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `String` | No | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `String` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `String` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `String` | No |  |
| `xeroBankFeedSyncStatus` | `String` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `Integer` | No | Indicates the number of unsynchronised transactions with Xero |

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
| `destinationUrl` | `String` | No | The destination URL for the webhook. |
| `emailAddress` | `String` | No | The recipient email address(es) for notifications. |
| `failedNotificationEmailAddress` | `String` | No | The email address to which notifications about failed webhook deliveries will be sent. |
| `id` | `String` | No |  |
| `isActive` | `Boolean` | No |  |
| `merchantID` | `String` | No | The ID of the merchant that the webhook is for. |
| `notificationMethod` | `String` | No | The type of notification that will be sent. |
| `resourceTypes` | `Array` | No | The resource types that the webhook will be generated for. |
| `retry` | `Boolean` | No |  |
| `secret` | `String` | No | The secret key required to authenticate webhook notifications. |
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

