package stoat

import (
	"context"
	"fmt"
	"net/http"
)

// doChannel executes a request and decodes the response as a Channel via RawChannel.
func (c *Client) doChannel(req *http.Request) (Channel, error) {
	var rc RawChannel
	if err := c.do(req, &rc); err != nil {
		return nil, err
	}
	return rc.Value, nil
}

// FetchChannel returns a channel by ID. Requires session token.
func (c *Client) FetchChannel(ctx context.Context, channelID string) (Channel, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/channels/%s", channelID), nil)
	if err != nil {
		return nil, err
	}
	return c.doChannel(req)
}

// EditChannel edits a channel. Requires session token.
func (c *Client) EditChannel(ctx context.Context, channelID string, data DataEditChannel) (Channel, error) {
	req, err := c.request(ctx, http.MethodPatch, fmt.Sprintf("/channels/%s", channelID), data)
	if err != nil {
		return nil, err
	}
	return c.doChannel(req)
}

// CloseChannel closes a DM, leaves a group, or deletes a server channel. Requires session token.
func (c *Client) CloseChannel(ctx context.Context, channelID string, leaveSilently bool) error {
	path := fmt.Sprintf("/channels/%s", channelID)
	if leaveSilently {
		path += "?leave_silently=true"
	}
	req, err := c.request(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// SetDefaultChannelPermissions sets default permission overrides for a channel. Requires session token.
func (c *Client) SetDefaultChannelPermissions(ctx context.Context, channelID string, data DataDefaultChannelPermissions) (Channel, error) {
	req, err := c.request(ctx, http.MethodPut, fmt.Sprintf("/channels/%s/permissions/default", channelID), data)
	if err != nil {
		return nil, err
	}
	return c.doChannel(req)
}

// SetRoleChannelPermissions sets role permissions on a channel. Requires session token.
func (c *Client) SetRoleChannelPermissions(ctx context.Context, channelID string, roleID string, data DataSetRolePermissions) (Channel, error) {
	req, err := c.request(ctx, http.MethodPut, fmt.Sprintf("/channels/%s/permissions/%s", channelID, roleID), data)
	if err != nil {
		return nil, err
	}
	return c.doChannel(req)
}
