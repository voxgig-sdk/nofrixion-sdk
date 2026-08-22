# Nofrixion Ruby SDK



The Ruby SDK for the Nofrixion API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Account` — with named operations (`list`/`load`/`create`/`update`/`remove`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/nofrixion-sdk/releases](https://github.com/voxgig-sdk/nofrixion-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "Nofrixion_sdk"

client = NofrixionSDK.new({
  "apikey" => ENV["NOFRIXION_APIKEY"],
})
```

### 2. List account records

```ruby
begin
  # list returns an Array of Account records — iterate directly.
  accounts = client.Account.list
  accounts.each do |item|
    puts "#{item["id"]} #{item["accountBalances"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```

### 3. Load a cardcustomertoken

CardCustomerToken is nested under customer_email_address, so provide the `customer_email_address`.

```ruby
begin
  # load returns the ENTITY — call data_get for the CardCustomerToken record (raises on error).
  cardcustomertoken = client.CardCustomerToken.load({ "customer_email_address" => "example_customer_email_address" })
  puts cardcustomertoken
rescue => err
  warn "load failed: #{err}"
end
```

### 4. Create, update, and remove

```ruby
# create returns the ENTITY — call data_get for the created Account record.
created = client.Account.create({ "account_id" => "example_account_id", "currency" => "example_currency", "createdBy" => {}, "identifier" => {} })

# Update — index the record via data_get (created.data_get["id"]).
client.Account.update({ "id" => created.data_get["id"], "accountBalances" => [], "accountID" => "example_accountID" })

# Remove
client.Account.remove({ "id" => created.data_get["id"] })
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  consents = client.Consent.list()
rescue => err
  warn "list failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```ruby
client = NofrixionSDK.test({
  "entity" => { "consent" => { "test01" => { "id" => "test01" } } },
})

# Entity ops return the ENTITY (raises on error);
# call data_get for the mock record.
consent = client.Consent.list()
puts consent
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = NofrixionSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
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
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### NofrixionSDK

```ruby
require_relative "Nofrixion_sdk"
client = NofrixionSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `String` | API key for authentication. |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = NofrixionSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### NofrixionSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
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
| `MerchantDirectDebitMandatePage` | `(data) -> MerchantDirectDebitMandatePageEntity` | Create a MerchantDirectDebitMandatePage entity instance. |
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
| `PayoutKeysetPage` | `(data) -> PayoutKeysetPageEntity` | Create a PayoutKeysetPage entity instance. |
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
| `list` | `(reqmatch = nil, ctrl) -> Array` | List entities matching the criteria (call with no argument to list all). Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `NofrixionError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

### Entities

#### Account

| Field | Description |
| --- | --- |
| `accountBalances` | The various balances for the account. |
| `accountID` | ID of the account. |
| `accountIdentifications` | The canoncial identifiers for the account. |
| `accountName` | Name for the account |
| `accountNames` | Optional account names set by the account holder. |
| `accountSupplierName` | The payment account supplier name. |
| `accountType` | The type of account e.g. |
| `availableBalance` | The current available balance of the account. |
| `availableBalanceMinorUnits` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | Balance of the account. |
| `balanceMinorUnits` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | The bank name for external accounts |
| `consentID` | The ID of the consent used to connect the external account. |
| `consolidatedAccountInformation` | Summary information regarding account balances of the overall account provided by the bank. |
| `createdBy` |  |
| `createdByDisplayName` | Either the name of the user, merchant token or api key that created the account |
| `currency` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | Indicates the default payment rail for this account. |
| `description` | Product name as defined by the financial institution for this account. |
| `details` | Supplementary specifications that might be provided by the Bank. |
| `displayName` | Gets a unique display name for the payment account. |
| `expiryDate` | The date that the external account will expire |
| `externalAccountIcon` | The Icon for external accounts |
| `format` | File format to save the statement as. |
| `fromDate` | Minimum transaction date for the statement. |
| `id` | Unique id for the account. |
| `identifier` |  |
| `inserted` | Timestamp when the account was created. |
| `isArchived` | Indicates whether the account is archived. |
| `isConnectedAccount` | Indicates if the payment account is an externally connected account. |
| `isDefault` | Is the default account |
| `isTrustAccount` | Indicates if the payment account is a trust account. |
| `isVirtual` | True if the account is a virtual account. |
| `lastTransaction` |  |
| `lastUpdated` | Timestamp when the account was last updated. |
| `merchantID` | The ID of the merchant that owns the account. |
| `merchantName` | The name of the merchant that owns the account. |
| `nickname` | Nickname of the account that was provided by the account owner. |
| `physicalAccountID` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `roleIDs` | Optional list of role IDs that will get access to the payment account when created. |
| `rules` | The list of rules associated with this account. |
| `submittedPayoutsBalance` | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | Gets a summary of the payments account's most important properties. |
| `supplierPhysicalAccountID` | For internal use only. |
| `supplierSepaInstantStatus` | Indicates the status of the SEPA Instant payment rail for this account. |
| `toDate` | Maximum transaction date for the statement. |
| `type` | Specifies the type of account e.g. |
| `usageType` |  |
| `xeroBankFeedConnectionStatus` | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` |  |
| `xeroBankFeedSyncLastFailedAt` |  |
| `xeroBankFeedSyncLastFailureReason` |  |
| `xeroBankFeedSyncStatus` |  |
| `xeroUnsynchronisedTransactionsCount` | Indicates the number of unsynchronised transactions with Xero |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/accounts/{accountID}/{currency}`

#### Batch

| Field | Description |
| --- | --- |
| `approveUrl` | This field is used when returning a batch payout record to a client. |
| `id` |  |
| `payouts` |  |

Operations: Create, Load.

API path: `/api/v1/payouts/batch`

#### Beneficiary

| Field | Description |
| --- | --- |
| `approvalCallbackUrl` |  |
| `authenticationMethods` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | The number of distinct authorisers that have authorised the beneficiary. |
| `authorisersRequiredCount` | The number of authorisers required for this beneficiary. |
| `beneficiaries` |  |
| `beneficiaryEvents` |  |
| `canAuthorise` | True if the beneficiary can be authorised by the user who loaded it. |
| `canUpdate` | True if the beneficiary can be updated by the user who loaded it. |
| `createdBy` |  |
| `createdByEmailAddress` |  |
| `currency` | Gets or Sets the currency. |
| `destination` |  |
| `failedBeneficiaries` |  |
| `hasCurrentUserAuthorised` | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `id` |  |
| `inserted` |  |
| `isEnabled` |  |
| `lastAuthorised` |  |
| `lastUpdated` |  |
| `merchantID` | Gets or Sets the merchant id. |
| `name` | The descriptive name for the beneficiary. |
| `nonce` |  |
| `sourceAccountIDs` | ID of the accounts which are authorised to act as a source for the beneficiary. |
| `sourceAccounts` |  |
| `theirReference` | The reference that will be used by default as TheirReference when creating payouts to this beneficiary if no TheirReference is specified for the payout. |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/beneficiaries/authorise/{id}`

#### BeneficiaryGroup

| Field | Description |
| --- | --- |
| `groupMembers` | The existing group members. |
| `groupName` | The descriptive name for the beneficiary group. |
| `id` |  |
| `inserted` | Timestamp indicating when the group was created. |
| `lastUpdated` | Timestamp indicating when the group was last updated. |
| `merchantID` | Gets or Sets the merchant id. |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/beneficiarygroups`

#### Card

| Field | Description |
| --- | --- |
| `authorizedAmount` |  |
| `currencyCode` |  |
| `isPayerAuthenticationRequired` | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` |  |
| `requestID` |  |
| `responseCode` |  |
| `responseType` |  |
| `status` |  |
| `threeDSRedirectUrl` | Checkout.com require a redirect for 3DS authentication. |
| `transactionID` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/{id}/card`

#### CardCustomerToken

| Field | Description |
| --- | --- |
| `cardType` | The type of the tokenised card, e.g. |
| `customerEmailAddress` | When creating a tokenised card the payer's email address must be supplied. |
| `expiryMonth` |  |
| `expiryYear` |  |
| `id` | The unique ID of the card token that has been stored for the customer. |
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
| `isPayerAuthenticationRequired` | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` |  |
| `requestID` |  |
| `responseCode` |  |
| `responseType` |  |
| `status` |  |
| `threeDSRedirectUrl` | Checkout.com require a redirect for 3DS authentication. |
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
| `authorisationUrl` | The URL the authorising user needs to be redirected to in order to get the open banking consent token. |
| `callbackUrl` | Optional callback URL that the end user performing the open banking authorisation will be redirected to on completion. |
| `consentID` | The ID of the open banking consent. |
| `emailAddress` | The email address that identifies the end user that will be authorising the open banking consent request. |
| `expiryDate` |  |
| `failureCallbackUrl` | Optional callback URL for open banking consent authorisation failure. |
| `id` |  |
| `inserted` |  |
| `institutionID` | The institution ID the open banking consent is being requested for. |
| `isConnectedAccounts` | Optional setting. |
| `isEnabled` |  |
| `merchantID` | The ID of the merchant the consent token is being created to be used with. |
| `provider` | Lists the supported card and PIS processors. |
| `successWebHookUrl` | A web hook URL to send an HTTP request to when the open banking consent is successfuly authorised. |

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
| `failedSubmissions` | Dictionary of failed submissions, keyed by the index (1-based) in the original request. |
| `successfulSubmissions` | List of successfully submitted direct debit payments. |

Operations: Create.

API path: `/api/v1/paymentrequests/directdebit/batchsubmit`

#### FxRate

| Field | Description |
| --- | --- |
| `destinationCurrency` |  |
| `exchangeRate` | The price at which the transaction will buy the source currency using the destination currency. |
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
| `accountNumber` | Account number of the customer's bank account in case of GBP account. |
| `addressLine1` | First line of the customer's address. |
| `addressLine2` | Second line of the customer's address. |
| `approvedAt` | Date at which the supplier approved this mandate. |
| `city` | Customer's city. |
| `countryCode` | 2-character country code of the customer's bank account. |
| `currency` | Currency of this mandate. |
| `customerAccountNumber` | Customer's account number in case of GBP account. |
| `customerCity` | Customer's city of residence. |
| `customerCountryCode` | Customer's country of residence code. |
| `customerCountryName` | Customer's country of residence. |
| `customerEmailAddress` | Customer's email address. |
| `customerFirstName` | Customer's first name. |
| `customerIban` | Customer's IBAN in case of EUR account. |
| `customerLastName` | Customer's last name. |
| `customerSortCode` | Customer's sort code in case of GBP account. |
| `emailAddress` | Customer's email address. |
| `firstName` | Customer's first name. |
| `iban` | IBAN of the customer's bank account in case of EUR account. |
| `id` | Internal ID of the mandate. |
| `inserted` | The timestamp this mandate was created at. |
| `isRecurring` | Whether this mandate is single-use or recurring. |
| `lastName` | Customer's last name. |
| `lastUpdated` | The timestamp this mandate was last updated at. |
| `merchantID` | Internal ID of this mandate's merchant. |
| `postalCode` | Customer's postal code. |
| `reference` | Reference assigned to this mandate. |
| `sortCode` | Sort code of the customer's bank account in case of GBP account. |
| `status` | General status of this mandate. |
| `supplierBankAccountID` | ID that the supplier assigned to this mandate's bank account. |
| `supplierCustomerID` | ID that the supplier assigned to this mandate's customer. |
| `supplierMandateID` | ID that the supplier assigned to this mandate. |
| `supplierName` | Name of the supplier used to create this mandate. |
| `supplierStatus` | Last status that the supplier reported for this mandate. |

Operations: Create, Load.

API path: `/api/v1/mandates`

#### Merchant

| Field | Description |
| --- | --- |
| `accountCurrencies` | The list of currencies that the merchant has accounts for. |
| `canHaveTrustAccounts` | Trust accounts are a special type of account that allow the account name to be trusted for use in statements and verification of payee checks. |
| `cardPaymentProcessor` | Name of the card payment processor. |
| `companyID` | The Company ID recorded in the Compliance system. |
| `displayQrOnHostedPay` | Indicates if a QR Code containing the payment link should be displayed on the hosted payment page. |
| `hostedPayVersion` | The version of the hosted payment page to use with the merchant. |
| `id` | Unique ID for the merchant. |
| `inserted` | Timestamp the merchant was added to MoneyMoov. |
| `isBlocked` | The merchant is blocked from making payments (payouts). |
| `isExited` | The merchant has formally terminated their relationship and is no longer a customer. |
| `isSuspended` | The merchant has temporarily suspended their own account. |
| `jurisdiction` | The jurisdiction the merchant entity is incorporated or established in. |
| `logoUrlPng` | The CDN URL of the merchant's logo in PNG format. |
| `logoUrlSvg` | The CDN URL of the merchant's logo in SVG format. |
| `merchantCategoryCode` | The industry code that represents the merchant's primary trading activity. |
| `name` | The registered business name of the merchant. |
| `notes` | The notes field is an optional free text field that can be used to store any additional information about the merchant. |
| `parentMerchant` |  |
| `paymentAccountLimit` | The maximum number of payment accounts that can be created for the Merchant. |
| `paymentAccounts` |  |
| `reason` | The reason for the suspension. |
| `shortName` | A URL friendly shortish name for the merchant. |
| `supportedPaymentMethodsList` | The payment methods that are configured and supported for this merchant. |
| `suspensionReason` | The reason for the suspension, provided by the merchant. |
| `tags` | An optional list of descriptive tags that can be used on merchant entities such as payment requests. |
| `timeZoneId` | The IANA (Internet Assigned Numbers Authority) time zone identifier of the merchant. |
| `tradingName` | An optional trading name. |
| `webHookLimit` | The maximum number of web hooks that can be created for the Merchant. |
| `yourRoleName` | The name of the role for the identity that loaded the merchant record. |

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
| `approvedAt` | Date at which the supplier approved this mandate. |
| `currency` | Currency of this mandate. |
| `customerAccountNumber` | Customer's account number in case of GBP account. |
| `customerCity` | Customer's city of residence. |
| `customerCountryCode` | Customer's country of residence code. |
| `customerCountryName` | Customer's country of residence. |
| `customerEmailAddress` | Customer's email address. |
| `customerFirstName` | Customer's first name. |
| `customerIban` | Customer's IBAN in case of EUR account. |
| `customerLastName` | Customer's last name. |
| `customerSortCode` | Customer's sort code in case of GBP account. |
| `id` | Internal ID of the mandate. |
| `inserted` | The timestamp this mandate was created at. |
| `isRecurring` | Whether this mandate is single-use or recurring. |
| `lastUpdated` | The timestamp this mandate was last updated at. |
| `merchantID` | Internal ID of this mandate's merchant. |
| `reference` | Reference assigned to this mandate. |
| `status` | General status of this mandate. |
| `supplierBankAccountID` | ID that the supplier assigned to this mandate's bank account. |
| `supplierCustomerID` | ID that the supplier assigned to this mandate's customer. |
| `supplierMandateID` | ID that the supplier assigned to this mandate. |
| `supplierName` | Name of the supplier used to create this mandate. |
| `supplierStatus` | Last status that the supplier reported for this mandate. |

Operations: List.

API path: `/api/v1/mandates`

#### MerchantPayByBankSetting

| Field | Description |
| --- | --- |
| `bankCountryCodes` | The list of country codes representing the banks the country supports. |
| `bankID` | ID of the bank to be configured for the merchant. |
| `bankName` | Name of the Bank/Institution. |
| `businessInstitutionID` | ID that the processor uses to identify the bank (business accounts). |
| `currency` | Currency supported by the bank. |
| `logo` | URL of the bank's logo. |
| `message` | Message relating to specific bank. |
| `messageImageUrl` | Optional image URL to be displayed with the message. |
| `order` | Order in which this setting will appear in the UI. |
| `personalInstitutionID` | ID that the processor uses to identify the bank (personal accounts). |
| `processor` | Name of the bank payment processor. |
| `warningHeading` | The heading for a warning message related to the bank institution to be displayed to the user. |
| `warningMessage` | The warning message related to the bank institution to be displayed to the user. |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/banksettings`

#### MerchantPaymentRequestTemplate

| Field | Description |
| --- | --- |
| `bankPaymentOptions` |  |
| `cardPaymentAddressOptions` |  |
| `cardPaymentCaptureOptions` |  |
| `customFields` | A list of custom fields that can be included in the payment request template. |
| `defaultFields` | A list of default fields that are included in the payment request template. |
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
| `authenticationMethods` | A list of authentication types allowed to authorise the merchant token. |
| `authorisations` | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | The number of distinct authorisers that have authorised the merchant token. |
| `authorisersRequiredCount` | The number of authorisers required for this merchant token. |
| `canAuthorise` | True if the merchant token can be authorised by the user who loaded it. |
| `description` | Token description |
| `expiresAt` | Optional. |
| `hasCurrentUserAuthorised` | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `hmacAlgorithm` | Optional shared secret algorithm to use for HMAC authentication. |
| `id` |  |
| `inserted` |  |
| `ipAddressWhitelist` | Optional. |
| `isArchived` | Indicates whether the merchant token is archived. |
| `isEnabled` | If set to false the merchant token will not be accepted to authorise a request. |
| `lastAuthorised` |  |
| `lastUpdated` |  |
| `merchantID` | The merchant id to add to the token |
| `nonce` |  |
| `permissionTypes` | The permissions that the merchant token supports. |
| `requestSignatureVersion` | Represent the version of the overall merchant token. |
| `sharedSecretAlgorithm` | Optional shared secret algorithm to use for HMAC authentication. |
| `sharedSecretBase64` | The base 64 encoded shared secret that is used for request authentication with an HMAC. |
| `token` | The JWT merchant token. |

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
| `accountName` | The name of the account to verify |
| `accountNumber` | The account number of the account to verify (for CoP checks) |
| `iban` | The IBAN of the account to verify (for VoP checks) |
| `payeeVerifiedAccountName` | The verified account name of the payee, if available (in case of a close match) |
| `result` | The result of the payee verification |
| `secondaryIdentification` | Optional secondary identifier for the account to verify. |
| `sortCode` | The sort code of the account to verify (for CoP checks) |

Operations: Create.

API path: `/api/v1/openbanking/payeeverification`

#### Payment

| Field | Description |
| --- | --- |
| `addresses` |  |
| `amount` | The amount of money to request. |
| `amountPending` | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | Total amount received for this payment request. |
| `amountRefunded` | Total amount refunded for this payment request. |
| `autoSendReceipt` | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardNoPayerAuthentication` | If set to true for card payments no attempt will be made to use payer authentication (3-D Secure and equivalent). |
| `cardProcessorMerchantID` | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `cardTransmitRawDetails` | If set to true for card payments the sensitive card number and card verification number will be transmitted directly rather than being tokenised. |
| `createdByUser` |  |
| `currency` | The currency of the request. |
| `customFields` | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | Optional email address for the customer. |
| `customerID` | An optional customer identifier for the payment request. |
| `customerName` |  |
| `description` | An optional description for the payment request. |
| `destinationAccount` |  |
| `directDebitPayment` | Contains information about a Direct Debit payment attempt for a payment request. |
| `dueDate` | The due date for the payment request. |
| `events` |  |
| `failureCallbackUrl` | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` |  |
| `hostedPayCheckoutUrl` | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` |  |
| `ignoreAddressVerification` | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | The timestamp the payment request was created at. |
| `insertedSortable` | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `isArchived` | Indicates whether the payment request is archived. |
| `jwk` | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | The timestamp the payment request was last updated at. |
| `lightningInvoice` | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | The ID of the merchant to create the payment request for. |
| `merchantTokenDescription` | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` |  |
| `notificationRoleIDs` | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | An optional order ID for the payment request. |
| `partialPaymentMethod` | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | The payment attempts made against this payment request. |
| `paymentMethods` | The payment methods that the payment request supports. |
| `paymentProcessor` | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `payrunID` | The ID of a payrun that needs an account top up. |
| `pispAccountID` | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` |  |
| `sandboxSettleDelayInSeconds` | Sandbox only. |
| `shippingAddress` |  |
| `shippingAddressCity` | Optionally the city of the customer's shipping address. |
| `shippingAddressCountryCode` | Optionally the country code of the customer's shipping address. |
| `shippingAddressCounty` | Optionally the state or county of the customer's shipping address. |
| `shippingAddressLine1` | Optionally the first line of the customer's shipping address. |
| `shippingAddressLine2` | Optionally the second line of the customer's shipping address. |
| `shippingAddressPostCode` | Optionally the post code of the customer's shipping address. |
| `shippingEmail` | Optionally the shipping email address for the customer. |
| `shippingFirstName` | Optionally the first name of the customer's shipping address. |
| `shippingLastName` | Optionally the last name of the customer's shipping address. |
| `shippingPhone` | Optionally the shipping phone number for the customer. |
| `status` | The current status of the payment request. |
| `successWebHookUrl` | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tagIds` | An optional list of tag ids to add to the payment request |
| `tags` | An optional list of descriptive tags attached to the payment request. |
| `title` | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` |  |
| `transactions` |  |
| `useHostedPaymentPage` | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

Operations: Create, Load, Update.

API path: `/api/v1/paymentrequests`

#### PaymentAccount

| Field | Description |
| --- | --- |
| `accountName` | Name for the account |
| `accountSupplierName` | The payment account supplier name. |
| `availableBalance` | The current available balance of the account. |
| `availableBalanceMinorUnits` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | Balance of the account. |
| `balanceMinorUnits` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | The bank name for external accounts |
| `consentID` | The ID of the consent used to connect the external account. |
| `createdBy` |  |
| `createdByDisplayName` | Either the name of the user, merchant token or api key that created the account |
| `currency` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | Indicates the default payment rail for this account. |
| `displayName` | Gets a unique display name for the payment account. |
| `expiryDate` | The date that the external account will expire |
| `externalAccountIcon` | The Icon for external accounts |
| `id` | Unique id for the account. |
| `identifier` |  |
| `inserted` | Timestamp when the account was created. |
| `isArchived` | Indicates whether the account is archived. |
| `isConnectedAccount` | Indicates if the payment account is an externally connected account. |
| `isDefault` | Is the default account |
| `isTrustAccount` | Indicates if the payment account is a trust account. |
| `isVirtual` | True if the account is a virtual account. |
| `lastTransaction` |  |
| `lastUpdated` | Timestamp when the account was last updated. |
| `merchantID` | The ID of the merchant that owns the account. |
| `merchantName` | The name of the merchant that owns the account. |
| `physicalAccountID` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | The list of rules associated with this account. |
| `submittedPayoutsBalance` | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` |  |
| `xeroBankFeedSyncLastFailedAt` |  |
| `xeroBankFeedSyncLastFailureReason` |  |
| `xeroBankFeedSyncStatus` |  |
| `xeroUnsynchronisedTransactionsCount` | Indicates the number of unsynchronised transactions with Xero |

Operations: List.

API path: `/api/v1/accounts/paged`

#### PaymentAccountMinimal

| Field | Description |
| --- | --- |
| `accountName` | Name for the account |
| `availableBalance` | The current available balance of the account. |
| `balance` | Balance of the account. |
| `balanceMinorUnits` | Balance of the account expressed in the currency’s minor units (e.g. |
| `currency` | Currency of the account in ISO 4217 format |
| `id` | Unique id for the account. |
| `identifier` |  |
| `isArchived` | Is the account archived |
| `isConnectedAccount` | Indicates if the payment account is an externally connected account. |
| `merchantID` | The ID of the merchant that owns the account. |
| `submittedPayoutsBalance` | Total of the payouts that have been submitted for processing. |

Operations: List.

API path: `/api/v1/accounts/minimal`

#### PaymentInitiation

| Field | Description |
| --- | --- |
| `paymentInitiationID` | The unique identifier of the payment initiation request. |
| `paymentRequestCallbackUrl` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` |  |
| `redirectUrl` | A redirect URL for the user to authorise the payment initiation request at the ASPSP |
| `responseType` |  |
| `specificErrorMessage` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/{id}/pisp`

#### PaymentRequest

| Field | Description |
| --- | --- |
| `addresses` |  |
| `amount` | The amount of money to request. |
| `amountPending` | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | Total amount received for this payment request. |
| `amountRefunded` | Total amount refunded for this payment request. |
| `autoSendReceipt` | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardProcessorMerchantID` | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `createdByUser` |  |
| `currency` | The currency of the request. |
| `customFields` | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | Optional email address for the customer. |
| `customerID` | An optional customer identifier for the payment request. |
| `customerName` |  |
| `description` | An optional description for the payment request. |
| `destinationAccount` |  |
| `directDebitPayment` | Contains information about a Direct Debit payment attempt for a payment request. |
| `doSimulateSettlementFailure` |  |
| `dueDate` | The due date for the payment request. |
| `errorDescription` |  |
| `events` |  |
| `failedPaymentRequests` |  |
| `failureCallbackUrl` | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` |  |
| `hostedPayCheckoutUrl` | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` |  |
| `ignoreAddressVerification` | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | The timestamp the payment request was created at. |
| `insertedSortable` | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `institution` |  |
| `isArchived` | Indicates whether the payment request is archived. |
| `jwk` | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | The timestamp the payment request was last updated at. |
| `lightningInvoice` | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` |  |
| `merchantTokenDescription` | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` |  |
| `notificationRoleIDs` | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | An optional order ID for the payment request. |
| `partialPaymentMethod` | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | The payment attempts made against this payment request. |
| `paymentInitiationID` |  |
| `paymentMethods` | The payment methods that the payment request supports. |
| `paymentProcessor` | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `paymentRequests` |  |
| `payrunID` | The ID of a payrun that needs an account top up. |
| `pispAccountID` | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` |  |
| `sandboxSettleDelayInSeconds` | Sandbox only. |
| `shippingAddress` |  |
| `status` | The current status of the payment request. |
| `successWebHookUrl` | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tags` | An optional list of descriptive tags attached to the payment request. |
| `title` | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` |  |
| `transactions` |  |
| `useHostedPaymentPage` | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/paymentrequests/{id}/directdebit`

#### PaymentRequestEvent

| Field | Description |
| --- | --- |
| `amount` |  |
| `applePayTransactionID` | Transaction ID received in Apple pay token. |
| `cardAuthorizationResponseID` | For a successful card authorization this field will hold the response ID. |
| `cardExpiryMonth` | For card payment events this field holds the payer's card expiry month. |
| `cardExpiryYear` | For card payment events this field holds the payer's card expiry year. |
| `cardIssuer` | For card payment events this field holds the payer's card issuer. |
| `cardIssuerCountry` | For card payment events this field holds the payer's card issuer country of origin. |
| `cardLastFourDigits` | For card payment events this field holds the payer's card last four digits. |
| `cardRequestID` |  |
| `cardScheme` | For card payment events this field holds the scheme of the payer's card, e.g. |
| `cardTokenCustomerID` | If the option to create a reusable token for card payments was set this field contains the token the merchant can store to use for repeat payments. |
| `cardTransactionID` |  |
| `currency` |  |
| `directDebitPaymentID` | Payment ID issued by the Direct Debit supplier. |
| `directDebitPaymentReference` | Reference string issued by the Direct Debit supplier. |
| `drirectDebitMandateID` | The ID of the mandate that was used wehn requesting payment. |
| `errorMessage` |  |
| `errorReason` |  |
| `eventType` |  |
| `id` |  |
| `inserted` |  |
| `lightningInvoice` | For Bitcoin Lightning payments this field holds the invoice presented to the payer. |
| `lightningRHash` | For Bitcoin Lightning payments the hash of the invoice presented to the payer. |
| `originUrl` | Optional field that can be set by payment methods, such as pay by bank, that may want to redirect back to the URL that initiated the attempt in the case of a failure condition. |
| `paymentMethodType` | The type of payment method the event relates to, e.g. |
| `paymentProcessorName` | If the event was for a card payment this is the name of the card processor, e.g. |
| `paymentRequestID` |  |
| `pispBankStatus` | For payment initiation attempts some providers (e.g. |
| `pispPaymentInitiationID` | For a payment initiation this is the ID returned by the service provider initiating the payment for us. |
| `pispPaymentInstitutionName` | For a payment initiation this is the name of the financial institution that is used to initiate and authorise the payment. |
| `pispPaymentServiceProviderID` | For a payment initiation this is the service provider ID selected by the payer, typically the ID for the bank or similar financial institution. |
| `pispRedirectUrl` | For a payment initiation this is the redirect URL returned by the service provider initiating the payment for us. |
| `reconciledTransactionID` | For settlement events (only relevant for non-card payments) this is the payin transaction that the payment request event was reconciled with. |
| `refundPayoutID` | ID of the Payout that was created for refund. |
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
| `amount` | The amount of money to request. |
| `amountPending` | The amount of money that was authorised but has not arrived in the account yet. |
| `amountReceived` | The amount of money that has been received for this payment request. |
| `amountRefunded` | The amount of money that has been refunded for this payment request. |
| `callbackUrl` |  |
| `cardStripePaymentIntentSecret` |  |
| `countryCode` | The country code associated with the payment. |
| `currency` | The currency of the request. |
| `customFieldsToDisplay` | Custom fields to display to the customer. |
| `description` | An optional description for the payment request. |
| `dueDate` | The due date of the payment request. |
| `fieldDisplaySettings` |  |
| `googlePayMerchantID` | Merchant ID from Google Pay |
| `id` |  |
| `jwk` | The jwk containing the public key |
| `merchantID` |  |
| `merchantLogoUrlPng` |  |
| `merchantLogoUrlSvg` |  |
| `merchantName` |  |
| `merchantShortName` |  |
| `partialPaymentMethod` |  |
| `paymentAttempts` | The payment attempts for this payment request. |
| `paymentMethodsList` | The payment methods that the payment request supports. |
| `paymentProcessor` | The card processor |
| `paymentProcessorKey` | The card processors public key |
| `pispError` | This is the error returned from the bank which is recorded in payment request events. |
| `priorityBankID` |  |
| `status` | The status of the payment request. |
| `stripeAccountID` | Account ID of connected customers in Stripe |
| `title` | The title of the payment request. |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/minimal`

#### PaymentRequestResult

| Field | Description |
| --- | --- |
| `amount` | The authorised payment amount. |
| `amountPending` |  |
| `amountReceived` |  |
| `amountRefunded` |  |
| `currency` | The authorised payment currency. |
| `customerID` | The customer id |
| `paymentRequestID` | The ID of the payment request the result is for. |
| `payments` | The list of payment attempts that have been received for the payment request. |
| `pispAuthorizations` |  |
| `requestedAmount` | The full original payment amount requested. |
| `result` | The result of the payment attempt. |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/result`

#### Payout

| Field | Description |
| --- | --- |
| `accountID` | Gets or Sets Account Id of sending account |
| `allowIncomplete` | If set to true the payout will get created even if the business validation rules fail. |
| `amount` | Gets or Sets payout amount |
| `amountMinorUnits` | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | This field is used when returning an payout record to a client. |
| `approverID` | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | The number of authorisers required for this payout. |
| `batchPayoutID` | The ID of the batch the payout is associated with. |
| `beneficiary` |  |
| `beneficiaryID` | Optional. |
| `canAuthorise` | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | Optional field to set who should pay any fees for the payout. |
| `createdBy` |  |
| `createdByEmailAddress` |  |
| `currency` | Gets or Sets Currency of payout request |
| `currentUserID` | The ID of the user that requested access to the PayOut record. |
| `description` | Gets or Sets description of payout request |
| `destination` |  |
| `documents` | Documents associated with the payout. |
| `events` | The activity associated with the payout. |
| `failedPayouts` |  |
| `formattedAmount` | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | FX destination currency and amount formatted string. |
| `formattedSchedule` |  |
| `formattedScheduleDayOnly` |  |
| `formattedSourceAccountAvailableBalance` | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | Optional. |
| `fxRate` | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | The ID for the payout. |
| `inserted` |  |
| `invoiceID` | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | Indicates whether the payout is archived. |
| `isFailed` | Set to true if a submitted payout subsequently fails. |
| `isSettled` | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` |  |
| `merchantID` | The ID of the merchant that owns the account. |
| `merchantTokenDescription` |  |
| `nonce` |  |
| `paymentProcessor` | The usptream payment processor for the payout. |
| `paymentRail` | Optional field to indicate the payment rail to use for the payout. |
| `payouts` |  |
| `payrunID` | The ID of the payrun that this payout is associated with. |
| `payrunName` | The name of the payrun that this payout is associated with. |
| `reason` |  |
| `rule` |  |
| `scheduleDate` | The date the payout should be submitted. |
| `scheduled` | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | The currency of the source account. |
| `sourceAccountIban` | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` |  |
| `sourceAccountName` | The name of the account the payout is being made from. |
| `sourceAccountNumber` | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | The sort code of the account the payout is being made from. |
| `status` | Gets or Sets the status of payout request |
| `tagIds` | An optional list of tag ids to add to the payout. |
| `tags` | An optional list of descriptive tags attached to the payout. |
| `theirReference` | Gets or Sets destination reference ID |
| `topupPayrunID` | The ID of a payrun that needs an account top up. |
| `transactedAmount` | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | Gets or Sets payout type |
| `userID` | Gets or Sets User ID of who created the payout request |
| `yourReference` | Gets or Sets your reference ID |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/payouts/batch/submit/{id}`

#### PayoutKeysetPage

| Field | Description |
| --- | --- |
| `accountID` | Gets or Sets Account Id of sending account |
| `amount` | Gets or Sets payout amount |
| `amountMinorUnits` | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | This field is used when returning an payout record to a client. |
| `approverID` | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | The number of authorisers required for this payout. |
| `batchPayoutID` | The ID of the batch the payout is associated with. |
| `beneficiary` |  |
| `canAuthorise` | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | Optional field to set who should pay any fees for the payout. |
| `createdBy` |  |
| `createdByEmailAddress` |  |
| `currency` | Gets or Sets Currency of payout request |
| `currentUserID` | The ID of the user that requested access to the PayOut record. |
| `description` | Gets or Sets description of payout request |
| `destination` |  |
| `documents` | Documents associated with the payout. |
| `events` | The activity associated with the payout. |
| `formattedAmount` | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | FX destination currency and amount formatted string. |
| `formattedSchedule` |  |
| `formattedScheduleDayOnly` |  |
| `formattedSourceAccountAvailableBalance` | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | Optional. |
| `fxRate` | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | The ID for the payout. |
| `inserted` |  |
| `invoiceID` | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | Indicates whether the payout is archived. |
| `isFailed` | Set to true if a submitted payout subsequently fails. |
| `isSettled` | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` |  |
| `merchantID` | The ID of the merchant that owns the account. |
| `merchantTokenDescription` |  |
| `nonce` |  |
| `paymentProcessor` | The usptream payment processor for the payout. |
| `paymentRail` | Optional field to indicate the payment rail to use for the payout. |
| `payrunID` | The ID of the payrun that this payout is associated with. |
| `payrunName` | The name of the payrun that this payout is associated with. |
| `rule` |  |
| `scheduleDate` | The date the payout should be submitted. |
| `scheduled` | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | The currency of the source account. |
| `sourceAccountIban` | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` |  |
| `sourceAccountName` | The name of the account the payout is being made from. |
| `sourceAccountNumber` | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | The sort code of the account the payout is being made from. |
| `status` | Gets or Sets the status of payout request |
| `tags` | An optional list of descriptive tags attached to the payout. |
| `theirReference` | Gets or Sets destination reference ID |
| `topupPayrunID` | The ID of a payrun that needs an account top up. |
| `transactedAmount` | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | Gets or Sets payout type |
| `userID` | Gets or Sets User ID of who created the payout request |
| `yourReference` | Gets or Sets your reference ID |

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
| `authorisations` | A list of the users who have successfully authorised the latest version of the payrun and when. |
| `authorisersCompletedCount` | The number of distinct authorisers that have authorised the payrun. |
| `authorisersRequiredCount` | The number of authorisers required for this payrun. |
| `batchPayoutID` |  |
| `canAuthorise` | True if the payrun can be authorised by the user who loaded it. |
| `canDelete` |  |
| `canEdit` |  |
| `events` |  |
| `hasCurrentUserAuthorised` | True if the payrun was loaded for a user and that user has already authorised the latest version of the payrun. |
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
| `accountID` | The ID of the account the rule will apply to. |
| `approveUrl` | If set this property holds the URL an approver needs to visit in order to complete a strong authentication check in order to approve the rule. |
| `approverID` |  |
| `authenticationMethods` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | A list of the users who have successfully authorised the latest version of the rule and when. |
| `authorisersCompletedCount` | The number of distinct authorisers that have authorised the rule. |
| `authorisersRequiredCount` | The number of authorisers required for this rule. |
| `canAuthorise` | True if the rule can be authorised by the user who loaded it. |
| `createdBy` |  |
| `description` | Arbitrary description for the rule. |
| `endAt` | Optional end time for rule executions. |
| `hasCurrentUserAuthorised` | True if the current user has authorised. |
| `id` |  |
| `inserted` |  |
| `isDisabled` | If set to true the rule will be disabled from executing. |
| `lastExecutedAt` |  |
| `lastRunAtTransactionDate` | The most recent transaction date when the rule was last run. |
| `lastUpdated` |  |
| `merchantID` | The ID of the merchant that owns the account. |
| `name` | A name to succinctly describe the rule. |
| `nonce` |  |
| `onApprovedWebHookUrl` | Optional URL to receive an HTTP request with the rule details when the rule status changes to approved. |
| `onExecutionErrorWebHookUrl` | Optional URL to receive an HTTP request when a rule execution attempt fails. |
| `onExecutionSuccessWebHookUrl` | Optional URL to receive an HTTP request when a rule execution attempt succeeds. |
| `startAt` | Optional start time for rule executions. |
| `status` |  |
| `sweepAction` |  |
| `timeZoneId` | If the rule should be executed on a recurring schedule this is the timezone that the CRON expression should be evaluated in. |
| `triggerCronExpression` | If the rule should be executed on a recurring schedule this is the expression that sets the schedule. |
| `triggerOnPayIn` | Set to true if the rule execution should be triggered when the account receives a pay in (credit). |
| `userID` |  |
| `webHookSecret` | If set this secret will be used to sign Web Hook requests. |

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
| `accountID` | The ID of the account the transaction belongs to. |
| `accountName` | The name of the account the transaction belongs to. |
| `accountSequenceNumber` | The sequence number of transaction on a per account basis. |
| `addressDetails` |  |
| `amount` | Amount of the transaction. |
| `amountMinorUnits` | Amount of the transaction expressed in the currency’s minor units (e.g. |
| `balance` | Balance left on the account after the transaction. |
| `balanceMinorUnits` | Balance on the account expressed in the currency’s minor units (e.g. |
| `bookingDateTime` |  |
| `chargeDetails` |  |
| `content` |  |
| `counterparty` |  |
| `counterpartySummary` | For pay in (credit) transactions this will contain a descriptive string with the most important fields about the counterparty. |
| `currency` | Currency of transaction. |
| `currencyExchange` | Provides details on the currency exchange. |
| `date` |  |
| `description` | Description of the transaction. |
| `enrichment` |  |
| `fxAmount` | For an FX payout this is the amound in the FX currency. |
| `fxCurrency` | For an FX payout this is the currency that was received or that was instructed. |
| `fxRate` | For an FX payout this is the exchange rate between the transaction currency and the FX currency. |
| `grossAmount` |  |
| `id` | Unique ID for the transaction. |
| `inserted` | Date when the transaction was inserted into the ledger. |
| `isoBankTransactionCode` |  |
| `merchant` |  |
| `merchantID` | The ID of the merchant that owns the account. |
| `pageNumber` | Current page number. |
| `pageSize` | Page size |
| `payeeDetails` | The Payee object contains details of the beneficiary, person or business. |
| `payerDetails` |  |
| `paymentRequestCustomFields` | The custom fields that were attached to the payment request that resulted in this transaction. |
| `paymentRequestID` | For Pay by Bank and Direct Debit transactions this will contain the ID of the payment request. |
| `payoutID` | ID of the payout that resulted in the transaction. |
| `proprietaryBankTransactionCode` |  |
| `rawReference` | The raw payment reference details as received from the payment processor. |
| `reference` |  |
| `ruleID` | ID of the rule that resulted in the transaction. |
| `statementReferences` |  |
| `status` |  |
| `supplementaryData` |  |
| `tags` | An optional list of descriptive tags attached to the transaction. |
| `theirReference` | For a pay out the reference that the payer attached for the receiving party. |
| `totalPages` | Total pages |
| `totalSize` | Total count |
| `transactionAmount` |  |
| `transactionDate` | Date when the transaction occurred. |
| `transactionInformation` |  |
| `transactionMutability` |  |
| `type` | Type of the transaction. |
| `valueDateTime` |  |
| `virtualIBAN` | If set it indicates the payin was to a virtual IBAN. |
| `yourReference` | For a pay in the reference the sending party attached. |

Operations: Create, List, Load, Remove.

API path: `/api/v1/transactions/{id}/tags`

#### User

| Field | Description |
| --- | --- |
| `clientSessionTimeouts` | The number of seconds a session for this user should last before expiring. |
| `emailAddress` |  |
| `firstName` |  |
| `id` |  |
| `lastName` |  |
| `passkeyAdded` |  |
| `permissions` |  |
| `profile` |  |
| `rolesWithScope` |  |
| `twoFactorEnabled` |  |
| `userInviteID` | Optional ID of the invite that is being accepted so the user can be assigned a role on a new merchant. |

Operations: List, Update.

API path: `/api/v1/user/{merchantID}/userspaged`

#### UserInvite

| Field | Description |
| --- | --- |
| `authorisationStatus` |  |
| `failedUserInvites` |  |
| `id` |  |
| `initialRoleID` | The role ID to automatically assign to the merchant’s very first user. |
| `inviteeEmailAddress` | Email address of the user being invited. |
| `inviteeFirstName` | First Name of the user being invited. |
| `inviteeLastName` | Last Name of the user being invited. |
| `inviterEmailAddress` |  |
| `inviterFirstName` |  |
| `inviterLastName` |  |
| `isAuthorised` | Will be set to true once the invite has met the authorisation requirements. |
| `isInviteeRegistered` | If true, indicates the invitee's email address corresponds to an existing MoneyMoov user. |
| `lastInvited` |  |
| `merchantID` | ID of the merchant the user is being invited to. |
| `merchantName` |  |
| `message` |  |
| `registrationUrl` |  |
| `sendInviteEmail` | If set to true an email will be sent to the invitee with instructions on how to accept the invite. |
| `status` |  |
| `user` |  |
| `userID` |  |
| `userInvites` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/userinvites/authorise/{id}`

#### Virtual

| Field | Description |
| --- | --- |
| `accountName` | Name for the account |
| `accountSupplierName` | The payment account supplier name. |
| `availableBalance` | The current available balance of the account. |
| `availableBalanceMinorUnits` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | Balance of the account. |
| `balanceMinorUnits` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | The bank name for external accounts |
| `consentID` | The ID of the consent used to connect the external account. |
| `createdBy` |  |
| `createdByDisplayName` | Either the name of the user, merchant token or api key that created the account |
| `currency` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | Indicates the default payment rail for this account. |
| `displayName` | Gets a unique display name for the payment account. |
| `expiryDate` | The date that the external account will expire |
| `externalAccountIcon` | The Icon for external accounts |
| `id` | Unique id for the account. |
| `identifier` |  |
| `inserted` | Timestamp when the account was created. |
| `isArchived` | Indicates whether the account is archived. |
| `isConnectedAccount` | Indicates if the payment account is an externally connected account. |
| `isDefault` | Is the default account |
| `isTrustAccount` | Indicates if the payment account is a trust account. |
| `isVirtual` | True if the account is a virtual account. |
| `lastTransaction` |  |
| `lastUpdated` | Timestamp when the account was last updated. |
| `merchantID` | The ID of the merchant that owns the account. |
| `merchantName` | The name of the merchant that owns the account. |
| `name` | The name of the virtual account. |
| `physicalAccountID` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | The list of rules associated with this account. |
| `submittedPayoutsBalance` | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` |  |
| `xeroBankFeedSyncLastFailedAt` |  |
| `xeroBankFeedSyncLastFailureReason` |  |
| `xeroBankFeedSyncStatus` |  |
| `xeroUnsynchronisedTransactionsCount` | Indicates the number of unsynchronised transactions with Xero |

Operations: Create, Update.

API path: `/api/v1/accounts/{accountID}/virtual`

#### Webhook

| Field | Description |
| --- | --- |
| `destinationUrl` | The destination URL for the webhook. |
| `emailAddress` | The recipient email address(es) for notifications. |
| `failedNotificationEmailAddress` | The email address to which notifications about failed webhook deliveries will be sent. |
| `id` |  |
| `isActive` |  |
| `merchantID` | The ID of the merchant that the webhook is for. |
| `notificationMethod` | The type of notification that will be sent. |
| `resourceTypes` | The resource types that the webhook will be generated for. |
| `retry` |  |
| `secret` | The secret key required to authenticate webhook notifications. |
| `version` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/webhooks`



## Entities


### Account

Create an instance: `account = client.Account`

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
| `accountBalances` | `Array` | The various balances for the account. |
| `accountID` | `String` | ID of the account. |
| `accountIdentifications` | `Array` | The canoncial identifiers for the account. |
| `accountName` | `String` | Name for the account |
| `accountNames` | `Array` | Optional account names set by the account holder. |
| `accountSupplierName` | `String` | The payment account supplier name. |
| `accountType` | `String` | The type of account e.g. |
| `availableBalance` | `Float` | The current available balance of the account. |
| `availableBalanceMinorUnits` | `Integer` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `Float` | Balance of the account. |
| `balanceMinorUnits` | `Integer` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `String` | The bank name for external accounts |
| `consentID` | `String` | The ID of the consent used to connect the external account. |
| `consolidatedAccountInformation` | `Hash` | Summary information regarding account balances of the overall account provided by the bank. |
| `createdBy` | `Hash` |  |
| `createdByDisplayName` | `String` | Either the name of the user, merchant token or api key that created the account |
| `currency` | `String` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `String` | Indicates the default payment rail for this account. |
| `description` | `String` | Product name as defined by the financial institution for this account. |
| `details` | `String` | Supplementary specifications that might be provided by the Bank. |
| `displayName` | `String` | Gets a unique display name for the payment account. |
| `expiryDate` | `String` | The date that the external account will expire |
| `externalAccountIcon` | `String` | The Icon for external accounts |
| `format` | `String` | File format to save the statement as. |
| `fromDate` | `String` | Minimum transaction date for the statement. |
| `id` | `String` | Unique id for the account. |
| `identifier` | `Hash` |  |
| `inserted` | `String` | Timestamp when the account was created. |
| `isArchived` | `Boolean` | Indicates whether the account is archived. |
| `isConnectedAccount` | `Boolean` | Indicates if the payment account is an externally connected account. |
| `isDefault` | `Boolean` | Is the default account |
| `isTrustAccount` | `Boolean` | Indicates if the payment account is a trust account. |
| `isVirtual` | `Boolean` | True if the account is a virtual account. |
| `lastTransaction` | `Hash` |  |
| `lastUpdated` | `String` | Timestamp when the account was last updated. |
| `merchantID` | `String` | The ID of the merchant that owns the account. |
| `merchantName` | `String` | The name of the merchant that owns the account. |
| `nickname` | `String` | Nickname of the account that was provided by the account owner. |
| `physicalAccountID` | `String` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `roleIDs` | `Array` | Optional list of role IDs that will get access to the payment account when created. |
| `rules` | `Array` | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `Float` | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `Integer` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `String` | Gets a summary of the payments account's most important properties. |
| `supplierPhysicalAccountID` | `String` | For internal use only. |
| `supplierSepaInstantStatus` | `String` | Indicates the status of the SEPA Instant payment rail for this account. |
| `toDate` | `String` | Maximum transaction date for the statement. |
| `type` | `String` | Specifies the type of account e.g. |
| `usageType` | `String` |  |
| `xeroBankFeedConnectionStatus` | `String` | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `String` |  |
| `xeroBankFeedSyncLastFailedAt` | `String` |  |
| `xeroBankFeedSyncLastFailureReason` | `String` |  |
| `xeroBankFeedSyncStatus` | `String` |  |
| `xeroUnsynchronisedTransactionsCount` | `Integer` | Indicates the number of unsynchronised transactions with Xero |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Account record (raises on error).
account = client.Account.load({ "id" => "account_id" })
```

#### Example: List

```ruby
# list returns an Array of Account records (raises on error).
accounts = client.Account.list
```

#### Example: Create

```ruby
account = client.Account.create({
  "account_id" => "example_account_id", # String
  "currency" => "example_currency", # String
  "createdBy" => {}, # Hash
  "identifier" => {}, # Hash
})
```


### Batch

Create an instance: `batch = client.Batch`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approveUrl` | `String` | This field is used when returning a batch payout record to a client. |
| `id` | `String` |  |
| `payouts` | `Array` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Batch record (raises on error).
batch = client.Batch.load({ "id" => "batch_id" })
```

#### Example: Create

```ruby
batch = client.Batch.create({
})
```


### Beneficiary

Create an instance: `beneficiary = client.Beneficiary`

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
| `approvalCallbackUrl` | `String` |  |
| `authenticationMethods` | `Array` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `Array` | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `Integer` | The number of distinct authorisers that have authorised the beneficiary. |
| `authorisersRequiredCount` | `Integer` | The number of authorisers required for this beneficiary. |
| `beneficiaries` | `Array` |  |
| `beneficiaryEvents` | `Array` |  |
| `canAuthorise` | `Boolean` | True if the beneficiary can be authorised by the user who loaded it. |
| `canUpdate` | `Boolean` | True if the beneficiary can be updated by the user who loaded it. |
| `createdBy` | `Hash` |  |
| `createdByEmailAddress` | `String` |  |
| `currency` | `String` | Gets or Sets the currency. |
| `destination` | `Hash` |  |
| `failedBeneficiaries` | `Hash` |  |
| `hasCurrentUserAuthorised` | `Boolean` | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `isEnabled` | `Boolean` |  |
| `lastAuthorised` | `String` |  |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` | Gets or Sets the merchant id. |
| `name` | `String` | The descriptive name for the beneficiary. |
| `nonce` | `String` |  |
| `sourceAccountIDs` | `Array` | ID of the accounts which are authorised to act as a source for the beneficiary. |
| `sourceAccounts` | `Array` |  |
| `theirReference` | `String` | The reference that will be used by default as TheirReference when creating payouts to this beneficiary if no TheirReference is specified for the payout. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Beneficiary record (raises on error).
beneficiary = client.Beneficiary.load({ "id" => "beneficiary_id" })
```

#### Example: List

```ruby
# list returns an Array of Beneficiary records (raises on error).
beneficiarys = client.Beneficiary.list
```

#### Example: Create

```ruby
beneficiary = client.Beneficiary.create({
  "id" => "example_id", # String
  "createdBy" => {}, # Hash
  "currency" => "example_currency", # String
  "name" => "example_name", # String
})
```


### BeneficiaryGroup

Create an instance: `beneficiary_group = client.BeneficiaryGroup`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `groupMembers` | `Array` | The existing group members. |
| `groupName` | `String` | The descriptive name for the beneficiary group. |
| `id` | `String` |  |
| `inserted` | `String` | Timestamp indicating when the group was created. |
| `lastUpdated` | `String` | Timestamp indicating when the group was last updated. |
| `merchantID` | `String` | Gets or Sets the merchant id. |

#### Example: List

```ruby
# list returns an Array of BeneficiaryGroup records (raises on error).
beneficiary_groups = client.BeneficiaryGroup.list
```


### Card

Create an instance: `card = client.Card`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorizedAmount` | `String` |  |
| `currencyCode` | `String` |  |
| `isPayerAuthenticationRequired` | `Boolean` | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `Boolean` | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `String` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `String` | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `String` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `Integer` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `Integer` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `String` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `String` |  |
| `requestID` | `String` |  |
| `responseCode` | `String` |  |
| `responseType` | `String` |  |
| `status` | `String` |  |
| `threeDSRedirectUrl` | `String` | Checkout.com require a redirect for 3DS authentication. |
| `transactionID` | `String` |  |

#### Example: Create

```ruby
card = client.Card.create({
  "paymentrequest_id" => "example_paymentrequest_id", # String
})
```


### CardCustomerToken

Create an instance: `card_customer_token = client.CardCustomerToken`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cardType` | `String` | The type of the tokenised card, e.g. |
| `customerEmailAddress` | `String` | When creating a tokenised card the payer's email address must be supplied. |
| `expiryMonth` | `String` |  |
| `expiryYear` | `String` |  |
| `id` | `String` | The unique ID of the card token that has been stored for the customer. |
| `inserted` | `String` |  |
| `lastFourDigits` | `String` |  |
| `lastUpdated` | `String` |  |
| `maskedCardNumber` | `String` |  |
| `merchantID` | `String` |  |
| `paymentRequestID` | `String` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the CardCustomerToken record (raises on error).
card_customer_token = client.CardCustomerToken.load({ "customer_email_address" => "customer_email_address" })
```

#### Example: List

```ruby
# list returns an Array of CardCustomerToken records (raises on error).
card_customer_tokens = client.CardCustomerToken.list
```


### CardPayment

Create an instance: `card_payment = client.CardPayment`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorizedAmount` | `String` |  |
| `currencyCode` | `String` |  |
| `isPayerAuthenticationRequired` | `Boolean` | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `Boolean` | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `String` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `String` | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `String` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `Integer` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `Integer` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `String` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `String` |  |
| `requestID` | `String` |  |
| `responseCode` | `String` |  |
| `responseType` | `String` |  |
| `status` | `String` |  |
| `threeDSRedirectUrl` | `String` | Checkout.com require a redirect for 3DS authentication. |
| `transactionID` | `String` |  |

#### Example: Create

```ruby
card_payment = client.CardPayment.create({
  "paymentrequest_id" => "example_paymentrequest_id", # String
})
```


### CardPublicKey

Create an instance: `card_public_key = client.CardPublicKey`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `jwt` | `String` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the CardPublicKey record (raises on error).
card_public_key = client.CardPublicKey.load({ "paymentrequest_id" => "paymentrequest_id" })
```


### Consent

Create an instance: `consent = client.Consent`

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
| `authorisationUrl` | `String` | The URL the authorising user needs to be redirected to in order to get the open banking consent token. |
| `callbackUrl` | `String` | Optional callback URL that the end user performing the open banking authorisation will be redirected to on completion. |
| `consentID` | `String` | The ID of the open banking consent. |
| `emailAddress` | `String` | The email address that identifies the end user that will be authorising the open banking consent request. |
| `expiryDate` | `String` |  |
| `failureCallbackUrl` | `String` | Optional callback URL for open banking consent authorisation failure. |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `institutionID` | `String` | The institution ID the open banking consent is being requested for. |
| `isConnectedAccounts` | `Boolean` | Optional setting. |
| `isEnabled` | `Boolean` |  |
| `merchantID` | `String` | The ID of the merchant the consent token is being created to be used with. |
| `provider` | `String` | Lists the supported card and PIS processors. |
| `successWebHookUrl` | `String` | A web hook URL to send an HTTP request to when the open banking consent is successfuly authorised. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Consent record (raises on error).
consent = client.Consent.load({ "id" => "consent_id" })
```

#### Example: List

```ruby
# list returns an Array of Consent records (raises on error).
consents = client.Consent.list
```

#### Example: Create

```ruby
consent = client.Consent.create({
})
```


### Currency

Create an instance: `currency = client.Currency`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `code` | `String` |  |
| `decimals` | `Integer` |  |
| `isFiat` | `Boolean` |  |
| `iso4217AlphaCode` | `String` |  |
| `iso4217NumericCode` | `String` |  |
| `symbol` | `String` |  |

#### Example: List

```ruby
# list returns an Array of Currency records (raises on error).
currencys = client.Currency.list
```


### DirectDebitBatchSubmit

Create an instance: `direct_debit_batch_submit = client.DirectDebitBatchSubmit`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failedSubmissions` | `Hash` | Dictionary of failed submissions, keyed by the index (1-based) in the original request. |
| `successfulSubmissions` | `Array` | List of successfully submitted direct debit payments. |

#### Example: Create

```ruby
direct_debit_batch_submit = client.DirectDebitBatchSubmit.create({
})
```


### FxRate

Create an instance: `fx_rate = client.FxRate`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `destinationCurrency` | `String` |  |
| `exchangeRate` | `Float` | The price at which the transaction will buy the source currency using the destination currency. |
| `expiryTime` | `String` |  |
| `quoteID` | `String` |  |
| `sourceCurrency` | `String` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the FxRate record (raises on error).
fx_rate = client.FxRate.load({ "destination" => "destination", "source" => "source", "valid_for_minute" => 1 })
```

#### Example: List

```ruby
# list returns an Array of FxRate records (raises on error).
fx_rates = client.FxRate.list
```


### IPayment

Create an instance: `i_payment = client.IPayment`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `paymentRequestID` | `String` |  |
| `responseType` | `String` |  |

#### Example: Create

```ruby
i_payment = client.IPayment.create({
})
```


### Mandate

Create an instance: `mandate = client.Mandate`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountNumber` | `String` | Account number of the customer's bank account in case of GBP account. |
| `addressLine1` | `String` | First line of the customer's address. |
| `addressLine2` | `String` | Second line of the customer's address. |
| `approvedAt` | `String` | Date at which the supplier approved this mandate. |
| `city` | `String` | Customer's city. |
| `countryCode` | `String` | 2-character country code of the customer's bank account. |
| `currency` | `String` | Currency of this mandate. |
| `customerAccountNumber` | `String` | Customer's account number in case of GBP account. |
| `customerCity` | `String` | Customer's city of residence. |
| `customerCountryCode` | `String` | Customer's country of residence code. |
| `customerCountryName` | `String` | Customer's country of residence. |
| `customerEmailAddress` | `String` | Customer's email address. |
| `customerFirstName` | `String` | Customer's first name. |
| `customerIban` | `String` | Customer's IBAN in case of EUR account. |
| `customerLastName` | `String` | Customer's last name. |
| `customerSortCode` | `String` | Customer's sort code in case of GBP account. |
| `emailAddress` | `String` | Customer's email address. |
| `firstName` | `String` | Customer's first name. |
| `iban` | `String` | IBAN of the customer's bank account in case of EUR account. |
| `id` | `String` | Internal ID of the mandate. |
| `inserted` | `String` | The timestamp this mandate was created at. |
| `isRecurring` | `Boolean` | Whether this mandate is single-use or recurring. |
| `lastName` | `String` | Customer's last name. |
| `lastUpdated` | `String` | The timestamp this mandate was last updated at. |
| `merchantID` | `String` | Internal ID of this mandate's merchant. |
| `postalCode` | `String` | Customer's postal code. |
| `reference` | `String` | Reference assigned to this mandate. |
| `sortCode` | `String` | Sort code of the customer's bank account in case of GBP account. |
| `status` | `String` | General status of this mandate. |
| `supplierBankAccountID` | `String` | ID that the supplier assigned to this mandate's bank account. |
| `supplierCustomerID` | `String` | ID that the supplier assigned to this mandate's customer. |
| `supplierMandateID` | `String` | ID that the supplier assigned to this mandate. |
| `supplierName` | `String` | Name of the supplier used to create this mandate. |
| `supplierStatus` | `String` | Last status that the supplier reported for this mandate. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Mandate record (raises on error).
mandate = client.Mandate.load({ "id" => "mandate_id" })
```

#### Example: Create

```ruby
mandate = client.Mandate.create({
  "addressLine1" => "example_addressLine1", # String
  "city" => "example_city", # String
  "countryCode" => "example_countryCode", # String
  "emailAddress" => "example_emailAddress", # String
  "firstName" => "example_firstName", # String
  "lastName" => "example_lastName", # String
  "postalCode" => "example_postalCode", # String
})
```


### Merchant

Create an instance: `merchant = client.Merchant`

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
| `accountCurrencies` | `Array` | The list of currencies that the merchant has accounts for. |
| `canHaveTrustAccounts` | `Boolean` | Trust accounts are a special type of account that allow the account name to be trusted for use in statements and verification of payee checks. |
| `cardPaymentProcessor` | `String` | Name of the card payment processor. |
| `companyID` | `String` | The Company ID recorded in the Compliance system. |
| `displayQrOnHostedPay` | `Boolean` | Indicates if a QR Code containing the payment link should be displayed on the hosted payment page. |
| `hostedPayVersion` | `Integer` | The version of the hosted payment page to use with the merchant. |
| `id` | `String` | Unique ID for the merchant. |
| `inserted` | `String` | Timestamp the merchant was added to MoneyMoov. |
| `isBlocked` | `Boolean` | The merchant is blocked from making payments (payouts). |
| `isExited` | `Boolean` | The merchant has formally terminated their relationship and is no longer a customer. |
| `isSuspended` | `Boolean` | The merchant has temporarily suspended their own account. |
| `jurisdiction` | `String` | The jurisdiction the merchant entity is incorporated or established in. |
| `logoUrlPng` | `String` | The CDN URL of the merchant's logo in PNG format. |
| `logoUrlSvg` | `String` | The CDN URL of the merchant's logo in SVG format. |
| `merchantCategoryCode` | `String` | The industry code that represents the merchant's primary trading activity. |
| `name` | `String` | The registered business name of the merchant. |
| `notes` | `String` | The notes field is an optional free text field that can be used to store any additional information about the merchant. |
| `parentMerchant` | `Hash` |  |
| `paymentAccountLimit` | `Integer` | The maximum number of payment accounts that can be created for the Merchant. |
| `paymentAccounts` | `Array` |  |
| `reason` | `String` | The reason for the suspension. |
| `shortName` | `String` | A URL friendly shortish name for the merchant. |
| `supportedPaymentMethodsList` | `Array` | The payment methods that are configured and supported for this merchant. |
| `suspensionReason` | `String` | The reason for the suspension, provided by the merchant. |
| `tags` | `Array` | An optional list of descriptive tags that can be used on merchant entities such as payment requests. |
| `timeZoneId` | `String` | The IANA (Internet Assigned Numbers Authority) time zone identifier of the merchant. |
| `tradingName` | `String` | An optional trading name. |
| `webHookLimit` | `Integer` | The maximum number of web hooks that can be created for the Merchant. |
| `yourRoleName` | `String` | The name of the role for the identity that loaded the merchant record. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Merchant record (raises on error).
merchant = client.Merchant.load({ "id" => "merchant_id" })
```

#### Example: List

```ruby
# list returns an Array of Merchant records (raises on error).
merchants = client.Merchant.list
```


### MerchantAuthorisationSetting

Create an instance: `merchant_authorisation_setting = client.MerchantAuthorisationSetting`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amountLower` | `Float` |  |
| `amountUpper` | `Float` |  |
| `authorisationType` | `String` |  |
| `beneficiariesOnly` | `Boolean` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `lastEditorCantAuthorise` | `Boolean` |  |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` |  |
| `numberOfAuthorisers` | `Integer` |  |
| `roleSettings` | `Array` |  |

#### Example: List

```ruby
# list returns an Array of MerchantAuthorisationSetting records (raises on error).
merchant_authorisation_settings = client.MerchantAuthorisationSetting.list
```


### MerchantDirectDebitMandatePage

Create an instance: `merchant_direct_debit_mandate_page = client.MerchantDirectDebitMandatePage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approvedAt` | `String` | Date at which the supplier approved this mandate. |
| `currency` | `String` | Currency of this mandate. |
| `customerAccountNumber` | `String` | Customer's account number in case of GBP account. |
| `customerCity` | `String` | Customer's city of residence. |
| `customerCountryCode` | `String` | Customer's country of residence code. |
| `customerCountryName` | `String` | Customer's country of residence. |
| `customerEmailAddress` | `String` | Customer's email address. |
| `customerFirstName` | `String` | Customer's first name. |
| `customerIban` | `String` | Customer's IBAN in case of EUR account. |
| `customerLastName` | `String` | Customer's last name. |
| `customerSortCode` | `String` | Customer's sort code in case of GBP account. |
| `id` | `String` | Internal ID of the mandate. |
| `inserted` | `String` | The timestamp this mandate was created at. |
| `isRecurring` | `Boolean` | Whether this mandate is single-use or recurring. |
| `lastUpdated` | `String` | The timestamp this mandate was last updated at. |
| `merchantID` | `String` | Internal ID of this mandate's merchant. |
| `reference` | `String` | Reference assigned to this mandate. |
| `status` | `String` | General status of this mandate. |
| `supplierBankAccountID` | `String` | ID that the supplier assigned to this mandate's bank account. |
| `supplierCustomerID` | `String` | ID that the supplier assigned to this mandate's customer. |
| `supplierMandateID` | `String` | ID that the supplier assigned to this mandate. |
| `supplierName` | `String` | Name of the supplier used to create this mandate. |
| `supplierStatus` | `String` | Last status that the supplier reported for this mandate. |

#### Example: List

```ruby
# list returns an Array of MerchantDirectDebitMandatePage records (raises on error).
merchant_direct_debit_mandate_pages = client.MerchantDirectDebitMandatePage.list
```


### MerchantPayByBankSetting

Create an instance: `merchant_pay_by_bank_setting = client.MerchantPayByBankSetting`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bankCountryCodes` | `Array` | The list of country codes representing the banks the country supports. |
| `bankID` | `String` | ID of the bank to be configured for the merchant. |
| `bankName` | `String` | Name of the Bank/Institution. |
| `businessInstitutionID` | `String` | ID that the processor uses to identify the bank (business accounts). |
| `currency` | `String` | Currency supported by the bank. |
| `logo` | `String` | URL of the bank's logo. |
| `message` | `String` | Message relating to specific bank. |
| `messageImageUrl` | `String` | Optional image URL to be displayed with the message. |
| `order` | `Integer` | Order in which this setting will appear in the UI. |
| `personalInstitutionID` | `String` | ID that the processor uses to identify the bank (personal accounts). |
| `processor` | `String` | Name of the bank payment processor. |
| `warningHeading` | `String` | The heading for a warning message related to the bank institution to be displayed to the user. |
| `warningMessage` | `String` | The warning message related to the bank institution to be displayed to the user. |

#### Example: List

```ruby
# list returns an Array of MerchantPayByBankSetting records (raises on error).
merchant_pay_by_bank_settings = client.MerchantPayByBankSetting.list
```


### MerchantPaymentRequestTemplate

Create an instance: `merchant_payment_request_template = client.MerchantPaymentRequestTemplate`

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
| `bankPaymentOptions` | `Hash` |  |
| `cardPaymentAddressOptions` | `Hash` |  |
| `cardPaymentCaptureOptions` | `Hash` |  |
| `customFields` | `Array` | A list of custom fields that can be included in the payment request template. |
| `defaultFields` | `Array` | A list of default fields that are included in the payment request template. |
| `description` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` |  |
| `name` | `String` |  |
| `notificationOptions` | `Hash` |  |
| `paymentMethods` | `Hash` |  |
| `paymentTerms` | `Hash` |  |
| `priorityBankOptions` | `Hash` |  |
| `template` | `Hash` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the MerchantPaymentRequestTemplate record (raises on error).
merchant_payment_request_template = client.MerchantPaymentRequestTemplate.load({ "id" => "merchant_payment_request_template_id", "paymentrequest_id" => "paymentrequest_id" })
```

#### Example: List

```ruby
# list returns an Array of MerchantPaymentRequestTemplate records (raises on error).
merchant_payment_request_templates = client.MerchantPaymentRequestTemplate.list
```


### MerchantToken

Create an instance: `merchant_token = client.MerchantToken`

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
| `authenticationMethods` | `Array` | A list of authentication types allowed to authorise the merchant token. |
| `authorisations` | `Array` | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `Integer` | The number of distinct authorisers that have authorised the merchant token. |
| `authorisersRequiredCount` | `Integer` | The number of authorisers required for this merchant token. |
| `canAuthorise` | `Boolean` | True if the merchant token can be authorised by the user who loaded it. |
| `description` | `String` | Token description |
| `expiresAt` | `String` | Optional. |
| `hasCurrentUserAuthorised` | `Boolean` | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `hmacAlgorithm` | `String` | Optional shared secret algorithm to use for HMAC authentication. |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `ipAddressWhitelist` | `String` | Optional. |
| `isArchived` | `Boolean` | Indicates whether the merchant token is archived. |
| `isEnabled` | `Boolean` | If set to false the merchant token will not be accepted to authorise a request. |
| `lastAuthorised` | `String` |  |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` | The merchant id to add to the token |
| `nonce` | `String` |  |
| `permissionTypes` | `Array` | The permissions that the merchant token supports. |
| `requestSignatureVersion` | `Integer` | Represent the version of the overall merchant token. |
| `sharedSecretAlgorithm` | `String` | Optional shared secret algorithm to use for HMAC authentication. |
| `sharedSecretBase64` | `String` | The base 64 encoded shared secret that is used for request authentication with an HMAC. |
| `token` | `String` | The JWT merchant token. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the MerchantToken record (raises on error).
merchant_token = client.MerchantToken.load({ "id" => "merchant_token_id" })
```

#### Example: List

```ruby
# list returns an Array of MerchantToken records (raises on error).
merchant_tokens = client.MerchantToken.list
```

#### Example: Create

```ruby
merchant_token = client.MerchantToken.create({
  "nonce" => "example_nonce", # String
})
```


### Metadata

Create an instance: `metadata = client.Metadata`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Metadata record (raises on error).
metadata = client.Metadata.load()
```


### NoFrixionVersion

Create an instance: `no_frixion_version = client.NoFrixionVersion`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `buildVersion` | `Integer` |  |
| `majorVersion` | `Integer` |  |
| `minorVersion` | `Integer` |  |
| `releaseName` | `String` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the NoFrixionVersion record (raises on error).
no_frixion_version = client.NoFrixionVersion.load()
```


### OpenBanking

Create an instance: `open_banking = client.OpenBanking`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```ruby
open_banking = client.OpenBanking.create({
  "account_id" => "example_account_id", # String
})
```


### Payeeverification

Create an instance: `payeeverification = client.Payeeverification`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountName` | `String` | The name of the account to verify |
| `accountNumber` | `String` | The account number of the account to verify (for CoP checks) |
| `iban` | `String` | The IBAN of the account to verify (for VoP checks) |
| `payeeVerifiedAccountName` | `String` | The verified account name of the payee, if available (in case of a close match) |
| `result` | `String` | The result of the payee verification |
| `secondaryIdentification` | `String` | Optional secondary identifier for the account to verify. |
| `sortCode` | `String` | The sort code of the account to verify (for CoP checks) |

#### Example: Create

```ruby
payeeverification = client.Payeeverification.create({
  "accountName" => "example_accountName", # String
  "iban" => "example_iban", # String
})
```


### Payment

Create an instance: `payment = client.Payment`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addresses` | `Array` |  |
| `amount` | `Float` | The amount of money to request. |
| `amountPending` | `Float` | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `Float` | Total amount received for this payment request. |
| `amountRefunded` | `Float` | Total amount refunded for this payment request. |
| `autoSendReceipt` | `Boolean` | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `String` | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `String` | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `Boolean` | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `Boolean` | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `String` | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `Boolean` | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardNoPayerAuthentication` | `Boolean` | If set to true for card payments no attempt will be made to use payer authentication (3-D Secure and equivalent). |
| `cardProcessorMerchantID` | `String` | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `String` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `String` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `cardTransmitRawDetails` | `Boolean` | If set to true for card payments the sensitive card number and card verification number will be transmitted directly rather than being tokenised. |
| `createdByUser` | `Hash` |  |
| `currency` | `String` | The currency of the request. |
| `customFields` | `Array` | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `String` | Optional email address for the customer. |
| `customerID` | `String` | An optional customer identifier for the payment request. |
| `customerName` | `String` |  |
| `description` | `String` | An optional description for the payment request. |
| `destinationAccount` | `Hash` |  |
| `directDebitPayment` | `Hash` | Contains information about a Direct Debit payment attempt for a payment request. |
| `dueDate` | `String` | The due date for the payment request. |
| `events` | `Array` |  |
| `failureCallbackUrl` | `String` | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `Array` | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `String` |  |
| `hostedPayCheckoutUrl` | `String` | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `String` |  |
| `ignoreAddressVerification` | `Boolean` | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `String` | The timestamp the payment request was created at. |
| `insertedSortable` | `String` | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `isArchived` | `Boolean` | Indicates whether the payment request is archived. |
| `jwk` | `String` | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `String` | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `String` | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `String` | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `String` | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `String` | The ID of the merchant to create the payment request for. |
| `merchantTokenDescription` | `String` | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `String` |  |
| `notificationRoleIDs` | `Array` | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `String` | An optional order ID for the payment request. |
| `partialPaymentMethod` | `String` | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `String` | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `Array` | The payment attempts made against this payment request. |
| `paymentMethods` | `Array` | The payment methods that the payment request supports. |
| `paymentProcessor` | `String` | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `payrunID` | `String` | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `String` | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `String` | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `Hash` |  |
| `sandboxSettleDelayInSeconds` | `Integer` | Sandbox only. |
| `shippingAddress` | `Hash` |  |
| `shippingAddressCity` | `String` | Optionally the city of the customer's shipping address. |
| `shippingAddressCountryCode` | `String` | Optionally the country code of the customer's shipping address. |
| `shippingAddressCounty` | `String` | Optionally the state or county of the customer's shipping address. |
| `shippingAddressLine1` | `String` | Optionally the first line of the customer's shipping address. |
| `shippingAddressLine2` | `String` | Optionally the second line of the customer's shipping address. |
| `shippingAddressPostCode` | `String` | Optionally the post code of the customer's shipping address. |
| `shippingEmail` | `String` | Optionally the shipping email address for the customer. |
| `shippingFirstName` | `String` | Optionally the first name of the customer's shipping address. |
| `shippingLastName` | `String` | Optionally the last name of the customer's shipping address. |
| `shippingPhone` | `String` | Optionally the shipping phone number for the customer. |
| `status` | `String` | The current status of the payment request. |
| `successWebHookUrl` | `String` | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tagIds` | `Array` | An optional list of tag ids to add to the payment request |
| `tags` | `Array` | An optional list of descriptive tags attached to the payment request. |
| `title` | `String` | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `Array` |  |
| `transactions` | `Array` |  |
| `useHostedPaymentPage` | `Boolean` | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Payment record (raises on error).
payment = client.Payment.load({ "id" => "payment_id" })
```

#### Example: Create

```ruby
payment = client.Payment.create({
  "createdByUser" => {}, # Hash
})
```


### PaymentAccount

Create an instance: `payment_account = client.PaymentAccount`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountName` | `String` | Name for the account |
| `accountSupplierName` | `String` | The payment account supplier name. |
| `availableBalance` | `Float` | The current available balance of the account. |
| `availableBalanceMinorUnits` | `Integer` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `Float` | Balance of the account. |
| `balanceMinorUnits` | `Integer` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `String` | The bank name for external accounts |
| `consentID` | `String` | The ID of the consent used to connect the external account. |
| `createdBy` | `Hash` |  |
| `createdByDisplayName` | `String` | Either the name of the user, merchant token or api key that created the account |
| `currency` | `String` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `String` | Indicates the default payment rail for this account. |
| `displayName` | `String` | Gets a unique display name for the payment account. |
| `expiryDate` | `String` | The date that the external account will expire |
| `externalAccountIcon` | `String` | The Icon for external accounts |
| `id` | `String` | Unique id for the account. |
| `identifier` | `Hash` |  |
| `inserted` | `String` | Timestamp when the account was created. |
| `isArchived` | `Boolean` | Indicates whether the account is archived. |
| `isConnectedAccount` | `Boolean` | Indicates if the payment account is an externally connected account. |
| `isDefault` | `Boolean` | Is the default account |
| `isTrustAccount` | `Boolean` | Indicates if the payment account is a trust account. |
| `isVirtual` | `Boolean` | True if the account is a virtual account. |
| `lastTransaction` | `Hash` |  |
| `lastUpdated` | `String` | Timestamp when the account was last updated. |
| `merchantID` | `String` | The ID of the merchant that owns the account. |
| `merchantName` | `String` | The name of the merchant that owns the account. |
| `physicalAccountID` | `String` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `Array` | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `Float` | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `Integer` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `String` | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `String` | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `String` | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `String` |  |
| `xeroBankFeedSyncLastFailedAt` | `String` |  |
| `xeroBankFeedSyncLastFailureReason` | `String` |  |
| `xeroBankFeedSyncStatus` | `String` |  |
| `xeroUnsynchronisedTransactionsCount` | `Integer` | Indicates the number of unsynchronised transactions with Xero |

#### Example: List

```ruby
# list returns an Array of PaymentAccount records (raises on error).
payment_accounts = client.PaymentAccount.list
```


### PaymentAccountMinimal

Create an instance: `payment_account_minimal = client.PaymentAccountMinimal`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountName` | `String` | Name for the account |
| `availableBalance` | `Float` | The current available balance of the account. |
| `balance` | `Float` | Balance of the account. |
| `balanceMinorUnits` | `Integer` | Balance of the account expressed in the currency’s minor units (e.g. |
| `currency` | `String` | Currency of the account in ISO 4217 format |
| `id` | `String` | Unique id for the account. |
| `identifier` | `Hash` |  |
| `isArchived` | `Boolean` | Is the account archived |
| `isConnectedAccount` | `Boolean` | Indicates if the payment account is an externally connected account. |
| `merchantID` | `String` | The ID of the merchant that owns the account. |
| `submittedPayoutsBalance` | `Float` | Total of the payouts that have been submitted for processing. |

#### Example: List

```ruby
# list returns an Array of PaymentAccountMinimal records (raises on error).
payment_account_minimals = client.PaymentAccountMinimal.list
```


### PaymentInitiation

Create an instance: `payment_initiation = client.PaymentInitiation`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `paymentInitiationID` | `String` | The unique identifier of the payment initiation request. |
| `paymentRequestCallbackUrl` | `String` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `String` |  |
| `redirectUrl` | `String` | A redirect URL for the user to authorise the payment initiation request at the ASPSP |
| `responseType` | `String` |  |
| `specificErrorMessage` | `String` |  |

#### Example: Create

```ruby
payment_initiation = client.PaymentInitiation.create({
  "paymentrequest_id" => "example_paymentrequest_id", # String
})
```


### PaymentRequest

Create an instance: `payment_request = client.PaymentRequest`

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
| `addresses` | `Array` |  |
| `amount` | `Float` | The amount of money to request. |
| `amountPending` | `Float` | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `Float` | Total amount received for this payment request. |
| `amountRefunded` | `Float` | Total amount refunded for this payment request. |
| `autoSendReceipt` | `Boolean` | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `String` | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `String` | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `Boolean` | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `Boolean` | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `String` | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `Boolean` | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardProcessorMerchantID` | `String` | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `String` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `String` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `createdByUser` | `Hash` |  |
| `currency` | `String` | The currency of the request. |
| `customFields` | `Array` | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `String` | Optional email address for the customer. |
| `customerID` | `String` | An optional customer identifier for the payment request. |
| `customerName` | `String` |  |
| `description` | `String` | An optional description for the payment request. |
| `destinationAccount` | `Hash` |  |
| `directDebitPayment` | `Hash` | Contains information about a Direct Debit payment attempt for a payment request. |
| `doSimulateSettlementFailure` | `Boolean` |  |
| `dueDate` | `String` | The due date for the payment request. |
| `errorDescription` | `String` |  |
| `events` | `Array` |  |
| `failedPaymentRequests` | `Hash` |  |
| `failureCallbackUrl` | `String` | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `Array` | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `String` |  |
| `hostedPayCheckoutUrl` | `String` | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `String` |  |
| `ignoreAddressVerification` | `Boolean` | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `String` | The timestamp the payment request was created at. |
| `insertedSortable` | `String` | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `institution` | `String` |  |
| `isArchived` | `Boolean` | Indicates whether the payment request is archived. |
| `jwk` | `String` | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `String` | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `String` | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `String` | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `String` | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `String` |  |
| `merchantTokenDescription` | `String` | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `String` |  |
| `notificationRoleIDs` | `Array` | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `String` | An optional order ID for the payment request. |
| `partialPaymentMethod` | `String` | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `String` | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `Array` | The payment attempts made against this payment request. |
| `paymentInitiationID` | `String` |  |
| `paymentMethods` | `Array` | The payment methods that the payment request supports. |
| `paymentProcessor` | `String` | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `paymentRequests` | `Array` |  |
| `payrunID` | `String` | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `String` | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `String` | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `Hash` |  |
| `sandboxSettleDelayInSeconds` | `Integer` | Sandbox only. |
| `shippingAddress` | `Hash` |  |
| `status` | `String` | The current status of the payment request. |
| `successWebHookUrl` | `String` | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tags` | `Array` | An optional list of descriptive tags attached to the payment request. |
| `title` | `String` | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `Array` |  |
| `transactions` | `Array` |  |
| `useHostedPaymentPage` | `Boolean` | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the PaymentRequest record (raises on error).
payment_request = client.PaymentRequest.load({ "id" => "payment_request_id" })
```

#### Example: List

```ruby
# list returns an Array of PaymentRequest records (raises on error).
payment_requests = client.PaymentRequest.list
```

#### Example: Create

```ruby
payment_request = client.PaymentRequest.create({
  "createdByUser" => {}, # Hash
})
```


### PaymentRequestEvent

Create an instance: `payment_request_event = client.PaymentRequestEvent`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `Float` |  |
| `applePayTransactionID` | `String` | Transaction ID received in Apple pay token. |
| `cardAuthorizationResponseID` | `String` | For a successful card authorization this field will hold the response ID. |
| `cardExpiryMonth` | `Integer` | For card payment events this field holds the payer's card expiry month. |
| `cardExpiryYear` | `Integer` | For card payment events this field holds the payer's card expiry year. |
| `cardIssuer` | `String` | For card payment events this field holds the payer's card issuer. |
| `cardIssuerCountry` | `String` | For card payment events this field holds the payer's card issuer country of origin. |
| `cardLastFourDigits` | `String` | For card payment events this field holds the payer's card last four digits. |
| `cardRequestID` | `String` |  |
| `cardScheme` | `String` | For card payment events this field holds the scheme of the payer's card, e.g. |
| `cardTokenCustomerID` | `String` | If the option to create a reusable token for card payments was set this field contains the token the merchant can store to use for repeat payments. |
| `cardTransactionID` | `String` |  |
| `currency` | `String` |  |
| `directDebitPaymentID` | `String` | Payment ID issued by the Direct Debit supplier. |
| `directDebitPaymentReference` | `String` | Reference string issued by the Direct Debit supplier. |
| `drirectDebitMandateID` | `String` | The ID of the mandate that was used wehn requesting payment. |
| `errorMessage` | `String` |  |
| `errorReason` | `String` |  |
| `eventType` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `lightningInvoice` | `String` | For Bitcoin Lightning payments this field holds the invoice presented to the payer. |
| `lightningRHash` | `String` | For Bitcoin Lightning payments the hash of the invoice presented to the payer. |
| `originUrl` | `String` | Optional field that can be set by payment methods, such as pay by bank, that may want to redirect back to the URL that initiated the attempt in the case of a failure condition. |
| `paymentMethodType` | `String` | The type of payment method the event relates to, e.g. |
| `paymentProcessorName` | `String` | If the event was for a card payment this is the name of the card processor, e.g. |
| `paymentRequestID` | `String` |  |
| `pispBankStatus` | `String` | For payment initiation attempts some providers (e.g. |
| `pispPaymentInitiationID` | `String` | For a payment initiation this is the ID returned by the service provider initiating the payment for us. |
| `pispPaymentInstitutionName` | `String` | For a payment initiation this is the name of the financial institution that is used to initiate and authorise the payment. |
| `pispPaymentServiceProviderID` | `String` | For a payment initiation this is the service provider ID selected by the payer, typically the ID for the bank or similar financial institution. |
| `pispRedirectUrl` | `String` | For a payment initiation this is the redirect URL returned by the service provider initiating the payment for us. |
| `reconciledTransactionID` | `String` | For settlement events (only relevant for non-card payments) this is the payin transaction that the payment request event was reconciled with. |
| `refundPayoutID` | `String` | ID of the Payout that was created for refund. |
| `status` | `String` |  |
| `walletName` | `String` |  |

#### Example: List

```ruby
# list returns an Array of PaymentRequestEvent records (raises on error).
payment_request_events = client.PaymentRequestEvent.list
```


### PaymentRequestMetric

Create an instance: `payment_request_metric = client.PaymentRequestMetric`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the PaymentRequestMetric record (raises on error).
payment_request_metric = client.PaymentRequestMetric.load()
```


### PaymentRequestMinimal

Create an instance: `payment_request_minimal = client.PaymentRequestMinimal`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `Float` | The amount of money to request. |
| `amountPending` | `Float` | The amount of money that was authorised but has not arrived in the account yet. |
| `amountReceived` | `Float` | The amount of money that has been received for this payment request. |
| `amountRefunded` | `Float` | The amount of money that has been refunded for this payment request. |
| `callbackUrl` | `String` |  |
| `cardStripePaymentIntentSecret` | `String` |  |
| `countryCode` | `String` | The country code associated with the payment. |
| `currency` | `String` | The currency of the request. |
| `customFieldsToDisplay` | `Array` | Custom fields to display to the customer. |
| `description` | `String` | An optional description for the payment request. |
| `dueDate` | `String` | The due date of the payment request. |
| `fieldDisplaySettings` | `Array` |  |
| `googlePayMerchantID` | `String` | Merchant ID from Google Pay |
| `id` | `String` |  |
| `jwk` | `String` | The jwk containing the public key |
| `merchantID` | `String` |  |
| `merchantLogoUrlPng` | `String` |  |
| `merchantLogoUrlSvg` | `String` |  |
| `merchantName` | `String` |  |
| `merchantShortName` | `String` |  |
| `partialPaymentMethod` | `String` |  |
| `paymentAttempts` | `Array` | The payment attempts for this payment request. |
| `paymentMethodsList` | `Array` | The payment methods that the payment request supports. |
| `paymentProcessor` | `String` | The card processor |
| `paymentProcessorKey` | `String` | The card processors public key |
| `pispError` | `String` | This is the error returned from the bank which is recorded in payment request events. |
| `priorityBankID` | `String` |  |
| `status` | `String` | The status of the payment request. |
| `stripeAccountID` | `String` | Account ID of connected customers in Stripe |
| `title` | `String` | The title of the payment request. |

#### Example: List

```ruby
# list returns an Array of PaymentRequestMinimal records (raises on error).
payment_request_minimals = client.PaymentRequestMinimal.list
```


### PaymentRequestResult

Create an instance: `payment_request_result = client.PaymentRequestResult`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `Float` | The authorised payment amount. |
| `amountPending` | `Float` |  |
| `amountReceived` | `Float` |  |
| `amountRefunded` | `Float` |  |
| `currency` | `String` | The authorised payment currency. |
| `customerID` | `String` | The customer id |
| `paymentRequestID` | `String` | The ID of the payment request the result is for. |
| `payments` | `Array` | The list of payment attempts that have been received for the payment request. |
| `pispAuthorizations` | `Array` |  |
| `requestedAmount` | `Float` | The full original payment amount requested. |
| `result` | `String` | The result of the payment attempt. |

#### Example: List

```ruby
# list returns an Array of PaymentRequestResult records (raises on error).
payment_request_results = client.PaymentRequestResult.list
```


### Payout

Create an instance: `payout = client.Payout`

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
| `accountID` | `String` | Gets or Sets Account Id of sending account |
| `allowIncomplete` | `Boolean` | If set to true the payout will get created even if the business validation rules fail. |
| `amount` | `Float` | Gets or Sets payout amount |
| `amountMinorUnits` | `Integer` | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `String` | This field is used when returning an payout record to a client. |
| `approverID` | `String` | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `Array` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `Array` | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `Integer` | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `Integer` | The number of authorisers required for this payout. |
| `batchPayoutID` | `String` | The ID of the batch the payout is associated with. |
| `beneficiary` | `Hash` |  |
| `beneficiaryID` | `String` | Optional. |
| `canAuthorise` | `Boolean` | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `Boolean` | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `Boolean` | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `String` | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `String` |  |
| `createdByEmailAddress` | `String` |  |
| `currency` | `String` | Gets or Sets Currency of payout request |
| `currentUserID` | `String` | The ID of the user that requested access to the PayOut record. |
| `description` | `String` | Gets or Sets description of payout request |
| `destination` | `Hash` |  |
| `documents` | `Array` | Documents associated with the payout. |
| `events` | `Array` | The activity associated with the payout. |
| `failedPayouts` | `Hash` |  |
| `formattedAmount` | `String` | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `String` | FX destination currency and amount formatted string. |
| `formattedSchedule` | `String` |  |
| `formattedScheduleDayOnly` | `String` |  |
| `formattedSourceAccountAvailableBalance` | `String` | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `Float` | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `Integer` | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `String` | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `String` | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `String` | Optional. |
| `fxRate` | `Float` | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `Boolean` | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `Boolean` | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `String` | The ID for the payout. |
| `inserted` | `String` |  |
| `invoiceID` | `String` | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `Boolean` | Indicates whether the payout is archived. |
| `isFailed` | `Boolean` | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `Boolean` | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `Boolean` | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `String` |  |
| `nonce` | `String` |  |
| `paymentProcessor` | `String` | The usptream payment processor for the payout. |
| `paymentRail` | `String` | Optional field to indicate the payment rail to use for the payout. |
| `payouts` | `Array` |  |
| `payrunID` | `String` | The ID of the payrun that this payout is associated with. |
| `payrunName` | `String` | The name of the payrun that this payout is associated with. |
| `reason` | `String` |  |
| `rule` | `Hash` |  |
| `scheduleDate` | `String` | The date the payout should be submitted. |
| `scheduled` | `Boolean` | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `Float` | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `Integer` | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `String` | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `String` | The currency of the source account. |
| `sourceAccountIban` | `String` | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `Hash` |  |
| `sourceAccountName` | `String` | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `String` | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `String` | The sort code of the account the payout is being made from. |
| `status` | `String` | Gets or Sets the status of payout request |
| `tagIds` | `Array` | An optional list of tag ids to add to the payout. |
| `tags` | `Array` | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `String` | Gets or Sets destination reference ID |
| `topupPayrunID` | `String` | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `Float` | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `Float` | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `Float` | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `String` | Gets or Sets payout type |
| `userID` | `String` | Gets or Sets User ID of who created the payout request |
| `yourReference` | `String` | Gets or Sets your reference ID |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Payout record (raises on error).
payout = client.Payout.load({ "id" => "payout_id" })
```

#### Example: List

```ruby
# list returns an Array of Payout records (raises on error).
payouts = client.Payout.list
```

#### Example: Create

```ruby
payout = client.Payout.create({
  "id" => "example_id", # String
  "beneficiary" => {}, # Hash
  "sourceAccountIdentifier" => {}, # Hash
})
```


### PayoutKeysetPage

Create an instance: `payout_keyset_page = client.PayoutKeysetPage`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountID` | `String` | Gets or Sets Account Id of sending account |
| `amount` | `Float` | Gets or Sets payout amount |
| `amountMinorUnits` | `Integer` | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `String` | This field is used when returning an payout record to a client. |
| `approverID` | `String` | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `Array` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `Array` | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `Integer` | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `Integer` | The number of authorisers required for this payout. |
| `batchPayoutID` | `String` | The ID of the batch the payout is associated with. |
| `beneficiary` | `Hash` |  |
| `canAuthorise` | `Boolean` | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `Boolean` | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `Boolean` | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `String` | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `String` |  |
| `createdByEmailAddress` | `String` |  |
| `currency` | `String` | Gets or Sets Currency of payout request |
| `currentUserID` | `String` | The ID of the user that requested access to the PayOut record. |
| `description` | `String` | Gets or Sets description of payout request |
| `destination` | `Hash` |  |
| `documents` | `Array` | Documents associated with the payout. |
| `events` | `Array` | The activity associated with the payout. |
| `formattedAmount` | `String` | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `String` | FX destination currency and amount formatted string. |
| `formattedSchedule` | `String` |  |
| `formattedScheduleDayOnly` | `String` |  |
| `formattedSourceAccountAvailableBalance` | `String` | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `Float` | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `Integer` | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `String` | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `String` | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `String` | Optional. |
| `fxRate` | `Float` | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `Boolean` | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `Boolean` | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `String` | The ID for the payout. |
| `inserted` | `String` |  |
| `invoiceID` | `String` | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `Boolean` | Indicates whether the payout is archived. |
| `isFailed` | `Boolean` | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `Boolean` | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `Boolean` | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `String` |  |
| `nonce` | `String` |  |
| `paymentProcessor` | `String` | The usptream payment processor for the payout. |
| `paymentRail` | `String` | Optional field to indicate the payment rail to use for the payout. |
| `payrunID` | `String` | The ID of the payrun that this payout is associated with. |
| `payrunName` | `String` | The name of the payrun that this payout is associated with. |
| `rule` | `Hash` |  |
| `scheduleDate` | `String` | The date the payout should be submitted. |
| `scheduled` | `Boolean` | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `Float` | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `Integer` | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `String` | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `String` | The currency of the source account. |
| `sourceAccountIban` | `String` | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `Hash` |  |
| `sourceAccountName` | `String` | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `String` | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `String` | The sort code of the account the payout is being made from. |
| `status` | `String` | Gets or Sets the status of payout request |
| `tags` | `Array` | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `String` | Gets or Sets destination reference ID |
| `topupPayrunID` | `String` | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `Float` | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `Float` | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `Float` | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `String` | Gets or Sets payout type |
| `userID` | `String` | Gets or Sets User ID of who created the payout request |
| `yourReference` | `String` | Gets or Sets your reference ID |

#### Example: List

```ruby
# list returns an Array of PayoutKeysetPage records (raises on error).
payout_keyset_pages = client.PayoutKeysetPage.list
```


### PayoutMetric

Create an instance: `payout_metric = client.PayoutMetric`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the PayoutMetric record (raises on error).
payout_metric = client.PayoutMetric.load()
```


### Payrun

Create an instance: `payrun = client.Payrun`

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
| `authorisationDate` | `String` |  |
| `authorisations` | `Array` | A list of the users who have successfully authorised the latest version of the payrun and when. |
| `authorisersCompletedCount` | `Integer` | The number of distinct authorisers that have authorised the payrun. |
| `authorisersRequiredCount` | `Integer` | The number of authorisers required for this payrun. |
| `batchPayoutID` | `String` |  |
| `canAuthorise` | `Boolean` | True if the payrun can be authorised by the user who loaded it. |
| `canDelete` | `Boolean` |  |
| `canEdit` | `Boolean` |  |
| `events` | `Array` |  |
| `hasCurrentUserAuthorised` | `Boolean` | True if the payrun was loaded for a user and that user has already authorised the latest version of the payrun. |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `invoices` | `Array` |  |
| `invoicesMinimal` | `Array` |  |
| `isArchived` | `Boolean` |  |
| `lastUpdated` | `String` |  |
| `lastUpdatedBy` | `Hash` |  |
| `merchantID` | `String` |  |
| `name` | `String` |  |
| `nonce` | `String` |  |
| `notes` | `String` |  |
| `payments` | `Array` |  |
| `payouts` | `Array` |  |
| `payoutsCount` | `Integer` |  |
| `reason` | `String` |  |
| `scheduleDate` | `String` |  |
| `scheduledDate` | `String` |  |
| `sourceAccounts` | `Array` |  |
| `status` | `String` |  |
| `totalEur` | `Float` |  |
| `totalGbp` | `Float` |  |
| `totalUsd` | `Float` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Payrun record (raises on error).
payrun = client.Payrun.load({ "id" => "payrun_id" })
```

#### Example: List

```ruby
# list returns an Array of Payrun records (raises on error).
payruns = client.Payrun.list
```

#### Example: Create

```ruby
payrun = client.Payrun.create({
  "id" => "example_id", # String
  "lastUpdatedBy" => {}, # Hash
})
```


### Report

Create an instance: `report = client.Report`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |


### ReportResult

Create an instance: `report_result = client.ReportResult`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `contentType` | `String` |  |
| `contents` | `String` |  |
| `lastCompletedAt` | `String` |  |
| `merchantID` | `String` |  |
| `reportName` | `String` |  |
| `reportType` | `String` |  |
| `statementNumber` | `Integer` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the ReportResult record (raises on error).
report_result = client.ReportResult.load({ "id" => 1, "report_id" => "report_id" })
```


### Role

Create an instance: `role = client.Role`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failedRoles` | `Hash` |  |
| `roles` | `Array` |  |

#### Example: Create

```ruby
role = client.Role.create({
  "merchant_id" => "example_merchant_id", # String
})
```


### Rule

Create an instance: `rule = client.Rule`

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
| `account` | `Hash` |  |
| `accountID` | `String` | The ID of the account the rule will apply to. |
| `approveUrl` | `String` | If set this property holds the URL an approver needs to visit in order to complete a strong authentication check in order to approve the rule. |
| `approverID` | `String` |  |
| `authenticationMethods` | `Array` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `Array` | A list of the users who have successfully authorised the latest version of the rule and when. |
| `authorisersCompletedCount` | `Integer` | The number of distinct authorisers that have authorised the rule. |
| `authorisersRequiredCount` | `Integer` | The number of authorisers required for this rule. |
| `canAuthorise` | `Boolean` | True if the rule can be authorised by the user who loaded it. |
| `createdBy` | `Hash` |  |
| `description` | `String` | Arbitrary description for the rule. |
| `endAt` | `String` | Optional end time for rule executions. |
| `hasCurrentUserAuthorised` | `Boolean` | True if the current user has authorised. |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `isDisabled` | `Boolean` | If set to true the rule will be disabled from executing. |
| `lastExecutedAt` | `String` |  |
| `lastRunAtTransactionDate` | `String` | The most recent transaction date when the rule was last run. |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` | The ID of the merchant that owns the account. |
| `name` | `String` | A name to succinctly describe the rule. |
| `nonce` | `String` |  |
| `onApprovedWebHookUrl` | `String` | Optional URL to receive an HTTP request with the rule details when the rule status changes to approved. |
| `onExecutionErrorWebHookUrl` | `String` | Optional URL to receive an HTTP request when a rule execution attempt fails. |
| `onExecutionSuccessWebHookUrl` | `String` | Optional URL to receive an HTTP request when a rule execution attempt succeeds. |
| `startAt` | `String` | Optional start time for rule executions. |
| `status` | `String` |  |
| `sweepAction` | `Hash` |  |
| `timeZoneId` | `String` | If the rule should be executed on a recurring schedule this is the timezone that the CRON expression should be evaluated in. |
| `triggerCronExpression` | `String` | If the rule should be executed on a recurring schedule this is the expression that sets the schedule. |
| `triggerOnPayIn` | `Boolean` | Set to true if the rule execution should be triggered when the account receives a pay in (credit). |
| `userID` | `String` |  |
| `webHookSecret` | `String` | If set this secret will be used to sign Web Hook requests. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Rule record (raises on error).
rule = client.Rule.load({ "id" => "rule_id" })
```

#### Example: List

```ruby
# list returns an Array of Rule records (raises on error).
rules = client.Rule.list
```

#### Example: Create

```ruby
rule = client.Rule.create({
  "createdBy" => {}, # Hash
  "nonce" => "example_nonce", # String
})
```


### RuleEvent

Create an instance: `rule_event = client.RuleEvent`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errorMessage` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `isAuthoriseToEnable` | `Boolean` |  |
| `message` | `String` |  |
| `rawResponse` | `String` |  |
| `ruleEventType` | `String` |  |
| `ruleID` | `String` |  |
| `user` | `Hash` |  |

#### Example: List

```ruby
# list returns an Array of RuleEvent records (raises on error).
rule_events = client.RuleEvent.list
```


### Tag

Create an instance: `tag = client.Tag`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `colourHex` | `String` |  |
| `description` | `String` |  |
| `id` | `String` |  |
| `merchantID` | `String` |  |
| `name` | `String` |  |

#### Example: List

```ruby
# list returns an Array of Tag records (raises on error).
tags = client.Tag.list
```

#### Example: Create

```ruby
tag = client.Tag.create({
  "merchant_id" => "example_merchant_id", # String
  "merchantID" => "example_merchantID", # String
  "name" => "example_name", # String
})
```


### Token

Create an instance: `token = client.Token`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```ruby
token = client.Token.create({
  "id" => "example_id", # String
})
```


### Transaction

Create an instance: `transaction = client.Transaction`

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
| `accountID` | `String` | The ID of the account the transaction belongs to. |
| `accountName` | `String` | The name of the account the transaction belongs to. |
| `accountSequenceNumber` | `Integer` | The sequence number of transaction on a per account basis. |
| `addressDetails` | `Hash` |  |
| `amount` | `Float` | Amount of the transaction. |
| `amountMinorUnits` | `Integer` | Amount of the transaction expressed in the currency’s minor units (e.g. |
| `balance` | `Float` | Balance left on the account after the transaction. |
| `balanceMinorUnits` | `Integer` | Balance on the account expressed in the currency’s minor units (e.g. |
| `bookingDateTime` | `String` |  |
| `chargeDetails` | `Hash` |  |
| `content` | `Array` |  |
| `counterparty` | `Hash` |  |
| `counterpartySummary` | `String` | For pay in (credit) transactions this will contain a descriptive string with the most important fields about the counterparty. |
| `currency` | `String` | Currency of transaction. |
| `currencyExchange` | `Hash` | Provides details on the currency exchange. |
| `date` | `String` |  |
| `description` | `String` | Description of the transaction. |
| `enrichment` | `Hash` |  |
| `fxAmount` | `Float` | For an FX payout this is the amound in the FX currency. |
| `fxCurrency` | `String` | For an FX payout this is the currency that was received or that was instructed. |
| `fxRate` | `Float` | For an FX payout this is the exchange rate between the transaction currency and the FX currency. |
| `grossAmount` | `Hash` |  |
| `id` | `String` | Unique ID for the transaction. |
| `inserted` | `String` | Date when the transaction was inserted into the ledger. |
| `isoBankTransactionCode` | `Hash` |  |
| `merchant` | `Hash` |  |
| `merchantID` | `String` | The ID of the merchant that owns the account. |
| `pageNumber` | `Integer` | Current page number. |
| `pageSize` | `Integer` | Page size |
| `payeeDetails` | `Hash` | The Payee object contains details of the beneficiary, person or business. |
| `payerDetails` | `Hash` |  |
| `paymentRequestCustomFields` | `Hash` | The custom fields that were attached to the payment request that resulted in this transaction. |
| `paymentRequestID` | `String` | For Pay by Bank and Direct Debit transactions this will contain the ID of the payment request. |
| `payoutID` | `String` | ID of the payout that resulted in the transaction. |
| `proprietaryBankTransactionCode` | `Hash` |  |
| `rawReference` | `String` | The raw payment reference details as received from the payment processor. |
| `reference` | `String` |  |
| `ruleID` | `String` | ID of the rule that resulted in the transaction. |
| `statementReferences` | `Array` |  |
| `status` | `String` |  |
| `supplementaryData` | `Object` |  |
| `tags` | `Array` | An optional list of descriptive tags attached to the transaction. |
| `theirReference` | `String` | For a pay out the reference that the payer attached for the receiving party. |
| `totalPages` | `Integer` | Total pages |
| `totalSize` | `Integer` | Total count |
| `transactionAmount` | `Hash` |  |
| `transactionDate` | `String` | Date when the transaction occurred. |
| `transactionInformation` | `Array` |  |
| `transactionMutability` | `String` |  |
| `type` | `String` | Type of the transaction. |
| `valueDateTime` | `String` |  |
| `virtualIBAN` | `String` | If set it indicates the payin was to a virtual IBAN. |
| `yourReference` | `String` | For a pay in the reference the sending party attached. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Transaction record (raises on error).
transaction = client.Transaction.load({ "id" => "transaction_id" })
```

#### Example: List

```ruby
# list returns an Array of Transaction records (raises on error).
transactions = client.Transaction.list
```

#### Example: Create

```ruby
transaction = client.Transaction.create({
  "id" => "example_id", # String
  "grossAmount" => {}, # Hash
  "payeeDetails" => {}, # Hash
  "payerDetails" => {}, # Hash
  "transactionAmount" => {}, # Hash
})
```


### User

Create an instance: `user = client.User`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `clientSessionTimeouts` | `Array` | The number of seconds a session for this user should last before expiring. |
| `emailAddress` | `String` |  |
| `firstName` | `String` |  |
| `id` | `String` |  |
| `lastName` | `String` |  |
| `passkeyAdded` | `Boolean` |  |
| `permissions` | `Hash` |  |
| `profile` | `String` |  |
| `rolesWithScope` | `Array` |  |
| `twoFactorEnabled` | `Boolean` |  |
| `userInviteID` | `String` | Optional ID of the invite that is being accepted so the user can be assigned a role on a new merchant. |

#### Example: List

```ruby
# list returns an Array of User records (raises on error).
users = client.User.list
```


### UserInvite

Create an instance: `user_invite = client.UserInvite`

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
| `authorisationStatus` | `Hash` |  |
| `failedUserInvites` | `Hash` |  |
| `id` | `String` |  |
| `initialRoleID` | `String` | The role ID to automatically assign to the merchant’s very first user. |
| `inviteeEmailAddress` | `String` | Email address of the user being invited. |
| `inviteeFirstName` | `String` | First Name of the user being invited. |
| `inviteeLastName` | `String` | Last Name of the user being invited. |
| `inviterEmailAddress` | `String` |  |
| `inviterFirstName` | `String` |  |
| `inviterLastName` | `String` |  |
| `isAuthorised` | `Boolean` | Will be set to true once the invite has met the authorisation requirements. |
| `isInviteeRegistered` | `Boolean` | If true, indicates the invitee's email address corresponds to an existing MoneyMoov user. |
| `lastInvited` | `String` |  |
| `merchantID` | `String` | ID of the merchant the user is being invited to. |
| `merchantName` | `String` |  |
| `message` | `String` |  |
| `registrationUrl` | `String` |  |
| `sendInviteEmail` | `Boolean` | If set to true an email will be sent to the invitee with instructions on how to accept the invite. |
| `status` | `String` |  |
| `user` | `Hash` |  |
| `userID` | `String` |  |
| `userInvites` | `Array` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the UserInvite record (raises on error).
user_invite = client.UserInvite.load({ "id" => "user_invite_id" })
```

#### Example: List

```ruby
# list returns an Array of UserInvite records (raises on error).
user_invites = client.UserInvite.list
```

#### Example: Create

```ruby
user_invite = client.UserInvite.create({
  "id" => "example_id", # String
  "user" => {}, # Hash
})
```


### Virtual

Create an instance: `virtual = client.Virtual`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountName` | `String` | Name for the account |
| `accountSupplierName` | `String` | The payment account supplier name. |
| `availableBalance` | `Float` | The current available balance of the account. |
| `availableBalanceMinorUnits` | `Integer` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `Float` | Balance of the account. |
| `balanceMinorUnits` | `Integer` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `String` | The bank name for external accounts |
| `consentID` | `String` | The ID of the consent used to connect the external account. |
| `createdBy` | `Hash` |  |
| `createdByDisplayName` | `String` | Either the name of the user, merchant token or api key that created the account |
| `currency` | `String` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `String` | Indicates the default payment rail for this account. |
| `displayName` | `String` | Gets a unique display name for the payment account. |
| `expiryDate` | `String` | The date that the external account will expire |
| `externalAccountIcon` | `String` | The Icon for external accounts |
| `id` | `String` | Unique id for the account. |
| `identifier` | `Hash` |  |
| `inserted` | `String` | Timestamp when the account was created. |
| `isArchived` | `Boolean` | Indicates whether the account is archived. |
| `isConnectedAccount` | `Boolean` | Indicates if the payment account is an externally connected account. |
| `isDefault` | `Boolean` | Is the default account |
| `isTrustAccount` | `Boolean` | Indicates if the payment account is a trust account. |
| `isVirtual` | `Boolean` | True if the account is a virtual account. |
| `lastTransaction` | `Hash` |  |
| `lastUpdated` | `String` | Timestamp when the account was last updated. |
| `merchantID` | `String` | The ID of the merchant that owns the account. |
| `merchantName` | `String` | The name of the merchant that owns the account. |
| `name` | `String` | The name of the virtual account. |
| `physicalAccountID` | `String` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `Array` | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `Float` | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `Integer` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `String` | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `String` | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `String` | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `String` |  |
| `xeroBankFeedSyncLastFailedAt` | `String` |  |
| `xeroBankFeedSyncLastFailureReason` | `String` |  |
| `xeroBankFeedSyncStatus` | `String` |  |
| `xeroUnsynchronisedTransactionsCount` | `Integer` | Indicates the number of unsynchronised transactions with Xero |

#### Example: Create

```ruby
virtual = client.Virtual.create({
  "account_id" => "example_account_id", # String
  "createdBy" => {}, # Hash
  "identifier" => {}, # Hash
  "name" => "example_name", # String
})
```


### Webhook

Create an instance: `webhook = client.Webhook`

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
| `destinationUrl` | `String` | The destination URL for the webhook. |
| `emailAddress` | `String` | The recipient email address(es) for notifications. |
| `failedNotificationEmailAddress` | `String` | The email address to which notifications about failed webhook deliveries will be sent. |
| `id` | `String` |  |
| `isActive` | `Boolean` |  |
| `merchantID` | `String` | The ID of the merchant that the webhook is for. |
| `notificationMethod` | `String` | The type of notification that will be sent. |
| `resourceTypes` | `Array` | The resource types that the webhook will be generated for. |
| `retry` | `Boolean` |  |
| `secret` | `String` | The secret key required to authenticate webhook notifications. |
| `version` | `Integer` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Webhook record (raises on error).
webhook = client.Webhook.load({ "id" => "webhook_id" })
```

#### Example: List

```ruby
# list returns an Array of Webhook records (raises on error).
webhooks = client.Webhook.list
```

#### Example: Create

```ruby
webhook = client.Webhook.create({
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

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── Nofrixion_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`Nofrixion_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```ruby
consent = client.Consent
consent.list()

# consent.data_get now returns the consent data from the last list
# consent.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
