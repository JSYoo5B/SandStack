# SandStack Agent Guidelines

## Language

Keep all durable project artifacts in English. This includes code, comments,
documentation, commit messages, test names, and agent guidance.

## Architecture

Keep Gophercloud compatibility as the strongest external contract. Isolate it
at API compatibility and test boundaries unless using Gophercloud public types
clearly reduces duplicate OpenStack wire structs.

Assume implementations may be replaced. Keep dependency direction explicit and
avoid coupling API handlers to storage, providers, clocks, fault execution, or
OpenStack wire details beyond their adapter role.

Target dependency direction:

```text
cmd -> internal/api -> internal/app -> internal/domain
                    -> internal/store
                    -> internal/provider
                    -> internal/platform
```

Layer responsibilities:

- `internal/api`: HTTP routing, authentication middleware, OpenStack wire DTOs,
  and response mapping.
- `internal/app`: use cases, orchestration, transaction boundaries, and
  dependency interfaces.
- `internal/domain`: resource lifecycle rules and service-independent domain
  models.
- `internal/store`: state persistence implementations.
- `internal/provider`: optional runtime backends such as Noop, Docker, or Lima.
- `internal/platform`: technical utilities such as config, clocks, IDs, and
  logging.

Do not create empty layers just to satisfy the target shape. Add a layer when
the current slice has behavior or dependencies that belong there.

## Implementation Scope

Build the next Gophercloud-compatible working slice first. If a concern does
not help the current slice run, document it as a future consideration instead
of implementing it early.

## Tests

Use Testify for readable assertions. Prefer suite tests when shared setup
improves clarity.

Use `require` for setup failures, call prerequisites, and checks that must pass
before the test can safely continue. Use `assert` for ordinary expected-value
comparisons.

Use whitespace and variable names to make arrange, act, and assert phases easy
to scan. Do not add explicit Given-When-Then comments unless the test would be
hard to follow without them.

## Style

Keep code and comments in English. Avoid long horizontal lines; split function
calls, struct literals, and assertions across multiple lines when readability
improves.

Prefer small files grouped by purpose over large files that mix routing, DTOs,
application logic, and helpers.

## Configuration

Keep simple environment loading for Phase 1. Reconsider Viper when configuration
needs nested files, multiple sources, or validation beyond simple defaults.
