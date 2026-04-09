package stoat

// RevoltConfig is the root configuration returned by GET /.
type RevoltConfig struct {
	Revolt   string          `json:"revolt"`
	Features RevoltFeatures  `json:"features"`
	WS       string          `json:"ws"`
	App      string          `json:"app"`
	VAPID    string          `json:"vapid"`
	Build    BuildInformation `json:"build"`
}

// RevoltFeatures describes features enabled on this node.
type RevoltFeatures struct {
	Captcha    CaptchaFeature `json:"captcha"`
	Email      bool           `json:"email"`
	InviteOnly bool           `json:"invite_only"`
	Autumn     Feature        `json:"autumn"`
	January    Feature        `json:"january"`
	LiveKit    VoiceFeature   `json:"livekit"`
	Limits     LimitsConfig   `json:"limits"`
}

// Feature represents a generic service feature (e.g. Autumn, January).
type Feature struct {
	Enabled bool   `json:"enabled"`
	URL     string `json:"url"`
}

// CaptchaFeature represents hCaptcha configuration.
type CaptchaFeature struct {
	Enabled bool   `json:"enabled"`
	Key     string `json:"key"`
}

// VoiceFeature represents voice server configuration.
type VoiceFeature struct {
	Enabled bool        `json:"enabled"`
	Nodes   []VoiceNode `json:"nodes"`
}

// VoiceNode represents a single LiveKit voice node.
type VoiceNode struct {
	Name      string  `json:"name"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	PublicURL string  `json:"public_url"`
}

// BuildInformation contains build metadata for the server.
type BuildInformation struct {
	CommitSHA       string `json:"commit_sha"`
	CommitTimestamp  string `json:"commit_timestamp"`
	Semver          string `json:"semver"`
	OriginURL       string `json:"origin_url"`
	Timestamp       string `json:"timestamp"`
}

// LimitsConfig contains all limit configurations.
type LimitsConfig struct {
	Global  GlobalLimits `json:"global"`
	NewUser UserLimits   `json:"new_user"`
	Default UserLimits   `json:"default"`
}

// GlobalLimits contains global instance limits.
type GlobalLimits struct {
	GroupSize              int      `json:"group_size"`
	MessageEmbeds          int      `json:"message_embeds"`
	MessageReplies         int      `json:"message_replies"`
	MessageReactions       int      `json:"message_reactions"`
	ServerEmoji            int      `json:"server_emoji"`
	ServerRoles            int      `json:"server_roles"`
	ServerChannels         int      `json:"server_channels"`
	BodyLimitSize          int      `json:"body_limit_size"`
	RestrictServerCreation []string `json:"restrict_server_creation"`
}

// UserLimits contains per-user limits.
type UserLimits struct {
	OutgoingFriendRequests int            `json:"outgoing_friend_requests"`
	Bots                   int            `json:"bots"`
	MessageLength          int            `json:"message_length"`
	MessageAttachments     int            `json:"message_attachments"`
	Servers                int            `json:"servers"`
	VoiceQuality           int            `json:"voice_quality"`
	Video                  bool           `json:"video"`
	VideoResolution        []int          `json:"video_resolution"`
	VideoAspectRatio       []float64      `json:"video_aspect_ratio"`
	FileUploadSizeLimits   map[string]int `json:"file_upload_size_limits"`
}
