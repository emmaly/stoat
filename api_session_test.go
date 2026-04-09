package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/auth/session/login" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "" {
			t.Error("unexpected session token")
		}
		var body DataLogin
		json.NewDecoder(r.Body).Decode(&body)
		if body.Email != "test@example.com" {
			t.Errorf("email = %q", body.Email)
		}
		if body.Password != "pass123" {
			t.Errorf("password = %q", body.Password)
		}
		if body.FriendlyName != "My Device" {
			t.Errorf("friendly_name = %q", body.FriendlyName)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"result":  "Success",
			"_id":     "sess01",
			"user_id": "user01",
			"token":   "tok123",
			"name":    "My Device",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	resp, err := c.Login(context.Background(), DataLogin{
		Email:        "test@example.com",
		Password:     "pass123",
		FriendlyName: "My Device",
	})
	if err != nil {
		t.Fatalf("Login: %v", err)
	}
	s, ok := resp.(*LoginSuccess)
	if !ok {
		t.Fatalf("expected *LoginSuccess, got %T", resp)
	}
	if s.Token != "tok123" {
		t.Errorf("Token = %q", s.Token)
	}
	if s.UserID != "user01" {
		t.Errorf("UserID = %q", s.UserID)
	}
}

func TestLoginMFA(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"result":          "MFA",
			"ticket":          "mfa-ticket",
			"allowed_methods": []string{"Password", "Totp"},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	resp, err := c.Login(context.Background(), DataLogin{
		Email:    "test@example.com",
		Password: "pass123",
	})
	if err != nil {
		t.Fatalf("Login: %v", err)
	}
	m, ok := resp.(*LoginMFA)
	if !ok {
		t.Fatalf("expected *LoginMFA, got %T", resp)
	}
	if m.Ticket != "mfa-ticket" {
		t.Errorf("Ticket = %q", m.Ticket)
	}
}

func TestLogout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/auth/session/logout" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	err := c.Logout(context.Background())
	if err != nil {
		t.Fatalf("Logout: %v", err)
	}
}

func TestFetchSessions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want GET", r.Method)
		}
		if r.URL.Path != "/auth/session/all" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]SessionInfo{
			{ID: "s1", Name: "Device 1"},
			{ID: "s2", Name: "Device 2"},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	sessions, err := c.FetchSessions(context.Background())
	if err != nil {
		t.Fatalf("FetchSessions: %v", err)
	}
	if len(sessions) != 2 {
		t.Fatalf("len = %d", len(sessions))
	}
	if sessions[0].ID != "s1" {
		t.Errorf("sessions[0].ID = %q", sessions[0].ID)
	}
}

func TestRevokeSession(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q, want DELETE", r.Method)
		}
		if r.URL.Path != "/auth/session/sess42" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	err := c.RevokeSession(context.Background(), "sess42")
	if err != nil {
		t.Fatalf("RevokeSession: %v", err)
	}
}

func TestEditSession(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q, want PATCH", r.Method)
		}
		if r.URL.Path != "/auth/session/sess42" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		var body DataEditSession
		json.NewDecoder(r.Body).Decode(&body)
		if body.FriendlyName != "New Name" {
			t.Errorf("friendly_name = %q", body.FriendlyName)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SessionInfo{ID: "sess42", Name: "New Name"})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	info, err := c.EditSession(context.Background(), "sess42", DataEditSession{FriendlyName: "New Name"})
	if err != nil {
		t.Fatalf("EditSession: %v", err)
	}
	if info.Name != "New Name" {
		t.Errorf("Name = %q", info.Name)
	}
}

func TestDeleteAllSessions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q, want DELETE", r.Method)
		}
		if r.URL.Path != "/auth/session/all" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		if r.URL.Query().Get("revoke_self") != "true" {
			t.Errorf("revoke_self = %q", r.URL.Query().Get("revoke_self"))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	err := c.DeleteAllSessions(context.Background(), true)
	if err != nil {
		t.Fatalf("DeleteAllSessions: %v", err)
	}
}

func TestDeleteAllSessionsNoRevokeSelf(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("revoke_self") != "" {
			t.Errorf("unexpected revoke_self param: %q", r.URL.Query().Get("revoke_self"))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	err := c.DeleteAllSessions(context.Background(), false)
	if err != nil {
		t.Fatalf("DeleteAllSessions: %v", err)
	}
}
