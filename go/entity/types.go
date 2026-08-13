// Typed models for the CustomsWindow SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/customs-window-sdk/go/core"
)

// BulkUpload is the typed data model for the bulk_upload entity.
type BulkUpload struct {
	ActiveTransportNationality *string `json:"active_transport_nationality,omitempty"`
	ActiveTransportNumber *string `json:"active_transport_number,omitempty"`
	ArrivalDatetime *string `json:"arrival_datetime,omitempty"`
	Client string `json:"client"`
	CompanyMember string `json:"company_member"`
	CreatedAt string `json:"created_at"`
	Declarant string `json:"declarant"`
	DeclarationsNo int `json:"declarations_no"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	DepartureDatetime *string `json:"departure_datetime,omitempty"`
	ErrorsFile map[string]any `json:"errors_file"`
	ExternalId *string `json:"external_id,omitempty"`
	FailedDeclarationsNo *int `json:"failed_declarations_no,omitempty"`
	File map[string]any `json:"file"`
	GreenRoutedNo *int `json:"green_routed_no,omitempty"`
	H1FallbackTemplate *string `json:"h1_fallback_template,omitempty"`
	HouseTransportDocRef *string `json:"house_transport_doc_ref,omitempty"`
	Id *string `json:"id,omitempty"`
	IssueDate *string `json:"issue_date,omitempty"`
	Mapping *string `json:"mapping,omitempty"`
	OrangeRoutedNo *int `json:"orange_routed_no,omitempty"`
	ParsedDeclarationsNo *int `json:"parsed_declarations_no,omitempty"`
	Parser *string `json:"parser,omitempty"`
	ParsingCompletedAt *string `json:"parsing_completed_at,omitempty"`
	ParsingStartedAt *string `json:"parsing_started_at,omitempty"`
	PassiveTransportNationality *string `json:"passive_transport_nationality,omitempty"`
	PassiveTransportNumber *string `json:"passive_transport_number,omitempty"`
	ProcessedDeclarationsNo *int `json:"processed_declarations_no,omitempty"`
	ProcessingEndedAt *string `json:"processing_ended_at,omitempty"`
	ProcessingStartedAt *string `json:"processing_started_at,omitempty"`
	ReceiptGeneratingStartedAt *string `json:"receipt_generating_started_at,omitempty"`
	ReceiptRequestStartedAt *string `json:"receipt_request_started_at,omitempty"`
	ReceiptRequestStatus *string `json:"receipt_request_status,omitempty"`
	ReceiptRequestUser *string `json:"receipt_request_user,omitempty"`
	ReceiptsZip *string `json:"receipts_zip,omitempty"`
	RedRoutedNo *int `json:"red_routed_no,omitempty"`
	RejectedStatusNo *int `json:"rejected_status_no,omitempty"`
	Status *string `json:"status,omitempty"`
	Template *string `json:"template,omitempty"`
	UpdatedAt string `json:"updated_at"`
	YellowRoutedNo *int `json:"yellow_routed_no,omitempty"`
}

// BulkUploadLoadMatch is the typed request payload for BulkUpload.LoadTyped.
type BulkUploadLoadMatch struct {
	Id string `json:"id"`
}

// BulkUploadListMatch is the typed request payload for BulkUpload.ListTyped.
type BulkUploadListMatch struct {
	ActiveTransportNationality *string `json:"active_transport_nationality,omitempty"`
	ActiveTransportNumber *string `json:"active_transport_number,omitempty"`
	ArrivalDatetime *string `json:"arrival_datetime,omitempty"`
	Client *string `json:"client,omitempty"`
	CompanyMember *string `json:"company_member,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Declarant *string `json:"declarant,omitempty"`
	DeclarationsNo *int `json:"declarations_no,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	DepartureDatetime *string `json:"departure_datetime,omitempty"`
	ErrorsFile *map[string]any `json:"errors_file,omitempty"`
	ExternalId *string `json:"external_id,omitempty"`
	FailedDeclarationsNo *int `json:"failed_declarations_no,omitempty"`
	File *map[string]any `json:"file,omitempty"`
	GreenRoutedNo *int `json:"green_routed_no,omitempty"`
	H1FallbackTemplate *string `json:"h1_fallback_template,omitempty"`
	HouseTransportDocRef *string `json:"house_transport_doc_ref,omitempty"`
	Id *string `json:"id,omitempty"`
	IssueDate *string `json:"issue_date,omitempty"`
	Mapping *string `json:"mapping,omitempty"`
	OrangeRoutedNo *int `json:"orange_routed_no,omitempty"`
	ParsedDeclarationsNo *int `json:"parsed_declarations_no,omitempty"`
	Parser *string `json:"parser,omitempty"`
	ParsingCompletedAt *string `json:"parsing_completed_at,omitempty"`
	ParsingStartedAt *string `json:"parsing_started_at,omitempty"`
	PassiveTransportNationality *string `json:"passive_transport_nationality,omitempty"`
	PassiveTransportNumber *string `json:"passive_transport_number,omitempty"`
	ProcessedDeclarationsNo *int `json:"processed_declarations_no,omitempty"`
	ProcessingEndedAt *string `json:"processing_ended_at,omitempty"`
	ProcessingStartedAt *string `json:"processing_started_at,omitempty"`
	ReceiptGeneratingStartedAt *string `json:"receipt_generating_started_at,omitempty"`
	ReceiptRequestStartedAt *string `json:"receipt_request_started_at,omitempty"`
	ReceiptRequestStatus *string `json:"receipt_request_status,omitempty"`
	ReceiptRequestUser *string `json:"receipt_request_user,omitempty"`
	ReceiptsZip *string `json:"receipts_zip,omitempty"`
	RedRoutedNo *int `json:"red_routed_no,omitempty"`
	RejectedStatusNo *int `json:"rejected_status_no,omitempty"`
	Status *string `json:"status,omitempty"`
	Template *string `json:"template,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	YellowRoutedNo *int `json:"yellow_routed_no,omitempty"`
}

// BulkUploadCreateData is the typed request payload for BulkUpload.CreateTyped.
type BulkUploadCreateData struct {
	ActiveTransportNationality *string `json:"active_transport_nationality,omitempty"`
	ActiveTransportNumber *string `json:"active_transport_number,omitempty"`
	ArrivalDatetime *string `json:"arrival_datetime,omitempty"`
	Client string `json:"client"`
	CompanyMember string `json:"company_member"`
	CreatedAt string `json:"created_at"`
	Declarant string `json:"declarant"`
	DeclarationsNo int `json:"declarations_no"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	DepartureDatetime *string `json:"departure_datetime,omitempty"`
	ErrorsFile map[string]any `json:"errors_file"`
	ExternalId *string `json:"external_id,omitempty"`
	FailedDeclarationsNo *int `json:"failed_declarations_no,omitempty"`
	File map[string]any `json:"file"`
	GreenRoutedNo *int `json:"green_routed_no,omitempty"`
	H1FallbackTemplate *string `json:"h1_fallback_template,omitempty"`
	HouseTransportDocRef *string `json:"house_transport_doc_ref,omitempty"`
	Id *string `json:"id,omitempty"`
	IssueDate *string `json:"issue_date,omitempty"`
	Mapping *string `json:"mapping,omitempty"`
	OrangeRoutedNo *int `json:"orange_routed_no,omitempty"`
	ParsedDeclarationsNo *int `json:"parsed_declarations_no,omitempty"`
	Parser *string `json:"parser,omitempty"`
	ParsingCompletedAt *string `json:"parsing_completed_at,omitempty"`
	ParsingStartedAt *string `json:"parsing_started_at,omitempty"`
	PassiveTransportNationality *string `json:"passive_transport_nationality,omitempty"`
	PassiveTransportNumber *string `json:"passive_transport_number,omitempty"`
	ProcessedDeclarationsNo *int `json:"processed_declarations_no,omitempty"`
	ProcessingEndedAt *string `json:"processing_ended_at,omitempty"`
	ProcessingStartedAt *string `json:"processing_started_at,omitempty"`
	ReceiptGeneratingStartedAt *string `json:"receipt_generating_started_at,omitempty"`
	ReceiptRequestStartedAt *string `json:"receipt_request_started_at,omitempty"`
	ReceiptRequestStatus *string `json:"receipt_request_status,omitempty"`
	ReceiptRequestUser *string `json:"receipt_request_user,omitempty"`
	ReceiptsZip *string `json:"receipts_zip,omitempty"`
	RedRoutedNo *int `json:"red_routed_no,omitempty"`
	RejectedStatusNo *int `json:"rejected_status_no,omitempty"`
	Status *string `json:"status,omitempty"`
	Template *string `json:"template,omitempty"`
	UpdatedAt string `json:"updated_at"`
	YellowRoutedNo *int `json:"yellow_routed_no,omitempty"`
}

// BulkUploadUpdateData is the typed request payload for BulkUpload.UpdateTyped.
type BulkUploadUpdateData struct {
	Id string `json:"id"`
	ActiveTransportNationality *string `json:"active_transport_nationality,omitempty"`
	ActiveTransportNumber *string `json:"active_transport_number,omitempty"`
	ArrivalDatetime *string `json:"arrival_datetime,omitempty"`
	Client *string `json:"client,omitempty"`
	CompanyMember *string `json:"company_member,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Declarant *string `json:"declarant,omitempty"`
	DeclarationsNo *int `json:"declarations_no,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	DepartureDatetime *string `json:"departure_datetime,omitempty"`
	ErrorsFile *map[string]any `json:"errors_file,omitempty"`
	ExternalId *string `json:"external_id,omitempty"`
	FailedDeclarationsNo *int `json:"failed_declarations_no,omitempty"`
	File *map[string]any `json:"file,omitempty"`
	GreenRoutedNo *int `json:"green_routed_no,omitempty"`
	H1FallbackTemplate *string `json:"h1_fallback_template,omitempty"`
	HouseTransportDocRef *string `json:"house_transport_doc_ref,omitempty"`
	IssueDate *string `json:"issue_date,omitempty"`
	Mapping *string `json:"mapping,omitempty"`
	OrangeRoutedNo *int `json:"orange_routed_no,omitempty"`
	ParsedDeclarationsNo *int `json:"parsed_declarations_no,omitempty"`
	Parser *string `json:"parser,omitempty"`
	ParsingCompletedAt *string `json:"parsing_completed_at,omitempty"`
	ParsingStartedAt *string `json:"parsing_started_at,omitempty"`
	PassiveTransportNationality *string `json:"passive_transport_nationality,omitempty"`
	PassiveTransportNumber *string `json:"passive_transport_number,omitempty"`
	ProcessedDeclarationsNo *int `json:"processed_declarations_no,omitempty"`
	ProcessingEndedAt *string `json:"processing_ended_at,omitempty"`
	ProcessingStartedAt *string `json:"processing_started_at,omitempty"`
	ReceiptGeneratingStartedAt *string `json:"receipt_generating_started_at,omitempty"`
	ReceiptRequestStartedAt *string `json:"receipt_request_started_at,omitempty"`
	ReceiptRequestStatus *string `json:"receipt_request_status,omitempty"`
	ReceiptRequestUser *string `json:"receipt_request_user,omitempty"`
	ReceiptsZip *string `json:"receipts_zip,omitempty"`
	RedRoutedNo *int `json:"red_routed_no,omitempty"`
	RejectedStatusNo *int `json:"rejected_status_no,omitempty"`
	Status *string `json:"status,omitempty"`
	Template *string `json:"template,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	YellowRoutedNo *int `json:"yellow_routed_no,omitempty"`
}

// BulkUploadRemoveMatch is the typed request payload for BulkUpload.RemoveTyped.
type BulkUploadRemoveMatch struct {
	Id string `json:"id"`
}

// File is the typed data model for the file entity.
type File struct {
	Company string `json:"company"`
	CreatedAt string `json:"created_at"`
	Extension string `json:"extension"`
	File string `json:"file"`
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Public bool `json:"public"`
	Size *int `json:"size,omitempty"`
	UpdatedAt string `json:"updated_at"`
	Url string `json:"url"`
}

// FileCreateData is the typed request payload for File.CreateTyped.
type FileCreateData struct {
	Company string `json:"company"`
	CreatedAt string `json:"created_at"`
	Extension string `json:"extension"`
	File string `json:"file"`
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	Public bool `json:"public"`
	Size *int `json:"size,omitempty"`
	UpdatedAt string `json:"updated_at"`
	Url string `json:"url"`
}

// PaginatedBulkUploadListList is the typed data model for the paginated_bulk_upload_list_list entity.
type PaginatedBulkUploadListList struct {
}

// PaginatedPartyListList is the typed data model for the paginated_party_list_list entity.
type PaginatedPartyListList struct {
}

// PaginatedSubmissionListList is the typed data model for the paginated_submission_list_list entity.
type PaginatedSubmissionListList struct {
}

// Party is the typed data model for the party entity.
type Party struct {
	AdditionalDeclarationType map[string]any `json:"additional_declaration_type"`
	Address map[string]any `json:"address"`
	Authorisation map[string]any `json:"authorisation"`
	BankDetails *string `json:"bank_details,omitempty"`
	Certificate map[string]any `json:"certificate"`
	CertificateType string `json:"certificate_type"`
	Company string `json:"company"`
	CreatedAt string `json:"created_at"`
	CustomsOfficeOfLodgement map[string]any `json:"customs_office_of_lodgement"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Email *string `json:"email,omitempty"`
	Id *string `json:"id,omitempty"`
	IdentificationNumber *string `json:"identification_number,omitempty"`
	IndirectRepresentative *bool `json:"indirect_representative,omitempty"`
	Name *string `json:"name,omitempty"`
	NhdLastSubmissionYear *int `json:"nhd_last_submission_year,omitempty"`
	NhdSubmissionCounter *int `json:"nhd_submission_counter,omitempty"`
	PersonPayingCustomsDuty *string `json:"person_paying_customs_duty,omitempty"`
	PhoneCountryCode *string `json:"phone_country_code,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	PreferredPaymentMethod map[string]any `json:"preferred_payment_method"`
	SignedForm map[string]any `json:"signed_form"`
	Type *string `json:"type,omitempty"`
	TypeOfPerson map[string]any `json:"type_of_person"`
	Unlocode *string `json:"unlocode,omitempty"`
	UpdatedAt string `json:"updated_at"`
}

// PartyLoadMatch is the typed request payload for Party.LoadTyped.
type PartyLoadMatch struct {
	Id string `json:"id"`
}

// PartyListMatch is the typed request payload for Party.ListTyped.
type PartyListMatch struct {
	AdditionalDeclarationType *map[string]any `json:"additional_declaration_type,omitempty"`
	Address *map[string]any `json:"address,omitempty"`
	Authorisation *map[string]any `json:"authorisation,omitempty"`
	BankDetails *string `json:"bank_details,omitempty"`
	Certificate *map[string]any `json:"certificate,omitempty"`
	CertificateType *string `json:"certificate_type,omitempty"`
	Company *string `json:"company,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CustomsOfficeOfLodgement *map[string]any `json:"customs_office_of_lodgement,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Email *string `json:"email,omitempty"`
	Id *string `json:"id,omitempty"`
	IdentificationNumber *string `json:"identification_number,omitempty"`
	IndirectRepresentative *bool `json:"indirect_representative,omitempty"`
	Name *string `json:"name,omitempty"`
	NhdLastSubmissionYear *int `json:"nhd_last_submission_year,omitempty"`
	NhdSubmissionCounter *int `json:"nhd_submission_counter,omitempty"`
	PersonPayingCustomsDuty *string `json:"person_paying_customs_duty,omitempty"`
	PhoneCountryCode *string `json:"phone_country_code,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	PreferredPaymentMethod *map[string]any `json:"preferred_payment_method,omitempty"`
	SignedForm *map[string]any `json:"signed_form,omitempty"`
	Type *string `json:"type,omitempty"`
	TypeOfPerson *map[string]any `json:"type_of_person,omitempty"`
	Unlocode *string `json:"unlocode,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// PartyCreateData is the typed request payload for Party.CreateTyped.
type PartyCreateData struct {
	AdditionalDeclarationType map[string]any `json:"additional_declaration_type"`
	Address map[string]any `json:"address"`
	Authorisation map[string]any `json:"authorisation"`
	BankDetails *string `json:"bank_details,omitempty"`
	Certificate map[string]any `json:"certificate"`
	CertificateType string `json:"certificate_type"`
	Company string `json:"company"`
	CreatedAt string `json:"created_at"`
	CustomsOfficeOfLodgement map[string]any `json:"customs_office_of_lodgement"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Email *string `json:"email,omitempty"`
	Id *string `json:"id,omitempty"`
	IdentificationNumber *string `json:"identification_number,omitempty"`
	IndirectRepresentative *bool `json:"indirect_representative,omitempty"`
	Name *string `json:"name,omitempty"`
	NhdLastSubmissionYear *int `json:"nhd_last_submission_year,omitempty"`
	NhdSubmissionCounter *int `json:"nhd_submission_counter,omitempty"`
	PersonPayingCustomsDuty *string `json:"person_paying_customs_duty,omitempty"`
	PhoneCountryCode *string `json:"phone_country_code,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	PreferredPaymentMethod map[string]any `json:"preferred_payment_method"`
	SignedForm map[string]any `json:"signed_form"`
	Type *string `json:"type,omitempty"`
	TypeOfPerson map[string]any `json:"type_of_person"`
	Unlocode *string `json:"unlocode,omitempty"`
	UpdatedAt string `json:"updated_at"`
}

// PartyUpdateData is the typed request payload for Party.UpdateTyped.
type PartyUpdateData struct {
	Id string `json:"id"`
	AdditionalDeclarationType *map[string]any `json:"additional_declaration_type,omitempty"`
	Address *map[string]any `json:"address,omitempty"`
	Authorisation *map[string]any `json:"authorisation,omitempty"`
	BankDetails *string `json:"bank_details,omitempty"`
	Certificate *map[string]any `json:"certificate,omitempty"`
	CertificateType *string `json:"certificate_type,omitempty"`
	Company *string `json:"company,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CustomsOfficeOfLodgement *map[string]any `json:"customs_office_of_lodgement,omitempty"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	Email *string `json:"email,omitempty"`
	IdentificationNumber *string `json:"identification_number,omitempty"`
	IndirectRepresentative *bool `json:"indirect_representative,omitempty"`
	Name *string `json:"name,omitempty"`
	NhdLastSubmissionYear *int `json:"nhd_last_submission_year,omitempty"`
	NhdSubmissionCounter *int `json:"nhd_submission_counter,omitempty"`
	PersonPayingCustomsDuty *string `json:"person_paying_customs_duty,omitempty"`
	PhoneCountryCode *string `json:"phone_country_code,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	PreferredPaymentMethod *map[string]any `json:"preferred_payment_method,omitempty"`
	SignedForm *map[string]any `json:"signed_form,omitempty"`
	Type *string `json:"type,omitempty"`
	TypeOfPerson *map[string]any `json:"type_of_person,omitempty"`
	Unlocode *string `json:"unlocode,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// PartyRemoveMatch is the typed request payload for Party.RemoveTyped.
type PartyRemoveMatch struct {
	Id string `json:"id"`
}

// Submission is the typed data model for the submission entity.
type Submission struct {
	AdditionalExternalIds *[]any `json:"additional_external_ids,omitempty"`
	AmendmentReason *string `json:"amendment_reason,omitempty"`
	AmendmentStatus *string `json:"amendment_status,omitempty"`
	Answers []any `json:"answers"`
	BypassRestrictedCode *bool `json:"bypass_restricted_code,omitempty"`
	ClearanceSlip map[string]any `json:"clearance_slip"`
	Client map[string]any `json:"client"`
	CompanyMember map[string]any `json:"company_member"`
	Consignee map[string]any `json:"consignee"`
	Consignor map[string]any `json:"consignor"`
	CreatedAt string `json:"created_at"`
	Declarant map[string]any `json:"declarant"`
	DocumentUploadStatus *string `json:"document_upload_status,omitempty"`
	DocumentsPresentationRequested *bool `json:"documents_presentation_requested,omitempty"`
	DocumentsUploadRequested *bool `json:"documents_upload_requested,omitempty"`
	ExternalId *string `json:"external_id,omitempty"`
	Form string `json:"form"`
	GoodsPresentationStatus *string `json:"goods_presentation_status,omitempty"`
	HrcmStatus *string `json:"hrcm_status,omitempty"`
	Id *string `json:"id,omitempty"`
	InvalidationStatus *string `json:"invalidation_status,omitempty"`
	IsGlobalTemplate *bool `json:"is_global_template,omitempty"`
	LatestNotificationItem *string `json:"latest_notification_item,omitempty"`
	LatestState map[string]any `json:"latest_state"`
	Lrn *string `json:"lrn,omitempty"`
	Mrn *string `json:"mrn,omitempty"`
	Name *string `json:"name,omitempty"`
	PartialAnswers *bool `json:"partial_answers,omitempty"`
	Receipt map[string]any `json:"receipt"`
	RefundApplicationStatus *string `json:"refund_application_status,omitempty"`
	Route *string `json:"route,omitempty"`
	ShipmentItemsNo *int `json:"shipment_items_no,omitempty"`
	ShipmentItemsQuantityNo *int `json:"shipment_items_quantity_no,omitempty"`
	Source *string `json:"source,omitempty"`
	SourceType *string `json:"source_type,omitempty"`
	Status *string `json:"status,omitempty"`
	Template *bool `json:"template,omitempty"`
	TemplateId *string `json:"template_id,omitempty"`
	TemplateProperties *[]any `json:"template_properties,omitempty"`
	TotalTaxAmount *string `json:"total_tax_amount,omitempty"`
	UpdatedAt string `json:"updated_at"`
	VerificationErrors *[]any `json:"verification_errors,omitempty"`
	VerificationStatus *string `json:"verification_status,omitempty"`
}

// SubmissionLoadMatch is the typed request payload for Submission.LoadTyped.
type SubmissionLoadMatch struct {
	Id string `json:"id"`
}

// SubmissionListMatch is the typed request payload for Submission.ListTyped.
type SubmissionListMatch struct {
	AdditionalExternalIds *[]any `json:"additional_external_ids,omitempty"`
	AmendmentReason *string `json:"amendment_reason,omitempty"`
	AmendmentStatus *string `json:"amendment_status,omitempty"`
	Answers *[]any `json:"answers,omitempty"`
	BypassRestrictedCode *bool `json:"bypass_restricted_code,omitempty"`
	ClearanceSlip *map[string]any `json:"clearance_slip,omitempty"`
	Client *map[string]any `json:"client,omitempty"`
	CompanyMember *map[string]any `json:"company_member,omitempty"`
	Consignee *map[string]any `json:"consignee,omitempty"`
	Consignor *map[string]any `json:"consignor,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Declarant *map[string]any `json:"declarant,omitempty"`
	DocumentUploadStatus *string `json:"document_upload_status,omitempty"`
	DocumentsPresentationRequested *bool `json:"documents_presentation_requested,omitempty"`
	DocumentsUploadRequested *bool `json:"documents_upload_requested,omitempty"`
	ExternalId *string `json:"external_id,omitempty"`
	Form *string `json:"form,omitempty"`
	GoodsPresentationStatus *string `json:"goods_presentation_status,omitempty"`
	HrcmStatus *string `json:"hrcm_status,omitempty"`
	Id *string `json:"id,omitempty"`
	InvalidationStatus *string `json:"invalidation_status,omitempty"`
	IsGlobalTemplate *bool `json:"is_global_template,omitempty"`
	LatestNotificationItem *string `json:"latest_notification_item,omitempty"`
	LatestState *map[string]any `json:"latest_state,omitempty"`
	Lrn *string `json:"lrn,omitempty"`
	Mrn *string `json:"mrn,omitempty"`
	Name *string `json:"name,omitempty"`
	PartialAnswers *bool `json:"partial_answers,omitempty"`
	Receipt *map[string]any `json:"receipt,omitempty"`
	RefundApplicationStatus *string `json:"refund_application_status,omitempty"`
	Route *string `json:"route,omitempty"`
	ShipmentItemsNo *int `json:"shipment_items_no,omitempty"`
	ShipmentItemsQuantityNo *int `json:"shipment_items_quantity_no,omitempty"`
	Source *string `json:"source,omitempty"`
	SourceType *string `json:"source_type,omitempty"`
	Status *string `json:"status,omitempty"`
	Template *bool `json:"template,omitempty"`
	TemplateId *string `json:"template_id,omitempty"`
	TemplateProperties *[]any `json:"template_properties,omitempty"`
	TotalTaxAmount *string `json:"total_tax_amount,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	VerificationErrors *[]any `json:"verification_errors,omitempty"`
	VerificationStatus *string `json:"verification_status,omitempty"`
}

// SubmissionCreateData is the typed request payload for Submission.CreateTyped.
type SubmissionCreateData struct {
	AdditionalExternalIds *[]any `json:"additional_external_ids,omitempty"`
	AmendmentReason *string `json:"amendment_reason,omitempty"`
	AmendmentStatus *string `json:"amendment_status,omitempty"`
	Answers []any `json:"answers"`
	BypassRestrictedCode *bool `json:"bypass_restricted_code,omitempty"`
	ClearanceSlip map[string]any `json:"clearance_slip"`
	Client map[string]any `json:"client"`
	CompanyMember map[string]any `json:"company_member"`
	Consignee map[string]any `json:"consignee"`
	Consignor map[string]any `json:"consignor"`
	CreatedAt string `json:"created_at"`
	Declarant map[string]any `json:"declarant"`
	DocumentUploadStatus *string `json:"document_upload_status,omitempty"`
	DocumentsPresentationRequested *bool `json:"documents_presentation_requested,omitempty"`
	DocumentsUploadRequested *bool `json:"documents_upload_requested,omitempty"`
	ExternalId *string `json:"external_id,omitempty"`
	Form string `json:"form"`
	GoodsPresentationStatus *string `json:"goods_presentation_status,omitempty"`
	HrcmStatus *string `json:"hrcm_status,omitempty"`
	Id *string `json:"id,omitempty"`
	InvalidationStatus *string `json:"invalidation_status,omitempty"`
	IsGlobalTemplate *bool `json:"is_global_template,omitempty"`
	LatestNotificationItem *string `json:"latest_notification_item,omitempty"`
	LatestState map[string]any `json:"latest_state"`
	Lrn *string `json:"lrn,omitempty"`
	Mrn *string `json:"mrn,omitempty"`
	Name *string `json:"name,omitempty"`
	PartialAnswers *bool `json:"partial_answers,omitempty"`
	Receipt map[string]any `json:"receipt"`
	RefundApplicationStatus *string `json:"refund_application_status,omitempty"`
	Route *string `json:"route,omitempty"`
	ShipmentItemsNo *int `json:"shipment_items_no,omitempty"`
	ShipmentItemsQuantityNo *int `json:"shipment_items_quantity_no,omitempty"`
	Source *string `json:"source,omitempty"`
	SourceType *string `json:"source_type,omitempty"`
	Status *string `json:"status,omitempty"`
	Template *bool `json:"template,omitempty"`
	TemplateId *string `json:"template_id,omitempty"`
	TemplateProperties *[]any `json:"template_properties,omitempty"`
	TotalTaxAmount *string `json:"total_tax_amount,omitempty"`
	UpdatedAt string `json:"updated_at"`
	VerificationErrors *[]any `json:"verification_errors,omitempty"`
	VerificationStatus *string `json:"verification_status,omitempty"`
}

// SubmissionUpdateData is the typed request payload for Submission.UpdateTyped.
type SubmissionUpdateData struct {
	Id string `json:"id"`
	AdditionalExternalIds *[]any `json:"additional_external_ids,omitempty"`
	AmendmentReason *string `json:"amendment_reason,omitempty"`
	AmendmentStatus *string `json:"amendment_status,omitempty"`
	Answers *[]any `json:"answers,omitempty"`
	BypassRestrictedCode *bool `json:"bypass_restricted_code,omitempty"`
	ClearanceSlip *map[string]any `json:"clearance_slip,omitempty"`
	Client *map[string]any `json:"client,omitempty"`
	CompanyMember *map[string]any `json:"company_member,omitempty"`
	Consignee *map[string]any `json:"consignee,omitempty"`
	Consignor *map[string]any `json:"consignor,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Declarant *map[string]any `json:"declarant,omitempty"`
	DocumentUploadStatus *string `json:"document_upload_status,omitempty"`
	DocumentsPresentationRequested *bool `json:"documents_presentation_requested,omitempty"`
	DocumentsUploadRequested *bool `json:"documents_upload_requested,omitempty"`
	ExternalId *string `json:"external_id,omitempty"`
	Form *string `json:"form,omitempty"`
	GoodsPresentationStatus *string `json:"goods_presentation_status,omitempty"`
	HrcmStatus *string `json:"hrcm_status,omitempty"`
	InvalidationStatus *string `json:"invalidation_status,omitempty"`
	IsGlobalTemplate *bool `json:"is_global_template,omitempty"`
	LatestNotificationItem *string `json:"latest_notification_item,omitempty"`
	LatestState *map[string]any `json:"latest_state,omitempty"`
	Lrn *string `json:"lrn,omitempty"`
	Mrn *string `json:"mrn,omitempty"`
	Name *string `json:"name,omitempty"`
	PartialAnswers *bool `json:"partial_answers,omitempty"`
	Receipt *map[string]any `json:"receipt,omitempty"`
	RefundApplicationStatus *string `json:"refund_application_status,omitempty"`
	Route *string `json:"route,omitempty"`
	ShipmentItemsNo *int `json:"shipment_items_no,omitempty"`
	ShipmentItemsQuantityNo *int `json:"shipment_items_quantity_no,omitempty"`
	Source *string `json:"source,omitempty"`
	SourceType *string `json:"source_type,omitempty"`
	Status *string `json:"status,omitempty"`
	Template *bool `json:"template,omitempty"`
	TemplateId *string `json:"template_id,omitempty"`
	TemplateProperties *[]any `json:"template_properties,omitempty"`
	TotalTaxAmount *string `json:"total_tax_amount,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	VerificationErrors *[]any `json:"verification_errors,omitempty"`
	VerificationStatus *string `json:"verification_status,omitempty"`
}

// SubmissionRemoveMatch is the typed request payload for Submission.RemoveTyped.
type SubmissionRemoveMatch struct {
	Id string `json:"id"`
}

// SubmissionDetail is the typed data model for the submission_detail entity.
type SubmissionDetail struct {
	AdditionalExternalIds *[]any `json:"additional_external_ids,omitempty"`
	AdditionalInformation []any `json:"additional_information"`
	AmendmentStatus *string `json:"amendment_status,omitempty"`
	ClearanceSlip map[string]any `json:"clearance_slip"`
	Client map[string]any `json:"client"`
	Company string `json:"company"`
	CompanyMember map[string]any `json:"company_member"`
	Consignee map[string]any `json:"consignee"`
	Consignor map[string]any `json:"consignor"`
	CreatedAt string `json:"created_at"`
	Declarant map[string]any `json:"declarant"`
	DocumentUploadStatus *string `json:"document_upload_status,omitempty"`
	DocumentsPresentationRequested *bool `json:"documents_presentation_requested,omitempty"`
	DocumentsUploadRequested *bool `json:"documents_upload_requested,omitempty"`
	ExternalId *string `json:"external_id,omitempty"`
	Form string `json:"form"`
	GoodsPresentationStatus *string `json:"goods_presentation_status,omitempty"`
	HrcmStatus *string `json:"hrcm_status,omitempty"`
	Id *string `json:"id,omitempty"`
	InvalidationStatus *string `json:"invalidation_status,omitempty"`
	IsGlobalTemplate *bool `json:"is_global_template,omitempty"`
	LatestNotificationItem *string `json:"latest_notification_item,omitempty"`
	LatestState map[string]any `json:"latest_state"`
	Lrn *string `json:"lrn,omitempty"`
	Mrn *string `json:"mrn,omitempty"`
	Name *string `json:"name,omitempty"`
	Receipt map[string]any `json:"receipt"`
	RefundApplicationStatus *string `json:"refund_application_status,omitempty"`
	Route *string `json:"route,omitempty"`
	ShipmentItemsNo *int `json:"shipment_items_no,omitempty"`
	ShipmentItemsQuantityNo *int `json:"shipment_items_quantity_no,omitempty"`
	Source *string `json:"source,omitempty"`
	SourceType *string `json:"source_type,omitempty"`
	Status *string `json:"status,omitempty"`
	Submission string `json:"submission"`
	SupportingDocuments *[]any `json:"supporting_documents,omitempty"`
	Template *bool `json:"template,omitempty"`
	TotalTaxAmount *string `json:"total_tax_amount,omitempty"`
	UpdatedAt string `json:"updated_at"`
	VerificationErrors *[]any `json:"verification_errors,omitempty"`
	VerificationStatus *string `json:"verification_status,omitempty"`
}

// SubmissionDetailCreateData is the typed request payload for SubmissionDetail.CreateTyped.
type SubmissionDetailCreateData struct {
	AdditionalExternalIds *[]any `json:"additional_external_ids,omitempty"`
	AdditionalInformation []any `json:"additional_information"`
	AmendmentStatus *string `json:"amendment_status,omitempty"`
	ClearanceSlip map[string]any `json:"clearance_slip"`
	Client map[string]any `json:"client"`
	Company string `json:"company"`
	CompanyMember map[string]any `json:"company_member"`
	Consignee map[string]any `json:"consignee"`
	Consignor map[string]any `json:"consignor"`
	CreatedAt string `json:"created_at"`
	Declarant map[string]any `json:"declarant"`
	DocumentUploadStatus *string `json:"document_upload_status,omitempty"`
	DocumentsPresentationRequested *bool `json:"documents_presentation_requested,omitempty"`
	DocumentsUploadRequested *bool `json:"documents_upload_requested,omitempty"`
	ExternalId *string `json:"external_id,omitempty"`
	Form string `json:"form"`
	GoodsPresentationStatus *string `json:"goods_presentation_status,omitempty"`
	HrcmStatus *string `json:"hrcm_status,omitempty"`
	Id *string `json:"id,omitempty"`
	InvalidationStatus *string `json:"invalidation_status,omitempty"`
	IsGlobalTemplate *bool `json:"is_global_template,omitempty"`
	LatestNotificationItem *string `json:"latest_notification_item,omitempty"`
	LatestState map[string]any `json:"latest_state"`
	Lrn *string `json:"lrn,omitempty"`
	Mrn *string `json:"mrn,omitempty"`
	Name *string `json:"name,omitempty"`
	Receipt map[string]any `json:"receipt"`
	RefundApplicationStatus *string `json:"refund_application_status,omitempty"`
	Route *string `json:"route,omitempty"`
	ShipmentItemsNo *int `json:"shipment_items_no,omitempty"`
	ShipmentItemsQuantityNo *int `json:"shipment_items_quantity_no,omitempty"`
	Source *string `json:"source,omitempty"`
	SourceType *string `json:"source_type,omitempty"`
	Status *string `json:"status,omitempty"`
	Submission string `json:"submission"`
	SupportingDocuments *[]any `json:"supporting_documents,omitempty"`
	Template *bool `json:"template,omitempty"`
	TotalTaxAmount *string `json:"total_tax_amount,omitempty"`
	UpdatedAt string `json:"updated_at"`
	VerificationErrors *[]any `json:"verification_errors,omitempty"`
	VerificationStatus *string `json:"verification_status,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
