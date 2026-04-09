package stoat

import (
	"encoding/json"
	"fmt"
)

// Emoji represents a custom emoji.
type Emoji struct {
	ID        string        `json:"_id"`
	Parent    RawEmojiParent `json:"parent"`
	CreatorID string        `json:"creator_id"`
	Name      string        `json:"name"`
	Animated  bool          `json:"animated,omitempty"`
	NSFW      bool          `json:"nsfw,omitempty"`
}

// EmojiParent is the interface for emoji parent variants.
type EmojiParent interface {
	emojiParentMarker()
	EmojiParentType() string
}

// ServerEmojiParent represents a server emoji parent (type="Server").
type ServerEmojiParent struct {
	ID string `json:"id"`
}

func (ServerEmojiParent) emojiParentMarker()      {}
func (ServerEmojiParent) EmojiParentType() string  { return "Server" }

// DetachedEmojiParent represents a detached emoji parent (type="Detached").
type DetachedEmojiParent struct{}

func (DetachedEmojiParent) emojiParentMarker()      {}
func (DetachedEmojiParent) EmojiParentType() string  { return "Detached" }

// RawEmojiParent is a wrapper that handles JSON marshalling/unmarshalling of the EmojiParent tagged union.
type RawEmojiParent struct {
	Value EmojiParent
}

// MarshalJSON implements json.Marshaler.
func (r RawEmojiParent) MarshalJSON() ([]byte, error) {
	switch v := r.Value.(type) {
	case *ServerEmojiParent:
		return json.Marshal(struct {
			Type string `json:"type"`
			*ServerEmojiParent
		}{Type: "Server", ServerEmojiParent: v})
	case *DetachedEmojiParent:
		return json.Marshal(struct {
			Type string `json:"type"`
		}{Type: "Detached"})
	default:
		return nil, fmt.Errorf("unknown EmojiParent variant: %T", r.Value)
	}
}

// UnmarshalJSON implements json.Unmarshaler for the EmojiParent tagged union.
func (r *RawEmojiParent) UnmarshalJSON(data []byte) error {
	var discriminator struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &discriminator); err != nil {
		return err
	}

	switch discriminator.Type {
	case "Server":
		var v ServerEmojiParent
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Value = &v
	case "Detached":
		r.Value = &DetachedEmojiParent{}
	default:
		return fmt.Errorf("unknown emoji parent type: %q", discriminator.Type)
	}
	return nil
}

// DataCreateEmoji is the request body for creating an emoji.
type DataCreateEmoji struct {
	Name   string         `json:"name"`
	Parent RawEmojiParent `json:"parent"`
	NSFW   bool           `json:"nsfw,omitempty"`
}
