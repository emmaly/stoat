package stoat

import (
	"encoding/json"
	"fmt"
)


// Message represents a message in a channel.
type Message struct {
	ID           string                `json:"_id"`
	Nonce        *string               `json:"nonce,omitempty"`
	Channel      string                `json:"channel"`
	Author       string                `json:"author"`
	User         *User                 `json:"user,omitempty"`
	Member       *Member                `json:"member,omitempty"`
	Webhook      *MessageWebhook       `json:"webhook,omitempty"`
	Content      *string               `json:"content,omitempty"`
	System       *RawSystemMessage     `json:"system,omitempty"`
	Attachments  []File                `json:"attachments,omitempty"`
	Edited       *string               `json:"edited,omitempty"`
	Embeds       []RawEmbed            `json:"embeds,omitempty"`
	Mentions     []string              `json:"mentions,omitempty"`
	RoleMentions []string              `json:"role_mentions,omitempty"`
	Replies      []string              `json:"replies,omitempty"`
	Reactions    map[string][]string   `json:"reactions,omitempty"`
	Interactions *Interactions         `json:"interactions,omitempty"`
	Masquerade   *Masquerade           `json:"masquerade,omitempty"`
	Pinned       bool                  `json:"pinned,omitempty"`
	Flags        int                   `json:"flags,omitempty"`
}

// Embed is the interface for embed variants.
type Embed interface {
	embedMarker()
	EmbedType() string
}

// WebsiteEmbed represents a website embed (type="Website").
type WebsiteEmbed struct {
	URL         *string  `json:"url,omitempty"`
	OriginalURL *string  `json:"original_url,omitempty"`
	Special     *Special `json:"special,omitempty"`
	Title       *string  `json:"title,omitempty"`
	Description *string  `json:"description,omitempty"`
	Image       *Image   `json:"image,omitempty"`
	Video       *Video   `json:"video,omitempty"`
	SiteName    *string  `json:"site_name,omitempty"`
	IconURL     *string  `json:"icon_url,omitempty"`
	Colour      *string  `json:"colour,omitempty"`
}

func (WebsiteEmbed) embedMarker()      {}
func (WebsiteEmbed) EmbedType() string { return "Website" }

// ImageEmbed represents an image embed (type="Image").
type ImageEmbed struct {
	URL    string    `json:"url"`
	Width  int       `json:"width"`
	Height int       `json:"height"`
	Size   ImageSize `json:"size"`
}

func (ImageEmbed) embedMarker()      {}
func (ImageEmbed) EmbedType() string { return "Image" }

// VideoEmbed represents a video embed (type="Video").
type VideoEmbed struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (VideoEmbed) embedMarker()      {}
func (VideoEmbed) EmbedType() string { return "Video" }

// TextEmbed represents a text embed (type="Text").
type TextEmbed struct {
	IconURL     *string `json:"icon_url,omitempty"`
	URL         *string `json:"url,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Media       *File   `json:"media,omitempty"`
	Colour      *string `json:"colour,omitempty"`
}

func (TextEmbed) embedMarker()      {}
func (TextEmbed) EmbedType() string { return "Text" }

// NoneEmbed represents an empty embed (type="None").
type NoneEmbed struct{}

func (NoneEmbed) embedMarker()      {}
func (NoneEmbed) EmbedType() string { return "None" }

// RawEmbed is a wrapper that handles JSON unmarshalling of the Embed tagged union.
type RawEmbed struct {
	Value Embed
}

// MarshalJSON implements json.Marshaler.
func (r RawEmbed) MarshalJSON() ([]byte, error) {
	switch v := r.Value.(type) {
	case *WebsiteEmbed:
		return json.Marshal(struct {
			Type string `json:"type"`
			*WebsiteEmbed
		}{Type: "Website", WebsiteEmbed: v})
	case *ImageEmbed:
		return json.Marshal(struct {
			Type string `json:"type"`
			*ImageEmbed
		}{Type: "Image", ImageEmbed: v})
	case *VideoEmbed:
		return json.Marshal(struct {
			Type string `json:"type"`
			*VideoEmbed
		}{Type: "Video", VideoEmbed: v})
	case *TextEmbed:
		return json.Marshal(struct {
			Type string `json:"type"`
			*TextEmbed
		}{Type: "Text", TextEmbed: v})
	case *NoneEmbed:
		return json.Marshal(struct {
			Type string `json:"type"`
		}{Type: "None"})
	default:
		return nil, fmt.Errorf("unknown Embed variant: %T", r.Value)
	}
}

// UnmarshalJSON implements json.Unmarshaler for the Embed tagged union.
func (r *RawEmbed) UnmarshalJSON(data []byte) error {
	var discriminator struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &discriminator); err != nil {
		return err
	}

	switch discriminator.Type {
	case "Website":
		var v WebsiteEmbed
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "Image":
		var v ImageEmbed
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "Video":
		var v VideoEmbed
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "Text":
		var v TextEmbed
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "None":
		r.Value = &NoneEmbed{}
	default:
		return fmt.Errorf("unknown embed type: %q", discriminator.Type)
	}
	return nil
}

// SystemMessage is the interface for system message variants.
type SystemMessage interface {
	systemMessageMarker()
	SystemMessageType() string
}

// SystemMessageText represents a system text message.
type SystemMessageText struct {
	Content string `json:"content"`
}

func (SystemMessageText) systemMessageMarker()      {}
func (SystemMessageText) SystemMessageType() string  { return "text" }

// SystemMessageUserAdded represents a user being added to a group.
type SystemMessageUserAdded struct {
	ID string `json:"id"`
	By string `json:"by"`
}

func (SystemMessageUserAdded) systemMessageMarker()      {}
func (SystemMessageUserAdded) SystemMessageType() string  { return "user_added" }

// SystemMessageUserRemove represents a user being removed from a group.
type SystemMessageUserRemove struct {
	ID string `json:"id"`
	By string `json:"by"`
}

func (SystemMessageUserRemove) systemMessageMarker()      {}
func (SystemMessageUserRemove) SystemMessageType() string  { return "user_remove" }

// SystemMessageUserJoined represents a user joining a server.
type SystemMessageUserJoined struct {
	ID string `json:"id"`
}

func (SystemMessageUserJoined) systemMessageMarker()      {}
func (SystemMessageUserJoined) SystemMessageType() string  { return "user_joined" }

// SystemMessageUserLeft represents a user leaving a server.
type SystemMessageUserLeft struct {
	ID string `json:"id"`
}

func (SystemMessageUserLeft) systemMessageMarker()      {}
func (SystemMessageUserLeft) SystemMessageType() string  { return "user_left" }

// SystemMessageUserKicked represents a user being kicked.
type SystemMessageUserKicked struct {
	ID string `json:"id"`
}

func (SystemMessageUserKicked) systemMessageMarker()      {}
func (SystemMessageUserKicked) SystemMessageType() string  { return "user_kicked" }

// SystemMessageUserBanned represents a user being banned.
type SystemMessageUserBanned struct {
	ID string `json:"id"`
}

func (SystemMessageUserBanned) systemMessageMarker()      {}
func (SystemMessageUserBanned) SystemMessageType() string  { return "user_banned" }

// SystemMessageChannelRenamed represents a channel being renamed.
type SystemMessageChannelRenamed struct {
	Name string `json:"name"`
	By   string `json:"by"`
}

func (SystemMessageChannelRenamed) systemMessageMarker()      {}
func (SystemMessageChannelRenamed) SystemMessageType() string  { return "channel_renamed" }

// SystemMessageChannelDescriptionChanged represents a channel description being changed.
type SystemMessageChannelDescriptionChanged struct {
	By string `json:"by"`
}

func (SystemMessageChannelDescriptionChanged) systemMessageMarker()      {}
func (SystemMessageChannelDescriptionChanged) SystemMessageType() string  { return "channel_description_changed" }

// SystemMessageChannelIconChanged represents a channel icon being changed.
type SystemMessageChannelIconChanged struct {
	By string `json:"by"`
}

func (SystemMessageChannelIconChanged) systemMessageMarker()      {}
func (SystemMessageChannelIconChanged) SystemMessageType() string  { return "channel_icon_changed" }

// SystemMessageChannelOwnershipChanged represents channel ownership being transferred.
type SystemMessageChannelOwnershipChanged struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func (SystemMessageChannelOwnershipChanged) systemMessageMarker()      {}
func (SystemMessageChannelOwnershipChanged) SystemMessageType() string  { return "channel_ownership_changed" }

// SystemMessageMessagePinned represents a message being pinned.
type SystemMessageMessagePinned struct {
	ID string `json:"id"`
	By string `json:"by"`
}

func (SystemMessageMessagePinned) systemMessageMarker()      {}
func (SystemMessageMessagePinned) SystemMessageType() string  { return "message_pinned" }

// SystemMessageMessageUnpinned represents a message being unpinned.
type SystemMessageMessageUnpinned struct {
	ID string `json:"id"`
	By string `json:"by"`
}

func (SystemMessageMessageUnpinned) systemMessageMarker()      {}
func (SystemMessageMessageUnpinned) SystemMessageType() string  { return "message_unpinned" }

// SystemMessageCallStarted represents a call being started.
type SystemMessageCallStarted struct {
	By         string  `json:"by"`
	FinishedAt *string `json:"finished_at,omitempty"`
}

func (SystemMessageCallStarted) systemMessageMarker()      {}
func (SystemMessageCallStarted) SystemMessageType() string  { return "call_started" }

// RawSystemMessage is a wrapper that handles JSON unmarshalling of the SystemMessage tagged union.
type RawSystemMessage struct {
	Value SystemMessage
}

// MarshalJSON implements json.Marshaler.
func (r RawSystemMessage) MarshalJSON() ([]byte, error) {
	switch v := r.Value.(type) {
	case *SystemMessageText:
		return json.Marshal(struct {
			Type string `json:"type"`
			*SystemMessageText
		}{Type: "text", SystemMessageText: v})
	case *SystemMessageUserAdded:
		return json.Marshal(struct {
			Type string `json:"type"`
			*SystemMessageUserAdded
		}{Type: "user_added", SystemMessageUserAdded: v})
	case *SystemMessageUserRemove:
		return json.Marshal(struct {
			Type string `json:"type"`
			*SystemMessageUserRemove
		}{Type: "user_remove", SystemMessageUserRemove: v})
	case *SystemMessageUserJoined:
		return json.Marshal(struct {
			Type string `json:"type"`
			*SystemMessageUserJoined
		}{Type: "user_joined", SystemMessageUserJoined: v})
	case *SystemMessageUserLeft:
		return json.Marshal(struct {
			Type string `json:"type"`
			*SystemMessageUserLeft
		}{Type: "user_left", SystemMessageUserLeft: v})
	case *SystemMessageUserKicked:
		return json.Marshal(struct {
			Type string `json:"type"`
			*SystemMessageUserKicked
		}{Type: "user_kicked", SystemMessageUserKicked: v})
	case *SystemMessageUserBanned:
		return json.Marshal(struct {
			Type string `json:"type"`
			*SystemMessageUserBanned
		}{Type: "user_banned", SystemMessageUserBanned: v})
	case *SystemMessageChannelRenamed:
		return json.Marshal(struct {
			Type string `json:"type"`
			*SystemMessageChannelRenamed
		}{Type: "channel_renamed", SystemMessageChannelRenamed: v})
	case *SystemMessageChannelDescriptionChanged:
		return json.Marshal(struct {
			Type string `json:"type"`
			*SystemMessageChannelDescriptionChanged
		}{Type: "channel_description_changed", SystemMessageChannelDescriptionChanged: v})
	case *SystemMessageChannelIconChanged:
		return json.Marshal(struct {
			Type string `json:"type"`
			*SystemMessageChannelIconChanged
		}{Type: "channel_icon_changed", SystemMessageChannelIconChanged: v})
	case *SystemMessageChannelOwnershipChanged:
		return json.Marshal(struct {
			Type string `json:"type"`
			*SystemMessageChannelOwnershipChanged
		}{Type: "channel_ownership_changed", SystemMessageChannelOwnershipChanged: v})
	case *SystemMessageMessagePinned:
		return json.Marshal(struct {
			Type string `json:"type"`
			*SystemMessageMessagePinned
		}{Type: "message_pinned", SystemMessageMessagePinned: v})
	case *SystemMessageMessageUnpinned:
		return json.Marshal(struct {
			Type string `json:"type"`
			*SystemMessageMessageUnpinned
		}{Type: "message_unpinned", SystemMessageMessageUnpinned: v})
	case *SystemMessageCallStarted:
		return json.Marshal(struct {
			Type string `json:"type"`
			*SystemMessageCallStarted
		}{Type: "call_started", SystemMessageCallStarted: v})
	default:
		return nil, fmt.Errorf("unknown SystemMessage variant: %T", r.Value)
	}
}

// UnmarshalJSON implements json.Unmarshaler for the SystemMessage tagged union.
func (r *RawSystemMessage) UnmarshalJSON(data []byte) error {
	var discriminator struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &discriminator); err != nil {
		return err
	}

	switch discriminator.Type {
	case "text":
		var v SystemMessageText
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "user_added":
		var v SystemMessageUserAdded
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "user_remove":
		var v SystemMessageUserRemove
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "user_joined":
		var v SystemMessageUserJoined
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "user_left":
		var v SystemMessageUserLeft
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "user_kicked":
		var v SystemMessageUserKicked
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "user_banned":
		var v SystemMessageUserBanned
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "channel_renamed":
		var v SystemMessageChannelRenamed
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "channel_description_changed":
		var v SystemMessageChannelDescriptionChanged
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "channel_icon_changed":
		var v SystemMessageChannelIconChanged
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "channel_ownership_changed":
		var v SystemMessageChannelOwnershipChanged
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "message_pinned":
		var v SystemMessageMessagePinned
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "message_unpinned":
		var v SystemMessageMessageUnpinned
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "call_started":
		var v SystemMessageCallStarted
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	default:
		return fmt.Errorf("unknown system message type: %q", discriminator.Type)
	}
	return nil
}

// DataMessageSend is the request body for sending a message.
type DataMessageSend struct {
	Nonce          *string         `json:"nonce,omitempty"`
	Content        *string         `json:"content,omitempty"`
	Attachments    []string        `json:"attachments,omitempty"`
	Replies        []ReplyIntent   `json:"replies,omitempty"`
	Embeds         []SendableEmbed `json:"embeds,omitempty"`
	Masquerade     *Masquerade     `json:"masquerade,omitempty"`
	Interactions   *Interactions   `json:"interactions,omitempty"`
	Flags          *int            `json:"flags,omitempty"`
	IdempotencyKey string          `json:"-"` // Set as Idempotency-Key header, not in body
}

// DataEditMessage is the request body for editing a message.
type DataEditMessage struct {
	Content *string         `json:"content,omitempty"`
	Embeds  []SendableEmbed `json:"embeds,omitempty"`
}

// DataMessageSearch is the request body for searching messages.
type DataMessageSearch struct {
	Query        *string      `json:"query,omitempty"`
	Pinned       *bool        `json:"pinned,omitempty"`
	Limit        *int         `json:"limit,omitempty"`
	Before       *string      `json:"before,omitempty"`
	After        *string      `json:"after,omitempty"`
	Sort         *MessageSort `json:"sort,omitempty"`
	IncludeUsers *bool        `json:"include_users,omitempty"`
}

// SendableEmbed is a text embed to include in a message.
type SendableEmbed struct {
	IconURL     *string `json:"icon_url,omitempty"`
	URL         *string `json:"url,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Media       *string `json:"media,omitempty"`
	Colour      *string `json:"colour,omitempty"`
}

// ReplyIntent specifies what a message should reply to.
type ReplyIntent struct {
	ID              string `json:"id"`
	Mention         bool   `json:"mention"`
	FailIfNotExists *bool  `json:"fail_if_not_exists,omitempty"`
}

// Masquerade represents name/avatar override information for a message.
type Masquerade struct {
	Name   *string `json:"name,omitempty"`
	Avatar *string `json:"avatar,omitempty"`
	Colour *string `json:"colour,omitempty"`
}

// Interactions represents information about how a message should be interacted with.
type Interactions struct {
	Reactions         []string `json:"reactions,omitempty"`
	RestrictReactions bool     `json:"restrict_reactions,omitempty"`
}

// MessageWebhook contains information about the webhook that sent a message.
type MessageWebhook struct {
	Name   string  `json:"name"`
	Avatar *string `json:"avatar,omitempty"`
}

// MessageSort is a string enum for message sort order.
type MessageSort string

const (
	MessageSortRelevance MessageSort = "Relevance"
	MessageSortLatest    MessageSort = "Latest"
	MessageSortOldest    MessageSort = "Oldest"
)

// BulkMessageResponse is the response from bulk message endpoints.
type BulkMessageResponse struct {
	Messages []Message          `json:"messages"`
	Users    []User             `json:"users,omitempty"`
	Members  []Member            `json:"members,omitempty"`
}

// OptionsBulkDelete is the request body for bulk deleting messages.
type OptionsBulkDelete struct {
	IDs []string `json:"ids"`
}

// FetchMessagesOptions are query parameters for fetching messages.
type FetchMessagesOptions struct {
	Limit        *int
	Before       *string
	After        *string
	Sort         *MessageSort
	Nearby       *string
	IncludeUsers *bool
}
