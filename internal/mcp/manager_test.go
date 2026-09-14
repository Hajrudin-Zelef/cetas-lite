package mcp

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func writeConfig(t *testing.T, dir string, mcpServers map[string]ServerConfig) string {
	t.Helper()
	path := filepath.Join(dir, "mcp.json")
	data, err := json.Marshal(Config{MCP: mcpServers})
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestManagerToolsAndCall(t *testing.T) {
	path := writeConfig(t, t.TempDir(), map[string]ServerConfig{
		"fake": {Command: os.Args[0], Env: map[string]string{"CETAS_MCP_HELPER": "1"}},
	})
	m, err := NewManager(path, "test")
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()
	if !m.Configured() {
		t.Fatal("doit etre configure")
	}

	ctx := context.Background()
	tools := m.Tools(ctx)
	if len(tools) != 2 {
		t.Fatalf("tools=%d", len(tools))
	}
	found := false
	for _, tl := range tools {
		if tl.Exposed == "mcp_fake_echo" {
			found = true
		}
	}
	if !found {
		t.Fatalf("mcp_fake_echo absent: %+v", tools)
	}

	out, err := m.Call(ctx, "mcp_fake_echo", map[string]any{"text": "salut"})
	if err != nil {
		t.Fatal(err)
	}
	if out != "salut" {
		t.Fatalf("call=%q", out)
	}

	servers := m.Servers()
	if len(servers) != 1 || !servers[0].Connected || servers[0].Tools != 2 {
		t.Fatalf("statut=%+v", servers)
	}
}

func TestManagerUnconfigured(t *testing.T) {
	m, err := NewManager(filepath.Join(t.TempDir(), "absent.json"), "test")
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()
	if m.Configured() {
		t.Fatal("ne doit pas etre configure")
	}
	if got := m.Tools(context.Background()); got != nil {
		t.Fatalf("tools attendus nuls: %+v", got)
	}
	if _, err := m.Call(context.Background(), "mcp_x_y", nil); err == nil {
		t.Fatal("erreur attendue")
	}
}

func TestManagerServerFailureIsSkipped(t *testing.T) {
	path := writeConfig(t, t.TempDir(), map[string]ServerConfig{
		"mort": {Command: "/chemin/inexistant/binaire"},
		"fake": {Command: os.Args[0], Env: map[string]string{"CETAS_MCP_HELPER": "1"}},
	})
	m, err := NewManager(path, "test")
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()
	tools := m.Tools(context.Background())
	if len(tools) != 2 {
		t.Fatalf("le serveur mort doit etre ignore, tools=%d", len(tools))
	}
	var foundErr bool
	for _, s := range m.Servers() {
		if s.Name == "mort" && s.Error != "" {
			foundErr = true
		}
	}
	if !foundErr {
		t.Fatalf("erreur attendue pour serveur mort: %+v", m.Servers())
	}
}
