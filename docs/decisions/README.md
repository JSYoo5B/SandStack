# Decisions

This directory holds the decisions carried into the current repository root.
Each record states one decision with its context and consequences. A decision
changes through a new record that supersedes the old one.

| ADR | Decision | Status |
|---|---|---|
| [0001](0001-router.md) | HTTP router: Chi | Accepted |
| [0002](0002-emulate-openstack-apis.md) | Emulate OpenStack APIs instead of deploying OpenStack | Accepted |
| [0003](0003-gophercloud-client-contract.md) | Gophercloud as the primary client contract | Accepted |
| [0004](0004-service-url-layout.md) | Single-port service URL layout | Accepted |
| [0005](0005-layered-architecture.md) | Layered architecture and dependency direction | Accepted |
| [0006](0006-runtime-determinism.md) | Injected clock and generators | Accepted |
| [0007](0007-dependency-policy.md) | Dependency policy | Accepted |

## Pending

The implementation scope policy and the verification policy are being
redefined. No implementation code is added until both are recorded here.

## Not Carried Over

The previous root (tag `archive/pre-reset-main`) and its unmerged pull request
stack (tag `archive/pre-reset-pr-stack`) also contained the following. They
were dropped on purpose and should not return without a new record.

- SQLite as a required durable store target. Durable state is undecided.
- The phase ordering that required broad API compatibility before fault
  injection.
- The phased implementation plan and the MVP contract, including fixture
  credentials, Keystone request payloads, the catalog contents, and the request
  ID format. These described OpenStack wire behavior by hand and are revisited
  with the verification policy.
- All implementation code and tests.
