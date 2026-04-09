# Server Customisation

## Fetch Server Emoji

List all custom emoji for a server.

```
GET /servers/{target}/emojis
Auth: Session Token
Path: target (string, required) — server ID
Response 200: Emoji[]
```

See [Emojis](emojis.md) for emoji CRUD operations and the Emoji type definition.
