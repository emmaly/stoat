package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateGroup(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/create" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body DataCreateGroup
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name != "Test Group" {
			t.Errorf("name = %q", body.Name)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"channel_type": "Group",
			"_id":          "ch01",
			"name":         "Test Group",
			"owner":        "user01",
			"recipients":   []string{"user01", "user02"},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	ch, err := c.CreateGroup(context.Background(), DataCreateGroup{
		Name:  "Test Group",
		Users: []string{"user02"},
	})
	if err != nil {
		t.Fatalf("CreateGroup: %v", err)
	}
	grp, ok := ch.(*GroupChannel)
	if !ok {
		t.Fatalf("type = %T", ch)
	}
	if grp.Name != "Test Group" {
		t.Errorf("Name = %q", grp.Name)
	}
}

func TestFetchGroupMembers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/members" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]map[string]any{
			{"_id": "user01", "username": "alice", "discriminator": "0001", "relationship": "User", "online": true},
			{"_id": "user02", "username": "bob", "discriminator": "0002", "relationship": "Friend", "online": false},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	users, err := c.FetchGroupMembers(context.Background(), "ch01")
	if err != nil {
		t.Fatalf("FetchGroupMembers: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("len = %d", len(users))
	}
	if users[0].Username != "alice" {
		t.Errorf("users[0].Username = %q", users[0].Username)
	}
}

func TestAddGroupMember(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/recipients/user02" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.AddGroupMember(context.Background(), "ch01", "user02"); err != nil {
		t.Fatalf("AddGroupMember: %v", err)
	}
}

func TestRemoveGroupMember(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/recipients/user02" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.RemoveGroupMember(context.Background(), "ch01", "user02"); err != nil {
		t.Fatalf("RemoveGroupMember: %v", err)
	}
}
