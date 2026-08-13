# CustomsWindow Python SDK



The Python SDK for the CustomsWindow API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.BulkUpload()` — each
carrying a small, uniform set of operations (`list`, `load`, `create`, `update`, `remove`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/customs-window-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from customswindow_sdk import CustomsWindowSDK

client = CustomsWindowSDK({
    "apikey": os.environ.get("CUSTOMS_WINDOW_APIKEY"),
})
```

### 2. List bulkupload records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    bulkuploads = client.BulkUpload().list()
    for bulkupload in bulkuploads:
        print(bulkupload)
except Exception as err:
    print(f"list failed: {err}")
```

### 3. Load a bulkupload

`load()` returns the ENTITY — call data_get() for the record — and raises on error.

```python
try:
    bulkupload = client.BulkUpload().load({"id": "example_id"})
    print(bulkupload)
except Exception as err:
    print(f"load failed: {err}")
```

### 4. Create, update, and remove

```python
# Create — returns the ENTITY (call data_get() for the record)
created = client.BulkUpload().create({"client": "example_client", "company_member": "example_company_member", "created_at": "example_created_at", "declarant": "example_declarant", "declarations_no": 1, "errors_file": {}, "file": {}, "updated_at": "example_updated_at"})

# Update — the created record's id is a plain dict key
client.BulkUpload().update({"id": created.data_get()["id"], "active_transport_nationality": "example_active_transport_nationality", "active_transport_number": "example_active_transport_number"})

# Remove
client.BulkUpload().remove({"id": created.data_get()["id"]})
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    submissions = client.Submission().list()
    print(submissions)
except Exception as err:
    print(f"list failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = CustomsWindowSDK.test()

# Entity ops return the ENTITY and raises on error;
# call data_get() for the record.
submission = client.Submission().list()
# submission contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = CustomsWindowSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
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
cd py && pytest test/
```


## Reference

### CustomsWindowSDK

```python
from customswindow_sdk import CustomsWindowSDK

client = CustomsWindowSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = CustomsWindowSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### CustomsWindowSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
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
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the ENTITY (call data_get() for the record) (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

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

Create an instance: `bulk_upload = client.BulkUpload()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `active_transport_nationality` | `str` |  |
| `active_transport_number` | `str` |  |
| `arrival_datetime` | `str` |  |
| `client` | `str` |  |
| `company_member` | `str` |  |
| `created_at` | `str` |  |
| `declarant` | `str` |  |
| `declarations_no` | `int` |  |
| `deleted_at` | `str` |  |
| `departure_datetime` | `str` |  |
| `errors_file` | `dict` |  |
| `external_id` | `str` |  |
| `failed_declarations_no` | `int` |  |
| `file` | `dict` |  |
| `green_routed_no` | `int` |  |
| `h1_fallback_template` | `str` |  |
| `house_transport_doc_ref` | `str` |  |
| `id` | `str` |  |
| `issue_date` | `str` |  |
| `mapping` | `str` |  |
| `orange_routed_no` | `int` |  |
| `parsed_declarations_no` | `int` |  |
| `parser` | `str` |  |
| `parsing_completed_at` | `str` |  |
| `parsing_started_at` | `str` |  |
| `passive_transport_nationality` | `str` |  |
| `passive_transport_number` | `str` |  |
| `processed_declarations_no` | `int` |  |
| `processing_ended_at` | `str` |  |
| `processing_started_at` | `str` |  |
| `receipt_generating_started_at` | `str` |  |
| `receipt_request_started_at` | `str` |  |
| `receipt_request_status` | `str` |  |
| `receipt_request_user` | `str` |  |
| `receipts_zip` | `str` |  |
| `red_routed_no` | `int` |  |
| `rejected_status_no` | `int` |  |
| `status` | `str` |  |
| `template` | `str` |  |
| `updated_at` | `str` |  |
| `yellow_routed_no` | `int` |  |

#### Example: Load

```python
bulk_upload = client.BulkUpload().load({"id": "bulk_upload_id"})
```

#### Example: List

```python
bulk_uploads = client.BulkUpload().list()
```

#### Example: Create

```python
bulk_upload = client.BulkUpload().create({
    "client": "example_client",  # str
    "company_member": "example_company_member",  # str
    "created_at": "example_created_at",  # str
    "declarant": "example_declarant",  # str
    "declarations_no": 1,  # int
    "errors_file": {},  # dict
    "file": {},  # dict
    "updated_at": "example_updated_at",  # str
})
```


### File

Create an instance: `file = client.File()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `company` | `str` |  |
| `created_at` | `str` |  |
| `extension` | `str` |  |
| `file` | `str` |  |
| `id` | `str` |  |
| `name` | `str` |  |
| `public` | `bool` |  |
| `size` | `int` |  |
| `updated_at` | `str` |  |
| `url` | `str` |  |

#### Example: Create

```python
file = client.File().create({
    "company": "example_company",  # str
    "created_at": "example_created_at",  # str
    "extension": "example_extension",  # str
    "file": "example_file",  # str
    "name": "example_name",  # str
    "public": True,  # bool
    "updated_at": "example_updated_at",  # str
    "url": "example_url",  # str
})
```


### PaginatedBulkUploadListList

Create an instance: `paginated_bulk_upload_list_list = client.PaginatedBulkUploadListList()`


### PaginatedPartyListList

Create an instance: `paginated_party_list_list = client.PaginatedPartyListList()`


### PaginatedSubmissionListList

Create an instance: `paginated_submission_list_list = client.PaginatedSubmissionListList()`


### Party

Create an instance: `party = client.Party()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_declaration_type` | `dict` |  |
| `address` | `dict` |  |
| `authorisation` | `dict` |  |
| `bank_details` | `str` |  |
| `certificate` | `dict` |  |
| `certificate_type` | `str` |  |
| `company` | `str` |  |
| `created_at` | `str` |  |
| `customs_office_of_lodgement` | `dict` |  |
| `deleted_at` | `str` |  |
| `email` | `str` |  |
| `id` | `str` |  |
| `identification_number` | `str` |  |
| `indirect_representative` | `bool` |  |
| `name` | `str` |  |
| `nhd_last_submission_year` | `int` |  |
| `nhd_submission_counter` | `int` |  |
| `person_paying_customs_duty` | `str` |  |
| `phone_country_code` | `str` |  |
| `phone_number` | `str` |  |
| `preferred_payment_method` | `dict` |  |
| `signed_form` | `dict` |  |
| `type` | `str` |  |
| `type_of_person` | `dict` |  |
| `unlocode` | `str` |  |
| `updated_at` | `str` |  |

#### Example: Load

```python
party = client.Party().load({"id": "party_id"})
```

#### Example: List

```python
partys = client.Party().list()
```

#### Example: Create

```python
party = client.Party().create({
    "additional_declaration_type": {},  # dict
    "address": {},  # dict
    "authorisation": {},  # dict
    "certificate": {},  # dict
    "certificate_type": "example_certificate_type",  # str
    "company": "example_company",  # str
    "created_at": "example_created_at",  # str
    "customs_office_of_lodgement": {},  # dict
    "preferred_payment_method": {},  # dict
    "signed_form": {},  # dict
    "type_of_person": {},  # dict
    "updated_at": "example_updated_at",  # str
})
```


### Submission

Create an instance: `submission = client.Submission()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_external_ids` | `list` |  |
| `amendment_reason` | `str` |  |
| `amendment_status` | `str` |  |
| `answers` | `list` |  |
| `bypass_restricted_code` | `bool` |  |
| `clearance_slip` | `dict` |  |
| `client` | `dict` |  |
| `company_member` | `dict` |  |
| `consignee` | `dict` |  |
| `consignor` | `dict` |  |
| `created_at` | `str` |  |
| `declarant` | `dict` |  |
| `document_upload_status` | `str` |  |
| `documents_presentation_requested` | `bool` |  |
| `documents_upload_requested` | `bool` |  |
| `external_id` | `str` |  |
| `form` | `str` |  |
| `goods_presentation_status` | `str` |  |
| `hrcm_status` | `str` |  |
| `id` | `str` |  |
| `invalidation_status` | `str` |  |
| `is_global_template` | `bool` |  |
| `latest_notification_item` | `str` |  |
| `latest_state` | `dict` |  |
| `lrn` | `str` |  |
| `mrn` | `str` |  |
| `name` | `str` |  |
| `partial_answers` | `bool` |  |
| `receipt` | `dict` |  |
| `refund_application_status` | `str` |  |
| `route` | `str` |  |
| `shipment_items_no` | `int` |  |
| `shipment_items_quantity_no` | `int` |  |
| `source` | `str` |  |
| `source_type` | `str` |  |
| `status` | `str` |  |
| `template` | `bool` |  |
| `template_id` | `str` |  |
| `template_properties` | `list` |  |
| `total_tax_amount` | `str` |  |
| `updated_at` | `str` |  |
| `verification_errors` | `list` |  |
| `verification_status` | `str` |  |

#### Example: Load

```python
submission = client.Submission().load({"id": "submission_id"})
```

#### Example: List

```python
submissions = client.Submission().list()
```

#### Example: Create

```python
submission = client.Submission().create({
    "answers": [],  # list
    "clearance_slip": {},  # dict
    "client": {},  # dict
    "company_member": {},  # dict
    "consignee": {},  # dict
    "consignor": {},  # dict
    "created_at": "example_created_at",  # str
    "declarant": {},  # dict
    "form": "example_form",  # str
    "latest_state": {},  # dict
    "receipt": {},  # dict
    "updated_at": "example_updated_at",  # str
})
```


### SubmissionDetail

Create an instance: `submission_detail = client.SubmissionDetail()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_external_ids` | `list` |  |
| `additional_information` | `list` |  |
| `amendment_status` | `str` |  |
| `clearance_slip` | `dict` |  |
| `client` | `dict` |  |
| `company` | `str` |  |
| `company_member` | `dict` |  |
| `consignee` | `dict` |  |
| `consignor` | `dict` |  |
| `created_at` | `str` |  |
| `declarant` | `dict` |  |
| `document_upload_status` | `str` |  |
| `documents_presentation_requested` | `bool` |  |
| `documents_upload_requested` | `bool` |  |
| `external_id` | `str` |  |
| `form` | `str` |  |
| `goods_presentation_status` | `str` |  |
| `hrcm_status` | `str` |  |
| `id` | `str` |  |
| `invalidation_status` | `str` |  |
| `is_global_template` | `bool` |  |
| `latest_notification_item` | `str` |  |
| `latest_state` | `dict` |  |
| `lrn` | `str` |  |
| `mrn` | `str` |  |
| `name` | `str` |  |
| `receipt` | `dict` |  |
| `refund_application_status` | `str` |  |
| `route` | `str` |  |
| `shipment_items_no` | `int` |  |
| `shipment_items_quantity_no` | `int` |  |
| `source` | `str` |  |
| `source_type` | `str` |  |
| `status` | `str` |  |
| `submission` | `str` |  |
| `supporting_documents` | `list` |  |
| `template` | `bool` |  |
| `total_tax_amount` | `str` |  |
| `updated_at` | `str` |  |
| `verification_errors` | `list` |  |
| `verification_status` | `str` |  |

#### Example: Create

```python
submission_detail = client.SubmissionDetail().create({
    "additional_information": [],  # list
    "clearance_slip": {},  # dict
    "client": {},  # dict
    "company": "example_company",  # str
    "company_member": {},  # dict
    "consignee": {},  # dict
    "consignor": {},  # dict
    "created_at": "example_created_at",  # str
    "declarant": {},  # dict
    "form": "example_form",  # str
    "latest_state": {},  # dict
    "receipt": {},  # dict
    "submission": "example_submission",  # str
    "updated_at": "example_updated_at",  # str
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

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── customswindow_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`customswindow_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```python
submission = client.Submission()
submission.list()

# submission.data_get() now returns the submission data from the last list
# submission.match_get() returns the last match criteria
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
