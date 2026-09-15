

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


describe('FxRateEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.FxRate()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['list', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'fx_rate.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"destinationCurrency","req":false,"type":"`$STRING`","index$":0},{"active":true,"format":"double","name":"exchangeRate","req":false,"short":"The price at which the transaction will buy the source currency using the destination currency.","type":"`$NUMBER`","index$":1},{"active":true,"format":"date-time","name":"expiryTime","req":false,"type":"`$STRING`","index$":2},{"active":true,"name":"id","req":false,"type":"`$STRING`","index$":3},{"active":true,"name":"quoteID","req":false,"type":"`$STRING`","index$":4},{"active":true,"name":"sourceCurrency","req":false,"type":"`$STRING`","index$":5}],"id":{"field":"id","name":"id","parts":["source","destination","valid_for_minute"],"sep":"/"},"name":"fx_rate","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"destination","orig":"destination","reqd":true,"type":"`$STRING`","index$":0},{"active":true,"kind":"param","name":"source","orig":"source","reqd":true,"type":"`$STRING`","index$":1}]},"contract":{"id":"GET /api/v1/payouts/fxallheldrates/{source}/{destination}","json":"{\"operationId\":\"GetAllFxHeldRates\",\"parameters\":[{\"description\":\"The source currency to get the the FX held rates for.\",\"in\":\"path\",\"name\":\"source\",\"required\":true,\"schema\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"}},{\"description\":\"The destination currency to get the FX held rates for.\",\"in\":\"path\",\"name\":\"destination\",\"required\":true,\"schema\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"additionalProperties\":false,\"description\":\"Represents a Foreign Exchange (FX) rate.\",\"properties\":{\"destinationCurrency\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"exchangeRate\":{\"description\":\"The price at which the transaction will buy the source currency \\r\\nusing the destination currency.\",\"format\":\"double\",\"type\":\"number\"},\"expiryTime\":{\"format\":\"date-time\",\"type\":\"string\"},\"quoteID\":{\"nullable\":true,\"type\":\"string\"},\"sourceCurrency\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}},\"text/json\":{\"schema\":{\"items\":{\"additionalProperties\":false,\"description\":\"Represents a Foreign Exchange (FX) rate.\",\"properties\":{\"destinationCurrency\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"exchangeRate\":{\"description\":\"The price at which the transaction will buy the source currency \\r\\nusing the destination currency.\",\"format\":\"double\",\"type\":\"number\"},\"expiryTime\":{\"format\":\"date-time\",\"type\":\"string\"},\"quoteID\":{\"nullable\":true,\"type\":\"string\"},\"sourceCurrency\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}},\"text/plain\":{\"schema\":{\"items\":{\"additionalProperties\":false,\"description\":\"Represents a Foreign Exchange (FX) rate.\",\"properties\":{\"destinationCurrency\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"exchangeRate\":{\"description\":\"The price at which the transaction will buy the source currency \\r\\nusing the destination currency.\",\"format\":\"double\",\"type\":\"number\"},\"expiryTime\":{\"format\":\"date-time\",\"type\":\"string\"},\"quoteID\":{\"nullable\":true,\"type\":\"string\"},\"sourceCurrency\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}}},\"description\":\"OK\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/api/v1/payouts/fxallheldrates/{source}/{destination}","segments":[{"lit":"api"},{"lit":"v1"},{"lit":"payouts"},{"lit":"fxallheldrates"},{"var":"source"},{"var":"destination"}],"select":{"exist":["destination","source"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"destination","orig":"destination","reqd":true,"type":"`$STRING`","index$":0},{"active":true,"kind":"param","name":"source","orig":"source","reqd":true,"type":"`$STRING`","index$":1},{"active":true,"kind":"param","name":"valid_for_minute","orig":"valid_for_minute","reqd":true,"type":"`$INTEGER`","index$":2}]},"contract":{"id":"GET /api/v1/payouts/fxheldrate/{source}/{destination}/{validForMinutes}","json":"{\"operationId\":\"GetFxHeldRate\",\"parameters\":[{\"description\":\"The source currency for the FX held rate.\",\"in\":\"path\",\"name\":\"source\",\"required\":true,\"schema\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"}},{\"description\":\"The destination currency for the FX held rate.\",\"in\":\"path\",\"name\":\"destination\",\"required\":true,\"schema\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"}},{\"description\":\"The number of minutes the held rate is being requested for.\",\"in\":\"path\",\"name\":\"validForMinutes\",\"required\":true,\"schema\":{\"format\":\"int32\",\"type\":\"integer\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"description\":\"Represents a Foreign Exchange (FX) rate.\",\"properties\":{\"destinationCurrency\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"exchangeRate\":{\"description\":\"The price at which the transaction will buy the source currency \\r\\nusing the destination currency.\",\"format\":\"double\",\"type\":\"number\"},\"expiryTime\":{\"format\":\"date-time\",\"type\":\"string\"},\"quoteID\":{\"nullable\":true,\"type\":\"string\"},\"sourceCurrency\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"}},\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"description\":\"Represents a Foreign Exchange (FX) rate.\",\"properties\":{\"destinationCurrency\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"exchangeRate\":{\"description\":\"The price at which the transaction will buy the source currency \\r\\nusing the destination currency.\",\"format\":\"double\",\"type\":\"number\"},\"expiryTime\":{\"format\":\"date-time\",\"type\":\"string\"},\"quoteID\":{\"nullable\":true,\"type\":\"string\"},\"sourceCurrency\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"}},\"type\":\"object\"}},\"text/plain\":{\"schema\":{\"additionalProperties\":false,\"description\":\"Represents a Foreign Exchange (FX) rate.\",\"properties\":{\"destinationCurrency\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"exchangeRate\":{\"description\":\"The price at which the transaction will buy the source currency \\r\\nusing the destination currency.\",\"format\":\"double\",\"type\":\"number\"},\"expiryTime\":{\"format\":\"date-time\",\"type\":\"string\"},\"quoteID\":{\"nullable\":true,\"type\":\"string\"},\"sourceCurrency\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"OK\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/api/v1/payouts/fxheldrate/{source}/{destination}/{validForMinutes}","rename":{"param":{"validForMinutes":"valid_for_minute"}},"segments":[{"lit":"api"},{"lit":"v1"},{"lit":"payouts"},{"lit":"fxheldrate"},{"var":"source"},{"var":"destination"},{"var":"valid_for_minute"}],"select":{"exist":["destination","source","valid_for_minute"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[["fxallheldrate"],["fxheldrate"]]},"key$":"fx_rate","name__orig":"fx_rate","Name":"FxRate","name_":"fx_rate","name-":"fx-rate","NAME":"FX_RATE","index$":11}, {"active":true,"entity":"fx_rate","key$":"BasicFxRateFlow","kind":"basic","name":"BasicFxRateFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{"destination":"destination01","source":"source01"},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"fx_rate_ref01"}}],"index$":0},{"active":true,"data":{},"input":{"ref":"fx_rate_ref01","srcdatavar":"fx_rate_ref01_data","suffix":"_dt0"},"match":{"destination":"destination01","id":"fx_rate01","source":"source01"},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-fx_rate_ref01"}}],"index$":1}]}, 'FxRate')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let fx_rate_ref01_data = Object.values(setup.data.existing.fx_rate)[0] as any

    // LIST
    const fx_rate_ref01_ent = client.FxRate()
    const fx_rate_ref01_match: any = {}
    fx_rate_ref01_match['destination'] = setup.idmap['destination01']
    fx_rate_ref01_match['source'] = setup.idmap['source01']

    const fx_rate_ref01_list = (await fx_rate_ref01_ent.list(fx_rate_ref01_match)).map((e: any) => e.data())


    // LOAD
    const fx_rate_ref01_match_dt0: any = {}
    fx_rate_ref01_match_dt0.id = fx_rate_ref01_data.id
    const fx_rate_ref01_data_dt0 = (await fx_rate_ref01_ent.load(fx_rate_ref01_match_dt0)).data()
    assert(fx_rate_ref01_data_dt0.id === fx_rate_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/fx_rate/FxRateTestData.json')

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
    ['fx_rate01','fx_rate02','fx_rate03','fxallheldrate01','fxallheldrate02','fxallheldrate03','fxheldrate01','fxheldrate02','fxheldrate03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'NOFRIXION_TEST_FX_RATE_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': '',
  })

  idmap = env['NOFRIXION_TEST_FX_RATE_ENTID']

  const live = 'TRUE' === env.NOFRIXION_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['NOFRIXION_TEST_FX_RATE_ENTID']
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
  
