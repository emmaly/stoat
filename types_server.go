package stoat

// Server represents a Revolt server.
type Server struct {
	ID                 string                    `json:"_id"`
	Owner              string                    `json:"owner"`
	Name               string                    `json:"name"`
	Description        *string                   `json:"description,omitempty"`
	Channels           []string                  `json:"channels"`
	Categories         []Category                `json:"categories,omitempty"`
	SystemMessages     *SystemMessageChannels    `json:"system_messages,omitempty"`
	Roles              map[string]Role           `json:"roles,omitempty"`
	DefaultPermissions int64                     `json:"default_permissions"`
	Icon               *File                     `json:"icon,omitempty"`
	Banner             *File                     `json:"banner,omitempty"`
	Flags              uint32                    `json:"flags,omitempty"`
	NSFW               bool                      `json:"nsfw,omitempty"`
	Analytics          bool                      `json:"analytics,omitempty"`
	Discoverable       bool                      `json:"discoverable,omitempty"`
}

// Role represents a role within a server.
type Role struct {
	Name        string        `json:"name"`
	Permissions OverrideField `json:"permissions"`
	Colour      *string       `json:"colour,omitempty"`
	Hoist       bool          `json:"hoist,omitempty"`
	Rank        int64         `json:"rank,omitempty"`
}

// SystemMessageChannels holds the channel IDs for system event messages.
type SystemMessageChannels struct {
	UserJoined *string `json:"user_joined,omitempty"`
	UserLeft   *string `json:"user_left,omitempty"`
	UserKicked *string `json:"user_kicked,omitempty"`
	UserBanned *string `json:"user_banned,omitempty"`
}

// DataCreateServer is the request body for creating a server.
type DataCreateServer struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	NSFW        *bool   `json:"nsfw,omitempty"`
}

// DataEditServer is the request body for editing a server.
type DataEditServer struct {
	Name           *string                `json:"name,omitempty"`
	Description    *string                `json:"description,omitempty"`
	Icon           *string                `json:"icon,omitempty"`
	Banner         *string                `json:"banner,omitempty"`
	Categories     []Category             `json:"categories,omitempty"`
	SystemMessages *SystemMessageChannels  `json:"system_messages,omitempty"`
	Flags          *uint32                `json:"flags,omitempty"`
	Discoverable   *bool                  `json:"discoverable,omitempty"`
	Analytics      *bool                  `json:"analytics,omitempty"`
	Owner          *string                `json:"owner,omitempty"`
	Remove         []FieldsServer         `json:"remove,omitempty"`
}

// CreateServerLegacyResponse is returned when creating a server.
type CreateServerLegacyResponse struct {
	Server   Server       `json:"server"`
	Channels []RawChannel `json:"channels"`
}

// FetchServerResponse is the rich response from fetching a server with channels.
type FetchServerResponse struct {
	Server   Server       `json:"server"`
	Channels []RawChannel `json:"channels"`
}

// FieldsServer is a string enum of optional fields on a server object that can be removed.
type FieldsServer string

const (
	FieldsServerIcon           FieldsServer = "Icon"
	FieldsServerBanner         FieldsServer = "Banner"
	FieldsServerDescription    FieldsServer = "Description"
	FieldsServerCategories     FieldsServer = "Categories"
	FieldsServerSystemMessages FieldsServer = "SystemMessages"
)
