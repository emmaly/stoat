package stoat

import (
	"context"
	"fmt"
	"net/http"
)

// SendFriendRequest sends a friend request by username. Requires session token.
func (c *Client) SendFriendRequest(ctx context.Context, data DataSendFriendRequest) (*User, error) {
	req, err := c.request(ctx, http.MethodPost, "/users/friend", data)
	if err != nil {
		return nil, err
	}
	var user User
	if err := c.do(req, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// AcceptFriend accepts an incoming friend request. Requires session token.
func (c *Client) AcceptFriend(ctx context.Context, targetID string) (*User, error) {
	req, err := c.request(ctx, http.MethodPut, fmt.Sprintf("/users/%s/friend", targetID), nil)
	if err != nil {
		return nil, err
	}
	var user User
	if err := c.do(req, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// RemoveFriend removes a friend or denies a friend request. Requires session token.
func (c *Client) RemoveFriend(ctx context.Context, targetID string) (*User, error) {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/users/%s/friend", targetID), nil)
	if err != nil {
		return nil, err
	}
	var user User
	if err := c.do(req, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// BlockUser blocks a user. Requires session token.
func (c *Client) BlockUser(ctx context.Context, targetID string) (*User, error) {
	req, err := c.request(ctx, http.MethodPut, fmt.Sprintf("/users/%s/block", targetID), nil)
	if err != nil {
		return nil, err
	}
	var user User
	if err := c.do(req, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// UnblockUser unblocks a user. Requires session token.
func (c *Client) UnblockUser(ctx context.Context, targetID string) (*User, error) {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/users/%s/block", targetID), nil)
	if err != nil {
		return nil, err
	}
	var user User
	if err := c.do(req, &user); err != nil {
		return nil, err
	}
	return &user, nil
}
