-- Nofrixion SDK configuration

-- Build a fresh, fully materialised config table. Every call rebuilds the
-- whole structure, so prefer require("config_shared") unless you need a
-- private copy you intend to mutate.
local function make_config()
  return {
    main = {
      name = "Nofrixion",
      slug = "nofrixion",
      version = "0.1.1",
      target = "lua",
    },
    feature = {
      ["test"] = {
        ["options"] = {
          ["active"] = false,
        },
      },
    },
    options = {
      base = "https://api-sandbox.nofrixion.com",
      auth = {
        prefix = "Bearer",
      },
      headers = {
        ["content-type"] = "application/json",
      },
      entity = {
        ["account"] = {},
        ["batch"] = {},
        ["beneficiary"] = {},
        ["beneficiary_group"] = {},
        ["card"] = {},
        ["card_customer_token"] = {},
        ["card_payment"] = {},
        ["card_public_key"] = {},
        ["consent"] = {},
        ["currency"] = {},
        ["direct_debit_batch_submit"] = {},
        ["fx_rate"] = {},
        ["i_payment"] = {},
        ["mandate"] = {},
        ["merchant"] = {},
        ["merchant_authorisation_setting"] = {},
        ["merchant_direct_debit_mandate_page"] = {},
        ["merchant_pay_by_bank_setting"] = {},
        ["merchant_payment_request_template"] = {},
        ["merchant_token"] = {},
        ["metadata"] = {},
        ["no_frixion_version"] = {},
        ["open_banking"] = {},
        ["payeeverification"] = {},
        ["payment"] = {},
        ["payment_account"] = {},
        ["payment_account_minimal"] = {},
        ["payment_initiation"] = {},
        ["payment_request"] = {},
        ["payment_request_event"] = {},
        ["payment_request_metric"] = {},
        ["payment_request_minimal"] = {},
        ["payment_request_result"] = {},
        ["payout"] = {},
        ["payout_keyset_page"] = {},
        ["payout_metric"] = {},
        ["payrun"] = {},
        ["report"] = {},
        ["report_result"] = {},
        ["role"] = {},
        ["rule"] = {},
        ["rule_event"] = {},
        ["tag"] = {},
        ["token"] = {},
        ["transaction"] = {},
        ["user"] = {},
        ["user_invite"] = {},
        ["virtual"] = {},
        ["webhook"] = {},
      },
    },
    entity = {
      ["account"] = {
        ["fields"] = {
          {
            ["name"] = "accountBalances",
            ["short"] = "The various balances for the account.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "accountID",
            ["short"] = "ID of the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "accountIdentifications",
            ["short"] = "The canoncial identifiers for the account.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "accountName",
            ["short"] = "Name for the account",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "accountNames",
            ["short"] = "Optional account names set by the account holder.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "accountSupplierName",
            ["short"] = "The payment account supplier name.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "accountType",
            ["short"] = "The type of account e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "availableBalance",
            ["short"] = "The current available balance of the account.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "availableBalanceMinorUnits",
            ["short"] = "The available balance expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "balance",
            ["short"] = "Balance of the account.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "balanceMinorUnits",
            ["short"] = "Balance of the account expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "bankName",
            ["short"] = "The bank name for external accounts",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "consentID",
            ["short"] = "The ID of the consent used to connect the external account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "consolidatedAccountInformation",
            ["short"] = "Summary information regarding account balances of the overall account provided by the bank.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "createdBy",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "createdByDisplayName",
            ["short"] = "Either the name of the user, merchant token or api key that created the account",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currency",
            ["short"] = "Currency of the account in ISO 4217 format",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "defaultPaymentRail",
            ["short"] = "Indicates the default payment rail for this account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "description",
            ["short"] = "Product name as defined by the financial institution for this account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "details",
            ["short"] = "Supplementary specifications that might be provided by the Bank.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "displayName",
            ["short"] = "Gets a unique display name for the payment account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "expiryDate",
            ["short"] = "The date that the external account will expire",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "externalAccountIcon",
            ["short"] = "The Icon for external accounts",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "format",
            ["short"] = "File format to save the statement as.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fromDate",
            ["short"] = "Minimum transaction date for the statement.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique id for the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "identifier",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "inserted",
            ["short"] = "Timestamp when the account was created.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isArchived",
            ["short"] = "Indicates whether the account is archived.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isConnectedAccount",
            ["short"] = "Indicates if the payment account is an externally connected account.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isDefault",
            ["short"] = "Is the default account",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isTrustAccount",
            ["short"] = "Indicates if the payment account is a trust account.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isVirtual",
            ["short"] = "True if the account is a virtual account.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lastTransaction",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "lastUpdated",
            ["short"] = "Timestamp when the account was last updated.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["short"] = "The ID of the merchant that owns the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantName",
            ["short"] = "The name of the merchant that owns the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nickname",
            ["short"] = "Nickname of the account that was provided by the account owner.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "physicalAccountID",
            ["short"] = "For virtual accounts this is the ID of the physical account that the virtual account is linked to.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "roleIDs",
            ["short"] = "Optional list of role IDs that will get access to the payment account when created.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "rules",
            ["short"] = "The list of rules associated with this account.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "submittedPayoutsBalance",
            ["short"] = "Total of the payouts that have been submitted for processing.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "submittedPayoutsBalanceMinorUnits",
            ["short"] = "The balance of the submitted payouts expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "summary",
            ["short"] = "Gets a summary of the payments account's most important properties.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supplierPhysicalAccountID",
            ["short"] = "For internal use only.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supplierSepaInstantStatus",
            ["short"] = "Indicates the status of the SEPA Instant payment rail for this account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "toDate",
            ["short"] = "Maximum transaction date for the statement.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "type",
            ["short"] = "Specifies the type of account e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "usageType",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroBankFeedConnectionStatus",
            ["short"] = "States the status of the Xero bank feed connection, if applicable.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroBankFeedLastSyncedAt",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroBankFeedSyncLastFailedAt",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroBankFeedSyncLastFailureReason",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroBankFeedSyncStatus",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroUnsynchronisedTransactionsCount",
            ["short"] = "Indicates the number of unsynchronised transactions with Xero",
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "account",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/accounts/{accountID}/{currency}",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "{account_id}",
                  "{currency}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "currency",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/accounts/{accountID}/statements",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "{account_id}",
                  "statements",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                  },
                },
                ["select"] = {
                  ["$action"] = "statement",
                  ["exist"] = {
                    "account_id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["accountID"] = "`reqdata.account_id`",
                    ["format"] = "`reqdata.format`",
                    ["fromDate"] = "`reqdata.from_date`",
                    ["toDate"] = "`reqdata.to_date`",
                  },
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/accounts",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["accountName"] = "`reqdata.account_name`",
                    ["accountType"] = "`reqdata.account_type`",
                    ["currency"] = "`reqdata.currency`",
                    ["isTrustAccount"] = "`reqdata.is_trust_account`",
                    ["merchantID"] = "`reqdata.merchant_id`",
                    ["physicalAccountID"] = "`reqdata.physical_account_id`",
                    ["roleIDs"] = "`reqdata.role_i_d`",
                    ["supplierPhysicalAccountID"] = "`reqdata.supplier_physical_account_id`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "connected_account",
                      ["orig"] = "connected_account",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_archived",
                      ["orig"] = "include_archived",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "only_connect_account",
                      ["orig"] = "only_connect_account",
                      ["type"] = "`$BOOLEAN`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/accounts",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                },
                ["select"] = {
                  ["exist"] = {
                    "connected_account",
                    "include_archived",
                    "merchant_id",
                    "only_connect_account",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "connected_account",
                      ["orig"] = "connected_account",
                      ["type"] = "`$BOOLEAN`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/accounts",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "accounts",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "connected_account",
                    "merchant_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "connected_account",
                      ["orig"] = "connected_account",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_archived",
                      ["orig"] = "include_archived",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_child_merchant",
                      ["orig"] = "include_child_merchant",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "only_archived",
                      ["orig"] = "only_archived",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "only_connect_account",
                      ["orig"] = "only_connect_account",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/accounts/export",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "export",
                },
                ["select"] = {
                  ["$action"] = "export",
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "credit_type",
                      ["orig"] = "credit_type",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "max_amount",
                      ["orig"] = "max_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "min_amount",
                      ["orig"] = "min_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/accounts/{accountID}/transactions/export",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "{account_id}",
                  "transactions",
                  "export",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                  },
                },
                ["select"] = {
                  ["$action"] = "transaction_export",
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/accounts/{accountID}/statements/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "{account_id}",
                  "statements",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/accounts/{accountID}",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "accounts",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "id",
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                    "merchant_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/accounts/{accountID}",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/openbanking/accounts/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "openbanking",
                  "accounts",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/accounts/statements",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "statements",
                },
                ["select"] = {
                  ["$action"] = "statement",
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/accounts/archive/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "archive",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/accounts/statements",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "statements",
                },
                ["select"] = {
                  ["$action"] = "statement",
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "amount",
                      ["orig"] = "amount",
                      ["reqd"] = true,
                      ["type"] = "`$NUMBER`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/accounts/{accountID}/topup/{amount}",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "{account_id}",
                  "topup",
                  "{amount}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "amount",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/accounts/unarchive/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "unarchive",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/accounts/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["accountID"] = "`reqdata.account_id`",
                    ["accountName"] = "`reqdata.account_name`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "account",
            },
            {
              "merchant",
            },
            {
              "account",
              "topup",
            },
          },
        },
      },
      ["batch"] = {
        ["fields"] = {
          {
            ["name"] = "approveUrl",
            ["short"] = "This field is used when returning a batch payout record to a client.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payouts",
            ["type"] = "`$ARRAY`",
          },
        },
        ["name"] = "batch",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/payouts/batch",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "batch",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/payouts/batch/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "batch",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["beneficiary"] = {
        ["fields"] = {
          {
            ["name"] = "approvalCallbackUrl",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "authenticationMethods",
            ["short"] = "A list of authentication types allowed to authorise the payout.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "authorisations",
            ["short"] = "A list of users who have successfully authorised the latest version of the beneficiary.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "authorisersCompletedCount",
            ["short"] = "The number of distinct authorisers that have authorised the beneficiary.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "authorisersRequiredCount",
            ["short"] = "The number of authorisers required for this beneficiary.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "beneficiaries",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "beneficiaryEvents",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "canAuthorise",
            ["short"] = "True if the beneficiary can be authorised by the user who loaded it.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "canUpdate",
            ["short"] = "True if the beneficiary can be updated by the user who loaded it.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "createdBy",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "createdByEmailAddress",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currency",
            ["op"] = {
              ["update"] = {
                ["type"] = "`$STRING`",
              },
            },
            ["req"] = true,
            ["short"] = "Gets or Sets the currency.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "destination",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$OBJECT`",
              },
            },
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "failedBeneficiaries",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "hasCurrentUserAuthorised",
            ["short"] = "True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isEnabled",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lastAuthorised",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lastUpdated",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "Gets or Sets the merchant id.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "name",
            ["op"] = {
              ["update"] = {
                ["type"] = "`$STRING`",
              },
            },
            ["req"] = true,
            ["short"] = "The descriptive name for the beneficiary.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nonce",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sourceAccountIDs",
            ["short"] = "ID of the accounts which are authorised to act as a source for the beneficiary.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "sourceAccounts",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "theirReference",
            ["short"] = "The reference that will be used by default as TheirReference when creating payouts to this beneficiary if no TheirReference is specified for the payout.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "beneficiary",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/beneficiaries/authorise/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "beneficiaries",
                  "authorise",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/beneficiaries",
                ["parts"] = {
                  "api",
                  "v1",
                  "beneficiaries",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["currency"] = "`reqdata.currency`",
                    ["destination"] = "`reqdata.destination`",
                    ["id"] = "`reqdata.id`",
                    ["merchantID"] = "`reqdata.merchant_id`",
                    ["name"] = "`reqdata.name`",
                    ["sourceAccountIDs"] = "`reqdata.source_account_i_d`",
                    ["theirReference"] = "`reqdata.their_reference`",
                  },
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/beneficiaries/batchcreate",
                ["parts"] = {
                  "api",
                  "v1",
                  "beneficiaries",
                  "batchcreate",
                },
                ["select"] = {
                  ["$action"] = "batchcreate",
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_disabled",
                      ["orig"] = "include_disabled",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "source_account_id",
                      ["orig"] = "source_account_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/beneficiaries",
                ["parts"] = {
                  "api",
                  "v1",
                  "beneficiaries",
                },
                ["select"] = {
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_disabled",
                      ["orig"] = "include_disabled",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "source_account_id",
                      ["orig"] = "source_account_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/beneficiaries",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "beneficiaries",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_disabled",
                      ["orig"] = "include_disabled",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/beneficiaries/export",
                ["parts"] = {
                  "api",
                  "v1",
                  "beneficiaries",
                  "export",
                },
                ["select"] = {
                  ["$action"] = "export",
                  ["exist"] = {
                    "currency",
                    "include_disabled",
                    "merchant_id",
                    "page_number",
                    "page_size",
                    "search",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/beneficiaries/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "beneficiaries",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                    "merchant_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/beneficiaries/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "beneficiaries",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/beneficiaries/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "beneficiaries",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/beneficiaries/disable/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "beneficiaries",
                  "disable",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/beneficiaries/enable/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "beneficiaries",
                  "enable",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/beneficiaries/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "beneficiaries",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["currency"] = "`reqdata.currency`",
                    ["destination"] = "`reqdata.destination`",
                    ["name"] = "`reqdata.name`",
                    ["sourceAccountIDs"] = "`reqdata.source_account_i_d`",
                    ["theirReference"] = "`reqdata.their_reference`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "merchant",
            },
          },
        },
      },
      ["beneficiary_group"] = {
        ["fields"] = {
          {
            ["name"] = "groupMembers",
            ["short"] = "The existing group members.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "groupName",
            ["req"] = true,
            ["short"] = "The descriptive name for the beneficiary group.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["short"] = "Timestamp indicating when the group was created.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lastUpdated",
            ["short"] = "Timestamp indicating when the group was last updated.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["req"] = true,
            ["short"] = "Gets or Sets the merchant id.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "beneficiary_group",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/beneficiarygroups",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "beneficiarygroups",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "merchant_id",
                    "page_number",
                    "page_size",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "merchant",
            },
          },
        },
      },
      ["card"] = {
        ["fields"] = {
          {
            ["name"] = "authorizedAmount",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currencyCode",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isPayerAuthenticationRequired",
            ["short"] = "Gets set to true if 3-D Secure payer authentication is required for a payment.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isSoftDecline",
            ["short"] = "Gets set to true if the card processor flagged the transaction as having failed address or card security number verification.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "payerAuthenticationAccessToken",
            ["short"] = "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payerAuthenticationMerchantData",
            ["short"] = "If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the \"authenticationcallback\" method that gets called automatically after a successful payer authenticati…",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payerAuthenticationUrl",
            ["short"] = "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payerAuthenticationWindowHeight",
            ["short"] = "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "payerAuthenticationWindowWidth",
            ["short"] = "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "paymentRequestCallbackUrl",
            ["short"] = "The callback URL that was set when the payment request was created.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentRequestID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "requestID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "responseCode",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "responseType",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "status",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "threeDSRedirectUrl",
            ["short"] = "Checkout.com require a redirect for 3DS authentication.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "transactionID",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "card",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/paymentrequests/{id}/card",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "card",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "paymentrequest_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "paymentrequest",
            },
          },
        },
      },
      ["card_customer_token"] = {
        ["fields"] = {
          {
            ["name"] = "cardType",
            ["short"] = "The type of the tokenised card, e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerEmailAddress",
            ["short"] = "When creating a tokenised card the payer's email address must be supplied.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "expiryMonth",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "expiryYear",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "The unique ID of the card token that has been stored for the customer.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lastFourDigits",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lastUpdated",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "maskedCardNumber",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentRequestID",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "card_customer_token",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "customer_email_address",
                      ["orig"] = "customer_email_address",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/paymentrequests/card/customertokens/{merchantID}/{customerEmailAddress}",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "card",
                  "customertokens",
                  "{merchant_id}",
                  "{customer_email_address}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["customerEmailAddress"] = "customer_email_address",
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "customer_email_address",
                    "merchant_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "customer_email_address",
                      ["orig"] = "customer_email_address",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/paymentrequests/card/customertokens/{customerEmailAddress}",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "card",
                  "customertokens",
                  "{customer_email_address}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["customerEmailAddress"] = "customer_email_address",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "customer_email_address",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "customer_email_address",
                      ["orig"] = "customer_email_address",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/paymentrequests/card/customertokens/removeall/{merchantID}/{customerEmailAddress}",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "card",
                  "customertokens",
                  "removeall",
                  "{merchant_id}",
                  "{customer_email_address}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["customerEmailAddress"] = "customer_email_address",
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "customer_email_address",
                    "merchant_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "customer_email_address",
                      ["orig"] = "customer_email_address",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/paymentrequests/card/customertokens/removeall/{customerEmailAddress}",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "card",
                  "customertokens",
                  "removeall",
                  "{customer_email_address}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["customerEmailAddress"] = "customer_email_address",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "customer_email_address",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/paymentrequests/card/customertokens/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "card",
                  "customertokens",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "removeall",
            },
            {
              "customertoken",
            },
          },
        },
      },
      ["card_payment"] = {
        ["fields"] = {
          {
            ["name"] = "authorizedAmount",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currencyCode",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isPayerAuthenticationRequired",
            ["short"] = "Gets set to true if 3-D Secure payer authentication is required for a payment.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isSoftDecline",
            ["short"] = "Gets set to true if the card processor flagged the transaction as having failed address or card security number verification.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "payerAuthenticationAccessToken",
            ["short"] = "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payerAuthenticationMerchantData",
            ["short"] = "If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the \"authenticationcallback\" method that gets called automatically after a successful payer authenticati…",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payerAuthenticationUrl",
            ["short"] = "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payerAuthenticationWindowHeight",
            ["short"] = "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "payerAuthenticationWindowWidth",
            ["short"] = "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "paymentRequestCallbackUrl",
            ["short"] = "The callback URL that was set when the payment request was created.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentRequestID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "requestID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "responseCode",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "responseType",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "status",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "threeDSRedirectUrl",
            ["short"] = "Checkout.com require a redirect for 3DS authentication.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "transactionID",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "card_payment",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "partial_refund_amount",
                      ["orig"] = "partial_refund_amount",
                      ["reqd"] = true,
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/paymentrequests/{id}/card/refund/{partialRefundAmount}",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "card",
                  "refund",
                  "{partial_refund_amount}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "paymentrequest_id",
                    ["partialRefundAmount"] = "partial_refund_amount",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "partial_refund_amount",
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/paymentrequests/{id}/card/capture",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "card",
                  "capture",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "paymentrequest_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/paymentrequests/{id}/card/paywithtoken",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "card",
                  "paywithtoken",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "paymentrequest_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/paymentrequests/{id}/card/void",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "card",
                  "void",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "paymentrequest_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/paymentrequests/{id}/card/voidpaymentrequest",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "card",
                  "voidpaymentrequest",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "paymentrequest_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "paymentrequest",
            },
            {
              "paymentrequest",
              "refund",
            },
          },
        },
      },
      ["card_public_key"] = {
        ["fields"] = {
          {
            ["name"] = "jwt",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "card_public_key",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/paymentrequests/{id}/card/publickey",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "card",
                  "publickey",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "paymentrequest_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "paymentrequest",
            },
          },
        },
      },
      ["consent"] = {
        ["fields"] = {
          {
            ["name"] = "authorisationUrl",
            ["short"] = "The URL the authorising user needs to be redirected to in order to get the open banking consent token.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "callbackUrl",
            ["short"] = "Optional callback URL that the end user performing the open banking authorisation will be redirected to on completion.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "consentID",
            ["short"] = "The ID of the open banking consent.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "emailAddress",
            ["short"] = "The email address that identifies the end user that will be authorising the open banking consent request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "expiryDate",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "failureCallbackUrl",
            ["short"] = "Optional callback URL for open banking consent authorisation failure.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "institutionID",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "The institution ID the open banking consent is being requested for.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isConnectedAccounts",
            ["short"] = "Optional setting.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isEnabled",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "merchantID",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "The ID of the merchant the consent token is being created to be used with.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "provider",
            ["short"] = "Lists the supported card and PIS processors.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "successWebHookUrl",
            ["short"] = "A web hook URL to send an HTTP request to when the open banking consent is successfuly authorised.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "consent",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/openbanking/consents",
                ["parts"] = {
                  "api",
                  "v1",
                  "openbanking",
                  "consents",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["callbackUrl"] = "`reqdata.callback_url`",
                    ["emailAddress"] = "`reqdata.email_address`",
                    ["failureCallbackUrl"] = "`reqdata.failure_callback_url`",
                    ["institutionID"] = "`reqdata.institution_id`",
                    ["isConnectedAccounts"] = "`reqdata.is_connected_account`",
                    ["merchantID"] = "`reqdata.merchant_id`",
                    ["successWebHookUrl"] = "`reqdata.success_web_hook_url`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "email",
                      ["orig"] = "email",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/openbanking/consents/{merchantID}/{email}",
                ["parts"] = {
                  "api",
                  "v1",
                  "openbanking",
                  "consents",
                  "{merchant_id}",
                  "{email}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "email",
                    "merchant_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/openbanking/consents/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "openbanking",
                  "consents",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/openbanking/consents/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "openbanking",
                  "consents",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PATCH",
                ["orig"] = "/api/v1/openbanking/consents/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "openbanking",
                  "consents",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "consent",
            },
          },
        },
      },
      ["currency"] = {
        ["fields"] = {
          {
            ["name"] = "code",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "decimals",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "isFiat",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "iso4217AlphaCode",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "iso4217NumericCode",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "symbol",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "currency",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "capability",
                      ["orig"] = "capability",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/currencies",
                ["parts"] = {
                  "api",
                  "v1",
                  "currencies",
                },
                ["select"] = {
                  ["exist"] = {
                    "capability",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["direct_debit_batch_submit"] = {
        ["fields"] = {
          {
            ["name"] = "failedSubmissions",
            ["short"] = "Dictionary of failed submissions, keyed by the index (1-based) in the original request.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "successfulSubmissions",
            ["short"] = "List of successfully submitted direct debit payments.",
            ["type"] = "`$ARRAY`",
          },
        },
        ["name"] = "direct_debit_batch_submit",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/paymentrequests/directdebit/batchsubmit",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "directdebit",
                  "batchsubmit",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["fx_rate"] = {
        ["fields"] = {
          {
            ["name"] = "destinationCurrency",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "exchangeRate",
            ["short"] = "The price at which the transaction will buy the source currency using the destination currency.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "expiryTime",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "quoteID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sourceCurrency",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "fx_rate",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "destination",
                      ["orig"] = "destination",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "source",
                      ["orig"] = "source",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/payouts/fxallheldrates/{source}/{destination}",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "fxallheldrates",
                  "{source}",
                  "{destination}",
                },
                ["select"] = {
                  ["exist"] = {
                    "destination",
                    "source",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "destination",
                      ["orig"] = "destination",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "source",
                      ["orig"] = "source",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "valid_for_minute",
                      ["orig"] = "valid_for_minute",
                      ["reqd"] = true,
                      ["type"] = "`$INTEGER`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/payouts/fxheldrate/{source}/{destination}/{validForMinutes}",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "fxheldrate",
                  "{source}",
                  "{destination}",
                  "{valid_for_minute}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["validForMinutes"] = "valid_for_minute",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "destination",
                    "source",
                    "valid_for_minute",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "fxallheldrate",
            },
            {
              "fxheldrate",
            },
          },
        },
      },
      ["i_payment"] = {
        ["fields"] = {
          {
            ["name"] = "paymentRequestID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "responseType",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "i_payment",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/paymentrequests/payondemand",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "payondemand",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["mandate"] = {
        ["fields"] = {
          {
            ["name"] = "accountNumber",
            ["short"] = "Account number of the customer's bank account in case of GBP account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "addressLine1",
            ["req"] = true,
            ["short"] = "First line of the customer's address.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "addressLine2",
            ["short"] = "Second line of the customer's address.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "approvedAt",
            ["short"] = "Date at which the supplier approved this mandate.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "city",
            ["req"] = true,
            ["short"] = "Customer's city.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "countryCode",
            ["req"] = true,
            ["short"] = "2-character country code of the customer's bank account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currency",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "Currency of this mandate.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerAccountNumber",
            ["short"] = "Customer's account number in case of GBP account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerCity",
            ["short"] = "Customer's city of residence.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerCountryCode",
            ["short"] = "Customer's country of residence code.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerCountryName",
            ["short"] = "Customer's country of residence.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerEmailAddress",
            ["short"] = "Customer's email address.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerFirstName",
            ["short"] = "Customer's first name.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerIban",
            ["short"] = "Customer's IBAN in case of EUR account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerLastName",
            ["short"] = "Customer's last name.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerSortCode",
            ["short"] = "Customer's sort code in case of GBP account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "emailAddress",
            ["req"] = true,
            ["short"] = "Customer's email address.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "firstName",
            ["req"] = true,
            ["short"] = "Customer's first name.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "iban",
            ["short"] = "IBAN of the customer's bank account in case of EUR account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "Internal ID of the mandate.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["short"] = "The timestamp this mandate was created at.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isRecurring",
            ["short"] = "Whether this mandate is single-use or recurring.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lastName",
            ["req"] = true,
            ["short"] = "Customer's last name.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lastUpdated",
            ["short"] = "The timestamp this mandate was last updated at.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "Internal ID of this mandate's merchant.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "postalCode",
            ["req"] = true,
            ["short"] = "Customer's postal code.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "reference",
            ["short"] = "Reference assigned to this mandate.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sortCode",
            ["short"] = "Sort code of the customer's bank account in case of GBP account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "status",
            ["short"] = "General status of this mandate.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supplierBankAccountID",
            ["short"] = "ID that the supplier assigned to this mandate's bank account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supplierCustomerID",
            ["short"] = "ID that the supplier assigned to this mandate's customer.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supplierMandateID",
            ["short"] = "ID that the supplier assigned to this mandate.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supplierName",
            ["short"] = "Name of the supplier used to create this mandate.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supplierStatus",
            ["short"] = "Last status that the supplier reported for this mandate.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "mandate",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/mandates",
                ["parts"] = {
                  "api",
                  "v1",
                  "mandates",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["accountNumber"] = "`reqdata.account_number`",
                    ["addressLine1"] = "`reqdata.address_line1`",
                    ["addressLine2"] = "`reqdata.address_line2`",
                    ["city"] = "`reqdata.city`",
                    ["countryCode"] = "`reqdata.country_code`",
                    ["currency"] = "`reqdata.currency`",
                    ["emailAddress"] = "`reqdata.email_address`",
                    ["firstName"] = "`reqdata.first_name`",
                    ["iban"] = "`reqdata.iban`",
                    ["isRecurring"] = "`reqdata.is_recurring`",
                    ["lastName"] = "`reqdata.last_name`",
                    ["merchantID"] = "`reqdata.merchant_id`",
                    ["postalCode"] = "`reqdata.postal_code`",
                    ["reference"] = "`reqdata.reference`",
                    ["sortCode"] = "`reqdata.sort_code`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/mandates/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "mandates",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["merchant"] = {
        ["fields"] = {
          {
            ["name"] = "accountCurrencies",
            ["short"] = "The list of currencies that the merchant has accounts for.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "canHaveTrustAccounts",
            ["short"] = "Trust accounts are a special type of account that allow the account name to be trusted for use in statements and verification of payee checks.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "cardPaymentProcessor",
            ["short"] = "Name of the card payment processor.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "companyID",
            ["short"] = "The Company ID recorded in the Compliance system.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "displayQrOnHostedPay",
            ["short"] = "Indicates if a QR Code containing the payment link should be displayed on the hosted payment page.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "hostedPayVersion",
            ["short"] = "The version of the hosted payment page to use with the merchant.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique ID for the merchant.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["short"] = "Timestamp the merchant was added to MoneyMoov.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isBlocked",
            ["short"] = "The merchant is blocked from making payments (payouts).",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isExited",
            ["short"] = "The merchant has formally terminated their relationship and is no longer a customer.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isSuspended",
            ["short"] = "The merchant has temporarily suspended their own account.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "jurisdiction",
            ["short"] = "The jurisdiction the merchant entity is incorporated or established in.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "logoUrlPng",
            ["short"] = "The CDN URL of the merchant's logo in PNG format.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "logoUrlSvg",
            ["short"] = "The CDN URL of the merchant's logo in SVG format.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantCategoryCode",
            ["short"] = "The industry code that represents the merchant's primary trading activity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "name",
            ["short"] = "The registered business name of the merchant.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "notes",
            ["short"] = "The notes field is an optional free text field that can be used to store any additional information about the merchant.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "parentMerchant",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "paymentAccountLimit",
            ["short"] = "The maximum number of payment accounts that can be created for the Merchant.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "paymentAccounts",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "reason",
            ["short"] = "The reason for the suspension.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "shortName",
            ["short"] = "A URL friendly shortish name for the merchant.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supportedPaymentMethodsList",
            ["short"] = "The payment methods that are configured and supported for this merchant.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "suspensionReason",
            ["short"] = "The reason for the suspension, provided by the merchant.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tags",
            ["short"] = "An optional list of descriptive tags that can be used on merchant entities such as payment requests.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "timeZoneId",
            ["short"] = "The IANA (Internet Assigned Numbers Authority) time zone identifier of the merchant.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tradingName",
            ["short"] = "An optional trading name.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "webHookLimit",
            ["short"] = "The maximum number of web hooks that can be created for the Merchant.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "yourRoleName",
            ["short"] = "The name of the role for the identity that loaded the merchant record.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "merchant",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["example"] = 1,
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/childmerchants",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "childmerchants",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["$action"] = "childmerchant",
                  ["exist"] = {
                    "merchant_id",
                    "page_number",
                    "page_size",
                    "search",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_suspended",
                      ["orig"] = "include_suspended",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["example"] = 1,
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = 10,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/paged",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "paged",
                },
                ["select"] = {
                  ["$action"] = "paged",
                  ["exist"] = {
                    "include_suspended",
                    "page_number",
                    "page_size",
                    "search",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/metadata/whoamimerchant",
                ["parts"] = {
                  "api",
                  "v1",
                  "metadata",
                  "whoamimerchant",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/metadata/whoamimerchantsigned",
                ["parts"] = {
                  "api",
                  "v1",
                  "metadata",
                  "whoamimerchantsigned",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/metadata/whoamimerchantwhitelist",
                ["parts"] = {
                  "api",
                  "v1",
                  "metadata",
                  "whoamimerchantwhitelist",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "max_amount",
                      ["orig"] = "max_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "min_amount",
                      ["orig"] = "min_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "status",
                      ["orig"] = "status",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "tag",
                      ["orig"] = "tag",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/payouts/export",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "payouts",
                  "export",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["$action"] = "payout_export",
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_disabled",
                      ["orig"] = "include_disabled",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/beneficiaries/export",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "beneficiaries",
                  "export",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["$action"] = "beneficiary_export",
                  ["exist"] = {
                    "currency",
                    "include_disabled",
                    "merchant_id",
                    "page_number",
                    "page_size",
                    "search",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "user_id",
                      ["orig"] = "user_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/merchants/{merchantId}/users/{userId}",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{id}",
                  "users",
                  "{user_id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantId"] = "id",
                    ["userId"] = "user_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                    "user_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "tag_id",
                      ["orig"] = "tag_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/merchants/{merchantID}/tags/{tagID}",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "tags",
                  "{tag_id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                    ["tagID"] = "tag_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "merchant_id",
                    "tag_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/merchants/{merchantID}",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["logoUrlPng"] = "`reqdata.logo_url_png`",
                    ["logoUrlSvg"] = "`reqdata.logo_url_svg`",
                    ["notes"] = "`reqdata.note`",
                    ["paymentAccountLimit"] = "`reqdata.payment_account_limit`",
                    ["shortName"] = "`reqdata.short_name`",
                  },
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/merchants/{merchantId}/suspend",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{id}",
                  "suspend",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantId"] = "id",
                  },
                },
                ["select"] = {
                  ["$action"] = "suspend",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["reason"] = "`reqdata.reason`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "merchant",
            },
            {
              "user",
            },
            {
              "merchant",
              "tag",
            },
          },
        },
      },
      ["merchant_authorisation_setting"] = {
        ["fields"] = {
          {
            ["name"] = "amountLower",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountUpper",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "authorisationType",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "beneficiariesOnly",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lastEditorCantAuthorise",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lastUpdated",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "numberOfAuthorisers",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "roleSettings",
            ["type"] = "`$ARRAY`",
          },
        },
        ["name"] = "merchant_authorisation_setting",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/authorisationsettings",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "authorisationsettings",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "merchant_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "merchant",
            },
          },
        },
      },
      ["merchant_direct_debit_mandate_page"] = {
        ["fields"] = {
          {
            ["name"] = "approvedAt",
            ["short"] = "Date at which the supplier approved this mandate.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currency",
            ["short"] = "Currency of this mandate.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerAccountNumber",
            ["short"] = "Customer's account number in case of GBP account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerCity",
            ["short"] = "Customer's city of residence.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerCountryCode",
            ["short"] = "Customer's country of residence code.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerCountryName",
            ["short"] = "Customer's country of residence.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerEmailAddress",
            ["short"] = "Customer's email address.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerFirstName",
            ["short"] = "Customer's first name.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerIban",
            ["short"] = "Customer's IBAN in case of EUR account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerLastName",
            ["short"] = "Customer's last name.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerSortCode",
            ["short"] = "Customer's sort code in case of GBP account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "Internal ID of the mandate.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["short"] = "The timestamp this mandate was created at.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isRecurring",
            ["short"] = "Whether this mandate is single-use or recurring.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lastUpdated",
            ["short"] = "The timestamp this mandate was last updated at.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["short"] = "Internal ID of this mandate's merchant.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "reference",
            ["short"] = "Reference assigned to this mandate.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "status",
            ["short"] = "General status of this mandate.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supplierBankAccountID",
            ["short"] = "ID that the supplier assigned to this mandate's bank account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supplierCustomerID",
            ["short"] = "ID that the supplier assigned to this mandate's customer.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supplierMandateID",
            ["short"] = "ID that the supplier assigned to this mandate.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supplierName",
            ["short"] = "Name of the supplier used to create this mandate.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supplierStatus",
            ["short"] = "Last status that the supplier reported for this mandate.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "merchant_direct_debit_mandate_page",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "mandate_i_d",
                      ["orig"] = "mandate_i_d",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "max_amount",
                      ["orig"] = "max_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "min_amount",
                      ["orig"] = "min_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = 1,
                      ["kind"] = "query",
                      ["name"] = "page",
                      ["orig"] = "page",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "status",
                      ["orig"] = "status",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/mandates",
                ["parts"] = {
                  "api",
                  "v1",
                  "mandates",
                },
                ["select"] = {
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["merchant_pay_by_bank_setting"] = {
        ["fields"] = {
          {
            ["name"] = "bankCountryCodes",
            ["short"] = "The list of country codes representing the banks the country supports.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "bankID",
            ["short"] = "ID of the bank to be configured for the merchant.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "bankName",
            ["short"] = "Name of the Bank/Institution.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "businessInstitutionID",
            ["short"] = "ID that the processor uses to identify the bank (business accounts).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currency",
            ["short"] = "Currency supported by the bank.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "logo",
            ["short"] = "URL of the bank's logo.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "message",
            ["short"] = "Message relating to specific bank.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "messageImageUrl",
            ["short"] = "Optional image URL to be displayed with the message.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "order",
            ["short"] = "Order in which this setting will appear in the UI.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "personalInstitutionID",
            ["short"] = "ID that the processor uses to identify the bank (personal accounts).",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "processor",
            ["short"] = "Name of the bank payment processor.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "warningHeading",
            ["short"] = "The heading for a warning message related to the bank institution to be displayed to the user.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "warningMessage",
            ["short"] = "The warning message related to the bank institution to be displayed to the user.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "merchant_pay_by_bank_setting",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "country_code",
                      ["orig"] = "country_code",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = "PIS",
                      ["kind"] = "query",
                      ["name"] = "open_banking_operation",
                      ["orig"] = "open_banking_operation",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/banksettings",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "banksettings",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "country_code",
                    "currency",
                    "merchant_id",
                    "open_banking_operation",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.payByBankSettings`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "merchant",
            },
          },
        },
      },
      ["merchant_payment_request_template"] = {
        ["fields"] = {
          {
            ["name"] = "bankPaymentOptions",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "cardPaymentAddressOptions",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "cardPaymentCaptureOptions",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "customFields",
            ["short"] = "A list of custom fields that can be included in the payment request template.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "defaultFields",
            ["short"] = "A list of default fields that are included in the payment request template.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "description",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lastUpdated",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "name",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "notificationOptions",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "paymentMethods",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "paymentTerms",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "priorityBankOptions",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "template",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
        },
        ["name"] = "merchant_payment_request_template",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/paymentrequests/{merchantID}/templates",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{merchant_id}",
                  "templates",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "merchant_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "template_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/paymentrequests/{merchantID}/templates/{templateID}",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "templates",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "paymentrequest_id",
                    ["templateID"] = "id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.template`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "template_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/paymentrequests/{merchantID}/templates/{templateID}",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "templates",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "paymentrequest_id",
                    ["templateID"] = "id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.template`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "template_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/paymentrequests/{merchantID}/templates/{templateID}",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "templates",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "paymentrequest_id",
                    ["templateID"] = "id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["description"] = "`reqdata.description`",
                    ["id"] = "`reqdata.id`",
                    ["merchantID"] = "`reqdata.merchant_id`",
                    ["name"] = "`reqdata.name`",
                    ["template"] = "`reqdata.template`",
                  },
                  ["res"] = "`body.template`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "paymentrequest",
            },
          },
        },
      },
      ["merchant_token"] = {
        ["fields"] = {
          {
            ["name"] = "authenticationMethods",
            ["short"] = "A list of authentication types allowed to authorise the merchant token.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "authorisations",
            ["short"] = "A list of users who have successfully authorised the latest version of the beneficiary.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "authorisersCompletedCount",
            ["short"] = "The number of distinct authorisers that have authorised the merchant token.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "authorisersRequiredCount",
            ["short"] = "The number of authorisers required for this merchant token.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "canAuthorise",
            ["short"] = "True if the merchant token can be authorised by the user who loaded it.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "description",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "Token description",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "expiresAt",
            ["short"] = "Optional.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "hasCurrentUserAuthorised",
            ["short"] = "True if the beneficiary was loaded for a user and that user has already authorised the latest version of the beneficiary.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "hmacAlgorithm",
            ["short"] = "Optional shared secret algorithm to use for HMAC authentication.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ipAddressWhitelist",
            ["short"] = "Optional.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isArchived",
            ["short"] = "Indicates whether the merchant token is archived.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isEnabled",
            ["short"] = "If set to false the merchant token will not be accepted to authorise a request.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lastAuthorised",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lastUpdated",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "The merchant id to add to the token",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nonce",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "permissionTypes",
            ["short"] = "The permissions that the merchant token supports.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "requestSignatureVersion",
            ["short"] = "Represent the version of the overall merchant token.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "sharedSecretAlgorithm",
            ["short"] = "Optional shared secret algorithm to use for HMAC authentication.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sharedSecretBase64",
            ["short"] = "The base 64 encoded shared secret that is used for request authentication with an HMAC.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "token",
            ["short"] = "The JWT merchant token.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "merchant_token",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/tokens",
                ["parts"] = {
                  "api",
                  "v1",
                  "tokens",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["description"] = "`reqdata.description`",
                    ["hmacAlgorithm"] = "`reqdata.hmac_algorithm`",
                    ["ipAddressWhitelist"] = "`reqdata.ip_address_whitelist`",
                    ["merchantID"] = "`reqdata.merchant_id`",
                    ["permissionTypes"] = "`reqdata.permission_type`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/tokens",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "tokens",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "merchant_id",
                    "page_number",
                    "page_size",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/tokens/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "tokens",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/tokens/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "tokens",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["description"] = "`reqdata.description`",
                    ["ipAddressWhitelist"] = "`reqdata.ip_address_whitelist`",
                    ["merchantID"] = "`reqdata.merchant_id`",
                    ["permissionTypes"] = "`reqdata.permission_type`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "merchant",
            },
          },
        },
      },
      ["metadata"] = {
        ["fields"] = {},
        ["name"] = "metadata",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "secret",
                      ["orig"] = "secret",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "url",
                      ["orig"] = "url",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/metadata/problemnotification",
                ["parts"] = {
                  "api",
                  "v1",
                  "metadata",
                  "problemnotification",
                },
                ["select"] = {
                  ["$action"] = "problemnotification",
                  ["exist"] = {
                    "id",
                    "secret",
                    "url",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/metadata/problem",
                ["parts"] = {
                  "api",
                  "v1",
                  "metadata",
                  "problem",
                },
                ["select"] = {
                  ["$action"] = "problem",
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["no_frixion_version"] = {
        ["fields"] = {
          {
            ["name"] = "buildVersion",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "majorVersion",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "minorVersion",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "releaseName",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "no_frixion_version",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/metadata/version",
                ["parts"] = {
                  "api",
                  "v1",
                  "metadata",
                  "version",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["open_banking"] = {
        ["fields"] = {},
        ["name"] = "open_banking",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/openbanking/account/{accountID}/synchronise",
                ["parts"] = {
                  "api",
                  "v1",
                  "openbanking",
                  "account",
                  "{account_id}",
                  "synchronise",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "email",
                      ["orig"] = "email",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/openbanking/consents/{merchantID}/{email}",
                ["parts"] = {
                  "api",
                  "v1",
                  "openbanking",
                  "consents",
                  "{merchant_id}",
                  "{email}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "email",
                    "merchant_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/openbanking/account/{accountID}",
                ["parts"] = {
                  "api",
                  "v1",
                  "openbanking",
                  "account",
                  "{account_id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "account",
            },
            {
              "consent",
            },
          },
        },
      },
      ["payeeverification"] = {
        ["fields"] = {
          {
            ["name"] = "accountName",
            ["req"] = true,
            ["short"] = "The name of the account to verify",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "accountNumber",
            ["short"] = "The account number of the account to verify (for CoP checks)",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "iban",
            ["req"] = true,
            ["short"] = "The IBAN of the account to verify (for VoP checks)",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payeeVerifiedAccountName",
            ["short"] = "The verified account name of the payee, if available (in case of a close match)",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "result",
            ["short"] = "The result of the payee verification",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "secondaryIdentification",
            ["short"] = "Optional secondary identifier for the account to verify.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sortCode",
            ["short"] = "The sort code of the account to verify (for CoP checks)",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "payeeverification",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/openbanking/payeeverification",
                ["parts"] = {
                  "api",
                  "v1",
                  "openbanking",
                  "payeeverification",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["accountName"] = "`reqdata.account_name`",
                    ["accountNumber"] = "`reqdata.account_number`",
                    ["iban"] = "`reqdata.iban`",
                    ["secondaryIdentification"] = "`reqdata.secondary_identification`",
                    ["sortCode"] = "`reqdata.sort_code`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["payment"] = {
        ["fields"] = {
          {
            ["name"] = "addresses",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "amount",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$NUMBER`",
              },
            },
            ["short"] = "The amount of money to request.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountPending",
            ["short"] = "Total amount that has been authorised but not settled for this payment request.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountReceived",
            ["short"] = "Total amount received for this payment request.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountRefunded",
            ["short"] = "Total amount refunded for this payment request.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "autoSendReceipt",
            ["short"] = "If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "baseOriginUrl",
            ["short"] = "For card payments the origin of the payment page needs to be set in advance.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "callbackUrl",
            ["short"] = "Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardAuthorizeOnly",
            ["short"] = "For card payments the default behaviour is to authorise and capture the payment at the same time.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "cardCreateToken",
            ["short"] = "For card payments a payment attempt can be used to create a reusable token for subsequent payments.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "cardCreateTokenMode",
            ["short"] = "This specifies whether user consent will be taken before tokenising card or not.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardIgnoreCVN",
            ["short"] = "If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "cardNoPayerAuthentication",
            ["short"] = "If set to true for card payments no attempt will be made to use payer authentication (3-D Secure and equivalent).",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "cardProcessorMerchantID",
            ["short"] = "Optional field that if specified indicates the processor merchant ID that should be used to process any card payments.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardStripePaymentIntentID",
            ["short"] = "If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardStripePaymentIntentSecret",
            ["short"] = "If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardTransmitRawDetails",
            ["short"] = "If set to true for card payments the sensitive card number and card verification number will be transmitted directly rather than being tokenised.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "createdByUser",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "currency",
            ["short"] = "The currency of the request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customFields",
            ["short"] = "A list of custom fields attached to the payment request.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "customerEmailAddress",
            ["short"] = "Optional email address for the customer.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerID",
            ["short"] = "An optional customer identifier for the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerName",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "description",
            ["short"] = "An optional description for the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "destinationAccount",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "directDebitPayment",
            ["short"] = "Contains information about a Direct Debit payment attempt for a payment request.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "dueDate",
            ["short"] = "The due date for the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "events",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "failureCallbackUrl",
            ["short"] = "Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fieldDisplaySettings",
            ["short"] = "A list of field display settings that control which fields are displayed to the payer.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "formattedAmount",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "hostedPayCheckoutUrl",
            ["short"] = "This is a convenience link generated for payment requests whose merchants are using hosted payment pages.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ignoreAddressVerification",
            ["short"] = "If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "inserted",
            ["short"] = "The timestamp the payment request was created at.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "insertedSortable",
            ["short"] = "The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se…",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isArchived",
            ["short"] = "Indicates whether the payment request is archived.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "jwk",
            ["short"] = "The jwk containing the public key used to verify the signature of the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lastUpdated",
            ["short"] = "The timestamp the payment request was last updated at.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lightningInvoice",
            ["short"] = "Bitcoin Lightning invoice for the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lightningInvoiceExpiresAt",
            ["short"] = "Date and time of expiration of the lightning invoice.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantDirectDebitMandateID",
            ["short"] = "Optional ID of the direct debit mandate associated with this payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["short"] = "The ID of the merchant to create the payment request for.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantTokenDescription",
            ["short"] = "Description of the merchant token in case the Payment request was created using a merchant token.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "notificationEmailAddresses",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "notificationRoleIDs",
            ["short"] = "A list of roles whose members will receive notifications about this payment request.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "orderID",
            ["short"] = "An optional order ID for the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "partialPaymentMethod",
            ["short"] = "The approach to use, or not, for accepting partial payments.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "partialPaymentSteps",
            ["short"] = "An optional comma separated list of partial payment amounts.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentAttempts",
            ["short"] = "The payment attempts made against this payment request.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "paymentMethods",
            ["short"] = "The payment methods that the payment request supports.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "paymentProcessor",
            ["short"] = "If the card payment option is enabled this field indicates which card processor the merchant is set up to use.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payrunID",
            ["short"] = "The ID of a payrun that needs an account top up.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "pispAccountID",
            ["short"] = "The payment account ID to use to receive payment initiation payments.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "priorityBankID",
            ["short"] = "The ID of the bank that is set as the priority bank for display on pay element.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "result",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sandboxSettleDelayInSeconds",
            ["short"] = "Sandbox only.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "shippingAddress",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "shippingAddressCity",
            ["short"] = "Optionally the city of the customer's shipping address.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "shippingAddressCountryCode",
            ["short"] = "Optionally the country code of the customer's shipping address.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "shippingAddressCounty",
            ["short"] = "Optionally the state or county of the customer's shipping address.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "shippingAddressLine1",
            ["short"] = "Optionally the first line of the customer's shipping address.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "shippingAddressLine2",
            ["short"] = "Optionally the second line of the customer's shipping address.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "shippingAddressPostCode",
            ["short"] = "Optionally the post code of the customer's shipping address.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "shippingEmail",
            ["short"] = "Optionally the shipping email address for the customer.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "shippingFirstName",
            ["short"] = "Optionally the first name of the customer's shipping address.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "shippingLastName",
            ["short"] = "Optionally the last name of the customer's shipping address.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "shippingPhone",
            ["short"] = "Optionally the shipping phone number for the customer.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "status",
            ["short"] = "The current status of the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "successWebHookUrl",
            ["short"] = "If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tagIds",
            ["short"] = "An optional list of tag ids to add to the payment request",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "tags",
            ["short"] = "An optional list of descriptive tags attached to the payment request.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "title",
            ["short"] = "A generic field to contain any additional data that the merchant wishes to store against the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tokenisedCards",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "transactions",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "useHostedPaymentPage",
            ["short"] = "If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page.",
            ["type"] = "`$BOOLEAN`",
          },
        },
        ["name"] = "payment",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/paymentrequests",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["amount"] = "`reqdata.amount`",
                    ["autoSendReceipt"] = "`reqdata.auto_send_receipt`",
                    ["callbackUrl"] = "`reqdata.callback_url`",
                    ["cardAuthorizeOnly"] = "`reqdata.card_authorize_only`",
                    ["cardCreateToken"] = "`reqdata.card_create_token`",
                    ["cardCreateTokenMode"] = "`reqdata.card_create_token_mode`",
                    ["cardIgnoreCVN"] = "`reqdata.card_ignore_cvn`",
                    ["cardNoPayerAuthentication"] = "`reqdata.card_no_payer_authentication`",
                    ["cardProcessorMerchantID"] = "`reqdata.card_processor_merchant_id`",
                    ["cardTransmitRawDetails"] = "`reqdata.card_transmit_raw_detail`",
                    ["currency"] = "`reqdata.currency`",
                    ["customFields"] = "`reqdata.custom_field`",
                    ["customerEmailAddress"] = "`reqdata.customer_email_address`",
                    ["customerID"] = "`reqdata.customer_id`",
                    ["description"] = "`reqdata.description`",
                    ["dueDate"] = "`reqdata.due_date`",
                    ["failureCallbackUrl"] = "`reqdata.failure_callback_url`",
                    ["fieldDisplaySettings"] = "`reqdata.field_display_setting`",
                    ["ignoreAddressVerification"] = "`reqdata.ignore_address_verification`",
                    ["merchantDirectDebitMandateID"] = "`reqdata.merchant_direct_debit_mandate_id`",
                    ["merchantID"] = "`reqdata.merchant_id`",
                    ["notificationEmailAddresses"] = "`reqdata.notification_email_address`",
                    ["notificationRoleIDs"] = "`reqdata.notification_role_i_d`",
                    ["orderID"] = "`reqdata.order_id`",
                    ["partialPaymentMethod"] = "`reqdata.partial_payment_method`",
                    ["partialPaymentSteps"] = "`reqdata.partial_payment_step`",
                    ["paymentMethods"] = "`reqdata.payment_method`",
                    ["payrunID"] = "`reqdata.payrun_id`",
                    ["pispAccountID"] = "`reqdata.pisp_account_id`",
                    ["priorityBankID"] = "`reqdata.priority_bank_id`",
                    ["sandboxSettleDelayInSeconds"] = "`reqdata.sandbox_settle_delay_in_second`",
                    ["shippingAddressCity"] = "`reqdata.shipping_address_city`",
                    ["shippingAddressCountryCode"] = "`reqdata.shipping_address_country_code`",
                    ["shippingAddressCounty"] = "`reqdata.shipping_address_county`",
                    ["shippingAddressLine1"] = "`reqdata.shipping_address_line1`",
                    ["shippingAddressLine2"] = "`reqdata.shipping_address_line2`",
                    ["shippingAddressPostCode"] = "`reqdata.shipping_address_post_code`",
                    ["shippingEmail"] = "`reqdata.shipping_email`",
                    ["shippingFirstName"] = "`reqdata.shipping_first_name`",
                    ["shippingLastName"] = "`reqdata.shipping_last_name`",
                    ["shippingPhone"] = "`reqdata.shipping_phone`",
                    ["successWebHookUrl"] = "`reqdata.success_web_hook_url`",
                    ["tagIds"] = "`reqdata.tag_id`",
                    ["tags"] = "`reqdata.tag`",
                    ["title"] = "`reqdata.title`",
                    ["useHostedPaymentPage"] = "`reqdata.use_hosted_payment_page`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_event",
                      ["orig"] = "include_event",
                      ["type"] = "`$BOOLEAN`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/paymentrequests/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                    "include_event",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "order_id",
                      ["orig"] = "order_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/paymentrequests/getbyorderid/{orderID}",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "getbyorderid",
                  "{order_id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["orderID"] = "order_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "order_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/paymentrequests/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["amount"] = "`reqdata.amount`",
                    ["autoSendReceipt"] = "`reqdata.auto_send_receipt`",
                    ["baseOriginUrl"] = "`reqdata.base_origin_url`",
                    ["callbackUrl"] = "`reqdata.callback_url`",
                    ["cardAuthorizeOnly"] = "`reqdata.card_authorize_only`",
                    ["cardCreateToken"] = "`reqdata.card_create_token`",
                    ["cardCreateTokenMode"] = "`reqdata.card_create_token_mode`",
                    ["cardIgnoreCVN"] = "`reqdata.card_ignore_cvn`",
                    ["cardProcessorMerchantID"] = "`reqdata.card_processor_merchant_id`",
                    ["currency"] = "`reqdata.currency`",
                    ["customFields"] = "`reqdata.custom_field`",
                    ["customerEmailAddress"] = "`reqdata.customer_email_address`",
                    ["customerID"] = "`reqdata.customer_id`",
                    ["description"] = "`reqdata.description`",
                    ["dueDate"] = "`reqdata.due_date`",
                    ["failureCallbackUrl"] = "`reqdata.failure_callback_url`",
                    ["ignoreAddressVerification"] = "`reqdata.ignore_address_verification`",
                    ["lightningInvoice"] = "`reqdata.lightning_invoice`",
                    ["lightningInvoiceExpiresAt"] = "`reqdata.lightning_invoice_expires_at`",
                    ["notificationEmailAddresses"] = "`reqdata.notification_email_address`",
                    ["orderID"] = "`reqdata.order_id`",
                    ["partialPaymentSteps"] = "`reqdata.partial_payment_step`",
                    ["paymentMethods"] = "`reqdata.payment_method`",
                    ["pispAccountID"] = "`reqdata.pisp_account_id`",
                    ["shippingAddressCity"] = "`reqdata.shipping_address_city`",
                    ["shippingAddressCountryCode"] = "`reqdata.shipping_address_country_code`",
                    ["shippingAddressCounty"] = "`reqdata.shipping_address_county`",
                    ["shippingAddressLine1"] = "`reqdata.shipping_address_line1`",
                    ["shippingAddressLine2"] = "`reqdata.shipping_address_line2`",
                    ["shippingAddressPostCode"] = "`reqdata.shipping_address_post_code`",
                    ["shippingEmail"] = "`reqdata.shipping_email`",
                    ["shippingFirstName"] = "`reqdata.shipping_first_name`",
                    ["shippingLastName"] = "`reqdata.shipping_last_name`",
                    ["shippingPhone"] = "`reqdata.shipping_phone`",
                    ["successWebHookUrl"] = "`reqdata.success_web_hook_url`",
                    ["tagIds"] = "`reqdata.tag_id`",
                    ["title"] = "`reqdata.title`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "getbyorderid",
            },
          },
        },
      },
      ["payment_account"] = {
        ["fields"] = {
          {
            ["name"] = "accountName",
            ["short"] = "Name for the account",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "accountSupplierName",
            ["short"] = "The payment account supplier name.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "availableBalance",
            ["short"] = "The current available balance of the account.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "availableBalanceMinorUnits",
            ["short"] = "The available balance expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "balance",
            ["short"] = "Balance of the account.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "balanceMinorUnits",
            ["short"] = "Balance of the account expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "bankName",
            ["short"] = "The bank name for external accounts",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "consentID",
            ["short"] = "The ID of the consent used to connect the external account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "createdBy",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "createdByDisplayName",
            ["short"] = "Either the name of the user, merchant token or api key that created the account",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currency",
            ["short"] = "Currency of the account in ISO 4217 format",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "defaultPaymentRail",
            ["short"] = "Indicates the default payment rail for this account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "displayName",
            ["short"] = "Gets a unique display name for the payment account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "expiryDate",
            ["short"] = "The date that the external account will expire",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "externalAccountIcon",
            ["short"] = "The Icon for external accounts",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique id for the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "identifier",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "inserted",
            ["short"] = "Timestamp when the account was created.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isArchived",
            ["short"] = "Indicates whether the account is archived.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isConnectedAccount",
            ["short"] = "Indicates if the payment account is an externally connected account.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isDefault",
            ["short"] = "Is the default account",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isTrustAccount",
            ["short"] = "Indicates if the payment account is a trust account.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isVirtual",
            ["short"] = "True if the account is a virtual account.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lastTransaction",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "lastUpdated",
            ["short"] = "Timestamp when the account was last updated.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["short"] = "The ID of the merchant that owns the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantName",
            ["short"] = "The name of the merchant that owns the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "physicalAccountID",
            ["short"] = "For virtual accounts this is the ID of the physical account that the virtual account is linked to.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "rules",
            ["short"] = "The list of rules associated with this account.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "submittedPayoutsBalance",
            ["short"] = "Total of the payouts that have been submitted for processing.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "submittedPayoutsBalanceMinorUnits",
            ["short"] = "The balance of the submitted payouts expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "summary",
            ["short"] = "Gets a summary of the payments account's most important properties.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supplierSepaInstantStatus",
            ["short"] = "Indicates the status of the SEPA Instant payment rail for this account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroBankFeedConnectionStatus",
            ["short"] = "States the status of the Xero bank feed connection, if applicable.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroBankFeedLastSyncedAt",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroBankFeedSyncLastFailedAt",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroBankFeedSyncLastFailureReason",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroBankFeedSyncStatus",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroUnsynchronisedTransactionsCount",
            ["short"] = "Indicates the number of unsynchronised transactions with Xero",
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "payment_account",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "connected_account",
                      ["orig"] = "connected_account",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_archived",
                      ["orig"] = "include_archived",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_child_merchant",
                      ["orig"] = "include_child_merchant",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "only_archived",
                      ["orig"] = "only_archived",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "only_connect_account",
                      ["orig"] = "only_connect_account",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/accounts/paged",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "paged",
                },
                ["select"] = {
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["example"] = 1,
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/accounts/{accountID}/virtual",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "{account_id}",
                  "virtual",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "page_number",
                    "page_size",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "account",
            },
          },
        },
      },
      ["payment_account_minimal"] = {
        ["fields"] = {
          {
            ["name"] = "accountName",
            ["short"] = "Name for the account",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "availableBalance",
            ["short"] = "The current available balance of the account.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "balance",
            ["short"] = "Balance of the account.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "balanceMinorUnits",
            ["short"] = "Balance of the account expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "currency",
            ["short"] = "Currency of the account in ISO 4217 format",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique id for the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "identifier",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "isArchived",
            ["short"] = "Is the account archived",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isConnectedAccount",
            ["short"] = "Indicates if the payment account is an externally connected account.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "merchantID",
            ["short"] = "The ID of the merchant that owns the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "submittedPayoutsBalance",
            ["short"] = "Total of the payouts that have been submitted for processing.",
            ["type"] = "`$NUMBER`",
          },
        },
        ["name"] = "payment_account_minimal",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "connected_account",
                      ["orig"] = "connected_account",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_archived",
                      ["orig"] = "include_archived",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "only_archived",
                      ["orig"] = "only_archived",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "only_connect_account",
                      ["orig"] = "only_connect_account",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["example"] = 1,
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = 10,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/accounts/minimal",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "minimal",
                },
                ["select"] = {
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["payment_initiation"] = {
        ["fields"] = {
          {
            ["name"] = "paymentInitiationID",
            ["short"] = "The unique identifier of the payment initiation request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentRequestCallbackUrl",
            ["short"] = "The callback URL that was set when the payment request was created.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentRequestID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "redirectUrl",
            ["short"] = "A redirect URL for the user to authorise the payment initiation request at the ASPSP",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "responseType",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "specificErrorMessage",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "payment_initiation",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/paymentrequests/{id}/pisp",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "pisp",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "paymentrequest_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "paymentrequest",
            },
          },
        },
      },
      ["payment_request"] = {
        ["fields"] = {
          {
            ["name"] = "addresses",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "amount",
            ["short"] = "The amount of money to request.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountPending",
            ["short"] = "Total amount that has been authorised but not settled for this payment request.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountReceived",
            ["short"] = "Total amount received for this payment request.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountRefunded",
            ["short"] = "Total amount refunded for this payment request.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "autoSendReceipt",
            ["short"] = "If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "baseOriginUrl",
            ["short"] = "For card payments the origin of the payment page needs to be set in advance.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "callbackUrl",
            ["short"] = "Once a payment is processed, or a notification of an inbound payment is received, a callback request will be made to this URL.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardAuthorizeOnly",
            ["short"] = "For card payments the default behaviour is to authorise and capture the payment at the same time.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "cardCreateToken",
            ["short"] = "For card payments a payment attempt can be used to create a reusable token for subsequent payments.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "cardCreateTokenMode",
            ["short"] = "This specifies whether user consent will be taken before tokenising card or not.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardIgnoreCVN",
            ["short"] = "If set to true the card payment gateway will be directed to proceed with a payment even if the card verification number check fails.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "cardProcessorMerchantID",
            ["short"] = "Optional field that if specified indicates the processor merchant ID that should be used to process any card payments.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardStripePaymentIntentID",
            ["short"] = "If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent ID.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardStripePaymentIntentSecret",
            ["short"] = "If Stripe is being used as the card payment processor this property is used to hold the Stripe payment intent client secret.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "createdByUser",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "currency",
            ["short"] = "The currency of the request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customFields",
            ["short"] = "A list of custom fields attached to the payment request.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "customerEmailAddress",
            ["short"] = "Optional email address for the customer.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerID",
            ["short"] = "An optional customer identifier for the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerName",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "description",
            ["short"] = "An optional description for the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "destinationAccount",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "directDebitPayment",
            ["short"] = "Contains information about a Direct Debit payment attempt for a payment request.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "doSimulateSettlementFailure",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "dueDate",
            ["short"] = "The due date for the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "errorDescription",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "events",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "failedPaymentRequests",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "failureCallbackUrl",
            ["short"] = "Optional callback URL for payment failures that can occur when the payer is redirected away from the payment page.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fieldDisplaySettings",
            ["short"] = "A list of field display settings that control which fields are displayed to the payer.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "formattedAmount",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "hostedPayCheckoutUrl",
            ["short"] = "This is a convenience link generated for payment requests whose merchants are using hosted payment pages.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ignoreAddressVerification",
            ["short"] = "If set to true the card payment gateway will be directed to proceed with a payment even if the address verification checks fails.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "inserted",
            ["short"] = "The timestamp the payment request was created at.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "insertedSortable",
            ["short"] = "The Inserted timestamp output as a sortable string https://docs.microsoft.com/en-us/dotnet/standard/base-types/standard-date-and-time-format-strings#UniversalSortable Format also supported natively by Javascript https://tc39.es/ecma262/#se…",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "institution",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isArchived",
            ["short"] = "Indicates whether the payment request is archived.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "jwk",
            ["short"] = "The jwk containing the public key used to verify the signature of the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lastUpdated",
            ["short"] = "The timestamp the payment request was last updated at.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lightningInvoice",
            ["short"] = "Bitcoin Lightning invoice for the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lightningInvoiceExpiresAt",
            ["short"] = "Date and time of expiration of the lightning invoice.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantDirectDebitMandateID",
            ["short"] = "Optional ID of the direct debit mandate associated with this payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantTokenDescription",
            ["short"] = "Description of the merchant token in case the Payment request was created using a merchant token.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "notificationEmailAddresses",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "notificationRoleIDs",
            ["short"] = "A list of roles whose members will receive notifications about this payment request.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "orderID",
            ["short"] = "An optional order ID for the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "partialPaymentMethod",
            ["short"] = "The approach to use, or not, for accepting partial payments.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "partialPaymentSteps",
            ["short"] = "An optional comma separated list of partial payment amounts.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentAttempts",
            ["short"] = "The payment attempts made against this payment request.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "paymentInitiationID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentMethods",
            ["short"] = "The payment methods that the payment request supports.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "paymentProcessor",
            ["short"] = "If the card payment option is enabled this field indicates which card processor the merchant is set up to use.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentRequests",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "payrunID",
            ["short"] = "The ID of a payrun that needs an account top up.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "pispAccountID",
            ["short"] = "The payment account ID to use to receive payment initiation payments.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "priorityBankID",
            ["short"] = "The ID of the bank that is set as the priority bank for display on pay element.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "result",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sandboxSettleDelayInSeconds",
            ["short"] = "Sandbox only.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "shippingAddress",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "status",
            ["short"] = "The current status of the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "successWebHookUrl",
            ["short"] = "If a payment event results in the payment request being classified as fully paid this success webhook URL will be invoked.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tags",
            ["short"] = "An optional list of descriptive tags attached to the payment request.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "title",
            ["short"] = "A generic field to contain any additional data that the merchant wishes to store against the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tokenisedCards",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "transactions",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "useHostedPaymentPage",
            ["short"] = "If set to true, and the merchant is configured for hosted payment pages, the base and callback URLs will be set to use the hosted payment page.",
            ["type"] = "`$BOOLEAN`",
          },
        },
        ["name"] = "payment_request",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "mandate_id",
                      ["orig"] = "mandate_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "submit_after",
                      ["orig"] = "submit_after",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/paymentrequests/{id}/directdebit",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "directdebit",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "paymentrequest_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "mandate_id",
                    "paymentrequest_id",
                    "submit_after",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/paymentrequests/batchcreate",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "batchcreate",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_archived",
                      ["orig"] = "include_archived",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "max_amount",
                      ["orig"] = "max_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "min_amount",
                      ["orig"] = "min_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = 1,
                      ["kind"] = "query",
                      ["name"] = "page",
                      ["orig"] = "page",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "payment_method",
                      ["orig"] = "payment_method",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = "All",
                      ["kind"] = "query",
                      ["name"] = "status",
                      ["orig"] = "status",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "tag",
                      ["orig"] = "tag",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/paymentrequests",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                },
                ["select"] = {
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_archived",
                      ["orig"] = "include_archived",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "max_amount",
                      ["orig"] = "max_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "min_amount",
                      ["orig"] = "min_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["example"] = 1,
                      ["kind"] = "query",
                      ["name"] = "page",
                      ["orig"] = "page",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "payment_method",
                      ["orig"] = "payment_method",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = "All",
                      ["kind"] = "query",
                      ["name"] = "status",
                      ["orig"] = "status",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "tag",
                      ["orig"] = "tag",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/paymentrequests/export",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "export",
                },
                ["select"] = {
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/paymentrequests/{id}/receipt",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "receipt",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "paymentrequest_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/paymentrequests/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/paymentrequests/{id}/pisp/sandboxcallback",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "pisp",
                  "sandboxcallback",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "paymentrequest_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["amount"] = "`reqdata.amount`",
                    ["doSimulateSettlementFailure"] = "`reqdata.do_simulate_settlement_failure`",
                    ["errorDescription"] = "`reqdata.error_description`",
                    ["institution"] = "`reqdata.institution`",
                    ["paymentInitiationID"] = "`reqdata.payment_initiation_id`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "paymentrequest",
            },
          },
        },
      },
      ["payment_request_event"] = {
        ["fields"] = {
          {
            ["name"] = "amount",
            ["req"] = true,
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "applePayTransactionID",
            ["short"] = "Transaction ID received in Apple pay token.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardAuthorizationResponseID",
            ["short"] = "For a successful card authorization this field will hold the response ID.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardExpiryMonth",
            ["short"] = "For card payment events this field holds the payer's card expiry month.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "cardExpiryYear",
            ["short"] = "For card payment events this field holds the payer's card expiry year.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "cardIssuer",
            ["short"] = "For card payment events this field holds the payer's card issuer.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardIssuerCountry",
            ["short"] = "For card payment events this field holds the payer's card issuer country of origin.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardLastFourDigits",
            ["short"] = "For card payment events this field holds the payer's card last four digits.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardRequestID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardScheme",
            ["short"] = "For card payment events this field holds the scheme of the payer's card, e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardTokenCustomerID",
            ["short"] = "If the option to create a reusable token for card payments was set this field contains the token the merchant can store to use for repeat payments.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardTransactionID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currency",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "directDebitPaymentID",
            ["short"] = "Payment ID issued by the Direct Debit supplier.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "directDebitPaymentReference",
            ["short"] = "Reference string issued by the Direct Debit supplier.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "drirectDebitMandateID",
            ["short"] = "The ID of the mandate that was used wehn requesting payment.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "errorMessage",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "errorReason",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "eventType",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lightningInvoice",
            ["short"] = "For Bitcoin Lightning payments this field holds the invoice presented to the payer.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lightningRHash",
            ["short"] = "For Bitcoin Lightning payments the hash of the invoice presented to the payer.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "originUrl",
            ["short"] = "Optional field that can be set by payment methods, such as pay by bank, that may want to redirect back to the URL that initiated the attempt in the case of a failure condition.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentMethodType",
            ["short"] = "The type of payment method the event relates to, e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentProcessorName",
            ["short"] = "If the event was for a card payment this is the name of the card processor, e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentRequestID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "pispBankStatus",
            ["short"] = "For payment initiation attempts some providers (e.g.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "pispPaymentInitiationID",
            ["short"] = "For a payment initiation this is the ID returned by the service provider initiating the payment for us.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "pispPaymentInstitutionName",
            ["short"] = "For a payment initiation this is the name of the financial institution that is used to initiate and authorise the payment.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "pispPaymentServiceProviderID",
            ["short"] = "For a payment initiation this is the service provider ID selected by the payer, typically the ID for the bank or similar financial institution.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "pispRedirectUrl",
            ["short"] = "For a payment initiation this is the redirect URL returned by the service provider initiating the payment for us.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "reconciledTransactionID",
            ["short"] = "For settlement events (only relevant for non-card payments) this is the payin transaction that the payment request event was reconciled with.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "refundPayoutID",
            ["short"] = "ID of the Payout that was created for refund.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "status",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "walletName",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "payment_request_event",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/paymentrequests/{id}/events",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "events",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "paymentrequest_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "paymentrequest",
            },
          },
        },
      },
      ["payment_request_metric"] = {
        ["fields"] = {},
        ["name"] = "payment_request_metric",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_archived",
                      ["orig"] = "include_archived",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "max_amount",
                      ["orig"] = "max_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "min_amount",
                      ["orig"] = "min_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "payment_method",
                      ["orig"] = "payment_method",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "tag",
                      ["orig"] = "tag",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/paymentrequests/metrics",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "metrics",
                },
                ["select"] = {
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.totalAmountsByCurrency`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["payment_request_minimal"] = {
        ["fields"] = {
          {
            ["name"] = "amount",
            ["short"] = "The amount of money to request.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountPending",
            ["short"] = "The amount of money that was authorised but has not arrived in the account yet.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountReceived",
            ["short"] = "The amount of money that has been received for this payment request.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountRefunded",
            ["short"] = "The amount of money that has been refunded for this payment request.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "callbackUrl",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "cardStripePaymentIntentSecret",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "countryCode",
            ["short"] = "The country code associated with the payment.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currency",
            ["short"] = "The currency of the request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customFieldsToDisplay",
            ["short"] = "Custom fields to display to the customer.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "description",
            ["short"] = "An optional description for the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "dueDate",
            ["short"] = "The due date of the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fieldDisplaySettings",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "googlePayMerchantID",
            ["short"] = "Merchant ID from Google Pay",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "jwk",
            ["short"] = "The jwk containing the public key",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantLogoUrlPng",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantLogoUrlSvg",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantName",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantShortName",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "partialPaymentMethod",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentAttempts",
            ["short"] = "The payment attempts for this payment request.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "paymentMethodsList",
            ["short"] = "The payment methods that the payment request supports.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "paymentProcessor",
            ["short"] = "The card processor",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentProcessorKey",
            ["short"] = "The card processors public key",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "pispError",
            ["short"] = "This is the error returned from the bank which is recorded in payment request events.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "priorityBankID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "status",
            ["short"] = "The status of the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "stripeAccountID",
            ["short"] = "Account ID of connected customers in Stripe",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "title",
            ["short"] = "The title of the payment request.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "payment_request_minimal",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/paymentrequests/{id}/minimal",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "minimal",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "paymentrequest_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "paymentrequest",
            },
          },
        },
      },
      ["payment_request_result"] = {
        ["fields"] = {
          {
            ["name"] = "amount",
            ["short"] = "The authorised payment amount.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountPending",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountReceived",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountRefunded",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "currency",
            ["short"] = "The authorised payment currency.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customerID",
            ["short"] = "The customer id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentRequestID",
            ["short"] = "The ID of the payment request the result is for.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payments",
            ["short"] = "The list of payment attempts that have been received for the payment request.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "pispAuthorizations",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "requestedAmount",
            ["short"] = "The full original payment amount requested.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "result",
            ["short"] = "The result of the payment attempt.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "payment_request_result",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "paymentrequest_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/paymentrequests/{id}/result",
                ["parts"] = {
                  "api",
                  "v1",
                  "paymentrequests",
                  "{paymentrequest_id}",
                  "result",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "paymentrequest_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "paymentrequest_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "paymentrequest",
            },
          },
        },
      },
      ["payout"] = {
        ["fields"] = {
          {
            ["name"] = "accountID",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "Gets or Sets Account Id of sending account",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "allowIncomplete",
            ["short"] = "If set to true the payout will get created even if the business validation rules fail.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "amount",
            ["short"] = "Gets or Sets payout amount",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountMinorUnits",
            ["short"] = "The payout amount expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "approvePayoutUrl",
            ["short"] = "This field is used when returning an payout record to a client.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "approverID",
            ["short"] = "Gets the User ID of person that approved the payout.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "authenticationMethods",
            ["short"] = "A list of authentication types allowed to authorise the payout.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "authorisations",
            ["short"] = "A list of the users who have successfully authorised the latest version of the payout and when.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "authorisersCompletedCount",
            ["short"] = "The number of distinct authorisers that have authorised the payout.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "authorisersRequiredCount",
            ["short"] = "The number of authorisers required for this payout.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "batchPayoutID",
            ["short"] = "The ID of the batch the payout is associated with.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "beneficiary",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "beneficiaryID",
            ["short"] = "Optional.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "canAuthorise",
            ["short"] = "True if the payout can be authorised by the user who loaded it.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "canProcess",
            ["short"] = "If set to true indicates the payout has been flagged as safe to process after transaction monitoring.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "canUpdate",
            ["short"] = "True if the payout can be updated by the user who loaded it.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "chargeBearer",
            ["short"] = "Optional field to set who should pay any fees for the payout.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "createdBy",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "createdByEmailAddress",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currency",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "Gets or Sets Currency of payout request",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currentUserID",
            ["short"] = "The ID of the user that requested access to the PayOut record.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "description",
            ["short"] = "Gets or Sets description of payout request",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "destination",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "documents",
            ["short"] = "Documents associated with the payout.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "events",
            ["short"] = "The activity associated with the payout.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "failedPayouts",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "formattedAmount",
            ["short"] = "Currency and formatted amount string.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "formattedFxDestinationAmount",
            ["short"] = "FX destination currency and amount formatted string.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "formattedSchedule",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "formattedScheduleDayOnly",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "formattedSourceAccountAvailableBalance",
            ["short"] = "The available balance of the account the payout is being made from.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fxDestinationAmount",
            ["short"] = "If specified this will be the amount sent to the payee.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "fxDestinationAmountMinorUnits",
            ["short"] = "The payout FxDestinationAmount expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "fxDestinationCurrency",
            ["short"] = "For an FX payout this is the currency to send to the beneficiary.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fxQuoteExpiresAt",
            ["short"] = "If an FX held rate quote ID is being used this is the time the quote expires.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fxQuoteID",
            ["short"] = "Optional.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fxRate",
            ["short"] = "For an FX payout this is the exchange rate to use for the payout.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "fxUseDestinationAmount",
            ["short"] = "For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "hasCurrentUserAuthorised",
            ["short"] = "True if the payout was loaded for a user and that user has already authorised the latest version of the payout.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "id",
            ["short"] = "The ID for the payout.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "invoiceID",
            ["short"] = "Optional field to associate the payout with the invoice from an external application such as Xero.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isArchived",
            ["short"] = "Indicates whether the payout is archived.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isFailed",
            ["short"] = "Set to true if a submitted payout subsequently fails.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isSettled",
            ["short"] = "Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isSubmitted",
            ["short"] = "Indicates whether the payout has been submitted for processing.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lastUpdated",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["short"] = "The ID of the merchant that owns the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantTokenDescription",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nonce",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentProcessor",
            ["short"] = "The usptream payment processor for the payout.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentRail",
            ["short"] = "Optional field to indicate the payment rail to use for the payout.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payouts",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "payrunID",
            ["short"] = "The ID of the payrun that this payout is associated with.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payrunName",
            ["short"] = "The name of the payrun that this payout is associated with.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "reason",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "rule",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "scheduleDate",
            ["short"] = "The date the payout should be submitted.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "scheduled",
            ["short"] = "Should this payout be scheduled for a future date?",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "sourceAccountAvailableBalance",
            ["short"] = "The available balance of the account the payout is being made from.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "sourceAccountAvailableBalanceMinorUnits",
            ["short"] = "The available balance of the source account expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "sourceAccountBic",
            ["short"] = "The BIC of the account the payout is being made from.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sourceAccountCurrency",
            ["short"] = "The currency of the source account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sourceAccountIban",
            ["short"] = "The IBAN of the account the payout is being made from.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sourceAccountIdentifier",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sourceAccountName",
            ["short"] = "The name of the account the payout is being made from.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sourceAccountNumber",
            ["short"] = "The account number of the account the payout is being made from.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sourceAccountSortcode",
            ["short"] = "The sort code of the account the payout is being made from.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "status",
            ["short"] = "Gets or Sets the status of payout request",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tagIds",
            ["short"] = "An optional list of tag ids to add to the payout.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "tags",
            ["short"] = "An optional list of descriptive tags attached to the payout.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "theirReference",
            ["short"] = "Gets or Sets destination reference ID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "topupPayrunID",
            ["short"] = "The ID of a payrun that needs an account top up.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "transactedAmount",
            ["short"] = "The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "transactedFxAmount",
            ["short"] = "The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "transactedFxRate",
            ["short"] = "The actual FX rate applied during settlement, as recorded on the associated transaction.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "type",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "Gets or Sets payout type",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "userID",
            ["short"] = "Gets or Sets User ID of who created the payout request",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "yourReference",
            ["short"] = "Gets or Sets your reference ID",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "payout",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/payouts/batch/submit/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "batch",
                  "submit",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/payouts/submit/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "submit",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/payouts",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["accountID"] = "`reqdata.account_id`",
                    ["allowIncomplete"] = "`reqdata.allow_incomplete`",
                    ["amount"] = "`reqdata.amount`",
                    ["batchPayoutID"] = "`reqdata.batch_payout_id`",
                    ["beneficiaryID"] = "`reqdata.beneficiary_id`",
                    ["chargeBearer"] = "`reqdata.charge_bearer`",
                    ["currency"] = "`reqdata.currency`",
                    ["description"] = "`reqdata.description`",
                    ["destination"] = "`reqdata.destination`",
                    ["documents"] = "`reqdata.document`",
                    ["fxDestinationAmount"] = "`reqdata.fx_destination_amount`",
                    ["fxDestinationCurrency"] = "`reqdata.fx_destination_currency`",
                    ["fxQuoteID"] = "`reqdata.fx_quote_id`",
                    ["fxUseDestinationAmount"] = "`reqdata.fx_use_destination_amount`",
                    ["invoiceID"] = "`reqdata.invoice_id`",
                    ["paymentRail"] = "`reqdata.payment_rail`",
                    ["scheduleDate"] = "`reqdata.schedule_date`",
                    ["scheduled"] = "`reqdata.scheduled`",
                    ["tagIds"] = "`reqdata.tag_id`",
                    ["tags"] = "`reqdata.tag`",
                    ["theirReference"] = "`reqdata.their_reference`",
                    ["topupPayrunID"] = "`reqdata.topup_payrun_id`",
                    ["type"] = "`reqdata.type`",
                    ["yourReference"] = "`reqdata.your_reference`",
                  },
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/payouts/batchcreate",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "batchcreate",
                },
                ["select"] = {
                  ["$action"] = "batchcreate",
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/payouts/send",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "send",
                },
                ["select"] = {
                  ["$action"] = "send",
                },
                ["transform"] = {
                  ["req"] = {
                    ["accountID"] = "`reqdata.account_id`",
                    ["allowIncomplete"] = "`reqdata.allow_incomplete`",
                    ["amount"] = "`reqdata.amount`",
                    ["batchPayoutID"] = "`reqdata.batch_payout_id`",
                    ["beneficiaryID"] = "`reqdata.beneficiary_id`",
                    ["chargeBearer"] = "`reqdata.charge_bearer`",
                    ["currency"] = "`reqdata.currency`",
                    ["description"] = "`reqdata.description`",
                    ["destination"] = "`reqdata.destination`",
                    ["documents"] = "`reqdata.document`",
                    ["fxDestinationAmount"] = "`reqdata.fx_destination_amount`",
                    ["fxDestinationCurrency"] = "`reqdata.fx_destination_currency`",
                    ["fxQuoteID"] = "`reqdata.fx_quote_id`",
                    ["fxUseDestinationAmount"] = "`reqdata.fx_use_destination_amount`",
                    ["invoiceID"] = "`reqdata.invoice_id`",
                    ["paymentRail"] = "`reqdata.payment_rail`",
                    ["scheduleDate"] = "`reqdata.schedule_date`",
                    ["scheduled"] = "`reqdata.scheduled`",
                    ["tagIds"] = "`reqdata.tag_id`",
                    ["tags"] = "`reqdata.tag`",
                    ["theirReference"] = "`reqdata.their_reference`",
                    ["topupPayrunID"] = "`reqdata.topup_payrun_id`",
                    ["type"] = "`reqdata.type`",
                    ["yourReference"] = "`reqdata.your_reference`",
                  },
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/payouts/sendbeneficiary",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "sendbeneficiary",
                },
                ["select"] = {
                  ["$action"] = "sendbeneficiary",
                },
                ["transform"] = {
                  ["req"] = {
                    ["accountID"] = "`reqdata.account_id`",
                    ["allowIncomplete"] = "`reqdata.allow_incomplete`",
                    ["amount"] = "`reqdata.amount`",
                    ["batchPayoutID"] = "`reqdata.batch_payout_id`",
                    ["beneficiaryID"] = "`reqdata.beneficiary_id`",
                    ["chargeBearer"] = "`reqdata.charge_bearer`",
                    ["currency"] = "`reqdata.currency`",
                    ["description"] = "`reqdata.description`",
                    ["destination"] = "`reqdata.destination`",
                    ["documents"] = "`reqdata.document`",
                    ["fxDestinationAmount"] = "`reqdata.fx_destination_amount`",
                    ["fxDestinationCurrency"] = "`reqdata.fx_destination_currency`",
                    ["fxQuoteID"] = "`reqdata.fx_quote_id`",
                    ["fxUseDestinationAmount"] = "`reqdata.fx_use_destination_amount`",
                    ["invoiceID"] = "`reqdata.invoice_id`",
                    ["paymentRail"] = "`reqdata.payment_rail`",
                    ["scheduleDate"] = "`reqdata.schedule_date`",
                    ["scheduled"] = "`reqdata.scheduled`",
                    ["tagIds"] = "`reqdata.tag_id`",
                    ["tags"] = "`reqdata.tag`",
                    ["theirReference"] = "`reqdata.their_reference`",
                    ["topupPayrunID"] = "`reqdata.topup_payrun_id`",
                    ["type"] = "`reqdata.type`",
                    ["yourReference"] = "`reqdata.your_reference`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_archived",
                      ["orig"] = "include_archived",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "max_amount",
                      ["orig"] = "max_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "min_amount",
                      ["orig"] = "min_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "status",
                      ["orig"] = "status",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "tag",
                      ["orig"] = "tag",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/payouts",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                },
                ["select"] = {
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "max_amount",
                      ["orig"] = "max_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "min_amount",
                      ["orig"] = "min_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "status",
                      ["orig"] = "status",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "tag",
                      ["orig"] = "tag",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/accounts/{accountID}/payouts",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "{account_id}",
                  "payouts",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "max_amount",
                      ["orig"] = "max_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "min_amount",
                      ["orig"] = "min_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "status",
                      ["orig"] = "status",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "tag",
                      ["orig"] = "tag",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/payouts",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "payouts",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_archived",
                      ["orig"] = "include_archived",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "max_amount",
                      ["orig"] = "max_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "min_amount",
                      ["orig"] = "min_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "status",
                      ["orig"] = "status",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "tag",
                      ["orig"] = "tag",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/payouts/export",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "export",
                },
                ["select"] = {
                  ["$action"] = "export",
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "amount",
                      ["orig"] = "amount",
                      ["reqd"] = true,
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "destination",
                      ["orig"] = "destination",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "source",
                      ["orig"] = "source",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/payouts/fxquote/{source}/{destination}/{amount}",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "fxquote",
                  "{source}",
                  "{destination}",
                  "{amount}",
                },
                ["select"] = {
                  ["exist"] = {
                    "amount",
                    "destination",
                    "source",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/payouts/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/payouts/{id}/proof",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "{id}",
                  "proof",
                },
                ["select"] = {
                  ["$action"] = "proof",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/payouts/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/payouts/batchdelete",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "batchdelete",
                },
                ["select"] = {
                  ["$action"] = "batchdelete",
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/payouts/cancel/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "cancel",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/payouts/reject/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "reject",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["reason"] = "`reqdata.reason`",
                  },
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/payouts/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["accountID"] = "`reqdata.account_id`",
                    ["allowIncomplete"] = "`reqdata.allow_incomplete`",
                    ["amount"] = "`reqdata.amount`",
                    ["chargeBearer"] = "`reqdata.charge_bearer`",
                    ["currency"] = "`reqdata.currency`",
                    ["description"] = "`reqdata.description`",
                    ["destination"] = "`reqdata.destination`",
                    ["fxDestinationAmount"] = "`reqdata.fx_destination_amount`",
                    ["fxDestinationCurrency"] = "`reqdata.fx_destination_currency`",
                    ["fxQuoteID"] = "`reqdata.fx_quote_id`",
                    ["fxUseDestinationAmount"] = "`reqdata.fx_use_destination_amount`",
                    ["paymentRail"] = "`reqdata.payment_rail`",
                    ["scheduleDate"] = "`reqdata.schedule_date`",
                    ["scheduled"] = "`reqdata.scheduled`",
                    ["tagIds"] = "`reqdata.tag_id`",
                    ["tags"] = "`reqdata.tag`",
                    ["theirReference"] = "`reqdata.their_reference`",
                    ["type"] = "`reqdata.type`",
                    ["yourReference"] = "`reqdata.your_reference`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "account",
            },
            {
              "merchant",
            },
            {
              "fxquote",
            },
          },
        },
      },
      ["payout_keyset_page"] = {
        ["fields"] = {
          {
            ["name"] = "accountID",
            ["short"] = "Gets or Sets Account Id of sending account",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "amount",
            ["short"] = "Gets or Sets payout amount",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountMinorUnits",
            ["short"] = "The payout amount expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "approvePayoutUrl",
            ["short"] = "This field is used when returning an payout record to a client.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "approverID",
            ["short"] = "Gets the User ID of person that approved the payout.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "authenticationMethods",
            ["short"] = "A list of authentication types allowed to authorise the payout.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "authorisations",
            ["short"] = "A list of the users who have successfully authorised the latest version of the payout and when.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "authorisersCompletedCount",
            ["short"] = "The number of distinct authorisers that have authorised the payout.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "authorisersRequiredCount",
            ["short"] = "The number of authorisers required for this payout.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "batchPayoutID",
            ["short"] = "The ID of the batch the payout is associated with.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "beneficiary",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "canAuthorise",
            ["short"] = "True if the payout can be authorised by the user who loaded it.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "canProcess",
            ["short"] = "If set to true indicates the payout has been flagged as safe to process after transaction monitoring.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "canUpdate",
            ["short"] = "True if the payout can be updated by the user who loaded it.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "chargeBearer",
            ["short"] = "Optional field to set who should pay any fees for the payout.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "createdBy",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "createdByEmailAddress",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currency",
            ["short"] = "Gets or Sets Currency of payout request",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currentUserID",
            ["short"] = "The ID of the user that requested access to the PayOut record.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "description",
            ["short"] = "Gets or Sets description of payout request",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "destination",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "documents",
            ["short"] = "Documents associated with the payout.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "events",
            ["short"] = "The activity associated with the payout.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "formattedAmount",
            ["short"] = "Currency and formatted amount string.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "formattedFxDestinationAmount",
            ["short"] = "FX destination currency and amount formatted string.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "formattedSchedule",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "formattedScheduleDayOnly",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "formattedSourceAccountAvailableBalance",
            ["short"] = "The available balance of the account the payout is being made from.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fxDestinationAmount",
            ["short"] = "If specified this will be the amount sent to the payee.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "fxDestinationAmountMinorUnits",
            ["short"] = "The payout FxDestinationAmount expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "fxDestinationCurrency",
            ["short"] = "For an FX payout this is the currency to send to the beneficiary.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fxQuoteExpiresAt",
            ["short"] = "If an FX held rate quote ID is being used this is the time the quote expires.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fxQuoteID",
            ["short"] = "Optional.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fxRate",
            ["short"] = "For an FX payout this is the exchange rate to use for the payout.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "fxUseDestinationAmount",
            ["short"] = "For a multi-currency payout this indicates how the Amount and FxDestinationAmount are treated.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "hasCurrentUserAuthorised",
            ["short"] = "True if the payout was loaded for a user and that user has already authorised the latest version of the payout.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "id",
            ["short"] = "The ID for the payout.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "invoiceID",
            ["short"] = "Optional field to associate the payout with the invoice from an external application such as Xero.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isArchived",
            ["short"] = "Indicates whether the payout is archived.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isFailed",
            ["short"] = "Set to true if a submitted payout subsequently fails.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isSettled",
            ["short"] = "Set to true if a payout was successfully processed and the corresponding transaction has been recorded on the ledger.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isSubmitted",
            ["short"] = "Indicates whether the payout has been submitted for processing.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lastUpdated",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["short"] = "The ID of the merchant that owns the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantTokenDescription",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nonce",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentProcessor",
            ["short"] = "The usptream payment processor for the payout.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "paymentRail",
            ["short"] = "Optional field to indicate the payment rail to use for the payout.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payrunID",
            ["short"] = "The ID of the payrun that this payout is associated with.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payrunName",
            ["short"] = "The name of the payrun that this payout is associated with.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "rule",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "scheduleDate",
            ["short"] = "The date the payout should be submitted.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "scheduled",
            ["short"] = "Should this payout be scheduled for a future date?",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "sourceAccountAvailableBalance",
            ["short"] = "The available balance of the account the payout is being made from.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "sourceAccountAvailableBalanceMinorUnits",
            ["short"] = "The available balance of the source account expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "sourceAccountBic",
            ["short"] = "The BIC of the account the payout is being made from.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sourceAccountCurrency",
            ["short"] = "The currency of the source account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sourceAccountIban",
            ["short"] = "The IBAN of the account the payout is being made from.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sourceAccountIdentifier",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "sourceAccountName",
            ["short"] = "The name of the account the payout is being made from.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sourceAccountNumber",
            ["short"] = "The account number of the account the payout is being made from.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sourceAccountSortcode",
            ["short"] = "The sort code of the account the payout is being made from.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "status",
            ["short"] = "Gets or Sets the status of payout request",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "tags",
            ["short"] = "An optional list of descriptive tags attached to the payout.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "theirReference",
            ["short"] = "Gets or Sets destination reference ID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "topupPayrunID",
            ["short"] = "The ID of a payrun that needs an account top up.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "transactedAmount",
            ["short"] = "The actual amount debited from the account in NoFrixion.MoneyMoov.Models.Payout.Currency, as recorded on the settled transaction.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "transactedFxAmount",
            ["short"] = "The actual amount received by the beneficiary in NoFrixion.MoneyMoov.Models.Payout.FxDestinationCurrency, as recorded on the settled transaction.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "transactedFxRate",
            ["short"] = "The actual FX rate applied during settlement, as recorded on the associated transaction.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "type",
            ["short"] = "Gets or Sets payout type",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "userID",
            ["short"] = "Gets or Sets User ID of who created the payout request",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "yourReference",
            ["short"] = "Gets or Sets your reference ID",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "payout_keyset_page",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date_utc",
                      ["orig"] = "from_date_utc",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/accounts/{accountID}/payouts/failed",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "{account_id}",
                  "payouts",
                  "failed",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "from_date_utc",
                    "page_size",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date_utc",
                      ["orig"] = "from_date_utc",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/payouts/failed",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "payouts",
                  "failed",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "from_date_utc",
                    "merchant_id",
                    "page_size",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date_utc",
                      ["orig"] = "from_date_utc",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/payouts/{merchantID}/failed",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "{merchant_id}",
                  "failed",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "from_date_utc",
                    "merchant_id",
                    "page_size",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "account",
            },
            {
              "merchant",
            },
            {
              "payout",
            },
          },
        },
      },
      ["payout_metric"] = {
        ["fields"] = {},
        ["name"] = "payout_metric",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "currency",
                      ["orig"] = "currency",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_archived",
                      ["orig"] = "include_archived",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "max_amount",
                      ["orig"] = "max_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "min_amount",
                      ["orig"] = "min_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "tag",
                      ["orig"] = "tag",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/payouts/metrics",
                ["parts"] = {
                  "api",
                  "v1",
                  "payouts",
                  "metrics",
                },
                ["select"] = {
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.totalAmountsByCurrency`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["payrun"] = {
        ["fields"] = {
          {
            ["name"] = "authorisationDate",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "authorisations",
            ["short"] = "A list of the users who have successfully authorised the latest version of the payrun and when.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "authorisersCompletedCount",
            ["short"] = "The number of distinct authorisers that have authorised the payrun.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "authorisersRequiredCount",
            ["short"] = "The number of authorisers required for this payrun.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "batchPayoutID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "canAuthorise",
            ["short"] = "True if the payrun can be authorised by the user who loaded it.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "canDelete",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "canEdit",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "events",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "hasCurrentUserAuthorised",
            ["short"] = "True if the payrun was loaded for a user and that user has already authorised the latest version of the payrun.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "invoices",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "invoicesMinimal",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "isArchived",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lastUpdated",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lastUpdatedBy",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "merchantID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "name",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nonce",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "notes",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payments",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "payouts",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "payoutsCount",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "reason",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "scheduleDate",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "scheduledDate",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sourceAccounts",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "status",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "totalEur",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "totalGbp",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "totalUsd",
            ["type"] = "`$NUMBER`",
          },
        },
        ["name"] = "payrun",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/payruns/{id}/request-authorisation",
                ["parts"] = {
                  "api",
                  "v1",
                  "payruns",
                  "{id}",
                  "request-authorisation",
                },
                ["select"] = {
                  ["$action"] = "request_authorisation",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["id"] = "`reqdata.id`",
                    ["notes"] = "`reqdata.note`",
                    ["scheduledDate"] = "`reqdata.scheduled_date`",
                  },
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/payruns/{id}/submit",
                ["parts"] = {
                  "api",
                  "v1",
                  "payruns",
                  "{id}",
                  "submit",
                },
                ["select"] = {
                  ["$action"] = "submit",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["scheduledDate"] = "`reqdata.scheduled_date`",
                  },
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/payruns/{merchantID}",
                ["parts"] = {
                  "api",
                  "v1",
                  "payruns",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["invoices"] = "`reqdata.invoice`",
                    ["name"] = "`reqdata.name`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "only_archived",
                      ["orig"] = "only_archived",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "status",
                      ["orig"] = "status",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/payruns",
                ["parts"] = {
                  "api",
                  "v1",
                  "payruns",
                },
                ["select"] = {
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/payruns/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "payruns",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/payruns/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "payruns",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/payruns/{id}/archive",
                ["parts"] = {
                  "api",
                  "v1",
                  "payruns",
                  "{id}",
                  "archive",
                },
                ["select"] = {
                  ["$action"] = "archive",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/payruns/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "payruns",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["id"] = "`reqdata.id`",
                    ["invoices"] = "`reqdata.invoice`",
                    ["name"] = "`reqdata.name`",
                    ["scheduledDate"] = "`reqdata.scheduled_date`",
                    ["sourceAccounts"] = "`reqdata.source_account`",
                  },
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/payruns/{id}/cancel",
                ["parts"] = {
                  "api",
                  "v1",
                  "payruns",
                  "{id}",
                  "cancel",
                },
                ["select"] = {
                  ["$action"] = "cancel",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/payruns/{id}/reject",
                ["parts"] = {
                  "api",
                  "v1",
                  "payruns",
                  "{id}",
                  "reject",
                },
                ["select"] = {
                  ["$action"] = "reject",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["id"] = "`reqdata.id`",
                    ["reason"] = "`reqdata.reason`",
                  },
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/payruns/{id}/unarchive",
                ["parts"] = {
                  "api",
                  "v1",
                  "payruns",
                  "{id}",
                  "unarchive",
                },
                ["select"] = {
                  ["$action"] = "unarchive",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["report"] = {
        ["fields"] = {},
        ["name"] = "report",
        ["op"] = {
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/reports/{id}/initiate",
                ["parts"] = {
                  "api",
                  "v1",
                  "reports",
                  "{id}",
                  "initiate",
                },
                ["select"] = {
                  ["$action"] = "initiate",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["report_result"] = {
        ["fields"] = {
          {
            ["name"] = "contentType",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "contents",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lastCompletedAt",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "reportName",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "reportType",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "statementNumber",
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "report_result",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "statement_number",
                      ["reqd"] = true,
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "report_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/reports/{id}/result/{statementNumber}",
                ["parts"] = {
                  "api",
                  "v1",
                  "reports",
                  "{report_id}",
                  "result",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "report_id",
                    ["statementNumber"] = "id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                    "report_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "report",
            },
          },
        },
      },
      ["role"] = {
        ["fields"] = {
          {
            ["name"] = "failedRoles",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "roles",
            ["type"] = "`$ARRAY`",
          },
        },
        ["name"] = "role",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/merchants/{merchantID}/roles/batchcreate",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "roles",
                  "batchcreate",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["$action"] = "batchcreate",
                  ["exist"] = {
                    "merchant_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "merchant",
            },
          },
        },
      },
      ["rule"] = {
        ["fields"] = {
          {
            ["name"] = "account",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "accountID",
            ["short"] = "The ID of the account the rule will apply to.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "approveUrl",
            ["short"] = "If set this property holds the URL an approver needs to visit in order to complete a strong authentication check in order to approve the rule.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "approverID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "authenticationMethods",
            ["short"] = "A list of authentication types allowed to authorise the payout.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "authorisations",
            ["short"] = "A list of the users who have successfully authorised the latest version of the rule and when.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "authorisersCompletedCount",
            ["short"] = "The number of distinct authorisers that have authorised the rule.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "authorisersRequiredCount",
            ["short"] = "The number of authorisers required for this rule.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "canAuthorise",
            ["short"] = "True if the rule can be authorised by the user who loaded it.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "createdBy",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "description",
            ["short"] = "Arbitrary description for the rule.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "endAt",
            ["short"] = "Optional end time for rule executions.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "hasCurrentUserAuthorised",
            ["short"] = "True if the current user has authorised.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isDisabled",
            ["short"] = "If set to true the rule will be disabled from executing.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lastExecutedAt",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lastRunAtTransactionDate",
            ["short"] = "The most recent transaction date when the rule was last run.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lastUpdated",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["short"] = "The ID of the merchant that owns the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "name",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "A name to succinctly describe the rule.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nonce",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "onApprovedWebHookUrl",
            ["short"] = "Optional URL to receive an HTTP request with the rule details when the rule status changes to approved.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "onExecutionErrorWebHookUrl",
            ["short"] = "Optional URL to receive an HTTP request when a rule execution attempt fails.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "onExecutionSuccessWebHookUrl",
            ["short"] = "Optional URL to receive an HTTP request when a rule execution attempt succeeds.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "startAt",
            ["short"] = "Optional start time for rule executions.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "status",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sweepAction",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$OBJECT`",
              },
            },
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "timeZoneId",
            ["short"] = "If the rule should be executed on a recurring schedule this is the timezone that the CRON expression should be evaluated in.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "triggerCronExpression",
            ["short"] = "If the rule should be executed on a recurring schedule this is the expression that sets the schedule.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "triggerOnPayIn",
            ["short"] = "Set to true if the rule execution should be triggered when the account receives a pay in (credit).",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "userID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "webHookSecret",
            ["short"] = "If set this secret will be used to sign Web Hook requests.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "rule",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/rules",
                ["parts"] = {
                  "api",
                  "v1",
                  "rules",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["accountID"] = "`reqdata.account_id`",
                    ["description"] = "`reqdata.description`",
                    ["endAt"] = "`reqdata.end_at`",
                    ["isDisabled"] = "`reqdata.is_disabled`",
                    ["name"] = "`reqdata.name`",
                    ["onApprovedWebHookUrl"] = "`reqdata.on_approved_web_hook_url`",
                    ["onExecutionErrorWebHookUrl"] = "`reqdata.on_execution_error_web_hook_url`",
                    ["onExecutionSuccessWebHookUrl"] = "`reqdata.on_execution_success_web_hook_url`",
                    ["startAt"] = "`reqdata.start_at`",
                    ["sweepAction"] = "`reqdata.sweep_action`",
                    ["timeZoneId"] = "`reqdata.time_zone_id`",
                    ["triggerCronExpression"] = "`reqdata.trigger_cron_expression`",
                    ["triggerOnPayIn"] = "`reqdata.trigger_on_pay_in`",
                    ["webHookSecret"] = "`reqdata.web_hook_secret`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "archived_only",
                      ["orig"] = "archived_only",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = 1,
                      ["kind"] = "query",
                      ["name"] = "page",
                      ["orig"] = "page",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/rules",
                ["parts"] = {
                  "api",
                  "v1",
                  "rules",
                },
                ["select"] = {
                  ["exist"] = {
                    "archived_only",
                    "merchant_id",
                    "page",
                    "search",
                    "size",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/rules/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "rules",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/rules/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "rules",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/rules/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "rules",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["accountID"] = "`reqdata.account_id`",
                    ["description"] = "`reqdata.description`",
                    ["endAt"] = "`reqdata.end_at`",
                    ["isDisabled"] = "`reqdata.is_disabled`",
                    ["name"] = "`reqdata.name`",
                    ["onApprovedWebHookUrl"] = "`reqdata.on_approved_web_hook_url`",
                    ["onExecutionErrorWebHookUrl"] = "`reqdata.on_execution_error_web_hook_url`",
                    ["onExecutionSuccessWebHookUrl"] = "`reqdata.on_execution_success_web_hook_url`",
                    ["startAt"] = "`reqdata.start_at`",
                    ["sweepAction"] = "`reqdata.sweep_action`",
                    ["timeZoneId"] = "`reqdata.time_zone_id`",
                    ["triggerCronExpression"] = "`reqdata.trigger_cron_expression`",
                    ["triggerOnPayIn"] = "`reqdata.trigger_on_pay_in`",
                    ["webHookSecret"] = "`reqdata.web_hook_secret`",
                  },
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/rules/{id}/disable",
                ["parts"] = {
                  "api",
                  "v1",
                  "rules",
                  "{id}",
                  "disable",
                },
                ["select"] = {
                  ["$action"] = "disable",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["rule_event"] = {
        ["fields"] = {
          {
            ["name"] = "errorMessage",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isAuthoriseToEnable",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "message",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "rawResponse",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ruleEventType",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ruleID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "user",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
        },
        ["name"] = "rule_event",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "event_type",
                      ["orig"] = "event_type",
                      ["type"] = "`$ARRAY`",
                    },
                    {
                      ["example"] = 1,
                      ["kind"] = "query",
                      ["name"] = "page",
                      ["orig"] = "page",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "size",
                      ["orig"] = "size",
                      ["type"] = "`$INTEGER`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/rules/{id}/events",
                ["parts"] = {
                  "api",
                  "v1",
                  "rules",
                  "{id}",
                  "events",
                },
                ["select"] = {
                  ["exist"] = {
                    "event_type",
                    "id",
                    "page",
                    "size",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["tag"] = {
        ["fields"] = {
          {
            ["name"] = "colourHex",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "description",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "name",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "tag",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/merchants/{merchantID}/tags",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "tags",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "merchant_id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["colourHex"] = "`reqdata.colour_hex`",
                    ["description"] = "`reqdata.description`",
                    ["id"] = "`reqdata.id`",
                    ["merchantID"] = "`reqdata.merchant_id`",
                    ["name"] = "`reqdata.name`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/tags",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "tags",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "merchant_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "merchant",
            },
          },
        },
      },
      ["token"] = {
        ["fields"] = {},
        ["name"] = "token",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/tokens/authorise/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "tokens",
                  "authorise",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/tokens/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "tokens",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["transaction"] = {
        ["fields"] = {
          {
            ["name"] = "accountID",
            ["short"] = "The ID of the account the transaction belongs to.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "accountName",
            ["short"] = "The name of the account the transaction belongs to.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "accountSequenceNumber",
            ["short"] = "The sequence number of transaction on a per account basis.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "addressDetails",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "amount",
            ["short"] = "Amount of the transaction.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "amountMinorUnits",
            ["short"] = "Amount of the transaction expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "balance",
            ["short"] = "Balance left on the account after the transaction.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "balanceMinorUnits",
            ["short"] = "Balance on the account expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "bookingDateTime",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "chargeDetails",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "content",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "counterparty",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "counterpartySummary",
            ["short"] = "For pay in (credit) transactions this will contain a descriptive string with the most important fields about the counterparty.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currency",
            ["short"] = "Currency of transaction.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currencyExchange",
            ["short"] = "Provides details on the currency exchange.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "date",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "description",
            ["short"] = "Description of the transaction.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "enrichment",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "fxAmount",
            ["short"] = "For an FX payout this is the amound in the FX currency.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "fxCurrency",
            ["short"] = "For an FX payout this is the currency that was received or that was instructed.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "fxRate",
            ["short"] = "For an FX payout this is the exchange rate between the transaction currency and the FX currency.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "grossAmount",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique ID for the transaction.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inserted",
            ["short"] = "Date when the transaction was inserted into the ledger.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isoBankTransactionCode",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "merchant",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "merchantID",
            ["short"] = "The ID of the merchant that owns the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "pageNumber",
            ["short"] = "Current page number.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "pageSize",
            ["short"] = "Page size",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "payeeDetails",
            ["req"] = true,
            ["short"] = "The Payee object contains details of the beneficiary, person or business.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "payerDetails",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "paymentRequestCustomFields",
            ["short"] = "The custom fields that were attached to the payment request that resulted in this transaction.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "paymentRequestID",
            ["short"] = "For Pay by Bank and Direct Debit transactions this will contain the ID of the payment request.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "payoutID",
            ["short"] = "ID of the payout that resulted in the transaction.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "proprietaryBankTransactionCode",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "rawReference",
            ["short"] = "The raw payment reference details as received from the payment processor.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "reference",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "ruleID",
            ["short"] = "ID of the rule that resulted in the transaction.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "statementReferences",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "status",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supplementaryData",
            ["type"] = "`$ANY`",
          },
          {
            ["name"] = "tags",
            ["short"] = "An optional list of descriptive tags attached to the transaction.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "theirReference",
            ["short"] = "For a pay out the reference that the payer attached for the receiving party.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "totalPages",
            ["short"] = "Total pages",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "totalSize",
            ["short"] = "Total count",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "transactionAmount",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "transactionDate",
            ["short"] = "Date when the transaction occurred.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "transactionInformation",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "transactionMutability",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "type",
            ["short"] = "Type of the transaction.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "valueDateTime",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "virtualIBAN",
            ["short"] = "If set it indicates the payin was to a virtual IBAN.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "yourReference",
            ["short"] = "For a pay in the reference the sending party attached.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "transaction",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/transactions/{id}/tags",
                ["parts"] = {
                  "api",
                  "v1",
                  "transactions",
                  "{id}",
                  "tags",
                },
                ["select"] = {
                  ["$action"] = "tag",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "credit_type",
                      ["orig"] = "credit_type",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "max_amount",
                      ["orig"] = "max_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "min_amount",
                      ["orig"] = "min_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/accounts/{accountID}/transactions",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "{account_id}",
                  "transactions",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "limit",
                      ["orig"] = "limit",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "offset",
                      ["orig"] = "offset",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort_inserted_ascending",
                      ["orig"] = "sort_inserted_ascending",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "transaction_from",
                      ["orig"] = "transaction_from",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "transaction_to",
                      ["orig"] = "transaction_to",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/openbanking/transactions/{id}/{accountID}",
                ["parts"] = {
                  "api",
                  "v1",
                  "openbanking",
                  "transactions",
                  "{id}",
                  "{account_id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "id",
                    "limit",
                    "offset",
                    "sort_inserted_ascending",
                    "transaction_from",
                    "transaction_to",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = false,
                      ["kind"] = "query",
                      ["name"] = "include_child_merchant",
                      ["orig"] = "include_child_merchant",
                      ["type"] = "`$BOOLEAN`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/transactions",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "transactions",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "from_date",
                    "include_child_merchant",
                    "merchant_id",
                    "page_number",
                    "page_size",
                    "to_date",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "credit_type",
                      ["orig"] = "credit_type",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/transactions",
                ["parts"] = {
                  "api",
                  "v1",
                  "transactions",
                },
                ["select"] = {
                  ["exist"] = {
                    "credit_type",
                    "from_date",
                    "page_number",
                    "page_size",
                    "to_date",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "credit_type",
                      ["orig"] = "credit_type",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "max_amount",
                      ["orig"] = "max_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "min_amount",
                      ["orig"] = "min_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/transactions/{accountID}/export",
                ["parts"] = {
                  "api",
                  "v1",
                  "transactions",
                  "{account_id}",
                  "export",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                  },
                },
                ["select"] = {
                  ["$action"] = "export",
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "credit_type",
                      ["orig"] = "credit_type",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "from_date",
                      ["orig"] = "from_date",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "max_amount",
                      ["orig"] = "max_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "min_amount",
                      ["orig"] = "min_amount",
                      ["type"] = "`$NUMBER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "to_date",
                      ["orig"] = "to_date",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/transactions/{accountID}",
                ["parts"] = {
                  "api",
                  "v1",
                  "transactions",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
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
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["example"] = 0,
                      ["kind"] = "param",
                      ["name"] = "sequence_number",
                      ["orig"] = "sequence_number",
                      ["reqd"] = true,
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "transaction_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/transactions/{accountID}/from/{sequenceNumber}",
                ["parts"] = {
                  "api",
                  "v1",
                  "transactions",
                  "{transaction_id}",
                  "from",
                  "{sequence_number}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "transaction_id",
                    ["sequenceNumber"] = "sequence_number",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "page_size",
                    "sequence_number",
                    "transaction_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/accounts/{accountID}/transactions/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "{account_id}",
                  "transactions",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/transactions/detail/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "transactions",
                  "detail",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/transactions/{id}/proof",
                ["parts"] = {
                  "api",
                  "v1",
                  "transactions",
                  "{id}",
                  "proof",
                },
                ["select"] = {
                  ["$action"] = "proof",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "tag_id",
                      ["orig"] = "tag_id",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/transactions/{id}/tag",
                ["parts"] = {
                  "api",
                  "v1",
                  "transactions",
                  "{id}",
                  "tag",
                },
                ["select"] = {
                  ["$action"] = "tag",
                  ["exist"] = {
                    "id",
                    "tag_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "account",
            },
            {
              "merchant",
            },
            {
              "transaction",
            },
            {
              "transaction",
              "from",
            },
          },
        },
      },
      ["user"] = {
        ["fields"] = {
          {
            ["name"] = "clientSessionTimeouts",
            ["short"] = "The number of seconds a session for this user should last before expiring.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "emailAddress",
            ["op"] = {
              ["update"] = {
                ["type"] = "`$STRING`",
              },
            },
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "firstName",
            ["op"] = {
              ["update"] = {
                ["type"] = "`$STRING`",
              },
            },
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "lastName",
            ["op"] = {
              ["update"] = {
                ["type"] = "`$STRING`",
              },
            },
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "passkeyAdded",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "permissions",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "profile",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "rolesWithScope",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "twoFactorEnabled",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "userInviteID",
            ["short"] = "Optional ID of the invite that is being accepted so the user can be assigned a role on a new merchant.",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "user",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["example"] = 1,
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/user/{merchantID}/userspaged",
                ["parts"] = {
                  "api",
                  "v1",
                  "user",
                  "{merchant_id}",
                  "userspaged",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["$action"] = "userspaged",
                  ["exist"] = {
                    "merchant_id",
                    "page_number",
                    "page_size",
                    "search",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/users",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "users",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "merchant_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/metadata/whoami",
                ["parts"] = {
                  "api",
                  "v1",
                  "metadata",
                  "whoami",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/metadata/whoamitrustedapp",
                ["parts"] = {
                  "api",
                  "v1",
                  "metadata",
                  "whoamitrustedapp",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/user",
                ["parts"] = {
                  "api",
                  "v1",
                  "user",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/user/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "user",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["emailAddress"] = "`reqdata.email_address`",
                    ["firstName"] = "`reqdata.first_name`",
                    ["lastName"] = "`reqdata.last_name`",
                    ["profile"] = "`reqdata.profile`",
                    ["userInviteID"] = "`reqdata.user_invite_id`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "merchant",
            },
            {
              "user",
            },
          },
        },
      },
      ["user_invite"] = {
        ["fields"] = {
          {
            ["name"] = "authorisationStatus",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "failedUserInvites",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "initialRoleID",
            ["short"] = "The role ID to automatically assign to the merchant’s very first user.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inviteeEmailAddress",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "Email address of the user being invited.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inviteeFirstName",
            ["short"] = "First Name of the user being invited.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inviteeLastName",
            ["short"] = "Last Name of the user being invited.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inviterEmailAddress",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inviterFirstName",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "inviterLastName",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isAuthorised",
            ["short"] = "Will be set to true once the invite has met the authorisation requirements.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isInviteeRegistered",
            ["short"] = "If true, indicates the invitee's email address corresponds to an existing MoneyMoov user.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lastInvited",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["short"] = "ID of the merchant the user is being invited to.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantName",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "message",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "registrationUrl",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "sendInviteEmail",
            ["short"] = "If set to true an email will be sent to the invitee with instructions on how to accept the invite.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "status",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "user",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "userID",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "userInvites",
            ["type"] = "`$ARRAY`",
          },
        },
        ["name"] = "user_invite",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/userinvites/authorise/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "userinvites",
                  "authorise",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/userinvites",
                ["parts"] = {
                  "api",
                  "v1",
                  "userinvites",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["initialRoleID"] = "`reqdata.initial_role_id`",
                    ["inviteeEmailAddress"] = "`reqdata.invitee_email_address`",
                    ["inviteeFirstName"] = "`reqdata.invitee_first_name`",
                    ["inviteeLastName"] = "`reqdata.invitee_last_name`",
                    ["merchantID"] = "`reqdata.merchant_id`",
                    ["sendInviteEmail"] = "`reqdata.send_invite_email`",
                  },
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/userinvites/batchcreate",
                ["parts"] = {
                  "api",
                  "v1",
                  "userinvites",
                  "batchcreate",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                  ["query"] = {
                    {
                      ["example"] = 1,
                      ["kind"] = "query",
                      ["name"] = "page_number",
                      ["orig"] = "page_number",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["example"] = 20,
                      ["kind"] = "query",
                      ["name"] = "page_size",
                      ["orig"] = "page_size",
                      ["type"] = "`$INTEGER`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "search",
                      ["orig"] = "search",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "sort",
                      ["orig"] = "sort",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/userinvitespaged",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "userinvitespaged",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "merchant_id",
                    "page_number",
                    "page_size",
                    "search",
                    "sort",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.content`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/userinvites/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "userinvites",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "userinvite_id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/userinvites/{id}/details",
                ["parts"] = {
                  "api",
                  "v1",
                  "userinvites",
                  "{userinvite_id}",
                  "details",
                },
                ["rename"] = {
                  ["param"] = {
                    ["id"] = "userinvite_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "userinvite_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/userinvites/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "userinvites",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/userinvites/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "userinvites",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "merchant",
            },
            {
              "userinvite",
            },
          },
        },
      },
      ["virtual"] = {
        ["fields"] = {
          {
            ["name"] = "accountName",
            ["short"] = "Name for the account",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "accountSupplierName",
            ["short"] = "The payment account supplier name.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "availableBalance",
            ["short"] = "The current available balance of the account.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "availableBalanceMinorUnits",
            ["short"] = "The available balance expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "balance",
            ["short"] = "Balance of the account.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "balanceMinorUnits",
            ["short"] = "Balance of the account expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "bankName",
            ["short"] = "The bank name for external accounts",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "consentID",
            ["short"] = "The ID of the consent used to connect the external account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "createdBy",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "createdByDisplayName",
            ["short"] = "Either the name of the user, merchant token or api key that created the account",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "currency",
            ["short"] = "Currency of the account in ISO 4217 format",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "defaultPaymentRail",
            ["short"] = "Indicates the default payment rail for this account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "displayName",
            ["short"] = "Gets a unique display name for the payment account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "expiryDate",
            ["short"] = "The date that the external account will expire",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "externalAccountIcon",
            ["short"] = "The Icon for external accounts",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "Unique id for the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "identifier",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "inserted",
            ["short"] = "Timestamp when the account was created.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isArchived",
            ["short"] = "Indicates whether the account is archived.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isConnectedAccount",
            ["short"] = "Indicates if the payment account is an externally connected account.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isDefault",
            ["short"] = "Is the default account",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isTrustAccount",
            ["short"] = "Indicates if the payment account is a trust account.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "isVirtual",
            ["short"] = "True if the account is a virtual account.",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "lastTransaction",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "lastUpdated",
            ["short"] = "Timestamp when the account was last updated.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantID",
            ["short"] = "The ID of the merchant that owns the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "merchantName",
            ["short"] = "The name of the merchant that owns the account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "name",
            ["req"] = true,
            ["short"] = "The name of the virtual account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "physicalAccountID",
            ["short"] = "For virtual accounts this is the ID of the physical account that the virtual account is linked to.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "rules",
            ["short"] = "The list of rules associated with this account.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "submittedPayoutsBalance",
            ["short"] = "Total of the payouts that have been submitted for processing.",
            ["type"] = "`$NUMBER`",
          },
          {
            ["name"] = "submittedPayoutsBalanceMinorUnits",
            ["short"] = "The balance of the submitted payouts expressed in the currency’s minor units (e.g.",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "summary",
            ["short"] = "Gets a summary of the payments account's most important properties.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supplierSepaInstantStatus",
            ["short"] = "Indicates the status of the SEPA Instant payment rail for this account.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroBankFeedConnectionStatus",
            ["short"] = "States the status of the Xero bank feed connection, if applicable.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroBankFeedLastSyncedAt",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroBankFeedSyncLastFailedAt",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroBankFeedSyncLastFailureReason",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroBankFeedSyncStatus",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "xeroUnsynchronisedTransactionsCount",
            ["short"] = "Indicates the number of unsynchronised transactions with Xero",
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "virtual",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/accounts/{accountID}/virtual",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "{account_id}",
                  "virtual",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["name"] = "`reqdata.name`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "account_id",
                      ["orig"] = "account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "virtual_account_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/accounts/{accountID}/virtual/{virtualAccountID}",
                ["parts"] = {
                  "api",
                  "v1",
                  "accounts",
                  "{account_id}",
                  "virtual",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["accountID"] = "account_id",
                    ["virtualAccountID"] = "id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "account_id",
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["name"] = "`reqdata.name`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "account",
            },
          },
        },
      },
      ["webhook"] = {
        ["fields"] = {
          {
            ["name"] = "destinationUrl",
            ["short"] = "The destination URL for the webhook.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "emailAddress",
            ["short"] = "The recipient email address(es) for notifications.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "failedNotificationEmailAddress",
            ["short"] = "The email address to which notifications about failed webhook deliveries will be sent.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "isActive",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "merchantID",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
              ["update"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "The ID of the merchant that the webhook is for.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "notificationMethod",
            ["op"] = {
              ["create"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
              ["update"] = {
                ["req"] = true,
                ["type"] = "`$STRING`",
              },
            },
            ["short"] = "The type of notification that will be sent.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "resourceTypes",
            ["short"] = "The resource types that the webhook will be generated for.",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "retry",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "secret",
            ["short"] = "The secret key required to authenticate webhook notifications.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "version",
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "webhook",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/api/v1/webhooks",
                ["parts"] = {
                  "api",
                  "v1",
                  "webhooks",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["destinationUrl"] = "`reqdata.destination_url`",
                    ["emailAddress"] = "`reqdata.email_address`",
                    ["failedNotificationEmailAddress"] = "`reqdata.failed_notification_email_address`",
                    ["id"] = "`reqdata.id`",
                    ["isActive"] = "`reqdata.is_active`",
                    ["merchantID"] = "`reqdata.merchant_id`",
                    ["notificationMethod"] = "`reqdata.notification_method`",
                    ["resourceTypes"] = "`reqdata.resource_type`",
                    ["retry"] = "`reqdata.retry`",
                    ["secret"] = "`reqdata.secret`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/webhooks",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "webhooks",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "merchant_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "param",
                      ["name"] = "merchant_id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/merchants/{merchantID}/webhooks/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "merchants",
                  "{merchant_id}",
                  "webhooks",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "merchant_id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                    "merchant_id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "merchant_id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/api/v1/webhooks/{merchantID}",
                ["parts"] = {
                  "api",
                  "v1",
                  "webhooks",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["merchantID"] = "id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/api/v1/webhooks/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "webhooks",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PUT",
                ["orig"] = "/api/v1/webhooks/{id}",
                ["parts"] = {
                  "api",
                  "v1",
                  "webhooks",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = {
                    ["destinationUrl"] = "`reqdata.destination_url`",
                    ["emailAddress"] = "`reqdata.email_address`",
                    ["failedNotificationEmailAddress"] = "`reqdata.failed_notification_email_address`",
                    ["id"] = "`reqdata.id`",
                    ["isActive"] = "`reqdata.is_active`",
                    ["merchantID"] = "`reqdata.merchant_id`",
                    ["notificationMethod"] = "`reqdata.notification_method`",
                    ["resourceTypes"] = "`reqdata.resource_type`",
                    ["retry"] = "`reqdata.retry`",
                    ["secret"] = "`reqdata.secret`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "merchant",
            },
          },
        },
      },
    },
  }
end


local function make_feature(name)
  local features = require("features")
  local factory = features[name]
  if factory ~= nil then
    return factory()
  end
  return features.base()
end


-- Attach make_feature to the SDK class
local function setup_sdk(SDK)
  SDK._make_feature = make_feature
end


return make_config
