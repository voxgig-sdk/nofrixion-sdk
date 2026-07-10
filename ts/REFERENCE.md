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
| `account_balance` | `any[]` | No |  |
| `account_id` | `string` | No |  |
| `account_identification` | `any[]` | No |  |
| `account_name` | `string` | No |  |
| `account_supplier_name` | `string` | No |  |
| `account_type` | `string` | No |  |
| `available_balance` | `number` | No |  |
| `available_balance_minor_unit` | `number` | No |  |
| `balance` | `number` | No |  |
| `balance_minor_unit` | `number` | No |  |
| `bank_name` | `string` | No |  |
| `consent_id` | `string` | No |  |
| `consolidated_account_information` | `Record<string, any>` | No |  |
| `created_by` | `Record<string, any>` | Yes |  |
| `created_by_display_name` | `string` | No |  |
| `currency` | `string` | No |  |
| `default_payment_rail` | `string` | No |  |
| `description` | `string` | No |  |
| `detail` | `string` | No |  |
| `display_name` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `external_account_icon` | `string` | No |  |
| `format` | `string` | No |  |
| `from_date` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `Record<string, any>` | Yes |  |
| `inserted` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_connected_account` | `boolean` | No |  |
| `is_default` | `boolean` | No |  |
| `is_trust_account` | `boolean` | No |  |
| `is_virtual` | `boolean` | No |  |
| `last_transaction` | `Record<string, any>` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `nickname` | `string` | No |  |
| `physical_account_id` | `string` | No |  |
| `role_i_d` | `any[]` | No |  |
| `rule` | `any[]` | No |  |
| `submitted_payouts_balance` | `number` | No |  |
| `submitted_payouts_balance_minor_unit` | `number` | No |  |
| `summary` | `string` | No |  |
| `supplier_physical_account_id` | `string` | No |  |
| `supplier_sepa_instant_status` | `string` | No |  |
| `to_date` | `string` | No |  |
| `type` | `string` | No |  |
| `usage_type` | `string` | No |  |
| `xero_bank_feed_connection_status` | `string` | No |  |
| `xero_bank_feed_last_synced_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` | No |  |
| `xero_bank_feed_sync_status` | `string` | No |  |
| `xero_unsynchronised_transactions_count` | `number` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Account().create({
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
| `approve_url` | `string` | No |  |
| `id` | `string` | No |  |
| `payout` | `any[]` | No |  |

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
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `beneficiary` | `any[]` | No |  |
| `beneficiary_event` | `any[]` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `created_by` | `Record<string, any>` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `Record<string, any>` | No |  |
| `failed_beneficiary` | `Record<string, any>` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `boolean` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `any[]` | No |  |
| `source_account_i_d` | `any[]` | No |  |
| `their_reference` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `approval_callback_url` | - | - | - | - | - |
| `authentication_method` | - | - | - | - | - |
| `authorisation` | - | - | - | - | - |
| `authorisers_completed_count` | - | - | - | - | - |
| `authorisers_required_count` | - | - | - | - | - |
| `beneficiary` | - | - | - | - | - |
| `beneficiary_event` | - | - | - | - | - |
| `can_authorise` | - | - | - | - | - |
| `can_update` | - | - | - | - | - |
| `created_by` | - | - | - | - | - |
| `created_by_email_address` | - | - | - | - | - |
| `currency` | - | - | - | Yes | - |
| `destination` | - | - | Yes | - | - |
| `failed_beneficiary` | - | - | - | - | - |
| `has_current_user_authorised` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `inserted` | - | - | - | - | - |
| `is_enabled` | - | - | - | - | - |
| `last_authorised` | - | - | - | - | - |
| `last_updated` | - | - | - | - | - |
| `merchant_id` | - | - | Yes | - | - |
| `name` | - | - | - | Yes | - |
| `nonce` | - | - | - | - | - |
| `source_account` | - | - | - | - | - |
| `source_account_i_d` | - | - | - | - | - |
| `their_reference` | - | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Beneficiary().create({
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
| `group_member` | `any[]` | No |  |
| `group_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.BeneficiaryGroup().list()
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
| `authorized_amount` | `string` | No |  |
| `currency_code` | `string` | No |  |
| `is_payer_authentication_required` | `boolean` | No |  |
| `is_soft_decline` | `boolean` | No |  |
| `payer_authentication_access_token` | `string` | No |  |
| `payer_authentication_merchant_data` | `string` | No |  |
| `payer_authentication_url` | `string` | No |  |
| `payer_authentication_window_height` | `number` | No |  |
| `payer_authentication_window_width` | `number` | No |  |
| `payment_request_callback_url` | `string` | No |  |
| `payment_request_id` | `string` | No |  |
| `request_id` | `string` | No |  |
| `response_code` | `string` | No |  |
| `response_type` | `string` | No |  |
| `status` | `string` | No |  |
| `three_ds_redirect_url` | `string` | No |  |
| `transaction_id` | `string` | No |  |

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
| `card_type` | `string` | No |  |
| `customer_email_address` | `string` | No |  |
| `expiry_month` | `string` | No |  |
| `expiry_year` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `last_four_digit` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `masked_card_number` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `payment_request_id` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.CardCustomerToken().list()
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
| `authorized_amount` | `string` | No |  |
| `currency_code` | `string` | No |  |
| `is_payer_authentication_required` | `boolean` | No |  |
| `is_soft_decline` | `boolean` | No |  |
| `payer_authentication_access_token` | `string` | No |  |
| `payer_authentication_merchant_data` | `string` | No |  |
| `payer_authentication_url` | `string` | No |  |
| `payer_authentication_window_height` | `number` | No |  |
| `payer_authentication_window_width` | `number` | No |  |
| `payment_request_callback_url` | `string` | No |  |
| `payment_request_id` | `string` | No |  |
| `request_id` | `string` | No |  |
| `response_code` | `string` | No |  |
| `response_type` | `string` | No |  |
| `status` | `string` | No |  |
| `three_ds_redirect_url` | `string` | No |  |
| `transaction_id` | `string` | No |  |

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
| `authorisation_url` | `string` | No |  |
| `callback_url` | `string` | No |  |
| `consent_id` | `string` | No |  |
| `email_address` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `failure_callback_url` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `institution_id` | `string` | No |  |
| `is_connected_account` | `boolean` | No |  |
| `is_enabled` | `boolean` | No |  |
| `merchant_id` | `string` | No |  |
| `provider` | `string` | No |  |
| `success_web_hook_url` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `authorisation_url` | - | - | - | - | - |
| `callback_url` | - | - | - | - | - |
| `consent_id` | - | - | - | - | - |
| `email_address` | - | - | - | - | - |
| `expiry_date` | - | - | - | - | - |
| `failure_callback_url` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `inserted` | - | - | - | - | - |
| `institution_id` | - | - | Yes | - | - |
| `is_connected_account` | - | - | - | - | - |
| `is_enabled` | - | - | - | - | - |
| `merchant_id` | - | - | Yes | - | - |
| `provider` | - | - | - | - | - |
| `success_web_hook_url` | - | - | - | - | - |

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
const results = await client.Consent().list()
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
| `decimal` | `number` | No |  |
| `is_fiat` | `boolean` | No |  |
| `iso4217_alpha_code` | `string` | No |  |
| `iso4217_numeric_code` | `string` | No |  |
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
| `failed_submission` | `Record<string, any>` | No |  |
| `successful_submission` | `any[]` | No |  |

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
| `destination_currency` | `string` | No |  |
| `exchange_rate` | `number` | No |  |
| `expiry_time` | `string` | No |  |
| `quote_id` | `string` | No |  |
| `source_currency` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.FxRate().list()
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
| `payment_request_id` | `string` | No |  |
| `response_type` | `string` | No |  |

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
| `account_number` | `string` | No |  |
| `address_line1` | `string` | Yes |  |
| `address_line2` | `string` | No |  |
| `approved_at` | `string` | No |  |
| `city` | `string` | Yes |  |
| `country_code` | `string` | Yes |  |
| `currency` | `string` | No |  |
| `customer_account_number` | `string` | No |  |
| `customer_city` | `string` | No |  |
| `customer_country_code` | `string` | No |  |
| `customer_country_name` | `string` | No |  |
| `customer_email_address` | `string` | No |  |
| `customer_first_name` | `string` | No |  |
| `customer_iban` | `string` | No |  |
| `customer_last_name` | `string` | No |  |
| `customer_sort_code` | `string` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `iban` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_recurring` | `boolean` | No |  |
| `last_name` | `string` | Yes |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `postal_code` | `string` | Yes |  |
| `reference` | `string` | No |  |
| `sort_code` | `string` | No |  |
| `status` | `string` | No |  |
| `supplier_bank_account_id` | `string` | No |  |
| `supplier_customer_id` | `string` | No |  |
| `supplier_mandate_id` | `string` | No |  |
| `supplier_name` | `string` | No |  |
| `supplier_status` | `string` | No |  |

### Field Usage by Operation

| Field | load | create |
| --- | --- | --- |
| `account_number` | - | - |
| `address_line1` | - | - |
| `address_line2` | - | - |
| `approved_at` | - | - |
| `city` | - | - |
| `country_code` | - | - |
| `currency` | - | Yes |
| `customer_account_number` | - | - |
| `customer_city` | - | - |
| `customer_country_code` | - | - |
| `customer_country_name` | - | - |
| `customer_email_address` | - | - |
| `customer_first_name` | - | - |
| `customer_iban` | - | - |
| `customer_last_name` | - | - |
| `customer_sort_code` | - | - |
| `email_address` | - | - |
| `first_name` | - | - |
| `iban` | - | - |
| `id` | - | - |
| `inserted` | - | - |
| `is_recurring` | - | - |
| `last_name` | - | - |
| `last_updated` | - | - |
| `merchant_id` | - | Yes |
| `postal_code` | - | - |
| `reference` | - | - |
| `sort_code` | - | - |
| `status` | - | - |
| `supplier_bank_account_id` | - | - |
| `supplier_customer_id` | - | - |
| `supplier_mandate_id` | - | - |
| `supplier_name` | - | - |
| `supplier_status` | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Mandate().create({
  address_line1: 'example_address_line1',
  city: 'example_city',
  country_code: 'example_country_code',
  email_address: 'example_email_address',
  first_name: 'example_first_name',
  last_name: 'example_last_name',
  postal_code: 'example_postal_code',
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
| `account_currency` | `any[]` | No |  |
| `can_have_trust_account` | `boolean` | No |  |
| `card_payment_processor` | `string` | No |  |
| `company_id` | `string` | No |  |
| `display_qr_on_hosted_pay` | `boolean` | No |  |
| `hosted_pay_version` | `number` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_blocked` | `boolean` | No |  |
| `is_exited` | `boolean` | No |  |
| `is_suspended` | `boolean` | No |  |
| `jurisdiction` | `string` | No |  |
| `logo_url_png` | `string` | No |  |
| `logo_url_svg` | `string` | No |  |
| `merchant_category_code` | `string` | No |  |
| `name` | `string` | No |  |
| `note` | `string` | No |  |
| `parent_merchant` | `Record<string, any>` | No |  |
| `payment_account` | `any[]` | No |  |
| `payment_account_limit` | `number` | No |  |
| `reason` | `string` | No |  |
| `short_name` | `string` | No |  |
| `supported_payment_methods_list` | `any[]` | No |  |
| `suspension_reason` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `time_zone_id` | `string` | No |  |
| `trading_name` | `string` | No |  |
| `web_hook_limit` | `number` | No |  |
| `your_role_name` | `string` | No |  |

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
| `amount_lower` | `number` | No |  |
| `amount_upper` | `number` | No |  |
| `authorisation_type` | `string` | No |  |
| `beneficiaries_only` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `last_editor_cant_authorise` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `number_of_authoriser` | `number` | No |  |
| `role_setting` | `any[]` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.MerchantAuthorisationSetting().list()
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
| `approved_at` | `string` | No |  |
| `currency` | `string` | No |  |
| `customer_account_number` | `string` | No |  |
| `customer_city` | `string` | No |  |
| `customer_country_code` | `string` | No |  |
| `customer_country_name` | `string` | No |  |
| `customer_email_address` | `string` | No |  |
| `customer_first_name` | `string` | No |  |
| `customer_iban` | `string` | No |  |
| `customer_last_name` | `string` | No |  |
| `customer_sort_code` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_recurring` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `reference` | `string` | No |  |
| `status` | `string` | No |  |
| `supplier_bank_account_id` | `string` | No |  |
| `supplier_customer_id` | `string` | No |  |
| `supplier_mandate_id` | `string` | No |  |
| `supplier_name` | `string` | No |  |
| `supplier_status` | `string` | No |  |

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
| `bank_country_code` | `any[]` | No |  |
| `bank_id` | `string` | No |  |
| `bank_name` | `string` | No |  |
| `business_institution_id` | `string` | No |  |
| `currency` | `string` | No |  |
| `logo` | `string` | No |  |
| `message` | `string` | No |  |
| `message_image_url` | `string` | No |  |
| `order` | `number` | No |  |
| `personal_institution_id` | `string` | No |  |
| `processor` | `string` | No |  |
| `warning_heading` | `string` | No |  |
| `warning_message` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.MerchantPayByBankSetting().list()
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
| `description` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `template` | `Record<string, any>` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.MerchantPaymentRequestTemplate().list()
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
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `can_authorise` | `boolean` | No |  |
| `description` | `string` | No |  |
| `expires_at` | `string` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `hmac_algorithm` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `ip_address_whitelist` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_enabled` | `boolean` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `nonce` | `string` | Yes |  |
| `permission_type` | `any[]` | No |  |
| `request_signature_version` | `number` | No |  |
| `shared_secret_algorithm` | `string` | No |  |
| `shared_secret_base64` | `string` | No |  |
| `token` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update |
| --- | --- | --- | --- | --- |
| `authentication_method` | - | - | - | - |
| `authorisation` | - | - | - | - |
| `authorisers_completed_count` | - | - | - | - |
| `authorisers_required_count` | - | - | - | - |
| `can_authorise` | - | - | - | - |
| `description` | - | - | Yes | - |
| `expires_at` | - | - | - | - |
| `has_current_user_authorised` | - | - | - | - |
| `hmac_algorithm` | - | - | - | - |
| `id` | - | - | - | - |
| `inserted` | - | - | - | - |
| `ip_address_whitelist` | - | - | - | - |
| `is_archived` | - | - | - | - |
| `is_enabled` | - | - | - | - |
| `last_authorised` | - | - | - | - |
| `last_updated` | - | - | - | - |
| `merchant_id` | - | - | Yes | - |
| `nonce` | - | - | - | - |
| `permission_type` | - | - | - | - |
| `request_signature_version` | - | - | - | - |
| `shared_secret_algorithm` | - | - | - | - |
| `shared_secret_base64` | - | - | - | - |
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
const results = await client.MerchantToken().list()
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
| `build_version` | `number` | No |  |
| `major_version` | `number` | No |  |
| `minor_version` | `number` | No |  |
| `release_name` | `string` | No |  |

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
| `account_name` | `string` | Yes |  |
| `account_number` | `string` | No |  |
| `iban` | `string` | Yes |  |
| `payee_verified_account_name` | `string` | No |  |
| `result` | `string` | No |  |
| `secondary_identification` | `string` | No |  |
| `sort_code` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Payeeverification().create({
  account_name: 'example_account_name',
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
| `address` | `any[]` | No |  |
| `amount` | `number` | No |  |
| `amount_pending` | `number` | No |  |
| `amount_received` | `number` | No |  |
| `amount_refunded` | `number` | No |  |
| `auto_send_receipt` | `boolean` | No |  |
| `base_origin_url` | `string` | No |  |
| `callback_url` | `string` | No |  |
| `card_authorize_only` | `boolean` | No |  |
| `card_create_token` | `boolean` | No |  |
| `card_create_token_mode` | `string` | No |  |
| `card_ignore_cvn` | `boolean` | No |  |
| `card_no_payer_authentication` | `boolean` | No |  |
| `card_processor_merchant_id` | `string` | No |  |
| `card_stripe_payment_intent_id` | `string` | No |  |
| `card_stripe_payment_intent_secret` | `string` | No |  |
| `card_transmit_raw_detail` | `boolean` | No |  |
| `created_by_user` | `Record<string, any>` | Yes |  |
| `currency` | `string` | No |  |
| `custom_field` | `any[]` | No |  |
| `customer_email_address` | `string` | No |  |
| `customer_id` | `string` | No |  |
| `customer_name` | `string` | No |  |
| `description` | `string` | No |  |
| `destination_account` | `Record<string, any>` | No |  |
| `direct_debit_payment` | `Record<string, any>` | No |  |
| `due_date` | `string` | No |  |
| `event` | `any[]` | No |  |
| `failure_callback_url` | `string` | No |  |
| `field_display_setting` | `any[]` | No |  |
| `formatted_amount` | `string` | No |  |
| `hosted_pay_checkout_url` | `string` | No |  |
| `id` | `string` | No |  |
| `ignore_address_verification` | `boolean` | No |  |
| `inserted` | `string` | No |  |
| `inserted_sortable` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `jwk` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `lightning_invoice` | `string` | No |  |
| `lightning_invoice_expires_at` | `string` | No |  |
| `merchant_direct_debit_mandate_id` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `notification_email_address` | `string` | No |  |
| `notification_role_i_d` | `any[]` | No |  |
| `order_id` | `string` | No |  |
| `partial_payment_method` | `string` | No |  |
| `partial_payment_step` | `string` | No |  |
| `payment_attempt` | `any[]` | No |  |
| `payment_method` | `any[]` | No |  |
| `payment_processor` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `pisp_account_id` | `string` | No |  |
| `priority_bank_id` | `string` | No |  |
| `result` | `Record<string, any>` | No |  |
| `sandbox_settle_delay_in_second` | `number` | No |  |
| `shipping_address` | `Record<string, any>` | No |  |
| `shipping_address_city` | `string` | No |  |
| `shipping_address_country_code` | `string` | No |  |
| `shipping_address_county` | `string` | No |  |
| `shipping_address_line1` | `string` | No |  |
| `shipping_address_line2` | `string` | No |  |
| `shipping_address_post_code` | `string` | No |  |
| `shipping_email` | `string` | No |  |
| `shipping_first_name` | `string` | No |  |
| `shipping_last_name` | `string` | No |  |
| `shipping_phone` | `string` | No |  |
| `status` | `string` | No |  |
| `success_web_hook_url` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `tag_id` | `any[]` | No |  |
| `title` | `string` | No |  |
| `tokenised_card` | `any[]` | No |  |
| `transaction` | `any[]` | No |  |
| `use_hosted_payment_page` | `boolean` | No |  |

### Field Usage by Operation

| Field | load | create | update |
| --- | --- | --- | --- |
| `address` | - | - | - |
| `amount` | - | Yes | - |
| `amount_pending` | - | - | - |
| `amount_received` | - | - | - |
| `amount_refunded` | - | - | - |
| `auto_send_receipt` | - | - | - |
| `base_origin_url` | - | - | - |
| `callback_url` | - | - | - |
| `card_authorize_only` | - | - | - |
| `card_create_token` | - | - | - |
| `card_create_token_mode` | - | - | - |
| `card_ignore_cvn` | - | - | - |
| `card_no_payer_authentication` | - | - | - |
| `card_processor_merchant_id` | - | - | - |
| `card_stripe_payment_intent_id` | - | - | - |
| `card_stripe_payment_intent_secret` | - | - | - |
| `card_transmit_raw_detail` | - | - | - |
| `created_by_user` | - | - | - |
| `currency` | - | - | - |
| `custom_field` | - | - | - |
| `customer_email_address` | - | - | - |
| `customer_id` | - | - | - |
| `customer_name` | - | - | - |
| `description` | - | - | - |
| `destination_account` | - | - | - |
| `direct_debit_payment` | - | - | - |
| `due_date` | - | - | - |
| `event` | - | - | - |
| `failure_callback_url` | - | - | - |
| `field_display_setting` | - | - | - |
| `formatted_amount` | - | - | - |
| `hosted_pay_checkout_url` | - | - | - |
| `id` | - | - | - |
| `ignore_address_verification` | - | - | - |
| `inserted` | - | - | - |
| `inserted_sortable` | - | - | - |
| `is_archived` | - | - | - |
| `jwk` | - | - | - |
| `last_updated` | - | - | - |
| `lightning_invoice` | - | - | - |
| `lightning_invoice_expires_at` | - | - | - |
| `merchant_direct_debit_mandate_id` | - | - | - |
| `merchant_id` | - | - | - |
| `merchant_token_description` | - | - | - |
| `notification_email_address` | - | - | - |
| `notification_role_i_d` | - | - | - |
| `order_id` | - | - | - |
| `partial_payment_method` | - | - | - |
| `partial_payment_step` | - | - | - |
| `payment_attempt` | - | - | - |
| `payment_method` | - | - | - |
| `payment_processor` | - | - | - |
| `payrun_id` | - | - | - |
| `pisp_account_id` | - | - | - |
| `priority_bank_id` | - | - | - |
| `result` | - | - | - |
| `sandbox_settle_delay_in_second` | - | - | - |
| `shipping_address` | - | - | - |
| `shipping_address_city` | - | - | - |
| `shipping_address_country_code` | - | - | - |
| `shipping_address_county` | - | - | - |
| `shipping_address_line1` | - | - | - |
| `shipping_address_line2` | - | - | - |
| `shipping_address_post_code` | - | - | - |
| `shipping_email` | - | - | - |
| `shipping_first_name` | - | - | - |
| `shipping_last_name` | - | - | - |
| `shipping_phone` | - | - | - |
| `status` | - | - | - |
| `success_web_hook_url` | - | - | - |
| `tag` | - | - | - |
| `tag_id` | - | - | - |
| `title` | - | - | - |
| `tokenised_card` | - | - | - |
| `transaction` | - | - | - |
| `use_hosted_payment_page` | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Payment().create({
  created_by_user: {},
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
| `account_name` | `string` | No |  |
| `account_supplier_name` | `string` | No |  |
| `available_balance` | `number` | No |  |
| `available_balance_minor_unit` | `number` | No |  |
| `balance` | `number` | No |  |
| `balance_minor_unit` | `number` | No |  |
| `bank_name` | `string` | No |  |
| `consent_id` | `string` | No |  |
| `created_by` | `Record<string, any>` | Yes |  |
| `created_by_display_name` | `string` | No |  |
| `currency` | `string` | No |  |
| `default_payment_rail` | `string` | No |  |
| `display_name` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `external_account_icon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `Record<string, any>` | Yes |  |
| `inserted` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_connected_account` | `boolean` | No |  |
| `is_default` | `boolean` | No |  |
| `is_trust_account` | `boolean` | No |  |
| `is_virtual` | `boolean` | No |  |
| `last_transaction` | `Record<string, any>` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `physical_account_id` | `string` | No |  |
| `rule` | `any[]` | No |  |
| `submitted_payouts_balance` | `number` | No |  |
| `submitted_payouts_balance_minor_unit` | `number` | No |  |
| `summary` | `string` | No |  |
| `supplier_sepa_instant_status` | `string` | No |  |
| `xero_bank_feed_connection_status` | `string` | No |  |
| `xero_bank_feed_last_synced_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` | No |  |
| `xero_bank_feed_sync_status` | `string` | No |  |
| `xero_unsynchronised_transactions_count` | `number` | No |  |

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
| `account_name` | `string` | No |  |
| `available_balance` | `number` | No |  |
| `balance` | `number` | No |  |
| `balance_minor_unit` | `number` | No |  |
| `currency` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `Record<string, any>` | Yes |  |
| `is_archived` | `boolean` | No |  |
| `is_connected_account` | `boolean` | No |  |
| `merchant_id` | `string` | No |  |
| `submitted_payouts_balance` | `number` | No |  |

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
| `payment_initiation_id` | `string` | No |  |
| `payment_request_callback_url` | `string` | No |  |
| `payment_request_id` | `string` | No |  |
| `redirect_url` | `string` | No |  |
| `response_type` | `string` | No |  |
| `specific_error_message` | `string` | No |  |

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
| `address` | `any[]` | No |  |
| `amount` | `number` | No |  |
| `amount_pending` | `number` | No |  |
| `amount_received` | `number` | No |  |
| `amount_refunded` | `number` | No |  |
| `auto_send_receipt` | `boolean` | No |  |
| `base_origin_url` | `string` | No |  |
| `callback_url` | `string` | No |  |
| `card_authorize_only` | `boolean` | No |  |
| `card_create_token` | `boolean` | No |  |
| `card_create_token_mode` | `string` | No |  |
| `card_ignore_cvn` | `boolean` | No |  |
| `card_processor_merchant_id` | `string` | No |  |
| `card_stripe_payment_intent_id` | `string` | No |  |
| `card_stripe_payment_intent_secret` | `string` | No |  |
| `created_by_user` | `Record<string, any>` | Yes |  |
| `currency` | `string` | No |  |
| `custom_field` | `any[]` | No |  |
| `customer_email_address` | `string` | No |  |
| `customer_id` | `string` | No |  |
| `customer_name` | `string` | No |  |
| `description` | `string` | No |  |
| `destination_account` | `Record<string, any>` | No |  |
| `direct_debit_payment` | `Record<string, any>` | No |  |
| `do_simulate_settlement_failure` | `boolean` | No |  |
| `due_date` | `string` | No |  |
| `error_description` | `string` | No |  |
| `event` | `any[]` | No |  |
| `failed_payment_request` | `Record<string, any>` | No |  |
| `failure_callback_url` | `string` | No |  |
| `field_display_setting` | `any[]` | No |  |
| `formatted_amount` | `string` | No |  |
| `hosted_pay_checkout_url` | `string` | No |  |
| `id` | `string` | No |  |
| `ignore_address_verification` | `boolean` | No |  |
| `inserted` | `string` | No |  |
| `inserted_sortable` | `string` | No |  |
| `institution` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `jwk` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `lightning_invoice` | `string` | No |  |
| `lightning_invoice_expires_at` | `string` | No |  |
| `merchant_direct_debit_mandate_id` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `notification_email_address` | `string` | No |  |
| `notification_role_i_d` | `any[]` | No |  |
| `order_id` | `string` | No |  |
| `partial_payment_method` | `string` | No |  |
| `partial_payment_step` | `string` | No |  |
| `payment_attempt` | `any[]` | No |  |
| `payment_initiation_id` | `string` | No |  |
| `payment_method` | `any[]` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_request` | `any[]` | No |  |
| `payrun_id` | `string` | No |  |
| `pisp_account_id` | `string` | No |  |
| `priority_bank_id` | `string` | No |  |
| `result` | `Record<string, any>` | No |  |
| `sandbox_settle_delay_in_second` | `number` | No |  |
| `shipping_address` | `Record<string, any>` | No |  |
| `status` | `string` | No |  |
| `success_web_hook_url` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `title` | `string` | No |  |
| `tokenised_card` | `any[]` | No |  |
| `transaction` | `any[]` | No |  |
| `use_hosted_payment_page` | `boolean` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PaymentRequest().create({
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
| `apple_pay_transaction_id` | `string` | No |  |
| `card_authorization_response_id` | `string` | No |  |
| `card_expiry_month` | `number` | No |  |
| `card_expiry_year` | `number` | No |  |
| `card_issuer` | `string` | No |  |
| `card_issuer_country` | `string` | No |  |
| `card_last_four_digit` | `string` | No |  |
| `card_request_id` | `string` | No |  |
| `card_scheme` | `string` | No |  |
| `card_token_customer_id` | `string` | No |  |
| `card_transaction_id` | `string` | No |  |
| `currency` | `string` | No |  |
| `direct_debit_payment_id` | `string` | No |  |
| `direct_debit_payment_reference` | `string` | No |  |
| `drirect_debit_mandate_id` | `string` | No |  |
| `error_message` | `string` | No |  |
| `error_reason` | `string` | No |  |
| `event_type` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `lightning_invoice` | `string` | No |  |
| `lightning_r_hash` | `string` | No |  |
| `origin_url` | `string` | No |  |
| `payment_method_type` | `string` | No |  |
| `payment_processor_name` | `string` | No |  |
| `payment_request_id` | `string` | No |  |
| `pisp_bank_status` | `string` | No |  |
| `pisp_payment_initiation_id` | `string` | No |  |
| `pisp_payment_institution_name` | `string` | No |  |
| `pisp_payment_service_provider_id` | `string` | No |  |
| `pisp_redirect_url` | `string` | No |  |
| `reconciled_transaction_id` | `string` | No |  |
| `refund_payout_id` | `string` | No |  |
| `status` | `string` | No |  |
| `wallet_name` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PaymentRequestEvent().list()
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

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `all` | `number` | No |  |
| `authorized` | `number` | No |  |
| `paid` | `number` | No |  |
| `partially_paid` | `number` | No |  |
| `total_amounts_by_currency` | `Record<string, any>` | No |  |
| `unpaid` | `number` | No |  |

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
| `amount_pending` | `number` | No |  |
| `amount_received` | `number` | No |  |
| `amount_refunded` | `number` | No |  |
| `callback_url` | `string` | No |  |
| `card_stripe_payment_intent_secret` | `string` | No |  |
| `country_code` | `string` | No |  |
| `currency` | `string` | No |  |
| `custom_fields_to_display` | `any[]` | No |  |
| `description` | `string` | No |  |
| `due_date` | `string` | No |  |
| `field_display_setting` | `any[]` | No |  |
| `google_pay_merchant_id` | `string` | No |  |
| `id` | `string` | No |  |
| `jwk` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_logo_url_png` | `string` | No |  |
| `merchant_logo_url_svg` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `merchant_short_name` | `string` | No |  |
| `partial_payment_method` | `string` | No |  |
| `payment_attempt` | `any[]` | No |  |
| `payment_methods_list` | `any[]` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_processor_key` | `string` | No |  |
| `pisp_error` | `string` | No |  |
| `priority_bank_id` | `string` | No |  |
| `status` | `string` | No |  |
| `stripe_account_id` | `string` | No |  |
| `title` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PaymentRequestMinimal().list()
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
| `amount_pending` | `number` | No |  |
| `amount_received` | `number` | No |  |
| `amount_refunded` | `number` | No |  |
| `currency` | `string` | No |  |
| `customer_id` | `string` | No |  |
| `payment` | `any[]` | No |  |
| `payment_request_id` | `string` | No |  |
| `pisp_authorization` | `any[]` | No |  |
| `requested_amount` | `number` | No |  |
| `result` | `string` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PaymentRequestResult().list()
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
| `account_id` | `string` | No |  |
| `allow_incomplete` | `boolean` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `Record<string, any>` | Yes |  |
| `beneficiary_id` | `string` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_process` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `Record<string, any>` | No |  |
| `document` | `any[]` | No |  |
| `event` | `any[]` | No |  |
| `failed_payout` | `Record<string, any>` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `number` | No |  |
| `fx_destination_amount_minor_unit` | `number` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `number` | No |  |
| `fx_use_destination_amount` | `boolean` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_failed` | `boolean` | No |  |
| `is_settled` | `boolean` | No |  |
| `is_submitted` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payout` | `any[]` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `reason` | `string` | No |  |
| `rule` | `Record<string, any>` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `Record<string, any>` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `tag_id` | `any[]` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `number` | No |  |
| `transacted_fx_amount` | `number` | No |  |
| `transacted_fx_rate` | `number` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `account_id` | - | - | Yes | - | - |
| `allow_incomplete` | - | - | - | - | - |
| `amount` | - | - | - | - | - |
| `amount_minor_unit` | - | - | - | - | - |
| `approve_payout_url` | - | - | - | - | - |
| `approver_id` | - | - | - | - | - |
| `authentication_method` | - | - | - | - | - |
| `authorisation` | - | - | - | - | - |
| `authorisers_completed_count` | - | - | - | - | - |
| `authorisers_required_count` | - | - | - | - | - |
| `batch_payout_id` | - | - | - | - | - |
| `beneficiary` | - | - | - | - | - |
| `beneficiary_id` | - | - | - | - | - |
| `can_authorise` | - | - | - | - | - |
| `can_process` | - | - | - | - | - |
| `can_update` | - | - | - | - | - |
| `charge_bearer` | - | - | - | - | - |
| `created_by` | - | - | - | - | - |
| `created_by_email_address` | - | - | - | - | - |
| `currency` | - | - | Yes | - | - |
| `current_user_id` | - | - | - | - | - |
| `description` | - | - | - | - | - |
| `destination` | - | - | - | - | - |
| `document` | - | - | - | - | - |
| `event` | - | - | - | - | - |
| `failed_payout` | - | - | - | - | - |
| `formatted_amount` | - | - | - | - | - |
| `formatted_fx_destination_amount` | - | - | - | - | - |
| `formatted_schedule` | - | - | - | - | - |
| `formatted_schedule_day_only` | - | - | - | - | - |
| `formatted_source_account_available_balance` | - | - | - | - | - |
| `fx_destination_amount` | - | - | - | - | - |
| `fx_destination_amount_minor_unit` | - | - | - | - | - |
| `fx_destination_currency` | - | - | - | - | - |
| `fx_quote_expires_at` | - | - | - | - | - |
| `fx_quote_id` | - | - | - | - | - |
| `fx_rate` | - | - | - | - | - |
| `fx_use_destination_amount` | - | - | - | - | - |
| `has_current_user_authorised` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `inserted` | - | - | - | - | - |
| `invoice_id` | - | - | - | - | - |
| `is_archived` | - | - | - | - | - |
| `is_failed` | - | - | - | - | - |
| `is_settled` | - | - | - | - | - |
| `is_submitted` | - | - | - | - | - |
| `last_updated` | - | - | - | - | - |
| `merchant_id` | - | - | - | - | - |
| `merchant_token_description` | - | - | - | - | - |
| `nonce` | - | - | - | - | - |
| `payment_processor` | - | - | - | - | - |
| `payment_rail` | - | - | - | - | - |
| `payout` | - | - | - | - | - |
| `payrun_id` | - | - | - | - | - |
| `payrun_name` | - | - | - | - | - |
| `reason` | - | - | - | - | - |
| `rule` | - | - | - | - | - |
| `schedule_date` | - | - | - | - | - |
| `scheduled` | - | - | - | - | - |
| `source_account_available_balance` | - | - | - | - | - |
| `source_account_available_balance_minor_unit` | - | - | - | - | - |
| `source_account_bic` | - | - | - | - | - |
| `source_account_currency` | - | - | - | - | - |
| `source_account_iban` | - | - | - | - | - |
| `source_account_identifier` | - | - | - | - | - |
| `source_account_name` | - | - | - | - | - |
| `source_account_number` | - | - | - | - | - |
| `source_account_sortcode` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `tag` | - | - | - | - | - |
| `tag_id` | - | - | - | - | - |
| `their_reference` | - | - | - | - | - |
| `topup_payrun_id` | - | - | - | - | - |
| `transacted_amount` | - | - | - | - | - |
| `transacted_fx_amount` | - | - | - | - | - |
| `transacted_fx_rate` | - | - | - | - | - |
| `type` | - | - | Yes | - | - |
| `user_id` | - | - | - | - | - |
| `your_reference` | - | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Payout().create({
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
| `account_id` | `string` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `Record<string, any>` | Yes |  |
| `can_authorise` | `boolean` | No |  |
| `can_process` | `boolean` | No |  |
| `can_update` | `boolean` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `Record<string, any>` | No |  |
| `document` | `any[]` | No |  |
| `event` | `any[]` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `number` | No |  |
| `fx_destination_amount_minor_unit` | `number` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `number` | No |  |
| `fx_use_destination_amount` | `boolean` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_failed` | `boolean` | No |  |
| `is_settled` | `boolean` | No |  |
| `is_submitted` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `Record<string, any>` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `boolean` | No |  |
| `source_account_available_balance` | `number` | No |  |
| `source_account_available_balance_minor_unit` | `number` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `Record<string, any>` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `any[]` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `number` | No |  |
| `transacted_fx_amount` | `number` | No |  |
| `transacted_fx_rate` | `number` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

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

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `all` | `number` | No |  |
| `failed` | `number` | No |  |
| `in_progress` | `number` | No |  |
| `paid` | `number` | No |  |
| `pending_approval` | `number` | No |  |
| `scheduled` | `number` | No |  |
| `total_amounts_by_currency` | `Record<string, any>` | No |  |

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
| `authorisation` | `any[]` | No |  |
| `authorisation_date` | `string` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `batch_payout_id` | `string` | No |  |
| `can_authorise` | `boolean` | No |  |
| `can_delete` | `boolean` | No |  |
| `can_edit` | `boolean` | No |  |
| `event` | `any[]` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice` | `any[]` | No |  |
| `invoices_minimal` | `any[]` | No |  |
| `is_archived` | `boolean` | No |  |
| `last_updated` | `string` | No |  |
| `last_updated_by` | `Record<string, any>` | Yes |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | No |  |
| `note` | `string` | No |  |
| `payment` | `any[]` | No |  |
| `payout` | `any[]` | No |  |
| `payouts_count` | `number` | No |  |
| `reason` | `string` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled_date` | `string` | No |  |
| `source_account` | `any[]` | No |  |
| `status` | `string` | No |  |
| `total_eur` | `number` | No |  |
| `total_gbp` | `number` | No |  |
| `total_usd` | `number` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Payrun().create({
  id: 'example_id',
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
| `content` | `string` | No |  |
| `content_type` | `string` | No |  |
| `last_completed_at` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `report_name` | `string` | No |  |
| `report_type` | `string` | No |  |
| `statement_number` | `number` | No |  |

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
| `failed_role` | `Record<string, any>` | No |  |
| `role` | `any[]` | No |  |

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
| `account_id` | `string` | No |  |
| `approve_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `any[]` | No |  |
| `authorisation` | `any[]` | No |  |
| `authorisers_completed_count` | `number` | No |  |
| `authorisers_required_count` | `number` | No |  |
| `can_authorise` | `boolean` | No |  |
| `created_by` | `Record<string, any>` | Yes |  |
| `description` | `string` | No |  |
| `end_at` | `string` | No |  |
| `has_current_user_authorised` | `boolean` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_disabled` | `boolean` | No |  |
| `last_executed_at` | `string` | No |  |
| `last_run_at_transaction_date` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | Yes |  |
| `on_approved_web_hook_url` | `string` | No |  |
| `on_execution_error_web_hook_url` | `string` | No |  |
| `on_execution_success_web_hook_url` | `string` | No |  |
| `start_at` | `string` | No |  |
| `status` | `string` | No |  |
| `sweep_action` | `Record<string, any>` | No |  |
| `time_zone_id` | `string` | No |  |
| `trigger_cron_expression` | `string` | No |  |
| `trigger_on_pay_in` | `boolean` | No |  |
| `user_id` | `string` | No |  |
| `web_hook_secret` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `account` | - | - | - | - | - |
| `account_id` | - | - | - | - | - |
| `approve_url` | - | - | - | - | - |
| `approver_id` | - | - | - | - | - |
| `authentication_method` | - | - | - | - | - |
| `authorisation` | - | - | - | - | - |
| `authorisers_completed_count` | - | - | - | - | - |
| `authorisers_required_count` | - | - | - | - | - |
| `can_authorise` | - | - | - | - | - |
| `created_by` | - | - | - | - | - |
| `description` | - | - | - | - | - |
| `end_at` | - | - | - | - | - |
| `has_current_user_authorised` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `inserted` | - | - | - | - | - |
| `is_disabled` | - | - | - | - | - |
| `last_executed_at` | - | - | - | - | - |
| `last_run_at_transaction_date` | - | - | - | - | - |
| `last_updated` | - | - | - | - | - |
| `merchant_id` | - | - | - | - | - |
| `name` | - | - | Yes | - | - |
| `nonce` | - | - | - | - | - |
| `on_approved_web_hook_url` | - | - | - | - | - |
| `on_execution_error_web_hook_url` | - | - | - | - | - |
| `on_execution_success_web_hook_url` | - | - | - | - | - |
| `start_at` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `sweep_action` | - | - | Yes | - | - |
| `time_zone_id` | - | - | - | - | - |
| `trigger_cron_expression` | - | - | - | - | - |
| `trigger_on_pay_in` | - | - | - | - | - |
| `user_id` | - | - | - | - | - |
| `web_hook_secret` | - | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Rule().create({
  created_by: {},
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
| `error_message` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_authorise_to_enable` | `boolean` | No |  |
| `message` | `string` | No |  |
| `raw_response` | `string` | No |  |
| `rule_event_type` | `string` | No |  |
| `rule_id` | `string` | No |  |
| `user` | `Record<string, any>` | Yes |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.RuleEvent().list()
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
| `colour_hex` | `string` | No |  |
| `description` | `string` | No |  |
| `id` | `string` | No |  |
| `merchant_id` | `string` | Yes |  |
| `name` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Tag().create({
  merchant_id: 'example_merchant_id',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Tag().list()
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
| `account_id` | `string` | No |  |
| `account_name` | `string` | No |  |
| `account_sequence_number` | `number` | No |  |
| `address_detail` | `Record<string, any>` | No |  |
| `amount` | `number` | No |  |
| `amount_minor_unit` | `number` | No |  |
| `balance` | `number` | No |  |
| `balance_minor_unit` | `number` | No |  |
| `booking_date_time` | `string` | No |  |
| `charge_detail` | `Record<string, any>` | No |  |
| `content` | `any[]` | No |  |
| `counterparty` | `Record<string, any>` | No |  |
| `counterparty_summary` | `string` | No |  |
| `currency` | `string` | No |  |
| `currency_exchange` | `Record<string, any>` | No |  |
| `date` | `string` | No |  |
| `description` | `string` | No |  |
| `enrichment` | `Record<string, any>` | No |  |
| `fx_amount` | `number` | No |  |
| `fx_currency` | `string` | No |  |
| `fx_rate` | `number` | No |  |
| `gross_amount` | `Record<string, any>` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `iso_bank_transaction_code` | `Record<string, any>` | No |  |
| `merchant` | `Record<string, any>` | No |  |
| `merchant_id` | `string` | No |  |
| `page_number` | `number` | No |  |
| `page_size` | `number` | No |  |
| `payee_detail` | `Record<string, any>` | Yes |  |
| `payer_detail` | `Record<string, any>` | Yes |  |
| `payment_request_custom_field` | `Record<string, any>` | No |  |
| `payment_request_id` | `string` | No |  |
| `payout_id` | `string` | No |  |
| `proprietary_bank_transaction_code` | `Record<string, any>` | No |  |
| `raw_reference` | `string` | No |  |
| `reference` | `string` | No |  |
| `rule_id` | `string` | No |  |
| `statement_reference` | `any[]` | No |  |
| `status` | `string` | No |  |
| `supplementary_data` | `any` | No |  |
| `tag` | `any[]` | No |  |
| `their_reference` | `string` | No |  |
| `total_page` | `number` | No |  |
| `total_size` | `number` | No |  |
| `transaction_amount` | `Record<string, any>` | Yes |  |
| `transaction_date` | `string` | No |  |
| `transaction_information` | `any[]` | No |  |
| `transaction_mutability` | `string` | No |  |
| `type` | `string` | No |  |
| `value_date_time` | `string` | No |  |
| `virtual_iban` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Transaction().create({
  id: 'example_id',
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
| `client_session_timeout` | `any[]` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `boolean` | No |  |
| `permission` | `Record<string, any>` | No |  |
| `profile` | `string` | No |  |
| `roles_with_scope` | `any[]` | No |  |
| `two_factor_enabled` | `boolean` | No |  |
| `user_invite_id` | `string` | No |  |

### Field Usage by Operation

| Field | list | update |
| --- | --- | --- |
| `client_session_timeout` | - | - |
| `email_address` | - | Yes |
| `first_name` | - | Yes |
| `id` | - | - |
| `last_name` | - | Yes |
| `passkey_added` | - | - |
| `permission` | - | - |
| `profile` | - | - |
| `roles_with_scope` | - | - |
| `two_factor_enabled` | - | - |
| `user_invite_id` | - | - |

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
| `authorisation_status` | `Record<string, any>` | No |  |
| `failed_user_invite` | `Record<string, any>` | No |  |
| `id` | `string` | No |  |
| `initial_role_id` | `string` | No |  |
| `invitee_email_address` | `string` | No |  |
| `invitee_first_name` | `string` | No |  |
| `invitee_last_name` | `string` | No |  |
| `inviter_email_address` | `string` | No |  |
| `inviter_first_name` | `string` | No |  |
| `inviter_last_name` | `string` | No |  |
| `is_authorised` | `boolean` | No |  |
| `is_invitee_registered` | `boolean` | No |  |
| `last_invited` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `message` | `string` | No |  |
| `registration_url` | `string` | No |  |
| `send_invite_email` | `boolean` | No |  |
| `status` | `string` | No |  |
| `user` | `Record<string, any>` | Yes |  |
| `user_id` | `string` | No |  |
| `user_invite` | `any[]` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `authorisation_status` | - | - | - | - | - |
| `failed_user_invite` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `initial_role_id` | - | - | - | - | - |
| `invitee_email_address` | - | - | Yes | - | - |
| `invitee_first_name` | - | - | - | - | - |
| `invitee_last_name` | - | - | - | - | - |
| `inviter_email_address` | - | - | - | - | - |
| `inviter_first_name` | - | - | - | - | - |
| `inviter_last_name` | - | - | - | - | - |
| `is_authorised` | - | - | - | - | - |
| `is_invitee_registered` | - | - | - | - | - |
| `last_invited` | - | - | - | - | - |
| `merchant_id` | - | - | - | - | - |
| `merchant_name` | - | - | - | - | - |
| `message` | - | - | - | - | - |
| `registration_url` | - | - | - | - | - |
| `send_invite_email` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `user` | - | - | - | - | - |
| `user_id` | - | - | - | - | - |
| `user_invite` | - | - | - | - | - |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.UserInvite().create({
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.UserInvite().list()
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
| `account_name` | `string` | No |  |
| `account_supplier_name` | `string` | No |  |
| `available_balance` | `number` | No |  |
| `available_balance_minor_unit` | `number` | No |  |
| `balance` | `number` | No |  |
| `balance_minor_unit` | `number` | No |  |
| `bank_name` | `string` | No |  |
| `consent_id` | `string` | No |  |
| `created_by` | `Record<string, any>` | Yes |  |
| `created_by_display_name` | `string` | No |  |
| `currency` | `string` | No |  |
| `default_payment_rail` | `string` | No |  |
| `display_name` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `external_account_icon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `Record<string, any>` | Yes |  |
| `inserted` | `string` | No |  |
| `is_archived` | `boolean` | No |  |
| `is_connected_account` | `boolean` | No |  |
| `is_default` | `boolean` | No |  |
| `is_trust_account` | `boolean` | No |  |
| `is_virtual` | `boolean` | No |  |
| `last_transaction` | `Record<string, any>` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `name` | `string` | Yes |  |
| `physical_account_id` | `string` | No |  |
| `rule` | `any[]` | No |  |
| `submitted_payouts_balance` | `number` | No |  |
| `submitted_payouts_balance_minor_unit` | `number` | No |  |
| `summary` | `string` | No |  |
| `supplier_sepa_instant_status` | `string` | No |  |
| `xero_bank_feed_connection_status` | `string` | No |  |
| `xero_bank_feed_last_synced_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` | No |  |
| `xero_bank_feed_sync_status` | `string` | No |  |
| `xero_unsynchronised_transactions_count` | `number` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Virtual().create({
  account_id: 'example_account_id',
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
| `destination_url` | `string` | No |  |
| `email_address` | `string` | No |  |
| `failed_notification_email_address` | `string` | No |  |
| `id` | `string` | No |  |
| `is_active` | `boolean` | No |  |
| `merchant_id` | `string` | No |  |
| `notification_method` | `string` | No |  |
| `resource_type` | `any[]` | No |  |
| `retry` | `boolean` | No |  |
| `secret` | `string` | No |  |
| `version` | `number` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `destination_url` | - | - | - | - | - |
| `email_address` | - | - | - | - | - |
| `failed_notification_email_address` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `is_active` | - | - | - | - | - |
| `merchant_id` | - | - | Yes | Yes | - |
| `notification_method` | - | - | Yes | Yes | - |
| `resource_type` | - | - | - | - | - |
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
const results = await client.Webhook().list()
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

