# Data Types (Schemas)

All 139 types from the Stoat API OpenAPI v3 spec (v0.12.0).

## Table of Contents

- [Core](#core)
- [Authentication](#authentication)
- [Users](#users)
- [Channels](#channels)
- [Messages](#messages)
- [Servers](#servers)
- [Members](#members)
- [Bans](#bans)
- [Roles and Permissions](#roles-and-permissions)
- [Bots](#bots)
- [Invites](#invites)
- [Emojis](#emojis)
- [Webhooks](#webhooks)
- [Files and Media](#files-and-media)
- [Safety](#safety)
- [Sync and Push](#sync-and-push)
- [Voice](#voice)
- [Misc](#misc)

---

## Core

### RevoltConfig

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `revolt` | string | Y | Revolt API Version |
| `features` | RevoltFeatures | Y | Features enabled on this Revolt node |
| `ws` | string | Y | WebSocket URL |
| `app` | string | Y | URL pointing to the client serving this node |
| `vapid` | string | Y | Web Push VAPID public key |
| `build` | BuildInformation | Y | Build information |

### RevoltFeatures

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `captcha` | CaptchaFeature | Y | hCaptcha configuration |
| `email` | boolean | Y | Whether email verification is enabled |
| `invite_only` | boolean | Y | Whether this server is invite only |
| `autumn` | Feature | Y | File server service configuration |
| `january` | Feature | Y | Proxy service configuration |
| `livekit` | VoiceFeature | Y | Voice server configuration |
| `limits` | LimitsConfig | Y | Limits |

### Feature

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `enabled` | boolean | Y | Whether the service is enabled |
| `url` | string | Y | URL pointing to the service |

### CaptchaFeature

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `enabled` | boolean | Y | Whether captcha is enabled |
| `key` | string | Y | Client key used for solving captcha |

### VoiceFeature

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `enabled` | boolean | Y | Whether voice is enabled |
| `nodes` | VoiceNode[] | Y | All livekit nodes |

### VoiceNode

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `name` | string | Y |  |
| `lat` | number | Y |  |
| `lon` | number | Y |  |
| `public_url` | string | Y |  |

### BuildInformation

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `commit_sha` | string | Y | Commit Hash |
| `commit_timestamp` | string | Y | Commit Timestamp |
| `semver` | string | Y | Git Semver |
| `origin_url` | string | Y | Git Origin URL |
| `timestamp` | string | Y | Build Timestamp |

### LimitsConfig

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `global` | GlobalLimits | Y | Global Limits |
| `new_user` | UserLimits | Y | New User Limits |
| `default` | UserLimits | Y | Default User Limits |

### GlobalLimits

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `group_size` | integer | Y | max group size |
| `message_embeds` | integer | Y | max message embeds |
| `message_replies` | integer | Y | max replies |
| `message_reactions` | integer | Y | max reactions per message |
| `server_emoji` | integer | Y | max server emoji |
| `server_roles` | integer | Y | max server roles |
| `server_channels` | integer | Y | max server channels |
| `body_limit_size` | integer | Y |  |
| `restrict_server_creation` | string[] | Y | restrict server creation to these users. if blank, all users can create servers |

### UserLimits

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `outgoing_friend_requests` | integer | Y | Max Outgoing Friend Requests |
| `bots` | integer | Y | Max Owned Bots |
| `message_length` | integer | Y | Max message content length |
| `message_attachments` | integer | Y | max message attachments |
| `servers` | integer | Y | max servers |
| `voice_quality` | integer | Y | max audio quality |
| `video` | boolean | Y | video streaming enabled |
| `video_resolution` | integer[] | Y | max video resolution (vertical, horizontal) |
| `video_aspect_ratio` | number[] | Y | min/max aspect ratios |
| `file_upload_size_limits` | map\<string, integer\> | Y |  |

### Error

Error information

**Tagged union:**

- **LabelMe**
- **AlreadyOnboarded**
- **UsernameTaken**
- **InvalidUsername**
- **DiscriminatorChangeRatelimited**
- **UnknownUser**
- **AlreadyFriends**
- **AlreadySentRequest**
- **Blocked**
- **BlockedByOther**
- **NotFriends**
- **TooManyPendingFriendRequests** — `max`: integer
- **UnknownChannel**
- **UnknownAttachment**
- **UnknownMessage**
- **CannotEditMessage**
- **CannotJoinCall**
- **TooManyAttachments** — `max`: integer
- **TooManyEmbeds** — `max`: integer
- **TooManyReplies** — `max`: integer
- **TooManyChannels** — `max`: integer
- **EmptyMessage**
- **PayloadTooLarge**
- **CannotRemoveYourself**
- **GroupTooLarge** — `max`: integer
- **AlreadyInGroup**
- **NotInGroup**
- **AlreadyPinned**
- **NotPinned**
- **InSlowmode** — `retry_after`: integer
- **CantCreateServers**
- **UnknownServer**
- **InvalidRole**
- **Banned**
- **TooManyServers** — `max`: integer
- **TooManyEmoji** — `max`: integer
- **TooManyRoles** — `max`: integer
- **AlreadyInServer**
- **CannotTimeoutYourself**
- **ReachedMaximumBots**
- **IsBot**
- **IsNotBot**
- **BotIsPrivate**
- **CannotReportYourself**
- **MissingPermission** — `permission`: string
- **MissingUserPermission** — `permission`: string
- **NotElevated**
- **NotPrivileged**
- **CannotGiveMissingPermissions**
- **NotOwner**
- **IsElevated**
- **DatabaseError** — `operation`: string, `collection`: string
- **InternalError**
- **InvalidOperation**
- **InvalidCredentials**
- **InvalidProperty**
- **InvalidSession**
- **InvalidFlagValue**
- **NotAuthenticated**
- **DuplicateNonce**
- **NotFound**
- **NoEffect**
- **FailedValidation** — `error`: string
- **LiveKitUnavailable**
- **NotAVoiceChannel**
- **AlreadyConnected**
- **NotConnected**
- **UnknownNode**
- **ProxyError**
- **FileTooSmall**
- **FileTooLarge** — `max`: integer
- **FileTypeNotAllowed**
- **ImageProcessingFailed**
- **NoEmbedData**
- **VosoUnavailable**
- **FeatureDisabled** — `feature`: string

## Authentication

### AccountInfo

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | string | Y |  |
| `email` | string | Y |  |

### DataCreateAccount

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `email` | string | Y | Valid email address |
| `password` | string | Y | Password |
| `invite` | string? |  | Invite code |
| `captcha` | string? |  | Captcha verification code |

### DataChangeEmail

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `email` | string | Y | Valid email address |
| `current_password` | string | Y | Current password |

### DataChangePassword

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `password` | string | Y | New password |
| `current_password` | string | Y | Current password |

### DataSendPasswordReset

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `email` | string | Y | Email associated with the account |
| `captcha` | string? |  | Captcha verification code |

### DataPasswordReset

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `token` | string | Y | Reset token |
| `password` | string | Y | New password |
| `remove_sessions` | boolean |  | Whether to logout all sessions |

### DataResendVerification

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `email` | string | Y | Email associated with the account |
| `captcha` | string? |  | Captcha verification code |

### DataAccountDeletion

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `token` | string | Y | Deletion token |

### DataLogin

Type: complex

### ResponseLogin

**Tagged union:**

- **Success** — `_id`: string, `user_id`: string, `token`: string, `name`: string, `last_seen`: string, `origin`: string, `subscription`: WebPushSubscription
- **MFA** — `ticket`: string, `allowed_methods`: MFAMethod[]
- **Disabled** — `user_id`: string

### ResponseVerify

Type: complex

### SessionInfo

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | string | Y |  |
| `name` | string | Y |  |

### DataEditSession

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `friendly_name` | string | Y | Session friendly name |

### MultiFactorStatus

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `email_otp` | boolean | Y |  |
| `trusted_handover` | boolean | Y |  |
| `email_mfa` | boolean | Y |  |
| `totp_mfa` | boolean | Y |  |
| `security_key_mfa` | boolean | Y |  |
| `recovery_active` | boolean | Y |  |

### MFAMethod

MFA method

**Enum:** `"Password"`, `"Recovery"`, `"Totp"`

### MFAResponse

MFA response

Type: complex

### MFATicket

Multi-factor auth ticket

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | string | Y | Unique Id |
| `account_id` | string | Y | Account Id |
| `token` | string | Y | Unique Token |
| `validated` | boolean | Y | Whether this ticket has been validated (can be used for account actions) |
| `authorised` | boolean | Y | Whether this ticket is authorised (can be used to log a user in) |
| `last_totp_code` | string? |  | TOTP code at time of ticket creation |

### ResponseTotpSecret

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `secret` | string | Y |  |

### Authifier Error

**Tagged union:**

- **IncorrectData** — `with`: string
- **DatabaseError** — `operation`: string, `with`: string
- **InternalError**
- **OperationFailed**
- **RenderFail**
- **MissingHeaders**
- **CaptchaFailed**
- **BlockedByShield**
- **InvalidSession**
- **UnverifiedAccount**
- **UnknownUser**
- **EmailFailed**
- **InvalidToken**
- **MissingInvite**
- **InvalidInvite**
- **InvalidCredentials**
- **CompromisedPassword**
- **ShortPassword**
- **Blacklisted**
- **LockedOut**
- **TotpAlreadyEnabled**
- **DisallowedMFAMethod**

## Users

### User

User

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | string | Y | Unique Id |
| `username` | string | Y | Username |
| `discriminator` | string | Y | Discriminator |
| `display_name` | string? |  | Display name |
| `avatar` | File? |  | Avatar attachment |
| `relations` | Relationship[] |  | Relationships with other users |
| `badges` | integer |  | Bitfield of user badges  https://docs.rs/revolt-models/latest/revolt_models/v0/e |
| `status` | UserStatus? |  | User's current status |
| `flags` | integer |  | Enum of user flags  https://docs.rs/revolt-models/latest/revolt_models/v0/enum.U |
| `privileged` | boolean |  | Whether this user is privileged |
| `bot` | BotInformation? |  | Bot information |
| `relationship` | RelationshipStatus | Y | Current session user's relationship with this user |
| `online` | boolean | Y | Whether this user is currently online |

### UserStatus

User's active status

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `text` | string? |  | Custom status text |
| `presence` | Presence? |  | Current presence option |

### UserProfile

User's profile

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `content` | string? |  | Text content on user's profile |
| `background` | File? |  | Background visible on user's profile |

### Presence

Presence status

**Tagged union:**

- `"Online"`
- `"Idle"`
- `"Focus"`
- `"Busy"`
- `"Invisible"`

### RelationshipStatus

User's relationship with another user (or themselves)

**Tagged union:**

- `"None"`
- `"User"`
- `"Friend"`
- `"Outgoing"`
- `"Incoming"`
- `"Blocked"`
- `"BlockedOther"`

### Relationship

Relationship entry indicating current status with other user

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | string | Y | Other user's Id |
| `status` | RelationshipStatus | Y | Relationship status with them |

### BotInformation

Bot information for if the user is a bot

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `owner` | string | Y | Id of the owner of this bot |

### FlagResponse

User flag reponse

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `flags` | integer | Y | Flags |

### MutualResponse

Mutual friends, servers, groups and DMs response

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `users` | string[] | Y | Array of mutual user IDs that both users are friends with |
| `servers` | string[] | Y | Array of mutual server IDs that both users are in |
| `channels` | string[] | Y | Array of mutual group and dm IDs that both users are in |

### DataEditUser

New user information

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `display_name` | string? |  | New display name |
| `avatar` | string? |  | Attachment Id for avatar |
| `status` | UserStatus? |  | New user status |
| `profile` | DataUserProfile? |  | New user profile data  This is applied as a partial. |
| `badges` | integer? |  | Bitfield of user badges |
| `flags` | integer? |  | Enum of user flags |
| `remove` | FieldsUser[] |  | Fields to remove from user object |

### DataUserProfile

New user profile data

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `content` | string? |  | Text to set as user profile description |
| `background` | string? |  | Attachment Id for background |

### DataChangeUsername

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `username` | string | Y | New username |
| `password` | string | Y | Current account password |

### DataSendFriendRequest

User lookup information

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `username` | string | Y | Username and discriminator combo separated by # |

### DataOnboard

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `username` | string | Y | New username which will be used to identify the user on the platform |

### DataHello

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `onboarding` | boolean | Y | Whether onboarding is required |

### FieldsUser

Optional fields on user object

**Tagged union:**

- `"Avatar"`
- `"StatusText"`
- `"StatusPresence"`
- `"ProfileContent"`
- `"ProfileBackground"`
- `"DisplayName"`
- `"Internal"`

## Channels

### Channel

Channel

**Tagged union:**

- **SavedMessages** — `_id`: string, `user`: string
- **DirectMessage** — `_id`: string, `active`: boolean, `recipients`: string[], `last_message_id`: string
- **Group** — `_id`: string, `name`: string, `owner`: string, `description`: string, `recipients`: string[], `icon`: File, `last_message_id`: string, `permissions`: integer, `nsfw`: boolean
- **TextChannel** — `_id`: string, `server`: string, `name`: string, `description`: string, `icon`: File, `last_message_id`: string, `default_permissions`: OverrideField, `role_permissions`: map\<string, OverrideField\>, `nsfw`: boolean, `voice`: VoiceInformation, `slowmode`: integer

### Category

Channel category

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `id` | string | Y | Unique ID for this category |
| `title` | string | Y | Title for this category |
| `channels` | string[] | Y | Channels in this category |

### ChannelCompositeKey

Composite primary key consisting of channel and user id

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `channel` | string | Y | Channel Id |
| `user` | string | Y | User Id |

### ChannelUnread

Channel Unread

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | ChannelCompositeKey | Y | Composite key pointing to a user's view of a channel |
| `last_id` | string? |  | Id of the last message read in this channel by a user |
| `mentions` | string[] |  | Array of message ids that mention the user |

### DataEditChannel

New webhook information

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `name` | string? |  | Channel name |
| `description` | string? |  | Channel description |
| `owner` | string? |  | Group owner |
| `icon` | string? |  | Icon  Provide an Autumn attachment Id. |
| `nsfw` | boolean? |  | Whether this channel is age-restricted |
| `archived` | boolean? |  | Whether this channel is archived |
| `voice` | VoiceInformation? |  | Voice Information for voice channels |
| `slowmode` | integer? |  | The channel's slow mode delay in seconds, up to 6 hours |
| `remove` | FieldsChannel[] |  | Fields to remove from channel |

### DataCreateGroup

Create new group

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `name` | string | Y | Group name |
| `description` | string? |  | Group description |
| `icon` | string? |  | Group icon |
| `users` | string[] |  | Array of user IDs to add to the group  Must be friends with these users. |
| `nsfw` | boolean? |  | Whether this group is age-restricted |

### DataCreateServerChannel

Create new server channel

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `type` | LegacyServerChannelType |  | Channel type |
| `name` | string | Y | Channel name |
| `description` | string? |  | Channel description |
| `nsfw` | boolean? |  | Whether this channel is age restricted |
| `voice` | VoiceInformation? |  | Voice Information for when this channel is also a voice channel |

### DataDefaultChannelPermissions

New default permissions

Type: complex

### DataSetRolePermissions

New role permissions

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `permissions` | Override | Y | Allow / deny values to set for this role |

### LegacyServerChannelType

Server Channel Type

**Tagged union:**

- `"Text"`
- `"Voice"`

### FieldsChannel

Optional fields on channel object

**Enum:** `"Description"`, `"Icon"`, `"DefaultPermissions"`, `"Voice"`

### VoiceInformation

Voice information for a channel

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `max_users` | integer? |  | Maximium amount of users allowed in the voice channel at once |

## Messages

### Message

Message

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | string | Y | Unique Id |
| `nonce` | string? |  | Unique value generated by client sending this message |
| `channel` | string | Y | Id of the channel this message was sent in |
| `author` | string | Y | Id of the user or webhook that sent this message |
| `user` | User? |  | The user that sent this message |
| `member` | Member? |  | The member that sent this message |
| `webhook` | MessageWebhook? |  | The webhook that sent this message |
| `content` | string? |  | Message content |
| `system` | SystemMessage? |  | System message |
| `attachments` | File[]? |  | Array of attachments |
| `edited` | ISO8601 Timestamp? |  | Time at which this message was last edited |
| `embeds` | Embed[]? |  | Attached embeds to this message |
| `mentions` | string[]? |  | Array of user ids mentioned in this message |
| `role_mentions` | string[]? |  | Array of role ids mentioned in this message |
| `replies` | string[]? |  | Array of message ids this message is replying to |
| `reactions` | map\<string, string[]\> |  | Hashmap of emoji IDs to array of user IDs |
| `interactions` | Interactions |  | Information about how this message should be interacted with |
| `masquerade` | Masquerade? |  | Name and / or avatar overrides for this message |
| `pinned` | boolean? |  | Whether or not the message in pinned |
| `flags` | integer |  | Bitfield of message flags  https://docs.rs/revolt-models/latest/revolt_models/v0 |

### DataMessageSend

Message to send

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `nonce` | string? |  | Unique token to prevent duplicate message sending  **This is deprecated and repl |
| `content` | string? |  | Message content to send |
| `attachments` | string[]? |  | Attachments to include in message |
| `replies` | ReplyIntent[]? |  | Messages to reply to |
| `embeds` | SendableEmbed[]? |  | Embeds to include in message  Text embed content contributes to the content leng |
| `masquerade` | Masquerade? |  | Masquerade to apply to this message |
| `interactions` | Interactions? |  | Information about how this message should be interacted with |
| `flags` | integer? |  | Bitfield of message flags  https://docs.rs/revolt-models/latest/revolt_models/v0 |

### DataEditMessage

Changes to make to message

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `content` | string? |  | New message content |
| `embeds` | SendableEmbed[]? |  | Embeds to include in the message |

### DataMessageSearch

Options for searching for messages

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `query` | string? |  | Full-text search query  See [MongoDB documentation](https://docs.mongodb.com/man |
| `pinned` | boolean? |  | Whether to only search for pinned messages, cannot be sent with `query`. |
| `limit` | integer? |  | Maximum number of messages to fetch |
| `before` | string? |  | Message id before which messages should be fetched |
| `after` | string? |  | Message id after which messages should be fetched |
| `sort` | MessageSort |  | Message sort direction  By default, it will be sorted by latest. |
| `include_users` | boolean? |  | Whether to include user (and member, if server channel) objects |

### SendableEmbed

Representation of a text embed before it is sent.

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `icon_url` | string? |  |  |
| `url` | string? |  |  |
| `title` | string? |  |  |
| `description` | string? |  |  |
| `media` | string? |  |  |
| `colour` | string? |  |  |

### Embed

Embed

**Tagged union:**

- **Website** — `url`: string, `original_url`: string, `special`: Special, `title`: string, `description`: string, `image`: Image, `video`: Video, `site_name`: string, `icon_url`: string, `colour`: string
- **Image** — `url`: string, `width`: integer, `height`: integer, `size`: ImageSize
- **Video** — `url`: string, `width`: integer, `height`: integer
- **Text** — `icon_url`: string, `url`: string, `title`: string, `description`: string, `media`: File, `colour`: string
- **None**

### ReplyIntent

What this message should reply to and how

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `id` | string | Y | Message Id |
| `mention` | boolean | Y | Whether this reply should mention the message's author |
| `fail_if_not_exists` | boolean? |  | Whether to error if the referenced message doesn't exist. Otherwise, send a mess |

### Masquerade

Name and / or avatar override information

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `name` | string? |  | Replace the display name shown on this message |
| `avatar` | string? |  | Replace the avatar shown on this message (URL to image file) |
| `colour` | string? |  | Replace the display role colour shown on this message  Must have `ManageRole` pe |

### Interactions

Information to guide interactions on this message

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `reactions` | string[]? |  | Reactions which should always appear and be distinct |
| `restrict_reactions` | boolean |  | Whether reactions should be restricted to the given list  Can only be set to tru |

### MessageWebhook

Information about the webhook bundled with Message

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `name` | string | Y |  |
| `avatar` | string? |  |  |

### SystemMessage

System Event

**Tagged union:**

- **text** — `content`: string
- **user_added** — `id`: string, `by`: string
- **user_remove** — `id`: string, `by`: string
- **user_joined** — `id`: string
- **user_left** — `id`: string
- **user_kicked** — `id`: string
- **user_banned** — `id`: string
- **channel_renamed** — `name`: string, `by`: string
- **channel_description_changed** — `by`: string
- **channel_icon_changed** — `by`: string
- **channel_ownership_changed** — `from`: string, `to`: string
- **message_pinned** — `id`: string, `by`: string
- **message_unpinned** — `id`: string, `by`: string
- **call_started** — `by`: string, `finished_at`: ISO8601 Timestamp

### MessageSort

Message Sort

Sort used for retrieving messages

**Tagged union:**

- `"Relevance"`
- `"Latest"`
- `"Oldest"`

### BulkMessageResponse

Bulk Message Response

Type: complex

### OptionsBulkDelete

Options for bulk deleting messages

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `ids` | string[] | Y | Message IDs |

## Servers

### Server

Server

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | string | Y | Unique Id |
| `owner` | string | Y | User id of the owner |
| `name` | string | Y | Name of the server |
| `description` | string? |  | Description for the server |
| `channels` | string[] | Y | Channels within this server |
| `categories` | Category[]? |  | Categories for this server |
| `system_messages` | SystemMessageChannels? |  | Configuration for sending system event messages |
| `roles` | map\<string, Role\> |  | Roles for this server |
| `default_permissions` | integer | Y | Default set of server and channel permissions |
| `icon` | File? |  | Icon attachment |
| `banner` | File? |  | Banner attachment |
| `flags` | integer |  | Bitfield of server flags |
| `nsfw` | boolean |  | Whether this server is flagged as not safe for work |
| `analytics` | boolean |  | Whether to enable analytics |
| `discoverable` | boolean |  | Whether this server should be publicly discoverable |

### DataCreateServer

Information about new server to create

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `name` | string | Y | Server name |
| `description` | string? |  | Server description |
| `nsfw` | boolean? |  | Whether this server is age-restricted |

### DataEditServer

New server information

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `name` | string? |  | Server name |
| `description` | string? |  | Server description |
| `icon` | string? |  | Attachment Id for icon |
| `banner` | string? |  | Attachment Id for banner |
| `categories` | Category[]? |  | Category structure for server |
| `system_messages` | SystemMessageChannels? |  | System message configuration |
| `flags` | integer? |  | Bitfield of server flags |
| `discoverable` | boolean? |  | Whether this server is public and should show up on [Revolt Discover](https://rv |
| `analytics` | boolean? |  | Whether analytics should be collected for this server  Must be enabled in order  |
| `owner` | string? |  | User id of the new owner |
| `remove` | FieldsServer[] |  | Fields to remove from server object |

### CreateServerLegacyResponse

Information returned when creating server

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `server` | Server | Y | Server object |
| `channels` | Channel[] | Y | Default channels |

### FetchServerResponse

Fetch server information

Type: complex

### SystemMessageChannels

System message channel assignments

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `user_joined` | string? |  | ID of channel to send user join messages in |
| `user_left` | string? |  | ID of channel to send user left messages in |
| `user_kicked` | string? |  | ID of channel to send user kicked messages in |
| `user_banned` | string? |  | ID of channel to send user banned messages in |

### FieldsServer

Optional fields on server object

**Enum:** `"Description"`, `"Categories"`, `"SystemMessages"`, `"Icon"`, `"Banner"`

## Members

### Member

Server Member

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | MemberCompositeKey | Y | Unique member id |
| `joined_at` | ISO8601 Timestamp | Y | Time at which this user joined the server |
| `nickname` | string? |  | Member's nickname |
| `avatar` | File? |  | Avatar attachment |
| `roles` | string[] |  | Member's roles |
| `timeout` | ISO8601 Timestamp? |  | Timestamp this member is timed out until |
| `can_publish` | boolean |  | Whether the member is server-wide voice muted |
| `can_receive` | boolean |  | Whether the member is server-wide voice deafened |

### MemberCompositeKey

Composite primary key consisting of server and user id

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `server` | string | Y | Server Id |
| `user` | string | Y | User Id |

### DataMemberEdit

New member information

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `nickname` | string? |  | Member nickname |
| `avatar` | string? |  | Attachment Id to set for avatar |
| `roles` | string[]? |  | Array of role ids |
| `timeout` | ISO8601 Timestamp? |  | Timestamp this member is timed out until |
| `can_publish` | boolean? |  | server-wide voice muted |
| `can_receive` | boolean? |  | server-wide voice deafened |
| `voice_channel` | string? |  | voice channel to move to if already in a voice channel |
| `remove` | FieldsMember[] |  | Fields to remove from channel object |

### AllMemberResponse

Response with all members

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `members` | Member[] | Y | List of members |
| `users` | User[] | Y | List of users |

### MemberResponse

Member response

Type: complex

### MemberQueryResponse

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `members` | Member[] | Y | List of members |
| `users` | User[] | Y | List of users |

### FieldsMember

Optional fields on server member object

**Enum:** `"Nickname"`, `"Avatar"`, `"Roles"`, `"Timeout"`, `"CanReceive"`, `"CanPublish"`, `"JoinedAt"`, `"VoiceChannel"`

## Bans

### ServerBan

Server Ban

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | MemberCompositeKey | Y | Unique member id |
| `reason` | string? |  | Reason for ban creation |

### BannedUser

Just enough information to list a ban

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | string | Y | Id of the banned user |
| `username` | string | Y | Username of the banned user |
| `discriminator` | string | Y | Discriminator of the banned user |
| `avatar` | File? |  | Avatar of the banned user |

### BanListResult

Ban list result

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `users` | BannedUser[] | Y | Users objects |
| `bans` | ServerBan[] | Y | Ban objects |

### DataBanCreate

Information for new server ban

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `reason` | string? |  | Ban reason |
| `delete_message_seconds` | integer? |  | Messages to delete in seconds |

## Roles and Permissions

### Role

Role

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | string | Y | Unique Id |
| `name` | string | Y | Role name |
| `permissions` | OverrideField | Y | Permissions available to this role |
| `colour` | string? |  | Colour used for this role  This can be any valid CSS colour |
| `hoist` | boolean |  | Whether this role should be shown separately on the member sidebar |
| `rank` | integer |  | Ranking of this role |

### DataCreateRole

Information about new role to create

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `name` | string | Y | Role name |
| `rank` | integer? |  | Ranking position  Smaller values take priority.  **Removed** - no effect, use th |

### DataEditRole

New role information

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `name` | string? |  | Role name |
| `colour` | string? |  | Role colour |
| `hoist` | boolean? |  | Whether this role should be displayed separately |
| `rank` | integer? |  | Ranking position  **Removed** - no effect, use the edit server role positions ro |
| `remove` | FieldsRole[] |  | Fields to remove from role object |

### DataEditRoleRanks

New role positions

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `ranks` | string[] | Y |  |

### NewRoleResponse

Response after creating new role

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `id` | string | Y | Id of the role |
| `role` | Role | Y | New role |

### Override

Representation of a single permission override

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `allow` | integer | Y | Allow bit flags |
| `deny` | integer | Y | Disallow bit flags |

### OverrideField

Representation of a single permission override as it appears on models and in the database

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `a` | integer | Y | Allow bit flags |
| `d` | integer | Y | Disallow bit flags |

### DataPermissionsValue

Data permissions Value - contains allow

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `permissions` | integer | Y |  |

### DataSetServerRolePermission

New role permissions

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `permissions` | Override | Y | Allow / deny values for the role in this server. |

### FieldsRole

Optional fields on server object

**Enum:** `"Colour"`

## Bots

### Bot

Bot

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | string | Y | Bot Id |
| `owner` | string | Y | User Id of the bot owner |
| `token` | string | Y | Token used to authenticate requests for this bot |
| `public` | boolean | Y | Whether the bot is public (may be invited by anyone) |
| `analytics` | boolean |  | Whether to enable analytics |
| `discoverable` | boolean |  | Whether this bot should be publicly discoverable |
| `interactions_url` | string |  | Reserved; URL for handling interactions |
| `terms_of_service_url` | string |  | URL for terms of service |
| `privacy_policy_url` | string |  | URL for privacy policy |
| `flags` | integer |  | Enum of bot flags |

### PublicBot

Public Bot

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | string | Y | Bot Id |
| `username` | string | Y | Bot Username |
| `avatar` | string |  | Profile Avatar |
| `description` | string |  | Profile Description |

### DataCreateBot

Bot Details

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `name` | string | Y | Bot username |

### DataEditBot

New Bot Details

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `name` | string? |  | Bot username |
| `public` | boolean? |  | Whether the bot can be added by anyone |
| `analytics` | boolean? |  | Whether analytics should be gathered for this bot  Must be enabled in order to s |
| `interactions_url` | string? |  | Interactions URL |
| `remove` | FieldsBot[] |  | Fields to remove from bot object |

### FetchBotResponse

Bot Response

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `bot` | Bot | Y | Bot object |
| `user` | User | Y | User object |

### BotWithUserResponse

Bot with user response

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `user` | User | Y |  |
| `_id` | string | Y | Bot Id |
| `owner` | string | Y | User Id of the bot owner |
| `token` | string | Y | Token used to authenticate requests for this bot |
| `public` | boolean | Y | Whether the bot is public (may be invited by anyone) |
| `analytics` | boolean |  | Whether to enable analytics |
| `discoverable` | boolean |  | Whether this bot should be publicly discoverable |
| `interactions_url` | string |  | Reserved; URL for handling interactions |
| `terms_of_service_url` | string |  | URL for terms of service |
| `privacy_policy_url` | string |  | URL for privacy policy |
| `flags` | integer |  | Enum of bot flags |

### OwnedBotsResponse

Owned Bots Response

Both lists are sorted by their IDs.

TODO: user should be in bot object

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `bots` | Bot[] | Y | Bot objects |
| `users` | User[] | Y | User objects |

### InviteBotDestination

Where we are inviting a bot to

Type: complex

### FieldsBot

Optional fields on bot object

**Enum:** `"Token"`, `"InteractionsURL"`

## Invites

### Invite

Invite

**Tagged union:**

- **Server** — `_id`: string, `server`: string, `creator`: string, `channel`: string
- **Group** — `_id`: string, `creator`: string, `channel`: string

### InviteResponse

Public invite response

**Tagged union:**

- **Server** — `code`: string, `server_id`: string, `server_name`: string, `server_icon`: File, `server_banner`: File, `server_flags`: integer, `channel_id`: string, `channel_name`: string, `channel_description`: string, `user_name`: string, `user_avatar`: File, `member_count`: integer
- **Group** — `code`: string, `channel_id`: string, `channel_name`: string, `channel_description`: string, `user_name`: string, `user_avatar`: File

### InviteJoinResponse

Invite join response

**Tagged union:**

- **Server** — `channels`: Channel[], `server`: Server
- **Group** — `channel`: Channel, `users`: User[]

## Emojis

### Emoji

Emoji

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | string | Y | Unique Id |
| `parent` | EmojiParent | Y | What owns this emoji |
| `creator_id` | string | Y | Uploader user id |
| `name` | string | Y | Emoji name |
| `animated` | boolean |  | Whether the emoji is animated |
| `nsfw` | boolean |  | Whether the emoji is marked as nsfw |

### EmojiParent

Parent Id of the emoji

**Tagged union:**

- **Server** — `id`: string
- **Detached**

### DataCreateEmoji

Create a new emoji

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `name` | string | Y | Server name |
| `parent` | EmojiParent | Y | Parent information |
| `nsfw` | boolean |  | Whether the emoji is mature |

## Webhooks

### Webhook

Webhook

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `id` | string | Y | Webhook Id |
| `name` | string | Y | The name of the webhook |
| `avatar` | File? |  | The avatar of the webhook |
| `creator_id` | string | Y | User that created this webhook |
| `channel_id` | string | Y | The channel this webhook belongs to |
| `permissions` | integer | Y | The permissions for the webhook |
| `token` | string? |  | The private token for the webhook |

### ResponseWebhook

Webhook information

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `id` | string | Y | Webhook Id |
| `name` | string | Y | Webhook name |
| `avatar` | string? |  | Avatar ID |
| `channel_id` | string | Y | The channel this webhook belongs to |
| `permissions` | integer | Y | The permissions for the webhook |

### CreateWebhookBody

Information for the webhook

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `name` | string | Y |  |
| `avatar` | string? |  |  |

### DataEditWebhook

New webhook information

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `name` | string? |  | Webhook name |
| `avatar` | string? |  | Avatar ID |
| `permissions` | integer? |  | Webhook permissions |
| `remove` | FieldsWebhook[] |  | Fields to remove from webhook |

### FieldsWebhook

Optional fields on webhook object

**Enum:** `"Avatar"`

## Files and Media

### File

File

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `_id` | string | Y | Unique Id |
| `tag` | string | Y | Tag / bucket this file was uploaded to |
| `filename` | string | Y | Original filename |
| `metadata` | Metadata | Y | Parsed metadata of this file |
| `content_type` | string | Y | Raw content type of this file |
| `size` | integer | Y | Size of this file (in bytes) |
| `deleted` | boolean? |  | Whether this file was deleted |
| `reported` | boolean? |  | Whether this file was reported |
| `message_id` | string? |  |  |
| `user_id` | string? |  |  |
| `server_id` | string? |  |  |
| `object_id` | string? |  | Id of the object this file is associated with |

### Metadata

Metadata associated with a file

**Tagged union:**

- **File**
- **Text**
- **Image** — `width`: integer, `height`: integer, `thumbhash`: integer[], `animated`: boolean
- **Video** — `width`: integer, `height`: integer
- **Audio**

### Image

Image

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `url` | string | Y | URL to the original image |
| `width` | integer | Y | Width of the image |
| `height` | integer | Y | Height of the image |
| `size` | ImageSize | Y | Positioning and size |

### Video

Video

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `url` | string | Y | URL to the original video |
| `width` | integer | Y | Width of the video |
| `height` | integer | Y | Height of the video |

### ImageSize

Image positioning and size

**Tagged union:**

- `"Large"`
- `"Preview"`

### Special

Information about special remote content

**Tagged union:**

- **None**
- **GIF**
- **YouTube** — `id`: string, `timestamp`: string
- **Lightspeed** — `content_type`: LightspeedType, `id`: string
- **Twitch** — `content_type`: TwitchType, `id`: string
- **Spotify** — `content_type`: string, `id`: string
- **Soundcloud**
- **Bandcamp** — `content_type`: BandcampType, `id`: string
- **AppleMusic** — `album_id`: string, `track_id`: string
- **Streamable** — `id`: string

### BandcampType

Type of remote Bandcamp content

**Enum:** `"Album"`, `"Track"`

### LightspeedType

Type of remote Lightspeed.tv content

**Enum:** `"Channel"`

### TwitchType

Type of remote Twitch content

**Enum:** `"Channel"`, `"Video"`, `"Clip"`

## Safety

### DataReportContent

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `content` | ReportedContent | Y | Content being reported |
| `additional_context` | string |  | Additional report description |

### ReportedContent

The content being reported

**Tagged union:**

- **Message** — `id`: string, `report_reason`: ContentReportReason
- **Server** — `id`: string, `report_reason`: ContentReportReason
- **User** — `id`: string, `report_reason`: UserReportReason, `message_id`: string

### ContentReportReason

Reason for reporting content (message or server)

**Tagged union:**

- `"NoneSpecified"`
- `"Illegal"`
- `"IllegalGoods"`
- `"IllegalExtortion"`
- `"IllegalPornography"`
- `"IllegalHacking"`
- `"ExtremeViolence"`
- `"PromotesHarm"`
- `"UnsolicitedSpam"`
- `"Raid"`
- `"SpamAbuse"`
- `"ScamsFraud"`
- `"Malware"`
- `"Harassment"`

### UserReportReason

Reason for reporting a user

**Tagged union:**

- `"NoneSpecified"`
- `"UnsolicitedSpam"`
- `"SpamAbuse"`
- `"InappropriateProfile"`
- `"Impersonation"`
- `"BanEvasion"`
- `"Underage"`

## Sync and Push

### OptionsFetchSettings

Options for fetching settings

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `keys` | string[] | Y | Keys to fetch |

### WebPushSubscription

Web Push subscription

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `endpoint` | string | Y |  |
| `p256dh` | string | Y |  |
| `auth` | string | Y |  |

## Voice

### DataJoinCall

Join a voice channel

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `node` | string? |  | Name of the node to join |
| `force_disconnect` | boolean? |  | Whether to force disconnect any other existing voice connections  Useful for dis |
| `recipients` | string[]? |  | Users which should be notified of the call starting  Only used when the user is  |

### CreateVoiceUserResponse

Voice server token response

| Field | Type | Req | Description |
|-------|------|-----|-------------|
| `token` | string | Y | Token for authenticating with the voice server |
| `url` | string | Y | Url of the livekit server to connect to |

## Misc

### Id

String

### ISO8601 Timestamp

ISO8601 formatted timestamp

String (format: date-time)
