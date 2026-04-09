package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendFriendRequest(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/users/friend" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("X-Session-Token") != "sess-tok" {
			t.Errorf("session token = %q", r.Header.Get("X-Session-Token"))
		}
		var body DataSendFriendRequest
		json.NewDecoder(r.Body).Decode(&body)
		if body.Username != "bob#0002" {
			t.Errorf("username = %q", body.Username)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":           "user02",
			"username":      "bob",
			"discriminator": "0002",
			"relationship":  "Outgoing",
			"online":        false,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	u, err := c.SendFriendRequest(context.Background(), DataSendFriendRequest{Username: "bob#0002"})
	if err != nil {
		t.Fatalf("SendFriendRequest: %v", err)
	}
	if u.ID != "user02" {
		t.Errorf("ID = %q", u.ID)
	}
	if u.Relationship != RelationshipStatusOutgoing {
		t.Errorf("Relationship = %q", u.Relationship)
	}
}

func TestAcceptFriend(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q, want PUT", r.Method)
		}
		if r.URL.Path != "/users/user02/friend" {
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
			"online":        true,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	u, err := c.AcceptFriend(context.Background(), "user02")
	if err != nil {
		t.Fatalf("AcceptFriend: %v", err)
	}
	if u.Relationship != RelationshipStatusFriend {
		t.Errorf("Relationship = %q", u.Relationship)
	}
}

func TestRemoveFriend(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q, want DELETE", r.Method)
		}
		if r.URL.Path != "/users/user02/friend" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":           "user02",
			"username":      "bob",
			"discriminator": "0002",
			"relationship":  "None",
			"online":        false,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	u, err := c.RemoveFriend(context.Background(), "user02")
	if err != nil {
		t.Fatalf("RemoveFriend: %v", err)
	}
	if u.Relationship != RelationshipStatusNone {
		t.Errorf("Relationship = %q", u.Relationship)
	}
}

func TestBlockUser(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q, want PUT", r.Method)
		}
		if r.URL.Path != "/users/user02/block" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":           "user02",
			"username":      "bob",
			"discriminator": "0002",
			"relationship":  "Blocked",
			"online":        false,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	u, err := c.BlockUser(context.Background(), "user02")
	if err != nil {
		t.Fatalf("BlockUser: %v", err)
	}
	if u.Relationship != RelationshipStatusBlocked {
		t.Errorf("Relationship = %q", u.Relationship)
	}
}

func TestUnblockUser(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q, want DELETE", r.Method)
		}
		if r.URL.Path != "/users/user02/block" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":           "user02",
			"username":      "bob",
			"discriminator": "0002",
			"relationship":  "None",
			"online":        false,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("sess-tok")
	u, err := c.UnblockUser(context.Background(), "user02")
	if err != nil {
		t.Fatalf("UnblockUser: %v", err)
	}
	if u.Relationship != RelationshipStatusNone {
		t.Errorf("Relationship = %q", u.Relationship)
	}
}
