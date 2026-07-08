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

#### `Beneficiary(data=None)`

Create a new `BeneficiaryEntity` instance. Pass `None` for no initial data.

#### `Cancel(data=None)`

Create a new `CancelEntity` instance. Pass `None` for no initial data.

#### `Disable(data=None)`

Create a new `DisableEntity` instance. Pass `None` for no initial data.

#### `Enable(data=None)`

Create a new `EnableEntity` instance. Pass `None` for no initial data.

#### `Merchant(data=None)`

Create a new `MerchantEntity` instance. Pass `None` for no initial data.

#### `Metadata(data=None)`

Create a new `MetadataEntity` instance. Pass `None` for no initial data.

#### `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(data=None)`

Create a new `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionBizBizModelsPagingPaymentRequestPage(data=None)`

Create a new `NoFrixionBizBizModelsPagingPaymentRequestPageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionBizBizModelsPagingPayoutPage(data=None)`

Create a new `NoFrixionBizBizModelsPagingPayoutPageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionBizBizModelsPagingPayrunPage(data=None)`

Create a new `NoFrixionBizBizModelsPagingPayrunPageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionBizBizModelsPagingRuleEventsPage(data=None)`

Create a new `NoFrixionBizBizModelsPagingRuleEventsPageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionBizBizModelsPagingRulesPage(data=None)`

Create a new `NoFrixionBizBizModelsPagingRulesPageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionBizBizModelsPaymentsCardPayment(data=None)`

Create a new `NoFrixionBizBizModelsPaymentsCardPaymentEntity` instance. Pass `None` for no initial data.

#### `NoFrixionBizBizModelsPaymentsCardPublicKey(data=None)`

Create a new `NoFrixionBizBizModelsPaymentsCardPublicKeyEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(data=None)`

Create a new `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(data=None)`

Create a new `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(data=None)`

Create a new `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(data=None)`

Create a new `NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(data=None)`

Create a new `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsBatchPayout(data=None)`

Create a new `NoFrixionMoneyMoovModelsBatchPayoutEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsBeneficiaryGroupPage(data=None)`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsBeneficiaryPage(data=None)`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryPageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsCardCustomerToken(data=None)`

Create a new `NoFrixionMoneyMoovModelsCardCustomerTokenEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(data=None)`

Create a new `NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(data=None)`

Create a new `NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsFxRate(data=None)`

Create a new `NoFrixionMoneyMoovModelsFxRateEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsIPayment(data=None)`

Create a new `NoFrixionMoneyMoovModelsIPaymentEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsMandatesMandate(data=None)`

Create a new `NoFrixionMoneyMoovModelsMandatesMandateEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchant(data=None)`

Create a new `NoFrixionMoneyMoovModelsMerchantEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantPage(data=None)`

Create a new `NoFrixionMoneyMoovModelsMerchantPageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantPayByBankSetting(data=None)`

Create a new `NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantToken(data=None)`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantTokenPage(data=None)`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenPageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsNoFrixionVersion(data=None)`

Create a new `NoFrixionMoneyMoovModelsNoFrixionVersionEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsOpenBankingAccount(data=None)`

Create a new `NoFrixionMoneyMoovModelsOpenBankingAccountEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsOpenBankingConsent(data=None)`

Create a new `NoFrixionMoneyMoovModelsOpenBankingConsentEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsOpenBankingTransaction(data=None)`

Create a new `NoFrixionMoneyMoovModelsOpenBankingTransactionEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPayment(data=None)`

Create a new `NoFrixionMoneyMoovModelsPaymentEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(data=None)`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentAccountPage(data=None)`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountPageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentInitiation(data=None)`

Create a new `NoFrixionMoneyMoovModelsPaymentInitiationEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestEvent(data=None)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestEventEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestMetric(data=None)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMetricEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestMinimal(data=None)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestResult(data=None)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestResultEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(data=None)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(data=None)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(data=None)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(data=None)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPayoutKeysetPage(data=None)`

Create a new `NoFrixionMoneyMoovModelsPayoutKeysetPageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPayoutMetric(data=None)`

Create a new `NoFrixionMoneyMoovModelsPayoutMetricEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(data=None)`

Create a new `NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsPayrun(data=None)`

Create a new `NoFrixionMoneyMoovModelsPayrunEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsReportResult(data=None)`

Create a new `NoFrixionMoneyMoovModelsReportResultEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsRule(data=None)`

Create a new `NoFrixionMoneyMoovModelsRuleEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsTransaction(data=None)`

Create a new `NoFrixionMoneyMoovModelsTransactionEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsTransactionPage(data=None)`

Create a new `NoFrixionMoneyMoovModelsTransactionPageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsUserInvite(data=None)`

Create a new `NoFrixionMoneyMoovModelsUserInviteEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsUserInvitePage(data=None)`

Create a new `NoFrixionMoneyMoovModelsUserInvitePageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsUserPage(data=None)`

Create a new `NoFrixionMoneyMoovModelsUserPageEntity` instance. Pass `None` for no initial data.

#### `NoFrixionMoneyMoovModelsWebhook(data=None)`

Create a new `NoFrixionMoneyMoovModelsWebhookEntity` instance. Pass `None` for no initial data.

#### `OpenBanking(data=None)`

Create a new `OpenBankingEntity` instance. Pass `None` for no initial data.

#### `Payeeverification(data=None)`

Create a new `PayeeverificationEntity` instance. Pass `None` for no initial data.

#### `PaymentRequest(data=None)`

Create a new `PaymentRequestEntity` instance. Pass `None` for no initial data.

#### `Payout(data=None)`

Create a new `PayoutEntity` instance. Pass `None` for no initial data.

#### `Payrun(data=None)`

Create a new `PayrunEntity` instance. Pass `None` for no initial data.

#### `Reject(data=None)`

Create a new `RejectEntity` instance. Pass `None` for no initial data.

#### `Report(data=None)`

Create a new `ReportEntity` instance. Pass `None` for no initial data.

#### `Rule(data=None)`

Create a new `RuleEntity` instance. Pass `None` for no initial data.

#### `Send(data=None)`

Create a new `SendEntity` instance. Pass `None` for no initial data.

#### `Sendbeneficiary(data=None)`

Create a new `SendbeneficiaryEntity` instance. Pass `None` for no initial data.

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

#### `Whoami(data=None)`

Create a new `WhoamiEntity` instance. Pass `None` for no initial data.

#### `Whoamitrustedapp(data=None)`

Create a new `WhoamitrustedappEntity` instance. Pass `None` for no initial data.

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
| `account_id` | `str` | No |  |
| `account_name` | `str` | No |  |
| `account_supplier_name` | `str` | No |  |
| `account_type` | `str` | No |  |
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
| `physical_account_id` | `str` | No |  |
| `role_i_d` | `list` | No |  |
| `rule` | `list` | No |  |
| `submitted_payouts_balance` | `float` | No |  |
| `submitted_payouts_balance_minor_unit` | `int` | No |  |
| `summary` | `str` | No |  |
| `supplier_physical_account_id` | `str` | No |  |
| `supplier_sepa_instant_status` | `str` | No |  |
| `to_date` | `str` | No |  |
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
| `beneficiary_event` | `list` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `created_by` | `dict` | Yes |  |
| `created_by_email_address` | `str` | No |  |
| `currency` | `str` | Yes |  |
| `destination` | `dict` | No |  |
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

| Field | load | create | update | remove |
| --- | --- | --- | --- | --- |
| `approval_callback_url` | - | - | - | - |
| `authentication_method` | - | - | - | - |
| `authorisation` | - | - | - | - |
| `authorisers_completed_count` | - | - | - | - |
| `authorisers_required_count` | - | - | - | - |
| `beneficiary_event` | - | - | - | - |
| `can_authorise` | - | - | - | - |
| `can_update` | - | - | - | - |
| `created_by` | - | - | - | - |
| `created_by_email_address` | - | - | - | - |
| `currency` | - | - | Yes | - |
| `destination` | - | Yes | - | - |
| `has_current_user_authorised` | - | - | - | - |
| `id` | - | - | - | - |
| `inserted` | - | - | - | - |
| `is_enabled` | - | - | - | - |
| `last_authorised` | - | - | - | - |
| `last_updated` | - | - | - | - |
| `merchant_id` | - | Yes | - | - |
| `name` | - | - | Yes | - |
| `nonce` | - | - | - | - |
| `source_account` | - | - | - | - |
| `source_account_i_d` | - | - | - | - |
| `their_reference` | - | - | - | - |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Beneficiary().create({
})
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

## CancelEntity

```python
cancel = client.Cancel()
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

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Cancel().update({
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

Create a new `CancelEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## DisableEntity

```python
disable = client.Disable()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `str` | No |  |
| `authentication_method` | `list` | No |  |
| `authorisation` | `list` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `beneficiary_event` | `list` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `created_by` | `dict` | Yes |  |
| `created_by_email_address` | `str` | No |  |
| `currency` | `str` | Yes |  |
| `destination` | `dict` | No |  |
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
| `their_reference` | `str` | No |  |

### Operations

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Disable().update({
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

Create a new `DisableEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## EnableEntity

```python
enable = client.Enable()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `str` | No |  |
| `authentication_method` | `list` | No |  |
| `authorisation` | `list` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `beneficiary_event` | `list` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `created_by` | `dict` | Yes |  |
| `created_by_email_address` | `str` | No |  |
| `currency` | `str` | Yes |  |
| `destination` | `dict` | No |  |
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
| `their_reference` | `str` | No |  |

### Operations

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Enable().update({
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

Create a new `EnableEntity` instance with the same options.

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
| `reason` | `str` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Merchant().load({"merchant_id": "merchant_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Merchant().remove()
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Merchant().update({
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

Create a new `MerchantEntity` instance with the same options.

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

## NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity

```python
no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page = client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage()
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
results = client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage().list()
for no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page in results:
    print(no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page)
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

Create a new `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionBizBizModelsPagingPaymentRequestPageEntity

```python
no_frixion_biz_biz_models_paging_payment_request_page = client.NoFrixionBizBizModelsPagingPaymentRequestPage()
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
| `status` | `str` | No |  |
| `success_web_hook_url` | `str` | No |  |
| `tag` | `list` | No |  |
| `title` | `str` | No |  |
| `tokenised_card` | `list` | No |  |
| `transaction` | `list` | No |  |
| `use_hosted_payment_page` | `bool` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.NoFrixionBizBizModelsPagingPaymentRequestPage().list()
for no_frixion_biz_biz_models_paging_payment_request_page in results:
    print(no_frixion_biz_biz_models_paging_payment_request_page)
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

Create a new `NoFrixionBizBizModelsPagingPaymentRequestPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionBizBizModelsPagingPayoutPageEntity

```python
no_frixion_biz_biz_models_paging_payout_page = client.NoFrixionBizBizModelsPagingPayoutPage()
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
results = client.NoFrixionBizBizModelsPagingPayoutPage().list()
for no_frixion_biz_biz_models_paging_payout_page in results:
    print(no_frixion_biz_biz_models_paging_payout_page)
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

Create a new `NoFrixionBizBizModelsPagingPayoutPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionBizBizModelsPagingPayrunPageEntity

```python
no_frixion_biz_biz_models_paging_payrun_page = client.NoFrixionBizBizModelsPagingPayrunPage()
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
| `payment` | `list` | No |  |
| `payout` | `list` | No |  |
| `payouts_count` | `int` | No |  |
| `schedule_date` | `str` | No |  |
| `source_account` | `list` | No |  |
| `status` | `str` | No |  |
| `total_eur` | `float` | No |  |
| `total_gbp` | `float` | No |  |
| `total_usd` | `float` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.NoFrixionBizBizModelsPagingPayrunPage().list()
for no_frixion_biz_biz_models_paging_payrun_page in results:
    print(no_frixion_biz_biz_models_paging_payrun_page)
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

Create a new `NoFrixionBizBizModelsPagingPayrunPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionBizBizModelsPagingRuleEventsPageEntity

```python
no_frixion_biz_biz_models_paging_rule_events_page = client.NoFrixionBizBizModelsPagingRuleEventsPage()
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
results = client.NoFrixionBizBizModelsPagingRuleEventsPage().list()
for no_frixion_biz_biz_models_paging_rule_events_page in results:
    print(no_frixion_biz_biz_models_paging_rule_events_page)
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

Create a new `NoFrixionBizBizModelsPagingRuleEventsPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionBizBizModelsPagingRulesPageEntity

```python
no_frixion_biz_biz_models_paging_rules_page = client.NoFrixionBizBizModelsPagingRulesPage()
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

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.NoFrixionBizBizModelsPagingRulesPage().list()
for no_frixion_biz_biz_models_paging_rules_page in results:
    print(no_frixion_biz_biz_models_paging_rules_page)
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

Create a new `NoFrixionBizBizModelsPagingRulesPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionBizBizModelsPaymentsCardPaymentEntity

```python
no_frixion_biz_biz_models_payments_card_payment = client.NoFrixionBizBizModelsPaymentsCardPayment()
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
result = client.NoFrixionBizBizModelsPaymentsCardPayment().create({
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

Create a new `NoFrixionBizBizModelsPaymentsCardPaymentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionBizBizModelsPaymentsCardPublicKeyEntity

```python
no_frixion_biz_biz_models_payments_card_public_key = client.NoFrixionBizBizModelsPaymentsCardPublicKey()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jwt` | `str` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionBizBizModelsPaymentsCardPublicKey().load({"paymentrequest_id": "paymentrequest_id"})
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

Create a new `NoFrixionBizBizModelsPaymentsCardPublicKeyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity

```python
no_frixion_money_moov_api_features_beneficiaries_beneficiaries = client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `beneficiary` | `list` | No |  |
| `failed_beneficiary` | `dict` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries().create({
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

Create a new `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity

```python
no_frixion_money_moov_api_features_payment_requests_payment = client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_payment_request` | `dict` | No |  |
| `payment_request` | `list` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment().create({
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

Create a new `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity

```python
no_frixion_money_moov_api_features_permissions_roles_create = client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate()
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
result = client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate().create({
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

Create a new `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity

```python
no_frixion_money_moov_api_features_user_invites_create = client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_user_invite` | `dict` | No |  |
| `user_invite` | `list` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate().create({
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

Create a new `NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity

```python
no_frixion_money_moov_models_authorisation_settings_merchant = client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant()
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
results = client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant().list()
for no_frixion_money_moov_models_authorisation_settings_merchant in results:
    print(no_frixion_money_moov_models_authorisation_settings_merchant)
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

Create a new `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsBatchPayoutEntity

```python
no_frixion_money_moov_models_batch_payout = client.NoFrixionMoneyMoovModelsBatchPayout()
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
result = client.NoFrixionMoneyMoovModelsBatchPayout().create({
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsBatchPayout().load({"id": "no_frixion_money_moov_models_batch_payout_id"})
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

Create a new `NoFrixionMoneyMoovModelsBatchPayoutEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity

```python
no_frixion_money_moov_models_beneficiary_group_page = client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage()
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
results = client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage().list()
for no_frixion_money_moov_models_beneficiary_group_page in results:
    print(no_frixion_money_moov_models_beneficiary_group_page)
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

Create a new `NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsBeneficiaryPageEntity

```python
no_frixion_money_moov_models_beneficiary_page = client.NoFrixionMoneyMoovModelsBeneficiaryPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `str` | No |  |
| `authentication_method` | `list` | No |  |
| `authorisation` | `list` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `beneficiary_event` | `list` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `created_by` | `dict` | Yes |  |
| `created_by_email_address` | `str` | No |  |
| `currency` | `str` | Yes |  |
| `destination` | `dict` | No |  |
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
| `their_reference` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.NoFrixionMoneyMoovModelsBeneficiaryPage().list()
for no_frixion_money_moov_models_beneficiary_page in results:
    print(no_frixion_money_moov_models_beneficiary_page)
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

Create a new `NoFrixionMoneyMoovModelsBeneficiaryPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsCardCustomerTokenEntity

```python
no_frixion_money_moov_models_card_customer_token = client.NoFrixionMoneyMoovModelsCardCustomerToken()
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
results = client.NoFrixionMoneyMoovModelsCardCustomerToken().list()
for no_frixion_money_moov_models_card_customer_token in results:
    print(no_frixion_money_moov_models_card_customer_token)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsCardCustomerToken().load({"customer_email_address": "customer_email_address"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.NoFrixionMoneyMoovModelsCardCustomerToken().remove()
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

Create a new `NoFrixionMoneyMoovModelsCardCustomerTokenEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity

```python
no_frixion_money_moov_models_currency_currency_info = client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo()
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
results = client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo().list()
for no_frixion_money_moov_models_currency_currency_info in results:
    print(no_frixion_money_moov_models_currency_currency_info)
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

Create a new `NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity

```python
no_frixion_money_moov_models_direct_debit_batch_submit = client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit()
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
result = client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit().create({
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

Create a new `NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsFxRateEntity

```python
no_frixion_money_moov_models_fx_rate = client.NoFrixionMoneyMoovModelsFxRate()
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
results = client.NoFrixionMoneyMoovModelsFxRate().list()
for no_frixion_money_moov_models_fx_rate in results:
    print(no_frixion_money_moov_models_fx_rate)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsFxRate().load({"destination": "destination", "source": "source", "valid_for_minute": 1})
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

Create a new `NoFrixionMoneyMoovModelsFxRateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsIPaymentEntity

```python
no_frixion_money_moov_models_i_payment = client.NoFrixionMoneyMoovModelsIPayment()
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
result = client.NoFrixionMoneyMoovModelsIPayment().create({
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

Create a new `NoFrixionMoneyMoovModelsIPaymentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMandatesMandateEntity

```python
no_frixion_money_moov_models_mandates_mandate = client.NoFrixionMoneyMoovModelsMandatesMandate()
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
result = client.NoFrixionMoneyMoovModelsMandatesMandate().create({
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
result = client.NoFrixionMoneyMoovModelsMandatesMandate().load({"id": "no_frixion_money_moov_models_mandates_mandate_id"})
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

Create a new `NoFrixionMoneyMoovModelsMandatesMandateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantEntity

```python
no_frixion_money_moov_models_merchant = client.NoFrixionMoneyMoovModelsMerchant()
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
results = client.NoFrixionMoneyMoovModelsMerchant().list()
for no_frixion_money_moov_models_merchant in results:
    print(no_frixion_money_moov_models_merchant)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsMerchant().load({"id": "no_frixion_money_moov_models_merchant_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsMerchant().update({
    "id": "no_frixion_money_moov_models_merchant_id",
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

Create a new `NoFrixionMoneyMoovModelsMerchantEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantPageEntity

```python
no_frixion_money_moov_models_merchant_page = client.NoFrixionMoneyMoovModelsMerchantPage()
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
results = client.NoFrixionMoneyMoovModelsMerchantPage().list()
for no_frixion_money_moov_models_merchant_page in results:
    print(no_frixion_money_moov_models_merchant_page)
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

Create a new `NoFrixionMoneyMoovModelsMerchantPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity

```python
no_frixion_money_moov_models_merchant_pay_by_bank_setting = client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting()
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
results = client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting().list()
for no_frixion_money_moov_models_merchant_pay_by_bank_setting in results:
    print(no_frixion_money_moov_models_merchant_pay_by_bank_setting)
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

Create a new `NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantTokenEntity

```python
no_frixion_money_moov_models_merchant_token = client.NoFrixionMoneyMoovModelsMerchantToken()
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

| Field | load | create | update |
| --- | --- | --- | --- |
| `authentication_method` | - | - | - |
| `authorisation` | - | - | - |
| `authorisers_completed_count` | - | - | - |
| `authorisers_required_count` | - | - | - |
| `can_authorise` | - | - | - |
| `description` | - | Yes | - |
| `expires_at` | - | - | - |
| `has_current_user_authorised` | - | - | - |
| `hmac_algorithm` | - | - | - |
| `id` | - | - | - |
| `inserted` | - | - | - |
| `ip_address_whitelist` | - | - | - |
| `is_archived` | - | - | - |
| `is_enabled` | - | - | - |
| `last_authorised` | - | - | - |
| `last_updated` | - | - | - |
| `merchant_id` | - | Yes | - |
| `nonce` | - | - | - |
| `permission_type` | - | - | - |
| `request_signature_version` | - | - | - |
| `shared_secret_algorithm` | - | - | - |
| `shared_secret_base64` | - | - | - |
| `token` | - | - | - |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsMerchantToken().create({
    "nonce": "example_nonce",  # str
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsMerchantToken().load({"id": "no_frixion_money_moov_models_merchant_token_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsMerchantToken().update({
    "id": "no_frixion_money_moov_models_merchant_token_id",
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

Create a new `NoFrixionMoneyMoovModelsMerchantTokenEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantTokenPageEntity

```python
no_frixion_money_moov_models_merchant_token_page = client.NoFrixionMoneyMoovModelsMerchantTokenPage()
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

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.NoFrixionMoneyMoovModelsMerchantTokenPage().list()
for no_frixion_money_moov_models_merchant_token_page in results:
    print(no_frixion_money_moov_models_merchant_token_page)
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

Create a new `NoFrixionMoneyMoovModelsMerchantTokenPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsNoFrixionVersionEntity

```python
no_frixion_money_moov_models_no_frixion_version = client.NoFrixionMoneyMoovModelsNoFrixionVersion()
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
result = client.NoFrixionMoneyMoovModelsNoFrixionVersion().load()
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

Create a new `NoFrixionMoneyMoovModelsNoFrixionVersionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsOpenBankingAccountEntity

```python
no_frixion_money_moov_models_open_banking_account = client.NoFrixionMoneyMoovModelsOpenBankingAccount()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_balance` | `list` | No |  |
| `account_identification` | `list` | No |  |
| `account_name` | `list` | No |  |
| `account_type` | `str` | No |  |
| `balance` | `float` | No |  |
| `consolidated_account_information` | `dict` | No |  |
| `currency` | `str` | No |  |
| `description` | `str` | No |  |
| `detail` | `str` | No |  |
| `id` | `str` | No |  |
| `nickname` | `str` | No |  |
| `type` | `str` | No |  |
| `usage_type` | `str` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsOpenBankingAccount().load({"id": "no_frixion_money_moov_models_open_banking_account_id"})
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

Create a new `NoFrixionMoneyMoovModelsOpenBankingAccountEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsOpenBankingConsentEntity

```python
no_frixion_money_moov_models_open_banking_consent = client.NoFrixionMoneyMoovModelsOpenBankingConsent()
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
result = client.NoFrixionMoneyMoovModelsOpenBankingConsent().create({
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.NoFrixionMoneyMoovModelsOpenBankingConsent().list()
for no_frixion_money_moov_models_open_banking_consent in results:
    print(no_frixion_money_moov_models_open_banking_consent)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsOpenBankingConsent().load({"id": "no_frixion_money_moov_models_open_banking_consent_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.NoFrixionMoneyMoovModelsOpenBankingConsent().remove({"id": "no_frixion_money_moov_models_open_banking_consent_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsOpenBankingConsent().update({
    "id": "no_frixion_money_moov_models_open_banking_consent_id",
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

Create a new `NoFrixionMoneyMoovModelsOpenBankingConsentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsOpenBankingTransactionEntity

```python
no_frixion_money_moov_models_open_banking_transaction = client.NoFrixionMoneyMoovModelsOpenBankingTransaction()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address_detail` | `dict` | No |  |
| `amount` | `float` | No |  |
| `balance` | `dict` | No |  |
| `booking_date_time` | `str` | No |  |
| `charge_detail` | `dict` | No |  |
| `currency` | `str` | No |  |
| `currency_exchange` | `dict` | No |  |
| `date` | `str` | No |  |
| `description` | `str` | No |  |
| `enrichment` | `dict` | No |  |
| `gross_amount` | `dict` | Yes |  |
| `id` | `str` | No |  |
| `iso_bank_transaction_code` | `dict` | No |  |
| `merchant` | `dict` | No |  |
| `payee_detail` | `dict` | Yes |  |
| `payer_detail` | `dict` | Yes |  |
| `proprietary_bank_transaction_code` | `dict` | No |  |
| `reference` | `str` | No |  |
| `statement_reference` | `list` | No |  |
| `status` | `str` | No |  |
| `supplementary_data` | `Any` | No |  |
| `transaction_amount` | `dict` | Yes |  |
| `transaction_information` | `list` | No |  |
| `transaction_mutability` | `str` | No |  |
| `value_date_time` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.NoFrixionMoneyMoovModelsOpenBankingTransaction().list()
for no_frixion_money_moov_models_open_banking_transaction in results:
    print(no_frixion_money_moov_models_open_banking_transaction)
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

Create a new `NoFrixionMoneyMoovModelsOpenBankingTransactionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentEntity

```python
no_frixion_money_moov_models_payment = client.NoFrixionMoneyMoovModelsPayment()
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
result = client.NoFrixionMoneyMoovModelsPayment().create({
    "created_by_user": {},  # dict
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsPayment().load({"id": "no_frixion_money_moov_models_payment_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsPayment().update({
    "id": "no_frixion_money_moov_models_payment_id",
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

Create a new `NoFrixionMoneyMoovModelsPaymentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity

```python
no_frixion_money_moov_models_payment_account_minimal_page = client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage()
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
results = client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage().list()
for no_frixion_money_moov_models_payment_account_minimal_page in results:
    print(no_frixion_money_moov_models_payment_account_minimal_page)
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

Create a new `NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentAccountPageEntity

```python
no_frixion_money_moov_models_payment_account_page = client.NoFrixionMoneyMoovModelsPaymentAccountPage()
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
results = client.NoFrixionMoneyMoovModelsPaymentAccountPage().list()
for no_frixion_money_moov_models_payment_account_page in results:
    print(no_frixion_money_moov_models_payment_account_page)
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

Create a new `NoFrixionMoneyMoovModelsPaymentAccountPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentInitiationEntity

```python
no_frixion_money_moov_models_payment_initiation = client.NoFrixionMoneyMoovModelsPaymentInitiation()
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
result = client.NoFrixionMoneyMoovModelsPaymentInitiation().create({
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

Create a new `NoFrixionMoneyMoovModelsPaymentInitiationEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestEventEntity

```python
no_frixion_money_moov_models_payment_request_event = client.NoFrixionMoneyMoovModelsPaymentRequestEvent()
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
results = client.NoFrixionMoneyMoovModelsPaymentRequestEvent().list()
for no_frixion_money_moov_models_payment_request_event in results:
    print(no_frixion_money_moov_models_payment_request_event)
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

Create a new `NoFrixionMoneyMoovModelsPaymentRequestEventEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestMetricEntity

```python
no_frixion_money_moov_models_payment_request_metric = client.NoFrixionMoneyMoovModelsPaymentRequestMetric()
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
result = client.NoFrixionMoneyMoovModelsPaymentRequestMetric().load()
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

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMetricEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity

```python
no_frixion_money_moov_models_payment_request_minimal = client.NoFrixionMoneyMoovModelsPaymentRequestMinimal()
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
results = client.NoFrixionMoneyMoovModelsPaymentRequestMinimal().list()
for no_frixion_money_moov_models_payment_request_minimal in results:
    print(no_frixion_money_moov_models_payment_request_minimal)
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

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestResultEntity

```python
no_frixion_money_moov_models_payment_request_result = client.NoFrixionMoneyMoovModelsPaymentRequestResult()
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
results = client.NoFrixionMoneyMoovModelsPaymentRequestResult().list()
for no_frixion_money_moov_models_payment_request_result in results:
    print(no_frixion_money_moov_models_payment_request_result)
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

Create a new `NoFrixionMoneyMoovModelsPaymentRequestResultEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity

```python
no_frixion_money_moov_models_payment_requests_merchant_payment = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment()
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
results = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment().list()
for no_frixion_money_moov_models_payment_requests_merchant_payment in results:
    print(no_frixion_money_moov_models_payment_requests_merchant_payment)
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

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity

```python
no_frixion_money_moov_models_payment_requests_merchant_payment2 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2()
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

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2().load({"paymentrequest_id": "paymentrequest_id", "template_id": "template_id"})
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

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity

```python
no_frixion_money_moov_models_payment_requests_merchant_payment3 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3()
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

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3().update({
    "paymentrequest_id": "paymentrequest_id",
    "template_id": "template_id",
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

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity

```python
no_frixion_money_moov_models_payment_requests_merchant_payment4 = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4()
```

### Operations

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4().remove({"paymentrequest_id": "paymentrequest_id", "template_id": "template_id"})
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

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayoutKeysetPageEntity

```python
no_frixion_money_moov_models_payout_keyset_page = client.NoFrixionMoneyMoovModelsPayoutKeysetPage()
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
results = client.NoFrixionMoneyMoovModelsPayoutKeysetPage().list()
for no_frixion_money_moov_models_payout_keyset_page in results:
    print(no_frixion_money_moov_models_payout_keyset_page)
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

Create a new `NoFrixionMoneyMoovModelsPayoutKeysetPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayoutMetricEntity

```python
no_frixion_money_moov_models_payout_metric = client.NoFrixionMoneyMoovModelsPayoutMetric()
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
result = client.NoFrixionMoneyMoovModelsPayoutMetric().load()
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

Create a new `NoFrixionMoneyMoovModelsPayoutMetricEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity

```python
no_frixion_money_moov_models_payouts_payouts_create = client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_payout` | `dict` | No |  |
| `payout` | `list` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate().create({
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

Create a new `NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayrunEntity

```python
no_frixion_money_moov_models_payrun = client.NoFrixionMoneyMoovModelsPayrun()
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
result = client.NoFrixionMoneyMoovModelsPayrun().create({
    "id": "example_id",  # str
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsPayrun().load({"id": "no_frixion_money_moov_models_payrun_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsPayrun().update({
    "id": "no_frixion_money_moov_models_payrun_id",
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

Create a new `NoFrixionMoneyMoovModelsPayrunEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsReportResultEntity

```python
no_frixion_money_moov_models_report_result = client.NoFrixionMoneyMoovModelsReportResult()
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
result = client.NoFrixionMoneyMoovModelsReportResult().load({"id": 1, "report_id": "report_id"})
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

Create a new `NoFrixionMoneyMoovModelsReportResultEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsRuleEntity

```python
no_frixion_money_moov_models_rule = client.NoFrixionMoneyMoovModelsRule()
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

| Field | load | create | update |
| --- | --- | --- | --- |
| `account` | - | - | - |
| `account_id` | - | - | - |
| `approve_url` | - | - | - |
| `approver_id` | - | - | - |
| `authentication_method` | - | - | - |
| `authorisation` | - | - | - |
| `authorisers_completed_count` | - | - | - |
| `authorisers_required_count` | - | - | - |
| `can_authorise` | - | - | - |
| `created_by` | - | - | - |
| `description` | - | - | - |
| `end_at` | - | - | - |
| `has_current_user_authorised` | - | - | - |
| `id` | - | - | - |
| `inserted` | - | - | - |
| `is_disabled` | - | - | - |
| `last_executed_at` | - | - | - |
| `last_run_at_transaction_date` | - | - | - |
| `last_updated` | - | - | - |
| `merchant_id` | - | - | - |
| `name` | - | Yes | - |
| `nonce` | - | - | - |
| `on_approved_web_hook_url` | - | - | - |
| `on_execution_error_web_hook_url` | - | - | - |
| `on_execution_success_web_hook_url` | - | - | - |
| `start_at` | - | - | - |
| `status` | - | - | - |
| `sweep_action` | - | Yes | - |
| `time_zone_id` | - | - | - |
| `trigger_cron_expression` | - | - | - |
| `trigger_on_pay_in` | - | - | - |
| `user_id` | - | - | - |
| `web_hook_secret` | - | - | - |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsRule().create({
    "created_by": {},  # dict
    "nonce": "example_nonce",  # str
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsRule().load({"id": "no_frixion_money_moov_models_rule_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsRule().update({
    "id": "no_frixion_money_moov_models_rule_id",
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

Create a new `NoFrixionMoneyMoovModelsRuleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsTransactionEntity

```python
no_frixion_money_moov_models_transaction = client.NoFrixionMoneyMoovModelsTransaction()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `str` | No |  |
| `account_name` | `str` | No |  |
| `account_sequence_number` | `int` | No |  |
| `amount` | `float` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `balance` | `float` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `counterparty` | `dict` | No |  |
| `counterparty_summary` | `str` | No |  |
| `currency` | `str` | No |  |
| `description` | `str` | No |  |
| `fx_amount` | `float` | No |  |
| `fx_currency` | `str` | No |  |
| `fx_rate` | `float` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `payment_request_custom_field` | `dict` | No |  |
| `payment_request_id` | `str` | No |  |
| `payout_id` | `str` | No |  |
| `raw_reference` | `str` | No |  |
| `rule_id` | `str` | No |  |
| `tag` | `list` | No |  |
| `their_reference` | `str` | No |  |
| `transaction_date` | `str` | No |  |
| `type` | `str` | No |  |
| `virtual_iban` | `str` | No |  |
| `your_reference` | `str` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsTransaction().load({"id": "no_frixion_money_moov_models_transaction_id"})
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

Create a new `NoFrixionMoneyMoovModelsTransactionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsTransactionPageEntity

```python
no_frixion_money_moov_models_transaction_page = client.NoFrixionMoneyMoovModelsTransactionPage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `str` | No |  |
| `account_name` | `str` | No |  |
| `account_sequence_number` | `int` | No |  |
| `amount` | `float` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `balance` | `float` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `content` | `list` | No |  |
| `counterparty` | `dict` | No |  |
| `counterparty_summary` | `str` | No |  |
| `currency` | `str` | No |  |
| `description` | `str` | No |  |
| `fx_amount` | `float` | No |  |
| `fx_currency` | `str` | No |  |
| `fx_rate` | `float` | No |  |
| `id` | `str` | No |  |
| `inserted` | `str` | No |  |
| `merchant_id` | `str` | No |  |
| `page_number` | `int` | No |  |
| `page_size` | `int` | No |  |
| `payment_request_custom_field` | `dict` | No |  |
| `payment_request_id` | `str` | No |  |
| `payout_id` | `str` | No |  |
| `raw_reference` | `str` | No |  |
| `rule_id` | `str` | No |  |
| `tag` | `list` | No |  |
| `their_reference` | `str` | No |  |
| `total_page` | `int` | No |  |
| `total_size` | `int` | No |  |
| `transaction_date` | `str` | No |  |
| `type` | `str` | No |  |
| `virtual_iban` | `str` | No |  |
| `your_reference` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.NoFrixionMoneyMoovModelsTransactionPage().list()
for no_frixion_money_moov_models_transaction_page in results:
    print(no_frixion_money_moov_models_transaction_page)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsTransactionPage().load({"account_id": "account_id"})
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

Create a new `NoFrixionMoneyMoovModelsTransactionPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsUserInviteEntity

```python
no_frixion_money_moov_models_user_invite = client.NoFrixionMoneyMoovModelsUserInvite()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_status` | `dict` | No |  |
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

### Field Usage by Operation

| Field | load | create |
| --- | --- | --- |
| `authorisation_status` | - | - |
| `id` | - | - |
| `initial_role_id` | - | - |
| `invitee_email_address` | - | Yes |
| `invitee_first_name` | - | - |
| `invitee_last_name` | - | - |
| `inviter_email_address` | - | - |
| `inviter_first_name` | - | - |
| `inviter_last_name` | - | - |
| `is_authorised` | - | - |
| `is_invitee_registered` | - | - |
| `last_invited` | - | - |
| `merchant_id` | - | - |
| `merchant_name` | - | - |
| `message` | - | - |
| `registration_url` | - | - |
| `send_invite_email` | - | - |
| `status` | - | - |
| `user` | - | - |
| `user_id` | - | - |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsUserInvite().create({
    "user": {},  # dict
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsUserInvite().load({"id": "no_frixion_money_moov_models_user_invite_id"})
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

Create a new `NoFrixionMoneyMoovModelsUserInviteEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsUserInvitePageEntity

```python
no_frixion_money_moov_models_user_invite_page = client.NoFrixionMoneyMoovModelsUserInvitePage()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_status` | `dict` | No |  |
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
| `status` | `str` | No |  |
| `user` | `dict` | Yes |  |
| `user_id` | `str` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.NoFrixionMoneyMoovModelsUserInvitePage().list()
for no_frixion_money_moov_models_user_invite_page in results:
    print(no_frixion_money_moov_models_user_invite_page)
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

Create a new `NoFrixionMoneyMoovModelsUserInvitePageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsUserPageEntity

```python
no_frixion_money_moov_models_user_page = client.NoFrixionMoneyMoovModelsUserPage()
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
| `roles_with_scope` | `list` | No |  |
| `two_factor_enabled` | `bool` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.NoFrixionMoneyMoovModelsUserPage().list()
for no_frixion_money_moov_models_user_page in results:
    print(no_frixion_money_moov_models_user_page)
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

Create a new `NoFrixionMoneyMoovModelsUserPageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## NoFrixionMoneyMoovModelsWebhookEntity

```python
no_frixion_money_moov_models_webhook = client.NoFrixionMoneyMoovModelsWebhook()
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

| Field | load | list | create | update |
| --- | --- | --- | --- | --- |
| `destination_url` | - | - | - | - |
| `email_address` | - | - | - | - |
| `failed_notification_email_address` | - | - | - | - |
| `id` | - | - | - | - |
| `is_active` | - | - | - | - |
| `merchant_id` | - | - | Yes | Yes |
| `notification_method` | - | - | Yes | Yes |
| `resource_type` | - | - | - | - |
| `retry` | - | - | - | - |
| `secret` | - | - | - | - |
| `version` | - | - | - | - |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsWebhook().create({
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.NoFrixionMoneyMoovModelsWebhook().list()
for no_frixion_money_moov_models_webhook in results:
    print(no_frixion_money_moov_models_webhook)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsWebhook().load({"id": "no_frixion_money_moov_models_webhook_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.NoFrixionMoneyMoovModelsWebhook().update({
    "id": "no_frixion_money_moov_models_webhook_id",
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

Create a new `NoFrixionMoneyMoovModelsWebhookEntity` instance with the same options.

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

## PaymentRequestEntity

```python
payment_request = client.PaymentRequest()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float` | No |  |
| `do_simulate_settlement_failure` | `bool` | No |  |
| `error_description` | `str` | No |  |
| `institution` | `str` | No |  |
| `payment_initiation_id` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PaymentRequest().create({
    "paymentrequest_id": "example_paymentrequest_id",  # str
})
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

| Field | load | create | update | remove |
| --- | --- | --- | --- | --- |
| `account_id` | - | Yes | - | - |
| `allow_incomplete` | - | - | - | - |
| `amount` | - | - | - | - |
| `amount_minor_unit` | - | - | - | - |
| `approve_payout_url` | - | - | - | - |
| `approver_id` | - | - | - | - |
| `authentication_method` | - | - | - | - |
| `authorisation` | - | - | - | - |
| `authorisers_completed_count` | - | - | - | - |
| `authorisers_required_count` | - | - | - | - |
| `batch_payout_id` | - | - | - | - |
| `beneficiary` | - | - | - | - |
| `beneficiary_id` | - | - | - | - |
| `can_authorise` | - | - | - | - |
| `can_process` | - | - | - | - |
| `can_update` | - | - | - | - |
| `charge_bearer` | - | - | - | - |
| `created_by` | - | - | - | - |
| `created_by_email_address` | - | - | - | - |
| `currency` | - | Yes | - | - |
| `current_user_id` | - | - | - | - |
| `description` | - | - | - | - |
| `destination` | - | - | - | - |
| `document` | - | - | - | - |
| `event` | - | - | - | - |
| `formatted_amount` | - | - | - | - |
| `formatted_fx_destination_amount` | - | - | - | - |
| `formatted_schedule` | - | - | - | - |
| `formatted_schedule_day_only` | - | - | - | - |
| `formatted_source_account_available_balance` | - | - | - | - |
| `fx_destination_amount` | - | - | - | - |
| `fx_destination_amount_minor_unit` | - | - | - | - |
| `fx_destination_currency` | - | - | - | - |
| `fx_quote_expires_at` | - | - | - | - |
| `fx_quote_id` | - | - | - | - |
| `fx_rate` | - | - | - | - |
| `fx_use_destination_amount` | - | - | - | - |
| `has_current_user_authorised` | - | - | - | - |
| `id` | - | - | - | - |
| `inserted` | - | - | - | - |
| `invoice_id` | - | - | - | - |
| `is_archived` | - | - | - | - |
| `is_failed` | - | - | - | - |
| `is_settled` | - | - | - | - |
| `is_submitted` | - | - | - | - |
| `last_updated` | - | - | - | - |
| `merchant_id` | - | - | - | - |
| `merchant_token_description` | - | - | - | - |
| `nonce` | - | - | - | - |
| `payment_processor` | - | - | - | - |
| `payment_rail` | - | - | - | - |
| `payrun_id` | - | - | - | - |
| `payrun_name` | - | - | - | - |
| `rule` | - | - | - | - |
| `schedule_date` | - | - | - | - |
| `scheduled` | - | - | - | - |
| `source_account_available_balance` | - | - | - | - |
| `source_account_available_balance_minor_unit` | - | - | - | - |
| `source_account_bic` | - | - | - | - |
| `source_account_currency` | - | - | - | - |
| `source_account_iban` | - | - | - | - |
| `source_account_identifier` | - | - | - | - |
| `source_account_name` | - | - | - | - |
| `source_account_number` | - | - | - | - |
| `source_account_sortcode` | - | - | - | - |
| `status` | - | - | - | - |
| `tag` | - | - | - | - |
| `tag_id` | - | - | - | - |
| `their_reference` | - | - | - | - |
| `topup_payrun_id` | - | - | - | - |
| `transacted_amount` | - | - | - | - |
| `transacted_fx_amount` | - | - | - | - |
| `transacted_fx_rate` | - | - | - | - |
| `type` | - | Yes | - | - |
| `user_id` | - | - | - | - |
| `your_reference` | - | - | - | - |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Payout().create({
})
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

## PayrunEntity

```python
payrun = client.Payrun()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `str` | No |  |
| `note` | `str` | No |  |
| `scheduled_date` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Payrun().create({
    "id": "example_id",  # str
})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Payrun().remove({"id": "id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Payrun().update({
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

Create a new `PayrunEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## RejectEntity

```python
reject = client.Reject()
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
| `their_reference` | `str` | No |  |
| `topup_payrun_id` | `str` | No |  |
| `transacted_amount` | `float` | No |  |
| `transacted_fx_amount` | `float` | No |  |
| `transacted_fx_rate` | `float` | No |  |
| `type` | `str` | No |  |
| `user_id` | `str` | No |  |
| `your_reference` | `str` | No |  |

### Operations

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Reject().update({
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

Create a new `RejectEntity` instance with the same options.

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

## RuleEntity

```python
rule = client.Rule()
```

### Operations

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Rule().remove({"id": "id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Rule().update({
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

Create a new `RuleEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SendEntity

```python
send = client.Send()
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

| Field | create |
| --- | --- |
| `account_id` | Yes |
| `allow_incomplete` | - |
| `amount` | - |
| `amount_minor_unit` | - |
| `approve_payout_url` | - |
| `approver_id` | - |
| `authentication_method` | - |
| `authorisation` | - |
| `authorisers_completed_count` | - |
| `authorisers_required_count` | - |
| `batch_payout_id` | - |
| `beneficiary` | - |
| `beneficiary_id` | - |
| `can_authorise` | - |
| `can_process` | - |
| `can_update` | - |
| `charge_bearer` | - |
| `created_by` | - |
| `created_by_email_address` | - |
| `currency` | Yes |
| `current_user_id` | - |
| `description` | - |
| `destination` | - |
| `document` | - |
| `event` | - |
| `formatted_amount` | - |
| `formatted_fx_destination_amount` | - |
| `formatted_schedule` | - |
| `formatted_schedule_day_only` | - |
| `formatted_source_account_available_balance` | - |
| `fx_destination_amount` | - |
| `fx_destination_amount_minor_unit` | - |
| `fx_destination_currency` | - |
| `fx_quote_expires_at` | - |
| `fx_quote_id` | - |
| `fx_rate` | - |
| `fx_use_destination_amount` | - |
| `has_current_user_authorised` | - |
| `id` | - |
| `inserted` | - |
| `invoice_id` | - |
| `is_archived` | - |
| `is_failed` | - |
| `is_settled` | - |
| `is_submitted` | - |
| `last_updated` | - |
| `merchant_id` | - |
| `merchant_token_description` | - |
| `nonce` | - |
| `payment_processor` | - |
| `payment_rail` | - |
| `payrun_id` | - |
| `payrun_name` | - |
| `rule` | - |
| `schedule_date` | - |
| `scheduled` | - |
| `source_account_available_balance` | - |
| `source_account_available_balance_minor_unit` | - |
| `source_account_bic` | - |
| `source_account_currency` | - |
| `source_account_iban` | - |
| `source_account_identifier` | - |
| `source_account_name` | - |
| `source_account_number` | - |
| `source_account_sortcode` | - |
| `status` | - |
| `tag` | - |
| `tag_id` | - |
| `their_reference` | - |
| `topup_payrun_id` | - |
| `transacted_amount` | - |
| `transacted_fx_amount` | - |
| `transacted_fx_rate` | - |
| `type` | Yes |
| `user_id` | - |
| `your_reference` | - |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Send().create({
    "beneficiary": {},  # dict
    "source_account_identifier": {},  # dict
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

Create a new `SendEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SendbeneficiaryEntity

```python
sendbeneficiary = client.Sendbeneficiary()
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

| Field | create |
| --- | --- |
| `account_id` | Yes |
| `allow_incomplete` | - |
| `amount` | - |
| `amount_minor_unit` | - |
| `approve_payout_url` | - |
| `approver_id` | - |
| `authentication_method` | - |
| `authorisation` | - |
| `authorisers_completed_count` | - |
| `authorisers_required_count` | - |
| `batch_payout_id` | - |
| `beneficiary` | - |
| `beneficiary_id` | - |
| `can_authorise` | - |
| `can_process` | - |
| `can_update` | - |
| `charge_bearer` | - |
| `created_by` | - |
| `created_by_email_address` | - |
| `currency` | Yes |
| `current_user_id` | - |
| `description` | - |
| `destination` | - |
| `document` | - |
| `event` | - |
| `formatted_amount` | - |
| `formatted_fx_destination_amount` | - |
| `formatted_schedule` | - |
| `formatted_schedule_day_only` | - |
| `formatted_source_account_available_balance` | - |
| `fx_destination_amount` | - |
| `fx_destination_amount_minor_unit` | - |
| `fx_destination_currency` | - |
| `fx_quote_expires_at` | - |
| `fx_quote_id` | - |
| `fx_rate` | - |
| `fx_use_destination_amount` | - |
| `has_current_user_authorised` | - |
| `id` | - |
| `inserted` | - |
| `invoice_id` | - |
| `is_archived` | - |
| `is_failed` | - |
| `is_settled` | - |
| `is_submitted` | - |
| `last_updated` | - |
| `merchant_id` | - |
| `merchant_token_description` | - |
| `nonce` | - |
| `payment_processor` | - |
| `payment_rail` | - |
| `payrun_id` | - |
| `payrun_name` | - |
| `rule` | - |
| `schedule_date` | - |
| `scheduled` | - |
| `source_account_available_balance` | - |
| `source_account_available_balance_minor_unit` | - |
| `source_account_bic` | - |
| `source_account_currency` | - |
| `source_account_iban` | - |
| `source_account_identifier` | - |
| `source_account_name` | - |
| `source_account_number` | - |
| `source_account_sortcode` | - |
| `status` | - |
| `tag` | - |
| `tag_id` | - |
| `their_reference` | - |
| `topup_payrun_id` | - |
| `transacted_amount` | - |
| `transacted_fx_amount` | - |
| `transacted_fx_rate` | - |
| `type` | Yes |
| `user_id` | - |
| `your_reference` | - |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Sendbeneficiary().create({
    "beneficiary": {},  # dict
    "source_account_identifier": {},  # dict
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

Create a new `SendbeneficiaryEntity` instance with the same options.

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

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Transaction().create({
    "id": "example_id",  # str
})
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Transaction().load({"sequence_number": 1, "transaction_id": "transaction_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Transaction().remove({"id": "id"})
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

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.UserInvite().create({
    "id": "example_id",  # str
})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.UserInvite().remove({"id": "id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.UserInvite().update({
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

### Operations

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Webhook().remove({"id": "id"})
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

## WhoamiEntity

```python
whoami = client.Whoami()
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
| `roles_with_scope` | `list` | No |  |
| `two_factor_enabled` | `bool` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Whoami().list()
for whoami in results:
    print(whoami)
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

Create a new `WhoamiEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## WhoamitrustedappEntity

```python
whoamitrustedapp = client.Whoamitrustedapp()
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
| `roles_with_scope` | `list` | No |  |
| `two_factor_enabled` | `bool` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Whoamitrustedapp().list()
for whoamitrustedapp in results:
    print(whoamitrustedapp)
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

Create a new `WhoamitrustedappEntity` instance with the same options.

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

