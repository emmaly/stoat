package stoat

import (
	"context"
	"fmt"
	"net/http"
)

// RemoveReactionOptions are optional query parameters for removing reactions.
type RemoveReactionOptions struct {
	UserID    *string
	RemoveAll *bool
}

// AddReaction adds a reaction to a message. Requires session token.
func (c *Client) AddReaction(ctx context.Context, channelID, msgID, emoji string) error {
	req, err := c.request(ctx, http.MethodPut, fmt.Sprintf("/channels/%s/messages/%s/reactions/%s", channelID, msgID, emoji), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// RemoveReaction removes a reaction from a message. Requires session token.
func (c *Client) RemoveReaction(ctx context.Context, channelID, msgID, emoji string, opts *RemoveReactionOptions) error {
	path := fmt.Sprintf("/channels/%s/messages/%s/reactions/%s", channelID, msgID, emoji)
	if opts != nil {
		q := ""
		sep := "?"
		if opts.UserID != nil {
			q += sep + "user_id=" + *opts.UserID
			sep = "&"
		}
		if opts.RemoveAll != nil && *opts.RemoveAll {
			q += sep + "remove_all=true"
		}
		path += q
	}
	req, err := c.request(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// ClearReactions removes all reactions from a message. Requires session token.
func (c *Client) ClearReactions(ctx context.Context, channelID, msgID string) error {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/channels/%s/messages/%s/reactions", channelID, msgID), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
