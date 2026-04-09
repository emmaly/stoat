package stoat

import (
	"encoding/json"
	"fmt"
)

// Invite represents a server invite.
type Invite struct {
	ID      string `json:"_id"`
	Server  string `json:"server"`
	Creator string `json:"creator"`
	Channel string `json:"channel"`
}

// InviteResponse is the interface for invite response variants.
type InviteResponse interface {
	inviteResponseMarker()
	InviteResponseType() string
}

// ServerInviteResponse represents a server invite response (type="Server").
type ServerInviteResponse struct {
	Code               string  `json:"code"`
	ServerID           string  `json:"server_id"`
	ServerName         string  `json:"server_name"`
	ServerIcon         *File   `json:"server_icon,omitempty"`
	ServerBanner       *File   `json:"server_banner,omitempty"`
	ServerFlags        uint32  `json:"server_flags,omitempty"`
	ChannelID          string  `json:"channel_id"`
	ChannelName        string  `json:"channel_name"`
	ChannelDescription *string `json:"channel_description,omitempty"`
	UserName           string  `json:"user_name"`
	UserAvatar         *File   `json:"user_avatar,omitempty"`
	MemberCount        int     `json:"member_count"`
	OnlineCount        *int    `json:"online_count,omitempty"`
}

func (ServerInviteResponse) inviteResponseMarker()      {}
func (ServerInviteResponse) InviteResponseType() string  { return "Server" }

// GroupInviteResponse represents a group invite response (type="Group").
type GroupInviteResponse struct {
	Code               string  `json:"code"`
	ChannelID          string  `json:"channel_id"`
	ChannelName        string  `json:"channel_name"`
	ChannelDescription *string `json:"channel_description,omitempty"`
	UserName           string  `json:"user_name"`
	UserAvatar         *File   `json:"user_avatar,omitempty"`
}

func (GroupInviteResponse) inviteResponseMarker()      {}
func (GroupInviteResponse) InviteResponseType() string  { return "Group" }

// RawInviteResponse is a wrapper that handles JSON unmarshalling of the InviteResponse tagged union.
type RawInviteResponse struct {
	Value InviteResponse
}

// MarshalJSON implements json.Marshaler.
func (r RawInviteResponse) MarshalJSON() ([]byte, error) {
	switch v := r.Value.(type) {
	case *ServerInviteResponse:
		return json.Marshal(struct {
			Type string `json:"type"`
			*ServerInviteResponse
		}{Type: "Server", ServerInviteResponse: v})
	case *GroupInviteResponse:
		return json.Marshal(struct {
			Type string `json:"type"`
			*GroupInviteResponse
		}{Type: "Group", GroupInviteResponse: v})
	default:
		return nil, fmt.Errorf("unknown InviteResponse variant: %T", r.Value)
	}
}

// UnmarshalJSON implements json.Unmarshaler for the InviteResponse tagged union.
func (r *RawInviteResponse) UnmarshalJSON(data []byte) error {
	var discriminator struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &discriminator); err != nil {
		return err
	}

	switch discriminator.Type {
	case "Server":
		var v ServerInviteResponse
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "Group":
		var v GroupInviteResponse
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	default:
		return fmt.Errorf("unknown invite response type: %q", discriminator.Type)
	}
	return nil
}

// InviteJoinResponse is the interface for invite join response variants.
type InviteJoinResponse interface {
	inviteJoinResponseMarker()
	InviteJoinResponseType() string
}

// ServerInviteJoinResponse represents a server invite join response (type="Server").
type ServerInviteJoinResponse struct {
	Channels []RawChannel `json:"channels"`
	Server   Server       `json:"server"`
}

func (ServerInviteJoinResponse) inviteJoinResponseMarker()      {}
func (ServerInviteJoinResponse) InviteJoinResponseType() string  { return "Server" }

// GroupInviteJoinResponse represents a group invite join response (type="Group").
type GroupInviteJoinResponse struct {
	Channel RawChannel `json:"channel"`
	Users   []User     `json:"users"`
}

func (GroupInviteJoinResponse) inviteJoinResponseMarker()      {}
func (GroupInviteJoinResponse) InviteJoinResponseType() string  { return "Group" }

// RawInviteJoinResponse is a wrapper that handles JSON unmarshalling of the InviteJoinResponse tagged union.
type RawInviteJoinResponse struct {
	Value InviteJoinResponse
}

// MarshalJSON implements json.Marshaler.
func (r RawInviteJoinResponse) MarshalJSON() ([]byte, error) {
	switch v := r.Value.(type) {
	case *ServerInviteJoinResponse:
		return json.Marshal(struct {
			Type string `json:"type"`
			*ServerInviteJoinResponse
		}{Type: "Server", ServerInviteJoinResponse: v})
	case *GroupInviteJoinResponse:
		return json.Marshal(struct {
			Type string `json:"type"`
			*GroupInviteJoinResponse
		}{Type: "Group", GroupInviteJoinResponse: v})
	default:
		return nil, fmt.Errorf("unknown InviteJoinResponse variant: %T", r.Value)
	}
}

// UnmarshalJSON implements json.Unmarshaler for the InviteJoinResponse tagged union.
func (r *RawInviteJoinResponse) UnmarshalJSON(data []byte) error {
	var discriminator struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &discriminator); err != nil {
		return err
	}

	switch discriminator.Type {
	case "Server":
		var v ServerInviteJoinResponse
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "Group":
		var v GroupInviteJoinResponse
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	default:
		return fmt.Errorf("unknown invite join response type: %q", discriminator.Type)
	}
	return nil
}
