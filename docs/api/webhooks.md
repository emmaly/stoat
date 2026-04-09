# Webhooks

Webhook CRUD and execution. Webhooks allow external services to post messages to channels.

## Create Webhook

```
POST /channels/{channel_id}/webhooks
Auth: Session Token
Path: channel_id (string, required)
Body: CreateWebhookBody
Response 200: Webhook
```

### CreateWebhookBody

```json
{
  "name": "string",   // required
  "avatar": "string"  // optional — file ID
}
```

## Fetch Webhooks for Channel

```
GET /channels/{channel_id}/webhooks
Auth: Session Token
Path: channel_id (string, required)
Response 200: Webhook[]
```

## Fetch Webhook (Authenticated)

```
GET /webhooks/{webhook_id}
Auth: Session Token
Path: webhook_id (string, required)
Response 200: Webhook
```

## Fetch Webhook (Token Auth)

No session required — uses webhook token.

```
GET /webhooks/{webhook_id}/{token}
Auth: None (token in path)
Path: webhook_id (string, required), token (string, required)
Response 200: Webhook
```

## Edit Webhook (Authenticated)

```
PATCH /webhooks/{webhook_id}
Auth: Session Token
Path: webhook_id (string, required)
Body: DataEditWebhook
Response 200: Webhook
```

## Edit Webhook (Token Auth)

```
PATCH /webhooks/{webhook_id}/{token}
Auth: None (token in path)
Path: webhook_id (string, required), token (string, required)
Body: DataEditWebhook
Response 200: Webhook
```

### DataEditWebhook

```json
{
  "name": "string",     // optional
  "avatar": "string",   // optional — file ID
  "remove": ["Avatar"]  // optional — fields to clear
}
```

Remove options (FieldsWebhook): `"Avatar"`

## Delete Webhook (Authenticated)

```
DELETE /webhooks/{webhook_id}
Auth: Session Token
Path: webhook_id (string, required)
Response 204: Success
```

## Delete Webhook (Token Auth)

```
DELETE /webhooks/{webhook_id}/{token}
Auth: None (token in path)
Path: webhook_id (string, required), token (string, required)
Response 204: Success
```

## Execute Webhook

Post a message via webhook.

```
POST /webhooks/{webhook_id}/{token}
Auth: None (token in path)
Path: webhook_id (string, required), token (string, required)
Body: DataMessageSend
Response 200: Message
```

The body uses the same `DataMessageSend` type as regular message sending. See [Messaging](messaging.md).

## Execute GitHub Webhook

Special endpoint for GitHub webhook payloads — formats GitHub events as messages.

```
POST /webhooks/{webhook_id}/{token}/github
Auth: None (token in path)
Path: webhook_id (string, required), token (string, required)
Body: (GitHub webhook payload)
Response 204: Success
```

## Webhook Type

```json
{
  "id": "string",
  "name": "string",
  "avatar": { /* File */ },   // nullable
  "channel_id": "string",
  "token": "string",          // nullable — only visible to webhook creator
  "permissions": 0             // int64 — permission bitfield
}
```

### ResponseWebhook

```json
{
  "id": "string",
  "name": "string",
  "avatar": { /* File */ },
  "channel_id": "string",
  "token": "string",
  "permissions": 0
}
```

### MessageWebhook

Included in messages sent via webhooks:

```json
{
  "name": "string",
  "avatar": "string"  // nullable
}
```

## Notes

- Webhook token authentication allows external services to interact without a session
- The GitHub endpoint specifically handles GitHub's webhook payload format
- Webhooks require the `ManageWebhooks` (1 << 24) permission to create/manage
