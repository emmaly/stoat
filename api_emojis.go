package stoat

import (
	"context"
	"fmt"
	"net/http"
)

// FetchEmoji fetches a custom emoji by ID. No auth required.
func (c *Client) FetchEmoji(ctx context.Context, emojiID string) (*Emoji, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/custom/emoji/%s", emojiID), nil)
	if err != nil {
		return nil, err
	}
	var e Emoji
	if err := c.do(req, &e); err != nil {
		return nil, err
	}
	return &e, nil
}

// CreateEmoji creates a custom emoji. The emojiID should be a file ID from
// a CDN upload to the "emojis" tag. Requires session token.
func (c *Client) CreateEmoji(ctx context.Context, emojiID string, data DataCreateEmoji) (*Emoji, error) {
	req, err := c.request(ctx, http.MethodPut, fmt.Sprintf("/custom/emoji/%s", emojiID), data)
	if err != nil {
		return nil, err
	}
	var e Emoji
	if err := c.do(req, &e); err != nil {
		return nil, err
	}
	return &e, nil
}

// DeleteEmoji deletes a custom emoji. Requires session token.
func (c *Client) DeleteEmoji(ctx context.Context, emojiID string) error {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/custom/emoji/%s", emojiID), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// FetchServerEmoji fetches all custom emoji for a server. Requires session token.
func (c *Client) FetchServerEmoji(ctx context.Context, serverID string) ([]Emoji, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/servers/%s/emojis", serverID), nil)
	if err != nil {
		return nil, err
	}
	var emojis []Emoji
	if err := c.do(req, &emojis); err != nil {
		return nil, err
	}
	return emojis, nil
}
