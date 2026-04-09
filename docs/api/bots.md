# Bots

Bot creation, management, and invitation endpoints.

## Create Bot

```
POST /bots/create
Auth: Session Token
Body: DataCreateBot
Response 200: BotWithUserResponse
```

### DataCreateBot

```json
{
  "name": "string"  // required — bot display name
}
```

### BotWithUserResponse

```json
{
  "bot": { /* Bot */ },
  "user": { /* User */ }
}
```

## Fetch Bot

Fetch a bot you own.

```
GET /bots/{bot_id}
Auth: Session Token
Path: bot_id (string, required)
Response 200: FetchBotResponse
```

### FetchBotResponse

```json
{
  "bot": { /* Bot */ },
  "user": { /* User */ }
}
```

## Edit Bot

```
PATCH /bots/{bot_id}
Auth: Session Token
Path: bot_id (string, required)
Body: DataEditBot
Response 200: BotWithUserResponse
```

### DataEditBot

```json
{
  "name": "string",              // optional
  "public": true,                // optional — whether bot is publicly invitable
  "analytics": true,             // optional
  "interactions_url": "string",  // optional — webhook URL for interactions
  "remove": ["Token", ...]       // optional — fields to clear
}
```

Remove options (FieldsBot): `"Token"`, `"InteractionsURL"`

## Delete Bot

```
DELETE /bots/{bot_id}
Auth: Session Token
Path: bot_id (string, required)
Response 204: Success
```

## Fetch Owned Bots

List all bots owned by the authenticated user.

```
GET /bots/@me
Auth: Session Token
Response 200: OwnedBotsResponse
```

### OwnedBotsResponse

```json
{
  "bots": [Bot, ...],
  "users": [User, ...]
}
```

Both lists are sorted by ID.

## Fetch Public Bot

Get public information about a bot (for invite pages).

```
GET /bots/{target}/invite
Auth: Session Token
Path: target (string, required) — bot ID
Response 200: PublicBot
```

### PublicBot

```json
{
  "_id": "string",
  "username": "string",
  "avatar": { /* File */ },
  "description": "string"
}
```

## Invite Bot

Add a bot to a server or group.

```
POST /bots/{target}/invite
Auth: Session Token
Path: target (string, required) — bot ID
Body: InviteBotDestination
Response 204: Success
```

### InviteBotDestination

Tagged union:

```json
{"server": "<server_id>"}
```
or
```json
{"group": "<group_channel_id>"}
```

## Bot Type

```json
{
  "_id": "string",              // bot user ID
  "owner": "string",            // owner user ID
  "token": "string",            // bot token (only visible to owner)
  "public": true,               // whether publicly invitable
  "analytics": false,
  "discoverable": false,
  "interactions_url": "string", // webhook URL for interactions
  "terms_of_service_url": "string",
  "privacy_policy_url": "string",
  "flags": 0                    // uint32 bitfield
}
```

### BotInformation

Appears on User objects when the user is a bot:

```json
{
  "owner": "string"  // owner user ID
}
```
