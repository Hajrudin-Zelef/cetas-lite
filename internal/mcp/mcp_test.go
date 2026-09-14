package mcp

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func helperConfig() ServerConfig {
	return ServerConfig{Command: os.Args[0], Env: map[string]string{"CETAS_MCP_HELPER": "1"}}
}

func TestLoadConfigMissing(t *testing.T) {
	cfg, err := LoadConfig(filepath.Join(t.TempDir(), "absent.json"))
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.MCP) != 0 {
		t.Fatalf("config vide attendue: %+v", cfg)
	}
}

func TestLoadConfigValid(t *testing.T) {
	path := filepath.Join(t.TempDir(), "mcp.json")
	if err := os.WriteFile(path, []byte(`{"mcp":{"a":{"command":"echo","args":["x"]},"b":{"type":"http","url":"http://x"}}}`), 0o600); err != nil {
		t.Fatal(err)
	}
	cfg, err := LoadConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.MCP) != 2 || cfg.MCP["a"].Command != "echo" || cfg.MCP["b"].URL != "http://x" {
		t.Fatalf("config inattendue: %+v", cfg)
	}
}

func TestLoadConfigInvalid(t *testing.T) {
	path := filepath.Join(t.TempDir(), "mcp.json")
	if err := os.WriteFile(path, []byte(`{not json`), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadConfig(path); err == nil {
		t.Fatal("erreur attendue pour JSON invalide")
	}
}

func TestTransportResolution(t *testing.T) {
	cases := []struct {
		cfg  ServerConfig
		want string
		ok   bool
	}{
		{ServerConfig{Command: "echo"}, "stdio", true},
		{ServerConfig{Type: "stdio", Command: "echo"}, "stdio", true},
		{ServerConfig{URL: "http://x"}, "http", true},
		{ServerConfig{Type: "http", URL: "http://x"}, "http", true},
		{ServerConfig{Type: "sse", URL: "http://x"}, "", false},
		{ServerConfig{}, "", false},
		{ServerConfig{Type: "stdio"}, "", false},
	}
	for i, c := range cases {
		got, err := c.cfg.transport()
		if c.ok && (err != nil || got != c.want) {
			t.Fatalf("cas %d: got=%q err=%v", i, got, err)
		}
		if !c.ok && err == nil {
			t.Fatalf("cas %d: erreur attendue", i)
		}
	}
}

func TestStdioInitializeListCall(t *testing.T) {
	c, err := NewClient("fake", "test", helperConfig())
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = c.Close() }()

	ctx := context.Background()
	tools, err := c.ListTools(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(tools) != 2 {
		t.Fatalf("tools=%d", len(tools))
	}
	out, err := c.CallTool(ctx, "echo", map[string]any{"text": "bonjour"})
	if err != nil {
		t.Fatal(err)
	}
	if out != "bonjour" {
		t.Fatalf("echo=%q", out)
	}
	out, err = c.CallTool(ctx, "add", map[string]any{"a": 2, "b": 3})
	if err != nil {
		t.Fatal(err)
	}
	if out != "5" {
		t.Fatalf("add=%q", out)
	}
}

func TestStdioCallError(t *testing.T) {
	c, err := NewClient("fake", "test", helperConfig())
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = c.Close() }()
	out, err := c.CallTool(context.Background(), "inconnu", map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(out, "[erreur]") {
		t.Fatalf("isError attendu: %q", out)
	}
}

func httpServer(t *testing.T, sse bool) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		var msg rpcMessage
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		resp := rpcMessage{JSONRPC: "2.0", ID: msg.ID}
		switch msg.Method {
		case "initialize":
			resp.Result = json.RawMessage(`{"protocolVersion":"2024-11-05","capabilities":{}}`)
		case "tools/list":
			resp.Result = json.RawMessage(`{"tools":[{"name":"ping","description":"Ping","inputSchema":{"type":"object"}}]}`)
		case "tools/call":
			resp.Result = json.RawMessage(`{"content":[{"type":"text","text":"pong"}],"isError":false}`)
		default:
			resp.Error = &rpcError{Code: -32601, Message: "nope"}
		}
		w.Header().Set("Mcp-Session-Id", "sess-1")
		if sse {
			w.Header().Set("Content-Type", "text/event-stream")
			data, _ := json.Marshal(resp)
			_, _ = w.Write([]byte("event: message\ndata: " + string(data) + "\n\n"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	}))
}

func TestHTTPTransportListCall(t *testing.T) {
	ts := httpServer(t, false)
	defer ts.Close()
	c, err := NewClient("remote", "test", ServerConfig{Type: "http", URL: ts.URL})
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = c.Close() }()
	ctx := context.Background()
	tools, err := c.ListTools(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(tools) != 1 || tools[0].Name != "ping" {
		t.Fatalf("tools=%+v", tools)
	}
	out, err := c.CallTool(ctx, "ping", nil)
	if err != nil {
		t.Fatal(err)
	}
	if out != "pong" {
		t.Fatalf("ping=%q", out)
	}
}

func TestHTTPTransportSSE(t *testing.T) {
	ts := httpServer(t, true)
	defer ts.Close()
	c, err := NewClient("remote", "test", ServerConfig{Type: "http", URL: ts.URL})
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = c.Close() }()
	out, err := c.CallTool(context.Background(), "ping", nil)
	if err != nil {
		t.Fatal(err)
	}
	if out != "pong" {
		t.Fatalf("ping=%q", out)
	}
}

func TestExposeName(t *testing.T) {
	if got := exposeName("context7", "resolve-library-id"); got != "mcp_context7_resolve-library-id" {
		t.Fatalf("got=%q", got)
	}
	if got := exposeName("a.b", "x/y"); got != "mcp_a_b_x_y" {
		t.Fatalf("sanitize got=%q", got)
	}
	long := exposeName(strings.Repeat("s", 80), "t")
	if len(long) != 64 {
		t.Fatalf("len=%d", len(long))
	}
}
