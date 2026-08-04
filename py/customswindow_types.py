# Typed models for the CustomsWindow SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class BulkUploadRequired(TypedDict):
    client: str
    company_member: str
    created_at: str
    declarant: str
    declarations_no: int
    errors_file: dict
    file: dict
    updated_at: str


class BulkUpload(BulkUploadRequired, total=False):
    active_transport_nationality: str
    active_transport_number: str
    arrival_datetime: str
    deleted_at: str
    departure_datetime: str
    external_id: str
    failed_declarations_no: int
    green_routed_no: int
    h1_fallback_template: str
    house_transport_doc_ref: str
    id: str
    issue_date: str
    mapping: str
    orange_routed_no: int
    parsed_declarations_no: int
    parser: str
    parsing_completed_at: str
    parsing_started_at: str
    passive_transport_nationality: str
    passive_transport_number: str
    processed_declarations_no: int
    processing_ended_at: str
    processing_started_at: str
    receipt_generating_started_at: str
    receipt_request_started_at: str
    receipt_request_status: str
    receipt_request_user: str
    receipts_zip: str
    red_routed_no: int
    rejected_status_no: int
    status: str
    template: str
    yellow_routed_no: int


class BulkUploadLoadMatch(TypedDict):
    id: str


class BulkUploadListMatch(TypedDict, total=False):
    active_transport_nationality: str
    active_transport_number: str
    arrival_datetime: str
    client: str
    company_member: str
    created_at: str
    declarant: str
    declarations_no: int
    deleted_at: str
    departure_datetime: str
    errors_file: dict
    external_id: str
    failed_declarations_no: int
    file: dict
    green_routed_no: int
    h1_fallback_template: str
    house_transport_doc_ref: str
    id: str
    issue_date: str
    mapping: str
    orange_routed_no: int
    parsed_declarations_no: int
    parser: str
    parsing_completed_at: str
    parsing_started_at: str
    passive_transport_nationality: str
    passive_transport_number: str
    processed_declarations_no: int
    processing_ended_at: str
    processing_started_at: str
    receipt_generating_started_at: str
    receipt_request_started_at: str
    receipt_request_status: str
    receipt_request_user: str
    receipts_zip: str
    red_routed_no: int
    rejected_status_no: int
    status: str
    template: str
    updated_at: str
    yellow_routed_no: int


class BulkUploadCreateDataRequired(TypedDict):
    client: str
    company_member: str
    created_at: str
    declarant: str
    declarations_no: int
    errors_file: dict
    file: dict
    updated_at: str


class BulkUploadCreateData(BulkUploadCreateDataRequired, total=False):
    active_transport_nationality: str
    active_transport_number: str
    arrival_datetime: str
    deleted_at: str
    departure_datetime: str
    external_id: str
    failed_declarations_no: int
    green_routed_no: int
    h1_fallback_template: str
    house_transport_doc_ref: str
    id: str
    issue_date: str
    mapping: str
    orange_routed_no: int
    parsed_declarations_no: int
    parser: str
    parsing_completed_at: str
    parsing_started_at: str
    passive_transport_nationality: str
    passive_transport_number: str
    processed_declarations_no: int
    processing_ended_at: str
    processing_started_at: str
    receipt_generating_started_at: str
    receipt_request_started_at: str
    receipt_request_status: str
    receipt_request_user: str
    receipts_zip: str
    red_routed_no: int
    rejected_status_no: int
    status: str
    template: str
    yellow_routed_no: int


class BulkUploadUpdateDataRequired(TypedDict):
    id: str


class BulkUploadUpdateData(BulkUploadUpdateDataRequired, total=False):
    active_transport_nationality: str
    active_transport_number: str
    arrival_datetime: str
    client: str
    company_member: str
    created_at: str
    declarant: str
    declarations_no: int
    deleted_at: str
    departure_datetime: str
    errors_file: dict
    external_id: str
    failed_declarations_no: int
    file: dict
    green_routed_no: int
    h1_fallback_template: str
    house_transport_doc_ref: str
    issue_date: str
    mapping: str
    orange_routed_no: int
    parsed_declarations_no: int
    parser: str
    parsing_completed_at: str
    parsing_started_at: str
    passive_transport_nationality: str
    passive_transport_number: str
    processed_declarations_no: int
    processing_ended_at: str
    processing_started_at: str
    receipt_generating_started_at: str
    receipt_request_started_at: str
    receipt_request_status: str
    receipt_request_user: str
    receipts_zip: str
    red_routed_no: int
    rejected_status_no: int
    status: str
    template: str
    updated_at: str
    yellow_routed_no: int


class BulkUploadRemoveMatch(TypedDict):
    id: str


class FileRequired(TypedDict):
    company: str
    created_at: str
    extension: str
    file: str
    name: str
    public: bool
    updated_at: str
    url: str


class File(FileRequired, total=False):
    id: str
    size: int


class FileCreateDataRequired(TypedDict):
    company: str
    created_at: str
    extension: str
    file: str
    name: str
    public: bool
    updated_at: str
    url: str


class FileCreateData(FileCreateDataRequired, total=False):
    id: str
    size: int


class PaginatedBulkUploadListList(TypedDict):
    pass


class PaginatedPartyListList(TypedDict):
    pass


class PaginatedSubmissionListList(TypedDict):
    pass


class PartyRequired(TypedDict):
    additional_declaration_type: dict
    address: dict
    authorisation: dict
    certificate: dict
    certificate_type: str
    company: str
    created_at: str
    customs_office_of_lodgement: dict
    preferred_payment_method: dict
    signed_form: dict
    type_of_person: dict
    updated_at: str


class Party(PartyRequired, total=False):
    bank_detail: str
    deleted_at: str
    email: str
    id: str
    identification_number: str
    indirect_representative: bool
    name: str
    nhd_last_submission_year: int
    nhd_submission_counter: int
    person_paying_customs_duty: str
    phone_country_code: str
    phone_number: str
    type: str
    unlocode: str


class PartyLoadMatch(TypedDict):
    id: str


class PartyListMatch(TypedDict, total=False):
    additional_declaration_type: dict
    address: dict
    authorisation: dict
    bank_detail: str
    certificate: dict
    certificate_type: str
    company: str
    created_at: str
    customs_office_of_lodgement: dict
    deleted_at: str
    email: str
    id: str
    identification_number: str
    indirect_representative: bool
    name: str
    nhd_last_submission_year: int
    nhd_submission_counter: int
    person_paying_customs_duty: str
    phone_country_code: str
    phone_number: str
    preferred_payment_method: dict
    signed_form: dict
    type: str
    type_of_person: dict
    unlocode: str
    updated_at: str


class PartyCreateDataRequired(TypedDict):
    additional_declaration_type: dict
    address: dict
    authorisation: dict
    certificate: dict
    certificate_type: str
    company: str
    created_at: str
    customs_office_of_lodgement: dict
    preferred_payment_method: dict
    signed_form: dict
    type_of_person: dict
    updated_at: str


class PartyCreateData(PartyCreateDataRequired, total=False):
    bank_detail: str
    deleted_at: str
    email: str
    id: str
    identification_number: str
    indirect_representative: bool
    name: str
    nhd_last_submission_year: int
    nhd_submission_counter: int
    person_paying_customs_duty: str
    phone_country_code: str
    phone_number: str
    type: str
    unlocode: str


class PartyUpdateDataRequired(TypedDict):
    id: str


class PartyUpdateData(PartyUpdateDataRequired, total=False):
    additional_declaration_type: dict
    address: dict
    authorisation: dict
    bank_detail: str
    certificate: dict
    certificate_type: str
    company: str
    created_at: str
    customs_office_of_lodgement: dict
    deleted_at: str
    email: str
    identification_number: str
    indirect_representative: bool
    name: str
    nhd_last_submission_year: int
    nhd_submission_counter: int
    person_paying_customs_duty: str
    phone_country_code: str
    phone_number: str
    preferred_payment_method: dict
    signed_form: dict
    type: str
    type_of_person: dict
    unlocode: str
    updated_at: str


class PartyRemoveMatch(TypedDict):
    id: str


class SubmissionRequired(TypedDict):
    answer: list
    clearance_slip: dict
    client: dict
    company_member: dict
    consignee: dict
    consignor: dict
    created_at: str
    declarant: dict
    form: str
    latest_state: dict
    receipt: dict
    updated_at: str


class Submission(SubmissionRequired, total=False):
    additional_external_id: list
    amendment_reason: str
    amendment_status: str
    bypass_restricted_code: bool
    document_upload_status: str
    documents_presentation_requested: bool
    documents_upload_requested: bool
    external_id: str
    goods_presentation_status: str
    hrcm_status: str
    id: str
    invalidation_status: str
    is_global_template: bool
    latest_notification_item: str
    lrn: str
    mrn: str
    name: str
    partial_answer: bool
    refund_application_status: str
    route: str
    shipment_items_no: int
    shipment_items_quantity_no: int
    source: str
    source_type: str
    status: str
    template: bool
    template_id: str
    template_property: list
    total_tax_amount: str
    verification_error: list
    verification_status: str


class SubmissionLoadMatch(TypedDict):
    id: str


class SubmissionListMatch(TypedDict, total=False):
    additional_external_id: list
    amendment_reason: str
    amendment_status: str
    answer: list
    bypass_restricted_code: bool
    clearance_slip: dict
    client: dict
    company_member: dict
    consignee: dict
    consignor: dict
    created_at: str
    declarant: dict
    document_upload_status: str
    documents_presentation_requested: bool
    documents_upload_requested: bool
    external_id: str
    form: str
    goods_presentation_status: str
    hrcm_status: str
    id: str
    invalidation_status: str
    is_global_template: bool
    latest_notification_item: str
    latest_state: dict
    lrn: str
    mrn: str
    name: str
    partial_answer: bool
    receipt: dict
    refund_application_status: str
    route: str
    shipment_items_no: int
    shipment_items_quantity_no: int
    source: str
    source_type: str
    status: str
    template: bool
    template_id: str
    template_property: list
    total_tax_amount: str
    updated_at: str
    verification_error: list
    verification_status: str


class SubmissionCreateDataRequired(TypedDict):
    answer: list
    clearance_slip: dict
    client: dict
    company_member: dict
    consignee: dict
    consignor: dict
    created_at: str
    declarant: dict
    form: str
    latest_state: dict
    receipt: dict
    updated_at: str


class SubmissionCreateData(SubmissionCreateDataRequired, total=False):
    additional_external_id: list
    amendment_reason: str
    amendment_status: str
    bypass_restricted_code: bool
    document_upload_status: str
    documents_presentation_requested: bool
    documents_upload_requested: bool
    external_id: str
    goods_presentation_status: str
    hrcm_status: str
    id: str
    invalidation_status: str
    is_global_template: bool
    latest_notification_item: str
    lrn: str
    mrn: str
    name: str
    partial_answer: bool
    refund_application_status: str
    route: str
    shipment_items_no: int
    shipment_items_quantity_no: int
    source: str
    source_type: str
    status: str
    template: bool
    template_id: str
    template_property: list
    total_tax_amount: str
    verification_error: list
    verification_status: str


class SubmissionUpdateDataRequired(TypedDict):
    id: str


class SubmissionUpdateData(SubmissionUpdateDataRequired, total=False):
    additional_external_id: list
    amendment_reason: str
    amendment_status: str
    answer: list
    bypass_restricted_code: bool
    clearance_slip: dict
    client: dict
    company_member: dict
    consignee: dict
    consignor: dict
    created_at: str
    declarant: dict
    document_upload_status: str
    documents_presentation_requested: bool
    documents_upload_requested: bool
    external_id: str
    form: str
    goods_presentation_status: str
    hrcm_status: str
    invalidation_status: str
    is_global_template: bool
    latest_notification_item: str
    latest_state: dict
    lrn: str
    mrn: str
    name: str
    partial_answer: bool
    receipt: dict
    refund_application_status: str
    route: str
    shipment_items_no: int
    shipment_items_quantity_no: int
    source: str
    source_type: str
    status: str
    template: bool
    template_id: str
    template_property: list
    total_tax_amount: str
    updated_at: str
    verification_error: list
    verification_status: str


class SubmissionRemoveMatch(TypedDict):
    id: str


class SubmissionDetailRequired(TypedDict):
    additional_information: list
    clearance_slip: dict
    client: dict
    company: str
    company_member: dict
    consignee: dict
    consignor: dict
    created_at: str
    declarant: dict
    form: str
    latest_state: dict
    receipt: dict
    submission: str
    updated_at: str


class SubmissionDetail(SubmissionDetailRequired, total=False):
    additional_external_id: list
    amendment_status: str
    document_upload_status: str
    documents_presentation_requested: bool
    documents_upload_requested: bool
    external_id: str
    goods_presentation_status: str
    hrcm_status: str
    id: str
    invalidation_status: str
    is_global_template: bool
    latest_notification_item: str
    lrn: str
    mrn: str
    name: str
    refund_application_status: str
    route: str
    shipment_items_no: int
    shipment_items_quantity_no: int
    source: str
    source_type: str
    status: str
    supporting_document: list
    template: bool
    total_tax_amount: str
    verification_error: list
    verification_status: str


class SubmissionDetailCreateDataRequired(TypedDict):
    additional_information: list
    clearance_slip: dict
    client: dict
    company: str
    company_member: dict
    consignee: dict
    consignor: dict
    created_at: str
    declarant: dict
    form: str
    latest_state: dict
    receipt: dict
    submission: str
    updated_at: str


class SubmissionDetailCreateData(SubmissionDetailCreateDataRequired, total=False):
    additional_external_id: list
    amendment_status: str
    document_upload_status: str
    documents_presentation_requested: bool
    documents_upload_requested: bool
    external_id: str
    goods_presentation_status: str
    hrcm_status: str
    id: str
    invalidation_status: str
    is_global_template: bool
    latest_notification_item: str
    lrn: str
    mrn: str
    name: str
    refund_application_status: str
    route: str
    shipment_items_no: int
    shipment_items_quantity_no: int
    source: str
    source_type: str
    status: str
    supporting_document: list
    template: bool
    total_tax_amount: str
    verification_error: list
    verification_status: str
