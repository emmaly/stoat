package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchChannel(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"channel_type": "TextChannel",
			"_id":          "ch01",
			"server":       "srv01",
			"name":         "general",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	ch, err := c.FetchChannel(context.Background(), "ch01")
	if err != nil {
		t.Fatalf("FetchChannel: %v", err)
	}
	tc, ok := ch.(*TextChannel)
	if !ok {
		t.Fatalf("type = %T", ch)
	}
	if tc.Name != "general" {
		t.Errorf("Name = %q", tc.Name)
	}
}

func TestEditChannel(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body DataEditChannel
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name == nil || *body.Name != "renamed" {
			t.Errorf("name = %v", body.Name)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"channel_type": "TextChannel",
			"_id":          "ch01",
			"server":       "srv01",
			"name":         "renamed",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	name := "renamed"
	ch, err := c.EditChannel(context.Background(), "ch01", DataEditChannel{Name: &name})
	if err != nil {
		t.Fatalf("EditChannel: %v", err)
	}
	tc := ch.(*TextChannel)
	if tc.Name != "renamed" {
		t.Errorf("Name = %q", tc.Name)
	}
}

func TestCloseChannel(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01" {
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
	if err := c.CloseChannel(context.Background(), "ch01", true); err != nil {
		t.Fatalf("CloseChannel: %v", err)
	}
}

func TestSetDefaultChannelPermissions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/permissions/default" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"channel_type": "TextChannel",
			"_id":          "ch01",
			"server":       "srv01",
			"name":         "general",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	ch, err := c.SetDefaultChannelPermissions(context.Background(), "ch01", DataDefaultChannelPermissions{
		Permissions: Override{Allow: 100, Deny: 50},
	})
	if err != nil {
		t.Fatalf("SetDefaultChannelPermissions: %v", err)
	}
	if ch.ChannelType() != "TextChannel" {
		t.Errorf("type = %q", ch.ChannelType())
	}
}

func TestSetRoleChannelPermissions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/permissions/role01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"channel_type": "TextChannel",
			"_id":          "ch01",
			"server":       "srv01",
			"name":         "general",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	ch, err := c.SetRoleChannelPermissions(context.Background(), "ch01", "role01", DataSetRolePermissions{
		Permissions: Override{Allow: 200, Deny: 100},
	})
	if err != nil {
		t.Fatalf("SetRoleChannelPermissions: %v", err)
	}
	if ch.ChannelID() != "ch01" {
		t.Errorf("id = %q", ch.ChannelID())
	}
}
