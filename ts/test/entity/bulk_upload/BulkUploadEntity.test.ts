

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


describe('BulkUploadEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when CUSTOMS_WINDOW_TEST_LIVE=TRUE.
  afterEach(liveDelay('CUSTOMS_WINDOW_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = CustomsWindowSDK.test()
    const ent = testsdk.BulkUpload()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.CUSTOMS_WINDOW_TEST_LIVE
    for (const op of ['create', 'list', 'update', 'load', 'remove']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'bulk_upload.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"active_transport_nationality":{"a":true,"h":"Active Transport Nationality","n":"active_transport_nationality","r":false,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"active_transport_nationality","index$":0},"active_transport_number":{"a":true,"h":"Active Transport Number","n":"active_transport_number","r":false,"t":"`$STRING`","key$":"active_transport_number","index$":1},"arrival_datetime":{"a":true,"fo":"date-time","h":"Arrival Datetime","n":"arrival_datetime","r":false,"t":"`$STRING`","key$":"arrival_datetime","index$":2},"client":{"a":true,"h":"Client","n":"client","r":true,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"client","index$":3},"company_member":{"a":true,"h":"Company Member","n":"company_member","r":true,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"company_member","index$":4},"created_at":{"a":true,"fo":"date-time","h":"Created At","n":"created_at","r":true,"ro":true,"t":"`$STRING`","key$":"created_at","index$":5},"declarant":{"a":true,"h":"Declarant","n":"declarant","r":true,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"declarant","index$":6},"declarations_no":{"a":true,"h":"Declarations No","n":"declarations_no","r":true,"t":"`$INTEGER`","key$":"declarations_no","index$":7},"deleted_at":{"a":true,"fo":"date-time","h":"Deleted At","n":"deleted_at","r":false,"t":"`$STRING`","key$":"deleted_at","index$":8},"departure_datetime":{"a":true,"fo":"date-time","h":"Departure Datetime","n":"departure_datetime","r":false,"t":"`$STRING`","key$":"departure_datetime","index$":9},"errors_file":{"a":true,"h":"Errors File","n":"errors_file","r":true,"t":"`$OBJECT`","key$":"errors_file","index$":10},"external_id":{"a":true,"h":"External Id","n":"external_id","r":false,"t":"`$STRING`","key$":"external_id","index$":11},"failed_declarations_no":{"a":true,"h":"Failed Declarations No","n":"failed_declarations_no","r":false,"t":"`$INTEGER`","key$":"failed_declarations_no","index$":12},"file":{"a":true,"h":"File","n":"file","r":true,"sh":"cuid-format identifier for this entity.","t":"`$OBJECT`","key$":"file","index$":13},"green_routed_no":{"a":true,"h":"Green Routed No","n":"green_routed_no","r":false,"t":"`$INTEGER`","key$":"green_routed_no","index$":14},"h1_fallback_template":{"a":true,"h":"H1 Fallback Template","n":"h1_fallback_template","r":false,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"h1_fallback_template","index$":15},"house_transport_doc_ref":{"a":true,"h":"House Transport Doc Ref","n":"house_transport_doc_ref","r":false,"t":"`$STRING`","key$":"house_transport_doc_ref","index$":16},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"id","index$":17},"issue_date":{"a":true,"fo":"date-time","h":"Issue Date","n":"issue_date","r":false,"t":"`$STRING`","key$":"issue_date","index$":18},"mapping":{"a":true,"h":"Mapping","n":"mapping","r":false,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"mapping","index$":19},"orange_routed_no":{"a":true,"h":"Orange Routed No","n":"orange_routed_no","r":false,"t":"`$INTEGER`","key$":"orange_routed_no","index$":20},"parsed_declarations_no":{"a":true,"h":"Parsed Declarations No","n":"parsed_declarations_no","r":false,"t":"`$INTEGER`","key$":"parsed_declarations_no","index$":21},"parser":{"a":true,"h":"Parser","n":"parser","r":false,"sh":"* `aes_platform` - aes_platform * `cds_platform` - cds_platform * `cds_export_platform` - cds_export_platform * `g4_g3` - g4_g3 * `nhd_platform` - nhd_platform * `platform` - platform * `birds` - birds * `ics2_platform` - ics2_platform","t":"`$STRING`","key$":"parser","index$":22},"parsing_completed_at":{"a":true,"fo":"date-time","h":"Parsing Completed At","n":"parsing_completed_at","r":false,"t":"`$STRING`","key$":"parsing_completed_at","index$":23},"parsing_started_at":{"a":true,"fo":"date-time","h":"Parsing Started At","n":"parsing_started_at","r":false,"t":"`$STRING`","key$":"parsing_started_at","index$":24},"passive_transport_nationality":{"a":true,"h":"Passive Transport Nationality","n":"passive_transport_nationality","r":false,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"passive_transport_nationality","index$":25},"passive_transport_number":{"a":true,"h":"Passive Transport Number","n":"passive_transport_number","r":false,"t":"`$STRING`","key$":"passive_transport_number","index$":26},"processed_declarations_no":{"a":true,"h":"Processed Declarations No","n":"processed_declarations_no","r":false,"t":"`$INTEGER`","key$":"processed_declarations_no","index$":27},"processing_ended_at":{"a":true,"fo":"date-time","h":"Processing Ended At","n":"processing_ended_at","r":false,"t":"`$STRING`","key$":"processing_ended_at","index$":28},"processing_started_at":{"a":true,"fo":"date-time","h":"Processing Started At","n":"processing_started_at","r":false,"t":"`$STRING`","key$":"processing_started_at","index$":29},"receipt_generating_started_at":{"a":true,"fo":"date-time","h":"Receipt Generating Started At","n":"receipt_generating_started_at","r":false,"t":"`$STRING`","key$":"receipt_generating_started_at","index$":30},"receipt_request_started_at":{"a":true,"fo":"date-time","h":"Receipt Request Started At","n":"receipt_request_started_at","r":false,"t":"`$STRING`","key$":"receipt_request_started_at","index$":31},"receipt_request_status":{"a":true,"h":"Receipt Request Status","n":"receipt_request_status","r":false,"sh":"* `pending` - pending * `processing` - processing * `generating` - generating * `completed` - completed","t":"`$STRING`","key$":"receipt_request_status","index$":32},"receipt_request_user":{"a":true,"h":"Receipt Request User","n":"receipt_request_user","r":false,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"receipt_request_user","index$":33},"receipts_zip":{"a":true,"h":"Receipts Zip","n":"receipts_zip","r":false,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"receipts_zip","index$":34},"red_routed_no":{"a":true,"h":"Red Routed No","n":"red_routed_no","r":false,"t":"`$INTEGER`","key$":"red_routed_no","index$":35},"rejected_status_no":{"a":true,"h":"Rejected Status No","n":"rejected_status_no","r":false,"t":"`$INTEGER`","key$":"rejected_status_no","index$":36},"status":{"a":true,"h":"Status","n":"status","r":false,"sh":"* `pending` - pending * `parsing` - parsing * `processing` - processing * `complete` - complete","t":"`$STRING`","key$":"status","index$":37},"template":{"a":true,"h":"Template","n":"template","r":false,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"template","index$":38},"updated_at":{"a":true,"fo":"date-time","h":"Updated At","n":"updated_at","r":true,"ro":true,"t":"`$STRING`","key$":"updated_at","index$":39},"yellow_routed_no":{"a":true,"h":"Yellow Routed No","n":"yellow_routed_no","r":false,"t":"`$INTEGER`","key$":"yellow_routed_no","index$":40}},"id":{"field":"id","name":"id"},"name":"bulk_upload","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /bulk-uploads","source":"openapi3","version":2},"g":{},"k":"http","m":"POST","o":"/bulk-uploads","q":{},"r":{},"s":[{"lit":"bulk-uploads"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"},"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /bulk-uploads","source":"openapi3","version":2},"g":{"query":[{"a":true,"k":"query","n":"cursor","or":"cursor","r":false,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/bulk-uploads","q":{"exist":["cursor"]},"r":{},"s":[{"lit":"bulk-uploads"}],"t":{"req":"`reqdata`","res":"`body.results`"},"index$":0}],"key$":"list"},"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /bulk-uploads/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/bulk-uploads/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"bulk-uploads"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0},{"a":true,"co":{"id":"GET /bulk-uploads/{id}/generate-pdfs","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"GET","o":"/bulk-uploads/{id}/generate-pdfs","q":{"$action":"generate_pdf","exist":["id"]},"r":{},"s":[{"lit":"bulk-uploads"},{"var":"id"},{"lit":"generate-pdfs"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"load"},"remove":{"input":"data","name":"remove","points":[{"a":true,"co":{"id":"DELETE /bulk-uploads/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"DELETE","o":"/bulk-uploads/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"bulk-uploads"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0},{"a":true,"co":{"id":"DELETE /bulk-uploads/{id}/generate-pdfs","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"DELETE","o":"/bulk-uploads/{id}/generate-pdfs","q":{"$action":"generate_pdf","exist":["id"]},"r":{},"s":[{"lit":"bulk-uploads"},{"var":"id"},{"lit":"generate-pdfs"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"remove"},"update":{"input":"data","name":"update","points":[{"a":true,"co":{"id":"PATCH /bulk-uploads/{id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"PATCH","o":"/bulk-uploads/{id}","q":{"exist":["id"]},"r":{},"s":[{"lit":"bulk-uploads"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0},{"a":true,"co":{"id":"PATCH /bulk-uploads/{id}/generate-pdfs","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"PATCH","o":"/bulk-uploads/{id}/generate-pdfs","q":{"$action":"generate_pdf","exist":["id"]},"r":{},"s":[{"lit":"bulk-uploads"},{"var":"id"},{"lit":"generate-pdfs"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":1}],"key$":"update"}},"relations":{"ancestors":[]},"key$":"bulk_upload","name__orig":"bulk_upload","Name":"BulkUpload","name_":"bulk_upload","name-":"bulk-upload","NAME":"BULK_UPLOAD","index$":0}, {"active":true,"entity":"bulk_upload","key$":"BasicBulkUploadFlow","kind":"basic","name":"BasicBulkUploadFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"bulk_upload_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0},{"a":true,"d":{},"i":{},"m":{},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"bulk_upload_ref01"}}],"index$":1},{"a":true,"d":{},"i":{"ref":"bulk_upload_ref01","srcdatavar":"bulk_upload_ref01_data","suffix":"_up0","textfield":"active_transport_nationality"},"m":{},"o":"update","s":[{"apply":"TextFieldMark","def":{"mark":"Mark01-bulk_upload_ref01"}}],"v":[],"index$":2},{"a":true,"d":{},"i":{"ref":"bulk_upload_ref01","srcdatavar":"bulk_upload_ref01_data","suffix":"_dt0"},"m":{"id":"bulk_upload01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-bulk_upload_ref01"}}],"index$":3},{"a":true,"d":{},"i":{"ref":"bulk_upload_ref01","suffix":"_rm0"},"m":{"id":"bulk_upload01"},"o":"remove","s":[],"v":[],"index$":4},{"a":true,"d":{},"i":{"suffix":"_rt0"},"m":{},"o":"list","s":[],"v":[{"apply":"ItemNotExists","def":{"ref":"bulk_upload_ref01"}}],"index$":5}]}, 'BulkUpload', {"POST /bulk-uploads":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","properties":{"file":{"type":"string","description":"cuid-format identifier for this entity.","key$":"file"},"parser":{"enum":["aes_platform","cds_platform","cds_export_platform","g4_g3","nhd_platform","platform","birds","ics2_platform"],"type":"string","description":"* `aes_platform` - aes_platform\n* `cds_platform` - cds_platform\n* `cds_export_platform` - cds_export_platform\n* `g4_g3` - g4_g3\n* `nhd_platform` - nhd_platform\n* `platform` - platform\n* `birds` - birds\n* `ics2_platform` - ics2_platform","x-spec-enum-id":"0bb684377b4b0393","key$":"parser"},"client":{"type":"string","description":"cuid-format identifier for this entity.","key$":"client"},"declarant":{"type":"string","description":"cuid-format identifier for this entity.","key$":"declarant"},"company_member":{"type":"string","description":"cuid-format identifier for this entity.","key$":"company_member"},"template":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"template"},"h1_fallback_template":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"h1_fallback_template"},"external_id":{"type":"string","nullable":true,"maxLength":255,"key$":"external_id"},"mapping":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"mapping"},"passive_transport_number":{"type":"string","nullable":true,"maxLength":255,"key$":"passive_transport_number"},"passive_transport_nationality":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"passive_transport_nationality"},"active_transport_number":{"type":"string","nullable":true,"maxLength":255,"key$":"active_transport_number"},"active_transport_nationality":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"active_transport_nationality"},"house_transport_doc_ref":{"type":"string","nullable":true,"maxLength":255,"key$":"house_transport_doc_ref"},"issue_date":{"type":"string","format":"date-time","nullable":true,"key$":"issue_date"},"departure_datetime":{"type":"string","format":"date-time","nullable":true,"key$":"departure_datetime"},"arrival_datetime":{"type":"string","format":"date-time","nullable":true,"key$":"arrival_datetime"}},"required":["client","company_member","declarant","file"],"title":"BulkUpload","x-ref":"#/components/schemas/BulkUploadCreate","index$":1}},"application/x-www-form-urlencoded":{"schema":{"type":"object","properties":{"file":{"type":"string","description":"cuid-format identifier for this entity.","key$":"file"},"parser":{"enum":["aes_platform","cds_platform","cds_export_platform","g4_g3","nhd_platform","platform","birds","ics2_platform"],"type":"string","description":"* `aes_platform` - aes_platform\n* `cds_platform` - cds_platform\n* `cds_export_platform` - cds_export_platform\n* `g4_g3` - g4_g3\n* `nhd_platform` - nhd_platform\n* `platform` - platform\n* `birds` - birds\n* `ics2_platform` - ics2_platform","x-spec-enum-id":"0bb684377b4b0393","key$":"parser"},"client":{"type":"string","description":"cuid-format identifier for this entity.","key$":"client"},"declarant":{"type":"string","description":"cuid-format identifier for this entity.","key$":"declarant"},"company_member":{"type":"string","description":"cuid-format identifier for this entity.","key$":"company_member"},"template":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"template"},"h1_fallback_template":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"h1_fallback_template"},"external_id":{"type":"string","nullable":true,"maxLength":255,"key$":"external_id"},"mapping":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"mapping"},"passive_transport_number":{"type":"string","nullable":true,"maxLength":255,"key$":"passive_transport_number"},"passive_transport_nationality":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"passive_transport_nationality"},"active_transport_number":{"type":"string","nullable":true,"maxLength":255,"key$":"active_transport_number"},"active_transport_nationality":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"active_transport_nationality"},"house_transport_doc_ref":{"type":"string","nullable":true,"maxLength":255,"key$":"house_transport_doc_ref"},"issue_date":{"type":"string","format":"date-time","nullable":true,"key$":"issue_date"},"departure_datetime":{"type":"string","format":"date-time","nullable":true,"key$":"departure_datetime"},"arrival_datetime":{"type":"string","format":"date-time","nullable":true,"key$":"arrival_datetime"}},"required":["client","company_member","declarant","file"],"title":"BulkUpload","x-ref":"#/components/schemas/BulkUploadCreate"}},"multipart/form-data":{"schema":{"type":"object","properties":{"file":{"type":"string","description":"cuid-format identifier for this entity.","key$":"file"},"parser":{"enum":["aes_platform","cds_platform","cds_export_platform","g4_g3","nhd_platform","platform","birds","ics2_platform"],"type":"string","description":"* `aes_platform` - aes_platform\n* `cds_platform` - cds_platform\n* `cds_export_platform` - cds_export_platform\n* `g4_g3` - g4_g3\n* `nhd_platform` - nhd_platform\n* `platform` - platform\n* `birds` - birds\n* `ics2_platform` - ics2_platform","x-spec-enum-id":"0bb684377b4b0393","key$":"parser"},"client":{"type":"string","description":"cuid-format identifier for this entity.","key$":"client"},"declarant":{"type":"string","description":"cuid-format identifier for this entity.","key$":"declarant"},"company_member":{"type":"string","description":"cuid-format identifier for this entity.","key$":"company_member"},"template":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"template"},"h1_fallback_template":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"h1_fallback_template"},"external_id":{"type":"string","nullable":true,"maxLength":255,"key$":"external_id"},"mapping":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"mapping"},"passive_transport_number":{"type":"string","nullable":true,"maxLength":255,"key$":"passive_transport_number"},"passive_transport_nationality":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"passive_transport_nationality"},"active_transport_number":{"type":"string","nullable":true,"maxLength":255,"key$":"active_transport_number"},"active_transport_nationality":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"active_transport_nationality"},"house_transport_doc_ref":{"type":"string","nullable":true,"maxLength":255,"key$":"house_transport_doc_ref"},"issue_date":{"type":"string","format":"date-time","nullable":true,"key$":"issue_date"},"departure_datetime":{"type":"string","format":"date-time","nullable":true,"key$":"departure_datetime"},"arrival_datetime":{"type":"string","format":"date-time","nullable":true,"key$":"arrival_datetime"}},"required":["client","company_member","declarant","file"],"title":"BulkUpload","x-ref":"#/components/schemas/BulkUploadCreate"}},"*/*":{"schema":{"type":"object","properties":{"file":{"type":"string","description":"cuid-format identifier for this entity.","key$":"file"},"parser":{"enum":["aes_platform","cds_platform","cds_export_platform","g4_g3","nhd_platform","platform","birds","ics2_platform"],"type":"string","description":"* `aes_platform` - aes_platform\n* `cds_platform` - cds_platform\n* `cds_export_platform` - cds_export_platform\n* `g4_g3` - g4_g3\n* `nhd_platform` - nhd_platform\n* `platform` - platform\n* `birds` - birds\n* `ics2_platform` - ics2_platform","x-spec-enum-id":"0bb684377b4b0393","key$":"parser"},"client":{"type":"string","description":"cuid-format identifier for this entity.","key$":"client"},"declarant":{"type":"string","description":"cuid-format identifier for this entity.","key$":"declarant"},"company_member":{"type":"string","description":"cuid-format identifier for this entity.","key$":"company_member"},"template":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"template"},"h1_fallback_template":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"h1_fallback_template"},"external_id":{"type":"string","nullable":true,"maxLength":255,"key$":"external_id"},"mapping":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"mapping"},"passive_transport_number":{"type":"string","nullable":true,"maxLength":255,"key$":"passive_transport_number"},"passive_transport_nationality":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"passive_transport_nationality"},"active_transport_number":{"type":"string","nullable":true,"maxLength":255,"key$":"active_transport_number"},"active_transport_nationality":{"type":"string","description":"cuid-format identifier for this entity.","nullable":true,"key$":"active_transport_nationality"},"house_transport_doc_ref":{"type":"string","nullable":true,"maxLength":255,"key$":"house_transport_doc_ref"},"issue_date":{"type":"string","format":"date-time","nullable":true,"key$":"issue_date"},"departure_datetime":{"type":"string","format":"date-time","nullable":true,"key$":"departure_datetime"},"arrival_datetime":{"type":"string","format":"date-time","nullable":true,"key$":"arrival_datetime"}},"required":["client","company_member","declarant","file"],"title":"BulkUpload","x-ref":"#/components/schemas/BulkUploadCreate"}}},"required":true},"parameters":[]},"GET /bulk-uploads":{"protocol":"http","parameters":[{"name":"cursor","required":false,"in":"query","description":"The pagination cursor value.","schema":{"type":"string"},"index$":0}]},"GET /bulk-uploads/{id}":{"protocol":"http","parameters":[{"in":"path","name":"id","schema":{"type":"string"},"required":true,"index$":0}]},"GET /bulk-uploads/{id}/generate-pdfs":{"protocol":"http","parameters":[{"in":"path","name":"id","schema":{"type":"string"},"required":true,"index$":0}]},"DELETE /bulk-uploads/{id}":{"protocol":"http","parameters":[{"in":"path","name":"id","schema":{"type":"string"},"required":true,"index$":0}]},"DELETE /bulk-uploads/{id}/generate-pdfs":{"protocol":"http","parameters":[{"in":"path","name":"id","schema":{"type":"string"},"required":true,"index$":0}]},"PATCH /bulk-uploads/{id}":{"protocol":"http","parameters":[{"in":"path","name":"id","schema":{"type":"string"},"required":true,"index$":0}]},"PATCH /bulk-uploads/{id}/generate-pdfs":{"protocol":"http","parameters":[{"in":"path","name":"id","schema":{"type":"string"},"required":true,"index$":0}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const bulk_upload_ref01_ent = client.BulkUpload()
    let bulk_upload_ref01_data = setup.data.new.bulk_upload['bulk_upload_ref01']

    bulk_upload_ref01_data = (await bulk_upload_ref01_ent.create(bulk_upload_ref01_data)).data()
    assert(null != bulk_upload_ref01_data.id)


    // LIST
    const bulk_upload_ref01_match: any = {}

    const bulk_upload_ref01_list = (await bulk_upload_ref01_ent.list(bulk_upload_ref01_match)).map((e: any) => e.data())

    assert(!isempty(select(bulk_upload_ref01_list, { id: bulk_upload_ref01_data.id })))


    // UPDATE
    const bulk_upload_ref01_data_up0: any = {}
    bulk_upload_ref01_data_up0.id = bulk_upload_ref01_data.id

    const bulk_upload_ref01_markdef_up0 = { name: 'active_transport_nationality', value: 'Mark01-bulk_upload_ref01_' + setup.now }
    ;(bulk_upload_ref01_data_up0 as any)[bulk_upload_ref01_markdef_up0.name] = bulk_upload_ref01_markdef_up0.value

    const bulk_upload_ref01_resdata_up0 = (await bulk_upload_ref01_ent.update(bulk_upload_ref01_data_up0)).data()
    assert(bulk_upload_ref01_resdata_up0.id === bulk_upload_ref01_data_up0.id)

    assert((bulk_upload_ref01_resdata_up0 as any)[bulk_upload_ref01_markdef_up0.name] === bulk_upload_ref01_markdef_up0.value)


    // LOAD
    const bulk_upload_ref01_match_dt0: any = {}
    bulk_upload_ref01_match_dt0.id = bulk_upload_ref01_data.id
    const bulk_upload_ref01_data_dt0 = (await bulk_upload_ref01_ent.load(bulk_upload_ref01_match_dt0)).data()
    assert(bulk_upload_ref01_data_dt0.id === bulk_upload_ref01_data.id)


    // REMOVE
    const bulk_upload_ref01_match_rm0: any = { id: bulk_upload_ref01_data.id }
    await bulk_upload_ref01_ent.remove(bulk_upload_ref01_match_rm0)
  

    // LIST
    const bulk_upload_ref01_match_rt0: any = {}

    const bulk_upload_ref01_list_rt0 = (await bulk_upload_ref01_ent.list(bulk_upload_ref01_match_rt0)).map((e: any) => e.data())

    assert(isempty(select(bulk_upload_ref01_list_rt0, { id: bulk_upload_ref01_data.id })))


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/bulk_upload/BulkUploadTestData.json')

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
    ['bulk_upload01','bulk_upload02','bulk_upload03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'CUSTOMS_WINDOW_TEST_BULK_UPLOAD_ENTID': idmap,
    'CUSTOMS_WINDOW_TEST_LIVE': 'FALSE',
    'CUSTOMS_WINDOW_TEST_EXPLAIN': 'FALSE',
    'CUSTOMS_WINDOW_APIKEY': '',
  })

  idmap = env['CUSTOMS_WINDOW_TEST_BULK_UPLOAD_ENTID']

  const live = 'TRUE' === env.CUSTOMS_WINDOW_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['CUSTOMS_WINDOW_TEST_BULK_UPLOAD_ENTID']
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
  
