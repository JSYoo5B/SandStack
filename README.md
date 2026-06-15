# SandStack

SandStack is a lightweight OpenStack-compatible test sandbox with programmable fault injection.

It is designed for automated tests of services that integrate with OpenStack APIs. SandStack can be started as an ephemeral test dependency, similar to Testcontainers-managed services. It provides Keystone-compatible authentication, selected Nova/Neutron/Glance/Cinder-compatible APIs, deterministic state reset, request tracing, and programmable faults for backend failures, stale reads, timing issues, and cross-service inconsistencies.

SandStack is not a full OpenStack distribution and is not a DevStack replacement. Optional local backends may use Docker or Lima to provide simplified resources, but SandStack does not guarantee real OpenStack virtualization, virtual networking, or distributed storage semantics.
