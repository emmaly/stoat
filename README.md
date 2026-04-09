# stoat

Go client library for the [Stoat](https://stoat.chat) chat platform API.

```
go get github.com/emmaly/stoat
```

## Coverage

| Area | Coverage |
|------|----------|
| REST API | 121 endpoints across 21 endpoint groups |
| Data Types | 139 schema types (including tagged unions) |
| WebSocket | 40 event types, connection management, event dispatch |
| CDN | File upload/download, tag constants |

Targets Stoat API v0.12.0. Built clean-room from the [OpenAPI spec](docs/OpenAPI.json) and [developer docs](https://developers.stoat.chat).

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/emmaly/stoat"
)

func main() {
    c, err := stoat.New("https://stoat.chat/api")
    if err != nil {
        log.Fatal(err)
    }
    c.SetBotToken("your-bot-token")

    // Fetch server config
    cfg, err := c.QueryNode(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("API version:", cfg.Revolt)
    fmt.Println("WebSocket:", cfg.WS)
}
```

## Authentication

```go
// Bot token (no expiry, no session management)
c.SetBotToken("bot-token-from-settings")

// User session (obtained via login)
resp, err := c.Login(ctx, stoat.DataLogin{
    Email:    "user@example.com",
    Password: "password",
})
// resp is a stoat.ResponseLogin — type-switch on the variant:
switch v := resp.(type) {
case *stoat.LoginSuccess:
    c.SetSessionToken(v.Token)
case *stoat.LoginMFA:
    // Handle MFA flow with v.Ticket
}
```

## WebSocket Events

```go
import "github.com/emmaly/stoat/ws"

conn, err := ws.Dial(ctx, "wss://stoat.chat/events", "your-token")
if err != nil {
    log.Fatal(err)
}
defer conn.Close()

// Implement ws.EventHandler (or embed ws.DefaultEventHandler)
type myHandler struct{ ws.DefaultEventHandler }

func (h *myHandler) OnMessage(e ws.MessageEvent) {
    fmt.Printf("[%s] %s: %s\n", e.Channel, e.Author, *e.Content)
}

conn.Listen(ctx, &myHandler{})
```

## CDN File Uploads

```go
import "github.com/emmaly/stoat/cdn"

cdnClient, _ := cdn.New("https://cdn.stoatusercontent.com")
cdnClient.SetBotToken("your-bot-token")

fileID, err := cdnClient.Upload(ctx, cdn.TagAttachments, "photo.png", file)
// Use fileID in message attachments
```

## Packages

| Package | Import | Description |
|---------|--------|-------------|
| `stoat` | `github.com/emmaly/stoat` | REST client, all API types |
| `stoat/ws` | `github.com/emmaly/stoat/ws` | WebSocket connection and events |
| `stoat/cdn` | `github.com/emmaly/stoat/cdn` | CDN file upload/download |

## Tagged Unions

Several API types are discriminated unions represented as Go interfaces:

| Type | Discriminator | Variants |
|------|---------------|----------|
| `Channel` | `channel_type` | SavedMessages, DirectMessage, Group, TextChannel, VoiceChannel |
| `Embed` | `type` | Website, Image, Video, Text, None |
| `SystemMessage` | `type` | 14 variants (text, user_added, user_joined, etc.) |
| `Metadata` | `type` | File, Text, Image, Video, Audio |
| `ResponseLogin` | `result` | Success, MFA, Disabled |
| `InviteResponse` | `type` | Server |
| `InviteJoinResponse` | `type` | Server |
| `EmojiParent` | `type` | Server, Detached |

Use `Raw*` wrapper types (e.g., `RawChannel`) for JSON unmarshaling, then access `.Value` for the interface.

## API Documentation

Full API documentation is in [`docs/`](docs/):

- [`docs/concepts/`](docs/concepts/) — Authentication, permissions, rate limits, file uploads
- [`docs/api/`](docs/api/) — All REST endpoint documentation
- [`docs/websocket/`](docs/websocket/) — WebSocket connection and event protocol
- [`docs/cdn/`](docs/cdn/) — CDN file server
- [`docs/types/`](docs/types/) — All 139 data type schemas
- [`docs/guides/`](docs/guides/) — Revolt to Stoat migration guide

## License

AGPLv3 — matching the Stoat platform license.
