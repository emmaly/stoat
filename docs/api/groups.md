# Groups

Group channel creation and member management.

## Create Group

```
POST /channels/create
Auth: Session Token
Body: DataCreateGroup
Response 200: Channel (Group type)
```

### DataCreateGroup

```json
{
  "name": "string",             // required
  "description": "string",      // optional
  "users": ["user_id", ...],    // optional — initial members to add
  "nsfw": false                  // optional
}
```

## Fetch Group Members

List all members of a group channel.

```
GET /channels/{target}/members
Auth: Session Token
Path: target (string, required) — group channel ID
Response 200: User[]
```

## Add Member to Group

```
PUT /channels/{group_id}/recipients/{member_id}
Auth: Session Token
Path:
  - group_id (string, required) — group channel ID
  - member_id (string, required) — user ID to add
Response 204: Success
```

## Remove Member from Group

```
DELETE /channels/{group_id}/recipients/{member_id}
Auth: Session Token
Path:
  - group_id (string, required) — group channel ID
  - member_id (string, required) — user ID to remove
Response 204: Success
```

## Notes

- The group owner can add/remove members
- To leave a group, use `DELETE /channels/{target}` (Close Channel)
- Groups use a flat permission integer, not the allow/deny override system
- Maximum group size is determined by server configuration
