# Direct Messaging

## Fetch Direct Message Channels

List all DM channels for the authenticated user.

```
GET /users/dms
Auth: Session Token
Response 200: Channel[] (DirectMessage and SavedMessages channels)
```

## Open Direct Message

Open or retrieve an existing DM channel with a user.

```
GET /users/{target}/dm
Auth: Session Token
Path: target (string, required) — user ID
Response 200: Channel (DirectMessage)
```

If a DM channel already exists, it is returned. If not, a new one is created.

## DM Channel Types

### SavedMessages

Personal "Saved Notes" channel.

```json
{
  "channel_type": "SavedMessages",
  "_id": "string",
  "user": "string"  // your user ID
}
```

### DirectMessage

DM between two users.

```json
{
  "channel_type": "DirectMessage",
  "_id": "string",
  "active": true,              // whether both sides have it open
  "recipients": ["user_id_1", "user_id_2"],
  "last_message_id": "string"  // nullable
}
```

Once a DM channel exists, use the standard [Messaging](messaging.md) endpoints to send/receive messages.
