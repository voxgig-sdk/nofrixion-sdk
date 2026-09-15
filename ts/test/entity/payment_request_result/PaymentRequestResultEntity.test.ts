

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { NofrixionSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
loadEnvLocal(__dirname + '/../../../.env.local')


describe('PaymentRequestResultEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.PaymentRequestResult()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'payment_request_result.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"format":"double","name":"amount","req":false,"short":"The authorised payment amount.","type":"`$NUMBER`","index$":0},{"active":true,"format":"double","name":"amountPending","req":false,"type":"`$NUMBER`","index$":1},{"active":true,"format":"double","name":"amountReceived","req":false,"type":"`$NUMBER`","index$":2},{"active":true,"format":"double","name":"amountRefunded","req":false,"type":"`$NUMBER`","index$":3},{"active":true,"name":"currency","req":false,"short":"The authorised payment currency.","type":"`$STRING`","index$":4},{"active":true,"name":"customerID","req":false,"short":"The customer id","type":"`$STRING`","index$":5},{"active":true,"format":"uuid","name":"paymentRequestID","req":false,"short":"The ID of the payment request the result is for.","type":"`$STRING`","index$":6},{"active":true,"name":"payments","req":false,"short":"The list of payment attempts that have been received for the payment request.","type":"`$ARRAY`","index$":7},{"active":true,"name":"pispAuthorizations","req":false,"type":"`$ARRAY`","index$":8},{"active":true,"format":"double","name":"requestedAmount","req":false,"short":"The full original payment amount requested.","type":"`$NUMBER`","index$":9},{"active":true,"name":"result","req":false,"short":"The result of the payment attempt.","type":"`$STRING`","index$":10}],"name":"payment_request_result","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"paymentrequest_id","orig":"id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /api/v1/paymentrequests/{id}/result","json":"{\"operationId\":\"GetPaymentRequestResult\",\"parameters\":[{\"description\":\"ID of the payment request to get the result for.\",\"in\":\"path\",\"name\":\"id\",\"required\":true,\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"amount\":{\"description\":\"The authorised payment amount.\",\"format\":\"double\",\"type\":\"number\"},\"amountPending\":{\"format\":\"double\",\"type\":\"number\"},\"amountReceived\":{\"format\":\"double\",\"type\":\"number\"},\"amountRefunded\":{\"format\":\"double\",\"type\":\"number\"},\"currency\":{\"description\":\"The authorised payment currency.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"customerID\":{\"description\":\"The customer id\",\"nullable\":true,\"type\":\"string\"},\"paymentRequestID\":{\"description\":\"The ID of the payment request the result is for.\",\"format\":\"uuid\",\"type\":\"string\"},\"payments\":{\"description\":\"The list of payment attempts that have been received for the payment request.\",\"items\":{\"additionalProperties\":false,\"properties\":{\"amount\":{\"description\":\"The authorised payment amount.\",\"format\":\"double\",\"type\":\"number\"},\"cardAuthorizationID\":{\"description\":\"For card payments this is the ID from the initial successful authorization or sale.\\r\\nRequired for voids and capture operations.\",\"nullable\":true,\"type\":\"string\"},\"cardCapturedAmount\":{\"description\":\"The captured amount for a card payment.\",\"format\":\"double\",\"type\":\"number\"},\"cardIsVoided\":{\"description\":\"If true indicates that the card payment was voided.\",\"type\":\"boolean\"},\"currency\":{\"description\":\"The authorised payment currency.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"occurredAt\":{\"description\":\"Timestamp the payment occurred. For cards this will be the time the\\r\\noriginal authorisation occurred.\",\"format\":\"date-time\",\"type\":\"string\"},\"paymentMethod\":{\"description\":\"The payment type for the received money.\",\"enum\":[\"None\",\"card\",\"pisp\",\"lightning\",\"cardtoken\",\"applePay\",\"googlePay\",\"directDebit\"],\"type\":\"string\"},\"paymentProcessor\":{\"description\":\"The card processor that was used for the payment event.\",\"enum\":[\"None\",\"CyberSource\",\"Checkout\",\"Stripe\",\"Modulr\",\"Plaid\",\"Yapily\",\"Nofrixion\",\"Bitcoin\",\"BitcoinTestnet\",\"BankingCircle\",\"BankingCircleAgency\",\"Simulator\",\"Lightning\",\"LightningTestnet\",\"BankingCircleDirectDebit\",\"Technoxander\"],\"type\":\"string\"},\"paymentRequestID\":{\"description\":\"The ID of the payment request the result is for.\",\"format\":\"uuid\",\"type\":\"string\"},\"refundedAmount\":{\"description\":\"Refunded Amount\",\"format\":\"double\",\"type\":\"number\"},\"tokenisedCardID\":{\"description\":\"For card payments the merchant can request a reusable token for this payer and\\r\\nuse it to submit subsequent merchant initiated payments.\",\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"},\"pispAuthorizations\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"amount\":{\"description\":\"The authorised payment amount.\",\"format\":\"double\",\"type\":\"number\"},\"currency\":{\"description\":\"The authorised payment currency.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"occurredAt\":{\"description\":\"Timestamp the pisp authorization occurred.\",\"format\":\"date-time\",\"type\":\"string\"},\"paymentMethod\":{\"description\":\"The payment type for the received money.\",\"enum\":[\"None\",\"card\",\"pisp\",\"lightning\",\"cardtoken\",\"applePay\",\"googlePay\",\"directDebit\"],\"type\":\"string\"},\"paymentRequestID\":{\"description\":\"The ID of the payment request the pisp authorization is for.\",\"format\":\"uuid\",\"type\":\"string\"},\"pispPaymentInitiationID\":{\"description\":\"For a payment initiation this is the ID returned by the service provider initiating\\r\\nthe payment for us.\",\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"},\"requestedAmount\":{\"description\":\"The full original payment amount requested.\",\"format\":\"double\",\"type\":\"number\"},\"result\":{\"description\":\"The result of the payment attempt.\",\"enum\":[\"None\",\"FullyPaid\",\"PartiallyPaid\",\"OverPaid\",\"Voided\",\"Authorized\"],\"type\":\"string\"}},\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"amount\":{\"description\":\"The authorised payment amount.\",\"format\":\"double\",\"type\":\"number\"},\"amountPending\":{\"format\":\"double\",\"type\":\"number\"},\"amountReceived\":{\"format\":\"double\",\"type\":\"number\"},\"amountRefunded\":{\"format\":\"double\",\"type\":\"number\"},\"currency\":{\"description\":\"The authorised payment currency.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"customerID\":{\"description\":\"The customer id\",\"nullable\":true,\"type\":\"string\"},\"paymentRequestID\":{\"description\":\"The ID of the payment request the result is for.\",\"format\":\"uuid\",\"type\":\"string\"},\"payments\":{\"description\":\"The list of payment attempts that have been received for the payment request.\",\"items\":{\"additionalProperties\":false,\"properties\":{\"amount\":{\"description\":\"The authorised payment amount.\",\"format\":\"double\",\"type\":\"number\"},\"cardAuthorizationID\":{\"description\":\"For card payments this is the ID from the initial successful authorization or sale.\\r\\nRequired for voids and capture operations.\",\"nullable\":true,\"type\":\"string\"},\"cardCapturedAmount\":{\"description\":\"The captured amount for a card payment.\",\"format\":\"double\",\"type\":\"number\"},\"cardIsVoided\":{\"description\":\"If true indicates that the card payment was voided.\",\"type\":\"boolean\"},\"currency\":{\"description\":\"The authorised payment currency.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"occurredAt\":{\"description\":\"Timestamp the payment occurred. For cards this will be the time the\\r\\noriginal authorisation occurred.\",\"format\":\"date-time\",\"type\":\"string\"},\"paymentMethod\":{\"description\":\"The payment type for the received money.\",\"enum\":[\"None\",\"card\",\"pisp\",\"lightning\",\"cardtoken\",\"applePay\",\"googlePay\",\"directDebit\"],\"type\":\"string\"},\"paymentProcessor\":{\"description\":\"The card processor that was used for the payment event.\",\"enum\":[\"None\",\"CyberSource\",\"Checkout\",\"Stripe\",\"Modulr\",\"Plaid\",\"Yapily\",\"Nofrixion\",\"Bitcoin\",\"BitcoinTestnet\",\"BankingCircle\",\"BankingCircleAgency\",\"Simulator\",\"Lightning\",\"LightningTestnet\",\"BankingCircleDirectDebit\",\"Technoxander\"],\"type\":\"string\"},\"paymentRequestID\":{\"description\":\"The ID of the payment request the result is for.\",\"format\":\"uuid\",\"type\":\"string\"},\"refundedAmount\":{\"description\":\"Refunded Amount\",\"format\":\"double\",\"type\":\"number\"},\"tokenisedCardID\":{\"description\":\"For card payments the merchant can request a reusable token for this payer and\\r\\nuse it to submit subsequent merchant initiated payments.\",\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"},\"pispAuthorizations\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"amount\":{\"description\":\"The authorised payment amount.\",\"format\":\"double\",\"type\":\"number\"},\"currency\":{\"description\":\"The authorised payment currency.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"occurredAt\":{\"description\":\"Timestamp the pisp authorization occurred.\",\"format\":\"date-time\",\"type\":\"string\"},\"paymentMethod\":{\"description\":\"The payment type for the received money.\",\"enum\":[\"None\",\"card\",\"pisp\",\"lightning\",\"cardtoken\",\"applePay\",\"googlePay\",\"directDebit\"],\"type\":\"string\"},\"paymentRequestID\":{\"description\":\"The ID of the payment request the pisp authorization is for.\",\"format\":\"uuid\",\"type\":\"string\"},\"pispPaymentInitiationID\":{\"description\":\"For a payment initiation this is the ID returned by the service provider initiating\\r\\nthe payment for us.\",\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"},\"requestedAmount\":{\"description\":\"The full original payment amount requested.\",\"format\":\"double\",\"type\":\"number\"},\"result\":{\"description\":\"The result of the payment attempt.\",\"enum\":[\"None\",\"FullyPaid\",\"PartiallyPaid\",\"OverPaid\",\"Voided\",\"Authorized\"],\"type\":\"string\"}},\"type\":\"object\"}},\"text/plain\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"amount\":{\"description\":\"The authorised payment amount.\",\"format\":\"double\",\"type\":\"number\"},\"amountPending\":{\"format\":\"double\",\"type\":\"number\"},\"amountReceived\":{\"format\":\"double\",\"type\":\"number\"},\"amountRefunded\":{\"format\":\"double\",\"type\":\"number\"},\"currency\":{\"description\":\"The authorised payment currency.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"customerID\":{\"description\":\"The customer id\",\"nullable\":true,\"type\":\"string\"},\"paymentRequestID\":{\"description\":\"The ID of the payment request the result is for.\",\"format\":\"uuid\",\"type\":\"string\"},\"payments\":{\"description\":\"The list of payment attempts that have been received for the payment request.\",\"items\":{\"additionalProperties\":false,\"properties\":{\"amount\":{\"description\":\"The authorised payment amount.\",\"format\":\"double\",\"type\":\"number\"},\"cardAuthorizationID\":{\"description\":\"For card payments this is the ID from the initial successful authorization or sale.\\r\\nRequired for voids and capture operations.\",\"nullable\":true,\"type\":\"string\"},\"cardCapturedAmount\":{\"description\":\"The captured amount for a card payment.\",\"format\":\"double\",\"type\":\"number\"},\"cardIsVoided\":{\"description\":\"If true indicates that the card payment was voided.\",\"type\":\"boolean\"},\"currency\":{\"description\":\"The authorised payment currency.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"occurredAt\":{\"description\":\"Timestamp the payment occurred. For cards this will be the time the\\r\\noriginal authorisation occurred.\",\"format\":\"date-time\",\"type\":\"string\"},\"paymentMethod\":{\"description\":\"The payment type for the received money.\",\"enum\":[\"None\",\"card\",\"pisp\",\"lightning\",\"cardtoken\",\"applePay\",\"googlePay\",\"directDebit\"],\"type\":\"string\"},\"paymentProcessor\":{\"description\":\"The card processor that was used for the payment event.\",\"enum\":[\"None\",\"CyberSource\",\"Checkout\",\"Stripe\",\"Modulr\",\"Plaid\",\"Yapily\",\"Nofrixion\",\"Bitcoin\",\"BitcoinTestnet\",\"BankingCircle\",\"BankingCircleAgency\",\"Simulator\",\"Lightning\",\"LightningTestnet\",\"BankingCircleDirectDebit\",\"Technoxander\"],\"type\":\"string\"},\"paymentRequestID\":{\"description\":\"The ID of the payment request the result is for.\",\"format\":\"uuid\",\"type\":\"string\"},\"refundedAmount\":{\"description\":\"Refunded Amount\",\"format\":\"double\",\"type\":\"number\"},\"tokenisedCardID\":{\"description\":\"For card payments the merchant can request a reusable token for this payer and\\r\\nuse it to submit subsequent merchant initiated payments.\",\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"},\"pispAuthorizations\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"amount\":{\"description\":\"The authorised payment amount.\",\"format\":\"double\",\"type\":\"number\"},\"currency\":{\"description\":\"The authorised payment currency.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"occurredAt\":{\"description\":\"Timestamp the pisp authorization occurred.\",\"format\":\"date-time\",\"type\":\"string\"},\"paymentMethod\":{\"description\":\"The payment type for the received money.\",\"enum\":[\"None\",\"card\",\"pisp\",\"lightning\",\"cardtoken\",\"applePay\",\"googlePay\",\"directDebit\"],\"type\":\"string\"},\"paymentRequestID\":{\"description\":\"The ID of the payment request the pisp authorization is for.\",\"format\":\"uuid\",\"type\":\"string\"},\"pispPaymentInitiationID\":{\"description\":\"For a payment initiation this is the ID returned by the service provider initiating\\r\\nthe payment for us.\",\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"},\"requestedAmount\":{\"description\":\"The full original payment amount requested.\",\"format\":\"double\",\"type\":\"number\"},\"result\":{\"description\":\"The result of the payment attempt.\",\"enum\":[\"None\",\"FullyPaid\",\"PartiallyPaid\",\"OverPaid\",\"Voided\",\"Authorized\"],\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Returns a payment request result record.\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/api/v1/paymentrequests/{id}/result","rename":{"param":{"id":"paymentrequest_id"}},"segments":[{"lit":"api"},{"lit":"v1"},{"lit":"paymentrequests"},{"var":"paymentrequest_id"},{"lit":"result"}],"select":{"exist":["paymentrequest_id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[["paymentrequest"]]},"key$":"payment_request_result","name__orig":"payment_request_result","Name":"PaymentRequestResult","name_":"payment_request_result","name-":"payment-request-result","NAME":"PAYMENT_REQUEST_RESULT","index$":32}, {"active":true,"entity":"payment_request_result","key$":"BasicPaymentRequestResultFlow","kind":"basic","name":"BasicPaymentRequestResultFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{"paymentrequest_id":"paymentrequest01"},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"payment_request_result_ref01"}}],"index$":0}]}, 'PaymentRequestResult')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let payment_request_result_ref01_data = Object.values(setup.data.existing.payment_request_result)[0] as any

    // LIST
    const payment_request_result_ref01_ent = client.PaymentRequestResult()
    const payment_request_result_ref01_match: any = {}
    payment_request_result_ref01_match['paymentrequest_id'] = setup.idmap['paymentrequest01']

    const payment_request_result_ref01_list = (await payment_request_result_ref01_ent.list(payment_request_result_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/payment_request_result/PaymentRequestResultTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = NofrixionSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['payment_request_result01','payment_request_result02','payment_request_result03','paymentrequest01','paymentrequest02','paymentrequest03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'NOFRIXION_TEST_PAYMENT_REQUEST_RESULT_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': '',
  })

  idmap = env['NOFRIXION_TEST_PAYMENT_REQUEST_RESULT_ENTID']

  const live = 'TRUE' === env.NOFRIXION_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['NOFRIXION_TEST_PAYMENT_REQUEST_RESULT_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new NofrixionSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
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
    ]))
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
  }

  return setup
}
  
