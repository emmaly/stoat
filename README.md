# Stoat API Documentation

Complete, structured documentation of the [Stoat](https://stoat.chat) chat platform API — reverse-engineered from the official OpenAPI v3 spec, developer docs, and WebSocket protocol documentation.

**Goal:** Serve as a comprehensive reference for building a full-coverage Go client library (`stoat-go`).

## API Overview

| Property | Value |
|----------|-------|
| API Version | 0.12.0 |
| Base URL (Production) | `https://stoat.chat/api` |
| WebSocket URL | `wss://stoat.chat/events` |
| CDN URL | `https://cdn.stoatusercontent.com` |
| Proxy URL | `https://external.stoatusercontent.com` |
| License | AGPLv3 |
| Spec Format | OpenAPI 3.0.0 |
| Total REST Endpoints | 83 |
| Total Schema Types | 139 |

## Documentation Structure

```
docs/
├── concepts/
│   ├── authentication.md      # Auth methods, tokens, headers
│   ├── permissions.md         # Bitfield permission system
│   ├── rate-limits.md         # Buckets, headers, retry behavior
│   └── file-uploads.md        # CDN upload flow
├── api/
│   ├── core.md                # Root query, config
│   ├── account.md             # Account CRUD, email, password
│   ├── session.md             # Login, logout, session management
│   ├── mfa.md                 # TOTP, recovery codes, MFA tickets
│   ├── onboarding.md          # Onboarding flow
│   ├── users.md               # User info, profiles, usernames
│   ├── relationships.md       # Friends, blocks
│   ├── direct-messaging.md    # DM channels
│   ├── bots.md                # Bot CRUD, invites
│   ├── servers.md             # Server CRUD, channels
│   ├── server-members.md      # Members, bans, invites
│   ├── server-permissions.md  # Roles, permissions
│   ├── server-customisation.md # Server emoji
│   ├── channels.md            # Channel CRUD, permissions
│   ├── messaging.md           # Messages, search, pins, bulk delete
│   ├── interactions.md        # Reactions
│   ├── groups.md              # Group channels
│   ├── voice.md               # Voice/call endpoints
│   ├── webhooks.md            # Webhook CRUD, execution
│   ├── invites.md             # Invite fetch, join, delete
│   ├── emojis.md              # Custom emoji CRUD
│   ├── sync.md                # Settings, unreads
│   ├── push.md                # Web push subscribe/unsubscribe
│   ├── safety.md              # Content reporting
│   └── policy.md              # Policy acknowledgement
├── websocket/
│   ├── connection.md          # Establishing connections
│   └── events.md              # All WS event types
├── cdn/
│   └── file-server.md         # Upload/download, tags, URLs
├── guides/
│   └── migration.md           # Revolt → Stoat migration
└── types/
    └── schemas.md             # All 139 data types
```

## Quick Links

- [Authentication](docs/concepts/authentication.md)
- [REST API Endpoints](docs/api/)
- [WebSocket Events](docs/websocket/events.md)
- [Data Types](docs/types/schemas.md)
- [Migration Guide](docs/guides/migration.md)

## Sources

- [Stoat Developer Docs](https://developers.stoat.chat)
- [OpenAPI Spec](https://github.com/stoatchat/javascript-client-api/blob/main/OpenAPI.json)
- [Stoat Backend Source](https://github.com/stoatchat/stoatchat)
- [Official JS SDK](https://github.com/stoatchat/javascript-client-sdk)
- [Official Python SDK](https://github.com/stoatchat/python-client-sdk)

## Existing Libraries

| Language | Library | URL |
|----------|---------|-----|
| JavaScript | stoat.js (official) | npm |
| Python | stoat.py (official) | [GitHub](https://github.com/stoatchat/python-client-sdk) |
| Go | revoltgo | GitHub |
| Rust | Rive, Seria, Stoat-rs | crates.io / GitHub |
| C# | Revolt.NET, StoatSharp | NuGet / GitHub |
| Swift | RevoltKit | GitHub |
| C | Ermine | Codeberg |
