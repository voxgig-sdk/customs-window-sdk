-- CustomsWindow SDK configuration

local function make_config()
  return {
    main = {
      name = "CustomsWindow",
    },
    feature = {
      ["test"] = {
        ["options"] = {
          ["active"] = false,
        },
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
            ["active"] = true,
            ["name"] = "active_transport_nationality",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 0,
          },
          {
            ["active"] = true,
            ["name"] = "active_transport_number",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 1,
          },
          {
            ["active"] = true,
            ["name"] = "arrival_datetime",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 2,
          },
          {
            ["active"] = true,
            ["name"] = "client",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 3,
          },
          {
            ["active"] = true,
            ["name"] = "company_member",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 4,
          },
          {
            ["active"] = true,
            ["name"] = "created_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 5,
          },
          {
            ["active"] = true,
            ["name"] = "declarant",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 6,
          },
          {
            ["active"] = true,
            ["name"] = "declarations_no",
            ["req"] = true,
            ["type"] = "`$INTEGER`",
            ["index$"] = 7,
          },
          {
            ["active"] = true,
            ["name"] = "deleted_at",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 8,
          },
          {
            ["active"] = true,
            ["name"] = "departure_datetime",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 9,
          },
          {
            ["active"] = true,
            ["name"] = "errors_file",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 10,
          },
          {
            ["active"] = true,
            ["name"] = "external_id",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 11,
          },
          {
            ["active"] = true,
            ["name"] = "failed_declarations_no",
            ["req"] = false,
            ["type"] = "`$INTEGER`",
            ["index$"] = 12,
          },
          {
            ["active"] = true,
            ["name"] = "file",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 13,
          },
          {
            ["active"] = true,
            ["name"] = "green_routed_no",
            ["req"] = false,
            ["type"] = "`$INTEGER`",
            ["index$"] = 14,
          },
          {
            ["active"] = true,
            ["name"] = "h1_fallback_template",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 15,
          },
          {
            ["active"] = true,
            ["name"] = "house_transport_doc_ref",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 16,
          },
          {
            ["active"] = true,
            ["name"] = "id",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 17,
          },
          {
            ["active"] = true,
            ["name"] = "issue_date",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 18,
          },
          {
            ["active"] = true,
            ["name"] = "mapping",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 19,
          },
          {
            ["active"] = true,
            ["name"] = "orange_routed_no",
            ["req"] = false,
            ["type"] = "`$INTEGER`",
            ["index$"] = 20,
          },
          {
            ["active"] = true,
            ["name"] = "parsed_declarations_no",
            ["req"] = false,
            ["type"] = "`$INTEGER`",
            ["index$"] = 21,
          },
          {
            ["active"] = true,
            ["name"] = "parser",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 22,
          },
          {
            ["active"] = true,
            ["name"] = "parsing_completed_at",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 23,
          },
          {
            ["active"] = true,
            ["name"] = "parsing_started_at",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 24,
          },
          {
            ["active"] = true,
            ["name"] = "passive_transport_nationality",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 25,
          },
          {
            ["active"] = true,
            ["name"] = "passive_transport_number",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 26,
          },
          {
            ["active"] = true,
            ["name"] = "processed_declarations_no",
            ["req"] = false,
            ["type"] = "`$INTEGER`",
            ["index$"] = 27,
          },
          {
            ["active"] = true,
            ["name"] = "processing_ended_at",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 28,
          },
          {
            ["active"] = true,
            ["name"] = "processing_started_at",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 29,
          },
          {
            ["active"] = true,
            ["name"] = "receipt_generating_started_at",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 30,
          },
          {
            ["active"] = true,
            ["name"] = "receipt_request_started_at",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 31,
          },
          {
            ["active"] = true,
            ["name"] = "receipt_request_status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 32,
          },
          {
            ["active"] = true,
            ["name"] = "receipt_request_user",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 33,
          },
          {
            ["active"] = true,
            ["name"] = "receipts_zip",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 34,
          },
          {
            ["active"] = true,
            ["name"] = "red_routed_no",
            ["req"] = false,
            ["type"] = "`$INTEGER`",
            ["index$"] = 35,
          },
          {
            ["active"] = true,
            ["name"] = "rejected_status_no",
            ["req"] = false,
            ["type"] = "`$INTEGER`",
            ["index$"] = 36,
          },
          {
            ["active"] = true,
            ["name"] = "status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 37,
          },
          {
            ["active"] = true,
            ["name"] = "template",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 38,
          },
          {
            ["active"] = true,
            ["name"] = "updated_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 39,
          },
          {
            ["active"] = true,
            ["name"] = "yellow_routed_no",
            ["req"] = false,
            ["type"] = "`$INTEGER`",
            ["index$"] = 40,
          },
        },
        ["name"] = "bulk_upload",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {},
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
                ["index$"] = 0,
              },
            },
            ["key$"] = "create",
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {
                  ["query"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "query",
                      ["name"] = "cursor",
                      ["orig"] = "cursor",
                      ["reqd"] = false,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                  ["res"] = "`body`",
                },
                ["index$"] = 0,
              },
            },
            ["key$"] = "list",
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                      ["index$"] = 0,
                    },
                  },
                },
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
                ["index$"] = 0,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                ["index$"] = 1,
              },
            },
            ["key$"] = "load",
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                      ["index$"] = 0,
                    },
                  },
                },
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
                ["index$"] = 0,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                ["index$"] = 1,
              },
            },
            ["key$"] = "remove",
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                      ["index$"] = 0,
                    },
                  },
                },
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
                ["index$"] = 0,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                      ["index$"] = 0,
                    },
                  },
                },
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
                ["index$"] = 1,
              },
            },
            ["key$"] = "update",
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["file"] = {
        ["fields"] = {
          {
            ["active"] = true,
            ["name"] = "company",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 0,
          },
          {
            ["active"] = true,
            ["name"] = "created_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 1,
          },
          {
            ["active"] = true,
            ["name"] = "extension",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 2,
          },
          {
            ["active"] = true,
            ["name"] = "file",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 3,
          },
          {
            ["active"] = true,
            ["name"] = "id",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 4,
          },
          {
            ["active"] = true,
            ["name"] = "name",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 5,
          },
          {
            ["active"] = true,
            ["name"] = "public",
            ["op"] = {
              ["create"] = {
                ["req"] = false,
                ["type"] = "`$BOOLEAN`",
              },
            },
            ["req"] = true,
            ["type"] = "`$BOOLEAN`",
            ["index$"] = 6,
          },
          {
            ["active"] = true,
            ["name"] = "size",
            ["req"] = false,
            ["type"] = "`$INTEGER`",
            ["index$"] = 7,
          },
          {
            ["active"] = true,
            ["name"] = "updated_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 8,
          },
          {
            ["active"] = true,
            ["name"] = "url",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 9,
          },
        },
        ["name"] = "file",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {},
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
                ["index$"] = 0,
              },
            },
            ["key$"] = "create",
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
            ["active"] = true,
            ["name"] = "additional_declaration_type",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 0,
          },
          {
            ["active"] = true,
            ["name"] = "address",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 1,
          },
          {
            ["active"] = true,
            ["name"] = "authorisation",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 2,
          },
          {
            ["active"] = true,
            ["name"] = "bank_detail",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 3,
          },
          {
            ["active"] = true,
            ["name"] = "certificate",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 4,
          },
          {
            ["active"] = true,
            ["name"] = "certificate_type",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 5,
          },
          {
            ["active"] = true,
            ["name"] = "company",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 6,
          },
          {
            ["active"] = true,
            ["name"] = "created_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 7,
          },
          {
            ["active"] = true,
            ["name"] = "customs_office_of_lodgement",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 8,
          },
          {
            ["active"] = true,
            ["name"] = "deleted_at",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 9,
          },
          {
            ["active"] = true,
            ["name"] = "email",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 10,
          },
          {
            ["active"] = true,
            ["name"] = "id",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 11,
          },
          {
            ["active"] = true,
            ["name"] = "identification_number",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 12,
          },
          {
            ["active"] = true,
            ["name"] = "indirect_representative",
            ["req"] = false,
            ["type"] = "`$BOOLEAN`",
            ["index$"] = 13,
          },
          {
            ["active"] = true,
            ["name"] = "name",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 14,
          },
          {
            ["active"] = true,
            ["name"] = "nhd_last_submission_year",
            ["req"] = false,
            ["type"] = "`$INTEGER`",
            ["index$"] = 15,
          },
          {
            ["active"] = true,
            ["name"] = "nhd_submission_counter",
            ["req"] = false,
            ["type"] = "`$INTEGER`",
            ["index$"] = 16,
          },
          {
            ["active"] = true,
            ["name"] = "person_paying_customs_duty",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 17,
          },
          {
            ["active"] = true,
            ["name"] = "phone_country_code",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 18,
          },
          {
            ["active"] = true,
            ["name"] = "phone_number",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 19,
          },
          {
            ["active"] = true,
            ["name"] = "preferred_payment_method",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 20,
          },
          {
            ["active"] = true,
            ["name"] = "signed_form",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 21,
          },
          {
            ["active"] = true,
            ["name"] = "type",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 22,
          },
          {
            ["active"] = true,
            ["name"] = "type_of_person",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 23,
          },
          {
            ["active"] = true,
            ["name"] = "unlocode",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 24,
          },
          {
            ["active"] = true,
            ["name"] = "updated_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 25,
          },
        },
        ["name"] = "party",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {},
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
                ["index$"] = 0,
              },
            },
            ["key$"] = "create",
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {
                  ["query"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "query",
                      ["name"] = "cursor",
                      ["orig"] = "cursor",
                      ["reqd"] = false,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["active"] = true,
                      ["kind"] = "query",
                      ["name"] = "type",
                      ["orig"] = "type",
                      ["reqd"] = false,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                  ["res"] = "`body`",
                },
                ["index$"] = 0,
              },
            },
            ["key$"] = "list",
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                      ["index$"] = 0,
                    },
                  },
                },
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
                ["index$"] = 0,
              },
            },
            ["key$"] = "load",
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                      ["index$"] = 0,
                    },
                  },
                },
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
                ["index$"] = 0,
              },
            },
            ["key$"] = "remove",
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                      ["index$"] = 0,
                    },
                  },
                },
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
                ["index$"] = 0,
              },
            },
            ["key$"] = "update",
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["submission"] = {
        ["fields"] = {
          {
            ["active"] = true,
            ["name"] = "additional_external_id",
            ["req"] = false,
            ["type"] = "`$ARRAY`",
            ["index$"] = 0,
          },
          {
            ["active"] = true,
            ["name"] = "amendment_reason",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 1,
          },
          {
            ["active"] = true,
            ["name"] = "amendment_status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 2,
          },
          {
            ["active"] = true,
            ["name"] = "answer",
            ["op"] = {
              ["update"] = {
                ["req"] = false,
                ["type"] = "`$ARRAY`",
              },
            },
            ["req"] = true,
            ["type"] = "`$ARRAY`",
            ["index$"] = 3,
          },
          {
            ["active"] = true,
            ["name"] = "bypass_restricted_code",
            ["req"] = false,
            ["type"] = "`$BOOLEAN`",
            ["index$"] = 4,
          },
          {
            ["active"] = true,
            ["name"] = "clearance_slip",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 5,
          },
          {
            ["active"] = true,
            ["name"] = "client",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 6,
          },
          {
            ["active"] = true,
            ["name"] = "company_member",
            ["op"] = {
              ["update"] = {
                ["req"] = false,
                ["type"] = "`$STRING`",
              },
            },
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 7,
          },
          {
            ["active"] = true,
            ["name"] = "consignee",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 8,
          },
          {
            ["active"] = true,
            ["name"] = "consignor",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 9,
          },
          {
            ["active"] = true,
            ["name"] = "created_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 10,
          },
          {
            ["active"] = true,
            ["name"] = "declarant",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 11,
          },
          {
            ["active"] = true,
            ["name"] = "document_upload_status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 12,
          },
          {
            ["active"] = true,
            ["name"] = "documents_presentation_requested",
            ["req"] = false,
            ["type"] = "`$BOOLEAN`",
            ["index$"] = 13,
          },
          {
            ["active"] = true,
            ["name"] = "documents_upload_requested",
            ["req"] = false,
            ["type"] = "`$BOOLEAN`",
            ["index$"] = 14,
          },
          {
            ["active"] = true,
            ["name"] = "external_id",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 15,
          },
          {
            ["active"] = true,
            ["name"] = "form",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 16,
          },
          {
            ["active"] = true,
            ["name"] = "goods_presentation_status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 17,
          },
          {
            ["active"] = true,
            ["name"] = "hrcm_status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 18,
          },
          {
            ["active"] = true,
            ["name"] = "id",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 19,
          },
          {
            ["active"] = true,
            ["name"] = "invalidation_status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 20,
          },
          {
            ["active"] = true,
            ["name"] = "is_global_template",
            ["req"] = false,
            ["type"] = "`$BOOLEAN`",
            ["index$"] = 21,
          },
          {
            ["active"] = true,
            ["name"] = "latest_notification_item",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 22,
          },
          {
            ["active"] = true,
            ["name"] = "latest_state",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 23,
          },
          {
            ["active"] = true,
            ["name"] = "lrn",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 24,
          },
          {
            ["active"] = true,
            ["name"] = "mrn",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 25,
          },
          {
            ["active"] = true,
            ["name"] = "name",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 26,
          },
          {
            ["active"] = true,
            ["name"] = "partial_answer",
            ["req"] = false,
            ["type"] = "`$BOOLEAN`",
            ["index$"] = 27,
          },
          {
            ["active"] = true,
            ["name"] = "receipt",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 28,
          },
          {
            ["active"] = true,
            ["name"] = "refund_application_status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 29,
          },
          {
            ["active"] = true,
            ["name"] = "route",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 30,
          },
          {
            ["active"] = true,
            ["name"] = "shipment_items_no",
            ["req"] = false,
            ["type"] = "`$INTEGER`",
            ["index$"] = 31,
          },
          {
            ["active"] = true,
            ["name"] = "shipment_items_quantity_no",
            ["req"] = false,
            ["type"] = "`$INTEGER`",
            ["index$"] = 32,
          },
          {
            ["active"] = true,
            ["name"] = "source",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 33,
          },
          {
            ["active"] = true,
            ["name"] = "source_type",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 34,
          },
          {
            ["active"] = true,
            ["name"] = "status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 35,
          },
          {
            ["active"] = true,
            ["name"] = "template",
            ["req"] = false,
            ["type"] = "`$BOOLEAN`",
            ["index$"] = 36,
          },
          {
            ["active"] = true,
            ["name"] = "template_id",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 37,
          },
          {
            ["active"] = true,
            ["name"] = "template_property",
            ["req"] = false,
            ["type"] = "`$ARRAY`",
            ["index$"] = 38,
          },
          {
            ["active"] = true,
            ["name"] = "total_tax_amount",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 39,
          },
          {
            ["active"] = true,
            ["name"] = "updated_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 40,
          },
          {
            ["active"] = true,
            ["name"] = "verification_error",
            ["req"] = false,
            ["type"] = "`$ARRAY`",
            ["index$"] = 41,
          },
          {
            ["active"] = true,
            ["name"] = "verification_status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 42,
          },
        },
        ["name"] = "submission",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                      ["index$"] = 0,
                    },
                  },
                },
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
                ["index$"] = 0,
              },
              {
                ["active"] = true,
                ["args"] = {},
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
                ["index$"] = 1,
              },
              {
                ["active"] = true,
                ["args"] = {},
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
                ["index$"] = 2,
              },
            },
            ["key$"] = "create",
          },
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {
                  ["query"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "query",
                      ["name"] = "cursor",
                      ["orig"] = "cursor",
                      ["reqd"] = false,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["active"] = true,
                      ["kind"] = "query",
                      ["name"] = "form_subtype",
                      ["orig"] = "form_subtype",
                      ["reqd"] = false,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["active"] = true,
                      ["example"] = "green",
                      ["kind"] = "query",
                      ["name"] = "route",
                      ["orig"] = "route",
                      ["reqd"] = false,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["active"] = true,
                      ["example"] = "released",
                      ["kind"] = "query",
                      ["name"] = "status",
                      ["orig"] = "status",
                      ["reqd"] = false,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["active"] = true,
                      ["kind"] = "query",
                      ["name"] = "template",
                      ["orig"] = "template",
                      ["reqd"] = false,
                      ["type"] = "`$BOOLEAN`",
                    },
                  },
                },
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
                  ["res"] = "`body`",
                },
                ["index$"] = 0,
              },
            },
            ["key$"] = "list",
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                      ["index$"] = 0,
                    },
                  },
                },
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
                ["index$"] = 0,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                ["index$"] = 1,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                ["index$"] = 2,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                ["index$"] = 3,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                ["index$"] = 4,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                ["index$"] = 5,
              },
              {
                ["active"] = true,
                ["args"] = {},
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
                ["index$"] = 6,
              },
            },
            ["key$"] = "load",
          },
          ["remove"] = {
            ["input"] = "data",
            ["name"] = "remove",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                      ["index$"] = 0,
                    },
                  },
                },
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
                ["index$"] = 0,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                ["index$"] = 1,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                ["index$"] = 2,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                ["index$"] = 3,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                ["index$"] = 4,
              },
            },
            ["key$"] = "remove",
          },
          ["update"] = {
            ["input"] = "data",
            ["name"] = "update",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                      ["index$"] = 0,
                    },
                  },
                },
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
                ["index$"] = 0,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                ["index$"] = 1,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                ["index$"] = 2,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                ["index$"] = 3,
              },
              {
                ["active"] = true,
                ["args"] = {
                  ["params"] = {
                    {
                      ["active"] = true,
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "id",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
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
                ["index$"] = 4,
              },
            },
            ["key$"] = "update",
          },
        },
        ["relations"] = {
          ["ancestors"] = {},
        },
      },
      ["submission_detail"] = {
        ["fields"] = {
          {
            ["active"] = true,
            ["name"] = "additional_external_id",
            ["req"] = false,
            ["type"] = "`$ARRAY`",
            ["index$"] = 0,
          },
          {
            ["active"] = true,
            ["name"] = "additional_information",
            ["req"] = true,
            ["type"] = "`$ARRAY`",
            ["index$"] = 1,
          },
          {
            ["active"] = true,
            ["name"] = "amendment_status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 2,
          },
          {
            ["active"] = true,
            ["name"] = "clearance_slip",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 3,
          },
          {
            ["active"] = true,
            ["name"] = "client",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 4,
          },
          {
            ["active"] = true,
            ["name"] = "company",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 5,
          },
          {
            ["active"] = true,
            ["name"] = "company_member",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 6,
          },
          {
            ["active"] = true,
            ["name"] = "consignee",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 7,
          },
          {
            ["active"] = true,
            ["name"] = "consignor",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 8,
          },
          {
            ["active"] = true,
            ["name"] = "created_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 9,
          },
          {
            ["active"] = true,
            ["name"] = "declarant",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 10,
          },
          {
            ["active"] = true,
            ["name"] = "document_upload_status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 11,
          },
          {
            ["active"] = true,
            ["name"] = "documents_presentation_requested",
            ["req"] = false,
            ["type"] = "`$BOOLEAN`",
            ["index$"] = 12,
          },
          {
            ["active"] = true,
            ["name"] = "documents_upload_requested",
            ["req"] = false,
            ["type"] = "`$BOOLEAN`",
            ["index$"] = 13,
          },
          {
            ["active"] = true,
            ["name"] = "external_id",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 14,
          },
          {
            ["active"] = true,
            ["name"] = "form",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 15,
          },
          {
            ["active"] = true,
            ["name"] = "goods_presentation_status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 16,
          },
          {
            ["active"] = true,
            ["name"] = "hrcm_status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 17,
          },
          {
            ["active"] = true,
            ["name"] = "id",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 18,
          },
          {
            ["active"] = true,
            ["name"] = "invalidation_status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 19,
          },
          {
            ["active"] = true,
            ["name"] = "is_global_template",
            ["req"] = false,
            ["type"] = "`$BOOLEAN`",
            ["index$"] = 20,
          },
          {
            ["active"] = true,
            ["name"] = "latest_notification_item",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 21,
          },
          {
            ["active"] = true,
            ["name"] = "latest_state",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 22,
          },
          {
            ["active"] = true,
            ["name"] = "lrn",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 23,
          },
          {
            ["active"] = true,
            ["name"] = "mrn",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 24,
          },
          {
            ["active"] = true,
            ["name"] = "name",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 25,
          },
          {
            ["active"] = true,
            ["name"] = "receipt",
            ["req"] = true,
            ["type"] = "`$OBJECT`",
            ["index$"] = 26,
          },
          {
            ["active"] = true,
            ["name"] = "refund_application_status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 27,
          },
          {
            ["active"] = true,
            ["name"] = "route",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 28,
          },
          {
            ["active"] = true,
            ["name"] = "shipment_items_no",
            ["req"] = false,
            ["type"] = "`$INTEGER`",
            ["index$"] = 29,
          },
          {
            ["active"] = true,
            ["name"] = "shipment_items_quantity_no",
            ["req"] = false,
            ["type"] = "`$INTEGER`",
            ["index$"] = 30,
          },
          {
            ["active"] = true,
            ["name"] = "source",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 31,
          },
          {
            ["active"] = true,
            ["name"] = "source_type",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 32,
          },
          {
            ["active"] = true,
            ["name"] = "status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 33,
          },
          {
            ["active"] = true,
            ["name"] = "submission",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 34,
          },
          {
            ["active"] = true,
            ["name"] = "supporting_document",
            ["req"] = false,
            ["type"] = "`$ARRAY`",
            ["index$"] = 35,
          },
          {
            ["active"] = true,
            ["name"] = "template",
            ["req"] = false,
            ["type"] = "`$BOOLEAN`",
            ["index$"] = 36,
          },
          {
            ["active"] = true,
            ["name"] = "total_tax_amount",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 37,
          },
          {
            ["active"] = true,
            ["name"] = "updated_at",
            ["req"] = true,
            ["type"] = "`$STRING`",
            ["index$"] = 38,
          },
          {
            ["active"] = true,
            ["name"] = "verification_error",
            ["req"] = false,
            ["type"] = "`$ARRAY`",
            ["index$"] = 39,
          },
          {
            ["active"] = true,
            ["name"] = "verification_status",
            ["req"] = false,
            ["type"] = "`$STRING`",
            ["index$"] = 40,
          },
        },
        ["name"] = "submission_detail",
        ["op"] = {
          ["create"] = {
            ["input"] = "data",
            ["name"] = "create",
            ["points"] = {
              {
                ["active"] = true,
                ["args"] = {},
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
                ["index$"] = 0,
              },
            },
            ["key$"] = "create",
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
