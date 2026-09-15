"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('MerchantDirectDebitMandatePageEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('NOFRIXION_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.NofrixionSDK.test();
        const ent = testsdk.MerchantDirectDebitMandatePage();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE;
        for (const op of ['list']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'merchant_direct_debit_mandate_page.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "format": "date-time", "name": "approvedAt", "req": false, "short": "Date at which the supplier approved this mandate.", "type": "`$STRING`", "index$": 0 }, { "active": true, "name": "currency", "req": false, "short": "Currency of this mandate.", "type": "`$STRING`", "index$": 1 }, { "active": true, "name": "customerAccountNumber", "req": false, "short": "Customer's account number in case of GBP account.", "type": "`$STRING`", "index$": 2 }, { "active": true, "name": "customerCity", "req": false, "short": "Customer's city of residence.", "type": "`$STRING`", "index$": 3 }, { "active": true, "name": "customerCountryCode", "req": false, "short": "Customer's country of residence code.", "type": "`$STRING`", "index$": 4 }, { "active": true, "name": "customerCountryName", "req": false, "short": "Customer's country of residence.", "type": "`$STRING`", "index$": 5 }, { "active": true, "name": "customerEmailAddress", "req": false, "short": "Customer's email address.", "type": "`$STRING`", "index$": 6 }, { "active": true, "name": "customerFirstName", "req": false, "short": "Customer's first name.", "type": "`$STRING`", "index$": 7 }, { "active": true, "name": "customerIban", "req": false, "short": "Customer's IBAN in case of EUR account.", "type": "`$STRING`", "index$": 8 }, { "active": true, "name": "customerLastName", "req": false, "short": "Customer's last name.", "type": "`$STRING`", "index$": 9 }, { "active": true, "name": "customerSortCode", "req": false, "short": "Customer's sort code in case of GBP account.", "type": "`$STRING`", "index$": 10 }, { "active": true, "format": "uuid", "name": "id", "req": false, "short": "Internal ID of the mandate.", "type": "`$STRING`", "index$": 11 }, { "active": true, "format": "date-time", "name": "inserted", "req": false, "short": "The timestamp this mandate was created at.", "type": "`$STRING`", "index$": 12 }, { "active": true, "name": "isRecurring", "req": false, "short": "Whether this mandate is single-use or recurring.", "type": "`$BOOLEAN`", "index$": 13 }, { "active": true, "format": "date-time", "name": "lastUpdated", "req": false, "short": "The timestamp this mandate was last updated at.", "type": "`$STRING`", "index$": 14 }, { "active": true, "format": "uuid", "name": "merchantID", "req": false, "short": "Internal ID of this mandate's merchant.", "type": "`$STRING`", "index$": 15 }, { "active": true, "name": "reference", "req": false, "short": "Reference assigned to this mandate.", "type": "`$STRING`", "index$": 16 }, { "active": true, "name": "status", "req": false, "short": "General status of this mandate.", "type": "`$STRING`", "index$": 17 }, { "active": true, "name": "supplierBankAccountID", "req": false, "short": "ID that the supplier assigned to this mandate's bank account.", "type": "`$STRING`", "index$": 18 }, { "active": true, "name": "supplierCustomerID", "req": false, "short": "ID that the supplier assigned to this mandate's customer.", "type": "`$STRING`", "index$": 19 }, { "active": true, "name": "supplierMandateID", "req": false, "short": "ID that the supplier assigned to this mandate.", "type": "`$STRING`", "index$": 20 }, { "active": true, "name": "supplierName", "req": false, "short": "Name of the supplier used to create this mandate.", "type": "`$STRING`", "index$": 21 }, { "active": true, "name": "supplierStatus", "req": false, "short": "Last status that the supplier reported for this mandate.", "type": "`$STRING`", "index$": 22 }], "id": { "field": "id", "name": "id" }, "name": "merchant_direct_debit_mandate_page", "op": { "list": { "input": "data", "name": "list", "points": [{ "active": true, "args": { "query": [{ "active": true, "kind": "query", "name": "currency", "orig": "currency", "reqd": false, "type": "`$STRING`", "index$": 0 }, { "active": true, "kind": "query", "name": "from_date", "orig": "from_date", "reqd": false, "type": "`$STRING`", "index$": 1 }, { "active": true, "kind": "query", "name": "mandate_i_d", "orig": "mandate_i_d", "reqd": false, "type": "`$ARRAY`", "index$": 2 }, { "active": true, "kind": "query", "name": "max_amount", "orig": "max_amount", "reqd": false, "type": "`$NUMBER`", "index$": 3 }, { "active": true, "kind": "query", "name": "merchant_id", "orig": "merchant_id", "reqd": false, "type": "`$STRING`", "index$": 4 }, { "active": true, "kind": "query", "name": "min_amount", "orig": "min_amount", "reqd": false, "type": "`$NUMBER`", "index$": 5 }, { "active": true, "example": 1, "kind": "query", "name": "page", "orig": "page", "reqd": false, "type": "`$INTEGER`", "index$": 6 }, { "active": true, "kind": "query", "name": "search", "orig": "search", "reqd": false, "type": "`$STRING`", "index$": 7 }, { "active": true, "example": 20, "kind": "query", "name": "size", "orig": "size", "reqd": false, "type": "`$INTEGER`", "index$": 8 }, { "active": true, "kind": "query", "name": "sort", "orig": "sort", "reqd": false, "type": "`$STRING`", "index$": 9 }, { "active": true, "kind": "query", "name": "status", "orig": "status", "reqd": false, "type": "`$STRING`", "index$": 10 }, { "active": true, "kind": "query", "name": "to_date", "orig": "to_date", "reqd": false, "type": "`$STRING`", "index$": 11 }] }, "contract": { "id": "GET /api/v1/mandates", "json": "{\"operationId\":\"GetMandatesPaged\",\"parameters\":[{\"description\":\"ID of the merchant to retrieve mandates for.\",\"in\":\"query\",\"name\":\"merchantID\",\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}},{\"description\":\"The page number from where records are retrieved. Note that the paging starts from page 1.\\r\\n            If a 0 is supplied it will be treated as a 1 and the first page is returned.\",\"in\":\"query\",\"name\":\"page\",\"schema\":{\"default\":1,\"format\":\"int32\",\"type\":\"integer\"}},{\"description\":\"The number of records to be retrieved from a page.\",\"in\":\"query\",\"name\":\"size\",\"schema\":{\"default\":20,\"format\":\"int32\",\"type\":\"integer\"}},{\"description\":\"The date filter to apply to retrieve mandates created after this date.\",\"in\":\"query\",\"name\":\"fromDate\",\"schema\":{\"format\":\"date-time\",\"type\":\"string\"}},{\"description\":\"The date filter to apply to retrieve mandates created up until this date.\",\"in\":\"query\",\"name\":\"toDate\",\"schema\":{\"format\":\"date-time\",\"type\":\"string\"}},{\"description\":\"The status filter to apply to retrieve mandates with this status\",\"in\":\"query\",\"name\":\"status\",\"schema\":{\"enum\":[\"Unknown\",\"Pending\",\"Active\",\"Expired\",\"Failed\"],\"type\":\"string\"}},{\"description\":\"The text filter to apply to retrieve mandates with similar customer, bank account or mandate information.\",\"in\":\"query\",\"name\":\"search\",\"schema\":{\"type\":\"string\"}},{\"description\":\"The currency filter to apply to retrieve mandates with this currency.\",\"in\":\"query\",\"name\":\"currency\",\"schema\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"}},{\"description\":\"The amount filter to apply to retrieve mandates that exceed this amount.\",\"in\":\"query\",\"name\":\"minAmount\",\"schema\":{\"format\":\"double\",\"type\":\"number\"}},{\"description\":\"The amount filter to apply to retrieve mandates that don't exceed this amount.\",\"in\":\"query\",\"name\":\"maxAmount\",\"schema\":{\"format\":\"double\",\"type\":\"number\"}},{\"description\":\"Optional list of mandate IDs to filter by. When supplied, only mandates whose ID is in this list are returned.\",\"in\":\"query\",\"name\":\"mandateIDs\",\"schema\":{\"items\":{\"format\":\"uuid\",\"type\":\"string\"},\"type\":\"array\"}},{\"description\":\"Optional expression to sort the order of the mandates. Example \\\"Amount desc,Inserted asc\\\".\",\"in\":\"query\",\"name\":\"sort\",\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"content\":{\"items\":{\"additionalProperties\":false,\"description\":\"Represents a Direct Debit mandate entity that contains customer,\\r\\n bank account and mandate information.\",\"properties\":{\"approvedAt\":{\"description\":\"Date at which the supplier approved this mandate.\",\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"},\"currency\":{\"description\":\"Currency of this mandate.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"customerAccountNumber\":{\"description\":\"Customer's account number in case of GBP account.\",\"nullable\":true,\"type\":\"string\"},\"customerCity\":{\"description\":\"Customer's city of residence.\",\"nullable\":true,\"type\":\"string\"},\"customerCountryCode\":{\"description\":\"Customer's country of residence code.\",\"nullable\":true,\"type\":\"string\"},\"customerCountryName\":{\"description\":\"Customer's country of residence.\",\"nullable\":true,\"type\":\"string\"},\"customerEmailAddress\":{\"description\":\"Customer's email address.\",\"nullable\":true,\"type\":\"string\"},\"customerFirstName\":{\"description\":\"Customer's first name.\",\"nullable\":true,\"type\":\"string\"},\"customerIban\":{\"description\":\"Customer's IBAN in case of EUR account.\",\"nullable\":true,\"type\":\"string\"},\"customerLastName\":{\"description\":\"Customer's last name.\",\"nullable\":true,\"type\":\"string\"},\"customerSortCode\":{\"description\":\"Customer's sort code in case of GBP account.\",\"nullable\":true,\"type\":\"string\"},\"id\":{\"description\":\"Internal ID of the mandate.\",\"format\":\"uuid\",\"type\":\"string\"},\"inserted\":{\"description\":\"The timestamp this mandate was created at.\",\"format\":\"date-time\",\"type\":\"string\"},\"isRecurring\":{\"description\":\"Whether this mandate is single-use or recurring.\",\"type\":\"boolean\"},\"lastUpdated\":{\"description\":\"The timestamp this mandate was last updated at.\",\"format\":\"date-time\",\"type\":\"string\"},\"merchantID\":{\"description\":\"Internal ID of this mandate's merchant.\",\"format\":\"uuid\",\"type\":\"string\"},\"reference\":{\"description\":\"Reference assigned to this mandate.\",\"nullable\":true,\"type\":\"string\"},\"status\":{\"description\":\"General status of this mandate.\",\"enum\":[\"Unknown\",\"Pending\",\"Active\",\"Expired\",\"Failed\"],\"type\":\"string\"},\"supplierBankAccountID\":{\"description\":\"ID that the supplier assigned to this mandate's\\r\\nbank account.\",\"nullable\":true,\"type\":\"string\"},\"supplierCustomerID\":{\"description\":\"ID that the supplier assigned to this mandate's\\r\\ncustomer.\",\"nullable\":true,\"type\":\"string\"},\"supplierMandateID\":{\"description\":\"ID that the supplier assigned to this mandate.\",\"nullable\":true,\"type\":\"string\"},\"supplierName\":{\"description\":\"Name of the supplier used to create this mandate.\",\"enum\":[\"None\",\"CyberSource\",\"Checkout\",\"Stripe\",\"Modulr\",\"Plaid\",\"Yapily\",\"Nofrixion\",\"Bitcoin\",\"BitcoinTestnet\",\"BankingCircle\",\"BankingCircleAgency\",\"Simulator\",\"Lightning\",\"LightningTestnet\",\"BankingCircleDirectDebit\",\"Technoxander\"],\"nullable\":true,\"type\":\"string\"},\"supplierStatus\":{\"description\":\"Last status that the supplier reported for this mandate.\",\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"},\"pageNumber\":{\"description\":\"Current page number. Its 1 based. i.e firstpage is 1, secondpage is 2\",\"format\":\"int32\",\"type\":\"integer\"},\"pageSize\":{\"description\":\"Page size\",\"format\":\"int32\",\"type\":\"integer\"},\"totalPages\":{\"description\":\"Total pages\",\"format\":\"int64\",\"type\":\"integer\"},\"totalSize\":{\"description\":\"Total count\",\"format\":\"int64\",\"type\":\"integer\"}},\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"content\":{\"items\":{\"additionalProperties\":false,\"description\":\"Represents a Direct Debit mandate entity that contains customer,\\r\\n bank account and mandate information.\",\"properties\":{\"approvedAt\":{\"description\":\"Date at which the supplier approved this mandate.\",\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"},\"currency\":{\"description\":\"Currency of this mandate.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"customerAccountNumber\":{\"description\":\"Customer's account number in case of GBP account.\",\"nullable\":true,\"type\":\"string\"},\"customerCity\":{\"description\":\"Customer's city of residence.\",\"nullable\":true,\"type\":\"string\"},\"customerCountryCode\":{\"description\":\"Customer's country of residence code.\",\"nullable\":true,\"type\":\"string\"},\"customerCountryName\":{\"description\":\"Customer's country of residence.\",\"nullable\":true,\"type\":\"string\"},\"customerEmailAddress\":{\"description\":\"Customer's email address.\",\"nullable\":true,\"type\":\"string\"},\"customerFirstName\":{\"description\":\"Customer's first name.\",\"nullable\":true,\"type\":\"string\"},\"customerIban\":{\"description\":\"Customer's IBAN in case of EUR account.\",\"nullable\":true,\"type\":\"string\"},\"customerLastName\":{\"description\":\"Customer's last name.\",\"nullable\":true,\"type\":\"string\"},\"customerSortCode\":{\"description\":\"Customer's sort code in case of GBP account.\",\"nullable\":true,\"type\":\"string\"},\"id\":{\"description\":\"Internal ID of the mandate.\",\"format\":\"uuid\",\"type\":\"string\"},\"inserted\":{\"description\":\"The timestamp this mandate was created at.\",\"format\":\"date-time\",\"type\":\"string\"},\"isRecurring\":{\"description\":\"Whether this mandate is single-use or recurring.\",\"type\":\"boolean\"},\"lastUpdated\":{\"description\":\"The timestamp this mandate was last updated at.\",\"format\":\"date-time\",\"type\":\"string\"},\"merchantID\":{\"description\":\"Internal ID of this mandate's merchant.\",\"format\":\"uuid\",\"type\":\"string\"},\"reference\":{\"description\":\"Reference assigned to this mandate.\",\"nullable\":true,\"type\":\"string\"},\"status\":{\"description\":\"General status of this mandate.\",\"enum\":[\"Unknown\",\"Pending\",\"Active\",\"Expired\",\"Failed\"],\"type\":\"string\"},\"supplierBankAccountID\":{\"description\":\"ID that the supplier assigned to this mandate's\\r\\nbank account.\",\"nullable\":true,\"type\":\"string\"},\"supplierCustomerID\":{\"description\":\"ID that the supplier assigned to this mandate's\\r\\ncustomer.\",\"nullable\":true,\"type\":\"string\"},\"supplierMandateID\":{\"description\":\"ID that the supplier assigned to this mandate.\",\"nullable\":true,\"type\":\"string\"},\"supplierName\":{\"description\":\"Name of the supplier used to create this mandate.\",\"enum\":[\"None\",\"CyberSource\",\"Checkout\",\"Stripe\",\"Modulr\",\"Plaid\",\"Yapily\",\"Nofrixion\",\"Bitcoin\",\"BitcoinTestnet\",\"BankingCircle\",\"BankingCircleAgency\",\"Simulator\",\"Lightning\",\"LightningTestnet\",\"BankingCircleDirectDebit\",\"Technoxander\"],\"nullable\":true,\"type\":\"string\"},\"supplierStatus\":{\"description\":\"Last status that the supplier reported for this mandate.\",\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"},\"pageNumber\":{\"description\":\"Current page number. Its 1 based. i.e firstpage is 1, secondpage is 2\",\"format\":\"int32\",\"type\":\"integer\"},\"pageSize\":{\"description\":\"Page size\",\"format\":\"int32\",\"type\":\"integer\"},\"totalPages\":{\"description\":\"Total pages\",\"format\":\"int64\",\"type\":\"integer\"},\"totalSize\":{\"description\":\"Total count\",\"format\":\"int64\",\"type\":\"integer\"}},\"type\":\"object\"}},\"text/plain\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"content\":{\"items\":{\"additionalProperties\":false,\"description\":\"Represents a Direct Debit mandate entity that contains customer,\\r\\n bank account and mandate information.\",\"properties\":{\"approvedAt\":{\"description\":\"Date at which the supplier approved this mandate.\",\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"},\"currency\":{\"description\":\"Currency of this mandate.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"customerAccountNumber\":{\"description\":\"Customer's account number in case of GBP account.\",\"nullable\":true,\"type\":\"string\"},\"customerCity\":{\"description\":\"Customer's city of residence.\",\"nullable\":true,\"type\":\"string\"},\"customerCountryCode\":{\"description\":\"Customer's country of residence code.\",\"nullable\":true,\"type\":\"string\"},\"customerCountryName\":{\"description\":\"Customer's country of residence.\",\"nullable\":true,\"type\":\"string\"},\"customerEmailAddress\":{\"description\":\"Customer's email address.\",\"nullable\":true,\"type\":\"string\"},\"customerFirstName\":{\"description\":\"Customer's first name.\",\"nullable\":true,\"type\":\"string\"},\"customerIban\":{\"description\":\"Customer's IBAN in case of EUR account.\",\"nullable\":true,\"type\":\"string\"},\"customerLastName\":{\"description\":\"Customer's last name.\",\"nullable\":true,\"type\":\"string\"},\"customerSortCode\":{\"description\":\"Customer's sort code in case of GBP account.\",\"nullable\":true,\"type\":\"string\"},\"id\":{\"description\":\"Internal ID of the mandate.\",\"format\":\"uuid\",\"type\":\"string\"},\"inserted\":{\"description\":\"The timestamp this mandate was created at.\",\"format\":\"date-time\",\"type\":\"string\"},\"isRecurring\":{\"description\":\"Whether this mandate is single-use or recurring.\",\"type\":\"boolean\"},\"lastUpdated\":{\"description\":\"The timestamp this mandate was last updated at.\",\"format\":\"date-time\",\"type\":\"string\"},\"merchantID\":{\"description\":\"Internal ID of this mandate's merchant.\",\"format\":\"uuid\",\"type\":\"string\"},\"reference\":{\"description\":\"Reference assigned to this mandate.\",\"nullable\":true,\"type\":\"string\"},\"status\":{\"description\":\"General status of this mandate.\",\"enum\":[\"Unknown\",\"Pending\",\"Active\",\"Expired\",\"Failed\"],\"type\":\"string\"},\"supplierBankAccountID\":{\"description\":\"ID that the supplier assigned to this mandate's\\r\\nbank account.\",\"nullable\":true,\"type\":\"string\"},\"supplierCustomerID\":{\"description\":\"ID that the supplier assigned to this mandate's\\r\\ncustomer.\",\"nullable\":true,\"type\":\"string\"},\"supplierMandateID\":{\"description\":\"ID that the supplier assigned to this mandate.\",\"nullable\":true,\"type\":\"string\"},\"supplierName\":{\"description\":\"Name of the supplier used to create this mandate.\",\"enum\":[\"None\",\"CyberSource\",\"Checkout\",\"Stripe\",\"Modulr\",\"Plaid\",\"Yapily\",\"Nofrixion\",\"Bitcoin\",\"BitcoinTestnet\",\"BankingCircle\",\"BankingCircleAgency\",\"Simulator\",\"Lightning\",\"LightningTestnet\",\"BankingCircleDirectDebit\",\"Technoxander\"],\"nullable\":true,\"type\":\"string\"},\"supplierStatus\":{\"description\":\"Last status that the supplier reported for this mandate.\",\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"},\"pageNumber\":{\"description\":\"Current page number. Its 1 based. i.e firstpage is 1, secondpage is 2\",\"format\":\"int32\",\"type\":\"integer\"},\"pageSize\":{\"description\":\"Page size\",\"format\":\"int32\",\"type\":\"integer\"},\"totalPages\":{\"description\":\"Total pages\",\"format\":\"int64\",\"type\":\"integer\"},\"totalSize\":{\"description\":\"Total count\",\"format\":\"int64\",\"type\":\"integer\"}},\"type\":\"object\"}}},\"description\":\"Mandate collection paged result.\"},\"403\":{\"description\":\"Unauthorized error.\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/api/v1/mandates", "segments": [{ "lit": "api" }, { "lit": "v1" }, { "lit": "mandates" }], "select": { "exist": ["currency", "from_date", "mandate_i_d", "max_amount", "merchant_id", "min_amount", "page", "search", "size", "sort", "status", "to_date"] }, "transform": { "req": "`reqdata`", "res": "`body.content`" }, "index$": 0 }], "key$": "list" } }, "relations": { "ancestors": [] }, "key$": "merchant_direct_debit_mandate_page", "name__orig": "merchant_direct_debit_mandate_page", "Name": "MerchantDirectDebitMandatePage", "name_": "merchant_direct_debit_mandate_page", "name-": "merchant-direct-debit-mandate-page", "NAME": "MERCHANT_DIRECT_DEBIT_MANDATE_PAGE", "index$": 16 }, { "active": true, "entity": "merchant_direct_debit_mandate_page", "key$": "BasicMerchantDirectDebitMandatePageFlow", "kind": "basic", "name": "BasicMerchantDirectDebitMandatePageFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": {}, "match": {}, "op": "list", "spec": [], "valid": [{ "apply": "ItemExists", "def": { "ref": "merchant_direct_debit_mandate_page_ref01" } }], "index$": 0 }] }, 'MerchantDirectDebitMandatePage');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let merchant_direct_debit_mandate_page_ref01_data = Object.values(setup.data.existing.merchant_direct_debit_mandate_page)[0];
        // LIST
        const merchant_direct_debit_mandate_page_ref01_ent = client.MerchantDirectDebitMandatePage();
        const merchant_direct_debit_mandate_page_ref01_match = {};
        const merchant_direct_debit_mandate_page_ref01_list = (await merchant_direct_debit_mandate_page_ref01_ent.list(merchant_direct_debit_mandate_page_ref01_match)).map((e) => e.data());
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/merchant_direct_debit_mandate_page/MerchantDirectDebitMandatePageTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.NofrixionSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['merchant_direct_debit_mandate_page01', 'merchant_direct_debit_mandate_page02', 'merchant_direct_debit_mandate_page03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'NOFRIXION_TEST_MERCHANT_DIRECT_DEBIT_MANDATE_PAGE_ENTID': idmap,
        'NOFRIXION_TEST_LIVE': 'FALSE',
        'NOFRIXION_TEST_EXPLAIN': 'FALSE',
        'NOFRIXION_APIKEY': '',
    });
    idmap = env['NOFRIXION_TEST_MERCHANT_DIRECT_DEBIT_MANDATE_PAGE_ENTID'];
    const live = 'TRUE' === env.NOFRIXION_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['NOFRIXION_TEST_MERCHANT_DIRECT_DEBIT_MANDATE_PAGE_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.NofrixionSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
            {
                apikey: env.NOFRIXION_APIKEY,
            },
            // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
            // last entry is undefined, and basicSetup is normally called with no
            // argument at all - so a bare 'extra' silently discarded the apikey
            // and server values above and handed the SDK undefined. Harmless
            // while there was nothing in that object; not harmless now.
            extra || {},
            { system: { fetch: transport.fetch } }
        ]));
    }
    const setup = {
        idmap,
        env,
        options,
        client,
        struct,
        data: entityData,
        explain: 'TRUE' === env.NOFRIXION_TEST_EXPLAIN,
        live,
        transport,
        now: Date.now(),
    };
    return setup;
}
//# sourceMappingURL=MerchantDirectDebitMandatePageEntity.test.js.map