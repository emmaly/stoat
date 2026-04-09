package stoat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReportContent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %q", r.Method)
		}
		if r.URL.Path != "/safety/report" {
			t.Errorf("path = %q", r.URL.Path)
		}
		var body DataReportContent
		json.NewDecoder(r.Body).Decode(&body)
		if body.Content.Type != "Message" {
			t.Errorf("content.type = %q", body.Content.Type)
		}
		if body.Content.ID != "msg01" {
			t.Errorf("content.id = %q", body.Content.ID)
		}
		if body.Content.ReportReason != string(ContentReportReasonHarassment) {
			t.Errorf("content.report_reason = %q", body.Content.ReportReason)
		}
		if body.AdditionalContext == nil || *body.AdditionalContext != "details" {
			t.Errorf("additional_context = %v", body.AdditionalContext)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	c, _ := New(srv.URL)
	c.SetSessionToken("tok")
	ctx := "details"
	err := c.ReportContent(context.Background(), DataReportContent{
		Content: ReportedContent{
			Type:         "Message",
			ID:           "msg01",
			ReportReason: string(ContentReportReasonHarassment),
		},
		AdditionalContext: &ctx,
	})
	if err != nil {
		t.Fatalf("ReportContent: %v", err)
	}
}
