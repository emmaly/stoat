package stoat

import (
	"net/http"
	"strconv"
	"time"
)

// RateLimit contains rate limit information parsed from API response headers.
type RateLimit struct {
	Limit      int    // Maximum calls allowed for this bucket.
	Bucket     string // Unique identifier for the bucket.
	Remaining  int    // Calls remaining in current window.
	ResetAfter int    // Milliseconds until bucket replenishes.
}

// ParseRateLimit extracts rate limit information from HTTP response headers.
// Returns nil if the required headers are not present.
func ParseRateLimit(h http.Header) *RateLimit {
	bucket := h.Get("X-RateLimit-Bucket")
	if bucket == "" {
		return nil
	}

	limit, err := strconv.Atoi(h.Get("X-RateLimit-Limit"))
	if err != nil {
		return nil
	}

	remaining, err := strconv.Atoi(h.Get("X-RateLimit-Remaining"))
	if err != nil {
		return nil
	}

	resetAfter, err := strconv.Atoi(h.Get("X-RateLimit-Reset-After"))
	if err != nil {
		return nil
	}

	return &RateLimit{
		Limit:      limit,
		Bucket:     bucket,
		Remaining:  remaining,
		ResetAfter: resetAfter,
	}
}

// ShouldWait returns the duration to wait before making the next request.
// Returns 0 if there are remaining calls in the current window.
func (rl *RateLimit) ShouldWait() time.Duration {
	if rl.Remaining > 0 {
		return 0
	}
	return time.Duration(rl.ResetAfter) * time.Millisecond
}
