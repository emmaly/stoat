package stoat

import (
	"net/http"
	"testing"
	"time"
)

func TestParseRateLimit(t *testing.T) {
	h := make(http.Header)
	h.Set("X-RateLimit-Limit", "10")
	h.Set("X-RateLimit-Bucket", "post_messages")
	h.Set("X-RateLimit-Remaining", "7")
	h.Set("X-RateLimit-Reset-After", "5000")

	rl := ParseRateLimit(h)
	if rl == nil {
		t.Fatal("expected non-nil RateLimit")
	}
	if rl.Limit != 10 {
		t.Errorf("Limit = %d, want 10", rl.Limit)
	}
	if rl.Bucket != "post_messages" {
		t.Errorf("Bucket = %q, want post_messages", rl.Bucket)
	}
	if rl.Remaining != 7 {
		t.Errorf("Remaining = %d, want 7", rl.Remaining)
	}
	if rl.ResetAfter != 5000 {
		t.Errorf("ResetAfter = %d, want 5000", rl.ResetAfter)
	}
}

func TestParseRateLimitMissingHeaders(t *testing.T) {
	h := make(http.Header)
	rl := ParseRateLimit(h)
	if rl != nil {
		t.Error("expected nil RateLimit for missing headers")
	}
}

func TestShouldWaitExhausted(t *testing.T) {
	rl := &RateLimit{
		Limit:      10,
		Bucket:     "test",
		Remaining:  0,
		ResetAfter: 3000,
	}
	d := rl.ShouldWait()
	if d != 3*time.Second {
		t.Errorf("ShouldWait = %v, want 3s", d)
	}
}

func TestShouldWaitNotExhausted(t *testing.T) {
	rl := &RateLimit{
		Limit:      10,
		Bucket:     "test",
		Remaining:  5,
		ResetAfter: 3000,
	}
	d := rl.ShouldWait()
	if d != 0 {
		t.Errorf("ShouldWait = %v, want 0", d)
	}
}
