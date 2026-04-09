package stoat

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/auth/account/create" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "" {
			t.Error("unexpected session token")
		}
		var body DataCreateAccount
		json.NewDecoder(r.Body).Decode(&body)
		if body.Email != "test@example.com" {
			t.Errorf("email = %q", body.Email)
		}
		if body.Password != "password123" {
			t.Errorf("password = %q", body.Password)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	err := c.CreateAccount(context.Background(), DataCreateAccount{
		Email:    "test@example.com",
		Password: "password123",
	})
	if err != nil {
		t.Fatalf("CreateAccount: %v", err)
	}
}

func TestFetchAccount(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want GET", r.Method)
		}
		if r.URL.Path != "/auth/account/" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		json.NewEncoder(w).Encode(AccountInfo{ID: "acct01", Email: "test@example.com"})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	info, err := c.FetchAccount(context.Background())
	if err != nil {
		t.Fatalf("FetchAccount: %v", err)
	}
	if info.ID != "acct01" {
		t.Errorf("ID = %q", info.ID)
	}
	if info.Email != "test@example.com" {
		t.Errorf("Email = %q", info.Email)
	}
}

func TestChangeEmail(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q, want PATCH", r.Method)
		}
		if r.URL.Path != "/auth/account/change/email" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		var body DataChangeEmail
		json.NewDecoder(r.Body).Decode(&body)
		if body.Email != "new@example.com" {
			t.Errorf("email = %q", body.Email)
		}
		if body.CurrentPassword != "oldpass" {
			t.Errorf("current_password = %q", body.CurrentPassword)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	err := c.ChangeEmail(context.Background(), DataChangeEmail{
		Email:           "new@example.com",
		CurrentPassword: "oldpass",
	})
	if err != nil {
		t.Fatalf("ChangeEmail: %v", err)
	}
}

func TestChangePassword(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q, want PATCH", r.Method)
		}
		if r.URL.Path != "/auth/account/change/password" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		var body DataChangePassword
		json.NewDecoder(r.Body).Decode(&body)
		if body.Password != "newpass" {
			t.Errorf("password = %q", body.Password)
		}
		if body.CurrentPassword != "oldpass" {
			t.Errorf("current_password = %q", body.CurrentPassword)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	err := c.ChangePassword(context.Background(), DataChangePassword{
		Password:        "newpass",
		CurrentPassword: "oldpass",
	})
	if err != nil {
		t.Fatalf("ChangePassword: %v", err)
	}
}

func TestSendPasswordReset(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/auth/account/reset_password" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "" {
			t.Error("unexpected session token")
		}
		var body DataSendPasswordReset
		json.NewDecoder(r.Body).Decode(&body)
		if body.Email != "test@example.com" {
			t.Errorf("email = %q", body.Email)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	err := c.SendPasswordReset(context.Background(), DataSendPasswordReset{
		Email: "test@example.com",
	})
	if err != nil {
		t.Fatalf("SendPasswordReset: %v", err)
	}
}

func TestPasswordReset(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q, want PATCH", r.Method)
		}
		if r.URL.Path != "/auth/account/reset_password" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body DataPasswordReset
		json.NewDecoder(r.Body).Decode(&body)
		if body.Token != "reset-tok" {
			t.Errorf("token = %q", body.Token)
		}
		if body.Password != "newpass" {
			t.Errorf("password = %q", body.Password)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	err := c.PasswordReset(context.Background(), DataPasswordReset{
		Token:    "reset-tok",
		Password: "newpass",
	})
	if err != nil {
		t.Fatalf("PasswordReset: %v", err)
	}
}

func TestVerifyEmail(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/auth/account/verify/abc123" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	err := c.VerifyEmail(context.Background(), "abc123")
	if err != nil {
		t.Fatalf("VerifyEmail: %v", err)
	}
}

func TestResendVerification(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/auth/account/reverify" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body DataResendVerification
		json.NewDecoder(r.Body).Decode(&body)
		if body.Email != "test@example.com" {
			t.Errorf("email = %q", body.Email)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	err := c.ResendVerification(context.Background(), DataResendVerification{
		Email: "test@example.com",
	})
	if err != nil {
		t.Fatalf("ResendVerification: %v", err)
	}
}

func TestDeleteAccount(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/auth/account/delete" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		if r.Header.Get("X-MFA-Ticket") != "mfa-tok" {
			t.Errorf("mfa ticket = %q", r.Header.Get("X-MFA-Ticket"))
		}
		// Verify no body
		body, _ := io.ReadAll(r.Body)
		if len(body) > 0 {
			t.Errorf("unexpected body: %s", body)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	c.SetMFATicket("mfa-tok")
	err := c.DeleteAccount(context.Background())
	if err != nil {
		t.Fatalf("DeleteAccount: %v", err)
	}
}

func TestConfirmDeletion(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q, want PUT", r.Method)
		}
		if r.URL.Path != "/auth/account/delete" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body DataAccountDeletion
		json.NewDecoder(r.Body).Decode(&body)
		if body.Token != "del-tok" {
			t.Errorf("token = %q", body.Token)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	err := c.ConfirmDeletion(context.Background(), DataAccountDeletion{Token: "del-tok"})
	if err != nil {
		t.Fatalf("ConfirmDeletion: %v", err)
	}
}

func TestDisableAccount(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/auth/account/disable" {
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
	err := c.DisableAccount(context.Background())
	if err != nil {
		t.Fatalf("DisableAccount: %v", err)
	}
}
