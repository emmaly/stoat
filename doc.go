// Package stoat provides a Go client for the Stoat chat platform API.
//
// It covers the full REST API (121 endpoints), all data types, and provides
// sub-packages for WebSocket real-time events ([github.com/emmaly/stoat/ws])
// and CDN file uploads ([github.com/emmaly/stoat/cdn]).
//
// # Quick Start
//
//	c, err := stoat.New("https://stoat.chat/api")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	c.SetBotToken("your-bot-token")
//
//	cfg, err := c.QueryNode(context.Background())
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println("Connected to", cfg.Revolt)
//
// # Authentication
//
// The client supports three authentication modes:
//
//   - Session token: [Client.SetSessionToken] (obtained via [Client.Login])
//   - Bot token: [Client.SetBotToken] (created in the Stoat client UI)
//   - MFA ticket: [Client.SetMFATicket] (for sensitive operations)
//
// # Tagged Unions
//
// Some API types are discriminated unions (e.g., [Channel], [Embed],
// [SystemMessage]). These are represented as Go interfaces with concrete
// types for each variant. Use type assertions or type switches to access
// variant-specific fields. JSON unmarshaling is handled by Raw* wrapper
// types (e.g., [RawChannel]).
package stoat
