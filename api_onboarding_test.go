package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckOnboarding(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want GET", r.Method)
		}
		if r.URL.Path != "/onboard/hello" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		json.NewEncoder(w).Encode(DataHello{Onboarding: true})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	hello, err := c.CheckOnboarding(context.Background())
	if err != nil {
		t.Fatalf("CheckOnboarding: %v", err)
	}
	if !hello.Onboarding {
		t.Error("expected onboarding = true")
	}
}

func TestCheckOnboardingComplete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(DataHello{Onboarding: false})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	hello, err := c.CheckOnboarding(context.Background())
	if err != nil {
		t.Fatalf("CheckOnboarding: %v", err)
	}
	if hello.Onboarding {
		t.Error("expected onboarding = false")
	}
}

func TestCompleteOnboarding(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/onboard/complete" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		var body DataOnboard
		json.NewDecoder(r.Body).Decode(&body)
		if body.Username != "cooluser" {
			t.Errorf("username = %q", body.Username)
		}
		// Server returns a User object, but we ignore it for now.
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":      "user01",
			"username": "cooluser",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	err := c.CompleteOnboarding(context.Background(), DataOnboard{Username: "cooluser"})
	if err != nil {
		t.Fatalf("CompleteOnboarding: %v", err)
	}
}
