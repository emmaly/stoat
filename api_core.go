package stoat

import (
	"context"
	"net/http"
)

// QueryNode fetches the server configuration. No authentication required.
func (c *Client) QueryNode(ctx context.Context) (*RevoltConfig, error) {
	req, err := c.request(ctx, http.MethodGet, "/", nil)
	if err != nil {
		return nil, err
	}

	var config RevoltConfig
	if err := c.do(req, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
