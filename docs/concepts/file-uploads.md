# File Uploads

File uploads go through the CDN service (formerly called "Autumn"), hosted at `https://cdn.stoatusercontent.com`.

## Upload Flow

1. **POST** multipart form data to `{cdn_base}/{tag}` with a `file` field
2. Include authentication header (`X-Session-Token` or `X-Bot-Token`)
3. Receive response: `{"id": "<file_id>"}`
4. Use the returned `id` in subsequent API calls (e.g., message attachments, avatar changes)

## Upload Tags

Tags determine the category/purpose of the uploaded file:

| Tag | Purpose | Used By |
|-----|---------|---------|
| `attachments` | Message attachments | `POST /channels/{target}/messages` |
| `avatars` | User/server/bot avatars | User/server/member edit endpoints |
| `backgrounds` | Profile backgrounds | `PATCH /users/{target}` (profile) |
| `icons` | Channel/server icons | Channel/server edit endpoints |
| `banners` | Server banners | `PATCH /servers/{target}` |
| `emojis` | Custom emoji | `PUT /custom/emoji/{emoji_id}` |

## File URLs

After upload, files are served at:

- **Preview/thumbnail:** `https://cdn.stoatusercontent.com/{tag}/{file_id}`
- **Original (animated):** `https://cdn.stoatusercontent.com/{tag}/{file_id}/original`

## File Object Schema

When files appear in API responses (e.g., as message attachments), they use the `File` type:

```json
{
  "_id": "string",           // Unique file ID
  "tag": "string",           // Upload tag (attachments, avatars, etc.)
  "filename": "string",      // Original filename
  "metadata": { ... },       // Metadata (type-dependent: Image, Video, Text, Audio, File)
  "content_type": "string",  // MIME type
  "size": 0,                 // File size in bytes
  "deleted": false,          // Whether the file has been deleted
  "reported": false,         // Whether the file has been reported
  "message_id": "string",    // Associated message ID (if attachment)
  "user_id": "string",       // Uploader's user ID
  "server_id": "string",     // Associated server ID (if applicable)
  "object_id": "string"      // Associated object ID
}
```

## Metadata Variants

The `metadata` field varies by file type:

| Type | Fields |
|------|--------|
| `File` | (no additional fields) |
| `Text` | (no additional fields) |
| `Image` | `width: int`, `height: int` |
| `Video` | `width: int`, `height: int` |
| `Audio` | (no additional fields) |

## Notes

- The CDN API docs at `https://cdn.stoatusercontent.com/scalar` are currently inaccessible (403)
- File size limits and supported MIME types are determined by server configuration
- The legacy CDN domain `https://cdn.revoltusercontent.com` redirects to the new Stoat domain
- The proxy service at `https://external.stoatusercontent.com` (formerly `jan.revolt.chat`) handles external URL proxying for embeds
