package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJoinCall(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/join_call" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"token": "voice-token-123",
			"url":   "wss://voice.example.com",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	resp, err := c.JoinCall(context.Background(), "ch01")
	if err != nil {
		t.Fatalf("JoinCall: %v", err)
	}
	if resp.Token != "voice-token-123" {
		t.Errorf("Token = %q", resp.Token)
	}
	if resp.URL != "wss://voice.example.com" {
		t.Errorf("URL = %q", resp.URL)
	}
}

func TestStopRing(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/end_ring/user01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.StopRing(context.Background(), "ch01", "user01"); err != nil {
		t.Fatalf("StopRing: %v", err)
	}
}
