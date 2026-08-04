# CustomsWindow Golang SDK



The Golang SDK for the CustomsWindow API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.BulkUpload(nil)` — each with the same small set of operations (`List`, `Load`, `Create`, `Update`, `Remove`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/customs-window-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/customs-window-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/customs-window-sdk/go=../customs-window-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    "os"
    sdk "github.com/voxgig-sdk/customs-window-sdk/go"
)

func main() {
    client := sdk.NewCustomsWindowSDK(map[string]any{
        "apikey": os.Getenv("CUSTOMS_WINDOW_APIKEY"),
    })

    // List bulkUpload records — the value is the array of records itself.
    bulkUploads, err := client.BulkUpload(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }
    for _, item := range bulkUploads.([]any) {
        fmt.Println(item)
    }

    // Load a single bulkUpload — the value is the loaded record.
    bulkUpload, err := client.BulkUpload(nil).Load(map[string]any{"id": "example_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(bulkUpload)

    // Create a bulkUpload.
    created, err := client.BulkUpload(nil).Create(map[string]any{"client": "example_client", "company_member": "example_company_member", "created_at": "example_created_at", "declarant": "example_declarant", "declarations_no": 1, "errors_file": map[string]any{}, "file": map[string]any{}, "updated_at": "example_updated_at"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(created)

    // Update a bulkUpload.
    updated, err := client.BulkUpload(nil).Update(map[string]any{"id": "example_id", "active_transport_nationality": "example_active_transport_nationality", "active_transport_number": "example_active_transport_number"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(updated)

    // Remove a bulkUpload.
    removed, err := client.BulkUpload(nil).Remove(map[string]any{"id": "example_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(removed)
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
submissions, err := client.Submission(nil).List(nil, nil)
if err != nil {
    // handle err
    return
}
_ = submissions
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

submission, err := client.Submission(nil).List(
    nil, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(submission) // the returned mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewCustomsWindowSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
CUSTOMS_WINDOW_TEST_LIVE=TRUE
CUSTOMS_WINDOW_APIKEY=<your-key>
```

Then run:

```bash
cd go && go test ./test/...
```


## Reference

### NewCustomsWindowSDK

```go
func NewCustomsWindowSDK(options map[string]any) *CustomsWindowSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"apikey"` | `string` | API key for authentication. |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *CustomsWindowSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### CustomsWindowSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `BulkUpload` | `(data map[string]any) CustomsWindowEntity` | Create a BulkUpload entity instance. |
| `File` | `(data map[string]any) CustomsWindowEntity` | Create a File entity instance. |
| `PaginatedBulkUploadListList` | `(data map[string]any) CustomsWindowEntity` | Create a PaginatedBulkUploadListList entity instance. |
| `PaginatedPartyListList` | `(data map[string]any) CustomsWindowEntity` | Create a PaginatedPartyListList entity instance. |
| `PaginatedSubmissionListList` | `(data map[string]any) CustomsWindowEntity` | Create a PaginatedSubmissionListList entity instance. |
| `Party` | `(data map[string]any) CustomsWindowEntity` | Create a Party entity instance. |
| `Submission` | `(data map[string]any) CustomsWindowEntity` | Create a Submission entity instance. |
| `SubmissionDetail` | `(data map[string]any) CustomsWindowEntity` | Create a SubmissionDetail entity instance. |

### Entity interface (CustomsWindowEntity)

All entities implement the `CustomsWindowEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Update` | `(reqdata, ctrl map[string]any) (any, error)` | Update an existing entity. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` / `Update` / `Remove` | the entity record (`map[string]any`) |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    bulkUpload, err := client.BulkUpload(nil).List(map[string]any{/* fields */}, nil)
    if err != nil { /* handle */ }
    // bulkUpload is the returned record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### BulkUpload

| Field | Description |
| --- | --- |
| `"active_transport_nationality"` |  |
| `"active_transport_number"` |  |
| `"arrival_datetime"` |  |
| `"client"` |  |
| `"company_member"` |  |
| `"created_at"` |  |
| `"declarant"` |  |
| `"declarations_no"` |  |
| `"deleted_at"` |  |
| `"departure_datetime"` |  |
| `"errors_file"` |  |
| `"external_id"` |  |
| `"failed_declarations_no"` |  |
| `"file"` |  |
| `"green_routed_no"` |  |
| `"h1_fallback_template"` |  |
| `"house_transport_doc_ref"` |  |
| `"id"` |  |
| `"issue_date"` |  |
| `"mapping"` |  |
| `"orange_routed_no"` |  |
| `"parsed_declarations_no"` |  |
| `"parser"` |  |
| `"parsing_completed_at"` |  |
| `"parsing_started_at"` |  |
| `"passive_transport_nationality"` |  |
| `"passive_transport_number"` |  |
| `"processed_declarations_no"` |  |
| `"processing_ended_at"` |  |
| `"processing_started_at"` |  |
| `"receipt_generating_started_at"` |  |
| `"receipt_request_started_at"` |  |
| `"receipt_request_status"` |  |
| `"receipt_request_user"` |  |
| `"receipts_zip"` |  |
| `"red_routed_no"` |  |
| `"rejected_status_no"` |  |
| `"status"` |  |
| `"template"` |  |
| `"updated_at"` |  |
| `"yellow_routed_no"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/bulk-uploads`

#### File

| Field | Description |
| --- | --- |
| `"company"` |  |
| `"created_at"` |  |
| `"extension"` |  |
| `"file"` |  |
| `"id"` |  |
| `"name"` |  |
| `"public"` |  |
| `"size"` |  |
| `"updated_at"` |  |
| `"url"` |  |

Operations: Create.

API path: `/files`

#### PaginatedBulkUploadListList

| Field | Description |
| --- | --- |

Operations: .

API path: ``

#### PaginatedPartyListList

| Field | Description |
| --- | --- |

Operations: .

API path: ``

#### PaginatedSubmissionListList

| Field | Description |
| --- | --- |

Operations: .

API path: ``

#### Party

| Field | Description |
| --- | --- |
| `"additional_declaration_type"` |  |
| `"address"` |  |
| `"authorisation"` |  |
| `"bank_detail"` |  |
| `"certificate"` |  |
| `"certificate_type"` |  |
| `"company"` |  |
| `"created_at"` |  |
| `"customs_office_of_lodgement"` |  |
| `"deleted_at"` |  |
| `"email"` |  |
| `"id"` |  |
| `"identification_number"` |  |
| `"indirect_representative"` |  |
| `"name"` |  |
| `"nhd_last_submission_year"` |  |
| `"nhd_submission_counter"` |  |
| `"person_paying_customs_duty"` |  |
| `"phone_country_code"` |  |
| `"phone_number"` |  |
| `"preferred_payment_method"` |  |
| `"signed_form"` |  |
| `"type"` |  |
| `"type_of_person"` |  |
| `"unlocode"` |  |
| `"updated_at"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/parties`

#### Submission

| Field | Description |
| --- | --- |
| `"additional_external_id"` |  |
| `"amendment_reason"` |  |
| `"amendment_status"` |  |
| `"answer"` |  |
| `"bypass_restricted_code"` |  |
| `"clearance_slip"` |  |
| `"client"` |  |
| `"company_member"` |  |
| `"consignee"` |  |
| `"consignor"` |  |
| `"created_at"` |  |
| `"declarant"` |  |
| `"document_upload_status"` |  |
| `"documents_presentation_requested"` |  |
| `"documents_upload_requested"` |  |
| `"external_id"` |  |
| `"form"` |  |
| `"goods_presentation_status"` |  |
| `"hrcm_status"` |  |
| `"id"` |  |
| `"invalidation_status"` |  |
| `"is_global_template"` |  |
| `"latest_notification_item"` |  |
| `"latest_state"` |  |
| `"lrn"` |  |
| `"mrn"` |  |
| `"name"` |  |
| `"partial_answer"` |  |
| `"receipt"` |  |
| `"refund_application_status"` |  |
| `"route"` |  |
| `"shipment_items_no"` |  |
| `"shipment_items_quantity_no"` |  |
| `"source"` |  |
| `"source_type"` |  |
| `"status"` |  |
| `"template"` |  |
| `"template_id"` |  |
| `"template_property"` |  |
| `"total_tax_amount"` |  |
| `"updated_at"` |  |
| `"verification_error"` |  |
| `"verification_status"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/submissions/{id}/refund`

#### SubmissionDetail

| Field | Description |
| --- | --- |
| `"additional_external_id"` |  |
| `"additional_information"` |  |
| `"amendment_status"` |  |
| `"clearance_slip"` |  |
| `"client"` |  |
| `"company"` |  |
| `"company_member"` |  |
| `"consignee"` |  |
| `"consignor"` |  |
| `"created_at"` |  |
| `"declarant"` |  |
| `"document_upload_status"` |  |
| `"documents_presentation_requested"` |  |
| `"documents_upload_requested"` |  |
| `"external_id"` |  |
| `"form"` |  |
| `"goods_presentation_status"` |  |
| `"hrcm_status"` |  |
| `"id"` |  |
| `"invalidation_status"` |  |
| `"is_global_template"` |  |
| `"latest_notification_item"` |  |
| `"latest_state"` |  |
| `"lrn"` |  |
| `"mrn"` |  |
| `"name"` |  |
| `"receipt"` |  |
| `"refund_application_status"` |  |
| `"route"` |  |
| `"shipment_items_no"` |  |
| `"shipment_items_quantity_no"` |  |
| `"source"` |  |
| `"source_type"` |  |
| `"status"` |  |
| `"submission"` |  |
| `"supporting_document"` |  |
| `"template"` |  |
| `"total_tax_amount"` |  |
| `"updated_at"` |  |
| `"verification_error"` |  |
| `"verification_status"` |  |

Operations: Create.

API path: `/documents-request`



## Entities


### BulkUpload

Create an instance: `bulkUpload := client.BulkUpload(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `active_transport_nationality` | `string` |  |
| `active_transport_number` | `string` |  |
| `arrival_datetime` | `string` |  |
| `client` | `string` |  |
| `company_member` | `string` |  |
| `created_at` | `string` |  |
| `declarant` | `string` |  |
| `declarations_no` | `int` |  |
| `deleted_at` | `string` |  |
| `departure_datetime` | `string` |  |
| `errors_file` | `map[string]any` |  |
| `external_id` | `string` |  |
| `failed_declarations_no` | `int` |  |
| `file` | `map[string]any` |  |
| `green_routed_no` | `int` |  |
| `h1_fallback_template` | `string` |  |
| `house_transport_doc_ref` | `string` |  |
| `id` | `string` |  |
| `issue_date` | `string` |  |
| `mapping` | `string` |  |
| `orange_routed_no` | `int` |  |
| `parsed_declarations_no` | `int` |  |
| `parser` | `string` |  |
| `parsing_completed_at` | `string` |  |
| `parsing_started_at` | `string` |  |
| `passive_transport_nationality` | `string` |  |
| `passive_transport_number` | `string` |  |
| `processed_declarations_no` | `int` |  |
| `processing_ended_at` | `string` |  |
| `processing_started_at` | `string` |  |
| `receipt_generating_started_at` | `string` |  |
| `receipt_request_started_at` | `string` |  |
| `receipt_request_status` | `string` |  |
| `receipt_request_user` | `string` |  |
| `receipts_zip` | `string` |  |
| `red_routed_no` | `int` |  |
| `rejected_status_no` | `int` |  |
| `status` | `string` |  |
| `template` | `string` |  |
| `updated_at` | `string` |  |
| `yellow_routed_no` | `int` |  |

#### Example: Load

```go
bulkUpload, err := client.BulkUpload(nil).Load(map[string]any{"id": "bulk_upload_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(bulkUpload) // the loaded record
```

#### Example: List

```go
bulkUploads, err := client.BulkUpload(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(bulkUploads) // the array of records
```

#### Example: Create

```go
result, err := client.BulkUpload(nil).Create(map[string]any{
    "client": "example_client",
    "company_member": "example_company_member",
    "created_at": "example_created_at",
    "declarant": "example_declarant",
    "declarations_no": 1,
    "errors_file": map[string]any{},
    "file": map[string]any{},
    "updated_at": "example_updated_at",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### File

Create an instance: `file := client.File(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `company` | `string` |  |
| `created_at` | `string` |  |
| `extension` | `string` |  |
| `file` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `public` | `bool` |  |
| `size` | `int` |  |
| `updated_at` | `string` |  |
| `url` | `string` |  |

#### Example: Create

```go
result, err := client.File(nil).Create(map[string]any{
    "company": "example_company",
    "created_at": "example_created_at",
    "extension": "example_extension",
    "file": "example_file",
    "name": "example_name",
    "public": true,
    "updated_at": "example_updated_at",
    "url": "example_url",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PaginatedBulkUploadListList

Create an instance: `paginatedBulkUploadListList := client.PaginatedBulkUploadListList(nil)`


### PaginatedPartyListList

Create an instance: `paginatedPartyListList := client.PaginatedPartyListList(nil)`


### PaginatedSubmissionListList

Create an instance: `paginatedSubmissionListList := client.PaginatedSubmissionListList(nil)`


### Party

Create an instance: `party := client.Party(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_declaration_type` | `map[string]any` |  |
| `address` | `map[string]any` |  |
| `authorisation` | `map[string]any` |  |
| `bank_detail` | `string` |  |
| `certificate` | `map[string]any` |  |
| `certificate_type` | `string` |  |
| `company` | `string` |  |
| `created_at` | `string` |  |
| `customs_office_of_lodgement` | `map[string]any` |  |
| `deleted_at` | `string` |  |
| `email` | `string` |  |
| `id` | `string` |  |
| `identification_number` | `string` |  |
| `indirect_representative` | `bool` |  |
| `name` | `string` |  |
| `nhd_last_submission_year` | `int` |  |
| `nhd_submission_counter` | `int` |  |
| `person_paying_customs_duty` | `string` |  |
| `phone_country_code` | `string` |  |
| `phone_number` | `string` |  |
| `preferred_payment_method` | `map[string]any` |  |
| `signed_form` | `map[string]any` |  |
| `type` | `string` |  |
| `type_of_person` | `map[string]any` |  |
| `unlocode` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```go
party, err := client.Party(nil).Load(map[string]any{"id": "party_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(party) // the loaded record
```

#### Example: List

```go
partys, err := client.Party(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(partys) // the array of records
```

#### Example: Create

```go
result, err := client.Party(nil).Create(map[string]any{
    "additional_declaration_type": map[string]any{},
    "address": map[string]any{},
    "authorisation": map[string]any{},
    "certificate": map[string]any{},
    "certificate_type": "example_certificate_type",
    "company": "example_company",
    "created_at": "example_created_at",
    "customs_office_of_lodgement": map[string]any{},
    "preferred_payment_method": map[string]any{},
    "signed_form": map[string]any{},
    "type_of_person": map[string]any{},
    "updated_at": "example_updated_at",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Submission

Create an instance: `submission := client.Submission(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_external_id` | `[]any` |  |
| `amendment_reason` | `string` |  |
| `amendment_status` | `string` |  |
| `answer` | `[]any` |  |
| `bypass_restricted_code` | `bool` |  |
| `clearance_slip` | `map[string]any` |  |
| `client` | `map[string]any` |  |
| `company_member` | `map[string]any` |  |
| `consignee` | `map[string]any` |  |
| `consignor` | `map[string]any` |  |
| `created_at` | `string` |  |
| `declarant` | `map[string]any` |  |
| `document_upload_status` | `string` |  |
| `documents_presentation_requested` | `bool` |  |
| `documents_upload_requested` | `bool` |  |
| `external_id` | `string` |  |
| `form` | `string` |  |
| `goods_presentation_status` | `string` |  |
| `hrcm_status` | `string` |  |
| `id` | `string` |  |
| `invalidation_status` | `string` |  |
| `is_global_template` | `bool` |  |
| `latest_notification_item` | `string` |  |
| `latest_state` | `map[string]any` |  |
| `lrn` | `string` |  |
| `mrn` | `string` |  |
| `name` | `string` |  |
| `partial_answer` | `bool` |  |
| `receipt` | `map[string]any` |  |
| `refund_application_status` | `string` |  |
| `route` | `string` |  |
| `shipment_items_no` | `int` |  |
| `shipment_items_quantity_no` | `int` |  |
| `source` | `string` |  |
| `source_type` | `string` |  |
| `status` | `string` |  |
| `template` | `bool` |  |
| `template_id` | `string` |  |
| `template_property` | `[]any` |  |
| `total_tax_amount` | `string` |  |
| `updated_at` | `string` |  |
| `verification_error` | `[]any` |  |
| `verification_status` | `string` |  |

#### Example: Load

```go
submission, err := client.Submission(nil).Load(map[string]any{"id": "submission_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(submission) // the loaded record
```

#### Example: List

```go
submissions, err := client.Submission(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(submissions) // the array of records
```

#### Example: Create

```go
result, err := client.Submission(nil).Create(map[string]any{
    "answer": []any{},
    "clearance_slip": map[string]any{},
    "client": map[string]any{},
    "company_member": map[string]any{},
    "consignee": map[string]any{},
    "consignor": map[string]any{},
    "created_at": "example_created_at",
    "declarant": map[string]any{},
    "form": "example_form",
    "latest_state": map[string]any{},
    "receipt": map[string]any{},
    "updated_at": "example_updated_at",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### SubmissionDetail

Create an instance: `submissionDetail := client.SubmissionDetail(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_external_id` | `[]any` |  |
| `additional_information` | `[]any` |  |
| `amendment_status` | `string` |  |
| `clearance_slip` | `map[string]any` |  |
| `client` | `map[string]any` |  |
| `company` | `string` |  |
| `company_member` | `map[string]any` |  |
| `consignee` | `map[string]any` |  |
| `consignor` | `map[string]any` |  |
| `created_at` | `string` |  |
| `declarant` | `map[string]any` |  |
| `document_upload_status` | `string` |  |
| `documents_presentation_requested` | `bool` |  |
| `documents_upload_requested` | `bool` |  |
| `external_id` | `string` |  |
| `form` | `string` |  |
| `goods_presentation_status` | `string` |  |
| `hrcm_status` | `string` |  |
| `id` | `string` |  |
| `invalidation_status` | `string` |  |
| `is_global_template` | `bool` |  |
| `latest_notification_item` | `string` |  |
| `latest_state` | `map[string]any` |  |
| `lrn` | `string` |  |
| `mrn` | `string` |  |
| `name` | `string` |  |
| `receipt` | `map[string]any` |  |
| `refund_application_status` | `string` |  |
| `route` | `string` |  |
| `shipment_items_no` | `int` |  |
| `shipment_items_quantity_no` | `int` |  |
| `source` | `string` |  |
| `source_type` | `string` |  |
| `status` | `string` |  |
| `submission` | `string` |  |
| `supporting_document` | `[]any` |  |
| `template` | `bool` |  |
| `total_tax_amount` | `string` |  |
| `updated_at` | `string` |  |
| `verification_error` | `[]any` |  |
| `verification_status` | `string` |  |

#### Example: Create

```go
result, err := client.SubmissionDetail(nil).Create(map[string]any{
    "additional_information": []any{},
    "clearance_slip": map[string]any{},
    "client": map[string]any{},
    "company": "example_company",
    "company_member": map[string]any{},
    "consignee": map[string]any{},
    "consignor": map[string]any{},
    "created_at": "example_created_at",
    "declarant": map[string]any{},
    "form": "example_form",
    "latest_state": map[string]any{},
    "receipt": map[string]any{},
    "submission": "example_submission",
    "updated_at": "example_updated_at",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/customs-window-sdk/go/
├── customs-window.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/customs-window-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `List`, the entity
stores the returned data and match criteria internally.

```go
submission := client.Submission(nil)
submission.List(nil, nil)

// submission.Data() now returns the submission data from the last list
// submission.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
