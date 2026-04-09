package stoat

// NOTE: Role is defined in types_server.go since Server depends on it.

// DataCreateRole is the request body for creating a role.
type DataCreateRole struct {
	Name string `json:"name"`
	Rank *int64 `json:"rank,omitempty"`
}

// DataEditRole is the request body for editing a role.
type DataEditRole struct {
	Name   *string      `json:"name,omitempty"`
	Colour *string      `json:"colour,omitempty"`
	Hoist  *bool        `json:"hoist,omitempty"`
	Rank   *int64       `json:"rank,omitempty"`
	Remove []FieldsRole `json:"remove,omitempty"`
}

// DataEditRoleRanks is the request body for reordering role positions.
type DataEditRoleRanks struct {
	Roles map[string]int64 `json:"roles"`
}

// NewRoleResponse is the response from creating a role.
type NewRoleResponse struct {
	ID   string `json:"id"`
	Role Role   `json:"role"`
}

// DataPermissionsValue is the request body for setting default server permissions.
type DataPermissionsValue struct {
	Permissions uint64 `json:"permissions"`
}

// DataSetServerRolePermission is the request body for setting role permissions on a server.
type DataSetServerRolePermission struct {
	Permissions Override `json:"permissions"`
}

// FieldsRole is a string enum of optional fields on a role object that can be removed.
type FieldsRole string

const (
	FieldsRoleColour FieldsRole = "Colour"
)
