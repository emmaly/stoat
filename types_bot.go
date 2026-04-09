package stoat

// Bot represents a Revolt bot.
type Bot struct {
	ID                string  `json:"_id"`
	Owner             string  `json:"owner"`
	Token             string  `json:"token"`
	Public            bool    `json:"public"`
	Analytics         bool    `json:"analytics,omitempty"`
	Discoverable      bool    `json:"discoverable,omitempty"`
	InteractionsURL   *string `json:"interactions_url,omitempty"`
	TermsOfServiceURL *string `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyURL  *string `json:"privacy_policy_url,omitempty"`
	Flags             uint32  `json:"flags,omitempty"`
}

// PublicBot represents public information about a bot.
type PublicBot struct {
	ID          string  `json:"_id"`
	Username    string  `json:"username"`
	Avatar      *File   `json:"avatar,omitempty"`
	Description *string `json:"description,omitempty"`
}

// DataCreateBot is the request body for creating a bot.
type DataCreateBot struct {
	Name string `json:"name"`
}

// DataEditBot is the request body for editing a bot.
type DataEditBot struct {
	Name            *string    `json:"name,omitempty"`
	Public          *bool      `json:"public,omitempty"`
	Analytics       *bool      `json:"analytics,omitempty"`
	InteractionsURL *string    `json:"interactions_url,omitempty"`
	Remove          []FieldsBot `json:"remove,omitempty"`
}

// FetchBotResponse is the response from fetching a bot.
type FetchBotResponse struct {
	Bot  Bot  `json:"bot"`
	User User `json:"user"`
}

// BotWithUserResponse is the response from creating or editing a bot.
// The bot fields are flattened at the top level alongside the user field.
type BotWithUserResponse struct {
	Bot
	User User `json:"user"`
}

// OwnedBotsResponse is the response from fetching all owned bots.
type OwnedBotsResponse struct {
	Bots  []Bot  `json:"bots"`
	Users []User `json:"users"`
}

// InviteBotDestination specifies where to invite a bot (server or group).
type InviteBotDestination struct {
	Server *string `json:"server,omitempty"`
	Group  *string `json:"group,omitempty"`
}

// FieldsBot is a string enum of optional fields on a bot object that can be removed.
type FieldsBot string

const (
	FieldsBotToken           FieldsBot = "Token"
	FieldsBotInteractionsURL FieldsBot = "InteractionsURL"
)
