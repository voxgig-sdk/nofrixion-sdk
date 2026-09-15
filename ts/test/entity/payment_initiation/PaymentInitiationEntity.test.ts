

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


describe('PaymentInitiationEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.PaymentInitiation()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['create']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'payment_initiation.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"paymentInitiationID","req":false,"short":"The unique identifier of the payment initiation request.","type":"`$STRING`","index$":0},{"active":true,"name":"paymentRequestCallbackUrl","req":false,"short":"The callback URL that was set when the payment request was created.","type":"`$STRING`","index$":1},{"active":true,"format":"uuid","name":"paymentRequestID","req":false,"type":"`$STRING`","index$":2},{"active":true,"name":"redirectUrl","req":false,"short":"A redirect URL for the user to authorise the payment initiation request at the ASPSP","type":"`$STRING`","index$":3},{"active":true,"name":"responseType","readOnly":true,"req":false,"type":"`$STRING`","index$":4},{"active":true,"name":"specificErrorMessage","req":false,"type":"`$STRING`","index$":5}],"name":"payment_initiation","op":{"create":{"input":"data","name":"create","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"paymentrequest_id","orig":"id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"POST /api/v1/paymentrequests/{id}/pisp","json":"{\"operationId\":\"SubmitPayByBank\",\"parameters\":[{\"description\":\"The ID of the payment request the payment initiation is being submitted for.\",\"in\":\"path\",\"name\":\"id\",\"required\":true,\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}}],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"multipart/form-data\":{\"encoding\":{\"OriginUrl\":{\"style\":\"form\"},\"PartialAmount\":{\"style\":\"form\"},\"ProviderID\":{\"style\":\"form\"},\"RedirectToOriginUrl\":{\"style\":\"form\"}},\"schema\":{\"properties\":{\"OriginUrl\":{\"description\":\"Optional. If set should indicate the origin URL the payer is making the \\r\\npayment from. If a pay by bank attempt fails and the payment request does not\\r\\nhave a FailureCallbackUrl set then the payer will be redirected to this URL.\",\"type\":\"string\"},\"PartialAmount\":{\"description\":\"Optional. If 0 the full amount is assumed.\",\"format\":\"double\",\"type\":\"number\"},\"ProviderID\":{\"description\":\"This is the ID of the institution (bank) that the payer ha chosen.\",\"type\":\"string\"},\"RedirectToOriginUrl\":{\"deprecated\":true,\"type\":\"string\"}},\"required\":[\"ProviderID\"],\"type\":\"object\"}}}},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"paymentInitiationID\":{\"description\":\"The unique identifier of the payment initiation request.\",\"nullable\":true,\"type\":\"string\"},\"paymentRequestCallbackUrl\":{\"description\":\"The callback URL that was set when the payment request was created. Payers will be \\r\\nredirected to this URL after a successful payment initiation.\",\"nullable\":true,\"type\":\"string\"},\"paymentRequestID\":{\"format\":\"uuid\",\"type\":\"string\"},\"redirectUrl\":{\"description\":\"A redirect URL for the user to authorise the payment initiation request at the ASPSP\",\"nullable\":true,\"type\":\"string\"},\"responseType\":{\"enum\":[\"None\",\"CardPayerAuthenticationSetupResponse\",\"CardPaymentResponse\",\"PaymentInitiationResponse\"],\"readOnly\":true,\"type\":\"string\"},\"specificErrorMessage\":{\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"paymentInitiationID\":{\"description\":\"The unique identifier of the payment initiation request.\",\"nullable\":true,\"type\":\"string\"},\"paymentRequestCallbackUrl\":{\"description\":\"The callback URL that was set when the payment request was created. Payers will be \\r\\nredirected to this URL after a successful payment initiation.\",\"nullable\":true,\"type\":\"string\"},\"paymentRequestID\":{\"format\":\"uuid\",\"type\":\"string\"},\"redirectUrl\":{\"description\":\"A redirect URL for the user to authorise the payment initiation request at the ASPSP\",\"nullable\":true,\"type\":\"string\"},\"responseType\":{\"enum\":[\"None\",\"CardPayerAuthenticationSetupResponse\",\"CardPaymentResponse\",\"PaymentInitiationResponse\"],\"readOnly\":true,\"type\":\"string\"},\"specificErrorMessage\":{\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"}},\"text/plain\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"paymentInitiationID\":{\"description\":\"The unique identifier of the payment initiation request.\",\"nullable\":true,\"type\":\"string\"},\"paymentRequestCallbackUrl\":{\"description\":\"The callback URL that was set when the payment request was created. Payers will be \\r\\nredirected to this URL after a successful payment initiation.\",\"nullable\":true,\"type\":\"string\"},\"paymentRequestID\":{\"format\":\"uuid\",\"type\":\"string\"},\"redirectUrl\":{\"description\":\"A redirect URL for the user to authorise the payment initiation request at the ASPSP\",\"nullable\":true,\"type\":\"string\"},\"responseType\":{\"enum\":[\"None\",\"CardPayerAuthenticationSetupResponse\",\"CardPaymentResponse\",\"PaymentInitiationResponse\"],\"readOnly\":true,\"type\":\"string\"},\"specificErrorMessage\":{\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Returns a payment initiation response that contains the payment ID and the payment link.\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"POST","orig":"/api/v1/paymentrequests/{id}/pisp","rename":{"param":{"id":"paymentrequest_id"}},"segments":[{"lit":"api"},{"lit":"v1"},{"lit":"paymentrequests"},{"var":"paymentrequest_id"},{"lit":"pisp"}],"select":{"exist":["paymentrequest_id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[["paymentrequest"]]},"key$":"payment_initiation","name__orig":"payment_initiation","Name":"PaymentInitiation","name_":"payment_initiation","name-":"payment-initiation","NAME":"PAYMENT_INITIATION","index$":27}, {"active":true,"entity":"payment_initiation","key$":"BasicPaymentInitiationFlow","kind":"basic","name":"BasicPaymentInitiationFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"payment_initiation_ref01"},"match":{"paymentrequest_id":"paymentrequest01"},"op":"create","spec":[],"valid":[],"index$":0}]}, 'PaymentInitiation')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const payment_initiation_ref01_ent = client.PaymentInitiation()
    let payment_initiation_ref01_data = setup.data.new.payment_initiation['payment_initiation_ref01']
    payment_initiation_ref01_data['paymentrequest_id'] = setup.idmap['paymentrequest01']

    payment_initiation_ref01_data = (await payment_initiation_ref01_ent.create(payment_initiation_ref01_data)).data()
    assert(null != payment_initiation_ref01_data)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/payment_initiation/PaymentInitiationTestData.json')

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
    ['payment_initiation01','payment_initiation02','payment_initiation03','paymentrequest01','paymentrequest02','paymentrequest03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'NOFRIXION_TEST_PAYMENT_INITIATION_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': '',
  })

  idmap = env['NOFRIXION_TEST_PAYMENT_INITIATION_ENTID']

  const live = 'TRUE' === env.NOFRIXION_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['NOFRIXION_TEST_PAYMENT_INITIATION_ENTID']
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
  
