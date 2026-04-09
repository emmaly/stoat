package stoat

import "fmt"

// APIError represents an error returned by the Stoat API.
// The Type field is the discriminator (e.g. "NotFound", "MissingPermission").
// Optional fields are populated depending on the error type.
type APIError struct {
	Type       string `json:"type"`
	Permission string `json:"permission,omitempty"`
	Operation  string `json:"operation,omitempty"`
	Collection string `json:"collection,omitempty"`
	Max        int    `json:"max,omitempty"`
	RetryAfter int    `json:"retry_after,omitempty"`
	Feature    string `json:"feature,omitempty"`
	Error_     string `json:"error,omitempty"` // for FailedValidation
}

// Error implements the error interface.
func (e *APIError) Error() string {
	base := e.Type
	switch {
	case e.Permission != "":
		return fmt.Sprintf("%s: permission=%s", base, e.Permission)
	case e.Max != 0:
		return fmt.Sprintf("%s: max=%d", base, e.Max)
	case e.Operation != "":
		s := fmt.Sprintf("%s: operation=%s", base, e.Operation)
		if e.Collection != "" {
			s += fmt.Sprintf(", collection=%s", e.Collection)
		}
		return s
	case e.RetryAfter != 0:
		return fmt.Sprintf("%s: retry_after=%d", base, e.RetryAfter)
	case e.Feature != "":
		return fmt.Sprintf("%s: feature=%s", base, e.Feature)
	case e.Error_ != "":
		return fmt.Sprintf("%s: error=%s", base, e.Error_)
	}
	return base
}

// Is supports errors.Is matching by comparing the Type field.
func (e *APIError) Is(target error) bool {
	if t, ok := target.(*APIError); ok {
		return e.Type == t.Type
	}
	return false
}

// Sentinel errors for common API error types.
var (
	ErrNotFound             = &APIError{Type: "NotFound"}
	ErrUnauthorised         = &APIError{Type: "Unauthorised"}
	ErrInvalidSession       = &APIError{Type: "InvalidSession"}
	ErrInternalError        = &APIError{Type: "InternalError"}
	ErrMissingPermission    = &APIError{Type: "MissingPermission"}
	ErrNotOwner             = &APIError{Type: "NotOwner"}
	ErrTooManyRequests      = &APIError{Type: "TooManyRequests"}
	ErrAlreadyAuthenticated = &APIError{Type: "AlreadyAuthenticated"}
	ErrOnboardingNotFinished = &APIError{Type: "OnboardingNotFinished"}
)
