# Nofrixion Python SDK Reference

Complete API reference for the Nofrixion Python SDK.


## NofrixionSDK

### Constructor

```python
from nofrixion_sdk import NofrixionSDK

client = NofrixionSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NofrixionSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = NofrixionSDK.test()
```


### Instance Methods

#### `Account(data=None)`

Create a new `AccountEntity` instance. Pass `None` for no initial data.

#### `Batch(data=None)`

Create a new `BatchEntity` instance. Pass `None` for no initial data.

#### `Beneficiary(data=None)`

Create a new `BeneficiaryEntity` instance. Pass `None` for no initial data.

#### `BeneficiaryGroup(data=None)`

Create a new `BeneficiaryGroupEntity` instance. Pass `None` for no initial data.

#### `Card(data=None)`

Create a new `CardEntity` instance. Pass `None` for no initial data.

#### `CardCustomerToken(data=None)`

Create a new `CardCustomerTokenEntity` instance. Pass `None` for no initial data.

#### `CardPayment(data=None)`

Create a new `CardPaymentEntity` instance. Pass `None` for no initial data.

#### `CardPublicKey(data=None)`

Create a new `CardPublicKeyEntity` instance. Pass `None` for no initial data.

#### `Consent(data=None)`

Create a new `ConsentEntity` instance. Pass `None` for no initial data.

#### `Currency(data=None)`

Create a new `CurrencyEntity` instance. Pass `None` for no initial data.

#### `DirectDebitBatchSubmit(data=None)`

Create a new `DirectDebitBatchSubmitEntity` instance. Pass `None` for no initial data.

#### `FxRate(data=None)`

Create a new `FxRateEntity` instance. Pass `None` for no initial data.

#### `IPayment(data=None)`

Create a new `IPaymentEntity` instance. Pass `None` for no initial data.

#### `Mandate(data=None)`

Create a new `MandateEntity` instance. Pass `None` for no initial data.

#### `Merchant(data=None)`

Create a new `MerchantEntity` instance. Pass `None` for no initial data.

#### `MerchantAuthorisationSetting(data=None)`

Create a new `MerchantAuthorisationSettingEntity` instance. Pass `None` for no initial data.

#### `MerchantDirectDebitMandatePage(data=None)`

Create a new `MerchantDirectDebitMandatePageEntity` instance. Pass `None` for no initial data.

#### `MerchantPayByBankSetting(data=None)`

Create a new `MerchantPayByBankSettingEntity` instance. Pass `None` for no initial data.

#### `MerchantPaymentRequestTemplate(data=None)`

Create a new `MerchantPaymentRequestTemplateEntity` instance. Pass `None` for no initial data.

#### `MerchantToken(data=None)`

Create a new `MerchantTokenEntity` instance. Pass `None` for no initial data.

#### `Metadata(data=None)`

Create a new `MetadataEntity` instance. Pass `None` for no initial data.

#### `NoFrixionVersion(data=None)`

Create a new `NoFrixionVersionEntity` instance. Pass `None` for no initial data.

#### `OpenBanking(data=None)`

Create a new `OpenBankingEntity` instance. Pass `None` for no initial data.

#### `Payeeverification(data=None)`

Create a new `PayeeverificationEntity` instance. Pass `None` for no initial data.

#### `Payment(data=None)`

Create a new `PaymentEntity` instance. Pass `None` for no initial data.

#### `PaymentAccount(data=None)`

Create a new `PaymentAccountEntity` instance. Pass `None` for no initial data.

#### `PaymentAccountMinimal(data=None)`

Create a new `PaymentAccountMinimalEntity` instance. Pass `None` for no initial data.

#### `PaymentInitiation(data=None)`

Create a new `PaymentInitiationEntity` instance. Pass `None` for no initial data.

#### `PaymentRequest(data=None)`

Create a new `PaymentRequestEntity` instance. Pass `None` for no initial data.

#### `PaymentRequestEvent(data=None)`

Create a new `PaymentRequestEventEntity` instance. Pass `None` for no initial data.

#### `PaymentRequestMetric(data=None)`

Create a new `PaymentRequestMetricEntity` instance. Pass `None` for no initial data.

#### `PaymentRequestMinimal(data=None)`

Create a new `PaymentRequestMinimalEntity` instance. Pass `None` for no initial data.

#### `PaymentRequestResult(data=None)`

Create a new `PaymentRequestResultEntity` instance. Pass `None` for no initial data.

#### `Payout(data=None)`

Create a new `PayoutEntity` instance. Pass `None` for no initial data.

#### `PayoutKeysetPage(data=None)`

Create a new `PayoutKeysetPageEntity` instance. Pass `None` for no initial data.

#### `PayoutMetric(data=None)`

Create a new `PayoutMetricEntity` instance. Pass `None` for no initial data.

#### `Payrun(data=None)`

Create a new `PayrunEntity` instance. Pass `None` for no initial data.

#### `Report(data=None)`

Create a new `ReportEntity` instance. Pass `None` for no initial data.

#### `ReportResult(data=None)`

Create a new `ReportResultEntity` instance. Pass `None` for no initial data.

#### `Role(data=None)`

Create a new `RoleEntity` instance. Pass `None` for no initial data.

#### `Rule(data=None)`

Create a new `RuleEntity` instance. Pass `None` for no initial data.

#### `RuleEvent(data=None)`

Create a new `RuleEventEntity` instance. Pass `None` for no initial data.

#### `Tag(data=None)`

Create a new `TagEntity` instance. Pass `None` for no initial data.

#### `Token(data=None)`

Create a new `TokenEntity` instance. Pass `None` for no initial data.

#### `Transaction(data=None)`

Create a new `TransactionEntity` instance. Pass `None` for no initial data.

#### `User(data=None)`

Create a new `UserEntity` instance. Pass `None` for no initial data.

#### `UserInvite(data=None)`

Create a new `UserInviteEntity` instance. Pass `None` for no initial data.

#### `Virtual(data=None)`

Create a new `VirtualEntity` instance. Pass `None` for no initial data.

#### `Webhook(data=None)`

Create a new `WebhookEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## AccountEntity

```python
account = client.Account()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountBalances` | `list` | No | The various balances for the account. |
| `accountID` | `str` | No | ID of the account. |
| `accountIdentifications` | `list` | No | The canoncial identifiers for the account. |
| `accountName` | `str` | No | Name for the account |
| `accountNames` | `list` | No | Optional account names set by the account holder. |
| `accountSupplierName` | `str` | No | The payment account supplier name. |
| `accountType` | `str` | No | The type of account e.g. |
| `availableBalance` | `float` | No | The current available balance of the account. |
| `availableBalanceMinorUnits` | `int` | No | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `float` | No | Balance of the account. |
| `balanceMinorUnits` | `int` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `str` | No | The bank name for external accounts |
| `consentID` | `str` | No | The ID of the consent used to connect the external account. |
| `consolidatedAccountInformation` | `dict` | No | Summary information regarding account balances of the overall account provided by the bank. |
| `createdBy` | `dict` | Yes |  |
| `createdByDisplayName` | `str` | No | Either the name of the user, merchant token or api key that created the account |
| `currency` | `str` | No | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `str` | No | Indicates the default payment rail for this account. |
| `description` | `str` | No | Product name as defined by the financial institution for this account. |
| `details` | `str` | No | Supplementary specifications that might be provided by the Bank. |
| `displayName` | `str` | No | Gets a unique display name for the payment account. |
| `expiryDate` | `str` | No | The date that the external account will expire |
| `externalAccountIcon` | `str` | No | The Icon for external accounts |
| `format` | `str` | No | File format to save the statement as. |
| `fromDate` | `str` | No | Minimum transaction date for the statement. |
| `id` | `str` | No | Unique id for the account. |
| `identifier` | `dict` | Yes |  |
| `inserted` | `str` | No | Timestamp when the account was created. |
| `isArchived` | `bool` | No | Indicates whether the account is archived. |
| `isConnectedAccount` | `bool` | No | Indicates if the payment account is an externally connected account. |
| `isDefault` | `bool` | No | Is the default account |
| `isTrustAccount` | `bool` | No | Indicates if the payment account is a trust account. |
| `isVirtual` | `bool` | No | True if the account is a virtual account. |
| `lastTransaction` | `dict` | No |  |
| `lastUpdated` | `str` | No | Timestamp when the account was last updated. |
| `merchantID` | `str` | No | The ID of the merchant that owns the account. |
| `merchantName` | `str` | No | The name of the merchant that owns the account. |
| `nickname` | `str` | No | Nickname of the account that was provided by the account owner. |
| `physicalAccountID` | `str` | No | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `roleIDs` | `list` | No | Optional list of role IDs that will get access to the payment account when created. |
| `rules` | `list` | No | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `float` | No | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `int` | No | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `str` | No | Gets a summary of the payments account's most important properties. |
| `supplierPhysicalAccountID` | `str` | No | For internal use only. |
| `supplierSepaInstantStatus` | `str` | No | Indicates the status of the SEPA Instant payment rail for this account. |
| `toDate` | `str` | No | Maximum transaction date for the statement. |
| `type` | `str` | No | Specifies the type of account e.g. |
| `usageType` | `str` | No |  |
| `xeroBankFeedConnectionStatus` | `str` | No | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `str` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `str` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `str` | No |  |
| `xeroBankFeedSyncStatus` | `str` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | No | Indicates the number of unsynchronised transactions with Xero |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Account().create({
    "account_id": "example_account_id",  # str
    "currency": "example_currency",  # str
    "createdBy": {},  # dict
    "identifier": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Account().list()
for account in results:
    print(account)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Account().load({"id": "account_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Account().remove({"id": "account_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Account().update({
    "id": "account_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `AccountEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## BatchEntity

```python
batch = client.Batch()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approveUrl` | `str` | No | This field is used when returning a batch payout record to a client. |
| `id` | `str` | No |  |
| `payouts` | `list` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Batch().create({
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Batch().load({"id": "batch_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BatchEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## BeneficiaryEntity

```python
beneficiary = client.Beneficiary()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approvalCallbackUrl` | `str` | No |  |
| `authenticationMethods` | `list` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `list` | No | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `int` | No | The number of distinct authorisers that have authorised the beneficiary. |
| `authorisersRequiredCount` | `int` | No | The number of authorisers required for this beneficiary. |
| `beneficiaries` | `list` | No |  |
| `beneficiaryEvents` | `list` | No |  |
| `canAuthorise` | `bool` | No | True if the beneficiary can be authorised by the user who loaded it. |
| `canUpdate` | `bool` | No | True if the beneficiary can be updated by the user who loaded it. |
| `createdBy` | `dict` | Yes |  |
| `createdByEmailAddress` | `str` | No |  |
| `currency` | `str` | Yes | Gets or Sets the currency. |
| `destination` | `dict` | No |  |
| `failedBeneficiaries` | `dict` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `isEnabled` | `bool` | No |  |
| `lastAuthorised` | `str` | No |  |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No | Gets or Sets the merchant id. |
| `name` | `str` | Yes | The descriptive name for the beneficiary. |
| `nonce` | `str` | No |  |
| `sourceAccountIDs` | `list` | No | ID of the accounts which are authorised to act as a source for the beneficiary. |
| `sourceAccounts` | `list` | No |  |
| `theirReference` | `str` | No | The reference that will be used by default as TheirReference when creating payouts to this beneficiary if no TheirReference is specified for the payout. |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Beneficiary().create({
    "id": "example_id",  # str
    "createdBy": {},  # dict
    "currency": "example_currency",  # str
    "name": "example_name",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Beneficiary().list()
for beneficiary in results:
    print(beneficiary)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Beneficiary().load({"id": "beneficiary_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Beneficiary().remove({"id": "beneficiary_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Beneficiary().update({
    "id": "beneficiary_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BeneficiaryEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## BeneficiaryGroupEntity

```python
beneficiary_group = client.BeneficiaryGroup()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `groupMembers` | `list` | No | The existing group members. |
| `groupName` | `str` | Yes | The descriptive name for the beneficiary group. |
| `id` | `str` | No |  |
| `inserted` | `str` | No | Timestamp indicating when the group was created. |
| `lastUpdated` | `str` | No | Timestamp indicating when the group was last updated. |
| `merchantID` | `str` | Yes | Gets or Sets the merchant id. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.BeneficiaryGroup().list({"merchant_id": "example"})
for beneficiary_group in results:
    print(beneficiary_group)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BeneficiaryGroupEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CardEntity

```python
card = client.Card()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorizedAmount` | `str` | No |  |
| `currencyCode` | `str` | No |  |
| `isPayerAuthenticationRequired` | `bool` | No | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `bool` | No | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `str` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `str` | No | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `str` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `int` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `int` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `str` | No | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `str` | No |  |
| `requestID` | `str` | No |  |
| `responseCode` | `str` | No |  |
| `responseType` | `str` | No |  |
| `status` | `str` | No |  |
| `threeDSRedirectUrl` | `str` | No | Checkout.com require a redirect for 3DS authentication. |
| `transactionID` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Card().create({
    "paymentrequest_id": "example_paymentrequest_id",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CardEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CardCustomerTokenEntity

```python
card_customer_token = client.CardCustomerToken()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cardType` | `str` | No | The type of the tokenised card, e.g. |
| `customerEmailAddress` | `str` | No | When creating a tokenised card the payer's email address must be supplied. |
| `expiryMonth` | `str` | No |  |
| `expiryYear` | `str` | No |  |
| `id` | `str` | No | The unique ID of the card token that has been stored for the customer. |
| `inserted` | `str` | No |  |
| `lastFourDigits` | `str` | No |  |
| `lastUpdated` | `str` | No |  |
| `maskedCardNumber` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `paymentRequestID` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.CardCustomerToken().list({"customer_email_address": "example", "merchant_id": "example"})
for card_customer_token in results:
    print(card_customer_token)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.CardCustomerToken().load({"customer_email_address": "customer_email_address"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.CardCustomerToken().remove({"id": "id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CardCustomerTokenEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CardPaymentEntity

```python
card_payment = client.CardPayment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorizedAmount` | `str` | No |  |
| `currencyCode` | `str` | No |  |
| `isPayerAuthenticationRequired` | `bool` | No | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `bool` | No | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `str` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `str` | No | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `str` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `int` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `int` | No | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `str` | No | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `str` | No |  |
| `requestID` | `str` | No |  |
| `responseCode` | `str` | No |  |
| `responseType` | `str` | No |  |
| `status` | `str` | No |  |
| `threeDSRedirectUrl` | `str` | No | Checkout.com require a redirect for 3DS authentication. |
| `transactionID` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.CardPayment().create({
    "paymentrequest_id": "example_paymentrequest_id",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CardPaymentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CardPublicKeyEntity

```python
card_public_key = client.CardPublicKey()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jwt` | `str` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.CardPublicKey().load({"paymentrequest_id": "paymentrequest_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CardPublicKeyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ConsentEntity

```python
consent = client.Consent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisationUrl` | `str` | No | The URL the authorising user needs to be redirected to in order to get the open banking consent token. |
| `callbackUrl` | `str` | No | Optional callback URL that the end user performing the open banking authorisation will be redirected to on completion. |
| `consentID` | `str` | No | The ID of the open banking consent. |
| `emailAddress` | `str` | No | The email address that identifies the end user that will be authorising the open banking consent request. |
| `expiryDate` | `str` | No |  |
| `failureCallbackUrl` | `str` | No | Optional callback URL for open banking consent authorisation failure. |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `institutionID` | `str` | No | The institution ID the open banking consent is being requested for. |
| `isConnectedAccounts` | `bool` | No | Optional setting. |
| `isEnabled` | `bool` | No |  |
| `merchantID` | `str` | No | The ID of the merchant the consent token is being created to be used with. |
| `provider` | `str` | No | Lists the supported card and PIS processors. |
| `successWebHookUrl` | `str` | No | A web hook URL to send an HTTP request to when the open banking consent is successfuly authorised. |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Consent().create({
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Consent().list({"email": "example", "merchant_id": "example"})
for consent in results:
    print(consent)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Consent().load({"id": "consent_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Consent().remove({"id": "consent_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Consent().update({
    "id": "consent_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ConsentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## CurrencyEntity

```python
currency = client.Currency()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `code` | `str` | No |  |
| `decimals` | `int` | No |  |
| `isFiat` | `bool` | No |  |
| `iso4217AlphaCode` | `str` | No |  |
| `iso4217NumericCode` | `str` | No |  |
| `symbol` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Currency().list()
for currency in results:
    print(currency)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `CurrencyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## DirectDebitBatchSubmitEntity

```python
direct_debit_batch_submit = client.DirectDebitBatchSubmit()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failedSubmissions` | `dict` | No | Dictionary of failed submissions, keyed by the index (1-based) in the original request. |
| `successfulSubmissions` | `list` | No | List of successfully submitted direct debit payments. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.DirectDebitBatchSubmit().create({
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `DirectDebitBatchSubmitEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## FxRateEntity

```python
fx_rate = client.FxRate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destinationCurrency` | `str` | No |  |
| `exchangeRate` | `float` | No | The price at which the transaction will buy the source currency using the destination currency. |
| `expiryTime` | `str` | No |  |
| `quoteID` | `str` | No |  |
| `sourceCurrency` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.FxRate().list({"destination": "example", "source": "example"})
for fx_rate in results:
    print(fx_rate)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.FxRate().load({"destination": "destination", "source": "source", "valid_for_minute": 1})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FxRateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## IPaymentEntity

```python
i_payment = client.IPayment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `paymentRequestID` | `str` | No |  |
| `responseType` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.IPayment().create({
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `IPaymentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MandateEntity

```python
mandate = client.Mandate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountNumber` | `str` | No | Account number of the customer's bank account in case of GBP account. |
| `addressLine1` | `str` | Yes | First line of the customer's address. |
| `addressLine2` | `str` | No | Second line of the customer's address. |
| `approvedAt` | `str` | No | Date at which the supplier approved this mandate. |
| `city` | `str` | Yes | Customer's city. |
| `countryCode` | `str` | Yes | 2-character country code of the customer's bank account. |
| `currency` | `str` | No | Currency of this mandate. |
| `customerAccountNumber` | `str` | No | Customer's account number in case of GBP account. |
| `customerCity` | `str` | No | Customer's city of residence. |
| `customerCountryCode` | `str` | No | Customer's country of residence code. |
| `customerCountryName` | `str` | No | Customer's country of residence. |
| `customerEmailAddress` | `str` | No | Customer's email address. |
| `customerFirstName` | `str` | No | Customer's first name. |
| `customerIban` | `str` | No | Customer's IBAN in case of EUR account. |
| `customerLastName` | `str` | No | Customer's last name. |
| `customerSortCode` | `str` | No | Customer's sort code in case of GBP account. |
| `emailAddress` | `str` | Yes | Customer's email address. |
| `firstName` | `str` | Yes | Customer's first name. |
| `iban` | `str` | No | IBAN of the customer's bank account in case of EUR account. |
| `id` | `str` | No | Internal ID of the mandate. |
| `inserted` | `str` | No | The timestamp this mandate was created at. |
| `isRecurring` | `bool` | No | Whether this mandate is single-use or recurring. |
| `lastName` | `str` | Yes | Customer's last name. |
| `lastUpdated` | `str` | No | The timestamp this mandate was last updated at. |
| `merchantID` | `str` | No | Internal ID of this mandate's merchant. |
| `postalCode` | `str` | Yes | Customer's postal code. |
| `reference` | `str` | No | Reference assigned to this mandate. |
| `sortCode` | `str` | No | Sort code of the customer's bank account in case of GBP account. |
| `status` | `str` | No | General status of this mandate. |
| `supplierBankAccountID` | `str` | No | ID that the supplier assigned to this mandate's bank account. |
| `supplierCustomerID` | `str` | No | ID that the supplier assigned to this mandate's customer. |
| `supplierMandateID` | `str` | No | ID that the supplier assigned to this mandate. |
| `supplierName` | `str` | No | Name of the supplier used to create this mandate. |
| `supplierStatus` | `str` | No | Last status that the supplier reported for this mandate. |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Mandate().create({
    "addressLine1": "example_addressLine1",  # str
    "city": "example_city",  # str
    "countryCode": "example_countryCode",  # str
    "emailAddress": "example_emailAddress",  # str
    "firstName": "example_firstName",  # str
    "lastName": "example_lastName",  # str
    "postalCode": "example_postalCode",  # str
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Mandate().load({"id": "mandate_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MandateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MerchantEntity

```python
merchant = client.Merchant()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountCurrencies` | `list` | No | The list of currencies that the merchant has accounts for. |
| `canHaveTrustAccounts` | `bool` | No | Trust accounts are a special type of account that allow the account name to be trusted for use in statements and verification of payee checks. |
| `cardPaymentProcessor` | `str` | No | Name of the card payment processor. |
| `companyID` | `str` | No | The Company ID recorded in the Compliance system. |
| `displayQrOnHostedPay` | `bool` | No | Indicates if a QR Code containing the payment link should be displayed on the hosted payment page. |
| `hostedPayVersion` | `int` | No | The version of the hosted payment page to use with the merchant. |
| `id` | `str` | No | Unique ID for the merchant. |
| `inserted` | `str` | No | Timestamp the merchant was added to MoneyMoov. |
| `isBlocked` | `bool` | No | The merchant is blocked from making payments (payouts). |
| `isExited` | `bool` | No | The merchant has formally terminated their relationship and is no longer a customer. |
| `isSuspended` | `bool` | No | The merchant has temporarily suspended their own account. |
| `jurisdiction` | `str` | No | The jurisdiction the merchant entity is incorporated or established in. |
| `logoUrlPng` | `str` | No | The CDN URL of the merchant's logo in PNG format. |
| `logoUrlSvg` | `str` | No | The CDN URL of the merchant's logo in SVG format. |
| `merchantCategoryCode` | `str` | No | The industry code that represents the merchant's primary trading activity. |
| `name` | `str` | No | The registered business name of the merchant. |
| `notes` | `str` | No | The notes field is an optional free text field that can be used to store any additional information about the merchant. |
| `parentMerchant` | `dict` | No |  |
| `paymentAccountLimit` | `int` | No | The maximum number of payment accounts that can be created for the Merchant. |
| `paymentAccounts` | `list` | No |  |
| `reason` | `str` | No | The reason for the suspension. |
| `shortName` | `str` | No | A URL friendly shortish name for the merchant. |
| `supportedPaymentMethodsList` | `list` | No | The payment methods that are configured and supported for this merchant. |
| `suspensionReason` | `str` | No | The reason for the suspension, provided by the merchant. |
| `tags` | `list` | No | An optional list of descriptive tags that can be used on merchant entities such as payment requests. |
| `timeZoneId` | `str` | No | The IANA (Internet Assigned Numbers Authority) time zone identifier of the merchant. |
| `tradingName` | `str` | No | An optional trading name. |
| `webHookLimit` | `int` | No | The maximum number of web hooks that can be created for the Merchant. |
| `yourRoleName` | `str` | No | The name of the role for the identity that loaded the merchant record. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Merchant().list()
for merchant in results:
    print(merchant)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Merchant().load({"id": "merchant_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Merchant().remove({"id": "merchant_id", "user_id": "user_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Merchant().update({
    "id": "merchant_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MerchantEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MerchantAuthorisationSettingEntity

```python
merchant_authorisation_setting = client.MerchantAuthorisationSetting()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amountLower` | `float` | No |  |
| `amountUpper` | `float` | No |  |
| `authorisationType` | `str` | No |  |
| `beneficiariesOnly` | `bool` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `lastEditorCantAuthorise` | `bool` | No |  |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `numberOfAuthorisers` | `int` | No |  |
| `roleSettings` | `list` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.MerchantAuthorisationSetting().list({"merchant_id": "example"})
for merchant_authorisation_setting in results:
    print(merchant_authorisation_setting)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MerchantAuthorisationSettingEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MerchantDirectDebitMandatePageEntity

```python
merchant_direct_debit_mandate_page = client.MerchantDirectDebitMandatePage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approvedAt` | `str` | No | Date at which the supplier approved this mandate. |
| `currency` | `str` | No | Currency of this mandate. |
| `customerAccountNumber` | `str` | No | Customer's account number in case of GBP account. |
| `customerCity` | `str` | No | Customer's city of residence. |
| `customerCountryCode` | `str` | No | Customer's country of residence code. |
| `customerCountryName` | `str` | No | Customer's country of residence. |
| `customerEmailAddress` | `str` | No | Customer's email address. |
| `customerFirstName` | `str` | No | Customer's first name. |
| `customerIban` | `str` | No | Customer's IBAN in case of EUR account. |
| `customerLastName` | `str` | No | Customer's last name. |
| `customerSortCode` | `str` | No | Customer's sort code in case of GBP account. |
| `id` | `str` | No | Internal ID of the mandate. |
| `inserted` | `str` | No | The timestamp this mandate was created at. |
| `isRecurring` | `bool` | No | Whether this mandate is single-use or recurring. |
| `lastUpdated` | `str` | No | The timestamp this mandate was last updated at. |
| `merchantID` | `str` | No | Internal ID of this mandate's merchant. |
| `reference` | `str` | No | Reference assigned to this mandate. |
| `status` | `str` | No | General status of this mandate. |
| `supplierBankAccountID` | `str` | No | ID that the supplier assigned to this mandate's bank account. |
| `supplierCustomerID` | `str` | No | ID that the supplier assigned to this mandate's customer. |
| `supplierMandateID` | `str` | No | ID that the supplier assigned to this mandate. |
| `supplierName` | `str` | No | Name of the supplier used to create this mandate. |
| `supplierStatus` | `str` | No | Last status that the supplier reported for this mandate. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.MerchantDirectDebitMandatePage().list()
for merchant_direct_debit_mandate_page in results:
    print(merchant_direct_debit_mandate_page)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MerchantDirectDebitMandatePageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MerchantPayByBankSettingEntity

```python
merchant_pay_by_bank_setting = client.MerchantPayByBankSetting()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bankCountryCodes` | `list` | No | The list of country codes representing the banks the country supports. |
| `bankID` | `str` | No | ID of the bank to be configured for the merchant. |
| `bankName` | `str` | No | Name of the Bank/Institution. |
| `businessInstitutionID` | `str` | No | ID that the processor uses to identify the bank (business accounts). |
| `currency` | `str` | No | Currency supported by the bank. |
| `logo` | `str` | No | URL of the bank's logo. |
| `message` | `str` | No | Message relating to specific bank. |
| `messageImageUrl` | `str` | No | Optional image URL to be displayed with the message. |
| `order` | `int` | No | Order in which this setting will appear in the UI. |
| `personalInstitutionID` | `str` | No | ID that the processor uses to identify the bank (personal accounts). |
| `processor` | `str` | No | Name of the bank payment processor. |
| `warningHeading` | `str` | No | The heading for a warning message related to the bank institution to be displayed to the user. |
| `warningMessage` | `str` | No | The warning message related to the bank institution to be displayed to the user. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.MerchantPayByBankSetting().list({"merchant_id": "example"})
for merchant_pay_by_bank_setting in results:
    print(merchant_pay_by_bank_setting)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MerchantPayByBankSettingEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MerchantPaymentRequestTemplateEntity

```python
merchant_payment_request_template = client.MerchantPaymentRequestTemplate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bankPaymentOptions` | `dict` | No |  |
| `cardPaymentAddressOptions` | `dict` | No |  |
| `cardPaymentCaptureOptions` | `dict` | No |  |
| `customFields` | `list` | No | A list of custom fields that can be included in the payment request template. |
| `defaultFields` | `list` | No | A list of default fields that are included in the payment request template. |
| `description` | `str` | Yes |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `name` | `str` | Yes |  |
| `notificationOptions` | `dict` | No |  |
| `paymentMethods` | `dict` | No |  |
| `paymentTerms` | `dict` | No |  |
| `priorityBankOptions` | `dict` | No |  |
| `template` | `dict` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.MerchantPaymentRequestTemplate().list({"merchant_id": "example"})
for merchant_payment_request_template in results:
    print(merchant_payment_request_template)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.MerchantPaymentRequestTemplate().load({"id": "merchant_payment_request_template_id", "paymentrequest_id": "paymentrequest_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.MerchantPaymentRequestTemplate().remove({"id": "merchant_payment_request_template_id", "paymentrequest_id": "paymentrequest_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.MerchantPaymentRequestTemplate().update({
    "id": "merchant_payment_request_template_id",
    "paymentrequest_id": "paymentrequest_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MerchantPaymentRequestTemplateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MerchantTokenEntity

```python
merchant_token = client.MerchantToken()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authenticationMethods` | `list` | No | A list of authentication types allowed to authorise the merchant token. |
| `authorisations` | `list` | No | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `int` | No | The number of distinct authorisers that have authorised the merchant token. |
| `authorisersRequiredCount` | `int` | No | The number of authorisers required for this merchant token. |
| `canAuthorise` | `bool` | No | True if the merchant token can be authorised by the user who loaded it. |
| `description` | `str` | No | Token description |
| `expiresAt` | `str` | No | Optional. |
| `hasCurrentUserAuthorised` | `bool` | No | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `hmacAlgorithm` | `str` | No | Optional shared secret algorithm to use for HMAC authentication. |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `ipAddressWhitelist` | `str` | No | Optional. |
| `isArchived` | `bool` | No | Indicates whether the merchant token is archived. |
| `isEnabled` | `bool` | No | If set to false the merchant token will not be accepted to authorise a request. |
| `lastAuthorised` | `str` | No |  |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No | The merchant id to add to the token |
| `nonce` | `str` | Yes |  |
| `permissionTypes` | `list` | No | The permissions that the merchant token supports. |
| `requestSignatureVersion` | `int` | No | Represent the version of the overall merchant token. |
| `sharedSecretAlgorithm` | `str` | No | Optional shared secret algorithm to use for HMAC authentication. |
| `sharedSecretBase64` | `str` | No | The base 64 encoded shared secret that is used for request authentication with an HMAC. |
| `token` | `str` | No | The JWT merchant token. |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.MerchantToken().create({
    "nonce": "example_nonce",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.MerchantToken().list({"merchant_id": "example"})
for merchant_token in results:
    print(merchant_token)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.MerchantToken().load({"id": "merchant_token_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.MerchantToken().update({
    "id": "merchant_token_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MerchantTokenEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MetadataEntity

```python
metadata = client.Metadata()
```

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Metadata().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MetadataEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionVersionEntity

```python
no_frixion_version = client.NoFrixionVersion()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `buildVersion` | `int` | No |  |
| `majorVersion` | `int` | No |  |
| `minorVersion` | `int` | No |  |
| `releaseName` | `str` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionVersion().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `NoFrixionVersionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## OpenBankingEntity

```python
open_banking = client.OpenBanking()
```

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.OpenBanking().create({
    "account_id": "example_account_id",  # str
})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.OpenBanking().remove({"account_id": "account_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `OpenBankingEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PayeeverificationEntity

```python
payeeverification = client.Payeeverification()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `str` | Yes | The name of the account to verify |
| `accountNumber` | `str` | No | The account number of the account to verify (for CoP checks) |
| `iban` | `str` | Yes | The IBAN of the account to verify (for VoP checks) |
| `payeeVerifiedAccountName` | `str` | No | The verified account name of the payee, if available (in case of a close match) |
| `result` | `str` | No | The result of the payee verification |
| `secondaryIdentification` | `str` | No | Optional secondary identifier for the account to verify. |
| `sortCode` | `str` | No | The sort code of the account to verify (for CoP checks) |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Payeeverification().create({
    "accountName": "example_accountName",  # str
    "iban": "example_iban",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PayeeverificationEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PaymentEntity

```python
payment = client.Payment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresses` | `list` | No |  |
| `amount` | `float` | No | The amount of money to request. |
| `amountPending` | `float` | No | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `float` | No | Total amount received for this payment request. |
| `amountRefunded` | `float` | No | Total amount refunded for this payment request. |
| `autoSendReceipt` | `bool` | No | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `str` | No | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `str` | No | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `bool` | No | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `bool` | No | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `str` | No | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `bool` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardNoPayerAuthentication` | `bool` | No | If set to true for card payments no attempt will be made to use payer authentication (3-D Secure and equivalent). |
| `cardProcessorMerchantID` | `str` | No | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `str` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `str` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `cardTransmitRawDetails` | `bool` | No | If set to true for card payments the sensitive card number and card verification number will be transmitted directly rather than being tokenised. |
| `createdByUser` | `dict` | Yes |  |
| `currency` | `str` | No | The currency of the request. |
| `customFields` | `list` | No | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `str` | No | Optional email address for the customer. |
| `customerID` | `str` | No | An optional customer identifier for the payment request. |
| `customerName` | `str` | No |  |
| `description` | `str` | No | An optional description for the payment request. |
| `destinationAccount` | `dict` | No |  |
| `directDebitPayment` | `dict` | No | Contains information about a Direct Debit payment attempt for a payment request. |
| `dueDate` | `str` | No | The due date for the payment request. |
| `events` | `list` | No |  |
| `failureCallbackUrl` | `str` | No | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `list` | No | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `str` | No |  |
| `hostedPayCheckoutUrl` | `str` | No | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `str` | No |  |
| `ignoreAddressVerification` | `bool` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `str` | No | The timestamp the payment request was created at. |
| `insertedSortable` | `str` | No | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `isArchived` | `bool` | No | Indicates whether the payment request is archived. |
| `jwk` | `str` | No | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `str` | No | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `str` | No | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `str` | No | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `str` | No | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `str` | No | The ID of the merchant to create the payment request for. |
| `merchantTokenDescription` | `str` | No | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `str` | No |  |
| `notificationRoleIDs` | `list` | No | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `str` | No | An optional order ID for the payment request. |
| `partialPaymentMethod` | `str` | No | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `str` | No | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `list` | No | The payment attempts made against this payment request. |
| `paymentMethods` | `list` | No | The payment methods that the payment request supports. |
| `paymentProcessor` | `str` | No | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `payrunID` | `str` | No | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `str` | No | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `str` | No | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `dict` | No |  |
| `sandboxSettleDelayInSeconds` | `int` | No | Sandbox only. |
| `shippingAddress` | `dict` | No |  |
| `shippingAddressCity` | `str` | No | Optionally the city of the customer's shipping address. |
| `shippingAddressCountryCode` | `str` | No | Optionally the country code of the customer's shipping address. |
| `shippingAddressCounty` | `str` | No | Optionally the state or county of the customer's shipping address. |
| `shippingAddressLine1` | `str` | No | Optionally the first line of the customer's shipping address. |
| `shippingAddressLine2` | `str` | No | Optionally the second line of the customer's shipping address. |
| `shippingAddressPostCode` | `str` | No | Optionally the post code of the customer's shipping address. |
| `shippingEmail` | `str` | No | Optionally the shipping email address for the customer. |
| `shippingFirstName` | `str` | No | Optionally the first name of the customer's shipping address. |
| `shippingLastName` | `str` | No | Optionally the last name of the customer's shipping address. |
| `shippingPhone` | `str` | No | Optionally the shipping phone number for the customer. |
| `status` | `str` | No | The current status of the payment request. |
| `successWebHookUrl` | `str` | No | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tagIds` | `list` | No | An optional list of tag ids to add to the payment request |
| `tags` | `list` | No | An optional list of descriptive tags attached to the payment request. |
| `title` | `str` | No | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `list` | No |  |
| `transactions` | `list` | No |  |
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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Payment().create({
    "createdByUser": {},  # dict
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Payment().load({"id": "payment_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Payment().update({
    "id": "payment_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PaymentAccountEntity

```python
payment_account = client.PaymentAccount()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `str` | No | Name for the account |
| `accountSupplierName` | `str` | No | The payment account supplier name. |
| `availableBalance` | `float` | No | The current available balance of the account. |
| `availableBalanceMinorUnits` | `int` | No | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `float` | No | Balance of the account. |
| `balanceMinorUnits` | `int` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `str` | No | The bank name for external accounts |
| `consentID` | `str` | No | The ID of the consent used to connect the external account. |
| `createdBy` | `dict` | Yes |  |
| `createdByDisplayName` | `str` | No | Either the name of the user, merchant token or api key that created the account |
| `currency` | `str` | No | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `str` | No | Indicates the default payment rail for this account. |
| `displayName` | `str` | No | Gets a unique display name for the payment account. |
| `expiryDate` | `str` | No | The date that the external account will expire |
| `externalAccountIcon` | `str` | No | The Icon for external accounts |
| `id` | `str` | No | Unique id for the account. |
| `identifier` | `dict` | Yes |  |
| `inserted` | `str` | No | Timestamp when the account was created. |
| `isArchived` | `bool` | No | Indicates whether the account is archived. |
| `isConnectedAccount` | `bool` | No | Indicates if the payment account is an externally connected account. |
| `isDefault` | `bool` | No | Is the default account |
| `isTrustAccount` | `bool` | No | Indicates if the payment account is a trust account. |
| `isVirtual` | `bool` | No | True if the account is a virtual account. |
| `lastTransaction` | `dict` | No |  |
| `lastUpdated` | `str` | No | Timestamp when the account was last updated. |
| `merchantID` | `str` | No | The ID of the merchant that owns the account. |
| `merchantName` | `str` | No | The name of the merchant that owns the account. |
| `physicalAccountID` | `str` | No | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `list` | No | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `float` | No | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `int` | No | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `str` | No | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `str` | No | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `str` | No | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `str` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `str` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `str` | No |  |
| `xeroBankFeedSyncStatus` | `str` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | No | Indicates the number of unsynchronised transactions with Xero |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PaymentAccount().list()
for payment_account in results:
    print(payment_account)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentAccountEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PaymentAccountMinimalEntity

```python
payment_account_minimal = client.PaymentAccountMinimal()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `str` | No | Name for the account |
| `availableBalance` | `float` | No | The current available balance of the account. |
| `balance` | `float` | No | Balance of the account. |
| `balanceMinorUnits` | `int` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `currency` | `str` | No | Currency of the account in ISO 4217 format |
| `id` | `str` | No | Unique id for the account. |
| `identifier` | `dict` | Yes |  |
| `isArchived` | `bool` | No | Is the account archived |
| `isConnectedAccount` | `bool` | No | Indicates if the payment account is an externally connected account. |
| `merchantID` | `str` | No | The ID of the merchant that owns the account. |
| `submittedPayoutsBalance` | `float` | No | Total of the payouts that have been submitted for processing. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PaymentAccountMinimal().list()
for payment_account_minimal in results:
    print(payment_account_minimal)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentAccountMinimalEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PaymentInitiationEntity

```python
payment_initiation = client.PaymentInitiation()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `paymentInitiationID` | `str` | No | The unique identifier of the payment initiation request. |
| `paymentRequestCallbackUrl` | `str` | No | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `str` | No |  |
| `redirectUrl` | `str` | No | A redirect URL for the user to authorise the payment initiation request at the ASPSP |
| `responseType` | `str` | No |  |
| `specificErrorMessage` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PaymentInitiation().create({
    "paymentrequest_id": "example_paymentrequest_id",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentInitiationEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PaymentRequestEntity

```python
payment_request = client.PaymentRequest()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresses` | `list` | No |  |
| `amount` | `float` | No | The amount of money to request. |
| `amountPending` | `float` | No | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `float` | No | Total amount received for this payment request. |
| `amountRefunded` | `float` | No | Total amount refunded for this payment request. |
| `autoSendReceipt` | `bool` | No | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `str` | No | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `str` | No | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `bool` | No | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `bool` | No | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `str` | No | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `bool` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardProcessorMerchantID` | `str` | No | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `str` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `str` | No | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `createdByUser` | `dict` | Yes |  |
| `currency` | `str` | No | The currency of the request. |
| `customFields` | `list` | No | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `str` | No | Optional email address for the customer. |
| `customerID` | `str` | No | An optional customer identifier for the payment request. |
| `customerName` | `str` | No |  |
| `description` | `str` | No | An optional description for the payment request. |
| `destinationAccount` | `dict` | No |  |
| `directDebitPayment` | `dict` | No | Contains information about a Direct Debit payment attempt for a payment request. |
| `doSimulateSettlementFailure` | `bool` | No |  |
| `dueDate` | `str` | No | The due date for the payment request. |
| `errorDescription` | `str` | No |  |
| `events` | `list` | No |  |
| `failedPaymentRequests` | `dict` | No |  |
| `failureCallbackUrl` | `str` | No | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `list` | No | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `str` | No |  |
| `hostedPayCheckoutUrl` | `str` | No | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `str` | No |  |
| `ignoreAddressVerification` | `bool` | No | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `str` | No | The timestamp the payment request was created at. |
| `insertedSortable` | `str` | No | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `institution` | `str` | No |  |
| `isArchived` | `bool` | No | Indicates whether the payment request is archived. |
| `jwk` | `str` | No | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `str` | No | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `str` | No | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `str` | No | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `str` | No | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `str` | No |  |
| `merchantTokenDescription` | `str` | No | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `str` | No |  |
| `notificationRoleIDs` | `list` | No | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `str` | No | An optional order ID for the payment request. |
| `partialPaymentMethod` | `str` | No | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `str` | No | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `list` | No | The payment attempts made against this payment request. |
| `paymentInitiationID` | `str` | No |  |
| `paymentMethods` | `list` | No | The payment methods that the payment request supports. |
| `paymentProcessor` | `str` | No | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `paymentRequests` | `list` | No |  |
| `payrunID` | `str` | No | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `str` | No | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `str` | No | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `dict` | No |  |
| `sandboxSettleDelayInSeconds` | `int` | No | Sandbox only. |
| `shippingAddress` | `dict` | No |  |
| `status` | `str` | No | The current status of the payment request. |
| `successWebHookUrl` | `str` | No | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tags` | `list` | No | An optional list of descriptive tags attached to the payment request. |
| `title` | `str` | No | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `list` | No |  |
| `transactions` | `list` | No |  |
| `useHostedPaymentPage` | `bool` | No | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PaymentRequest().create({
    "createdByUser": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PaymentRequest().list()
for payment_request in results:
    print(payment_request)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.PaymentRequest().load({"id": "payment_request_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.PaymentRequest().remove({"id": "payment_request_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.PaymentRequest().update({
    "id": "payment_request_id",
    "paymentrequest_id": "paymentrequest_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentRequestEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PaymentRequestEventEntity

```python
payment_request_event = client.PaymentRequestEvent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float` | Yes |  |
| `applePayTransactionID` | `str` | No | Transaction ID received in Apple pay token. |
| `cardAuthorizationResponseID` | `str` | No | For a successful card authorization this field will hold the response ID. |
| `cardExpiryMonth` | `int` | No | For card payment events this field holds the payer's card expiry month. |
| `cardExpiryYear` | `int` | No | For card payment events this field holds the payer's card expiry year. |
| `cardIssuer` | `str` | No | For card payment events this field holds the payer's card issuer. |
| `cardIssuerCountry` | `str` | No | For card payment events this field holds the payer's card issuer country of origin. |
| `cardLastFourDigits` | `str` | No | For card payment events this field holds the payer's card last four digits. |
| `cardRequestID` | `str` | No |  |
| `cardScheme` | `str` | No | For card payment events this field holds the scheme of the payer's card, e.g. |
| `cardTokenCustomerID` | `str` | No | If the option to create a reusable token for card payments was set this field contains the token the merchant can store to use for repeat payments. |
| `cardTransactionID` | `str` | No |  |
| `currency` | `str` | No |  |
| `directDebitPaymentID` | `str` | No | Payment ID issued by the Direct Debit supplier. |
| `directDebitPaymentReference` | `str` | No | Reference string issued by the Direct Debit supplier. |
| `drirectDebitMandateID` | `str` | No | The ID of the mandate that was used wehn requesting payment. |
| `errorMessage` | `str` | No |  |
| `errorReason` | `str` | No |  |
| `eventType` | `str` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `lightningInvoice` | `str` | No | For Bitcoin Lightning payments this field holds the invoice presented to the payer. |
| `lightningRHash` | `str` | No | For Bitcoin Lightning payments the hash of the invoice presented to the payer. |
| `originUrl` | `str` | No | Optional field that can be set by payment methods, such as pay by bank, that may want to redirect back to the URL that initiated the attempt in the case of a failure condition. |
| `paymentMethodType` | `str` | No | The type of payment method the event relates to, e.g. |
| `paymentProcessorName` | `str` | No | If the event was for a card payment this is the name of the card processor, e.g. |
| `paymentRequestID` | `str` | No |  |
| `pispBankStatus` | `str` | No | For payment initiation attempts some providers (e.g. |
| `pispPaymentInitiationID` | `str` | No | For a payment initiation this is the ID returned by the service provider initiating the payment for us. |
| `pispPaymentInstitutionName` | `str` | No | For a payment initiation this is the name of the financial institution that is used to initiate and authorise the payment. |
| `pispPaymentServiceProviderID` | `str` | No | For a payment initiation this is the service provider ID selected by the payer, typically the ID for the bank or similar financial institution. |
| `pispRedirectUrl` | `str` | No | For a payment initiation this is the redirect URL returned by the service provider initiating the payment for us. |
| `reconciledTransactionID` | `str` | No | For settlement events (only relevant for non-card payments) this is the payin transaction that the payment request event was reconciled with. |
| `refundPayoutID` | `str` | No | ID of the Payout that was created for refund. |
| `status` | `str` | No |  |
| `walletName` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PaymentRequestEvent().list({"paymentrequest_id": "example"})
for payment_request_event in results:
    print(payment_request_event)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentRequestEventEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PaymentRequestMetricEntity

```python
payment_request_metric = client.PaymentRequestMetric()
```

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.PaymentRequestMetric().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentRequestMetricEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PaymentRequestMinimalEntity

```python
payment_request_minimal = client.PaymentRequestMinimal()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float` | No | The amount of money to request. |
| `amountPending` | `float` | No | The amount of money that was authorised but has not arrived in the account yet. |
| `amountReceived` | `float` | No | The amount of money that has been received for this payment request. |
| `amountRefunded` | `float` | No | The amount of money that has been refunded for this payment request. |
| `callbackUrl` | `str` | No |  |
| `cardStripePaymentIntentSecret` | `str` | No |  |
| `countryCode` | `str` | No | The country code associated with the payment. |
| `currency` | `str` | No | The currency of the request. |
| `customFieldsToDisplay` | `list` | No | Custom fields to display to the customer. |
| `description` | `str` | No | An optional description for the payment request. |
| `dueDate` | `str` | No | The due date of the payment request. |
| `fieldDisplaySettings` | `list` | No |  |
| `googlePayMerchantID` | `str` | No | Merchant ID from Google Pay |
| `id` | `str` | No |  |
| `jwk` | `str` | No | The jwk containing the public key |
| `merchantID` | `str` | No |  |
| `merchantLogoUrlPng` | `str` | No |  |
| `merchantLogoUrlSvg` | `str` | No |  |
| `merchantName` | `str` | No |  |
| `merchantShortName` | `str` | No |  |
| `partialPaymentMethod` | `str` | No |  |
| `paymentAttempts` | `list` | No | The payment attempts for this payment request. |
| `paymentMethodsList` | `list` | No | The payment methods that the payment request supports. |
| `paymentProcessor` | `str` | No | The card processor |
| `paymentProcessorKey` | `str` | No | The card processors public key |
| `pispError` | `str` | No | This is the error returned from the bank which is recorded in payment request events. |
| `priorityBankID` | `str` | No |  |
| `status` | `str` | No | The status of the payment request. |
| `stripeAccountID` | `str` | No | Account ID of connected customers in Stripe |
| `title` | `str` | No | The title of the payment request. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PaymentRequestMinimal().list({"paymentrequest_id": "example"})
for payment_request_minimal in results:
    print(payment_request_minimal)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentRequestMinimalEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PaymentRequestResultEntity

```python
payment_request_result = client.PaymentRequestResult()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float` | No | The authorised payment amount. |
| `amountPending` | `float` | No |  |
| `amountReceived` | `float` | No |  |
| `amountRefunded` | `float` | No |  |
| `currency` | `str` | No | The authorised payment currency. |
| `customerID` | `str` | No | The customer id |
| `paymentRequestID` | `str` | No | The ID of the payment request the result is for. |
| `payments` | `list` | No | The list of payment attempts that have been received for the payment request. |
| `pispAuthorizations` | `list` | No |  |
| `requestedAmount` | `float` | No | The full original payment amount requested. |
| `result` | `str` | No | The result of the payment attempt. |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PaymentRequestResult().list({"paymentrequest_id": "example"})
for payment_request_result in results:
    print(payment_request_result)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaymentRequestResultEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PayoutEntity

```python
payout = client.Payout()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountID` | `str` | No | Gets or Sets Account Id of sending account |
| `allowIncomplete` | `bool` | No | If set to true the payout will get created even if the business validation rules fail. |
| `amount` | `float` | No | Gets or Sets payout amount |
| `amountMinorUnits` | `int` | No | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `str` | No | This field is used when returning an payout record to a client. |
| `approverID` | `str` | No | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `list` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `list` | No | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `int` | No | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `int` | No | The number of authorisers required for this payout. |
| `batchPayoutID` | `str` | No | The ID of the batch the payout is associated with. |
| `beneficiary` | `dict` | Yes |  |
| `beneficiaryID` | `str` | No | Optional. |
| `canAuthorise` | `bool` | No | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `bool` | No | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `bool` | No | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `str` | No | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `str` | No |  |
| `createdByEmailAddress` | `str` | No |  |
| `currency` | `str` | No | Gets or Sets Currency of payout request |
| `currentUserID` | `str` | No | The ID of the user that requested access to the PayOut record. |
| `description` | `str` | No | Gets or Sets description of payout request |
| `destination` | `dict` | No |  |
| `documents` | `list` | No | Documents associated with the payout. |
| `events` | `list` | No | The activity associated with the payout. |
| `failedPayouts` | `dict` | No |  |
| `formattedAmount` | `str` | No | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `str` | No | FX destination currency and amount formatted string. |
| `formattedSchedule` | `str` | No |  |
| `formattedScheduleDayOnly` | `str` | No |  |
| `formattedSourceAccountAvailableBalance` | `str` | No | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `float` | No | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `int` | No | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `str` | No | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `str` | No | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `str` | No | Optional. |
| `fxRate` | `float` | No | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `bool` | No | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `bool` | No | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `str` | No | The ID for the payout. |
| `inserted` | `str` | No |  |
| `invoiceID` | `str` | No | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `bool` | No | Indicates whether the payout is archived. |
| `isFailed` | `bool` | No | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `bool` | No | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `bool` | No | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `str` | No |  |
| `nonce` | `str` | No |  |
| `paymentProcessor` | `str` | No | The usptream payment processor for the payout. |
| `paymentRail` | `str` | No | Optional field to indicate the payment rail to use for the payout. |
| `payouts` | `list` | No |  |
| `payrunID` | `str` | No | The ID of the payrun that this payout is associated with. |
| `payrunName` | `str` | No | The name of the payrun that this payout is associated with. |
| `reason` | `str` | No |  |
| `rule` | `dict` | No |  |
| `scheduleDate` | `str` | No | The date the payout should be submitted. |
| `scheduled` | `bool` | No | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `float` | No | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `int` | No | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `str` | No | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `str` | No | The currency of the source account. |
| `sourceAccountIban` | `str` | No | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `dict` | Yes |  |
| `sourceAccountName` | `str` | No | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `str` | No | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `str` | No | The sort code of the account the payout is being made from. |
| `status` | `str` | No | Gets or Sets the status of payout request |
| `tagIds` | `list` | No | An optional list of tag ids to add to the payout. |
| `tags` | `list` | No | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `str` | No | Gets or Sets destination reference ID |
| `topupPayrunID` | `str` | No | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `float` | No | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `float` | No | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `float` | No | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `str` | No | Gets or Sets payout type |
| `userID` | `str` | No | Gets or Sets User ID of who created the payout request |
| `yourReference` | `str` | No | Gets or Sets your reference ID |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Payout().create({
    "id": "example_id",  # str
    "beneficiary": {},  # dict
    "sourceAccountIdentifier": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Payout().list()
for payout in results:
    print(payout)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Payout().load({"id": "payout_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Payout().remove({"id": "payout_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Payout().update({
    "id": "payout_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PayoutEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PayoutKeysetPageEntity

```python
payout_keyset_page = client.PayoutKeysetPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountID` | `str` | No | Gets or Sets Account Id of sending account |
| `amount` | `float` | No | Gets or Sets payout amount |
| `amountMinorUnits` | `int` | No | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `str` | No | This field is used when returning an payout record to a client. |
| `approverID` | `str` | No | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `list` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `list` | No | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `int` | No | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `int` | No | The number of authorisers required for this payout. |
| `batchPayoutID` | `str` | No | The ID of the batch the payout is associated with. |
| `beneficiary` | `dict` | Yes |  |
| `canAuthorise` | `bool` | No | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `bool` | No | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `bool` | No | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `str` | No | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `str` | No |  |
| `createdByEmailAddress` | `str` | No |  |
| `currency` | `str` | No | Gets or Sets Currency of payout request |
| `currentUserID` | `str` | No | The ID of the user that requested access to the PayOut record. |
| `description` | `str` | No | Gets or Sets description of payout request |
| `destination` | `dict` | No |  |
| `documents` | `list` | No | Documents associated with the payout. |
| `events` | `list` | No | The activity associated with the payout. |
| `formattedAmount` | `str` | No | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `str` | No | FX destination currency and amount formatted string. |
| `formattedSchedule` | `str` | No |  |
| `formattedScheduleDayOnly` | `str` | No |  |
| `formattedSourceAccountAvailableBalance` | `str` | No | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `float` | No | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `int` | No | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `str` | No | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `str` | No | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `str` | No | Optional. |
| `fxRate` | `float` | No | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `bool` | No | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `bool` | No | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `str` | No | The ID for the payout. |
| `inserted` | `str` | No |  |
| `invoiceID` | `str` | No | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `bool` | No | Indicates whether the payout is archived. |
| `isFailed` | `bool` | No | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `bool` | No | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `bool` | No | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `str` | No |  |
| `nonce` | `str` | No |  |
| `paymentProcessor` | `str` | No | The usptream payment processor for the payout. |
| `paymentRail` | `str` | No | Optional field to indicate the payment rail to use for the payout. |
| `payrunID` | `str` | No | The ID of the payrun that this payout is associated with. |
| `payrunName` | `str` | No | The name of the payrun that this payout is associated with. |
| `rule` | `dict` | No |  |
| `scheduleDate` | `str` | No | The date the payout should be submitted. |
| `scheduled` | `bool` | No | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `float` | No | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `int` | No | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `str` | No | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `str` | No | The currency of the source account. |
| `sourceAccountIban` | `str` | No | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `dict` | Yes |  |
| `sourceAccountName` | `str` | No | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `str` | No | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `str` | No | The sort code of the account the payout is being made from. |
| `status` | `str` | No | Gets or Sets the status of payout request |
| `tags` | `list` | No | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `str` | No | Gets or Sets destination reference ID |
| `topupPayrunID` | `str` | No | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `float` | No | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `float` | No | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `float` | No | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `str` | No | Gets or Sets payout type |
| `userID` | `str` | No | Gets or Sets User ID of who created the payout request |
| `yourReference` | `str` | No | Gets or Sets your reference ID |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PayoutKeysetPage().list({"merchant_id": "example"})
for payout_keyset_page in results:
    print(payout_keyset_page)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PayoutKeysetPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PayoutMetricEntity

```python
payout_metric = client.PayoutMetric()
```

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.PayoutMetric().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PayoutMetricEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PayrunEntity

```python
payrun = client.Payrun()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisationDate` | `str` | No |  |
| `authorisations` | `list` | No | A list of the users who have successfully authorised the latest version of the payrun and when. |
| `authorisersCompletedCount` | `int` | No | The number of distinct authorisers that have authorised the payrun. |
| `authorisersRequiredCount` | `int` | No | The number of authorisers required for this payrun. |
| `batchPayoutID` | `str` | No |  |
| `canAuthorise` | `bool` | No | True if the payrun can be authorised by the user who loaded it. |
| `canDelete` | `bool` | No |  |
| `canEdit` | `bool` | No |  |
| `events` | `list` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No | True if the payrun was loaded for a user and that user has already authorised the latest version of the payrun. |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `invoices` | `list` | No |  |
| `invoicesMinimal` | `list` | No |  |
| `isArchived` | `bool` | No |  |
| `lastUpdated` | `str` | No |  |
| `lastUpdatedBy` | `dict` | Yes |  |
| `merchantID` | `str` | No |  |
| `name` | `str` | No |  |
| `nonce` | `str` | No |  |
| `notes` | `str` | No |  |
| `payments` | `list` | No |  |
| `payouts` | `list` | No |  |
| `payoutsCount` | `int` | No |  |
| `reason` | `str` | No |  |
| `scheduleDate` | `str` | No |  |
| `scheduledDate` | `str` | No |  |
| `sourceAccounts` | `list` | No |  |
| `status` | `str` | No |  |
| `totalEur` | `float` | No |  |
| `totalGbp` | `float` | No |  |
| `totalUsd` | `float` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Payrun().create({
    "id": "example_id",  # str
    "lastUpdatedBy": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Payrun().list()
for payrun in results:
    print(payrun)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Payrun().load({"id": "payrun_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Payrun().remove({"id": "payrun_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Payrun().update({
    "id": "payrun_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PayrunEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ReportEntity

```python
report = client.Report()
```

### Operations

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Report().update({
    "id": "id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ReportEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ReportResultEntity

```python
report_result = client.ReportResult()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `contentType` | `str` | No |  |
| `contents` | `str` | No |  |
| `lastCompletedAt` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `reportName` | `str` | No |  |
| `reportType` | `str` | No |  |
| `statementNumber` | `int` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.ReportResult().load({"id": 1, "report_id": "report_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ReportResultEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RoleEntity

```python
role = client.Role()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failedRoles` | `dict` | No |  |
| `roles` | `list` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Role().create({
    "merchant_id": "example_merchant_id",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RoleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RuleEntity

```python
rule = client.Rule()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account` | `dict` | No |  |
| `accountID` | `str` | No | The ID of the account the rule will apply to. |
| `approveUrl` | `str` | No | If set this property holds the URL an approver needs to visit in order to complete a strong authentication check in order to approve the rule. |
| `approverID` | `str` | No |  |
| `authenticationMethods` | `list` | No | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `list` | No | A list of the users who have successfully authorised the latest version of the rule and when. |
| `authorisersCompletedCount` | `int` | No | The number of distinct authorisers that have authorised the rule. |
| `authorisersRequiredCount` | `int` | No | The number of authorisers required for this rule. |
| `canAuthorise` | `bool` | No | True if the rule can be authorised by the user who loaded it. |
| `createdBy` | `dict` | Yes |  |
| `description` | `str` | No | Arbitrary description for the rule. |
| `endAt` | `str` | No | Optional end time for rule executions. |
| `hasCurrentUserAuthorised` | `bool` | No | True if the current user has authorised. |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `isDisabled` | `bool` | No | If set to true the rule will be disabled from executing. |
| `lastExecutedAt` | `str` | No |  |
| `lastRunAtTransactionDate` | `str` | No | The most recent transaction date when the rule was last run. |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No | The ID of the merchant that owns the account. |
| `name` | `str` | No | A name to succinctly describe the rule. |
| `nonce` | `str` | Yes |  |
| `onApprovedWebHookUrl` | `str` | No | Optional URL to receive an HTTP request with the rule details when the rule status changes to approved. |
| `onExecutionErrorWebHookUrl` | `str` | No | Optional URL to receive an HTTP request when a rule execution attempt fails. |
| `onExecutionSuccessWebHookUrl` | `str` | No | Optional URL to receive an HTTP request when a rule execution attempt succeeds. |
| `startAt` | `str` | No | Optional start time for rule executions. |
| `status` | `str` | No |  |
| `sweepAction` | `dict` | No |  |
| `timeZoneId` | `str` | No | If the rule should be executed on a recurring schedule this is the timezone that the CRON expression should be evaluated in. |
| `triggerCronExpression` | `str` | No | If the rule should be executed on a recurring schedule this is the expression that sets the schedule. |
| `triggerOnPayIn` | `bool` | No | Set to true if the rule execution should be triggered when the account receives a pay in (credit). |
| `userID` | `str` | No |  |
| `webHookSecret` | `str` | No | If set this secret will be used to sign Web Hook requests. |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Rule().create({
    "createdBy": {},  # dict
    "nonce": "example_nonce",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Rule().list()
for rule in results:
    print(rule)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Rule().load({"id": "rule_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Rule().remove({"id": "rule_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Rule().update({
    "id": "rule_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RuleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RuleEventEntity

```python
rule_event = client.RuleEvent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `errorMessage` | `str` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `isAuthoriseToEnable` | `bool` | No |  |
| `message` | `str` | No |  |
| `rawResponse` | `str` | No |  |
| `ruleEventType` | `str` | No |  |
| `ruleID` | `str` | No |  |
| `user` | `dict` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.RuleEvent().list({"id": "example"})
for rule_event in results:
    print(rule_event)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `RuleEventEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TagEntity

```python
tag = client.Tag()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `colourHex` | `str` | No |  |
| `description` | `str` | No |  |
| `id` | `str` | No |  |
| `merchantID` | `str` | Yes |  |
| `name` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Tag().create({
    "merchant_id": "example_merchant_id",  # str
    "merchantID": "example_merchantID",  # str
    "name": "example_name",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Tag().list({"merchant_id": "example"})
for tag in results:
    print(tag)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TagEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TokenEntity

```python
token = client.Token()
```

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Token().create({
    "id": "example_id",  # str
})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Token().remove({"id": "id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TokenEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TransactionEntity

```python
transaction = client.Transaction()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountID` | `str` | No | The ID of the account the transaction belongs to. |
| `accountName` | `str` | No | The name of the account the transaction belongs to. |
| `accountSequenceNumber` | `int` | No | The sequence number of transaction on a per account basis. |
| `addressDetails` | `dict` | No |  |
| `amount` | `float` | No | Amount of the transaction. |
| `amountMinorUnits` | `int` | No | Amount of the transaction expressed in the currency’s minor units (e.g. |
| `balance` | `float` | No | Balance left on the account after the transaction. |
| `balanceMinorUnits` | `int` | No | Balance on the account expressed in the currency’s minor units (e.g. |
| `bookingDateTime` | `str` | No |  |
| `chargeDetails` | `dict` | No |  |
| `content` | `list` | No |  |
| `counterparty` | `dict` | No |  |
| `counterpartySummary` | `str` | No | For pay in (credit) transactions this will contain a descriptive string with the most important fields about the counterparty. |
| `currency` | `str` | No | Currency of transaction. |
| `currencyExchange` | `dict` | No | Provides details on the currency exchange. |
| `date` | `str` | No |  |
| `description` | `str` | No | Description of the transaction. |
| `enrichment` | `dict` | No |  |
| `fxAmount` | `float` | No | For an FX payout this is the amound in the FX currency. |
| `fxCurrency` | `str` | No | For an FX payout this is the currency that was received or that was instructed. |
| `fxRate` | `float` | No | For an FX payout this is the exchange rate between the transaction currency and the FX currency. |
| `grossAmount` | `dict` | Yes |  |
| `id` | `str` | No | Unique ID for the transaction. |
| `inserted` | `str` | No | Date when the transaction was inserted into the ledger. |
| `isoBankTransactionCode` | `dict` | No |  |
| `merchant` | `dict` | No |  |
| `merchantID` | `str` | No | The ID of the merchant that owns the account. |
| `pageNumber` | `int` | No | Current page number. |
| `pageSize` | `int` | No | Page size |
| `payeeDetails` | `dict` | Yes | The Payee object contains details of the beneficiary, person or business. |
| `payerDetails` | `dict` | Yes |  |
| `paymentRequestCustomFields` | `dict` | No | The custom fields that were attached to the payment request that resulted in this transaction. |
| `paymentRequestID` | `str` | No | For Pay by Bank and Direct Debit transactions this will contain the ID of the payment request. |
| `payoutID` | `str` | No | ID of the payout that resulted in the transaction. |
| `proprietaryBankTransactionCode` | `dict` | No |  |
| `rawReference` | `str` | No | The raw payment reference details as received from the payment processor. |
| `reference` | `str` | No |  |
| `ruleID` | `str` | No | ID of the rule that resulted in the transaction. |
| `statementReferences` | `list` | No |  |
| `status` | `str` | No |  |
| `supplementaryData` | `Any` | No |  |
| `tags` | `list` | No | An optional list of descriptive tags attached to the transaction. |
| `theirReference` | `str` | No | For a pay out the reference that the payer attached for the receiving party. |
| `totalPages` | `int` | No | Total pages |
| `totalSize` | `int` | No | Total count |
| `transactionAmount` | `dict` | Yes |  |
| `transactionDate` | `str` | No | Date when the transaction occurred. |
| `transactionInformation` | `list` | No |  |
| `transactionMutability` | `str` | No |  |
| `type` | `str` | No | Type of the transaction. |
| `valueDateTime` | `str` | No |  |
| `virtualIBAN` | `str` | No | If set it indicates the payin was to a virtual IBAN. |
| `yourReference` | `str` | No | For a pay in the reference the sending party attached. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Transaction().create({
    "id": "example_id",  # str
    "grossAmount": {},  # dict
    "payeeDetails": {},  # dict
    "payerDetails": {},  # dict
    "transactionAmount": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Transaction().list({"account_id": "example", "id": "example_id"})
for transaction in results:
    print(transaction)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Transaction().load({"id": "transaction_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Transaction().remove({"id": "transaction_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TransactionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UserEntity

```python
user = client.User()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `clientSessionTimeouts` | `list` | No | The number of seconds a session for this user should last before expiring. |
| `emailAddress` | `str` | Yes |  |
| `firstName` | `str` | Yes |  |
| `id` | `str` | No |  |
| `lastName` | `str` | Yes |  |
| `passkeyAdded` | `bool` | No |  |
| `permissions` | `dict` | No |  |
| `profile` | `str` | No |  |
| `rolesWithScope` | `list` | No |  |
| `twoFactorEnabled` | `bool` | No |  |
| `userInviteID` | `str` | No | Optional ID of the invite that is being accepted so the user can be assigned a role on a new merchant. |

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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.User().list()
for user in results:
    print(user)
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.User().update({
    "id": "id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UserEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UserInviteEntity

```python
user_invite = client.UserInvite()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisationStatus` | `dict` | No |  |
| `failedUserInvites` | `dict` | No |  |
| `id` | `str` | No |  |
| `initialRoleID` | `str` | No | The role ID to automatically assign to the merchant’s very first user. |
| `inviteeEmailAddress` | `str` | No | Email address of the user being invited. |
| `inviteeFirstName` | `str` | No | First Name of the user being invited. |
| `inviteeLastName` | `str` | No | Last Name of the user being invited. |
| `inviterEmailAddress` | `str` | No |  |
| `inviterFirstName` | `str` | No |  |
| `inviterLastName` | `str` | No |  |
| `isAuthorised` | `bool` | No | Will be set to true once the invite has met the authorisation requirements. |
| `isInviteeRegistered` | `bool` | No | If true, indicates the invitee's email address corresponds to an existing MoneyMoov user. |
| `lastInvited` | `str` | No |  |
| `merchantID` | `str` | No | ID of the merchant the user is being invited to. |
| `merchantName` | `str` | No |  |
| `message` | `str` | No |  |
| `registrationUrl` | `str` | No |  |
| `sendInviteEmail` | `bool` | No | If set to true an email will be sent to the invitee with instructions on how to accept the invite. |
| `status` | `str` | No |  |
| `user` | `dict` | Yes |  |
| `userID` | `str` | No |  |
| `userInvites` | `list` | No |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.UserInvite().create({
    "id": "example_id",  # str
    "user": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.UserInvite().list({"merchant_id": "example"})
for user_invite in results:
    print(user_invite)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.UserInvite().load({"id": "user_invite_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.UserInvite().remove({"id": "user_invite_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.UserInvite().update({
    "id": "user_invite_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UserInviteEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## VirtualEntity

```python
virtual = client.Virtual()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `str` | No | Name for the account |
| `accountSupplierName` | `str` | No | The payment account supplier name. |
| `availableBalance` | `float` | No | The current available balance of the account. |
| `availableBalanceMinorUnits` | `int` | No | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `float` | No | Balance of the account. |
| `balanceMinorUnits` | `int` | No | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `str` | No | The bank name for external accounts |
| `consentID` | `str` | No | The ID of the consent used to connect the external account. |
| `createdBy` | `dict` | Yes |  |
| `createdByDisplayName` | `str` | No | Either the name of the user, merchant token or api key that created the account |
| `currency` | `str` | No | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `str` | No | Indicates the default payment rail for this account. |
| `displayName` | `str` | No | Gets a unique display name for the payment account. |
| `expiryDate` | `str` | No | The date that the external account will expire |
| `externalAccountIcon` | `str` | No | The Icon for external accounts |
| `id` | `str` | No | Unique id for the account. |
| `identifier` | `dict` | Yes |  |
| `inserted` | `str` | No | Timestamp when the account was created. |
| `isArchived` | `bool` | No | Indicates whether the account is archived. |
| `isConnectedAccount` | `bool` | No | Indicates if the payment account is an externally connected account. |
| `isDefault` | `bool` | No | Is the default account |
| `isTrustAccount` | `bool` | No | Indicates if the payment account is a trust account. |
| `isVirtual` | `bool` | No | True if the account is a virtual account. |
| `lastTransaction` | `dict` | No |  |
| `lastUpdated` | `str` | No | Timestamp when the account was last updated. |
| `merchantID` | `str` | No | The ID of the merchant that owns the account. |
| `merchantName` | `str` | No | The name of the merchant that owns the account. |
| `name` | `str` | Yes | The name of the virtual account. |
| `physicalAccountID` | `str` | No | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `list` | No | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `float` | No | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `int` | No | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `str` | No | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `str` | No | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `str` | No | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `str` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `str` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `str` | No |  |
| `xeroBankFeedSyncStatus` | `str` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | No | Indicates the number of unsynchronised transactions with Xero |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Virtual().create({
    "account_id": "example_account_id",  # str
    "createdBy": {},  # dict
    "identifier": {},  # dict
    "name": "example_name",  # str
})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Virtual().update({
    "account_id": "account_id",
    "id": "id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `VirtualEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## WebhookEntity

```python
webhook = client.Webhook()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destinationUrl` | `str` | No | The destination URL for the webhook. |
| `emailAddress` | `str` | No | The recipient email address(es) for notifications. |
| `failedNotificationEmailAddress` | `str` | No | The email address to which notifications about failed webhook deliveries will be sent. |
| `id` | `str` | No |  |
| `isActive` | `bool` | No |  |
| `merchantID` | `str` | No | The ID of the merchant that the webhook is for. |
| `notificationMethod` | `str` | No | The type of notification that will be sent. |
| `resourceTypes` | `list` | No | The resource types that the webhook will be generated for. |
| `retry` | `bool` | No |  |
| `secret` | `str` | No | The secret key required to authenticate webhook notifications. |
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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Webhook().create({
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Webhook().list({"merchant_id": "example"})
for webhook in results:
    print(webhook)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Webhook().load({"id": "webhook_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Webhook().remove({"id": "webhook_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Webhook().update({
    "id": "webhook_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `WebhookEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = NofrixionSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

