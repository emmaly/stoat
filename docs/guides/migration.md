# Revolt → Stoat Migration Guide

Stoat is a fork/continuation of the Revolt chat platform. This guide documents the changes relevant to API consumers.

## Domain Changes

| Service | Old (Revolt) | New (Stoat) | Redirect? |
|---------|-------------|-------------|-----------|
| REST API | `https://api.revolt.chat` | `https://stoat.chat/api` | Yes |
| REST API (alt) | `https://app.revolt.chat/api` | `https://stoat.chat/api` | Yes |
| REST API (alt) | `https://revolt.chat/api` | `https://stoat.chat/api` | Yes |
| WebSocket | `wss://ws.revolt.chat` | `wss://stoat.chat/events` | Yes |
| WebSocket (alt) | `wss://app.revolt.chat/events` | `wss://stoat.chat/events` | Yes |
| WebSocket (alt) | `wss://revolt.chat/events` | `wss://stoat.chat/events` | Yes |
| CDN (files) | `https://autumn.revolt.chat` | `https://cdn.stoatusercontent.com` | Yes |
| CDN (legacy) | `https://cdn.revoltusercontent.com` | `https://cdn.stoatusercontent.com` | Yes |
| Proxy | `https://jan.revolt.chat` | `https://external.stoatusercontent.com` | — |
| Voice | `https://vortex.revolt.chat` | Superseded by Voice Chats v2 | — |

## What Changed

### Branding

- "Revolt" → "Stoat" in all user-facing contexts
- GitHub org: `revoltchat` → `stoatchat`
- Developer docs: `developers.revolt.chat` → `developers.stoat.chat`
- Support: `support.revolt.chat` → `support.stoat.chat`
- Translation: `translate.revolt.chat` → `translate.stoat.chat`

### API

- **Base URL changed** — the API itself is structurally the same
- The OpenAPI spec still references "Revolt API" internally (as of v0.12.0)
- No endpoint paths changed
- No schema changes related to the rebrand
- Server config (`GET /`) returns Stoat URLs

### Client Recommendations

1. **Use dynamic discovery:** Always fetch `GET /` first and use the returned URLs for WebSocket, CDN, and proxy — never hardcode domains
2. **Update hardcoded URLs:** If your client has hardcoded Revolt domains, update them to Stoat domains
3. **Legacy redirects work:** Old Revolt URLs currently redirect, but don't rely on this long-term
4. **Voice migration:** If you used the old Vortex voice system, migrate to Voice Chats v2

## Status

This migration guide is marked as "not yet finished" in the official docs. Additional changes may be documented as the rebrand continues.
