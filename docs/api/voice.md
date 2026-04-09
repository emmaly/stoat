# Voice

Voice/call endpoints.

## Join Call

Join a voice channel or start a call in a DM/group.

```
POST /channels/{target}/join_call
Auth: Session Token
Path: target (string, required) — channel ID
Body: DataJoinCall
Response 200: CreateVoiceUserResponse
```

### DataJoinCall

```json
{
  // (empty object or minimal fields — specifics underdocumented)
}
```

### CreateVoiceUserResponse

```json
{
  "token": "string"  // voice server authentication token
}
```

## Stop Ring

Stop ringing a specific user (for DM/group calls).

```
PUT /channels/{target}/end_ring/{target_user}
Auth: Session Token
Path:
  - target (string, required) — channel ID
  - target_user (string, required) — user ID to stop ringing
Response 204: Success
```

## Voice Architecture

Voice is handled by a separate service:
- **Production URL:** Provided in the root config (`GET /`) under `features.voso.url` and `features.voso.ws`
- The voice server uses its own WebSocket connection separate from the main events WebSocket
- Voice was previously powered by "Vortex" (`vortex.revolt.chat`) but has been superseded by Voice Chats v2

## VoiceInformation

Appears on channels that have an active voice session:

```json
{
  "voice": {
    // Voice session metadata (structure not fully documented in spec)
  }
}
```

## Related Permissions

| Permission | Bit | Description |
|------------|-----|-------------|
| Connect | 1 << 30 | Join voice channels |
| Speak | 1 << 31 | Transmit audio |
| Video | 1 << 32 | Share video |
| MuteMembers | 1 << 33 | Server-mute others |
| DeafenMembers | 1 << 34 | Server-deafen others |
| MoveMembers | 1 << 35 | Move others between voice channels |
