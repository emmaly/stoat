package stoat

import "testing"

func TestPermissionConstants(t *testing.T) {
	tests := []struct {
		name string
		perm Permission
		want uint64
	}{
		// Server Management (bits 0-13)
		{"ManageChannel", ManageChannel, 1 << 0},
		{"ManageServer", ManageServer, 1 << 1},
		{"ManagePermissions", ManagePermissions, 1 << 2},
		{"ManageRole", ManageRole, 1 << 3},
		{"ManageCustomisation", ManageCustomisation, 1 << 4},
		{"KickMembers", KickMembers, 1 << 6},
		{"BanMembers", BanMembers, 1 << 7},
		{"TimeoutMembers", TimeoutMembers, 1 << 8},
		{"AssignRoles", AssignRoles, 1 << 9},
		{"ChangeNickname", ChangeNickname, 1 << 10},
		{"ManageNicknames", ManageNicknames, 1 << 11},
		{"ChangeAvatar", ChangeAvatar, 1 << 12},
		{"RemoveAvatars", RemoveAvatars, 1 << 13},

		// Channel (bits 20-29)
		{"ViewChannel", ViewChannel, 1 << 20},
		{"ReadMessageHistory", ReadMessageHistory, 1 << 21},
		{"SendMessage", SendMessage, 1 << 22},
		{"ManageMessages", ManageMessages, 1 << 23},
		{"ManageWebhooks", ManageWebhooks, 1 << 24},
		{"InviteOthers", InviteOthers, 1 << 25},
		{"SendEmbeds", SendEmbeds, 1 << 26},
		{"UploadFiles", UploadFiles, 1 << 27},
		{"UseMasquerade", UseMasquerade, 1 << 28},
		{"React", React, 1 << 29},

		// Voice (bits 30-35)
		{"Connect", Connect, 1 << 30},
		{"Speak", Speak, 1 << 31},
		{"ShareVideo", ShareVideo, 1 << 32},
		{"MuteMembers", MuteMembers, 1 << 33},
		{"DeafenMembers", DeafenMembers, 1 << 34},
		{"MoveMembers", MoveMembers, 1 << 35},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if uint64(tt.perm) != tt.want {
				t.Errorf("%s = %d, want %d", tt.name, tt.perm, tt.want)
			}
		})
	}
}

func TestPermissionHas(t *testing.T) {
	p := ManageChannel | ManageServer | ViewChannel
	if !p.Has(ManageChannel) {
		t.Error("expected Has(ManageChannel) to be true")
	}
	if !p.Has(ManageServer) {
		t.Error("expected Has(ManageServer) to be true")
	}
	if p.Has(BanMembers) {
		t.Error("expected Has(BanMembers) to be false")
	}
	// Has with multiple flags: all must be present
	if !p.Has(ManageChannel | ManageServer) {
		t.Error("expected Has(ManageChannel|ManageServer) to be true")
	}
	if p.Has(ManageChannel | BanMembers) {
		t.Error("expected Has(ManageChannel|BanMembers) to be false")
	}
}

func TestPermissionAdd(t *testing.T) {
	p := ManageChannel
	p = p.Add(ViewChannel)
	if !p.Has(ManageChannel) {
		t.Error("lost ManageChannel after Add")
	}
	if !p.Has(ViewChannel) {
		t.Error("expected ViewChannel after Add")
	}
}

func TestPermissionRemove(t *testing.T) {
	p := ManageChannel | ViewChannel | SendMessage
	p = p.Remove(ViewChannel)
	if p.Has(ViewChannel) {
		t.Error("expected ViewChannel to be removed")
	}
	if !p.Has(ManageChannel) {
		t.Error("lost ManageChannel after Remove")
	}
	if !p.Has(SendMessage) {
		t.Error("lost SendMessage after Remove")
	}
}
