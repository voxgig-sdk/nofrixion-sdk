
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


describe('MerchantPaymentRequestTemplateEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.MerchantPaymentRequestTemplate()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['list', 'update', 'load']) {
      if (maybeSkipControl(t, 'entityOp', 'merchant_payment_request_template.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let merchant_payment_request_template_ref01_data = Object.values(setup.data.existing.merchant_payment_request_template)[0] as any

    // LIST
    const merchant_payment_request_template_ref01_ent = client.MerchantPaymentRequestTemplate()
    const merchant_payment_request_template_ref01_match: any = {}
    merchant_payment_request_template_ref01_match['merchant_id'] = setup.idmap['merchant01']

    const merchant_payment_request_template_ref01_list = await merchant_payment_request_template_ref01_ent.list(merchant_payment_request_template_ref01_match)


    // UPDATE
    const merchant_payment_request_template_ref01_data_up0: any = {}
    merchant_payment_request_template_ref01_data_up0.id = merchant_payment_request_template_ref01_data.id
    merchant_payment_request_template_ref01_data_up0 ['paymentrequest_id'] = setup.idmap['paymentrequest_id']

    const merchant_payment_request_template_ref01_markdef_up0 = { name: 'description', value: 'Mark01-merchant_payment_request_template_ref01_' + setup.now }
    ;(merchant_payment_request_template_ref01_data_up0 as any)[merchant_payment_request_template_ref01_markdef_up0.name] = merchant_payment_request_template_ref01_markdef_up0.value

    const merchant_payment_request_template_ref01_resdata_up0 = await merchant_payment_request_template_ref01_ent.update(merchant_payment_request_template_ref01_data_up0)
    assert(merchant_payment_request_template_ref01_resdata_up0.id === merchant_payment_request_template_ref01_data_up0.id)

    assert((merchant_payment_request_template_ref01_resdata_up0 as any)[merchant_payment_request_template_ref01_markdef_up0.name] === merchant_payment_request_template_ref01_markdef_up0.value)


    // LOAD
    const merchant_payment_request_template_ref01_match_dt0: any = {}
    merchant_payment_request_template_ref01_match_dt0.id = merchant_payment_request_template_ref01_data.id
    const merchant_payment_request_template_ref01_data_dt0 = await merchant_payment_request_template_ref01_ent.load(merchant_payment_request_template_ref01_match_dt0)
    assert(merchant_payment_request_template_ref01_data_dt0.id === merchant_payment_request_template_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/merchant_payment_request_template/MerchantPaymentRequestTemplateTestData.json')

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
    ['merchant_payment_request_template01','merchant_payment_request_template02','merchant_payment_request_template03','paymentrequest01','paymentrequest02','paymentrequest03'],
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
  const idmapEnvVal = process.env['NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': 'NONE',
  })

  idmap = env['NOFRIXION_TEST_MERCHANT_PAYMENT_REQUEST_TEMPLATE_ENTID']

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
  
