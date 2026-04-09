package ws

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"nhooyr.io/websocket"
)

// testServer creates a test WebSocket server that handles authentication and echoes pings.
func testServer(t *testing.T, handler func(ctx context.Context, conn *websocket.Conn)) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			t.Logf("accept error: %v", err)
			return
		}
		defer conn.Close(websocket.StatusNormalClosure, "")
		handler(r.Context(), conn)
	}))
}

func wsURL(s *httptest.Server) string {
	return "ws" + strings.TrimPrefix(s.URL, "http")
}

func TestDialAndClose(t *testing.T) {
	srv := testServer(t, func(ctx context.Context, conn *websocket.Conn) {
		// Read the Authenticate event
		_, data, err := conn.Read(ctx)
		if err != nil {
			t.Logf("read error: %v", err)
			return
		}
		var m map[string]any
		json.Unmarshal(data, &m)
		if m["type"] != "Authenticate" {
			t.Errorf("expected Authenticate, got %v", m["type"])
			return
		}

		// Send Authenticated
		conn.Write(ctx, websocket.MessageText, []byte(`{"type":"Authenticated"}`))

		// Keep connection open until client closes
		for {
			_, _, err := conn.Read(ctx)
			if err != nil {
				return
			}
		}
	})
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c, err := Dial(ctx, wsURL(srv), "test-token")
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestDialAuthError(t *testing.T) {
	srv := testServer(t, func(ctx context.Context, conn *websocket.Conn) {
		// Read Authenticate
		conn.Read(ctx)
		// Send Error instead of Authenticated
		conn.Write(ctx, websocket.MessageText, []byte(`{"type":"Error","error":"InvalidSession"}`))
	})
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := Dial(ctx, wsURL(srv), "bad-token")
	if err == nil {
		t.Fatal("expected error for invalid session")
	}
	if !strings.Contains(err.Error(), "InvalidSession") {
		t.Errorf("expected InvalidSession in error, got: %v", err)
	}
}

func TestSendAndReadEvent(t *testing.T) {
	srv := testServer(t, func(ctx context.Context, conn *websocket.Conn) {
		// Auth handshake
		conn.Read(ctx)
		conn.Write(ctx, websocket.MessageText, []byte(`{"type":"Authenticated"}`))

		// Read a Ping from client, respond with Pong
		_, data, err := conn.Read(ctx)
		if err != nil {
			return
		}
		var m map[string]any
		json.Unmarshal(data, &m)
		if m["type"] == "Ping" {
			pongData := int64(m["data"].(float64))
			resp, _ := json.Marshal(map[string]any{"type": "Pong", "data": pongData})
			conn.Write(ctx, websocket.MessageText, resp)
		}

		// Keep alive
		for {
			_, _, err := conn.Read(ctx)
			if err != nil {
				return
			}
		}
	})
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c, err := Dial(ctx, wsURL(srv), "test-token")
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// Send Ping
	if err := c.Send(ctx, PingEvent{Data: 42}); err != nil {
		t.Fatal(err)
	}

	// Read Pong
	ev, err := c.ReadEvent(ctx)
	if err != nil {
		t.Fatal(err)
	}
	pong, ok := ev.(*PongEvent)
	if !ok {
		t.Fatalf("expected *PongEvent, got %T", ev)
	}
	if pong.Data != 42 {
		t.Errorf("expected data 42, got %d", pong.Data)
	}
}

func TestSendGoroutineSafety(t *testing.T) {
	srv := testServer(t, func(ctx context.Context, conn *websocket.Conn) {
		conn.Read(ctx)
		conn.Write(ctx, websocket.MessageText, []byte(`{"type":"Authenticated"}`))
		for {
			_, _, err := conn.Read(ctx)
			if err != nil {
				return
			}
		}
	})
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c, err := Dial(ctx, wsURL(srv), "test-token")
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// Send from multiple goroutines concurrently
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			c.Send(ctx, PingEvent{Data: int64(n)})
		}(i)
	}
	wg.Wait()
}

func TestHeartbeat(t *testing.T) {
	var mu sync.Mutex
	pingsReceived := 0

	srv := testServer(t, func(ctx context.Context, conn *websocket.Conn) {
		conn.Read(ctx)
		conn.Write(ctx, websocket.MessageText, []byte(`{"type":"Authenticated"}`))
		for {
			_, data, err := conn.Read(ctx)
			if err != nil {
				return
			}
			var m map[string]any
			json.Unmarshal(data, &m)
			if m["type"] == "Ping" {
				mu.Lock()
				pingsReceived++
				mu.Unlock()
				resp, _ := json.Marshal(map[string]any{"type": "Pong", "data": m["data"]})
				conn.Write(ctx, websocket.MessageText, resp)
			}
		}
	})
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c, err := DialWithHeartbeat(ctx, wsURL(srv), "test-token", 100*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	// Wait for a few heartbeats
	time.Sleep(350 * time.Millisecond)

	mu.Lock()
	n := pingsReceived
	mu.Unlock()

	if n < 2 {
		t.Errorf("expected at least 2 pings, got %d", n)
	}
}
