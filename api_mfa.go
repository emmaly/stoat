package stoat

import (
	"context"
	"net/http"
)

// MFAStatus checks whether the authenticated user has MFA enabled. Requires session token.
func (c *Client) MFAStatus(ctx context.Context) (*MultiFactorStatus, error) {
	req, err := c.request(ctx, http.MethodGet, "/auth/mfa/", nil)
	if err != nil {
		return nil, err
	}
	var status MultiFactorStatus
	if err := c.do(req, &status); err != nil {
		return nil, err
	}
	return &status, nil
}

// GetMFAMethods lists available MFA methods. Requires session token.
func (c *Client) GetMFAMethods(ctx context.Context) ([]MFAMethod, error) {
	req, err := c.request(ctx, http.MethodGet, "/auth/mfa/methods", nil)
	if err != nil {
		return nil, err
	}
	var methods []MFAMethod
	if err := c.do(req, &methods); err != nil {
		return nil, err
	}
	return methods, nil
}

// CreateMFATicket creates a validated MFA ticket. Requires session token or unvalidated MFA ticket.
func (c *Client) CreateMFATicket(ctx context.Context, data MFAResponse) (*MFATicket, error) {
	req, err := c.request(ctx, http.MethodPut, "/auth/mfa/ticket", data)
	if err != nil {
		return nil, err
	}
	var ticket MFATicket
	if err := c.do(req, &ticket); err != nil {
		return nil, err
	}
	return &ticket, nil
}

// GenerateTOTPSecret generates a new TOTP secret for 2FA setup. Requires session token and valid MFA ticket.
func (c *Client) GenerateTOTPSecret(ctx context.Context) (*ResponseTotpSecret, error) {
	req, err := c.request(ctx, http.MethodPost, "/auth/mfa/totp", nil)
	if err != nil {
		return nil, err
	}
	var resp ResponseTotpSecret
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableTOTP enables TOTP 2FA by providing a verification code. Requires session token.
func (c *Client) EnableTOTP(ctx context.Context, data MFAResponse) error {
	req, err := c.request(ctx, http.MethodPut, "/auth/mfa/totp", data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// DisableTOTP disables TOTP 2FA. Requires session token and valid MFA ticket.
func (c *Client) DisableTOTP(ctx context.Context) error {
	req, err := c.request(ctx, http.MethodDelete, "/auth/mfa/totp", nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// FetchRecoveryCodes fetches existing recovery codes. Requires session token and valid MFA ticket.
func (c *Client) FetchRecoveryCodes(ctx context.Context) ([]string, error) {
	req, err := c.request(ctx, http.MethodPost, "/auth/mfa/recovery", nil)
	if err != nil {
		return nil, err
	}
	var codes []string
	if err := c.do(req, &codes); err != nil {
		return nil, err
	}
	return codes, nil
}

// GenerateRecoveryCodes generates new recovery codes, invalidating old ones.
// Requires session token and valid MFA ticket.
func (c *Client) GenerateRecoveryCodes(ctx context.Context) ([]string, error) {
	req, err := c.request(ctx, http.MethodPatch, "/auth/mfa/recovery", nil)
	if err != nil {
		return nil, err
	}
	var codes []string
	if err := c.do(req, &codes); err != nil {
		return nil, err
	}
	return codes, nil
}
