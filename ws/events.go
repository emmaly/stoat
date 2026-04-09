// Package ws provides WebSocket client support for the Stoat chat API.
package ws

import (
	"encoding/json"
	"fmt"

	"github.com/emmaly/stoat"
)

// Event is the interface implemented by all WebSocket events.
type Event interface {
	eventMarker()
	EventType() string
}

// --- Client-to-server events ---

// AuthenticateEvent authenticates a WebSocket connection.
type AuthenticateEvent struct {
	Token string `json:"token"`
}

func (AuthenticateEvent) eventMarker()        {}
func (AuthenticateEvent) EventType() string   { return "Authenticate" }
func (e AuthenticateEvent) MarshalJSON() ([]byte, error) {
	type Alias AuthenticateEvent
	return json.Marshal(struct {
		Type string `json:"type"`
		Alias
	}{Type: "Authenticate", Alias: Alias(e)})
}

// PingEvent sends a keepalive ping to the server.
type PingEvent struct {
	Data int64 `json:"data"`
}

func (PingEvent) eventMarker()        {}
func (PingEvent) EventType() string   { return "Ping" }
func (e PingEvent) MarshalJSON() ([]byte, error) {
	type Alias PingEvent
	return json.Marshal(struct {
		Type string `json:"type"`
		Alias
	}{Type: "Ping", Alias: Alias(e)})
}

// BeginTypingEvent notifies the server that the user started typing.
type BeginTypingEvent struct {
	Channel string `json:"channel"`
}

func (BeginTypingEvent) eventMarker()        {}
func (BeginTypingEvent) EventType() string   { return "BeginTyping" }
func (e BeginTypingEvent) MarshalJSON() ([]byte, error) {
	type Alias BeginTypingEvent
	return json.Marshal(struct {
		Type string `json:"type"`
		Alias
	}{Type: "BeginTyping", Alias: Alias(e)})
}

// EndTypingEvent notifies the server that the user stopped typing.
type EndTypingEvent struct {
	Channel string `json:"channel"`
}

func (EndTypingEvent) eventMarker()        {}
func (EndTypingEvent) EventType() string   { return "EndTyping" }
func (e EndTypingEvent) MarshalJSON() ([]byte, error) {
	type Alias EndTypingEvent
	return json.Marshal(struct {
		Type string `json:"type"`
		Alias
	}{Type: "EndTyping", Alias: Alias(e)})
}

// SubscribeEvent subscribes to a server's UserUpdate events.
type SubscribeEvent struct {
	ServerID string `json:"server_id"`
}

func (SubscribeEvent) eventMarker()        {}
func (SubscribeEvent) EventType() string   { return "Subscribe" }
func (e SubscribeEvent) MarshalJSON() ([]byte, error) {
	type Alias SubscribeEvent
	return json.Marshal(struct {
		Type string `json:"type"`
		Alias
	}{Type: "Subscribe", Alias: Alias(e)})
}

// --- Server-to-client events ---

// AuthenticatedEvent confirms successful authentication.
type AuthenticatedEvent struct{}

func (AuthenticatedEvent) eventMarker()      {}
func (AuthenticatedEvent) EventType() string { return "Authenticated" }

// ReadyEvent is the initial state dump after authentication.
type ReadyEvent struct {
	Users          []stoat.User                   `json:"users,omitempty"`
	Servers        []stoat.Server                 `json:"servers,omitempty"`
	Channels       []stoat.RawChannel             `json:"channels,omitempty"`
	Members        []stoat.Member                 `json:"members,omitempty"`
	Emojis         []stoat.Emoji                  `json:"emojis,omitempty"`
	UserSettings   map[string]json.RawMessage     `json:"user_settings,omitempty"`
	ChannelUnreads []stoat.ChannelUnread           `json:"channel_unreads,omitempty"`
}

func (ReadyEvent) eventMarker()      {}
func (ReadyEvent) EventType() string { return "Ready" }

// ErrorEvent indicates an error from the server.
type ErrorEvent struct {
	Error string `json:"error"`
}

func (ErrorEvent) eventMarker()      {}
func (ErrorEvent) EventType() string { return "Error" }

// LogoutEvent indicates the session has been invalidated.
type LogoutEvent struct{}

func (LogoutEvent) eventMarker()      {}
func (LogoutEvent) EventType() string { return "Logout" }

// PongEvent is the server's response to a PingEvent.
type PongEvent struct {
	Data int64 `json:"data"`
}

func (PongEvent) eventMarker()      {}
func (PongEvent) EventType() string { return "Pong" }

// BulkEvent contains multiple events bundled together.
type BulkEvent struct {
	V []RawEvent `json:"v"`
}

func (BulkEvent) eventMarker()      {}
func (BulkEvent) EventType() string { return "Bulk" }

// --- Message events ---

// MessageEvent represents a new message.
type MessageEvent struct {
	stoat.Message
}

func (MessageEvent) eventMarker()      {}
func (MessageEvent) EventType() string { return "Message" }

// MessageUpdateEvent indicates a message was edited.
type MessageUpdateEvent struct {
	ID      string          `json:"id"`
	Channel string          `json:"channel"`
	Data    json.RawMessage `json:"data"`
}

func (MessageUpdateEvent) eventMarker()      {}
func (MessageUpdateEvent) EventType() string { return "MessageUpdate" }

// MessageAppendEvent indicates data was appended to a message.
type MessageAppendEvent struct {
	ID      string             `json:"id"`
	Channel string             `json:"channel"`
	Append  MessageAppendBody  `json:"append"`
}

// MessageAppendBody contains the appended data.
type MessageAppendBody struct {
	Embeds []stoat.RawEmbed `json:"embeds,omitempty"`
}

func (MessageAppendEvent) eventMarker()      {}
func (MessageAppendEvent) EventType() string { return "MessageAppend" }

// MessageDeleteEvent indicates a message was deleted.
type MessageDeleteEvent struct {
	ID      string `json:"id"`
	Channel string `json:"channel"`
}

func (MessageDeleteEvent) eventMarker()      {}
func (MessageDeleteEvent) EventType() string { return "MessageDelete" }

// MessageReactEvent indicates a reaction was added.
type MessageReactEvent struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel_id"`
	UserID    string `json:"user_id"`
	EmojiID   string `json:"emoji_id"`
}

func (MessageReactEvent) eventMarker()      {}
func (MessageReactEvent) EventType() string { return "MessageReact" }

// MessageUnreactEvent indicates a user removed their reaction.
type MessageUnreactEvent struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel_id"`
	UserID    string `json:"user_id"`
	EmojiID   string `json:"emoji_id"`
}

func (MessageUnreactEvent) eventMarker()      {}
func (MessageUnreactEvent) EventType() string { return "MessageUnreact" }

// MessageRemoveReactionEvent indicates all reactions of an emoji were removed.
type MessageRemoveReactionEvent struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel_id"`
	EmojiID   string `json:"emoji_id"`
}

func (MessageRemoveReactionEvent) eventMarker()      {}
func (MessageRemoveReactionEvent) EventType() string { return "MessageRemoveReaction" }

// --- Channel events ---

// ChannelCreateEvent indicates a new channel was created.
type ChannelCreateEvent struct {
	raw stoat.RawChannel
}

func (ChannelCreateEvent) eventMarker()      {}
func (ChannelCreateEvent) EventType() string { return "ChannelCreate" }

// Channel returns the typed channel value.
func (e *ChannelCreateEvent) Channel() stoat.Channel {
	return e.raw.Value
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *ChannelCreateEvent) UnmarshalJSON(data []byte) error {
	return e.raw.UnmarshalJSON(data)
}

// MarshalJSON implements json.Marshaler.
func (e ChannelCreateEvent) MarshalJSON() ([]byte, error) {
	return e.raw.MarshalJSON()
}

// ChannelUpdateEvent indicates a channel was edited.
type ChannelUpdateEvent struct {
	ID    string          `json:"id"`
	Data  json.RawMessage `json:"data"`
	Clear []string        `json:"clear,omitempty"`
}

func (ChannelUpdateEvent) eventMarker()      {}
func (ChannelUpdateEvent) EventType() string { return "ChannelUpdate" }

// ChannelDeleteEvent indicates a channel was deleted.
type ChannelDeleteEvent struct {
	ID string `json:"id"`
}

func (ChannelDeleteEvent) eventMarker()      {}
func (ChannelDeleteEvent) EventType() string { return "ChannelDelete" }

// ChannelGroupJoinEvent indicates a user joined a group channel.
type ChannelGroupJoinEvent struct {
	ID   string `json:"id"`
	User string `json:"user"`
}

func (ChannelGroupJoinEvent) eventMarker()      {}
func (ChannelGroupJoinEvent) EventType() string { return "ChannelGroupJoin" }

// ChannelGroupLeaveEvent indicates a user left a group channel.
type ChannelGroupLeaveEvent struct {
	ID   string `json:"id"`
	User string `json:"user"`
}

func (ChannelGroupLeaveEvent) eventMarker()      {}
func (ChannelGroupLeaveEvent) EventType() string { return "ChannelGroupLeave" }

// ChannelStartTypingEvent indicates a user started typing.
type ChannelStartTypingEvent struct {
	ID   string `json:"id"`
	User string `json:"user"`
}

func (ChannelStartTypingEvent) eventMarker()      {}
func (ChannelStartTypingEvent) EventType() string { return "ChannelStartTyping" }

// ChannelStopTypingEvent indicates a user stopped typing.
type ChannelStopTypingEvent struct {
	ID   string `json:"id"`
	User string `json:"user"`
}

func (ChannelStopTypingEvent) eventMarker()      {}
func (ChannelStopTypingEvent) EventType() string { return "ChannelStopTyping" }

// ChannelAckEvent indicates messages were acknowledged.
type ChannelAckEvent struct {
	ID        string `json:"id"`
	User      string `json:"user"`
	MessageID string `json:"message_id"`
}

func (ChannelAckEvent) eventMarker()      {}
func (ChannelAckEvent) EventType() string { return "ChannelAck" }

// --- Server events ---

// ServerCreateEvent indicates a server was created or joined.
type ServerCreateEvent struct {
	stoat.Server
}

func (ServerCreateEvent) eventMarker()      {}
func (ServerCreateEvent) EventType() string { return "ServerCreate" }

// ServerUpdateEvent indicates a server was edited.
type ServerUpdateEvent struct {
	ID    string          `json:"id"`
	Data  json.RawMessage `json:"data"`
	Clear []string        `json:"clear,omitempty"`
}

func (ServerUpdateEvent) eventMarker()      {}
func (ServerUpdateEvent) EventType() string { return "ServerUpdate" }

// ServerDeleteEvent indicates a server was deleted or left.
type ServerDeleteEvent struct {
	ID string `json:"id"`
}

func (ServerDeleteEvent) eventMarker()      {}
func (ServerDeleteEvent) EventType() string { return "ServerDelete" }

// ServerMemberUpdateEvent indicates a server member was edited.
type ServerMemberUpdateEvent struct {
	ID    stoat.MemberCompositeKey `json:"id"`
	Data  json.RawMessage          `json:"data"`
	Clear []string                 `json:"clear,omitempty"`
}

func (ServerMemberUpdateEvent) eventMarker()      {}
func (ServerMemberUpdateEvent) EventType() string { return "ServerMemberUpdate" }

// ServerMemberJoinEvent indicates a user joined a server.
type ServerMemberJoinEvent struct {
	ID     string       `json:"id"`
	User   string       `json:"user"`
	Member stoat.Member `json:"member"`
}

func (ServerMemberJoinEvent) eventMarker()      {}
func (ServerMemberJoinEvent) EventType() string { return "ServerMemberJoin" }

// ServerMemberLeaveEvent indicates a user left a server.
type ServerMemberLeaveEvent struct {
	ID   string `json:"id"`
	User string `json:"user"`
}

func (ServerMemberLeaveEvent) eventMarker()      {}
func (ServerMemberLeaveEvent) EventType() string { return "ServerMemberLeave" }

// ServerRoleUpdateEvent indicates a role was edited.
type ServerRoleUpdateEvent struct {
	ID     string          `json:"id"`
	RoleID string          `json:"role_id"`
	Data   json.RawMessage `json:"data"`
	Clear  []string        `json:"clear,omitempty"`
}

func (ServerRoleUpdateEvent) eventMarker()      {}
func (ServerRoleUpdateEvent) EventType() string { return "ServerRoleUpdate" }

// ServerRoleDeleteEvent indicates a role was deleted.
type ServerRoleDeleteEvent struct {
	ID     string `json:"id"`
	RoleID string `json:"role_id"`
}

func (ServerRoleDeleteEvent) eventMarker()      {}
func (ServerRoleDeleteEvent) EventType() string { return "ServerRoleDelete" }

// --- User events ---

// UserUpdateEvent indicates a user's profile or status changed.
type UserUpdateEvent struct {
	ID    string          `json:"id"`
	Data  json.RawMessage `json:"data"`
	Clear []string        `json:"clear,omitempty"`
}

func (UserUpdateEvent) eventMarker()      {}
func (UserUpdateEvent) EventType() string { return "UserUpdate" }

// UserRelationshipEvent indicates a relationship change.
type UserRelationshipEvent struct {
	ID     string                   `json:"id"`
	User   stoat.User               `json:"user"`
	Status stoat.RelationshipStatus `json:"status"`
}

func (UserRelationshipEvent) eventMarker()      {}
func (UserRelationshipEvent) EventType() string { return "UserRelationship" }

// UserPlatformWipeEvent indicates a user was platform-banned or deleted.
type UserPlatformWipeEvent struct {
	UserID string `json:"user_id"`
	Flags  int    `json:"flags"`
}

func (UserPlatformWipeEvent) eventMarker()      {}
func (UserPlatformWipeEvent) EventType() string { return "UserPlatformWipe" }

// --- Emoji events ---

// EmojiCreateEvent indicates a custom emoji was created.
type EmojiCreateEvent struct {
	stoat.Emoji
}

func (EmojiCreateEvent) eventMarker()      {}
func (EmojiCreateEvent) EventType() string { return "EmojiCreate" }

// EmojiDeleteEvent indicates a custom emoji was deleted.
type EmojiDeleteEvent struct {
	ID string `json:"id"`
}

func (EmojiDeleteEvent) eventMarker()      {}
func (EmojiDeleteEvent) EventType() string { return "EmojiDelete" }

// --- Auth events ---

// AuthEvent is a forwarded authentication system event.
type AuthEvent struct {
	AuthEventType    string `json:"event_type"`
	UserID           string `json:"user_id,omitempty"`
	SessionID        string `json:"session_id,omitempty"`
	ExcludeSessionID string `json:"exclude_session_id,omitempty"`
}

func (AuthEvent) eventMarker()      {}
func (AuthEvent) EventType() string { return "Auth" }

// --- RawEvent ---

// RawEvent wraps an Event for JSON unmarshalling with type discrimination.
type RawEvent struct {
	Value Event
}

// UnmarshalJSON implements json.Unmarshaler. It reads the "type" field to determine
// which concrete event struct to unmarshal into.
func (r *RawEvent) UnmarshalJSON(data []byte) error {
	var disc struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &disc); err != nil {
		return err
	}

	switch disc.Type {
	// Connection
	case "Authenticated":
		r.Value = &AuthenticatedEvent{}
	case "Ready":
		var v ReadyEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "Error":
		var v ErrorEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "Logout":
		r.Value = &LogoutEvent{}
	case "Pong":
		var v PongEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "Bulk":
		var v BulkEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v

	// Messages
	case "Message":
		var v MessageEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "MessageUpdate":
		var v MessageUpdateEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "MessageAppend":
		var v MessageAppendEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "MessageDelete":
		var v MessageDeleteEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "MessageReact":
		var v MessageReactEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "MessageUnreact":
		var v MessageUnreactEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "MessageRemoveReaction":
		var v MessageRemoveReactionEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v

	// Channels
	case "ChannelCreate":
		var v ChannelCreateEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "ChannelUpdate":
		var v ChannelUpdateEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "ChannelDelete":
		var v ChannelDeleteEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "ChannelGroupJoin":
		var v ChannelGroupJoinEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "ChannelGroupLeave":
		var v ChannelGroupLeaveEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "ChannelStartTyping":
		var v ChannelStartTypingEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "ChannelStopTyping":
		var v ChannelStopTypingEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "ChannelAck":
		var v ChannelAckEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v

	// Servers
	case "ServerCreate":
		var v ServerCreateEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "ServerUpdate":
		var v ServerUpdateEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "ServerDelete":
		var v ServerDeleteEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "ServerMemberUpdate":
		var v ServerMemberUpdateEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "ServerMemberJoin":
		var v ServerMemberJoinEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "ServerMemberLeave":
		var v ServerMemberLeaveEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "ServerRoleUpdate":
		var v ServerRoleUpdateEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "ServerRoleDelete":
		var v ServerRoleDeleteEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v

	// Users
	case "UserUpdate":
		var v UserUpdateEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "UserRelationship":
		var v UserRelationshipEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "UserPlatformWipe":
		var v UserPlatformWipeEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v

	// Emoji
	case "EmojiCreate":
		var v EmojiCreateEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "EmojiDelete":
		var v EmojiDeleteEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v

	// Auth
	case "Auth":
		var v AuthEvent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v

	default:
		return fmt.Errorf("unknown event type: %q", disc.Type)
	}

	return nil
}
