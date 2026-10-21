import { describe, test } from 'node:test'
import { SDK } from '..'
import { runDefinitionPoint } from './definition-runner'
import { isControlSkipped } from './utility'


// Generated from the API definition, not from the model this SDK was built
// from: the route, the declared query parameters, the credential the security
// scheme names, and the definition's own response example.
const PLAN: any[] = [
  {
    "entity": "bulk_upload",
    "accessor": "BulkUpload",
    "op": "create",
    "method": "POST",
    "path": "/bulk-uploads",
    "args": [],
    "select": {},
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": {
      "id": "x",
      "file": {
        "id": "x",
        "url": "x",
        "name": "x",
        "extension": "x",
        "size": 1,
        "public": true,
        "updated_at": "2026-01-01T00:00:00Z",
        "created_at": "2026-01-01T00:00:00Z",
        "company": "x"
      },
      "errors_file": {
        "id": "x",
        "url": "x",
        "name": "x",
        "extension": "x",
        "size": 1,
        "public": true,
        "updated_at": "2026-01-01T00:00:00Z",
        "created_at": "2026-01-01T00:00:00Z",
        "company": "x"
      },
      "deleted_at": "2026-01-01T00:00:00Z",
      "updated_at": "2026-01-01T00:00:00Z",
      "created_at": "2026-01-01T00:00:00Z",
      "status": "pending",
      "receipt_request_status": "pending",
      "receipt_request_started_at": "2026-01-01T00:00:00Z",
      "receipt_generating_started_at": "2026-01-01T00:00:00Z",
      "parser": "aes_platform",
      "declarations_no": 1,
      "processed_declarations_no": 1,
      "parsed_declarations_no": 1,
      "failed_declarations_no": 1,
      "processing_started_at": "2026-01-01T00:00:00Z",
      "processing_ended_at": "2026-01-01T00:00:00Z",
      "parsing_started_at": "2026-01-01T00:00:00Z",
      "parsing_completed_at": "2026-01-01T00:00:00Z",
      "green_routed_no": 1,
      "red_routed_no": 1,
      "orange_routed_no": 1,
      "yellow_routed_no": 1,
      "rejected_status_no": 1,
      "external_id": "x",
      "passive_transport_number": "x",
      "active_transport_number": "x",
      "house_transport_doc_ref": "x",
      "issue_date": "2026-01-01T00:00:00Z",
      "departure_datetime": "2026-01-01T00:00:00Z",
      "arrival_datetime": "2026-01-01T00:00:00Z",
      "client": "x",
      "declarant": "x",
      "company_member": "x",
      "template": "x",
      "h1_fallback_template": "x",
      "mapping": "x",
      "receipts_zip": "x",
      "receipt_request_user": "x",
      "passive_transport_nationality": "x",
      "active_transport_nationality": "x"
    },
    "idField": "id"
  },
  {
    "entity": "bulk_upload",
    "accessor": "BulkUpload",
    "op": "list",
    "method": "GET",
    "path": "/bulk-uploads",
    "args": [],
    "select": {
      "cursor": "v1"
    },
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [
      "cursor"
    ],
    "queryArgs": [
      {
        "name": "cursor",
        "wire": "cursor"
      }
    ],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": {
      "next": "x",
      "previous": "x",
      "results": [
        {
          "id": "x",
          "file": {
            "id": "x",
            "url": "x",
            "name": "x",
            "extension": "x",
            "size": 1,
            "public": true,
            "updated_at": "2026-01-01T00:00:00Z",
            "created_at": "2026-01-01T00:00:00Z",
            "company": "x"
          },
          "errors_file": {
            "id": "x",
            "url": "x",
            "name": "x",
            "extension": "x",
            "size": 1,
            "public": true,
            "updated_at": "2026-01-01T00:00:00Z",
            "created_at": "2026-01-01T00:00:00Z",
            "company": "x"
          },
          "deleted_at": "2026-01-01T00:00:00Z",
          "updated_at": "2026-01-01T00:00:00Z",
          "created_at": "2026-01-01T00:00:00Z",
          "status": "pending",
          "receipt_request_status": "pending",
          "receipt_request_started_at": "2026-01-01T00:00:00Z",
          "receipt_generating_started_at": "2026-01-01T00:00:00Z",
          "parser": "aes_platform",
          "declarations_no": 1,
          "processed_declarations_no": 1,
          "parsed_declarations_no": 1,
          "failed_declarations_no": 1,
          "processing_started_at": "2026-01-01T00:00:00Z",
          "processing_ended_at": "2026-01-01T00:00:00Z",
          "parsing_started_at": "2026-01-01T00:00:00Z",
          "parsing_completed_at": "2026-01-01T00:00:00Z",
          "green_routed_no": 1,
          "red_routed_no": 1,
          "orange_routed_no": 1,
          "yellow_routed_no": 1,
          "rejected_status_no": 1,
          "external_id": "x",
          "passive_transport_number": "x",
          "active_transport_number": "x",
          "house_transport_doc_ref": "x",
          "issue_date": "2026-01-01T00:00:00Z",
          "departure_datetime": "2026-01-01T00:00:00Z",
          "arrival_datetime": "2026-01-01T00:00:00Z",
          "client": "x",
          "declarant": "x",
          "company_member": "x",
          "template": "x",
          "h1_fallback_template": "x",
          "mapping": "x",
          "receipts_zip": "x",
          "receipt_request_user": "x",
          "passive_transport_nationality": "x",
          "active_transport_nationality": "x"
        }
      ]
    },
    "idField": "id"
  },
  {
    "entity": "bulk_upload",
    "accessor": "BulkUpload",
    "op": "load",
    "method": "GET",
    "path": "/bulk-uploads/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "bulk_upload",
    "accessor": "BulkUpload",
    "op": "load",
    "method": "GET",
    "path": "/bulk-uploads/{id}/generate-pdfs",
    "action": "generate_pdf",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "bulk_upload",
    "accessor": "BulkUpload",
    "op": "remove",
    "method": "DELETE",
    "path": "/bulk-uploads/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 204,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "bulk_upload",
    "accessor": "BulkUpload",
    "op": "remove",
    "method": "DELETE",
    "path": "/bulk-uploads/{id}/generate-pdfs",
    "action": "generate_pdf",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 204,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "bulk_upload",
    "accessor": "BulkUpload",
    "op": "update",
    "method": "PATCH",
    "path": "/bulk-uploads/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "bulk_upload",
    "accessor": "BulkUpload",
    "op": "update",
    "method": "PATCH",
    "path": "/bulk-uploads/{id}/generate-pdfs",
    "action": "generate_pdf",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "file",
    "accessor": "File",
    "op": "create",
    "method": "POST",
    "path": "/files",
    "args": [],
    "select": {},
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": {
      "id": "x",
      "url": "x",
      "name": "x",
      "extension": "x",
      "size": 1,
      "public": true,
      "updated_at": "2026-01-01T00:00:00Z",
      "created_at": "2026-01-01T00:00:00Z",
      "company": "x"
    },
    "idField": "id"
  },
  {
    "entity": "party",
    "accessor": "Party",
    "op": "create",
    "method": "POST",
    "path": "/parties",
    "args": [],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "party",
    "accessor": "Party",
    "op": "list",
    "method": "GET",
    "path": "/parties",
    "args": [],
    "select": {
      "cursor": "v1",
      "type": "v1"
    },
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [
      "cursor",
      "type"
    ],
    "queryArgs": [
      {
        "name": "cursor",
        "wire": "cursor"
      },
      {
        "name": "type",
        "wire": "type"
      }
    ],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": {
      "next": "x",
      "previous": "x",
      "results": [
        {
          "id": "x",
          "address": {
            "id": "x",
            "deleted_at": "2026-01-01T00:00:00Z",
            "updated_at": "2026-01-01T00:00:00Z",
            "created_at": "2026-01-01T00:00:00Z",
            "raw": "x",
            "line_1": "x",
            "city": "x",
            "postal_code": "x",
            "country_code": "x",
            "sub_division": "x",
            "number": "x",
            "po_box": "x"
          },
          "authorisation": {
            "id": "x",
            "deleted_at": "2026-01-01T00:00:00Z",
            "updated_at": "2026-01-01T00:00:00Z",
            "created_at": "2026-01-01T00:00:00Z",
            "value": "x",
            "name": "x",
            "description": "x",
            "list": "x"
          },
          "additional_declaration_type": {
            "id": "x",
            "deleted_at": "2026-01-01T00:00:00Z",
            "updated_at": "2026-01-01T00:00:00Z",
            "created_at": "2026-01-01T00:00:00Z",
            "value": "x",
            "name": "x",
            "description": "x",
            "list": "x"
          },
          "preferred_payment_method": {
            "id": "x",
            "deleted_at": "2026-01-01T00:00:00Z",
            "updated_at": "2026-01-01T00:00:00Z",
            "created_at": "2026-01-01T00:00:00Z",
            "value": "x",
            "name": "x",
            "description": "x",
            "list": "x"
          },
          "customs_office_of_lodgement": {
            "id": "x",
            "deleted_at": "2026-01-01T00:00:00Z",
            "updated_at": "2026-01-01T00:00:00Z",
            "created_at": "2026-01-01T00:00:00Z",
            "value": "x",
            "name": "x",
            "description": "x",
            "list": "x"
          },
          "certificate_type": "x",
          "signed_form": {
            "id": "x",
            "url": "x",
            "name": "x",
            "extension": "x",
            "size": 1,
            "public": true,
            "updated_at": "2026-01-01T00:00:00Z",
            "created_at": "2026-01-01T00:00:00Z",
            "company": "x"
          },
          "certificate": {
            "type": "ros",
            "last_fetched_at": "2026-01-01T00:00:00Z",
            "processing_start_at": "2026-01-01T00:00:00Z",
            "frequent_processing_ends_at": "2026-01-01T00:00:00Z"
          },
          "type_of_person": {
            "id": "x",
            "deleted_at": "2026-01-01T00:00:00Z",
            "updated_at": "2026-01-01T00:00:00Z",
            "created_at": "2026-01-01T00:00:00Z",
            "value": "x",
            "name": "x",
            "description": "x",
            "list": "x"
          },
          "deleted_at": "2026-01-01T00:00:00Z",
          "updated_at": "2026-01-01T00:00:00Z",
          "created_at": "2026-01-01T00:00:00Z",
          "person_paying_customs_duty": "x",
          "unlocode": "x",
          "identification_number": "x",
          "indirect_representative": true,
          "name": "x",
          "type": "exporter",
          "email": "x",
          "phone_country_code": "x",
          "phone_number": "x",
          "nhd_submission_counter": 1,
          "nhd_last_submission_year": 1,
          "bank_details": "x",
          "company": "x"
        }
      ]
    },
    "idField": "id"
  },
  {
    "entity": "party",
    "accessor": "Party",
    "op": "load",
    "method": "GET",
    "path": "/parties/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "party",
    "accessor": "Party",
    "op": "remove",
    "method": "DELETE",
    "path": "/parties/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 204,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "party",
    "accessor": "Party",
    "op": "update",
    "method": "PATCH",
    "path": "/parties/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "create",
    "method": "POST",
    "path": "/submissions/{id}/refund",
    "action": "refund",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "create",
    "method": "POST",
    "path": "/submissions",
    "args": [],
    "select": {},
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": {
      "id": "x",
      "lrn": "x",
      "mrn": "x",
      "form": "x",
      "name": "x",
      "route": "green",
      "company_member": {
        "user": {
          "id": "x",
          "role": "company_member",
          "first_name": "x",
          "last_name": "x",
          "email": "x"
        },
        "role": "staff",
        "company": {
          "id": "x",
          "name": "x",
          "description": "x",
          "updated_at": "2026-01-01T00:00:00Z",
          "created_at": "2026-01-01T00:00:00Z"
        },
        "created_at": "2026-01-01T00:00:00Z",
        "hidden_submission_statuses": [
          "draft"
        ]
      },
      "declarant": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "consignor": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "consignee": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "client": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "status": "draft",
      "template": true,
      "is_global_template": true,
      "latest_state": {
        "id": "x",
        "deleted_at": "2026-01-01T00:00:00Z",
        "updated_at": "2026-01-01T00:00:00Z",
        "created_at": "2026-01-01T00:00:00Z",
        "type": "created",
        "status": "draft",
        "amendment_status": "requested",
        "invalidation_status": "in_review",
        "goods_presentation_status": "in_review",
        "hrcm_status": "requested",
        "document_upload_status": "requested",
        "refund_application_status": "processing",
        "route": "green",
        "reason": "x",
        "message": "x",
        "lrn": "x",
        "transaction_id": "x",
        "documents_upload_requested": true,
        "documents_presentation_requested": true,
        "processing_time": 1,
        "submission": "x",
        "company_member": "x"
      },
      "shipment_items_no": 1,
      "shipment_items_quantity_no": 1,
      "invalidation_status": "in_review",
      "goods_presentation_status": "in_review",
      "documents_upload_requested": true,
      "documents_presentation_requested": true,
      "amendment_status": "requested",
      "hrcm_status": "requested",
      "document_upload_status": "requested",
      "refund_application_status": "processing",
      "total_tax_amount": "x",
      "external_id": "x",
      "additional_external_ids": [
        "x"
      ],
      "source": "x",
      "source_type": "template",
      "clearance_slip": {
        "id": "x",
        "url": "x",
        "name": "x",
        "extension": "x",
        "size": 1,
        "public": true,
        "updated_at": "2026-01-01T00:00:00Z",
        "created_at": "2026-01-01T00:00:00Z",
        "company": "x"
      },
      "receipt": {
        "id": "x",
        "url": "x",
        "name": "x",
        "extension": "x",
        "size": 1,
        "public": true,
        "updated_at": "2026-01-01T00:00:00Z",
        "created_at": "2026-01-01T00:00:00Z",
        "company": "x"
      },
      "verification_status": "passed",
      "verification_errors": [
        "x"
      ],
      "created_at": "2026-01-01T00:00:00Z",
      "updated_at": "2026-01-01T00:00:00Z",
      "latest_notification_item": "x"
    },
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "create",
    "method": "POST",
    "path": "/submissions/retrieve",
    "action": "retrieve",
    "args": [],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "list",
    "method": "GET",
    "path": "/submissions",
    "args": [],
    "select": {
      "cursor": "v1",
      "form_subtype": "v1",
      "route": "green",
      "status": "released",
      "template": "v1"
    },
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [
      "cursor",
      "form_subtype",
      "route",
      "status",
      "template"
    ],
    "queryArgs": [
      {
        "name": "cursor",
        "wire": "cursor"
      },
      {
        "name": "form_subtype",
        "wire": "form_subtype"
      },
      {
        "name": "route",
        "wire": "route"
      },
      {
        "name": "status",
        "wire": "status"
      },
      {
        "name": "template",
        "wire": "template"
      }
    ],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": {
      "next": "x",
      "previous": "x",
      "results": [
        {
          "id": "x",
          "lrn": "x",
          "mrn": "x",
          "form": {
            "id": "x",
            "name": "x",
            "type": "ais",
            "subtype": "H1",
            "description": "x"
          },
          "name": "x",
          "route": "green",
          "company_member": {
            "user": {
              "id": "x",
              "role": "company_member",
              "first_name": "x",
              "last_name": "x",
              "email": "x"
            },
            "role": "staff",
            "company": {
              "id": "x",
              "name": "x",
              "description": "x",
              "updated_at": "2026-01-01T00:00:00Z",
              "created_at": "2026-01-01T00:00:00Z"
            },
            "created_at": "2026-01-01T00:00:00Z",
            "hidden_submission_statuses": [
              "draft"
            ]
          },
          "declarant": {
            "id": "x",
            "name": "x",
            "identification_number": "x"
          },
          "consignor": {
            "id": "x",
            "name": "x",
            "identification_number": "x"
          },
          "consignee": {
            "id": "x",
            "name": "x",
            "identification_number": "x"
          },
          "client": {
            "id": "x",
            "name": "x",
            "identification_number": "x"
          },
          "status": "draft",
          "template": true,
          "is_global_template": true,
          "latest_state": {
            "id": "x",
            "company_member": {
              "user": {
                "id": "x",
                "role": "company_member",
                "first_name": "x",
                "last_name": "x",
                "email": "x"
              },
              "role": "staff",
              "company": {
                "id": "x",
                "name": "x",
                "description": "x",
                "updated_at": "2026-01-01T00:00:00Z",
                "created_at": "2026-01-01T00:00:00Z"
              },
              "created_at": "2026-01-01T00:00:00Z",
              "hidden_submission_statuses": [
                "draft"
              ]
            },
            "deleted_at": "2026-01-01T00:00:00Z",
            "updated_at": "2026-01-01T00:00:00Z",
            "created_at": "2026-01-01T00:00:00Z",
            "type": "created",
            "status": "draft",
            "amendment_status": "requested",
            "invalidation_status": "in_review",
            "goods_presentation_status": "in_review",
            "hrcm_status": "requested",
            "document_upload_status": "requested",
            "refund_application_status": "processing",
            "route": "green",
            "reason": "x",
            "message": "x",
            "lrn": "x",
            "transaction_id": "x",
            "documents_upload_requested": true,
            "documents_presentation_requested": true,
            "processing_time": 1
          },
          "shipment_items_no": 1,
          "shipment_items_quantity_no": 1,
          "invalidation_status": "in_review",
          "goods_presentation_status": "in_review",
          "documents_upload_requested": true,
          "documents_presentation_requested": true,
          "amendment_status": "requested",
          "hrcm_status": "requested",
          "document_upload_status": "requested",
          "refund_application_status": "processing",
          "total_tax_amount": "x",
          "external_id": "x",
          "additional_external_ids": [
            "x"
          ],
          "source": "x",
          "source_type": "template",
          "clearance_slip": {
            "id": "x",
            "url": "x",
            "name": "x",
            "extension": "x",
            "size": 1,
            "public": true,
            "updated_at": "2026-01-01T00:00:00Z",
            "created_at": "2026-01-01T00:00:00Z",
            "company": "x"
          },
          "receipt": {
            "id": "x",
            "url": "x",
            "name": "x",
            "extension": "x",
            "size": 1,
            "public": true,
            "updated_at": "2026-01-01T00:00:00Z",
            "created_at": "2026-01-01T00:00:00Z",
            "company": "x"
          },
          "verification_status": "passed",
          "verification_errors": [
            "x"
          ],
          "created_at": "2026-01-01T00:00:00Z",
          "updated_at": "2026-01-01T00:00:00Z",
          "latest_notification_item": "x"
        }
      ]
    },
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "load",
    "method": "GET",
    "path": "/submissions/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": {
      "id": "x",
      "lrn": "x",
      "mrn": "x",
      "form": "x",
      "name": "x",
      "route": "green",
      "company_member": {
        "user": {
          "id": "x",
          "role": "company_member",
          "first_name": "x",
          "last_name": "x",
          "email": "x"
        },
        "role": "staff",
        "company": {
          "id": "x",
          "name": "x",
          "description": "x",
          "updated_at": "2026-01-01T00:00:00Z",
          "created_at": "2026-01-01T00:00:00Z"
        },
        "created_at": "2026-01-01T00:00:00Z",
        "hidden_submission_statuses": [
          "draft"
        ]
      },
      "declarant": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "consignor": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "consignee": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "client": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "status": "draft",
      "template": true,
      "is_global_template": true,
      "latest_state": {
        "id": "x",
        "deleted_at": "2026-01-01T00:00:00Z",
        "updated_at": "2026-01-01T00:00:00Z",
        "created_at": "2026-01-01T00:00:00Z",
        "type": "created",
        "status": "draft",
        "amendment_status": "requested",
        "invalidation_status": "in_review",
        "goods_presentation_status": "in_review",
        "hrcm_status": "requested",
        "document_upload_status": "requested",
        "refund_application_status": "processing",
        "route": "green",
        "reason": "x",
        "message": "x",
        "lrn": "x",
        "transaction_id": "x",
        "documents_upload_requested": true,
        "documents_presentation_requested": true,
        "processing_time": 1,
        "submission": "x",
        "company_member": "x"
      },
      "shipment_items_no": 1,
      "shipment_items_quantity_no": 1,
      "invalidation_status": "in_review",
      "goods_presentation_status": "in_review",
      "documents_upload_requested": true,
      "documents_presentation_requested": true,
      "amendment_status": "requested",
      "hrcm_status": "requested",
      "document_upload_status": "requested",
      "refund_application_status": "processing",
      "total_tax_amount": "x",
      "external_id": "x",
      "additional_external_ids": [
        "x"
      ],
      "source": "x",
      "source_type": "template",
      "clearance_slip": {
        "id": "x",
        "url": "x",
        "name": "x",
        "extension": "x",
        "size": 1,
        "public": true,
        "updated_at": "2026-01-01T00:00:00Z",
        "created_at": "2026-01-01T00:00:00Z",
        "company": "x"
      },
      "receipt": {
        "id": "x",
        "url": "x",
        "name": "x",
        "extension": "x",
        "size": 1,
        "public": true,
        "updated_at": "2026-01-01T00:00:00Z",
        "created_at": "2026-01-01T00:00:00Z",
        "company": "x"
      },
      "verification_status": "passed",
      "verification_errors": [
        "x"
      ],
      "created_at": "2026-01-01T00:00:00Z",
      "updated_at": "2026-01-01T00:00:00Z",
      "latest_notification_item": "x"
    },
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "load",
    "method": "GET",
    "path": "/submissions/{id}/clearance-slip",
    "action": "clearance_slip",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "load",
    "method": "GET",
    "path": "/submissions/{id}/notification-read",
    "action": "notification_read",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "load",
    "method": "GET",
    "path": "/submissions/{id}/pbn-applicable",
    "action": "pbn_applicable",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "load",
    "method": "GET",
    "path": "/submissions/{id}/receipt",
    "action": "receipt",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "load",
    "method": "GET",
    "path": "/submissions/{id}/refund",
    "action": "refund",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "load",
    "method": "GET",
    "path": "/submissions/retrieve",
    "action": "retrieve",
    "args": [],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "remove",
    "method": "DELETE",
    "path": "/submissions/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 204,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "remove",
    "method": "DELETE",
    "path": "/submissions/{id}/clearance-slip",
    "action": "clearance_slip",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 204,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "remove",
    "method": "DELETE",
    "path": "/submissions/{id}/notification-read",
    "action": "notification_read",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 204,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "remove",
    "method": "DELETE",
    "path": "/submissions/{id}/pbn-applicable",
    "action": "pbn_applicable",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 204,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "remove",
    "method": "DELETE",
    "path": "/submissions/{id}/receipt",
    "action": "receipt",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 204,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "update",
    "method": "PATCH",
    "path": "/submissions/{id}",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": {
      "id": "x",
      "lrn": "x",
      "mrn": "x",
      "form": "x",
      "name": "x",
      "route": "green",
      "company_member": {
        "user": {
          "id": "x",
          "role": "company_member",
          "first_name": "x",
          "last_name": "x",
          "email": "x"
        },
        "role": "staff",
        "company": {
          "id": "x",
          "name": "x",
          "description": "x",
          "updated_at": "2026-01-01T00:00:00Z",
          "created_at": "2026-01-01T00:00:00Z"
        },
        "created_at": "2026-01-01T00:00:00Z",
        "hidden_submission_statuses": [
          "draft"
        ]
      },
      "declarant": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "consignor": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "consignee": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "client": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "status": "draft",
      "template": true,
      "is_global_template": true,
      "latest_state": {
        "id": "x",
        "deleted_at": "2026-01-01T00:00:00Z",
        "updated_at": "2026-01-01T00:00:00Z",
        "created_at": "2026-01-01T00:00:00Z",
        "type": "created",
        "status": "draft",
        "amendment_status": "requested",
        "invalidation_status": "in_review",
        "goods_presentation_status": "in_review",
        "hrcm_status": "requested",
        "document_upload_status": "requested",
        "refund_application_status": "processing",
        "route": "green",
        "reason": "x",
        "message": "x",
        "lrn": "x",
        "transaction_id": "x",
        "documents_upload_requested": true,
        "documents_presentation_requested": true,
        "processing_time": 1,
        "submission": "x",
        "company_member": "x"
      },
      "shipment_items_no": 1,
      "shipment_items_quantity_no": 1,
      "invalidation_status": "in_review",
      "goods_presentation_status": "in_review",
      "documents_upload_requested": true,
      "documents_presentation_requested": true,
      "amendment_status": "requested",
      "hrcm_status": "requested",
      "document_upload_status": "requested",
      "refund_application_status": "processing",
      "total_tax_amount": "x",
      "external_id": "x",
      "additional_external_ids": [
        "x"
      ],
      "source": "x",
      "source_type": "template",
      "clearance_slip": {
        "id": "x",
        "url": "x",
        "name": "x",
        "extension": "x",
        "size": 1,
        "public": true,
        "updated_at": "2026-01-01T00:00:00Z",
        "created_at": "2026-01-01T00:00:00Z",
        "company": "x"
      },
      "receipt": {
        "id": "x",
        "url": "x",
        "name": "x",
        "extension": "x",
        "size": 1,
        "public": true,
        "updated_at": "2026-01-01T00:00:00Z",
        "created_at": "2026-01-01T00:00:00Z",
        "company": "x"
      },
      "verification_status": "passed",
      "verification_errors": [
        "x"
      ],
      "created_at": "2026-01-01T00:00:00Z",
      "updated_at": "2026-01-01T00:00:00Z",
      "latest_notification_item": "x"
    },
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "update",
    "method": "PATCH",
    "path": "/submissions/{id}/clearance-slip",
    "action": "clearance_slip",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "update",
    "method": "PATCH",
    "path": "/submissions/{id}/notification-read",
    "action": "notification_read",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "update",
    "method": "PATCH",
    "path": "/submissions/{id}/pbn-applicable",
    "action": "pbn_applicable",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission",
    "accessor": "Submission",
    "op": "update",
    "method": "PATCH",
    "path": "/submissions/{id}/receipt",
    "action": "receipt",
    "args": [
      {
        "name": "id",
        "wire": "id",
        "value": "p1"
      }
    ],
    "select": {},
    "headers": [],
    "cookies": [],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": null,
    "idField": "id"
  },
  {
    "entity": "submission_detail",
    "accessor": "SubmissionDetail",
    "op": "create",
    "method": "POST",
    "path": "/documents-request",
    "args": [],
    "select": {},
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [],
    "queryArgs": [],
    "auth": [
      [
        {
          "in": "header",
          "name": "authorization"
        }
      ]
    ],
    "status": 200,
    "sample": {
      "id": "x",
      "lrn": "x",
      "mrn": "x",
      "form": "x",
      "name": "x",
      "route": "green",
      "company_member": {
        "user": {
          "id": "x",
          "role": "company_member",
          "first_name": "x",
          "last_name": "x",
          "email": "x"
        },
        "role": "staff",
        "company": {
          "id": "x",
          "name": "x",
          "description": "x",
          "updated_at": "2026-01-01T00:00:00Z",
          "created_at": "2026-01-01T00:00:00Z"
        },
        "created_at": "2026-01-01T00:00:00Z",
        "hidden_submission_statuses": [
          "draft"
        ]
      },
      "declarant": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "consignor": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "consignee": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "client": {
        "id": "x",
        "name": "x",
        "identification_number": "x"
      },
      "status": "draft",
      "template": true,
      "is_global_template": true,
      "latest_state": {
        "id": "x",
        "deleted_at": "2026-01-01T00:00:00Z",
        "updated_at": "2026-01-01T00:00:00Z",
        "created_at": "2026-01-01T00:00:00Z",
        "type": "created",
        "status": "draft",
        "amendment_status": "requested",
        "invalidation_status": "in_review",
        "goods_presentation_status": "in_review",
        "hrcm_status": "requested",
        "document_upload_status": "requested",
        "refund_application_status": "processing",
        "route": "green",
        "reason": "x",
        "message": "x",
        "lrn": "x",
        "transaction_id": "x",
        "documents_upload_requested": true,
        "documents_presentation_requested": true,
        "processing_time": 1,
        "submission": "x",
        "company_member": "x"
      },
      "shipment_items_no": 1,
      "shipment_items_quantity_no": 1,
      "invalidation_status": "in_review",
      "goods_presentation_status": "in_review",
      "documents_upload_requested": true,
      "documents_presentation_requested": true,
      "amendment_status": "requested",
      "hrcm_status": "requested",
      "document_upload_status": "requested",
      "refund_application_status": "processing",
      "total_tax_amount": "x",
      "external_id": "x",
      "additional_external_ids": [
        "x"
      ],
      "source": "x",
      "source_type": "template",
      "clearance_slip": {
        "id": "x",
        "url": "x",
        "name": "x",
        "extension": "x",
        "size": 1,
        "public": true,
        "updated_at": "2026-01-01T00:00:00Z",
        "created_at": "2026-01-01T00:00:00Z",
        "company": "x"
      },
      "receipt": {
        "id": "x",
        "url": "x",
        "name": "x",
        "extension": "x",
        "size": 1,
        "public": true,
        "updated_at": "2026-01-01T00:00:00Z",
        "created_at": "2026-01-01T00:00:00Z",
        "company": "x"
      },
      "verification_status": "passed",
      "verification_errors": [
        "x"
      ],
      "created_at": "2026-01-01T00:00:00Z",
      "updated_at": "2026-01-01T00:00:00Z",
      "latest_notification_item": "x"
    },
    "idField": "id"
  }
]


describe('definition', () => {
  for (const point of PLAN) {
    test(point.entity + '.' + point.op + ' ' + point.method + ' ' + point.path, async (t) => {
      const control = isControlSkipped('entityOp', point.entity + '.' + point.op, 'definition')
      if (control.skip) {
        t.skip(control.reason || 'skipped via sdk-test-control.json')
        return
      }
      await runDefinitionPoint(SDK, point)
    })
  }
})
