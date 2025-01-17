package ovh

import "fmt"

// APIError represents an error that can occurred while calling the API.
type APIError struct {
	// Error class
	Class string
	// Error message.
	Message string
	// Error details
	Details map[string]string
	// HTTP code.
	Code int
	// ID of the request
	QueryID string
}

func (err *APIError) Error() string {
	return fmt.Sprintf("Error %d: %q", err.Code, err.Message)
}
