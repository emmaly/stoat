package stoat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

// Client is the Stoat API client.
type Client struct {
	baseURL    string
	httpClient *http.Client

	mu           sync.RWMutex
	sessionToken string
	botToken     string
	mfaTicket    string
	lastRate     *RateLimit
}

// Option configures the Client.
type Option func(*Client)

// WithHTTPClient sets a custom http.Client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.httpClient = hc
	}
}

// New creates a new Stoat API client.
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

// SetMFATicket sets the MFA ticket header.
func (c *Client) SetMFATicket(ticket string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.mfaTicket = ticket
}

// LastRateLimit returns the rate limit info from the most recent response.
func (c *Client) LastRateLimit() *RateLimit {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.lastRate
}

// request builds an HTTP request with auth headers.
func (c *Client) request(ctx context.Context, method, path string, body any) (*http.Request, error) {
	url := c.baseURL + path

	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.sessionToken != "" {
		req.Header.Set("X-Session-Token", c.sessionToken)
	}
	if c.botToken != "" {
		req.Header.Set("X-Bot-Token", c.botToken)
	}
	if c.mfaTicket != "" {
		req.Header.Set("X-MFA-Ticket", c.mfaTicket)
	}

	return req, nil
}

// do executes the request, parses rate limits, and decodes the response.
// If result is non-nil, the response body is JSON-decoded into it.
func (c *Client) do(req *http.Request, result any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Parse rate limit headers from every response.
	if rl := ParseRateLimit(resp.Header); rl != nil {
		c.mu.Lock()
		c.lastRate = rl
		c.mu.Unlock()
	}

	// Read body.
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	// Check for error status codes.
	if resp.StatusCode >= 400 {
		var apiErr APIError
		if err := json.Unmarshal(data, &apiErr); err != nil {
			return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(data))
		}
		return &apiErr
	}

	// Decode success response.
	if result != nil && len(data) > 0 {
		if err := json.Unmarshal(data, result); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}
