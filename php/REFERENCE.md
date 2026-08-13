# CustomsWindow PHP SDK Reference

Complete API reference for the CustomsWindow PHP SDK.


## CustomsWindowSDK

### Constructor

```php
require_once __DIR__ . '/customswindow_sdk.php';

$client = new CustomsWindowSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CustomsWindowSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = CustomsWindowSDK::test();
```


### Instance Methods

#### `BulkUpload($data = null)`

Create a new `BulkUploadEntity` instance. Pass `null` for no initial data.

#### `File($data = null)`

Create a new `FileEntity` instance. Pass `null` for no initial data.

#### `PaginatedBulkUploadListList($data = null)`

Create a new `PaginatedBulkUploadListListEntity` instance. Pass `null` for no initial data.

#### `PaginatedPartyListList($data = null)`

Create a new `PaginatedPartyListListEntity` instance. Pass `null` for no initial data.

#### `PaginatedSubmissionListList($data = null)`

Create a new `PaginatedSubmissionListListEntity` instance. Pass `null` for no initial data.

#### `Party($data = null)`

Create a new `PartyEntity` instance. Pass `null` for no initial data.

#### `Submission($data = null)`

Create a new `SubmissionEntity` instance. Pass `null` for no initial data.

#### `SubmissionDetail($data = null)`

Create a new `SubmissionDetailEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): CustomsWindowUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## BulkUploadEntity

```php
$bulk_upload = $client->BulkUpload();
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
| `declarations_no` | `int` | Yes |  |
| `deleted_at` | `string` | No |  |
| `departure_datetime` | `string` | No |  |
| `errors_file` | `array` | Yes |  |
| `external_id` | `string` | No |  |
| `failed_declarations_no` | `int` | No |  |
| `file` | `array` | Yes |  |
| `green_routed_no` | `int` | No |  |
| `h1_fallback_template` | `string` | No |  |
| `house_transport_doc_ref` | `string` | No |  |
| `id` | `string` | No |  |
| `issue_date` | `string` | No |  |
| `mapping` | `string` | No |  |
| `orange_routed_no` | `int` | No |  |
| `parsed_declarations_no` | `int` | No |  |
| `parser` | `string` | No |  |
| `parsing_completed_at` | `string` | No |  |
| `parsing_started_at` | `string` | No |  |
| `passive_transport_nationality` | `string` | No |  |
| `passive_transport_number` | `string` | No |  |
| `processed_declarations_no` | `int` | No |  |
| `processing_ended_at` | `string` | No |  |
| `processing_started_at` | `string` | No |  |
| `receipt_generating_started_at` | `string` | No |  |
| `receipt_request_started_at` | `string` | No |  |
| `receipt_request_status` | `string` | No |  |
| `receipt_request_user` | `string` | No |  |
| `receipts_zip` | `string` | No |  |
| `red_routed_no` | `int` | No |  |
| `rejected_status_no` | `int` | No |  |
| `status` | `string` | No |  |
| `template` | `string` | No |  |
| `updated_at` | `string` | Yes |  |
| `yellow_routed_no` | `int` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->BulkUpload()->create([
  "client" => null, // string
  "company_member" => null, // string
  "created_at" => null, // string
  "declarant" => null, // string
  "declarations_no" => null, // int
  "errors_file" => null, // array
  "file" => null, // array
  "updated_at" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->BulkUpload()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->BulkUpload()->load(["id" => "bulk_upload_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->BulkUpload()->remove(["id" => "bulk_upload_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->BulkUpload()->update([
  "id" => "bulk_upload_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): BulkUploadEntity`

Create a new `BulkUploadEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## FileEntity

```php
$file = $client->File();
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->File()->create([
  "company" => null, // string
  "created_at" => null, // string
  "extension" => null, // string
  "file" => null, // string
  "name" => null, // string
  "public" => null, // bool
  "updated_at" => null, // string
  "url" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): FileEntity`

Create a new `FileEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PaginatedBulkUploadListListEntity

```php
$paginated_bulk_upload_list_list = $client->PaginatedBulkUploadListList();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PaginatedBulkUploadListListEntity`

Create a new `PaginatedBulkUploadListListEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PaginatedPartyListListEntity

```php
$paginated_party_list_list = $client->PaginatedPartyListList();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PaginatedPartyListListEntity`

Create a new `PaginatedPartyListListEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PaginatedSubmissionListListEntity

```php
$paginated_submission_list_list = $client->PaginatedSubmissionListList();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PaginatedSubmissionListListEntity`

Create a new `PaginatedSubmissionListListEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PartyEntity

```php
$party = $client->Party();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_declaration_type` | `array` | Yes |  |
| `address` | `array` | Yes |  |
| `authorisation` | `array` | Yes |  |
| `bank_details` | `string` | No |  |
| `certificate` | `array` | Yes |  |
| `certificate_type` | `string` | Yes |  |
| `company` | `string` | Yes |  |
| `created_at` | `string` | Yes |  |
| `customs_office_of_lodgement` | `array` | Yes |  |
| `deleted_at` | `string` | No |  |
| `email` | `string` | No |  |
| `id` | `string` | No |  |
| `identification_number` | `string` | No |  |
| `indirect_representative` | `bool` | No |  |
| `name` | `string` | No |  |
| `nhd_last_submission_year` | `int` | No |  |
| `nhd_submission_counter` | `int` | No |  |
| `person_paying_customs_duty` | `string` | No |  |
| `phone_country_code` | `string` | No |  |
| `phone_number` | `string` | No |  |
| `preferred_payment_method` | `array` | Yes |  |
| `signed_form` | `array` | Yes |  |
| `type` | `string` | No |  |
| `type_of_person` | `array` | Yes |  |
| `unlocode` | `string` | No |  |
| `updated_at` | `string` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Party()->create([
  "additional_declaration_type" => null, // array
  "address" => null, // array
  "authorisation" => null, // array
  "certificate" => null, // array
  "certificate_type" => null, // string
  "company" => null, // string
  "created_at" => null, // string
  "customs_office_of_lodgement" => null, // array
  "preferred_payment_method" => null, // array
  "signed_form" => null, // array
  "type_of_person" => null, // array
  "updated_at" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Party()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Party()->load(["id" => "party_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Party()->remove(["id" => "party_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Party()->update([
  "id" => "party_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PartyEntity`

Create a new `PartyEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SubmissionEntity

```php
$submission = $client->Submission();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_external_ids` | `array` | No |  |
| `amendment_reason` | `string` | No |  |
| `amendment_status` | `string` | No |  |
| `answers` | `array` | Yes |  |
| `bypass_restricted_code` | `bool` | No |  |
| `clearance_slip` | `array` | Yes |  |
| `client` | `array` | Yes |  |
| `company_member` | `array` | Yes |  |
| `consignee` | `array` | Yes |  |
| `consignor` | `array` | Yes |  |
| `created_at` | `string` | Yes |  |
| `declarant` | `array` | Yes |  |
| `document_upload_status` | `string` | No |  |
| `documents_presentation_requested` | `bool` | No |  |
| `documents_upload_requested` | `bool` | No |  |
| `external_id` | `string` | No |  |
| `form` | `string` | Yes |  |
| `goods_presentation_status` | `string` | No |  |
| `hrcm_status` | `string` | No |  |
| `id` | `string` | No |  |
| `invalidation_status` | `string` | No |  |
| `is_global_template` | `bool` | No |  |
| `latest_notification_item` | `string` | No |  |
| `latest_state` | `array` | Yes |  |
| `lrn` | `string` | No |  |
| `mrn` | `string` | No |  |
| `name` | `string` | No |  |
| `partial_answers` | `bool` | No |  |
| `receipt` | `array` | Yes |  |
| `refund_application_status` | `string` | No |  |
| `route` | `string` | No |  |
| `shipment_items_no` | `int` | No |  |
| `shipment_items_quantity_no` | `int` | No |  |
| `source` | `string` | No |  |
| `source_type` | `string` | No |  |
| `status` | `string` | No |  |
| `template` | `bool` | No |  |
| `template_id` | `string` | No |  |
| `template_properties` | `array` | No |  |
| `total_tax_amount` | `string` | No |  |
| `updated_at` | `string` | Yes |  |
| `verification_errors` | `array` | No |  |
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

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Submission()->create([
  "answers" => null, // array
  "clearance_slip" => null, // array
  "client" => null, // array
  "company_member" => null, // array
  "consignee" => null, // array
  "consignor" => null, // array
  "created_at" => null, // string
  "declarant" => null, // array
  "form" => null, // string
  "latest_state" => null, // array
  "receipt" => null, // array
  "updated_at" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Submission()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Submission()->load(["id" => "submission_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Submission()->remove(["id" => "submission_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Submission()->update([
  "id" => "submission_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SubmissionEntity`

Create a new `SubmissionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SubmissionDetailEntity

```php
$submission_detail = $client->SubmissionDetail();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_external_ids` | `array` | No |  |
| `additional_information` | `array` | Yes |  |
| `amendment_status` | `string` | No |  |
| `clearance_slip` | `array` | Yes |  |
| `client` | `array` | Yes |  |
| `company` | `string` | Yes |  |
| `company_member` | `array` | Yes |  |
| `consignee` | `array` | Yes |  |
| `consignor` | `array` | Yes |  |
| `created_at` | `string` | Yes |  |
| `declarant` | `array` | Yes |  |
| `document_upload_status` | `string` | No |  |
| `documents_presentation_requested` | `bool` | No |  |
| `documents_upload_requested` | `bool` | No |  |
| `external_id` | `string` | No |  |
| `form` | `string` | Yes |  |
| `goods_presentation_status` | `string` | No |  |
| `hrcm_status` | `string` | No |  |
| `id` | `string` | No |  |
| `invalidation_status` | `string` | No |  |
| `is_global_template` | `bool` | No |  |
| `latest_notification_item` | `string` | No |  |
| `latest_state` | `array` | Yes |  |
| `lrn` | `string` | No |  |
| `mrn` | `string` | No |  |
| `name` | `string` | No |  |
| `receipt` | `array` | Yes |  |
| `refund_application_status` | `string` | No |  |
| `route` | `string` | No |  |
| `shipment_items_no` | `int` | No |  |
| `shipment_items_quantity_no` | `int` | No |  |
| `source` | `string` | No |  |
| `source_type` | `string` | No |  |
| `status` | `string` | No |  |
| `submission` | `string` | Yes |  |
| `supporting_documents` | `array` | No |  |
| `template` | `bool` | No |  |
| `total_tax_amount` | `string` | No |  |
| `updated_at` | `string` | Yes |  |
| `verification_errors` | `array` | No |  |
| `verification_status` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->SubmissionDetail()->create([
  "additional_information" => null, // array
  "clearance_slip" => null, // array
  "client" => null, // array
  "company" => null, // string
  "company_member" => null, // array
  "consignee" => null, // array
  "consignor" => null, // array
  "created_at" => null, // string
  "declarant" => null, // array
  "form" => null, // string
  "latest_state" => null, // array
  "receipt" => null, // array
  "submission" => null, // string
  "updated_at" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SubmissionDetailEntity`

Create a new `SubmissionDetailEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new CustomsWindowSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

