

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
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"additional_declaration_type":{"a":true,"h":"Additional Declaration Type","n":"additional_declaration_type","r":true,"t":"`$OBJECT`","key$":"additional_declaration_type","index$":0},"address":{"a":true,"h":"Address","n":"address","r":true,"t":"`$OBJECT`","key$":"address","index$":1},"authorisation":{"a":true,"h":"Authorisation","n":"authorisation","r":true,"t":"`$OBJECT`","key$":"authorisation","index$":2},"bank_details":{"a":true,"h":"Bank Details","n":"bank_details","r":false,"t":"`$STRING`","key$":"bank_details","index$":3},"certificate":{"a":true,"h":"Certificate","n":"certificate","r":true,"t":"`$OBJECT`","key$":"certificate","index$":4},"certificate_type":{"a":true,"h":"Certificate Type","n":"certificate_type","r":true,"ro":true,"t":"`$STRING`","key$":"certificate_type","index$":5},"company":{"a":true,"h":"Company","n":"company","r":true,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"company","index$":6},"created_at":{"a":true,"fo":"date-time","h":"Created At","n":"created_at","r":true,"ro":true,"t":"`$STRING`","key$":"created_at","index$":7},"customs_office_of_lodgement":{"a":true,"h":"Customs Office Of Lodgement","n":"customs_office_of_lodgement","r":true,"t":"`$OBJECT`","key$":"customs_office_of_lodgement","index$":8},"deleted_at":{"a":true,"fo":"date-time","h":"Deleted At","n":"deleted_at","r":false,"t":"`$STRING`","key$":"deleted_at","index$":9},"email":{"a":true,"fo":"email","h":"Email","n":"email","r":false,"t":"`$STRING`","key$":"email","index$":10},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"id","index$":11},"identification_number":{"a":true,"h":"Identification Number","n":"identification_number","r":false,"t":"`$STRING`","key$":"identification_number","index$":12},"indirect_representative":{"a":true,"h":"Indirect Representative","n":"indirect_representative","r":false,"t":"`$BOOLEAN`","key$":"indirect_representative","index$":13},"name":{"a":true,"h":"Name","n":"name","r":false,"t":"`$STRING`","key$":"name","index$":14},"nhd_last_submission_year":{"a":true,"h":"Nhd Last Submission Year","n":"nhd_last_submission_year","r":false,"t":"`$INTEGER`","key$":"nhd_last_submission_year","index$":15},"nhd_submission_counter":{"a":true,"h":"Nhd Submission Counter","n":"nhd_submission_counter","r":false,"t":"`$INTEGER`","key$":"nhd_submission_counter","index$":16},"person_paying_customs_duty":{"a":true,"h":"Person Paying Customs Duty","n":"person_paying_customs_duty","r":false,"t":"`$STRING`","key$":"person_paying_customs_duty","index$":17},"phone_country_code":{"a":true,"h":"Phone Country Code","n":"phone_country_code","r":false,"t":"`$STRING`","key$":"phone_country_code","index$":18},"phone_number":{"a":true,"h":"Phone Number","n":"phone_number","r":false,"t":"`$STRING`","key$":"phone_number","index$":19},"preferred_payment_method":{"a":true,"h":"Preferred Payment Method","n":"preferred_payment_method","r":true,"t":"`$OBJECT`","key$":"preferred_payment_method","index$":20},"signed_form":{"a":true,"h":"Signed Form","n":"signed_form","r":true,"t":"`$OBJECT`","key$":"signed_form","index$":21},"type":{"a":true,"h":"Type","n":"type","r":false,"sh":"* `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `co…","t":"`$STRING`","key$":"type","index$":22},"type_of_person":{"a":true,"h":"Type Of Person","n":"type_of_person","r":true,"t":"`$OBJECT`","key$":"type_of_person","index$":23},"unlocode":{"a":true,"h":"Unlocode","n":"unlocode","r":false,"t":"`$STRING`","key$":"unlocode","index$":24},"updated_at":{"a":true,"fo":"date-time","h":"Updated At","n":"updated_at","r":true,"ro":true,"t":"`$STRING`","key$":"updated_at","index$":25}},"id":{"field":"id","name":"id"},"name":"party","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /parties","source":"openapi3","version":2},"g":{},"k":"http","m":"POST","o":"/parties","q":{},"r":{},"s":[{"lit":"parties"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"},"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /parties","source":"openapi3","version":2},"g":{"query":[{"a":true,"k":"query","n":"cursor","or":"cursor","r":false,"t":"`$STRING`","index$":0},{"a":true,"k":"query","n":"type","or":"type","r":false,"t":"`$STRING`","index$":1}]},"k":"http","m":"GET","o":"/parties","q":{"exist":["cursor","type"]},"r":{},"s":[{"lit":"parties"}],"t":{"req":"`reqdata`","res":"`body.results`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /parties/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/parties/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"parties"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"},"remove":{"input":"data","name":"remove","points":[{"a":true,"co":{"id":"DELETE /parties/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"DELETE","o":"/parties/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"parties"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"remove"},"update":{"input":"data","name":"update","points":[{"a":true,"co":{"id":"PATCH /parties/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"PATCH","o":"/parties/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"parties"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"update"}},"relations":{"ancestors":[]},"key$":"party","name__orig":"party","Name":"Party","name_":"party","name-":"party","NAME":"PARTY","index$":2}, {"active":true,"entity":"party","key$":"BasicPartyFlow","kind":"basic","name":"BasicPartyFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"party_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0},{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"party_ref01"}}],"index$":1},{"a":true,"d":{},"i":{"ref":"party_ref01","srcdatavar":"party_ref01_data","suffix":"_up0","textfield":"bank_details"},"m":{},"o":"update","s":[{"apply":"TextFieldMark","def":{"mark":"Mark01-party_ref01"}}],"v":[],"index$":2},{"a":true,"d":{},"i":{"ref":"party_ref01","srcdatavar":"party_ref01_data","suffix":"_dt0"},"m":{"id":"party01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-party_ref01"}}],"index$":3},{"a":true,"d":{},"i":{"ref":"party_ref01","suffix":"_rm0"},"m":{"id":"party01"},"o":"remove","s":[],"v":[],"index$":4},{"a":true,"d":{},"i":{"suffix":"_rt0"},"m":{},"o":"list","s":[],"v":[{"apply":"ItemNotExists","def":{"ref":"party_ref01"}}],"index$":5}]}, 'Party', {"POST /parties":{"protocol":"http","parameters":[]},"GET /parties":{"protocol":"http","parameters":[{"name":"cursor","required":false,"in":"query","description":"The pagination cursor value.","schema":{"type":"string"},"index$":0},{"in":"query","name":"type","schema":{"type":"string","enum":["authorisation_holder","buyer","carrier","client","consignee","consignor","contact","declarant","exporter","importer","owner","representative","seller"]},"description":"To filter the parties returned to a specific type, e.g. declarant, client using the following","index$":1}]},"GET /parties/{id}":{"protocol":"http","parameters":[{"in":"path","name":"id","schema":{"type":"string"},"required":true,"index$":0}]},"DELETE /parties/{id}":{"protocol":"http","parameters":[{"in":"path","name":"id","schema":{"type":"string"},"required":true,"index$":0}]},"PATCH /parties/{id}":{"protocol":"http","parameters":[{"in":"path","name":"id","schema":{"type":"string"},"required":true,"index$":0}]}})
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
  
