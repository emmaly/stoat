# Multi-Factor Authentication (MFA)

## MFA Status

Check whether the authenticated user has MFA enabled.

```
GET /auth/mfa/
Auth: Session Token
Response 200: MultiFactorStatus
```

### MultiFactorStatus

```json
{
  "email_otp": false,
  "trusted_handover": false,
  "email_mfa": false,
  "totp_mfa": false,
  "security_key_mfa": false,
  "recovery_active": false
}
```

## Get MFA Methods

List available MFA methods for the authenticated user.

```
GET /auth/mfa/methods
Auth: Session Token
Response 200: MFAMethod[]
```

### MFAMethod

String enum: `"Password"`, `"Totp"`, `"Recovery"`

## Create MFA Ticket

Create a validated MFA ticket by providing an MFA response.

```
PUT /auth/mfa/ticket
Auth: Session Token OR Unvalidated MFA Ticket
Body: MFAResponse
Response 200: MFATicket
```

### MFAResponse

Tagged union — provide one of:

```json
{"password": "string"}
```
```json
{"totp_code": "string"}
```
```json
{"recovery_code": "string"}
```

### MFATicket

```json
{
  "_id": "string",
  "account_id": "string",
  "token": "string",
  "validated": true,
  "authorised": true,
  "last_totp_code": "string"
}
```

## Generate TOTP Secret

Generate a new TOTP secret for setting up 2FA.

```
POST /auth/mfa/totp
Auth: Valid MFA Ticket + Session Token
Response 200: ResponseTotpSecret
```

### ResponseTotpSecret

```json
{
  "secret": "string"  // Base32-encoded TOTP secret
}
```

## Enable TOTP 2FA

Enable TOTP by providing the generated code to prove the secret was saved.

```
PUT /auth/mfa/totp
Auth: Session Token
Body: MFAResponse (with totp_code)
Response 204: Success
```

## Disable TOTP 2FA

```
DELETE /auth/mfa/totp
Auth: Valid MFA Ticket + Session Token
Response 204: Success
```

## Fetch Recovery Codes

```
POST /auth/mfa/recovery
Auth: Valid MFA Ticket + Session Token
Response 200: string[] (array of recovery codes)
```

## Generate Recovery Codes

Generate new recovery codes (invalidates old ones).

```
PATCH /auth/mfa/recovery
Auth: Valid MFA Ticket + Session Token
Response 200: string[] (array of new recovery codes)
```
