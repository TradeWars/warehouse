package types

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// Status is a custom status object returned by all endpoints. This is due
// to the fact that HTTP status codes do not match the use-case of this service
// so all endpoints will return either 200 or 500 with this object wrapping any
// response object with a message and error state.
type Status struct {
	Result  interface{} `json:"result"`
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
}

// NewStatus creates and returns a Status
// always use this in order to ensure all fields are filled
// message may be left blank however
func NewStatus(result interface{}, success bool, message string) Status {
	return Status{result, success, message}
}

// NewStatusValidationError returns a Status from a set of validation errors
func NewStatusValidationError(ve validator.ValidationErrors) Status {
	var errors []string
	for _, e := range ve {
		errors = append(errors, fmt.Sprintf("validation '%s' failed for '%s'", e.Tag(), e.Field()))
	}

	return Status{
		errors,
		false,
		"payload validation failed",
	}
}

// ExampleStatus returns an example of Status
func ExampleStatus(result interface{}, success bool) Status {
	if !success {
		return Status{
			Result:  result,
			Success: false,
			Message: "error occurred",
		}
	}
	return Status{
		Result:  result,
		Success: true,
	}
}
