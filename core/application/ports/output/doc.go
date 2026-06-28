/*
Package outputport defines the driven ports (outbound interfaces) of the application core.

In accordance with Hexagonal Architecture (Ports & Adapters) and Clean Architecture principles,
this package contains strictly abstract contracts (interfaces) that declare what the application
core requires from the outside world to fulfill its business use cases.

Output ports act as an isolation barrier between pure business logic and volatile infrastructure
technologies. They ensure that the core remains agnostic

# Error Handling Policy

Implementations of these ports (adapters) must never leak technology-specific errors
(such as driver connection drops or raw HTTP status codes) into the core.

Any operational failure encountered by an adapter must be mapped and wrapped into
standardized application errors (e.g., apperrors.InfrastructureError) before being
propagated back through these interfaces.
*/
package outputport
