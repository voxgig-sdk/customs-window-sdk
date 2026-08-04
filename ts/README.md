# CustomsWindow TypeScript SDK



The TypeScript SDK for the CustomsWindow API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.BulkUpload()` — each with a small set of operations (`list`, `load`, `create`, `update`, `remove`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
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

`list()` resolves to an array of BulkUpload objects — iterate it directly:

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
// Create — returns the created BulkUpload
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

// Update — the id comes straight off the returned entity
const updated = await client.BulkUpload().update({
  id: created.id!,
  active_transport_nationality: 'example_active_transport_nationality',
  active_transport_number: 'example_active_transport_number',
})

// Remove
await client.BulkUpload().remove({
  id: created.id!,
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
// submission is a bare entity populated with mock response data
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

Operations: create, list, load, remove, update.

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
| `bank_detail` |  |
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

Operations: create, list, load, remove, update.

API path: `/parties`

#### Submission

| Field | Description |
| --- | --- |
| `additional_external_id` |  |
| `amendment_reason` |  |
| `amendment_status` |  |
| `answer` |  |
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
| `partial_answer` |  |
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
| `template_property` |  |
| `total_tax_amount` |  |
| `updated_at` |  |
| `verification_error` |  |
| `verification_status` |  |

Operations: create, list, load, remove, update.

API path: `/submissions/{id}/refund`

#### SubmissionDetail

| Field | Description |
| --- | --- |
| `additional_external_id` |  |
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
| `supporting_document` |  |
| `template` |  |
| `total_tax_amount` |  |
| `updated_at` |  |
| `verification_error` |  |
| `verification_status` |  |

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
| `active_transport_nationality` | `string` |  |
| `active_transport_number` | `string` |  |
| `arrival_datetime` | `string` |  |
| `client` | `string` |  |
| `company_member` | `string` |  |
| `created_at` | `string` |  |
| `declarant` | `string` |  |
| `declarations_no` | `number` |  |
| `deleted_at` | `string` |  |
| `departure_datetime` | `string` |  |
| `errors_file` | `Record<string, any>` |  |
| `external_id` | `string` |  |
| `failed_declarations_no` | `number` |  |
| `file` | `Record<string, any>` |  |
| `green_routed_no` | `number` |  |
| `h1_fallback_template` | `string` |  |
| `house_transport_doc_ref` | `string` |  |
| `id` | `string` |  |
| `issue_date` | `string` |  |
| `mapping` | `string` |  |
| `orange_routed_no` | `number` |  |
| `parsed_declarations_no` | `number` |  |
| `parser` | `string` |  |
| `parsing_completed_at` | `string` |  |
| `parsing_started_at` | `string` |  |
| `passive_transport_nationality` | `string` |  |
| `passive_transport_number` | `string` |  |
| `processed_declarations_no` | `number` |  |
| `processing_ended_at` | `string` |  |
| `processing_started_at` | `string` |  |
| `receipt_generating_started_at` | `string` |  |
| `receipt_request_started_at` | `string` |  |
| `receipt_request_status` | `string` |  |
| `receipt_request_user` | `string` |  |
| `receipts_zip` | `string` |  |
| `red_routed_no` | `number` |  |
| `rejected_status_no` | `number` |  |
| `status` | `string` |  |
| `template` | `string` |  |
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
| `company` | `string` |  |
| `created_at` | `string` |  |
| `extension` | `string` |  |
| `file` | `string` |  |
| `id` | `string` |  |
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
| `bank_detail` | `string` |  |
| `certificate` | `Record<string, any>` |  |
| `certificate_type` | `string` |  |
| `company` | `string` |  |
| `created_at` | `string` |  |
| `customs_office_of_lodgement` | `Record<string, any>` |  |
| `deleted_at` | `string` |  |
| `email` | `string` |  |
| `id` | `string` |  |
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
| `type` | `string` |  |
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
| `additional_external_id` | `any[]` |  |
| `amendment_reason` | `string` |  |
| `amendment_status` | `string` |  |
| `answer` | `any[]` |  |
| `bypass_restricted_code` | `boolean` |  |
| `clearance_slip` | `Record<string, any>` |  |
| `client` | `Record<string, any>` |  |
| `company_member` | `Record<string, any>` |  |
| `consignee` | `Record<string, any>` |  |
| `consignor` | `Record<string, any>` |  |
| `created_at` | `string` |  |
| `declarant` | `Record<string, any>` |  |
| `document_upload_status` | `string` |  |
| `documents_presentation_requested` | `boolean` |  |
| `documents_upload_requested` | `boolean` |  |
| `external_id` | `string` |  |
| `form` | `string` |  |
| `goods_presentation_status` | `string` |  |
| `hrcm_status` | `string` |  |
| `id` | `string` |  |
| `invalidation_status` | `string` |  |
| `is_global_template` | `boolean` |  |
| `latest_notification_item` | `string` |  |
| `latest_state` | `Record<string, any>` |  |
| `lrn` | `string` |  |
| `mrn` | `string` |  |
| `name` | `string` |  |
| `partial_answer` | `boolean` |  |
| `receipt` | `Record<string, any>` |  |
| `refund_application_status` | `string` |  |
| `route` | `string` |  |
| `shipment_items_no` | `number` |  |
| `shipment_items_quantity_no` | `number` |  |
| `source` | `string` |  |
| `source_type` | `string` |  |
| `status` | `string` |  |
| `template` | `boolean` |  |
| `template_id` | `string` |  |
| `template_property` | `any[]` |  |
| `total_tax_amount` | `string` |  |
| `updated_at` | `string` |  |
| `verification_error` | `any[]` |  |
| `verification_status` | `string` |  |

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
  answer: [],
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
| `additional_external_id` | `any[]` |  |
| `additional_information` | `any[]` |  |
| `amendment_status` | `string` |  |
| `clearance_slip` | `Record<string, any>` |  |
| `client` | `Record<string, any>` |  |
| `company` | `string` |  |
| `company_member` | `Record<string, any>` |  |
| `consignee` | `Record<string, any>` |  |
| `consignor` | `Record<string, any>` |  |
| `created_at` | `string` |  |
| `declarant` | `Record<string, any>` |  |
| `document_upload_status` | `string` |  |
| `documents_presentation_requested` | `boolean` |  |
| `documents_upload_requested` | `boolean` |  |
| `external_id` | `string` |  |
| `form` | `string` |  |
| `goods_presentation_status` | `string` |  |
| `hrcm_status` | `string` |  |
| `id` | `string` |  |
| `invalidation_status` | `string` |  |
| `is_global_template` | `boolean` |  |
| `latest_notification_item` | `string` |  |
| `latest_state` | `Record<string, any>` |  |
| `lrn` | `string` |  |
| `mrn` | `string` |  |
| `name` | `string` |  |
| `receipt` | `Record<string, any>` |  |
| `refund_application_status` | `string` |  |
| `route` | `string` |  |
| `shipment_items_no` | `number` |  |
| `shipment_items_quantity_no` | `number` |  |
| `source` | `string` |  |
| `source_type` | `string` |  |
| `status` | `string` |  |
| `submission` | `string` |  |
| `supporting_document` | `any[]` |  |
| `template` | `boolean` |  |
| `total_tax_amount` | `string` |  |
| `updated_at` | `string` |  |
| `verification_error` | `any[]` |  |
| `verification_status` | `string` |  |

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

- **TestFeature**: In-memory mock transport for testing without a live server

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
