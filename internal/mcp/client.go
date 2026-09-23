package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

const protocolVersion = "2024-11-05"

type Tool struct {
	Exposed     string
	Server      string
	Name        string
	Description string
	InputSchema json.RawMessage
}

type Client struct {
	name    string
	version string
	tr      transport

	mu     sync.Mutex
	inited bool
}

func NewClient(name, version string, cfg ServerConfig) (*Client, error) {
	tr, err := newTransport(cfg)
	if err != nil {
		return nil, err
	}
	return &Client{name: name, version: version, tr: tr}, nil
}

func (c *Client) Initialized() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.inited
}

func (c *Client) ensure(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.inited {
		return nil
	}
	params := map[string]any{
		"protocolVersion": protocolVersion,
		"capabilities":    map[string]any{},
		"clientInfo":      map[string]any{"name": "cetas-lite", "version": c.version},
	}
	if _, err := c.tr.Call(ctx, "initialize", params); err != nil {
		return fmt.Errorf("initialize: %w", err)
	}
	if err := c.tr.Notify(ctx, "notifications/initialized", nil); err != nil {
		return fmt.Errorf("initialized: %w", err)
	}
	c.inited = true
	return nil
}

func (c *Client) ListTools(ctx context.Context) ([]Tool, error) {
	if err := c.ensure(ctx); err != nil {
		return nil, err
	}
	raw, err := c.tr.Call(ctx, "tools/list", map[string]any{})
	if err != nil {
		return nil, err
	}
	var res struct {
		Tools []struct {
			Name        string          `json:"name"`
			Description string          `json:"description"`
			InputSchema json.RawMessage `json:"inputSchema"`
		} `json:"tools"`
	}
	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, fmt.Errorf("tools/list: %w", err)
	}
	out := make([]Tool, 0, len(res.Tools))
	for _, t := range res.Tools {
		if strings.TrimSpace(t.Name) == "" {
			continue
		}
		out = append(out, Tool{Server: c.name, Name: t.Name, Description: t.Description, InputSchema: t.InputSchema})
	}
	return out, nil
}

func (c *Client) CallTool(ctx context.Context, name string, args map[string]any) (string, error) {
	if err := c.ensure(ctx); err != nil {
		return "", err
	}
	if args == nil {
		args = map[string]any{}
	}
	raw, err := c.tr.Call(ctx, "tools/call", map[string]any{"name": name, "arguments": args})
	if err != nil {
		return "", err
	}
	var res struct {
		Content []struct {
			Type     string `json:"type"`
			Text     string `json:"text"`
			Data     string `json:"data"`
			MimeType string `json:"mimeType"`
		} `json:"content"`
		IsError bool `json:"isError"`
	}
	if err := json.Unmarshal(raw, &res); err != nil {
		return string(raw), nil
	}
	parts := make([]string, 0, len(res.Content))
	for _, item := range res.Content {
		switch {
		case item.Text != "":
			parts = append(parts, item.Text)
		case item.Data != "":
			parts = append(parts, "[image "+item.MimeType+"]")
		default:
			b, _ := json.Marshal(item)
			parts = append(parts, string(b))
		}
	}
	text := strings.Join(parts, "\n")
	if text == "" {
		text = "(resultat vide)"
	}
	if res.IsError {
		text = "[error] " + text
	}
	return text, nil
}

func (c *Client) Close() error { return c.tr.Close() }
