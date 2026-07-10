# nofrixion-mcp

MCP server exposing the Nofrixion SDK as tools, built on the
[official Go MCP SDK](https://github.com/modelcontextprotocol/go-sdk)
and the sibling Go SDK at `../go`.

## Tools

| Tool | Args | Returns |
|------|------|---------|
| `nofrixion_list` | `entity`, optional `query` map | First page of records as JSON |
| `nofrixion_load` | `entity`, `query` (e.g. `{id:N}`) | Single record as JSON |

JSON schemas for both tools are emitted by the SDK from the `Args`
struct's `json` / `jsonschema` tags — no schema is hand-written.

## Entities

account | beneficiary | cancel | disable | enable | merchant | metadata | no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page | no_frixion_biz_biz_models_paging_payment_request_page | no_frixion_biz_biz_models_paging_payout_page | no_frixion_biz_biz_models_paging_payrun_page | no_frixion_biz_biz_models_paging_rule_events_page | no_frixion_biz_biz_models_paging_rules_page | no_frixion_biz_biz_models_payments_card_payment | no_frixion_biz_biz_models_payments_card_public_key | no_frixion_money_moov_api_features_beneficiaries_beneficiaries | no_frixion_money_moov_api_features_payment_requests_payment | no_frixion_money_moov_api_features_permissions_roles_create | no_frixion_money_moov_api_features_user_invites_create | no_frixion_money_moov_models_authorisation_settings_merchant | no_frixion_money_moov_models_batch_payout | no_frixion_money_moov_models_beneficiary_group_page | no_frixion_money_moov_models_beneficiary_page | no_frixion_money_moov_models_card_customer_token | no_frixion_money_moov_models_currency_currency_info | no_frixion_money_moov_models_direct_debit_batch_submit | no_frixion_money_moov_models_fx_rate | no_frixion_money_moov_models_i_payment | no_frixion_money_moov_models_mandates_mandate | no_frixion_money_moov_models_merchant | no_frixion_money_moov_models_merchant_page | no_frixion_money_moov_models_merchant_pay_by_bank_setting | no_frixion_money_moov_models_merchant_token | no_frixion_money_moov_models_merchant_token_page | no_frixion_money_moov_models_no_frixion_version | no_frixion_money_moov_models_open_banking_account | no_frixion_money_moov_models_open_banking_consent | no_frixion_money_moov_models_open_banking_transaction | no_frixion_money_moov_models_payment | no_frixion_money_moov_models_payment_account_minimal_page | no_frixion_money_moov_models_payment_account_page | no_frixion_money_moov_models_payment_initiation | no_frixion_money_moov_models_payment_request_event | no_frixion_money_moov_models_payment_request_metric | no_frixion_money_moov_models_payment_request_minimal | no_frixion_money_moov_models_payment_request_result | no_frixion_money_moov_models_payment_requests_merchant_payment | no_frixion_money_moov_models_payment_requests_merchant_payment2 | no_frixion_money_moov_models_payment_requests_merchant_payment3 | no_frixion_money_moov_models_payment_requests_merchant_payment4 | no_frixion_money_moov_models_payout_keyset_page | no_frixion_money_moov_models_payout_metric | no_frixion_money_moov_models_payouts_payouts_create | no_frixion_money_moov_models_payrun | no_frixion_money_moov_models_report_result | no_frixion_money_moov_models_rule | no_frixion_money_moov_models_transaction | no_frixion_money_moov_models_transaction_page | no_frixion_money_moov_models_user_invite | no_frixion_money_moov_models_user_invite_page | no_frixion_money_moov_models_user_page | no_frixion_money_moov_models_webhook | open_banking | payeeverification | payment_request | payout | payrun | reject | report | rule | send | sendbeneficiary | tag | token | transaction | user | user_invite | virtual | webhook | whoami | whoamitrustedapp

## Build

```sh
go build -o nofrixion-mcp ./...
```

## Run

```sh
# stdio transport — for spawnable agent installs (default).
./nofrixion-mcp

# streamable HTTP transport — share one running server between agents.
./nofrixion-mcp -transport http -addr :8080
```

## Install for Claude Code

```sh
claude mcp add --scope user nofrixion \
  -- /absolute/path/to/nofrixion-mcp -transport stdio
```

After install, restart Claude Code; the `nofrixion_list` and
`nofrixion_load` tools become available in new sessions.

## Smoke test via HTTP

```sh
./nofrixion-mcp -transport http -addr :18080 &

# initialize, grab the session id
curl -sN -X POST http://localhost:18080 \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json, text/event-stream' \
  -D headers \
  -d '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-06-18","capabilities":{},"clientInfo":{"name":"smoke","version":"0"}}}'

SESSION=$(awk '/Mcp-Session-Id/ {print $2}' headers | tr -d '\r')

curl -sN -X POST http://localhost:18080 \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json, text/event-stream' \
  -H "Mcp-Session-Id: $SESSION" \
  -d '{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"nofrixion_load","arguments":{"entity":"account","query":{"id":1}}}}'
```

## Generated by

sdkgen `go-mcp` target. See the target source under
`.sdk/src/cmp/go-mcp/` in this repo, or upstream at
`github.com/voxgig/sdkgen/project/.sdk/src/cmp/go-mcp/`.
