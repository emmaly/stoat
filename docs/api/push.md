# Web Push

Web push notification subscription management.

## Subscribe

Register a web push subscription for the current session.

```
POST /push/subscribe
Auth: Session Token
Body: WebPushSubscription
Response 204: Success
```

### WebPushSubscription

```json
{
  "endpoint": "string",   // required — push service endpoint URL
  "p256dh": "string",     // required — P-256 ECDH public key
  "auth": "string"        // required — authentication secret
}
```

These values come from the browser's Push API (`PushSubscription` object).

## Unsubscribe

Remove the web push subscription for the current session.

```
POST /push/unsubscribe
Auth: Session Token
Response 204: Success
```

## Notes

- The VAPID public key is available from `GET /` (root config) in the `vapid` field
- Push notifications are sent by the `pushd` daemon service
- The subscription is tied to the current session — logging out removes it
