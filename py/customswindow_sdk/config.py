# CustomsWindow SDK configuration


# The sekreto plugin DEFINITIONS the model selected per feature, imported
# above by name from the modules the catalogue's active `plugin.def`
# entries declare. Handed to each feature (secrets builds its Sekreto
# with them): a provider kind not listed here is unknown to that SDK.
FEATURE_PLUGINS = {
}


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "CustomsWindow",
            "slug": "customs-window",
            "version": "0.1.1",
            "target": "py",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
        "transport": "base",
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
            "name": "active_transport_nationality",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "active_transport_number",
            "type": "`$STRING`",
          },
          {
            "format": "date-time",
            "name": "arrival_datetime",
            "type": "`$STRING`",
          },
          {
            "name": "client",
            "req": True,
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "company_member",
            "req": True,
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "format": "date-time",
            "name": "created_at",
            "readOnly": True,
            "req": True,
            "type": "`$STRING`",
          },
          {
            "name": "declarant",
            "req": True,
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "declarations_no",
            "req": True,
            "type": "`$INTEGER`",
          },
          {
            "format": "date-time",
            "name": "deleted_at",
            "type": "`$STRING`",
          },
          {
            "format": "date-time",
            "name": "departure_datetime",
            "type": "`$STRING`",
          },
          {
            "name": "errors_file",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "external_id",
            "type": "`$STRING`",
          },
          {
            "name": "failed_declarations_no",
            "type": "`$INTEGER`",
          },
          {
            "name": "file",
            "req": True,
            "short": "cuid-format identifier for this entity.",
            "type": "`$OBJECT`",
          },
          {
            "name": "green_routed_no",
            "type": "`$INTEGER`",
          },
          {
            "name": "h1_fallback_template",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "house_transport_doc_ref",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "format": "date-time",
            "name": "issue_date",
            "type": "`$STRING`",
          },
          {
            "name": "mapping",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "orange_routed_no",
            "type": "`$INTEGER`",
          },
          {
            "name": "parsed_declarations_no",
            "type": "`$INTEGER`",
          },
          {
            "name": "parser",
            "short": "* `aes_platform` - aes_platform * `cds_platform` - cds_platform * `cds_export_platform` - cds_export_platform * `g4_g3` - g4_g3 * `nhd_platform` - nhd_platform * `platform` - platform * `birds` - birds * `ics2_platform` - ics2_platform",
            "type": "`$STRING`",
          },
          {
            "format": "date-time",
            "name": "parsing_completed_at",
            "type": "`$STRING`",
          },
          {
            "format": "date-time",
            "name": "parsing_started_at",
            "type": "`$STRING`",
          },
          {
            "name": "passive_transport_nationality",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "passive_transport_number",
            "type": "`$STRING`",
          },
          {
            "name": "processed_declarations_no",
            "type": "`$INTEGER`",
          },
          {
            "format": "date-time",
            "name": "processing_ended_at",
            "type": "`$STRING`",
          },
          {
            "format": "date-time",
            "name": "processing_started_at",
            "type": "`$STRING`",
          },
          {
            "format": "date-time",
            "name": "receipt_generating_started_at",
            "type": "`$STRING`",
          },
          {
            "format": "date-time",
            "name": "receipt_request_started_at",
            "type": "`$STRING`",
          },
          {
            "name": "receipt_request_status",
            "short": "* `pending` - pending * `processing` - processing * `generating` - generating * `completed` - completed",
            "type": "`$STRING`",
          },
          {
            "name": "receipt_request_user",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "receipts_zip",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "red_routed_no",
            "type": "`$INTEGER`",
          },
          {
            "name": "rejected_status_no",
            "type": "`$INTEGER`",
          },
          {
            "name": "status",
            "short": "* `pending` - pending * `parsing` - parsing * `processing` - processing * `complete` - complete",
            "type": "`$STRING`",
          },
          {
            "name": "template",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "format": "date-time",
            "name": "updated_at",
            "readOnly": True,
            "req": True,
            "type": "`$STRING`",
          },
          {
            "name": "yellow_routed_no",
            "type": "`$INTEGER`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "bulk_upload",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "args": {},
                "kind": "http",
                "method": "POST",
                "orig": "/bulk-uploads",
                "segments": [
                  {
                    "lit": "bulk-uploads",
                  },
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "bulk-uploads",
                ],
              },
            ],
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "query": [
                    {
                      "kind": "query",
                      "name": "cursor",
                      "orig": "cursor",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/bulk-uploads",
                "segments": [
                  {
                    "lit": "bulk-uploads",
                  },
                ],
                "select": {
                  "exist": [
                    "cursor",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.results`",
                },
                "parts": [
                  "bulk-uploads",
                ],
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/bulk-uploads/{id}",
                "segments": [
                  {
                    "lit": "bulk-uploads",
                  },
                  {
                    "var": "id",
                  },
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
                "parts": [
                  "bulk-uploads",
                  "{id}",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/bulk-uploads/{id}/generate-pdfs",
                "segments": [
                  {
                    "lit": "bulk-uploads",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "generate-pdfs",
                  },
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
                "parts": [
                  "bulk-uploads",
                  "{id}",
                  "generate-pdfs",
                ],
              },
            ],
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "DELETE",
                "orig": "/bulk-uploads/{id}",
                "segments": [
                  {
                    "lit": "bulk-uploads",
                  },
                  {
                    "var": "id",
                  },
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
                "parts": [
                  "bulk-uploads",
                  "{id}",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "DELETE",
                "orig": "/bulk-uploads/{id}/generate-pdfs",
                "segments": [
                  {
                    "lit": "bulk-uploads",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "generate-pdfs",
                  },
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
                "parts": [
                  "bulk-uploads",
                  "{id}",
                  "generate-pdfs",
                ],
              },
            ],
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "PATCH",
                "orig": "/bulk-uploads/{id}",
                "segments": [
                  {
                    "lit": "bulk-uploads",
                  },
                  {
                    "var": "id",
                  },
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
                "parts": [
                  "bulk-uploads",
                  "{id}",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "PATCH",
                "orig": "/bulk-uploads/{id}/generate-pdfs",
                "segments": [
                  {
                    "lit": "bulk-uploads",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "generate-pdfs",
                  },
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
                "parts": [
                  "bulk-uploads",
                  "{id}",
                  "generate-pdfs",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "file": {
        "fields": [
          {
            "name": "company",
            "req": True,
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "format": "date-time",
            "name": "created_at",
            "readOnly": True,
            "req": True,
            "type": "`$STRING`",
          },
          {
            "name": "extension",
            "readOnly": True,
            "req": True,
            "type": "`$STRING`",
          },
          {
            "format": "uri",
            "name": "file",
            "req": True,
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "name",
            "readOnly": True,
            "req": True,
            "type": "`$STRING`",
          },
          {
            "name": "public",
            "op": {
              "create": {
                "type": "`$BOOLEAN`",
              },
            },
            "req": True,
            "type": "`$BOOLEAN`",
          },
          {
            "name": "size",
            "type": "`$INTEGER`",
          },
          {
            "format": "date-time",
            "name": "updated_at",
            "readOnly": True,
            "req": True,
            "type": "`$STRING`",
          },
          {
            "name": "url",
            "readOnly": True,
            "req": True,
            "type": "`$STRING`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "file",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "args": {},
                "kind": "http",
                "method": "POST",
                "orig": "/files",
                "segments": [
                  {
                    "lit": "files",
                  },
                ],
                "select": {},
                "transform": {
                  "req": {
                    "file": "`reqdata`",
                  },
                  "res": "`body`",
                },
                "parts": [
                  "files",
                ],
              },
            ],
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
            "name": "additional_declaration_type",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "address",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "authorisation",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "bank_details",
            "type": "`$STRING`",
          },
          {
            "name": "certificate",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "certificate_type",
            "readOnly": True,
            "req": True,
            "type": "`$STRING`",
          },
          {
            "name": "company",
            "req": True,
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "format": "date-time",
            "name": "created_at",
            "readOnly": True,
            "req": True,
            "type": "`$STRING`",
          },
          {
            "name": "customs_office_of_lodgement",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "format": "date-time",
            "name": "deleted_at",
            "type": "`$STRING`",
          },
          {
            "format": "email",
            "name": "email",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "identification_number",
            "type": "`$STRING`",
          },
          {
            "name": "indirect_representative",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "name",
            "type": "`$STRING`",
          },
          {
            "name": "nhd_last_submission_year",
            "type": "`$INTEGER`",
          },
          {
            "name": "nhd_submission_counter",
            "type": "`$INTEGER`",
          },
          {
            "name": "person_paying_customs_duty",
            "type": "`$STRING`",
          },
          {
            "name": "phone_country_code",
            "type": "`$STRING`",
          },
          {
            "name": "phone_number",
            "type": "`$STRING`",
          },
          {
            "name": "preferred_payment_method",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "signed_form",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "type",
            "short": "* `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `co…",
            "type": "`$STRING`",
          },
          {
            "name": "type_of_person",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "unlocode",
            "type": "`$STRING`",
          },
          {
            "format": "date-time",
            "name": "updated_at",
            "readOnly": True,
            "req": True,
            "type": "`$STRING`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "party",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "args": {},
                "kind": "http",
                "method": "POST",
                "orig": "/parties",
                "segments": [
                  {
                    "lit": "parties",
                  },
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "parties",
                ],
              },
            ],
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "query": [
                    {
                      "kind": "query",
                      "name": "cursor",
                      "orig": "cursor",
                      "type": "`$STRING`",
                    },
                    {
                      "kind": "query",
                      "name": "type",
                      "orig": "type",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/parties",
                "segments": [
                  {
                    "lit": "parties",
                  },
                ],
                "select": {
                  "exist": [
                    "cursor",
                    "type",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.results`",
                },
                "parts": [
                  "parties",
                ],
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/parties/{id}",
                "segments": [
                  {
                    "lit": "parties",
                  },
                  {
                    "var": "id",
                  },
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
                "parts": [
                  "parties",
                  "{id}",
                ],
              },
            ],
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "DELETE",
                "orig": "/parties/{id}",
                "segments": [
                  {
                    "lit": "parties",
                  },
                  {
                    "var": "id",
                  },
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
                "parts": [
                  "parties",
                  "{id}",
                ],
              },
            ],
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "PATCH",
                "orig": "/parties/{id}",
                "segments": [
                  {
                    "lit": "parties",
                  },
                  {
                    "var": "id",
                  },
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
                "parts": [
                  "parties",
                  "{id}",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "submission": {
        "fields": [
          {
            "name": "additional_external_ids",
            "type": "`$ARRAY`",
          },
          {
            "name": "amendment_reason",
            "type": "`$STRING`",
          },
          {
            "name": "amendment_status",
            "short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
            "type": "`$STRING`",
          },
          {
            "name": "answers",
            "op": {
              "update": {
                "type": "`$ARRAY`",
              },
            },
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "bypass_restricted_code",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "clearance_slip",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "client",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "company_member",
            "op": {
              "update": {
                "type": "`$STRING`",
              },
            },
            "req": True,
            "short": "cuid-format identifier for this entity.",
            "type": "`$OBJECT`",
          },
          {
            "name": "consignee",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "consignor",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "format": "date-time",
            "name": "created_at",
            "readOnly": True,
            "req": True,
            "type": "`$STRING`",
          },
          {
            "name": "declarant",
            "req": True,
            "short": "cuid-format identifier for this entity.",
            "type": "`$OBJECT`",
          },
          {
            "name": "document_upload_status",
            "short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
            "type": "`$STRING`",
          },
          {
            "name": "documents_presentation_requested",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "documents_upload_requested",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "external_id",
            "type": "`$STRING`",
          },
          {
            "name": "form",
            "readOnly": True,
            "req": True,
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "goods_presentation_status",
            "short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested",
            "type": "`$STRING`",
          },
          {
            "name": "hrcm_status",
            "short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "invalidation_status",
            "short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled",
            "type": "`$STRING`",
          },
          {
            "name": "is_global_template",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "latest_notification_item",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "latest_state",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "lrn",
            "type": "`$STRING`",
          },
          {
            "name": "mrn",
            "type": "`$STRING`",
          },
          {
            "name": "name",
            "type": "`$STRING`",
          },
          {
            "name": "partial_answers",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "receipt",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "refund_application_status",
            "short": "* `processing` - processing * `rejected` - rejected * `accepted` - accepted",
            "type": "`$STRING`",
          },
          {
            "name": "route",
            "short": "* `green` - green * `orange` - orange * `red` - red * `yellow` - yellow",
            "type": "`$STRING`",
          },
          {
            "name": "shipment_items_no",
            "type": "`$INTEGER`",
          },
          {
            "name": "shipment_items_quantity_no",
            "type": "`$INTEGER`",
          },
          {
            "name": "source",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "source_type",
            "short": "* `template` - template * `automated_import` - automated_import",
            "type": "`$STRING`",
          },
          {
            "name": "status",
            "short": "* `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re…",
            "type": "`$STRING`",
          },
          {
            "name": "template",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "template_id",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "template_properties",
            "type": "`$ARRAY`",
          },
          {
            "format": "decimal",
            "name": "total_tax_amount",
            "type": "`$STRING`",
          },
          {
            "format": "date-time",
            "name": "updated_at",
            "readOnly": True,
            "req": True,
            "type": "`$STRING`",
          },
          {
            "name": "verification_errors",
            "type": "`$ARRAY`",
          },
          {
            "name": "verification_status",
            "short": "* `passed` - passed * `failed` - failed",
            "type": "`$STRING`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "submission",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "POST",
                "orig": "/submissions/{id}/refund",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "refund",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                  "refund",
                ],
              },
              {
                "args": {},
                "kind": "http",
                "method": "POST",
                "orig": "/submissions",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "submissions",
                ],
              },
              {
                "args": {},
                "kind": "http",
                "method": "POST",
                "orig": "/submissions/retrieve",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "lit": "retrieve",
                  },
                ],
                "select": {
                  "$action": "retrieve",
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "submissions",
                  "retrieve",
                ],
              },
            ],
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {
                  "query": [
                    {
                      "kind": "query",
                      "name": "cursor",
                      "orig": "cursor",
                      "type": "`$STRING`",
                    },
                    {
                      "kind": "query",
                      "name": "form_subtype",
                      "orig": "form_subtype",
                      "type": "`$STRING`",
                    },
                    {
                      "example": "green",
                      "kind": "query",
                      "name": "route",
                      "orig": "route",
                      "type": "`$STRING`",
                    },
                    {
                      "example": "released",
                      "kind": "query",
                      "name": "status",
                      "orig": "status",
                      "type": "`$STRING`",
                    },
                    {
                      "kind": "query",
                      "name": "template",
                      "orig": "template",
                      "type": "`$BOOLEAN`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/submissions",
                "segments": [
                  {
                    "lit": "submissions",
                  },
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
                  "res": "`body.results`",
                },
                "parts": [
                  "submissions",
                ],
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/submissions/{id}",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/submissions/{id}/clearance-slip",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "clearance-slip",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                  "clearance-slip",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/submissions/{id}/notification-read",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "notification-read",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                  "notification-read",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/submissions/{id}/pbn-applicable",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "pbn-applicable",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                  "pbn-applicable",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/submissions/{id}/receipt",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "receipt",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                  "receipt",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/submissions/{id}/refund",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "refund",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                  "refund",
                ],
              },
              {
                "args": {},
                "kind": "http",
                "method": "GET",
                "orig": "/submissions/retrieve",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "lit": "retrieve",
                  },
                ],
                "select": {
                  "$action": "retrieve",
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "submissions",
                  "retrieve",
                ],
              },
            ],
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "DELETE",
                "orig": "/submissions/{id}",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "DELETE",
                "orig": "/submissions/{id}/clearance-slip",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "clearance-slip",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                  "clearance-slip",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "DELETE",
                "orig": "/submissions/{id}/notification-read",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "notification-read",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                  "notification-read",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "DELETE",
                "orig": "/submissions/{id}/pbn-applicable",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "pbn-applicable",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                  "pbn-applicable",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "DELETE",
                "orig": "/submissions/{id}/receipt",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "receipt",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                  "receipt",
                ],
              },
            ],
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "PATCH",
                "orig": "/submissions/{id}",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "PATCH",
                "orig": "/submissions/{id}/clearance-slip",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "clearance-slip",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                  "clearance-slip",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "PATCH",
                "orig": "/submissions/{id}/notification-read",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "notification-read",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                  "notification-read",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "PATCH",
                "orig": "/submissions/{id}/pbn-applicable",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "pbn-applicable",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                  "pbn-applicable",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "kind": "param",
                      "name": "id",
                      "orig": "id",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "PATCH",
                "orig": "/submissions/{id}/receipt",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                  {
                    "var": "id",
                  },
                  {
                    "lit": "receipt",
                  },
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
                "parts": [
                  "submissions",
                  "{id}",
                  "receipt",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "submission_detail": {
        "fields": [
          {
            "name": "additional_external_ids",
            "type": "`$ARRAY`",
          },
          {
            "name": "additional_information",
            "req": True,
            "type": "`$ARRAY`",
          },
          {
            "name": "amendment_status",
            "short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
            "type": "`$STRING`",
          },
          {
            "name": "clearance_slip",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "client",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "company",
            "req": True,
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "company_member",
            "req": True,
            "short": "cuid-format identifier for this entity.",
            "type": "`$OBJECT`",
          },
          {
            "name": "consignee",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "consignor",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "format": "date-time",
            "name": "created_at",
            "readOnly": True,
            "req": True,
            "type": "`$STRING`",
          },
          {
            "name": "declarant",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "document_upload_status",
            "short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
            "type": "`$STRING`",
          },
          {
            "name": "documents_presentation_requested",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "documents_upload_requested",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "external_id",
            "type": "`$STRING`",
          },
          {
            "name": "form",
            "readOnly": True,
            "req": True,
            "type": "`$STRING`",
          },
          {
            "name": "goods_presentation_status",
            "short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested",
            "type": "`$STRING`",
          },
          {
            "name": "hrcm_status",
            "short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "invalidation_status",
            "short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled",
            "type": "`$STRING`",
          },
          {
            "name": "is_global_template",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "latest_notification_item",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "latest_state",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "lrn",
            "type": "`$STRING`",
          },
          {
            "name": "mrn",
            "type": "`$STRING`",
          },
          {
            "name": "name",
            "type": "`$STRING`",
          },
          {
            "name": "receipt",
            "req": True,
            "type": "`$OBJECT`",
          },
          {
            "name": "refund_application_status",
            "short": "* `processing` - processing * `rejected` - rejected * `accepted` - accepted",
            "type": "`$STRING`",
          },
          {
            "name": "route",
            "short": "* `green` - green * `orange` - orange * `red` - red * `yellow` - yellow",
            "type": "`$STRING`",
          },
          {
            "name": "shipment_items_no",
            "type": "`$INTEGER`",
          },
          {
            "name": "shipment_items_quantity_no",
            "type": "`$INTEGER`",
          },
          {
            "name": "source",
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "source_type",
            "short": "* `template` - template * `automated_import` - automated_import",
            "type": "`$STRING`",
          },
          {
            "name": "status",
            "short": "* `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re…",
            "type": "`$STRING`",
          },
          {
            "name": "submission",
            "req": True,
            "short": "cuid-format identifier for this entity.",
            "type": "`$STRING`",
          },
          {
            "name": "supporting_documents",
            "type": "`$ARRAY`",
          },
          {
            "name": "template",
            "type": "`$BOOLEAN`",
          },
          {
            "format": "decimal",
            "name": "total_tax_amount",
            "type": "`$STRING`",
          },
          {
            "format": "date-time",
            "name": "updated_at",
            "readOnly": True,
            "req": True,
            "type": "`$STRING`",
          },
          {
            "name": "verification_errors",
            "type": "`$ARRAY`",
          },
          {
            "name": "verification_status",
            "short": "* `passed` - passed * `failed` - failed",
            "type": "`$STRING`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "submission_detail",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "args": {},
                "kind": "http",
                "method": "POST",
                "orig": "/documents-request",
                "segments": [
                  {
                    "lit": "documents-request",
                  },
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "documents-request",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }
