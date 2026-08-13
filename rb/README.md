# CustomsWindow Ruby SDK



The Ruby SDK for the CustomsWindow API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.BulkUpload` — with named operations (`list`/`load`/`create`/`update`/`remove`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/customs-window-sdk/releases](https://github.com/voxgig-sdk/customs-window-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "CustomsWindow_sdk"

client = CustomsWindowSDK.new({
  "apikey" => ENV["CUSTOMS_WINDOW_APIKEY"],
})
```

### 2. List bulkupload records

```ruby
begin
  # list returns an Array of BulkUpload records — iterate directly.
  bulkuploads = client.BulkUpload.list
  bulkuploads.each do |item|
    puts "#{item["id"]} #{item["active_transport_nationality"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```

### 3. Load a bulkupload

```ruby
begin
  # load returns the ENTITY — call data_get for the BulkUpload record (raises on error).
  bulkupload = client.BulkUpload.load({ "id" => "example_id" })
  puts bulkupload
rescue => err
  warn "load failed: #{err}"
end
```

### 4. Create, update, and remove

```ruby
# create returns the ENTITY — call data_get for the created BulkUpload record.
created = client.BulkUpload.create({ "client" => "example_client", "company_member" => "example_company_member", "created_at" => "example_created_at", "declarant" => "example_declarant", "declarations_no" => 1, "errors_file" => {}, "file" => {}, "updated_at" => "example_updated_at" })

# Update — index the record via data_get (created.data_get["id"]).
client.BulkUpload.update({ "id" => created.data_get["id"], "active_transport_nationality" => "example_active_transport_nationality", "active_transport_number" => "example_active_transport_number" })

# Remove
client.BulkUpload.remove({ "id" => created.data_get["id"] })
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  submissions = client.Submission.list()
rescue => err
  warn "list failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```ruby
client = CustomsWindowSDK.test({
  "entity" => { "submission" => { "test01" => { "id" => "test01" } } },
})

# Entity ops return the ENTITY (raises on error);
# call data_get for the mock record.
submission = client.Submission.list()
puts submission
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = CustomsWindowSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
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
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### CustomsWindowSDK

```ruby
require_relative "CustomsWindow_sdk"
client = CustomsWindowSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `String` | API key for authentication. |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = CustomsWindowSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### CustomsWindowSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
| `BulkUpload` | `(data) -> BulkUploadEntity` | Create a BulkUpload entity instance. |
| `File` | `(data) -> FileEntity` | Create a File entity instance. |
| `PaginatedBulkUploadListList` | `(data) -> PaginatedBulkUploadListListEntity` | Create a PaginatedBulkUploadListList entity instance. |
| `PaginatedPartyListList` | `(data) -> PaginatedPartyListListEntity` | Create a PaginatedPartyListList entity instance. |
| `PaginatedSubmissionListList` | `(data) -> PaginatedSubmissionListListEntity` | Create a PaginatedSubmissionListList entity instance. |
| `Party` | `(data) -> PartyEntity` | Create a Party entity instance. |
| `Submission` | `(data) -> SubmissionEntity` | Create a Submission entity instance. |
| `SubmissionDetail` | `(data) -> SubmissionDetailEntity` | Create a SubmissionDetail entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch = nil, ctrl) -> Array` | List entities matching the criteria (call with no argument to list all). Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `CustomsWindowError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

### Entities

#### BulkUpload

| Field | Description |
| --- | --- |
| `active_transport_nationality` |  |
| `active_transport_number` |  |
| `arrival_datetime` |  |
| `client` |  |
| `company_member` |  |
| `created_at` |  |
| `declarant` |  |
| `declarations_no` |  |
| `deleted_at` |  |
| `departure_datetime` |  |
| `errors_file` |  |
| `external_id` |  |
| `failed_declarations_no` |  |
| `file` |  |
| `green_routed_no` |  |
| `h1_fallback_template` |  |
| `house_transport_doc_ref` |  |
| `id` |  |
| `issue_date` |  |
| `mapping` |  |
| `orange_routed_no` |  |
| `parsed_declarations_no` |  |
| `parser` |  |
| `parsing_completed_at` |  |
| `parsing_started_at` |  |
| `passive_transport_nationality` |  |
| `passive_transport_number` |  |
| `processed_declarations_no` |  |
| `processing_ended_at` |  |
| `processing_started_at` |  |
| `receipt_generating_started_at` |  |
| `receipt_request_started_at` |  |
| `receipt_request_status` |  |
| `receipt_request_user` |  |
| `receipts_zip` |  |
| `red_routed_no` |  |
| `rejected_status_no` |  |
| `status` |  |
| `template` |  |
| `updated_at` |  |
| `yellow_routed_no` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/bulk-uploads`

#### File

| Field | Description |
| --- | --- |
| `company` |  |
| `created_at` |  |
| `extension` |  |
| `file` |  |
| `id` |  |
| `name` |  |
| `public` |  |
| `size` |  |
| `updated_at` |  |
| `url` |  |

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
| `additional_declaration_type` |  |
| `address` |  |
| `authorisation` |  |
| `bank_details` |  |
| `certificate` |  |
| `certificate_type` |  |
| `company` |  |
| `created_at` |  |
| `customs_office_of_lodgement` |  |
| `deleted_at` |  |
| `email` |  |
| `id` |  |
| `identification_number` |  |
| `indirect_representative` |  |
| `name` |  |
| `nhd_last_submission_year` |  |
| `nhd_submission_counter` |  |
| `person_paying_customs_duty` |  |
| `phone_country_code` |  |
| `phone_number` |  |
| `preferred_payment_method` |  |
| `signed_form` |  |
| `type` |  |
| `type_of_person` |  |
| `unlocode` |  |
| `updated_at` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/parties`

#### Submission

| Field | Description |
| --- | --- |
| `additional_external_ids` |  |
| `amendment_reason` |  |
| `amendment_status` |  |
| `answers` |  |
| `bypass_restricted_code` |  |
| `clearance_slip` |  |
| `client` |  |
| `company_member` |  |
| `consignee` |  |
| `consignor` |  |
| `created_at` |  |
| `declarant` |  |
| `document_upload_status` |  |
| `documents_presentation_requested` |  |
| `documents_upload_requested` |  |
| `external_id` |  |
| `form` |  |
| `goods_presentation_status` |  |
| `hrcm_status` |  |
| `id` |  |
| `invalidation_status` |  |
| `is_global_template` |  |
| `latest_notification_item` |  |
| `latest_state` |  |
| `lrn` |  |
| `mrn` |  |
| `name` |  |
| `partial_answers` |  |
| `receipt` |  |
| `refund_application_status` |  |
| `route` |  |
| `shipment_items_no` |  |
| `shipment_items_quantity_no` |  |
| `source` |  |
| `source_type` |  |
| `status` |  |
| `template` |  |
| `template_id` |  |
| `template_properties` |  |
| `total_tax_amount` |  |
| `updated_at` |  |
| `verification_errors` |  |
| `verification_status` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/submissions/{id}/refund`

#### SubmissionDetail

| Field | Description |
| --- | --- |
| `additional_external_ids` |  |
| `additional_information` |  |
| `amendment_status` |  |
| `clearance_slip` |  |
| `client` |  |
| `company` |  |
| `company_member` |  |
| `consignee` |  |
| `consignor` |  |
| `created_at` |  |
| `declarant` |  |
| `document_upload_status` |  |
| `documents_presentation_requested` |  |
| `documents_upload_requested` |  |
| `external_id` |  |
| `form` |  |
| `goods_presentation_status` |  |
| `hrcm_status` |  |
| `id` |  |
| `invalidation_status` |  |
| `is_global_template` |  |
| `latest_notification_item` |  |
| `latest_state` |  |
| `lrn` |  |
| `mrn` |  |
| `name` |  |
| `receipt` |  |
| `refund_application_status` |  |
| `route` |  |
| `shipment_items_no` |  |
| `shipment_items_quantity_no` |  |
| `source` |  |
| `source_type` |  |
| `status` |  |
| `submission` |  |
| `supporting_documents` |  |
| `template` |  |
| `total_tax_amount` |  |
| `updated_at` |  |
| `verification_errors` |  |
| `verification_status` |  |

Operations: Create.

API path: `/documents-request`



## Entities


### BulkUpload

Create an instance: `bulk_upload = client.BulkUpload`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `active_transport_nationality` | `String` |  |
| `active_transport_number` | `String` |  |
| `arrival_datetime` | `String` |  |
| `client` | `String` |  |
| `company_member` | `String` |  |
| `created_at` | `String` |  |
| `declarant` | `String` |  |
| `declarations_no` | `Integer` |  |
| `deleted_at` | `String` |  |
| `departure_datetime` | `String` |  |
| `errors_file` | `Hash` |  |
| `external_id` | `String` |  |
| `failed_declarations_no` | `Integer` |  |
| `file` | `Hash` |  |
| `green_routed_no` | `Integer` |  |
| `h1_fallback_template` | `String` |  |
| `house_transport_doc_ref` | `String` |  |
| `id` | `String` |  |
| `issue_date` | `String` |  |
| `mapping` | `String` |  |
| `orange_routed_no` | `Integer` |  |
| `parsed_declarations_no` | `Integer` |  |
| `parser` | `String` |  |
| `parsing_completed_at` | `String` |  |
| `parsing_started_at` | `String` |  |
| `passive_transport_nationality` | `String` |  |
| `passive_transport_number` | `String` |  |
| `processed_declarations_no` | `Integer` |  |
| `processing_ended_at` | `String` |  |
| `processing_started_at` | `String` |  |
| `receipt_generating_started_at` | `String` |  |
| `receipt_request_started_at` | `String` |  |
| `receipt_request_status` | `String` |  |
| `receipt_request_user` | `String` |  |
| `receipts_zip` | `String` |  |
| `red_routed_no` | `Integer` |  |
| `rejected_status_no` | `Integer` |  |
| `status` | `String` |  |
| `template` | `String` |  |
| `updated_at` | `String` |  |
| `yellow_routed_no` | `Integer` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the BulkUpload record (raises on error).
bulk_upload = client.BulkUpload.load({ "id" => "bulk_upload_id" })
```

#### Example: List

```ruby
# list returns an Array of BulkUpload records (raises on error).
bulk_uploads = client.BulkUpload.list
```

#### Example: Create

```ruby
bulk_upload = client.BulkUpload.create({
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


### File

Create an instance: `file = client.File`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `company` | `String` |  |
| `created_at` | `String` |  |
| `extension` | `String` |  |
| `file` | `String` |  |
| `id` | `String` |  |
| `name` | `String` |  |
| `public` | `Boolean` |  |
| `size` | `Integer` |  |
| `updated_at` | `String` |  |
| `url` | `String` |  |

#### Example: Create

```ruby
file = client.File.create({
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


### PaginatedBulkUploadListList

Create an instance: `paginated_bulk_upload_list_list = client.PaginatedBulkUploadListList`


### PaginatedPartyListList

Create an instance: `paginated_party_list_list = client.PaginatedPartyListList`


### PaginatedSubmissionListList

Create an instance: `paginated_submission_list_list = client.PaginatedSubmissionListList`


### Party

Create an instance: `party = client.Party`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_declaration_type` | `Hash` |  |
| `address` | `Hash` |  |
| `authorisation` | `Hash` |  |
| `bank_details` | `String` |  |
| `certificate` | `Hash` |  |
| `certificate_type` | `String` |  |
| `company` | `String` |  |
| `created_at` | `String` |  |
| `customs_office_of_lodgement` | `Hash` |  |
| `deleted_at` | `String` |  |
| `email` | `String` |  |
| `id` | `String` |  |
| `identification_number` | `String` |  |
| `indirect_representative` | `Boolean` |  |
| `name` | `String` |  |
| `nhd_last_submission_year` | `Integer` |  |
| `nhd_submission_counter` | `Integer` |  |
| `person_paying_customs_duty` | `String` |  |
| `phone_country_code` | `String` |  |
| `phone_number` | `String` |  |
| `preferred_payment_method` | `Hash` |  |
| `signed_form` | `Hash` |  |
| `type` | `String` |  |
| `type_of_person` | `Hash` |  |
| `unlocode` | `String` |  |
| `updated_at` | `String` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Party record (raises on error).
party = client.Party.load({ "id" => "party_id" })
```

#### Example: List

```ruby
# list returns an Array of Party records (raises on error).
partys = client.Party.list
```

#### Example: Create

```ruby
party = client.Party.create({
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


### Submission

Create an instance: `submission = client.Submission`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_external_ids` | `Array` |  |
| `amendment_reason` | `String` |  |
| `amendment_status` | `String` |  |
| `answers` | `Array` |  |
| `bypass_restricted_code` | `Boolean` |  |
| `clearance_slip` | `Hash` |  |
| `client` | `Hash` |  |
| `company_member` | `Hash` |  |
| `consignee` | `Hash` |  |
| `consignor` | `Hash` |  |
| `created_at` | `String` |  |
| `declarant` | `Hash` |  |
| `document_upload_status` | `String` |  |
| `documents_presentation_requested` | `Boolean` |  |
| `documents_upload_requested` | `Boolean` |  |
| `external_id` | `String` |  |
| `form` | `String` |  |
| `goods_presentation_status` | `String` |  |
| `hrcm_status` | `String` |  |
| `id` | `String` |  |
| `invalidation_status` | `String` |  |
| `is_global_template` | `Boolean` |  |
| `latest_notification_item` | `String` |  |
| `latest_state` | `Hash` |  |
| `lrn` | `String` |  |
| `mrn` | `String` |  |
| `name` | `String` |  |
| `partial_answers` | `Boolean` |  |
| `receipt` | `Hash` |  |
| `refund_application_status` | `String` |  |
| `route` | `String` |  |
| `shipment_items_no` | `Integer` |  |
| `shipment_items_quantity_no` | `Integer` |  |
| `source` | `String` |  |
| `source_type` | `String` |  |
| `status` | `String` |  |
| `template` | `Boolean` |  |
| `template_id` | `String` |  |
| `template_properties` | `Array` |  |
| `total_tax_amount` | `String` |  |
| `updated_at` | `String` |  |
| `verification_errors` | `Array` |  |
| `verification_status` | `String` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Submission record (raises on error).
submission = client.Submission.load({ "id" => "submission_id" })
```

#### Example: List

```ruby
# list returns an Array of Submission records (raises on error).
submissions = client.Submission.list
```

#### Example: Create

```ruby
submission = client.Submission.create({
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


### SubmissionDetail

Create an instance: `submission_detail = client.SubmissionDetail`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_external_ids` | `Array` |  |
| `additional_information` | `Array` |  |
| `amendment_status` | `String` |  |
| `clearance_slip` | `Hash` |  |
| `client` | `Hash` |  |
| `company` | `String` |  |
| `company_member` | `Hash` |  |
| `consignee` | `Hash` |  |
| `consignor` | `Hash` |  |
| `created_at` | `String` |  |
| `declarant` | `Hash` |  |
| `document_upload_status` | `String` |  |
| `documents_presentation_requested` | `Boolean` |  |
| `documents_upload_requested` | `Boolean` |  |
| `external_id` | `String` |  |
| `form` | `String` |  |
| `goods_presentation_status` | `String` |  |
| `hrcm_status` | `String` |  |
| `id` | `String` |  |
| `invalidation_status` | `String` |  |
| `is_global_template` | `Boolean` |  |
| `latest_notification_item` | `String` |  |
| `latest_state` | `Hash` |  |
| `lrn` | `String` |  |
| `mrn` | `String` |  |
| `name` | `String` |  |
| `receipt` | `Hash` |  |
| `refund_application_status` | `String` |  |
| `route` | `String` |  |
| `shipment_items_no` | `Integer` |  |
| `shipment_items_quantity_no` | `Integer` |  |
| `source` | `String` |  |
| `source_type` | `String` |  |
| `status` | `String` |  |
| `submission` | `String` |  |
| `supporting_documents` | `Array` |  |
| `template` | `Boolean` |  |
| `total_tax_amount` | `String` |  |
| `updated_at` | `String` |  |
| `verification_errors` | `Array` |  |
| `verification_status` | `String` |  |

#### Example: Create

```ruby
submission_detail = client.SubmissionDetail.create({
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

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── CustomsWindow_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`CustomsWindow_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```ruby
submission = client.Submission
submission.list()

# submission.data_get now returns the submission data from the last list
# submission.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
