# CDN / File Server (Autumn)

The Stoat CDN service handles file uploads, storage, and serving. It runs as a separate service from the main API.

## URLs

| Environment | URL |
|-------------|-----|
| Production | `https://cdn.stoatusercontent.com` |
| Legacy (redirect) | `https://cdn.revoltusercontent.com` |
| Local dev | `http://local.revolt.chat:14009` |

The CDN URL is provided dynamically via `GET /` (root config) under `features.autumn.url`.

## Upload

### Endpoint

```
POST {cdn_url}/{tag}
Content-Type: multipart/form-data
Auth: X-Session-Token or X-Bot-Token header
```

### Request

Multipart form data with a single `file` field containing the file binary.

### Response

```json
{"id": "<file_id>"}
```

### Upload Tags

| Tag | Purpose | Max Size | Used By |
|-----|---------|----------|---------|
| `attachments` | Message file attachments | Server-configured | Send Message |
| `avatars` | User/bot/member avatars | Server-configured | Edit User, Edit Member, Edit Bot |
| `backgrounds` | Profile backgrounds | Server-configured | Edit User (profile) |
| `icons` | Channel/server/group icons | Server-configured | Edit Channel, Edit Server |
| `banners` | Server banners | Server-configured | Edit Server |
| `emojis` | Custom emoji images | Server-configured | Create Emoji |

Note: Exact size limits are determined by the server's configuration and are not published in the API spec. Check the instance's `GET /` response for `LimitsConfig` details.

## Download / Serve

### File URL Patterns

```
GET {cdn_url}/{tag}/{file_id}                    → Processed/preview version
GET {cdn_url}/{tag}/{file_id}/{filename}          → With filename hint
GET {cdn_url}/{tag}/{file_id}/original            → Original file (animated)
GET {cdn_url}/{tag}/{file_id}/original/{filename} → Original with filename hint
```

### Examples

```
https://cdn.stoatusercontent.com/attachments/abc123
https://cdn.stoatusercontent.com/avatars/def456
https://cdn.stoatusercontent.com/emojis/ghi789/original
```

## File Object

When files appear in API responses (attachments, avatars, icons, etc.), they use the `File` type:

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `_id` | string | yes | Unique file ID |
| `tag` | string | yes | Upload tag |
| `filename` | string | yes | Original filename |
| `metadata` | Metadata | yes | File metadata (varies by type) |
| `content_type` | string | yes | MIME type |
| `size` | integer | yes | File size in bytes |
| `deleted` | boolean | no | Whether file was deleted |
| `reported` | boolean | no | Whether file was reported |
| `message_id` | string | no | Associated message ID |
| `user_id` | string | no | Uploader's user ID |
| `server_id` | string | no | Associated server ID |
| `object_id` | string | no | Associated object ID |

## Metadata Variants

The `metadata` field is a tagged union on `type`:

| Type | Additional Fields | Description |
|------|-------------------|-------------|
| `File` | (none) | Generic file |
| `Text` | (none) | Text file |
| `Image` | `width`, `height` | Image with dimensions |
| `Video` | `width`, `height` | Video with dimensions |
| `Audio` | (none) | Audio file |

## Proxy Service (January)

External URLs embedded in messages are proxied through the January service:

| Environment | URL |
|-------------|-----|
| Production | `https://external.stoatusercontent.com` |
| Legacy (redirect) | `https://jan.revolt.chat` |

The proxy URL is provided via `GET /` under `features.january.url`. It handles:
- URL unfurling for embeds (link previews)
- Image proxying for external images in embeds
- Content safety filtering

## Implementation Notes

1. **Always use the CDN URL from the root config** — don't hardcode domains
2. Upload returns only the file ID — construct the full URL using `{cdn_url}/{tag}/{id}`
3. For animated content (GIFs, animated WebP), use the `/original` path to get the animated version
4. The CDN may return resized/optimized versions at the base path
5. Authentication is required for uploads but not for downloads (files are served publicly by ID)
