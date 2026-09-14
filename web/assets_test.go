package webassets

import (
	"io/fs"
	"strings"
	"testing"
)

func TestAssetsEmbedded(t *testing.T) {
	for _, p := range []string{
		"index.html",
		"css/cetas.css", "css/cetas-ocean.css", "css/cetas-lite.css",
		"css/base/variables.css", "css/base/layout.css",
		"css/features/chat.css", "css/components/components.css",
		"css/themes/ocean.css",
		"images/icons.svg", "images/Cetas42.png", "images/cetas-thinking.png",
		"js/app.js", "js/api.js", "js/auth.js", "js/chat.js",
		"js/model-select.js", "js/settings.js", "js/theme.js", "js/theme-init.js",
		"js/conversations.js", "js/markdown.js", "js/stream-render.js",
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
		`id="composer"`, `id="prompt-input"`, `id="send-btn"`, `id="stop-btn"`, `id="plus-menu-btn"`, `id="plus-menu-dropdown"`,
		`id="settings-overlay"`, `id="aliases-editor"`, `id="aliases-save"`, `id="aliases-status"`,
		`id="agent-toggle"`, `id="web-toggle"`, `id="mcp-toggle"`,
		`id="thinking-toggle"`, `id="effort-select"`, `id="thinking-state"`,
		`id="conv-list"`, `id="stats-badge"`, `id="export-btn"`,
		`id="attach-chips"`, `id="attach-input"`, `id="attach-btn"`, `id="mic-btn"`,
		`id="caps-panel"`, `id="caps-save"`, `id="caps-add"`, `id="caps-add-input"`,
		`class="sidebar"`, `class="main"`, `class="chat-container"`,
		`class="input-area"`, `class="input-wrapper"`,
	} {
		if !strings.Contains(html, id) {
			t.Errorf("index.html manque %s", id)
		}
	}
	for _, src := range []string{
		"/js/vendor/marked.umd.min.js", "/js/vendor/purify.min.js", "/js/app.js",
		"/css/cetas.css", "/css/cetas-lite.css",
	} {
		if !strings.Contains(html, src) {
			t.Errorf("index.html manque %s", src)
		}
	}
}
