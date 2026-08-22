# CustomsWindow Lua SDK



The Lua SDK for the CustomsWindow API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:BulkUpload()` — each with the same small set of operations (`list`, `load`, `create`, `update`, `remove`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/customs-window-sdk/releases)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("customs-window_sdk")

local client = sdk.new({
  apikey = os.getenv("CUSTOMS_WINDOW_APIKEY"),
})
```

### 2. List bulkupload records

Entity operations return `(value, err)`. For `list`, `value` is the
array of records itself — iterate it directly (there is no wrapper).

```lua
local bulkuploads, err = client:BulkUpload():list()
if err then error(err) end

for _, item in ipairs(bulkuploads) do
  print(item["id"], item["active_transport_nationality"])
end
```

### 3. Load a bulkupload

```lua
local bulkupload, err = client:BulkUpload():load({ id = "example_id" })
if err then error(err) end
print(bulkupload)
```

### 4. Create, update, and remove

```lua
-- Create
local created, err = client:BulkUpload():create({ client = "example_client", company_member = "example_company_member", created_at = "example_created_at", declarant = "example_declarant", declarations_no = 1, errors_file = {}, file = {}, updated_at = "example_updated_at" })
if err then error(err) end

-- Update
client:BulkUpload():update({ id = created:data_get()["id"], active_transport_nationality = "example_active_transport_nationality", active_transport_number = "example_active_transport_number" })

-- Remove
client:BulkUpload():remove({ id = created:data_get()["id"] })
```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local submissions, err = client:Submission():list()
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:Submission():list()
-- result is the returned data; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
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
cd lua && busted test/
```


## Reference

### CustomsWindowSDK

```lua
local sdk = require("customs-window_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### CustomsWindowSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
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
| `load` | `(reqmatch, ctrl) -> any, err` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> any, err` | List entities matching the criteria. |
| `create` | `(reqdata, ctrl) -> any, err` | Create a new entity. |
| `update` | `(reqdata, ctrl) -> any, err` | Update an existing entity. |
| `remove` | `(reqmatch, ctrl) -> any, err` | Remove an entity. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the operation's
data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `load` / `create` / `update` / `remove` | the entity record (a `table`) |
| `list` | an array (`table`) of entity records |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local bulk_upload, err = client:BulkUpload():load({ id = "example_id" })
    if err then error(err) end
    -- bulk_upload is the loaded record

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

### Entities

#### BulkUpload

| Field | Description |
| --- | --- |
| `active_transport_nationality` | cuid-format identifier for this entity. |
| `active_transport_number` |  |
| `arrival_datetime` |  |
| `client` | cuid-format identifier for this entity. |
| `company_member` | cuid-format identifier for this entity. |
| `created_at` |  |
| `declarant` | cuid-format identifier for this entity. |
| `declarations_no` |  |
| `deleted_at` |  |
| `departure_datetime` |  |
| `errors_file` |  |
| `external_id` |  |
| `failed_declarations_no` |  |
| `file` | cuid-format identifier for this entity. |
| `green_routed_no` |  |
| `h1_fallback_template` | cuid-format identifier for this entity. |
| `house_transport_doc_ref` |  |
| `id` | cuid-format identifier for this entity. |
| `issue_date` |  |
| `mapping` | cuid-format identifier for this entity. |
| `orange_routed_no` |  |
| `parsed_declarations_no` |  |
| `parser` | * `aes_platform` - aes_platform * `cds_platform` - cds_platform * `cds_export_platform` - cds_export_platform * `g4_g3` - g4_g3 * `nhd_platform` - nhd_platform * `platform` - platform * `birds` - birds * `ics2_platform` - ics2_platform |
| `parsing_completed_at` |  |
| `parsing_started_at` |  |
| `passive_transport_nationality` | cuid-format identifier for this entity. |
| `passive_transport_number` |  |
| `processed_declarations_no` |  |
| `processing_ended_at` |  |
| `processing_started_at` |  |
| `receipt_generating_started_at` |  |
| `receipt_request_started_at` |  |
| `receipt_request_status` | * `pending` - pending * `processing` - processing * `generating` - generating * `completed` - completed |
| `receipt_request_user` | cuid-format identifier for this entity. |
| `receipts_zip` | cuid-format identifier for this entity. |
| `red_routed_no` |  |
| `rejected_status_no` |  |
| `status` | * `pending` - pending * `parsing` - parsing * `processing` - processing * `complete` - complete |
| `template` | cuid-format identifier for this entity. |
| `updated_at` |  |
| `yellow_routed_no` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/bulk-uploads`

#### File

| Field | Description |
| --- | --- |
| `company` | cuid-format identifier for this entity. |
| `created_at` |  |
| `extension` |  |
| `file` |  |
| `id` | cuid-format identifier for this entity. |
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
| `company` | cuid-format identifier for this entity. |
| `created_at` |  |
| `customs_office_of_lodgement` |  |
| `deleted_at` |  |
| `email` |  |
| `id` | cuid-format identifier for this entity. |
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
| `type` | * `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `co… |
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
| `amendment_status` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `answers` |  |
| `bypass_restricted_code` |  |
| `clearance_slip` |  |
| `client` |  |
| `company_member` | cuid-format identifier for this entity. |
| `consignee` |  |
| `consignor` |  |
| `created_at` |  |
| `declarant` | cuid-format identifier for this entity. |
| `document_upload_status` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `documents_presentation_requested` |  |
| `documents_upload_requested` |  |
| `external_id` |  |
| `form` | cuid-format identifier for this entity. |
| `goods_presentation_status` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested |
| `hrcm_status` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `id` | cuid-format identifier for this entity. |
| `invalidation_status` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled |
| `is_global_template` |  |
| `latest_notification_item` | cuid-format identifier for this entity. |
| `latest_state` |  |
| `lrn` |  |
| `mrn` |  |
| `name` |  |
| `partial_answers` |  |
| `receipt` |  |
| `refund_application_status` | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `route` | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `shipment_items_no` |  |
| `shipment_items_quantity_no` |  |
| `source` | cuid-format identifier for this entity. |
| `source_type` | * `template` - template * `automated_import` - automated_import |
| `status` | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `template` |  |
| `template_id` | cuid-format identifier for this entity. |
| `template_properties` |  |
| `total_tax_amount` |  |
| `updated_at` |  |
| `verification_errors` |  |
| `verification_status` | * `passed` - passed * `failed` - failed |

Operations: Create, List, Load, Remove, Update.

API path: `/submissions/{id}/refund`

#### SubmissionDetail

| Field | Description |
| --- | --- |
| `additional_external_ids` |  |
| `additional_information` |  |
| `amendment_status` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `clearance_slip` |  |
| `client` |  |
| `company` | cuid-format identifier for this entity. |
| `company_member` | cuid-format identifier for this entity. |
| `consignee` |  |
| `consignor` |  |
| `created_at` |  |
| `declarant` |  |
| `document_upload_status` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `documents_presentation_requested` |  |
| `documents_upload_requested` |  |
| `external_id` |  |
| `form` |  |
| `goods_presentation_status` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested |
| `hrcm_status` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `id` | cuid-format identifier for this entity. |
| `invalidation_status` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled |
| `is_global_template` |  |
| `latest_notification_item` | cuid-format identifier for this entity. |
| `latest_state` |  |
| `lrn` |  |
| `mrn` |  |
| `name` |  |
| `receipt` |  |
| `refund_application_status` | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `route` | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `shipment_items_no` |  |
| `shipment_items_quantity_no` |  |
| `source` | cuid-format identifier for this entity. |
| `source_type` | * `template` - template * `automated_import` - automated_import |
| `status` | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `submission` | cuid-format identifier for this entity. |
| `supporting_documents` |  |
| `template` |  |
| `total_tax_amount` |  |
| `updated_at` |  |
| `verification_errors` |  |
| `verification_status` | * `passed` - passed * `failed` - failed |

Operations: Create.

API path: `/documents-request`



## Entities


### BulkUpload

Create an instance: `local bulk_upload = client:BulkUpload(nil)`

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
| `active_transport_nationality` | `string` | cuid-format identifier for this entity. |
| `active_transport_number` | `string` |  |
| `arrival_datetime` | `string` |  |
| `client` | `string` | cuid-format identifier for this entity. |
| `company_member` | `string` | cuid-format identifier for this entity. |
| `created_at` | `string` |  |
| `declarant` | `string` | cuid-format identifier for this entity. |
| `declarations_no` | `number` |  |
| `deleted_at` | `string` |  |
| `departure_datetime` | `string` |  |
| `errors_file` | `table` |  |
| `external_id` | `string` |  |
| `failed_declarations_no` | `number` |  |
| `file` | `table` | cuid-format identifier for this entity. |
| `green_routed_no` | `number` |  |
| `h1_fallback_template` | `string` | cuid-format identifier for this entity. |
| `house_transport_doc_ref` | `string` |  |
| `id` | `string` | cuid-format identifier for this entity. |
| `issue_date` | `string` |  |
| `mapping` | `string` | cuid-format identifier for this entity. |
| `orange_routed_no` | `number` |  |
| `parsed_declarations_no` | `number` |  |
| `parser` | `string` | * `aes_platform` - aes_platform * `cds_platform` - cds_platform * `cds_export_platform` - cds_export_platform * `g4_g3` - g4_g3 * `nhd_platform` - nhd_platform * `platform` - platform * `birds` - birds * `ics2_platform` - ics2_platform |
| `parsing_completed_at` | `string` |  |
| `parsing_started_at` | `string` |  |
| `passive_transport_nationality` | `string` | cuid-format identifier for this entity. |
| `passive_transport_number` | `string` |  |
| `processed_declarations_no` | `number` |  |
| `processing_ended_at` | `string` |  |
| `processing_started_at` | `string` |  |
| `receipt_generating_started_at` | `string` |  |
| `receipt_request_started_at` | `string` |  |
| `receipt_request_status` | `string` | * `pending` - pending * `processing` - processing * `generating` - generating * `completed` - completed |
| `receipt_request_user` | `string` | cuid-format identifier for this entity. |
| `receipts_zip` | `string` | cuid-format identifier for this entity. |
| `red_routed_no` | `number` |  |
| `rejected_status_no` | `number` |  |
| `status` | `string` | * `pending` - pending * `parsing` - parsing * `processing` - processing * `complete` - complete |
| `template` | `string` | cuid-format identifier for this entity. |
| `updated_at` | `string` |  |
| `yellow_routed_no` | `number` |  |

#### Example: Load

```lua
local bulk_upload, err = client:BulkUpload():load({ id = "bulk_upload_id" })
```

#### Example: List

```lua
local bulk_uploads, err = client:BulkUpload():list()
```

#### Example: Create

```lua
local bulk_upload, err = client:BulkUpload():create({
  client = "example_client", -- string
  company_member = "example_company_member", -- string
  created_at = "example_created_at", -- string
  declarant = "example_declarant", -- string
  declarations_no = 1, -- number
  errors_file = {}, -- table
  file = {}, -- table
  updated_at = "example_updated_at", -- string
})
```


### File

Create an instance: `local file = client:File(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `company` | `string` | cuid-format identifier for this entity. |
| `created_at` | `string` |  |
| `extension` | `string` |  |
| `file` | `string` |  |
| `id` | `string` | cuid-format identifier for this entity. |
| `name` | `string` |  |
| `public` | `boolean` |  |
| `size` | `number` |  |
| `updated_at` | `string` |  |
| `url` | `string` |  |

#### Example: Create

```lua
local file, err = client:File():create({
  company = "example_company", -- string
  created_at = "example_created_at", -- string
  extension = "example_extension", -- string
  file = "example_file", -- string
  name = "example_name", -- string
  public = true, -- boolean
  updated_at = "example_updated_at", -- string
  url = "example_url", -- string
})
```


### PaginatedBulkUploadListList

Create an instance: `local paginated_bulk_upload_list_list = client:PaginatedBulkUploadListList(nil)`


### PaginatedPartyListList

Create an instance: `local paginated_party_list_list = client:PaginatedPartyListList(nil)`


### PaginatedSubmissionListList

Create an instance: `local paginated_submission_list_list = client:PaginatedSubmissionListList(nil)`


### Party

Create an instance: `local party = client:Party(nil)`

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
| `additional_declaration_type` | `table` |  |
| `address` | `table` |  |
| `authorisation` | `table` |  |
| `bank_details` | `string` |  |
| `certificate` | `table` |  |
| `certificate_type` | `string` |  |
| `company` | `string` | cuid-format identifier for this entity. |
| `created_at` | `string` |  |
| `customs_office_of_lodgement` | `table` |  |
| `deleted_at` | `string` |  |
| `email` | `string` |  |
| `id` | `string` | cuid-format identifier for this entity. |
| `identification_number` | `string` |  |
| `indirect_representative` | `boolean` |  |
| `name` | `string` |  |
| `nhd_last_submission_year` | `number` |  |
| `nhd_submission_counter` | `number` |  |
| `person_paying_customs_duty` | `string` |  |
| `phone_country_code` | `string` |  |
| `phone_number` | `string` |  |
| `preferred_payment_method` | `table` |  |
| `signed_form` | `table` |  |
| `type` | `string` | * `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `co… |
| `type_of_person` | `table` |  |
| `unlocode` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```lua
local party, err = client:Party():load({ id = "party_id" })
```

#### Example: List

```lua
local partys, err = client:Party():list()
```

#### Example: Create

```lua
local party, err = client:Party():create({
  additional_declaration_type = {}, -- table
  address = {}, -- table
  authorisation = {}, -- table
  certificate = {}, -- table
  certificate_type = "example_certificate_type", -- string
  company = "example_company", -- string
  created_at = "example_created_at", -- string
  customs_office_of_lodgement = {}, -- table
  preferred_payment_method = {}, -- table
  signed_form = {}, -- table
  type_of_person = {}, -- table
  updated_at = "example_updated_at", -- string
})
```


### Submission

Create an instance: `local submission = client:Submission(nil)`

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
| `additional_external_ids` | `table` |  |
| `amendment_reason` | `string` |  |
| `amendment_status` | `string` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `answers` | `table` |  |
| `bypass_restricted_code` | `boolean` |  |
| `clearance_slip` | `table` |  |
| `client` | `table` |  |
| `company_member` | `table` | cuid-format identifier for this entity. |
| `consignee` | `table` |  |
| `consignor` | `table` |  |
| `created_at` | `string` |  |
| `declarant` | `table` | cuid-format identifier for this entity. |
| `document_upload_status` | `string` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `documents_presentation_requested` | `boolean` |  |
| `documents_upload_requested` | `boolean` |  |
| `external_id` | `string` |  |
| `form` | `string` | cuid-format identifier for this entity. |
| `goods_presentation_status` | `string` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested |
| `hrcm_status` | `string` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `id` | `string` | cuid-format identifier for this entity. |
| `invalidation_status` | `string` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled |
| `is_global_template` | `boolean` |  |
| `latest_notification_item` | `string` | cuid-format identifier for this entity. |
| `latest_state` | `table` |  |
| `lrn` | `string` |  |
| `mrn` | `string` |  |
| `name` | `string` |  |
| `partial_answers` | `boolean` |  |
| `receipt` | `table` |  |
| `refund_application_status` | `string` | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `route` | `string` | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `shipment_items_no` | `number` |  |
| `shipment_items_quantity_no` | `number` |  |
| `source` | `string` | cuid-format identifier for this entity. |
| `source_type` | `string` | * `template` - template * `automated_import` - automated_import |
| `status` | `string` | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `template` | `boolean` |  |
| `template_id` | `string` | cuid-format identifier for this entity. |
| `template_properties` | `table` |  |
| `total_tax_amount` | `string` |  |
| `updated_at` | `string` |  |
| `verification_errors` | `table` |  |
| `verification_status` | `string` | * `passed` - passed * `failed` - failed |

#### Example: Load

```lua
local submission, err = client:Submission():load({ id = "submission_id" })
```

#### Example: List

```lua
local submissions, err = client:Submission():list()
```

#### Example: Create

```lua
local submission, err = client:Submission():create({
  answers = {}, -- table
  clearance_slip = {}, -- table
  client = {}, -- table
  company_member = {}, -- table
  consignee = {}, -- table
  consignor = {}, -- table
  created_at = "example_created_at", -- string
  declarant = {}, -- table
  form = "example_form", -- string
  latest_state = {}, -- table
  receipt = {}, -- table
  updated_at = "example_updated_at", -- string
})
```


### SubmissionDetail

Create an instance: `local submission_detail = client:SubmissionDetail(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_external_ids` | `table` |  |
| `additional_information` | `table` |  |
| `amendment_status` | `string` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `clearance_slip` | `table` |  |
| `client` | `table` |  |
| `company` | `string` | cuid-format identifier for this entity. |
| `company_member` | `table` | cuid-format identifier for this entity. |
| `consignee` | `table` |  |
| `consignor` | `table` |  |
| `created_at` | `string` |  |
| `declarant` | `table` |  |
| `document_upload_status` | `string` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `documents_presentation_requested` | `boolean` |  |
| `documents_upload_requested` | `boolean` |  |
| `external_id` | `string` |  |
| `form` | `string` |  |
| `goods_presentation_status` | `string` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested |
| `hrcm_status` | `string` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `id` | `string` | cuid-format identifier for this entity. |
| `invalidation_status` | `string` | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled |
| `is_global_template` | `boolean` |  |
| `latest_notification_item` | `string` | cuid-format identifier for this entity. |
| `latest_state` | `table` |  |
| `lrn` | `string` |  |
| `mrn` | `string` |  |
| `name` | `string` |  |
| `receipt` | `table` |  |
| `refund_application_status` | `string` | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `route` | `string` | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `shipment_items_no` | `number` |  |
| `shipment_items_quantity_no` | `number` |  |
| `source` | `string` | cuid-format identifier for this entity. |
| `source_type` | `string` | * `template` - template * `automated_import` - automated_import |
| `status` | `string` | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `submission` | `string` | cuid-format identifier for this entity. |
| `supporting_documents` | `table` |  |
| `template` | `boolean` |  |
| `total_tax_amount` | `string` |  |
| `updated_at` | `string` |  |
| `verification_errors` | `table` |  |
| `verification_status` | `string` | * `passed` - passed * `failed` - failed |

#### Example: Create

```lua
local submission_detail, err = client:SubmissionDetail():create({
  additional_information = {}, -- table
  clearance_slip = {}, -- table
  client = {}, -- table
  company = "example_company", -- string
  company_member = {}, -- table
  consignee = {}, -- table
  consignor = {}, -- table
  created_at = "example_created_at", -- string
  declarant = {}, -- table
  form = "example_form", -- string
  latest_state = {}, -- table
  receipt = {}, -- table
  submission = "example_submission", -- string
  updated_at = "example_updated_at", -- string
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

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── customs-window_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`customs-window_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```lua
local submission = client:Submission()
submission:list()

-- submission:data_get() now returns the submission data from the last list
-- submission:match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
