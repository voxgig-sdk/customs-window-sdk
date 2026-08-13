# CustomsWindow Lua SDK Reference

Complete API reference for the CustomsWindow Lua SDK.


## CustomsWindowSDK

### Constructor

```lua
local sdk = require("customs-window_sdk")
local client = sdk.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `table` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `table` | Custom headers for all requests. |
| `options.feature` | `table` | Feature configuration. |
| `options.system` | `table` | System overrides (e.g. custom fetch). |


### Static Methods

#### `sdk.test(testopts?, sdkopts?)`

Create a test client with mock features active. Both arguments are optional.

```lua
local client = sdk.test()
```


### Instance Methods

#### `BulkUpload(data)`

Create a new `BulkUpload` entity instance. Pass `nil` for no initial data.

#### `File(data)`

Create a new `File` entity instance. Pass `nil` for no initial data.

#### `PaginatedBulkUploadListList(data)`

Create a new `PaginatedBulkUploadListList` entity instance. Pass `nil` for no initial data.

#### `PaginatedPartyListList(data)`

Create a new `PaginatedPartyListList` entity instance. Pass `nil` for no initial data.

#### `PaginatedSubmissionListList(data)`

Create a new `PaginatedSubmissionListList` entity instance. Pass `nil` for no initial data.

#### `Party(data)`

Create a new `Party` entity instance. Pass `nil` for no initial data.

#### `Submission(data)`

Create a new `Submission` entity instance. Pass `nil` for no initial data.

#### `SubmissionDetail(data)`

Create a new `SubmissionDetail` entity instance. Pass `nil` for no initial data.

#### `options_map() -> table`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs) -> table, err`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs.params` | `table` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `table` | Query string parameters. |
| `fetchargs.headers` | `table` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (tables are JSON-serialized). |
| `fetchargs.ctrl` | `table` | Control options (e.g. `{ explain = true }`). |

**Returns:** `table, err`

#### `prepare(fetchargs) -> table, err`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `table, err`


---

## BulkUploadEntity

```lua
local bulk_upload = client:BulkUpload(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `active_transport_nationality` | `string` | No |  |
| `active_transport_number` | `string` | No |  |
| `arrival_datetime` | `string` | No |  |
| `client` | `string` | Yes |  |
| `company_member` | `string` | Yes |  |
| `created_at` | `string` | Yes |  |
| `declarant` | `string` | Yes |  |
| `declarations_no` | `number` | Yes |  |
| `deleted_at` | `string` | No |  |
| `departure_datetime` | `string` | No |  |
| `errors_file` | `table` | Yes |  |
| `external_id` | `string` | No |  |
| `failed_declarations_no` | `number` | No |  |
| `file` | `table` | Yes |  |
| `green_routed_no` | `number` | No |  |
| `h1_fallback_template` | `string` | No |  |
| `house_transport_doc_ref` | `string` | No |  |
| `id` | `string` | No |  |
| `issue_date` | `string` | No |  |
| `mapping` | `string` | No |  |
| `orange_routed_no` | `number` | No |  |
| `parsed_declarations_no` | `number` | No |  |
| `parser` | `string` | No |  |
| `parsing_completed_at` | `string` | No |  |
| `parsing_started_at` | `string` | No |  |
| `passive_transport_nationality` | `string` | No |  |
| `passive_transport_number` | `string` | No |  |
| `processed_declarations_no` | `number` | No |  |
| `processing_ended_at` | `string` | No |  |
| `processing_started_at` | `string` | No |  |
| `receipt_generating_started_at` | `string` | No |  |
| `receipt_request_started_at` | `string` | No |  |
| `receipt_request_status` | `string` | No |  |
| `receipt_request_user` | `string` | No |  |
| `receipts_zip` | `string` | No |  |
| `red_routed_no` | `number` | No |  |
| `rejected_status_no` | `number` | No |  |
| `status` | `string` | No |  |
| `template` | `string` | No |  |
| `updated_at` | `string` | Yes |  |
| `yellow_routed_no` | `number` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:BulkUpload():create({
  client = --[[ string ]],
  company_member = --[[ string ]],
  created_at = --[[ string ]],
  declarant = --[[ string ]],
  declarations_no = --[[ number ]],
  errors_file = --[[ table ]],
  file = --[[ table ]],
  updated_at = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:BulkUpload():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:BulkUpload():load({ id = "bulk_upload_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:BulkUpload():remove({ id = "bulk_upload_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:BulkUpload():update({
  id = "bulk_upload_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BulkUploadEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## FileEntity

```lua
local file = client:File(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `company` | `string` | Yes |  |
| `created_at` | `string` | Yes |  |
| `extension` | `string` | Yes |  |
| `file` | `string` | Yes |  |
| `id` | `string` | No |  |
| `name` | `string` | Yes |  |
| `public` | `boolean` | Yes |  |
| `size` | `number` | No |  |
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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:File():create({
  company = --[[ string ]],
  created_at = --[[ string ]],
  extension = --[[ string ]],
  file = --[[ string ]],
  name = --[[ string ]],
  public = --[[ boolean ]],
  updated_at = --[[ string ]],
  url = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FileEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PaginatedBulkUploadListListEntity

```lua
local paginated_bulk_upload_list_list = client:PaginatedBulkUploadListList(nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaginatedBulkUploadListListEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PaginatedPartyListListEntity

```lua
local paginated_party_list_list = client:PaginatedPartyListList(nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaginatedPartyListListEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PaginatedSubmissionListListEntity

```lua
local paginated_submission_list_list = client:PaginatedSubmissionListList(nil)
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaginatedSubmissionListListEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## PartyEntity

```lua
local party = client:Party(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_declaration_type` | `table` | Yes |  |
| `address` | `table` | Yes |  |
| `authorisation` | `table` | Yes |  |
| `bank_details` | `string` | No |  |
| `certificate` | `table` | Yes |  |
| `certificate_type` | `string` | Yes |  |
| `company` | `string` | Yes |  |
| `created_at` | `string` | Yes |  |
| `customs_office_of_lodgement` | `table` | Yes |  |
| `deleted_at` | `string` | No |  |
| `email` | `string` | No |  |
| `id` | `string` | No |  |
| `identification_number` | `string` | No |  |
| `indirect_representative` | `boolean` | No |  |
| `name` | `string` | No |  |
| `nhd_last_submission_year` | `number` | No |  |
| `nhd_submission_counter` | `number` | No |  |
| `person_paying_customs_duty` | `string` | No |  |
| `phone_country_code` | `string` | No |  |
| `phone_number` | `string` | No |  |
| `preferred_payment_method` | `table` | Yes |  |
| `signed_form` | `table` | Yes |  |
| `type` | `string` | No |  |
| `type_of_person` | `table` | Yes |  |
| `unlocode` | `string` | No |  |
| `updated_at` | `string` | Yes |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Party():create({
  additional_declaration_type = --[[ table ]],
  address = --[[ table ]],
  authorisation = --[[ table ]],
  certificate = --[[ table ]],
  certificate_type = --[[ string ]],
  company = --[[ string ]],
  created_at = --[[ string ]],
  customs_office_of_lodgement = --[[ table ]],
  preferred_payment_method = --[[ table ]],
  signed_form = --[[ table ]],
  type_of_person = --[[ table ]],
  updated_at = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Party():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Party():load({ id = "party_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Party():remove({ id = "party_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Party():update({
  id = "party_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PartyEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SubmissionEntity

```lua
local submission = client:Submission(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_external_ids` | `table` | No |  |
| `amendment_reason` | `string` | No |  |
| `amendment_status` | `string` | No |  |
| `answers` | `table` | Yes |  |
| `bypass_restricted_code` | `boolean` | No |  |
| `clearance_slip` | `table` | Yes |  |
| `client` | `table` | Yes |  |
| `company_member` | `table` | Yes |  |
| `consignee` | `table` | Yes |  |
| `consignor` | `table` | Yes |  |
| `created_at` | `string` | Yes |  |
| `declarant` | `table` | Yes |  |
| `document_upload_status` | `string` | No |  |
| `documents_presentation_requested` | `boolean` | No |  |
| `documents_upload_requested` | `boolean` | No |  |
| `external_id` | `string` | No |  |
| `form` | `string` | Yes |  |
| `goods_presentation_status` | `string` | No |  |
| `hrcm_status` | `string` | No |  |
| `id` | `string` | No |  |
| `invalidation_status` | `string` | No |  |
| `is_global_template` | `boolean` | No |  |
| `latest_notification_item` | `string` | No |  |
| `latest_state` | `table` | Yes |  |
| `lrn` | `string` | No |  |
| `mrn` | `string` | No |  |
| `name` | `string` | No |  |
| `partial_answers` | `boolean` | No |  |
| `receipt` | `table` | Yes |  |
| `refund_application_status` | `string` | No |  |
| `route` | `string` | No |  |
| `shipment_items_no` | `number` | No |  |
| `shipment_items_quantity_no` | `number` | No |  |
| `source` | `string` | No |  |
| `source_type` | `string` | No |  |
| `status` | `string` | No |  |
| `template` | `boolean` | No |  |
| `template_id` | `string` | No |  |
| `template_properties` | `table` | No |  |
| `total_tax_amount` | `string` | No |  |
| `updated_at` | `string` | Yes |  |
| `verification_errors` | `table` | No |  |
| `verification_status` | `string` | No |  |

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

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:Submission():create({
  answers = --[[ table ]],
  clearance_slip = --[[ table ]],
  client = --[[ table ]],
  company_member = --[[ table ]],
  consignee = --[[ table ]],
  consignor = --[[ table ]],
  created_at = --[[ string ]],
  declarant = --[[ table ]],
  form = --[[ string ]],
  latest_state = --[[ table ]],
  receipt = --[[ table ]],
  updated_at = --[[ string ]],
})
```

#### `list(reqmatch, ctrl) -> any, err`

List entities matching the given criteria. Returns an array.

```lua
local results, err = client:Submission():list()
```

#### `load(reqmatch, ctrl) -> any, err`

Load a single entity matching the given criteria.

```lua
local result, err = client:Submission():load({ id = "submission_id" })
```

#### `remove(reqmatch, ctrl) -> any, err`

Remove the entity matching the given criteria.

```lua
local result, err = client:Submission():remove({ id = "submission_id" })
```

#### `update(reqdata, ctrl) -> any, err`

Update an existing entity. The data must include the entity `id`.

```lua
local result, err = client:Submission():update({
  id = "submission_id",
  -- Fields to update
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SubmissionEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## SubmissionDetailEntity

```lua
local submission_detail = client:SubmissionDetail(nil)
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_external_ids` | `table` | No |  |
| `additional_information` | `table` | Yes |  |
| `amendment_status` | `string` | No |  |
| `clearance_slip` | `table` | Yes |  |
| `client` | `table` | Yes |  |
| `company` | `string` | Yes |  |
| `company_member` | `table` | Yes |  |
| `consignee` | `table` | Yes |  |
| `consignor` | `table` | Yes |  |
| `created_at` | `string` | Yes |  |
| `declarant` | `table` | Yes |  |
| `document_upload_status` | `string` | No |  |
| `documents_presentation_requested` | `boolean` | No |  |
| `documents_upload_requested` | `boolean` | No |  |
| `external_id` | `string` | No |  |
| `form` | `string` | Yes |  |
| `goods_presentation_status` | `string` | No |  |
| `hrcm_status` | `string` | No |  |
| `id` | `string` | No |  |
| `invalidation_status` | `string` | No |  |
| `is_global_template` | `boolean` | No |  |
| `latest_notification_item` | `string` | No |  |
| `latest_state` | `table` | Yes |  |
| `lrn` | `string` | No |  |
| `mrn` | `string` | No |  |
| `name` | `string` | No |  |
| `receipt` | `table` | Yes |  |
| `refund_application_status` | `string` | No |  |
| `route` | `string` | No |  |
| `shipment_items_no` | `number` | No |  |
| `shipment_items_quantity_no` | `number` | No |  |
| `source` | `string` | No |  |
| `source_type` | `string` | No |  |
| `status` | `string` | No |  |
| `submission` | `string` | Yes |  |
| `supporting_documents` | `table` | No |  |
| `template` | `boolean` | No |  |
| `total_tax_amount` | `string` | No |  |
| `updated_at` | `string` | Yes |  |
| `verification_errors` | `table` | No |  |
| `verification_status` | `string` | No |  |

### Operations

#### `create(reqdata, ctrl) -> any, err`

Create a new entity with the given data.

```lua
local result, err = client:SubmissionDetail():create({
  additional_information = --[[ table ]],
  clearance_slip = --[[ table ]],
  client = --[[ table ]],
  company = --[[ string ]],
  company_member = --[[ table ]],
  consignee = --[[ table ]],
  consignor = --[[ table ]],
  created_at = --[[ string ]],
  declarant = --[[ table ]],
  form = --[[ string ]],
  latest_state = --[[ table ]],
  receipt = --[[ table ]],
  submission = --[[ string ]],
  updated_at = --[[ string ]],
})
```

### Common Methods

#### `data_get() -> table`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> table`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SubmissionDetailEntity` instance with the same client and
options.

#### `get_name() -> string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```lua
local client = sdk.new({
  feature = {
    test = { active = true },
  },
})
```

