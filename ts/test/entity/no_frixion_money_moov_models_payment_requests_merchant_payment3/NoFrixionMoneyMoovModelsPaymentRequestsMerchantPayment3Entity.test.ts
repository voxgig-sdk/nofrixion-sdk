
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'


import { NofrixionSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveDelay,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


describe('NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['update']) {
      if (maybeSkipControl(t, 'entityOp', 'no_frixion_money_moov_models_payment_requests_merchant_payment3.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data = Object.values(setup.data.existing.no_frixion_money_moov_models_payment_requests_merchant_payment3)[0] as any

    // UPDATE
    const no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_ent = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3()
    const no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data_up0: any = {}
    no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data_up0.id = no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data.id
    no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data_up0 ['paymentrequest_id'] = setup.idmap['paymentrequest_id']

    const no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_markdef_up0 = { name: 'description', value: 'Mark01-no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_' + setup.now }
    ;(no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data_up0 as any)[no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_markdef_up0.name] = no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_markdef_up0.value

    const no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_resdata_up0 = await no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_ent.update(no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data_up0)
    assert(no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_resdata_up0.id === no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_data_up0.id)

    assert((no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_resdata_up0 as any)[no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_markdef_up0.name] === no_frixion_money_moov_models_payment_requests_merchant_payment3_ref01_markdef_up0.value)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/no_frixion_money_moov_models_payment_requests_merchant_payment3/NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3TestData.json')

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
    ['no_frixion_money_moov_models_payment_requests_merchant_payment301','no_frixion_money_moov_models_payment_requests_merchant_payment302','no_frixion_money_moov_models_payment_requests_merchant_payment303','paymentrequest01','paymentrequest02','paymentrequest03','template01','template02','template03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  // Detect whether the user provided a real ENTID JSON via env var. The
  // basic flow consumes synthetic IDs from the fixture file; without an
  // override those synthetic IDs reach the live API and 4xx. Surface this
  // to the test so it can skip rather than fail.
  const idmapEnvVal = process.env['NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': 'NONE',
  })

  idmap = env['NOFRIXION_TEST_NO_FRIXION_MONEY_MOOV_MODELS_PAYMENT_REQUESTS_MERCHANT_PAYMENT__ENTID']

  const live = 'TRUE' === env.NOFRIXION_TEST_LIVE

  if (live) {
    client = new NofrixionSDK(merge([
      {
        apikey: env.NOFRIXION_APIKEY,
      },
      extra
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
    syntheticOnly: live && !idmapOverridden,
    now: Date.now(),
  }

  return setup
}
  
