# Session

Session management endpoints for login, logout, and session enumeration.

## Login

```
POST /auth/session/login
Auth: None
Body: DataLogin
Response 200: ResponseLogin
```

### DataLogin

The login body is a tagged union. Two variants:

**Email + Password:**
```json
{
  "email": "string",
  "password": "string",
  "friendly_name": "string"  // optional — device/session name
}
```

### ResponseLogin

The response is a tagged union with these variants:

**Success:**
```json
{
  "result": "Success",
  "_id": "<session_id>",
  "user_id": "<user_id>",
  "token": "<session_token>",
  "name": "<friendly_name>"
}
```

**MFA Required:**
```json
{
  "result": "MFA",
  "ticket": "<unvalidated_mfa_ticket>",
  "allowed_methods": ["Password", "Totp", "Recovery"]
}
```

**Account Disabled:**
```json
{
  "result": "Disabled",
  "user_id": "<user_id>"
}
```

## Logout

```
POST /auth/session/logout
Auth: Session Token
Response 204: Success
```

## Fetch Sessions

List all active sessions for the authenticated user.

```
GET /auth/session/all
Auth: Session Token
Response 200: SessionInfo[]
```

### SessionInfo

```json
{
  "_id": "string",
  "name": "string"     // friendly name
}
```

## Revoke Session

Delete a specific session by ID.

```
DELETE /auth/session/{id}
Auth: Session Token
Path: id (string, required)
Response 204: Success
```

## Edit Session

Update a session's friendly name.

```
PATCH /auth/session/{id}
Auth: Session Token
Path: id (string, required)
Body: DataEditSession
Response 200: SessionInfo
```

### DataEditSession

```json
{
  "friendly_name": "string"  // required
}
```

## Delete All Sessions

Revoke all sessions for the authenticated user.

```
DELETE /auth/session/all
Auth: Session Token
Query: revoke_self (boolean, optional) — whether to also revoke current session
Response 204: Success
```
