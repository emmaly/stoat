package ws

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"nhooyr.io/websocket"
)

type testHandler struct {
	DefaultEventHandler
	mu       sync.Mutex
	messages []MessageEvent
	errors   []ErrorEvent
	ready    []ReadyEvent
	pongs    []PongEvent
	deletes  []MessageDeleteEvent
}

func (h *testHandler) OnMessage(ev MessageEvent) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.messages = append(h.messages, ev)
}

func (h *testHandler) OnError(ev ErrorEvent) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.errors = append(h.errors, ev)
}

func (h *testHandler) OnReady(ev ReadyEvent) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.ready = append(h.ready, ev)
}

func (h *testHandler) OnPong(ev PongEvent) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.pongs = append(h.pongs, ev)
}

func (h *testHandler) OnMessageDelete(ev MessageDeleteEvent) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.deletes = append(h.deletes, ev)
}

func TestListenDispatchesEvents(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close(websocket.StatusNormalClosure, "")

		ctx := r.Context()

		// Auth handshake
		conn.Read(ctx)
		conn.Write(ctx, websocket.MessageText, []byte(`{"type":"Authenticated"}`))

		// Send a Message event
		conn.Write(ctx, websocket.MessageText, []byte(`{"type":"Message","_id":"m1","channel":"c1","author":"u1","content":"hello"}`))

		// Send an Error event
		conn.Write(ctx, websocket.MessageText, []byte(`{"type":"Error","error":"LabelMe"}`))

		// Drain pings until close
		for {
			_, _, err := conn.Read(ctx)
			if err != nil {
				return
			}
		}
	}))
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c, err := DialWithHeartbeat(ctx, wsURL(srv), "tok", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	handler := &testHandler{}

	// Listen in a goroutine, cancel after receiving events
	listenCtx, listenCancel := context.WithCancel(ctx)
	go func() {
		// Give it time to process events
		time.Sleep(200 * time.Millisecond)
		listenCancel()
	}()

	_ = c.Listen(listenCtx, handler)

	handler.mu.Lock()
	defer handler.mu.Unlock()

	if len(handler.messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(handler.messages))
	}
	if handler.messages[0].ID != "m1" {
		t.Errorf("expected message id m1, got %s", handler.messages[0].ID)
	}
	if len(handler.errors) != 1 {
		t.Fatalf("expected 1 error, got %d", len(handler.errors))
	}
	if handler.errors[0].Error != "LabelMe" {
		t.Errorf("expected error LabelMe, got %s", handler.errors[0].Error)
	}
}

func TestListenBulkEvent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close(websocket.StatusNormalClosure, "")

		ctx := r.Context()

		// Auth
		conn.Read(ctx)
		conn.Write(ctx, websocket.MessageText, []byte(`{"type":"Authenticated"}`))

		// Send Bulk with two sub-events
		bulk := map[string]any{
			"type": "Bulk",
			"v": []any{
				map[string]any{"type": "MessageDelete", "id": "m1", "channel": "c1"},
				map[string]any{"type": "MessageDelete", "id": "m2", "channel": "c2"},
			},
		}
		data, _ := json.Marshal(bulk)
		conn.Write(ctx, websocket.MessageText, data)

		for {
			_, _, err := conn.Read(ctx)
			if err != nil {
				return
			}
		}
	}))
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c, err := DialWithHeartbeat(ctx, wsURL(srv), "tok", 10*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	handler := &testHandler{}
	listenCtx, listenCancel := context.WithCancel(ctx)
	go func() {
		time.Sleep(200 * time.Millisecond)
		listenCancel()
	}()

	_ = c.Listen(listenCtx, handler)

	handler.mu.Lock()
	defer handler.mu.Unlock()

	if len(handler.deletes) != 2 {
		t.Fatalf("expected 2 deletes from bulk, got %d", len(handler.deletes))
	}
	if handler.deletes[0].ID != "m1" || handler.deletes[1].ID != "m2" {
		t.Errorf("unexpected delete ids: %s, %s", handler.deletes[0].ID, handler.deletes[1].ID)
	}
}

func TestDefaultEventHandlerCompiles(t *testing.T) {
	// Verify DefaultEventHandler satisfies EventHandler
	var _ EventHandler = &DefaultEventHandler{}
}
