package stoat

import (
	"context"
	"fmt"
	"net/http"
)

// CreateInvite creates an invite for a channel. Requires session token.
func (c *Client) CreateInvite(ctx context.Context, channelID string) (*Invite, error) {
	req, err := c.request(ctx, http.MethodPost, fmt.Sprintf("/channels/%s/invites", channelID), nil)
	if err != nil {
		return nil, err
	}
	var inv Invite
	if err := c.do(req, &inv); err != nil {
		return nil, err
	}
	return &inv, nil
}

// FetchInvite fetches public information about an invite. No auth required.
func (c *Client) FetchInvite(ctx context.Context, inviteCode string) (InviteResponse, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/invites/%s", inviteCode), nil)
	if err != nil {
		return nil, err
	}
	var raw RawInviteResponse
	if err := c.do(req, &raw); err != nil {
		return nil, err
	}
	return raw.Value, nil
}

// JoinInvite accepts an invite and joins the server or group. Requires session token.
func (c *Client) JoinInvite(ctx context.Context, inviteCode string) (InviteJoinResponse, error) {
	req, err := c.request(ctx, http.MethodPost, fmt.Sprintf("/invites/%s", inviteCode), nil)
	if err != nil {
		return nil, err
	}
	var raw RawInviteJoinResponse
	if err := c.do(req, &raw); err != nil {
		return nil, err
	}
	return raw.Value, nil
}

// DeleteInvite deletes an invite. Requires session token.
func (c *Client) DeleteInvite(ctx context.Context, inviteCode string) error {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/invites/%s", inviteCode), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
