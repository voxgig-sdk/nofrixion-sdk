# Nofrixion PHP SDK Reference

Complete API reference for the Nofrixion PHP SDK.


## NofrixionSDK

### Constructor

```php
require_once __DIR__ . '/nofrixion_sdk.php';

$client = new NofrixionSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `NofrixionSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = NofrixionSDK::test();
```


### Instance Methods

#### `Account($data = null)`

Create a new `AccountEntity` instance. Pass `null` for no initial data.

#### `Beneficiary($data = null)`

Create a new `BeneficiaryEntity` instance. Pass `null` for no initial data.

#### `Cancel($data = null)`

Create a new `CancelEntity` instance. Pass `null` for no initial data.

#### `Disable($data = null)`

Create a new `DisableEntity` instance. Pass `null` for no initial data.

#### `Enable($data = null)`

Create a new `EnableEntity` instance. Pass `null` for no initial data.

#### `Merchant($data = null)`

Create a new `MerchantEntity` instance. Pass `null` for no initial data.

#### `Metadata($data = null)`

Create a new `MetadataEntity` instance. Pass `null` for no initial data.

#### `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage($data = null)`

Create a new `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionBizBizModelsPagingPaymentRequestPage($data = null)`

Create a new `NoFrixionBizBizModelsPagingPaymentRequestPageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionBizBizModelsPagingPayoutPage($data = null)`

Create a new `NoFrixionBizBizModelsPagingPayoutPageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionBizBizModelsPagingPayrunPage($data = null)`

Create a new `NoFrixionBizBizModelsPagingPayrunPageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionBizBizModelsPagingRuleEventsPage($data = null)`

Create a new `NoFrixionBizBizModelsPagingRuleEventsPageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionBizBizModelsPagingRulesPage($data = null)`

Create a new `NoFrixionBizBizModelsPagingRulesPageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionBizBizModelsPaymentsCardPayment($data = null)`

Create a new `NoFrixionBizBizModelsPaymentsCardPaymentEntity` instance. Pass `null` for no initial data.

#### `NoFrixionBizBizModelsPaymentsCardPublicKey($data = null)`

Create a new `NoFrixionBizBizModelsPaymentsCardPublicKeyEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries($data = null)`

Create a new `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment($data = null)`

Create a new `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate($data = null)`

Create a new `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovApiFeaturesUserInvitesCreate($data = null)`

Create a new `NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant($data = null)`

Create a new `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsBatchPayout($data = null)`

Create a new `NoFrixionMoneyMoovModelsBatchPayoutEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsBeneficiaryGroupPage($data = null)`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsBeneficiaryPage($data = null)`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryPageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsCardCustomerToken($data = null)`

Create a new `NoFrixionMoneyMoovModelsCardCustomerTokenEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsCurrencyCurrencyInfo($data = null)`

Create a new `NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsDirectDebitBatchSubmit($data = null)`

Create a new `NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsFxRate($data = null)`

Create a new `NoFrixionMoneyMoovModelsFxRateEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsIPayment($data = null)`

Create a new `NoFrixionMoneyMoovModelsIPaymentEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsMandatesMandate($data = null)`

Create a new `NoFrixionMoneyMoovModelsMandatesMandateEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchant($data = null)`

Create a new `NoFrixionMoneyMoovModelsMerchantEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantPage($data = null)`

Create a new `NoFrixionMoneyMoovModelsMerchantPageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantPayByBankSetting($data = null)`

Create a new `NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantToken($data = null)`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsMerchantTokenPage($data = null)`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenPageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsNoFrixionVersion($data = null)`

Create a new `NoFrixionMoneyMoovModelsNoFrixionVersionEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsOpenBankingAccount($data = null)`

Create a new `NoFrixionMoneyMoovModelsOpenBankingAccountEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsOpenBankingConsent($data = null)`

Create a new `NoFrixionMoneyMoovModelsOpenBankingConsentEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsOpenBankingTransaction($data = null)`

Create a new `NoFrixionMoneyMoovModelsOpenBankingTransactionEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPayment($data = null)`

Create a new `NoFrixionMoneyMoovModelsPaymentEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentAccountMinimalPage($data = null)`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentAccountPage($data = null)`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountPageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentInitiation($data = null)`

Create a new `NoFrixionMoneyMoovModelsPaymentInitiationEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestEvent($data = null)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestEventEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestMetric($data = null)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMetricEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestMinimal($data = null)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestResult($data = null)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestResultEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment($data = null)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2($data = null)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3($data = null)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4($data = null)`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPayoutKeysetPage($data = null)`

Create a new `NoFrixionMoneyMoovModelsPayoutKeysetPageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPayoutMetric($data = null)`

Create a new `NoFrixionMoneyMoovModelsPayoutMetricEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPayoutsPayoutsCreate($data = null)`

Create a new `NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsPayrun($data = null)`

Create a new `NoFrixionMoneyMoovModelsPayrunEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsReportResult($data = null)`

Create a new `NoFrixionMoneyMoovModelsReportResultEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsRule($data = null)`

Create a new `NoFrixionMoneyMoovModelsRuleEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsTransaction($data = null)`

Create a new `NoFrixionMoneyMoovModelsTransactionEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsTransactionPage($data = null)`

Create a new `NoFrixionMoneyMoovModelsTransactionPageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsUserInvite($data = null)`

Create a new `NoFrixionMoneyMoovModelsUserInviteEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsUserInvitePage($data = null)`

Create a new `NoFrixionMoneyMoovModelsUserInvitePageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsUserPage($data = null)`

Create a new `NoFrixionMoneyMoovModelsUserPageEntity` instance. Pass `null` for no initial data.

#### `NoFrixionMoneyMoovModelsWebhook($data = null)`

Create a new `NoFrixionMoneyMoovModelsWebhookEntity` instance. Pass `null` for no initial data.

#### `OpenBanking($data = null)`

Create a new `OpenBankingEntity` instance. Pass `null` for no initial data.

#### `Payeeverification($data = null)`

Create a new `PayeeverificationEntity` instance. Pass `null` for no initial data.

#### `PaymentRequest($data = null)`

Create a new `PaymentRequestEntity` instance. Pass `null` for no initial data.

#### `Payout($data = null)`

Create a new `PayoutEntity` instance. Pass `null` for no initial data.

#### `Payrun($data = null)`

Create a new `PayrunEntity` instance. Pass `null` for no initial data.

#### `Reject($data = null)`

Create a new `RejectEntity` instance. Pass `null` for no initial data.

#### `Report($data = null)`

Create a new `ReportEntity` instance. Pass `null` for no initial data.

#### `Rule($data = null)`

Create a new `RuleEntity` instance. Pass `null` for no initial data.

#### `Send($data = null)`

Create a new `SendEntity` instance. Pass `null` for no initial data.

#### `Sendbeneficiary($data = null)`

Create a new `SendbeneficiaryEntity` instance. Pass `null` for no initial data.

#### `Tag($data = null)`

Create a new `TagEntity` instance. Pass `null` for no initial data.

#### `Token($data = null)`

Create a new `TokenEntity` instance. Pass `null` for no initial data.

#### `Transaction($data = null)`

Create a new `TransactionEntity` instance. Pass `null` for no initial data.

#### `User($data = null)`

Create a new `UserEntity` instance. Pass `null` for no initial data.

#### `UserInvite($data = null)`

Create a new `UserInviteEntity` instance. Pass `null` for no initial data.

#### `Virtual($data = null)`

Create a new `VirtualEntity` instance. Pass `null` for no initial data.

#### `Webhook($data = null)`

Create a new `WebhookEntity` instance. Pass `null` for no initial data.

#### `Whoami($data = null)`

Create a new `WhoamiEntity` instance. Pass `null` for no initial data.

#### `Whoamitrustedapp($data = null)`

Create a new `WhoamitrustedappEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): NofrixionUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## AccountEntity

```php
$account = $client->Account();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `account_name` | `string` | No |  |
| `account_supplier_name` | `string` | No |  |
| `account_type` | `string` | No |  |
| `available_balance` | `float` | No |  |
| `available_balance_minor_unit` | `int` | No |  |
| `balance` | `float` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `bank_name` | `string` | No |  |
| `consent_id` | `string` | No |  |
| `created_by` | `array` | Yes |  |
| `created_by_display_name` | `string` | No |  |
| `currency` | `string` | No |  |
| `default_payment_rail` | `string` | No |  |
| `display_name` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `external_account_icon` | `string` | No |  |
| `format` | `string` | No |  |
| `from_date` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `array` | Yes |  |
| `inserted` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_connected_account` | `bool` | No |  |
| `is_default` | `bool` | No |  |
| `is_trust_account` | `bool` | No |  |
| `is_virtual` | `bool` | No |  |
| `last_transaction` | `array` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `physical_account_id` | `string` | No |  |
| `role_i_d` | `array` | No |  |
| `rule` | `array` | No |  |
| `submitted_payouts_balance` | `float` | No |  |
| `submitted_payouts_balance_minor_unit` | `int` | No |  |
| `summary` | `string` | No |  |
| `supplier_physical_account_id` | `string` | No |  |
| `supplier_sepa_instant_status` | `string` | No |  |
| `to_date` | `string` | No |  |
| `xero_bank_feed_connection_status` | `string` | No |  |
| `xero_bank_feed_last_synced_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` | No |  |
| `xero_bank_feed_sync_status` | `string` | No |  |
| `xero_unsynchronised_transactions_count` | `int` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Account()->create([
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Account()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Account()->load(["id" => "account_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Account()->remove(["id" => "account_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Account()->update([
  "id" => "account_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): AccountEntity`

Create a new `AccountEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## BeneficiaryEntity

```php
$beneficiary = $client->Beneficiary();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `array` | No |  |
| `authorisation` | `array` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `beneficiary_event` | `array` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `created_by` | `array` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `array` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `bool` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `array` | No |  |
| `source_account_i_d` | `array` | No |  |
| `their_reference` | `string` | No |  |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Beneficiary()->create([
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Beneficiary()->load(["id" => "beneficiary_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Beneficiary()->remove(["id" => "beneficiary_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Beneficiary()->update([
  "id" => "beneficiary_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): BeneficiaryEntity`

Create a new `BeneficiaryEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## CancelEntity

```php
$cancel = $client->Cancel();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `float` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `array` | No |  |
| `authorisation` | `array` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `array` | Yes |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `array` | No |  |
| `document` | `array` | No |  |
| `event` | `array` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `float` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `float` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `array` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `array` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `array` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `float` | No |  |
| `transacted_fx_amount` | `float` | No |  |
| `transacted_fx_rate` | `float` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Cancel()->update([
  "id" => "id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): CancelEntity`

Create a new `CancelEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## DisableEntity

```php
$disable = $client->Disable();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `array` | No |  |
| `authorisation` | `array` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `beneficiary_event` | `array` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `created_by` | `array` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `array` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `bool` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `array` | No |  |
| `their_reference` | `string` | No |  |

### Operations

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Disable()->update([
  "id" => "id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): DisableEntity`

Create a new `DisableEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## EnableEntity

```php
$enable = $client->Enable();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `array` | No |  |
| `authorisation` | `array` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `beneficiary_event` | `array` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `created_by` | `array` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `array` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `bool` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `array` | No |  |
| `their_reference` | `string` | No |  |

### Operations

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Enable()->update([
  "id" => "id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): EnableEntity`

Create a new `EnableEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MerchantEntity

```php
$merchant = $client->Merchant();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `reason` | `string` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Merchant()->load(["merchant_id" => "merchant_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Merchant()->remove();
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Merchant()->update([
  "id" => "id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MerchantEntity`

Create a new `MerchantEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MetadataEntity

```php
$metadata = $client->Metadata();
```

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Metadata()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MetadataEntity`

Create a new `MetadataEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity

```php
$no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page = $client->NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage();
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
| `is_recurring` | `bool` | No |  |
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

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity`

Create a new `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingPaymentRequestPageEntity

```php
$no_frixion_biz_biz_models_paging_payment_request_page = $client->NoFrixionBizBizModelsPagingPaymentRequestPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `array` | No |  |
| `amount` | `float` | No |  |
| `amount_pending` | `float` | No |  |
| `amount_received` | `float` | No |  |
| `amount_refunded` | `float` | No |  |
| `auto_send_receipt` | `bool` | No |  |
| `base_origin_url` | `string` | No |  |
| `callback_url` | `string` | No |  |
| `card_authorize_only` | `bool` | No |  |
| `card_create_token` | `bool` | No |  |
| `card_create_token_mode` | `string` | No |  |
| `card_ignore_cvn` | `bool` | No |  |
| `card_processor_merchant_id` | `string` | No |  |
| `card_stripe_payment_intent_id` | `string` | No |  |
| `card_stripe_payment_intent_secret` | `string` | No |  |
| `created_by_user` | `array` | Yes |  |
| `currency` | `string` | No |  |
| `custom_field` | `array` | No |  |
| `customer_email_address` | `string` | No |  |
| `customer_id` | `string` | No |  |
| `customer_name` | `string` | No |  |
| `description` | `string` | No |  |
| `destination_account` | `array` | No |  |
| `direct_debit_payment` | `array` | No |  |
| `due_date` | `string` | No |  |
| `event` | `array` | No |  |
| `failure_callback_url` | `string` | No |  |
| `field_display_setting` | `array` | No |  |
| `formatted_amount` | `string` | No |  |
| `hosted_pay_checkout_url` | `string` | No |  |
| `id` | `string` | No |  |
| `ignore_address_verification` | `bool` | No |  |
| `inserted` | `string` | No |  |
| `inserted_sortable` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `jwk` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `lightning_invoice` | `string` | No |  |
| `lightning_invoice_expires_at` | `string` | No |  |
| `merchant_direct_debit_mandate_id` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `notification_email_address` | `string` | No |  |
| `notification_role_i_d` | `array` | No |  |
| `order_id` | `string` | No |  |
| `partial_payment_method` | `string` | No |  |
| `partial_payment_step` | `string` | No |  |
| `payment_attempt` | `array` | No |  |
| `payment_method` | `array` | No |  |
| `payment_processor` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `pisp_account_id` | `string` | No |  |
| `priority_bank_id` | `string` | No |  |
| `result` | `array` | No |  |
| `sandbox_settle_delay_in_second` | `int` | No |  |
| `shipping_address` | `array` | No |  |
| `status` | `string` | No |  |
| `success_web_hook_url` | `string` | No |  |
| `tag` | `array` | No |  |
| `title` | `string` | No |  |
| `tokenised_card` | `array` | No |  |
| `transaction` | `array` | No |  |
| `use_hosted_payment_page` | `bool` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionBizBizModelsPagingPaymentRequestPage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionBizBizModelsPagingPaymentRequestPageEntity`

Create a new `NoFrixionBizBizModelsPagingPaymentRequestPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingPayoutPageEntity

```php
$no_frixion_biz_biz_models_paging_payout_page = $client->NoFrixionBizBizModelsPagingPayoutPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `float` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `array` | No |  |
| `authorisation` | `array` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `array` | Yes |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `array` | No |  |
| `document` | `array` | No |  |
| `event` | `array` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `float` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `float` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `array` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `array` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `array` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `float` | No |  |
| `transacted_fx_amount` | `float` | No |  |
| `transacted_fx_rate` | `float` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionBizBizModelsPagingPayoutPage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionBizBizModelsPagingPayoutPageEntity`

Create a new `NoFrixionBizBizModelsPagingPayoutPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingPayrunPageEntity

```php
$no_frixion_biz_biz_models_paging_payrun_page = $client->NoFrixionBizBizModelsPagingPayrunPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation` | `array` | No |  |
| `authorisation_date` | `string` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_delete` | `bool` | No |  |
| `can_edit` | `bool` | No |  |
| `event` | `array` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice` | `array` | No |  |
| `invoices_minimal` | `array` | No |  |
| `is_archived` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `last_updated_by` | `array` | Yes |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment` | `array` | No |  |
| `payout` | `array` | No |  |
| `payouts_count` | `int` | No |  |
| `schedule_date` | `string` | No |  |
| `source_account` | `array` | No |  |
| `status` | `string` | No |  |
| `total_eur` | `float` | No |  |
| `total_gbp` | `float` | No |  |
| `total_usd` | `float` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionBizBizModelsPagingPayrunPage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionBizBizModelsPagingPayrunPageEntity`

Create a new `NoFrixionBizBizModelsPagingPayrunPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingRuleEventsPageEntity

```php
$no_frixion_biz_biz_models_paging_rule_events_page = $client->NoFrixionBizBizModelsPagingRuleEventsPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `error_message` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_authorise_to_enable` | `bool` | No |  |
| `message` | `string` | No |  |
| `raw_response` | `string` | No |  |
| `rule_event_type` | `string` | No |  |
| `rule_id` | `string` | No |  |
| `user` | `array` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionBizBizModelsPagingRuleEventsPage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionBizBizModelsPagingRuleEventsPageEntity`

Create a new `NoFrixionBizBizModelsPagingRuleEventsPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionBizBizModelsPagingRulesPageEntity

```php
$no_frixion_biz_biz_models_paging_rules_page = $client->NoFrixionBizBizModelsPagingRulesPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account` | `array` | No |  |
| `account_id` | `string` | No |  |
| `approve_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `array` | No |  |
| `authorisation` | `array` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `can_authorise` | `bool` | No |  |
| `created_by` | `array` | Yes |  |
| `description` | `string` | No |  |
| `end_at` | `string` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_disabled` | `bool` | No |  |
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
| `sweep_action` | `array` | No |  |
| `time_zone_id` | `string` | No |  |
| `trigger_cron_expression` | `string` | No |  |
| `trigger_on_pay_in` | `bool` | No |  |
| `user_id` | `string` | No |  |
| `web_hook_secret` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionBizBizModelsPagingRulesPage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionBizBizModelsPagingRulesPageEntity`

Create a new `NoFrixionBizBizModelsPagingRulesPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionBizBizModelsPaymentsCardPaymentEntity

```php
$no_frixion_biz_biz_models_payments_card_payment = $client->NoFrixionBizBizModelsPaymentsCardPayment();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorized_amount` | `string` | No |  |
| `currency_code` | `string` | No |  |
| `is_payer_authentication_required` | `bool` | No |  |
| `is_soft_decline` | `bool` | No |  |
| `payer_authentication_access_token` | `string` | No |  |
| `payer_authentication_merchant_data` | `string` | No |  |
| `payer_authentication_url` | `string` | No |  |
| `payer_authentication_window_height` | `int` | No |  |
| `payer_authentication_window_width` | `int` | No |  |
| `payment_request_callback_url` | `string` | No |  |
| `payment_request_id` | `string` | No |  |
| `request_id` | `string` | No |  |
| `response_code` | `string` | No |  |
| `response_type` | `string` | No |  |
| `status` | `string` | No |  |
| `three_ds_redirect_url` | `string` | No |  |
| `transaction_id` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionBizBizModelsPaymentsCardPayment()->create([
  "paymentrequest_id" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionBizBizModelsPaymentsCardPaymentEntity`

Create a new `NoFrixionBizBizModelsPaymentsCardPaymentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionBizBizModelsPaymentsCardPublicKeyEntity

```php
$no_frixion_biz_biz_models_payments_card_public_key = $client->NoFrixionBizBizModelsPaymentsCardPublicKey();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `jwt` | `string` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionBizBizModelsPaymentsCardPublicKey()->load(["paymentrequest_id" => "paymentrequest_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionBizBizModelsPaymentsCardPublicKeyEntity`

Create a new `NoFrixionBizBizModelsPaymentsCardPublicKeyEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity

```php
$no_frixion_money_moov_api_features_beneficiaries_beneficiaries = $client->NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `beneficiary` | `array` | No |  |
| `failed_beneficiary` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries()->create([
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity`

Create a new `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity

```php
$no_frixion_money_moov_api_features_payment_requests_payment = $client->NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_payment_request` | `array` | No |  |
| `payment_request` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment()->create([
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity`

Create a new `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity

```php
$no_frixion_money_moov_api_features_permissions_roles_create = $client->NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_role` | `array` | No |  |
| `role` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate()->create([
  "merchant_id" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity`

Create a new `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity

```php
$no_frixion_money_moov_api_features_user_invites_create = $client->NoFrixionMoneyMoovApiFeaturesUserInvitesCreate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_user_invite` | `array` | No |  |
| `user_invite` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovApiFeaturesUserInvitesCreate()->create([
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity`

Create a new `NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity

```php
$no_frixion_money_moov_models_authorisation_settings_merchant = $client->NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount_lower` | `float` | No |  |
| `amount_upper` | `float` | No |  |
| `authorisation_type` | `string` | No |  |
| `beneficiaries_only` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `last_editor_cant_authorise` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `number_of_authoriser` | `int` | No |  |
| `role_setting` | `array` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity`

Create a new `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsBatchPayoutEntity

```php
$no_frixion_money_moov_models_batch_payout = $client->NoFrixionMoneyMoovModelsBatchPayout();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approve_url` | `string` | No |  |
| `id` | `string` | No |  |
| `payout` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsBatchPayout()->create([
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsBatchPayout()->load(["id" => "no_frixion_money_moov_models_batch_payout_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsBatchPayoutEntity`

Create a new `NoFrixionMoneyMoovModelsBatchPayoutEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity

```php
$no_frixion_money_moov_models_beneficiary_group_page = $client->NoFrixionMoneyMoovModelsBeneficiaryGroupPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `group_member` | `array` | No |  |
| `group_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsBeneficiaryGroupPage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsBeneficiaryPageEntity

```php
$no_frixion_money_moov_models_beneficiary_page = $client->NoFrixionMoneyMoovModelsBeneficiaryPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `approval_callback_url` | `string` | No |  |
| `authentication_method` | `array` | No |  |
| `authorisation` | `array` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `beneficiary_event` | `array` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `created_by` | `array` | Yes |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | Yes |  |
| `destination` | `array` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_enabled` | `bool` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `nonce` | `string` | No |  |
| `source_account` | `array` | No |  |
| `their_reference` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsBeneficiaryPage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsBeneficiaryPageEntity`

Create a new `NoFrixionMoneyMoovModelsBeneficiaryPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsCardCustomerTokenEntity

```php
$no_frixion_money_moov_models_card_customer_token = $client->NoFrixionMoneyMoovModelsCardCustomerToken();
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

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsCardCustomerToken()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsCardCustomerToken()->load(["customer_email_address" => "customer_email_address"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsCardCustomerToken()->remove();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsCardCustomerTokenEntity`

Create a new `NoFrixionMoneyMoovModelsCardCustomerTokenEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity

```php
$no_frixion_money_moov_models_currency_currency_info = $client->NoFrixionMoneyMoovModelsCurrencyCurrencyInfo();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `code` | `string` | No |  |
| `decimal` | `int` | No |  |
| `is_fiat` | `bool` | No |  |
| `iso4217_alpha_code` | `string` | No |  |
| `iso4217_numeric_code` | `string` | No |  |
| `symbol` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsCurrencyCurrencyInfo()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity`

Create a new `NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity

```php
$no_frixion_money_moov_models_direct_debit_batch_submit = $client->NoFrixionMoneyMoovModelsDirectDebitBatchSubmit();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_submission` | `array` | No |  |
| `successful_submission` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsDirectDebitBatchSubmit()->create([
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity`

Create a new `NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsFxRateEntity

```php
$no_frixion_money_moov_models_fx_rate = $client->NoFrixionMoneyMoovModelsFxRate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destination_currency` | `string` | No |  |
| `exchange_rate` | `float` | No |  |
| `expiry_time` | `string` | No |  |
| `quote_id` | `string` | No |  |
| `source_currency` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsFxRate()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsFxRate()->load(["destination" => "destination", "source" => "source", "valid_for_minute" => 1]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsFxRateEntity`

Create a new `NoFrixionMoneyMoovModelsFxRateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsIPaymentEntity

```php
$no_frixion_money_moov_models_i_payment = $client->NoFrixionMoneyMoovModelsIPayment();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `payment_request_id` | `string` | No |  |
| `response_type` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsIPayment()->create([
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsIPaymentEntity`

Create a new `NoFrixionMoneyMoovModelsIPaymentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMandatesMandateEntity

```php
$no_frixion_money_moov_models_mandates_mandate = $client->NoFrixionMoneyMoovModelsMandatesMandate();
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
| `is_recurring` | `bool` | No |  |
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsMandatesMandate()->create([
  "address_line1" => null, // string
  "city" => null, // string
  "country_code" => null, // string
  "email_address" => null, // string
  "first_name" => null, // string
  "last_name" => null, // string
  "postal_code" => null, // string
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsMandatesMandate()->load(["id" => "no_frixion_money_moov_models_mandates_mandate_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsMandatesMandateEntity`

Create a new `NoFrixionMoneyMoovModelsMandatesMandateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantEntity

```php
$no_frixion_money_moov_models_merchant = $client->NoFrixionMoneyMoovModelsMerchant();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_currency` | `array` | No |  |
| `can_have_trust_account` | `bool` | No |  |
| `card_payment_processor` | `string` | No |  |
| `company_id` | `string` | No |  |
| `display_qr_on_hosted_pay` | `bool` | No |  |
| `hosted_pay_version` | `int` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_blocked` | `bool` | No |  |
| `is_exited` | `bool` | No |  |
| `is_suspended` | `bool` | No |  |
| `jurisdiction` | `string` | No |  |
| `logo_url_png` | `string` | No |  |
| `logo_url_svg` | `string` | No |  |
| `merchant_category_code` | `string` | No |  |
| `name` | `string` | No |  |
| `note` | `string` | No |  |
| `parent_merchant` | `array` | No |  |
| `payment_account` | `array` | No |  |
| `payment_account_limit` | `int` | No |  |
| `short_name` | `string` | No |  |
| `supported_payment_methods_list` | `array` | No |  |
| `suspension_reason` | `string` | No |  |
| `tag` | `array` | No |  |
| `time_zone_id` | `string` | No |  |
| `trading_name` | `string` | No |  |
| `web_hook_limit` | `int` | No |  |
| `your_role_name` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsMerchant()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsMerchant()->load(["id" => "no_frixion_money_moov_models_merchant_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsMerchant()->update([
  "id" => "no_frixion_money_moov_models_merchant_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsMerchantEntity`

Create a new `NoFrixionMoneyMoovModelsMerchantEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantPageEntity

```php
$no_frixion_money_moov_models_merchant_page = $client->NoFrixionMoneyMoovModelsMerchantPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_currency` | `array` | No |  |
| `can_have_trust_account` | `bool` | No |  |
| `card_payment_processor` | `string` | No |  |
| `company_id` | `string` | No |  |
| `display_qr_on_hosted_pay` | `bool` | No |  |
| `hosted_pay_version` | `int` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_blocked` | `bool` | No |  |
| `is_exited` | `bool` | No |  |
| `is_suspended` | `bool` | No |  |
| `jurisdiction` | `string` | No |  |
| `logo_url_png` | `string` | No |  |
| `logo_url_svg` | `string` | No |  |
| `merchant_category_code` | `string` | No |  |
| `name` | `string` | No |  |
| `note` | `string` | No |  |
| `parent_merchant` | `array` | No |  |
| `payment_account` | `array` | No |  |
| `payment_account_limit` | `int` | No |  |
| `short_name` | `string` | No |  |
| `supported_payment_methods_list` | `array` | No |  |
| `suspension_reason` | `string` | No |  |
| `tag` | `array` | No |  |
| `time_zone_id` | `string` | No |  |
| `trading_name` | `string` | No |  |
| `web_hook_limit` | `int` | No |  |
| `your_role_name` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsMerchantPage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsMerchantPageEntity`

Create a new `NoFrixionMoneyMoovModelsMerchantPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity

```php
$no_frixion_money_moov_models_merchant_pay_by_bank_setting = $client->NoFrixionMoneyMoovModelsMerchantPayByBankSetting();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `bank_country_code` | `array` | No |  |
| `bank_id` | `string` | No |  |
| `bank_name` | `string` | No |  |
| `business_institution_id` | `string` | No |  |
| `currency` | `string` | No |  |
| `logo` | `string` | No |  |
| `message` | `string` | No |  |
| `message_image_url` | `string` | No |  |
| `order` | `int` | No |  |
| `personal_institution_id` | `string` | No |  |
| `processor` | `string` | No |  |
| `warning_heading` | `string` | No |  |
| `warning_message` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsMerchantPayByBankSetting()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity`

Create a new `NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantTokenEntity

```php
$no_frixion_money_moov_models_merchant_token = $client->NoFrixionMoneyMoovModelsMerchantToken();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authentication_method` | `array` | No |  |
| `authorisation` | `array` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `can_authorise` | `bool` | No |  |
| `description` | `string` | No |  |
| `expires_at` | `string` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `hmac_algorithm` | `string` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `ip_address_whitelist` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_enabled` | `bool` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `nonce` | `string` | Yes |  |
| `permission_type` | `array` | No |  |
| `request_signature_version` | `int` | No |  |
| `shared_secret_algorithm` | `string` | No |  |
| `shared_secret_base64` | `string` | No |  |
| `token` | `string` | No |  |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsMerchantToken()->create([
  "nonce" => null, // string
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsMerchantToken()->load(["id" => "no_frixion_money_moov_models_merchant_token_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsMerchantToken()->update([
  "id" => "no_frixion_money_moov_models_merchant_token_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsMerchantTokenEntity`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsMerchantTokenPageEntity

```php
$no_frixion_money_moov_models_merchant_token_page = $client->NoFrixionMoneyMoovModelsMerchantTokenPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authentication_method` | `array` | No |  |
| `authorisation` | `array` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `can_authorise` | `bool` | No |  |
| `description` | `string` | No |  |
| `expires_at` | `string` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `ip_address_whitelist` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_enabled` | `bool` | No |  |
| `last_authorised` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `nonce` | `string` | Yes |  |
| `permission_type` | `array` | No |  |
| `request_signature_version` | `int` | No |  |
| `shared_secret_algorithm` | `string` | No |  |
| `shared_secret_base64` | `string` | No |  |
| `token` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsMerchantTokenPage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsMerchantTokenPageEntity`

Create a new `NoFrixionMoneyMoovModelsMerchantTokenPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsNoFrixionVersionEntity

```php
$no_frixion_money_moov_models_no_frixion_version = $client->NoFrixionMoneyMoovModelsNoFrixionVersion();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `build_version` | `int` | No |  |
| `major_version` | `int` | No |  |
| `minor_version` | `int` | No |  |
| `release_name` | `string` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsNoFrixionVersion()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsNoFrixionVersionEntity`

Create a new `NoFrixionMoneyMoovModelsNoFrixionVersionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsOpenBankingAccountEntity

```php
$no_frixion_money_moov_models_open_banking_account = $client->NoFrixionMoneyMoovModelsOpenBankingAccount();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_balance` | `array` | No |  |
| `account_identification` | `array` | No |  |
| `account_name` | `array` | No |  |
| `account_type` | `string` | No |  |
| `balance` | `float` | No |  |
| `consolidated_account_information` | `array` | No |  |
| `currency` | `string` | No |  |
| `description` | `string` | No |  |
| `detail` | `string` | No |  |
| `id` | `string` | No |  |
| `nickname` | `string` | No |  |
| `type` | `string` | No |  |
| `usage_type` | `string` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsOpenBankingAccount()->load(["id" => "no_frixion_money_moov_models_open_banking_account_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsOpenBankingAccountEntity`

Create a new `NoFrixionMoneyMoovModelsOpenBankingAccountEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsOpenBankingConsentEntity

```php
$no_frixion_money_moov_models_open_banking_consent = $client->NoFrixionMoneyMoovModelsOpenBankingConsent();
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
| `is_connected_account` | `bool` | No |  |
| `is_enabled` | `bool` | No |  |
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsOpenBankingConsent()->create([
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsOpenBankingConsent()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsOpenBankingConsent()->load(["id" => "no_frixion_money_moov_models_open_banking_consent_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsOpenBankingConsent()->remove(["id" => "no_frixion_money_moov_models_open_banking_consent_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsOpenBankingConsent()->update([
  "id" => "no_frixion_money_moov_models_open_banking_consent_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsOpenBankingConsentEntity`

Create a new `NoFrixionMoneyMoovModelsOpenBankingConsentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsOpenBankingTransactionEntity

```php
$no_frixion_money_moov_models_open_banking_transaction = $client->NoFrixionMoneyMoovModelsOpenBankingTransaction();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address_detail` | `array` | No |  |
| `amount` | `float` | No |  |
| `balance` | `array` | No |  |
| `booking_date_time` | `string` | No |  |
| `charge_detail` | `array` | No |  |
| `currency` | `string` | No |  |
| `currency_exchange` | `array` | No |  |
| `date` | `string` | No |  |
| `description` | `string` | No |  |
| `enrichment` | `array` | No |  |
| `gross_amount` | `array` | Yes |  |
| `id` | `string` | No |  |
| `iso_bank_transaction_code` | `array` | No |  |
| `merchant` | `array` | No |  |
| `payee_detail` | `array` | Yes |  |
| `payer_detail` | `array` | Yes |  |
| `proprietary_bank_transaction_code` | `array` | No |  |
| `reference` | `string` | No |  |
| `statement_reference` | `array` | No |  |
| `status` | `string` | No |  |
| `supplementary_data` | `mixed` | No |  |
| `transaction_amount` | `array` | Yes |  |
| `transaction_information` | `array` | No |  |
| `transaction_mutability` | `string` | No |  |
| `value_date_time` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsOpenBankingTransaction()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsOpenBankingTransactionEntity`

Create a new `NoFrixionMoneyMoovModelsOpenBankingTransactionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentEntity

```php
$no_frixion_money_moov_models_payment = $client->NoFrixionMoneyMoovModelsPayment();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `address` | `array` | No |  |
| `amount` | `float` | No |  |
| `amount_pending` | `float` | No |  |
| `amount_received` | `float` | No |  |
| `amount_refunded` | `float` | No |  |
| `auto_send_receipt` | `bool` | No |  |
| `base_origin_url` | `string` | No |  |
| `callback_url` | `string` | No |  |
| `card_authorize_only` | `bool` | No |  |
| `card_create_token` | `bool` | No |  |
| `card_create_token_mode` | `string` | No |  |
| `card_ignore_cvn` | `bool` | No |  |
| `card_no_payer_authentication` | `bool` | No |  |
| `card_processor_merchant_id` | `string` | No |  |
| `card_stripe_payment_intent_id` | `string` | No |  |
| `card_stripe_payment_intent_secret` | `string` | No |  |
| `card_transmit_raw_detail` | `bool` | No |  |
| `created_by_user` | `array` | Yes |  |
| `currency` | `string` | No |  |
| `custom_field` | `array` | No |  |
| `customer_email_address` | `string` | No |  |
| `customer_id` | `string` | No |  |
| `customer_name` | `string` | No |  |
| `description` | `string` | No |  |
| `destination_account` | `array` | No |  |
| `direct_debit_payment` | `array` | No |  |
| `due_date` | `string` | No |  |
| `event` | `array` | No |  |
| `failure_callback_url` | `string` | No |  |
| `field_display_setting` | `array` | No |  |
| `formatted_amount` | `string` | No |  |
| `hosted_pay_checkout_url` | `string` | No |  |
| `id` | `string` | No |  |
| `ignore_address_verification` | `bool` | No |  |
| `inserted` | `string` | No |  |
| `inserted_sortable` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `jwk` | `string` | No |  |
| `last_updated` | `string` | No |  |
| `lightning_invoice` | `string` | No |  |
| `lightning_invoice_expires_at` | `string` | No |  |
| `merchant_direct_debit_mandate_id` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `notification_email_address` | `string` | No |  |
| `notification_role_i_d` | `array` | No |  |
| `order_id` | `string` | No |  |
| `partial_payment_method` | `string` | No |  |
| `partial_payment_step` | `string` | No |  |
| `payment_attempt` | `array` | No |  |
| `payment_method` | `array` | No |  |
| `payment_processor` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `pisp_account_id` | `string` | No |  |
| `priority_bank_id` | `string` | No |  |
| `result` | `array` | No |  |
| `sandbox_settle_delay_in_second` | `int` | No |  |
| `shipping_address` | `array` | No |  |
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
| `tag` | `array` | No |  |
| `tag_id` | `array` | No |  |
| `title` | `string` | No |  |
| `tokenised_card` | `array` | No |  |
| `transaction` | `array` | No |  |
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsPayment()->create([
  "created_by_user" => null, // array
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsPayment()->load(["id" => "no_frixion_money_moov_models_payment_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsPayment()->update([
  "id" => "no_frixion_money_moov_models_payment_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPaymentEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity

```php
$no_frixion_money_moov_models_payment_account_minimal_page = $client->NoFrixionMoneyMoovModelsPaymentAccountMinimalPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_name` | `string` | No |  |
| `available_balance` | `float` | No |  |
| `balance` | `float` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `currency` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `array` | Yes |  |
| `is_archived` | `bool` | No |  |
| `is_connected_account` | `bool` | No |  |
| `merchant_id` | `string` | No |  |
| `submitted_payouts_balance` | `float` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsPaymentAccountMinimalPage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentAccountPageEntity

```php
$no_frixion_money_moov_models_payment_account_page = $client->NoFrixionMoneyMoovModelsPaymentAccountPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_name` | `string` | No |  |
| `account_supplier_name` | `string` | No |  |
| `available_balance` | `float` | No |  |
| `available_balance_minor_unit` | `int` | No |  |
| `balance` | `float` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `bank_name` | `string` | No |  |
| `consent_id` | `string` | No |  |
| `created_by` | `array` | Yes |  |
| `created_by_display_name` | `string` | No |  |
| `currency` | `string` | No |  |
| `default_payment_rail` | `string` | No |  |
| `display_name` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `external_account_icon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `array` | Yes |  |
| `inserted` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_connected_account` | `bool` | No |  |
| `is_default` | `bool` | No |  |
| `is_trust_account` | `bool` | No |  |
| `is_virtual` | `bool` | No |  |
| `last_transaction` | `array` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `physical_account_id` | `string` | No |  |
| `rule` | `array` | No |  |
| `submitted_payouts_balance` | `float` | No |  |
| `submitted_payouts_balance_minor_unit` | `int` | No |  |
| `summary` | `string` | No |  |
| `supplier_sepa_instant_status` | `string` | No |  |
| `xero_bank_feed_connection_status` | `string` | No |  |
| `xero_bank_feed_last_synced_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` | No |  |
| `xero_bank_feed_sync_status` | `string` | No |  |
| `xero_unsynchronised_transactions_count` | `int` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsPaymentAccountPage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPaymentAccountPageEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentAccountPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentInitiationEntity

```php
$no_frixion_money_moov_models_payment_initiation = $client->NoFrixionMoneyMoovModelsPaymentInitiation();
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsPaymentInitiation()->create([
  "paymentrequest_id" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPaymentInitiationEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentInitiationEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestEventEntity

```php
$no_frixion_money_moov_models_payment_request_event = $client->NoFrixionMoneyMoovModelsPaymentRequestEvent();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float` | Yes |  |
| `apple_pay_transaction_id` | `string` | No |  |
| `card_authorization_response_id` | `string` | No |  |
| `card_expiry_month` | `int` | No |  |
| `card_expiry_year` | `int` | No |  |
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

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsPaymentRequestEvent()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPaymentRequestEventEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestEventEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestMetricEntity

```php
$no_frixion_money_moov_models_payment_request_metric = $client->NoFrixionMoneyMoovModelsPaymentRequestMetric();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `all` | `int` | No |  |
| `authorized` | `int` | No |  |
| `paid` | `int` | No |  |
| `partially_paid` | `int` | No |  |
| `total_amounts_by_currency` | `array` | No |  |
| `unpaid` | `int` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsPaymentRequestMetric()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPaymentRequestMetricEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMetricEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity

```php
$no_frixion_money_moov_models_payment_request_minimal = $client->NoFrixionMoneyMoovModelsPaymentRequestMinimal();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float` | No |  |
| `amount_pending` | `float` | No |  |
| `amount_received` | `float` | No |  |
| `amount_refunded` | `float` | No |  |
| `callback_url` | `string` | No |  |
| `card_stripe_payment_intent_secret` | `string` | No |  |
| `country_code` | `string` | No |  |
| `currency` | `string` | No |  |
| `custom_fields_to_display` | `array` | No |  |
| `description` | `string` | No |  |
| `due_date` | `string` | No |  |
| `field_display_setting` | `array` | No |  |
| `google_pay_merchant_id` | `string` | No |  |
| `id` | `string` | No |  |
| `jwk` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_logo_url_png` | `string` | No |  |
| `merchant_logo_url_svg` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `merchant_short_name` | `string` | No |  |
| `partial_payment_method` | `string` | No |  |
| `payment_attempt` | `array` | No |  |
| `payment_methods_list` | `array` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_processor_key` | `string` | No |  |
| `pisp_error` | `string` | No |  |
| `priority_bank_id` | `string` | No |  |
| `status` | `string` | No |  |
| `stripe_account_id` | `string` | No |  |
| `title` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsPaymentRequestMinimal()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestResultEntity

```php
$no_frixion_money_moov_models_payment_request_result = $client->NoFrixionMoneyMoovModelsPaymentRequestResult();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float` | No |  |
| `amount_pending` | `float` | No |  |
| `amount_received` | `float` | No |  |
| `amount_refunded` | `float` | No |  |
| `currency` | `string` | No |  |
| `customer_id` | `string` | No |  |
| `payment` | `array` | No |  |
| `payment_request_id` | `string` | No |  |
| `pisp_authorization` | `array` | No |  |
| `requested_amount` | `float` | No |  |
| `result` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsPaymentRequestResult()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPaymentRequestResultEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestResultEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity

```php
$no_frixion_money_moov_models_payment_requests_merchant_payment = $client->NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment();
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
| `template` | `array` | Yes |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity

```php
$no_frixion_money_moov_models_payment_requests_merchant_payment2 = $client->NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2();
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
| `template` | `array` | Yes |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2()->load(["paymentrequest_id" => "paymentrequest_id", "template_id" => "template_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity

```php
$no_frixion_money_moov_models_payment_requests_merchant_payment3 = $client->NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3();
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
| `template` | `array` | Yes |  |

### Operations

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3()->update([
  "paymentrequest_id" => "paymentrequest_id",
  "template_id" => "template_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity

```php
$no_frixion_money_moov_models_payment_requests_merchant_payment4 = $client->NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4();
```

### Operations

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4()->remove(["paymentrequest_id" => "paymentrequest_id", "template_id" => "template_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity`

Create a new `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayoutKeysetPageEntity

```php
$no_frixion_money_moov_models_payout_keyset_page = $client->NoFrixionMoneyMoovModelsPayoutKeysetPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `float` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `array` | No |  |
| `authorisation` | `array` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `array` | Yes |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `array` | No |  |
| `document` | `array` | No |  |
| `event` | `array` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `float` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `float` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `array` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `array` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `array` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `float` | No |  |
| `transacted_fx_amount` | `float` | No |  |
| `transacted_fx_rate` | `float` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsPayoutKeysetPage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPayoutKeysetPageEntity`

Create a new `NoFrixionMoneyMoovModelsPayoutKeysetPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayoutMetricEntity

```php
$no_frixion_money_moov_models_payout_metric = $client->NoFrixionMoneyMoovModelsPayoutMetric();
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
| `total_amounts_by_currency` | `array` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsPayoutMetric()->load();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPayoutMetricEntity`

Create a new `NoFrixionMoneyMoovModelsPayoutMetricEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity

```php
$no_frixion_money_moov_models_payouts_payouts_create = $client->NoFrixionMoneyMoovModelsPayoutsPayoutsCreate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `failed_payout` | `array` | No |  |
| `payout` | `array` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsPayoutsPayoutsCreate()->create([
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity`

Create a new `NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsPayrunEntity

```php
$no_frixion_money_moov_models_payrun = $client->NoFrixionMoneyMoovModelsPayrun();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation` | `array` | No |  |
| `authorisation_date` | `string` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_delete` | `bool` | No |  |
| `can_edit` | `bool` | No |  |
| `event` | `array` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice` | `array` | No |  |
| `invoices_minimal` | `array` | No |  |
| `is_archived` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `last_updated_by` | `array` | Yes |  |
| `merchant_id` | `string` | No |  |
| `name` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment` | `array` | No |  |
| `payout` | `array` | No |  |
| `payouts_count` | `int` | No |  |
| `reason` | `string` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled_date` | `string` | No |  |
| `source_account` | `array` | No |  |
| `status` | `string` | No |  |
| `total_eur` | `float` | No |  |
| `total_gbp` | `float` | No |  |
| `total_usd` | `float` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsPayrun()->create([
  "id" => null, // string
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsPayrun()->load(["id" => "no_frixion_money_moov_models_payrun_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsPayrun()->update([
  "id" => "no_frixion_money_moov_models_payrun_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsPayrunEntity`

Create a new `NoFrixionMoneyMoovModelsPayrunEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsReportResultEntity

```php
$no_frixion_money_moov_models_report_result = $client->NoFrixionMoneyMoovModelsReportResult();
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
| `statement_number` | `int` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsReportResult()->load(["id" => 1, "report_id" => "report_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsReportResultEntity`

Create a new `NoFrixionMoneyMoovModelsReportResultEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsRuleEntity

```php
$no_frixion_money_moov_models_rule = $client->NoFrixionMoneyMoovModelsRule();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account` | `array` | No |  |
| `account_id` | `string` | No |  |
| `approve_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `array` | No |  |
| `authorisation` | `array` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `can_authorise` | `bool` | No |  |
| `created_by` | `array` | Yes |  |
| `description` | `string` | No |  |
| `end_at` | `string` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `is_disabled` | `bool` | No |  |
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
| `sweep_action` | `array` | No |  |
| `time_zone_id` | `string` | No |  |
| `trigger_cron_expression` | `string` | No |  |
| `trigger_on_pay_in` | `bool` | No |  |
| `user_id` | `string` | No |  |
| `web_hook_secret` | `string` | No |  |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsRule()->create([
  "created_by" => null, // array
  "nonce" => null, // string
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsRule()->load(["id" => "no_frixion_money_moov_models_rule_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsRule()->update([
  "id" => "no_frixion_money_moov_models_rule_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsRuleEntity`

Create a new `NoFrixionMoneyMoovModelsRuleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsTransactionEntity

```php
$no_frixion_money_moov_models_transaction = $client->NoFrixionMoneyMoovModelsTransaction();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `account_name` | `string` | No |  |
| `account_sequence_number` | `int` | No |  |
| `amount` | `float` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `balance` | `float` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `counterparty` | `array` | No |  |
| `counterparty_summary` | `string` | No |  |
| `currency` | `string` | No |  |
| `description` | `string` | No |  |
| `fx_amount` | `float` | No |  |
| `fx_currency` | `string` | No |  |
| `fx_rate` | `float` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `payment_request_custom_field` | `array` | No |  |
| `payment_request_id` | `string` | No |  |
| `payout_id` | `string` | No |  |
| `raw_reference` | `string` | No |  |
| `rule_id` | `string` | No |  |
| `tag` | `array` | No |  |
| `their_reference` | `string` | No |  |
| `transaction_date` | `string` | No |  |
| `type` | `string` | No |  |
| `virtual_iban` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsTransaction()->load(["id" => "no_frixion_money_moov_models_transaction_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsTransactionEntity`

Create a new `NoFrixionMoneyMoovModelsTransactionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsTransactionPageEntity

```php
$no_frixion_money_moov_models_transaction_page = $client->NoFrixionMoneyMoovModelsTransactionPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `account_name` | `string` | No |  |
| `account_sequence_number` | `int` | No |  |
| `amount` | `float` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `balance` | `float` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `content` | `array` | No |  |
| `counterparty` | `array` | No |  |
| `counterparty_summary` | `string` | No |  |
| `currency` | `string` | No |  |
| `description` | `string` | No |  |
| `fx_amount` | `float` | No |  |
| `fx_currency` | `string` | No |  |
| `fx_rate` | `float` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `page_number` | `int` | No |  |
| `page_size` | `int` | No |  |
| `payment_request_custom_field` | `array` | No |  |
| `payment_request_id` | `string` | No |  |
| `payout_id` | `string` | No |  |
| `raw_reference` | `string` | No |  |
| `rule_id` | `string` | No |  |
| `tag` | `array` | No |  |
| `their_reference` | `string` | No |  |
| `total_page` | `int` | No |  |
| `total_size` | `int` | No |  |
| `transaction_date` | `string` | No |  |
| `type` | `string` | No |  |
| `virtual_iban` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsTransactionPage()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsTransactionPage()->load(["account_id" => "account_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsTransactionPageEntity`

Create a new `NoFrixionMoneyMoovModelsTransactionPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsUserInviteEntity

```php
$no_frixion_money_moov_models_user_invite = $client->NoFrixionMoneyMoovModelsUserInvite();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_status` | `array` | No |  |
| `id` | `string` | No |  |
| `initial_role_id` | `string` | No |  |
| `invitee_email_address` | `string` | No |  |
| `invitee_first_name` | `string` | No |  |
| `invitee_last_name` | `string` | No |  |
| `inviter_email_address` | `string` | No |  |
| `inviter_first_name` | `string` | No |  |
| `inviter_last_name` | `string` | No |  |
| `is_authorised` | `bool` | No |  |
| `is_invitee_registered` | `bool` | No |  |
| `last_invited` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `message` | `string` | No |  |
| `registration_url` | `string` | No |  |
| `send_invite_email` | `bool` | No |  |
| `status` | `string` | No |  |
| `user` | `array` | Yes |  |
| `user_id` | `string` | No |  |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsUserInvite()->create([
  "user" => null, // array
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsUserInvite()->load(["id" => "no_frixion_money_moov_models_user_invite_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsUserInviteEntity`

Create a new `NoFrixionMoneyMoovModelsUserInviteEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsUserInvitePageEntity

```php
$no_frixion_money_moov_models_user_invite_page = $client->NoFrixionMoneyMoovModelsUserInvitePage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `authorisation_status` | `array` | No |  |
| `id` | `string` | No |  |
| `initial_role_id` | `string` | No |  |
| `invitee_email_address` | `string` | No |  |
| `invitee_first_name` | `string` | No |  |
| `invitee_last_name` | `string` | No |  |
| `inviter_email_address` | `string` | No |  |
| `inviter_first_name` | `string` | No |  |
| `inviter_last_name` | `string` | No |  |
| `is_authorised` | `bool` | No |  |
| `is_invitee_registered` | `bool` | No |  |
| `last_invited` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `message` | `string` | No |  |
| `registration_url` | `string` | No |  |
| `status` | `string` | No |  |
| `user` | `array` | Yes |  |
| `user_id` | `string` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsUserInvitePage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsUserInvitePageEntity`

Create a new `NoFrixionMoneyMoovModelsUserInvitePageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsUserPageEntity

```php
$no_frixion_money_moov_models_user_page = $client->NoFrixionMoneyMoovModelsUserPage();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `array` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `bool` | No |  |
| `permission` | `array` | No |  |
| `roles_with_scope` | `array` | No |  |
| `two_factor_enabled` | `bool` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsUserPage()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsUserPageEntity`

Create a new `NoFrixionMoneyMoovModelsUserPageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## NoFrixionMoneyMoovModelsWebhookEntity

```php
$no_frixion_money_moov_models_webhook = $client->NoFrixionMoneyMoovModelsWebhook();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `destination_url` | `string` | No |  |
| `email_address` | `string` | No |  |
| `failed_notification_email_address` | `string` | No |  |
| `id` | `string` | No |  |
| `is_active` | `bool` | No |  |
| `merchant_id` | `string` | No |  |
| `notification_method` | `string` | No |  |
| `resource_type` | `array` | No |  |
| `retry` | `bool` | No |  |
| `secret` | `string` | No |  |
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsWebhook()->create([
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->NoFrixionMoneyMoovModelsWebhook()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsWebhook()->load(["id" => "no_frixion_money_moov_models_webhook_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->NoFrixionMoneyMoovModelsWebhook()->update([
  "id" => "no_frixion_money_moov_models_webhook_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): NoFrixionMoneyMoovModelsWebhookEntity`

Create a new `NoFrixionMoneyMoovModelsWebhookEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## OpenBankingEntity

```php
$open_banking = $client->OpenBanking();
```

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->OpenBanking()->create([
  "account_id" => null, // string
]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->OpenBanking()->remove();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): OpenBankingEntity`

Create a new `OpenBankingEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PayeeverificationEntity

```php
$payeeverification = $client->Payeeverification();
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Payeeverification()->create([
  "account_name" => null, // string
  "iban" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PayeeverificationEntity`

Create a new `PayeeverificationEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PaymentRequestEntity

```php
$payment_request = $client->PaymentRequest();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `float` | No |  |
| `do_simulate_settlement_failure` | `bool` | No |  |
| `error_description` | `string` | No |  |
| `institution` | `string` | No |  |
| `payment_initiation_id` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PaymentRequest()->create([
  "paymentrequest_id" => null, // string
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->PaymentRequest()->load();
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->PaymentRequest()->remove(["id" => "id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->PaymentRequest()->update([
  "paymentrequest_id" => "paymentrequest_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PaymentRequestEntity`

Create a new `PaymentRequestEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PayoutEntity

```php
$payout = $client->Payout();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `allow_incomplete` | `bool` | No |  |
| `amount` | `float` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `array` | No |  |
| `authorisation` | `array` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `array` | Yes |  |
| `beneficiary_id` | `string` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `array` | No |  |
| `document` | `array` | No |  |
| `event` | `array` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `float` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `float` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `array` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `array` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `array` | No |  |
| `tag_id` | `array` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `float` | No |  |
| `transacted_fx_amount` | `float` | No |  |
| `transacted_fx_rate` | `float` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Payout()->create([
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Payout()->load(["id" => "payout_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Payout()->remove(["id" => "payout_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Payout()->update([
  "id" => "payout_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PayoutEntity`

Create a new `PayoutEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PayrunEntity

```php
$payrun = $client->Payrun();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |
| `note` | `string` | No |  |
| `scheduled_date` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Payrun()->create([
  "id" => null, // string
]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Payrun()->remove(["id" => "id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Payrun()->update([
  "id" => "id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PayrunEntity`

Create a new `PayrunEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## RejectEntity

```php
$reject = $client->Reject();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `amount` | `float` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `array` | No |  |
| `authorisation` | `array` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `array` | Yes |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `array` | No |  |
| `document` | `array` | No |  |
| `event` | `array` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `float` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `float` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `reason` | `string` | No |  |
| `rule` | `array` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `array` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `array` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `float` | No |  |
| `transacted_fx_amount` | `float` | No |  |
| `transacted_fx_rate` | `float` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

### Operations

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Reject()->update([
  "id" => "id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): RejectEntity`

Create a new `RejectEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ReportEntity

```php
$report = $client->Report();
```

### Operations

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Report()->update([
  "id" => "id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ReportEntity`

Create a new `ReportEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## RuleEntity

```php
$rule = $client->Rule();
```

### Operations

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Rule()->remove(["id" => "id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Rule()->update([
  "id" => "id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): RuleEntity`

Create a new `RuleEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SendEntity

```php
$send = $client->Send();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `allow_incomplete` | `bool` | No |  |
| `amount` | `float` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `array` | No |  |
| `authorisation` | `array` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `array` | Yes |  |
| `beneficiary_id` | `string` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `array` | No |  |
| `document` | `array` | No |  |
| `event` | `array` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `float` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `float` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `array` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `array` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `array` | No |  |
| `tag_id` | `array` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `float` | No |  |
| `transacted_fx_amount` | `float` | No |  |
| `transacted_fx_rate` | `float` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Send()->create([
  "beneficiary" => null, // array
  "source_account_identifier" => null, // array
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SendEntity`

Create a new `SendEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SendbeneficiaryEntity

```php
$sendbeneficiary = $client->Sendbeneficiary();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_id` | `string` | No |  |
| `allow_incomplete` | `bool` | No |  |
| `amount` | `float` | No |  |
| `amount_minor_unit` | `int` | No |  |
| `approve_payout_url` | `string` | No |  |
| `approver_id` | `string` | No |  |
| `authentication_method` | `array` | No |  |
| `authorisation` | `array` | No |  |
| `authorisers_completed_count` | `int` | No |  |
| `authorisers_required_count` | `int` | No |  |
| `batch_payout_id` | `string` | No |  |
| `beneficiary` | `array` | Yes |  |
| `beneficiary_id` | `string` | No |  |
| `can_authorise` | `bool` | No |  |
| `can_process` | `bool` | No |  |
| `can_update` | `bool` | No |  |
| `charge_bearer` | `string` | No |  |
| `created_by` | `string` | No |  |
| `created_by_email_address` | `string` | No |  |
| `currency` | `string` | No |  |
| `current_user_id` | `string` | No |  |
| `description` | `string` | No |  |
| `destination` | `array` | No |  |
| `document` | `array` | No |  |
| `event` | `array` | No |  |
| `formatted_amount` | `string` | No |  |
| `formatted_fx_destination_amount` | `string` | No |  |
| `formatted_schedule` | `string` | No |  |
| `formatted_schedule_day_only` | `string` | No |  |
| `formatted_source_account_available_balance` | `string` | No |  |
| `fx_destination_amount` | `float` | No |  |
| `fx_destination_amount_minor_unit` | `int` | No |  |
| `fx_destination_currency` | `string` | No |  |
| `fx_quote_expires_at` | `string` | No |  |
| `fx_quote_id` | `string` | No |  |
| `fx_rate` | `float` | No |  |
| `fx_use_destination_amount` | `bool` | No |  |
| `has_current_user_authorised` | `bool` | No |  |
| `id` | `string` | No |  |
| `inserted` | `string` | No |  |
| `invoice_id` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_failed` | `bool` | No |  |
| `is_settled` | `bool` | No |  |
| `is_submitted` | `bool` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_token_description` | `string` | No |  |
| `nonce` | `string` | No |  |
| `payment_processor` | `string` | No |  |
| `payment_rail` | `string` | No |  |
| `payrun_id` | `string` | No |  |
| `payrun_name` | `string` | No |  |
| `rule` | `array` | No |  |
| `schedule_date` | `string` | No |  |
| `scheduled` | `bool` | No |  |
| `source_account_available_balance` | `float` | No |  |
| `source_account_available_balance_minor_unit` | `int` | No |  |
| `source_account_bic` | `string` | No |  |
| `source_account_currency` | `string` | No |  |
| `source_account_iban` | `string` | No |  |
| `source_account_identifier` | `array` | Yes |  |
| `source_account_name` | `string` | No |  |
| `source_account_number` | `string` | No |  |
| `source_account_sortcode` | `string` | No |  |
| `status` | `string` | No |  |
| `tag` | `array` | No |  |
| `tag_id` | `array` | No |  |
| `their_reference` | `string` | No |  |
| `topup_payrun_id` | `string` | No |  |
| `transacted_amount` | `float` | No |  |
| `transacted_fx_amount` | `float` | No |  |
| `transacted_fx_rate` | `float` | No |  |
| `type` | `string` | No |  |
| `user_id` | `string` | No |  |
| `your_reference` | `string` | No |  |

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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Sendbeneficiary()->create([
  "beneficiary" => null, // array
  "source_account_identifier" => null, // array
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SendbeneficiaryEntity`

Create a new `SendbeneficiaryEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TagEntity

```php
$tag = $client->Tag();
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Tag()->create([
  "merchant_id" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Tag()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TagEntity`

Create a new `TagEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TokenEntity

```php
$token = $client->Token();
```

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Token()->create([
  "id" => null, // string
]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Token()->remove(["id" => "id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TokenEntity`

Create a new `TokenEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TransactionEntity

```php
$transaction = $client->Transaction();
```

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Transaction()->create([
  "id" => null, // string
]);
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Transaction()->load(["sequence_number" => 1, "transaction_id" => "transaction_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Transaction()->remove(["id" => "id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TransactionEntity`

Create a new `TransactionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## UserEntity

```php
$user = $client->User();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `array` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `bool` | No |  |
| `permission` | `array` | No |  |
| `profile` | `string` | No |  |
| `roles_with_scope` | `array` | No |  |
| `two_factor_enabled` | `bool` | No |  |
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

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->User()->list();
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->User()->update([
  "id" => "id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): UserEntity`

Create a new `UserEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## UserInviteEntity

```php
$user_invite = $client->UserInvite();
```

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->UserInvite()->create([
  "id" => null, // string
]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->UserInvite()->remove(["id" => "id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->UserInvite()->update([
  "id" => "id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): UserInviteEntity`

Create a new `UserInviteEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## VirtualEntity

```php
$virtual = $client->Virtual();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `account_name` | `string` | No |  |
| `account_supplier_name` | `string` | No |  |
| `available_balance` | `float` | No |  |
| `available_balance_minor_unit` | `int` | No |  |
| `balance` | `float` | No |  |
| `balance_minor_unit` | `int` | No |  |
| `bank_name` | `string` | No |  |
| `consent_id` | `string` | No |  |
| `created_by` | `array` | Yes |  |
| `created_by_display_name` | `string` | No |  |
| `currency` | `string` | No |  |
| `default_payment_rail` | `string` | No |  |
| `display_name` | `string` | No |  |
| `expiry_date` | `string` | No |  |
| `external_account_icon` | `string` | No |  |
| `id` | `string` | No |  |
| `identifier` | `array` | Yes |  |
| `inserted` | `string` | No |  |
| `is_archived` | `bool` | No |  |
| `is_connected_account` | `bool` | No |  |
| `is_default` | `bool` | No |  |
| `is_trust_account` | `bool` | No |  |
| `is_virtual` | `bool` | No |  |
| `last_transaction` | `array` | No |  |
| `last_updated` | `string` | No |  |
| `merchant_id` | `string` | No |  |
| `merchant_name` | `string` | No |  |
| `name` | `string` | Yes |  |
| `physical_account_id` | `string` | No |  |
| `rule` | `array` | No |  |
| `submitted_payouts_balance` | `float` | No |  |
| `submitted_payouts_balance_minor_unit` | `int` | No |  |
| `summary` | `string` | No |  |
| `supplier_sepa_instant_status` | `string` | No |  |
| `xero_bank_feed_connection_status` | `string` | No |  |
| `xero_bank_feed_last_synced_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failed_at` | `string` | No |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` | No |  |
| `xero_bank_feed_sync_status` | `string` | No |  |
| `xero_unsynchronised_transactions_count` | `int` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Virtual()->create([
  "account_id" => null, // string
]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Virtual()->update([
  "account_id" => "account_id",
  "id" => "id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): VirtualEntity`

Create a new `VirtualEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## WebhookEntity

```php
$webhook = $client->Webhook();
```

### Operations

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Webhook()->remove(["id" => "id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): WebhookEntity`

Create a new `WebhookEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## WhoamiEntity

```php
$whoami = $client->Whoami();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `array` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `bool` | No |  |
| `permission` | `array` | No |  |
| `roles_with_scope` | `array` | No |  |
| `two_factor_enabled` | `bool` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Whoami()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): WhoamiEntity`

Create a new `WhoamiEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## WhoamitrustedappEntity

```php
$whoamitrustedapp = $client->Whoamitrustedapp();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `client_session_timeout` | `array` | No |  |
| `email_address` | `string` | Yes |  |
| `first_name` | `string` | Yes |  |
| `id` | `string` | No |  |
| `last_name` | `string` | Yes |  |
| `passkey_added` | `bool` | No |  |
| `permission` | `array` | No |  |
| `roles_with_scope` | `array` | No |  |
| `two_factor_enabled` | `bool` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Whoamitrustedapp()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): WhoamitrustedappEntity`

Create a new `WhoamitrustedappEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new NofrixionSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

