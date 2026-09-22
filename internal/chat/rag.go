package chat

import (
	"context"
	"fmt"
	"strings"

	"cetas-lite/internal/provider"
	"cetas-lite/internal/rag"
)

// RagTools : acces a l'index documentaire local (lecture seule, in-process).
// *rag.Manager l'implemente ; les tests injectent un faux.
type RagTools interface {
	Ready() bool
	Search(ctx context.Context, query string, limit int) rag.Result
	Read(rel string, offset, limit int) (string, error)
	Stats() rag.Stats
}

const (
	ragContextHits   = 3
	ragContextBudget = 3600
)

// RagToolSchemas : outils rag_search / rag_read exposes a l'agent.
func RagToolSchemas() []provider.Tool {
	str := func(props map[string]any, required ...string) map[string]any {
		return map[string]any{"type": "object", "properties": props, "required": required}
	}
	return []provider.Tool{
		{Type: "function", Function: provider.ToolFunction{
			Name: "rag_search",
			Description: "Cherche dans la base documentaire locale (actualites tech/IA, notes perso). " +
				"Renvoie les passages les plus pertinents avec leur chemin. A utiliser avant toute recherche web.",
			Parameters: str(map[string]any{
				"query": map[string]any{"type": "string"},
				"limit": map[string]any{"type": "integer", "description": "Default 8, max 30"},
			}, "query")}},
		{Type: "function", Function: provider.ToolFunction{
			Name:        "rag_read",
			Description: "Relit un passage de la base locale par son chemin (issu de rag_search), lignes numerotees.",
			Parameters: str(map[string]any{
				"path":   map[string]any{"type": "string"},
				"offset": map[string]any{"type": "integer", "description": "First line (1-based)"},
				"limit":  map[string]any{"type": "integer"},
			}, "path")}},
	}
}

func (e *Engine) ragExecute(ctx context.Context, name, argsJSON string) ToolResult {
	rt := e.ragTools()
	if rt == nil || !rt.Ready() {
		return ToolResult{Text: "[info] base documentaire locale indisponible"}
	}
	args := parseArgs(argsJSON)
	switch name {
	case "rag_search":
		query := strings.TrimSpace(strArg(args, "query"))
		if query == "" {
			return ToolResult{Text: "[erreur] requete vide"}
		}
		res := rt.Search(ctx, query, intArg(args, "limit"))
		if len(res.Hits) == 0 {
			return ToolResult{Text: "[info] aucun passage pertinent dans la base locale"}
		}
		return ToolResult{Text: formatRagHits(res)}
	case "rag_read":
		rel := strings.TrimSpace(strArg(args, "path"))
		if rel == "" {
			return ToolResult{Text: "[erreur] chemin vide"}
		}
		out, err := rt.Read(rel, intArg(args, "offset"), intArg(args, "limit"))
		if err != nil {
			return ToolResult{Text: "[erreur] lecture: " + err.Error()}
		}
		return ToolResult{Text: truncate(out, toolMaxOutput)}
	}
	return ToolResult{Text: "[erreur] outil inconnu: " + name}
}

func formatRagHits(res rag.Result) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Base locale (%d passage(s) pertinent(s)):\n", len(res.Hits))
	for i, h := range res.Hits {
		fmt.Fprintf(&b, "[%d] %s — %s\n%s\n\n", i+1, h.Title, h.Path, h.Snippet)
	}
	return truncate(b.String(), toolMaxOutput)
}

// ragHits interroge l'index local. Resultat vide si le RAG est inactif
// (fail-open : l'appelant traite le vide comme "pas de contexte").
func (e *Engine) ragHits(ctx context.Context, query string) rag.Result {
	rt := e.ragTools()
	if rt == nil || !rt.Ready() || strings.TrimSpace(query) == "" {
		return rag.Result{}
	}
	return rt.Search(ctx, query, ragContextHits)
}

// ragContextFrom construit le message systeme depuis un resultat deja obtenu.
func ragContextFrom(res rag.Result) (string, bool) {
	if len(res.Hits) == 0 {
		return "", false
	}
	var b strings.Builder
	b.WriteString("Extraits de la base documentaire locale — source prioritaire pour les sujets couverts. " +
		"Reponds depuis ces extraits et cite [1], [2]… Ne fais une recherche web que si l'information manque. " +
		"Utilise rag_read pour lire un passage entier.\n\n")
	budget := ragContextBudget
	for i, h := range res.Hits {
		ex := strings.TrimSpace(h.Excerpt)
		if ex == "" {
			ex = h.Snippet
		}
		entry := fmt.Sprintf("[%d] %s (%s)\n%s\n\n", i+1, h.Title, h.Path, ex)
		if len(entry) > budget {
			break
		}
		b.WriteString(entry)
		budget -= len(entry)
	}
	return b.String(), true
}

// ragStrongScore : au-dela de ce score BM25, la base locale est consideree
// comme couvrant la requete. Sert a economiser la pre-recherche web du tour.
const ragStrongScore = 2.5

// ragCovered : la base locale repond a la requete.
func ragCovered(res rag.Result) bool {
	return len(res.Hits) > 0 && res.Hits[0].Score >= ragStrongScore
}

// localFirstDirective remplace la directive de recherche web quand la base
// locale couvre deja la requete : evite une recherche inutile (cout + latence).
func localFirstDirective() string {
	return "A local document base already provides relevant extracts for this request " +
		"(see the local-base system message). Answer from those extracts first and cite them [1], [2]… " +
		"Only search the web if the local extracts clearly do not contain the answer."
}
