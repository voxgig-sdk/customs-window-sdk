# customs-window-cli

boru-driven command-line client **and** interactive REPL for the CustomsWindow
SDK. Each command line is parsed as a single [boru](https://github.com/boru-lang/boru)
expression and evaluated against the live API; run it with no arguments to drop
into a REPL. Built on `github.com/boru-lang/boru/eng/go` and the sibling Go SDK
at `../go`.

## Examples

```sh
# 1. Build a native binary (-> dist/<os>-<arch>/customs-window-cli)
make build

# 2. See usage (words, entities, env vars)
./customs-window-cli --help

# 3. Provide credentials once, via the environment
export CUSTOMS_WINDOW_APIKEY=sk_live_xxx

# 4. Each command line is ONE boru expression, run against the API:
./customs-window-cli list bulk_upload
./customs-window-cli load 1 bulk_upload            # {id:1} shorthand
./customs-window-cli load '{id:1}' bulk_upload       # explicit match map
./customs-window-cli update '{name:"x"}' bulk_upload
./customs-window-cli list file

# 5. Override the API base URL for a single call
CUSTOMS_WINDOW_BASE=https://api.example.com ./customs-window-cli list bulk_upload

# 6. No arguments -> interactive REPL
./customs-window-cli
customs-window> list bulk_upload
customs-window> /quit
```

> The rest of this guide follows the [Diátaxis](https://diataxis.fr) framework:
> a hands-on **Tutorial**, task-focused **How-to guides**, a factual
> **Reference**, and background **Explanation**.

## Tutorial: your first query in under a minute

1. **Build the binary.** From this `go-cli/` directory:

   ```sh
   make build          # -> dist/<os>-<arch>/customs-window-cli
   ```

2. **Set your API key** (read from the environment):

   ```sh
   export CUSTOMS_WINDOW_APIKEY=sk_live_xxx
   ```

3. **Run a query.** Evaluate an boru expression against the API (or run with no
   arguments to open the REPL):

   ```sh
   ./dist/*/customs-window-cli list bulk_upload
   ```

4. **Go interactive.** Run the binary with no arguments to open the REPL, then
   type `/help` for the word and entity lists and `/quit` to leave.

That is the whole loop: *build → set key → evaluate boru expressions*.

## How-to guides

### List the records of an entity

```sh
./customs-window-cli list bulk_upload
```

`list <entity>` returns the first page of records. `<entity>` is a bareword —
it is auto-quoted as an boru atom, so no quotes are needed.

### Load a single record

```sh
./customs-window-cli load 1 bulk_upload          # scalar shorthand for {id:1}
./customs-window-cli load '{id:1}' bulk_upload     # explicit match map
```

The query is either a **scalar** (`1`, treated as `{id:1}`) or a **match map**
(`{id:1}`, `{slug:"acme"}`). Quote the map so your shell passes it through intact.

### Update a record

```sh
./customs-window-cli update '{id:1,name:"new"}' bulk_upload
```

The match map carries both the selector and the new field values; the updated
record is printed back.

### Authenticate and choose an environment

Configuration is read from the environment — nothing is written to disk:

```sh
export CUSTOMS_WINDOW_APIKEY=sk_live_xxx            # API key
export CUSTOMS_WINDOW_BASE=https://api.example.com  # optional: override the API base URL
./customs-window-cli list bulk_upload
```

Both are injectable by a secrets vault, so the key never has to be typed inline.

### Explore interactively with the REPL

Run with no arguments to open a REPL (prompt `customs-window>`). Each line is
evaluated as its own boru expression:

```text
$ ./customs-window-cli
customs-window> list bulk_upload
customs-window> /help
customs-window> /quit
```

### Cross-compile release binaries

```sh
make build       # native binary for this machine
make build-all   # linux/darwin/windows x amd64/arm64, under dist/<os>-<arch>/
```

### Discover the available entities

`/help` in the REPL prints the full entity list, or see [Entities](#entities)
below — this SDK exposes 8 entities.

## Reference

### Words

The CLI registers these boru words, each bound to the SDK:

| Word     | Signatures                                    | Returns                        |
|----------|-----------------------------------------------|--------------------------------|
| `list`   | `list <entity>` · `list <query> <entity>`     | First page of records          |
| `load`   | `load <entity>` · `load <query> <entity>`     | A single record                |
| `update` | `update <query> <entity>`                     | Update a record, return it     |

- `<entity>` is a bareword, auto-quoted as an boru atom (e.g. `bulk_upload`).
- `<query>` is either a **Map** (`{id:1}`) or a **Scalar** (`1`, treated as
  `{id:1}`). A scalar is always wrapped as `{id:<value>}`.

### Environment variables

| Variable | Purpose |
|----------|---------|
| `CUSTOMS_WINDOW_APIKEY` | API key sent with every request. |
| `CUSTOMS_WINDOW_BASE` | Optional override of the API base URL. |

Unset variables fall back to the SDK's built-in defaults.

### CLI flags

- `--help` / `-h` — print usage (words, entities, env vars) and exit.

### REPL commands

Meta-commands use the `/` prefix (everything else on a line is evaluated as boru):

- `/quit` / `/q` / `/exit` — exit the REPL
- `/help` / `/h` / `/?`     — show the word list, entity list and meta commands

### Exit codes

| Code | Meaning |
|------|---------|
| `0` | Success (also the normal REPL exit). |
| `1` | Parse error, word-registration error, or an API/evaluation error. |

### Build targets

| Target | Result |
|--------|--------|
| `make build` | Native binary at `dist/<os>-<arch>/customs-window-cli`. |
| `make build-all` | linux/darwin/windows x amd64/arm64, each under its own `dist/<os>-<arch>/`. |
| `make clean` | Remove `dist/` and any stray binaries. |

### Entities

The 8 entities this SDK exposes (any is valid as `<entity>`):

bulk_upload file paginated_bulk_upload_list_list paginated_party_list_list paginated_submission_list_list party submission submission_detail

## Explanation

### Why boru?

The whole command line is one [boru](https://github.com/boru-lang/boru) expression,
not a fixed `verb --flag` grammar. That means the same binary works one-shot
(`./customs-window-cli <expr>`) and interactively (the REPL), and expressions compose the
same way in both. `list` / `load` / `update` are ordinary boru *words* bound to
the SDK — adding SDK operations is adding words, not re-parsing flags.

### How it is wired

`main.go` builds the SDK client (configured from the environment), creates an
boru registry, and `words.go` registers `list` / `load` / `update` as native
words that dispatch on the entity atom and call the sibling Go SDK at `../go`.
Results are unwrapped from their `Entity` wrappers to plain data before being
printed.

### Output format

Each result value is printed as its boru string form (a JSON-like rendering of
the record or list of records). One-shot mode prints to stdout; errors go to
stderr with a non-zero exit code.

## Generated by

sdkgen `go-cli` target. See the target source under `.sdk/src/cmp/go-cli/` in
this repo, or upstream at
`github.com/voxgig/sdkgen/project/.sdk/src/cmp/go-cli/`.
