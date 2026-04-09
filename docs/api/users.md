# User Information

## Fetch Self

Get the currently authenticated user.

```
GET /users/@me
Auth: Session Token
Response 200: User
```

## Fetch User

Get a user by ID.

```
GET /users/{target}
Auth: Session Token
Path: target (string, required) — user ID
Response 200: User
```

## Edit User

Edit the authenticated user's profile.

```
PATCH /users/{target}
Auth: Session Token
Path: target (string, required) — user ID
Body: DataEditUser
Response 200: User
```

### DataEditUser

```json
{
  "display_name": "string",   // optional
  "avatar": "string",         // optional — file ID from CDN upload
  "status": {                  // optional
    "text": "string",         // optional — status text
    "presence": "Online"      // optional — Online, Idle, Focus, Busy, Invisible
  },
  "profile": {                 // optional
    "content": "string",      // optional — profile bio (markdown)
    "background": "string"    // optional — file ID for background
  },
  "badges": 0,                // optional — bitfield (admin only)
  "flags": 0,                 // optional — bitfield (admin only)
  "remove": ["Avatar", ...]   // optional — fields to clear
}
```

Remove options (FieldsUser): `"Avatar"`, `"StatusText"`, `"StatusPresence"`, `"ProfileContent"`, `"ProfileBackground"`, `"DisplayName"`

## Change Username

```
PATCH /users/@me/username
Auth: Session Token
Body: DataChangeUsername
Response 200: User
```

### DataChangeUsername

```json
{
  "username": "string",   // required
  "password": "string"    // required
}
```

## Fetch User Profile

Get a user's full profile (bio, background).

```
GET /users/{target}/profile
Auth: Session Token
Path: target (string, required)
Response 200: UserProfile
```

### UserProfile

```json
{
  "content": "string",        // profile bio (markdown)
  "background": { /* File */ } // background image
}
```

## Fetch Default Avatar

Get a user's default generated avatar (no authentication needed).

```
GET /users/{target}/default_avatar
Auth: None
Path: target (string, required)
Response 200: PNG image
```

## Fetch User Flags

```
GET /users/{target}/flags
Auth: Session Token
Path: target (string, required)
Response 200: FlagResponse
```

### FlagResponse

```json
{
  "flags": 0  // uint32 bitfield
}
```

### User Flags

| Flag | Value | Description |
|------|-------|-------------|
| Suspended | 1 | User is suspended |
| Deleted | 2 | User is deleted |
| Banned | 4 | User is banned |
| Spam | 8 | User flagged as spam |

### User Badges

| Badge | Value | Bit |
|-------|-------|-----|
| Developer | 1 | 1 << 0 |
| Translator | 2 | 1 << 1 |
| Supporter | 4 | 1 << 2 |
| ResponsibleDisclosure | 8 | 1 << 3 |
| Founder | 16 | 1 << 4 |
| PlatformModeration | 32 | 1 << 5 |
| ActiveSupporter | 64 | 1 << 6 |
| Paw | 128 | 1 << 7 |
| EarlyAdopter | 256 | 1 << 8 |
| ReservedRelevantJokeBadge1 | 512 | 1 << 9 |
| ReservedRelevantJokeBadge2 | 1024 | 1 << 10 |

## Fetch Mutual Friends, Servers, Groups, and DMs

```
GET /users/{target}/mutual
Auth: Session Token
Path: target (string, required)
Response 200: MutualResponse
```

### MutualResponse

```json
{
  "users": ["user_id", ...],     // mutual friend IDs
  "servers": ["server_id", ...], // mutual server IDs
  "groups": ["channel_id", ...], // mutual group IDs (undocumented)
  "dms": ["channel_id", ...]     // mutual DM channels (undocumented)
}
```

## User Type

See [schemas.md](../types/schemas.md#user) for the full User type definition.

### Presence

String enum: `"Online"`, `"Idle"`, `"Focus"`, `"Busy"`, `"Invisible"`

### RelationshipStatus

String enum: `"None"`, `"User"`, `"Friend"`, `"Outgoing"`, `"Incoming"`, `"Blocked"`, `"BlockedOther"`
