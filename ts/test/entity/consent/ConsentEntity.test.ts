
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


describe('ConsentEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NOFRIXION_TEST_LIVE=TRUE.
  afterEach(liveDelay('NOFRIXION_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NofrixionSDK.test()
    const ent = testsdk.Consent()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NOFRIXION_TEST_LIVE
    for (const op of ['create', 'list', 'update', 'load', 'remove']) {
      if (maybeSkipControl(t, 'entityOp', 'consent.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set NOFRIXION_TEST_CONSENT_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const consent_ref01_ent = client.Consent()
    let consent_ref01_data = setup.data.new.consent['consent_ref01']
    consent_ref01_data['email'] = setup.idmap['email01']
    consent_ref01_data['merchant_id'] = setup.idmap['merchant01']

    consent_ref01_data = await consent_ref01_ent.create(consent_ref01_data)
    assert(null != consent_ref01_data.id)


    // LIST
    const consent_ref01_match: any = {}
    consent_ref01_match['email'] = setup.idmap['email01']
    consent_ref01_match['merchant_id'] = setup.idmap['merchant01']

    const consent_ref01_list = await consent_ref01_ent.list(consent_ref01_match)

    assert(!isempty(select(consent_ref01_list, { id: consent_ref01_data.id })))


    // UPDATE
    const consent_ref01_data_up0: any = {}
    consent_ref01_data_up0.id = consent_ref01_data.id

    const consent_ref01_markdef_up0 = { name: 'authorisation_url', value: 'Mark01-consent_ref01_' + setup.now }
    ;(consent_ref01_data_up0 as any)[consent_ref01_markdef_up0.name] = consent_ref01_markdef_up0.value

    const consent_ref01_resdata_up0 = await consent_ref01_ent.update(consent_ref01_data_up0)
    assert(consent_ref01_resdata_up0.id === consent_ref01_data_up0.id)

    assert((consent_ref01_resdata_up0 as any)[consent_ref01_markdef_up0.name] === consent_ref01_markdef_up0.value)


    // LOAD
    const consent_ref01_match_dt0: any = {}
    consent_ref01_match_dt0.id = consent_ref01_data.id
    const consent_ref01_data_dt0 = await consent_ref01_ent.load(consent_ref01_match_dt0)
    assert(consent_ref01_data_dt0.id === consent_ref01_data.id)


    // REMOVE
    const consent_ref01_match_rm0: any = { id: consent_ref01_data.id }
    await consent_ref01_ent.remove(consent_ref01_match_rm0)
  

    // LIST
    const consent_ref01_match_rt0: any = {}
    consent_ref01_match_rt0['email'] = setup.idmap['email01']
    consent_ref01_match_rt0['merchant_id'] = setup.idmap['merchant01']

    const consent_ref01_list_rt0 = await consent_ref01_ent.list(consent_ref01_match_rt0)

    assert(isempty(select(consent_ref01_list_rt0, { id: consent_ref01_data.id })))


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/consent/ConsentTestData.json')

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
    ['consent01','consent02','consent03','consent01','consent02','consent03'],
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
  const idmapEnvVal = process.env['NOFRIXION_TEST_CONSENT_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'NOFRIXION_TEST_CONSENT_ENTID': idmap,
    'NOFRIXION_TEST_LIVE': 'FALSE',
    'NOFRIXION_TEST_EXPLAIN': 'FALSE',
    'NOFRIXION_APIKEY': 'NONE',
  })

  idmap = env['NOFRIXION_TEST_CONSENT_ENTID']

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
  
