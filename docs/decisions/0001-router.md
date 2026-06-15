# ADR 0001: HTTP Router

## Status

Accepted for the initial implementation.

## Decision

Use `github.com/go-chi/chi/v5` as the HTTP router.

## Context

SandStack needs many OpenStack-compatible route groups with service-specific middleware, shared request recording, fault hooks, authentication rules, custom 404 and 405 behavior, and readable route composition.

Raw `net/http` is viable, but route grouping and sub-router composition would become repetitive as Keystone, Nova, Neutron, Glance, Cinder, Placement, and Admin APIs grow.

Gin is familiar and productive, but it is a web framework with its own context and binding conventions. SandStack benefits from keeping handlers close to `net/http` because the API layer is an adapter over application services, not the center of the design.

Echo has flexible routing and middleware, but it is also a broader framework with a framework-specific context.

## Rationale

Chi is the best fit because:

- It is compatible with standard `net/http` handlers and middleware.
- It supports route groups, sub-routers, method routing, custom not-found handlers, and custom method-not-allowed handlers.
- It keeps the API layer thin and does not force framework-specific context through application code.
- Its design focuses on project structure, maintainability, standard HTTP handlers, and modular APIs.
- It has no external dependencies beyond the standard library.

Performance is not the primary driver for SandStack, but Chi is still lightweight enough for this use case.

## Consequences

Handlers should accept standard `http.ResponseWriter` and `*http.Request`.

API packages may use `chi.URLParam` at the adapter boundary only. Application and domain packages must not depend on Chi.

Middleware must be written as standard `func(http.Handler) http.Handler` middleware.

Route registration should stay service-local:

```text
internal/api/identity
internal/api/compute
internal/api/network
internal/api/image
internal/api/volume
internal/api/placement
internal/api/admin
```

The server package composes those service routers into the final root router.
