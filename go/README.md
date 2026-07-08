# Nofrixion Golang SDK



The Golang SDK for the Nofrixion API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.Account(nil)` — each with the same small set of operations (`List`, `Load`, `Create`, `Update`, `Remove`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/nofrixion-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/nofrixion-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/nofrixion-sdk/go=../nofrixion-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    "os"
    sdk "github.com/voxgig-sdk/nofrixion-sdk/go"
)

func main() {
    client := sdk.NewNofrixionSDK(map[string]any{
        "apikey": os.Getenv("NOFRIXION_APIKEY"),
    })

    // List account records — the value is the array of records itself.
    accounts, err := client.Account(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }
    for _, item := range accounts.([]any) {
        fmt.Println(item)
    }

    // Load a single account — the value is the loaded record.
    account, err := client.Account(nil).Load(map[string]any{"id": "example_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(account)

    // Create a account.
    created, err := client.Account(nil).Create(map[string]any{"account_id": "example_account_id", "currency": "example_currency"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(created)

    // Update a account.
    updated, err := client.Account(nil).Update(map[string]any{"id": "example_id", "account_id": "example_account_id", "amount": 1}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(updated)

    // Remove a account.
    removed, err := client.Account(nil).Remove(map[string]any{"id": "example_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(removed)
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
accounts, err := client.Account(nil).List(nil, nil)
if err != nil {
    // handle err
    return
}
_ = accounts
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

account, err := client.Account(nil).List(
    nil, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(account) // the returned mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewNofrixionSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
    },
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
cd go && go test ./test/...
```


## Reference

### NewNofrixionSDK

```go
func NewNofrixionSDK(options map[string]any) *NofrixionSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"apikey"` | `string` | API key for authentication. |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *NofrixionSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### NofrixionSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `Account` | `(data map[string]any) NofrixionEntity` | Create an Account entity instance. |
| `Beneficiary` | `(data map[string]any) NofrixionEntity` | Create a Beneficiary entity instance. |
| `Cancel` | `(data map[string]any) NofrixionEntity` | Create a Cancel entity instance. |
| `Disable` | `(data map[string]any) NofrixionEntity` | Create a Disable entity instance. |
| `Enable` | `(data map[string]any) NofrixionEntity` | Create an Enable entity instance. |
| `Merchant` | `(data map[string]any) NofrixionEntity` | Create a Merchant entity instance. |
| `Metadata` | `(data map[string]any) NofrixionEntity` | Create a Metadata entity instance. |
| `NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage entity instance. |
| `NoFrixionBizBizModelsPagingPaymentRequestPage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionBizBizModelsPagingPaymentRequestPage entity instance. |
| `NoFrixionBizBizModelsPagingPayoutPage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionBizBizModelsPagingPayoutPage entity instance. |
| `NoFrixionBizBizModelsPagingPayrunPage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionBizBizModelsPagingPayrunPage entity instance. |
| `NoFrixionBizBizModelsPagingRuleEventsPage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionBizBizModelsPagingRuleEventsPage entity instance. |
| `NoFrixionBizBizModelsPagingRulesPage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionBizBizModelsPagingRulesPage entity instance. |
| `NoFrixionBizBizModelsPaymentsCardPayment` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionBizBizModelsPaymentsCardPayment entity instance. |
| `NoFrixionBizBizModelsPaymentsCardPublicKey` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionBizBizModelsPaymentsCardPublicKey entity instance. |
| `NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries entity instance. |
| `NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment entity instance. |
| `NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate entity instance. |
| `NoFrixionMoneyMoovApiFeaturesUserInvitesCreate` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovApiFeaturesUserInvitesCreate entity instance. |
| `NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant entity instance. |
| `NoFrixionMoneyMoovModelsBatchPayout` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsBatchPayout entity instance. |
| `NoFrixionMoneyMoovModelsBeneficiaryGroupPage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsBeneficiaryGroupPage entity instance. |
| `NoFrixionMoneyMoovModelsBeneficiaryPage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsBeneficiaryPage entity instance. |
| `NoFrixionMoneyMoovModelsCardCustomerToken` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsCardCustomerToken entity instance. |
| `NoFrixionMoneyMoovModelsCurrencyCurrencyInfo` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsCurrencyCurrencyInfo entity instance. |
| `NoFrixionMoneyMoovModelsDirectDebitBatchSubmit` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsDirectDebitBatchSubmit entity instance. |
| `NoFrixionMoneyMoovModelsFxRate` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsFxRate entity instance. |
| `NoFrixionMoneyMoovModelsIPayment` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsIPayment entity instance. |
| `NoFrixionMoneyMoovModelsMandatesMandate` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsMandatesMandate entity instance. |
| `NoFrixionMoneyMoovModelsMerchant` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsMerchant entity instance. |
| `NoFrixionMoneyMoovModelsMerchantPage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsMerchantPage entity instance. |
| `NoFrixionMoneyMoovModelsMerchantPayByBankSetting` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsMerchantPayByBankSetting entity instance. |
| `NoFrixionMoneyMoovModelsMerchantToken` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsMerchantToken entity instance. |
| `NoFrixionMoneyMoovModelsMerchantTokenPage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsMerchantTokenPage entity instance. |
| `NoFrixionMoneyMoovModelsNoFrixionVersion` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsNoFrixionVersion entity instance. |
| `NoFrixionMoneyMoovModelsOpenBankingAccount` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsOpenBankingAccount entity instance. |
| `NoFrixionMoneyMoovModelsOpenBankingConsent` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsOpenBankingConsent entity instance. |
| `NoFrixionMoneyMoovModelsOpenBankingTransaction` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsOpenBankingTransaction entity instance. |
| `NoFrixionMoneyMoovModelsPayment` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPayment entity instance. |
| `NoFrixionMoneyMoovModelsPaymentAccountMinimalPage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPaymentAccountMinimalPage entity instance. |
| `NoFrixionMoneyMoovModelsPaymentAccountPage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPaymentAccountPage entity instance. |
| `NoFrixionMoneyMoovModelsPaymentInitiation` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPaymentInitiation entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestEvent` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPaymentRequestEvent entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestMetric` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPaymentRequestMetric entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestMinimal` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPaymentRequestMinimal entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestResult` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPaymentRequestResult entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2 entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3 entity instance. |
| `NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4 entity instance. |
| `NoFrixionMoneyMoovModelsPayoutKeysetPage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPayoutKeysetPage entity instance. |
| `NoFrixionMoneyMoovModelsPayoutMetric` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPayoutMetric entity instance. |
| `NoFrixionMoneyMoovModelsPayoutsPayoutsCreate` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPayoutsPayoutsCreate entity instance. |
| `NoFrixionMoneyMoovModelsPayrun` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsPayrun entity instance. |
| `NoFrixionMoneyMoovModelsReportResult` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsReportResult entity instance. |
| `NoFrixionMoneyMoovModelsRule` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsRule entity instance. |
| `NoFrixionMoneyMoovModelsTransaction` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsTransaction entity instance. |
| `NoFrixionMoneyMoovModelsTransactionPage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsTransactionPage entity instance. |
| `NoFrixionMoneyMoovModelsUserInvite` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsUserInvite entity instance. |
| `NoFrixionMoneyMoovModelsUserInvitePage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsUserInvitePage entity instance. |
| `NoFrixionMoneyMoovModelsUserPage` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsUserPage entity instance. |
| `NoFrixionMoneyMoovModelsWebhook` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionMoneyMoovModelsWebhook entity instance. |
| `OpenBanking` | `(data map[string]any) NofrixionEntity` | Create an OpenBanking entity instance. |
| `Payeeverification` | `(data map[string]any) NofrixionEntity` | Create a Payeeverification entity instance. |
| `PaymentRequest` | `(data map[string]any) NofrixionEntity` | Create a PaymentRequest entity instance. |
| `Payout` | `(data map[string]any) NofrixionEntity` | Create a Payout entity instance. |
| `Payrun` | `(data map[string]any) NofrixionEntity` | Create a Payrun entity instance. |
| `Reject` | `(data map[string]any) NofrixionEntity` | Create a Reject entity instance. |
| `Report` | `(data map[string]any) NofrixionEntity` | Create a Report entity instance. |
| `Rule` | `(data map[string]any) NofrixionEntity` | Create a Rule entity instance. |
| `Send` | `(data map[string]any) NofrixionEntity` | Create a Send entity instance. |
| `Sendbeneficiary` | `(data map[string]any) NofrixionEntity` | Create a Sendbeneficiary entity instance. |
| `Tag` | `(data map[string]any) NofrixionEntity` | Create a Tag entity instance. |
| `Token` | `(data map[string]any) NofrixionEntity` | Create a Token entity instance. |
| `Transaction` | `(data map[string]any) NofrixionEntity` | Create a Transaction entity instance. |
| `User` | `(data map[string]any) NofrixionEntity` | Create an User entity instance. |
| `UserInvite` | `(data map[string]any) NofrixionEntity` | Create an UserInvite entity instance. |
| `Virtual` | `(data map[string]any) NofrixionEntity` | Create a Virtual entity instance. |
| `Webhook` | `(data map[string]any) NofrixionEntity` | Create a Webhook entity instance. |
| `Whoami` | `(data map[string]any) NofrixionEntity` | Create a Whoami entity instance. |
| `Whoamitrustedapp` | `(data map[string]any) NofrixionEntity` | Create a Whoamitrustedapp entity instance. |

### Entity interface (NofrixionEntity)

All entities implement the `NofrixionEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Update` | `(reqdata, ctrl map[string]any) (any, error)` | Update an existing entity. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` / `Update` / `Remove` | the entity record (`map[string]any`) |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    account, err := client.Account(nil).List(map[string]any{/* fields */}, nil)
    if err != nil { /* handle */ }
    // account is the returned record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### Account

| Field | Description |
| --- | --- |
| `"account_id"` |  |
| `"account_name"` |  |
| `"account_supplier_name"` |  |
| `"account_type"` |  |
| `"available_balance"` |  |
| `"available_balance_minor_unit"` |  |
| `"balance"` |  |
| `"balance_minor_unit"` |  |
| `"bank_name"` |  |
| `"consent_id"` |  |
| `"created_by"` |  |
| `"created_by_display_name"` |  |
| `"currency"` |  |
| `"default_payment_rail"` |  |
| `"display_name"` |  |
| `"expiry_date"` |  |
| `"external_account_icon"` |  |
| `"format"` |  |
| `"from_date"` |  |
| `"id"` |  |
| `"identifier"` |  |
| `"inserted"` |  |
| `"is_archived"` |  |
| `"is_connected_account"` |  |
| `"is_default"` |  |
| `"is_trust_account"` |  |
| `"is_virtual"` |  |
| `"last_transaction"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"merchant_name"` |  |
| `"physical_account_id"` |  |
| `"role_i_d"` |  |
| `"rule"` |  |
| `"submitted_payouts_balance"` |  |
| `"submitted_payouts_balance_minor_unit"` |  |
| `"summary"` |  |
| `"supplier_physical_account_id"` |  |
| `"supplier_sepa_instant_status"` |  |
| `"to_date"` |  |
| `"xero_bank_feed_connection_status"` |  |
| `"xero_bank_feed_last_synced_at"` |  |
| `"xero_bank_feed_sync_last_failed_at"` |  |
| `"xero_bank_feed_sync_last_failure_reason"` |  |
| `"xero_bank_feed_sync_status"` |  |
| `"xero_unsynchronised_transactions_count"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/accounts/{accountID}/{currency}`

#### Beneficiary

| Field | Description |
| --- | --- |
| `"approval_callback_url"` |  |
| `"authentication_method"` |  |
| `"authorisation"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"beneficiary_event"` |  |
| `"can_authorise"` |  |
| `"can_update"` |  |
| `"created_by"` |  |
| `"created_by_email_address"` |  |
| `"currency"` |  |
| `"destination"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"is_enabled"` |  |
| `"last_authorised"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"name"` |  |
| `"nonce"` |  |
| `"source_account"` |  |
| `"source_account_i_d"` |  |
| `"their_reference"` |  |

Operations: Create, Load, Remove, Update.

API path: `/api/v1/beneficiaries/authorise/{id}`

#### Cancel

| Field | Description |
| --- | --- |
| `"account_id"` |  |
| `"amount"` |  |
| `"amount_minor_unit"` |  |
| `"approve_payout_url"` |  |
| `"approver_id"` |  |
| `"authentication_method"` |  |
| `"authorisation"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"batch_payout_id"` |  |
| `"beneficiary"` |  |
| `"can_authorise"` |  |
| `"can_process"` |  |
| `"can_update"` |  |
| `"charge_bearer"` |  |
| `"created_by"` |  |
| `"created_by_email_address"` |  |
| `"currency"` |  |
| `"current_user_id"` |  |
| `"description"` |  |
| `"destination"` |  |
| `"document"` |  |
| `"event"` |  |
| `"formatted_amount"` |  |
| `"formatted_fx_destination_amount"` |  |
| `"formatted_schedule"` |  |
| `"formatted_schedule_day_only"` |  |
| `"formatted_source_account_available_balance"` |  |
| `"fx_destination_amount"` |  |
| `"fx_destination_amount_minor_unit"` |  |
| `"fx_destination_currency"` |  |
| `"fx_quote_expires_at"` |  |
| `"fx_quote_id"` |  |
| `"fx_rate"` |  |
| `"fx_use_destination_amount"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"invoice_id"` |  |
| `"is_archived"` |  |
| `"is_failed"` |  |
| `"is_settled"` |  |
| `"is_submitted"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"merchant_token_description"` |  |
| `"nonce"` |  |
| `"payment_processor"` |  |
| `"payment_rail"` |  |
| `"payrun_id"` |  |
| `"payrun_name"` |  |
| `"rule"` |  |
| `"schedule_date"` |  |
| `"scheduled"` |  |
| `"source_account_available_balance"` |  |
| `"source_account_available_balance_minor_unit"` |  |
| `"source_account_bic"` |  |
| `"source_account_currency"` |  |
| `"source_account_iban"` |  |
| `"source_account_identifier"` |  |
| `"source_account_name"` |  |
| `"source_account_number"` |  |
| `"source_account_sortcode"` |  |
| `"status"` |  |
| `"tag"` |  |
| `"their_reference"` |  |
| `"topup_payrun_id"` |  |
| `"transacted_amount"` |  |
| `"transacted_fx_amount"` |  |
| `"transacted_fx_rate"` |  |
| `"type"` |  |
| `"user_id"` |  |
| `"your_reference"` |  |

Operations: Update.

API path: `/api/v1/payouts/cancel/{id}`

#### Disable

| Field | Description |
| --- | --- |
| `"approval_callback_url"` |  |
| `"authentication_method"` |  |
| `"authorisation"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"beneficiary_event"` |  |
| `"can_authorise"` |  |
| `"can_update"` |  |
| `"created_by"` |  |
| `"created_by_email_address"` |  |
| `"currency"` |  |
| `"destination"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"is_enabled"` |  |
| `"last_authorised"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"name"` |  |
| `"nonce"` |  |
| `"source_account"` |  |
| `"their_reference"` |  |

Operations: Update.

API path: `/api/v1/beneficiaries/disable/{id}`

#### Enable

| Field | Description |
| --- | --- |
| `"approval_callback_url"` |  |
| `"authentication_method"` |  |
| `"authorisation"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"beneficiary_event"` |  |
| `"can_authorise"` |  |
| `"can_update"` |  |
| `"created_by"` |  |
| `"created_by_email_address"` |  |
| `"currency"` |  |
| `"destination"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"is_enabled"` |  |
| `"last_authorised"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"name"` |  |
| `"nonce"` |  |
| `"source_account"` |  |
| `"their_reference"` |  |

Operations: Update.

API path: `/api/v1/beneficiaries/enable/{id}`

#### Merchant

| Field | Description |
| --- | --- |
| `"reason"` |  |

Operations: Load, Remove, Update.

API path: `/api/v1/merchants/{merchantID}/payouts/export`

#### Metadata

| Field | Description |
| --- | --- |

Operations: Load.

API path: `/api/v1/metadata/problemnotification`

#### NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage

| Field | Description |
| --- | --- |
| `"approved_at"` |  |
| `"currency"` |  |
| `"customer_account_number"` |  |
| `"customer_city"` |  |
| `"customer_country_code"` |  |
| `"customer_country_name"` |  |
| `"customer_email_address"` |  |
| `"customer_first_name"` |  |
| `"customer_iban"` |  |
| `"customer_last_name"` |  |
| `"customer_sort_code"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"is_recurring"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"reference"` |  |
| `"status"` |  |
| `"supplier_bank_account_id"` |  |
| `"supplier_customer_id"` |  |
| `"supplier_mandate_id"` |  |
| `"supplier_name"` |  |
| `"supplier_status"` |  |

Operations: List.

API path: `/api/v1/mandates`

#### NoFrixionBizBizModelsPagingPaymentRequestPage

| Field | Description |
| --- | --- |
| `"address"` |  |
| `"amount"` |  |
| `"amount_pending"` |  |
| `"amount_received"` |  |
| `"amount_refunded"` |  |
| `"auto_send_receipt"` |  |
| `"base_origin_url"` |  |
| `"callback_url"` |  |
| `"card_authorize_only"` |  |
| `"card_create_token"` |  |
| `"card_create_token_mode"` |  |
| `"card_ignore_cvn"` |  |
| `"card_processor_merchant_id"` |  |
| `"card_stripe_payment_intent_id"` |  |
| `"card_stripe_payment_intent_secret"` |  |
| `"created_by_user"` |  |
| `"currency"` |  |
| `"custom_field"` |  |
| `"customer_email_address"` |  |
| `"customer_id"` |  |
| `"customer_name"` |  |
| `"description"` |  |
| `"destination_account"` |  |
| `"direct_debit_payment"` |  |
| `"due_date"` |  |
| `"event"` |  |
| `"failure_callback_url"` |  |
| `"field_display_setting"` |  |
| `"formatted_amount"` |  |
| `"hosted_pay_checkout_url"` |  |
| `"id"` |  |
| `"ignore_address_verification"` |  |
| `"inserted"` |  |
| `"inserted_sortable"` |  |
| `"is_archived"` |  |
| `"jwk"` |  |
| `"last_updated"` |  |
| `"lightning_invoice"` |  |
| `"lightning_invoice_expires_at"` |  |
| `"merchant_direct_debit_mandate_id"` |  |
| `"merchant_id"` |  |
| `"merchant_token_description"` |  |
| `"notification_email_address"` |  |
| `"notification_role_i_d"` |  |
| `"order_id"` |  |
| `"partial_payment_method"` |  |
| `"partial_payment_step"` |  |
| `"payment_attempt"` |  |
| `"payment_method"` |  |
| `"payment_processor"` |  |
| `"payrun_id"` |  |
| `"pisp_account_id"` |  |
| `"priority_bank_id"` |  |
| `"result"` |  |
| `"sandbox_settle_delay_in_second"` |  |
| `"shipping_address"` |  |
| `"status"` |  |
| `"success_web_hook_url"` |  |
| `"tag"` |  |
| `"title"` |  |
| `"tokenised_card"` |  |
| `"transaction"` |  |
| `"use_hosted_payment_page"` |  |

Operations: List.

API path: `/api/v1/paymentrequests`

#### NoFrixionBizBizModelsPagingPayoutPage

| Field | Description |
| --- | --- |
| `"account_id"` |  |
| `"amount"` |  |
| `"amount_minor_unit"` |  |
| `"approve_payout_url"` |  |
| `"approver_id"` |  |
| `"authentication_method"` |  |
| `"authorisation"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"batch_payout_id"` |  |
| `"beneficiary"` |  |
| `"can_authorise"` |  |
| `"can_process"` |  |
| `"can_update"` |  |
| `"charge_bearer"` |  |
| `"created_by"` |  |
| `"created_by_email_address"` |  |
| `"currency"` |  |
| `"current_user_id"` |  |
| `"description"` |  |
| `"destination"` |  |
| `"document"` |  |
| `"event"` |  |
| `"formatted_amount"` |  |
| `"formatted_fx_destination_amount"` |  |
| `"formatted_schedule"` |  |
| `"formatted_schedule_day_only"` |  |
| `"formatted_source_account_available_balance"` |  |
| `"fx_destination_amount"` |  |
| `"fx_destination_amount_minor_unit"` |  |
| `"fx_destination_currency"` |  |
| `"fx_quote_expires_at"` |  |
| `"fx_quote_id"` |  |
| `"fx_rate"` |  |
| `"fx_use_destination_amount"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"invoice_id"` |  |
| `"is_archived"` |  |
| `"is_failed"` |  |
| `"is_settled"` |  |
| `"is_submitted"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"merchant_token_description"` |  |
| `"nonce"` |  |
| `"payment_processor"` |  |
| `"payment_rail"` |  |
| `"payrun_id"` |  |
| `"payrun_name"` |  |
| `"rule"` |  |
| `"schedule_date"` |  |
| `"scheduled"` |  |
| `"source_account_available_balance"` |  |
| `"source_account_available_balance_minor_unit"` |  |
| `"source_account_bic"` |  |
| `"source_account_currency"` |  |
| `"source_account_iban"` |  |
| `"source_account_identifier"` |  |
| `"source_account_name"` |  |
| `"source_account_number"` |  |
| `"source_account_sortcode"` |  |
| `"status"` |  |
| `"tag"` |  |
| `"their_reference"` |  |
| `"topup_payrun_id"` |  |
| `"transacted_amount"` |  |
| `"transacted_fx_amount"` |  |
| `"transacted_fx_rate"` |  |
| `"type"` |  |
| `"user_id"` |  |
| `"your_reference"` |  |

Operations: List.

API path: `/api/v1/payouts`

#### NoFrixionBizBizModelsPagingPayrunPage

| Field | Description |
| --- | --- |
| `"authorisation"` |  |
| `"authorisation_date"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"batch_payout_id"` |  |
| `"can_authorise"` |  |
| `"can_delete"` |  |
| `"can_edit"` |  |
| `"event"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"invoice"` |  |
| `"invoices_minimal"` |  |
| `"is_archived"` |  |
| `"last_updated"` |  |
| `"last_updated_by"` |  |
| `"merchant_id"` |  |
| `"name"` |  |
| `"nonce"` |  |
| `"payment"` |  |
| `"payout"` |  |
| `"payouts_count"` |  |
| `"schedule_date"` |  |
| `"source_account"` |  |
| `"status"` |  |
| `"total_eur"` |  |
| `"total_gbp"` |  |
| `"total_usd"` |  |

Operations: List.

API path: `/api/v1/payruns`

#### NoFrixionBizBizModelsPagingRuleEventsPage

| Field | Description |
| --- | --- |
| `"error_message"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"is_authorise_to_enable"` |  |
| `"message"` |  |
| `"raw_response"` |  |
| `"rule_event_type"` |  |
| `"rule_id"` |  |
| `"user"` |  |

Operations: List.

API path: `/api/v1/rules/{id}/events`

#### NoFrixionBizBizModelsPagingRulesPage

| Field | Description |
| --- | --- |
| `"account"` |  |
| `"account_id"` |  |
| `"approve_url"` |  |
| `"approver_id"` |  |
| `"authentication_method"` |  |
| `"authorisation"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"can_authorise"` |  |
| `"created_by"` |  |
| `"description"` |  |
| `"end_at"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"is_disabled"` |  |
| `"last_executed_at"` |  |
| `"last_run_at_transaction_date"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"name"` |  |
| `"nonce"` |  |
| `"on_approved_web_hook_url"` |  |
| `"on_execution_error_web_hook_url"` |  |
| `"on_execution_success_web_hook_url"` |  |
| `"start_at"` |  |
| `"status"` |  |
| `"sweep_action"` |  |
| `"time_zone_id"` |  |
| `"trigger_cron_expression"` |  |
| `"trigger_on_pay_in"` |  |
| `"user_id"` |  |
| `"web_hook_secret"` |  |

Operations: List.

API path: `/api/v1/rules`

#### NoFrixionBizBizModelsPaymentsCardPayment

| Field | Description |
| --- | --- |
| `"authorized_amount"` |  |
| `"currency_code"` |  |
| `"is_payer_authentication_required"` |  |
| `"is_soft_decline"` |  |
| `"payer_authentication_access_token"` |  |
| `"payer_authentication_merchant_data"` |  |
| `"payer_authentication_url"` |  |
| `"payer_authentication_window_height"` |  |
| `"payer_authentication_window_width"` |  |
| `"payment_request_callback_url"` |  |
| `"payment_request_id"` |  |
| `"request_id"` |  |
| `"response_code"` |  |
| `"response_type"` |  |
| `"status"` |  |
| `"three_ds_redirect_url"` |  |
| `"transaction_id"` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/{id}/card/refund/{partialRefundAmount}`

#### NoFrixionBizBizModelsPaymentsCardPublicKey

| Field | Description |
| --- | --- |
| `"jwt"` |  |

Operations: Load.

API path: `/api/v1/paymentrequests/{id}/card/publickey`

#### NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries

| Field | Description |
| --- | --- |
| `"beneficiary"` |  |
| `"failed_beneficiary"` |  |

Operations: Create.

API path: `/api/v1/beneficiaries/batchcreate`

#### NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment

| Field | Description |
| --- | --- |
| `"failed_payment_request"` |  |
| `"payment_request"` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/batchcreate`

#### NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate

| Field | Description |
| --- | --- |
| `"failed_role"` |  |
| `"role"` |  |

Operations: Create.

API path: `/api/v1/merchants/{merchantID}/roles/batchcreate`

#### NoFrixionMoneyMoovApiFeaturesUserInvitesCreate

| Field | Description |
| --- | --- |
| `"failed_user_invite"` |  |
| `"user_invite"` |  |

Operations: Create.

API path: `/api/v1/userinvites/batchcreate`

#### NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant

| Field | Description |
| --- | --- |
| `"amount_lower"` |  |
| `"amount_upper"` |  |
| `"authorisation_type"` |  |
| `"beneficiaries_only"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"last_editor_cant_authorise"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"number_of_authoriser"` |  |
| `"role_setting"` |  |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/authorisationsettings`

#### NoFrixionMoneyMoovModelsBatchPayout

| Field | Description |
| --- | --- |
| `"approve_url"` |  |
| `"id"` |  |
| `"payout"` |  |

Operations: Create, Load.

API path: `/api/v1/payouts/batch`

#### NoFrixionMoneyMoovModelsBeneficiaryGroupPage

| Field | Description |
| --- | --- |
| `"group_member"` |  |
| `"group_name"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/beneficiarygroups`

#### NoFrixionMoneyMoovModelsBeneficiaryPage

| Field | Description |
| --- | --- |
| `"approval_callback_url"` |  |
| `"authentication_method"` |  |
| `"authorisation"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"beneficiary_event"` |  |
| `"can_authorise"` |  |
| `"can_update"` |  |
| `"created_by"` |  |
| `"created_by_email_address"` |  |
| `"currency"` |  |
| `"destination"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"is_enabled"` |  |
| `"last_authorised"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"name"` |  |
| `"nonce"` |  |
| `"source_account"` |  |
| `"their_reference"` |  |

Operations: List.

API path: `/api/v1/beneficiaries`

#### NoFrixionMoneyMoovModelsCardCustomerToken

| Field | Description |
| --- | --- |
| `"card_type"` |  |
| `"customer_email_address"` |  |
| `"expiry_month"` |  |
| `"expiry_year"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"last_four_digit"` |  |
| `"last_updated"` |  |
| `"masked_card_number"` |  |
| `"merchant_id"` |  |
| `"payment_request_id"` |  |

Operations: List, Load, Remove.

API path: `/api/v1/paymentrequests/card/customertokens/{merchantID}/{customerEmailAddress}`

#### NoFrixionMoneyMoovModelsCurrencyCurrencyInfo

| Field | Description |
| --- | --- |
| `"code"` |  |
| `"decimal"` |  |
| `"is_fiat"` |  |
| `"iso4217_alpha_code"` |  |
| `"iso4217_numeric_code"` |  |
| `"symbol"` |  |

Operations: List.

API path: `/api/v1/currencies`

#### NoFrixionMoneyMoovModelsDirectDebitBatchSubmit

| Field | Description |
| --- | --- |
| `"failed_submission"` |  |
| `"successful_submission"` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/directdebit/batchsubmit`

#### NoFrixionMoneyMoovModelsFxRate

| Field | Description |
| --- | --- |
| `"destination_currency"` |  |
| `"exchange_rate"` |  |
| `"expiry_time"` |  |
| `"quote_id"` |  |
| `"source_currency"` |  |

Operations: List, Load.

API path: `/api/v1/payouts/fxallheldrates/{source}/{destination}`

#### NoFrixionMoneyMoovModelsIPayment

| Field | Description |
| --- | --- |
| `"payment_request_id"` |  |
| `"response_type"` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/payondemand`

#### NoFrixionMoneyMoovModelsMandatesMandate

| Field | Description |
| --- | --- |
| `"account_number"` |  |
| `"address_line1"` |  |
| `"address_line2"` |  |
| `"approved_at"` |  |
| `"city"` |  |
| `"country_code"` |  |
| `"currency"` |  |
| `"customer_account_number"` |  |
| `"customer_city"` |  |
| `"customer_country_code"` |  |
| `"customer_country_name"` |  |
| `"customer_email_address"` |  |
| `"customer_first_name"` |  |
| `"customer_iban"` |  |
| `"customer_last_name"` |  |
| `"customer_sort_code"` |  |
| `"email_address"` |  |
| `"first_name"` |  |
| `"iban"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"is_recurring"` |  |
| `"last_name"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"postal_code"` |  |
| `"reference"` |  |
| `"sort_code"` |  |
| `"status"` |  |
| `"supplier_bank_account_id"` |  |
| `"supplier_customer_id"` |  |
| `"supplier_mandate_id"` |  |
| `"supplier_name"` |  |
| `"supplier_status"` |  |

Operations: Create, Load.

API path: `/api/v1/mandates`

#### NoFrixionMoneyMoovModelsMerchant

| Field | Description |
| --- | --- |
| `"account_currency"` |  |
| `"can_have_trust_account"` |  |
| `"card_payment_processor"` |  |
| `"company_id"` |  |
| `"display_qr_on_hosted_pay"` |  |
| `"hosted_pay_version"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"is_blocked"` |  |
| `"is_exited"` |  |
| `"is_suspended"` |  |
| `"jurisdiction"` |  |
| `"logo_url_png"` |  |
| `"logo_url_svg"` |  |
| `"merchant_category_code"` |  |
| `"name"` |  |
| `"note"` |  |
| `"parent_merchant"` |  |
| `"payment_account"` |  |
| `"payment_account_limit"` |  |
| `"short_name"` |  |
| `"supported_payment_methods_list"` |  |
| `"suspension_reason"` |  |
| `"tag"` |  |
| `"time_zone_id"` |  |
| `"trading_name"` |  |
| `"web_hook_limit"` |  |
| `"your_role_name"` |  |

Operations: List, Load, Update.

API path: `/api/v1/merchants`

#### NoFrixionMoneyMoovModelsMerchantPage

| Field | Description |
| --- | --- |
| `"account_currency"` |  |
| `"can_have_trust_account"` |  |
| `"card_payment_processor"` |  |
| `"company_id"` |  |
| `"display_qr_on_hosted_pay"` |  |
| `"hosted_pay_version"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"is_blocked"` |  |
| `"is_exited"` |  |
| `"is_suspended"` |  |
| `"jurisdiction"` |  |
| `"logo_url_png"` |  |
| `"logo_url_svg"` |  |
| `"merchant_category_code"` |  |
| `"name"` |  |
| `"note"` |  |
| `"parent_merchant"` |  |
| `"payment_account"` |  |
| `"payment_account_limit"` |  |
| `"short_name"` |  |
| `"supported_payment_methods_list"` |  |
| `"suspension_reason"` |  |
| `"tag"` |  |
| `"time_zone_id"` |  |
| `"trading_name"` |  |
| `"web_hook_limit"` |  |
| `"your_role_name"` |  |

Operations: List.

API path: `/api/v1/merchants/paged`

#### NoFrixionMoneyMoovModelsMerchantPayByBankSetting

| Field | Description |
| --- | --- |
| `"bank_country_code"` |  |
| `"bank_id"` |  |
| `"bank_name"` |  |
| `"business_institution_id"` |  |
| `"currency"` |  |
| `"logo"` |  |
| `"message"` |  |
| `"message_image_url"` |  |
| `"order"` |  |
| `"personal_institution_id"` |  |
| `"processor"` |  |
| `"warning_heading"` |  |
| `"warning_message"` |  |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/banksettings`

#### NoFrixionMoneyMoovModelsMerchantToken

| Field | Description |
| --- | --- |
| `"authentication_method"` |  |
| `"authorisation"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"can_authorise"` |  |
| `"description"` |  |
| `"expires_at"` |  |
| `"has_current_user_authorised"` |  |
| `"hmac_algorithm"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"ip_address_whitelist"` |  |
| `"is_archived"` |  |
| `"is_enabled"` |  |
| `"last_authorised"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"nonce"` |  |
| `"permission_type"` |  |
| `"request_signature_version"` |  |
| `"shared_secret_algorithm"` |  |
| `"shared_secret_base64"` |  |
| `"token"` |  |

Operations: Create, Load, Update.

API path: `/api/v1/tokens`

#### NoFrixionMoneyMoovModelsMerchantTokenPage

| Field | Description |
| --- | --- |
| `"authentication_method"` |  |
| `"authorisation"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"can_authorise"` |  |
| `"description"` |  |
| `"expires_at"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"ip_address_whitelist"` |  |
| `"is_archived"` |  |
| `"is_enabled"` |  |
| `"last_authorised"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"nonce"` |  |
| `"permission_type"` |  |
| `"request_signature_version"` |  |
| `"shared_secret_algorithm"` |  |
| `"shared_secret_base64"` |  |
| `"token"` |  |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/tokens`

#### NoFrixionMoneyMoovModelsNoFrixionVersion

| Field | Description |
| --- | --- |
| `"build_version"` |  |
| `"major_version"` |  |
| `"minor_version"` |  |
| `"release_name"` |  |

Operations: Load.

API path: `/api/v1/metadata/version`

#### NoFrixionMoneyMoovModelsOpenBankingAccount

| Field | Description |
| --- | --- |
| `"account_balance"` |  |
| `"account_identification"` |  |
| `"account_name"` |  |
| `"account_type"` |  |
| `"balance"` |  |
| `"consolidated_account_information"` |  |
| `"currency"` |  |
| `"description"` |  |
| `"detail"` |  |
| `"id"` |  |
| `"nickname"` |  |
| `"type"` |  |
| `"usage_type"` |  |

Operations: Load.

API path: `/api/v1/openbanking/accounts/{id}`

#### NoFrixionMoneyMoovModelsOpenBankingConsent

| Field | Description |
| --- | --- |
| `"authorisation_url"` |  |
| `"callback_url"` |  |
| `"consent_id"` |  |
| `"email_address"` |  |
| `"expiry_date"` |  |
| `"failure_callback_url"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"institution_id"` |  |
| `"is_connected_account"` |  |
| `"is_enabled"` |  |
| `"merchant_id"` |  |
| `"provider"` |  |
| `"success_web_hook_url"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/openbanking/consents`

#### NoFrixionMoneyMoovModelsOpenBankingTransaction

| Field | Description |
| --- | --- |
| `"address_detail"` |  |
| `"amount"` |  |
| `"balance"` |  |
| `"booking_date_time"` |  |
| `"charge_detail"` |  |
| `"currency"` |  |
| `"currency_exchange"` |  |
| `"date"` |  |
| `"description"` |  |
| `"enrichment"` |  |
| `"gross_amount"` |  |
| `"id"` |  |
| `"iso_bank_transaction_code"` |  |
| `"merchant"` |  |
| `"payee_detail"` |  |
| `"payer_detail"` |  |
| `"proprietary_bank_transaction_code"` |  |
| `"reference"` |  |
| `"statement_reference"` |  |
| `"status"` |  |
| `"supplementary_data"` |  |
| `"transaction_amount"` |  |
| `"transaction_information"` |  |
| `"transaction_mutability"` |  |
| `"value_date_time"` |  |

Operations: List.

API path: `/api/v1/openbanking/transactions/{id}/{accountID}`

#### NoFrixionMoneyMoovModelsPayment

| Field | Description |
| --- | --- |
| `"address"` |  |
| `"amount"` |  |
| `"amount_pending"` |  |
| `"amount_received"` |  |
| `"amount_refunded"` |  |
| `"auto_send_receipt"` |  |
| `"base_origin_url"` |  |
| `"callback_url"` |  |
| `"card_authorize_only"` |  |
| `"card_create_token"` |  |
| `"card_create_token_mode"` |  |
| `"card_ignore_cvn"` |  |
| `"card_no_payer_authentication"` |  |
| `"card_processor_merchant_id"` |  |
| `"card_stripe_payment_intent_id"` |  |
| `"card_stripe_payment_intent_secret"` |  |
| `"card_transmit_raw_detail"` |  |
| `"created_by_user"` |  |
| `"currency"` |  |
| `"custom_field"` |  |
| `"customer_email_address"` |  |
| `"customer_id"` |  |
| `"customer_name"` |  |
| `"description"` |  |
| `"destination_account"` |  |
| `"direct_debit_payment"` |  |
| `"due_date"` |  |
| `"event"` |  |
| `"failure_callback_url"` |  |
| `"field_display_setting"` |  |
| `"formatted_amount"` |  |
| `"hosted_pay_checkout_url"` |  |
| `"id"` |  |
| `"ignore_address_verification"` |  |
| `"inserted"` |  |
| `"inserted_sortable"` |  |
| `"is_archived"` |  |
| `"jwk"` |  |
| `"last_updated"` |  |
| `"lightning_invoice"` |  |
| `"lightning_invoice_expires_at"` |  |
| `"merchant_direct_debit_mandate_id"` |  |
| `"merchant_id"` |  |
| `"merchant_token_description"` |  |
| `"notification_email_address"` |  |
| `"notification_role_i_d"` |  |
| `"order_id"` |  |
| `"partial_payment_method"` |  |
| `"partial_payment_step"` |  |
| `"payment_attempt"` |  |
| `"payment_method"` |  |
| `"payment_processor"` |  |
| `"payrun_id"` |  |
| `"pisp_account_id"` |  |
| `"priority_bank_id"` |  |
| `"result"` |  |
| `"sandbox_settle_delay_in_second"` |  |
| `"shipping_address"` |  |
| `"shipping_address_city"` |  |
| `"shipping_address_country_code"` |  |
| `"shipping_address_county"` |  |
| `"shipping_address_line1"` |  |
| `"shipping_address_line2"` |  |
| `"shipping_address_post_code"` |  |
| `"shipping_email"` |  |
| `"shipping_first_name"` |  |
| `"shipping_last_name"` |  |
| `"shipping_phone"` |  |
| `"status"` |  |
| `"success_web_hook_url"` |  |
| `"tag"` |  |
| `"tag_id"` |  |
| `"title"` |  |
| `"tokenised_card"` |  |
| `"transaction"` |  |
| `"use_hosted_payment_page"` |  |

Operations: Create, Load, Update.

API path: `/api/v1/paymentrequests`

#### NoFrixionMoneyMoovModelsPaymentAccountMinimalPage

| Field | Description |
| --- | --- |
| `"account_name"` |  |
| `"available_balance"` |  |
| `"balance"` |  |
| `"balance_minor_unit"` |  |
| `"currency"` |  |
| `"id"` |  |
| `"identifier"` |  |
| `"is_archived"` |  |
| `"is_connected_account"` |  |
| `"merchant_id"` |  |
| `"submitted_payouts_balance"` |  |

Operations: List.

API path: `/api/v1/accounts/minimal`

#### NoFrixionMoneyMoovModelsPaymentAccountPage

| Field | Description |
| --- | --- |
| `"account_name"` |  |
| `"account_supplier_name"` |  |
| `"available_balance"` |  |
| `"available_balance_minor_unit"` |  |
| `"balance"` |  |
| `"balance_minor_unit"` |  |
| `"bank_name"` |  |
| `"consent_id"` |  |
| `"created_by"` |  |
| `"created_by_display_name"` |  |
| `"currency"` |  |
| `"default_payment_rail"` |  |
| `"display_name"` |  |
| `"expiry_date"` |  |
| `"external_account_icon"` |  |
| `"id"` |  |
| `"identifier"` |  |
| `"inserted"` |  |
| `"is_archived"` |  |
| `"is_connected_account"` |  |
| `"is_default"` |  |
| `"is_trust_account"` |  |
| `"is_virtual"` |  |
| `"last_transaction"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"merchant_name"` |  |
| `"physical_account_id"` |  |
| `"rule"` |  |
| `"submitted_payouts_balance"` |  |
| `"submitted_payouts_balance_minor_unit"` |  |
| `"summary"` |  |
| `"supplier_sepa_instant_status"` |  |
| `"xero_bank_feed_connection_status"` |  |
| `"xero_bank_feed_last_synced_at"` |  |
| `"xero_bank_feed_sync_last_failed_at"` |  |
| `"xero_bank_feed_sync_last_failure_reason"` |  |
| `"xero_bank_feed_sync_status"` |  |
| `"xero_unsynchronised_transactions_count"` |  |

Operations: List.

API path: `/api/v1/accounts/paged`

#### NoFrixionMoneyMoovModelsPaymentInitiation

| Field | Description |
| --- | --- |
| `"payment_initiation_id"` |  |
| `"payment_request_callback_url"` |  |
| `"payment_request_id"` |  |
| `"redirect_url"` |  |
| `"response_type"` |  |
| `"specific_error_message"` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/{id}/pisp`

#### NoFrixionMoneyMoovModelsPaymentRequestEvent

| Field | Description |
| --- | --- |
| `"amount"` |  |
| `"apple_pay_transaction_id"` |  |
| `"card_authorization_response_id"` |  |
| `"card_expiry_month"` |  |
| `"card_expiry_year"` |  |
| `"card_issuer"` |  |
| `"card_issuer_country"` |  |
| `"card_last_four_digit"` |  |
| `"card_request_id"` |  |
| `"card_scheme"` |  |
| `"card_token_customer_id"` |  |
| `"card_transaction_id"` |  |
| `"currency"` |  |
| `"direct_debit_payment_id"` |  |
| `"direct_debit_payment_reference"` |  |
| `"drirect_debit_mandate_id"` |  |
| `"error_message"` |  |
| `"error_reason"` |  |
| `"event_type"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"lightning_invoice"` |  |
| `"lightning_r_hash"` |  |
| `"origin_url"` |  |
| `"payment_method_type"` |  |
| `"payment_processor_name"` |  |
| `"payment_request_id"` |  |
| `"pisp_bank_status"` |  |
| `"pisp_payment_initiation_id"` |  |
| `"pisp_payment_institution_name"` |  |
| `"pisp_payment_service_provider_id"` |  |
| `"pisp_redirect_url"` |  |
| `"reconciled_transaction_id"` |  |
| `"refund_payout_id"` |  |
| `"status"` |  |
| `"wallet_name"` |  |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/events`

#### NoFrixionMoneyMoovModelsPaymentRequestMetric

| Field | Description |
| --- | --- |
| `"all"` |  |
| `"authorized"` |  |
| `"paid"` |  |
| `"partially_paid"` |  |
| `"total_amounts_by_currency"` |  |
| `"unpaid"` |  |

Operations: Load.

API path: `/api/v1/paymentrequests/metrics`

#### NoFrixionMoneyMoovModelsPaymentRequestMinimal

| Field | Description |
| --- | --- |
| `"amount"` |  |
| `"amount_pending"` |  |
| `"amount_received"` |  |
| `"amount_refunded"` |  |
| `"callback_url"` |  |
| `"card_stripe_payment_intent_secret"` |  |
| `"country_code"` |  |
| `"currency"` |  |
| `"custom_fields_to_display"` |  |
| `"description"` |  |
| `"due_date"` |  |
| `"field_display_setting"` |  |
| `"google_pay_merchant_id"` |  |
| `"id"` |  |
| `"jwk"` |  |
| `"merchant_id"` |  |
| `"merchant_logo_url_png"` |  |
| `"merchant_logo_url_svg"` |  |
| `"merchant_name"` |  |
| `"merchant_short_name"` |  |
| `"partial_payment_method"` |  |
| `"payment_attempt"` |  |
| `"payment_methods_list"` |  |
| `"payment_processor"` |  |
| `"payment_processor_key"` |  |
| `"pisp_error"` |  |
| `"priority_bank_id"` |  |
| `"status"` |  |
| `"stripe_account_id"` |  |
| `"title"` |  |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/minimal`

#### NoFrixionMoneyMoovModelsPaymentRequestResult

| Field | Description |
| --- | --- |
| `"amount"` |  |
| `"amount_pending"` |  |
| `"amount_received"` |  |
| `"amount_refunded"` |  |
| `"currency"` |  |
| `"customer_id"` |  |
| `"payment"` |  |
| `"payment_request_id"` |  |
| `"pisp_authorization"` |  |
| `"requested_amount"` |  |
| `"result"` |  |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/result`

#### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment

| Field | Description |
| --- | --- |
| `"description"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"name"` |  |
| `"template"` |  |

Operations: List.

API path: `/api/v1/paymentrequests/{merchantID}/templates`

#### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2

| Field | Description |
| --- | --- |
| `"description"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"name"` |  |
| `"template"` |  |

Operations: Load.

API path: `/api/v1/paymentrequests/{merchantID}/templates/{templateID}`

#### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3

| Field | Description |
| --- | --- |
| `"description"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"name"` |  |
| `"template"` |  |

Operations: Update.

API path: `/api/v1/paymentrequests/{merchantID}/templates/{templateID}`

#### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4

| Field | Description |
| --- | --- |

Operations: Remove.

API path: `/api/v1/paymentrequests/{merchantID}/templates/{templateID}`

#### NoFrixionMoneyMoovModelsPayoutKeysetPage

| Field | Description |
| --- | --- |
| `"account_id"` |  |
| `"amount"` |  |
| `"amount_minor_unit"` |  |
| `"approve_payout_url"` |  |
| `"approver_id"` |  |
| `"authentication_method"` |  |
| `"authorisation"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"batch_payout_id"` |  |
| `"beneficiary"` |  |
| `"can_authorise"` |  |
| `"can_process"` |  |
| `"can_update"` |  |
| `"charge_bearer"` |  |
| `"created_by"` |  |
| `"created_by_email_address"` |  |
| `"currency"` |  |
| `"current_user_id"` |  |
| `"description"` |  |
| `"destination"` |  |
| `"document"` |  |
| `"event"` |  |
| `"formatted_amount"` |  |
| `"formatted_fx_destination_amount"` |  |
| `"formatted_schedule"` |  |
| `"formatted_schedule_day_only"` |  |
| `"formatted_source_account_available_balance"` |  |
| `"fx_destination_amount"` |  |
| `"fx_destination_amount_minor_unit"` |  |
| `"fx_destination_currency"` |  |
| `"fx_quote_expires_at"` |  |
| `"fx_quote_id"` |  |
| `"fx_rate"` |  |
| `"fx_use_destination_amount"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"invoice_id"` |  |
| `"is_archived"` |  |
| `"is_failed"` |  |
| `"is_settled"` |  |
| `"is_submitted"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"merchant_token_description"` |  |
| `"nonce"` |  |
| `"payment_processor"` |  |
| `"payment_rail"` |  |
| `"payrun_id"` |  |
| `"payrun_name"` |  |
| `"rule"` |  |
| `"schedule_date"` |  |
| `"scheduled"` |  |
| `"source_account_available_balance"` |  |
| `"source_account_available_balance_minor_unit"` |  |
| `"source_account_bic"` |  |
| `"source_account_currency"` |  |
| `"source_account_iban"` |  |
| `"source_account_identifier"` |  |
| `"source_account_name"` |  |
| `"source_account_number"` |  |
| `"source_account_sortcode"` |  |
| `"status"` |  |
| `"tag"` |  |
| `"their_reference"` |  |
| `"topup_payrun_id"` |  |
| `"transacted_amount"` |  |
| `"transacted_fx_amount"` |  |
| `"transacted_fx_rate"` |  |
| `"type"` |  |
| `"user_id"` |  |
| `"your_reference"` |  |

Operations: List.

API path: `/api/v1/accounts/{accountID}/payouts/failed`

#### NoFrixionMoneyMoovModelsPayoutMetric

| Field | Description |
| --- | --- |
| `"all"` |  |
| `"failed"` |  |
| `"in_progress"` |  |
| `"paid"` |  |
| `"pending_approval"` |  |
| `"scheduled"` |  |
| `"total_amounts_by_currency"` |  |

Operations: Load.

API path: `/api/v1/payouts/metrics`

#### NoFrixionMoneyMoovModelsPayoutsPayoutsCreate

| Field | Description |
| --- | --- |
| `"failed_payout"` |  |
| `"payout"` |  |

Operations: Create.

API path: `/api/v1/payouts/batchcreate`

#### NoFrixionMoneyMoovModelsPayrun

| Field | Description |
| --- | --- |
| `"authorisation"` |  |
| `"authorisation_date"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"batch_payout_id"` |  |
| `"can_authorise"` |  |
| `"can_delete"` |  |
| `"can_edit"` |  |
| `"event"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"invoice"` |  |
| `"invoices_minimal"` |  |
| `"is_archived"` |  |
| `"last_updated"` |  |
| `"last_updated_by"` |  |
| `"merchant_id"` |  |
| `"name"` |  |
| `"nonce"` |  |
| `"payment"` |  |
| `"payout"` |  |
| `"payouts_count"` |  |
| `"reason"` |  |
| `"schedule_date"` |  |
| `"scheduled_date"` |  |
| `"source_account"` |  |
| `"status"` |  |
| `"total_eur"` |  |
| `"total_gbp"` |  |
| `"total_usd"` |  |

Operations: Create, Load, Update.

API path: `/api/v1/payruns/{merchantID}`

#### NoFrixionMoneyMoovModelsReportResult

| Field | Description |
| --- | --- |
| `"content"` |  |
| `"content_type"` |  |
| `"last_completed_at"` |  |
| `"merchant_id"` |  |
| `"report_name"` |  |
| `"report_type"` |  |
| `"statement_number"` |  |

Operations: Load.

API path: `/api/v1/reports/{id}/result/{statementNumber}`

#### NoFrixionMoneyMoovModelsRule

| Field | Description |
| --- | --- |
| `"account"` |  |
| `"account_id"` |  |
| `"approve_url"` |  |
| `"approver_id"` |  |
| `"authentication_method"` |  |
| `"authorisation"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"can_authorise"` |  |
| `"created_by"` |  |
| `"description"` |  |
| `"end_at"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"is_disabled"` |  |
| `"last_executed_at"` |  |
| `"last_run_at_transaction_date"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"name"` |  |
| `"nonce"` |  |
| `"on_approved_web_hook_url"` |  |
| `"on_execution_error_web_hook_url"` |  |
| `"on_execution_success_web_hook_url"` |  |
| `"start_at"` |  |
| `"status"` |  |
| `"sweep_action"` |  |
| `"time_zone_id"` |  |
| `"trigger_cron_expression"` |  |
| `"trigger_on_pay_in"` |  |
| `"user_id"` |  |
| `"web_hook_secret"` |  |

Operations: Create, Load, Update.

API path: `/api/v1/rules`

#### NoFrixionMoneyMoovModelsTransaction

| Field | Description |
| --- | --- |
| `"account_id"` |  |
| `"account_name"` |  |
| `"account_sequence_number"` |  |
| `"amount"` |  |
| `"amount_minor_unit"` |  |
| `"balance"` |  |
| `"balance_minor_unit"` |  |
| `"counterparty"` |  |
| `"counterparty_summary"` |  |
| `"currency"` |  |
| `"description"` |  |
| `"fx_amount"` |  |
| `"fx_currency"` |  |
| `"fx_rate"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"merchant_id"` |  |
| `"payment_request_custom_field"` |  |
| `"payment_request_id"` |  |
| `"payout_id"` |  |
| `"raw_reference"` |  |
| `"rule_id"` |  |
| `"tag"` |  |
| `"their_reference"` |  |
| `"transaction_date"` |  |
| `"type"` |  |
| `"virtual_iban"` |  |
| `"your_reference"` |  |

Operations: Load.

API path: `/api/v1/accounts/{accountID}/transactions/{id}`

#### NoFrixionMoneyMoovModelsTransactionPage

| Field | Description |
| --- | --- |
| `"account_id"` |  |
| `"account_name"` |  |
| `"account_sequence_number"` |  |
| `"amount"` |  |
| `"amount_minor_unit"` |  |
| `"balance"` |  |
| `"balance_minor_unit"` |  |
| `"content"` |  |
| `"counterparty"` |  |
| `"counterparty_summary"` |  |
| `"currency"` |  |
| `"description"` |  |
| `"fx_amount"` |  |
| `"fx_currency"` |  |
| `"fx_rate"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"merchant_id"` |  |
| `"page_number"` |  |
| `"page_size"` |  |
| `"payment_request_custom_field"` |  |
| `"payment_request_id"` |  |
| `"payout_id"` |  |
| `"raw_reference"` |  |
| `"rule_id"` |  |
| `"tag"` |  |
| `"their_reference"` |  |
| `"total_page"` |  |
| `"total_size"` |  |
| `"transaction_date"` |  |
| `"type"` |  |
| `"virtual_iban"` |  |
| `"your_reference"` |  |

Operations: List, Load.

API path: `/api/v1/accounts/{accountID}/transactions`

#### NoFrixionMoneyMoovModelsUserInvite

| Field | Description |
| --- | --- |
| `"authorisation_status"` |  |
| `"id"` |  |
| `"initial_role_id"` |  |
| `"invitee_email_address"` |  |
| `"invitee_first_name"` |  |
| `"invitee_last_name"` |  |
| `"inviter_email_address"` |  |
| `"inviter_first_name"` |  |
| `"inviter_last_name"` |  |
| `"is_authorised"` |  |
| `"is_invitee_registered"` |  |
| `"last_invited"` |  |
| `"merchant_id"` |  |
| `"merchant_name"` |  |
| `"message"` |  |
| `"registration_url"` |  |
| `"send_invite_email"` |  |
| `"status"` |  |
| `"user"` |  |
| `"user_id"` |  |

Operations: Create, Load.

API path: `/api/v1/userinvites`

#### NoFrixionMoneyMoovModelsUserInvitePage

| Field | Description |
| --- | --- |
| `"authorisation_status"` |  |
| `"id"` |  |
| `"initial_role_id"` |  |
| `"invitee_email_address"` |  |
| `"invitee_first_name"` |  |
| `"invitee_last_name"` |  |
| `"inviter_email_address"` |  |
| `"inviter_first_name"` |  |
| `"inviter_last_name"` |  |
| `"is_authorised"` |  |
| `"is_invitee_registered"` |  |
| `"last_invited"` |  |
| `"merchant_id"` |  |
| `"merchant_name"` |  |
| `"message"` |  |
| `"registration_url"` |  |
| `"status"` |  |
| `"user"` |  |
| `"user_id"` |  |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/userinvitespaged`

#### NoFrixionMoneyMoovModelsUserPage

| Field | Description |
| --- | --- |
| `"client_session_timeout"` |  |
| `"email_address"` |  |
| `"first_name"` |  |
| `"id"` |  |
| `"last_name"` |  |
| `"passkey_added"` |  |
| `"permission"` |  |
| `"roles_with_scope"` |  |
| `"two_factor_enabled"` |  |

Operations: List.

API path: `/api/v1/user/{merchantID}/userspaged`

#### NoFrixionMoneyMoovModelsWebhook

| Field | Description |
| --- | --- |
| `"destination_url"` |  |
| `"email_address"` |  |
| `"failed_notification_email_address"` |  |
| `"id"` |  |
| `"is_active"` |  |
| `"merchant_id"` |  |
| `"notification_method"` |  |
| `"resource_type"` |  |
| `"retry"` |  |
| `"secret"` |  |
| `"version"` |  |

Operations: Create, List, Load, Update.

API path: `/api/v1/webhooks`

#### OpenBanking

| Field | Description |
| --- | --- |

Operations: Create, Remove.

API path: `/api/v1/openbanking/account/{accountID}/synchronise`

#### Payeeverification

| Field | Description |
| --- | --- |
| `"account_name"` |  |
| `"account_number"` |  |
| `"iban"` |  |
| `"payee_verified_account_name"` |  |
| `"result"` |  |
| `"secondary_identification"` |  |
| `"sort_code"` |  |

Operations: Create.

API path: `/api/v1/openbanking/payeeverification`

#### PaymentRequest

| Field | Description |
| --- | --- |
| `"amount"` |  |
| `"do_simulate_settlement_failure"` |  |
| `"error_description"` |  |
| `"institution"` |  |
| `"payment_initiation_id"` |  |

Operations: Create, Load, Remove, Update.

API path: `/api/v1/paymentrequests/{id}/directdebit`

#### Payout

| Field | Description |
| --- | --- |
| `"account_id"` |  |
| `"allow_incomplete"` |  |
| `"amount"` |  |
| `"amount_minor_unit"` |  |
| `"approve_payout_url"` |  |
| `"approver_id"` |  |
| `"authentication_method"` |  |
| `"authorisation"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"batch_payout_id"` |  |
| `"beneficiary"` |  |
| `"beneficiary_id"` |  |
| `"can_authorise"` |  |
| `"can_process"` |  |
| `"can_update"` |  |
| `"charge_bearer"` |  |
| `"created_by"` |  |
| `"created_by_email_address"` |  |
| `"currency"` |  |
| `"current_user_id"` |  |
| `"description"` |  |
| `"destination"` |  |
| `"document"` |  |
| `"event"` |  |
| `"formatted_amount"` |  |
| `"formatted_fx_destination_amount"` |  |
| `"formatted_schedule"` |  |
| `"formatted_schedule_day_only"` |  |
| `"formatted_source_account_available_balance"` |  |
| `"fx_destination_amount"` |  |
| `"fx_destination_amount_minor_unit"` |  |
| `"fx_destination_currency"` |  |
| `"fx_quote_expires_at"` |  |
| `"fx_quote_id"` |  |
| `"fx_rate"` |  |
| `"fx_use_destination_amount"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"invoice_id"` |  |
| `"is_archived"` |  |
| `"is_failed"` |  |
| `"is_settled"` |  |
| `"is_submitted"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"merchant_token_description"` |  |
| `"nonce"` |  |
| `"payment_processor"` |  |
| `"payment_rail"` |  |
| `"payrun_id"` |  |
| `"payrun_name"` |  |
| `"rule"` |  |
| `"schedule_date"` |  |
| `"scheduled"` |  |
| `"source_account_available_balance"` |  |
| `"source_account_available_balance_minor_unit"` |  |
| `"source_account_bic"` |  |
| `"source_account_currency"` |  |
| `"source_account_iban"` |  |
| `"source_account_identifier"` |  |
| `"source_account_name"` |  |
| `"source_account_number"` |  |
| `"source_account_sortcode"` |  |
| `"status"` |  |
| `"tag"` |  |
| `"tag_id"` |  |
| `"their_reference"` |  |
| `"topup_payrun_id"` |  |
| `"transacted_amount"` |  |
| `"transacted_fx_amount"` |  |
| `"transacted_fx_rate"` |  |
| `"type"` |  |
| `"user_id"` |  |
| `"your_reference"` |  |

Operations: Create, Load, Remove, Update.

API path: `/api/v1/payouts/batch/submit/{id}`

#### Payrun

| Field | Description |
| --- | --- |
| `"id"` |  |
| `"note"` |  |
| `"scheduled_date"` |  |

Operations: Create, Remove, Update.

API path: `/api/v1/payruns/{id}/request-authorisation`

#### Reject

| Field | Description |
| --- | --- |
| `"account_id"` |  |
| `"amount"` |  |
| `"amount_minor_unit"` |  |
| `"approve_payout_url"` |  |
| `"approver_id"` |  |
| `"authentication_method"` |  |
| `"authorisation"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"batch_payout_id"` |  |
| `"beneficiary"` |  |
| `"can_authorise"` |  |
| `"can_process"` |  |
| `"can_update"` |  |
| `"charge_bearer"` |  |
| `"created_by"` |  |
| `"created_by_email_address"` |  |
| `"currency"` |  |
| `"current_user_id"` |  |
| `"description"` |  |
| `"destination"` |  |
| `"document"` |  |
| `"event"` |  |
| `"formatted_amount"` |  |
| `"formatted_fx_destination_amount"` |  |
| `"formatted_schedule"` |  |
| `"formatted_schedule_day_only"` |  |
| `"formatted_source_account_available_balance"` |  |
| `"fx_destination_amount"` |  |
| `"fx_destination_amount_minor_unit"` |  |
| `"fx_destination_currency"` |  |
| `"fx_quote_expires_at"` |  |
| `"fx_quote_id"` |  |
| `"fx_rate"` |  |
| `"fx_use_destination_amount"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"invoice_id"` |  |
| `"is_archived"` |  |
| `"is_failed"` |  |
| `"is_settled"` |  |
| `"is_submitted"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"merchant_token_description"` |  |
| `"nonce"` |  |
| `"payment_processor"` |  |
| `"payment_rail"` |  |
| `"payrun_id"` |  |
| `"payrun_name"` |  |
| `"reason"` |  |
| `"rule"` |  |
| `"schedule_date"` |  |
| `"scheduled"` |  |
| `"source_account_available_balance"` |  |
| `"source_account_available_balance_minor_unit"` |  |
| `"source_account_bic"` |  |
| `"source_account_currency"` |  |
| `"source_account_iban"` |  |
| `"source_account_identifier"` |  |
| `"source_account_name"` |  |
| `"source_account_number"` |  |
| `"source_account_sortcode"` |  |
| `"status"` |  |
| `"tag"` |  |
| `"their_reference"` |  |
| `"topup_payrun_id"` |  |
| `"transacted_amount"` |  |
| `"transacted_fx_amount"` |  |
| `"transacted_fx_rate"` |  |
| `"type"` |  |
| `"user_id"` |  |
| `"your_reference"` |  |

Operations: Update.

API path: `/api/v1/payouts/reject/{id}`

#### Report

| Field | Description |
| --- | --- |

Operations: Update.

API path: `/api/v1/reports/{id}/initiate`

#### Rule

| Field | Description |
| --- | --- |

Operations: Remove, Update.

API path: `/api/v1/rules/{id}`

#### Send

| Field | Description |
| --- | --- |
| `"account_id"` |  |
| `"allow_incomplete"` |  |
| `"amount"` |  |
| `"amount_minor_unit"` |  |
| `"approve_payout_url"` |  |
| `"approver_id"` |  |
| `"authentication_method"` |  |
| `"authorisation"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"batch_payout_id"` |  |
| `"beneficiary"` |  |
| `"beneficiary_id"` |  |
| `"can_authorise"` |  |
| `"can_process"` |  |
| `"can_update"` |  |
| `"charge_bearer"` |  |
| `"created_by"` |  |
| `"created_by_email_address"` |  |
| `"currency"` |  |
| `"current_user_id"` |  |
| `"description"` |  |
| `"destination"` |  |
| `"document"` |  |
| `"event"` |  |
| `"formatted_amount"` |  |
| `"formatted_fx_destination_amount"` |  |
| `"formatted_schedule"` |  |
| `"formatted_schedule_day_only"` |  |
| `"formatted_source_account_available_balance"` |  |
| `"fx_destination_amount"` |  |
| `"fx_destination_amount_minor_unit"` |  |
| `"fx_destination_currency"` |  |
| `"fx_quote_expires_at"` |  |
| `"fx_quote_id"` |  |
| `"fx_rate"` |  |
| `"fx_use_destination_amount"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"invoice_id"` |  |
| `"is_archived"` |  |
| `"is_failed"` |  |
| `"is_settled"` |  |
| `"is_submitted"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"merchant_token_description"` |  |
| `"nonce"` |  |
| `"payment_processor"` |  |
| `"payment_rail"` |  |
| `"payrun_id"` |  |
| `"payrun_name"` |  |
| `"rule"` |  |
| `"schedule_date"` |  |
| `"scheduled"` |  |
| `"source_account_available_balance"` |  |
| `"source_account_available_balance_minor_unit"` |  |
| `"source_account_bic"` |  |
| `"source_account_currency"` |  |
| `"source_account_iban"` |  |
| `"source_account_identifier"` |  |
| `"source_account_name"` |  |
| `"source_account_number"` |  |
| `"source_account_sortcode"` |  |
| `"status"` |  |
| `"tag"` |  |
| `"tag_id"` |  |
| `"their_reference"` |  |
| `"topup_payrun_id"` |  |
| `"transacted_amount"` |  |
| `"transacted_fx_amount"` |  |
| `"transacted_fx_rate"` |  |
| `"type"` |  |
| `"user_id"` |  |
| `"your_reference"` |  |

Operations: Create.

API path: `/api/v1/payouts/send`

#### Sendbeneficiary

| Field | Description |
| --- | --- |
| `"account_id"` |  |
| `"allow_incomplete"` |  |
| `"amount"` |  |
| `"amount_minor_unit"` |  |
| `"approve_payout_url"` |  |
| `"approver_id"` |  |
| `"authentication_method"` |  |
| `"authorisation"` |  |
| `"authorisers_completed_count"` |  |
| `"authorisers_required_count"` |  |
| `"batch_payout_id"` |  |
| `"beneficiary"` |  |
| `"beneficiary_id"` |  |
| `"can_authorise"` |  |
| `"can_process"` |  |
| `"can_update"` |  |
| `"charge_bearer"` |  |
| `"created_by"` |  |
| `"created_by_email_address"` |  |
| `"currency"` |  |
| `"current_user_id"` |  |
| `"description"` |  |
| `"destination"` |  |
| `"document"` |  |
| `"event"` |  |
| `"formatted_amount"` |  |
| `"formatted_fx_destination_amount"` |  |
| `"formatted_schedule"` |  |
| `"formatted_schedule_day_only"` |  |
| `"formatted_source_account_available_balance"` |  |
| `"fx_destination_amount"` |  |
| `"fx_destination_amount_minor_unit"` |  |
| `"fx_destination_currency"` |  |
| `"fx_quote_expires_at"` |  |
| `"fx_quote_id"` |  |
| `"fx_rate"` |  |
| `"fx_use_destination_amount"` |  |
| `"has_current_user_authorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"invoice_id"` |  |
| `"is_archived"` |  |
| `"is_failed"` |  |
| `"is_settled"` |  |
| `"is_submitted"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"merchant_token_description"` |  |
| `"nonce"` |  |
| `"payment_processor"` |  |
| `"payment_rail"` |  |
| `"payrun_id"` |  |
| `"payrun_name"` |  |
| `"rule"` |  |
| `"schedule_date"` |  |
| `"scheduled"` |  |
| `"source_account_available_balance"` |  |
| `"source_account_available_balance_minor_unit"` |  |
| `"source_account_bic"` |  |
| `"source_account_currency"` |  |
| `"source_account_iban"` |  |
| `"source_account_identifier"` |  |
| `"source_account_name"` |  |
| `"source_account_number"` |  |
| `"source_account_sortcode"` |  |
| `"status"` |  |
| `"tag"` |  |
| `"tag_id"` |  |
| `"their_reference"` |  |
| `"topup_payrun_id"` |  |
| `"transacted_amount"` |  |
| `"transacted_fx_amount"` |  |
| `"transacted_fx_rate"` |  |
| `"type"` |  |
| `"user_id"` |  |
| `"your_reference"` |  |

Operations: Create.

API path: `/api/v1/payouts/sendbeneficiary`

#### Tag

| Field | Description |
| --- | --- |
| `"colour_hex"` |  |
| `"description"` |  |
| `"id"` |  |
| `"merchant_id"` |  |
| `"name"` |  |

Operations: Create, List.

API path: `/api/v1/merchants/{merchantID}/tags`

#### Token

| Field | Description |
| --- | --- |

Operations: Create, Remove.

API path: `/api/v1/tokens/authorise/{id}`

#### Transaction

| Field | Description |
| --- | --- |

Operations: Create, Load, Remove.

API path: `/api/v1/transactions/{id}/tags`

#### User

| Field | Description |
| --- | --- |
| `"client_session_timeout"` |  |
| `"email_address"` |  |
| `"first_name"` |  |
| `"id"` |  |
| `"last_name"` |  |
| `"passkey_added"` |  |
| `"permission"` |  |
| `"profile"` |  |
| `"roles_with_scope"` |  |
| `"two_factor_enabled"` |  |
| `"user_invite_id"` |  |

Operations: List, Update.

API path: `/api/v1/merchants/{merchantID}/users`

#### UserInvite

| Field | Description |
| --- | --- |

Operations: Create, Remove, Update.

API path: `/api/v1/userinvites/authorise/{id}`

#### Virtual

| Field | Description |
| --- | --- |
| `"account_name"` |  |
| `"account_supplier_name"` |  |
| `"available_balance"` |  |
| `"available_balance_minor_unit"` |  |
| `"balance"` |  |
| `"balance_minor_unit"` |  |
| `"bank_name"` |  |
| `"consent_id"` |  |
| `"created_by"` |  |
| `"created_by_display_name"` |  |
| `"currency"` |  |
| `"default_payment_rail"` |  |
| `"display_name"` |  |
| `"expiry_date"` |  |
| `"external_account_icon"` |  |
| `"id"` |  |
| `"identifier"` |  |
| `"inserted"` |  |
| `"is_archived"` |  |
| `"is_connected_account"` |  |
| `"is_default"` |  |
| `"is_trust_account"` |  |
| `"is_virtual"` |  |
| `"last_transaction"` |  |
| `"last_updated"` |  |
| `"merchant_id"` |  |
| `"merchant_name"` |  |
| `"name"` |  |
| `"physical_account_id"` |  |
| `"rule"` |  |
| `"submitted_payouts_balance"` |  |
| `"submitted_payouts_balance_minor_unit"` |  |
| `"summary"` |  |
| `"supplier_sepa_instant_status"` |  |
| `"xero_bank_feed_connection_status"` |  |
| `"xero_bank_feed_last_synced_at"` |  |
| `"xero_bank_feed_sync_last_failed_at"` |  |
| `"xero_bank_feed_sync_last_failure_reason"` |  |
| `"xero_bank_feed_sync_status"` |  |
| `"xero_unsynchronised_transactions_count"` |  |

Operations: Create, Update.

API path: `/api/v1/accounts/{accountID}/virtual`

#### Webhook

| Field | Description |
| --- | --- |

Operations: Remove.

API path: `/api/v1/webhooks/{id}`

#### Whoami

| Field | Description |
| --- | --- |
| `"client_session_timeout"` |  |
| `"email_address"` |  |
| `"first_name"` |  |
| `"id"` |  |
| `"last_name"` |  |
| `"passkey_added"` |  |
| `"permission"` |  |
| `"roles_with_scope"` |  |
| `"two_factor_enabled"` |  |

Operations: List.

API path: `/api/v1/metadata/whoami`

#### Whoamitrustedapp

| Field | Description |
| --- | --- |
| `"client_session_timeout"` |  |
| `"email_address"` |  |
| `"first_name"` |  |
| `"id"` |  |
| `"last_name"` |  |
| `"passkey_added"` |  |
| `"permission"` |  |
| `"roles_with_scope"` |  |
| `"two_factor_enabled"` |  |

Operations: List.

API path: `/api/v1/metadata/whoamitrustedapp`



## Entities


### Account

Create an instance: `account := client.Account(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `account_name` | `string` |  |
| `account_supplier_name` | `string` |  |
| `account_type` | `string` |  |
| `available_balance` | `float64` |  |
| `available_balance_minor_unit` | `int` |  |
| `balance` | `float64` |  |
| `balance_minor_unit` | `int` |  |
| `bank_name` | `string` |  |
| `consent_id` | `string` |  |
| `created_by` | `map[string]any` |  |
| `created_by_display_name` | `string` |  |
| `currency` | `string` |  |
| `default_payment_rail` | `string` |  |
| `display_name` | `string` |  |
| `expiry_date` | `string` |  |
| `external_account_icon` | `string` |  |
| `format` | `string` |  |
| `from_date` | `string` |  |
| `id` | `string` |  |
| `identifier` | `map[string]any` |  |
| `inserted` | `string` |  |
| `is_archived` | `bool` |  |
| `is_connected_account` | `bool` |  |
| `is_default` | `bool` |  |
| `is_trust_account` | `bool` |  |
| `is_virtual` | `bool` |  |
| `last_transaction` | `map[string]any` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `physical_account_id` | `string` |  |
| `role_i_d` | `[]any` |  |
| `rule` | `[]any` |  |
| `submitted_payouts_balance` | `float64` |  |
| `submitted_payouts_balance_minor_unit` | `int` |  |
| `summary` | `string` |  |
| `supplier_physical_account_id` | `string` |  |
| `supplier_sepa_instant_status` | `string` |  |
| `to_date` | `string` |  |
| `xero_bank_feed_connection_status` | `string` |  |
| `xero_bank_feed_last_synced_at` | `string` |  |
| `xero_bank_feed_sync_last_failed_at` | `string` |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` |  |
| `xero_bank_feed_sync_status` | `string` |  |
| `xero_unsynchronised_transactions_count` | `int` |  |

#### Example: Load

```go
account, err := client.Account(nil).Load(map[string]any{"id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(account) // the loaded record
```

#### Example: List

```go
accounts, err := client.Account(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(accounts) // the array of records
```

#### Example: Create

```go
result, err := client.Account(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Beneficiary

Create an instance: `beneficiary := client.Beneficiary(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_callback_url` | `string` |  |
| `authentication_method` | `[]any` |  |
| `authorisation` | `[]any` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `beneficiary_event` | `[]any` |  |
| `can_authorise` | `bool` |  |
| `can_update` | `bool` |  |
| `created_by` | `map[string]any` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `destination` | `map[string]any` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_enabled` | `bool` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `source_account` | `[]any` |  |
| `source_account_i_d` | `[]any` |  |
| `their_reference` | `string` |  |

#### Example: Load

```go
beneficiary, err := client.Beneficiary(nil).Load(map[string]any{"id": "beneficiary_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(beneficiary) // the loaded record
```

#### Example: Create

```go
result, err := client.Beneficiary(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Cancel

Create an instance: `cancel := client.Cancel(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `amount` | `float64` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `[]any` |  |
| `authorisation` | `[]any` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `map[string]any` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `map[string]any` |  |
| `document` | `[]any` |  |
| `event` | `[]any` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `float64` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `float64` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `rule` | `map[string]any` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float64` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `map[string]any` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `[]any` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `float64` |  |
| `transacted_fx_amount` | `float64` |  |
| `transacted_fx_rate` | `float64` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |


### Disable

Create an instance: `disable := client.Disable(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_callback_url` | `string` |  |
| `authentication_method` | `[]any` |  |
| `authorisation` | `[]any` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `beneficiary_event` | `[]any` |  |
| `can_authorise` | `bool` |  |
| `can_update` | `bool` |  |
| `created_by` | `map[string]any` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `destination` | `map[string]any` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_enabled` | `bool` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `source_account` | `[]any` |  |
| `their_reference` | `string` |  |


### Enable

Create an instance: `enable := client.Enable(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_callback_url` | `string` |  |
| `authentication_method` | `[]any` |  |
| `authorisation` | `[]any` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `beneficiary_event` | `[]any` |  |
| `can_authorise` | `bool` |  |
| `can_update` | `bool` |  |
| `created_by` | `map[string]any` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `destination` | `map[string]any` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_enabled` | `bool` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `source_account` | `[]any` |  |
| `their_reference` | `string` |  |


### Merchant

Create an instance: `merchant := client.Merchant(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `reason` | `string` |  |

#### Example: Load

```go
merchant, err := client.Merchant(nil).Load(map[string]any{"merchant_id": "merchant_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(merchant) // the loaded record
```


### Metadata

Create an instance: `metadata := client.Metadata(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Example: Load

```go
metadata, err := client.Metadata(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(metadata) // the loaded record
```


### NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage

Create an instance: `noFrixionBizBizModelsPagingMerchantDirectDebitMandatePage := client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

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
| `is_recurring` | `bool` |  |
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

```go
noFrixionBizBizModelsPagingMerchantDirectDebitMandatePages, err := client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionBizBizModelsPagingMerchantDirectDebitMandatePages) // the array of records
```


### NoFrixionBizBizModelsPagingPaymentRequestPage

Create an instance: `noFrixionBizBizModelsPagingPaymentRequestPage := client.NoFrixionBizBizModelsPagingPaymentRequestPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `[]any` |  |
| `amount` | `float64` |  |
| `amount_pending` | `float64` |  |
| `amount_received` | `float64` |  |
| `amount_refunded` | `float64` |  |
| `auto_send_receipt` | `bool` |  |
| `base_origin_url` | `string` |  |
| `callback_url` | `string` |  |
| `card_authorize_only` | `bool` |  |
| `card_create_token` | `bool` |  |
| `card_create_token_mode` | `string` |  |
| `card_ignore_cvn` | `bool` |  |
| `card_processor_merchant_id` | `string` |  |
| `card_stripe_payment_intent_id` | `string` |  |
| `card_stripe_payment_intent_secret` | `string` |  |
| `created_by_user` | `map[string]any` |  |
| `currency` | `string` |  |
| `custom_field` | `[]any` |  |
| `customer_email_address` | `string` |  |
| `customer_id` | `string` |  |
| `customer_name` | `string` |  |
| `description` | `string` |  |
| `destination_account` | `map[string]any` |  |
| `direct_debit_payment` | `map[string]any` |  |
| `due_date` | `string` |  |
| `event` | `[]any` |  |
| `failure_callback_url` | `string` |  |
| `field_display_setting` | `[]any` |  |
| `formatted_amount` | `string` |  |
| `hosted_pay_checkout_url` | `string` |  |
| `id` | `string` |  |
| `ignore_address_verification` | `bool` |  |
| `inserted` | `string` |  |
| `inserted_sortable` | `string` |  |
| `is_archived` | `bool` |  |
| `jwk` | `string` |  |
| `last_updated` | `string` |  |
| `lightning_invoice` | `string` |  |
| `lightning_invoice_expires_at` | `string` |  |
| `merchant_direct_debit_mandate_id` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `notification_email_address` | `string` |  |
| `notification_role_i_d` | `[]any` |  |
| `order_id` | `string` |  |
| `partial_payment_method` | `string` |  |
| `partial_payment_step` | `string` |  |
| `payment_attempt` | `[]any` |  |
| `payment_method` | `[]any` |  |
| `payment_processor` | `string` |  |
| `payrun_id` | `string` |  |
| `pisp_account_id` | `string` |  |
| `priority_bank_id` | `string` |  |
| `result` | `map[string]any` |  |
| `sandbox_settle_delay_in_second` | `int` |  |
| `shipping_address` | `map[string]any` |  |
| `status` | `string` |  |
| `success_web_hook_url` | `string` |  |
| `tag` | `[]any` |  |
| `title` | `string` |  |
| `tokenised_card` | `[]any` |  |
| `transaction` | `[]any` |  |
| `use_hosted_payment_page` | `bool` |  |

#### Example: List

```go
noFrixionBizBizModelsPagingPaymentRequestPages, err := client.NoFrixionBizBizModelsPagingPaymentRequestPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionBizBizModelsPagingPaymentRequestPages) // the array of records
```


### NoFrixionBizBizModelsPagingPayoutPage

Create an instance: `noFrixionBizBizModelsPagingPayoutPage := client.NoFrixionBizBizModelsPagingPayoutPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `amount` | `float64` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `[]any` |  |
| `authorisation` | `[]any` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `map[string]any` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `map[string]any` |  |
| `document` | `[]any` |  |
| `event` | `[]any` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `float64` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `float64` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `rule` | `map[string]any` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float64` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `map[string]any` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `[]any` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `float64` |  |
| `transacted_fx_amount` | `float64` |  |
| `transacted_fx_rate` | `float64` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |

#### Example: List

```go
noFrixionBizBizModelsPagingPayoutPages, err := client.NoFrixionBizBizModelsPagingPayoutPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionBizBizModelsPagingPayoutPages) // the array of records
```


### NoFrixionBizBizModelsPagingPayrunPage

Create an instance: `noFrixionBizBizModelsPagingPayrunPage := client.NoFrixionBizBizModelsPagingPayrunPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation` | `[]any` |  |
| `authorisation_date` | `string` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `string` |  |
| `can_authorise` | `bool` |  |
| `can_delete` | `bool` |  |
| `can_edit` | `bool` |  |
| `event` | `[]any` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice` | `[]any` |  |
| `invoices_minimal` | `[]any` |  |
| `is_archived` | `bool` |  |
| `last_updated` | `string` |  |
| `last_updated_by` | `map[string]any` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `payment` | `[]any` |  |
| `payout` | `[]any` |  |
| `payouts_count` | `int` |  |
| `schedule_date` | `string` |  |
| `source_account` | `[]any` |  |
| `status` | `string` |  |
| `total_eur` | `float64` |  |
| `total_gbp` | `float64` |  |
| `total_usd` | `float64` |  |

#### Example: List

```go
noFrixionBizBizModelsPagingPayrunPages, err := client.NoFrixionBizBizModelsPagingPayrunPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionBizBizModelsPagingPayrunPages) // the array of records
```


### NoFrixionBizBizModelsPagingRuleEventsPage

Create an instance: `noFrixionBizBizModelsPagingRuleEventsPage := client.NoFrixionBizBizModelsPagingRuleEventsPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `error_message` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_authorise_to_enable` | `bool` |  |
| `message` | `string` |  |
| `raw_response` | `string` |  |
| `rule_event_type` | `string` |  |
| `rule_id` | `string` |  |
| `user` | `map[string]any` |  |

#### Example: List

```go
noFrixionBizBizModelsPagingRuleEventsPages, err := client.NoFrixionBizBizModelsPagingRuleEventsPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionBizBizModelsPagingRuleEventsPages) // the array of records
```


### NoFrixionBizBizModelsPagingRulesPage

Create an instance: `noFrixionBizBizModelsPagingRulesPage := client.NoFrixionBizBizModelsPagingRulesPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account` | `map[string]any` |  |
| `account_id` | `string` |  |
| `approve_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `[]any` |  |
| `authorisation` | `[]any` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `can_authorise` | `bool` |  |
| `created_by` | `map[string]any` |  |
| `description` | `string` |  |
| `end_at` | `string` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_disabled` | `bool` |  |
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
| `sweep_action` | `map[string]any` |  |
| `time_zone_id` | `string` |  |
| `trigger_cron_expression` | `string` |  |
| `trigger_on_pay_in` | `bool` |  |
| `user_id` | `string` |  |
| `web_hook_secret` | `string` |  |

#### Example: List

```go
noFrixionBizBizModelsPagingRulesPages, err := client.NoFrixionBizBizModelsPagingRulesPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionBizBizModelsPagingRulesPages) // the array of records
```


### NoFrixionBizBizModelsPaymentsCardPayment

Create an instance: `noFrixionBizBizModelsPaymentsCardPayment := client.NoFrixionBizBizModelsPaymentsCardPayment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorized_amount` | `string` |  |
| `currency_code` | `string` |  |
| `is_payer_authentication_required` | `bool` |  |
| `is_soft_decline` | `bool` |  |
| `payer_authentication_access_token` | `string` |  |
| `payer_authentication_merchant_data` | `string` |  |
| `payer_authentication_url` | `string` |  |
| `payer_authentication_window_height` | `int` |  |
| `payer_authentication_window_width` | `int` |  |
| `payment_request_callback_url` | `string` |  |
| `payment_request_id` | `string` |  |
| `request_id` | `string` |  |
| `response_code` | `string` |  |
| `response_type` | `string` |  |
| `status` | `string` |  |
| `three_ds_redirect_url` | `string` |  |
| `transaction_id` | `string` |  |

#### Example: Create

```go
result, err := client.NoFrixionBizBizModelsPaymentsCardPayment(nil).Create(map[string]any{
    "paymentrequest_id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionBizBizModelsPaymentsCardPublicKey

Create an instance: `noFrixionBizBizModelsPaymentsCardPublicKey := client.NoFrixionBizBizModelsPaymentsCardPublicKey(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `jwt` | `string` |  |

#### Example: Load

```go
noFrixionBizBizModelsPaymentsCardPublicKey, err := client.NoFrixionBizBizModelsPaymentsCardPublicKey(nil).Load(map[string]any{"paymentrequest_id": "paymentrequest_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionBizBizModelsPaymentsCardPublicKey) // the loaded record
```


### NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries

Create an instance: `noFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries := client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `beneficiary` | `[]any` |  |
| `failed_beneficiary` | `map[string]any` |  |

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment

Create an instance: `noFrixionMoneyMoovApiFeaturesPaymentRequestsPayment := client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_payment_request` | `map[string]any` |  |
| `payment_request` | `[]any` |  |

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate

Create an instance: `noFrixionMoneyMoovApiFeaturesPermissionsRolesCreate := client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_role` | `map[string]any` |  |
| `role` | `[]any` |  |

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(nil).Create(map[string]any{
    "merchant_id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovApiFeaturesUserInvitesCreate

Create an instance: `noFrixionMoneyMoovApiFeaturesUserInvitesCreate := client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_user_invite` | `map[string]any` |  |
| `user_invite` | `[]any` |  |

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant

Create an instance: `noFrixionMoneyMoovModelsAuthorisationSettingsMerchant := client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount_lower` | `float64` |  |
| `amount_upper` | `float64` |  |
| `authorisation_type` | `string` |  |
| `beneficiaries_only` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_editor_cant_authorise` | `bool` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `number_of_authoriser` | `int` |  |
| `role_setting` | `[]any` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsAuthorisationSettingsMerchants, err := client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsAuthorisationSettingsMerchants) // the array of records
```


### NoFrixionMoneyMoovModelsBatchPayout

Create an instance: `noFrixionMoneyMoovModelsBatchPayout := client.NoFrixionMoneyMoovModelsBatchPayout(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approve_url` | `string` |  |
| `id` | `string` |  |
| `payout` | `[]any` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsBatchPayout, err := client.NoFrixionMoneyMoovModelsBatchPayout(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_batch_payout_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsBatchPayout) // the loaded record
```

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovModelsBatchPayout(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovModelsBeneficiaryGroupPage

Create an instance: `noFrixionMoneyMoovModelsBeneficiaryGroupPage := client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `group_member` | `[]any` |  |
| `group_name` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsBeneficiaryGroupPages, err := client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsBeneficiaryGroupPages) // the array of records
```


### NoFrixionMoneyMoovModelsBeneficiaryPage

Create an instance: `noFrixionMoneyMoovModelsBeneficiaryPage := client.NoFrixionMoneyMoovModelsBeneficiaryPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approval_callback_url` | `string` |  |
| `authentication_method` | `[]any` |  |
| `authorisation` | `[]any` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `beneficiary_event` | `[]any` |  |
| `can_authorise` | `bool` |  |
| `can_update` | `bool` |  |
| `created_by` | `map[string]any` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `destination` | `map[string]any` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_enabled` | `bool` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `source_account` | `[]any` |  |
| `their_reference` | `string` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsBeneficiaryPages, err := client.NoFrixionMoneyMoovModelsBeneficiaryPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsBeneficiaryPages) // the array of records
```


### NoFrixionMoneyMoovModelsCardCustomerToken

Create an instance: `noFrixionMoneyMoovModelsCardCustomerToken := client.NoFrixionMoneyMoovModelsCardCustomerToken(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Remove(match, ctrl)` | Remove the matching entity. |

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

```go
noFrixionMoneyMoovModelsCardCustomerToken, err := client.NoFrixionMoneyMoovModelsCardCustomerToken(nil).Load(map[string]any{"customer_email_address": "customer_email_address"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsCardCustomerToken) // the loaded record
```

#### Example: List

```go
noFrixionMoneyMoovModelsCardCustomerTokens, err := client.NoFrixionMoneyMoovModelsCardCustomerToken(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsCardCustomerTokens) // the array of records
```


### NoFrixionMoneyMoovModelsCurrencyCurrencyInfo

Create an instance: `noFrixionMoneyMoovModelsCurrencyCurrencyInfo := client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `code` | `string` |  |
| `decimal` | `int` |  |
| `is_fiat` | `bool` |  |
| `iso4217_alpha_code` | `string` |  |
| `iso4217_numeric_code` | `string` |  |
| `symbol` | `string` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsCurrencyCurrencyInfos, err := client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsCurrencyCurrencyInfos) // the array of records
```


### NoFrixionMoneyMoovModelsDirectDebitBatchSubmit

Create an instance: `noFrixionMoneyMoovModelsDirectDebitBatchSubmit := client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_submission` | `map[string]any` |  |
| `successful_submission` | `[]any` |  |

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovModelsFxRate

Create an instance: `noFrixionMoneyMoovModelsFxRate := client.NoFrixionMoneyMoovModelsFxRate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `destination_currency` | `string` |  |
| `exchange_rate` | `float64` |  |
| `expiry_time` | `string` |  |
| `quote_id` | `string` |  |
| `source_currency` | `string` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsFxRate, err := client.NoFrixionMoneyMoovModelsFxRate(nil).Load(map[string]any{"destination": "destination", "source": "source", "valid_for_minute": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsFxRate) // the loaded record
```

#### Example: List

```go
noFrixionMoneyMoovModelsFxRates, err := client.NoFrixionMoneyMoovModelsFxRate(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsFxRates) // the array of records
```


### NoFrixionMoneyMoovModelsIPayment

Create an instance: `noFrixionMoneyMoovModelsIPayment := client.NoFrixionMoneyMoovModelsIPayment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `payment_request_id` | `string` |  |
| `response_type` | `string` |  |

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovModelsIPayment(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovModelsMandatesMandate

Create an instance: `noFrixionMoneyMoovModelsMandatesMandate := client.NoFrixionMoneyMoovModelsMandatesMandate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

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
| `is_recurring` | `bool` |  |
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

```go
noFrixionMoneyMoovModelsMandatesMandate, err := client.NoFrixionMoneyMoovModelsMandatesMandate(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_mandates_mandate_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsMandatesMandate) // the loaded record
```

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovModelsMandatesMandate(nil).Create(map[string]any{
    "address_line1": /* string */,
    "city": /* string */,
    "country_code": /* string */,
    "email_address": /* string */,
    "first_name": /* string */,
    "last_name": /* string */,
    "postal_code": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovModelsMerchant

Create an instance: `noFrixionMoneyMoovModelsMerchant := client.NoFrixionMoneyMoovModelsMerchant(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_currency` | `[]any` |  |
| `can_have_trust_account` | `bool` |  |
| `card_payment_processor` | `string` |  |
| `company_id` | `string` |  |
| `display_qr_on_hosted_pay` | `bool` |  |
| `hosted_pay_version` | `int` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_blocked` | `bool` |  |
| `is_exited` | `bool` |  |
| `is_suspended` | `bool` |  |
| `jurisdiction` | `string` |  |
| `logo_url_png` | `string` |  |
| `logo_url_svg` | `string` |  |
| `merchant_category_code` | `string` |  |
| `name` | `string` |  |
| `note` | `string` |  |
| `parent_merchant` | `map[string]any` |  |
| `payment_account` | `[]any` |  |
| `payment_account_limit` | `int` |  |
| `short_name` | `string` |  |
| `supported_payment_methods_list` | `[]any` |  |
| `suspension_reason` | `string` |  |
| `tag` | `[]any` |  |
| `time_zone_id` | `string` |  |
| `trading_name` | `string` |  |
| `web_hook_limit` | `int` |  |
| `your_role_name` | `string` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsMerchant, err := client.NoFrixionMoneyMoovModelsMerchant(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_merchant_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsMerchant) // the loaded record
```

#### Example: List

```go
noFrixionMoneyMoovModelsMerchants, err := client.NoFrixionMoneyMoovModelsMerchant(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsMerchants) // the array of records
```


### NoFrixionMoneyMoovModelsMerchantPage

Create an instance: `noFrixionMoneyMoovModelsMerchantPage := client.NoFrixionMoneyMoovModelsMerchantPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_currency` | `[]any` |  |
| `can_have_trust_account` | `bool` |  |
| `card_payment_processor` | `string` |  |
| `company_id` | `string` |  |
| `display_qr_on_hosted_pay` | `bool` |  |
| `hosted_pay_version` | `int` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_blocked` | `bool` |  |
| `is_exited` | `bool` |  |
| `is_suspended` | `bool` |  |
| `jurisdiction` | `string` |  |
| `logo_url_png` | `string` |  |
| `logo_url_svg` | `string` |  |
| `merchant_category_code` | `string` |  |
| `name` | `string` |  |
| `note` | `string` |  |
| `parent_merchant` | `map[string]any` |  |
| `payment_account` | `[]any` |  |
| `payment_account_limit` | `int` |  |
| `short_name` | `string` |  |
| `supported_payment_methods_list` | `[]any` |  |
| `suspension_reason` | `string` |  |
| `tag` | `[]any` |  |
| `time_zone_id` | `string` |  |
| `trading_name` | `string` |  |
| `web_hook_limit` | `int` |  |
| `your_role_name` | `string` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsMerchantPages, err := client.NoFrixionMoneyMoovModelsMerchantPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsMerchantPages) // the array of records
```


### NoFrixionMoneyMoovModelsMerchantPayByBankSetting

Create an instance: `noFrixionMoneyMoovModelsMerchantPayByBankSetting := client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bank_country_code` | `[]any` |  |
| `bank_id` | `string` |  |
| `bank_name` | `string` |  |
| `business_institution_id` | `string` |  |
| `currency` | `string` |  |
| `logo` | `string` |  |
| `message` | `string` |  |
| `message_image_url` | `string` |  |
| `order` | `int` |  |
| `personal_institution_id` | `string` |  |
| `processor` | `string` |  |
| `warning_heading` | `string` |  |
| `warning_message` | `string` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsMerchantPayByBankSettings, err := client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsMerchantPayByBankSettings) // the array of records
```


### NoFrixionMoneyMoovModelsMerchantToken

Create an instance: `noFrixionMoneyMoovModelsMerchantToken := client.NoFrixionMoneyMoovModelsMerchantToken(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authentication_method` | `[]any` |  |
| `authorisation` | `[]any` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `can_authorise` | `bool` |  |
| `description` | `string` |  |
| `expires_at` | `string` |  |
| `has_current_user_authorised` | `bool` |  |
| `hmac_algorithm` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `ip_address_whitelist` | `string` |  |
| `is_archived` | `bool` |  |
| `is_enabled` | `bool` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `nonce` | `string` |  |
| `permission_type` | `[]any` |  |
| `request_signature_version` | `int` |  |
| `shared_secret_algorithm` | `string` |  |
| `shared_secret_base64` | `string` |  |
| `token` | `string` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsMerchantToken, err := client.NoFrixionMoneyMoovModelsMerchantToken(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_merchant_token_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsMerchantToken) // the loaded record
```

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovModelsMerchantToken(nil).Create(map[string]any{
    "nonce": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovModelsMerchantTokenPage

Create an instance: `noFrixionMoneyMoovModelsMerchantTokenPage := client.NoFrixionMoneyMoovModelsMerchantTokenPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authentication_method` | `[]any` |  |
| `authorisation` | `[]any` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `can_authorise` | `bool` |  |
| `description` | `string` |  |
| `expires_at` | `string` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `ip_address_whitelist` | `string` |  |
| `is_archived` | `bool` |  |
| `is_enabled` | `bool` |  |
| `last_authorised` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `nonce` | `string` |  |
| `permission_type` | `[]any` |  |
| `request_signature_version` | `int` |  |
| `shared_secret_algorithm` | `string` |  |
| `shared_secret_base64` | `string` |  |
| `token` | `string` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsMerchantTokenPages, err := client.NoFrixionMoneyMoovModelsMerchantTokenPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsMerchantTokenPages) // the array of records
```


### NoFrixionMoneyMoovModelsNoFrixionVersion

Create an instance: `noFrixionMoneyMoovModelsNoFrixionVersion := client.NoFrixionMoneyMoovModelsNoFrixionVersion(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `build_version` | `int` |  |
| `major_version` | `int` |  |
| `minor_version` | `int` |  |
| `release_name` | `string` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsNoFrixionVersion, err := client.NoFrixionMoneyMoovModelsNoFrixionVersion(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsNoFrixionVersion) // the loaded record
```


### NoFrixionMoneyMoovModelsOpenBankingAccount

Create an instance: `noFrixionMoneyMoovModelsOpenBankingAccount := client.NoFrixionMoneyMoovModelsOpenBankingAccount(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_balance` | `[]any` |  |
| `account_identification` | `[]any` |  |
| `account_name` | `[]any` |  |
| `account_type` | `string` |  |
| `balance` | `float64` |  |
| `consolidated_account_information` | `map[string]any` |  |
| `currency` | `string` |  |
| `description` | `string` |  |
| `detail` | `string` |  |
| `id` | `string` |  |
| `nickname` | `string` |  |
| `type` | `string` |  |
| `usage_type` | `string` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsOpenBankingAccount, err := client.NoFrixionMoneyMoovModelsOpenBankingAccount(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_open_banking_account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsOpenBankingAccount) // the loaded record
```


### NoFrixionMoneyMoovModelsOpenBankingConsent

Create an instance: `noFrixionMoneyMoovModelsOpenBankingConsent := client.NoFrixionMoneyMoovModelsOpenBankingConsent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

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
| `is_connected_account` | `bool` |  |
| `is_enabled` | `bool` |  |
| `merchant_id` | `string` |  |
| `provider` | `string` |  |
| `success_web_hook_url` | `string` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsOpenBankingConsent, err := client.NoFrixionMoneyMoovModelsOpenBankingConsent(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_open_banking_consent_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsOpenBankingConsent) // the loaded record
```

#### Example: List

```go
noFrixionMoneyMoovModelsOpenBankingConsents, err := client.NoFrixionMoneyMoovModelsOpenBankingConsent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsOpenBankingConsents) // the array of records
```

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovModelsOpenBankingConsent(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovModelsOpenBankingTransaction

Create an instance: `noFrixionMoneyMoovModelsOpenBankingTransaction := client.NoFrixionMoneyMoovModelsOpenBankingTransaction(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address_detail` | `map[string]any` |  |
| `amount` | `float64` |  |
| `balance` | `map[string]any` |  |
| `booking_date_time` | `string` |  |
| `charge_detail` | `map[string]any` |  |
| `currency` | `string` |  |
| `currency_exchange` | `map[string]any` |  |
| `date` | `string` |  |
| `description` | `string` |  |
| `enrichment` | `map[string]any` |  |
| `gross_amount` | `map[string]any` |  |
| `id` | `string` |  |
| `iso_bank_transaction_code` | `map[string]any` |  |
| `merchant` | `map[string]any` |  |
| `payee_detail` | `map[string]any` |  |
| `payer_detail` | `map[string]any` |  |
| `proprietary_bank_transaction_code` | `map[string]any` |  |
| `reference` | `string` |  |
| `statement_reference` | `[]any` |  |
| `status` | `string` |  |
| `supplementary_data` | `any` |  |
| `transaction_amount` | `map[string]any` |  |
| `transaction_information` | `[]any` |  |
| `transaction_mutability` | `string` |  |
| `value_date_time` | `string` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsOpenBankingTransactions, err := client.NoFrixionMoneyMoovModelsOpenBankingTransaction(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsOpenBankingTransactions) // the array of records
```


### NoFrixionMoneyMoovModelsPayment

Create an instance: `noFrixionMoneyMoovModelsPayment := client.NoFrixionMoneyMoovModelsPayment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `address` | `[]any` |  |
| `amount` | `float64` |  |
| `amount_pending` | `float64` |  |
| `amount_received` | `float64` |  |
| `amount_refunded` | `float64` |  |
| `auto_send_receipt` | `bool` |  |
| `base_origin_url` | `string` |  |
| `callback_url` | `string` |  |
| `card_authorize_only` | `bool` |  |
| `card_create_token` | `bool` |  |
| `card_create_token_mode` | `string` |  |
| `card_ignore_cvn` | `bool` |  |
| `card_no_payer_authentication` | `bool` |  |
| `card_processor_merchant_id` | `string` |  |
| `card_stripe_payment_intent_id` | `string` |  |
| `card_stripe_payment_intent_secret` | `string` |  |
| `card_transmit_raw_detail` | `bool` |  |
| `created_by_user` | `map[string]any` |  |
| `currency` | `string` |  |
| `custom_field` | `[]any` |  |
| `customer_email_address` | `string` |  |
| `customer_id` | `string` |  |
| `customer_name` | `string` |  |
| `description` | `string` |  |
| `destination_account` | `map[string]any` |  |
| `direct_debit_payment` | `map[string]any` |  |
| `due_date` | `string` |  |
| `event` | `[]any` |  |
| `failure_callback_url` | `string` |  |
| `field_display_setting` | `[]any` |  |
| `formatted_amount` | `string` |  |
| `hosted_pay_checkout_url` | `string` |  |
| `id` | `string` |  |
| `ignore_address_verification` | `bool` |  |
| `inserted` | `string` |  |
| `inserted_sortable` | `string` |  |
| `is_archived` | `bool` |  |
| `jwk` | `string` |  |
| `last_updated` | `string` |  |
| `lightning_invoice` | `string` |  |
| `lightning_invoice_expires_at` | `string` |  |
| `merchant_direct_debit_mandate_id` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `notification_email_address` | `string` |  |
| `notification_role_i_d` | `[]any` |  |
| `order_id` | `string` |  |
| `partial_payment_method` | `string` |  |
| `partial_payment_step` | `string` |  |
| `payment_attempt` | `[]any` |  |
| `payment_method` | `[]any` |  |
| `payment_processor` | `string` |  |
| `payrun_id` | `string` |  |
| `pisp_account_id` | `string` |  |
| `priority_bank_id` | `string` |  |
| `result` | `map[string]any` |  |
| `sandbox_settle_delay_in_second` | `int` |  |
| `shipping_address` | `map[string]any` |  |
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
| `tag` | `[]any` |  |
| `tag_id` | `[]any` |  |
| `title` | `string` |  |
| `tokenised_card` | `[]any` |  |
| `transaction` | `[]any` |  |
| `use_hosted_payment_page` | `bool` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsPayment, err := client.NoFrixionMoneyMoovModelsPayment(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_payment_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsPayment) // the loaded record
```

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovModelsPayment(nil).Create(map[string]any{
    "created_by_user": /* map[string]any */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovModelsPaymentAccountMinimalPage

Create an instance: `noFrixionMoneyMoovModelsPaymentAccountMinimalPage := client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `string` |  |
| `available_balance` | `float64` |  |
| `balance` | `float64` |  |
| `balance_minor_unit` | `int` |  |
| `currency` | `string` |  |
| `id` | `string` |  |
| `identifier` | `map[string]any` |  |
| `is_archived` | `bool` |  |
| `is_connected_account` | `bool` |  |
| `merchant_id` | `string` |  |
| `submitted_payouts_balance` | `float64` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsPaymentAccountMinimalPages, err := client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsPaymentAccountMinimalPages) // the array of records
```


### NoFrixionMoneyMoovModelsPaymentAccountPage

Create an instance: `noFrixionMoneyMoovModelsPaymentAccountPage := client.NoFrixionMoneyMoovModelsPaymentAccountPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `string` |  |
| `account_supplier_name` | `string` |  |
| `available_balance` | `float64` |  |
| `available_balance_minor_unit` | `int` |  |
| `balance` | `float64` |  |
| `balance_minor_unit` | `int` |  |
| `bank_name` | `string` |  |
| `consent_id` | `string` |  |
| `created_by` | `map[string]any` |  |
| `created_by_display_name` | `string` |  |
| `currency` | `string` |  |
| `default_payment_rail` | `string` |  |
| `display_name` | `string` |  |
| `expiry_date` | `string` |  |
| `external_account_icon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `map[string]any` |  |
| `inserted` | `string` |  |
| `is_archived` | `bool` |  |
| `is_connected_account` | `bool` |  |
| `is_default` | `bool` |  |
| `is_trust_account` | `bool` |  |
| `is_virtual` | `bool` |  |
| `last_transaction` | `map[string]any` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `physical_account_id` | `string` |  |
| `rule` | `[]any` |  |
| `submitted_payouts_balance` | `float64` |  |
| `submitted_payouts_balance_minor_unit` | `int` |  |
| `summary` | `string` |  |
| `supplier_sepa_instant_status` | `string` |  |
| `xero_bank_feed_connection_status` | `string` |  |
| `xero_bank_feed_last_synced_at` | `string` |  |
| `xero_bank_feed_sync_last_failed_at` | `string` |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` |  |
| `xero_bank_feed_sync_status` | `string` |  |
| `xero_unsynchronised_transactions_count` | `int` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsPaymentAccountPages, err := client.NoFrixionMoneyMoovModelsPaymentAccountPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsPaymentAccountPages) // the array of records
```


### NoFrixionMoneyMoovModelsPaymentInitiation

Create an instance: `noFrixionMoneyMoovModelsPaymentInitiation := client.NoFrixionMoneyMoovModelsPaymentInitiation(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

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

```go
result, err := client.NoFrixionMoneyMoovModelsPaymentInitiation(nil).Create(map[string]any{
    "paymentrequest_id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovModelsPaymentRequestEvent

Create an instance: `noFrixionMoneyMoovModelsPaymentRequestEvent := client.NoFrixionMoneyMoovModelsPaymentRequestEvent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float64` |  |
| `apple_pay_transaction_id` | `string` |  |
| `card_authorization_response_id` | `string` |  |
| `card_expiry_month` | `int` |  |
| `card_expiry_year` | `int` |  |
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

```go
noFrixionMoneyMoovModelsPaymentRequestEvents, err := client.NoFrixionMoneyMoovModelsPaymentRequestEvent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsPaymentRequestEvents) // the array of records
```


### NoFrixionMoneyMoovModelsPaymentRequestMetric

Create an instance: `noFrixionMoneyMoovModelsPaymentRequestMetric := client.NoFrixionMoneyMoovModelsPaymentRequestMetric(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `all` | `int` |  |
| `authorized` | `int` |  |
| `paid` | `int` |  |
| `partially_paid` | `int` |  |
| `total_amounts_by_currency` | `map[string]any` |  |
| `unpaid` | `int` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsPaymentRequestMetric, err := client.NoFrixionMoneyMoovModelsPaymentRequestMetric(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsPaymentRequestMetric) // the loaded record
```


### NoFrixionMoneyMoovModelsPaymentRequestMinimal

Create an instance: `noFrixionMoneyMoovModelsPaymentRequestMinimal := client.NoFrixionMoneyMoovModelsPaymentRequestMinimal(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float64` |  |
| `amount_pending` | `float64` |  |
| `amount_received` | `float64` |  |
| `amount_refunded` | `float64` |  |
| `callback_url` | `string` |  |
| `card_stripe_payment_intent_secret` | `string` |  |
| `country_code` | `string` |  |
| `currency` | `string` |  |
| `custom_fields_to_display` | `[]any` |  |
| `description` | `string` |  |
| `due_date` | `string` |  |
| `field_display_setting` | `[]any` |  |
| `google_pay_merchant_id` | `string` |  |
| `id` | `string` |  |
| `jwk` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_logo_url_png` | `string` |  |
| `merchant_logo_url_svg` | `string` |  |
| `merchant_name` | `string` |  |
| `merchant_short_name` | `string` |  |
| `partial_payment_method` | `string` |  |
| `payment_attempt` | `[]any` |  |
| `payment_methods_list` | `[]any` |  |
| `payment_processor` | `string` |  |
| `payment_processor_key` | `string` |  |
| `pisp_error` | `string` |  |
| `priority_bank_id` | `string` |  |
| `status` | `string` |  |
| `stripe_account_id` | `string` |  |
| `title` | `string` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsPaymentRequestMinimals, err := client.NoFrixionMoneyMoovModelsPaymentRequestMinimal(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsPaymentRequestMinimals) // the array of records
```


### NoFrixionMoneyMoovModelsPaymentRequestResult

Create an instance: `noFrixionMoneyMoovModelsPaymentRequestResult := client.NoFrixionMoneyMoovModelsPaymentRequestResult(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float64` |  |
| `amount_pending` | `float64` |  |
| `amount_received` | `float64` |  |
| `amount_refunded` | `float64` |  |
| `currency` | `string` |  |
| `customer_id` | `string` |  |
| `payment` | `[]any` |  |
| `payment_request_id` | `string` |  |
| `pisp_authorization` | `[]any` |  |
| `requested_amount` | `float64` |  |
| `result` | `string` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsPaymentRequestResults, err := client.NoFrixionMoneyMoovModelsPaymentRequestResult(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsPaymentRequestResults) // the array of records
```


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment

Create an instance: `noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `description` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `template` | `map[string]any` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsPaymentRequestsMerchantPayments, err := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsPaymentRequestsMerchantPayments) // the array of records
```


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2

Create an instance: `noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2 := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `description` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `template` | `map[string]any` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2, err := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(nil).Load(map[string]any{"paymentrequest_id": "paymentrequest_id", "template_id": "template_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2) // the loaded record
```


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3

Create an instance: `noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3 := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `description` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `template` | `map[string]any` |  |


### NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4

Create an instance: `noFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4 := client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |


### NoFrixionMoneyMoovModelsPayoutKeysetPage

Create an instance: `noFrixionMoneyMoovModelsPayoutKeysetPage := client.NoFrixionMoneyMoovModelsPayoutKeysetPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `amount` | `float64` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `[]any` |  |
| `authorisation` | `[]any` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `map[string]any` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `map[string]any` |  |
| `document` | `[]any` |  |
| `event` | `[]any` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `float64` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `float64` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `rule` | `map[string]any` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float64` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `map[string]any` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `[]any` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `float64` |  |
| `transacted_fx_amount` | `float64` |  |
| `transacted_fx_rate` | `float64` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsPayoutKeysetPages, err := client.NoFrixionMoneyMoovModelsPayoutKeysetPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsPayoutKeysetPages) // the array of records
```


### NoFrixionMoneyMoovModelsPayoutMetric

Create an instance: `noFrixionMoneyMoovModelsPayoutMetric := client.NoFrixionMoneyMoovModelsPayoutMetric(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `all` | `float64` |  |
| `failed` | `float64` |  |
| `in_progress` | `float64` |  |
| `paid` | `float64` |  |
| `pending_approval` | `float64` |  |
| `scheduled` | `float64` |  |
| `total_amounts_by_currency` | `map[string]any` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsPayoutMetric, err := client.NoFrixionMoneyMoovModelsPayoutMetric(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsPayoutMetric) // the loaded record
```


### NoFrixionMoneyMoovModelsPayoutsPayoutsCreate

Create an instance: `noFrixionMoneyMoovModelsPayoutsPayoutsCreate := client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failed_payout` | `map[string]any` |  |
| `payout` | `[]any` |  |

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovModelsPayrun

Create an instance: `noFrixionMoneyMoovModelsPayrun := client.NoFrixionMoneyMoovModelsPayrun(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation` | `[]any` |  |
| `authorisation_date` | `string` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `string` |  |
| `can_authorise` | `bool` |  |
| `can_delete` | `bool` |  |
| `can_edit` | `bool` |  |
| `event` | `[]any` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice` | `[]any` |  |
| `invoices_minimal` | `[]any` |  |
| `is_archived` | `bool` |  |
| `last_updated` | `string` |  |
| `last_updated_by` | `map[string]any` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `payment` | `[]any` |  |
| `payout` | `[]any` |  |
| `payouts_count` | `int` |  |
| `reason` | `string` |  |
| `schedule_date` | `string` |  |
| `scheduled_date` | `string` |  |
| `source_account` | `[]any` |  |
| `status` | `string` |  |
| `total_eur` | `float64` |  |
| `total_gbp` | `float64` |  |
| `total_usd` | `float64` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsPayrun, err := client.NoFrixionMoneyMoovModelsPayrun(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_payrun_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsPayrun) // the loaded record
```

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovModelsPayrun(nil).Create(map[string]any{
    "id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovModelsReportResult

Create an instance: `noFrixionMoneyMoovModelsReportResult := client.NoFrixionMoneyMoovModelsReportResult(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `content` | `string` |  |
| `content_type` | `string` |  |
| `last_completed_at` | `string` |  |
| `merchant_id` | `string` |  |
| `report_name` | `string` |  |
| `report_type` | `string` |  |
| `statement_number` | `int` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsReportResult, err := client.NoFrixionMoneyMoovModelsReportResult(nil).Load(map[string]any{"id": 1, "report_id": "report_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsReportResult) // the loaded record
```


### NoFrixionMoneyMoovModelsRule

Create an instance: `noFrixionMoneyMoovModelsRule := client.NoFrixionMoneyMoovModelsRule(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account` | `map[string]any` |  |
| `account_id` | `string` |  |
| `approve_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `[]any` |  |
| `authorisation` | `[]any` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `can_authorise` | `bool` |  |
| `created_by` | `map[string]any` |  |
| `description` | `string` |  |
| `end_at` | `string` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `is_disabled` | `bool` |  |
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
| `sweep_action` | `map[string]any` |  |
| `time_zone_id` | `string` |  |
| `trigger_cron_expression` | `string` |  |
| `trigger_on_pay_in` | `bool` |  |
| `user_id` | `string` |  |
| `web_hook_secret` | `string` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsRule, err := client.NoFrixionMoneyMoovModelsRule(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_rule_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsRule) // the loaded record
```

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovModelsRule(nil).Create(map[string]any{
    "created_by": /* map[string]any */,
    "nonce": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovModelsTransaction

Create an instance: `noFrixionMoneyMoovModelsTransaction := client.NoFrixionMoneyMoovModelsTransaction(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `account_name` | `string` |  |
| `account_sequence_number` | `int` |  |
| `amount` | `float64` |  |
| `amount_minor_unit` | `int` |  |
| `balance` | `float64` |  |
| `balance_minor_unit` | `int` |  |
| `counterparty` | `map[string]any` |  |
| `counterparty_summary` | `string` |  |
| `currency` | `string` |  |
| `description` | `string` |  |
| `fx_amount` | `float64` |  |
| `fx_currency` | `string` |  |
| `fx_rate` | `float64` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `merchant_id` | `string` |  |
| `payment_request_custom_field` | `map[string]any` |  |
| `payment_request_id` | `string` |  |
| `payout_id` | `string` |  |
| `raw_reference` | `string` |  |
| `rule_id` | `string` |  |
| `tag` | `[]any` |  |
| `their_reference` | `string` |  |
| `transaction_date` | `string` |  |
| `type` | `string` |  |
| `virtual_iban` | `string` |  |
| `your_reference` | `string` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsTransaction, err := client.NoFrixionMoneyMoovModelsTransaction(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_transaction_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsTransaction) // the loaded record
```


### NoFrixionMoneyMoovModelsTransactionPage

Create an instance: `noFrixionMoneyMoovModelsTransactionPage := client.NoFrixionMoneyMoovModelsTransactionPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `account_name` | `string` |  |
| `account_sequence_number` | `int` |  |
| `amount` | `float64` |  |
| `amount_minor_unit` | `int` |  |
| `balance` | `float64` |  |
| `balance_minor_unit` | `int` |  |
| `content` | `[]any` |  |
| `counterparty` | `map[string]any` |  |
| `counterparty_summary` | `string` |  |
| `currency` | `string` |  |
| `description` | `string` |  |
| `fx_amount` | `float64` |  |
| `fx_currency` | `string` |  |
| `fx_rate` | `float64` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `merchant_id` | `string` |  |
| `page_number` | `int` |  |
| `page_size` | `int` |  |
| `payment_request_custom_field` | `map[string]any` |  |
| `payment_request_id` | `string` |  |
| `payout_id` | `string` |  |
| `raw_reference` | `string` |  |
| `rule_id` | `string` |  |
| `tag` | `[]any` |  |
| `their_reference` | `string` |  |
| `total_page` | `int` |  |
| `total_size` | `int` |  |
| `transaction_date` | `string` |  |
| `type` | `string` |  |
| `virtual_iban` | `string` |  |
| `your_reference` | `string` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsTransactionPage, err := client.NoFrixionMoneyMoovModelsTransactionPage(nil).Load(map[string]any{"account_id": "account_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsTransactionPage) // the loaded record
```

#### Example: List

```go
noFrixionMoneyMoovModelsTransactionPages, err := client.NoFrixionMoneyMoovModelsTransactionPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsTransactionPages) // the array of records
```


### NoFrixionMoneyMoovModelsUserInvite

Create an instance: `noFrixionMoneyMoovModelsUserInvite := client.NoFrixionMoneyMoovModelsUserInvite(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_status` | `map[string]any` |  |
| `id` | `string` |  |
| `initial_role_id` | `string` |  |
| `invitee_email_address` | `string` |  |
| `invitee_first_name` | `string` |  |
| `invitee_last_name` | `string` |  |
| `inviter_email_address` | `string` |  |
| `inviter_first_name` | `string` |  |
| `inviter_last_name` | `string` |  |
| `is_authorised` | `bool` |  |
| `is_invitee_registered` | `bool` |  |
| `last_invited` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `message` | `string` |  |
| `registration_url` | `string` |  |
| `send_invite_email` | `bool` |  |
| `status` | `string` |  |
| `user` | `map[string]any` |  |
| `user_id` | `string` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsUserInvite, err := client.NoFrixionMoneyMoovModelsUserInvite(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_user_invite_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsUserInvite) // the loaded record
```

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovModelsUserInvite(nil).Create(map[string]any{
    "user": /* map[string]any */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### NoFrixionMoneyMoovModelsUserInvitePage

Create an instance: `noFrixionMoneyMoovModelsUserInvitePage := client.NoFrixionMoneyMoovModelsUserInvitePage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisation_status` | `map[string]any` |  |
| `id` | `string` |  |
| `initial_role_id` | `string` |  |
| `invitee_email_address` | `string` |  |
| `invitee_first_name` | `string` |  |
| `invitee_last_name` | `string` |  |
| `inviter_email_address` | `string` |  |
| `inviter_first_name` | `string` |  |
| `inviter_last_name` | `string` |  |
| `is_authorised` | `bool` |  |
| `is_invitee_registered` | `bool` |  |
| `last_invited` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `message` | `string` |  |
| `registration_url` | `string` |  |
| `status` | `string` |  |
| `user` | `map[string]any` |  |
| `user_id` | `string` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsUserInvitePages, err := client.NoFrixionMoneyMoovModelsUserInvitePage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsUserInvitePages) // the array of records
```


### NoFrixionMoneyMoovModelsUserPage

Create an instance: `noFrixionMoneyMoovModelsUserPage := client.NoFrixionMoneyMoovModelsUserPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client_session_timeout` | `[]any` |  |
| `email_address` | `string` |  |
| `first_name` | `string` |  |
| `id` | `string` |  |
| `last_name` | `string` |  |
| `passkey_added` | `bool` |  |
| `permission` | `map[string]any` |  |
| `roles_with_scope` | `[]any` |  |
| `two_factor_enabled` | `bool` |  |

#### Example: List

```go
noFrixionMoneyMoovModelsUserPages, err := client.NoFrixionMoneyMoovModelsUserPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsUserPages) // the array of records
```


### NoFrixionMoneyMoovModelsWebhook

Create an instance: `noFrixionMoneyMoovModelsWebhook := client.NoFrixionMoneyMoovModelsWebhook(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `destination_url` | `string` |  |
| `email_address` | `string` |  |
| `failed_notification_email_address` | `string` |  |
| `id` | `string` |  |
| `is_active` | `bool` |  |
| `merchant_id` | `string` |  |
| `notification_method` | `string` |  |
| `resource_type` | `[]any` |  |
| `retry` | `bool` |  |
| `secret` | `string` |  |
| `version` | `int` |  |

#### Example: Load

```go
noFrixionMoneyMoovModelsWebhook, err := client.NoFrixionMoneyMoovModelsWebhook(nil).Load(map[string]any{"id": "no_frixion_money_moov_models_webhook_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsWebhook) // the loaded record
```

#### Example: List

```go
noFrixionMoneyMoovModelsWebhooks, err := client.NoFrixionMoneyMoovModelsWebhook(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionMoneyMoovModelsWebhooks) // the array of records
```

#### Example: Create

```go
result, err := client.NoFrixionMoneyMoovModelsWebhook(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### OpenBanking

Create an instance: `openBanking := client.OpenBanking(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Example: Create

```go
result, err := client.OpenBanking(nil).Create(map[string]any{
    "account_id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Payeeverification

Create an instance: `payeeverification := client.Payeeverification(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

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

```go
result, err := client.Payeeverification(nil).Create(map[string]any{
    "account_name": /* string */,
    "iban": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PaymentRequest

Create an instance: `paymentRequest := client.PaymentRequest(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float64` |  |
| `do_simulate_settlement_failure` | `bool` |  |
| `error_description` | `string` |  |
| `institution` | `string` |  |
| `payment_initiation_id` | `string` |  |

#### Example: Load

```go
paymentRequest, err := client.PaymentRequest(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(paymentRequest) // the loaded record
```

#### Example: Create

```go
result, err := client.PaymentRequest(nil).Create(map[string]any{
    "paymentrequest_id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Payout

Create an instance: `payout := client.Payout(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `allow_incomplete` | `bool` |  |
| `amount` | `float64` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `[]any` |  |
| `authorisation` | `[]any` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `map[string]any` |  |
| `beneficiary_id` | `string` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `map[string]any` |  |
| `document` | `[]any` |  |
| `event` | `[]any` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `float64` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `float64` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `rule` | `map[string]any` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float64` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `map[string]any` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `[]any` |  |
| `tag_id` | `[]any` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `float64` |  |
| `transacted_fx_amount` | `float64` |  |
| `transacted_fx_rate` | `float64` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |

#### Example: Load

```go
payout, err := client.Payout(nil).Load(map[string]any{"id": "payout_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(payout) // the loaded record
```

#### Example: Create

```go
result, err := client.Payout(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Payrun

Create an instance: `payrun := client.Payrun(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |
| `note` | `string` |  |
| `scheduled_date` | `string` |  |

#### Example: Create

```go
result, err := client.Payrun(nil).Create(map[string]any{
    "id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Reject

Create an instance: `reject := client.Reject(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `amount` | `float64` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `[]any` |  |
| `authorisation` | `[]any` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `map[string]any` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `map[string]any` |  |
| `document` | `[]any` |  |
| `event` | `[]any` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `float64` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `float64` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `reason` | `string` |  |
| `rule` | `map[string]any` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float64` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `map[string]any` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `[]any` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `float64` |  |
| `transacted_fx_amount` | `float64` |  |
| `transacted_fx_rate` | `float64` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |


### Report

Create an instance: `report := client.Report(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |


### Rule

Create an instance: `rule := client.Rule(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |


### Send

Create an instance: `send := client.Send(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `allow_incomplete` | `bool` |  |
| `amount` | `float64` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `[]any` |  |
| `authorisation` | `[]any` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `map[string]any` |  |
| `beneficiary_id` | `string` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `map[string]any` |  |
| `document` | `[]any` |  |
| `event` | `[]any` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `float64` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `float64` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `rule` | `map[string]any` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float64` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `map[string]any` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `[]any` |  |
| `tag_id` | `[]any` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `float64` |  |
| `transacted_fx_amount` | `float64` |  |
| `transacted_fx_rate` | `float64` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |

#### Example: Create

```go
result, err := client.Send(nil).Create(map[string]any{
    "beneficiary": /* map[string]any */,
    "source_account_identifier": /* map[string]any */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Sendbeneficiary

Create an instance: `sendbeneficiary := client.Sendbeneficiary(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_id` | `string` |  |
| `allow_incomplete` | `bool` |  |
| `amount` | `float64` |  |
| `amount_minor_unit` | `int` |  |
| `approve_payout_url` | `string` |  |
| `approver_id` | `string` |  |
| `authentication_method` | `[]any` |  |
| `authorisation` | `[]any` |  |
| `authorisers_completed_count` | `int` |  |
| `authorisers_required_count` | `int` |  |
| `batch_payout_id` | `string` |  |
| `beneficiary` | `map[string]any` |  |
| `beneficiary_id` | `string` |  |
| `can_authorise` | `bool` |  |
| `can_process` | `bool` |  |
| `can_update` | `bool` |  |
| `charge_bearer` | `string` |  |
| `created_by` | `string` |  |
| `created_by_email_address` | `string` |  |
| `currency` | `string` |  |
| `current_user_id` | `string` |  |
| `description` | `string` |  |
| `destination` | `map[string]any` |  |
| `document` | `[]any` |  |
| `event` | `[]any` |  |
| `formatted_amount` | `string` |  |
| `formatted_fx_destination_amount` | `string` |  |
| `formatted_schedule` | `string` |  |
| `formatted_schedule_day_only` | `string` |  |
| `formatted_source_account_available_balance` | `string` |  |
| `fx_destination_amount` | `float64` |  |
| `fx_destination_amount_minor_unit` | `int` |  |
| `fx_destination_currency` | `string` |  |
| `fx_quote_expires_at` | `string` |  |
| `fx_quote_id` | `string` |  |
| `fx_rate` | `float64` |  |
| `fx_use_destination_amount` | `bool` |  |
| `has_current_user_authorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoice_id` | `string` |  |
| `is_archived` | `bool` |  |
| `is_failed` | `bool` |  |
| `is_settled` | `bool` |  |
| `is_submitted` | `bool` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_token_description` | `string` |  |
| `nonce` | `string` |  |
| `payment_processor` | `string` |  |
| `payment_rail` | `string` |  |
| `payrun_id` | `string` |  |
| `payrun_name` | `string` |  |
| `rule` | `map[string]any` |  |
| `schedule_date` | `string` |  |
| `scheduled` | `bool` |  |
| `source_account_available_balance` | `float64` |  |
| `source_account_available_balance_minor_unit` | `int` |  |
| `source_account_bic` | `string` |  |
| `source_account_currency` | `string` |  |
| `source_account_iban` | `string` |  |
| `source_account_identifier` | `map[string]any` |  |
| `source_account_name` | `string` |  |
| `source_account_number` | `string` |  |
| `source_account_sortcode` | `string` |  |
| `status` | `string` |  |
| `tag` | `[]any` |  |
| `tag_id` | `[]any` |  |
| `their_reference` | `string` |  |
| `topup_payrun_id` | `string` |  |
| `transacted_amount` | `float64` |  |
| `transacted_fx_amount` | `float64` |  |
| `transacted_fx_rate` | `float64` |  |
| `type` | `string` |  |
| `user_id` | `string` |  |
| `your_reference` | `string` |  |

#### Example: Create

```go
result, err := client.Sendbeneficiary(nil).Create(map[string]any{
    "beneficiary": /* map[string]any */,
    "source_account_identifier": /* map[string]any */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Tag

Create an instance: `tag := client.Tag(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `colour_hex` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `merchant_id` | `string` |  |
| `name` | `string` |  |

#### Example: List

```go
tags, err := client.Tag(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(tags) // the array of records
```

#### Example: Create

```go
result, err := client.Tag(nil).Create(map[string]any{
    "merchant_id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Token

Create an instance: `token := client.Token(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Example: Create

```go
result, err := client.Token(nil).Create(map[string]any{
    "id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Transaction

Create an instance: `transaction := client.Transaction(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Example: Load

```go
transaction, err := client.Transaction(nil).Load(map[string]any{"sequence_number": 1, "transaction_id": "transaction_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(transaction) // the loaded record
```

#### Example: Create

```go
result, err := client.Transaction(nil).Create(map[string]any{
    "id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### User

Create an instance: `user := client.User(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client_session_timeout` | `[]any` |  |
| `email_address` | `string` |  |
| `first_name` | `string` |  |
| `id` | `string` |  |
| `last_name` | `string` |  |
| `passkey_added` | `bool` |  |
| `permission` | `map[string]any` |  |
| `profile` | `string` |  |
| `roles_with_scope` | `[]any` |  |
| `two_factor_enabled` | `bool` |  |
| `user_invite_id` | `string` |  |

#### Example: List

```go
users, err := client.User(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(users) // the array of records
```


### UserInvite

Create an instance: `userInvite := client.UserInvite(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Example: Create

```go
result, err := client.UserInvite(nil).Create(map[string]any{
    "id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Virtual

Create an instance: `virtual := client.Virtual(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `account_name` | `string` |  |
| `account_supplier_name` | `string` |  |
| `available_balance` | `float64` |  |
| `available_balance_minor_unit` | `int` |  |
| `balance` | `float64` |  |
| `balance_minor_unit` | `int` |  |
| `bank_name` | `string` |  |
| `consent_id` | `string` |  |
| `created_by` | `map[string]any` |  |
| `created_by_display_name` | `string` |  |
| `currency` | `string` |  |
| `default_payment_rail` | `string` |  |
| `display_name` | `string` |  |
| `expiry_date` | `string` |  |
| `external_account_icon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `map[string]any` |  |
| `inserted` | `string` |  |
| `is_archived` | `bool` |  |
| `is_connected_account` | `bool` |  |
| `is_default` | `bool` |  |
| `is_trust_account` | `bool` |  |
| `is_virtual` | `bool` |  |
| `last_transaction` | `map[string]any` |  |
| `last_updated` | `string` |  |
| `merchant_id` | `string` |  |
| `merchant_name` | `string` |  |
| `name` | `string` |  |
| `physical_account_id` | `string` |  |
| `rule` | `[]any` |  |
| `submitted_payouts_balance` | `float64` |  |
| `submitted_payouts_balance_minor_unit` | `int` |  |
| `summary` | `string` |  |
| `supplier_sepa_instant_status` | `string` |  |
| `xero_bank_feed_connection_status` | `string` |  |
| `xero_bank_feed_last_synced_at` | `string` |  |
| `xero_bank_feed_sync_last_failed_at` | `string` |  |
| `xero_bank_feed_sync_last_failure_reason` | `string` |  |
| `xero_bank_feed_sync_status` | `string` |  |
| `xero_unsynchronised_transactions_count` | `int` |  |

#### Example: Create

```go
result, err := client.Virtual(nil).Create(map[string]any{
    "account_id": /* string */,
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Webhook

Create an instance: `webhook := client.Webhook(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |


### Whoami

Create an instance: `whoami := client.Whoami(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client_session_timeout` | `[]any` |  |
| `email_address` | `string` |  |
| `first_name` | `string` |  |
| `id` | `string` |  |
| `last_name` | `string` |  |
| `passkey_added` | `bool` |  |
| `permission` | `map[string]any` |  |
| `roles_with_scope` | `[]any` |  |
| `two_factor_enabled` | `bool` |  |

#### Example: List

```go
whoamis, err := client.Whoami(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(whoamis) // the array of records
```


### Whoamitrustedapp

Create an instance: `whoamitrustedapp := client.Whoamitrustedapp(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `client_session_timeout` | `[]any` |  |
| `email_address` | `string` |  |
| `first_name` | `string` |  |
| `id` | `string` |  |
| `last_name` | `string` |  |
| `passkey_added` | `bool` |  |
| `permission` | `map[string]any` |  |
| `roles_with_scope` | `[]any` |  |
| `two_factor_enabled` | `bool` |  |

#### Example: List

```go
whoamitrustedapps, err := client.Whoamitrustedapp(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(whoamitrustedapps) // the array of records
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

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/nofrixion-sdk/go/
├── nofrixion.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/nofrixion-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `List`, the entity
stores the returned data and match criteria internally.

```go
account := client.Account(nil)
account.List(nil, nil)

// account.Data() now returns the account data from the last list
// account.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
