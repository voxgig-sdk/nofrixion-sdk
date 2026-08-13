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
created = client.Account.create({ "createdBy" => {}, "identifier" => {} })

# Update — index the record via data_get (created.data_get["id"]).
client.Account.update({ "id" => created.data_get["id"], "account_id" => "example_account_id", "amount" => 1 })

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
| `accountBalances` | `Array` |  |
| `accountID` | `String` |  |
| `accountIdentifications` | `Array` |  |
| `accountName` | `String` |  |
| `accountNames` | `Array` |  |
| `accountSupplierName` | `String` |  |
| `accountType` | `String` |  |
| `availableBalance` | `Float` |  |
| `availableBalanceMinorUnits` | `Integer` |  |
| `balance` | `Float` |  |
| `balanceMinorUnits` | `Integer` |  |
| `bankName` | `String` |  |
| `consentID` | `String` |  |
| `consolidatedAccountInformation` | `Hash` |  |
| `createdBy` | `Hash` |  |
| `createdByDisplayName` | `String` |  |
| `currency` | `String` |  |
| `defaultPaymentRail` | `String` |  |
| `description` | `String` |  |
| `details` | `String` |  |
| `displayName` | `String` |  |
| `expiryDate` | `String` |  |
| `externalAccountIcon` | `String` |  |
| `format` | `String` |  |
| `fromDate` | `String` |  |
| `id` | `String` |  |
| `identifier` | `Hash` |  |
| `inserted` | `String` |  |
| `isArchived` | `Boolean` |  |
| `isConnectedAccount` | `Boolean` |  |
| `isDefault` | `Boolean` |  |
| `isTrustAccount` | `Boolean` |  |
| `isVirtual` | `Boolean` |  |
| `lastTransaction` | `Hash` |  |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` |  |
| `merchantName` | `String` |  |
| `nickname` | `String` |  |
| `physicalAccountID` | `String` |  |
| `roleIDs` | `Array` |  |
| `rules` | `Array` |  |
| `submittedPayoutsBalance` | `Float` |  |
| `submittedPayoutsBalanceMinorUnits` | `Integer` |  |
| `summary` | `String` |  |
| `supplierPhysicalAccountID` | `String` |  |
| `supplierSepaInstantStatus` | `String` |  |
| `toDate` | `String` |  |
| `type` | `String` |  |
| `usageType` | `String` |  |
| `xeroBankFeedConnectionStatus` | `String` |  |
| `xeroBankFeedLastSyncedAt` | `String` |  |
| `xeroBankFeedSyncLastFailedAt` | `String` |  |
| `xeroBankFeedSyncLastFailureReason` | `String` |  |
| `xeroBankFeedSyncStatus` | `String` |  |
| `xeroUnsynchronisedTransactionsCount` | `Integer` |  |

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
| `approveUrl` | `String` |  |
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
| `authenticationMethods` | `Array` |  |
| `authorisations` | `Array` |  |
| `authorisersCompletedCount` | `Integer` |  |
| `authorisersRequiredCount` | `Integer` |  |
| `beneficiaries` | `Array` |  |
| `beneficiaryEvents` | `Array` |  |
| `canAuthorise` | `Boolean` |  |
| `canUpdate` | `Boolean` |  |
| `createdBy` | `Hash` |  |
| `createdByEmailAddress` | `String` |  |
| `currency` | `String` |  |
| `destination` | `Hash` |  |
| `failedBeneficiaries` | `Hash` |  |
| `hasCurrentUserAuthorised` | `Boolean` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `isEnabled` | `Boolean` |  |
| `lastAuthorised` | `String` |  |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` |  |
| `name` | `String` |  |
| `nonce` | `String` |  |
| `sourceAccountIDs` | `Array` |  |
| `sourceAccounts` | `Array` |  |
| `theirReference` | `String` |  |

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
| `groupMembers` | `Array` |  |
| `groupName` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` |  |

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
| `isPayerAuthenticationRequired` | `Boolean` |  |
| `isSoftDecline` | `Boolean` |  |
| `payerAuthenticationAccessToken` | `String` |  |
| `payerAuthenticationMerchantData` | `String` |  |
| `payerAuthenticationUrl` | `String` |  |
| `payerAuthenticationWindowHeight` | `Integer` |  |
| `payerAuthenticationWindowWidth` | `Integer` |  |
| `paymentRequestCallbackUrl` | `String` |  |
| `paymentRequestID` | `String` |  |
| `requestID` | `String` |  |
| `responseCode` | `String` |  |
| `responseType` | `String` |  |
| `status` | `String` |  |
| `threeDSRedirectUrl` | `String` |  |
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
| `cardType` | `String` |  |
| `customerEmailAddress` | `String` |  |
| `expiryMonth` | `String` |  |
| `expiryYear` | `String` |  |
| `id` | `String` |  |
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
| `isPayerAuthenticationRequired` | `Boolean` |  |
| `isSoftDecline` | `Boolean` |  |
| `payerAuthenticationAccessToken` | `String` |  |
| `payerAuthenticationMerchantData` | `String` |  |
| `payerAuthenticationUrl` | `String` |  |
| `payerAuthenticationWindowHeight` | `Integer` |  |
| `payerAuthenticationWindowWidth` | `Integer` |  |
| `paymentRequestCallbackUrl` | `String` |  |
| `paymentRequestID` | `String` |  |
| `requestID` | `String` |  |
| `responseCode` | `String` |  |
| `responseType` | `String` |  |
| `status` | `String` |  |
| `threeDSRedirectUrl` | `String` |  |
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
| `authorisationUrl` | `String` |  |
| `callbackUrl` | `String` |  |
| `consentID` | `String` |  |
| `emailAddress` | `String` |  |
| `expiryDate` | `String` |  |
| `failureCallbackUrl` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `institutionID` | `String` |  |
| `isConnectedAccounts` | `Boolean` |  |
| `isEnabled` | `Boolean` |  |
| `merchantID` | `String` |  |
| `provider` | `String` |  |
| `successWebHookUrl` | `String` |  |

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
| `failedSubmissions` | `Hash` |  |
| `successfulSubmissions` | `Array` |  |

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
| `exchangeRate` | `Float` |  |
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
| `accountNumber` | `String` |  |
| `addressLine1` | `String` |  |
| `addressLine2` | `String` |  |
| `approvedAt` | `String` |  |
| `city` | `String` |  |
| `countryCode` | `String` |  |
| `currency` | `String` |  |
| `customerAccountNumber` | `String` |  |
| `customerCity` | `String` |  |
| `customerCountryCode` | `String` |  |
| `customerCountryName` | `String` |  |
| `customerEmailAddress` | `String` |  |
| `customerFirstName` | `String` |  |
| `customerIban` | `String` |  |
| `customerLastName` | `String` |  |
| `customerSortCode` | `String` |  |
| `emailAddress` | `String` |  |
| `firstName` | `String` |  |
| `iban` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `isRecurring` | `Boolean` |  |
| `lastName` | `String` |  |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` |  |
| `postalCode` | `String` |  |
| `reference` | `String` |  |
| `sortCode` | `String` |  |
| `status` | `String` |  |
| `supplierBankAccountID` | `String` |  |
| `supplierCustomerID` | `String` |  |
| `supplierMandateID` | `String` |  |
| `supplierName` | `String` |  |
| `supplierStatus` | `String` |  |

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
| `accountCurrencies` | `Array` |  |
| `canHaveTrustAccounts` | `Boolean` |  |
| `cardPaymentProcessor` | `String` |  |
| `companyID` | `String` |  |
| `displayQrOnHostedPay` | `Boolean` |  |
| `hostedPayVersion` | `Integer` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `isBlocked` | `Boolean` |  |
| `isExited` | `Boolean` |  |
| `isSuspended` | `Boolean` |  |
| `jurisdiction` | `String` |  |
| `logoUrlPng` | `String` |  |
| `logoUrlSvg` | `String` |  |
| `merchantCategoryCode` | `String` |  |
| `name` | `String` |  |
| `notes` | `String` |  |
| `parentMerchant` | `Hash` |  |
| `paymentAccountLimit` | `Integer` |  |
| `paymentAccounts` | `Array` |  |
| `reason` | `String` |  |
| `shortName` | `String` |  |
| `supportedPaymentMethodsList` | `Array` |  |
| `suspensionReason` | `String` |  |
| `tags` | `Array` |  |
| `timeZoneId` | `String` |  |
| `tradingName` | `String` |  |
| `webHookLimit` | `Integer` |  |
| `yourRoleName` | `String` |  |

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
| `approvedAt` | `String` |  |
| `currency` | `String` |  |
| `customerAccountNumber` | `String` |  |
| `customerCity` | `String` |  |
| `customerCountryCode` | `String` |  |
| `customerCountryName` | `String` |  |
| `customerEmailAddress` | `String` |  |
| `customerFirstName` | `String` |  |
| `customerIban` | `String` |  |
| `customerLastName` | `String` |  |
| `customerSortCode` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `isRecurring` | `Boolean` |  |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` |  |
| `reference` | `String` |  |
| `status` | `String` |  |
| `supplierBankAccountID` | `String` |  |
| `supplierCustomerID` | `String` |  |
| `supplierMandateID` | `String` |  |
| `supplierName` | `String` |  |
| `supplierStatus` | `String` |  |

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
| `bankCountryCodes` | `Array` |  |
| `bankID` | `String` |  |
| `bankName` | `String` |  |
| `businessInstitutionID` | `String` |  |
| `currency` | `String` |  |
| `logo` | `String` |  |
| `message` | `String` |  |
| `messageImageUrl` | `String` |  |
| `order` | `Integer` |  |
| `personalInstitutionID` | `String` |  |
| `processor` | `String` |  |
| `warningHeading` | `String` |  |
| `warningMessage` | `String` |  |

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
| `customFields` | `Array` |  |
| `defaultFields` | `Array` |  |
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
| `authenticationMethods` | `Array` |  |
| `authorisations` | `Array` |  |
| `authorisersCompletedCount` | `Integer` |  |
| `authorisersRequiredCount` | `Integer` |  |
| `canAuthorise` | `Boolean` |  |
| `description` | `String` |  |
| `expiresAt` | `String` |  |
| `hasCurrentUserAuthorised` | `Boolean` |  |
| `hmacAlgorithm` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `ipAddressWhitelist` | `String` |  |
| `isArchived` | `Boolean` |  |
| `isEnabled` | `Boolean` |  |
| `lastAuthorised` | `String` |  |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` |  |
| `nonce` | `String` |  |
| `permissionTypes` | `Array` |  |
| `requestSignatureVersion` | `Integer` |  |
| `sharedSecretAlgorithm` | `String` |  |
| `sharedSecretBase64` | `String` |  |
| `token` | `String` |  |

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
| `accountName` | `String` |  |
| `accountNumber` | `String` |  |
| `iban` | `String` |  |
| `payeeVerifiedAccountName` | `String` |  |
| `result` | `String` |  |
| `secondaryIdentification` | `String` |  |
| `sortCode` | `String` |  |

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
| `amount` | `Float` |  |
| `amountPending` | `Float` |  |
| `amountReceived` | `Float` |  |
| `amountRefunded` | `Float` |  |
| `autoSendReceipt` | `Boolean` |  |
| `baseOriginUrl` | `String` |  |
| `callbackUrl` | `String` |  |
| `cardAuthorizeOnly` | `Boolean` |  |
| `cardCreateToken` | `Boolean` |  |
| `cardCreateTokenMode` | `String` |  |
| `cardIgnoreCVN` | `Boolean` |  |
| `cardNoPayerAuthentication` | `Boolean` |  |
| `cardProcessorMerchantID` | `String` |  |
| `cardStripePaymentIntentID` | `String` |  |
| `cardStripePaymentIntentSecret` | `String` |  |
| `cardTransmitRawDetails` | `Boolean` |  |
| `createdByUser` | `Hash` |  |
| `currency` | `String` |  |
| `customFields` | `Array` |  |
| `customerEmailAddress` | `String` |  |
| `customerID` | `String` |  |
| `customerName` | `String` |  |
| `description` | `String` |  |
| `destinationAccount` | `Hash` |  |
| `directDebitPayment` | `Hash` |  |
| `dueDate` | `String` |  |
| `events` | `Array` |  |
| `failureCallbackUrl` | `String` |  |
| `fieldDisplaySettings` | `Array` |  |
| `formattedAmount` | `String` |  |
| `hostedPayCheckoutUrl` | `String` |  |
| `id` | `String` |  |
| `ignoreAddressVerification` | `Boolean` |  |
| `inserted` | `String` |  |
| `insertedSortable` | `String` |  |
| `isArchived` | `Boolean` |  |
| `jwk` | `String` |  |
| `lastUpdated` | `String` |  |
| `lightningInvoice` | `String` |  |
| `lightningInvoiceExpiresAt` | `String` |  |
| `merchantDirectDebitMandateID` | `String` |  |
| `merchantID` | `String` |  |
| `merchantTokenDescription` | `String` |  |
| `notificationEmailAddresses` | `String` |  |
| `notificationRoleIDs` | `Array` |  |
| `orderID` | `String` |  |
| `partialPaymentMethod` | `String` |  |
| `partialPaymentSteps` | `String` |  |
| `paymentAttempts` | `Array` |  |
| `paymentMethods` | `Array` |  |
| `paymentProcessor` | `String` |  |
| `payrunID` | `String` |  |
| `pispAccountID` | `String` |  |
| `priorityBankID` | `String` |  |
| `result` | `Hash` |  |
| `sandboxSettleDelayInSeconds` | `Integer` |  |
| `shippingAddress` | `Hash` |  |
| `shippingAddressCity` | `String` |  |
| `shippingAddressCountryCode` | `String` |  |
| `shippingAddressCounty` | `String` |  |
| `shippingAddressLine1` | `String` |  |
| `shippingAddressLine2` | `String` |  |
| `shippingAddressPostCode` | `String` |  |
| `shippingEmail` | `String` |  |
| `shippingFirstName` | `String` |  |
| `shippingLastName` | `String` |  |
| `shippingPhone` | `String` |  |
| `status` | `String` |  |
| `successWebHookUrl` | `String` |  |
| `tagIds` | `Array` |  |
| `tags` | `Array` |  |
| `title` | `String` |  |
| `tokenisedCards` | `Array` |  |
| `transactions` | `Array` |  |
| `useHostedPaymentPage` | `Boolean` |  |

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
| `accountName` | `String` |  |
| `accountSupplierName` | `String` |  |
| `availableBalance` | `Float` |  |
| `availableBalanceMinorUnits` | `Integer` |  |
| `balance` | `Float` |  |
| `balanceMinorUnits` | `Integer` |  |
| `bankName` | `String` |  |
| `consentID` | `String` |  |
| `createdBy` | `Hash` |  |
| `createdByDisplayName` | `String` |  |
| `currency` | `String` |  |
| `defaultPaymentRail` | `String` |  |
| `displayName` | `String` |  |
| `expiryDate` | `String` |  |
| `externalAccountIcon` | `String` |  |
| `id` | `String` |  |
| `identifier` | `Hash` |  |
| `inserted` | `String` |  |
| `isArchived` | `Boolean` |  |
| `isConnectedAccount` | `Boolean` |  |
| `isDefault` | `Boolean` |  |
| `isTrustAccount` | `Boolean` |  |
| `isVirtual` | `Boolean` |  |
| `lastTransaction` | `Hash` |  |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` |  |
| `merchantName` | `String` |  |
| `physicalAccountID` | `String` |  |
| `rules` | `Array` |  |
| `submittedPayoutsBalance` | `Float` |  |
| `submittedPayoutsBalanceMinorUnits` | `Integer` |  |
| `summary` | `String` |  |
| `supplierSepaInstantStatus` | `String` |  |
| `xeroBankFeedConnectionStatus` | `String` |  |
| `xeroBankFeedLastSyncedAt` | `String` |  |
| `xeroBankFeedSyncLastFailedAt` | `String` |  |
| `xeroBankFeedSyncLastFailureReason` | `String` |  |
| `xeroBankFeedSyncStatus` | `String` |  |
| `xeroUnsynchronisedTransactionsCount` | `Integer` |  |

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
| `accountName` | `String` |  |
| `availableBalance` | `Float` |  |
| `balance` | `Float` |  |
| `balanceMinorUnits` | `Integer` |  |
| `currency` | `String` |  |
| `id` | `String` |  |
| `identifier` | `Hash` |  |
| `isArchived` | `Boolean` |  |
| `isConnectedAccount` | `Boolean` |  |
| `merchantID` | `String` |  |
| `submittedPayoutsBalance` | `Float` |  |

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
| `paymentInitiationID` | `String` |  |
| `paymentRequestCallbackUrl` | `String` |  |
| `paymentRequestID` | `String` |  |
| `redirectUrl` | `String` |  |
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
| `amount` | `Float` |  |
| `amountPending` | `Float` |  |
| `amountReceived` | `Float` |  |
| `amountRefunded` | `Float` |  |
| `autoSendReceipt` | `Boolean` |  |
| `baseOriginUrl` | `String` |  |
| `callbackUrl` | `String` |  |
| `cardAuthorizeOnly` | `Boolean` |  |
| `cardCreateToken` | `Boolean` |  |
| `cardCreateTokenMode` | `String` |  |
| `cardIgnoreCVN` | `Boolean` |  |
| `cardProcessorMerchantID` | `String` |  |
| `cardStripePaymentIntentID` | `String` |  |
| `cardStripePaymentIntentSecret` | `String` |  |
| `createdByUser` | `Hash` |  |
| `currency` | `String` |  |
| `customFields` | `Array` |  |
| `customerEmailAddress` | `String` |  |
| `customerID` | `String` |  |
| `customerName` | `String` |  |
| `description` | `String` |  |
| `destinationAccount` | `Hash` |  |
| `directDebitPayment` | `Hash` |  |
| `doSimulateSettlementFailure` | `Boolean` |  |
| `dueDate` | `String` |  |
| `errorDescription` | `String` |  |
| `events` | `Array` |  |
| `failedPaymentRequests` | `Hash` |  |
| `failureCallbackUrl` | `String` |  |
| `fieldDisplaySettings` | `Array` |  |
| `formattedAmount` | `String` |  |
| `hostedPayCheckoutUrl` | `String` |  |
| `id` | `String` |  |
| `ignoreAddressVerification` | `Boolean` |  |
| `inserted` | `String` |  |
| `insertedSortable` | `String` |  |
| `institution` | `String` |  |
| `isArchived` | `Boolean` |  |
| `jwk` | `String` |  |
| `lastUpdated` | `String` |  |
| `lightningInvoice` | `String` |  |
| `lightningInvoiceExpiresAt` | `String` |  |
| `merchantDirectDebitMandateID` | `String` |  |
| `merchantID` | `String` |  |
| `merchantTokenDescription` | `String` |  |
| `notificationEmailAddresses` | `String` |  |
| `notificationRoleIDs` | `Array` |  |
| `orderID` | `String` |  |
| `partialPaymentMethod` | `String` |  |
| `partialPaymentSteps` | `String` |  |
| `paymentAttempts` | `Array` |  |
| `paymentInitiationID` | `String` |  |
| `paymentMethods` | `Array` |  |
| `paymentProcessor` | `String` |  |
| `paymentRequests` | `Array` |  |
| `payrunID` | `String` |  |
| `pispAccountID` | `String` |  |
| `priorityBankID` | `String` |  |
| `result` | `Hash` |  |
| `sandboxSettleDelayInSeconds` | `Integer` |  |
| `shippingAddress` | `Hash` |  |
| `status` | `String` |  |
| `successWebHookUrl` | `String` |  |
| `tags` | `Array` |  |
| `title` | `String` |  |
| `tokenisedCards` | `Array` |  |
| `transactions` | `Array` |  |
| `useHostedPaymentPage` | `Boolean` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the PaymentRequest record (raises on error).
payment_request = client.PaymentRequest.load()
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
| `applePayTransactionID` | `String` |  |
| `cardAuthorizationResponseID` | `String` |  |
| `cardExpiryMonth` | `Integer` |  |
| `cardExpiryYear` | `Integer` |  |
| `cardIssuer` | `String` |  |
| `cardIssuerCountry` | `String` |  |
| `cardLastFourDigits` | `String` |  |
| `cardRequestID` | `String` |  |
| `cardScheme` | `String` |  |
| `cardTokenCustomerID` | `String` |  |
| `cardTransactionID` | `String` |  |
| `currency` | `String` |  |
| `directDebitPaymentID` | `String` |  |
| `directDebitPaymentReference` | `String` |  |
| `drirectDebitMandateID` | `String` |  |
| `errorMessage` | `String` |  |
| `errorReason` | `String` |  |
| `eventType` | `String` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `lightningInvoice` | `String` |  |
| `lightningRHash` | `String` |  |
| `originUrl` | `String` |  |
| `paymentMethodType` | `String` |  |
| `paymentProcessorName` | `String` |  |
| `paymentRequestID` | `String` |  |
| `pispBankStatus` | `String` |  |
| `pispPaymentInitiationID` | `String` |  |
| `pispPaymentInstitutionName` | `String` |  |
| `pispPaymentServiceProviderID` | `String` |  |
| `pispRedirectUrl` | `String` |  |
| `reconciledTransactionID` | `String` |  |
| `refundPayoutID` | `String` |  |
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
| `amount` | `Float` |  |
| `amountPending` | `Float` |  |
| `amountReceived` | `Float` |  |
| `amountRefunded` | `Float` |  |
| `callbackUrl` | `String` |  |
| `cardStripePaymentIntentSecret` | `String` |  |
| `countryCode` | `String` |  |
| `currency` | `String` |  |
| `customFieldsToDisplay` | `Array` |  |
| `description` | `String` |  |
| `dueDate` | `String` |  |
| `fieldDisplaySettings` | `Array` |  |
| `googlePayMerchantID` | `String` |  |
| `id` | `String` |  |
| `jwk` | `String` |  |
| `merchantID` | `String` |  |
| `merchantLogoUrlPng` | `String` |  |
| `merchantLogoUrlSvg` | `String` |  |
| `merchantName` | `String` |  |
| `merchantShortName` | `String` |  |
| `partialPaymentMethod` | `String` |  |
| `paymentAttempts` | `Array` |  |
| `paymentMethodsList` | `Array` |  |
| `paymentProcessor` | `String` |  |
| `paymentProcessorKey` | `String` |  |
| `pispError` | `String` |  |
| `priorityBankID` | `String` |  |
| `status` | `String` |  |
| `stripeAccountID` | `String` |  |
| `title` | `String` |  |

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
| `amount` | `Float` |  |
| `amountPending` | `Float` |  |
| `amountReceived` | `Float` |  |
| `amountRefunded` | `Float` |  |
| `currency` | `String` |  |
| `customerID` | `String` |  |
| `paymentRequestID` | `String` |  |
| `payments` | `Array` |  |
| `pispAuthorizations` | `Array` |  |
| `requestedAmount` | `Float` |  |
| `result` | `String` |  |

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
| `accountID` | `String` |  |
| `allowIncomplete` | `Boolean` |  |
| `amount` | `Float` |  |
| `amountMinorUnits` | `Integer` |  |
| `approvePayoutUrl` | `String` |  |
| `approverID` | `String` |  |
| `authenticationMethods` | `Array` |  |
| `authorisations` | `Array` |  |
| `authorisersCompletedCount` | `Integer` |  |
| `authorisersRequiredCount` | `Integer` |  |
| `batchPayoutID` | `String` |  |
| `beneficiary` | `Hash` |  |
| `beneficiaryID` | `String` |  |
| `canAuthorise` | `Boolean` |  |
| `canProcess` | `Boolean` |  |
| `canUpdate` | `Boolean` |  |
| `chargeBearer` | `String` |  |
| `createdBy` | `String` |  |
| `createdByEmailAddress` | `String` |  |
| `currency` | `String` |  |
| `currentUserID` | `String` |  |
| `description` | `String` |  |
| `destination` | `Hash` |  |
| `documents` | `Array` |  |
| `events` | `Array` |  |
| `failedPayouts` | `Hash` |  |
| `formattedAmount` | `String` |  |
| `formattedFxDestinationAmount` | `String` |  |
| `formattedSchedule` | `String` |  |
| `formattedScheduleDayOnly` | `String` |  |
| `formattedSourceAccountAvailableBalance` | `String` |  |
| `fxDestinationAmount` | `Float` |  |
| `fxDestinationAmountMinorUnits` | `Integer` |  |
| `fxDestinationCurrency` | `String` |  |
| `fxQuoteExpiresAt` | `String` |  |
| `fxQuoteID` | `String` |  |
| `fxRate` | `Float` |  |
| `fxUseDestinationAmount` | `Boolean` |  |
| `hasCurrentUserAuthorised` | `Boolean` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `invoiceID` | `String` |  |
| `isArchived` | `Boolean` |  |
| `isFailed` | `Boolean` |  |
| `isSettled` | `Boolean` |  |
| `isSubmitted` | `Boolean` |  |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` |  |
| `merchantTokenDescription` | `String` |  |
| `nonce` | `String` |  |
| `paymentProcessor` | `String` |  |
| `paymentRail` | `String` |  |
| `payouts` | `Array` |  |
| `payrunID` | `String` |  |
| `payrunName` | `String` |  |
| `reason` | `String` |  |
| `rule` | `Hash` |  |
| `scheduleDate` | `String` |  |
| `scheduled` | `Boolean` |  |
| `sourceAccountAvailableBalance` | `Float` |  |
| `sourceAccountAvailableBalanceMinorUnits` | `Integer` |  |
| `sourceAccountBic` | `String` |  |
| `sourceAccountCurrency` | `String` |  |
| `sourceAccountIban` | `String` |  |
| `sourceAccountIdentifier` | `Hash` |  |
| `sourceAccountName` | `String` |  |
| `sourceAccountNumber` | `String` |  |
| `sourceAccountSortcode` | `String` |  |
| `status` | `String` |  |
| `tagIds` | `Array` |  |
| `tags` | `Array` |  |
| `theirReference` | `String` |  |
| `topupPayrunID` | `String` |  |
| `transactedAmount` | `Float` |  |
| `transactedFxAmount` | `Float` |  |
| `transactedFxRate` | `Float` |  |
| `type` | `String` |  |
| `userID` | `String` |  |
| `yourReference` | `String` |  |

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
| `accountID` | `String` |  |
| `amount` | `Float` |  |
| `amountMinorUnits` | `Integer` |  |
| `approvePayoutUrl` | `String` |  |
| `approverID` | `String` |  |
| `authenticationMethods` | `Array` |  |
| `authorisations` | `Array` |  |
| `authorisersCompletedCount` | `Integer` |  |
| `authorisersRequiredCount` | `Integer` |  |
| `batchPayoutID` | `String` |  |
| `beneficiary` | `Hash` |  |
| `canAuthorise` | `Boolean` |  |
| `canProcess` | `Boolean` |  |
| `canUpdate` | `Boolean` |  |
| `chargeBearer` | `String` |  |
| `createdBy` | `String` |  |
| `createdByEmailAddress` | `String` |  |
| `currency` | `String` |  |
| `currentUserID` | `String` |  |
| `description` | `String` |  |
| `destination` | `Hash` |  |
| `documents` | `Array` |  |
| `events` | `Array` |  |
| `formattedAmount` | `String` |  |
| `formattedFxDestinationAmount` | `String` |  |
| `formattedSchedule` | `String` |  |
| `formattedScheduleDayOnly` | `String` |  |
| `formattedSourceAccountAvailableBalance` | `String` |  |
| `fxDestinationAmount` | `Float` |  |
| `fxDestinationAmountMinorUnits` | `Integer` |  |
| `fxDestinationCurrency` | `String` |  |
| `fxQuoteExpiresAt` | `String` |  |
| `fxQuoteID` | `String` |  |
| `fxRate` | `Float` |  |
| `fxUseDestinationAmount` | `Boolean` |  |
| `hasCurrentUserAuthorised` | `Boolean` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `invoiceID` | `String` |  |
| `isArchived` | `Boolean` |  |
| `isFailed` | `Boolean` |  |
| `isSettled` | `Boolean` |  |
| `isSubmitted` | `Boolean` |  |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` |  |
| `merchantTokenDescription` | `String` |  |
| `nonce` | `String` |  |
| `paymentProcessor` | `String` |  |
| `paymentRail` | `String` |  |
| `payrunID` | `String` |  |
| `payrunName` | `String` |  |
| `rule` | `Hash` |  |
| `scheduleDate` | `String` |  |
| `scheduled` | `Boolean` |  |
| `sourceAccountAvailableBalance` | `Float` |  |
| `sourceAccountAvailableBalanceMinorUnits` | `Integer` |  |
| `sourceAccountBic` | `String` |  |
| `sourceAccountCurrency` | `String` |  |
| `sourceAccountIban` | `String` |  |
| `sourceAccountIdentifier` | `Hash` |  |
| `sourceAccountName` | `String` |  |
| `sourceAccountNumber` | `String` |  |
| `sourceAccountSortcode` | `String` |  |
| `status` | `String` |  |
| `tags` | `Array` |  |
| `theirReference` | `String` |  |
| `topupPayrunID` | `String` |  |
| `transactedAmount` | `Float` |  |
| `transactedFxAmount` | `Float` |  |
| `transactedFxRate` | `Float` |  |
| `type` | `String` |  |
| `userID` | `String` |  |
| `yourReference` | `String` |  |

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
| `authorisations` | `Array` |  |
| `authorisersCompletedCount` | `Integer` |  |
| `authorisersRequiredCount` | `Integer` |  |
| `batchPayoutID` | `String` |  |
| `canAuthorise` | `Boolean` |  |
| `canDelete` | `Boolean` |  |
| `canEdit` | `Boolean` |  |
| `events` | `Array` |  |
| `hasCurrentUserAuthorised` | `Boolean` |  |
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
| `accountID` | `String` |  |
| `approveUrl` | `String` |  |
| `approverID` | `String` |  |
| `authenticationMethods` | `Array` |  |
| `authorisations` | `Array` |  |
| `authorisersCompletedCount` | `Integer` |  |
| `authorisersRequiredCount` | `Integer` |  |
| `canAuthorise` | `Boolean` |  |
| `createdBy` | `Hash` |  |
| `description` | `String` |  |
| `endAt` | `String` |  |
| `hasCurrentUserAuthorised` | `Boolean` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `isDisabled` | `Boolean` |  |
| `lastExecutedAt` | `String` |  |
| `lastRunAtTransactionDate` | `String` |  |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` |  |
| `name` | `String` |  |
| `nonce` | `String` |  |
| `onApprovedWebHookUrl` | `String` |  |
| `onExecutionErrorWebHookUrl` | `String` |  |
| `onExecutionSuccessWebHookUrl` | `String` |  |
| `startAt` | `String` |  |
| `status` | `String` |  |
| `sweepAction` | `Hash` |  |
| `timeZoneId` | `String` |  |
| `triggerCronExpression` | `String` |  |
| `triggerOnPayIn` | `Boolean` |  |
| `userID` | `String` |  |
| `webHookSecret` | `String` |  |

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
| `accountID` | `String` |  |
| `accountName` | `String` |  |
| `accountSequenceNumber` | `Integer` |  |
| `addressDetails` | `Hash` |  |
| `amount` | `Float` |  |
| `amountMinorUnits` | `Integer` |  |
| `balance` | `Float` |  |
| `balanceMinorUnits` | `Integer` |  |
| `bookingDateTime` | `String` |  |
| `chargeDetails` | `Hash` |  |
| `content` | `Array` |  |
| `counterparty` | `Hash` |  |
| `counterpartySummary` | `String` |  |
| `currency` | `String` |  |
| `currencyExchange` | `Hash` |  |
| `date` | `String` |  |
| `description` | `String` |  |
| `enrichment` | `Hash` |  |
| `fxAmount` | `Float` |  |
| `fxCurrency` | `String` |  |
| `fxRate` | `Float` |  |
| `grossAmount` | `Hash` |  |
| `id` | `String` |  |
| `inserted` | `String` |  |
| `isoBankTransactionCode` | `Hash` |  |
| `merchant` | `Hash` |  |
| `merchantID` | `String` |  |
| `pageNumber` | `Integer` |  |
| `pageSize` | `Integer` |  |
| `payeeDetails` | `Hash` |  |
| `payerDetails` | `Hash` |  |
| `paymentRequestCustomFields` | `Hash` |  |
| `paymentRequestID` | `String` |  |
| `payoutID` | `String` |  |
| `proprietaryBankTransactionCode` | `Hash` |  |
| `rawReference` | `String` |  |
| `reference` | `String` |  |
| `ruleID` | `String` |  |
| `statementReferences` | `Array` |  |
| `status` | `String` |  |
| `supplementaryData` | `Object` |  |
| `tags` | `Array` |  |
| `theirReference` | `String` |  |
| `totalPages` | `Integer` |  |
| `totalSize` | `Integer` |  |
| `transactionAmount` | `Hash` |  |
| `transactionDate` | `String` |  |
| `transactionInformation` | `Array` |  |
| `transactionMutability` | `String` |  |
| `type` | `String` |  |
| `valueDateTime` | `String` |  |
| `virtualIBAN` | `String` |  |
| `yourReference` | `String` |  |

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
| `clientSessionTimeouts` | `Array` |  |
| `emailAddress` | `String` |  |
| `firstName` | `String` |  |
| `id` | `String` |  |
| `lastName` | `String` |  |
| `passkeyAdded` | `Boolean` |  |
| `permissions` | `Hash` |  |
| `profile` | `String` |  |
| `rolesWithScope` | `Array` |  |
| `twoFactorEnabled` | `Boolean` |  |
| `userInviteID` | `String` |  |

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
| `initialRoleID` | `String` |  |
| `inviteeEmailAddress` | `String` |  |
| `inviteeFirstName` | `String` |  |
| `inviteeLastName` | `String` |  |
| `inviterEmailAddress` | `String` |  |
| `inviterFirstName` | `String` |  |
| `inviterLastName` | `String` |  |
| `isAuthorised` | `Boolean` |  |
| `isInviteeRegistered` | `Boolean` |  |
| `lastInvited` | `String` |  |
| `merchantID` | `String` |  |
| `merchantName` | `String` |  |
| `message` | `String` |  |
| `registrationUrl` | `String` |  |
| `sendInviteEmail` | `Boolean` |  |
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
| `accountName` | `String` |  |
| `accountSupplierName` | `String` |  |
| `availableBalance` | `Float` |  |
| `availableBalanceMinorUnits` | `Integer` |  |
| `balance` | `Float` |  |
| `balanceMinorUnits` | `Integer` |  |
| `bankName` | `String` |  |
| `consentID` | `String` |  |
| `createdBy` | `Hash` |  |
| `createdByDisplayName` | `String` |  |
| `currency` | `String` |  |
| `defaultPaymentRail` | `String` |  |
| `displayName` | `String` |  |
| `expiryDate` | `String` |  |
| `externalAccountIcon` | `String` |  |
| `id` | `String` |  |
| `identifier` | `Hash` |  |
| `inserted` | `String` |  |
| `isArchived` | `Boolean` |  |
| `isConnectedAccount` | `Boolean` |  |
| `isDefault` | `Boolean` |  |
| `isTrustAccount` | `Boolean` |  |
| `isVirtual` | `Boolean` |  |
| `lastTransaction` | `Hash` |  |
| `lastUpdated` | `String` |  |
| `merchantID` | `String` |  |
| `merchantName` | `String` |  |
| `name` | `String` |  |
| `physicalAccountID` | `String` |  |
| `rules` | `Array` |  |
| `submittedPayoutsBalance` | `Float` |  |
| `submittedPayoutsBalanceMinorUnits` | `Integer` |  |
| `summary` | `String` |  |
| `supplierSepaInstantStatus` | `String` |  |
| `xeroBankFeedConnectionStatus` | `String` |  |
| `xeroBankFeedLastSyncedAt` | `String` |  |
| `xeroBankFeedSyncLastFailedAt` | `String` |  |
| `xeroBankFeedSyncLastFailureReason` | `String` |  |
| `xeroBankFeedSyncStatus` | `String` |  |
| `xeroUnsynchronisedTransactionsCount` | `Integer` |  |

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
| `destinationUrl` | `String` |  |
| `emailAddress` | `String` |  |
| `failedNotificationEmailAddress` | `String` |  |
| `id` | `String` |  |
| `isActive` | `Boolean` |  |
| `merchantID` | `String` |  |
| `notificationMethod` | `String` |  |
| `resourceTypes` | `Array` |  |
| `retry` | `Boolean` |  |
| `secret` | `String` |  |
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
