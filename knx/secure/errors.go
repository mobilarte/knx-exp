package secure

import "fmt"

// DataSecureError represents an error during Data Secure operations
type DataSecureError struct {
	Message string
	Cause   error
}

// NewDataSecureError creates a new DataSecureError
func NewDataSecureError(message string) *DataSecureError {
	return &DataSecureError{Message: message}
}

// NewDataSecureErrorWithCause creates a new DataSecureError with a cause
func NewDataSecureErrorWithCause(message string, cause error) *DataSecureError {
	return &DataSecureError{Message: message, Cause: cause}
}

func (e *DataSecureError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("data secure error: %s: %v", e.Message, e.Cause)
	}
	return fmt.Sprintf("data secure error: %s", e.Message)
}
