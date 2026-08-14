package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "CustomsWindow",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://api.customswindow.com",
			"auth": map[string]any{
				"prefix": "",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"bulk_upload": map[string]any{},
				"file": map[string]any{},
				"paginated_bulk_upload_list_list": map[string]any{},
				"paginated_party_list_list": map[string]any{},
				"paginated_submission_list_list": map[string]any{},
				"party": map[string]any{},
				"submission": map[string]any{},
				"submission_detail": map[string]any{},
			},
		},
		"entity": map[string]any{
			"bulk_upload": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "active_transport_nationality",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "active_transport_number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "arrival_datetime",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "client",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "company_member",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_at",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "declarant",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "declarations_no",
						"req": true,
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "deleted_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "departure_datetime",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "errors_file",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "external_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "failed_declarations_no",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "file",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "green_routed_no",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "h1_fallback_template",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "house_transport_doc_ref",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "issue_date",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "mapping",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "orange_routed_no",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "parsed_declarations_no",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "parser",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "parsing_completed_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "parsing_started_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "passive_transport_nationality",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "passive_transport_number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "processed_declarations_no",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "processing_ended_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "processing_started_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "receipt_generating_started_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "receipt_request_started_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "receipt_request_status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "receipt_request_user",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "receipts_zip",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "red_routed_no",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "rejected_status_no",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "template",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "yellow_routed_no",
						"type": "`$INTEGER`",
					},
				},
				"name": "bulk_upload",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/bulk-uploads",
								"parts": []any{
									"bulk-uploads",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "cursor",
											"orig": "cursor",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/bulk-uploads",
								"parts": []any{
									"bulk-uploads",
								},
								"select": map[string]any{
									"exist": []any{
										"cursor",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.results`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/bulk-uploads/{id}",
								"parts": []any{
									"bulk-uploads",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/bulk-uploads/{id}/generate-pdfs",
								"parts": []any{
									"bulk-uploads",
									"{id}",
									"generate-pdfs",
								},
								"select": map[string]any{
									"$action": "generate_pdf",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/bulk-uploads/{id}",
								"parts": []any{
									"bulk-uploads",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/bulk-uploads/{id}/generate-pdfs",
								"parts": []any{
									"bulk-uploads",
									"{id}",
									"generate-pdfs",
								},
								"select": map[string]any{
									"$action": "generate_pdf",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/bulk-uploads/{id}",
								"parts": []any{
									"bulk-uploads",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/bulk-uploads/{id}/generate-pdfs",
								"parts": []any{
									"bulk-uploads",
									"{id}",
									"generate-pdfs",
								},
								"select": map[string]any{
									"$action": "generate_pdf",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"file": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "company",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_at",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "extension",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "file",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "public",
						"op": map[string]any{
							"create": map[string]any{
								"type": "`$BOOLEAN`",
							},
						},
						"req": true,
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "size",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "updated_at",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "url",
						"req": true,
						"type": "`$STRING`",
					},
				},
				"name": "file",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/files",
								"parts": []any{
									"files",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"file": "`reqdata`",
									},
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"paginated_bulk_upload_list_list": map[string]any{
				"fields": []any{},
				"name": "paginated_bulk_upload_list_list",
				"op": map[string]any{},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"paginated_party_list_list": map[string]any{
				"fields": []any{},
				"name": "paginated_party_list_list",
				"op": map[string]any{},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"paginated_submission_list_list": map[string]any{
				"fields": []any{},
				"name": "paginated_submission_list_list",
				"op": map[string]any{},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"party": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "additional_declaration_type",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "address",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "authorisation",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "bank_details",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "certificate",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "certificate_type",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "company",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "created_at",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customs_office_of_lodgement",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "deleted_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "email",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "identification_number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "indirect_representative",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nhd_last_submission_year",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "nhd_submission_counter",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "person_paying_customs_duty",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "phone_country_code",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "phone_number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "preferred_payment_method",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "signed_form",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type_of_person",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "unlocode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"req": true,
						"type": "`$STRING`",
					},
				},
				"name": "party",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/parties",
								"parts": []any{
									"parties",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "cursor",
											"orig": "cursor",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "type",
											"orig": "type",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/parties",
								"parts": []any{
									"parties",
								},
								"select": map[string]any{
									"exist": []any{
										"cursor",
										"type",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.results`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/parties/{id}",
								"parts": []any{
									"parties",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/parties/{id}",
								"parts": []any{
									"parties",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/parties/{id}",
								"parts": []any{
									"parties",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"submission": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "additional_external_ids",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "amendment_reason",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "amendment_status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "answers",
						"op": map[string]any{
							"update": map[string]any{
								"type": "`$ARRAY`",
							},
						},
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "bypass_restricted_code",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "clearance_slip",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "client",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "company_member",
						"op": map[string]any{
							"update": map[string]any{
								"type": "`$STRING`",
							},
						},
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "consignee",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "consignor",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "created_at",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "declarant",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "document_upload_status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "documents_presentation_requested",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "documents_upload_requested",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "external_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "form",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "goods_presentation_status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hrcm_status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "invalidation_status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_global_template",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "latest_notification_item",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "latest_state",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "lrn",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "mrn",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "partial_answers",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "receipt",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "refund_application_status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "route",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shipment_items_no",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "shipment_items_quantity_no",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "template",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "template_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "template_properties",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "total_tax_amount",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "verification_errors",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "verification_status",
						"type": "`$STRING`",
					},
				},
				"name": "submission",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "POST",
								"orig": "/submissions/{id}/refund",
								"parts": []any{
									"submissions",
									"{id}",
									"refund",
								},
								"select": map[string]any{
									"$action": "refund",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/submissions",
								"parts": []any{
									"submissions",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/submissions/retrieve",
								"parts": []any{
									"submissions",
									"retrieve",
								},
								"select": map[string]any{
									"$action": "retrieve",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "cursor",
											"orig": "cursor",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "form_subtype",
											"orig": "form_subtype",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "green",
											"kind": "query",
											"name": "route",
											"orig": "route",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "released",
											"kind": "query",
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "template",
											"orig": "template",
											"type": "`$BOOLEAN`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/submissions",
								"parts": []any{
									"submissions",
								},
								"select": map[string]any{
									"exist": []any{
										"cursor",
										"form_subtype",
										"route",
										"status",
										"template",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.results`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/submissions/{id}",
								"parts": []any{
									"submissions",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/submissions/{id}/clearance-slip",
								"parts": []any{
									"submissions",
									"{id}",
									"clearance-slip",
								},
								"select": map[string]any{
									"$action": "clearance_slip",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/submissions/{id}/notification-read",
								"parts": []any{
									"submissions",
									"{id}",
									"notification-read",
								},
								"select": map[string]any{
									"$action": "notification_read",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/submissions/{id}/pbn-applicable",
								"parts": []any{
									"submissions",
									"{id}",
									"pbn-applicable",
								},
								"select": map[string]any{
									"$action": "pbn_applicable",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/submissions/{id}/receipt",
								"parts": []any{
									"submissions",
									"{id}",
									"receipt",
								},
								"select": map[string]any{
									"$action": "receipt",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/submissions/{id}/refund",
								"parts": []any{
									"submissions",
									"{id}",
									"refund",
								},
								"select": map[string]any{
									"$action": "refund",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/submissions/retrieve",
								"parts": []any{
									"submissions",
									"retrieve",
								},
								"select": map[string]any{
									"$action": "retrieve",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/submissions/{id}",
								"parts": []any{
									"submissions",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/submissions/{id}/clearance-slip",
								"parts": []any{
									"submissions",
									"{id}",
									"clearance-slip",
								},
								"select": map[string]any{
									"$action": "clearance_slip",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/submissions/{id}/notification-read",
								"parts": []any{
									"submissions",
									"{id}",
									"notification-read",
								},
								"select": map[string]any{
									"$action": "notification_read",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/submissions/{id}/pbn-applicable",
								"parts": []any{
									"submissions",
									"{id}",
									"pbn-applicable",
								},
								"select": map[string]any{
									"$action": "pbn_applicable",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "DELETE",
								"orig": "/submissions/{id}/receipt",
								"parts": []any{
									"submissions",
									"{id}",
									"receipt",
								},
								"select": map[string]any{
									"$action": "receipt",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/submissions/{id}",
								"parts": []any{
									"submissions",
									"{id}",
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/submissions/{id}/clearance-slip",
								"parts": []any{
									"submissions",
									"{id}",
									"clearance-slip",
								},
								"select": map[string]any{
									"$action": "clearance_slip",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/submissions/{id}/notification-read",
								"parts": []any{
									"submissions",
									"{id}",
									"notification-read",
								},
								"select": map[string]any{
									"$action": "notification_read",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/submissions/{id}/pbn-applicable",
								"parts": []any{
									"submissions",
									"{id}",
									"pbn-applicable",
								},
								"select": map[string]any{
									"$action": "pbn_applicable",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"kind": "param",
											"name": "id",
											"orig": "id",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "PATCH",
								"orig": "/submissions/{id}/receipt",
								"parts": []any{
									"submissions",
									"{id}",
									"receipt",
								},
								"select": map[string]any{
									"$action": "receipt",
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"submission_detail": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "additional_external_ids",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "additional_information",
						"req": true,
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "amendment_status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "clearance_slip",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "client",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "company",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "company_member",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "consignee",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "consignor",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "created_at",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "declarant",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "document_upload_status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "documents_presentation_requested",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "documents_upload_requested",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "external_id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "form",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "goods_presentation_status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hrcm_status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "invalidation_status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_global_template",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "latest_notification_item",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "latest_state",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "lrn",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "mrn",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "receipt",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "refund_application_status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "route",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "shipment_items_no",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "shipment_items_quantity_no",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_type",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "submission",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "supporting_documents",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "template",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "total_tax_amount",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "verification_errors",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "verification_status",
						"type": "`$STRING`",
					},
				},
				"name": "submission_detail",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/documents-request",
								"parts": []any{
									"documents-request",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
