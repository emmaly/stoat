# Server Permissions

Role CRUD and server-level permission management.

## Create Role

```
POST /servers/{target}/roles
Auth: Session Token
Path: target (string, required) — server ID
Body: DataCreateRole
Response 200: NewRoleResponse
```

### DataCreateRole

```json
{
  "name": "string",  // required
  "rank": 0          // optional — integer rank position
}
```

### NewRoleResponse

```json
{
  "id": "string",
  "role": { /* Role */ }
}
```

## Fetch Role

```
GET /servers/{target}/roles/{role_id}
Auth: Session Token
Path: target (string, required), role_id (string, required)
Response 200: Role
```

## Edit Role

```
PATCH /servers/{target}/roles/{role_id}
Auth: Session Token
Path: target (string, required), role_id (string, required)
Body: DataEditRole
Response 200: Role
```

### DataEditRole

```json
{
  "name": "string",    // optional
  "colour": "string",  // optional — CSS color string
  "hoist": true,        // optional — display separately in member list
  "rank": 0,            // optional
  "remove": ["Colour"]  // optional — fields to clear
}
```

Remove options (FieldsRole): `"Colour"`

## Delete Role

```
DELETE /servers/{target}/roles/{role_id}
Auth: Session Token
Path: target (string, required), role_id (string, required)
Response 204: Success
```

## Edit Role Ranks

Reorder role positions.

```
PATCH /servers/{target}/roles/ranks
Auth: Session Token
Path: target (string, required)
Body: DataEditRoleRanks
Response 200: Server
```

### DataEditRoleRanks

```json
{
  "roles": {
    "role_id": 0,   // role_id → new rank position
    "role_id": 1
  }
}
```

## Set Default Server Permissions

Set the default permission value for all members.

```
PUT /servers/{target}/permissions/default
Auth: Session Token
Path: target (string, required)
Body: DataPermissionsValue
Response 200: Server
```

### DataPermissionsValue

```json
{
  "permissions": 0  // required — int64 permission bitfield
}
```

## Set Role Permission on Server

```
PUT /servers/{target}/permissions/{role_id}
Auth: Session Token
Path: target (string, required), role_id (string, required)
Body: DataSetServerRolePermission
Response 200: Server
```

### DataSetServerRolePermission

```json
{
  "permissions": {
    "allow": 0,  // int64 — permission bits to allow
    "deny": 0    // int64 — permission bits to deny
  }
}
```

## Role Type

```json
{
  "name": "string",
  "permissions": {
    "a": 0,  // allow bits (int64)
    "d": 0   // deny bits (int64)
  },
  "colour": "string",  // nullable — CSS color
  "hoist": false,       // display separately in member list
  "rank": 0             // int64 — lower = higher priority
}
```

See [Permissions](../concepts/permissions.md) for the full permission flag reference.
