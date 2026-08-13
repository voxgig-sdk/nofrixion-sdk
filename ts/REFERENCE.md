# Nofrixion TypeScript SDK Reference

Complete API reference for the Nofrixion TypeScript SDK.


## NofrixionSDK

### Constructor

```ts
new NofrixionSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NofrixionSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = NofrixionSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `NofrixionSDK` instance in test mode.


### Instance Methods

#### `Account(data?: object)`

Create a new `Account` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `AccountEntity` instance.

#### `Batch(data?: object)`

Create a new `Batch` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `BatchEntity` instance.

#### `Beneficiary(data?: object)`

Create a new `Beneficiary` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `BeneficiaryEntity` instance.

#### `BeneficiaryGroup(data?: object)`

Create a new `BeneficiaryGroup` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `BeneficiaryGroupEntity` instance.

#### `Card(data?: object)`

Create a new `Card` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CardEntity` instance.

#### `CardCustomerToken(data?: object)`

Create a new `CardCustomerToken` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CardCustomerTokenEntity` instance.

#### `CardPayment(data?: object)`

Create a new `CardPayment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CardPaymentEntity` instance.

#### `CardPublicKey(data?: object)`

Create a new `CardPublicKey` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CardPublicKeyEntity` instance.

#### `Consent(data?: object)`

Create a new `Consent` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ConsentEntity` instance.

#### `Currency(data?: object)`

Create a new `Currency` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `CurrencyEntity` instance.

#### `DirectDebitBatchSubmit(data?: object)`

Create a new `DirectDebitBatchSubmit` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `DirectDebitBatchSubmitEntity` instance.

#### `FxRate(data?: object)`

Create a new `FxRate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `FxRateEntity` instance.

#### `IPayment(data?: object)`

Create a new `IPayment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `IPaymentEntity` instance.

#### `Mandate(data?: object)`

Create a new `Mandate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MandateEntity` instance.

#### `Merchant(data?: object)`

Create a new `Merchant` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MerchantEntity` instance.

#### `MerchantAuthorisationSetting(data?: object)`

Create a new `MerchantAuthorisationSetting` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MerchantAuthorisationSettingEntity` instance.

#### `MerchantDirectDebitMandatePage(data?: object)`

Create a new `MerchantDirectDebitMandatePage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MerchantDirectDebitMandatePageEntity` instance.

#### `MerchantPayByBankSetting(data?: object)`

Create a new `MerchantPayByBankSetting` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MerchantPayByBankSettingEntity` instance.

#### `MerchantPaymentRequestTemplate(data?: object)`

Create a new `MerchantPaymentRequestTemplate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MerchantPaymentRequestTemplateEntity` instance.

#### `MerchantToken(data?: object)`

Create a new `MerchantToken` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MerchantTokenEntity` instance.

#### `Metadata(data?: object)`

Create a new `Metadata` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MetadataEntity` instance.

#### `NoFrixionVersion(data?: object)`

Create a new `NoFrixionVersion` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `NoFrixionVersionEntity` instance.

#### `OpenBanking(data?: object)`

Create a new `OpenBanking` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `OpenBankingEntity` instance.

#### `Payeeverification(data?: object)`

Create a new `Payeeverification` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PayeeverificationEntity` instance.

#### `Payment(data?: object)`

Create a new `Payment` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PaymentEntity` instance.

#### `PaymentAccount(data?: object)`

Create a new `PaymentAccount` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PaymentAccountEntity` instance.

#### `PaymentAccountMinimal(data?: object)`

Create a new `PaymentAccountMinimal` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PaymentAccountMinimalEntity` instance.

#### `PaymentInitiation(data?: object)`

Create a new `PaymentInitiation` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PaymentInitiationEntity` instance.

#### `PaymentRequest(data?: object)`

Create a new `PaymentRequest` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PaymentRequestEntity` instance.

#### `PaymentRequestEvent(data?: object)`

Create a new `PaymentRequestEvent` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PaymentRequestEventEntity` instance.

#### `PaymentRequestMetric(data?: object)`

Create a new `PaymentRequestMetric` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PaymentRequestMetricEntity` instance.

#### `PaymentRequestMinimal(data?: object)`

Create a new `PaymentRequestMinimal` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PaymentRequestMinimalEntity` instance.

#### `PaymentRequestResult(data?: object)`

Create a new `PaymentRequestResult` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PaymentRequestResultEntity` instance.

#### `Payout(data?: object)`

Create a new `Payout` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PayoutEntity` instance.

#### `PayoutKeysetPage(data?: object)`

Create a new `PayoutKeysetPage` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PayoutKeysetPageEntity` instance.

#### `PayoutMetric(data?: object)`

Create a new `PayoutMetric` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PayoutMetricEntity` instance.

#### `Payrun(data?: object)`

Create a new `Payrun` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PayrunEntity` instance.

#### `Report(data?: object)`

Create a new `Report` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ReportEntity` instance.

#### `ReportResult(data?: object)`

Create a new `ReportResult` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ReportResultEntity` instance.

#### `Role(data?: object)`

Create a new `Role` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `RoleEntity` instance.

#### `Rule(data?: object)`

Create a new `Rule` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `RuleEntity` instance.

#### `RuleEvent(data?: object)`

Create a new `RuleEvent` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `RuleEventEntity` instance.

#### `Tag(data?: object)`

Create a new `Tag` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TagEntity` instance.

#### `Token(data?: object)`

Create a new `Token` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TokenEntity` instance.

#### `Transaction(data?: object)`

Create a new `Transaction` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TransactionEntity` instance.

#### `User(data?: object)`

Create a new `User` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `UserEntity` instance.

#### `UserInvite(data?: object)`

Create a new `UserInvite` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `UserInviteEntity` instance.

#### `Virtual(data?: object)`

Create a new `Virtual` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `VirtualEntity` instance.

#### `Webhook(data?: object)`

Create a new `Webhook` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `WebhookEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `NofrixionSDK.test()`.

**Returns:** `NofrixionSDK` instance in test mode.


---

## AccountEntity

```ts
const account = client.Account()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountBalances` | `any[]` | No |  |
| `accountID` | `string` | No |  |
| `accountIdentifications` | `any[]` | No |  |
| `accountName` | `string` | No |  |
| `accountNames` | `any[]` | No |  |
| `accountSupplierName` | `string` | No |  |
| `accountType` | `string` | No |  |
| `availableBalance` | `number` | No |  |
| `availableBalanceMinorUnits` | `number` | No |  |
| `balance` | `number` | No |  |
| `balanceMinorUnits` | `number` | No |  |
| `bankName` | `string` | No |  |
| `consentID` | `string` | No |  |
| `consolidatedAccountInformation` | `Record<string, any>` | No |  |
| `createdBy` | `Record<string, any>` | Yes |  |
| `createdByDisplayName` | `string` | No |  |
| `currency` | `string` | No |  |
| `defaultPaymentRail` | `string` | No |  |
| `description` | `string` | No |  |
| `details` | `string` | No |  |
| `displayName` | `string` | No |  |
| `expiryDate` | `string` | No |  |
| `externalAccountIcon` | `string` | No |  |
| `format` | `string` | No |  |
| `fromDate` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `Record<string, any>` | Yes |  |
| `inserted` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `isConnectedAccount` | `boolean` | No |  |
| `isDefault` | `boolean` | No |  |
| `isTrustAccount` | `boolean` | No |  |
| `isVirtual` | `boolean` | No |  |
| `lastTransaction` | `Record<string, any>` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `nickname` | `string` | No |  |
| `physicalAccountID` | `string` | No |  |
| `roleIDs` | `any[]` | No |  |
| `rules` | `any[]` | No |  |
| `submittedPayoutsBalance` | `number` | No |  |
| `submittedPayoutsBalanceMinorUnits` | `number` | No |  |
| `summary` | `string` | No |  |
| `supplierPhysicalAccountID` | `string` | No |  |
| `supplierSepaInstantStatus` | `string` | No |  |
| `toDate` | `string` | No |  |
| `type` | `string` | No |  |
| `usageType` | `string` | No |  |
| `xeroBankFeedConnectionStatus` | `string` | No |  |
| `xeroBankFeedLastSyncedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `string` | No |  |
| `xeroBankFeedSyncStatus` | `string` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `number` | No |  |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `statement` | `/api/v1/accounts/{accountID}/statements` | `client.Account().create({ $action: 'statement', ... })` |
| `export` | `/api/v1/accounts/export` | `client.Account().load({ $action: 'export', ... })` |
| `statement` | `/api/v1/accounts/statements` | `client.Account().load({ $action: 'statement', ... })` |
| `transaction_export` | `/api/v1/accounts/{accountID}/transactions/export` | `client.Account().load({ $action: 'transaction_export', ... })` |
| `statement` | `/api/v1/accounts/statements` | `client.Account().remove({ $action: 'statement', ... })` |

An action returns that action's OWN response, which is not necessarily a
Account record — check the API definition for its shape.

```ts
const result = await client.Account().create({
  $action: 'statement',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Account().create({
  createdBy: {},
  identifier: {},
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Account().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Account().load({ id: 'account_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Account().remove({ id: 'account_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Account().update({
  id: 'account_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `AccountEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## BatchEntity

```ts
const batch = client.Batch()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approveUrl` | `string` | No |  |
| `id` | `string` | No |  |
| `payouts` | `any[]` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Batch().create({
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Batch().load({ id: 'batch_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `BatchEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## BeneficiaryEntity

```ts
const beneficiary = client.Beneficiary()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approvalCallbackUrl` | `string` | No |  |
| `authenticationMethods` | `any[]` | No |  |
| `authorisations` | `any[]` | No |  |
| `authorisersCompletedCount` | `number` | No |  |
| `authorisersRequiredCount` | `number` | No |  |
| `beneficiaries` | `any[]` | No |  |
| `beneficiaryEvents` | `any[]` | No |  |
| `canAuthorise` | `boolean` | No |  |
| `canUpdate` | `boolean` | No |  |
| `createdBy` | `Record<string, any>` | Yes |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `Record<string, any>` | No |  |
| `failedBeneficiaries` | `Record<string, any>` | No |  |
| `hasCurrentUserAuthorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isEnabled` | `boolean` | No |  |
| `lastAuthorised` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `sourceAccountIDs` | `any[]` | No |  |
| `sourceAccounts` | `any[]` | No |  |
| `theirReference` | `string` | No |  |

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

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `batchcreate` | `/api/v1/beneficiaries/batchcreate` | `client.Beneficiary().create({ $action: 'batchcreate', ... })` |
| `export` | `/api/v1/beneficiaries/export` | `client.Beneficiary().load({ $action: 'export', ... })` |

An action returns that action's OWN response, which is not necessarily a
Beneficiary record — check the API definition for its shape.

```ts
const result = await client.Beneficiary().create({
  $action: 'batchcreate',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Beneficiary().create({
  createdBy: {},
  currency: 'example_currency',
  name: 'example_name',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Beneficiary().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Beneficiary().load({ id: 'beneficiary_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Beneficiary().remove({ id: 'beneficiary_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Beneficiary().update({
  id: 'beneficiary_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `BeneficiaryEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## BeneficiaryGroupEntity

```ts
const beneficiary_group = client.BeneficiaryGroup()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `groupMembers` | `any[]` | No |  |
| `groupName` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.BeneficiaryGroup().list({ merchant_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `BeneficiaryGroupEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CardEntity

```ts
const card = client.Card()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorizedAmount` | `string` | No |  |
| `currencyCode` | `string` | No |  |
| `isPayerAuthenticationRequired` | `boolean` | No |  |
| `isSoftDecline` | `boolean` | No |  |
| `payerAuthenticationAccessToken` | `string` | No |  |
| `payerAuthenticationMerchantData` | `string` | No |  |
| `payerAuthenticationUrl` | `string` | No |  |
| `payerAuthenticationWindowHeight` | `number` | No |  |
| `payerAuthenticationWindowWidth` | `number` | No |  |
| `paymentRequestCallbackUrl` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |
| `requestID` | `string` | No |  |
| `responseCode` | `string` | No |  |
| `responseType` | `string` | No |  |
| `status` | `string` | No |  |
| `threeDSRedirectUrl` | `string` | No |  |
| `transactionID` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Card().create({
  paymentrequest_id: 'example_paymentrequest_id',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CardEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CardCustomerTokenEntity

```ts
const card_customer_token = client.CardCustomerToken()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cardType` | `string` | No |  |
| `customerEmailAddress` | `string` | No |  |
| `expiryMonth` | `string` | No |  |
| `expiryYear` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lastFourDigits` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `maskedCardNumber` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.CardCustomerToken().list({ customer_email_address: "example", merchant_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.CardCustomerToken().load({ customer_email_address: 'customer_email_address' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.CardCustomerToken().remove()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CardCustomerTokenEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CardPaymentEntity

```ts
const card_payment = client.CardPayment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorizedAmount` | `string` | No |  |
| `currencyCode` | `string` | No |  |
| `isPayerAuthenticationRequired` | `boolean` | No |  |
| `isSoftDecline` | `boolean` | No |  |
| `payerAuthenticationAccessToken` | `string` | No |  |
| `payerAuthenticationMerchantData` | `string` | No |  |
| `payerAuthenticationUrl` | `string` | No |  |
| `payerAuthenticationWindowHeight` | `number` | No |  |
| `payerAuthenticationWindowWidth` | `number` | No |  |
| `paymentRequestCallbackUrl` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |
| `requestID` | `string` | No |  |
| `responseCode` | `string` | No |  |
| `responseType` | `string` | No |  |
| `status` | `string` | No |  |
| `threeDSRedirectUrl` | `string` | No |  |
| `transactionID` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.CardPayment().create({
  paymentrequest_id: 'example_paymentrequest_id',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CardPaymentEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CardPublicKeyEntity

```ts
const card_public_key = client.CardPublicKey()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jwt` | `string` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.CardPublicKey().load({ paymentrequest_id: 'paymentrequest_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CardPublicKeyEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ConsentEntity

```ts
const consent = client.Consent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisationUrl` | `string` | No |  |
| `callbackUrl` | `string` | No |  |
| `consentID` | `string` | No |  |
| `emailAddress` | `string` | No |  |
| `expiryDate` | `string` | No |  |
| `failureCallbackUrl` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `institutionID` | `string` | No |  |
| `isConnectedAccounts` | `boolean` | No |  |
| `isEnabled` | `boolean` | No |  |
| `merchantID` | `string` | No |  |
| `provider` | `string` | No |  |
| `successWebHookUrl` | `string` | No |  |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Consent().create({
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Consent().list({ email: "example", merchant_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Consent().load({ id: 'consent_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Consent().remove({ id: 'consent_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Consent().update({
  id: 'consent_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ConsentEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## CurrencyEntity

```ts
const currency = client.Currency()
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

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Currency().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `CurrencyEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## DirectDebitBatchSubmitEntity

```ts
const direct_debit_batch_submit = client.DirectDebitBatchSubmit()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failedSubmissions` | `Record<string, any>` | No |  |
| `successfulSubmissions` | `any[]` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.DirectDebitBatchSubmit().create({
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `DirectDebitBatchSubmitEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## FxRateEntity

```ts
const fx_rate = client.FxRate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destinationCurrency` | `string` | No |  |
| `exchangeRate` | `number` | No |  |
| `expiryTime` | `string` | No |  |
| `quoteID` | `string` | No |  |
| `sourceCurrency` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.FxRate().list({ destination: "example", source: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.FxRate().load({ destination: 'destination', source: 'source', valid_for_minute: 1 })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `FxRateEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## IPaymentEntity

```ts
const i_payment = client.IPayment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `paymentRequestID` | `string` | No |  |
| `responseType` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.IPayment().create({
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `IPaymentEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MandateEntity

```ts
const mandate = client.Mandate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountNumber` | `string` | No |  |
| `addressLine1` | `string` | Yes |  |
| `addressLine2` | `string` | No |  |
| `approvedAt` | `string` | No |  |
| `city` | `string` | Yes |  |
| `countryCode` | `string` | Yes |  |
| `currency` | `string` | No |  |
| `customerAccountNumber` | `string` | No |  |
| `customerCity` | `string` | No |  |
| `customerCountryCode` | `string` | No |  |
| `customerCountryName` | `string` | No |  |
| `customerEmailAddress` | `string` | No |  |
| `customerFirstName` | `string` | No |  |
| `customerIban` | `string` | No |  |
| `customerLastName` | `string` | No |  |
| `customerSortCode` | `string` | No |  |
| `emailAddress` | `string` | Yes |  |
| `firstName` | `string` | Yes |  |
| `iban` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isRecurring` | `boolean` | No |  |
| `lastName` | `string` | Yes |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `postalCode` | `string` | Yes |  |
| `reference` | `string` | No |  |
| `sortCode` | `string` | No |  |
| `status` | `string` | No |  |
| `supplierBankAccountID` | `string` | No |  |
| `supplierCustomerID` | `string` | No |  |
| `supplierMandateID` | `string` | No |  |
| `supplierName` | `string` | No |  |
| `supplierStatus` | `string` | No |  |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Mandate().create({
  addressLine1: 'example_addressLine1',
  city: 'example_city',
  countryCode: 'example_countryCode',
  emailAddress: 'example_emailAddress',
  firstName: 'example_firstName',
  lastName: 'example_lastName',
  postalCode: 'example_postalCode',
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Mandate().load({ id: 'mandate_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MandateEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MerchantEntity

```ts
const merchant = client.Merchant()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountCurrencies` | `any[]` | No |  |
| `canHaveTrustAccounts` | `boolean` | No |  |
| `cardPaymentProcessor` | `string` | No |  |
| `companyID` | `string` | No |  |
| `displayQrOnHostedPay` | `boolean` | No |  |
| `hostedPayVersion` | `number` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isBlocked` | `boolean` | No |  |
| `isExited` | `boolean` | No |  |
| `isSuspended` | `boolean` | No |  |
| `jurisdiction` | `string` | No |  |
| `logoUrlPng` | `string` | No |  |
| `logoUrlSvg` | `string` | No |  |
| `merchantCategoryCode` | `string` | No |  |
| `name` | `string` | No |  |
| `notes` | `string` | No |  |
| `parentMerchant` | `Record<string, any>` | No |  |
| `paymentAccountLimit` | `number` | No |  |
| `paymentAccounts` | `any[]` | No |  |
| `reason` | `string` | No |  |
| `shortName` | `string` | No |  |
| `supportedPaymentMethodsList` | `any[]` | No |  |
| `suspensionReason` | `string` | No |  |
| `tags` | `any[]` | No |  |
| `timeZoneId` | `string` | No |  |
| `tradingName` | `string` | No |  |
| `webHookLimit` | `number` | No |  |
| `yourRoleName` | `string` | No |  |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `childmerchant` | `/api/v1/merchants/{merchantID}/childmerchants` | `client.Merchant().list({ $action: 'childmerchant', ... })` |
| `paged` | `/api/v1/merchants/paged` | `client.Merchant().list({ $action: 'paged', ... })` |
| `beneficiary_export` | `/api/v1/merchants/{merchantID}/beneficiaries/export` | `client.Merchant().load({ $action: 'beneficiary_export', ... })` |
| `payout_export` | `/api/v1/merchants/{merchantID}/payouts/export` | `client.Merchant().load({ $action: 'payout_export', ... })` |
| `suspend` | `/api/v1/merchants/{merchantId}/suspend` | `client.Merchant().update({ $action: 'suspend', ... })` |

An action returns that action's OWN response, which is not necessarily a
Merchant record — check the API definition for its shape.

```ts
const result = await client.Merchant().list({
  $action: 'childmerchant',
  /* ...the action's own arguments */
})
```

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Merchant().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Merchant().load({ id: 'merchant_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Merchant().remove({ id: 'merchant_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Merchant().update({
  id: 'merchant_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MerchantEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MerchantAuthorisationSettingEntity

```ts
const merchant_authorisation_setting = client.MerchantAuthorisationSetting()
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
| `roleSettings` | `any[]` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.MerchantAuthorisationSetting().list({ merchant_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MerchantAuthorisationSettingEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MerchantDirectDebitMandatePageEntity

```ts
const merchant_direct_debit_mandate_page = client.MerchantDirectDebitMandatePage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approvedAt` | `string` | No |  |
| `currency` | `string` | No |  |
| `customerAccountNumber` | `string` | No |  |
| `customerCity` | `string` | No |  |
| `customerCountryCode` | `string` | No |  |
| `customerCountryName` | `string` | No |  |
| `customerEmailAddress` | `string` | No |  |
| `customerFirstName` | `string` | No |  |
| `customerIban` | `string` | No |  |
| `customerLastName` | `string` | No |  |
| `customerSortCode` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isRecurring` | `boolean` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `reference` | `string` | No |  |
| `status` | `string` | No |  |
| `supplierBankAccountID` | `string` | No |  |
| `supplierCustomerID` | `string` | No |  |
| `supplierMandateID` | `string` | No |  |
| `supplierName` | `string` | No |  |
| `supplierStatus` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.MerchantDirectDebitMandatePage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MerchantDirectDebitMandatePageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MerchantPayByBankSettingEntity

```ts
const merchant_pay_by_bank_setting = client.MerchantPayByBankSetting()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bankCountryCodes` | `any[]` | No |  |
| `bankID` | `string` | No |  |
| `bankName` | `string` | No |  |
| `businessInstitutionID` | `string` | No |  |
| `currency` | `string` | No |  |
| `logo` | `string` | No |  |
| `message` | `string` | No |  |
| `messageImageUrl` | `string` | No |  |
| `order` | `number` | No |  |
| `personalInstitutionID` | `string` | No |  |
| `processor` | `string` | No |  |
| `warningHeading` | `string` | No |  |
| `warningMessage` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.MerchantPayByBankSetting().list({ merchant_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MerchantPayByBankSettingEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MerchantPaymentRequestTemplateEntity

```ts
const merchant_payment_request_template = client.MerchantPaymentRequestTemplate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bankPaymentOptions` | `Record<string, any>` | No |  |
| `cardPaymentAddressOptions` | `Record<string, any>` | No |  |
| `cardPaymentCaptureOptions` | `Record<string, any>` | No |  |
| `customFields` | `any[]` | No |  |
| `defaultFields` | `any[]` | No |  |
| `description` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `name` | `string` | Yes |  |
| `notificationOptions` | `Record<string, any>` | No |  |
| `paymentMethods` | `Record<string, any>` | No |  |
| `paymentTerms` | `Record<string, any>` | No |  |
| `priorityBankOptions` | `Record<string, any>` | No |  |
| `template` | `Record<string, any>` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.MerchantPaymentRequestTemplate().list({ merchant_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.MerchantPaymentRequestTemplate().load({ id: 'merchant_payment_request_template_id', paymentrequest_id: 'paymentrequest_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.MerchantPaymentRequestTemplate().remove({ id: 'merchant_payment_request_template_id', paymentrequest_id: 'paymentrequest_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.MerchantPaymentRequestTemplate().update({
  id: 'merchant_payment_request_template_id',
  paymentrequest_id: 'paymentrequest_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MerchantPaymentRequestTemplateEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MerchantTokenEntity

```ts
const merchant_token = client.MerchantToken()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authenticationMethods` | `any[]` | No |  |
| `authorisations` | `any[]` | No |  |
| `authorisersCompletedCount` | `number` | No |  |
| `authorisersRequiredCount` | `number` | No |  |
| `canAuthorise` | `boolean` | No |  |
| `description` | `string` | No |  |
| `expiresAt` | `string` | No |  |
| `hasCurrentUserAuthorised` | `boolean` | No |  |
| `hmacAlgorithm` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `ipAddressWhitelist` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `isEnabled` | `boolean` | No |  |
| `lastAuthorised` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `nonce` | `string` | Yes |  |
| `permissionTypes` | `any[]` | No |  |
| `requestSignatureVersion` | `number` | No |  |
| `sharedSecretAlgorithm` | `string` | No |  |
| `sharedSecretBase64` | `string` | No |  |
| `token` | `string` | No |  |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.MerchantToken().create({
  nonce: 'example_nonce',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.MerchantToken().list({ merchant_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.MerchantToken().load({ id: 'merchant_token_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.MerchantToken().update({
  id: 'merchant_token_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MerchantTokenEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MetadataEntity

```ts
const metadata = client.Metadata()
```

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `problem` | `/api/v1/metadata/problem` | `client.Metadata().load({ $action: 'problem', ... })` |
| `problemnotification` | `/api/v1/metadata/problemnotification` | `client.Metadata().load({ $action: 'problemnotification', ... })` |

An action returns that action's OWN response, which is not necessarily a
Metadata record — check the API definition for its shape.

```ts
const result = await client.Metadata().load({
  $action: 'problem',
  /* ...the action's own arguments */
})
```

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Metadata().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MetadataEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## NoFrixionVersionEntity

```ts
const no_frixion_version = client.NoFrixionVersion()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `buildVersion` | `number` | No |  |
| `majorVersion` | `number` | No |  |
| `minorVersion` | `number` | No |  |
| `releaseName` | `string` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.NoFrixionVersion().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `NoFrixionVersionEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## OpenBankingEntity

```ts
const open_banking = client.OpenBanking()
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.OpenBanking().create({
  account_id: 'example_account_id',
})
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.OpenBanking().remove()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `OpenBankingEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PayeeverificationEntity

```ts
const payeeverification = client.Payeeverification()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `string` | Yes |  |
| `accountNumber` | `string` | No |  |
| `iban` | `string` | Yes |  |
| `payeeVerifiedAccountName` | `string` | No |  |
| `result` | `string` | No |  |
| `secondaryIdentification` | `string` | No |  |
| `sortCode` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Payeeverification().create({
  accountName: 'example_accountName',
  iban: 'example_iban',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PayeeverificationEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PaymentEntity

```ts
const payment = client.Payment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresses` | `any[]` | No |  |
| `amount` | `number` | No |  |
| `amountPending` | `number` | No |  |
| `amountReceived` | `number` | No |  |
| `amountRefunded` | `number` | No |  |
| `autoSendReceipt` | `boolean` | No |  |
| `baseOriginUrl` | `string` | No |  |
| `callbackUrl` | `string` | No |  |
| `cardAuthorizeOnly` | `boolean` | No |  |
| `cardCreateToken` | `boolean` | No |  |
| `cardCreateTokenMode` | `string` | No |  |
| `cardIgnoreCVN` | `boolean` | No |  |
| `cardNoPayerAuthentication` | `boolean` | No |  |
| `cardProcessorMerchantID` | `string` | No |  |
| `cardStripePaymentIntentID` | `string` | No |  |
| `cardStripePaymentIntentSecret` | `string` | No |  |
| `cardTransmitRawDetails` | `boolean` | No |  |
| `createdByUser` | `Record<string, any>` | Yes |  |
| `currency` | `string` | No |  |
| `customFields` | `any[]` | No |  |
| `customerEmailAddress` | `string` | No |  |
| `customerID` | `string` | No |  |
| `customerName` | `string` | No |  |
| `description` | `string` | No |  |
| `destinationAccount` | `Record<string, any>` | No |  |
| `directDebitPayment` | `Record<string, any>` | No |  |
| `dueDate` | `string` | No |  |
| `events` | `any[]` | No |  |
| `failureCallbackUrl` | `string` | No |  |
| `fieldDisplaySettings` | `any[]` | No |  |
| `formattedAmount` | `string` | No |  |
| `hostedPayCheckoutUrl` | `string` | No |  |
| `id` | `string` | No |  |
| `ignoreAddressVerification` | `boolean` | No |  |
| `inserted` | `string` | No |  |
| `insertedSortable` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `jwk` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `lightningInvoice` | `string` | No |  |
| `lightningInvoiceExpiresAt` | `string` | No |  |
| `merchantDirectDebitMandateID` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantTokenDescription` | `string` | No |  |
| `notificationEmailAddresses` | `string` | No |  |
| `notificationRoleIDs` | `any[]` | No |  |
| `orderID` | `string` | No |  |
| `partialPaymentMethod` | `string` | No |  |
| `partialPaymentSteps` | `string` | No |  |
| `paymentAttempts` | `any[]` | No |  |
| `paymentMethods` | `any[]` | No |  |
| `paymentProcessor` | `string` | No |  |
| `payrunID` | `string` | No |  |
| `pispAccountID` | `string` | No |  |
| `priorityBankID` | `string` | No |  |
| `result` | `Record<string, any>` | No |  |
| `sandboxSettleDelayInSeconds` | `number` | No |  |
| `shippingAddress` | `Record<string, any>` | No |  |
| `shippingAddressCity` | `string` | No |  |
| `shippingAddressCountryCode` | `string` | No |  |
| `shippingAddressCounty` | `string` | No |  |
| `shippingAddressLine1` | `string` | No |  |
| `shippingAddressLine2` | `string` | No |  |
| `shippingAddressPostCode` | `string` | No |  |
| `shippingEmail` | `string` | No |  |
| `shippingFirstName` | `string` | No |  |
| `shippingLastName` | `string` | No |  |
| `shippingPhone` | `string` | No |  |
| `status` | `string` | No |  |
| `successWebHookUrl` | `string` | No |  |
| `tagIds` | `any[]` | No |  |
| `tags` | `any[]` | No |  |
| `title` | `string` | No |  |
| `tokenisedCards` | `any[]` | No |  |
| `transactions` | `any[]` | No |  |
| `useHostedPaymentPage` | `boolean` | No |  |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Payment().create({
  createdByUser: {},
})
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Payment().load({ id: 'payment_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Payment().update({
  id: 'payment_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PaymentEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PaymentAccountEntity

```ts
const payment_account = client.PaymentAccount()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `string` | No |  |
| `accountSupplierName` | `string` | No |  |
| `availableBalance` | `number` | No |  |
| `availableBalanceMinorUnits` | `number` | No |  |
| `balance` | `number` | No |  |
| `balanceMinorUnits` | `number` | No |  |
| `bankName` | `string` | No |  |
| `consentID` | `string` | No |  |
| `createdBy` | `Record<string, any>` | Yes |  |
| `createdByDisplayName` | `string` | No |  |
| `currency` | `string` | No |  |
| `defaultPaymentRail` | `string` | No |  |
| `displayName` | `string` | No |  |
| `expiryDate` | `string` | No |  |
| `externalAccountIcon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `Record<string, any>` | Yes |  |
| `inserted` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `isConnectedAccount` | `boolean` | No |  |
| `isDefault` | `boolean` | No |  |
| `isTrustAccount` | `boolean` | No |  |
| `isVirtual` | `boolean` | No |  |
| `lastTransaction` | `Record<string, any>` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `physicalAccountID` | `string` | No |  |
| `rules` | `any[]` | No |  |
| `submittedPayoutsBalance` | `number` | No |  |
| `submittedPayoutsBalanceMinorUnits` | `number` | No |  |
| `summary` | `string` | No |  |
| `supplierSepaInstantStatus` | `string` | No |  |
| `xeroBankFeedConnectionStatus` | `string` | No |  |
| `xeroBankFeedLastSyncedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `string` | No |  |
| `xeroBankFeedSyncStatus` | `string` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `number` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PaymentAccount().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PaymentAccountEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PaymentAccountMinimalEntity

```ts
const payment_account_minimal = client.PaymentAccountMinimal()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `string` | No |  |
| `availableBalance` | `number` | No |  |
| `balance` | `number` | No |  |
| `balanceMinorUnits` | `number` | No |  |
| `currency` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `Record<string, any>` | Yes |  |
| `isArchived` | `boolean` | No |  |
| `isConnectedAccount` | `boolean` | No |  |
| `merchantID` | `string` | No |  |
| `submittedPayoutsBalance` | `number` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PaymentAccountMinimal().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PaymentAccountMinimalEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PaymentInitiationEntity

```ts
const payment_initiation = client.PaymentInitiation()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `paymentInitiationID` | `string` | No |  |
| `paymentRequestCallbackUrl` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |
| `redirectUrl` | `string` | No |  |
| `responseType` | `string` | No |  |
| `specificErrorMessage` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PaymentInitiation().create({
  paymentrequest_id: 'example_paymentrequest_id',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PaymentInitiationEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PaymentRequestEntity

```ts
const payment_request = client.PaymentRequest()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `addresses` | `any[]` | No |  |
| `amount` | `number` | No |  |
| `amountPending` | `number` | No |  |
| `amountReceived` | `number` | No |  |
| `amountRefunded` | `number` | No |  |
| `autoSendReceipt` | `boolean` | No |  |
| `baseOriginUrl` | `string` | No |  |
| `callbackUrl` | `string` | No |  |
| `cardAuthorizeOnly` | `boolean` | No |  |
| `cardCreateToken` | `boolean` | No |  |
| `cardCreateTokenMode` | `string` | No |  |
| `cardIgnoreCVN` | `boolean` | No |  |
| `cardProcessorMerchantID` | `string` | No |  |
| `cardStripePaymentIntentID` | `string` | No |  |
| `cardStripePaymentIntentSecret` | `string` | No |  |
| `createdByUser` | `Record<string, any>` | Yes |  |
| `currency` | `string` | No |  |
| `customFields` | `any[]` | No |  |
| `customerEmailAddress` | `string` | No |  |
| `customerID` | `string` | No |  |
| `customerName` | `string` | No |  |
| `description` | `string` | No |  |
| `destinationAccount` | `Record<string, any>` | No |  |
| `directDebitPayment` | `Record<string, any>` | No |  |
| `doSimulateSettlementFailure` | `boolean` | No |  |
| `dueDate` | `string` | No |  |
| `errorDescription` | `string` | No |  |
| `events` | `any[]` | No |  |
| `failedPaymentRequests` | `Record<string, any>` | No |  |
| `failureCallbackUrl` | `string` | No |  |
| `fieldDisplaySettings` | `any[]` | No |  |
| `formattedAmount` | `string` | No |  |
| `hostedPayCheckoutUrl` | `string` | No |  |
| `id` | `string` | No |  |
| `ignoreAddressVerification` | `boolean` | No |  |
| `inserted` | `string` | No |  |
| `insertedSortable` | `string` | No |  |
| `institution` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `jwk` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `lightningInvoice` | `string` | No |  |
| `lightningInvoiceExpiresAt` | `string` | No |  |
| `merchantDirectDebitMandateID` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantTokenDescription` | `string` | No |  |
| `notificationEmailAddresses` | `string` | No |  |
| `notificationRoleIDs` | `any[]` | No |  |
| `orderID` | `string` | No |  |
| `partialPaymentMethod` | `string` | No |  |
| `partialPaymentSteps` | `string` | No |  |
| `paymentAttempts` | `any[]` | No |  |
| `paymentInitiationID` | `string` | No |  |
| `paymentMethods` | `any[]` | No |  |
| `paymentProcessor` | `string` | No |  |
| `paymentRequests` | `any[]` | No |  |
| `payrunID` | `string` | No |  |
| `pispAccountID` | `string` | No |  |
| `priorityBankID` | `string` | No |  |
| `result` | `Record<string, any>` | No |  |
| `sandboxSettleDelayInSeconds` | `number` | No |  |
| `shippingAddress` | `Record<string, any>` | No |  |
| `status` | `string` | No |  |
| `successWebHookUrl` | `string` | No |  |
| `tags` | `any[]` | No |  |
| `title` | `string` | No |  |
| `tokenisedCards` | `any[]` | No |  |
| `transactions` | `any[]` | No |  |
| `useHostedPaymentPage` | `boolean` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PaymentRequest().create({
  createdByUser: {},
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PaymentRequest().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.PaymentRequest().load()
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.PaymentRequest().remove({ id: 'id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.PaymentRequest().update({
  paymentrequest_id: 'paymentrequest_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PaymentRequestEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PaymentRequestEventEntity

```ts
const payment_request_event = client.PaymentRequestEvent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `number` | Yes |  |
| `applePayTransactionID` | `string` | No |  |
| `cardAuthorizationResponseID` | `string` | No |  |
| `cardExpiryMonth` | `number` | No |  |
| `cardExpiryYear` | `number` | No |  |
| `cardIssuer` | `string` | No |  |
| `cardIssuerCountry` | `string` | No |  |
| `cardLastFourDigits` | `string` | No |  |
| `cardRequestID` | `string` | No |  |
| `cardScheme` | `string` | No |  |
| `cardTokenCustomerID` | `string` | No |  |
| `cardTransactionID` | `string` | No |  |
| `currency` | `string` | No |  |
| `directDebitPaymentID` | `string` | No |  |
| `directDebitPaymentReference` | `string` | No |  |
| `drirectDebitMandateID` | `string` | No |  |
| `errorMessage` | `string` | No |  |
| `errorReason` | `string` | No |  |
| `eventType` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lightningInvoice` | `string` | No |  |
| `lightningRHash` | `string` | No |  |
| `originUrl` | `string` | No |  |
| `paymentMethodType` | `string` | No |  |
| `paymentProcessorName` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |
| `pispBankStatus` | `string` | No |  |
| `pispPaymentInitiationID` | `string` | No |  |
| `pispPaymentInstitutionName` | `string` | No |  |
| `pispPaymentServiceProviderID` | `string` | No |  |
| `pispRedirectUrl` | `string` | No |  |
| `reconciledTransactionID` | `string` | No |  |
| `refundPayoutID` | `string` | No |  |
| `status` | `string` | No |  |
| `walletName` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PaymentRequestEvent().list({ paymentrequest_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PaymentRequestEventEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PaymentRequestMetricEntity

```ts
const payment_request_metric = client.PaymentRequestMetric()
```

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.PaymentRequestMetric().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PaymentRequestMetricEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PaymentRequestMinimalEntity

```ts
const payment_request_minimal = client.PaymentRequestMinimal()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `number` | No |  |
| `amountPending` | `number` | No |  |
| `amountReceived` | `number` | No |  |
| `amountRefunded` | `number` | No |  |
| `callbackUrl` | `string` | No |  |
| `cardStripePaymentIntentSecret` | `string` | No |  |
| `countryCode` | `string` | No |  |
| `currency` | `string` | No |  |
| `customFieldsToDisplay` | `any[]` | No |  |
| `description` | `string` | No |  |
| `dueDate` | `string` | No |  |
| `fieldDisplaySettings` | `any[]` | No |  |
| `googlePayMerchantID` | `string` | No |  |
| `id` | `string` | No |  |
| `jwk` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantLogoUrlPng` | `string` | No |  |
| `merchantLogoUrlSvg` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `merchantShortName` | `string` | No |  |
| `partialPaymentMethod` | `string` | No |  |
| `paymentAttempts` | `any[]` | No |  |
| `paymentMethodsList` | `any[]` | No |  |
| `paymentProcessor` | `string` | No |  |
| `paymentProcessorKey` | `string` | No |  |
| `pispError` | `string` | No |  |
| `priorityBankID` | `string` | No |  |
| `status` | `string` | No |  |
| `stripeAccountID` | `string` | No |  |
| `title` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PaymentRequestMinimal().list({ paymentrequest_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PaymentRequestMinimalEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PaymentRequestResultEntity

```ts
const payment_request_result = client.PaymentRequestResult()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `number` | No |  |
| `amountPending` | `number` | No |  |
| `amountReceived` | `number` | No |  |
| `amountRefunded` | `number` | No |  |
| `currency` | `string` | No |  |
| `customerID` | `string` | No |  |
| `paymentRequestID` | `string` | No |  |
| `payments` | `any[]` | No |  |
| `pispAuthorizations` | `any[]` | No |  |
| `requestedAmount` | `number` | No |  |
| `result` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PaymentRequestResult().list({ paymentrequest_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PaymentRequestResultEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PayoutEntity

```ts
const payout = client.Payout()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountID` | `string` | No |  |
| `allowIncomplete` | `boolean` | No |  |
| `amount` | `number` | No |  |
| `amountMinorUnits` | `number` | No |  |
| `approvePayoutUrl` | `string` | No |  |
| `approverID` | `string` | No |  |
| `authenticationMethods` | `any[]` | No |  |
| `authorisations` | `any[]` | No |  |
| `authorisersCompletedCount` | `number` | No |  |
| `authorisersRequiredCount` | `number` | No |  |
| `batchPayoutID` | `string` | No |  |
| `beneficiary` | `Record<string, any>` | Yes |  |
| `beneficiaryID` | `string` | No |  |
| `canAuthorise` | `boolean` | No |  |
| `canProcess` | `boolean` | No |  |
| `canUpdate` | `boolean` | No |  |
| `chargeBearer` | `string` | No |  |
| `createdBy` | `string` | No |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | No |  |
| `currentUserID` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `Record<string, any>` | No |  |
| `documents` | `any[]` | No |  |
| `events` | `any[]` | No |  |
| `failedPayouts` | `Record<string, any>` | No |  |
| `formattedAmount` | `string` | No |  |
| `formattedFxDestinationAmount` | `string` | No |  |
| `formattedSchedule` | `string` | No |  |
| `formattedScheduleDayOnly` | `string` | No |  |
| `formattedSourceAccountAvailableBalance` | `string` | No |  |
| `fxDestinationAmount` | `number` | No |  |
| `fxDestinationAmountMinorUnits` | `number` | No |  |
| `fxDestinationCurrency` | `string` | No |  |
| `fxQuoteExpiresAt` | `string` | No |  |
| `fxQuoteID` | `string` | No |  |
| `fxRate` | `number` | No |  |
| `fxUseDestinationAmount` | `boolean` | No |  |
| `hasCurrentUserAuthorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoiceID` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `isFailed` | `boolean` | No |  |
| `isSettled` | `boolean` | No |  |
| `isSubmitted` | `boolean` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantTokenDescription` | `string` | No |  |
| `nonce` | `string` | No |  |
| `paymentProcessor` | `string` | No |  |
| `paymentRail` | `string` | No |  |
| `payouts` | `any[]` | No |  |
| `payrunID` | `string` | No |  |
| `payrunName` | `string` | No |  |
| `reason` | `string` | No |  |
| `rule` | `Record<string, any>` | No |  |
| `scheduleDate` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `sourceAccountAvailableBalance` | `number` | No |  |
| `sourceAccountAvailableBalanceMinorUnits` | `number` | No |  |
| `sourceAccountBic` | `string` | No |  |
| `sourceAccountCurrency` | `string` | No |  |
| `sourceAccountIban` | `string` | No |  |
| `sourceAccountIdentifier` | `Record<string, any>` | Yes |  |
| `sourceAccountName` | `string` | No |  |
| `sourceAccountNumber` | `string` | No |  |
| `sourceAccountSortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tagIds` | `any[]` | No |  |
| `tags` | `any[]` | No |  |
| `theirReference` | `string` | No |  |
| `topupPayrunID` | `string` | No |  |
| `transactedAmount` | `number` | No |  |
| `transactedFxAmount` | `number` | No |  |
| `transactedFxRate` | `number` | No |  |
| `type` | `string` | No |  |
| `userID` | `string` | No |  |
| `yourReference` | `string` | No |  |

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

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `batchcreate` | `/api/v1/payouts/batchcreate` | `client.Payout().create({ $action: 'batchcreate', ... })` |
| `send` | `/api/v1/payouts/send` | `client.Payout().create({ $action: 'send', ... })` |
| `sendbeneficiary` | `/api/v1/payouts/sendbeneficiary` | `client.Payout().create({ $action: 'sendbeneficiary', ... })` |
| `export` | `/api/v1/payouts/export` | `client.Payout().load({ $action: 'export', ... })` |
| `proof` | `/api/v1/payouts/{id}/proof` | `client.Payout().load({ $action: 'proof', ... })` |
| `batchdelete` | `/api/v1/payouts/batchdelete` | `client.Payout().remove({ $action: 'batchdelete', ... })` |

An action returns that action's OWN response, which is not necessarily a
Payout record — check the API definition for its shape.

```ts
const result = await client.Payout().create({
  $action: 'batchcreate',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Payout().create({
  beneficiary: {},
  sourceAccountIdentifier: {},
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Payout().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Payout().load({ id: 'payout_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Payout().remove({ id: 'payout_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Payout().update({
  id: 'payout_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PayoutEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PayoutKeysetPageEntity

```ts
const payout_keyset_page = client.PayoutKeysetPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountID` | `string` | No |  |
| `amount` | `number` | No |  |
| `amountMinorUnits` | `number` | No |  |
| `approvePayoutUrl` | `string` | No |  |
| `approverID` | `string` | No |  |
| `authenticationMethods` | `any[]` | No |  |
| `authorisations` | `any[]` | No |  |
| `authorisersCompletedCount` | `number` | No |  |
| `authorisersRequiredCount` | `number` | No |  |
| `batchPayoutID` | `string` | No |  |
| `beneficiary` | `Record<string, any>` | Yes |  |
| `canAuthorise` | `boolean` | No |  |
| `canProcess` | `boolean` | No |  |
| `canUpdate` | `boolean` | No |  |
| `chargeBearer` | `string` | No |  |
| `createdBy` | `string` | No |  |
| `createdByEmailAddress` | `string` | No |  |
| `currency` | `string` | No |  |
| `currentUserID` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `Record<string, any>` | No |  |
| `documents` | `any[]` | No |  |
| `events` | `any[]` | No |  |
| `formattedAmount` | `string` | No |  |
| `formattedFxDestinationAmount` | `string` | No |  |
| `formattedSchedule` | `string` | No |  |
| `formattedScheduleDayOnly` | `string` | No |  |
| `formattedSourceAccountAvailableBalance` | `string` | No |  |
| `fxDestinationAmount` | `number` | No |  |
| `fxDestinationAmountMinorUnits` | `number` | No |  |
| `fxDestinationCurrency` | `string` | No |  |
| `fxQuoteExpiresAt` | `string` | No |  |
| `fxQuoteID` | `string` | No |  |
| `fxRate` | `number` | No |  |
| `fxUseDestinationAmount` | `boolean` | No |  |
| `hasCurrentUserAuthorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoiceID` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `isFailed` | `boolean` | No |  |
| `isSettled` | `boolean` | No |  |
| `isSubmitted` | `boolean` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantTokenDescription` | `string` | No |  |
| `nonce` | `string` | No |  |
| `paymentProcessor` | `string` | No |  |
| `paymentRail` | `string` | No |  |
| `payrunID` | `string` | No |  |
| `payrunName` | `string` | No |  |
| `rule` | `Record<string, any>` | No |  |
| `scheduleDate` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `sourceAccountAvailableBalance` | `number` | No |  |
| `sourceAccountAvailableBalanceMinorUnits` | `number` | No |  |
| `sourceAccountBic` | `string` | No |  |
| `sourceAccountCurrency` | `string` | No |  |
| `sourceAccountIban` | `string` | No |  |
| `sourceAccountIdentifier` | `Record<string, any>` | Yes |  |
| `sourceAccountName` | `string` | No |  |
| `sourceAccountNumber` | `string` | No |  |
| `sourceAccountSortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tags` | `any[]` | No |  |
| `theirReference` | `string` | No |  |
| `topupPayrunID` | `string` | No |  |
| `transactedAmount` | `number` | No |  |
| `transactedFxAmount` | `number` | No |  |
| `transactedFxRate` | `number` | No |  |
| `type` | `string` | No |  |
| `userID` | `string` | No |  |
| `yourReference` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PayoutKeysetPage().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PayoutKeysetPageEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PayoutMetricEntity

```ts
const payout_metric = client.PayoutMetric()
```

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.PayoutMetric().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PayoutMetricEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PayrunEntity

```ts
const payrun = client.Payrun()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisationDate` | `string` | No |  |
| `authorisations` | `any[]` | No |  |
| `authorisersCompletedCount` | `number` | No |  |
| `authorisersRequiredCount` | `number` | No |  |
| `batchPayoutID` | `string` | No |  |
| `canAuthorise` | `boolean` | No |  |
| `canDelete` | `boolean` | No |  |
| `canEdit` | `boolean` | No |  |
| `events` | `any[]` | No |  |
| `hasCurrentUserAuthorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoices` | `any[]` | No |  |
| `invoicesMinimal` | `any[]` | No |  |
| `isArchived` | `boolean` | No |  |
| `lastUpdated` | `string` | No |  |
| `lastUpdatedBy` | `Record<string, any>` | Yes |  |
| `merchantID` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | No |  |
| `notes` | `string` | No |  |
| `payments` | `any[]` | No |  |
| `payouts` | `any[]` | No |  |
| `payoutsCount` | `number` | No |  |
| `reason` | `string` | No |  |
| `scheduleDate` | `string` | No |  |
| `scheduledDate` | `string` | No |  |
| `sourceAccounts` | `any[]` | No |  |
| `status` | `string` | No |  |
| `totalEur` | `number` | No |  |
| `totalGbp` | `number` | No |  |
| `totalUsd` | `number` | No |  |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `request_authorisation` | `/api/v1/payruns/{id}/request-authorisation` | `client.Payrun().create({ $action: 'request_authorisation', ... })` |
| `submit` | `/api/v1/payruns/{id}/submit` | `client.Payrun().create({ $action: 'submit', ... })` |
| `archive` | `/api/v1/payruns/{id}/archive` | `client.Payrun().remove({ $action: 'archive', ... })` |
| `cancel` | `/api/v1/payruns/{id}/cancel` | `client.Payrun().update({ $action: 'cancel', ... })` |
| `reject` | `/api/v1/payruns/{id}/reject` | `client.Payrun().update({ $action: 'reject', ... })` |
| `unarchive` | `/api/v1/payruns/{id}/unarchive` | `client.Payrun().update({ $action: 'unarchive', ... })` |

An action returns that action's OWN response, which is not necessarily a
Payrun record — check the API definition for its shape.

```ts
const result = await client.Payrun().create({
  $action: 'request_authorisation',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Payrun().create({
  id: 'example_id',
  lastUpdatedBy: {},
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Payrun().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Payrun().load({ id: 'payrun_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Payrun().remove({ id: 'payrun_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Payrun().update({
  id: 'payrun_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PayrunEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ReportEntity

```ts
const report = client.Report()
```

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `initiate` | `/api/v1/reports/{id}/initiate` | `client.Report().update({ $action: 'initiate', ... })` |

An action returns that action's OWN response, which is not necessarily a
Report record — check the API definition for its shape.

```ts
const result = await client.Report().update({
  $action: 'initiate',
  /* ...the action's own arguments */
})
```

### Operations

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Report().update({
  id: 'id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ReportEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ReportResultEntity

```ts
const report_result = client.ReportResult()
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

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.ReportResult().load({ id: 1, report_id: 'report_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ReportResultEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## RoleEntity

```ts
const role = client.Role()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failedRoles` | `Record<string, any>` | No |  |
| `roles` | `any[]` | No |  |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `batchcreate` | `/api/v1/merchants/{merchantID}/roles/batchcreate` | `client.Role().create({ $action: 'batchcreate', ... })` |

An action returns that action's OWN response, which is not necessarily a
Role record — check the API definition for its shape.

```ts
const result = await client.Role().create({
  $action: 'batchcreate',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Role().create({
  merchant_id: 'example_merchant_id',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `RoleEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## RuleEntity

```ts
const rule = client.Rule()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account` | `Record<string, any>` | No |  |
| `accountID` | `string` | No |  |
| `approveUrl` | `string` | No |  |
| `approverID` | `string` | No |  |
| `authenticationMethods` | `any[]` | No |  |
| `authorisations` | `any[]` | No |  |
| `authorisersCompletedCount` | `number` | No |  |
| `authorisersRequiredCount` | `number` | No |  |
| `canAuthorise` | `boolean` | No |  |
| `createdBy` | `Record<string, any>` | Yes |  |
| `description` | `string` | No |  |
| `endAt` | `string` | No |  |
| `hasCurrentUserAuthorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isDisabled` | `boolean` | No |  |
| `lastExecutedAt` | `string` | No |  |
| `lastRunAtTransactionDate` | `string` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | Yes |  |
| `onApprovedWebHookUrl` | `string` | No |  |
| `onExecutionErrorWebHookUrl` | `string` | No |  |
| `onExecutionSuccessWebHookUrl` | `string` | No |  |
| `startAt` | `string` | No |  |
| `status` | `string` | No |  |
| `sweepAction` | `Record<string, any>` | No |  |
| `timeZoneId` | `string` | No |  |
| `triggerCronExpression` | `string` | No |  |
| `triggerOnPayIn` | `boolean` | No |  |
| `userID` | `string` | No |  |
| `webHookSecret` | `string` | No |  |

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

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `disable` | `/api/v1/rules/{id}/disable` | `client.Rule().update({ $action: 'disable', ... })` |

An action returns that action's OWN response, which is not necessarily a
Rule record — check the API definition for its shape.

```ts
const result = await client.Rule().update({
  $action: 'disable',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Rule().create({
  createdBy: {},
  nonce: 'example_nonce',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Rule().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Rule().load({ id: 'rule_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Rule().remove({ id: 'rule_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Rule().update({
  id: 'rule_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `RuleEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## RuleEventEntity

```ts
const rule_event = client.RuleEvent()
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
| `user` | `Record<string, any>` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.RuleEvent().list({ id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `RuleEventEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TagEntity

```ts
const tag = client.Tag()
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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Tag().create({
  merchant_id: 'example_merchant_id',
  merchantID: 'example_merchantID',
  name: 'example_name',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Tag().list({ merchant_id: "example" })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TagEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TokenEntity

```ts
const token = client.Token()
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Token().create({
  id: 'example_id',
})
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Token().remove({ id: 'id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TokenEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TransactionEntity

```ts
const transaction = client.Transaction()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountID` | `string` | No |  |
| `accountName` | `string` | No |  |
| `accountSequenceNumber` | `number` | No |  |
| `addressDetails` | `Record<string, any>` | No |  |
| `amount` | `number` | No |  |
| `amountMinorUnits` | `number` | No |  |
| `balance` | `number` | No |  |
| `balanceMinorUnits` | `number` | No |  |
| `bookingDateTime` | `string` | No |  |
| `chargeDetails` | `Record<string, any>` | No |  |
| `content` | `any[]` | No |  |
| `counterparty` | `Record<string, any>` | No |  |
| `counterpartySummary` | `string` | No |  |
| `currency` | `string` | No |  |
| `currencyExchange` | `Record<string, any>` | No |  |
| `date` | `string` | No |  |
| `description` | `string` | No |  |
| `enrichment` | `Record<string, any>` | No |  |
| `fxAmount` | `number` | No |  |
| `fxCurrency` | `string` | No |  |
| `fxRate` | `number` | No |  |
| `grossAmount` | `Record<string, any>` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `isoBankTransactionCode` | `Record<string, any>` | No |  |
| `merchant` | `Record<string, any>` | No |  |
| `merchantID` | `string` | No |  |
| `pageNumber` | `number` | No |  |
| `pageSize` | `number` | No |  |
| `payeeDetails` | `Record<string, any>` | Yes |  |
| `payerDetails` | `Record<string, any>` | Yes |  |
| `paymentRequestCustomFields` | `Record<string, any>` | No |  |
| `paymentRequestID` | `string` | No |  |
| `payoutID` | `string` | No |  |
| `proprietaryBankTransactionCode` | `Record<string, any>` | No |  |
| `rawReference` | `string` | No |  |
| `reference` | `string` | No |  |
| `ruleID` | `string` | No |  |
| `statementReferences` | `any[]` | No |  |
| `status` | `string` | No |  |
| `supplementaryData` | `any` | No |  |
| `tags` | `any[]` | No |  |
| `theirReference` | `string` | No |  |
| `totalPages` | `number` | No |  |
| `totalSize` | `number` | No |  |
| `transactionAmount` | `Record<string, any>` | Yes |  |
| `transactionDate` | `string` | No |  |
| `transactionInformation` | `any[]` | No |  |
| `transactionMutability` | `string` | No |  |
| `type` | `string` | No |  |
| `valueDateTime` | `string` | No |  |
| `virtualIBAN` | `string` | No |  |
| `yourReference` | `string` | No |  |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `tag` | `/api/v1/transactions/{id}/tags` | `client.Transaction().create({ $action: 'tag', ... })` |
| `export` | `/api/v1/transactions/{accountID}/export` | `client.Transaction().load({ $action: 'export', ... })` |
| `proof` | `/api/v1/transactions/{id}/proof` | `client.Transaction().load({ $action: 'proof', ... })` |
| `tag` | `/api/v1/transactions/{id}/tag` | `client.Transaction().remove({ $action: 'tag', ... })` |

An action returns that action's OWN response, which is not necessarily a
Transaction record — check the API definition for its shape.

```ts
const result = await client.Transaction().create({
  $action: 'tag',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Transaction().create({
  id: 'example_id',
  grossAmount: {},
  payeeDetails: {},
  payerDetails: {},
  transactionAmount: {},
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Transaction().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Transaction().load({ id: 'transaction_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Transaction().remove({ id: 'transaction_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TransactionEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## UserEntity

```ts
const user = client.User()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `clientSessionTimeouts` | `any[]` | No |  |
| `emailAddress` | `string` | Yes |  |
| `firstName` | `string` | Yes |  |
| `id` | `string` | No |  |
| `lastName` | `string` | Yes |  |
| `passkeyAdded` | `boolean` | No |  |
| `permissions` | `Record<string, any>` | No |  |
| `profile` | `string` | No |  |
| `rolesWithScope` | `any[]` | No |  |
| `twoFactorEnabled` | `boolean` | No |  |
| `userInviteID` | `string` | No |  |

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

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `userspaged` | `/api/v1/user/{merchantID}/userspaged` | `client.User().list({ $action: 'userspaged', ... })` |

An action returns that action's OWN response, which is not necessarily a
User record — check the API definition for its shape.

```ts
const result = await client.User().list({
  $action: 'userspaged',
  /* ...the action's own arguments */
})
```

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.User().list()
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.User().update({
  id: 'id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `UserEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## UserInviteEntity

```ts
const user_invite = client.UserInvite()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisationStatus` | `Record<string, any>` | No |  |
| `failedUserInvites` | `Record<string, any>` | No |  |
| `id` | `string` | No |  |
| `initialRoleID` | `string` | No |  |
| `inviteeEmailAddress` | `string` | No |  |
| `inviteeFirstName` | `string` | No |  |
| `inviteeLastName` | `string` | No |  |
| `inviterEmailAddress` | `string` | No |  |
| `inviterFirstName` | `string` | No |  |
| `inviterLastName` | `string` | No |  |
| `isAuthorised` | `boolean` | No |  |
| `isInviteeRegistered` | `boolean` | No |  |
| `lastInvited` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `message` | `string` | No |  |
| `registrationUrl` | `string` | No |  |
| `sendInviteEmail` | `boolean` | No |  |
| `status` | `string` | No |  |
| `user` | `Record<string, any>` | Yes |  |
| `userID` | `string` | No |  |
| `userInvites` | `any[]` | No |  |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.UserInvite().create({
  user: {},
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.UserInvite().list({ merchant_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.UserInvite().load({ id: 'user_invite_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.UserInvite().remove({ id: 'user_invite_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.UserInvite().update({
  id: 'user_invite_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `UserInviteEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## VirtualEntity

```ts
const virtual = client.Virtual()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `accountName` | `string` | No |  |
| `accountSupplierName` | `string` | No |  |
| `availableBalance` | `number` | No |  |
| `availableBalanceMinorUnits` | `number` | No |  |
| `balance` | `number` | No |  |
| `balanceMinorUnits` | `number` | No |  |
| `bankName` | `string` | No |  |
| `consentID` | `string` | No |  |
| `createdBy` | `Record<string, any>` | Yes |  |
| `createdByDisplayName` | `string` | No |  |
| `currency` | `string` | No |  |
| `defaultPaymentRail` | `string` | No |  |
| `displayName` | `string` | No |  |
| `expiryDate` | `string` | No |  |
| `externalAccountIcon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `Record<string, any>` | Yes |  |
| `inserted` | `string` | No |  |
| `isArchived` | `boolean` | No |  |
| `isConnectedAccount` | `boolean` | No |  |
| `isDefault` | `boolean` | No |  |
| `isTrustAccount` | `boolean` | No |  |
| `isVirtual` | `boolean` | No |  |
| `lastTransaction` | `Record<string, any>` | No |  |
| `lastUpdated` | `string` | No |  |
| `merchantID` | `string` | No |  |
| `merchantName` | `string` | No |  |
| `name` | `string` | Yes |  |
| `physicalAccountID` | `string` | No |  |
| `rules` | `any[]` | No |  |
| `submittedPayoutsBalance` | `number` | No |  |
| `submittedPayoutsBalanceMinorUnits` | `number` | No |  |
| `summary` | `string` | No |  |
| `supplierSepaInstantStatus` | `string` | No |  |
| `xeroBankFeedConnectionStatus` | `string` | No |  |
| `xeroBankFeedLastSyncedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailedAt` | `string` | No |  |
| `xeroBankFeedSyncLastFailureReason` | `string` | No |  |
| `xeroBankFeedSyncStatus` | `string` | No |  |
| `xeroUnsynchronisedTransactionsCount` | `number` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Virtual().create({
  account_id: 'example_account_id',
  createdBy: {},
  identifier: {},
  name: 'example_name',
})
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Virtual().update({
  account_id: 'account_id',
  id: 'id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `VirtualEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## WebhookEntity

```ts
const webhook = client.Webhook()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destinationUrl` | `string` | No |  |
| `emailAddress` | `string` | No |  |
| `failedNotificationEmailAddress` | `string` | No |  |
| `id` | `string` | No |  |
| `isActive` | `boolean` | No |  |
| `merchantID` | `string` | No |  |
| `notificationMethod` | `string` | No |  |
| `resourceTypes` | `any[]` | No |  |
| `retry` | `boolean` | No |  |
| `secret` | `string` | No |  |
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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Webhook().create({
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Webhook().list({ merchant_id: "example" })
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Webhook().load({ id: 'webhook_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Webhook().remove({ id: 'webhook_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Webhook().update({
  id: 'webhook_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `WebhookEntity` instance with the same client and
options.

#### `client()`

Return the parent `NofrixionSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new NofrixionSDK({
  feature: {
    test: { active: true },
  }
})
```

