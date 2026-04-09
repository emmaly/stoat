package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchDMs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want GET", r.Method)
		}
		if r.URL.Path != "/users/dms" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]map[string]any{
			{
				"channel_type": "SavedMessages",
				"_id":          "ch01",
				"user":         "user01",
			},
			{
				"channel_type":    "DirectMessage",
				"_id":             "ch02",
				"active":          true,
				"recipients":      []string{"user01", "user02"},
				"last_message_id": "msg01",
			},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	channels, err := c.FetchDMs(context.Background())
	if err != nil {
		t.Fatalf("FetchDMs: %v", err)
	}
	if len(channels) != 2 {
		t.Fatalf("len = %d, want 2", len(channels))
	}

	// Verify we can unmarshal the raw messages
	var ch0 map[string]any
	if err := json.Unmarshal(channels[0], &ch0); err != nil {
		t.Fatalf("unmarshal ch0: %v", err)
	}
	if ch0["channel_type"] != "SavedMessages" {
		t.Errorf("ch0 channel_type = %v", ch0["channel_type"])
	}

	var ch1 map[string]any
	if err := json.Unmarshal(channels[1], &ch1); err != nil {
		t.Fatalf("unmarshal ch1: %v", err)
	}
	if ch1["channel_type"] != "DirectMessage" {
		t.Errorf("ch1 channel_type = %v", ch1["channel_type"])
	}
}

func TestOpenDM(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want GET", r.Method)
		}
		if r.URL.Path != "/users/user02/dm" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"channel_type":    "DirectMessage",
			"_id":             "ch02",
			"active":          true,
			"recipients":      []string{"user01", "user02"},
			"last_message_id": "msg01",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	ch, err := c.OpenDM(context.Background(), "user02")
	if err != nil {
		t.Fatalf("OpenDM: %v", err)
	}

	var m map[string]any
	if err := json.Unmarshal(ch, &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if m["channel_type"] != "DirectMessage" {
		t.Errorf("channel_type = %v", m["channel_type"])
	}
	if m["_id"] != "ch02" {
		t.Errorf("_id = %v", m["_id"])
	}
}
