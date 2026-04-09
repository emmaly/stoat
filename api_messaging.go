package stoat

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

// SendMessage sends a message to a channel. Requires session token.
func (c *Client) SendMessage(ctx context.Context, channelID string, data DataMessageSend) (*Message, error) {
	req, err := c.request(ctx, http.MethodPost, fmt.Sprintf("/channels/%s/messages", channelID), data)
	if err != nil {
		return nil, err
	}
	if data.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", data.IdempotencyKey)
	}
	var msg Message
	if err := c.do(req, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

// FetchMessages fetches messages from a channel. Requires session token.
func (c *Client) FetchMessages(ctx context.Context, channelID string, opts *FetchMessagesOptions) (*BulkMessageResponse, error) {
	path := fmt.Sprintf("/channels/%s/messages", channelID)
	if opts != nil {
		q := ""
		sep := "?"
		addParam := func(key, value string) {
			q += sep + key + "=" + value
			sep = "&"
		}
		if opts.Limit != nil {
			addParam("limit", strconv.Itoa(*opts.Limit))
		}
		if opts.Before != nil {
			addParam("before", *opts.Before)
		}
		if opts.After != nil {
			addParam("after", *opts.After)
		}
		if opts.Sort != nil {
			addParam("sort", string(*opts.Sort))
		}
		if opts.Nearby != nil {
			addParam("nearby", *opts.Nearby)
		}
		if opts.IncludeUsers != nil && *opts.IncludeUsers {
			addParam("include_users", "true")
		}
		path += q
	}
	req, err := c.request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var resp BulkMessageResponse
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FetchMessage fetches a single message by ID. Requires session token.
func (c *Client) FetchMessage(ctx context.Context, channelID, msgID string) (*Message, error) {
	req, err := c.request(ctx, http.MethodGet, fmt.Sprintf("/channels/%s/messages/%s", channelID, msgID), nil)
	if err != nil {
		return nil, err
	}
	var msg Message
	if err := c.do(req, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

// EditMessage edits a message. Requires session token.
func (c *Client) EditMessage(ctx context.Context, channelID, msgID string, data DataEditMessage) (*Message, error) {
	req, err := c.request(ctx, http.MethodPatch, fmt.Sprintf("/channels/%s/messages/%s", channelID, msgID), data)
	if err != nil {
		return nil, err
	}
	var msg Message
	if err := c.do(req, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

// DeleteMessage deletes a message. Requires session token.
func (c *Client) DeleteMessage(ctx context.Context, channelID, msgID string) error {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/channels/%s/messages/%s", channelID, msgID), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// BulkDeleteMessages deletes multiple messages. Requires session token.
func (c *Client) BulkDeleteMessages(ctx context.Context, channelID string, data OptionsBulkDelete) error {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/channels/%s/messages/bulk", channelID), data)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// SearchMessages searches for messages in a channel. Requires session token.
func (c *Client) SearchMessages(ctx context.Context, channelID string, data DataMessageSearch) (*BulkMessageResponse, error) {
	req, err := c.request(ctx, http.MethodPost, fmt.Sprintf("/channels/%s/search", channelID), data)
	if err != nil {
		return nil, err
	}
	var resp BulkMessageResponse
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PinMessage pins a message in a channel. Requires session token.
func (c *Client) PinMessage(ctx context.Context, channelID, msgID string) error {
	req, err := c.request(ctx, http.MethodPost, fmt.Sprintf("/channels/%s/messages/%s/pin", channelID, msgID), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// UnpinMessage unpins a message from a channel. Requires session token.
func (c *Client) UnpinMessage(ctx context.Context, channelID, msgID string) error {
	req, err := c.request(ctx, http.MethodDelete, fmt.Sprintf("/channels/%s/messages/%s/pin", channelID, msgID), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// AcknowledgeMessage marks messages as read up to a given message. Requires session token.
func (c *Client) AcknowledgeMessage(ctx context.Context, channelID, msgID string) error {
	req, err := c.request(ctx, http.MethodPut, fmt.Sprintf("/channels/%s/ack/%s", channelID, msgID), nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}
