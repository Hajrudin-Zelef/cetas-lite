package chat

import (
	"errors"
	"fmt"
	"strings"

	"cetas-lite/internal/memory"
	"cetas-lite/internal/provider"
)

type MemoryTools interface {
	List(user string) []memory.Page
	Read(user, name string, offset, limit int) (string, error)
	Add(user, name, content string) error
	Edit(user, name, oldText, newText string) error
	Delete(user, name string) error
	Search(user, query string, limit int) []memory.Hit
	Index(user string) string
}

const memIndexPrefix = "Index memoire"

func MemoryToolSchemas() []provider.Tool {
	str := func(props map[string]any, required ...string) map[string]any {
		return map[string]any{"type": "object", "properties": props, "required": required}
	}
	return []provider.Tool{
		{Type: "function", Function: provider.ToolFunction{Name: "mem_search", Description: "Search your persistent memory pages. Returns files, titles and snippets.", Parameters: str(map[string]any{
			"query": map[string]any{"type": "string"},
			"limit": map[string]any{"type": "integer", "description": "Defaut 8, max 30"},
		}, "query")}},
		{Type: "function", Function: provider.ToolFunction{Name: "mem_read", Description: "Read a memory page, line-numbered, with offset/limit.", Parameters: str(map[string]any{
			"name":   map[string]any{"type": "string", "description": "Ex: notes.md"},
			"offset": map[string]any{"type": "integer", "description": "Premiere ligne (1-based)"},
			"limit":  map[string]any{"type": "integer"},
		}, "name")}},
		{Type: "function", Function: provider.ToolFunction{Name: "mem_add", Description: "Create a new memory page (fails if it exists).", Parameters: str(map[string]any{
			"name":    map[string]any{"type": "string"},
			"content": map[string]any{"type": "string"},
		}, "name", "content")}},
		{Type: "function", Function: provider.ToolFunction{Name: "mem_edit", Description: "Replace one occurrence of text in a memory page.", Parameters: str(map[string]any{
			"name": map[string]any{"type": "string"},
			"old":  map[string]any{"type": "string"},
			"new":  map[string]any{"type": "string"},
		}, "name", "old", "new")}},
		{Type: "function", Function: provider.ToolFunction{Name: "mem_delete", Description: "Delete a memory page.", Parameters: str(map[string]any{
			"name": map[string]any{"type": "string"},
		}, "name")}},
	}
}

func (e *Engine) memExecute(user, name, argsJSON string) ToolResult {
	m := e.memoryTools()
	if m == nil || user == "" {
		return ToolResult{Text: "[erreur] memoire indisponible"}
	}
	args := parseArgs(argsJSON)
	switch name {
	case "mem_search":
		query := strings.TrimSpace(strArg(args, "query"))
		if query == "" {
			return ToolResult{Text: "[erreur] requete vide"}
		}
		hits := m.Search(user, query, intArg(args, "limit"))
		if len(hits) == 0 {
			return ToolResult{Text: "[aucun resultat]"}
		}
		var b strings.Builder
		fmt.Fprintf(&b, "%d resultats:\n", len(hits))
		for _, h := range hits {
			fmt.Fprintf(&b, "- %s -- %s\n  %s\n", h.File, h.Title, h.Snippet)
		}
		return ToolResult{Text: truncate(b.String(), toolMaxOutput)}
	case "mem_read":
		page := strings.TrimSpace(strArg(args, "name"))
		out, err := m.Read(user, page, intArg(args, "offset"), intArg(args, "limit"))
		if err != nil {
			return ToolResult{Text: "[erreur] " + err.Error()}
		}
		return ToolResult{Text: truncate(out, toolMaxOutput)}
	case "mem_add":
		page := strings.TrimSpace(strArg(args, "name"))
		if err := m.Add(user, page, strArg(args, "content")); err != nil {
			return ToolResult{Text: "[erreur] " + err.Error()}
		}
		return ToolResult{Text: "[ok] page " + page + " creee"}
	case "mem_edit":
		page := strings.TrimSpace(strArg(args, "name"))
		err := m.Edit(user, page, strArg(args, "old"), strArg(args, "new"))
		if errors.Is(err, memory.ErrAlreadyApplied) {
			return ToolResult{Text: "[ok] deja a jour"}
		}
		if err != nil {
			return ToolResult{Text: "[erreur] " + err.Error()}
		}
		return ToolResult{Text: "[ok] page " + page + " modifiee"}
	case "mem_delete":
		page := strings.TrimSpace(strArg(args, "name"))
		if err := m.Delete(user, page); err != nil {
			return ToolResult{Text: "[erreur] " + err.Error()}
		}
		return ToolResult{Text: "[ok] page " + page + " supprimee"}
	}
	return ToolResult{Text: "[erreur] outil inconnu: " + name}
}

func (e *Engine) memoryIndexMessage(user string) (provider.Message, bool) {
	m := e.memoryTools()
	if m == nil || user == "" {
		return provider.Message{}, false
	}
	idx := strings.TrimSpace(m.Index(user))
	if idx == "" {
		return provider.Message{}, false
	}
	content := memIndexPrefix + " (titres seulement ; utilise mem_read pour lire une page, " +
		"mem_search pour chercher, mem_add/mem_edit/mem_delete pour la maintenir).\n\n" + idx
	return provider.Message{Role: "system", Content: content}, true
}
