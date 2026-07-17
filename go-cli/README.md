# nofrixion-cli

AQL-driven command-line client **and** interactive REPL for the Nofrixion
SDK. Each command line is parsed as a single [AQL](https://github.com/aql-lang/aql)
expression and evaluated against the live API; run it with no arguments to drop
into a REPL. Built on `github.com/aql-lang/aql/eng/go` and the sibling Go SDK
at `../go`.

## Examples

```sh
# 1. Build a native binary (-> dist/<os>-<arch>/nofrixion-cli)
make build

# 2. Provide credentials once, via the environment
export NOFRIXION_APIKEY=sk_live_xxx

# 3. Each command line is ONE AQL expression, run against the API:
./nofrixion-cli list account
./nofrixion-cli load 1 account            # {id:1} shorthand
./nofrixion-cli load '{id:1}' account       # explicit match map
./nofrixion-cli update '{name:"x"}' account
./nofrixion-cli list batch

# 4. Override the API base URL for a single call
NOFRIXION_BASE=https://api.example.com ./nofrixion-cli list account

# 5. No arguments -> interactive REPL
./nofrixion-cli
nofrixion> list account
nofrixion> :quit
```

> The rest of this guide follows the [Diátaxis](https://diataxis.fr) framework:
> a hands-on **Tutorial**, task-focused **How-to guides**, a factual
> **Reference**, and background **Explanation**.

## Tutorial: your first query in under a minute

1. **Build the binary.** From this `go-cli/` directory:

   ```sh
   make build          # -> dist/<os>-<arch>/nofrixion-cli
   ```

2. **Set your API key** (read from the environment):

   ```sh
   export NOFRIXION_APIKEY=sk_live_xxx
   ```

3. **Run a query.** Evaluate an AQL expression against the API (or run with no
   arguments to open the REPL):

   ```sh
   ./dist/*/nofrixion-cli list account
   ```

4. **Go interactive.** Run the binary with no arguments to open the REPL, then
   type `:help` for the word and entity lists and `:quit` to leave.

That is the whole loop: *build → set key → evaluate AQL expressions*.

## How-to guides

### List the records of an entity

```sh
./nofrixion-cli list account
```

`list <entity>` returns the first page of records. `<entity>` is a bareword —
it is auto-quoted as an AQL atom, so no quotes are needed.

### Load a single record

```sh
./nofrixion-cli load 1 account          # scalar shorthand for {id:1}
./nofrixion-cli load '{id:1}' account     # explicit match map
```

The query is either a **scalar** (`1`, treated as `{id:1}`) or a **match map**
(`{id:1}`, `{slug:"acme"}`). Quote the map so your shell passes it through intact.

### Update a record

```sh
./nofrixion-cli update '{id:1,name:"new"}' account
```

The match map carries both the selector and the new field values; the updated
record is printed back.

### Authenticate and choose an environment

Configuration is read from the environment — nothing is written to disk:

```sh
export NOFRIXION_APIKEY=sk_live_xxx            # API key
export NOFRIXION_BASE=https://api.example.com  # optional: override the API base URL
./nofrixion-cli list account
```

Both are injectable by a secrets vault, so the key never has to be typed inline.

### Explore interactively with the REPL

Run with no arguments to open a REPL (prompt `nofrixion>`). Each line is
evaluated as its own AQL expression:

```text
$ ./nofrixion-cli
nofrixion> list account
nofrixion> :help
nofrixion> :quit
```

### Cross-compile release binaries

```sh
make build       # native binary for this machine
make build-all   # linux/darwin/windows x amd64/arm64, under dist/<os>-<arch>/
```

### Discover the available entities

`:help` in the REPL prints the full entity list, or see [Entities](#entities)
below — this SDK exposes 49 entities.

## Reference

### Words

The CLI registers these AQL words, each bound to the SDK:

| Word     | Signatures                                    | Returns                        |
|----------|-----------------------------------------------|--------------------------------|
| `list`   | `list <entity>` · `list <query> <entity>`     | First page of records          |
| `load`   | `load <entity>` · `load <query> <entity>`     | A single record                |
| `update` | `update <query> <entity>`                     | Update a record, return it     |

- `<entity>` is a bareword, auto-quoted as an AQL atom (e.g. `account`).
- `<query>` is either a **Map** (`{id:1}`) or a **Scalar** (`1`, treated as
  `{id:1}`). A scalar is always wrapped as `{id:<value>}`.

### Environment variables

| Variable | Purpose |
|----------|---------|
| `NOFRIXION_APIKEY` | API key sent with every request. |
| `NOFRIXION_BASE` | Optional override of the API base URL. |

Unset variables fall back to the SDK's built-in defaults.

### REPL commands

- `:quit` / `:q` / `:exit` — exit the REPL
- `:help` / `:h` / `:?`     — show the word list, entity list and meta commands

### Exit codes

| Code | Meaning |
|------|---------|
| `0` | Success (also the normal REPL exit). |
| `1` | Parse error, word-registration error, or an API/evaluation error. |

### Build targets

| Target | Result |
|--------|--------|
| `make build` | Native binary at `dist/<os>-<arch>/nofrixion-cli`. |
| `make build-all` | linux/darwin/windows x amd64/arm64, each under its own `dist/<os>-<arch>/`. |
| `make clean` | Remove `dist/` and any stray binaries. |

### Entities

The 49 entities this SDK exposes (any is valid as `<entity>`):

account batch beneficiary beneficiary_group card card_customer_token card_payment card_public_key consent currency direct_debit_batch_submit fx_rate i_payment mandate merchant merchant_authorisation_setting merchant_direct_debit_mandate_page merchant_pay_by_bank_setting merchant_payment_request_template merchant_token metadata no_frixion_version open_banking payeeverification payment payment_account payment_account_minimal payment_initiation payment_request payment_request_event payment_request_metric payment_request_minimal payment_request_result payout payout_keyset_page payout_metric payrun report report_result role rule rule_event tag token transaction user user_invite virtual webhook

## Explanation

### Why AQL?

The whole command line is one [AQL](https://github.com/aql-lang/aql) expression,
not a fixed `verb --flag` grammar. That means the same binary works one-shot
(`./nofrixion-cli <expr>`) and interactively (the REPL), and expressions compose the
same way in both. `list` / `load` / `update` are ordinary AQL *words* bound to
the SDK — adding SDK operations is adding words, not re-parsing flags.

### How it is wired

`main.go` builds the SDK client (configured from the environment), creates an
AQL registry, and `words.go` registers `list` / `load` / `update` as native
words that dispatch on the entity atom and call the sibling Go SDK at `../go`.
Results are unwrapped from their `Entity` wrappers to plain data before being
printed.

### Output format

Each result value is printed as its AQL string form (a JSON-like rendering of
the record or list of records). One-shot mode prints to stdout; errors go to
stderr with a non-zero exit code.

## Generated by

sdkgen `go-cli` target. See the target source under `.sdk/src/cmp/go-cli/` in
this repo, or upstream at
`github.com/voxgig/sdkgen/project/.sdk/src/cmp/go-cli/`.
