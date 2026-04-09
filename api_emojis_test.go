package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchEmoji(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/custom/emoji/em01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":        "em01",
			"parent":     map[string]any{"type": "Server", "id": "srv01"},
			"creator_id": "user01",
			"name":       "pepe",
			"animated":   false,
			"nsfw":       false,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	e, err := c.FetchEmoji(context.Background(), "em01")
	if err != nil {
		t.Fatalf("FetchEmoji: %v", err)
	}
	if e.ID != "em01" {
		t.Errorf("ID = %q", e.ID)
	}
	if e.Name != "pepe" {
		t.Errorf("Name = %q", e.Name)
	}
	sp, ok := e.Parent.Value.(*ServerEmojiParent)
	if !ok {
		t.Fatalf("Parent type = %T", e.Parent.Value)
	}
	if sp.ID != "srv01" {
		t.Errorf("Parent.ID = %q", sp.ID)
	}
}

func TestCreateEmoji(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/custom/emoji/em01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body DataCreateEmoji
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name != "pepe" {
			t.Errorf("name = %q", body.Name)
		}
		sp, ok := body.Parent.Value.(*ServerEmojiParent)
		if !ok {
			t.Errorf("parent type = %T", body.Parent.Value)
		} else if sp.ID != "srv01" {
			t.Errorf("parent.id = %q", sp.ID)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":        "em01",
			"parent":     map[string]any{"type": "Server", "id": "srv01"},
			"creator_id": "user01",
			"name":       "pepe",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	e, err := c.CreateEmoji(context.Background(), "em01", DataCreateEmoji{
		Name:   "pepe",
		Parent: RawEmojiParent{Value: &ServerEmojiParent{ID: "srv01"}},
	})
	if err != nil {
		t.Fatalf("CreateEmoji: %v", err)
	}
	if e.Name != "pepe" {
		t.Errorf("Name = %q", e.Name)
	}
}

func TestDeleteEmoji(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/custom/emoji/em01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.DeleteEmoji(context.Background(), "em01"); err != nil {
		t.Fatalf("DeleteEmoji: %v", err)
	}
}

func TestFetchServerEmoji(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/emojis" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]map[string]any{
			{
				"_id":        "em01",
				"parent":     map[string]any{"type": "Server", "id": "srv01"},
				"creator_id": "user01",
				"name":       "pepe",
			},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	emojis, err := c.FetchServerEmoji(context.Background(), "srv01")
	if err != nil {
		t.Fatalf("FetchServerEmoji: %v", err)
	}
	if len(emojis) != 1 {
		t.Fatalf("len = %d", len(emojis))
	}
	if emojis[0].Name != "pepe" {
		t.Errorf("Name = %q", emojis[0].Name)
	}
}
