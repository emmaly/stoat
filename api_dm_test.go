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

	if channels[0].ChannelType() != "SavedMessages" {
		t.Errorf("ch0 type = %q", channels[0].ChannelType())
	}
	if channels[0].ChannelID() != "ch01" {
		t.Errorf("ch0 id = %q", channels[0].ChannelID())
	}

	dm, ok := channels[1].(*DirectMessageChannel)
	if !ok {
		t.Fatalf("ch1 type = %T, want *DirectMessageChannel", channels[1])
	}
	if dm.ID != "ch02" {
		t.Errorf("ch1 id = %q", dm.ID)
	}
	if !dm.Active {
		t.Error("ch1 expected active = true")
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

	dm, ok := ch.(*DirectMessageChannel)
	if !ok {
		t.Fatalf("type = %T, want *DirectMessageChannel", ch)
	}
	if dm.ID != "ch02" {
		t.Errorf("id = %q", dm.ID)
	}
	if dm.ChannelType() != "DirectMessage" {
		t.Errorf("ChannelType = %q", dm.ChannelType())
	}
}
