
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'


import { CustomsWindowSDK, BaseFeature, stdutil } from '../../..'

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


describe('SubmissionDetailEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when CUSTOMS_WINDOW_TEST_LIVE=TRUE.
  afterEach(liveDelay('CUSTOMS_WINDOW_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = CustomsWindowSDK.test()
    const ent = testsdk.SubmissionDetail()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.CUSTOMS_WINDOW_TEST_LIVE
    for (const op of ['create']) {
      if (maybeSkipControl(t, 'entityOp', 'submission_detail.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set CUSTOMS_WINDOW_TEST_SUBMISSION_DETAIL_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const submission_detail_ref01_ent = client.SubmissionDetail()
    let submission_detail_ref01_data = setup.data.new.submission_detail['submission_detail_ref01']

    submission_detail_ref01_data = (await submission_detail_ref01_ent.create(submission_detail_ref01_data)).data()
    assert(null != submission_detail_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/submission_detail/SubmissionDetailTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = CustomsWindowSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['submission_detail01','submission_detail02','submission_detail03'],
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
  const idmapEnvVal = process.env['CUSTOMS_WINDOW_TEST_SUBMISSION_DETAIL_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'CUSTOMS_WINDOW_TEST_SUBMISSION_DETAIL_ENTID': idmap,
    'CUSTOMS_WINDOW_TEST_LIVE': 'FALSE',
    'CUSTOMS_WINDOW_TEST_EXPLAIN': 'FALSE',
    'CUSTOMS_WINDOW_APIKEY': 'NONE',
  })

  idmap = env['CUSTOMS_WINDOW_TEST_SUBMISSION_DETAIL_ENTID']

  const live = 'TRUE' === env.CUSTOMS_WINDOW_TEST_LIVE

  if (live) {
    client = new CustomsWindowSDK(merge([
      {
        apikey: env.CUSTOMS_WINDOW_APIKEY,
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
    explain: 'TRUE' === env.CUSTOMS_WINDOW_TEST_EXPLAIN,
    live,
    syntheticOnly: live && !idmapOverridden,
    now: Date.now(),
  }

  return setup
}
  
