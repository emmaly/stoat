package stoat

import (
	"context"
	"fmt"
	"net/http"
)

// CreateBot creates a new bot. Requires session token.
func (c *Client) CreateBot(ctx context.Context, data DataCreateBot) (*BotWithUserResponse, error) {
	req, err := c.request(ctx, http.MethodPost, "/bots/create", data)
	if err != nil {
		return nil, err
	}
	var resp BotWithUserResponse
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchBot fetches a bot by ID. Requires session token.
func (c *Client) FetchBot(ctx context.Context, botID string) (*FetchBotResponse, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/bots/%s", botID), nil)
	if err != nil {
		return nil, err
	}
	var resp FetchBotResponse
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EditBot edits a bot. Requires session token.
func (c *Client) EditBot(ctx context.Context, botID string, data DataEditBot) (*BotWithUserResponse, error) {
	req, err := c.request(ctx, http.MethodPatch, fmt.Sprintf("/bots/%s", botID), data)
	if err != nil {
		return nil, err
	}
	var resp BotWithUserResponse
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBot deletes a bot. Requires session token.
func (c *Client) DeleteBot(ctx context.Context, botID string) error {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/bots/%s", botID), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// FetchOwnedBots fetches all bots owned by the authenticated user. Requires session token.
func (c *Client) FetchOwnedBots(ctx context.Context) (*OwnedBotsResponse, error) {
	req, err := c.request(ctx, http.MethodGet, "/bots/@me", nil)
	if err != nil {
		return nil, err
	}
	var resp OwnedBotsResponse
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchPublicBot fetches public information about a bot. Requires session token.
func (c *Client) FetchPublicBot(ctx context.Context, botID string) (*PublicBot, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/bots/%s/invite", botID), nil)
	if err != nil {
		return nil, err
	}
	var resp PublicBot
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// InviteBot invites a bot to a server or group. Requires session token.
func (c *Client) InviteBot(ctx context.Context, botID string, dest InviteBotDestination) error {
	req, err := c.request(ctx, http.MethodPost, fmt.Sprintf("/bots/%s/invite", botID), dest)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
