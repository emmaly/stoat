package stoat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchDMs returns all DM channels for the authenticated user. Requires session token.
func (c *Client) FetchDMs(ctx context.Context) ([]Channel, error) {
	req, err := c.request(ctx, http.MethodGet, "/users/dms", nil)
	if err != nil {
		return nil, err
	}
	var rawChannels []json.RawMessage
	if err := c.do(req, &rawChannels); err != nil {
		return nil, err
	}
	channels := make([]Channel, 0, len(rawChannels))
	for _, raw := range rawChannels {
		var rc RawChannel
		if err := json.Unmarshal(raw, &rc); err != nil {
			return nil, fmt.Errorf("unmarshal channel: %w", err)
		}
		channels = append(channels, rc.Value)
	}
	return channels, nil
}

// OpenDM opens or retrieves an existing DM channel with a user. Requires session token.
func (c *Client) OpenDM(ctx context.Context, targetID string) (Channel, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/users/%s/dm", targetID), nil)
	if err != nil {
		return nil, err
	}
	var rc RawChannel
	if err := c.do(req, &rc); err != nil {
		return nil, err
	}
	return rc.Value, nil
}
