# Open Questions

This file tracks decisions that are not needed for the next working slice.

## Gophercloud Version

Which `github.com/gophercloud/gophercloud/v2` version should be pinned after local verification?

Needed before: Phase 1 implementation.

## Durable State

Should SandStack eventually support SQLite, MySQL, or PostgreSQL state backends?

Current stance: future consideration. Do not introduce GORM or database-shaped
domain models for the MVP.

## Advanced Faults

Which fault behaviors are needed after `http_error`, `nth`, and `once`?

Current stance: let real compatibility and integration tests choose the next behaviors.

## Scenario Format

Should scenario files have an explicit schema version?

Current stance: defer until the first scenario loader exists. Add a version
field if scenarios become durable user-facing test assets.

## Admin API Scope

Which admin APIs are necessary for the MVP?

Current stance: implement only the admin endpoints needed by the current phase.
Request logs, state dumps, scenario management, and clock control can wait
until their consumers exist.
