# Server Information

Server CRUD and server channel management.

## Create Server

```
POST /servers/create
Auth: Session Token
Body: DataCreateServer
Response 200: CreateServerLegacyResponse
```

### DataCreateServer

```json
{
  "name": "string",         // required
  "description": "string",  // optional
  "nsfw": false              // optional
}
```

### CreateServerLegacyResponse

```json
{
  "server": { /* Server */ },
  "channels": [Channel, ...]
}
```

## Fetch Server

```
GET /servers/{target}
Auth: Session Token
Path: target (string, required)
Query: include_channels (boolean, optional) — include full channel objects
Response 200: FetchServerResponse
```

### FetchServerResponse

When `include_channels` is false (default):
```json
{ /* Server object */ }
```

When `include_channels` is true:
```json
{
  "server": { /* Server */ },
  "channels": [Channel, ...]
}
```

(This is a tagged union — the response shape changes based on the query parameter.)

## Edit Server

```
PATCH /servers/{target}
Auth: Session Token + Valid MFA Ticket
Path: target (string, required)
Body: DataEditServer
Response 200: Server
```

### DataEditServer

```json
{
  "name": "string",              // optional
  "description": "string",       // optional
  "icon": "string",              // optional — file ID
  "banner": "string",            // optional — file ID
  "categories": [Category, ...], // optional
  "system_messages": { ... },    // optional — SystemMessageChannels
  "flags": 0,                    // optional
  "discoverable": false,         // optional
  "analytics": false,            // optional
  "remove": ["Icon", ...]        // optional — fields to clear
}
```

Remove options (FieldsServer): `"Icon"`, `"Banner"`, `"Description"`

### Category

```json
{
  "id": "string",
  "title": "string",
  "channels": ["channel_id", ...]
}
```

### SystemMessageChannels

```json
{
  "user_joined": "string",     // channel ID — nullable
  "user_left": "string",       // channel ID — nullable
  "user_kicked": "string",     // channel ID — nullable
  "user_banned": "string"      // channel ID — nullable
}
```

## Delete / Leave Server

```
DELETE /servers/{target}
Auth: Session Token
Path: target (string, required)
Query: leave_silently (boolean, optional) — don't send leave message
Response 204: Success
```

If you are the owner, this **deletes** the server. Otherwise, you **leave** the server.

## Mark Server As Read

Mark all channels in the server as read.

```
PUT /servers/{target}/ack
Auth: Session Token
Path: target (string, required)
Response 204: Success
```

## Create Server Channel

```
POST /servers/{server}/channels
Auth: Session Token
Path: server (string, required)
Body: DataCreateServerChannel
Response 200: Channel
```

### DataCreateServerChannel

```json
{
  "type": "Text",           // optional — "Text" or "Voice" (default: "Text")
  "name": "string",         // required
  "description": "string",  // optional
  "nsfw": false              // optional
}
```

### LegacyServerChannelType

String enum: `"Text"`, `"Voice"`

## Server Type

```json
{
  "_id": "string",
  "owner": "string",
  "name": "string",
  "description": "string",             // nullable
  "channels": ["channel_id", ...],
  "categories": [Category, ...],       // nullable
  "system_messages": { ... },          // nullable — SystemMessageChannels
  "roles": {                           // map of role_id → Role
    "role_id": { /* Role */ }
  },
  "default_permissions": 0,            // int64 permission bitfield
  "icon": { /* File */ },              // nullable
  "banner": { /* File */ },            // nullable
  "flags": 0,                          // uint32
  "nsfw": false,
  "analytics": false,
  "discoverable": false
}
```

### Server Flags

| Flag | Value | Description |
|------|-------|-------------|
| Verified | 1 | Officially verified server |
| Official | 2 | Official Stoat server |
