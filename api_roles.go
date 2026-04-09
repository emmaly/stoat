package stoat

import (
	"context"
	"fmt"
	"net/http"
)

// CreateRole creates a new role in a server. Requires session token.
func (c *Client) CreateRole(ctx context.Context, serverID string, data DataCreateRole) (*NewRoleResponse, error) {
	req, err := c.request(ctx, http.MethodPost, fmt.Sprintf("/servers/%s/roles", serverID), data)
	if err != nil {
		return nil, err
	}
	var resp NewRoleResponse
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchRole fetches a role by ID. Requires session token.
func (c *Client) FetchRole(ctx context.Context, serverID, roleID string) (*Role, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/servers/%s/roles/%s", serverID, roleID), nil)
	if err != nil {
		return nil, err
	}
	var role Role
	if err := c.do(req, &role); err != nil {
		return nil, err
	}
	return &role, nil
}

// EditRole edits a role. Requires session token.
func (c *Client) EditRole(ctx context.Context, serverID, roleID string, data DataEditRole) (*Role, error) {
	req, err := c.request(ctx, http.MethodPatch, fmt.Sprintf("/servers/%s/roles/%s", serverID, roleID), data)
	if err != nil {
		return nil, err
	}
	var role Role
	if err := c.do(req, &role); err != nil {
		return nil, err
	}
	return &role, nil
}

// DeleteRole deletes a role. Requires session token.
func (c *Client) DeleteRole(ctx context.Context, serverID, roleID string) error {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/servers/%s/roles/%s", serverID, roleID), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// EditRoleRanks reorders role positions. Requires session token.
func (c *Client) EditRoleRanks(ctx context.Context, serverID string, data DataEditRoleRanks) (*Server, error) {
	req, err := c.request(ctx, http.MethodPatch, fmt.Sprintf("/servers/%s/roles/ranks", serverID), data)
	if err != nil {
		return nil, err
	}
	var s Server
	if err := c.do(req, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// SetDefaultServerPermissions sets the default permission value for all members. Requires session token.
func (c *Client) SetDefaultServerPermissions(ctx context.Context, serverID string, data DataPermissionsValue) (*Server, error) {
	req, err := c.request(ctx, http.MethodPut, fmt.Sprintf("/servers/%s/permissions/default", serverID), data)
	if err != nil {
		return nil, err
	}
	var s Server
	if err := c.do(req, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// SetRoleServerPermission sets a role's permission overrides on a server. Requires session token.
func (c *Client) SetRoleServerPermission(ctx context.Context, serverID, roleID string, data DataSetServerRolePermission) (*Server, error) {
	req, err := c.request(ctx, http.MethodPut, fmt.Sprintf("/servers/%s/permissions/%s", serverID, roleID), data)
	if err != nil {
		return nil, err
	}
	var s Server
	if err := c.do(req, &s); err != nil {
		return nil, err
	}
	return &s, nil
}
