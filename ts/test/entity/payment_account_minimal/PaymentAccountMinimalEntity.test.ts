

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


describe('PaymentAccountMinimalEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.PaymentAccountMinimal()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'payment_account_minimal.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"accountName","req":false,"short":"Name for the account","type":"`$STRING`","index$":0},{"active":true,"format":"double","name":"availableBalance","readOnly":true,"req":false,"short":"The current available balance of the account.","type":"`$NUMBER`","index$":1},{"active":true,"format":"double","name":"balance","req":false,"short":"Balance of the account.","type":"`$NUMBER`","index$":2},{"active":true,"format":"int64","name":"balanceMinorUnits","readOnly":true,"req":false,"short":"Balance of the account expressed in the currency’s minor units (e.g.","type":"`$INTEGER`","index$":3},{"active":true,"name":"currency","req":false,"short":"Currency of the account in ISO 4217 format","type":"`$STRING`","index$":4},{"active":true,"format":"uuid","name":"id","req":false,"short":"Unique id for the account.","type":"`$STRING`","index$":5},{"active":true,"name":"identifier","req":true,"type":"`$OBJECT`","index$":6},{"active":true,"name":"isArchived","req":false,"short":"Is the account archived","type":"`$BOOLEAN`","index$":7},{"active":true,"name":"isConnectedAccount","req":false,"short":"Indicates if the payment account is an externally connected account.","type":"`$BOOLEAN`","index$":8},{"active":true,"format":"uuid","name":"merchantID","req":false,"short":"The ID of the merchant that owns the account.","type":"`$STRING`","index$":9},{"active":true,"format":"double","name":"submittedPayoutsBalance","req":false,"short":"Total of the payouts that have been submitted for processing.","type":"`$NUMBER`","index$":10}],"id":{"field":"id","name":"id"},"name":"payment_account_minimal","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{"query":[{"active":true,"example":false,"kind":"query","name":"connected_account","orig":"connected_account","reqd":false,"type":"`$BOOLEAN`","index$":0},{"active":true,"kind":"query","name":"currency","orig":"currency","reqd":false,"type":"`$ARRAY`","index$":1},{"active":true,"example":false,"kind":"query","name":"include_archived","orig":"include_archived","reqd":false,"type":"`$BOOLEAN`","index$":2},{"active":true,"kind":"query","name":"merchant_id","orig":"merchant_id","reqd":false,"type":"`$STRING`","index$":3},{"active":true,"example":false,"kind":"query","name":"only_archived","orig":"only_archived","reqd":false,"type":"`$BOOLEAN`","index$":4},{"active":true,"example":false,"kind":"query","name":"only_connect_account","orig":"only_connect_account","reqd":false,"type":"`$BOOLEAN`","index$":5},{"active":true,"example":1,"kind":"query","name":"page_number","orig":"page_number","reqd":false,"type":"`$INTEGER`","index$":6},{"active":true,"example":10,"kind":"query","name":"page_size","orig":"page_size","reqd":false,"type":"`$INTEGER`","index$":7},{"active":true,"kind":"query","name":"search","orig":"search","reqd":false,"type":"`$STRING`","index$":8},{"active":true,"kind":"query","name":"sort","orig":"sort","reqd":false,"type":"`$STRING`","index$":9}]},"contract":{"id":"GET /api/v1/accounts/minimal","json":"{\"operationId\":\"GetAccountsMinimal\",\"parameters\":[{\"description\":\"The merchantID of the accounts to retrieve.\",\"in\":\"query\",\"name\":\"merchantID\",\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}},{\"description\":\"Optional. Array. If specified will only return accounts for these currencies.\",\"in\":\"query\",\"name\":\"currency\",\"schema\":{\"items\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"type\":\"array\"}},{\"description\":\"Optional include connected accounts along with payment accounts.\",\"in\":\"query\",\"name\":\"connectedAccounts\",\"schema\":{\"default\":false,\"type\":\"boolean\"}},{\"description\":\"Optional. The page number to retrieve.\",\"in\":\"query\",\"name\":\"pageNumber\",\"schema\":{\"default\":1,\"format\":\"int32\",\"type\":\"integer\"}},{\"description\":\"Optional. The number of accounts per page.\",\"in\":\"query\",\"name\":\"pageSize\",\"schema\":{\"default\":10,\"format\":\"int32\",\"type\":\"integer\"}},{\"description\":\"The text filter to apply to retrieve accounts with a similar account name, IBAN etc.\",\"in\":\"query\",\"name\":\"search\",\"schema\":{\"type\":\"string\"}},{\"description\":\"Optional expression to sort the order of the accounts. Example \\\"AvailableBalance desc,Inserted asc\\\".\",\"in\":\"query\",\"name\":\"sort\",\"schema\":{\"type\":\"string\"}},{\"description\":\"Only return connected accounts\",\"in\":\"query\",\"name\":\"onlyConnectAccounts\",\"schema\":{\"default\":false,\"type\":\"boolean\"}},{\"description\":\"Flag that indicates whether to fetch only archived accounts or not.\",\"in\":\"query\",\"name\":\"onlyArchived\",\"schema\":{\"default\":false,\"type\":\"boolean\"}},{\"description\":\"Flag that indicates whether to fetch archived accounts or not.\",\"in\":\"query\",\"name\":\"includeArchived\",\"schema\":{\"default\":false,\"type\":\"boolean\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"content\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"accountName\":{\"description\":\"Name for the account\",\"nullable\":true,\"type\":\"string\"},\"availableBalance\":{\"description\":\"The current available balance of the account. Calculated by subtracting any submitted payments from the current balance.\",\"format\":\"double\",\"readOnly\":true,\"type\":\"number\"},\"balance\":{\"description\":\"Balance of the account.\",\"format\":\"double\",\"type\":\"number\"},\"balanceMinorUnits\":{\"description\":\"Balance of the account expressed in the currency’s minor units (e.g. cents, pence).\",\"format\":\"int64\",\"readOnly\":true,\"type\":\"integer\"},\"currency\":{\"description\":\"Currency of the account in ISO 4217 format\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"id\":{\"description\":\"Unique id for the account.\",\"format\":\"uuid\",\"type\":\"string\"},\"identifier\":{\"additionalProperties\":false,\"properties\":{\"accountNumber\":{\"nullable\":true,\"type\":\"string\"},\"bic\":{\"nullable\":true,\"type\":\"string\"},\"currency\":{\"description\":\"The currency for the account.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"displayBicSummary\":{\"nullable\":true,\"readOnly\":true,\"type\":\"string\"},\"displayScanSummary\":{\"nullable\":true,\"readOnly\":true,\"type\":\"string\"},\"displaySummary\":{\"description\":\"Summary of the account identifier's most important properties.\",\"nullable\":true,\"readOnly\":true,\"type\":\"string\"},\"iban\":{\"nullable\":true,\"type\":\"string\"},\"sortCode\":{\"nullable\":true,\"type\":\"string\"},\"summary\":{\"description\":\"Summary of the account identifier's most important properties.\",\"nullable\":true,\"readOnly\":true,\"type\":\"string\"},\"type\":{\"description\":\"The type of the account identifier.\",\"enum\":[\"Unknown\",\"SCAN\",\"IBAN\",\"DD\",\"BTC\",\"BIC\"],\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"currency\"],\"type\":\"object\"},\"isArchived\":{\"description\":\"Is the account archived\",\"type\":\"boolean\"},\"isConnectedAccount\":{\"description\":\"Indicates if the payment account is an externally connected account.\\r\\nExternally connected account can be used to view account balances and transactions.\",\"type\":\"boolean\"},\"merchantID\":{\"description\":\"The ID of the merchant that owns the account.\",\"format\":\"uuid\",\"type\":\"string\"},\"submittedPayoutsBalance\":{\"description\":\"Total of the payouts that have been submitted for processing.\",\"format\":\"double\",\"type\":\"number\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"},\"pageNumber\":{\"description\":\"Current page number. Its 1 based. i.e firstpage is 1, secondpage is 2\",\"format\":\"int32\",\"type\":\"integer\"},\"pageSize\":{\"description\":\"Page size\",\"format\":\"int32\",\"type\":\"integer\"},\"totalPages\":{\"description\":\"Total pages\",\"format\":\"int64\",\"type\":\"integer\"},\"totalSize\":{\"description\":\"Total count\",\"format\":\"int64\",\"type\":\"integer\"}},\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"content\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"accountName\":{\"description\":\"Name for the account\",\"nullable\":true,\"type\":\"string\"},\"availableBalance\":{\"description\":\"The current available balance of the account. Calculated by subtracting any submitted payments from the current balance.\",\"format\":\"double\",\"readOnly\":true,\"type\":\"number\"},\"balance\":{\"description\":\"Balance of the account.\",\"format\":\"double\",\"type\":\"number\"},\"balanceMinorUnits\":{\"description\":\"Balance of the account expressed in the currency’s minor units (e.g. cents, pence).\",\"format\":\"int64\",\"readOnly\":true,\"type\":\"integer\"},\"currency\":{\"description\":\"Currency of the account in ISO 4217 format\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"id\":{\"description\":\"Unique id for the account.\",\"format\":\"uuid\",\"type\":\"string\"},\"identifier\":{\"additionalProperties\":false,\"properties\":{\"accountNumber\":{\"nullable\":true,\"type\":\"string\"},\"bic\":{\"nullable\":true,\"type\":\"string\"},\"currency\":{\"description\":\"The currency for the account.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"displayBicSummary\":{\"nullable\":true,\"readOnly\":true,\"type\":\"string\"},\"displayScanSummary\":{\"nullable\":true,\"readOnly\":true,\"type\":\"string\"},\"displaySummary\":{\"description\":\"Summary of the account identifier's most important properties.\",\"nullable\":true,\"readOnly\":true,\"type\":\"string\"},\"iban\":{\"nullable\":true,\"type\":\"string\"},\"sortCode\":{\"nullable\":true,\"type\":\"string\"},\"summary\":{\"description\":\"Summary of the account identifier's most important properties.\",\"nullable\":true,\"readOnly\":true,\"type\":\"string\"},\"type\":{\"description\":\"The type of the account identifier.\",\"enum\":[\"Unknown\",\"SCAN\",\"IBAN\",\"DD\",\"BTC\",\"BIC\"],\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"currency\"],\"type\":\"object\"},\"isArchived\":{\"description\":\"Is the account archived\",\"type\":\"boolean\"},\"isConnectedAccount\":{\"description\":\"Indicates if the payment account is an externally connected account.\\r\\nExternally connected account can be used to view account balances and transactions.\",\"type\":\"boolean\"},\"merchantID\":{\"description\":\"The ID of the merchant that owns the account.\",\"format\":\"uuid\",\"type\":\"string\"},\"submittedPayoutsBalance\":{\"description\":\"Total of the payouts that have been submitted for processing.\",\"format\":\"double\",\"type\":\"number\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"},\"pageNumber\":{\"description\":\"Current page number. Its 1 based. i.e firstpage is 1, secondpage is 2\",\"format\":\"int32\",\"type\":\"integer\"},\"pageSize\":{\"description\":\"Page size\",\"format\":\"int32\",\"type\":\"integer\"},\"totalPages\":{\"description\":\"Total pages\",\"format\":\"int64\",\"type\":\"integer\"},\"totalSize\":{\"description\":\"Total count\",\"format\":\"int64\",\"type\":\"integer\"}},\"type\":\"object\"}},\"text/plain\":{\"schema\":{\"additionalProperties\":false,\"properties\":{\"content\":{\"items\":{\"additionalProperties\":false,\"properties\":{\"accountName\":{\"description\":\"Name for the account\",\"nullable\":true,\"type\":\"string\"},\"availableBalance\":{\"description\":\"The current available balance of the account. Calculated by subtracting any submitted payments from the current balance.\",\"format\":\"double\",\"readOnly\":true,\"type\":\"number\"},\"balance\":{\"description\":\"Balance of the account.\",\"format\":\"double\",\"type\":\"number\"},\"balanceMinorUnits\":{\"description\":\"Balance of the account expressed in the currency’s minor units (e.g. cents, pence).\",\"format\":\"int64\",\"readOnly\":true,\"type\":\"integer\"},\"currency\":{\"description\":\"Currency of the account in ISO 4217 format\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"id\":{\"description\":\"Unique id for the account.\",\"format\":\"uuid\",\"type\":\"string\"},\"identifier\":{\"additionalProperties\":false,\"properties\":{\"accountNumber\":{\"nullable\":true,\"type\":\"string\"},\"bic\":{\"nullable\":true,\"type\":\"string\"},\"currency\":{\"description\":\"The currency for the account.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"displayBicSummary\":{\"nullable\":true,\"readOnly\":true,\"type\":\"string\"},\"displayScanSummary\":{\"nullable\":true,\"readOnly\":true,\"type\":\"string\"},\"displaySummary\":{\"description\":\"Summary of the account identifier's most important properties.\",\"nullable\":true,\"readOnly\":true,\"type\":\"string\"},\"iban\":{\"nullable\":true,\"type\":\"string\"},\"sortCode\":{\"nullable\":true,\"type\":\"string\"},\"summary\":{\"description\":\"Summary of the account identifier's most important properties.\",\"nullable\":true,\"readOnly\":true,\"type\":\"string\"},\"type\":{\"description\":\"The type of the account identifier.\",\"enum\":[\"Unknown\",\"SCAN\",\"IBAN\",\"DD\",\"BTC\",\"BIC\"],\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"currency\"],\"type\":\"object\"},\"isArchived\":{\"description\":\"Is the account archived\",\"type\":\"boolean\"},\"isConnectedAccount\":{\"description\":\"Indicates if the payment account is an externally connected account.\\r\\nExternally connected account can be used to view account balances and transactions.\",\"type\":\"boolean\"},\"merchantID\":{\"description\":\"The ID of the merchant that owns the account.\",\"format\":\"uuid\",\"type\":\"string\"},\"submittedPayoutsBalance\":{\"description\":\"Total of the payouts that have been submitted for processing.\",\"format\":\"double\",\"type\":\"number\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"},\"pageNumber\":{\"description\":\"Current page number. Its 1 based. i.e firstpage is 1, secondpage is 2\",\"format\":\"int32\",\"type\":\"integer\"},\"pageSize\":{\"description\":\"Page size\",\"format\":\"int32\",\"type\":\"integer\"},\"totalPages\":{\"description\":\"Total pages\",\"format\":\"int64\",\"type\":\"integer\"},\"totalSize\":{\"description\":\"Total count\",\"format\":\"int64\",\"type\":\"integer\"}},\"type\":\"object\"}}},\"description\":\"Returns the list of accounts for the merchant.\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/api/v1/accounts/minimal","segments":[{"lit":"api"},{"lit":"v1"},{"lit":"accounts"},{"lit":"minimal"}],"select":{"exist":["connected_account","currency","include_archived","merchant_id","only_archived","only_connect_account","page_number","page_size","search","sort"]},"transform":{"req":"`reqdata`","res":"`body.content`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[]},"key$":"payment_account_minimal","name__orig":"payment_account_minimal","Name":"PaymentAccountMinimal","name_":"payment_account_minimal","name-":"payment-account-minimal","NAME":"PAYMENT_ACCOUNT_MINIMAL","index$":26}, {"active":true,"entity":"payment_account_minimal","key$":"BasicPaymentAccountMinimalFlow","kind":"basic","name":"BasicPaymentAccountMinimalFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"payment_account_minimal_ref01"}}],"index$":0}]}, 'PaymentAccountMinimal')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let payment_account_minimal_ref01_data = Object.values(setup.data.existing.payment_account_minimal)[0] as any

    // LIST
    const payment_account_minimal_ref01_ent = client.PaymentAccountMinimal()
    const payment_account_minimal_ref01_match: any = {}

    const payment_account_minimal_ref01_list = (await payment_account_minimal_ref01_ent.list(payment_account_minimal_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/payment_account_minimal/PaymentAccountMinimalTestData.json')

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
    ['payment_account_minimal01','payment_account_minimal02','payment_account_minimal03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'NOFRIXION_TEST_PAYMENT_ACCOUNT_MINIMAL_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': '',
  })

  idmap = env['NOFRIXION_TEST_PAYMENT_ACCOUNT_MINIMAL_ENTID']

  const live = 'TRUE' === env.NOFRIXION_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['NOFRIXION_TEST_PAYMENT_ACCOUNT_MINIMAL_ENTID']
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
  
