# Policy

## Acknowledge Policy Changes

Acknowledge updated terms of service or privacy policy.

```
POST /policy/acknowledge
Auth: Session Token
Response 204: Success
```

The `policy_changes` field in the WebSocket `Ready` event indicates whether the user needs to acknowledge updated policies. Clients should prompt the user and call this endpoint.
