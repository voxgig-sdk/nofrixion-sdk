

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


describe('MerchantAuthorisationSettingEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.MerchantAuthorisationSetting()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'merchant_authorisation_setting.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"format":"double","name":"amountLower","req":false,"type":"`$NUMBER`","index$":0},{"active":true,"format":"double","name":"amountUpper","req":false,"type":"`$NUMBER`","index$":1},{"active":true,"name":"authorisationType","req":false,"type":"`$STRING`","index$":2},{"active":true,"name":"beneficiariesOnly","req":false,"type":"`$BOOLEAN`","index$":3},{"active":true,"format":"uuid","name":"id","req":false,"type":"`$STRING`","index$":4},{"active":true,"format":"date-time","name":"inserted","req":false,"type":"`$STRING`","index$":5},{"active":true,"name":"lastEditorCantAuthorise","req":false,"type":"`$BOOLEAN`","index$":6},{"active":true,"format":"date-time","name":"lastUpdated","req":false,"type":"`$STRING`","index$":7},{"active":true,"format":"uuid","name":"merchantID","req":false,"type":"`$STRING`","index$":8},{"active":true,"format":"int32","name":"numberOfAuthorisers","req":false,"type":"`$INTEGER`","index$":9},{"active":true,"name":"roleSettings","req":false,"type":"`$ARRAY`","index$":10}],"id":{"field":"id","name":"id"},"name":"merchant_authorisation_setting","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"merchant_id","orig":"merchant_id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /api/v1/merchants/{merchantID}/authorisationsettings","json":"{\"operationId\":\"GetAuthorisationSettings\",\"parameters\":[{\"description\":\"The merchant id\",\"in\":\"path\",\"name\":\"merchantID\",\"required\":true,\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"amountLower\":{\"format\":\"double\",\"nullable\":true,\"type\":\"number\"},\"amountUpper\":{\"format\":\"double\",\"nullable\":true,\"type\":\"number\"},\"authorisationType\":{\"enum\":[\"None\",\"Payout\",\"Rule\",\"Beneficiary\",\"Payrun\",\"MerchantToken\",\"UserInvite\",\"RoleUser\"],\"type\":\"string\"},\"beneficiariesOnly\":{\"type\":\"boolean\"},\"id\":{\"format\":\"uuid\",\"type\":\"string\"},\"inserted\":{\"format\":\"date-time\",\"type\":\"string\"},\"lastEditorCantAuthorise\":{\"type\":\"boolean\"},\"lastUpdated\":{\"format\":\"date-time\",\"type\":\"string\"},\"merchantID\":{\"format\":\"uuid\",\"type\":\"string\"},\"numberOfAuthorisers\":{\"format\":\"int32\",\"type\":\"integer\"},\"roleSettings\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"description\":{\"nullable\":true,\"type\":\"string\"},\"inserted\":{\"format\":\"date-time\",\"type\":\"string\"},\"lastUpdated\":{\"format\":\"date-time\",\"type\":\"string\"},\"maxNumberAuthorisers\":{\"format\":\"int32\",\"nullable\":true,\"type\":\"integer\"},\"merchantID\":{\"format\":\"uuid\",\"type\":\"string\"},\"minNumberAuthorisers\":{\"format\":\"int32\",\"nullable\":true,\"type\":\"integer\"},\"name\":{\"nullable\":true,\"type\":\"string\"},\"roleID\":{\"format\":\"uuid\",\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"}},\"type\":\"object\"},\"type\":\"array\"}},\"text/json\":{\"schema\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"amountLower\":{\"format\":\"double\",\"nullable\":true,\"type\":\"number\"},\"amountUpper\":{\"format\":\"double\",\"nullable\":true,\"type\":\"number\"},\"authorisationType\":{\"enum\":[\"None\",\"Payout\",\"Rule\",\"Beneficiary\",\"Payrun\",\"MerchantToken\",\"UserInvite\",\"RoleUser\"],\"type\":\"string\"},\"beneficiariesOnly\":{\"type\":\"boolean\"},\"id\":{\"format\":\"uuid\",\"type\":\"string\"},\"inserted\":{\"format\":\"date-time\",\"type\":\"string\"},\"lastEditorCantAuthorise\":{\"type\":\"boolean\"},\"lastUpdated\":{\"format\":\"date-time\",\"type\":\"string\"},\"merchantID\":{\"format\":\"uuid\",\"type\":\"string\"},\"numberOfAuthorisers\":{\"format\":\"int32\",\"type\":\"integer\"},\"roleSettings\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"description\":{\"nullable\":true,\"type\":\"string\"},\"inserted\":{\"format\":\"date-time\",\"type\":\"string\"},\"lastUpdated\":{\"format\":\"date-time\",\"type\":\"string\"},\"maxNumberAuthorisers\":{\"format\":\"int32\",\"nullable\":true,\"type\":\"integer\"},\"merchantID\":{\"format\":\"uuid\",\"type\":\"string\"},\"minNumberAuthorisers\":{\"format\":\"int32\",\"nullable\":true,\"type\":\"integer\"},\"name\":{\"nullable\":true,\"type\":\"string\"},\"roleID\":{\"format\":\"uuid\",\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"}},\"type\":\"object\"},\"type\":\"array\"}},\"text/plain\":{\"schema\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"amountLower\":{\"format\":\"double\",\"nullable\":true,\"type\":\"number\"},\"amountUpper\":{\"format\":\"double\",\"nullable\":true,\"type\":\"number\"},\"authorisationType\":{\"enum\":[\"None\",\"Payout\",\"Rule\",\"Beneficiary\",\"Payrun\",\"MerchantToken\",\"UserInvite\",\"RoleUser\"],\"type\":\"string\"},\"beneficiariesOnly\":{\"type\":\"boolean\"},\"id\":{\"format\":\"uuid\",\"type\":\"string\"},\"inserted\":{\"format\":\"date-time\",\"type\":\"string\"},\"lastEditorCantAuthorise\":{\"type\":\"boolean\"},\"lastUpdated\":{\"format\":\"date-time\",\"type\":\"string\"},\"merchantID\":{\"format\":\"uuid\",\"type\":\"string\"},\"numberOfAuthorisers\":{\"format\":\"int32\",\"type\":\"integer\"},\"roleSettings\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"description\":{\"nullable\":true,\"type\":\"string\"},\"inserted\":{\"format\":\"date-time\",\"type\":\"string\"},\"lastUpdated\":{\"format\":\"date-time\",\"type\":\"string\"},\"maxNumberAuthorisers\":{\"format\":\"int32\",\"nullable\":true,\"type\":\"integer\"},\"merchantID\":{\"format\":\"uuid\",\"type\":\"string\"},\"minNumberAuthorisers\":{\"format\":\"int32\",\"nullable\":true,\"type\":\"integer\"},\"name\":{\"nullable\":true,\"type\":\"string\"},\"roleID\":{\"format\":\"uuid\",\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"}},\"type\":\"object\"},\"type\":\"array\"}}},\"description\":\"A list of merchant authorisation settings.\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/api/v1/merchants/{merchantID}/authorisationsettings","rename":{"param":{"merchantID":"merchant_id"}},"segments":[{"lit":"api"},{"lit":"v1"},{"lit":"merchants"},{"var":"merchant_id"},{"lit":"authorisationsettings"}],"select":{"exist":["merchant_id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[["merchant"]]},"key$":"merchant_authorisation_setting","name__orig":"merchant_authorisation_setting","Name":"MerchantAuthorisationSetting","name_":"merchant_authorisation_setting","name-":"merchant-authorisation-setting","NAME":"MERCHANT_AUTHORISATION_SETTING","index$":15}, {"active":true,"entity":"merchant_authorisation_setting","key$":"BasicMerchantAuthorisationSettingFlow","kind":"basic","name":"BasicMerchantAuthorisationSettingFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{"merchant_id":"merchant01"},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"merchant_authorisation_setting_ref01"}}],"index$":0}]}, 'MerchantAuthorisationSetting')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let merchant_authorisation_setting_ref01_data = Object.values(setup.data.existing.merchant_authorisation_setting)[0] as any

    // LIST
    const merchant_authorisation_setting_ref01_ent = client.MerchantAuthorisationSetting()
    const merchant_authorisation_setting_ref01_match: any = {}
    merchant_authorisation_setting_ref01_match['merchant_id'] = setup.idmap['merchant01']

    const merchant_authorisation_setting_ref01_list = (await merchant_authorisation_setting_ref01_ent.list(merchant_authorisation_setting_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/merchant_authorisation_setting/MerchantAuthorisationSettingTestData.json')

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
    ['merchant_authorisation_setting01','merchant_authorisation_setting02','merchant_authorisation_setting03','merchant01','merchant02','merchant03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'NOFRIXION_TEST_MERCHANT_AUTHORISATION_SETTING_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': '',
  })

  idmap = env['NOFRIXION_TEST_MERCHANT_AUTHORISATION_SETTING_ENTID']

  const live = 'TRUE' === env.NOFRIXION_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['NOFRIXION_TEST_MERCHANT_AUTHORISATION_SETTING_ENTID']
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
  
