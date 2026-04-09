package stoat

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestAPIErrorInterface(t *testing.T) {
	var err error = &APIError{Type: "NotFound"}
	if err.Error() != "NotFound" {
		t.Errorf("Error() = %q, want %q", err.Error(), "NotFound")
	}
}

func TestAPIErrorWithFields(t *testing.T) {
	err := &APIError{
		Type:       "MissingPermission",
		Permission: "ManageChannel",
	}
	got := err.Error()
	if got != "MissingPermission: permission=ManageChannel" {
		t.Errorf("Error() = %q", got)
	}
}

func TestAPIErrorWithMax(t *testing.T) {
	err := &APIError{
		Type: "TooManyReplies",
		Max:  5,
	}
	got := err.Error()
	if got != "TooManyReplies: max=5" {
		t.Errorf("Error() = %q", got)
	}
}

func TestAPIErrorIs(t *testing.T) {
	err := &APIError{Type: "NotFound"}
	if !errors.Is(err, ErrNotFound) {
		t.Error("expected errors.Is(err, ErrNotFound) to be true")
	}
	if errors.Is(err, ErrInternalError) {
		t.Error("expected errors.Is(err, ErrInternalError) to be false")
	}
}

func TestAPIErrorIsSentinels(t *testing.T) {
	tests := []struct {
		typ      string
		sentinel error
	}{
		{"NotFound", ErrNotFound},
		{"Unauthorised", ErrUnauthorised},
		{"InvalidSession", ErrInvalidSession},
		{"InternalError", ErrInternalError},
		{"MissingPermission", ErrMissingPermission},
		{"NotOwner", ErrNotOwner},
		{"TooManyRequests", ErrTooManyRequests},
		{"AlreadyAuthenticated", ErrAlreadyAuthenticated},
		{"OnboardingNotFinished", ErrOnboardingNotFinished},
	}
	for _, tt := range tests {
		t.Run(tt.typ, func(t *testing.T) {
			err := &APIError{Type: tt.typ}
			if !errors.Is(err, tt.sentinel) {
				t.Errorf("expected errors.Is(&APIError{Type: %q}, sentinel) to be true", tt.typ)
			}
		})
	}
}

func TestAPIErrorUnmarshal(t *testing.T) {
	raw := `{"type":"MissingPermission","permission":"ManageChannel"}`
	var apiErr APIError
	if err := json.Unmarshal([]byte(raw), &apiErr); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if apiErr.Type != "MissingPermission" {
		t.Errorf("Type = %q, want MissingPermission", apiErr.Type)
	}
	if apiErr.Permission != "ManageChannel" {
		t.Errorf("Permission = %q, want ManageChannel", apiErr.Permission)
	}
}

func TestAPIErrorUnmarshalWithMax(t *testing.T) {
	raw := `{"type":"TooManyAttachments","max":10}`
	var apiErr APIError
	if err := json.Unmarshal([]byte(raw), &apiErr); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if apiErr.Type != "TooManyAttachments" {
		t.Errorf("Type = %q", apiErr.Type)
	}
	if apiErr.Max != 10 {
		t.Errorf("Max = %d, want 10", apiErr.Max)
	}
}

func TestAPIErrorUnmarshalDatabaseError(t *testing.T) {
	raw := `{"type":"DatabaseError","operation":"find","collection":"messages"}`
	var apiErr APIError
	if err := json.Unmarshal([]byte(raw), &apiErr); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if apiErr.Operation != "find" {
		t.Errorf("Operation = %q", apiErr.Operation)
	}
	if apiErr.Collection != "messages" {
		t.Errorf("Collection = %q", apiErr.Collection)
	}
}
