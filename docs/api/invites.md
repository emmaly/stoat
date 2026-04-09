# Invites

Server invite management.

## Create Invite

Create an invite for a channel (which invites to the parent server).

```
POST /channels/{target}/invites
Auth: Session Token
Path: target (string, required) — channel ID
Response 200: Invite
```

## Fetch Invite

Get public information about an invite (no auth required).

```
GET /invites/{target}
Auth: None
Path: target (string, required) — invite code
Response 200: InviteResponse
```

### InviteResponse

Tagged union:

**Server invite:**
```json
{
  "type": "Server",
  "server_id": "string",
  "server_name": "string",
  "server_icon": { /* File */ },   // nullable
  "server_banner": { /* File */ }, // nullable
  "server_flags": 0,
  "channel_id": "string",
  "channel_name": "string",
  "channel_description": "string", // nullable
  "user_name": "string",
  "user_avatar": { /* File */ },   // nullable
  "member_count": 0,
  "online_count": 0                // nullable
}
```

## Join Invite

Accept an invite and join the server.

```
POST /invites/{target}
Auth: Session Token
Path: target (string, required) — invite code
Response 200: InviteJoinResponse
```

### InviteJoinResponse

Tagged union:

**Server:**
```json
{
  "type": "Server",
  "channels": [Channel, ...],
  "server": { /* Server */ }
}
```

## Delete Invite

```
DELETE /invites/{target}
Auth: Session Token
Path: target (string, required) — invite code
Response 204: Success
```

## Invite Type

```json
{
  "_id": "string",       // invite code
  "server": "string",    // server ID
  "creator": "string",   // user ID who created the invite
  "channel": "string"    // channel ID the invite points to
}
```
