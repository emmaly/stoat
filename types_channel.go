package stoat

import (
	"encoding/json"
	"fmt"
)

// Channel is the interface for channel variants.
type Channel interface {
	channelMarker()
	ChannelType() string
	ChannelID() string
}

// SavedMessagesChannel represents a user's saved messages channel.
type SavedMessagesChannel struct {
	ID   string `json:"_id"`
	User string `json:"user"`
}

func (SavedMessagesChannel) channelMarker()      {}
func (SavedMessagesChannel) ChannelType() string  { return "SavedMessages" }
func (c SavedMessagesChannel) ChannelID() string  { return c.ID }

// DirectMessageChannel represents a direct message channel between two users.
type DirectMessageChannel struct {
	ID            string   `json:"_id"`
	Active        bool     `json:"active"`
	Recipients    []string `json:"recipients"`
	LastMessageID *string  `json:"last_message_id,omitempty"`
}

func (DirectMessageChannel) channelMarker()      {}
func (DirectMessageChannel) ChannelType() string  { return "DirectMessage" }
func (c DirectMessageChannel) ChannelID() string  { return c.ID }

// GroupChannel represents a group channel.
type GroupChannel struct {
	ID            string   `json:"_id"`
	Name          string   `json:"name"`
	Owner         string   `json:"owner"`
	Description   *string  `json:"description,omitempty"`
	Recipients    []string `json:"recipients"`
	Icon          *File    `json:"icon,omitempty"`
	LastMessageID *string  `json:"last_message_id,omitempty"`
	Permissions   *int64   `json:"permissions,omitempty"`
	NSFW          bool     `json:"nsfw,omitempty"`
}

func (GroupChannel) channelMarker()      {}
func (GroupChannel) ChannelType() string  { return "Group" }
func (c GroupChannel) ChannelID() string  { return c.ID }

// TextChannel represents a text channel in a server.
type TextChannel struct {
	ID                 string              `json:"_id"`
	Server             string              `json:"server"`
	Name               string              `json:"name"`
	Description        *string             `json:"description,omitempty"`
	Icon               *File               `json:"icon,omitempty"`
	LastMessageID      *string             `json:"last_message_id,omitempty"`
	DefaultPermissions *OverrideField      `json:"default_permissions,omitempty"`
	RolePermissions    map[string]OverrideField `json:"role_permissions,omitempty"`
	NSFW               bool                `json:"nsfw,omitempty"`
}

func (TextChannel) channelMarker()      {}
func (TextChannel) ChannelType() string  { return "TextChannel" }
func (c TextChannel) ChannelID() string  { return c.ID }

// VoiceChannel represents a voice channel in a server.
type VoiceChannel struct {
	ID                 string              `json:"_id"`
	Server             string              `json:"server"`
	Name               string              `json:"name"`
	Description        *string             `json:"description,omitempty"`
	Icon               *File               `json:"icon,omitempty"`
	DefaultPermissions *OverrideField      `json:"default_permissions,omitempty"`
	RolePermissions    map[string]OverrideField `json:"role_permissions,omitempty"`
	NSFW               bool                `json:"nsfw,omitempty"`
}

func (VoiceChannel) channelMarker()      {}
func (VoiceChannel) ChannelType() string  { return "VoiceChannel" }
func (c VoiceChannel) ChannelID() string  { return c.ID }

// RawChannel is a wrapper that handles JSON unmarshalling of the Channel tagged union.
type RawChannel struct {
	Value Channel
}

// MarshalJSON implements json.Marshaler.
func (r RawChannel) MarshalJSON() ([]byte, error) {
	switch v := r.Value.(type) {
	case *SavedMessagesChannel:
		return json.Marshal(struct {
			ChannelType string `json:"channel_type"`
			*SavedMessagesChannel
		}{ChannelType: "SavedMessages", SavedMessagesChannel: v})
	case *DirectMessageChannel:
		return json.Marshal(struct {
			ChannelType string `json:"channel_type"`
			*DirectMessageChannel
		}{ChannelType: "DirectMessage", DirectMessageChannel: v})
	case *GroupChannel:
		return json.Marshal(struct {
			ChannelType string `json:"channel_type"`
			*GroupChannel
		}{ChannelType: "Group", GroupChannel: v})
	case *TextChannel:
		return json.Marshal(struct {
			ChannelType string `json:"channel_type"`
			*TextChannel
		}{ChannelType: "TextChannel", TextChannel: v})
	case *VoiceChannel:
		return json.Marshal(struct {
			ChannelType string `json:"channel_type"`
			*VoiceChannel
		}{ChannelType: "VoiceChannel", VoiceChannel: v})
	default:
		return nil, fmt.Errorf("unknown Channel variant: %T", r.Value)
	}
}

// UnmarshalJSON implements json.Unmarshaler for the Channel tagged union.
func (r *RawChannel) UnmarshalJSON(data []byte) error {
	var discriminator struct {
		ChannelType string `json:"channel_type"`
	}
	if err := json.Unmarshal(data, &discriminator); err != nil {
		return err
	}

	switch discriminator.ChannelType {
	case "SavedMessages":
		var v SavedMessagesChannel
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "DirectMessage":
		var v DirectMessageChannel
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "Group":
		var v GroupChannel
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "TextChannel":
		var v TextChannel
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "VoiceChannel":
		var v VoiceChannel
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	default:
		return fmt.Errorf("unknown channel type: %q", discriminator.ChannelType)
	}
	return nil
}

// Override represents a permission override with allow and deny bit flags.
// Used in API request bodies (e.g., DataSetRolePermissions).
type Override struct {
	Allow uint64 `json:"allow"`
	Deny  uint64 `json:"deny"`
}

// OverrideField represents a permission override as stored in the database.
// Uses short field names "a" and "d".
type OverrideField struct {
	A uint64 `json:"a"`
	D uint64 `json:"d"`
}

// Category represents a channel category in a server.
type Category struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Channels []string `json:"channels"`
}

// ChannelCompositeKey is a composite primary key consisting of channel and user ID.
type ChannelCompositeKey struct {
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// DataEditChannel is the request body for editing a channel.
type DataEditChannel struct {
	Name        *string       `json:"name,omitempty"`
	Description *string       `json:"description,omitempty"`
	Owner       *string       `json:"owner,omitempty"`
	Icon        *string       `json:"icon,omitempty"`
	NSFW        *bool         `json:"nsfw,omitempty"`
	Archived    *bool         `json:"archived,omitempty"`
	Remove      []FieldsChannel `json:"remove,omitempty"`
}

// DataCreateGroup is the request body for creating a group.
type DataCreateGroup struct {
	Name        string   `json:"name"`
	Description *string  `json:"description,omitempty"`
	Icon        *string  `json:"icon,omitempty"`
	Users       []string `json:"users,omitempty"`
	NSFW        *bool    `json:"nsfw,omitempty"`
}

// DataCreateServerChannel is the request body for creating a server channel.
type DataCreateServerChannel struct {
	Type        *string `json:"type,omitempty"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	NSFW        *bool   `json:"nsfw,omitempty"`
}

// DataDefaultChannelPermissions is the request body for setting default channel permissions.
type DataDefaultChannelPermissions struct {
	Permissions Override `json:"permissions"`
}

// DataSetRolePermissions is the request body for setting role permissions on a channel.
type DataSetRolePermissions struct {
	Permissions Override `json:"permissions"`
}

// FieldsChannel is a string enum of optional fields on a channel object.
type FieldsChannel string

const (
	FieldsChannelIcon               FieldsChannel = "Icon"
	FieldsChannelDescription        FieldsChannel = "Description"
	FieldsChannelDefaultPermissions FieldsChannel = "DefaultPermissions"
	FieldsChannelVoice              FieldsChannel = "Voice"
)

// VoiceInformation represents voice information for a channel.
type VoiceInformation struct {
	MaxUsers *int `json:"max_users,omitempty"`
}

// LegacyServerChannelType is a string enum for server channel types.
type LegacyServerChannelType string

const (
	LegacyServerChannelTypeText  LegacyServerChannelType = "Text"
	LegacyServerChannelTypeVoice LegacyServerChannelType = "Voice"
)
