/*
Package apperrors centralizes all custom errors for the application layer.

This package defines and manages application-specific errors, ensuring a
unified way to handle use case execution failures, application validation issues,
and wrapped infrastructure errors. By isolating these failures here, the system
maintains a clean separation of concerns, allowing input/output adapters and
orchestration layers to interpret and respond to application errors predictably
using Go's standard errors.Is() and errors.As() mechanisms.

RECOMMENDATION: Always use the provided factory functions to instantiate application
errors instead of initializing the error structs directly.

The following factories must be used depending on the failure context:

  - NewInfrastructureError
  - NewEmailInvalidFormatError
  - NewEmptyFieldError
*/
package apperrors
