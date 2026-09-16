# CustomsWindow TypeScript SDK



The TypeScript SDK for the CustomsWindow API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.BulkUpload()` — each with a small set of operations (`list`, `load`, `create`, `update`, `remove`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Also generated from this model: `go`, `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb` — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/customs-window-sdk/releases](https://github.com/voxgig-sdk/customs-window-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { CustomsWindowSDK } from '@voxgig-sdk/customs-window'

const client = new CustomsWindowSDK({
  apikey: process.env.CUSTOMS_WINDOW_APIKEY,
})
```

### 2. List bulkupload records

`list()` resolves to an array of BulkUpload ENTITIES — every operation
resolves to entities, not raw records. Iterate them directly, and call
`.data()` on one for the record it holds:

```ts
const bulkuploads = await client.BulkUpload().list()

for (const bulkupload of bulkuploads) {
  console.log(bulkupload)
}
```

### 3. Load a bulkupload

`load()` returns the entity directly and throws on failure:

```ts
try {
  const bulkupload = await client.BulkUpload().load({ id: 'example_id' })
  console.log(bulkupload)
} catch (err) {
  console.error('load failed:', err)
}
```

### 4. Create, update, and remove

```ts
// Create — returns the created BulkUpload ENTITY (.data() for the record)
const created = await client.BulkUpload().create({
  client: 'example_client',
  company_member: 'example_company_member',
  created_at: 'example_created_at',
  declarant: 'example_declarant',
  declarations_no: 1,
  errors_file: {},
  file: {},
  updated_at: 'example_updated_at',
})

// Update — the id comes off the returned entity's data()
const updated = await client.BulkUpload().update({
  id: created.data().id!,
  active_transport_nationality: 'example_active_transport_nationality',
  active_transport_number: 'example_active_transport_number',
})

// Remove
await client.BulkUpload().remove({
  id: created.data().id!,
})
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const submissions = await client.Submission().list()
  console.log(submissions)
} catch (err) {
  console.error('list failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = CustomsWindowSDK.test()

const submission = await client.Submission().list()
// submission is the entity, populated with mock response data
// — call submission.data() for the record itself
console.log(submission)
```

You can also use the instance method:

```ts
const client = new CustomsWindowSDK({ apikey: '...' })
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.Submission()

// First call runs the operation and stores its result
await entity.list()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data.id)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new CustomsWindowSDK({
  apikey: '...',
  extend: [logger],
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
cd ts && npm test
```

Live entity tests continue independent operations after errors and attempt
supported cleanup. Their final result reports failures and missing prerequisites
after the remaining work completes. The model and test inputs determine which
API operations the generated scenarios cover.


## Reference

### CustomsWindowSDK

#### Constructor

```ts
new CustomsWindowSDK(options?: {
  apikey?: string
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `BulkUpload(data?)` | `BulkUploadEntity` | Create a BulkUpload entity instance. |
| `File(data?)` | `FileEntity` | Create a File entity instance. |
| `PaginatedBulkUploadListList(data?)` | `PaginatedBulkUploadListListEntity` | Create a PaginatedBulkUploadListList entity instance. |
| `PaginatedPartyListList(data?)` | `PaginatedPartyListListEntity` | Create a PaginatedPartyListList entity instance. |
| `PaginatedSubmissionListList(data?)` | `PaginatedSubmissionListListEntity` | Create a PaginatedSubmissionListList entity instance. |
| `Party(data?)` | `PartyEntity` | Create a Party entity instance. |
| `Submission(data?)` | `SubmissionEntity` | Create a Submission entity instance. |
| `SubmissionDetail(data?)` | `SubmissionDetailEntity` | Create a SubmissionDetail entity instance. |
| `tester(testopts?, sdkopts?)` | `CustomsWindowSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `CustomsWindowSDK.test(testopts?, sdkopts?)` | `CustomsWindowSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity. |
| `update` | `update(reqdata?, ctrl?): Promise<Entity>` | Update an existing entity. |
| `remove` | `remove(reqmatch?, ctrl?): Promise<void>` | Remove an entity. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): CustomsWindowSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load`, `create` and `update` resolve to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).
- `remove` resolves to `void`.

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

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

Operations: create, list, load, remove, update.

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

Operations: create.

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

Operations: create, list, load, remove, update.

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

Operations: create, list, load, remove, update.

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

Operations: create.

API path: `/documents-request`



## Entities


### BulkUpload

Create an instance: `const bulk_upload = client.BulkUpload()`

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
| `errors_file` | `Record<string, any>` |  |
| `external_id` | `string` |  |
| `failed_declarations_no` | `number` |  |
| `file` | `Record<string, any>` | cuid-format identifier for this entity. |
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

```ts
const bulk_upload = await client.BulkUpload().load({ id: 'bulk_upload_id' })
```

#### Example: List

```ts
const bulk_uploads = await client.BulkUpload().list()
```

#### Example: Create

```ts
const bulk_upload = await client.BulkUpload().create({
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


### File

Create an instance: `const file = client.File()`

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

```ts
const file = await client.File().create({
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


### PaginatedBulkUploadListList

Create an instance: `const paginated_bulk_upload_list_list = client.PaginatedBulkUploadListList()`


### PaginatedPartyListList

Create an instance: `const paginated_party_list_list = client.PaginatedPartyListList()`


### PaginatedSubmissionListList

Create an instance: `const paginated_submission_list_list = client.PaginatedSubmissionListList()`


### Party

Create an instance: `const party = client.Party()`

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
| `additional_declaration_type` | `Record<string, any>` |  |
| `address` | `Record<string, any>` |  |
| `authorisation` | `Record<string, any>` |  |
| `bank_details` | `string` |  |
| `certificate` | `Record<string, any>` |  |
| `certificate_type` | `string` |  |
| `company` | `string` | cuid-format identifier for this entity. |
| `created_at` | `string` |  |
| `customs_office_of_lodgement` | `Record<string, any>` |  |
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
| `preferred_payment_method` | `Record<string, any>` |  |
| `signed_form` | `Record<string, any>` |  |
| `type` | `string` | * `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `co… |
| `type_of_person` | `Record<string, any>` |  |
| `unlocode` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```ts
const party = await client.Party().load({ id: 'party_id' })
```

#### Example: List

```ts
const partys = await client.Party().list()
```

#### Example: Create

```ts
const party = await client.Party().create({
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


### Submission

Create an instance: `const submission = client.Submission()`

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
| `additional_external_ids` | `any[]` |  |
| `amendment_reason` | `string` |  |
| `amendment_status` | `string` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `answers` | `any[]` |  |
| `bypass_restricted_code` | `boolean` |  |
| `clearance_slip` | `Record<string, any>` |  |
| `client` | `Record<string, any>` |  |
| `company_member` | `Record<string, any>` | cuid-format identifier for this entity. |
| `consignee` | `Record<string, any>` |  |
| `consignor` | `Record<string, any>` |  |
| `created_at` | `string` |  |
| `declarant` | `Record<string, any>` | cuid-format identifier for this entity. |
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
| `latest_state` | `Record<string, any>` |  |
| `lrn` | `string` |  |
| `mrn` | `string` |  |
| `name` | `string` |  |
| `partial_answers` | `boolean` |  |
| `receipt` | `Record<string, any>` |  |
| `refund_application_status` | `string` | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `route` | `string` | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `shipment_items_no` | `number` |  |
| `shipment_items_quantity_no` | `number` |  |
| `source` | `string` | cuid-format identifier for this entity. |
| `source_type` | `string` | * `template` - template * `automated_import` - automated_import |
| `status` | `string` | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `template` | `boolean` |  |
| `template_id` | `string` | cuid-format identifier for this entity. |
| `template_properties` | `any[]` |  |
| `total_tax_amount` | `string` |  |
| `updated_at` | `string` |  |
| `verification_errors` | `any[]` |  |
| `verification_status` | `string` | * `passed` - passed * `failed` - failed |

#### Example: Load

```ts
const submission = await client.Submission().load({ id: 'submission_id' })
```

#### Example: List

```ts
const submissions = await client.Submission().list()
```

#### Example: Create

```ts
const submission = await client.Submission().create({
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


### SubmissionDetail

Create an instance: `const submission_detail = client.SubmissionDetail()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `additional_external_ids` | `any[]` |  |
| `additional_information` | `any[]` |  |
| `amendment_status` | `string` | * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled |
| `clearance_slip` | `Record<string, any>` |  |
| `client` | `Record<string, any>` |  |
| `company` | `string` | cuid-format identifier for this entity. |
| `company_member` | `Record<string, any>` | cuid-format identifier for this entity. |
| `consignee` | `Record<string, any>` |  |
| `consignor` | `Record<string, any>` |  |
| `created_at` | `string` |  |
| `declarant` | `Record<string, any>` |  |
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
| `latest_state` | `Record<string, any>` |  |
| `lrn` | `string` |  |
| `mrn` | `string` |  |
| `name` | `string` |  |
| `receipt` | `Record<string, any>` |  |
| `refund_application_status` | `string` | * `processing` - processing * `rejected` - rejected * `accepted` - accepted |
| `route` | `string` | * `green` - green * `orange` - orange * `red` - red * `yellow` - yellow |
| `shipment_items_no` | `number` |  |
| `shipment_items_quantity_no` | `number` |  |
| `source` | `string` | cuid-format identifier for this entity. |
| `source_type` | `string` | * `template` - template * `automated_import` - automated_import |
| `status` | `string` | * `draft` - draft * `complete` - complete * `processing` - processing * `rejected` - rejected * `registered` - registered * `accepted` - accepted * `under_review` - under_review * `insufficient_funds` - insufficient_funds * `released` - re… |
| `submission` | `string` | cuid-format identifier for this entity. |
| `supporting_documents` | `any[]` |  |
| `template` | `boolean` |  |
| `total_tax_amount` | `string` |  |
| `updated_at` | `string` |  |
| `verification_errors` | `any[]` |  |
| `verification_status` | `string` | * `passed` - passed * `failed` - failed |

#### Example: Create

```ts
const submission_detail = await client.SubmissionDetail().create({
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

## Features

This SDK ships 8 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`debug`](#debug) | Request/response capture ring buffer for debugging |
| [`idempotency`](#idempotency) | Idempotency keys for safe retries of mutating operations |
| [`metrics`](#metrics) | Statistics capture: per-operation counters and latency |
| [`paging`](#paging) | Pagination signals for list operations |
| [`ratelimit`](#ratelimit) | Client-side rate limiting via a token bucket |
| [`retry`](#retry) | Automatic retry of transient failures with exponential backoff |
| [`test`](#test) | In-memory mock transport for testing without a live server |
| [`timeout`](#timeout) | Per-request timeout with transport abort |

> **Order matters for `ratelimit`, `retry`, `timeout`.** These wrap the
> transport, so each one wraps whatever is already installed: the order you
> activate them in IS the nesting order. Activating them as an ordered list
> rather than a map is what fixes that order.

### debug

Request/response capture ring buffer for debugging.

| Option | Default |
|---|---|
| `active` | `false` |
| `max` | `100` |
| `redact` | `['authorization', 'cookie', 'set-cookie', 'api-key', 'apikey', 'x-api-key', 'idempotency-key']` |

Set `feature.debug.active` to enable it, then override any of the options above.

### idempotency

Idempotency keys for safe retries of mutating operations.

| Option | Default |
|---|---|
| `active` | `false` |
| `header` | `'Idempotency-Key'` |
| `methods` | `['POST', 'PUT', 'PATCH', 'DELETE']` |
| `ops` | `['create', 'update', 'remove']` |

Set `feature.idempotency.active` to enable it, then override any of the options above.

### metrics

Statistics capture: per-operation counters and latency.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.metrics.active` to enable it, then override any of the options above.

### paging

Pagination signals for list operations.

| Option | Default |
|---|---|
| `active` | `false` |
| `afterVar` | `'after'` |
| `cursorParam` | `'cursor'` |
| `firstVar` | `'first'` |
| `limitParam` | `'limit'` |
| `pageParam` | `'page'` |
| `startPage` | `1` |

Set `feature.paging.active` to enable it, then override any of the options above.

### ratelimit

Client-side rate limiting via a token bucket.

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

Set `feature.ratelimit.active` to enable it, then override any of the options above.

`ratelimit` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### retry

Automatic retry of transient failures with exponential backoff.

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

Set `feature.retry.active` to enable it, then override any of the options above.

`retry` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### test

In-memory mock transport for testing without a live server.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.

### timeout

Per-request timeout with transport abort.

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

Set `feature.timeout.active` to enable it, then override any of the options above.

`timeout` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.


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

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **DebugFeature**: Request/response capture ring buffer for debugging
- **IdempotencyFeature**: Idempotency keys for safe retries of mutating operations
- **MetricsFeature**: Statistics capture: per-operation counters and latency
- **PagingFeature**: Pagination signals for list operations
- **RatelimitFeature**: Client-side rate limiting via a token bucket
- **RetryFeature**: Automatic retry of transient failures with exponential backoff
- **TestFeature**: In-memory mock transport for testing without a live server
- **TimeoutFeature**: Per-request timeout with transport abort

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
customs-window/
├── src/
│   ├── CustomsWindowSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { CustomsWindowSDK } from '@voxgig-sdk/customs-window'
```

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const submission = client.Submission()
await submission.list()

// submission.data() now returns the submission data from the last `list`
// submission.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
