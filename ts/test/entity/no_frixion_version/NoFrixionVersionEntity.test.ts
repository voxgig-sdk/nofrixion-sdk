

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


describe('NoFrixionVersionEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.NoFrixionVersion()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'no_frixion_version.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"format":"int32","name":"buildVersion","req":false,"type":"`$INTEGER`","index$":0},{"active":true,"format":"int32","name":"majorVersion","req":false,"type":"`$INTEGER`","index$":1},{"active":true,"format":"int32","name":"minorVersion","req":false,"type":"`$INTEGER`","index$":2},{"active":true,"name":"releaseName","req":false,"type":"`$STRING`","index$":3}],"name":"no_frixion_version","op":{"load":{"input":"data","name":"load","points":[{"active":true,"args":{},"contract":{"id":"GET /api/v1/metadata/version","json":"{\"operationId\":\"Version\",\"parameters\":[],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"buildVersion\":{\"format\":\"int32\",\"type\":\"integer\"},\"majorVersion\":{\"format\":\"int32\",\"type\":\"integer\"},\"minorVersion\":{\"format\":\"int32\",\"type\":\"integer\"},\"releaseName\":{\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"buildVersion\":{\"format\":\"int32\",\"type\":\"integer\"},\"majorVersion\":{\"format\":\"int32\",\"type\":\"integer\"},\"minorVersion\":{\"format\":\"int32\",\"type\":\"integer\"},\"releaseName\":{\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"}},\"text/plain\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"buildVersion\":{\"format\":\"int32\",\"type\":\"integer\"},\"majorVersion\":{\"format\":\"int32\",\"type\":\"integer\"},\"minorVersion\":{\"format\":\"int32\",\"type\":\"integer\"},\"releaseName\":{\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Returns a model with fields describing the current API version.\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/api/v1/metadata/version","segments":[{"lit":"api"},{"lit":"v1"},{"lit":"metadata"},{"lit":"version"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"no_frixion_version","name__orig":"no_frixion_version","Name":"NoFrixionVersion","name_":"no_frixion_version","name-":"no-frixion-version","NAME":"NO_FRIXION_VERSION","index$":21}, {"active":true,"entity":"no_frixion_version","key$":"BasicNoFrixionVersionFlow","kind":"basic","name":"BasicNoFrixionVersionFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"no_frixion_version_ref01","srcdatavar":"no_frixion_version_ref01_data","suffix":"_dt0"},"match":{},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-no_frixion_version_ref01"}}],"index$":0}]}, 'NoFrixionVersion')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let no_frixion_version_ref01_data = Object.values(setup.data.existing.no_frixion_version)[0] as any

    // LOAD
    const no_frixion_version_ref01_ent = client.NoFrixionVersion()
    const no_frixion_version_ref01_match_dt0: any = {}
    const no_frixion_version_ref01_data_dt0 = (await no_frixion_version_ref01_ent.load(no_frixion_version_ref01_match_dt0)).data()
    assert(null != no_frixion_version_ref01_data_dt0)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/no_frixion_version/NoFrixionVersionTestData.json')

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
    ['no_frixion_version01','no_frixion_version02','no_frixion_version03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'NOFRIXION_TEST_NO_FRIXION_VERSION_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': '',
  })

  idmap = env['NOFRIXION_TEST_NO_FRIXION_VERSION_ENTID']

  const live = 'TRUE' === env.NOFRIXION_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['NOFRIXION_TEST_NO_FRIXION_VERSION_ENTID']
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
  
