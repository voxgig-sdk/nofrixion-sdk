

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


describe('PayoutMetricEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.PayoutMetric()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'payout_metric.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[],"name":"payout_metric","op":{"load":{"input":"data","name":"load","points":[{"active":true,"args":{"query":[{"active":true,"kind":"query","name":"currency","orig":"currency","reqd":false,"type":"`$STRING`","index$":0},{"active":true,"kind":"query","name":"from_date","orig":"from_date","reqd":false,"type":"`$STRING`","index$":1},{"active":true,"example":false,"kind":"query","name":"include_archived","orig":"include_archived","reqd":false,"type":"`$BOOLEAN`","index$":2},{"active":true,"kind":"query","name":"max_amount","orig":"max_amount","reqd":false,"type":"`$NUMBER`","index$":3},{"active":true,"kind":"query","name":"merchant_id","orig":"merchant_id","reqd":false,"type":"`$STRING`","index$":4},{"active":true,"kind":"query","name":"min_amount","orig":"min_amount","reqd":false,"type":"`$NUMBER`","index$":5},{"active":true,"kind":"query","name":"search","orig":"search","reqd":false,"type":"`$STRING`","index$":6},{"active":true,"kind":"query","name":"tag","orig":"tag","reqd":false,"type":"`$ARRAY`","index$":7},{"active":true,"kind":"query","name":"to_date","orig":"to_date","reqd":false,"type":"`$STRING`","index$":8}]},"contract":{"id":"GET /api/v1/payouts/metrics","json":"{\"operationId\":\"GetPayoutMetrics\",\"parameters\":[{\"description\":\"Required. The ID of the merchant to get the payout metrics for.\",\"in\":\"query\",\"name\":\"merchantID\",\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}},{\"description\":\"The date filter to apply to retrieve payouts created after this date as metrics.\",\"in\":\"query\",\"name\":\"fromDate\",\"schema\":{\"format\":\"date-time\",\"type\":\"string\"}},{\"description\":\"The date filter to apply to retrieve payouts created up until this date as metrics.\",\"in\":\"query\",\"name\":\"toDate\",\"schema\":{\"format\":\"date-time\",\"type\":\"string\"}},{\"description\":\"The text filter to apply to retrieve payouts with a similar title, description, merchant name or contact information as metrics.\",\"in\":\"query\",\"name\":\"search\",\"schema\":{\"type\":\"string\"}},{\"description\":\"The currency filter to apply to retrieve payouts with this currency as metrics.\",\"in\":\"query\",\"name\":\"currency\",\"schema\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"}},{\"description\":\"The amount filter to apply to retrieve payouts that exceed this amount as metrics.\",\"in\":\"query\",\"name\":\"minAmount\",\"schema\":{\"format\":\"double\",\"type\":\"number\"}},{\"description\":\"The amount filter to apply to retrieve payouts that don't exceed this amount as metrics.\",\"in\":\"query\",\"name\":\"maxAmount\",\"schema\":{\"format\":\"double\",\"type\":\"number\"}},{\"description\":\"The tag filter to apply to retrieve payouts with at least one of these tags as metrics.\",\"in\":\"query\",\"name\":\"tags\",\"schema\":{\"items\":{\"format\":\"uuid\",\"type\":\"string\"},\"type\":\"array\"}},{\"description\":\"Flag that indicates whether to include archived payouts in the metrics.\",\"in\":\"query\",\"name\":\"includeArchived\",\"schema\":{\"default\":false,\"type\":\"boolean\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"all\":{\"description\":\"Total payout count.\",\"format\":\"double\",\"type\":\"number\"},\"failed\":{\"description\":\"Payouts with Failed, Rejected or Unknown status.\",\"format\":\"double\",\"type\":\"number\"},\"inProgress\":{\"description\":\"Payouts with Pending, Queued or QueuedUpstream status.\",\"format\":\"double\",\"type\":\"number\"},\"paid\":{\"description\":\"Payouts with Processed status.\",\"format\":\"double\",\"type\":\"number\"},\"pendingApproval\":{\"description\":\"Payouts with PendingApproval or PendingInput status.\",\"format\":\"double\",\"type\":\"number\"},\"scheduled\":{\"description\":\"Payouts with Scheduled status.\",\"format\":\"double\",\"type\":\"number\"},\"totalAmountsByCurrency\":{\"additionalProperties\":{\"additionalProperties\":{\"format\":\"double\",\"type\":\"number\"},\"type\":\"object\"},\"description\":\"The total amounts by status and currency.\",\"nullable\":true,\"type\":\"object\"}},\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"all\":{\"description\":\"Total payout count.\",\"format\":\"double\",\"type\":\"number\"},\"failed\":{\"description\":\"Payouts with Failed, Rejected or Unknown status.\",\"format\":\"double\",\"type\":\"number\"},\"inProgress\":{\"description\":\"Payouts with Pending, Queued or QueuedUpstream status.\",\"format\":\"double\",\"type\":\"number\"},\"paid\":{\"description\":\"Payouts with Processed status.\",\"format\":\"double\",\"type\":\"number\"},\"pendingApproval\":{\"description\":\"Payouts with PendingApproval or PendingInput status.\",\"format\":\"double\",\"type\":\"number\"},\"scheduled\":{\"description\":\"Payouts with Scheduled status.\",\"format\":\"double\",\"type\":\"number\"},\"totalAmountsByCurrency\":{\"additionalProperties\":{\"additionalProperties\":{\"format\":\"double\",\"type\":\"number\"},\"type\":\"object\"},\"description\":\"The total amounts by status and currency.\",\"nullable\":true,\"type\":\"object\"}},\"type\":\"object\"}},\"text/plain\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"all\":{\"description\":\"Total payout count.\",\"format\":\"double\",\"type\":\"number\"},\"failed\":{\"description\":\"Payouts with Failed, Rejected or Unknown status.\",\"format\":\"double\",\"type\":\"number\"},\"inProgress\":{\"description\":\"Payouts with Pending, Queued or QueuedUpstream status.\",\"format\":\"double\",\"type\":\"number\"},\"paid\":{\"description\":\"Payouts with Processed status.\",\"format\":\"double\",\"type\":\"number\"},\"pendingApproval\":{\"description\":\"Payouts with PendingApproval or PendingInput status.\",\"format\":\"double\",\"type\":\"number\"},\"scheduled\":{\"description\":\"Payouts with Scheduled status.\",\"format\":\"double\",\"type\":\"number\"},\"totalAmountsByCurrency\":{\"additionalProperties\":{\"additionalProperties\":{\"format\":\"double\",\"type\":\"number\"},\"type\":\"object\"},\"description\":\"The total amounts by status and currency.\",\"nullable\":true,\"type\":\"object\"}},\"type\":\"object\"}}},\"description\":\"OK\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/api/v1/payouts/metrics","segments":[{"lit":"api"},{"lit":"v1"},{"lit":"payouts"},{"lit":"metrics"}],"select":{"exist":["currency","from_date","include_archived","max_amount","merchant_id","min_amount","search","tag","to_date"]},"transform":{"req":"`reqdata`","res":"`body.totalAmountsByCurrency`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"payout_metric","name__orig":"payout_metric","Name":"PayoutMetric","name_":"payout_metric","name-":"payout-metric","NAME":"PAYOUT_METRIC","index$":35}, {"active":true,"entity":"payout_metric","key$":"BasicPayoutMetricFlow","kind":"basic","name":"BasicPayoutMetricFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"payout_metric_ref01","srcdatavar":"payout_metric_ref01_data","suffix":"_dt0"},"match":{},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-payout_metric_ref01"}}],"index$":0}]}, 'PayoutMetric')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let payout_metric_ref01_data = Object.values(setup.data.existing.payout_metric)[0] as any

    // LOAD
    const payout_metric_ref01_ent = client.PayoutMetric()
    const payout_metric_ref01_match_dt0: any = {}
    const payout_metric_ref01_data_dt0 = (await payout_metric_ref01_ent.load(payout_metric_ref01_match_dt0)).data()
    assert(null != payout_metric_ref01_data_dt0)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/payout_metric/PayoutMetricTestData.json')

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
    ['payout_metric01','payout_metric02','payout_metric03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'NOFRIXION_TEST_PAYOUT_METRIC_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': '',
  })

  idmap = env['NOFRIXION_TEST_PAYOUT_METRIC_ENTID']

  const live = 'TRUE' === env.NOFRIXION_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['NOFRIXION_TEST_PAYOUT_METRIC_ENTID']
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
  
