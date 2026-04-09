# WebSocket Events

All events are JSON objects with a `type` field as the discriminator.

## Client → Server Events

### Authenticate

Authenticate with the server after connecting.

```json
{"type": "Authenticate", "token": "<token>"}
```

### Ping

Keep-alive ping. Server responds with Pong echoing the `data` field.

```json
{"type": "Ping", "data": 0}
```

### BeginTyping

Notify others that you started typing in a channel. Must be a member of the channel.

```json
{"type": "BeginTyping", "channel": "<channel_id>"}
```

### EndTyping

Notify others that you stopped typing.

```json
{"type": "EndTyping", "channel": "<channel_id>"}
```

### Subscribe

Subscribe to a server's UserUpdate events. Only needed for normal users (not bots).

```json
{"type": "Subscribe", "server_id": "<server_id>"}
```

- Expires after 15 minutes
- Max 5 active subscriptions
- Resend at most every 10 minutes per server
- Only send when app is in focus

---

## Server → Client Events

### Connection Events

#### Authenticated

Confirms successful authentication. No additional fields.

```json
{"type": "Authenticated"}
```

#### Ready

Initial state dump after authentication. All fields are optional based on `ready` query parameter.

```json
{
  "type": "Ready",
  "users": [User, ...],
  "servers": [Server, ...],
  "channels": [Channel, ...],
  "members": [Member, ...],
  "emojis": [Emoji, ...],
  "user_settings": { ... },
  "channel_unreads": [ChannelUnread, ...],
  "policy_changes": { ... }
}
```

#### Error

An error occurred.

```json
{"type": "Error", "error": "<error_id>"}
```

Error IDs:
| ID | Description |
|----|-------------|
| `LabelMe` | Uncategorized error |
| `InternalError` | Server internal error |
| `InvalidSession` | Invalid or expired token |
| `OnboardingNotFinished` | User must complete onboarding first |
| `AlreadyAuthenticated` | Connection is already authenticated |

#### Logout

Session has been invalidated. Connection will close shortly.

```json
{"type": "Logout"}
```

#### Pong

Response to client Ping.

```json
{"type": "Pong", "data": 0}
```

#### Bulk

Multiple events bundled together. Process each item in `v` as a separate event.

```json
{"type": "Bulk", "v": [Event, ...]}
```

---

### Message Events

#### Message

New message received.

```json
{
  "type": "Message",
  // ... all Message object fields
  "_id": "<message_id>",
  "channel": "<channel_id>",
  "author": "<user_id>",
  "content": "Hello!",
  // etc.
}
```

#### MessageUpdate

A message was edited.

```json
{
  "type": "MessageUpdate",
  "id": "<message_id>",
  "channel": "<channel_id>",
  "data": {
    // Partial Message object — only changed fields
  }
}
```

#### MessageAppend

Data appended to a message (e.g., embeds loaded after send).

```json
{
  "type": "MessageAppend",
  "id": "<message_id>",
  "channel": "<channel_id>",
  "append": {
    "embeds": [Embed, ...]
  }
}
```

#### MessageDelete

A message was deleted.

```json
{
  "type": "MessageDelete",
  "id": "<message_id>",
  "channel": "<channel_id>"
}
```

#### MessageReact

A reaction was added to a message.

```json
{
  "type": "MessageReact",
  "id": "<message_id>",
  "channel_id": "<channel_id>",
  "user_id": "<user_id>",
  "emoji_id": "<emoji_id>"
}
```

#### MessageUnreact

A specific user removed their reaction.

```json
{
  "type": "MessageUnreact",
  "id": "<message_id>",
  "channel_id": "<channel_id>",
  "user_id": "<user_id>",
  "emoji_id": "<emoji_id>"
}
```

#### MessageRemoveReaction

All reactions of a specific emoji removed from a message.

```json
{
  "type": "MessageRemoveReaction",
  "id": "<message_id>",
  "channel_id": "<channel_id>",
  "emoji_id": "<emoji_id>"
}
```

---

### Channel Events

#### ChannelCreate

A new channel was created.

```json
{
  "type": "ChannelCreate",
  // ... all Channel object fields
}
```

#### ChannelUpdate

A channel was edited. The `clear` array lists fields that were removed/reset.

```json
{
  "type": "ChannelUpdate",
  "id": "<channel_id>",
  "data": {
    // Partial Channel object
  },
  "clear": ["Icon", "Description"]
}
```

Clearable fields: `Icon`, `Description`

#### ChannelDelete

A channel was deleted.

```json
{
  "type": "ChannelDelete",
  "id": "<channel_id>"
}
```

#### ChannelGroupJoin

A user joined a group channel.

```json
{
  "type": "ChannelGroupJoin",
  "id": "<channel_id>",
  "user": "<user_id>"
}
```

#### ChannelGroupLeave

A user left a group channel.

```json
{
  "type": "ChannelGroupLeave",
  "id": "<channel_id>",
  "user": "<user_id>"
}
```

#### ChannelStartTyping

A user started typing in a channel.

```json
{
  "type": "ChannelStartTyping",
  "id": "<channel_id>",
  "user": "<user_id>"
}
```

#### ChannelStopTyping

A user stopped typing.

```json
{
  "type": "ChannelStopTyping",
  "id": "<channel_id>",
  "user": "<user_id>"
}
```

#### ChannelAck

You acknowledged messages in a channel up to a specific message.

```json
{
  "type": "ChannelAck",
  "id": "<channel_id>",
  "user": "<user_id>",
  "message_id": "<message_id>"
}
```

---

### Server Events

#### ServerCreate

A new server was created (or you joined one).

```json
{
  "type": "ServerCreate",
  // ... all Server object fields
}
```

#### ServerUpdate

A server was edited.

```json
{
  "type": "ServerUpdate",
  "id": "<server_id>",
  "data": {
    // Partial Server object
  },
  "clear": ["Icon", "Banner", "Description"]
}
```

Clearable fields: `Icon`, `Banner`, `Description`

#### ServerDelete

A server was deleted (or you left it).

```json
{
  "type": "ServerDelete",
  "id": "<server_id>"
}
```

#### ServerMemberUpdate

A server member was edited.

```json
{
  "type": "ServerMemberUpdate",
  "id": {
    "server": "<server_id>",
    "user": "<user_id>"
  },
  "data": {
    // Partial Member object
  },
  "clear": ["Nickname", "Avatar"]
}
```

Clearable fields: `Nickname`, `Avatar`

#### ServerMemberJoin

A user joined a server.

```json
{
  "type": "ServerMemberJoin",
  "id": "<server_id>",
  "user": "<user_id>",
  "member": { /* Member object */ }
}
```

#### ServerMemberLeave

A user left a server (or was kicked/banned).

```json
{
  "type": "ServerMemberLeave",
  "id": "<server_id>",
  "user": "<user_id>"
}
```

#### ServerRoleUpdate

A role was edited.

```json
{
  "type": "ServerRoleUpdate",
  "id": "<server_id>",
  "role_id": "<role_id>",
  "data": {
    // Partial Role object
  },
  "clear": ["Colour"]
}
```

Clearable fields: `Colour`

#### ServerRoleDelete

A role was deleted.

```json
{
  "type": "ServerRoleDelete",
  "id": "<server_id>",
  "role_id": "<role_id>"
}
```

---

### User Events

#### UserUpdate

A user's profile or status changed.

```json
{
  "type": "UserUpdate",
  "id": "<user_id>",
  "data": {
    // Partial User object
  },
  "clear": ["ProfileContent", "ProfileBackground", "StatusText", "Avatar", "DisplayName"]
}
```

Clearable fields: `ProfileContent`, `ProfileBackground`, `StatusText`, `Avatar`, `DisplayName`

#### UserRelationship

Your relationship with another user changed.

```json
{
  "type": "UserRelationship",
  "id": "<your_user_id>",
  "user": { /* User object */ },
  "status": "<RelationshipStatus>"
}
```

#### UserPlatformWipe

A user was platform-banned or deleted their account. Clients should remove all traces of this user: messages, DM channels, relationships, server memberships.

```json
{
  "type": "UserPlatformWipe",
  "user_id": "<user_id>",
  "flags": 0
}
```

---

### Emoji Events

#### EmojiCreate

A new custom emoji was created.

```json
{
  "type": "EmojiCreate",
  // ... all Emoji object fields
}
```

#### EmojiDelete

A custom emoji was deleted.

```json
{
  "type": "EmojiDelete",
  "id": "<emoji_id>"
}
```

---

### Auth Events

#### Auth

Forwarded events from the authentication system (Authifier). Currently only session deletion events are forwarded.

##### DeleteSession

A specific session was revoked.

```json
{
  "type": "Auth",
  "event_type": "DeleteSession",
  "user_id": "<user_id>",
  "session_id": "<session_id>"
}
```

##### DeleteAllSessions

All sessions were revoked, optionally excluding one.

```json
{
  "type": "Auth",
  "event_type": "DeleteAllSessions",
  "user_id": "<user_id>",
  "exclude_session_id": "<session_id>"
}
```
