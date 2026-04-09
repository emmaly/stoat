package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPushSubscribe(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/push/subscribe" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body WebPushSubscription
		json.NewDecoder(r.Body).Decode(&body)
		if body.Endpoint != "https://push.example.com" {
			t.Errorf("endpoint = %q", body.Endpoint)
		}
		if body.P256DH != "key123" {
			t.Errorf("p256dh = %q", body.P256DH)
		}
		if body.Auth != "auth456" {
			t.Errorf("auth = %q", body.Auth)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	err := c.PushSubscribe(context.Background(), WebPushSubscription{
		Endpoint: "https://push.example.com",
		P256DH:   "key123",
		Auth:     "auth456",
	})
	if err != nil {
		t.Fatalf("PushSubscribe: %v", err)
	}
}

func TestPushUnsubscribe(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/push/unsubscribe" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.PushUnsubscribe(context.Background()); err != nil {
		t.Fatalf("PushUnsubscribe: %v", err)
	}
}
