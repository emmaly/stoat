package stoat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchMembers fetches all members of a server. Requires session token.
func (c *Client) FetchMembers(ctx context.Context, serverID string, excludeOffline bool) (*AllMemberResponse, error) {
	path := fmt.Sprintf("/servers/%s/members", serverID)
	if excludeOffline {
		path += "?exclude_offline=true"
	}
	req, err := c.request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var resp AllMemberResponse
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchMember fetches a specific member. When includeRoles is true, the response
// includes role information. Requires session token.
func (c *Client) FetchMember(ctx context.Context, serverID, memberID string, includeRoles bool) (*MemberResponse, error) {
	path := fmt.Sprintf("/servers/%s/members/%s", serverID, memberID)
	if includeRoles {
		path += "?roles=true"
	}
	req, err := c.request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var resp MemberResponse
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EditMember edits a server member. Requires session token.
func (c *Client) EditMember(ctx context.Context, serverID, memberID string, data DataMemberEdit) (*Member, error) {
	req, err := c.request(ctx, http.MethodPatch, fmt.Sprintf("/servers/%s/members/%s", serverID, memberID), data)
	if err != nil {
		return nil, err
	}
	var m Member
	if err := c.do(req, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// KickMember kicks a member from a server. Requires session token.
func (c *Client) KickMember(ctx context.Context, serverID, memberID string) error {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/servers/%s/members/%s", serverID, memberID), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// QueryMembers searches for members by name (experimental). Requires session token.
func (c *Client) QueryMembers(ctx context.Context, serverID, query string) (*MemberQueryResponse, error) {
	path := fmt.Sprintf("/servers/%s/members_experimental_query?query=%s&experimental_api=true", serverID, query)
	req, err := c.request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var resp MemberQueryResponse
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BanUser bans a user from a server. Requires session token.
func (c *Client) BanUser(ctx context.Context, serverID, targetID string, data DataBanCreate) (*ServerBan, error) {
	req, err := c.request(ctx, http.MethodPut, fmt.Sprintf("/servers/%s/bans/%s", serverID, targetID), data)
	if err != nil {
		return nil, err
	}
	var ban ServerBan
	if err := c.do(req, &ban); err != nil {
		return nil, err
	}
	return &ban, nil
}

// UnbanUser unbans a user from a server. Requires session token.
func (c *Client) UnbanUser(ctx context.Context, serverID, targetID string) error {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/servers/%s/bans/%s", serverID, targetID), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// FetchBans fetches all bans for a server. Requires session token.
func (c *Client) FetchBans(ctx context.Context, serverID string) (*BanListResult, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/servers/%s/bans", serverID), nil)
	if err != nil {
		return nil, err
	}
	var resp BanListResult
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchServerInvites fetches all invites for a server. Requires session token.
// TODO: Use proper Invite type when available (Phase 6).
func (c *Client) FetchServerInvites(ctx context.Context, serverID string) ([]json.RawMessage, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/servers/%s/invites", serverID), nil)
	if err != nil {
		return nil, err
	}
	var resp []json.RawMessage
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
