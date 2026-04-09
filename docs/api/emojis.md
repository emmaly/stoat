# Emojis

Custom emoji management.

## Fetch Emoji

Get a custom emoji by ID (no auth required).

```
GET /custom/emoji/{emoji_id}
Auth: None
Path: emoji_id (string, required)
Response 200: Emoji
```

## Create Emoji

Upload a new custom emoji. The `emoji_id` in the path should be the file ID from a CDN upload to the `emojis` tag.

```
PUT /custom/emoji/{emoji_id}
Auth: Session Token
Path: emoji_id (string, required) — file ID from CDN upload
Body: DataCreateEmoji
Response 200: Emoji
```

### DataCreateEmoji

```json
{
  "name": "string",    // required — emoji name
  "parent": {          // required — where the emoji belongs
    "type": "Server",
    "id": "server_id"
  },
  "nsfw": false         // optional
}
```

### EmojiParent

Tagged union:

```json
{"type": "Server", "id": "server_id"}
```
or
```json
{"type": "Detached"}
```

## Delete Emoji

```
DELETE /custom/emoji/{emoji_id}
Auth: Session Token
Path: emoji_id (string, required)
Response 204: Success
```

## Emoji Type

```json
{
  "_id": "string",
  "parent": { /* EmojiParent */ },
  "creator_id": "string",
  "name": "string",
  "animated": false,
  "nsfw": false
}
```

## Usage

- Custom emoji can be used in messages by referencing their ID
- Emoji images are served from the CDN at the `emojis` tag path
- The `ManageCustomisation` (1 << 4) permission is required to create/delete server emoji
- Emoji are also delivered via WebSocket events (`EmojiCreate`, `EmojiDelete`)
