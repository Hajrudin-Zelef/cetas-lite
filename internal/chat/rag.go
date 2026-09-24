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
	// SearchCorpus : recherche restreinte aux chunks d'un corpus
	// (focus au clic sur une question suggérée).
	SearchCorpus(ctx context.Context, query, corpus string, limit int) rag.Result
	// SearchBoosted : recherche avec facteur de boost par terme (entités
	// reprises de l'historique, itération 6b).
	SearchBoosted(ctx context.Context, query string, boost map[string]float64, limit int) rag.Result
	Read(rel string, offset, limit int) (string, error)
	Stats() rag.Stats
}

const (
	ragContextHits   = 3
	ragContextBudget = 3600
	// entityBoostFactor : boost appliqué aux tokens d'entité ajoutés par
	// l'enrichissement (itération 6b). Le corpus réel (1819 chunks) :
	// sans boost, l'entité à idf faible perdait face aux termes
	// conversationnels rares (RTX Spark 12.99 > Kimi 8.86, aucun Kimi
	// dans le top-10). Calibré à ×3 sur ce corpus ; mesurer avant de
	// bouger.
	entityBoostFactor = 3.0
)

// RagToolSchemas : outils rag_search / rag_read exposes a l'agent.
func RagToolSchemas() []provider.Tool {
	str := func(props map[string]any, required ...string) map[string]any {
		return map[string]any{"type": "object", "properties": props, "required": required}
	}
	return []provider.Tool{
		{Type: "function", Function: provider.ToolFunction{
			Name: "rag_search",
			Description: "Search the local document base (tech/AI news, personal notes). " +
				"Returns the most relevant passages with their paths. Use before any web search.",
			Parameters: str(map[string]any{
				"query": map[string]any{"type": "string"},
				"limit": map[string]any{"type": "integer", "description": "Default 8, max 30"},
			}, "query")}},
		{Type: "function", Function: provider.ToolFunction{
			Name:        "rag_read",
			Description: "Read a passage from the local base by its path (from rag_search), numbered lines. Server limit: default 80 lines, max 200 lines.",
			Parameters: str(map[string]any{
				"path":   map[string]any{"type": "string"},
				"offset": map[string]any{"type": "integer", "description": "First line (1-based), default 1"},
				"limit":  map[string]any{"type": "integer", "description": "Default 80 lines, max 200"},
			}, "path")}},
	}
}

func (e *Engine) ragExecute(ctx context.Context, name, argsJSON string) ToolResult {
	rt := e.ragTools()
	if rt == nil || !rt.Ready() {
		return ToolResult{Text: "[info] local document base unavailable"}
	}
	args := parseArgs(argsJSON)
	switch name {
	case "rag_search":
		query := strings.TrimSpace(strArg(args, "query"))
		if query == "" {
			return ToolResult{Text: "[error] empty query"}
		}
		res := rt.Search(ctx, query, intArg(args, "limit"))
		if len(res.Hits) == 0 {
			return ToolResult{Text: "[info] no relevant passage in the local base"}
		}
		return ToolResult{Text: formatRagHits(res)}
	case "rag_read":
		rel := strings.TrimSpace(strArg(args, "path"))
		if rel == "" {
			return ToolResult{Text: "[error] empty path"}
		}
		// Bornage serveur : limite par defaut raisonnable et plafond strict
		// pour eviter qu'une lecture mal bornee n'injecte un document entier
		// (jusqu'a ~12k tokens) dans le contexte.
		offset := intArg(args, "offset")
		if offset < 0 {
			offset = 0
		}
		limit := intArg(args, "limit")
		if limit <= 0 {
			limit = ragReadDefaultLines
		} else if limit > ragReadMaxLines {
			limit = ragReadMaxLines
		}
		out, err := rt.Read(rel, offset, limit)
		if err != nil {
			return ToolResult{Text: "[error] read: " + err.Error()}
		}
		return ToolResult{Text: truncate(out, toolMaxOutput)}
	}
	return ToolResult{Text: "[error] unknown tool: " + name}
}

func formatRagHits(res rag.Result) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Local base (%d relevant passage(s)):\n", len(res.Hits))
	for i, h := range res.Hits {
		fmt.Fprintf(&b, "[%d] %s — %s\n%s\n\n", i+1, h.Title, h.Path, h.Snippet)
	}
	return truncate(b.String(), toolMaxOutput)
}

// ragHits interroge l'index local. La requete est enrichie avec le sujet
// de la conversation quand elle est elliptique (iteration 6 : anaphore
// « la » resolue au lieu d'etre jetee comme mot-outil). Resultat vide si
// le RAG est inactif (fail-open : l'appelant traite le vide comme
// "pas de contexte").
func (e *Engine) ragHits(ctx context.Context, query string, history []provider.Message) rag.Result {
	rt := e.ragTools()
	if rt == nil || !rt.Ready() || strings.TrimSpace(query) == "" {
		return rag.Result{}
	}
	var hist []string
	for _, m := range history {
		if s, ok := m.Content.(string); ok && strings.TrimSpace(s) != "" {
			hist = append(hist, s)
		}
	}
	q, ents := rag.EnrichQueryWithHistory(query, hist)
	if len(ents) > 0 {
		return rt.SearchBoosted(ctx, q, boostMap(ents, entityBoostFactor), ragContextHits)
	}
	return rt.Search(ctx, q, ragContextHits)
}

// boostMap : facteur de boost par token d'entité (itération 6b).
func boostMap(entities []string, factor float64) map[string]float64 {
	m := make(map[string]float64, len(entities))
	for _, t := range entities {
		m[t] = factor
	}
	return m
}

// ragHitsCorpus interroge l'index local restreint à un corpus.
// Résultat vide si le RAG est inactif ou le corpus inconnu (fail-open).
func (e *Engine) ragHitsCorpus(ctx context.Context, query, corpus string) rag.Result {
	rt := e.ragTools()
	if rt == nil || !rt.Ready() || strings.TrimSpace(query) == "" || strings.TrimSpace(corpus) == "" {
		return rag.Result{}
	}
	return rt.SearchCorpus(ctx, query, corpus, ragContextHits)
}

// ragContextFrom construit le message systeme depuis un resultat deja obtenu.
// Porte de score (iteration 1) : des hits faibles ne doivent pas etre
// presentes comme « source prioritaire ». Sans couverture, rien n'est
// injecte (fail-open) et la recherche web prend le relais si activee.
func ragContextFrom(res rag.Result) (string, bool) {
	if !ragCovered(res) {
		return "", false
	}
	return buildRagContext(res), true
}

// ragContextForced : comme ragContextFrom mais sans porte de score.
// Réservé au focus corpus explicite (clic sur une question suggérée) :
// la question est curée pour ce corpus, le RAG est forcément sollicité.
// Fail-open : aucun hit => rien d'injecté.
func ragContextForced(res rag.Result) (string, bool) {
	if len(res.Hits) == 0 {
		return "", false
	}
	return buildRagContext(res), true
}

func buildRagContext(res rag.Result) string {
	var b strings.Builder
	b.WriteString("The following is your document knowledge base: treat it as things you know, not as documents handed to you.\n" +
		"Response guidelines (must follow):\n" +
		"- Answer the question directly, in your own words, in continuous natural prose, like in a conversation.\n" +
		"- Never mention the excerpts, the documents, or the base: \"according to the excerpts\", \"the provided documents\", \"the base does (not) cover\" are banned. No visible citations ([1], [2]…), no internal paths, unless the user explicitly asks where the information comes from.\n" +
		"- Never describe the content or limits of the sources; never send the user back to reading (\"tell me if I should read it\", \"go look for yourself\" are banned).\n" +
		"- Weave the information together instead of juxtaposing it: no section per excerpt, no per-source subtitles, no mechanical enumeration.\n" +
		"- Structure your answer: main idea first, then key points, then limits or uncertainties — without turning it into a formal report.\n" +
		"- These extracts are your primary source and factual anchor: ground your answer in them, and never present internal knowledge or guesses as coming from the local document base. Reason across the extracts, connect them, and synthesize in your own words — do not write one section per extract. Where your own knowledge genuinely helps beyond what the extracts cover, you may use it to reason further, staying on the topic of the question.\n" +
		"- If a requested point is not there, say you don't know rather than inventing it — without presenting invention as coming from the base, and without verbalizing this rule.\n" +
		"- Only search the web if the information is truly missing above.\n" +
		"- Only use rag_read if the passages above are insufficient, and in small passages (offset/limit) rather than the whole document.\n\n")
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
	return b.String()
}

// Bornes serveur de rag_read : un appel sans limit (ou mal bornee) ne doit
// pas injecter un document entier dans le contexte. Le plafond reste bien en
// dessous de maxReadLines (400) pour eviter les lectures a ~12k tokens.
const (
	ragReadDefaultLines = 80
	ragReadMaxLines     = 200
)

// ragStrongScore : au-dela de ce score BM25, la base locale est consideree
// comme couvrant la requete. Sert a economiser la pre-recherche web du tour.
const ragStrongScore = 2.5

// insertBeforeLastUser insere des messages dynamiques juste avant le dernier
// message utilisateur (le tour courant), donc apres l'historique. Le
// prefixe [prompts systeme stables + historique] reste ainsi identique d'un
// tour a l'autre, ce qui favorise le prompt caching (prefixe automatique
// cote provider). Si le dernier message n'est pas un message utilisateur,
// les messages dynamiques sont ajoutes a la fin.
func insertBeforeLastUser(msgs []provider.Message, dyn ...provider.Message) []provider.Message {
	if len(dyn) == 0 {
		return msgs
	}
	if n := len(msgs); n > 0 && msgs[n-1].Role == "user" {
		out := make([]provider.Message, 0, n+len(dyn))
		out = append(out, msgs[:n-1]...)
		out = append(out, dyn...)
		out = append(out, msgs[n-1])
		return out
	}
	return append(msgs, dyn...)
}

// ragCovered : la base locale repond a la requete.
func ragCovered(res rag.Result) bool {
	return len(res.Hits) > 0 && res.Hits[0].Score >= ragStrongScore
}

// localFirstDirective remplace la directive de recherche web quand la base
// locale couvre deja la requete : evite une recherche inutile (cout + latence).
func localFirstDirective() string {
	return "A local document base already provides background knowledge for this request " +
		"(see the local-base system message) — treat it as things you know, not as documents handed to you. " +
		"Answer directly in natural prose, weaving the facts together, without visible citations and without ever mentioning " +
		"the extracts, the documents or the base. Only search the web if that knowledge clearly does not contain the answer."
}

// ragNotCoveredNote : note systeme injectee quand la base locale est active
// mais ne couvre pas la requete. Texte strictement stable (pas de contenu
// dynamique) pour ne pas casser le préfixe de prompt caching. Dit
// immediatement la non-couverture et interdit de presenter une invention
// comme issue du corpus — en modelisant le ton naturel attendu (iteration 6 :
// pas de rapport formel ni d'enumeration de sources).
func ragNotCoveredNote() string {
	return "The local document base does not cover this request: do not mention it " +
		"and do not present anything as coming from it. If you answer, do so simply and " +
		"naturally, without enumerating sources or writing a formal report."
}
