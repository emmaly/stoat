package stoat

import (
	"encoding/json"
	"testing"
)

func TestSavedMessagesChannelRoundTrip(t *testing.T) {
	orig := RawChannel{Value: &SavedMessagesChannel{ID: "ch01", User: "user01"}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawChannel
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	ch, ok := got.Value.(*SavedMessagesChannel)
	if !ok {
		t.Fatalf("type = %T, want *SavedMessagesChannel", got.Value)
	}
	if ch.ID != "ch01" || ch.User != "user01" {
		t.Errorf("got ID=%q User=%q", ch.ID, ch.User)
	}
	if ch.ChannelType() != "SavedMessages" {
		t.Errorf("ChannelType = %q", ch.ChannelType())
	}
	if ch.ChannelID() != "ch01" {
		t.Errorf("ChannelID = %q", ch.ChannelID())
	}
}

func TestDirectMessageChannelRoundTrip(t *testing.T) {
	lastMsg := "msg01"
	orig := RawChannel{Value: &DirectMessageChannel{
		ID:            "ch02",
		Active:        true,
		Recipients:    []string{"user01", "user02"},
		LastMessageID: &lastMsg,
	}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawChannel
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	ch, ok := got.Value.(*DirectMessageChannel)
	if !ok {
		t.Fatalf("type = %T, want *DirectMessageChannel", got.Value)
	}
	if ch.ID != "ch02" {
		t.Errorf("ID = %q", ch.ID)
	}
	if !ch.Active {
		t.Error("expected active = true")
	}
	if len(ch.Recipients) != 2 {
		t.Errorf("recipients len = %d", len(ch.Recipients))
	}
	if ch.LastMessageID == nil || *ch.LastMessageID != "msg01" {
		t.Errorf("last_message_id = %v", ch.LastMessageID)
	}
}

func TestGroupChannelRoundTrip(t *testing.T) {
	desc := "A test group"
	orig := RawChannel{Value: &GroupChannel{
		ID:          "ch03",
		Name:        "Test Group",
		Owner:       "user01",
		Description: &desc,
		Recipients:  []string{"user01", "user02", "user03"},
		NSFW:        false,
	}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawChannel
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	ch, ok := got.Value.(*GroupChannel)
	if !ok {
		t.Fatalf("type = %T, want *GroupChannel", got.Value)
	}
	if ch.Name != "Test Group" {
		t.Errorf("Name = %q", ch.Name)
	}
	if ch.Description == nil || *ch.Description != "A test group" {
		t.Errorf("Description = %v", ch.Description)
	}
}

func TestTextChannelRoundTrip(t *testing.T) {
	orig := RawChannel{Value: &TextChannel{
		ID:     "ch04",
		Server: "srv01",
		Name:   "general",
		RolePermissions: map[string]OverrideField{
			"role01": {A: 100, D: 50},
		},
	}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawChannel
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	ch, ok := got.Value.(*TextChannel)
	if !ok {
		t.Fatalf("type = %T, want *TextChannel", got.Value)
	}
	if ch.Server != "srv01" {
		t.Errorf("Server = %q", ch.Server)
	}
	if ch.Name != "general" {
		t.Errorf("Name = %q", ch.Name)
	}
	if rp, ok := ch.RolePermissions["role01"]; !ok || rp.A != 100 || rp.D != 50 {
		t.Errorf("RolePermissions = %+v", ch.RolePermissions)
	}
}

func TestVoiceChannelRoundTrip(t *testing.T) {
	desc := "Voice chat"
	orig := RawChannel{Value: &VoiceChannel{
		ID:          "ch05",
		Server:      "srv01",
		Name:        "voice-1",
		Description: &desc,
	}}
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var got RawChannel
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	ch, ok := got.Value.(*VoiceChannel)
	if !ok {
		t.Fatalf("type = %T, want *VoiceChannel", got.Value)
	}
	if ch.Name != "voice-1" {
		t.Errorf("Name = %q", ch.Name)
	}
	if ch.ChannelType() != "VoiceChannel" {
		t.Errorf("ChannelType = %q", ch.ChannelType())
	}
}
