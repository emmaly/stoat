# Channel Information

Channel CRUD and channel-level permission management.

## Fetch Channel

```
GET /channels/{target}
Auth: Session Token
Path: target (string, required) — channel ID
Response 200: Channel
```

## Edit Channel

```
PATCH /channels/{target}
Auth: Session Token
Path: target (string, required)
Body: DataEditChannel
Response 200: Channel
```

### DataEditChannel

```json
{
  "name": "string",         // optional
  "description": "string",  // optional
  "icon": "string",         // optional — file ID
  "nsfw": false,             // optional
  "archived": false,         // optional
  "remove": ["Icon", ...]   // optional — fields to clear
}
```

Remove options (FieldsChannel): `"Icon"`, `"Description"`, `"DefaultPermissions"`

## Close Channel

Close a DM, leave a group, or delete a server channel.

```
DELETE /channels/{target}
Auth: Session Token
Path: target (string, required)
Query: leave_silently (boolean, optional) — don't send leave message (groups only)
Response 204: Success
```

Behavior depends on channel type:
- **SavedMessages**: Not deletable
- **DirectMessage**: Closes the DM (sets `active` to false)
- **Group**: Leave the group (or delete if owner)
- **TextChannel / VoiceChannel**: Delete the channel (requires ManageChannel permission)

## Set Default Channel Permissions

Set default permission overrides for a channel.

```
PUT /channels/{target}/permissions/default
Auth: Session Token
Path: target (string, required)
Body: DataDefaultChannelPermissions
Response 200: Channel
```

### DataDefaultChannelPermissions

This is a tagged union:

**Override variant** (for server channels):
```json
{
  "permissions": {
    "allow": 0,
    "deny": 0
  }
}
```

**Value variant** (for groups):
```json
{
  "permissions": 0
}
```

## Set Role Permission on Channel

```
PUT /channels/{target}/permissions/{role_id}
Auth: Session Token
Path: target (string, required), role_id (string, required)
Body: DataSetRolePermissions
Response 200: Channel
```

### DataSetRolePermissions

```json
{
  "permissions": {
    "allow": 0,
    "deny": 0
  }
}
```

## Channel Type

Channel is a tagged union (discriminated by `channel_type`):

### SavedMessages

```json
{
  "channel_type": "SavedMessages",
  "_id": "string",
  "user": "string"
}
```

### DirectMessage

```json
{
  "channel_type": "DirectMessage",
  "_id": "string",
  "active": true,
  "recipients": ["user_id", "user_id"],
  "last_message_id": "string"  // nullable
}
```

### Group

```json
{
  "channel_type": "Group",
  "_id": "string",
  "name": "string",
  "owner": "string",
  "description": "string",           // nullable
  "recipients": ["user_id", ...],
  "icon": { /* File */ },             // nullable
  "last_message_id": "string",        // nullable
  "permissions": 0,                   // nullable — int64
  "nsfw": false
}
```

### TextChannel

```json
{
  "channel_type": "TextChannel",
  "_id": "string",
  "server": "string",
  "name": "string",
  "description": "string",              // nullable
  "icon": { /* File */ },                // nullable
  "last_message_id": "string",          // nullable
  "default_permissions": { /* Override */ }, // nullable
  "role_permissions": {                  // map of role_id → Override
    "role_id": {"a": 0, "d": 0}
  },
  "nsfw": false
}
```

### VoiceChannel

```json
{
  "channel_type": "VoiceChannel",
  "_id": "string",
  "server": "string",
  "name": "string",
  "description": "string",              // nullable
  "icon": { /* File */ },                // nullable
  "default_permissions": { /* Override */ }, // nullable
  "role_permissions": {
    "role_id": {"a": 0, "d": 0}
  },
  "nsfw": false
}
```
