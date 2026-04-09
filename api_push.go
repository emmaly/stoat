package stoat

import (
	"context"
	"net/http"
)

// PushSubscribe registers a web push subscription for the current session. Requires session token.
func (c *Client) PushSubscribe(ctx context.Context, sub WebPushSubscription) error {
	req, err := c.request(ctx, http.MethodPost, "/push/subscribe", sub)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// PushUnsubscribe removes the web push subscription for the current session. Requires session token.
func (c *Client) PushUnsubscribe(ctx context.Context) error {
	req, err := c.request(ctx, http.MethodPost, "/push/unsubscribe", nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
