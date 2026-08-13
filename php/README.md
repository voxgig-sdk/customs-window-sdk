# CustomsWindow PHP SDK



The PHP SDK for the CustomsWindow API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->BulkUpload()` — with named operations (`list`/`load`/`create`/`update`/`remove`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/customs-window-sdk/releases](https://github.com/voxgig-sdk/customs-window-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'customswindow_sdk.php';

$client = new CustomsWindowSDK([
    "apikey" => getenv("CUSTOMS_WINDOW_APIKEY"),
]);
```

### 2. List bulkupload records

```php
try {
    // list() returns an array of BulkUpload records — iterate directly.
    $bulkuploads = $client->BulkUpload()->list();
    foreach ($bulkuploads as $item) {
        echo $item["id"] . " " . $item["active_transport_nationality"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 3. Load a bulkupload

```php
try {
    // load() returns the ENTITY — call data_get() for the BulkUpload record (throws on error).
    $bulkupload = $client->BulkUpload()->load(["id" => "example_id"]);
    print_r($bulkupload);
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 4. Create, update, and remove

```php
// create() returns the ENTITY — call data_get() for the created BulkUpload record.
$created = $client->BulkUpload()->create(["client" => "example_client", "company_member" => "example_company_member", "created_at" => "example_created_at", "declarant" => "example_declarant", "declarations_no" => 1, "errors_file" => [], "file" => [], "updated_at" => "example_updated_at"]);

// Update — index the record via data_get() ($created->data_get()["id"]).
$client->BulkUpload()->update(["id" => $created->data_get()["id"], "active_transport_nationality" => "example_active_transport_nationality", "active_transport_number" => "example_active_transport_number"]);

// Remove
$client->BulkUpload()->remove(["id" => $created->data_get()["id"]]);
```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $submissions = $client->Submission()->list();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```php
$client = CustomsWindowSDK::test([
    "entity" => ["submission" => ["test01" => ["id" => "test01"]]],
]);

// Entity ops return the ENTITY (throws on error);
// call data_get() for the mock record.
$submission = $client->Submission()->list();
print_r($submission);
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new CustomsWindowSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
CUSTOMS_WINDOW_TEST_LIVE=TRUE
CUSTOMS_WINDOW_APIKEY=<your-key>
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### CustomsWindowSDK

```php
require_once 'customswindow_sdk.php';
$client = new CustomsWindowSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = CustomsWindowSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### CustomsWindowSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `BulkUpload` | `($data): BulkUploadEntity` | Create a BulkUpload entity instance. |
| `File` | `($data): FileEntity` | Create a File entity instance. |
| `PaginatedBulkUploadListList` | `($data): PaginatedBulkUploadListListEntity` | Create a PaginatedBulkUploadListList entity instance. |
| `PaginatedPartyListList` | `($data): PaginatedPartyListListEntity` | Create a PaginatedPartyListList entity instance. |
| `PaginatedSubmissionListList` | `($data): PaginatedSubmissionListListEntity` | Create a PaginatedSubmissionListList entity instance. |
| `Party` | `($data): PartyEntity` | Create a Party entity instance. |
| `Submission` | `($data): SubmissionEntity` | Create a Submission entity instance. |
| `SubmissionDetail` | `($data): SubmissionDetailEntity` | Create a SubmissionDetail entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `create` | `($reqdata, $ctrl): array` | Create a new entity. |
| `update` | `($reqdata, $ctrl): array` | Update an existing entity. |
| `remove` | `($reqmatch, $ctrl): array` | Remove an entity. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

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

Create an instance: `$bulk_upload = $client->BulkUpload();`

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
| `errors_file` | `array` |  |
| `external_id` | `string` |  |
| `failed_declarations_no` | `int` |  |
| `file` | `array` |  |
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

```php
// load() returns the ENTITY — call data_get() for the BulkUpload record (throws on error).
$bulk_upload = $client->BulkUpload()->load(["id" => "bulk_upload_id"]);
```

#### Example: List

```php
// list() returns an array of BulkUpload records (throws on error).
$bulk_uploads = $client->BulkUpload()->list();
```

#### Example: Create

```php
$bulk_upload = $client->BulkUpload()->create([
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


### File

Create an instance: `$file = $client->File();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

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

```php
$file = $client->File()->create([
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


### PaginatedBulkUploadListList

Create an instance: `$paginated_bulk_upload_list_list = $client->PaginatedBulkUploadListList();`


### PaginatedPartyListList

Create an instance: `$paginated_party_list_list = $client->PaginatedPartyListList();`


### PaginatedSubmissionListList

Create an instance: `$paginated_submission_list_list = $client->PaginatedSubmissionListList();`


### Party

Create an instance: `$party = $client->Party();`

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
| `additional_declaration_type` | `array` |  |
| `address` | `array` |  |
| `authorisation` | `array` |  |
| `bank_details` | `string` |  |
| `certificate` | `array` |  |
| `certificate_type` | `string` |  |
| `company` | `string` |  |
| `created_at` | `string` |  |
| `customs_office_of_lodgement` | `array` |  |
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
| `preferred_payment_method` | `array` |  |
| `signed_form` | `array` |  |
| `type` | `string` |  |
| `type_of_person` | `array` |  |
| `unlocode` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Party record (throws on error).
$party = $client->Party()->load(["id" => "party_id"]);
```

#### Example: List

```php
// list() returns an array of Party records (throws on error).
$partys = $client->Party()->list();
```

#### Example: Create

```php
$party = $client->Party()->create([
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


### Submission

Create an instance: `$submission = $client->Submission();`

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
| `additional_external_ids` | `array` |  |
| `amendment_reason` | `string` |  |
| `amendment_status` | `string` |  |
| `answers` | `array` |  |
| `bypass_restricted_code` | `bool` |  |
| `clearance_slip` | `array` |  |
| `client` | `array` |  |
| `company_member` | `array` |  |
| `consignee` | `array` |  |
| `consignor` | `array` |  |
| `created_at` | `string` |  |
| `declarant` | `array` |  |
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
| `latest_state` | `array` |  |
| `lrn` | `string` |  |
| `mrn` | `string` |  |
| `name` | `string` |  |
| `partial_answers` | `bool` |  |
| `receipt` | `array` |  |
| `refund_application_status` | `string` |  |
| `route` | `string` |  |
| `shipment_items_no` | `int` |  |
| `shipment_items_quantity_no` | `int` |  |
| `source` | `string` |  |
| `source_type` | `string` |  |
| `status` | `string` |  |
| `template` | `bool` |  |
| `template_id` | `string` |  |
| `template_properties` | `array` |  |
| `total_tax_amount` | `string` |  |
| `updated_at` | `string` |  |
| `verification_errors` | `array` |  |
| `verification_status` | `string` |  |

#### Example: Load

```php
// load() returns the ENTITY — call data_get() for the Submission record (throws on error).
$submission = $client->Submission()->load(["id" => "submission_id"]);
```

#### Example: List

```php
// list() returns an array of Submission records (throws on error).
$submissions = $client->Submission()->list();
```

#### Example: Create

```php
$submission = $client->Submission()->create([
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


### SubmissionDetail

Create an instance: `$submission_detail = $client->SubmissionDetail();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_external_ids` | `array` |  |
| `additional_information` | `array` |  |
| `amendment_status` | `string` |  |
| `clearance_slip` | `array` |  |
| `client` | `array` |  |
| `company` | `string` |  |
| `company_member` | `array` |  |
| `consignee` | `array` |  |
| `consignor` | `array` |  |
| `created_at` | `string` |  |
| `declarant` | `array` |  |
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
| `latest_state` | `array` |  |
| `lrn` | `string` |  |
| `mrn` | `string` |  |
| `name` | `string` |  |
| `receipt` | `array` |  |
| `refund_application_status` | `string` |  |
| `route` | `string` |  |
| `shipment_items_no` | `int` |  |
| `shipment_items_quantity_no` | `int` |  |
| `source` | `string` |  |
| `source_type` | `string` |  |
| `status` | `string` |  |
| `submission` | `string` |  |
| `supporting_documents` | `array` |  |
| `template` | `bool` |  |
| `total_tax_amount` | `string` |  |
| `updated_at` | `string` |  |
| `verification_errors` | `array` |  |
| `verification_status` | `string` |  |

#### Example: Create

```php
$submission_detail = $client->SubmissionDetail()->create([
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

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── customswindow_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`customswindow_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```php
$submission = $client->Submission();
$submission->list();

// $submission->data_get() now returns the submission data from the last list
// $submission->match_get() returns the last match criteria
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
