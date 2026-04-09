# Rate Limits

Stoat uses a **fixed-window rate limiting** algorithm. Each endpoint belongs to a named bucket with a fixed call limit. Buckets replenish **10 seconds** after the first request in the window.

## Rate Limit Buckets

| Method | Path Pattern | Bucket Limit |
|--------|-------------|--------------|
| GET | `/users` | 20 |
| PATCH | `/users/:id` | 2 |
| GET | `/users/:id/default_avatar` | 255 |
| GET | `/bots` | 10 |
| GET | `/channels` | 15 |
| POST | `/channels/:id/messages` | 10 |
| GET | `/servers` | 5 |
| POST | `/auth` | 3 |
| DELETE | `/auth` | 255 |
| GET | `/safety` | 15 |
| POST | `/safety/report` | 3 |
| GET | `/swagger` | 100 |
| ANY | `/*` (catch-all) | 20 |

## Response Headers

Every response includes these rate limit headers:

| Header | Type | Description |
|--------|------|-------------|
| `X-RateLimit-Limit` | integer | Maximum calls allowed for this bucket |
| `X-RateLimit-Bucket` | string | Unique identifier for the bucket |
| `X-RateLimit-Remaining` | integer | Calls remaining in current window |
| `X-RateLimit-Reset-After` | integer | Milliseconds until bucket replenishes |

## Rate Limited Response

When the limit is exceeded, the server responds with:

```
HTTP 429 Too Many Requests

{
  "retry_after": <milliseconds until replenishment>
}
```

## Behavior Notes

- Buckets are **independent** — exhausting one bucket does not affect others
- All buckets within the same window reset on the same 10-second cycle
- The catch-all bucket (`/*`, limit 20) applies to any endpoint not explicitly listed
- Message sending (`POST /channels/:id/messages`) is limited to 10 per window — important for bot developers
- User editing (`PATCH /users/:id`) has an extremely tight limit of 2 per window
- Auth endpoints are tightly limited (3 per window) to prevent brute-force attacks

## Client Implementation Recommendations

1. Track `X-RateLimit-Remaining` per bucket
2. When `Remaining` reaches 0, wait for `Reset-After` milliseconds before retrying
3. On receiving 429, respect `retry_after` from the response body
4. For bots sending messages, implement a per-channel queue with 10/10s throttling
