# Sync

Client settings and unread state synchronization.

## Fetch Settings

Retrieve client-specific settings (stored as opaque key-value pairs).

```
POST /sync/settings/fetch
Auth: Session Token
Body: OptionsFetchSettings
Response 200: object (key → [timestamp, value] pairs)
```

### OptionsFetchSettings

```json
{
  "keys": ["key1", "key2"]  // required — array of setting keys to fetch
}
```

The response is a map of key → `[timestamp, value]` tuples. The timestamp indicates when the setting was last modified.

## Set Settings

Store client-specific settings.

```
POST /sync/settings/set
Auth: Session Token
Body: object (key → value pairs)
Response 204: Success
```

The body is a map of setting key → string value.

## Fetch Unreads

Get unread message state for all channels.

```
GET /sync/unreads
Auth: Session Token
Response 200: ChannelUnread[]
```

### ChannelUnread

```json
{
  "_id": {
    "channel": "string",
    "user": "string"
  },
  "last_id": "string",              // nullable — last read message ID
  "mentions": ["message_id", ...]   // nullable — unread mention message IDs
}
```

## Notes

- Settings are opaque to the server — clients define their own key names
- The sync system allows multiple clients to share configuration
- Unread state is per-channel, tracking both the last read message and any mentions
