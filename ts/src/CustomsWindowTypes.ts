// Typed models for the CustomsWindow SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface BulkUpload {
  active_transport_nationality?: string
  active_transport_number?: string
  arrival_datetime?: string
  client: string
  company_member: string
  created_at: string
  declarant: string
  declarations_no: number
  deleted_at?: string
  departure_datetime?: string
  errors_file: Record<string, any>
  external_id?: string
  failed_declarations_no?: number
  file: Record<string, any>
  green_routed_no?: number
  h1_fallback_template?: string
  house_transport_doc_ref?: string
  id?: string
  issue_date?: string
  mapping?: string
  orange_routed_no?: number
  parsed_declarations_no?: number
  parser?: string
  parsing_completed_at?: string
  parsing_started_at?: string
  passive_transport_nationality?: string
  passive_transport_number?: string
  processed_declarations_no?: number
  processing_ended_at?: string
  processing_started_at?: string
  receipt_generating_started_at?: string
  receipt_request_started_at?: string
  receipt_request_status?: string
  receipt_request_user?: string
  receipts_zip?: string
  red_routed_no?: number
  rejected_status_no?: number
  status?: string
  template?: string
  updated_at: string
  yellow_routed_no?: number
}

export interface BulkUploadLoadMatch {
  id: string

  // Selects a custom action instead of the plain load:
  //   'generate_pdf'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface BulkUploadListMatch {
  cursor?: string
}

export interface BulkUploadCreateData {
  active_transport_nationality?: string
  active_transport_number?: string
  arrival_datetime?: string
  client: string
  company_member: string
  created_at: string
  declarant: string
  declarations_no: number
  deleted_at?: string
  departure_datetime?: string
  errors_file: Record<string, any>
  external_id?: string
  failed_declarations_no?: number
  file: Record<string, any>
  green_routed_no?: number
  h1_fallback_template?: string
  house_transport_doc_ref?: string
  id?: string
  issue_date?: string
  mapping?: string
  orange_routed_no?: number
  parsed_declarations_no?: number
  parser?: string
  parsing_completed_at?: string
  parsing_started_at?: string
  passive_transport_nationality?: string
  passive_transport_number?: string
  processed_declarations_no?: number
  processing_ended_at?: string
  processing_started_at?: string
  receipt_generating_started_at?: string
  receipt_request_started_at?: string
  receipt_request_status?: string
  receipt_request_user?: string
  receipts_zip?: string
  red_routed_no?: number
  rejected_status_no?: number
  status?: string
  template?: string
  updated_at: string
  yellow_routed_no?: number
}

export interface BulkUploadUpdateData {
  id: string
  active_transport_nationality?: string
  active_transport_number?: string
  arrival_datetime?: string
  client?: string
  company_member?: string
  created_at?: string
  declarant?: string
  declarations_no?: number
  deleted_at?: string
  departure_datetime?: string
  errors_file?: Record<string, any>
  external_id?: string
  failed_declarations_no?: number
  file?: Record<string, any>
  green_routed_no?: number
  h1_fallback_template?: string
  house_transport_doc_ref?: string
  issue_date?: string
  mapping?: string
  orange_routed_no?: number
  parsed_declarations_no?: number
  parser?: string
  parsing_completed_at?: string
  parsing_started_at?: string
  passive_transport_nationality?: string
  passive_transport_number?: string
  processed_declarations_no?: number
  processing_ended_at?: string
  processing_started_at?: string
  receipt_generating_started_at?: string
  receipt_request_started_at?: string
  receipt_request_status?: string
  receipt_request_user?: string
  receipts_zip?: string
  red_routed_no?: number
  rejected_status_no?: number
  status?: string
  template?: string
  updated_at?: string
  yellow_routed_no?: number

  // Selects a custom action instead of the plain update:
  //   'generate_pdf'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface BulkUploadRemoveMatch {
  id: string

  // Selects a custom action instead of the plain remove:
  //   'generate_pdf'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface File {
  company: string
  created_at: string
  extension: string
  file: string
  id?: string
  name: string
  public: boolean
  size?: number
  updated_at: string
  url: string
}

export interface FileCreateData {
  company: string
  created_at: string
  extension: string
  file: string
  id?: string
  name: string
  public: boolean
  size?: number
  updated_at: string
  url: string
}

export interface PaginatedBulkUploadListList {
}

export interface PaginatedPartyListList {
}

export interface PaginatedSubmissionListList {
}

export interface Party {
  additional_declaration_type: Record<string, any>
  address: Record<string, any>
  authorisation: Record<string, any>
  bank_details?: string
  certificate: Record<string, any>
  certificate_type: string
  company: string
  created_at: string
  customs_office_of_lodgement: Record<string, any>
  deleted_at?: string
  email?: string
  id?: string
  identification_number?: string
  indirect_representative?: boolean
  name?: string
  nhd_last_submission_year?: number
  nhd_submission_counter?: number
  person_paying_customs_duty?: string
  phone_country_code?: string
  phone_number?: string
  preferred_payment_method: Record<string, any>
  signed_form: Record<string, any>
  type?: string
  type_of_person: Record<string, any>
  unlocode?: string
  updated_at: string
}

export interface PartyLoadMatch {
  id: string
}

export interface PartyListMatch {
  cursor?: string
  type?: string
}

export interface PartyCreateData {
  additional_declaration_type: Record<string, any>
  address: Record<string, any>
  authorisation: Record<string, any>
  bank_details?: string
  certificate: Record<string, any>
  certificate_type: string
  company: string
  created_at: string
  customs_office_of_lodgement: Record<string, any>
  deleted_at?: string
  email?: string
  id?: string
  identification_number?: string
  indirect_representative?: boolean
  name?: string
  nhd_last_submission_year?: number
  nhd_submission_counter?: number
  person_paying_customs_duty?: string
  phone_country_code?: string
  phone_number?: string
  preferred_payment_method: Record<string, any>
  signed_form: Record<string, any>
  type?: string
  type_of_person: Record<string, any>
  unlocode?: string
  updated_at: string
}

export interface PartyUpdateData {
  id: string
  additional_declaration_type?: Record<string, any>
  address?: Record<string, any>
  authorisation?: Record<string, any>
  bank_details?: string
  certificate?: Record<string, any>
  certificate_type?: string
  company?: string
  created_at?: string
  customs_office_of_lodgement?: Record<string, any>
  deleted_at?: string
  email?: string
  identification_number?: string
  indirect_representative?: boolean
  name?: string
  nhd_last_submission_year?: number
  nhd_submission_counter?: number
  person_paying_customs_duty?: string
  phone_country_code?: string
  phone_number?: string
  preferred_payment_method?: Record<string, any>
  signed_form?: Record<string, any>
  type?: string
  type_of_person?: Record<string, any>
  unlocode?: string
  updated_at?: string
}

export interface PartyRemoveMatch {
  id: string
}

export interface Submission {
  additional_external_ids?: any[]
  amendment_reason?: string
  amendment_status?: string
  answers: any[]
  bypass_restricted_code?: boolean
  clearance_slip: Record<string, any>
  client: Record<string, any>
  company_member: Record<string, any>
  consignee: Record<string, any>
  consignor: Record<string, any>
  created_at: string
  declarant: Record<string, any>
  document_upload_status?: string
  documents_presentation_requested?: boolean
  documents_upload_requested?: boolean
  external_id?: string
  form: string
  goods_presentation_status?: string
  hrcm_status?: string
  id?: string
  invalidation_status?: string
  is_global_template?: boolean
  latest_notification_item?: string
  latest_state: Record<string, any>
  lrn?: string
  mrn?: string
  name?: string
  partial_answers?: boolean
  receipt: Record<string, any>
  refund_application_status?: string
  route?: string
  shipment_items_no?: number
  shipment_items_quantity_no?: number
  source?: string
  source_type?: string
  status?: string
  template?: boolean
  template_id?: string
  template_properties?: any[]
  total_tax_amount?: string
  updated_at: string
  verification_errors?: any[]
  verification_status?: string
}

export interface SubmissionLoadMatch {
  id: string

  // Selects a custom action instead of the plain load:
  //   'clearance_slip' | 'notification_read' | 'pbn_applicable' | 'receipt' | 'refund' | 'retrieve'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface SubmissionListMatch {
  cursor?: string
  form_subtype?: string
  route?: string
  status?: string
  template?: boolean
}

export interface SubmissionCreateData {
  additional_external_ids?: any[]
  amendment_reason?: string
  amendment_status?: string
  answers: any[]
  bypass_restricted_code?: boolean
  clearance_slip: Record<string, any>
  client: Record<string, any>
  company_member: Record<string, any>
  consignee: Record<string, any>
  consignor: Record<string, any>
  created_at: string
  declarant: Record<string, any>
  document_upload_status?: string
  documents_presentation_requested?: boolean
  documents_upload_requested?: boolean
  external_id?: string
  form: string
  goods_presentation_status?: string
  hrcm_status?: string
  id?: string
  invalidation_status?: string
  is_global_template?: boolean
  latest_notification_item?: string
  latest_state: Record<string, any>
  lrn?: string
  mrn?: string
  name?: string
  partial_answers?: boolean
  receipt: Record<string, any>
  refund_application_status?: string
  route?: string
  shipment_items_no?: number
  shipment_items_quantity_no?: number
  source?: string
  source_type?: string
  status?: string
  template?: boolean
  template_id?: string
  template_properties?: any[]
  total_tax_amount?: string
  updated_at: string
  verification_errors?: any[]
  verification_status?: string

  // Selects a custom action instead of the plain create:
  //   'refund' | 'retrieve'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface SubmissionUpdateData {
  id: string
  additional_external_ids?: any[]
  amendment_reason?: string
  amendment_status?: string
  answers?: any[]
  bypass_restricted_code?: boolean
  clearance_slip?: Record<string, any>
  client?: Record<string, any>
  company_member?: Record<string, any>
  consignee?: Record<string, any>
  consignor?: Record<string, any>
  created_at?: string
  declarant?: Record<string, any>
  document_upload_status?: string
  documents_presentation_requested?: boolean
  documents_upload_requested?: boolean
  external_id?: string
  form?: string
  goods_presentation_status?: string
  hrcm_status?: string
  invalidation_status?: string
  is_global_template?: boolean
  latest_notification_item?: string
  latest_state?: Record<string, any>
  lrn?: string
  mrn?: string
  name?: string
  partial_answers?: boolean
  receipt?: Record<string, any>
  refund_application_status?: string
  route?: string
  shipment_items_no?: number
  shipment_items_quantity_no?: number
  source?: string
  source_type?: string
  status?: string
  template?: boolean
  template_id?: string
  template_properties?: any[]
  total_tax_amount?: string
  updated_at?: string
  verification_errors?: any[]
  verification_status?: string

  // Selects a custom action instead of the plain update:
  //   'clearance_slip' | 'notification_read' | 'pbn_applicable' | 'receipt'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface SubmissionRemoveMatch {
  id: string

  // Selects a custom action instead of the plain remove:
  //   'clearance_slip' | 'notification_read' | 'pbn_applicable' | 'receipt'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface SubmissionDetail {
  additional_external_ids?: any[]
  additional_information: any[]
  amendment_status?: string
  clearance_slip: Record<string, any>
  client: Record<string, any>
  company: string
  company_member: Record<string, any>
  consignee: Record<string, any>
  consignor: Record<string, any>
  created_at: string
  declarant: Record<string, any>
  document_upload_status?: string
  documents_presentation_requested?: boolean
  documents_upload_requested?: boolean
  external_id?: string
  form: string
  goods_presentation_status?: string
  hrcm_status?: string
  id?: string
  invalidation_status?: string
  is_global_template?: boolean
  latest_notification_item?: string
  latest_state: Record<string, any>
  lrn?: string
  mrn?: string
  name?: string
  receipt: Record<string, any>
  refund_application_status?: string
  route?: string
  shipment_items_no?: number
  shipment_items_quantity_no?: number
  source?: string
  source_type?: string
  status?: string
  submission: string
  supporting_documents?: any[]
  template?: boolean
  total_tax_amount?: string
  updated_at: string
  verification_errors?: any[]
  verification_status?: string
}

export interface SubmissionDetailCreateData {
  additional_external_ids?: any[]
  additional_information: any[]
  amendment_status?: string
  clearance_slip: Record<string, any>
  client: Record<string, any>
  company: string
  company_member: Record<string, any>
  consignee: Record<string, any>
  consignor: Record<string, any>
  created_at: string
  declarant: Record<string, any>
  document_upload_status?: string
  documents_presentation_requested?: boolean
  documents_upload_requested?: boolean
  external_id?: string
  form: string
  goods_presentation_status?: string
  hrcm_status?: string
  id?: string
  invalidation_status?: string
  is_global_template?: boolean
  latest_notification_item?: string
  latest_state: Record<string, any>
  lrn?: string
  mrn?: string
  name?: string
  receipt: Record<string, any>
  refund_application_status?: string
  route?: string
  shipment_items_no?: number
  shipment_items_quantity_no?: number
  source?: string
  source_type?: string
  status?: string
  submission: string
  supporting_documents?: any[]
  template?: boolean
  total_tax_amount?: string
  updated_at: string
  verification_errors?: any[]
  verification_status?: string
}

