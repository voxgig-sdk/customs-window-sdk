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
			"slug": "customs-window",
			"version": "0.1.1",
			"target": "go",
		},
		"feature": map[string]any{
			"debug": map[string]any{
				"options": map[string]any{
					"active": false,
					"max": 100,
					"redact": []any{
						"authorization",
						"cookie",
						"set-cookie",
						"api-key",
						"apikey",
						"x-api-key",
						"idempotency-key",
					},
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"onEntry": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"idempotency": map[string]any{
				"options": map[string]any{
					"active": false,
					"header": "Idempotency-Key",
					"methods": []any{
						"POST",
						"PUT",
						"PATCH",
						"DELETE",
					},
					"ops": []any{
						"create",
						"update",
						"remove",
					},
				},
				"optspec": map[string]any{
					"keygen": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"metrics": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "none",
			},
			"paging": map[string]any{
				"options": map[string]any{
					"active": false,
					"afterVar": "after",
					"cursorParam": "cursor",
					"firstVar": "first",
					"limitParam": "limit",
					"pageParam": "page",
					"startPage": 1,
				},
				"optspec": map[string]any{
					"limit": "`$NUMBER`",
					"ops": "`$LIST`",
				},
				"strict": false,
				"transport": "none",
			},
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
		},
		"options": map[string]any{
			"base": "https://api.customswindow.com",
			"auth": map[string]any{
				"prefix": "",
				"name": "Basic Auth",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"bulk_upload": map[string]any{},
				"file": map[string]any{},
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
						"title": "Active Transport Nationality",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "active_transport_number",
						"title": "Active Transport Number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "arrival_datetime",
						"title": "Arrival Datetime",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "client",
						"title": "Client",
						"type": "`$STRING`",
						"req": true,
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "company_member",
						"title": "Company Member",
						"type": "`$STRING`",
						"req": true,
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "declarant",
						"title": "Declarant",
						"type": "`$STRING`",
						"req": true,
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "declarations_no",
						"title": "Declarations No",
						"type": "`$INTEGER`",
						"req": true,
					},
					map[string]any{
						"name": "deleted_at",
						"title": "Deleted At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "departure_datetime",
						"title": "Departure Datetime",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "errors_file",
						"title": "Errors File",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "external_id",
						"title": "External Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "failed_declarations_no",
						"title": "Failed Declarations No",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "file",
						"title": "File",
						"type": "`$OBJECT`",
						"req": true,
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "green_routed_no",
						"title": "Green Routed No",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "h1_fallback_template",
						"title": "H1 Fallback Template",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "house_transport_doc_ref",
						"title": "House Transport Doc Ref",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "issue_date",
						"title": "Issue Date",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "mapping",
						"title": "Mapping",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "orange_routed_no",
						"title": "Orange Routed No",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "parsed_declarations_no",
						"title": "Parsed Declarations No",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "parser",
						"title": "Parser",
						"type": "`$STRING`",
						"short": "* `aes_platform` - aes_platform * `cds_platform` - cds_platform * `cds_export_platform` - cds_export_platform * `g4_g3` - g4_g3 * `nhd_platform` - nhd_platform * `platform` - platform * `birds` - birds * `ics2_platform` - ics2_platform",
					},
					map[string]any{
						"name": "parsing_completed_at",
						"title": "Parsing Completed At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "parsing_started_at",
						"title": "Parsing Started At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "passive_transport_nationality",
						"title": "Passive Transport Nationality",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "passive_transport_number",
						"title": "Passive Transport Number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "processed_declarations_no",
						"title": "Processed Declarations No",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "processing_ended_at",
						"title": "Processing Ended At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "processing_started_at",
						"title": "Processing Started At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "receipt_generating_started_at",
						"title": "Receipt Generating Started At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "receipt_request_started_at",
						"title": "Receipt Request Started At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "receipt_request_status",
						"title": "Receipt Request Status",
						"type": "`$STRING`",
						"short": "* `pending` - pending * `processing` - processing * `generating` - generating * `completed` - completed",
					},
					map[string]any{
						"name": "receipt_request_user",
						"title": "Receipt Request User",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "receipts_zip",
						"title": "Receipts Zip",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "red_routed_no",
						"title": "Red Routed No",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "rejected_status_no",
						"title": "Rejected Status No",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "* `pending` - pending * `parsing` - parsing * `processing` - processing * `complete` - complete",
					},
					map[string]any{
						"name": "template",
						"title": "Template",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "yellow_routed_no",
						"title": "Yellow Routed No",
						"type": "`$INTEGER`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "bulk_upload",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/bulk-uploads",
								"segments": []any{
									map[string]any{
										"lit": "bulk-uploads",
									},
								},
								"parts": []any{
									"bulk-uploads",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/bulk-uploads",
								"segments": []any{
									map[string]any{
										"lit": "bulk-uploads",
									},
								},
								"parts": []any{
									"bulk-uploads",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.results`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "cursor",
											"orig": "cursor",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"cursor",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/bulk-uploads/{id}",
								"segments": []any{
									map[string]any{
										"lit": "bulk-uploads",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"bulk-uploads",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/bulk-uploads/{id}/generate-pdfs",
								"segments": []any{
									map[string]any{
										"lit": "bulk-uploads",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "generate-pdfs",
									},
								},
								"parts": []any{
									"bulk-uploads",
									"{id}",
									"generate-pdfs",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "generate_pdf",
									"exist": []any{
										"id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/bulk-uploads/{id}",
								"segments": []any{
									map[string]any{
										"lit": "bulk-uploads",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"bulk-uploads",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/bulk-uploads/{id}/generate-pdfs",
								"segments": []any{
									map[string]any{
										"lit": "bulk-uploads",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "generate-pdfs",
									},
								},
								"parts": []any{
									"bulk-uploads",
									"{id}",
									"generate-pdfs",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "generate_pdf",
									"exist": []any{
										"id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/bulk-uploads/{id}",
								"segments": []any{
									map[string]any{
										"lit": "bulk-uploads",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"bulk-uploads",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/bulk-uploads/{id}/generate-pdfs",
								"segments": []any{
									map[string]any{
										"lit": "bulk-uploads",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "generate-pdfs",
									},
								},
								"parts": []any{
									"bulk-uploads",
									"{id}",
									"generate-pdfs",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "generate_pdf",
									"exist": []any{
										"id",
									},
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
						"title": "Company",
						"type": "`$STRING`",
						"req": true,
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "extension",
						"title": "Extension",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
					},
					map[string]any{
						"name": "file",
						"title": "File",
						"type": "`$STRING`",
						"req": true,
						"format": "uri",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
					},
					map[string]any{
						"name": "public",
						"title": "Public",
						"type": "`$BOOLEAN`",
						"req": true,
						"op": map[string]any{
							"create": map[string]any{
								"type": "`$BOOLEAN`",
							},
						},
					},
					map[string]any{
						"name": "size",
						"title": "Size",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "url",
						"title": "Url",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "file",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/files",
								"segments": []any{
									map[string]any{
										"lit": "files",
									},
								},
								"parts": []any{
									"files",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"file": "`reqdata`",
									},
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"party": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "additional_declaration_type",
						"title": "Additional Declaration Type",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "address",
						"title": "Address",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "authorisation",
						"title": "Authorisation",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "bank_details",
						"title": "Bank Details",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "certificate",
						"title": "Certificate",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "certificate_type",
						"title": "Certificate Type",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
					},
					map[string]any{
						"name": "company",
						"title": "Company",
						"type": "`$STRING`",
						"req": true,
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "customs_office_of_lodgement",
						"title": "Customs Office Of Lodgement",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "deleted_at",
						"title": "Deleted At",
						"type": "`$STRING`",
						"format": "date-time",
					},
					map[string]any{
						"name": "email",
						"title": "Email",
						"type": "`$STRING`",
						"format": "email",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "identification_number",
						"title": "Identification Number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "indirect_representative",
						"title": "Indirect Representative",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "nhd_last_submission_year",
						"title": "Nhd Last Submission Year",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "nhd_submission_counter",
						"title": "Nhd Submission Counter",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "person_paying_customs_duty",
						"title": "Person Paying Customs Duty",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "phone_country_code",
						"title": "Phone Country Code",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "phone_number",
						"title": "Phone Number",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "preferred_payment_method",
						"title": "Preferred Payment Method",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "signed_form",
						"title": "Signed Form",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "type",
						"title": "Type",
						"type": "`$STRING`",
						"short": "* `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `co…",
					},
					map[string]any{
						"name": "type_of_person",
						"title": "Type Of Person",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "unlocode",
						"title": "Unlocode",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
						"format": "date-time",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "party",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/parties",
								"segments": []any{
									map[string]any{
										"lit": "parties",
									},
								},
								"parts": []any{
									"parties",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/parties",
								"segments": []any{
									map[string]any{
										"lit": "parties",
									},
								},
								"parts": []any{
									"parties",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.results`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "cursor",
											"orig": "cursor",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "type",
											"orig": "type",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"cursor",
										"type",
									},
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/parties/{id}",
								"segments": []any{
									map[string]any{
										"lit": "parties",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"parties",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/parties/{id}",
								"segments": []any{
									map[string]any{
										"lit": "parties",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"parties",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/parties/{id}",
								"segments": []any{
									map[string]any{
										"lit": "parties",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"parties",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
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
						"title": "Additional External Ids",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "amendment_reason",
						"title": "Amendment Reason",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "amendment_status",
						"title": "Amendment Status",
						"type": "`$STRING`",
						"short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
					},
					map[string]any{
						"name": "answers",
						"title": "Answers",
						"type": "`$ARRAY`",
						"req": true,
						"op": map[string]any{
							"update": map[string]any{
								"type": "`$ARRAY`",
							},
						},
					},
					map[string]any{
						"name": "bypass_restricted_code",
						"title": "Bypass Restricted Code",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "clearance_slip",
						"title": "Clearance Slip",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "client",
						"title": "Client",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "company_member",
						"title": "Company Member",
						"type": "`$OBJECT`",
						"req": true,
						"op": map[string]any{
							"update": map[string]any{
								"type": "`$STRING`",
							},
						},
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "consignee",
						"title": "Consignee",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "consignor",
						"title": "Consignor",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "declarant",
						"title": "Declarant",
						"type": "`$OBJECT`",
						"req": true,
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "document_upload_status",
						"title": "Document Upload Status",
						"type": "`$STRING`",
						"short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
					},
					map[string]any{
						"name": "documents_presentation_requested",
						"title": "Documents Presentation Requested",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "documents_upload_requested",
						"title": "Documents Upload Requested",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "external_id",
						"title": "External Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "form",
						"title": "Form",
						"type": "`$STRING`",
						"req": true,
						"short": "cuid-format identifier for this entity.",
						"readOnly": true,
					},
					map[string]any{
						"name": "goods_presentation_status",
						"title": "Goods Presentation Status",
						"type": "`$STRING`",
						"short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested",
					},
					map[string]any{
						"name": "hrcm_status",
						"title": "Hrcm Status",
						"type": "`$STRING`",
						"short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "invalidation_status",
						"title": "Invalidation Status",
						"type": "`$STRING`",
						"short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled",
					},
					map[string]any{
						"name": "is_global_template",
						"title": "Is Global Template",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "latest_notification_item",
						"title": "Latest Notification Item",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "latest_state",
						"title": "Latest State",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "lrn",
						"title": "Lrn",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "mrn",
						"title": "Mrn",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "partial_answers",
						"title": "Partial Answers",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "receipt",
						"title": "Receipt",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "refund_application_status",
						"title": "Refund Application Status",
						"type": "`$STRING`",
						"short": "* `processing` - processing * `rejected` - rejected * `accepted` - accepted",
					},
					map[string]any{
						"name": "route",
						"title": "Route",
						"type": "`$STRING`",
						"short": "* `green` - green * `orange` - orange * `red` - red * `yellow` - yellow",
					},
					map[string]any{
						"name": "shipment_items_no",
						"title": "Shipment Items No",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "shipment_items_quantity_no",
						"title": "Shipment Items Quantity No",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "source",
						"title": "Source",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "source_type",
						"title": "Source Type",
						"type": "`$STRING`",
						"short": "* `template` - template * `automated_import` - automated_import",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "* `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re…",
					},
					map[string]any{
						"name": "template",
						"title": "Template",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "template_id",
						"title": "Template Id",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "template_properties",
						"title": "Template Properties",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "total_tax_amount",
						"title": "Total Tax Amount",
						"type": "`$STRING`",
						"format": "decimal",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "verification_errors",
						"title": "Verification Errors",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "verification_status",
						"title": "Verification Status",
						"type": "`$STRING`",
						"short": "* `passed` - passed * `failed` - failed",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "submission",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/submissions/{id}/refund",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "refund",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
									"refund",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "refund",
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/submissions",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
								},
								"parts": []any{
									"submissions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/submissions/retrieve",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"lit": "retrieve",
									},
								},
								"parts": []any{
									"submissions",
									"retrieve",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{
									"$action": "retrieve",
								},
							},
						},
					},
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/submissions",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
								},
								"parts": []any{
									"submissions",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.results`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "cursor",
											"orig": "cursor",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "form_subtype",
											"orig": "form_subtype",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "route",
											"orig": "route",
											"type": "`$STRING`",
											"kind": "query",
											"example": "green",
										},
										map[string]any{
											"name": "status",
											"orig": "status",
											"type": "`$STRING`",
											"kind": "query",
											"example": "released",
										},
										map[string]any{
											"name": "template",
											"orig": "template",
											"type": "`$BOOLEAN`",
											"kind": "query",
										},
									},
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
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/submissions/{id}",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/submissions/{id}/clearance-slip",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "clearance-slip",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
									"clearance-slip",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "clearance_slip",
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/submissions/{id}/notification-read",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "notification-read",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
									"notification-read",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "notification_read",
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/submissions/{id}/pbn-applicable",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "pbn-applicable",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
									"pbn-applicable",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "pbn_applicable",
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/submissions/{id}/receipt",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "receipt",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
									"receipt",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "receipt",
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/submissions/{id}/refund",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "refund",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
									"refund",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "refund",
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/submissions/retrieve",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"lit": "retrieve",
									},
								},
								"parts": []any{
									"submissions",
									"retrieve",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{
									"$action": "retrieve",
								},
							},
						},
					},
					"remove": map[string]any{
						"input": "data",
						"name": "remove",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/submissions/{id}",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/submissions/{id}/clearance-slip",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "clearance-slip",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
									"clearance-slip",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "clearance_slip",
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/submissions/{id}/notification-read",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "notification-read",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
									"notification-read",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "notification_read",
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/submissions/{id}/pbn-applicable",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "pbn-applicable",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
									"pbn-applicable",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "pbn_applicable",
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "DELETE",
								"orig": "/submissions/{id}/receipt",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "receipt",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
									"receipt",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "receipt",
									"exist": []any{
										"id",
									},
								},
							},
						},
					},
					"update": map[string]any{
						"input": "data",
						"name": "update",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/submissions/{id}",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/submissions/{id}/clearance-slip",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "clearance-slip",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
									"clearance-slip",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "clearance_slip",
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/submissions/{id}/notification-read",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "notification-read",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
									"notification-read",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "notification_read",
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/submissions/{id}/pbn-applicable",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "pbn-applicable",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
									"pbn-applicable",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "pbn_applicable",
									"exist": []any{
										"id",
									},
								},
							},
							map[string]any{
								"kind": "http",
								"method": "PATCH",
								"orig": "/submissions/{id}/receipt",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
									map[string]any{
										"lit": "receipt",
									},
								},
								"parts": []any{
									"submissions",
									"{id}",
									"receipt",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"name": "id",
											"orig": "id",
											"type": "`$STRING`",
											"kind": "param",
											"reqd": true,
										},
									},
								},
								"select": map[string]any{
									"$action": "receipt",
									"exist": []any{
										"id",
									},
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
						"title": "Additional External Ids",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "additional_information",
						"title": "Additional Information",
						"type": "`$ARRAY`",
						"req": true,
					},
					map[string]any{
						"name": "amendment_status",
						"title": "Amendment Status",
						"type": "`$STRING`",
						"short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
					},
					map[string]any{
						"name": "clearance_slip",
						"title": "Clearance Slip",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "client",
						"title": "Client",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "company",
						"title": "Company",
						"type": "`$STRING`",
						"req": true,
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "company_member",
						"title": "Company Member",
						"type": "`$OBJECT`",
						"req": true,
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "consignee",
						"title": "Consignee",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "consignor",
						"title": "Consignor",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "created_at",
						"title": "Created At",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "declarant",
						"title": "Declarant",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "document_upload_status",
						"title": "Document Upload Status",
						"type": "`$STRING`",
						"short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
					},
					map[string]any{
						"name": "documents_presentation_requested",
						"title": "Documents Presentation Requested",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "documents_upload_requested",
						"title": "Documents Upload Requested",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "external_id",
						"title": "External Id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "form",
						"title": "Form",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
					},
					map[string]any{
						"name": "goods_presentation_status",
						"title": "Goods Presentation Status",
						"type": "`$STRING`",
						"short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested",
					},
					map[string]any{
						"name": "hrcm_status",
						"title": "Hrcm Status",
						"type": "`$STRING`",
						"short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
					},
					map[string]any{
						"name": "id",
						"title": "Id",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "invalidation_status",
						"title": "Invalidation Status",
						"type": "`$STRING`",
						"short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled",
					},
					map[string]any{
						"name": "is_global_template",
						"title": "Is Global Template",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "latest_notification_item",
						"title": "Latest Notification Item",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "latest_state",
						"title": "Latest State",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "lrn",
						"title": "Lrn",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "mrn",
						"title": "Mrn",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"title": "Name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "receipt",
						"title": "Receipt",
						"type": "`$OBJECT`",
						"req": true,
					},
					map[string]any{
						"name": "refund_application_status",
						"title": "Refund Application Status",
						"type": "`$STRING`",
						"short": "* `processing` - processing * `rejected` - rejected * `accepted` - accepted",
					},
					map[string]any{
						"name": "route",
						"title": "Route",
						"type": "`$STRING`",
						"short": "* `green` - green * `orange` - orange * `red` - red * `yellow` - yellow",
					},
					map[string]any{
						"name": "shipment_items_no",
						"title": "Shipment Items No",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "shipment_items_quantity_no",
						"title": "Shipment Items Quantity No",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "source",
						"title": "Source",
						"type": "`$STRING`",
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "source_type",
						"title": "Source Type",
						"type": "`$STRING`",
						"short": "* `template` - template * `automated_import` - automated_import",
					},
					map[string]any{
						"name": "status",
						"title": "Status",
						"type": "`$STRING`",
						"short": "* `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re…",
					},
					map[string]any{
						"name": "submission",
						"title": "Submission",
						"type": "`$STRING`",
						"req": true,
						"short": "cuid-format identifier for this entity.",
					},
					map[string]any{
						"name": "supporting_documents",
						"title": "Supporting Documents",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "template",
						"title": "Template",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "total_tax_amount",
						"title": "Total Tax Amount",
						"type": "`$STRING`",
						"format": "decimal",
					},
					map[string]any{
						"name": "updated_at",
						"title": "Updated At",
						"type": "`$STRING`",
						"req": true,
						"readOnly": true,
						"format": "date-time",
					},
					map[string]any{
						"name": "verification_errors",
						"title": "Verification Errors",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "verification_status",
						"title": "Verification Status",
						"type": "`$STRING`",
						"short": "* `passed` - passed * `failed` - failed",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "submission_detail",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "POST",
								"orig": "/documents-request",
								"segments": []any{
									map[string]any{
										"lit": "documents-request",
									},
								},
								"parts": []any{
									"documents-request",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
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

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
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
	case "debug":
		if NewDebugFeatureFunc != nil {
			return NewDebugFeatureFunc()
		}
	case "idempotency":
		if NewIdempotencyFeatureFunc != nil {
			return NewIdempotencyFeatureFunc()
		}
	case "metrics":
		if NewMetricsFeatureFunc != nil {
			return NewMetricsFeatureFunc()
		}
	case "paging":
		if NewPagingFeatureFunc != nil {
			return NewPagingFeatureFunc()
		}
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
