# Authentication

The Stoat API uses API key authentication via HTTP headers. There are no OAuth2 flows — tokens are obtained through the login endpoint or created in the client UI.

## Token Types

| Type | Header | How to Obtain |
|------|--------|---------------|
| User Session Token | `X-Session-Token` | `POST /auth/session/login` or copy from client |
| Bot Token | `X-Bot-Token` | Created in user settings UI |
| MFA Ticket (validated) | `X-MFA-Ticket` | `PUT /auth/mfa/ticket` after providing TOTP/recovery code |
| MFA Ticket (unvalidated) | `X-MFA-Ticket` | Returned from login when MFA is required |

## Security Schemes (from OpenAPI spec)

```
Session Token:
  type: apiKey
  name: x-session-token
  in: header
  description: Used to authenticate as a user.

Valid MFA Ticket:
  type: apiKey
  name: x-mfa-ticket
  in: header
  description: Used to authorise a request (validated MFA ticket).

Unvalidated MFA Ticket:
  type: apiKey
  name: x-mfa-ticket
  in: header
  description: Used to authorise a request (unvalidated MFA ticket).
```

## Login Flow

### Basic Login (no MFA)

```
POST /auth/session/login
Body: { "email": "...", "password": "..." }
→ 200: ResponseLogin (contains session token)
```

### Login with MFA

```
POST /auth/session/login
Body: { "email": "...", "password": "..." }
→ 200: { "result": "MFA", "ticket": "<unvalidated-ticket>", ... }

PUT /auth/mfa/ticket
Header: X-MFA-Ticket: <unvalidated-ticket>
Body: { "totp_code": "123456" }  (or recovery code)
→ 200: MFATicket (validated)

POST /auth/session/login  (retry with validated ticket)
```

## Bot Authentication

Bot tokens are sent via the `X-Bot-Token` header. Bot tokens do not expire and do not require session management.

Bots authenticate to WebSocket connections by including the token in the connection URL query parameter (`?token=<bot-token>`) or by sending an `Authenticate` event after connecting.

## Endpoints Requiring Authentication

Nearly all endpoints require `Session Token` authentication. Exceptions:

| Endpoint | Auth Required | Notes |
|----------|---------------|-------|
| `GET /` | No | Returns server config |
| `POST /auth/account/create` | No | Account creation |
| `POST /auth/session/login` | No | Login |
| `POST /auth/account/verify/{code}` | No | Email verification |
| `POST /auth/account/reset_password` | No | Request password reset |
| `PATCH /auth/account/reset_password` | No | Execute password reset |
| `POST /auth/account/reverify` | No | Resend verification |
| `PUT /auth/account/delete` | No | Confirm deletion (uses token in body) |
| `GET /custom/emoji/{emoji_id}` | No | Public emoji fetch |
| `GET /invites/{target}` | No | Public invite info |
| `GET /users/{target}/default_avatar` | No | Default avatar image |

## MFA-Protected Endpoints

Some sensitive operations require both a session token AND a valid MFA ticket:

- `POST /auth/account/delete` — Delete account
- `POST /auth/account/disable` — Disable account
- `DELETE /auth/mfa/totp` — Disable TOTP
- `POST /auth/mfa/totp` — Generate TOTP secret
- `POST /auth/mfa/recovery` — Fetch recovery codes
- `PATCH /auth/mfa/recovery` — Regenerate recovery codes
- `PATCH /servers/{target}` — Edit server (requires MFA ticket)
