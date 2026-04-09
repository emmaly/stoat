# Account

Account management endpoints for registration, email/password changes, verification, and deletion.

## Create Account

```
POST /auth/account/create
Auth: None
Body: DataCreateAccount
Response 204: Success
```

### DataCreateAccount

```json
{
  "email": "string",         // required
  "password": "string",      // required
  "invite": "string",        // optional — invite code (if instance is invite-only)
  "captcha": "string"        // optional — CAPTCHA verification token
}
```

## Fetch Account

```
GET /auth/account/
Auth: Session Token
Response 200: AccountInfo
```

### AccountInfo

```json
{
  "_id": "string",
  "email": "string"
}
```

## Change Email

```
PATCH /auth/account/change/email
Auth: Session Token
Body: DataChangeEmail
Response 204: Success
```

### DataChangeEmail

```json
{
  "email": "string",          // required — new email
  "current_password": "string" // required
}
```

## Change Password

```
PATCH /auth/account/change/password
Auth: Session Token
Body: DataChangePassword
Response 204: Success
```

### DataChangePassword

```json
{
  "password": "string",         // required — new password
  "current_password": "string"  // required
}
```

## Send Password Reset

Request a password reset email.

```
POST /auth/account/reset_password
Auth: None
Body: DataSendPasswordReset
Response 204: Success
```

### DataSendPasswordReset

```json
{
  "email": "string",    // required
  "captcha": "string"   // optional
}
```

## Password Reset

Execute a password reset using the code from email.

```
PATCH /auth/account/reset_password
Auth: None
Body: DataPasswordReset
Response 204: Success
```

### DataPasswordReset

```json
{
  "token": "string",     // required — reset token from email
  "password": "string",  // required — new password
  "remove_sessions": true // optional — revoke all sessions
}
```

## Verify Email

```
POST /auth/account/verify/{code}
Auth: None
Path: code (string, required) — verification code from email
Response 200: ResponseVerify
```

## Resend Verification

```
POST /auth/account/reverify
Auth: None
Body: DataResendVerification
Response 204: Success
```

### DataResendVerification

```json
{
  "email": "string",    // required
  "captcha": "string"   // optional
}
```

## Delete Account

Initiate account deletion (requires MFA).

```
POST /auth/account/delete
Auth: Valid MFA Ticket + Session Token
Response 204: Success
```

## Confirm Account Deletion

Confirm deletion using the token from email.

```
PUT /auth/account/delete
Auth: None
Body: DataAccountDeletion
Response 204: Success
```

### DataAccountDeletion

```json
{
  "token": "string"  // required — deletion confirmation token
}
```

## Disable Account

```
POST /auth/account/disable
Auth: Valid MFA Ticket + Session Token
Response 204: Success
```
