package stoat

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddReaction(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q", r.Method)
		}
		expected := "/channels/ch01/messages/msg01/reactions/\U0001f44d"
		if r.URL.Path != expected {
			t.Errorf("path = %q, want %q", r.URL.Path, expected)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.AddReaction(context.Background(), "ch01", "msg01", "\U0001f44d"); err != nil {
		t.Fatalf("AddReaction: %v", err)
	}
}

func TestRemoveReaction(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Query().Get("user_id") != "user01" {
			t.Errorf("user_id = %q", r.URL.Query().Get("user_id"))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	uid := "user01"
	if err := c.RemoveReaction(context.Background(), "ch01", "msg01", "emoji01", &RemoveReactionOptions{UserID: &uid}); err != nil {
		t.Fatalf("RemoveReaction: %v", err)
	}
}

func TestRemoveReactionRemoveAll(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("remove_all") != "true" {
			t.Errorf("remove_all = %q", r.URL.Query().Get("remove_all"))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	removeAll := true
	if err := c.RemoveReaction(context.Background(), "ch01", "msg01", "emoji01", &RemoveReactionOptions{RemoveAll: &removeAll}); err != nil {
		t.Fatalf("RemoveReaction: %v", err)
	}
}

func TestClearReactions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/messages/msg01/reactions" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.ClearReactions(context.Background(), "ch01", "msg01"); err != nil {
		t.Fatalf("ClearReactions: %v", err)
	}
}
