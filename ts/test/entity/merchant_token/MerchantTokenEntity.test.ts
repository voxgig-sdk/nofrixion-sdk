
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


describe('MerchantTokenEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.MerchantToken()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['create', 'list', 'update', 'load']) {
      if (maybeSkipControl(t, 'entityOp', 'merchant_token.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_MERCHANT_TOKEN_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const merchant_token_ref01_ent = client.MerchantToken()
    let merchant_token_ref01_data = setup.data.new.merchant_token['merchant_token_ref01']
    merchant_token_ref01_data['merchant_id'] = setup.idmap['merchant01']

    merchant_token_ref01_data = await merchant_token_ref01_ent.create(merchant_token_ref01_data)
    assert(null != merchant_token_ref01_data.id)


    // LIST
    const merchant_token_ref01_match: any = {}
    merchant_token_ref01_match['merchant_id'] = setup.idmap['merchant01']

    const merchant_token_ref01_list = await merchant_token_ref01_ent.list(merchant_token_ref01_match)

    assert(!isempty(select(merchant_token_ref01_list, { id: merchant_token_ref01_data.id })))


    // UPDATE
    const merchant_token_ref01_data_up0: any = {}
    merchant_token_ref01_data_up0.id = merchant_token_ref01_data.id

    const merchant_token_ref01_markdef_up0 = { name: 'description', value: 'Mark01-merchant_token_ref01_' + setup.now }
    ;(merchant_token_ref01_data_up0 as any)[merchant_token_ref01_markdef_up0.name] = merchant_token_ref01_markdef_up0.value

    const merchant_token_ref01_resdata_up0 = await merchant_token_ref01_ent.update(merchant_token_ref01_data_up0)
    assert(merchant_token_ref01_resdata_up0.id === merchant_token_ref01_data_up0.id)

    assert((merchant_token_ref01_resdata_up0 as any)[merchant_token_ref01_markdef_up0.name] === merchant_token_ref01_markdef_up0.value)


    // LOAD
    const merchant_token_ref01_match_dt0: any = {}
    merchant_token_ref01_match_dt0.id = merchant_token_ref01_data.id
    const merchant_token_ref01_data_dt0 = await merchant_token_ref01_ent.load(merchant_token_ref01_match_dt0)
    assert(merchant_token_ref01_data_dt0.id === merchant_token_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/merchant_token/MerchantTokenTestData.json')

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
    ['merchant_token01','merchant_token02','merchant_token03','merchant01','merchant02','merchant03'],
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
  const idmapEnvVal = process.env['NOFRIXION_TEST_MERCHANT_TOKEN_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'NOFRIXION_TEST_MERCHANT_TOKEN_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': 'NONE',
  })

  idmap = env['NOFRIXION_TEST_MERCHANT_TOKEN_ENTID']

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
  
