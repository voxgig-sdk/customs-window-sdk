<?php
declare(strict_types=1);

// Typed models for the CustomsWindow SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** BulkUpload entity data model. */
class BulkUpload
{
    public ?string $active_transport_nationality = null;
    public ?string $active_transport_number = null;
    public ?string $arrival_datetime = null;
    public string $client;
    public string $company_member;
    public string $created_at;
    public string $declarant;
    public int $declarations_no;
    public ?string $deleted_at = null;
    public ?string $departure_datetime = null;
    public array $errors_file;
    public ?string $external_id = null;
    public ?int $failed_declarations_no = null;
    public array $file;
    public ?int $green_routed_no = null;
    public ?string $h1_fallback_template = null;
    public ?string $house_transport_doc_ref = null;
    public ?string $id = null;
    public ?string $issue_date = null;
    public ?string $mapping = null;
    public ?int $orange_routed_no = null;
    public ?int $parsed_declarations_no = null;
    public ?string $parser = null;
    public ?string $parsing_completed_at = null;
    public ?string $parsing_started_at = null;
    public ?string $passive_transport_nationality = null;
    public ?string $passive_transport_number = null;
    public ?int $processed_declarations_no = null;
    public ?string $processing_ended_at = null;
    public ?string $processing_started_at = null;
    public ?string $receipt_generating_started_at = null;
    public ?string $receipt_request_started_at = null;
    public ?string $receipt_request_status = null;
    public ?string $receipt_request_user = null;
    public ?string $receipts_zip = null;
    public ?int $red_routed_no = null;
    public ?int $rejected_status_no = null;
    public ?string $status = null;
    public ?string $template = null;
    public string $updated_at;
    public ?int $yellow_routed_no = null;
}

/** Request payload for BulkUpload#load. */
class BulkUploadLoadMatch
{
    public string $id;
}

/** Request payload for BulkUpload#list. */
class BulkUploadListMatch
{
    public ?string $cursor = null;
}

/** Request payload for BulkUpload#create. */
class BulkUploadCreateData
{
    public ?string $active_transport_nationality = null;
    public ?string $active_transport_number = null;
    public ?string $arrival_datetime = null;
    public string $client;
    public string $company_member;
    public string $created_at;
    public string $declarant;
    public int $declarations_no;
    public ?string $deleted_at = null;
    public ?string $departure_datetime = null;
    public array $errors_file;
    public ?string $external_id = null;
    public ?int $failed_declarations_no = null;
    public array $file;
    public ?int $green_routed_no = null;
    public ?string $h1_fallback_template = null;
    public ?string $house_transport_doc_ref = null;
    public ?string $id = null;
    public ?string $issue_date = null;
    public ?string $mapping = null;
    public ?int $orange_routed_no = null;
    public ?int $parsed_declarations_no = null;
    public ?string $parser = null;
    public ?string $parsing_completed_at = null;
    public ?string $parsing_started_at = null;
    public ?string $passive_transport_nationality = null;
    public ?string $passive_transport_number = null;
    public ?int $processed_declarations_no = null;
    public ?string $processing_ended_at = null;
    public ?string $processing_started_at = null;
    public ?string $receipt_generating_started_at = null;
    public ?string $receipt_request_started_at = null;
    public ?string $receipt_request_status = null;
    public ?string $receipt_request_user = null;
    public ?string $receipts_zip = null;
    public ?int $red_routed_no = null;
    public ?int $rejected_status_no = null;
    public ?string $status = null;
    public ?string $template = null;
    public string $updated_at;
    public ?int $yellow_routed_no = null;
}

/** Request payload for BulkUpload#update. */
class BulkUploadUpdateData
{
    public string $id;
    public ?string $active_transport_nationality = null;
    public ?string $active_transport_number = null;
    public ?string $arrival_datetime = null;
    public ?string $client = null;
    public ?string $company_member = null;
    public ?string $created_at = null;
    public ?string $declarant = null;
    public ?int $declarations_no = null;
    public ?string $deleted_at = null;
    public ?string $departure_datetime = null;
    public ?array $errors_file = null;
    public ?string $external_id = null;
    public ?int $failed_declarations_no = null;
    public ?array $file = null;
    public ?int $green_routed_no = null;
    public ?string $h1_fallback_template = null;
    public ?string $house_transport_doc_ref = null;
    public ?string $issue_date = null;
    public ?string $mapping = null;
    public ?int $orange_routed_no = null;
    public ?int $parsed_declarations_no = null;
    public ?string $parser = null;
    public ?string $parsing_completed_at = null;
    public ?string $parsing_started_at = null;
    public ?string $passive_transport_nationality = null;
    public ?string $passive_transport_number = null;
    public ?int $processed_declarations_no = null;
    public ?string $processing_ended_at = null;
    public ?string $processing_started_at = null;
    public ?string $receipt_generating_started_at = null;
    public ?string $receipt_request_started_at = null;
    public ?string $receipt_request_status = null;
    public ?string $receipt_request_user = null;
    public ?string $receipts_zip = null;
    public ?int $red_routed_no = null;
    public ?int $rejected_status_no = null;
    public ?string $status = null;
    public ?string $template = null;
    public ?string $updated_at = null;
    public ?int $yellow_routed_no = null;
}

/** Request payload for BulkUpload#remove. */
class BulkUploadRemoveMatch
{
    public string $id;
}

/** File entity data model. */
class File
{
    public string $company;
    public string $created_at;
    public string $extension;
    public string $file;
    public ?string $id = null;
    public string $name;
    public bool $public;
    public ?int $size = null;
    public string $updated_at;
    public string $url;
}

/** Request payload for File#create. */
class FileCreateData
{
    public string $company;
    public string $created_at;
    public string $extension;
    public string $file;
    public ?string $id = null;
    public string $name;
    public bool $public;
    public ?int $size = null;
    public string $updated_at;
    public string $url;
}

/** PaginatedBulkUploadListList entity data model. */
class PaginatedBulkUploadListList
{
}

/** PaginatedPartyListList entity data model. */
class PaginatedPartyListList
{
}

/** PaginatedSubmissionListList entity data model. */
class PaginatedSubmissionListList
{
}

/** Party entity data model. */
class Party
{
    public array $additional_declaration_type;
    public array $address;
    public array $authorisation;
    public ?string $bank_details = null;
    public array $certificate;
    public string $certificate_type;
    public string $company;
    public string $created_at;
    public array $customs_office_of_lodgement;
    public ?string $deleted_at = null;
    public ?string $email = null;
    public ?string $id = null;
    public ?string $identification_number = null;
    public ?bool $indirect_representative = null;
    public ?string $name = null;
    public ?int $nhd_last_submission_year = null;
    public ?int $nhd_submission_counter = null;
    public ?string $person_paying_customs_duty = null;
    public ?string $phone_country_code = null;
    public ?string $phone_number = null;
    public array $preferred_payment_method;
    public array $signed_form;
    public ?string $type = null;
    public array $type_of_person;
    public ?string $unlocode = null;
    public string $updated_at;
}

/** Request payload for Party#load. */
class PartyLoadMatch
{
    public string $id;
}

/** Request payload for Party#list. */
class PartyListMatch
{
    public ?string $cursor = null;
    public ?string $type = null;
}

/** Request payload for Party#create. */
class PartyCreateData
{
    public array $additional_declaration_type;
    public array $address;
    public array $authorisation;
    public ?string $bank_details = null;
    public array $certificate;
    public string $certificate_type;
    public string $company;
    public string $created_at;
    public array $customs_office_of_lodgement;
    public ?string $deleted_at = null;
    public ?string $email = null;
    public ?string $id = null;
    public ?string $identification_number = null;
    public ?bool $indirect_representative = null;
    public ?string $name = null;
    public ?int $nhd_last_submission_year = null;
    public ?int $nhd_submission_counter = null;
    public ?string $person_paying_customs_duty = null;
    public ?string $phone_country_code = null;
    public ?string $phone_number = null;
    public array $preferred_payment_method;
    public array $signed_form;
    public ?string $type = null;
    public array $type_of_person;
    public ?string $unlocode = null;
    public string $updated_at;
}

/** Request payload for Party#update. */
class PartyUpdateData
{
    public string $id;
    public ?array $additional_declaration_type = null;
    public ?array $address = null;
    public ?array $authorisation = null;
    public ?string $bank_details = null;
    public ?array $certificate = null;
    public ?string $certificate_type = null;
    public ?string $company = null;
    public ?string $created_at = null;
    public ?array $customs_office_of_lodgement = null;
    public ?string $deleted_at = null;
    public ?string $email = null;
    public ?string $identification_number = null;
    public ?bool $indirect_representative = null;
    public ?string $name = null;
    public ?int $nhd_last_submission_year = null;
    public ?int $nhd_submission_counter = null;
    public ?string $person_paying_customs_duty = null;
    public ?string $phone_country_code = null;
    public ?string $phone_number = null;
    public ?array $preferred_payment_method = null;
    public ?array $signed_form = null;
    public ?string $type = null;
    public ?array $type_of_person = null;
    public ?string $unlocode = null;
    public ?string $updated_at = null;
}

/** Request payload for Party#remove. */
class PartyRemoveMatch
{
    public string $id;
}

/** Submission entity data model. */
class Submission
{
    public ?array $additional_external_ids = null;
    public ?string $amendment_reason = null;
    public ?string $amendment_status = null;
    public array $answers;
    public ?bool $bypass_restricted_code = null;
    public array $clearance_slip;
    public array $client;
    public array $company_member;
    public array $consignee;
    public array $consignor;
    public string $created_at;
    public array $declarant;
    public ?string $document_upload_status = null;
    public ?bool $documents_presentation_requested = null;
    public ?bool $documents_upload_requested = null;
    public ?string $external_id = null;
    public string $form;
    public ?string $goods_presentation_status = null;
    public ?string $hrcm_status = null;
    public ?string $id = null;
    public ?string $invalidation_status = null;
    public ?bool $is_global_template = null;
    public ?string $latest_notification_item = null;
    public array $latest_state;
    public ?string $lrn = null;
    public ?string $mrn = null;
    public ?string $name = null;
    public ?bool $partial_answers = null;
    public array $receipt;
    public ?string $refund_application_status = null;
    public ?string $route = null;
    public ?int $shipment_items_no = null;
    public ?int $shipment_items_quantity_no = null;
    public ?string $source = null;
    public ?string $source_type = null;
    public ?string $status = null;
    public ?bool $template = null;
    public ?string $template_id = null;
    public ?array $template_properties = null;
    public ?string $total_tax_amount = null;
    public string $updated_at;
    public ?array $verification_errors = null;
    public ?string $verification_status = null;
}

/** Request payload for Submission#load. */
class SubmissionLoadMatch
{
    public string $id;
}

/** Request payload for Submission#list. */
class SubmissionListMatch
{
    public ?string $cursor = null;
    public ?string $form_subtype = null;
    public ?string $route = null;
    public ?string $status = null;
    public ?bool $template = null;
}

/** Request payload for Submission#create. */
class SubmissionCreateData
{
    public ?array $additional_external_ids = null;
    public ?string $amendment_reason = null;
    public ?string $amendment_status = null;
    public array $answers;
    public ?bool $bypass_restricted_code = null;
    public array $clearance_slip;
    public array $client;
    public array $company_member;
    public array $consignee;
    public array $consignor;
    public string $created_at;
    public array $declarant;
    public ?string $document_upload_status = null;
    public ?bool $documents_presentation_requested = null;
    public ?bool $documents_upload_requested = null;
    public ?string $external_id = null;
    public string $form;
    public ?string $goods_presentation_status = null;
    public ?string $hrcm_status = null;
    public ?string $id = null;
    public ?string $invalidation_status = null;
    public ?bool $is_global_template = null;
    public ?string $latest_notification_item = null;
    public array $latest_state;
    public ?string $lrn = null;
    public ?string $mrn = null;
    public ?string $name = null;
    public ?bool $partial_answers = null;
    public array $receipt;
    public ?string $refund_application_status = null;
    public ?string $route = null;
    public ?int $shipment_items_no = null;
    public ?int $shipment_items_quantity_no = null;
    public ?string $source = null;
    public ?string $source_type = null;
    public ?string $status = null;
    public ?bool $template = null;
    public ?string $template_id = null;
    public ?array $template_properties = null;
    public ?string $total_tax_amount = null;
    public string $updated_at;
    public ?array $verification_errors = null;
    public ?string $verification_status = null;
}

/** Request payload for Submission#update. */
class SubmissionUpdateData
{
    public string $id;
    public ?array $additional_external_ids = null;
    public ?string $amendment_reason = null;
    public ?string $amendment_status = null;
    public ?array $answers = null;
    public ?bool $bypass_restricted_code = null;
    public ?array $clearance_slip = null;
    public ?array $client = null;
    public ?array $company_member = null;
    public ?array $consignee = null;
    public ?array $consignor = null;
    public ?string $created_at = null;
    public ?array $declarant = null;
    public ?string $document_upload_status = null;
    public ?bool $documents_presentation_requested = null;
    public ?bool $documents_upload_requested = null;
    public ?string $external_id = null;
    public ?string $form = null;
    public ?string $goods_presentation_status = null;
    public ?string $hrcm_status = null;
    public ?string $invalidation_status = null;
    public ?bool $is_global_template = null;
    public ?string $latest_notification_item = null;
    public ?array $latest_state = null;
    public ?string $lrn = null;
    public ?string $mrn = null;
    public ?string $name = null;
    public ?bool $partial_answers = null;
    public ?array $receipt = null;
    public ?string $refund_application_status = null;
    public ?string $route = null;
    public ?int $shipment_items_no = null;
    public ?int $shipment_items_quantity_no = null;
    public ?string $source = null;
    public ?string $source_type = null;
    public ?string $status = null;
    public ?bool $template = null;
    public ?string $template_id = null;
    public ?array $template_properties = null;
    public ?string $total_tax_amount = null;
    public ?string $updated_at = null;
    public ?array $verification_errors = null;
    public ?string $verification_status = null;
}

/** Request payload for Submission#remove. */
class SubmissionRemoveMatch
{
    public string $id;
}

/** SubmissionDetail entity data model. */
class SubmissionDetail
{
    public ?array $additional_external_ids = null;
    public array $additional_information;
    public ?string $amendment_status = null;
    public array $clearance_slip;
    public array $client;
    public string $company;
    public array $company_member;
    public array $consignee;
    public array $consignor;
    public string $created_at;
    public array $declarant;
    public ?string $document_upload_status = null;
    public ?bool $documents_presentation_requested = null;
    public ?bool $documents_upload_requested = null;
    public ?string $external_id = null;
    public string $form;
    public ?string $goods_presentation_status = null;
    public ?string $hrcm_status = null;
    public ?string $id = null;
    public ?string $invalidation_status = null;
    public ?bool $is_global_template = null;
    public ?string $latest_notification_item = null;
    public array $latest_state;
    public ?string $lrn = null;
    public ?string $mrn = null;
    public ?string $name = null;
    public array $receipt;
    public ?string $refund_application_status = null;
    public ?string $route = null;
    public ?int $shipment_items_no = null;
    public ?int $shipment_items_quantity_no = null;
    public ?string $source = null;
    public ?string $source_type = null;
    public ?string $status = null;
    public string $submission;
    public ?array $supporting_documents = null;
    public ?bool $template = null;
    public ?string $total_tax_amount = null;
    public string $updated_at;
    public ?array $verification_errors = null;
    public ?string $verification_status = null;
}

/** Request payload for SubmissionDetail#create. */
class SubmissionDetailCreateData
{
    public ?array $additional_external_ids = null;
    public array $additional_information;
    public ?string $amendment_status = null;
    public array $clearance_slip;
    public array $client;
    public string $company;
    public array $company_member;
    public array $consignee;
    public array $consignor;
    public string $created_at;
    public array $declarant;
    public ?string $document_upload_status = null;
    public ?bool $documents_presentation_requested = null;
    public ?bool $documents_upload_requested = null;
    public ?string $external_id = null;
    public string $form;
    public ?string $goods_presentation_status = null;
    public ?string $hrcm_status = null;
    public ?string $id = null;
    public ?string $invalidation_status = null;
    public ?bool $is_global_template = null;
    public ?string $latest_notification_item = null;
    public array $latest_state;
    public ?string $lrn = null;
    public ?string $mrn = null;
    public ?string $name = null;
    public array $receipt;
    public ?string $refund_application_status = null;
    public ?string $route = null;
    public ?int $shipment_items_no = null;
    public ?int $shipment_items_quantity_no = null;
    public ?string $source = null;
    public ?string $source_type = null;
    public ?string $status = null;
    public string $submission;
    public ?array $supporting_documents = null;
    public ?bool $template = null;
    public ?string $total_tax_amount = null;
    public string $updated_at;
    public ?array $verification_errors = null;
    public ?string $verification_status = null;
}

