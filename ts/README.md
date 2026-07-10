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

`list()` resolves to an array of Account objects — iterate it directly:

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
// Create — returns the created Account
const created = await client.Account().create({
  account_id: 'example_account_id',
  currency: 'example_currency',
})

// Update — the id comes straight off the returned entity
const updated = await client.Account().update({
  id: created.id!,
  account_id: 'example_account_id',
  amount: 1,
})

// Remove
await client.Account().remove({
  id: created.id!,
})
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const accounts = await client.Account().list()
  console.log(accounts)
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

const account = await client.Account().list()
// account is a bare entity populated with mock response data
console.log(account)
```

You can also use the instance method:

```ts
const client = new NofrixionSDK({ apikey: '...' })
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.Account()

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
| `MerchantDirectDebitMandate(data?)` | `MerchantDirectDebitMandateEntity` | Create a MerchantDirectDebitMandate entity instance. |
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
| `PayoutKeyset(data?)` | `PayoutKeysetEntity` | Create a PayoutKeyset entity instance. |
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
| `account_balance` |  |
| `account_id` |  |
| `account_identification` |  |
| `account_name` |  |
| `account_supplier_name` |  |
| `account_type` |  |
| `available_balance` |  |
| `available_balance_minor_unit` |  |
| `balance` |  |
| `balance_minor_unit` |  |
| `bank_name` |  |
| `consent_id` |  |
| `consolidated_account_information` |  |
| `created_by` |  |
| `created_by_display_name` |  |
| `currency` |  |
| `default_payment_rail` |  |
| `description` |  |
| `detail` |  |
| `display_name` |  |
| `expiry_date` |  |
| `external_account_icon` |  |
| `format` |  |
| `from_date` |  |
| `id` |  |
| `identifier` |  |
| `inserted` |  |
| `is_archived` |  |
| `is_connected_account` |  |
| `is_default` |  |
| `is_trust_account` |  |
| `is_virtual` |  |
| `last_transaction` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `merchant_name` |  |
| `nickname` |  |
| `physical_account_id` |  |
| `role_i_d` |  |
| `rule` |  |
| `submitted_payouts_balance` |  |
| `submitted_payouts_balance_minor_unit` |  |
| `summary` |  |
| `supplier_physical_account_id` |  |
| `supplier_sepa_instant_status` |  |
| `to_date` |  |
| `type` |  |
| `usage_type` |  |
| `xero_bank_feed_connection_status` |  |
| `xero_bank_feed_last_synced_at` |  |
| `xero_bank_feed_sync_last_failed_at` |  |
| `xero_bank_feed_sync_last_failure_reason` |  |
| `xero_bank_feed_sync_status` |  |
| `xero_unsynchronised_transactions_count` |  |

Operations: create, list, load, remove, update.

API path: `/api/v1/accounts/{accountID}/{currency}`

#### Batch

| Field | Description |
| --- | --- |
| `approve_url` |  |
| `id` |  |
| `payout` |  |

Operations: create, load.

API path: `/api/v1/payouts/batch`

#### Beneficiary

| Field | Description |
| --- | --- |
| `approval_callback_url` |  |
| `authentication_method` |  |
| `authorisation` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `beneficiary` |  |
| `beneficiary_event` |  |
| `can_authorise` |  |
| `can_update` |  |
| `created_by` |  |
| `created_by_email_address` |  |
| `currency` |  |
| `destination` |  |
| `failed_beneficiary` |  |
| `has_current_user_authorised` |  |
| `id` |  |
| `inserted` |  |
| `is_enabled` |  |
| `last_authorised` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `name` |  |
| `nonce` |  |
| `source_account` |  |
| `source_account_i_d` |  |
| `their_reference` |  |

Operations: create, list, load, remove, update.

API path: `/api/v1/beneficiaries/authorise/{id}`

#### BeneficiaryGroup

| Field | Description |
| --- | --- |
| `group_member` |  |
| `group_name` |  |
| `id` |  |
| `inserted` |  |
| `last_updated` |  |
| `merchant_id` |  |

Operations: list.

API path: `/api/v1/merchants/{merchantID}/beneficiarygroups`

#### Card

| Field | Description |
| --- | --- |
| `authorized_amount` |  |
| `currency_code` |  |
| `is_payer_authentication_required` |  |
| `is_soft_decline` |  |
| `payer_authentication_access_token` |  |
| `payer_authentication_merchant_data` |  |
| `payer_authentication_url` |  |
| `payer_authentication_window_height` |  |
| `payer_authentication_window_width` |  |
| `payment_request_callback_url` |  |
| `payment_request_id` |  |
| `request_id` |  |
| `response_code` |  |
| `response_type` |  |
| `status` |  |
| `three_ds_redirect_url` |  |
| `transaction_id` |  |

Operations: create.

API path: `/api/v1/paymentrequests/{id}/card`

#### CardCustomerToken

| Field | Description |
| --- | --- |
| `card_type` |  |
| `customer_email_address` |  |
| `expiry_month` |  |
| `expiry_year` |  |
| `id` |  |
| `inserted` |  |
| `last_four_digit` |  |
| `last_updated` |  |
| `masked_card_number` |  |
| `merchant_id` |  |
| `payment_request_id` |  |

Operations: list, load, remove.

API path: `/api/v1/paymentrequests/card/customertokens/{merchantID}/{customerEmailAddress}`

#### CardPayment

| Field | Description |
| --- | --- |
| `authorized_amount` |  |
| `currency_code` |  |
| `is_payer_authentication_required` |  |
| `is_soft_decline` |  |
| `payer_authentication_access_token` |  |
| `payer_authentication_merchant_data` |  |
| `payer_authentication_url` |  |
| `payer_authentication_window_height` |  |
| `payer_authentication_window_width` |  |
| `payment_request_callback_url` |  |
| `payment_request_id` |  |
| `request_id` |  |
| `response_code` |  |
| `response_type` |  |
| `status` |  |
| `three_ds_redirect_url` |  |
| `transaction_id` |  |

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
| `authorisation_url` |  |
| `callback_url` |  |
| `consent_id` |  |
| `email_address` |  |
| `expiry_date` |  |
| `failure_callback_url` |  |
| `id` |  |
| `inserted` |  |
| `institution_id` |  |
| `is_connected_account` |  |
| `is_enabled` |  |
| `merchant_id` |  |
| `provider` |  |
| `success_web_hook_url` |  |

Operations: create, list, load, remove, update.

API path: `/api/v1/openbanking/consents`

#### Currency

| Field | Description |
| --- | --- |
| `code` |  |
| `decimal` |  |
| `is_fiat` |  |
| `iso4217_alpha_code` |  |
| `iso4217_numeric_code` |  |
| `symbol` |  |

Operations: list.

API path: `/api/v1/currencies`

#### DirectDebitBatchSubmit

| Field | Description |
| --- | --- |
| `failed_submission` |  |
| `successful_submission` |  |

Operations: create.

API path: `/api/v1/paymentrequests/directdebit/batchsubmit`

#### FxRate

| Field | Description |
| --- | --- |
| `destination_currency` |  |
| `exchange_rate` |  |
| `expiry_time` |  |
| `quote_id` |  |
| `source_currency` |  |

Operations: list, load.

API path: `/api/v1/payouts/fxallheldrates/{source}/{destination}`

#### IPayment

| Field | Description |
| --- | --- |
| `payment_request_id` |  |
| `response_type` |  |

Operations: create.

API path: `/api/v1/paymentrequests/payondemand`

#### Mandate

| Field | Description |
| --- | --- |
| `account_number` |  |
| `address_line1` |  |
| `address_line2` |  |
| `approved_at` |  |
| `city` |  |
| `country_code` |  |
| `currency` |  |
| `customer_account_number` |  |
| `customer_city` |  |
| `customer_country_code` |  |
| `customer_country_name` |  |
| `customer_email_address` |  |
| `customer_first_name` |  |
| `customer_iban` |  |
| `customer_last_name` |  |
| `customer_sort_code` |  |
| `email_address` |  |
| `first_name` |  |
| `iban` |  |
| `id` |  |
| `inserted` |  |
| `is_recurring` |  |
| `last_name` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `postal_code` |  |
| `reference` |  |
| `sort_code` |  |
| `status` |  |
| `supplier_bank_account_id` |  |
| `supplier_customer_id` |  |
| `supplier_mandate_id` |  |
| `supplier_name` |  |
| `supplier_status` |  |

Operations: create, load.

API path: `/api/v1/mandates`

#### Merchant

| Field | Description |
| --- | --- |
| `account_currency` |  |
| `can_have_trust_account` |  |
| `card_payment_processor` |  |
| `company_id` |  |
| `display_qr_on_hosted_pay` |  |
| `hosted_pay_version` |  |
| `id` |  |
| `inserted` |  |
| `is_blocked` |  |
| `is_exited` |  |
| `is_suspended` |  |
| `jurisdiction` |  |
| `logo_url_png` |  |
| `logo_url_svg` |  |
| `merchant_category_code` |  |
| `name` |  |
| `note` |  |
| `parent_merchant` |  |
| `payment_account` |  |
| `payment_account_limit` |  |
| `reason` |  |
| `short_name` |  |
| `supported_payment_methods_list` |  |
| `suspension_reason` |  |
| `tag` |  |
| `time_zone_id` |  |
| `trading_name` |  |
| `web_hook_limit` |  |
| `your_role_name` |  |

Operations: list, load, remove, update.

API path: `/api/v1/merchants/{merchantID}/childmerchants`

#### MerchantAuthorisationSetting

| Field | Description |
| --- | --- |
| `amount_lower` |  |
| `amount_upper` |  |
| `authorisation_type` |  |
| `beneficiaries_only` |  |
| `id` |  |
| `inserted` |  |
| `last_editor_cant_authorise` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `number_of_authoriser` |  |
| `role_setting` |  |

Operations: list.

API path: `/api/v1/merchants/{merchantID}/authorisationsettings`

#### MerchantDirectDebitMandate

| Field | Description |
| --- | --- |
| `approved_at` |  |
| `currency` |  |
| `customer_account_number` |  |
| `customer_city` |  |
| `customer_country_code` |  |
| `customer_country_name` |  |
| `customer_email_address` |  |
| `customer_first_name` |  |
| `customer_iban` |  |
| `customer_last_name` |  |
| `customer_sort_code` |  |
| `id` |  |
| `inserted` |  |
| `is_recurring` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `reference` |  |
| `status` |  |
| `supplier_bank_account_id` |  |
| `supplier_customer_id` |  |
| `supplier_mandate_id` |  |
| `supplier_name` |  |
| `supplier_status` |  |

Operations: list.

API path: `/api/v1/mandates`

#### MerchantPayByBankSetting

| Field | Description |
| --- | --- |
| `bank_country_code` |  |
| `bank_id` |  |
| `bank_name` |  |
| `business_institution_id` |  |
| `currency` |  |
| `logo` |  |
| `message` |  |
| `message_image_url` |  |
| `order` |  |
| `personal_institution_id` |  |
| `processor` |  |
| `warning_heading` |  |
| `warning_message` |  |

Operations: list.

API path: `/api/v1/merchants/{merchantID}/banksettings`

#### MerchantPaymentRequestTemplate

| Field | Description |
| --- | --- |
| `description` |  |
| `id` |  |
| `inserted` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `name` |  |
| `template` |  |

Operations: list, load, remove, update.

API path: `/api/v1/paymentrequests/{merchantID}/templates`

#### MerchantToken

| Field | Description |
| --- | --- |
| `authentication_method` |  |
| `authorisation` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `can_authorise` |  |
| `description` |  |
| `expires_at` |  |
| `has_current_user_authorised` |  |
| `hmac_algorithm` |  |
| `id` |  |
| `inserted` |  |
| `ip_address_whitelist` |  |
| `is_archived` |  |
| `is_enabled` |  |
| `last_authorised` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `nonce` |  |
| `permission_type` |  |
| `request_signature_version` |  |
| `shared_secret_algorithm` |  |
| `shared_secret_base64` |  |
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
| `build_version` |  |
| `major_version` |  |
| `minor_version` |  |
| `release_name` |  |

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
| `account_name` |  |
| `account_number` |  |
| `iban` |  |
| `payee_verified_account_name` |  |
| `result` |  |
| `secondary_identification` |  |
| `sort_code` |  |

Operations: create.

API path: `/api/v1/openbanking/payeeverification`

#### Payment

| Field | Description |
| --- | --- |
| `address` |  |
| `amount` |  |
| `amount_pending` |  |
| `amount_received` |  |
| `amount_refunded` |  |
| `auto_send_receipt` |  |
| `base_origin_url` |  |
| `callback_url` |  |
| `card_authorize_only` |  |
| `card_create_token` |  |
| `card_create_token_mode` |  |
| `card_ignore_cvn` |  |
| `card_no_payer_authentication` |  |
| `card_processor_merchant_id` |  |
| `card_stripe_payment_intent_id` |  |
| `card_stripe_payment_intent_secret` |  |
| `card_transmit_raw_detail` |  |
| `created_by_user` |  |
| `currency` |  |
| `custom_field` |  |
| `customer_email_address` |  |
| `customer_id` |  |
| `customer_name` |  |
| `description` |  |
| `destination_account` |  |
| `direct_debit_payment` |  |
| `due_date` |  |
| `event` |  |
| `failure_callback_url` |  |
| `field_display_setting` |  |
| `formatted_amount` |  |
| `hosted_pay_checkout_url` |  |
| `id` |  |
| `ignore_address_verification` |  |
| `inserted` |  |
| `inserted_sortable` |  |
| `is_archived` |  |
| `jwk` |  |
| `last_updated` |  |
| `lightning_invoice` |  |
| `lightning_invoice_expires_at` |  |
| `merchant_direct_debit_mandate_id` |  |
| `merchant_id` |  |
| `merchant_token_description` |  |
| `notification_email_address` |  |
| `notification_role_i_d` |  |
| `order_id` |  |
| `partial_payment_method` |  |
| `partial_payment_step` |  |
| `payment_attempt` |  |
| `payment_method` |  |
| `payment_processor` |  |
| `payrun_id` |  |
| `pisp_account_id` |  |
| `priority_bank_id` |  |
| `result` |  |
| `sandbox_settle_delay_in_second` |  |
| `shipping_address` |  |
| `shipping_address_city` |  |
| `shipping_address_country_code` |  |
| `shipping_address_county` |  |
| `shipping_address_line1` |  |
| `shipping_address_line2` |  |
| `shipping_address_post_code` |  |
| `shipping_email` |  |
| `shipping_first_name` |  |
| `shipping_last_name` |  |
| `shipping_phone` |  |
| `status` |  |
| `success_web_hook_url` |  |
| `tag` |  |
| `tag_id` |  |
| `title` |  |
| `tokenised_card` |  |
| `transaction` |  |
| `use_hosted_payment_page` |  |

Operations: create, load, update.

API path: `/api/v1/paymentrequests`

#### PaymentAccount

| Field | Description |
| --- | --- |
| `account_name` |  |
| `account_supplier_name` |  |
| `available_balance` |  |
| `available_balance_minor_unit` |  |
| `balance` |  |
| `balance_minor_unit` |  |
| `bank_name` |  |
| `consent_id` |  |
| `created_by` |  |
| `created_by_display_name` |  |
| `currency` |  |
| `default_payment_rail` |  |
| `display_name` |  |
| `expiry_date` |  |
| `external_account_icon` |  |
| `id` |  |
| `identifier` |  |
| `inserted` |  |
| `is_archived` |  |
| `is_connected_account` |  |
| `is_default` |  |
| `is_trust_account` |  |
| `is_virtual` |  |
| `last_transaction` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `merchant_name` |  |
| `physical_account_id` |  |
| `rule` |  |
| `submitted_payouts_balance` |  |
| `submitted_payouts_balance_minor_unit` |  |
| `summary` |  |
| `supplier_sepa_instant_status` |  |
| `xero_bank_feed_connection_status` |  |
| `xero_bank_feed_last_synced_at` |  |
| `xero_bank_feed_sync_last_failed_at` |  |
| `xero_bank_feed_sync_last_failure_reason` |  |
| `xero_bank_feed_sync_status` |  |
| `xero_unsynchronised_transactions_count` |  |

Operations: list.

API path: `/api/v1/accounts/paged`

#### PaymentAccountMinimal

| Field | Description |
| --- | --- |
| `account_name` |  |
| `available_balance` |  |
| `balance` |  |
| `balance_minor_unit` |  |
| `currency` |  |
| `id` |  |
| `identifier` |  |
| `is_archived` |  |
| `is_connected_account` |  |
| `merchant_id` |  |
| `submitted_payouts_balance` |  |

Operations: list.

API path: `/api/v1/accounts/minimal`

#### PaymentInitiation

| Field | Description |
| --- | --- |
| `payment_initiation_id` |  |
| `payment_request_callback_url` |  |
| `payment_request_id` |  |
| `redirect_url` |  |
| `response_type` |  |
| `specific_error_message` |  |

Operations: create.

API path: `/api/v1/paymentrequests/{id}/pisp`

#### PaymentRequest

| Field | Description |
| --- | --- |
| `address` |  |
| `amount` |  |
| `amount_pending` |  |
| `amount_received` |  |
| `amount_refunded` |  |
| `auto_send_receipt` |  |
| `base_origin_url` |  |
| `callback_url` |  |
| `card_authorize_only` |  |
| `card_create_token` |  |
| `card_create_token_mode` |  |
| `card_ignore_cvn` |  |
| `card_processor_merchant_id` |  |
| `card_stripe_payment_intent_id` |  |
| `card_stripe_payment_intent_secret` |  |
| `created_by_user` |  |
| `currency` |  |
| `custom_field` |  |
| `customer_email_address` |  |
| `customer_id` |  |
| `customer_name` |  |
| `description` |  |
| `destination_account` |  |
| `direct_debit_payment` |  |
| `do_simulate_settlement_failure` |  |
| `due_date` |  |
| `error_description` |  |
| `event` |  |
| `failed_payment_request` |  |
| `failure_callback_url` |  |
| `field_display_setting` |  |
| `formatted_amount` |  |
| `hosted_pay_checkout_url` |  |
| `id` |  |
| `ignore_address_verification` |  |
| `inserted` |  |
| `inserted_sortable` |  |
| `institution` |  |
| `is_archived` |  |
| `jwk` |  |
| `last_updated` |  |
| `lightning_invoice` |  |
| `lightning_invoice_expires_at` |  |
| `merchant_direct_debit_mandate_id` |  |
| `merchant_id` |  |
| `merchant_token_description` |  |
| `notification_email_address` |  |
| `notification_role_i_d` |  |
| `order_id` |  |
| `partial_payment_method` |  |
| `partial_payment_step` |  |
| `payment_attempt` |  |
| `payment_initiation_id` |  |
| `payment_method` |  |
| `payment_processor` |  |
| `payment_request` |  |
| `payrun_id` |  |
| `pisp_account_id` |  |
| `priority_bank_id` |  |
| `result` |  |
| `sandbox_settle_delay_in_second` |  |
| `shipping_address` |  |
| `status` |  |
| `success_web_hook_url` |  |
| `tag` |  |
| `title` |  |
| `tokenised_card` |  |
| `transaction` |  |
| `use_hosted_payment_page` |  |

Operations: create, list, load, remove, update.

API path: `/api/v1/paymentrequests/{id}/directdebit`

#### PaymentRequestEvent

| Field | Description |
| --- | --- |
| `amount` |  |
| `apple_pay_transaction_id` |  |
| `card_authorization_response_id` |  |
| `card_expiry_month` |  |
| `card_expiry_year` |  |
| `card_issuer` |  |
| `card_issuer_country` |  |
| `card_last_four_digit` |  |
| `card_request_id` |  |
| `card_scheme` |  |
| `card_token_customer_id` |  |
| `card_transaction_id` |  |
| `currency` |  |
| `direct_debit_payment_id` |  |
| `direct_debit_payment_reference` |  |
| `drirect_debit_mandate_id` |  |
| `error_message` |  |
| `error_reason` |  |
| `event_type` |  |
| `id` |  |
| `inserted` |  |
| `lightning_invoice` |  |
| `lightning_r_hash` |  |
| `origin_url` |  |
| `payment_method_type` |  |
| `payment_processor_name` |  |
| `payment_request_id` |  |
| `pisp_bank_status` |  |
| `pisp_payment_initiation_id` |  |
| `pisp_payment_institution_name` |  |
| `pisp_payment_service_provider_id` |  |
| `pisp_redirect_url` |  |
| `reconciled_transaction_id` |  |
| `refund_payout_id` |  |
| `status` |  |
| `wallet_name` |  |

Operations: list.

API path: `/api/v1/paymentrequests/{id}/events`

#### PaymentRequestMetric

| Field | Description |
| --- | --- |
| `all` |  |
| `authorized` |  |
| `paid` |  |
| `partially_paid` |  |
| `total_amounts_by_currency` |  |
| `unpaid` |  |

Operations: load.

API path: `/api/v1/paymentrequests/metrics`

#### PaymentRequestMinimal

| Field | Description |
| --- | --- |
| `amount` |  |
| `amount_pending` |  |
| `amount_received` |  |
| `amount_refunded` |  |
| `callback_url` |  |
| `card_stripe_payment_intent_secret` |  |
| `country_code` |  |
| `currency` |  |
| `custom_fields_to_display` |  |
| `description` |  |
| `due_date` |  |
| `field_display_setting` |  |
| `google_pay_merchant_id` |  |
| `id` |  |
| `jwk` |  |
| `merchant_id` |  |
| `merchant_logo_url_png` |  |
| `merchant_logo_url_svg` |  |
| `merchant_name` |  |
| `merchant_short_name` |  |
| `partial_payment_method` |  |
| `payment_attempt` |  |
| `payment_methods_list` |  |
| `payment_processor` |  |
| `payment_processor_key` |  |
| `pisp_error` |  |
| `priority_bank_id` |  |
| `status` |  |
| `stripe_account_id` |  |
| `title` |  |

Operations: list.

API path: `/api/v1/paymentrequests/{id}/minimal`

#### PaymentRequestResult

| Field | Description |
| --- | --- |
| `amount` |  |
| `amount_pending` |  |
| `amount_received` |  |
| `amount_refunded` |  |
| `currency` |  |
| `customer_id` |  |
| `payment` |  |
| `payment_request_id` |  |
| `pisp_authorization` |  |
| `requested_amount` |  |
| `result` |  |

Operations: list.

API path: `/api/v1/paymentrequests/{id}/result`

#### Payout

| Field | Description |
| --- | --- |
| `account_id` |  |
| `allow_incomplete` |  |
| `amount` |  |
| `amount_minor_unit` |  |
| `approve_payout_url` |  |
| `approver_id` |  |
| `authentication_method` |  |
| `authorisation` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `batch_payout_id` |  |
| `beneficiary` |  |
| `beneficiary_id` |  |
| `can_authorise` |  |
| `can_process` |  |
| `can_update` |  |
| `charge_bearer` |  |
| `created_by` |  |
| `created_by_email_address` |  |
| `currency` |  |
| `current_user_id` |  |
| `description` |  |
| `destination` |  |
| `document` |  |
| `event` |  |
| `failed_payout` |  |
| `formatted_amount` |  |
| `formatted_fx_destination_amount` |  |
| `formatted_schedule` |  |
| `formatted_schedule_day_only` |  |
| `formatted_source_account_available_balance` |  |
| `fx_destination_amount` |  |
| `fx_destination_amount_minor_unit` |  |
| `fx_destination_currency` |  |
| `fx_quote_expires_at` |  |
| `fx_quote_id` |  |
| `fx_rate` |  |
| `fx_use_destination_amount` |  |
| `has_current_user_authorised` |  |
| `id` |  |
| `inserted` |  |
| `invoice_id` |  |
| `is_archived` |  |
| `is_failed` |  |
| `is_settled` |  |
| `is_submitted` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `merchant_token_description` |  |
| `nonce` |  |
| `payment_processor` |  |
| `payment_rail` |  |
| `payout` |  |
| `payrun_id` |  |
| `payrun_name` |  |
| `reason` |  |
| `rule` |  |
| `schedule_date` |  |
| `scheduled` |  |
| `source_account_available_balance` |  |
| `source_account_available_balance_minor_unit` |  |
| `source_account_bic` |  |
| `source_account_currency` |  |
| `source_account_iban` |  |
| `source_account_identifier` |  |
| `source_account_name` |  |
| `source_account_number` |  |
| `source_account_sortcode` |  |
| `status` |  |
| `tag` |  |
| `tag_id` |  |
| `their_reference` |  |
| `topup_payrun_id` |  |
| `transacted_amount` |  |
| `transacted_fx_amount` |  |
| `transacted_fx_rate` |  |
| `type` |  |
| `user_id` |  |
| `your_reference` |  |

Operations: create, list, load, remove, update.

API path: `/api/v1/payouts/batch/submit/{id}`

#### PayoutKeyset

| Field | Description |
| --- | --- |
| `account_id` |  |
| `amount` |  |
| `amount_minor_unit` |  |
| `approve_payout_url` |  |
| `approver_id` |  |
| `authentication_method` |  |
| `authorisation` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `batch_payout_id` |  |
| `beneficiary` |  |
| `can_authorise` |  |
| `can_process` |  |
| `can_update` |  |
| `charge_bearer` |  |
| `created_by` |  |
| `created_by_email_address` |  |
| `currency` |  |
| `current_user_id` |  |
| `description` |  |
| `destination` |  |
| `document` |  |
| `event` |  |
| `formatted_amount` |  |
| `formatted_fx_destination_amount` |  |
| `formatted_schedule` |  |
| `formatted_schedule_day_only` |  |
| `formatted_source_account_available_balance` |  |
| `fx_destination_amount` |  |
| `fx_destination_amount_minor_unit` |  |
| `fx_destination_currency` |  |
| `fx_quote_expires_at` |  |
| `fx_quote_id` |  |
| `fx_rate` |  |
| `fx_use_destination_amount` |  |
| `has_current_user_authorised` |  |
| `id` |  |
| `inserted` |  |
| `invoice_id` |  |
| `is_archived` |  |
| `is_failed` |  |
| `is_settled` |  |
| `is_submitted` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `merchant_token_description` |  |
| `nonce` |  |
| `payment_processor` |  |
| `payment_rail` |  |
| `payrun_id` |  |
| `payrun_name` |  |
| `rule` |  |
| `schedule_date` |  |
| `scheduled` |  |
| `source_account_available_balance` |  |
| `source_account_available_balance_minor_unit` |  |
| `source_account_bic` |  |
| `source_account_currency` |  |
| `source_account_iban` |  |
| `source_account_identifier` |  |
| `source_account_name` |  |
| `source_account_number` |  |
| `source_account_sortcode` |  |
| `status` |  |
| `tag` |  |
| `their_reference` |  |
| `topup_payrun_id` |  |
| `transacted_amount` |  |
| `transacted_fx_amount` |  |
| `transacted_fx_rate` |  |
| `type` |  |
| `user_id` |  |
| `your_reference` |  |

Operations: list.

API path: `/api/v1/accounts/{accountID}/payouts/failed`

#### PayoutMetric

| Field | Description |
| --- | --- |
| `all` |  |
| `failed` |  |
| `in_progress` |  |
| `paid` |  |
| `pending_approval` |  |
| `scheduled` |  |
| `total_amounts_by_currency` |  |

Operations: load.

API path: `/api/v1/payouts/metrics`

#### Payrun

| Field | Description |
| --- | --- |
| `authorisation` |  |
| `authorisation_date` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `batch_payout_id` |  |
| `can_authorise` |  |
| `can_delete` |  |
| `can_edit` |  |
| `event` |  |
| `has_current_user_authorised` |  |
| `id` |  |
| `inserted` |  |
| `invoice` |  |
| `invoices_minimal` |  |
| `is_archived` |  |
| `last_updated` |  |
| `last_updated_by` |  |
| `merchant_id` |  |
| `name` |  |
| `nonce` |  |
| `note` |  |
| `payment` |  |
| `payout` |  |
| `payouts_count` |  |
| `reason` |  |
| `schedule_date` |  |
| `scheduled_date` |  |
| `source_account` |  |
| `status` |  |
| `total_eur` |  |
| `total_gbp` |  |
| `total_usd` |  |

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
| `content` |  |
| `content_type` |  |
| `last_completed_at` |  |
| `merchant_id` |  |
| `report_name` |  |
| `report_type` |  |
| `statement_number` |  |

Operations: load.

API path: `/api/v1/reports/{id}/result/{statementNumber}`

#### Role

| Field | Description |
| --- | --- |
| `failed_role` |  |
| `role` |  |

Operations: create.

API path: `/api/v1/merchants/{merchantID}/roles/batchcreate`

#### Rule

| Field | Description |
| --- | --- |
| `account` |  |
| `account_id` |  |
| `approve_url` |  |
| `approver_id` |  |
| `authentication_method` |  |
| `authorisation` |  |
| `authorisers_completed_count` |  |
| `authorisers_required_count` |  |
| `can_authorise` |  |
| `created_by` |  |
| `description` |  |
| `end_at` |  |
| `has_current_user_authorised` |  |
| `id` |  |
| `inserted` |  |
| `is_disabled` |  |
| `last_executed_at` |  |
| `last_run_at_transaction_date` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `name` |  |
| `nonce` |  |
| `on_approved_web_hook_url` |  |
| `on_execution_error_web_hook_url` |  |
| `on_execution_success_web_hook_url` |  |
| `start_at` |  |
| `status` |  |
| `sweep_action` |  |
| `time_zone_id` |  |
| `trigger_cron_expression` |  |
| `trigger_on_pay_in` |  |
| `user_id` |  |
| `web_hook_secret` |  |

Operations: create, list, load, remove, update.

API path: `/api/v1/rules`

#### RuleEvent

| Field | Description |
| --- | --- |
| `error_message` |  |
| `id` |  |
| `inserted` |  |
| `is_authorise_to_enable` |  |
| `message` |  |
| `raw_response` |  |
| `rule_event_type` |  |
| `rule_id` |  |
| `user` |  |

Operations: list.

API path: `/api/v1/rules/{id}/events`

#### Tag

| Field | Description |
| --- | --- |
| `colour_hex` |  |
| `description` |  |
| `id` |  |
| `merchant_id` |  |
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
| `account_id` |  |
| `account_name` |  |
| `account_sequence_number` |  |
| `address_detail` |  |
| `amount` |  |
| `amount_minor_unit` |  |
| `balance` |  |
| `balance_minor_unit` |  |
| `booking_date_time` |  |
| `charge_detail` |  |
| `content` |  |
| `counterparty` |  |
| `counterparty_summary` |  |
| `currency` |  |
| `currency_exchange` |  |
| `date` |  |
| `description` |  |
| `enrichment` |  |
| `fx_amount` |  |
| `fx_currency` |  |
| `fx_rate` |  |
| `gross_amount` |  |
| `id` |  |
| `inserted` |  |
| `iso_bank_transaction_code` |  |
| `merchant` |  |
| `merchant_id` |  |
| `page_number` |  |
| `page_size` |  |
| `payee_detail` |  |
| `payer_detail` |  |
| `payment_request_custom_field` |  |
| `payment_request_id` |  |
| `payout_id` |  |
| `proprietary_bank_transaction_code` |  |
| `raw_reference` |  |
| `reference` |  |
| `rule_id` |  |
| `statement_reference` |  |
| `status` |  |
| `supplementary_data` |  |
| `tag` |  |
| `their_reference` |  |
| `total_page` |  |
| `total_size` |  |
| `transaction_amount` |  |
| `transaction_date` |  |
| `transaction_information` |  |
| `transaction_mutability` |  |
| `type` |  |
| `value_date_time` |  |
| `virtual_iban` |  |
| `your_reference` |  |

Operations: create, list, load, remove.

API path: `/api/v1/transactions/{id}/tags`

#### User

| Field | Description |
| --- | --- |
| `client_session_timeout` |  |
| `email_address` |  |
| `first_name` |  |
| `id` |  |
| `last_name` |  |
| `passkey_added` |  |
| `permission` |  |
| `profile` |  |
| `roles_with_scope` |  |
| `two_factor_enabled` |  |
| `user_invite_id` |  |

Operations: list, update.

API path: `/api/v1/user/{merchantID}/userspaged`

#### UserInvite

| Field | Description |
| --- | --- |
| `authorisation_status` |  |
| `failed_user_invite` |  |
| `id` |  |
| `initial_role_id` |  |
| `invitee_email_address` |  |
| `invitee_first_name` |  |
| `invitee_last_name` |  |
| `inviter_email_address` |  |
| `inviter_first_name` |  |
| `inviter_last_name` |  |
| `is_authorised` |  |
| `is_invitee_registered` |  |
| `last_invited` |  |
| `merchant_id` |  |
| `merchant_name` |  |
| `message` |  |
| `registration_url` |  |
| `send_invite_email` |  |
| `status` |  |
| `user` |  |
| `user_id` |  |
| `user_invite` |  |

Operations: create, list, load, remove, update.

API path: `/api/v1/userinvites/authorise/{id}`

#### Virtual

| Field | Description |
| --- | --- |
| `account_name` |  |
| `account_supplier_name` |  |
| `available_balance` |  |
| `available_balance_minor_unit` |  |
| `balance` |  |
| `balance_minor_unit` |  |
| `bank_name` |  |
| `consent_id` |  |
| `created_by` |  |
| `created_by_display_name` |  |
| `currency` |  |
| `default_payment_rail` |  |
| `display_name` |  |
| `expiry_date` |  |
| `external_account_icon` |  |
| `id` |  |
| `identifier` |  |
| `inserted` |  |
| `is_archived` |  |
| `is_connected_account` |  |
| `is_default` |  |
| `is_trust_account` |  |
| `is_virtual` |  |
| `last_transaction` |  |
| `last_updated` |  |
| `merchant_id` |  |
| `merchant_name` |  |
| `name` |  |
| `physical_account_id` |  |
| `rule` |  |
| `submitted_payouts_balance` |  |
| `submitted_payouts_balance_minor_unit` |  |
| `summary` |  |
| `supplier_sepa_instant_status` |  |
| `xero_bank_feed_connection_status` |  |
| `xero_bank_feed_last_synced_at` |  |
| `xero_bank_feed_sync_last_failed_at` |  |
| `xero_bank_feed_sync_last_failure_reason` |  |
| `xero_bank_feed_sync_status` |  |
| `xero_unsynchronised_transactions_count` |  |

Operations: create, update.

API path: `/api/v1/accounts/{accountID}/virtual`

#### Webhook

| Field | Description |
| --- | --- |
| `destination_url` |  |
| `email_address` |  |
| `failed_notification_email_address` |  |
| `id` |  |
| `is_active` |  |
| `merchant_id` |  |
| `notification_method` |  |
| `resource_type` |  |
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
| `account_balance` | `any[]` |  |
| `account_id` | `string` |  |
| `account_identification` | `any[]` |  |
| `account_name` | `string` |  |
| `account_supplier_name` | `string` |  |
| `account_type` | `string` |  |
| `available_balance` | `number` |  |
| `available_balance_minor_unit` | `number` |  |
| `balance` | `number` |  |
| `balance_minor_unit` | `number` |  |
| `bank_name` | `string` |  |
| `consent_id` | `string` |  |
| `consolidated_account_information` | `Record<string, any>` |  |
| `created_by` | `Record<string, any>` |  |
| `created_by_display_name` | `string` |  |
| `currency` | `string` |  |
| `default_payment_rail` | `string` |  |
| `description` | `string` |  |
| `detail` | `string` |  |
| `display_name` | `string` |  |
| `expiry_date` | `string` |  |
| `external_account_icon` | `string` |  |
| `format` | `string` |  |
| `from_date` | `string` |  |
| `id` | `string` |  |
| `identifier` | `Record<string, any>` |  |
| `inserted` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_connected_account` | `boolean` |  |
| `is_default` | `boolean` |  |
| `is_trust_account` | `boolean` |  |
| `is_virtual` | `boolean` |  |
| `last_transaction` | `Record<string, any>` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `nickname` | `string` |  |
| `physical_account_id` | `string` |  |
| `role_i_d` | `any[]` |  |
| `rule` | `any[]` |  |
| `submitted_payouts_balance` | `number` |  |
| `submitted_payouts_balance_minor_unit` | `number` |  |
| `summary` | `string` |  |
| `supplier_physical_account_id` | `string` |  |
| `supplier_sepa_instant_status` | `string` |  |
| `to_date` | `string` |  |
| `type` | `string` |  |
| `usage_type` | `string` |  |
| `xero_bank_feed_connection_status` | `string` |  |
| `xero_bank_feed_last_synced_at` | `string` |  |
| `xero_bank_feed_sync_last_failed_at` | `string` |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` |  |
| `xero_bank_feed_sync_status` | `string` |  |
| `xero_unsynchronised_transactions_count` | `number` |  |

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
| `approve_url` | `string` |  |
| `id` | `string` |  |
| `payout` | `any[]` |  |

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
| `approval_callback_url` | `string` |  |
| `authentication_method` | `any[]` |  |
| `authorisation` | `any[]` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `beneficiary` | `any[]` |  |
| `beneficiary_event` | `any[]` |  |
| `can_authorise` | `boolean` |  |
| `can_update` | `boolean` |  |
| `created_by` | `Record<string, any>` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `destination` | `Record<string, any>` |  |
| `failed_beneficiary` | `Record<string, any>` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_enabled` | `boolean` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `source_account` | `any[]` |  |
| `source_account_i_d` | `any[]` |  |
| `their_reference` | `string` |  |

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
| `group_member` | `any[]` |  |
| `group_name` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |

#### Example: List

```ts
const beneficiary_groups = await client.BeneficiaryGroup().list()
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
| `authorized_amount` | `string` |  |
| `currency_code` | `string` |  |
| `is_payer_authentication_required` | `boolean` |  |
| `is_soft_decline` | `boolean` |  |
| `payer_authentication_access_token` | `string` |  |
| `payer_authentication_merchant_data` | `string` |  |
| `payer_authentication_url` | `string` |  |
| `payer_authentication_window_height` | `number` |  |
| `payer_authentication_window_width` | `number` |  |
| `payment_request_callback_url` | `string` |  |
| `payment_request_id` | `string` |  |
| `request_id` | `string` |  |
| `response_code` | `string` |  |
| `response_type` | `string` |  |
| `status` | `string` |  |
| `three_ds_redirect_url` | `string` |  |
| `transaction_id` | `string` |  |

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
| `card_type` | `string` |  |
| `customer_email_address` | `string` |  |
| `expiry_month` | `string` |  |
| `expiry_year` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_four_digit` | `string` |  |
| `last_updated` | `string` |  |
| `masked_card_number` | `string` |  |
| `merchant_id` | `string` |  |
| `payment_request_id` | `string` |  |

#### Example: Load

```ts
const card_customer_token = await client.CardCustomerToken().load({ customer_email_address: 'customer_email_address' })
```

#### Example: List

```ts
const card_customer_tokens = await client.CardCustomerToken().list()
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
| `authorized_amount` | `string` |  |
| `currency_code` | `string` |  |
| `is_payer_authentication_required` | `boolean` |  |
| `is_soft_decline` | `boolean` |  |
| `payer_authentication_access_token` | `string` |  |
| `payer_authentication_merchant_data` | `string` |  |
| `payer_authentication_url` | `string` |  |
| `payer_authentication_window_height` | `number` |  |
| `payer_authentication_window_width` | `number` |  |
| `payment_request_callback_url` | `string` |  |
| `payment_request_id` | `string` |  |
| `request_id` | `string` |  |
| `response_code` | `string` |  |
| `response_type` | `string` |  |
| `status` | `string` |  |
| `three_ds_redirect_url` | `string` |  |
| `transaction_id` | `string` |  |

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
| `authorisation_url` | `string` |  |
| `callback_url` | `string` |  |
| `consent_id` | `string` |  |
| `email_address` | `string` |  |
| `expiry_date` | `string` |  |
| `failure_callback_url` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `institution_id` | `string` |  |
| `is_connected_account` | `boolean` |  |
| `is_enabled` | `boolean` |  |
| `merchant_id` | `string` |  |
| `provider` | `string` |  |
| `success_web_hook_url` | `string` |  |

#### Example: Load

```ts
const consent = await client.Consent().load({ id: 'consent_id' })
```

#### Example: List

```ts
const consents = await client.Consent().list()
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
| `decimal` | `number` |  |
| `is_fiat` | `boolean` |  |
| `iso4217_alpha_code` | `string` |  |
| `iso4217_numeric_code` | `string` |  |
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
| `failed_submission` | `Record<string, any>` |  |
| `successful_submission` | `any[]` |  |

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
| `destination_currency` | `string` |  |
| `exchange_rate` | `number` |  |
| `expiry_time` | `string` |  |
| `quote_id` | `string` |  |
| `source_currency` | `string` |  |

#### Example: Load

```ts
const fx_rate = await client.FxRate().load({ destination: 'destination', source: 'source', valid_for_minute: 1 })
```

#### Example: List

```ts
const fx_rates = await client.FxRate().list()
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
| `payment_request_id` | `string` |  |
| `response_type` | `string` |  |

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
| `account_number` | `string` |  |
| `address_line1` | `string` |  |
| `address_line2` | `string` |  |
| `approved_at` | `string` |  |
| `city` | `string` |  |
| `country_code` | `string` |  |
| `currency` | `string` |  |
| `customer_account_number` | `string` |  |
| `customer_city` | `string` |  |
| `customer_country_code` | `string` |  |
| `customer_country_name` | `string` |  |
| `customer_email_address` | `string` |  |
| `customer_first_name` | `string` |  |
| `customer_iban` | `string` |  |
| `customer_last_name` | `string` |  |
| `customer_sort_code` | `string` |  |
| `email_address` | `string` |  |
| `first_name` | `string` |  |
| `iban` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_recurring` | `boolean` |  |
| `last_name` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `postal_code` | `string` |  |
| `reference` | `string` |  |
| `sort_code` | `string` |  |
| `status` | `string` |  |
| `supplier_bank_account_id` | `string` |  |
| `supplier_customer_id` | `string` |  |
| `supplier_mandate_id` | `string` |  |
| `supplier_name` | `string` |  |
| `supplier_status` | `string` |  |

#### Example: Load

```ts
const mandate = await client.Mandate().load({ id: 'mandate_id' })
```

#### Example: Create

```ts
const mandate = await client.Mandate().create({
  address_line1: 'example_address_line1',
  city: 'example_city',
  country_code: 'example_country_code',
  email_address: 'example_email_address',
  first_name: 'example_first_name',
  last_name: 'example_last_name',
  postal_code: 'example_postal_code',
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
| `account_currency` | `any[]` |  |
| `can_have_trust_account` | `boolean` |  |
| `card_payment_processor` | `string` |  |
| `company_id` | `string` |  |
| `display_qr_on_hosted_pay` | `boolean` |  |
| `hosted_pay_version` | `number` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_blocked` | `boolean` |  |
| `is_exited` | `boolean` |  |
| `is_suspended` | `boolean` |  |
| `jurisdiction` | `string` |  |
| `logo_url_png` | `string` |  |
| `logo_url_svg` | `string` |  |
| `merchant_category_code` | `string` |  |
| `name` | `string` |  |
| `note` | `string` |  |
| `parent_merchant` | `Record<string, any>` |  |
| `payment_account` | `any[]` |  |
| `payment_account_limit` | `number` |  |
| `reason` | `string` |  |
| `short_name` | `string` |  |
| `supported_payment_methods_list` | `any[]` |  |
| `suspension_reason` | `string` |  |
| `tag` | `any[]` |  |
| `time_zone_id` | `string` |  |
| `trading_name` | `string` |  |
| `web_hook_limit` | `number` |  |
| `your_role_name` | `string` |  |

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
| `amount_lower` | `number` |  |
| `amount_upper` | `number` |  |
| `authorisation_type` | `string` |  |
| `beneficiaries_only` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_editor_cant_authorise` | `boolean` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `number_of_authoriser` | `number` |  |
| `role_setting` | `any[]` |  |

#### Example: List

```ts
const merchant_authorisation_settings = await client.MerchantAuthorisationSetting().list()
```


### MerchantDirectDebitMandate

Create an instance: `const merchant_direct_debit_mandate = client.MerchantDirectDebitMandate()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approved_at` | `string` |  |
| `currency` | `string` |  |
| `customer_account_number` | `string` |  |
| `customer_city` | `string` |  |
| `customer_country_code` | `string` |  |
| `customer_country_name` | `string` |  |
| `customer_email_address` | `string` |  |
| `customer_first_name` | `string` |  |
| `customer_iban` | `string` |  |
| `customer_last_name` | `string` |  |
| `customer_sort_code` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_recurring` | `boolean` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `reference` | `string` |  |
| `status` | `string` |  |
| `supplier_bank_account_id` | `string` |  |
| `supplier_customer_id` | `string` |  |
| `supplier_mandate_id` | `string` |  |
| `supplier_name` | `string` |  |
| `supplier_status` | `string` |  |

#### Example: List

```ts
const merchant_direct_debit_mandates = await client.MerchantDirectDebitMandate().list()
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
| `bank_country_code` | `any[]` |  |
| `bank_id` | `string` |  |
| `bank_name` | `string` |  |
| `business_institution_id` | `string` |  |
| `currency` | `string` |  |
| `logo` | `string` |  |
| `message` | `string` |  |
| `message_image_url` | `string` |  |
| `order` | `number` |  |
| `personal_institution_id` | `string` |  |
| `processor` | `string` |  |
| `warning_heading` | `string` |  |
| `warning_message` | `string` |  |

#### Example: List

```ts
const merchant_pay_by_bank_settings = await client.MerchantPayByBankSetting().list()
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
| `description` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `template` | `Record<string, any>` |  |

#### Example: Load

```ts
const merchant_payment_request_template = await client.MerchantPaymentRequestTemplate().load({ id: 'merchant_payment_request_template_id', paymentrequest_id: 'paymentrequest_id' })
```

#### Example: List

```ts
const merchant_payment_request_templates = await client.MerchantPaymentRequestTemplate().list()
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
| `authentication_method` | `any[]` |  |
| `authorisation` | `any[]` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `can_authorise` | `boolean` |  |
| `description` | `string` |  |
| `expires_at` | `string` |  |
| `has_current_user_authorised` | `boolean` |  |
| `hmac_algorithm` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `ip_address_whitelist` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_enabled` | `boolean` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `nonce` | `string` |  |
| `permission_type` | `any[]` |  |
| `request_signature_version` | `number` |  |
| `shared_secret_algorithm` | `string` |  |
| `shared_secret_base64` | `string` |  |
| `token` | `string` |  |

#### Example: Load

```ts
const merchant_token = await client.MerchantToken().load({ id: 'merchant_token_id' })
```

#### Example: List

```ts
const merchant_tokens = await client.MerchantToken().list()
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
| `build_version` | `number` |  |
| `major_version` | `number` |  |
| `minor_version` | `number` |  |
| `release_name` | `string` |  |

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
| `account_name` | `string` |  |
| `account_number` | `string` |  |
| `iban` | `string` |  |
| `payee_verified_account_name` | `string` |  |
| `result` | `string` |  |
| `secondary_identification` | `string` |  |
| `sort_code` | `string` |  |

#### Example: Create

```ts
const payeeverification = await client.Payeeverification().create({
  account_name: 'example_account_name',
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
| `address` | `any[]` |  |
| `amount` | `number` |  |
| `amount_pending` | `number` |  |
| `amount_received` | `number` |  |
| `amount_refunded` | `number` |  |
| `auto_send_receipt` | `boolean` |  |
| `base_origin_url` | `string` |  |
| `callback_url` | `string` |  |
| `card_authorize_only` | `boolean` |  |
| `card_create_token` | `boolean` |  |
| `card_create_token_mode` | `string` |  |
| `card_ignore_cvn` | `boolean` |  |
| `card_no_payer_authentication` | `boolean` |  |
| `card_processor_merchant_id` | `string` |  |
| `card_stripe_payment_intent_id` | `string` |  |
| `card_stripe_payment_intent_secret` | `string` |  |
| `card_transmit_raw_detail` | `boolean` |  |
| `created_by_user` | `Record<string, any>` |  |
| `currency` | `string` |  |
| `custom_field` | `any[]` |  |
| `customer_email_address` | `string` |  |
| `customer_id` | `string` |  |
| `customer_name` | `string` |  |
| `description` | `string` |  |
| `destination_account` | `Record<string, any>` |  |
| `direct_debit_payment` | `Record<string, any>` |  |
| `due_date` | `string` |  |
| `event` | `any[]` |  |
| `failure_callback_url` | `string` |  |
| `field_display_setting` | `any[]` |  |
| `formatted_amount` | `string` |  |
| `hosted_pay_checkout_url` | `string` |  |
| `id` | `string` |  |
| `ignore_address_verification` | `boolean` |  |
| `inserted` | `string` |  |
| `inserted_sortable` | `string` |  |
| `is_archived` | `boolean` |  |
| `jwk` | `string` |  |
| `last_updated` | `string` |  |
| `lightning_invoice` | `string` |  |
| `lightning_invoice_expires_at` | `string` |  |
| `merchant_direct_debit_mandate_id` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `notification_email_address` | `string` |  |
| `notification_role_i_d` | `any[]` |  |
| `order_id` | `string` |  |
| `partial_payment_method` | `string` |  |
| `partial_payment_step` | `string` |  |
| `payment_attempt` | `any[]` |  |
| `payment_method` | `any[]` |  |
| `payment_processor` | `string` |  |
| `payrun_id` | `string` |  |
| `pisp_account_id` | `string` |  |
| `priority_bank_id` | `string` |  |
| `result` | `Record<string, any>` |  |
| `sandbox_settle_delay_in_second` | `number` |  |
| `shipping_address` | `Record<string, any>` |  |
| `shipping_address_city` | `string` |  |
| `shipping_address_country_code` | `string` |  |
| `shipping_address_county` | `string` |  |
| `shipping_address_line1` | `string` |  |
| `shipping_address_line2` | `string` |  |
| `shipping_address_post_code` | `string` |  |
| `shipping_email` | `string` |  |
| `shipping_first_name` | `string` |  |
| `shipping_last_name` | `string` |  |
| `shipping_phone` | `string` |  |
| `status` | `string` |  |
| `success_web_hook_url` | `string` |  |
| `tag` | `any[]` |  |
| `tag_id` | `any[]` |  |
| `title` | `string` |  |
| `tokenised_card` | `any[]` |  |
| `transaction` | `any[]` |  |
| `use_hosted_payment_page` | `boolean` |  |

#### Example: Load

```ts
const payment = await client.Payment().load({ id: 'payment_id' })
```

#### Example: Create

```ts
const payment = await client.Payment().create({
  created_by_user: {},
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
| `account_name` | `string` |  |
| `account_supplier_name` | `string` |  |
| `available_balance` | `number` |  |
| `available_balance_minor_unit` | `number` |  |
| `balance` | `number` |  |
| `balance_minor_unit` | `number` |  |
| `bank_name` | `string` |  |
| `consent_id` | `string` |  |
| `created_by` | `Record<string, any>` |  |
| `created_by_display_name` | `string` |  |
| `currency` | `string` |  |
| `default_payment_rail` | `string` |  |
| `display_name` | `string` |  |
| `expiry_date` | `string` |  |
| `external_account_icon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `Record<string, any>` |  |
| `inserted` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_connected_account` | `boolean` |  |
| `is_default` | `boolean` |  |
| `is_trust_account` | `boolean` |  |
| `is_virtual` | `boolean` |  |
| `last_transaction` | `Record<string, any>` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `physical_account_id` | `string` |  |
| `rule` | `any[]` |  |
| `submitted_payouts_balance` | `number` |  |
| `submitted_payouts_balance_minor_unit` | `number` |  |
| `summary` | `string` |  |
| `supplier_sepa_instant_status` | `string` |  |
| `xero_bank_feed_connection_status` | `string` |  |
| `xero_bank_feed_last_synced_at` | `string` |  |
| `xero_bank_feed_sync_last_failed_at` | `string` |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` |  |
| `xero_bank_feed_sync_status` | `string` |  |
| `xero_unsynchronised_transactions_count` | `number` |  |

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
| `account_name` | `string` |  |
| `available_balance` | `number` |  |
| `balance` | `number` |  |
| `balance_minor_unit` | `number` |  |
| `currency` | `string` |  |
| `id` | `string` |  |
| `identifier` | `Record<string, any>` |  |
| `is_archived` | `boolean` |  |
| `is_connected_account` | `boolean` |  |
| `merchant_id` | `string` |  |
| `submitted_payouts_balance` | `number` |  |

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
| `payment_initiation_id` | `string` |  |
| `payment_request_callback_url` | `string` |  |
| `payment_request_id` | `string` |  |
| `redirect_url` | `string` |  |
| `response_type` | `string` |  |
| `specific_error_message` | `string` |  |

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
| `address` | `any[]` |  |
| `amount` | `number` |  |
| `amount_pending` | `number` |  |
| `amount_received` | `number` |  |
| `amount_refunded` | `number` |  |
| `auto_send_receipt` | `boolean` |  |
| `base_origin_url` | `string` |  |
| `callback_url` | `string` |  |
| `card_authorize_only` | `boolean` |  |
| `card_create_token` | `boolean` |  |
| `card_create_token_mode` | `string` |  |
| `card_ignore_cvn` | `boolean` |  |
| `card_processor_merchant_id` | `string` |  |
| `card_stripe_payment_intent_id` | `string` |  |
| `card_stripe_payment_intent_secret` | `string` |  |
| `created_by_user` | `Record<string, any>` |  |
| `currency` | `string` |  |
| `custom_field` | `any[]` |  |
| `customer_email_address` | `string` |  |
| `customer_id` | `string` |  |
| `customer_name` | `string` |  |
| `description` | `string` |  |
| `destination_account` | `Record<string, any>` |  |
| `direct_debit_payment` | `Record<string, any>` |  |
| `do_simulate_settlement_failure` | `boolean` |  |
| `due_date` | `string` |  |
| `error_description` | `string` |  |
| `event` | `any[]` |  |
| `failed_payment_request` | `Record<string, any>` |  |
| `failure_callback_url` | `string` |  |
| `field_display_setting` | `any[]` |  |
| `formatted_amount` | `string` |  |
| `hosted_pay_checkout_url` | `string` |  |
| `id` | `string` |  |
| `ignore_address_verification` | `boolean` |  |
| `inserted` | `string` |  |
| `inserted_sortable` | `string` |  |
| `institution` | `string` |  |
| `is_archived` | `boolean` |  |
| `jwk` | `string` |  |
| `last_updated` | `string` |  |
| `lightning_invoice` | `string` |  |
| `lightning_invoice_expires_at` | `string` |  |
| `merchant_direct_debit_mandate_id` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `notification_email_address` | `string` |  |
| `notification_role_i_d` | `any[]` |  |
| `order_id` | `string` |  |
| `partial_payment_method` | `string` |  |
| `partial_payment_step` | `string` |  |
| `payment_attempt` | `any[]` |  |
| `payment_initiation_id` | `string` |  |
| `payment_method` | `any[]` |  |
| `payment_processor` | `string` |  |
| `payment_request` | `any[]` |  |
| `payrun_id` | `string` |  |
| `pisp_account_id` | `string` |  |
| `priority_bank_id` | `string` |  |
| `result` | `Record<string, any>` |  |
| `sandbox_settle_delay_in_second` | `number` |  |
| `shipping_address` | `Record<string, any>` |  |
| `status` | `string` |  |
| `success_web_hook_url` | `string` |  |
| `tag` | `any[]` |  |
| `title` | `string` |  |
| `tokenised_card` | `any[]` |  |
| `transaction` | `any[]` |  |
| `use_hosted_payment_page` | `boolean` |  |

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
| `apple_pay_transaction_id` | `string` |  |
| `card_authorization_response_id` | `string` |  |
| `card_expiry_month` | `number` |  |
| `card_expiry_year` | `number` |  |
| `card_issuer` | `string` |  |
| `card_issuer_country` | `string` |  |
| `card_last_four_digit` | `string` |  |
| `card_request_id` | `string` |  |
| `card_scheme` | `string` |  |
| `card_token_customer_id` | `string` |  |
| `card_transaction_id` | `string` |  |
| `currency` | `string` |  |
| `direct_debit_payment_id` | `string` |  |
| `direct_debit_payment_reference` | `string` |  |
| `drirect_debit_mandate_id` | `string` |  |
| `error_message` | `string` |  |
| `error_reason` | `string` |  |
| `event_type` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lightning_invoice` | `string` |  |
| `lightning_r_hash` | `string` |  |
| `origin_url` | `string` |  |
| `payment_method_type` | `string` |  |
| `payment_processor_name` | `string` |  |
| `payment_request_id` | `string` |  |
| `pisp_bank_status` | `string` |  |
| `pisp_payment_initiation_id` | `string` |  |
| `pisp_payment_institution_name` | `string` |  |
| `pisp_payment_service_provider_id` | `string` |  |
| `pisp_redirect_url` | `string` |  |
| `reconciled_transaction_id` | `string` |  |
| `refund_payout_id` | `string` |  |
| `status` | `string` |  |
| `wallet_name` | `string` |  |

#### Example: List

```ts
const payment_request_events = await client.PaymentRequestEvent().list()
```


### PaymentRequestMetric

Create an instance: `const payment_request_metric = client.PaymentRequestMetric()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `all` | `number` |  |
| `authorized` | `number` |  |
| `paid` | `number` |  |
| `partially_paid` | `number` |  |
| `total_amounts_by_currency` | `Record<string, any>` |  |
| `unpaid` | `number` |  |

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
| `amount_pending` | `number` |  |
| `amount_received` | `number` |  |
| `amount_refunded` | `number` |  |
| `callback_url` | `string` |  |
| `card_stripe_payment_intent_secret` | `string` |  |
| `country_code` | `string` |  |
| `currency` | `string` |  |
| `custom_fields_to_display` | `any[]` |  |
| `description` | `string` |  |
| `due_date` | `string` |  |
| `field_display_setting` | `any[]` |  |
| `google_pay_merchant_id` | `string` |  |
| `id` | `string` |  |
| `jwk` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_logo_url_png` | `string` |  |
| `merchant_logo_url_svg` | `string` |  |
| `merchant_name` | `string` |  |
| `merchant_short_name` | `string` |  |
| `partial_payment_method` | `string` |  |
| `payment_attempt` | `any[]` |  |
| `payment_methods_list` | `any[]` |  |
| `payment_processor` | `string` |  |
| `payment_processor_key` | `string` |  |
| `pisp_error` | `string` |  |
| `priority_bank_id` | `string` |  |
| `status` | `string` |  |
| `stripe_account_id` | `string` |  |
| `title` | `string` |  |

#### Example: List

```ts
const payment_request_minimals = await client.PaymentRequestMinimal().list()
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
| `amount_pending` | `number` |  |
| `amount_received` | `number` |  |
| `amount_refunded` | `number` |  |
| `currency` | `string` |  |
| `customer_id` | `string` |  |
| `payment` | `any[]` |  |
| `payment_request_id` | `string` |  |
| `pisp_authorization` | `any[]` |  |
| `requested_amount` | `number` |  |
| `result` | `string` |  |

#### Example: List

```ts
const payment_request_results = await client.PaymentRequestResult().list()
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
| `account_id` | `string` |  |
| `allow_incomplete` | `boolean` |  |
| `amount` | `number` |  |
| `amount_minor_unit` | `number` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `any[]` |  |
| `authorisation` | `any[]` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `Record<string, any>` |  |
| `beneficiary_id` | `string` |  |
| `can_authorise` | `boolean` |  |
| `can_process` | `boolean` |  |
| `can_update` | `boolean` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `Record<string, any>` |  |
| `document` | `any[]` |  |
| `event` | `any[]` |  |
| `failed_payout` | `Record<string, any>` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `number` |  |
| `fx_destination_amount_minor_unit` | `number` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `number` |  |
| `fx_use_destination_amount` | `boolean` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_failed` | `boolean` |  |
| `is_settled` | `boolean` |  |
| `is_submitted` | `boolean` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payout` | `any[]` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `reason` | `string` |  |
| `rule` | `Record<string, any>` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `boolean` |  |
| `source_account_available_balance` | `number` |  |
| `source_account_available_balance_minor_unit` | `number` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `Record<string, any>` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `any[]` |  |
| `tag_id` | `any[]` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `number` |  |
| `transacted_fx_amount` | `number` |  |
| `transacted_fx_rate` | `number` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |

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
})
```


### PayoutKeyset

Create an instance: `const payout_keyset = client.PayoutKeyset()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `amount` | `number` |  |
| `amount_minor_unit` | `number` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `any[]` |  |
| `authorisation` | `any[]` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `Record<string, any>` |  |
| `can_authorise` | `boolean` |  |
| `can_process` | `boolean` |  |
| `can_update` | `boolean` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `Record<string, any>` |  |
| `document` | `any[]` |  |
| `event` | `any[]` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `number` |  |
| `fx_destination_amount_minor_unit` | `number` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `number` |  |
| `fx_use_destination_amount` | `boolean` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_failed` | `boolean` |  |
| `is_settled` | `boolean` |  |
| `is_submitted` | `boolean` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `rule` | `Record<string, any>` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `boolean` |  |
| `source_account_available_balance` | `number` |  |
| `source_account_available_balance_minor_unit` | `number` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `Record<string, any>` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `any[]` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `number` |  |
| `transacted_fx_amount` | `number` |  |
| `transacted_fx_rate` | `number` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |

#### Example: List

```ts
const payout_keysets = await client.PayoutKeyset().list()
```


### PayoutMetric

Create an instance: `const payout_metric = client.PayoutMetric()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `all` | `number` |  |
| `failed` | `number` |  |
| `in_progress` | `number` |  |
| `paid` | `number` |  |
| `pending_approval` | `number` |  |
| `scheduled` | `number` |  |
| `total_amounts_by_currency` | `Record<string, any>` |  |

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
| `authorisation` | `any[]` |  |
| `authorisation_date` | `string` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `batch_payout_id` | `string` |  |
| `can_authorise` | `boolean` |  |
| `can_delete` | `boolean` |  |
| `can_edit` | `boolean` |  |
| `event` | `any[]` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice` | `any[]` |  |
| `invoices_minimal` | `any[]` |  |
| `is_archived` | `boolean` |  |
| `last_updated` | `string` |  |
| `last_updated_by` | `Record<string, any>` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `note` | `string` |  |
| `payment` | `any[]` |  |
| `payout` | `any[]` |  |
| `payouts_count` | `number` |  |
| `reason` | `string` |  |
| `schedule_date` | `string` |  |
| `scheduled_date` | `string` |  |
| `source_account` | `any[]` |  |
| `status` | `string` |  |
| `total_eur` | `number` |  |
| `total_gbp` | `number` |  |
| `total_usd` | `number` |  |

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
| `content` | `string` |  |
| `content_type` | `string` |  |
| `last_completed_at` | `string` |  |
| `merchant_id` | `string` |  |
| `report_name` | `string` |  |
| `report_type` | `string` |  |
| `statement_number` | `number` |  |

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
| `failed_role` | `Record<string, any>` |  |
| `role` | `any[]` |  |

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
| `account_id` | `string` |  |
| `approve_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `any[]` |  |
| `authorisation` | `any[]` |  |
| `authorisers_completed_count` | `number` |  |
| `authorisers_required_count` | `number` |  |
| `can_authorise` | `boolean` |  |
| `created_by` | `Record<string, any>` |  |
| `description` | `string` |  |
| `end_at` | `string` |  |
| `has_current_user_authorised` | `boolean` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_disabled` | `boolean` |  |
| `last_executed_at` | `string` |  |
| `last_run_at_transaction_date` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `on_approved_web_hook_url` | `string` |  |
| `on_execution_error_web_hook_url` | `string` |  |
| `on_execution_success_web_hook_url` | `string` |  |
| `start_at` | `string` |  |
| `status` | `string` |  |
| `sweep_action` | `Record<string, any>` |  |
| `time_zone_id` | `string` |  |
| `trigger_cron_expression` | `string` |  |
| `trigger_on_pay_in` | `boolean` |  |
| `user_id` | `string` |  |
| `web_hook_secret` | `string` |  |

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
  created_by: {},
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
| `error_message` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_authorise_to_enable` | `boolean` |  |
| `message` | `string` |  |
| `raw_response` | `string` |  |
| `rule_event_type` | `string` |  |
| `rule_id` | `string` |  |
| `user` | `Record<string, any>` |  |

#### Example: List

```ts
const rule_events = await client.RuleEvent().list()
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
| `colour_hex` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |

#### Example: List

```ts
const tags = await client.Tag().list()
```

#### Example: Create

```ts
const tag = await client.Tag().create({
  merchant_id: 'example_merchant_id',
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
| `account_id` | `string` |  |
| `account_name` | `string` |  |
| `account_sequence_number` | `number` |  |
| `address_detail` | `Record<string, any>` |  |
| `amount` | `number` |  |
| `amount_minor_unit` | `number` |  |
| `balance` | `number` |  |
| `balance_minor_unit` | `number` |  |
| `booking_date_time` | `string` |  |
| `charge_detail` | `Record<string, any>` |  |
| `content` | `any[]` |  |
| `counterparty` | `Record<string, any>` |  |
| `counterparty_summary` | `string` |  |
| `currency` | `string` |  |
| `currency_exchange` | `Record<string, any>` |  |
| `date` | `string` |  |
| `description` | `string` |  |
| `enrichment` | `Record<string, any>` |  |
| `fx_amount` | `number` |  |
| `fx_currency` | `string` |  |
| `fx_rate` | `number` |  |
| `gross_amount` | `Record<string, any>` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `iso_bank_transaction_code` | `Record<string, any>` |  |
| `merchant` | `Record<string, any>` |  |
| `merchant_id` | `string` |  |
| `page_number` | `number` |  |
| `page_size` | `number` |  |
| `payee_detail` | `Record<string, any>` |  |
| `payer_detail` | `Record<string, any>` |  |
| `payment_request_custom_field` | `Record<string, any>` |  |
| `payment_request_id` | `string` |  |
| `payout_id` | `string` |  |
| `proprietary_bank_transaction_code` | `Record<string, any>` |  |
| `raw_reference` | `string` |  |
| `reference` | `string` |  |
| `rule_id` | `string` |  |
| `statement_reference` | `any[]` |  |
| `status` | `string` |  |
| `supplementary_data` | `any` |  |
| `tag` | `any[]` |  |
| `their_reference` | `string` |  |
| `total_page` | `number` |  |
| `total_size` | `number` |  |
| `transaction_amount` | `Record<string, any>` |  |
| `transaction_date` | `string` |  |
| `transaction_information` | `any[]` |  |
| `transaction_mutability` | `string` |  |
| `type` | `string` |  |
| `value_date_time` | `string` |  |
| `virtual_iban` | `string` |  |
| `your_reference` | `string` |  |

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
| `client_session_timeout` | `any[]` |  |
| `email_address` | `string` |  |
| `first_name` | `string` |  |
| `id` | `string` |  |
| `last_name` | `string` |  |
| `passkey_added` | `boolean` |  |
| `permission` | `Record<string, any>` |  |
| `profile` | `string` |  |
| `roles_with_scope` | `any[]` |  |
| `two_factor_enabled` | `boolean` |  |
| `user_invite_id` | `string` |  |

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
| `authorisation_status` | `Record<string, any>` |  |
| `failed_user_invite` | `Record<string, any>` |  |
| `id` | `string` |  |
| `initial_role_id` | `string` |  |
| `invitee_email_address` | `string` |  |
| `invitee_first_name` | `string` |  |
| `invitee_last_name` | `string` |  |
| `inviter_email_address` | `string` |  |
| `inviter_first_name` | `string` |  |
| `inviter_last_name` | `string` |  |
| `is_authorised` | `boolean` |  |
| `is_invitee_registered` | `boolean` |  |
| `last_invited` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `message` | `string` |  |
| `registration_url` | `string` |  |
| `send_invite_email` | `boolean` |  |
| `status` | `string` |  |
| `user` | `Record<string, any>` |  |
| `user_id` | `string` |  |
| `user_invite` | `any[]` |  |

#### Example: Load

```ts
const user_invite = await client.UserInvite().load({ id: 'user_invite_id' })
```

#### Example: List

```ts
const user_invites = await client.UserInvite().list()
```

#### Example: Create

```ts
const user_invite = await client.UserInvite().create({
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
| `account_name` | `string` |  |
| `account_supplier_name` | `string` |  |
| `available_balance` | `number` |  |
| `available_balance_minor_unit` | `number` |  |
| `balance` | `number` |  |
| `balance_minor_unit` | `number` |  |
| `bank_name` | `string` |  |
| `consent_id` | `string` |  |
| `created_by` | `Record<string, any>` |  |
| `created_by_display_name` | `string` |  |
| `currency` | `string` |  |
| `default_payment_rail` | `string` |  |
| `display_name` | `string` |  |
| `expiry_date` | `string` |  |
| `external_account_icon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `Record<string, any>` |  |
| `inserted` | `string` |  |
| `is_archived` | `boolean` |  |
| `is_connected_account` | `boolean` |  |
| `is_default` | `boolean` |  |
| `is_trust_account` | `boolean` |  |
| `is_virtual` | `boolean` |  |
| `last_transaction` | `Record<string, any>` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `name` | `string` |  |
| `physical_account_id` | `string` |  |
| `rule` | `any[]` |  |
| `submitted_payouts_balance` | `number` |  |
| `submitted_payouts_balance_minor_unit` | `number` |  |
| `summary` | `string` |  |
| `supplier_sepa_instant_status` | `string` |  |
| `xero_bank_feed_connection_status` | `string` |  |
| `xero_bank_feed_last_synced_at` | `string` |  |
| `xero_bank_feed_sync_last_failed_at` | `string` |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` |  |
| `xero_bank_feed_sync_status` | `string` |  |
| `xero_unsynchronised_transactions_count` | `number` |  |

#### Example: Create

```ts
const virtual = await client.Virtual().create({
  account_id: 'example_account_id',
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
| `destination_url` | `string` |  |
| `email_address` | `string` |  |
| `failed_notification_email_address` | `string` |  |
| `id` | `string` |  |
| `is_active` | `boolean` |  |
| `merchant_id` | `string` |  |
| `notification_method` | `string` |  |
| `resource_type` | `any[]` |  |
| `retry` | `boolean` |  |
| `secret` | `string` |  |
| `version` | `number` |  |

#### Example: Load

```ts
const webhook = await client.Webhook().load({ id: 'webhook_id' })
```

#### Example: List

```ts
const webhooks = await client.Webhook().list()
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
const account = client.Account()
await account.list()

// account.data() now returns the account data from the last `list`
// account.match() returns the last match criteria
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
