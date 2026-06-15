# SandStack Architecture

## Product Intent

SandStack is an OpenStack-compatible API sandbox for integration tests. It should be easy to start as a temporary dependency, authenticate with Keystone v3-compatible flows, create selected OpenStack resources, inspect internal state, reset state between tests, and inject programmable faults that are difficult to reproduce against a real OpenStack deployment.

SandStack is not a real OpenStack distribution. It optimizes for client compatibility, repeatability, observability, and controlled failure behavior over virtualization, networking, or storage fidelity.

## Architectural Style

SandStack uses DDD-inspired layered architecture without making the MVP artificially abstract. Most implementation details are expected to change as compatibility behavior, fault injection, and provider boundaries become clearer.

Gophercloud compatibility is the only intentionally strong external constraint. Even there, Gophercloud should be isolated at the API compatibility and test boundary rather than leaking throughout application or domain code.

The architecture should optimize for reversible decisions. When a concern is not required for the next working slice, document it as a future consideration instead of building for it immediately.

The strongest internal boundaries are:

- OpenStack HTTP compatibility adapters.
- Application use cases.
- Domain state and transitions.
- Infrastructure concerns such as in-memory storage, event scheduling, request recording, and providers.

Recommended dependency direction:

```text
cmd -> internal/api -> internal/app -> internal/domain
                    -> internal/store
                    -> internal/provider
                    -> internal/platform
```

The API layer knows OpenStack JSON shapes and Gophercloud compatibility requirements. The application layer owns use-case orchestration, transaction boundaries, and dependency interfaces. The domain layer owns resource lifecycle rules, cross-service references, and state transitions. Store packages provide state persistence implementations. Provider packages isolate optional runtime backends. Platform packages provide technical utilities such as configuration, clocks, IDs, and logging.

Early phases may omit empty layers, but they should not violate the target dependency direction. When behavior grows, add the missing layer instead of folding storage, provider, clock, or fault concerns into API handlers.

## Package Layout

```text
cmd/
  sandstack-server/
internal/
  api/
    admin/
    compute/
    identity/
    image/
    network/
    placement/
    volume/
    middleware/
    openstack/
  app/
    compute/
    identity/
    image/
    network/
    volume/
    scenario/
    faults/
  domain/
    catalog/
    clock/
    compute/
    event/
    faults/
    identity/
    image/
    network/
    requestlog/
    state/
    volume/
  platform/
    config/
    httpserver/
    idgen/
    log/
  provider/
    noop/
  store/
    memory/
scenarios/
test/
  compatibility/
  integration/
docs/
```

Internal API packages are organized by OpenStack service name. External service paths should stay explicit and service-prefixed where possible. For example, `internal/api/identity` should be mounted at `/identity/v3`.

## Request Flow

This is the target request flow. Early phases may implement only the steps required by the current working slice.

```text
HTTP request
-> route match
-> request context creation
-> OpenStack or admin authentication
-> request recorder start
-> pre-handler fault evaluation
-> API adapter parses compatibility DTO
-> application use case
-> domain transition and state mutation
-> event enqueue
-> post-handler fault evaluation
-> API adapter renders compatibility DTO
-> request recorder finish
-> response
```

Fault rules receive normalized metadata such as service, operation, project ID, user ID, resource type, resource ID, request body, and request headers.

## Compatibility Boundary

OpenStack-facing request and response compatibility must be validated against Gophercloud. SandStack treats Gophercloud's accepted request shapes, extracted response structs, status-code expectations, endpoint discovery behavior, and error handling as primary compatibility contracts.

Gophercloud is also a Go representation of many OpenStack component request and response shapes. SandStack should still keep its own API DTOs because OpenStack responses may vary by service, API microversion, extension fields, backend behavior, or client extraction expectations. Gophercloud defines the compatibility target, but it should not become the domain model.

API DTOs are separate from domain models:

- `internal/api/*` DTOs use OpenStack-compatible JSON names.
- `internal/domain/*` models use clean Go names and lifecycle-focused structures.
- Mapping functions are explicit and tested with Gophercloud clients.

This protects domain design when OpenStack response shapes require unusual fields such as `OS-EXT-STS:vm_state`, `tenant_id`, or service-specific wrappers.

## State Store

The MVP store is an in-memory repository with a single state aggregate guarded by `sync.RWMutex`. This keeps the first implementation easy to run, reset, and reason about.

Durable state backends are a future consideration. Some external integration tests may eventually benefit from preserving a prepared SandStack state across process restarts or sharing a controlled state between test processes. Future backends may include SQLite, MySQL, or PostgreSQL, and a library such as GORM can be reconsidered at that time.

For now, persistence should be treated as a store boundary concern, not as a reason to introduce database abstractions early.

The store exposes operation-level methods instead of leaking maps directly.

Mutable resources should keep only the metadata needed by the current implementation. Version history and soft-delete metadata can be added when stale reads, deleted-resource inspection, or durable storage require them.

The initial store interface should be narrow and domain-specific. It should not design for SQL too early, but it should avoid APIs that make future durable backends impossible.

## Randomness and Reproducibility

SandStack should be random by default where real OpenStack would naturally generate values, but reproducible when a seed is configured. Determinism is a test control, not the only runtime mode.

A runtime generation component owns:

- UUID generation.
- MAC generation.
- IP allocation order.
- Fault probability decisions.
- Event jitter and ordering.

Domain code should use injected generators for values that tests may need to reproduce. The first implementation only needs to cover generated IDs and timestamps; MAC addresses, IP allocation, probabilistic fault decisions, and event ordering can be added as those features appear.

## Event and Clock Model

Use a clock interface wherever application or domain logic needs time. The default implementation uses wall-clock time. Tests and admin-controlled runtime modes may inject a controllable clock.

Async-looking state transitions should be represented as events once the basic synchronous behavior is working:

- `server.activate`
- `server.delete.complete`
- `image.upload.complete`
- `volume.available`
- `volume.attach.complete`
- `port.activate`

Wall-clock time is the default mental model. SandStack should behave like a normal service unless tests explicitly choose a controlled clock implementation.

Controlled clock mode exists to accelerate tests and reproduce timing-sensitive behavior. Admin endpoints for advancing time can be added after the core clock interface and event behavior are stable.

Auto mode processes due events according to the active clock. Manual mode can be added when explicit clock control becomes necessary.

The first implementation can keep transitions simple and synchronous if that gets Gophercloud compatibility working faster. A queue or background worker can be introduced after real tests need observable asynchronous behavior.

## Fault Engine

Fault evaluation is a separate application service, not handler-local logic.

The architecture should keep explicit extension points for fault injection and avoid prematurely deciding details that can be expressed as fault behaviors later. If a behavior can be modeled as a fault rule, prefer keeping the normal implementation simple until real compatibility tests require more detail.

The first useful fault behavior is an injected HTTP error. More advanced behaviors such as stale reads, stuck states, inconsistent states, event reordering, and partial failures should be added after the normal resource APIs exist.

## Provider Layer

Only NoopProvider is implemented for MVP.

Provider interfaces should return provider-level outcomes and errors that application services map into SandStack state transitions or fault-compatible errors. Docker and Lima packages should not exist until they are implemented or explicitly stubbed behind build tags.
