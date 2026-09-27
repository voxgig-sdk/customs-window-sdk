

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
      if (!live && maybeSkipControl(t, 'entityOp', 'submission_detail.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"additional_external_ids":{"a":true,"h":"Additional External Ids","n":"additional_external_ids","r":false,"t":"`$ARRAY`","key$":"additional_external_ids","index$":0},"additional_information":{"a":true,"h":"Additional Information","n":"additional_information","r":true,"t":"`$ARRAY`","key$":"additional_information","index$":1},"amendment_status":{"a":true,"h":"Amendment Status","n":"amendment_status","r":false,"sh":"* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled","t":"`$STRING`","key$":"amendment_status","index$":2},"clearance_slip":{"a":true,"h":"Clearance Slip","n":"clearance_slip","r":true,"t":"`$OBJECT`","key$":"clearance_slip","index$":3},"client":{"a":true,"h":"Client","n":"client","r":true,"t":"`$OBJECT`","key$":"client","index$":4},"company":{"a":true,"h":"Company","n":"company","r":true,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"company","index$":5},"company_member":{"a":true,"h":"Company Member","n":"company_member","r":true,"sh":"cuid-format identifier for this entity.","t":"`$OBJECT`","key$":"company_member","index$":6},"consignee":{"a":true,"h":"Consignee","n":"consignee","r":true,"t":"`$OBJECT`","key$":"consignee","index$":7},"consignor":{"a":true,"h":"Consignor","n":"consignor","r":true,"t":"`$OBJECT`","key$":"consignor","index$":8},"created_at":{"a":true,"fo":"date-time","h":"Created At","n":"created_at","r":true,"ro":true,"t":"`$STRING`","key$":"created_at","index$":9},"declarant":{"a":true,"h":"Declarant","n":"declarant","r":true,"t":"`$OBJECT`","key$":"declarant","index$":10},"document_upload_status":{"a":true,"h":"Document Upload Status","n":"document_upload_status","r":false,"sh":"* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled","t":"`$STRING`","key$":"document_upload_status","index$":11},"documents_presentation_requested":{"a":true,"h":"Documents Presentation Requested","n":"documents_presentation_requested","r":false,"t":"`$BOOLEAN`","key$":"documents_presentation_requested","index$":12},"documents_upload_requested":{"a":true,"h":"Documents Upload Requested","n":"documents_upload_requested","r":false,"t":"`$BOOLEAN`","key$":"documents_upload_requested","index$":13},"external_id":{"a":true,"h":"External Id","n":"external_id","r":false,"t":"`$STRING`","key$":"external_id","index$":14},"form":{"a":true,"h":"Form","n":"form","r":true,"ro":true,"t":"`$STRING`","key$":"form","index$":15},"goods_presentation_status":{"a":true,"h":"Goods Presentation Status","n":"goods_presentation_status","r":false,"sh":"* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested","t":"`$STRING`","key$":"goods_presentation_status","index$":16},"hrcm_status":{"a":true,"h":"Hrcm Status","n":"hrcm_status","r":false,"sh":"* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled","t":"`$STRING`","key$":"hrcm_status","index$":17},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"id","index$":18},"invalidation_status":{"a":true,"h":"Invalidation Status","n":"invalidation_status","r":false,"sh":"* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled","t":"`$STRING`","key$":"invalidation_status","index$":19},"is_global_template":{"a":true,"h":"Is Global Template","n":"is_global_template","r":false,"t":"`$BOOLEAN`","key$":"is_global_template","index$":20},"latest_notification_item":{"a":true,"h":"Latest Notification Item","n":"latest_notification_item","r":false,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"latest_notification_item","index$":21},"latest_state":{"a":true,"h":"Latest State","n":"latest_state","r":true,"t":"`$OBJECT`","key$":"latest_state","index$":22},"lrn":{"a":true,"h":"Lrn","n":"lrn","r":false,"t":"`$STRING`","key$":"lrn","index$":23},"mrn":{"a":true,"h":"Mrn","n":"mrn","r":false,"t":"`$STRING`","key$":"mrn","index$":24},"name":{"a":true,"h":"Name","n":"name","r":false,"t":"`$STRING`","key$":"name","index$":25},"receipt":{"a":true,"h":"Receipt","n":"receipt","r":true,"t":"`$OBJECT`","key$":"receipt","index$":26},"refund_application_status":{"a":true,"h":"Refund Application Status","n":"refund_application_status","r":false,"sh":"* `processing` - processing * `rejected` - rejected * `accepted` - accepted","t":"`$STRING`","key$":"refund_application_status","index$":27},"route":{"a":true,"h":"Route","n":"route","r":false,"sh":"* `green` - green * `orange` - orange * `red` - red * `yellow` - yellow","t":"`$STRING`","key$":"route","index$":28},"shipment_items_no":{"a":true,"h":"Shipment Items No","n":"shipment_items_no","r":false,"t":"`$INTEGER`","key$":"shipment_items_no","index$":29},"shipment_items_quantity_no":{"a":true,"h":"Shipment Items Quantity No","n":"shipment_items_quantity_no","r":false,"t":"`$INTEGER`","key$":"shipment_items_quantity_no","index$":30},"source":{"a":true,"h":"Source","n":"source","r":false,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"source","index$":31},"source_type":{"a":true,"h":"Source Type","n":"source_type","r":false,"sh":"* `template` - template * `automated_import` - automated_import","t":"`$STRING`","key$":"source_type","index$":32},"status":{"a":true,"h":"Status","n":"status","r":false,"sh":"* `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re…","t":"`$STRING`","key$":"status","index$":33},"submission":{"a":true,"h":"Submission","n":"submission","r":true,"sh":"cuid-format identifier for this entity.","t":"`$STRING`","key$":"submission","index$":34},"supporting_documents":{"a":true,"h":"Supporting Documents","n":"supporting_documents","r":false,"t":"`$ARRAY`","key$":"supporting_documents","index$":35},"template":{"a":true,"h":"Template","n":"template","r":false,"t":"`$BOOLEAN`","key$":"template","index$":36},"total_tax_amount":{"a":true,"fo":"decimal","h":"Total Tax Amount","n":"total_tax_amount","r":false,"t":"`$STRING`","key$":"total_tax_amount","index$":37},"updated_at":{"a":true,"fo":"date-time","h":"Updated At","n":"updated_at","r":true,"ro":true,"t":"`$STRING`","key$":"updated_at","index$":38},"verification_errors":{"a":true,"h":"Verification Errors","n":"verification_errors","r":false,"t":"`$ARRAY`","key$":"verification_errors","index$":39},"verification_status":{"a":true,"h":"Verification Status","n":"verification_status","r":false,"sh":"* `passed` - passed * `failed` - failed","t":"`$STRING`","key$":"verification_status","index$":40}},"id":{"field":"id","name":"id"},"name":"submission_detail","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /documents-request","source":"openapi3","version":2},"g":{},"k":"http","m":"POST","o":"/documents-request","q":{},"r":{},"s":[{"lit":"documents-request"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"}},"relations":{"ancestors":[]},"key$":"submission_detail","name__orig":"submission_detail","Name":"SubmissionDetail","name_":"submission_detail","name-":"submission-detail","NAME":"SUBMISSION_DETAIL","index$":4}, {"active":true,"entity":"submission_detail","key$":"BasicSubmissionDetailFlow","kind":"basic","name":"BasicSubmissionDetailFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"submission_detail_ref01"},"m":{},"o":"create","s":[],"v":[],"index$":0}]}, 'SubmissionDetail', {"POST /documents-request":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","properties":{"additional_information":{"type":"array","items":{"type":"object","properties":{"type":{"type":"string","description":"cuid-format identifier for this entity."},"info":{"type":"string","maxLength":512},"referral_request_reference":{"type":"string","maxLength":512}},"required":["info","type"],"x-ref":"#/components/schemas/AdditionalInformationValidate"},"key$":"additional_information"},"supporting_documents":{"type":"array","items":{"type":"object","properties":{"file":{"type":"string","description":"cuid-format identifier for this entity."},"id":{"type":"string","maxLength":512},"description":{"type":"string","maxLength":512}},"required":["description","file","id"],"x-ref":"#/components/schemas/SupportingDocumentValidate"},"key$":"supporting_documents"},"submission":{"type":"string","description":"cuid-format identifier for this entity.","key$":"submission"},"company":{"type":"string","description":"cuid-format identifier for this entity.","key$":"company"},"company_member":{"type":"string","description":"cuid-format identifier for this entity.","key$":"company_member"}},"required":["additional_information","company","company_member","submission"],"x-ref":"#/components/schemas/DocumentsRequest","index$":1}},"application/x-www-form-urlencoded":{"schema":{"type":"object","properties":{"additional_information":{"type":"array","items":{"type":"object","properties":{"type":{"type":"string","description":"cuid-format identifier for this entity."},"info":{"type":"string","maxLength":512},"referral_request_reference":{"type":"string","maxLength":512}},"required":["info","type"],"x-ref":"#/components/schemas/AdditionalInformationValidate"},"key$":"additional_information"},"supporting_documents":{"type":"array","items":{"type":"object","properties":{"file":{"type":"string","description":"cuid-format identifier for this entity."},"id":{"type":"string","maxLength":512},"description":{"type":"string","maxLength":512}},"required":["description","file","id"],"x-ref":"#/components/schemas/SupportingDocumentValidate"},"key$":"supporting_documents"},"submission":{"type":"string","description":"cuid-format identifier for this entity.","key$":"submission"},"company":{"type":"string","description":"cuid-format identifier for this entity.","key$":"company"},"company_member":{"type":"string","description":"cuid-format identifier for this entity.","key$":"company_member"}},"required":["additional_information","company","company_member","submission"],"x-ref":"#/components/schemas/DocumentsRequest"}},"multipart/form-data":{"schema":{"type":"object","properties":{"additional_information":{"type":"array","items":{"type":"object","properties":{"type":{"type":"string","description":"cuid-format identifier for this entity."},"info":{"type":"string","maxLength":512},"referral_request_reference":{"type":"string","maxLength":512}},"required":["info","type"],"x-ref":"#/components/schemas/AdditionalInformationValidate"},"key$":"additional_information"},"supporting_documents":{"type":"array","items":{"type":"object","properties":{"file":{"type":"string","description":"cuid-format identifier for this entity."},"id":{"type":"string","maxLength":512},"description":{"type":"string","maxLength":512}},"required":["description","file","id"],"x-ref":"#/components/schemas/SupportingDocumentValidate"},"key$":"supporting_documents"},"submission":{"type":"string","description":"cuid-format identifier for this entity.","key$":"submission"},"company":{"type":"string","description":"cuid-format identifier for this entity.","key$":"company"},"company_member":{"type":"string","description":"cuid-format identifier for this entity.","key$":"company_member"}},"required":["additional_information","company","company_member","submission"],"x-ref":"#/components/schemas/DocumentsRequest"}},"*/*":{"schema":{"type":"object","properties":{"additional_information":{"type":"array","items":{"type":"object","properties":{"type":{"type":"string","description":"cuid-format identifier for this entity."},"info":{"type":"string","maxLength":512},"referral_request_reference":{"type":"string","maxLength":512}},"required":["info","type"],"x-ref":"#/components/schemas/AdditionalInformationValidate"},"key$":"additional_information"},"supporting_documents":{"type":"array","items":{"type":"object","properties":{"file":{"type":"string","description":"cuid-format identifier for this entity."},"id":{"type":"string","maxLength":512},"description":{"type":"string","maxLength":512}},"required":["description","file","id"],"x-ref":"#/components/schemas/SupportingDocumentValidate"},"key$":"supporting_documents"},"submission":{"type":"string","description":"cuid-format identifier for this entity.","key$":"submission"},"company":{"type":"string","description":"cuid-format identifier for this entity.","key$":"company"},"company_member":{"type":"string","description":"cuid-format identifier for this entity.","key$":"company_member"}},"required":["additional_information","company","company_member","submission"],"x-ref":"#/components/schemas/DocumentsRequest"}}},"required":true},"parameters":[]}})
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

  const env = envOverride({
    'CUSTOMS_WINDOW_TEST_SUBMISSION_DETAIL_ENTID': idmap,
    'CUSTOMS_WINDOW_TEST_LIVE': 'FALSE',
    'CUSTOMS_WINDOW_TEST_EXPLAIN': 'FALSE',
    'CUSTOMS_WINDOW_APIKEY': '',
  })

  idmap = env['CUSTOMS_WINDOW_TEST_SUBMISSION_DETAIL_ENTID']

  const live = 'TRUE' === env.CUSTOMS_WINDOW_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['CUSTOMS_WINDOW_TEST_SUBMISSION_DETAIL_ENTID']
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
  
