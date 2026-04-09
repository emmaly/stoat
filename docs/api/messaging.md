# Messaging

Message sending, fetching, editing, deleting, searching, and pinning.

## Send Message

```
POST /channels/{target}/messages
Auth: Session Token
Path: target (string, required) — channel ID
Header: Idempotency-Key (string, optional) — prevent duplicate sends
Body: DataMessageSend
Response 200: Message
```

### DataMessageSend

```json
{
  "content": "string",              // optional — message text (markdown)
  "attachments": ["file_id", ...],  // optional — file IDs from CDN upload
  "replies": [                       // optional — reply references
    {
      "id": "message_id",
      "mention": true               // whether to mention the author
    }
  ],
  "embeds": [SendableEmbed, ...],   // optional — text embeds
  "masquerade": {                    // optional — name/avatar override
    "name": "string",
    "avatar": "url",
    "colour": "string"              // CSS color for username
  },
  "interactions": {                  // optional
    "reactions": ["emoji_id", ...], // restrict reactions to these emoji
    "restrict_reactions": false      // if true, only listed reactions allowed
  },
  "flags": 0                        // optional — MessageFlags bitfield
}
```

Either `content` or `attachments` must be provided (at least one).

### SendableEmbed

```json
{
  "icon_url": "string",    // optional
  "url": "string",         // optional
  "title": "string",       // optional
  "description": "string", // optional
  "media": "string",       // optional — file ID
  "colour": "string"       // optional — CSS color
}
```

### ReplyIntent

```json
{
  "id": "string",     // message ID to reply to
  "mention": true     // whether to ping the author
}
```

### Masquerade

```json
{
  "name": "string",   // nullable — override display name
  "avatar": "string", // nullable — override avatar URL
  "colour": "string"  // nullable — override username color
}
```

## Fetch Messages

```
GET /channels/{target}/messages
Auth: Session Token
Path: target (string, required)
Query:
  - limit (integer, optional) — max messages to return (default 50, max 100)
  - before (string, optional) — message ID, fetch messages before this
  - after (string, optional) — message ID, fetch messages after this
  - sort (string, optional) — "Latest" or "Oldest" (default "Latest")
  - nearby (string, optional) — message ID, fetch messages around this
  - include_users (boolean, optional) — include user/member objects
Response 200: BulkMessageResponse
```

### BulkMessageResponse

When `include_users` is false:
```json
[Message, ...]
```

When `include_users` is true:
```json
{
  "messages": [Message, ...],
  "users": [User, ...],
  "members": [Member, ...]
}
```

### MessageSort

String enum: `"Relevance"`, `"Latest"`, `"Oldest"`

## Fetch Message

```
GET /channels/{target}/messages/{msg}
Auth: Session Token
Path: target (string, required), msg (string, required)
Response 200: Message
```

## Edit Message

```
PATCH /channels/{target}/messages/{msg}
Auth: Session Token
Path: target (string, required), msg (string, required)
Body: DataEditMessage
Response 200: Message
```

### DataEditMessage

```json
{
  "content": "string",            // optional
  "embeds": [SendableEmbed, ...]  // optional
}
```

## Delete Message

```
DELETE /channels/{target}/messages/{msg}
Auth: Session Token
Path: target (string, required), msg (string, required)
Response 204: Success
```

## Bulk Delete Messages

```
DELETE /channels/{target}/messages/bulk
Auth: Session Token
Path: target (string, required)
Body: OptionsBulkDelete
Response 204: Success
```

### OptionsBulkDelete

```json
{
  "ids": ["message_id", ...]  // required — array of message IDs to delete
}
```

## Search for Messages

```
POST /channels/{target}/search
Auth: Session Token
Path: target (string, required)
Body: DataMessageSearch
Response 200: BulkMessageResponse
```

### DataMessageSearch

```json
{
  "query": "string",           // required — search text
  "limit": 0,                  // optional
  "before": "string",         // optional — message ID
  "after": "string",          // optional — message ID
  "sort": "Relevance",        // optional — Relevance, Latest, Oldest
  "include_users": false       // optional
}
```

## Pin Message

```
POST /channels/{target}/messages/{msg}/pin
Auth: Session Token
Path: target (string, required), msg (string, required)
Response 204: Success
```

## Unpin Message

```
DELETE /channels/{target}/messages/{msg}/pin
Auth: Session Token
Path: target (string, required), msg (string, required)
Response 204: Success
```

## Acknowledge Message

Mark messages as read up to a given message.

```
PUT /channels/{target}/ack/{message}
Auth: Session Token
Path: target (string, required), message (string, required)
Response 204: Success
```

## Message Type

```json
{
  "_id": "string",
  "nonce": "string",                    // nullable — client-generated dedup key
  "channel": "string",
  "author": "string",
  "user": { /* User */ },               // nullable — populated when include_users
  "member": { /* Member */ },           // nullable — populated when include_users
  "webhook": {                          // nullable
    "name": "string",
    "avatar": "string"
  },
  "content": "string",                  // nullable
  "system": { /* SystemMessage */ },    // nullable — system event
  "attachments": [File, ...],           // nullable
  "edited": "ISO8601",                  // nullable
  "embeds": [Embed, ...],              // nullable
  "mentions": ["user_id", ...],        // nullable
  "role_mentions": ["role_id", ...],   // nullable
  "replies": ["message_id", ...],      // nullable
  "reactions": {                        // emoji_id → [user_id, ...]
    "emoji_id": ["user_id", ...]
  },
  "interactions": {
    "reactions": ["emoji_id", ...],
    "restrict_reactions": false
  },
  "pinned": false,
  "flags": 0
}
```

### SystemMessage Variants

| Type | Fields | Description |
|------|--------|-------------|
| `text` | `content` | System text |
| `user_added` | `id`, `by` | User added to group |
| `user_remove` | `id`, `by` | User removed from group |
| `user_joined` | `id` | User joined server |
| `user_left` | `id` | User left server |
| `user_kicked` | `id` | User kicked |
| `user_banned` | `id` | User banned |
| `channel_renamed` | `name`, `by` | Channel renamed |
| `channel_description_changed` | `by` | Description changed |
| `channel_icon_changed` | `by` | Icon changed |
| `message_pinned` | `id`, `by` | Message pinned |
| `message_unpinned` | `id`, `by` | Message unpinned |
