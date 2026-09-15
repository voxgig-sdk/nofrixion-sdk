

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


describe('OpenBankingEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.OpenBanking()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['create', 'remove']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'open_banking.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"id","req":false,"type":"`$STRING`","index$":0}],"id":{"field":"id","name":"id","parts":["merchant_id","email"],"sep":"/"},"name":"open_banking","op":{"create":{"input":"data","name":"create","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"account_id","orig":"account_id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"POST /api/v1/openbanking/account/{accountID}/synchronise","json":"{\"operationId\":\"SynchroniseConnectedAccount\",\"parameters\":[{\"description\":\"The ID of the connected account.\",\"in\":\"path\",\"name\":\"accountID\",\"required\":true,\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"OK\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"POST","orig":"/api/v1/openbanking/account/{accountID}/synchronise","rename":{"param":{"accountID":"account_id"}},"segments":[{"lit":"api"},{"lit":"v1"},{"lit":"openbanking"},{"lit":"account"},{"var":"account_id"},{"lit":"synchronise"}],"select":{"$action":"synchronise","exist":["account_id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"},"remove":{"input":"data","name":"remove","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"email","orig":"email","reqd":true,"type":"`$STRING`","index$":0},{"active":true,"kind":"param","name":"merchant_id","orig":"merchant_id","reqd":true,"type":"`$STRING`","index$":1}]},"contract":{"id":"DELETE /api/v1/openbanking/consents/{merchantID}/{email}","json":"{\"operationId\":\"DeleteAllConsents\",\"parameters\":[{\"description\":\"The ID of the merchant to delete the consents for.\",\"in\":\"path\",\"name\":\"merchantID\",\"required\":true,\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}},{\"description\":\"The email address of the end user to delete the consents for.\",\"in\":\"path\",\"name\":\"email\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"OK\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"DELETE","orig":"/api/v1/openbanking/consents/{merchantID}/{email}","rename":{"param":{"merchantID":"merchant_id"}},"segments":[{"lit":"api"},{"lit":"v1"},{"lit":"openbanking"},{"lit":"consents"},{"var":"merchant_id"},{"var":"email"}],"select":{"exist":["email","merchant_id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0},{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"account_id","orig":"account_id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"DELETE /api/v1/openbanking/account/{accountID}","json":"{\"operationId\":\"DeleteConnectedAccount\",\"parameters\":[{\"description\":\"The ID of the connected account.\",\"in\":\"path\",\"name\":\"accountID\",\"required\":true,\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"OK\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"DELETE","orig":"/api/v1/openbanking/account/{accountID}","rename":{"param":{"accountID":"account_id"}},"segments":[{"lit":"api"},{"lit":"v1"},{"lit":"openbanking"},{"lit":"account"},{"var":"account_id"}],"select":{"exist":["account_id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"remove"}},"relations":{"ancestors":[["account"],["consent"]]},"key$":"open_banking","name__orig":"open_banking","Name":"OpenBanking","name_":"open_banking","name-":"open-banking","NAME":"OPEN_BANKING","index$":22}, {"active":true,"entity":"open_banking","key$":"BasicOpenBankingFlow","kind":"basic","name":"BasicOpenBankingFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"open_banking_ref01"},"match":{"account_id":"account01","merchant_id":"merchant01"},"op":"create","spec":[],"valid":[],"index$":0},{"active":true,"data":{},"input":{"ref":"open_banking_ref01","suffix":"_rm0"},"match":{"id":"open_banking01"},"op":"remove","spec":[],"valid":[],"index$":1}]}, 'OpenBanking')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const open_banking_ref01_ent = client.OpenBanking()
    let open_banking_ref01_data = setup.data.new.open_banking['open_banking_ref01']
    open_banking_ref01_data['account_id'] = setup.idmap['account01']
    open_banking_ref01_data['merchant_id'] = setup.idmap['merchant01']

    open_banking_ref01_data = (await open_banking_ref01_ent.create(open_banking_ref01_data)).data()
    assert(null != open_banking_ref01_data.id)


    // REMOVE
    const open_banking_ref01_match_rm0: any = { id: open_banking_ref01_data.id }
    await open_banking_ref01_ent.remove(open_banking_ref01_match_rm0)
  

  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/open_banking/OpenBankingTestData.json')

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
    ['open_banking01','open_banking02','open_banking03','account01','account02','account03','consent01','consent02','consent03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'NOFRIXION_TEST_OPEN_BANKING_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': '',
  })

  idmap = env['NOFRIXION_TEST_OPEN_BANKING_ENTID']

  const live = 'TRUE' === env.NOFRIXION_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['NOFRIXION_TEST_OPEN_BANKING_ENTID']
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
  
