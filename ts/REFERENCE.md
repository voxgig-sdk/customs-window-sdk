# CustomsWindow TypeScript SDK Reference

Complete API reference for the CustomsWindow TypeScript SDK.


## CustomsWindowSDK

### Constructor

```ts
new CustomsWindowSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `CustomsWindowSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = CustomsWindowSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `CustomsWindowSDK` instance in test mode.


### Instance Methods

#### `BulkUpload(data?: object)`

Create a new `BulkUpload` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `BulkUploadEntity` instance.

#### `File(data?: object)`

Create a new `File` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `FileEntity` instance.

#### `PaginatedBulkUploadListList(data?: object)`

Create a new `PaginatedBulkUploadListList` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PaginatedBulkUploadListListEntity` instance.

#### `PaginatedPartyListList(data?: object)`

Create a new `PaginatedPartyListList` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PaginatedPartyListListEntity` instance.

#### `PaginatedSubmissionListList(data?: object)`

Create a new `PaginatedSubmissionListList` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PaginatedSubmissionListListEntity` instance.

#### `Party(data?: object)`

Create a new `Party` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PartyEntity` instance.

#### `Submission(data?: object)`

Create a new `Submission` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SubmissionEntity` instance.

#### `SubmissionDetail(data?: object)`

Create a new `SubmissionDetail` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SubmissionDetailEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `CustomsWindowSDK.test()`.

**Returns:** `CustomsWindowSDK` instance in test mode.


---

## BulkUploadEntity

```ts
const bulk_upload = client.BulkUpload()
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
| `declarations_no` | `number` | Yes |  |
| `deleted_at` | `string` | No |  |
| `departure_datetime` | `string` | No |  |
| `errors_file` | `Record<string, any>` | Yes |  |
| `external_id` | `string` | No |  |
| `failed_declarations_no` | `number` | No |  |
| `file` | `Record<string, any>` | Yes | cuid-format identifier for this entity. |
| `green_routed_no` | `number` | No |  |
| `h1_fallback_template` | `string` | No | cuid-format identifier for this entity. |
| `house_transport_doc_ref` | `string` | No |  |
| `id` | `string` | No | cuid-format identifier for this entity. |
| `issue_date` | `string` | No |  |
| `mapping` | `string` | No | cuid-format identifier for this entity. |
| `orange_routed_no` | `number` | No |  |
| `parsed_declarations_no` | `number` | No |  |
| `parser` | `string` | No | * `aes_platform` - aes_platform * `cds_platform` - cds_platform * `cds_export_platform` - cds_export_platform * `g4_g3` - g4_g3 * `nhd_platform` - nhd_platform * `platform` - platform * `birds` - birds * `ics2_platform` - ics2_platform |
| `parsing_completed_at` | `string` | No |  |
| `parsing_started_at` | `string` | No |  |
| `passive_transport_nationality` | `string` | No | cuid-format identifier for this entity. |
| `passive_transport_number` | `string` | No |  |
| `processed_declarations_no` | `number` | No |  |
| `processing_ended_at` | `string` | No |  |
| `processing_started_at` | `string` | No |  |
| `receipt_generating_started_at` | `string` | No |  |
| `receipt_request_started_at` | `string` | No |  |
| `receipt_request_status` | `string` | No | * `pending` - pending * `processing` - processing * `generating` - generating * `completed` - completed |
| `receipt_request_user` | `string` | No | cuid-format identifier for this entity. |
| `receipts_zip` | `string` | No | cuid-format identifier for this entity. |
| `red_routed_no` | `number` | No |  |
| `rejected_status_no` | `number` | No |  |
| `status` | `string` | No | * `pending` - pending * `parsing` - parsing * `processing` - processing * `complete` - complete |
| `template` | `string` | No | cuid-format identifier for this entity. |
| `updated_at` | `string` | Yes |  |
| `yellow_routed_no` | `number` | No |  |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `generate_pdf` | `/bulk-uploads/{id}/generate-pdfs` | `client.BulkUpload().load({ $action: 'generate_pdf', ... })` |
| `generate_pdf` | `/bulk-uploads/{id}/generate-pdfs` | `client.BulkUpload().remove({ $action: 'generate_pdf', ... })` |
| `generate_pdf` | `/bulk-uploads/{id}/generate-pdfs` | `client.BulkUpload().update({ $action: 'generate_pdf', ... })` |

An action returns that action's OWN response, which is not necessarily a
BulkUpload record — check the API definition for its shape.

```ts
const result = await client.BulkUpload().load({
  $action: 'generate_pdf',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.BulkUpload().create({
  client: 'example_client',
  company_member: 'example_company_member',
  created_at: 'example_created_at',
  declarant: 'example_declarant',
  declarations_no: 1,
  errors_file: {},
  file: {},
  updated_at: 'example_updated_at',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.BulkUpload().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.BulkUpload().load({ id: 'bulk_upload_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.BulkUpload().remove({ id: 'bulk_upload_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.BulkUpload().update({
  id: 'bulk_upload_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `BulkUploadEntity` instance with the same client and
options.

#### `client()`

Return the parent `CustomsWindowSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## FileEntity

```ts
const file = client.File()
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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.File().create({
  company: 'example_company',
  created_at: 'example_created_at',
  extension: 'example_extension',
  file: 'example_file',
  name: 'example_name',
  public: true,
  updated_at: 'example_updated_at',
  url: 'example_url',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `FileEntity` instance with the same client and
options.

#### `client()`

Return the parent `CustomsWindowSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PaginatedBulkUploadListListEntity

```ts
const paginated_bulk_upload_list_list = client.PaginatedBulkUploadListList()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PaginatedBulkUploadListListEntity` instance with the same client and
options.

#### `client()`

Return the parent `CustomsWindowSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PaginatedPartyListListEntity

```ts
const paginated_party_list_list = client.PaginatedPartyListList()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PaginatedPartyListListEntity` instance with the same client and
options.

#### `client()`

Return the parent `CustomsWindowSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PaginatedSubmissionListListEntity

```ts
const paginated_submission_list_list = client.PaginatedSubmissionListList()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PaginatedSubmissionListListEntity` instance with the same client and
options.

#### `client()`

Return the parent `CustomsWindowSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PartyEntity

```ts
const party = client.Party()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_declaration_type` | `Record<string, any>` | Yes |  |
| `address` | `Record<string, any>` | Yes |  |
| `authorisation` | `Record<string, any>` | Yes |  |
| `bank_details` | `string` | No |  |
| `certificate` | `Record<string, any>` | Yes |  |
| `certificate_type` | `string` | Yes |  |
| `company` | `string` | Yes | cuid-format identifier for this entity. |
| `created_at` | `string` | Yes |  |
| `customs_office_of_lodgement` | `Record<string, any>` | Yes |  |
| `deleted_at` | `string` | No |  |
| `email` | `string` | No |  |
| `id` | `string` | No | cuid-format identifier for this entity. |
| `identification_number` | `string` | No |  |
| `indirect_representative` | `boolean` | No |  |
| `name` | `string` | No |  |
| `nhd_last_submission_year` | `number` | No |  |
| `nhd_submission_counter` | `number` | No |  |
| `person_paying_customs_duty` | `string` | No |  |
| `phone_country_code` | `string` | No |  |
| `phone_number` | `string` | No |  |
| `preferred_payment_method` | `Record<string, any>` | Yes |  |
| `signed_form` | `Record<string, any>` | Yes |  |
| `type` | `string` | No | * `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `co… |
| `type_of_person` | `Record<string, any>` | Yes |  |
| `unlocode` | `string` | No |  |
| `updated_at` | `string` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Party().create({
  additional_declaration_type: {},
  address: {},
  authorisation: {},
  certificate: {},
  certificate_type: 'example_certificate_type',
  company: 'example_company',
  created_at: 'example_created_at',
  customs_office_of_lodgement: {},
  preferred_payment_method: {},
  signed_form: {},
  type_of_person: {},
  updated_at: 'example_updated_at',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Party().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Party().load({ id: 'party_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Party().remove({ id: 'party_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Party().update({
  id: 'party_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PartyEntity` instance with the same client and
options.

#### `client()`

Return the parent `CustomsWindowSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SubmissionEntity

```ts
const submission = client.Submission()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_external_ids` | `any[]` | No |  |
| `amendment_reason` | `string` | No |  |
| `amendment_status` | `string` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `answers` | `any[]` | Yes |  |
| `bypass_restricted_code` | `boolean` | No |  |
| `clearance_slip` | `Record<string, any>` | Yes |  |
| `client` | `Record<string, any>` | Yes |  |
| `company_member` | `Record<string, any>` | Yes | cuid-format identifier for this entity. |
| `consignee` | `Record<string, any>` | Yes |  |
| `consignor` | `Record<string, any>` | Yes |  |
| `created_at` | `string` | Yes |  |
| `declarant` | `Record<string, any>` | Yes | cuid-format identifier for this entity. |
| `document_upload_status` | `string` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `documents_presentation_requested` | `boolean` | No |  |
| `documents_upload_requested` | `boolean` | No |  |
| `external_id` | `string` | No |  |
| `form` | `string` | Yes | cuid-format identifier for this entity. |
| `goods_presentation_status` | `string` | No | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested |
| `hrcm_status` | `string` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `id` | `string` | No | cuid-format identifier for this entity. |
| `invalidation_status` | `string` | No | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled |
| `is_global_template` | `boolean` | No |  |
| `latest_notification_item` | `string` | No | cuid-format identifier for this entity. |
| `latest_state` | `Record<string, any>` | Yes |  |
| `lrn` | `string` | No |  |
| `mrn` | `string` | No |  |
| `name` | `string` | No |  |
| `partial_answers` | `boolean` | No |  |
| `receipt` | `Record<string, any>` | Yes |  |
| `refund_application_status` | `string` | No | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `route` | `string` | No | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `shipment_items_no` | `number` | No |  |
| `shipment_items_quantity_no` | `number` | No |  |
| `source` | `string` | No | cuid-format identifier for this entity. |
| `source_type` | `string` | No | * `template` - template * `automated_import` - automated_import |
| `status` | `string` | No | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `template` | `boolean` | No |  |
| `template_id` | `string` | No | cuid-format identifier for this entity. |
| `template_properties` | `any[]` | No |  |
| `total_tax_amount` | `string` | No |  |
| `updated_at` | `string` | Yes |  |
| `verification_errors` | `any[]` | No |  |
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

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `refund` | `/submissions/{id}/refund` | `client.Submission().create({ $action: 'refund', ... })` |
| `retrieve` | `/submissions/retrieve` | `client.Submission().create({ $action: 'retrieve', ... })` |
| `clearance_slip` | `/submissions/{id}/clearance-slip` | `client.Submission().load({ $action: 'clearance_slip', ... })` |
| `notification_read` | `/submissions/{id}/notification-read` | `client.Submission().load({ $action: 'notification_read', ... })` |
| `pbn_applicable` | `/submissions/{id}/pbn-applicable` | `client.Submission().load({ $action: 'pbn_applicable', ... })` |
| `receipt` | `/submissions/{id}/receipt` | `client.Submission().load({ $action: 'receipt', ... })` |
| `refund` | `/submissions/{id}/refund` | `client.Submission().load({ $action: 'refund', ... })` |
| `retrieve` | `/submissions/retrieve` | `client.Submission().load({ $action: 'retrieve', ... })` |
| `clearance_slip` | `/submissions/{id}/clearance-slip` | `client.Submission().remove({ $action: 'clearance_slip', ... })` |
| `notification_read` | `/submissions/{id}/notification-read` | `client.Submission().remove({ $action: 'notification_read', ... })` |
| `pbn_applicable` | `/submissions/{id}/pbn-applicable` | `client.Submission().remove({ $action: 'pbn_applicable', ... })` |
| `receipt` | `/submissions/{id}/receipt` | `client.Submission().remove({ $action: 'receipt', ... })` |
| `clearance_slip` | `/submissions/{id}/clearance-slip` | `client.Submission().update({ $action: 'clearance_slip', ... })` |
| `notification_read` | `/submissions/{id}/notification-read` | `client.Submission().update({ $action: 'notification_read', ... })` |
| `pbn_applicable` | `/submissions/{id}/pbn-applicable` | `client.Submission().update({ $action: 'pbn_applicable', ... })` |
| `receipt` | `/submissions/{id}/receipt` | `client.Submission().update({ $action: 'receipt', ... })` |

An action returns that action's OWN response, which is not necessarily a
Submission record — check the API definition for its shape.

```ts
const result = await client.Submission().create({
  $action: 'refund',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Submission().create({
  answers: [],
  clearance_slip: {},
  client: {},
  company_member: {},
  consignee: {},
  consignor: {},
  created_at: 'example_created_at',
  declarant: {},
  form: 'example_form',
  latest_state: {},
  receipt: {},
  updated_at: 'example_updated_at',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Submission().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Submission().load({ id: 'submission_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Submission().remove({ id: 'submission_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Submission().update({
  id: 'submission_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SubmissionEntity` instance with the same client and
options.

#### `client()`

Return the parent `CustomsWindowSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SubmissionDetailEntity

```ts
const submission_detail = client.SubmissionDetail()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `additional_external_ids` | `any[]` | No |  |
| `additional_information` | `any[]` | Yes |  |
| `amendment_status` | `string` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `clearance_slip` | `Record<string, any>` | Yes |  |
| `client` | `Record<string, any>` | Yes |  |
| `company` | `string` | Yes | cuid-format identifier for this entity. |
| `company_member` | `Record<string, any>` | Yes | cuid-format identifier for this entity. |
| `consignee` | `Record<string, any>` | Yes |  |
| `consignor` | `Record<string, any>` | Yes |  |
| `created_at` | `string` | Yes |  |
| `declarant` | `Record<string, any>` | Yes |  |
| `document_upload_status` | `string` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `documents_presentation_requested` | `boolean` | No |  |
| `documents_upload_requested` | `boolean` | No |  |
| `external_id` | `string` | No |  |
| `form` | `string` | Yes |  |
| `goods_presentation_status` | `string` | No | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested |
| `hrcm_status` | `string` | No | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `id` | `string` | No | cuid-format identifier for this entity. |
| `invalidation_status` | `string` | No | * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested * `cancelled` - cancelled |
| `is_global_template` | `boolean` | No |  |
| `latest_notification_item` | `string` | No | cuid-format identifier for this entity. |
| `latest_state` | `Record<string, any>` | Yes |  |
| `lrn` | `string` | No |  |
| `mrn` | `string` | No |  |
| `name` | `string` | No |  |
| `receipt` | `Record<string, any>` | Yes |  |
| `refund_application_status` | `string` | No | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `route` | `string` | No | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `shipment_items_no` | `number` | No |  |
| `shipment_items_quantity_no` | `number` | No |  |
| `source` | `string` | No | cuid-format identifier for this entity. |
| `source_type` | `string` | No | * `template` - template * `automated_import` - automated_import |
| `status` | `string` | No | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `submission` | `string` | Yes | cuid-format identifier for this entity. |
| `supporting_documents` | `any[]` | No |  |
| `template` | `boolean` | No |  |
| `total_tax_amount` | `string` | No |  |
| `updated_at` | `string` | Yes |  |
| `verification_errors` | `any[]` | No |  |
| `verification_status` | `string` | No | * `passed` - passed * `failed` - failed |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.SubmissionDetail().create({
  additional_information: [],
  clearance_slip: {},
  client: {},
  company: 'example_company',
  company_member: {},
  consignee: {},
  consignor: {},
  created_at: 'example_created_at',
  declarant: {},
  form: 'example_form',
  latest_state: {},
  receipt: {},
  submission: 'example_submission',
  updated_at: 'example_updated_at',
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SubmissionDetailEntity` instance with the same client and
options.

#### `client()`

Return the parent `CustomsWindowSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new CustomsWindowSDK({
  feature: {
    test: { active: true },
  }
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

Options above are those the model carries a default for. A feature may
also accept callback options — a `sink` to receive each record, for
instance — which have no default and are covered in the full feature
reference.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

