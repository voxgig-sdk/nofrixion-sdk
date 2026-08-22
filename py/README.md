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
created = client.Account().create({"account_id": "example_account_id", "currency": "example_currency", "createdBy": {}, "identifier": {}})

# Update — the created record's id is a plain dict key
client.Account().update({"id": created.data_get()["id"], "accountBalances": [], "accountID": "example_accountID"})

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
| `accountBalances` | `list` | The various balances for the account. |
| `accountID` | `str` | ID of the account. |
| `accountIdentifications` | `list` | The canoncial identifiers for the account. |
| `accountName` | `str` | Name for the account |
| `accountNames` | `list` | Optional account names set by the account holder. |
| `accountSupplierName` | `str` | The payment account supplier name. |
| `accountType` | `str` | The type of account e.g. |
| `availableBalance` | `float` | The current available balance of the account. |
| `availableBalanceMinorUnits` | `int` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `float` | Balance of the account. |
| `balanceMinorUnits` | `int` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `str` | The bank name for external accounts |
| `consentID` | `str` | The ID of the consent used to connect the external account. |
| `consolidatedAccountInformation` | `dict` | Summary information regarding account balances of the overall account provided by the bank. |
| `createdBy` | `dict` |  |
| `createdByDisplayName` | `str` | Either the name of the user, merchant token or api key that created the account |
| `currency` | `str` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `str` | Indicates the default payment rail for this account. |
| `description` | `str` | Product name as defined by the financial institution for this account. |
| `details` | `str` | Supplementary specifications that might be provided by the Bank. |
| `displayName` | `str` | Gets a unique display name for the payment account. |
| `expiryDate` | `str` | The date that the external account will expire |
| `externalAccountIcon` | `str` | The Icon for external accounts |
| `format` | `str` | File format to save the statement as. |
| `fromDate` | `str` | Minimum transaction date for the statement. |
| `id` | `str` | Unique id for the account. |
| `identifier` | `dict` |  |
| `inserted` | `str` | Timestamp when the account was created. |
| `isArchived` | `bool` | Indicates whether the account is archived. |
| `isConnectedAccount` | `bool` | Indicates if the payment account is an externally connected account. |
| `isDefault` | `bool` | Is the default account |
| `isTrustAccount` | `bool` | Indicates if the payment account is a trust account. |
| `isVirtual` | `bool` | True if the account is a virtual account. |
| `lastTransaction` | `dict` |  |
| `lastUpdated` | `str` | Timestamp when the account was last updated. |
| `merchantID` | `str` | The ID of the merchant that owns the account. |
| `merchantName` | `str` | The name of the merchant that owns the account. |
| `nickname` | `str` | Nickname of the account that was provided by the account owner. |
| `physicalAccountID` | `str` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `roleIDs` | `list` | Optional list of role IDs that will get access to the payment account when created. |
| `rules` | `list` | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `float` | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `int` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `str` | Gets a summary of the payments account's most important properties. |
| `supplierPhysicalAccountID` | `str` | For internal use only. |
| `supplierSepaInstantStatus` | `str` | Indicates the status of the SEPA Instant payment rail for this account. |
| `toDate` | `str` | Maximum transaction date for the statement. |
| `type` | `str` | Specifies the type of account e.g. |
| `usageType` | `str` |  |
| `xeroBankFeedConnectionStatus` | `str` | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `str` |  |
| `xeroBankFeedSyncLastFailedAt` | `str` |  |
| `xeroBankFeedSyncLastFailureReason` | `str` |  |
| `xeroBankFeedSyncStatus` | `str` |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | Indicates the number of unsynchronised transactions with Xero |

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
    "account_id": "example_account_id",  # str
    "currency": "example_currency",  # str
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
| `approveUrl` | `str` | This field is used when returning a batch payout record to a client. |
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
| `authenticationMethods` | `list` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `list` | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `int` | The number of distinct authorisers that have authorised the beneficiary. |
| `authorisersRequiredCount` | `int` | The number of authorisers required for this beneficiary. |
| `beneficiaries` | `list` |  |
| `beneficiaryEvents` | `list` |  |
| `canAuthorise` | `bool` | True if the beneficiary can be authorised by the user who loaded it. |
| `canUpdate` | `bool` | True if the beneficiary can be updated by the user who loaded it. |
| `createdBy` | `dict` |  |
| `createdByEmailAddress` | `str` |  |
| `currency` | `str` | Gets or Sets the currency. |
| `destination` | `dict` |  |
| `failedBeneficiaries` | `dict` |  |
| `hasCurrentUserAuthorised` | `bool` | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `isEnabled` | `bool` |  |
| `lastAuthorised` | `str` |  |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` | Gets or Sets the merchant id. |
| `name` | `str` | The descriptive name for the beneficiary. |
| `nonce` | `str` |  |
| `sourceAccountIDs` | `list` | ID of the accounts which are authorised to act as a source for the beneficiary. |
| `sourceAccounts` | `list` |  |
| `theirReference` | `str` | The reference that will be used by default as TheirReference when creating payouts to this beneficiary if no TheirReference is specified for the payout. |

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
    "id": "example_id",  # str
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
| `groupMembers` | `list` | The existing group members. |
| `groupName` | `str` | The descriptive name for the beneficiary group. |
| `id` | `str` |  |
| `inserted` | `str` | Timestamp indicating when the group was created. |
| `lastUpdated` | `str` | Timestamp indicating when the group was last updated. |
| `merchantID` | `str` | Gets or Sets the merchant id. |

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
| `isPayerAuthenticationRequired` | `bool` | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `bool` | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `str` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `str` | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `str` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `int` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `int` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `str` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `str` |  |
| `requestID` | `str` |  |
| `responseCode` | `str` |  |
| `responseType` | `str` |  |
| `status` | `str` |  |
| `threeDSRedirectUrl` | `str` | Checkout.com require a redirect for 3DS authentication. |
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
| `cardType` | `str` | The type of the tokenised card, e.g. |
| `customerEmailAddress` | `str` | When creating a tokenised card the payer's email address must be supplied. |
| `expiryMonth` | `str` |  |
| `expiryYear` | `str` |  |
| `id` | `str` | The unique ID of the card token that has been stored for the customer. |
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
| `isPayerAuthenticationRequired` | `bool` | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `bool` | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `str` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `str` | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `str` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `int` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `int` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `str` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `str` |  |
| `requestID` | `str` |  |
| `responseCode` | `str` |  |
| `responseType` | `str` |  |
| `status` | `str` |  |
| `threeDSRedirectUrl` | `str` | Checkout.com require a redirect for 3DS authentication. |
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
| `authorisationUrl` | `str` | The URL the authorising user needs to be redirected to in order to get the open banking consent token. |
| `callbackUrl` | `str` | Optional callback URL that the end user performing the open banking authorisation will be redirected to on completion. |
| `consentID` | `str` | The ID of the open banking consent. |
| `emailAddress` | `str` | The email address that identifies the end user that will be authorising the open banking consent request. |
| `expiryDate` | `str` |  |
| `failureCallbackUrl` | `str` | Optional callback URL for open banking consent authorisation failure. |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `institutionID` | `str` | The institution ID the open banking consent is being requested for. |
| `isConnectedAccounts` | `bool` | Optional setting. |
| `isEnabled` | `bool` |  |
| `merchantID` | `str` | The ID of the merchant the consent token is being created to be used with. |
| `provider` | `str` | Lists the supported card and PIS processors. |
| `successWebHookUrl` | `str` | A web hook URL to send an HTTP request to when the open banking consent is successfuly authorised. |

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
| `failedSubmissions` | `dict` | Dictionary of failed submissions, keyed by the index (1-based) in the original request. |
| `successfulSubmissions` | `list` | List of successfully submitted direct debit payments. |

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
| `exchangeRate` | `float` | The price at which the transaction will buy the source currency using the destination currency. |
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
| `accountNumber` | `str` | Account number of the customer's bank account in case of GBP account. |
| `addressLine1` | `str` | First line of the customer's address. |
| `addressLine2` | `str` | Second line of the customer's address. |
| `approvedAt` | `str` | Date at which the supplier approved this mandate. |
| `city` | `str` | Customer's city. |
| `countryCode` | `str` | 2-character country code of the customer's bank account. |
| `currency` | `str` | Currency of this mandate. |
| `customerAccountNumber` | `str` | Customer's account number in case of GBP account. |
| `customerCity` | `str` | Customer's city of residence. |
| `customerCountryCode` | `str` | Customer's country of residence code. |
| `customerCountryName` | `str` | Customer's country of residence. |
| `customerEmailAddress` | `str` | Customer's email address. |
| `customerFirstName` | `str` | Customer's first name. |
| `customerIban` | `str` | Customer's IBAN in case of EUR account. |
| `customerLastName` | `str` | Customer's last name. |
| `customerSortCode` | `str` | Customer's sort code in case of GBP account. |
| `emailAddress` | `str` | Customer's email address. |
| `firstName` | `str` | Customer's first name. |
| `iban` | `str` | IBAN of the customer's bank account in case of EUR account. |
| `id` | `str` | Internal ID of the mandate. |
| `inserted` | `str` | The timestamp this mandate was created at. |
| `isRecurring` | `bool` | Whether this mandate is single-use or recurring. |
| `lastName` | `str` | Customer's last name. |
| `lastUpdated` | `str` | The timestamp this mandate was last updated at. |
| `merchantID` | `str` | Internal ID of this mandate's merchant. |
| `postalCode` | `str` | Customer's postal code. |
| `reference` | `str` | Reference assigned to this mandate. |
| `sortCode` | `str` | Sort code of the customer's bank account in case of GBP account. |
| `status` | `str` | General status of this mandate. |
| `supplierBankAccountID` | `str` | ID that the supplier assigned to this mandate's bank account. |
| `supplierCustomerID` | `str` | ID that the supplier assigned to this mandate's customer. |
| `supplierMandateID` | `str` | ID that the supplier assigned to this mandate. |
| `supplierName` | `str` | Name of the supplier used to create this mandate. |
| `supplierStatus` | `str` | Last status that the supplier reported for this mandate. |

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
| `accountCurrencies` | `list` | The list of currencies that the merchant has accounts for. |
| `canHaveTrustAccounts` | `bool` | Trust accounts are a special type of account that allow the account name to be trusted for use in statements and verification of payee checks. |
| `cardPaymentProcessor` | `str` | Name of the card payment processor. |
| `companyID` | `str` | The Company ID recorded in the Compliance system. |
| `displayQrOnHostedPay` | `bool` | Indicates if a QR Code containing the payment link should be displayed on the hosted payment page. |
| `hostedPayVersion` | `int` | The version of the hosted payment page to use with the merchant. |
| `id` | `str` | Unique ID for the merchant. |
| `inserted` | `str` | Timestamp the merchant was added to MoneyMoov. |
| `isBlocked` | `bool` | The merchant is blocked from making payments (payouts). |
| `isExited` | `bool` | The merchant has formally terminated their relationship and is no longer a customer. |
| `isSuspended` | `bool` | The merchant has temporarily suspended their own account. |
| `jurisdiction` | `str` | The jurisdiction the merchant entity is incorporated or established in. |
| `logoUrlPng` | `str` | The CDN URL of the merchant's logo in PNG format. |
| `logoUrlSvg` | `str` | The CDN URL of the merchant's logo in SVG format. |
| `merchantCategoryCode` | `str` | The industry code that represents the merchant's primary trading activity. |
| `name` | `str` | The registered business name of the merchant. |
| `notes` | `str` | The notes field is an optional free text field that can be used to store any additional information about the merchant. |
| `parentMerchant` | `dict` |  |
| `paymentAccountLimit` | `int` | The maximum number of payment accounts that can be created for the Merchant. |
| `paymentAccounts` | `list` |  |
| `reason` | `str` | The reason for the suspension. |
| `shortName` | `str` | A URL friendly shortish name for the merchant. |
| `supportedPaymentMethodsList` | `list` | The payment methods that are configured and supported for this merchant. |
| `suspensionReason` | `str` | The reason for the suspension, provided by the merchant. |
| `tags` | `list` | An optional list of descriptive tags that can be used on merchant entities such as payment requests. |
| `timeZoneId` | `str` | The IANA (Internet Assigned Numbers Authority) time zone identifier of the merchant. |
| `tradingName` | `str` | An optional trading name. |
| `webHookLimit` | `int` | The maximum number of web hooks that can be created for the Merchant. |
| `yourRoleName` | `str` | The name of the role for the identity that loaded the merchant record. |

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
| `approvedAt` | `str` | Date at which the supplier approved this mandate. |
| `currency` | `str` | Currency of this mandate. |
| `customerAccountNumber` | `str` | Customer's account number in case of GBP account. |
| `customerCity` | `str` | Customer's city of residence. |
| `customerCountryCode` | `str` | Customer's country of residence code. |
| `customerCountryName` | `str` | Customer's country of residence. |
| `customerEmailAddress` | `str` | Customer's email address. |
| `customerFirstName` | `str` | Customer's first name. |
| `customerIban` | `str` | Customer's IBAN in case of EUR account. |
| `customerLastName` | `str` | Customer's last name. |
| `customerSortCode` | `str` | Customer's sort code in case of GBP account. |
| `id` | `str` | Internal ID of the mandate. |
| `inserted` | `str` | The timestamp this mandate was created at. |
| `isRecurring` | `bool` | Whether this mandate is single-use or recurring. |
| `lastUpdated` | `str` | The timestamp this mandate was last updated at. |
| `merchantID` | `str` | Internal ID of this mandate's merchant. |
| `reference` | `str` | Reference assigned to this mandate. |
| `status` | `str` | General status of this mandate. |
| `supplierBankAccountID` | `str` | ID that the supplier assigned to this mandate's bank account. |
| `supplierCustomerID` | `str` | ID that the supplier assigned to this mandate's customer. |
| `supplierMandateID` | `str` | ID that the supplier assigned to this mandate. |
| `supplierName` | `str` | Name of the supplier used to create this mandate. |
| `supplierStatus` | `str` | Last status that the supplier reported for this mandate. |

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
| `bankCountryCodes` | `list` | The list of country codes representing the banks the country supports. |
| `bankID` | `str` | ID of the bank to be configured for the merchant. |
| `bankName` | `str` | Name of the Bank/Institution. |
| `businessInstitutionID` | `str` | ID that the processor uses to identify the bank (business accounts). |
| `currency` | `str` | Currency supported by the bank. |
| `logo` | `str` | URL of the bank's logo. |
| `message` | `str` | Message relating to specific bank. |
| `messageImageUrl` | `str` | Optional image URL to be displayed with the message. |
| `order` | `int` | Order in which this setting will appear in the UI. |
| `personalInstitutionID` | `str` | ID that the processor uses to identify the bank (personal accounts). |
| `processor` | `str` | Name of the bank payment processor. |
| `warningHeading` | `str` | The heading for a warning message related to the bank institution to be displayed to the user. |
| `warningMessage` | `str` | The warning message related to the bank institution to be displayed to the user. |

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
| `customFields` | `list` | A list of custom fields that can be included in the payment request template. |
| `defaultFields` | `list` | A list of default fields that are included in the payment request template. |
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
| `authenticationMethods` | `list` | A list of authentication types allowed to authorise the merchant token. |
| `authorisations` | `list` | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `int` | The number of distinct authorisers that have authorised the merchant token. |
| `authorisersRequiredCount` | `int` | The number of authorisers required for this merchant token. |
| `canAuthorise` | `bool` | True if the merchant token can be authorised by the user who loaded it. |
| `description` | `str` | Token description |
| `expiresAt` | `str` | Optional. |
| `hasCurrentUserAuthorised` | `bool` | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `hmacAlgorithm` | `str` | Optional shared secret algorithm to use for HMAC authentication. |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `ipAddressWhitelist` | `str` | Optional. |
| `isArchived` | `bool` | Indicates whether the merchant token is archived. |
| `isEnabled` | `bool` | If set to false the merchant token will not be accepted to authorise a request. |
| `lastAuthorised` | `str` |  |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` | The merchant id to add to the token |
| `nonce` | `str` |  |
| `permissionTypes` | `list` | The permissions that the merchant token supports. |
| `requestSignatureVersion` | `int` | Represent the version of the overall merchant token. |
| `sharedSecretAlgorithm` | `str` | Optional shared secret algorithm to use for HMAC authentication. |
| `sharedSecretBase64` | `str` | The base 64 encoded shared secret that is used for request authentication with an HMAC. |
| `token` | `str` | The JWT merchant token. |

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
| `accountName` | `str` | The name of the account to verify |
| `accountNumber` | `str` | The account number of the account to verify (for CoP checks) |
| `iban` | `str` | The IBAN of the account to verify (for VoP checks) |
| `payeeVerifiedAccountName` | `str` | The verified account name of the payee, if available (in case of a close match) |
| `result` | `str` | The result of the payee verification |
| `secondaryIdentification` | `str` | Optional secondary identifier for the account to verify. |
| `sortCode` | `str` | The sort code of the account to verify (for CoP checks) |

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
| `amount` | `float` | The amount of money to request. |
| `amountPending` | `float` | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `float` | Total amount received for this payment request. |
| `amountRefunded` | `float` | Total amount refunded for this payment request. |
| `autoSendReceipt` | `bool` | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `str` | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `str` | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `bool` | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `bool` | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `str` | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `bool` | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardNoPayerAuthentication` | `bool` | If set to true for card payments no attempt will be made to use payer authentication (3-D Secure and equivalent). |
| `cardProcessorMerchantID` | `str` | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `str` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `str` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `cardTransmitRawDetails` | `bool` | If set to true for card payments the sensitive card number and card verification number will be transmitted directly rather than being tokenised. |
| `createdByUser` | `dict` |  |
| `currency` | `str` | The currency of the request. |
| `customFields` | `list` | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `str` | Optional email address for the customer. |
| `customerID` | `str` | An optional customer identifier for the payment request. |
| `customerName` | `str` |  |
| `description` | `str` | An optional description for the payment request. |
| `destinationAccount` | `dict` |  |
| `directDebitPayment` | `dict` | Contains information about a Direct Debit payment attempt for a payment request. |
| `dueDate` | `str` | The due date for the payment request. |
| `events` | `list` |  |
| `failureCallbackUrl` | `str` | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `list` | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `str` |  |
| `hostedPayCheckoutUrl` | `str` | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `str` |  |
| `ignoreAddressVerification` | `bool` | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `str` | The timestamp the payment request was created at. |
| `insertedSortable` | `str` | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `isArchived` | `bool` | Indicates whether the payment request is archived. |
| `jwk` | `str` | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `str` | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `str` | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `str` | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `str` | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `str` | The ID of the merchant to create the payment request for. |
| `merchantTokenDescription` | `str` | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `str` |  |
| `notificationRoleIDs` | `list` | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `str` | An optional order ID for the payment request. |
| `partialPaymentMethod` | `str` | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `str` | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `list` | The payment attempts made against this payment request. |
| `paymentMethods` | `list` | The payment methods that the payment request supports. |
| `paymentProcessor` | `str` | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `payrunID` | `str` | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `str` | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `str` | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `dict` |  |
| `sandboxSettleDelayInSeconds` | `int` | Sandbox only. |
| `shippingAddress` | `dict` |  |
| `shippingAddressCity` | `str` | Optionally the city of the customer's shipping address. |
| `shippingAddressCountryCode` | `str` | Optionally the country code of the customer's shipping address. |
| `shippingAddressCounty` | `str` | Optionally the state or county of the customer's shipping address. |
| `shippingAddressLine1` | `str` | Optionally the first line of the customer's shipping address. |
| `shippingAddressLine2` | `str` | Optionally the second line of the customer's shipping address. |
| `shippingAddressPostCode` | `str` | Optionally the post code of the customer's shipping address. |
| `shippingEmail` | `str` | Optionally the shipping email address for the customer. |
| `shippingFirstName` | `str` | Optionally the first name of the customer's shipping address. |
| `shippingLastName` | `str` | Optionally the last name of the customer's shipping address. |
| `shippingPhone` | `str` | Optionally the shipping phone number for the customer. |
| `status` | `str` | The current status of the payment request. |
| `successWebHookUrl` | `str` | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tagIds` | `list` | An optional list of tag ids to add to the payment request |
| `tags` | `list` | An optional list of descriptive tags attached to the payment request. |
| `title` | `str` | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `list` |  |
| `transactions` | `list` |  |
| `useHostedPaymentPage` | `bool` | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

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
| `accountName` | `str` | Name for the account |
| `accountSupplierName` | `str` | The payment account supplier name. |
| `availableBalance` | `float` | The current available balance of the account. |
| `availableBalanceMinorUnits` | `int` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `float` | Balance of the account. |
| `balanceMinorUnits` | `int` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `str` | The bank name for external accounts |
| `consentID` | `str` | The ID of the consent used to connect the external account. |
| `createdBy` | `dict` |  |
| `createdByDisplayName` | `str` | Either the name of the user, merchant token or api key that created the account |
| `currency` | `str` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `str` | Indicates the default payment rail for this account. |
| `displayName` | `str` | Gets a unique display name for the payment account. |
| `expiryDate` | `str` | The date that the external account will expire |
| `externalAccountIcon` | `str` | The Icon for external accounts |
| `id` | `str` | Unique id for the account. |
| `identifier` | `dict` |  |
| `inserted` | `str` | Timestamp when the account was created. |
| `isArchived` | `bool` | Indicates whether the account is archived. |
| `isConnectedAccount` | `bool` | Indicates if the payment account is an externally connected account. |
| `isDefault` | `bool` | Is the default account |
| `isTrustAccount` | `bool` | Indicates if the payment account is a trust account. |
| `isVirtual` | `bool` | True if the account is a virtual account. |
| `lastTransaction` | `dict` |  |
| `lastUpdated` | `str` | Timestamp when the account was last updated. |
| `merchantID` | `str` | The ID of the merchant that owns the account. |
| `merchantName` | `str` | The name of the merchant that owns the account. |
| `physicalAccountID` | `str` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `list` | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `float` | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `int` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `str` | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `str` | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `str` | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `str` |  |
| `xeroBankFeedSyncLastFailedAt` | `str` |  |
| `xeroBankFeedSyncLastFailureReason` | `str` |  |
| `xeroBankFeedSyncStatus` | `str` |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | Indicates the number of unsynchronised transactions with Xero |

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
| `accountName` | `str` | Name for the account |
| `availableBalance` | `float` | The current available balance of the account. |
| `balance` | `float` | Balance of the account. |
| `balanceMinorUnits` | `int` | Balance of the account expressed in the currency’s minor units (e.g. |
| `currency` | `str` | Currency of the account in ISO 4217 format |
| `id` | `str` | Unique id for the account. |
| `identifier` | `dict` |  |
| `isArchived` | `bool` | Is the account archived |
| `isConnectedAccount` | `bool` | Indicates if the payment account is an externally connected account. |
| `merchantID` | `str` | The ID of the merchant that owns the account. |
| `submittedPayoutsBalance` | `float` | Total of the payouts that have been submitted for processing. |

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
| `paymentInitiationID` | `str` | The unique identifier of the payment initiation request. |
| `paymentRequestCallbackUrl` | `str` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `str` |  |
| `redirectUrl` | `str` | A redirect URL for the user to authorise the payment initiation request at the ASPSP |
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
| `amount` | `float` | The amount of money to request. |
| `amountPending` | `float` | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `float` | Total amount received for this payment request. |
| `amountRefunded` | `float` | Total amount refunded for this payment request. |
| `autoSendReceipt` | `bool` | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `str` | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `str` | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `bool` | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `bool` | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `str` | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `bool` | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardProcessorMerchantID` | `str` | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `str` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `str` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `createdByUser` | `dict` |  |
| `currency` | `str` | The currency of the request. |
| `customFields` | `list` | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `str` | Optional email address for the customer. |
| `customerID` | `str` | An optional customer identifier for the payment request. |
| `customerName` | `str` |  |
| `description` | `str` | An optional description for the payment request. |
| `destinationAccount` | `dict` |  |
| `directDebitPayment` | `dict` | Contains information about a Direct Debit payment attempt for a payment request. |
| `doSimulateSettlementFailure` | `bool` |  |
| `dueDate` | `str` | The due date for the payment request. |
| `errorDescription` | `str` |  |
| `events` | `list` |  |
| `failedPaymentRequests` | `dict` |  |
| `failureCallbackUrl` | `str` | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `list` | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `str` |  |
| `hostedPayCheckoutUrl` | `str` | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `str` |  |
| `ignoreAddressVerification` | `bool` | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `str` | The timestamp the payment request was created at. |
| `insertedSortable` | `str` | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `institution` | `str` |  |
| `isArchived` | `bool` | Indicates whether the payment request is archived. |
| `jwk` | `str` | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `str` | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `str` | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `str` | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `str` | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `str` |  |
| `merchantTokenDescription` | `str` | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `str` |  |
| `notificationRoleIDs` | `list` | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `str` | An optional order ID for the payment request. |
| `partialPaymentMethod` | `str` | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `str` | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `list` | The payment attempts made against this payment request. |
| `paymentInitiationID` | `str` |  |
| `paymentMethods` | `list` | The payment methods that the payment request supports. |
| `paymentProcessor` | `str` | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `paymentRequests` | `list` |  |
| `payrunID` | `str` | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `str` | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `str` | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `dict` |  |
| `sandboxSettleDelayInSeconds` | `int` | Sandbox only. |
| `shippingAddress` | `dict` |  |
| `status` | `str` | The current status of the payment request. |
| `successWebHookUrl` | `str` | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tags` | `list` | An optional list of descriptive tags attached to the payment request. |
| `title` | `str` | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `list` |  |
| `transactions` | `list` |  |
| `useHostedPaymentPage` | `bool` | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

#### Example: Load

```python
payment_request = client.PaymentRequest().load({"id": "payment_request_id"})
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
| `applePayTransactionID` | `str` | Transaction ID received in Apple pay token. |
| `cardAuthorizationResponseID` | `str` | For a successful card authorization this field will hold the response ID. |
| `cardExpiryMonth` | `int` | For card payment events this field holds the payer's card expiry month. |
| `cardExpiryYear` | `int` | For card payment events this field holds the payer's card expiry year. |
| `cardIssuer` | `str` | For card payment events this field holds the payer's card issuer. |
| `cardIssuerCountry` | `str` | For card payment events this field holds the payer's card issuer country of origin. |
| `cardLastFourDigits` | `str` | For card payment events this field holds the payer's card last four digits. |
| `cardRequestID` | `str` |  |
| `cardScheme` | `str` | For card payment events this field holds the scheme of the payer's card, e.g. |
| `cardTokenCustomerID` | `str` | If the option to create a reusable token for card payments was set this field contains the token the merchant can store to use for repeat payments. |
| `cardTransactionID` | `str` |  |
| `currency` | `str` |  |
| `directDebitPaymentID` | `str` | Payment ID issued by the Direct Debit supplier. |
| `directDebitPaymentReference` | `str` | Reference string issued by the Direct Debit supplier. |
| `drirectDebitMandateID` | `str` | The ID of the mandate that was used wehn requesting payment. |
| `errorMessage` | `str` |  |
| `errorReason` | `str` |  |
| `eventType` | `str` |  |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `lightningInvoice` | `str` | For Bitcoin Lightning payments this field holds the invoice presented to the payer. |
| `lightningRHash` | `str` | For Bitcoin Lightning payments the hash of the invoice presented to the payer. |
| `originUrl` | `str` | Optional field that can be set by payment methods, such as pay by bank, that may want to redirect back to the URL that initiated the attempt in the case of a failure condition. |
| `paymentMethodType` | `str` | The type of payment method the event relates to, e.g. |
| `paymentProcessorName` | `str` | If the event was for a card payment this is the name of the card processor, e.g. |
| `paymentRequestID` | `str` |  |
| `pispBankStatus` | `str` | For payment initiation attempts some providers (e.g. |
| `pispPaymentInitiationID` | `str` | For a payment initiation this is the ID returned by the service provider initiating the payment for us. |
| `pispPaymentInstitutionName` | `str` | For a payment initiation this is the name of the financial institution that is used to initiate and authorise the payment. |
| `pispPaymentServiceProviderID` | `str` | For a payment initiation this is the service provider ID selected by the payer, typically the ID for the bank or similar financial institution. |
| `pispRedirectUrl` | `str` | For a payment initiation this is the redirect URL returned by the service provider initiating the payment for us. |
| `reconciledTransactionID` | `str` | For settlement events (only relevant for non-card payments) this is the payin transaction that the payment request event was reconciled with. |
| `refundPayoutID` | `str` | ID of the Payout that was created for refund. |
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
| `amount` | `float` | The amount of money to request. |
| `amountPending` | `float` | The amount of money that was authorised but has not arrived in the account yet. |
| `amountReceived` | `float` | The amount of money that has been received for this payment request. |
| `amountRefunded` | `float` | The amount of money that has been refunded for this payment request. |
| `callbackUrl` | `str` |  |
| `cardStripePaymentIntentSecret` | `str` |  |
| `countryCode` | `str` | The country code associated with the payment. |
| `currency` | `str` | The currency of the request. |
| `customFieldsToDisplay` | `list` | Custom fields to display to the customer. |
| `description` | `str` | An optional description for the payment request. |
| `dueDate` | `str` | The due date of the payment request. |
| `fieldDisplaySettings` | `list` |  |
| `googlePayMerchantID` | `str` | Merchant ID from Google Pay |
| `id` | `str` |  |
| `jwk` | `str` | The jwk containing the public key |
| `merchantID` | `str` |  |
| `merchantLogoUrlPng` | `str` |  |
| `merchantLogoUrlSvg` | `str` |  |
| `merchantName` | `str` |  |
| `merchantShortName` | `str` |  |
| `partialPaymentMethod` | `str` |  |
| `paymentAttempts` | `list` | The payment attempts for this payment request. |
| `paymentMethodsList` | `list` | The payment methods that the payment request supports. |
| `paymentProcessor` | `str` | The card processor |
| `paymentProcessorKey` | `str` | The card processors public key |
| `pispError` | `str` | This is the error returned from the bank which is recorded in payment request events. |
| `priorityBankID` | `str` |  |
| `status` | `str` | The status of the payment request. |
| `stripeAccountID` | `str` | Account ID of connected customers in Stripe |
| `title` | `str` | The title of the payment request. |

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
| `amount` | `float` | The authorised payment amount. |
| `amountPending` | `float` |  |
| `amountReceived` | `float` |  |
| `amountRefunded` | `float` |  |
| `currency` | `str` | The authorised payment currency. |
| `customerID` | `str` | The customer id |
| `paymentRequestID` | `str` | The ID of the payment request the result is for. |
| `payments` | `list` | The list of payment attempts that have been received for the payment request. |
| `pispAuthorizations` | `list` |  |
| `requestedAmount` | `float` | The full original payment amount requested. |
| `result` | `str` | The result of the payment attempt. |

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
| `accountID` | `str` | Gets or Sets Account Id of sending account |
| `allowIncomplete` | `bool` | If set to true the payout will get created even if the business validation rules fail. |
| `amount` | `float` | Gets or Sets payout amount |
| `amountMinorUnits` | `int` | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `str` | This field is used when returning an payout record to a client. |
| `approverID` | `str` | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `list` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `list` | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `int` | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `int` | The number of authorisers required for this payout. |
| `batchPayoutID` | `str` | The ID of the batch the payout is associated with. |
| `beneficiary` | `dict` |  |
| `beneficiaryID` | `str` | Optional. |
| `canAuthorise` | `bool` | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `bool` | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `bool` | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `str` | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `str` |  |
| `createdByEmailAddress` | `str` |  |
| `currency` | `str` | Gets or Sets Currency of payout request |
| `currentUserID` | `str` | The ID of the user that requested access to the PayOut record. |
| `description` | `str` | Gets or Sets description of payout request |
| `destination` | `dict` |  |
| `documents` | `list` | Documents associated with the payout. |
| `events` | `list` | The activity associated with the payout. |
| `failedPayouts` | `dict` |  |
| `formattedAmount` | `str` | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `str` | FX destination currency and amount formatted string. |
| `formattedSchedule` | `str` |  |
| `formattedScheduleDayOnly` | `str` |  |
| `formattedSourceAccountAvailableBalance` | `str` | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `float` | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `int` | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `str` | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `str` | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `str` | Optional. |
| `fxRate` | `float` | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `bool` | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `bool` | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `str` | The ID for the payout. |
| `inserted` | `str` |  |
| `invoiceID` | `str` | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `bool` | Indicates whether the payout is archived. |
| `isFailed` | `bool` | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `bool` | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `bool` | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `str` |  |
| `nonce` | `str` |  |
| `paymentProcessor` | `str` | The usptream payment processor for the payout. |
| `paymentRail` | `str` | Optional field to indicate the payment rail to use for the payout. |
| `payouts` | `list` |  |
| `payrunID` | `str` | The ID of the payrun that this payout is associated with. |
| `payrunName` | `str` | The name of the payrun that this payout is associated with. |
| `reason` | `str` |  |
| `rule` | `dict` |  |
| `scheduleDate` | `str` | The date the payout should be submitted. |
| `scheduled` | `bool` | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `float` | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `int` | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `str` | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `str` | The currency of the source account. |
| `sourceAccountIban` | `str` | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `dict` |  |
| `sourceAccountName` | `str` | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `str` | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `str` | The sort code of the account the payout is being made from. |
| `status` | `str` | Gets or Sets the status of payout request |
| `tagIds` | `list` | An optional list of tag ids to add to the payout. |
| `tags` | `list` | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `str` | Gets or Sets destination reference ID |
| `topupPayrunID` | `str` | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `float` | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `float` | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `float` | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `str` | Gets or Sets payout type |
| `userID` | `str` | Gets or Sets User ID of who created the payout request |
| `yourReference` | `str` | Gets or Sets your reference ID |

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
    "id": "example_id",  # str
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
| `accountID` | `str` | Gets or Sets Account Id of sending account |
| `amount` | `float` | Gets or Sets payout amount |
| `amountMinorUnits` | `int` | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `str` | This field is used when returning an payout record to a client. |
| `approverID` | `str` | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `list` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `list` | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `int` | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `int` | The number of authorisers required for this payout. |
| `batchPayoutID` | `str` | The ID of the batch the payout is associated with. |
| `beneficiary` | `dict` |  |
| `canAuthorise` | `bool` | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `bool` | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `bool` | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `str` | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `str` |  |
| `createdByEmailAddress` | `str` |  |
| `currency` | `str` | Gets or Sets Currency of payout request |
| `currentUserID` | `str` | The ID of the user that requested access to the PayOut record. |
| `description` | `str` | Gets or Sets description of payout request |
| `destination` | `dict` |  |
| `documents` | `list` | Documents associated with the payout. |
| `events` | `list` | The activity associated with the payout. |
| `formattedAmount` | `str` | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `str` | FX destination currency and amount formatted string. |
| `formattedSchedule` | `str` |  |
| `formattedScheduleDayOnly` | `str` |  |
| `formattedSourceAccountAvailableBalance` | `str` | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `float` | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `int` | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `str` | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `str` | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `str` | Optional. |
| `fxRate` | `float` | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `bool` | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `bool` | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `str` | The ID for the payout. |
| `inserted` | `str` |  |
| `invoiceID` | `str` | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `bool` | Indicates whether the payout is archived. |
| `isFailed` | `bool` | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `bool` | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `bool` | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `str` |  |
| `nonce` | `str` |  |
| `paymentProcessor` | `str` | The usptream payment processor for the payout. |
| `paymentRail` | `str` | Optional field to indicate the payment rail to use for the payout. |
| `payrunID` | `str` | The ID of the payrun that this payout is associated with. |
| `payrunName` | `str` | The name of the payrun that this payout is associated with. |
| `rule` | `dict` |  |
| `scheduleDate` | `str` | The date the payout should be submitted. |
| `scheduled` | `bool` | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `float` | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `int` | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `str` | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `str` | The currency of the source account. |
| `sourceAccountIban` | `str` | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `dict` |  |
| `sourceAccountName` | `str` | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `str` | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `str` | The sort code of the account the payout is being made from. |
| `status` | `str` | Gets or Sets the status of payout request |
| `tags` | `list` | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `str` | Gets or Sets destination reference ID |
| `topupPayrunID` | `str` | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `float` | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `float` | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `float` | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `str` | Gets or Sets payout type |
| `userID` | `str` | Gets or Sets User ID of who created the payout request |
| `yourReference` | `str` | Gets or Sets your reference ID |

#### Example: List

```python
payout_keyset_pages = client.PayoutKeysetPage().list({"merchant_id": "example"})
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
| `authorisations` | `list` | A list of the users who have successfully authorised the latest version of the payrun and when. |
| `authorisersCompletedCount` | `int` | The number of distinct authorisers that have authorised the payrun. |
| `authorisersRequiredCount` | `int` | The number of authorisers required for this payrun. |
| `batchPayoutID` | `str` |  |
| `canAuthorise` | `bool` | True if the payrun can be authorised by the user who loaded it. |
| `canDelete` | `bool` |  |
| `canEdit` | `bool` |  |
| `events` | `list` |  |
| `hasCurrentUserAuthorised` | `bool` | True if the payrun was loaded for a user and that user has already authorised the latest version of the payrun. |
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
| `accountID` | `str` | The ID of the account the rule will apply to. |
| `approveUrl` | `str` | If set this property holds the URL an approver needs to visit in order to complete a strong authentication check in order to approve the rule. |
| `approverID` | `str` |  |
| `authenticationMethods` | `list` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `list` | A list of the users who have successfully authorised the latest version of the rule and when. |
| `authorisersCompletedCount` | `int` | The number of distinct authorisers that have authorised the rule. |
| `authorisersRequiredCount` | `int` | The number of authorisers required for this rule. |
| `canAuthorise` | `bool` | True if the rule can be authorised by the user who loaded it. |
| `createdBy` | `dict` |  |
| `description` | `str` | Arbitrary description for the rule. |
| `endAt` | `str` | Optional end time for rule executions. |
| `hasCurrentUserAuthorised` | `bool` | True if the current user has authorised. |
| `id` | `str` |  |
| `inserted` | `str` |  |
| `isDisabled` | `bool` | If set to true the rule will be disabled from executing. |
| `lastExecutedAt` | `str` |  |
| `lastRunAtTransactionDate` | `str` | The most recent transaction date when the rule was last run. |
| `lastUpdated` | `str` |  |
| `merchantID` | `str` | The ID of the merchant that owns the account. |
| `name` | `str` | A name to succinctly describe the rule. |
| `nonce` | `str` |  |
| `onApprovedWebHookUrl` | `str` | Optional URL to receive an HTTP request with the rule details when the rule status changes to approved. |
| `onExecutionErrorWebHookUrl` | `str` | Optional URL to receive an HTTP request when a rule execution attempt fails. |
| `onExecutionSuccessWebHookUrl` | `str` | Optional URL to receive an HTTP request when a rule execution attempt succeeds. |
| `startAt` | `str` | Optional start time for rule executions. |
| `status` | `str` |  |
| `sweepAction` | `dict` |  |
| `timeZoneId` | `str` | If the rule should be executed on a recurring schedule this is the timezone that the CRON expression should be evaluated in. |
| `triggerCronExpression` | `str` | If the rule should be executed on a recurring schedule this is the expression that sets the schedule. |
| `triggerOnPayIn` | `bool` | Set to true if the rule execution should be triggered when the account receives a pay in (credit). |
| `userID` | `str` |  |
| `webHookSecret` | `str` | If set this secret will be used to sign Web Hook requests. |

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
| `accountID` | `str` | The ID of the account the transaction belongs to. |
| `accountName` | `str` | The name of the account the transaction belongs to. |
| `accountSequenceNumber` | `int` | The sequence number of transaction on a per account basis. |
| `addressDetails` | `dict` |  |
| `amount` | `float` | Amount of the transaction. |
| `amountMinorUnits` | `int` | Amount of the transaction expressed in the currency’s minor units (e.g. |
| `balance` | `float` | Balance left on the account after the transaction. |
| `balanceMinorUnits` | `int` | Balance on the account expressed in the currency’s minor units (e.g. |
| `bookingDateTime` | `str` |  |
| `chargeDetails` | `dict` |  |
| `content` | `list` |  |
| `counterparty` | `dict` |  |
| `counterpartySummary` | `str` | For pay in (credit) transactions this will contain a descriptive string with the most important fields about the counterparty. |
| `currency` | `str` | Currency of transaction. |
| `currencyExchange` | `dict` | Provides details on the currency exchange. |
| `date` | `str` |  |
| `description` | `str` | Description of the transaction. |
| `enrichment` | `dict` |  |
| `fxAmount` | `float` | For an FX payout this is the amound in the FX currency. |
| `fxCurrency` | `str` | For an FX payout this is the currency that was received or that was instructed. |
| `fxRate` | `float` | For an FX payout this is the exchange rate between the transaction currency and the FX currency. |
| `grossAmount` | `dict` |  |
| `id` | `str` | Unique ID for the transaction. |
| `inserted` | `str` | Date when the transaction was inserted into the ledger. |
| `isoBankTransactionCode` | `dict` |  |
| `merchant` | `dict` |  |
| `merchantID` | `str` | The ID of the merchant that owns the account. |
| `pageNumber` | `int` | Current page number. |
| `pageSize` | `int` | Page size |
| `payeeDetails` | `dict` | The Payee object contains details of the beneficiary, person or business. |
| `payerDetails` | `dict` |  |
| `paymentRequestCustomFields` | `dict` | The custom fields that were attached to the payment request that resulted in this transaction. |
| `paymentRequestID` | `str` | For Pay by Bank and Direct Debit transactions this will contain the ID of the payment request. |
| `payoutID` | `str` | ID of the payout that resulted in the transaction. |
| `proprietaryBankTransactionCode` | `dict` |  |
| `rawReference` | `str` | The raw payment reference details as received from the payment processor. |
| `reference` | `str` |  |
| `ruleID` | `str` | ID of the rule that resulted in the transaction. |
| `statementReferences` | `list` |  |
| `status` | `str` |  |
| `supplementaryData` | `Any` |  |
| `tags` | `list` | An optional list of descriptive tags attached to the transaction. |
| `theirReference` | `str` | For a pay out the reference that the payer attached for the receiving party. |
| `totalPages` | `int` | Total pages |
| `totalSize` | `int` | Total count |
| `transactionAmount` | `dict` |  |
| `transactionDate` | `str` | Date when the transaction occurred. |
| `transactionInformation` | `list` |  |
| `transactionMutability` | `str` |  |
| `type` | `str` | Type of the transaction. |
| `valueDateTime` | `str` |  |
| `virtualIBAN` | `str` | If set it indicates the payin was to a virtual IBAN. |
| `yourReference` | `str` | For a pay in the reference the sending party attached. |

#### Example: Load

```python
transaction = client.Transaction().load({"id": "transaction_id"})
```

#### Example: List

```python
transactions = client.Transaction().list({"account_id": "example", "id": "example_id"})
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
| `clientSessionTimeouts` | `list` | The number of seconds a session for this user should last before expiring. |
| `emailAddress` | `str` |  |
| `firstName` | `str` |  |
| `id` | `str` |  |
| `lastName` | `str` |  |
| `passkeyAdded` | `bool` |  |
| `permissions` | `dict` |  |
| `profile` | `str` |  |
| `rolesWithScope` | `list` |  |
| `twoFactorEnabled` | `bool` |  |
| `userInviteID` | `str` | Optional ID of the invite that is being accepted so the user can be assigned a role on a new merchant. |

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
| `initialRoleID` | `str` | The role ID to automatically assign to the merchant’s very first user. |
| `inviteeEmailAddress` | `str` | Email address of the user being invited. |
| `inviteeFirstName` | `str` | First Name of the user being invited. |
| `inviteeLastName` | `str` | Last Name of the user being invited. |
| `inviterEmailAddress` | `str` |  |
| `inviterFirstName` | `str` |  |
| `inviterLastName` | `str` |  |
| `isAuthorised` | `bool` | Will be set to true once the invite has met the authorisation requirements. |
| `isInviteeRegistered` | `bool` | If true, indicates the invitee's email address corresponds to an existing MoneyMoov user. |
| `lastInvited` | `str` |  |
| `merchantID` | `str` | ID of the merchant the user is being invited to. |
| `merchantName` | `str` |  |
| `message` | `str` |  |
| `registrationUrl` | `str` |  |
| `sendInviteEmail` | `bool` | If set to true an email will be sent to the invitee with instructions on how to accept the invite. |
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
    "id": "example_id",  # str
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
| `accountName` | `str` | Name for the account |
| `accountSupplierName` | `str` | The payment account supplier name. |
| `availableBalance` | `float` | The current available balance of the account. |
| `availableBalanceMinorUnits` | `int` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `float` | Balance of the account. |
| `balanceMinorUnits` | `int` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `str` | The bank name for external accounts |
| `consentID` | `str` | The ID of the consent used to connect the external account. |
| `createdBy` | `dict` |  |
| `createdByDisplayName` | `str` | Either the name of the user, merchant token or api key that created the account |
| `currency` | `str` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `str` | Indicates the default payment rail for this account. |
| `displayName` | `str` | Gets a unique display name for the payment account. |
| `expiryDate` | `str` | The date that the external account will expire |
| `externalAccountIcon` | `str` | The Icon for external accounts |
| `id` | `str` | Unique id for the account. |
| `identifier` | `dict` |  |
| `inserted` | `str` | Timestamp when the account was created. |
| `isArchived` | `bool` | Indicates whether the account is archived. |
| `isConnectedAccount` | `bool` | Indicates if the payment account is an externally connected account. |
| `isDefault` | `bool` | Is the default account |
| `isTrustAccount` | `bool` | Indicates if the payment account is a trust account. |
| `isVirtual` | `bool` | True if the account is a virtual account. |
| `lastTransaction` | `dict` |  |
| `lastUpdated` | `str` | Timestamp when the account was last updated. |
| `merchantID` | `str` | The ID of the merchant that owns the account. |
| `merchantName` | `str` | The name of the merchant that owns the account. |
| `name` | `str` | The name of the virtual account. |
| `physicalAccountID` | `str` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `list` | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `float` | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `int` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `str` | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `str` | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `str` | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `str` |  |
| `xeroBankFeedSyncLastFailedAt` | `str` |  |
| `xeroBankFeedSyncLastFailureReason` | `str` |  |
| `xeroBankFeedSyncStatus` | `str` |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | Indicates the number of unsynchronised transactions with Xero |

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
| `destinationUrl` | `str` | The destination URL for the webhook. |
| `emailAddress` | `str` | The recipient email address(es) for notifications. |
| `failedNotificationEmailAddress` | `str` | The email address to which notifications about failed webhook deliveries will be sent. |
| `id` | `str` |  |
| `isActive` | `bool` |  |
| `merchantID` | `str` | The ID of the merchant that the webhook is for. |
| `notificationMethod` | `str` | The type of notification that will be sent. |
| `resourceTypes` | `list` | The resource types that the webhook will be generated for. |
| `retry` | `bool` |  |
| `secret` | `str` | The secret key required to authenticate webhook notifications. |
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
