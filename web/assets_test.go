package webassets

import (
	"io/fs"
	"strings"
	"testing"
)

func TestAssetsEmbedded(t *testing.T) {
	for _, p := range []string{
		"index.html",
		"css/cetas-lite.css",
		"css/features/agents-terminal.css",
		"css/ref/base/variables.css", "css/ref/base/layout.css",
		"css/ref/features/chat.css", "css/ref/components/components.css",
		"css/ref/themes/ocean.css", "css/ref/ocean.css",
		"images/icons.svg", "images/Cetas42.png", "images/cetas-thinking.png",
		"js/app.js", "js/api.js", "js/auth.js", "js/chat.js",
		"js/model-select.js", "js/sidebar.js", "js/right-panel.js", "js/modals.js",
		"js/agents.js", "js/terminal.js", "js/thread-view.js",
		"js/theme-init.js",
		"js/markdown.js", "js/stream-render.js",
		"js/vendor/marked.umd.min.js", "js/vendor/purify.min.js",
		"js/vendor/highlight.esm.min.js", "js/vendor/highlight-github-dark.min.css",
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
		`id="kiro-splash"`, `id="sidebar"`, `id="sidebar-toggle"`, `id="conv-list"`, `id="conv-search"`,
		`id="new-chat-btn"`, `id="logout-btn"`, `id="apikeys-btn"`, `id="clear-all-btn"`,
		`id="family-select"`, `id="mode-select"`,
		`id="chat-container"`, `id="empty-chat-placeholder"`, `id="model-alert"`,
		`id="prompt-input"`, `id="send-btn"`, `id="stop-btn"`, `id="mic-btn"`,
		`id="plus-menu-btn"`, `id="plus-menu-dropdown"`, `id="file-input"`, `id="attach-preview"`,
		`id="token-info"`, `id="token-bar"`, `id="input-hint"`,
		`id="share-btn"`, `id="share-menu"`,
		`id="chat-header-settings"`, `id="right-panel"`,
		`id="apikeys-modal-overlay"`, `id="providers-list"`, `id="families-list"`,
		`id="save-modal-overlay"`, `id="cat-modal-overlay"`,
		`id="roles-manage-overlay"`, `id="prompts-manage-overlay"`,
		`id="custom-dialog-overlay"`,
		`class="input-area"`, `class="input-wrapper"`,
	} {
		if !strings.Contains(html, id) {
			t.Errorf("index.html manque %s", id)
		}
	}
	for _, src := range []string{
		"/js/vendor/marked.umd.min.js", "/js/vendor/purify.min.js", "/js/app.js",
		"/css/ref/base/variables.css", "/css/cetas-lite.css", "/css/features/agents-terminal.css",
	} {
		if !strings.Contains(html, src) {
			t.Errorf("index.html manque %s", src)
		}
	}
}
