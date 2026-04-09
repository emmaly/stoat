package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

const defaultHeartbeatInterval = 20 * time.Second

// Conn is a WebSocket connection to the Stoat chat API.
type Conn struct {
	ws     *websocket.Conn
	token  string
	mu     sync.Mutex // protects writes
	closed chan struct{}
}

// Dial connects to the WebSocket server, authenticates, and starts the heartbeat.
func Dial(ctx context.Context, url string, token string) (*Conn, error) {
	return DialWithHeartbeat(ctx, url, token, defaultHeartbeatInterval)
}

// DialWithHeartbeat connects with a custom heartbeat interval. Useful for testing.
func DialWithHeartbeat(ctx context.Context, url string, token string, heartbeatInterval time.Duration) (*Conn, error) {
	ws, _, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		return nil, fmt.Errorf("ws dial: %w", err)
	}

	c := &Conn{
		ws:     ws,
		token:  token,
		closed: make(chan struct{}),
	}

	// Send Authenticate event
	if err := c.Send(ctx, AuthenticateEvent{Token: token}); err != nil {
		ws.Close(websocket.StatusInternalError, "auth send failed")
		return nil, fmt.Errorf("ws authenticate send: %w", err)
	}

	// Wait for Authenticated or Error
	ev, err := c.ReadEvent(ctx)
	if err != nil {
		ws.Close(websocket.StatusInternalError, "auth read failed")
		return nil, fmt.Errorf("ws authenticate read: %w", err)
	}

	switch e := ev.(type) {
	case *AuthenticatedEvent:
		// Success
	case *ErrorEvent:
		ws.Close(websocket.StatusNormalClosure, "")
		return nil, fmt.Errorf("ws authentication error: %s", e.Error)
	default:
		ws.Close(websocket.StatusInternalError, "unexpected event")
		return nil, fmt.Errorf("ws expected Authenticated, got %T", ev)
	}

	c.startHeartbeat(ctx, heartbeatInterval)

	return c, nil
}

// Send marshals an event to JSON and writes it to the WebSocket. It is safe
// for concurrent use.
func (c *Conn) Send(ctx context.Context, event Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("ws marshal: %w", err)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	return c.ws.Write(ctx, websocket.MessageText, data)
}

// ReadEvent reads a single event from the WebSocket.
func (c *Conn) ReadEvent(ctx context.Context) (Event, error) {
	_, data, err := c.ws.Read(ctx)
	if err != nil {
		return nil, err
	}

	var raw RawEvent
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("ws unmarshal: %w", err)
	}

	return raw.Value, nil
}

// Close closes the WebSocket connection.
func (c *Conn) Close() error {
	select {
	case <-c.closed:
		return nil
	default:
		close(c.closed)
	}
	return c.ws.Close(websocket.StatusNormalClosure, "")
}

// startHeartbeat starts a goroutine that sends PingEvent at the given interval.
func (c *Conn) startHeartbeat(ctx context.Context, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-c.closed:
				return
			case <-ctx.Done():
				return
			case t := <-ticker.C:
				c.Send(ctx, PingEvent{Data: t.UnixMilli()})
			}
		}
	}()
}
