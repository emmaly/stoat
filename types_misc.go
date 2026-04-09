package stoat

// ChannelUnread represents the unread state for a channel.
type ChannelUnread struct {
	ID       ChannelCompositeKey `json:"_id"`
	LastID   *string             `json:"last_id,omitempty"`
	Mentions []string            `json:"mentions,omitempty"`
}

// OptionsFetchSettings is the request body for fetching settings.
type OptionsFetchSettings struct {
	Keys []string `json:"keys"`
}

// WebPushSubscription represents a web push subscription.
type WebPushSubscription struct {
	Endpoint string `json:"endpoint"`
	P256DH   string `json:"p256dh"`
	Auth     string `json:"auth"`
}

// DataReportContent is the request body for reporting content.
type DataReportContent struct {
	Content           ReportedContent `json:"content"`
	AdditionalContext *string         `json:"additional_context,omitempty"`
}

// ReportedContent represents the content being reported.
// This is a simplified struct rather than a full tagged union since it is send-only.
type ReportedContent struct {
	Type         string `json:"type"`
	ID           string `json:"id"`
	ReportReason string `json:"report_reason"`
	MessageID    string `json:"message_id,omitempty"`
}

// ContentReportReason is a string enum for content report reasons.
type ContentReportReason string

const (
	ContentReportReasonNoneSpecified      ContentReportReason = "NoneSpecified"
	ContentReportReasonIllegal            ContentReportReason = "Illegal"
	ContentReportReasonIllegalGoods       ContentReportReason = "IllegalGoods"
	ContentReportReasonIllegalExtortion   ContentReportReason = "IllegalExtortion"
	ContentReportReasonIllegalPornography ContentReportReason = "IllegalPornography"
	ContentReportReasonIllegalHacking     ContentReportReason = "IllegalHacking"
	ContentReportReasonExtremeViolence    ContentReportReason = "ExtremeViolence"
	ContentReportReasonPromotesHarm       ContentReportReason = "PromotesHarm"
	ContentReportReasonUnsolicitedSpam    ContentReportReason = "UnsolicitedSpam"
	ContentReportReasonRaid               ContentReportReason = "Raid"
	ContentReportReasonSpamAbuse          ContentReportReason = "SpamAbuse"
	ContentReportReasonScamsFraud         ContentReportReason = "ScamsFraud"
	ContentReportReasonMalware            ContentReportReason = "Malware"
	ContentReportReasonHarassment         ContentReportReason = "Harassment"
)

// UserReportReason is a string enum for user report reasons.
type UserReportReason string

const (
	UserReportReasonNoneSpecified       UserReportReason = "NoneSpecified"
	UserReportReasonUnsolicitedSpam     UserReportReason = "UnsolicitedSpam"
	UserReportReasonSpamAbuse           UserReportReason = "SpamAbuse"
	UserReportReasonInappropriateProfile UserReportReason = "InappropriateProfile"
	UserReportReasonImpersonation       UserReportReason = "Impersonation"
	UserReportReasonBanEvasion          UserReportReason = "BanEvasion"
	UserReportReasonUnderage            UserReportReason = "Underage"
)
