

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


describe('ReportResultEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.ReportResult()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'report_result.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"contentType","readOnly":true,"req":false,"type":"`$STRING`","index$":0},{"active":true,"name":"contents","req":false,"type":"`$STRING`","index$":1},{"active":true,"name":"id","req":false,"type":"`$STRING`","index$":2},{"active":true,"format":"date-time","name":"lastCompletedAt","req":false,"type":"`$STRING`","index$":3},{"active":true,"format":"uuid","name":"merchantID","req":false,"type":"`$STRING`","index$":4},{"active":true,"name":"reportName","req":false,"type":"`$STRING`","index$":5},{"active":true,"name":"reportType","req":false,"type":"`$STRING`","index$":6},{"active":true,"format":"int32","name":"statementNumber","req":false,"type":"`$INTEGER`","index$":7}],"id":{"field":"id","name":"id"},"name":"report_result","op":{"load":{"input":"data","name":"load","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"id","orig":"statement_number","reqd":true,"type":"`$INTEGER`","index$":0},{"active":true,"kind":"param","name":"report_id","orig":"id","reqd":true,"type":"`$STRING`","index$":1}]},"contract":{"id":"GET /api/v1/reports/{id}/result/{statementNumber}","json":"{\"operationId\":\"GetReportResult\",\"parameters\":[{\"description\":\"The ID of the merchant report to get the result for.\",\"in\":\"path\",\"name\":\"id\",\"required\":true,\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}},{\"description\":\"Optional. If set the result for this statement number \\r\\n            will be returned. If not set the latest report will be returned.\",\"in\":\"path\",\"name\":\"statementNumber\",\"required\":true,\"schema\":{\"format\":\"int32\",\"type\":\"integer\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"contentType\":{\"nullable\":true,\"readOnly\":true,\"type\":\"string\"},\"contents\":{\"nullable\":true,\"type\":\"string\"},\"lastCompletedAt\":{\"format\":\"date-time\",\"type\":\"string\"},\"merchantID\":{\"format\":\"uuid\",\"nullable\":true,\"type\":\"string\"},\"reportName\":{\"nullable\":true,\"type\":\"string\"},\"reportType\":{\"enum\":[\"Unknown\",\"SwiftCustomerStatement\",\"CustomerActivity\",\"SafeGuardingReconciliation\",\"MerchantAccountsBalance\",\"MerchantAccountsTransaction\",\"VisionBlueTransaction\",\"MerchantSafeGuardingReconciliation\"],\"type\":\"string\"},\"statementNumber\":{\"format\":\"int32\",\"type\":\"integer\"}},\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"contentType\":{\"nullable\":true,\"readOnly\":true,\"type\":\"string\"},\"contents\":{\"nullable\":true,\"type\":\"string\"},\"lastCompletedAt\":{\"format\":\"date-time\",\"type\":\"string\"},\"merchantID\":{\"format\":\"uuid\",\"nullable\":true,\"type\":\"string\"},\"reportName\":{\"nullable\":true,\"type\":\"string\"},\"reportType\":{\"enum\":[\"Unknown\",\"SwiftCustomerStatement\",\"CustomerActivity\",\"SafeGuardingReconciliation\",\"MerchantAccountsBalance\",\"MerchantAccountsTransaction\",\"VisionBlueTransaction\",\"MerchantSafeGuardingReconciliation\"],\"type\":\"string\"},\"statementNumber\":{\"format\":\"int32\",\"type\":\"integer\"}},\"type\":\"object\"}},\"text/plain\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"contentType\":{\"nullable\":true,\"readOnly\":true,\"type\":\"string\"},\"contents\":{\"nullable\":true,\"type\":\"string\"},\"lastCompletedAt\":{\"format\":\"date-time\",\"type\":\"string\"},\"merchantID\":{\"format\":\"uuid\",\"nullable\":true,\"type\":\"string\"},\"reportName\":{\"nullable\":true,\"type\":\"string\"},\"reportType\":{\"enum\":[\"Unknown\",\"SwiftCustomerStatement\",\"CustomerActivity\",\"SafeGuardingReconciliation\",\"MerchantAccountsBalance\",\"MerchantAccountsTransaction\",\"VisionBlueTransaction\",\"MerchantSafeGuardingReconciliation\"],\"type\":\"string\"},\"statementNumber\":{\"format\":\"int32\",\"type\":\"integer\"}},\"type\":\"object\"}}},\"description\":\"OK\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/api/v1/reports/{id}/result/{statementNumber}","rename":{"param":{"id":"report_id","statementNumber":"id"}},"segments":[{"lit":"api"},{"lit":"v1"},{"lit":"reports"},{"var":"report_id"},{"lit":"result"},{"var":"id"}],"select":{"exist":["id","report_id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[["report"]]},"key$":"report_result","name__orig":"report_result","Name":"ReportResult","name_":"report_result","name-":"report-result","NAME":"REPORT_RESULT","index$":38}, {"active":true,"entity":"report_result","key$":"BasicReportResultFlow","kind":"basic","name":"BasicReportResultFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"report_result_ref01","srcdatavar":"report_result_ref01_data","suffix":"_dt0"},"match":{"id":"report_result01","report_id":"report01"},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-report_result_ref01"}}],"index$":0}]}, 'ReportResult')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let report_result_ref01_data = Object.values(setup.data.existing.report_result)[0] as any

    // LOAD
    const report_result_ref01_ent = client.ReportResult()
    const report_result_ref01_match_dt0: any = {}
    report_result_ref01_match_dt0.id = report_result_ref01_data.id
    const report_result_ref01_data_dt0 = (await report_result_ref01_ent.load(report_result_ref01_match_dt0)).data()
    assert(report_result_ref01_data_dt0.id === report_result_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/report_result/ReportResultTestData.json')

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
    ['report_result01','report_result02','report_result03','report01','report02','report03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'NOFRIXION_TEST_REPORT_RESULT_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': '',
  })

  idmap = env['NOFRIXION_TEST_REPORT_RESULT_ENTID']

  const live = 'TRUE' === env.NOFRIXION_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['NOFRIXION_TEST_REPORT_RESULT_ENTID']
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
  
