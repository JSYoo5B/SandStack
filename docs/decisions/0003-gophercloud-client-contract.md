# ADR 0003: Gophercloud as the Primary Client Contract

## Status

Accepted.

## Context

SandStack needs one concrete client to define what "compatible" means.
`github.com/gophercloud/gophercloud/v2` is that client. Gophercloud also
encodes many OpenStack request and response shapes in Go, which makes it
tempting to reuse its types throughout the code.

## Decision

Gophercloud compatibility is the strongest external client contract. A
Gophercloud client must be able to authenticate, discover the service catalog,
send requests, extract responses, and handle errors against SandStack.

Gophercloud stays at the API compatibility and test boundary:

- API packages keep their own OpenStack wire DTOs with OpenStack JSON names.
- Domain models use plain Go names and lifecycle-focused structures.
- Mapping between DTOs and domain models is explicit.
- Gophercloud public types may be reused only where they clearly remove a
  duplicate wire struct.

## Consequences

- Wire quirks such as `OS-EXT-STS:vm_state`, `tenant_id`, and
  service-specific wrappers stay in the API layer.
- Gophercloud decodes responses leniently. Missing fields become zero values,
  unknown fields are ignored, and state transitions and error semantics are
  not checked. Passing Gophercloud tests shows client compatibility, not
  behavioral fidelity to OpenStack.
