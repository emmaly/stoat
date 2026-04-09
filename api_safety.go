package stoat

import (
	"context"
	"net/http"
)

// ReportContent reports a message, server, or user. Requires session token.
func (c *Client) ReportContent(ctx context.Context, data DataReportContent) error {
	req, err := c.request(ctx, http.MethodPost, "/safety/report", data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
