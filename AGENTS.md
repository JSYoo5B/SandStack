# SandStack Agent Guidelines

## Decisions

Accepted decisions live in `docs/decisions/`. Read its index before changing
code or documents. Do not reintroduce anything listed there under
"Not Carried Over" without a new decision record.

## Pending Policy

The implementation scope policy and the verification policy are being
redefined. Do not add implementation code until both are recorded in
`docs/decisions/`.

## Language

Keep all durable project artifacts in English. This includes code, comments,
documentation, commit messages, test names, and agent guidance.

## Architecture

Assume implementations may be replaced. Follow the dependency direction and
layer responsibilities in ADR 0005, and keep Gophercloud at the API
compatibility and test boundary as described in ADR 0003.

## Tests

Use Testify for readable assertions. Prefer suite tests when shared setup
improves clarity.

Use `require` for setup failures, call prerequisites, and checks that must pass
before the test can safely continue. Use `assert` for ordinary expected-value
comparisons.

Use whitespace and variable names to make arrange, act, and assert phases easy
to scan. Do not add explicit Given-When-Then comments unless the test would be
hard to follow without them.

## Style

Keep code and comments in English. Avoid long horizontal lines; split function
calls, struct literals, and assertions across multiple lines when readability
improves.

Prefer small files grouped by purpose over large files that mix routing, DTOs,
application logic, and helpers.

## Configuration

Keep simple environment loading. Reconsider Viper only under the conditions in
ADR 0007.
