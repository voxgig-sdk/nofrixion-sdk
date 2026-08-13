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
    created, err := client.Account(nil).Create(map[string]any{"createdBy": map[string]any{}, "identifier": map[string]any{}}, nil)
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
consents, err := client.Consent(nil).List(nil, nil)
if err != nil {
    // handle err
    return
}
_ = consents
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

consent, err := client.Consent(nil).List(
    nil, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(consent) // the returned mock data
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
| `Batch` | `(data map[string]any) NofrixionEntity` | Create a Batch entity instance. |
| `Beneficiary` | `(data map[string]any) NofrixionEntity` | Create a Beneficiary entity instance. |
| `BeneficiaryGroup` | `(data map[string]any) NofrixionEntity` | Create a BeneficiaryGroup entity instance. |
| `Card` | `(data map[string]any) NofrixionEntity` | Create a Card entity instance. |
| `CardCustomerToken` | `(data map[string]any) NofrixionEntity` | Create a CardCustomerToken entity instance. |
| `CardPayment` | `(data map[string]any) NofrixionEntity` | Create a CardPayment entity instance. |
| `CardPublicKey` | `(data map[string]any) NofrixionEntity` | Create a CardPublicKey entity instance. |
| `Consent` | `(data map[string]any) NofrixionEntity` | Create a Consent entity instance. |
| `Currency` | `(data map[string]any) NofrixionEntity` | Create a Currency entity instance. |
| `DirectDebitBatchSubmit` | `(data map[string]any) NofrixionEntity` | Create a DirectDebitBatchSubmit entity instance. |
| `FxRate` | `(data map[string]any) NofrixionEntity` | Create a FxRate entity instance. |
| `IPayment` | `(data map[string]any) NofrixionEntity` | Create an IPayment entity instance. |
| `Mandate` | `(data map[string]any) NofrixionEntity` | Create a Mandate entity instance. |
| `Merchant` | `(data map[string]any) NofrixionEntity` | Create a Merchant entity instance. |
| `MerchantAuthorisationSetting` | `(data map[string]any) NofrixionEntity` | Create a MerchantAuthorisationSetting entity instance. |
| `MerchantDirectDebitMandatePage` | `(data map[string]any) NofrixionEntity` | Create a MerchantDirectDebitMandatePage entity instance. |
| `MerchantPayByBankSetting` | `(data map[string]any) NofrixionEntity` | Create a MerchantPayByBankSetting entity instance. |
| `MerchantPaymentRequestTemplate` | `(data map[string]any) NofrixionEntity` | Create a MerchantPaymentRequestTemplate entity instance. |
| `MerchantToken` | `(data map[string]any) NofrixionEntity` | Create a MerchantToken entity instance. |
| `Metadata` | `(data map[string]any) NofrixionEntity` | Create a Metadata entity instance. |
| `NoFrixionVersion` | `(data map[string]any) NofrixionEntity` | Create a NoFrixionVersion entity instance. |
| `OpenBanking` | `(data map[string]any) NofrixionEntity` | Create an OpenBanking entity instance. |
| `Payeeverification` | `(data map[string]any) NofrixionEntity` | Create a Payeeverification entity instance. |
| `Payment` | `(data map[string]any) NofrixionEntity` | Create a Payment entity instance. |
| `PaymentAccount` | `(data map[string]any) NofrixionEntity` | Create a PaymentAccount entity instance. |
| `PaymentAccountMinimal` | `(data map[string]any) NofrixionEntity` | Create a PaymentAccountMinimal entity instance. |
| `PaymentInitiation` | `(data map[string]any) NofrixionEntity` | Create a PaymentInitiation entity instance. |
| `PaymentRequest` | `(data map[string]any) NofrixionEntity` | Create a PaymentRequest entity instance. |
| `PaymentRequestEvent` | `(data map[string]any) NofrixionEntity` | Create a PaymentRequestEvent entity instance. |
| `PaymentRequestMetric` | `(data map[string]any) NofrixionEntity` | Create a PaymentRequestMetric entity instance. |
| `PaymentRequestMinimal` | `(data map[string]any) NofrixionEntity` | Create a PaymentRequestMinimal entity instance. |
| `PaymentRequestResult` | `(data map[string]any) NofrixionEntity` | Create a PaymentRequestResult entity instance. |
| `Payout` | `(data map[string]any) NofrixionEntity` | Create a Payout entity instance. |
| `PayoutKeysetPage` | `(data map[string]any) NofrixionEntity` | Create a PayoutKeysetPage entity instance. |
| `PayoutMetric` | `(data map[string]any) NofrixionEntity` | Create a PayoutMetric entity instance. |
| `Payrun` | `(data map[string]any) NofrixionEntity` | Create a Payrun entity instance. |
| `Report` | `(data map[string]any) NofrixionEntity` | Create a Report entity instance. |
| `ReportResult` | `(data map[string]any) NofrixionEntity` | Create a ReportResult entity instance. |
| `Role` | `(data map[string]any) NofrixionEntity` | Create a Role entity instance. |
| `Rule` | `(data map[string]any) NofrixionEntity` | Create a Rule entity instance. |
| `RuleEvent` | `(data map[string]any) NofrixionEntity` | Create a RuleEvent entity instance. |
| `Tag` | `(data map[string]any) NofrixionEntity` | Create a Tag entity instance. |
| `Token` | `(data map[string]any) NofrixionEntity` | Create a Token entity instance. |
| `Transaction` | `(data map[string]any) NofrixionEntity` | Create a Transaction entity instance. |
| `User` | `(data map[string]any) NofrixionEntity` | Create an User entity instance. |
| `UserInvite` | `(data map[string]any) NofrixionEntity` | Create an UserInvite entity instance. |
| `Virtual` | `(data map[string]any) NofrixionEntity` | Create a Virtual entity instance. |
| `Webhook` | `(data map[string]any) NofrixionEntity` | Create a Webhook entity instance. |

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
| `"accountBalances"` |  |
| `"accountID"` |  |
| `"accountIdentifications"` |  |
| `"accountName"` |  |
| `"accountNames"` |  |
| `"accountSupplierName"` |  |
| `"accountType"` |  |
| `"availableBalance"` |  |
| `"availableBalanceMinorUnits"` |  |
| `"balance"` |  |
| `"balanceMinorUnits"` |  |
| `"bankName"` |  |
| `"consentID"` |  |
| `"consolidatedAccountInformation"` |  |
| `"createdBy"` |  |
| `"createdByDisplayName"` |  |
| `"currency"` |  |
| `"defaultPaymentRail"` |  |
| `"description"` |  |
| `"details"` |  |
| `"displayName"` |  |
| `"expiryDate"` |  |
| `"externalAccountIcon"` |  |
| `"format"` |  |
| `"fromDate"` |  |
| `"id"` |  |
| `"identifier"` |  |
| `"inserted"` |  |
| `"isArchived"` |  |
| `"isConnectedAccount"` |  |
| `"isDefault"` |  |
| `"isTrustAccount"` |  |
| `"isVirtual"` |  |
| `"lastTransaction"` |  |
| `"lastUpdated"` |  |
| `"merchantID"` |  |
| `"merchantName"` |  |
| `"nickname"` |  |
| `"physicalAccountID"` |  |
| `"roleIDs"` |  |
| `"rules"` |  |
| `"submittedPayoutsBalance"` |  |
| `"submittedPayoutsBalanceMinorUnits"` |  |
| `"summary"` |  |
| `"supplierPhysicalAccountID"` |  |
| `"supplierSepaInstantStatus"` |  |
| `"toDate"` |  |
| `"type"` |  |
| `"usageType"` |  |
| `"xeroBankFeedConnectionStatus"` |  |
| `"xeroBankFeedLastSyncedAt"` |  |
| `"xeroBankFeedSyncLastFailedAt"` |  |
| `"xeroBankFeedSyncLastFailureReason"` |  |
| `"xeroBankFeedSyncStatus"` |  |
| `"xeroUnsynchronisedTransactionsCount"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/accounts/{accountID}/{currency}`

#### Batch

| Field | Description |
| --- | --- |
| `"approveUrl"` |  |
| `"id"` |  |
| `"payouts"` |  |

Operations: Create, Load.

API path: `/api/v1/payouts/batch`

#### Beneficiary

| Field | Description |
| --- | --- |
| `"approvalCallbackUrl"` |  |
| `"authenticationMethods"` |  |
| `"authorisations"` |  |
| `"authorisersCompletedCount"` |  |
| `"authorisersRequiredCount"` |  |
| `"beneficiaries"` |  |
| `"beneficiaryEvents"` |  |
| `"canAuthorise"` |  |
| `"canUpdate"` |  |
| `"createdBy"` |  |
| `"createdByEmailAddress"` |  |
| `"currency"` |  |
| `"destination"` |  |
| `"failedBeneficiaries"` |  |
| `"hasCurrentUserAuthorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"isEnabled"` |  |
| `"lastAuthorised"` |  |
| `"lastUpdated"` |  |
| `"merchantID"` |  |
| `"name"` |  |
| `"nonce"` |  |
| `"sourceAccountIDs"` |  |
| `"sourceAccounts"` |  |
| `"theirReference"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/beneficiaries/authorise/{id}`

#### BeneficiaryGroup

| Field | Description |
| --- | --- |
| `"groupMembers"` |  |
| `"groupName"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"lastUpdated"` |  |
| `"merchantID"` |  |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/beneficiarygroups`

#### Card

| Field | Description |
| --- | --- |
| `"authorizedAmount"` |  |
| `"currencyCode"` |  |
| `"isPayerAuthenticationRequired"` |  |
| `"isSoftDecline"` |  |
| `"payerAuthenticationAccessToken"` |  |
| `"payerAuthenticationMerchantData"` |  |
| `"payerAuthenticationUrl"` |  |
| `"payerAuthenticationWindowHeight"` |  |
| `"payerAuthenticationWindowWidth"` |  |
| `"paymentRequestCallbackUrl"` |  |
| `"paymentRequestID"` |  |
| `"requestID"` |  |
| `"responseCode"` |  |
| `"responseType"` |  |
| `"status"` |  |
| `"threeDSRedirectUrl"` |  |
| `"transactionID"` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/{id}/card`

#### CardCustomerToken

| Field | Description |
| --- | --- |
| `"cardType"` |  |
| `"customerEmailAddress"` |  |
| `"expiryMonth"` |  |
| `"expiryYear"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"lastFourDigits"` |  |
| `"lastUpdated"` |  |
| `"maskedCardNumber"` |  |
| `"merchantID"` |  |
| `"paymentRequestID"` |  |

Operations: List, Load, Remove.

API path: `/api/v1/paymentrequests/card/customertokens/{merchantID}/{customerEmailAddress}`

#### CardPayment

| Field | Description |
| --- | --- |
| `"authorizedAmount"` |  |
| `"currencyCode"` |  |
| `"isPayerAuthenticationRequired"` |  |
| `"isSoftDecline"` |  |
| `"payerAuthenticationAccessToken"` |  |
| `"payerAuthenticationMerchantData"` |  |
| `"payerAuthenticationUrl"` |  |
| `"payerAuthenticationWindowHeight"` |  |
| `"payerAuthenticationWindowWidth"` |  |
| `"paymentRequestCallbackUrl"` |  |
| `"paymentRequestID"` |  |
| `"requestID"` |  |
| `"responseCode"` |  |
| `"responseType"` |  |
| `"status"` |  |
| `"threeDSRedirectUrl"` |  |
| `"transactionID"` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/{id}/card/refund/{partialRefundAmount}`

#### CardPublicKey

| Field | Description |
| --- | --- |
| `"jwt"` |  |

Operations: Load.

API path: `/api/v1/paymentrequests/{id}/card/publickey`

#### Consent

| Field | Description |
| --- | --- |
| `"authorisationUrl"` |  |
| `"callbackUrl"` |  |
| `"consentID"` |  |
| `"emailAddress"` |  |
| `"expiryDate"` |  |
| `"failureCallbackUrl"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"institutionID"` |  |
| `"isConnectedAccounts"` |  |
| `"isEnabled"` |  |
| `"merchantID"` |  |
| `"provider"` |  |
| `"successWebHookUrl"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/openbanking/consents`

#### Currency

| Field | Description |
| --- | --- |
| `"code"` |  |
| `"decimals"` |  |
| `"isFiat"` |  |
| `"iso4217AlphaCode"` |  |
| `"iso4217NumericCode"` |  |
| `"symbol"` |  |

Operations: List.

API path: `/api/v1/currencies`

#### DirectDebitBatchSubmit

| Field | Description |
| --- | --- |
| `"failedSubmissions"` |  |
| `"successfulSubmissions"` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/directdebit/batchsubmit`

#### FxRate

| Field | Description |
| --- | --- |
| `"destinationCurrency"` |  |
| `"exchangeRate"` |  |
| `"expiryTime"` |  |
| `"quoteID"` |  |
| `"sourceCurrency"` |  |

Operations: List, Load.

API path: `/api/v1/payouts/fxallheldrates/{source}/{destination}`

#### IPayment

| Field | Description |
| --- | --- |
| `"paymentRequestID"` |  |
| `"responseType"` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/payondemand`

#### Mandate

| Field | Description |
| --- | --- |
| `"accountNumber"` |  |
| `"addressLine1"` |  |
| `"addressLine2"` |  |
| `"approvedAt"` |  |
| `"city"` |  |
| `"countryCode"` |  |
| `"currency"` |  |
| `"customerAccountNumber"` |  |
| `"customerCity"` |  |
| `"customerCountryCode"` |  |
| `"customerCountryName"` |  |
| `"customerEmailAddress"` |  |
| `"customerFirstName"` |  |
| `"customerIban"` |  |
| `"customerLastName"` |  |
| `"customerSortCode"` |  |
| `"emailAddress"` |  |
| `"firstName"` |  |
| `"iban"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"isRecurring"` |  |
| `"lastName"` |  |
| `"lastUpdated"` |  |
| `"merchantID"` |  |
| `"postalCode"` |  |
| `"reference"` |  |
| `"sortCode"` |  |
| `"status"` |  |
| `"supplierBankAccountID"` |  |
| `"supplierCustomerID"` |  |
| `"supplierMandateID"` |  |
| `"supplierName"` |  |
| `"supplierStatus"` |  |

Operations: Create, Load.

API path: `/api/v1/mandates`

#### Merchant

| Field | Description |
| --- | --- |
| `"accountCurrencies"` |  |
| `"canHaveTrustAccounts"` |  |
| `"cardPaymentProcessor"` |  |
| `"companyID"` |  |
| `"displayQrOnHostedPay"` |  |
| `"hostedPayVersion"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"isBlocked"` |  |
| `"isExited"` |  |
| `"isSuspended"` |  |
| `"jurisdiction"` |  |
| `"logoUrlPng"` |  |
| `"logoUrlSvg"` |  |
| `"merchantCategoryCode"` |  |
| `"name"` |  |
| `"notes"` |  |
| `"parentMerchant"` |  |
| `"paymentAccountLimit"` |  |
| `"paymentAccounts"` |  |
| `"reason"` |  |
| `"shortName"` |  |
| `"supportedPaymentMethodsList"` |  |
| `"suspensionReason"` |  |
| `"tags"` |  |
| `"timeZoneId"` |  |
| `"tradingName"` |  |
| `"webHookLimit"` |  |
| `"yourRoleName"` |  |

Operations: List, Load, Remove, Update.

API path: `/api/v1/merchants/{merchantID}/childmerchants`

#### MerchantAuthorisationSetting

| Field | Description |
| --- | --- |
| `"amountLower"` |  |
| `"amountUpper"` |  |
| `"authorisationType"` |  |
| `"beneficiariesOnly"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"lastEditorCantAuthorise"` |  |
| `"lastUpdated"` |  |
| `"merchantID"` |  |
| `"numberOfAuthorisers"` |  |
| `"roleSettings"` |  |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/authorisationsettings`

#### MerchantDirectDebitMandatePage

| Field | Description |
| --- | --- |
| `"approvedAt"` |  |
| `"currency"` |  |
| `"customerAccountNumber"` |  |
| `"customerCity"` |  |
| `"customerCountryCode"` |  |
| `"customerCountryName"` |  |
| `"customerEmailAddress"` |  |
| `"customerFirstName"` |  |
| `"customerIban"` |  |
| `"customerLastName"` |  |
| `"customerSortCode"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"isRecurring"` |  |
| `"lastUpdated"` |  |
| `"merchantID"` |  |
| `"reference"` |  |
| `"status"` |  |
| `"supplierBankAccountID"` |  |
| `"supplierCustomerID"` |  |
| `"supplierMandateID"` |  |
| `"supplierName"` |  |
| `"supplierStatus"` |  |

Operations: List.

API path: `/api/v1/mandates`

#### MerchantPayByBankSetting

| Field | Description |
| --- | --- |
| `"bankCountryCodes"` |  |
| `"bankID"` |  |
| `"bankName"` |  |
| `"businessInstitutionID"` |  |
| `"currency"` |  |
| `"logo"` |  |
| `"message"` |  |
| `"messageImageUrl"` |  |
| `"order"` |  |
| `"personalInstitutionID"` |  |
| `"processor"` |  |
| `"warningHeading"` |  |
| `"warningMessage"` |  |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/banksettings`

#### MerchantPaymentRequestTemplate

| Field | Description |
| --- | --- |
| `"bankPaymentOptions"` |  |
| `"cardPaymentAddressOptions"` |  |
| `"cardPaymentCaptureOptions"` |  |
| `"customFields"` |  |
| `"defaultFields"` |  |
| `"description"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"lastUpdated"` |  |
| `"merchantID"` |  |
| `"name"` |  |
| `"notificationOptions"` |  |
| `"paymentMethods"` |  |
| `"paymentTerms"` |  |
| `"priorityBankOptions"` |  |
| `"template"` |  |

Operations: List, Load, Remove, Update.

API path: `/api/v1/paymentrequests/{merchantID}/templates`

#### MerchantToken

| Field | Description |
| --- | --- |
| `"authenticationMethods"` |  |
| `"authorisations"` |  |
| `"authorisersCompletedCount"` |  |
| `"authorisersRequiredCount"` |  |
| `"canAuthorise"` |  |
| `"description"` |  |
| `"expiresAt"` |  |
| `"hasCurrentUserAuthorised"` |  |
| `"hmacAlgorithm"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"ipAddressWhitelist"` |  |
| `"isArchived"` |  |
| `"isEnabled"` |  |
| `"lastAuthorised"` |  |
| `"lastUpdated"` |  |
| `"merchantID"` |  |
| `"nonce"` |  |
| `"permissionTypes"` |  |
| `"requestSignatureVersion"` |  |
| `"sharedSecretAlgorithm"` |  |
| `"sharedSecretBase64"` |  |
| `"token"` |  |

Operations: Create, List, Load, Update.

API path: `/api/v1/tokens`

#### Metadata

| Field | Description |
| --- | --- |

Operations: Load.

API path: `/api/v1/metadata/problemnotification`

#### NoFrixionVersion

| Field | Description |
| --- | --- |
| `"buildVersion"` |  |
| `"majorVersion"` |  |
| `"minorVersion"` |  |
| `"releaseName"` |  |

Operations: Load.

API path: `/api/v1/metadata/version`

#### OpenBanking

| Field | Description |
| --- | --- |

Operations: Create, Remove.

API path: `/api/v1/openbanking/account/{accountID}/synchronise`

#### Payeeverification

| Field | Description |
| --- | --- |
| `"accountName"` |  |
| `"accountNumber"` |  |
| `"iban"` |  |
| `"payeeVerifiedAccountName"` |  |
| `"result"` |  |
| `"secondaryIdentification"` |  |
| `"sortCode"` |  |

Operations: Create.

API path: `/api/v1/openbanking/payeeverification`

#### Payment

| Field | Description |
| --- | --- |
| `"addresses"` |  |
| `"amount"` |  |
| `"amountPending"` |  |
| `"amountReceived"` |  |
| `"amountRefunded"` |  |
| `"autoSendReceipt"` |  |
| `"baseOriginUrl"` |  |
| `"callbackUrl"` |  |
| `"cardAuthorizeOnly"` |  |
| `"cardCreateToken"` |  |
| `"cardCreateTokenMode"` |  |
| `"cardIgnoreCVN"` |  |
| `"cardNoPayerAuthentication"` |  |
| `"cardProcessorMerchantID"` |  |
| `"cardStripePaymentIntentID"` |  |
| `"cardStripePaymentIntentSecret"` |  |
| `"cardTransmitRawDetails"` |  |
| `"createdByUser"` |  |
| `"currency"` |  |
| `"customFields"` |  |
| `"customerEmailAddress"` |  |
| `"customerID"` |  |
| `"customerName"` |  |
| `"description"` |  |
| `"destinationAccount"` |  |
| `"directDebitPayment"` |  |
| `"dueDate"` |  |
| `"events"` |  |
| `"failureCallbackUrl"` |  |
| `"fieldDisplaySettings"` |  |
| `"formattedAmount"` |  |
| `"hostedPayCheckoutUrl"` |  |
| `"id"` |  |
| `"ignoreAddressVerification"` |  |
| `"inserted"` |  |
| `"insertedSortable"` |  |
| `"isArchived"` |  |
| `"jwk"` |  |
| `"lastUpdated"` |  |
| `"lightningInvoice"` |  |
| `"lightningInvoiceExpiresAt"` |  |
| `"merchantDirectDebitMandateID"` |  |
| `"merchantID"` |  |
| `"merchantTokenDescription"` |  |
| `"notificationEmailAddresses"` |  |
| `"notificationRoleIDs"` |  |
| `"orderID"` |  |
| `"partialPaymentMethod"` |  |
| `"partialPaymentSteps"` |  |
| `"paymentAttempts"` |  |
| `"paymentMethods"` |  |
| `"paymentProcessor"` |  |
| `"payrunID"` |  |
| `"pispAccountID"` |  |
| `"priorityBankID"` |  |
| `"result"` |  |
| `"sandboxSettleDelayInSeconds"` |  |
| `"shippingAddress"` |  |
| `"shippingAddressCity"` |  |
| `"shippingAddressCountryCode"` |  |
| `"shippingAddressCounty"` |  |
| `"shippingAddressLine1"` |  |
| `"shippingAddressLine2"` |  |
| `"shippingAddressPostCode"` |  |
| `"shippingEmail"` |  |
| `"shippingFirstName"` |  |
| `"shippingLastName"` |  |
| `"shippingPhone"` |  |
| `"status"` |  |
| `"successWebHookUrl"` |  |
| `"tagIds"` |  |
| `"tags"` |  |
| `"title"` |  |
| `"tokenisedCards"` |  |
| `"transactions"` |  |
| `"useHostedPaymentPage"` |  |

Operations: Create, Load, Update.

API path: `/api/v1/paymentrequests`

#### PaymentAccount

| Field | Description |
| --- | --- |
| `"accountName"` |  |
| `"accountSupplierName"` |  |
| `"availableBalance"` |  |
| `"availableBalanceMinorUnits"` |  |
| `"balance"` |  |
| `"balanceMinorUnits"` |  |
| `"bankName"` |  |
| `"consentID"` |  |
| `"createdBy"` |  |
| `"createdByDisplayName"` |  |
| `"currency"` |  |
| `"defaultPaymentRail"` |  |
| `"displayName"` |  |
| `"expiryDate"` |  |
| `"externalAccountIcon"` |  |
| `"id"` |  |
| `"identifier"` |  |
| `"inserted"` |  |
| `"isArchived"` |  |
| `"isConnectedAccount"` |  |
| `"isDefault"` |  |
| `"isTrustAccount"` |  |
| `"isVirtual"` |  |
| `"lastTransaction"` |  |
| `"lastUpdated"` |  |
| `"merchantID"` |  |
| `"merchantName"` |  |
| `"physicalAccountID"` |  |
| `"rules"` |  |
| `"submittedPayoutsBalance"` |  |
| `"submittedPayoutsBalanceMinorUnits"` |  |
| `"summary"` |  |
| `"supplierSepaInstantStatus"` |  |
| `"xeroBankFeedConnectionStatus"` |  |
| `"xeroBankFeedLastSyncedAt"` |  |
| `"xeroBankFeedSyncLastFailedAt"` |  |
| `"xeroBankFeedSyncLastFailureReason"` |  |
| `"xeroBankFeedSyncStatus"` |  |
| `"xeroUnsynchronisedTransactionsCount"` |  |

Operations: List.

API path: `/api/v1/accounts/paged`

#### PaymentAccountMinimal

| Field | Description |
| --- | --- |
| `"accountName"` |  |
| `"availableBalance"` |  |
| `"balance"` |  |
| `"balanceMinorUnits"` |  |
| `"currency"` |  |
| `"id"` |  |
| `"identifier"` |  |
| `"isArchived"` |  |
| `"isConnectedAccount"` |  |
| `"merchantID"` |  |
| `"submittedPayoutsBalance"` |  |

Operations: List.

API path: `/api/v1/accounts/minimal`

#### PaymentInitiation

| Field | Description |
| --- | --- |
| `"paymentInitiationID"` |  |
| `"paymentRequestCallbackUrl"` |  |
| `"paymentRequestID"` |  |
| `"redirectUrl"` |  |
| `"responseType"` |  |
| `"specificErrorMessage"` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/{id}/pisp`

#### PaymentRequest

| Field | Description |
| --- | --- |
| `"addresses"` |  |
| `"amount"` |  |
| `"amountPending"` |  |
| `"amountReceived"` |  |
| `"amountRefunded"` |  |
| `"autoSendReceipt"` |  |
| `"baseOriginUrl"` |  |
| `"callbackUrl"` |  |
| `"cardAuthorizeOnly"` |  |
| `"cardCreateToken"` |  |
| `"cardCreateTokenMode"` |  |
| `"cardIgnoreCVN"` |  |
| `"cardProcessorMerchantID"` |  |
| `"cardStripePaymentIntentID"` |  |
| `"cardStripePaymentIntentSecret"` |  |
| `"createdByUser"` |  |
| `"currency"` |  |
| `"customFields"` |  |
| `"customerEmailAddress"` |  |
| `"customerID"` |  |
| `"customerName"` |  |
| `"description"` |  |
| `"destinationAccount"` |  |
| `"directDebitPayment"` |  |
| `"doSimulateSettlementFailure"` |  |
| `"dueDate"` |  |
| `"errorDescription"` |  |
| `"events"` |  |
| `"failedPaymentRequests"` |  |
| `"failureCallbackUrl"` |  |
| `"fieldDisplaySettings"` |  |
| `"formattedAmount"` |  |
| `"hostedPayCheckoutUrl"` |  |
| `"id"` |  |
| `"ignoreAddressVerification"` |  |
| `"inserted"` |  |
| `"insertedSortable"` |  |
| `"institution"` |  |
| `"isArchived"` |  |
| `"jwk"` |  |
| `"lastUpdated"` |  |
| `"lightningInvoice"` |  |
| `"lightningInvoiceExpiresAt"` |  |
| `"merchantDirectDebitMandateID"` |  |
| `"merchantID"` |  |
| `"merchantTokenDescription"` |  |
| `"notificationEmailAddresses"` |  |
| `"notificationRoleIDs"` |  |
| `"orderID"` |  |
| `"partialPaymentMethod"` |  |
| `"partialPaymentSteps"` |  |
| `"paymentAttempts"` |  |
| `"paymentInitiationID"` |  |
| `"paymentMethods"` |  |
| `"paymentProcessor"` |  |
| `"paymentRequests"` |  |
| `"payrunID"` |  |
| `"pispAccountID"` |  |
| `"priorityBankID"` |  |
| `"result"` |  |
| `"sandboxSettleDelayInSeconds"` |  |
| `"shippingAddress"` |  |
| `"status"` |  |
| `"successWebHookUrl"` |  |
| `"tags"` |  |
| `"title"` |  |
| `"tokenisedCards"` |  |
| `"transactions"` |  |
| `"useHostedPaymentPage"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/paymentrequests/{id}/directdebit`

#### PaymentRequestEvent

| Field | Description |
| --- | --- |
| `"amount"` |  |
| `"applePayTransactionID"` |  |
| `"cardAuthorizationResponseID"` |  |
| `"cardExpiryMonth"` |  |
| `"cardExpiryYear"` |  |
| `"cardIssuer"` |  |
| `"cardIssuerCountry"` |  |
| `"cardLastFourDigits"` |  |
| `"cardRequestID"` |  |
| `"cardScheme"` |  |
| `"cardTokenCustomerID"` |  |
| `"cardTransactionID"` |  |
| `"currency"` |  |
| `"directDebitPaymentID"` |  |
| `"directDebitPaymentReference"` |  |
| `"drirectDebitMandateID"` |  |
| `"errorMessage"` |  |
| `"errorReason"` |  |
| `"eventType"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"lightningInvoice"` |  |
| `"lightningRHash"` |  |
| `"originUrl"` |  |
| `"paymentMethodType"` |  |
| `"paymentProcessorName"` |  |
| `"paymentRequestID"` |  |
| `"pispBankStatus"` |  |
| `"pispPaymentInitiationID"` |  |
| `"pispPaymentInstitutionName"` |  |
| `"pispPaymentServiceProviderID"` |  |
| `"pispRedirectUrl"` |  |
| `"reconciledTransactionID"` |  |
| `"refundPayoutID"` |  |
| `"status"` |  |
| `"walletName"` |  |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/events`

#### PaymentRequestMetric

| Field | Description |
| --- | --- |

Operations: Load.

API path: `/api/v1/paymentrequests/metrics`

#### PaymentRequestMinimal

| Field | Description |
| --- | --- |
| `"amount"` |  |
| `"amountPending"` |  |
| `"amountReceived"` |  |
| `"amountRefunded"` |  |
| `"callbackUrl"` |  |
| `"cardStripePaymentIntentSecret"` |  |
| `"countryCode"` |  |
| `"currency"` |  |
| `"customFieldsToDisplay"` |  |
| `"description"` |  |
| `"dueDate"` |  |
| `"fieldDisplaySettings"` |  |
| `"googlePayMerchantID"` |  |
| `"id"` |  |
| `"jwk"` |  |
| `"merchantID"` |  |
| `"merchantLogoUrlPng"` |  |
| `"merchantLogoUrlSvg"` |  |
| `"merchantName"` |  |
| `"merchantShortName"` |  |
| `"partialPaymentMethod"` |  |
| `"paymentAttempts"` |  |
| `"paymentMethodsList"` |  |
| `"paymentProcessor"` |  |
| `"paymentProcessorKey"` |  |
| `"pispError"` |  |
| `"priorityBankID"` |  |
| `"status"` |  |
| `"stripeAccountID"` |  |
| `"title"` |  |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/minimal`

#### PaymentRequestResult

| Field | Description |
| --- | --- |
| `"amount"` |  |
| `"amountPending"` |  |
| `"amountReceived"` |  |
| `"amountRefunded"` |  |
| `"currency"` |  |
| `"customerID"` |  |
| `"paymentRequestID"` |  |
| `"payments"` |  |
| `"pispAuthorizations"` |  |
| `"requestedAmount"` |  |
| `"result"` |  |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/result`

#### Payout

| Field | Description |
| --- | --- |
| `"accountID"` |  |
| `"allowIncomplete"` |  |
| `"amount"` |  |
| `"amountMinorUnits"` |  |
| `"approvePayoutUrl"` |  |
| `"approverID"` |  |
| `"authenticationMethods"` |  |
| `"authorisations"` |  |
| `"authorisersCompletedCount"` |  |
| `"authorisersRequiredCount"` |  |
| `"batchPayoutID"` |  |
| `"beneficiary"` |  |
| `"beneficiaryID"` |  |
| `"canAuthorise"` |  |
| `"canProcess"` |  |
| `"canUpdate"` |  |
| `"chargeBearer"` |  |
| `"createdBy"` |  |
| `"createdByEmailAddress"` |  |
| `"currency"` |  |
| `"currentUserID"` |  |
| `"description"` |  |
| `"destination"` |  |
| `"documents"` |  |
| `"events"` |  |
| `"failedPayouts"` |  |
| `"formattedAmount"` |  |
| `"formattedFxDestinationAmount"` |  |
| `"formattedSchedule"` |  |
| `"formattedScheduleDayOnly"` |  |
| `"formattedSourceAccountAvailableBalance"` |  |
| `"fxDestinationAmount"` |  |
| `"fxDestinationAmountMinorUnits"` |  |
| `"fxDestinationCurrency"` |  |
| `"fxQuoteExpiresAt"` |  |
| `"fxQuoteID"` |  |
| `"fxRate"` |  |
| `"fxUseDestinationAmount"` |  |
| `"hasCurrentUserAuthorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"invoiceID"` |  |
| `"isArchived"` |  |
| `"isFailed"` |  |
| `"isSettled"` |  |
| `"isSubmitted"` |  |
| `"lastUpdated"` |  |
| `"merchantID"` |  |
| `"merchantTokenDescription"` |  |
| `"nonce"` |  |
| `"paymentProcessor"` |  |
| `"paymentRail"` |  |
| `"payouts"` |  |
| `"payrunID"` |  |
| `"payrunName"` |  |
| `"reason"` |  |
| `"rule"` |  |
| `"scheduleDate"` |  |
| `"scheduled"` |  |
| `"sourceAccountAvailableBalance"` |  |
| `"sourceAccountAvailableBalanceMinorUnits"` |  |
| `"sourceAccountBic"` |  |
| `"sourceAccountCurrency"` |  |
| `"sourceAccountIban"` |  |
| `"sourceAccountIdentifier"` |  |
| `"sourceAccountName"` |  |
| `"sourceAccountNumber"` |  |
| `"sourceAccountSortcode"` |  |
| `"status"` |  |
| `"tagIds"` |  |
| `"tags"` |  |
| `"theirReference"` |  |
| `"topupPayrunID"` |  |
| `"transactedAmount"` |  |
| `"transactedFxAmount"` |  |
| `"transactedFxRate"` |  |
| `"type"` |  |
| `"userID"` |  |
| `"yourReference"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/payouts/batch/submit/{id}`

#### PayoutKeysetPage

| Field | Description |
| --- | --- |
| `"accountID"` |  |
| `"amount"` |  |
| `"amountMinorUnits"` |  |
| `"approvePayoutUrl"` |  |
| `"approverID"` |  |
| `"authenticationMethods"` |  |
| `"authorisations"` |  |
| `"authorisersCompletedCount"` |  |
| `"authorisersRequiredCount"` |  |
| `"batchPayoutID"` |  |
| `"beneficiary"` |  |
| `"canAuthorise"` |  |
| `"canProcess"` |  |
| `"canUpdate"` |  |
| `"chargeBearer"` |  |
| `"createdBy"` |  |
| `"createdByEmailAddress"` |  |
| `"currency"` |  |
| `"currentUserID"` |  |
| `"description"` |  |
| `"destination"` |  |
| `"documents"` |  |
| `"events"` |  |
| `"formattedAmount"` |  |
| `"formattedFxDestinationAmount"` |  |
| `"formattedSchedule"` |  |
| `"formattedScheduleDayOnly"` |  |
| `"formattedSourceAccountAvailableBalance"` |  |
| `"fxDestinationAmount"` |  |
| `"fxDestinationAmountMinorUnits"` |  |
| `"fxDestinationCurrency"` |  |
| `"fxQuoteExpiresAt"` |  |
| `"fxQuoteID"` |  |
| `"fxRate"` |  |
| `"fxUseDestinationAmount"` |  |
| `"hasCurrentUserAuthorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"invoiceID"` |  |
| `"isArchived"` |  |
| `"isFailed"` |  |
| `"isSettled"` |  |
| `"isSubmitted"` |  |
| `"lastUpdated"` |  |
| `"merchantID"` |  |
| `"merchantTokenDescription"` |  |
| `"nonce"` |  |
| `"paymentProcessor"` |  |
| `"paymentRail"` |  |
| `"payrunID"` |  |
| `"payrunName"` |  |
| `"rule"` |  |
| `"scheduleDate"` |  |
| `"scheduled"` |  |
| `"sourceAccountAvailableBalance"` |  |
| `"sourceAccountAvailableBalanceMinorUnits"` |  |
| `"sourceAccountBic"` |  |
| `"sourceAccountCurrency"` |  |
| `"sourceAccountIban"` |  |
| `"sourceAccountIdentifier"` |  |
| `"sourceAccountName"` |  |
| `"sourceAccountNumber"` |  |
| `"sourceAccountSortcode"` |  |
| `"status"` |  |
| `"tags"` |  |
| `"theirReference"` |  |
| `"topupPayrunID"` |  |
| `"transactedAmount"` |  |
| `"transactedFxAmount"` |  |
| `"transactedFxRate"` |  |
| `"type"` |  |
| `"userID"` |  |
| `"yourReference"` |  |

Operations: List.

API path: `/api/v1/accounts/{accountID}/payouts/failed`

#### PayoutMetric

| Field | Description |
| --- | --- |

Operations: Load.

API path: `/api/v1/payouts/metrics`

#### Payrun

| Field | Description |
| --- | --- |
| `"authorisationDate"` |  |
| `"authorisations"` |  |
| `"authorisersCompletedCount"` |  |
| `"authorisersRequiredCount"` |  |
| `"batchPayoutID"` |  |
| `"canAuthorise"` |  |
| `"canDelete"` |  |
| `"canEdit"` |  |
| `"events"` |  |
| `"hasCurrentUserAuthorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"invoices"` |  |
| `"invoicesMinimal"` |  |
| `"isArchived"` |  |
| `"lastUpdated"` |  |
| `"lastUpdatedBy"` |  |
| `"merchantID"` |  |
| `"name"` |  |
| `"nonce"` |  |
| `"notes"` |  |
| `"payments"` |  |
| `"payouts"` |  |
| `"payoutsCount"` |  |
| `"reason"` |  |
| `"scheduleDate"` |  |
| `"scheduledDate"` |  |
| `"sourceAccounts"` |  |
| `"status"` |  |
| `"totalEur"` |  |
| `"totalGbp"` |  |
| `"totalUsd"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/payruns/{id}/request-authorisation`

#### Report

| Field | Description |
| --- | --- |

Operations: Update.

API path: `/api/v1/reports/{id}/initiate`

#### ReportResult

| Field | Description |
| --- | --- |
| `"contentType"` |  |
| `"contents"` |  |
| `"lastCompletedAt"` |  |
| `"merchantID"` |  |
| `"reportName"` |  |
| `"reportType"` |  |
| `"statementNumber"` |  |

Operations: Load.

API path: `/api/v1/reports/{id}/result/{statementNumber}`

#### Role

| Field | Description |
| --- | --- |
| `"failedRoles"` |  |
| `"roles"` |  |

Operations: Create.

API path: `/api/v1/merchants/{merchantID}/roles/batchcreate`

#### Rule

| Field | Description |
| --- | --- |
| `"account"` |  |
| `"accountID"` |  |
| `"approveUrl"` |  |
| `"approverID"` |  |
| `"authenticationMethods"` |  |
| `"authorisations"` |  |
| `"authorisersCompletedCount"` |  |
| `"authorisersRequiredCount"` |  |
| `"canAuthorise"` |  |
| `"createdBy"` |  |
| `"description"` |  |
| `"endAt"` |  |
| `"hasCurrentUserAuthorised"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"isDisabled"` |  |
| `"lastExecutedAt"` |  |
| `"lastRunAtTransactionDate"` |  |
| `"lastUpdated"` |  |
| `"merchantID"` |  |
| `"name"` |  |
| `"nonce"` |  |
| `"onApprovedWebHookUrl"` |  |
| `"onExecutionErrorWebHookUrl"` |  |
| `"onExecutionSuccessWebHookUrl"` |  |
| `"startAt"` |  |
| `"status"` |  |
| `"sweepAction"` |  |
| `"timeZoneId"` |  |
| `"triggerCronExpression"` |  |
| `"triggerOnPayIn"` |  |
| `"userID"` |  |
| `"webHookSecret"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/rules`

#### RuleEvent

| Field | Description |
| --- | --- |
| `"errorMessage"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"isAuthoriseToEnable"` |  |
| `"message"` |  |
| `"rawResponse"` |  |
| `"ruleEventType"` |  |
| `"ruleID"` |  |
| `"user"` |  |

Operations: List.

API path: `/api/v1/rules/{id}/events`

#### Tag

| Field | Description |
| --- | --- |
| `"colourHex"` |  |
| `"description"` |  |
| `"id"` |  |
| `"merchantID"` |  |
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
| `"accountID"` |  |
| `"accountName"` |  |
| `"accountSequenceNumber"` |  |
| `"addressDetails"` |  |
| `"amount"` |  |
| `"amountMinorUnits"` |  |
| `"balance"` |  |
| `"balanceMinorUnits"` |  |
| `"bookingDateTime"` |  |
| `"chargeDetails"` |  |
| `"content"` |  |
| `"counterparty"` |  |
| `"counterpartySummary"` |  |
| `"currency"` |  |
| `"currencyExchange"` |  |
| `"date"` |  |
| `"description"` |  |
| `"enrichment"` |  |
| `"fxAmount"` |  |
| `"fxCurrency"` |  |
| `"fxRate"` |  |
| `"grossAmount"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"isoBankTransactionCode"` |  |
| `"merchant"` |  |
| `"merchantID"` |  |
| `"pageNumber"` |  |
| `"pageSize"` |  |
| `"payeeDetails"` |  |
| `"payerDetails"` |  |
| `"paymentRequestCustomFields"` |  |
| `"paymentRequestID"` |  |
| `"payoutID"` |  |
| `"proprietaryBankTransactionCode"` |  |
| `"rawReference"` |  |
| `"reference"` |  |
| `"ruleID"` |  |
| `"statementReferences"` |  |
| `"status"` |  |
| `"supplementaryData"` |  |
| `"tags"` |  |
| `"theirReference"` |  |
| `"totalPages"` |  |
| `"totalSize"` |  |
| `"transactionAmount"` |  |
| `"transactionDate"` |  |
| `"transactionInformation"` |  |
| `"transactionMutability"` |  |
| `"type"` |  |
| `"valueDateTime"` |  |
| `"virtualIBAN"` |  |
| `"yourReference"` |  |

Operations: Create, List, Load, Remove.

API path: `/api/v1/transactions/{id}/tags`

#### User

| Field | Description |
| --- | --- |
| `"clientSessionTimeouts"` |  |
| `"emailAddress"` |  |
| `"firstName"` |  |
| `"id"` |  |
| `"lastName"` |  |
| `"passkeyAdded"` |  |
| `"permissions"` |  |
| `"profile"` |  |
| `"rolesWithScope"` |  |
| `"twoFactorEnabled"` |  |
| `"userInviteID"` |  |

Operations: List, Update.

API path: `/api/v1/user/{merchantID}/userspaged`

#### UserInvite

| Field | Description |
| --- | --- |
| `"authorisationStatus"` |  |
| `"failedUserInvites"` |  |
| `"id"` |  |
| `"initialRoleID"` |  |
| `"inviteeEmailAddress"` |  |
| `"inviteeFirstName"` |  |
| `"inviteeLastName"` |  |
| `"inviterEmailAddress"` |  |
| `"inviterFirstName"` |  |
| `"inviterLastName"` |  |
| `"isAuthorised"` |  |
| `"isInviteeRegistered"` |  |
| `"lastInvited"` |  |
| `"merchantID"` |  |
| `"merchantName"` |  |
| `"message"` |  |
| `"registrationUrl"` |  |
| `"sendInviteEmail"` |  |
| `"status"` |  |
| `"user"` |  |
| `"userID"` |  |
| `"userInvites"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/userinvites/authorise/{id}`

#### Virtual

| Field | Description |
| --- | --- |
| `"accountName"` |  |
| `"accountSupplierName"` |  |
| `"availableBalance"` |  |
| `"availableBalanceMinorUnits"` |  |
| `"balance"` |  |
| `"balanceMinorUnits"` |  |
| `"bankName"` |  |
| `"consentID"` |  |
| `"createdBy"` |  |
| `"createdByDisplayName"` |  |
| `"currency"` |  |
| `"defaultPaymentRail"` |  |
| `"displayName"` |  |
| `"expiryDate"` |  |
| `"externalAccountIcon"` |  |
| `"id"` |  |
| `"identifier"` |  |
| `"inserted"` |  |
| `"isArchived"` |  |
| `"isConnectedAccount"` |  |
| `"isDefault"` |  |
| `"isTrustAccount"` |  |
| `"isVirtual"` |  |
| `"lastTransaction"` |  |
| `"lastUpdated"` |  |
| `"merchantID"` |  |
| `"merchantName"` |  |
| `"name"` |  |
| `"physicalAccountID"` |  |
| `"rules"` |  |
| `"submittedPayoutsBalance"` |  |
| `"submittedPayoutsBalanceMinorUnits"` |  |
| `"summary"` |  |
| `"supplierSepaInstantStatus"` |  |
| `"xeroBankFeedConnectionStatus"` |  |
| `"xeroBankFeedLastSyncedAt"` |  |
| `"xeroBankFeedSyncLastFailedAt"` |  |
| `"xeroBankFeedSyncLastFailureReason"` |  |
| `"xeroBankFeedSyncStatus"` |  |
| `"xeroUnsynchronisedTransactionsCount"` |  |

Operations: Create, Update.

API path: `/api/v1/accounts/{accountID}/virtual`

#### Webhook

| Field | Description |
| --- | --- |
| `"destinationUrl"` |  |
| `"emailAddress"` |  |
| `"failedNotificationEmailAddress"` |  |
| `"id"` |  |
| `"isActive"` |  |
| `"merchantID"` |  |
| `"notificationMethod"` |  |
| `"resourceTypes"` |  |
| `"retry"` |  |
| `"secret"` |  |
| `"version"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/webhooks`



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
| `accountBalances` | `[]any` |  |
| `accountID` | `string` |  |
| `accountIdentifications` | `[]any` |  |
| `accountName` | `string` |  |
| `accountNames` | `[]any` |  |
| `accountSupplierName` | `string` |  |
| `accountType` | `string` |  |
| `availableBalance` | `float64` |  |
| `availableBalanceMinorUnits` | `int` |  |
| `balance` | `float64` |  |
| `balanceMinorUnits` | `int` |  |
| `bankName` | `string` |  |
| `consentID` | `string` |  |
| `consolidatedAccountInformation` | `map[string]any` |  |
| `createdBy` | `map[string]any` |  |
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
| `identifier` | `map[string]any` |  |
| `inserted` | `string` |  |
| `isArchived` | `bool` |  |
| `isConnectedAccount` | `bool` |  |
| `isDefault` | `bool` |  |
| `isTrustAccount` | `bool` |  |
| `isVirtual` | `bool` |  |
| `lastTransaction` | `map[string]any` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantName` | `string` |  |
| `nickname` | `string` |  |
| `physicalAccountID` | `string` |  |
| `roleIDs` | `[]any` |  |
| `rules` | `[]any` |  |
| `submittedPayoutsBalance` | `float64` |  |
| `submittedPayoutsBalanceMinorUnits` | `int` |  |
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
| `xeroUnsynchronisedTransactionsCount` | `int` |  |

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
    "createdBy": map[string]any{},
    "identifier": map[string]any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Batch

Create an instance: `batch := client.Batch(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approveUrl` | `string` |  |
| `id` | `string` |  |
| `payouts` | `[]any` |  |

#### Example: Load

```go
batch, err := client.Batch(nil).Load(map[string]any{"id": "batch_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(batch) // the loaded record
```

#### Example: Create

```go
result, err := client.Batch(nil).Create(map[string]any{
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
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `approvalCallbackUrl` | `string` |  |
| `authenticationMethods` | `[]any` |  |
| `authorisations` | `[]any` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `beneficiaries` | `[]any` |  |
| `beneficiaryEvents` | `[]any` |  |
| `canAuthorise` | `bool` |  |
| `canUpdate` | `bool` |  |
| `createdBy` | `map[string]any` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` |  |
| `destination` | `map[string]any` |  |
| `failedBeneficiaries` | `map[string]any` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isEnabled` | `bool` |  |
| `lastAuthorised` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `sourceAccountIDs` | `[]any` |  |
| `sourceAccounts` | `[]any` |  |
| `theirReference` | `string` |  |

#### Example: Load

```go
beneficiary, err := client.Beneficiary(nil).Load(map[string]any{"id": "beneficiary_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(beneficiary) // the loaded record
```

#### Example: List

```go
beneficiarys, err := client.Beneficiary(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(beneficiarys) // the array of records
```

#### Example: Create

```go
result, err := client.Beneficiary(nil).Create(map[string]any{
    "createdBy": map[string]any{},
    "currency": "example_currency",
    "name": "example_name",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### BeneficiaryGroup

Create an instance: `beneficiaryGroup := client.BeneficiaryGroup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `groupMembers` | `[]any` |  |
| `groupName` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |

#### Example: List

```go
beneficiaryGroups, err := client.BeneficiaryGroup(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(beneficiaryGroups) // the array of records
```


### Card

Create an instance: `card := client.Card(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorizedAmount` | `string` |  |
| `currencyCode` | `string` |  |
| `isPayerAuthenticationRequired` | `bool` |  |
| `isSoftDecline` | `bool` |  |
| `payerAuthenticationAccessToken` | `string` |  |
| `payerAuthenticationMerchantData` | `string` |  |
| `payerAuthenticationUrl` | `string` |  |
| `payerAuthenticationWindowHeight` | `int` |  |
| `payerAuthenticationWindowWidth` | `int` |  |
| `paymentRequestCallbackUrl` | `string` |  |
| `paymentRequestID` | `string` |  |
| `requestID` | `string` |  |
| `responseCode` | `string` |  |
| `responseType` | `string` |  |
| `status` | `string` |  |
| `threeDSRedirectUrl` | `string` |  |
| `transactionID` | `string` |  |

#### Example: Create

```go
result, err := client.Card(nil).Create(map[string]any{
    "paymentrequest_id": "example_paymentrequest_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### CardCustomerToken

Create an instance: `cardCustomerToken := client.CardCustomerToken(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Remove(match, ctrl)` | Remove the matching entity. |

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

```go
cardCustomerToken, err := client.CardCustomerToken(nil).Load(map[string]any{"customer_email_address": "customer_email_address"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(cardCustomerToken) // the loaded record
```

#### Example: List

```go
cardCustomerTokens, err := client.CardCustomerToken(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(cardCustomerTokens) // the array of records
```


### CardPayment

Create an instance: `cardPayment := client.CardPayment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorizedAmount` | `string` |  |
| `currencyCode` | `string` |  |
| `isPayerAuthenticationRequired` | `bool` |  |
| `isSoftDecline` | `bool` |  |
| `payerAuthenticationAccessToken` | `string` |  |
| `payerAuthenticationMerchantData` | `string` |  |
| `payerAuthenticationUrl` | `string` |  |
| `payerAuthenticationWindowHeight` | `int` |  |
| `payerAuthenticationWindowWidth` | `int` |  |
| `paymentRequestCallbackUrl` | `string` |  |
| `paymentRequestID` | `string` |  |
| `requestID` | `string` |  |
| `responseCode` | `string` |  |
| `responseType` | `string` |  |
| `status` | `string` |  |
| `threeDSRedirectUrl` | `string` |  |
| `transactionID` | `string` |  |

#### Example: Create

```go
result, err := client.CardPayment(nil).Create(map[string]any{
    "paymentrequest_id": "example_paymentrequest_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### CardPublicKey

Create an instance: `cardPublicKey := client.CardPublicKey(nil)`

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
cardPublicKey, err := client.CardPublicKey(nil).Load(map[string]any{"paymentrequest_id": "paymentrequest_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(cardPublicKey) // the loaded record
```


### Consent

Create an instance: `consent := client.Consent(nil)`

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
| `authorisationUrl` | `string` |  |
| `callbackUrl` | `string` |  |
| `consentID` | `string` |  |
| `emailAddress` | `string` |  |
| `expiryDate` | `string` |  |
| `failureCallbackUrl` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `institutionID` | `string` |  |
| `isConnectedAccounts` | `bool` |  |
| `isEnabled` | `bool` |  |
| `merchantID` | `string` |  |
| `provider` | `string` |  |
| `successWebHookUrl` | `string` |  |

#### Example: Load

```go
consent, err := client.Consent(nil).Load(map[string]any{"id": "consent_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(consent) // the loaded record
```

#### Example: List

```go
consents, err := client.Consent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(consents) // the array of records
```

#### Example: Create

```go
result, err := client.Consent(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Currency

Create an instance: `currency := client.Currency(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `code` | `string` |  |
| `decimals` | `int` |  |
| `isFiat` | `bool` |  |
| `iso4217AlphaCode` | `string` |  |
| `iso4217NumericCode` | `string` |  |
| `symbol` | `string` |  |

#### Example: List

```go
currencys, err := client.Currency(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(currencys) // the array of records
```


### DirectDebitBatchSubmit

Create an instance: `directDebitBatchSubmit := client.DirectDebitBatchSubmit(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failedSubmissions` | `map[string]any` |  |
| `successfulSubmissions` | `[]any` |  |

#### Example: Create

```go
result, err := client.DirectDebitBatchSubmit(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### FxRate

Create an instance: `fxRate := client.FxRate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `destinationCurrency` | `string` |  |
| `exchangeRate` | `float64` |  |
| `expiryTime` | `string` |  |
| `quoteID` | `string` |  |
| `sourceCurrency` | `string` |  |

#### Example: Load

```go
fxRate, err := client.FxRate(nil).Load(map[string]any{"destination": "destination", "source": "source", "valid_for_minute": 1}, nil)
if err != nil {
    panic(err)
}
fmt.Println(fxRate) // the loaded record
```

#### Example: List

```go
fxRates, err := client.FxRate(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(fxRates) // the array of records
```


### IPayment

Create an instance: `iPayment := client.IPayment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `paymentRequestID` | `string` |  |
| `responseType` | `string` |  |

#### Example: Create

```go
result, err := client.IPayment(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Mandate

Create an instance: `mandate := client.Mandate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

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
| `isRecurring` | `bool` |  |
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

```go
mandate, err := client.Mandate(nil).Load(map[string]any{"id": "mandate_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(mandate) // the loaded record
```

#### Example: Create

```go
result, err := client.Mandate(nil).Create(map[string]any{
    "addressLine1": "example_addressLine1",
    "city": "example_city",
    "countryCode": "example_countryCode",
    "emailAddress": "example_emailAddress",
    "firstName": "example_firstName",
    "lastName": "example_lastName",
    "postalCode": "example_postalCode",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Merchant

Create an instance: `merchant := client.Merchant(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountCurrencies` | `[]any` |  |
| `canHaveTrustAccounts` | `bool` |  |
| `cardPaymentProcessor` | `string` |  |
| `companyID` | `string` |  |
| `displayQrOnHostedPay` | `bool` |  |
| `hostedPayVersion` | `int` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isBlocked` | `bool` |  |
| `isExited` | `bool` |  |
| `isSuspended` | `bool` |  |
| `jurisdiction` | `string` |  |
| `logoUrlPng` | `string` |  |
| `logoUrlSvg` | `string` |  |
| `merchantCategoryCode` | `string` |  |
| `name` | `string` |  |
| `notes` | `string` |  |
| `parentMerchant` | `map[string]any` |  |
| `paymentAccountLimit` | `int` |  |
| `paymentAccounts` | `[]any` |  |
| `reason` | `string` |  |
| `shortName` | `string` |  |
| `supportedPaymentMethodsList` | `[]any` |  |
| `suspensionReason` | `string` |  |
| `tags` | `[]any` |  |
| `timeZoneId` | `string` |  |
| `tradingName` | `string` |  |
| `webHookLimit` | `int` |  |
| `yourRoleName` | `string` |  |

#### Example: Load

```go
merchant, err := client.Merchant(nil).Load(map[string]any{"id": "merchant_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(merchant) // the loaded record
```

#### Example: List

```go
merchants, err := client.Merchant(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(merchants) // the array of records
```


### MerchantAuthorisationSetting

Create an instance: `merchantAuthorisationSetting := client.MerchantAuthorisationSetting(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amountLower` | `float64` |  |
| `amountUpper` | `float64` |  |
| `authorisationType` | `string` |  |
| `beneficiariesOnly` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lastEditorCantAuthorise` | `bool` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `numberOfAuthorisers` | `int` |  |
| `roleSettings` | `[]any` |  |

#### Example: List

```go
merchantAuthorisationSettings, err := client.MerchantAuthorisationSetting(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(merchantAuthorisationSettings) // the array of records
```


### MerchantDirectDebitMandatePage

Create an instance: `merchantDirectDebitMandatePage := client.MerchantDirectDebitMandatePage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

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
| `isRecurring` | `bool` |  |
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

```go
merchantDirectDebitMandatePages, err := client.MerchantDirectDebitMandatePage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(merchantDirectDebitMandatePages) // the array of records
```


### MerchantPayByBankSetting

Create an instance: `merchantPayByBankSetting := client.MerchantPayByBankSetting(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bankCountryCodes` | `[]any` |  |
| `bankID` | `string` |  |
| `bankName` | `string` |  |
| `businessInstitutionID` | `string` |  |
| `currency` | `string` |  |
| `logo` | `string` |  |
| `message` | `string` |  |
| `messageImageUrl` | `string` |  |
| `order` | `int` |  |
| `personalInstitutionID` | `string` |  |
| `processor` | `string` |  |
| `warningHeading` | `string` |  |
| `warningMessage` | `string` |  |

#### Example: List

```go
merchantPayByBankSettings, err := client.MerchantPayByBankSetting(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(merchantPayByBankSettings) // the array of records
```


### MerchantPaymentRequestTemplate

Create an instance: `merchantPaymentRequestTemplate := client.MerchantPaymentRequestTemplate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `bankPaymentOptions` | `map[string]any` |  |
| `cardPaymentAddressOptions` | `map[string]any` |  |
| `cardPaymentCaptureOptions` | `map[string]any` |  |
| `customFields` | `[]any` |  |
| `defaultFields` | `[]any` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `name` | `string` |  |
| `notificationOptions` | `map[string]any` |  |
| `paymentMethods` | `map[string]any` |  |
| `paymentTerms` | `map[string]any` |  |
| `priorityBankOptions` | `map[string]any` |  |
| `template` | `map[string]any` |  |

#### Example: Load

```go
merchantPaymentRequestTemplate, err := client.MerchantPaymentRequestTemplate(nil).Load(map[string]any{"id": "merchant_payment_request_template_id", "paymentrequest_id": "paymentrequest_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(merchantPaymentRequestTemplate) // the loaded record
```

#### Example: List

```go
merchantPaymentRequestTemplates, err := client.MerchantPaymentRequestTemplate(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(merchantPaymentRequestTemplates) // the array of records
```


### MerchantToken

Create an instance: `merchantToken := client.MerchantToken(nil)`

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
| `authenticationMethods` | `[]any` |  |
| `authorisations` | `[]any` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `canAuthorise` | `bool` |  |
| `description` | `string` |  |
| `expiresAt` | `string` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `hmacAlgorithm` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `ipAddressWhitelist` | `string` |  |
| `isArchived` | `bool` |  |
| `isEnabled` | `bool` |  |
| `lastAuthorised` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `nonce` | `string` |  |
| `permissionTypes` | `[]any` |  |
| `requestSignatureVersion` | `int` |  |
| `sharedSecretAlgorithm` | `string` |  |
| `sharedSecretBase64` | `string` |  |
| `token` | `string` |  |

#### Example: Load

```go
merchantToken, err := client.MerchantToken(nil).Load(map[string]any{"id": "merchant_token_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(merchantToken) // the loaded record
```

#### Example: List

```go
merchantTokens, err := client.MerchantToken(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(merchantTokens) // the array of records
```

#### Example: Create

```go
result, err := client.MerchantToken(nil).Create(map[string]any{
    "nonce": "example_nonce",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
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


### NoFrixionVersion

Create an instance: `noFrixionVersion := client.NoFrixionVersion(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `buildVersion` | `int` |  |
| `majorVersion` | `int` |  |
| `minorVersion` | `int` |  |
| `releaseName` | `string` |  |

#### Example: Load

```go
noFrixionVersion, err := client.NoFrixionVersion(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(noFrixionVersion) // the loaded record
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
    "account_id": "example_account_id",
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
| `accountName` | `string` |  |
| `accountNumber` | `string` |  |
| `iban` | `string` |  |
| `payeeVerifiedAccountName` | `string` |  |
| `result` | `string` |  |
| `secondaryIdentification` | `string` |  |
| `sortCode` | `string` |  |

#### Example: Create

```go
result, err := client.Payeeverification(nil).Create(map[string]any{
    "accountName": "example_accountName",
    "iban": "example_iban",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Payment

Create an instance: `payment := client.Payment(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addresses` | `[]any` |  |
| `amount` | `float64` |  |
| `amountPending` | `float64` |  |
| `amountReceived` | `float64` |  |
| `amountRefunded` | `float64` |  |
| `autoSendReceipt` | `bool` |  |
| `baseOriginUrl` | `string` |  |
| `callbackUrl` | `string` |  |
| `cardAuthorizeOnly` | `bool` |  |
| `cardCreateToken` | `bool` |  |
| `cardCreateTokenMode` | `string` |  |
| `cardIgnoreCVN` | `bool` |  |
| `cardNoPayerAuthentication` | `bool` |  |
| `cardProcessorMerchantID` | `string` |  |
| `cardStripePaymentIntentID` | `string` |  |
| `cardStripePaymentIntentSecret` | `string` |  |
| `cardTransmitRawDetails` | `bool` |  |
| `createdByUser` | `map[string]any` |  |
| `currency` | `string` |  |
| `customFields` | `[]any` |  |
| `customerEmailAddress` | `string` |  |
| `customerID` | `string` |  |
| `customerName` | `string` |  |
| `description` | `string` |  |
| `destinationAccount` | `map[string]any` |  |
| `directDebitPayment` | `map[string]any` |  |
| `dueDate` | `string` |  |
| `events` | `[]any` |  |
| `failureCallbackUrl` | `string` |  |
| `fieldDisplaySettings` | `[]any` |  |
| `formattedAmount` | `string` |  |
| `hostedPayCheckoutUrl` | `string` |  |
| `id` | `string` |  |
| `ignoreAddressVerification` | `bool` |  |
| `inserted` | `string` |  |
| `insertedSortable` | `string` |  |
| `isArchived` | `bool` |  |
| `jwk` | `string` |  |
| `lastUpdated` | `string` |  |
| `lightningInvoice` | `string` |  |
| `lightningInvoiceExpiresAt` | `string` |  |
| `merchantDirectDebitMandateID` | `string` |  |
| `merchantID` | `string` |  |
| `merchantTokenDescription` | `string` |  |
| `notificationEmailAddresses` | `string` |  |
| `notificationRoleIDs` | `[]any` |  |
| `orderID` | `string` |  |
| `partialPaymentMethod` | `string` |  |
| `partialPaymentSteps` | `string` |  |
| `paymentAttempts` | `[]any` |  |
| `paymentMethods` | `[]any` |  |
| `paymentProcessor` | `string` |  |
| `payrunID` | `string` |  |
| `pispAccountID` | `string` |  |
| `priorityBankID` | `string` |  |
| `result` | `map[string]any` |  |
| `sandboxSettleDelayInSeconds` | `int` |  |
| `shippingAddress` | `map[string]any` |  |
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
| `tagIds` | `[]any` |  |
| `tags` | `[]any` |  |
| `title` | `string` |  |
| `tokenisedCards` | `[]any` |  |
| `transactions` | `[]any` |  |
| `useHostedPaymentPage` | `bool` |  |

#### Example: Load

```go
payment, err := client.Payment(nil).Load(map[string]any{"id": "payment_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(payment) // the loaded record
```

#### Example: Create

```go
result, err := client.Payment(nil).Create(map[string]any{
    "createdByUser": map[string]any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PaymentAccount

Create an instance: `paymentAccount := client.PaymentAccount(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountName` | `string` |  |
| `accountSupplierName` | `string` |  |
| `availableBalance` | `float64` |  |
| `availableBalanceMinorUnits` | `int` |  |
| `balance` | `float64` |  |
| `balanceMinorUnits` | `int` |  |
| `bankName` | `string` |  |
| `consentID` | `string` |  |
| `createdBy` | `map[string]any` |  |
| `createdByDisplayName` | `string` |  |
| `currency` | `string` |  |
| `defaultPaymentRail` | `string` |  |
| `displayName` | `string` |  |
| `expiryDate` | `string` |  |
| `externalAccountIcon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `map[string]any` |  |
| `inserted` | `string` |  |
| `isArchived` | `bool` |  |
| `isConnectedAccount` | `bool` |  |
| `isDefault` | `bool` |  |
| `isTrustAccount` | `bool` |  |
| `isVirtual` | `bool` |  |
| `lastTransaction` | `map[string]any` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantName` | `string` |  |
| `physicalAccountID` | `string` |  |
| `rules` | `[]any` |  |
| `submittedPayoutsBalance` | `float64` |  |
| `submittedPayoutsBalanceMinorUnits` | `int` |  |
| `summary` | `string` |  |
| `supplierSepaInstantStatus` | `string` |  |
| `xeroBankFeedConnectionStatus` | `string` |  |
| `xeroBankFeedLastSyncedAt` | `string` |  |
| `xeroBankFeedSyncLastFailedAt` | `string` |  |
| `xeroBankFeedSyncLastFailureReason` | `string` |  |
| `xeroBankFeedSyncStatus` | `string` |  |
| `xeroUnsynchronisedTransactionsCount` | `int` |  |

#### Example: List

```go
paymentAccounts, err := client.PaymentAccount(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(paymentAccounts) // the array of records
```


### PaymentAccountMinimal

Create an instance: `paymentAccountMinimal := client.PaymentAccountMinimal(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountName` | `string` |  |
| `availableBalance` | `float64` |  |
| `balance` | `float64` |  |
| `balanceMinorUnits` | `int` |  |
| `currency` | `string` |  |
| `id` | `string` |  |
| `identifier` | `map[string]any` |  |
| `isArchived` | `bool` |  |
| `isConnectedAccount` | `bool` |  |
| `merchantID` | `string` |  |
| `submittedPayoutsBalance` | `float64` |  |

#### Example: List

```go
paymentAccountMinimals, err := client.PaymentAccountMinimal(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(paymentAccountMinimals) // the array of records
```


### PaymentInitiation

Create an instance: `paymentInitiation := client.PaymentInitiation(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

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

```go
result, err := client.PaymentInitiation(nil).Create(map[string]any{
    "paymentrequest_id": "example_paymentrequest_id",
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
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `addresses` | `[]any` |  |
| `amount` | `float64` |  |
| `amountPending` | `float64` |  |
| `amountReceived` | `float64` |  |
| `amountRefunded` | `float64` |  |
| `autoSendReceipt` | `bool` |  |
| `baseOriginUrl` | `string` |  |
| `callbackUrl` | `string` |  |
| `cardAuthorizeOnly` | `bool` |  |
| `cardCreateToken` | `bool` |  |
| `cardCreateTokenMode` | `string` |  |
| `cardIgnoreCVN` | `bool` |  |
| `cardProcessorMerchantID` | `string` |  |
| `cardStripePaymentIntentID` | `string` |  |
| `cardStripePaymentIntentSecret` | `string` |  |
| `createdByUser` | `map[string]any` |  |
| `currency` | `string` |  |
| `customFields` | `[]any` |  |
| `customerEmailAddress` | `string` |  |
| `customerID` | `string` |  |
| `customerName` | `string` |  |
| `description` | `string` |  |
| `destinationAccount` | `map[string]any` |  |
| `directDebitPayment` | `map[string]any` |  |
| `doSimulateSettlementFailure` | `bool` |  |
| `dueDate` | `string` |  |
| `errorDescription` | `string` |  |
| `events` | `[]any` |  |
| `failedPaymentRequests` | `map[string]any` |  |
| `failureCallbackUrl` | `string` |  |
| `fieldDisplaySettings` | `[]any` |  |
| `formattedAmount` | `string` |  |
| `hostedPayCheckoutUrl` | `string` |  |
| `id` | `string` |  |
| `ignoreAddressVerification` | `bool` |  |
| `inserted` | `string` |  |
| `insertedSortable` | `string` |  |
| `institution` | `string` |  |
| `isArchived` | `bool` |  |
| `jwk` | `string` |  |
| `lastUpdated` | `string` |  |
| `lightningInvoice` | `string` |  |
| `lightningInvoiceExpiresAt` | `string` |  |
| `merchantDirectDebitMandateID` | `string` |  |
| `merchantID` | `string` |  |
| `merchantTokenDescription` | `string` |  |
| `notificationEmailAddresses` | `string` |  |
| `notificationRoleIDs` | `[]any` |  |
| `orderID` | `string` |  |
| `partialPaymentMethod` | `string` |  |
| `partialPaymentSteps` | `string` |  |
| `paymentAttempts` | `[]any` |  |
| `paymentInitiationID` | `string` |  |
| `paymentMethods` | `[]any` |  |
| `paymentProcessor` | `string` |  |
| `paymentRequests` | `[]any` |  |
| `payrunID` | `string` |  |
| `pispAccountID` | `string` |  |
| `priorityBankID` | `string` |  |
| `result` | `map[string]any` |  |
| `sandboxSettleDelayInSeconds` | `int` |  |
| `shippingAddress` | `map[string]any` |  |
| `status` | `string` |  |
| `successWebHookUrl` | `string` |  |
| `tags` | `[]any` |  |
| `title` | `string` |  |
| `tokenisedCards` | `[]any` |  |
| `transactions` | `[]any` |  |
| `useHostedPaymentPage` | `bool` |  |

#### Example: Load

```go
paymentRequest, err := client.PaymentRequest(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(paymentRequest) // the loaded record
```

#### Example: List

```go
paymentRequests, err := client.PaymentRequest(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(paymentRequests) // the array of records
```

#### Example: Create

```go
result, err := client.PaymentRequest(nil).Create(map[string]any{
    "createdByUser": map[string]any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PaymentRequestEvent

Create an instance: `paymentRequestEvent := client.PaymentRequestEvent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float64` |  |
| `applePayTransactionID` | `string` |  |
| `cardAuthorizationResponseID` | `string` |  |
| `cardExpiryMonth` | `int` |  |
| `cardExpiryYear` | `int` |  |
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

```go
paymentRequestEvents, err := client.PaymentRequestEvent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(paymentRequestEvents) // the array of records
```


### PaymentRequestMetric

Create an instance: `paymentRequestMetric := client.PaymentRequestMetric(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Example: Load

```go
paymentRequestMetric, err := client.PaymentRequestMetric(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(paymentRequestMetric) // the loaded record
```


### PaymentRequestMinimal

Create an instance: `paymentRequestMinimal := client.PaymentRequestMinimal(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float64` |  |
| `amountPending` | `float64` |  |
| `amountReceived` | `float64` |  |
| `amountRefunded` | `float64` |  |
| `callbackUrl` | `string` |  |
| `cardStripePaymentIntentSecret` | `string` |  |
| `countryCode` | `string` |  |
| `currency` | `string` |  |
| `customFieldsToDisplay` | `[]any` |  |
| `description` | `string` |  |
| `dueDate` | `string` |  |
| `fieldDisplaySettings` | `[]any` |  |
| `googlePayMerchantID` | `string` |  |
| `id` | `string` |  |
| `jwk` | `string` |  |
| `merchantID` | `string` |  |
| `merchantLogoUrlPng` | `string` |  |
| `merchantLogoUrlSvg` | `string` |  |
| `merchantName` | `string` |  |
| `merchantShortName` | `string` |  |
| `partialPaymentMethod` | `string` |  |
| `paymentAttempts` | `[]any` |  |
| `paymentMethodsList` | `[]any` |  |
| `paymentProcessor` | `string` |  |
| `paymentProcessorKey` | `string` |  |
| `pispError` | `string` |  |
| `priorityBankID` | `string` |  |
| `status` | `string` |  |
| `stripeAccountID` | `string` |  |
| `title` | `string` |  |

#### Example: List

```go
paymentRequestMinimals, err := client.PaymentRequestMinimal(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(paymentRequestMinimals) // the array of records
```


### PaymentRequestResult

Create an instance: `paymentRequestResult := client.PaymentRequestResult(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `float64` |  |
| `amountPending` | `float64` |  |
| `amountReceived` | `float64` |  |
| `amountRefunded` | `float64` |  |
| `currency` | `string` |  |
| `customerID` | `string` |  |
| `paymentRequestID` | `string` |  |
| `payments` | `[]any` |  |
| `pispAuthorizations` | `[]any` |  |
| `requestedAmount` | `float64` |  |
| `result` | `string` |  |

#### Example: List

```go
paymentRequestResults, err := client.PaymentRequestResult(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(paymentRequestResults) // the array of records
```


### Payout

Create an instance: `payout := client.Payout(nil)`

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
| `accountID` | `string` |  |
| `allowIncomplete` | `bool` |  |
| `amount` | `float64` |  |
| `amountMinorUnits` | `int` |  |
| `approvePayoutUrl` | `string` |  |
| `approverID` | `string` |  |
| `authenticationMethods` | `[]any` |  |
| `authorisations` | `[]any` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `batchPayoutID` | `string` |  |
| `beneficiary` | `map[string]any` |  |
| `beneficiaryID` | `string` |  |
| `canAuthorise` | `bool` |  |
| `canProcess` | `bool` |  |
| `canUpdate` | `bool` |  |
| `chargeBearer` | `string` |  |
| `createdBy` | `string` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` |  |
| `currentUserID` | `string` |  |
| `description` | `string` |  |
| `destination` | `map[string]any` |  |
| `documents` | `[]any` |  |
| `events` | `[]any` |  |
| `failedPayouts` | `map[string]any` |  |
| `formattedAmount` | `string` |  |
| `formattedFxDestinationAmount` | `string` |  |
| `formattedSchedule` | `string` |  |
| `formattedScheduleDayOnly` | `string` |  |
| `formattedSourceAccountAvailableBalance` | `string` |  |
| `fxDestinationAmount` | `float64` |  |
| `fxDestinationAmountMinorUnits` | `int` |  |
| `fxDestinationCurrency` | `string` |  |
| `fxQuoteExpiresAt` | `string` |  |
| `fxQuoteID` | `string` |  |
| `fxRate` | `float64` |  |
| `fxUseDestinationAmount` | `bool` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoiceID` | `string` |  |
| `isArchived` | `bool` |  |
| `isFailed` | `bool` |  |
| `isSettled` | `bool` |  |
| `isSubmitted` | `bool` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantTokenDescription` | `string` |  |
| `nonce` | `string` |  |
| `paymentProcessor` | `string` |  |
| `paymentRail` | `string` |  |
| `payouts` | `[]any` |  |
| `payrunID` | `string` |  |
| `payrunName` | `string` |  |
| `reason` | `string` |  |
| `rule` | `map[string]any` |  |
| `scheduleDate` | `string` |  |
| `scheduled` | `bool` |  |
| `sourceAccountAvailableBalance` | `float64` |  |
| `sourceAccountAvailableBalanceMinorUnits` | `int` |  |
| `sourceAccountBic` | `string` |  |
| `sourceAccountCurrency` | `string` |  |
| `sourceAccountIban` | `string` |  |
| `sourceAccountIdentifier` | `map[string]any` |  |
| `sourceAccountName` | `string` |  |
| `sourceAccountNumber` | `string` |  |
| `sourceAccountSortcode` | `string` |  |
| `status` | `string` |  |
| `tagIds` | `[]any` |  |
| `tags` | `[]any` |  |
| `theirReference` | `string` |  |
| `topupPayrunID` | `string` |  |
| `transactedAmount` | `float64` |  |
| `transactedFxAmount` | `float64` |  |
| `transactedFxRate` | `float64` |  |
| `type` | `string` |  |
| `userID` | `string` |  |
| `yourReference` | `string` |  |

#### Example: Load

```go
payout, err := client.Payout(nil).Load(map[string]any{"id": "payout_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(payout) // the loaded record
```

#### Example: List

```go
payouts, err := client.Payout(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(payouts) // the array of records
```

#### Example: Create

```go
result, err := client.Payout(nil).Create(map[string]any{
    "beneficiary": map[string]any{},
    "sourceAccountIdentifier": map[string]any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PayoutKeysetPage

Create an instance: `payoutKeysetPage := client.PayoutKeysetPage(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountID` | `string` |  |
| `amount` | `float64` |  |
| `amountMinorUnits` | `int` |  |
| `approvePayoutUrl` | `string` |  |
| `approverID` | `string` |  |
| `authenticationMethods` | `[]any` |  |
| `authorisations` | `[]any` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `batchPayoutID` | `string` |  |
| `beneficiary` | `map[string]any` |  |
| `canAuthorise` | `bool` |  |
| `canProcess` | `bool` |  |
| `canUpdate` | `bool` |  |
| `chargeBearer` | `string` |  |
| `createdBy` | `string` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` |  |
| `currentUserID` | `string` |  |
| `description` | `string` |  |
| `destination` | `map[string]any` |  |
| `documents` | `[]any` |  |
| `events` | `[]any` |  |
| `formattedAmount` | `string` |  |
| `formattedFxDestinationAmount` | `string` |  |
| `formattedSchedule` | `string` |  |
| `formattedScheduleDayOnly` | `string` |  |
| `formattedSourceAccountAvailableBalance` | `string` |  |
| `fxDestinationAmount` | `float64` |  |
| `fxDestinationAmountMinorUnits` | `int` |  |
| `fxDestinationCurrency` | `string` |  |
| `fxQuoteExpiresAt` | `string` |  |
| `fxQuoteID` | `string` |  |
| `fxRate` | `float64` |  |
| `fxUseDestinationAmount` | `bool` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoiceID` | `string` |  |
| `isArchived` | `bool` |  |
| `isFailed` | `bool` |  |
| `isSettled` | `bool` |  |
| `isSubmitted` | `bool` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantTokenDescription` | `string` |  |
| `nonce` | `string` |  |
| `paymentProcessor` | `string` |  |
| `paymentRail` | `string` |  |
| `payrunID` | `string` |  |
| `payrunName` | `string` |  |
| `rule` | `map[string]any` |  |
| `scheduleDate` | `string` |  |
| `scheduled` | `bool` |  |
| `sourceAccountAvailableBalance` | `float64` |  |
| `sourceAccountAvailableBalanceMinorUnits` | `int` |  |
| `sourceAccountBic` | `string` |  |
| `sourceAccountCurrency` | `string` |  |
| `sourceAccountIban` | `string` |  |
| `sourceAccountIdentifier` | `map[string]any` |  |
| `sourceAccountName` | `string` |  |
| `sourceAccountNumber` | `string` |  |
| `sourceAccountSortcode` | `string` |  |
| `status` | `string` |  |
| `tags` | `[]any` |  |
| `theirReference` | `string` |  |
| `topupPayrunID` | `string` |  |
| `transactedAmount` | `float64` |  |
| `transactedFxAmount` | `float64` |  |
| `transactedFxRate` | `float64` |  |
| `type` | `string` |  |
| `userID` | `string` |  |
| `yourReference` | `string` |  |

#### Example: List

```go
payoutKeysetPages, err := client.PayoutKeysetPage(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(payoutKeysetPages) // the array of records
```


### PayoutMetric

Create an instance: `payoutMetric := client.PayoutMetric(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Example: Load

```go
payoutMetric, err := client.PayoutMetric(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(payoutMetric) // the loaded record
```


### Payrun

Create an instance: `payrun := client.Payrun(nil)`

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
| `authorisationDate` | `string` |  |
| `authorisations` | `[]any` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `batchPayoutID` | `string` |  |
| `canAuthorise` | `bool` |  |
| `canDelete` | `bool` |  |
| `canEdit` | `bool` |  |
| `events` | `[]any` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `invoices` | `[]any` |  |
| `invoicesMinimal` | `[]any` |  |
| `isArchived` | `bool` |  |
| `lastUpdated` | `string` |  |
| `lastUpdatedBy` | `map[string]any` |  |
| `merchantID` | `string` |  |
| `name` | `string` |  |
| `nonce` | `string` |  |
| `notes` | `string` |  |
| `payments` | `[]any` |  |
| `payouts` | `[]any` |  |
| `payoutsCount` | `int` |  |
| `reason` | `string` |  |
| `scheduleDate` | `string` |  |
| `scheduledDate` | `string` |  |
| `sourceAccounts` | `[]any` |  |
| `status` | `string` |  |
| `totalEur` | `float64` |  |
| `totalGbp` | `float64` |  |
| `totalUsd` | `float64` |  |

#### Example: Load

```go
payrun, err := client.Payrun(nil).Load(map[string]any{"id": "payrun_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(payrun) // the loaded record
```

#### Example: List

```go
payruns, err := client.Payrun(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(payruns) // the array of records
```

#### Example: Create

```go
result, err := client.Payrun(nil).Create(map[string]any{
    "id": "example_id",
    "lastUpdatedBy": map[string]any{},
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Report

Create an instance: `report := client.Report(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |


### ReportResult

Create an instance: `reportResult := client.ReportResult(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `contentType` | `string` |  |
| `contents` | `string` |  |
| `lastCompletedAt` | `string` |  |
| `merchantID` | `string` |  |
| `reportName` | `string` |  |
| `reportType` | `string` |  |
| `statementNumber` | `int` |  |

#### Example: Load

```go
reportResult, err := client.ReportResult(nil).Load(map[string]any{"id": 1, "report_id": "report_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(reportResult) // the loaded record
```


### Role

Create an instance: `role := client.Role(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `failedRoles` | `map[string]any` |  |
| `roles` | `[]any` |  |

#### Example: Create

```go
result, err := client.Role(nil).Create(map[string]any{
    "merchant_id": "example_merchant_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Rule

Create an instance: `rule := client.Rule(nil)`

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
| `account` | `map[string]any` |  |
| `accountID` | `string` |  |
| `approveUrl` | `string` |  |
| `approverID` | `string` |  |
| `authenticationMethods` | `[]any` |  |
| `authorisations` | `[]any` |  |
| `authorisersCompletedCount` | `int` |  |
| `authorisersRequiredCount` | `int` |  |
| `canAuthorise` | `bool` |  |
| `createdBy` | `map[string]any` |  |
| `description` | `string` |  |
| `endAt` | `string` |  |
| `hasCurrentUserAuthorised` | `bool` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isDisabled` | `bool` |  |
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
| `sweepAction` | `map[string]any` |  |
| `timeZoneId` | `string` |  |
| `triggerCronExpression` | `string` |  |
| `triggerOnPayIn` | `bool` |  |
| `userID` | `string` |  |
| `webHookSecret` | `string` |  |

#### Example: Load

```go
rule, err := client.Rule(nil).Load(map[string]any{"id": "rule_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(rule) // the loaded record
```

#### Example: List

```go
rules, err := client.Rule(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(rules) // the array of records
```

#### Example: Create

```go
result, err := client.Rule(nil).Create(map[string]any{
    "createdBy": map[string]any{},
    "nonce": "example_nonce",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### RuleEvent

Create an instance: `ruleEvent := client.RuleEvent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `errorMessage` | `string` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isAuthoriseToEnable` | `bool` |  |
| `message` | `string` |  |
| `rawResponse` | `string` |  |
| `ruleEventType` | `string` |  |
| `ruleID` | `string` |  |
| `user` | `map[string]any` |  |

#### Example: List

```go
ruleEvents, err := client.RuleEvent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(ruleEvents) // the array of records
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
| `colourHex` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `merchantID` | `string` |  |
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
    "merchant_id": "example_merchant_id",
    "merchantID": "example_merchantID",
    "name": "example_name",
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
    "id": "example_id",
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
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `accountID` | `string` |  |
| `accountName` | `string` |  |
| `accountSequenceNumber` | `int` |  |
| `addressDetails` | `map[string]any` |  |
| `amount` | `float64` |  |
| `amountMinorUnits` | `int` |  |
| `balance` | `float64` |  |
| `balanceMinorUnits` | `int` |  |
| `bookingDateTime` | `string` |  |
| `chargeDetails` | `map[string]any` |  |
| `content` | `[]any` |  |
| `counterparty` | `map[string]any` |  |
| `counterpartySummary` | `string` |  |
| `currency` | `string` |  |
| `currencyExchange` | `map[string]any` |  |
| `date` | `string` |  |
| `description` | `string` |  |
| `enrichment` | `map[string]any` |  |
| `fxAmount` | `float64` |  |
| `fxCurrency` | `string` |  |
| `fxRate` | `float64` |  |
| `grossAmount` | `map[string]any` |  |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isoBankTransactionCode` | `map[string]any` |  |
| `merchant` | `map[string]any` |  |
| `merchantID` | `string` |  |
| `pageNumber` | `int` |  |
| `pageSize` | `int` |  |
| `payeeDetails` | `map[string]any` |  |
| `payerDetails` | `map[string]any` |  |
| `paymentRequestCustomFields` | `map[string]any` |  |
| `paymentRequestID` | `string` |  |
| `payoutID` | `string` |  |
| `proprietaryBankTransactionCode` | `map[string]any` |  |
| `rawReference` | `string` |  |
| `reference` | `string` |  |
| `ruleID` | `string` |  |
| `statementReferences` | `[]any` |  |
| `status` | `string` |  |
| `supplementaryData` | `any` |  |
| `tags` | `[]any` |  |
| `theirReference` | `string` |  |
| `totalPages` | `int` |  |
| `totalSize` | `int` |  |
| `transactionAmount` | `map[string]any` |  |
| `transactionDate` | `string` |  |
| `transactionInformation` | `[]any` |  |
| `transactionMutability` | `string` |  |
| `type` | `string` |  |
| `valueDateTime` | `string` |  |
| `virtualIBAN` | `string` |  |
| `yourReference` | `string` |  |

#### Example: Load

```go
transaction, err := client.Transaction(nil).Load(map[string]any{"id": "transaction_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(transaction) // the loaded record
```

#### Example: List

```go
transactions, err := client.Transaction(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(transactions) // the array of records
```

#### Example: Create

```go
result, err := client.Transaction(nil).Create(map[string]any{
    "id": "example_id",
    "grossAmount": map[string]any{},
    "payeeDetails": map[string]any{},
    "payerDetails": map[string]any{},
    "transactionAmount": map[string]any{},
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
| `clientSessionTimeouts` | `[]any` |  |
| `emailAddress` | `string` |  |
| `firstName` | `string` |  |
| `id` | `string` |  |
| `lastName` | `string` |  |
| `passkeyAdded` | `bool` |  |
| `permissions` | `map[string]any` |  |
| `profile` | `string` |  |
| `rolesWithScope` | `[]any` |  |
| `twoFactorEnabled` | `bool` |  |
| `userInviteID` | `string` |  |

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
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `authorisationStatus` | `map[string]any` |  |
| `failedUserInvites` | `map[string]any` |  |
| `id` | `string` |  |
| `initialRoleID` | `string` |  |
| `inviteeEmailAddress` | `string` |  |
| `inviteeFirstName` | `string` |  |
| `inviteeLastName` | `string` |  |
| `inviterEmailAddress` | `string` |  |
| `inviterFirstName` | `string` |  |
| `inviterLastName` | `string` |  |
| `isAuthorised` | `bool` |  |
| `isInviteeRegistered` | `bool` |  |
| `lastInvited` | `string` |  |
| `merchantID` | `string` |  |
| `merchantName` | `string` |  |
| `message` | `string` |  |
| `registrationUrl` | `string` |  |
| `sendInviteEmail` | `bool` |  |
| `status` | `string` |  |
| `user` | `map[string]any` |  |
| `userID` | `string` |  |
| `userInvites` | `[]any` |  |

#### Example: Load

```go
userInvite, err := client.UserInvite(nil).Load(map[string]any{"id": "user_invite_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(userInvite) // the loaded record
```

#### Example: List

```go
userInvites, err := client.UserInvite(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(userInvites) // the array of records
```

#### Example: Create

```go
result, err := client.UserInvite(nil).Create(map[string]any{
    "user": map[string]any{},
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
| `accountName` | `string` |  |
| `accountSupplierName` | `string` |  |
| `availableBalance` | `float64` |  |
| `availableBalanceMinorUnits` | `int` |  |
| `balance` | `float64` |  |
| `balanceMinorUnits` | `int` |  |
| `bankName` | `string` |  |
| `consentID` | `string` |  |
| `createdBy` | `map[string]any` |  |
| `createdByDisplayName` | `string` |  |
| `currency` | `string` |  |
| `defaultPaymentRail` | `string` |  |
| `displayName` | `string` |  |
| `expiryDate` | `string` |  |
| `externalAccountIcon` | `string` |  |
| `id` | `string` |  |
| `identifier` | `map[string]any` |  |
| `inserted` | `string` |  |
| `isArchived` | `bool` |  |
| `isConnectedAccount` | `bool` |  |
| `isDefault` | `bool` |  |
| `isTrustAccount` | `bool` |  |
| `isVirtual` | `bool` |  |
| `lastTransaction` | `map[string]any` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` |  |
| `merchantName` | `string` |  |
| `name` | `string` |  |
| `physicalAccountID` | `string` |  |
| `rules` | `[]any` |  |
| `submittedPayoutsBalance` | `float64` |  |
| `submittedPayoutsBalanceMinorUnits` | `int` |  |
| `summary` | `string` |  |
| `supplierSepaInstantStatus` | `string` |  |
| `xeroBankFeedConnectionStatus` | `string` |  |
| `xeroBankFeedLastSyncedAt` | `string` |  |
| `xeroBankFeedSyncLastFailedAt` | `string` |  |
| `xeroBankFeedSyncLastFailureReason` | `string` |  |
| `xeroBankFeedSyncStatus` | `string` |  |
| `xeroUnsynchronisedTransactionsCount` | `int` |  |

#### Example: Create

```go
result, err := client.Virtual(nil).Create(map[string]any{
    "account_id": "example_account_id",
    "createdBy": map[string]any{},
    "identifier": map[string]any{},
    "name": "example_name",
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
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `destinationUrl` | `string` |  |
| `emailAddress` | `string` |  |
| `failedNotificationEmailAddress` | `string` |  |
| `id` | `string` |  |
| `isActive` | `bool` |  |
| `merchantID` | `string` |  |
| `notificationMethod` | `string` |  |
| `resourceTypes` | `[]any` |  |
| `retry` | `bool` |  |
| `secret` | `string` |  |
| `version` | `int` |  |

#### Example: Load

```go
webhook, err := client.Webhook(nil).Load(map[string]any{"id": "webhook_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(webhook) // the loaded record
```

#### Example: List

```go
webhooks, err := client.Webhook(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(webhooks) // the array of records
```

#### Example: Create

```go
result, err := client.Webhook(nil).Create(map[string]any{
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
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
consent := client.Consent(nil)
consent.List(nil, nil)

// consent.Data() now returns the consent data from the last list
// consent.Match() returns the last match criteria
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
