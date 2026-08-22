# Nofrixion TypeScript SDK



The TypeScript SDK for the Nofrixion API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.Account()` — each with a small set of operations (`list`, `load`, `create`, `update`, `remove`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Also generated from this model: `go`, `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb` — see
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
  account_id: 'example_account_id',
  currency: 'example_currency',
  createdBy: {},
  identifier: {},
})

// Update — the id comes off the returned entity's data()
const updated = await client.Account().update({
  id: created.data().id!,
  accountBalances: [],
  accountID: 'example_accountID',
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

Operations: create, list, load, remove, update.

API path: `/api/v1/accounts/{accountID}/{currency}`

#### Batch

| Field | Description |
| --- | --- |
| `approveUrl` | This field is used when returning a batch payout record to a client. |
| `id` |  |
| `payouts` |  |

Operations: create, load.

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

Operations: create, list, load, remove, update.

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

Operations: list.

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

Operations: create.

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

Operations: list, load, remove.

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
| `failedSubmissions` | Dictionary of failed submissions, keyed by the index (1-based) in the original request. |
| `successfulSubmissions` | List of successfully submitted direct debit payments. |

Operations: create.

API path: `/api/v1/paymentrequests/directdebit/batchsubmit`

#### FxRate

| Field | Description |
| --- | --- |
| `destinationCurrency` |  |
| `exchangeRate` | The price at which the transaction will buy the source currency using the destination currency. |
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

Operations: create, load.

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

Operations: list.

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

Operations: list.

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

Operations: list, load, remove, update.

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
| `accountName` | The name of the account to verify |
| `accountNumber` | The account number of the account to verify (for CoP checks) |
| `iban` | The IBAN of the account to verify (for VoP checks) |
| `payeeVerifiedAccountName` | The verified account name of the payee, if available (in case of a close match) |
| `result` | The result of the payee verification |
| `secondaryIdentification` | Optional secondary identifier for the account to verify. |
| `sortCode` | The sort code of the account to verify (for CoP checks) |

Operations: create.

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

Operations: create, load, update.

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

Operations: list.

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

Operations: list.

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

Operations: create.

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

Operations: create, list, load, remove, update.

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

Operations: list.

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

Operations: list.

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

Operations: create, list, load, remove, update.

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

Operations: create, list, load, remove.

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

Operations: list, update.

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

Operations: create, list, load, remove, update.

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

Operations: create, update.

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
| `accountBalances` | `any[]` | The various balances for the account. |
| `accountID` | `string` | ID of the account. |
| `accountIdentifications` | `any[]` | The canoncial identifiers for the account. |
| `accountName` | `string` | Name for the account |
| `accountNames` | `any[]` | Optional account names set by the account holder. |
| `accountSupplierName` | `string` | The payment account supplier name. |
| `accountType` | `string` | The type of account e.g. |
| `availableBalance` | `number` | The current available balance of the account. |
| `availableBalanceMinorUnits` | `number` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `number` | Balance of the account. |
| `balanceMinorUnits` | `number` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `string` | The bank name for external accounts |
| `consentID` | `string` | The ID of the consent used to connect the external account. |
| `consolidatedAccountInformation` | `Record<string, any>` | Summary information regarding account balances of the overall account provided by the bank. |
| `createdBy` | `Record<string, any>` |  |
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
| `identifier` | `Record<string, any>` |  |
| `inserted` | `string` | Timestamp when the account was created. |
| `isArchived` | `boolean` | Indicates whether the account is archived. |
| `isConnectedAccount` | `boolean` | Indicates if the payment account is an externally connected account. |
| `isDefault` | `boolean` | Is the default account |
| `isTrustAccount` | `boolean` | Indicates if the payment account is a trust account. |
| `isVirtual` | `boolean` | True if the account is a virtual account. |
| `lastTransaction` | `Record<string, any>` |  |
| `lastUpdated` | `string` | Timestamp when the account was last updated. |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `merchantName` | `string` | The name of the merchant that owns the account. |
| `nickname` | `string` | Nickname of the account that was provided by the account owner. |
| `physicalAccountID` | `string` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `roleIDs` | `any[]` | Optional list of role IDs that will get access to the payment account when created. |
| `rules` | `any[]` | The list of rules associated with this account. |
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
  account_id: 'example_account_id',
  currency: 'example_currency',
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
| `approveUrl` | `string` | This field is used when returning a batch payout record to a client. |
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
| `authenticationMethods` | `any[]` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `any[]` | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `number` | The number of distinct authorisers that have authorised the beneficiary. |
| `authorisersRequiredCount` | `number` | The number of authorisers required for this beneficiary. |
| `beneficiaries` | `any[]` |  |
| `beneficiaryEvents` | `any[]` |  |
| `canAuthorise` | `boolean` | True if the beneficiary can be authorised by the user who loaded it. |
| `canUpdate` | `boolean` | True if the beneficiary can be updated by the user who loaded it. |
| `createdBy` | `Record<string, any>` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` | Gets or Sets the currency. |
| `destination` | `Record<string, any>` |  |
| `failedBeneficiaries` | `Record<string, any>` |  |
| `hasCurrentUserAuthorised` | `boolean` | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isEnabled` | `boolean` |  |
| `lastAuthorised` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` | Gets or Sets the merchant id. |
| `name` | `string` | The descriptive name for the beneficiary. |
| `nonce` | `string` |  |
| `sourceAccountIDs` | `any[]` | ID of the accounts which are authorised to act as a source for the beneficiary. |
| `sourceAccounts` | `any[]` |  |
| `theirReference` | `string` | The reference that will be used by default as TheirReference when creating payouts to this beneficiary if no TheirReference is specified for the payout. |

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
  id: 'example_id',
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
| `groupMembers` | `any[]` | The existing group members. |
| `groupName` | `string` | The descriptive name for the beneficiary group. |
| `id` | `string` |  |
| `inserted` | `string` | Timestamp indicating when the group was created. |
| `lastUpdated` | `string` | Timestamp indicating when the group was last updated. |
| `merchantID` | `string` | Gets or Sets the merchant id. |

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
| `failedSubmissions` | `Record<string, any>` | Dictionary of failed submissions, keyed by the index (1-based) in the original request. |
| `successfulSubmissions` | `any[]` | List of successfully submitted direct debit payments. |

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
| `exchangeRate` | `number` | The price at which the transaction will buy the source currency using the destination currency. |
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
| `accountCurrencies` | `any[]` | The list of currencies that the merchant has accounts for. |
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
| `parentMerchant` | `Record<string, any>` |  |
| `paymentAccountLimit` | `number` | The maximum number of payment accounts that can be created for the Merchant. |
| `paymentAccounts` | `any[]` |  |
| `reason` | `string` | The reason for the suspension. |
| `shortName` | `string` | A URL friendly shortish name for the merchant. |
| `supportedPaymentMethodsList` | `any[]` | The payment methods that are configured and supported for this merchant. |
| `suspensionReason` | `string` | The reason for the suspension, provided by the merchant. |
| `tags` | `any[]` | An optional list of descriptive tags that can be used on merchant entities such as payment requests. |
| `timeZoneId` | `string` | The IANA (Internet Assigned Numbers Authority) time zone identifier of the merchant. |
| `tradingName` | `string` | An optional trading name. |
| `webHookLimit` | `number` | The maximum number of web hooks that can be created for the Merchant. |
| `yourRoleName` | `string` | The name of the role for the identity that loaded the merchant record. |

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
| `bankCountryCodes` | `any[]` | The list of country codes representing the banks the country supports. |
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
| `customFields` | `any[]` | A list of custom fields that can be included in the payment request template. |
| `defaultFields` | `any[]` | A list of default fields that are included in the payment request template. |
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
| `authenticationMethods` | `any[]` | A list of authentication types allowed to authorise the merchant token. |
| `authorisations` | `any[]` | A list of users who have successfully authorised the latest version of the beneficiary. |
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
| `permissionTypes` | `any[]` | The permissions that the merchant token supports. |
| `requestSignatureVersion` | `number` | Represent the version of the overall merchant token. |
| `sharedSecretAlgorithm` | `string` | Optional shared secret algorithm to use for HMAC authentication. |
| `sharedSecretBase64` | `string` | The base 64 encoded shared secret that is used for request authentication with an HMAC. |
| `token` | `string` | The JWT merchant token. |

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
| `accountName` | `string` | The name of the account to verify |
| `accountNumber` | `string` | The account number of the account to verify (for CoP checks) |
| `iban` | `string` | The IBAN of the account to verify (for VoP checks) |
| `payeeVerifiedAccountName` | `string` | The verified account name of the payee, if available (in case of a close match) |
| `result` | `string` | The result of the payee verification |
| `secondaryIdentification` | `string` | Optional secondary identifier for the account to verify. |
| `sortCode` | `string` | The sort code of the account to verify (for CoP checks) |

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
| `createdByUser` | `Record<string, any>` |  |
| `currency` | `string` | The currency of the request. |
| `customFields` | `any[]` | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `string` | Optional email address for the customer. |
| `customerID` | `string` | An optional customer identifier for the payment request. |
| `customerName` | `string` |  |
| `description` | `string` | An optional description for the payment request. |
| `destinationAccount` | `Record<string, any>` |  |
| `directDebitPayment` | `Record<string, any>` | Contains information about a Direct Debit payment attempt for a payment request. |
| `dueDate` | `string` | The due date for the payment request. |
| `events` | `any[]` |  |
| `failureCallbackUrl` | `string` | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `any[]` | A list of field display settings that control which fields are displayed to the payer. |
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
| `notificationRoleIDs` | `any[]` | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `string` | An optional order ID for the payment request. |
| `partialPaymentMethod` | `string` | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `string` | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `any[]` | The payment attempts made against this payment request. |
| `paymentMethods` | `any[]` | The payment methods that the payment request supports. |
| `paymentProcessor` | `string` | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `payrunID` | `string` | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `string` | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `string` | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `Record<string, any>` |  |
| `sandboxSettleDelayInSeconds` | `number` | Sandbox only. |
| `shippingAddress` | `Record<string, any>` |  |
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
| `tagIds` | `any[]` | An optional list of tag ids to add to the payment request |
| `tags` | `any[]` | An optional list of descriptive tags attached to the payment request. |
| `title` | `string` | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `any[]` |  |
| `transactions` | `any[]` |  |
| `useHostedPaymentPage` | `boolean` | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

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
| `accountName` | `string` | Name for the account |
| `accountSupplierName` | `string` | The payment account supplier name. |
| `availableBalance` | `number` | The current available balance of the account. |
| `availableBalanceMinorUnits` | `number` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `number` | Balance of the account. |
| `balanceMinorUnits` | `number` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `string` | The bank name for external accounts |
| `consentID` | `string` | The ID of the consent used to connect the external account. |
| `createdBy` | `Record<string, any>` |  |
| `createdByDisplayName` | `string` | Either the name of the user, merchant token or api key that created the account |
| `currency` | `string` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `string` | Indicates the default payment rail for this account. |
| `displayName` | `string` | Gets a unique display name for the payment account. |
| `expiryDate` | `string` | The date that the external account will expire |
| `externalAccountIcon` | `string` | The Icon for external accounts |
| `id` | `string` | Unique id for the account. |
| `identifier` | `Record<string, any>` |  |
| `inserted` | `string` | Timestamp when the account was created. |
| `isArchived` | `boolean` | Indicates whether the account is archived. |
| `isConnectedAccount` | `boolean` | Indicates if the payment account is an externally connected account. |
| `isDefault` | `boolean` | Is the default account |
| `isTrustAccount` | `boolean` | Indicates if the payment account is a trust account. |
| `isVirtual` | `boolean` | True if the account is a virtual account. |
| `lastTransaction` | `Record<string, any>` |  |
| `lastUpdated` | `string` | Timestamp when the account was last updated. |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `merchantName` | `string` | The name of the merchant that owns the account. |
| `physicalAccountID` | `string` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `any[]` | The list of rules associated with this account. |
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
| `accountName` | `string` | Name for the account |
| `availableBalance` | `number` | The current available balance of the account. |
| `balance` | `number` | Balance of the account. |
| `balanceMinorUnits` | `number` | Balance of the account expressed in the currency’s minor units (e.g. |
| `currency` | `string` | Currency of the account in ISO 4217 format |
| `id` | `string` | Unique id for the account. |
| `identifier` | `Record<string, any>` |  |
| `isArchived` | `boolean` | Is the account archived |
| `isConnectedAccount` | `boolean` | Indicates if the payment account is an externally connected account. |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `submittedPayoutsBalance` | `number` | Total of the payouts that have been submitted for processing. |

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
| `paymentInitiationID` | `string` | The unique identifier of the payment initiation request. |
| `paymentRequestCallbackUrl` | `string` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `string` |  |
| `redirectUrl` | `string` | A redirect URL for the user to authorise the payment initiation request at the ASPSP |
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
| `createdByUser` | `Record<string, any>` |  |
| `currency` | `string` | The currency of the request. |
| `customFields` | `any[]` | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `string` | Optional email address for the customer. |
| `customerID` | `string` | An optional customer identifier for the payment request. |
| `customerName` | `string` |  |
| `description` | `string` | An optional description for the payment request. |
| `destinationAccount` | `Record<string, any>` |  |
| `directDebitPayment` | `Record<string, any>` | Contains information about a Direct Debit payment attempt for a payment request. |
| `doSimulateSettlementFailure` | `boolean` |  |
| `dueDate` | `string` | The due date for the payment request. |
| `errorDescription` | `string` |  |
| `events` | `any[]` |  |
| `failedPaymentRequests` | `Record<string, any>` |  |
| `failureCallbackUrl` | `string` | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `any[]` | A list of field display settings that control which fields are displayed to the payer. |
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
| `notificationRoleIDs` | `any[]` | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `string` | An optional order ID for the payment request. |
| `partialPaymentMethod` | `string` | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `string` | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `any[]` | The payment attempts made against this payment request. |
| `paymentInitiationID` | `string` |  |
| `paymentMethods` | `any[]` | The payment methods that the payment request supports. |
| `paymentProcessor` | `string` | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `paymentRequests` | `any[]` |  |
| `payrunID` | `string` | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `string` | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `string` | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `Record<string, any>` |  |
| `sandboxSettleDelayInSeconds` | `number` | Sandbox only. |
| `shippingAddress` | `Record<string, any>` |  |
| `status` | `string` | The current status of the payment request. |
| `successWebHookUrl` | `string` | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tags` | `any[]` | An optional list of descriptive tags attached to the payment request. |
| `title` | `string` | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `any[]` |  |
| `transactions` | `any[]` |  |
| `useHostedPaymentPage` | `boolean` | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

#### Example: Load

```ts
const payment_request = await client.PaymentRequest().load({ id: 'payment_request_id' })
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
| `amount` | `number` | The amount of money to request. |
| `amountPending` | `number` | The amount of money that was authorised but has not arrived in the account yet. |
| `amountReceived` | `number` | The amount of money that has been received for this payment request. |
| `amountRefunded` | `number` | The amount of money that has been refunded for this payment request. |
| `callbackUrl` | `string` |  |
| `cardStripePaymentIntentSecret` | `string` |  |
| `countryCode` | `string` | The country code associated with the payment. |
| `currency` | `string` | The currency of the request. |
| `customFieldsToDisplay` | `any[]` | Custom fields to display to the customer. |
| `description` | `string` | An optional description for the payment request. |
| `dueDate` | `string` | The due date of the payment request. |
| `fieldDisplaySettings` | `any[]` |  |
| `googlePayMerchantID` | `string` | Merchant ID from Google Pay |
| `id` | `string` |  |
| `jwk` | `string` | The jwk containing the public key |
| `merchantID` | `string` |  |
| `merchantLogoUrlPng` | `string` |  |
| `merchantLogoUrlSvg` | `string` |  |
| `merchantName` | `string` |  |
| `merchantShortName` | `string` |  |
| `partialPaymentMethod` | `string` |  |
| `paymentAttempts` | `any[]` | The payment attempts for this payment request. |
| `paymentMethodsList` | `any[]` | The payment methods that the payment request supports. |
| `paymentProcessor` | `string` | The card processor |
| `paymentProcessorKey` | `string` | The card processors public key |
| `pispError` | `string` | This is the error returned from the bank which is recorded in payment request events. |
| `priorityBankID` | `string` |  |
| `status` | `string` | The status of the payment request. |
| `stripeAccountID` | `string` | Account ID of connected customers in Stripe |
| `title` | `string` | The title of the payment request. |

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
| `amount` | `number` | The authorised payment amount. |
| `amountPending` | `number` |  |
| `amountReceived` | `number` |  |
| `amountRefunded` | `number` |  |
| `currency` | `string` | The authorised payment currency. |
| `customerID` | `string` | The customer id |
| `paymentRequestID` | `string` | The ID of the payment request the result is for. |
| `payments` | `any[]` | The list of payment attempts that have been received for the payment request. |
| `pispAuthorizations` | `any[]` |  |
| `requestedAmount` | `number` | The full original payment amount requested. |
| `result` | `string` | The result of the payment attempt. |

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
| `accountID` | `string` | Gets or Sets Account Id of sending account |
| `allowIncomplete` | `boolean` | If set to true the payout will get created even if the business validation rules fail. |
| `amount` | `number` | Gets or Sets payout amount |
| `amountMinorUnits` | `number` | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `string` | This field is used when returning an payout record to a client. |
| `approverID` | `string` | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `any[]` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `any[]` | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `number` | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `number` | The number of authorisers required for this payout. |
| `batchPayoutID` | `string` | The ID of the batch the payout is associated with. |
| `beneficiary` | `Record<string, any>` |  |
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
| `destination` | `Record<string, any>` |  |
| `documents` | `any[]` | Documents associated with the payout. |
| `events` | `any[]` | The activity associated with the payout. |
| `failedPayouts` | `Record<string, any>` |  |
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
| `payouts` | `any[]` |  |
| `payrunID` | `string` | The ID of the payrun that this payout is associated with. |
| `payrunName` | `string` | The name of the payrun that this payout is associated with. |
| `reason` | `string` |  |
| `rule` | `Record<string, any>` |  |
| `scheduleDate` | `string` | The date the payout should be submitted. |
| `scheduled` | `boolean` | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `number` | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `number` | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `string` | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `string` | The currency of the source account. |
| `sourceAccountIban` | `string` | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `Record<string, any>` |  |
| `sourceAccountName` | `string` | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `string` | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `string` | The sort code of the account the payout is being made from. |
| `status` | `string` | Gets or Sets the status of payout request |
| `tagIds` | `any[]` | An optional list of tag ids to add to the payout. |
| `tags` | `any[]` | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `string` | Gets or Sets destination reference ID |
| `topupPayrunID` | `string` | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `number` | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `number` | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `number` | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `string` | Gets or Sets payout type |
| `userID` | `string` | Gets or Sets User ID of who created the payout request |
| `yourReference` | `string` | Gets or Sets your reference ID |

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
  id: 'example_id',
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
| `accountID` | `string` | Gets or Sets Account Id of sending account |
| `amount` | `number` | Gets or Sets payout amount |
| `amountMinorUnits` | `number` | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `string` | This field is used when returning an payout record to a client. |
| `approverID` | `string` | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `any[]` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `any[]` | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `number` | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `number` | The number of authorisers required for this payout. |
| `batchPayoutID` | `string` | The ID of the batch the payout is associated with. |
| `beneficiary` | `Record<string, any>` |  |
| `canAuthorise` | `boolean` | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `boolean` | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `boolean` | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `string` | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `string` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` | Gets or Sets Currency of payout request |
| `currentUserID` | `string` | The ID of the user that requested access to the PayOut record. |
| `description` | `string` | Gets or Sets description of payout request |
| `destination` | `Record<string, any>` |  |
| `documents` | `any[]` | Documents associated with the payout. |
| `events` | `any[]` | The activity associated with the payout. |
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
| `rule` | `Record<string, any>` |  |
| `scheduleDate` | `string` | The date the payout should be submitted. |
| `scheduled` | `boolean` | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `number` | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `number` | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `string` | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `string` | The currency of the source account. |
| `sourceAccountIban` | `string` | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `Record<string, any>` |  |
| `sourceAccountName` | `string` | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `string` | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `string` | The sort code of the account the payout is being made from. |
| `status` | `string` | Gets or Sets the status of payout request |
| `tags` | `any[]` | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `string` | Gets or Sets destination reference ID |
| `topupPayrunID` | `string` | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `number` | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `number` | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `number` | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `string` | Gets or Sets payout type |
| `userID` | `string` | Gets or Sets User ID of who created the payout request |
| `yourReference` | `string` | Gets or Sets your reference ID |

#### Example: List

```ts
const payout_keyset_pages = await client.PayoutKeysetPage().list({ merchant_id: "example" })
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
| `authorisations` | `any[]` | A list of the users who have successfully authorised the latest version of the payrun and when. |
| `authorisersCompletedCount` | `number` | The number of distinct authorisers that have authorised the payrun. |
| `authorisersRequiredCount` | `number` | The number of authorisers required for this payrun. |
| `batchPayoutID` | `string` |  |
| `canAuthorise` | `boolean` | True if the payrun can be authorised by the user who loaded it. |
| `canDelete` | `boolean` |  |
| `canEdit` | `boolean` |  |
| `events` | `any[]` |  |
| `hasCurrentUserAuthorised` | `boolean` | True if the payrun was loaded for a user and that user has already authorised the latest version of the payrun. |
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
| `accountID` | `string` | The ID of the account the rule will apply to. |
| `approveUrl` | `string` | If set this property holds the URL an approver needs to visit in order to complete a strong authentication check in order to approve the rule. |
| `approverID` | `string` |  |
| `authenticationMethods` | `any[]` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `any[]` | A list of the users who have successfully authorised the latest version of the rule and when. |
| `authorisersCompletedCount` | `number` | The number of distinct authorisers that have authorised the rule. |
| `authorisersRequiredCount` | `number` | The number of authorisers required for this rule. |
| `canAuthorise` | `boolean` | True if the rule can be authorised by the user who loaded it. |
| `createdBy` | `Record<string, any>` |  |
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
| `sweepAction` | `Record<string, any>` |  |
| `timeZoneId` | `string` | If the rule should be executed on a recurring schedule this is the timezone that the CRON expression should be evaluated in. |
| `triggerCronExpression` | `string` | If the rule should be executed on a recurring schedule this is the expression that sets the schedule. |
| `triggerOnPayIn` | `boolean` | Set to true if the rule execution should be triggered when the account receives a pay in (credit). |
| `userID` | `string` |  |
| `webHookSecret` | `string` | If set this secret will be used to sign Web Hook requests. |

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
| `accountID` | `string` | The ID of the account the transaction belongs to. |
| `accountName` | `string` | The name of the account the transaction belongs to. |
| `accountSequenceNumber` | `number` | The sequence number of transaction on a per account basis. |
| `addressDetails` | `Record<string, any>` |  |
| `amount` | `number` | Amount of the transaction. |
| `amountMinorUnits` | `number` | Amount of the transaction expressed in the currency’s minor units (e.g. |
| `balance` | `number` | Balance left on the account after the transaction. |
| `balanceMinorUnits` | `number` | Balance on the account expressed in the currency’s minor units (e.g. |
| `bookingDateTime` | `string` |  |
| `chargeDetails` | `Record<string, any>` |  |
| `content` | `any[]` |  |
| `counterparty` | `Record<string, any>` |  |
| `counterpartySummary` | `string` | For pay in (credit) transactions this will contain a descriptive string with the most important fields about the counterparty. |
| `currency` | `string` | Currency of transaction. |
| `currencyExchange` | `Record<string, any>` | Provides details on the currency exchange. |
| `date` | `string` |  |
| `description` | `string` | Description of the transaction. |
| `enrichment` | `Record<string, any>` |  |
| `fxAmount` | `number` | For an FX payout this is the amound in the FX currency. |
| `fxCurrency` | `string` | For an FX payout this is the currency that was received or that was instructed. |
| `fxRate` | `number` | For an FX payout this is the exchange rate between the transaction currency and the FX currency. |
| `grossAmount` | `Record<string, any>` |  |
| `id` | `string` | Unique ID for the transaction. |
| `inserted` | `string` | Date when the transaction was inserted into the ledger. |
| `isoBankTransactionCode` | `Record<string, any>` |  |
| `merchant` | `Record<string, any>` |  |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `pageNumber` | `number` | Current page number. |
| `pageSize` | `number` | Page size |
| `payeeDetails` | `Record<string, any>` | The Payee object contains details of the beneficiary, person or business. |
| `payerDetails` | `Record<string, any>` |  |
| `paymentRequestCustomFields` | `Record<string, any>` | The custom fields that were attached to the payment request that resulted in this transaction. |
| `paymentRequestID` | `string` | For Pay by Bank and Direct Debit transactions this will contain the ID of the payment request. |
| `payoutID` | `string` | ID of the payout that resulted in the transaction. |
| `proprietaryBankTransactionCode` | `Record<string, any>` |  |
| `rawReference` | `string` | The raw payment reference details as received from the payment processor. |
| `reference` | `string` |  |
| `ruleID` | `string` | ID of the rule that resulted in the transaction. |
| `statementReferences` | `any[]` |  |
| `status` | `string` |  |
| `supplementaryData` | `any` |  |
| `tags` | `any[]` | An optional list of descriptive tags attached to the transaction. |
| `theirReference` | `string` | For a pay out the reference that the payer attached for the receiving party. |
| `totalPages` | `number` | Total pages |
| `totalSize` | `number` | Total count |
| `transactionAmount` | `Record<string, any>` |  |
| `transactionDate` | `string` | Date when the transaction occurred. |
| `transactionInformation` | `any[]` |  |
| `transactionMutability` | `string` |  |
| `type` | `string` | Type of the transaction. |
| `valueDateTime` | `string` |  |
| `virtualIBAN` | `string` | If set it indicates the payin was to a virtual IBAN. |
| `yourReference` | `string` | For a pay in the reference the sending party attached. |

#### Example: Load

```ts
const transaction = await client.Transaction().load({ id: 'transaction_id' })
```

#### Example: List

```ts
const transactions = await client.Transaction().list({ account_id: "example", id: "example_id" })
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
| `clientSessionTimeouts` | `any[]` | The number of seconds a session for this user should last before expiring. |
| `emailAddress` | `string` |  |
| `firstName` | `string` |  |
| `id` | `string` |  |
| `lastName` | `string` |  |
| `passkeyAdded` | `boolean` |  |
| `permissions` | `Record<string, any>` |  |
| `profile` | `string` |  |
| `rolesWithScope` | `any[]` |  |
| `twoFactorEnabled` | `boolean` |  |
| `userInviteID` | `string` | Optional ID of the invite that is being accepted so the user can be assigned a role on a new merchant. |

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
  id: 'example_id',
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
| `accountName` | `string` | Name for the account |
| `accountSupplierName` | `string` | The payment account supplier name. |
| `availableBalance` | `number` | The current available balance of the account. |
| `availableBalanceMinorUnits` | `number` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `number` | Balance of the account. |
| `balanceMinorUnits` | `number` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `string` | The bank name for external accounts |
| `consentID` | `string` | The ID of the consent used to connect the external account. |
| `createdBy` | `Record<string, any>` |  |
| `createdByDisplayName` | `string` | Either the name of the user, merchant token or api key that created the account |
| `currency` | `string` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `string` | Indicates the default payment rail for this account. |
| `displayName` | `string` | Gets a unique display name for the payment account. |
| `expiryDate` | `string` | The date that the external account will expire |
| `externalAccountIcon` | `string` | The Icon for external accounts |
| `id` | `string` | Unique id for the account. |
| `identifier` | `Record<string, any>` |  |
| `inserted` | `string` | Timestamp when the account was created. |
| `isArchived` | `boolean` | Indicates whether the account is archived. |
| `isConnectedAccount` | `boolean` | Indicates if the payment account is an externally connected account. |
| `isDefault` | `boolean` | Is the default account |
| `isTrustAccount` | `boolean` | Indicates if the payment account is a trust account. |
| `isVirtual` | `boolean` | True if the account is a virtual account. |
| `lastTransaction` | `Record<string, any>` |  |
| `lastUpdated` | `string` | Timestamp when the account was last updated. |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `merchantName` | `string` | The name of the merchant that owns the account. |
| `name` | `string` | The name of the virtual account. |
| `physicalAccountID` | `string` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `any[]` | The list of rules associated with this account. |
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
| `destinationUrl` | `string` | The destination URL for the webhook. |
| `emailAddress` | `string` | The recipient email address(es) for notifications. |
| `failedNotificationEmailAddress` | `string` | The email address to which notifications about failed webhook deliveries will be sent. |
| `id` | `string` |  |
| `isActive` | `boolean` |  |
| `merchantID` | `string` | The ID of the merchant that the webhook is for. |
| `notificationMethod` | `string` | The type of notification that will be sent. |
| `resourceTypes` | `any[]` | The resource types that the webhook will be generated for. |
| `retry` | `boolean` |  |
| `secret` | `string` | The secret key required to authenticate webhook notifications. |
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
