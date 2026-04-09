package stoat

import (
	"context"
	"fmt"
	"net/http"
)

// DataJoinCall is the request body for joining a voice call.
type DataJoinCall struct {
	Node             *string  `json:"node,omitempty"`
	ForceDisconnect  *bool    `json:"force_disconnect,omitempty"`
	Recipients       []string `json:"recipients,omitempty"`
}

// CreateVoiceUserResponse is the response from joining a voice call.
type CreateVoiceUserResponse struct {
	Token string `json:"token"`
	URL   string `json:"url"`
}

// JoinCall joins a voice channel or starts a call. Requires session token.
func (c *Client) JoinCall(ctx context.Context, channelID string) (*CreateVoiceUserResponse, error) {
	req, err := c.request(ctx, http.MethodPost, fmt.Sprintf("/channels/%s/join_call", channelID), DataJoinCall{})
	if err != nil {
		return nil, err
	}
	var resp CreateVoiceUserResponse
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// StopRing stops ringing a specific user in a DM/group call. Requires session token.
func (c *Client) StopRing(ctx context.Context, channelID, userID string) error {
	req, err := c.request(ctx, http.MethodPut, fmt.Sprintf("/channels/%s/end_ring/%s", channelID, userID), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
