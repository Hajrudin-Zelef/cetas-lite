package chat

import (
	"encoding/json"
	"fmt"
	"strings"

	"cetas-lite/internal/provider"
)

func (e *Engine) ExportActive(user, format string) (string, string, bool) {
	if e.st == nil {
		return "", "", false
	}
	s := e.Conversation(user).marshal()
	if format == "json" {
		data, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			return "", "", false
		}
		return s.ID, string(data), true
	}
	return s.ID, exportMarkdown(s.Messages), true
}

func (e *Engine) ExportArchive(user, id, format string) (string, bool) {
	if e.st == nil {
		return "", false
	}
	raw, ok := e.st.GetArchive(user, id)
	if !ok {
		return "", false
	}
	if format == "json" {
		return string(raw), true
	}
	var s snapshot
	if json.Unmarshal(raw, &s) != nil {
		return "", false
	}
	return exportMarkdown(s.Messages), true
}

func exportMarkdown(msgs []provider.Message) string {
	var b strings.Builder
	b.WriteString("# Conversation\n\n")
	for _, m := range msgs {
		switch m.Role {
		case "user":
			if s, ok := m.Content.(string); ok {
				fmt.Fprintf(&b, "## Vous\n\n%s\n\n", s)
			}
		case "assistant":
			if s, ok := m.Content.(string); ok && strings.TrimSpace(s) != "" {
				fmt.Fprintf(&b, "## Assistant\n\n%s\n\n", s)
			}
			for _, tc := range m.ToolCalls {
				fmt.Fprintf(&b, "**Outil `%s`**\n\n```json\n%s\n```\n\n", tc.Function.Name, tc.Function.Arguments)
			}
		case "tool":
			if s, ok := m.Content.(string); ok {
				fmt.Fprintf(&b, "> Résultat d'outil\n\n```\n%s\n```\n\n", s)
			}
		}
	}
	return strings.TrimRight(b.String(), "\n") + "\n"
}
