

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


describe('CurrencyEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.Currency()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'currency.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"code","req":false,"type":"`$STRING`","index$":0},{"active":true,"format":"int32","name":"decimals","req":false,"type":"`$INTEGER`","index$":1},{"active":true,"name":"isFiat","req":false,"type":"`$BOOLEAN`","index$":2},{"active":true,"name":"iso4217AlphaCode","req":false,"type":"`$STRING`","index$":3},{"active":true,"name":"iso4217NumericCode","req":false,"type":"`$STRING`","index$":4},{"active":true,"name":"symbol","req":false,"type":"`$STRING`","index$":5}],"name":"currency","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{"query":[{"active":true,"kind":"query","name":"capability","orig":"capability","reqd":false,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /api/v1/currencies","json":"{\"operationId\":\"GetCurrencies\",\"parameters\":[{\"description\":\"Optional. When supplied, only currencies configured for the requested capability are returned.\\n\\nCapability values:\\n- `None`: No capability flags are selected. Using this value as a filter returns no currencies.\\n- `Holding`: Currency is supported for held balances and payment accounts.\\n- `Inbound`: Currency is supported for inbound payment flows.\\n- `FxConversion`: Currency is supported for FX conversion flows.\",\"in\":\"query\",\"name\":\"capability\",\"schema\":{\"description\":\"Capability values:\\n- `None`: No capability flags are selected. Using this value as a filter returns no currencies.\\n- `Holding`: Currency is supported for held balances and payment accounts.\\n- `Inbound`: Currency is supported for inbound payment flows.\\n- `FxConversion`: Currency is supported for FX conversion flows.\",\"enum\":[\"None\",\"Holding\",\"Inbound\",\"FxConversion\"],\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"code\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"decimals\":{\"format\":\"int32\",\"type\":\"integer\"},\"isFiat\":{\"type\":\"boolean\"},\"iso4217AlphaCode\":{\"nullable\":true,\"type\":\"string\"},\"iso4217NumericCode\":{\"nullable\":true,\"type\":\"string\"},\"symbol\":{\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}},\"text/json\":{\"schema\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"code\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"decimals\":{\"format\":\"int32\",\"type\":\"integer\"},\"isFiat\":{\"type\":\"boolean\"},\"iso4217AlphaCode\":{\"nullable\":true,\"type\":\"string\"},\"iso4217NumericCode\":{\"nullable\":true,\"type\":\"string\"},\"symbol\":{\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}},\"text/plain\":{\"schema\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"code\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"decimals\":{\"format\":\"int32\",\"type\":\"integer\"},\"isFiat\":{\"type\":\"boolean\"},\"iso4217AlphaCode\":{\"nullable\":true,\"type\":\"string\"},\"iso4217NumericCode\":{\"nullable\":true,\"type\":\"string\"},\"symbol\":{\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}}},\"description\":\"Returns the list of currencies and their static metadata.\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/api/v1/currencies","segments":[{"lit":"api"},{"lit":"v1"},{"lit":"currencies"}],"select":{"exist":["capability"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[]},"key$":"currency","name__orig":"currency","Name":"Currency","name_":"currency","name-":"currency","NAME":"CURRENCY","index$":9}, {"active":true,"entity":"currency","key$":"BasicCurrencyFlow","kind":"basic","name":"BasicCurrencyFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"currency_ref01"}}],"index$":0}]}, 'Currency')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let currency_ref01_data = Object.values(setup.data.existing.currency)[0] as any

    // LIST
    const currency_ref01_ent = client.Currency()
    const currency_ref01_match: any = {}

    const currency_ref01_list = (await currency_ref01_ent.list(currency_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/currency/CurrencyTestData.json')

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
    ['currency01','currency02','currency03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'NOFRIXION_TEST_CURRENCY_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': '',
  })

  idmap = env['NOFRIXION_TEST_CURRENCY_ENTID']

  const live = 'TRUE' === env.NOFRIXION_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['NOFRIXION_TEST_CURRENCY_ENTID']
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
  
