package customtools

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	maxResponse = 64 << 10
	maxTimeout  = 60
)

type Tool struct {
	Description string            `json:"description,omitempty"`
	Method      string            `json:"method,omitempty"`
	URL         string            `json:"url"`
	Headers     map[string]string `json:"headers,omitempty"`
	Body        string            `json:"body,omitempty"`
	TimeoutSec  int               `json:"timeout_sec,omitempty"`
	Parameters  map[string]any    `json:"parameters,omitempty"`
	Required    []string          `json:"required,omitempty"`
}

type Config struct {
	Tools map[string]Tool `json:"tools"`
}

type Def struct {
	Name        string
	Description string
	Schema      map[string]any
}

type Manager struct {
	tools  map[string]Tool
	order  []string
	client *http.Client
}

func LoadConfig(path string) (Config, error) {
	empty := Config{Tools: map[string]Tool{}}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return empty, nil
		}
		return empty, fmt.Errorf("lecture %s: %w", path, err)
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return empty, fmt.Errorf("tools.json invalide: %w", err)
	}
	if cfg.Tools == nil {
		cfg.Tools = map[string]Tool{}
	}
	return cfg, nil
}

func NewManager(path string, client *http.Client) (*Manager, error) {
	cfg, err := LoadConfig(path)
	if err != nil {
		return nil, err
	}
	if client == nil {
		client = &http.Client{}
	}
	m := &Manager{tools: map[string]Tool{}, client: client}
	for name, t := range cfg.Tools {
		if strings.TrimSpace(t.URL) == "" {
			continue
		}
		m.tools[ExposedName(name)] = t
	}
	for exposed := range m.tools {
		m.order = append(m.order, exposed)
	}
	sort.Strings(m.order)
	return m, nil
}

func (m *Manager) Configured() bool { return m != nil && len(m.tools) > 0 }

func (m *Manager) Defs() []Def {
	if m == nil {
		return nil
	}
	out := make([]Def, 0, len(m.order))
	for _, exposed := range m.order {
		t := m.tools[exposed]
		props := t.Parameters
		if props == nil {
			props = map[string]any{}
		}
		schema := map[string]any{"type": "object", "properties": props}
		if len(t.Required) > 0 {
			schema["required"] = t.Required
		}
		desc := t.Description
		if desc == "" {
			desc = "Appel HTTP " + strings.ToUpper(methodOf(t)) + " " + t.URL
		}
		out = append(out, Def{Name: exposed, Description: desc, Schema: schema})
	}
	return out
}

func (m *Manager) Call(ctx context.Context, name, argsJSON string) (string, error) {
	if m == nil {
		return "", fmt.Errorf("outils personnalises indisponibles")
	}
	t, ok := m.tools[name]
	if !ok {
		return "", fmt.Errorf("outil inconnu: %s", name)
	}
	args := map[string]any{}
	if strings.TrimSpace(argsJSON) != "" {
		_ = json.Unmarshal([]byte(argsJSON), &args)
	}
	method := methodOf(t)
	rawURL, err := buildURL(t.URL, method, args)
	if err != nil {
		return "", err
	}
	var bodyReader io.Reader
	contentType := ""
	if t.Body != "" {
		body := substitute(t.Body, args, func(s string) string { return s })
		bodyReader = strings.NewReader(body)
		contentType = "application/json"
	} else if method != http.MethodGet && method != http.MethodHead && method != http.MethodDelete && len(args) > 0 {
		data, err := json.Marshal(args)
		if err != nil {
			return "", err
		}
		bodyReader = bytes.NewReader(data)
		contentType = "application/json"
	}

	timeout := t.TimeoutSec
	if timeout <= 0 {
		timeout = 30
	}
	if timeout > maxTimeout {
		timeout = maxTimeout
	}
	cctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(cctx, method, rawURL, bodyReader)
	if err != nil {
		return "", err
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Set("Accept", "application/json, text/plain;q=0.9, */*;q=0.8")
	for k, v := range t.Headers {
		req.Header.Set(k, v)
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(io.LimitReader(resp.Body, maxResponse))
	text := strings.TrimSpace(string(data))
	if text == "" {
		text = "(reponse vide)"
	}
	if resp.StatusCode/100 != 2 {
		return fmt.Sprintf("[http %d] %s", resp.StatusCode, text), nil
	}
	return text, nil
}

func methodOf(t Tool) string {
	m := strings.ToUpper(strings.TrimSpace(t.Method))
	if m == "" {
		return http.MethodGet
	}
	return m
}

func buildURL(tpl, method string, args map[string]any) (string, error) {
	referenced := referencedKeys(tpl)
	raw := substitute(tpl, args, url.PathEscape)
	u, err := url.Parse(raw)
	if err != nil {
		return "", fmt.Errorf("url invalide: %w", err)
	}
	if method == http.MethodGet || method == http.MethodHead || method == http.MethodDelete {
		q := u.Query()
		for k, v := range args {
			if referenced[k] {
				continue
			}
			q.Set(k, fmt.Sprint(v))
		}
		u.RawQuery = q.Encode()
	}
	return u.String(), nil
}

var placeholderRE = regexp.MustCompile(`\{([A-Za-z0-9_]+)\}`)

func substitute(tpl string, args map[string]any, escape func(string) string) string {
	return placeholderRE.ReplaceAllStringFunc(tpl, func(m string) string {
		key := m[1 : len(m)-1]
		v, ok := args[key]
		if !ok {
			return m
		}
		return escape(fmt.Sprint(v))
	})
}

func referencedKeys(tpl string) map[string]bool {
	out := map[string]bool{}
	for _, m := range placeholderRE.FindAllStringSubmatch(tpl, -1) {
		out[m[1]] = true
	}
	return out
}

func ExposedName(name string) string {
	var b strings.Builder
	b.WriteString("custom_")
	for _, r := range name {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '_', r == '-':
			b.WriteRune(r)
		default:
			b.WriteByte('_')
		}
	}
	out := b.String()
	if len(out) > 64 {
		out = out[:64]
	}
	return out
}
