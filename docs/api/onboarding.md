# Onboarding

New users must complete onboarding (choose a username) before using the platform.

## Check Onboarding Status

```
GET /onboard/hello
Auth: Session Token
Response 200: DataHello
```

### DataHello

```json
{
  "onboarding": true  // whether onboarding is required
}
```

## Complete Onboarding

```
POST /onboard/complete
Auth: Session Token
Body: DataOnboard
Response 200: User
```

### DataOnboard

```json
{
  "username": "string"  // required — chosen username
}
```

If the WebSocket returns an `OnboardingNotFinished` error, the client must call this endpoint before proceeding.
