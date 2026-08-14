package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "Nofrixion",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://api-sandbox.nofrixion.com",
			"auth": map[string]any{
				"prefix": "Bearer",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"account": map[string]any{},
				"batch": map[string]any{},
				"beneficiary": map[string]any{},
				"beneficiary_group": map[string]any{},
				"card": map[string]any{},
				"card_customer_token": map[string]any{},
				"card_payment": map[string]any{},
				"card_public_key": map[string]any{},
				"consent": map[string]any{},
				"currency": map[string]any{},
				"direct_debit_batch_submit": map[string]any{},
				"fx_rate": map[string]any{},
				"i_payment": map[string]any{},
				"mandate": map[string]any{},
				"merchant": map[string]any{},
				"merchant_authorisation_setting": map[string]any{},
				"merchant_direct_debit_mandate_page": map[string]any{},
				"merchant_pay_by_bank_setting": map[string]any{},
				"merchant_payment_request_template": map[string]any{},
				"merchant_token": map[string]any{},
				"metadata": map[string]any{},
				"no_frixion_version": map[string]any{},
				"open_banking": map[string]any{},
				"payeeverification": map[string]any{},
				"payment": map[string]any{},
				"payment_account": map[string]any{},
				"payment_account_minimal": map[string]any{},
				"payment_initiation": map[string]any{},
				"payment_request": map[string]any{},
				"payment_request_event": map[string]any{},
				"payment_request_metric": map[string]any{},
				"payment_request_minimal": map[string]any{},
				"payment_request_result": map[string]any{},
				"payout": map[string]any{},
				"payout_keyset_page": map[string]any{},
				"payout_metric": map[string]any{},
				"payrun": map[string]any{},
				"report": map[string]any{},
				"report_result": map[string]any{},
				"role": map[string]any{},
				"rule": map[string]any{},
				"rule_event": map[string]any{},
				"tag": map[string]any{},
				"token": map[string]any{},
				"transaction": map[string]any{},
				"user": map[string]any{},
				"user_invite": map[string]any{},
				"virtual": map[string]any{},
				"webhook": map[string]any{},
			},
		},
		"entity": map[string]any{
			"account": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "accountBalances",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "accountID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountIdentifications",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "accountName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountNames",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "accountSupplierName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountType",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "availableBalance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "availableBalanceMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "balance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "balanceMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "bankName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "consentID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "consolidatedAccountInformation",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "createdBy",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "createdByDisplayName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "defaultPaymentRail",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "details",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "displayName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "expiryDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "externalAccountIcon",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "format",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fromDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "identifier",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isConnectedAccount",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isDefault",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isTrustAccount",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isVirtual",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastTransaction",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nickname",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "physicalAccountID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "roleIDs",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "rules",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "submittedPayoutsBalance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "submittedPayoutsBalanceMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "summary",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierPhysicalAccountID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierSepaInstantStatus",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "toDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "usageType",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedConnectionStatus",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedLastSyncedAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedSyncLastFailedAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedSyncLastFailureReason",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedSyncStatus",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroUnsynchronisedTransactionsCount",
						"type": "`$INTEGER`",
					},
				},
				"name": "account",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "currency",
											"orig": "currency",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/accounts/{accountID}/{currency}",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"{account_id}",
									"{currency}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"currency",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/accounts/{accountID}/statements",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"{account_id}",
									"statements",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
									},
								},
								"select": map[string]any{
									"$action": "statement",
									"exist": []any{
										"account_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/accounts",
								"parts": []any{
									"api",
									"v1",
									"accounts",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "connected_account",
											"orig": "connected_account",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_archived",
											"orig": "include_archived",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "only_connect_account",
											"orig": "only_connect_account",
											"type": "`$BOOLEAN`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/accounts",
								"parts": []any{
									"api",
									"v1",
									"accounts",
								},
								"select": map[string]any{
									"exist": []any{
										"connected_account",
										"include_archived",
										"merchant_id",
										"only_connect_account",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "connected_account",
											"orig": "connected_account",
											"type": "`$BOOLEAN`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/accounts",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"accounts",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"connected_account",
										"merchant_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "connected_account",
											"orig": "connected_account",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_archived",
											"orig": "include_archived",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_child_merchant",
											"orig": "include_child_merchant",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "only_archived",
											"orig": "only_archived",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "only_connect_account",
											"orig": "only_connect_account",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/accounts/export",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"export",
								},
								"select": map[string]any{
									"$action": "export",
									"exist": []any{
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
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "credit_type",
											"orig": "credit_type",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "max_amount",
											"orig": "max_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "min_amount",
											"orig": "min_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/accounts/{accountID}/transactions/export",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"{account_id}",
									"transactions",
									"export",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
									},
								},
								"select": map[string]any{
									"$action": "transaction_export",
									"exist": []any{
										"account_id",
										"credit_type",
										"from_date",
										"max_amount",
										"min_amount",
										"page_number",
										"page_size",
										"search",
										"sort",
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/accounts/{accountID}/statements/{id}",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"{account_id}",
									"statements",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/accounts/{accountID}",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"accounts",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "id",
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"merchant_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/accounts/{accountID}",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/openbanking/accounts/{id}",
								"parts": []any{
									"api",
									"v1",
									"openbanking",
									"accounts",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/accounts/statements",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"statements",
								},
								"select": map[string]any{
									"$action": "statement",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/accounts/archive/{id}",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"archive",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/accounts/statements",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"statements",
								},
								"select": map[string]any{
									"$action": "statement",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "amount",
											"orig": "amount",
											"reqd": true,
											"type": "`$NUMBER`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/accounts/{accountID}/topup/{amount}",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"{account_id}",
									"topup",
									"{amount}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"amount",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/accounts/unarchive/{id}",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"unarchive",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/accounts/{id}",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"account",
						},
						[]any{
							"merchant",
						},
						[]any{
							"account",
							"topup",
						},
					},
				},
			},
			"batch": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "approveUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payouts",
						"type": "`$ARRAY`",
					},
				},
				"name": "batch",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/payouts/batch",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"batch",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/payouts/batch/{id}",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"batch",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"beneficiary": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "approvalCallbackUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "authenticationMethods",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisations",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisersCompletedCount",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "authorisersRequiredCount",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "beneficiaries",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "beneficiaryEvents",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "canAuthorise",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "canUpdate",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "createdBy",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "createdByEmailAddress",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"op": map[string]any{
							"update": map[string]any{
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "destination",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$OBJECT`",
							},
						},
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "failedBeneficiaries",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "hasCurrentUserAuthorised",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isEnabled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastAuthorised",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"op": map[string]any{
							"update": map[string]any{
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nonce",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountIDs",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "sourceAccounts",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "theirReference",
						"type": "`$STRING`",
					},
				},
				"name": "beneficiary",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/beneficiaries/authorise/{id}",
								"parts": []any{
									"api",
									"v1",
									"beneficiaries",
									"authorise",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/beneficiaries",
								"parts": []any{
									"api",
									"v1",
									"beneficiaries",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/beneficiaries/batchcreate",
								"parts": []any{
									"api",
									"v1",
									"beneficiaries",
									"batchcreate",
								},
								"select": map[string]any{
									"$action": "batchcreate",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_disabled",
											"orig": "include_disabled",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "source_account_id",
											"orig": "source_account_id",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/beneficiaries",
								"parts": []any{
									"api",
									"v1",
									"beneficiaries",
								},
								"select": map[string]any{
									"exist": []any{
										"currency",
										"include_disabled",
										"merchant_id",
										"page_number",
										"page_size",
										"search",
										"sort",
										"source_account_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_disabled",
											"orig": "include_disabled",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "source_account_id",
											"orig": "source_account_id",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/beneficiaries",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"beneficiaries",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"currency",
										"include_disabled",
										"merchant_id",
										"page_number",
										"page_size",
										"search",
										"sort",
										"source_account_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_disabled",
											"orig": "include_disabled",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/beneficiaries/export",
								"parts": []any{
									"api",
									"v1",
									"beneficiaries",
									"export",
								},
								"select": map[string]any{
									"$action": "export",
									"exist": []any{
										"currency",
										"include_disabled",
										"merchant_id",
										"page_number",
										"page_size",
										"search",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/beneficiaries/{id}",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"beneficiaries",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"merchant_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/beneficiaries/{id}",
								"parts": []any{
									"api",
									"v1",
									"beneficiaries",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/beneficiaries/{id}",
								"parts": []any{
									"api",
									"v1",
									"beneficiaries",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/beneficiaries/disable/{id}",
								"parts": []any{
									"api",
									"v1",
									"beneficiaries",
									"disable",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/beneficiaries/enable/{id}",
								"parts": []any{
									"api",
									"v1",
									"beneficiaries",
									"enable",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/beneficiaries/{id}",
								"parts": []any{
									"api",
									"v1",
									"beneficiaries",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"merchant",
						},
					},
				},
			},
			"beneficiary_group": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "groupMembers",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "groupName",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"req": true,
						"type": "`$STRING`",
					},
				},
				"name": "beneficiary_group",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/beneficiarygroups",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"beneficiarygroups",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"merchant_id",
										"page_number",
										"page_size",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"merchant",
						},
					},
				},
			},
			"card": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "authorizedAmount",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currencyCode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isPayerAuthenticationRequired",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isSoftDecline",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "payerAuthenticationAccessToken",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payerAuthenticationMerchantData",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payerAuthenticationUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payerAuthenticationWindowHeight",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "payerAuthenticationWindowWidth",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "paymentRequestCallbackUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRequestID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "requestID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "responseCode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "responseType",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "threeDSRedirectUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "transactionID",
						"type": "`$STRING`",
					},
				},
				"name": "card",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/paymentrequests/{id}/card",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"card",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "paymentrequest_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"paymentrequest",
						},
					},
				},
			},
			"card_customer_token": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "cardType",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerEmailAddress",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "expiryMonth",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "expiryYear",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastFourDigits",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "maskedCardNumber",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRequestID",
						"type": "`$STRING`",
					},
				},
				"name": "card_customer_token",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "customer_email_address",
											"orig": "customer_email_address",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/paymentrequests/card/customertokens/{merchantID}/{customerEmailAddress}",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"card",
									"customertokens",
									"{merchant_id}",
									"{customer_email_address}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"customerEmailAddress": "customer_email_address",
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"customer_email_address",
										"merchant_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "customer_email_address",
											"orig": "customer_email_address",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/paymentrequests/card/customertokens/{customerEmailAddress}",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"card",
									"customertokens",
									"{customer_email_address}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"customerEmailAddress": "customer_email_address",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"customer_email_address",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "customer_email_address",
											"orig": "customer_email_address",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/paymentrequests/card/customertokens/removeall/{merchantID}/{customerEmailAddress}",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"card",
									"customertokens",
									"removeall",
									"{merchant_id}",
									"{customer_email_address}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"customerEmailAddress": "customer_email_address",
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"customer_email_address",
										"merchant_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "customer_email_address",
											"orig": "customer_email_address",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/paymentrequests/card/customertokens/removeall/{customerEmailAddress}",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"card",
									"customertokens",
									"removeall",
									"{customer_email_address}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"customerEmailAddress": "customer_email_address",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"customer_email_address",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/paymentrequests/card/customertokens/{id}",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"card",
									"customertokens",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"removeall",
						},
						[]any{
							"customertoken",
						},
					},
				},
			},
			"card_payment": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "authorizedAmount",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currencyCode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isPayerAuthenticationRequired",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isSoftDecline",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "payerAuthenticationAccessToken",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payerAuthenticationMerchantData",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payerAuthenticationUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payerAuthenticationWindowHeight",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "payerAuthenticationWindowWidth",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "paymentRequestCallbackUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRequestID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "requestID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "responseCode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "responseType",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "threeDSRedirectUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "transactionID",
						"type": "`$STRING`",
					},
				},
				"name": "card_payment",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "partial_refund_amount",
											"orig": "partial_refund_amount",
											"reqd": true,
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/paymentrequests/{id}/card/refund/{partialRefundAmount}",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"card",
									"refund",
									"{partial_refund_amount}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "paymentrequest_id",
										"partialRefundAmount": "partial_refund_amount",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"partial_refund_amount",
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/paymentrequests/{id}/card/capture",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"card",
									"capture",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "paymentrequest_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/paymentrequests/{id}/card/paywithtoken",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"card",
									"paywithtoken",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "paymentrequest_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/paymentrequests/{id}/card/void",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"card",
									"void",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "paymentrequest_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/paymentrequests/{id}/card/voidpaymentrequest",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"card",
									"voidpaymentrequest",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "paymentrequest_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"paymentrequest",
						},
						[]any{
							"paymentrequest",
							"refund",
						},
					},
				},
			},
			"card_public_key": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "jwt",
						"type": "`$STRING`",
					},
				},
				"name": "card_public_key",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/paymentrequests/{id}/card/publickey",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"card",
									"publickey",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "paymentrequest_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"paymentrequest",
						},
					},
				},
			},
			"consent": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "authorisationUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "callbackUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "consentID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "emailAddress",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "expiryDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "failureCallbackUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "institutionID",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isConnectedAccounts",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isEnabled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "merchantID",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "provider",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "successWebHookUrl",
						"type": "`$STRING`",
					},
				},
				"name": "consent",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/openbanking/consents",
								"parts": []any{
									"api",
									"v1",
									"openbanking",
									"consents",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "email",
											"orig": "email",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/openbanking/consents/{merchantID}/{email}",
								"parts": []any{
									"api",
									"v1",
									"openbanking",
									"consents",
									"{merchant_id}",
									"{email}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"email",
										"merchant_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/openbanking/consents/{id}",
								"parts": []any{
									"api",
									"v1",
									"openbanking",
									"consents",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/openbanking/consents/{id}",
								"parts": []any{
									"api",
									"v1",
									"openbanking",
									"consents",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/api/v1/openbanking/consents/{id}",
								"parts": []any{
									"api",
									"v1",
									"openbanking",
									"consents",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"consent",
						},
					},
				},
			},
			"currency": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "code",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "decimals",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "isFiat",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "iso4217AlphaCode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "iso4217NumericCode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "symbol",
						"type": "`$STRING`",
					},
				},
				"name": "currency",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "capability",
											"orig": "capability",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/currencies",
								"parts": []any{
									"api",
									"v1",
									"currencies",
								},
								"select": map[string]any{
									"exist": []any{
										"capability",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"direct_debit_batch_submit": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "failedSubmissions",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "successfulSubmissions",
						"type": "`$ARRAY`",
					},
				},
				"name": "direct_debit_batch_submit",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/paymentrequests/directdebit/batchsubmit",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"directdebit",
									"batchsubmit",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"fx_rate": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "destinationCurrency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "exchangeRate",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "expiryTime",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "quoteID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceCurrency",
						"type": "`$STRING`",
					},
				},
				"name": "fx_rate",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "destination",
											"orig": "destination",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "source",
											"orig": "source",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/payouts/fxallheldrates/{source}/{destination}",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"fxallheldrates",
									"{source}",
									"{destination}",
								},
								"select": map[string]any{
									"exist": []any{
										"destination",
										"source",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "destination",
											"orig": "destination",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "source",
											"orig": "source",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "valid_for_minute",
											"orig": "valid_for_minute",
											"reqd": true,
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/payouts/fxheldrate/{source}/{destination}/{validForMinutes}",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"fxheldrate",
									"{source}",
									"{destination}",
									"{valid_for_minute}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"validForMinutes": "valid_for_minute",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"destination",
										"source",
										"valid_for_minute",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"fxallheldrate",
						},
						[]any{
							"fxheldrate",
						},
					},
				},
			},
			"i_payment": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "paymentRequestID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "responseType",
						"type": "`$STRING`",
					},
				},
				"name": "i_payment",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/paymentrequests/payondemand",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"payondemand",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"mandate": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "accountNumber",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "addressLine1",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "addressLine2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "approvedAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "city",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "countryCode",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerAccountNumber",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerCity",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerCountryCode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerCountryName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerEmailAddress",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerFirstName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerIban",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerLastName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerSortCode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "emailAddress",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "firstName",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "iban",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isRecurring",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastName",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "postalCode",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reference",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sortCode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierBankAccountID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierCustomerID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierMandateID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierStatus",
						"type": "`$STRING`",
					},
				},
				"name": "mandate",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/mandates",
								"parts": []any{
									"api",
									"v1",
									"mandates",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/mandates/{id}",
								"parts": []any{
									"api",
									"v1",
									"mandates",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"merchant": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "accountCurrencies",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "canHaveTrustAccounts",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardPaymentProcessor",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "companyID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "displayQrOnHostedPay",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "hostedPayVersion",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isBlocked",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isExited",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isSuspended",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "jurisdiction",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "logoUrlPng",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "logoUrlSvg",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantCategoryCode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notes",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "parentMerchant",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "paymentAccountLimit",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "paymentAccounts",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "reason",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shortName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supportedPaymentMethodsList",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "suspensionReason",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tags",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "timeZoneId",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tradingName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "webHookLimit",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "yourRoleName",
						"type": "`$STRING`",
					},
				},
				"name": "merchant",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/childmerchants",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"childmerchants",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"$action": "childmerchant",
									"exist": []any{
										"merchant_id",
										"page_number",
										"page_size",
										"search",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_suspended",
											"orig": "include_suspended",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 10,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/paged",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"paged",
								},
								"select": map[string]any{
									"$action": "paged",
									"exist": []any{
										"include_suspended",
										"page_number",
										"page_size",
										"search",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants",
								"parts": []any{
									"api",
									"v1",
									"merchants",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/metadata/whoamimerchant",
								"parts": []any{
									"api",
									"v1",
									"metadata",
									"whoamimerchant",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/metadata/whoamimerchantsigned",
								"parts": []any{
									"api",
									"v1",
									"metadata",
									"whoamimerchantsigned",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/metadata/whoamimerchantwhitelist",
								"parts": []any{
									"api",
									"v1",
									"metadata",
									"whoamimerchantwhitelist",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "max_amount",
											"orig": "max_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "min_amount",
											"orig": "min_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "status",
											"orig": "status",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "tag",
											"orig": "tag",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/payouts/export",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"payouts",
									"export",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"$action": "payout_export",
									"exist": []any{
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
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_disabled",
											"orig": "include_disabled",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/beneficiaries/export",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"beneficiaries",
									"export",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"$action": "beneficiary_export",
									"exist": []any{
										"currency",
										"include_disabled",
										"merchant_id",
										"page_number",
										"page_size",
										"search",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "user_id",
											"orig": "user_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/merchants/{merchantId}/users/{userId}",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{id}",
									"users",
									"{user_id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantId": "id",
										"userId": "user_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"user_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "tag_id",
											"orig": "tag_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/merchants/{merchantID}/tags/{tagID}",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"tags",
									"{tag_id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
										"tagID": "tag_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"merchant_id",
										"tag_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/merchants/{merchantID}",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/merchants/{merchantId}/suspend",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{id}",
									"suspend",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantId": "id",
									},
								},
								"select": map[string]any{
									"$action": "suspend",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"merchant",
						},
						[]any{
							"user",
						},
						[]any{
							"merchant",
							"tag",
						},
					},
				},
			},
			"merchant_authorisation_setting": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "amountLower",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountUpper",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "authorisationType",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "beneficiariesOnly",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastEditorCantAuthorise",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "numberOfAuthorisers",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "roleSettings",
						"type": "`$ARRAY`",
					},
				},
				"name": "merchant_authorisation_setting",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/authorisationsettings",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"authorisationsettings",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"merchant_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"merchant",
						},
					},
				},
			},
			"merchant_direct_debit_mandate_page": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "approvedAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerAccountNumber",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerCity",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerCountryCode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerCountryName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerEmailAddress",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerFirstName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerIban",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerLastName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerSortCode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isRecurring",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reference",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierBankAccountID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierCustomerID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierMandateID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierStatus",
						"type": "`$STRING`",
					},
				},
				"name": "merchant_direct_debit_mandate_page",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "mandate_i_d",
											"orig": "mandate_i_d",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "max_amount",
											"orig": "max_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "min_amount",
											"orig": "min_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/mandates",
								"parts": []any{
									"api",
									"v1",
									"mandates",
								},
								"select": map[string]any{
									"exist": []any{
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
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"merchant_pay_by_bank_setting": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "bankCountryCodes",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "bankID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "bankName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "businessInstitutionID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "logo",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "message",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "messageImageUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "order",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "personalInstitutionID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "processor",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "warningHeading",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "warningMessage",
						"type": "`$STRING`",
					},
				},
				"name": "merchant_pay_by_bank_setting",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "country_code",
											"orig": "country_code",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "PIS",
											"kind": "query",
											"name": "open_banking_operation",
											"orig": "open_banking_operation",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/banksettings",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"banksettings",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"country_code",
										"currency",
										"merchant_id",
										"open_banking_operation",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"merchant",
						},
					},
				},
			},
			"merchant_payment_request_template": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "bankPaymentOptions",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "cardPaymentAddressOptions",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "cardPaymentCaptureOptions",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "customFields",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "defaultFields",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "description",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notificationOptions",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "paymentMethods",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "paymentTerms",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "priorityBankOptions",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "template",
						"req": true,
						"type": "`$OBJECT`",
					},
				},
				"name": "merchant_payment_request_template",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/paymentrequests/{merchantID}/templates",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{merchant_id}",
									"templates",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"merchant_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "template_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/paymentrequests/{merchantID}/templates/{templateID}",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"templates",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "paymentrequest_id",
										"templateID": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "template_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/paymentrequests/{merchantID}/templates/{templateID}",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"templates",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "paymentrequest_id",
										"templateID": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "template_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/paymentrequests/{merchantID}/templates/{templateID}",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"templates",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "paymentrequest_id",
										"templateID": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"paymentrequest",
						},
					},
				},
			},
			"merchant_token": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "authenticationMethods",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisations",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisersCompletedCount",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "authorisersRequiredCount",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "canAuthorise",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "description",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "expiresAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hasCurrentUserAuthorised",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "hmacAlgorithm",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ipAddressWhitelist",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isEnabled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastAuthorised",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nonce",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "permissionTypes",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "requestSignatureVersion",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "sharedSecretAlgorithm",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sharedSecretBase64",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "token",
						"type": "`$STRING`",
					},
				},
				"name": "merchant_token",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/tokens",
								"parts": []any{
									"api",
									"v1",
									"tokens",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/tokens",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"tokens",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"merchant_id",
										"page_number",
										"page_size",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/tokens/{id}",
								"parts": []any{
									"api",
									"v1",
									"tokens",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/tokens/{id}",
								"parts": []any{
									"api",
									"v1",
									"tokens",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"merchant",
						},
					},
				},
			},
			"metadata": map[string]any{
				"fields": []any{},
				"name": "metadata",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "secret",
											"orig": "secret",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "url",
											"orig": "url",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/metadata/problemnotification",
								"parts": []any{
									"api",
									"v1",
									"metadata",
									"problemnotification",
								},
								"select": map[string]any{
									"$action": "problemnotification",
									"exist": []any{
										"id",
										"secret",
										"url",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/metadata/problem",
								"parts": []any{
									"api",
									"v1",
									"metadata",
									"problem",
								},
								"select": map[string]any{
									"$action": "problem",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"no_frixion_version": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "buildVersion",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "majorVersion",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "minorVersion",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "releaseName",
						"type": "`$STRING`",
					},
				},
				"name": "no_frixion_version",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/metadata/version",
								"parts": []any{
									"api",
									"v1",
									"metadata",
									"version",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"open_banking": map[string]any{
				"fields": []any{},
				"name": "open_banking",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/openbanking/account/{accountID}/synchronise",
								"parts": []any{
									"api",
									"v1",
									"openbanking",
									"account",
									"{account_id}",
									"synchronise",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "email",
											"orig": "email",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/openbanking/consents/{merchantID}/{email}",
								"parts": []any{
									"api",
									"v1",
									"openbanking",
									"consents",
									"{merchant_id}",
									"{email}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"email",
										"merchant_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/openbanking/account/{accountID}",
								"parts": []any{
									"api",
									"v1",
									"openbanking",
									"account",
									"{account_id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"account",
						},
						[]any{
							"consent",
						},
					},
				},
			},
			"payeeverification": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "accountName",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountNumber",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "iban",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payeeVerifiedAccountName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "secondaryIdentification",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sortCode",
						"type": "`$STRING`",
					},
				},
				"name": "payeeverification",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/openbanking/payeeverification",
								"parts": []any{
									"api",
									"v1",
									"openbanking",
									"payeeverification",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"payment": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "addresses",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "amount",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$NUMBER`",
							},
						},
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountPending",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountReceived",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountRefunded",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "autoSendReceipt",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "baseOriginUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "callbackUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardAuthorizeOnly",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardCreateToken",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardCreateTokenMode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardIgnoreCVN",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardNoPayerAuthentication",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardProcessorMerchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardStripePaymentIntentID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardStripePaymentIntentSecret",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardTransmitRawDetails",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "createdByUser",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "currency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customFields",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "customerEmailAddress",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "destinationAccount",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "directDebitPayment",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "dueDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "events",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "failureCallbackUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fieldDisplaySettings",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "formattedAmount",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hostedPayCheckoutUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ignoreAddressVerification",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "insertedSortable",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "jwk",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lightningInvoice",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lightningInvoiceExpiresAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantDirectDebitMandateID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantTokenDescription",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notificationEmailAddresses",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notificationRoleIDs",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "orderID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "partialPaymentMethod",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "partialPaymentSteps",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentAttempts",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "paymentMethods",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "paymentProcessor",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payrunID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispAccountID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "priorityBankID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sandboxSettleDelayInSeconds",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "shippingAddress",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "shippingAddressCity",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingAddressCountryCode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingAddressCounty",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingAddressLine1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingAddressLine2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingAddressPostCode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingEmail",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingFirstName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingLastName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingPhone",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "successWebHookUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tagIds",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "tags",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tokenisedCards",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "transactions",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "useHostedPaymentPage",
						"type": "`$BOOLEAN`",
					},
				},
				"name": "payment",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/paymentrequests",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_event",
											"orig": "include_event",
											"type": "`$BOOLEAN`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/paymentrequests/{id}",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"include_event",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "order_id",
											"orig": "order_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/paymentrequests/getbyorderid/{orderID}",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"getbyorderid",
									"{order_id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"orderID": "order_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"order_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/paymentrequests/{id}",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"getbyorderid",
						},
					},
				},
			},
			"payment_account": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "accountName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountSupplierName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "availableBalance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "availableBalanceMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "balance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "balanceMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "bankName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "consentID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "createdBy",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "createdByDisplayName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "defaultPaymentRail",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "displayName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "expiryDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "externalAccountIcon",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "identifier",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isConnectedAccount",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isDefault",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isTrustAccount",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isVirtual",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastTransaction",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "physicalAccountID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "rules",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "submittedPayoutsBalance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "submittedPayoutsBalanceMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "summary",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierSepaInstantStatus",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedConnectionStatus",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedLastSyncedAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedSyncLastFailedAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedSyncLastFailureReason",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedSyncStatus",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroUnsynchronisedTransactionsCount",
						"type": "`$INTEGER`",
					},
				},
				"name": "payment_account",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "connected_account",
											"orig": "connected_account",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_archived",
											"orig": "include_archived",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_child_merchant",
											"orig": "include_child_merchant",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "only_archived",
											"orig": "only_archived",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "only_connect_account",
											"orig": "only_connect_account",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/accounts/paged",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"paged",
								},
								"select": map[string]any{
									"exist": []any{
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
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/accounts/{accountID}/virtual",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"{account_id}",
									"virtual",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"page_number",
										"page_size",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"account",
						},
					},
				},
			},
			"payment_account_minimal": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "accountName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "availableBalance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "balance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "balanceMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "currency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "identifier",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "isArchived",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isConnectedAccount",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "submittedPayoutsBalance",
						"type": "`$NUMBER`",
					},
				},
				"name": "payment_account_minimal",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "connected_account",
											"orig": "connected_account",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_archived",
											"orig": "include_archived",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "only_archived",
											"orig": "only_archived",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "only_connect_account",
											"orig": "only_connect_account",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 10,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/accounts/minimal",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"minimal",
								},
								"select": map[string]any{
									"exist": []any{
										"connected_account",
										"currency",
										"include_archived",
										"merchant_id",
										"only_archived",
										"only_connect_account",
										"page_number",
										"page_size",
										"search",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"payment_initiation": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "paymentInitiationID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRequestCallbackUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRequestID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "redirectUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "responseType",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "specificErrorMessage",
						"type": "`$STRING`",
					},
				},
				"name": "payment_initiation",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/paymentrequests/{id}/pisp",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"pisp",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "paymentrequest_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"paymentrequest",
						},
					},
				},
			},
			"payment_request": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "addresses",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "amount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountPending",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountReceived",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountRefunded",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "autoSendReceipt",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "baseOriginUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "callbackUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardAuthorizeOnly",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardCreateToken",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardCreateTokenMode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardIgnoreCVN",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardProcessorMerchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardStripePaymentIntentID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardStripePaymentIntentSecret",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "createdByUser",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "currency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customFields",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "customerEmailAddress",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "destinationAccount",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "directDebitPayment",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "doSimulateSettlementFailure",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "dueDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "errorDescription",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "events",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "failedPaymentRequests",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "failureCallbackUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fieldDisplaySettings",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "formattedAmount",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hostedPayCheckoutUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ignoreAddressVerification",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "insertedSortable",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "institution",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "jwk",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lightningInvoice",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lightningInvoiceExpiresAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantDirectDebitMandateID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantTokenDescription",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notificationEmailAddresses",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notificationRoleIDs",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "orderID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "partialPaymentMethod",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "partialPaymentSteps",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentAttempts",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "paymentInitiationID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentMethods",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "paymentProcessor",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRequests",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "payrunID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispAccountID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "priorityBankID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sandboxSettleDelayInSeconds",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "shippingAddress",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "successWebHookUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tags",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tokenisedCards",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "transactions",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "useHostedPaymentPage",
						"type": "`$BOOLEAN`",
					},
				},
				"name": "payment_request",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "mandate_id",
											"orig": "mandate_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "submit_after",
											"orig": "submit_after",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/paymentrequests/{id}/directdebit",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"directdebit",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "paymentrequest_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"mandate_id",
										"paymentrequest_id",
										"submit_after",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/paymentrequests/batchcreate",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"batchcreate",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_archived",
											"orig": "include_archived",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "max_amount",
											"orig": "max_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "min_amount",
											"orig": "min_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "payment_method",
											"orig": "payment_method",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "All",
											"kind": "query",
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "tag",
											"orig": "tag",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/paymentrequests",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
								},
								"select": map[string]any{
									"exist": []any{
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
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_archived",
											"orig": "include_archived",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "max_amount",
											"orig": "max_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "min_amount",
											"orig": "min_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "payment_method",
											"orig": "payment_method",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "All",
											"kind": "query",
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "tag",
											"orig": "tag",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/paymentrequests/export",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"export",
								},
								"select": map[string]any{
									"exist": []any{
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
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/paymentrequests/{id}/receipt",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"receipt",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "paymentrequest_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/paymentrequests/{id}",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/paymentrequests/{id}/pisp/sandboxcallback",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"pisp",
									"sandboxcallback",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "paymentrequest_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"paymentrequest",
						},
					},
				},
			},
			"payment_request_event": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "amount",
						"req": true,
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "applePayTransactionID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardAuthorizationResponseID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardExpiryMonth",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "cardExpiryYear",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "cardIssuer",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardIssuerCountry",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardLastFourDigits",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardRequestID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardScheme",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardTokenCustomerID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardTransactionID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "directDebitPaymentID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "directDebitPaymentReference",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "drirectDebitMandateID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "errorMessage",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "errorReason",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "eventType",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lightningInvoice",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lightningRHash",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "originUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentMethodType",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentProcessorName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRequestID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispBankStatus",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispPaymentInitiationID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispPaymentInstitutionName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispPaymentServiceProviderID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispRedirectUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reconciledTransactionID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "refundPayoutID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "walletName",
						"type": "`$STRING`",
					},
				},
				"name": "payment_request_event",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/paymentrequests/{id}/events",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"events",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "paymentrequest_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"paymentrequest",
						},
					},
				},
			},
			"payment_request_metric": map[string]any{
				"fields": []any{},
				"name": "payment_request_metric",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_archived",
											"orig": "include_archived",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "max_amount",
											"orig": "max_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "min_amount",
											"orig": "min_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "payment_method",
											"orig": "payment_method",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "tag",
											"orig": "tag",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/paymentrequests/metrics",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"metrics",
								},
								"select": map[string]any{
									"exist": []any{
										"currency",
										"from_date",
										"include_archived",
										"max_amount",
										"merchant_id",
										"min_amount",
										"payment_method",
										"search",
										"tag",
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"payment_request_minimal": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "amount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountPending",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountReceived",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountRefunded",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "callbackUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardStripePaymentIntentSecret",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "countryCode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customFieldsToDisplay",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "dueDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fieldDisplaySettings",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "googlePayMerchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "jwk",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantLogoUrlPng",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantLogoUrlSvg",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantShortName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "partialPaymentMethod",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentAttempts",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "paymentMethodsList",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "paymentProcessor",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentProcessorKey",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispError",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "priorityBankID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "stripeAccountID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "title",
						"type": "`$STRING`",
					},
				},
				"name": "payment_request_minimal",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/paymentrequests/{id}/minimal",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"minimal",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "paymentrequest_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"paymentrequest",
						},
					},
				},
			},
			"payment_request_result": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "amount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountPending",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountReceived",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountRefunded",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "currency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRequestID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payments",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "pispAuthorizations",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "requestedAmount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "result",
						"type": "`$STRING`",
					},
				},
				"name": "payment_request_result",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "paymentrequest_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/paymentrequests/{id}/result",
								"parts": []any{
									"api",
									"v1",
									"paymentrequests",
									"{paymentrequest_id}",
									"result",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "paymentrequest_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"paymentrequest_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"paymentrequest",
						},
					},
				},
			},
			"payout": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "accountID",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "allowIncomplete",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "amount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "approvePayoutUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "approverID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "authenticationMethods",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisations",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisersCompletedCount",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "authorisersRequiredCount",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "batchPayoutID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "beneficiary",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "beneficiaryID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "canAuthorise",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "canProcess",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "canUpdate",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "chargeBearer",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "createdBy",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "createdByEmailAddress",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currentUserID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "destination",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "documents",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "events",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "failedPayouts",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "formattedAmount",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "formattedFxDestinationAmount",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "formattedSchedule",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "formattedScheduleDayOnly",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "formattedSourceAccountAvailableBalance",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxDestinationAmount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fxDestinationAmountMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "fxDestinationCurrency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxQuoteExpiresAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxQuoteID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxRate",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fxUseDestinationAmount",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "hasCurrentUserAuthorised",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "invoiceID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isFailed",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isSettled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isSubmitted",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantTokenDescription",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nonce",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentProcessor",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRail",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payouts",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "payrunID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payrunName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reason",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "rule",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "scheduleDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "scheduled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "sourceAccountAvailableBalance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "sourceAccountAvailableBalanceMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "sourceAccountBic",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountCurrency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountIban",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountIdentifier",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sourceAccountName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountNumber",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountSortcode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tagIds",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "tags",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "theirReference",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "topupPayrunID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "transactedAmount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "transactedFxAmount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "transactedFxRate",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "type",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "userID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "yourReference",
						"type": "`$STRING`",
					},
				},
				"name": "payout",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/payouts/batch/submit/{id}",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"batch",
									"submit",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/payouts/submit/{id}",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"submit",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/payouts",
								"parts": []any{
									"api",
									"v1",
									"payouts",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/payouts/batchcreate",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"batchcreate",
								},
								"select": map[string]any{
									"$action": "batchcreate",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/payouts/send",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"send",
								},
								"select": map[string]any{
									"$action": "send",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/payouts/sendbeneficiary",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"sendbeneficiary",
								},
								"select": map[string]any{
									"$action": "sendbeneficiary",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_archived",
											"orig": "include_archived",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "max_amount",
											"orig": "max_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "min_amount",
											"orig": "min_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "status",
											"orig": "status",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "tag",
											"orig": "tag",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/payouts",
								"parts": []any{
									"api",
									"v1",
									"payouts",
								},
								"select": map[string]any{
									"exist": []any{
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
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "max_amount",
											"orig": "max_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "min_amount",
											"orig": "min_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "status",
											"orig": "status",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "tag",
											"orig": "tag",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/accounts/{accountID}/payouts",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"{account_id}",
									"payouts",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
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
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "max_amount",
											"orig": "max_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "min_amount",
											"orig": "min_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "status",
											"orig": "status",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "tag",
											"orig": "tag",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/payouts",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"payouts",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
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
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_archived",
											"orig": "include_archived",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "max_amount",
											"orig": "max_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "min_amount",
											"orig": "min_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "status",
											"orig": "status",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "tag",
											"orig": "tag",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/payouts/export",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"export",
								},
								"select": map[string]any{
									"$action": "export",
									"exist": []any{
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
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "amount",
											"orig": "amount",
											"reqd": true,
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "param",
											"name": "destination",
											"orig": "destination",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "source",
											"orig": "source",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/payouts/fxquote/{source}/{destination}/{amount}",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"fxquote",
									"{source}",
									"{destination}",
									"{amount}",
								},
								"select": map[string]any{
									"exist": []any{
										"amount",
										"destination",
										"source",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/payouts/{id}",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/payouts/{id}/proof",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"{id}",
									"proof",
								},
								"select": map[string]any{
									"$action": "proof",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/payouts/{id}",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/payouts/batchdelete",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"batchdelete",
								},
								"select": map[string]any{
									"$action": "batchdelete",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/payouts/cancel/{id}",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"cancel",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/payouts/reject/{id}",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"reject",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/payouts/{id}",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"account",
						},
						[]any{
							"merchant",
						},
						[]any{
							"fxquote",
						},
					},
				},
			},
			"payout_keyset_page": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "accountID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "amount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "approvePayoutUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "approverID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "authenticationMethods",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisations",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisersCompletedCount",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "authorisersRequiredCount",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "batchPayoutID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "beneficiary",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "canAuthorise",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "canProcess",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "canUpdate",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "chargeBearer",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "createdBy",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "createdByEmailAddress",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currentUserID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "destination",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "documents",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "events",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "formattedAmount",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "formattedFxDestinationAmount",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "formattedSchedule",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "formattedScheduleDayOnly",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "formattedSourceAccountAvailableBalance",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxDestinationAmount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fxDestinationAmountMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "fxDestinationCurrency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxQuoteExpiresAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxQuoteID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxRate",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fxUseDestinationAmount",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "hasCurrentUserAuthorised",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "invoiceID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isFailed",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isSettled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isSubmitted",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantTokenDescription",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nonce",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentProcessor",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRail",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payrunID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payrunName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "rule",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "scheduleDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "scheduled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "sourceAccountAvailableBalance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "sourceAccountAvailableBalanceMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "sourceAccountBic",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountCurrency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountIban",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountIdentifier",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sourceAccountName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountNumber",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountSortcode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tags",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "theirReference",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "topupPayrunID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "transactedAmount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "transactedFxAmount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "transactedFxRate",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "userID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "yourReference",
						"type": "`$STRING`",
					},
				},
				"name": "payout_keyset_page",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "from_date_utc",
											"orig": "from_date_utc",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/accounts/{accountID}/payouts/failed",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"{account_id}",
									"payouts",
									"failed",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"from_date_utc",
										"page_size",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "from_date_utc",
											"orig": "from_date_utc",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/payouts/failed",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"payouts",
									"failed",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"from_date_utc",
										"merchant_id",
										"page_size",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "from_date_utc",
											"orig": "from_date_utc",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/payouts/{merchantID}/failed",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"{merchant_id}",
									"failed",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"from_date_utc",
										"merchant_id",
										"page_size",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"account",
						},
						[]any{
							"merchant",
						},
						[]any{
							"payout",
						},
					},
				},
			},
			"payout_metric": map[string]any{
				"fields": []any{},
				"name": "payout_metric",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "currency",
											"orig": "currency",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_archived",
											"orig": "include_archived",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "max_amount",
											"orig": "max_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "min_amount",
											"orig": "min_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "tag",
											"orig": "tag",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/payouts/metrics",
								"parts": []any{
									"api",
									"v1",
									"payouts",
									"metrics",
								},
								"select": map[string]any{
									"exist": []any{
										"currency",
										"from_date",
										"include_archived",
										"max_amount",
										"merchant_id",
										"min_amount",
										"search",
										"tag",
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"payrun": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "authorisationDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "authorisations",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisersCompletedCount",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "authorisersRequiredCount",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "batchPayoutID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "canAuthorise",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "canDelete",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "canEdit",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "events",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "hasCurrentUserAuthorised",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "invoices",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "invoicesMinimal",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "isArchived",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastUpdatedBy",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nonce",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notes",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payments",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "payouts",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "payoutsCount",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "reason",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "scheduleDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "scheduledDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccounts",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "totalEur",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "totalGbp",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "totalUsd",
						"type": "`$NUMBER`",
					},
				},
				"name": "payrun",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/payruns/{id}/request-authorisation",
								"parts": []any{
									"api",
									"v1",
									"payruns",
									"{id}",
									"request-authorisation",
								},
								"select": map[string]any{
									"$action": "request_authorisation",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/payruns/{id}/submit",
								"parts": []any{
									"api",
									"v1",
									"payruns",
									"{id}",
									"submit",
								},
								"select": map[string]any{
									"$action": "submit",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/payruns/{merchantID}",
								"parts": []any{
									"api",
									"v1",
									"payruns",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "only_archived",
											"orig": "only_archived",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "status",
											"orig": "status",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/payruns",
								"parts": []any{
									"api",
									"v1",
									"payruns",
								},
								"select": map[string]any{
									"exist": []any{
										"from_date",
										"merchant_id",
										"only_archived",
										"page_number",
										"page_size",
										"search",
										"sort",
										"status",
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/payruns/{id}",
								"parts": []any{
									"api",
									"v1",
									"payruns",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/payruns/{id}",
								"parts": []any{
									"api",
									"v1",
									"payruns",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/payruns/{id}/archive",
								"parts": []any{
									"api",
									"v1",
									"payruns",
									"{id}",
									"archive",
								},
								"select": map[string]any{
									"$action": "archive",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/payruns/{id}",
								"parts": []any{
									"api",
									"v1",
									"payruns",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/payruns/{id}/cancel",
								"parts": []any{
									"api",
									"v1",
									"payruns",
									"{id}",
									"cancel",
								},
								"select": map[string]any{
									"$action": "cancel",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/payruns/{id}/reject",
								"parts": []any{
									"api",
									"v1",
									"payruns",
									"{id}",
									"reject",
								},
								"select": map[string]any{
									"$action": "reject",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/payruns/{id}/unarchive",
								"parts": []any{
									"api",
									"v1",
									"payruns",
									"{id}",
									"unarchive",
								},
								"select": map[string]any{
									"$action": "unarchive",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"report": map[string]any{
				"fields": []any{},
				"name": "report",
				"op": map[string]any{
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/reports/{id}/initiate",
								"parts": []any{
									"api",
									"v1",
									"reports",
									"{id}",
									"initiate",
								},
								"select": map[string]any{
									"$action": "initiate",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"report_result": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "contentType",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "contents",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastCompletedAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reportName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reportType",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "statementNumber",
						"type": "`$INTEGER`",
					},
				},
				"name": "report_result",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "statement_number",
											"reqd": true,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "param",
											"name": "report_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/reports/{id}/result/{statementNumber}",
								"parts": []any{
									"api",
									"v1",
									"reports",
									"{report_id}",
									"result",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "report_id",
										"statementNumber": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"report_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"report",
						},
					},
				},
			},
			"role": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "failedRoles",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "roles",
						"type": "`$ARRAY`",
					},
				},
				"name": "role",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/merchants/{merchantID}/roles/batchcreate",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"roles",
									"batchcreate",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"$action": "batchcreate",
									"exist": []any{
										"merchant_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"merchant",
						},
					},
				},
			},
			"rule": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "account",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "accountID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "approveUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "approverID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "authenticationMethods",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisations",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisersCompletedCount",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "authorisersRequiredCount",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "canAuthorise",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "createdBy",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "endAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hasCurrentUserAuthorised",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isDisabled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastExecutedAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastRunAtTransactionDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nonce",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "onApprovedWebHookUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "onExecutionErrorWebHookUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "onExecutionSuccessWebHookUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "startAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sweepAction",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$OBJECT`",
							},
						},
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "timeZoneId",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "triggerCronExpression",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "triggerOnPayIn",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "userID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "webHookSecret",
						"type": "`$STRING`",
					},
				},
				"name": "rule",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/rules",
								"parts": []any{
									"api",
									"v1",
									"rules",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "archived_only",
											"orig": "archived_only",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "merchant_id",
											"orig": "merchant_id",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/rules",
								"parts": []any{
									"api",
									"v1",
									"rules",
								},
								"select": map[string]any{
									"exist": []any{
										"archived_only",
										"merchant_id",
										"page",
										"search",
										"size",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/rules/{id}",
								"parts": []any{
									"api",
									"v1",
									"rules",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/rules/{id}",
								"parts": []any{
									"api",
									"v1",
									"rules",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/rules/{id}",
								"parts": []any{
									"api",
									"v1",
									"rules",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/rules/{id}/disable",
								"parts": []any{
									"api",
									"v1",
									"rules",
									"{id}",
									"disable",
								},
								"select": map[string]any{
									"$action": "disable",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"rule_event": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "errorMessage",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isAuthoriseToEnable",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "message",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "rawResponse",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ruleEventType",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ruleID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "user",
						"req": true,
						"type": "`$OBJECT`",
					},
				},
				"name": "rule_event",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "event_type",
											"orig": "event_type",
											"type": "`$ARRAY`",
										},
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page",
											"orig": "page",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "size",
											"orig": "size",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/rules/{id}/events",
								"parts": []any{
									"api",
									"v1",
									"rules",
									"{id}",
									"events",
								},
								"select": map[string]any{
									"exist": []any{
										"event_type",
										"id",
										"page",
										"size",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"tag": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "colourHex",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"req": true,
						"type": "`$STRING`",
					},
				},
				"name": "tag",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/merchants/{merchantID}/tags",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"tags",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"merchant_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/tags",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"tags",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"merchant_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"merchant",
						},
					},
				},
			},
			"token": map[string]any{
				"fields": []any{},
				"name": "token",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/tokens/authorise/{id}",
								"parts": []any{
									"api",
									"v1",
									"tokens",
									"authorise",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/tokens/{id}",
								"parts": []any{
									"api",
									"v1",
									"tokens",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"transaction": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "accountID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountSequenceNumber",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "addressDetails",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "amount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "balance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "balanceMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "bookingDateTime",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "chargeDetails",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "content",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "counterparty",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "counterpartySummary",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currencyExchange",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "date",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "enrichment",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "fxAmount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fxCurrency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxRate",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "grossAmount",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isoBankTransactionCode",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "merchant",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pageNumber",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "pageSize",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "payeeDetails",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "payerDetails",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "paymentRequestCustomFields",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "paymentRequestID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payoutID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "proprietaryBankTransactionCode",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "rawReference",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reference",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ruleID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "statementReferences",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplementaryData",
						"type": "`$ANY`",
					},
					map[string]any{
						"name": "tags",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "theirReference",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "totalPages",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "totalSize",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "transactionAmount",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "transactionDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "transactionInformation",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "transactionMutability",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "valueDateTime",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "virtualIBAN",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "yourReference",
						"type": "`$STRING`",
					},
				},
				"name": "transaction",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/transactions/{id}/tags",
								"parts": []any{
									"api",
									"v1",
									"transactions",
									"{id}",
									"tags",
								},
								"select": map[string]any{
									"$action": "tag",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "credit_type",
											"orig": "credit_type",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "max_amount",
											"orig": "max_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "min_amount",
											"orig": "min_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/accounts/{accountID}/transactions",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"{account_id}",
									"transactions",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"credit_type",
										"from_date",
										"max_amount",
										"min_amount",
										"page_number",
										"page_size",
										"search",
										"sort",
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "limit",
											"orig": "limit",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "offset",
											"orig": "offset",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort_inserted_ascending",
											"orig": "sort_inserted_ascending",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "transaction_from",
											"orig": "transaction_from",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "transaction_to",
											"orig": "transaction_to",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/openbanking/transactions/{id}/{accountID}",
								"parts": []any{
									"api",
									"v1",
									"openbanking",
									"transactions",
									"{id}",
									"{account_id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
										"limit",
										"offset",
										"sort_inserted_ascending",
										"transaction_from",
										"transaction_to",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": false,
											"kind": "query",
											"name": "include_child_merchant",
											"orig": "include_child_merchant",
											"type": "`$BOOLEAN`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/transactions",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"transactions",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"from_date",
										"include_child_merchant",
										"merchant_id",
										"page_number",
										"page_size",
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "credit_type",
											"orig": "credit_type",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/transactions",
								"parts": []any{
									"api",
									"v1",
									"transactions",
								},
								"select": map[string]any{
									"exist": []any{
										"credit_type",
										"from_date",
										"page_number",
										"page_size",
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "credit_type",
											"orig": "credit_type",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "max_amount",
											"orig": "max_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "min_amount",
											"orig": "min_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/transactions/{accountID}/export",
								"parts": []any{
									"api",
									"v1",
									"transactions",
									"{account_id}",
									"export",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
									},
								},
								"select": map[string]any{
									"$action": "export",
									"exist": []any{
										"account_id",
										"credit_type",
										"from_date",
										"max_amount",
										"min_amount",
										"page_number",
										"page_size",
										"search",
										"sort",
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "credit_type",
											"orig": "credit_type",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "from_date",
											"orig": "from_date",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "max_amount",
											"orig": "max_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "min_amount",
											"orig": "min_amount",
											"type": "`$NUMBER`",
										},
										map[string]any{
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "to_date",
											"orig": "to_date",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/transactions/{accountID}",
								"parts": []any{
									"api",
									"v1",
									"transactions",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"credit_type",
										"from_date",
										"id",
										"max_amount",
										"min_amount",
										"page_number",
										"page_size",
										"search",
										"sort",
										"to_date",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": 0,
											"kind": "param",
											"name": "sequence_number",
											"orig": "sequence_number",
											"reqd": true,
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "param",
											"name": "transaction_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/transactions/{accountID}/from/{sequenceNumber}",
								"parts": []any{
									"api",
									"v1",
									"transactions",
									"{transaction_id}",
									"from",
									"{sequence_number}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "transaction_id",
										"sequenceNumber": "sequence_number",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"page_size",
										"sequence_number",
										"transaction_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/accounts/{accountID}/transactions/{id}",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"{account_id}",
									"transactions",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/transactions/detail/{id}",
								"parts": []any{
									"api",
									"v1",
									"transactions",
									"detail",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/transactions/{id}/proof",
								"parts": []any{
									"api",
									"v1",
									"transactions",
									"{id}",
									"proof",
								},
								"select": map[string]any{
									"$action": "proof",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "tag_id",
											"orig": "tag_id",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/transactions/{id}/tag",
								"parts": []any{
									"api",
									"v1",
									"transactions",
									"{id}",
									"tag",
								},
								"select": map[string]any{
									"$action": "tag",
									"exist": []any{
										"id",
										"tag_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"account",
						},
						[]any{
							"merchant",
						},
						[]any{
							"transaction",
						},
						[]any{
							"transaction",
							"from",
						},
					},
				},
			},
			"user": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "clientSessionTimeouts",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "emailAddress",
						"op": map[string]any{
							"update": map[string]any{
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "firstName",
						"op": map[string]any{
							"update": map[string]any{
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastName",
						"op": map[string]any{
							"update": map[string]any{
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "passkeyAdded",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "permissions",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "profile",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "rolesWithScope",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "twoFactorEnabled",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "userInviteID",
						"type": "`$STRING`",
					},
				},
				"name": "user",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/user/{merchantID}/userspaged",
								"parts": []any{
									"api",
									"v1",
									"user",
									"{merchant_id}",
									"userspaged",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"$action": "userspaged",
									"exist": []any{
										"merchant_id",
										"page_number",
										"page_size",
										"search",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/users",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"users",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"merchant_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/metadata/whoami",
								"parts": []any{
									"api",
									"v1",
									"metadata",
									"whoami",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/metadata/whoamitrustedapp",
								"parts": []any{
									"api",
									"v1",
									"metadata",
									"whoamitrustedapp",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/user",
								"parts": []any{
									"api",
									"v1",
									"user",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/user/{id}",
								"parts": []any{
									"api",
									"v1",
									"user",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"merchant",
						},
						[]any{
							"user",
						},
					},
				},
			},
			"user_invite": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "authorisationStatus",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "failedUserInvites",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "initialRoleID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inviteeEmailAddress",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inviteeFirstName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inviteeLastName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inviterEmailAddress",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inviterFirstName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inviterLastName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isAuthorised",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isInviteeRegistered",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastInvited",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "message",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "registrationUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sendInviteEmail",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "user",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "userID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "userInvites",
						"type": "`$ARRAY`",
					},
				},
				"name": "user_invite",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/userinvites/authorise/{id}",
								"parts": []any{
									"api",
									"v1",
									"userinvites",
									"authorise",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/userinvites",
								"parts": []any{
									"api",
									"v1",
									"userinvites",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/userinvites/batchcreate",
								"parts": []any{
									"api",
									"v1",
									"userinvites",
									"batchcreate",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": 1,
											"kind": "query",
											"name": "page_number",
											"orig": "page_number",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"example": 20,
											"kind": "query",
											"name": "page_size",
											"orig": "page_size",
											"type": "`$INTEGER`",
										},
										map[string]any{
											"kind": "query",
											"name": "search",
											"orig": "search",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "sort",
											"orig": "sort",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/userinvitespaged",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"userinvitespaged",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"merchant_id",
										"page_number",
										"page_size",
										"search",
										"sort",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/userinvites/{id}",
								"parts": []any{
									"api",
									"v1",
									"userinvites",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "userinvite_id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/userinvites/{id}/details",
								"parts": []any{
									"api",
									"v1",
									"userinvites",
									"{userinvite_id}",
									"details",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"id": "userinvite_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"userinvite_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/userinvites/{id}",
								"parts": []any{
									"api",
									"v1",
									"userinvites",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/userinvites/{id}",
								"parts": []any{
									"api",
									"v1",
									"userinvites",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"merchant",
						},
						[]any{
							"userinvite",
						},
					},
				},
			},
			"virtual": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "accountName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountSupplierName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "availableBalance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "availableBalanceMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "balance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "balanceMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "bankName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "consentID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "createdBy",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "createdByDisplayName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "defaultPaymentRail",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "displayName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "expiryDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "externalAccountIcon",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "identifier",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isConnectedAccount",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isDefault",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isTrustAccount",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isVirtual",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastTransaction",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "physicalAccountID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "rules",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "submittedPayoutsBalance",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "submittedPayoutsBalanceMinorUnits",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "summary",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierSepaInstantStatus",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedConnectionStatus",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedLastSyncedAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedSyncLastFailedAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedSyncLastFailureReason",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedSyncStatus",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroUnsynchronisedTransactionsCount",
						"type": "`$INTEGER`",
					},
				},
				"name": "virtual",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/accounts/{accountID}/virtual",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"{account_id}",
									"virtual",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "account_id",
											"orig": "account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "virtual_account_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/accounts/{accountID}/virtual/{virtualAccountID}",
								"parts": []any{
									"api",
									"v1",
									"accounts",
									"{account_id}",
									"virtual",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"accountID": "account_id",
										"virtualAccountID": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"account_id",
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"account",
						},
					},
				},
			},
			"webhook": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "destinationUrl",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "emailAddress",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "failedNotificationEmailAddress",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isActive",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "merchantID",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
							"update": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notificationMethod",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
							"update": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "resourceTypes",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "retry",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "secret",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "version",
						"type": "`$INTEGER`",
					},
				},
				"name": "webhook",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/api/v1/webhooks",
								"parts": []any{
									"api",
									"v1",
									"webhooks",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/webhooks",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"webhooks",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"merchant_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "param",
											"name": "merchant_id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/merchants/{merchantID}/webhooks/{id}",
								"parts": []any{
									"api",
									"v1",
									"merchants",
									"{merchant_id}",
									"webhooks",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "merchant_id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
										"merchant_id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "merchant_id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/api/v1/webhooks/{merchantID}",
								"parts": []any{
									"api",
									"v1",
									"webhooks",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"merchantID": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/api/v1/webhooks/{id}",
								"parts": []any{
									"api",
									"v1",
									"webhooks",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PUT",
								"orig": "/api/v1/webhooks/{id}",
								"parts": []any{
									"api",
									"v1",
									"webhooks",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"merchant",
						},
					},
				},
			},
		},
	}
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
