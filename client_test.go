package stoat

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	c, err := New("https://api.example.com")
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if c == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestNewClientWithHTTPClient(t *testing.T) {
	custom := &http.Client{}
	c, err := New("https://api.example.com", WithHTTPClient(custom))
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if c.httpClient != custom {
		t.Error("expected custom HTTP client to be used")
	}
}

func TestSessionTokenHeader(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got := r.Header.Get("X-Session-Token")
		if got != "my-session-token" {
			t.Errorf("X-Session-Token = %q, want my-session-token", got)
		}
		w.Header().Set("X-RateLimit-Limit", "20")
		w.Header().Set("X-RateLimit-Bucket", "test")
		w.Header().Set("X-RateLimit-Remaining", "19")
		w.Header().Set("X-RateLimit-Reset-After", "10000")
		json.NewEncoder(w).Encode(map[string]string{"ok": "true"})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("my-session-token")

	req, err := c.request(context.Background(), http.MethodGet, "/test", nil)
	if err != nil {
		t.Fatalf("request: %v", err)
	}

	var result map[string]string
	if err := c.do(req, &result); err != nil {
		t.Fatalf("do: %v", err)
	}
	if result["ok"] != "true" {
		t.Errorf("result = %v", result)
	}
}

func TestBotTokenHeader(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got := r.Header.Get("X-Bot-Token")
		if got != "my-bot-token" {
			t.Errorf("X-Bot-Token = %q, want my-bot-token", got)
		}
		w.Header().Set("X-RateLimit-Limit", "20")
		w.Header().Set("X-RateLimit-Bucket", "test")
		w.Header().Set("X-RateLimit-Remaining", "19")
		w.Header().Set("X-RateLimit-Reset-After", "10000")
		json.NewEncoder(w).Encode(map[string]string{})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetBotToken("my-bot-token")

	req, _ := c.request(context.Background(), http.MethodGet, "/test", nil)
	if err := c.do(req, nil); err != nil {
		t.Fatalf("do: %v", err)
	}
}

func TestMFATicketHeader(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got := r.Header.Get("X-MFA-Ticket")
		if got != "my-mfa-ticket" {
			t.Errorf("X-MFA-Ticket = %q, want my-mfa-ticket", got)
		}
		w.Header().Set("X-RateLimit-Limit", "20")
		w.Header().Set("X-RateLimit-Bucket", "test")
		w.Header().Set("X-RateLimit-Remaining", "19")
		w.Header().Set("X-RateLimit-Reset-After", "10000")
		json.NewEncoder(w).Encode(map[string]string{})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetMFATicket("my-mfa-ticket")

	req, _ := c.request(context.Background(), http.MethodGet, "/test", nil)
	if err := c.do(req, nil); err != nil {
		t.Fatalf("do: %v", err)
	}
}

func TestErrorResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-RateLimit-Limit", "20")
		w.Header().Set("X-RateLimit-Bucket", "test")
		w.Header().Set("X-RateLimit-Remaining", "19")
		w.Header().Set("X-RateLimit-Reset-After", "10000")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"type": "NotFound"})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	req, _ := c.request(context.Background(), http.MethodGet, "/missing", nil)
	err := c.do(req, nil)
	if err == nil {
		t.Fatal("expected error")
	}

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected *APIError, got %T: %v", err, err)
	}
	if apiErr.Type != "NotFound" {
		t.Errorf("Type = %q, want NotFound", apiErr.Type)
	}
	if !errors.Is(err, ErrNotFound) {
		t.Error("expected errors.Is(err, ErrNotFound)")
	}
}

func TestRateLimitParsing(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-RateLimit-Limit", "10")
		w.Header().Set("X-RateLimit-Bucket", "messages")
		w.Header().Set("X-RateLimit-Remaining", "3")
		w.Header().Set("X-RateLimit-Reset-After", "7000")
		json.NewEncoder(w).Encode(map[string]string{"ok": "true"})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	req, _ := c.request(context.Background(), http.MethodGet, "/test", nil)
	if err := c.do(req, nil); err != nil {
		t.Fatalf("do: %v", err)
	}

	rl := c.LastRateLimit()
	if rl == nil {
		t.Fatal("expected non-nil LastRateLimit")
	}
	if rl.Limit != 10 {
		t.Errorf("Limit = %d, want 10", rl.Limit)
	}
	if rl.Bucket != "messages" {
		t.Errorf("Bucket = %q, want messages", rl.Bucket)
	}
	if rl.Remaining != 3 {
		t.Errorf("Remaining = %d, want 3", rl.Remaining)
	}
	if rl.ResetAfter != 7000 {
		t.Errorf("ResetAfter = %d, want 7000", rl.ResetAfter)
	}
}

func TestRequestWithBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Content-Type = %q", r.Header.Get("Content-Type"))
		}
		var body map[string]string
		json.NewDecoder(r.Body).Decode(&body)
		if body["name"] != "test" {
			t.Errorf("body name = %q", body["name"])
		}
		w.Header().Set("X-RateLimit-Limit", "20")
		w.Header().Set("X-RateLimit-Bucket", "test")
		w.Header().Set("X-RateLimit-Remaining", "19")
		w.Header().Set("X-RateLimit-Reset-After", "10000")
		json.NewEncoder(w).Encode(map[string]string{"ok": "true"})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	req, _ := c.request(context.Background(), http.MethodPost, "/test", map[string]string{"name": "test"})
	if err := c.do(req, nil); err != nil {
		t.Fatalf("do: %v", err)
	}
}
