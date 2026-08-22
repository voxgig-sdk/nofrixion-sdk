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
local created, err = client:Account():create({ account_id = "example_account_id", currency = "example_currency", createdBy = {}, identifier = {} })
if err then error(err) end

-- Update
client:Account():update({ id = created:data_get()["id"], accountBalances = {}, accountID = "example_accountID" })

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
| `accountBalances` | `table` | The various balances for the account. |
| `accountID` | `string` | ID of the account. |
| `accountIdentifications` | `table` | The canoncial identifiers for the account. |
| `accountName` | `string` | Name for the account |
| `accountNames` | `table` | Optional account names set by the account holder. |
| `accountSupplierName` | `string` | The payment account supplier name. |
| `accountType` | `string` | The type of account e.g. |
| `availableBalance` | `number` | The current available balance of the account. |
| `availableBalanceMinorUnits` | `number` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `number` | Balance of the account. |
| `balanceMinorUnits` | `number` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `string` | The bank name for external accounts |
| `consentID` | `string` | The ID of the consent used to connect the external account. |
| `consolidatedAccountInformation` | `table` | Summary information regarding account balances of the overall account provided by the bank. |
| `createdBy` | `table` |  |
| `createdByDisplayName` | `string` | Either the name of the user, merchant token or api key that created the account |
| `currency` | `string` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `string` | Indicates the default payment rail for this account. |
| `description` | `string` | Product name as defined by the financial institution for this account. |
| `details` | `string` | Supplementary specifications that might be provided by the Bank. |
| `displayName` | `string` | Gets a unique display name for the payment account. |
| `expiryDate` | `string` | The date that the external account will expire |
| `externalAccountIcon` | `string` | The Icon for external accounts |
| `format` | `string` | File format to save the statement as. |
| `fromDate` | `string` | Minimum transaction date for the statement. |
| `id` | `string` | Unique id for the account. |
| `identifier` | `table` |  |
| `inserted` | `string` | Timestamp when the account was created. |
| `isArchived` | `boolean` | Indicates whether the account is archived. |
| `isConnectedAccount` | `boolean` | Indicates if the payment account is an externally connected account. |
| `isDefault` | `boolean` | Is the default account |
| `isTrustAccount` | `boolean` | Indicates if the payment account is a trust account. |
| `isVirtual` | `boolean` | True if the account is a virtual account. |
| `lastTransaction` | `table` |  |
| `lastUpdated` | `string` | Timestamp when the account was last updated. |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `merchantName` | `string` | The name of the merchant that owns the account. |
| `nickname` | `string` | Nickname of the account that was provided by the account owner. |
| `physicalAccountID` | `string` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `roleIDs` | `table` | Optional list of role IDs that will get access to the payment account when created. |
| `rules` | `table` | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `number` | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `number` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `string` | Gets a summary of the payments account's most important properties. |
| `supplierPhysicalAccountID` | `string` | For internal use only. |
| `supplierSepaInstantStatus` | `string` | Indicates the status of the SEPA Instant payment rail for this account. |
| `toDate` | `string` | Maximum transaction date for the statement. |
| `type` | `string` | Specifies the type of account e.g. |
| `usageType` | `string` |  |
| `xeroBankFeedConnectionStatus` | `string` | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `string` |  |
| `xeroBankFeedSyncLastFailedAt` | `string` |  |
| `xeroBankFeedSyncLastFailureReason` | `string` |  |
| `xeroBankFeedSyncStatus` | `string` |  |
| `xeroUnsynchronisedTransactionsCount` | `number` | Indicates the number of unsynchronised transactions with Xero |

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
  account_id = "example_account_id", -- string
  currency = "example_currency", -- string
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
| `approveUrl` | `string` | This field is used when returning a batch payout record to a client. |
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
| `authenticationMethods` | `table` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `table` | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `number` | The number of distinct authorisers that have authorised the beneficiary. |
| `authorisersRequiredCount` | `number` | The number of authorisers required for this beneficiary. |
| `beneficiaries` | `table` |  |
| `beneficiaryEvents` | `table` |  |
| `canAuthorise` | `boolean` | True if the beneficiary can be authorised by the user who loaded it. |
| `canUpdate` | `boolean` | True if the beneficiary can be updated by the user who loaded it. |
| `createdBy` | `table` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` | Gets or Sets the currency. |
| `destination` | `table` |  |
| `failedBeneficiaries` | `table` |  |
| `hasCurrentUserAuthorised` | `boolean` | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isEnabled` | `boolean` |  |
| `lastAuthorised` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` | Gets or Sets the merchant id. |
| `name` | `string` | The descriptive name for the beneficiary. |
| `nonce` | `string` |  |
| `sourceAccountIDs` | `table` | ID of the accounts which are authorised to act as a source for the beneficiary. |
| `sourceAccounts` | `table` |  |
| `theirReference` | `string` | The reference that will be used by default as TheirReference when creating payouts to this beneficiary if no TheirReference is specified for the payout. |

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
  id = "example_id", -- string
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
| `groupMembers` | `table` | The existing group members. |
| `groupName` | `string` | The descriptive name for the beneficiary group. |
| `id` | `string` |  |
| `inserted` | `string` | Timestamp indicating when the group was created. |
| `lastUpdated` | `string` | Timestamp indicating when the group was last updated. |
| `merchantID` | `string` | Gets or Sets the merchant id. |

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
| `isPayerAuthenticationRequired` | `boolean` | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `boolean` | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `string` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `string` | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `string` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `number` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `number` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `string` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `string` |  |
| `requestID` | `string` |  |
| `responseCode` | `string` |  |
| `responseType` | `string` |  |
| `status` | `string` |  |
| `threeDSRedirectUrl` | `string` | Checkout.com require a redirect for 3DS authentication. |
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
| `cardType` | `string` | The type of the tokenised card, e.g. |
| `customerEmailAddress` | `string` | When creating a tokenised card the payer's email address must be supplied. |
| `expiryMonth` | `string` |  |
| `expiryYear` | `string` |  |
| `id` | `string` | The unique ID of the card token that has been stored for the customer. |
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
| `isPayerAuthenticationRequired` | `boolean` | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `boolean` | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `string` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `string` | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `string` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `number` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `number` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `string` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `string` |  |
| `requestID` | `string` |  |
| `responseCode` | `string` |  |
| `responseType` | `string` |  |
| `status` | `string` |  |
| `threeDSRedirectUrl` | `string` | Checkout.com require a redirect for 3DS authentication. |
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
| `authorisationUrl` | `string` | The URL the authorising user needs to be redirected to in order to get the open banking consent token. |
| `callbackUrl` | `string` | Optional callback URL that the end user performing the open banking authorisation will be redirected to on completion. |
| `consentID` | `string` | The ID of the open banking consent. |
| `emailAddress` | `string` | The email address that identifies the end user that will be authorising the open banking consent request. |
| `expiryDate` | `string` |  |
| `failureCallbackUrl` | `string` | Optional callback URL for open banking consent authorisation failure. |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `institutionID` | `string` | The institution ID the open banking consent is being requested for. |
| `isConnectedAccounts` | `boolean` | Optional setting. |
| `isEnabled` | `boolean` |  |
| `merchantID` | `string` | The ID of the merchant the consent token is being created to be used with. |
| `provider` | `string` | Lists the supported card and PIS processors. |
| `successWebHookUrl` | `string` | A web hook URL to send an HTTP request to when the open banking consent is successfuly authorised. |

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
| `failedSubmissions` | `table` | Dictionary of failed submissions, keyed by the index (1-based) in the original request. |
| `successfulSubmissions` | `table` | List of successfully submitted direct debit payments. |

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
| `exchangeRate` | `number` | The price at which the transaction will buy the source currency using the destination currency. |
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
| `accountNumber` | `string` | Account number of the customer's bank account in case of GBP account. |
| `addressLine1` | `string` | First line of the customer's address. |
| `addressLine2` | `string` | Second line of the customer's address. |
| `approvedAt` | `string` | Date at which the supplier approved this mandate. |
| `city` | `string` | Customer's city. |
| `countryCode` | `string` | 2-character country code of the customer's bank account. |
| `currency` | `string` | Currency of this mandate. |
| `customerAccountNumber` | `string` | Customer's account number in case of GBP account. |
| `customerCity` | `string` | Customer's city of residence. |
| `customerCountryCode` | `string` | Customer's country of residence code. |
| `customerCountryName` | `string` | Customer's country of residence. |
| `customerEmailAddress` | `string` | Customer's email address. |
| `customerFirstName` | `string` | Customer's first name. |
| `customerIban` | `string` | Customer's IBAN in case of EUR account. |
| `customerLastName` | `string` | Customer's last name. |
| `customerSortCode` | `string` | Customer's sort code in case of GBP account. |
| `emailAddress` | `string` | Customer's email address. |
| `firstName` | `string` | Customer's first name. |
| `iban` | `string` | IBAN of the customer's bank account in case of EUR account. |
| `id` | `string` | Internal ID of the mandate. |
| `inserted` | `string` | The timestamp this mandate was created at. |
| `isRecurring` | `boolean` | Whether this mandate is single-use or recurring. |
| `lastName` | `string` | Customer's last name. |
| `lastUpdated` | `string` | The timestamp this mandate was last updated at. |
| `merchantID` | `string` | Internal ID of this mandate's merchant. |
| `postalCode` | `string` | Customer's postal code. |
| `reference` | `string` | Reference assigned to this mandate. |
| `sortCode` | `string` | Sort code of the customer's bank account in case of GBP account. |
| `status` | `string` | General status of this mandate. |
| `supplierBankAccountID` | `string` | ID that the supplier assigned to this mandate's bank account. |
| `supplierCustomerID` | `string` | ID that the supplier assigned to this mandate's customer. |
| `supplierMandateID` | `string` | ID that the supplier assigned to this mandate. |
| `supplierName` | `string` | Name of the supplier used to create this mandate. |
| `supplierStatus` | `string` | Last status that the supplier reported for this mandate. |

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
| `accountCurrencies` | `table` | The list of currencies that the merchant has accounts for. |
| `canHaveTrustAccounts` | `boolean` | Trust accounts are a special type of account that allow the account name to be trusted for use in statements and verification of payee checks. |
| `cardPaymentProcessor` | `string` | Name of the card payment processor. |
| `companyID` | `string` | The Company ID recorded in the Compliance system. |
| `displayQrOnHostedPay` | `boolean` | Indicates if a QR Code containing the payment link should be displayed on the hosted payment page. |
| `hostedPayVersion` | `number` | The version of the hosted payment page to use with the merchant. |
| `id` | `string` | Unique ID for the merchant. |
| `inserted` | `string` | Timestamp the merchant was added to MoneyMoov. |
| `isBlocked` | `boolean` | The merchant is blocked from making payments (payouts). |
| `isExited` | `boolean` | The merchant has formally terminated their relationship and is no longer a customer. |
| `isSuspended` | `boolean` | The merchant has temporarily suspended their own account. |
| `jurisdiction` | `string` | The jurisdiction the merchant entity is incorporated or established in. |
| `logoUrlPng` | `string` | The CDN URL of the merchant's logo in PNG format. |
| `logoUrlSvg` | `string` | The CDN URL of the merchant's logo in SVG format. |
| `merchantCategoryCode` | `string` | The industry code that represents the merchant's primary trading activity. |
| `name` | `string` | The registered business name of the merchant. |
| `notes` | `string` | The notes field is an optional free text field that can be used to store any additional information about the merchant. |
| `parentMerchant` | `table` |  |
| `paymentAccountLimit` | `number` | The maximum number of payment accounts that can be created for the Merchant. |
| `paymentAccounts` | `table` |  |
| `reason` | `string` | The reason for the suspension. |
| `shortName` | `string` | A URL friendly shortish name for the merchant. |
| `supportedPaymentMethodsList` | `table` | The payment methods that are configured and supported for this merchant. |
| `suspensionReason` | `string` | The reason for the suspension, provided by the merchant. |
| `tags` | `table` | An optional list of descriptive tags that can be used on merchant entities such as payment requests. |
| `timeZoneId` | `string` | The IANA (Internet Assigned Numbers Authority) time zone identifier of the merchant. |
| `tradingName` | `string` | An optional trading name. |
| `webHookLimit` | `number` | The maximum number of web hooks that can be created for the Merchant. |
| `yourRoleName` | `string` | The name of the role for the identity that loaded the merchant record. |

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
| `approvedAt` | `string` | Date at which the supplier approved this mandate. |
| `currency` | `string` | Currency of this mandate. |
| `customerAccountNumber` | `string` | Customer's account number in case of GBP account. |
| `customerCity` | `string` | Customer's city of residence. |
| `customerCountryCode` | `string` | Customer's country of residence code. |
| `customerCountryName` | `string` | Customer's country of residence. |
| `customerEmailAddress` | `string` | Customer's email address. |
| `customerFirstName` | `string` | Customer's first name. |
| `customerIban` | `string` | Customer's IBAN in case of EUR account. |
| `customerLastName` | `string` | Customer's last name. |
| `customerSortCode` | `string` | Customer's sort code in case of GBP account. |
| `id` | `string` | Internal ID of the mandate. |
| `inserted` | `string` | The timestamp this mandate was created at. |
| `isRecurring` | `boolean` | Whether this mandate is single-use or recurring. |
| `lastUpdated` | `string` | The timestamp this mandate was last updated at. |
| `merchantID` | `string` | Internal ID of this mandate's merchant. |
| `reference` | `string` | Reference assigned to this mandate. |
| `status` | `string` | General status of this mandate. |
| `supplierBankAccountID` | `string` | ID that the supplier assigned to this mandate's bank account. |
| `supplierCustomerID` | `string` | ID that the supplier assigned to this mandate's customer. |
| `supplierMandateID` | `string` | ID that the supplier assigned to this mandate. |
| `supplierName` | `string` | Name of the supplier used to create this mandate. |
| `supplierStatus` | `string` | Last status that the supplier reported for this mandate. |

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
| `bankCountryCodes` | `table` | The list of country codes representing the banks the country supports. |
| `bankID` | `string` | ID of the bank to be configured for the merchant. |
| `bankName` | `string` | Name of the Bank/Institution. |
| `businessInstitutionID` | `string` | ID that the processor uses to identify the bank (business accounts). |
| `currency` | `string` | Currency supported by the bank. |
| `logo` | `string` | URL of the bank's logo. |
| `message` | `string` | Message relating to specific bank. |
| `messageImageUrl` | `string` | Optional image URL to be displayed with the message. |
| `order` | `number` | Order in which this setting will appear in the UI. |
| `personalInstitutionID` | `string` | ID that the processor uses to identify the bank (personal accounts). |
| `processor` | `string` | Name of the bank payment processor. |
| `warningHeading` | `string` | The heading for a warning message related to the bank institution to be displayed to the user. |
| `warningMessage` | `string` | The warning message related to the bank institution to be displayed to the user. |

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
| `customFields` | `table` | A list of custom fields that can be included in the payment request template. |
| `defaultFields` | `table` | A list of default fields that are included in the payment request template. |
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
| `authenticationMethods` | `table` | A list of authentication types allowed to authorise the merchant token. |
| `authorisations` | `table` | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `number` | The number of distinct authorisers that have authorised the merchant token. |
| `authorisersRequiredCount` | `number` | The number of authorisers required for this merchant token. |
| `canAuthorise` | `boolean` | True if the merchant token can be authorised by the user who loaded it. |
| `description` | `string` | Token description |
| `expiresAt` | `string` | Optional. |
| `hasCurrentUserAuthorised` | `boolean` | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `hmacAlgorithm` | `string` | Optional shared secret algorithm to use for HMAC authentication. |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `ipAddressWhitelist` | `string` | Optional. |
| `isArchived` | `boolean` | Indicates whether the merchant token is archived. |
| `isEnabled` | `boolean` | If set to false the merchant token will not be accepted to authorise a request. |
| `lastAuthorised` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` | The merchant id to add to the token |
| `nonce` | `string` |  |
| `permissionTypes` | `table` | The permissions that the merchant token supports. |
| `requestSignatureVersion` | `number` | Represent the version of the overall merchant token. |
| `sharedSecretAlgorithm` | `string` | Optional shared secret algorithm to use for HMAC authentication. |
| `sharedSecretBase64` | `string` | The base 64 encoded shared secret that is used for request authentication with an HMAC. |
| `token` | `string` | The JWT merchant token. |

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
| `accountName` | `string` | The name of the account to verify |
| `accountNumber` | `string` | The account number of the account to verify (for CoP checks) |
| `iban` | `string` | The IBAN of the account to verify (for VoP checks) |
| `payeeVerifiedAccountName` | `string` | The verified account name of the payee, if available (in case of a close match) |
| `result` | `string` | The result of the payee verification |
| `secondaryIdentification` | `string` | Optional secondary identifier for the account to verify. |
| `sortCode` | `string` | The sort code of the account to verify (for CoP checks) |

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
| `amount` | `number` | The amount of money to request. |
| `amountPending` | `number` | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `number` | Total amount received for this payment request. |
| `amountRefunded` | `number` | Total amount refunded for this payment request. |
| `autoSendReceipt` | `boolean` | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `string` | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `string` | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `boolean` | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `boolean` | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `string` | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `boolean` | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardNoPayerAuthentication` | `boolean` | If set to true for card payments no attempt will be made to use payer authentication (3-D Secure and equivalent). |
| `cardProcessorMerchantID` | `string` | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `string` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `string` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `cardTransmitRawDetails` | `boolean` | If set to true for card payments the sensitive card number and card verification number will be transmitted directly rather than being tokenised. |
| `createdByUser` | `table` |  |
| `currency` | `string` | The currency of the request. |
| `customFields` | `table` | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `string` | Optional email address for the customer. |
| `customerID` | `string` | An optional customer identifier for the payment request. |
| `customerName` | `string` |  |
| `description` | `string` | An optional description for the payment request. |
| `destinationAccount` | `table` |  |
| `directDebitPayment` | `table` | Contains information about a Direct Debit payment attempt for a payment request. |
| `dueDate` | `string` | The due date for the payment request. |
| `events` | `table` |  |
| `failureCallbackUrl` | `string` | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `table` | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `string` |  |
| `hostedPayCheckoutUrl` | `string` | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `string` |  |
| `ignoreAddressVerification` | `boolean` | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `string` | The timestamp the payment request was created at. |
| `insertedSortable` | `string` | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `isArchived` | `boolean` | Indicates whether the payment request is archived. |
| `jwk` | `string` | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `string` | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `string` | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `string` | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `string` | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `string` | The ID of the merchant to create the payment request for. |
| `merchantTokenDescription` | `string` | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `string` |  |
| `notificationRoleIDs` | `table` | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `string` | An optional order ID for the payment request. |
| `partialPaymentMethod` | `string` | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `string` | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `table` | The payment attempts made against this payment request. |
| `paymentMethods` | `table` | The payment methods that the payment request supports. |
| `paymentProcessor` | `string` | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `payrunID` | `string` | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `string` | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `string` | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `table` |  |
| `sandboxSettleDelayInSeconds` | `number` | Sandbox only. |
| `shippingAddress` | `table` |  |
| `shippingAddressCity` | `string` | Optionally the city of the customer's shipping address. |
| `shippingAddressCountryCode` | `string` | Optionally the country code of the customer's shipping address. |
| `shippingAddressCounty` | `string` | Optionally the state or county of the customer's shipping address. |
| `shippingAddressLine1` | `string` | Optionally the first line of the customer's shipping address. |
| `shippingAddressLine2` | `string` | Optionally the second line of the customer's shipping address. |
| `shippingAddressPostCode` | `string` | Optionally the post code of the customer's shipping address. |
| `shippingEmail` | `string` | Optionally the shipping email address for the customer. |
| `shippingFirstName` | `string` | Optionally the first name of the customer's shipping address. |
| `shippingLastName` | `string` | Optionally the last name of the customer's shipping address. |
| `shippingPhone` | `string` | Optionally the shipping phone number for the customer. |
| `status` | `string` | The current status of the payment request. |
| `successWebHookUrl` | `string` | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tagIds` | `table` | An optional list of tag ids to add to the payment request |
| `tags` | `table` | An optional list of descriptive tags attached to the payment request. |
| `title` | `string` | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `table` |  |
| `transactions` | `table` |  |
| `useHostedPaymentPage` | `boolean` | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

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
| `accountName` | `string` | Name for the account |
| `accountSupplierName` | `string` | The payment account supplier name. |
| `availableBalance` | `number` | The current available balance of the account. |
| `availableBalanceMinorUnits` | `number` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `number` | Balance of the account. |
| `balanceMinorUnits` | `number` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `string` | The bank name for external accounts |
| `consentID` | `string` | The ID of the consent used to connect the external account. |
| `createdBy` | `table` |  |
| `createdByDisplayName` | `string` | Either the name of the user, merchant token or api key that created the account |
| `currency` | `string` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `string` | Indicates the default payment rail for this account. |
| `displayName` | `string` | Gets a unique display name for the payment account. |
| `expiryDate` | `string` | The date that the external account will expire |
| `externalAccountIcon` | `string` | The Icon for external accounts |
| `id` | `string` | Unique id for the account. |
| `identifier` | `table` |  |
| `inserted` | `string` | Timestamp when the account was created. |
| `isArchived` | `boolean` | Indicates whether the account is archived. |
| `isConnectedAccount` | `boolean` | Indicates if the payment account is an externally connected account. |
| `isDefault` | `boolean` | Is the default account |
| `isTrustAccount` | `boolean` | Indicates if the payment account is a trust account. |
| `isVirtual` | `boolean` | True if the account is a virtual account. |
| `lastTransaction` | `table` |  |
| `lastUpdated` | `string` | Timestamp when the account was last updated. |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `merchantName` | `string` | The name of the merchant that owns the account. |
| `physicalAccountID` | `string` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `table` | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `number` | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `number` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `string` | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `string` | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `string` | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `string` |  |
| `xeroBankFeedSyncLastFailedAt` | `string` |  |
| `xeroBankFeedSyncLastFailureReason` | `string` |  |
| `xeroBankFeedSyncStatus` | `string` |  |
| `xeroUnsynchronisedTransactionsCount` | `number` | Indicates the number of unsynchronised transactions with Xero |

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
| `accountName` | `string` | Name for the account |
| `availableBalance` | `number` | The current available balance of the account. |
| `balance` | `number` | Balance of the account. |
| `balanceMinorUnits` | `number` | Balance of the account expressed in the currency’s minor units (e.g. |
| `currency` | `string` | Currency of the account in ISO 4217 format |
| `id` | `string` | Unique id for the account. |
| `identifier` | `table` |  |
| `isArchived` | `boolean` | Is the account archived |
| `isConnectedAccount` | `boolean` | Indicates if the payment account is an externally connected account. |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `submittedPayoutsBalance` | `number` | Total of the payouts that have been submitted for processing. |

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
| `paymentInitiationID` | `string` | The unique identifier of the payment initiation request. |
| `paymentRequestCallbackUrl` | `string` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `string` |  |
| `redirectUrl` | `string` | A redirect URL for the user to authorise the payment initiation request at the ASPSP |
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
| `amount` | `number` | The amount of money to request. |
| `amountPending` | `number` | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `number` | Total amount received for this payment request. |
| `amountRefunded` | `number` | Total amount refunded for this payment request. |
| `autoSendReceipt` | `boolean` | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `string` | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `string` | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `boolean` | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `boolean` | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `string` | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `boolean` | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardProcessorMerchantID` | `string` | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `string` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `string` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `createdByUser` | `table` |  |
| `currency` | `string` | The currency of the request. |
| `customFields` | `table` | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `string` | Optional email address for the customer. |
| `customerID` | `string` | An optional customer identifier for the payment request. |
| `customerName` | `string` |  |
| `description` | `string` | An optional description for the payment request. |
| `destinationAccount` | `table` |  |
| `directDebitPayment` | `table` | Contains information about a Direct Debit payment attempt for a payment request. |
| `doSimulateSettlementFailure` | `boolean` |  |
| `dueDate` | `string` | The due date for the payment request. |
| `errorDescription` | `string` |  |
| `events` | `table` |  |
| `failedPaymentRequests` | `table` |  |
| `failureCallbackUrl` | `string` | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `table` | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `string` |  |
| `hostedPayCheckoutUrl` | `string` | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `string` |  |
| `ignoreAddressVerification` | `boolean` | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `string` | The timestamp the payment request was created at. |
| `insertedSortable` | `string` | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `institution` | `string` |  |
| `isArchived` | `boolean` | Indicates whether the payment request is archived. |
| `jwk` | `string` | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `string` | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `string` | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `string` | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `string` | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `string` |  |
| `merchantTokenDescription` | `string` | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `string` |  |
| `notificationRoleIDs` | `table` | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `string` | An optional order ID for the payment request. |
| `partialPaymentMethod` | `string` | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `string` | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `table` | The payment attempts made against this payment request. |
| `paymentInitiationID` | `string` |  |
| `paymentMethods` | `table` | The payment methods that the payment request supports. |
| `paymentProcessor` | `string` | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `paymentRequests` | `table` |  |
| `payrunID` | `string` | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `string` | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `string` | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `table` |  |
| `sandboxSettleDelayInSeconds` | `number` | Sandbox only. |
| `shippingAddress` | `table` |  |
| `status` | `string` | The current status of the payment request. |
| `successWebHookUrl` | `string` | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tags` | `table` | An optional list of descriptive tags attached to the payment request. |
| `title` | `string` | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `table` |  |
| `transactions` | `table` |  |
| `useHostedPaymentPage` | `boolean` | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

#### Example: Load

```lua
local payment_request, err = client:PaymentRequest():load({ id = "payment_request_id" })
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
| `applePayTransactionID` | `string` | Transaction ID received in Apple pay token. |
| `cardAuthorizationResponseID` | `string` | For a successful card authorization this field will hold the response ID. |
| `cardExpiryMonth` | `number` | For card payment events this field holds the payer's card expiry month. |
| `cardExpiryYear` | `number` | For card payment events this field holds the payer's card expiry year. |
| `cardIssuer` | `string` | For card payment events this field holds the payer's card issuer. |
| `cardIssuerCountry` | `string` | For card payment events this field holds the payer's card issuer country of origin. |
| `cardLastFourDigits` | `string` | For card payment events this field holds the payer's card last four digits. |
| `cardRequestID` | `string` |  |
| `cardScheme` | `string` | For card payment events this field holds the scheme of the payer's card, e.g. |
| `cardTokenCustomerID` | `string` | If the option to create a reusable token for card payments was set this field contains the token the merchant can store to use for repeat payments. |
| `cardTransactionID` | `string` |  |
| `currency` | `string` |  |
| `directDebitPaymentID` | `string` | Payment ID issued by the Direct Debit supplier. |
| `directDebitPaymentReference` | `string` | Reference string issued by the Direct Debit supplier. |
| `drirectDebitMandateID` | `string` | The ID of the mandate that was used wehn requesting payment. |
| `errorMessage` | `string` |  |
| `errorReason` | `string` |  |
| `eventType` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lightningInvoice` | `string` | For Bitcoin Lightning payments this field holds the invoice presented to the payer. |
| `lightningRHash` | `string` | For Bitcoin Lightning payments the hash of the invoice presented to the payer. |
| `originUrl` | `string` | Optional field that can be set by payment methods, such as pay by bank, that may want to redirect back to the URL that initiated the attempt in the case of a failure condition. |
| `paymentMethodType` | `string` | The type of payment method the event relates to, e.g. |
| `paymentProcessorName` | `string` | If the event was for a card payment this is the name of the card processor, e.g. |
| `paymentRequestID` | `string` |  |
| `pispBankStatus` | `string` | For payment initiation attempts some providers (e.g. |
| `pispPaymentInitiationID` | `string` | For a payment initiation this is the ID returned by the service provider initiating the payment for us. |
| `pispPaymentInstitutionName` | `string` | For a payment initiation this is the name of the financial institution that is used to initiate and authorise the payment. |
| `pispPaymentServiceProviderID` | `string` | For a payment initiation this is the service provider ID selected by the payer, typically the ID for the bank or similar financial institution. |
| `pispRedirectUrl` | `string` | For a payment initiation this is the redirect URL returned by the service provider initiating the payment for us. |
| `reconciledTransactionID` | `string` | For settlement events (only relevant for non-card payments) this is the payin transaction that the payment request event was reconciled with. |
| `refundPayoutID` | `string` | ID of the Payout that was created for refund. |
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
| `amount` | `number` | The amount of money to request. |
| `amountPending` | `number` | The amount of money that was authorised but has not arrived in the account yet. |
| `amountReceived` | `number` | The amount of money that has been received for this payment request. |
| `amountRefunded` | `number` | The amount of money that has been refunded for this payment request. |
| `callbackUrl` | `string` |  |
| `cardStripePaymentIntentSecret` | `string` |  |
| `countryCode` | `string` | The country code associated with the payment. |
| `currency` | `string` | The currency of the request. |
| `customFieldsToDisplay` | `table` | Custom fields to display to the customer. |
| `description` | `string` | An optional description for the payment request. |
| `dueDate` | `string` | The due date of the payment request. |
| `fieldDisplaySettings` | `table` |  |
| `googlePayMerchantID` | `string` | Merchant ID from Google Pay |
| `id` | `string` |  |
| `jwk` | `string` | The jwk containing the public key |
| `merchantID` | `string` |  |
| `merchantLogoUrlPng` | `string` |  |
| `merchantLogoUrlSvg` | `string` |  |
| `merchantName` | `string` |  |
| `merchantShortName` | `string` |  |
| `partialPaymentMethod` | `string` |  |
| `paymentAttempts` | `table` | The payment attempts for this payment request. |
| `paymentMethodsList` | `table` | The payment methods that the payment request supports. |
| `paymentProcessor` | `string` | The card processor |
| `paymentProcessorKey` | `string` | The card processors public key |
| `pispError` | `string` | This is the error returned from the bank which is recorded in payment request events. |
| `priorityBankID` | `string` |  |
| `status` | `string` | The status of the payment request. |
| `stripeAccountID` | `string` | Account ID of connected customers in Stripe |
| `title` | `string` | The title of the payment request. |

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
| `amount` | `number` | The authorised payment amount. |
| `amountPending` | `number` |  |
| `amountReceived` | `number` |  |
| `amountRefunded` | `number` |  |
| `currency` | `string` | The authorised payment currency. |
| `customerID` | `string` | The customer id |
| `paymentRequestID` | `string` | The ID of the payment request the result is for. |
| `payments` | `table` | The list of payment attempts that have been received for the payment request. |
| `pispAuthorizations` | `table` |  |
| `requestedAmount` | `number` | The full original payment amount requested. |
| `result` | `string` | The result of the payment attempt. |

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
| `accountID` | `string` | Gets or Sets Account Id of sending account |
| `allowIncomplete` | `boolean` | If set to true the payout will get created even if the business validation rules fail. |
| `amount` | `number` | Gets or Sets payout amount |
| `amountMinorUnits` | `number` | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `string` | This field is used when returning an payout record to a client. |
| `approverID` | `string` | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `table` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `table` | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `number` | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `number` | The number of authorisers required for this payout. |
| `batchPayoutID` | `string` | The ID of the batch the payout is associated with. |
| `beneficiary` | `table` |  |
| `beneficiaryID` | `string` | Optional. |
| `canAuthorise` | `boolean` | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `boolean` | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `boolean` | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `string` | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `string` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` | Gets or Sets Currency of payout request |
| `currentUserID` | `string` | The ID of the user that requested access to the PayOut record. |
| `description` | `string` | Gets or Sets description of payout request |
| `destination` | `table` |  |
| `documents` | `table` | Documents associated with the payout. |
| `events` | `table` | The activity associated with the payout. |
| `failedPayouts` | `table` |  |
| `formattedAmount` | `string` | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `string` | FX destination currency and amount formatted string. |
| `formattedSchedule` | `string` |  |
| `formattedScheduleDayOnly` | `string` |  |
| `formattedSourceAccountAvailableBalance` | `string` | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `number` | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `number` | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `string` | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `string` | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `string` | Optional. |
| `fxRate` | `number` | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `boolean` | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `boolean` | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `string` | The ID for the payout. |
| `inserted` | `string` |  |
| `invoiceID` | `string` | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `boolean` | Indicates whether the payout is archived. |
| `isFailed` | `boolean` | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `boolean` | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `boolean` | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `string` |  |
| `nonce` | `string` |  |
| `paymentProcessor` | `string` | The usptream payment processor for the payout. |
| `paymentRail` | `string` | Optional field to indicate the payment rail to use for the payout. |
| `payouts` | `table` |  |
| `payrunID` | `string` | The ID of the payrun that this payout is associated with. |
| `payrunName` | `string` | The name of the payrun that this payout is associated with. |
| `reason` | `string` |  |
| `rule` | `table` |  |
| `scheduleDate` | `string` | The date the payout should be submitted. |
| `scheduled` | `boolean` | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `number` | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `number` | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `string` | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `string` | The currency of the source account. |
| `sourceAccountIban` | `string` | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `table` |  |
| `sourceAccountName` | `string` | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `string` | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `string` | The sort code of the account the payout is being made from. |
| `status` | `string` | Gets or Sets the status of payout request |
| `tagIds` | `table` | An optional list of tag ids to add to the payout. |
| `tags` | `table` | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `string` | Gets or Sets destination reference ID |
| `topupPayrunID` | `string` | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `number` | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `number` | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `number` | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `string` | Gets or Sets payout type |
| `userID` | `string` | Gets or Sets User ID of who created the payout request |
| `yourReference` | `string` | Gets or Sets your reference ID |

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
  id = "example_id", -- string
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
| `accountID` | `string` | Gets or Sets Account Id of sending account |
| `amount` | `number` | Gets or Sets payout amount |
| `amountMinorUnits` | `number` | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `string` | This field is used when returning an payout record to a client. |
| `approverID` | `string` | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `table` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `table` | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `number` | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `number` | The number of authorisers required for this payout. |
| `batchPayoutID` | `string` | The ID of the batch the payout is associated with. |
| `beneficiary` | `table` |  |
| `canAuthorise` | `boolean` | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `boolean` | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `boolean` | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `string` | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `string` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` | Gets or Sets Currency of payout request |
| `currentUserID` | `string` | The ID of the user that requested access to the PayOut record. |
| `description` | `string` | Gets or Sets description of payout request |
| `destination` | `table` |  |
| `documents` | `table` | Documents associated with the payout. |
| `events` | `table` | The activity associated with the payout. |
| `formattedAmount` | `string` | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `string` | FX destination currency and amount formatted string. |
| `formattedSchedule` | `string` |  |
| `formattedScheduleDayOnly` | `string` |  |
| `formattedSourceAccountAvailableBalance` | `string` | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `number` | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `number` | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `string` | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `string` | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `string` | Optional. |
| `fxRate` | `number` | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `boolean` | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `boolean` | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `string` | The ID for the payout. |
| `inserted` | `string` |  |
| `invoiceID` | `string` | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `boolean` | Indicates whether the payout is archived. |
| `isFailed` | `boolean` | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `boolean` | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `boolean` | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `string` |  |
| `nonce` | `string` |  |
| `paymentProcessor` | `string` | The usptream payment processor for the payout. |
| `paymentRail` | `string` | Optional field to indicate the payment rail to use for the payout. |
| `payrunID` | `string` | The ID of the payrun that this payout is associated with. |
| `payrunName` | `string` | The name of the payrun that this payout is associated with. |
| `rule` | `table` |  |
| `scheduleDate` | `string` | The date the payout should be submitted. |
| `scheduled` | `boolean` | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `number` | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `number` | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `string` | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `string` | The currency of the source account. |
| `sourceAccountIban` | `string` | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `table` |  |
| `sourceAccountName` | `string` | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `string` | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `string` | The sort code of the account the payout is being made from. |
| `status` | `string` | Gets or Sets the status of payout request |
| `tags` | `table` | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `string` | Gets or Sets destination reference ID |
| `topupPayrunID` | `string` | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `number` | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `number` | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `number` | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `string` | Gets or Sets payout type |
| `userID` | `string` | Gets or Sets User ID of who created the payout request |
| `yourReference` | `string` | Gets or Sets your reference ID |

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
| `authorisations` | `table` | A list of the users who have successfully authorised the latest version of the payrun and when. |
| `authorisersCompletedCount` | `number` | The number of distinct authorisers that have authorised the payrun. |
| `authorisersRequiredCount` | `number` | The number of authorisers required for this payrun. |
| `batchPayoutID` | `string` |  |
| `canAuthorise` | `boolean` | True if the payrun can be authorised by the user who loaded it. |
| `canDelete` | `boolean` |  |
| `canEdit` | `boolean` |  |
| `events` | `table` |  |
| `hasCurrentUserAuthorised` | `boolean` | True if the payrun was loaded for a user and that user has already authorised the latest version of the payrun. |
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
| `accountID` | `string` | The ID of the account the rule will apply to. |
| `approveUrl` | `string` | If set this property holds the URL an approver needs to visit in order to complete a strong authentication check in order to approve the rule. |
| `approverID` | `string` |  |
| `authenticationMethods` | `table` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `table` | A list of the users who have successfully authorised the latest version of the rule and when. |
| `authorisersCompletedCount` | `number` | The number of distinct authorisers that have authorised the rule. |
| `authorisersRequiredCount` | `number` | The number of authorisers required for this rule. |
| `canAuthorise` | `boolean` | True if the rule can be authorised by the user who loaded it. |
| `createdBy` | `table` |  |
| `description` | `string` | Arbitrary description for the rule. |
| `endAt` | `string` | Optional end time for rule executions. |
| `hasCurrentUserAuthorised` | `boolean` | True if the current user has authorised. |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isDisabled` | `boolean` | If set to true the rule will be disabled from executing. |
| `lastExecutedAt` | `string` |  |
| `lastRunAtTransactionDate` | `string` | The most recent transaction date when the rule was last run. |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `name` | `string` | A name to succinctly describe the rule. |
| `nonce` | `string` |  |
| `onApprovedWebHookUrl` | `string` | Optional URL to receive an HTTP request with the rule details when the rule status changes to approved. |
| `onExecutionErrorWebHookUrl` | `string` | Optional URL to receive an HTTP request when a rule execution attempt fails. |
| `onExecutionSuccessWebHookUrl` | `string` | Optional URL to receive an HTTP request when a rule execution attempt succeeds. |
| `startAt` | `string` | Optional start time for rule executions. |
| `status` | `string` |  |
| `sweepAction` | `table` |  |
| `timeZoneId` | `string` | If the rule should be executed on a recurring schedule this is the timezone that the CRON expression should be evaluated in. |
| `triggerCronExpression` | `string` | If the rule should be executed on a recurring schedule this is the expression that sets the schedule. |
| `triggerOnPayIn` | `boolean` | Set to true if the rule execution should be triggered when the account receives a pay in (credit). |
| `userID` | `string` |  |
| `webHookSecret` | `string` | If set this secret will be used to sign Web Hook requests. |

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
| `accountID` | `string` | The ID of the account the transaction belongs to. |
| `accountName` | `string` | The name of the account the transaction belongs to. |
| `accountSequenceNumber` | `number` | The sequence number of transaction on a per account basis. |
| `addressDetails` | `table` |  |
| `amount` | `number` | Amount of the transaction. |
| `amountMinorUnits` | `number` | Amount of the transaction expressed in the currency’s minor units (e.g. |
| `balance` | `number` | Balance left on the account after the transaction. |
| `balanceMinorUnits` | `number` | Balance on the account expressed in the currency’s minor units (e.g. |
| `bookingDateTime` | `string` |  |
| `chargeDetails` | `table` |  |
| `content` | `table` |  |
| `counterparty` | `table` |  |
| `counterpartySummary` | `string` | For pay in (credit) transactions this will contain a descriptive string with the most important fields about the counterparty. |
| `currency` | `string` | Currency of transaction. |
| `currencyExchange` | `table` | Provides details on the currency exchange. |
| `date` | `string` |  |
| `description` | `string` | Description of the transaction. |
| `enrichment` | `table` |  |
| `fxAmount` | `number` | For an FX payout this is the amound in the FX currency. |
| `fxCurrency` | `string` | For an FX payout this is the currency that was received or that was instructed. |
| `fxRate` | `number` | For an FX payout this is the exchange rate between the transaction currency and the FX currency. |
| `grossAmount` | `table` |  |
| `id` | `string` | Unique ID for the transaction. |
| `inserted` | `string` | Date when the transaction was inserted into the ledger. |
| `isoBankTransactionCode` | `table` |  |
| `merchant` | `table` |  |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `pageNumber` | `number` | Current page number. |
| `pageSize` | `number` | Page size |
| `payeeDetails` | `table` | The Payee object contains details of the beneficiary, person or business. |
| `payerDetails` | `table` |  |
| `paymentRequestCustomFields` | `table` | The custom fields that were attached to the payment request that resulted in this transaction. |
| `paymentRequestID` | `string` | For Pay by Bank and Direct Debit transactions this will contain the ID of the payment request. |
| `payoutID` | `string` | ID of the payout that resulted in the transaction. |
| `proprietaryBankTransactionCode` | `table` |  |
| `rawReference` | `string` | The raw payment reference details as received from the payment processor. |
| `reference` | `string` |  |
| `ruleID` | `string` | ID of the rule that resulted in the transaction. |
| `statementReferences` | `table` |  |
| `status` | `string` |  |
| `supplementaryData` | `any` |  |
| `tags` | `table` | An optional list of descriptive tags attached to the transaction. |
| `theirReference` | `string` | For a pay out the reference that the payer attached for the receiving party. |
| `totalPages` | `number` | Total pages |
| `totalSize` | `number` | Total count |
| `transactionAmount` | `table` |  |
| `transactionDate` | `string` | Date when the transaction occurred. |
| `transactionInformation` | `table` |  |
| `transactionMutability` | `string` |  |
| `type` | `string` | Type of the transaction. |
| `valueDateTime` | `string` |  |
| `virtualIBAN` | `string` | If set it indicates the payin was to a virtual IBAN. |
| `yourReference` | `string` | For a pay in the reference the sending party attached. |

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
| `clientSessionTimeouts` | `table` | The number of seconds a session for this user should last before expiring. |
| `emailAddress` | `string` |  |
| `firstName` | `string` |  |
| `id` | `string` |  |
| `lastName` | `string` |  |
| `passkeyAdded` | `boolean` |  |
| `permissions` | `table` |  |
| `profile` | `string` |  |
| `rolesWithScope` | `table` |  |
| `twoFactorEnabled` | `boolean` |  |
| `userInviteID` | `string` | Optional ID of the invite that is being accepted so the user can be assigned a role on a new merchant. |

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
| `initialRoleID` | `string` | The role ID to automatically assign to the merchant’s very first user. |
| `inviteeEmailAddress` | `string` | Email address of the user being invited. |
| `inviteeFirstName` | `string` | First Name of the user being invited. |
| `inviteeLastName` | `string` | Last Name of the user being invited. |
| `inviterEmailAddress` | `string` |  |
| `inviterFirstName` | `string` |  |
| `inviterLastName` | `string` |  |
| `isAuthorised` | `boolean` | Will be set to true once the invite has met the authorisation requirements. |
| `isInviteeRegistered` | `boolean` | If true, indicates the invitee's email address corresponds to an existing MoneyMoov user. |
| `lastInvited` | `string` |  |
| `merchantID` | `string` | ID of the merchant the user is being invited to. |
| `merchantName` | `string` |  |
| `message` | `string` |  |
| `registrationUrl` | `string` |  |
| `sendInviteEmail` | `boolean` | If set to true an email will be sent to the invitee with instructions on how to accept the invite. |
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
  id = "example_id", -- string
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
| `accountName` | `string` | Name for the account |
| `accountSupplierName` | `string` | The payment account supplier name. |
| `availableBalance` | `number` | The current available balance of the account. |
| `availableBalanceMinorUnits` | `number` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `number` | Balance of the account. |
| `balanceMinorUnits` | `number` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `string` | The bank name for external accounts |
| `consentID` | `string` | The ID of the consent used to connect the external account. |
| `createdBy` | `table` |  |
| `createdByDisplayName` | `string` | Either the name of the user, merchant token or api key that created the account |
| `currency` | `string` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `string` | Indicates the default payment rail for this account. |
| `displayName` | `string` | Gets a unique display name for the payment account. |
| `expiryDate` | `string` | The date that the external account will expire |
| `externalAccountIcon` | `string` | The Icon for external accounts |
| `id` | `string` | Unique id for the account. |
| `identifier` | `table` |  |
| `inserted` | `string` | Timestamp when the account was created. |
| `isArchived` | `boolean` | Indicates whether the account is archived. |
| `isConnectedAccount` | `boolean` | Indicates if the payment account is an externally connected account. |
| `isDefault` | `boolean` | Is the default account |
| `isTrustAccount` | `boolean` | Indicates if the payment account is a trust account. |
| `isVirtual` | `boolean` | True if the account is a virtual account. |
| `lastTransaction` | `table` |  |
| `lastUpdated` | `string` | Timestamp when the account was last updated. |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `merchantName` | `string` | The name of the merchant that owns the account. |
| `name` | `string` | The name of the virtual account. |
| `physicalAccountID` | `string` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `table` | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `number` | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `number` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `string` | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `string` | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `string` | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `string` |  |
| `xeroBankFeedSyncLastFailedAt` | `string` |  |
| `xeroBankFeedSyncLastFailureReason` | `string` |  |
| `xeroBankFeedSyncStatus` | `string` |  |
| `xeroUnsynchronisedTransactionsCount` | `number` | Indicates the number of unsynchronised transactions with Xero |

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
| `destinationUrl` | `string` | The destination URL for the webhook. |
| `emailAddress` | `string` | The recipient email address(es) for notifications. |
| `failedNotificationEmailAddress` | `string` | The email address to which notifications about failed webhook deliveries will be sent. |
| `id` | `string` |  |
| `isActive` | `boolean` |  |
| `merchantID` | `string` | The ID of the merchant that the webhook is for. |
| `notificationMethod` | `string` | The type of notification that will be sent. |
| `resourceTypes` | `table` | The resource types that the webhook will be generated for. |
| `retry` | `boolean` |  |
| `secret` | `string` | The secret key required to authenticate webhook notifications. |
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
