# Permissions

Stoat uses a bitfield-based permission system. Permissions are applied sequentially: **allows first, then denies**.

## Permission Flags

All permissions are represented as bits in a 64-bit integer.

### Server Management Permissions (bits 0–13)

| Permission | Value | Bit | Description |
|------------|-------|-----|-------------|
| ManageChannel | 1 | 1 << 0 | Create, edit, delete channels |
| ManageServer | 2 | 1 << 1 | Edit server settings |
| ManagePermissions | 4 | 1 << 2 | Edit permissions and role assignments |
| ManageRole | 8 | 1 << 3 | Create, edit, delete roles |
| ManageCustomisation | 16 | 1 << 4 | Manage server emoji |
| KickMembers | 64 | 1 << 6 | Kick members from server |
| BanMembers | 128 | 1 << 7 | Ban members from server |
| TimeoutMembers | 256 | 1 << 8 | Timeout members |
| AssignRoles | 512 | 1 << 9 | Assign roles to members |
| ChangeNickname | 1024 | 1 << 10 | Change own nickname |
| ManageNicknames | 2048 | 1 << 11 | Change other members' nicknames |
| ChangeAvatar | 4096 | 1 << 12 | Change own server avatar |
| RemoveAvatars | 8192 | 1 << 13 | Remove other members' avatars |

### Channel Permissions (bits 20–29)

| Permission | Value | Bit | Description |
|------------|-------|-----|-------------|
| ViewChannel | 1048576 | 1 << 20 | View the channel |
| ReadMessageHistory | 2097152 | 1 << 21 | Read message history |
| SendMessage | 4194304 | 1 << 22 | Send messages |
| ManageMessages | 8388608 | 1 << 23 | Delete others' messages, manage pins |
| ManageWebhooks | 16777216 | 1 << 24 | Create, edit, delete webhooks |
| InviteOthers | 33554432 | 1 << 25 | Create invites |
| SendEmbeds | 67108864 | 1 << 26 | Send embedded content |
| UploadFiles | 134217728 | 1 << 27 | Upload file attachments |
| Masquerade | 268435456 | 1 << 28 | Use masquerade (name/avatar override) |
| React | 536870912 | 1 << 29 | Add reactions to messages |

### Voice Permissions (bits 30–35)

| Permission | Value | Bit | Description |
|------------|-------|-----|-------------|
| Connect | 1073741824 | 1 << 30 | Connect to voice channel |
| Speak | 2147483648 | 1 << 31 | Speak in voice channel |
| Video | 4294967296 | 1 << 32 | Share video |
| MuteMembers | 8589934592 | 1 << 33 | Server-mute members |
| DeafenMembers | 17179869184 | 1 << 34 | Server-deafen members |
| MoveMembers | 34359738368 | 1 << 35 | Move members between voice channels |

## Permission Resolution

Permissions are calculated in this order:

1. Start with the server's `default_permissions` value
2. For each role the member has (in rank order), apply `allow` bits then `deny` bits
3. For channel-specific overrides, apply channel default override, then per-role overrides

### Override Structure

```json
{
  "a": 0,  // allow bits
  "d": 0   // deny bits
}
```

The `Override` type contains `allow` (int64) and `deny` (int64) fields.

## API Endpoints

### Set Default Channel Permissions

```
PUT /channels/{target}/permissions/default
Body: DataDefaultChannelPermissions
→ 200: Channel
```

### Set Role Permissions on Channel

```
PUT /channels/{target}/permissions/{role_id}
Body: DataSetRolePermissions { "permissions": { "allow": 0, "deny": 0 } }
→ 200: Channel
```

### Set Default Server Permissions

```
PUT /servers/{target}/permissions/default
Body: DataPermissionsValue { "permissions": 0 }
→ 200: Server
```

### Set Role Permissions on Server

```
PUT /servers/{target}/permissions/{role_id}
Body: DataSetServerRolePermission { "permissions": { "allow": 0, "deny": 0 } }
→ 200: Server
```

## Notes

- Bit 5 (value 32) is unused / reserved
- Bits 14–19 are unused / reserved
- The server owner implicitly has all permissions
- Permission calculation logic is implemented in the backend's `core/permissions` crate
- For authoritative implementation details, see the [Stoat backend source](https://github.com/stoatchat/stoatchat/tree/main/crates/core/permissions)
