package stoat

// Member represents a server member.
type Member struct {
	ID         MemberCompositeKey `json:"_id"`
	JoinedAt   string             `json:"joined_at"`
	Nickname   *string            `json:"nickname,omitempty"`
	Avatar     *File              `json:"avatar,omitempty"`
	Roles      []string           `json:"roles,omitempty"`
	Timeout    *string            `json:"timeout,omitempty"`
	CanPublish bool               `json:"can_publish,omitempty"`
	CanReceive bool               `json:"can_receive,omitempty"`
}

// MemberCompositeKey is a composite primary key consisting of server and user ID.
type MemberCompositeKey struct {
	Server string `json:"server"`
	User   string `json:"user"`
}

// DataMemberEdit is the request body for editing a member.
type DataMemberEdit struct {
	Nickname     *string        `json:"nickname,omitempty"`
	Avatar       *string        `json:"avatar,omitempty"`
	Roles        []string       `json:"roles,omitempty"`
	Timeout      *string        `json:"timeout,omitempty"`
	CanPublish   *bool          `json:"can_publish,omitempty"`
	CanReceive   *bool          `json:"can_receive,omitempty"`
	VoiceChannel *string        `json:"voice_channel,omitempty"`
	Remove       []FieldsMember `json:"remove,omitempty"`
}

// AllMemberResponse is the response from fetching all members.
type AllMemberResponse struct {
	Members []Member `json:"members"`
	Users   []User   `json:"users"`
}

// MemberResponse is the rich response from fetching a member with roles.
type MemberResponse struct {
	Member Member `json:"member"`
	Roles  []Role `json:"roles,omitempty"`
}

// MemberQueryResponse is the response from querying members by name.
type MemberQueryResponse struct {
	Members []Member `json:"members"`
	Users   []User   `json:"users"`
}

// FieldsMember is a string enum of optional fields on a member object that can be removed.
type FieldsMember string

const (
	FieldsMemberNickname     FieldsMember = "Nickname"
	FieldsMemberAvatar       FieldsMember = "Avatar"
	FieldsMemberRoles        FieldsMember = "Roles"
	FieldsMemberTimeout      FieldsMember = "Timeout"
	FieldsMemberCanReceive   FieldsMember = "CanReceive"
	FieldsMemberCanPublish   FieldsMember = "CanPublish"
	FieldsMemberJoinedAt     FieldsMember = "JoinedAt"
	FieldsMemberVoiceChannel FieldsMember = "VoiceChannel"
)

// ServerBan represents a server ban.
type ServerBan struct {
	ID     MemberCompositeKey `json:"_id"`
	Reason *string            `json:"reason,omitempty"`
}

// BannedUser contains just enough information to list a ban.
type BannedUser struct {
	ID            string `json:"_id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        *File  `json:"avatar,omitempty"`
}

// BanListResult is the response from fetching bans.
type BanListResult struct {
	Users []BannedUser `json:"users"`
	Bans  []ServerBan  `json:"bans"`
}

// DataBanCreate is the request body for banning a user.
type DataBanCreate struct {
	Reason               *string `json:"reason,omitempty"`
	DeleteMessageSeconds *int    `json:"delete_message_seconds,omitempty"`
}
