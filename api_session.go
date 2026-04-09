package stoat

import (
	"context"
	"fmt"
	"net/http"
)

// Login authenticates with email and password. No authentication required.
// Returns a ResponseLogin tagged union (LoginSuccess, LoginMFA, or LoginDisabled).
func (c *Client) Login(ctx context.Context, data DataLogin) (ResponseLogin, error) {
	req, err := c.request(ctx, http.MethodPost, "/auth/session/login", data)
	if err != nil {
		return nil, err
	}

	var raw RawResponseLogin
	if err := c.do(req, &raw); err != nil {
		return nil, err
	}
	return raw.Result, nil
}

// Logout ends the current session. Requires session token.
func (c *Client) Logout(ctx context.Context) error {
	req, err := c.request(ctx, http.MethodPost, "/auth/session/logout", nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// FetchSessions lists all active sessions. Requires session token.
func (c *Client) FetchSessions(ctx context.Context) ([]SessionInfo, error) {
	req, err := c.request(ctx, http.MethodGet, "/auth/session/all", nil)
	if err != nil {
		return nil, err
	}

	var sessions []SessionInfo
	if err := c.do(req, &sessions); err != nil {
		return nil, err
	}
	return sessions, nil
}

// RevokeSession deletes a specific session by ID. Requires session token.
func (c *Client) RevokeSession(ctx context.Context, id string) error {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/auth/session/%s", id), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// EditSession updates a session's friendly name. Requires session token.
func (c *Client) EditSession(ctx context.Context, id string, data DataEditSession) (*SessionInfo, error) {
	req, err := c.request(ctx, http.MethodPatch, fmt.Sprintf("/auth/session/%s", id), data)
	if err != nil {
		return nil, err
	}

	var info SessionInfo
	if err := c.do(req, &info); err != nil {
		return nil, err
	}
	return &info, nil
}

// DeleteAllSessions revokes all sessions. Requires session token.
// If revokeSelf is true, the current session is also revoked.
func (c *Client) DeleteAllSessions(ctx context.Context, revokeSelf bool) error {
	path := "/auth/session/all"
	if revokeSelf {
		path += "?revoke_self=true"
	}
	req, err := c.request(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
