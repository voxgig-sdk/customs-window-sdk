

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { CustomsWindowSDK, BaseFeature, stdutil } from '../../..'

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


describe('FileEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when CUSTOMS_WINDOW_TEST_LIVE=TRUE.
  afterEach(liveDelay('CUSTOMS_WINDOW_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = CustomsWindowSDK.test()
    const ent = testsdk.File()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.CUSTOMS_WINDOW_TEST_LIVE
    for (const op of ['create']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'file.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"company","req":true,"short":"cuid-format identifier for this entity.","type":"`$STRING`","index$":0},{"active":true,"format":"date-time","name":"created_at","readOnly":true,"req":true,"type":"`$STRING`","index$":1},{"active":true,"name":"extension","readOnly":true,"req":true,"type":"`$STRING`","index$":2},{"active":true,"format":"uri","name":"file","req":true,"type":"`$STRING`","index$":3},{"active":true,"name":"id","req":false,"short":"cuid-format identifier for this entity.","type":"`$STRING`","index$":4},{"active":true,"name":"name","readOnly":true,"req":true,"type":"`$STRING`","index$":5},{"active":true,"name":"public","op":{"create":{"req":false,"type":"`$BOOLEAN`"}},"req":true,"type":"`$BOOLEAN`","index$":6},{"active":true,"name":"size","req":false,"type":"`$INTEGER`","index$":7},{"active":true,"format":"date-time","name":"updated_at","readOnly":true,"req":true,"type":"`$STRING`","index$":8},{"active":true,"name":"url","readOnly":true,"req":true,"type":"`$STRING`","index$":9}],"id":{"field":"id","name":"id"},"name":"file","op":{"create":{"input":"data","name":"create","points":[{"active":true,"args":{},"contract":{"id":"POST /files","json":"{\"operationId\":\"Create a File\",\"parameters\":[],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"*/*\":{\"schema\":{\"properties\":{\"company\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"extension\":{\"readOnly\":true,\"type\":\"string\"},\"file\":{\"format\":\"uri\",\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"name\":{\"readOnly\":true,\"type\":\"string\"},\"public\":{\"default\":true,\"type\":\"boolean\"},\"size\":{\"maximum\":2147483647,\"minimum\":-2147483648,\"nullable\":true,\"type\":\"integer\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"url\":{\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"company\",\"created_at\",\"extension\",\"file\",\"name\",\"updated_at\",\"url\"],\"title\":\"File\",\"type\":\"object\"}},\"application/json\":{\"schema\":{\"properties\":{\"company\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"extension\":{\"readOnly\":true,\"type\":\"string\"},\"file\":{\"format\":\"uri\",\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"name\":{\"readOnly\":true,\"type\":\"string\"},\"public\":{\"default\":true,\"type\":\"boolean\"},\"size\":{\"maximum\":2147483647,\"minimum\":-2147483648,\"nullable\":true,\"type\":\"integer\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"url\":{\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"company\",\"created_at\",\"extension\",\"file\",\"name\",\"updated_at\",\"url\"],\"title\":\"File\",\"type\":\"object\"}},\"application/x-www-form-urlencoded\":{\"schema\":{\"properties\":{\"company\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"extension\":{\"readOnly\":true,\"type\":\"string\"},\"file\":{\"format\":\"uri\",\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"name\":{\"readOnly\":true,\"type\":\"string\"},\"public\":{\"default\":true,\"type\":\"boolean\"},\"size\":{\"maximum\":2147483647,\"minimum\":-2147483648,\"nullable\":true,\"type\":\"integer\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"url\":{\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"company\",\"created_at\",\"extension\",\"file\",\"name\",\"updated_at\",\"url\"],\"title\":\"File\",\"type\":\"object\"}},\"multipart/form-data\":{\"schema\":{\"properties\":{\"company\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"extension\":{\"readOnly\":true,\"type\":\"string\"},\"file\":{\"format\":\"uri\",\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"name\":{\"readOnly\":true,\"type\":\"string\"},\"public\":{\"default\":true,\"type\":\"boolean\"},\"size\":{\"maximum\":2147483647,\"minimum\":-2147483648,\"nullable\":true,\"type\":\"integer\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"url\":{\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"company\",\"created_at\",\"extension\",\"file\",\"name\",\"updated_at\",\"url\"],\"title\":\"File\",\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"company\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"extension\":{\"maxLength\":255,\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"name\":{\"maxLength\":255,\"type\":\"string\"},\"public\":{\"type\":\"boolean\"},\"size\":{\"maximum\":2147483647,\"minimum\":-2147483648,\"nullable\":true,\"type\":\"integer\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"url\":{\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"company\",\"created_at\",\"extension\",\"name\",\"public\",\"updated_at\",\"url\"],\"title\":\"File\",\"type\":\"object\"}}},\"description\":\"\"}},\"security\":[{\"Secret Key\":[]}],\"securitySchemes\":{\"Secret Key\":{\"description\":\"Basic authentication with required secret-key\",\"in\":\"header\",\"name\":\"Basic Auth\",\"type\":\"apiKey\"}},\"securitySource\":\"operation\"}","source":"openapi3","version":1},"kind":"http","method":"POST","orig":"/files","segments":[{"lit":"files"}],"select":{},"transform":{"req":{"file":"`reqdata`"},"res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"file","name__orig":"file","Name":"File","name_":"file","name-":"file","NAME":"FILE","index$":1}, {"active":true,"entity":"file","key$":"BasicFileFlow","kind":"basic","name":"BasicFileFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"file_ref01"},"match":{},"op":"create","spec":[],"valid":[],"index$":0}]}, 'File')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const file_ref01_ent = client.File()
    let file_ref01_data = setup.data.new.file['file_ref01']

    file_ref01_data = (await file_ref01_ent.create(file_ref01_data)).data()
    assert(null != file_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/file/FileTestData.json')

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
    ['file01','file02','file03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'CUSTOMS_WINDOW_TEST_FILE_ENTID': idmap,
    'CUSTOMS_WINDOW_TEST_LIVE': 'FALSE',
    'CUSTOMS_WINDOW_TEST_EXPLAIN': 'FALSE',
    'CUSTOMS_WINDOW_APIKEY': '',
  })

  idmap = env['CUSTOMS_WINDOW_TEST_FILE_ENTID']

  const live = 'TRUE' === env.CUSTOMS_WINDOW_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['CUSTOMS_WINDOW_TEST_FILE_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new CustomsWindowSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
        apikey: env.CUSTOMS_WINDOW_APIKEY,
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
    explain: 'TRUE' === env.CUSTOMS_WINDOW_TEST_EXPLAIN,
    live,
    transport,
    now: Date.now(),
  }

  return setup
}
  
