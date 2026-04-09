package stoat

import (
	"context"
	"fmt"
	"net/http"
)

// CreateServer creates a new server. Requires session token.
func (c *Client) CreateServer(ctx context.Context, data DataCreateServer) (*CreateServerLegacyResponse, error) {
	req, err := c.request(ctx, http.MethodPost, "/servers/create", data)
	if err != nil {
		return nil, err
	}
	var resp CreateServerLegacyResponse
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchServer fetches a server by ID. When includeChannels is true, the response
// includes full channel objects. Requires session token.
func (c *Client) FetchServer(ctx context.Context, serverID string, includeChannels bool) (*FetchServerResponse, error) {
	path := fmt.Sprintf("/servers/%s", serverID)
	if includeChannels {
		path += "?include_channels=true"
	}
	req, err := c.request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var resp FetchServerResponse
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EditServer edits a server. Requires session token and valid MFA ticket.
func (c *Client) EditServer(ctx context.Context, serverID string, data DataEditServer) (*Server, error) {
	req, err := c.request(ctx, http.MethodPatch, fmt.Sprintf("/servers/%s", serverID), data)
	if err != nil {
		return nil, err
	}
	var s Server
	if err := c.do(req, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// DeleteServer deletes or leaves a server. If you are the owner, this deletes
// the server; otherwise you leave. Requires session token.
func (c *Client) DeleteServer(ctx context.Context, serverID string, leaveSilently bool) error {
	path := fmt.Sprintf("/servers/%s", serverID)
	if leaveSilently {
		path += "?leave_silently=true"
	}
	req, err := c.request(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// MarkServerRead marks all channels in a server as read. Requires session token.
func (c *Client) MarkServerRead(ctx context.Context, serverID string) error {
	req, err := c.request(ctx, http.MethodPut, fmt.Sprintf("/servers/%s/ack", serverID), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// CreateServerChannel creates a channel in a server. Requires session token.
func (c *Client) CreateServerChannel(ctx context.Context, serverID string, data DataCreateServerChannel) (Channel, error) {
	req, err := c.request(ctx, http.MethodPost, fmt.Sprintf("/servers/%s/channels", serverID), data)
	if err != nil {
		return nil, err
	}
	return c.doChannel(req)
}
