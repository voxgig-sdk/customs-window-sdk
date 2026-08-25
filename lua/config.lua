-- CustomsWindow SDK configuration

-- Build a fresh, fully materialised config table. Every call rebuilds the
-- whole structure, so prefer require("config_shared") unless you need a
-- private copy you intend to mutate.
local function make_config()
  return {
    main = {
      name = "CustomsWindow",
      slug = "customs-window",
      version = "0.1.1",
      target = "lua",
    },
    feature = {
      ["test"] = {
        ["options"] = {
          ["active"] = false,
        },
        ["transport"] = "base",
      },
    },
    options = {
      base = "https://api.customswindow.com",
      auth = {
        prefix = "",
      },
      headers = {
        ["content-type"] = "application/json",
      },
      entity = {
        ["bulk_upload"] = {},
        ["file"] = {},
        ["paginated_bulk_upload_list_list"] = {},
        ["paginated_party_list_list"] = {},
        ["paginated_submission_list_list"] = {},
        ["party"] = {},
        ["submission"] = {},
        ["submission_detail"] = {},
      },
    },
    entity = {
      ["bulk_upload"] = {
        ["fields"] = {
          {
            ["name"] = "active_transport_nationality",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "active_transport_number",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "arrival_datetime",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "client",
            ["req"] = true,
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "company_member",
            ["req"] = true,
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "created_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "declarant",
            ["req"] = true,
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "declarations_no",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "deleted_at",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "departure_datetime",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "errors_file",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "external_id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "failed_declarations_no",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "file",
            ["req"] = true,
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "green_routed_no",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "h1_fallback_template",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "house_transport_doc_ref",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "issue_date",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "mapping",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "orange_routed_no",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "parsed_declarations_no",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "parser",
            ["short"] = "* `aes_platform` - aes_platform * `cds_platform` - cds_platform * `cds_export_platform` - cds_export_platform * `g4_g3` - g4_g3 * `nhd_platform` - nhd_platform * `platform` - platform * `birds` - birds * `ics2_platform` - ics2_platform",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "parsing_completed_at",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "parsing_started_at",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "passive_transport_nationality",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "passive_transport_number",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "processed_declarations_no",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "processing_ended_at",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "processing_started_at",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "receipt_generating_started_at",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "receipt_request_started_at",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "receipt_request_status",
            ["short"] = "* `pending` - pending * `processing` - processing * `generating` - generating * `completed` - completed",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "receipt_request_user",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "receipts_zip",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "red_routed_no",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "rejected_status_no",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "status",
            ["short"] = "* `pending` - pending * `parsing` - parsing * `processing` - processing * `complete` - complete",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "template",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "yellow_routed_no",
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "bulk_upload",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/bulk-uploads",
                ["parts"] = {
                  "bulk-uploads",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "cursor",
                      ["orig"] = "cursor",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/bulk-uploads",
                ["parts"] = {
                  "bulk-uploads",
                },
                ["select"] = {
                  ["exist"] = {
                    "cursor",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.results`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/bulk-uploads/{id}",
                ["parts"] = {
                  "bulk-uploads",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/bulk-uploads/{id}/generate-pdfs",
                ["parts"] = {
                  "bulk-uploads",
                  "{id}",
                  "generate-pdfs",
                },
                ["select"] = {
                  ["$action"] = "generate_pdf",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/bulk-uploads/{id}",
                ["parts"] = {
                  "bulk-uploads",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/bulk-uploads/{id}/generate-pdfs",
                ["parts"] = {
                  "bulk-uploads",
                  "{id}",
                  "generate-pdfs",
                },
                ["select"] = {
                  ["$action"] = "generate_pdf",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PATCH",
                ["orig"] = "/bulk-uploads/{id}",
                ["parts"] = {
                  "bulk-uploads",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PATCH",
                ["orig"] = "/bulk-uploads/{id}/generate-pdfs",
                ["parts"] = {
                  "bulk-uploads",
                  "{id}",
                  "generate-pdfs",
                },
                ["select"] = {
                  ["$action"] = "generate_pdf",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["file"] = {
        ["fields"] = {
          {
            ["name"] = "company",
            ["req"] = true,
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "created_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "extension",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "file",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "name",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "public",
            ["op"] = {
              ["create"] = {
                ["type"] = "`$BOOLEAN`",
              },
            },
            ["req"] = true,
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "size",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "updated_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "url",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "file",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/files",
                ["parts"] = {
                  "files",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = {
                    ["file"] = "`reqdata`",
                  },
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["paginated_bulk_upload_list_list"] = {
        ["fields"] = {},
        ["name"] = "paginated_bulk_upload_list_list",
        ["op"] = {},
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["paginated_party_list_list"] = {
        ["fields"] = {},
        ["name"] = "paginated_party_list_list",
        ["op"] = {},
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["paginated_submission_list_list"] = {
        ["fields"] = {},
        ["name"] = "paginated_submission_list_list",
        ["op"] = {},
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["party"] = {
        ["fields"] = {
          {
            ["name"] = "additional_declaration_type",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "address",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "authorisation",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "bank_details",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "certificate",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "certificate_type",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "company",
            ["req"] = true,
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "created_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "customs_office_of_lodgement",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "deleted_at",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "email",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "identification_number",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "indirect_representative",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "name",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "nhd_last_submission_year",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "nhd_submission_counter",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "person_paying_customs_duty",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "phone_country_code",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "phone_number",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "preferred_payment_method",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "signed_form",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "type",
            ["short"] = "* `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `co…",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "type_of_person",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "unlocode",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "party",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/parties",
                ["parts"] = {
                  "parties",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "cursor",
                      ["orig"] = "cursor",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "type",
                      ["orig"] = "type",
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/parties",
                ["parts"] = {
                  "parties",
                },
                ["select"] = {
                  ["exist"] = {
                    "cursor",
                    "type",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.results`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/parties/{id}",
                ["parts"] = {
                  "parties",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/parties/{id}",
                ["parts"] = {
                  "parties",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PATCH",
                ["orig"] = "/parties/{id}",
                ["parts"] = {
                  "parties",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["submission"] = {
        ["fields"] = {
          {
            ["name"] = "additional_external_ids",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "amendment_reason",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "amendment_status",
            ["short"] = "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "answers",
            ["op"] = {
              ["update"] = {
                ["type"] = "`$ARRAY`",
              },
            },
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "bypass_restricted_code",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "clearance_slip",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "client",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "company_member",
            ["op"] = {
              ["update"] = {
                ["type"] = "`$STRING`",
              },
            },
            ["req"] = true,
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "consignee",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "consignor",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "created_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "declarant",
            ["req"] = true,
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "document_upload_status",
            ["short"] = "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "documents_presentation_requested",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "documents_upload_requested",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "external_id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "form",
            ["req"] = true,
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "goods_presentation_status",
            ["short"] = "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "hrcm_status",
            ["short"] = "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "invalidation_status",
            ["short"] = "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "is_global_template",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "latest_notification_item",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "latest_state",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "lrn",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "mrn",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "name",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "partial_answers",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "receipt",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "refund_application_status",
            ["short"] = "* `processing` - processing * `rejected` - rejected * `accepted` - accepted",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "route",
            ["short"] = "* `green` - green * `orange` - orange * `red` - red * `yellow` - yellow",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "shipment_items_no",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "shipment_items_quantity_no",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "source",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_type",
            ["short"] = "* `template` - template * `automated_import` - automated_import",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "status",
            ["short"] = "* `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re…",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "template",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "template_id",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "template_properties",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "total_tax_amount",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "verification_errors",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "verification_status",
            ["short"] = "* `passed` - passed * `failed` - failed",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "submission",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/submissions/{id}/refund",
                ["parts"] = {
                  "submissions",
                  "{id}",
                  "refund",
                },
                ["select"] = {
                  ["$action"] = "refund",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/submissions",
                ["parts"] = {
                  "submissions",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/submissions/retrieve",
                ["parts"] = {
                  "submissions",
                  "retrieve",
                },
                ["select"] = {
                  ["$action"] = "retrieve",
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {
                  ["query"] = {
                    {
                      ["kind"] = "query",
                      ["name"] = "cursor",
                      ["orig"] = "cursor",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "form_subtype",
                      ["orig"] = "form_subtype",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = "green",
                      ["kind"] = "query",
                      ["name"] = "route",
                      ["orig"] = "route",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = "released",
                      ["kind"] = "query",
                      ["name"] = "status",
                      ["orig"] = "status",
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["kind"] = "query",
                      ["name"] = "template",
                      ["orig"] = "template",
                      ["type"] = "`$BOOLEAN`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/submissions",
                ["parts"] = {
                  "submissions",
                },
                ["select"] = {
                  ["exist"] = {
                    "cursor",
                    "form_subtype",
                    "route",
                    "status",
                    "template",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body.results`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/submissions/{id}",
                ["parts"] = {
                  "submissions",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/submissions/{id}/clearance-slip",
                ["parts"] = {
                  "submissions",
                  "{id}",
                  "clearance-slip",
                },
                ["select"] = {
                  ["$action"] = "clearance_slip",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/submissions/{id}/notification-read",
                ["parts"] = {
                  "submissions",
                  "{id}",
                  "notification-read",
                },
                ["select"] = {
                  ["$action"] = "notification_read",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/submissions/{id}/pbn-applicable",
                ["parts"] = {
                  "submissions",
                  "{id}",
                  "pbn-applicable",
                },
                ["select"] = {
                  ["$action"] = "pbn_applicable",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/submissions/{id}/receipt",
                ["parts"] = {
                  "submissions",
                  "{id}",
                  "receipt",
                },
                ["select"] = {
                  ["$action"] = "receipt",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/submissions/{id}/refund",
                ["parts"] = {
                  "submissions",
                  "{id}",
                  "refund",
                },
                ["select"] = {
                  ["$action"] = "refund",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/submissions/retrieve",
                ["parts"] = {
                  "submissions",
                  "retrieve",
                },
                ["select"] = {
                  ["$action"] = "retrieve",
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/submissions/{id}",
                ["parts"] = {
                  "submissions",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/submissions/{id}/clearance-slip",
                ["parts"] = {
                  "submissions",
                  "{id}",
                  "clearance-slip",
                },
                ["select"] = {
                  ["$action"] = "clearance_slip",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/submissions/{id}/notification-read",
                ["parts"] = {
                  "submissions",
                  "{id}",
                  "notification-read",
                },
                ["select"] = {
                  ["$action"] = "notification_read",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/submissions/{id}/pbn-applicable",
                ["parts"] = {
                  "submissions",
                  "{id}",
                  "pbn-applicable",
                },
                ["select"] = {
                  ["$action"] = "pbn_applicable",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "DELETE",
                ["orig"] = "/submissions/{id}/receipt",
                ["parts"] = {
                  "submissions",
                  "{id}",
                  "receipt",
                },
                ["select"] = {
                  ["$action"] = "receipt",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PATCH",
                ["orig"] = "/submissions/{id}",
                ["parts"] = {
                  "submissions",
                  "{id}",
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PATCH",
                ["orig"] = "/submissions/{id}/clearance-slip",
                ["parts"] = {
                  "submissions",
                  "{id}",
                  "clearance-slip",
                },
                ["select"] = {
                  ["$action"] = "clearance_slip",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PATCH",
                ["orig"] = "/submissions/{id}/notification-read",
                ["parts"] = {
                  "submissions",
                  "{id}",
                  "notification-read",
                },
                ["select"] = {
                  ["$action"] = "notification_read",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PATCH",
                ["orig"] = "/submissions/{id}/pbn-applicable",
                ["parts"] = {
                  "submissions",
                  "{id}",
                  "pbn-applicable",
                },
                ["select"] = {
                  ["$action"] = "pbn_applicable",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "PATCH",
                ["orig"] = "/submissions/{id}/receipt",
                ["parts"] = {
                  "submissions",
                  "{id}",
                  "receipt",
                },
                ["select"] = {
                  ["$action"] = "receipt",
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["submission_detail"] = {
        ["fields"] = {
          {
            ["name"] = "additional_external_ids",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "additional_information",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "amendment_status",
            ["short"] = "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "clearance_slip",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "client",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "company",
            ["req"] = true,
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "company_member",
            ["req"] = true,
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "consignee",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "consignor",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "created_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "declarant",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "document_upload_status",
            ["short"] = "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "documents_presentation_requested",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "documents_upload_requested",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "external_id",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "form",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "goods_presentation_status",
            ["short"] = "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "hrcm_status",
            ["short"] = "* `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "id",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "invalidation_status",
            ["short"] = "* `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "is_global_template",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "latest_notification_item",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "latest_state",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "lrn",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "mrn",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "name",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "receipt",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
          },
          {
            ["name"] = "refund_application_status",
            ["short"] = "* `processing` - processing * `rejected` - rejected * `accepted` - accepted",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "route",
            ["short"] = "* `green` - green * `orange` - orange * `red` - red * `yellow` - yellow",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "shipment_items_no",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "shipment_items_quantity_no",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "source",
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "source_type",
            ["short"] = "* `template` - template * `automated_import` - automated_import",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "status",
            ["short"] = "* `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re…",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "submission",
            ["req"] = true,
            ["short"] = "cuid-format identifier for this entity.",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "supporting_documents",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "template",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "total_tax_amount",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "updated_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "verification_errors",
            ["type"] = "`$ARRAY`",
          },
          {
            ["name"] = "verification_status",
            ["short"] = "* `passed` - passed * `failed` - failed",
            ["type"] = "`$STRING`",
          },
        },
        ["name"] = "submission_detail",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "POST",
                ["orig"] = "/documents-request",
                ["parts"] = {
                  "documents-request",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
    },
  }
end


local function make_feature(name)
  local features = require("features")
  local factory = features[name]
  if factory ~= nil then
    return factory()
  end
  return features.base()
end


-- Attach make_feature to the SDK class
local function setup_sdk(SDK)
  SDK._make_feature = make_feature
end


return make_config
