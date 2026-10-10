# ADR 0004: Single-Port Service URL Layout

## Status

Accepted.

## Context

Real OpenStack deployments expose each service on its own endpoint. Tests need
one address to start, wait on, and point clients at.

## Decision

SandStack serves every API on one HTTP port, separated by service path
prefixes:

```text
/identity/v3
/compute/v2.1
/compute/v2.1/{project_id}
/network/v2.0
/image/v2
/volume/v3/{project_id}
/placement
/_sandstack
```

- `/identity/v3` is the canonical auth URL given to clients.
- `/_sandstack` holds SandStack endpoints that are not part of any OpenStack
  API, such as health, readiness, and test controls.
- OpenStack APIs authenticate with `X-Auth-Token`. SandStack endpoints
  authenticate with `X-SandStack-Admin-Token`. Health and readiness stay
  unauthenticated.
- Service catalog URLs come from explicit configuration when provided.
  Otherwise they are derived from the request host.

## Consequences

- One container port and one readiness check are enough for a test harness.
- Clients reach services other than Identity through catalog discovery, so
  those prefixes can change without breaking clients that do not hard-code
  service URLs.
