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
(0, node_test_1.describe)('DirectDebitBatchSubmitEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('NOFRIXION_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.NofrixionSDK.test();
        const ent = testsdk.DirectDebitBatchSubmit();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE;
        for (const op of ['create']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'direct_debit_batch_submit.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "failedSubmissions", "req": false, "short": "Dictionary of failed submissions, keyed by the index (1-based) in the original request.", "type": "`$OBJECT`", "index$": 0 }, { "active": true, "name": "successfulSubmissions", "req": false, "short": "List of successfully submitted direct debit payments.", "type": "`$ARRAY`", "index$": 1 }], "name": "direct_debit_batch_submit", "op": { "create": { "input": "data", "name": "create", "points": [{ "active": true, "args": {}, "contract": { "id": "POST /api/v1/paymentrequests/directdebit/batchsubmit", "json": "{\"operationId\":\"BatchSubmitDirectDebit\",\"parameters\":[],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"mandateID\":{\"description\":\"The ID of the Direct Debit mandate to use for this payment.\",\"format\":\"uuid\",\"type\":\"string\"},\"paymentRequestID\":{\"description\":\"The ID of the payment request to submit for direct debit.\",\"format\":\"uuid\",\"type\":\"string\"},\"submitAfter\":{\"description\":\"Optional. Defines when this payment should be earliest submitted\\r\\nto the customer's bank account. If not specified, payments will be\\r\\nsubmitted as soon as possible.\",\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"}},\"required\":[\"mandateID\",\"paymentRequestID\"],\"type\":\"object\"},\"type\":\"array\"}}},\"description\":\"A list of objects containing payment request IDs and their associated mandate IDs to submit for direct debit.\"},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"failedSubmissions\":{\"additionalProperties\":{\"additionalProperties\":false,\"properties\":{\"problem\":{\"additionalProperties\":false,\"properties\":{\"detail\":{\"description\":\"A human-readable explanation specific to this occurrence of the problem.\",\"example\":\"\\n            Something went wrong with your request and this is a useful explanation of what happened.\\n            \",\"nullable\":true,\"type\":\"string\"},\"errors\":{\"additionalProperties\":{\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"description\":\"Optional validation error messages. These are \\r\\nintended to provide an additional error message about a specific field.\\r\\nNote that the approach to recording validation errors has been deliberately\\r\\nchosen to result in the same serialised JSON as will be returned by ASP.NET \\r\\nwhen the parameter on a controller action fails validation.\",\"nullable\":true,\"type\":\"object\"},\"instance\":{\"description\":\"A URI reference that identifies the specific occurrence of the problem.It may\\r\\nor may not yield further information if dereferenced.\",\"nullable\":true,\"type\":\"string\"},\"status\":{\"description\":\"The HTTP status code([RFC7231], Section 6) generated by the origin server for\\r\\nthis occurrence of the problem.\",\"format\":\"int32\",\"type\":\"integer\"},\"title\":{\"description\":\"A short, human-readable summary of the problem type.It SHOULD NOT change from\\r\\noccurrence to occurrence of the problem, except for purposes of localization (e.g.,\\r\\nusing proactive content negotiation; see[RFC7231], Section 3.4).\",\"example\":\"\\n            Bad Request\\n            \",\"nullable\":true,\"type\":\"string\"},\"traceID\":{\"nullable\":true,\"type\":\"string\"},\"type\":{\"description\":\"A URI reference [RFC3986] that identifies the problem type. This specification\\r\\n encourages that, when dereferenced, it provide human-readable documentation for\\r\\n the problem type (e.g., using HTML [W3C.REC-html5-20141028]). When this member\\r\\n is not present, its value is assumed to be \\\"about:blank\\\".\",\"example\":\"\\n            https://tools.ietf.org/html/rfc7231#section-6.5.1\\n            \",\"nullable\":true,\"type\":\"string\"},\"wasParseFailure\":{\"description\":\"Will be true if an attempt was made to deserialise a problem JSON response.\",\"readOnly\":true,\"type\":\"boolean\"}},\"type\":\"object\"},\"submissionItem\":{\"additionalProperties\":false,\"properties\":{\"mandateID\":{\"description\":\"The ID of the Direct Debit mandate to use for this payment.\",\"format\":\"uuid\",\"type\":\"string\"},\"paymentRequestID\":{\"description\":\"The ID of the payment request to submit for direct debit.\",\"format\":\"uuid\",\"type\":\"string\"},\"submitAfter\":{\"description\":\"Optional. Defines when this payment should be earliest submitted\\r\\nto the customer's bank account. If not specified, payments will be\\r\\nsubmitted as soon as possible.\",\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"}},\"required\":[\"mandateID\",\"paymentRequestID\"],\"type\":\"object\"}},\"required\":[\"problem\",\"submissionItem\"],\"type\":\"object\"},\"description\":\"Dictionary of failed submissions, keyed by the index (1-based) in the original request.\",\"nullable\":true,\"type\":\"object\"},\"successfulSubmissions\":{\"description\":\"List of successfully submitted direct debit payments.\",\"items\":{\"additionalProperties\":false,\"properties\":{\"paymentRequestID\":{\"description\":\"The ID of the payment request that was successfully submitted.\",\"format\":\"uuid\",\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"}},\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"failedSubmissions\":{\"additionalProperties\":{\"additionalProperties\":false,\"properties\":{\"problem\":{\"additionalProperties\":false,\"properties\":{\"detail\":{\"description\":\"A human-readable explanation specific to this occurrence of the problem.\",\"example\":\"\\n            Something went wrong with your request and this is a useful explanation of what happened.\\n            \",\"nullable\":true,\"type\":\"string\"},\"errors\":{\"additionalProperties\":{\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"description\":\"Optional validation error messages. These are \\r\\nintended to provide an additional error message about a specific field.\\r\\nNote that the approach to recording validation errors has been deliberately\\r\\nchosen to result in the same serialised JSON as will be returned by ASP.NET \\r\\nwhen the parameter on a controller action fails validation.\",\"nullable\":true,\"type\":\"object\"},\"instance\":{\"description\":\"A URI reference that identifies the specific occurrence of the problem.It may\\r\\nor may not yield further information if dereferenced.\",\"nullable\":true,\"type\":\"string\"},\"status\":{\"description\":\"The HTTP status code([RFC7231], Section 6) generated by the origin server for\\r\\nthis occurrence of the problem.\",\"format\":\"int32\",\"type\":\"integer\"},\"title\":{\"description\":\"A short, human-readable summary of the problem type.It SHOULD NOT change from\\r\\noccurrence to occurrence of the problem, except for purposes of localization (e.g.,\\r\\nusing proactive content negotiation; see[RFC7231], Section 3.4).\",\"example\":\"\\n            Bad Request\\n            \",\"nullable\":true,\"type\":\"string\"},\"traceID\":{\"nullable\":true,\"type\":\"string\"},\"type\":{\"description\":\"A URI reference [RFC3986] that identifies the problem type. This specification\\r\\n encourages that, when dereferenced, it provide human-readable documentation for\\r\\n the problem type (e.g., using HTML [W3C.REC-html5-20141028]). When this member\\r\\n is not present, its value is assumed to be \\\"about:blank\\\".\",\"example\":\"\\n            https://tools.ietf.org/html/rfc7231#section-6.5.1\\n            \",\"nullable\":true,\"type\":\"string\"},\"wasParseFailure\":{\"description\":\"Will be true if an attempt was made to deserialise a problem JSON response.\",\"readOnly\":true,\"type\":\"boolean\"}},\"type\":\"object\"},\"submissionItem\":{\"additionalProperties\":false,\"properties\":{\"mandateID\":{\"description\":\"The ID of the Direct Debit mandate to use for this payment.\",\"format\":\"uuid\",\"type\":\"string\"},\"paymentRequestID\":{\"description\":\"The ID of the payment request to submit for direct debit.\",\"format\":\"uuid\",\"type\":\"string\"},\"submitAfter\":{\"description\":\"Optional. Defines when this payment should be earliest submitted\\r\\nto the customer's bank account. If not specified, payments will be\\r\\nsubmitted as soon as possible.\",\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"}},\"required\":[\"mandateID\",\"paymentRequestID\"],\"type\":\"object\"}},\"required\":[\"problem\",\"submissionItem\"],\"type\":\"object\"},\"description\":\"Dictionary of failed submissions, keyed by the index (1-based) in the original request.\",\"nullable\":true,\"type\":\"object\"},\"successfulSubmissions\":{\"description\":\"List of successfully submitted direct debit payments.\",\"items\":{\"additionalProperties\":false,\"properties\":{\"paymentRequestID\":{\"description\":\"The ID of the payment request that was successfully submitted.\",\"format\":\"uuid\",\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"}},\"type\":\"object\"}},\"text/plain\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"failedSubmissions\":{\"additionalProperties\":{\"additionalProperties\":false,\"properties\":{\"problem\":{\"additionalProperties\":false,\"properties\":{\"detail\":{\"description\":\"A human-readable explanation specific to this occurrence of the problem.\",\"example\":\"\\n            Something went wrong with your request and this is a useful explanation of what happened.\\n            \",\"nullable\":true,\"type\":\"string\"},\"errors\":{\"additionalProperties\":{\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"description\":\"Optional validation error messages. These are \\r\\nintended to provide an additional error message about a specific field.\\r\\nNote that the approach to recording validation errors has been deliberately\\r\\nchosen to result in the same serialised JSON as will be returned by ASP.NET \\r\\nwhen the parameter on a controller action fails validation.\",\"nullable\":true,\"type\":\"object\"},\"instance\":{\"description\":\"A URI reference that identifies the specific occurrence of the problem.It may\\r\\nor may not yield further information if dereferenced.\",\"nullable\":true,\"type\":\"string\"},\"status\":{\"description\":\"The HTTP status code([RFC7231], Section 6) generated by the origin server for\\r\\nthis occurrence of the problem.\",\"format\":\"int32\",\"type\":\"integer\"},\"title\":{\"description\":\"A short, human-readable summary of the problem type.It SHOULD NOT change from\\r\\noccurrence to occurrence of the problem, except for purposes of localization (e.g.,\\r\\nusing proactive content negotiation; see[RFC7231], Section 3.4).\",\"example\":\"\\n            Bad Request\\n            \",\"nullable\":true,\"type\":\"string\"},\"traceID\":{\"nullable\":true,\"type\":\"string\"},\"type\":{\"description\":\"A URI reference [RFC3986] that identifies the problem type. This specification\\r\\n encourages that, when dereferenced, it provide human-readable documentation for\\r\\n the problem type (e.g., using HTML [W3C.REC-html5-20141028]). When this member\\r\\n is not present, its value is assumed to be \\\"about:blank\\\".\",\"example\":\"\\n            https://tools.ietf.org/html/rfc7231#section-6.5.1\\n            \",\"nullable\":true,\"type\":\"string\"},\"wasParseFailure\":{\"description\":\"Will be true if an attempt was made to deserialise a problem JSON response.\",\"readOnly\":true,\"type\":\"boolean\"}},\"type\":\"object\"},\"submissionItem\":{\"additionalProperties\":false,\"properties\":{\"mandateID\":{\"description\":\"The ID of the Direct Debit mandate to use for this payment.\",\"format\":\"uuid\",\"type\":\"string\"},\"paymentRequestID\":{\"description\":\"The ID of the payment request to submit for direct debit.\",\"format\":\"uuid\",\"type\":\"string\"},\"submitAfter\":{\"description\":\"Optional. Defines when this payment should be earliest submitted\\r\\nto the customer's bank account. If not specified, payments will be\\r\\nsubmitted as soon as possible.\",\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"}},\"required\":[\"mandateID\",\"paymentRequestID\"],\"type\":\"object\"}},\"required\":[\"problem\",\"submissionItem\"],\"type\":\"object\"},\"description\":\"Dictionary of failed submissions, keyed by the index (1-based) in the original request.\",\"nullable\":true,\"type\":\"object\"},\"successfulSubmissions\":{\"description\":\"List of successfully submitted direct debit payments.\",\"items\":{\"additionalProperties\":false,\"properties\":{\"paymentRequestID\":{\"description\":\"The ID of the payment request that was successfully submitted.\",\"format\":\"uuid\",\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"}},\"type\":\"object\"}}},\"description\":\"A DirectDebitBatchSubmitResponse.\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "POST", "orig": "/api/v1/paymentrequests/directdebit/batchsubmit", "segments": [{ "lit": "api" }, { "lit": "v1" }, { "lit": "paymentrequests" }, { "lit": "directdebit" }, { "lit": "batchsubmit" }], "select": {}, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "create" } }, "relations": { "ancestors": [] }, "key$": "direct_debit_batch_submit", "name__orig": "direct_debit_batch_submit", "Name": "DirectDebitBatchSubmit", "name_": "direct_debit_batch_submit", "name-": "direct-debit-batch-submit", "NAME": "DIRECT_DEBIT_BATCH_SUBMIT", "index$": 10 }, { "active": true, "entity": "direct_debit_batch_submit", "key$": "BasicDirectDebitBatchSubmitFlow", "kind": "basic", "name": "BasicDirectDebitBatchSubmitFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": { "ref": "direct_debit_batch_submit_ref01" }, "match": {}, "op": "create", "spec": [], "valid": [], "index$": 0 }] }, 'DirectDebitBatchSubmit');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const direct_debit_batch_submit_ref01_ent = client.DirectDebitBatchSubmit();
        let direct_debit_batch_submit_ref01_data = setup.data.new.direct_debit_batch_submit['direct_debit_batch_submit_ref01'];
        direct_debit_batch_submit_ref01_data = (await direct_debit_batch_submit_ref01_ent.create(direct_debit_batch_submit_ref01_data)).data();
        (0, node_assert_1.default)(null != direct_debit_batch_submit_ref01_data);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/direct_debit_batch_submit/DirectDebitBatchSubmitTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.NofrixionSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['direct_debit_batch_submit01', 'direct_debit_batch_submit02', 'direct_debit_batch_submit03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'NOFRIXION_TEST_DIRECT_DEBIT_BATCH_SUBMIT_ENTID': idmap,
        'NOFRIXION_TEST_LIVE': 'FALSE',
        'NOFRIXION_TEST_EXPLAIN': 'FALSE',
        'NOFRIXION_APIKEY': '',
    });
    idmap = env['NOFRIXION_TEST_DIRECT_DEBIT_BATCH_SUBMIT_ENTID'];
    const live = 'TRUE' === env.NOFRIXION_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['NOFRIXION_TEST_DIRECT_DEBIT_BATCH_SUBMIT_ENTID'];
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
//# sourceMappingURL=DirectDebitBatchSubmitEntity.test.js.map