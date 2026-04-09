# Core

## Query Node

Fetch server configuration and feature flags.

```
GET /
Auth: None
Response 200: RevoltConfig
```

### Response: RevoltConfig

```json
{
  "revolt": "0.12.0",
  "features": {
    "captcha": {
      "enabled": true,
      "key": "<hcaptcha_site_key>"
    },
    "email": true,
    "invite_only": false,
    "autumn": {
      "enabled": true,
      "url": "https://cdn.stoatusercontent.com"
    },
    "january": {
      "enabled": true,
      "url": "https://external.stoatusercontent.com"
    },
    "voso": {
      "enabled": true,
      "url": "<voice_server_url>",
      "ws": "<voice_ws_url>"
    }
  },
  "ws": "wss://stoat.chat/events",
  "app": "https://stoat.chat",
  "vapid": "<vapid_public_key>",
  "build": {
    "commit_sha": "string",
    "commit_timestamp": "string",
    "semver": "string",
    "origin_url": "string",
    "timestamp": "string"
  }
}
```

This is the first endpoint any client should call — it provides:
- WebSocket URL for real-time events
- CDN URL for file uploads/downloads
- Proxy URL for external content
- Voice server configuration
- CAPTCHA site key (if CAPTCHA is enabled)
- VAPID key for web push notifications
- Whether the instance is invite-only
