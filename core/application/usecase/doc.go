/*
Package usecase coordinates and implements the core business logic of the application.

In Clean Architecture, this layer embodies the application-specific business rules. It orchestrates
the flow of data to and from the domain, directing input ports to accept data and triggering
output ports (driven interfaces) to interact with external systems without knowing their
underlying technologies.

# Architectural Integrity & Isolation

To preserve a decoupled architecture, components within this package must never import or depend
on concrete infrastructure details (such as HTTP routing, database drivers, or specific cloud providers).
They rely strictly on the abstract contracts declared in the ports layers.

# Error Handling Principle

Use cases act as a safe haven from low-level technical complexities. They only interpret, process,
and return application-centric errors (such as apperrors.InvalidFieldError or apperrors.InfrastructureError),
ensuring the core business flow remains predictable, highly testable, and isolated.
*/
package usecase
