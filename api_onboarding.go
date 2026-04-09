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
// Requires session token.
//
// TODO: Update return type to *User when User type is added in Phase 3.
func (c *Client) CompleteOnboarding(ctx context.Context, data DataOnboard) error {
	req, err := c.request(ctx, http.MethodPost, "/onboard/complete", data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
