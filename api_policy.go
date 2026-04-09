package stoat

import (
	"context"
	"net/http"
)

// AcknowledgePolicy acknowledges updated terms of service or privacy policy. Requires session token.
func (c *Client) AcknowledgePolicy(ctx context.Context) error {
	req, err := c.request(ctx, http.MethodPost, "/policy/acknowledge", nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
