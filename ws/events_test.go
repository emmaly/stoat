package ws

import (
	"encoding/json"
	"testing"

	"github.com/emmaly/stoat"
)

func TestPingMarshalJSON(t *testing.T) {
	ev := PingEvent{Data: 1234567890}
	data, err := json.Marshal(ev)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	if m["type"] != "Ping" {
		t.Errorf("expected type Ping, got %v", m["type"])
	}
	if int64(m["data"].(float64)) != 1234567890 {
		t.Errorf("expected data 1234567890, got %v", m["data"])
	}
}

func TestPongUnmarshalJSON(t *testing.T) {
	raw := `{"type":"Pong","data":1234567890}`
	var re RawEvent
	if err := json.Unmarshal([]byte(raw), &re); err != nil {
		t.Fatal(err)
	}
	ev, ok := re.Value.(*PongEvent)
	if !ok {
		t.Fatalf("expected *PongEvent, got %T", re.Value)
	}
	if ev.Data != 1234567890 {
		t.Errorf("expected data 1234567890, got %d", ev.Data)
	}
}

func TestErrorUnmarshalJSON(t *testing.T) {
	raw := `{"type":"Error","error":"InvalidSession"}`
	var re RawEvent
	if err := json.Unmarshal([]byte(raw), &re); err != nil {
		t.Fatal(err)
	}
	ev, ok := re.Value.(*ErrorEvent)
	if !ok {
		t.Fatalf("expected *ErrorEvent, got %T", re.Value)
	}
	if ev.Error != "InvalidSession" {
		t.Errorf("expected error InvalidSession, got %s", ev.Error)
	}
}

func TestReadyUnmarshalJSON(t *testing.T) {
	raw := `{
		"type": "Ready",
		"users": [{"_id":"u1","username":"alice","discriminator":"0001","relationship":"Friend","online":true}],
		"servers": [{"_id":"s1","owner":"u1","name":"Test","channels":[],"default_permissions":0}],
		"channels": [{"channel_type":"TextChannel","_id":"c1","server":"s1","name":"general"}],
		"members": [{"_id":{"server":"s1","user":"u1"},"joined_at":"2024-01-01T00:00:00Z"}],
		"emojis": [],
		"user_settings": {"theme": "dark"},
		"channel_unreads": []
	}`
	var re RawEvent
	if err := json.Unmarshal([]byte(raw), &re); err != nil {
		t.Fatal(err)
	}
	ev, ok := re.Value.(*ReadyEvent)
	if !ok {
		t.Fatalf("expected *ReadyEvent, got %T", re.Value)
	}
	if len(ev.Users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(ev.Users))
	}
	if ev.Users[0].ID != "u1" {
		t.Errorf("expected user id u1, got %s", ev.Users[0].ID)
	}
	if len(ev.Servers) != 1 {
		t.Fatalf("expected 1 server, got %d", len(ev.Servers))
	}
	if ev.Servers[0].Name != "Test" {
		t.Errorf("expected server name Test, got %s", ev.Servers[0].Name)
	}
	if len(ev.Channels) != 1 {
		t.Fatalf("expected 1 channel, got %d", len(ev.Channels))
	}
	if ev.Channels[0].Value == nil {
		t.Fatal("expected non-nil channel value")
	}
	if ev.UserSettings == nil {
		t.Fatal("expected non-nil user_settings")
	}
	if _, ok := ev.UserSettings["theme"]; !ok {
		t.Error("expected theme key in user_settings")
	}
}

func TestMessageEventUnmarshalJSON(t *testing.T) {
	raw := `{
		"type": "Message",
		"_id": "msg1",
		"channel": "c1",
		"author": "u1",
		"content": "Hello, world!"
	}`
	var re RawEvent
	if err := json.Unmarshal([]byte(raw), &re); err != nil {
		t.Fatal(err)
	}
	ev, ok := re.Value.(*MessageEvent)
	if !ok {
		t.Fatalf("expected *MessageEvent, got %T", re.Value)
	}
	if ev.ID != "msg1" {
		t.Errorf("expected id msg1, got %s", ev.ID)
	}
	if ev.Content == nil || *ev.Content != "Hello, world!" {
		t.Errorf("expected content 'Hello, world!', got %v", ev.Content)
	}
}

func TestChannelCreateEventUnmarshalJSON(t *testing.T) {
	raw := `{
		"type": "ChannelCreate",
		"channel_type": "TextChannel",
		"_id": "c2",
		"server": "s1",
		"name": "new-channel"
	}`
	var re RawEvent
	if err := json.Unmarshal([]byte(raw), &re); err != nil {
		t.Fatal(err)
	}
	ev, ok := re.Value.(*ChannelCreateEvent)
	if !ok {
		t.Fatalf("expected *ChannelCreateEvent, got %T", re.Value)
	}
	ch := ev.Channel()
	if ch == nil {
		t.Fatal("expected non-nil channel")
	}
	tc, ok := ch.(*stoat.TextChannel)
	if !ok {
		t.Fatalf("expected *stoat.TextChannel, got %T", ch)
	}
	if tc.Name != "new-channel" {
		t.Errorf("expected name new-channel, got %s", tc.Name)
	}
}

func TestServerUpdateEventUnmarshalJSON(t *testing.T) {
	raw := `{
		"type": "ServerUpdate",
		"id": "s1",
		"data": {"name": "Updated Server"},
		"clear": ["Description"]
	}`
	var re RawEvent
	if err := json.Unmarshal([]byte(raw), &re); err != nil {
		t.Fatal(err)
	}
	ev, ok := re.Value.(*ServerUpdateEvent)
	if !ok {
		t.Fatalf("expected *ServerUpdateEvent, got %T", re.Value)
	}
	if ev.ID != "s1" {
		t.Errorf("expected id s1, got %s", ev.ID)
	}
	if len(ev.Clear) != 1 || ev.Clear[0] != "Description" {
		t.Errorf("expected clear [Description], got %v", ev.Clear)
	}
}

func TestBulkEventUnmarshalJSON(t *testing.T) {
	raw := `{
		"type": "Bulk",
		"v": [
			{"type": "Authenticated"},
			{"type": "Pong", "data": 42}
		]
	}`
	var re RawEvent
	if err := json.Unmarshal([]byte(raw), &re); err != nil {
		t.Fatal(err)
	}
	ev, ok := re.Value.(*BulkEvent)
	if !ok {
		t.Fatalf("expected *BulkEvent, got %T", re.Value)
	}
	if len(ev.V) != 2 {
		t.Fatalf("expected 2 events, got %d", len(ev.V))
	}
	if _, ok := ev.V[0].Value.(*AuthenticatedEvent); !ok {
		t.Errorf("expected *AuthenticatedEvent, got %T", ev.V[0].Value)
	}
	pong, ok := ev.V[1].Value.(*PongEvent)
	if !ok {
		t.Fatalf("expected *PongEvent, got %T", ev.V[1].Value)
	}
	if pong.Data != 42 {
		t.Errorf("expected data 42, got %d", pong.Data)
	}
}

func TestAuthenticateMarshalJSON(t *testing.T) {
	ev := AuthenticateEvent{Token: "my-token"}
	data, err := json.Marshal(ev)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	if m["type"] != "Authenticate" {
		t.Errorf("expected type Authenticate, got %v", m["type"])
	}
	if m["token"] != "my-token" {
		t.Errorf("expected token my-token, got %v", m["token"])
	}
}

func TestBeginTypingMarshalJSON(t *testing.T) {
	ev := BeginTypingEvent{Channel: "c1"}
	data, err := json.Marshal(ev)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	if m["type"] != "BeginTyping" {
		t.Errorf("expected type BeginTyping, got %v", m["type"])
	}
	if m["channel"] != "c1" {
		t.Errorf("expected channel c1, got %v", m["channel"])
	}
}

func TestSubscribeMarshalJSON(t *testing.T) {
	ev := SubscribeEvent{ServerID: "s1"}
	data, err := json.Marshal(ev)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	if m["type"] != "Subscribe" {
		t.Errorf("expected type Subscribe, got %v", m["type"])
	}
	if m["server_id"] != "s1" {
		t.Errorf("expected server_id s1, got %v", m["server_id"])
	}
}

func TestAllEventTypesUnmarshal(t *testing.T) {
	// Test that all event types can be unmarshalled without error
	tests := []struct {
		name string
		json string
		want string // expected Go type suffix
	}{
		{"Authenticated", `{"type":"Authenticated"}`, "*ws.AuthenticatedEvent"},
		{"Logout", `{"type":"Logout"}`, "*ws.LogoutEvent"},
		{"MessageDelete", `{"type":"MessageDelete","id":"m1","channel":"c1"}`, "*ws.MessageDeleteEvent"},
		{"MessageReact", `{"type":"MessageReact","id":"m1","channel_id":"c1","user_id":"u1","emoji_id":"e1"}`, "*ws.MessageReactEvent"},
		{"MessageUnreact", `{"type":"MessageUnreact","id":"m1","channel_id":"c1","user_id":"u1","emoji_id":"e1"}`, "*ws.MessageUnreactEvent"},
		{"MessageRemoveReaction", `{"type":"MessageRemoveReaction","id":"m1","channel_id":"c1","emoji_id":"e1"}`, "*ws.MessageRemoveReactionEvent"},
		{"ChannelUpdate", `{"type":"ChannelUpdate","id":"c1","data":{},"clear":[]}`, "*ws.ChannelUpdateEvent"},
		{"ChannelDelete", `{"type":"ChannelDelete","id":"c1"}`, "*ws.ChannelDeleteEvent"},
		{"ChannelGroupJoin", `{"type":"ChannelGroupJoin","id":"c1","user":"u1"}`, "*ws.ChannelGroupJoinEvent"},
		{"ChannelGroupLeave", `{"type":"ChannelGroupLeave","id":"c1","user":"u1"}`, "*ws.ChannelGroupLeaveEvent"},
		{"ChannelStartTyping", `{"type":"ChannelStartTyping","id":"c1","user":"u1"}`, "*ws.ChannelStartTypingEvent"},
		{"ChannelStopTyping", `{"type":"ChannelStopTyping","id":"c1","user":"u1"}`, "*ws.ChannelStopTypingEvent"},
		{"ChannelAck", `{"type":"ChannelAck","id":"c1","user":"u1","message_id":"m1"}`, "*ws.ChannelAckEvent"},
		{"ServerDelete", `{"type":"ServerDelete","id":"s1"}`, "*ws.ServerDeleteEvent"},
		{"ServerMemberUpdate", `{"type":"ServerMemberUpdate","id":{"server":"s1","user":"u1"},"data":{},"clear":[]}`, "*ws.ServerMemberUpdateEvent"},
		{"ServerMemberJoin", `{"type":"ServerMemberJoin","id":"s1","user":"u1","member":{"_id":{"server":"s1","user":"u1"},"joined_at":"2024-01-01T00:00:00Z"}}`, "*ws.ServerMemberJoinEvent"},
		{"ServerMemberLeave", `{"type":"ServerMemberLeave","id":"s1","user":"u1"}`, "*ws.ServerMemberLeaveEvent"},
		{"ServerRoleUpdate", `{"type":"ServerRoleUpdate","id":"s1","role_id":"r1","data":{},"clear":[]}`, "*ws.ServerRoleUpdateEvent"},
		{"ServerRoleDelete", `{"type":"ServerRoleDelete","id":"s1","role_id":"r1"}`, "*ws.ServerRoleDeleteEvent"},
		{"UserUpdate", `{"type":"UserUpdate","id":"u1","data":{},"clear":[]}`, "*ws.UserUpdateEvent"},
		{"UserRelationship", `{"type":"UserRelationship","id":"u1","user":{"_id":"u2","username":"bob","discriminator":"0002","relationship":"Friend","online":true},"status":"Friend"}`, "*ws.UserRelationshipEvent"},
		{"UserPlatformWipe", `{"type":"UserPlatformWipe","user_id":"u1","flags":1}`, "*ws.UserPlatformWipeEvent"},
		{"EmojiCreate", `{"type":"EmojiCreate","_id":"e1","parent":{"type":"Server","id":"s1"},"creator_id":"u1","name":"test"}`, "*ws.EmojiCreateEvent"},
		{"EmojiDelete", `{"type":"EmojiDelete","id":"e1"}`, "*ws.EmojiDeleteEvent"},
		{"Auth", `{"type":"Auth","event_type":"DeleteSession","user_id":"u1","session_id":"sess1"}`, "*ws.AuthEvent"},
		{"MessageUpdate", `{"type":"MessageUpdate","id":"m1","channel":"c1","data":{"content":"edited"}}`, "*ws.MessageUpdateEvent"},
		{"MessageAppend", `{"type":"MessageAppend","id":"m1","channel":"c1","append":{"embeds":[]}}`, "*ws.MessageAppendEvent"},
		{"ServerCreate", `{"type":"ServerCreate","_id":"s1","owner":"u1","name":"New","channels":[],"default_permissions":0}`, "*ws.ServerCreateEvent"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var re RawEvent
			if err := json.Unmarshal([]byte(tt.json), &re); err != nil {
				t.Fatalf("unmarshal failed: %v", err)
			}
			if re.Value == nil {
				t.Fatal("expected non-nil event")
			}
			if re.Value.EventType() != tt.name {
				t.Errorf("expected EventType() %q, got %q", tt.name, re.Value.EventType())
			}
		})
	}
}

func TestUnknownEventType(t *testing.T) {
	raw := `{"type":"UnknownFutureEvent","foo":"bar"}`
	var re RawEvent
	err := json.Unmarshal([]byte(raw), &re)
	if err == nil {
		t.Fatal("expected error for unknown event type")
	}
}
