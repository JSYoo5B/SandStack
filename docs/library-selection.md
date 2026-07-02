# Library Selection

## Required

### Gophercloud

Use `github.com/gophercloud/gophercloud/v2`.

Gophercloud is the primary compatibility target. SandStack should use it in
tests to verify authentication, service catalog discovery, request shapes,
response extraction, and error handling.

The exact version should be pinned during implementation after checking
available module versions locally.

## Selected

### Chi

Use `github.com/go-chi/chi/v5` for HTTP routing.

Chi gives SandStack readable route grouping and middleware composition while
keeping handlers compatible with `net/http`.

## Deferred Until Needed

### YAML

Use `gopkg.in/yaml.v3` for scenario parsing when scenario loading becomes part
of the implementation.

The scenario format is YAML-first. Strict decoding should be used where
possible so invalid scenario files fail early.

### Test Assertions

Use `github.com/stretchr/testify` for readable assertions and suites when tests
benefit from explicit setup and Given-When-Then flow.

Use `github.com/google/go-cmp/cmp` for deterministic state comparisons.

### Configuration

Consider Viper when configuration grows beyond simple environment defaults.
Do not introduce it while the server only needs a small set of environment
variables.

### Durable Persistence

GORM or database drivers should be considered only when a durable state backend
is actually implemented.

## Avoid Initially

- Full web frameworks.
- ORM or database libraries.
- Background job frameworks.
- Generic DDD or CQRS frameworks.

These can be reconsidered when the implementation has real pressure that justifies them.
