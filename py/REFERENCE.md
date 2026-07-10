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

#### `MerchantDirectDebitMandate(data=None)`

Create a new `MerchantDirectDebitMandateEntity` instance. Pass `None` for no initial data.

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

#### `PayoutKeyset(data=None)`

Create a new `PayoutKeysetEntity` instance. Pass `None` for no initial data.

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
| `account_balance` | `list` | No |  |
| `account_id` | `str` | No |  |
| `account_identification` | `list` | No |  |
| `account_name` | `str` | No |  |
| `account_supplier_name` | `str` | No |  |
| `account_type` | `str` | No |  |
| `available_balance` | `float` | No |  |
| `available_balance_minor_unit` | `int` | No |  |
| `balance` | `float` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `bank_name` | `str` | No |  |
| `consent_id` | `str` | No |  |
| `consolidated_account_information` | `dict` | No |  |
| `created_by` | `dict` | Yes |  |
| `created_by_display_name` | `str` | No |  |
| `currency` | `str` | No |  |
| `default_payment_rail` | `str` | No |  |
| `description` | `str` | No |  |
| `detail` | `str` | No |  |
| `display_name` | `str` | No |  |
| `expiry_date` | `str` | No |  |
| `external_account_icon` | `str` | No |  |
| `format` | `str` | No |  |
| `from_date` | `str` | No |  |
| `id` | `str` | No |  |
| `identifier` | `dict` | Yes |  |
| `inserted` | `str` | No |  |
| `is_archived` | `bool` | No |  |
| `is_connected_account` | `bool` | No |  |
| `is_default` | `bool` | No |  |
| `is_trust_account` | `bool` | No |  |
| `is_virtual` | `bool` | No |  |
| `last_transaction` | `dict` | No |  |
| `last_updated` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `merchant_name` | `str` | No |  |
| `nickname` | `str` | No |  |
| `physical_account_id` | `str` | No |  |
| `role_i_d` | `list` | No |  |
| `rule` | `list` | No |  |
| `submitted_payouts_balance` | `float` | No |  |
| `submitted_payouts_balance_minor_unit` | `int` | No |  |
| `summary` | `str` | No |  |
| `supplier_physical_account_id` | `str` | No |  |
| `supplier_sepa_instant_status` | `str` | No |  |
| `to_date` | `str` | No |  |
| `type` | `str` | No |  |
| `usage_type` | `str` | No |  |
| `xero_bank_feed_connection_status` | `str` | No |  |
| `xero_bank_feed_last_synced_at` | `str` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `str` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `str` | No |  |
| `xero_bank_feed_sync_status` | `str` | No |  |
| `xero_unsynchronised_transactions_count` | `int` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Account().create({
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
| `approve_url` | `str` | No |  |
| `id` | `str` | No |  |
| `payout` | `list` | No |  |

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
| `approval_callback_url` | `str` | No |  |
| `authentication_method` | `list` | No |  |
| `authorisation` | `list` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `beneficiary` | `list` | No |  |
| `beneficiary_event` | `list` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `created_by` | `dict` | Yes |  |
| `created_by_email_address` | `str` | No |  |
| `currency` | `str` | Yes |  |
| `destination` | `dict` | No |  |
| `failed_beneficiary` | `dict` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `is_enabled` | `bool` | No |  |
| `last_authorised` | `str` | No |  |
| `last_updated` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `name` | `str` | Yes |  |
| `nonce` | `str` | No |  |
| `source_account` | `list` | No |  |
| `source_account_i_d` | `list` | No |  |
| `their_reference` | `str` | No |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Beneficiary().create({
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
| `group_member` | `list` | No |  |
| `group_name` | `str` | Yes |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `last_updated` | `str` | No |  |
| `merchant_id` | `str` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.BeneficiaryGroup().list()
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
| `authorized_amount` | `str` | No |  |
| `currency_code` | `str` | No |  |
| `is_payer_authentication_required` | `bool` | No |  |
| `is_soft_decline` | `bool` | No |  |
| `payer_authentication_access_token` | `str` | No |  |
| `payer_authentication_merchant_data` | `str` | No |  |
| `payer_authentication_url` | `str` | No |  |
| `payer_authentication_window_height` | `int` | No |  |
| `payer_authentication_window_width` | `int` | No |  |
| `payment_request_callback_url` | `str` | No |  |
| `payment_request_id` | `str` | No |  |
| `request_id` | `str` | No |  |
| `response_code` | `str` | No |  |
| `response_type` | `str` | No |  |
| `status` | `str` | No |  |
| `three_ds_redirect_url` | `str` | No |  |
| `transaction_id` | `str` | No |  |

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
| `card_type` | `str` | No |  |
| `customer_email_address` | `str` | No |  |
| `expiry_month` | `str` | No |  |
| `expiry_year` | `str` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `last_four_digit` | `str` | No |  |
| `last_updated` | `str` | No |  |
| `masked_card_number` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `payment_request_id` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.CardCustomerToken().list()
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
| `authorized_amount` | `str` | No |  |
| `currency_code` | `str` | No |  |
| `is_payer_authentication_required` | `bool` | No |  |
| `is_soft_decline` | `bool` | No |  |
| `payer_authentication_access_token` | `str` | No |  |
| `payer_authentication_merchant_data` | `str` | No |  |
| `payer_authentication_url` | `str` | No |  |
| `payer_authentication_window_height` | `int` | No |  |
| `payer_authentication_window_width` | `int` | No |  |
| `payment_request_callback_url` | `str` | No |  |
| `payment_request_id` | `str` | No |  |
| `request_id` | `str` | No |  |
| `response_code` | `str` | No |  |
| `response_type` | `str` | No |  |
| `status` | `str` | No |  |
| `three_ds_redirect_url` | `str` | No |  |
| `transaction_id` | `str` | No |  |

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
| `authorisation_url` | `str` | No |  |
| `callback_url` | `str` | No |  |
| `consent_id` | `str` | No |  |
| `email_address` | `str` | No |  |
| `expiry_date` | `str` | No |  |
| `failure_callback_url` | `str` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `institution_id` | `str` | No |  |
| `is_connected_account` | `bool` | No |  |
| `is_enabled` | `bool` | No |  |
| `merchant_id` | `str` | No |  |
| `provider` | `str` | No |  |
| `success_web_hook_url` | `str` | No |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Consent().create({
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Consent().list()
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
| `decimal` | `int` | No |  |
| `is_fiat` | `bool` | No |  |
| `iso4217_alpha_code` | `str` | No |  |
| `iso4217_numeric_code` | `str` | No |  |
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
| `failed_submission` | `dict` | No |  |
| `successful_submission` | `list` | No |  |

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
| `destination_currency` | `str` | No |  |
| `exchange_rate` | `float` | No |  |
| `expiry_time` | `str` | No |  |
| `quote_id` | `str` | No |  |
| `source_currency` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.FxRate().list()
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
| `payment_request_id` | `str` | No |  |
| `response_type` | `str` | No |  |

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
| `account_number` | `str` | No |  |
| `address_line1` | `str` | Yes |  |
| `address_line2` | `str` | No |  |
| `approved_at` | `str` | No |  |
| `city` | `str` | Yes |  |
| `country_code` | `str` | Yes |  |
| `currency` | `str` | No |  |
| `customer_account_number` | `str` | No |  |
| `customer_city` | `str` | No |  |
| `customer_country_code` | `str` | No |  |
| `customer_country_name` | `str` | No |  |
| `customer_email_address` | `str` | No |  |
| `customer_first_name` | `str` | No |  |
| `customer_iban` | `str` | No |  |
| `customer_last_name` | `str` | No |  |
| `customer_sort_code` | `str` | No |  |
| `email_address` | `str` | Yes |  |
| `first_name` | `str` | Yes |  |
| `iban` | `str` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `is_recurring` | `bool` | No |  |
| `last_name` | `str` | Yes |  |
| `last_updated` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `postal_code` | `str` | Yes |  |
| `reference` | `str` | No |  |
| `sort_code` | `str` | No |  |
| `status` | `str` | No |  |
| `supplier_bank_account_id` | `str` | No |  |
| `supplier_customer_id` | `str` | No |  |
| `supplier_mandate_id` | `str` | No |  |
| `supplier_name` | `str` | No |  |
| `supplier_status` | `str` | No |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Mandate().create({
    "address_line1": "example_address_line1",  # str
    "city": "example_city",  # str
    "country_code": "example_country_code",  # str
    "email_address": "example_email_address",  # str
    "first_name": "example_first_name",  # str
    "last_name": "example_last_name",  # str
    "postal_code": "example_postal_code",  # str
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
| `account_currency` | `list` | No |  |
| `can_have_trust_account` | `bool` | No |  |
| `card_payment_processor` | `str` | No |  |
| `company_id` | `str` | No |  |
| `display_qr_on_hosted_pay` | `bool` | No |  |
| `hosted_pay_version` | `int` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `is_blocked` | `bool` | No |  |
| `is_exited` | `bool` | No |  |
| `is_suspended` | `bool` | No |  |
| `jurisdiction` | `str` | No |  |
| `logo_url_png` | `str` | No |  |
| `logo_url_svg` | `str` | No |  |
| `merchant_category_code` | `str` | No |  |
| `name` | `str` | No |  |
| `note` | `str` | No |  |
| `parent_merchant` | `dict` | No |  |
| `payment_account` | `list` | No |  |
| `payment_account_limit` | `int` | No |  |
| `reason` | `str` | No |  |
| `short_name` | `str` | No |  |
| `supported_payment_methods_list` | `list` | No |  |
| `suspension_reason` | `str` | No |  |
| `tag` | `list` | No |  |
| `time_zone_id` | `str` | No |  |
| `trading_name` | `str` | No |  |
| `web_hook_limit` | `int` | No |  |
| `your_role_name` | `str` | No |  |

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
| `amount_lower` | `float` | No |  |
| `amount_upper` | `float` | No |  |
| `authorisation_type` | `str` | No |  |
| `beneficiaries_only` | `bool` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `last_editor_cant_authorise` | `bool` | No |  |
| `last_updated` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `number_of_authoriser` | `int` | No |  |
| `role_setting` | `list` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.MerchantAuthorisationSetting().list()
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

## MerchantDirectDebitMandateEntity

```python
merchant_direct_debit_mandate = client.MerchantDirectDebitMandate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approved_at` | `str` | No |  |
| `currency` | `str` | No |  |
| `customer_account_number` | `str` | No |  |
| `customer_city` | `str` | No |  |
| `customer_country_code` | `str` | No |  |
| `customer_country_name` | `str` | No |  |
| `customer_email_address` | `str` | No |  |
| `customer_first_name` | `str` | No |  |
| `customer_iban` | `str` | No |  |
| `customer_last_name` | `str` | No |  |
| `customer_sort_code` | `str` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `is_recurring` | `bool` | No |  |
| `last_updated` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `reference` | `str` | No |  |
| `status` | `str` | No |  |
| `supplier_bank_account_id` | `str` | No |  |
| `supplier_customer_id` | `str` | No |  |
| `supplier_mandate_id` | `str` | No |  |
| `supplier_name` | `str` | No |  |
| `supplier_status` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.MerchantDirectDebitMandate().list()
for merchant_direct_debit_mandate in results:
    print(merchant_direct_debit_mandate)
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

Create a new `MerchantDirectDebitMandateEntity` instance with the same options.

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
| `bank_country_code` | `list` | No |  |
| `bank_id` | `str` | No |  |
| `bank_name` | `str` | No |  |
| `business_institution_id` | `str` | No |  |
| `currency` | `str` | No |  |
| `logo` | `str` | No |  |
| `message` | `str` | No |  |
| `message_image_url` | `str` | No |  |
| `order` | `int` | No |  |
| `personal_institution_id` | `str` | No |  |
| `processor` | `str` | No |  |
| `warning_heading` | `str` | No |  |
| `warning_message` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.MerchantPayByBankSetting().list()
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
| `description` | `str` | Yes |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `last_updated` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `name` | `str` | Yes |  |
| `template` | `dict` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.MerchantPaymentRequestTemplate().list()
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
| `authentication_method` | `list` | No |  |
| `authorisation` | `list` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `can_authorise` | `bool` | No |  |
| `description` | `str` | No |  |
| `expires_at` | `str` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `hmac_algorithm` | `str` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `ip_address_whitelist` | `str` | No |  |
| `is_archived` | `bool` | No |  |
| `is_enabled` | `bool` | No |  |
| `last_authorised` | `str` | No |  |
| `last_updated` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `nonce` | `str` | Yes |  |
| `permission_type` | `list` | No |  |
| `request_signature_version` | `int` | No |  |
| `shared_secret_algorithm` | `str` | No |  |
| `shared_secret_base64` | `str` | No |  |
| `token` | `str` | No |  |

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
results = client.MerchantToken().list()
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
| `build_version` | `int` | No |  |
| `major_version` | `int` | No |  |
| `minor_version` | `int` | No |  |
| `release_name` | `str` | No |  |

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
| `account_name` | `str` | Yes |  |
| `account_number` | `str` | No |  |
| `iban` | `str` | Yes |  |
| `payee_verified_account_name` | `str` | No |  |
| `result` | `str` | No |  |
| `secondary_identification` | `str` | No |  |
| `sort_code` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Payeeverification().create({
    "account_name": "example_account_name",  # str
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
| `address` | `list` | No |  |
| `amount` | `float` | No |  |
| `amount_pending` | `float` | No |  |
| `amount_received` | `float` | No |  |
| `amount_refunded` | `float` | No |  |
| `auto_send_receipt` | `bool` | No |  |
| `base_origin_url` | `str` | No |  |
| `callback_url` | `str` | No |  |
| `card_authorize_only` | `bool` | No |  |
| `card_create_token` | `bool` | No |  |
| `card_create_token_mode` | `str` | No |  |
| `card_ignore_cvn` | `bool` | No |  |
| `card_no_payer_authentication` | `bool` | No |  |
| `card_processor_merchant_id` | `str` | No |  |
| `card_stripe_payment_intent_id` | `str` | No |  |
| `card_stripe_payment_intent_secret` | `str` | No |  |
| `card_transmit_raw_detail` | `bool` | No |  |
| `created_by_user` | `dict` | Yes |  |
| `currency` | `str` | No |  |
| `custom_field` | `list` | No |  |
| `customer_email_address` | `str` | No |  |
| `customer_id` | `str` | No |  |
| `customer_name` | `str` | No |  |
| `description` | `str` | No |  |
| `destination_account` | `dict` | No |  |
| `direct_debit_payment` | `dict` | No |  |
| `due_date` | `str` | No |  |
| `event` | `list` | No |  |
| `failure_callback_url` | `str` | No |  |
| `field_display_setting` | `list` | No |  |
| `formatted_amount` | `str` | No |  |
| `hosted_pay_checkout_url` | `str` | No |  |
| `id` | `str` | No |  |
| `ignore_address_verification` | `bool` | No |  |
| `inserted` | `str` | No |  |
| `inserted_sortable` | `str` | No |  |
| `is_archived` | `bool` | No |  |
| `jwk` | `str` | No |  |
| `last_updated` | `str` | No |  |
| `lightning_invoice` | `str` | No |  |
| `lightning_invoice_expires_at` | `str` | No |  |
| `merchant_direct_debit_mandate_id` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `merchant_token_description` | `str` | No |  |
| `notification_email_address` | `str` | No |  |
| `notification_role_i_d` | `list` | No |  |
| `order_id` | `str` | No |  |
| `partial_payment_method` | `str` | No |  |
| `partial_payment_step` | `str` | No |  |
| `payment_attempt` | `list` | No |  |
| `payment_method` | `list` | No |  |
| `payment_processor` | `str` | No |  |
| `payrun_id` | `str` | No |  |
| `pisp_account_id` | `str` | No |  |
| `priority_bank_id` | `str` | No |  |
| `result` | `dict` | No |  |
| `sandbox_settle_delay_in_second` | `int` | No |  |
| `shipping_address` | `dict` | No |  |
| `shipping_address_city` | `str` | No |  |
| `shipping_address_country_code` | `str` | No |  |
| `shipping_address_county` | `str` | No |  |
| `shipping_address_line1` | `str` | No |  |
| `shipping_address_line2` | `str` | No |  |
| `shipping_address_post_code` | `str` | No |  |
| `shipping_email` | `str` | No |  |
| `shipping_first_name` | `str` | No |  |
| `shipping_last_name` | `str` | No |  |
| `shipping_phone` | `str` | No |  |
| `status` | `str` | No |  |
| `success_web_hook_url` | `str` | No |  |
| `tag` | `list` | No |  |
| `tag_id` | `list` | No |  |
| `title` | `str` | No |  |
| `tokenised_card` | `list` | No |  |
| `transaction` | `list` | No |  |
| `use_hosted_payment_page` | `bool` | No |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Payment().create({
    "created_by_user": {},  # dict
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
| `account_name` | `str` | No |  |
| `account_supplier_name` | `str` | No |  |
| `available_balance` | `float` | No |  |
| `available_balance_minor_unit` | `int` | No |  |
| `balance` | `float` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `bank_name` | `str` | No |  |
| `consent_id` | `str` | No |  |
| `created_by` | `dict` | Yes |  |
| `created_by_display_name` | `str` | No |  |
| `currency` | `str` | No |  |
| `default_payment_rail` | `str` | No |  |
| `display_name` | `str` | No |  |
| `expiry_date` | `str` | No |  |
| `external_account_icon` | `str` | No |  |
| `id` | `str` | No |  |
| `identifier` | `dict` | Yes |  |
| `inserted` | `str` | No |  |
| `is_archived` | `bool` | No |  |
| `is_connected_account` | `bool` | No |  |
| `is_default` | `bool` | No |  |
| `is_trust_account` | `bool` | No |  |
| `is_virtual` | `bool` | No |  |
| `last_transaction` | `dict` | No |  |
| `last_updated` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `merchant_name` | `str` | No |  |
| `physical_account_id` | `str` | No |  |
| `rule` | `list` | No |  |
| `submitted_payouts_balance` | `float` | No |  |
| `submitted_payouts_balance_minor_unit` | `int` | No |  |
| `summary` | `str` | No |  |
| `supplier_sepa_instant_status` | `str` | No |  |
| `xero_bank_feed_connection_status` | `str` | No |  |
| `xero_bank_feed_last_synced_at` | `str` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `str` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `str` | No |  |
| `xero_bank_feed_sync_status` | `str` | No |  |
| `xero_unsynchronised_transactions_count` | `int` | No |  |

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
| `account_name` | `str` | No |  |
| `available_balance` | `float` | No |  |
| `balance` | `float` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `currency` | `str` | No |  |
| `id` | `str` | No |  |
| `identifier` | `dict` | Yes |  |
| `is_archived` | `bool` | No |  |
| `is_connected_account` | `bool` | No |  |
| `merchant_id` | `str` | No |  |
| `submitted_payouts_balance` | `float` | No |  |

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
| `payment_initiation_id` | `str` | No |  |
| `payment_request_callback_url` | `str` | No |  |
| `payment_request_id` | `str` | No |  |
| `redirect_url` | `str` | No |  |
| `response_type` | `str` | No |  |
| `specific_error_message` | `str` | No |  |

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
| `address` | `list` | No |  |
| `amount` | `float` | No |  |
| `amount_pending` | `float` | No |  |
| `amount_received` | `float` | No |  |
| `amount_refunded` | `float` | No |  |
| `auto_send_receipt` | `bool` | No |  |
| `base_origin_url` | `str` | No |  |
| `callback_url` | `str` | No |  |
| `card_authorize_only` | `bool` | No |  |
| `card_create_token` | `bool` | No |  |
| `card_create_token_mode` | `str` | No |  |
| `card_ignore_cvn` | `bool` | No |  |
| `card_processor_merchant_id` | `str` | No |  |
| `card_stripe_payment_intent_id` | `str` | No |  |
| `card_stripe_payment_intent_secret` | `str` | No |  |
| `created_by_user` | `dict` | Yes |  |
| `currency` | `str` | No |  |
| `custom_field` | `list` | No |  |
| `customer_email_address` | `str` | No |  |
| `customer_id` | `str` | No |  |
| `customer_name` | `str` | No |  |
| `description` | `str` | No |  |
| `destination_account` | `dict` | No |  |
| `direct_debit_payment` | `dict` | No |  |
| `do_simulate_settlement_failure` | `bool` | No |  |
| `due_date` | `str` | No |  |
| `error_description` | `str` | No |  |
| `event` | `list` | No |  |
| `failed_payment_request` | `dict` | No |  |
| `failure_callback_url` | `str` | No |  |
| `field_display_setting` | `list` | No |  |
| `formatted_amount` | `str` | No |  |
| `hosted_pay_checkout_url` | `str` | No |  |
| `id` | `str` | No |  |
| `ignore_address_verification` | `bool` | No |  |
| `inserted` | `str` | No |  |
| `inserted_sortable` | `str` | No |  |
| `institution` | `str` | No |  |
| `is_archived` | `bool` | No |  |
| `jwk` | `str` | No |  |
| `last_updated` | `str` | No |  |
| `lightning_invoice` | `str` | No |  |
| `lightning_invoice_expires_at` | `str` | No |  |
| `merchant_direct_debit_mandate_id` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `merchant_token_description` | `str` | No |  |
| `notification_email_address` | `str` | No |  |
| `notification_role_i_d` | `list` | No |  |
| `order_id` | `str` | No |  |
| `partial_payment_method` | `str` | No |  |
| `partial_payment_step` | `str` | No |  |
| `payment_attempt` | `list` | No |  |
| `payment_initiation_id` | `str` | No |  |
| `payment_method` | `list` | No |  |
| `payment_processor` | `str` | No |  |
| `payment_request` | `list` | No |  |
| `payrun_id` | `str` | No |  |
| `pisp_account_id` | `str` | No |  |
| `priority_bank_id` | `str` | No |  |
| `result` | `dict` | No |  |
| `sandbox_settle_delay_in_second` | `int` | No |  |
| `shipping_address` | `dict` | No |  |
| `status` | `str` | No |  |
| `success_web_hook_url` | `str` | No |  |
| `tag` | `list` | No |  |
| `title` | `str` | No |  |
| `tokenised_card` | `list` | No |  |
| `transaction` | `list` | No |  |
| `use_hosted_payment_page` | `bool` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PaymentRequest().create({
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
| `apple_pay_transaction_id` | `str` | No |  |
| `card_authorization_response_id` | `str` | No |  |
| `card_expiry_month` | `int` | No |  |
| `card_expiry_year` | `int` | No |  |
| `card_issuer` | `str` | No |  |
| `card_issuer_country` | `str` | No |  |
| `card_last_four_digit` | `str` | No |  |
| `card_request_id` | `str` | No |  |
| `card_scheme` | `str` | No |  |
| `card_token_customer_id` | `str` | No |  |
| `card_transaction_id` | `str` | No |  |
| `currency` | `str` | No |  |
| `direct_debit_payment_id` | `str` | No |  |
| `direct_debit_payment_reference` | `str` | No |  |
| `drirect_debit_mandate_id` | `str` | No |  |
| `error_message` | `str` | No |  |
| `error_reason` | `str` | No |  |
| `event_type` | `str` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `lightning_invoice` | `str` | No |  |
| `lightning_r_hash` | `str` | No |  |
| `origin_url` | `str` | No |  |
| `payment_method_type` | `str` | No |  |
| `payment_processor_name` | `str` | No |  |
| `payment_request_id` | `str` | No |  |
| `pisp_bank_status` | `str` | No |  |
| `pisp_payment_initiation_id` | `str` | No |  |
| `pisp_payment_institution_name` | `str` | No |  |
| `pisp_payment_service_provider_id` | `str` | No |  |
| `pisp_redirect_url` | `str` | No |  |
| `reconciled_transaction_id` | `str` | No |  |
| `refund_payout_id` | `str` | No |  |
| `status` | `str` | No |  |
| `wallet_name` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PaymentRequestEvent().list()
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

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `all` | `int` | No |  |
| `authorized` | `int` | No |  |
| `paid` | `int` | No |  |
| `partially_paid` | `int` | No |  |
| `total_amounts_by_currency` | `dict` | No |  |
| `unpaid` | `int` | No |  |

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
| `amount_pending` | `float` | No |  |
| `amount_received` | `float` | No |  |
| `amount_refunded` | `float` | No |  |
| `callback_url` | `str` | No |  |
| `card_stripe_payment_intent_secret` | `str` | No |  |
| `country_code` | `str` | No |  |
| `currency` | `str` | No |  |
| `custom_fields_to_display` | `list` | No |  |
| `description` | `str` | No |  |
| `due_date` | `str` | No |  |
| `field_display_setting` | `list` | No |  |
| `google_pay_merchant_id` | `str` | No |  |
| `id` | `str` | No |  |
| `jwk` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `merchant_logo_url_png` | `str` | No |  |
| `merchant_logo_url_svg` | `str` | No |  |
| `merchant_name` | `str` | No |  |
| `merchant_short_name` | `str` | No |  |
| `partial_payment_method` | `str` | No |  |
| `payment_attempt` | `list` | No |  |
| `payment_methods_list` | `list` | No |  |
| `payment_processor` | `str` | No |  |
| `payment_processor_key` | `str` | No |  |
| `pisp_error` | `str` | No |  |
| `priority_bank_id` | `str` | No |  |
| `status` | `str` | No |  |
| `stripe_account_id` | `str` | No |  |
| `title` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PaymentRequestMinimal().list()
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
| `amount_pending` | `float` | No |  |
| `amount_received` | `float` | No |  |
| `amount_refunded` | `float` | No |  |
| `currency` | `str` | No |  |
| `customer_id` | `str` | No |  |
| `payment` | `list` | No |  |
| `payment_request_id` | `str` | No |  |
| `pisp_authorization` | `list` | No |  |
| `requested_amount` | `float` | No |  |
| `result` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PaymentRequestResult().list()
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
| `account_id` | `str` | No |  |
| `allow_incomplete` | `bool` | No |  |
| `amount` | `float` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `str` | No |  |
| `approver_id` | `str` | No |  |
| `authentication_method` | `list` | No |  |
| `authorisation` | `list` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `str` | No |  |
| `beneficiary` | `dict` | Yes |  |
| `beneficiary_id` | `str` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `str` | No |  |
| `created_by` | `str` | No |  |
| `created_by_email_address` | `str` | No |  |
| `currency` | `str` | No |  |
| `current_user_id` | `str` | No |  |
| `description` | `str` | No |  |
| `destination` | `dict` | No |  |
| `document` | `list` | No |  |
| `event` | `list` | No |  |
| `failed_payout` | `dict` | No |  |
| `formatted_amount` | `str` | No |  |
| `formatted_fx_destination_amount` | `str` | No |  |
| `formatted_schedule` | `str` | No |  |
| `formatted_schedule_day_only` | `str` | No |  |
| `formatted_source_account_available_balance` | `str` | No |  |
| `fx_destination_amount` | `float` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `str` | No |  |
| `fx_quote_expires_at` | `str` | No |  |
| `fx_quote_id` | `str` | No |  |
| `fx_rate` | `float` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `invoice_id` | `str` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `merchant_token_description` | `str` | No |  |
| `nonce` | `str` | No |  |
| `payment_processor` | `str` | No |  |
| `payment_rail` | `str` | No |  |
| `payout` | `list` | No |  |
| `payrun_id` | `str` | No |  |
| `payrun_name` | `str` | No |  |
| `reason` | `str` | No |  |
| `rule` | `dict` | No |  |
| `schedule_date` | `str` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `str` | No |  |
| `source_account_currency` | `str` | No |  |
| `source_account_iban` | `str` | No |  |
| `source_account_identifier` | `dict` | Yes |  |
| `source_account_name` | `str` | No |  |
| `source_account_number` | `str` | No |  |
| `source_account_sortcode` | `str` | No |  |
| `status` | `str` | No |  |
| `tag` | `list` | No |  |
| `tag_id` | `list` | No |  |
| `their_reference` | `str` | No |  |
| `topup_payrun_id` | `str` | No |  |
| `transacted_amount` | `float` | No |  |
| `transacted_fx_amount` | `float` | No |  |
| `transacted_fx_rate` | `float` | No |  |
| `type` | `str` | No |  |
| `user_id` | `str` | No |  |
| `your_reference` | `str` | No |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Payout().create({
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

## PayoutKeysetEntity

```python
payout_keyset = client.PayoutKeyset()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `str` | No |  |
| `amount` | `float` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `str` | No |  |
| `approver_id` | `str` | No |  |
| `authentication_method` | `list` | No |  |
| `authorisation` | `list` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `str` | No |  |
| `beneficiary` | `dict` | Yes |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `str` | No |  |
| `created_by` | `str` | No |  |
| `created_by_email_address` | `str` | No |  |
| `currency` | `str` | No |  |
| `current_user_id` | `str` | No |  |
| `description` | `str` | No |  |
| `destination` | `dict` | No |  |
| `document` | `list` | No |  |
| `event` | `list` | No |  |
| `formatted_amount` | `str` | No |  |
| `formatted_fx_destination_amount` | `str` | No |  |
| `formatted_schedule` | `str` | No |  |
| `formatted_schedule_day_only` | `str` | No |  |
| `formatted_source_account_available_balance` | `str` | No |  |
| `fx_destination_amount` | `float` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `str` | No |  |
| `fx_quote_expires_at` | `str` | No |  |
| `fx_quote_id` | `str` | No |  |
| `fx_rate` | `float` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `invoice_id` | `str` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `merchant_token_description` | `str` | No |  |
| `nonce` | `str` | No |  |
| `payment_processor` | `str` | No |  |
| `payment_rail` | `str` | No |  |
| `payrun_id` | `str` | No |  |
| `payrun_name` | `str` | No |  |
| `rule` | `dict` | No |  |
| `schedule_date` | `str` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `str` | No |  |
| `source_account_currency` | `str` | No |  |
| `source_account_iban` | `str` | No |  |
| `source_account_identifier` | `dict` | Yes |  |
| `source_account_name` | `str` | No |  |
| `source_account_number` | `str` | No |  |
| `source_account_sortcode` | `str` | No |  |
| `status` | `str` | No |  |
| `tag` | `list` | No |  |
| `their_reference` | `str` | No |  |
| `topup_payrun_id` | `str` | No |  |
| `transacted_amount` | `float` | No |  |
| `transacted_fx_amount` | `float` | No |  |
| `transacted_fx_rate` | `float` | No |  |
| `type` | `str` | No |  |
| `user_id` | `str` | No |  |
| `your_reference` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PayoutKeyset().list()
for payout_keyset in results:
    print(payout_keyset)
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

Create a new `PayoutKeysetEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PayoutMetricEntity

```python
payout_metric = client.PayoutMetric()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `all` | `float` | No |  |
| `failed` | `float` | No |  |
| `in_progress` | `float` | No |  |
| `paid` | `float` | No |  |
| `pending_approval` | `float` | No |  |
| `scheduled` | `float` | No |  |
| `total_amounts_by_currency` | `dict` | No |  |

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
| `authorisation` | `list` | No |  |
| `authorisation_date` | `str` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `str` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_delete` | `bool` | No |  |
| `can_edit` | `bool` | No |  |
| `event` | `list` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `invoice` | `list` | No |  |
| `invoices_minimal` | `list` | No |  |
| `is_archived` | `bool` | No |  |
| `last_updated` | `str` | No |  |
| `last_updated_by` | `dict` | Yes |  |
| `merchant_id` | `str` | No |  |
| `name` | `str` | No |  |
| `nonce` | `str` | No |  |
| `note` | `str` | No |  |
| `payment` | `list` | No |  |
| `payout` | `list` | No |  |
| `payouts_count` | `int` | No |  |
| `reason` | `str` | No |  |
| `schedule_date` | `str` | No |  |
| `scheduled_date` | `str` | No |  |
| `source_account` | `list` | No |  |
| `status` | `str` | No |  |
| `total_eur` | `float` | No |  |
| `total_gbp` | `float` | No |  |
| `total_usd` | `float` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Payrun().create({
    "id": "example_id",  # str
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
| `content` | `str` | No |  |
| `content_type` | `str` | No |  |
| `last_completed_at` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `report_name` | `str` | No |  |
| `report_type` | `str` | No |  |
| `statement_number` | `int` | No |  |

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
| `failed_role` | `dict` | No |  |
| `role` | `list` | No |  |

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
| `account_id` | `str` | No |  |
| `approve_url` | `str` | No |  |
| `approver_id` | `str` | No |  |
| `authentication_method` | `list` | No |  |
| `authorisation` | `list` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `can_authorise` | `bool` | No |  |
| `created_by` | `dict` | Yes |  |
| `description` | `str` | No |  |
| `end_at` | `str` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `is_disabled` | `bool` | No |  |
| `last_executed_at` | `str` | No |  |
| `last_run_at_transaction_date` | `str` | No |  |
| `last_updated` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `name` | `str` | No |  |
| `nonce` | `str` | Yes |  |
| `on_approved_web_hook_url` | `str` | No |  |
| `on_execution_error_web_hook_url` | `str` | No |  |
| `on_execution_success_web_hook_url` | `str` | No |  |
| `start_at` | `str` | No |  |
| `status` | `str` | No |  |
| `sweep_action` | `dict` | No |  |
| `time_zone_id` | `str` | No |  |
| `trigger_cron_expression` | `str` | No |  |
| `trigger_on_pay_in` | `bool` | No |  |
| `user_id` | `str` | No |  |
| `web_hook_secret` | `str` | No |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Rule().create({
    "created_by": {},  # dict
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
| `error_message` | `str` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `is_authorise_to_enable` | `bool` | No |  |
| `message` | `str` | No |  |
| `raw_response` | `str` | No |  |
| `rule_event_type` | `str` | No |  |
| `rule_id` | `str` | No |  |
| `user` | `dict` | Yes |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.RuleEvent().list()
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
| `colour_hex` | `str` | No |  |
| `description` | `str` | No |  |
| `id` | `str` | No |  |
| `merchant_id` | `str` | Yes |  |
| `name` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Tag().create({
    "merchant_id": "example_merchant_id",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Tag().list()
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
| `account_id` | `str` | No |  |
| `account_name` | `str` | No |  |
| `account_sequence_number` | `int` | No |  |
| `address_detail` | `dict` | No |  |
| `amount` | `float` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `balance` | `float` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `booking_date_time` | `str` | No |  |
| `charge_detail` | `dict` | No |  |
| `content` | `list` | No |  |
| `counterparty` | `dict` | No |  |
| `counterparty_summary` | `str` | No |  |
| `currency` | `str` | No |  |
| `currency_exchange` | `dict` | No |  |
| `date` | `str` | No |  |
| `description` | `str` | No |  |
| `enrichment` | `dict` | No |  |
| `fx_amount` | `float` | No |  |
| `fx_currency` | `str` | No |  |
| `fx_rate` | `float` | No |  |
| `gross_amount` | `dict` | Yes |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `iso_bank_transaction_code` | `dict` | No |  |
| `merchant` | `dict` | No |  |
| `merchant_id` | `str` | No |  |
| `page_number` | `int` | No |  |
| `page_size` | `int` | No |  |
| `payee_detail` | `dict` | Yes |  |
| `payer_detail` | `dict` | Yes |  |
| `payment_request_custom_field` | `dict` | No |  |
| `payment_request_id` | `str` | No |  |
| `payout_id` | `str` | No |  |
| `proprietary_bank_transaction_code` | `dict` | No |  |
| `raw_reference` | `str` | No |  |
| `reference` | `str` | No |  |
| `rule_id` | `str` | No |  |
| `statement_reference` | `list` | No |  |
| `status` | `str` | No |  |
| `supplementary_data` | `Any` | No |  |
| `tag` | `list` | No |  |
| `their_reference` | `str` | No |  |
| `total_page` | `int` | No |  |
| `total_size` | `int` | No |  |
| `transaction_amount` | `dict` | Yes |  |
| `transaction_date` | `str` | No |  |
| `transaction_information` | `list` | No |  |
| `transaction_mutability` | `str` | No |  |
| `type` | `str` | No |  |
| `value_date_time` | `str` | No |  |
| `virtual_iban` | `str` | No |  |
| `your_reference` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Transaction().create({
    "id": "example_id",  # str
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
| `client_session_timeout` | `list` | No |  |
| `email_address` | `str` | Yes |  |
| `first_name` | `str` | Yes |  |
| `id` | `str` | No |  |
| `last_name` | `str` | Yes |  |
| `passkey_added` | `bool` | No |  |
| `permission` | `dict` | No |  |
| `profile` | `str` | No |  |
| `roles_with_scope` | `list` | No |  |
| `two_factor_enabled` | `bool` | No |  |
| `user_invite_id` | `str` | No |  |

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
| `authorisation_status` | `dict` | No |  |
| `failed_user_invite` | `dict` | No |  |
| `id` | `str` | No |  |
| `initial_role_id` | `str` | No |  |
| `invitee_email_address` | `str` | No |  |
| `invitee_first_name` | `str` | No |  |
| `invitee_last_name` | `str` | No |  |
| `inviter_email_address` | `str` | No |  |
| `inviter_first_name` | `str` | No |  |
| `inviter_last_name` | `str` | No |  |
| `is_authorised` | `bool` | No |  |
| `is_invitee_registered` | `bool` | No |  |
| `last_invited` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `merchant_name` | `str` | No |  |
| `message` | `str` | No |  |
| `registration_url` | `str` | No |  |
| `send_invite_email` | `bool` | No |  |
| `status` | `str` | No |  |
| `user` | `dict` | Yes |  |
| `user_id` | `str` | No |  |
| `user_invite` | `list` | No |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.UserInvite().create({
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.UserInvite().list()
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
| `account_name` | `str` | No |  |
| `account_supplier_name` | `str` | No |  |
| `available_balance` | `float` | No |  |
| `available_balance_minor_unit` | `int` | No |  |
| `balance` | `float` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `bank_name` | `str` | No |  |
| `consent_id` | `str` | No |  |
| `created_by` | `dict` | Yes |  |
| `created_by_display_name` | `str` | No |  |
| `currency` | `str` | No |  |
| `default_payment_rail` | `str` | No |  |
| `display_name` | `str` | No |  |
| `expiry_date` | `str` | No |  |
| `external_account_icon` | `str` | No |  |
| `id` | `str` | No |  |
| `identifier` | `dict` | Yes |  |
| `inserted` | `str` | No |  |
| `is_archived` | `bool` | No |  |
| `is_connected_account` | `bool` | No |  |
| `is_default` | `bool` | No |  |
| `is_trust_account` | `bool` | No |  |
| `is_virtual` | `bool` | No |  |
| `last_transaction` | `dict` | No |  |
| `last_updated` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `merchant_name` | `str` | No |  |
| `name` | `str` | Yes |  |
| `physical_account_id` | `str` | No |  |
| `rule` | `list` | No |  |
| `submitted_payouts_balance` | `float` | No |  |
| `submitted_payouts_balance_minor_unit` | `int` | No |  |
| `summary` | `str` | No |  |
| `supplier_sepa_instant_status` | `str` | No |  |
| `xero_bank_feed_connection_status` | `str` | No |  |
| `xero_bank_feed_last_synced_at` | `str` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `str` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `str` | No |  |
| `xero_bank_feed_sync_status` | `str` | No |  |
| `xero_unsynchronised_transactions_count` | `int` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Virtual().create({
    "account_id": "example_account_id",  # str
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
| `destination_url` | `str` | No |  |
| `email_address` | `str` | No |  |
| `failed_notification_email_address` | `str` | No |  |
| `id` | `str` | No |  |
| `is_active` | `bool` | No |  |
| `merchant_id` | `str` | No |  |
| `notification_method` | `str` | No |  |
| `resource_type` | `list` | No |  |
| `retry` | `bool` | No |  |
| `secret` | `str` | No |  |
| `version` | `int` | No |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Webhook().create({
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Webhook().list()
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

