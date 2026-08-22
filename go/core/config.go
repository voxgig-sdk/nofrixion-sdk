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
			"slug": "nofrixion",
			"version": "0.0.1",
			"target": "go",
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
						"short": "The various balances for the account.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "accountID",
						"short": "ID of the account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountIdentifications",
						"short": "The canoncial identifiers for the account.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "accountName",
						"short": "Name for the account",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountNames",
						"short": "Optional account names set by the account holder.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "accountSupplierName",
						"short": "The payment account supplier name.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountType",
						"short": "The type of account e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "availableBalance",
						"short": "The current available balance of the account.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "availableBalanceMinorUnits",
						"short": "The available balance expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "balance",
						"short": "Balance of the account.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "balanceMinorUnits",
						"short": "Balance of the account expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "bankName",
						"short": "The bank name for external accounts",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "consentID",
						"short": "The ID of the consent used to connect the external account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "consolidatedAccountInformation",
						"short": "Summary information regarding account balances of the overall account provided by the bank.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "createdBy",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "createdByDisplayName",
						"short": "Either the name of the user, merchant token or api key that created the account",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"short": "Currency of the account in ISO 4217 format",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "defaultPaymentRail",
						"short": "Indicates the default payment rail for this account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"short": "Product name as defined by the financial institution for this account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "details",
						"short": "Supplementary specifications that might be provided by the Bank.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "displayName",
						"short": "Gets a unique display name for the payment account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "expiryDate",
						"short": "The date that the external account will expire",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "externalAccountIcon",
						"short": "The Icon for external accounts",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "format",
						"short": "File format to save the statement as.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fromDate",
						"short": "Minimum transaction date for the statement.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique id for the account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "identifier",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "inserted",
						"short": "Timestamp when the account was created.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"short": "Indicates whether the account is archived.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isConnectedAccount",
						"short": "Indicates if the payment account is an externally connected account.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isDefault",
						"short": "Is the default account",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isTrustAccount",
						"short": "Indicates if the payment account is a trust account.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isVirtual",
						"short": "True if the account is a virtual account.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastTransaction",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "lastUpdated",
						"short": "Timestamp when the account was last updated.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"short": "The ID of the merchant that owns the account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantName",
						"short": "The name of the merchant that owns the account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nickname",
						"short": "Nickname of the account that was provided by the account owner.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "physicalAccountID",
						"short": "For virtual accounts this is the ID of the physical account that the virtual account is linked to.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "roleIDs",
						"short": "Optional list of role IDs that will get access to the payment account when created.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "rules",
						"short": "The list of rules associated with this account.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "submittedPayoutsBalance",
						"short": "Total of the payouts that have been submitted for processing.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "submittedPayoutsBalanceMinorUnits",
						"short": "The balance of the submitted payouts expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "summary",
						"short": "Gets a summary of the payments account's most important properties.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierPhysicalAccountID",
						"short": "For internal use only.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierSepaInstantStatus",
						"short": "Indicates the status of the SEPA Instant payment rail for this account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "toDate",
						"short": "Maximum transaction date for the statement.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"short": "Specifies the type of account e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "usageType",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedConnectionStatus",
						"short": "States the status of the Xero bank feed connection, if applicable.",
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
						"short": "Indicates the number of unsynchronised transactions with Xero",
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
						"short": "This field is used when returning a batch payout record to a client.",
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
						"short": "A list of authentication types allowed to authorise the payout.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisations",
						"short": "A list of users who have successfully authorised the latest version of the beneficiary.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisersCompletedCount",
						"short": "The number of distinct authorisers that have authorised the beneficiary.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "authorisersRequiredCount",
						"short": "The number of authorisers required for this beneficiary.",
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
						"short": "True if the beneficiary can be authorised by the user who loaded it.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "canUpdate",
						"short": "True if the beneficiary can be updated by the user who loaded it.",
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
						"short": "Gets or Sets the currency.",
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
						"short": "True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary.",
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
						"short": "Gets or Sets the merchant id.",
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
						"short": "The descriptive name for the beneficiary.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nonce",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountIDs",
						"short": "ID of the accounts which are authorised to act as a source for the beneficiary.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "sourceAccounts",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "theirReference",
						"short": "The reference that will be used by default as TheirReference when creating payouts to this beneficiary if no TheirReference is specified for the payout.",
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
						"short": "The existing group members.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "groupName",
						"req": true,
						"short": "The descriptive name for the beneficiary group.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"short": "Timestamp indicating when the group was created.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastUpdated",
						"short": "Timestamp indicating when the group was last updated.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"req": true,
						"short": "Gets or Sets the merchant id.",
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
						"short": "Gets set to true if 3-D Secure payer authentication is required for a payment.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isSoftDecline",
						"short": "Gets set to true if the card processor flagged the transaction as having failed address or card security number verification.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "payerAuthenticationAccessToken",
						"short": "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payerAuthenticationMerchantData",
						"short": "If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the \"authenticationcallback\" method that gets called automatically after a successful payer authenticati…",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payerAuthenticationUrl",
						"short": "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payerAuthenticationWindowHeight",
						"short": "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "payerAuthenticationWindowWidth",
						"short": "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "paymentRequestCallbackUrl",
						"short": "The callback URL that was set when the payment request was created.",
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
						"short": "Checkout.com require a redirect for 3DS authentication.",
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
						"short": "The type of the tokenised card, e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerEmailAddress",
						"short": "When creating a tokenised card the payer's email address must be supplied.",
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
						"short": "The unique ID of the card token that has been stored for the customer.",
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
						"short": "Gets set to true if 3-D Secure payer authentication is required for a payment.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isSoftDecline",
						"short": "Gets set to true if the card processor flagged the transaction as having failed address or card security number verification.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "payerAuthenticationAccessToken",
						"short": "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payerAuthenticationMerchantData",
						"short": "If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the \"authenticationcallback\" method that gets called automatically after a successful payer authenticati…",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payerAuthenticationUrl",
						"short": "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payerAuthenticationWindowHeight",
						"short": "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "payerAuthenticationWindowWidth",
						"short": "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "paymentRequestCallbackUrl",
						"short": "The callback URL that was set when the payment request was created.",
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
						"short": "Checkout.com require a redirect for 3DS authentication.",
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
						"short": "The URL the authorising user needs to be redirected to in order to get the open banking consent token.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "callbackUrl",
						"short": "Optional callback URL that the end user performing the open banking authorisation will be redirected to on completion.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "consentID",
						"short": "The ID of the open banking consent.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "emailAddress",
						"short": "The email address that identifies the end user that will be authorising the open banking consent request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "expiryDate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "failureCallbackUrl",
						"short": "Optional callback URL for open banking consent authorisation failure.",
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
						"short": "The institution ID the open banking consent is being requested for.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isConnectedAccounts",
						"short": "Optional setting.",
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
						"short": "The ID of the merchant the consent token is being created to be used with.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "provider",
						"short": "Lists the supported card and PIS processors.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "successWebHookUrl",
						"short": "A web hook URL to send an HTTP request to when the open banking consent is successfuly authorised.",
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
						"short": "Dictionary of failed submissions, keyed by the index (1-based) in the original request.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "successfulSubmissions",
						"short": "List of successfully submitted direct debit payments.",
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
						"short": "The price at which the transaction will buy the source currency using the destination currency.",
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
						"short": "Account number of the customer's bank account in case of GBP account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "addressLine1",
						"req": true,
						"short": "First line of the customer's address.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "addressLine2",
						"short": "Second line of the customer's address.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "approvedAt",
						"short": "Date at which the supplier approved this mandate.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "city",
						"req": true,
						"short": "Customer's city.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "countryCode",
						"req": true,
						"short": "2-character country code of the customer's bank account.",
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
						"short": "Currency of this mandate.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerAccountNumber",
						"short": "Customer's account number in case of GBP account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerCity",
						"short": "Customer's city of residence.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerCountryCode",
						"short": "Customer's country of residence code.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerCountryName",
						"short": "Customer's country of residence.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerEmailAddress",
						"short": "Customer's email address.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerFirstName",
						"short": "Customer's first name.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerIban",
						"short": "Customer's IBAN in case of EUR account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerLastName",
						"short": "Customer's last name.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerSortCode",
						"short": "Customer's sort code in case of GBP account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "emailAddress",
						"req": true,
						"short": "Customer's email address.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "firstName",
						"req": true,
						"short": "Customer's first name.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "iban",
						"short": "IBAN of the customer's bank account in case of EUR account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Internal ID of the mandate.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"short": "The timestamp this mandate was created at.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isRecurring",
						"short": "Whether this mandate is single-use or recurring.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastName",
						"req": true,
						"short": "Customer's last name.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastUpdated",
						"short": "The timestamp this mandate was last updated at.",
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
						"short": "Internal ID of this mandate's merchant.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "postalCode",
						"req": true,
						"short": "Customer's postal code.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reference",
						"short": "Reference assigned to this mandate.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sortCode",
						"short": "Sort code of the customer's bank account in case of GBP account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"short": "General status of this mandate.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierBankAccountID",
						"short": "ID that the supplier assigned to this mandate's bank account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierCustomerID",
						"short": "ID that the supplier assigned to this mandate's customer.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierMandateID",
						"short": "ID that the supplier assigned to this mandate.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierName",
						"short": "Name of the supplier used to create this mandate.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierStatus",
						"short": "Last status that the supplier reported for this mandate.",
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
						"short": "The list of currencies that the merchant has accounts for.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "canHaveTrustAccounts",
						"short": "Trust accounts are a special type of account that allow the account name to be trusted for use in statements and verification of payee checks.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardPaymentProcessor",
						"short": "Name of the card payment processor.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "companyID",
						"short": "The Company ID recorded in the Compliance system.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "displayQrOnHostedPay",
						"short": "Indicates if a QR Code containing the payment link should be displayed on the hosted payment page.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "hostedPayVersion",
						"short": "The version of the hosted payment page to use with the merchant.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique ID for the merchant.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"short": "Timestamp the merchant was added to MoneyMoov.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isBlocked",
						"short": "The merchant is blocked from making payments (payouts).",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isExited",
						"short": "The merchant has formally terminated their relationship and is no longer a customer.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isSuspended",
						"short": "The merchant has temporarily suspended their own account.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "jurisdiction",
						"short": "The jurisdiction the merchant entity is incorporated or established in.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "logoUrlPng",
						"short": "The CDN URL of the merchant's logo in PNG format.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "logoUrlSvg",
						"short": "The CDN URL of the merchant's logo in SVG format.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantCategoryCode",
						"short": "The industry code that represents the merchant's primary trading activity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"short": "The registered business name of the merchant.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notes",
						"short": "The notes field is an optional free text field that can be used to store any additional information about the merchant.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "parentMerchant",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "paymentAccountLimit",
						"short": "The maximum number of payment accounts that can be created for the Merchant.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "paymentAccounts",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "reason",
						"short": "The reason for the suspension.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shortName",
						"short": "A URL friendly shortish name for the merchant.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supportedPaymentMethodsList",
						"short": "The payment methods that are configured and supported for this merchant.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "suspensionReason",
						"short": "The reason for the suspension, provided by the merchant.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tags",
						"short": "An optional list of descriptive tags that can be used on merchant entities such as payment requests.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "timeZoneId",
						"short": "The IANA (Internet Assigned Numbers Authority) time zone identifier of the merchant.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tradingName",
						"short": "An optional trading name.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "webHookLimit",
						"short": "The maximum number of web hooks that can be created for the Merchant.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "yourRoleName",
						"short": "The name of the role for the identity that loaded the merchant record.",
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
						"short": "Date at which the supplier approved this mandate.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"short": "Currency of this mandate.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerAccountNumber",
						"short": "Customer's account number in case of GBP account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerCity",
						"short": "Customer's city of residence.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerCountryCode",
						"short": "Customer's country of residence code.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerCountryName",
						"short": "Customer's country of residence.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerEmailAddress",
						"short": "Customer's email address.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerFirstName",
						"short": "Customer's first name.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerIban",
						"short": "Customer's IBAN in case of EUR account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerLastName",
						"short": "Customer's last name.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerSortCode",
						"short": "Customer's sort code in case of GBP account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Internal ID of the mandate.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"short": "The timestamp this mandate was created at.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isRecurring",
						"short": "Whether this mandate is single-use or recurring.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastUpdated",
						"short": "The timestamp this mandate was last updated at.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"short": "Internal ID of this mandate's merchant.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reference",
						"short": "Reference assigned to this mandate.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"short": "General status of this mandate.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierBankAccountID",
						"short": "ID that the supplier assigned to this mandate's bank account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierCustomerID",
						"short": "ID that the supplier assigned to this mandate's customer.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierMandateID",
						"short": "ID that the supplier assigned to this mandate.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierName",
						"short": "Name of the supplier used to create this mandate.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierStatus",
						"short": "Last status that the supplier reported for this mandate.",
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
						"short": "The list of country codes representing the banks the country supports.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "bankID",
						"short": "ID of the bank to be configured for the merchant.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "bankName",
						"short": "Name of the Bank/Institution.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "businessInstitutionID",
						"short": "ID that the processor uses to identify the bank (business accounts).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"short": "Currency supported by the bank.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "logo",
						"short": "URL of the bank's logo.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "message",
						"short": "Message relating to specific bank.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "messageImageUrl",
						"short": "Optional image URL to be displayed with the message.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "order",
						"short": "Order in which this setting will appear in the UI.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "personalInstitutionID",
						"short": "ID that the processor uses to identify the bank (personal accounts).",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "processor",
						"short": "Name of the bank payment processor.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "warningHeading",
						"short": "The heading for a warning message related to the bank institution to be displayed to the user.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "warningMessage",
						"short": "The warning message related to the bank institution to be displayed to the user.",
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
						"short": "A list of custom fields that can be included in the payment request template.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "defaultFields",
						"short": "A list of default fields that are included in the payment request template.",
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
						"short": "A list of authentication types allowed to authorise the merchant token.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisations",
						"short": "A list of users who have successfully authorised the latest version of the beneficiary.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisersCompletedCount",
						"short": "The number of distinct authorisers that have authorised the merchant token.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "authorisersRequiredCount",
						"short": "The number of authorisers required for this merchant token.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "canAuthorise",
						"short": "True if the merchant token can be authorised by the user who loaded it.",
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
						"short": "Token description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "expiresAt",
						"short": "Optional.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hasCurrentUserAuthorised",
						"short": "True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "hmacAlgorithm",
						"short": "Optional shared secret algorithm to use for HMAC authentication.",
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
						"short": "Optional.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"short": "Indicates whether the merchant token is archived.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isEnabled",
						"short": "If set to false the merchant token will not be accepted to authorise a request.",
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
						"short": "The merchant id to add to the token",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nonce",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "permissionTypes",
						"short": "The permissions that the merchant token supports.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "requestSignatureVersion",
						"short": "Represent the version of the overall merchant token.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "sharedSecretAlgorithm",
						"short": "Optional shared secret algorithm to use for HMAC authentication.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sharedSecretBase64",
						"short": "The base 64 encoded shared secret that is used for request authentication with an HMAC.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "token",
						"short": "The JWT merchant token.",
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
						"short": "The name of the account to verify",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountNumber",
						"short": "The account number of the account to verify (for CoP checks)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "iban",
						"req": true,
						"short": "The IBAN of the account to verify (for VoP checks)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payeeVerifiedAccountName",
						"short": "The verified account name of the payee, if available (in case of a close match)",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"short": "The result of the payee verification",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "secondaryIdentification",
						"short": "Optional secondary identifier for the account to verify.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sortCode",
						"short": "The sort code of the account to verify (for CoP checks)",
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
						"short": "The amount of money to request.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountPending",
						"short": "Total amount that has been authorised but not settled for this payment request.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountReceived",
						"short": "Total amount received for this payment request.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountRefunded",
						"short": "Total amount refunded for this payment request.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "autoSendReceipt",
						"short": "If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "baseOriginUrl",
						"short": "For card payments the origin of the payment page needs to be set in advance.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "callbackUrl",
						"short": "Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardAuthorizeOnly",
						"short": "For card payments the default behaviour is to authorise and capture the payment at the same time.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardCreateToken",
						"short": "For card payments a payment attempt can be used to create a reusable token for subsequent payments.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardCreateTokenMode",
						"short": "This specifies whether user consent will be taken before tokenising card or not.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardIgnoreCVN",
						"short": "If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardNoPayerAuthentication",
						"short": "If set to true for card payments no attempt will be made to use payer authentication (3-D Secure and equivalent).",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardProcessorMerchantID",
						"short": "Optional field that if specified indicates the processor merchant ID that should be used to process any card payments.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardStripePaymentIntentID",
						"short": "If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardStripePaymentIntentSecret",
						"short": "If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardTransmitRawDetails",
						"short": "If set to true for card payments the sensitive card number and card verification number will be transmitted directly rather than being tokenised.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "createdByUser",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "currency",
						"short": "The currency of the request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customFields",
						"short": "A list of custom fields attached to the payment request.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "customerEmailAddress",
						"short": "Optional email address for the customer.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerID",
						"short": "An optional customer identifier for the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"short": "An optional description for the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "destinationAccount",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "directDebitPayment",
						"short": "Contains information about a Direct Debit payment attempt for a payment request.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "dueDate",
						"short": "The due date for the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "events",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "failureCallbackUrl",
						"short": "Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fieldDisplaySettings",
						"short": "A list of field display settings that control which fields are displayed to the payer.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "formattedAmount",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hostedPayCheckoutUrl",
						"short": "This is a convenience link generated for payment requests whose merchants are using hosted payment pages.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ignoreAddressVerification",
						"short": "If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "inserted",
						"short": "The timestamp the payment request was created at.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "insertedSortable",
						"short": "The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se…",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"short": "Indicates whether the payment request is archived.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "jwk",
						"short": "The jwk containing the public key used to verify the signature of the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastUpdated",
						"short": "The timestamp the payment request was last updated at.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lightningInvoice",
						"short": "Bitcoin Lightning invoice for the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lightningInvoiceExpiresAt",
						"short": "Date and time of expiration of the lightning invoice.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantDirectDebitMandateID",
						"short": "Optional ID of the direct debit mandate associated with this payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"short": "The ID of the merchant to create the payment request for.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantTokenDescription",
						"short": "Description of the merchant token in case the Payment request was created using a merchant token.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notificationEmailAddresses",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notificationRoleIDs",
						"short": "A list of roles whose members will receive notifications about this payment request.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "orderID",
						"short": "An optional order ID for the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "partialPaymentMethod",
						"short": "The approach to use, or not, for accepting partial payments.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "partialPaymentSteps",
						"short": "An optional comma separated list of partial payment amounts.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentAttempts",
						"short": "The payment attempts made against this payment request.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "paymentMethods",
						"short": "The payment methods that the payment request supports.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "paymentProcessor",
						"short": "If the card payment option is enabled this field indicates which card processor the merchant is set up to use.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payrunID",
						"short": "The ID of a payrun that needs an account top up.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispAccountID",
						"short": "The payment account ID to use to receive payment initiation payments.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "priorityBankID",
						"short": "The ID of the bank that is set as the priority bank for display on pay element.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sandboxSettleDelayInSeconds",
						"short": "Sandbox only.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "shippingAddress",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "shippingAddressCity",
						"short": "Optionally the city of the customer's shipping address.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingAddressCountryCode",
						"short": "Optionally the country code of the customer's shipping address.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingAddressCounty",
						"short": "Optionally the state or county of the customer's shipping address.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingAddressLine1",
						"short": "Optionally the first line of the customer's shipping address.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingAddressLine2",
						"short": "Optionally the second line of the customer's shipping address.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingAddressPostCode",
						"short": "Optionally the post code of the customer's shipping address.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingEmail",
						"short": "Optionally the shipping email address for the customer.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingFirstName",
						"short": "Optionally the first name of the customer's shipping address.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingLastName",
						"short": "Optionally the last name of the customer's shipping address.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shippingPhone",
						"short": "Optionally the shipping phone number for the customer.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"short": "The current status of the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "successWebHookUrl",
						"short": "If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tagIds",
						"short": "An optional list of tag ids to add to the payment request",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "tags",
						"short": "An optional list of descriptive tags attached to the payment request.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "title",
						"short": "A generic field to contain any additional data that the merchant wishes to store against the payment request.",
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
						"short": "If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page.",
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
						"short": "Name for the account",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountSupplierName",
						"short": "The payment account supplier name.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "availableBalance",
						"short": "The current available balance of the account.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "availableBalanceMinorUnits",
						"short": "The available balance expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "balance",
						"short": "Balance of the account.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "balanceMinorUnits",
						"short": "Balance of the account expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "bankName",
						"short": "The bank name for external accounts",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "consentID",
						"short": "The ID of the consent used to connect the external account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "createdBy",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "createdByDisplayName",
						"short": "Either the name of the user, merchant token or api key that created the account",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"short": "Currency of the account in ISO 4217 format",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "defaultPaymentRail",
						"short": "Indicates the default payment rail for this account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "displayName",
						"short": "Gets a unique display name for the payment account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "expiryDate",
						"short": "The date that the external account will expire",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "externalAccountIcon",
						"short": "The Icon for external accounts",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique id for the account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "identifier",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "inserted",
						"short": "Timestamp when the account was created.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"short": "Indicates whether the account is archived.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isConnectedAccount",
						"short": "Indicates if the payment account is an externally connected account.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isDefault",
						"short": "Is the default account",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isTrustAccount",
						"short": "Indicates if the payment account is a trust account.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isVirtual",
						"short": "True if the account is a virtual account.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastTransaction",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "lastUpdated",
						"short": "Timestamp when the account was last updated.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"short": "The ID of the merchant that owns the account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantName",
						"short": "The name of the merchant that owns the account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "physicalAccountID",
						"short": "For virtual accounts this is the ID of the physical account that the virtual account is linked to.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "rules",
						"short": "The list of rules associated with this account.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "submittedPayoutsBalance",
						"short": "Total of the payouts that have been submitted for processing.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "submittedPayoutsBalanceMinorUnits",
						"short": "The balance of the submitted payouts expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "summary",
						"short": "Gets a summary of the payments account's most important properties.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierSepaInstantStatus",
						"short": "Indicates the status of the SEPA Instant payment rail for this account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedConnectionStatus",
						"short": "States the status of the Xero bank feed connection, if applicable.",
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
						"short": "Indicates the number of unsynchronised transactions with Xero",
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
						"short": "Name for the account",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "availableBalance",
						"short": "The current available balance of the account.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "balance",
						"short": "Balance of the account.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "balanceMinorUnits",
						"short": "Balance of the account expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "currency",
						"short": "Currency of the account in ISO 4217 format",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique id for the account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "identifier",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "isArchived",
						"short": "Is the account archived",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isConnectedAccount",
						"short": "Indicates if the payment account is an externally connected account.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "merchantID",
						"short": "The ID of the merchant that owns the account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "submittedPayoutsBalance",
						"short": "Total of the payouts that have been submitted for processing.",
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
						"short": "The unique identifier of the payment initiation request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRequestCallbackUrl",
						"short": "The callback URL that was set when the payment request was created.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRequestID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "redirectUrl",
						"short": "A redirect URL for the user to authorise the payment initiation request at the ASPSP",
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
						"short": "The amount of money to request.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountPending",
						"short": "Total amount that has been authorised but not settled for this payment request.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountReceived",
						"short": "Total amount received for this payment request.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountRefunded",
						"short": "Total amount refunded for this payment request.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "autoSendReceipt",
						"short": "If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "baseOriginUrl",
						"short": "For card payments the origin of the payment page needs to be set in advance.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "callbackUrl",
						"short": "Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardAuthorizeOnly",
						"short": "For card payments the default behaviour is to authorise and capture the payment at the same time.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardCreateToken",
						"short": "For card payments a payment attempt can be used to create a reusable token for subsequent payments.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardCreateTokenMode",
						"short": "This specifies whether user consent will be taken before tokenising card or not.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardIgnoreCVN",
						"short": "If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "cardProcessorMerchantID",
						"short": "Optional field that if specified indicates the processor merchant ID that should be used to process any card payments.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardStripePaymentIntentID",
						"short": "If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardStripePaymentIntentSecret",
						"short": "If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "createdByUser",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "currency",
						"short": "The currency of the request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customFields",
						"short": "A list of custom fields attached to the payment request.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "customerEmailAddress",
						"short": "Optional email address for the customer.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerID",
						"short": "An optional customer identifier for the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerName",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"short": "An optional description for the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "destinationAccount",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "directDebitPayment",
						"short": "Contains information about a Direct Debit payment attempt for a payment request.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "doSimulateSettlementFailure",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "dueDate",
						"short": "The due date for the payment request.",
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
						"short": "Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fieldDisplaySettings",
						"short": "A list of field display settings that control which fields are displayed to the payer.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "formattedAmount",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hostedPayCheckoutUrl",
						"short": "This is a convenience link generated for payment requests whose merchants are using hosted payment pages.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ignoreAddressVerification",
						"short": "If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "inserted",
						"short": "The timestamp the payment request was created at.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "insertedSortable",
						"short": "The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se…",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "institution",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"short": "Indicates whether the payment request is archived.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "jwk",
						"short": "The jwk containing the public key used to verify the signature of the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastUpdated",
						"short": "The timestamp the payment request was last updated at.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lightningInvoice",
						"short": "Bitcoin Lightning invoice for the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lightningInvoiceExpiresAt",
						"short": "Date and time of expiration of the lightning invoice.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantDirectDebitMandateID",
						"short": "Optional ID of the direct debit mandate associated with this payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantTokenDescription",
						"short": "Description of the merchant token in case the Payment request was created using a merchant token.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notificationEmailAddresses",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "notificationRoleIDs",
						"short": "A list of roles whose members will receive notifications about this payment request.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "orderID",
						"short": "An optional order ID for the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "partialPaymentMethod",
						"short": "The approach to use, or not, for accepting partial payments.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "partialPaymentSteps",
						"short": "An optional comma separated list of partial payment amounts.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentAttempts",
						"short": "The payment attempts made against this payment request.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "paymentInitiationID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentMethods",
						"short": "The payment methods that the payment request supports.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "paymentProcessor",
						"short": "If the card payment option is enabled this field indicates which card processor the merchant is set up to use.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRequests",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "payrunID",
						"short": "The ID of a payrun that needs an account top up.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispAccountID",
						"short": "The payment account ID to use to receive payment initiation payments.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "priorityBankID",
						"short": "The ID of the bank that is set as the priority bank for display on pay element.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "result",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sandboxSettleDelayInSeconds",
						"short": "Sandbox only.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "shippingAddress",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "status",
						"short": "The current status of the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "successWebHookUrl",
						"short": "If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tags",
						"short": "An optional list of descriptive tags attached to the payment request.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "title",
						"short": "A generic field to contain any additional data that the merchant wishes to store against the payment request.",
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
						"short": "If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page.",
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
						"short": "Transaction ID received in Apple pay token.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardAuthorizationResponseID",
						"short": "For a successful card authorization this field will hold the response ID.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardExpiryMonth",
						"short": "For card payment events this field holds the payer's card expiry month.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "cardExpiryYear",
						"short": "For card payment events this field holds the payer's card expiry year.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "cardIssuer",
						"short": "For card payment events this field holds the payer's card issuer.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardIssuerCountry",
						"short": "For card payment events this field holds the payer's card issuer country of origin.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardLastFourDigits",
						"short": "For card payment events this field holds the payer's card last four digits.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardRequestID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardScheme",
						"short": "For card payment events this field holds the scheme of the payer's card, e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "cardTokenCustomerID",
						"short": "If the option to create a reusable token for card payments was set this field contains the token the merchant can store to use for repeat payments.",
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
						"short": "Payment ID issued by the Direct Debit supplier.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "directDebitPaymentReference",
						"short": "Reference string issued by the Direct Debit supplier.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "drirectDebitMandateID",
						"short": "The ID of the mandate that was used wehn requesting payment.",
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
						"short": "For Bitcoin Lightning payments this field holds the invoice presented to the payer.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lightningRHash",
						"short": "For Bitcoin Lightning payments the hash of the invoice presented to the payer.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "originUrl",
						"short": "Optional field that can be set by payment methods, such as pay by bank, that may want to redirect back to the URL that initiated the attempt in the case of a failure condition.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentMethodType",
						"short": "The type of payment method the event relates to, e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentProcessorName",
						"short": "If the event was for a card payment this is the name of the card processor, e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRequestID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispBankStatus",
						"short": "For payment initiation attempts some providers (e.g.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispPaymentInitiationID",
						"short": "For a payment initiation this is the ID returned by the service provider initiating the payment for us.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispPaymentInstitutionName",
						"short": "For a payment initiation this is the name of the financial institution that is used to initiate and authorise the payment.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispPaymentServiceProviderID",
						"short": "For a payment initiation this is the service provider ID selected by the payer, typically the ID for the bank or similar financial institution.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispRedirectUrl",
						"short": "For a payment initiation this is the redirect URL returned by the service provider initiating the payment for us.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reconciledTransactionID",
						"short": "For settlement events (only relevant for non-card payments) this is the payin transaction that the payment request event was reconciled with.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "refundPayoutID",
						"short": "ID of the Payout that was created for refund.",
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
						"short": "The amount of money to request.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountPending",
						"short": "The amount of money that was authorised but has not arrived in the account yet.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountReceived",
						"short": "The amount of money that has been received for this payment request.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountRefunded",
						"short": "The amount of money that has been refunded for this payment request.",
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
						"short": "The country code associated with the payment.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"short": "The currency of the request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customFieldsToDisplay",
						"short": "Custom fields to display to the customer.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "description",
						"short": "An optional description for the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "dueDate",
						"short": "The due date of the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fieldDisplaySettings",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "googlePayMerchantID",
						"short": "Merchant ID from Google Pay",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "jwk",
						"short": "The jwk containing the public key",
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
						"short": "The payment attempts for this payment request.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "paymentMethodsList",
						"short": "The payment methods that the payment request supports.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "paymentProcessor",
						"short": "The card processor",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentProcessorKey",
						"short": "The card processors public key",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pispError",
						"short": "This is the error returned from the bank which is recorded in payment request events.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "priorityBankID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"short": "The status of the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "stripeAccountID",
						"short": "Account ID of connected customers in Stripe",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "title",
						"short": "The title of the payment request.",
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
						"short": "The authorised payment amount.",
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
						"short": "The authorised payment currency.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customerID",
						"short": "The customer id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRequestID",
						"short": "The ID of the payment request the result is for.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payments",
						"short": "The list of payment attempts that have been received for the payment request.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "pispAuthorizations",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "requestedAmount",
						"short": "The full original payment amount requested.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "result",
						"short": "The result of the payment attempt.",
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
						"short": "Gets or Sets Account Id of sending account",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "allowIncomplete",
						"short": "If set to true the payout will get created even if the business validation rules fail.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "amount",
						"short": "Gets or Sets payout amount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountMinorUnits",
						"short": "The payout amount expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "approvePayoutUrl",
						"short": "This field is used when returning an payout record to a client.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "approverID",
						"short": "Gets the User ID of person that approved the payout.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "authenticationMethods",
						"short": "A list of authentication types allowed to authorise the payout.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisations",
						"short": "A list of the users who have successfully authorised the latest version of the payout and when.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisersCompletedCount",
						"short": "The number of distinct authorisers that have authorised the payout.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "authorisersRequiredCount",
						"short": "The number of authorisers required for this payout.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "batchPayoutID",
						"short": "The ID of the batch the payout is associated with.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "beneficiary",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "beneficiaryID",
						"short": "Optional.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "canAuthorise",
						"short": "True if the payout can be authorised by the user who loaded it.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "canProcess",
						"short": "If set to true indicates the payout has been flagged as safe to process after transaction monitoring.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "canUpdate",
						"short": "True if the payout can be updated by the user who loaded it.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "chargeBearer",
						"short": "Optional field to set who should pay any fees for the payout.",
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
						"short": "Gets or Sets Currency of payout request",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currentUserID",
						"short": "The ID of the user that requested access to the PayOut record.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"short": "Gets or Sets description of payout request",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "destination",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "documents",
						"short": "Documents associated with the payout.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "events",
						"short": "The activity associated with the payout.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "failedPayouts",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "formattedAmount",
						"short": "Currency and formatted amount string.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "formattedFxDestinationAmount",
						"short": "FX destination currency and amount formatted string.",
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
						"short": "The available balance of the account the payout is being made from.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxDestinationAmount",
						"short": "If specified this will be the amount sent to the payee.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fxDestinationAmountMinorUnits",
						"short": "The payout FxDestinationAmount expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "fxDestinationCurrency",
						"short": "For an FX payout this is the currency to send to the beneficiary.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxQuoteExpiresAt",
						"short": "If an FX held rate quote ID is being used this is the time the quote expires.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxQuoteID",
						"short": "Optional.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxRate",
						"short": "For an FX payout this is the exchange rate to use for the payout.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fxUseDestinationAmount",
						"short": "For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "hasCurrentUserAuthorised",
						"short": "True if the payout was loaded for a user and that user has already authorised the latest version of the payout.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"short": "The ID for the payout.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "invoiceID",
						"short": "Optional field to associate the payout with the invoice from an external application such as Xero.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"short": "Indicates whether the payout is archived.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isFailed",
						"short": "Set to true if a submitted payout subsequently fails.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isSettled",
						"short": "Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isSubmitted",
						"short": "Indicates whether the payout has been submitted for processing.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"short": "The ID of the merchant that owns the account.",
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
						"short": "The usptream payment processor for the payout.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRail",
						"short": "Optional field to indicate the payment rail to use for the payout.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payouts",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "payrunID",
						"short": "The ID of the payrun that this payout is associated with.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payrunName",
						"short": "The name of the payrun that this payout is associated with.",
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
						"short": "The date the payout should be submitted.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "scheduled",
						"short": "Should this payout be scheduled for a future date?",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "sourceAccountAvailableBalance",
						"short": "The available balance of the account the payout is being made from.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "sourceAccountAvailableBalanceMinorUnits",
						"short": "The available balance of the source account expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "sourceAccountBic",
						"short": "The BIC of the account the payout is being made from.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountCurrency",
						"short": "The currency of the source account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountIban",
						"short": "The IBAN of the account the payout is being made from.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountIdentifier",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sourceAccountName",
						"short": "The name of the account the payout is being made from.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountNumber",
						"short": "The account number of the account the payout is being made from.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountSortcode",
						"short": "The sort code of the account the payout is being made from.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"short": "Gets or Sets the status of payout request",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tagIds",
						"short": "An optional list of tag ids to add to the payout.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "tags",
						"short": "An optional list of descriptive tags attached to the payout.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "theirReference",
						"short": "Gets or Sets destination reference ID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "topupPayrunID",
						"short": "The ID of a payrun that needs an account top up.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "transactedAmount",
						"short": "The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "transactedFxAmount",
						"short": "The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "transactedFxRate",
						"short": "The actual FX rate applied during settlement, as recorded on the associated transaction.",
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
						"short": "Gets or Sets payout type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "userID",
						"short": "Gets or Sets User ID of who created the payout request",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "yourReference",
						"short": "Gets or Sets your reference ID",
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
						"short": "Gets or Sets Account Id of sending account",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "amount",
						"short": "Gets or Sets payout amount",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountMinorUnits",
						"short": "The payout amount expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "approvePayoutUrl",
						"short": "This field is used when returning an payout record to a client.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "approverID",
						"short": "Gets the User ID of person that approved the payout.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "authenticationMethods",
						"short": "A list of authentication types allowed to authorise the payout.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisations",
						"short": "A list of the users who have successfully authorised the latest version of the payout and when.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisersCompletedCount",
						"short": "The number of distinct authorisers that have authorised the payout.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "authorisersRequiredCount",
						"short": "The number of authorisers required for this payout.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "batchPayoutID",
						"short": "The ID of the batch the payout is associated with.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "beneficiary",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "canAuthorise",
						"short": "True if the payout can be authorised by the user who loaded it.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "canProcess",
						"short": "If set to true indicates the payout has been flagged as safe to process after transaction monitoring.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "canUpdate",
						"short": "True if the payout can be updated by the user who loaded it.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "chargeBearer",
						"short": "Optional field to set who should pay any fees for the payout.",
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
						"short": "Gets or Sets Currency of payout request",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currentUserID",
						"short": "The ID of the user that requested access to the PayOut record.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"short": "Gets or Sets description of payout request",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "destination",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "documents",
						"short": "Documents associated with the payout.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "events",
						"short": "The activity associated with the payout.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "formattedAmount",
						"short": "Currency and formatted amount string.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "formattedFxDestinationAmount",
						"short": "FX destination currency and amount formatted string.",
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
						"short": "The available balance of the account the payout is being made from.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxDestinationAmount",
						"short": "If specified this will be the amount sent to the payee.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fxDestinationAmountMinorUnits",
						"short": "The payout FxDestinationAmount expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "fxDestinationCurrency",
						"short": "For an FX payout this is the currency to send to the beneficiary.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxQuoteExpiresAt",
						"short": "If an FX held rate quote ID is being used this is the time the quote expires.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxQuoteID",
						"short": "Optional.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxRate",
						"short": "For an FX payout this is the exchange rate to use for the payout.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fxUseDestinationAmount",
						"short": "For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "hasCurrentUserAuthorised",
						"short": "True if the payout was loaded for a user and that user has already authorised the latest version of the payout.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "id",
						"short": "The ID for the payout.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "invoiceID",
						"short": "Optional field to associate the payout with the invoice from an external application such as Xero.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"short": "Indicates whether the payout is archived.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isFailed",
						"short": "Set to true if a submitted payout subsequently fails.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isSettled",
						"short": "Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isSubmitted",
						"short": "Indicates whether the payout has been submitted for processing.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"short": "The ID of the merchant that owns the account.",
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
						"short": "The usptream payment processor for the payout.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "paymentRail",
						"short": "Optional field to indicate the payment rail to use for the payout.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payrunID",
						"short": "The ID of the payrun that this payout is associated with.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payrunName",
						"short": "The name of the payrun that this payout is associated with.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "rule",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "scheduleDate",
						"short": "The date the payout should be submitted.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "scheduled",
						"short": "Should this payout be scheduled for a future date?",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "sourceAccountAvailableBalance",
						"short": "The available balance of the account the payout is being made from.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "sourceAccountAvailableBalanceMinorUnits",
						"short": "The available balance of the source account expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "sourceAccountBic",
						"short": "The BIC of the account the payout is being made from.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountCurrency",
						"short": "The currency of the source account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountIban",
						"short": "The IBAN of the account the payout is being made from.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountIdentifier",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "sourceAccountName",
						"short": "The name of the account the payout is being made from.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountNumber",
						"short": "The account number of the account the payout is being made from.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "sourceAccountSortcode",
						"short": "The sort code of the account the payout is being made from.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"short": "Gets or Sets the status of payout request",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "tags",
						"short": "An optional list of descriptive tags attached to the payout.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "theirReference",
						"short": "Gets or Sets destination reference ID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "topupPayrunID",
						"short": "The ID of a payrun that needs an account top up.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "transactedAmount",
						"short": "The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "transactedFxAmount",
						"short": "The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "transactedFxRate",
						"short": "The actual FX rate applied during settlement, as recorded on the associated transaction.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "type",
						"short": "Gets or Sets payout type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "userID",
						"short": "Gets or Sets User ID of who created the payout request",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "yourReference",
						"short": "Gets or Sets your reference ID",
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
						"short": "A list of the users who have successfully authorised the latest version of the payrun and when.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisersCompletedCount",
						"short": "The number of distinct authorisers that have authorised the payrun.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "authorisersRequiredCount",
						"short": "The number of authorisers required for this payrun.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "batchPayoutID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "canAuthorise",
						"short": "True if the payrun can be authorised by the user who loaded it.",
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
						"short": "True if the payrun was loaded for a user and that user has already authorised the latest version of the payrun.",
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
						"short": "The ID of the account the rule will apply to.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "approveUrl",
						"short": "If set this property holds the URL an approver needs to visit in order to complete a strong authentication check in order to approve the rule.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "approverID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "authenticationMethods",
						"short": "A list of authentication types allowed to authorise the payout.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisations",
						"short": "A list of the users who have successfully authorised the latest version of the rule and when.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "authorisersCompletedCount",
						"short": "The number of distinct authorisers that have authorised the rule.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "authorisersRequiredCount",
						"short": "The number of authorisers required for this rule.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "canAuthorise",
						"short": "True if the rule can be authorised by the user who loaded it.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "createdBy",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "description",
						"short": "Arbitrary description for the rule.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "endAt",
						"short": "Optional end time for rule executions.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hasCurrentUserAuthorised",
						"short": "True if the current user has authorised.",
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
						"short": "If set to true the rule will be disabled from executing.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastExecutedAt",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastRunAtTransactionDate",
						"short": "The most recent transaction date when the rule was last run.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "lastUpdated",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"short": "The ID of the merchant that owns the account.",
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
						"short": "A name to succinctly describe the rule.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nonce",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "onApprovedWebHookUrl",
						"short": "Optional URL to receive an HTTP request with the rule details when the rule status changes to approved.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "onExecutionErrorWebHookUrl",
						"short": "Optional URL to receive an HTTP request when a rule execution attempt fails.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "onExecutionSuccessWebHookUrl",
						"short": "Optional URL to receive an HTTP request when a rule execution attempt succeeds.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "startAt",
						"short": "Optional start time for rule executions.",
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
						"short": "If the rule should be executed on a recurring schedule this is the timezone that the CRON expression should be evaluated in.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "triggerCronExpression",
						"short": "If the rule should be executed on a recurring schedule this is the expression that sets the schedule.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "triggerOnPayIn",
						"short": "Set to true if the rule execution should be triggered when the account receives a pay in (credit).",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "userID",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "webHookSecret",
						"short": "If set this secret will be used to sign Web Hook requests.",
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
						"short": "The ID of the account the transaction belongs to.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountName",
						"short": "The name of the account the transaction belongs to.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountSequenceNumber",
						"short": "The sequence number of transaction on a per account basis.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "addressDetails",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "amount",
						"short": "Amount of the transaction.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "amountMinorUnits",
						"short": "Amount of the transaction expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "balance",
						"short": "Balance left on the account after the transaction.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "balanceMinorUnits",
						"short": "Balance on the account expressed in the currency’s minor units (e.g.",
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
						"short": "For pay in (credit) transactions this will contain a descriptive string with the most important fields about the counterparty.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"short": "Currency of transaction.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currencyExchange",
						"short": "Provides details on the currency exchange.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "date",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "description",
						"short": "Description of the transaction.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "enrichment",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "fxAmount",
						"short": "For an FX payout this is the amound in the FX currency.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "fxCurrency",
						"short": "For an FX payout this is the currency that was received or that was instructed.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "fxRate",
						"short": "For an FX payout this is the exchange rate between the transaction currency and the FX currency.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "grossAmount",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique ID for the transaction.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inserted",
						"short": "Date when the transaction was inserted into the ledger.",
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
						"short": "The ID of the merchant that owns the account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pageNumber",
						"short": "Current page number.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "pageSize",
						"short": "Page size",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "payeeDetails",
						"req": true,
						"short": "The Payee object contains details of the beneficiary, person or business.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "payerDetails",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "paymentRequestCustomFields",
						"short": "The custom fields that were attached to the payment request that resulted in this transaction.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "paymentRequestID",
						"short": "For Pay by Bank and Direct Debit transactions this will contain the ID of the payment request.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "payoutID",
						"short": "ID of the payout that resulted in the transaction.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "proprietaryBankTransactionCode",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "rawReference",
						"short": "The raw payment reference details as received from the payment processor.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "reference",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "ruleID",
						"short": "ID of the rule that resulted in the transaction.",
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
						"short": "An optional list of descriptive tags attached to the transaction.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "theirReference",
						"short": "For a pay out the reference that the payer attached for the receiving party.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "totalPages",
						"short": "Total pages",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "totalSize",
						"short": "Total count",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "transactionAmount",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "transactionDate",
						"short": "Date when the transaction occurred.",
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
						"short": "Type of the transaction.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "valueDateTime",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "virtualIBAN",
						"short": "If set it indicates the payin was to a virtual IBAN.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "yourReference",
						"short": "For a pay in the reference the sending party attached.",
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
						"short": "The number of seconds a session for this user should last before expiring.",
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
						"short": "Optional ID of the invite that is being accepted so the user can be assigned a role on a new merchant.",
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
						"short": "The role ID to automatically assign to the merchant’s very first user.",
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
						"short": "Email address of the user being invited.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inviteeFirstName",
						"short": "First Name of the user being invited.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "inviteeLastName",
						"short": "Last Name of the user being invited.",
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
						"short": "Will be set to true once the invite has met the authorisation requirements.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isInviteeRegistered",
						"short": "If true, indicates the invitee's email address corresponds to an existing MoneyMoov user.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastInvited",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"short": "ID of the merchant the user is being invited to.",
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
						"short": "If set to true an email will be sent to the invitee with instructions on how to accept the invite.",
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
						"short": "Name for the account",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "accountSupplierName",
						"short": "The payment account supplier name.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "availableBalance",
						"short": "The current available balance of the account.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "availableBalanceMinorUnits",
						"short": "The available balance expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "balance",
						"short": "Balance of the account.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "balanceMinorUnits",
						"short": "Balance of the account expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "bankName",
						"short": "The bank name for external accounts",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "consentID",
						"short": "The ID of the consent used to connect the external account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "createdBy",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "createdByDisplayName",
						"short": "Either the name of the user, merchant token or api key that created the account",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "currency",
						"short": "Currency of the account in ISO 4217 format",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "defaultPaymentRail",
						"short": "Indicates the default payment rail for this account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "displayName",
						"short": "Gets a unique display name for the payment account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "expiryDate",
						"short": "The date that the external account will expire",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "externalAccountIcon",
						"short": "The Icon for external accounts",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "Unique id for the account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "identifier",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "inserted",
						"short": "Timestamp when the account was created.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "isArchived",
						"short": "Indicates whether the account is archived.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isConnectedAccount",
						"short": "Indicates if the payment account is an externally connected account.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isDefault",
						"short": "Is the default account",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isTrustAccount",
						"short": "Indicates if the payment account is a trust account.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "isVirtual",
						"short": "True if the account is a virtual account.",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "lastTransaction",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "lastUpdated",
						"short": "Timestamp when the account was last updated.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantID",
						"short": "The ID of the merchant that owns the account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "merchantName",
						"short": "The name of the merchant that owns the account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"req": true,
						"short": "The name of the virtual account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "physicalAccountID",
						"short": "For virtual accounts this is the ID of the physical account that the virtual account is linked to.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "rules",
						"short": "The list of rules associated with this account.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "submittedPayoutsBalance",
						"short": "Total of the payouts that have been submitted for processing.",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "submittedPayoutsBalanceMinorUnits",
						"short": "The balance of the submitted payouts expressed in the currency’s minor units (e.g.",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "summary",
						"short": "Gets a summary of the payments account's most important properties.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supplierSepaInstantStatus",
						"short": "Indicates the status of the SEPA Instant payment rail for this account.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "xeroBankFeedConnectionStatus",
						"short": "States the status of the Xero bank feed connection, if applicable.",
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
						"short": "Indicates the number of unsynchronised transactions with Xero",
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
						"short": "The destination URL for the webhook.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "emailAddress",
						"short": "The recipient email address(es) for notifications.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "failedNotificationEmailAddress",
						"short": "The email address to which notifications about failed webhook deliveries will be sent.",
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
						"short": "The ID of the merchant that the webhook is for.",
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
						"short": "The type of notification that will be sent.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "resourceTypes",
						"short": "The resource types that the webhook will be generated for.",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "retry",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "secret",
						"short": "The secret key required to authenticate webhook notifications.",
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
