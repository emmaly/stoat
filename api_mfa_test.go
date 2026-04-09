package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMFAStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want GET", r.Method)
		}
		if r.URL.Path != "/auth/mfa/" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		json.NewEncoder(w).Encode(MultiFactorStatus{
			TOTPMFA:        true,
			RecoveryActive: true,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	status, err := c.MFAStatus(context.Background())
	if err != nil {
		t.Fatalf("MFAStatus: %v", err)
	}
	if !status.TOTPMFA {
		t.Error("expected totp_mfa = true")
	}
	if !status.RecoveryActive {
		t.Error("expected recovery_active = true")
	}
	if status.EmailOTP {
		t.Error("expected email_otp = false")
	}
}

func TestGetMFAMethods(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want GET", r.Method)
		}
		if r.URL.Path != "/auth/mfa/methods" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		json.NewEncoder(w).Encode([]MFAMethod{MFAMethodPassword, MFAMethodTOTP})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	methods, err := c.GetMFAMethods(context.Background())
	if err != nil {
		t.Fatalf("GetMFAMethods: %v", err)
	}
	if len(methods) != 2 {
		t.Fatalf("len = %d", len(methods))
	}
	if methods[0] != MFAMethodPassword {
		t.Errorf("methods[0] = %q", methods[0])
	}
	if methods[1] != MFAMethodTOTP {
		t.Errorf("methods[1] = %q", methods[1])
	}
}

func TestCreateMFATicket(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q, want PUT", r.Method)
		}
		if r.URL.Path != "/auth/mfa/ticket" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		var body MFAResponse
		json.NewDecoder(r.Body).Decode(&body)
		if body.Password != "mypassword" {
			t.Errorf("password = %q", body.Password)
		}
		json.NewEncoder(w).Encode(MFATicket{
			ID:         "ticket01",
			AccountID:  "acct01",
			Token:      "tok-mfa",
			Validated:  true,
			Authorised: true,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	ticket, err := c.CreateMFATicket(context.Background(), MFAResponse{Password: "mypassword"})
	if err != nil {
		t.Fatalf("CreateMFATicket: %v", err)
	}
	if ticket.ID != "ticket01" {
		t.Errorf("ID = %q", ticket.ID)
	}
	if !ticket.Validated {
		t.Error("expected validated = true")
	}
	if ticket.Token != "tok-mfa" {
		t.Errorf("Token = %q", ticket.Token)
	}
}

func TestGenerateTOTPSecret(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/auth/mfa/totp" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		if r.Header.Get("X-MFA-Ticket") != "mfa-tok" {
			t.Errorf("mfa ticket = %q", r.Header.Get("X-MFA-Ticket"))
		}
		json.NewEncoder(w).Encode(ResponseTotpSecret{Secret: "JBSWY3DPEHPK3PXP"})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	c.SetMFATicket("mfa-tok")
	resp, err := c.GenerateTOTPSecret(context.Background())
	if err != nil {
		t.Fatalf("GenerateTOTPSecret: %v", err)
	}
	if resp.Secret != "JBSWY3DPEHPK3PXP" {
		t.Errorf("Secret = %q", resp.Secret)
	}
}

func TestEnableTOTP(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q, want PUT", r.Method)
		}
		if r.URL.Path != "/auth/mfa/totp" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		var body MFAResponse
		json.NewDecoder(r.Body).Decode(&body)
		if body.TOTPCode != "123456" {
			t.Errorf("totp_code = %q", body.TOTPCode)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	err := c.EnableTOTP(context.Background(), MFAResponse{TOTPCode: "123456"})
	if err != nil {
		t.Fatalf("EnableTOTP: %v", err)
	}
}

func TestDisableTOTP(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q, want DELETE", r.Method)
		}
		if r.URL.Path != "/auth/mfa/totp" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		if r.Header.Get("X-MFA-Ticket") != "mfa-tok" {
			t.Errorf("mfa ticket = %q", r.Header.Get("X-MFA-Ticket"))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	c.SetMFATicket("mfa-tok")
	err := c.DisableTOTP(context.Background())
	if err != nil {
		t.Fatalf("DisableTOTP: %v", err)
	}
}

func TestFetchRecoveryCodes(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/auth/mfa/recovery" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		if r.Header.Get("X-MFA-Ticket") != "mfa-tok" {
			t.Errorf("mfa ticket = %q", r.Header.Get("X-MFA-Ticket"))
		}
		json.NewEncoder(w).Encode([]string{"code1", "code2", "code3"})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	c.SetMFATicket("mfa-tok")
	codes, err := c.FetchRecoveryCodes(context.Background())
	if err != nil {
		t.Fatalf("FetchRecoveryCodes: %v", err)
	}
	if len(codes) != 3 {
		t.Fatalf("len = %d", len(codes))
	}
	if codes[0] != "code1" {
		t.Errorf("codes[0] = %q", codes[0])
	}
}

func TestGenerateRecoveryCodes(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q, want PATCH", r.Method)
		}
		if r.URL.Path != "/auth/mfa/recovery" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		if r.Header.Get("X-MFA-Ticket") != "mfa-tok" {
			t.Errorf("mfa ticket = %q", r.Header.Get("X-MFA-Ticket"))
		}
		json.NewEncoder(w).Encode([]string{"new1", "new2", "new3", "new4"})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	c.SetMFATicket("mfa-tok")
	codes, err := c.GenerateRecoveryCodes(context.Background())
	if err != nil {
		t.Fatalf("GenerateRecoveryCodes: %v", err)
	}
	if len(codes) != 4 {
		t.Fatalf("len = %d", len(codes))
	}
	if codes[0] != "new1" {
		t.Errorf("codes[0] = %q", codes[0])
	}
}
