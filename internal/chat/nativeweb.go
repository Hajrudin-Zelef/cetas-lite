package chat

import (
	"encoding/json"
	"strings"

	"cetas-lite/internal/provider"
)

// Modes de recherche web (reglage utilisateur "websearch_mode").
//   - auto   : recherche native du provider quand elle existe (OpenRouter),
//     sinon outils web_search/web_fetch (defaut).
//   - natif  : recherche native uniquement (OpenRouter) ; aucun repli.
//   - outils : toujours les outils web_search/web_fetch, jamais le natif.
//   - off    : aucune recherche web.
const (
	WebSearchAuto   = "auto"
	WebSearchNative = "natif"
	WebSearchTools  = "outils"
	WebSearchOff    = "off"
)

// webSearchMode lit le reglage de recherche web de l'utilisateur
// (defaut "auto").
func (e *Engine) webSearchMode(user string) string {
	if e.st == nil || user == "" {
		return WebSearchAuto
	}
	raw, ok := e.st.GetSetting(user, "ui")
	if !ok {
		return WebSearchAuto
	}
	var ui struct {
		WebSearchMode string `json:"websearch_mode"`
	}
	if err := json.Unmarshal(raw, &ui); err != nil {
		return WebSearchAuto
	}
	switch ui.WebSearchMode {
	case WebSearchNative, WebSearchTools, WebSearchOff:
		return ui.WebSearchMode
	}
	return WebSearchAuto
}

// nativeWebFor indique si la recherche web native du provider doit etre
// utilisee. Aujourd'hui seul OpenRouter expose une recherche native (plugin
// "web" : ~0,007 $/requete en moteur Exa auto, ou tarif du provider en
// moteur natif). DeepSeek n'offre pas de recherche web native via son API :
// le repli par outils reste le seul chemin honnete.
func (e *Engine) nativeWebFor(user, providerID, modelID string) bool {
	if providerID != "openrouter" {
		return false
	}
	if strings.HasSuffix(strings.ToLower(modelID), ":online") {
		return false // modele deja "en ligne" : rien a injecter
	}
	switch e.webSearchMode(user) {
	case WebSearchAuto, WebSearchNative:
		return true
	}
	return false
}

// useNativeWebSearch decide, pour un tour donne, si la recherche web passe
// par le natif du provider (avec consommation du limiteur de debit partage).
func (e *Engine) useNativeWebSearch(in TurnInput, providerID, modelID string) bool {
	if !in.Web || in.User == "" {
		return false
	}
	if !e.nativeWebFor(in.User, providerID, modelID) {
		return false
	}
	return e.allowWeb(in.User)
}

// webToolsFor indique si les outils web_search/web_fetch sont proposes au
// modele pour ce tour (chemin de repli).
func (e *Engine) webToolsFor(in TurnInput) bool {
	if !in.Web || in.User == "" || e.webTools() == nil {
		return false
	}
	switch e.webSearchMode(in.User) {
	case WebSearchOff, WebSearchNative:
		return false
	}
	return true
}

// nativeWebExtra construit le parametre "plugins" d'OpenRouter qui active
// la recherche web native cote provider.
func nativeWebExtra() map[string]any {
	return map[string]any{
		"plugins": []any{map[string]any{"id": "web", "max_results": 5}},
	}
}

// dropWebTools retire web_search/web_fetch d'une liste d'outils (quand le
// natif du provider prend le relais, pour eviter une double recherche).
func dropWebTools(tools []provider.Tool) []provider.Tool {
	out := make([]provider.Tool, 0, len(tools))
	for _, t := range tools {
		name := t.Function.Name
		if name == "web_search" || name == "web_fetch" {
			continue
		}
		out = append(out, t)
	}
	return out
}

// searchSourcesDelta convertit des annotations natives en sources pour le
// delta SSE {"search": {...}} consomme par le frontend.
func searchSourcesDelta(annotations []provider.WebAnnotation, native bool) map[string]any {
	srcs := make([]map[string]any, 0, len(annotations))
	for _, a := range annotations {
		srcs = append(srcs, map[string]any{"title": a.Title, "url": a.URL})
	}
	return map[string]any{
		"phase":   "done",
		"native":  native,
		"sources": srcs,
	}
}
