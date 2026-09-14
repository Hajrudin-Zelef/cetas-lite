package chat

import (
	"os"
	"path/filepath"
	"strings"

	"cetas-lite/internal/provider"
)

const marexMaxBytes = 16 << 10

func (e *Engine) marexMessage(user string) (provider.Message, bool) {
	ws, err := userWorkspacePath(e.workspace, user)
	if err != nil {
		return provider.Message{}, false
	}
	var parts []string
	global := filepath.Join(filepath.Dir(e.workspace), "MAREX.md")
	project := filepath.Join(ws, "MAREX.md")
	for _, path := range []string{global, project} {
		if s := readCapped(path); s != "" {
			parts = append(parts, s)
		}
	}
	if len(parts) == 0 {
		return provider.Message{}, false
	}
	return provider.Message{Role: "system", Content: "[MAREX.md]\n\n" + strings.Join(parts, "\n\n")}, true
}

func readCapped(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	if len(b) > marexMaxBytes {
		b = b[:marexMaxBytes]
	}
	return strings.TrimSpace(string(b))
}
