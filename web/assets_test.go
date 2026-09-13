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
	} {
		if _, err := fs.Stat(FS, p); err != nil {
			t.Errorf("asset manquant %s: %v", p, err)
		}
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
}
