package stoat

// Permission represents a bitfield of permission flags.
type Permission uint64

// Has reports whether p contains all the flags in other.
func (p Permission) Has(other Permission) bool {
	return p&other == other
}

// Add returns p with the flags in other set.
func (p Permission) Add(other Permission) Permission {
	return p | other
}

// Remove returns p with the flags in other cleared.
func (p Permission) Remove(other Permission) Permission {
	return p &^ other
}

// Server Management Permissions (bits 0–13)
const (
	ManageChannel       Permission = 1 << 0
	ManageServer        Permission = 1 << 1
	ManagePermissions   Permission = 1 << 2
	ManageRole          Permission = 1 << 3
	ManageCustomisation Permission = 1 << 4
	// bit 5 is reserved
	KickMembers    Permission = 1 << 6
	BanMembers     Permission = 1 << 7
	TimeoutMembers Permission = 1 << 8
	AssignRoles    Permission = 1 << 9
	ChangeNickname Permission = 1 << 10
	ManageNicknames Permission = 1 << 11
	ChangeAvatar   Permission = 1 << 12
	RemoveAvatars  Permission = 1 << 13
)

// Channel Permissions (bits 20–29)
const (
	ViewChannel        Permission = 1 << 20
	ReadMessageHistory Permission = 1 << 21
	SendMessage        Permission = 1 << 22
	ManageMessages     Permission = 1 << 23
	ManageWebhooks     Permission = 1 << 24
	InviteOthers       Permission = 1 << 25
	SendEmbeds         Permission = 1 << 26
	UploadFiles        Permission = 1 << 27
	UseMasquerade      Permission = 1 << 28
	React              Permission = 1 << 29
)

// Voice Permissions (bits 30–35)
const (
	Connect       Permission = 1 << 30
	Speak         Permission = 1 << 31
	ShareVideo    Permission = 1 << 32
	MuteMembers   Permission = 1 << 33
	DeafenMembers Permission = 1 << 34
	MoveMembers   Permission = 1 << 35
)
