
import { BaseFeature } from './feature/base/BaseFeature'
import { TestFeature } from './feature/test/TestFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   test: TestFeature,

}


class Config {

  makeFeature(this: any, fn: string) {
    const fc = FEATURE_CLASS[fn]
    const fi = new fc()
    // TODO: errors etc
    return fi
  }


  main = {
    name: 'ProjectName',
  }


  feature = {
     test:     {
      "options": {
        "active": false
      }
    },

  }


  options = {
    base: 'https://api-sandbox.nofrixion.com',

    auth: {
      prefix: 'Bearer',
    },

    headers: {
      "content-type": "application/json"
    },

    entity: {
      
      account: {
      },

      batch: {
      },

      beneficiary: {
      },

      beneficiary_group: {
      },

      card: {
      },

      card_customer_token: {
      },

      card_payment: {
      },

      card_public_key: {
      },

      consent: {
      },

      currency: {
      },

      direct_debit_batch_submit: {
      },

      fx_rate: {
      },

      i_payment: {
      },

      mandate: {
      },

      merchant: {
      },

      merchant_authorisation_setting: {
      },

      merchant_direct_debit_mandate_page: {
      },

      merchant_pay_by_bank_setting: {
      },

      merchant_payment_request_template: {
      },

      merchant_token: {
      },

      metadata: {
      },

      no_frixion_version: {
      },

      open_banking: {
      },

      payeeverification: {
      },

      payment: {
      },

      payment_account: {
      },

      payment_account_minimal: {
      },

      payment_initiation: {
      },

      payment_request: {
      },

      payment_request_event: {
      },

      payment_request_metric: {
      },

      payment_request_minimal: {
      },

      payment_request_result: {
      },

      payout: {
      },

      payout_keyset_page: {
      },

      payout_metric: {
      },

      payrun: {
      },

      report: {
      },

      report_result: {
      },

      role: {
      },

      rule: {
      },

      rule_event: {
      },

      tag: {
      },

      token: {
      },

      transaction: {
      },

      user: {
      },

      user_invite: {
      },

      virtual: {
      },

      webhook: {
      },

    }
  }


  entity = {
    "account": {
      "fields": [
        {
          "active": true,
          "name": "account_balance",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "account_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "account_identification",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 2
        },
        {
          "active": true,
          "name": "account_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "account_supplier_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "account_type",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "available_balance",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 6
        },
        {
          "active": true,
          "name": "available_balance_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 7
        },
        {
          "active": true,
          "name": "balance",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 8
        },
        {
          "active": true,
          "name": "balance_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 9
        },
        {
          "active": true,
          "name": "bank_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "consent_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "consolidated_account_information",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 12
        },
        {
          "active": true,
          "name": "created_by",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 13
        },
        {
          "active": true,
          "name": "created_by_display_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "default_payment_rail",
          "req": false,
          "type": "`$STRING`",
          "index$": 16
        },
        {
          "active": true,
          "name": "description",
          "req": false,
          "type": "`$STRING`",
          "index$": 17
        },
        {
          "active": true,
          "name": "detail",
          "req": false,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "display_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 19
        },
        {
          "active": true,
          "name": "expiry_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 20
        },
        {
          "active": true,
          "name": "external_account_icon",
          "req": false,
          "type": "`$STRING`",
          "index$": 21
        },
        {
          "active": true,
          "name": "format",
          "req": false,
          "type": "`$STRING`",
          "index$": 22
        },
        {
          "active": true,
          "name": "from_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 23
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 24
        },
        {
          "active": true,
          "name": "identifier",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 25
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 26
        },
        {
          "active": true,
          "name": "is_archived",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 27
        },
        {
          "active": true,
          "name": "is_connected_account",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 28
        },
        {
          "active": true,
          "name": "is_default",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 29
        },
        {
          "active": true,
          "name": "is_trust_account",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 30
        },
        {
          "active": true,
          "name": "is_virtual",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 31
        },
        {
          "active": true,
          "name": "last_transaction",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 32
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 33
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 34
        },
        {
          "active": true,
          "name": "merchant_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 35
        },
        {
          "active": true,
          "name": "nickname",
          "req": false,
          "type": "`$STRING`",
          "index$": 36
        },
        {
          "active": true,
          "name": "physical_account_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 37
        },
        {
          "active": true,
          "name": "role_i_d",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 38
        },
        {
          "active": true,
          "name": "rule",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 39
        },
        {
          "active": true,
          "name": "submitted_payouts_balance",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 40
        },
        {
          "active": true,
          "name": "submitted_payouts_balance_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 41
        },
        {
          "active": true,
          "name": "summary",
          "req": false,
          "type": "`$STRING`",
          "index$": 42
        },
        {
          "active": true,
          "name": "supplier_physical_account_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 43
        },
        {
          "active": true,
          "name": "supplier_sepa_instant_status",
          "req": false,
          "type": "`$STRING`",
          "index$": 44
        },
        {
          "active": true,
          "name": "to_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 45
        },
        {
          "active": true,
          "name": "type",
          "req": false,
          "type": "`$STRING`",
          "index$": 46
        },
        {
          "active": true,
          "name": "usage_type",
          "req": false,
          "type": "`$STRING`",
          "index$": 47
        },
        {
          "active": true,
          "name": "xero_bank_feed_connection_status",
          "req": false,
          "type": "`$STRING`",
          "index$": 48
        },
        {
          "active": true,
          "name": "xero_bank_feed_last_synced_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 49
        },
        {
          "active": true,
          "name": "xero_bank_feed_sync_last_failed_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 50
        },
        {
          "active": true,
          "name": "xero_bank_feed_sync_last_failure_reason",
          "req": false,
          "type": "`$STRING`",
          "index$": 51
        },
        {
          "active": true,
          "name": "xero_bank_feed_sync_status",
          "req": false,
          "type": "`$STRING`",
          "index$": 52
        },
        {
          "active": true,
          "name": "xero_unsynchronised_transactions_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 53
        }
      ],
      "name": "account",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/accounts/{accountID}/{currency}",
              "parts": [
                "api",
                "v1",
                "accounts",
                "{account_id}",
                "{currency}"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id"
                }
              },
              "select": {
                "exist": [
                  "account_id",
                  "currency"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/accounts/{accountID}/statements",
              "parts": [
                "api",
                "v1",
                "accounts",
                "{account_id}",
                "statements"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id"
                }
              },
              "select": {
                "$action": "statement",
                "exist": [
                  "account_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/accounts",
              "parts": [
                "api",
                "v1",
                "accounts"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "connected_account",
                    "orig": "connected_account",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "only_connect_account",
                    "orig": "only_connect_account",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/accounts",
              "parts": [
                "api",
                "v1",
                "accounts"
              ],
              "select": {
                "exist": [
                  "connected_account",
                  "include_archived",
                  "merchant_id",
                  "only_connect_account"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "connected_account",
                    "orig": "connected_account",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/accounts",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "accounts"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "connected_account",
                  "merchant_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "connected_account",
                    "orig": "connected_account",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_child_merchant",
                    "orig": "include_child_merchant",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "only_archived",
                    "orig": "only_archived",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "only_connect_account",
                    "orig": "only_connect_account",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/accounts/export",
              "parts": [
                "api",
                "v1",
                "accounts",
                "export"
              ],
              "select": {
                "$action": "export",
                "exist": [
                  "connected_account",
                  "currency",
                  "include_archived",
                  "include_child_merchant",
                  "merchant_id",
                  "only_archived",
                  "only_connect_account",
                  "page_number",
                  "page_size",
                  "search",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "credit_type",
                    "orig": "credit_type",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/accounts/{accountID}/transactions/export",
              "parts": [
                "api",
                "v1",
                "accounts",
                "{account_id}",
                "transactions",
                "export"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id"
                }
              },
              "select": {
                "$action": "transaction_export",
                "exist": [
                  "account_id",
                  "credit_type",
                  "from_date",
                  "max_amount",
                  "min_amount",
                  "page_number",
                  "page_size",
                  "search",
                  "sort",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/accounts/{accountID}/statements/{id}",
              "parts": [
                "api",
                "v1",
                "accounts",
                "{account_id}",
                "statements",
                "{id}"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id"
                }
              },
              "select": {
                "exist": [
                  "account_id",
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/accounts/{accountID}",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "accounts",
                "{id}"
              ],
              "rename": {
                "param": {
                  "accountID": "id",
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "merchant_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 3
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/accounts/{accountID}",
              "parts": [
                "api",
                "v1",
                "accounts",
                "{id}"
              ],
              "rename": {
                "param": {
                  "accountID": "id"
                }
              },
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 4
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/openbanking/accounts/{id}",
              "parts": [
                "api",
                "v1",
                "openbanking",
                "accounts",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 5
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/accounts/statements",
              "parts": [
                "api",
                "v1",
                "accounts",
                "statements"
              ],
              "select": {
                "$action": "statement"
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 6
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/accounts/archive/{id}",
              "parts": [
                "api",
                "v1",
                "accounts",
                "archive",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {},
              "method": "DELETE",
              "orig": "/api/v1/accounts/statements",
              "parts": [
                "api",
                "v1",
                "accounts",
                "statements"
              ],
              "select": {
                "$action": "statement"
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "amount",
                    "orig": "amount",
                    "reqd": true,
                    "type": "`$NUMBER`",
                    "index$": 1
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/accounts/{accountID}/topup/{amount}",
              "parts": [
                "api",
                "v1",
                "accounts",
                "{account_id}",
                "topup",
                "{amount}"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id"
                }
              },
              "select": {
                "exist": [
                  "account_id",
                  "amount"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/accounts/unarchive/{id}",
              "parts": [
                "api",
                "v1",
                "accounts",
                "unarchive",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/accounts/{id}",
              "parts": [
                "api",
                "v1",
                "accounts",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": [
          [
            "account"
          ],
          [
            "merchant"
          ],
          [
            "account",
            "topup"
          ]
        ]
      }
    },
    "batch": {
      "fields": [
        {
          "active": true,
          "name": "approve_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "payout",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 2
        }
      ],
      "name": "batch",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/payouts/batch",
              "parts": [
                "api",
                "v1",
                "payouts",
                "batch"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/payouts/batch/{id}",
              "parts": [
                "api",
                "v1",
                "payouts",
                "batch",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "beneficiary": {
      "fields": [
        {
          "active": true,
          "name": "approval_callback_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "authentication_method",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 1
        },
        {
          "active": true,
          "name": "authorisation",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 2
        },
        {
          "active": true,
          "name": "authorisers_completed_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 3
        },
        {
          "active": true,
          "name": "authorisers_required_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 4
        },
        {
          "active": true,
          "name": "beneficiary",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 5
        },
        {
          "active": true,
          "name": "beneficiary_event",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 6
        },
        {
          "active": true,
          "name": "can_authorise",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 7
        },
        {
          "active": true,
          "name": "can_update",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 8
        },
        {
          "active": true,
          "name": "created_by",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 9
        },
        {
          "active": true,
          "name": "created_by_email_address",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "currency",
          "op": {
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "destination",
          "op": {
            "create": {
              "req": true,
              "type": "`$OBJECT`"
            }
          },
          "req": false,
          "type": "`$OBJECT`",
          "index$": 12
        },
        {
          "active": true,
          "name": "failed_beneficiary",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 13
        },
        {
          "active": true,
          "name": "has_current_user_authorised",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 14
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 16
        },
        {
          "active": true,
          "name": "is_enabled",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 17
        },
        {
          "active": true,
          "name": "last_authorised",
          "req": false,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 19
        },
        {
          "active": true,
          "name": "merchant_id",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$STRING`",
          "index$": 20
        },
        {
          "active": true,
          "name": "name",
          "op": {
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 21
        },
        {
          "active": true,
          "name": "nonce",
          "req": false,
          "type": "`$STRING`",
          "index$": 22
        },
        {
          "active": true,
          "name": "source_account",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 23
        },
        {
          "active": true,
          "name": "source_account_i_d",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 24
        },
        {
          "active": true,
          "name": "their_reference",
          "req": false,
          "type": "`$STRING`",
          "index$": 25
        }
      ],
      "name": "beneficiary",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/beneficiaries/authorise/{id}",
              "parts": [
                "api",
                "v1",
                "beneficiaries",
                "authorise",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/beneficiaries",
              "parts": [
                "api",
                "v1",
                "beneficiaries"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/beneficiaries/batchcreate",
              "parts": [
                "api",
                "v1",
                "beneficiaries",
                "batchcreate"
              ],
              "select": {
                "$action": "batchcreate"
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_disabled",
                    "orig": "include_disabled",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "source_account_id",
                    "orig": "source_account_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/beneficiaries",
              "parts": [
                "api",
                "v1",
                "beneficiaries"
              ],
              "select": {
                "exist": [
                  "currency",
                  "include_disabled",
                  "merchant_id",
                  "page_number",
                  "page_size",
                  "search",
                  "sort",
                  "source_account_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_disabled",
                    "orig": "include_disabled",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "source_account_id",
                    "orig": "source_account_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/beneficiaries",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "beneficiaries"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "currency",
                  "include_disabled",
                  "merchant_id",
                  "page_number",
                  "page_size",
                  "search",
                  "sort",
                  "source_account_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_disabled",
                    "orig": "include_disabled",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/beneficiaries/export",
              "parts": [
                "api",
                "v1",
                "beneficiaries",
                "export"
              ],
              "select": {
                "$action": "export",
                "exist": [
                  "currency",
                  "include_disabled",
                  "merchant_id",
                  "page_number",
                  "page_size",
                  "search",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/beneficiaries/{id}",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "beneficiaries",
                "{id}"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "merchant_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/beneficiaries/{id}",
              "parts": [
                "api",
                "v1",
                "beneficiaries",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/beneficiaries/{id}",
              "parts": [
                "api",
                "v1",
                "beneficiaries",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/beneficiaries/disable/{id}",
              "parts": [
                "api",
                "v1",
                "beneficiaries",
                "disable",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/beneficiaries/enable/{id}",
              "parts": [
                "api",
                "v1",
                "beneficiaries",
                "enable",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/beneficiaries/{id}",
              "parts": [
                "api",
                "v1",
                "beneficiaries",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": [
          [
            "merchant"
          ]
        ]
      }
    },
    "beneficiary_group": {
      "fields": [
        {
          "active": true,
          "name": "group_member",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "group_name",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 5
        }
      ],
      "name": "beneficiary_group",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/beneficiarygroups",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "beneficiarygroups"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "merchant_id",
                  "page_number",
                  "page_size"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": [
          [
            "merchant"
          ]
        ]
      }
    },
    "card": {
      "fields": [
        {
          "active": true,
          "name": "authorized_amount",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "currency_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "is_payer_authentication_required",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 2
        },
        {
          "active": true,
          "name": "is_soft_decline",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 3
        },
        {
          "active": true,
          "name": "payer_authentication_access_token",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "payer_authentication_merchant_data",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "payer_authentication_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "payer_authentication_window_height",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 7
        },
        {
          "active": true,
          "name": "payer_authentication_window_width",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 8
        },
        {
          "active": true,
          "name": "payment_request_callback_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "payment_request_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "request_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "response_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        },
        {
          "active": true,
          "name": "response_type",
          "req": false,
          "type": "`$STRING`",
          "index$": 13
        },
        {
          "active": true,
          "name": "status",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "three_ds_redirect_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "transaction_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 16
        }
      ],
      "name": "card",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/paymentrequests/{id}/card",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "card"
              ],
              "rename": {
                "param": {
                  "id": "paymentrequest_id"
                }
              },
              "select": {
                "exist": [
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        }
      },
      "relations": {
        "ancestors": [
          [
            "paymentrequest"
          ]
        ]
      }
    },
    "card_customer_token": {
      "fields": [
        {
          "active": true,
          "name": "card_type",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "customer_email_address",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "expiry_month",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "expiry_year",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "last_four_digit",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "masked_card_number",
          "req": false,
          "type": "`$STRING`",
          "index$": 8
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "payment_request_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        }
      ],
      "name": "card_customer_token",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "customer_email_address",
                    "orig": "customer_email_address",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/paymentrequests/card/customertokens/{merchantID}/{customerEmailAddress}",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "card",
                "customertokens",
                "{merchant_id}",
                "{customer_email_address}"
              ],
              "rename": {
                "param": {
                  "customerEmailAddress": "customer_email_address",
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "customer_email_address",
                  "merchant_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "customer_email_address",
                    "orig": "customer_email_address",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/paymentrequests/card/customertokens/{customerEmailAddress}",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "card",
                "customertokens",
                "{customer_email_address}"
              ],
              "rename": {
                "param": {
                  "customerEmailAddress": "customer_email_address"
                }
              },
              "select": {
                "exist": [
                  "customer_email_address"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "customer_email_address",
                    "orig": "customer_email_address",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/paymentrequests/card/customertokens/removeall/{merchantID}/{customerEmailAddress}",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "card",
                "customertokens",
                "removeall",
                "{merchant_id}",
                "{customer_email_address}"
              ],
              "rename": {
                "param": {
                  "customerEmailAddress": "customer_email_address",
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "customer_email_address",
                  "merchant_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "customer_email_address",
                    "orig": "customer_email_address",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/paymentrequests/card/customertokens/removeall/{customerEmailAddress}",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "card",
                "customertokens",
                "removeall",
                "{customer_email_address}"
              ],
              "rename": {
                "param": {
                  "customerEmailAddress": "customer_email_address"
                }
              },
              "select": {
                "exist": [
                  "customer_email_address"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/paymentrequests/card/customertokens/{id}",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "card",
                "customertokens",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            }
          ],
          "key$": "remove"
        }
      },
      "relations": {
        "ancestors": [
          [
            "removeall"
          ],
          [
            "customertoken"
          ]
        ]
      }
    },
    "card_payment": {
      "fields": [
        {
          "active": true,
          "name": "authorized_amount",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "currency_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "is_payer_authentication_required",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 2
        },
        {
          "active": true,
          "name": "is_soft_decline",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 3
        },
        {
          "active": true,
          "name": "payer_authentication_access_token",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "payer_authentication_merchant_data",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "payer_authentication_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "payer_authentication_window_height",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 7
        },
        {
          "active": true,
          "name": "payer_authentication_window_width",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 8
        },
        {
          "active": true,
          "name": "payment_request_callback_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "payment_request_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "request_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "response_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        },
        {
          "active": true,
          "name": "response_type",
          "req": false,
          "type": "`$STRING`",
          "index$": 13
        },
        {
          "active": true,
          "name": "status",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "three_ds_redirect_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "transaction_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 16
        }
      ],
      "name": "card_payment",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "partial_refund_amount",
                    "orig": "partial_refund_amount",
                    "reqd": true,
                    "type": "`$NUMBER`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/paymentrequests/{id}/card/refund/{partialRefundAmount}",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "card",
                "refund",
                "{partial_refund_amount}"
              ],
              "rename": {
                "param": {
                  "id": "paymentrequest_id",
                  "partialRefundAmount": "partial_refund_amount"
                }
              },
              "select": {
                "exist": [
                  "partial_refund_amount",
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/paymentrequests/{id}/card/capture",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "card",
                "capture"
              ],
              "rename": {
                "param": {
                  "id": "paymentrequest_id"
                }
              },
              "select": {
                "exist": [
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/paymentrequests/{id}/card/paywithtoken",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "card",
                "paywithtoken"
              ],
              "rename": {
                "param": {
                  "id": "paymentrequest_id"
                }
              },
              "select": {
                "exist": [
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/paymentrequests/{id}/card/void",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "card",
                "void"
              ],
              "rename": {
                "param": {
                  "id": "paymentrequest_id"
                }
              },
              "select": {
                "exist": [
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 3
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/paymentrequests/{id}/card/voidpaymentrequest",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "card",
                "voidpaymentrequest"
              ],
              "rename": {
                "param": {
                  "id": "paymentrequest_id"
                }
              },
              "select": {
                "exist": [
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 4
            }
          ],
          "key$": "create"
        }
      },
      "relations": {
        "ancestors": [
          [
            "paymentrequest"
          ],
          [
            "paymentrequest",
            "refund"
          ]
        ]
      }
    },
    "card_public_key": {
      "fields": [
        {
          "active": true,
          "name": "jwt",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        }
      ],
      "name": "card_public_key",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/paymentrequests/{id}/card/publickey",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "card",
                "publickey"
              ],
              "rename": {
                "param": {
                  "id": "paymentrequest_id"
                }
              },
              "select": {
                "exist": [
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": [
          [
            "paymentrequest"
          ]
        ]
      }
    },
    "consent": {
      "fields": [
        {
          "active": true,
          "name": "authorisation_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "callback_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "consent_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "email_address",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "expiry_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "failure_callback_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "institution_id",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$STRING`",
          "index$": 8
        },
        {
          "active": true,
          "name": "is_connected_account",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 9
        },
        {
          "active": true,
          "name": "is_enabled",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 10
        },
        {
          "active": true,
          "name": "merchant_id",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "provider",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        },
        {
          "active": true,
          "name": "success_web_hook_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 13
        }
      ],
      "name": "consent",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/openbanking/consents",
              "parts": [
                "api",
                "v1",
                "openbanking",
                "consents"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "email",
                    "orig": "email",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/openbanking/consents/{merchantID}/{email}",
              "parts": [
                "api",
                "v1",
                "openbanking",
                "consents",
                "{merchant_id}",
                "{email}"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "email",
                  "merchant_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/openbanking/consents/{id}",
              "parts": [
                "api",
                "v1",
                "openbanking",
                "consents",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/openbanking/consents/{id}",
              "parts": [
                "api",
                "v1",
                "openbanking",
                "consents",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PATCH",
              "orig": "/api/v1/openbanking/consents/{id}",
              "parts": [
                "api",
                "v1",
                "openbanking",
                "consents",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": [
          [
            "consent"
          ]
        ]
      }
    },
    "currency": {
      "fields": [
        {
          "active": true,
          "name": "code",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "decimal",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 1
        },
        {
          "active": true,
          "name": "is_fiat",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 2
        },
        {
          "active": true,
          "name": "iso4217_alpha_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "iso4217_numeric_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "symbol",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        }
      ],
      "name": "currency",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "capability",
                    "orig": "capability",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/currencies",
              "parts": [
                "api",
                "v1",
                "currencies"
              ],
              "select": {
                "exist": [
                  "capability"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "direct_debit_batch_submit": {
      "fields": [
        {
          "active": true,
          "name": "failed_submission",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 0
        },
        {
          "active": true,
          "name": "successful_submission",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 1
        }
      ],
      "name": "direct_debit_batch_submit",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/paymentrequests/directdebit/batchsubmit",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "directdebit",
                "batchsubmit"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "fx_rate": {
      "fields": [
        {
          "active": true,
          "name": "destination_currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "exchange_rate",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 1
        },
        {
          "active": true,
          "name": "expiry_time",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "quote_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "source_currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        }
      ],
      "name": "fx_rate",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "destination",
                    "orig": "destination",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "source",
                    "orig": "source",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/payouts/fxallheldrates/{source}/{destination}",
              "parts": [
                "api",
                "v1",
                "payouts",
                "fxallheldrates",
                "{source}",
                "{destination}"
              ],
              "select": {
                "exist": [
                  "destination",
                  "source"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "destination",
                    "orig": "destination",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "source",
                    "orig": "source",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "valid_for_minute",
                    "orig": "valid_for_minute",
                    "reqd": true,
                    "type": "`$INTEGER`",
                    "index$": 2
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/payouts/fxheldrate/{source}/{destination}/{validForMinutes}",
              "parts": [
                "api",
                "v1",
                "payouts",
                "fxheldrate",
                "{source}",
                "{destination}",
                "{valid_for_minute}"
              ],
              "rename": {
                "param": {
                  "validForMinutes": "valid_for_minute"
                }
              },
              "select": {
                "exist": [
                  "destination",
                  "source",
                  "valid_for_minute"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": [
          [
            "fxallheldrate"
          ],
          [
            "fxheldrate"
          ]
        ]
      }
    },
    "i_payment": {
      "fields": [
        {
          "active": true,
          "name": "payment_request_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "response_type",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        }
      ],
      "name": "i_payment",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/paymentrequests/payondemand",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "payondemand"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "mandate": {
      "fields": [
        {
          "active": true,
          "name": "account_number",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "address_line1",
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "address_line2",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "approved_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "city",
          "req": true,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "country_code",
          "req": true,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "currency",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "customer_account_number",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "customer_city",
          "req": false,
          "type": "`$STRING`",
          "index$": 8
        },
        {
          "active": true,
          "name": "customer_country_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "customer_country_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "customer_email_address",
          "req": false,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "customer_first_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        },
        {
          "active": true,
          "name": "customer_iban",
          "req": false,
          "type": "`$STRING`",
          "index$": 13
        },
        {
          "active": true,
          "name": "customer_last_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "customer_sort_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "email_address",
          "req": true,
          "type": "`$STRING`",
          "index$": 16
        },
        {
          "active": true,
          "name": "first_name",
          "req": true,
          "type": "`$STRING`",
          "index$": 17
        },
        {
          "active": true,
          "name": "iban",
          "req": false,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 19
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 20
        },
        {
          "active": true,
          "name": "is_recurring",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 21
        },
        {
          "active": true,
          "name": "last_name",
          "req": true,
          "type": "`$STRING`",
          "index$": 22
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 23
        },
        {
          "active": true,
          "name": "merchant_id",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$STRING`",
          "index$": 24
        },
        {
          "active": true,
          "name": "postal_code",
          "req": true,
          "type": "`$STRING`",
          "index$": 25
        },
        {
          "active": true,
          "name": "reference",
          "req": false,
          "type": "`$STRING`",
          "index$": 26
        },
        {
          "active": true,
          "name": "sort_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 27
        },
        {
          "active": true,
          "name": "status",
          "req": false,
          "type": "`$STRING`",
          "index$": 28
        },
        {
          "active": true,
          "name": "supplier_bank_account_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 29
        },
        {
          "active": true,
          "name": "supplier_customer_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 30
        },
        {
          "active": true,
          "name": "supplier_mandate_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 31
        },
        {
          "active": true,
          "name": "supplier_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 32
        },
        {
          "active": true,
          "name": "supplier_status",
          "req": false,
          "type": "`$STRING`",
          "index$": 33
        }
      ],
      "name": "mandate",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/mandates",
              "parts": [
                "api",
                "v1",
                "mandates"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/mandates/{id}",
              "parts": [
                "api",
                "v1",
                "mandates",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "merchant": {
      "fields": [
        {
          "active": true,
          "name": "account_currency",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "can_have_trust_account",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 1
        },
        {
          "active": true,
          "name": "card_payment_processor",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "company_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "display_qr_on_hosted_pay",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 4
        },
        {
          "active": true,
          "name": "hosted_pay_version",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 5
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "is_blocked",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 8
        },
        {
          "active": true,
          "name": "is_exited",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 9
        },
        {
          "active": true,
          "name": "is_suspended",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 10
        },
        {
          "active": true,
          "name": "jurisdiction",
          "req": false,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "logo_url_png",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        },
        {
          "active": true,
          "name": "logo_url_svg",
          "req": false,
          "type": "`$STRING`",
          "index$": 13
        },
        {
          "active": true,
          "name": "merchant_category_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "name",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "note",
          "req": false,
          "type": "`$STRING`",
          "index$": 16
        },
        {
          "active": true,
          "name": "parent_merchant",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 17
        },
        {
          "active": true,
          "name": "payment_account",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 18
        },
        {
          "active": true,
          "name": "payment_account_limit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 19
        },
        {
          "active": true,
          "name": "reason",
          "req": false,
          "type": "`$STRING`",
          "index$": 20
        },
        {
          "active": true,
          "name": "short_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 21
        },
        {
          "active": true,
          "name": "supported_payment_methods_list",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 22
        },
        {
          "active": true,
          "name": "suspension_reason",
          "req": false,
          "type": "`$STRING`",
          "index$": 23
        },
        {
          "active": true,
          "name": "tag",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 24
        },
        {
          "active": true,
          "name": "time_zone_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 25
        },
        {
          "active": true,
          "name": "trading_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 26
        },
        {
          "active": true,
          "name": "web_hook_limit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 27
        },
        {
          "active": true,
          "name": "your_role_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 28
        }
      ],
      "name": "merchant",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "example": 1,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/childmerchants",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "childmerchants"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "$action": "childmerchant",
                "exist": [
                  "merchant_id",
                  "page_number",
                  "page_size",
                  "search",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_suspended",
                    "orig": "include_suspended",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "example": 1,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 10,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/paged",
              "parts": [
                "api",
                "v1",
                "merchants",
                "paged"
              ],
              "select": {
                "$action": "paged",
                "exist": [
                  "include_suspended",
                  "page_number",
                  "page_size",
                  "search",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/merchants",
              "parts": [
                "api",
                "v1",
                "merchants"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/metadata/whoamimerchant",
              "parts": [
                "api",
                "v1",
                "metadata",
                "whoamimerchant"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 3
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/metadata/whoamimerchantsigned",
              "parts": [
                "api",
                "v1",
                "metadata",
                "whoamimerchantsigned"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 4
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/metadata/whoamimerchantwhitelist",
              "parts": [
                "api",
                "v1",
                "metadata",
                "whoamimerchantwhitelist"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 5
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/payouts/export",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "payouts",
                "export"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "$action": "payout_export",
                "exist": [
                  "currency",
                  "from_date",
                  "max_amount",
                  "merchant_id",
                  "min_amount",
                  "page_number",
                  "page_size",
                  "search",
                  "sort",
                  "status",
                  "tag",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_disabled",
                    "orig": "include_disabled",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/beneficiaries/export",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "beneficiaries",
                "export"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "$action": "beneficiary_export",
                "exist": [
                  "currency",
                  "include_disabled",
                  "merchant_id",
                  "page_number",
                  "page_size",
                  "search",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{id}"
              ],
              "rename": {
                "param": {
                  "merchantID": "id"
                }
              },
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "user_id",
                    "orig": "user_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/merchants/{merchantId}/users/{userId}",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{id}",
                "users",
                "{user_id}"
              ],
              "rename": {
                "param": {
                  "merchantId": "id",
                  "userId": "user_id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "user_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "tag_id",
                    "orig": "tag_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/merchants/{merchantID}/tags/{tagID}",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "tags",
                "{tag_id}"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id",
                  "tagID": "tag_id"
                }
              },
              "select": {
                "exist": [
                  "merchant_id",
                  "tag_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/merchants/{merchantID}",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{id}"
              ],
              "rename": {
                "param": {
                  "merchantID": "id"
                }
              },
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/merchants/{merchantId}/suspend",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{id}",
                "suspend"
              ],
              "rename": {
                "param": {
                  "merchantId": "id"
                }
              },
              "select": {
                "$action": "suspend",
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": [
          [
            "merchant"
          ],
          [
            "user"
          ],
          [
            "merchant",
            "tag"
          ]
        ]
      }
    },
    "merchant_authorisation_setting": {
      "fields": [
        {
          "active": true,
          "name": "amount_lower",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 0
        },
        {
          "active": true,
          "name": "amount_upper",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 1
        },
        {
          "active": true,
          "name": "authorisation_type",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "beneficiaries_only",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 3
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "last_editor_cant_authorise",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 6
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 8
        },
        {
          "active": true,
          "name": "number_of_authoriser",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 9
        },
        {
          "active": true,
          "name": "role_setting",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 10
        }
      ],
      "name": "merchant_authorisation_setting",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/authorisationsettings",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "authorisationsettings"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "merchant_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": [
          [
            "merchant"
          ]
        ]
      }
    },
    "merchant_direct_debit_mandate_page": {
      "fields": [
        {
          "active": true,
          "name": "approved_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "customer_account_number",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "customer_city",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "customer_country_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "customer_country_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "customer_email_address",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "customer_first_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "customer_iban",
          "req": false,
          "type": "`$STRING`",
          "index$": 8
        },
        {
          "active": true,
          "name": "customer_last_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "customer_sort_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        },
        {
          "active": true,
          "name": "is_recurring",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 13
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "reference",
          "req": false,
          "type": "`$STRING`",
          "index$": 16
        },
        {
          "active": true,
          "name": "status",
          "req": false,
          "type": "`$STRING`",
          "index$": 17
        },
        {
          "active": true,
          "name": "supplier_bank_account_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "supplier_customer_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 19
        },
        {
          "active": true,
          "name": "supplier_mandate_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 20
        },
        {
          "active": true,
          "name": "supplier_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 21
        },
        {
          "active": true,
          "name": "supplier_status",
          "req": false,
          "type": "`$STRING`",
          "index$": 22
        }
      ],
      "name": "merchant_direct_debit_mandate_page",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "mandate_i_d",
                    "orig": "mandate_i_d",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/mandates",
              "parts": [
                "api",
                "v1",
                "mandates"
              ],
              "select": {
                "exist": [
                  "currency",
                  "from_date",
                  "mandate_i_d",
                  "max_amount",
                  "merchant_id",
                  "min_amount",
                  "page",
                  "search",
                  "size",
                  "sort",
                  "status",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "merchant_pay_by_bank_setting": {
      "fields": [
        {
          "active": true,
          "name": "bank_country_code",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "bank_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "bank_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "business_institution_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "logo",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "message",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "message_image_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "order",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 8
        },
        {
          "active": true,
          "name": "personal_institution_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "processor",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "warning_heading",
          "req": false,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "warning_message",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        }
      ],
      "name": "merchant_pay_by_bank_setting",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "country_code",
                    "orig": "country_code",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": "PIS",
                    "kind": "query",
                    "name": "open_banking_operation",
                    "orig": "open_banking_operation",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/banksettings",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "banksettings"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "country_code",
                  "currency",
                  "merchant_id",
                  "open_banking_operation"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": [
          [
            "merchant"
          ]
        ]
      }
    },
    "merchant_payment_request_template": {
      "fields": [
        {
          "active": true,
          "name": "description",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "name",
          "req": true,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "template",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 6
        }
      ],
      "name": "merchant_payment_request_template",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/paymentrequests/{merchantID}/templates",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{merchant_id}",
                "templates"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "merchant_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "template_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/paymentrequests/{merchantID}/templates/{templateID}",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "templates",
                "{id}"
              ],
              "rename": {
                "param": {
                  "merchantID": "paymentrequest_id",
                  "templateID": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "template_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/paymentrequests/{merchantID}/templates/{templateID}",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "templates",
                "{id}"
              ],
              "rename": {
                "param": {
                  "merchantID": "paymentrequest_id",
                  "templateID": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "template_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/paymentrequests/{merchantID}/templates/{templateID}",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "templates",
                "{id}"
              ],
              "rename": {
                "param": {
                  "merchantID": "paymentrequest_id",
                  "templateID": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": [
          [
            "paymentrequest"
          ]
        ]
      }
    },
    "merchant_token": {
      "fields": [
        {
          "active": true,
          "name": "authentication_method",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "authorisation",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 1
        },
        {
          "active": true,
          "name": "authorisers_completed_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "authorisers_required_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 3
        },
        {
          "active": true,
          "name": "can_authorise",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 4
        },
        {
          "active": true,
          "name": "description",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "expires_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "has_current_user_authorised",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 7
        },
        {
          "active": true,
          "name": "hmac_algorithm",
          "req": false,
          "type": "`$STRING`",
          "index$": 8
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "ip_address_whitelist",
          "req": false,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "is_archived",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 12
        },
        {
          "active": true,
          "name": "is_enabled",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 13
        },
        {
          "active": true,
          "name": "last_authorised",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "merchant_id",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$STRING`",
          "index$": 16
        },
        {
          "active": true,
          "name": "nonce",
          "req": true,
          "type": "`$STRING`",
          "index$": 17
        },
        {
          "active": true,
          "name": "permission_type",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 18
        },
        {
          "active": true,
          "name": "request_signature_version",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 19
        },
        {
          "active": true,
          "name": "shared_secret_algorithm",
          "req": false,
          "type": "`$STRING`",
          "index$": 20
        },
        {
          "active": true,
          "name": "shared_secret_base64",
          "req": false,
          "type": "`$STRING`",
          "index$": 21
        },
        {
          "active": true,
          "name": "token",
          "req": false,
          "type": "`$STRING`",
          "index$": 22
        }
      ],
      "name": "merchant_token",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/tokens",
              "parts": [
                "api",
                "v1",
                "tokens"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/tokens",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "tokens"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "merchant_id",
                  "page_number",
                  "page_size"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/tokens/{id}",
              "parts": [
                "api",
                "v1",
                "tokens",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/tokens/{id}",
              "parts": [
                "api",
                "v1",
                "tokens",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": [
          [
            "merchant"
          ]
        ]
      }
    },
    "metadata": {
      "fields": [],
      "name": "metadata",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "id",
                    "orig": "id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "secret",
                    "orig": "secret",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "url",
                    "orig": "url",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/metadata/problemnotification",
              "parts": [
                "api",
                "v1",
                "metadata",
                "problemnotification"
              ],
              "select": {
                "$action": "problemnotification",
                "exist": [
                  "id",
                  "secret",
                  "url"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/metadata/problem",
              "parts": [
                "api",
                "v1",
                "metadata",
                "problem"
              ],
              "select": {
                "$action": "problem"
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "no_frixion_version": {
      "fields": [
        {
          "active": true,
          "name": "build_version",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 0
        },
        {
          "active": true,
          "name": "major_version",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 1
        },
        {
          "active": true,
          "name": "minor_version",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "release_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        }
      ],
      "name": "no_frixion_version",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/metadata/version",
              "parts": [
                "api",
                "v1",
                "metadata",
                "version"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "open_banking": {
      "fields": [],
      "name": "open_banking",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/openbanking/account/{accountID}/synchronise",
              "parts": [
                "api",
                "v1",
                "openbanking",
                "account",
                "{account_id}",
                "synchronise"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id"
                }
              },
              "select": {
                "exist": [
                  "account_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "email",
                    "orig": "email",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/openbanking/consents/{merchantID}/{email}",
              "parts": [
                "api",
                "v1",
                "openbanking",
                "consents",
                "{merchant_id}",
                "{email}"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "email",
                  "merchant_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/openbanking/account/{accountID}",
              "parts": [
                "api",
                "v1",
                "openbanking",
                "account",
                "{account_id}"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id"
                }
              },
              "select": {
                "exist": [
                  "account_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "remove"
        }
      },
      "relations": {
        "ancestors": [
          [
            "account"
          ],
          [
            "consent"
          ]
        ]
      }
    },
    "payeeverification": {
      "fields": [
        {
          "active": true,
          "name": "account_name",
          "req": true,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "account_number",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "iban",
          "req": true,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "payee_verified_account_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "result",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "secondary_identification",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "sort_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        }
      ],
      "name": "payeeverification",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/openbanking/payeeverification",
              "parts": [
                "api",
                "v1",
                "openbanking",
                "payeeverification"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "payment": {
      "fields": [
        {
          "active": true,
          "name": "address",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "amount",
          "op": {
            "create": {
              "req": true,
              "type": "`$NUMBER`"
            }
          },
          "req": false,
          "type": "`$NUMBER`",
          "index$": 1
        },
        {
          "active": true,
          "name": "amount_pending",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "amount_received",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 3
        },
        {
          "active": true,
          "name": "amount_refunded",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 4
        },
        {
          "active": true,
          "name": "auto_send_receipt",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 5
        },
        {
          "active": true,
          "name": "base_origin_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "callback_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "card_authorize_only",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 8
        },
        {
          "active": true,
          "name": "card_create_token",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 9
        },
        {
          "active": true,
          "name": "card_create_token_mode",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "card_ignore_cvn",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 11
        },
        {
          "active": true,
          "name": "card_no_payer_authentication",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 12
        },
        {
          "active": true,
          "name": "card_processor_merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 13
        },
        {
          "active": true,
          "name": "card_stripe_payment_intent_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "card_stripe_payment_intent_secret",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "card_transmit_raw_detail",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 16
        },
        {
          "active": true,
          "name": "created_by_user",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 17
        },
        {
          "active": true,
          "name": "currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "custom_field",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 19
        },
        {
          "active": true,
          "name": "customer_email_address",
          "req": false,
          "type": "`$STRING`",
          "index$": 20
        },
        {
          "active": true,
          "name": "customer_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 21
        },
        {
          "active": true,
          "name": "customer_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 22
        },
        {
          "active": true,
          "name": "description",
          "req": false,
          "type": "`$STRING`",
          "index$": 23
        },
        {
          "active": true,
          "name": "destination_account",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 24
        },
        {
          "active": true,
          "name": "direct_debit_payment",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 25
        },
        {
          "active": true,
          "name": "due_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 26
        },
        {
          "active": true,
          "name": "event",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 27
        },
        {
          "active": true,
          "name": "failure_callback_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 28
        },
        {
          "active": true,
          "name": "field_display_setting",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 29
        },
        {
          "active": true,
          "name": "formatted_amount",
          "req": false,
          "type": "`$STRING`",
          "index$": 30
        },
        {
          "active": true,
          "name": "hosted_pay_checkout_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 31
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 32
        },
        {
          "active": true,
          "name": "ignore_address_verification",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 33
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 34
        },
        {
          "active": true,
          "name": "inserted_sortable",
          "req": false,
          "type": "`$STRING`",
          "index$": 35
        },
        {
          "active": true,
          "name": "is_archived",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 36
        },
        {
          "active": true,
          "name": "jwk",
          "req": false,
          "type": "`$STRING`",
          "index$": 37
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 38
        },
        {
          "active": true,
          "name": "lightning_invoice",
          "req": false,
          "type": "`$STRING`",
          "index$": 39
        },
        {
          "active": true,
          "name": "lightning_invoice_expires_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 40
        },
        {
          "active": true,
          "name": "merchant_direct_debit_mandate_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 41
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 42
        },
        {
          "active": true,
          "name": "merchant_token_description",
          "req": false,
          "type": "`$STRING`",
          "index$": 43
        },
        {
          "active": true,
          "name": "notification_email_address",
          "req": false,
          "type": "`$STRING`",
          "index$": 44
        },
        {
          "active": true,
          "name": "notification_role_i_d",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 45
        },
        {
          "active": true,
          "name": "order_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 46
        },
        {
          "active": true,
          "name": "partial_payment_method",
          "req": false,
          "type": "`$STRING`",
          "index$": 47
        },
        {
          "active": true,
          "name": "partial_payment_step",
          "req": false,
          "type": "`$STRING`",
          "index$": 48
        },
        {
          "active": true,
          "name": "payment_attempt",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 49
        },
        {
          "active": true,
          "name": "payment_method",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 50
        },
        {
          "active": true,
          "name": "payment_processor",
          "req": false,
          "type": "`$STRING`",
          "index$": 51
        },
        {
          "active": true,
          "name": "payrun_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 52
        },
        {
          "active": true,
          "name": "pisp_account_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 53
        },
        {
          "active": true,
          "name": "priority_bank_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 54
        },
        {
          "active": true,
          "name": "result",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 55
        },
        {
          "active": true,
          "name": "sandbox_settle_delay_in_second",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 56
        },
        {
          "active": true,
          "name": "shipping_address",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 57
        },
        {
          "active": true,
          "name": "shipping_address_city",
          "req": false,
          "type": "`$STRING`",
          "index$": 58
        },
        {
          "active": true,
          "name": "shipping_address_country_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 59
        },
        {
          "active": true,
          "name": "shipping_address_county",
          "req": false,
          "type": "`$STRING`",
          "index$": 60
        },
        {
          "active": true,
          "name": "shipping_address_line1",
          "req": false,
          "type": "`$STRING`",
          "index$": 61
        },
        {
          "active": true,
          "name": "shipping_address_line2",
          "req": false,
          "type": "`$STRING`",
          "index$": 62
        },
        {
          "active": true,
          "name": "shipping_address_post_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 63
        },
        {
          "active": true,
          "name": "shipping_email",
          "req": false,
          "type": "`$STRING`",
          "index$": 64
        },
        {
          "active": true,
          "name": "shipping_first_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 65
        },
        {
          "active": true,
          "name": "shipping_last_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 66
        },
        {
          "active": true,
          "name": "shipping_phone",
          "req": false,
          "type": "`$STRING`",
          "index$": 67
        },
        {
          "active": true,
          "name": "status",
          "req": false,
          "type": "`$STRING`",
          "index$": 68
        },
        {
          "active": true,
          "name": "success_web_hook_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 69
        },
        {
          "active": true,
          "name": "tag",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 70
        },
        {
          "active": true,
          "name": "tag_id",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 71
        },
        {
          "active": true,
          "name": "title",
          "req": false,
          "type": "`$STRING`",
          "index$": 72
        },
        {
          "active": true,
          "name": "tokenised_card",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 73
        },
        {
          "active": true,
          "name": "transaction",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 74
        },
        {
          "active": true,
          "name": "use_hosted_payment_page",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 75
        }
      ],
      "name": "payment",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/paymentrequests",
              "parts": [
                "api",
                "v1",
                "paymentrequests"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_event",
                    "orig": "include_event",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/paymentrequests/{id}",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id",
                  "include_event"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "order_id",
                    "orig": "order_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/paymentrequests/getbyorderid/{orderID}",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "getbyorderid",
                "{order_id}"
              ],
              "rename": {
                "param": {
                  "orderID": "order_id"
                }
              },
              "select": {
                "exist": [
                  "order_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "load"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/paymentrequests/{id}",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": [
          [
            "getbyorderid"
          ]
        ]
      }
    },
    "payment_account": {
      "fields": [
        {
          "active": true,
          "name": "account_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "account_supplier_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "available_balance",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "available_balance_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 3
        },
        {
          "active": true,
          "name": "balance",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 4
        },
        {
          "active": true,
          "name": "balance_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 5
        },
        {
          "active": true,
          "name": "bank_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "consent_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "created_by",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 8
        },
        {
          "active": true,
          "name": "created_by_display_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "default_payment_rail",
          "req": false,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "display_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        },
        {
          "active": true,
          "name": "expiry_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 13
        },
        {
          "active": true,
          "name": "external_account_icon",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "identifier",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 16
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 17
        },
        {
          "active": true,
          "name": "is_archived",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 18
        },
        {
          "active": true,
          "name": "is_connected_account",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 19
        },
        {
          "active": true,
          "name": "is_default",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 20
        },
        {
          "active": true,
          "name": "is_trust_account",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 21
        },
        {
          "active": true,
          "name": "is_virtual",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 22
        },
        {
          "active": true,
          "name": "last_transaction",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 23
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 24
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 25
        },
        {
          "active": true,
          "name": "merchant_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 26
        },
        {
          "active": true,
          "name": "physical_account_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 27
        },
        {
          "active": true,
          "name": "rule",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 28
        },
        {
          "active": true,
          "name": "submitted_payouts_balance",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 29
        },
        {
          "active": true,
          "name": "submitted_payouts_balance_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 30
        },
        {
          "active": true,
          "name": "summary",
          "req": false,
          "type": "`$STRING`",
          "index$": 31
        },
        {
          "active": true,
          "name": "supplier_sepa_instant_status",
          "req": false,
          "type": "`$STRING`",
          "index$": 32
        },
        {
          "active": true,
          "name": "xero_bank_feed_connection_status",
          "req": false,
          "type": "`$STRING`",
          "index$": 33
        },
        {
          "active": true,
          "name": "xero_bank_feed_last_synced_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 34
        },
        {
          "active": true,
          "name": "xero_bank_feed_sync_last_failed_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 35
        },
        {
          "active": true,
          "name": "xero_bank_feed_sync_last_failure_reason",
          "req": false,
          "type": "`$STRING`",
          "index$": 36
        },
        {
          "active": true,
          "name": "xero_bank_feed_sync_status",
          "req": false,
          "type": "`$STRING`",
          "index$": 37
        },
        {
          "active": true,
          "name": "xero_unsynchronised_transactions_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 38
        }
      ],
      "name": "payment_account",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "connected_account",
                    "orig": "connected_account",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_child_merchant",
                    "orig": "include_child_merchant",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "only_archived",
                    "orig": "only_archived",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "only_connect_account",
                    "orig": "only_connect_account",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/accounts/paged",
              "parts": [
                "api",
                "v1",
                "accounts",
                "paged"
              ],
              "select": {
                "exist": [
                  "connected_account",
                  "currency",
                  "include_archived",
                  "include_child_merchant",
                  "merchant_id",
                  "only_archived",
                  "only_connect_account",
                  "page_number",
                  "page_size",
                  "search",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "example": 1,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/accounts/{accountID}/virtual",
              "parts": [
                "api",
                "v1",
                "accounts",
                "{account_id}",
                "virtual"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id"
                }
              },
              "select": {
                "exist": [
                  "account_id",
                  "page_number",
                  "page_size"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": [
          [
            "account"
          ]
        ]
      }
    },
    "payment_account_minimal": {
      "fields": [
        {
          "active": true,
          "name": "account_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "available_balance",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 1
        },
        {
          "active": true,
          "name": "balance",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "balance_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 3
        },
        {
          "active": true,
          "name": "currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "identifier",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 6
        },
        {
          "active": true,
          "name": "is_archived",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 7
        },
        {
          "active": true,
          "name": "is_connected_account",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 8
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "submitted_payouts_balance",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 10
        }
      ],
      "name": "payment_account_minimal",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "connected_account",
                    "orig": "connected_account",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "only_archived",
                    "orig": "only_archived",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "only_connect_account",
                    "orig": "only_connect_account",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "example": 1,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 10,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/accounts/minimal",
              "parts": [
                "api",
                "v1",
                "accounts",
                "minimal"
              ],
              "select": {
                "exist": [
                  "connected_account",
                  "currency",
                  "include_archived",
                  "merchant_id",
                  "only_archived",
                  "only_connect_account",
                  "page_number",
                  "page_size",
                  "search",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "payment_initiation": {
      "fields": [
        {
          "active": true,
          "name": "payment_initiation_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "payment_request_callback_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "payment_request_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "redirect_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "response_type",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "specific_error_message",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        }
      ],
      "name": "payment_initiation",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/paymentrequests/{id}/pisp",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "pisp"
              ],
              "rename": {
                "param": {
                  "id": "paymentrequest_id"
                }
              },
              "select": {
                "exist": [
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        }
      },
      "relations": {
        "ancestors": [
          [
            "paymentrequest"
          ]
        ]
      }
    },
    "payment_request": {
      "fields": [
        {
          "active": true,
          "name": "address",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "amount",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 1
        },
        {
          "active": true,
          "name": "amount_pending",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "amount_received",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 3
        },
        {
          "active": true,
          "name": "amount_refunded",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 4
        },
        {
          "active": true,
          "name": "auto_send_receipt",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 5
        },
        {
          "active": true,
          "name": "base_origin_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "callback_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "card_authorize_only",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 8
        },
        {
          "active": true,
          "name": "card_create_token",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 9
        },
        {
          "active": true,
          "name": "card_create_token_mode",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "card_ignore_cvn",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 11
        },
        {
          "active": true,
          "name": "card_processor_merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        },
        {
          "active": true,
          "name": "card_stripe_payment_intent_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 13
        },
        {
          "active": true,
          "name": "card_stripe_payment_intent_secret",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "created_by_user",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 15
        },
        {
          "active": true,
          "name": "currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 16
        },
        {
          "active": true,
          "name": "custom_field",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 17
        },
        {
          "active": true,
          "name": "customer_email_address",
          "req": false,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "customer_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 19
        },
        {
          "active": true,
          "name": "customer_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 20
        },
        {
          "active": true,
          "name": "description",
          "req": false,
          "type": "`$STRING`",
          "index$": 21
        },
        {
          "active": true,
          "name": "destination_account",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 22
        },
        {
          "active": true,
          "name": "direct_debit_payment",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 23
        },
        {
          "active": true,
          "name": "do_simulate_settlement_failure",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 24
        },
        {
          "active": true,
          "name": "due_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 25
        },
        {
          "active": true,
          "name": "error_description",
          "req": false,
          "type": "`$STRING`",
          "index$": 26
        },
        {
          "active": true,
          "name": "event",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 27
        },
        {
          "active": true,
          "name": "failed_payment_request",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 28
        },
        {
          "active": true,
          "name": "failure_callback_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 29
        },
        {
          "active": true,
          "name": "field_display_setting",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 30
        },
        {
          "active": true,
          "name": "formatted_amount",
          "req": false,
          "type": "`$STRING`",
          "index$": 31
        },
        {
          "active": true,
          "name": "hosted_pay_checkout_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 32
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 33
        },
        {
          "active": true,
          "name": "ignore_address_verification",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 34
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 35
        },
        {
          "active": true,
          "name": "inserted_sortable",
          "req": false,
          "type": "`$STRING`",
          "index$": 36
        },
        {
          "active": true,
          "name": "institution",
          "req": false,
          "type": "`$STRING`",
          "index$": 37
        },
        {
          "active": true,
          "name": "is_archived",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 38
        },
        {
          "active": true,
          "name": "jwk",
          "req": false,
          "type": "`$STRING`",
          "index$": 39
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 40
        },
        {
          "active": true,
          "name": "lightning_invoice",
          "req": false,
          "type": "`$STRING`",
          "index$": 41
        },
        {
          "active": true,
          "name": "lightning_invoice_expires_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 42
        },
        {
          "active": true,
          "name": "merchant_direct_debit_mandate_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 43
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 44
        },
        {
          "active": true,
          "name": "merchant_token_description",
          "req": false,
          "type": "`$STRING`",
          "index$": 45
        },
        {
          "active": true,
          "name": "notification_email_address",
          "req": false,
          "type": "`$STRING`",
          "index$": 46
        },
        {
          "active": true,
          "name": "notification_role_i_d",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 47
        },
        {
          "active": true,
          "name": "order_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 48
        },
        {
          "active": true,
          "name": "partial_payment_method",
          "req": false,
          "type": "`$STRING`",
          "index$": 49
        },
        {
          "active": true,
          "name": "partial_payment_step",
          "req": false,
          "type": "`$STRING`",
          "index$": 50
        },
        {
          "active": true,
          "name": "payment_attempt",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 51
        },
        {
          "active": true,
          "name": "payment_initiation_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 52
        },
        {
          "active": true,
          "name": "payment_method",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 53
        },
        {
          "active": true,
          "name": "payment_processor",
          "req": false,
          "type": "`$STRING`",
          "index$": 54
        },
        {
          "active": true,
          "name": "payment_request",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 55
        },
        {
          "active": true,
          "name": "payrun_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 56
        },
        {
          "active": true,
          "name": "pisp_account_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 57
        },
        {
          "active": true,
          "name": "priority_bank_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 58
        },
        {
          "active": true,
          "name": "result",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 59
        },
        {
          "active": true,
          "name": "sandbox_settle_delay_in_second",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 60
        },
        {
          "active": true,
          "name": "shipping_address",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 61
        },
        {
          "active": true,
          "name": "status",
          "req": false,
          "type": "`$STRING`",
          "index$": 62
        },
        {
          "active": true,
          "name": "success_web_hook_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 63
        },
        {
          "active": true,
          "name": "tag",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 64
        },
        {
          "active": true,
          "name": "title",
          "req": false,
          "type": "`$STRING`",
          "index$": 65
        },
        {
          "active": true,
          "name": "tokenised_card",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 66
        },
        {
          "active": true,
          "name": "transaction",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 67
        },
        {
          "active": true,
          "name": "use_hosted_payment_page",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 68
        }
      ],
      "name": "payment_request",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "mandate_id",
                    "orig": "mandate_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "submit_after",
                    "orig": "submit_after",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/paymentrequests/{id}/directdebit",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "directdebit"
              ],
              "rename": {
                "param": {
                  "id": "paymentrequest_id"
                }
              },
              "select": {
                "exist": [
                  "mandate_id",
                  "paymentrequest_id",
                  "submit_after"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/paymentrequests/batchcreate",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "batchcreate"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "payment_method",
                    "orig": "payment_method",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": "All",
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/paymentrequests",
              "parts": [
                "api",
                "v1",
                "paymentrequests"
              ],
              "select": {
                "exist": [
                  "currency",
                  "from_date",
                  "include_archived",
                  "max_amount",
                  "merchant_id",
                  "min_amount",
                  "page",
                  "payment_method",
                  "search",
                  "size",
                  "sort",
                  "status",
                  "tag",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "payment_method",
                    "orig": "payment_method",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": "All",
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/paymentrequests/export",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "export"
              ],
              "select": {
                "exist": [
                  "currency",
                  "from_date",
                  "include_archived",
                  "max_amount",
                  "merchant_id",
                  "min_amount",
                  "page",
                  "payment_method",
                  "search",
                  "size",
                  "sort",
                  "status",
                  "tag",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/paymentrequests/{id}/receipt",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "receipt"
              ],
              "rename": {
                "param": {
                  "id": "paymentrequest_id"
                }
              },
              "select": {
                "exist": [
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/paymentrequests/{id}",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/paymentrequests/{id}/pisp/sandboxcallback",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "pisp",
                "sandboxcallback"
              ],
              "rename": {
                "param": {
                  "id": "paymentrequest_id"
                }
              },
              "select": {
                "exist": [
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": [
          [
            "paymentrequest"
          ]
        ]
      }
    },
    "payment_request_event": {
      "fields": [
        {
          "active": true,
          "name": "amount",
          "req": true,
          "type": "`$NUMBER`",
          "index$": 0
        },
        {
          "active": true,
          "name": "apple_pay_transaction_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "card_authorization_response_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "card_expiry_month",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 3
        },
        {
          "active": true,
          "name": "card_expiry_year",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 4
        },
        {
          "active": true,
          "name": "card_issuer",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "card_issuer_country",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "card_last_four_digit",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "card_request_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 8
        },
        {
          "active": true,
          "name": "card_scheme",
          "req": false,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "card_token_customer_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "card_transaction_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        },
        {
          "active": true,
          "name": "direct_debit_payment_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 13
        },
        {
          "active": true,
          "name": "direct_debit_payment_reference",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "drirect_debit_mandate_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "error_message",
          "req": false,
          "type": "`$STRING`",
          "index$": 16
        },
        {
          "active": true,
          "name": "error_reason",
          "req": false,
          "type": "`$STRING`",
          "index$": 17
        },
        {
          "active": true,
          "name": "event_type",
          "req": false,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 19
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 20
        },
        {
          "active": true,
          "name": "lightning_invoice",
          "req": false,
          "type": "`$STRING`",
          "index$": 21
        },
        {
          "active": true,
          "name": "lightning_r_hash",
          "req": false,
          "type": "`$STRING`",
          "index$": 22
        },
        {
          "active": true,
          "name": "origin_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 23
        },
        {
          "active": true,
          "name": "payment_method_type",
          "req": false,
          "type": "`$STRING`",
          "index$": 24
        },
        {
          "active": true,
          "name": "payment_processor_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 25
        },
        {
          "active": true,
          "name": "payment_request_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 26
        },
        {
          "active": true,
          "name": "pisp_bank_status",
          "req": false,
          "type": "`$STRING`",
          "index$": 27
        },
        {
          "active": true,
          "name": "pisp_payment_initiation_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 28
        },
        {
          "active": true,
          "name": "pisp_payment_institution_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 29
        },
        {
          "active": true,
          "name": "pisp_payment_service_provider_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 30
        },
        {
          "active": true,
          "name": "pisp_redirect_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 31
        },
        {
          "active": true,
          "name": "reconciled_transaction_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 32
        },
        {
          "active": true,
          "name": "refund_payout_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 33
        },
        {
          "active": true,
          "name": "status",
          "req": false,
          "type": "`$STRING`",
          "index$": 34
        },
        {
          "active": true,
          "name": "wallet_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 35
        }
      ],
      "name": "payment_request_event",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/paymentrequests/{id}/events",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "events"
              ],
              "rename": {
                "param": {
                  "id": "paymentrequest_id"
                }
              },
              "select": {
                "exist": [
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": [
          [
            "paymentrequest"
          ]
        ]
      }
    },
    "payment_request_metric": {
      "fields": [
        {
          "active": true,
          "name": "all",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 0
        },
        {
          "active": true,
          "name": "authorized",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 1
        },
        {
          "active": true,
          "name": "paid",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "partially_paid",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 3
        },
        {
          "active": true,
          "name": "total_amounts_by_currency",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 4
        },
        {
          "active": true,
          "name": "unpaid",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 5
        }
      ],
      "name": "payment_request_metric",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "payment_method",
                    "orig": "payment_method",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/paymentrequests/metrics",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "metrics"
              ],
              "select": {
                "exist": [
                  "currency",
                  "from_date",
                  "include_archived",
                  "max_amount",
                  "merchant_id",
                  "min_amount",
                  "payment_method",
                  "search",
                  "tag",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "payment_request_minimal": {
      "fields": [
        {
          "active": true,
          "name": "amount",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 0
        },
        {
          "active": true,
          "name": "amount_pending",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 1
        },
        {
          "active": true,
          "name": "amount_received",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "amount_refunded",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 3
        },
        {
          "active": true,
          "name": "callback_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "card_stripe_payment_intent_secret",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "country_code",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "custom_fields_to_display",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 8
        },
        {
          "active": true,
          "name": "description",
          "req": false,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "due_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "field_display_setting",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 11
        },
        {
          "active": true,
          "name": "google_pay_merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 13
        },
        {
          "active": true,
          "name": "jwk",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "merchant_logo_url_png",
          "req": false,
          "type": "`$STRING`",
          "index$": 16
        },
        {
          "active": true,
          "name": "merchant_logo_url_svg",
          "req": false,
          "type": "`$STRING`",
          "index$": 17
        },
        {
          "active": true,
          "name": "merchant_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "merchant_short_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 19
        },
        {
          "active": true,
          "name": "partial_payment_method",
          "req": false,
          "type": "`$STRING`",
          "index$": 20
        },
        {
          "active": true,
          "name": "payment_attempt",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 21
        },
        {
          "active": true,
          "name": "payment_methods_list",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 22
        },
        {
          "active": true,
          "name": "payment_processor",
          "req": false,
          "type": "`$STRING`",
          "index$": 23
        },
        {
          "active": true,
          "name": "payment_processor_key",
          "req": false,
          "type": "`$STRING`",
          "index$": 24
        },
        {
          "active": true,
          "name": "pisp_error",
          "req": false,
          "type": "`$STRING`",
          "index$": 25
        },
        {
          "active": true,
          "name": "priority_bank_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 26
        },
        {
          "active": true,
          "name": "status",
          "req": false,
          "type": "`$STRING`",
          "index$": 27
        },
        {
          "active": true,
          "name": "stripe_account_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 28
        },
        {
          "active": true,
          "name": "title",
          "req": false,
          "type": "`$STRING`",
          "index$": 29
        }
      ],
      "name": "payment_request_minimal",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/paymentrequests/{id}/minimal",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "minimal"
              ],
              "rename": {
                "param": {
                  "id": "paymentrequest_id"
                }
              },
              "select": {
                "exist": [
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": [
          [
            "paymentrequest"
          ]
        ]
      }
    },
    "payment_request_result": {
      "fields": [
        {
          "active": true,
          "name": "amount",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 0
        },
        {
          "active": true,
          "name": "amount_pending",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 1
        },
        {
          "active": true,
          "name": "amount_received",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "amount_refunded",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 3
        },
        {
          "active": true,
          "name": "currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "customer_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "payment",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 6
        },
        {
          "active": true,
          "name": "payment_request_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "pisp_authorization",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 8
        },
        {
          "active": true,
          "name": "requested_amount",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 9
        },
        {
          "active": true,
          "name": "result",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        }
      ],
      "name": "payment_request_result",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/paymentrequests/{id}/result",
              "parts": [
                "api",
                "v1",
                "paymentrequests",
                "{paymentrequest_id}",
                "result"
              ],
              "rename": {
                "param": {
                  "id": "paymentrequest_id"
                }
              },
              "select": {
                "exist": [
                  "paymentrequest_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": [
          [
            "paymentrequest"
          ]
        ]
      }
    },
    "payout": {
      "fields": [
        {
          "active": true,
          "name": "account_id",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "allow_incomplete",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 1
        },
        {
          "active": true,
          "name": "amount",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "amount_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 3
        },
        {
          "active": true,
          "name": "approve_payout_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "approver_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "authentication_method",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 6
        },
        {
          "active": true,
          "name": "authorisation",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 7
        },
        {
          "active": true,
          "name": "authorisers_completed_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 8
        },
        {
          "active": true,
          "name": "authorisers_required_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 9
        },
        {
          "active": true,
          "name": "batch_payout_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "beneficiary",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 11
        },
        {
          "active": true,
          "name": "beneficiary_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        },
        {
          "active": true,
          "name": "can_authorise",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 13
        },
        {
          "active": true,
          "name": "can_process",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 14
        },
        {
          "active": true,
          "name": "can_update",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 15
        },
        {
          "active": true,
          "name": "charge_bearer",
          "req": false,
          "type": "`$STRING`",
          "index$": 16
        },
        {
          "active": true,
          "name": "created_by",
          "req": false,
          "type": "`$STRING`",
          "index$": 17
        },
        {
          "active": true,
          "name": "created_by_email_address",
          "req": false,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "currency",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$STRING`",
          "index$": 19
        },
        {
          "active": true,
          "name": "current_user_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 20
        },
        {
          "active": true,
          "name": "description",
          "req": false,
          "type": "`$STRING`",
          "index$": 21
        },
        {
          "active": true,
          "name": "destination",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 22
        },
        {
          "active": true,
          "name": "document",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 23
        },
        {
          "active": true,
          "name": "event",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 24
        },
        {
          "active": true,
          "name": "failed_payout",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 25
        },
        {
          "active": true,
          "name": "formatted_amount",
          "req": false,
          "type": "`$STRING`",
          "index$": 26
        },
        {
          "active": true,
          "name": "formatted_fx_destination_amount",
          "req": false,
          "type": "`$STRING`",
          "index$": 27
        },
        {
          "active": true,
          "name": "formatted_schedule",
          "req": false,
          "type": "`$STRING`",
          "index$": 28
        },
        {
          "active": true,
          "name": "formatted_schedule_day_only",
          "req": false,
          "type": "`$STRING`",
          "index$": 29
        },
        {
          "active": true,
          "name": "formatted_source_account_available_balance",
          "req": false,
          "type": "`$STRING`",
          "index$": 30
        },
        {
          "active": true,
          "name": "fx_destination_amount",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 31
        },
        {
          "active": true,
          "name": "fx_destination_amount_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 32
        },
        {
          "active": true,
          "name": "fx_destination_currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 33
        },
        {
          "active": true,
          "name": "fx_quote_expires_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 34
        },
        {
          "active": true,
          "name": "fx_quote_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 35
        },
        {
          "active": true,
          "name": "fx_rate",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 36
        },
        {
          "active": true,
          "name": "fx_use_destination_amount",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 37
        },
        {
          "active": true,
          "name": "has_current_user_authorised",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 38
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 39
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 40
        },
        {
          "active": true,
          "name": "invoice_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 41
        },
        {
          "active": true,
          "name": "is_archived",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 42
        },
        {
          "active": true,
          "name": "is_failed",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 43
        },
        {
          "active": true,
          "name": "is_settled",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 44
        },
        {
          "active": true,
          "name": "is_submitted",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 45
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 46
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 47
        },
        {
          "active": true,
          "name": "merchant_token_description",
          "req": false,
          "type": "`$STRING`",
          "index$": 48
        },
        {
          "active": true,
          "name": "nonce",
          "req": false,
          "type": "`$STRING`",
          "index$": 49
        },
        {
          "active": true,
          "name": "payment_processor",
          "req": false,
          "type": "`$STRING`",
          "index$": 50
        },
        {
          "active": true,
          "name": "payment_rail",
          "req": false,
          "type": "`$STRING`",
          "index$": 51
        },
        {
          "active": true,
          "name": "payout",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 52
        },
        {
          "active": true,
          "name": "payrun_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 53
        },
        {
          "active": true,
          "name": "payrun_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 54
        },
        {
          "active": true,
          "name": "reason",
          "req": false,
          "type": "`$STRING`",
          "index$": 55
        },
        {
          "active": true,
          "name": "rule",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 56
        },
        {
          "active": true,
          "name": "schedule_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 57
        },
        {
          "active": true,
          "name": "scheduled",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 58
        },
        {
          "active": true,
          "name": "source_account_available_balance",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 59
        },
        {
          "active": true,
          "name": "source_account_available_balance_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 60
        },
        {
          "active": true,
          "name": "source_account_bic",
          "req": false,
          "type": "`$STRING`",
          "index$": 61
        },
        {
          "active": true,
          "name": "source_account_currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 62
        },
        {
          "active": true,
          "name": "source_account_iban",
          "req": false,
          "type": "`$STRING`",
          "index$": 63
        },
        {
          "active": true,
          "name": "source_account_identifier",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 64
        },
        {
          "active": true,
          "name": "source_account_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 65
        },
        {
          "active": true,
          "name": "source_account_number",
          "req": false,
          "type": "`$STRING`",
          "index$": 66
        },
        {
          "active": true,
          "name": "source_account_sortcode",
          "req": false,
          "type": "`$STRING`",
          "index$": 67
        },
        {
          "active": true,
          "name": "status",
          "req": false,
          "type": "`$STRING`",
          "index$": 68
        },
        {
          "active": true,
          "name": "tag",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 69
        },
        {
          "active": true,
          "name": "tag_id",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 70
        },
        {
          "active": true,
          "name": "their_reference",
          "req": false,
          "type": "`$STRING`",
          "index$": 71
        },
        {
          "active": true,
          "name": "topup_payrun_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 72
        },
        {
          "active": true,
          "name": "transacted_amount",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 73
        },
        {
          "active": true,
          "name": "transacted_fx_amount",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 74
        },
        {
          "active": true,
          "name": "transacted_fx_rate",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 75
        },
        {
          "active": true,
          "name": "type",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$STRING`",
          "index$": 76
        },
        {
          "active": true,
          "name": "user_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 77
        },
        {
          "active": true,
          "name": "your_reference",
          "req": false,
          "type": "`$STRING`",
          "index$": 78
        }
      ],
      "name": "payout",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/payouts/batch/submit/{id}",
              "parts": [
                "api",
                "v1",
                "payouts",
                "batch",
                "submit",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/payouts/submit/{id}",
              "parts": [
                "api",
                "v1",
                "payouts",
                "submit",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/payouts",
              "parts": [
                "api",
                "v1",
                "payouts"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            },
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/payouts/batchcreate",
              "parts": [
                "api",
                "v1",
                "payouts",
                "batchcreate"
              ],
              "select": {
                "$action": "batchcreate"
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 3
            },
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/payouts/send",
              "parts": [
                "api",
                "v1",
                "payouts",
                "send"
              ],
              "select": {
                "$action": "send"
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 4
            },
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/payouts/sendbeneficiary",
              "parts": [
                "api",
                "v1",
                "payouts",
                "sendbeneficiary"
              ],
              "select": {
                "$action": "sendbeneficiary"
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 5
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/payouts",
              "parts": [
                "api",
                "v1",
                "payouts"
              ],
              "select": {
                "exist": [
                  "currency",
                  "from_date",
                  "include_archived",
                  "max_amount",
                  "merchant_id",
                  "min_amount",
                  "page_number",
                  "page_size",
                  "search",
                  "sort",
                  "status",
                  "tag",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/accounts/{accountID}/payouts",
              "parts": [
                "api",
                "v1",
                "accounts",
                "{account_id}",
                "payouts"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id"
                }
              },
              "select": {
                "exist": [
                  "account_id",
                  "currency",
                  "from_date",
                  "max_amount",
                  "min_amount",
                  "page_number",
                  "page_size",
                  "search",
                  "sort",
                  "status",
                  "tag",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/payouts",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "payouts"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "currency",
                  "from_date",
                  "max_amount",
                  "merchant_id",
                  "min_amount",
                  "page_number",
                  "page_size",
                  "search",
                  "sort",
                  "status",
                  "tag",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/payouts/export",
              "parts": [
                "api",
                "v1",
                "payouts",
                "export"
              ],
              "select": {
                "$action": "export",
                "exist": [
                  "currency",
                  "from_date",
                  "include_archived",
                  "max_amount",
                  "merchant_id",
                  "min_amount",
                  "page_number",
                  "page_size",
                  "search",
                  "sort",
                  "status",
                  "tag",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "amount",
                    "orig": "amount",
                    "reqd": true,
                    "type": "`$NUMBER`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "destination",
                    "orig": "destination",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "source",
                    "orig": "source",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 2
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/payouts/fxquote/{source}/{destination}/{amount}",
              "parts": [
                "api",
                "v1",
                "payouts",
                "fxquote",
                "{source}",
                "{destination}",
                "{amount}"
              ],
              "select": {
                "exist": [
                  "amount",
                  "destination",
                  "source"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/payouts/{id}",
              "parts": [
                "api",
                "v1",
                "payouts",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/payouts/{id}/proof",
              "parts": [
                "api",
                "v1",
                "payouts",
                "{id}",
                "proof"
              ],
              "select": {
                "$action": "proof",
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 3
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/payouts/{id}",
              "parts": [
                "api",
                "v1",
                "payouts",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {},
              "method": "DELETE",
              "orig": "/api/v1/payouts/batchdelete",
              "parts": [
                "api",
                "v1",
                "payouts",
                "batchdelete"
              ],
              "select": {
                "$action": "batchdelete"
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/payouts/cancel/{id}",
              "parts": [
                "api",
                "v1",
                "payouts",
                "cancel",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/payouts/reject/{id}",
              "parts": [
                "api",
                "v1",
                "payouts",
                "reject",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/payouts/{id}",
              "parts": [
                "api",
                "v1",
                "payouts",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": [
          [
            "account"
          ],
          [
            "merchant"
          ],
          [
            "fxquote"
          ]
        ]
      }
    },
    "payout_keyset_page": {
      "fields": [
        {
          "active": true,
          "name": "account_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "amount",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 1
        },
        {
          "active": true,
          "name": "amount_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "approve_payout_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "approver_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "authentication_method",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 5
        },
        {
          "active": true,
          "name": "authorisation",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 6
        },
        {
          "active": true,
          "name": "authorisers_completed_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 7
        },
        {
          "active": true,
          "name": "authorisers_required_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 8
        },
        {
          "active": true,
          "name": "batch_payout_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "beneficiary",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 10
        },
        {
          "active": true,
          "name": "can_authorise",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 11
        },
        {
          "active": true,
          "name": "can_process",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 12
        },
        {
          "active": true,
          "name": "can_update",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 13
        },
        {
          "active": true,
          "name": "charge_bearer",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "created_by",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "created_by_email_address",
          "req": false,
          "type": "`$STRING`",
          "index$": 16
        },
        {
          "active": true,
          "name": "currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 17
        },
        {
          "active": true,
          "name": "current_user_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "description",
          "req": false,
          "type": "`$STRING`",
          "index$": 19
        },
        {
          "active": true,
          "name": "destination",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 20
        },
        {
          "active": true,
          "name": "document",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 21
        },
        {
          "active": true,
          "name": "event",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 22
        },
        {
          "active": true,
          "name": "formatted_amount",
          "req": false,
          "type": "`$STRING`",
          "index$": 23
        },
        {
          "active": true,
          "name": "formatted_fx_destination_amount",
          "req": false,
          "type": "`$STRING`",
          "index$": 24
        },
        {
          "active": true,
          "name": "formatted_schedule",
          "req": false,
          "type": "`$STRING`",
          "index$": 25
        },
        {
          "active": true,
          "name": "formatted_schedule_day_only",
          "req": false,
          "type": "`$STRING`",
          "index$": 26
        },
        {
          "active": true,
          "name": "formatted_source_account_available_balance",
          "req": false,
          "type": "`$STRING`",
          "index$": 27
        },
        {
          "active": true,
          "name": "fx_destination_amount",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 28
        },
        {
          "active": true,
          "name": "fx_destination_amount_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 29
        },
        {
          "active": true,
          "name": "fx_destination_currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 30
        },
        {
          "active": true,
          "name": "fx_quote_expires_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 31
        },
        {
          "active": true,
          "name": "fx_quote_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 32
        },
        {
          "active": true,
          "name": "fx_rate",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 33
        },
        {
          "active": true,
          "name": "fx_use_destination_amount",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 34
        },
        {
          "active": true,
          "name": "has_current_user_authorised",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 35
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 36
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 37
        },
        {
          "active": true,
          "name": "invoice_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 38
        },
        {
          "active": true,
          "name": "is_archived",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 39
        },
        {
          "active": true,
          "name": "is_failed",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 40
        },
        {
          "active": true,
          "name": "is_settled",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 41
        },
        {
          "active": true,
          "name": "is_submitted",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 42
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 43
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 44
        },
        {
          "active": true,
          "name": "merchant_token_description",
          "req": false,
          "type": "`$STRING`",
          "index$": 45
        },
        {
          "active": true,
          "name": "nonce",
          "req": false,
          "type": "`$STRING`",
          "index$": 46
        },
        {
          "active": true,
          "name": "payment_processor",
          "req": false,
          "type": "`$STRING`",
          "index$": 47
        },
        {
          "active": true,
          "name": "payment_rail",
          "req": false,
          "type": "`$STRING`",
          "index$": 48
        },
        {
          "active": true,
          "name": "payrun_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 49
        },
        {
          "active": true,
          "name": "payrun_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 50
        },
        {
          "active": true,
          "name": "rule",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 51
        },
        {
          "active": true,
          "name": "schedule_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 52
        },
        {
          "active": true,
          "name": "scheduled",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 53
        },
        {
          "active": true,
          "name": "source_account_available_balance",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 54
        },
        {
          "active": true,
          "name": "source_account_available_balance_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 55
        },
        {
          "active": true,
          "name": "source_account_bic",
          "req": false,
          "type": "`$STRING`",
          "index$": 56
        },
        {
          "active": true,
          "name": "source_account_currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 57
        },
        {
          "active": true,
          "name": "source_account_iban",
          "req": false,
          "type": "`$STRING`",
          "index$": 58
        },
        {
          "active": true,
          "name": "source_account_identifier",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 59
        },
        {
          "active": true,
          "name": "source_account_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 60
        },
        {
          "active": true,
          "name": "source_account_number",
          "req": false,
          "type": "`$STRING`",
          "index$": 61
        },
        {
          "active": true,
          "name": "source_account_sortcode",
          "req": false,
          "type": "`$STRING`",
          "index$": 62
        },
        {
          "active": true,
          "name": "status",
          "req": false,
          "type": "`$STRING`",
          "index$": 63
        },
        {
          "active": true,
          "name": "tag",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 64
        },
        {
          "active": true,
          "name": "their_reference",
          "req": false,
          "type": "`$STRING`",
          "index$": 65
        },
        {
          "active": true,
          "name": "topup_payrun_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 66
        },
        {
          "active": true,
          "name": "transacted_amount",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 67
        },
        {
          "active": true,
          "name": "transacted_fx_amount",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 68
        },
        {
          "active": true,
          "name": "transacted_fx_rate",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 69
        },
        {
          "active": true,
          "name": "type",
          "req": false,
          "type": "`$STRING`",
          "index$": 70
        },
        {
          "active": true,
          "name": "user_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 71
        },
        {
          "active": true,
          "name": "your_reference",
          "req": false,
          "type": "`$STRING`",
          "index$": 72
        }
      ],
      "name": "payout_keyset_page",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date_utc",
                    "orig": "from_date_utc",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/accounts/{accountID}/payouts/failed",
              "parts": [
                "api",
                "v1",
                "accounts",
                "{account_id}",
                "payouts",
                "failed"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id"
                }
              },
              "select": {
                "exist": [
                  "account_id",
                  "from_date_utc",
                  "page_size"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date_utc",
                    "orig": "from_date_utc",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/payouts/failed",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "payouts",
                "failed"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "from_date_utc",
                  "merchant_id",
                  "page_size"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date_utc",
                    "orig": "from_date_utc",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/payouts/{merchantID}/failed",
              "parts": [
                "api",
                "v1",
                "payouts",
                "{merchant_id}",
                "failed"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "from_date_utc",
                  "merchant_id",
                  "page_size"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": [
          [
            "account"
          ],
          [
            "merchant"
          ],
          [
            "payout"
          ]
        ]
      }
    },
    "payout_metric": {
      "fields": [
        {
          "active": true,
          "name": "all",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 0
        },
        {
          "active": true,
          "name": "failed",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 1
        },
        {
          "active": true,
          "name": "in_progress",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "paid",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 3
        },
        {
          "active": true,
          "name": "pending_approval",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 4
        },
        {
          "active": true,
          "name": "scheduled",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 5
        },
        {
          "active": true,
          "name": "total_amounts_by_currency",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 6
        }
      ],
      "name": "payout_metric",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/payouts/metrics",
              "parts": [
                "api",
                "v1",
                "payouts",
                "metrics"
              ],
              "select": {
                "exist": [
                  "currency",
                  "from_date",
                  "include_archived",
                  "max_amount",
                  "merchant_id",
                  "min_amount",
                  "search",
                  "tag",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "payrun": {
      "fields": [
        {
          "active": true,
          "name": "authorisation",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "authorisation_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "authorisers_completed_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "authorisers_required_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 3
        },
        {
          "active": true,
          "name": "batch_payout_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "can_authorise",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 5
        },
        {
          "active": true,
          "name": "can_delete",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 6
        },
        {
          "active": true,
          "name": "can_edit",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 7
        },
        {
          "active": true,
          "name": "event",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 8
        },
        {
          "active": true,
          "name": "has_current_user_authorised",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 9
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "invoice",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 12
        },
        {
          "active": true,
          "name": "invoices_minimal",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 13
        },
        {
          "active": true,
          "name": "is_archived",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 14
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "last_updated_by",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 16
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 17
        },
        {
          "active": true,
          "name": "name",
          "req": false,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "nonce",
          "req": false,
          "type": "`$STRING`",
          "index$": 19
        },
        {
          "active": true,
          "name": "note",
          "req": false,
          "type": "`$STRING`",
          "index$": 20
        },
        {
          "active": true,
          "name": "payment",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 21
        },
        {
          "active": true,
          "name": "payout",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 22
        },
        {
          "active": true,
          "name": "payouts_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 23
        },
        {
          "active": true,
          "name": "reason",
          "req": false,
          "type": "`$STRING`",
          "index$": 24
        },
        {
          "active": true,
          "name": "schedule_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 25
        },
        {
          "active": true,
          "name": "scheduled_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 26
        },
        {
          "active": true,
          "name": "source_account",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 27
        },
        {
          "active": true,
          "name": "status",
          "req": false,
          "type": "`$STRING`",
          "index$": 28
        },
        {
          "active": true,
          "name": "total_eur",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 29
        },
        {
          "active": true,
          "name": "total_gbp",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 30
        },
        {
          "active": true,
          "name": "total_usd",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 31
        }
      ],
      "name": "payrun",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/payruns/{id}/request-authorisation",
              "parts": [
                "api",
                "v1",
                "payruns",
                "{id}",
                "request-authorisation"
              ],
              "select": {
                "$action": "request_authorisation",
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/payruns/{id}/submit",
              "parts": [
                "api",
                "v1",
                "payruns",
                "{id}",
                "submit"
              ],
              "select": {
                "$action": "submit",
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/payruns/{merchantID}",
              "parts": [
                "api",
                "v1",
                "payruns",
                "{id}"
              ],
              "rename": {
                "param": {
                  "merchantID": "id"
                }
              },
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "only_archived",
                    "orig": "only_archived",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/payruns",
              "parts": [
                "api",
                "v1",
                "payruns"
              ],
              "select": {
                "exist": [
                  "from_date",
                  "merchant_id",
                  "only_archived",
                  "page_number",
                  "page_size",
                  "search",
                  "sort",
                  "status",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/payruns/{id}",
              "parts": [
                "api",
                "v1",
                "payruns",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/payruns/{id}",
              "parts": [
                "api",
                "v1",
                "payruns",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/payruns/{id}/archive",
              "parts": [
                "api",
                "v1",
                "payruns",
                "{id}",
                "archive"
              ],
              "select": {
                "$action": "archive",
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/payruns/{id}",
              "parts": [
                "api",
                "v1",
                "payruns",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/payruns/{id}/cancel",
              "parts": [
                "api",
                "v1",
                "payruns",
                "{id}",
                "cancel"
              ],
              "select": {
                "$action": "cancel",
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/payruns/{id}/reject",
              "parts": [
                "api",
                "v1",
                "payruns",
                "{id}",
                "reject"
              ],
              "select": {
                "$action": "reject",
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/payruns/{id}/unarchive",
              "parts": [
                "api",
                "v1",
                "payruns",
                "{id}",
                "unarchive"
              ],
              "select": {
                "$action": "unarchive",
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 3
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "report": {
      "fields": [],
      "name": "report",
      "op": {
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/reports/{id}/initiate",
              "parts": [
                "api",
                "v1",
                "reports",
                "{id}",
                "initiate"
              ],
              "select": {
                "$action": "initiate",
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "report_result": {
      "fields": [
        {
          "active": true,
          "name": "content",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "content_type",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "last_completed_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "report_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "report_type",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "statement_number",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 6
        }
      ],
      "name": "report_result",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "statement_number",
                    "reqd": true,
                    "type": "`$INTEGER`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "report_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/reports/{id}/result/{statementNumber}",
              "parts": [
                "api",
                "v1",
                "reports",
                "{report_id}",
                "result",
                "{id}"
              ],
              "rename": {
                "param": {
                  "id": "report_id",
                  "statementNumber": "id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "report_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        }
      },
      "relations": {
        "ancestors": [
          [
            "report"
          ]
        ]
      }
    },
    "role": {
      "fields": [
        {
          "active": true,
          "name": "failed_role",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 0
        },
        {
          "active": true,
          "name": "role",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 1
        }
      ],
      "name": "role",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/merchants/{merchantID}/roles/batchcreate",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "roles",
                "batchcreate"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "$action": "batchcreate",
                "exist": [
                  "merchant_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        }
      },
      "relations": {
        "ancestors": [
          [
            "merchant"
          ]
        ]
      }
    },
    "rule": {
      "fields": [
        {
          "active": true,
          "name": "account",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 0
        },
        {
          "active": true,
          "name": "account_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "approve_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "approver_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "authentication_method",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 4
        },
        {
          "active": true,
          "name": "authorisation",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 5
        },
        {
          "active": true,
          "name": "authorisers_completed_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 6
        },
        {
          "active": true,
          "name": "authorisers_required_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 7
        },
        {
          "active": true,
          "name": "can_authorise",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 8
        },
        {
          "active": true,
          "name": "created_by",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 9
        },
        {
          "active": true,
          "name": "description",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "end_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "has_current_user_authorised",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 12
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 13
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "is_disabled",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 15
        },
        {
          "active": true,
          "name": "last_executed_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 16
        },
        {
          "active": true,
          "name": "last_run_at_transaction_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 17
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 19
        },
        {
          "active": true,
          "name": "name",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$STRING`",
          "index$": 20
        },
        {
          "active": true,
          "name": "nonce",
          "req": true,
          "type": "`$STRING`",
          "index$": 21
        },
        {
          "active": true,
          "name": "on_approved_web_hook_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 22
        },
        {
          "active": true,
          "name": "on_execution_error_web_hook_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 23
        },
        {
          "active": true,
          "name": "on_execution_success_web_hook_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 24
        },
        {
          "active": true,
          "name": "start_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 25
        },
        {
          "active": true,
          "name": "status",
          "req": false,
          "type": "`$STRING`",
          "index$": 26
        },
        {
          "active": true,
          "name": "sweep_action",
          "op": {
            "create": {
              "req": true,
              "type": "`$OBJECT`"
            }
          },
          "req": false,
          "type": "`$OBJECT`",
          "index$": 27
        },
        {
          "active": true,
          "name": "time_zone_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 28
        },
        {
          "active": true,
          "name": "trigger_cron_expression",
          "req": false,
          "type": "`$STRING`",
          "index$": 29
        },
        {
          "active": true,
          "name": "trigger_on_pay_in",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 30
        },
        {
          "active": true,
          "name": "user_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 31
        },
        {
          "active": true,
          "name": "web_hook_secret",
          "req": false,
          "type": "`$STRING`",
          "index$": 32
        }
      ],
      "name": "rule",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/rules",
              "parts": [
                "api",
                "v1",
                "rules"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "archived_only",
                    "orig": "archived_only",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/rules",
              "parts": [
                "api",
                "v1",
                "rules"
              ],
              "select": {
                "exist": [
                  "archived_only",
                  "merchant_id",
                  "page",
                  "search",
                  "size",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/rules/{id}",
              "parts": [
                "api",
                "v1",
                "rules",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/rules/{id}",
              "parts": [
                "api",
                "v1",
                "rules",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/rules/{id}",
              "parts": [
                "api",
                "v1",
                "rules",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/rules/{id}/disable",
              "parts": [
                "api",
                "v1",
                "rules",
                "{id}",
                "disable"
              ],
              "select": {
                "$action": "disable",
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "rule_event": {
      "fields": [
        {
          "active": true,
          "name": "error_message",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "is_authorise_to_enable",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 3
        },
        {
          "active": true,
          "name": "message",
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "raw_response",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "rule_event_type",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "rule_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "user",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 8
        }
      ],
      "name": "rule_event",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "event_type",
                    "orig": "event_type",
                    "reqd": false,
                    "type": "`$ARRAY`"
                  },
                  {
                    "active": true,
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/rules/{id}/events",
              "parts": [
                "api",
                "v1",
                "rules",
                "{id}",
                "events"
              ],
              "select": {
                "exist": [
                  "event_type",
                  "id",
                  "page",
                  "size"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "tag": {
      "fields": [
        {
          "active": true,
          "name": "colour_hex",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "description",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": true,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "name",
          "req": true,
          "type": "`$STRING`",
          "index$": 4
        }
      ],
      "name": "tag",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/merchants/{merchantID}/tags",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "tags"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "merchant_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/tags",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "tags"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "merchant_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        }
      },
      "relations": {
        "ancestors": [
          [
            "merchant"
          ]
        ]
      }
    },
    "token": {
      "fields": [],
      "name": "token",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/tokens/authorise/{id}",
              "parts": [
                "api",
                "v1",
                "tokens",
                "authorise",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/tokens/{id}",
              "parts": [
                "api",
                "v1",
                "tokens",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "transaction": {
      "fields": [
        {
          "active": true,
          "name": "account_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "account_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "account_sequence_number",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "address_detail",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 3
        },
        {
          "active": true,
          "name": "amount",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 4
        },
        {
          "active": true,
          "name": "amount_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 5
        },
        {
          "active": true,
          "name": "balance",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 6
        },
        {
          "active": true,
          "name": "balance_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 7
        },
        {
          "active": true,
          "name": "booking_date_time",
          "req": false,
          "type": "`$STRING`",
          "index$": 8
        },
        {
          "active": true,
          "name": "charge_detail",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 9
        },
        {
          "active": true,
          "name": "content",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 10
        },
        {
          "active": true,
          "name": "counterparty",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 11
        },
        {
          "active": true,
          "name": "counterparty_summary",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        },
        {
          "active": true,
          "name": "currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 13
        },
        {
          "active": true,
          "name": "currency_exchange",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 14
        },
        {
          "active": true,
          "name": "date",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "description",
          "req": false,
          "type": "`$STRING`",
          "index$": 16
        },
        {
          "active": true,
          "name": "enrichment",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 17
        },
        {
          "active": true,
          "name": "fx_amount",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 18
        },
        {
          "active": true,
          "name": "fx_currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 19
        },
        {
          "active": true,
          "name": "fx_rate",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 20
        },
        {
          "active": true,
          "name": "gross_amount",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 21
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 22
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 23
        },
        {
          "active": true,
          "name": "iso_bank_transaction_code",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 24
        },
        {
          "active": true,
          "name": "merchant",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 25
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 26
        },
        {
          "active": true,
          "name": "page_number",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 27
        },
        {
          "active": true,
          "name": "page_size",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 28
        },
        {
          "active": true,
          "name": "payee_detail",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 29
        },
        {
          "active": true,
          "name": "payer_detail",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 30
        },
        {
          "active": true,
          "name": "payment_request_custom_field",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 31
        },
        {
          "active": true,
          "name": "payment_request_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 32
        },
        {
          "active": true,
          "name": "payout_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 33
        },
        {
          "active": true,
          "name": "proprietary_bank_transaction_code",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 34
        },
        {
          "active": true,
          "name": "raw_reference",
          "req": false,
          "type": "`$STRING`",
          "index$": 35
        },
        {
          "active": true,
          "name": "reference",
          "req": false,
          "type": "`$STRING`",
          "index$": 36
        },
        {
          "active": true,
          "name": "rule_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 37
        },
        {
          "active": true,
          "name": "statement_reference",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 38
        },
        {
          "active": true,
          "name": "status",
          "req": false,
          "type": "`$STRING`",
          "index$": 39
        },
        {
          "active": true,
          "name": "supplementary_data",
          "req": false,
          "type": "`$ANY`",
          "index$": 40
        },
        {
          "active": true,
          "name": "tag",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 41
        },
        {
          "active": true,
          "name": "their_reference",
          "req": false,
          "type": "`$STRING`",
          "index$": 42
        },
        {
          "active": true,
          "name": "total_page",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 43
        },
        {
          "active": true,
          "name": "total_size",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 44
        },
        {
          "active": true,
          "name": "transaction_amount",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 45
        },
        {
          "active": true,
          "name": "transaction_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 46
        },
        {
          "active": true,
          "name": "transaction_information",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 47
        },
        {
          "active": true,
          "name": "transaction_mutability",
          "req": false,
          "type": "`$STRING`",
          "index$": 48
        },
        {
          "active": true,
          "name": "type",
          "req": false,
          "type": "`$STRING`",
          "index$": 49
        },
        {
          "active": true,
          "name": "value_date_time",
          "req": false,
          "type": "`$STRING`",
          "index$": 50
        },
        {
          "active": true,
          "name": "virtual_iban",
          "req": false,
          "type": "`$STRING`",
          "index$": 51
        },
        {
          "active": true,
          "name": "your_reference",
          "req": false,
          "type": "`$STRING`",
          "index$": 52
        }
      ],
      "name": "transaction",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/transactions/{id}/tags",
              "parts": [
                "api",
                "v1",
                "transactions",
                "{id}",
                "tags"
              ],
              "select": {
                "$action": "tag",
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "credit_type",
                    "orig": "credit_type",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/accounts/{accountID}/transactions",
              "parts": [
                "api",
                "v1",
                "accounts",
                "{account_id}",
                "transactions"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id"
                }
              },
              "select": {
                "exist": [
                  "account_id",
                  "credit_type",
                  "from_date",
                  "max_amount",
                  "min_amount",
                  "page_number",
                  "page_size",
                  "search",
                  "sort",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "limit",
                    "orig": "limit",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "offset",
                    "orig": "offset",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort_inserted_ascending",
                    "orig": "sort_inserted_ascending",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "transaction_from",
                    "orig": "transaction_from",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "transaction_to",
                    "orig": "transaction_to",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/openbanking/transactions/{id}/{accountID}",
              "parts": [
                "api",
                "v1",
                "openbanking",
                "transactions",
                "{id}",
                "{account_id}"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id"
                }
              },
              "select": {
                "exist": [
                  "account_id",
                  "id",
                  "limit",
                  "offset",
                  "sort_inserted_ascending",
                  "transaction_from",
                  "transaction_to"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "example": false,
                    "kind": "query",
                    "name": "include_child_merchant",
                    "orig": "include_child_merchant",
                    "reqd": false,
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/transactions",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "transactions"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "from_date",
                  "include_child_merchant",
                  "merchant_id",
                  "page_number",
                  "page_size",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            },
            {
              "active": true,
              "args": {
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "credit_type",
                    "orig": "credit_type",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/transactions",
              "parts": [
                "api",
                "v1",
                "transactions"
              ],
              "select": {
                "exist": [
                  "credit_type",
                  "from_date",
                  "page_number",
                  "page_size",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 3
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "credit_type",
                    "orig": "credit_type",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/transactions/{accountID}/export",
              "parts": [
                "api",
                "v1",
                "transactions",
                "{account_id}",
                "export"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id"
                }
              },
              "select": {
                "$action": "export",
                "exist": [
                  "account_id",
                  "credit_type",
                  "from_date",
                  "max_amount",
                  "min_amount",
                  "page_number",
                  "page_size",
                  "search",
                  "sort",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "credit_type",
                    "orig": "credit_type",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "reqd": false,
                    "type": "`$NUMBER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/transactions/{accountID}",
              "parts": [
                "api",
                "v1",
                "transactions",
                "{id}"
              ],
              "rename": {
                "param": {
                  "accountID": "id"
                }
              },
              "select": {
                "exist": [
                  "credit_type",
                  "from_date",
                  "id",
                  "max_amount",
                  "min_amount",
                  "page_number",
                  "page_size",
                  "search",
                  "sort",
                  "to_date"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "example": 0,
                    "kind": "param",
                    "name": "sequence_number",
                    "orig": "sequence_number",
                    "reqd": true,
                    "type": "`$INTEGER`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "transaction_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/transactions/{accountID}/from/{sequenceNumber}",
              "parts": [
                "api",
                "v1",
                "transactions",
                "{transaction_id}",
                "from",
                "{sequence_number}"
              ],
              "rename": {
                "param": {
                  "accountID": "transaction_id",
                  "sequenceNumber": "sequence_number"
                }
              },
              "select": {
                "exist": [
                  "page_size",
                  "sequence_number",
                  "transaction_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/accounts/{accountID}/transactions/{id}",
              "parts": [
                "api",
                "v1",
                "accounts",
                "{account_id}",
                "transactions",
                "{id}"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id"
                }
              },
              "select": {
                "exist": [
                  "account_id",
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 3
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/transactions/detail/{id}",
              "parts": [
                "api",
                "v1",
                "transactions",
                "detail",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 4
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/transactions/{id}/proof",
              "parts": [
                "api",
                "v1",
                "transactions",
                "{id}",
                "proof"
              ],
              "select": {
                "$action": "proof",
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 5
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "kind": "query",
                    "name": "tag_id",
                    "orig": "tag_id",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/transactions/{id}/tag",
              "parts": [
                "api",
                "v1",
                "transactions",
                "{id}",
                "tag"
              ],
              "select": {
                "$action": "tag",
                "exist": [
                  "id",
                  "tag_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        }
      },
      "relations": {
        "ancestors": [
          [
            "account"
          ],
          [
            "merchant"
          ],
          [
            "transaction"
          ],
          [
            "transaction",
            "from"
          ]
        ]
      }
    },
    "user": {
      "fields": [
        {
          "active": true,
          "name": "client_session_timeout",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 0
        },
        {
          "active": true,
          "name": "email_address",
          "op": {
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "first_name",
          "op": {
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "last_name",
          "op": {
            "update": {
              "req": false,
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "passkey_added",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 5
        },
        {
          "active": true,
          "name": "permission",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 6
        },
        {
          "active": true,
          "name": "profile",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "roles_with_scope",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 8
        },
        {
          "active": true,
          "name": "two_factor_enabled",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 9
        },
        {
          "active": true,
          "name": "user_invite_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        }
      ],
      "name": "user",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "example": 1,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/user/{merchantID}/userspaged",
              "parts": [
                "api",
                "v1",
                "user",
                "{merchant_id}",
                "userspaged"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "$action": "userspaged",
                "exist": [
                  "merchant_id",
                  "page_number",
                  "page_size",
                  "search",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/users",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "users"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "merchant_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/metadata/whoami",
              "parts": [
                "api",
                "v1",
                "metadata",
                "whoami"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/metadata/whoamitrustedapp",
              "parts": [
                "api",
                "v1",
                "metadata",
                "whoamitrustedapp"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 3
            },
            {
              "active": true,
              "args": {},
              "method": "GET",
              "orig": "/api/v1/user",
              "parts": [
                "api",
                "v1",
                "user"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 4
            }
          ],
          "key$": "list"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/user/{id}",
              "parts": [
                "api",
                "v1",
                "user",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": [
          [
            "merchant"
          ],
          [
            "user"
          ]
        ]
      }
    },
    "user_invite": {
      "fields": [
        {
          "active": true,
          "name": "authorisation_status",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 0
        },
        {
          "active": true,
          "name": "failed_user_invite",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 1
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "initial_role_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "invitee_email_address",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$STRING`",
          "index$": 4
        },
        {
          "active": true,
          "name": "invitee_first_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "invitee_last_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "inviter_email_address",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "inviter_first_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 8
        },
        {
          "active": true,
          "name": "inviter_last_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "is_authorised",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 10
        },
        {
          "active": true,
          "name": "is_invitee_registered",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 11
        },
        {
          "active": true,
          "name": "last_invited",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 13
        },
        {
          "active": true,
          "name": "merchant_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "message",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "registration_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 16
        },
        {
          "active": true,
          "name": "send_invite_email",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 17
        },
        {
          "active": true,
          "name": "status",
          "req": false,
          "type": "`$STRING`",
          "index$": 18
        },
        {
          "active": true,
          "name": "user",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 19
        },
        {
          "active": true,
          "name": "user_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 20
        },
        {
          "active": true,
          "name": "user_invite",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 21
        }
      ],
      "name": "user_invite",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/userinvites/authorise/{id}",
              "parts": [
                "api",
                "v1",
                "userinvites",
                "authorise",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/userinvites",
              "parts": [
                "api",
                "v1",
                "userinvites"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            },
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/userinvites/batchcreate",
              "parts": [
                "api",
                "v1",
                "userinvites",
                "batchcreate"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 2
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ],
                "query": [
                  {
                    "active": true,
                    "example": 1,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "reqd": false,
                    "type": "`$INTEGER`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "reqd": false,
                    "type": "`$STRING`"
                  },
                  {
                    "active": true,
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "reqd": false,
                    "type": "`$STRING`"
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/userinvitespaged",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "userinvitespaged"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "merchant_id",
                  "page_number",
                  "page_size",
                  "search",
                  "sort"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/userinvites/{id}",
              "parts": [
                "api",
                "v1",
                "userinvites",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "userinvite_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/userinvites/{id}/details",
              "parts": [
                "api",
                "v1",
                "userinvites",
                "{userinvite_id}",
                "details"
              ],
              "rename": {
                "param": {
                  "id": "userinvite_id"
                }
              },
              "select": {
                "exist": [
                  "userinvite_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/userinvites/{id}",
              "parts": [
                "api",
                "v1",
                "userinvites",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/userinvites/{id}",
              "parts": [
                "api",
                "v1",
                "userinvites",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": [
          [
            "merchant"
          ],
          [
            "userinvite"
          ]
        ]
      }
    },
    "virtual": {
      "fields": [
        {
          "active": true,
          "name": "account_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "account_supplier_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "available_balance",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 2
        },
        {
          "active": true,
          "name": "available_balance_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 3
        },
        {
          "active": true,
          "name": "balance",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 4
        },
        {
          "active": true,
          "name": "balance_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 5
        },
        {
          "active": true,
          "name": "bank_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "consent_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 7
        },
        {
          "active": true,
          "name": "created_by",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 8
        },
        {
          "active": true,
          "name": "created_by_display_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "currency",
          "req": false,
          "type": "`$STRING`",
          "index$": 10
        },
        {
          "active": true,
          "name": "default_payment_rail",
          "req": false,
          "type": "`$STRING`",
          "index$": 11
        },
        {
          "active": true,
          "name": "display_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 12
        },
        {
          "active": true,
          "name": "expiry_date",
          "req": false,
          "type": "`$STRING`",
          "index$": 13
        },
        {
          "active": true,
          "name": "external_account_icon",
          "req": false,
          "type": "`$STRING`",
          "index$": 14
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 15
        },
        {
          "active": true,
          "name": "identifier",
          "req": true,
          "type": "`$OBJECT`",
          "index$": 16
        },
        {
          "active": true,
          "name": "inserted",
          "req": false,
          "type": "`$STRING`",
          "index$": 17
        },
        {
          "active": true,
          "name": "is_archived",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 18
        },
        {
          "active": true,
          "name": "is_connected_account",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 19
        },
        {
          "active": true,
          "name": "is_default",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 20
        },
        {
          "active": true,
          "name": "is_trust_account",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 21
        },
        {
          "active": true,
          "name": "is_virtual",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 22
        },
        {
          "active": true,
          "name": "last_transaction",
          "req": false,
          "type": "`$OBJECT`",
          "index$": 23
        },
        {
          "active": true,
          "name": "last_updated",
          "req": false,
          "type": "`$STRING`",
          "index$": 24
        },
        {
          "active": true,
          "name": "merchant_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 25
        },
        {
          "active": true,
          "name": "merchant_name",
          "req": false,
          "type": "`$STRING`",
          "index$": 26
        },
        {
          "active": true,
          "name": "name",
          "req": true,
          "type": "`$STRING`",
          "index$": 27
        },
        {
          "active": true,
          "name": "physical_account_id",
          "req": false,
          "type": "`$STRING`",
          "index$": 28
        },
        {
          "active": true,
          "name": "rule",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 29
        },
        {
          "active": true,
          "name": "submitted_payouts_balance",
          "req": false,
          "type": "`$NUMBER`",
          "index$": 30
        },
        {
          "active": true,
          "name": "submitted_payouts_balance_minor_unit",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 31
        },
        {
          "active": true,
          "name": "summary",
          "req": false,
          "type": "`$STRING`",
          "index$": 32
        },
        {
          "active": true,
          "name": "supplier_sepa_instant_status",
          "req": false,
          "type": "`$STRING`",
          "index$": 33
        },
        {
          "active": true,
          "name": "xero_bank_feed_connection_status",
          "req": false,
          "type": "`$STRING`",
          "index$": 34
        },
        {
          "active": true,
          "name": "xero_bank_feed_last_synced_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 35
        },
        {
          "active": true,
          "name": "xero_bank_feed_sync_last_failed_at",
          "req": false,
          "type": "`$STRING`",
          "index$": 36
        },
        {
          "active": true,
          "name": "xero_bank_feed_sync_last_failure_reason",
          "req": false,
          "type": "`$STRING`",
          "index$": 37
        },
        {
          "active": true,
          "name": "xero_bank_feed_sync_status",
          "req": false,
          "type": "`$STRING`",
          "index$": 38
        },
        {
          "active": true,
          "name": "xero_unsynchronised_transactions_count",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 39
        }
      ],
      "name": "virtual",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "POST",
              "orig": "/api/v1/accounts/{accountID}/virtual",
              "parts": [
                "api",
                "v1",
                "accounts",
                "{account_id}",
                "virtual"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id"
                }
              },
              "select": {
                "exist": [
                  "account_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "virtual_account_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/accounts/{accountID}/virtual/{virtualAccountID}",
              "parts": [
                "api",
                "v1",
                "accounts",
                "{account_id}",
                "virtual",
                "{id}"
              ],
              "rename": {
                "param": {
                  "accountID": "account_id",
                  "virtualAccountID": "id"
                }
              },
              "select": {
                "exist": [
                  "account_id",
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": [
          [
            "account"
          ]
        ]
      }
    },
    "webhook": {
      "fields": [
        {
          "active": true,
          "name": "destination_url",
          "req": false,
          "type": "`$STRING`",
          "index$": 0
        },
        {
          "active": true,
          "name": "email_address",
          "req": false,
          "type": "`$STRING`",
          "index$": 1
        },
        {
          "active": true,
          "name": "failed_notification_email_address",
          "req": false,
          "type": "`$STRING`",
          "index$": 2
        },
        {
          "active": true,
          "name": "id",
          "req": false,
          "type": "`$STRING`",
          "index$": 3
        },
        {
          "active": true,
          "name": "is_active",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 4
        },
        {
          "active": true,
          "name": "merchant_id",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            },
            "update": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$STRING`",
          "index$": 5
        },
        {
          "active": true,
          "name": "notification_method",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            },
            "update": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "req": false,
          "type": "`$STRING`",
          "index$": 6
        },
        {
          "active": true,
          "name": "resource_type",
          "req": false,
          "type": "`$ARRAY`",
          "index$": 7
        },
        {
          "active": true,
          "name": "retry",
          "req": false,
          "type": "`$BOOLEAN`",
          "index$": 8
        },
        {
          "active": true,
          "name": "secret",
          "req": false,
          "type": "`$STRING`",
          "index$": 9
        },
        {
          "active": true,
          "name": "version",
          "req": false,
          "type": "`$INTEGER`",
          "index$": 10
        }
      ],
      "name": "webhook",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "active": true,
              "args": {},
              "method": "POST",
              "orig": "/api/v1/webhooks",
              "parts": [
                "api",
                "v1",
                "webhooks"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "create"
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/webhooks",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "webhooks"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "merchant_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "list"
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  },
                  {
                    "active": true,
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 1
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/merchants/{merchantID}/webhooks/{id}",
              "parts": [
                "api",
                "v1",
                "merchants",
                "{merchant_id}",
                "webhooks",
                "{id}"
              ],
              "rename": {
                "param": {
                  "merchantID": "merchant_id"
                }
              },
              "select": {
                "exist": [
                  "id",
                  "merchant_id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            },
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "GET",
              "orig": "/api/v1/webhooks/{merchantID}",
              "parts": [
                "api",
                "v1",
                "webhooks",
                "{id}"
              ],
              "rename": {
                "param": {
                  "merchantID": "id"
                }
              },
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 1
            }
          ],
          "key$": "load"
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "DELETE",
              "orig": "/api/v1/webhooks/{id}",
              "parts": [
                "api",
                "v1",
                "webhooks",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "remove"
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "active": true,
              "args": {
                "params": [
                  {
                    "active": true,
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`",
                    "index$": 0
                  }
                ]
              },
              "method": "PUT",
              "orig": "/api/v1/webhooks/{id}",
              "parts": [
                "api",
                "v1",
                "webhooks",
                "{id}"
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "index$": 0
            }
          ],
          "key$": "update"
        }
      },
      "relations": {
        "ancestors": [
          [
            "merchant"
          ]
        ]
      }
    }
  }
}


const config = new Config()

export {
  config
}

