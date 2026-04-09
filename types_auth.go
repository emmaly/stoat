package stoat

import (
	"encoding/json"
	"fmt"
)

// AccountInfo represents a user's account information.
type AccountInfo struct {
	ID    string `json:"_id"`
	Email string `json:"email"`
}

// DataCreateAccount is the request body for creating an account.
type DataCreateAccount struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Invite   string `json:"invite,omitempty"`
	Captcha  string `json:"captcha,omitempty"`
}

// DataChangeEmail is the request body for changing email address.
type DataChangeEmail struct {
	Email           string `json:"email"`
	CurrentPassword string `json:"current_password"`
}

// DataChangePassword is the request body for changing password.
type DataChangePassword struct {
	Password        string `json:"password"`
	CurrentPassword string `json:"current_password"`
}

// DataSendPasswordReset is the request body for requesting a password reset email.
type DataSendPasswordReset struct {
	Email   string `json:"email"`
	Captcha string `json:"captcha,omitempty"`
}

// DataPasswordReset is the request body for executing a password reset.
type DataPasswordReset struct {
	Token          string `json:"token"`
	Password       string `json:"password"`
	RemoveSessions bool   `json:"remove_sessions,omitempty"`
}

// DataResendVerification is the request body for resending a verification email.
type DataResendVerification struct {
	Email   string `json:"email"`
	Captcha string `json:"captcha,omitempty"`
}

// DataAccountDeletion is the request body for confirming account deletion.
type DataAccountDeletion struct {
	Token string `json:"token"`
}

// DataLogin is the request body for logging in.
type DataLogin struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	FriendlyName string `json:"friendly_name,omitempty"`
}

// ResponseLogin is the interface for login response variants.
type ResponseLogin interface {
	loginResult() string
}

// LoginSuccess is returned when login succeeds.
type LoginSuccess struct {
	ID     string `json:"_id"`
	UserID string `json:"user_id"`
	Token  string `json:"token"`
	Name   string `json:"name"`
}

func (LoginSuccess) loginResult() string { return "Success" }

// LoginMFA is returned when MFA is required.
type LoginMFA struct {
	Ticket         string      `json:"ticket"`
	AllowedMethods []MFAMethod `json:"allowed_methods"`
}

func (LoginMFA) loginResult() string { return "MFA" }

// LoginDisabled is returned when the account is disabled.
type LoginDisabled struct {
	UserID string `json:"user_id"`
}

func (LoginDisabled) loginResult() string { return "Disabled" }

// RawResponseLogin is a wrapper that handles JSON unmarshalling of the tagged union.
type RawResponseLogin struct {
	Result ResponseLogin
}

// MarshalJSON implements json.Marshaler.
func (r RawResponseLogin) MarshalJSON() ([]byte, error) {
	switch v := r.Result.(type) {
	case *LoginSuccess:
		return json.Marshal(struct {
			Result string `json:"result"`
			*LoginSuccess
		}{Result: "Success", LoginSuccess: v})
	case *LoginMFA:
		return json.Marshal(struct {
			Result string `json:"result"`
			*LoginMFA
		}{Result: "MFA", LoginMFA: v})
	case *LoginDisabled:
		return json.Marshal(struct {
			Result string `json:"result"`
			*LoginDisabled
		}{Result: "Disabled", LoginDisabled: v})
	default:
		return nil, fmt.Errorf("unknown ResponseLogin variant: %T", r.Result)
	}
}

// UnmarshalJSON implements json.Unmarshaler for the tagged union.
func (r *RawResponseLogin) UnmarshalJSON(data []byte) error {
	var discriminator struct {
		Result string `json:"result"`
	}
	if err := json.Unmarshal(data, &discriminator); err != nil {
		return err
	}

	switch discriminator.Result {
	case "Success":
		var v LoginSuccess
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Result = &v
	case "MFA":
		var v LoginMFA
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Result = &v
	case "Disabled":
		var v LoginDisabled
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		r.Result = &v
	default:
		return fmt.Errorf("unknown login result: %q", discriminator.Result)
	}
	return nil
}

// SessionInfo represents a session.
type SessionInfo struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

// DataEditSession is the request body for editing a session.
type DataEditSession struct {
	FriendlyName string `json:"friendly_name"`
}

// MultiFactorStatus represents the MFA status of an account.
type MultiFactorStatus struct {
	EmailOTP       bool `json:"email_otp"`
	TrustedHandover bool `json:"trusted_handover"`
	EmailMFA       bool `json:"email_mfa"`
	TOTPMFA        bool `json:"totp_mfa"`
	SecurityKeyMFA bool `json:"security_key_mfa"`
	RecoveryActive bool `json:"recovery_active"`
}

// MFAMethod is a string type representing an MFA method.
type MFAMethod string

const (
	MFAMethodPassword MFAMethod = "Password"
	MFAMethodTOTP     MFAMethod = "Totp"
	MFAMethodRecovery MFAMethod = "Recovery"
)

// MFAResponse is the request body for MFA verification.
// Provide exactly one field.
type MFAResponse struct {
	Password     string `json:"password,omitempty"`
	TOTPCode     string `json:"totp_code,omitempty"`
	RecoveryCode string `json:"recovery_code,omitempty"`
}

// MFATicket represents a multi-factor auth ticket.
type MFATicket struct {
	ID           string `json:"_id"`
	AccountID    string `json:"account_id"`
	Token        string `json:"token"`
	Validated    bool   `json:"validated"`
	Authorised   bool   `json:"authorised"`
	LastTOTPCode string `json:"last_totp_code,omitempty"`
}

// ResponseTotpSecret is the response when generating a TOTP secret.
type ResponseTotpSecret struct {
	Secret string `json:"secret"`
}

// DataOnboard is the request body for completing onboarding.
type DataOnboard struct {
	Username string `json:"username"`
}

// DataHello is the response for checking onboarding status.
type DataHello struct {
	Onboarding bool `json:"onboarding"`
}
