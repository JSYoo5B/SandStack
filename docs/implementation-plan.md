# SandStack Implementation Plan

## Baseline Decisions

- Go module path: `github.com/JSYoo5B/SandStack`
- Router: `github.com/go-chi/chi/v5`
- Public documentation language: English
- Primary OpenStack compatibility target: Gophercloud
- MVP state backend: in-memory
- MVP contract source: `docs/mvp-contract.md`

The exact Gophercloud version should be pinned after running:

```text
go list -m -versions github.com/gophercloud/gophercloud/v2
```

## Guiding Rule

Build the next Gophercloud-compatible working slice first. If a concern does not help the next slice run, keep it as a future consideration rather than implementing it early.

## Phase 1: Minimal Server and Identity

Goal: a Gophercloud client can authenticate and discover service endpoints.

- Create `go.mod`.
- Add `cmd/sandstack-server`.
- Add configuration defaults.
- Add Chi root router.
- Add `/`, `/_sandstack/health`, and `/_sandstack/ready`.
- Add Identity router mounted at `/identity/v3`.
- Add `POST /identity/v3/auth/tokens`.
- Return a minimal service catalog.
- Add Gophercloud auth compatibility test.

## Phase 2: Minimal Resource APIs

Goal: a Gophercloud client can create and read the core resources needed by common workflows.

- Add in-memory state.
- Add minimal Image API.
- Add minimal Network API for networks, subnets, and ports.
- Add minimal Compute API for flavors and servers.
- Add minimal Volume API for volumes and volume types.
- Add Placement stub.
- Add compatibility tests for one happy path across image, network, server, and volume.

## Phase 3: Basic Runtime Behavior

Goal: resources behave enough like OpenStack for service-level tests.

- Add injectable clock interface with wall-clock default.
- Add generated IDs and timestamps through injected runtime helpers.
- Add simple state transitions such as server `BUILD -> ACTIVE` and volume `creating -> available`.
- Add reset support for test reuse.
- Add request ID headers and basic request recording.

## Phase 4: First Fault Injection Slice

Goal: tests can force a predictable OpenStack API failure.

- Add fault rule model.
- Add pre-handler fault hook.
- Add operation matching.
- Add `http_error`.
- Add `nth` and `once`.
- Add minimal fault admin endpoints only as needed to create, list, enable, and disable rules.
- Add compatibility test for "third server create returns 503".

## Phase 5: Expand Only Where Tests Demand

Goal: grow from real usage instead of speculative design.

Candidate work:

- More fault behaviors: delay, conflict, quota exceeded, stale read, stuck state, inconsistent state.
- Event queue and controlled clock endpoints.
- Request log API and state dump API.
- Scenario YAML loader.
- Durable state backend investigation.
- Dockerfile and Testcontainers examples.
- Additional compatibility profiles beyond Gophercloud.

## Acceptance Strategy

Each phase should leave the repository runnable and testable. Compatibility tests using Gophercloud should be added as soon as the relevant OpenStack API surface exists.

Commits are intentionally not part of this plan. Changes should be reviewed before the user commits them.
