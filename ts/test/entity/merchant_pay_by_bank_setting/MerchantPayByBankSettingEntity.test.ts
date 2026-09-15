

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


describe('MerchantPayByBankSettingEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.MerchantPayByBankSetting()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'merchant_pay_by_bank_setting.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"bankCountryCodes","req":false,"short":"The list of country codes representing the banks the country supports.","type":"`$ARRAY`","index$":0},{"active":true,"format":"uuid","name":"bankID","req":false,"short":"ID of the bank to be configured for the merchant.","type":"`$STRING`","index$":1},{"active":true,"name":"bankName","req":false,"short":"Name of the Bank/Institution.","type":"`$STRING`","index$":2},{"active":true,"name":"businessInstitutionID","req":false,"short":"ID that the processor uses to identify the bank (business accounts).","type":"`$STRING`","index$":3},{"active":true,"name":"currency","req":false,"short":"Currency supported by the bank.","type":"`$STRING`","index$":4},{"active":true,"name":"logo","req":false,"short":"URL of the bank's logo.","type":"`$STRING`","index$":5},{"active":true,"name":"message","req":false,"short":"Message relating to specific bank.","type":"`$STRING`","index$":6},{"active":true,"name":"messageImageUrl","req":false,"short":"Optional image URL to be displayed with the message.","type":"`$STRING`","index$":7},{"active":true,"format":"int32","name":"order","req":false,"short":"Order in which this setting will appear in the UI.","type":"`$INTEGER`","index$":8},{"active":true,"name":"personalInstitutionID","req":false,"short":"ID that the processor uses to identify the bank (personal accounts).","type":"`$STRING`","index$":9},{"active":true,"name":"processor","req":false,"short":"Name of the bank payment processor.","type":"`$STRING`","index$":10},{"active":true,"name":"warningHeading","req":false,"short":"The heading for a warning message related to the bank institution to be displayed to the user.","type":"`$STRING`","index$":11},{"active":true,"name":"warningMessage","req":false,"short":"The warning message related to the bank institution to be displayed to the user.","type":"`$STRING`","index$":12}],"name":"merchant_pay_by_bank_setting","op":{"list":{"input":"data","name":"list","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"merchant_id","orig":"merchant_id","reqd":true,"type":"`$STRING`","index$":0}],"query":[{"active":true,"kind":"query","name":"country_code","orig":"country_code","reqd":false,"type":"`$STRING`","index$":0},{"active":true,"kind":"query","name":"currency","orig":"currency","reqd":false,"type":"`$STRING`","index$":1},{"active":true,"example":"PIS","kind":"query","name":"open_banking_operation","orig":"open_banking_operation","reqd":false,"type":"`$STRING`","index$":2}]},"contract":{"id":"GET /api/v1/merchants/{merchantID}/banksettings","json":"{\"operationId\":\"GetMerchantBankSettings\",\"parameters\":[{\"description\":\"ID of the merchant.\",\"in\":\"path\",\"name\":\"merchantID\",\"required\":true,\"schema\":{\"format\":\"uuid\",\"type\":\"string\"}},{\"description\":\"Optional. The bank currency.\",\"in\":\"query\",\"name\":\"currency\",\"schema\":{\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"}},{\"description\":\"Optional. Country code for country specific banks.\",\"in\":\"query\",\"name\":\"countryCode\",\"schema\":{\"type\":\"string\"}},{\"description\":\"The open banking operation type.\",\"in\":\"query\",\"name\":\"openBankingOperation\",\"schema\":{\"default\":\"PIS\",\"enum\":[\"None\",\"PIS\",\"AIS\"],\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"additionalProperties\":false,\"description\":\"Represents a collection of merchant bank payment settings.\",\"properties\":{\"merchantID\":{\"description\":\"Merchant to which the settings will be configured.\",\"format\":\"uuid\",\"type\":\"string\"},\"payByBankSettings\":{\"description\":\"Collection of bank payment settings.\",\"items\":{\"additionalProperties\":false,\"description\":\"Represents an individual bank payment setting.\",\"properties\":{\"bankCountryCodes\":{\"description\":\"The list of country codes representing the banks the country supports.\",\"items\":{\"type\":\"string\"},\"nullable\":true,\"type\":\"array\"},\"bankID\":{\"description\":\"ID of the bank to be configured for the merchant.\",\"format\":\"uuid\",\"type\":\"string\"},\"bankName\":{\"description\":\"Name of the Bank/Institution.\",\"nullable\":true,\"type\":\"string\"},\"businessInstitutionID\":{\"description\":\"ID that the processor uses to identify the bank (business accounts).\",\"nullable\":true,\"type\":\"string\"},\"currency\":{\"description\":\"Currency supported by the bank.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"logo\":{\"description\":\"URL of the bank's logo.\",\"nullable\":true,\"type\":\"string\"},\"message\":{\"description\":\"Message relating to specific bank.\",\"nullable\":true,\"type\":\"string\"},\"messageImageUrl\":{\"description\":\"Optional image URL to be displayed with the message.\",\"nullable\":true,\"type\":\"string\"},\"order\":{\"description\":\"Order in which this setting will appear in the UI.\",\"format\":\"int32\",\"type\":\"integer\"},\"personalInstitutionID\":{\"description\":\"ID that the processor uses to identify the bank (personal accounts).\",\"nullable\":true,\"type\":\"string\"},\"processor\":{\"description\":\"Name of the bank payment processor.\",\"enum\":[\"None\",\"CyberSource\",\"Checkout\",\"Stripe\",\"Modulr\",\"Plaid\",\"Yapily\",\"Nofrixion\",\"Bitcoin\",\"BitcoinTestnet\",\"BankingCircle\",\"BankingCircleAgency\",\"Simulator\",\"Lightning\",\"LightningTestnet\",\"BankingCircleDirectDebit\",\"Technoxander\"],\"type\":\"string\"},\"warningHeading\":{\"description\":\"The heading for a warning message related to the bank institution to be displayed to the user.\",\"nullable\":true,\"type\":\"string\"},\"warningMessage\":{\"description\":\"The warning message related to the bank institution to be displayed to the user.\",\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"}},\"type\":\"object\"}},\"text/json\":{\"schema\":{\"additionalProperties\":false,\"description\":\"Represents a collection of merchant bank payment settings.\",\"properties\":{\"merchantID\":{\"description\":\"Merchant to which the settings will be configured.\",\"format\":\"uuid\",\"type\":\"string\"},\"payByBankSettings\":{\"description\":\"Collection of bank payment settings.\",\"items\":{\"additionalProperties\":false,\"description\":\"Represents an individual bank payment setting.\",\"properties\":{\"bankCountryCodes\":{\"description\":\"The list of country codes representing the banks the country supports.\",\"items\":{\"type\":\"string\"},\"nullable\":true,\"type\":\"array\"},\"bankID\":{\"description\":\"ID of the bank to be configured for the merchant.\",\"format\":\"uuid\",\"type\":\"string\"},\"bankName\":{\"description\":\"Name of the Bank/Institution.\",\"nullable\":true,\"type\":\"string\"},\"businessInstitutionID\":{\"description\":\"ID that the processor uses to identify the bank (business accounts).\",\"nullable\":true,\"type\":\"string\"},\"currency\":{\"description\":\"Currency supported by the bank.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"logo\":{\"description\":\"URL of the bank's logo.\",\"nullable\":true,\"type\":\"string\"},\"message\":{\"description\":\"Message relating to specific bank.\",\"nullable\":true,\"type\":\"string\"},\"messageImageUrl\":{\"description\":\"Optional image URL to be displayed with the message.\",\"nullable\":true,\"type\":\"string\"},\"order\":{\"description\":\"Order in which this setting will appear in the UI.\",\"format\":\"int32\",\"type\":\"integer\"},\"personalInstitutionID\":{\"description\":\"ID that the processor uses to identify the bank (personal accounts).\",\"nullable\":true,\"type\":\"string\"},\"processor\":{\"description\":\"Name of the bank payment processor.\",\"enum\":[\"None\",\"CyberSource\",\"Checkout\",\"Stripe\",\"Modulr\",\"Plaid\",\"Yapily\",\"Nofrixion\",\"Bitcoin\",\"BitcoinTestnet\",\"BankingCircle\",\"BankingCircleAgency\",\"Simulator\",\"Lightning\",\"LightningTestnet\",\"BankingCircleDirectDebit\",\"Technoxander\"],\"type\":\"string\"},\"warningHeading\":{\"description\":\"The heading for a warning message related to the bank institution to be displayed to the user.\",\"nullable\":true,\"type\":\"string\"},\"warningMessage\":{\"description\":\"The warning message related to the bank institution to be displayed to the user.\",\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"}},\"type\":\"object\"}},\"text/plain\":{\"schema\":{\"additionalProperties\":false,\"description\":\"Represents a collection of merchant bank payment settings.\",\"properties\":{\"merchantID\":{\"description\":\"Merchant to which the settings will be configured.\",\"format\":\"uuid\",\"type\":\"string\"},\"payByBankSettings\":{\"description\":\"Collection of bank payment settings.\",\"items\":{\"additionalProperties\":false,\"description\":\"Represents an individual bank payment setting.\",\"properties\":{\"bankCountryCodes\":{\"description\":\"The list of country codes representing the banks the country supports.\",\"items\":{\"type\":\"string\"},\"nullable\":true,\"type\":\"array\"},\"bankID\":{\"description\":\"ID of the bank to be configured for the merchant.\",\"format\":\"uuid\",\"type\":\"string\"},\"bankName\":{\"description\":\"Name of the Bank/Institution.\",\"nullable\":true,\"type\":\"string\"},\"businessInstitutionID\":{\"description\":\"ID that the processor uses to identify the bank (business accounts).\",\"nullable\":true,\"type\":\"string\"},\"currency\":{\"description\":\"Currency supported by the bank.\",\"enum\":[\"NONE\",\"GBP\",\"EUR\",\"USD\",\"AUD\",\"BGN\",\"CAD\",\"CZK\",\"DKK\",\"HUF\",\"ISK\",\"CHF\",\"NOK\",\"PLN\",\"RON\",\"AED\",\"CNH\",\"HKD\",\"ILS\",\"JPY\",\"MXN\",\"NZD\",\"SAR\",\"SEK\",\"SGD\",\"TRY\",\"ZAR\",\"BTC\"],\"type\":\"string\"},\"logo\":{\"description\":\"URL of the bank's logo.\",\"nullable\":true,\"type\":\"string\"},\"message\":{\"description\":\"Message relating to specific bank.\",\"nullable\":true,\"type\":\"string\"},\"messageImageUrl\":{\"description\":\"Optional image URL to be displayed with the message.\",\"nullable\":true,\"type\":\"string\"},\"order\":{\"description\":\"Order in which this setting will appear in the UI.\",\"format\":\"int32\",\"type\":\"integer\"},\"personalInstitutionID\":{\"description\":\"ID that the processor uses to identify the bank (personal accounts).\",\"nullable\":true,\"type\":\"string\"},\"processor\":{\"description\":\"Name of the bank payment processor.\",\"enum\":[\"None\",\"CyberSource\",\"Checkout\",\"Stripe\",\"Modulr\",\"Plaid\",\"Yapily\",\"Nofrixion\",\"Bitcoin\",\"BitcoinTestnet\",\"BankingCircle\",\"BankingCircleAgency\",\"Simulator\",\"Lightning\",\"LightningTestnet\",\"BankingCircleDirectDebit\",\"Technoxander\"],\"type\":\"string\"},\"warningHeading\":{\"description\":\"The heading for a warning message related to the bank institution to be displayed to the user.\",\"nullable\":true,\"type\":\"string\"},\"warningMessage\":{\"description\":\"The warning message related to the bank institution to be displayed to the user.\",\"nullable\":true,\"type\":\"string\"}},\"type\":\"object\"},\"nullable\":true,\"type\":\"array\"}},\"type\":\"object\"}}},\"description\":\"OK\"}},\"security\":[{\"Bearer\":[]}],\"securitySchemes\":{\"Bearer\":{\"description\":\"JWT Authorization header using the Bearer scheme.<br/>\\r\\n                      Enter your JWT access token in the text input below.<br/>\\r\\n                      Example: Bearer eyJhbGciOiJ...\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/api/v1/merchants/{merchantID}/banksettings","rename":{"param":{"merchantID":"merchant_id"}},"segments":[{"lit":"api"},{"lit":"v1"},{"lit":"merchants"},{"var":"merchant_id"},{"lit":"banksettings"}],"select":{"exist":["country_code","currency","merchant_id","open_banking_operation"]},"transform":{"req":"`reqdata`","res":"`body.payByBankSettings`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[["merchant"]]},"key$":"merchant_pay_by_bank_setting","name__orig":"merchant_pay_by_bank_setting","Name":"MerchantPayByBankSetting","name_":"merchant_pay_by_bank_setting","name-":"merchant-pay-by-bank-setting","NAME":"MERCHANT_PAY_BY_BANK_SETTING","index$":17}, {"active":true,"entity":"merchant_pay_by_bank_setting","key$":"BasicMerchantPayByBankSettingFlow","kind":"basic","name":"BasicMerchantPayByBankSettingFlow","param":{},"step":[{"active":true,"data":{},"input":{},"match":{"merchant_id":"merchant01"},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"merchant_pay_by_bank_setting_ref01"}}],"index$":0}]}, 'MerchantPayByBankSetting')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let merchant_pay_by_bank_setting_ref01_data = Object.values(setup.data.existing.merchant_pay_by_bank_setting)[0] as any

    // LIST
    const merchant_pay_by_bank_setting_ref01_ent = client.MerchantPayByBankSetting()
    const merchant_pay_by_bank_setting_ref01_match: any = {}
    merchant_pay_by_bank_setting_ref01_match['merchant_id'] = setup.idmap['merchant01']

    const merchant_pay_by_bank_setting_ref01_list = (await merchant_pay_by_bank_setting_ref01_ent.list(merchant_pay_by_bank_setting_ref01_match)).map((e: any) => e.data())


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/merchant_pay_by_bank_setting/MerchantPayByBankSettingTestData.json')

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
    ['merchant_pay_by_bank_setting01','merchant_pay_by_bank_setting02','merchant_pay_by_bank_setting03','merchant01','merchant02','merchant03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'NOFRIXION_TEST_MERCHANT_PAY_BY_BANK_SETTING_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': '',
  })

  idmap = env['NOFRIXION_TEST_MERCHANT_PAY_BY_BANK_SETTING_ENTID']

  const live = 'TRUE' === env.NOFRIXION_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['NOFRIXION_TEST_MERCHANT_PAY_BY_BANK_SETTING_ENTID']
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
  
