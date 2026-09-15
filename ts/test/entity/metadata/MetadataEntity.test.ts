

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


describe('MetadataEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.Metadata()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'metadata.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[],"name":"metadata","op":{"load":{"input":"data","name":"load","points":[{"active":true,"args":{"query":[{"active":true,"kind":"query","name":"id","orig":"id","reqd":false,"type":"`$STRING`","index$":0},{"active":true,"kind":"query","name":"secret","orig":"secret","reqd":false,"type":"`$STRING`","index$":1},{"active":true,"kind":"query","name":"url","orig":"url","reqd":false,"type":"`$STRING`","index$":2}]},"contract":{"id":"GET /api/v1/metadata/problemnotification","json":"{\"operationId\":\"ProblemNotification\",\"parameters\":[{\"description\":\"The URL to send the HTTP POST request to with the problem notification.\",\"in\":\"query\",\"name\":\"url\",\"schema\":{\"type\":\"string\"}},{\"description\":\"Optional. If the id and secret are set the notification request will have the\\r\\n            HMAC HTTP authentication headers added.\",\"in\":\"query\",\"name\":\"id\",\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}},{\"description\":\"Optional. If the id and secret are set the notification request will have the\\r\\n            HMAC HTTP authentication headers added.\",\"in\":\"query\",\"name\":\"secret\",\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"Returns OK if the problem notification request is accepted.\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/api/v1/metadata/problemnotification","segments":[{"lit":"api"},{"lit":"v1"},{"lit":"metadata"},{"lit":"problemnotification"}],"select":{"$action":"problemnotification","exist":["id","secret","url"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0},{"active":true,"args":{},"contract":{"id":"GET /api/v1/metadata/problem","json":"{\"operationId\":\"Problem\",\"parameters\":[],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"OK\"},\"400\":{\"description\":\"Returns an example problem response.\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/api/v1/metadata/problem","segments":[{"lit":"api"},{"lit":"v1"},{"lit":"metadata"},{"lit":"problem"}],"select":{"$action":"problem"},"transform":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"metadata","name__orig":"metadata","Name":"Metadata","name_":"metadata","name-":"metadata","NAME":"METADATA","index$":20}, {"active":true,"entity":"metadata","key$":"BasicMetadataFlow","kind":"basic","name":"BasicMetadataFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"metadata_ref01","srcdatavar":"metadata_ref01_data","suffix":"_dt0"},"match":{},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-metadata_ref01"}}],"index$":0}]}, 'Metadata')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let metadata_ref01_data = Object.values(setup.data.existing.metadata)[0] as any

    // LOAD
    const metadata_ref01_ent = client.Metadata()
    const metadata_ref01_match_dt0: any = {}
    const metadata_ref01_data_dt0 = (await metadata_ref01_ent.load(metadata_ref01_match_dt0)).data()
    assert(null != metadata_ref01_data_dt0)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/metadata/MetadataTestData.json')

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
    ['metadata01','metadata02','metadata03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'NOFRIXION_TEST_METADATA_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': '',
  })

  idmap = env['NOFRIXION_TEST_METADATA_ENTID']

  const live = 'TRUE' === env.NOFRIXION_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['NOFRIXION_TEST_METADATA_ENTID']
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
  
