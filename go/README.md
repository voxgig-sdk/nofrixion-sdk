# Nofrixion Golang SDK



The Golang SDK for the Nofrixion API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.Account(nil)` — each with the same small set of operations (`List`, `Load`, `Create`, `Update`, `Remove`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Also generated from this model: `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb`, `ts` — see
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
    created, err := client.Account(nil).Create(map[string]any{"account_id": "example_account_id", "currency": "example_currency", "createdBy": map[string]any{}, "identifier": map[string]any{}}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(created)

    // Update a account.
    updated, err := client.Account(nil).Update(map[string]any{"id": "example_id", "accountBalances": []any{}, "accountID": "example_accountID"}, nil)
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
| `"accountBalances"` | The various balances for the account. |
| `"accountID"` | ID of the account. |
| `"accountIdentifications"` | The canoncial identifiers for the account. |
| `"accountName"` | Name for the account |
| `"accountNames"` | Optional account names set by the account holder. |
| `"accountSupplierName"` | The payment account supplier name. |
| `"accountType"` | The type of account e.g. |
| `"availableBalance"` | The current available balance of the account. |
| `"availableBalanceMinorUnits"` | The available balance expressed in the currency’s minor units (e.g. |
| `"balance"` | Balance of the account. |
| `"balanceMinorUnits"` | Balance of the account expressed in the currency’s minor units (e.g. |
| `"bankName"` | The bank name for external accounts |
| `"consentID"` | The ID of the consent used to connect the external account. |
| `"consolidatedAccountInformation"` | Summary information regarding account balances of the overall account provided by the bank. |
| `"createdBy"` |  |
| `"createdByDisplayName"` | Either the name of the user, merchant token or api key that created the account |
| `"currency"` | Currency of the account in ISO 4217 format |
| `"defaultPaymentRail"` | Indicates the default payment rail for this account. |
| `"description"` | Product name as defined by the financial institution for this account. |
| `"details"` | Supplementary specifications that might be provided by the Bank. |
| `"displayName"` | Gets a unique display name for the payment account. |
| `"expiryDate"` | The date that the external account will expire |
| `"externalAccountIcon"` | The Icon for external accounts |
| `"format"` | File format to save the statement as. |
| `"fromDate"` | Minimum transaction date for the statement. |
| `"id"` | Unique id for the account. |
| `"identifier"` |  |
| `"inserted"` | Timestamp when the account was created. |
| `"isArchived"` | Indicates whether the account is archived. |
| `"isConnectedAccount"` | Indicates if the payment account is an externally connected account. |
| `"isDefault"` | Is the default account |
| `"isTrustAccount"` | Indicates if the payment account is a trust account. |
| `"isVirtual"` | True if the account is a virtual account. |
| `"lastTransaction"` |  |
| `"lastUpdated"` | Timestamp when the account was last updated. |
| `"merchantID"` | The ID of the merchant that owns the account. |
| `"merchantName"` | The name of the merchant that owns the account. |
| `"nickname"` | Nickname of the account that was provided by the account owner. |
| `"physicalAccountID"` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `"roleIDs"` | Optional list of role IDs that will get access to the payment account when created. |
| `"rules"` | The list of rules associated with this account. |
| `"submittedPayoutsBalance"` | Total of the payouts that have been submitted for processing. |
| `"submittedPayoutsBalanceMinorUnits"` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `"summary"` | Gets a summary of the payments account's most important properties. |
| `"supplierPhysicalAccountID"` | For internal use only. |
| `"supplierSepaInstantStatus"` | Indicates the status of the SEPA Instant payment rail for this account. |
| `"toDate"` | Maximum transaction date for the statement. |
| `"type"` | Specifies the type of account e.g. |
| `"usageType"` |  |
| `"xeroBankFeedConnectionStatus"` | States the status of the Xero bank feed connection, if applicable. |
| `"xeroBankFeedLastSyncedAt"` |  |
| `"xeroBankFeedSyncLastFailedAt"` |  |
| `"xeroBankFeedSyncLastFailureReason"` |  |
| `"xeroBankFeedSyncStatus"` |  |
| `"xeroUnsynchronisedTransactionsCount"` | Indicates the number of unsynchronised transactions with Xero |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/accounts/{accountID}/{currency}`

#### Batch

| Field | Description |
| --- | --- |
| `"approveUrl"` | This field is used when returning a batch payout record to a client. |
| `"id"` |  |
| `"payouts"` |  |

Operations: Create, Load.

API path: `/api/v1/payouts/batch`

#### Beneficiary

| Field | Description |
| --- | --- |
| `"approvalCallbackUrl"` |  |
| `"authenticationMethods"` | A list of authentication types allowed to authorise the payout. |
| `"authorisations"` | A list of users who have successfully authorised the latest version of the beneficiary. |
| `"authorisersCompletedCount"` | The number of distinct authorisers that have authorised the beneficiary. |
| `"authorisersRequiredCount"` | The number of authorisers required for this beneficiary. |
| `"beneficiaries"` |  |
| `"beneficiaryEvents"` |  |
| `"canAuthorise"` | True if the beneficiary can be authorised by the user who loaded it. |
| `"canUpdate"` | True if the beneficiary can be updated by the user who loaded it. |
| `"createdBy"` |  |
| `"createdByEmailAddress"` |  |
| `"currency"` | Gets or Sets the currency. |
| `"destination"` |  |
| `"failedBeneficiaries"` |  |
| `"hasCurrentUserAuthorised"` | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `"id"` |  |
| `"inserted"` |  |
| `"isEnabled"` |  |
| `"lastAuthorised"` |  |
| `"lastUpdated"` |  |
| `"merchantID"` | Gets or Sets the merchant id. |
| `"name"` | The descriptive name for the beneficiary. |
| `"nonce"` |  |
| `"sourceAccountIDs"` | ID of the accounts which are authorised to act as a source for the beneficiary. |
| `"sourceAccounts"` |  |
| `"theirReference"` | The reference that will be used by default as TheirReference when creating payouts to this beneficiary if no TheirReference is specified for the payout. |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/beneficiaries/authorise/{id}`

#### BeneficiaryGroup

| Field | Description |
| --- | --- |
| `"groupMembers"` | The existing group members. |
| `"groupName"` | The descriptive name for the beneficiary group. |
| `"id"` |  |
| `"inserted"` | Timestamp indicating when the group was created. |
| `"lastUpdated"` | Timestamp indicating when the group was last updated. |
| `"merchantID"` | Gets or Sets the merchant id. |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/beneficiarygroups`

#### Card

| Field | Description |
| --- | --- |
| `"authorizedAmount"` |  |
| `"currencyCode"` |  |
| `"isPayerAuthenticationRequired"` | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `"isSoftDecline"` | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `"payerAuthenticationAccessToken"` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `"payerAuthenticationMerchantData"` | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `"payerAuthenticationUrl"` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `"payerAuthenticationWindowHeight"` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `"payerAuthenticationWindowWidth"` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `"paymentRequestCallbackUrl"` | The callback URL that was set when the payment request was created. |
| `"paymentRequestID"` |  |
| `"requestID"` |  |
| `"responseCode"` |  |
| `"responseType"` |  |
| `"status"` |  |
| `"threeDSRedirectUrl"` | Checkout.com require a redirect for 3DS authentication. |
| `"transactionID"` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/{id}/card`

#### CardCustomerToken

| Field | Description |
| --- | --- |
| `"cardType"` | The type of the tokenised card, e.g. |
| `"customerEmailAddress"` | When creating a tokenised card the payer's email address must be supplied. |
| `"expiryMonth"` |  |
| `"expiryYear"` |  |
| `"id"` | The unique ID of the card token that has been stored for the customer. |
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
| `"isPayerAuthenticationRequired"` | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `"isSoftDecline"` | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `"payerAuthenticationAccessToken"` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `"payerAuthenticationMerchantData"` | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `"payerAuthenticationUrl"` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `"payerAuthenticationWindowHeight"` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `"payerAuthenticationWindowWidth"` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `"paymentRequestCallbackUrl"` | The callback URL that was set when the payment request was created. |
| `"paymentRequestID"` |  |
| `"requestID"` |  |
| `"responseCode"` |  |
| `"responseType"` |  |
| `"status"` |  |
| `"threeDSRedirectUrl"` | Checkout.com require a redirect for 3DS authentication. |
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
| `"authorisationUrl"` | The URL the authorising user needs to be redirected to in order to get the open banking consent token. |
| `"callbackUrl"` | Optional callback URL that the end user performing the open banking authorisation will be redirected to on completion. |
| `"consentID"` | The ID of the open banking consent. |
| `"emailAddress"` | The email address that identifies the end user that will be authorising the open banking consent request. |
| `"expiryDate"` |  |
| `"failureCallbackUrl"` | Optional callback URL for open banking consent authorisation failure. |
| `"id"` |  |
| `"inserted"` |  |
| `"institutionID"` | The institution ID the open banking consent is being requested for. |
| `"isConnectedAccounts"` | Optional setting. |
| `"isEnabled"` |  |
| `"merchantID"` | The ID of the merchant the consent token is being created to be used with. |
| `"provider"` | Lists the supported card and PIS processors. |
| `"successWebHookUrl"` | A web hook URL to send an HTTP request to when the open banking consent is successfuly authorised. |

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
| `"failedSubmissions"` | Dictionary of failed submissions, keyed by the index (1-based) in the original request. |
| `"successfulSubmissions"` | List of successfully submitted direct debit payments. |

Operations: Create.

API path: `/api/v1/paymentrequests/directdebit/batchsubmit`

#### FxRate

| Field | Description |
| --- | --- |
| `"destinationCurrency"` |  |
| `"exchangeRate"` | The price at which the transaction will buy the source currency using the destination currency. |
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
| `"accountNumber"` | Account number of the customer's bank account in case of GBP account. |
| `"addressLine1"` | First line of the customer's address. |
| `"addressLine2"` | Second line of the customer's address. |
| `"approvedAt"` | Date at which the supplier approved this mandate. |
| `"city"` | Customer's city. |
| `"countryCode"` | 2-character country code of the customer's bank account. |
| `"currency"` | Currency of this mandate. |
| `"customerAccountNumber"` | Customer's account number in case of GBP account. |
| `"customerCity"` | Customer's city of residence. |
| `"customerCountryCode"` | Customer's country of residence code. |
| `"customerCountryName"` | Customer's country of residence. |
| `"customerEmailAddress"` | Customer's email address. |
| `"customerFirstName"` | Customer's first name. |
| `"customerIban"` | Customer's IBAN in case of EUR account. |
| `"customerLastName"` | Customer's last name. |
| `"customerSortCode"` | Customer's sort code in case of GBP account. |
| `"emailAddress"` | Customer's email address. |
| `"firstName"` | Customer's first name. |
| `"iban"` | IBAN of the customer's bank account in case of EUR account. |
| `"id"` | Internal ID of the mandate. |
| `"inserted"` | The timestamp this mandate was created at. |
| `"isRecurring"` | Whether this mandate is single-use or recurring. |
| `"lastName"` | Customer's last name. |
| `"lastUpdated"` | The timestamp this mandate was last updated at. |
| `"merchantID"` | Internal ID of this mandate's merchant. |
| `"postalCode"` | Customer's postal code. |
| `"reference"` | Reference assigned to this mandate. |
| `"sortCode"` | Sort code of the customer's bank account in case of GBP account. |
| `"status"` | General status of this mandate. |
| `"supplierBankAccountID"` | ID that the supplier assigned to this mandate's bank account. |
| `"supplierCustomerID"` | ID that the supplier assigned to this mandate's customer. |
| `"supplierMandateID"` | ID that the supplier assigned to this mandate. |
| `"supplierName"` | Name of the supplier used to create this mandate. |
| `"supplierStatus"` | Last status that the supplier reported for this mandate. |

Operations: Create, Load.

API path: `/api/v1/mandates`

#### Merchant

| Field | Description |
| --- | --- |
| `"accountCurrencies"` | The list of currencies that the merchant has accounts for. |
| `"canHaveTrustAccounts"` | Trust accounts are a special type of account that allow the account name to be trusted for use in statements and verification of payee checks. |
| `"cardPaymentProcessor"` | Name of the card payment processor. |
| `"companyID"` | The Company ID recorded in the Compliance system. |
| `"displayQrOnHostedPay"` | Indicates if a QR Code containing the payment link should be displayed on the hosted payment page. |
| `"hostedPayVersion"` | The version of the hosted payment page to use with the merchant. |
| `"id"` | Unique ID for the merchant. |
| `"inserted"` | Timestamp the merchant was added to MoneyMoov. |
| `"isBlocked"` | The merchant is blocked from making payments (payouts). |
| `"isExited"` | The merchant has formally terminated their relationship and is no longer a customer. |
| `"isSuspended"` | The merchant has temporarily suspended their own account. |
| `"jurisdiction"` | The jurisdiction the merchant entity is incorporated or established in. |
| `"logoUrlPng"` | The CDN URL of the merchant's logo in PNG format. |
| `"logoUrlSvg"` | The CDN URL of the merchant's logo in SVG format. |
| `"merchantCategoryCode"` | The industry code that represents the merchant's primary trading activity. |
| `"name"` | The registered business name of the merchant. |
| `"notes"` | The notes field is an optional free text field that can be used to store any additional information about the merchant. |
| `"parentMerchant"` |  |
| `"paymentAccountLimit"` | The maximum number of payment accounts that can be created for the Merchant. |
| `"paymentAccounts"` |  |
| `"reason"` | The reason for the suspension. |
| `"shortName"` | A URL friendly shortish name for the merchant. |
| `"supportedPaymentMethodsList"` | The payment methods that are configured and supported for this merchant. |
| `"suspensionReason"` | The reason for the suspension, provided by the merchant. |
| `"tags"` | An optional list of descriptive tags that can be used on merchant entities such as payment requests. |
| `"timeZoneId"` | The IANA (Internet Assigned Numbers Authority) time zone identifier of the merchant. |
| `"tradingName"` | An optional trading name. |
| `"webHookLimit"` | The maximum number of web hooks that can be created for the Merchant. |
| `"yourRoleName"` | The name of the role for the identity that loaded the merchant record. |

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
| `"approvedAt"` | Date at which the supplier approved this mandate. |
| `"currency"` | Currency of this mandate. |
| `"customerAccountNumber"` | Customer's account number in case of GBP account. |
| `"customerCity"` | Customer's city of residence. |
| `"customerCountryCode"` | Customer's country of residence code. |
| `"customerCountryName"` | Customer's country of residence. |
| `"customerEmailAddress"` | Customer's email address. |
| `"customerFirstName"` | Customer's first name. |
| `"customerIban"` | Customer's IBAN in case of EUR account. |
| `"customerLastName"` | Customer's last name. |
| `"customerSortCode"` | Customer's sort code in case of GBP account. |
| `"id"` | Internal ID of the mandate. |
| `"inserted"` | The timestamp this mandate was created at. |
| `"isRecurring"` | Whether this mandate is single-use or recurring. |
| `"lastUpdated"` | The timestamp this mandate was last updated at. |
| `"merchantID"` | Internal ID of this mandate's merchant. |
| `"reference"` | Reference assigned to this mandate. |
| `"status"` | General status of this mandate. |
| `"supplierBankAccountID"` | ID that the supplier assigned to this mandate's bank account. |
| `"supplierCustomerID"` | ID that the supplier assigned to this mandate's customer. |
| `"supplierMandateID"` | ID that the supplier assigned to this mandate. |
| `"supplierName"` | Name of the supplier used to create this mandate. |
| `"supplierStatus"` | Last status that the supplier reported for this mandate. |

Operations: List.

API path: `/api/v1/mandates`

#### MerchantPayByBankSetting

| Field | Description |
| --- | --- |
| `"bankCountryCodes"` | The list of country codes representing the banks the country supports. |
| `"bankID"` | ID of the bank to be configured for the merchant. |
| `"bankName"` | Name of the Bank/Institution. |
| `"businessInstitutionID"` | ID that the processor uses to identify the bank (business accounts). |
| `"currency"` | Currency supported by the bank. |
| `"logo"` | URL of the bank's logo. |
| `"message"` | Message relating to specific bank. |
| `"messageImageUrl"` | Optional image URL to be displayed with the message. |
| `"order"` | Order in which this setting will appear in the UI. |
| `"personalInstitutionID"` | ID that the processor uses to identify the bank (personal accounts). |
| `"processor"` | Name of the bank payment processor. |
| `"warningHeading"` | The heading for a warning message related to the bank institution to be displayed to the user. |
| `"warningMessage"` | The warning message related to the bank institution to be displayed to the user. |

Operations: List.

API path: `/api/v1/merchants/{merchantID}/banksettings`

#### MerchantPaymentRequestTemplate

| Field | Description |
| --- | --- |
| `"bankPaymentOptions"` |  |
| `"cardPaymentAddressOptions"` |  |
| `"cardPaymentCaptureOptions"` |  |
| `"customFields"` | A list of custom fields that can be included in the payment request template. |
| `"defaultFields"` | A list of default fields that are included in the payment request template. |
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
| `"authenticationMethods"` | A list of authentication types allowed to authorise the merchant token. |
| `"authorisations"` | A list of users who have successfully authorised the latest version of the beneficiary. |
| `"authorisersCompletedCount"` | The number of distinct authorisers that have authorised the merchant token. |
| `"authorisersRequiredCount"` | The number of authorisers required for this merchant token. |
| `"canAuthorise"` | True if the merchant token can be authorised by the user who loaded it. |
| `"description"` | Token description |
| `"expiresAt"` | Optional. |
| `"hasCurrentUserAuthorised"` | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `"hmacAlgorithm"` | Optional shared secret algorithm to use for HMAC authentication. |
| `"id"` |  |
| `"inserted"` |  |
| `"ipAddressWhitelist"` | Optional. |
| `"isArchived"` | Indicates whether the merchant token is archived. |
| `"isEnabled"` | If set to false the merchant token will not be accepted to authorise a request. |
| `"lastAuthorised"` |  |
| `"lastUpdated"` |  |
| `"merchantID"` | The merchant id to add to the token |
| `"nonce"` |  |
| `"permissionTypes"` | The permissions that the merchant token supports. |
| `"requestSignatureVersion"` | Represent the version of the overall merchant token. |
| `"sharedSecretAlgorithm"` | Optional shared secret algorithm to use for HMAC authentication. |
| `"sharedSecretBase64"` | The base 64 encoded shared secret that is used for request authentication with an HMAC. |
| `"token"` | The JWT merchant token. |

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
| `"accountName"` | The name of the account to verify |
| `"accountNumber"` | The account number of the account to verify (for CoP checks) |
| `"iban"` | The IBAN of the account to verify (for VoP checks) |
| `"payeeVerifiedAccountName"` | The verified account name of the payee, if available (in case of a close match) |
| `"result"` | The result of the payee verification |
| `"secondaryIdentification"` | Optional secondary identifier for the account to verify. |
| `"sortCode"` | The sort code of the account to verify (for CoP checks) |

Operations: Create.

API path: `/api/v1/openbanking/payeeverification`

#### Payment

| Field | Description |
| --- | --- |
| `"addresses"` |  |
| `"amount"` | The amount of money to request. |
| `"amountPending"` | Total amount that has been authorised but not settled for this payment request. |
| `"amountReceived"` | Total amount received for this payment request. |
| `"amountRefunded"` | Total amount refunded for this payment request. |
| `"autoSendReceipt"` | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `"baseOriginUrl"` | For card payments the origin of the payment page needs to be set in advance. |
| `"callbackUrl"` | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `"cardAuthorizeOnly"` | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `"cardCreateToken"` | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `"cardCreateTokenMode"` | This specifies whether user consent will be taken before tokenising card or not. |
| `"cardIgnoreCVN"` | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `"cardNoPayerAuthentication"` | If set to true for card payments no attempt will be made to use payer authentication (3-D Secure and equivalent). |
| `"cardProcessorMerchantID"` | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `"cardStripePaymentIntentID"` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `"cardStripePaymentIntentSecret"` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `"cardTransmitRawDetails"` | If set to true for card payments the sensitive card number and card verification number will be transmitted directly rather than being tokenised. |
| `"createdByUser"` |  |
| `"currency"` | The currency of the request. |
| `"customFields"` | A list of custom fields attached to the payment request. |
| `"customerEmailAddress"` | Optional email address for the customer. |
| `"customerID"` | An optional customer identifier for the payment request. |
| `"customerName"` |  |
| `"description"` | An optional description for the payment request. |
| `"destinationAccount"` |  |
| `"directDebitPayment"` | Contains information about a Direct Debit payment attempt for a payment request. |
| `"dueDate"` | The due date for the payment request. |
| `"events"` |  |
| `"failureCallbackUrl"` | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `"fieldDisplaySettings"` | A list of field display settings that control which fields are displayed to the payer. |
| `"formattedAmount"` |  |
| `"hostedPayCheckoutUrl"` | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `"id"` |  |
| `"ignoreAddressVerification"` | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `"inserted"` | The timestamp the payment request was created at. |
| `"insertedSortable"` | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `"isArchived"` | Indicates whether the payment request is archived. |
| `"jwk"` | The jwk containing the public key used to verify the signature of the payment request. |
| `"lastUpdated"` | The timestamp the payment request was last updated at. |
| `"lightningInvoice"` | Bitcoin Lightning invoice for the payment request. |
| `"lightningInvoiceExpiresAt"` | Date and time of expiration of the lightning invoice. |
| `"merchantDirectDebitMandateID"` | Optional ID of the direct debit mandate associated with this payment request. |
| `"merchantID"` | The ID of the merchant to create the payment request for. |
| `"merchantTokenDescription"` | Description of the merchant token in case the Payment request was created using a merchant token. |
| `"notificationEmailAddresses"` |  |
| `"notificationRoleIDs"` | A list of roles whose members will receive notifications about this payment request. |
| `"orderID"` | An optional order ID for the payment request. |
| `"partialPaymentMethod"` | The approach to use, or not, for accepting partial payments. |
| `"partialPaymentSteps"` | An optional comma separated list of partial payment amounts. |
| `"paymentAttempts"` | The payment attempts made against this payment request. |
| `"paymentMethods"` | The payment methods that the payment request supports. |
| `"paymentProcessor"` | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `"payrunID"` | The ID of a payrun that needs an account top up. |
| `"pispAccountID"` | The payment account ID to use to receive payment initiation payments. |
| `"priorityBankID"` | The ID of the bank that is set as the priority bank for display on pay element. |
| `"result"` |  |
| `"sandboxSettleDelayInSeconds"` | Sandbox only. |
| `"shippingAddress"` |  |
| `"shippingAddressCity"` | Optionally the city of the customer's shipping address. |
| `"shippingAddressCountryCode"` | Optionally the country code of the customer's shipping address. |
| `"shippingAddressCounty"` | Optionally the state or county of the customer's shipping address. |
| `"shippingAddressLine1"` | Optionally the first line of the customer's shipping address. |
| `"shippingAddressLine2"` | Optionally the second line of the customer's shipping address. |
| `"shippingAddressPostCode"` | Optionally the post code of the customer's shipping address. |
| `"shippingEmail"` | Optionally the shipping email address for the customer. |
| `"shippingFirstName"` | Optionally the first name of the customer's shipping address. |
| `"shippingLastName"` | Optionally the last name of the customer's shipping address. |
| `"shippingPhone"` | Optionally the shipping phone number for the customer. |
| `"status"` | The current status of the payment request. |
| `"successWebHookUrl"` | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `"tagIds"` | An optional list of tag ids to add to the payment request |
| `"tags"` | An optional list of descriptive tags attached to the payment request. |
| `"title"` | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `"tokenisedCards"` |  |
| `"transactions"` |  |
| `"useHostedPaymentPage"` | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

Operations: Create, Load, Update.

API path: `/api/v1/paymentrequests`

#### PaymentAccount

| Field | Description |
| --- | --- |
| `"accountName"` | Name for the account |
| `"accountSupplierName"` | The payment account supplier name. |
| `"availableBalance"` | The current available balance of the account. |
| `"availableBalanceMinorUnits"` | The available balance expressed in the currency’s minor units (e.g. |
| `"balance"` | Balance of the account. |
| `"balanceMinorUnits"` | Balance of the account expressed in the currency’s minor units (e.g. |
| `"bankName"` | The bank name for external accounts |
| `"consentID"` | The ID of the consent used to connect the external account. |
| `"createdBy"` |  |
| `"createdByDisplayName"` | Either the name of the user, merchant token or api key that created the account |
| `"currency"` | Currency of the account in ISO 4217 format |
| `"defaultPaymentRail"` | Indicates the default payment rail for this account. |
| `"displayName"` | Gets a unique display name for the payment account. |
| `"expiryDate"` | The date that the external account will expire |
| `"externalAccountIcon"` | The Icon for external accounts |
| `"id"` | Unique id for the account. |
| `"identifier"` |  |
| `"inserted"` | Timestamp when the account was created. |
| `"isArchived"` | Indicates whether the account is archived. |
| `"isConnectedAccount"` | Indicates if the payment account is an externally connected account. |
| `"isDefault"` | Is the default account |
| `"isTrustAccount"` | Indicates if the payment account is a trust account. |
| `"isVirtual"` | True if the account is a virtual account. |
| `"lastTransaction"` |  |
| `"lastUpdated"` | Timestamp when the account was last updated. |
| `"merchantID"` | The ID of the merchant that owns the account. |
| `"merchantName"` | The name of the merchant that owns the account. |
| `"physicalAccountID"` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `"rules"` | The list of rules associated with this account. |
| `"submittedPayoutsBalance"` | Total of the payouts that have been submitted for processing. |
| `"submittedPayoutsBalanceMinorUnits"` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `"summary"` | Gets a summary of the payments account's most important properties. |
| `"supplierSepaInstantStatus"` | Indicates the status of the SEPA Instant payment rail for this account. |
| `"xeroBankFeedConnectionStatus"` | States the status of the Xero bank feed connection, if applicable. |
| `"xeroBankFeedLastSyncedAt"` |  |
| `"xeroBankFeedSyncLastFailedAt"` |  |
| `"xeroBankFeedSyncLastFailureReason"` |  |
| `"xeroBankFeedSyncStatus"` |  |
| `"xeroUnsynchronisedTransactionsCount"` | Indicates the number of unsynchronised transactions with Xero |

Operations: List.

API path: `/api/v1/accounts/paged`

#### PaymentAccountMinimal

| Field | Description |
| --- | --- |
| `"accountName"` | Name for the account |
| `"availableBalance"` | The current available balance of the account. |
| `"balance"` | Balance of the account. |
| `"balanceMinorUnits"` | Balance of the account expressed in the currency’s minor units (e.g. |
| `"currency"` | Currency of the account in ISO 4217 format |
| `"id"` | Unique id for the account. |
| `"identifier"` |  |
| `"isArchived"` | Is the account archived |
| `"isConnectedAccount"` | Indicates if the payment account is an externally connected account. |
| `"merchantID"` | The ID of the merchant that owns the account. |
| `"submittedPayoutsBalance"` | Total of the payouts that have been submitted for processing. |

Operations: List.

API path: `/api/v1/accounts/minimal`

#### PaymentInitiation

| Field | Description |
| --- | --- |
| `"paymentInitiationID"` | The unique identifier of the payment initiation request. |
| `"paymentRequestCallbackUrl"` | The callback URL that was set when the payment request was created. |
| `"paymentRequestID"` |  |
| `"redirectUrl"` | A redirect URL for the user to authorise the payment initiation request at the ASPSP |
| `"responseType"` |  |
| `"specificErrorMessage"` |  |

Operations: Create.

API path: `/api/v1/paymentrequests/{id}/pisp`

#### PaymentRequest

| Field | Description |
| --- | --- |
| `"addresses"` |  |
| `"amount"` | The amount of money to request. |
| `"amountPending"` | Total amount that has been authorised but not settled for this payment request. |
| `"amountReceived"` | Total amount received for this payment request. |
| `"amountRefunded"` | Total amount refunded for this payment request. |
| `"autoSendReceipt"` | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `"baseOriginUrl"` | For card payments the origin of the payment page needs to be set in advance. |
| `"callbackUrl"` | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `"cardAuthorizeOnly"` | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `"cardCreateToken"` | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `"cardCreateTokenMode"` | This specifies whether user consent will be taken before tokenising card or not. |
| `"cardIgnoreCVN"` | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `"cardProcessorMerchantID"` | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `"cardStripePaymentIntentID"` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `"cardStripePaymentIntentSecret"` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `"createdByUser"` |  |
| `"currency"` | The currency of the request. |
| `"customFields"` | A list of custom fields attached to the payment request. |
| `"customerEmailAddress"` | Optional email address for the customer. |
| `"customerID"` | An optional customer identifier for the payment request. |
| `"customerName"` |  |
| `"description"` | An optional description for the payment request. |
| `"destinationAccount"` |  |
| `"directDebitPayment"` | Contains information about a Direct Debit payment attempt for a payment request. |
| `"doSimulateSettlementFailure"` |  |
| `"dueDate"` | The due date for the payment request. |
| `"errorDescription"` |  |
| `"events"` |  |
| `"failedPaymentRequests"` |  |
| `"failureCallbackUrl"` | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `"fieldDisplaySettings"` | A list of field display settings that control which fields are displayed to the payer. |
| `"formattedAmount"` |  |
| `"hostedPayCheckoutUrl"` | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `"id"` |  |
| `"ignoreAddressVerification"` | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `"inserted"` | The timestamp the payment request was created at. |
| `"insertedSortable"` | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `"institution"` |  |
| `"isArchived"` | Indicates whether the payment request is archived. |
| `"jwk"` | The jwk containing the public key used to verify the signature of the payment request. |
| `"lastUpdated"` | The timestamp the payment request was last updated at. |
| `"lightningInvoice"` | Bitcoin Lightning invoice for the payment request. |
| `"lightningInvoiceExpiresAt"` | Date and time of expiration of the lightning invoice. |
| `"merchantDirectDebitMandateID"` | Optional ID of the direct debit mandate associated with this payment request. |
| `"merchantID"` |  |
| `"merchantTokenDescription"` | Description of the merchant token in case the Payment request was created using a merchant token. |
| `"notificationEmailAddresses"` |  |
| `"notificationRoleIDs"` | A list of roles whose members will receive notifications about this payment request. |
| `"orderID"` | An optional order ID for the payment request. |
| `"partialPaymentMethod"` | The approach to use, or not, for accepting partial payments. |
| `"partialPaymentSteps"` | An optional comma separated list of partial payment amounts. |
| `"paymentAttempts"` | The payment attempts made against this payment request. |
| `"paymentInitiationID"` |  |
| `"paymentMethods"` | The payment methods that the payment request supports. |
| `"paymentProcessor"` | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `"paymentRequests"` |  |
| `"payrunID"` | The ID of a payrun that needs an account top up. |
| `"pispAccountID"` | The payment account ID to use to receive payment initiation payments. |
| `"priorityBankID"` | The ID of the bank that is set as the priority bank for display on pay element. |
| `"result"` |  |
| `"sandboxSettleDelayInSeconds"` | Sandbox only. |
| `"shippingAddress"` |  |
| `"status"` | The current status of the payment request. |
| `"successWebHookUrl"` | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `"tags"` | An optional list of descriptive tags attached to the payment request. |
| `"title"` | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `"tokenisedCards"` |  |
| `"transactions"` |  |
| `"useHostedPaymentPage"` | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/paymentrequests/{id}/directdebit`

#### PaymentRequestEvent

| Field | Description |
| --- | --- |
| `"amount"` |  |
| `"applePayTransactionID"` | Transaction ID received in Apple pay token. |
| `"cardAuthorizationResponseID"` | For a successful card authorization this field will hold the response ID. |
| `"cardExpiryMonth"` | For card payment events this field holds the payer's card expiry month. |
| `"cardExpiryYear"` | For card payment events this field holds the payer's card expiry year. |
| `"cardIssuer"` | For card payment events this field holds the payer's card issuer. |
| `"cardIssuerCountry"` | For card payment events this field holds the payer's card issuer country of origin. |
| `"cardLastFourDigits"` | For card payment events this field holds the payer's card last four digits. |
| `"cardRequestID"` |  |
| `"cardScheme"` | For card payment events this field holds the scheme of the payer's card, e.g. |
| `"cardTokenCustomerID"` | If the option to create a reusable token for card payments was set this field contains the token the merchant can store to use for repeat payments. |
| `"cardTransactionID"` |  |
| `"currency"` |  |
| `"directDebitPaymentID"` | Payment ID issued by the Direct Debit supplier. |
| `"directDebitPaymentReference"` | Reference string issued by the Direct Debit supplier. |
| `"drirectDebitMandateID"` | The ID of the mandate that was used wehn requesting payment. |
| `"errorMessage"` |  |
| `"errorReason"` |  |
| `"eventType"` |  |
| `"id"` |  |
| `"inserted"` |  |
| `"lightningInvoice"` | For Bitcoin Lightning payments this field holds the invoice presented to the payer. |
| `"lightningRHash"` | For Bitcoin Lightning payments the hash of the invoice presented to the payer. |
| `"originUrl"` | Optional field that can be set by payment methods, such as pay by bank, that may want to redirect back to the URL that initiated the attempt in the case of a failure condition. |
| `"paymentMethodType"` | The type of payment method the event relates to, e.g. |
| `"paymentProcessorName"` | If the event was for a card payment this is the name of the card processor, e.g. |
| `"paymentRequestID"` |  |
| `"pispBankStatus"` | For payment initiation attempts some providers (e.g. |
| `"pispPaymentInitiationID"` | For a payment initiation this is the ID returned by the service provider initiating the payment for us. |
| `"pispPaymentInstitutionName"` | For a payment initiation this is the name of the financial institution that is used to initiate and authorise the payment. |
| `"pispPaymentServiceProviderID"` | For a payment initiation this is the service provider ID selected by the payer, typically the ID for the bank or similar financial institution. |
| `"pispRedirectUrl"` | For a payment initiation this is the redirect URL returned by the service provider initiating the payment for us. |
| `"reconciledTransactionID"` | For settlement events (only relevant for non-card payments) this is the payin transaction that the payment request event was reconciled with. |
| `"refundPayoutID"` | ID of the Payout that was created for refund. |
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
| `"amount"` | The amount of money to request. |
| `"amountPending"` | The amount of money that was authorised but has not arrived in the account yet. |
| `"amountReceived"` | The amount of money that has been received for this payment request. |
| `"amountRefunded"` | The amount of money that has been refunded for this payment request. |
| `"callbackUrl"` |  |
| `"cardStripePaymentIntentSecret"` |  |
| `"countryCode"` | The country code associated with the payment. |
| `"currency"` | The currency of the request. |
| `"customFieldsToDisplay"` | Custom fields to display to the customer. |
| `"description"` | An optional description for the payment request. |
| `"dueDate"` | The due date of the payment request. |
| `"fieldDisplaySettings"` |  |
| `"googlePayMerchantID"` | Merchant ID from Google Pay |
| `"id"` |  |
| `"jwk"` | The jwk containing the public key |
| `"merchantID"` |  |
| `"merchantLogoUrlPng"` |  |
| `"merchantLogoUrlSvg"` |  |
| `"merchantName"` |  |
| `"merchantShortName"` |  |
| `"partialPaymentMethod"` |  |
| `"paymentAttempts"` | The payment attempts for this payment request. |
| `"paymentMethodsList"` | The payment methods that the payment request supports. |
| `"paymentProcessor"` | The card processor |
| `"paymentProcessorKey"` | The card processors public key |
| `"pispError"` | This is the error returned from the bank which is recorded in payment request events. |
| `"priorityBankID"` |  |
| `"status"` | The status of the payment request. |
| `"stripeAccountID"` | Account ID of connected customers in Stripe |
| `"title"` | The title of the payment request. |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/minimal`

#### PaymentRequestResult

| Field | Description |
| --- | --- |
| `"amount"` | The authorised payment amount. |
| `"amountPending"` |  |
| `"amountReceived"` |  |
| `"amountRefunded"` |  |
| `"currency"` | The authorised payment currency. |
| `"customerID"` | The customer id |
| `"paymentRequestID"` | The ID of the payment request the result is for. |
| `"payments"` | The list of payment attempts that have been received for the payment request. |
| `"pispAuthorizations"` |  |
| `"requestedAmount"` | The full original payment amount requested. |
| `"result"` | The result of the payment attempt. |

Operations: List.

API path: `/api/v1/paymentrequests/{id}/result`

#### Payout

| Field | Description |
| --- | --- |
| `"accountID"` | Gets or Sets Account Id of sending account |
| `"allowIncomplete"` | If set to true the payout will get created even if the business validation rules fail. |
| `"amount"` | Gets or Sets payout amount |
| `"amountMinorUnits"` | The payout amount expressed in the currency’s minor units (e.g. |
| `"approvePayoutUrl"` | This field is used when returning an payout record to a client. |
| `"approverID"` | Gets the User ID of person that approved the payout. |
| `"authenticationMethods"` | A list of authentication types allowed to authorise the payout. |
| `"authorisations"` | A list of the users who have successfully authorised the latest version of the payout and when. |
| `"authorisersCompletedCount"` | The number of distinct authorisers that have authorised the payout. |
| `"authorisersRequiredCount"` | The number of authorisers required for this payout. |
| `"batchPayoutID"` | The ID of the batch the payout is associated with. |
| `"beneficiary"` |  |
| `"beneficiaryID"` | Optional. |
| `"canAuthorise"` | True if the payout can be authorised by the user who loaded it. |
| `"canProcess"` | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `"canUpdate"` | True if the payout can be updated by the user who loaded it. |
| `"chargeBearer"` | Optional field to set who should pay any fees for the payout. |
| `"createdBy"` |  |
| `"createdByEmailAddress"` |  |
| `"currency"` | Gets or Sets Currency of payout request |
| `"currentUserID"` | The ID of the user that requested access to the PayOut record. |
| `"description"` | Gets or Sets description of payout request |
| `"destination"` |  |
| `"documents"` | Documents associated with the payout. |
| `"events"` | The activity associated with the payout. |
| `"failedPayouts"` |  |
| `"formattedAmount"` | Currency and formatted amount string. |
| `"formattedFxDestinationAmount"` | FX destination currency and amount formatted string. |
| `"formattedSchedule"` |  |
| `"formattedScheduleDayOnly"` |  |
| `"formattedSourceAccountAvailableBalance"` | The available balance of the account the payout is being made from. |
| `"fxDestinationAmount"` | If specified this will be the amount sent to the payee. |
| `"fxDestinationAmountMinorUnits"` | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `"fxDestinationCurrency"` | For an FX payout this is the currency to send to the beneficiary. |
| `"fxQuoteExpiresAt"` | If an FX held rate quote ID is being used this is the time the quote expires. |
| `"fxQuoteID"` | Optional. |
| `"fxRate"` | For an FX payout this is the exchange rate to use for the payout. |
| `"fxUseDestinationAmount"` | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `"hasCurrentUserAuthorised"` | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `"id"` | The ID for the payout. |
| `"inserted"` |  |
| `"invoiceID"` | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `"isArchived"` | Indicates whether the payout is archived. |
| `"isFailed"` | Set to true if a submitted payout subsequently fails. |
| `"isSettled"` | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `"isSubmitted"` | Indicates whether the payout has been submitted for processing. |
| `"lastUpdated"` |  |
| `"merchantID"` | The ID of the merchant that owns the account. |
| `"merchantTokenDescription"` |  |
| `"nonce"` |  |
| `"paymentProcessor"` | The usptream payment processor for the payout. |
| `"paymentRail"` | Optional field to indicate the payment rail to use for the payout. |
| `"payouts"` |  |
| `"payrunID"` | The ID of the payrun that this payout is associated with. |
| `"payrunName"` | The name of the payrun that this payout is associated with. |
| `"reason"` |  |
| `"rule"` |  |
| `"scheduleDate"` | The date the payout should be submitted. |
| `"scheduled"` | Should this payout be scheduled for a future date? |
| `"sourceAccountAvailableBalance"` | The available balance of the account the payout is being made from. |
| `"sourceAccountAvailableBalanceMinorUnits"` | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `"sourceAccountBic"` | The BIC of the account the payout is being made from. |
| `"sourceAccountCurrency"` | The currency of the source account. |
| `"sourceAccountIban"` | The IBAN of the account the payout is being made from. |
| `"sourceAccountIdentifier"` |  |
| `"sourceAccountName"` | The name of the account the payout is being made from. |
| `"sourceAccountNumber"` | The account number of the account the payout is being made from. |
| `"sourceAccountSortcode"` | The sort code of the account the payout is being made from. |
| `"status"` | Gets or Sets the status of payout request |
| `"tagIds"` | An optional list of tag ids to add to the payout. |
| `"tags"` | An optional list of descriptive tags attached to the payout. |
| `"theirReference"` | Gets or Sets destination reference ID |
| `"topupPayrunID"` | The ID of a payrun that needs an account top up. |
| `"transactedAmount"` | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `"transactedFxAmount"` | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `"transactedFxRate"` | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `"type"` | Gets or Sets payout type |
| `"userID"` | Gets or Sets User ID of who created the payout request |
| `"yourReference"` | Gets or Sets your reference ID |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/payouts/batch/submit/{id}`

#### PayoutKeysetPage

| Field | Description |
| --- | --- |
| `"accountID"` | Gets or Sets Account Id of sending account |
| `"amount"` | Gets or Sets payout amount |
| `"amountMinorUnits"` | The payout amount expressed in the currency’s minor units (e.g. |
| `"approvePayoutUrl"` | This field is used when returning an payout record to a client. |
| `"approverID"` | Gets the User ID of person that approved the payout. |
| `"authenticationMethods"` | A list of authentication types allowed to authorise the payout. |
| `"authorisations"` | A list of the users who have successfully authorised the latest version of the payout and when. |
| `"authorisersCompletedCount"` | The number of distinct authorisers that have authorised the payout. |
| `"authorisersRequiredCount"` | The number of authorisers required for this payout. |
| `"batchPayoutID"` | The ID of the batch the payout is associated with. |
| `"beneficiary"` |  |
| `"canAuthorise"` | True if the payout can be authorised by the user who loaded it. |
| `"canProcess"` | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `"canUpdate"` | True if the payout can be updated by the user who loaded it. |
| `"chargeBearer"` | Optional field to set who should pay any fees for the payout. |
| `"createdBy"` |  |
| `"createdByEmailAddress"` |  |
| `"currency"` | Gets or Sets Currency of payout request |
| `"currentUserID"` | The ID of the user that requested access to the PayOut record. |
| `"description"` | Gets or Sets description of payout request |
| `"destination"` |  |
| `"documents"` | Documents associated with the payout. |
| `"events"` | The activity associated with the payout. |
| `"formattedAmount"` | Currency and formatted amount string. |
| `"formattedFxDestinationAmount"` | FX destination currency and amount formatted string. |
| `"formattedSchedule"` |  |
| `"formattedScheduleDayOnly"` |  |
| `"formattedSourceAccountAvailableBalance"` | The available balance of the account the payout is being made from. |
| `"fxDestinationAmount"` | If specified this will be the amount sent to the payee. |
| `"fxDestinationAmountMinorUnits"` | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `"fxDestinationCurrency"` | For an FX payout this is the currency to send to the beneficiary. |
| `"fxQuoteExpiresAt"` | If an FX held rate quote ID is being used this is the time the quote expires. |
| `"fxQuoteID"` | Optional. |
| `"fxRate"` | For an FX payout this is the exchange rate to use for the payout. |
| `"fxUseDestinationAmount"` | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `"hasCurrentUserAuthorised"` | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `"id"` | The ID for the payout. |
| `"inserted"` |  |
| `"invoiceID"` | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `"isArchived"` | Indicates whether the payout is archived. |
| `"isFailed"` | Set to true if a submitted payout subsequently fails. |
| `"isSettled"` | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `"isSubmitted"` | Indicates whether the payout has been submitted for processing. |
| `"lastUpdated"` |  |
| `"merchantID"` | The ID of the merchant that owns the account. |
| `"merchantTokenDescription"` |  |
| `"nonce"` |  |
| `"paymentProcessor"` | The usptream payment processor for the payout. |
| `"paymentRail"` | Optional field to indicate the payment rail to use for the payout. |
| `"payrunID"` | The ID of the payrun that this payout is associated with. |
| `"payrunName"` | The name of the payrun that this payout is associated with. |
| `"rule"` |  |
| `"scheduleDate"` | The date the payout should be submitted. |
| `"scheduled"` | Should this payout be scheduled for a future date? |
| `"sourceAccountAvailableBalance"` | The available balance of the account the payout is being made from. |
| `"sourceAccountAvailableBalanceMinorUnits"` | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `"sourceAccountBic"` | The BIC of the account the payout is being made from. |
| `"sourceAccountCurrency"` | The currency of the source account. |
| `"sourceAccountIban"` | The IBAN of the account the payout is being made from. |
| `"sourceAccountIdentifier"` |  |
| `"sourceAccountName"` | The name of the account the payout is being made from. |
| `"sourceAccountNumber"` | The account number of the account the payout is being made from. |
| `"sourceAccountSortcode"` | The sort code of the account the payout is being made from. |
| `"status"` | Gets or Sets the status of payout request |
| `"tags"` | An optional list of descriptive tags attached to the payout. |
| `"theirReference"` | Gets or Sets destination reference ID |
| `"topupPayrunID"` | The ID of a payrun that needs an account top up. |
| `"transactedAmount"` | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `"transactedFxAmount"` | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `"transactedFxRate"` | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `"type"` | Gets or Sets payout type |
| `"userID"` | Gets or Sets User ID of who created the payout request |
| `"yourReference"` | Gets or Sets your reference ID |

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
| `"authorisations"` | A list of the users who have successfully authorised the latest version of the payrun and when. |
| `"authorisersCompletedCount"` | The number of distinct authorisers that have authorised the payrun. |
| `"authorisersRequiredCount"` | The number of authorisers required for this payrun. |
| `"batchPayoutID"` |  |
| `"canAuthorise"` | True if the payrun can be authorised by the user who loaded it. |
| `"canDelete"` |  |
| `"canEdit"` |  |
| `"events"` |  |
| `"hasCurrentUserAuthorised"` | True if the payrun was loaded for a user and that user has already authorised the latest version of the payrun. |
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
| `"accountID"` | The ID of the account the rule will apply to. |
| `"approveUrl"` | If set this property holds the URL an approver needs to visit in order to complete a strong authentication check in order to approve the rule. |
| `"approverID"` |  |
| `"authenticationMethods"` | A list of authentication types allowed to authorise the payout. |
| `"authorisations"` | A list of the users who have successfully authorised the latest version of the rule and when. |
| `"authorisersCompletedCount"` | The number of distinct authorisers that have authorised the rule. |
| `"authorisersRequiredCount"` | The number of authorisers required for this rule. |
| `"canAuthorise"` | True if the rule can be authorised by the user who loaded it. |
| `"createdBy"` |  |
| `"description"` | Arbitrary description for the rule. |
| `"endAt"` | Optional end time for rule executions. |
| `"hasCurrentUserAuthorised"` | True if the current user has authorised. |
| `"id"` |  |
| `"inserted"` |  |
| `"isDisabled"` | If set to true the rule will be disabled from executing. |
| `"lastExecutedAt"` |  |
| `"lastRunAtTransactionDate"` | The most recent transaction date when the rule was last run. |
| `"lastUpdated"` |  |
| `"merchantID"` | The ID of the merchant that owns the account. |
| `"name"` | A name to succinctly describe the rule. |
| `"nonce"` |  |
| `"onApprovedWebHookUrl"` | Optional URL to receive an HTTP request with the rule details when the rule status changes to approved. |
| `"onExecutionErrorWebHookUrl"` | Optional URL to receive an HTTP request when a rule execution attempt fails. |
| `"onExecutionSuccessWebHookUrl"` | Optional URL to receive an HTTP request when a rule execution attempt succeeds. |
| `"startAt"` | Optional start time for rule executions. |
| `"status"` |  |
| `"sweepAction"` |  |
| `"timeZoneId"` | If the rule should be executed on a recurring schedule this is the timezone that the CRON expression should be evaluated in. |
| `"triggerCronExpression"` | If the rule should be executed on a recurring schedule this is the expression that sets the schedule. |
| `"triggerOnPayIn"` | Set to true if the rule execution should be triggered when the account receives a pay in (credit). |
| `"userID"` |  |
| `"webHookSecret"` | If set this secret will be used to sign Web Hook requests. |

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
| `"accountID"` | The ID of the account the transaction belongs to. |
| `"accountName"` | The name of the account the transaction belongs to. |
| `"accountSequenceNumber"` | The sequence number of transaction on a per account basis. |
| `"addressDetails"` |  |
| `"amount"` | Amount of the transaction. |
| `"amountMinorUnits"` | Amount of the transaction expressed in the currency’s minor units (e.g. |
| `"balance"` | Balance left on the account after the transaction. |
| `"balanceMinorUnits"` | Balance on the account expressed in the currency’s minor units (e.g. |
| `"bookingDateTime"` |  |
| `"chargeDetails"` |  |
| `"content"` |  |
| `"counterparty"` |  |
| `"counterpartySummary"` | For pay in (credit) transactions this will contain a descriptive string with the most important fields about the counterparty. |
| `"currency"` | Currency of transaction. |
| `"currencyExchange"` | Provides details on the currency exchange. |
| `"date"` |  |
| `"description"` | Description of the transaction. |
| `"enrichment"` |  |
| `"fxAmount"` | For an FX payout this is the amound in the FX currency. |
| `"fxCurrency"` | For an FX payout this is the currency that was received or that was instructed. |
| `"fxRate"` | For an FX payout this is the exchange rate between the transaction currency and the FX currency. |
| `"grossAmount"` |  |
| `"id"` | Unique ID for the transaction. |
| `"inserted"` | Date when the transaction was inserted into the ledger. |
| `"isoBankTransactionCode"` |  |
| `"merchant"` |  |
| `"merchantID"` | The ID of the merchant that owns the account. |
| `"pageNumber"` | Current page number. |
| `"pageSize"` | Page size |
| `"payeeDetails"` | The Payee object contains details of the beneficiary, person or business. |
| `"payerDetails"` |  |
| `"paymentRequestCustomFields"` | The custom fields that were attached to the payment request that resulted in this transaction. |
| `"paymentRequestID"` | For Pay by Bank and Direct Debit transactions this will contain the ID of the payment request. |
| `"payoutID"` | ID of the payout that resulted in the transaction. |
| `"proprietaryBankTransactionCode"` |  |
| `"rawReference"` | The raw payment reference details as received from the payment processor. |
| `"reference"` |  |
| `"ruleID"` | ID of the rule that resulted in the transaction. |
| `"statementReferences"` |  |
| `"status"` |  |
| `"supplementaryData"` |  |
| `"tags"` | An optional list of descriptive tags attached to the transaction. |
| `"theirReference"` | For a pay out the reference that the payer attached for the receiving party. |
| `"totalPages"` | Total pages |
| `"totalSize"` | Total count |
| `"transactionAmount"` |  |
| `"transactionDate"` | Date when the transaction occurred. |
| `"transactionInformation"` |  |
| `"transactionMutability"` |  |
| `"type"` | Type of the transaction. |
| `"valueDateTime"` |  |
| `"virtualIBAN"` | If set it indicates the payin was to a virtual IBAN. |
| `"yourReference"` | For a pay in the reference the sending party attached. |

Operations: Create, List, Load, Remove.

API path: `/api/v1/transactions/{id}/tags`

#### User

| Field | Description |
| --- | --- |
| `"clientSessionTimeouts"` | The number of seconds a session for this user should last before expiring. |
| `"emailAddress"` |  |
| `"firstName"` |  |
| `"id"` |  |
| `"lastName"` |  |
| `"passkeyAdded"` |  |
| `"permissions"` |  |
| `"profile"` |  |
| `"rolesWithScope"` |  |
| `"twoFactorEnabled"` |  |
| `"userInviteID"` | Optional ID of the invite that is being accepted so the user can be assigned a role on a new merchant. |

Operations: List, Update.

API path: `/api/v1/user/{merchantID}/userspaged`

#### UserInvite

| Field | Description |
| --- | --- |
| `"authorisationStatus"` |  |
| `"failedUserInvites"` |  |
| `"id"` |  |
| `"initialRoleID"` | The role ID to automatically assign to the merchant’s very first user. |
| `"inviteeEmailAddress"` | Email address of the user being invited. |
| `"inviteeFirstName"` | First Name of the user being invited. |
| `"inviteeLastName"` | Last Name of the user being invited. |
| `"inviterEmailAddress"` |  |
| `"inviterFirstName"` |  |
| `"inviterLastName"` |  |
| `"isAuthorised"` | Will be set to true once the invite has met the authorisation requirements. |
| `"isInviteeRegistered"` | If true, indicates the invitee's email address corresponds to an existing MoneyMoov user. |
| `"lastInvited"` |  |
| `"merchantID"` | ID of the merchant the user is being invited to. |
| `"merchantName"` |  |
| `"message"` |  |
| `"registrationUrl"` |  |
| `"sendInviteEmail"` | If set to true an email will be sent to the invitee with instructions on how to accept the invite. |
| `"status"` |  |
| `"user"` |  |
| `"userID"` |  |
| `"userInvites"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/api/v1/userinvites/authorise/{id}`

#### Virtual

| Field | Description |
| --- | --- |
| `"accountName"` | Name for the account |
| `"accountSupplierName"` | The payment account supplier name. |
| `"availableBalance"` | The current available balance of the account. |
| `"availableBalanceMinorUnits"` | The available balance expressed in the currency’s minor units (e.g. |
| `"balance"` | Balance of the account. |
| `"balanceMinorUnits"` | Balance of the account expressed in the currency’s minor units (e.g. |
| `"bankName"` | The bank name for external accounts |
| `"consentID"` | The ID of the consent used to connect the external account. |
| `"createdBy"` |  |
| `"createdByDisplayName"` | Either the name of the user, merchant token or api key that created the account |
| `"currency"` | Currency of the account in ISO 4217 format |
| `"defaultPaymentRail"` | Indicates the default payment rail for this account. |
| `"displayName"` | Gets a unique display name for the payment account. |
| `"expiryDate"` | The date that the external account will expire |
| `"externalAccountIcon"` | The Icon for external accounts |
| `"id"` | Unique id for the account. |
| `"identifier"` |  |
| `"inserted"` | Timestamp when the account was created. |
| `"isArchived"` | Indicates whether the account is archived. |
| `"isConnectedAccount"` | Indicates if the payment account is an externally connected account. |
| `"isDefault"` | Is the default account |
| `"isTrustAccount"` | Indicates if the payment account is a trust account. |
| `"isVirtual"` | True if the account is a virtual account. |
| `"lastTransaction"` |  |
| `"lastUpdated"` | Timestamp when the account was last updated. |
| `"merchantID"` | The ID of the merchant that owns the account. |
| `"merchantName"` | The name of the merchant that owns the account. |
| `"name"` | The name of the virtual account. |
| `"physicalAccountID"` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `"rules"` | The list of rules associated with this account. |
| `"submittedPayoutsBalance"` | Total of the payouts that have been submitted for processing. |
| `"submittedPayoutsBalanceMinorUnits"` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `"summary"` | Gets a summary of the payments account's most important properties. |
| `"supplierSepaInstantStatus"` | Indicates the status of the SEPA Instant payment rail for this account. |
| `"xeroBankFeedConnectionStatus"` | States the status of the Xero bank feed connection, if applicable. |
| `"xeroBankFeedLastSyncedAt"` |  |
| `"xeroBankFeedSyncLastFailedAt"` |  |
| `"xeroBankFeedSyncLastFailureReason"` |  |
| `"xeroBankFeedSyncStatus"` |  |
| `"xeroUnsynchronisedTransactionsCount"` | Indicates the number of unsynchronised transactions with Xero |

Operations: Create, Update.

API path: `/api/v1/accounts/{accountID}/virtual`

#### Webhook

| Field | Description |
| --- | --- |
| `"destinationUrl"` | The destination URL for the webhook. |
| `"emailAddress"` | The recipient email address(es) for notifications. |
| `"failedNotificationEmailAddress"` | The email address to which notifications about failed webhook deliveries will be sent. |
| `"id"` |  |
| `"isActive"` |  |
| `"merchantID"` | The ID of the merchant that the webhook is for. |
| `"notificationMethod"` | The type of notification that will be sent. |
| `"resourceTypes"` | The resource types that the webhook will be generated for. |
| `"retry"` |  |
| `"secret"` | The secret key required to authenticate webhook notifications. |
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
| `accountBalances` | `[]any` | The various balances for the account. |
| `accountID` | `string` | ID of the account. |
| `accountIdentifications` | `[]any` | The canoncial identifiers for the account. |
| `accountName` | `string` | Name for the account |
| `accountNames` | `[]any` | Optional account names set by the account holder. |
| `accountSupplierName` | `string` | The payment account supplier name. |
| `accountType` | `string` | The type of account e.g. |
| `availableBalance` | `float64` | The current available balance of the account. |
| `availableBalanceMinorUnits` | `int` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `float64` | Balance of the account. |
| `balanceMinorUnits` | `int` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `string` | The bank name for external accounts |
| `consentID` | `string` | The ID of the consent used to connect the external account. |
| `consolidatedAccountInformation` | `map[string]any` | Summary information regarding account balances of the overall account provided by the bank. |
| `createdBy` | `map[string]any` |  |
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
| `identifier` | `map[string]any` |  |
| `inserted` | `string` | Timestamp when the account was created. |
| `isArchived` | `bool` | Indicates whether the account is archived. |
| `isConnectedAccount` | `bool` | Indicates if the payment account is an externally connected account. |
| `isDefault` | `bool` | Is the default account |
| `isTrustAccount` | `bool` | Indicates if the payment account is a trust account. |
| `isVirtual` | `bool` | True if the account is a virtual account. |
| `lastTransaction` | `map[string]any` |  |
| `lastUpdated` | `string` | Timestamp when the account was last updated. |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `merchantName` | `string` | The name of the merchant that owns the account. |
| `nickname` | `string` | Nickname of the account that was provided by the account owner. |
| `physicalAccountID` | `string` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `roleIDs` | `[]any` | Optional list of role IDs that will get access to the payment account when created. |
| `rules` | `[]any` | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `float64` | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `int` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
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
| `xeroUnsynchronisedTransactionsCount` | `int` | Indicates the number of unsynchronised transactions with Xero |

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
    "account_id": "example_account_id",
    "currency": "example_currency",
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
| `approveUrl` | `string` | This field is used when returning a batch payout record to a client. |
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
| `authenticationMethods` | `[]any` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `[]any` | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `int` | The number of distinct authorisers that have authorised the beneficiary. |
| `authorisersRequiredCount` | `int` | The number of authorisers required for this beneficiary. |
| `beneficiaries` | `[]any` |  |
| `beneficiaryEvents` | `[]any` |  |
| `canAuthorise` | `bool` | True if the beneficiary can be authorised by the user who loaded it. |
| `canUpdate` | `bool` | True if the beneficiary can be updated by the user who loaded it. |
| `createdBy` | `map[string]any` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` | Gets or Sets the currency. |
| `destination` | `map[string]any` |  |
| `failedBeneficiaries` | `map[string]any` |  |
| `hasCurrentUserAuthorised` | `bool` | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isEnabled` | `bool` |  |
| `lastAuthorised` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` | Gets or Sets the merchant id. |
| `name` | `string` | The descriptive name for the beneficiary. |
| `nonce` | `string` |  |
| `sourceAccountIDs` | `[]any` | ID of the accounts which are authorised to act as a source for the beneficiary. |
| `sourceAccounts` | `[]any` |  |
| `theirReference` | `string` | The reference that will be used by default as TheirReference when creating payouts to this beneficiary if no TheirReference is specified for the payout. |

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
    "id": "example_id",
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
| `groupMembers` | `[]any` | The existing group members. |
| `groupName` | `string` | The descriptive name for the beneficiary group. |
| `id` | `string` |  |
| `inserted` | `string` | Timestamp indicating when the group was created. |
| `lastUpdated` | `string` | Timestamp indicating when the group was last updated. |
| `merchantID` | `string` | Gets or Sets the merchant id. |

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
| `isPayerAuthenticationRequired` | `bool` | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `bool` | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `string` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `string` | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `string` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `int` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `int` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `string` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `string` |  |
| `requestID` | `string` |  |
| `responseCode` | `string` |  |
| `responseType` | `string` |  |
| `status` | `string` |  |
| `threeDSRedirectUrl` | `string` | Checkout.com require a redirect for 3DS authentication. |
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
| `isPayerAuthenticationRequired` | `bool` | Gets set to true if 3-D Secure payer authentication is required for a payment. |
| `isSoftDecline` | `bool` | Gets set to true if the card processor flagged the transaction as having failed address or card security number verification. |
| `payerAuthenticationAccessToken` | `string` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect. |
| `payerAuthenticationMerchantData` | `string` | If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the "authenticationcallback" method that gets called automatically after a successful payer authenticati… |
| `payerAuthenticationUrl` | `string` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank. |
| `payerAuthenticationWindowHeight` | `int` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge. |
| `payerAuthenticationWindowWidth` | `int` | If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge. |
| `paymentRequestCallbackUrl` | `string` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `string` |  |
| `requestID` | `string` |  |
| `responseCode` | `string` |  |
| `responseType` | `string` |  |
| `status` | `string` |  |
| `threeDSRedirectUrl` | `string` | Checkout.com require a redirect for 3DS authentication. |
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
| `authorisationUrl` | `string` | The URL the authorising user needs to be redirected to in order to get the open banking consent token. |
| `callbackUrl` | `string` | Optional callback URL that the end user performing the open banking authorisation will be redirected to on completion. |
| `consentID` | `string` | The ID of the open banking consent. |
| `emailAddress` | `string` | The email address that identifies the end user that will be authorising the open banking consent request. |
| `expiryDate` | `string` |  |
| `failureCallbackUrl` | `string` | Optional callback URL for open banking consent authorisation failure. |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `institutionID` | `string` | The institution ID the open banking consent is being requested for. |
| `isConnectedAccounts` | `bool` | Optional setting. |
| `isEnabled` | `bool` |  |
| `merchantID` | `string` | The ID of the merchant the consent token is being created to be used with. |
| `provider` | `string` | Lists the supported card and PIS processors. |
| `successWebHookUrl` | `string` | A web hook URL to send an HTTP request to when the open banking consent is successfuly authorised. |

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
| `failedSubmissions` | `map[string]any` | Dictionary of failed submissions, keyed by the index (1-based) in the original request. |
| `successfulSubmissions` | `[]any` | List of successfully submitted direct debit payments. |

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
| `exchangeRate` | `float64` | The price at which the transaction will buy the source currency using the destination currency. |
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
| `isRecurring` | `bool` | Whether this mandate is single-use or recurring. |
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
| `accountCurrencies` | `[]any` | The list of currencies that the merchant has accounts for. |
| `canHaveTrustAccounts` | `bool` | Trust accounts are a special type of account that allow the account name to be trusted for use in statements and verification of payee checks. |
| `cardPaymentProcessor` | `string` | Name of the card payment processor. |
| `companyID` | `string` | The Company ID recorded in the Compliance system. |
| `displayQrOnHostedPay` | `bool` | Indicates if a QR Code containing the payment link should be displayed on the hosted payment page. |
| `hostedPayVersion` | `int` | The version of the hosted payment page to use with the merchant. |
| `id` | `string` | Unique ID for the merchant. |
| `inserted` | `string` | Timestamp the merchant was added to MoneyMoov. |
| `isBlocked` | `bool` | The merchant is blocked from making payments (payouts). |
| `isExited` | `bool` | The merchant has formally terminated their relationship and is no longer a customer. |
| `isSuspended` | `bool` | The merchant has temporarily suspended their own account. |
| `jurisdiction` | `string` | The jurisdiction the merchant entity is incorporated or established in. |
| `logoUrlPng` | `string` | The CDN URL of the merchant's logo in PNG format. |
| `logoUrlSvg` | `string` | The CDN URL of the merchant's logo in SVG format. |
| `merchantCategoryCode` | `string` | The industry code that represents the merchant's primary trading activity. |
| `name` | `string` | The registered business name of the merchant. |
| `notes` | `string` | The notes field is an optional free text field that can be used to store any additional information about the merchant. |
| `parentMerchant` | `map[string]any` |  |
| `paymentAccountLimit` | `int` | The maximum number of payment accounts that can be created for the Merchant. |
| `paymentAccounts` | `[]any` |  |
| `reason` | `string` | The reason for the suspension. |
| `shortName` | `string` | A URL friendly shortish name for the merchant. |
| `supportedPaymentMethodsList` | `[]any` | The payment methods that are configured and supported for this merchant. |
| `suspensionReason` | `string` | The reason for the suspension, provided by the merchant. |
| `tags` | `[]any` | An optional list of descriptive tags that can be used on merchant entities such as payment requests. |
| `timeZoneId` | `string` | The IANA (Internet Assigned Numbers Authority) time zone identifier of the merchant. |
| `tradingName` | `string` | An optional trading name. |
| `webHookLimit` | `int` | The maximum number of web hooks that can be created for the Merchant. |
| `yourRoleName` | `string` | The name of the role for the identity that loaded the merchant record. |

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
| `isRecurring` | `bool` | Whether this mandate is single-use or recurring. |
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
| `bankCountryCodes` | `[]any` | The list of country codes representing the banks the country supports. |
| `bankID` | `string` | ID of the bank to be configured for the merchant. |
| `bankName` | `string` | Name of the Bank/Institution. |
| `businessInstitutionID` | `string` | ID that the processor uses to identify the bank (business accounts). |
| `currency` | `string` | Currency supported by the bank. |
| `logo` | `string` | URL of the bank's logo. |
| `message` | `string` | Message relating to specific bank. |
| `messageImageUrl` | `string` | Optional image URL to be displayed with the message. |
| `order` | `int` | Order in which this setting will appear in the UI. |
| `personalInstitutionID` | `string` | ID that the processor uses to identify the bank (personal accounts). |
| `processor` | `string` | Name of the bank payment processor. |
| `warningHeading` | `string` | The heading for a warning message related to the bank institution to be displayed to the user. |
| `warningMessage` | `string` | The warning message related to the bank institution to be displayed to the user. |

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
| `customFields` | `[]any` | A list of custom fields that can be included in the payment request template. |
| `defaultFields` | `[]any` | A list of default fields that are included in the payment request template. |
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
| `authenticationMethods` | `[]any` | A list of authentication types allowed to authorise the merchant token. |
| `authorisations` | `[]any` | A list of users who have successfully authorised the latest version of the beneficiary. |
| `authorisersCompletedCount` | `int` | The number of distinct authorisers that have authorised the merchant token. |
| `authorisersRequiredCount` | `int` | The number of authorisers required for this merchant token. |
| `canAuthorise` | `bool` | True if the merchant token can be authorised by the user who loaded it. |
| `description` | `string` | Token description |
| `expiresAt` | `string` | Optional. |
| `hasCurrentUserAuthorised` | `bool` | True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary. |
| `hmacAlgorithm` | `string` | Optional shared secret algorithm to use for HMAC authentication. |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `ipAddressWhitelist` | `string` | Optional. |
| `isArchived` | `bool` | Indicates whether the merchant token is archived. |
| `isEnabled` | `bool` | If set to false the merchant token will not be accepted to authorise a request. |
| `lastAuthorised` | `string` |  |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` | The merchant id to add to the token |
| `nonce` | `string` |  |
| `permissionTypes` | `[]any` | The permissions that the merchant token supports. |
| `requestSignatureVersion` | `int` | Represent the version of the overall merchant token. |
| `sharedSecretAlgorithm` | `string` | Optional shared secret algorithm to use for HMAC authentication. |
| `sharedSecretBase64` | `string` | The base 64 encoded shared secret that is used for request authentication with an HMAC. |
| `token` | `string` | The JWT merchant token. |

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
| `accountName` | `string` | The name of the account to verify |
| `accountNumber` | `string` | The account number of the account to verify (for CoP checks) |
| `iban` | `string` | The IBAN of the account to verify (for VoP checks) |
| `payeeVerifiedAccountName` | `string` | The verified account name of the payee, if available (in case of a close match) |
| `result` | `string` | The result of the payee verification |
| `secondaryIdentification` | `string` | Optional secondary identifier for the account to verify. |
| `sortCode` | `string` | The sort code of the account to verify (for CoP checks) |

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
| `amount` | `float64` | The amount of money to request. |
| `amountPending` | `float64` | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `float64` | Total amount received for this payment request. |
| `amountRefunded` | `float64` | Total amount refunded for this payment request. |
| `autoSendReceipt` | `bool` | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `string` | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `string` | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `bool` | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `bool` | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `string` | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `bool` | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardNoPayerAuthentication` | `bool` | If set to true for card payments no attempt will be made to use payer authentication (3-D Secure and equivalent). |
| `cardProcessorMerchantID` | `string` | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `string` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `string` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `cardTransmitRawDetails` | `bool` | If set to true for card payments the sensitive card number and card verification number will be transmitted directly rather than being tokenised. |
| `createdByUser` | `map[string]any` |  |
| `currency` | `string` | The currency of the request. |
| `customFields` | `[]any` | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `string` | Optional email address for the customer. |
| `customerID` | `string` | An optional customer identifier for the payment request. |
| `customerName` | `string` |  |
| `description` | `string` | An optional description for the payment request. |
| `destinationAccount` | `map[string]any` |  |
| `directDebitPayment` | `map[string]any` | Contains information about a Direct Debit payment attempt for a payment request. |
| `dueDate` | `string` | The due date for the payment request. |
| `events` | `[]any` |  |
| `failureCallbackUrl` | `string` | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `[]any` | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `string` |  |
| `hostedPayCheckoutUrl` | `string` | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `string` |  |
| `ignoreAddressVerification` | `bool` | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `string` | The timestamp the payment request was created at. |
| `insertedSortable` | `string` | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `isArchived` | `bool` | Indicates whether the payment request is archived. |
| `jwk` | `string` | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `string` | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `string` | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `string` | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `string` | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `string` | The ID of the merchant to create the payment request for. |
| `merchantTokenDescription` | `string` | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `string` |  |
| `notificationRoleIDs` | `[]any` | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `string` | An optional order ID for the payment request. |
| `partialPaymentMethod` | `string` | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `string` | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `[]any` | The payment attempts made against this payment request. |
| `paymentMethods` | `[]any` | The payment methods that the payment request supports. |
| `paymentProcessor` | `string` | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `payrunID` | `string` | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `string` | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `string` | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `map[string]any` |  |
| `sandboxSettleDelayInSeconds` | `int` | Sandbox only. |
| `shippingAddress` | `map[string]any` |  |
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
| `tagIds` | `[]any` | An optional list of tag ids to add to the payment request |
| `tags` | `[]any` | An optional list of descriptive tags attached to the payment request. |
| `title` | `string` | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `[]any` |  |
| `transactions` | `[]any` |  |
| `useHostedPaymentPage` | `bool` | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

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
| `accountName` | `string` | Name for the account |
| `accountSupplierName` | `string` | The payment account supplier name. |
| `availableBalance` | `float64` | The current available balance of the account. |
| `availableBalanceMinorUnits` | `int` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `float64` | Balance of the account. |
| `balanceMinorUnits` | `int` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `string` | The bank name for external accounts |
| `consentID` | `string` | The ID of the consent used to connect the external account. |
| `createdBy` | `map[string]any` |  |
| `createdByDisplayName` | `string` | Either the name of the user, merchant token or api key that created the account |
| `currency` | `string` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `string` | Indicates the default payment rail for this account. |
| `displayName` | `string` | Gets a unique display name for the payment account. |
| `expiryDate` | `string` | The date that the external account will expire |
| `externalAccountIcon` | `string` | The Icon for external accounts |
| `id` | `string` | Unique id for the account. |
| `identifier` | `map[string]any` |  |
| `inserted` | `string` | Timestamp when the account was created. |
| `isArchived` | `bool` | Indicates whether the account is archived. |
| `isConnectedAccount` | `bool` | Indicates if the payment account is an externally connected account. |
| `isDefault` | `bool` | Is the default account |
| `isTrustAccount` | `bool` | Indicates if the payment account is a trust account. |
| `isVirtual` | `bool` | True if the account is a virtual account. |
| `lastTransaction` | `map[string]any` |  |
| `lastUpdated` | `string` | Timestamp when the account was last updated. |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `merchantName` | `string` | The name of the merchant that owns the account. |
| `physicalAccountID` | `string` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `[]any` | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `float64` | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `int` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `string` | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `string` | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `string` | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `string` |  |
| `xeroBankFeedSyncLastFailedAt` | `string` |  |
| `xeroBankFeedSyncLastFailureReason` | `string` |  |
| `xeroBankFeedSyncStatus` | `string` |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | Indicates the number of unsynchronised transactions with Xero |

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
| `accountName` | `string` | Name for the account |
| `availableBalance` | `float64` | The current available balance of the account. |
| `balance` | `float64` | Balance of the account. |
| `balanceMinorUnits` | `int` | Balance of the account expressed in the currency’s minor units (e.g. |
| `currency` | `string` | Currency of the account in ISO 4217 format |
| `id` | `string` | Unique id for the account. |
| `identifier` | `map[string]any` |  |
| `isArchived` | `bool` | Is the account archived |
| `isConnectedAccount` | `bool` | Indicates if the payment account is an externally connected account. |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `submittedPayoutsBalance` | `float64` | Total of the payouts that have been submitted for processing. |

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
| `paymentInitiationID` | `string` | The unique identifier of the payment initiation request. |
| `paymentRequestCallbackUrl` | `string` | The callback URL that was set when the payment request was created. |
| `paymentRequestID` | `string` |  |
| `redirectUrl` | `string` | A redirect URL for the user to authorise the payment initiation request at the ASPSP |
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
| `amount` | `float64` | The amount of money to request. |
| `amountPending` | `float64` | Total amount that has been authorised but not settled for this payment request. |
| `amountReceived` | `float64` | Total amount received for this payment request. |
| `amountRefunded` | `float64` | Total amount refunded for this payment request. |
| `autoSendReceipt` | `bool` | If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received. |
| `baseOriginUrl` | `string` | For card payments the origin of the payment page needs to be set in advance. |
| `callbackUrl` | `string` | Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL. |
| `cardAuthorizeOnly` | `bool` | For card payments the default behaviour is to authorise and capture the payment at the same time. |
| `cardCreateToken` | `bool` | For card payments a payment attempt can be used to create a reusable token for subsequent payments. |
| `cardCreateTokenMode` | `string` | This specifies whether user consent will be taken before tokenising card or not. |
| `cardIgnoreCVN` | `bool` | If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails. |
| `cardProcessorMerchantID` | `string` | Optional field that if specified indicates the processor merchant ID that should be used to process any card payments. |
| `cardStripePaymentIntentID` | `string` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID. |
| `cardStripePaymentIntentSecret` | `string` | If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret. |
| `createdByUser` | `map[string]any` |  |
| `currency` | `string` | The currency of the request. |
| `customFields` | `[]any` | A list of custom fields attached to the payment request. |
| `customerEmailAddress` | `string` | Optional email address for the customer. |
| `customerID` | `string` | An optional customer identifier for the payment request. |
| `customerName` | `string` |  |
| `description` | `string` | An optional description for the payment request. |
| `destinationAccount` | `map[string]any` |  |
| `directDebitPayment` | `map[string]any` | Contains information about a Direct Debit payment attempt for a payment request. |
| `doSimulateSettlementFailure` | `bool` |  |
| `dueDate` | `string` | The due date for the payment request. |
| `errorDescription` | `string` |  |
| `events` | `[]any` |  |
| `failedPaymentRequests` | `map[string]any` |  |
| `failureCallbackUrl` | `string` | Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page. |
| `fieldDisplaySettings` | `[]any` | A list of field display settings that control which fields are displayed to the payer. |
| `formattedAmount` | `string` |  |
| `hostedPayCheckoutUrl` | `string` | This is a convenience link generated for payment requests whose merchants are using hosted payment pages. |
| `id` | `string` |  |
| `ignoreAddressVerification` | `bool` | If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails. |
| `inserted` | `string` | The timestamp the payment request was created at. |
| `insertedSortable` | `string` | The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se… |
| `institution` | `string` |  |
| `isArchived` | `bool` | Indicates whether the payment request is archived. |
| `jwk` | `string` | The jwk containing the public key used to verify the signature of the payment request. |
| `lastUpdated` | `string` | The timestamp the payment request was last updated at. |
| `lightningInvoice` | `string` | Bitcoin Lightning invoice for the payment request. |
| `lightningInvoiceExpiresAt` | `string` | Date and time of expiration of the lightning invoice. |
| `merchantDirectDebitMandateID` | `string` | Optional ID of the direct debit mandate associated with this payment request. |
| `merchantID` | `string` |  |
| `merchantTokenDescription` | `string` | Description of the merchant token in case the Payment request was created using a merchant token. |
| `notificationEmailAddresses` | `string` |  |
| `notificationRoleIDs` | `[]any` | A list of roles whose members will receive notifications about this payment request. |
| `orderID` | `string` | An optional order ID for the payment request. |
| `partialPaymentMethod` | `string` | The approach to use, or not, for accepting partial payments. |
| `partialPaymentSteps` | `string` | An optional comma separated list of partial payment amounts. |
| `paymentAttempts` | `[]any` | The payment attempts made against this payment request. |
| `paymentInitiationID` | `string` |  |
| `paymentMethods` | `[]any` | The payment methods that the payment request supports. |
| `paymentProcessor` | `string` | If the card payment option is enabled this field indicates which card processor the merchant is set up to use. |
| `paymentRequests` | `[]any` |  |
| `payrunID` | `string` | The ID of a payrun that needs an account top up. |
| `pispAccountID` | `string` | The payment account ID to use to receive payment initiation payments. |
| `priorityBankID` | `string` | The ID of the bank that is set as the priority bank for display on pay element. |
| `result` | `map[string]any` |  |
| `sandboxSettleDelayInSeconds` | `int` | Sandbox only. |
| `shippingAddress` | `map[string]any` |  |
| `status` | `string` | The current status of the payment request. |
| `successWebHookUrl` | `string` | If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked. |
| `tags` | `[]any` | An optional list of descriptive tags attached to the payment request. |
| `title` | `string` | A generic field to contain any additional data that the merchant wishes to store against the payment request. |
| `tokenisedCards` | `[]any` |  |
| `transactions` | `[]any` |  |
| `useHostedPaymentPage` | `bool` | If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page. |

#### Example: Load

```go
paymentRequest, err := client.PaymentRequest(nil).Load(map[string]any{"id": "payment_request_id"}, nil)
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
| `applePayTransactionID` | `string` | Transaction ID received in Apple pay token. |
| `cardAuthorizationResponseID` | `string` | For a successful card authorization this field will hold the response ID. |
| `cardExpiryMonth` | `int` | For card payment events this field holds the payer's card expiry month. |
| `cardExpiryYear` | `int` | For card payment events this field holds the payer's card expiry year. |
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
| `amount` | `float64` | The amount of money to request. |
| `amountPending` | `float64` | The amount of money that was authorised but has not arrived in the account yet. |
| `amountReceived` | `float64` | The amount of money that has been received for this payment request. |
| `amountRefunded` | `float64` | The amount of money that has been refunded for this payment request. |
| `callbackUrl` | `string` |  |
| `cardStripePaymentIntentSecret` | `string` |  |
| `countryCode` | `string` | The country code associated with the payment. |
| `currency` | `string` | The currency of the request. |
| `customFieldsToDisplay` | `[]any` | Custom fields to display to the customer. |
| `description` | `string` | An optional description for the payment request. |
| `dueDate` | `string` | The due date of the payment request. |
| `fieldDisplaySettings` | `[]any` |  |
| `googlePayMerchantID` | `string` | Merchant ID from Google Pay |
| `id` | `string` |  |
| `jwk` | `string` | The jwk containing the public key |
| `merchantID` | `string` |  |
| `merchantLogoUrlPng` | `string` |  |
| `merchantLogoUrlSvg` | `string` |  |
| `merchantName` | `string` |  |
| `merchantShortName` | `string` |  |
| `partialPaymentMethod` | `string` |  |
| `paymentAttempts` | `[]any` | The payment attempts for this payment request. |
| `paymentMethodsList` | `[]any` | The payment methods that the payment request supports. |
| `paymentProcessor` | `string` | The card processor |
| `paymentProcessorKey` | `string` | The card processors public key |
| `pispError` | `string` | This is the error returned from the bank which is recorded in payment request events. |
| `priorityBankID` | `string` |  |
| `status` | `string` | The status of the payment request. |
| `stripeAccountID` | `string` | Account ID of connected customers in Stripe |
| `title` | `string` | The title of the payment request. |

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
| `amount` | `float64` | The authorised payment amount. |
| `amountPending` | `float64` |  |
| `amountReceived` | `float64` |  |
| `amountRefunded` | `float64` |  |
| `currency` | `string` | The authorised payment currency. |
| `customerID` | `string` | The customer id |
| `paymentRequestID` | `string` | The ID of the payment request the result is for. |
| `payments` | `[]any` | The list of payment attempts that have been received for the payment request. |
| `pispAuthorizations` | `[]any` |  |
| `requestedAmount` | `float64` | The full original payment amount requested. |
| `result` | `string` | The result of the payment attempt. |

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
| `accountID` | `string` | Gets or Sets Account Id of sending account |
| `allowIncomplete` | `bool` | If set to true the payout will get created even if the business validation rules fail. |
| `amount` | `float64` | Gets or Sets payout amount |
| `amountMinorUnits` | `int` | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `string` | This field is used when returning an payout record to a client. |
| `approverID` | `string` | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `[]any` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `[]any` | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `int` | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `int` | The number of authorisers required for this payout. |
| `batchPayoutID` | `string` | The ID of the batch the payout is associated with. |
| `beneficiary` | `map[string]any` |  |
| `beneficiaryID` | `string` | Optional. |
| `canAuthorise` | `bool` | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `bool` | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `bool` | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `string` | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `string` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` | Gets or Sets Currency of payout request |
| `currentUserID` | `string` | The ID of the user that requested access to the PayOut record. |
| `description` | `string` | Gets or Sets description of payout request |
| `destination` | `map[string]any` |  |
| `documents` | `[]any` | Documents associated with the payout. |
| `events` | `[]any` | The activity associated with the payout. |
| `failedPayouts` | `map[string]any` |  |
| `formattedAmount` | `string` | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `string` | FX destination currency and amount formatted string. |
| `formattedSchedule` | `string` |  |
| `formattedScheduleDayOnly` | `string` |  |
| `formattedSourceAccountAvailableBalance` | `string` | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `float64` | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `int` | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `string` | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `string` | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `string` | Optional. |
| `fxRate` | `float64` | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `bool` | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `bool` | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `string` | The ID for the payout. |
| `inserted` | `string` |  |
| `invoiceID` | `string` | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `bool` | Indicates whether the payout is archived. |
| `isFailed` | `bool` | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `bool` | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `bool` | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `string` |  |
| `nonce` | `string` |  |
| `paymentProcessor` | `string` | The usptream payment processor for the payout. |
| `paymentRail` | `string` | Optional field to indicate the payment rail to use for the payout. |
| `payouts` | `[]any` |  |
| `payrunID` | `string` | The ID of the payrun that this payout is associated with. |
| `payrunName` | `string` | The name of the payrun that this payout is associated with. |
| `reason` | `string` |  |
| `rule` | `map[string]any` |  |
| `scheduleDate` | `string` | The date the payout should be submitted. |
| `scheduled` | `bool` | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `float64` | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `int` | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `string` | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `string` | The currency of the source account. |
| `sourceAccountIban` | `string` | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `map[string]any` |  |
| `sourceAccountName` | `string` | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `string` | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `string` | The sort code of the account the payout is being made from. |
| `status` | `string` | Gets or Sets the status of payout request |
| `tagIds` | `[]any` | An optional list of tag ids to add to the payout. |
| `tags` | `[]any` | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `string` | Gets or Sets destination reference ID |
| `topupPayrunID` | `string` | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `float64` | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `float64` | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `float64` | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `string` | Gets or Sets payout type |
| `userID` | `string` | Gets or Sets User ID of who created the payout request |
| `yourReference` | `string` | Gets or Sets your reference ID |

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
    "id": "example_id",
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
| `accountID` | `string` | Gets or Sets Account Id of sending account |
| `amount` | `float64` | Gets or Sets payout amount |
| `amountMinorUnits` | `int` | The payout amount expressed in the currency’s minor units (e.g. |
| `approvePayoutUrl` | `string` | This field is used when returning an payout record to a client. |
| `approverID` | `string` | Gets the User ID of person that approved the payout. |
| `authenticationMethods` | `[]any` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `[]any` | A list of the users who have successfully authorised the latest version of the payout and when. |
| `authorisersCompletedCount` | `int` | The number of distinct authorisers that have authorised the payout. |
| `authorisersRequiredCount` | `int` | The number of authorisers required for this payout. |
| `batchPayoutID` | `string` | The ID of the batch the payout is associated with. |
| `beneficiary` | `map[string]any` |  |
| `canAuthorise` | `bool` | True if the payout can be authorised by the user who loaded it. |
| `canProcess` | `bool` | If set to true indicates the payout has been flagged as safe to process after transaction monitoring. |
| `canUpdate` | `bool` | True if the payout can be updated by the user who loaded it. |
| `chargeBearer` | `string` | Optional field to set who should pay any fees for the payout. |
| `createdBy` | `string` |  |
| `createdByEmailAddress` | `string` |  |
| `currency` | `string` | Gets or Sets Currency of payout request |
| `currentUserID` | `string` | The ID of the user that requested access to the PayOut record. |
| `description` | `string` | Gets or Sets description of payout request |
| `destination` | `map[string]any` |  |
| `documents` | `[]any` | Documents associated with the payout. |
| `events` | `[]any` | The activity associated with the payout. |
| `formattedAmount` | `string` | Currency and formatted amount string. |
| `formattedFxDestinationAmount` | `string` | FX destination currency and amount formatted string. |
| `formattedSchedule` | `string` |  |
| `formattedScheduleDayOnly` | `string` |  |
| `formattedSourceAccountAvailableBalance` | `string` | The available balance of the account the payout is being made from. |
| `fxDestinationAmount` | `float64` | If specified this will be the amount sent to the payee. |
| `fxDestinationAmountMinorUnits` | `int` | The payout FxDestinationAmount expressed in the currency’s minor units (e.g. |
| `fxDestinationCurrency` | `string` | For an FX payout this is the currency to send to the beneficiary. |
| `fxQuoteExpiresAt` | `string` | If an FX held rate quote ID is being used this is the time the quote expires. |
| `fxQuoteID` | `string` | Optional. |
| `fxRate` | `float64` | For an FX payout this is the exchange rate to use for the payout. |
| `fxUseDestinationAmount` | `bool` | For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated. |
| `hasCurrentUserAuthorised` | `bool` | True if the payout was loaded for a user and that user has already authorised the latest version of the payout. |
| `id` | `string` | The ID for the payout. |
| `inserted` | `string` |  |
| `invoiceID` | `string` | Optional field to associate the payout with the invoice from an external application such as Xero. |
| `isArchived` | `bool` | Indicates whether the payout is archived. |
| `isFailed` | `bool` | Set to true if a submitted payout subsequently fails. |
| `isSettled` | `bool` | Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger. |
| `isSubmitted` | `bool` | Indicates whether the payout has been submitted for processing. |
| `lastUpdated` | `string` |  |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `merchantTokenDescription` | `string` |  |
| `nonce` | `string` |  |
| `paymentProcessor` | `string` | The usptream payment processor for the payout. |
| `paymentRail` | `string` | Optional field to indicate the payment rail to use for the payout. |
| `payrunID` | `string` | The ID of the payrun that this payout is associated with. |
| `payrunName` | `string` | The name of the payrun that this payout is associated with. |
| `rule` | `map[string]any` |  |
| `scheduleDate` | `string` | The date the payout should be submitted. |
| `scheduled` | `bool` | Should this payout be scheduled for a future date? |
| `sourceAccountAvailableBalance` | `float64` | The available balance of the account the payout is being made from. |
| `sourceAccountAvailableBalanceMinorUnits` | `int` | The available balance of the source account expressed in the currency’s minor units (e.g. |
| `sourceAccountBic` | `string` | The BIC of the account the payout is being made from. |
| `sourceAccountCurrency` | `string` | The currency of the source account. |
| `sourceAccountIban` | `string` | The IBAN of the account the payout is being made from. |
| `sourceAccountIdentifier` | `map[string]any` |  |
| `sourceAccountName` | `string` | The name of the account the payout is being made from. |
| `sourceAccountNumber` | `string` | The account number of the account the payout is being made from. |
| `sourceAccountSortcode` | `string` | The sort code of the account the payout is being made from. |
| `status` | `string` | Gets or Sets the status of payout request |
| `tags` | `[]any` | An optional list of descriptive tags attached to the payout. |
| `theirReference` | `string` | Gets or Sets destination reference ID |
| `topupPayrunID` | `string` | The ID of a payrun that needs an account top up. |
| `transactedAmount` | `float64` | The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction. |
| `transactedFxAmount` | `float64` | The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction. |
| `transactedFxRate` | `float64` | The actual FX rate applied during settlement, as recorded on the associated transaction. |
| `type` | `string` | Gets or Sets payout type |
| `userID` | `string` | Gets or Sets User ID of who created the payout request |
| `yourReference` | `string` | Gets or Sets your reference ID |

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
| `authorisations` | `[]any` | A list of the users who have successfully authorised the latest version of the payrun and when. |
| `authorisersCompletedCount` | `int` | The number of distinct authorisers that have authorised the payrun. |
| `authorisersRequiredCount` | `int` | The number of authorisers required for this payrun. |
| `batchPayoutID` | `string` |  |
| `canAuthorise` | `bool` | True if the payrun can be authorised by the user who loaded it. |
| `canDelete` | `bool` |  |
| `canEdit` | `bool` |  |
| `events` | `[]any` |  |
| `hasCurrentUserAuthorised` | `bool` | True if the payrun was loaded for a user and that user has already authorised the latest version of the payrun. |
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
| `accountID` | `string` | The ID of the account the rule will apply to. |
| `approveUrl` | `string` | If set this property holds the URL an approver needs to visit in order to complete a strong authentication check in order to approve the rule. |
| `approverID` | `string` |  |
| `authenticationMethods` | `[]any` | A list of authentication types allowed to authorise the payout. |
| `authorisations` | `[]any` | A list of the users who have successfully authorised the latest version of the rule and when. |
| `authorisersCompletedCount` | `int` | The number of distinct authorisers that have authorised the rule. |
| `authorisersRequiredCount` | `int` | The number of authorisers required for this rule. |
| `canAuthorise` | `bool` | True if the rule can be authorised by the user who loaded it. |
| `createdBy` | `map[string]any` |  |
| `description` | `string` | Arbitrary description for the rule. |
| `endAt` | `string` | Optional end time for rule executions. |
| `hasCurrentUserAuthorised` | `bool` | True if the current user has authorised. |
| `id` | `string` |  |
| `inserted` | `string` |  |
| `isDisabled` | `bool` | If set to true the rule will be disabled from executing. |
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
| `sweepAction` | `map[string]any` |  |
| `timeZoneId` | `string` | If the rule should be executed on a recurring schedule this is the timezone that the CRON expression should be evaluated in. |
| `triggerCronExpression` | `string` | If the rule should be executed on a recurring schedule this is the expression that sets the schedule. |
| `triggerOnPayIn` | `bool` | Set to true if the rule execution should be triggered when the account receives a pay in (credit). |
| `userID` | `string` |  |
| `webHookSecret` | `string` | If set this secret will be used to sign Web Hook requests. |

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
| `accountID` | `string` | The ID of the account the transaction belongs to. |
| `accountName` | `string` | The name of the account the transaction belongs to. |
| `accountSequenceNumber` | `int` | The sequence number of transaction on a per account basis. |
| `addressDetails` | `map[string]any` |  |
| `amount` | `float64` | Amount of the transaction. |
| `amountMinorUnits` | `int` | Amount of the transaction expressed in the currency’s minor units (e.g. |
| `balance` | `float64` | Balance left on the account after the transaction. |
| `balanceMinorUnits` | `int` | Balance on the account expressed in the currency’s minor units (e.g. |
| `bookingDateTime` | `string` |  |
| `chargeDetails` | `map[string]any` |  |
| `content` | `[]any` |  |
| `counterparty` | `map[string]any` |  |
| `counterpartySummary` | `string` | For pay in (credit) transactions this will contain a descriptive string with the most important fields about the counterparty. |
| `currency` | `string` | Currency of transaction. |
| `currencyExchange` | `map[string]any` | Provides details on the currency exchange. |
| `date` | `string` |  |
| `description` | `string` | Description of the transaction. |
| `enrichment` | `map[string]any` |  |
| `fxAmount` | `float64` | For an FX payout this is the amound in the FX currency. |
| `fxCurrency` | `string` | For an FX payout this is the currency that was received or that was instructed. |
| `fxRate` | `float64` | For an FX payout this is the exchange rate between the transaction currency and the FX currency. |
| `grossAmount` | `map[string]any` |  |
| `id` | `string` | Unique ID for the transaction. |
| `inserted` | `string` | Date when the transaction was inserted into the ledger. |
| `isoBankTransactionCode` | `map[string]any` |  |
| `merchant` | `map[string]any` |  |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `pageNumber` | `int` | Current page number. |
| `pageSize` | `int` | Page size |
| `payeeDetails` | `map[string]any` | The Payee object contains details of the beneficiary, person or business. |
| `payerDetails` | `map[string]any` |  |
| `paymentRequestCustomFields` | `map[string]any` | The custom fields that were attached to the payment request that resulted in this transaction. |
| `paymentRequestID` | `string` | For Pay by Bank and Direct Debit transactions this will contain the ID of the payment request. |
| `payoutID` | `string` | ID of the payout that resulted in the transaction. |
| `proprietaryBankTransactionCode` | `map[string]any` |  |
| `rawReference` | `string` | The raw payment reference details as received from the payment processor. |
| `reference` | `string` |  |
| `ruleID` | `string` | ID of the rule that resulted in the transaction. |
| `statementReferences` | `[]any` |  |
| `status` | `string` |  |
| `supplementaryData` | `any` |  |
| `tags` | `[]any` | An optional list of descriptive tags attached to the transaction. |
| `theirReference` | `string` | For a pay out the reference that the payer attached for the receiving party. |
| `totalPages` | `int` | Total pages |
| `totalSize` | `int` | Total count |
| `transactionAmount` | `map[string]any` |  |
| `transactionDate` | `string` | Date when the transaction occurred. |
| `transactionInformation` | `[]any` |  |
| `transactionMutability` | `string` |  |
| `type` | `string` | Type of the transaction. |
| `valueDateTime` | `string` |  |
| `virtualIBAN` | `string` | If set it indicates the payin was to a virtual IBAN. |
| `yourReference` | `string` | For a pay in the reference the sending party attached. |

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
| `clientSessionTimeouts` | `[]any` | The number of seconds a session for this user should last before expiring. |
| `emailAddress` | `string` |  |
| `firstName` | `string` |  |
| `id` | `string` |  |
| `lastName` | `string` |  |
| `passkeyAdded` | `bool` |  |
| `permissions` | `map[string]any` |  |
| `profile` | `string` |  |
| `rolesWithScope` | `[]any` |  |
| `twoFactorEnabled` | `bool` |  |
| `userInviteID` | `string` | Optional ID of the invite that is being accepted so the user can be assigned a role on a new merchant. |

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
| `initialRoleID` | `string` | The role ID to automatically assign to the merchant’s very first user. |
| `inviteeEmailAddress` | `string` | Email address of the user being invited. |
| `inviteeFirstName` | `string` | First Name of the user being invited. |
| `inviteeLastName` | `string` | Last Name of the user being invited. |
| `inviterEmailAddress` | `string` |  |
| `inviterFirstName` | `string` |  |
| `inviterLastName` | `string` |  |
| `isAuthorised` | `bool` | Will be set to true once the invite has met the authorisation requirements. |
| `isInviteeRegistered` | `bool` | If true, indicates the invitee's email address corresponds to an existing MoneyMoov user. |
| `lastInvited` | `string` |  |
| `merchantID` | `string` | ID of the merchant the user is being invited to. |
| `merchantName` | `string` |  |
| `message` | `string` |  |
| `registrationUrl` | `string` |  |
| `sendInviteEmail` | `bool` | If set to true an email will be sent to the invitee with instructions on how to accept the invite. |
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
    "id": "example_id",
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
| `accountName` | `string` | Name for the account |
| `accountSupplierName` | `string` | The payment account supplier name. |
| `availableBalance` | `float64` | The current available balance of the account. |
| `availableBalanceMinorUnits` | `int` | The available balance expressed in the currency’s minor units (e.g. |
| `balance` | `float64` | Balance of the account. |
| `balanceMinorUnits` | `int` | Balance of the account expressed in the currency’s minor units (e.g. |
| `bankName` | `string` | The bank name for external accounts |
| `consentID` | `string` | The ID of the consent used to connect the external account. |
| `createdBy` | `map[string]any` |  |
| `createdByDisplayName` | `string` | Either the name of the user, merchant token or api key that created the account |
| `currency` | `string` | Currency of the account in ISO 4217 format |
| `defaultPaymentRail` | `string` | Indicates the default payment rail for this account. |
| `displayName` | `string` | Gets a unique display name for the payment account. |
| `expiryDate` | `string` | The date that the external account will expire |
| `externalAccountIcon` | `string` | The Icon for external accounts |
| `id` | `string` | Unique id for the account. |
| `identifier` | `map[string]any` |  |
| `inserted` | `string` | Timestamp when the account was created. |
| `isArchived` | `bool` | Indicates whether the account is archived. |
| `isConnectedAccount` | `bool` | Indicates if the payment account is an externally connected account. |
| `isDefault` | `bool` | Is the default account |
| `isTrustAccount` | `bool` | Indicates if the payment account is a trust account. |
| `isVirtual` | `bool` | True if the account is a virtual account. |
| `lastTransaction` | `map[string]any` |  |
| `lastUpdated` | `string` | Timestamp when the account was last updated. |
| `merchantID` | `string` | The ID of the merchant that owns the account. |
| `merchantName` | `string` | The name of the merchant that owns the account. |
| `name` | `string` | The name of the virtual account. |
| `physicalAccountID` | `string` | For virtual accounts this is the ID of the physical account that the virtual account is linked to. |
| `rules` | `[]any` | The list of rules associated with this account. |
| `submittedPayoutsBalance` | `float64` | Total of the payouts that have been submitted for processing. |
| `submittedPayoutsBalanceMinorUnits` | `int` | The balance of the submitted payouts expressed in the currency’s minor units (e.g. |
| `summary` | `string` | Gets a summary of the payments account's most important properties. |
| `supplierSepaInstantStatus` | `string` | Indicates the status of the SEPA Instant payment rail for this account. |
| `xeroBankFeedConnectionStatus` | `string` | States the status of the Xero bank feed connection, if applicable. |
| `xeroBankFeedLastSyncedAt` | `string` |  |
| `xeroBankFeedSyncLastFailedAt` | `string` |  |
| `xeroBankFeedSyncLastFailureReason` | `string` |  |
| `xeroBankFeedSyncStatus` | `string` |  |
| `xeroUnsynchronisedTransactionsCount` | `int` | Indicates the number of unsynchronised transactions with Xero |

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
| `destinationUrl` | `string` | The destination URL for the webhook. |
| `emailAddress` | `string` | The recipient email address(es) for notifications. |
| `failedNotificationEmailAddress` | `string` | The email address to which notifications about failed webhook deliveries will be sent. |
| `id` | `string` |  |
| `isActive` | `bool` |  |
| `merchantID` | `string` | The ID of the merchant that the webhook is for. |
| `notificationMethod` | `string` | The type of notification that will be sent. |
| `resourceTypes` | `[]any` | The resource types that the webhook will be generated for. |
| `retry` | `bool` |  |
| `secret` | `string` | The secret key required to authenticate webhook notifications. |
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
