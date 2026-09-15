---
theme: none
title: "NoFrixion MoneyMoov API"
fonts:
  provider: none
mdc: false
colorSchema: light
canvasWidth: 1120
aspectRatio: 16/9
transition: none
---

# NoFrixion MoneyMoov API

The NoFrixion MoneyMoov API.

---

# What this SDK gives you

- 49 entities across 187 HTTP routes
- 6 language SDKs and 2 companion tools
- 1 optional feature built in, all off until you enable them

An entity groups related operations, and each operation may cover several routes. The SDK exposes both using the conventions of your language, so you write `client` calls rather than HTTP requests.

---

# API capabilities 1 / 3

- Account: `create`, `list`, `load`, `remove`, `update`
- Batch: `create`, `load`
- Beneficiary: `create`, `list`, `load`, `remove`, `update`
- BeneficiaryGroup: `list`
- Card: `create`
- CardCustomerToken: `list`, `load`, `remove`

---

# API capabilities 2 / 3

- CardPayment: `create`
- CardPublicKey: `load`
- Consent: `create`, `list`, `load`, `remove`, `update`
- Currency: `list`
- DirectDebitBatchSubmit: `create`
- FxRate: `list`, `load`

---

# API capabilities 3 / 3

- IPayment: `create`
- Mandate: `create`, `load`
- Merchant: `list`, `load`, `remove`, `update`
- MerchantAuthorisationSetting: `list`
- MerchantDirectDebitMandatePage: `list`
- and 32 more, in the API reference

---

# Authentication

The API uses `apiKey` authentication. Pass the credential when you construct the client, and read it from the environment rather than from source.

---

# Features come with the SDK

Retries, caching, paging, logging and the rest are compiled in and switched off. You turn on the ones your application needs in the client configuration, and pay for nothing you leave alone.

---

# Built-in features

- In-memory mock transport for testing without a live server

---

# SDKs and tools 1 / 2

- Golang: build from `go/`
- Go CLI: build from `go-cli/`
- Go MCP server: build from `go-mcp/`
- Lua: build from `lua/`
- PHP: build from `php/`

---

# SDKs and tools 2 / 2

- Python: build from `py/`
- Ruby: build from `rb/`
- TypeScript: build from `ts/`

---

# Tutorial: your first call

Five steps, in TypeScript. Every other language in this repository follows the same shape.

1. Install the SDK.
2. Construct a client.
3. Call an operation.
4. Handle the failure.
5. Turn on a feature.

---

# Step 1: install

This target is not published yet. Build it from `ts/` in the repository.

---

# Step 2: construct a client

```ts
import { NofrixionSDK } from '@voxgig-sdk/nofrixion'

const client = new NofrixionSDK({ apikey: process.env.NOFRIXION_APIKEY })
```

The credential is read from the environment, so nothing secret reaches your source tree.

---

# Step 3: call an operation

```ts
const result = await client.Account().create({
  // the input this operation requires
})
```

That reaches `POST /api/v1/accounts/{accountID}/{currency}`. The reference page for Account lists every field it expects.

---

# Step 4: handle the failure

```ts
try {
  const result = await client.Account().create({ /* input */ })
} catch (err) {
  // authentication, validation, or the API itself
}
```

Operations fail loudly rather than returning an empty result, so an error is never mistaken for no data.

---

# Step 5: turn on a feature

```json
{ "test": { "active": true } }
```

Pass that in the client configuration and in-memory mock transport for testing without a live server is live. The feature pages document every option and its default.

---

# This SDK is generated

Nothing here was written by hand. An API definition goes in, and every language, test suite and page of documentation comes out.

So the useful question is not how to patch it. It is what to regenerate, and what survives when you do.

---

# Regenerate it

```sh
npm run generate
```

Update the API definition in `.sdk/def/` and regenerate. New routes, new fields and new entities reach every target at once.

---

# Add a language

```sh
voxgig-sdkgen target add <language>
```

This repository already builds 6 of them. A target added today is generated from the same model as the ones already here, so it arrives complete rather than as a stub.

---

# Add a feature

```sh
voxgig-sdkgen feature add <feature>
```

A feature is behaviour around the call, not a change to it: retries, caching, audit trails, secret resolution. It lands in every target that supports it.

---

# What survives a regeneration

- `.sdk/model/project.aon` is yours. Settings belong there.
- Everything generated is overwritten, every time.

An edit to a generated file works until the next regeneration and then disappears without warning, which is the one failure mode worth knowing in advance.

---

# Where to go next

- Read the API reference for the operation you need.
- Pick a language and follow its installation page.
- Review the feature pages before enabling anything in production.
