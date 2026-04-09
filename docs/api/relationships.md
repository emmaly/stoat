# Relationships

Friend requests, blocking, and relationship management.

## Send Friend Request

Send a friend request by username.

```
POST /users/friend
Auth: Session Token
Body: DataSendFriendRequest
Response 200: User
```

### DataSendFriendRequest

```json
{
  "username": "string"  // required — target username
}
```

## Accept Friend Request

Accept an incoming friend request (or send one if none exists).

```
PUT /users/{target}/friend
Auth: Session Token
Path: target (string, required) — user ID
Response 200: User
```

## Deny Friend Request / Remove Friend

```
DELETE /users/{target}/friend
Auth: Session Token
Path: target (string, required) — user ID
Response 200: User
```

## Block User

```
PUT /users/{target}/block
Auth: Session Token
Path: target (string, required) — user ID
Response 200: User
```

## Unblock User

```
DELETE /users/{target}/block
Auth: Session Token
Path: target (string, required) — user ID
Response 200: User
```

## Relationship States

The `relationship` field on User objects indicates the current authenticated user's relationship with that user:

| Status | Description |
|--------|-------------|
| `None` | No relationship |
| `User` | This is yourself |
| `Friend` | Mutual friends |
| `Outgoing` | You sent a friend request |
| `Incoming` | They sent you a friend request |
| `Blocked` | You blocked this user |
| `BlockedOther` | This user blocked you |

## Relationship Type

```json
{
  "_id": "string",           // other user's ID
  "status": "Friend"         // RelationshipStatus enum
}
```

The `relations` array on the User object contains all relationships.
