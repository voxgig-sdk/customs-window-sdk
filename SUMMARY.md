# API Reference

The API Reference.

## Start here

This guide introduces the API, the client libraries, and the companion tools in this repository. Start with the API capabilities, choose a client for your application, and use the linked reference when you need exact request and response details.

The selected API surface contains 8 entities and 36 HTTP routes. There are 6 SDK targets and 2 companion tools.

An entity groups related API operations. An operation can have several routes with different inputs or authentication requirements. The SDK exposes the entity and its operations using the conventions of the selected language.

## What the API provides

### [BulkUpload](docs/api/bulk_upload.html)

Results: No response body.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `active_transport_nationality`: cuid-format identifier for this entity.
- `client`: cuid-format identifier for this entity.
- `company_member`: cuid-format identifier for this entity.
- `declarant`: cuid-format identifier for this entity.
- `file`: cuid-format identifier for this entity.

### [File](docs/api/file.html)

SDK operations: `create`.

Key fields to recognise:

- `company`: cuid-format identifier for this entity.
- `id`: cuid-format identifier for this entity.

### [PaginatedBulkUploadListList](docs/api/paginated_bulk_upload_list_list.html)

SDK operations: .

### [PaginatedPartyListList](docs/api/paginated_party_list_list.html)

SDK operations: .

### [PaginatedSubmissionListList](docs/api/paginated_submission_list_list.html)

SDK operations: .

### [Party](docs/api/party.html)

Results: No response body.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `company`: cuid-format identifier for this entity.
- `id`: cuid-format identifier for this entity.
- `type`: * `exporter` - exporter * `importer` - importer * `buyer` - buyer * `seller` - seller * `representative` - representative * `declarant` - declarant * `owner` - owner * `authorisation_holder` - authorisation_holder * `client` - client * `contact` - contact * `consignee` - consignee * `consignor` - consignor * `carrier` - carrier

### [Submission](docs/api/submission.html)

Results: No response body.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `amendment_status`: * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled
- `company_member`: cuid-format identifier for this entity.
- `declarant`: cuid-format identifier for this entity.
- `document_upload_status`: * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled
- `form`: cuid-format identifier for this entity.

### [SubmissionDetail](docs/api/submission_detail.html)

SDK operations: `create`.

Key fields to recognise:

- `amendment_status`: * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled
- `company`: cuid-format identifier for this entity.
- `company_member`: cuid-format identifier for this entity.
- `document_upload_status`: * `requested` - requested * `draft` - draft * `complete` - complete * `processing` - processing * `accepted` - accepted * `rejected` - rejected * `cancelled` - cancelled
- `goods_presentation_status`: * `in_review` - in_review * `rejected` - rejected * `accepted` - accepted * `requested` - requested

### Route map

Use this map to locate a capability. Consult the entity reference before supplying request data; routes for the same operation can require different fields.

| Entity | SDK operation | HTTP route | Authentication |
| --- | --- | --- | --- |
| [BulkUpload](docs/api/bulk_upload.html) | `create` | `POST /bulk-uploads` | Required |
| [BulkUpload](docs/api/bulk_upload.html) | `list` | `GET /bulk-uploads` | Required |
| [BulkUpload](docs/api/bulk_upload.html) | `load` | `GET /bulk-uploads/{id}` | Required |
| [BulkUpload](docs/api/bulk_upload.html) | `load` | `GET /bulk-uploads/{id}/generate-pdfs` | Required |
| [BulkUpload](docs/api/bulk_upload.html) | `remove` | `DELETE /bulk-uploads/{id}` | Required |
| [BulkUpload](docs/api/bulk_upload.html) | `remove` | `DELETE /bulk-uploads/{id}/generate-pdfs` | Required |
| [BulkUpload](docs/api/bulk_upload.html) | `update` | `PATCH /bulk-uploads/{id}` | Required |
| [BulkUpload](docs/api/bulk_upload.html) | `update` | `PATCH /bulk-uploads/{id}/generate-pdfs` | Required |
| [File](docs/api/file.html) | `create` | `POST /files` | Required |
| [Party](docs/api/party.html) | `create` | `POST /parties` | Required |
| [Party](docs/api/party.html) | `list` | `GET /parties` | Required |
| [Party](docs/api/party.html) | `load` | `GET /parties/{id}` | Required |
| [Party](docs/api/party.html) | `remove` | `DELETE /parties/{id}` | Required |
| [Party](docs/api/party.html) | `update` | `PATCH /parties/{id}` | Required |
| [Submission](docs/api/submission.html) | `create` | `POST /submissions/{id}/refund` | Required |
| [Submission](docs/api/submission.html) | `create` | `POST /submissions` | Required |
| [Submission](docs/api/submission.html) | `create` | `POST /submissions/retrieve` | Required |
| [Submission](docs/api/submission.html) | `list` | `GET /submissions` | Required |
| [Submission](docs/api/submission.html) | `load` | `GET /submissions/{id}` | Required |
| [Submission](docs/api/submission.html) | `load` | `GET /submissions/{id}/clearance-slip` | Required |
| [Submission](docs/api/submission.html) | `load` | `GET /submissions/{id}/notification-read` | Required |
| [Submission](docs/api/submission.html) | `load` | `GET /submissions/{id}/pbn-applicable` | Required |
| [Submission](docs/api/submission.html) | `load` | `GET /submissions/{id}/receipt` | Required |
| [Submission](docs/api/submission.html) | `load` | `GET /submissions/{id}/refund` | Required |
| [Submission](docs/api/submission.html) | `load` | `GET /submissions/retrieve` | Required |
| [Submission](docs/api/submission.html) | `remove` | `DELETE /submissions/{id}` | Required |
| [Submission](docs/api/submission.html) | `remove` | `DELETE /submissions/{id}/clearance-slip` | Required |
| [Submission](docs/api/submission.html) | `remove` | `DELETE /submissions/{id}/notification-read` | Required |
| [Submission](docs/api/submission.html) | `remove` | `DELETE /submissions/{id}/pbn-applicable` | Required |
| [Submission](docs/api/submission.html) | `remove` | `DELETE /submissions/{id}/receipt` | Required |
| [Submission](docs/api/submission.html) | `update` | `PATCH /submissions/{id}` | Required |
| [Submission](docs/api/submission.html) | `update` | `PATCH /submissions/{id}/clearance-slip` | Required |
| [Submission](docs/api/submission.html) | `update` | `PATCH /submissions/{id}/notification-read` | Required |
| [Submission](docs/api/submission.html) | `update` | `PATCH /submissions/{id}/pbn-applicable` | Required |
| [Submission](docs/api/submission.html) | `update` | `PATCH /submissions/{id}/receipt` | Required |
| [SubmissionDetail](docs/api/submission_detail.html) | `create` | `POST /documents-request` | Required |

## Connect to the API

- Customs Window production API: `https://api.customswindow.com`

The default credential is sent in the `Basic Auth` header.

Basic authentication with required secret-key

Check authentication for the route you plan to call. A route that declares no authentication can be used without credentials; this does not change the requirements of other routes. Keep credentials in environment variables or a configured secret provider, and keep them out of source control and logs.

## Make a first request

1. Choose the API server and an operation that matches your task.
2. Check the operation’s required input and authentication. Use values valid for your account and environment.
3. Send one request and inspect the returned data before adding retries, concurrency, or a larger batch.

For an SDK call, install or build the chosen client, create a client instance with its documented configuration, and call the required entity operation. Language references describe the argument shape, asynchronous behaviour, and returned values.

## Choose an SDK

Choose the language already used by your application or service. The clients represent the same API model, while package setup, naming, and return types follow each language. Check the selected client’s reference and tests before integrating it into an existing application.

| Client | Repository directory | Distribution |
| --- | --- | --- |
| [Golang](docs/sdks/go.html) | `go/` | Build from source |
| [Lua](docs/sdks/lua.html) | `lua/` | Build from source |
| [PHP](docs/sdks/php.html) | `php/` | Build from source |
| [Python](docs/sdks/py.html) | `py/` | Build from source |
| [Ruby](docs/sdks/rb.html) | `rb/` | Build from source |
| [TypeScript](docs/sdks/ts.html) | `ts/` | Build from source |

Build-from-source entries are not marked as published in the project model. Follow the build instructions in that target’s README, then consume the resulting package using your language’s local dependency mechanism. Published entries give the installation command recorded for that client.

## Companion tools

These targets provide another way to use the API. Their available commands or tools can cover a smaller set of operations than the client libraries.

### [Go CLI](docs/tools/go-cli.html)

Use the command-line interface for shell-based tasks and scripts.

Repository directory: `go-cli/`. Not published. Build from the go-cli directory.


### [Go MCP server](docs/tools/go-mcp.html)

Use the MCP server to expose supported API operations to an MCP client.

Repository directory: `go-mcp/`. Not published. Build from the go-mcp directory.

- `customs-window_list`: List records for an entity. Supported entities: `bulk_upload`, `party`, `submission`.
- `customs-window_load`: Load one record for an entity. Supported entities: `bulk_upload`, `party`, `submission`.

## Operational features

Features supply behaviour around API calls, such as request handling, diagnostics, or local testing. Inclusion in this project does not mean a feature is enabled at runtime. Check the selected SDK’s supported features and configuration defaults, then enable the behaviour your application needs.

- [`debug`](docs/features/debug.html): Request/response capture ring buffer for debugging
- [`idempotency`](docs/features/idempotency.html): Idempotency keys for safe retries of mutating operations
- [`metrics`](docs/features/metrics.html): Statistics capture: per-operation counters and latency
- [`paging`](docs/features/paging.html): Pagination signals for list operations
- [`ratelimit`](docs/features/ratelimit.html): Client-side rate limiting via a token bucket
- [`retry`](docs/features/retry.html): Automatic retry of transient failures with exponential backoff
- [`test`](docs/features/test.html): In-memory mock transport for testing without a live server
- [`timeout`](docs/features/timeout.html): Per-request timeout with transport abort

Start with the default client configuration. Add request limits and diagnostics as needed, test error paths, and review retry behaviour before using operations that change data. A retry can repeat an operation unless the API provides a suitable guarantee.

## Continue with the documentation

- Follow the [first-call guide](docs/guides/first-call.html) for the setup sequence.
- Read the [authentication guide](docs/guides/authentication.html) before using protected routes.
- Use the [API reference](docs/api/index.html) for request schemas, response formats, and status codes.
- Check the chosen SDK or companion tool reference for its configuration and supported operations.

