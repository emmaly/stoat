package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchMembers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/members" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.URL.Query().Get("exclude_offline") != "true" {
			t.Errorf("exclude_offline = %q", r.URL.Query().Get("exclude_offline"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"members": []map[string]any{
				{"_id": map[string]any{"server": "srv01", "user": "user01"}, "joined_at": "2024-01-01T00:00:00Z"},
			},
			"users": []map[string]any{
				{"_id": "user01", "username": "alice", "discriminator": "0001", "relationship": "None", "online": false},
			},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	resp, err := c.FetchMembers(context.Background(), "srv01", true)
	if err != nil {
		t.Fatalf("FetchMembers: %v", err)
	}
	if len(resp.Members) != 1 || resp.Members[0].ID.User != "user01" {
		t.Errorf("Members = %+v", resp.Members)
	}
	if len(resp.Users) != 1 {
		t.Errorf("Users len = %d", len(resp.Users))
	}
}

func TestFetchMember(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/members/user01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.URL.Query().Get("roles") != "true" {
			t.Errorf("roles = %q", r.URL.Query().Get("roles"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"member": map[string]any{
				"_id":       map[string]any{"server": "srv01", "user": "user01"},
				"joined_at": "2024-01-01T00:00:00Z",
			},
			"roles": []map[string]any{
				{"name": "Admin", "permissions": map[string]any{"a": 255, "d": 0}},
			},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	resp, err := c.FetchMember(context.Background(), "srv01", "user01", true)
	if err != nil {
		t.Fatalf("FetchMember: %v", err)
	}
	if resp.Member.ID.User != "user01" {
		t.Errorf("Member.ID.User = %q", resp.Member.ID.User)
	}
	if len(resp.Roles) != 1 || resp.Roles[0].Name != "Admin" {
		t.Errorf("Roles = %+v", resp.Roles)
	}
}

func TestEditMember(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/members/user01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		nick := "NewNick"
		json.NewEncoder(w).Encode(Member{
			ID:       MemberCompositeKey{Server: "srv01", User: "user01"},
			JoinedAt: "2024-01-01T00:00:00Z",
			Nickname: &nick,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	nick := "NewNick"
	m, err := c.EditMember(context.Background(), "srv01", "user01", DataMemberEdit{Nickname: &nick})
	if err != nil {
		t.Fatalf("EditMember: %v", err)
	}
	if m.Nickname == nil || *m.Nickname != "NewNick" {
		t.Errorf("Nickname = %v", m.Nickname)
	}
}

func TestKickMember(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/members/user01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.KickMember(context.Background(), "srv01", "user01"); err != nil {
		t.Fatalf("KickMember: %v", err)
	}
}

func TestQueryMembers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/members_experimental_query" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.URL.Query().Get("query") != "alice" {
			t.Errorf("query = %q", r.URL.Query().Get("query"))
		}
		if r.URL.Query().Get("experimental_api") != "true" {
			t.Errorf("experimental_api = %q", r.URL.Query().Get("experimental_api"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"members": []map[string]any{},
			"users":   []map[string]any{},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	resp, err := c.QueryMembers(context.Background(), "srv01", "alice")
	if err != nil {
		t.Fatalf("QueryMembers: %v", err)
	}
	if resp == nil {
		t.Fatal("resp is nil")
	}
}

func TestBanUser(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/bans/user01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		reason := "spam"
		json.NewEncoder(w).Encode(ServerBan{
			ID:     MemberCompositeKey{Server: "srv01", User: "user01"},
			Reason: &reason,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	reason := "spam"
	ban, err := c.BanUser(context.Background(), "srv01", "user01", DataBanCreate{Reason: &reason})
	if err != nil {
		t.Fatalf("BanUser: %v", err)
	}
	if ban.Reason == nil || *ban.Reason != "spam" {
		t.Errorf("Reason = %v", ban.Reason)
	}
}

func TestUnbanUser(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/bans/user01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.UnbanUser(context.Background(), "srv01", "user01"); err != nil {
		t.Fatalf("UnbanUser: %v", err)
	}
}

func TestFetchBans(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/bans" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"users": []map[string]any{
				{"_id": "user01", "username": "alice", "discriminator": "0001"},
			},
			"bans": []map[string]any{
				{"_id": map[string]any{"server": "srv01", "user": "user01"}, "reason": "spam"},
			},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	resp, err := c.FetchBans(context.Background(), "srv01")
	if err != nil {
		t.Fatalf("FetchBans: %v", err)
	}
	if len(resp.Users) != 1 {
		t.Errorf("Users len = %d", len(resp.Users))
	}
	if len(resp.Bans) != 1 {
		t.Errorf("Bans len = %d", len(resp.Bans))
	}
}

func TestFetchServerInvites(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/invites" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]map[string]any{
			{"_id": "inv01", "server": "srv01", "creator": "user01", "channel": "ch01"},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	invites, err := c.FetchServerInvites(context.Background(), "srv01")
	if err != nil {
		t.Fatalf("FetchServerInvites: %v", err)
	}
	if len(invites) != 1 {
		t.Errorf("invites len = %d", len(invites))
	}
}
