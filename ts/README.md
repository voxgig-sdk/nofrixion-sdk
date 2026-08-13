# Nofrixion TypeScript SDK



The TypeScript SDK for the Nofrixion API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.Account()` — each with a small set of operations (`list`, `load`, `create`, `update`, `remove`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/nofrixion-sdk/releases](https://github.com/voxgig-sdk/nofrixion-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { NofrixionSDK } from '@voxgig-sdk/nofrixion'

const client = new NofrixionSDK({
  apikey: process.env.NOFRIXION_APIKEY,
})
```

### 2. List account records

`list()` resolves to an array of Account ENTITIES — every operation
resolves to entities, not raw records. Iterate them directly, and call
`.data()` on one for the record it holds:

```ts
const accounts = await client.Account().list()

for (const account of accounts) {
  console.log(account)
}
```

### 3. Load a cardcustomertoken

CardCustomerToken is nested under customer_email_address, so provide the `customer_email_address`.
`load()` returns the entity directly and throws on failure:

```ts
try {
  const cardcustomertoken = await client.CardCustomerToken().load({
    customer_email_address: 'example_customer_email_address',
  })
  console.log(cardcustomertoken)
} catch (err) {
  console.error('load failed:', err)
}
```

### 4. Create, update, and remove

```ts
// Create — returns the created Account ENTITY (.data() for the record)
const created = await client.Account().create({
  createdBy: {},
  identifier: {},
})

// Update — the id comes off the returned entity's data()
const updated = await client.Account().update({
  id: created.data().id!,
  account_id: 'example_account_id',
  amount: 1,
})

// Remove
await client.Account().remove({
  id: created.data().id!,
})
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const consents = await client.Consent().list()
  console.log(consents)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = NofrixionSDK.test()

const consent = await client.Consent().list()
// consent is the entity, populated with mock response data
// — call consent.data() for the record itself
console.log(consent)
```

You can also use the instance method:

```ts
const client = new NofrixionSDK({ apikey: '...' })
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.Consent()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data.id)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new NofrixionSDK({
  apikey: '...',
  extend: [logger],
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
cd ts && npm test
```


## Reference

### NofrixionSDK

#### Constructor

```ts
new NofrixionSDK(options?: {
  apikey?: string
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `Account(data?)` | `AccountEntity` | Create an Account entity instance. |
| `Batch(data?)` | `BatchEntity` | Create a Batch entity instance. |
| `Beneficiary(data?)` | `BeneficiaryEntity` | Create a Beneficiary entity instance. |
| `BeneficiaryGroup(data?)` | `BeneficiaryGroupEntity` | Create a BeneficiaryGroup entity instance. |
| `Card(data?)` | `CardEntity` | Create a Card entity instance. |
| `CardCustomerToken(data?)` | `CardCustomerTokenEntity` | Create a CardCustomerToken entity instance. |
| `CardPayment(data?)` | `CardPaymentEntity` | Create a CardPayment entity instance. |
| `CardPublicKey(data?)` | `CardPublicKeyEntity` | Create a CardPublicKey entity instance. |
| `Consent(data?)` | `ConsentEntity` | Create a Consent entity instance. |
| `Currency(data?)` | `CurrencyEntity` | Create a Currency entity instance. |
| `DirectDebitBatchSubmit(data?)` | `DirectDebitBatchSubmitEntity` | Create a DirectDebitBatchSubmit entity instance. |
| `FxRate(data?)` | `FxRateEntity` | Create a FxRate entity instance. |
| `IPayment(data?)` | `IPaymentEntity` | Create an IPayment entity instance. |
| `Mandate(data?)` | `MandateEntity` | Create a Mandate entity instance. |
| `Merchant(data?)` | `MerchantEntity` | Create a Merchant entity instance. |
| `MerchantAuthorisationSetting(data?)` | `MerchantAuthorisationSettingEntity` | Create a MerchantAuthorisationSetting entity instance. |
| `MerchantDirectDebitMandatePage(data?)` | `MerchantDirectDebitMandatePageEntity` | Create a MerchantDirectDebitMandatePage entity instance. |
| `MerchantPayByBankSetting(data?)` | `MerchantPayByBankSettingEntity` | Create a MerchantPayByBankSetting entity instance. |
| `MerchantPaymentRequestTemplate(data?)` | `MerchantPaymentRequestTemplateEntity` | Create a MerchantPaymentRequestTemplate entity instance. |
| `MerchantToken(data?)` | `MerchantTokenEntity` | Create a MerchantToken entity instance. |
| `Metadata(data?)` | `MetadataEntity` | Create a Metadata entity instance. |
| `NoFrixionVersion(data?)` | `NoFrixionVersionEntity` | Create a NoFrixionVersion entity instance. |
| `OpenBanking(data?)` | `OpenBankingEntity` | Create an OpenBanking entity instance. |
| `Payeeverification(data?)` | `PayeeverificationEntity` | Create a Payeeverification entity instance. |
| `Payment(data?)` | `PaymentEntity` | Create a Payment entity instance. |
| `PaymentAccount(data?)` | `PaymentAccountEntity` | Create a PaymentAccount entity instance. |
| `PaymentAccountMinimal(data?)` | `PaymentAccountMinimalEntity` | Create a PaymentAccountMinimal entity instance. |
| `PaymentInitiation(data?)` | `PaymentInitiationEntity` | Create a PaymentInitiation entity instance. |
| `PaymentRequest(data?)` | `PaymentRequestEntity` | Create a PaymentRequest entity instance. |
| `PaymentRequestEvent(data?)` | `PaymentRequestEventEntity` | Create a PaymentRequestEvent entity instance. |
| `PaymentRequestMetric(data?)` | `PaymentRequestMetricEntity` | Create a PaymentRequestMetric entity instance. |
| `PaymentRequestMinimal(data?)` | `PaymentRequestMinimalEntity` | Create a PaymentRequestMinimal entity instance. |
| `PaymentRequestResult(data?)` | `PaymentRequestResultEntity` | Create a PaymentRequestResult entity instance. |
| `Payout(data?)` | `PayoutEntity` | Create a Payout entity instance. |
| `PayoutKeysetPage(data?)` | `PayoutKeysetPageEntity` | Create a PayoutKeysetPage entity instance. |
| `PayoutMetric(data?)` | `PayoutMetricEntity` | Create a PayoutMetric entity instance. |
| `Payrun(data?)` | `PayrunEntity` | Create a Payrun entity instance. |
| `Report(data?)` | `ReportEntity` | Create a Report entity instance. |
| `ReportResult(data?)` | `ReportResultEntity` | Create a ReportResult entity instance. |
| `Role(data?)` | `RoleEntity` | Create a Role entity instance. |
| `Rule(data?)` | `RuleEntity` | Create a Rule entity instance. |
| `RuleEvent(data?)` | `RuleEventEntity` | Create a RuleEvent entity instance. |
| `Tag(data?)` | `TagEntity` | Create a Tag entity instance. |
| `Token(data?)` | `TokenEntity` | Create a Token entity instance. |
| `Transaction(data?)` | `TransactionEntity` | Create a Transaction entity instance. |
| `User(data?)` | `UserEntity` | Create an User entity instance. |
| `UserInvite(data?)` | `UserInviteEntity` | Create an UserInvite entity instance. |
| `Virtual(data?)` | `VirtualEntity` | Create a Virtual entity instance. |
| `Webhook(data?)` | `WebhookEntity` | Create a Webhook entity instance. |
| `tester(testopts?, sdkopts?)` | `NofrixionSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `NofrixionSDK.test(testopts?, sdkopts?)` | `NofrixionSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity. |
| `update` | `update(reqdata?, ctrl?): Promise<Entity>` | Update an existing entity. |
| `remove` | `remove(reqmatch?, ctrl?): Promise<void>` | Remove an entity. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): NofrixionSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load`, `create` and `update` resolve to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).
- `remove` resolves to `void`.

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

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

Operations: create, list, load, remove, update.

API path: `/api/v1/accounts/{accountID}/{currency}`

#### Batch

| Field | Description |
| --- | --- |
| `approveUrl` |  |
| `id` |  |
| `payouts` |  |

Operations: create, load.

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

Operations: create, list, load, remove, update.

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

Operations: list.

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

Operations: create.

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

Operations: list, load, remove.

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

Operations: create.

API path: `/api/v1/paymentrequests/{id}/card/refund/{partialRefundAmount}`

#### CardPublicKey

| Field | Description |
| --- | --- |
| `jwt` |  |

Operations: load.

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

Operations: create, list, load, remove, update.

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

Operations: list.

API path: `/api/v1/currencies`

#### DirectDebitBatchSubmit

| Field | Description |
| --- | --- |
| `failedSubmissions` |  |
| `successfulSubmissions` |  |

Operations: create.

API path: `/api/v1/paymentrequests/directdebit/batchsubmit`

#### FxRate

| Field | Description |
| --- | --- |
| `destinationCurrency` |  |
| `exchangeRate` |  |
| `expiryTime` |  |
| `quoteID` |  |
| `sourceCurrency` |  |

Operations: list, load.

API path: `/api/v1/payouts/fxallheldrates/{source}/{destination}`

#### IPayment

| Field | Description |
| --- | --- |
| `paymentRequestID` |  |
| `responseType` |  |

Operations: create.

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

Operations: create, load.

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

Operations: list, load, remove, update.

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

Operations: list.

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

Operations: list.

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

Operations: list.

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

Operations: list, load, remove, update.

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

Operations: create, list, load, update.

API path: `/api/v1/tokens`

#### Metadata

| Field | Description |
| --- | --- |

Operations: load.

API path: `/api/v1/metadata/problemnotification`

#### NoFrixionVersion

| Field | Description |
| --- | --- |
| `buildVersion` |  |
| `majorVersion` |  |
| `minorVersion` |  |
| `releaseName` |  |

Operations: load.

API path: `/api/v1/metadata/version`

#### OpenBanking

| Field | Description |
| --- | --- |

Operations: create, remove.

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

Operations: create.

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

Operations: create, load, update.

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

Operations: list.

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

Operations: list.

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

Operations: create.

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

Operations: create, list, load, remove, update.

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

Operations: list.

API path: `/api/v1/paymentrequests/{id}/events`

#### PaymentRequestMetric

| Field | Description |
| --- | --- |

Operations: load.

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

Operations: list.

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

Operations: list.

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

Operations: create, list, load, remove, update.

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

Operations: list.

API path: `/api/v1/accounts/{accountID}/payouts/failed`

#### PayoutMetric

| Field | Description |
| --- | --- |

Operations: load.

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

Operations: create, list, load, remove, update.

API path: `/api/v1/payruns/{id}/request-authorisation`

#### Report

| Field | Description |
| --- | --- |

Operations: update.

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

Operations: load.

API path: `/api/v1/reports/{id}/result/{statementNumber}`

#### Role

| Field | Description |
| --- | --- |
| `failedRoles` |  |
| `roles` |  |

Operations: create.

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

Operations: create, list, load, remove, update.

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

Operations: list.

API path: `/api/v1/rules/{id}/events`

#### Tag

| Field | Description |
| --- | --- |
| `colourHex` |  |
| `description` |  |
| `id` |  |
| `merchantID` |  |
| `name` |  |

Operations: create, list.

API path: `/api/v1/merchants/{merchantID}/tags`

#### Token

| Field | Description |
| --- | --- |

Operations: create, remove.

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

Operations: create, list, load, remove.

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

Operations: list, update.

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

Operations: create, list, load, remove, update.

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

Operations: create, update.

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

Operations: create, list, load, remove, update.

API path: `/api/v1/webhooks`



## Entities


### Account

Create an instance: `const account = client.Account()`

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
| `accountBalances` | `any[]` |  |
| `accountID` | `string` |  |
| `accountIdentifications` | `any[]` |  |
| `accountName` | `string` |  |
| `accountNames` | `any[]` |  |
| `accountSupplierName` | `string` |  |
| `accountType` | `string` |  |
| `availableBalance` | `number` |  |
| `availableBalanceMinorUnits` | `number` |  |
| `balance` | `number` |  |
| `balanceMinorUnits` | `number` |  |
| `bankName` | `string` |  |
| `consentID` | `string` |  |
| `consolidatedAccountInformation` | `Record<string, any>` |  |
| `createdBy` | `Record<string, any>` |  |
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
| `identifier` | `Record<string, any>` |  |
| `inserted` | `string` |  |
| `isArchived` | `boolean` |  |
| `isConnectedAccount` | `boolean` |  |
| `isDefault` | `boolean` |  |
| `isTrustAccount` | `boolean` |  |
| `isVirtual` | `boolean` |  |
| `lastTransaction` | `Record<string, any>` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantName` | `string` |  |
| `nickname` | `string` |  |
| `physicalAccountID` | `string` |  |
| `roleIDs` | `any[]` |  |
| `rules` | `any[]` |  |
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

```ts
const account = await client.Account().load({ id: 'account_id' })
```

#### Example: List

```ts
const accounts = await client.Account().list()
```

#### Example: Create

```ts
const account = await client.Account().create({
  createdBy: {},
  identifier: {},
})
```


### Batch

Create an instance: `const batch = client.Batch()`

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
| `payouts` | `any[]` |  |

#### Example: Load

```ts
const batch = await client.Batch().load({ id: 'batch_id' })
```

#### Example: Create

```ts
const batch = await client.Batch().create({
})
```


### Beneficiary

Create an instance: `const beneficiary = client.Beneficiary()`

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
| `authenticationMethods` | `any[]` |  |
| `authorisations` | `any[]` |  |
| `authorisersCompletedCount` | `number` |  |
| `authorisersRequiredCount` | `number` |  |
| `beneficiaries` | `any[]` |  |
| `beneficiaryEvents` | `any[]` |  |
| `canAuthorise` | `boolean` |  |
| `canUpdate` | `boolean` |  |
| `createdBy` | `Record<string, any>` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` |  |
| `destination` | `Record<string, any>` |  |
| `failedBeneficiaries` | `Record<string, any>` |  |
| `hasCurrentUserAuthorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isEnabled` | `boolean` |  |
| `lastAuthorised` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `sourceAccountIDs` | `any[]` |  |
| `sourceAccounts` | `any[]` |  |
| `theirReference` | `string` |  |

#### Example: Load

```ts
const beneficiary = await client.Beneficiary().load({ id: 'beneficiary_id' })
```

#### Example: List

```ts
const beneficiarys = await client.Beneficiary().list()
```

#### Example: Create

```ts
const beneficiary = await client.Beneficiary().create({
  createdBy: {},
  currency: 'example_currency',
  name: 'example_name',
})
```


### BeneficiaryGroup

Create an instance: `const beneficiary_group = client.BeneficiaryGroup()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `groupMembers` | `any[]` |  |
| `groupName` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |

#### Example: List

```ts
const beneficiary_groups = await client.BeneficiaryGroup().list({ merchant_id: "example" })
```


### Card

Create an instance: `const card = client.Card()`

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

```ts
const card = await client.Card().create({
  paymentrequest_id: 'example_paymentrequest_id',
})
```


### CardCustomerToken

Create an instance: `const card_customer_token = client.CardCustomerToken()`

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

```ts
const card_customer_token = await client.CardCustomerToken().load({ customer_email_address: 'customer_email_address' })
```

#### Example: List

```ts
const card_customer_tokens = await client.CardCustomerToken().list({ customer_email_address: "example", merchant_id: "example" })
```


### CardPayment

Create an instance: `const card_payment = client.CardPayment()`

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

```ts
const card_payment = await client.CardPayment().create({
  paymentrequest_id: 'example_paymentrequest_id',
})
```


### CardPublicKey

Create an instance: `const card_public_key = client.CardPublicKey()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `jwt` | `string` |  |

#### Example: Load

```ts
const card_public_key = await client.CardPublicKey().load({ paymentrequest_id: 'paymentrequest_id' })
```


### Consent

Create an instance: `const consent = client.Consent()`

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

```ts
const consent = await client.Consent().load({ id: 'consent_id' })
```

#### Example: List

```ts
const consents = await client.Consent().list({ email: "example", merchant_id: "example" })
```

#### Example: Create

```ts
const consent = await client.Consent().create({
})
```


### Currency

Create an instance: `const currency = client.Currency()`

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

```ts
const currencys = await client.Currency().list()
```


### DirectDebitBatchSubmit

Create an instance: `const direct_debit_batch_submit = client.DirectDebitBatchSubmit()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failedSubmissions` | `Record<string, any>` |  |
| `successfulSubmissions` | `any[]` |  |

#### Example: Create

```ts
const direct_debit_batch_submit = await client.DirectDebitBatchSubmit().create({
})
```


### FxRate

Create an instance: `const fx_rate = client.FxRate()`

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

```ts
const fx_rate = await client.FxRate().load({ destination: 'destination', source: 'source', valid_for_minute: 1 })
```

#### Example: List

```ts
const fx_rates = await client.FxRate().list({ destination: "example", source: "example" })
```


### IPayment

Create an instance: `const i_payment = client.IPayment()`

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

```ts
const i_payment = await client.IPayment().create({
})
```


### Mandate

Create an instance: `const mandate = client.Mandate()`

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

```ts
const mandate = await client.Mandate().load({ id: 'mandate_id' })
```

#### Example: Create

```ts
const mandate = await client.Mandate().create({
  addressLine1: 'example_addressLine1',
  city: 'example_city',
  countryCode: 'example_countryCode',
  emailAddress: 'example_emailAddress',
  firstName: 'example_firstName',
  lastName: 'example_lastName',
  postalCode: 'example_postalCode',
})
```


### Merchant

Create an instance: `const merchant = client.Merchant()`

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
| `accountCurrencies` | `any[]` |  |
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
| `parentMerchant` | `Record<string, any>` |  |
| `paymentAccountLimit` | `number` |  |
| `paymentAccounts` | `any[]` |  |
| `reason` | `string` |  |
| `shortName` | `string` |  |
| `supportedPaymentMethodsList` | `any[]` |  |
| `suspensionReason` | `string` |  |
| `tags` | `any[]` |  |
| `timeZoneId` | `string` |  |
| `tradingName` | `string` |  |
| `webHookLimit` | `number` |  |
| `yourRoleName` | `string` |  |

#### Example: Load

```ts
const merchant = await client.Merchant().load({ id: 'merchant_id' })
```

#### Example: List

```ts
const merchants = await client.Merchant().list()
```


### MerchantAuthorisationSetting

Create an instance: `const merchant_authorisation_setting = client.MerchantAuthorisationSetting()`

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
| `roleSettings` | `any[]` |  |

#### Example: List

```ts
const merchant_authorisation_settings = await client.MerchantAuthorisationSetting().list({ merchant_id: "example" })
```


### MerchantDirectDebitMandatePage

Create an instance: `const merchant_direct_debit_mandate_page = client.MerchantDirectDebitMandatePage()`

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

```ts
const merchant_direct_debit_mandate_pages = await client.MerchantDirectDebitMandatePage().list()
```


### MerchantPayByBankSetting

Create an instance: `const merchant_pay_by_bank_setting = client.MerchantPayByBankSetting()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bankCountryCodes` | `any[]` |  |
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

```ts
const merchant_pay_by_bank_settings = await client.MerchantPayByBankSetting().list({ merchant_id: "example" })
```


### MerchantPaymentRequestTemplate

Create an instance: `const merchant_payment_request_template = client.MerchantPaymentRequestTemplate()`

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
| `bankPaymentOptions` | `Record<string, any>` |  |
| `cardPaymentAddressOptions` | `Record<string, any>` |  |
| `cardPaymentCaptureOptions` | `Record<string, any>` |  |
| `customFields` | `any[]` |  |
| `defaultFields` | `any[]` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `name` | `string` |  |
| `notificationOptions` | `Record<string, any>` |  |
| `paymentMethods` | `Record<string, any>` |  |
| `paymentTerms` | `Record<string, any>` |  |
| `priorityBankOptions` | `Record<string, any>` |  |
| `template` | `Record<string, any>` |  |

#### Example: Load

```ts
const merchant_payment_request_template = await client.MerchantPaymentRequestTemplate().load({ id: 'merchant_payment_request_template_id', paymentrequest_id: 'paymentrequest_id' })
```

#### Example: List

```ts
const merchant_payment_request_templates = await client.MerchantPaymentRequestTemplate().list({ merchant_id: "example" })
```


### MerchantToken

Create an instance: `const merchant_token = client.MerchantToken()`

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
| `authenticationMethods` | `any[]` |  |
| `authorisations` | `any[]` |  |
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
| `permissionTypes` | `any[]` |  |
| `requestSignatureVersion` | `number` |  |
| `sharedSecretAlgorithm` | `string` |  |
| `sharedSecretBase64` | `string` |  |
| `token` | `string` |  |

#### Example: Load

```ts
const merchant_token = await client.MerchantToken().load({ id: 'merchant_token_id' })
```

#### Example: List

```ts
const merchant_tokens = await client.MerchantToken().list({ merchant_id: "example" })
```

#### Example: Create

```ts
const merchant_token = await client.MerchantToken().create({
  nonce: 'example_nonce',
})
```


### Metadata

Create an instance: `const metadata = client.Metadata()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```ts
const metadata = await client.Metadata().load()
```


### NoFrixionVersion

Create an instance: `const no_frixion_version = client.NoFrixionVersion()`

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

```ts
const no_frixion_version = await client.NoFrixionVersion().load()
```


### OpenBanking

Create an instance: `const open_banking = client.OpenBanking()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```ts
const open_banking = await client.OpenBanking().create({
  account_id: 'example_account_id',
})
```


### Payeeverification

Create an instance: `const payeeverification = client.Payeeverification()`

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

```ts
const payeeverification = await client.Payeeverification().create({
  accountName: 'example_accountName',
  iban: 'example_iban',
})
```


### Payment

Create an instance: `const payment = client.Payment()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addresses` | `any[]` |  |
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
| `createdByUser` | `Record<string, any>` |  |
| `currency` | `string` |  |
| `customFields` | `any[]` |  |
| `customerEmailAddress` | `string` |  |
| `customerID` | `string` |  |
| `customerName` | `string` |  |
| `description` | `string` |  |
| `destinationAccount` | `Record<string, any>` |  |
| `directDebitPayment` | `Record<string, any>` |  |
| `dueDate` | `string` |  |
| `events` | `any[]` |  |
| `failureCallbackUrl` | `string` |  |
| `fieldDisplaySettings` | `any[]` |  |
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
| `notificationRoleIDs` | `any[]` |  |
| `orderID` | `string` |  |
| `partialPaymentMethod` | `string` |  |
| `partialPaymentSteps` | `string` |  |
| `paymentAttempts` | `any[]` |  |
| `paymentMethods` | `any[]` |  |
| `paymentProcessor` | `string` |  |
| `payrunID` | `string` |  |
| `pispAccountID` | `string` |  |
| `priorityBankID` | `string` |  |
| `result` | `Record<string, any>` |  |
| `sandboxSettleDelayInSeconds` | `number` |  |
| `shippingAddress` | `Record<string, any>` |  |
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
| `tagIds` | `any[]` |  |
| `tags` | `any[]` |  |
| `title` | `string` |  |
| `tokenisedCards` | `any[]` |  |
| `transactions` | `any[]` |  |
| `useHostedPaymentPage` | `boolean` |  |

#### Example: Load

```ts
const payment = await client.Payment().load({ id: 'payment_id' })
```

#### Example: Create

```ts
const payment = await client.Payment().create({
  createdByUser: {},
})
```


### PaymentAccount

Create an instance: `const payment_account = client.PaymentAccount()`

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
| `createdBy` | `Record<string, any>` |  |
| `createdByDisplayName` | `string` |  |
| `currency` | `string` |  |
| `defaultPaymentRail` | `string` |  |
| `displayName` | `string` |  |
| `expiryDate` | `string` |  |
| `externalAccountIcon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `Record<string, any>` |  |
| `inserted` | `string` |  |
| `isArchived` | `boolean` |  |
| `isConnectedAccount` | `boolean` |  |
| `isDefault` | `boolean` |  |
| `isTrustAccount` | `boolean` |  |
| `isVirtual` | `boolean` |  |
| `lastTransaction` | `Record<string, any>` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantName` | `string` |  |
| `physicalAccountID` | `string` |  |
| `rules` | `any[]` |  |
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

```ts
const payment_accounts = await client.PaymentAccount().list()
```


### PaymentAccountMinimal

Create an instance: `const payment_account_minimal = client.PaymentAccountMinimal()`

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
| `identifier` | `Record<string, any>` |  |
| `isArchived` | `boolean` |  |
| `isConnectedAccount` | `boolean` |  |
| `merchantID` | `string` |  |
| `submittedPayoutsBalance` | `number` |  |

#### Example: List

```ts
const payment_account_minimals = await client.PaymentAccountMinimal().list()
```


### PaymentInitiation

Create an instance: `const payment_initiation = client.PaymentInitiation()`

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

```ts
const payment_initiation = await client.PaymentInitiation().create({
  paymentrequest_id: 'example_paymentrequest_id',
})
```


### PaymentRequest

Create an instance: `const payment_request = client.PaymentRequest()`

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
| `addresses` | `any[]` |  |
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
| `createdByUser` | `Record<string, any>` |  |
| `currency` | `string` |  |
| `customFields` | `any[]` |  |
| `customerEmailAddress` | `string` |  |
| `customerID` | `string` |  |
| `customerName` | `string` |  |
| `description` | `string` |  |
| `destinationAccount` | `Record<string, any>` |  |
| `directDebitPayment` | `Record<string, any>` |  |
| `doSimulateSettlementFailure` | `boolean` |  |
| `dueDate` | `string` |  |
| `errorDescription` | `string` |  |
| `events` | `any[]` |  |
| `failedPaymentRequests` | `Record<string, any>` |  |
| `failureCallbackUrl` | `string` |  |
| `fieldDisplaySettings` | `any[]` |  |
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
| `notificationRoleIDs` | `any[]` |  |
| `orderID` | `string` |  |
| `partialPaymentMethod` | `string` |  |
| `partialPaymentSteps` | `string` |  |
| `paymentAttempts` | `any[]` |  |
| `paymentInitiationID` | `string` |  |
| `paymentMethods` | `any[]` |  |
| `paymentProcessor` | `string` |  |
| `paymentRequests` | `any[]` |  |
| `payrunID` | `string` |  |
| `pispAccountID` | `string` |  |
| `priorityBankID` | `string` |  |
| `result` | `Record<string, any>` |  |
| `sandboxSettleDelayInSeconds` | `number` |  |
| `shippingAddress` | `Record<string, any>` |  |
| `status` | `string` |  |
| `successWebHookUrl` | `string` |  |
| `tags` | `any[]` |  |
| `title` | `string` |  |
| `tokenisedCards` | `any[]` |  |
| `transactions` | `any[]` |  |
| `useHostedPaymentPage` | `boolean` |  |

#### Example: Load

```ts
const payment_request = await client.PaymentRequest().load()
```

#### Example: List

```ts
const payment_requests = await client.PaymentRequest().list()
```

#### Example: Create

```ts
const payment_request = await client.PaymentRequest().create({
  createdByUser: {},
})
```


### PaymentRequestEvent

Create an instance: `const payment_request_event = client.PaymentRequestEvent()`

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

```ts
const payment_request_events = await client.PaymentRequestEvent().list({ paymentrequest_id: "example" })
```


### PaymentRequestMetric

Create an instance: `const payment_request_metric = client.PaymentRequestMetric()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```ts
const payment_request_metric = await client.PaymentRequestMetric().load()
```


### PaymentRequestMinimal

Create an instance: `const payment_request_minimal = client.PaymentRequestMinimal()`

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
| `customFieldsToDisplay` | `any[]` |  |
| `description` | `string` |  |
| `dueDate` | `string` |  |
| `fieldDisplaySettings` | `any[]` |  |
| `googlePayMerchantID` | `string` |  |
| `id` | `string` |  |
| `jwk` | `string` |  |
| `merchantID` | `string` |  |
| `merchantLogoUrlPng` | `string` |  |
| `merchantLogoUrlSvg` | `string` |  |
| `merchantName` | `string` |  |
| `merchantShortName` | `string` |  |
| `partialPaymentMethod` | `string` |  |
| `paymentAttempts` | `any[]` |  |
| `paymentMethodsList` | `any[]` |  |
| `paymentProcessor` | `string` |  |
| `paymentProcessorKey` | `string` |  |
| `pispError` | `string` |  |
| `priorityBankID` | `string` |  |
| `status` | `string` |  |
| `stripeAccountID` | `string` |  |
| `title` | `string` |  |

#### Example: List

```ts
const payment_request_minimals = await client.PaymentRequestMinimal().list({ paymentrequest_id: "example" })
```


### PaymentRequestResult

Create an instance: `const payment_request_result = client.PaymentRequestResult()`

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
| `payments` | `any[]` |  |
| `pispAuthorizations` | `any[]` |  |
| `requestedAmount` | `number` |  |
| `result` | `string` |  |

#### Example: List

```ts
const payment_request_results = await client.PaymentRequestResult().list({ paymentrequest_id: "example" })
```


### Payout

Create an instance: `const payout = client.Payout()`

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
| `authenticationMethods` | `any[]` |  |
| `authorisations` | `any[]` |  |
| `authorisersCompletedCount` | `number` |  |
| `authorisersRequiredCount` | `number` |  |
| `batchPayoutID` | `string` |  |
| `beneficiary` | `Record<string, any>` |  |
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
| `destination` | `Record<string, any>` |  |
| `documents` | `any[]` |  |
| `events` | `any[]` |  |
| `failedPayouts` | `Record<string, any>` |  |
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
| `payouts` | `any[]` |  |
| `payrunID` | `string` |  |
| `payrunName` | `string` |  |
| `reason` | `string` |  |
| `rule` | `Record<string, any>` |  |
| `scheduleDate` | `string` |  |
| `scheduled` | `boolean` |  |
| `sourceAccountAvailableBalance` | `number` |  |
| `sourceAccountAvailableBalanceMinorUnits` | `number` |  |
| `sourceAccountBic` | `string` |  |
| `sourceAccountCurrency` | `string` |  |
| `sourceAccountIban` | `string` |  |
| `sourceAccountIdentifier` | `Record<string, any>` |  |
| `sourceAccountName` | `string` |  |
| `sourceAccountNumber` | `string` |  |
| `sourceAccountSortcode` | `string` |  |
| `status` | `string` |  |
| `tagIds` | `any[]` |  |
| `tags` | `any[]` |  |
| `theirReference` | `string` |  |
| `topupPayrunID` | `string` |  |
| `transactedAmount` | `number` |  |
| `transactedFxAmount` | `number` |  |
| `transactedFxRate` | `number` |  |
| `type` | `string` |  |
| `userID` | `string` |  |
| `yourReference` | `string` |  |

#### Example: Load

```ts
const payout = await client.Payout().load({ id: 'payout_id' })
```

#### Example: List

```ts
const payouts = await client.Payout().list()
```

#### Example: Create

```ts
const payout = await client.Payout().create({
  beneficiary: {},
  sourceAccountIdentifier: {},
})
```


### PayoutKeysetPage

Create an instance: `const payout_keyset_page = client.PayoutKeysetPage()`

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
| `authenticationMethods` | `any[]` |  |
| `authorisations` | `any[]` |  |
| `authorisersCompletedCount` | `number` |  |
| `authorisersRequiredCount` | `number` |  |
| `batchPayoutID` | `string` |  |
| `beneficiary` | `Record<string, any>` |  |
| `canAuthorise` | `boolean` |  |
| `canProcess` | `boolean` |  |
| `canUpdate` | `boolean` |  |
| `chargeBearer` | `string` |  |
| `createdBy` | `string` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` |  |
| `currentUserID` | `string` |  |
| `description` | `string` |  |
| `destination` | `Record<string, any>` |  |
| `documents` | `any[]` |  |
| `events` | `any[]` |  |
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
| `rule` | `Record<string, any>` |  |
| `scheduleDate` | `string` |  |
| `scheduled` | `boolean` |  |
| `sourceAccountAvailableBalance` | `number` |  |
| `sourceAccountAvailableBalanceMinorUnits` | `number` |  |
| `sourceAccountBic` | `string` |  |
| `sourceAccountCurrency` | `string` |  |
| `sourceAccountIban` | `string` |  |
| `sourceAccountIdentifier` | `Record<string, any>` |  |
| `sourceAccountName` | `string` |  |
| `sourceAccountNumber` | `string` |  |
| `sourceAccountSortcode` | `string` |  |
| `status` | `string` |  |
| `tags` | `any[]` |  |
| `theirReference` | `string` |  |
| `topupPayrunID` | `string` |  |
| `transactedAmount` | `number` |  |
| `transactedFxAmount` | `number` |  |
| `transactedFxRate` | `number` |  |
| `type` | `string` |  |
| `userID` | `string` |  |
| `yourReference` | `string` |  |

#### Example: List

```ts
const payout_keyset_pages = await client.PayoutKeysetPage().list()
```


### PayoutMetric

Create an instance: `const payout_metric = client.PayoutMetric()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Example: Load

```ts
const payout_metric = await client.PayoutMetric().load()
```


### Payrun

Create an instance: `const payrun = client.Payrun()`

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
| `authorisations` | `any[]` |  |
| `authorisersCompletedCount` | `number` |  |
| `authorisersRequiredCount` | `number` |  |
| `batchPayoutID` | `string` |  |
| `canAuthorise` | `boolean` |  |
| `canDelete` | `boolean` |  |
| `canEdit` | `boolean` |  |
| `events` | `any[]` |  |
| `hasCurrentUserAuthorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoices` | `any[]` |  |
| `invoicesMinimal` | `any[]` |  |
| `isArchived` | `boolean` |  |
| `lastUpdated` | `string` |  |
| `lastUpdatedBy` | `Record<string, any>` |  |
| `merchantID` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `notes` | `string` |  |
| `payments` | `any[]` |  |
| `payouts` | `any[]` |  |
| `payoutsCount` | `number` |  |
| `reason` | `string` |  |
| `scheduleDate` | `string` |  |
| `scheduledDate` | `string` |  |
| `sourceAccounts` | `any[]` |  |
| `status` | `string` |  |
| `totalEur` | `number` |  |
| `totalGbp` | `number` |  |
| `totalUsd` | `number` |  |

#### Example: Load

```ts
const payrun = await client.Payrun().load({ id: 'payrun_id' })
```

#### Example: List

```ts
const payruns = await client.Payrun().list()
```

#### Example: Create

```ts
const payrun = await client.Payrun().create({
  id: 'example_id',
  lastUpdatedBy: {},
})
```


### Report

Create an instance: `const report = client.Report()`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |


### ReportResult

Create an instance: `const report_result = client.ReportResult()`

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

```ts
const report_result = await client.ReportResult().load({ id: 1, report_id: 'report_id' })
```


### Role

Create an instance: `const role = client.Role()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failedRoles` | `Record<string, any>` |  |
| `roles` | `any[]` |  |

#### Example: Create

```ts
const role = await client.Role().create({
  merchant_id: 'example_merchant_id',
})
```


### Rule

Create an instance: `const rule = client.Rule()`

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
| `account` | `Record<string, any>` |  |
| `accountID` | `string` |  |
| `approveUrl` | `string` |  |
| `approverID` | `string` |  |
| `authenticationMethods` | `any[]` |  |
| `authorisations` | `any[]` |  |
| `authorisersCompletedCount` | `number` |  |
| `authorisersRequiredCount` | `number` |  |
| `canAuthorise` | `boolean` |  |
| `createdBy` | `Record<string, any>` |  |
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
| `sweepAction` | `Record<string, any>` |  |
| `timeZoneId` | `string` |  |
| `triggerCronExpression` | `string` |  |
| `triggerOnPayIn` | `boolean` |  |
| `userID` | `string` |  |
| `webHookSecret` | `string` |  |

#### Example: Load

```ts
const rule = await client.Rule().load({ id: 'rule_id' })
```

#### Example: List

```ts
const rules = await client.Rule().list()
```

#### Example: Create

```ts
const rule = await client.Rule().create({
  createdBy: {},
  nonce: 'example_nonce',
})
```


### RuleEvent

Create an instance: `const rule_event = client.RuleEvent()`

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
| `user` | `Record<string, any>` |  |

#### Example: List

```ts
const rule_events = await client.RuleEvent().list({ id: "example" })
```


### Tag

Create an instance: `const tag = client.Tag()`

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

```ts
const tags = await client.Tag().list({ merchant_id: "example" })
```

#### Example: Create

```ts
const tag = await client.Tag().create({
  merchant_id: 'example_merchant_id',
  merchantID: 'example_merchantID',
  name: 'example_name',
})
```


### Token

Create an instance: `const token = client.Token()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `remove(match)` | Remove the matching entity. |

#### Example: Create

```ts
const token = await client.Token().create({
  id: 'example_id',
})
```


### Transaction

Create an instance: `const transaction = client.Transaction()`

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
| `addressDetails` | `Record<string, any>` |  |
| `amount` | `number` |  |
| `amountMinorUnits` | `number` |  |
| `balance` | `number` |  |
| `balanceMinorUnits` | `number` |  |
| `bookingDateTime` | `string` |  |
| `chargeDetails` | `Record<string, any>` |  |
| `content` | `any[]` |  |
| `counterparty` | `Record<string, any>` |  |
| `counterpartySummary` | `string` |  |
| `currency` | `string` |  |
| `currencyExchange` | `Record<string, any>` |  |
| `date` | `string` |  |
| `description` | `string` |  |
| `enrichment` | `Record<string, any>` |  |
| `fxAmount` | `number` |  |
| `fxCurrency` | `string` |  |
| `fxRate` | `number` |  |
| `grossAmount` | `Record<string, any>` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isoBankTransactionCode` | `Record<string, any>` |  |
| `merchant` | `Record<string, any>` |  |
| `merchantID` | `string` |  |
| `pageNumber` | `number` |  |
| `pageSize` | `number` |  |
| `payeeDetails` | `Record<string, any>` |  |
| `payerDetails` | `Record<string, any>` |  |
| `paymentRequestCustomFields` | `Record<string, any>` |  |
| `paymentRequestID` | `string` |  |
| `payoutID` | `string` |  |
| `proprietaryBankTransactionCode` | `Record<string, any>` |  |
| `rawReference` | `string` |  |
| `reference` | `string` |  |
| `ruleID` | `string` |  |
| `statementReferences` | `any[]` |  |
| `status` | `string` |  |
| `supplementaryData` | `any` |  |
| `tags` | `any[]` |  |
| `theirReference` | `string` |  |
| `totalPages` | `number` |  |
| `totalSize` | `number` |  |
| `transactionAmount` | `Record<string, any>` |  |
| `transactionDate` | `string` |  |
| `transactionInformation` | `any[]` |  |
| `transactionMutability` | `string` |  |
| `type` | `string` |  |
| `valueDateTime` | `string` |  |
| `virtualIBAN` | `string` |  |
| `yourReference` | `string` |  |

#### Example: Load

```ts
const transaction = await client.Transaction().load({ id: 'transaction_id' })
```

#### Example: List

```ts
const transactions = await client.Transaction().list()
```

#### Example: Create

```ts
const transaction = await client.Transaction().create({
  id: 'example_id',
  grossAmount: {},
  payeeDetails: {},
  payerDetails: {},
  transactionAmount: {},
})
```


### User

Create an instance: `const user = client.User()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `clientSessionTimeouts` | `any[]` |  |
| `emailAddress` | `string` |  |
| `firstName` | `string` |  |
| `id` | `string` |  |
| `lastName` | `string` |  |
| `passkeyAdded` | `boolean` |  |
| `permissions` | `Record<string, any>` |  |
| `profile` | `string` |  |
| `rolesWithScope` | `any[]` |  |
| `twoFactorEnabled` | `boolean` |  |
| `userInviteID` | `string` |  |

#### Example: List

```ts
const users = await client.User().list()
```


### UserInvite

Create an instance: `const user_invite = client.UserInvite()`

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
| `authorisationStatus` | `Record<string, any>` |  |
| `failedUserInvites` | `Record<string, any>` |  |
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
| `user` | `Record<string, any>` |  |
| `userID` | `string` |  |
| `userInvites` | `any[]` |  |

#### Example: Load

```ts
const user_invite = await client.UserInvite().load({ id: 'user_invite_id' })
```

#### Example: List

```ts
const user_invites = await client.UserInvite().list({ merchant_id: "example" })
```

#### Example: Create

```ts
const user_invite = await client.UserInvite().create({
  user: {},
})
```


### Virtual

Create an instance: `const virtual = client.Virtual()`

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
| `createdBy` | `Record<string, any>` |  |
| `createdByDisplayName` | `string` |  |
| `currency` | `string` |  |
| `defaultPaymentRail` | `string` |  |
| `displayName` | `string` |  |
| `expiryDate` | `string` |  |
| `externalAccountIcon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `Record<string, any>` |  |
| `inserted` | `string` |  |
| `isArchived` | `boolean` |  |
| `isConnectedAccount` | `boolean` |  |
| `isDefault` | `boolean` |  |
| `isTrustAccount` | `boolean` |  |
| `isVirtual` | `boolean` |  |
| `lastTransaction` | `Record<string, any>` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantName` | `string` |  |
| `name` | `string` |  |
| `physicalAccountID` | `string` |  |
| `rules` | `any[]` |  |
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

```ts
const virtual = await client.Virtual().create({
  account_id: 'example_account_id',
  createdBy: {},
  identifier: {},
  name: 'example_name',
})
```


### Webhook

Create an instance: `const webhook = client.Webhook()`

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
| `resourceTypes` | `any[]` |  |
| `retry` | `boolean` |  |
| `secret` | `string` |  |
| `version` | `number` |  |

#### Example: Load

```ts
const webhook = await client.Webhook().load({ id: 'webhook_id' })
```

#### Example: List

```ts
const webhooks = await client.Webhook().list({ merchant_id: "example" })
```

#### Example: Create

```ts
const webhook = await client.Webhook().create({
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

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
nofrixion/
├── src/
│   ├── NofrixionSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { NofrixionSDK } from '@voxgig-sdk/nofrixion'
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const consent = client.Consent()
await consent.list()

// consent.data() now returns the consent data from the last `list`
// consent.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
