# ADR 0007: Dependency Policy

## Status

Accepted.

## Context

SandStack keeps its API layer as a thin adapter and expects implementations to
be replaced. Large frameworks would pull their own conventions into every
layer.

## Decision

Selected:

- `github.com/gophercloud/gophercloud/v2` for compatibility tests
  (ADR 0003).
- `github.com/go-chi/chi/v5` for HTTP routing (ADR 0001).
- `github.com/stretchr/testify` for assertions and suites.
- `github.com/google/go-cmp/cmp` for deterministic state comparisons.

Selected when the need appears:

- `gopkg.in/yaml.v3` with strict decoding for scenario files. Scenario files
  are YAML-first.
- Viper only when configuration needs nested files, multiple sources, or
  validation beyond simple environment defaults.
- Database drivers only when a durable state backend is actually implemented.

Avoided:

- Full web frameworks.
- ORM libraries.
- Background job frameworks.
- Generic DDD or CQRS frameworks.

## Consequences

- Avoided dependencies can be reconsidered when the implementation has real
  pressure that justifies them, through a new record.
