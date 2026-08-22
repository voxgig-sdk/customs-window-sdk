# CustomsWindow Golang SDK Reference

Complete API reference for the CustomsWindow Golang SDK.


## CustomsWindowSDK

### Constructor

```go
func NewCustomsWindowSDK(options map[string]any) *CustomsWindowSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *CustomsWindowSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *CustomsWindowSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `BulkUpload(data map[string]any) CustomsWindowEntity`

Create a new `BulkUpload` entity instance. Pass `nil` for no initial data.

#### `File(data map[string]any) CustomsWindowEntity`

Create a new `File` entity instance. Pass `nil` for no initial data.

#### `PaginatedBulkUploadListList(data map[string]any) CustomsWindowEntity`

Create a new `PaginatedBulkUploadListList` entity instance. Pass `nil` for no initial data.

#### `PaginatedPartyListList(data map[string]any) CustomsWindowEntity`

Create a new `PaginatedPartyListList` entity instance. Pass `nil` for no initial data.

#### `PaginatedSubmissionListList(data map[string]any) CustomsWindowEntity`

Create a new `PaginatedSubmissionListList` entity instance. Pass `nil` for no initial data.

#### `Party(data map[string]any) CustomsWindowEntity`

Create a new `Party` entity instance. Pass `nil` for no initial data.

#### `Submission(data map[string]any) CustomsWindowEntity`

Create a new `Submission` entity instance. Pass `nil` for no initial data.

#### `SubmissionDetail(data map[string]any) CustomsWindowEntity`

Create a new `SubmissionDetail` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## BulkUploadEntity

```go
bulkUpload := client.BulkUpload(nil)
fmt.Println(bulkUpload.GetName()) // "bulk_upload"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `active_transport_nationality` | `string` | No | cuid-format identifier for this entity. |
| `active_transport_number` | `string` | No |  |
| `arrival_datetime` | `string` | No |  |
| `client` | `string` | Yes | cuid-format identifier for this entity. |
| `company_member` | `string` | Yes | cuid-format identifier for this entity. |
| `created_at` | `string` | Yes |  |
| `declarant` | `string` | Yes | cuid-format identifier for this entity. |
| `declarations_no` | `int` | Yes |  |
| `deleted_at` | `string` | No |  |
| `departure_datetime` | `string` | No |  |
| `errors_file` | `map[string]any` | Yes |  |
| `external_id` | `string` | No |  |
| `failed_declarations_no` | `int` | No |  |
| `file` | `map[string]any` | Yes | cuid-format identifier for this entity. |
| `green_routed_no` | `int` | No |  |
| `h1_fallback_template` | `string` | No | cuid-format identifier for this entity. |
| `house_transport_doc_ref` | `string` | No |  |
| `id` | `string` | No | cuid-format identifier for this entity. |
| `issue_date` | `string` | No |  |
| `mapping` | `string` | No | cuid-format identifier for this entity. |
| `orange_routed_no` | `int` | No |  |
| `parsed_declarations_no` | `int` | No |  |
| `parser` | `string` | No | * `aes_platform` - aes_platform * `cds_platform` - cds_platform * `cds_export_platform` - cds_export_platform * `g4_g3` - g4_g3 * `nhd_platform` - nhd_platform * `platform` - platform * `birds` - birds * `ics2_platform` - ics2_platform |
| `parsing_completed_at` | `string` | No |  |
| `parsing_started_at` | `string` | No |  |
| `passive_transport_nationality` | `string` | No | cuid-format identifier for this entity. |
| `passive_transport_number` | `string` | No |  |
| `processed_declarations_no` | `int` | No |  |
| `processing_ended_at` | `string` | No |  |
| `processing_started_at` | `string` | No |  |
| `receipt_generating_started_at` | `string` | No |  |
| `receipt_request_started_at` | `string` | No |  |
| `receipt_request_status` | `string` | No | * `pending` - pending * `processing` - processing * `generating` - generating * `completed` - completed |
| `receipt_request_user` | `string` | No | cuid-format identifier for this entity. |
| `receipts_zip` | `string` | No | cuid-format identifier for this entity. |
| `red_routed_no` | `int` | No |  |
| `rejected_status_no` | `int` | No |  |
| `status` | `string` | No | * `pending` - pending * `parsing` - parsing * `processing` - processing * `complete` - complete |
| `template` | `string` | No | cuid-format identifier for this entity. |
| `updated_at` | `string` | Yes |  |
| `yellow_routed_no` | `int` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.BulkUpload(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.BulkUpload(nil).Load(map[string]any{"id": "bulk_upload_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

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

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.BulkUpload(nil).Update(map[string]any{
    "id": "bulk_upload_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.BulkUpload(nil).Remove(map[string]any{"id": "bulk_upload_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `BulkUploadEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## FileEntity

```go
file := client.File(nil)
fmt.Println(file.GetName()) // "file"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `company` | `string` | Yes | cuid-format identifier for this entity. |
| `created_at` | `string` | Yes |  |
| `extension` | `string` | Yes |  |
| `file` | `string` | Yes |  |
| `id` | `string` | No | cuid-format identifier for this entity. |
| `name` | `string` | Yes |  |
| `public` | `bool` | Yes |  |
| `size` | `int` | No |  |
| `updated_at` | `string` | Yes |  |
| `url` | `string` | Yes |  |

### Field Usage by Operation

| Field | create |
| --- | --- |
| `company` | - |
| `created_at` | - |
| `extension` | - |
| `file` | - |
| `id` | - |
| `name` | - |
| `public` | Yes |
| `size` | - |
| `updated_at` | - |
| `url` | - |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

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

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `FileEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PaginatedBulkUploadListListEntity

```go
paginatedBulkUploadListList := client.PaginatedBulkUploadListList(nil)
fmt.Println(paginatedBulkUploadListList.GetName()) // "paginated_bulk_upload_list_list"
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PaginatedBulkUploadListListEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PaginatedPartyListListEntity

```go
paginatedPartyListList := client.PaginatedPartyListList(nil)
fmt.Println(paginatedPartyListList.GetName()) // "paginated_party_list_list"
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PaginatedPartyListListEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PaginatedSubmissionListListEntity

```go
paginatedSubmissionListList := client.PaginatedSubmissionListList(nil)
fmt.Println(paginatedSubmissionListList.GetName()) // "paginated_submission_list_list"
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PaginatedSubmissionListListEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PartyEntity

```go
party := client.Party(nil)
fmt.Println(party.GetName()) // "party"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_declaration_type` | `map[string]any` | Yes |  |
| `address` | `map[string]any` | Yes |  |
| `authorisation` | `map[string]any` | Yes |  |
| `bank_details` | `string` | No |  |
| `certificate` | `map[string]any` | Yes |  |
| `certificate_type` | `string` | Yes |  |
| `company` | `string` | Yes | cuid-format identifier for this entity. |
| `created_at` | `string` | Yes |  |
| `customs_office_of_lodgement` | `map[string]any` | Yes |  |
| `deleted_at` | `string` | No |  |
| `email` | `string` | No |  |
| `id` | `string` | No | cuid-format identifier for this entity. |
| `identification_number` | `string` | No |  |
| `indirect_representative` | `bool` | No |  |
| `name` | `string` | No |  |
| `nhd_last_submission_year` | `int` | No |  |
| `nhd_submission_counter` | `int` | No |  |
| `person_paying_customs_duty` | `string` | No |  |
| `phone_country_code` | `string` | No |  |
| `phone_number` | `string` | No |  |
| `preferred_payment_method` | `map[string]any` | Yes |  |
| `signed_form` | `map[string]any` | Yes |  |
| `type` | `string` | No | * `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `co… |
| `type_of_person` | `map[string]any` | Yes |  |
| `unlocode` | `string` | No |  |
| `updated_at` | `string` | Yes |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Party(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Party(nil).Load(map[string]any{"id": "party_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

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

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Party(nil).Update(map[string]any{
    "id": "party_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Party(nil).Remove(map[string]any{"id": "party_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PartyEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SubmissionEntity

```go
submission := client.Submission(nil)
fmt.Println(submission.GetName()) // "submission"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_external_ids` | `[]any` | No |  |
| `amendment_reason` | `string` | No |  |
| `amendment_status` | `string` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `answers` | `[]any` | Yes |  |
| `bypass_restricted_code` | `bool` | No |  |
| `clearance_slip` | `map[string]any` | Yes |  |
| `client` | `map[string]any` | Yes |  |
| `company_member` | `map[string]any` | Yes | cuid-format identifier for this entity. |
| `consignee` | `map[string]any` | Yes |  |
| `consignor` | `map[string]any` | Yes |  |
| `created_at` | `string` | Yes |  |
| `declarant` | `map[string]any` | Yes | cuid-format identifier for this entity. |
| `document_upload_status` | `string` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `documents_presentation_requested` | `bool` | No |  |
| `documents_upload_requested` | `bool` | No |  |
| `external_id` | `string` | No |  |
| `form` | `string` | Yes | cuid-format identifier for this entity. |
| `goods_presentation_status` | `string` | No | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested |
| `hrcm_status` | `string` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `id` | `string` | No | cuid-format identifier for this entity. |
| `invalidation_status` | `string` | No | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled |
| `is_global_template` | `bool` | No |  |
| `latest_notification_item` | `string` | No | cuid-format identifier for this entity. |
| `latest_state` | `map[string]any` | Yes |  |
| `lrn` | `string` | No |  |
| `mrn` | `string` | No |  |
| `name` | `string` | No |  |
| `partial_answers` | `bool` | No |  |
| `receipt` | `map[string]any` | Yes |  |
| `refund_application_status` | `string` | No | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `route` | `string` | No | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `shipment_items_no` | `int` | No |  |
| `shipment_items_quantity_no` | `int` | No |  |
| `source` | `string` | No | cuid-format identifier for this entity. |
| `source_type` | `string` | No | * `template` - template * `automated_import` - automated_import |
| `status` | `string` | No | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `template` | `bool` | No |  |
| `template_id` | `string` | No | cuid-format identifier for this entity. |
| `template_properties` | `[]any` | No |  |
| `total_tax_amount` | `string` | No |  |
| `updated_at` | `string` | Yes |  |
| `verification_errors` | `[]any` | No |  |
| `verification_status` | `string` | No | * `passed` - passed * `failed` - failed |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `additional_external_ids` | - | - | - | - | - |
| `amendment_reason` | - | - | - | - | - |
| `amendment_status` | - | - | - | - | - |
| `answers` | - | - | - | Yes | - |
| `bypass_restricted_code` | - | - | - | - | - |
| `clearance_slip` | - | - | - | - | - |
| `client` | - | - | - | - | - |
| `company_member` | - | - | - | Yes | - |
| `consignee` | - | - | - | - | - |
| `consignor` | - | - | - | - | - |
| `created_at` | - | - | - | - | - |
| `declarant` | - | - | - | - | - |
| `document_upload_status` | - | - | - | - | - |
| `documents_presentation_requested` | - | - | - | - | - |
| `documents_upload_requested` | - | - | - | - | - |
| `external_id` | - | - | - | - | - |
| `form` | - | - | - | - | - |
| `goods_presentation_status` | - | - | - | - | - |
| `hrcm_status` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `invalidation_status` | - | - | - | - | - |
| `is_global_template` | - | - | - | - | - |
| `latest_notification_item` | - | - | - | - | - |
| `latest_state` | - | - | - | - | - |
| `lrn` | - | - | - | - | - |
| `mrn` | - | - | - | - | - |
| `name` | - | - | - | - | - |
| `partial_answers` | - | - | - | - | - |
| `receipt` | - | - | - | - | - |
| `refund_application_status` | - | - | - | - | - |
| `route` | - | - | - | - | - |
| `shipment_items_no` | - | - | - | - | - |
| `shipment_items_quantity_no` | - | - | - | - | - |
| `source` | - | - | - | - | - |
| `source_type` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `template` | - | - | - | - | - |
| `template_id` | - | - | - | - | - |
| `template_properties` | - | - | - | - | - |
| `total_tax_amount` | - | - | - | - | - |
| `updated_at` | - | - | - | - | - |
| `verification_errors` | - | - | - | - | - |
| `verification_status` | - | - | - | - | - |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Submission(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Submission(nil).Load(map[string]any{"id": "submission_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

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

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Submission(nil).Update(map[string]any{
    "id": "submission_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Submission(nil).Remove(map[string]any{"id": "submission_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SubmissionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SubmissionDetailEntity

```go
submissionDetail := client.SubmissionDetail(nil)
fmt.Println(submissionDetail.GetName()) // "submission_detail"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_external_ids` | `[]any` | No |  |
| `additional_information` | `[]any` | Yes |  |
| `amendment_status` | `string` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `clearance_slip` | `map[string]any` | Yes |  |
| `client` | `map[string]any` | Yes |  |
| `company` | `string` | Yes | cuid-format identifier for this entity. |
| `company_member` | `map[string]any` | Yes | cuid-format identifier for this entity. |
| `consignee` | `map[string]any` | Yes |  |
| `consignor` | `map[string]any` | Yes |  |
| `created_at` | `string` | Yes |  |
| `declarant` | `map[string]any` | Yes |  |
| `document_upload_status` | `string` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `documents_presentation_requested` | `bool` | No |  |
| `documents_upload_requested` | `bool` | No |  |
| `external_id` | `string` | No |  |
| `form` | `string` | Yes |  |
| `goods_presentation_status` | `string` | No | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested |
| `hrcm_status` | `string` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `id` | `string` | No | cuid-format identifier for this entity. |
| `invalidation_status` | `string` | No | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled |
| `is_global_template` | `bool` | No |  |
| `latest_notification_item` | `string` | No | cuid-format identifier for this entity. |
| `latest_state` | `map[string]any` | Yes |  |
| `lrn` | `string` | No |  |
| `mrn` | `string` | No |  |
| `name` | `string` | No |  |
| `receipt` | `map[string]any` | Yes |  |
| `refund_application_status` | `string` | No | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `route` | `string` | No | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `shipment_items_no` | `int` | No |  |
| `shipment_items_quantity_no` | `int` | No |  |
| `source` | `string` | No | cuid-format identifier for this entity. |
| `source_type` | `string` | No | * `template` - template * `automated_import` - automated_import |
| `status` | `string` | No | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `submission` | `string` | Yes | cuid-format identifier for this entity. |
| `supporting_documents` | `[]any` | No |  |
| `template` | `bool` | No |  |
| `total_tax_amount` | `string` | No |  |
| `updated_at` | `string` | Yes |  |
| `verification_errors` | `[]any` | No |  |
| `verification_status` | `string` | No | * `passed` - passed * `failed` - failed |

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

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

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SubmissionDetailEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewCustomsWindowSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

