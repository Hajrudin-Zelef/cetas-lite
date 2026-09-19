package chat

import (
	"context"
	"fmt"
	"strings"

	"cetas-lite/internal/provider"
	"cetas-lite/internal/search"
)

type WebTools interface {
	Search(ctx context.Context, query string, maxResults int) search.Result
	Fetch(ctx context.Context, rawURL string) (string, string, error)
}

func WebToolSchemas() []provider.Tool {
	return []provider.Tool{
		{Type: "function", Function: provider.ToolFunction{
			Name:        "web_search",
			Description: "Cherche des informations récentes sur le web. Renvoie titres, URLs et extraits.",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"query":       map[string]any{"type": "string"},
					"max_results": map[string]any{"type": "integer", "description": "Default 5, max 10"},
				},
				"required": []string{"query"},
			},
		}},
		{Type: "function", Function: provider.ToolFunction{
			Name:        "web_fetch",
			Description: "Récupère une URL et renvoie son contenu lisible en Markdown.",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"url": map[string]any{"type": "string"},
				},
				"required": []string{"url"},
			},
		}},
	}
}

func (e *Engine) webExecute(ctx context.Context, user, name, argsJSON string) ToolResult {
	wt := e.webTools()
	if wt == nil {
		return ToolResult{Text: "[erreur] recherche web indisponible"}
	}
	if !e.allowWeb(user) {
		return ToolResult{Text: "[erreur] limite de recherche web atteinte, reessaie dans une minute"}
	}
	args := parseArgs(argsJSON)
	switch name {
	case "web_search":
		query := strings.TrimSpace(strArg(args, "query"))
		if query == "" {
			return ToolResult{Text: "[erreur] requete vide"}
		}
		res := wt.Search(ctx, query, intArg(args, "max_results"))
		if len(res.Hits) == 0 {
			msg := res.Error
			if msg == "" {
				msg = "aucun resultat"
			}
			return ToolResult{Text: "[info] recherche web: " + msg}
		}
		return ToolResult{Text: formatHits(res), Meta: searchResultMeta(res)}
	case "web_fetch":
		raw := strings.TrimSpace(strArg(args, "url"))
		if raw == "" {
			return ToolResult{Text: "[erreur] url vide"}
		}
		title, md, err := wt.Fetch(ctx, raw)
		if err != nil {
			return ToolResult{Text: "[erreur] fetch: " + err.Error()}
		}
		return ToolResult{Text: formatFetch(title, md)}
	}
	return ToolResult{Text: "[erreur] outil inconnu: " + name}
}

func formatHits(res search.Result) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Resultats %s (%d):\n", res.Provider, len(res.Hits))
	for i, h := range res.Hits {
		fmt.Fprintf(&b, "[%d] %s\n%s\n%s\n", i+1, h.Title, h.URL, h.Description)
	}
	return truncate(b.String(), toolMaxOutput)
}

// searchResultMeta expose les sources d'une recherche au frontend (panneau
// "Sources" + indicateur visuel), en plus du texte pour le modele.
func searchResultMeta(res search.Result) map[string]any {
	srcs := make([]map[string]any, 0, len(res.Hits))
	for _, h := range res.Hits {
		srcs = append(srcs, map[string]any{"title": h.Title, "url": h.URL})
	}
	return map[string]any{"sources": srcs, "search_provider": res.Provider}
}

func formatFetch(title, md string) string {
	if title != "" {
		return truncate(title+"\n\n"+md, toolMaxOutput)
	}
	return truncate(md, toolMaxOutput)
}

func (e *Engine) webContext(ctx context.Context, user, query string) string {
	wt := e.webTools()
	if wt == nil || strings.TrimSpace(query) == "" || !e.allowWeb(user) {
		return ""
	}
	res := wt.Search(ctx, query, 5)
	if len(res.Hits) == 0 {
		return ""
	}
	lines := make([]string, 0, len(res.Hits))
	for i, h := range res.Hits {
		lines = append(lines, fmt.Sprintf("[%d] %s\n%s\n%s", i+1, h.Title, h.URL, h.Description))
	}
	return "Web search results for the query: \"" + query + "\"\n\n" +
		strings.Join(lines, "\n\n") +
		"\n\nUse this information to answer and cite your sources by their number [1], [2], etc."
}
