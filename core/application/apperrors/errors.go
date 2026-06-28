package apperrors

import (
	"errors"
)

var (
	ErrTemporaryFailure = errors.New("temporary failure")
	ErrPermanentFailure = errors.New("permanent failure")
)

// ================ InfrastructureError ======================

// InfrastructureError represents a technical or operational failure that occurs
// specifically within the output adapters layer (Output/Driven Adapters),
// such as database repositories, external API clients, or file systems.
//
// This structure is used to capture unexpected infrastructure errors
// (e.g., timeouts, connection drops, disk failures) and "shield" the application's
// Core/Use Cases, preventing technology-specific implementation details
// (such as SQL driver or HTTP errors) from leaking into the business logic layer.
type InfrastructureError struct {
	// Message provides a clear, high-level context of what the application
	// was trying to achieve when the failure occurred (e.g.,"failed to save user to database").
	Message string

	// OriginalCause stores the raw, unexpected technical error that triggered the issue.
	// This preserves the original underlying error for debugging and logging purposes.
	OriginalCause error
}

// Error formats the error message for display. If an original cause is present,
// it concatenates the high-level application context with the underlying technical error.
func (i *InfrastructureError) Error() string {
	if i.OriginalCause != nil {
		return i.Message + ": " + i.OriginalCause.Error()
	}
	return i.Message
}

// Unwrap exposes the underlying error that caused the infrastructure failure.
// This allows higher layers to use standard Go functions, such as errors.Is()
// and errors.As(), to inspect the root cause if necessary.
func (i *InfrastructureError) Unwrap() error {
	return i.OriginalCause
}

// ================ InvalidFieldError ======================

// InvalidFieldError marks errors related to invalid field.
type InvalidFieldError struct {
	Message string

	// FieldName stores the name of the invalid field.
	//
	// This is particularly useful for building structured HTTP responses,
	// allowing API clients to identify exactly which field caused the error
	// (e.g., mapping fields in a 400 Bad Request JSON payload).
	FieldName string
}

func (i *InvalidFieldError) Error() string {
	return i.Message
}
