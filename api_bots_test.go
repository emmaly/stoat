package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBot(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/bots/create" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body DataCreateBot
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name != "TestBot" {
			t.Errorf("name = %q", body.Name)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":    "bot01",
			"owner":  "user01",
			"token":  "tok123",
			"public": false,
			"user": map[string]any{
				"_id":           "bot01",
				"username":      "TestBot",
				"discriminator": "0000",
				"relationship":  "None",
				"online":        false,
			},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	resp, err := c.CreateBot(context.Background(), DataCreateBot{Name: "TestBot"})
	if err != nil {
		t.Fatalf("CreateBot: %v", err)
	}
	if resp.Bot.ID != "bot01" {
		t.Errorf("Bot.ID = %q", resp.Bot.ID)
	}
	if resp.User.Username != "TestBot" {
		t.Errorf("User.Username = %q", resp.User.Username)
	}
}

func TestFetchBot(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/bots/bot01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"bot": map[string]any{
				"_id":    "bot01",
				"owner":  "user01",
				"token":  "tok123",
				"public": true,
			},
			"user": map[string]any{
				"_id":           "bot01",
				"username":      "TestBot",
				"discriminator": "0000",
				"relationship":  "None",
				"online":        false,
			},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	resp, err := c.FetchBot(context.Background(), "bot01")
	if err != nil {
		t.Fatalf("FetchBot: %v", err)
	}
	if resp.Bot.ID != "bot01" {
		t.Errorf("Bot.ID = %q", resp.Bot.ID)
	}
	if !resp.Bot.Public {
		t.Error("Bot.Public = false")
	}
	if resp.User.Username != "TestBot" {
		t.Errorf("User.Username = %q", resp.User.Username)
	}
}

func TestEditBot(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/bots/bot01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body DataEditBot
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name == nil || *body.Name != "Renamed" {
			t.Errorf("name = %v", body.Name)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":    "bot01",
			"owner":  "user01",
			"token":  "tok123",
			"public": false,
			"user": map[string]any{
				"_id":           "bot01",
				"username":      "Renamed",
				"discriminator": "0000",
				"relationship":  "None",
				"online":        false,
			},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	name := "Renamed"
	resp, err := c.EditBot(context.Background(), "bot01", DataEditBot{Name: &name})
	if err != nil {
		t.Fatalf("EditBot: %v", err)
	}
	if resp.User.Username != "Renamed" {
		t.Errorf("User.Username = %q", resp.User.Username)
	}
}

func TestDeleteBot(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/bots/bot01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.DeleteBot(context.Background(), "bot01"); err != nil {
		t.Fatalf("DeleteBot: %v", err)
	}
}

func TestFetchOwnedBots(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/bots/@me" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"bots": []map[string]any{
				{"_id": "bot01", "owner": "user01", "token": "tok123", "public": false},
			},
			"users": []map[string]any{
				{"_id": "bot01", "username": "TestBot", "discriminator": "0000", "relationship": "None", "online": false},
			},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	resp, err := c.FetchOwnedBots(context.Background())
	if err != nil {
		t.Fatalf("FetchOwnedBots: %v", err)
	}
	if len(resp.Bots) != 1 {
		t.Fatalf("Bots len = %d", len(resp.Bots))
	}
	if resp.Bots[0].ID != "bot01" {
		t.Errorf("Bots[0].ID = %q", resp.Bots[0].ID)
	}
	if len(resp.Users) != 1 {
		t.Fatalf("Users len = %d", len(resp.Users))
	}
}

func TestFetchPublicBot(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/bots/bot01/invite" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		desc := "A test bot"
		json.NewEncoder(w).Encode(PublicBot{
			ID:          "bot01",
			Username:    "TestBot",
			Description: &desc,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	pb, err := c.FetchPublicBot(context.Background(), "bot01")
	if err != nil {
		t.Fatalf("FetchPublicBot: %v", err)
	}
	if pb.Username != "TestBot" {
		t.Errorf("Username = %q", pb.Username)
	}
	if pb.Description == nil || *pb.Description != "A test bot" {
		t.Errorf("Description = %v", pb.Description)
	}
}

func TestInviteBot(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/bots/bot01/invite" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body InviteBotDestination
		json.NewDecoder(r.Body).Decode(&body)
		if body.Server == nil || *body.Server != "srv01" {
			t.Errorf("Server = %v", body.Server)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	srvID := "srv01"
	if err := c.InviteBot(context.Background(), "bot01", InviteBotDestination{Server: &srvID}); err != nil {
		t.Fatalf("InviteBot: %v", err)
	}
}
