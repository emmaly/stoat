# Server Members

Member management, bans, and server invites.

## Fetch Members

Get all members of a server.

```
GET /servers/{target}/members
Auth: Session Token
Path: target (string, required)
Query: exclude_offline (boolean, optional) — exclude offline members
Response 200: AllMemberResponse
```

### AllMemberResponse

```json
{
  "members": [Member, ...],
  "users": [User, ...]
}
```

## Fetch Member

Get a specific member.

```
GET /servers/{server_id}/members/{member_id}
Auth: Session Token
Path: server_id (string, required), member_id (string, required)
Query: roles (boolean, optional) — include role information
Response 200: MemberResponse
```

### MemberResponse

When `roles` is false:
```json
{ /* Member object */ }
```

When `roles` is true:
```json
{
  "member": { /* Member */ },
  "roles": [Role, ...]
}
```

## Edit Member

```
PATCH /servers/{server_id}/members/{member_id}
Auth: Session Token
Path: server_id (string, required), member_id (string, required)
Body: DataMemberEdit
Response 200: Member
```

### DataMemberEdit

```json
{
  "nickname": "string",       // optional
  "avatar": "string",         // optional — file ID
  "roles": ["role_id", ...],  // optional
  "timeout": "ISO8601",       // optional — timeout until timestamp
  "remove": ["Nickname", ...] // optional — fields to clear
}
```

Remove options (FieldsMember): `"Nickname"`, `"Avatar"`, `"Roles"`, `"Timeout"`

## Kick Member

```
DELETE /servers/{server_id}/members/{member_id}
Auth: Session Token
Path: server_id (string, required), member_id (string, required)
Response 204: Success
```

## Query Members by Name (Experimental)

```
GET /servers/{target}/members_experimental_query
Auth: Session Token
Path: target (string, required)
Query:
  - query (string, required) — search string
  - experimental_api (boolean, required) — must be true
Response 200: MemberQueryResponse
```

### MemberQueryResponse

```json
{
  "members": [Member, ...],
  "users": [User, ...]
}
```

## Ban User

```
PUT /servers/{server}/bans/{target}
Auth: Session Token
Path: server (string, required), target (string, required) — user ID to ban
Body: DataBanCreate
Response 200: ServerBan
```

### DataBanCreate

```json
{
  "reason": "string"  // optional — ban reason
}
```

### ServerBan

```json
{
  "_id": {
    "server": "string",
    "user": "string"
  },
  "reason": "string"  // nullable
}
```

## Unban User

```
DELETE /servers/{server}/bans/{target}
Auth: Session Token
Path: server (string, required), target (string, required)
Response 204: Success
```

## Fetch Bans

```
GET /servers/{target}/bans
Auth: Session Token
Path: target (string, required)
Response 200: BanListResult
```

### BanListResult

```json
{
  "users": [BannedUser, ...],
  "bans": [ServerBan, ...]
}
```

### BannedUser

```json
{
  "_id": "string",
  "username": "string",
  "discriminator": "string",
  "avatar": { /* File */ }  // nullable
}
```

## Fetch Invites

List all invites for a server.

```
GET /servers/{target}/invites
Auth: Session Token
Path: target (string, required)
Response 200: Invite[]
```

## Member Type

```json
{
  "_id": {
    "server": "string",
    "user": "string"
  },
  "joined_at": "ISO8601",
  "nickname": "string",          // nullable
  "avatar": { /* File */ },      // nullable
  "roles": ["role_id", ...],
  "timeout": "ISO8601"           // nullable — timed out until
}
```

### MemberCompositeKey

```json
{
  "server": "string",
  "user": "string"
}
```
