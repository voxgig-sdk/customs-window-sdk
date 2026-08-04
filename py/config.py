# CustomsWindow SDK configuration


def make_config():
    return {
        "main": {
            "name": "CustomsWindow",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
      },
        },
        "options": {
            "base": "https://api.customswindow.com",
            "auth": {
                "prefix": "",
            },
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "bulk_upload": {},
                "file": {},
                "paginated_bulk_upload_list_list": {},
                "paginated_party_list_list": {},
                "paginated_submission_list_list": {},
                "party": {},
                "submission": {},
                "submission_detail": {},
            },
        },
        "entity": {
      "bulk_upload": {
        "fields": [
          {
            "active": True,
            "name": "active_transport_nationality",
            "req": False,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "active_transport_number",
            "req": False,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "arrival_datetime",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "client",
            "req": True,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "company_member",
            "req": True,
            "type": "`$STRING`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "created_at",
            "req": True,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "declarant",
            "req": True,
            "type": "`$STRING`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "declarations_no",
            "req": True,
            "type": "`$INTEGER`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "deleted_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "departure_datetime",
            "req": False,
            "type": "`$STRING`",
            "index$": 9,
          },
          {
            "active": True,
            "name": "errors_file",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 10,
          },
          {
            "active": True,
            "name": "external_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 11,
          },
          {
            "active": True,
            "name": "failed_declarations_no",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 12,
          },
          {
            "active": True,
            "name": "file",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 13,
          },
          {
            "active": True,
            "name": "green_routed_no",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 14,
          },
          {
            "active": True,
            "name": "h1_fallback_template",
            "req": False,
            "type": "`$STRING`",
            "index$": 15,
          },
          {
            "active": True,
            "name": "house_transport_doc_ref",
            "req": False,
            "type": "`$STRING`",
            "index$": 16,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 17,
          },
          {
            "active": True,
            "name": "issue_date",
            "req": False,
            "type": "`$STRING`",
            "index$": 18,
          },
          {
            "active": True,
            "name": "mapping",
            "req": False,
            "type": "`$STRING`",
            "index$": 19,
          },
          {
            "active": True,
            "name": "orange_routed_no",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 20,
          },
          {
            "active": True,
            "name": "parsed_declarations_no",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 21,
          },
          {
            "active": True,
            "name": "parser",
            "req": False,
            "type": "`$STRING`",
            "index$": 22,
          },
          {
            "active": True,
            "name": "parsing_completed_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 23,
          },
          {
            "active": True,
            "name": "parsing_started_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 24,
          },
          {
            "active": True,
            "name": "passive_transport_nationality",
            "req": False,
            "type": "`$STRING`",
            "index$": 25,
          },
          {
            "active": True,
            "name": "passive_transport_number",
            "req": False,
            "type": "`$STRING`",
            "index$": 26,
          },
          {
            "active": True,
            "name": "processed_declarations_no",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 27,
          },
          {
            "active": True,
            "name": "processing_ended_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 28,
          },
          {
            "active": True,
            "name": "processing_started_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 29,
          },
          {
            "active": True,
            "name": "receipt_generating_started_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 30,
          },
          {
            "active": True,
            "name": "receipt_request_started_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 31,
          },
          {
            "active": True,
            "name": "receipt_request_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 32,
          },
          {
            "active": True,
            "name": "receipt_request_user",
            "req": False,
            "type": "`$STRING`",
            "index$": 33,
          },
          {
            "active": True,
            "name": "receipts_zip",
            "req": False,
            "type": "`$STRING`",
            "index$": 34,
          },
          {
            "active": True,
            "name": "red_routed_no",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 35,
          },
          {
            "active": True,
            "name": "rejected_status_no",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 36,
          },
          {
            "active": True,
            "name": "status",
            "req": False,
            "type": "`$STRING`",
            "index$": 37,
          },
          {
            "active": True,
            "name": "template",
            "req": False,
            "type": "`$STRING`",
            "index$": 38,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": True,
            "type": "`$STRING`",
            "index$": 39,
          },
          {
            "active": True,
            "name": "yellow_routed_no",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 40,
          },
        ],
        "name": "bulk_upload",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {},
                "method": "POST",
                "orig": "/bulk-uploads",
                "parts": [
                  "bulk-uploads",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "create",
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "cursor",
                      "orig": "cursor",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/bulk-uploads",
                "parts": [
                  "bulk-uploads",
                ],
                "select": {
                  "exist": [
                    "cursor",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "list",
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/bulk-uploads/{id}",
                "parts": [
                  "bulk-uploads",
                  "{id}",
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/bulk-uploads/{id}/generate-pdfs",
                "parts": [
                  "bulk-uploads",
                  "{id}",
                  "generate-pdfs",
                ],
                "select": {
                  "$action": "generate_pdf",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
            ],
            "key$": "load",
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/bulk-uploads/{id}",
                "parts": [
                  "bulk-uploads",
                  "{id}",
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/bulk-uploads/{id}/generate-pdfs",
                "parts": [
                  "bulk-uploads",
                  "{id}",
                  "generate-pdfs",
                ],
                "select": {
                  "$action": "generate_pdf",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
            ],
            "key$": "remove",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/bulk-uploads/{id}",
                "parts": [
                  "bulk-uploads",
                  "{id}",
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/bulk-uploads/{id}/generate-pdfs",
                "parts": [
                  "bulk-uploads",
                  "{id}",
                  "generate-pdfs",
                ],
                "select": {
                  "$action": "generate_pdf",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "file": {
        "fields": [
          {
            "active": True,
            "name": "company",
            "req": True,
            "type": "`$STRING`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "created_at",
            "req": True,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "extension",
            "req": True,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "file",
            "req": True,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "name",
            "req": True,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "public",
            "op": {
              "create": {
                "req": False,
                "type": "`$BOOLEAN`",
              },
            },
            "req": True,
            "type": "`$BOOLEAN`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "size",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": True,
            "type": "`$STRING`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "url",
            "req": True,
            "type": "`$STRING`",
            "index$": 9,
          },
        ],
        "name": "file",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {},
                "method": "POST",
                "orig": "/files",
                "parts": [
                  "files",
                ],
                "select": {},
                "transform": {
                  "req": {
                    "file": "`reqdata`",
                  },
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "create",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "paginated_bulk_upload_list_list": {
        "fields": [],
        "name": "paginated_bulk_upload_list_list",
        "op": {},
        "relations": {
          "ancestors": [],
        },
      },
      "paginated_party_list_list": {
        "fields": [],
        "name": "paginated_party_list_list",
        "op": {},
        "relations": {
          "ancestors": [],
        },
      },
      "paginated_submission_list_list": {
        "fields": [],
        "name": "paginated_submission_list_list",
        "op": {},
        "relations": {
          "ancestors": [],
        },
      },
      "party": {
        "fields": [
          {
            "active": True,
            "name": "additional_declaration_type",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "address",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "authorisation",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "bank_detail",
            "req": False,
            "type": "`$STRING`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "certificate",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "certificate_type",
            "req": True,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "company",
            "req": True,
            "type": "`$STRING`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "created_at",
            "req": True,
            "type": "`$STRING`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "customs_office_of_lodgement",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "deleted_at",
            "req": False,
            "type": "`$STRING`",
            "index$": 9,
          },
          {
            "active": True,
            "name": "email",
            "req": False,
            "type": "`$STRING`",
            "index$": 10,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 11,
          },
          {
            "active": True,
            "name": "identification_number",
            "req": False,
            "type": "`$STRING`",
            "index$": 12,
          },
          {
            "active": True,
            "name": "indirect_representative",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 13,
          },
          {
            "active": True,
            "name": "name",
            "req": False,
            "type": "`$STRING`",
            "index$": 14,
          },
          {
            "active": True,
            "name": "nhd_last_submission_year",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 15,
          },
          {
            "active": True,
            "name": "nhd_submission_counter",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 16,
          },
          {
            "active": True,
            "name": "person_paying_customs_duty",
            "req": False,
            "type": "`$STRING`",
            "index$": 17,
          },
          {
            "active": True,
            "name": "phone_country_code",
            "req": False,
            "type": "`$STRING`",
            "index$": 18,
          },
          {
            "active": True,
            "name": "phone_number",
            "req": False,
            "type": "`$STRING`",
            "index$": 19,
          },
          {
            "active": True,
            "name": "preferred_payment_method",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 20,
          },
          {
            "active": True,
            "name": "signed_form",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 21,
          },
          {
            "active": True,
            "name": "type",
            "req": False,
            "type": "`$STRING`",
            "index$": 22,
          },
          {
            "active": True,
            "name": "type_of_person",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 23,
          },
          {
            "active": True,
            "name": "unlocode",
            "req": False,
            "type": "`$STRING`",
            "index$": 24,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": True,
            "type": "`$STRING`",
            "index$": 25,
          },
        ],
        "name": "party",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {},
                "method": "POST",
                "orig": "/parties",
                "parts": [
                  "parties",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "create",
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "cursor",
                      "orig": "cursor",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "type",
                      "orig": "type",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/parties",
                "parts": [
                  "parties",
                ],
                "select": {
                  "exist": [
                    "cursor",
                    "type",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "list",
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/parties/{id}",
                "parts": [
                  "parties",
                  "{id}",
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "load",
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/parties/{id}",
                "parts": [
                  "parties",
                  "{id}",
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "remove",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/parties/{id}",
                "parts": [
                  "parties",
                  "{id}",
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "submission": {
        "fields": [
          {
            "active": True,
            "name": "additional_external_id",
            "req": False,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "amendment_reason",
            "req": False,
            "type": "`$STRING`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "amendment_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "answer",
            "op": {
              "update": {
                "req": False,
                "type": "`$ARRAY`",
              },
            },
            "req": True,
            "type": "`$ARRAY`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "bypass_restricted_code",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "clearance_slip",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "client",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "company_member",
            "op": {
              "update": {
                "req": False,
                "type": "`$STRING`",
              },
            },
            "req": True,
            "type": "`$OBJECT`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "consignee",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "consignor",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 9,
          },
          {
            "active": True,
            "name": "created_at",
            "req": True,
            "type": "`$STRING`",
            "index$": 10,
          },
          {
            "active": True,
            "name": "declarant",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 11,
          },
          {
            "active": True,
            "name": "document_upload_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 12,
          },
          {
            "active": True,
            "name": "documents_presentation_requested",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 13,
          },
          {
            "active": True,
            "name": "documents_upload_requested",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 14,
          },
          {
            "active": True,
            "name": "external_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 15,
          },
          {
            "active": True,
            "name": "form",
            "req": True,
            "type": "`$STRING`",
            "index$": 16,
          },
          {
            "active": True,
            "name": "goods_presentation_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 17,
          },
          {
            "active": True,
            "name": "hrcm_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 18,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 19,
          },
          {
            "active": True,
            "name": "invalidation_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 20,
          },
          {
            "active": True,
            "name": "is_global_template",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 21,
          },
          {
            "active": True,
            "name": "latest_notification_item",
            "req": False,
            "type": "`$STRING`",
            "index$": 22,
          },
          {
            "active": True,
            "name": "latest_state",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 23,
          },
          {
            "active": True,
            "name": "lrn",
            "req": False,
            "type": "`$STRING`",
            "index$": 24,
          },
          {
            "active": True,
            "name": "mrn",
            "req": False,
            "type": "`$STRING`",
            "index$": 25,
          },
          {
            "active": True,
            "name": "name",
            "req": False,
            "type": "`$STRING`",
            "index$": 26,
          },
          {
            "active": True,
            "name": "partial_answer",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 27,
          },
          {
            "active": True,
            "name": "receipt",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 28,
          },
          {
            "active": True,
            "name": "refund_application_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 29,
          },
          {
            "active": True,
            "name": "route",
            "req": False,
            "type": "`$STRING`",
            "index$": 30,
          },
          {
            "active": True,
            "name": "shipment_items_no",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 31,
          },
          {
            "active": True,
            "name": "shipment_items_quantity_no",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 32,
          },
          {
            "active": True,
            "name": "source",
            "req": False,
            "type": "`$STRING`",
            "index$": 33,
          },
          {
            "active": True,
            "name": "source_type",
            "req": False,
            "type": "`$STRING`",
            "index$": 34,
          },
          {
            "active": True,
            "name": "status",
            "req": False,
            "type": "`$STRING`",
            "index$": 35,
          },
          {
            "active": True,
            "name": "template",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 36,
          },
          {
            "active": True,
            "name": "template_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 37,
          },
          {
            "active": True,
            "name": "template_property",
            "req": False,
            "type": "`$ARRAY`",
            "index$": 38,
          },
          {
            "active": True,
            "name": "total_tax_amount",
            "req": False,
            "type": "`$STRING`",
            "index$": 39,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": True,
            "type": "`$STRING`",
            "index$": 40,
          },
          {
            "active": True,
            "name": "verification_error",
            "req": False,
            "type": "`$ARRAY`",
            "index$": 41,
          },
          {
            "active": True,
            "name": "verification_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 42,
          },
        ],
        "name": "submission",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "POST",
                "orig": "/submissions/{id}/refund",
                "parts": [
                  "submissions",
                  "{id}",
                  "refund",
                ],
                "select": {
                  "$action": "refund",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {},
                "method": "POST",
                "orig": "/submissions",
                "parts": [
                  "submissions",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {},
                "method": "POST",
                "orig": "/submissions/retrieve",
                "parts": [
                  "submissions",
                  "retrieve",
                ],
                "select": {
                  "$action": "retrieve",
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
            ],
            "key$": "create",
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "active": True,
                "args": {
                  "query": [
                    {
                      "active": True,
                      "kind": "query",
                      "name": "cursor",
                      "orig": "cursor",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "form_subtype",
                      "orig": "form_subtype",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "example": "green",
                      "kind": "query",
                      "name": "route",
                      "orig": "route",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "example": "released",
                      "kind": "query",
                      "name": "status",
                      "orig": "status",
                      "reqd": False,
                      "type": "`$STRING`",
                    },
                    {
                      "active": True,
                      "kind": "query",
                      "name": "template",
                      "orig": "template",
                      "reqd": False,
                      "type": "`$BOOLEAN`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/submissions",
                "parts": [
                  "submissions",
                ],
                "select": {
                  "exist": [
                    "cursor",
                    "form_subtype",
                    "route",
                    "status",
                    "template",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "list",
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "GET",
                "orig": "/submissions/{id}",
                "parts": [
                  "submissions",
                  "{id}",
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/submissions/{id}/clearance-slip",
                "parts": [
                  "submissions",
                  "{id}",
                  "clearance-slip",
                ],
                "select": {
                  "$action": "clearance_slip",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/submissions/{id}/notification-read",
                "parts": [
                  "submissions",
                  "{id}",
                  "notification-read",
                ],
                "select": {
                  "$action": "notification_read",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/submissions/{id}/pbn-applicable",
                "parts": [
                  "submissions",
                  "{id}",
                  "pbn-applicable",
                ],
                "select": {
                  "$action": "pbn_applicable",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 3,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/submissions/{id}/receipt",
                "parts": [
                  "submissions",
                  "{id}",
                  "receipt",
                ],
                "select": {
                  "$action": "receipt",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 4,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "GET",
                "orig": "/submissions/{id}/refund",
                "parts": [
                  "submissions",
                  "{id}",
                  "refund",
                ],
                "select": {
                  "$action": "refund",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 5,
              },
              {
                "active": True,
                "args": {},
                "method": "GET",
                "orig": "/submissions/retrieve",
                "parts": [
                  "submissions",
                  "retrieve",
                ],
                "select": {
                  "$action": "retrieve",
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 6,
              },
            ],
            "key$": "load",
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/submissions/{id}",
                "parts": [
                  "submissions",
                  "{id}",
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/submissions/{id}/clearance-slip",
                "parts": [
                  "submissions",
                  "{id}",
                  "clearance-slip",
                ],
                "select": {
                  "$action": "clearance_slip",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/submissions/{id}/notification-read",
                "parts": [
                  "submissions",
                  "{id}",
                  "notification-read",
                ],
                "select": {
                  "$action": "notification_read",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/submissions/{id}/pbn-applicable",
                "parts": [
                  "submissions",
                  "{id}",
                  "pbn-applicable",
                ],
                "select": {
                  "$action": "pbn_applicable",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 3,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "DELETE",
                "orig": "/submissions/{id}/receipt",
                "parts": [
                  "submissions",
                  "{id}",
                  "receipt",
                ],
                "select": {
                  "$action": "receipt",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 4,
              },
            ],
            "key$": "remove",
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                      "index$": 0,
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/submissions/{id}",
                "parts": [
                  "submissions",
                  "{id}",
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/submissions/{id}/clearance-slip",
                "parts": [
                  "submissions",
                  "{id}",
                  "clearance-slip",
                ],
                "select": {
                  "$action": "clearance_slip",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 1,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/submissions/{id}/notification-read",
                "parts": [
                  "submissions",
                  "{id}",
                  "notification-read",
                ],
                "select": {
                  "$action": "notification_read",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 2,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/submissions/{id}/pbn-applicable",
                "parts": [
                  "submissions",
                  "{id}",
                  "pbn-applicable",
                ],
                "select": {
                  "$action": "pbn_applicable",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 3,
              },
              {
                "active": True,
                "args": {
                  "params": [
                    {
                      "active": True,
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "method": "PATCH",
                "orig": "/submissions/{id}/receipt",
                "parts": [
                  "submissions",
                  "{id}",
                  "receipt",
                ],
                "select": {
                  "$action": "receipt",
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 4,
              },
            ],
            "key$": "update",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "submission_detail": {
        "fields": [
          {
            "active": True,
            "name": "additional_external_id",
            "req": False,
            "type": "`$ARRAY`",
            "index$": 0,
          },
          {
            "active": True,
            "name": "additional_information",
            "req": True,
            "type": "`$ARRAY`",
            "index$": 1,
          },
          {
            "active": True,
            "name": "amendment_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 2,
          },
          {
            "active": True,
            "name": "clearance_slip",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 3,
          },
          {
            "active": True,
            "name": "client",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 4,
          },
          {
            "active": True,
            "name": "company",
            "req": True,
            "type": "`$STRING`",
            "index$": 5,
          },
          {
            "active": True,
            "name": "company_member",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 6,
          },
          {
            "active": True,
            "name": "consignee",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 7,
          },
          {
            "active": True,
            "name": "consignor",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 8,
          },
          {
            "active": True,
            "name": "created_at",
            "req": True,
            "type": "`$STRING`",
            "index$": 9,
          },
          {
            "active": True,
            "name": "declarant",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 10,
          },
          {
            "active": True,
            "name": "document_upload_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 11,
          },
          {
            "active": True,
            "name": "documents_presentation_requested",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 12,
          },
          {
            "active": True,
            "name": "documents_upload_requested",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 13,
          },
          {
            "active": True,
            "name": "external_id",
            "req": False,
            "type": "`$STRING`",
            "index$": 14,
          },
          {
            "active": True,
            "name": "form",
            "req": True,
            "type": "`$STRING`",
            "index$": 15,
          },
          {
            "active": True,
            "name": "goods_presentation_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 16,
          },
          {
            "active": True,
            "name": "hrcm_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 17,
          },
          {
            "active": True,
            "name": "id",
            "req": False,
            "type": "`$STRING`",
            "index$": 18,
          },
          {
            "active": True,
            "name": "invalidation_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 19,
          },
          {
            "active": True,
            "name": "is_global_template",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 20,
          },
          {
            "active": True,
            "name": "latest_notification_item",
            "req": False,
            "type": "`$STRING`",
            "index$": 21,
          },
          {
            "active": True,
            "name": "latest_state",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 22,
          },
          {
            "active": True,
            "name": "lrn",
            "req": False,
            "type": "`$STRING`",
            "index$": 23,
          },
          {
            "active": True,
            "name": "mrn",
            "req": False,
            "type": "`$STRING`",
            "index$": 24,
          },
          {
            "active": True,
            "name": "name",
            "req": False,
            "type": "`$STRING`",
            "index$": 25,
          },
          {
            "active": True,
            "name": "receipt",
            "req": True,
            "type": "`$OBJECT`",
            "index$": 26,
          },
          {
            "active": True,
            "name": "refund_application_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 27,
          },
          {
            "active": True,
            "name": "route",
            "req": False,
            "type": "`$STRING`",
            "index$": 28,
          },
          {
            "active": True,
            "name": "shipment_items_no",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 29,
          },
          {
            "active": True,
            "name": "shipment_items_quantity_no",
            "req": False,
            "type": "`$INTEGER`",
            "index$": 30,
          },
          {
            "active": True,
            "name": "source",
            "req": False,
            "type": "`$STRING`",
            "index$": 31,
          },
          {
            "active": True,
            "name": "source_type",
            "req": False,
            "type": "`$STRING`",
            "index$": 32,
          },
          {
            "active": True,
            "name": "status",
            "req": False,
            "type": "`$STRING`",
            "index$": 33,
          },
          {
            "active": True,
            "name": "submission",
            "req": True,
            "type": "`$STRING`",
            "index$": 34,
          },
          {
            "active": True,
            "name": "supporting_document",
            "req": False,
            "type": "`$ARRAY`",
            "index$": 35,
          },
          {
            "active": True,
            "name": "template",
            "req": False,
            "type": "`$BOOLEAN`",
            "index$": 36,
          },
          {
            "active": True,
            "name": "total_tax_amount",
            "req": False,
            "type": "`$STRING`",
            "index$": 37,
          },
          {
            "active": True,
            "name": "updated_at",
            "req": True,
            "type": "`$STRING`",
            "index$": 38,
          },
          {
            "active": True,
            "name": "verification_error",
            "req": False,
            "type": "`$ARRAY`",
            "index$": 39,
          },
          {
            "active": True,
            "name": "verification_status",
            "req": False,
            "type": "`$STRING`",
            "index$": 40,
          },
        ],
        "name": "submission_detail",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "active": True,
                "args": {},
                "method": "POST",
                "orig": "/documents-request",
                "parts": [
                  "documents-request",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "index$": 0,
              },
            ],
            "key$": "create",
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
