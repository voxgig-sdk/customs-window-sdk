# CustomsWindow Golang SDK



The Golang SDK for the CustomsWindow API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.BulkUpload(nil)` — each with the same small set of operations (`List`, `Load`, `Create`, `Update`, `Remove`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Also generated from this model: `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb`, `ts` — see
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
| `"active_transport_nationality"` | cuid-format identifier for this entity. |
| `"active_transport_number"` |  |
| `"arrival_datetime"` |  |
| `"client"` | cuid-format identifier for this entity. |
| `"company_member"` | cuid-format identifier for this entity. |
| `"created_at"` |  |
| `"declarant"` | cuid-format identifier for this entity. |
| `"declarations_no"` |  |
| `"deleted_at"` |  |
| `"departure_datetime"` |  |
| `"errors_file"` |  |
| `"external_id"` |  |
| `"failed_declarations_no"` |  |
| `"file"` | cuid-format identifier for this entity. |
| `"green_routed_no"` |  |
| `"h1_fallback_template"` | cuid-format identifier for this entity. |
| `"house_transport_doc_ref"` |  |
| `"id"` | cuid-format identifier for this entity. |
| `"issue_date"` |  |
| `"mapping"` | cuid-format identifier for this entity. |
| `"orange_routed_no"` |  |
| `"parsed_declarations_no"` |  |
| `"parser"` | * `aes_platform` - aes_platform * `cds_platform` - cds_platform * `cds_export_platform` - cds_export_platform * `g4_g3` - g4_g3 * `nhd_platform` - nhd_platform * `platform` - platform * `birds` - birds * `ics2_platform` - ics2_platform |
| `"parsing_completed_at"` |  |
| `"parsing_started_at"` |  |
| `"passive_transport_nationality"` | cuid-format identifier for this entity. |
| `"passive_transport_number"` |  |
| `"processed_declarations_no"` |  |
| `"processing_ended_at"` |  |
| `"processing_started_at"` |  |
| `"receipt_generating_started_at"` |  |
| `"receipt_request_started_at"` |  |
| `"receipt_request_status"` | * `pending` - pending * `processing` - processing * `generating` - generating * `completed` - completed |
| `"receipt_request_user"` | cuid-format identifier for this entity. |
| `"receipts_zip"` | cuid-format identifier for this entity. |
| `"red_routed_no"` |  |
| `"rejected_status_no"` |  |
| `"status"` | * `pending` - pending * `parsing` - parsing * `processing` - processing * `complete` - complete |
| `"template"` | cuid-format identifier for this entity. |
| `"updated_at"` |  |
| `"yellow_routed_no"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/bulk-uploads`

#### File

| Field | Description |
| --- | --- |
| `"company"` | cuid-format identifier for this entity. |
| `"created_at"` |  |
| `"extension"` |  |
| `"file"` |  |
| `"id"` | cuid-format identifier for this entity. |
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
| `"bank_details"` |  |
| `"certificate"` |  |
| `"certificate_type"` |  |
| `"company"` | cuid-format identifier for this entity. |
| `"created_at"` |  |
| `"customs_office_of_lodgement"` |  |
| `"deleted_at"` |  |
| `"email"` |  |
| `"id"` | cuid-format identifier for this entity. |
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
| `"type"` | * `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `co… |
| `"type_of_person"` |  |
| `"unlocode"` |  |
| `"updated_at"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/parties`

#### Submission

| Field | Description |
| --- | --- |
| `"additional_external_ids"` |  |
| `"amendment_reason"` |  |
| `"amendment_status"` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `"answers"` |  |
| `"bypass_restricted_code"` |  |
| `"clearance_slip"` |  |
| `"client"` |  |
| `"company_member"` | cuid-format identifier for this entity. |
| `"consignee"` |  |
| `"consignor"` |  |
| `"created_at"` |  |
| `"declarant"` | cuid-format identifier for this entity. |
| `"document_upload_status"` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `"documents_presentation_requested"` |  |
| `"documents_upload_requested"` |  |
| `"external_id"` |  |
| `"form"` | cuid-format identifier for this entity. |
| `"goods_presentation_status"` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested |
| `"hrcm_status"` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `"id"` | cuid-format identifier for this entity. |
| `"invalidation_status"` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled |
| `"is_global_template"` |  |
| `"latest_notification_item"` | cuid-format identifier for this entity. |
| `"latest_state"` |  |
| `"lrn"` |  |
| `"mrn"` |  |
| `"name"` |  |
| `"partial_answers"` |  |
| `"receipt"` |  |
| `"refund_application_status"` | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `"route"` | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `"shipment_items_no"` |  |
| `"shipment_items_quantity_no"` |  |
| `"source"` | cuid-format identifier for this entity. |
| `"source_type"` | * `template` - template * `automated_import` - automated_import |
| `"status"` | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `"template"` |  |
| `"template_id"` | cuid-format identifier for this entity. |
| `"template_properties"` |  |
| `"total_tax_amount"` |  |
| `"updated_at"` |  |
| `"verification_errors"` |  |
| `"verification_status"` | * `passed` - passed * `failed` - failed |

Operations: Create, List, Load, Remove, Update.

API path: `/submissions/{id}/refund`

#### SubmissionDetail

| Field | Description |
| --- | --- |
| `"additional_external_ids"` |  |
| `"additional_information"` |  |
| `"amendment_status"` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `"clearance_slip"` |  |
| `"client"` |  |
| `"company"` | cuid-format identifier for this entity. |
| `"company_member"` | cuid-format identifier for this entity. |
| `"consignee"` |  |
| `"consignor"` |  |
| `"created_at"` |  |
| `"declarant"` |  |
| `"document_upload_status"` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `"documents_presentation_requested"` |  |
| `"documents_upload_requested"` |  |
| `"external_id"` |  |
| `"form"` |  |
| `"goods_presentation_status"` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested |
| `"hrcm_status"` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `"id"` | cuid-format identifier for this entity. |
| `"invalidation_status"` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled |
| `"is_global_template"` |  |
| `"latest_notification_item"` | cuid-format identifier for this entity. |
| `"latest_state"` |  |
| `"lrn"` |  |
| `"mrn"` |  |
| `"name"` |  |
| `"receipt"` |  |
| `"refund_application_status"` | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `"route"` | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `"shipment_items_no"` |  |
| `"shipment_items_quantity_no"` |  |
| `"source"` | cuid-format identifier for this entity. |
| `"source_type"` | * `template` - template * `automated_import` - automated_import |
| `"status"` | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `"submission"` | cuid-format identifier for this entity. |
| `"supporting_documents"` |  |
| `"template"` |  |
| `"total_tax_amount"` |  |
| `"updated_at"` |  |
| `"verification_errors"` |  |
| `"verification_status"` | * `passed` - passed * `failed` - failed |

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
| `active_transport_nationality` | `string` | cuid-format identifier for this entity. |
| `active_transport_number` | `string` |  |
| `arrival_datetime` | `string` |  |
| `client` | `string` | cuid-format identifier for this entity. |
| `company_member` | `string` | cuid-format identifier for this entity. |
| `created_at` | `string` |  |
| `declarant` | `string` | cuid-format identifier for this entity. |
| `declarations_no` | `int` |  |
| `deleted_at` | `string` |  |
| `departure_datetime` | `string` |  |
| `errors_file` | `map[string]any` |  |
| `external_id` | `string` |  |
| `failed_declarations_no` | `int` |  |
| `file` | `map[string]any` | cuid-format identifier for this entity. |
| `green_routed_no` | `int` |  |
| `h1_fallback_template` | `string` | cuid-format identifier for this entity. |
| `house_transport_doc_ref` | `string` |  |
| `id` | `string` | cuid-format identifier for this entity. |
| `issue_date` | `string` |  |
| `mapping` | `string` | cuid-format identifier for this entity. |
| `orange_routed_no` | `int` |  |
| `parsed_declarations_no` | `int` |  |
| `parser` | `string` | * `aes_platform` - aes_platform * `cds_platform` - cds_platform * `cds_export_platform` - cds_export_platform * `g4_g3` - g4_g3 * `nhd_platform` - nhd_platform * `platform` - platform * `birds` - birds * `ics2_platform` - ics2_platform |
| `parsing_completed_at` | `string` |  |
| `parsing_started_at` | `string` |  |
| `passive_transport_nationality` | `string` | cuid-format identifier for this entity. |
| `passive_transport_number` | `string` |  |
| `processed_declarations_no` | `int` |  |
| `processing_ended_at` | `string` |  |
| `processing_started_at` | `string` |  |
| `receipt_generating_started_at` | `string` |  |
| `receipt_request_started_at` | `string` |  |
| `receipt_request_status` | `string` | * `pending` - pending * `processing` - processing * `generating` - generating * `completed` - completed |
| `receipt_request_user` | `string` | cuid-format identifier for this entity. |
| `receipts_zip` | `string` | cuid-format identifier for this entity. |
| `red_routed_no` | `int` |  |
| `rejected_status_no` | `int` |  |
| `status` | `string` | * `pending` - pending * `parsing` - parsing * `processing` - processing * `complete` - complete |
| `template` | `string` | cuid-format identifier for this entity. |
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
| `company` | `string` | cuid-format identifier for this entity. |
| `created_at` | `string` |  |
| `extension` | `string` |  |
| `file` | `string` |  |
| `id` | `string` | cuid-format identifier for this entity. |
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
| `bank_details` | `string` |  |
| `certificate` | `map[string]any` |  |
| `certificate_type` | `string` |  |
| `company` | `string` | cuid-format identifier for this entity. |
| `created_at` | `string` |  |
| `customs_office_of_lodgement` | `map[string]any` |  |
| `deleted_at` | `string` |  |
| `email` | `string` |  |
| `id` | `string` | cuid-format identifier for this entity. |
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
| `type` | `string` | * `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `co… |
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
| `additional_external_ids` | `[]any` |  |
| `amendment_reason` | `string` |  |
| `amendment_status` | `string` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `answers` | `[]any` |  |
| `bypass_restricted_code` | `bool` |  |
| `clearance_slip` | `map[string]any` |  |
| `client` | `map[string]any` |  |
| `company_member` | `map[string]any` | cuid-format identifier for this entity. |
| `consignee` | `map[string]any` |  |
| `consignor` | `map[string]any` |  |
| `created_at` | `string` |  |
| `declarant` | `map[string]any` | cuid-format identifier for this entity. |
| `document_upload_status` | `string` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `documents_presentation_requested` | `bool` |  |
| `documents_upload_requested` | `bool` |  |
| `external_id` | `string` |  |
| `form` | `string` | cuid-format identifier for this entity. |
| `goods_presentation_status` | `string` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested |
| `hrcm_status` | `string` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `id` | `string` | cuid-format identifier for this entity. |
| `invalidation_status` | `string` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled |
| `is_global_template` | `bool` |  |
| `latest_notification_item` | `string` | cuid-format identifier for this entity. |
| `latest_state` | `map[string]any` |  |
| `lrn` | `string` |  |
| `mrn` | `string` |  |
| `name` | `string` |  |
| `partial_answers` | `bool` |  |
| `receipt` | `map[string]any` |  |
| `refund_application_status` | `string` | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `route` | `string` | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `shipment_items_no` | `int` |  |
| `shipment_items_quantity_no` | `int` |  |
| `source` | `string` | cuid-format identifier for this entity. |
| `source_type` | `string` | * `template` - template * `automated_import` - automated_import |
| `status` | `string` | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `template` | `bool` |  |
| `template_id` | `string` | cuid-format identifier for this entity. |
| `template_properties` | `[]any` |  |
| `total_tax_amount` | `string` |  |
| `updated_at` | `string` |  |
| `verification_errors` | `[]any` |  |
| `verification_status` | `string` | * `passed` - passed * `failed` - failed |

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
    "answers": []any{},
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
| `additional_external_ids` | `[]any` |  |
| `additional_information` | `[]any` |  |
| `amendment_status` | `string` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `clearance_slip` | `map[string]any` |  |
| `client` | `map[string]any` |  |
| `company` | `string` | cuid-format identifier for this entity. |
| `company_member` | `map[string]any` | cuid-format identifier for this entity. |
| `consignee` | `map[string]any` |  |
| `consignor` | `map[string]any` |  |
| `created_at` | `string` |  |
| `declarant` | `map[string]any` |  |
| `document_upload_status` | `string` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `documents_presentation_requested` | `bool` |  |
| `documents_upload_requested` | `bool` |  |
| `external_id` | `string` |  |
| `form` | `string` |  |
| `goods_presentation_status` | `string` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested |
| `hrcm_status` | `string` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `id` | `string` | cuid-format identifier for this entity. |
| `invalidation_status` | `string` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled |
| `is_global_template` | `bool` |  |
| `latest_notification_item` | `string` | cuid-format identifier for this entity. |
| `latest_state` | `map[string]any` |  |
| `lrn` | `string` |  |
| `mrn` | `string` |  |
| `name` | `string` |  |
| `receipt` | `map[string]any` |  |
| `refund_application_status` | `string` | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `route` | `string` | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `shipment_items_no` | `int` |  |
| `shipment_items_quantity_no` | `int` |  |
| `source` | `string` | cuid-format identifier for this entity. |
| `source_type` | `string` | * `template` - template * `automated_import` - automated_import |
| `status` | `string` | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `submission` | `string` | cuid-format identifier for this entity. |
| `supporting_documents` | `[]any` |  |
| `template` | `bool` |  |
| `total_tax_amount` | `string` |  |
| `updated_at` | `string` |  |
| `verification_errors` | `[]any` |  |
| `verification_status` | `string` | * `passed` - passed * `failed` - failed |

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

## Features

This SDK ships 8 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`debug`](#debug) | Request/response capture ring buffer for debugging |
| [`idempotency`](#idempotency) | Idempotency keys for safe retries of mutating operations |
| [`metrics`](#metrics) | Statistics capture: per-operation counters and latency |
| [`paging`](#paging) | Pagination signals for list operations |
| [`ratelimit`](#ratelimit) | Client-side rate limiting via a token bucket |
| [`retry`](#retry) | Automatic retry of transient failures with exponential backoff |
| [`test`](#test) | In-memory mock transport for testing without a live server |
| [`timeout`](#timeout) | Per-request timeout with transport abort |

> **Order matters for `ratelimit`, `retry`, `timeout`.** These wrap the
> transport, so each one wraps whatever is already installed: the order you
> activate them in IS the nesting order. Activating them as an ordered list
> rather than a map is what fixes that order.

### debug

Request/response capture ring buffer for debugging.

| Option | Default |
|---|---|
| `active` | `false` |
| `max` | `100` |
| `redact` | `['authorization', 'cookie', 'set-cookie', 'api-key', 'apikey', 'x-api-key', 'idempotency-key']` |

Set `feature.debug.active` to enable it, then override any of the options above.

### idempotency

Idempotency keys for safe retries of mutating operations.

| Option | Default |
|---|---|
| `active` | `false` |
| `header` | `'Idempotency-Key'` |
| `methods` | `['POST', 'PUT', 'PATCH', 'DELETE']` |
| `ops` | `['create', 'update', 'remove']` |

Set `feature.idempotency.active` to enable it, then override any of the options above.

### metrics

Statistics capture: per-operation counters and latency.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.metrics.active` to enable it, then override any of the options above.

### paging

Pagination signals for list operations.

| Option | Default |
|---|---|
| `active` | `false` |
| `afterVar` | `'after'` |
| `cursorParam` | `'cursor'` |
| `firstVar` | `'first'` |
| `limitParam` | `'limit'` |
| `pageParam` | `'page'` |
| `startPage` | `1` |

Set `feature.paging.active` to enable it, then override any of the options above.

### ratelimit

Client-side rate limiting via a token bucket.

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

Set `feature.ratelimit.active` to enable it, then override any of the options above.

`ratelimit` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### retry

Automatic retry of transient failures with exponential backoff.

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

Set `feature.retry.active` to enable it, then override any of the options above.

`retry` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### test

In-memory mock transport for testing without a live server.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.

### timeout

Per-request timeout with transport abort.

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

Set `feature.timeout.active` to enable it, then override any of the options above.

`timeout` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.


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

- **DebugFeature**: Request/response capture ring buffer for debugging
- **IdempotencyFeature**: Idempotency keys for safe retries of mutating operations
- **MetricsFeature**: Statistics capture: per-operation counters and latency
- **PagingFeature**: Pagination signals for list operations
- **RatelimitFeature**: Client-side rate limiting via a token bucket
- **RetryFeature**: Automatic retry of transient failures with exponential backoff
- **TestFeature**: In-memory mock transport for testing without a live server
- **TimeoutFeature**: Per-request timeout with transport abort

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
