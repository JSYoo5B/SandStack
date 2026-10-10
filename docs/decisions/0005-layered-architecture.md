# ADR 0005: Layered Architecture and Dependency Direction

## Status

Accepted.

## Context

Storage, runtime backends, clocks, fault execution, and OpenStack wire details
are all expected to be replaced or extended. Implementation code is expected
to be rewritten more often than it is reviewed line by line, so the structure
has to be stated as rules rather than left to review.

## Decision

SandStack uses a DDD-inspired layered architecture with this dependency
direction:

```text
cmd -> internal/api -> internal/app -> internal/domain
                    -> internal/store
                    -> internal/provider
                    -> internal/platform
```

Layer responsibilities:

- `internal/api`: HTTP routing, authentication middleware, OpenStack wire
  DTOs, and response mapping.
- `internal/app`: use cases, orchestration, transaction boundaries, and
  dependency interfaces.
- `internal/domain`: resource lifecycle rules and service-independent domain
  models.
- `internal/store`: state persistence implementations.
- `internal/provider`: optional runtime backends.
- `internal/platform`: technical utilities such as configuration, clocks, IDs,
  and logging.

Rules:

- API handlers depend on application services, not on repositories.
- Repository interfaces are narrow, domain-specific, and operation-oriented.
  They do not expose maps, slices, or SQL-shaped APIs.
- Fault evaluation is an application service, not handler-local logic.
- Internal API packages are organized by OpenStack service name.
- Empty layers are not created ahead of need, but no code may violate the
  dependency direction.

## Consequences

- Persistence and runtime backends can be swapped without touching handlers.
- When behavior grows, the missing layer is added instead of folding storage,
  provider, clock, or fault concerns into API handlers.
