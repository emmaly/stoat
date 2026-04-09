// Package ws provides a WebSocket client for Stoat real-time events.
//
// It handles connection management, authentication handshake, heartbeat,
// and event dispatch for all 40 Stoat WebSocket event types.
//
// # Quick Start
//
//	conn, err := ws.Dial(ctx, "wss://stoat.chat/events", "your-token")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer conn.Close()
//
//	handler := &myHandler{}
//	if err := conn.Listen(ctx, handler); err != nil {
//	    log.Fatal(err)
//	}
//
// # Event Handling
//
// Implement the [EventHandler] interface (or embed [DefaultEventHandler] and
// override specific methods) to handle events. Pass the handler to
// [Conn.Listen] for automatic dispatch.
package ws
