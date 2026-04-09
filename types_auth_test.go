package stoat

import (
	"encoding/json"
	"testing"
)

func TestResponseLoginSuccess(t *testing.T) {
	raw := `{"result":"Success","_id":"sess01","user_id":"user01","token":"tok123","name":"My Device"}`

	var r RawResponseLogin
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	s, ok := r.Result.(*LoginSuccess)
	if !ok {
		t.Fatalf("expected *LoginSuccess, got %T", r.Result)
	}
	if s.ID != "sess01" {
		t.Errorf("ID = %q, want %q", s.ID, "sess01")
	}
	if s.UserID != "user01" {
		t.Errorf("UserID = %q", s.UserID)
	}
	if s.Token != "tok123" {
		t.Errorf("Token = %q", s.Token)
	}
	if s.Name != "My Device" {
		t.Errorf("Name = %q", s.Name)
	}

	// Round-trip
	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var r2 RawResponseLogin
	if err := json.Unmarshal(b, &r2); err != nil {
		t.Fatalf("Unmarshal round-trip: %v", err)
	}
	s2, ok := r2.Result.(*LoginSuccess)
	if !ok {
		t.Fatalf("round-trip: expected *LoginSuccess, got %T", r2.Result)
	}
	if s2.Token != "tok123" {
		t.Errorf("round-trip Token = %q", s2.Token)
	}
}

func TestResponseLoginMFA(t *testing.T) {
	raw := `{"result":"MFA","ticket":"mfa-ticket-123","allowed_methods":["Password","Totp"]}`

	var r RawResponseLogin
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	m, ok := r.Result.(*LoginMFA)
	if !ok {
		t.Fatalf("expected *LoginMFA, got %T", r.Result)
	}
	if m.Ticket != "mfa-ticket-123" {
		t.Errorf("Ticket = %q", m.Ticket)
	}
	if len(m.AllowedMethods) != 2 {
		t.Fatalf("AllowedMethods len = %d", len(m.AllowedMethods))
	}
	if m.AllowedMethods[0] != MFAMethodPassword {
		t.Errorf("method[0] = %q", m.AllowedMethods[0])
	}
	if m.AllowedMethods[1] != MFAMethodTOTP {
		t.Errorf("method[1] = %q", m.AllowedMethods[1])
	}

	// Round-trip
	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var r2 RawResponseLogin
	if err := json.Unmarshal(b, &r2); err != nil {
		t.Fatalf("Unmarshal round-trip: %v", err)
	}
	m2, ok := r2.Result.(*LoginMFA)
	if !ok {
		t.Fatalf("round-trip: expected *LoginMFA, got %T", r2.Result)
	}
	if m2.Ticket != "mfa-ticket-123" {
		t.Errorf("round-trip Ticket = %q", m2.Ticket)
	}
}

func TestResponseLoginDisabled(t *testing.T) {
	raw := `{"result":"Disabled","user_id":"user99"}`

	var r RawResponseLogin
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	d, ok := r.Result.(*LoginDisabled)
	if !ok {
		t.Fatalf("expected *LoginDisabled, got %T", r.Result)
	}
	if d.UserID != "user99" {
		t.Errorf("UserID = %q", d.UserID)
	}

	// Round-trip
	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var r2 RawResponseLogin
	if err := json.Unmarshal(b, &r2); err != nil {
		t.Fatalf("Unmarshal round-trip: %v", err)
	}
	d2, ok := r2.Result.(*LoginDisabled)
	if !ok {
		t.Fatalf("round-trip: expected *LoginDisabled, got %T", r2.Result)
	}
	if d2.UserID != "user99" {
		t.Errorf("round-trip UserID = %q", d2.UserID)
	}
}

func TestResponseLoginUnknownResult(t *testing.T) {
	raw := `{"result":"Unknown"}`

	var r RawResponseLogin
	err := json.Unmarshal([]byte(raw), &r)
	if err == nil {
		t.Fatal("expected error for unknown result")
	}
}
