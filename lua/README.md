# Nofrixion Lua SDK



The Lua SDK for the Nofrixion API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:Account()` — each with the same small set of operations (`list`, `load`, `create`, `update`, `remove`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/nofrixion-sdk/releases)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("nofrixion_sdk")

local client = sdk.new({
  apikey = os.getenv("NOFRIXION_APIKEY"),
})
```

### 2. List account records

Entity operations return `(value, err)`. For `list`, `value` is the
array of records itself — iterate it directly (there is no wrapper).

```lua
local accounts, err = client:Account():list()
if err then error(err) end

for _, item in ipairs(accounts) do
  print(item["id"], item["accountID"])
end
```

### 3. Load a cardcustomertoken

CardCustomerToken is nested under customer_email_address, so provide the `customer_email_address`.

```lua
local cardcustomertoken, err = client:CardCustomerToken():load({ customer_email_address = "example_customer_email_address" })
if err then error(err) end
print(cardcustomertoken)
```

### 4. Create, update, and remove

```lua
-- Create
local created, err = client:Account():create({ createdBy = {}, identifier = {} })
if err then error(err) end

-- Update
client:Account():update({ id = created:data_get()["id"], account_id = "example_account_id", amount = 1 })

-- Remove
client:Account():remove({ id = created:data_get()["id"] })
```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local consents, err = client:Consent():list()
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:Consent():list()
-- result is the returned data; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
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
cd lua && busted test/
```


## Reference

### NofrixionSDK

```lua
local sdk = require("nofrixion_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### NofrixionSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
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
| `load` | `(reqmatch, ctrl) -> any, err` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> any, err` | List entities matching the criteria. |
| `create` | `(reqdata, ctrl) -> any, err` | Create a new entity. |
| `update` | `(reqdata, ctrl) -> any, err` | Update an existing entity. |
| `remove` | `(reqmatch, ctrl) -> any, err` | Remove an entity. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the operation's
data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `load` / `create` / `update` / `remove` | the entity record (a `table`) |
| `list` | an array (`table`) of entity records |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local account, err = client:Account():load({ id = "example_id" })
    if err then error(err) end
    -- account is the loaded record

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

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

Create an instance: `local account = client:Account(nil)`

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
| `accountBalances` | `table` |  |
| `accountID` | `string` |  |
| `accountIdentifications` | `table` |  |
| `accountName` | `string` |  |
| `accountNames` | `table` |  |
| `accountSupplierName` | `string` |  |
| `accountType` | `string` |  |
| `availableBalance` | `number` |  |
| `availableBalanceMinorUnits` | `number` |  |
| `balance` | `number` |  |
| `balanceMinorUnits` | `number` |  |
| `bankName` | `string` |  |
| `consentID` | `string` |  |
| `consolidatedAccountInformation` | `table` |  |
| `createdBy` | `table` |  |
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
| `identifier` | `table` |  |
| `inserted` | `string` |  |
| `isArchived` | `boolean` |  |
| `isConnectedAccount` | `boolean` |  |
| `isDefault` | `boolean` |  |
| `isTrustAccount` | `boolean` |  |
| `isVirtual` | `boolean` |  |
| `lastTransaction` | `table` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantName` | `string` |  |
| `nickname` | `string` |  |
| `physicalAccountID` | `string` |  |
| `roleIDs` | `table` |  |
| `rules` | `table` |  |
| `submittedPayoutsBalance` | `number` |  |
| `submittedPayoutsBalanceMinorUnits` | `number` |  |
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
| `xeroUnsynchronisedTransactionsCount` | `number` |  |

#### Example: Load

```lua
local account, err = client:Account():load({ id = "account_id" })
```

#### Example: List

```lua
local accounts, err = client:Account():list()
```

#### Example: Create

```lua
local account, err = client:Account():create({
  createdBy = {}, -- table
  identifier = {}, -- table
})
```


### Batch

Create an instance: `local batch = client:Batch(nil)`

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
| `payouts` | `table` |  |

#### Example: Load

```lua
local batch, err = client:Batch():load({ id = "batch_id" })
```

#### Example: Create

```lua
local batch, err = client:Batch():create({
})
```


### Beneficiary

Create an instance: `local beneficiary = client:Beneficiary(nil)`

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
| `authenticationMethods` | `table` |  |
| `authorisations` | `table` |  |
| `authorisersCompletedCount` | `number` |  |
| `authorisersRequiredCount` | `number` |  |
| `beneficiaries` | `table` |  |
| `beneficiaryEvents` | `table` |  |
| `canAuthorise` | `boolean` |  |
| `canUpdate` | `boolean` |  |
| `createdBy` | `table` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` |  |
| `destination` | `table` |  |
| `failedBeneficiaries` | `table` |  |
| `hasCurrentUserAuthorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isEnabled` | `boolean` |  |
| `lastAuthorised` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `sourceAccountIDs` | `table` |  |
| `sourceAccounts` | `table` |  |
| `theirReference` | `string` |  |

#### Example: Load

```lua
local beneficiary, err = client:Beneficiary():load({ id = "beneficiary_id" })
```

#### Example: List

```lua
local beneficiarys, err = client:Beneficiary():list()
```

#### Example: Create

```lua
local beneficiary, err = client:Beneficiary():create({
  createdBy = {}, -- table
  currency = "example_currency", -- string
  name = "example_name", -- string
})
```


### BeneficiaryGroup

Create an instance: `local beneficiary_group = client:BeneficiaryGroup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `groupMembers` | `table` |  |
| `groupName` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |

#### Example: List

```lua
local beneficiary_groups, err = client:BeneficiaryGroup():list()
```


### Card

Create an instance: `local card = client:Card(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorizedAmount` | `string` |  |
| `currencyCode` | `string` |  |
| `isPayerAuthenticationRequired` | `boolean` |  |
| `isSoftDecline` | `boolean` |  |
| `payerAuthenticationAccessToken` | `string` |  |
| `payerAuthenticationMerchantData` | `string` |  |
| `payerAuthenticationUrl` | `string` |  |
| `payerAuthenticationWindowHeight` | `number` |  |
| `payerAuthenticationWindowWidth` | `number` |  |
| `paymentRequestCallbackUrl` | `string` |  |
| `paymentRequestID` | `string` |  |
| `requestID` | `string` |  |
| `responseCode` | `string` |  |
| `responseType` | `string` |  |
| `status` | `string` |  |
| `threeDSRedirectUrl` | `string` |  |
| `transactionID` | `string` |  |

#### Example: Create

```lua
local card, err = client:Card():create({
  paymentrequest_id = "example_paymentrequest_id", -- string
})
```


### CardCustomerToken

Create an instance: `local card_customer_token = client:CardCustomerToken(nil)`

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

```lua
local card_customer_token, err = client:CardCustomerToken():load({ customer_email_address = "customer_email_address" })
```

#### Example: List

```lua
local card_customer_tokens, err = client:CardCustomerToken():list()
```


### CardPayment

Create an instance: `local card_payment = client:CardPayment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorizedAmount` | `string` |  |
| `currencyCode` | `string` |  |
| `isPayerAuthenticationRequired` | `boolean` |  |
| `isSoftDecline` | `boolean` |  |
| `payerAuthenticationAccessToken` | `string` |  |
| `payerAuthenticationMerchantData` | `string` |  |
| `payerAuthenticationUrl` | `string` |  |
| `payerAuthenticationWindowHeight` | `number` |  |
| `payerAuthenticationWindowWidth` | `number` |  |
| `paymentRequestCallbackUrl` | `string` |  |
| `paymentRequestID` | `string` |  |
| `requestID` | `string` |  |
| `responseCode` | `string` |  |
| `responseType` | `string` |  |
| `status` | `string` |  |
| `threeDSRedirectUrl` | `string` |  |
| `transactionID` | `string` |  |

#### Example: Create

```lua
local card_payment, err = client:CardPayment():create({
  paymentrequest_id = "example_paymentrequest_id", -- string
})
```


### CardPublicKey

Create an instance: `local card_public_key = client:CardPublicKey(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `jwt` | `string` |  |

#### Example: Load

```lua
local card_public_key, err = client:CardPublicKey():load({ paymentrequest_id = "paymentrequest_id" })
```


### Consent

Create an instance: `local consent = client:Consent(nil)`

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
| `isConnectedAccounts` | `boolean` |  |
| `isEnabled` | `boolean` |  |
| `merchantID` | `string` |  |
| `provider` | `string` |  |
| `successWebHookUrl` | `string` |  |

#### Example: Load

```lua
local consent, err = client:Consent():load({ id = "consent_id" })
```

#### Example: List

```lua
local consents, err = client:Consent():list()
```

#### Example: Create

```lua
local consent, err = client:Consent():create({
})
```


### Currency

Create an instance: `local currency = client:Currency(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `code` | `string` |  |
| `decimals` | `number` |  |
| `isFiat` | `boolean` |  |
| `iso4217AlphaCode` | `string` |  |
| `iso4217NumericCode` | `string` |  |
| `symbol` | `string` |  |

#### Example: List

```lua
local currencys, err = client:Currency():list()
```


### DirectDebitBatchSubmit

Create an instance: `local direct_debit_batch_submit = client:DirectDebitBatchSubmit(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failedSubmissions` | `table` |  |
| `successfulSubmissions` | `table` |  |

#### Example: Create

```lua
local direct_debit_batch_submit, err = client:DirectDebitBatchSubmit():create({
})
```


### FxRate

Create an instance: `local fx_rate = client:FxRate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `destinationCurrency` | `string` |  |
| `exchangeRate` | `number` |  |
| `expiryTime` | `string` |  |
| `quoteID` | `string` |  |
| `sourceCurrency` | `string` |  |

#### Example: Load

```lua
local fx_rate, err = client:FxRate():load({ destination = "destination", source = "source", valid_for_minute = 1 })
```

#### Example: List

```lua
local fx_rates, err = client:FxRate():list()
```


### IPayment

Create an instance: `local i_payment = client:IPayment(nil)`

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

```lua
local i_payment, err = client:IPayment():create({
})
```


### Mandate

Create an instance: `local mandate = client:Mandate(nil)`

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
| `isRecurring` | `boolean` |  |
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

```lua
local mandate, err = client:Mandate():load({ id = "mandate_id" })
```

#### Example: Create

```lua
local mandate, err = client:Mandate():create({
  addressLine1 = "example_addressLine1", -- string
  city = "example_city", -- string
  countryCode = "example_countryCode", -- string
  emailAddress = "example_emailAddress", -- string
  firstName = "example_firstName", -- string
  lastName = "example_lastName", -- string
  postalCode = "example_postalCode", -- string
})
```


### Merchant

Create an instance: `local merchant = client:Merchant(nil)`

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
| `accountCurrencies` | `table` |  |
| `canHaveTrustAccounts` | `boolean` |  |
| `cardPaymentProcessor` | `string` |  |
| `companyID` | `string` |  |
| `displayQrOnHostedPay` | `boolean` |  |
| `hostedPayVersion` | `number` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isBlocked` | `boolean` |  |
| `isExited` | `boolean` |  |
| `isSuspended` | `boolean` |  |
| `jurisdiction` | `string` |  |
| `logoUrlPng` | `string` |  |
| `logoUrlSvg` | `string` |  |
| `merchantCategoryCode` | `string` |  |
| `name` | `string` |  |
| `notes` | `string` |  |
| `parentMerchant` | `table` |  |
| `paymentAccountLimit` | `number` |  |
| `paymentAccounts` | `table` |  |
| `reason` | `string` |  |
| `shortName` | `string` |  |
| `supportedPaymentMethodsList` | `table` |  |
| `suspensionReason` | `string` |  |
| `tags` | `table` |  |
| `timeZoneId` | `string` |  |
| `tradingName` | `string` |  |
| `webHookLimit` | `number` |  |
| `yourRoleName` | `string` |  |

#### Example: Load

```lua
local merchant, err = client:Merchant():load({ id = "merchant_id" })
```

#### Example: List

```lua
local merchants, err = client:Merchant():list()
```


### MerchantAuthorisationSetting

Create an instance: `local merchant_authorisation_setting = client:MerchantAuthorisationSetting(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amountLower` | `number` |  |
| `amountUpper` | `number` |  |
| `authorisationType` | `string` |  |
| `beneficiariesOnly` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lastEditorCantAuthorise` | `boolean` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `numberOfAuthorisers` | `number` |  |
| `roleSettings` | `table` |  |

#### Example: List

```lua
local merchant_authorisation_settings, err = client:MerchantAuthorisationSetting():list()
```


### MerchantDirectDebitMandatePage

Create an instance: `local merchant_direct_debit_mandate_page = client:MerchantDirectDebitMandatePage(nil)`

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
| `isRecurring` | `boolean` |  |
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

```lua
local merchant_direct_debit_mandate_pages, err = client:MerchantDirectDebitMandatePage():list()
```


### MerchantPayByBankSetting

Create an instance: `local merchant_pay_by_bank_setting = client:MerchantPayByBankSetting(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bankCountryCodes` | `table` |  |
| `bankID` | `string` |  |
| `bankName` | `string` |  |
| `businessInstitutionID` | `string` |  |
| `currency` | `string` |  |
| `logo` | `string` |  |
| `message` | `string` |  |
| `messageImageUrl` | `string` |  |
| `order` | `number` |  |
| `personalInstitutionID` | `string` |  |
| `processor` | `string` |  |
| `warningHeading` | `string` |  |
| `warningMessage` | `string` |  |

#### Example: List

```lua
local merchant_pay_by_bank_settings, err = client:MerchantPayByBankSetting():list()
```


### MerchantPaymentRequestTemplate

Create an instance: `local merchant_payment_request_template = client:MerchantPaymentRequestTemplate(nil)`

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
| `bankPaymentOptions` | `table` |  |
| `cardPaymentAddressOptions` | `table` |  |
| `cardPaymentCaptureOptions` | `table` |  |
| `customFields` | `table` |  |
| `defaultFields` | `table` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `name` | `string` |  |
| `notificationOptions` | `table` |  |
| `paymentMethods` | `table` |  |
| `paymentTerms` | `table` |  |
| `priorityBankOptions` | `table` |  |
| `template` | `table` |  |

#### Example: Load

```lua
local merchant_payment_request_template, err = client:MerchantPaymentRequestTemplate():load({ id = "merchant_payment_request_template_id", paymentrequest_id = "paymentrequest_id" })
```

#### Example: List

```lua
local merchant_payment_request_templates, err = client:MerchantPaymentRequestTemplate():list()
```


### MerchantToken

Create an instance: `local merchant_token = client:MerchantToken(nil)`

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
| `authenticationMethods` | `table` |  |
| `authorisations` | `table` |  |
| `authorisersCompletedCount` | `number` |  |
| `authorisersRequiredCount` | `number` |  |
| `canAuthorise` | `boolean` |  |
| `description` | `string` |  |
| `expiresAt` | `string` |  |
| `hasCurrentUserAuthorised` | `boolean` |  |
| `hmacAlgorithm` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `ipAddressWhitelist` | `string` |  |
| `isArchived` | `boolean` |  |
| `isEnabled` | `boolean` |  |
| `lastAuthorised` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `nonce` | `string` |  |
| `permissionTypes` | `table` |  |
| `requestSignatureVersion` | `number` |  |
| `sharedSecretAlgorithm` | `string` |  |
| `sharedSecretBase64` | `string` |  |
| `token` | `string` |  |

#### Example: Load

```lua
local merchant_token, err = client:MerchantToken():load({ id = "merchant_token_id" })
```

#### Example: List

```lua
local merchant_tokens, err = client:MerchantToken():list()
```

#### Example: Create

```lua
local merchant_token, err = client:MerchantToken():create({
  nonce = "example_nonce", -- string
})
```


### Metadata

Create an instance: `local metadata = client:Metadata(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```lua
local metadata, err = client:Metadata():load()
```


### NoFrixionVersion

Create an instance: `local no_frixion_version = client:NoFrixionVersion(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `buildVersion` | `number` |  |
| `majorVersion` | `number` |  |
| `minorVersion` | `number` |  |
| `releaseName` | `string` |  |

#### Example: Load

```lua
local no_frixion_version, err = client:NoFrixionVersion():load()
```


### OpenBanking

Create an instance: `local open_banking = client:OpenBanking(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```lua
local open_banking, err = client:OpenBanking():create({
  account_id = "example_account_id", -- string
})
```


### Payeeverification

Create an instance: `local payeeverification = client:Payeeverification(nil)`

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

```lua
local payeeverification, err = client:Payeeverification():create({
  accountName = "example_accountName", -- string
  iban = "example_iban", -- string
})
```


### Payment

Create an instance: `local payment = client:Payment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addresses` | `table` |  |
| `amount` | `number` |  |
| `amountPending` | `number` |  |
| `amountReceived` | `number` |  |
| `amountRefunded` | `number` |  |
| `autoSendReceipt` | `boolean` |  |
| `baseOriginUrl` | `string` |  |
| `callbackUrl` | `string` |  |
| `cardAuthorizeOnly` | `boolean` |  |
| `cardCreateToken` | `boolean` |  |
| `cardCreateTokenMode` | `string` |  |
| `cardIgnoreCVN` | `boolean` |  |
| `cardNoPayerAuthentication` | `boolean` |  |
| `cardProcessorMerchantID` | `string` |  |
| `cardStripePaymentIntentID` | `string` |  |
| `cardStripePaymentIntentSecret` | `string` |  |
| `cardTransmitRawDetails` | `boolean` |  |
| `createdByUser` | `table` |  |
| `currency` | `string` |  |
| `customFields` | `table` |  |
| `customerEmailAddress` | `string` |  |
| `customerID` | `string` |  |
| `customerName` | `string` |  |
| `description` | `string` |  |
| `destinationAccount` | `table` |  |
| `directDebitPayment` | `table` |  |
| `dueDate` | `string` |  |
| `events` | `table` |  |
| `failureCallbackUrl` | `string` |  |
| `fieldDisplaySettings` | `table` |  |
| `formattedAmount` | `string` |  |
| `hostedPayCheckoutUrl` | `string` |  |
| `id` | `string` |  |
| `ignoreAddressVerification` | `boolean` |  |
| `inserted` | `string` |  |
| `insertedSortable` | `string` |  |
| `isArchived` | `boolean` |  |
| `jwk` | `string` |  |
| `lastUpdated` | `string` |  |
| `lightningInvoice` | `string` |  |
| `lightningInvoiceExpiresAt` | `string` |  |
| `merchantDirectDebitMandateID` | `string` |  |
| `merchantID` | `string` |  |
| `merchantTokenDescription` | `string` |  |
| `notificationEmailAddresses` | `string` |  |
| `notificationRoleIDs` | `table` |  |
| `orderID` | `string` |  |
| `partialPaymentMethod` | `string` |  |
| `partialPaymentSteps` | `string` |  |
| `paymentAttempts` | `table` |  |
| `paymentMethods` | `table` |  |
| `paymentProcessor` | `string` |  |
| `payrunID` | `string` |  |
| `pispAccountID` | `string` |  |
| `priorityBankID` | `string` |  |
| `result` | `table` |  |
| `sandboxSettleDelayInSeconds` | `number` |  |
| `shippingAddress` | `table` |  |
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
| `tagIds` | `table` |  |
| `tags` | `table` |  |
| `title` | `string` |  |
| `tokenisedCards` | `table` |  |
| `transactions` | `table` |  |
| `useHostedPaymentPage` | `boolean` |  |

#### Example: Load

```lua
local payment, err = client:Payment():load({ id = "payment_id" })
```

#### Example: Create

```lua
local payment, err = client:Payment():create({
  createdByUser = {}, -- table
})
```


### PaymentAccount

Create an instance: `local payment_account = client:PaymentAccount(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountName` | `string` |  |
| `accountSupplierName` | `string` |  |
| `availableBalance` | `number` |  |
| `availableBalanceMinorUnits` | `number` |  |
| `balance` | `number` |  |
| `balanceMinorUnits` | `number` |  |
| `bankName` | `string` |  |
| `consentID` | `string` |  |
| `createdBy` | `table` |  |
| `createdByDisplayName` | `string` |  |
| `currency` | `string` |  |
| `defaultPaymentRail` | `string` |  |
| `displayName` | `string` |  |
| `expiryDate` | `string` |  |
| `externalAccountIcon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `table` |  |
| `inserted` | `string` |  |
| `isArchived` | `boolean` |  |
| `isConnectedAccount` | `boolean` |  |
| `isDefault` | `boolean` |  |
| `isTrustAccount` | `boolean` |  |
| `isVirtual` | `boolean` |  |
| `lastTransaction` | `table` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantName` | `string` |  |
| `physicalAccountID` | `string` |  |
| `rules` | `table` |  |
| `submittedPayoutsBalance` | `number` |  |
| `submittedPayoutsBalanceMinorUnits` | `number` |  |
| `summary` | `string` |  |
| `supplierSepaInstantStatus` | `string` |  |
| `xeroBankFeedConnectionStatus` | `string` |  |
| `xeroBankFeedLastSyncedAt` | `string` |  |
| `xeroBankFeedSyncLastFailedAt` | `string` |  |
| `xeroBankFeedSyncLastFailureReason` | `string` |  |
| `xeroBankFeedSyncStatus` | `string` |  |
| `xeroUnsynchronisedTransactionsCount` | `number` |  |

#### Example: List

```lua
local payment_accounts, err = client:PaymentAccount():list()
```


### PaymentAccountMinimal

Create an instance: `local payment_account_minimal = client:PaymentAccountMinimal(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountName` | `string` |  |
| `availableBalance` | `number` |  |
| `balance` | `number` |  |
| `balanceMinorUnits` | `number` |  |
| `currency` | `string` |  |
| `id` | `string` |  |
| `identifier` | `table` |  |
| `isArchived` | `boolean` |  |
| `isConnectedAccount` | `boolean` |  |
| `merchantID` | `string` |  |
| `submittedPayoutsBalance` | `number` |  |

#### Example: List

```lua
local payment_account_minimals, err = client:PaymentAccountMinimal():list()
```


### PaymentInitiation

Create an instance: `local payment_initiation = client:PaymentInitiation(nil)`

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

```lua
local payment_initiation, err = client:PaymentInitiation():create({
  paymentrequest_id = "example_paymentrequest_id", -- string
})
```


### PaymentRequest

Create an instance: `local payment_request = client:PaymentRequest(nil)`

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
| `addresses` | `table` |  |
| `amount` | `number` |  |
| `amountPending` | `number` |  |
| `amountReceived` | `number` |  |
| `amountRefunded` | `number` |  |
| `autoSendReceipt` | `boolean` |  |
| `baseOriginUrl` | `string` |  |
| `callbackUrl` | `string` |  |
| `cardAuthorizeOnly` | `boolean` |  |
| `cardCreateToken` | `boolean` |  |
| `cardCreateTokenMode` | `string` |  |
| `cardIgnoreCVN` | `boolean` |  |
| `cardProcessorMerchantID` | `string` |  |
| `cardStripePaymentIntentID` | `string` |  |
| `cardStripePaymentIntentSecret` | `string` |  |
| `createdByUser` | `table` |  |
| `currency` | `string` |  |
| `customFields` | `table` |  |
| `customerEmailAddress` | `string` |  |
| `customerID` | `string` |  |
| `customerName` | `string` |  |
| `description` | `string` |  |
| `destinationAccount` | `table` |  |
| `directDebitPayment` | `table` |  |
| `doSimulateSettlementFailure` | `boolean` |  |
| `dueDate` | `string` |  |
| `errorDescription` | `string` |  |
| `events` | `table` |  |
| `failedPaymentRequests` | `table` |  |
| `failureCallbackUrl` | `string` |  |
| `fieldDisplaySettings` | `table` |  |
| `formattedAmount` | `string` |  |
| `hostedPayCheckoutUrl` | `string` |  |
| `id` | `string` |  |
| `ignoreAddressVerification` | `boolean` |  |
| `inserted` | `string` |  |
| `insertedSortable` | `string` |  |
| `institution` | `string` |  |
| `isArchived` | `boolean` |  |
| `jwk` | `string` |  |
| `lastUpdated` | `string` |  |
| `lightningInvoice` | `string` |  |
| `lightningInvoiceExpiresAt` | `string` |  |
| `merchantDirectDebitMandateID` | `string` |  |
| `merchantID` | `string` |  |
| `merchantTokenDescription` | `string` |  |
| `notificationEmailAddresses` | `string` |  |
| `notificationRoleIDs` | `table` |  |
| `orderID` | `string` |  |
| `partialPaymentMethod` | `string` |  |
| `partialPaymentSteps` | `string` |  |
| `paymentAttempts` | `table` |  |
| `paymentInitiationID` | `string` |  |
| `paymentMethods` | `table` |  |
| `paymentProcessor` | `string` |  |
| `paymentRequests` | `table` |  |
| `payrunID` | `string` |  |
| `pispAccountID` | `string` |  |
| `priorityBankID` | `string` |  |
| `result` | `table` |  |
| `sandboxSettleDelayInSeconds` | `number` |  |
| `shippingAddress` | `table` |  |
| `status` | `string` |  |
| `successWebHookUrl` | `string` |  |
| `tags` | `table` |  |
| `title` | `string` |  |
| `tokenisedCards` | `table` |  |
| `transactions` | `table` |  |
| `useHostedPaymentPage` | `boolean` |  |

#### Example: Load

```lua
local payment_request, err = client:PaymentRequest():load()
```

#### Example: List

```lua
local payment_requests, err = client:PaymentRequest():list()
```

#### Example: Create

```lua
local payment_request, err = client:PaymentRequest():create({
  createdByUser = {}, -- table
})
```


### PaymentRequestEvent

Create an instance: `local payment_request_event = client:PaymentRequestEvent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `number` |  |
| `applePayTransactionID` | `string` |  |
| `cardAuthorizationResponseID` | `string` |  |
| `cardExpiryMonth` | `number` |  |
| `cardExpiryYear` | `number` |  |
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

```lua
local payment_request_events, err = client:PaymentRequestEvent():list()
```


### PaymentRequestMetric

Create an instance: `local payment_request_metric = client:PaymentRequestMetric(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```lua
local payment_request_metric, err = client:PaymentRequestMetric():load()
```


### PaymentRequestMinimal

Create an instance: `local payment_request_minimal = client:PaymentRequestMinimal(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `number` |  |
| `amountPending` | `number` |  |
| `amountReceived` | `number` |  |
| `amountRefunded` | `number` |  |
| `callbackUrl` | `string` |  |
| `cardStripePaymentIntentSecret` | `string` |  |
| `countryCode` | `string` |  |
| `currency` | `string` |  |
| `customFieldsToDisplay` | `table` |  |
| `description` | `string` |  |
| `dueDate` | `string` |  |
| `fieldDisplaySettings` | `table` |  |
| `googlePayMerchantID` | `string` |  |
| `id` | `string` |  |
| `jwk` | `string` |  |
| `merchantID` | `string` |  |
| `merchantLogoUrlPng` | `string` |  |
| `merchantLogoUrlSvg` | `string` |  |
| `merchantName` | `string` |  |
| `merchantShortName` | `string` |  |
| `partialPaymentMethod` | `string` |  |
| `paymentAttempts` | `table` |  |
| `paymentMethodsList` | `table` |  |
| `paymentProcessor` | `string` |  |
| `paymentProcessorKey` | `string` |  |
| `pispError` | `string` |  |
| `priorityBankID` | `string` |  |
| `status` | `string` |  |
| `stripeAccountID` | `string` |  |
| `title` | `string` |  |

#### Example: List

```lua
local payment_request_minimals, err = client:PaymentRequestMinimal():list()
```


### PaymentRequestResult

Create an instance: `local payment_request_result = client:PaymentRequestResult(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `number` |  |
| `amountPending` | `number` |  |
| `amountReceived` | `number` |  |
| `amountRefunded` | `number` |  |
| `currency` | `string` |  |
| `customerID` | `string` |  |
| `paymentRequestID` | `string` |  |
| `payments` | `table` |  |
| `pispAuthorizations` | `table` |  |
| `requestedAmount` | `number` |  |
| `result` | `string` |  |

#### Example: List

```lua
local payment_request_results, err = client:PaymentRequestResult():list()
```


### Payout

Create an instance: `local payout = client:Payout(nil)`

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
| `allowIncomplete` | `boolean` |  |
| `amount` | `number` |  |
| `amountMinorUnits` | `number` |  |
| `approvePayoutUrl` | `string` |  |
| `approverID` | `string` |  |
| `authenticationMethods` | `table` |  |
| `authorisations` | `table` |  |
| `authorisersCompletedCount` | `number` |  |
| `authorisersRequiredCount` | `number` |  |
| `batchPayoutID` | `string` |  |
| `beneficiary` | `table` |  |
| `beneficiaryID` | `string` |  |
| `canAuthorise` | `boolean` |  |
| `canProcess` | `boolean` |  |
| `canUpdate` | `boolean` |  |
| `chargeBearer` | `string` |  |
| `createdBy` | `string` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` |  |
| `currentUserID` | `string` |  |
| `description` | `string` |  |
| `destination` | `table` |  |
| `documents` | `table` |  |
| `events` | `table` |  |
| `failedPayouts` | `table` |  |
| `formattedAmount` | `string` |  |
| `formattedFxDestinationAmount` | `string` |  |
| `formattedSchedule` | `string` |  |
| `formattedScheduleDayOnly` | `string` |  |
| `formattedSourceAccountAvailableBalance` | `string` |  |
| `fxDestinationAmount` | `number` |  |
| `fxDestinationAmountMinorUnits` | `number` |  |
| `fxDestinationCurrency` | `string` |  |
| `fxQuoteExpiresAt` | `string` |  |
| `fxQuoteID` | `string` |  |
| `fxRate` | `number` |  |
| `fxUseDestinationAmount` | `boolean` |  |
| `hasCurrentUserAuthorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoiceID` | `string` |  |
| `isArchived` | `boolean` |  |
| `isFailed` | `boolean` |  |
| `isSettled` | `boolean` |  |
| `isSubmitted` | `boolean` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantTokenDescription` | `string` |  |
| `nonce` | `string` |  |
| `paymentProcessor` | `string` |  |
| `paymentRail` | `string` |  |
| `payouts` | `table` |  |
| `payrunID` | `string` |  |
| `payrunName` | `string` |  |
| `reason` | `string` |  |
| `rule` | `table` |  |
| `scheduleDate` | `string` |  |
| `scheduled` | `boolean` |  |
| `sourceAccountAvailableBalance` | `number` |  |
| `sourceAccountAvailableBalanceMinorUnits` | `number` |  |
| `sourceAccountBic` | `string` |  |
| `sourceAccountCurrency` | `string` |  |
| `sourceAccountIban` | `string` |  |
| `sourceAccountIdentifier` | `table` |  |
| `sourceAccountName` | `string` |  |
| `sourceAccountNumber` | `string` |  |
| `sourceAccountSortcode` | `string` |  |
| `status` | `string` |  |
| `tagIds` | `table` |  |
| `tags` | `table` |  |
| `theirReference` | `string` |  |
| `topupPayrunID` | `string` |  |
| `transactedAmount` | `number` |  |
| `transactedFxAmount` | `number` |  |
| `transactedFxRate` | `number` |  |
| `type` | `string` |  |
| `userID` | `string` |  |
| `yourReference` | `string` |  |

#### Example: Load

```lua
local payout, err = client:Payout():load({ id = "payout_id" })
```

#### Example: List

```lua
local payouts, err = client:Payout():list()
```

#### Example: Create

```lua
local payout, err = client:Payout():create({
  beneficiary = {}, -- table
  sourceAccountIdentifier = {}, -- table
})
```


### PayoutKeysetPage

Create an instance: `local payout_keyset_page = client:PayoutKeysetPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountID` | `string` |  |
| `amount` | `number` |  |
| `amountMinorUnits` | `number` |  |
| `approvePayoutUrl` | `string` |  |
| `approverID` | `string` |  |
| `authenticationMethods` | `table` |  |
| `authorisations` | `table` |  |
| `authorisersCompletedCount` | `number` |  |
| `authorisersRequiredCount` | `number` |  |
| `batchPayoutID` | `string` |  |
| `beneficiary` | `table` |  |
| `canAuthorise` | `boolean` |  |
| `canProcess` | `boolean` |  |
| `canUpdate` | `boolean` |  |
| `chargeBearer` | `string` |  |
| `createdBy` | `string` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` |  |
| `currentUserID` | `string` |  |
| `description` | `string` |  |
| `destination` | `table` |  |
| `documents` | `table` |  |
| `events` | `table` |  |
| `formattedAmount` | `string` |  |
| `formattedFxDestinationAmount` | `string` |  |
| `formattedSchedule` | `string` |  |
| `formattedScheduleDayOnly` | `string` |  |
| `formattedSourceAccountAvailableBalance` | `string` |  |
| `fxDestinationAmount` | `number` |  |
| `fxDestinationAmountMinorUnits` | `number` |  |
| `fxDestinationCurrency` | `string` |  |
| `fxQuoteExpiresAt` | `string` |  |
| `fxQuoteID` | `string` |  |
| `fxRate` | `number` |  |
| `fxUseDestinationAmount` | `boolean` |  |
| `hasCurrentUserAuthorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoiceID` | `string` |  |
| `isArchived` | `boolean` |  |
| `isFailed` | `boolean` |  |
| `isSettled` | `boolean` |  |
| `isSubmitted` | `boolean` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantTokenDescription` | `string` |  |
| `nonce` | `string` |  |
| `paymentProcessor` | `string` |  |
| `paymentRail` | `string` |  |
| `payrunID` | `string` |  |
| `payrunName` | `string` |  |
| `rule` | `table` |  |
| `scheduleDate` | `string` |  |
| `scheduled` | `boolean` |  |
| `sourceAccountAvailableBalance` | `number` |  |
| `sourceAccountAvailableBalanceMinorUnits` | `number` |  |
| `sourceAccountBic` | `string` |  |
| `sourceAccountCurrency` | `string` |  |
| `sourceAccountIban` | `string` |  |
| `sourceAccountIdentifier` | `table` |  |
| `sourceAccountName` | `string` |  |
| `sourceAccountNumber` | `string` |  |
| `sourceAccountSortcode` | `string` |  |
| `status` | `string` |  |
| `tags` | `table` |  |
| `theirReference` | `string` |  |
| `topupPayrunID` | `string` |  |
| `transactedAmount` | `number` |  |
| `transactedFxAmount` | `number` |  |
| `transactedFxRate` | `number` |  |
| `type` | `string` |  |
| `userID` | `string` |  |
| `yourReference` | `string` |  |

#### Example: List

```lua
local payout_keyset_pages, err = client:PayoutKeysetPage():list()
```


### PayoutMetric

Create an instance: `local payout_metric = client:PayoutMetric(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```lua
local payout_metric, err = client:PayoutMetric():load()
```


### Payrun

Create an instance: `local payrun = client:Payrun(nil)`

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
| `authorisations` | `table` |  |
| `authorisersCompletedCount` | `number` |  |
| `authorisersRequiredCount` | `number` |  |
| `batchPayoutID` | `string` |  |
| `canAuthorise` | `boolean` |  |
| `canDelete` | `boolean` |  |
| `canEdit` | `boolean` |  |
| `events` | `table` |  |
| `hasCurrentUserAuthorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoices` | `table` |  |
| `invoicesMinimal` | `table` |  |
| `isArchived` | `boolean` |  |
| `lastUpdated` | `string` |  |
| `lastUpdatedBy` | `table` |  |
| `merchantID` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `notes` | `string` |  |
| `payments` | `table` |  |
| `payouts` | `table` |  |
| `payoutsCount` | `number` |  |
| `reason` | `string` |  |
| `scheduleDate` | `string` |  |
| `scheduledDate` | `string` |  |
| `sourceAccounts` | `table` |  |
| `status` | `string` |  |
| `totalEur` | `number` |  |
| `totalGbp` | `number` |  |
| `totalUsd` | `number` |  |

#### Example: Load

```lua
local payrun, err = client:Payrun():load({ id = "payrun_id" })
```

#### Example: List

```lua
local payruns, err = client:Payrun():list()
```

#### Example: Create

```lua
local payrun, err = client:Payrun():create({
  id = "example_id", -- string
  lastUpdatedBy = {}, -- table
})
```


### Report

Create an instance: `local report = client:Report(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |


### ReportResult

Create an instance: `local report_result = client:ReportResult(nil)`

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
| `statementNumber` | `number` |  |

#### Example: Load

```lua
local report_result, err = client:ReportResult():load({ id = 1, report_id = "report_id" })
```


### Role

Create an instance: `local role = client:Role(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failedRoles` | `table` |  |
| `roles` | `table` |  |

#### Example: Create

```lua
local role, err = client:Role():create({
  merchant_id = "example_merchant_id", -- string
})
```


### Rule

Create an instance: `local rule = client:Rule(nil)`

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
| `account` | `table` |  |
| `accountID` | `string` |  |
| `approveUrl` | `string` |  |
| `approverID` | `string` |  |
| `authenticationMethods` | `table` |  |
| `authorisations` | `table` |  |
| `authorisersCompletedCount` | `number` |  |
| `authorisersRequiredCount` | `number` |  |
| `canAuthorise` | `boolean` |  |
| `createdBy` | `table` |  |
| `description` | `string` |  |
| `endAt` | `string` |  |
| `hasCurrentUserAuthorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isDisabled` | `boolean` |  |
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
| `sweepAction` | `table` |  |
| `timeZoneId` | `string` |  |
| `triggerCronExpression` | `string` |  |
| `triggerOnPayIn` | `boolean` |  |
| `userID` | `string` |  |
| `webHookSecret` | `string` |  |

#### Example: Load

```lua
local rule, err = client:Rule():load({ id = "rule_id" })
```

#### Example: List

```lua
local rules, err = client:Rule():list()
```

#### Example: Create

```lua
local rule, err = client:Rule():create({
  createdBy = {}, -- table
  nonce = "example_nonce", -- string
})
```


### RuleEvent

Create an instance: `local rule_event = client:RuleEvent(nil)`

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
| `isAuthoriseToEnable` | `boolean` |  |
| `message` | `string` |  |
| `rawResponse` | `string` |  |
| `ruleEventType` | `string` |  |
| `ruleID` | `string` |  |
| `user` | `table` |  |

#### Example: List

```lua
local rule_events, err = client:RuleEvent():list()
```


### Tag

Create an instance: `local tag = client:Tag(nil)`

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

```lua
local tags, err = client:Tag():list()
```

#### Example: Create

```lua
local tag, err = client:Tag():create({
  merchant_id = "example_merchant_id", -- string
  merchantID = "example_merchantID", -- string
  name = "example_name", -- string
})
```


### Token

Create an instance: `local token = client:Token(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```lua
local token, err = client:Token():create({
  id = "example_id", -- string
})
```


### Transaction

Create an instance: `local transaction = client:Transaction(nil)`

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
| `accountSequenceNumber` | `number` |  |
| `addressDetails` | `table` |  |
| `amount` | `number` |  |
| `amountMinorUnits` | `number` |  |
| `balance` | `number` |  |
| `balanceMinorUnits` | `number` |  |
| `bookingDateTime` | `string` |  |
| `chargeDetails` | `table` |  |
| `content` | `table` |  |
| `counterparty` | `table` |  |
| `counterpartySummary` | `string` |  |
| `currency` | `string` |  |
| `currencyExchange` | `table` |  |
| `date` | `string` |  |
| `description` | `string` |  |
| `enrichment` | `table` |  |
| `fxAmount` | `number` |  |
| `fxCurrency` | `string` |  |
| `fxRate` | `number` |  |
| `grossAmount` | `table` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isoBankTransactionCode` | `table` |  |
| `merchant` | `table` |  |
| `merchantID` | `string` |  |
| `pageNumber` | `number` |  |
| `pageSize` | `number` |  |
| `payeeDetails` | `table` |  |
| `payerDetails` | `table` |  |
| `paymentRequestCustomFields` | `table` |  |
| `paymentRequestID` | `string` |  |
| `payoutID` | `string` |  |
| `proprietaryBankTransactionCode` | `table` |  |
| `rawReference` | `string` |  |
| `reference` | `string` |  |
| `ruleID` | `string` |  |
| `statementReferences` | `table` |  |
| `status` | `string` |  |
| `supplementaryData` | `any` |  |
| `tags` | `table` |  |
| `theirReference` | `string` |  |
| `totalPages` | `number` |  |
| `totalSize` | `number` |  |
| `transactionAmount` | `table` |  |
| `transactionDate` | `string` |  |
| `transactionInformation` | `table` |  |
| `transactionMutability` | `string` |  |
| `type` | `string` |  |
| `valueDateTime` | `string` |  |
| `virtualIBAN` | `string` |  |
| `yourReference` | `string` |  |

#### Example: Load

```lua
local transaction, err = client:Transaction():load({ id = "transaction_id" })
```

#### Example: List

```lua
local transactions, err = client:Transaction():list()
```

#### Example: Create

```lua
local transaction, err = client:Transaction():create({
  id = "example_id", -- string
  grossAmount = {}, -- table
  payeeDetails = {}, -- table
  payerDetails = {}, -- table
  transactionAmount = {}, -- table
})
```


### User

Create an instance: `local user = client:User(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `clientSessionTimeouts` | `table` |  |
| `emailAddress` | `string` |  |
| `firstName` | `string` |  |
| `id` | `string` |  |
| `lastName` | `string` |  |
| `passkeyAdded` | `boolean` |  |
| `permissions` | `table` |  |
| `profile` | `string` |  |
| `rolesWithScope` | `table` |  |
| `twoFactorEnabled` | `boolean` |  |
| `userInviteID` | `string` |  |

#### Example: List

```lua
local users, err = client:User():list()
```


### UserInvite

Create an instance: `local user_invite = client:UserInvite(nil)`

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
| `authorisationStatus` | `table` |  |
| `failedUserInvites` | `table` |  |
| `id` | `string` |  |
| `initialRoleID` | `string` |  |
| `inviteeEmailAddress` | `string` |  |
| `inviteeFirstName` | `string` |  |
| `inviteeLastName` | `string` |  |
| `inviterEmailAddress` | `string` |  |
| `inviterFirstName` | `string` |  |
| `inviterLastName` | `string` |  |
| `isAuthorised` | `boolean` |  |
| `isInviteeRegistered` | `boolean` |  |
| `lastInvited` | `string` |  |
| `merchantID` | `string` |  |
| `merchantName` | `string` |  |
| `message` | `string` |  |
| `registrationUrl` | `string` |  |
| `sendInviteEmail` | `boolean` |  |
| `status` | `string` |  |
| `user` | `table` |  |
| `userID` | `string` |  |
| `userInvites` | `table` |  |

#### Example: Load

```lua
local user_invite, err = client:UserInvite():load({ id = "user_invite_id" })
```

#### Example: List

```lua
local user_invites, err = client:UserInvite():list()
```

#### Example: Create

```lua
local user_invite, err = client:UserInvite():create({
  user = {}, -- table
})
```


### Virtual

Create an instance: `local virtual = client:Virtual(nil)`

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
| `availableBalance` | `number` |  |
| `availableBalanceMinorUnits` | `number` |  |
| `balance` | `number` |  |
| `balanceMinorUnits` | `number` |  |
| `bankName` | `string` |  |
| `consentID` | `string` |  |
| `createdBy` | `table` |  |
| `createdByDisplayName` | `string` |  |
| `currency` | `string` |  |
| `defaultPaymentRail` | `string` |  |
| `displayName` | `string` |  |
| `expiryDate` | `string` |  |
| `externalAccountIcon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `table` |  |
| `inserted` | `string` |  |
| `isArchived` | `boolean` |  |
| `isConnectedAccount` | `boolean` |  |
| `isDefault` | `boolean` |  |
| `isTrustAccount` | `boolean` |  |
| `isVirtual` | `boolean` |  |
| `lastTransaction` | `table` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantName` | `string` |  |
| `name` | `string` |  |
| `physicalAccountID` | `string` |  |
| `rules` | `table` |  |
| `submittedPayoutsBalance` | `number` |  |
| `submittedPayoutsBalanceMinorUnits` | `number` |  |
| `summary` | `string` |  |
| `supplierSepaInstantStatus` | `string` |  |
| `xeroBankFeedConnectionStatus` | `string` |  |
| `xeroBankFeedLastSyncedAt` | `string` |  |
| `xeroBankFeedSyncLastFailedAt` | `string` |  |
| `xeroBankFeedSyncLastFailureReason` | `string` |  |
| `xeroBankFeedSyncStatus` | `string` |  |
| `xeroUnsynchronisedTransactionsCount` | `number` |  |

#### Example: Create

```lua
local virtual, err = client:Virtual():create({
  account_id = "example_account_id", -- string
  createdBy = {}, -- table
  identifier = {}, -- table
  name = "example_name", -- string
})
```


### Webhook

Create an instance: `local webhook = client:Webhook(nil)`

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
| `isActive` | `boolean` |  |
| `merchantID` | `string` |  |
| `notificationMethod` | `string` |  |
| `resourceTypes` | `table` |  |
| `retry` | `boolean` |  |
| `secret` | `string` |  |
| `version` | `number` |  |

#### Example: Load

```lua
local webhook, err = client:Webhook():load({ id = "webhook_id" })
```

#### Example: List

```lua
local webhooks, err = client:Webhook():list()
```

#### Example: Create

```lua
local webhook, err = client:Webhook():create({
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

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── nofrixion_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`nofrixion_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```lua
local consent = client:Consent()
consent:list()

-- consent:data_get() now returns the consent data from the last list
-- consent:match_get() returns the last match criteria
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
