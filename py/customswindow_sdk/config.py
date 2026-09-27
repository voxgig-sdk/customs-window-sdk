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
            "debug": {
        "options": {
          "active": False,
          "max": 100,
          "redact": [
            "authorization",
            "cookie",
            "set-cookie",
            "api-key",
            "apikey",
            "x-api-key",
            "idempotency-key",
          ],
        },
        "optspec": {
          "now": "`$FUNCTION`",
          "onEntry": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "none",
      },
            "idempotency": {
        "options": {
          "active": False,
          "header": "Idempotency-Key",
          "methods": [
            "POST",
            "PUT",
            "PATCH",
            "DELETE",
          ],
          "ops": [
            "create",
            "update",
            "remove",
          ],
        },
        "optspec": {
          "keygen": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "none",
      },
            "metrics": {
        "options": {
          "active": False,
        },
        "optspec": {
          "now": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "none",
      },
            "paging": {
        "options": {
          "active": False,
          "afterVar": "after",
          "cursorParam": "cursor",
          "firstVar": "first",
          "limitParam": "limit",
          "pageParam": "page",
          "startPage": 1,
        },
        "optspec": {
          "limit": "`$NUMBER`",
          "ops": "`$LIST`",
        },
        "strict": False,
        "transport": "none",
      },
            "ratelimit": {
        "options": {
          "active": False,
          "burst": 5,
          "rate": 5,
        },
        "optspec": {
          "now": "`$FUNCTION`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "retry": {
        "options": {
          "active": False,
          "factor": 2,
          "maxDelay": 2000,
          "minDelay": 50,
          "retries": 2,
          "statuses": [
            408,
            425,
            429,
            500,
            502,
            503,
            504,
          ],
        },
        "optspec": {
          "jitter": "`$BOOLEAN`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "test": {
        "options": {
          "active": False,
        },
        "optspec": {
          "entity": "`$MAP`",
          "net": "`$MAP`",
        },
        "strict": False,
        "transport": "base",
      },
            "timeout": {
        "options": {
          "active": False,
          "ms": 30000,
        },
        "optspec": {
          "clearTimer": "`$FUNCTION`",
          "setTimer": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
        },
        "options": {
            "base": "https://api.customswindow.com",
            "auth": {
                "prefix": "",
                "name": "Basic Auth",
            },
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "bulk_upload": {},
                "file": {},
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
            "title": "Active Transport Nationality",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "active_transport_number",
            "title": "Active Transport Number",
            "type": "`$STRING`",
          },
          {
            "name": "arrival_datetime",
            "title": "Arrival Datetime",
            "type": "`$STRING`",
            "format": "date-time",
          },
          {
            "name": "client",
            "title": "Client",
            "type": "`$STRING`",
            "req": True,
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "company_member",
            "title": "Company Member",
            "type": "`$STRING`",
            "req": True,
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "created_at",
            "title": "Created At",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
            "format": "date-time",
          },
          {
            "name": "declarant",
            "title": "Declarant",
            "type": "`$STRING`",
            "req": True,
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "declarations_no",
            "title": "Declarations No",
            "type": "`$INTEGER`",
            "req": True,
          },
          {
            "name": "deleted_at",
            "title": "Deleted At",
            "type": "`$STRING`",
            "format": "date-time",
          },
          {
            "name": "departure_datetime",
            "title": "Departure Datetime",
            "type": "`$STRING`",
            "format": "date-time",
          },
          {
            "name": "errors_file",
            "title": "Errors File",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "external_id",
            "title": "External Id",
            "type": "`$STRING`",
          },
          {
            "name": "failed_declarations_no",
            "title": "Failed Declarations No",
            "type": "`$INTEGER`",
          },
          {
            "name": "file",
            "title": "File",
            "type": "`$OBJECT`",
            "req": True,
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "green_routed_no",
            "title": "Green Routed No",
            "type": "`$INTEGER`",
          },
          {
            "name": "h1_fallback_template",
            "title": "H1 Fallback Template",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "house_transport_doc_ref",
            "title": "House Transport Doc Ref",
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "issue_date",
            "title": "Issue Date",
            "type": "`$STRING`",
            "format": "date-time",
          },
          {
            "name": "mapping",
            "title": "Mapping",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "orange_routed_no",
            "title": "Orange Routed No",
            "type": "`$INTEGER`",
          },
          {
            "name": "parsed_declarations_no",
            "title": "Parsed Declarations No",
            "type": "`$INTEGER`",
          },
          {
            "name": "parser",
            "title": "Parser",
            "type": "`$STRING`",
            "short": "* `aes_platform` - aes_platform * `cds_platform` - cds_platform * `cds_export_platform` - cds_export_platform * `g4_g3` - g4_g3 * `nhd_platform` - nhd_platform * `platform` - platform * `birds` - birds * `ics2_platform` - ics2_platform",
          },
          {
            "name": "parsing_completed_at",
            "title": "Parsing Completed At",
            "type": "`$STRING`",
            "format": "date-time",
          },
          {
            "name": "parsing_started_at",
            "title": "Parsing Started At",
            "type": "`$STRING`",
            "format": "date-time",
          },
          {
            "name": "passive_transport_nationality",
            "title": "Passive Transport Nationality",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "passive_transport_number",
            "title": "Passive Transport Number",
            "type": "`$STRING`",
          },
          {
            "name": "processed_declarations_no",
            "title": "Processed Declarations No",
            "type": "`$INTEGER`",
          },
          {
            "name": "processing_ended_at",
            "title": "Processing Ended At",
            "type": "`$STRING`",
            "format": "date-time",
          },
          {
            "name": "processing_started_at",
            "title": "Processing Started At",
            "type": "`$STRING`",
            "format": "date-time",
          },
          {
            "name": "receipt_generating_started_at",
            "title": "Receipt Generating Started At",
            "type": "`$STRING`",
            "format": "date-time",
          },
          {
            "name": "receipt_request_started_at",
            "title": "Receipt Request Started At",
            "type": "`$STRING`",
            "format": "date-time",
          },
          {
            "name": "receipt_request_status",
            "title": "Receipt Request Status",
            "type": "`$STRING`",
            "short": "* `pending` - pending * `processing` - processing * `generating` - generating * `completed` - completed",
          },
          {
            "name": "receipt_request_user",
            "title": "Receipt Request User",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "receipts_zip",
            "title": "Receipts Zip",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "red_routed_no",
            "title": "Red Routed No",
            "type": "`$INTEGER`",
          },
          {
            "name": "rejected_status_no",
            "title": "Rejected Status No",
            "type": "`$INTEGER`",
          },
          {
            "name": "status",
            "title": "Status",
            "type": "`$STRING`",
            "short": "* `pending` - pending * `parsing` - parsing * `processing` - processing * `complete` - complete",
          },
          {
            "name": "template",
            "title": "Template",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "updated_at",
            "title": "Updated At",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
            "format": "date-time",
          },
          {
            "name": "yellow_routed_no",
            "title": "Yellow Routed No",
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
                "kind": "http",
                "method": "POST",
                "orig": "/bulk-uploads",
                "segments": [
                  {
                    "lit": "bulk-uploads",
                  },
                ],
                "parts": [
                  "bulk-uploads",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
            ],
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/bulk-uploads",
                "segments": [
                  {
                    "lit": "bulk-uploads",
                  },
                ],
                "parts": [
                  "bulk-uploads",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.results`",
                },
                "args": {
                  "query": [
                    {
                      "name": "cursor",
                      "orig": "cursor",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "cursor",
                  ],
                },
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
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
                "parts": [
                  "bulk-uploads",
                  "{id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "bulk-uploads",
                  "{id}",
                  "generate-pdfs",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "generate_pdf",
                  "exist": [
                    "id",
                  ],
                },
              },
            ],
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
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
                "parts": [
                  "bulk-uploads",
                  "{id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "bulk-uploads",
                  "{id}",
                  "generate-pdfs",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "generate_pdf",
                  "exist": [
                    "id",
                  ],
                },
              },
            ],
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
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
                "parts": [
                  "bulk-uploads",
                  "{id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "bulk-uploads",
                  "{id}",
                  "generate-pdfs",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "generate_pdf",
                  "exist": [
                    "id",
                  ],
                },
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
            "title": "Company",
            "type": "`$STRING`",
            "req": True,
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "created_at",
            "title": "Created At",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
            "format": "date-time",
          },
          {
            "name": "extension",
            "title": "Extension",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
          },
          {
            "name": "file",
            "title": "File",
            "type": "`$STRING`",
            "req": True,
            "format": "uri",
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "name",
            "title": "Name",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
          },
          {
            "name": "public",
            "title": "Public",
            "type": "`$BOOLEAN`",
            "req": True,
            "op": {
              "create": {
                "type": "`$BOOLEAN`",
              },
            },
          },
          {
            "name": "size",
            "title": "Size",
            "type": "`$INTEGER`",
          },
          {
            "name": "updated_at",
            "title": "Updated At",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
            "format": "date-time",
          },
          {
            "name": "url",
            "title": "Url",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
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
                "kind": "http",
                "method": "POST",
                "orig": "/files",
                "segments": [
                  {
                    "lit": "files",
                  },
                ],
                "parts": [
                  "files",
                ],
                "rename": {},
                "transform": {
                  "req": {
                    "file": "`reqdata`",
                  },
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "party": {
        "fields": [
          {
            "name": "additional_declaration_type",
            "title": "Additional Declaration Type",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "address",
            "title": "Address",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "authorisation",
            "title": "Authorisation",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "bank_details",
            "title": "Bank Details",
            "type": "`$STRING`",
          },
          {
            "name": "certificate",
            "title": "Certificate",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "certificate_type",
            "title": "Certificate Type",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
          },
          {
            "name": "company",
            "title": "Company",
            "type": "`$STRING`",
            "req": True,
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "created_at",
            "title": "Created At",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
            "format": "date-time",
          },
          {
            "name": "customs_office_of_lodgement",
            "title": "Customs Office Of Lodgement",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "deleted_at",
            "title": "Deleted At",
            "type": "`$STRING`",
            "format": "date-time",
          },
          {
            "name": "email",
            "title": "Email",
            "type": "`$STRING`",
            "format": "email",
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "identification_number",
            "title": "Identification Number",
            "type": "`$STRING`",
          },
          {
            "name": "indirect_representative",
            "title": "Indirect Representative",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "name",
            "title": "Name",
            "type": "`$STRING`",
          },
          {
            "name": "nhd_last_submission_year",
            "title": "Nhd Last Submission Year",
            "type": "`$INTEGER`",
          },
          {
            "name": "nhd_submission_counter",
            "title": "Nhd Submission Counter",
            "type": "`$INTEGER`",
          },
          {
            "name": "person_paying_customs_duty",
            "title": "Person Paying Customs Duty",
            "type": "`$STRING`",
          },
          {
            "name": "phone_country_code",
            "title": "Phone Country Code",
            "type": "`$STRING`",
          },
          {
            "name": "phone_number",
            "title": "Phone Number",
            "type": "`$STRING`",
          },
          {
            "name": "preferred_payment_method",
            "title": "Preferred Payment Method",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "signed_form",
            "title": "Signed Form",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "type",
            "title": "Type",
            "type": "`$STRING`",
            "short": "* `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `co…",
          },
          {
            "name": "type_of_person",
            "title": "Type Of Person",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "unlocode",
            "title": "Unlocode",
            "type": "`$STRING`",
          },
          {
            "name": "updated_at",
            "title": "Updated At",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
            "format": "date-time",
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
                "kind": "http",
                "method": "POST",
                "orig": "/parties",
                "segments": [
                  {
                    "lit": "parties",
                  },
                ],
                "parts": [
                  "parties",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
            ],
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/parties",
                "segments": [
                  {
                    "lit": "parties",
                  },
                ],
                "parts": [
                  "parties",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.results`",
                },
                "args": {
                  "query": [
                    {
                      "name": "cursor",
                      "orig": "cursor",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "type",
                      "orig": "type",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "cursor",
                    "type",
                  ],
                },
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
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
                "parts": [
                  "parties",
                  "{id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
              },
            ],
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
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
                "parts": [
                  "parties",
                  "{id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
              },
            ],
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
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
                "parts": [
                  "parties",
                  "{id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
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
            "title": "Additional External Ids",
            "type": "`$ARRAY`",
          },
          {
            "name": "amendment_reason",
            "title": "Amendment Reason",
            "type": "`$STRING`",
          },
          {
            "name": "amendment_status",
            "title": "Amendment Status",
            "type": "`$STRING`",
            "short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
          },
          {
            "name": "answers",
            "title": "Answers",
            "type": "`$ARRAY`",
            "req": True,
            "op": {
              "update": {
                "type": "`$ARRAY`",
              },
            },
          },
          {
            "name": "bypass_restricted_code",
            "title": "Bypass Restricted Code",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "clearance_slip",
            "title": "Clearance Slip",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "client",
            "title": "Client",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "company_member",
            "title": "Company Member",
            "type": "`$OBJECT`",
            "req": True,
            "op": {
              "update": {
                "type": "`$STRING`",
              },
            },
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "consignee",
            "title": "Consignee",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "consignor",
            "title": "Consignor",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "created_at",
            "title": "Created At",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
            "format": "date-time",
          },
          {
            "name": "declarant",
            "title": "Declarant",
            "type": "`$OBJECT`",
            "req": True,
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "document_upload_status",
            "title": "Document Upload Status",
            "type": "`$STRING`",
            "short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
          },
          {
            "name": "documents_presentation_requested",
            "title": "Documents Presentation Requested",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "documents_upload_requested",
            "title": "Documents Upload Requested",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "external_id",
            "title": "External Id",
            "type": "`$STRING`",
          },
          {
            "name": "form",
            "title": "Form",
            "type": "`$STRING`",
            "req": True,
            "short": "cuid-format identifier for this entity.",
            "readOnly": True,
          },
          {
            "name": "goods_presentation_status",
            "title": "Goods Presentation Status",
            "type": "`$STRING`",
            "short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested",
          },
          {
            "name": "hrcm_status",
            "title": "Hrcm Status",
            "type": "`$STRING`",
            "short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "invalidation_status",
            "title": "Invalidation Status",
            "type": "`$STRING`",
            "short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled",
          },
          {
            "name": "is_global_template",
            "title": "Is Global Template",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "latest_notification_item",
            "title": "Latest Notification Item",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "latest_state",
            "title": "Latest State",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "lrn",
            "title": "Lrn",
            "type": "`$STRING`",
          },
          {
            "name": "mrn",
            "title": "Mrn",
            "type": "`$STRING`",
          },
          {
            "name": "name",
            "title": "Name",
            "type": "`$STRING`",
          },
          {
            "name": "partial_answers",
            "title": "Partial Answers",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "receipt",
            "title": "Receipt",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "refund_application_status",
            "title": "Refund Application Status",
            "type": "`$STRING`",
            "short": "* `processing` - processing * `rejected` - rejected * `accepted` - accepted",
          },
          {
            "name": "route",
            "title": "Route",
            "type": "`$STRING`",
            "short": "* `green` - green * `orange` - orange * `red` - red * `yellow` - yellow",
          },
          {
            "name": "shipment_items_no",
            "title": "Shipment Items No",
            "type": "`$INTEGER`",
          },
          {
            "name": "shipment_items_quantity_no",
            "title": "Shipment Items Quantity No",
            "type": "`$INTEGER`",
          },
          {
            "name": "source",
            "title": "Source",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "source_type",
            "title": "Source Type",
            "type": "`$STRING`",
            "short": "* `template` - template * `automated_import` - automated_import",
          },
          {
            "name": "status",
            "title": "Status",
            "type": "`$STRING`",
            "short": "* `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re…",
          },
          {
            "name": "template",
            "title": "Template",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "template_id",
            "title": "Template Id",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "template_properties",
            "title": "Template Properties",
            "type": "`$ARRAY`",
          },
          {
            "name": "total_tax_amount",
            "title": "Total Tax Amount",
            "type": "`$STRING`",
            "format": "decimal",
          },
          {
            "name": "updated_at",
            "title": "Updated At",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
            "format": "date-time",
          },
          {
            "name": "verification_errors",
            "title": "Verification Errors",
            "type": "`$ARRAY`",
          },
          {
            "name": "verification_status",
            "title": "Verification Status",
            "type": "`$STRING`",
            "short": "* `passed` - passed * `failed` - failed",
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
                "parts": [
                  "submissions",
                  "{id}",
                  "refund",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "refund",
                  "exist": [
                    "id",
                  ],
                },
              },
              {
                "kind": "http",
                "method": "POST",
                "orig": "/submissions",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                ],
                "parts": [
                  "submissions",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
              },
              {
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
                "parts": [
                  "submissions",
                  "retrieve",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {
                  "$action": "retrieve",
                },
              },
            ],
          },
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "kind": "http",
                "method": "GET",
                "orig": "/submissions",
                "segments": [
                  {
                    "lit": "submissions",
                  },
                ],
                "parts": [
                  "submissions",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.results`",
                },
                "args": {
                  "query": [
                    {
                      "name": "cursor",
                      "orig": "cursor",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "form_subtype",
                      "orig": "form_subtype",
                      "type": "`$STRING`",
                      "kind": "query",
                    },
                    {
                      "name": "route",
                      "orig": "route",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "green",
                    },
                    {
                      "name": "status",
                      "orig": "status",
                      "type": "`$STRING`",
                      "kind": "query",
                      "example": "released",
                    },
                    {
                      "name": "template",
                      "orig": "template",
                      "type": "`$BOOLEAN`",
                      "kind": "query",
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "cursor",
                    "form_subtype",
                    "route",
                    "status",
                    "template",
                  ],
                },
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                  "clearance-slip",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "clearance_slip",
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                  "notification-read",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "notification_read",
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                  "pbn-applicable",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "pbn_applicable",
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                  "receipt",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "receipt",
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                  "refund",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "refund",
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "submissions",
                  "retrieve",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {
                  "$action": "retrieve",
                },
              },
            ],
          },
          "remove": {
            "input": "data",
            "name": "remove",
            "points": [
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                  "clearance-slip",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "clearance_slip",
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                  "notification-read",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "notification_read",
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                  "pbn-applicable",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "pbn_applicable",
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                  "receipt",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "receipt",
                  "exist": [
                    "id",
                  ],
                },
              },
            ],
          },
          "update": {
            "input": "data",
            "name": "update",
            "points": [
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                  "clearance-slip",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "clearance_slip",
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                  "notification-read",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "notification_read",
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                  "pbn-applicable",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "pbn_applicable",
                  "exist": [
                    "id",
                  ],
                },
              },
              {
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
                "parts": [
                  "submissions",
                  "{id}",
                  "receipt",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {
                  "params": [
                    {
                      "name": "id",
                      "orig": "id",
                      "type": "`$STRING`",
                      "kind": "param",
                      "reqd": True,
                    },
                  ],
                },
                "select": {
                  "$action": "receipt",
                  "exist": [
                    "id",
                  ],
                },
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
            "title": "Additional External Ids",
            "type": "`$ARRAY`",
          },
          {
            "name": "additional_information",
            "title": "Additional Information",
            "type": "`$ARRAY`",
            "req": True,
          },
          {
            "name": "amendment_status",
            "title": "Amendment Status",
            "type": "`$STRING`",
            "short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
          },
          {
            "name": "clearance_slip",
            "title": "Clearance Slip",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "client",
            "title": "Client",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "company",
            "title": "Company",
            "type": "`$STRING`",
            "req": True,
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "company_member",
            "title": "Company Member",
            "type": "`$OBJECT`",
            "req": True,
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "consignee",
            "title": "Consignee",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "consignor",
            "title": "Consignor",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "created_at",
            "title": "Created At",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
            "format": "date-time",
          },
          {
            "name": "declarant",
            "title": "Declarant",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "document_upload_status",
            "title": "Document Upload Status",
            "type": "`$STRING`",
            "short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
          },
          {
            "name": "documents_presentation_requested",
            "title": "Documents Presentation Requested",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "documents_upload_requested",
            "title": "Documents Upload Requested",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "external_id",
            "title": "External Id",
            "type": "`$STRING`",
          },
          {
            "name": "form",
            "title": "Form",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
          },
          {
            "name": "goods_presentation_status",
            "title": "Goods Presentation Status",
            "type": "`$STRING`",
            "short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested",
          },
          {
            "name": "hrcm_status",
            "title": "Hrcm Status",
            "type": "`$STRING`",
            "short": "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
          },
          {
            "name": "id",
            "title": "Id",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "invalidation_status",
            "title": "Invalidation Status",
            "type": "`$STRING`",
            "short": "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled",
          },
          {
            "name": "is_global_template",
            "title": "Is Global Template",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "latest_notification_item",
            "title": "Latest Notification Item",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "latest_state",
            "title": "Latest State",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "lrn",
            "title": "Lrn",
            "type": "`$STRING`",
          },
          {
            "name": "mrn",
            "title": "Mrn",
            "type": "`$STRING`",
          },
          {
            "name": "name",
            "title": "Name",
            "type": "`$STRING`",
          },
          {
            "name": "receipt",
            "title": "Receipt",
            "type": "`$OBJECT`",
            "req": True,
          },
          {
            "name": "refund_application_status",
            "title": "Refund Application Status",
            "type": "`$STRING`",
            "short": "* `processing` - processing * `rejected` - rejected * `accepted` - accepted",
          },
          {
            "name": "route",
            "title": "Route",
            "type": "`$STRING`",
            "short": "* `green` - green * `orange` - orange * `red` - red * `yellow` - yellow",
          },
          {
            "name": "shipment_items_no",
            "title": "Shipment Items No",
            "type": "`$INTEGER`",
          },
          {
            "name": "shipment_items_quantity_no",
            "title": "Shipment Items Quantity No",
            "type": "`$INTEGER`",
          },
          {
            "name": "source",
            "title": "Source",
            "type": "`$STRING`",
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "source_type",
            "title": "Source Type",
            "type": "`$STRING`",
            "short": "* `template` - template * `automated_import` - automated_import",
          },
          {
            "name": "status",
            "title": "Status",
            "type": "`$STRING`",
            "short": "* `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re…",
          },
          {
            "name": "submission",
            "title": "Submission",
            "type": "`$STRING`",
            "req": True,
            "short": "cuid-format identifier for this entity.",
          },
          {
            "name": "supporting_documents",
            "title": "Supporting Documents",
            "type": "`$ARRAY`",
          },
          {
            "name": "template",
            "title": "Template",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "total_tax_amount",
            "title": "Total Tax Amount",
            "type": "`$STRING`",
            "format": "decimal",
          },
          {
            "name": "updated_at",
            "title": "Updated At",
            "type": "`$STRING`",
            "req": True,
            "readOnly": True,
            "format": "date-time",
          },
          {
            "name": "verification_errors",
            "title": "Verification Errors",
            "type": "`$ARRAY`",
          },
          {
            "name": "verification_status",
            "title": "Verification Status",
            "type": "`$STRING`",
            "short": "* `passed` - passed * `failed` - failed",
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
                "kind": "http",
                "method": "POST",
                "orig": "/documents-request",
                "segments": [
                  {
                    "lit": "documents-request",
                  },
                ],
                "parts": [
                  "documents-request",
                ],
                "rename": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "args": {},
                "select": {},
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
