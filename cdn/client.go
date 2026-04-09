// Package cdn provides a client for the Stoat CDN (Autumn) file server.
package cdn

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"
)

// Client is a CDN client for uploading and referencing files.
type Client struct {
	baseURL    string
	httpClient *http.Client

	mu           sync.RWMutex
	sessionToken string
	botToken     string
}

// Option configures the Client.
type Option func(*Client)

// WithHTTPClient sets a custom http.Client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.httpClient = hc
	}
}

// New creates a new CDN client for the given base URL.
func New(baseURL string, opts ...Option) (*Client, error) {
	baseURL = strings.TrimRight(baseURL, "/")
	c := &Client{
		baseURL:    baseURL,
		httpClient: http.DefaultClient,
	}
	for _, o := range opts {
		o(c)
	}
	return c, nil
}

// SetSessionToken sets the session token for user authentication.
func (c *Client) SetSessionToken(token string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sessionToken = token
}

// SetBotToken sets the bot token for bot authentication.
func (c *Client) SetBotToken(token string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.botToken = token
}

// Upload uploads a file to the CDN under the given tag and returns the file ID.
func (c *Client) Upload(ctx context.Context, tag string, filename string, r io.Reader) (string, error) {
	c.mu.RLock()
	sessionToken := c.sessionToken
	botToken := c.botToken
	c.mu.RUnlock()

	if sessionToken == "" && botToken == "" {
		return "", fmt.Errorf("cdn: no authentication token set")
	}

	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)

	go func() {
		part, err := mw.CreateFormFile("file", filename)
		if err != nil {
			pw.CloseWithError(err)
			return
		}
		if _, err := io.Copy(part, r); err != nil {
			pw.CloseWithError(err)
			return
		}
		pw.CloseWithError(mw.Close())
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/"+tag, pr)
	if err != nil {
		return "", fmt.Errorf("cdn: create request: %w", err)
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())

	if botToken != "" {
		req.Header.Set("X-Bot-Token", botToken)
	} else {
		req.Header.Set("X-Session-Token", sessionToken)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("cdn: upload: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("cdn: upload failed (%d): %s", resp.StatusCode, string(body))
	}

	var result struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("cdn: decode response: %w", err)
	}
	return result.ID, nil
}

// URL returns the CDN URL for a file.
func (c *Client) URL(tag, fileID string) string {
	return c.baseURL + "/" + tag + "/" + fileID
}

// OriginalURL returns the CDN URL for the original version of a file.
func (c *Client) OriginalURL(tag, fileID string) string {
	return c.baseURL + "/" + tag + "/" + fileID + "/original"
}
