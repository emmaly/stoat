package stoat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchDMs returns all DM channels for the authenticated user. Requires session token.
// TODO: return []Channel when Channel type is defined (Phase 4).
func (c *Client) FetchDMs(ctx context.Context) ([]json.RawMessage, error) {
	req, err := c.request(ctx, http.MethodGet, "/users/dms", nil)
	if err != nil {
		return nil, err
	}
	var channels []json.RawMessage
	if err := c.do(req, &channels); err != nil {
		return nil, err
	}
	return channels, nil
}

// OpenDM opens or retrieves an existing DM channel with a user. Requires session token.
// TODO: return Channel when Channel type is defined (Phase 4).
func (c *Client) OpenDM(ctx context.Context, targetID string) (json.RawMessage, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/users/%s/dm", targetID), nil)
	if err != nil {
		return nil, err
	}
	var channel json.RawMessage
	if err := c.do(req, &channel); err != nil {
		return nil, err
	}
	return channel, nil
}
