

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


describe('TagEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.Tag()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['create', 'list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'tag.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"colourHex","req":false,"type":"`$STRING`","index$":0},{"active":true,"name":"description","req":false,"type":"`$STRING`","index$":1},{"active":true,"format":"uuid","name":"id","req":false,"type":"`$STRING`","index$":2},{"active":true,"format":"uuid","name":"merchantID","req":true,"type":"`$STRING`","index$":3},{"active":true,"name":"name","req":true,"type":"`$STRING`","index$":4}],"id":{"field":"id","name":"id"},"name":"tag","op":{"create":{"input":"data","name":"create","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"merchant_id","orig":"merchant_id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"POST /api/v1/merchants/{merchantID}/tags","json":"{\"operationId\":\"CreateMerchantTag\",\"parameters\":[{\"description\":\"The merchant id\",\"in\":\"path\",\"name\":\"merchantID\",\"required\":true,\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}}],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"colourHex\":{\"nullable\":true,\"type\":\"string\"},\"description\":{\"nullable\":true,\"type\":\"string\"},\"id\":{\"format\":\"uuid\",\"type\":\"string\"},\"merchantID\":{\"format\":\"uuid\",\"type\":\"string\"},\"name\":{\"minLength\":1,\"type\":\"string\"}},\"required\":[\"merchantID\",\"name\"],\"type\":\"object\"}}},\"description\":\"The tag to add\"},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"colourHex\":{\"nullable\":true,\"type\":\"string\"},\"description\":{\"nullable\":true,\"type\":\"string\"},\"id\":{\"format\":\"uuid\",\"type\":\"string\"},\"merchantID\":{\"format\":\"uuid\",\"type\":\"string\"},\"name\":{\"minLength\":1,\"type\":\"string\"}},\"required\":[\"merchantID\",\"name\"],\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"colourHex\":{\"nullable\":true,\"type\":\"string\"},\"description\":{\"nullable\":true,\"type\":\"string\"},\"id\":{\"format\":\"uuid\",\"type\":\"string\"},\"merchantID\":{\"format\":\"uuid\",\"type\":\"string\"},\"name\":{\"minLength\":1,\"type\":\"string\"}},\"required\":[\"merchantID\",\"name\"],\"type\":\"object\"}},\"text/plain\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"colourHex\":{\"nullable\":true,\"type\":\"string\"},\"description\":{\"nullable\":true,\"type\":\"string\"},\"id\":{\"format\":\"uuid\",\"type\":\"string\"},\"merchantID\":{\"format\":\"uuid\",\"type\":\"string\"},\"name\":{\"minLength\":1,\"type\":\"string\"}},\"required\":[\"merchantID\",\"name\"],\"type\":\"object\"}}},\"description\":\"OK\"},\"201\":{\"description\":\"The newly created tag.\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"POST","orig":"/api/v1/merchants/{merchantID}/tags","rename":{"param":{"merchantID":"merchant_id"}},"segments":[{"lit":"api"},{"lit":"v1"},{"lit":"merchants"},{"var":"merchant_id"},{"lit":"tags"}],"select":{"exist":["merchant_id"]},"transform":{"req":{"colourHex":"`reqdata.colour_hex`","description":"`reqdata.description`","id":"`reqdata.id`","merchantID":"`reqdata.merchant_id`","name":"`reqdata.name`"},"res":"`body`"},"index$":0}],"key$":"create"},"list":{"input":"data","name":"list","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"merchant_id","orig":"merchant_id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /api/v1/merchants/{merchantID}/tags","json":"{\"operationId\":\"GetMerchantTags\",\"parameters\":[{\"description\":\"The merchant id\",\"in\":\"path\",\"name\":\"merchantID\",\"required\":true,\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"colourHex\":{\"nullable\":true,\"type\":\"string\"},\"description\":{\"nullable\":true,\"type\":\"string\"},\"id\":{\"format\":\"uuid\",\"type\":\"string\"},\"merchantID\":{\"format\":\"uuid\",\"type\":\"string\"},\"name\":{\"minLength\":1,\"type\":\"string\"}},\"required\":[\"merchantID\",\"name\"],\"type\":\"object\"},\"type\":\"array\"}},\"text/json\":{\"schema\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"colourHex\":{\"nullable\":true,\"type\":\"string\"},\"description\":{\"nullable\":true,\"type\":\"string\"},\"id\":{\"format\":\"uuid\",\"type\":\"string\"},\"merchantID\":{\"format\":\"uuid\",\"type\":\"string\"},\"name\":{\"minLength\":1,\"type\":\"string\"}},\"required\":[\"merchantID\",\"name\"],\"type\":\"object\"},\"type\":\"array\"}},\"text/plain\":{\"schema\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"colourHex\":{\"nullable\":true,\"type\":\"string\"},\"description\":{\"nullable\":true,\"type\":\"string\"},\"id\":{\"format\":\"uuid\",\"type\":\"string\"},\"merchantID\":{\"format\":\"uuid\",\"type\":\"string\"},\"name\":{\"minLength\":1,\"type\":\"string\"}},\"required\":[\"merchantID\",\"name\"],\"type\":\"object\"},\"type\":\"array\"}}},\"description\":\"Returns a list of merchant tags.\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/api/v1/merchants/{merchantID}/tags","rename":{"param":{"merchantID":"merchant_id"}},"segments":[{"lit":"api"},{"lit":"v1"},{"lit":"merchants"},{"var":"merchant_id"},{"lit":"tags"}],"select":{"exist":["merchant_id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[["merchant"]]},"key$":"tag","name__orig":"tag","Name":"Tag","name_":"tag","name-":"tag","NAME":"TAG","index$":42}, {"active":true,"entity":"tag","key$":"BasicTagFlow","kind":"basic","name":"BasicTagFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"tag_ref01"},"match":{"merchant_id":"merchant01"},"op":"create","spec":[],"valid":[],"index$":0},{"active":true,"data":{},"input":{},"match":{"merchant_id":"merchant01"},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"tag_ref01"}}],"index$":1}]}, 'Tag')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const tag_ref01_ent = client.Tag()
    let tag_ref01_data = setup.data.new.tag['tag_ref01']
    tag_ref01_data['merchant_id'] = setup.idmap['merchant01']

    tag_ref01_data = (await tag_ref01_ent.create(tag_ref01_data)).data()
    assert(null != tag_ref01_data.id)


    // LIST
    const tag_ref01_match: any = {}
    tag_ref01_match['merchant_id'] = setup.idmap['merchant01']

    const tag_ref01_list = (await tag_ref01_ent.list(tag_ref01_match)).map((e: any) => e.data())

    assert(!isempty(select(tag_ref01_list, { id: tag_ref01_data.id })))


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/tag/TagTestData.json')

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
    ['tag01','tag02','tag03','merchant01','merchant02','merchant03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'NOFRIXION_TEST_TAG_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': '',
  })

  idmap = env['NOFRIXION_TEST_TAG_ENTID']

  const live = 'TRUE' === env.NOFRIXION_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['NOFRIXION_TEST_TAG_ENTID']
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
  
