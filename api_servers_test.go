package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateServer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/create" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body DataCreateServer
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name != "My Server" {
			t.Errorf("name = %q", body.Name)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"server": map[string]any{
				"_id":                 "srv01",
				"owner":              "user01",
				"name":               "My Server",
				"channels":           []string{"ch01"},
				"default_permissions": 0,
			},
			"channels": []map[string]any{
				{"channel_type": "TextChannel", "_id": "ch01", "server": "srv01", "name": "general"},
			},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	resp, err := c.CreateServer(context.Background(), DataCreateServer{Name: "My Server"})
	if err != nil {
		t.Fatalf("CreateServer: %v", err)
	}
	if resp.Server.ID != "srv01" {
		t.Errorf("Server.ID = %q", resp.Server.ID)
	}
	if len(resp.Channels) != 1 {
		t.Fatalf("Channels len = %d", len(resp.Channels))
	}
}

func TestFetchServer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.URL.Query().Get("include_channels") != "true" {
			t.Errorf("include_channels = %q", r.URL.Query().Get("include_channels"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"server": map[string]any{
				"_id":                 "srv01",
				"owner":              "user01",
				"name":               "My Server",
				"channels":           []string{"ch01"},
				"default_permissions": 0,
			},
			"channels": []map[string]any{
				{"channel_type": "TextChannel", "_id": "ch01", "server": "srv01", "name": "general"},
			},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	resp, err := c.FetchServer(context.Background(), "srv01", true)
	if err != nil {
		t.Fatalf("FetchServer: %v", err)
	}
	if resp.Server.Name != "My Server" {
		t.Errorf("Server.Name = %q", resp.Server.Name)
	}
	if len(resp.Channels) != 1 {
		t.Fatalf("Channels len = %d", len(resp.Channels))
	}
}

func TestEditServer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body DataEditServer
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name == nil || *body.Name != "Renamed" {
			t.Errorf("name = %v", body.Name)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":                 "srv01",
			"owner":              "user01",
			"name":               "Renamed",
			"channels":           []string{},
			"default_permissions": 0,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	name := "Renamed"
	s, err := c.EditServer(context.Background(), "srv01", DataEditServer{Name: &name})
	if err != nil {
		t.Fatalf("EditServer: %v", err)
	}
	if s.Name != "Renamed" {
		t.Errorf("Name = %q", s.Name)
	}
}

func TestDeleteServer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.URL.Query().Get("leave_silently") != "true" {
			t.Errorf("leave_silently = %q", r.URL.Query().Get("leave_silently"))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.DeleteServer(context.Background(), "srv01", true); err != nil {
		t.Fatalf("DeleteServer: %v", err)
	}
}

func TestMarkServerRead(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/ack" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.MarkServerRead(context.Background(), "srv01"); err != nil {
		t.Fatalf("MarkServerRead: %v", err)
	}
}

func TestCreateServerChannel(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/channels" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body DataCreateServerChannel
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name != "new-channel" {
			t.Errorf("name = %q", body.Name)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"channel_type": "TextChannel",
			"_id":          "ch02",
			"server":       "srv01",
			"name":         "new-channel",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	ch, err := c.CreateServerChannel(context.Background(), "srv01", DataCreateServerChannel{Name: "new-channel"})
	if err != nil {
		t.Fatalf("CreateServerChannel: %v", err)
	}
	tc, ok := ch.(*TextChannel)
	if !ok {
		t.Fatalf("type = %T", ch)
	}
	if tc.Name != "new-channel" {
		t.Errorf("Name = %q", tc.Name)
	}
}
