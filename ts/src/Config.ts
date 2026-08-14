
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
    name: 'Nofrixion',
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
          "name": "accountBalances",
          "type": "`$ARRAY`"
        },
        {
          "name": "accountID",
          "type": "`$STRING`"
        },
        {
          "name": "accountIdentifications",
          "type": "`$ARRAY`"
        },
        {
          "name": "accountName",
          "type": "`$STRING`"
        },
        {
          "name": "accountNames",
          "type": "`$ARRAY`"
        },
        {
          "name": "accountSupplierName",
          "type": "`$STRING`"
        },
        {
          "name": "accountType",
          "type": "`$STRING`"
        },
        {
          "name": "availableBalance",
          "type": "`$NUMBER`"
        },
        {
          "name": "availableBalanceMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "balance",
          "type": "`$NUMBER`"
        },
        {
          "name": "balanceMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "bankName",
          "type": "`$STRING`"
        },
        {
          "name": "consentID",
          "type": "`$STRING`"
        },
        {
          "name": "consolidatedAccountInformation",
          "type": "`$OBJECT`"
        },
        {
          "name": "createdBy",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "createdByDisplayName",
          "type": "`$STRING`"
        },
        {
          "name": "currency",
          "type": "`$STRING`"
        },
        {
          "name": "defaultPaymentRail",
          "type": "`$STRING`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "details",
          "type": "`$STRING`"
        },
        {
          "name": "displayName",
          "type": "`$STRING`"
        },
        {
          "name": "expiryDate",
          "type": "`$STRING`"
        },
        {
          "name": "externalAccountIcon",
          "type": "`$STRING`"
        },
        {
          "name": "format",
          "type": "`$STRING`"
        },
        {
          "name": "fromDate",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "identifier",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "isArchived",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isConnectedAccount",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isDefault",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isTrustAccount",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isVirtual",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lastTransaction",
          "type": "`$OBJECT`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "merchantName",
          "type": "`$STRING`"
        },
        {
          "name": "nickname",
          "type": "`$STRING`"
        },
        {
          "name": "physicalAccountID",
          "type": "`$STRING`"
        },
        {
          "name": "roleIDs",
          "type": "`$ARRAY`"
        },
        {
          "name": "rules",
          "type": "`$ARRAY`"
        },
        {
          "name": "submittedPayoutsBalance",
          "type": "`$NUMBER`"
        },
        {
          "name": "submittedPayoutsBalanceMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "summary",
          "type": "`$STRING`"
        },
        {
          "name": "supplierPhysicalAccountID",
          "type": "`$STRING`"
        },
        {
          "name": "supplierSepaInstantStatus",
          "type": "`$STRING`"
        },
        {
          "name": "toDate",
          "type": "`$STRING`"
        },
        {
          "name": "type",
          "type": "`$STRING`"
        },
        {
          "name": "usageType",
          "type": "`$STRING`"
        },
        {
          "name": "xeroBankFeedConnectionStatus",
          "type": "`$STRING`"
        },
        {
          "name": "xeroBankFeedLastSyncedAt",
          "type": "`$STRING`"
        },
        {
          "name": "xeroBankFeedSyncLastFailedAt",
          "type": "`$STRING`"
        },
        {
          "name": "xeroBankFeedSyncLastFailureReason",
          "type": "`$STRING`"
        },
        {
          "name": "xeroBankFeedSyncStatus",
          "type": "`$STRING`"
        },
        {
          "name": "xeroUnsynchronisedTransactionsCount",
          "type": "`$INTEGER`"
        }
      ],
      "name": "account",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "currency",
                    "orig": "currency",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "example": false,
                    "kind": "query",
                    "name": "connected_account",
                    "orig": "connected_account",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "only_connect_account",
                    "orig": "only_connect_account",
                    "type": "`$BOOLEAN`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "example": false,
                    "kind": "query",
                    "name": "connected_account",
                    "orig": "connected_account",
                    "type": "`$BOOLEAN`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "example": false,
                    "kind": "query",
                    "name": "connected_account",
                    "orig": "connected_account",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$ARRAY`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_child_merchant",
                    "orig": "include_child_merchant",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "only_archived",
                    "orig": "only_archived",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "only_connect_account",
                    "orig": "only_connect_account",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "credit_type",
                    "orig": "credit_type",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "amount",
                    "orig": "amount",
                    "reqd": true,
                    "type": "`$NUMBER`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "approveUrl",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "payouts",
          "type": "`$ARRAY`"
        }
      ],
      "name": "batch",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "beneficiary": {
      "fields": [
        {
          "name": "approvalCallbackUrl",
          "type": "`$STRING`"
        },
        {
          "name": "authenticationMethods",
          "type": "`$ARRAY`"
        },
        {
          "name": "authorisations",
          "type": "`$ARRAY`"
        },
        {
          "name": "authorisersCompletedCount",
          "type": "`$INTEGER`"
        },
        {
          "name": "authorisersRequiredCount",
          "type": "`$INTEGER`"
        },
        {
          "name": "beneficiaries",
          "type": "`$ARRAY`"
        },
        {
          "name": "beneficiaryEvents",
          "type": "`$ARRAY`"
        },
        {
          "name": "canAuthorise",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "canUpdate",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "createdBy",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "createdByEmailAddress",
          "type": "`$STRING`"
        },
        {
          "name": "currency",
          "op": {
            "update": {
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "destination",
          "op": {
            "create": {
              "req": true,
              "type": "`$OBJECT`"
            }
          },
          "type": "`$OBJECT`"
        },
        {
          "name": "failedBeneficiaries",
          "type": "`$OBJECT`"
        },
        {
          "name": "hasCurrentUserAuthorised",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "isEnabled",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lastAuthorised",
          "type": "`$STRING`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "op": {
            "update": {
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "nonce",
          "type": "`$STRING`"
        },
        {
          "name": "sourceAccountIDs",
          "type": "`$ARRAY`"
        },
        {
          "name": "sourceAccounts",
          "type": "`$ARRAY`"
        },
        {
          "name": "theirReference",
          "type": "`$STRING`"
        }
      ],
      "name": "beneficiary",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_disabled",
                    "orig": "include_disabled",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "source_account_id",
                    "orig": "source_account_id",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_disabled",
                    "orig": "include_disabled",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "source_account_id",
                    "orig": "source_account_id",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_disabled",
                    "orig": "include_disabled",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "groupMembers",
          "type": "`$ARRAY`"
        },
        {
          "name": "groupName",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "req": true,
          "type": "`$STRING`"
        }
      ],
      "name": "beneficiary_group",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "authorizedAmount",
          "type": "`$STRING`"
        },
        {
          "name": "currencyCode",
          "type": "`$STRING`"
        },
        {
          "name": "isPayerAuthenticationRequired",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isSoftDecline",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "payerAuthenticationAccessToken",
          "type": "`$STRING`"
        },
        {
          "name": "payerAuthenticationMerchantData",
          "type": "`$STRING`"
        },
        {
          "name": "payerAuthenticationUrl",
          "type": "`$STRING`"
        },
        {
          "name": "payerAuthenticationWindowHeight",
          "type": "`$INTEGER`"
        },
        {
          "name": "payerAuthenticationWindowWidth",
          "type": "`$INTEGER`"
        },
        {
          "name": "paymentRequestCallbackUrl",
          "type": "`$STRING`"
        },
        {
          "name": "paymentRequestID",
          "type": "`$STRING`"
        },
        {
          "name": "requestID",
          "type": "`$STRING`"
        },
        {
          "name": "responseCode",
          "type": "`$STRING`"
        },
        {
          "name": "responseType",
          "type": "`$STRING`"
        },
        {
          "name": "status",
          "type": "`$STRING`"
        },
        {
          "name": "threeDSRedirectUrl",
          "type": "`$STRING`"
        },
        {
          "name": "transactionID",
          "type": "`$STRING`"
        }
      ],
      "name": "card",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "cardType",
          "type": "`$STRING`"
        },
        {
          "name": "customerEmailAddress",
          "type": "`$STRING`"
        },
        {
          "name": "expiryMonth",
          "type": "`$STRING`"
        },
        {
          "name": "expiryYear",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "lastFourDigits",
          "type": "`$STRING`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "maskedCardNumber",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "paymentRequestID",
          "type": "`$STRING`"
        }
      ],
      "name": "card_customer_token",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "customer_email_address",
                    "orig": "customer_email_address",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "customer_email_address",
                    "orig": "customer_email_address",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "customer_email_address",
                    "orig": "customer_email_address",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "customer_email_address",
                    "orig": "customer_email_address",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "authorizedAmount",
          "type": "`$STRING`"
        },
        {
          "name": "currencyCode",
          "type": "`$STRING`"
        },
        {
          "name": "isPayerAuthenticationRequired",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isSoftDecline",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "payerAuthenticationAccessToken",
          "type": "`$STRING`"
        },
        {
          "name": "payerAuthenticationMerchantData",
          "type": "`$STRING`"
        },
        {
          "name": "payerAuthenticationUrl",
          "type": "`$STRING`"
        },
        {
          "name": "payerAuthenticationWindowHeight",
          "type": "`$INTEGER`"
        },
        {
          "name": "payerAuthenticationWindowWidth",
          "type": "`$INTEGER`"
        },
        {
          "name": "paymentRequestCallbackUrl",
          "type": "`$STRING`"
        },
        {
          "name": "paymentRequestID",
          "type": "`$STRING`"
        },
        {
          "name": "requestID",
          "type": "`$STRING`"
        },
        {
          "name": "responseCode",
          "type": "`$STRING`"
        },
        {
          "name": "responseType",
          "type": "`$STRING`"
        },
        {
          "name": "status",
          "type": "`$STRING`"
        },
        {
          "name": "threeDSRedirectUrl",
          "type": "`$STRING`"
        },
        {
          "name": "transactionID",
          "type": "`$STRING`"
        }
      ],
      "name": "card_payment",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "partial_refund_amount",
                    "orig": "partial_refund_amount",
                    "reqd": true,
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "jwt",
          "type": "`$STRING`"
        }
      ],
      "name": "card_public_key",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "authorisationUrl",
          "type": "`$STRING`"
        },
        {
          "name": "callbackUrl",
          "type": "`$STRING`"
        },
        {
          "name": "consentID",
          "type": "`$STRING`"
        },
        {
          "name": "emailAddress",
          "type": "`$STRING`"
        },
        {
          "name": "expiryDate",
          "type": "`$STRING`"
        },
        {
          "name": "failureCallbackUrl",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "institutionID",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "type": "`$STRING`"
        },
        {
          "name": "isConnectedAccounts",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isEnabled",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "merchantID",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "type": "`$STRING`"
        },
        {
          "name": "provider",
          "type": "`$STRING`"
        },
        {
          "name": "successWebHookUrl",
          "type": "`$STRING`"
        }
      ],
      "name": "consent",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "email",
                    "orig": "email",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "code",
          "type": "`$STRING`"
        },
        {
          "name": "decimals",
          "type": "`$INTEGER`"
        },
        {
          "name": "isFiat",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "iso4217AlphaCode",
          "type": "`$STRING`"
        },
        {
          "name": "iso4217NumericCode",
          "type": "`$STRING`"
        },
        {
          "name": "symbol",
          "type": "`$STRING`"
        }
      ],
      "name": "currency",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "capability",
                    "orig": "capability",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "direct_debit_batch_submit": {
      "fields": [
        {
          "name": "failedSubmissions",
          "type": "`$OBJECT`"
        },
        {
          "name": "successfulSubmissions",
          "type": "`$ARRAY`"
        }
      ],
      "name": "direct_debit_batch_submit",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "fx_rate": {
      "fields": [
        {
          "name": "destinationCurrency",
          "type": "`$STRING`"
        },
        {
          "name": "exchangeRate",
          "type": "`$NUMBER`"
        },
        {
          "name": "expiryTime",
          "type": "`$STRING`"
        },
        {
          "name": "quoteID",
          "type": "`$STRING`"
        },
        {
          "name": "sourceCurrency",
          "type": "`$STRING`"
        }
      ],
      "name": "fx_rate",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "destination",
                    "orig": "destination",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "source",
                    "orig": "source",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "destination",
                    "orig": "destination",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "source",
                    "orig": "source",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "valid_for_minute",
                    "orig": "valid_for_minute",
                    "reqd": true,
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "paymentRequestID",
          "type": "`$STRING`"
        },
        {
          "name": "responseType",
          "type": "`$STRING`"
        }
      ],
      "name": "i_payment",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "mandate": {
      "fields": [
        {
          "name": "accountNumber",
          "type": "`$STRING`"
        },
        {
          "name": "addressLine1",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "addressLine2",
          "type": "`$STRING`"
        },
        {
          "name": "approvedAt",
          "type": "`$STRING`"
        },
        {
          "name": "city",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "countryCode",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "currency",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "type": "`$STRING`"
        },
        {
          "name": "customerAccountNumber",
          "type": "`$STRING`"
        },
        {
          "name": "customerCity",
          "type": "`$STRING`"
        },
        {
          "name": "customerCountryCode",
          "type": "`$STRING`"
        },
        {
          "name": "customerCountryName",
          "type": "`$STRING`"
        },
        {
          "name": "customerEmailAddress",
          "type": "`$STRING`"
        },
        {
          "name": "customerFirstName",
          "type": "`$STRING`"
        },
        {
          "name": "customerIban",
          "type": "`$STRING`"
        },
        {
          "name": "customerLastName",
          "type": "`$STRING`"
        },
        {
          "name": "customerSortCode",
          "type": "`$STRING`"
        },
        {
          "name": "emailAddress",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "firstName",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "iban",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "isRecurring",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lastName",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "type": "`$STRING`"
        },
        {
          "name": "postalCode",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "reference",
          "type": "`$STRING`"
        },
        {
          "name": "sortCode",
          "type": "`$STRING`"
        },
        {
          "name": "status",
          "type": "`$STRING`"
        },
        {
          "name": "supplierBankAccountID",
          "type": "`$STRING`"
        },
        {
          "name": "supplierCustomerID",
          "type": "`$STRING`"
        },
        {
          "name": "supplierMandateID",
          "type": "`$STRING`"
        },
        {
          "name": "supplierName",
          "type": "`$STRING`"
        },
        {
          "name": "supplierStatus",
          "type": "`$STRING`"
        }
      ],
      "name": "mandate",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "merchant": {
      "fields": [
        {
          "name": "accountCurrencies",
          "type": "`$ARRAY`"
        },
        {
          "name": "canHaveTrustAccounts",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "cardPaymentProcessor",
          "type": "`$STRING`"
        },
        {
          "name": "companyID",
          "type": "`$STRING`"
        },
        {
          "name": "displayQrOnHostedPay",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "hostedPayVersion",
          "type": "`$INTEGER`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "isBlocked",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isExited",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isSuspended",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "jurisdiction",
          "type": "`$STRING`"
        },
        {
          "name": "logoUrlPng",
          "type": "`$STRING`"
        },
        {
          "name": "logoUrlSvg",
          "type": "`$STRING`"
        },
        {
          "name": "merchantCategoryCode",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "type": "`$STRING`"
        },
        {
          "name": "notes",
          "type": "`$STRING`"
        },
        {
          "name": "parentMerchant",
          "type": "`$OBJECT`"
        },
        {
          "name": "paymentAccountLimit",
          "type": "`$INTEGER`"
        },
        {
          "name": "paymentAccounts",
          "type": "`$ARRAY`"
        },
        {
          "name": "reason",
          "type": "`$STRING`"
        },
        {
          "name": "shortName",
          "type": "`$STRING`"
        },
        {
          "name": "supportedPaymentMethodsList",
          "type": "`$ARRAY`"
        },
        {
          "name": "suspensionReason",
          "type": "`$STRING`"
        },
        {
          "name": "tags",
          "type": "`$ARRAY`"
        },
        {
          "name": "timeZoneId",
          "type": "`$STRING`"
        },
        {
          "name": "tradingName",
          "type": "`$STRING`"
        },
        {
          "name": "webHookLimit",
          "type": "`$INTEGER`"
        },
        {
          "name": "yourRoleName",
          "type": "`$STRING`"
        }
      ],
      "name": "merchant",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "query": [
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_suspended",
                    "orig": "include_suspended",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 10,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_disabled",
                    "orig": "include_disabled",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "user_id",
                    "orig": "user_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "tag_id",
                    "orig": "tag_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "amountLower",
          "type": "`$NUMBER`"
        },
        {
          "name": "amountUpper",
          "type": "`$NUMBER`"
        },
        {
          "name": "authorisationType",
          "type": "`$STRING`"
        },
        {
          "name": "beneficiariesOnly",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "lastEditorCantAuthorise",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "numberOfAuthorisers",
          "type": "`$INTEGER`"
        },
        {
          "name": "roleSettings",
          "type": "`$ARRAY`"
        }
      ],
      "name": "merchant_authorisation_setting",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "approvedAt",
          "type": "`$STRING`"
        },
        {
          "name": "currency",
          "type": "`$STRING`"
        },
        {
          "name": "customerAccountNumber",
          "type": "`$STRING`"
        },
        {
          "name": "customerCity",
          "type": "`$STRING`"
        },
        {
          "name": "customerCountryCode",
          "type": "`$STRING`"
        },
        {
          "name": "customerCountryName",
          "type": "`$STRING`"
        },
        {
          "name": "customerEmailAddress",
          "type": "`$STRING`"
        },
        {
          "name": "customerFirstName",
          "type": "`$STRING`"
        },
        {
          "name": "customerIban",
          "type": "`$STRING`"
        },
        {
          "name": "customerLastName",
          "type": "`$STRING`"
        },
        {
          "name": "customerSortCode",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "isRecurring",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "reference",
          "type": "`$STRING`"
        },
        {
          "name": "status",
          "type": "`$STRING`"
        },
        {
          "name": "supplierBankAccountID",
          "type": "`$STRING`"
        },
        {
          "name": "supplierCustomerID",
          "type": "`$STRING`"
        },
        {
          "name": "supplierMandateID",
          "type": "`$STRING`"
        },
        {
          "name": "supplierName",
          "type": "`$STRING`"
        },
        {
          "name": "supplierStatus",
          "type": "`$STRING`"
        }
      ],
      "name": "merchant_direct_debit_mandate_page",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "mandate_i_d",
                    "orig": "mandate_i_d",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "merchant_pay_by_bank_setting": {
      "fields": [
        {
          "name": "bankCountryCodes",
          "type": "`$ARRAY`"
        },
        {
          "name": "bankID",
          "type": "`$STRING`"
        },
        {
          "name": "bankName",
          "type": "`$STRING`"
        },
        {
          "name": "businessInstitutionID",
          "type": "`$STRING`"
        },
        {
          "name": "currency",
          "type": "`$STRING`"
        },
        {
          "name": "logo",
          "type": "`$STRING`"
        },
        {
          "name": "message",
          "type": "`$STRING`"
        },
        {
          "name": "messageImageUrl",
          "type": "`$STRING`"
        },
        {
          "name": "order",
          "type": "`$INTEGER`"
        },
        {
          "name": "personalInstitutionID",
          "type": "`$STRING`"
        },
        {
          "name": "processor",
          "type": "`$STRING`"
        },
        {
          "name": "warningHeading",
          "type": "`$STRING`"
        },
        {
          "name": "warningMessage",
          "type": "`$STRING`"
        }
      ],
      "name": "merchant_pay_by_bank_setting",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "country_code",
                    "orig": "country_code",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "PIS",
                    "kind": "query",
                    "name": "open_banking_operation",
                    "orig": "open_banking_operation",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "bankPaymentOptions",
          "type": "`$OBJECT`"
        },
        {
          "name": "cardPaymentAddressOptions",
          "type": "`$OBJECT`"
        },
        {
          "name": "cardPaymentCaptureOptions",
          "type": "`$OBJECT`"
        },
        {
          "name": "customFields",
          "type": "`$ARRAY`"
        },
        {
          "name": "defaultFields",
          "type": "`$ARRAY`"
        },
        {
          "name": "description",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "notificationOptions",
          "type": "`$OBJECT`"
        },
        {
          "name": "paymentMethods",
          "type": "`$OBJECT`"
        },
        {
          "name": "paymentTerms",
          "type": "`$OBJECT`"
        },
        {
          "name": "priorityBankOptions",
          "type": "`$OBJECT`"
        },
        {
          "name": "template",
          "req": true,
          "type": "`$OBJECT`"
        }
      ],
      "name": "merchant_payment_request_template",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "template_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "template_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "template_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "authenticationMethods",
          "type": "`$ARRAY`"
        },
        {
          "name": "authorisations",
          "type": "`$ARRAY`"
        },
        {
          "name": "authorisersCompletedCount",
          "type": "`$INTEGER`"
        },
        {
          "name": "authorisersRequiredCount",
          "type": "`$INTEGER`"
        },
        {
          "name": "canAuthorise",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "description",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "type": "`$STRING`"
        },
        {
          "name": "expiresAt",
          "type": "`$STRING`"
        },
        {
          "name": "hasCurrentUserAuthorised",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "hmacAlgorithm",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "ipAddressWhitelist",
          "type": "`$STRING`"
        },
        {
          "name": "isArchived",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isEnabled",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lastAuthorised",
          "type": "`$STRING`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "type": "`$STRING`"
        },
        {
          "name": "nonce",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "permissionTypes",
          "type": "`$ARRAY`"
        },
        {
          "name": "requestSignatureVersion",
          "type": "`$INTEGER`"
        },
        {
          "name": "sharedSecretAlgorithm",
          "type": "`$STRING`"
        },
        {
          "name": "sharedSecretBase64",
          "type": "`$STRING`"
        },
        {
          "name": "token",
          "type": "`$STRING`"
        }
      ],
      "name": "merchant_token",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "id",
                    "orig": "id",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "secret",
                    "orig": "secret",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "url",
                    "orig": "url",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "no_frixion_version": {
      "fields": [
        {
          "name": "buildVersion",
          "type": "`$INTEGER`"
        },
        {
          "name": "majorVersion",
          "type": "`$INTEGER`"
        },
        {
          "name": "minorVersion",
          "type": "`$INTEGER`"
        },
        {
          "name": "releaseName",
          "type": "`$STRING`"
        }
      ],
      "name": "no_frixion_version",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
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
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "email",
                    "orig": "email",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "accountName",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "accountNumber",
          "type": "`$STRING`"
        },
        {
          "name": "iban",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "payeeVerifiedAccountName",
          "type": "`$STRING`"
        },
        {
          "name": "result",
          "type": "`$STRING`"
        },
        {
          "name": "secondaryIdentification",
          "type": "`$STRING`"
        },
        {
          "name": "sortCode",
          "type": "`$STRING`"
        }
      ],
      "name": "payeeverification",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "payment": {
      "fields": [
        {
          "name": "addresses",
          "type": "`$ARRAY`"
        },
        {
          "name": "amount",
          "op": {
            "create": {
              "req": true,
              "type": "`$NUMBER`"
            }
          },
          "type": "`$NUMBER`"
        },
        {
          "name": "amountPending",
          "type": "`$NUMBER`"
        },
        {
          "name": "amountReceived",
          "type": "`$NUMBER`"
        },
        {
          "name": "amountRefunded",
          "type": "`$NUMBER`"
        },
        {
          "name": "autoSendReceipt",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "baseOriginUrl",
          "type": "`$STRING`"
        },
        {
          "name": "callbackUrl",
          "type": "`$STRING`"
        },
        {
          "name": "cardAuthorizeOnly",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "cardCreateToken",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "cardCreateTokenMode",
          "type": "`$STRING`"
        },
        {
          "name": "cardIgnoreCVN",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "cardNoPayerAuthentication",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "cardProcessorMerchantID",
          "type": "`$STRING`"
        },
        {
          "name": "cardStripePaymentIntentID",
          "type": "`$STRING`"
        },
        {
          "name": "cardStripePaymentIntentSecret",
          "type": "`$STRING`"
        },
        {
          "name": "cardTransmitRawDetails",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "createdByUser",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "currency",
          "type": "`$STRING`"
        },
        {
          "name": "customFields",
          "type": "`$ARRAY`"
        },
        {
          "name": "customerEmailAddress",
          "type": "`$STRING`"
        },
        {
          "name": "customerID",
          "type": "`$STRING`"
        },
        {
          "name": "customerName",
          "type": "`$STRING`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "destinationAccount",
          "type": "`$OBJECT`"
        },
        {
          "name": "directDebitPayment",
          "type": "`$OBJECT`"
        },
        {
          "name": "dueDate",
          "type": "`$STRING`"
        },
        {
          "name": "events",
          "type": "`$ARRAY`"
        },
        {
          "name": "failureCallbackUrl",
          "type": "`$STRING`"
        },
        {
          "name": "fieldDisplaySettings",
          "type": "`$ARRAY`"
        },
        {
          "name": "formattedAmount",
          "type": "`$STRING`"
        },
        {
          "name": "hostedPayCheckoutUrl",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "ignoreAddressVerification",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "insertedSortable",
          "type": "`$STRING`"
        },
        {
          "name": "isArchived",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "jwk",
          "type": "`$STRING`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "lightningInvoice",
          "type": "`$STRING`"
        },
        {
          "name": "lightningInvoiceExpiresAt",
          "type": "`$STRING`"
        },
        {
          "name": "merchantDirectDebitMandateID",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "merchantTokenDescription",
          "type": "`$STRING`"
        },
        {
          "name": "notificationEmailAddresses",
          "type": "`$STRING`"
        },
        {
          "name": "notificationRoleIDs",
          "type": "`$ARRAY`"
        },
        {
          "name": "orderID",
          "type": "`$STRING`"
        },
        {
          "name": "partialPaymentMethod",
          "type": "`$STRING`"
        },
        {
          "name": "partialPaymentSteps",
          "type": "`$STRING`"
        },
        {
          "name": "paymentAttempts",
          "type": "`$ARRAY`"
        },
        {
          "name": "paymentMethods",
          "type": "`$ARRAY`"
        },
        {
          "name": "paymentProcessor",
          "type": "`$STRING`"
        },
        {
          "name": "payrunID",
          "type": "`$STRING`"
        },
        {
          "name": "pispAccountID",
          "type": "`$STRING`"
        },
        {
          "name": "priorityBankID",
          "type": "`$STRING`"
        },
        {
          "name": "result",
          "type": "`$OBJECT`"
        },
        {
          "name": "sandboxSettleDelayInSeconds",
          "type": "`$INTEGER`"
        },
        {
          "name": "shippingAddress",
          "type": "`$OBJECT`"
        },
        {
          "name": "shippingAddressCity",
          "type": "`$STRING`"
        },
        {
          "name": "shippingAddressCountryCode",
          "type": "`$STRING`"
        },
        {
          "name": "shippingAddressCounty",
          "type": "`$STRING`"
        },
        {
          "name": "shippingAddressLine1",
          "type": "`$STRING`"
        },
        {
          "name": "shippingAddressLine2",
          "type": "`$STRING`"
        },
        {
          "name": "shippingAddressPostCode",
          "type": "`$STRING`"
        },
        {
          "name": "shippingEmail",
          "type": "`$STRING`"
        },
        {
          "name": "shippingFirstName",
          "type": "`$STRING`"
        },
        {
          "name": "shippingLastName",
          "type": "`$STRING`"
        },
        {
          "name": "shippingPhone",
          "type": "`$STRING`"
        },
        {
          "name": "status",
          "type": "`$STRING`"
        },
        {
          "name": "successWebHookUrl",
          "type": "`$STRING`"
        },
        {
          "name": "tagIds",
          "type": "`$ARRAY`"
        },
        {
          "name": "tags",
          "type": "`$ARRAY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "tokenisedCards",
          "type": "`$ARRAY`"
        },
        {
          "name": "transactions",
          "type": "`$ARRAY`"
        },
        {
          "name": "useHostedPaymentPage",
          "type": "`$BOOLEAN`"
        }
      ],
      "name": "payment",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_event",
                    "orig": "include_event",
                    "type": "`$BOOLEAN`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "order_id",
                    "orig": "order_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "accountName",
          "type": "`$STRING`"
        },
        {
          "name": "accountSupplierName",
          "type": "`$STRING`"
        },
        {
          "name": "availableBalance",
          "type": "`$NUMBER`"
        },
        {
          "name": "availableBalanceMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "balance",
          "type": "`$NUMBER`"
        },
        {
          "name": "balanceMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "bankName",
          "type": "`$STRING`"
        },
        {
          "name": "consentID",
          "type": "`$STRING`"
        },
        {
          "name": "createdBy",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "createdByDisplayName",
          "type": "`$STRING`"
        },
        {
          "name": "currency",
          "type": "`$STRING`"
        },
        {
          "name": "defaultPaymentRail",
          "type": "`$STRING`"
        },
        {
          "name": "displayName",
          "type": "`$STRING`"
        },
        {
          "name": "expiryDate",
          "type": "`$STRING`"
        },
        {
          "name": "externalAccountIcon",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "identifier",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "isArchived",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isConnectedAccount",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isDefault",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isTrustAccount",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isVirtual",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lastTransaction",
          "type": "`$OBJECT`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "merchantName",
          "type": "`$STRING`"
        },
        {
          "name": "physicalAccountID",
          "type": "`$STRING`"
        },
        {
          "name": "rules",
          "type": "`$ARRAY`"
        },
        {
          "name": "submittedPayoutsBalance",
          "type": "`$NUMBER`"
        },
        {
          "name": "submittedPayoutsBalanceMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "summary",
          "type": "`$STRING`"
        },
        {
          "name": "supplierSepaInstantStatus",
          "type": "`$STRING`"
        },
        {
          "name": "xeroBankFeedConnectionStatus",
          "type": "`$STRING`"
        },
        {
          "name": "xeroBankFeedLastSyncedAt",
          "type": "`$STRING`"
        },
        {
          "name": "xeroBankFeedSyncLastFailedAt",
          "type": "`$STRING`"
        },
        {
          "name": "xeroBankFeedSyncLastFailureReason",
          "type": "`$STRING`"
        },
        {
          "name": "xeroBankFeedSyncStatus",
          "type": "`$STRING`"
        },
        {
          "name": "xeroUnsynchronisedTransactionsCount",
          "type": "`$INTEGER`"
        }
      ],
      "name": "payment_account",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "example": false,
                    "kind": "query",
                    "name": "connected_account",
                    "orig": "connected_account",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$ARRAY`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_child_merchant",
                    "orig": "include_child_merchant",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "only_archived",
                    "orig": "only_archived",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "only_connect_account",
                    "orig": "only_connect_account",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "accountName",
          "type": "`$STRING`"
        },
        {
          "name": "availableBalance",
          "type": "`$NUMBER`"
        },
        {
          "name": "balance",
          "type": "`$NUMBER`"
        },
        {
          "name": "balanceMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "currency",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "identifier",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "isArchived",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isConnectedAccount",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "submittedPayoutsBalance",
          "type": "`$NUMBER`"
        }
      ],
      "name": "payment_account_minimal",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "example": false,
                    "kind": "query",
                    "name": "connected_account",
                    "orig": "connected_account",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$ARRAY`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "only_archived",
                    "orig": "only_archived",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "only_connect_account",
                    "orig": "only_connect_account",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 10,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "payment_initiation": {
      "fields": [
        {
          "name": "paymentInitiationID",
          "type": "`$STRING`"
        },
        {
          "name": "paymentRequestCallbackUrl",
          "type": "`$STRING`"
        },
        {
          "name": "paymentRequestID",
          "type": "`$STRING`"
        },
        {
          "name": "redirectUrl",
          "type": "`$STRING`"
        },
        {
          "name": "responseType",
          "type": "`$STRING`"
        },
        {
          "name": "specificErrorMessage",
          "type": "`$STRING`"
        }
      ],
      "name": "payment_initiation",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "addresses",
          "type": "`$ARRAY`"
        },
        {
          "name": "amount",
          "type": "`$NUMBER`"
        },
        {
          "name": "amountPending",
          "type": "`$NUMBER`"
        },
        {
          "name": "amountReceived",
          "type": "`$NUMBER`"
        },
        {
          "name": "amountRefunded",
          "type": "`$NUMBER`"
        },
        {
          "name": "autoSendReceipt",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "baseOriginUrl",
          "type": "`$STRING`"
        },
        {
          "name": "callbackUrl",
          "type": "`$STRING`"
        },
        {
          "name": "cardAuthorizeOnly",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "cardCreateToken",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "cardCreateTokenMode",
          "type": "`$STRING`"
        },
        {
          "name": "cardIgnoreCVN",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "cardProcessorMerchantID",
          "type": "`$STRING`"
        },
        {
          "name": "cardStripePaymentIntentID",
          "type": "`$STRING`"
        },
        {
          "name": "cardStripePaymentIntentSecret",
          "type": "`$STRING`"
        },
        {
          "name": "createdByUser",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "currency",
          "type": "`$STRING`"
        },
        {
          "name": "customFields",
          "type": "`$ARRAY`"
        },
        {
          "name": "customerEmailAddress",
          "type": "`$STRING`"
        },
        {
          "name": "customerID",
          "type": "`$STRING`"
        },
        {
          "name": "customerName",
          "type": "`$STRING`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "destinationAccount",
          "type": "`$OBJECT`"
        },
        {
          "name": "directDebitPayment",
          "type": "`$OBJECT`"
        },
        {
          "name": "doSimulateSettlementFailure",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "dueDate",
          "type": "`$STRING`"
        },
        {
          "name": "errorDescription",
          "type": "`$STRING`"
        },
        {
          "name": "events",
          "type": "`$ARRAY`"
        },
        {
          "name": "failedPaymentRequests",
          "type": "`$OBJECT`"
        },
        {
          "name": "failureCallbackUrl",
          "type": "`$STRING`"
        },
        {
          "name": "fieldDisplaySettings",
          "type": "`$ARRAY`"
        },
        {
          "name": "formattedAmount",
          "type": "`$STRING`"
        },
        {
          "name": "hostedPayCheckoutUrl",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "ignoreAddressVerification",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "insertedSortable",
          "type": "`$STRING`"
        },
        {
          "name": "institution",
          "type": "`$STRING`"
        },
        {
          "name": "isArchived",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "jwk",
          "type": "`$STRING`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "lightningInvoice",
          "type": "`$STRING`"
        },
        {
          "name": "lightningInvoiceExpiresAt",
          "type": "`$STRING`"
        },
        {
          "name": "merchantDirectDebitMandateID",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "merchantTokenDescription",
          "type": "`$STRING`"
        },
        {
          "name": "notificationEmailAddresses",
          "type": "`$STRING`"
        },
        {
          "name": "notificationRoleIDs",
          "type": "`$ARRAY`"
        },
        {
          "name": "orderID",
          "type": "`$STRING`"
        },
        {
          "name": "partialPaymentMethod",
          "type": "`$STRING`"
        },
        {
          "name": "partialPaymentSteps",
          "type": "`$STRING`"
        },
        {
          "name": "paymentAttempts",
          "type": "`$ARRAY`"
        },
        {
          "name": "paymentInitiationID",
          "type": "`$STRING`"
        },
        {
          "name": "paymentMethods",
          "type": "`$ARRAY`"
        },
        {
          "name": "paymentProcessor",
          "type": "`$STRING`"
        },
        {
          "name": "paymentRequests",
          "type": "`$ARRAY`"
        },
        {
          "name": "payrunID",
          "type": "`$STRING`"
        },
        {
          "name": "pispAccountID",
          "type": "`$STRING`"
        },
        {
          "name": "priorityBankID",
          "type": "`$STRING`"
        },
        {
          "name": "result",
          "type": "`$OBJECT`"
        },
        {
          "name": "sandboxSettleDelayInSeconds",
          "type": "`$INTEGER`"
        },
        {
          "name": "shippingAddress",
          "type": "`$OBJECT`"
        },
        {
          "name": "status",
          "type": "`$STRING`"
        },
        {
          "name": "successWebHookUrl",
          "type": "`$STRING`"
        },
        {
          "name": "tags",
          "type": "`$ARRAY`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        },
        {
          "name": "tokenisedCards",
          "type": "`$ARRAY`"
        },
        {
          "name": "transactions",
          "type": "`$ARRAY`"
        },
        {
          "name": "useHostedPaymentPage",
          "type": "`$BOOLEAN`"
        }
      ],
      "name": "payment_request",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "mandate_id",
                    "orig": "mandate_id",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "submit_after",
                    "orig": "submit_after",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "payment_method",
                    "orig": "payment_method",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "All",
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "payment_method",
                    "orig": "payment_method",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "All",
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "amount",
          "req": true,
          "type": "`$NUMBER`"
        },
        {
          "name": "applePayTransactionID",
          "type": "`$STRING`"
        },
        {
          "name": "cardAuthorizationResponseID",
          "type": "`$STRING`"
        },
        {
          "name": "cardExpiryMonth",
          "type": "`$INTEGER`"
        },
        {
          "name": "cardExpiryYear",
          "type": "`$INTEGER`"
        },
        {
          "name": "cardIssuer",
          "type": "`$STRING`"
        },
        {
          "name": "cardIssuerCountry",
          "type": "`$STRING`"
        },
        {
          "name": "cardLastFourDigits",
          "type": "`$STRING`"
        },
        {
          "name": "cardRequestID",
          "type": "`$STRING`"
        },
        {
          "name": "cardScheme",
          "type": "`$STRING`"
        },
        {
          "name": "cardTokenCustomerID",
          "type": "`$STRING`"
        },
        {
          "name": "cardTransactionID",
          "type": "`$STRING`"
        },
        {
          "name": "currency",
          "type": "`$STRING`"
        },
        {
          "name": "directDebitPaymentID",
          "type": "`$STRING`"
        },
        {
          "name": "directDebitPaymentReference",
          "type": "`$STRING`"
        },
        {
          "name": "drirectDebitMandateID",
          "type": "`$STRING`"
        },
        {
          "name": "errorMessage",
          "type": "`$STRING`"
        },
        {
          "name": "errorReason",
          "type": "`$STRING`"
        },
        {
          "name": "eventType",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "lightningInvoice",
          "type": "`$STRING`"
        },
        {
          "name": "lightningRHash",
          "type": "`$STRING`"
        },
        {
          "name": "originUrl",
          "type": "`$STRING`"
        },
        {
          "name": "paymentMethodType",
          "type": "`$STRING`"
        },
        {
          "name": "paymentProcessorName",
          "type": "`$STRING`"
        },
        {
          "name": "paymentRequestID",
          "type": "`$STRING`"
        },
        {
          "name": "pispBankStatus",
          "type": "`$STRING`"
        },
        {
          "name": "pispPaymentInitiationID",
          "type": "`$STRING`"
        },
        {
          "name": "pispPaymentInstitutionName",
          "type": "`$STRING`"
        },
        {
          "name": "pispPaymentServiceProviderID",
          "type": "`$STRING`"
        },
        {
          "name": "pispRedirectUrl",
          "type": "`$STRING`"
        },
        {
          "name": "reconciledTransactionID",
          "type": "`$STRING`"
        },
        {
          "name": "refundPayoutID",
          "type": "`$STRING`"
        },
        {
          "name": "status",
          "type": "`$STRING`"
        },
        {
          "name": "walletName",
          "type": "`$STRING`"
        }
      ],
      "name": "payment_request_event",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
      "fields": [],
      "name": "payment_request_metric",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "payment_method",
                    "orig": "payment_method",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "payment_request_minimal": {
      "fields": [
        {
          "name": "amount",
          "type": "`$NUMBER`"
        },
        {
          "name": "amountPending",
          "type": "`$NUMBER`"
        },
        {
          "name": "amountReceived",
          "type": "`$NUMBER`"
        },
        {
          "name": "amountRefunded",
          "type": "`$NUMBER`"
        },
        {
          "name": "callbackUrl",
          "type": "`$STRING`"
        },
        {
          "name": "cardStripePaymentIntentSecret",
          "type": "`$STRING`"
        },
        {
          "name": "countryCode",
          "type": "`$STRING`"
        },
        {
          "name": "currency",
          "type": "`$STRING`"
        },
        {
          "name": "customFieldsToDisplay",
          "type": "`$ARRAY`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "dueDate",
          "type": "`$STRING`"
        },
        {
          "name": "fieldDisplaySettings",
          "type": "`$ARRAY`"
        },
        {
          "name": "googlePayMerchantID",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "jwk",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "merchantLogoUrlPng",
          "type": "`$STRING`"
        },
        {
          "name": "merchantLogoUrlSvg",
          "type": "`$STRING`"
        },
        {
          "name": "merchantName",
          "type": "`$STRING`"
        },
        {
          "name": "merchantShortName",
          "type": "`$STRING`"
        },
        {
          "name": "partialPaymentMethod",
          "type": "`$STRING`"
        },
        {
          "name": "paymentAttempts",
          "type": "`$ARRAY`"
        },
        {
          "name": "paymentMethodsList",
          "type": "`$ARRAY`"
        },
        {
          "name": "paymentProcessor",
          "type": "`$STRING`"
        },
        {
          "name": "paymentProcessorKey",
          "type": "`$STRING`"
        },
        {
          "name": "pispError",
          "type": "`$STRING`"
        },
        {
          "name": "priorityBankID",
          "type": "`$STRING`"
        },
        {
          "name": "status",
          "type": "`$STRING`"
        },
        {
          "name": "stripeAccountID",
          "type": "`$STRING`"
        },
        {
          "name": "title",
          "type": "`$STRING`"
        }
      ],
      "name": "payment_request_minimal",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "amount",
          "type": "`$NUMBER`"
        },
        {
          "name": "amountPending",
          "type": "`$NUMBER`"
        },
        {
          "name": "amountReceived",
          "type": "`$NUMBER`"
        },
        {
          "name": "amountRefunded",
          "type": "`$NUMBER`"
        },
        {
          "name": "currency",
          "type": "`$STRING`"
        },
        {
          "name": "customerID",
          "type": "`$STRING`"
        },
        {
          "name": "paymentRequestID",
          "type": "`$STRING`"
        },
        {
          "name": "payments",
          "type": "`$ARRAY`"
        },
        {
          "name": "pispAuthorizations",
          "type": "`$ARRAY`"
        },
        {
          "name": "requestedAmount",
          "type": "`$NUMBER`"
        },
        {
          "name": "result",
          "type": "`$STRING`"
        }
      ],
      "name": "payment_request_result",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "paymentrequest_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "accountID",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "type": "`$STRING`"
        },
        {
          "name": "allowIncomplete",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "amount",
          "type": "`$NUMBER`"
        },
        {
          "name": "amountMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "approvePayoutUrl",
          "type": "`$STRING`"
        },
        {
          "name": "approverID",
          "type": "`$STRING`"
        },
        {
          "name": "authenticationMethods",
          "type": "`$ARRAY`"
        },
        {
          "name": "authorisations",
          "type": "`$ARRAY`"
        },
        {
          "name": "authorisersCompletedCount",
          "type": "`$INTEGER`"
        },
        {
          "name": "authorisersRequiredCount",
          "type": "`$INTEGER`"
        },
        {
          "name": "batchPayoutID",
          "type": "`$STRING`"
        },
        {
          "name": "beneficiary",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "beneficiaryID",
          "type": "`$STRING`"
        },
        {
          "name": "canAuthorise",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "canProcess",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "canUpdate",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "chargeBearer",
          "type": "`$STRING`"
        },
        {
          "name": "createdBy",
          "type": "`$STRING`"
        },
        {
          "name": "createdByEmailAddress",
          "type": "`$STRING`"
        },
        {
          "name": "currency",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "type": "`$STRING`"
        },
        {
          "name": "currentUserID",
          "type": "`$STRING`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "destination",
          "type": "`$OBJECT`"
        },
        {
          "name": "documents",
          "type": "`$ARRAY`"
        },
        {
          "name": "events",
          "type": "`$ARRAY`"
        },
        {
          "name": "failedPayouts",
          "type": "`$OBJECT`"
        },
        {
          "name": "formattedAmount",
          "type": "`$STRING`"
        },
        {
          "name": "formattedFxDestinationAmount",
          "type": "`$STRING`"
        },
        {
          "name": "formattedSchedule",
          "type": "`$STRING`"
        },
        {
          "name": "formattedScheduleDayOnly",
          "type": "`$STRING`"
        },
        {
          "name": "formattedSourceAccountAvailableBalance",
          "type": "`$STRING`"
        },
        {
          "name": "fxDestinationAmount",
          "type": "`$NUMBER`"
        },
        {
          "name": "fxDestinationAmountMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "fxDestinationCurrency",
          "type": "`$STRING`"
        },
        {
          "name": "fxQuoteExpiresAt",
          "type": "`$STRING`"
        },
        {
          "name": "fxQuoteID",
          "type": "`$STRING`"
        },
        {
          "name": "fxRate",
          "type": "`$NUMBER`"
        },
        {
          "name": "fxUseDestinationAmount",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "hasCurrentUserAuthorised",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "invoiceID",
          "type": "`$STRING`"
        },
        {
          "name": "isArchived",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isFailed",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isSettled",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isSubmitted",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "merchantTokenDescription",
          "type": "`$STRING`"
        },
        {
          "name": "nonce",
          "type": "`$STRING`"
        },
        {
          "name": "paymentProcessor",
          "type": "`$STRING`"
        },
        {
          "name": "paymentRail",
          "type": "`$STRING`"
        },
        {
          "name": "payouts",
          "type": "`$ARRAY`"
        },
        {
          "name": "payrunID",
          "type": "`$STRING`"
        },
        {
          "name": "payrunName",
          "type": "`$STRING`"
        },
        {
          "name": "reason",
          "type": "`$STRING`"
        },
        {
          "name": "rule",
          "type": "`$OBJECT`"
        },
        {
          "name": "scheduleDate",
          "type": "`$STRING`"
        },
        {
          "name": "scheduled",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "sourceAccountAvailableBalance",
          "type": "`$NUMBER`"
        },
        {
          "name": "sourceAccountAvailableBalanceMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "sourceAccountBic",
          "type": "`$STRING`"
        },
        {
          "name": "sourceAccountCurrency",
          "type": "`$STRING`"
        },
        {
          "name": "sourceAccountIban",
          "type": "`$STRING`"
        },
        {
          "name": "sourceAccountIdentifier",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "sourceAccountName",
          "type": "`$STRING`"
        },
        {
          "name": "sourceAccountNumber",
          "type": "`$STRING`"
        },
        {
          "name": "sourceAccountSortcode",
          "type": "`$STRING`"
        },
        {
          "name": "status",
          "type": "`$STRING`"
        },
        {
          "name": "tagIds",
          "type": "`$ARRAY`"
        },
        {
          "name": "tags",
          "type": "`$ARRAY`"
        },
        {
          "name": "theirReference",
          "type": "`$STRING`"
        },
        {
          "name": "topupPayrunID",
          "type": "`$STRING`"
        },
        {
          "name": "transactedAmount",
          "type": "`$NUMBER`"
        },
        {
          "name": "transactedFxAmount",
          "type": "`$NUMBER`"
        },
        {
          "name": "transactedFxRate",
          "type": "`$NUMBER`"
        },
        {
          "name": "type",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "type": "`$STRING`"
        },
        {
          "name": "userID",
          "type": "`$STRING`"
        },
        {
          "name": "yourReference",
          "type": "`$STRING`"
        }
      ],
      "name": "payout",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "amount",
                    "orig": "amount",
                    "reqd": true,
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "param",
                    "name": "destination",
                    "orig": "destination",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "source",
                    "orig": "source",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "accountID",
          "type": "`$STRING`"
        },
        {
          "name": "amount",
          "type": "`$NUMBER`"
        },
        {
          "name": "amountMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "approvePayoutUrl",
          "type": "`$STRING`"
        },
        {
          "name": "approverID",
          "type": "`$STRING`"
        },
        {
          "name": "authenticationMethods",
          "type": "`$ARRAY`"
        },
        {
          "name": "authorisations",
          "type": "`$ARRAY`"
        },
        {
          "name": "authorisersCompletedCount",
          "type": "`$INTEGER`"
        },
        {
          "name": "authorisersRequiredCount",
          "type": "`$INTEGER`"
        },
        {
          "name": "batchPayoutID",
          "type": "`$STRING`"
        },
        {
          "name": "beneficiary",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "canAuthorise",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "canProcess",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "canUpdate",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "chargeBearer",
          "type": "`$STRING`"
        },
        {
          "name": "createdBy",
          "type": "`$STRING`"
        },
        {
          "name": "createdByEmailAddress",
          "type": "`$STRING`"
        },
        {
          "name": "currency",
          "type": "`$STRING`"
        },
        {
          "name": "currentUserID",
          "type": "`$STRING`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "destination",
          "type": "`$OBJECT`"
        },
        {
          "name": "documents",
          "type": "`$ARRAY`"
        },
        {
          "name": "events",
          "type": "`$ARRAY`"
        },
        {
          "name": "formattedAmount",
          "type": "`$STRING`"
        },
        {
          "name": "formattedFxDestinationAmount",
          "type": "`$STRING`"
        },
        {
          "name": "formattedSchedule",
          "type": "`$STRING`"
        },
        {
          "name": "formattedScheduleDayOnly",
          "type": "`$STRING`"
        },
        {
          "name": "formattedSourceAccountAvailableBalance",
          "type": "`$STRING`"
        },
        {
          "name": "fxDestinationAmount",
          "type": "`$NUMBER`"
        },
        {
          "name": "fxDestinationAmountMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "fxDestinationCurrency",
          "type": "`$STRING`"
        },
        {
          "name": "fxQuoteExpiresAt",
          "type": "`$STRING`"
        },
        {
          "name": "fxQuoteID",
          "type": "`$STRING`"
        },
        {
          "name": "fxRate",
          "type": "`$NUMBER`"
        },
        {
          "name": "fxUseDestinationAmount",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "hasCurrentUserAuthorised",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "invoiceID",
          "type": "`$STRING`"
        },
        {
          "name": "isArchived",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isFailed",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isSettled",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isSubmitted",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "merchantTokenDescription",
          "type": "`$STRING`"
        },
        {
          "name": "nonce",
          "type": "`$STRING`"
        },
        {
          "name": "paymentProcessor",
          "type": "`$STRING`"
        },
        {
          "name": "paymentRail",
          "type": "`$STRING`"
        },
        {
          "name": "payrunID",
          "type": "`$STRING`"
        },
        {
          "name": "payrunName",
          "type": "`$STRING`"
        },
        {
          "name": "rule",
          "type": "`$OBJECT`"
        },
        {
          "name": "scheduleDate",
          "type": "`$STRING`"
        },
        {
          "name": "scheduled",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "sourceAccountAvailableBalance",
          "type": "`$NUMBER`"
        },
        {
          "name": "sourceAccountAvailableBalanceMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "sourceAccountBic",
          "type": "`$STRING`"
        },
        {
          "name": "sourceAccountCurrency",
          "type": "`$STRING`"
        },
        {
          "name": "sourceAccountIban",
          "type": "`$STRING`"
        },
        {
          "name": "sourceAccountIdentifier",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "sourceAccountName",
          "type": "`$STRING`"
        },
        {
          "name": "sourceAccountNumber",
          "type": "`$STRING`"
        },
        {
          "name": "sourceAccountSortcode",
          "type": "`$STRING`"
        },
        {
          "name": "status",
          "type": "`$STRING`"
        },
        {
          "name": "tags",
          "type": "`$ARRAY`"
        },
        {
          "name": "theirReference",
          "type": "`$STRING`"
        },
        {
          "name": "topupPayrunID",
          "type": "`$STRING`"
        },
        {
          "name": "transactedAmount",
          "type": "`$NUMBER`"
        },
        {
          "name": "transactedFxAmount",
          "type": "`$NUMBER`"
        },
        {
          "name": "transactedFxRate",
          "type": "`$NUMBER`"
        },
        {
          "name": "type",
          "type": "`$STRING`"
        },
        {
          "name": "userID",
          "type": "`$STRING`"
        },
        {
          "name": "yourReference",
          "type": "`$STRING`"
        }
      ],
      "name": "payout_keyset_page",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "from_date_utc",
                    "orig": "from_date_utc",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "from_date_utc",
                    "orig": "from_date_utc",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "from_date_utc",
                    "orig": "from_date_utc",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
      "fields": [],
      "name": "payout_metric",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "currency",
                    "orig": "currency",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_archived",
                    "orig": "include_archived",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "tag",
                    "orig": "tag",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "payrun": {
      "fields": [
        {
          "name": "authorisationDate",
          "type": "`$STRING`"
        },
        {
          "name": "authorisations",
          "type": "`$ARRAY`"
        },
        {
          "name": "authorisersCompletedCount",
          "type": "`$INTEGER`"
        },
        {
          "name": "authorisersRequiredCount",
          "type": "`$INTEGER`"
        },
        {
          "name": "batchPayoutID",
          "type": "`$STRING`"
        },
        {
          "name": "canAuthorise",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "canDelete",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "canEdit",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "events",
          "type": "`$ARRAY`"
        },
        {
          "name": "hasCurrentUserAuthorised",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "invoices",
          "type": "`$ARRAY`"
        },
        {
          "name": "invoicesMinimal",
          "type": "`$ARRAY`"
        },
        {
          "name": "isArchived",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "lastUpdatedBy",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "type": "`$STRING`"
        },
        {
          "name": "nonce",
          "type": "`$STRING`"
        },
        {
          "name": "notes",
          "type": "`$STRING`"
        },
        {
          "name": "payments",
          "type": "`$ARRAY`"
        },
        {
          "name": "payouts",
          "type": "`$ARRAY`"
        },
        {
          "name": "payoutsCount",
          "type": "`$INTEGER`"
        },
        {
          "name": "reason",
          "type": "`$STRING`"
        },
        {
          "name": "scheduleDate",
          "type": "`$STRING`"
        },
        {
          "name": "scheduledDate",
          "type": "`$STRING`"
        },
        {
          "name": "sourceAccounts",
          "type": "`$ARRAY`"
        },
        {
          "name": "status",
          "type": "`$STRING`"
        },
        {
          "name": "totalEur",
          "type": "`$NUMBER`"
        },
        {
          "name": "totalGbp",
          "type": "`$NUMBER`"
        },
        {
          "name": "totalUsd",
          "type": "`$NUMBER`"
        }
      ],
      "name": "payrun",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "only_archived",
                    "orig": "only_archived",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "status",
                    "orig": "status",
                    "type": "`$ARRAY`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "report_result": {
      "fields": [
        {
          "name": "contentType",
          "type": "`$STRING`"
        },
        {
          "name": "contents",
          "type": "`$STRING`"
        },
        {
          "name": "lastCompletedAt",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "reportName",
          "type": "`$STRING`"
        },
        {
          "name": "reportType",
          "type": "`$STRING`"
        },
        {
          "name": "statementNumber",
          "type": "`$INTEGER`"
        }
      ],
      "name": "report_result",
      "op": {
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "statement_number",
                    "reqd": true,
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "param",
                    "name": "report_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "failedRoles",
          "type": "`$OBJECT`"
        },
        {
          "name": "roles",
          "type": "`$ARRAY`"
        }
      ],
      "name": "role",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "account",
          "type": "`$OBJECT`"
        },
        {
          "name": "accountID",
          "type": "`$STRING`"
        },
        {
          "name": "approveUrl",
          "type": "`$STRING`"
        },
        {
          "name": "approverID",
          "type": "`$STRING`"
        },
        {
          "name": "authenticationMethods",
          "type": "`$ARRAY`"
        },
        {
          "name": "authorisations",
          "type": "`$ARRAY`"
        },
        {
          "name": "authorisersCompletedCount",
          "type": "`$INTEGER`"
        },
        {
          "name": "authorisersRequiredCount",
          "type": "`$INTEGER`"
        },
        {
          "name": "canAuthorise",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "createdBy",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "endAt",
          "type": "`$STRING`"
        },
        {
          "name": "hasCurrentUserAuthorised",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "isDisabled",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lastExecutedAt",
          "type": "`$STRING`"
        },
        {
          "name": "lastRunAtTransactionDate",
          "type": "`$STRING`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "type": "`$STRING`"
        },
        {
          "name": "nonce",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "onApprovedWebHookUrl",
          "type": "`$STRING`"
        },
        {
          "name": "onExecutionErrorWebHookUrl",
          "type": "`$STRING`"
        },
        {
          "name": "onExecutionSuccessWebHookUrl",
          "type": "`$STRING`"
        },
        {
          "name": "startAt",
          "type": "`$STRING`"
        },
        {
          "name": "status",
          "type": "`$STRING`"
        },
        {
          "name": "sweepAction",
          "op": {
            "create": {
              "req": true,
              "type": "`$OBJECT`"
            }
          },
          "type": "`$OBJECT`"
        },
        {
          "name": "timeZoneId",
          "type": "`$STRING`"
        },
        {
          "name": "triggerCronExpression",
          "type": "`$STRING`"
        },
        {
          "name": "triggerOnPayIn",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "userID",
          "type": "`$STRING`"
        },
        {
          "name": "webHookSecret",
          "type": "`$STRING`"
        }
      ],
      "name": "rule",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "example": false,
                    "kind": "query",
                    "name": "archived_only",
                    "orig": "archived_only",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "rule_event": {
      "fields": [
        {
          "name": "errorMessage",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "isAuthoriseToEnable",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "message",
          "type": "`$STRING`"
        },
        {
          "name": "rawResponse",
          "type": "`$STRING`"
        },
        {
          "name": "ruleEventType",
          "type": "`$STRING`"
        },
        {
          "name": "ruleID",
          "type": "`$STRING`"
        },
        {
          "name": "user",
          "req": true,
          "type": "`$OBJECT`"
        }
      ],
      "name": "rule_event",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "event_type",
                    "orig": "event_type",
                    "type": "`$ARRAY`"
                  },
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page",
                    "orig": "page",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "size",
                    "orig": "size",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "tag": {
      "fields": [
        {
          "name": "colourHex",
          "type": "`$STRING`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "req": true,
          "type": "`$STRING`"
        }
      ],
      "name": "tag",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "transaction": {
      "fields": [
        {
          "name": "accountID",
          "type": "`$STRING`"
        },
        {
          "name": "accountName",
          "type": "`$STRING`"
        },
        {
          "name": "accountSequenceNumber",
          "type": "`$INTEGER`"
        },
        {
          "name": "addressDetails",
          "type": "`$OBJECT`"
        },
        {
          "name": "amount",
          "type": "`$NUMBER`"
        },
        {
          "name": "amountMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "balance",
          "type": "`$NUMBER`"
        },
        {
          "name": "balanceMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "bookingDateTime",
          "type": "`$STRING`"
        },
        {
          "name": "chargeDetails",
          "type": "`$OBJECT`"
        },
        {
          "name": "content",
          "type": "`$ARRAY`"
        },
        {
          "name": "counterparty",
          "type": "`$OBJECT`"
        },
        {
          "name": "counterpartySummary",
          "type": "`$STRING`"
        },
        {
          "name": "currency",
          "type": "`$STRING`"
        },
        {
          "name": "currencyExchange",
          "type": "`$OBJECT`"
        },
        {
          "name": "date",
          "type": "`$STRING`"
        },
        {
          "name": "description",
          "type": "`$STRING`"
        },
        {
          "name": "enrichment",
          "type": "`$OBJECT`"
        },
        {
          "name": "fxAmount",
          "type": "`$NUMBER`"
        },
        {
          "name": "fxCurrency",
          "type": "`$STRING`"
        },
        {
          "name": "fxRate",
          "type": "`$NUMBER`"
        },
        {
          "name": "grossAmount",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "isoBankTransactionCode",
          "type": "`$OBJECT`"
        },
        {
          "name": "merchant",
          "type": "`$OBJECT`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "pageNumber",
          "type": "`$INTEGER`"
        },
        {
          "name": "pageSize",
          "type": "`$INTEGER`"
        },
        {
          "name": "payeeDetails",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "payerDetails",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "paymentRequestCustomFields",
          "type": "`$OBJECT`"
        },
        {
          "name": "paymentRequestID",
          "type": "`$STRING`"
        },
        {
          "name": "payoutID",
          "type": "`$STRING`"
        },
        {
          "name": "proprietaryBankTransactionCode",
          "type": "`$OBJECT`"
        },
        {
          "name": "rawReference",
          "type": "`$STRING`"
        },
        {
          "name": "reference",
          "type": "`$STRING`"
        },
        {
          "name": "ruleID",
          "type": "`$STRING`"
        },
        {
          "name": "statementReferences",
          "type": "`$ARRAY`"
        },
        {
          "name": "status",
          "type": "`$STRING`"
        },
        {
          "name": "supplementaryData",
          "type": "`$ANY`"
        },
        {
          "name": "tags",
          "type": "`$ARRAY`"
        },
        {
          "name": "theirReference",
          "type": "`$STRING`"
        },
        {
          "name": "totalPages",
          "type": "`$INTEGER`"
        },
        {
          "name": "totalSize",
          "type": "`$INTEGER`"
        },
        {
          "name": "transactionAmount",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "transactionDate",
          "type": "`$STRING`"
        },
        {
          "name": "transactionInformation",
          "type": "`$ARRAY`"
        },
        {
          "name": "transactionMutability",
          "type": "`$STRING`"
        },
        {
          "name": "type",
          "type": "`$STRING`"
        },
        {
          "name": "valueDateTime",
          "type": "`$STRING`"
        },
        {
          "name": "virtualIBAN",
          "type": "`$STRING`"
        },
        {
          "name": "yourReference",
          "type": "`$STRING`"
        }
      ],
      "name": "transaction",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "credit_type",
                    "orig": "credit_type",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "limit",
                    "orig": "limit",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "offset",
                    "orig": "offset",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "sort_inserted_ascending",
                    "orig": "sort_inserted_ascending",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "transaction_from",
                    "orig": "transaction_from",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "transaction_to",
                    "orig": "transaction_to",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "example": false,
                    "kind": "query",
                    "name": "include_child_merchant",
                    "orig": "include_child_merchant",
                    "type": "`$BOOLEAN`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "credit_type",
                    "orig": "credit_type",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "credit_type",
                    "orig": "credit_type",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "credit_type",
                    "orig": "credit_type",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "from_date",
                    "orig": "from_date",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "max_amount",
                    "orig": "max_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "min_amount",
                    "orig": "min_amount",
                    "type": "`$NUMBER`"
                  },
                  {
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "to_date",
                    "orig": "to_date",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "example": 0,
                    "kind": "param",
                    "name": "sequence_number",
                    "orig": "sequence_number",
                    "reqd": true,
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "param",
                    "name": "transaction_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "kind": "query",
                    "name": "tag_id",
                    "orig": "tag_id",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "clientSessionTimeouts",
          "type": "`$ARRAY`"
        },
        {
          "name": "emailAddress",
          "op": {
            "update": {
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "firstName",
          "op": {
            "update": {
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "lastName",
          "op": {
            "update": {
              "type": "`$STRING`"
            }
          },
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "passkeyAdded",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "permissions",
          "type": "`$OBJECT`"
        },
        {
          "name": "profile",
          "type": "`$STRING`"
        },
        {
          "name": "rolesWithScope",
          "type": "`$ARRAY`"
        },
        {
          "name": "twoFactorEnabled",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "userInviteID",
          "type": "`$STRING`"
        }
      ],
      "name": "user",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "authorisationStatus",
          "type": "`$OBJECT`"
        },
        {
          "name": "failedUserInvites",
          "type": "`$OBJECT`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "initialRoleID",
          "type": "`$STRING`"
        },
        {
          "name": "inviteeEmailAddress",
          "op": {
            "create": {
              "req": true,
              "type": "`$STRING`"
            }
          },
          "type": "`$STRING`"
        },
        {
          "name": "inviteeFirstName",
          "type": "`$STRING`"
        },
        {
          "name": "inviteeLastName",
          "type": "`$STRING`"
        },
        {
          "name": "inviterEmailAddress",
          "type": "`$STRING`"
        },
        {
          "name": "inviterFirstName",
          "type": "`$STRING`"
        },
        {
          "name": "inviterLastName",
          "type": "`$STRING`"
        },
        {
          "name": "isAuthorised",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isInviteeRegistered",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lastInvited",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "merchantName",
          "type": "`$STRING`"
        },
        {
          "name": "message",
          "type": "`$STRING`"
        },
        {
          "name": "registrationUrl",
          "type": "`$STRING`"
        },
        {
          "name": "sendInviteEmail",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "status",
          "type": "`$STRING`"
        },
        {
          "name": "user",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "userID",
          "type": "`$STRING`"
        },
        {
          "name": "userInvites",
          "type": "`$ARRAY`"
        }
      ],
      "name": "user_invite",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            },
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ],
                "query": [
                  {
                    "example": 1,
                    "kind": "query",
                    "name": "page_number",
                    "orig": "page_number",
                    "type": "`$INTEGER`"
                  },
                  {
                    "example": 20,
                    "kind": "query",
                    "name": "page_size",
                    "orig": "page_size",
                    "type": "`$INTEGER`"
                  },
                  {
                    "kind": "query",
                    "name": "search",
                    "orig": "search",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "sort",
                    "orig": "sort",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "userinvite_id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "accountName",
          "type": "`$STRING`"
        },
        {
          "name": "accountSupplierName",
          "type": "`$STRING`"
        },
        {
          "name": "availableBalance",
          "type": "`$NUMBER`"
        },
        {
          "name": "availableBalanceMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "balance",
          "type": "`$NUMBER`"
        },
        {
          "name": "balanceMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "bankName",
          "type": "`$STRING`"
        },
        {
          "name": "consentID",
          "type": "`$STRING`"
        },
        {
          "name": "createdBy",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "createdByDisplayName",
          "type": "`$STRING`"
        },
        {
          "name": "currency",
          "type": "`$STRING`"
        },
        {
          "name": "defaultPaymentRail",
          "type": "`$STRING`"
        },
        {
          "name": "displayName",
          "type": "`$STRING`"
        },
        {
          "name": "expiryDate",
          "type": "`$STRING`"
        },
        {
          "name": "externalAccountIcon",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "identifier",
          "req": true,
          "type": "`$OBJECT`"
        },
        {
          "name": "inserted",
          "type": "`$STRING`"
        },
        {
          "name": "isArchived",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isConnectedAccount",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isDefault",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isTrustAccount",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "isVirtual",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "lastTransaction",
          "type": "`$OBJECT`"
        },
        {
          "name": "lastUpdated",
          "type": "`$STRING`"
        },
        {
          "name": "merchantID",
          "type": "`$STRING`"
        },
        {
          "name": "merchantName",
          "type": "`$STRING`"
        },
        {
          "name": "name",
          "req": true,
          "type": "`$STRING`"
        },
        {
          "name": "physicalAccountID",
          "type": "`$STRING`"
        },
        {
          "name": "rules",
          "type": "`$ARRAY`"
        },
        {
          "name": "submittedPayoutsBalance",
          "type": "`$NUMBER`"
        },
        {
          "name": "submittedPayoutsBalanceMinorUnits",
          "type": "`$INTEGER`"
        },
        {
          "name": "summary",
          "type": "`$STRING`"
        },
        {
          "name": "supplierSepaInstantStatus",
          "type": "`$STRING`"
        },
        {
          "name": "xeroBankFeedConnectionStatus",
          "type": "`$STRING`"
        },
        {
          "name": "xeroBankFeedLastSyncedAt",
          "type": "`$STRING`"
        },
        {
          "name": "xeroBankFeedSyncLastFailedAt",
          "type": "`$STRING`"
        },
        {
          "name": "xeroBankFeedSyncLastFailureReason",
          "type": "`$STRING`"
        },
        {
          "name": "xeroBankFeedSyncStatus",
          "type": "`$STRING`"
        },
        {
          "name": "xeroUnsynchronisedTransactionsCount",
          "type": "`$INTEGER`"
        }
      ],
      "name": "virtual",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "account_id",
                    "orig": "account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "virtual_account_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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
          "name": "destinationUrl",
          "type": "`$STRING`"
        },
        {
          "name": "emailAddress",
          "type": "`$STRING`"
        },
        {
          "name": "failedNotificationEmailAddress",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "isActive",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "merchantID",
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
          "type": "`$STRING`"
        },
        {
          "name": "notificationMethod",
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
          "type": "`$STRING`"
        },
        {
          "name": "resourceTypes",
          "type": "`$ARRAY`"
        },
        {
          "name": "retry",
          "type": "`$BOOLEAN`"
        },
        {
          "name": "secret",
          "type": "`$STRING`"
        },
        {
          "name": "version",
          "type": "`$INTEGER`"
        }
      ],
      "name": "webhook",
      "op": {
        "create": {
          "input": "data",
          "name": "create",
          "points": [
            {
              "args": {},
              "kind": "http",
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
              }
            }
          ]
        },
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "load": {
          "input": "data",
          "name": "load",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "param",
                    "name": "merchant_id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            },
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "merchant_id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "remove": {
          "input": "data",
          "name": "remove",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
        },
        "update": {
          "input": "data",
          "name": "update",
          "points": [
            {
              "args": {
                "params": [
                  {
                    "kind": "param",
                    "name": "id",
                    "orig": "id",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
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
              }
            }
          ]
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

