package stoat

import (
	"context"
	"net/http"
)

// CheckOnboarding checks whether the authenticated user needs to complete onboarding.
// Requires session token.
func (c *Client) CheckOnboarding(ctx context.Context) (*DataHello, error) {
	req, err := c.request(ctx, http.MethodGet, "/onboard/hello", nil)
	if err != nil {
		return nil, err
	}
	var hello DataHello
	if err := c.do(req, &hello); err != nil {
		return nil, err
	}
	return &hello, nil
}

// CompleteOnboarding completes the onboarding process by choosing a username.
// Requires session token. Returns the created user.
func (c *Client) CompleteOnboarding(ctx context.Context, data DataOnboard) (*User, error) {
	req, err := c.request(ctx, http.MethodPost, "/onboard/complete", data)
	if err != nil {
		return nil, err
	}
	var user User
	if err := c.do(req, &user); err != nil {
		return nil, err
	}
	return &user, nil
}
