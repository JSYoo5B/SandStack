# ADR 0002: Emulate OpenStack APIs Instead of Deploying OpenStack

## Status

Accepted.

## Context

Services built on OpenStack need an OpenStack dependency in automated tests.
Running real OpenStack as a Testcontainers-style dependency is impractical:

- nova-compute needs libvirt and hardware virtualization. Neutron agents need
  Open vSwitch, network namespaces, and iptables. Both need privileged
  containers, and developer machines often cannot provide KVM acceleration.
- An API-only deployment with fake backends still needs a database, a message
  queue, a cache, and many service containers. It takes minutes to start and
  gigabytes of memory, which does not fit a per-suite test dependency.
- A real deployment cannot reproduce state faults on demand. Servers stuck in
  a transitional state, stale reads, and inconsistencies between services
  happen only by chance.

Transport faults such as timeouts, connection resets, and injected 5xx
responses do not need an emulator. A fault proxy in front of any endpoint can
produce them.

## Decision

SandStack emulates OpenStack HTTP APIs over its own state. It is not an
OpenStack distribution and not a DevStack replacement.

SandStack owns resource state so that tests can reset it, inspect it, and
inject state faults deterministically. It does not provide real
virtualization, virtual networking, or distributed storage semantics. The
default runtime backend only keeps state. Optional backends such as Docker or
Lima may back selected resources later.

## Consequences

- SandStack starts in seconds without a hypervisor, a container runtime, or an
  external database.
- Emulated behavior can drift from real OpenStack. Fidelity has to be checked
  against sources that do not come from the SandStack implementation itself.
  How that is verified is decided in a separate record.
- Every emulated API adds surface that has to be verified and maintained.
