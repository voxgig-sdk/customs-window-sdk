# CustomsWindow Ruby SDK Reference

Complete API reference for the CustomsWindow Ruby SDK.


## CustomsWindowSDK

### Constructor

```ruby
require_relative 'CustomsWindow_sdk'

client = CustomsWindowSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["apikey"]` | `String` | API key for authentication. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CustomsWindowSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = CustomsWindowSDK.test
```


### Instance Methods

#### `BulkUpload(data = nil)`

Create a new `BulkUpload` entity instance. Pass `nil` for no initial data.

#### `File(data = nil)`

Create a new `File` entity instance. Pass `nil` for no initial data.

#### `PaginatedBulkUploadListList(data = nil)`

Create a new `PaginatedBulkUploadListList` entity instance. Pass `nil` for no initial data.

#### `PaginatedPartyListList(data = nil)`

Create a new `PaginatedPartyListList` entity instance. Pass `nil` for no initial data.

#### `PaginatedSubmissionListList(data = nil)`

Create a new `PaginatedSubmissionListList` entity instance. Pass `nil` for no initial data.

#### `Party(data = nil)`

Create a new `Party` entity instance. Pass `nil` for no initial data.

#### `Submission(data = nil)`

Create a new `Submission` entity instance. Pass `nil` for no initial data.

#### `SubmissionDetail(data = nil)`

Create a new `SubmissionDetail` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## BulkUploadEntity

```ruby
bulk_upload = client.BulkUpload
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `active_transport_nationality` | `String` | No | cuid-format identifier for this entity. |
| `active_transport_number` | `String` | No |  |
| `arrival_datetime` | `String` | No |  |
| `client` | `String` | Yes | cuid-format identifier for this entity. |
| `company_member` | `String` | Yes | cuid-format identifier for this entity. |
| `created_at` | `String` | Yes |  |
| `declarant` | `String` | Yes | cuid-format identifier for this entity. |
| `declarations_no` | `Integer` | Yes |  |
| `deleted_at` | `String` | No |  |
| `departure_datetime` | `String` | No |  |
| `errors_file` | `Hash` | Yes |  |
| `external_id` | `String` | No |  |
| `failed_declarations_no` | `Integer` | No |  |
| `file` | `Hash` | Yes | cuid-format identifier for this entity. |
| `green_routed_no` | `Integer` | No |  |
| `h1_fallback_template` | `String` | No | cuid-format identifier for this entity. |
| `house_transport_doc_ref` | `String` | No |  |
| `id` | `String` | No | cuid-format identifier for this entity. |
| `issue_date` | `String` | No |  |
| `mapping` | `String` | No | cuid-format identifier for this entity. |
| `orange_routed_no` | `Integer` | No |  |
| `parsed_declarations_no` | `Integer` | No |  |
| `parser` | `String` | No | * `aes_platform` - aes_platform * `cds_platform` - cds_platform * `cds_export_platform` - cds_export_platform * `g4_g3` - g4_g3 * `nhd_platform` - nhd_platform * `platform` - platform * `birds` - birds * `ics2_platform` - ics2_platform |
| `parsing_completed_at` | `String` | No |  |
| `parsing_started_at` | `String` | No |  |
| `passive_transport_nationality` | `String` | No | cuid-format identifier for this entity. |
| `passive_transport_number` | `String` | No |  |
| `processed_declarations_no` | `Integer` | No |  |
| `processing_ended_at` | `String` | No |  |
| `processing_started_at` | `String` | No |  |
| `receipt_generating_started_at` | `String` | No |  |
| `receipt_request_started_at` | `String` | No |  |
| `receipt_request_status` | `String` | No | * `pending` - pending * `processing` - processing * `generating` - generating * `completed` - completed |
| `receipt_request_user` | `String` | No | cuid-format identifier for this entity. |
| `receipts_zip` | `String` | No | cuid-format identifier for this entity. |
| `red_routed_no` | `Integer` | No |  |
| `rejected_status_no` | `Integer` | No |  |
| `status` | `String` | No | * `pending` - pending * `parsing` - parsing * `processing` - processing * `complete` - complete |
| `template` | `String` | No | cuid-format identifier for this entity. |
| `updated_at` | `String` | Yes |  |
| `yellow_routed_no` | `Integer` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.BulkUpload.create({
  "client" => "example_client", # String
  "company_member" => "example_company_member", # String
  "created_at" => "example_created_at", # String
  "declarant" => "example_declarant", # String
  "declarations_no" => 1, # Integer
  "errors_file" => {}, # Hash
  "file" => {}, # Hash
  "updated_at" => "example_updated_at", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.BulkUpload.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.BulkUpload.load({ "id" => "bulk_upload_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.BulkUpload.remove({ "id" => "bulk_upload_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.BulkUpload.update({
  "id" => "bulk_upload_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `BulkUploadEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## FileEntity

```ruby
file = client.File
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `company` | `String` | Yes | cuid-format identifier for this entity. |
| `created_at` | `String` | Yes |  |
| `extension` | `String` | Yes |  |
| `file` | `String` | Yes |  |
| `id` | `String` | No | cuid-format identifier for this entity. |
| `name` | `String` | Yes |  |
| `public` | `Boolean` | Yes |  |
| `size` | `Integer` | No |  |
| `updated_at` | `String` | Yes |  |
| `url` | `String` | Yes |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.File.create({
  "company" => "example_company", # String
  "created_at" => "example_created_at", # String
  "extension" => "example_extension", # String
  "file" => "example_file", # String
  "name" => "example_name", # String
  "public" => true, # Boolean
  "updated_at" => "example_updated_at", # String
  "url" => "example_url", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `FileEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PaginatedBulkUploadListListEntity

```ruby
paginated_bulk_upload_list_list = client.PaginatedBulkUploadListList
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PaginatedBulkUploadListListEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PaginatedPartyListListEntity

```ruby
paginated_party_list_list = client.PaginatedPartyListList
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PaginatedPartyListListEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PaginatedSubmissionListListEntity

```ruby
paginated_submission_list_list = client.PaginatedSubmissionListList
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PaginatedSubmissionListListEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PartyEntity

```ruby
party = client.Party
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_declaration_type` | `Hash` | Yes |  |
| `address` | `Hash` | Yes |  |
| `authorisation` | `Hash` | Yes |  |
| `bank_details` | `String` | No |  |
| `certificate` | `Hash` | Yes |  |
| `certificate_type` | `String` | Yes |  |
| `company` | `String` | Yes | cuid-format identifier for this entity. |
| `created_at` | `String` | Yes |  |
| `customs_office_of_lodgement` | `Hash` | Yes |  |
| `deleted_at` | `String` | No |  |
| `email` | `String` | No |  |
| `id` | `String` | No | cuid-format identifier for this entity. |
| `identification_number` | `String` | No |  |
| `indirect_representative` | `Boolean` | No |  |
| `name` | `String` | No |  |
| `nhd_last_submission_year` | `Integer` | No |  |
| `nhd_submission_counter` | `Integer` | No |  |
| `person_paying_customs_duty` | `String` | No |  |
| `phone_country_code` | `String` | No |  |
| `phone_number` | `String` | No |  |
| `preferred_payment_method` | `Hash` | Yes |  |
| `signed_form` | `Hash` | Yes |  |
| `type` | `String` | No | * `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `co… |
| `type_of_person` | `Hash` | Yes |  |
| `unlocode` | `String` | No |  |
| `updated_at` | `String` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Party.create({
  "additional_declaration_type" => {}, # Hash
  "address" => {}, # Hash
  "authorisation" => {}, # Hash
  "certificate" => {}, # Hash
  "certificate_type" => "example_certificate_type", # String
  "company" => "example_company", # String
  "created_at" => "example_created_at", # String
  "customs_office_of_lodgement" => {}, # Hash
  "preferred_payment_method" => {}, # Hash
  "signed_form" => {}, # Hash
  "type_of_person" => {}, # Hash
  "updated_at" => "example_updated_at", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Party.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Party.load({ "id" => "party_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Party.remove({ "id" => "party_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Party.update({
  "id" => "party_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PartyEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SubmissionEntity

```ruby
submission = client.Submission
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_external_ids` | `Array` | No |  |
| `amendment_reason` | `String` | No |  |
| `amendment_status` | `String` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `answers` | `Array` | Yes |  |
| `bypass_restricted_code` | `Boolean` | No |  |
| `clearance_slip` | `Hash` | Yes |  |
| `client` | `Hash` | Yes |  |
| `company_member` | `Hash` | Yes | cuid-format identifier for this entity. |
| `consignee` | `Hash` | Yes |  |
| `consignor` | `Hash` | Yes |  |
| `created_at` | `String` | Yes |  |
| `declarant` | `Hash` | Yes | cuid-format identifier for this entity. |
| `document_upload_status` | `String` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `documents_presentation_requested` | `Boolean` | No |  |
| `documents_upload_requested` | `Boolean` | No |  |
| `external_id` | `String` | No |  |
| `form` | `String` | Yes | cuid-format identifier for this entity. |
| `goods_presentation_status` | `String` | No | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested |
| `hrcm_status` | `String` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `id` | `String` | No | cuid-format identifier for this entity. |
| `invalidation_status` | `String` | No | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled |
| `is_global_template` | `Boolean` | No |  |
| `latest_notification_item` | `String` | No | cuid-format identifier for this entity. |
| `latest_state` | `Hash` | Yes |  |
| `lrn` | `String` | No |  |
| `mrn` | `String` | No |  |
| `name` | `String` | No |  |
| `partial_answers` | `Boolean` | No |  |
| `receipt` | `Hash` | Yes |  |
| `refund_application_status` | `String` | No | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `route` | `String` | No | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `shipment_items_no` | `Integer` | No |  |
| `shipment_items_quantity_no` | `Integer` | No |  |
| `source` | `String` | No | cuid-format identifier for this entity. |
| `source_type` | `String` | No | * `template` - template * `automated_import` - automated_import |
| `status` | `String` | No | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `template` | `Boolean` | No |  |
| `template_id` | `String` | No | cuid-format identifier for this entity. |
| `template_properties` | `Array` | No |  |
| `total_tax_amount` | `String` | No |  |
| `updated_at` | `String` | Yes |  |
| `verification_errors` | `Array` | No |  |
| `verification_status` | `String` | No | * `passed` - passed * `failed` - failed |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Submission.create({
  "answers" => [], # Array
  "clearance_slip" => {}, # Hash
  "client" => {}, # Hash
  "company_member" => {}, # Hash
  "consignee" => {}, # Hash
  "consignor" => {}, # Hash
  "created_at" => "example_created_at", # String
  "declarant" => {}, # Hash
  "form" => "example_form", # String
  "latest_state" => {}, # Hash
  "receipt" => {}, # Hash
  "updated_at" => "example_updated_at", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Submission.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Submission.load({ "id" => "submission_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Submission.remove({ "id" => "submission_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Submission.update({
  "id" => "submission_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SubmissionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SubmissionDetailEntity

```ruby
submission_detail = client.SubmissionDetail
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_external_ids` | `Array` | No |  |
| `additional_information` | `Array` | Yes |  |
| `amendment_status` | `String` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `clearance_slip` | `Hash` | Yes |  |
| `client` | `Hash` | Yes |  |
| `company` | `String` | Yes | cuid-format identifier for this entity. |
| `company_member` | `Hash` | Yes | cuid-format identifier for this entity. |
| `consignee` | `Hash` | Yes |  |
| `consignor` | `Hash` | Yes |  |
| `created_at` | `String` | Yes |  |
| `declarant` | `Hash` | Yes |  |
| `document_upload_status` | `String` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `documents_presentation_requested` | `Boolean` | No |  |
| `documents_upload_requested` | `Boolean` | No |  |
| `external_id` | `String` | No |  |
| `form` | `String` | Yes |  |
| `goods_presentation_status` | `String` | No | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested |
| `hrcm_status` | `String` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `id` | `String` | No | cuid-format identifier for this entity. |
| `invalidation_status` | `String` | No | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled |
| `is_global_template` | `Boolean` | No |  |
| `latest_notification_item` | `String` | No | cuid-format identifier for this entity. |
| `latest_state` | `Hash` | Yes |  |
| `lrn` | `String` | No |  |
| `mrn` | `String` | No |  |
| `name` | `String` | No |  |
| `receipt` | `Hash` | Yes |  |
| `refund_application_status` | `String` | No | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `route` | `String` | No | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `shipment_items_no` | `Integer` | No |  |
| `shipment_items_quantity_no` | `Integer` | No |  |
| `source` | `String` | No | cuid-format identifier for this entity. |
| `source_type` | `String` | No | * `template` - template * `automated_import` - automated_import |
| `status` | `String` | No | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `submission` | `String` | Yes | cuid-format identifier for this entity. |
| `supporting_documents` | `Array` | No |  |
| `template` | `Boolean` | No |  |
| `total_tax_amount` | `String` | No |  |
| `updated_at` | `String` | Yes |  |
| `verification_errors` | `Array` | No |  |
| `verification_status` | `String` | No | * `passed` - passed * `failed` - failed |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.SubmissionDetail.create({
  "additional_information" => [], # Array
  "clearance_slip" => {}, # Hash
  "client" => {}, # Hash
  "company" => "example_company", # String
  "company_member" => {}, # Hash
  "consignee" => {}, # Hash
  "consignor" => {}, # Hash
  "created_at" => "example_created_at", # String
  "declarant" => {}, # Hash
  "form" => "example_form", # String
  "latest_state" => {}, # Hash
  "receipt" => {}, # Hash
  "submission" => "example_submission", # String
  "updated_at" => "example_updated_at", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SubmissionDetailEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = CustomsWindowSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```

