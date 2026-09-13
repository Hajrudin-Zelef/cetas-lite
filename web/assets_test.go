package webassets

import (
	"io/fs"
	"strings"
	"testing"
)

func TestAssetsEmbedded(t *testing.T) {
	for _, p := range []string{
		"index.html",
		"css/variables.css", "css/ocean.css", "css/app.css",
		"js/app.js", "js/api.js", "js/auth.js", "js/chat.js",
		"js/model-select.js", "js/settings.js", "js/theme.js",
		"js/markdown.js", "js/stream-render.js",
		"js/vendor/marked.umd.min.js", "js/vendor/purify.min.js",
		"js/package.json",
	} {
		if _, err := fs.Stat(FS, p); err != nil {
			t.Errorf("asset manquant %s: %v", p, err)
		}
	}
}

func TestTestsNotEmbedded(t *testing.T) {
	if _, err := fs.Stat(FS, "js/_tests/stream-render.test.mjs"); err == nil {
		t.Fatal("le dossier _tests ne doit pas etre embarque")
	}
}

func TestIndexHasIDs(t *testing.T) {
	b, err := fs.ReadFile(FS, "index.html")
	if err != nil {
		t.Fatal(err)
	}
	html := string(b)
	for _, id := range []string{
		`id="login-overlay"`, `id="login-form"`, `id="login-username"`, `id="login-password"`, `id="login-error"`,
		`id="app"`, `id="sidebar"`, `id="chat-log"`, `id="new-chat-btn"`, `id="logout-btn"`, `id="settings-btn"`, `id="theme-select"`,
		`id="family-select"`, `id="mode-select"`, `id="route-badge"`,
		`id="composer"`, `id="prompt-input"`, `id="send-btn"`, `id="stop-btn"`,
		`id="settings-overlay"`, `id="aliases-editor"`, `id="aliases-save"`, `id="aliases-status"`,
		`id="agent-toggle"`,
	} {
		if !strings.Contains(html, id) {
			t.Errorf("index.html manque %s", id)
		}
	}
	for _, src := range []string{"/js/vendor/marked.umd.min.js", "/js/vendor/purify.min.js", "/js/app.js"} {
		if !strings.Contains(html, src) {
			t.Errorf("index.html manque le script %s", src)
		}
	}
}
