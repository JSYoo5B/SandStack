# ADR 0006: Injected Clock and Generators

## Status

Accepted.

## Context

Fault injection and timing-sensitive tests need reproducible behavior, while
normal use should feel like a regular service with random IDs and wall-clock
time.

## Decision

- Code that needs time uses a clock interface. The default implementation is
  the wall clock. Tests may inject a controllable clock.
- Values that real OpenStack generates, such as UUIDs and timestamps, come
  from injected generators. MAC addresses, IP allocation order, fault
  probability, and event ordering follow the same rule when they appear.
- SandStack is random by default and reproducible when a seed is configured.
  Determinism is a test control, not the only runtime mode.

## Consequences

- Domain and application code never read the wall clock or a global random
  source directly.
- Clock control endpoints can be added later without changing domain code.
