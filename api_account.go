package stoat

import (
	"context"
	"fmt"
	"net/http"
)

// CreateAccount registers a new account. No authentication required.
func (c *Client) CreateAccount(ctx context.Context, data DataCreateAccount) error {
	req, err := c.request(ctx, http.MethodPost, "/auth/account/create", data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// FetchAccount returns the authenticated user's account info. Requires session token.
func (c *Client) FetchAccount(ctx context.Context) (*AccountInfo, error) {
	req, err := c.request(ctx, http.MethodGet, "/auth/account/", nil)
	if err != nil {
		return nil, err
	}
	var info AccountInfo
	if err := c.do(req, &info); err != nil {
		return nil, err
	}
	return &info, nil
}

// ChangeEmail changes the account email. Requires session token.
func (c *Client) ChangeEmail(ctx context.Context, data DataChangeEmail) error {
	req, err := c.request(ctx, http.MethodPatch, "/auth/account/change/email", data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// ChangePassword changes the account password. Requires session token.
func (c *Client) ChangePassword(ctx context.Context, data DataChangePassword) error {
	req, err := c.request(ctx, http.MethodPatch, "/auth/account/change/password", data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// SendPasswordReset sends a password reset email. No authentication required.
func (c *Client) SendPasswordReset(ctx context.Context, data DataSendPasswordReset) error {
	req, err := c.request(ctx, http.MethodPost, "/auth/account/reset_password", data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// PasswordReset executes a password reset using a token. No authentication required.
func (c *Client) PasswordReset(ctx context.Context, data DataPasswordReset) error {
	req, err := c.request(ctx, http.MethodPatch, "/auth/account/reset_password", data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// VerifyEmail verifies an email address using a code. No authentication required.
func (c *Client) VerifyEmail(ctx context.Context, code string) error {
	req, err := c.request(ctx, http.MethodPost, fmt.Sprintf("/auth/account/verify/%s", code), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// ResendVerification resends a verification email. No authentication required.
func (c *Client) ResendVerification(ctx context.Context, data DataResendVerification) error {
	req, err := c.request(ctx, http.MethodPost, "/auth/account/reverify", data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// DeleteAccount initiates account deletion. Requires session token and valid MFA ticket.
func (c *Client) DeleteAccount(ctx context.Context) error {
	req, err := c.request(ctx, http.MethodPost, "/auth/account/delete", nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// ConfirmDeletion confirms account deletion using a token. No authentication required.
func (c *Client) ConfirmDeletion(ctx context.Context, data DataAccountDeletion) error {
	req, err := c.request(ctx, http.MethodPut, "/auth/account/delete", data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// DisableAccount disables the account. Requires session token and valid MFA ticket.
func (c *Client) DisableAccount(ctx context.Context) error {
	req, err := c.request(ctx, http.MethodPost, "/auth/account/disable", nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
