package stoat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateWebhook creates a webhook for a channel. Requires session token.
func (c *Client) CreateWebhook(ctx context.Context, channelID string, data CreateWebhookBody) (*Webhook, error) {
	req, err := c.request(ctx, http.MethodPost, fmt.Sprintf("/channels/%s/webhooks", channelID), data)
	if err != nil {
		return nil, err
	}
	var wh Webhook
	if err := c.do(req, &wh); err != nil {
		return nil, err
	}
	return &wh, nil
}

// FetchChannelWebhooks fetches all webhooks for a channel. Requires session token.
func (c *Client) FetchChannelWebhooks(ctx context.Context, channelID string) ([]Webhook, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/channels/%s/webhooks", channelID), nil)
	if err != nil {
		return nil, err
	}
	var whs []Webhook
	if err := c.do(req, &whs); err != nil {
		return nil, err
	}
	return whs, nil
}

// FetchWebhook fetches a webhook by ID. Requires session token.
func (c *Client) FetchWebhook(ctx context.Context, webhookID string) (*Webhook, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/webhooks/%s", webhookID), nil)
	if err != nil {
		return nil, err
	}
	var wh Webhook
	if err := c.do(req, &wh); err != nil {
		return nil, err
	}
	return &wh, nil
}

// FetchWebhookWithToken fetches a webhook using its token. No auth required.
func (c *Client) FetchWebhookWithToken(ctx context.Context, webhookID, token string) (*Webhook, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/webhooks/%s/%s", webhookID, token), nil)
	if err != nil {
		return nil, err
	}
	var wh Webhook
	if err := c.do(req, &wh); err != nil {
		return nil, err
	}
	return &wh, nil
}

// EditWebhook edits a webhook. Requires session token.
func (c *Client) EditWebhook(ctx context.Context, webhookID string, data DataEditWebhook) (*Webhook, error) {
	req, err := c.request(ctx, http.MethodPatch, fmt.Sprintf("/webhooks/%s", webhookID), data)
	if err != nil {
		return nil, err
	}
	var wh Webhook
	if err := c.do(req, &wh); err != nil {
		return nil, err
	}
	return &wh, nil
}

// EditWebhookWithToken edits a webhook using its token. No auth required.
func (c *Client) EditWebhookWithToken(ctx context.Context, webhookID, token string, data DataEditWebhook) (*Webhook, error) {
	req, err := c.request(ctx, http.MethodPatch, fmt.Sprintf("/webhooks/%s/%s", webhookID, token), data)
	if err != nil {
		return nil, err
	}
	var wh Webhook
	if err := c.do(req, &wh); err != nil {
		return nil, err
	}
	return &wh, nil
}

// DeleteWebhook deletes a webhook. Requires session token.
func (c *Client) DeleteWebhook(ctx context.Context, webhookID string) error {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/webhooks/%s", webhookID), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// DeleteWebhookWithToken deletes a webhook using its token. No auth required.
func (c *Client) DeleteWebhookWithToken(ctx context.Context, webhookID, token string) error {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/webhooks/%s/%s", webhookID, token), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// ExecuteWebhook sends a message via webhook. No auth required.
func (c *Client) ExecuteWebhook(ctx context.Context, webhookID, token string, data DataMessageSend) (*Message, error) {
	req, err := c.request(ctx, http.MethodPost, fmt.Sprintf("/webhooks/%s/%s", webhookID, token), data)
	if err != nil {
		return nil, err
	}
	var msg Message
	if err := c.do(req, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

// ExecuteGitHubWebhook sends a GitHub webhook payload. No auth required.
func (c *Client) ExecuteGitHubWebhook(ctx context.Context, webhookID, token string, payload json.RawMessage) error {
	req, err := c.request(ctx, http.MethodPost, fmt.Sprintf("/webhooks/%s/%s/github", webhookID, token), payload)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
