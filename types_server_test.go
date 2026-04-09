package stoat

import (
	"encoding/json"
	"testing"
)

func TestServerUnmarshal(t *testing.T) {
	data := `{
		"_id": "srv01",
		"owner": "user01",
		"name": "Test Server",
		"description": "A test server",
		"channels": ["ch01", "ch02"],
		"categories": [{"id": "cat01", "title": "General", "channels": ["ch01"]}],
		"system_messages": {
			"user_joined": "ch01",
			"user_left": null,
			"user_kicked": null,
			"user_banned": null
		},
		"roles": {
			"role01": {
				"name": "Admin",
				"permissions": {"a": 255, "d": 0},
				"colour": "#ff0000",
				"hoist": true,
				"rank": 1
			}
		},
		"default_permissions": 1048576,
		"icon": {"_id": "icon01", "tag": "icons", "filename": "icon.png", "metadata": {"type": "Image", "width": 64, "height": 64}, "content_type": "image/png", "size": 1024},
		"banner": {"_id": "banner01", "tag": "banners", "filename": "banner.png", "metadata": {"type": "Image", "width": 800, "height": 200}, "content_type": "image/png", "size": 4096},
		"flags": 1,
		"nsfw": false,
		"analytics": true,
		"discoverable": true
	}`

	var s Server
	if err := json.Unmarshal([]byte(data), &s); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if s.ID != "srv01" {
		t.Errorf("ID = %q", s.ID)
	}
	if s.Owner != "user01" {
		t.Errorf("Owner = %q", s.Owner)
	}
	if s.Name != "Test Server" {
		t.Errorf("Name = %q", s.Name)
	}
	if s.Description == nil || *s.Description != "A test server" {
		t.Errorf("Description = %v", s.Description)
	}
	if len(s.Channels) != 2 {
		t.Errorf("Channels = %v", s.Channels)
	}
	if len(s.Categories) != 1 || s.Categories[0].ID != "cat01" {
		t.Errorf("Categories = %v", s.Categories)
	}
	if s.SystemMessages == nil || s.SystemMessages.UserJoined == nil || *s.SystemMessages.UserJoined != "ch01" {
		t.Errorf("SystemMessages.UserJoined = %v", s.SystemMessages)
	}
	if s.SystemMessages.UserLeft != nil {
		t.Errorf("SystemMessages.UserLeft should be nil")
	}
	role, ok := s.Roles["role01"]
	if !ok {
		t.Fatalf("role01 not found")
	}
	if role.Name != "Admin" {
		t.Errorf("Role.Name = %q", role.Name)
	}
	if role.Permissions.A != 255 {
		t.Errorf("Role.Permissions.A = %d", role.Permissions.A)
	}
	if role.Colour == nil || *role.Colour != "#ff0000" {
		t.Errorf("Role.Colour = %v", role.Colour)
	}
	if !role.Hoist {
		t.Error("Role.Hoist = false")
	}
	if role.Rank != 1 {
		t.Errorf("Role.Rank = %d", role.Rank)
	}
	if s.DefaultPermissions != 1048576 {
		t.Errorf("DefaultPermissions = %d", s.DefaultPermissions)
	}
	if s.Icon == nil || s.Icon.ID != "icon01" {
		t.Errorf("Icon = %v", s.Icon)
	}
	if s.Banner == nil || s.Banner.ID != "banner01" {
		t.Errorf("Banner = %v", s.Banner)
	}
	if s.Flags != 1 {
		t.Errorf("Flags = %d", s.Flags)
	}
	if s.NSFW {
		t.Error("NSFW = true")
	}
	if !s.Analytics {
		t.Error("Analytics = false")
	}
	if !s.Discoverable {
		t.Error("Discoverable = false")
	}
}

func TestServerMarshalRoundTrip(t *testing.T) {
	desc := "desc"
	colour := "#ff0000"
	s := Server{
		ID:    "srv01",
		Owner: "user01",
		Name:  "Test",
		Description: &desc,
		Channels: []string{"ch01"},
		Roles: map[string]Role{
			"r1": {Name: "Admin", Permissions: OverrideField{A: 255, D: 0}, Colour: &colour, Hoist: true, Rank: 1},
		},
		DefaultPermissions: 100,
	}
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var s2 Server
	if err := json.Unmarshal(b, &s2); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if s2.ID != s.ID || s2.Name != s.Name {
		t.Errorf("roundtrip mismatch: %+v", s2)
	}
}

func TestCreateServerLegacyResponseUnmarshal(t *testing.T) {
	data := `{
		"server": {
			"_id": "srv01",
			"owner": "user01",
			"name": "New Server",
			"channels": ["ch01"],
			"default_permissions": 0
		},
		"channels": [
			{"channel_type": "TextChannel", "_id": "ch01", "server": "srv01", "name": "general"}
		]
	}`
	var resp CreateServerLegacyResponse
	if err := json.Unmarshal([]byte(data), &resp); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if resp.Server.ID != "srv01" {
		t.Errorf("Server.ID = %q", resp.Server.ID)
	}
	if len(resp.Channels) != 1 {
		t.Fatalf("Channels len = %d", len(resp.Channels))
	}
	ch, ok := resp.Channels[0].Value.(*TextChannel)
	if !ok {
		t.Fatalf("channel type = %T", resp.Channels[0].Value)
	}
	if ch.Name != "general" {
		t.Errorf("channel name = %q", ch.Name)
	}
}

func TestFetchServerResponseUnmarshal(t *testing.T) {
	data := `{
		"server": {
			"_id": "srv01",
			"owner": "user01",
			"name": "My Server",
			"channels": ["ch01"],
			"default_permissions": 0
		},
		"channels": [
			{"channel_type": "TextChannel", "_id": "ch01", "server": "srv01", "name": "general"}
		]
	}`
	var resp FetchServerResponse
	if err := json.Unmarshal([]byte(data), &resp); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if resp.Server.ID != "srv01" {
		t.Errorf("Server.ID = %q", resp.Server.ID)
	}
	if len(resp.Channels) != 1 {
		t.Fatalf("Channels len = %d", len(resp.Channels))
	}
}

func TestSystemMessageChannelsNullable(t *testing.T) {
	data := `{"user_joined": "ch01", "user_left": null}`
	var sm SystemMessageChannels
	if err := json.Unmarshal([]byte(data), &sm); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if sm.UserJoined == nil || *sm.UserJoined != "ch01" {
		t.Errorf("UserJoined = %v", sm.UserJoined)
	}
	if sm.UserLeft != nil {
		t.Errorf("UserLeft = %v", sm.UserLeft)
	}
}

func TestFieldsServerValues(t *testing.T) {
	if FieldsServerIcon != "Icon" {
		t.Errorf("FieldsServerIcon = %q", FieldsServerIcon)
	}
	if FieldsServerBanner != "Banner" {
		t.Errorf("FieldsServerBanner = %q", FieldsServerBanner)
	}
	if FieldsServerDescription != "Description" {
		t.Errorf("FieldsServerDescription = %q", FieldsServerDescription)
	}
}

func TestDataCreateServerMarshal(t *testing.T) {
	desc := "my server"
	d := DataCreateServer{
		Name:        "Test",
		Description: &desc,
		NSFW:        nil,
	}
	b, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var m map[string]any
	json.Unmarshal(b, &m)
	if m["name"] != "Test" {
		t.Errorf("name = %v", m["name"])
	}
	if m["description"] != "my server" {
		t.Errorf("description = %v", m["description"])
	}
	if _, ok := m["nsfw"]; ok {
		t.Error("nsfw should be omitted")
	}
}

func TestDataEditServerMarshal(t *testing.T) {
	name := "Renamed"
	d := DataEditServer{
		Name:   &name,
		Remove: []FieldsServer{FieldsServerIcon},
	}
	b, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var m map[string]any
	json.Unmarshal(b, &m)
	if m["name"] != "Renamed" {
		t.Errorf("name = %v", m["name"])
	}
	r := m["remove"].([]any)
	if r[0] != "Icon" {
		t.Errorf("remove[0] = %v", r[0])
	}
}
