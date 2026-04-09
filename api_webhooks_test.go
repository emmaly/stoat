package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateWebhook(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/webhooks" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body CreateWebhookBody
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name != "MyHook" {
			t.Errorf("name = %q", body.Name)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Webhook{
			ID:          "wh01",
			Name:        "MyHook",
			CreatorID:   "user01",
			ChannelID:   "ch01",
			Permissions: 0,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	wh, err := c.CreateWebhook(context.Background(), "ch01", CreateWebhookBody{Name: "MyHook"})
	if err != nil {
		t.Fatalf("CreateWebhook: %v", err)
	}
	if wh.ID != "wh01" {
		t.Errorf("ID = %q", wh.ID)
	}
	if wh.Name != "MyHook" {
		t.Errorf("Name = %q", wh.Name)
	}
}

func TestFetchChannelWebhooks(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/channels/ch01/webhooks" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]Webhook{
			{ID: "wh01", Name: "Hook1", CreatorID: "user01", ChannelID: "ch01", Permissions: 0},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	whs, err := c.FetchChannelWebhooks(context.Background(), "ch01")
	if err != nil {
		t.Fatalf("FetchChannelWebhooks: %v", err)
	}
	if len(whs) != 1 {
		t.Fatalf("len = %d", len(whs))
	}
	if whs[0].ID != "wh01" {
		t.Errorf("ID = %q", whs[0].ID)
	}
}

func TestFetchWebhook(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/webhooks/wh01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		tok := "whtoken"
		json.NewEncoder(w).Encode(Webhook{
			ID:          "wh01",
			Name:        "Hook1",
			CreatorID:   "user01",
			ChannelID:   "ch01",
			Permissions: 0,
			Token:       &tok,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	wh, err := c.FetchWebhook(context.Background(), "wh01")
	if err != nil {
		t.Fatalf("FetchWebhook: %v", err)
	}
	if wh.Token == nil || *wh.Token != "whtoken" {
		t.Errorf("Token = %v", wh.Token)
	}
}

func TestFetchWebhookWithToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/webhooks/wh01/whtoken" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Webhook{
			ID:          "wh01",
			Name:        "Hook1",
			CreatorID:   "user01",
			ChannelID:   "ch01",
			Permissions: 0,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	wh, err := c.FetchWebhookWithToken(context.Background(), "wh01", "whtoken")
	if err != nil {
		t.Fatalf("FetchWebhookWithToken: %v", err)
	}
	if wh.ID != "wh01" {
		t.Errorf("ID = %q", wh.ID)
	}
}

func TestEditWebhook(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/webhooks/wh01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Webhook{
			ID:          "wh01",
			Name:        "Renamed",
			CreatorID:   "user01",
			ChannelID:   "ch01",
			Permissions: 0,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	name := "Renamed"
	wh, err := c.EditWebhook(context.Background(), "wh01", DataEditWebhook{Name: &name})
	if err != nil {
		t.Fatalf("EditWebhook: %v", err)
	}
	if wh.Name != "Renamed" {
		t.Errorf("Name = %q", wh.Name)
	}
}

func TestEditWebhookWithToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/webhooks/wh01/whtoken" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Webhook{
			ID:          "wh01",
			Name:        "Renamed",
			CreatorID:   "user01",
			ChannelID:   "ch01",
			Permissions: 0,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	name := "Renamed"
	wh, err := c.EditWebhookWithToken(context.Background(), "wh01", "whtoken", DataEditWebhook{Name: &name})
	if err != nil {
		t.Fatalf("EditWebhookWithToken: %v", err)
	}
	if wh.Name != "Renamed" {
		t.Errorf("Name = %q", wh.Name)
	}
}

func TestDeleteWebhook(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/webhooks/wh01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.DeleteWebhook(context.Background(), "wh01"); err != nil {
		t.Fatalf("DeleteWebhook: %v", err)
	}
}

func TestDeleteWebhookWithToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/webhooks/wh01/whtoken" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	if err := c.DeleteWebhookWithToken(context.Background(), "wh01", "whtoken"); err != nil {
		t.Fatalf("DeleteWebhookWithToken: %v", err)
	}
}

func TestExecuteWebhook(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/webhooks/wh01/whtoken" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body DataMessageSend
		json.NewDecoder(r.Body).Decode(&body)
		if body.Content == nil || *body.Content != "hello" {
			t.Errorf("content = %v", body.Content)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":     "msg01",
			"channel": "ch01",
			"author":  "wh01",
			"content": "hello",
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	content := "hello"
	msg, err := c.ExecuteWebhook(context.Background(), "wh01", "whtoken", DataMessageSend{Content: &content})
	if err != nil {
		t.Fatalf("ExecuteWebhook: %v", err)
	}
	if msg.ID != "msg01" {
		t.Errorf("ID = %q", msg.ID)
	}
}

func TestExecuteGitHubWebhook(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/webhooks/wh01/whtoken/github" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	payload := json.RawMessage(`{"action":"push"}`)
	if err := c.ExecuteGitHubWebhook(context.Background(), "wh01", "whtoken", payload); err != nil {
		t.Fatalf("ExecuteGitHubWebhook: %v", err)
	}
}
