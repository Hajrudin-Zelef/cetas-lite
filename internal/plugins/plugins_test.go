package plugins

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func writePlugin(t *testing.T, dir, name, manifest string, files map[string]string) string {
	t.Helper()
	pdir := filepath.Join(dir, name)
	if err := os.MkdirAll(pdir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(pdir, "plugin.json"), []byte(manifest), 0o644); err != nil {
		t.Fatal(err)
	}
	for f, c := range files {
		if err := os.WriteFile(filepath.Join(pdir, f), []byte(c), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func TestLoadValid(t *testing.T) {
	dir := writePlugin(t, t.TempDir(), "demo", `{
		"name": "demo", "version": "1.0.0", "description": "Demo",
		"tools": [{"name": "echo", "description": "Renvoie le texte",
			"parameters": {"type": "object", "properties": {"texte": {"type": "string"}}},
			"exec": {"command": ["sh", "echo.sh"]}}]
	}`, map[string]string{"echo.sh": "#!/bin/sh\ncat\n"})
	m := NewManager(dir, nil)
	if err := m.Load(); err != nil {
		t.Fatal(err)
	}
	if len(m.Errors()) != 0 {
		t.Fatalf("erreurs inattendues: %v", m.Errors())
	}
	infos := m.Plugins()
	if len(infos) != 1 || infos[0].Name != "demo" {
		t.Fatalf("plugins: %+v", infos)
	}
	defs := m.Defs()
	if len(defs) != 1 || defs[0].Name != "plugin_demo_echo" {
		t.Fatalf("defs: %+v", defs)
	}
}

func TestExecRawText(t *testing.T) {
	dir := writePlugin(t, t.TempDir(), "txt", `{
		"name": "txt", "version": "1.0.0",
		"tools": [{"name": "maj", "description": "Majuscules",
			"exec": {"command": ["sh", "-c", "tr a-z A-Z"]}}]
	}`, nil)
	m := NewManager(dir, nil)
	if err := m.Load(); err != nil {
		t.Fatal(err)
	}
	out, err := m.Execute(context.Background(), "plugin_txt_maj", `{"x":"abc"}`)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "ABC") {
		t.Fatalf("sortie inattendue: %q", out)
	}
}

func TestExecJSONProtocol(t *testing.T) {
	script := "#!/bin/sh\n" +
		"cat > /dev/null\n" +
		"printf '{\"result\": \"pong\"}'\n"
	dir := writePlugin(t, t.TempDir(), "js", `{
		"name": "js", "version": "1.0.0",
		"tools": [{"name": "ping", "description": "Ping", "exec": {"command": ["sh", "ping.sh"]}}]
	}`, map[string]string{"ping.sh": script})
	m := NewManager(dir, nil)
	if err := m.Load(); err != nil {
		t.Fatal(err)
	}
	out, err := m.Execute(context.Background(), "plugin_js_ping", `{"a":1}`)
	if err != nil {
		t.Fatal(err)
	}
	if out != "pong" {
		t.Fatalf("protocole JSON non respecte: %q", out)
	}
}

func TestExecErrorProtocol(t *testing.T) {
	dir := writePlugin(t, t.TempDir(), "ko", `{
		"name": "ko", "version": "1.0.0",
		"tools": [{"name": "boom", "description": "Echec", "exec": {"command": ["sh", "-c", "printf '{\"error\": \"rate !\"}'"]}}]
	}`, nil)
	m := NewManager(dir, nil)
	if err := m.Load(); err != nil {
		t.Fatal(err)
	}
	if _, err := m.Execute(context.Background(), "plugin_ko_boom", `{}`); err == nil {
		t.Fatal("erreur plugin attendue")
	}
}

func TestExecTimeout(t *testing.T) {
	dir := writePlugin(t, t.TempDir(), "lent", `{
		"name": "lent", "version": "1.0.0",
		"tools": [{"name": "dors", "description": "Lent", "exec": {"command": ["sh", "-c", "sleep 30"], "timeout_sec": 1}}]
	}`, nil)
	m := NewManager(dir, nil)
	if err := m.Load(); err != nil {
		t.Fatal(err)
	}
	start := time.Now()
	_, err := m.Execute(context.Background(), "plugin_lent_dors", `{}`)
	if err == nil {
		t.Fatal("timeout attendu")
	}
	if time.Since(start) > 10*time.Second {
		t.Fatal("le timeout n'a pas tue le processus assez vite")
	}
}

func TestLoadRejeteInvalides(t *testing.T) {
	dir := t.TempDir()
	// Nom de dossier != nom du manifeste.
	writePlugin(t, dir, "dossier", `{"name": "autre", "version": "1.0.0",
		"tools": [{"name": "t", "description": "d", "exec": {"command": ["sh"]}}]}`, nil)
	// Ni exec ni http.
	writePlugin(t, dir, "vide", `{"name": "vide", "version": "1.0.0",
		"tools": [{"name": "t", "description": "d"}]}`, nil)
	// Evasion de dossier.
	writePlugin(t, dir, "evil", `{"name": "evil", "version": "1.0.0",
		"tools": [{"name": "t", "description": "d", "exec": {"command": ["sh"], "dir": "../.."}}]}`, nil)
	// JSON invalide.
	pdir := filepath.Join(dir, "cassé")
	_ = os.MkdirAll(pdir, 0o755)
	_ = os.WriteFile(filepath.Join(pdir, "plugin.json"), []byte("{pas json"), 0o644)

	m := NewManager(dir, nil)
	if err := m.Load(); err != nil {
		t.Fatal(err)
	}
	if len(m.Plugins()) != 0 {
		t.Fatalf("aucun plugin valide attendu, obtenu: %+v", m.Plugins())
	}
	if len(m.Errors()) != 4 {
		t.Fatalf("4 erreurs attendues, obtenues: %v", m.Errors())
	}
}

func TestHTTPTool(t *testing.T) {
	var gotBody string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw := make([]byte, r.ContentLength)
		_, _ = r.Body.Read(raw)
		gotBody = string(raw)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"result": "ok-http"})
	}))
	defer srv.Close()
	dir := writePlugin(t, t.TempDir(), "web", `{
		"name": "web", "version": "2.0.0",
		"tools": [{"name": "appel", "description": "Appel HTTP",
			"http": {"url": "`+srv.URL+`", "method": "POST"}}]
	}`, nil)
	m := NewManager(dir, srv.Client())
	if err := m.Load(); err != nil {
		t.Fatal(err)
	}
	out, err := m.Execute(context.Background(), "plugin_web_appel", `{"q":"x"}`)
	if err != nil {
		t.Fatal(err)
	}
	if out != "ok-http" {
		t.Fatalf("resultat http inattendu: %q", out)
	}
	if !strings.Contains(gotBody, `"q":"x"`) {
		t.Fatalf("corps non transmis: %q", gotBody)
	}
}

func TestExecuteInconnu(t *testing.T) {
	m := NewManager(t.TempDir(), nil)
	if err := m.Load(); err != nil {
		t.Fatal(err)
	}
	if _, err := m.Execute(context.Background(), "plugin_x_y", `{}`); err == nil {
		t.Fatal("erreur attendue pour outil inconnu")
	}
}

func TestConfinedDir(t *testing.T) {
	base := t.TempDir()
	if _, err := confinedDir(base, ""); err != nil {
		t.Fatal(err)
	}
	if _, err := confinedDir(base, "sous/dossier"); err != nil {
		t.Fatal(err)
	}
	if _, err := confinedDir(base, "../echappe"); err == nil {
		t.Fatal("evasion .. doit etre refusee")
	}
	if _, err := confinedDir(base, "/absolu"); err == nil {
		t.Fatal("chemin absolu doit etre refuse")
	}
}
