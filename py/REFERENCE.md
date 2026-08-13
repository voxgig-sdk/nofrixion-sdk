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
| `accountBalances` | `list` | No |  |
| `accountID` | `str` | No |  |
| `accountIdentifications` | `list` | No |  |
| `accountName` | `str` | No |  |
| `accountNames` | `list` | No |  |
| `accountSupplierName` | `str` | No |  |
| `accountType` | `str` | No |  |
| `availableBalance` | `float` | No |  |
| `availableBalanceMinorUnits` | `int` | No |  |
| `balance` | `float` | No |  |
| `balanceMinorUnits` | `int` | No |  |
| `bankName` | `str` | No |  |
| `consentID` | `str` | No |  |
| `consolidatedAccountInformation` | `dict` | No |  |
| `createdBy` | `dict` | Yes |  |
| `createdByDisplayName` | `str` | No |  |
| `currency` | `str` | No |  |
| `defaultPaymentRail` | `str` | No |  |
| `description` | `str` | No |  |
| `details` | `str` | No |  |
| `displayName` | `str` | No |  |
| `expiryDate` | `str` | No |  |
| `externalAccountIcon` | `str` | No |  |
| `format` | `str` | No |  |
| `fromDate` | `str` | No |  |
| `id` | `str` | No |  |
| `identifier` | `dict` | Yes |  |
| `inserted` | `str` | No |  |
| `isArchived` | `bool` | No |  |
| `isConnectedAccount` | `bool` | No |  |
| `isDefault` | `bool` | No |  |
| `isTrustAccount` | `bool` | No |  |
| `isVirtual` | `bool` | No |  |
| `lastTransaction` | `dict` | No |  |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `merchantName` | `str` | No |  |
| `nickname` | `str` | No |  |
| `physicalAccountID` | `str` | No |  |
| `roleIDs` | `list` | No |  |
| `rules` | `list` | No |  |
| `submittedPayoutsBalance` | `float` | No |  |
| `submittedPayoutsBalanceMinorUnits` | `int` | No |  |
| `summary` | `str` | No |  |
| `supplierPhysicalAccountID` | `str` | No |  |
| `supplierSepaInstantStatus` | `str` | No |  |
| `toDate` | `str` | No |  |
| `type` | `str` | No |  |
| `usageType` | `str` | No |  |
| `xeroBankFeedConnectionStatus` | `str` | No |  |
| `xeroBankFeedLastSyncedAt` | `str` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `str` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `str` | No |  |
| `xeroBankFeedSyncStatus` | `str` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Account().create({
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
| `approveUrl` | `str` | No |  |
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
| `authenticationMethods` | `list` | No |  |
| `authorisations` | `list` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `beneficiaries` | `list` | No |  |
| `beneficiaryEvents` | `list` | No |  |
| `canAuthorise` | `bool` | No |  |
| `canUpdate` | `bool` | No |  |
| `createdBy` | `dict` | Yes |  |
| `createdByEmailAddress` | `str` | No |  |
| `currency` | `str` | Yes |  |
| `destination` | `dict` | No |  |
| `failedBeneficiaries` | `dict` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `isEnabled` | `bool` | No |  |
| `lastAuthorised` | `str` | No |  |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `name` | `str` | Yes |  |
| `nonce` | `str` | No |  |
| `sourceAccountIDs` | `list` | No |  |
| `sourceAccounts` | `list` | No |  |
| `theirReference` | `str` | No |  |

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
| `groupMembers` | `list` | No |  |
| `groupName` | `str` | Yes |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | Yes |  |

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
| `isPayerAuthenticationRequired` | `bool` | No |  |
| `isSoftDecline` | `bool` | No |  |
| `payerAuthenticationAccessToken` | `str` | No |  |
| `payerAuthenticationMerchantData` | `str` | No |  |
| `payerAuthenticationUrl` | `str` | No |  |
| `payerAuthenticationWindowHeight` | `int` | No |  |
| `payerAuthenticationWindowWidth` | `int` | No |  |
| `paymentRequestCallbackUrl` | `str` | No |  |
| `paymentRequestID` | `str` | No |  |
| `requestID` | `str` | No |  |
| `responseCode` | `str` | No |  |
| `responseType` | `str` | No |  |
| `status` | `str` | No |  |
| `threeDSRedirectUrl` | `str` | No |  |
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
| `cardType` | `str` | No |  |
| `customerEmailAddress` | `str` | No |  |
| `expiryMonth` | `str` | No |  |
| `expiryYear` | `str` | No |  |
| `id` | `str` | No |  |
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
result = client.CardCustomerToken().remove()
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
| `isPayerAuthenticationRequired` | `bool` | No |  |
| `isSoftDecline` | `bool` | No |  |
| `payerAuthenticationAccessToken` | `str` | No |  |
| `payerAuthenticationMerchantData` | `str` | No |  |
| `payerAuthenticationUrl` | `str` | No |  |
| `payerAuthenticationWindowHeight` | `int` | No |  |
| `payerAuthenticationWindowWidth` | `int` | No |  |
| `paymentRequestCallbackUrl` | `str` | No |  |
| `paymentRequestID` | `str` | No |  |
| `requestID` | `str` | No |  |
| `responseCode` | `str` | No |  |
| `responseType` | `str` | No |  |
| `status` | `str` | No |  |
| `threeDSRedirectUrl` | `str` | No |  |
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
| `authorisationUrl` | `str` | No |  |
| `callbackUrl` | `str` | No |  |
| `consentID` | `str` | No |  |
| `emailAddress` | `str` | No |  |
| `expiryDate` | `str` | No |  |
| `failureCallbackUrl` | `str` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `institutionID` | `str` | No |  |
| `isConnectedAccounts` | `bool` | No |  |
| `isEnabled` | `bool` | No |  |
| `merchantID` | `str` | No |  |
| `provider` | `str` | No |  |
| `successWebHookUrl` | `str` | No |  |

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
| `failedSubmissions` | `dict` | No |  |
| `successfulSubmissions` | `list` | No |  |

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
| `exchangeRate` | `float` | No |  |
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
| `accountNumber` | `str` | No |  |
| `addressLine1` | `str` | Yes |  |
| `addressLine2` | `str` | No |  |
| `approvedAt` | `str` | No |  |
| `city` | `str` | Yes |  |
| `countryCode` | `str` | Yes |  |
| `currency` | `str` | No |  |
| `customerAccountNumber` | `str` | No |  |
| `customerCity` | `str` | No |  |
| `customerCountryCode` | `str` | No |  |
| `customerCountryName` | `str` | No |  |
| `customerEmailAddress` | `str` | No |  |
| `customerFirstName` | `str` | No |  |
| `customerIban` | `str` | No |  |
| `customerLastName` | `str` | No |  |
| `customerSortCode` | `str` | No |  |
| `emailAddress` | `str` | Yes |  |
| `firstName` | `str` | Yes |  |
| `iban` | `str` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `isRecurring` | `bool` | No |  |
| `lastName` | `str` | Yes |  |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `postalCode` | `str` | Yes |  |
| `reference` | `str` | No |  |
| `sortCode` | `str` | No |  |
| `status` | `str` | No |  |
| `supplierBankAccountID` | `str` | No |  |
| `supplierCustomerID` | `str` | No |  |
| `supplierMandateID` | `str` | No |  |
| `supplierName` | `str` | No |  |
| `supplierStatus` | `str` | No |  |

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
| `accountCurrencies` | `list` | No |  |
| `canHaveTrustAccounts` | `bool` | No |  |
| `cardPaymentProcessor` | `str` | No |  |
| `companyID` | `str` | No |  |
| `displayQrOnHostedPay` | `bool` | No |  |
| `hostedPayVersion` | `int` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `isBlocked` | `bool` | No |  |
| `isExited` | `bool` | No |  |
| `isSuspended` | `bool` | No |  |
| `jurisdiction` | `str` | No |  |
| `logoUrlPng` | `str` | No |  |
| `logoUrlSvg` | `str` | No |  |
| `merchantCategoryCode` | `str` | No |  |
| `name` | `str` | No |  |
| `notes` | `str` | No |  |
| `parentMerchant` | `dict` | No |  |
| `paymentAccountLimit` | `int` | No |  |
| `paymentAccounts` | `list` | No |  |
| `reason` | `str` | No |  |
| `shortName` | `str` | No |  |
| `supportedPaymentMethodsList` | `list` | No |  |
| `suspensionReason` | `str` | No |  |
| `tags` | `list` | No |  |
| `timeZoneId` | `str` | No |  |
| `tradingName` | `str` | No |  |
| `webHookLimit` | `int` | No |  |
| `yourRoleName` | `str` | No |  |

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
result = client.Merchant().remove({"id": "merchant_id"})
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
| `approvedAt` | `str` | No |  |
| `currency` | `str` | No |  |
| `customerAccountNumber` | `str` | No |  |
| `customerCity` | `str` | No |  |
| `customerCountryCode` | `str` | No |  |
| `customerCountryName` | `str` | No |  |
| `customerEmailAddress` | `str` | No |  |
| `customerFirstName` | `str` | No |  |
| `customerIban` | `str` | No |  |
| `customerLastName` | `str` | No |  |
| `customerSortCode` | `str` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `isRecurring` | `bool` | No |  |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `reference` | `str` | No |  |
| `status` | `str` | No |  |
| `supplierBankAccountID` | `str` | No |  |
| `supplierCustomerID` | `str` | No |  |
| `supplierMandateID` | `str` | No |  |
| `supplierName` | `str` | No |  |
| `supplierStatus` | `str` | No |  |

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
| `bankCountryCodes` | `list` | No |  |
| `bankID` | `str` | No |  |
| `bankName` | `str` | No |  |
| `businessInstitutionID` | `str` | No |  |
| `currency` | `str` | No |  |
| `logo` | `str` | No |  |
| `message` | `str` | No |  |
| `messageImageUrl` | `str` | No |  |
| `order` | `int` | No |  |
| `personalInstitutionID` | `str` | No |  |
| `processor` | `str` | No |  |
| `warningHeading` | `str` | No |  |
| `warningMessage` | `str` | No |  |

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
| `customFields` | `list` | No |  |
| `defaultFields` | `list` | No |  |
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
| `authenticationMethods` | `list` | No |  |
| `authorisations` | `list` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `canAuthorise` | `bool` | No |  |
| `description` | `str` | No |  |
| `expiresAt` | `str` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No |  |
| `hmacAlgorithm` | `str` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `ipAddressWhitelist` | `str` | No |  |
| `isArchived` | `bool` | No |  |
| `isEnabled` | `bool` | No |  |
| `lastAuthorised` | `str` | No |  |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `nonce` | `str` | Yes |  |
| `permissionTypes` | `list` | No |  |
| `requestSignatureVersion` | `int` | No |  |
| `sharedSecretAlgorithm` | `str` | No |  |
| `sharedSecretBase64` | `str` | No |  |
| `token` | `str` | No |  |

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
result = client.OpenBanking().remove()
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
| `accountName` | `str` | Yes |  |
| `accountNumber` | `str` | No |  |
| `iban` | `str` | Yes |  |
| `payeeVerifiedAccountName` | `str` | No |  |
| `result` | `str` | No |  |
| `secondaryIdentification` | `str` | No |  |
| `sortCode` | `str` | No |  |

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
| `amount` | `float` | No |  |
| `amountPending` | `float` | No |  |
| `amountReceived` | `float` | No |  |
| `amountRefunded` | `float` | No |  |
| `autoSendReceipt` | `bool` | No |  |
| `baseOriginUrl` | `str` | No |  |
| `callbackUrl` | `str` | No |  |
| `cardAuthorizeOnly` | `bool` | No |  |
| `cardCreateToken` | `bool` | No |  |
| `cardCreateTokenMode` | `str` | No |  |
| `cardIgnoreCVN` | `bool` | No |  |
| `cardNoPayerAuthentication` | `bool` | No |  |
| `cardProcessorMerchantID` | `str` | No |  |
| `cardStripePaymentIntentID` | `str` | No |  |
| `cardStripePaymentIntentSecret` | `str` | No |  |
| `cardTransmitRawDetails` | `bool` | No |  |
| `createdByUser` | `dict` | Yes |  |
| `currency` | `str` | No |  |
| `customFields` | `list` | No |  |
| `customerEmailAddress` | `str` | No |  |
| `customerID` | `str` | No |  |
| `customerName` | `str` | No |  |
| `description` | `str` | No |  |
| `destinationAccount` | `dict` | No |  |
| `directDebitPayment` | `dict` | No |  |
| `dueDate` | `str` | No |  |
| `events` | `list` | No |  |
| `failureCallbackUrl` | `str` | No |  |
| `fieldDisplaySettings` | `list` | No |  |
| `formattedAmount` | `str` | No |  |
| `hostedPayCheckoutUrl` | `str` | No |  |
| `id` | `str` | No |  |
| `ignoreAddressVerification` | `bool` | No |  |
| `inserted` | `str` | No |  |
| `insertedSortable` | `str` | No |  |
| `isArchived` | `bool` | No |  |
| `jwk` | `str` | No |  |
| `lastUpdated` | `str` | No |  |
| `lightningInvoice` | `str` | No |  |
| `lightningInvoiceExpiresAt` | `str` | No |  |
| `merchantDirectDebitMandateID` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `merchantTokenDescription` | `str` | No |  |
| `notificationEmailAddresses` | `str` | No |  |
| `notificationRoleIDs` | `list` | No |  |
| `orderID` | `str` | No |  |
| `partialPaymentMethod` | `str` | No |  |
| `partialPaymentSteps` | `str` | No |  |
| `paymentAttempts` | `list` | No |  |
| `paymentMethods` | `list` | No |  |
| `paymentProcessor` | `str` | No |  |
| `payrunID` | `str` | No |  |
| `pispAccountID` | `str` | No |  |
| `priorityBankID` | `str` | No |  |
| `result` | `dict` | No |  |
| `sandboxSettleDelayInSeconds` | `int` | No |  |
| `shippingAddress` | `dict` | No |  |
| `shippingAddressCity` | `str` | No |  |
| `shippingAddressCountryCode` | `str` | No |  |
| `shippingAddressCounty` | `str` | No |  |
| `shippingAddressLine1` | `str` | No |  |
| `shippingAddressLine2` | `str` | No |  |
| `shippingAddressPostCode` | `str` | No |  |
| `shippingEmail` | `str` | No |  |
| `shippingFirstName` | `str` | No |  |
| `shippingLastName` | `str` | No |  |
| `shippingPhone` | `str` | No |  |
| `status` | `str` | No |  |
| `successWebHookUrl` | `str` | No |  |
| `tagIds` | `list` | No |  |
| `tags` | `list` | No |  |
| `title` | `str` | No |  |
| `tokenisedCards` | `list` | No |  |
| `transactions` | `list` | No |  |
| `useHostedPaymentPage` | `bool` | No |  |

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
| `accountName` | `str` | No |  |
| `accountSupplierName` | `str` | No |  |
| `availableBalance` | `float` | No |  |
| `availableBalanceMinorUnits` | `int` | No |  |
| `balance` | `float` | No |  |
| `balanceMinorUnits` | `int` | No |  |
| `bankName` | `str` | No |  |
| `consentID` | `str` | No |  |
| `createdBy` | `dict` | Yes |  |
| `createdByDisplayName` | `str` | No |  |
| `currency` | `str` | No |  |
| `defaultPaymentRail` | `str` | No |  |
| `displayName` | `str` | No |  |
| `expiryDate` | `str` | No |  |
| `externalAccountIcon` | `str` | No |  |
| `id` | `str` | No |  |
| `identifier` | `dict` | Yes |  |
| `inserted` | `str` | No |  |
| `isArchived` | `bool` | No |  |
| `isConnectedAccount` | `bool` | No |  |
| `isDefault` | `bool` | No |  |
| `isTrustAccount` | `bool` | No |  |
| `isVirtual` | `bool` | No |  |
| `lastTransaction` | `dict` | No |  |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `merchantName` | `str` | No |  |
| `physicalAccountID` | `str` | No |  |
| `rules` | `list` | No |  |
| `submittedPayoutsBalance` | `float` | No |  |
| `submittedPayoutsBalanceMinorUnits` | `int` | No |  |
| `summary` | `str` | No |  |
| `supplierSepaInstantStatus` | `str` | No |  |
| `xeroBankFeedConnectionStatus` | `str` | No |  |
| `xeroBankFeedLastSyncedAt` | `str` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `str` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `str` | No |  |
| `xeroBankFeedSyncStatus` | `str` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | No |  |

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
| `accountName` | `str` | No |  |
| `availableBalance` | `float` | No |  |
| `balance` | `float` | No |  |
| `balanceMinorUnits` | `int` | No |  |
| `currency` | `str` | No |  |
| `id` | `str` | No |  |
| `identifier` | `dict` | Yes |  |
| `isArchived` | `bool` | No |  |
| `isConnectedAccount` | `bool` | No |  |
| `merchantID` | `str` | No |  |
| `submittedPayoutsBalance` | `float` | No |  |

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
| `paymentInitiationID` | `str` | No |  |
| `paymentRequestCallbackUrl` | `str` | No |  |
| `paymentRequestID` | `str` | No |  |
| `redirectUrl` | `str` | No |  |
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
| `amount` | `float` | No |  |
| `amountPending` | `float` | No |  |
| `amountReceived` | `float` | No |  |
| `amountRefunded` | `float` | No |  |
| `autoSendReceipt` | `bool` | No |  |
| `baseOriginUrl` | `str` | No |  |
| `callbackUrl` | `str` | No |  |
| `cardAuthorizeOnly` | `bool` | No |  |
| `cardCreateToken` | `bool` | No |  |
| `cardCreateTokenMode` | `str` | No |  |
| `cardIgnoreCVN` | `bool` | No |  |
| `cardProcessorMerchantID` | `str` | No |  |
| `cardStripePaymentIntentID` | `str` | No |  |
| `cardStripePaymentIntentSecret` | `str` | No |  |
| `createdByUser` | `dict` | Yes |  |
| `currency` | `str` | No |  |
| `customFields` | `list` | No |  |
| `customerEmailAddress` | `str` | No |  |
| `customerID` | `str` | No |  |
| `customerName` | `str` | No |  |
| `description` | `str` | No |  |
| `destinationAccount` | `dict` | No |  |
| `directDebitPayment` | `dict` | No |  |
| `doSimulateSettlementFailure` | `bool` | No |  |
| `dueDate` | `str` | No |  |
| `errorDescription` | `str` | No |  |
| `events` | `list` | No |  |
| `failedPaymentRequests` | `dict` | No |  |
| `failureCallbackUrl` | `str` | No |  |
| `fieldDisplaySettings` | `list` | No |  |
| `formattedAmount` | `str` | No |  |
| `hostedPayCheckoutUrl` | `str` | No |  |
| `id` | `str` | No |  |
| `ignoreAddressVerification` | `bool` | No |  |
| `inserted` | `str` | No |  |
| `insertedSortable` | `str` | No |  |
| `institution` | `str` | No |  |
| `isArchived` | `bool` | No |  |
| `jwk` | `str` | No |  |
| `lastUpdated` | `str` | No |  |
| `lightningInvoice` | `str` | No |  |
| `lightningInvoiceExpiresAt` | `str` | No |  |
| `merchantDirectDebitMandateID` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `merchantTokenDescription` | `str` | No |  |
| `notificationEmailAddresses` | `str` | No |  |
| `notificationRoleIDs` | `list` | No |  |
| `orderID` | `str` | No |  |
| `partialPaymentMethod` | `str` | No |  |
| `partialPaymentSteps` | `str` | No |  |
| `paymentAttempts` | `list` | No |  |
| `paymentInitiationID` | `str` | No |  |
| `paymentMethods` | `list` | No |  |
| `paymentProcessor` | `str` | No |  |
| `paymentRequests` | `list` | No |  |
| `payrunID` | `str` | No |  |
| `pispAccountID` | `str` | No |  |
| `priorityBankID` | `str` | No |  |
| `result` | `dict` | No |  |
| `sandboxSettleDelayInSeconds` | `int` | No |  |
| `shippingAddress` | `dict` | No |  |
| `status` | `str` | No |  |
| `successWebHookUrl` | `str` | No |  |
| `tags` | `list` | No |  |
| `title` | `str` | No |  |
| `tokenisedCards` | `list` | No |  |
| `transactions` | `list` | No |  |
| `useHostedPaymentPage` | `bool` | No |  |

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
result = client.PaymentRequest().load()
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.PaymentRequest().remove({"id": "id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.PaymentRequest().update({
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
| `applePayTransactionID` | `str` | No |  |
| `cardAuthorizationResponseID` | `str` | No |  |
| `cardExpiryMonth` | `int` | No |  |
| `cardExpiryYear` | `int` | No |  |
| `cardIssuer` | `str` | No |  |
| `cardIssuerCountry` | `str` | No |  |
| `cardLastFourDigits` | `str` | No |  |
| `cardRequestID` | `str` | No |  |
| `cardScheme` | `str` | No |  |
| `cardTokenCustomerID` | `str` | No |  |
| `cardTransactionID` | `str` | No |  |
| `currency` | `str` | No |  |
| `directDebitPaymentID` | `str` | No |  |
| `directDebitPaymentReference` | `str` | No |  |
| `drirectDebitMandateID` | `str` | No |  |
| `errorMessage` | `str` | No |  |
| `errorReason` | `str` | No |  |
| `eventType` | `str` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `lightningInvoice` | `str` | No |  |
| `lightningRHash` | `str` | No |  |
| `originUrl` | `str` | No |  |
| `paymentMethodType` | `str` | No |  |
| `paymentProcessorName` | `str` | No |  |
| `paymentRequestID` | `str` | No |  |
| `pispBankStatus` | `str` | No |  |
| `pispPaymentInitiationID` | `str` | No |  |
| `pispPaymentInstitutionName` | `str` | No |  |
| `pispPaymentServiceProviderID` | `str` | No |  |
| `pispRedirectUrl` | `str` | No |  |
| `reconciledTransactionID` | `str` | No |  |
| `refundPayoutID` | `str` | No |  |
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
| `amount` | `float` | No |  |
| `amountPending` | `float` | No |  |
| `amountReceived` | `float` | No |  |
| `amountRefunded` | `float` | No |  |
| `callbackUrl` | `str` | No |  |
| `cardStripePaymentIntentSecret` | `str` | No |  |
| `countryCode` | `str` | No |  |
| `currency` | `str` | No |  |
| `customFieldsToDisplay` | `list` | No |  |
| `description` | `str` | No |  |
| `dueDate` | `str` | No |  |
| `fieldDisplaySettings` | `list` | No |  |
| `googlePayMerchantID` | `str` | No |  |
| `id` | `str` | No |  |
| `jwk` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `merchantLogoUrlPng` | `str` | No |  |
| `merchantLogoUrlSvg` | `str` | No |  |
| `merchantName` | `str` | No |  |
| `merchantShortName` | `str` | No |  |
| `partialPaymentMethod` | `str` | No |  |
| `paymentAttempts` | `list` | No |  |
| `paymentMethodsList` | `list` | No |  |
| `paymentProcessor` | `str` | No |  |
| `paymentProcessorKey` | `str` | No |  |
| `pispError` | `str` | No |  |
| `priorityBankID` | `str` | No |  |
| `status` | `str` | No |  |
| `stripeAccountID` | `str` | No |  |
| `title` | `str` | No |  |

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
| `amount` | `float` | No |  |
| `amountPending` | `float` | No |  |
| `amountReceived` | `float` | No |  |
| `amountRefunded` | `float` | No |  |
| `currency` | `str` | No |  |
| `customerID` | `str` | No |  |
| `paymentRequestID` | `str` | No |  |
| `payments` | `list` | No |  |
| `pispAuthorizations` | `list` | No |  |
| `requestedAmount` | `float` | No |  |
| `result` | `str` | No |  |

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
| `accountID` | `str` | No |  |
| `allowIncomplete` | `bool` | No |  |
| `amount` | `float` | No |  |
| `amountMinorUnits` | `int` | No |  |
| `approvePayoutUrl` | `str` | No |  |
| `approverID` | `str` | No |  |
| `authenticationMethods` | `list` | No |  |
| `authorisations` | `list` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `batchPayoutID` | `str` | No |  |
| `beneficiary` | `dict` | Yes |  |
| `beneficiaryID` | `str` | No |  |
| `canAuthorise` | `bool` | No |  |
| `canProcess` | `bool` | No |  |
| `canUpdate` | `bool` | No |  |
| `chargeBearer` | `str` | No |  |
| `createdBy` | `str` | No |  |
| `createdByEmailAddress` | `str` | No |  |
| `currency` | `str` | No |  |
| `currentUserID` | `str` | No |  |
| `description` | `str` | No |  |
| `destination` | `dict` | No |  |
| `documents` | `list` | No |  |
| `events` | `list` | No |  |
| `failedPayouts` | `dict` | No |  |
| `formattedAmount` | `str` | No |  |
| `formattedFxDestinationAmount` | `str` | No |  |
| `formattedSchedule` | `str` | No |  |
| `formattedScheduleDayOnly` | `str` | No |  |
| `formattedSourceAccountAvailableBalance` | `str` | No |  |
| `fxDestinationAmount` | `float` | No |  |
| `fxDestinationAmountMinorUnits` | `int` | No |  |
| `fxDestinationCurrency` | `str` | No |  |
| `fxQuoteExpiresAt` | `str` | No |  |
| `fxQuoteID` | `str` | No |  |
| `fxRate` | `float` | No |  |
| `fxUseDestinationAmount` | `bool` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `invoiceID` | `str` | No |  |
| `isArchived` | `bool` | No |  |
| `isFailed` | `bool` | No |  |
| `isSettled` | `bool` | No |  |
| `isSubmitted` | `bool` | No |  |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `merchantTokenDescription` | `str` | No |  |
| `nonce` | `str` | No |  |
| `paymentProcessor` | `str` | No |  |
| `paymentRail` | `str` | No |  |
| `payouts` | `list` | No |  |
| `payrunID` | `str` | No |  |
| `payrunName` | `str` | No |  |
| `reason` | `str` | No |  |
| `rule` | `dict` | No |  |
| `scheduleDate` | `str` | No |  |
| `scheduled` | `bool` | No |  |
| `sourceAccountAvailableBalance` | `float` | No |  |
| `sourceAccountAvailableBalanceMinorUnits` | `int` | No |  |
| `sourceAccountBic` | `str` | No |  |
| `sourceAccountCurrency` | `str` | No |  |
| `sourceAccountIban` | `str` | No |  |
| `sourceAccountIdentifier` | `dict` | Yes |  |
| `sourceAccountName` | `str` | No |  |
| `sourceAccountNumber` | `str` | No |  |
| `sourceAccountSortcode` | `str` | No |  |
| `status` | `str` | No |  |
| `tagIds` | `list` | No |  |
| `tags` | `list` | No |  |
| `theirReference` | `str` | No |  |
| `topupPayrunID` | `str` | No |  |
| `transactedAmount` | `float` | No |  |
| `transactedFxAmount` | `float` | No |  |
| `transactedFxRate` | `float` | No |  |
| `type` | `str` | No |  |
| `userID` | `str` | No |  |
| `yourReference` | `str` | No |  |

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
| `accountID` | `str` | No |  |
| `amount` | `float` | No |  |
| `amountMinorUnits` | `int` | No |  |
| `approvePayoutUrl` | `str` | No |  |
| `approverID` | `str` | No |  |
| `authenticationMethods` | `list` | No |  |
| `authorisations` | `list` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `batchPayoutID` | `str` | No |  |
| `beneficiary` | `dict` | Yes |  |
| `canAuthorise` | `bool` | No |  |
| `canProcess` | `bool` | No |  |
| `canUpdate` | `bool` | No |  |
| `chargeBearer` | `str` | No |  |
| `createdBy` | `str` | No |  |
| `createdByEmailAddress` | `str` | No |  |
| `currency` | `str` | No |  |
| `currentUserID` | `str` | No |  |
| `description` | `str` | No |  |
| `destination` | `dict` | No |  |
| `documents` | `list` | No |  |
| `events` | `list` | No |  |
| `formattedAmount` | `str` | No |  |
| `formattedFxDestinationAmount` | `str` | No |  |
| `formattedSchedule` | `str` | No |  |
| `formattedScheduleDayOnly` | `str` | No |  |
| `formattedSourceAccountAvailableBalance` | `str` | No |  |
| `fxDestinationAmount` | `float` | No |  |
| `fxDestinationAmountMinorUnits` | `int` | No |  |
| `fxDestinationCurrency` | `str` | No |  |
| `fxQuoteExpiresAt` | `str` | No |  |
| `fxQuoteID` | `str` | No |  |
| `fxRate` | `float` | No |  |
| `fxUseDestinationAmount` | `bool` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `invoiceID` | `str` | No |  |
| `isArchived` | `bool` | No |  |
| `isFailed` | `bool` | No |  |
| `isSettled` | `bool` | No |  |
| `isSubmitted` | `bool` | No |  |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `merchantTokenDescription` | `str` | No |  |
| `nonce` | `str` | No |  |
| `paymentProcessor` | `str` | No |  |
| `paymentRail` | `str` | No |  |
| `payrunID` | `str` | No |  |
| `payrunName` | `str` | No |  |
| `rule` | `dict` | No |  |
| `scheduleDate` | `str` | No |  |
| `scheduled` | `bool` | No |  |
| `sourceAccountAvailableBalance` | `float` | No |  |
| `sourceAccountAvailableBalanceMinorUnits` | `int` | No |  |
| `sourceAccountBic` | `str` | No |  |
| `sourceAccountCurrency` | `str` | No |  |
| `sourceAccountIban` | `str` | No |  |
| `sourceAccountIdentifier` | `dict` | Yes |  |
| `sourceAccountName` | `str` | No |  |
| `sourceAccountNumber` | `str` | No |  |
| `sourceAccountSortcode` | `str` | No |  |
| `status` | `str` | No |  |
| `tags` | `list` | No |  |
| `theirReference` | `str` | No |  |
| `topupPayrunID` | `str` | No |  |
| `transactedAmount` | `float` | No |  |
| `transactedFxAmount` | `float` | No |  |
| `transactedFxRate` | `float` | No |  |
| `type` | `str` | No |  |
| `userID` | `str` | No |  |
| `yourReference` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PayoutKeysetPage().list()
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
| `authorisations` | `list` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `batchPayoutID` | `str` | No |  |
| `canAuthorise` | `bool` | No |  |
| `canDelete` | `bool` | No |  |
| `canEdit` | `bool` | No |  |
| `events` | `list` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No |  |
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
| `accountID` | `str` | No |  |
| `approveUrl` | `str` | No |  |
| `approverID` | `str` | No |  |
| `authenticationMethods` | `list` | No |  |
| `authorisations` | `list` | No |  |
| `authorisersCompletedCount` | `int` | No |  |
| `authorisersRequiredCount` | `int` | No |  |
| `canAuthorise` | `bool` | No |  |
| `createdBy` | `dict` | Yes |  |
| `description` | `str` | No |  |
| `endAt` | `str` | No |  |
| `hasCurrentUserAuthorised` | `bool` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `isDisabled` | `bool` | No |  |
| `lastExecutedAt` | `str` | No |  |
| `lastRunAtTransactionDate` | `str` | No |  |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `name` | `str` | No |  |
| `nonce` | `str` | Yes |  |
| `onApprovedWebHookUrl` | `str` | No |  |
| `onExecutionErrorWebHookUrl` | `str` | No |  |
| `onExecutionSuccessWebHookUrl` | `str` | No |  |
| `startAt` | `str` | No |  |
| `status` | `str` | No |  |
| `sweepAction` | `dict` | No |  |
| `timeZoneId` | `str` | No |  |
| `triggerCronExpression` | `str` | No |  |
| `triggerOnPayIn` | `bool` | No |  |
| `userID` | `str` | No |  |
| `webHookSecret` | `str` | No |  |

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
| `accountID` | `str` | No |  |
| `accountName` | `str` | No |  |
| `accountSequenceNumber` | `int` | No |  |
| `addressDetails` | `dict` | No |  |
| `amount` | `float` | No |  |
| `amountMinorUnits` | `int` | No |  |
| `balance` | `float` | No |  |
| `balanceMinorUnits` | `int` | No |  |
| `bookingDateTime` | `str` | No |  |
| `chargeDetails` | `dict` | No |  |
| `content` | `list` | No |  |
| `counterparty` | `dict` | No |  |
| `counterpartySummary` | `str` | No |  |
| `currency` | `str` | No |  |
| `currencyExchange` | `dict` | No |  |
| `date` | `str` | No |  |
| `description` | `str` | No |  |
| `enrichment` | `dict` | No |  |
| `fxAmount` | `float` | No |  |
| `fxCurrency` | `str` | No |  |
| `fxRate` | `float` | No |  |
| `grossAmount` | `dict` | Yes |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `isoBankTransactionCode` | `dict` | No |  |
| `merchant` | `dict` | No |  |
| `merchantID` | `str` | No |  |
| `pageNumber` | `int` | No |  |
| `pageSize` | `int` | No |  |
| `payeeDetails` | `dict` | Yes |  |
| `payerDetails` | `dict` | Yes |  |
| `paymentRequestCustomFields` | `dict` | No |  |
| `paymentRequestID` | `str` | No |  |
| `payoutID` | `str` | No |  |
| `proprietaryBankTransactionCode` | `dict` | No |  |
| `rawReference` | `str` | No |  |
| `reference` | `str` | No |  |
| `ruleID` | `str` | No |  |
| `statementReferences` | `list` | No |  |
| `status` | `str` | No |  |
| `supplementaryData` | `Any` | No |  |
| `tags` | `list` | No |  |
| `theirReference` | `str` | No |  |
| `totalPages` | `int` | No |  |
| `totalSize` | `int` | No |  |
| `transactionAmount` | `dict` | Yes |  |
| `transactionDate` | `str` | No |  |
| `transactionInformation` | `list` | No |  |
| `transactionMutability` | `str` | No |  |
| `type` | `str` | No |  |
| `valueDateTime` | `str` | No |  |
| `virtualIBAN` | `str` | No |  |
| `yourReference` | `str` | No |  |

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
results = client.Transaction().list()
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
| `clientSessionTimeouts` | `list` | No |  |
| `emailAddress` | `str` | Yes |  |
| `firstName` | `str` | Yes |  |
| `id` | `str` | No |  |
| `lastName` | `str` | Yes |  |
| `passkeyAdded` | `bool` | No |  |
| `permissions` | `dict` | No |  |
| `profile` | `str` | No |  |
| `rolesWithScope` | `list` | No |  |
| `twoFactorEnabled` | `bool` | No |  |
| `userInviteID` | `str` | No |  |

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
| `initialRoleID` | `str` | No |  |
| `inviteeEmailAddress` | `str` | No |  |
| `inviteeFirstName` | `str` | No |  |
| `inviteeLastName` | `str` | No |  |
| `inviterEmailAddress` | `str` | No |  |
| `inviterFirstName` | `str` | No |  |
| `inviterLastName` | `str` | No |  |
| `isAuthorised` | `bool` | No |  |
| `isInviteeRegistered` | `bool` | No |  |
| `lastInvited` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `merchantName` | `str` | No |  |
| `message` | `str` | No |  |
| `registrationUrl` | `str` | No |  |
| `sendInviteEmail` | `bool` | No |  |
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
| `accountName` | `str` | No |  |
| `accountSupplierName` | `str` | No |  |
| `availableBalance` | `float` | No |  |
| `availableBalanceMinorUnits` | `int` | No |  |
| `balance` | `float` | No |  |
| `balanceMinorUnits` | `int` | No |  |
| `bankName` | `str` | No |  |
| `consentID` | `str` | No |  |
| `createdBy` | `dict` | Yes |  |
| `createdByDisplayName` | `str` | No |  |
| `currency` | `str` | No |  |
| `defaultPaymentRail` | `str` | No |  |
| `displayName` | `str` | No |  |
| `expiryDate` | `str` | No |  |
| `externalAccountIcon` | `str` | No |  |
| `id` | `str` | No |  |
| `identifier` | `dict` | Yes |  |
| `inserted` | `str` | No |  |
| `isArchived` | `bool` | No |  |
| `isConnectedAccount` | `bool` | No |  |
| `isDefault` | `bool` | No |  |
| `isTrustAccount` | `bool` | No |  |
| `isVirtual` | `bool` | No |  |
| `lastTransaction` | `dict` | No |  |
| `lastUpdated` | `str` | No |  |
| `merchantID` | `str` | No |  |
| `merchantName` | `str` | No |  |
| `name` | `str` | Yes |  |
| `physicalAccountID` | `str` | No |  |
| `rules` | `list` | No |  |
| `submittedPayoutsBalance` | `float` | No |  |
| `submittedPayoutsBalanceMinorUnits` | `int` | No |  |
| `summary` | `str` | No |  |
| `supplierSepaInstantStatus` | `str` | No |  |
| `xeroBankFeedConnectionStatus` | `str` | No |  |
| `xeroBankFeedLastSyncedAt` | `str` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `str` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `str` | No |  |
| `xeroBankFeedSyncStatus` | `str` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | No |  |

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
| `destinationUrl` | `str` | No |  |
| `emailAddress` | `str` | No |  |
| `failedNotificationEmailAddress` | `str` | No |  |
| `id` | `str` | No |  |
| `isActive` | `bool` | No |  |
| `merchantID` | `str` | No |  |
| `notificationMethod` | `str` | No |  |
| `resourceTypes` | `list` | No |  |
| `retry` | `bool` | No |  |
| `secret` | `str` | No |  |
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

