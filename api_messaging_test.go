package stoat

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendMessage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/messages" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.Header.Get("Idempotency-Key") != "idem-123" {
			t.Errorf("idempotency key = %q", r.Header.Get("Idempotency-Key"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":     "msg01",
			"channel": "ch01",
			"author":  "user01",
			"content": "hello",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	content := "hello"
	msg, err := c.SendMessage(context.Background(), "ch01", DataMessageSend{
		Content:        &content,
		IdempotencyKey: "idem-123",
	})
	if err != nil {
		t.Fatalf("SendMessage: %v", err)
	}
	if msg.ID != "msg01" {
		t.Errorf("ID = %q", msg.ID)
	}
	if msg.Content == nil || *msg.Content != "hello" {
		t.Errorf("Content = %v", msg.Content)
	}
}

func TestFetchMessages(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/messages" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.URL.Query().Get("limit") != "10" {
			t.Errorf("limit = %q", r.URL.Query().Get("limit"))
		}
		if r.URL.Query().Get("include_users") != "true" {
			t.Errorf("include_users = %q", r.URL.Query().Get("include_users"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"messages": []map[string]any{
				{"_id": "msg01", "channel": "ch01", "author": "user01", "content": "hi"},
			},
			"users": []map[string]any{
				{"_id": "user01", "username": "alice", "discriminator": "0001", "relationship": "User", "online": true},
			},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	limit := 10
	includeUsers := true
	resp, err := c.FetchMessages(context.Background(), "ch01", &FetchMessagesOptions{
		Limit:        &limit,
		IncludeUsers: &includeUsers,
	})
	if err != nil {
		t.Fatalf("FetchMessages: %v", err)
	}
	if len(resp.Messages) != 1 {
		t.Errorf("messages len = %d", len(resp.Messages))
	}
	if len(resp.Users) != 1 {
		t.Errorf("users len = %d", len(resp.Users))
	}
}

func TestFetchMessage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/channels/ch01/messages/msg01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":     "msg01",
			"channel": "ch01",
			"author":  "user01",
			"content": "hello",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	msg, err := c.FetchMessage(context.Background(), "ch01", "msg01")
	if err != nil {
		t.Fatalf("FetchMessage: %v", err)
	}
	if msg.ID != "msg01" {
		t.Errorf("ID = %q", msg.ID)
	}
}

func TestEditMessage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/messages/msg01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":     "msg01",
			"channel": "ch01",
			"author":  "user01",
			"content": "edited",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	content := "edited"
	msg, err := c.EditMessage(context.Background(), "ch01", "msg01", DataEditMessage{Content: &content})
	if err != nil {
		t.Fatalf("EditMessage: %v", err)
	}
	if msg.Content == nil || *msg.Content != "edited" {
		t.Errorf("Content = %v", msg.Content)
	}
}

func TestDeleteMessage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/messages/msg01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.DeleteMessage(context.Background(), "ch01", "msg01"); err != nil {
		t.Fatalf("DeleteMessage: %v", err)
	}
}

func TestBulkDeleteMessages(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/messages/bulk" {
			t.Errorf("path = %q", r.URL.Path)
		}
		body, _ := io.ReadAll(r.Body)
		var data OptionsBulkDelete
		json.Unmarshal(body, &data)
		if len(data.IDs) != 2 {
			t.Errorf("ids len = %d", len(data.IDs))
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.BulkDeleteMessages(context.Background(), "ch01", OptionsBulkDelete{IDs: []string{"msg01", "msg02"}}); err != nil {
		t.Fatalf("BulkDeleteMessages: %v", err)
	}
}

func TestSearchMessages(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/search" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"messages": []map[string]any{
				{"_id": "msg01", "channel": "ch01", "author": "user01", "content": "found"},
			},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	query := "found"
	resp, err := c.SearchMessages(context.Background(), "ch01", DataMessageSearch{Query: &query})
	if err != nil {
		t.Fatalf("SearchMessages: %v", err)
	}
	if len(resp.Messages) != 1 {
		t.Errorf("messages len = %d", len(resp.Messages))
	}
}

func TestPinMessage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/messages/msg01/pin" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.PinMessage(context.Background(), "ch01", "msg01"); err != nil {
		t.Fatalf("PinMessage: %v", err)
	}
}

func TestUnpinMessage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/messages/msg01/pin" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.UnpinMessage(context.Background(), "ch01", "msg01"); err != nil {
		t.Fatalf("UnpinMessage: %v", err)
	}
}

func TestAcknowledgeMessage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/ack/msg01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.AcknowledgeMessage(context.Background(), "ch01", "msg01"); err != nil {
		t.Fatalf("AcknowledgeMessage: %v", err)
	}
}
