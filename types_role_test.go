package stoat

import (
	"encoding/json"
	"testing"
)

func TestDataCreateRoleMarshal(t *testing.T) {
	rank := int64(5)
	d := DataCreateRole{
		Name: "Moderator",
		Rank: &rank,
	}
	b, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var m map[string]any
	json.Unmarshal(b, &m)
	if m["name"] != "Moderator" {
		t.Errorf("name = %v", m["name"])
	}
	if m["rank"] != float64(5) {
		t.Errorf("rank = %v", m["rank"])
	}
}

func TestDataCreateRoleOmitRank(t *testing.T) {
	d := DataCreateRole{Name: "Member"}
	b, _ := json.Marshal(d)
	var m map[string]any
	json.Unmarshal(b, &m)
	if _, ok := m["rank"]; ok {
		t.Error("rank should be omitted")
	}
}

func TestDataEditRoleMarshal(t *testing.T) {
	name := "Admin"
	colour := "#00ff00"
	hoist := true
	d := DataEditRole{
		Name:   &name,
		Colour: &colour,
		Hoist:  &hoist,
		Remove: []FieldsRole{FieldsRoleColour},
	}
	b, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var m map[string]any
	json.Unmarshal(b, &m)
	if m["name"] != "Admin" {
		t.Errorf("name = %v", m["name"])
	}
	if m["colour"] != "#00ff00" {
		t.Errorf("colour = %v", m["colour"])
	}
	if m["hoist"] != true {
		t.Errorf("hoist = %v", m["hoist"])
	}
	r := m["remove"].([]any)
	if r[0] != "Colour" {
		t.Errorf("remove[0] = %v", r[0])
	}
}

func TestDataEditRoleRanksMarshal(t *testing.T) {
	d := DataEditRoleRanks{
		Roles: map[string]int64{"role01": 0, "role02": 1},
	}
	b, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var m map[string]any
	json.Unmarshal(b, &m)
	roles := m["roles"].(map[string]any)
	if roles["role01"] != float64(0) {
		t.Errorf("role01 = %v", roles["role01"])
	}
	if roles["role02"] != float64(1) {
		t.Errorf("role02 = %v", roles["role02"])
	}
}

func TestNewRoleResponseUnmarshal(t *testing.T) {
	data := `{
		"id": "role01",
		"role": {
			"name": "Admin",
			"permissions": {"a": 255, "d": 0},
			"hoist": true,
			"rank": 0
		}
	}`
	var resp NewRoleResponse
	if err := json.Unmarshal([]byte(data), &resp); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if resp.ID != "role01" {
		t.Errorf("ID = %q", resp.ID)
	}
	if resp.Role.Name != "Admin" {
		t.Errorf("Role.Name = %q", resp.Role.Name)
	}
}

func TestDataPermissionsValueMarshal(t *testing.T) {
	d := DataPermissionsValue{Permissions: 1048576}
	b, _ := json.Marshal(d)
	var m map[string]any
	json.Unmarshal(b, &m)
	if m["permissions"] != float64(1048576) {
		t.Errorf("permissions = %v", m["permissions"])
	}
}

func TestDataSetServerRolePermissionMarshal(t *testing.T) {
	d := DataSetServerRolePermission{
		Permissions: Override{Allow: 255, Deny: 10},
	}
	b, _ := json.Marshal(d)
	var m map[string]any
	json.Unmarshal(b, &m)
	perms := m["permissions"].(map[string]any)
	if perms["allow"] != float64(255) {
		t.Errorf("allow = %v", perms["allow"])
	}
	if perms["deny"] != float64(10) {
		t.Errorf("deny = %v", perms["deny"])
	}
}

func TestFieldsRoleValue(t *testing.T) {
	if FieldsRoleColour != "Colour" {
		t.Errorf("FieldsRoleColour = %q", FieldsRoleColour)
	}
}
