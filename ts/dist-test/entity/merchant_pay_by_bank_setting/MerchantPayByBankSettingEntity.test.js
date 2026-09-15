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
(0, node_test_1.describe)('MerchantPayByBankSettingEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('NOFRIXION_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.NofrixionSDK.test();
        const ent = testsdk.MerchantPayByBankSetting();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE;
        for (const op of ['list']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'merchant_pay_by_bank_setting.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "bankCountryCodes", "req": false, "short": "The list of country codes representing the banks the country supports.", "type": "`$ARRAY`", "index$": 0 }, { "active": true, "format": "uuid", "name": "bankID", "req": false, "short": "ID of the bank to be configured for the merchant.", "type": "`$STRING`", "index$": 1 }, { "active": true, "name": "bankName", "req": false, "short": "Name of the Bank/Institution.", "type": "`$STRING`", "index$": 2 }, { "active": true, "name": "businessInstitutionID", "req": false, "short": "ID that the processor uses to identify the bank (business accounts).", "type": "`$STRING`", "index$": 3 }, { "active": true, "name": "currency", "req": false, "short": "Currency supported by the bank.", "type": "`$STRING`", "index$": 4 }, { "active": true, "name": "logo", "req": false, "short": "URL of the bank's logo.", "type": "`$STRING`", "index$": 5 }, { "active": true, "name": "message", "req": false, "short": "Message relating to specific bank.", "type": "`$STRING`", "index$": 6 }, { "active": true, "name": "messageImageUrl", "req": false, "short": "Optional image URL to be displayed with the message.", "type": "`$STRING`", "index$": 7 }, { "active": true, "format": "int32", "name": "order", "req": false, "short": "Order in which this setting will appear in the UI.", "type": "`$INTEGER`", "index$": 8 }, { "active": true, "name": "personalInstitutionID", "req": false, "short": "ID that the processor uses to identify the bank (personal accounts).", "type": "`$STRING`", "index$": 9 }, { "active": true, "name": "processor", "req": false, "short": "Name of the bank payment processor.", "type": "`$STRING`", "index$": 10 }, { "active": true, "name": "warningHeading", "req": false, "short": "The heading for a warning message related to the bank institution to be displayed to the user.", "type": "`$STRING`", "index$": 11 }, { "active": true, "name": "warningMessage", "req": false, "short": "The warning message related to the bank institution to be displayed to the user.", "type": "`$STRING`", "index$": 12 }], "name": "merchant_pay_by_bank_setting", "op": { "list": { "input": "data", "name": "list", "points": [{ "active": true, "args": { "params": [{ "active": true, "kind": "param", "name": "merchant_id", "orig": "merchant_id", "reqd": true, "type": "`$STRING`", "index$": 0 }], "query": [{ "active": true, "kind": "query", "name": "country_code", "orig": "country_code", "reqd": false, "type": "`$STRING`", "index$": 0 }, { "active": true, "kind": "query", "name": "currency", "orig": "currency", "reqd": false, "type": "`$STRING`", "index$": 1 }, { "active": true, "example": "PIS", "kind": "query", "name": "open_banking_operation", "orig": "open_banking_operation", "reqd": false, "type": "`$STRING`", "index$": 2 }] }, "contract": { "id": "GET /api/v1/merchants/{merchantID}/banksettings", "json": "{\"operationId\":\"GetMerchantBankSettings\",\"parameters\":[{\"description\":\"ID of the merchant.\",\"in\":\"path\",\"name\":\"merchantID\",\"required\":true,\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}},{\"description\":\"Optional. The bank currency.\",\"in\":\"query\",\"name\":\"currency\",\"schema\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"}},{\"description\":\"Optional. Country code for country specific banks.\",\"in\":\"query\",\"name\":\"countryCode\",\"schema\":{\"type\":\"string\"}},{\"description\":\"The open banking operation type.\",\"in\":\"query\",\"name\":\"openBankingOperation\",\"schema\":{\"default\":\"PIS\",\"enum\":[\"None\",\"PIS\",\"AIS\"],\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"description\":\"Represents a collection of merchant bank payment settings.\",\"properties\":{\"merchantID\":{\"description\":\"Merchant to which the settings will be configured.\",\"format\":\"uuid\",\"type\":\"string\"},\"payByBankSettings\":{\"description\":\"Collection of bank payment settings.\",\"items\":{\"additionalProperties\":false,\"description\":\"Represents an individual bank payment setting.\",\"properties\":{\"bankCountryCodes\":{\"description\":\"The list of country codes representing the banks the country supports.\",\"items\":{\"type\":\"string\"},\"nullable\":true,\"type\":\"array\"},\"bankID\":{\"description\":\"ID of the bank to be configured for the merchant.\",\"format\":\"uuid\",\"type\":\"string\"},\"bankName\":{\"description\":\"Name of the Bank/Institution.\",\"nullable\":true,\"type\":\"string\"},\"businessInstitutionID\":{\"description\":\"ID that the processor uses to identify the bank (business accounts).\",\"nullable\":true,\"type\":\"string\"},\"currency\":{\"description\":\"Currency supported by the bank.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"logo\":{\"description\":\"URL of the bank's logo.\",\"nullable\":true,\"type\":\"string\"},\"message\":{\"description\":\"Message relating to specific bank.\",\"nullable\":true,\"type\":\"string\"},\"messageImageUrl\":{\"description\":\"Optional image URL to be displayed with the message.\",\"nullable\":true,\"type\":\"string\"},\"order\":{\"description\":\"Order in which this setting will appear in the UI.\",\"format\":\"int32\",\"type\":\"integer\"},\"personalInstitutionID\":{\"description\":\"ID that the processor uses to identify the bank (personal accounts).\",\"nullable\":true,\"type\":\"string\"},\"processor\":{\"description\":\"Name of the bank payment processor.\",\"enum\":[\"None\",\"CyberSource\",\"Checkout\",\"Stripe\",\"Modulr\",\"Plaid\",\"Yapily\",\"Nofrixion\",\"Bitcoin\",\"BitcoinTestnet\",\"BankingCircle\",\"BankingCircleAgency\",\"Simulator\",\"Lightning\",\"LightningTestnet\",\"BankingCircleDirectDebit\",\"Technoxander\"],\"type\":\"string\"},\"warningHeading\":{\"description\":\"The heading for a warning message related to the bank institution to be displayed to the user.\",\"nullable\":true,\"type\":\"string\"},\"warningMessage\":{\"description\":\"The warning message related to the bank institution to be displayed to the user.\",\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"}},\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"description\":\"Represents a collection of merchant bank payment settings.\",\"properties\":{\"merchantID\":{\"description\":\"Merchant to which the settings will be configured.\",\"format\":\"uuid\",\"type\":\"string\"},\"payByBankSettings\":{\"description\":\"Collection of bank payment settings.\",\"items\":{\"additionalProperties\":false,\"description\":\"Represents an individual bank payment setting.\",\"properties\":{\"bankCountryCodes\":{\"description\":\"The list of country codes representing the banks the country supports.\",\"items\":{\"type\":\"string\"},\"nullable\":true,\"type\":\"array\"},\"bankID\":{\"description\":\"ID of the bank to be configured for the merchant.\",\"format\":\"uuid\",\"type\":\"string\"},\"bankName\":{\"description\":\"Name of the Bank/Institution.\",\"nullable\":true,\"type\":\"string\"},\"businessInstitutionID\":{\"description\":\"ID that the processor uses to identify the bank (business accounts).\",\"nullable\":true,\"type\":\"string\"},\"currency\":{\"description\":\"Currency supported by the bank.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"logo\":{\"description\":\"URL of the bank's logo.\",\"nullable\":true,\"type\":\"string\"},\"message\":{\"description\":\"Message relating to specific bank.\",\"nullable\":true,\"type\":\"string\"},\"messageImageUrl\":{\"description\":\"Optional image URL to be displayed with the message.\",\"nullable\":true,\"type\":\"string\"},\"order\":{\"description\":\"Order in which this setting will appear in the UI.\",\"format\":\"int32\",\"type\":\"integer\"},\"personalInstitutionID\":{\"description\":\"ID that the processor uses to identify the bank (personal accounts).\",\"nullable\":true,\"type\":\"string\"},\"processor\":{\"description\":\"Name of the bank payment processor.\",\"enum\":[\"None\",\"CyberSource\",\"Checkout\",\"Stripe\",\"Modulr\",\"Plaid\",\"Yapily\",\"Nofrixion\",\"Bitcoin\",\"BitcoinTestnet\",\"BankingCircle\",\"BankingCircleAgency\",\"Simulator\",\"Lightning\",\"LightningTestnet\",\"BankingCircleDirectDebit\",\"Technoxander\"],\"type\":\"string\"},\"warningHeading\":{\"description\":\"The heading for a warning message related to the bank institution to be displayed to the user.\",\"nullable\":true,\"type\":\"string\"},\"warningMessage\":{\"description\":\"The warning message related to the bank institution to be displayed to the user.\",\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"}},\"type\":\"object\"}},\"text/plain\":{\"schema\":{\"additionalProperties\":false,\"description\":\"Represents a collection of merchant bank payment settings.\",\"properties\":{\"merchantID\":{\"description\":\"Merchant to which the settings will be configured.\",\"format\":\"uuid\",\"type\":\"string\"},\"payByBankSettings\":{\"description\":\"Collection of bank payment settings.\",\"items\":{\"additionalProperties\":false,\"description\":\"Represents an individual bank payment setting.\",\"properties\":{\"bankCountryCodes\":{\"description\":\"The list of country codes representing the banks the country supports.\",\"items\":{\"type\":\"string\"},\"nullable\":true,\"type\":\"array\"},\"bankID\":{\"description\":\"ID of the bank to be configured for the merchant.\",\"format\":\"uuid\",\"type\":\"string\"},\"bankName\":{\"description\":\"Name of the Bank/Institution.\",\"nullable\":true,\"type\":\"string\"},\"businessInstitutionID\":{\"description\":\"ID that the processor uses to identify the bank (business accounts).\",\"nullable\":true,\"type\":\"string\"},\"currency\":{\"description\":\"Currency supported by the bank.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"logo\":{\"description\":\"URL of the bank's logo.\",\"nullable\":true,\"type\":\"string\"},\"message\":{\"description\":\"Message relating to specific bank.\",\"nullable\":true,\"type\":\"string\"},\"messageImageUrl\":{\"description\":\"Optional image URL to be displayed with the message.\",\"nullable\":true,\"type\":\"string\"},\"order\":{\"description\":\"Order in which this setting will appear in the UI.\",\"format\":\"int32\",\"type\":\"integer\"},\"personalInstitutionID\":{\"description\":\"ID that the processor uses to identify the bank (personal accounts).\",\"nullable\":true,\"type\":\"string\"},\"processor\":{\"description\":\"Name of the bank payment processor.\",\"enum\":[\"None\",\"CyberSource\",\"Checkout\",\"Stripe\",\"Modulr\",\"Plaid\",\"Yapily\",\"Nofrixion\",\"Bitcoin\",\"BitcoinTestnet\",\"BankingCircle\",\"BankingCircleAgency\",\"Simulator\",\"Lightning\",\"LightningTestnet\",\"BankingCircleDirectDebit\",\"Technoxander\"],\"type\":\"string\"},\"warningHeading\":{\"description\":\"The heading for a warning message related to the bank institution to be displayed to the user.\",\"nullable\":true,\"type\":\"string\"},\"warningMessage\":{\"description\":\"The warning message related to the bank institution to be displayed to the user.\",\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"}},\"type\":\"object\"}}},\"description\":\"OK\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/api/v1/merchants/{merchantID}/banksettings", "rename": { "param": { "merchantID": "merchant_id" } }, "segments": [{ "lit": "api" }, { "lit": "v1" }, { "lit": "merchants" }, { "var": "merchant_id" }, { "lit": "banksettings" }], "select": { "exist": ["country_code", "currency", "merchant_id", "open_banking_operation"] }, "transform": { "req": "`reqdata`", "res": "`body.payByBankSettings`" }, "index$": 0 }], "key$": "list" } }, "relations": { "ancestors": [["merchant"]] }, "key$": "merchant_pay_by_bank_setting", "name__orig": "merchant_pay_by_bank_setting", "Name": "MerchantPayByBankSetting", "name_": "merchant_pay_by_bank_setting", "name-": "merchant-pay-by-bank-setting", "NAME": "MERCHANT_PAY_BY_BANK_SETTING", "index$": 17 }, { "active": true, "entity": "merchant_pay_by_bank_setting", "key$": "BasicMerchantPayByBankSettingFlow", "kind": "basic", "name": "BasicMerchantPayByBankSettingFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": {}, "match": { "merchant_id": "merchant01" }, "op": "list", "spec": [], "valid": [{ "apply": "ItemExists", "def": { "ref": "merchant_pay_by_bank_setting_ref01" } }], "index$": 0 }] }, 'MerchantPayByBankSetting');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let merchant_pay_by_bank_setting_ref01_data = Object.values(setup.data.existing.merchant_pay_by_bank_setting)[0];
        // LIST
        const merchant_pay_by_bank_setting_ref01_ent = client.MerchantPayByBankSetting();
        const merchant_pay_by_bank_setting_ref01_match = {};
        merchant_pay_by_bank_setting_ref01_match['merchant_id'] = setup.idmap['merchant01'];
        const merchant_pay_by_bank_setting_ref01_list = (await merchant_pay_by_bank_setting_ref01_ent.list(merchant_pay_by_bank_setting_ref01_match)).map((e) => e.data());
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/merchant_pay_by_bank_setting/MerchantPayByBankSettingTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.NofrixionSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['merchant_pay_by_bank_setting01', 'merchant_pay_by_bank_setting02', 'merchant_pay_by_bank_setting03', 'merchant01', 'merchant02', 'merchant03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'NOFRIXION_TEST_MERCHANT_PAY_BY_BANK_SETTING_ENTID': idmap,
        'NOFRIXION_TEST_LIVE': 'FALSE',
        'NOFRIXION_TEST_EXPLAIN': 'FALSE',
        'NOFRIXION_APIKEY': '',
    });
    idmap = env['NOFRIXION_TEST_MERCHANT_PAY_BY_BANK_SETTING_ENTID'];
    const live = 'TRUE' === env.NOFRIXION_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['NOFRIXION_TEST_MERCHANT_PAY_BY_BANK_SETTING_ENTID'];
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
//# sourceMappingURL=MerchantPayByBankSettingEntity.test.js.map