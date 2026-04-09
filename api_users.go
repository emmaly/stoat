package stoat

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// FetchSelf returns the currently authenticated user. Requires session token.
func (c *Client) FetchSelf(ctx context.Context) (*User, error) {
	req, err := c.request(ctx, http.MethodGet, "/users/@me", nil)
	if err != nil {
		return nil, err
	}
	var user User
	if err := c.do(req, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// FetchUser returns a user by ID. Requires session token.
func (c *Client) FetchUser(ctx context.Context, targetID string) (*User, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/users/%s", targetID), nil)
	if err != nil {
		return nil, err
	}
	var user User
	if err := c.do(req, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// EditUser edits a user's profile. Requires session token.
func (c *Client) EditUser(ctx context.Context, targetID string, data DataEditUser) (*User, error) {
	req, err := c.request(ctx, http.MethodPatch, fmt.Sprintf("/users/%s", targetID), data)
	if err != nil {
		return nil, err
	}
	var user User
	if err := c.do(req, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// ChangeUsername changes the authenticated user's username. Requires session token.
func (c *Client) ChangeUsername(ctx context.Context, data DataChangeUsername) (*User, error) {
	req, err := c.request(ctx, http.MethodPatch, "/users/@me/username", data)
	if err != nil {
		return nil, err
	}
	var user User
	if err := c.do(req, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// FetchUserProfile returns a user's full profile. Requires session token.
func (c *Client) FetchUserProfile(ctx context.Context, targetID string) (*UserProfile, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/users/%s/profile", targetID), nil)
	if err != nil {
		return nil, err
	}
	var profile UserProfile
	if err := c.do(req, &profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

// FetchDefaultAvatar returns a user's default generated avatar as PNG bytes.
// No authentication required.
func (c *Client) FetchDefaultAvatar(ctx context.Context, targetID string) ([]byte, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/users/%s/default_avatar", targetID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(data))
	}

	return data, nil
}

// FetchUserFlags returns a user's flags. Requires session token.
func (c *Client) FetchUserFlags(ctx context.Context, targetID string) (*FlagResponse, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/users/%s/flags", targetID), nil)
	if err != nil {
		return nil, err
	}
	var flags FlagResponse
	if err := c.do(req, &flags); err != nil {
		return nil, err
	}
	return &flags, nil
}

// FetchMutual returns mutual friends, servers, and channels with a user. Requires session token.
func (c *Client) FetchMutual(ctx context.Context, targetID string) (*MutualResponse, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/users/%s/mutual", targetID), nil)
	if err != nil {
		return nil, err
	}
	var mutual MutualResponse
	if err := c.do(req, &mutual); err != nil {
		return nil, err
	}
	return &mutual, nil
}
