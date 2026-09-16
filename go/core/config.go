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
			"version": "0.0.1",
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
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "active_transport_number",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "arrival_datetime",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "client",
						"req": true,
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "company_member",
						"req": true,
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "created_at",
						"readOnly": true,
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "declarant",
						"req": true,
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "declarations_no",
						"req": true,
						"type": "`$INTEGER`",
					},
					map[string]any{
						"format": "date-time",
						"name": "deleted_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
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
						"short": "cuid-format identifier for this entity.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "green_routed_no",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "h1_fallback_template",
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "house_transport_doc_ref",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "issue_date",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "mapping",
						"short": "cuid-format identifier for this entity.",
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
						"short": "* `aes_platform` - aes_platform * `cds_platform` - cds_platform * `cds_export_platform` - cds_export_platform * `g4_g3` - g4_g3 * `nhd_platform` - nhd_platform * `platform` - platform * `birds` - birds * `ics2_platform` - ics2_platform",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "parsing_completed_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "parsing_started_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "passive_transport_nationality",
						"short": "cuid-format identifier for this entity.",
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
						"format": "date-time",
						"name": "processing_ended_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "processing_started_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "receipt_generating_started_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "receipt_request_started_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "receipt_request_status",
						"short": "* `pending` - pending * `processing` - processing * `generating` - generating * `completed` - completed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "receipt_request_user",
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "receipts_zip",
						"short": "cuid-format identifier for this entity.",
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
						"short": "* `pending` - pending * `parsing` - parsing * `processing` - processing * `complete` - complete",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "template",
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "updated_at",
						"readOnly": true,
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "yellow_routed_no",
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
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/bulk-uploads",
								"segments": []any{
									map[string]any{
										"lit": "bulk-uploads",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"bulk-uploads",
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
								"segments": []any{
									map[string]any{
										"lit": "bulk-uploads",
									},
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
								"parts": []any{
									"bulk-uploads",
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
								"segments": []any{
									map[string]any{
										"lit": "bulk-uploads",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"bulk-uploads",
									"{id}",
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
								"parts": []any{
									"bulk-uploads",
									"{id}",
									"generate-pdfs",
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
								"segments": []any{
									map[string]any{
										"lit": "bulk-uploads",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"bulk-uploads",
									"{id}",
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
								"parts": []any{
									"bulk-uploads",
									"{id}",
									"generate-pdfs",
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
								"segments": []any{
									map[string]any{
										"lit": "bulk-uploads",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"bulk-uploads",
									"{id}",
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
								"parts": []any{
									"bulk-uploads",
									"{id}",
									"generate-pdfs",
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
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "created_at",
						"readOnly": true,
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "extension",
						"readOnly": true,
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "uri",
						"name": "file",
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "name",
						"readOnly": true,
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
						"format": "date-time",
						"name": "updated_at",
						"readOnly": true,
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "url",
						"readOnly": true,
						"req": true,
						"type": "`$STRING`",
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
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/files",
								"segments": []any{
									map[string]any{
										"lit": "files",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": map[string]any{
										"file": "`reqdata`",
									},
									"res": "`body`",
								},
								"parts": []any{
									"files",
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
						"readOnly": true,
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "company",
						"req": true,
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "created_at",
						"readOnly": true,
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "customs_office_of_lodgement",
						"req": true,
						"type": "`$OBJECT`",
					},
					map[string]any{
						"format": "date-time",
						"name": "deleted_at",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "email",
						"name": "email",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "cuid-format identifier for this entity.",
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
						"short": "* `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `co…",
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
						"format": "date-time",
						"name": "updated_at",
						"readOnly": true,
						"req": true,
						"type": "`$STRING`",
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
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/parties",
								"segments": []any{
									map[string]any{
										"lit": "parties",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"parties",
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
								"segments": []any{
									map[string]any{
										"lit": "parties",
									},
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
								"parts": []any{
									"parties",
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
								"segments": []any{
									map[string]any{
										"lit": "parties",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"parties",
									"{id}",
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
								"segments": []any{
									map[string]any{
										"lit": "parties",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"parties",
									"{id}",
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
								"segments": []any{
									map[string]any{
										"lit": "parties",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"parties",
									"{id}",
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
						"short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
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
						"short": "cuid-format identifier for this entity.",
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
						"format": "date-time",
						"name": "created_at",
						"readOnly": true,
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "declarant",
						"req": true,
						"short": "cuid-format identifier for this entity.",
						"type": "`$OBJECT`",
					},
					map[string]any{
						"name": "document_upload_status",
						"short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
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
						"readOnly": true,
						"req": true,
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "goods_presentation_status",
						"short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hrcm_status",
						"short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "invalidation_status",
						"short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_global_template",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "latest_notification_item",
						"short": "cuid-format identifier for this entity.",
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
						"short": "* `processing` - processing * `rejected` - rejected * `accepted` - accepted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "route",
						"short": "* `green` - green * `orange` - orange * `red` - red * `yellow` - yellow",
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
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_type",
						"short": "* `template` - template * `automated_import` - automated_import",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"short": "* `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re…",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "template",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "template_id",
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "template_properties",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"format": "decimal",
						"name": "total_tax_amount",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "updated_at",
						"readOnly": true,
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "verification_errors",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "verification_status",
						"short": "* `passed` - passed * `failed` - failed",
						"type": "`$STRING`",
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
								"parts": []any{
									"submissions",
									"{id}",
									"refund",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/submissions",
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"submissions",
								},
							},
							map[string]any{
								"args": map[string]any{},
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
								"select": map[string]any{
									"$action": "retrieve",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"submissions",
									"retrieve",
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
								"segments": []any{
									map[string]any{
										"lit": "submissions",
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
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.results`",
								},
								"parts": []any{
									"submissions",
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
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"submissions",
									"{id}",
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
								"parts": []any{
									"submissions",
									"{id}",
									"clearance-slip",
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
								"parts": []any{
									"submissions",
									"{id}",
									"notification-read",
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
								"parts": []any{
									"submissions",
									"{id}",
									"pbn-applicable",
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
								"parts": []any{
									"submissions",
									"{id}",
									"receipt",
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
								"parts": []any{
									"submissions",
									"{id}",
									"refund",
								},
							},
							map[string]any{
								"args": map[string]any{},
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
								"select": map[string]any{
									"$action": "retrieve",
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"submissions",
									"retrieve",
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
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"submissions",
									"{id}",
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
								"parts": []any{
									"submissions",
									"{id}",
									"clearance-slip",
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
								"parts": []any{
									"submissions",
									"{id}",
									"notification-read",
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
								"parts": []any{
									"submissions",
									"{id}",
									"pbn-applicable",
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
								"parts": []any{
									"submissions",
									"{id}",
									"receipt",
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
								"segments": []any{
									map[string]any{
										"lit": "submissions",
									},
									map[string]any{
										"var": "id",
									},
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
								"parts": []any{
									"submissions",
									"{id}",
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
								"parts": []any{
									"submissions",
									"{id}",
									"clearance-slip",
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
								"parts": []any{
									"submissions",
									"{id}",
									"notification-read",
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
								"parts": []any{
									"submissions",
									"{id}",
									"pbn-applicable",
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
								"parts": []any{
									"submissions",
									"{id}",
									"receipt",
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
						"short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
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
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "company_member",
						"req": true,
						"short": "cuid-format identifier for this entity.",
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
						"format": "date-time",
						"name": "created_at",
						"readOnly": true,
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
						"short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
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
						"readOnly": true,
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "goods_presentation_status",
						"short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "hrcm_status",
						"short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "invalidation_status",
						"short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "is_global_template",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "latest_notification_item",
						"short": "cuid-format identifier for this entity.",
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
						"short": "* `processing` - processing * `rejected` - rejected * `accepted` - accepted",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "route",
						"short": "* `green` - green * `orange` - orange * `red` - red * `yellow` - yellow",
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
						"short": "cuid-format identifier for this entity.",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "source_type",
						"short": "* `template` - template * `automated_import` - automated_import",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "status",
						"short": "* `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re…",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "submission",
						"req": true,
						"short": "cuid-format identifier for this entity.",
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
						"format": "decimal",
						"name": "total_tax_amount",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "updated_at",
						"readOnly": true,
						"req": true,
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "verification_errors",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "verification_status",
						"short": "* `passed` - passed * `failed` - failed",
						"type": "`$STRING`",
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
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/documents-request",
								"segments": []any{
									map[string]any{
										"lit": "documents-request",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"documents-request",
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
