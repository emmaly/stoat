# WebSocket Connection

The Stoat WebSocket provides real-time events for all platform activity. It is the primary mechanism for receiving updates about messages, channels, servers, users, and more.

## Connection URL

```
wss://stoat.chat/events?version=1&format=json
```

### Query Parameters

| Parameter | Required | Default | Description |
|-----------|----------|---------|-------------|
| `version` | Recommended | — | Protocol version (use `1`) |
| `format` | No | `json` | Packet format: `json` or `msgpack` |
| `token` | No | — | Session token or bot token (alternative to Authenticate event) |
| `ready` | No | all fields | Comma-separated list of fields to include in Ready event |

### Ready Event Field Options

The `ready` parameter accepts these values to customize the Ready payload:

- `users`
- `servers`
- `channels`
- `members`
- `emojis`
- `user_settings`
- `channel_unreads`
- `policy_changes`

## Authentication Methods

### Method 1: URL Query Parameter

Include the token directly in the connection URL:

```
wss://stoat.chat/events?version=1&format=json&token=<your-token>
```

### Method 2: Authenticate Event

Connect without a token, then send an Authenticate event:

```json
{"type": "Authenticate", "token": "<your-token>"}
```

## Connection Handshake

After successful authentication, the server sends events in this order:

1. **`Authenticated`** — Connection has been authenticated
2. **`Ready`** — Initial state dump with users, servers, channels, members, etc.
3. Server begins sending real-time events

If authentication fails, the server sends an **`Error`** event with one of:
- `InvalidSession` — Token is invalid or expired
- `OnboardingNotFinished` — User hasn't completed onboarding
- `AlreadyAuthenticated` — Connection already authenticated
- `InternalError` — Server error

## Heartbeat

**You must ping the server every 10–30 seconds** to keep the connection alive.

```json
// Client sends:
{"type": "Ping", "data": 1234567890}

// Server responds:
{"type": "Pong", "data": 1234567890}
```

The `data` field is echoed back — use it as a timestamp to measure latency.

## Subscriptions

By default, **normal users do not receive `UserUpdate` events** fanned out through servers. To receive them, explicitly subscribe:

```json
{"type": "Subscribe", "server_id": "<server_id>"}
```

### Subscription Rules

- Subscriptions **expire after 15 minutes**
- Maximum **5 active subscriptions** per client
- Resend every 10 minutes maximum per server
- Only send when the app/client is **in focus**
- No effect on bot sessions (bots receive all events automatically)

## Bots vs. Normal Users

| Behavior | Normal Users | Bots |
|----------|-------------|------|
| Receive all events | No — must subscribe to servers | Yes |
| UserUpdate fan-out | Only for subscribed servers | All servers |
| Subscribe needed | Yes | No (no effect) |

## Reconnection

The documentation does not specify an official reconnection protocol. Recommended client behavior:

1. On disconnect, attempt reconnection with exponential backoff
2. Re-authenticate on reconnect
3. The new `Ready` event provides fresh state
4. Re-establish any active subscriptions

## Legacy URLs

| Old URL | Status |
|---------|--------|
| `wss://ws.revolt.chat` | Redirects to `wss://stoat.chat/events` |
| `wss://app.revolt.chat/events` | Redirects to `wss://stoat.chat/events` |
| `wss://revolt.chat/events` | Redirects to `wss://stoat.chat/events` |
