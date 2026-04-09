package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchSelf(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want GET", r.Method)
		}
		if r.URL.Path != "/users/@me" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":           "user01",
			"username":      "alice",
			"discriminator": "0001",
			"relationship":  "User",
			"online":        true,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	u, err := c.FetchSelf(context.Background())
	if err != nil {
		t.Fatalf("FetchSelf: %v", err)
	}
	if u.ID != "user01" {
		t.Errorf("ID = %q", u.ID)
	}
	if u.Username != "alice" {
		t.Errorf("Username = %q", u.Username)
	}
	if u.Discriminator != "0001" {
		t.Errorf("Discriminator = %q", u.Discriminator)
	}
	if u.Relationship != RelationshipStatusUser {
		t.Errorf("Relationship = %q", u.Relationship)
	}
	if !u.Online {
		t.Error("expected online = true")
	}
}

func TestFetchUser(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want GET", r.Method)
		}
		if r.URL.Path != "/users/user02" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":           "user02",
			"username":      "bob",
			"discriminator": "0002",
			"relationship":  "Friend",
			"online":        false,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	u, err := c.FetchUser(context.Background(), "user02")
	if err != nil {
		t.Fatalf("FetchUser: %v", err)
	}
	if u.ID != "user02" {
		t.Errorf("ID = %q", u.ID)
	}
	if u.Relationship != RelationshipStatusFriend {
		t.Errorf("Relationship = %q", u.Relationship)
	}
}

func TestEditUser(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q, want PATCH", r.Method)
		}
		if r.URL.Path != "/users/user01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		var body DataEditUser
		json.NewDecoder(r.Body).Decode(&body)
		if body.DisplayName == nil || *body.DisplayName != "Alice W" {
			t.Errorf("display_name = %v", body.DisplayName)
		}
		w.Header().Set("Content-Type", "application/json")
		dn := "Alice W"
		json.NewEncoder(w).Encode(map[string]any{
			"_id":           "user01",
			"username":      "alice",
			"discriminator": "0001",
			"display_name":  dn,
			"relationship":  "User",
			"online":        true,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	dn := "Alice W"
	u, err := c.EditUser(context.Background(), "user01", DataEditUser{DisplayName: &dn})
	if err != nil {
		t.Fatalf("EditUser: %v", err)
	}
	if u.DisplayName == nil || *u.DisplayName != "Alice W" {
		t.Errorf("DisplayName = %v", u.DisplayName)
	}
}

func TestChangeUsername(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q, want PATCH", r.Method)
		}
		if r.URL.Path != "/users/@me/username" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body DataChangeUsername
		json.NewDecoder(r.Body).Decode(&body)
		if body.Username != "alice2" {
			t.Errorf("username = %q", body.Username)
		}
		if body.Password != "pass123" {
			t.Errorf("password = %q", body.Password)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":           "user01",
			"username":      "alice2",
			"discriminator": "0001",
			"relationship":  "User",
			"online":        true,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	u, err := c.ChangeUsername(context.Background(), DataChangeUsername{Username: "alice2", Password: "pass123"})
	if err != nil {
		t.Fatalf("ChangeUsername: %v", err)
	}
	if u.Username != "alice2" {
		t.Errorf("Username = %q", u.Username)
	}
}

func TestFetchUserProfile(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want GET", r.Method)
		}
		if r.URL.Path != "/users/user02/profile" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"content": "Hello, world!",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	p, err := c.FetchUserProfile(context.Background(), "user02")
	if err != nil {
		t.Fatalf("FetchUserProfile: %v", err)
	}
	if p.Content == nil || *p.Content != "Hello, world!" {
		t.Errorf("Content = %v", p.Content)
	}
}

func TestFetchDefaultAvatar(t *testing.T) {
	pngData := []byte{0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a} // PNG magic bytes
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want GET", r.Method)
		}
		if r.URL.Path != "/users/user02/default_avatar" {
			t.Errorf("path = %q", r.URL.Path)
		}
		// No auth required
		if r.Header.Get("X-Session-Token") != "" {
			t.Error("unexpected session token")
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngData)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	data, err := c.FetchDefaultAvatar(context.Background(), "user02")
	if err != nil {
		t.Fatalf("FetchDefaultAvatar: %v", err)
	}
	if len(data) != len(pngData) {
		t.Fatalf("len = %d, want %d", len(data), len(pngData))
	}
	for i, b := range data {
		if b != pngData[i] {
			t.Errorf("byte[%d] = %x, want %x", i, b, pngData[i])
		}
	}
}

func TestFetchDefaultAvatarError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"type":"NotFound"}`))
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	_, err := c.FetchDefaultAvatar(context.Background(), "baduser")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestFetchUserFlags(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want GET", r.Method)
		}
		if r.URL.Path != "/users/user02/flags" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"flags": 3})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	f, err := c.FetchUserFlags(context.Background(), "user02")
	if err != nil {
		t.Fatalf("FetchUserFlags: %v", err)
	}
	if f.Flags != 3 {
		t.Errorf("Flags = %d", f.Flags)
	}
}

func TestFetchMutual(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q, want GET", r.Method)
		}
		if r.URL.Path != "/users/user02/mutual" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"users":    []string{"user03", "user04"},
			"servers":  []string{"srv01"},
			"channels": []string{"ch01"},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	m, err := c.FetchMutual(context.Background(), "user02")
	if err != nil {
		t.Fatalf("FetchMutual: %v", err)
	}
	if len(m.Users) != 2 {
		t.Errorf("Users len = %d", len(m.Users))
	}
	if len(m.Servers) != 1 {
		t.Errorf("Servers len = %d", len(m.Servers))
	}
	if len(m.Channels) != 1 {
		t.Errorf("Channels len = %d", len(m.Channels))
	}
}
