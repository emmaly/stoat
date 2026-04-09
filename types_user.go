package stoat

import "encoding/json"

// Presence represents a user's presence status.
type Presence string

const (
	PresenceOnline    Presence = "Online"
	PresenceIdle      Presence = "Idle"
	PresenceFocus     Presence = "Focus"
	PresenceBusy      Presence = "Busy"
	PresenceInvisible Presence = "Invisible"
)

// RelationshipStatus represents a user's relationship with another user.
type RelationshipStatus string

const (
	RelationshipStatusNone         RelationshipStatus = "None"
	RelationshipStatusUser         RelationshipStatus = "User"
	RelationshipStatusFriend       RelationshipStatus = "Friend"
	RelationshipStatusOutgoing     RelationshipStatus = "Outgoing"
	RelationshipStatusIncoming     RelationshipStatus = "Incoming"
	RelationshipStatusBlocked      RelationshipStatus = "Blocked"
	RelationshipStatusBlockedOther RelationshipStatus = "BlockedOther"
)

// FieldsUser represents optional fields on the user object that can be removed.
type FieldsUser string

const (
	FieldsUserAvatar            FieldsUser = "Avatar"
	FieldsUserStatusText        FieldsUser = "StatusText"
	FieldsUserStatusPresence    FieldsUser = "StatusPresence"
	FieldsUserProfileContent    FieldsUser = "ProfileContent"
	FieldsUserProfileBackground FieldsUser = "ProfileBackground"
	FieldsUserDisplayName       FieldsUser = "DisplayName"
)

// User represents a Revolt user.
type User struct {
	ID            string             `json:"_id"`
	Username      string             `json:"username"`
	Discriminator string             `json:"discriminator"`
	DisplayName   *string            `json:"display_name,omitempty"`
	Avatar        *json.RawMessage   `json:"avatar,omitempty"` // TODO: use *File when File type is defined (Phase 4)
	Relations     []Relationship     `json:"relations,omitempty"`
	Badges        uint32             `json:"badges,omitempty"`
	Status        *UserStatus        `json:"status,omitempty"`
	Flags         uint32             `json:"flags,omitempty"`
	Privileged    bool               `json:"privileged,omitempty"`
	Bot           *BotInformation    `json:"bot,omitempty"`
	Relationship  RelationshipStatus `json:"relationship"`
	Online        bool               `json:"online"`
}

// UserStatus represents a user's active status.
type UserStatus struct {
	Text     *string   `json:"text,omitempty"`
	Presence *Presence `json:"presence,omitempty"`
}

// UserProfile represents a user's profile.
type UserProfile struct {
	Content    *string          `json:"content,omitempty"`
	Background *json.RawMessage `json:"background,omitempty"` // TODO: use *File when File type is defined (Phase 4)
}

// Relationship represents a relationship entry with another user.
type Relationship struct {
	ID     string             `json:"_id"`
	Status RelationshipStatus `json:"status"`
}

// BotInformation contains bot information for bot users.
type BotInformation struct {
	Owner string `json:"owner"`
}

// FlagResponse is the response for fetching user flags.
type FlagResponse struct {
	Flags uint32 `json:"flags"`
}

// MutualResponse is the response for fetching mutual friends, servers, groups, and DMs.
type MutualResponse struct {
	Users    []string `json:"users"`
	Servers  []string `json:"servers"`
	Channels []string `json:"channels,omitempty"`
}

// DataEditUser is the request body for editing a user.
type DataEditUser struct {
	DisplayName *string          `json:"display_name,omitempty"`
	Avatar      *string          `json:"avatar,omitempty"`
	Status      *UserStatus      `json:"status,omitempty"`
	Profile     *DataUserProfile `json:"profile,omitempty"`
	Badges      *uint32          `json:"badges,omitempty"`
	Flags       *uint32          `json:"flags,omitempty"`
	Remove      []FieldsUser     `json:"remove,omitempty"`
}

// DataUserProfile is the request body for editing a user's profile.
type DataUserProfile struct {
	Content    *string `json:"content,omitempty"`
	Background *string `json:"background,omitempty"`
}

// DataChangeUsername is the request body for changing a username.
type DataChangeUsername struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// DataSendFriendRequest is the request body for sending a friend request.
type DataSendFriendRequest struct {
	Username string `json:"username"`
}
