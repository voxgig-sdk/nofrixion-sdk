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
(0, node_test_1.describe)('CardEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('NOFRIXION_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.NofrixionSDK.test();
        const ent = testsdk.Card();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE;
        for (const op of ['create']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'card.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "authorizedAmount", "req": false, "type": "`$STRING`", "index$": 0 }, { "active": true, "name": "currencyCode", "req": false, "type": "`$STRING`", "index$": 1 }, { "active": true, "name": "isPayerAuthenticationRequired", "req": false, "short": "Gets set to true if 3-D Secure payer authentication is required for a payment.", "type": "`$BOOLEAN`", "index$": 2 }, { "active": true, "name": "isSoftDecline", "req": false, "short": "Gets set to true if the card processor flagged the transaction as having failed address or card security number verification.", "type": "`$BOOLEAN`", "index$": 3 }, { "active": true, "name": "payerAuthenticationAccessToken", "req": false, "short": "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the access token to POST when performing the redirect.", "type": "`$STRING`", "index$": 4 }, { "active": true, "name": "payerAuthenticationMerchantData", "req": false, "short": "If a card payment response indicates a 3-D Secure payer authentication this field may get set in order to transfer information back to the \"authenticationcallback\" method that gets called automatically after a successful payer authenticati…", "type": "`$STRING`", "index$": 5 }, { "active": true, "name": "payerAuthenticationUrl", "req": false, "short": "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the URL to redirect the payer to their issuing bank.", "type": "`$STRING`", "index$": 6 }, { "active": true, "format": "int32", "name": "payerAuthenticationWindowHeight", "req": false, "short": "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested height of the iframe used to hold the challenge.", "type": "`$INTEGER`", "index$": 7 }, { "active": true, "format": "int32", "name": "payerAuthenticationWindowWidth", "req": false, "short": "If a card payment response indicates a 3-D Secure payer authentication is required this field holds the requested width of the iframe used to hold the challenge.", "type": "`$INTEGER`", "index$": 8 }, { "active": true, "name": "paymentRequestCallbackUrl", "req": false, "short": "The callback URL that was set when the payment request was created.", "type": "`$STRING`", "index$": 9 }, { "active": true, "format": "uuid", "name": "paymentRequestID", "req": false, "type": "`$STRING`", "index$": 10 }, { "active": true, "name": "requestID", "req": false, "type": "`$STRING`", "index$": 11 }, { "active": true, "name": "responseCode", "req": false, "type": "`$STRING`", "index$": 12 }, { "active": true, "name": "responseType", "readOnly": true, "req": false, "type": "`$STRING`", "index$": 13 }, { "active": true, "name": "status", "req": false, "type": "`$STRING`", "index$": 14 }, { "active": true, "name": "threeDSRedirectUrl", "req": false, "short": "Checkout.com require a redirect for 3DS authentication.", "type": "`$STRING`", "index$": 15 }, { "active": true, "name": "transactionID", "req": false, "type": "`$STRING`", "index$": 16 }], "name": "card", "op": { "create": { "input": "data", "name": "create", "points": [{ "active": true, "args": { "params": [{ "active": true, "kind": "param", "name": "paymentrequest_id", "orig": "id", "reqd": true, "type": "`$STRING`", "index$": 0 }] }, "contract": { "id": "POST /api/v1/paymentrequests/{id}/card", "json": "{\"operationId\":\"SubmitCardPayment\",\"parameters\":[{\"description\":\"The ID of the payment request the card payment is being submitted for.\",\"in\":\"path\",\"name\":\"id\",\"required\":true,\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}}],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"multipart/form-data\":{\"encoding\":{\"BillingAddressCity\":{\"style\":\"form\"},\"BillingAddressCountryCode\":{\"style\":\"form\"},\"BillingAddressCounty\":{\"style\":\"form\"},\"BillingAddressLine1\":{\"style\":\"form\"},\"BillingAddressLine2\":{\"style\":\"form\"},\"BillingAddressPostCode\":{\"style\":\"form\"},\"BillingEmail\":{\"style\":\"form\"},\"BillingPhone\":{\"style\":\"form\"},\"CardHolderFirstName\":{\"style\":\"form\"},\"CardHolderLastName\":{\"style\":\"form\"},\"CvcToken\":{\"style\":\"form\"},\"ExpiryMonth\":{\"style\":\"form\"},\"ExpiryYear\":{\"style\":\"form\"},\"JwtToken\":{\"style\":\"form\"},\"NumberToken\":{\"style\":\"form\"},\"PartialAmount\":{\"style\":\"form\"},\"SaveCardConsent\":{\"style\":\"form\"}},\"schema\":{\"properties\":{\"BillingAddressCity\":{\"type\":\"string\"},\"BillingAddressCountryCode\":{\"type\":\"string\"},\"BillingAddressCounty\":{\"type\":\"string\"},\"BillingAddressLine1\":{\"type\":\"string\"},\"BillingAddressLine2\":{\"type\":\"string\"},\"BillingAddressPostCode\":{\"type\":\"string\"},\"BillingEmail\":{\"format\":\"email\",\"type\":\"string\"},\"BillingPhone\":{\"type\":\"string\"},\"CardHolderFirstName\":{\"type\":\"string\"},\"CardHolderLastName\":{\"type\":\"string\"},\"CvcToken\":{\"type\":\"string\"},\"ExpiryMonth\":{\"type\":\"string\"},\"ExpiryYear\":{\"type\":\"string\"},\"JwtToken\":{\"description\":\"The tokenised Card Number and CardVerification Number (CVN) details.\\r\\nTypically generated by the CyberSource javascript library.\",\"type\":\"string\"},\"NumberToken\":{\"type\":\"string\"},\"PartialAmount\":{\"description\":\"Optional partial payment amount. If not specified then the full payment request\\r\\namount will be attempted.\",\"format\":\"double\",\"type\":\"number\"},\"SaveCardConsent\":{\"type\":\"boolean\"}},\"required\":[\"ExpiryMonth\",\"ExpiryYear\"],\"type\":\"object\"}}}},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"authorizedAmount\":{\"nullable\":true,\"type\":\"string\"},\"currencyCode\":{\"nullable\":true,\"type\":\"string\"},\"isPayerAuthenticationRequired\":{\"description\":\"Gets set to true if 3-D Secure payer authentication is required for a payment.\",\"type\":\"boolean\"},\"isSoftDecline\":{\"description\":\"Gets set to true if the card processor flagged the transaction as having failed address or\\r\\ncard security number verification. If the payment was for a sale a soft decline will prevent\\r\\nthe transaction being captured. The merchant can set the ignore address verification flag \\r\\nif they want to allow soft declines to be accepted or false if not.\",\"type\":\"boolean\"},\"payerAuthenticationAccessToken\":{\"description\":\"If a card payment response indicates a 3-D Secure payer authentication is required\\r\\nthis field holds the access token to POST when performing the redirect.\",\"nullable\":true,\"type\":\"string\"},\"payerAuthenticationMerchantData\":{\"description\":\"If a card payment response indicates a 3-D Secure payer authentication this field may \\r\\nget set in order to transfer information back to the \\\"authenticationcallback\\\" method\\r\\nthat gets called automatically after a successful payer authentication attempt.\",\"nullable\":true,\"type\":\"string\"},\"payerAuthenticationUrl\":{\"description\":\"If a card payment response indicates a 3-D Secure payer authentication is required\\r\\nthis field holds the URL to redirect the payer to their issuing bank.\",\"nullable\":true,\"type\":\"string\"},\"payerAuthenticationWindowHeight\":{\"description\":\"If a card payment response indicates a 3-D Secure payer authentication is required\\r\\nthis field holds the requested height of the iframe used to hold the challenge.\",\"format\":\"int32\",\"type\":\"integer\"},\"payerAuthenticationWindowWidth\":{\"description\":\"If a card payment response indicates a 3-D Secure payer authentication is required\\r\\nthis field holds the requested width of the iframe used to hold the challenge.\",\"format\":\"int32\",\"type\":\"integer\"},\"paymentRequestCallbackUrl\":{\"description\":\"The callback URL that was set when the payment request was created. Payers will be \\r\\nredirected to this URL after a successful card authorisation.\",\"nullable\":true,\"type\":\"string\"},\"paymentRequestID\":{\"format\":\"uuid\",\"type\":\"string\"},\"requestID\":{\"nullable\":true,\"type\":\"string\"},\"responseCode\":{\"nullable\":true,\"type\":\"string\"},\"responseType\":{\"enum\":[\"None\",\"CardPayerAuthenticationSetupResponse\",\"CardPaymentResponse\",\"PaymentInitiationResponse\"],\"readOnly\":true,\"type\":\"string\"},\"status\":{\"nullable\":true,\"type\":\"string\"},\"threeDSRedirectUrl\":{\"description\":\"Checkout.com require a redirect for 3DS authentication. This link will contain the \\r\\nlink the payer needs to be redirected to.\",\"nullable\":true,\"type\":\"string\"},\"transactionID\":{\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"authorizedAmount\":{\"nullable\":true,\"type\":\"string\"},\"currencyCode\":{\"nullable\":true,\"type\":\"string\"},\"isPayerAuthenticationRequired\":{\"description\":\"Gets set to true if 3-D Secure payer authentication is required for a payment.\",\"type\":\"boolean\"},\"isSoftDecline\":{\"description\":\"Gets set to true if the card processor flagged the transaction as having failed address or\\r\\ncard security number verification. If the payment was for a sale a soft decline will prevent\\r\\nthe transaction being captured. The merchant can set the ignore address verification flag \\r\\nif they want to allow soft declines to be accepted or false if not.\",\"type\":\"boolean\"},\"payerAuthenticationAccessToken\":{\"description\":\"If a card payment response indicates a 3-D Secure payer authentication is required\\r\\nthis field holds the access token to POST when performing the redirect.\",\"nullable\":true,\"type\":\"string\"},\"payerAuthenticationMerchantData\":{\"description\":\"If a card payment response indicates a 3-D Secure payer authentication this field may \\r\\nget set in order to transfer information back to the \\\"authenticationcallback\\\" method\\r\\nthat gets called automatically after a successful payer authentication attempt.\",\"nullable\":true,\"type\":\"string\"},\"payerAuthenticationUrl\":{\"description\":\"If a card payment response indicates a 3-D Secure payer authentication is required\\r\\nthis field holds the URL to redirect the payer to their issuing bank.\",\"nullable\":true,\"type\":\"string\"},\"payerAuthenticationWindowHeight\":{\"description\":\"If a card payment response indicates a 3-D Secure payer authentication is required\\r\\nthis field holds the requested height of the iframe used to hold the challenge.\",\"format\":\"int32\",\"type\":\"integer\"},\"payerAuthenticationWindowWidth\":{\"description\":\"If a card payment response indicates a 3-D Secure payer authentication is required\\r\\nthis field holds the requested width of the iframe used to hold the challenge.\",\"format\":\"int32\",\"type\":\"integer\"},\"paymentRequestCallbackUrl\":{\"description\":\"The callback URL that was set when the payment request was created. Payers will be \\r\\nredirected to this URL after a successful card authorisation.\",\"nullable\":true,\"type\":\"string\"},\"paymentRequestID\":{\"format\":\"uuid\",\"type\":\"string\"},\"requestID\":{\"nullable\":true,\"type\":\"string\"},\"responseCode\":{\"nullable\":true,\"type\":\"string\"},\"responseType\":{\"enum\":[\"None\",\"CardPayerAuthenticationSetupResponse\",\"CardPaymentResponse\",\"PaymentInitiationResponse\"],\"readOnly\":true,\"type\":\"string\"},\"status\":{\"nullable\":true,\"type\":\"string\"},\"threeDSRedirectUrl\":{\"description\":\"Checkout.com require a redirect for 3DS authentication. This link will contain the \\r\\nlink the payer needs to be redirected to.\",\"nullable\":true,\"type\":\"string\"},\"transactionID\":{\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"}},\"text/plain\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"authorizedAmount\":{\"nullable\":true,\"type\":\"string\"},\"currencyCode\":{\"nullable\":true,\"type\":\"string\"},\"isPayerAuthenticationRequired\":{\"description\":\"Gets set to true if 3-D Secure payer authentication is required for a payment.\",\"type\":\"boolean\"},\"isSoftDecline\":{\"description\":\"Gets set to true if the card processor flagged the transaction as having failed address or\\r\\ncard security number verification. If the payment was for a sale a soft decline will prevent\\r\\nthe transaction being captured. The merchant can set the ignore address verification flag \\r\\nif they want to allow soft declines to be accepted or false if not.\",\"type\":\"boolean\"},\"payerAuthenticationAccessToken\":{\"description\":\"If a card payment response indicates a 3-D Secure payer authentication is required\\r\\nthis field holds the access token to POST when performing the redirect.\",\"nullable\":true,\"type\":\"string\"},\"payerAuthenticationMerchantData\":{\"description\":\"If a card payment response indicates a 3-D Secure payer authentication this field may \\r\\nget set in order to transfer information back to the \\\"authenticationcallback\\\" method\\r\\nthat gets called automatically after a successful payer authentication attempt.\",\"nullable\":true,\"type\":\"string\"},\"payerAuthenticationUrl\":{\"description\":\"If a card payment response indicates a 3-D Secure payer authentication is required\\r\\nthis field holds the URL to redirect the payer to their issuing bank.\",\"nullable\":true,\"type\":\"string\"},\"payerAuthenticationWindowHeight\":{\"description\":\"If a card payment response indicates a 3-D Secure payer authentication is required\\r\\nthis field holds the requested height of the iframe used to hold the challenge.\",\"format\":\"int32\",\"type\":\"integer\"},\"payerAuthenticationWindowWidth\":{\"description\":\"If a card payment response indicates a 3-D Secure payer authentication is required\\r\\nthis field holds the requested width of the iframe used to hold the challenge.\",\"format\":\"int32\",\"type\":\"integer\"},\"paymentRequestCallbackUrl\":{\"description\":\"The callback URL that was set when the payment request was created. Payers will be \\r\\nredirected to this URL after a successful card authorisation.\",\"nullable\":true,\"type\":\"string\"},\"paymentRequestID\":{\"format\":\"uuid\",\"type\":\"string\"},\"requestID\":{\"nullable\":true,\"type\":\"string\"},\"responseCode\":{\"nullable\":true,\"type\":\"string\"},\"responseType\":{\"enum\":[\"None\",\"CardPayerAuthenticationSetupResponse\",\"CardPaymentResponse\",\"PaymentInitiationResponse\"],\"readOnly\":true,\"type\":\"string\"},\"status\":{\"nullable\":true,\"type\":\"string\"},\"threeDSRedirectUrl\":{\"description\":\"Checkout.com require a redirect for 3DS authentication. This link will contain the \\r\\nlink the payer needs to be redirected to.\",\"nullable\":true,\"type\":\"string\"},\"transactionID\":{\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Returns a card payment response model that indicates the result of the payment attempt.\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "POST", "orig": "/api/v1/paymentrequests/{id}/card", "rename": { "param": { "id": "paymentrequest_id" } }, "segments": [{ "lit": "api" }, { "lit": "v1" }, { "lit": "paymentrequests" }, { "var": "paymentrequest_id" }, { "lit": "card" }], "select": { "exist": ["paymentrequest_id"] }, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "create" } }, "relations": { "ancestors": [["paymentrequest"]] }, "key$": "card", "name__orig": "card", "Name": "Card", "name_": "card", "name-": "card", "NAME": "CARD", "index$": 4 }, { "active": true, "entity": "card", "key$": "BasicCardFlow", "kind": "basic", "name": "BasicCardFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": { "ref": "card_ref01" }, "match": { "paymentrequest_id": "paymentrequest01" }, "op": "create", "spec": [], "valid": [], "index$": 0 }] }, 'Card');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const card_ref01_ent = client.Card();
        let card_ref01_data = setup.data.new.card['card_ref01'];
        card_ref01_data['paymentrequest_id'] = setup.idmap['paymentrequest01'];
        card_ref01_data = (await card_ref01_ent.create(card_ref01_data)).data();
        (0, node_assert_1.default)(null != card_ref01_data);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/card/CardTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.NofrixionSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['card01', 'card02', 'card03', 'paymentrequest01', 'paymentrequest02', 'paymentrequest03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'NOFRIXION_TEST_CARD_ENTID': idmap,
        'NOFRIXION_TEST_LIVE': 'FALSE',
        'NOFRIXION_TEST_EXPLAIN': 'FALSE',
        'NOFRIXION_APIKEY': '',
    });
    idmap = env['NOFRIXION_TEST_CARD_ENTID'];
    const live = 'TRUE' === env.NOFRIXION_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['NOFRIXION_TEST_CARD_ENTID'];
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
//# sourceMappingURL=CardEntity.test.js.map