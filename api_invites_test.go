package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateInvite(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/invites" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Invite{
			ID:      "inv01",
			Server:  "srv01",
			Creator: "user01",
			Channel: "ch01",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	inv, err := c.CreateInvite(context.Background(), "ch01")
	if err != nil {
		t.Fatalf("CreateInvite: %v", err)
	}
	if inv.ID != "inv01" {
		t.Errorf("ID = %q", inv.ID)
	}
	if inv.Server != "srv01" {
		t.Errorf("Server = %q", inv.Server)
	}
}

func TestFetchInvite(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/invites/inv01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"type":         "Server",
			"code":         "inv01",
			"server_id":    "srv01",
			"server_name":  "My Server",
			"server_flags": 0,
			"channel_id":   "ch01",
			"channel_name": "general",
			"user_name":    "alice",
			"member_count": 42,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	resp, err := c.FetchInvite(context.Background(), "inv01")
	if err != nil {
		t.Fatalf("FetchInvite: %v", err)
	}
	si, ok := resp.(*ServerInviteResponse)
	if !ok {
		t.Fatalf("type = %T", resp)
	}
	if si.ServerName != "My Server" {
		t.Errorf("ServerName = %q", si.ServerName)
	}
	if si.MemberCount != 42 {
		t.Errorf("MemberCount = %d", si.MemberCount)
	}
}

func TestJoinInvite(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/invites/inv01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"type": "Server",
			"channels": []map[string]any{
				{"channel_type": "TextChannel", "_id": "ch01", "server": "srv01", "name": "general"},
			},
			"server": map[string]any{
				"_id":                 "srv01",
				"owner":              "user01",
				"name":               "My Server",
				"channels":           []string{"ch01"},
				"default_permissions": 0,
			},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	resp, err := c.JoinInvite(context.Background(), "inv01")
	if err != nil {
		t.Fatalf("JoinInvite: %v", err)
	}
	sj, ok := resp.(*ServerInviteJoinResponse)
	if !ok {
		t.Fatalf("type = %T", resp)
	}
	if sj.Server.Name != "My Server" {
		t.Errorf("Server.Name = %q", sj.Server.Name)
	}
	if len(sj.Channels) != 1 {
		t.Fatalf("Channels len = %d", len(sj.Channels))
	}
}

func TestDeleteInvite(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/invites/inv01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.DeleteInvite(context.Background(), "inv01"); err != nil {
		t.Fatalf("DeleteInvite: %v", err)
	}
}
