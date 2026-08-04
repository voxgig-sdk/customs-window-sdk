# CustomsWindow Python SDK Reference

Complete API reference for the CustomsWindow Python SDK.


## CustomsWindowSDK

### Constructor

```python
from customswindow_sdk import CustomsWindowSDK

client = CustomsWindowSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CustomsWindowSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = CustomsWindowSDK.test()
```


### Instance Methods

#### `BulkUpload(data=None)`

Create a new `BulkUploadEntity` instance. Pass `None` for no initial data.

#### `File(data=None)`

Create a new `FileEntity` instance. Pass `None` for no initial data.

#### `PaginatedBulkUploadListList(data=None)`

Create a new `PaginatedBulkUploadListListEntity` instance. Pass `None` for no initial data.

#### `PaginatedPartyListList(data=None)`

Create a new `PaginatedPartyListListEntity` instance. Pass `None` for no initial data.

#### `PaginatedSubmissionListList(data=None)`

Create a new `PaginatedSubmissionListListEntity` instance. Pass `None` for no initial data.

#### `Party(data=None)`

Create a new `PartyEntity` instance. Pass `None` for no initial data.

#### `Submission(data=None)`

Create a new `SubmissionEntity` instance. Pass `None` for no initial data.

#### `SubmissionDetail(data=None)`

Create a new `SubmissionDetailEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## BulkUploadEntity

```python
bulk_upload = client.BulkUpload()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `active_transport_nationality` | `str` | No |  |
| `active_transport_number` | `str` | No |  |
| `arrival_datetime` | `str` | No |  |
| `client` | `str` | Yes |  |
| `company_member` | `str` | Yes |  |
| `created_at` | `str` | Yes |  |
| `declarant` | `str` | Yes |  |
| `declarations_no` | `int` | Yes |  |
| `deleted_at` | `str` | No |  |
| `departure_datetime` | `str` | No |  |
| `errors_file` | `dict` | Yes |  |
| `external_id` | `str` | No |  |
| `failed_declarations_no` | `int` | No |  |
| `file` | `dict` | Yes |  |
| `green_routed_no` | `int` | No |  |
| `h1_fallback_template` | `str` | No |  |
| `house_transport_doc_ref` | `str` | No |  |
| `id` | `str` | No |  |
| `issue_date` | `str` | No |  |
| `mapping` | `str` | No |  |
| `orange_routed_no` | `int` | No |  |
| `parsed_declarations_no` | `int` | No |  |
| `parser` | `str` | No |  |
| `parsing_completed_at` | `str` | No |  |
| `parsing_started_at` | `str` | No |  |
| `passive_transport_nationality` | `str` | No |  |
| `passive_transport_number` | `str` | No |  |
| `processed_declarations_no` | `int` | No |  |
| `processing_ended_at` | `str` | No |  |
| `processing_started_at` | `str` | No |  |
| `receipt_generating_started_at` | `str` | No |  |
| `receipt_request_started_at` | `str` | No |  |
| `receipt_request_status` | `str` | No |  |
| `receipt_request_user` | `str` | No |  |
| `receipts_zip` | `str` | No |  |
| `red_routed_no` | `int` | No |  |
| `rejected_status_no` | `int` | No |  |
| `status` | `str` | No |  |
| `template` | `str` | No |  |
| `updated_at` | `str` | Yes |  |
| `yellow_routed_no` | `int` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.BulkUpload().create({
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.BulkUpload().list()
for bulk_upload in results:
    print(bulk_upload)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.BulkUpload().load({"id": "bulk_upload_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.BulkUpload().remove({"id": "bulk_upload_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.BulkUpload().update({
    "id": "bulk_upload_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `BulkUploadEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## FileEntity

```python
file = client.File()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `company` | `str` | Yes |  |
| `created_at` | `str` | Yes |  |
| `extension` | `str` | Yes |  |
| `file` | `str` | Yes |  |
| `id` | `str` | No |  |
| `name` | `str` | Yes |  |
| `public` | `bool` | Yes |  |
| `size` | `int` | No |  |
| `updated_at` | `str` | Yes |  |
| `url` | `str` | Yes |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.File().create({
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

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `FileEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PaginatedBulkUploadListListEntity

```python
paginated_bulk_upload_list_list = client.PaginatedBulkUploadListList()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaginatedBulkUploadListListEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PaginatedPartyListListEntity

```python
paginated_party_list_list = client.PaginatedPartyListList()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaginatedPartyListListEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PaginatedSubmissionListListEntity

```python
paginated_submission_list_list = client.PaginatedSubmissionListList()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PaginatedSubmissionListListEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PartyEntity

```python
party = client.Party()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_declaration_type` | `dict` | Yes |  |
| `address` | `dict` | Yes |  |
| `authorisation` | `dict` | Yes |  |
| `bank_detail` | `str` | No |  |
| `certificate` | `dict` | Yes |  |
| `certificate_type` | `str` | Yes |  |
| `company` | `str` | Yes |  |
| `created_at` | `str` | Yes |  |
| `customs_office_of_lodgement` | `dict` | Yes |  |
| `deleted_at` | `str` | No |  |
| `email` | `str` | No |  |
| `id` | `str` | No |  |
| `identification_number` | `str` | No |  |
| `indirect_representative` | `bool` | No |  |
| `name` | `str` | No |  |
| `nhd_last_submission_year` | `int` | No |  |
| `nhd_submission_counter` | `int` | No |  |
| `person_paying_customs_duty` | `str` | No |  |
| `phone_country_code` | `str` | No |  |
| `phone_number` | `str` | No |  |
| `preferred_payment_method` | `dict` | Yes |  |
| `signed_form` | `dict` | Yes |  |
| `type` | `str` | No |  |
| `type_of_person` | `dict` | Yes |  |
| `unlocode` | `str` | No |  |
| `updated_at` | `str` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Party().create({
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Party().list()
for party in results:
    print(party)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Party().load({"id": "party_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Party().remove({"id": "party_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Party().update({
    "id": "party_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PartyEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SubmissionEntity

```python
submission = client.Submission()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_external_id` | `list` | No |  |
| `amendment_reason` | `str` | No |  |
| `amendment_status` | `str` | No |  |
| `answer` | `list` | Yes |  |
| `bypass_restricted_code` | `bool` | No |  |
| `clearance_slip` | `dict` | Yes |  |
| `client` | `dict` | Yes |  |
| `company_member` | `dict` | Yes |  |
| `consignee` | `dict` | Yes |  |
| `consignor` | `dict` | Yes |  |
| `created_at` | `str` | Yes |  |
| `declarant` | `dict` | Yes |  |
| `document_upload_status` | `str` | No |  |
| `documents_presentation_requested` | `bool` | No |  |
| `documents_upload_requested` | `bool` | No |  |
| `external_id` | `str` | No |  |
| `form` | `str` | Yes |  |
| `goods_presentation_status` | `str` | No |  |
| `hrcm_status` | `str` | No |  |
| `id` | `str` | No |  |
| `invalidation_status` | `str` | No |  |
| `is_global_template` | `bool` | No |  |
| `latest_notification_item` | `str` | No |  |
| `latest_state` | `dict` | Yes |  |
| `lrn` | `str` | No |  |
| `mrn` | `str` | No |  |
| `name` | `str` | No |  |
| `partial_answer` | `bool` | No |  |
| `receipt` | `dict` | Yes |  |
| `refund_application_status` | `str` | No |  |
| `route` | `str` | No |  |
| `shipment_items_no` | `int` | No |  |
| `shipment_items_quantity_no` | `int` | No |  |
| `source` | `str` | No |  |
| `source_type` | `str` | No |  |
| `status` | `str` | No |  |
| `template` | `bool` | No |  |
| `template_id` | `str` | No |  |
| `template_property` | `list` | No |  |
| `total_tax_amount` | `str` | No |  |
| `updated_at` | `str` | Yes |  |
| `verification_error` | `list` | No |  |
| `verification_status` | `str` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `additional_external_id` | - | - | - | - | - |
| `amendment_reason` | - | - | - | - | - |
| `amendment_status` | - | - | - | - | - |
| `answer` | - | - | - | Yes | - |
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
| `partial_answer` | - | - | - | - | - |
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
| `template_property` | - | - | - | - | - |
| `total_tax_amount` | - | - | - | - | - |
| `updated_at` | - | - | - | - | - |
| `verification_error` | - | - | - | - | - |
| `verification_status` | - | - | - | - | - |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Submission().create({
    "answer": [],  # list
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

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Submission().list()
for submission in results:
    print(submission)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Submission().load({"id": "submission_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Submission().remove({"id": "submission_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Submission().update({
    "id": "submission_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SubmissionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SubmissionDetailEntity

```python
submission_detail = client.SubmissionDetail()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_external_id` | `list` | No |  |
| `additional_information` | `list` | Yes |  |
| `amendment_status` | `str` | No |  |
| `clearance_slip` | `dict` | Yes |  |
| `client` | `dict` | Yes |  |
| `company` | `str` | Yes |  |
| `company_member` | `dict` | Yes |  |
| `consignee` | `dict` | Yes |  |
| `consignor` | `dict` | Yes |  |
| `created_at` | `str` | Yes |  |
| `declarant` | `dict` | Yes |  |
| `document_upload_status` | `str` | No |  |
| `documents_presentation_requested` | `bool` | No |  |
| `documents_upload_requested` | `bool` | No |  |
| `external_id` | `str` | No |  |
| `form` | `str` | Yes |  |
| `goods_presentation_status` | `str` | No |  |
| `hrcm_status` | `str` | No |  |
| `id` | `str` | No |  |
| `invalidation_status` | `str` | No |  |
| `is_global_template` | `bool` | No |  |
| `latest_notification_item` | `str` | No |  |
| `latest_state` | `dict` | Yes |  |
| `lrn` | `str` | No |  |
| `mrn` | `str` | No |  |
| `name` | `str` | No |  |
| `receipt` | `dict` | Yes |  |
| `refund_application_status` | `str` | No |  |
| `route` | `str` | No |  |
| `shipment_items_no` | `int` | No |  |
| `shipment_items_quantity_no` | `int` | No |  |
| `source` | `str` | No |  |
| `source_type` | `str` | No |  |
| `status` | `str` | No |  |
| `submission` | `str` | Yes |  |
| `supporting_document` | `list` | No |  |
| `template` | `bool` | No |  |
| `total_tax_amount` | `str` | No |  |
| `updated_at` | `str` | Yes |  |
| `verification_error` | `list` | No |  |
| `verification_status` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.SubmissionDetail().create({
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

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SubmissionDetailEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = CustomsWindowSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

