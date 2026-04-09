package stoat

import (
	"encoding/json"
	"testing"
)

func TestMemberUnmarshal(t *testing.T) {
	data := `{
		"_id": {"server": "srv01", "user": "user01"},
		"joined_at": "2024-01-15T10:30:00Z",
		"nickname": "Cool Nick",
		"avatar": {"_id": "av01", "tag": "avatars", "filename": "a.png", "metadata": {"type": "Image", "width": 64, "height": 64}, "content_type": "image/png", "size": 512},
		"roles": ["role01", "role02"],
		"timeout": "2024-01-16T10:30:00Z"
	}`
	var m Member
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if m.ID.Server != "srv01" || m.ID.User != "user01" {
		t.Errorf("ID = %+v", m.ID)
	}
	if m.JoinedAt != "2024-01-15T10:30:00Z" {
		t.Errorf("JoinedAt = %q", m.JoinedAt)
	}
	if m.Nickname == nil || *m.Nickname != "Cool Nick" {
		t.Errorf("Nickname = %v", m.Nickname)
	}
	if m.Avatar == nil || m.Avatar.ID != "av01" {
		t.Errorf("Avatar = %v", m.Avatar)
	}
	if len(m.Roles) != 2 {
		t.Errorf("Roles = %v", m.Roles)
	}
	if m.Timeout == nil || *m.Timeout != "2024-01-16T10:30:00Z" {
		t.Errorf("Timeout = %v", m.Timeout)
	}
}

func TestMemberCompositeKeyMarshal(t *testing.T) {
	k := MemberCompositeKey{Server: "srv01", User: "user01"}
	b, err := json.Marshal(k)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var k2 MemberCompositeKey
	if err := json.Unmarshal(b, &k2); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if k2 != k {
		t.Errorf("roundtrip mismatch: %+v", k2)
	}
}

func TestAllMemberResponseUnmarshal(t *testing.T) {
	data := `{
		"members": [
			{"_id": {"server": "srv01", "user": "user01"}, "joined_at": "2024-01-01T00:00:00Z"}
		],
		"users": [
			{"_id": "user01", "username": "alice", "discriminator": "0001", "relationship": "None", "online": false}
		]
	}`
	var resp AllMemberResponse
	if err := json.Unmarshal([]byte(data), &resp); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if len(resp.Members) != 1 || resp.Members[0].ID.User != "user01" {
		t.Errorf("Members = %+v", resp.Members)
	}
	if len(resp.Users) != 1 || resp.Users[0].ID != "user01" {
		t.Errorf("Users = %+v", resp.Users)
	}
}

func TestMemberResponseUnmarshal(t *testing.T) {
	data := `{
		"member": {"_id": {"server": "srv01", "user": "user01"}, "joined_at": "2024-01-01T00:00:00Z"},
		"roles": [{"name": "Admin", "permissions": {"a": 255, "d": 0}, "hoist": true, "rank": 1}]
	}`
	var resp MemberResponse
	if err := json.Unmarshal([]byte(data), &resp); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if resp.Member.ID.Server != "srv01" {
		t.Errorf("Member.ID = %+v", resp.Member.ID)
	}
	if len(resp.Roles) != 1 || resp.Roles[0].Name != "Admin" {
		t.Errorf("Roles = %+v", resp.Roles)
	}
}

func TestDataMemberEditMarshal(t *testing.T) {
	nick := "NewNick"
	d := DataMemberEdit{
		Nickname: &nick,
		Remove:   []FieldsMember{FieldsMemberNickname},
	}
	b, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var m map[string]any
	json.Unmarshal(b, &m)
	if m["nickname"] != "NewNick" {
		t.Errorf("nickname = %v", m["nickname"])
	}
	r := m["remove"].([]any)
	if r[0] != "Nickname" {
		t.Errorf("remove[0] = %v", r[0])
	}
}

func TestServerBanUnmarshal(t *testing.T) {
	data := `{
		"_id": {"server": "srv01", "user": "user01"},
		"reason": "spam"
	}`
	var ban ServerBan
	if err := json.Unmarshal([]byte(data), &ban); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if ban.ID.Server != "srv01" || ban.ID.User != "user01" {
		t.Errorf("ID = %+v", ban.ID)
	}
	if ban.Reason == nil || *ban.Reason != "spam" {
		t.Errorf("Reason = %v", ban.Reason)
	}
}

func TestBanListResultUnmarshal(t *testing.T) {
	data := `{
		"users": [{"_id": "user01", "username": "alice", "discriminator": "0001"}],
		"bans": [{"_id": {"server": "srv01", "user": "user01"}, "reason": "spam"}]
	}`
	var resp BanListResult
	if err := json.Unmarshal([]byte(data), &resp); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if len(resp.Users) != 1 || resp.Users[0].ID != "user01" {
		t.Errorf("Users = %+v", resp.Users)
	}
	if len(resp.Bans) != 1 || resp.Bans[0].ID.Server != "srv01" {
		t.Errorf("Bans = %+v", resp.Bans)
	}
}

func TestFieldsMemberValues(t *testing.T) {
	if FieldsMemberNickname != "Nickname" {
		t.Errorf("FieldsMemberNickname = %q", FieldsMemberNickname)
	}
	if FieldsMemberAvatar != "Avatar" {
		t.Errorf("FieldsMemberAvatar = %q", FieldsMemberAvatar)
	}
	if FieldsMemberRoles != "Roles" {
		t.Errorf("FieldsMemberRoles = %q", FieldsMemberRoles)
	}
	if FieldsMemberTimeout != "Timeout" {
		t.Errorf("FieldsMemberTimeout = %q", FieldsMemberTimeout)
	}
}

func TestMessageWithMemberUnmarshal(t *testing.T) {
	data := `{
		"_id": "msg01",
		"channel": "ch01",
		"author": "user01",
		"content": "hello",
		"member": {
			"_id": {"server": "srv01", "user": "user01"},
			"joined_at": "2024-01-01T00:00:00Z",
			"nickname": "alice"
		}
	}`
	var msg Message
	if err := json.Unmarshal([]byte(data), &msg); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if msg.Member == nil {
		t.Fatal("Member is nil")
	}
	if msg.Member.ID.Server != "srv01" {
		t.Errorf("Member.ID.Server = %q", msg.Member.ID.Server)
	}
	if msg.Member.Nickname == nil || *msg.Member.Nickname != "alice" {
		t.Errorf("Member.Nickname = %v", msg.Member.Nickname)
	}
}

func TestBulkMessageResponseWithMembers(t *testing.T) {
	data := `{
		"messages": [{"_id": "msg01", "channel": "ch01", "author": "user01"}],
		"users": [{"_id": "user01", "username": "alice", "discriminator": "0001", "relationship": "None", "online": false}],
		"members": [{"_id": {"server": "srv01", "user": "user01"}, "joined_at": "2024-01-01T00:00:00Z"}]
	}`
	var resp BulkMessageResponse
	if err := json.Unmarshal([]byte(data), &resp); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if len(resp.Members) != 1 || resp.Members[0].ID.User != "user01" {
		t.Errorf("Members = %+v", resp.Members)
	}
}
