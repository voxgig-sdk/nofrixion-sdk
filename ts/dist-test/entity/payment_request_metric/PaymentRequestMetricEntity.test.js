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
(0, node_test_1.describe)('PaymentRequestMetricEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('NOFRIXION_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.NofrixionSDK.test();
        const ent = testsdk.PaymentRequestMetric();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE;
        for (const op of ['load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'payment_request_metric.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [], "name": "payment_request_metric", "op": { "load": { "input": "data", "name": "load", "points": [{ "active": true, "args": { "query": [{ "active": true, "kind": "query", "name": "currency", "orig": "currency", "reqd": false, "type": "`$STRING`", "index$": 0 }, { "active": true, "kind": "query", "name": "from_date", "orig": "from_date", "reqd": false, "type": "`$STRING`", "index$": 1 }, { "active": true, "example": false, "kind": "query", "name": "include_archived", "orig": "include_archived", "reqd": false, "type": "`$BOOLEAN`", "index$": 2 }, { "active": true, "kind": "query", "name": "max_amount", "orig": "max_amount", "reqd": false, "type": "`$NUMBER`", "index$": 3 }, { "active": true, "kind": "query", "name": "merchant_id", "orig": "merchant_id", "reqd": false, "type": "`$STRING`", "index$": 4 }, { "active": true, "kind": "query", "name": "min_amount", "orig": "min_amount", "reqd": false, "type": "`$NUMBER`", "index$": 5 }, { "active": true, "kind": "query", "name": "payment_method", "orig": "payment_method", "reqd": false, "type": "`$ARRAY`", "index$": 6 }, { "active": true, "kind": "query", "name": "search", "orig": "search", "reqd": false, "type": "`$STRING`", "index$": 7 }, { "active": true, "kind": "query", "name": "tag", "orig": "tag", "reqd": false, "type": "`$ARRAY`", "index$": 8 }, { "active": true, "kind": "query", "name": "to_date", "orig": "to_date", "reqd": false, "type": "`$STRING`", "index$": 9 }] }, "contract": { "id": "GET /api/v1/paymentrequests/metrics", "json": "{\"operationId\":\"GetPaymentRequestMetricsForMerchant\",\"parameters\":[{\"description\":\"The merchant ID to get the payment request metrics. This field is ignored for\\r\\n            merchant tokens and required for user tokens.\",\"in\":\"query\",\"name\":\"merchantID\",\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}},{\"description\":\"The date filter to apply to retrieve payment request metrics created after this date.\",\"in\":\"query\",\"name\":\"fromDate\",\"schema\":{\"format\":\"date-time\",\"type\":\"string\"}},{\"description\":\"The date filter to apply to retrieve payment request metrics created up until this date.\",\"in\":\"query\",\"name\":\"toDate\",\"schema\":{\"format\":\"date-time\",\"type\":\"string\"}},{\"description\":\"The text filter to apply to retrieve payment requests with a similar title, description or contact information.k\",\"in\":\"query\",\"name\":\"search\",\"schema\":{\"type\":\"string\"}},{\"description\":\"The currency filter to apply to retrieve payment request metrics with this currency.\",\"in\":\"query\",\"name\":\"currency\",\"schema\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"}},{\"description\":\"The amount filter to apply to retrieve payment request metrics that exceed this amount.\",\"in\":\"query\",\"name\":\"minAmount\",\"schema\":{\"format\":\"double\",\"type\":\"number\"}},{\"description\":\"The amount filter to apply to retrieve payment request metrics that don't exceed this amount.\",\"in\":\"query\",\"name\":\"maxAmount\",\"schema\":{\"format\":\"double\",\"type\":\"number\"}},{\"description\":\"The tag filter to apply to retrieve payment request metrics with at least one of these tags.\",\"in\":\"query\",\"name\":\"tags\",\"schema\":{\"items\":{\"format\":\"uuid\",\"type\":\"string\"},\"type\":\"array\"}},{\"description\":\"The payment method filter to apply to retrieve payment request metrics with any of these payment methods.\",\"in\":\"query\",\"name\":\"paymentMethods\",\"schema\":{\"items\":{\"enum\":[\"None\",\"card\",\"pisp\",\"lightning\",\"cardtoken\",\"applePay\",\"googlePay\",\"directDebit\"],\"type\":\"string\"},\"type\":\"array\"}},{\"description\":\"Flag that indicates whether to include archived payment requests or not.\",\"in\":\"query\",\"name\":\"includeArchived\",\"schema\":{\"default\":false,\"type\":\"boolean\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"all\":{\"description\":\"Total payment request count.\",\"format\":\"int32\",\"type\":\"integer\"},\"authorized\":{\"description\":\"Total payment request count with status Authorized.\",\"format\":\"int32\",\"type\":\"integer\"},\"paid\":{\"description\":\"Total payment request count with status FullyPaid.\",\"format\":\"int32\",\"type\":\"integer\"},\"partiallyPaid\":{\"description\":\"Total payment request count with status PartiallyPaid.\",\"format\":\"int32\",\"type\":\"integer\"},\"totalAmountsByCurrency\":{\"additionalProperties\":{\"additionalProperties\":{\"format\":\"double\",\"type\":\"number\"},\"type\":\"object\"},\"description\":\"The total amounts by status and currency.\",\"nullable\":true,\"type\":\"object\"},\"unpaid\":{\"description\":\"Total payment request count with status None.\",\"format\":\"int32\",\"type\":\"integer\"}},\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"all\":{\"description\":\"Total payment request count.\",\"format\":\"int32\",\"type\":\"integer\"},\"authorized\":{\"description\":\"Total payment request count with status Authorized.\",\"format\":\"int32\",\"type\":\"integer\"},\"paid\":{\"description\":\"Total payment request count with status FullyPaid.\",\"format\":\"int32\",\"type\":\"integer\"},\"partiallyPaid\":{\"description\":\"Total payment request count with status PartiallyPaid.\",\"format\":\"int32\",\"type\":\"integer\"},\"totalAmountsByCurrency\":{\"additionalProperties\":{\"additionalProperties\":{\"format\":\"double\",\"type\":\"number\"},\"type\":\"object\"},\"description\":\"The total amounts by status and currency.\",\"nullable\":true,\"type\":\"object\"},\"unpaid\":{\"description\":\"Total payment request count with status None.\",\"format\":\"int32\",\"type\":\"integer\"}},\"type\":\"object\"}},\"text/plain\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"all\":{\"description\":\"Total payment request count.\",\"format\":\"int32\",\"type\":\"integer\"},\"authorized\":{\"description\":\"Total payment request count with status Authorized.\",\"format\":\"int32\",\"type\":\"integer\"},\"paid\":{\"description\":\"Total payment request count with status FullyPaid.\",\"format\":\"int32\",\"type\":\"integer\"},\"partiallyPaid\":{\"description\":\"Total payment request count with status PartiallyPaid.\",\"format\":\"int32\",\"type\":\"integer\"},\"totalAmountsByCurrency\":{\"additionalProperties\":{\"additionalProperties\":{\"format\":\"double\",\"type\":\"number\"},\"type\":\"object\"},\"description\":\"The total amounts by status and currency.\",\"nullable\":true,\"type\":\"object\"},\"unpaid\":{\"description\":\"Total payment request count with status None.\",\"format\":\"int32\",\"type\":\"integer\"}},\"type\":\"object\"}}},\"description\":\"OK\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/api/v1/paymentrequests/metrics", "segments": [{ "lit": "api" }, { "lit": "v1" }, { "lit": "paymentrequests" }, { "lit": "metrics" }], "select": { "exist": ["currency", "from_date", "include_archived", "max_amount", "merchant_id", "min_amount", "payment_method", "search", "tag", "to_date"] }, "transform": { "req": "`reqdata`", "res": "`body.totalAmountsByCurrency`" }, "index$": 0 }], "key$": "load" } }, "relations": { "ancestors": [] }, "key$": "payment_request_metric", "name__orig": "payment_request_metric", "Name": "PaymentRequestMetric", "name_": "payment_request_metric", "name-": "payment-request-metric", "NAME": "PAYMENT_REQUEST_METRIC", "index$": 30 }, { "active": true, "entity": "payment_request_metric", "key$": "BasicPaymentRequestMetricFlow", "kind": "basic", "name": "BasicPaymentRequestMetricFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": { "ref": "payment_request_metric_ref01", "srcdatavar": "payment_request_metric_ref01_data", "suffix": "_dt0" }, "match": {}, "op": "load", "spec": [], "valid": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-payment_request_metric_ref01" } }], "index$": 0 }] }, 'PaymentRequestMetric');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let payment_request_metric_ref01_data = Object.values(setup.data.existing.payment_request_metric)[0];
        // LOAD
        const payment_request_metric_ref01_ent = client.PaymentRequestMetric();
        const payment_request_metric_ref01_match_dt0 = {};
        const payment_request_metric_ref01_data_dt0 = (await payment_request_metric_ref01_ent.load(payment_request_metric_ref01_match_dt0)).data();
        (0, node_assert_1.default)(null != payment_request_metric_ref01_data_dt0);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/payment_request_metric/PaymentRequestMetricTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.NofrixionSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['payment_request_metric01', 'payment_request_metric02', 'payment_request_metric03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'NOFRIXION_TEST_PAYMENT_REQUEST_METRIC_ENTID': idmap,
        'NOFRIXION_TEST_LIVE': 'FALSE',
        'NOFRIXION_TEST_EXPLAIN': 'FALSE',
        'NOFRIXION_APIKEY': '',
    });
    idmap = env['NOFRIXION_TEST_PAYMENT_REQUEST_METRIC_ENTID'];
    const live = 'TRUE' === env.NOFRIXION_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['NOFRIXION_TEST_PAYMENT_REQUEST_METRIC_ENTID'];
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
//# sourceMappingURL=PaymentRequestMetricEntity.test.js.map