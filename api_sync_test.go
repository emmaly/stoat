package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchSettings(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/sync/settings/fetch" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body OptionsFetchSettings
		json.NewDecoder(r.Body).Decode(&body)
		if len(body.Keys) != 1 || body.Keys[0] != "theme" {
			t.Errorf("keys = %v", body.Keys)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"theme": []any{1234567890, "dark"},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	settings, err := c.FetchSettings(context.Background(), OptionsFetchSettings{Keys: []string{"theme"}})
	if err != nil {
		t.Fatalf("FetchSettings: %v", err)
	}
	if _, ok := settings["theme"]; !ok {
		t.Error("missing 'theme' key")
	}
}

func TestSetSettings(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/sync/settings/set" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body map[string]string
		json.NewDecoder(r.Body).Decode(&body)
		if body["theme"] != "dark" {
			t.Errorf("theme = %q", body["theme"])
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.SetSettings(context.Background(), map[string]string{"theme": "dark"}); err != nil {
		t.Fatalf("SetSettings: %v", err)
	}
}

func TestFetchUnreads(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/sync/unreads" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]map[string]any{
			{
				"_id":      map[string]any{"channel": "ch01", "user": "user01"},
				"last_id":  "msg99",
				"mentions": []string{"msg50"},
			},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	unreads, err := c.FetchUnreads(context.Background())
	if err != nil {
		t.Fatalf("FetchUnreads: %v", err)
	}
	if len(unreads) != 1 {
		t.Fatalf("len = %d", len(unreads))
	}
	if unreads[0].ID.Channel != "ch01" {
		t.Errorf("Channel = %q", unreads[0].ID.Channel)
	}
	if unreads[0].LastID == nil || *unreads[0].LastID != "msg99" {
		t.Errorf("LastID = %v", unreads[0].LastID)
	}
	if len(unreads[0].Mentions) != 1 || unreads[0].Mentions[0] != "msg50" {
		t.Errorf("Mentions = %v", unreads[0].Mentions)
	}
}
