# Nofrixion Python SDK



The Python SDK for the Nofrixion API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Account()` — each
carrying a small, uniform set of operations (`list`, `load`, `create`, `update`, `remove`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/nofrixion-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from nofrixion_sdk import NofrixionSDK

client = NofrixionSDK({
    "apikey": os.environ.get("NOFRIXION_APIKEY"),
})
```

### 2. List account records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    accounts = client.Account().list()
    for account in accounts:
        print(account)
except Exception as err:
    print(f"list failed: {err}")
```

### 3. Load a cardcustomertoken

CardCustomerToken is nested under customer_email_address, so provide the `customer_email_address`.
`load()` returns the ENTITY — call data_get() for the record — and raises on error.

```python
try:
    cardcustomertoken = client.CardCustomerToken().load({"customer_email_address": "example_customer_email_address"})
    print(cardcustomertoken)
except Exception as err:
    print(f"load failed: {err}")
```

### 4. Create, update, and remove

```python
# Create — returns the ENTITY (call data_get() for the record)
created = client.Account().create({"createdBy": {}, "identifier": {}})

# Update — the created record's id is a plain dict key
client.Account().update({"id": created.data_get()["id"], "account_id": "example_account_id", "amount": 1})

# Remove
client.Account().remove({"id": created.data_get()["id"]})
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    consents = client.Consent().list()
    print(consents)
except Exception as err:
    print(f"list failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = NofrixionSDK.test()

# Entity ops return the ENTITY and raises on error;
# call data_get() for the record.
consent = client.Consent().list()
# consent contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = NofrixionSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### NofrixionSDK

```python
from nofrixion_sdk import NofrixionSDK

client = NofrixionSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = NofrixionSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### NofrixionSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
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
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

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

Create an instance: `account = client.Account()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountBalances` | `list` |  |
| `accountID` | `str` |  |
| `accountIdentifications` | `list` |  |
| `accountName` | `str` |  |
| `accountNames` | `list` |  |
| `accountSupplierName` | `str` |  |
| `accountType` | `str` |  |
| `availableBalance` | `float` |  |
| `availableBalanceMinorUnits` | `int` |  |
| `balance` | `float` |  |
| `balanceMinorUnits` | `int` |  |
| `bankName` | `str` |  |
| `consentID` | `str` |  |
| `consolidatedAccountInformation` | `dict` |  |
| `createdBy` | `dict` |  |
| `createdByDisplayName` | `str` |  |
| `currency` | `str` |  |
| `defaultPaymentRail` | `str` |  |
| `description` | `str` |  |
| `details` | `str` |  |
| `displayName` | `str` |  |
| `expiryDate` | `str` |  |
| `externalAccountIcon` | `str` |  |
| `format` | `str` |  |
| `fromDate` | `str` |  |
| `id` | `str` |  |
| `identifier` | `dict` |  |
| `inserted` | `str` |  |
| `isArchived` | `bool` |  |
| `isConnectedAccount` | `bool` |  |
| `isDefault` | `bool` |  |
| `isTrustAccount` | `bool` |  |
| `isVirtual` | `bool` |  |
| `lastTransaction` | `dict` |  |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` |  |
| `merchantName` | `str` |  |
| `nickname` | `str` |  |
| `physicalAccountID` | `str` |  |
| `roleIDs` | `list` |  |
| `rules` | `list` |  |
| `submittedPayoutsBalance` | `float` |  |
| `submittedPayoutsBalanceMinorUnits` | `int` |  |
| `summary` | `str` |  |
| `supplierPhysicalAccountID` | `str` |  |
| `supplierSepaInstantStatus` | `str` |  |
| `toDate` | `str` |  |
| `type` | `str` |  |
| `usageType` | `str` |  |
| `xeroBankFeedConnectionStatus` | `str` |  |
| `xeroBankFeedLastSyncedAt` | `str` |  |
| `xeroBankFeedSyncLastFailedAt` | `str` |  |
| `xeroBankFeedSyncLastFailureReason` | `str` |  |
| `xeroBankFeedSyncStatus` | `str` |  |
| `xeroUnsynchronisedTransactionsCount` | `int` |  |

#### Example: Load

```python
account = client.Account().load({"id": "account_id"})
```

#### Example: List

```python
accounts = client.Account().list()
```

#### Example: Create

```python
account = client.Account().create({
    "createdBy": {},  # dict
    "identifier": {},  # dict
})
```


### Batch

Create an instance: `batch = client.Batch()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approveUrl` | `str` |  |
| `id` | `str` |  |
| `payouts` | `list` |  |

#### Example: Load

```python
batch = client.Batch().load({"id": "batch_id"})
```

#### Example: Create

```python
batch = client.Batch().create({
})
```


### Beneficiary

Create an instance: `beneficiary = client.Beneficiary()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approvalCallbackUrl` | `str` |  |
| `authenticationMethods` | `list` |  |
| `authorisations` | `list` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `beneficiaries` | `list` |  |
| `beneficiaryEvents` | `list` |  |
| `canAuthorise` | `bool` |  |
| `canUpdate` | `bool` |  |
| `createdBy` | `dict` |  |
| `createdByEmailAddress` | `str` |  |
| `currency` | `str` |  |
| `destination` | `dict` |  |
| `failedBeneficiaries` | `dict` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `isEnabled` | `bool` |  |
| `lastAuthorised` | `str` |  |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` |  |
| `name` | `str` |  |
| `nonce` | `str` |  |
| `sourceAccountIDs` | `list` |  |
| `sourceAccounts` | `list` |  |
| `theirReference` | `str` |  |

#### Example: Load

```python
beneficiary = client.Beneficiary().load({"id": "beneficiary_id"})
```

#### Example: List

```python
beneficiarys = client.Beneficiary().list()
```

#### Example: Create

```python
beneficiary = client.Beneficiary().create({
    "createdBy": {},  # dict
    "currency": "example_currency",  # str
    "name": "example_name",  # str
})
```


### BeneficiaryGroup

Create an instance: `beneficiary_group = client.BeneficiaryGroup()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `groupMembers` | `list` |  |
| `groupName` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` |  |

#### Example: List

```python
beneficiary_groups = client.BeneficiaryGroup().list({"merchant_id": "example"})
```


### Card

Create an instance: `card = client.Card()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorizedAmount` | `str` |  |
| `currencyCode` | `str` |  |
| `isPayerAuthenticationRequired` | `bool` |  |
| `isSoftDecline` | `bool` |  |
| `payerAuthenticationAccessToken` | `str` |  |
| `payerAuthenticationMerchantData` | `str` |  |
| `payerAuthenticationUrl` | `str` |  |
| `payerAuthenticationWindowHeight` | `int` |  |
| `payerAuthenticationWindowWidth` | `int` |  |
| `paymentRequestCallbackUrl` | `str` |  |
| `paymentRequestID` | `str` |  |
| `requestID` | `str` |  |
| `responseCode` | `str` |  |
| `responseType` | `str` |  |
| `status` | `str` |  |
| `threeDSRedirectUrl` | `str` |  |
| `transactionID` | `str` |  |

#### Example: Create

```python
card = client.Card().create({
    "paymentrequest_id": "example_paymentrequest_id",  # str
})
```


### CardCustomerToken

Create an instance: `card_customer_token = client.CardCustomerToken()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `cardType` | `str` |  |
| `customerEmailAddress` | `str` |  |
| `expiryMonth` | `str` |  |
| `expiryYear` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `lastFourDigits` | `str` |  |
| `lastUpdated` | `str` |  |
| `maskedCardNumber` | `str` |  |
| `merchantID` | `str` |  |
| `paymentRequestID` | `str` |  |

#### Example: Load

```python
card_customer_token = client.CardCustomerToken().load({"customer_email_address": "customer_email_address"})
```

#### Example: List

```python
card_customer_tokens = client.CardCustomerToken().list({"customer_email_address": "example", "merchant_id": "example"})
```


### CardPayment

Create an instance: `card_payment = client.CardPayment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorizedAmount` | `str` |  |
| `currencyCode` | `str` |  |
| `isPayerAuthenticationRequired` | `bool` |  |
| `isSoftDecline` | `bool` |  |
| `payerAuthenticationAccessToken` | `str` |  |
| `payerAuthenticationMerchantData` | `str` |  |
| `payerAuthenticationUrl` | `str` |  |
| `payerAuthenticationWindowHeight` | `int` |  |
| `payerAuthenticationWindowWidth` | `int` |  |
| `paymentRequestCallbackUrl` | `str` |  |
| `paymentRequestID` | `str` |  |
| `requestID` | `str` |  |
| `responseCode` | `str` |  |
| `responseType` | `str` |  |
| `status` | `str` |  |
| `threeDSRedirectUrl` | `str` |  |
| `transactionID` | `str` |  |

#### Example: Create

```python
card_payment = client.CardPayment().create({
    "paymentrequest_id": "example_paymentrequest_id",  # str
})
```


### CardPublicKey

Create an instance: `card_public_key = client.CardPublicKey()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `jwt` | `str` |  |

#### Example: Load

```python
card_public_key = client.CardPublicKey().load({"paymentrequest_id": "paymentrequest_id"})
```


### Consent

Create an instance: `consent = client.Consent()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisationUrl` | `str` |  |
| `callbackUrl` | `str` |  |
| `consentID` | `str` |  |
| `emailAddress` | `str` |  |
| `expiryDate` | `str` |  |
| `failureCallbackUrl` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `institutionID` | `str` |  |
| `isConnectedAccounts` | `bool` |  |
| `isEnabled` | `bool` |  |
| `merchantID` | `str` |  |
| `provider` | `str` |  |
| `successWebHookUrl` | `str` |  |

#### Example: Load

```python
consent = client.Consent().load({"id": "consent_id"})
```

#### Example: List

```python
consents = client.Consent().list({"email": "example", "merchant_id": "example"})
```

#### Example: Create

```python
consent = client.Consent().create({
})
```


### Currency

Create an instance: `currency = client.Currency()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `code` | `str` |  |
| `decimals` | `int` |  |
| `isFiat` | `bool` |  |
| `iso4217AlphaCode` | `str` |  |
| `iso4217NumericCode` | `str` |  |
| `symbol` | `str` |  |

#### Example: List

```python
currencys = client.Currency().list()
```


### DirectDebitBatchSubmit

Create an instance: `direct_debit_batch_submit = client.DirectDebitBatchSubmit()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failedSubmissions` | `dict` |  |
| `successfulSubmissions` | `list` |  |

#### Example: Create

```python
direct_debit_batch_submit = client.DirectDebitBatchSubmit().create({
})
```


### FxRate

Create an instance: `fx_rate = client.FxRate()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `destinationCurrency` | `str` |  |
| `exchangeRate` | `float` |  |
| `expiryTime` | `str` |  |
| `quoteID` | `str` |  |
| `sourceCurrency` | `str` |  |

#### Example: Load

```python
fx_rate = client.FxRate().load({"destination": "destination", "source": "source", "valid_for_minute": 1})
```

#### Example: List

```python
fx_rates = client.FxRate().list({"destination": "example", "source": "example"})
```


### IPayment

Create an instance: `i_payment = client.IPayment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `paymentRequestID` | `str` |  |
| `responseType` | `str` |  |

#### Example: Create

```python
i_payment = client.IPayment().create({
})
```


### Mandate

Create an instance: `mandate = client.Mandate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountNumber` | `str` |  |
| `addressLine1` | `str` |  |
| `addressLine2` | `str` |  |
| `approvedAt` | `str` |  |
| `city` | `str` |  |
| `countryCode` | `str` |  |
| `currency` | `str` |  |
| `customerAccountNumber` | `str` |  |
| `customerCity` | `str` |  |
| `customerCountryCode` | `str` |  |
| `customerCountryName` | `str` |  |
| `customerEmailAddress` | `str` |  |
| `customerFirstName` | `str` |  |
| `customerIban` | `str` |  |
| `customerLastName` | `str` |  |
| `customerSortCode` | `str` |  |
| `emailAddress` | `str` |  |
| `firstName` | `str` |  |
| `iban` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `isRecurring` | `bool` |  |
| `lastName` | `str` |  |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` |  |
| `postalCode` | `str` |  |
| `reference` | `str` |  |
| `sortCode` | `str` |  |
| `status` | `str` |  |
| `supplierBankAccountID` | `str` |  |
| `supplierCustomerID` | `str` |  |
| `supplierMandateID` | `str` |  |
| `supplierName` | `str` |  |
| `supplierStatus` | `str` |  |

#### Example: Load

```python
mandate = client.Mandate().load({"id": "mandate_id"})
```

#### Example: Create

```python
mandate = client.Mandate().create({
    "addressLine1": "example_addressLine1",  # str
    "city": "example_city",  # str
    "countryCode": "example_countryCode",  # str
    "emailAddress": "example_emailAddress",  # str
    "firstName": "example_firstName",  # str
    "lastName": "example_lastName",  # str
    "postalCode": "example_postalCode",  # str
})
```


### Merchant

Create an instance: `merchant = client.Merchant()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountCurrencies` | `list` |  |
| `canHaveTrustAccounts` | `bool` |  |
| `cardPaymentProcessor` | `str` |  |
| `companyID` | `str` |  |
| `displayQrOnHostedPay` | `bool` |  |
| `hostedPayVersion` | `int` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `isBlocked` | `bool` |  |
| `isExited` | `bool` |  |
| `isSuspended` | `bool` |  |
| `jurisdiction` | `str` |  |
| `logoUrlPng` | `str` |  |
| `logoUrlSvg` | `str` |  |
| `merchantCategoryCode` | `str` |  |
| `name` | `str` |  |
| `notes` | `str` |  |
| `parentMerchant` | `dict` |  |
| `paymentAccountLimit` | `int` |  |
| `paymentAccounts` | `list` |  |
| `reason` | `str` |  |
| `shortName` | `str` |  |
| `supportedPaymentMethodsList` | `list` |  |
| `suspensionReason` | `str` |  |
| `tags` | `list` |  |
| `timeZoneId` | `str` |  |
| `tradingName` | `str` |  |
| `webHookLimit` | `int` |  |
| `yourRoleName` | `str` |  |

#### Example: Load

```python
merchant = client.Merchant().load({"id": "merchant_id"})
```

#### Example: List

```python
merchants = client.Merchant().list()
```


### MerchantAuthorisationSetting

Create an instance: `merchant_authorisation_setting = client.MerchantAuthorisationSetting()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amountLower` | `float` |  |
| `amountUpper` | `float` |  |
| `authorisationType` | `str` |  |
| `beneficiariesOnly` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `lastEditorCantAuthorise` | `bool` |  |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` |  |
| `numberOfAuthorisers` | `int` |  |
| `roleSettings` | `list` |  |

#### Example: List

```python
merchant_authorisation_settings = client.MerchantAuthorisationSetting().list({"merchant_id": "example"})
```


### MerchantDirectDebitMandatePage

Create an instance: `merchant_direct_debit_mandate_page = client.MerchantDirectDebitMandatePage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approvedAt` | `str` |  |
| `currency` | `str` |  |
| `customerAccountNumber` | `str` |  |
| `customerCity` | `str` |  |
| `customerCountryCode` | `str` |  |
| `customerCountryName` | `str` |  |
| `customerEmailAddress` | `str` |  |
| `customerFirstName` | `str` |  |
| `customerIban` | `str` |  |
| `customerLastName` | `str` |  |
| `customerSortCode` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `isRecurring` | `bool` |  |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` |  |
| `reference` | `str` |  |
| `status` | `str` |  |
| `supplierBankAccountID` | `str` |  |
| `supplierCustomerID` | `str` |  |
| `supplierMandateID` | `str` |  |
| `supplierName` | `str` |  |
| `supplierStatus` | `str` |  |

#### Example: List

```python
merchant_direct_debit_mandate_pages = client.MerchantDirectDebitMandatePage().list()
```


### MerchantPayByBankSetting

Create an instance: `merchant_pay_by_bank_setting = client.MerchantPayByBankSetting()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bankCountryCodes` | `list` |  |
| `bankID` | `str` |  |
| `bankName` | `str` |  |
| `businessInstitutionID` | `str` |  |
| `currency` | `str` |  |
| `logo` | `str` |  |
| `message` | `str` |  |
| `messageImageUrl` | `str` |  |
| `order` | `int` |  |
| `personalInstitutionID` | `str` |  |
| `processor` | `str` |  |
| `warningHeading` | `str` |  |
| `warningMessage` | `str` |  |

#### Example: List

```python
merchant_pay_by_bank_settings = client.MerchantPayByBankSetting().list({"merchant_id": "example"})
```


### MerchantPaymentRequestTemplate

Create an instance: `merchant_payment_request_template = client.MerchantPaymentRequestTemplate()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bankPaymentOptions` | `dict` |  |
| `cardPaymentAddressOptions` | `dict` |  |
| `cardPaymentCaptureOptions` | `dict` |  |
| `customFields` | `list` |  |
| `defaultFields` | `list` |  |
| `description` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` |  |
| `name` | `str` |  |
| `notificationOptions` | `dict` |  |
| `paymentMethods` | `dict` |  |
| `paymentTerms` | `dict` |  |
| `priorityBankOptions` | `dict` |  |
| `template` | `dict` |  |

#### Example: Load

```python
merchant_payment_request_template = client.MerchantPaymentRequestTemplate().load({"id": "merchant_payment_request_template_id", "paymentrequest_id": "paymentrequest_id"})
```

#### Example: List

```python
merchant_payment_request_templates = client.MerchantPaymentRequestTemplate().list({"merchant_id": "example"})
```


### MerchantToken

Create an instance: `merchant_token = client.MerchantToken()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authenticationMethods` | `list` |  |
| `authorisations` | `list` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `canAuthorise` | `bool` |  |
| `description` | `str` |  |
| `expiresAt` | `str` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `hmacAlgorithm` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `ipAddressWhitelist` | `str` |  |
| `isArchived` | `bool` |  |
| `isEnabled` | `bool` |  |
| `lastAuthorised` | `str` |  |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` |  |
| `nonce` | `str` |  |
| `permissionTypes` | `list` |  |
| `requestSignatureVersion` | `int` |  |
| `sharedSecretAlgorithm` | `str` |  |
| `sharedSecretBase64` | `str` |  |
| `token` | `str` |  |

#### Example: Load

```python
merchant_token = client.MerchantToken().load({"id": "merchant_token_id"})
```

#### Example: List

```python
merchant_tokens = client.MerchantToken().list({"merchant_id": "example"})
```

#### Example: Create

```python
merchant_token = client.MerchantToken().create({
    "nonce": "example_nonce",  # str
})
```


### Metadata

Create an instance: `metadata = client.Metadata()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```python
metadata = client.Metadata().load()
```


### NoFrixionVersion

Create an instance: `no_frixion_version = client.NoFrixionVersion()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `buildVersion` | `int` |  |
| `majorVersion` | `int` |  |
| `minorVersion` | `int` |  |
| `releaseName` | `str` |  |

#### Example: Load

```python
no_frixion_version = client.NoFrixionVersion().load()
```


### OpenBanking

Create an instance: `open_banking = client.OpenBanking()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```python
open_banking = client.OpenBanking().create({
    "account_id": "example_account_id",  # str
})
```


### Payeeverification

Create an instance: `payeeverification = client.Payeeverification()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountName` | `str` |  |
| `accountNumber` | `str` |  |
| `iban` | `str` |  |
| `payeeVerifiedAccountName` | `str` |  |
| `result` | `str` |  |
| `secondaryIdentification` | `str` |  |
| `sortCode` | `str` |  |

#### Example: Create

```python
payeeverification = client.Payeeverification().create({
    "accountName": "example_accountName",  # str
    "iban": "example_iban",  # str
})
```


### Payment

Create an instance: `payment = client.Payment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addresses` | `list` |  |
| `amount` | `float` |  |
| `amountPending` | `float` |  |
| `amountReceived` | `float` |  |
| `amountRefunded` | `float` |  |
| `autoSendReceipt` | `bool` |  |
| `baseOriginUrl` | `str` |  |
| `callbackUrl` | `str` |  |
| `cardAuthorizeOnly` | `bool` |  |
| `cardCreateToken` | `bool` |  |
| `cardCreateTokenMode` | `str` |  |
| `cardIgnoreCVN` | `bool` |  |
| `cardNoPayerAuthentication` | `bool` |  |
| `cardProcessorMerchantID` | `str` |  |
| `cardStripePaymentIntentID` | `str` |  |
| `cardStripePaymentIntentSecret` | `str` |  |
| `cardTransmitRawDetails` | `bool` |  |
| `createdByUser` | `dict` |  |
| `currency` | `str` |  |
| `customFields` | `list` |  |
| `customerEmailAddress` | `str` |  |
| `customerID` | `str` |  |
| `customerName` | `str` |  |
| `description` | `str` |  |
| `destinationAccount` | `dict` |  |
| `directDebitPayment` | `dict` |  |
| `dueDate` | `str` |  |
| `events` | `list` |  |
| `failureCallbackUrl` | `str` |  |
| `fieldDisplaySettings` | `list` |  |
| `formattedAmount` | `str` |  |
| `hostedPayCheckoutUrl` | `str` |  |
| `id` | `str` |  |
| `ignoreAddressVerification` | `bool` |  |
| `inserted` | `str` |  |
| `insertedSortable` | `str` |  |
| `isArchived` | `bool` |  |
| `jwk` | `str` |  |
| `lastUpdated` | `str` |  |
| `lightningInvoice` | `str` |  |
| `lightningInvoiceExpiresAt` | `str` |  |
| `merchantDirectDebitMandateID` | `str` |  |
| `merchantID` | `str` |  |
| `merchantTokenDescription` | `str` |  |
| `notificationEmailAddresses` | `str` |  |
| `notificationRoleIDs` | `list` |  |
| `orderID` | `str` |  |
| `partialPaymentMethod` | `str` |  |
| `partialPaymentSteps` | `str` |  |
| `paymentAttempts` | `list` |  |
| `paymentMethods` | `list` |  |
| `paymentProcessor` | `str` |  |
| `payrunID` | `str` |  |
| `pispAccountID` | `str` |  |
| `priorityBankID` | `str` |  |
| `result` | `dict` |  |
| `sandboxSettleDelayInSeconds` | `int` |  |
| `shippingAddress` | `dict` |  |
| `shippingAddressCity` | `str` |  |
| `shippingAddressCountryCode` | `str` |  |
| `shippingAddressCounty` | `str` |  |
| `shippingAddressLine1` | `str` |  |
| `shippingAddressLine2` | `str` |  |
| `shippingAddressPostCode` | `str` |  |
| `shippingEmail` | `str` |  |
| `shippingFirstName` | `str` |  |
| `shippingLastName` | `str` |  |
| `shippingPhone` | `str` |  |
| `status` | `str` |  |
| `successWebHookUrl` | `str` |  |
| `tagIds` | `list` |  |
| `tags` | `list` |  |
| `title` | `str` |  |
| `tokenisedCards` | `list` |  |
| `transactions` | `list` |  |
| `useHostedPaymentPage` | `bool` |  |

#### Example: Load

```python
payment = client.Payment().load({"id": "payment_id"})
```

#### Example: Create

```python
payment = client.Payment().create({
    "createdByUser": {},  # dict
})
```


### PaymentAccount

Create an instance: `payment_account = client.PaymentAccount()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountName` | `str` |  |
| `accountSupplierName` | `str` |  |
| `availableBalance` | `float` |  |
| `availableBalanceMinorUnits` | `int` |  |
| `balance` | `float` |  |
| `balanceMinorUnits` | `int` |  |
| `bankName` | `str` |  |
| `consentID` | `str` |  |
| `createdBy` | `dict` |  |
| `createdByDisplayName` | `str` |  |
| `currency` | `str` |  |
| `defaultPaymentRail` | `str` |  |
| `displayName` | `str` |  |
| `expiryDate` | `str` |  |
| `externalAccountIcon` | `str` |  |
| `id` | `str` |  |
| `identifier` | `dict` |  |
| `inserted` | `str` |  |
| `isArchived` | `bool` |  |
| `isConnectedAccount` | `bool` |  |
| `isDefault` | `bool` |  |
| `isTrustAccount` | `bool` |  |
| `isVirtual` | `bool` |  |
| `lastTransaction` | `dict` |  |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` |  |
| `merchantName` | `str` |  |
| `physicalAccountID` | `str` |  |
| `rules` | `list` |  |
| `submittedPayoutsBalance` | `float` |  |
| `submittedPayoutsBalanceMinorUnits` | `int` |  |
| `summary` | `str` |  |
| `supplierSepaInstantStatus` | `str` |  |
| `xeroBankFeedConnectionStatus` | `str` |  |
| `xeroBankFeedLastSyncedAt` | `str` |  |
| `xeroBankFeedSyncLastFailedAt` | `str` |  |
| `xeroBankFeedSyncLastFailureReason` | `str` |  |
| `xeroBankFeedSyncStatus` | `str` |  |
| `xeroUnsynchronisedTransactionsCount` | `int` |  |

#### Example: List

```python
payment_accounts = client.PaymentAccount().list()
```


### PaymentAccountMinimal

Create an instance: `payment_account_minimal = client.PaymentAccountMinimal()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountName` | `str` |  |
| `availableBalance` | `float` |  |
| `balance` | `float` |  |
| `balanceMinorUnits` | `int` |  |
| `currency` | `str` |  |
| `id` | `str` |  |
| `identifier` | `dict` |  |
| `isArchived` | `bool` |  |
| `isConnectedAccount` | `bool` |  |
| `merchantID` | `str` |  |
| `submittedPayoutsBalance` | `float` |  |

#### Example: List

```python
payment_account_minimals = client.PaymentAccountMinimal().list()
```


### PaymentInitiation

Create an instance: `payment_initiation = client.PaymentInitiation()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `paymentInitiationID` | `str` |  |
| `paymentRequestCallbackUrl` | `str` |  |
| `paymentRequestID` | `str` |  |
| `redirectUrl` | `str` |  |
| `responseType` | `str` |  |
| `specificErrorMessage` | `str` |  |

#### Example: Create

```python
payment_initiation = client.PaymentInitiation().create({
    "paymentrequest_id": "example_paymentrequest_id",  # str
})
```


### PaymentRequest

Create an instance: `payment_request = client.PaymentRequest()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addresses` | `list` |  |
| `amount` | `float` |  |
| `amountPending` | `float` |  |
| `amountReceived` | `float` |  |
| `amountRefunded` | `float` |  |
| `autoSendReceipt` | `bool` |  |
| `baseOriginUrl` | `str` |  |
| `callbackUrl` | `str` |  |
| `cardAuthorizeOnly` | `bool` |  |
| `cardCreateToken` | `bool` |  |
| `cardCreateTokenMode` | `str` |  |
| `cardIgnoreCVN` | `bool` |  |
| `cardProcessorMerchantID` | `str` |  |
| `cardStripePaymentIntentID` | `str` |  |
| `cardStripePaymentIntentSecret` | `str` |  |
| `createdByUser` | `dict` |  |
| `currency` | `str` |  |
| `customFields` | `list` |  |
| `customerEmailAddress` | `str` |  |
| `customerID` | `str` |  |
| `customerName` | `str` |  |
| `description` | `str` |  |
| `destinationAccount` | `dict` |  |
| `directDebitPayment` | `dict` |  |
| `doSimulateSettlementFailure` | `bool` |  |
| `dueDate` | `str` |  |
| `errorDescription` | `str` |  |
| `events` | `list` |  |
| `failedPaymentRequests` | `dict` |  |
| `failureCallbackUrl` | `str` |  |
| `fieldDisplaySettings` | `list` |  |
| `formattedAmount` | `str` |  |
| `hostedPayCheckoutUrl` | `str` |  |
| `id` | `str` |  |
| `ignoreAddressVerification` | `bool` |  |
| `inserted` | `str` |  |
| `insertedSortable` | `str` |  |
| `institution` | `str` |  |
| `isArchived` | `bool` |  |
| `jwk` | `str` |  |
| `lastUpdated` | `str` |  |
| `lightningInvoice` | `str` |  |
| `lightningInvoiceExpiresAt` | `str` |  |
| `merchantDirectDebitMandateID` | `str` |  |
| `merchantID` | `str` |  |
| `merchantTokenDescription` | `str` |  |
| `notificationEmailAddresses` | `str` |  |
| `notificationRoleIDs` | `list` |  |
| `orderID` | `str` |  |
| `partialPaymentMethod` | `str` |  |
| `partialPaymentSteps` | `str` |  |
| `paymentAttempts` | `list` |  |
| `paymentInitiationID` | `str` |  |
| `paymentMethods` | `list` |  |
| `paymentProcessor` | `str` |  |
| `paymentRequests` | `list` |  |
| `payrunID` | `str` |  |
| `pispAccountID` | `str` |  |
| `priorityBankID` | `str` |  |
| `result` | `dict` |  |
| `sandboxSettleDelayInSeconds` | `int` |  |
| `shippingAddress` | `dict` |  |
| `status` | `str` |  |
| `successWebHookUrl` | `str` |  |
| `tags` | `list` |  |
| `title` | `str` |  |
| `tokenisedCards` | `list` |  |
| `transactions` | `list` |  |
| `useHostedPaymentPage` | `bool` |  |

#### Example: Load

```python
payment_request = client.PaymentRequest().load()
```

#### Example: List

```python
payment_requests = client.PaymentRequest().list()
```

#### Example: Create

```python
payment_request = client.PaymentRequest().create({
    "createdByUser": {},  # dict
})
```


### PaymentRequestEvent

Create an instance: `payment_request_event = client.PaymentRequestEvent()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float` |  |
| `applePayTransactionID` | `str` |  |
| `cardAuthorizationResponseID` | `str` |  |
| `cardExpiryMonth` | `int` |  |
| `cardExpiryYear` | `int` |  |
| `cardIssuer` | `str` |  |
| `cardIssuerCountry` | `str` |  |
| `cardLastFourDigits` | `str` |  |
| `cardRequestID` | `str` |  |
| `cardScheme` | `str` |  |
| `cardTokenCustomerID` | `str` |  |
| `cardTransactionID` | `str` |  |
| `currency` | `str` |  |
| `directDebitPaymentID` | `str` |  |
| `directDebitPaymentReference` | `str` |  |
| `drirectDebitMandateID` | `str` |  |
| `errorMessage` | `str` |  |
| `errorReason` | `str` |  |
| `eventType` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `lightningInvoice` | `str` |  |
| `lightningRHash` | `str` |  |
| `originUrl` | `str` |  |
| `paymentMethodType` | `str` |  |
| `paymentProcessorName` | `str` |  |
| `paymentRequestID` | `str` |  |
| `pispBankStatus` | `str` |  |
| `pispPaymentInitiationID` | `str` |  |
| `pispPaymentInstitutionName` | `str` |  |
| `pispPaymentServiceProviderID` | `str` |  |
| `pispRedirectUrl` | `str` |  |
| `reconciledTransactionID` | `str` |  |
| `refundPayoutID` | `str` |  |
| `status` | `str` |  |
| `walletName` | `str` |  |

#### Example: List

```python
payment_request_events = client.PaymentRequestEvent().list({"paymentrequest_id": "example"})
```


### PaymentRequestMetric

Create an instance: `payment_request_metric = client.PaymentRequestMetric()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```python
payment_request_metric = client.PaymentRequestMetric().load()
```


### PaymentRequestMinimal

Create an instance: `payment_request_minimal = client.PaymentRequestMinimal()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float` |  |
| `amountPending` | `float` |  |
| `amountReceived` | `float` |  |
| `amountRefunded` | `float` |  |
| `callbackUrl` | `str` |  |
| `cardStripePaymentIntentSecret` | `str` |  |
| `countryCode` | `str` |  |
| `currency` | `str` |  |
| `customFieldsToDisplay` | `list` |  |
| `description` | `str` |  |
| `dueDate` | `str` |  |
| `fieldDisplaySettings` | `list` |  |
| `googlePayMerchantID` | `str` |  |
| `id` | `str` |  |
| `jwk` | `str` |  |
| `merchantID` | `str` |  |
| `merchantLogoUrlPng` | `str` |  |
| `merchantLogoUrlSvg` | `str` |  |
| `merchantName` | `str` |  |
| `merchantShortName` | `str` |  |
| `partialPaymentMethod` | `str` |  |
| `paymentAttempts` | `list` |  |
| `paymentMethodsList` | `list` |  |
| `paymentProcessor` | `str` |  |
| `paymentProcessorKey` | `str` |  |
| `pispError` | `str` |  |
| `priorityBankID` | `str` |  |
| `status` | `str` |  |
| `stripeAccountID` | `str` |  |
| `title` | `str` |  |

#### Example: List

```python
payment_request_minimals = client.PaymentRequestMinimal().list({"paymentrequest_id": "example"})
```


### PaymentRequestResult

Create an instance: `payment_request_result = client.PaymentRequestResult()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float` |  |
| `amountPending` | `float` |  |
| `amountReceived` | `float` |  |
| `amountRefunded` | `float` |  |
| `currency` | `str` |  |
| `customerID` | `str` |  |
| `paymentRequestID` | `str` |  |
| `payments` | `list` |  |
| `pispAuthorizations` | `list` |  |
| `requestedAmount` | `float` |  |
| `result` | `str` |  |

#### Example: List

```python
payment_request_results = client.PaymentRequestResult().list({"paymentrequest_id": "example"})
```


### Payout

Create an instance: `payout = client.Payout()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountID` | `str` |  |
| `allowIncomplete` | `bool` |  |
| `amount` | `float` |  |
| `amountMinorUnits` | `int` |  |
| `approvePayoutUrl` | `str` |  |
| `approverID` | `str` |  |
| `authenticationMethods` | `list` |  |
| `authorisations` | `list` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `batchPayoutID` | `str` |  |
| `beneficiary` | `dict` |  |
| `beneficiaryID` | `str` |  |
| `canAuthorise` | `bool` |  |
| `canProcess` | `bool` |  |
| `canUpdate` | `bool` |  |
| `chargeBearer` | `str` |  |
| `createdBy` | `str` |  |
| `createdByEmailAddress` | `str` |  |
| `currency` | `str` |  |
| `currentUserID` | `str` |  |
| `description` | `str` |  |
| `destination` | `dict` |  |
| `documents` | `list` |  |
| `events` | `list` |  |
| `failedPayouts` | `dict` |  |
| `formattedAmount` | `str` |  |
| `formattedFxDestinationAmount` | `str` |  |
| `formattedSchedule` | `str` |  |
| `formattedScheduleDayOnly` | `str` |  |
| `formattedSourceAccountAvailableBalance` | `str` |  |
| `fxDestinationAmount` | `float` |  |
| `fxDestinationAmountMinorUnits` | `int` |  |
| `fxDestinationCurrency` | `str` |  |
| `fxQuoteExpiresAt` | `str` |  |
| `fxQuoteID` | `str` |  |
| `fxRate` | `float` |  |
| `fxUseDestinationAmount` | `bool` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `invoiceID` | `str` |  |
| `isArchived` | `bool` |  |
| `isFailed` | `bool` |  |
| `isSettled` | `bool` |  |
| `isSubmitted` | `bool` |  |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` |  |
| `merchantTokenDescription` | `str` |  |
| `nonce` | `str` |  |
| `paymentProcessor` | `str` |  |
| `paymentRail` | `str` |  |
| `payouts` | `list` |  |
| `payrunID` | `str` |  |
| `payrunName` | `str` |  |
| `reason` | `str` |  |
| `rule` | `dict` |  |
| `scheduleDate` | `str` |  |
| `scheduled` | `bool` |  |
| `sourceAccountAvailableBalance` | `float` |  |
| `sourceAccountAvailableBalanceMinorUnits` | `int` |  |
| `sourceAccountBic` | `str` |  |
| `sourceAccountCurrency` | `str` |  |
| `sourceAccountIban` | `str` |  |
| `sourceAccountIdentifier` | `dict` |  |
| `sourceAccountName` | `str` |  |
| `sourceAccountNumber` | `str` |  |
| `sourceAccountSortcode` | `str` |  |
| `status` | `str` |  |
| `tagIds` | `list` |  |
| `tags` | `list` |  |
| `theirReference` | `str` |  |
| `topupPayrunID` | `str` |  |
| `transactedAmount` | `float` |  |
| `transactedFxAmount` | `float` |  |
| `transactedFxRate` | `float` |  |
| `type` | `str` |  |
| `userID` | `str` |  |
| `yourReference` | `str` |  |

#### Example: Load

```python
payout = client.Payout().load({"id": "payout_id"})
```

#### Example: List

```python
payouts = client.Payout().list()
```

#### Example: Create

```python
payout = client.Payout().create({
    "beneficiary": {},  # dict
    "sourceAccountIdentifier": {},  # dict
})
```


### PayoutKeysetPage

Create an instance: `payout_keyset_page = client.PayoutKeysetPage()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountID` | `str` |  |
| `amount` | `float` |  |
| `amountMinorUnits` | `int` |  |
| `approvePayoutUrl` | `str` |  |
| `approverID` | `str` |  |
| `authenticationMethods` | `list` |  |
| `authorisations` | `list` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `batchPayoutID` | `str` |  |
| `beneficiary` | `dict` |  |
| `canAuthorise` | `bool` |  |
| `canProcess` | `bool` |  |
| `canUpdate` | `bool` |  |
| `chargeBearer` | `str` |  |
| `createdBy` | `str` |  |
| `createdByEmailAddress` | `str` |  |
| `currency` | `str` |  |
| `currentUserID` | `str` |  |
| `description` | `str` |  |
| `destination` | `dict` |  |
| `documents` | `list` |  |
| `events` | `list` |  |
| `formattedAmount` | `str` |  |
| `formattedFxDestinationAmount` | `str` |  |
| `formattedSchedule` | `str` |  |
| `formattedScheduleDayOnly` | `str` |  |
| `formattedSourceAccountAvailableBalance` | `str` |  |
| `fxDestinationAmount` | `float` |  |
| `fxDestinationAmountMinorUnits` | `int` |  |
| `fxDestinationCurrency` | `str` |  |
| `fxQuoteExpiresAt` | `str` |  |
| `fxQuoteID` | `str` |  |
| `fxRate` | `float` |  |
| `fxUseDestinationAmount` | `bool` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `invoiceID` | `str` |  |
| `isArchived` | `bool` |  |
| `isFailed` | `bool` |  |
| `isSettled` | `bool` |  |
| `isSubmitted` | `bool` |  |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` |  |
| `merchantTokenDescription` | `str` |  |
| `nonce` | `str` |  |
| `paymentProcessor` | `str` |  |
| `paymentRail` | `str` |  |
| `payrunID` | `str` |  |
| `payrunName` | `str` |  |
| `rule` | `dict` |  |
| `scheduleDate` | `str` |  |
| `scheduled` | `bool` |  |
| `sourceAccountAvailableBalance` | `float` |  |
| `sourceAccountAvailableBalanceMinorUnits` | `int` |  |
| `sourceAccountBic` | `str` |  |
| `sourceAccountCurrency` | `str` |  |
| `sourceAccountIban` | `str` |  |
| `sourceAccountIdentifier` | `dict` |  |
| `sourceAccountName` | `str` |  |
| `sourceAccountNumber` | `str` |  |
| `sourceAccountSortcode` | `str` |  |
| `status` | `str` |  |
| `tags` | `list` |  |
| `theirReference` | `str` |  |
| `topupPayrunID` | `str` |  |
| `transactedAmount` | `float` |  |
| `transactedFxAmount` | `float` |  |
| `transactedFxRate` | `float` |  |
| `type` | `str` |  |
| `userID` | `str` |  |
| `yourReference` | `str` |  |

#### Example: List

```python
payout_keyset_pages = client.PayoutKeysetPage().list()
```


### PayoutMetric

Create an instance: `payout_metric = client.PayoutMetric()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```python
payout_metric = client.PayoutMetric().load()
```


### Payrun

Create an instance: `payrun = client.Payrun()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisationDate` | `str` |  |
| `authorisations` | `list` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `batchPayoutID` | `str` |  |
| `canAuthorise` | `bool` |  |
| `canDelete` | `bool` |  |
| `canEdit` | `bool` |  |
| `events` | `list` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `invoices` | `list` |  |
| `invoicesMinimal` | `list` |  |
| `isArchived` | `bool` |  |
| `lastUpdated` | `str` |  |
| `lastUpdatedBy` | `dict` |  |
| `merchantID` | `str` |  |
| `name` | `str` |  |
| `nonce` | `str` |  |
| `notes` | `str` |  |
| `payments` | `list` |  |
| `payouts` | `list` |  |
| `payoutsCount` | `int` |  |
| `reason` | `str` |  |
| `scheduleDate` | `str` |  |
| `scheduledDate` | `str` |  |
| `sourceAccounts` | `list` |  |
| `status` | `str` |  |
| `totalEur` | `float` |  |
| `totalGbp` | `float` |  |
| `totalUsd` | `float` |  |

#### Example: Load

```python
payrun = client.Payrun().load({"id": "payrun_id"})
```

#### Example: List

```python
payruns = client.Payrun().list()
```

#### Example: Create

```python
payrun = client.Payrun().create({
    "id": "example_id",  # str
    "lastUpdatedBy": {},  # dict
})
```


### Report

Create an instance: `report = client.Report()`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |


### ReportResult

Create an instance: `report_result = client.ReportResult()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `contentType` | `str` |  |
| `contents` | `str` |  |
| `lastCompletedAt` | `str` |  |
| `merchantID` | `str` |  |
| `reportName` | `str` |  |
| `reportType` | `str` |  |
| `statementNumber` | `int` |  |

#### Example: Load

```python
report_result = client.ReportResult().load({"id": 1, "report_id": "report_id"})
```


### Role

Create an instance: `role = client.Role()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failedRoles` | `dict` |  |
| `roles` | `list` |  |

#### Example: Create

```python
role = client.Role().create({
    "merchant_id": "example_merchant_id",  # str
})
```


### Rule

Create an instance: `rule = client.Rule()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account` | `dict` |  |
| `accountID` | `str` |  |
| `approveUrl` | `str` |  |
| `approverID` | `str` |  |
| `authenticationMethods` | `list` |  |
| `authorisations` | `list` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `canAuthorise` | `bool` |  |
| `createdBy` | `dict` |  |
| `description` | `str` |  |
| `endAt` | `str` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `isDisabled` | `bool` |  |
| `lastExecutedAt` | `str` |  |
| `lastRunAtTransactionDate` | `str` |  |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` |  |
| `name` | `str` |  |
| `nonce` | `str` |  |
| `onApprovedWebHookUrl` | `str` |  |
| `onExecutionErrorWebHookUrl` | `str` |  |
| `onExecutionSuccessWebHookUrl` | `str` |  |
| `startAt` | `str` |  |
| `status` | `str` |  |
| `sweepAction` | `dict` |  |
| `timeZoneId` | `str` |  |
| `triggerCronExpression` | `str` |  |
| `triggerOnPayIn` | `bool` |  |
| `userID` | `str` |  |
| `webHookSecret` | `str` |  |

#### Example: Load

```python
rule = client.Rule().load({"id": "rule_id"})
```

#### Example: List

```python
rules = client.Rule().list()
```

#### Example: Create

```python
rule = client.Rule().create({
    "createdBy": {},  # dict
    "nonce": "example_nonce",  # str
})
```


### RuleEvent

Create an instance: `rule_event = client.RuleEvent()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errorMessage` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `isAuthoriseToEnable` | `bool` |  |
| `message` | `str` |  |
| `rawResponse` | `str` |  |
| `ruleEventType` | `str` |  |
| `ruleID` | `str` |  |
| `user` | `dict` |  |

#### Example: List

```python
rule_events = client.RuleEvent().list({"id": "example"})
```


### Tag

Create an instance: `tag = client.Tag()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `colourHex` | `str` |  |
| `description` | `str` |  |
| `id` | `str` |  |
| `merchantID` | `str` |  |
| `name` | `str` |  |

#### Example: List

```python
tags = client.Tag().list({"merchant_id": "example"})
```

#### Example: Create

```python
tag = client.Tag().create({
    "merchant_id": "example_merchant_id",  # str
    "merchantID": "example_merchantID",  # str
    "name": "example_name",  # str
})
```


### Token

Create an instance: `token = client.Token()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```python
token = client.Token().create({
    "id": "example_id",  # str
})
```


### Transaction

Create an instance: `transaction = client.Transaction()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountID` | `str` |  |
| `accountName` | `str` |  |
| `accountSequenceNumber` | `int` |  |
| `addressDetails` | `dict` |  |
| `amount` | `float` |  |
| `amountMinorUnits` | `int` |  |
| `balance` | `float` |  |
| `balanceMinorUnits` | `int` |  |
| `bookingDateTime` | `str` |  |
| `chargeDetails` | `dict` |  |
| `content` | `list` |  |
| `counterparty` | `dict` |  |
| `counterpartySummary` | `str` |  |
| `currency` | `str` |  |
| `currencyExchange` | `dict` |  |
| `date` | `str` |  |
| `description` | `str` |  |
| `enrichment` | `dict` |  |
| `fxAmount` | `float` |  |
| `fxCurrency` | `str` |  |
| `fxRate` | `float` |  |
| `grossAmount` | `dict` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `isoBankTransactionCode` | `dict` |  |
| `merchant` | `dict` |  |
| `merchantID` | `str` |  |
| `pageNumber` | `int` |  |
| `pageSize` | `int` |  |
| `payeeDetails` | `dict` |  |
| `payerDetails` | `dict` |  |
| `paymentRequestCustomFields` | `dict` |  |
| `paymentRequestID` | `str` |  |
| `payoutID` | `str` |  |
| `proprietaryBankTransactionCode` | `dict` |  |
| `rawReference` | `str` |  |
| `reference` | `str` |  |
| `ruleID` | `str` |  |
| `statementReferences` | `list` |  |
| `status` | `str` |  |
| `supplementaryData` | `Any` |  |
| `tags` | `list` |  |
| `theirReference` | `str` |  |
| `totalPages` | `int` |  |
| `totalSize` | `int` |  |
| `transactionAmount` | `dict` |  |
| `transactionDate` | `str` |  |
| `transactionInformation` | `list` |  |
| `transactionMutability` | `str` |  |
| `type` | `str` |  |
| `valueDateTime` | `str` |  |
| `virtualIBAN` | `str` |  |
| `yourReference` | `str` |  |

#### Example: Load

```python
transaction = client.Transaction().load({"id": "transaction_id"})
```

#### Example: List

```python
transactions = client.Transaction().list()
```

#### Example: Create

```python
transaction = client.Transaction().create({
    "id": "example_id",  # str
    "grossAmount": {},  # dict
    "payeeDetails": {},  # dict
    "payerDetails": {},  # dict
    "transactionAmount": {},  # dict
})
```


### User

Create an instance: `user = client.User()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `clientSessionTimeouts` | `list` |  |
| `emailAddress` | `str` |  |
| `firstName` | `str` |  |
| `id` | `str` |  |
| `lastName` | `str` |  |
| `passkeyAdded` | `bool` |  |
| `permissions` | `dict` |  |
| `profile` | `str` |  |
| `rolesWithScope` | `list` |  |
| `twoFactorEnabled` | `bool` |  |
| `userInviteID` | `str` |  |

#### Example: List

```python
users = client.User().list()
```


### UserInvite

Create an instance: `user_invite = client.UserInvite()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisationStatus` | `dict` |  |
| `failedUserInvites` | `dict` |  |
| `id` | `str` |  |
| `initialRoleID` | `str` |  |
| `inviteeEmailAddress` | `str` |  |
| `inviteeFirstName` | `str` |  |
| `inviteeLastName` | `str` |  |
| `inviterEmailAddress` | `str` |  |
| `inviterFirstName` | `str` |  |
| `inviterLastName` | `str` |  |
| `isAuthorised` | `bool` |  |
| `isInviteeRegistered` | `bool` |  |
| `lastInvited` | `str` |  |
| `merchantID` | `str` |  |
| `merchantName` | `str` |  |
| `message` | `str` |  |
| `registrationUrl` | `str` |  |
| `sendInviteEmail` | `bool` |  |
| `status` | `str` |  |
| `user` | `dict` |  |
| `userID` | `str` |  |
| `userInvites` | `list` |  |

#### Example: Load

```python
user_invite = client.UserInvite().load({"id": "user_invite_id"})
```

#### Example: List

```python
user_invites = client.UserInvite().list({"merchant_id": "example"})
```

#### Example: Create

```python
user_invite = client.UserInvite().create({
    "user": {},  # dict
})
```


### Virtual

Create an instance: `virtual = client.Virtual()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountName` | `str` |  |
| `accountSupplierName` | `str` |  |
| `availableBalance` | `float` |  |
| `availableBalanceMinorUnits` | `int` |  |
| `balance` | `float` |  |
| `balanceMinorUnits` | `int` |  |
| `bankName` | `str` |  |
| `consentID` | `str` |  |
| `createdBy` | `dict` |  |
| `createdByDisplayName` | `str` |  |
| `currency` | `str` |  |
| `defaultPaymentRail` | `str` |  |
| `displayName` | `str` |  |
| `expiryDate` | `str` |  |
| `externalAccountIcon` | `str` |  |
| `id` | `str` |  |
| `identifier` | `dict` |  |
| `inserted` | `str` |  |
| `isArchived` | `bool` |  |
| `isConnectedAccount` | `bool` |  |
| `isDefault` | `bool` |  |
| `isTrustAccount` | `bool` |  |
| `isVirtual` | `bool` |  |
| `lastTransaction` | `dict` |  |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` |  |
| `merchantName` | `str` |  |
| `name` | `str` |  |
| `physicalAccountID` | `str` |  |
| `rules` | `list` |  |
| `submittedPayoutsBalance` | `float` |  |
| `submittedPayoutsBalanceMinorUnits` | `int` |  |
| `summary` | `str` |  |
| `supplierSepaInstantStatus` | `str` |  |
| `xeroBankFeedConnectionStatus` | `str` |  |
| `xeroBankFeedLastSyncedAt` | `str` |  |
| `xeroBankFeedSyncLastFailedAt` | `str` |  |
| `xeroBankFeedSyncLastFailureReason` | `str` |  |
| `xeroBankFeedSyncStatus` | `str` |  |
| `xeroUnsynchronisedTransactionsCount` | `int` |  |

#### Example: Create

```python
virtual = client.Virtual().create({
    "account_id": "example_account_id",  # str
    "createdBy": {},  # dict
    "identifier": {},  # dict
    "name": "example_name",  # str
})
```


### Webhook

Create an instance: `webhook = client.Webhook()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `destinationUrl` | `str` |  |
| `emailAddress` | `str` |  |
| `failedNotificationEmailAddress` | `str` |  |
| `id` | `str` |  |
| `isActive` | `bool` |  |
| `merchantID` | `str` |  |
| `notificationMethod` | `str` |  |
| `resourceTypes` | `list` |  |
| `retry` | `bool` |  |
| `secret` | `str` |  |
| `version` | `int` |  |

#### Example: Load

```python
webhook = client.Webhook().load({"id": "webhook_id"})
```

#### Example: List

```python
webhooks = client.Webhook().list({"merchant_id": "example"})
```

#### Example: Create

```python
webhook = client.Webhook().create({
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

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── nofrixion_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`nofrixion_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```python
consent = client.Consent()
consent.list()

# consent.data_get() now returns the consent data from the last list
# consent.match_get() returns the last match criteria
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
