# Interactions (Reactions)

## Add Reaction to Message

```
PUT /channels/{target}/messages/{msg}/reactions/{emoji}
Auth: Session Token
Path:
  - target (string, required) — channel ID
  - msg (string, required) — message ID
  - emoji (string, required) — emoji ID (custom emoji ID or Unicode emoji)
Response 204: Success
```

## Remove Reaction(s) from Message

Remove your own reaction, or optionally remove all of a specific emoji or a specific user's reaction.

```
DELETE /channels/{target}/messages/{msg}/reactions/{emoji}
Auth: Session Token
Path:
  - target (string, required) — channel ID
  - msg (string, required) — message ID
  - emoji (string, required) — emoji ID
Query:
  - user_id (string, optional) — remove this user's reaction (requires ManageMessages)
  - remove_all (boolean, optional) — remove all reactions of this emoji (requires ManageMessages)
Response 204: Success
```

## Remove All Reactions from Message

Clear all reactions from a message.

```
DELETE /channels/{target}/messages/{msg}/reactions
Auth: Session Token
Path:
  - target (string, required) — channel ID
  - msg (string, required) — message ID
Response 204: Success
```

## Reaction Data Model

Reactions are stored on the Message object as a map:

```json
{
  "reactions": {
    "emoji_id_or_unicode": ["user_id_1", "user_id_2"],
    "👍": ["user_id_3"]
  }
}
```

## Interaction Controls

Messages can restrict which reactions are allowed via the `interactions` field:

```json
{
  "interactions": {
    "reactions": ["emoji_id_1", "emoji_id_2"],
    "restrict_reactions": true
  }
}
```

When `restrict_reactions` is true, only the listed emoji can be used as reactions on that message.
