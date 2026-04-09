package stoat

// Webhook represents a webhook.
type Webhook struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Avatar      *File   `json:"avatar,omitempty"`
	CreatorID   string  `json:"creator_id"`
	ChannelID   string  `json:"channel_id"`
	Permissions uint64  `json:"permissions"`
	Token       *string `json:"token,omitempty"`
}

// CreateWebhookBody is the request body for creating a webhook.
type CreateWebhookBody struct {
	Name   string  `json:"name"`
	Avatar *string `json:"avatar,omitempty"`
}

// DataEditWebhook is the request body for editing a webhook.
type DataEditWebhook struct {
	Name        *string         `json:"name,omitempty"`
	Avatar      *string         `json:"avatar,omitempty"`
	Permissions *uint64         `json:"permissions,omitempty"`
	Remove      []FieldsWebhook `json:"remove,omitempty"`
}

// FieldsWebhook is a string enum of optional fields on a webhook object that can be removed.
type FieldsWebhook string

const (
	FieldsWebhookAvatar FieldsWebhook = "Avatar"
)
