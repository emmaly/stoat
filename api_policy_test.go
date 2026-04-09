package stoat

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAcknowledgePolicy(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/policy/acknowledge" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	if err := c.AcknowledgePolicy(context.Background()); err != nil {
		t.Fatalf("AcknowledgePolicy: %v", err)
	}
}
