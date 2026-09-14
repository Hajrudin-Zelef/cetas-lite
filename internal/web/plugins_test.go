package web

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"cetas-lite/internal/plugins"
)

func writeTestPlugin(t *testing.T, dir, name string) {
	t.Helper()
	pdir := filepath.Join(dir, name)
	if err := os.MkdirAll(pdir, 0o755); err != nil {
		t.Fatal(err)
	}
	manifest := `{"name": "` + name + `", "version": "1.0.0", "tools": [{"name": "ping", "description": "Ping", "exec": {"command": ["sh", "-c", "printf 'PONG'"]}}]}`
	if err := os.WriteFile(filepath.Join(pdir, "plugin.json"), []byte(manifest), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestPluginsStatusEmpty(t *testing.T) {
	s := newTestServer(t, true)
	s.engine.SetPlugins(plugins.NewManager(t.TempDir(), nil))
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	rec := doJSON(t, h, http.MethodGet, "/api/plugins", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d (%s)", rec.Code, rec.Body.String())
	}
	body := decode(t, rec)
	if got, ok := body["plugins"].([]any); !ok || len(got) != 0 {
		t.Fatalf("plugins attendus vides: %v", body["plugins"])
	}
	if got, ok := body["errors"].([]any); !ok || len(got) != 0 {
		t.Fatalf("errors attendues vides: %v", body["errors"])
	}
}

func TestPluginsStatusLoaded(t *testing.T) {
	s := newTestServer(t, true)
	dir := t.TempDir()
	writeTestPlugin(t, dir, "demo")
	pm := plugins.NewManager(dir, nil)
	if err := pm.Load(); err != nil {
		t.Fatal(err)
	}
	s.engine.SetPlugins(pm)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	rec := doJSON(t, h, http.MethodGet, "/api/plugins", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d (%s)", rec.Code, rec.Body.String())
	}
	body := decode(t, rec)
	list, ok := body["plugins"].([]any)
	if !ok || len(list) != 1 {
		t.Fatalf("1 plugin attendu: %v", body["plugins"])
	}
	first, _ := list[0].(map[string]any)
	if first["name"] != "demo" {
		t.Fatalf("nom inattendu: %v", first)
	}
	tools, _ := first["tools"].([]any)
	if len(tools) != 1 || tools[0] != "plugin_demo_ping" {
		t.Fatalf("outils inattendus: %v", tools)
	}
}

func TestPluginsReload(t *testing.T) {
	s := newTestServer(t, true)
	dir := t.TempDir()
	pm := plugins.NewManager(dir, nil)
	if err := pm.Load(); err != nil {
		t.Fatal(err)
	}
	s.engine.SetPlugins(pm)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	// Ajout d'un plugin APRES le chargement initial : le reload doit le voir.
	writeTestPlugin(t, dir, "nouveau")
	rec := doJSON(t, h, http.MethodPost, "/api/plugins/reload", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d (%s)", rec.Code, rec.Body.String())
	}
	if got, _ := decode(t, rec)["plugins"].([]any); len(got) != 1 {
		t.Fatalf("1 plugin attendu apres reload: %v", decode(t, rec)["plugins"])
	}
}

func TestPluginsUnauthorized(t *testing.T) {
	s := newTestServer(t, true)
	s.engine.SetPlugins(plugins.NewManager(t.TempDir(), nil))
	h := s.Handler()
	if rec := doJSON(t, h, http.MethodGet, "/api/plugins", "", nil); rec.Code != http.StatusUnauthorized {
		t.Fatalf("GET sans token = %d", rec.Code)
	}
	if rec := doJSON(t, h, http.MethodPost, "/api/plugins/reload", "", nil); rec.Code != http.StatusUnauthorized {
		t.Fatalf("POST sans token = %d", rec.Code)
	}
}
