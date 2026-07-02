# SandStack MVP Contract

This document captures the implementation facts that should be available without
reopening the original product specification.

## Product Summary

SandStack is a lightweight OpenStack-compatible API sandbox.

It is intended to be started as an ephemeral test dependency for services that
integrate with OpenStack APIs. It should not require a real OpenStack
deployment, Docker socket access, Lima, or a hypervisor for the MVP.

Fault injection is an important long-term capability, but it is not part of the
current MVP contract until its matching semantics, rule lifecycle, and
OpenStack-compatible error behavior are designed.

Before fault injection becomes a main implementation stream, SandStack should
first broaden normal OpenStack API compatibility across the major services used
by Gophercloud clients: Keystone, Nova, Neutron, Glance, Cinder, and Placement.
Fault behavior should be layered onto a useful compatibility surface, not used
as a substitute for missing normal APIs.

## Strong Constraint

Gophercloud compatibility is the primary implementation contract.

The server should accept request shapes and return response shapes that
Gophercloud can use for authentication, service discovery, and the selected MVP
OpenStack resource APIs.

## External Service URL Layout

SandStack uses one HTTP port with service path prefixes:

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

`/identity/v3` is the canonical SandStack Identity endpoint. Gophercloud should
receive this value as the auth URL.

Service catalog URLs should be based on explicit configuration when provided.
Otherwise, they may be derived from the request host.

Runtime defaults such as port, credentials, project, region, and admin token are
defined by configuration during implementation. They do not need to be treated
as architecture decisions.

Phase 1 may use a simple well-known compatibility fixture:

```text
username=admin
password=password
project=demo
role=admin
region=RegionOne
```

## Public Endpoints

These endpoints must be available without OpenStack authentication:

```text
GET  /identity/v3
GET  /_sandstack/health
GET  /_sandstack/ready
POST /identity/v3/auth/tokens
```

Additional unauthenticated version/root endpoints may be added for discovered
services as needed.

## Authentication

OpenStack-compatible service APIs require:

```text
X-Auth-Token: {token}
```

Admin APIs require:

```text
X-SandStack-Admin-Token: {admin_token}
```

The full Admin API can be implemented later. Health and readiness should stay unauthenticated.

## Keystone Phase 1 Contract

`GET /identity/v3` should return a minimal Keystone v3-compatible version
document. It only needs enough fields for client discovery during Phase 1:

```text
version.id
version.status
version.links
```

`POST /identity/v3/auth/tokens` must accept password authentication for the
default user and project.

Minimum accepted request shape:

```json
{
  "auth": {
    "identity": {
      "methods": ["password"],
      "password": {
        "user": {
          "name": "admin",
          "domain": {"id": "default"},
          "password": "password"
        }
      }
    },
    "scope": {
      "project": {
        "name": "demo",
        "domain": {"id": "default"}
      }
    }
  }
}
```

Successful responses must include:

```text
X-Subject-Token
```

The response body must include a token with:

```text
user
project
roles
catalog
expires_at
```

The service catalog must include:

```text
identity
compute
network
image
volumev3
placement
```

Each service should include `public`, `internal`, and `admin` endpoint
interfaces. The MVP may point all interfaces to the same URL.

Phase 1 catalog URLs should follow this mapping:

```text
identity:  /identity/v3
compute:   /compute/v2.1/{project_id}
network:   /network/v2.0
image:     /image/v2
volumev3:  /volume/v3/{project_id}
placement: /placement
```

## Minimal Resource API Contract

The first resource API slice should cover create/read/list/delete behavior only
where needed by Gophercloud compatibility tests.

Initial services:

```text
Image:   images and placeholder upload metadata
Network: networks, subnets, ports
Compute: flavors, servers
Volume:  volumes, volume types
Placement: JSON stub responses
```

Avoid advanced OpenStack behavior until a compatibility test requires it.

## State and Repository Contract

The MVP implementation should use repository boundaries for mutable OpenStack
resources. Application services should not expose or depend on raw maps, slices,
or SQL details.

The first repository implementation may be in-memory. This is an implementation
choice, not the long-term state model.

SandStack should provide SQLite as the minimum durable repository target after
the repository interfaces are established. SQLite should support both `:memory:`
and file-backed modes when implemented.

Longer term, MySQL and PostgreSQL may be added as repository implementations for
workflows that need to start SandStack from a manually prepared or persistent
state.

State design should preserve:

```text
reset support
resource ownership by project and user
cross-resource references
seeded startup from a specified state
compatibility with generated IDs and timestamps
```

## Request IDs

OpenStack-compatible API responses should include:

```text
X-Openstack-Request-Id: req-{id}
```

The exact ID format only needs to be stable enough for tests at first.

## Phase 1 Acceptance

Phase 1 is complete when:

- The server starts on the configured port.
- `GET /_sandstack/health` returns 200.
- `GET /_sandstack/ready` returns 200 after initialization.
- Gophercloud can authenticate against `/identity/v3`.
- Gophercloud can discover the service catalog.
- No real OpenStack, Docker, Lima, or persistent database is required.
