# Implementation Status

This document is a lightweight review aid. It records what the current branch
has implemented without changing the implementation plan.

## Completed Working Slices

### Phase 1: Server Skeleton and Identity

- Chi-based root router.
- SandStack admin status endpoints.
- Keystone v3-compatible identity discovery and password authentication.
- Service catalog entries for identity, compute, image, network, placement, and
  volume.
- Per-commit GitHub Actions validation for Go changes.

### Phase 2: Minimal Resource APIs

- Image create, list, get, and delete.
- Network create, list, get, and delete.
- Subnet create, list, get, and delete.
- Port create, list, get, and delete.
- Compute flavor list and get.
- Compute server create, list, get, and delete.
- Volume create, list, get, and delete.
- Volume type list and get.
- Placement version stub.
- Root-router compatibility test covering image, network, compute, and volume.

### Phase 3: Basic Runtime Behavior

- Injectable clock helper with wall-clock default.
- Injectable ID generator helper with random default.
- App service reset support.
- `POST /_sandstack/reset`.
- Request ID middleware.
- Basic request recording via `GET /_sandstack/requests`.
- Server read transition from `BUILD` to `ACTIVE`.
- Volume read transition from `creating` to `available`.

### Phase 4: First Fault Injection Slice

- Fault rule model and in-memory evaluator.
- `http_error` behavior.
- Operation matching for `compute/server.create`.
- `nth` and `once` triggers.
- Fault admin endpoints to create, list, enable, and disable rules.
- Compatibility test for "third server create returns 503".

## Next Candidates

Phase 5 should stay usage-driven. Good next slices are:

- State dump endpoint if external tests need to inspect all in-memory resources.
- More fault operations after a real test needs them.
- Additional fault behaviors such as delay, quota exceeded, stale read, or stuck
  state after a real compatibility test asks for them.
- Controlled clock endpoints only when tests need time to move without waiting.
