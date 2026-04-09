package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateRole(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/roles" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body DataCreateRole
		json.NewDecoder(r.Body).Decode(&body)
		if body.Name != "Moderator" {
			t.Errorf("name = %q", body.Name)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"id":   "role01",
			"role": map[string]any{"name": "Moderator", "permissions": map[string]any{"a": 0, "d": 0}},
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	resp, err := c.CreateRole(context.Background(), "srv01", DataCreateRole{Name: "Moderator"})
	if err != nil {
		t.Fatalf("CreateRole: %v", err)
	}
	if resp.ID != "role01" {
		t.Errorf("ID = %q", resp.ID)
	}
	if resp.Role.Name != "Moderator" {
		t.Errorf("Role.Name = %q", resp.Role.Name)
	}
}

func TestFetchRole(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/roles/role01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"name":        "Admin",
			"permissions": map[string]any{"a": 255, "d": 0},
			"hoist":       true,
			"rank":        0,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	role, err := c.FetchRole(context.Background(), "srv01", "role01")
	if err != nil {
		t.Fatalf("FetchRole: %v", err)
	}
	if role.Name != "Admin" {
		t.Errorf("Name = %q", role.Name)
	}
}

func TestEditRole(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/roles/role01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		colour := "#00ff00"
		json.NewEncoder(w).Encode(Role{
			Name:        "Renamed",
			Permissions: OverrideField{A: 255, D: 0},
			Colour:      &colour,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	name := "Renamed"
	colour := "#00ff00"
	role, err := c.EditRole(context.Background(), "srv01", "role01", DataEditRole{Name: &name, Colour: &colour})
	if err != nil {
		t.Fatalf("EditRole: %v", err)
	}
	if role.Name != "Renamed" {
		t.Errorf("Name = %q", role.Name)
	}
	if role.Colour == nil || *role.Colour != "#00ff00" {
		t.Errorf("Colour = %v", role.Colour)
	}
}

func TestDeleteRole(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/roles/role01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.DeleteRole(context.Background(), "srv01", "role01"); err != nil {
		t.Fatalf("DeleteRole: %v", err)
	}
}

func TestEditRoleRanks(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/roles/ranks" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":                 "srv01",
			"owner":              "user01",
			"name":               "My Server",
			"channels":           []string{},
			"default_permissions": 0,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	s, err := c.EditRoleRanks(context.Background(), "srv01", DataEditRoleRanks{
		Roles: map[string]int64{"role01": 0, "role02": 1},
	})
	if err != nil {
		t.Fatalf("EditRoleRanks: %v", err)
	}
	if s.ID != "srv01" {
		t.Errorf("Server.ID = %q", s.ID)
	}
}

func TestSetDefaultServerPermissions(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/permissions/default" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":                 "srv01",
			"owner":              "user01",
			"name":               "My Server",
			"channels":           []string{},
			"default_permissions": 1048576,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	s, err := c.SetDefaultServerPermissions(context.Background(), "srv01", DataPermissionsValue{Permissions: 1048576})
	if err != nil {
		t.Fatalf("SetDefaultServerPermissions: %v", err)
	}
	if s.DefaultPermissions != 1048576 {
		t.Errorf("DefaultPermissions = %d", s.DefaultPermissions)
	}
}

func TestSetRoleServerPermission(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/servers/srv01/permissions/role01" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"_id":                 "srv01",
			"owner":              "user01",
			"name":               "My Server",
			"channels":           []string{},
			"default_permissions": 0,
		})
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	s, err := c.SetRoleServerPermission(context.Background(), "srv01", "role01", DataSetServerRolePermission{
		Permissions: Override{Allow: 255, Deny: 10},
	})
	if err != nil {
		t.Fatalf("SetRoleServerPermission: %v", err)
	}
	if s.ID != "srv01" {
		t.Errorf("Server.ID = %q", s.ID)
	}
}
