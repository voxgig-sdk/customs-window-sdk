# frozen_string_literal: true

# Typed models for the CustomsWindow SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# BulkUpload entity data model.
#
# @!attribute [rw] active_transport_nationality
#   @return [String, nil]
#
# @!attribute [rw] active_transport_number
#   @return [String, nil]
#
# @!attribute [rw] arrival_datetime
#   @return [String, nil]
#
# @!attribute [rw] client
#   @return [String]
#
# @!attribute [rw] company_member
#   @return [String]
#
# @!attribute [rw] created_at
#   @return [String]
#
# @!attribute [rw] declarant
#   @return [String]
#
# @!attribute [rw] declarations_no
#   @return [Integer]
#
# @!attribute [rw] deleted_at
#   @return [String, nil]
#
# @!attribute [rw] departure_datetime
#   @return [String, nil]
#
# @!attribute [rw] errors_file
#   @return [Hash]
#
# @!attribute [rw] external_id
#   @return [String, nil]
#
# @!attribute [rw] failed_declarations_no
#   @return [Integer, nil]
#
# @!attribute [rw] file
#   @return [Hash]
#
# @!attribute [rw] green_routed_no
#   @return [Integer, nil]
#
# @!attribute [rw] h1_fallback_template
#   @return [String, nil]
#
# @!attribute [rw] house_transport_doc_ref
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] issue_date
#   @return [String, nil]
#
# @!attribute [rw] mapping
#   @return [String, nil]
#
# @!attribute [rw] orange_routed_no
#   @return [Integer, nil]
#
# @!attribute [rw] parsed_declarations_no
#   @return [Integer, nil]
#
# @!attribute [rw] parser
#   @return [String, nil]
#
# @!attribute [rw] parsing_completed_at
#   @return [String, nil]
#
# @!attribute [rw] parsing_started_at
#   @return [String, nil]
#
# @!attribute [rw] passive_transport_nationality
#   @return [String, nil]
#
# @!attribute [rw] passive_transport_number
#   @return [String, nil]
#
# @!attribute [rw] processed_declarations_no
#   @return [Integer, nil]
#
# @!attribute [rw] processing_ended_at
#   @return [String, nil]
#
# @!attribute [rw] processing_started_at
#   @return [String, nil]
#
# @!attribute [rw] receipt_generating_started_at
#   @return [String, nil]
#
# @!attribute [rw] receipt_request_started_at
#   @return [String, nil]
#
# @!attribute [rw] receipt_request_status
#   @return [String, nil]
#
# @!attribute [rw] receipt_request_user
#   @return [String, nil]
#
# @!attribute [rw] receipts_zip
#   @return [String, nil]
#
# @!attribute [rw] red_routed_no
#   @return [Integer, nil]
#
# @!attribute [rw] rejected_status_no
#   @return [Integer, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] template
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String]
#
# @!attribute [rw] yellow_routed_no
#   @return [Integer, nil]
BulkUpload = Struct.new(
  :active_transport_nationality,
  :active_transport_number,
  :arrival_datetime,
  :client,
  :company_member,
  :created_at,
  :declarant,
  :declarations_no,
  :deleted_at,
  :departure_datetime,
  :errors_file,
  :external_id,
  :failed_declarations_no,
  :file,
  :green_routed_no,
  :h1_fallback_template,
  :house_transport_doc_ref,
  :id,
  :issue_date,
  :mapping,
  :orange_routed_no,
  :parsed_declarations_no,
  :parser,
  :parsing_completed_at,
  :parsing_started_at,
  :passive_transport_nationality,
  :passive_transport_number,
  :processed_declarations_no,
  :processing_ended_at,
  :processing_started_at,
  :receipt_generating_started_at,
  :receipt_request_started_at,
  :receipt_request_status,
  :receipt_request_user,
  :receipts_zip,
  :red_routed_no,
  :rejected_status_no,
  :status,
  :template,
  :updated_at,
  :yellow_routed_no,
  keyword_init: true
)

# Request payload for BulkUpload#load.
#
# @!attribute [rw] id
#   @return [String]
BulkUploadLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for BulkUpload#list.
#
# @!attribute [rw] active_transport_nationality
#   @return [String, nil]
#
# @!attribute [rw] active_transport_number
#   @return [String, nil]
#
# @!attribute [rw] arrival_datetime
#   @return [String, nil]
#
# @!attribute [rw] client
#   @return [String, nil]
#
# @!attribute [rw] company_member
#   @return [String, nil]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] declarant
#   @return [String, nil]
#
# @!attribute [rw] declarations_no
#   @return [Integer, nil]
#
# @!attribute [rw] deleted_at
#   @return [String, nil]
#
# @!attribute [rw] departure_datetime
#   @return [String, nil]
#
# @!attribute [rw] errors_file
#   @return [Hash, nil]
#
# @!attribute [rw] external_id
#   @return [String, nil]
#
# @!attribute [rw] failed_declarations_no
#   @return [Integer, nil]
#
# @!attribute [rw] file
#   @return [Hash, nil]
#
# @!attribute [rw] green_routed_no
#   @return [Integer, nil]
#
# @!attribute [rw] h1_fallback_template
#   @return [String, nil]
#
# @!attribute [rw] house_transport_doc_ref
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] issue_date
#   @return [String, nil]
#
# @!attribute [rw] mapping
#   @return [String, nil]
#
# @!attribute [rw] orange_routed_no
#   @return [Integer, nil]
#
# @!attribute [rw] parsed_declarations_no
#   @return [Integer, nil]
#
# @!attribute [rw] parser
#   @return [String, nil]
#
# @!attribute [rw] parsing_completed_at
#   @return [String, nil]
#
# @!attribute [rw] parsing_started_at
#   @return [String, nil]
#
# @!attribute [rw] passive_transport_nationality
#   @return [String, nil]
#
# @!attribute [rw] passive_transport_number
#   @return [String, nil]
#
# @!attribute [rw] processed_declarations_no
#   @return [Integer, nil]
#
# @!attribute [rw] processing_ended_at
#   @return [String, nil]
#
# @!attribute [rw] processing_started_at
#   @return [String, nil]
#
# @!attribute [rw] receipt_generating_started_at
#   @return [String, nil]
#
# @!attribute [rw] receipt_request_started_at
#   @return [String, nil]
#
# @!attribute [rw] receipt_request_status
#   @return [String, nil]
#
# @!attribute [rw] receipt_request_user
#   @return [String, nil]
#
# @!attribute [rw] receipts_zip
#   @return [String, nil]
#
# @!attribute [rw] red_routed_no
#   @return [Integer, nil]
#
# @!attribute [rw] rejected_status_no
#   @return [Integer, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] template
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
#
# @!attribute [rw] yellow_routed_no
#   @return [Integer, nil]
BulkUploadListMatch = Struct.new(
  :active_transport_nationality,
  :active_transport_number,
  :arrival_datetime,
  :client,
  :company_member,
  :created_at,
  :declarant,
  :declarations_no,
  :deleted_at,
  :departure_datetime,
  :errors_file,
  :external_id,
  :failed_declarations_no,
  :file,
  :green_routed_no,
  :h1_fallback_template,
  :house_transport_doc_ref,
  :id,
  :issue_date,
  :mapping,
  :orange_routed_no,
  :parsed_declarations_no,
  :parser,
  :parsing_completed_at,
  :parsing_started_at,
  :passive_transport_nationality,
  :passive_transport_number,
  :processed_declarations_no,
  :processing_ended_at,
  :processing_started_at,
  :receipt_generating_started_at,
  :receipt_request_started_at,
  :receipt_request_status,
  :receipt_request_user,
  :receipts_zip,
  :red_routed_no,
  :rejected_status_no,
  :status,
  :template,
  :updated_at,
  :yellow_routed_no,
  keyword_init: true
)

# Request payload for BulkUpload#create.
#
# @!attribute [rw] active_transport_nationality
#   @return [String, nil]
#
# @!attribute [rw] active_transport_number
#   @return [String, nil]
#
# @!attribute [rw] arrival_datetime
#   @return [String, nil]
#
# @!attribute [rw] client
#   @return [String]
#
# @!attribute [rw] company_member
#   @return [String]
#
# @!attribute [rw] created_at
#   @return [String]
#
# @!attribute [rw] declarant
#   @return [String]
#
# @!attribute [rw] declarations_no
#   @return [Integer]
#
# @!attribute [rw] deleted_at
#   @return [String, nil]
#
# @!attribute [rw] departure_datetime
#   @return [String, nil]
#
# @!attribute [rw] errors_file
#   @return [Hash]
#
# @!attribute [rw] external_id
#   @return [String, nil]
#
# @!attribute [rw] failed_declarations_no
#   @return [Integer, nil]
#
# @!attribute [rw] file
#   @return [Hash]
#
# @!attribute [rw] green_routed_no
#   @return [Integer, nil]
#
# @!attribute [rw] h1_fallback_template
#   @return [String, nil]
#
# @!attribute [rw] house_transport_doc_ref
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] issue_date
#   @return [String, nil]
#
# @!attribute [rw] mapping
#   @return [String, nil]
#
# @!attribute [rw] orange_routed_no
#   @return [Integer, nil]
#
# @!attribute [rw] parsed_declarations_no
#   @return [Integer, nil]
#
# @!attribute [rw] parser
#   @return [String, nil]
#
# @!attribute [rw] parsing_completed_at
#   @return [String, nil]
#
# @!attribute [rw] parsing_started_at
#   @return [String, nil]
#
# @!attribute [rw] passive_transport_nationality
#   @return [String, nil]
#
# @!attribute [rw] passive_transport_number
#   @return [String, nil]
#
# @!attribute [rw] processed_declarations_no
#   @return [Integer, nil]
#
# @!attribute [rw] processing_ended_at
#   @return [String, nil]
#
# @!attribute [rw] processing_started_at
#   @return [String, nil]
#
# @!attribute [rw] receipt_generating_started_at
#   @return [String, nil]
#
# @!attribute [rw] receipt_request_started_at
#   @return [String, nil]
#
# @!attribute [rw] receipt_request_status
#   @return [String, nil]
#
# @!attribute [rw] receipt_request_user
#   @return [String, nil]
#
# @!attribute [rw] receipts_zip
#   @return [String, nil]
#
# @!attribute [rw] red_routed_no
#   @return [Integer, nil]
#
# @!attribute [rw] rejected_status_no
#   @return [Integer, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] template
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String]
#
# @!attribute [rw] yellow_routed_no
#   @return [Integer, nil]
BulkUploadCreateData = Struct.new(
  :active_transport_nationality,
  :active_transport_number,
  :arrival_datetime,
  :client,
  :company_member,
  :created_at,
  :declarant,
  :declarations_no,
  :deleted_at,
  :departure_datetime,
  :errors_file,
  :external_id,
  :failed_declarations_no,
  :file,
  :green_routed_no,
  :h1_fallback_template,
  :house_transport_doc_ref,
  :id,
  :issue_date,
  :mapping,
  :orange_routed_no,
  :parsed_declarations_no,
  :parser,
  :parsing_completed_at,
  :parsing_started_at,
  :passive_transport_nationality,
  :passive_transport_number,
  :processed_declarations_no,
  :processing_ended_at,
  :processing_started_at,
  :receipt_generating_started_at,
  :receipt_request_started_at,
  :receipt_request_status,
  :receipt_request_user,
  :receipts_zip,
  :red_routed_no,
  :rejected_status_no,
  :status,
  :template,
  :updated_at,
  :yellow_routed_no,
  keyword_init: true
)

# Request payload for BulkUpload#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] active_transport_nationality
#   @return [String, nil]
#
# @!attribute [rw] active_transport_number
#   @return [String, nil]
#
# @!attribute [rw] arrival_datetime
#   @return [String, nil]
#
# @!attribute [rw] client
#   @return [String, nil]
#
# @!attribute [rw] company_member
#   @return [String, nil]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] declarant
#   @return [String, nil]
#
# @!attribute [rw] declarations_no
#   @return [Integer, nil]
#
# @!attribute [rw] deleted_at
#   @return [String, nil]
#
# @!attribute [rw] departure_datetime
#   @return [String, nil]
#
# @!attribute [rw] errors_file
#   @return [Hash, nil]
#
# @!attribute [rw] external_id
#   @return [String, nil]
#
# @!attribute [rw] failed_declarations_no
#   @return [Integer, nil]
#
# @!attribute [rw] file
#   @return [Hash, nil]
#
# @!attribute [rw] green_routed_no
#   @return [Integer, nil]
#
# @!attribute [rw] h1_fallback_template
#   @return [String, nil]
#
# @!attribute [rw] house_transport_doc_ref
#   @return [String, nil]
#
# @!attribute [rw] issue_date
#   @return [String, nil]
#
# @!attribute [rw] mapping
#   @return [String, nil]
#
# @!attribute [rw] orange_routed_no
#   @return [Integer, nil]
#
# @!attribute [rw] parsed_declarations_no
#   @return [Integer, nil]
#
# @!attribute [rw] parser
#   @return [String, nil]
#
# @!attribute [rw] parsing_completed_at
#   @return [String, nil]
#
# @!attribute [rw] parsing_started_at
#   @return [String, nil]
#
# @!attribute [rw] passive_transport_nationality
#   @return [String, nil]
#
# @!attribute [rw] passive_transport_number
#   @return [String, nil]
#
# @!attribute [rw] processed_declarations_no
#   @return [Integer, nil]
#
# @!attribute [rw] processing_ended_at
#   @return [String, nil]
#
# @!attribute [rw] processing_started_at
#   @return [String, nil]
#
# @!attribute [rw] receipt_generating_started_at
#   @return [String, nil]
#
# @!attribute [rw] receipt_request_started_at
#   @return [String, nil]
#
# @!attribute [rw] receipt_request_status
#   @return [String, nil]
#
# @!attribute [rw] receipt_request_user
#   @return [String, nil]
#
# @!attribute [rw] receipts_zip
#   @return [String, nil]
#
# @!attribute [rw] red_routed_no
#   @return [Integer, nil]
#
# @!attribute [rw] rejected_status_no
#   @return [Integer, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] template
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
#
# @!attribute [rw] yellow_routed_no
#   @return [Integer, nil]
BulkUploadUpdateData = Struct.new(
  :id,
  :active_transport_nationality,
  :active_transport_number,
  :arrival_datetime,
  :client,
  :company_member,
  :created_at,
  :declarant,
  :declarations_no,
  :deleted_at,
  :departure_datetime,
  :errors_file,
  :external_id,
  :failed_declarations_no,
  :file,
  :green_routed_no,
  :h1_fallback_template,
  :house_transport_doc_ref,
  :issue_date,
  :mapping,
  :orange_routed_no,
  :parsed_declarations_no,
  :parser,
  :parsing_completed_at,
  :parsing_started_at,
  :passive_transport_nationality,
  :passive_transport_number,
  :processed_declarations_no,
  :processing_ended_at,
  :processing_started_at,
  :receipt_generating_started_at,
  :receipt_request_started_at,
  :receipt_request_status,
  :receipt_request_user,
  :receipts_zip,
  :red_routed_no,
  :rejected_status_no,
  :status,
  :template,
  :updated_at,
  :yellow_routed_no,
  keyword_init: true
)

# Request payload for BulkUpload#remove.
#
# @!attribute [rw] id
#   @return [String]
BulkUploadRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# File entity data model.
#
# @!attribute [rw] company
#   @return [String]
#
# @!attribute [rw] created_at
#   @return [String]
#
# @!attribute [rw] extension
#   @return [String]
#
# @!attribute [rw] file
#   @return [String]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] public
#   @return [Boolean]
#
# @!attribute [rw] size
#   @return [Integer, nil]
#
# @!attribute [rw] updated_at
#   @return [String]
#
# @!attribute [rw] url
#   @return [String]
FileType = Struct.new(
  :company,
  :created_at,
  :extension,
  :file,
  :id,
  :name,
  :public,
  :size,
  :updated_at,
  :url,
  keyword_init: true
)

# Request payload for File#create.
#
# @!attribute [rw] company
#   @return [String]
#
# @!attribute [rw] created_at
#   @return [String]
#
# @!attribute [rw] extension
#   @return [String]
#
# @!attribute [rw] file
#   @return [String]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] public
#   @return [Boolean]
#
# @!attribute [rw] size
#   @return [Integer, nil]
#
# @!attribute [rw] updated_at
#   @return [String]
#
# @!attribute [rw] url
#   @return [String]
FileCreateData = Struct.new(
  :company,
  :created_at,
  :extension,
  :file,
  :id,
  :name,
  :public,
  :size,
  :updated_at,
  :url,
  keyword_init: true
)

# PaginatedBulkUploadListList entity data model.
class PaginatedBulkUploadListList
end

# PaginatedPartyListList entity data model.
class PaginatedPartyListList
end

# PaginatedSubmissionListList entity data model.
class PaginatedSubmissionListList
end

# Party entity data model.
#
# @!attribute [rw] additional_declaration_type
#   @return [Hash]
#
# @!attribute [rw] address
#   @return [Hash]
#
# @!attribute [rw] authorisation
#   @return [Hash]
#
# @!attribute [rw] bank_details
#   @return [String, nil]
#
# @!attribute [rw] certificate
#   @return [Hash]
#
# @!attribute [rw] certificate_type
#   @return [String]
#
# @!attribute [rw] company
#   @return [String]
#
# @!attribute [rw] created_at
#   @return [String]
#
# @!attribute [rw] customs_office_of_lodgement
#   @return [Hash]
#
# @!attribute [rw] deleted_at
#   @return [String, nil]
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identification_number
#   @return [String, nil]
#
# @!attribute [rw] indirect_representative
#   @return [Boolean, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nhd_last_submission_year
#   @return [Integer, nil]
#
# @!attribute [rw] nhd_submission_counter
#   @return [Integer, nil]
#
# @!attribute [rw] person_paying_customs_duty
#   @return [String, nil]
#
# @!attribute [rw] phone_country_code
#   @return [String, nil]
#
# @!attribute [rw] phone_number
#   @return [String, nil]
#
# @!attribute [rw] preferred_payment_method
#   @return [Hash]
#
# @!attribute [rw] signed_form
#   @return [Hash]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] type_of_person
#   @return [Hash]
#
# @!attribute [rw] unlocode
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String]
Party = Struct.new(
  :additional_declaration_type,
  :address,
  :authorisation,
  :bank_details,
  :certificate,
  :certificate_type,
  :company,
  :created_at,
  :customs_office_of_lodgement,
  :deleted_at,
  :email,
  :id,
  :identification_number,
  :indirect_representative,
  :name,
  :nhd_last_submission_year,
  :nhd_submission_counter,
  :person_paying_customs_duty,
  :phone_country_code,
  :phone_number,
  :preferred_payment_method,
  :signed_form,
  :type,
  :type_of_person,
  :unlocode,
  :updated_at,
  keyword_init: true
)

# Request payload for Party#load.
#
# @!attribute [rw] id
#   @return [String]
PartyLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Party#list.
#
# @!attribute [rw] additional_declaration_type
#   @return [Hash, nil]
#
# @!attribute [rw] address
#   @return [Hash, nil]
#
# @!attribute [rw] authorisation
#   @return [Hash, nil]
#
# @!attribute [rw] bank_details
#   @return [String, nil]
#
# @!attribute [rw] certificate
#   @return [Hash, nil]
#
# @!attribute [rw] certificate_type
#   @return [String, nil]
#
# @!attribute [rw] company
#   @return [String, nil]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] customs_office_of_lodgement
#   @return [Hash, nil]
#
# @!attribute [rw] deleted_at
#   @return [String, nil]
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identification_number
#   @return [String, nil]
#
# @!attribute [rw] indirect_representative
#   @return [Boolean, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nhd_last_submission_year
#   @return [Integer, nil]
#
# @!attribute [rw] nhd_submission_counter
#   @return [Integer, nil]
#
# @!attribute [rw] person_paying_customs_duty
#   @return [String, nil]
#
# @!attribute [rw] phone_country_code
#   @return [String, nil]
#
# @!attribute [rw] phone_number
#   @return [String, nil]
#
# @!attribute [rw] preferred_payment_method
#   @return [Hash, nil]
#
# @!attribute [rw] signed_form
#   @return [Hash, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] type_of_person
#   @return [Hash, nil]
#
# @!attribute [rw] unlocode
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
PartyListMatch = Struct.new(
  :additional_declaration_type,
  :address,
  :authorisation,
  :bank_details,
  :certificate,
  :certificate_type,
  :company,
  :created_at,
  :customs_office_of_lodgement,
  :deleted_at,
  :email,
  :id,
  :identification_number,
  :indirect_representative,
  :name,
  :nhd_last_submission_year,
  :nhd_submission_counter,
  :person_paying_customs_duty,
  :phone_country_code,
  :phone_number,
  :preferred_payment_method,
  :signed_form,
  :type,
  :type_of_person,
  :unlocode,
  :updated_at,
  keyword_init: true
)

# Request payload for Party#create.
#
# @!attribute [rw] additional_declaration_type
#   @return [Hash]
#
# @!attribute [rw] address
#   @return [Hash]
#
# @!attribute [rw] authorisation
#   @return [Hash]
#
# @!attribute [rw] bank_details
#   @return [String, nil]
#
# @!attribute [rw] certificate
#   @return [Hash]
#
# @!attribute [rw] certificate_type
#   @return [String]
#
# @!attribute [rw] company
#   @return [String]
#
# @!attribute [rw] created_at
#   @return [String]
#
# @!attribute [rw] customs_office_of_lodgement
#   @return [Hash]
#
# @!attribute [rw] deleted_at
#   @return [String, nil]
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identification_number
#   @return [String, nil]
#
# @!attribute [rw] indirect_representative
#   @return [Boolean, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nhd_last_submission_year
#   @return [Integer, nil]
#
# @!attribute [rw] nhd_submission_counter
#   @return [Integer, nil]
#
# @!attribute [rw] person_paying_customs_duty
#   @return [String, nil]
#
# @!attribute [rw] phone_country_code
#   @return [String, nil]
#
# @!attribute [rw] phone_number
#   @return [String, nil]
#
# @!attribute [rw] preferred_payment_method
#   @return [Hash]
#
# @!attribute [rw] signed_form
#   @return [Hash]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] type_of_person
#   @return [Hash]
#
# @!attribute [rw] unlocode
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String]
PartyCreateData = Struct.new(
  :additional_declaration_type,
  :address,
  :authorisation,
  :bank_details,
  :certificate,
  :certificate_type,
  :company,
  :created_at,
  :customs_office_of_lodgement,
  :deleted_at,
  :email,
  :id,
  :identification_number,
  :indirect_representative,
  :name,
  :nhd_last_submission_year,
  :nhd_submission_counter,
  :person_paying_customs_duty,
  :phone_country_code,
  :phone_number,
  :preferred_payment_method,
  :signed_form,
  :type,
  :type_of_person,
  :unlocode,
  :updated_at,
  keyword_init: true
)

# Request payload for Party#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] additional_declaration_type
#   @return [Hash, nil]
#
# @!attribute [rw] address
#   @return [Hash, nil]
#
# @!attribute [rw] authorisation
#   @return [Hash, nil]
#
# @!attribute [rw] bank_details
#   @return [String, nil]
#
# @!attribute [rw] certificate
#   @return [Hash, nil]
#
# @!attribute [rw] certificate_type
#   @return [String, nil]
#
# @!attribute [rw] company
#   @return [String, nil]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] customs_office_of_lodgement
#   @return [Hash, nil]
#
# @!attribute [rw] deleted_at
#   @return [String, nil]
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] identification_number
#   @return [String, nil]
#
# @!attribute [rw] indirect_representative
#   @return [Boolean, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nhd_last_submission_year
#   @return [Integer, nil]
#
# @!attribute [rw] nhd_submission_counter
#   @return [Integer, nil]
#
# @!attribute [rw] person_paying_customs_duty
#   @return [String, nil]
#
# @!attribute [rw] phone_country_code
#   @return [String, nil]
#
# @!attribute [rw] phone_number
#   @return [String, nil]
#
# @!attribute [rw] preferred_payment_method
#   @return [Hash, nil]
#
# @!attribute [rw] signed_form
#   @return [Hash, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] type_of_person
#   @return [Hash, nil]
#
# @!attribute [rw] unlocode
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
PartyUpdateData = Struct.new(
  :id,
  :additional_declaration_type,
  :address,
  :authorisation,
  :bank_details,
  :certificate,
  :certificate_type,
  :company,
  :created_at,
  :customs_office_of_lodgement,
  :deleted_at,
  :email,
  :identification_number,
  :indirect_representative,
  :name,
  :nhd_last_submission_year,
  :nhd_submission_counter,
  :person_paying_customs_duty,
  :phone_country_code,
  :phone_number,
  :preferred_payment_method,
  :signed_form,
  :type,
  :type_of_person,
  :unlocode,
  :updated_at,
  keyword_init: true
)

# Request payload for Party#remove.
#
# @!attribute [rw] id
#   @return [String]
PartyRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# Submission entity data model.
#
# @!attribute [rw] additional_external_ids
#   @return [Array, nil]
#
# @!attribute [rw] amendment_reason
#   @return [String, nil]
#
# @!attribute [rw] amendment_status
#   @return [String, nil]
#
# @!attribute [rw] answers
#   @return [Array]
#
# @!attribute [rw] bypass_restricted_code
#   @return [Boolean, nil]
#
# @!attribute [rw] clearance_slip
#   @return [Hash]
#
# @!attribute [rw] client
#   @return [Hash]
#
# @!attribute [rw] company_member
#   @return [Hash]
#
# @!attribute [rw] consignee
#   @return [Hash]
#
# @!attribute [rw] consignor
#   @return [Hash]
#
# @!attribute [rw] created_at
#   @return [String]
#
# @!attribute [rw] declarant
#   @return [Hash]
#
# @!attribute [rw] document_upload_status
#   @return [String, nil]
#
# @!attribute [rw] documents_presentation_requested
#   @return [Boolean, nil]
#
# @!attribute [rw] documents_upload_requested
#   @return [Boolean, nil]
#
# @!attribute [rw] external_id
#   @return [String, nil]
#
# @!attribute [rw] form
#   @return [String]
#
# @!attribute [rw] goods_presentation_status
#   @return [String, nil]
#
# @!attribute [rw] hrcm_status
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] invalidation_status
#   @return [String, nil]
#
# @!attribute [rw] is_global_template
#   @return [Boolean, nil]
#
# @!attribute [rw] latest_notification_item
#   @return [String, nil]
#
# @!attribute [rw] latest_state
#   @return [Hash]
#
# @!attribute [rw] lrn
#   @return [String, nil]
#
# @!attribute [rw] mrn
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] partial_answers
#   @return [Boolean, nil]
#
# @!attribute [rw] receipt
#   @return [Hash]
#
# @!attribute [rw] refund_application_status
#   @return [String, nil]
#
# @!attribute [rw] route
#   @return [String, nil]
#
# @!attribute [rw] shipment_items_no
#   @return [Integer, nil]
#
# @!attribute [rw] shipment_items_quantity_no
#   @return [Integer, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
#
# @!attribute [rw] source_type
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] template
#   @return [Boolean, nil]
#
# @!attribute [rw] template_id
#   @return [String, nil]
#
# @!attribute [rw] template_properties
#   @return [Array, nil]
#
# @!attribute [rw] total_tax_amount
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String]
#
# @!attribute [rw] verification_errors
#   @return [Array, nil]
#
# @!attribute [rw] verification_status
#   @return [String, nil]
Submission = Struct.new(
  :additional_external_ids,
  :amendment_reason,
  :amendment_status,
  :answers,
  :bypass_restricted_code,
  :clearance_slip,
  :client,
  :company_member,
  :consignee,
  :consignor,
  :created_at,
  :declarant,
  :document_upload_status,
  :documents_presentation_requested,
  :documents_upload_requested,
  :external_id,
  :form,
  :goods_presentation_status,
  :hrcm_status,
  :id,
  :invalidation_status,
  :is_global_template,
  :latest_notification_item,
  :latest_state,
  :lrn,
  :mrn,
  :name,
  :partial_answers,
  :receipt,
  :refund_application_status,
  :route,
  :shipment_items_no,
  :shipment_items_quantity_no,
  :source,
  :source_type,
  :status,
  :template,
  :template_id,
  :template_properties,
  :total_tax_amount,
  :updated_at,
  :verification_errors,
  :verification_status,
  keyword_init: true
)

# Request payload for Submission#load.
#
# @!attribute [rw] id
#   @return [String]
SubmissionLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Submission#list.
#
# @!attribute [rw] additional_external_ids
#   @return [Array, nil]
#
# @!attribute [rw] amendment_reason
#   @return [String, nil]
#
# @!attribute [rw] amendment_status
#   @return [String, nil]
#
# @!attribute [rw] answers
#   @return [Array, nil]
#
# @!attribute [rw] bypass_restricted_code
#   @return [Boolean, nil]
#
# @!attribute [rw] clearance_slip
#   @return [Hash, nil]
#
# @!attribute [rw] client
#   @return [Hash, nil]
#
# @!attribute [rw] company_member
#   @return [Hash, nil]
#
# @!attribute [rw] consignee
#   @return [Hash, nil]
#
# @!attribute [rw] consignor
#   @return [Hash, nil]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] declarant
#   @return [Hash, nil]
#
# @!attribute [rw] document_upload_status
#   @return [String, nil]
#
# @!attribute [rw] documents_presentation_requested
#   @return [Boolean, nil]
#
# @!attribute [rw] documents_upload_requested
#   @return [Boolean, nil]
#
# @!attribute [rw] external_id
#   @return [String, nil]
#
# @!attribute [rw] form
#   @return [String, nil]
#
# @!attribute [rw] goods_presentation_status
#   @return [String, nil]
#
# @!attribute [rw] hrcm_status
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] invalidation_status
#   @return [String, nil]
#
# @!attribute [rw] is_global_template
#   @return [Boolean, nil]
#
# @!attribute [rw] latest_notification_item
#   @return [String, nil]
#
# @!attribute [rw] latest_state
#   @return [Hash, nil]
#
# @!attribute [rw] lrn
#   @return [String, nil]
#
# @!attribute [rw] mrn
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] partial_answers
#   @return [Boolean, nil]
#
# @!attribute [rw] receipt
#   @return [Hash, nil]
#
# @!attribute [rw] refund_application_status
#   @return [String, nil]
#
# @!attribute [rw] route
#   @return [String, nil]
#
# @!attribute [rw] shipment_items_no
#   @return [Integer, nil]
#
# @!attribute [rw] shipment_items_quantity_no
#   @return [Integer, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
#
# @!attribute [rw] source_type
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] template
#   @return [Boolean, nil]
#
# @!attribute [rw] template_id
#   @return [String, nil]
#
# @!attribute [rw] template_properties
#   @return [Array, nil]
#
# @!attribute [rw] total_tax_amount
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
#
# @!attribute [rw] verification_errors
#   @return [Array, nil]
#
# @!attribute [rw] verification_status
#   @return [String, nil]
SubmissionListMatch = Struct.new(
  :additional_external_ids,
  :amendment_reason,
  :amendment_status,
  :answers,
  :bypass_restricted_code,
  :clearance_slip,
  :client,
  :company_member,
  :consignee,
  :consignor,
  :created_at,
  :declarant,
  :document_upload_status,
  :documents_presentation_requested,
  :documents_upload_requested,
  :external_id,
  :form,
  :goods_presentation_status,
  :hrcm_status,
  :id,
  :invalidation_status,
  :is_global_template,
  :latest_notification_item,
  :latest_state,
  :lrn,
  :mrn,
  :name,
  :partial_answers,
  :receipt,
  :refund_application_status,
  :route,
  :shipment_items_no,
  :shipment_items_quantity_no,
  :source,
  :source_type,
  :status,
  :template,
  :template_id,
  :template_properties,
  :total_tax_amount,
  :updated_at,
  :verification_errors,
  :verification_status,
  keyword_init: true
)

# Request payload for Submission#create.
#
# @!attribute [rw] additional_external_ids
#   @return [Array, nil]
#
# @!attribute [rw] amendment_reason
#   @return [String, nil]
#
# @!attribute [rw] amendment_status
#   @return [String, nil]
#
# @!attribute [rw] answers
#   @return [Array]
#
# @!attribute [rw] bypass_restricted_code
#   @return [Boolean, nil]
#
# @!attribute [rw] clearance_slip
#   @return [Hash]
#
# @!attribute [rw] client
#   @return [Hash]
#
# @!attribute [rw] company_member
#   @return [Hash]
#
# @!attribute [rw] consignee
#   @return [Hash]
#
# @!attribute [rw] consignor
#   @return [Hash]
#
# @!attribute [rw] created_at
#   @return [String]
#
# @!attribute [rw] declarant
#   @return [Hash]
#
# @!attribute [rw] document_upload_status
#   @return [String, nil]
#
# @!attribute [rw] documents_presentation_requested
#   @return [Boolean, nil]
#
# @!attribute [rw] documents_upload_requested
#   @return [Boolean, nil]
#
# @!attribute [rw] external_id
#   @return [String, nil]
#
# @!attribute [rw] form
#   @return [String]
#
# @!attribute [rw] goods_presentation_status
#   @return [String, nil]
#
# @!attribute [rw] hrcm_status
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] invalidation_status
#   @return [String, nil]
#
# @!attribute [rw] is_global_template
#   @return [Boolean, nil]
#
# @!attribute [rw] latest_notification_item
#   @return [String, nil]
#
# @!attribute [rw] latest_state
#   @return [Hash]
#
# @!attribute [rw] lrn
#   @return [String, nil]
#
# @!attribute [rw] mrn
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] partial_answers
#   @return [Boolean, nil]
#
# @!attribute [rw] receipt
#   @return [Hash]
#
# @!attribute [rw] refund_application_status
#   @return [String, nil]
#
# @!attribute [rw] route
#   @return [String, nil]
#
# @!attribute [rw] shipment_items_no
#   @return [Integer, nil]
#
# @!attribute [rw] shipment_items_quantity_no
#   @return [Integer, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
#
# @!attribute [rw] source_type
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] template
#   @return [Boolean, nil]
#
# @!attribute [rw] template_id
#   @return [String, nil]
#
# @!attribute [rw] template_properties
#   @return [Array, nil]
#
# @!attribute [rw] total_tax_amount
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String]
#
# @!attribute [rw] verification_errors
#   @return [Array, nil]
#
# @!attribute [rw] verification_status
#   @return [String, nil]
SubmissionCreateData = Struct.new(
  :additional_external_ids,
  :amendment_reason,
  :amendment_status,
  :answers,
  :bypass_restricted_code,
  :clearance_slip,
  :client,
  :company_member,
  :consignee,
  :consignor,
  :created_at,
  :declarant,
  :document_upload_status,
  :documents_presentation_requested,
  :documents_upload_requested,
  :external_id,
  :form,
  :goods_presentation_status,
  :hrcm_status,
  :id,
  :invalidation_status,
  :is_global_template,
  :latest_notification_item,
  :latest_state,
  :lrn,
  :mrn,
  :name,
  :partial_answers,
  :receipt,
  :refund_application_status,
  :route,
  :shipment_items_no,
  :shipment_items_quantity_no,
  :source,
  :source_type,
  :status,
  :template,
  :template_id,
  :template_properties,
  :total_tax_amount,
  :updated_at,
  :verification_errors,
  :verification_status,
  keyword_init: true
)

# Request payload for Submission#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] additional_external_ids
#   @return [Array, nil]
#
# @!attribute [rw] amendment_reason
#   @return [String, nil]
#
# @!attribute [rw] amendment_status
#   @return [String, nil]
#
# @!attribute [rw] answers
#   @return [Array, nil]
#
# @!attribute [rw] bypass_restricted_code
#   @return [Boolean, nil]
#
# @!attribute [rw] clearance_slip
#   @return [Hash, nil]
#
# @!attribute [rw] client
#   @return [Hash, nil]
#
# @!attribute [rw] company_member
#   @return [Hash, nil]
#
# @!attribute [rw] consignee
#   @return [Hash, nil]
#
# @!attribute [rw] consignor
#   @return [Hash, nil]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] declarant
#   @return [Hash, nil]
#
# @!attribute [rw] document_upload_status
#   @return [String, nil]
#
# @!attribute [rw] documents_presentation_requested
#   @return [Boolean, nil]
#
# @!attribute [rw] documents_upload_requested
#   @return [Boolean, nil]
#
# @!attribute [rw] external_id
#   @return [String, nil]
#
# @!attribute [rw] form
#   @return [String, nil]
#
# @!attribute [rw] goods_presentation_status
#   @return [String, nil]
#
# @!attribute [rw] hrcm_status
#   @return [String, nil]
#
# @!attribute [rw] invalidation_status
#   @return [String, nil]
#
# @!attribute [rw] is_global_template
#   @return [Boolean, nil]
#
# @!attribute [rw] latest_notification_item
#   @return [String, nil]
#
# @!attribute [rw] latest_state
#   @return [Hash, nil]
#
# @!attribute [rw] lrn
#   @return [String, nil]
#
# @!attribute [rw] mrn
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] partial_answers
#   @return [Boolean, nil]
#
# @!attribute [rw] receipt
#   @return [Hash, nil]
#
# @!attribute [rw] refund_application_status
#   @return [String, nil]
#
# @!attribute [rw] route
#   @return [String, nil]
#
# @!attribute [rw] shipment_items_no
#   @return [Integer, nil]
#
# @!attribute [rw] shipment_items_quantity_no
#   @return [Integer, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
#
# @!attribute [rw] source_type
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] template
#   @return [Boolean, nil]
#
# @!attribute [rw] template_id
#   @return [String, nil]
#
# @!attribute [rw] template_properties
#   @return [Array, nil]
#
# @!attribute [rw] total_tax_amount
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
#
# @!attribute [rw] verification_errors
#   @return [Array, nil]
#
# @!attribute [rw] verification_status
#   @return [String, nil]
SubmissionUpdateData = Struct.new(
  :id,
  :additional_external_ids,
  :amendment_reason,
  :amendment_status,
  :answers,
  :bypass_restricted_code,
  :clearance_slip,
  :client,
  :company_member,
  :consignee,
  :consignor,
  :created_at,
  :declarant,
  :document_upload_status,
  :documents_presentation_requested,
  :documents_upload_requested,
  :external_id,
  :form,
  :goods_presentation_status,
  :hrcm_status,
  :invalidation_status,
  :is_global_template,
  :latest_notification_item,
  :latest_state,
  :lrn,
  :mrn,
  :name,
  :partial_answers,
  :receipt,
  :refund_application_status,
  :route,
  :shipment_items_no,
  :shipment_items_quantity_no,
  :source,
  :source_type,
  :status,
  :template,
  :template_id,
  :template_properties,
  :total_tax_amount,
  :updated_at,
  :verification_errors,
  :verification_status,
  keyword_init: true
)

# Request payload for Submission#remove.
#
# @!attribute [rw] id
#   @return [String]
SubmissionRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# SubmissionDetail entity data model.
#
# @!attribute [rw] additional_external_ids
#   @return [Array, nil]
#
# @!attribute [rw] additional_information
#   @return [Array]
#
# @!attribute [rw] amendment_status
#   @return [String, nil]
#
# @!attribute [rw] clearance_slip
#   @return [Hash]
#
# @!attribute [rw] client
#   @return [Hash]
#
# @!attribute [rw] company
#   @return [String]
#
# @!attribute [rw] company_member
#   @return [Hash]
#
# @!attribute [rw] consignee
#   @return [Hash]
#
# @!attribute [rw] consignor
#   @return [Hash]
#
# @!attribute [rw] created_at
#   @return [String]
#
# @!attribute [rw] declarant
#   @return [Hash]
#
# @!attribute [rw] document_upload_status
#   @return [String, nil]
#
# @!attribute [rw] documents_presentation_requested
#   @return [Boolean, nil]
#
# @!attribute [rw] documents_upload_requested
#   @return [Boolean, nil]
#
# @!attribute [rw] external_id
#   @return [String, nil]
#
# @!attribute [rw] form
#   @return [String]
#
# @!attribute [rw] goods_presentation_status
#   @return [String, nil]
#
# @!attribute [rw] hrcm_status
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] invalidation_status
#   @return [String, nil]
#
# @!attribute [rw] is_global_template
#   @return [Boolean, nil]
#
# @!attribute [rw] latest_notification_item
#   @return [String, nil]
#
# @!attribute [rw] latest_state
#   @return [Hash]
#
# @!attribute [rw] lrn
#   @return [String, nil]
#
# @!attribute [rw] mrn
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] receipt
#   @return [Hash]
#
# @!attribute [rw] refund_application_status
#   @return [String, nil]
#
# @!attribute [rw] route
#   @return [String, nil]
#
# @!attribute [rw] shipment_items_no
#   @return [Integer, nil]
#
# @!attribute [rw] shipment_items_quantity_no
#   @return [Integer, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
#
# @!attribute [rw] source_type
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] submission
#   @return [String]
#
# @!attribute [rw] supporting_documents
#   @return [Array, nil]
#
# @!attribute [rw] template
#   @return [Boolean, nil]
#
# @!attribute [rw] total_tax_amount
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String]
#
# @!attribute [rw] verification_errors
#   @return [Array, nil]
#
# @!attribute [rw] verification_status
#   @return [String, nil]
SubmissionDetail = Struct.new(
  :additional_external_ids,
  :additional_information,
  :amendment_status,
  :clearance_slip,
  :client,
  :company,
  :company_member,
  :consignee,
  :consignor,
  :created_at,
  :declarant,
  :document_upload_status,
  :documents_presentation_requested,
  :documents_upload_requested,
  :external_id,
  :form,
  :goods_presentation_status,
  :hrcm_status,
  :id,
  :invalidation_status,
  :is_global_template,
  :latest_notification_item,
  :latest_state,
  :lrn,
  :mrn,
  :name,
  :receipt,
  :refund_application_status,
  :route,
  :shipment_items_no,
  :shipment_items_quantity_no,
  :source,
  :source_type,
  :status,
  :submission,
  :supporting_documents,
  :template,
  :total_tax_amount,
  :updated_at,
  :verification_errors,
  :verification_status,
  keyword_init: true
)

# Request payload for SubmissionDetail#create.
#
# @!attribute [rw] additional_external_ids
#   @return [Array, nil]
#
# @!attribute [rw] additional_information
#   @return [Array]
#
# @!attribute [rw] amendment_status
#   @return [String, nil]
#
# @!attribute [rw] clearance_slip
#   @return [Hash]
#
# @!attribute [rw] client
#   @return [Hash]
#
# @!attribute [rw] company
#   @return [String]
#
# @!attribute [rw] company_member
#   @return [Hash]
#
# @!attribute [rw] consignee
#   @return [Hash]
#
# @!attribute [rw] consignor
#   @return [Hash]
#
# @!attribute [rw] created_at
#   @return [String]
#
# @!attribute [rw] declarant
#   @return [Hash]
#
# @!attribute [rw] document_upload_status
#   @return [String, nil]
#
# @!attribute [rw] documents_presentation_requested
#   @return [Boolean, nil]
#
# @!attribute [rw] documents_upload_requested
#   @return [Boolean, nil]
#
# @!attribute [rw] external_id
#   @return [String, nil]
#
# @!attribute [rw] form
#   @return [String]
#
# @!attribute [rw] goods_presentation_status
#   @return [String, nil]
#
# @!attribute [rw] hrcm_status
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] invalidation_status
#   @return [String, nil]
#
# @!attribute [rw] is_global_template
#   @return [Boolean, nil]
#
# @!attribute [rw] latest_notification_item
#   @return [String, nil]
#
# @!attribute [rw] latest_state
#   @return [Hash]
#
# @!attribute [rw] lrn
#   @return [String, nil]
#
# @!attribute [rw] mrn
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] receipt
#   @return [Hash]
#
# @!attribute [rw] refund_application_status
#   @return [String, nil]
#
# @!attribute [rw] route
#   @return [String, nil]
#
# @!attribute [rw] shipment_items_no
#   @return [Integer, nil]
#
# @!attribute [rw] shipment_items_quantity_no
#   @return [Integer, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
#
# @!attribute [rw] source_type
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] submission
#   @return [String]
#
# @!attribute [rw] supporting_documents
#   @return [Array, nil]
#
# @!attribute [rw] template
#   @return [Boolean, nil]
#
# @!attribute [rw] total_tax_amount
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String]
#
# @!attribute [rw] verification_errors
#   @return [Array, nil]
#
# @!attribute [rw] verification_status
#   @return [String, nil]
SubmissionDetailCreateData = Struct.new(
  :additional_external_ids,
  :additional_information,
  :amendment_status,
  :clearance_slip,
  :client,
  :company,
  :company_member,
  :consignee,
  :consignor,
  :created_at,
  :declarant,
  :document_upload_status,
  :documents_presentation_requested,
  :documents_upload_requested,
  :external_id,
  :form,
  :goods_presentation_status,
  :hrcm_status,
  :id,
  :invalidation_status,
  :is_global_template,
  :latest_notification_item,
  :latest_state,
  :lrn,
  :mrn,
  :name,
  :receipt,
  :refund_application_status,
  :route,
  :shipment_items_no,
  :shipment_items_quantity_no,
  :source,
  :source_type,
  :status,
  :submission,
  :supporting_documents,
  :template,
  :total_tax_amount,
  :updated_at,
  :verification_errors,
  :verification_status,
  keyword_init: true
)

