package stoat

import (
	"context"
	"encoding/json"
	"net/http"
)

// FetchSettings retrieves client-specific settings. Requires session token.
// The response is a map of key to [timestamp, value] tuples represented as json.RawMessage.
func (c *Client) FetchSettings(ctx context.Context, opts OptionsFetchSettings) (map[string]json.RawMessage, error) {
	req, err := c.request(ctx, http.MethodPost, "/sync/settings/fetch", opts)
	if err != nil {
		return nil, err
	}
	var resp map[string]json.RawMessage
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// SetSettings stores client-specific settings. Requires session token.
func (c *Client) SetSettings(ctx context.Context, settings map[string]string) error {
	req, err := c.request(ctx, http.MethodPost, "/sync/settings/set", settings)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// FetchUnreads fetches unread message state for all channels. Requires session token.
func (c *Client) FetchUnreads(ctx context.Context) ([]ChannelUnread, error) {
	req, err := c.request(ctx, http.MethodGet, "/sync/unreads", nil)
	if err != nil {
		return nil, err
	}
	var resp []ChannelUnread
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
