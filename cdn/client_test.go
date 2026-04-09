package cdn

import (
	"context"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUploadSessionToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/attachments" {
			t.Errorf("expected path /attachments, got %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Session-Token"); got != "sess-tok" {
			t.Errorf("expected session token sess-tok, got %q", got)
		}

		ct := r.Header.Get("Content-Type")
		mediaType, params, err := mime.ParseMediaType(ct)
		if err != nil {
			t.Fatalf("parse content-type: %v", err)
		}
		if mediaType != "multipart/form-data" {
			t.Fatalf("expected multipart/form-data, got %s", mediaType)
		}

		mr := multipart.NewReader(r.Body, params["boundary"])
		part, err := mr.NextPart()
		if err != nil {
			t.Fatalf("next part: %v", err)
		}
		if part.FormName() != "file" {
			t.Errorf("expected form name file, got %s", part.FormName())
		}
		if part.FileName() != "test.txt" {
			t.Errorf("expected filename test.txt, got %s", part.FileName())
		}
		data, _ := io.ReadAll(part)
		if string(data) != "hello world" {
			t.Errorf("expected body hello world, got %s", string(data))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"id": "file123"})
	}))
	defer srv.Close()

	c, err := New(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	c.SetSessionToken("sess-tok")

	id, err := c.Upload(context.Background(), TagAttachments, "test.txt", strings.NewReader("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	if id != "file123" {
		t.Errorf("expected file123, got %s", id)
	}
}

func TestUploadBotToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("X-Bot-Token"); got != "bot-tok" {
			t.Errorf("expected bot token bot-tok, got %q", got)
		}
		if got := r.Header.Get("X-Session-Token"); got != "" {
			t.Errorf("expected no session token, got %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"id": "file456"})
	}))
	defer srv.Close()

	c, err := New(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	c.SetBotToken("bot-tok")

	id, err := c.Upload(context.Background(), TagAvatars, "avatar.png", strings.NewReader("png-data"))
	if err != nil {
		t.Fatal(err)
	}
	if id != "file456" {
		t.Errorf("expected file456, got %s", id)
	}
}

func TestUploadErrorResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		w.Write([]byte("file too large"))
	}))
	defer srv.Close()

	c, err := New(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	c.SetSessionToken("tok")

	_, err = c.Upload(context.Background(), TagAttachments, "big.bin", strings.NewReader("data"))
	if err == nil {
		t.Fatal("expected error for 413 response")
	}
	if !strings.Contains(err.Error(), "413") {
		t.Errorf("expected error to mention 413, got: %v", err)
	}
}

func TestUploadNoAuth(t *testing.T) {
	c, err := New("http://localhost:9999")
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.Upload(context.Background(), TagAttachments, "f.txt", strings.NewReader("x"))
	if err == nil {
		t.Fatal("expected error when no auth token set")
	}
}

func TestURL(t *testing.T) {
	c, err := New("https://cdn.example.com")
	if err != nil {
		t.Fatal(err)
	}

	got := c.URL(TagAttachments, "abc123")
	want := "https://cdn.example.com/attachments/abc123"
	if got != want {
		t.Errorf("URL() = %q, want %q", got, want)
	}
}

func TestURLTrailingSlash(t *testing.T) {
	c, err := New("https://cdn.example.com/")
	if err != nil {
		t.Fatal(err)
	}

	got := c.URL(TagAvatars, "def456")
	want := "https://cdn.example.com/avatars/def456"
	if got != want {
		t.Errorf("URL() = %q, want %q", got, want)
	}
}

func TestOriginalURL(t *testing.T) {
	c, err := New("https://cdn.example.com")
	if err != nil {
		t.Fatal(err)
	}

	got := c.OriginalURL(TagEmojis, "ghi789")
	want := "https://cdn.example.com/emojis/ghi789/original"
	if got != want {
		t.Errorf("OriginalURL() = %q, want %q", got, want)
	}
}

func TestWithHTTPClient(t *testing.T) {
	custom := &http.Client{}
	c, err := New("https://cdn.example.com", WithHTTPClient(custom))
	if err != nil {
		t.Fatal(err)
	}
	if c.httpClient != custom {
		t.Error("expected custom http client to be set")
	}
}
