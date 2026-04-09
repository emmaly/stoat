package stoat

import (
	"context"
	"fmt"
	"net/http"
)

// CreateGroup creates a new group channel. Requires session token.
func (c *Client) CreateGroup(ctx context.Context, data DataCreateGroup) (Channel, error) {
	req, err := c.request(ctx, http.MethodPost, "/channels/create", data)
	if err != nil {
		return nil, err
	}
	return c.doChannel(req)
}

// FetchGroupMembers lists all members of a group channel. Requires session token.
func (c *Client) FetchGroupMembers(ctx context.Context, channelID string) ([]User, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/channels/%s/members", channelID), nil)
	if err != nil {
		return nil, err
	}
	var users []User
	if err := c.do(req, &users); err != nil {
		return nil, err
	}
	return users, nil
}

// AddGroupMember adds a user to a group channel. Requires session token.
func (c *Client) AddGroupMember(ctx context.Context, groupID, memberID string) error {
	req, err := c.request(ctx, http.MethodPut, fmt.Sprintf("/channels/%s/recipients/%s", groupID, memberID), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// RemoveGroupMember removes a user from a group channel. Requires session token.
func (c *Client) RemoveGroupMember(ctx context.Context, groupID, memberID string) error {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/channels/%s/recipients/%s", groupID, memberID), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
