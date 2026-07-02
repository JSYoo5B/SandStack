# SandStack Implementation Plan

## Baseline Decisions

- Go module path: `github.com/JSYoo5B/SandStack`
- Router: `github.com/go-chi/chi/v5`
- Public documentation language: English
- Primary OpenStack compatibility target: Gophercloud
- State strategy: repository-first, with in-memory as the first implementation
- Required durable store target: SQLite
- MVP contract source: `docs/mvp-contract.md`

The exact Gophercloud version should be pinned after running:

```text
go list -m -versions github.com/gophercloud/gophercloud/v2
```

## Guiding Rule

Build the next Gophercloud-compatible working slice first. If a concern does
not help the next slice run, keep it as a future consideration rather than
implementing it early.

## Phase 0: Compatibility Backlog

Goal: make API expansion reviewable before trying to implement "all OpenStack
APIs".

- Add a Gophercloud-oriented API coverage matrix.
- Track OpenStack service package, Gophercloud package, operation, HTTP method
  and path, required request and response fields, support status, behavior gaps,
  and priority.
- Add a small conformance test harness for selected Gophercloud workflows.

## Phase 1: Minimal Server and Identity

Goal: a Gophercloud client can authenticate and discover service endpoints.

- Create `go.mod`.
- Add `cmd/sandstack-server`.
- Add configuration defaults.
- Add Chi root router.
- Add `/_sandstack/health` and `/_sandstack/ready`.
- Add Identity router mounted at `/identity/v3`.
- Add `POST /identity/v3/auth/tokens`.
- Return a minimal service catalog.
- Add Gophercloud auth compatibility test.

## Phase 2: Minimal Resource APIs

Goal: a Gophercloud client can create and read the core resources needed by
common workflows.

- Add minimal Image API.
- Add minimal Network API for networks, subnets, and ports.
- Add minimal Compute API for flavors and servers.
- Add minimal Volume API for volumes and volume types.
- Add Placement stub.
- Add compatibility tests for one happy path across image, network, server, and volume.

## Phase 3: Repository Boundaries

Goal: application services stop owning storage details directly.

- Define narrow resource-level repository interfaces from current application
  behavior.
- Move current in-memory maps and slices behind in-memory repository
  implementations.
- Keep API handlers dependent on application services, not repositories.
- Keep repository interfaces domain-specific and operation-oriented.
- Start with one small resource, such as images, before repeating the pattern
  across network, compute, and volume.

## Phase 4: SQLite and Seeded State

Goal: SandStack can run with a durable local state implementation and can start
from a specified state.

- Add SQLite repositories after in-memory repositories prove the interfaces.
- Support both SQLite `:memory:` and file-backed databases.
- Add schema initialization in Go code first.
- Add migration tooling only when schema changes become hard to manage.
- Add seed/import support for starting SandStack with a prepared state.
- Keep MySQL and PostgreSQL as later repository implementations for external
  RDB workflows and manually prepared state experiments.

## Phase 5: Basic Runtime Behavior

Goal: resources behave enough like OpenStack for service-level tests.

- Add injectable clock interface with wall-clock default.
- Add generated IDs and timestamps through injected runtime helpers.
- Add simple state transitions such as server `BUILD -> ACTIVE` and volume
  `creating -> available`.
- Add reset support for test reuse.
- Add request ID headers and basic request recording.

## Phase 6: Complete Major OpenStack API Compatibility First

Goal: fill the major Gophercloud-facing OpenStack API surfaces before returning
to fault injection design.

This phase should prioritize broad normal-behavior compatibility for Keystone,
Nova, Neutron, Glance, Cinder, and Placement. The intent is that SandStack can
serve as a useful OpenStack-compatible test dependency even without fault
injection. Extension APIs and rarely used operations can still be prioritized
through the compatibility matrix, but fault injection should not become the next
main workstream while major component API gaps remain.

- Add Identity read APIs for projects, users, roles, and service catalog
  inspection.
- Add Compute server actions, metadata, addresses, keypairs, and security
  groups.
- Add Network security groups, routers, floating IPs, and router interfaces.
- Add Volume snapshots, attachments, volume actions, and type extra specs.
- Add Image upload/download, update, member APIs, and lifecycle actions.
- Add Placement resource providers and inventories only when a client needs
  them.

## Phase 7: Lightweight Provider Backends

Goal: optionally back selected resources with local Docker or Lima behavior.

- Define provider interfaces around a small workflow, not around all OpenStack
  resources.
- Keep Noop provider as the default.
- Start with image-to-container server lifecycle.
- Add Docker before Lima unless macOS VM isolation is the first real need.
- Add Lima after server lifecycle and image mapping are proven.

## Phase 8: Fault Injection Revisit

Goal: design and implement fault injection only after normal API and state
behavior are stable enough to define meaningful failure semantics.

This phase should start only after the major OpenStack component APIs are broad
enough for regular client integration tests. Fault behavior should be designed
against real compatibility gaps and observed workflows, not ahead of the normal
API surface.

- Document operation identity and matching rules.
- Decide whether matching happens before request decoding, after request
  decoding, or after application execution.
- Decide rule scope: global, project, user, resource, or a combination.
- Decide rule lifecycle, persistence, and seed interaction.
- Define service-specific OpenStack-compatible error bodies.
- Implement one operation and one behavior first after the design is accepted.

## Continuous Expansion

Goal: grow from real usage instead of speculative design.

Candidate work:

- More fault behaviors: delay, conflict, quota exceeded, stale read, stuck
  state, inconsistent state.
- Event queue and controlled clock endpoints.
- Request log API and state dump API.
- Scenario YAML loader.
- Dockerfile and Testcontainers examples.
- Additional compatibility profiles beyond Gophercloud.

## Acceptance Strategy

Each phase should leave the repository runnable and testable. Compatibility
tests using Gophercloud should be added as soon as the relevant OpenStack API
surface exists.

Commits are intentionally not part of this plan. Changes should be reviewed
before the user commits them.
