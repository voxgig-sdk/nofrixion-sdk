
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


describe('NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['load']) {
      if (maybeSkipControl(t, 'entityOp', 'no_frixion_money_moov_models_payment_requests_merchant_payment2.' + op, live)) return
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

    let no_frixion_money_moov_models_payment_requests_merchant_payment2_ref01_data = Object.values(setup.data.existing.no_frixion_money_moov_models_payment_requests_merchant_payment2)[0] as any

    // LOAD
    const no_frixion_money_moov_models_payment_requests_merchant_payment2_ref01_ent = client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2()
    const no_frixion_money_moov_models_payment_requests_merchant_payment2_ref01_match_dt0: any = {}
    no_frixion_money_moov_models_payment_requests_merchant_payment2_ref01_match_dt0.id = no_frixion_money_moov_models_payment_requests_merchant_payment2_ref01_data.id
    const no_frixion_money_moov_models_payment_requests_merchant_payment2_ref01_data_dt0 = await no_frixion_money_moov_models_payment_requests_merchant_payment2_ref01_ent.load(no_frixion_money_moov_models_payment_requests_merchant_payment2_ref01_match_dt0)
    assert(no_frixion_money_moov_models_payment_requests_merchant_payment2_ref01_data_dt0.id === no_frixion_money_moov_models_payment_requests_merchant_payment2_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/no_frixion_money_moov_models_payment_requests_merchant_payment2/NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2TestData.json')

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
    ['no_frixion_money_moov_models_payment_requests_merchant_payment201','no_frixion_money_moov_models_payment_requests_merchant_payment202','no_frixion_money_moov_models_payment_requests_merchant_payment203','paymentrequest01','paymentrequest02','paymentrequest03','template01','template02','template03'],
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
  
