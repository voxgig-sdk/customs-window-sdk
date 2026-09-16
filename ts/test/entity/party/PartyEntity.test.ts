

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


describe('PartyEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when CUSTOMS_WINDOW_TEST_LIVE=TRUE.
  afterEach(liveDelay('CUSTOMS_WINDOW_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = CustomsWindowSDK.test()
    const ent = testsdk.Party()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.CUSTOMS_WINDOW_TEST_LIVE
    for (const op of ['create', 'list', 'update', 'load', 'remove']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'party.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"additional_declaration_type","req":true,"type":"`$OBJECT`","index$":0},{"active":true,"name":"address","req":true,"type":"`$OBJECT`","index$":1},{"active":true,"name":"authorisation","req":true,"type":"`$OBJECT`","index$":2},{"active":true,"name":"bank_details","req":false,"type":"`$STRING`","index$":3},{"active":true,"name":"certificate","req":true,"type":"`$OBJECT`","index$":4},{"active":true,"name":"certificate_type","readOnly":true,"req":true,"type":"`$STRING`","index$":5},{"active":true,"name":"company","req":true,"short":"cuid-format identifier for this entity.","type":"`$STRING`","index$":6},{"active":true,"format":"date-time","name":"created_at","readOnly":true,"req":true,"type":"`$STRING`","index$":7},{"active":true,"name":"customs_office_of_lodgement","req":true,"type":"`$OBJECT`","index$":8},{"active":true,"format":"date-time","name":"deleted_at","req":false,"type":"`$STRING`","index$":9},{"active":true,"format":"email","name":"email","req":false,"type":"`$STRING`","index$":10},{"active":true,"name":"id","req":false,"short":"cuid-format identifier for this entity.","type":"`$STRING`","index$":11},{"active":true,"name":"identification_number","req":false,"type":"`$STRING`","index$":12},{"active":true,"name":"indirect_representative","req":false,"type":"`$BOOLEAN`","index$":13},{"active":true,"name":"name","req":false,"type":"`$STRING`","index$":14},{"active":true,"name":"nhd_last_submission_year","req":false,"type":"`$INTEGER`","index$":15},{"active":true,"name":"nhd_submission_counter","req":false,"type":"`$INTEGER`","index$":16},{"active":true,"name":"person_paying_customs_duty","req":false,"type":"`$STRING`","index$":17},{"active":true,"name":"phone_country_code","req":false,"type":"`$STRING`","index$":18},{"active":true,"name":"phone_number","req":false,"type":"`$STRING`","index$":19},{"active":true,"name":"preferred_payment_method","req":true,"type":"`$OBJECT`","index$":20},{"active":true,"name":"signed_form","req":true,"type":"`$OBJECT`","index$":21},{"active":true,"name":"type","req":false,"short":"* `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `co…","type":"`$STRING`","index$":22},{"active":true,"name":"type_of_person","req":true,"type":"`$OBJECT`","index$":23},{"active":true,"name":"unlocode","req":false,"type":"`$STRING`","index$":24},{"active":true,"format":"date-time","name":"updated_at","readOnly":true,"req":true,"type":"`$STRING`","index$":25}],"id":{"field":"id","name":"id"},"name":"party","op":{"create":{"input":"data","name":"create","points":[{"active":true,"args":{},"contract":{"id":"POST /parties","json":"{\"operationId\":\"parties_create\",\"parameters\":[],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"No response body\"}},\"security\":[{\"Secret Key\":[]}],\"securitySchemes\":{\"Secret Key\":{\"description\":\"Basic authentication with required secret-key\",\"in\":\"header\",\"name\":\"Basic Auth\",\"type\":\"apiKey\"}},\"securitySource\":\"operation\"}","source":"openapi3","version":1},"kind":"http","method":"POST","orig":"/parties","segments":[{"lit":"parties"}],"select":{},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"},"list":{"input":"data","name":"list","points":[{"active":true,"args":{"query":[{"active":true,"kind":"query","name":"cursor","orig":"cursor","reqd":false,"type":"`$STRING`","index$":0},{"active":true,"kind":"query","name":"type","orig":"type","reqd":false,"type":"`$STRING`","index$":1}]},"contract":{"id":"GET /parties","json":"{\"operationId\":\"List all Parties\",\"parameters\":[{\"description\":\"The pagination cursor value.\",\"in\":\"query\",\"name\":\"cursor\",\"required\":false,\"schema\":{\"type\":\"string\"}},{\"description\":\"To filter the parties returned to a specific type, e.g. declarant, client using the following\",\"in\":\"query\",\"name\":\"type\",\"schema\":{\"enum\":[\"authorisation_holder\",\"buyer\",\"carrier\",\"client\",\"consignee\",\"consignor\",\"contact\",\"declarant\",\"exporter\",\"importer\",\"owner\",\"representative\",\"seller\"],\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"next\":{\"nullable\":true,\"type\":\"string\"},\"previous\":{\"nullable\":true,\"type\":\"string\"},\"results\":{\"items\":{\"properties\":{\"additional_declaration_type\":{\"properties\":{\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"deleted_at\":{\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"},\"description\":{\"nullable\":true,\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"list\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"value\":{\"maxLength\":255,\"type\":\"string\"}},\"required\":[\"created_at\",\"list\",\"name\",\"updated_at\",\"value\"],\"title\":\"Code\",\"type\":\"object\"},\"address\":{\"properties\":{\"city\":{\"maxLength\":35,\"nullable\":true,\"type\":\"string\"},\"country_code\":{\"maxLength\":2,\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"deleted_at\":{\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"line_1\":{\"maxLength\":70,\"nullable\":true,\"type\":\"string\"},\"number\":{\"maxLength\":35,\"nullable\":true,\"type\":\"string\"},\"po_box\":{\"maxLength\":70,\"nullable\":true,\"type\":\"string\"},\"postal_code\":{\"maxLength\":9,\"nullable\":true,\"type\":\"string\"},\"raw\":{\"type\":\"string\"},\"sub_division\":{\"maxLength\":35,\"nullable\":true,\"type\":\"string\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"country_code\",\"created_at\",\"raw\",\"updated_at\"],\"title\":\"Address\",\"type\":\"object\"},\"authorisation\":{\"properties\":{\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"deleted_at\":{\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"},\"description\":{\"nullable\":true,\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"list\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"value\":{\"maxLength\":255,\"type\":\"string\"}},\"required\":[\"created_at\",\"list\",\"name\",\"updated_at\",\"value\"],\"title\":\"Code\",\"type\":\"object\"},\"bank_details\":{\"nullable\":true,\"type\":\"string\"},\"certificate\":{\"properties\":{\"frequent_processing_ends_at\":{\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"},\"last_fetched_at\":{\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"},\"processing_start_at\":{\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"},\"type\":{\"description\":\"* `ros` - ros\\n* `hmrc` - hmrc\\n* `nhd` - nhd\\n* `ics2` - ics2\",\"enum\":[\"ros\",\"hmrc\",\"nhd\",\"ics2\"],\"type\":\"string\"}},\"title\":\"Certificate\",\"type\":\"object\"},\"certificate_type\":{\"readOnly\":true,\"type\":\"string\"},\"company\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"customs_office_of_lodgement\":{\"properties\":{\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"deleted_at\":{\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"},\"description\":{\"nullable\":true,\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"list\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"value\":{\"maxLength\":255,\"type\":\"string\"}},\"required\":[\"created_at\",\"list\",\"name\",\"updated_at\",\"value\"],\"title\":\"Code\",\"type\":\"object\"},\"deleted_at\":{\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"},\"email\":{\"format\":\"email\",\"maxLength\":255,\"nullable\":true,\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"identification_number\":{\"maxLength\":17,\"nullable\":true,\"type\":\"string\"},\"indirect_representative\":{\"nullable\":true,\"type\":\"boolean\"},\"name\":{\"maxLength\":70,\"nullable\":true,\"type\":\"string\"},\"nhd_last_submission_year\":{\"maximum\":2147483647,\"minimum\":-2147483648,\"nullable\":true,\"type\":\"integer\"},\"nhd_submission_counter\":{\"maximum\":2147483647,\"minimum\":-2147483648,\"type\":\"integer\"},\"person_paying_customs_duty\":{\"maxLength\":17,\"nullable\":true,\"type\":\"string\"},\"phone_country_code\":{\"maxLength\":255,\"nullable\":true,\"type\":\"string\"},\"phone_number\":{\"maxLength\":255,\"nullable\":true,\"type\":\"string\"},\"preferred_payment_method\":{\"properties\":{\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"deleted_at\":{\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"},\"description\":{\"nullable\":true,\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"list\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"value\":{\"maxLength\":255,\"type\":\"string\"}},\"required\":[\"created_at\",\"list\",\"name\",\"updated_at\",\"value\"],\"title\":\"Code\",\"type\":\"object\"},\"signed_form\":{\"properties\":{\"company\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"extension\":{\"maxLength\":255,\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"name\":{\"maxLength\":255,\"type\":\"string\"},\"public\":{\"type\":\"boolean\"},\"size\":{\"maximum\":2147483647,\"minimum\":-2147483648,\"nullable\":true,\"type\":\"integer\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"url\":{\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"company\",\"created_at\",\"extension\",\"name\",\"public\",\"updated_at\",\"url\"],\"title\":\"File\",\"type\":\"object\"},\"type\":{\"description\":\"* `exporter` - exporter\\n* `importer` - importer\\n* `buyer` - buyer\\n* `seller` - seller\\n* `representative` - representative\\n* `declarant` - declarant\\n* `owner` - owner\\n* `authorisation_holder` - authorisation_holder\\n* `client` - client\\n* `contact` - contact\\n* `consignee` - consignee\\n* `consignor` - consignor\\n* `carrier` - carrier\",\"enum\":[\"exporter\",\"importer\",\"buyer\",\"seller\",\"representative\",\"declarant\",\"owner\",\"authorisation_holder\",\"client\",\"contact\",\"consignee\",\"consignor\",\"carrier\"],\"type\":\"string\"},\"type_of_person\":{\"properties\":{\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"deleted_at\":{\"format\":\"date-time\",\"nullable\":true,\"type\":\"string\"},\"description\":{\"nullable\":true,\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"list\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"value\":{\"maxLength\":255,\"type\":\"string\"}},\"required\":[\"created_at\",\"list\",\"name\",\"updated_at\",\"value\"],\"title\":\"Code\",\"type\":\"object\"},\"unlocode\":{\"maxLength\":17,\"nullable\":true,\"type\":\"string\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"additional_declaration_type\",\"address\",\"authorisation\",\"certificate\",\"certificate_type\",\"company\",\"created_at\",\"customs_office_of_lodgement\",\"preferred_payment_method\",\"signed_form\",\"type_of_person\",\"updated_at\"],\"title\":\"Party\",\"type\":\"object\"},\"type\":\"array\"}},\"title\":\"Party\",\"type\":\"object\"}}},\"description\":\"\"}},\"security\":[{\"Secret Key\":[]}],\"securitySchemes\":{\"Secret Key\":{\"description\":\"Basic authentication with required secret-key\",\"in\":\"header\",\"name\":\"Basic Auth\",\"type\":\"apiKey\"}},\"securitySource\":\"operation\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/parties","segments":[{"lit":"parties"}],"select":{"exist":["cursor","type"]},"transform":{"req":"`reqdata`","res":"`body.results`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"id","orig":"id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /parties/{id}","json":"{\"operationId\":\"parties_retrieve\",\"parameters\":[{\"in\":\"path\",\"name\":\"id\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"No response body\"}},\"security\":[{\"Secret Key\":[]}],\"securitySchemes\":{\"Secret Key\":{\"description\":\"Basic authentication with required secret-key\",\"in\":\"header\",\"name\":\"Basic Auth\",\"type\":\"apiKey\"}},\"securitySource\":\"operation\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/parties/{id}","segments":[{"lit":"parties"},{"var":"id"}],"select":{"exist":["id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"},"remove":{"input":"data","name":"remove","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"id","orig":"id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"DELETE /parties/{id}","json":"{\"operationId\":\"parties_destroy\",\"parameters\":[{\"in\":\"path\",\"name\":\"id\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"204\":{\"description\":\"No response body\"}},\"security\":[{\"Secret Key\":[]}],\"securitySchemes\":{\"Secret Key\":{\"description\":\"Basic authentication with required secret-key\",\"in\":\"header\",\"name\":\"Basic Auth\",\"type\":\"apiKey\"}},\"securitySource\":\"operation\"}","source":"openapi3","version":1},"kind":"http","method":"DELETE","orig":"/parties/{id}","segments":[{"lit":"parties"},{"var":"id"}],"select":{"exist":["id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"remove"},"update":{"input":"data","name":"update","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"id","orig":"id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"PATCH /parties/{id}","json":"{\"operationId\":\"parties_partial_update\",\"parameters\":[{\"in\":\"path\",\"name\":\"id\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"description\":\"No response body\"}},\"security\":[{\"Secret Key\":[]}],\"securitySchemes\":{\"Secret Key\":{\"description\":\"Basic authentication with required secret-key\",\"in\":\"header\",\"name\":\"Basic Auth\",\"type\":\"apiKey\"}},\"securitySource\":\"operation\"}","source":"openapi3","version":1},"kind":"http","method":"PATCH","orig":"/parties/{id}","segments":[{"lit":"parties"},{"var":"id"}],"select":{"exist":["id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"update"}},"relations":{"ancestors":[]},"key$":"party","name__orig":"party","Name":"Party","name_":"party","name-":"party","NAME":"PARTY","index$":5}, {"active":true,"entity":"party","key$":"BasicPartyFlow","kind":"basic","name":"BasicPartyFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"party_ref01"},"match":{},"op":"create","spec":[],"valid":[],"index$":0},{"active":true,"data":{},"input":{},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"party_ref01"}}],"index$":1},{"active":true,"data":{},"input":{"ref":"party_ref01","srcdatavar":"party_ref01_data","suffix":"_up0","textfield":"bank_details"},"match":{},"op":"update","spec":[{"apply":"TextFieldMark","def":{"mark":"Mark01-party_ref01"}}],"valid":[],"index$":2},{"active":true,"data":{},"input":{"ref":"party_ref01","srcdatavar":"party_ref01_data","suffix":"_dt0"},"match":{"id":"party01"},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-party_ref01"}}],"index$":3},{"active":true,"data":{},"input":{"ref":"party_ref01","suffix":"_rm0"},"match":{"id":"party01"},"op":"remove","spec":[],"valid":[],"index$":4},{"active":true,"data":{},"input":{"suffix":"_rt0"},"match":{},"op":"list","spec":[],"valid":[{"apply":"ItemNotExists","def":{"ref":"party_ref01"}}],"index$":5}]}, 'Party')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const party_ref01_ent = client.Party()
    let party_ref01_data = setup.data.new.party['party_ref01']

    party_ref01_data = (await party_ref01_ent.create(party_ref01_data)).data()
    assert(null != party_ref01_data.id)


    // LIST
    const party_ref01_match: any = {}

    const party_ref01_list = (await party_ref01_ent.list(party_ref01_match)).map((e: any) => e.data())

    assert(!isempty(select(party_ref01_list, { id: party_ref01_data.id })))


    // UPDATE
    const party_ref01_data_up0: any = {}
    party_ref01_data_up0.id = party_ref01_data.id

    const party_ref01_markdef_up0 = { name: 'bank_details', value: 'Mark01-party_ref01_' + setup.now }
    ;(party_ref01_data_up0 as any)[party_ref01_markdef_up0.name] = party_ref01_markdef_up0.value

    const party_ref01_resdata_up0 = (await party_ref01_ent.update(party_ref01_data_up0)).data()
    assert(party_ref01_resdata_up0.id === party_ref01_data_up0.id)

    assert((party_ref01_resdata_up0 as any)[party_ref01_markdef_up0.name] === party_ref01_markdef_up0.value)


    // LOAD
    const party_ref01_match_dt0: any = {}
    party_ref01_match_dt0.id = party_ref01_data.id
    const party_ref01_data_dt0 = (await party_ref01_ent.load(party_ref01_match_dt0)).data()
    assert(party_ref01_data_dt0.id === party_ref01_data.id)


    // REMOVE
    const party_ref01_match_rm0: any = { id: party_ref01_data.id }
    await party_ref01_ent.remove(party_ref01_match_rm0)
  

    // LIST
    const party_ref01_match_rt0: any = {}

    const party_ref01_list_rt0 = (await party_ref01_ent.list(party_ref01_match_rt0)).map((e: any) => e.data())

    assert(isempty(select(party_ref01_list_rt0, { id: party_ref01_data.id })))


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/party/PartyTestData.json')

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
    ['party01','party02','party03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'CUSTOMS_WINDOW_TEST_PARTY_ENTID': idmap,
    'CUSTOMS_WINDOW_TEST_LIVE': 'FALSE',
    'CUSTOMS_WINDOW_TEST_EXPLAIN': 'FALSE',
    'CUSTOMS_WINDOW_APIKEY': '',
  })

  idmap = env['CUSTOMS_WINDOW_TEST_PARTY_ENTID']

  const live = 'TRUE' === env.CUSTOMS_WINDOW_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['CUSTOMS_WINDOW_TEST_PARTY_ENTID']
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
  
