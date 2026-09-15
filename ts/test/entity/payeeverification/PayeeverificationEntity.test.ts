

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


describe('PayeeverificationEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.Payeeverification()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['create']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'payeeverification.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"accountName","req":true,"short":"The name of the account to verify","type":"`$STRING`","index$":0},{"active":true,"name":"accountNumber","req":false,"short":"The account number of the account to verify (for CoP checks)","type":"`$STRING`","index$":1},{"active":true,"name":"iban","req":true,"short":"The IBAN of the account to verify (for VoP checks)","type":"`$STRING`","index$":2},{"active":true,"name":"payeeVerifiedAccountName","req":false,"short":"The verified account name of the payee, if available (in case of a close match)","type":"`$STRING`","index$":3},{"active":true,"name":"result","req":false,"short":"The result of the payee verification","type":"`$STRING`","index$":4},{"active":true,"name":"secondaryIdentification","req":false,"short":"Optional secondary identifier for the account to verify.","type":"`$STRING`","index$":5},{"active":true,"name":"sortCode","req":false,"short":"The sort code of the account to verify (for CoP checks)","type":"`$STRING`","index$":6}],"name":"payeeverification","op":{"create":{"input":"data","name":"create","points":[{"active":true,"args":{},"contract":{"id":"POST /api/v1/openbanking/payeeverification","json":"{\"operationId\":\"VerifyPayee\",\"parameters\":[],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/*+json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"accountName\":{\"description\":\"The name of the account to verify\",\"minLength\":1,\"type\":\"string\"},\"accountNumber\":{\"description\":\"The account number of the account to verify (for CoP checks)\",\"nullable\":true,\"type\":\"string\"},\"iban\":{\"description\":\"The IBAN of the account to verify (for VoP checks)\",\"minLength\":1,\"type\":\"string\"},\"secondaryIdentification\":{\"description\":\"Optional secondary identifier for the account to verify.\\r\\nIt is usually the reason why the payment is being made or what invoice or obligation it relates to.\\r\\nSome responders may require this where just the identifier is not sufficient to uniquely identify the account.\",\"nullable\":true,\"type\":\"string\"},\"sortCode\":{\"description\":\"The sort code of the account to verify (for CoP checks)\",\"nullable\":true,\"type\":\"string\"}},\"required\":[\"accountName\",\"iban\"],\"type\":\"object\"}},\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"accountName\":{\"description\":\"The name of the account to verify\",\"minLength\":1,\"type\":\"string\"},\"accountNumber\":{\"description\":\"The account number of the account to verify (for CoP checks)\",\"nullable\":true,\"type\":\"string\"},\"iban\":{\"description\":\"The IBAN of the account to verify (for VoP checks)\",\"minLength\":1,\"type\":\"string\"},\"secondaryIdentification\":{\"description\":\"Optional secondary identifier for the account to verify.\\r\\nIt is usually the reason why the payment is being made or what invoice or obligation it relates to.\\r\\nSome responders may require this where just the identifier is not sufficient to uniquely identify the account.\",\"nullable\":true,\"type\":\"string\"},\"sortCode\":{\"description\":\"The sort code of the account to verify (for CoP checks)\",\"nullable\":true,\"type\":\"string\"}},\"required\":[\"accountName\",\"iban\"],\"type\":\"object\"}},\"application/json-patch+json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"accountName\":{\"description\":\"The name of the account to verify\",\"minLength\":1,\"type\":\"string\"},\"accountNumber\":{\"description\":\"The account number of the account to verify (for CoP checks)\",\"nullable\":true,\"type\":\"string\"},\"iban\":{\"description\":\"The IBAN of the account to verify (for VoP checks)\",\"minLength\":1,\"type\":\"string\"},\"secondaryIdentification\":{\"description\":\"Optional secondary identifier for the account to verify.\\r\\nIt is usually the reason why the payment is being made or what invoice or obligation it relates to.\\r\\nSome responders may require this where just the identifier is not sufficient to uniquely identify the account.\",\"nullable\":true,\"type\":\"string\"},\"sortCode\":{\"description\":\"The sort code of the account to verify (for CoP checks)\",\"nullable\":true,\"type\":\"string\"}},\"required\":[\"accountName\",\"iban\"],\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"accountName\":{\"description\":\"The name of the account to verify\",\"minLength\":1,\"type\":\"string\"},\"accountNumber\":{\"description\":\"The account number of the account to verify (for CoP checks)\",\"nullable\":true,\"type\":\"string\"},\"iban\":{\"description\":\"The IBAN of the account to verify (for VoP checks)\",\"minLength\":1,\"type\":\"string\"},\"secondaryIdentification\":{\"description\":\"Optional secondary identifier for the account to verify.\\r\\nIt is usually the reason why the payment is being made or what invoice or obligation it relates to.\\r\\nSome responders may require this where just the identifier is not sufficient to uniquely identify the account.\",\"nullable\":true,\"type\":\"string\"},\"sortCode\":{\"description\":\"The sort code of the account to verify (for CoP checks)\",\"nullable\":true,\"type\":\"string\"}},\"required\":[\"accountName\",\"iban\"],\"type\":\"object\"}}},\"description\":\"The payee to verify.\"},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"payeeVerifiedAccountName\":{\"description\":\"The verified account name of the payee, if available (in case of a close match)\",\"nullable\":true,\"type\":\"string\"},\"result\":{\"description\":\"The result of the payee verification\",\"enum\":[\"Unknown\",\"Match\",\"NoMatch\",\"CloseMatch\"],\"type\":\"string\"}},\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"payeeVerifiedAccountName\":{\"description\":\"The verified account name of the payee, if available (in case of a close match)\",\"nullable\":true,\"type\":\"string\"},\"result\":{\"description\":\"The result of the payee verification\",\"enum\":[\"Unknown\",\"Match\",\"NoMatch\",\"CloseMatch\"],\"type\":\"string\"}},\"type\":\"object\"}},\"text/plain\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"payeeVerifiedAccountName\":{\"description\":\"The verified account name of the payee, if available (in case of a close match)\",\"nullable\":true,\"type\":\"string\"},\"result\":{\"description\":\"The result of the payee verification\",\"enum\":[\"Unknown\",\"Match\",\"NoMatch\",\"CloseMatch\"],\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"OK\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"POST","orig":"/api/v1/openbanking/payeeverification","segments":[{"lit":"api"},{"lit":"v1"},{"lit":"openbanking"},{"lit":"payeeverification"}],"select":{},"transform":{"req":{"accountName":"`reqdata.account_name`","accountNumber":"`reqdata.account_number`","iban":"`reqdata.iban`","secondaryIdentification":"`reqdata.secondary_identification`","sortCode":"`reqdata.sort_code`"},"res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"payeeverification","name__orig":"payeeverification","Name":"Payeeverification","name_":"payeeverification","name-":"payeeverification","NAME":"PAYEEVERIFICATION","index$":23}, {"active":true,"entity":"payeeverification","key$":"BasicPayeeverificationFlow","kind":"basic","name":"BasicPayeeverificationFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"payeeverification_ref01"},"match":{},"op":"create","spec":[],"valid":[],"index$":0}]}, 'Payeeverification')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const payeeverification_ref01_ent = client.Payeeverification()
    let payeeverification_ref01_data = setup.data.new.payeeverification['payeeverification_ref01']

    payeeverification_ref01_data = (await payeeverification_ref01_ent.create(payeeverification_ref01_data)).data()
    assert(null != payeeverification_ref01_data)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/payeeverification/PayeeverificationTestData.json')

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
    ['payeeverification01','payeeverification02','payeeverification03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'NOFRIXION_TEST_PAYEEVERIFICATION_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': '',
  })

  idmap = env['NOFRIXION_TEST_PAYEEVERIFICATION_ENTID']

  const live = 'TRUE' === env.NOFRIXION_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['NOFRIXION_TEST_PAYEEVERIFICATION_ENTID']
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
  
