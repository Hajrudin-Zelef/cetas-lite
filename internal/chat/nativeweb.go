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
// utilisee.
//   - OpenRouter : plugin "web" (recherche integree, ~0,007 $/requete en
//     moteur Exa auto, ou tarif du provider en moteur natif).
//   - DeepSeek : parametre "enable_search" (recherche integree cote DeepSeek).
//
// Dans les deux cas le modele declenche la recherche lui-meme quand il
// doute, atteint son cutoff, ne connait pas la reponse, que l'information
// est recente, ou que l'utilisateur la demande explicitement.
func (e *Engine) nativeWebFor(user, providerID, modelID string) bool {
	switch providerID {
	case "openrouter":
		if strings.HasSuffix(strings.ToLower(modelID), ":online") {
			return false // modele deja "en ligne" : rien a injecter
		}
	case "deepseek":
	default:
		return false
	}
	switch e.webSearchMode(user) {
	case WebSearchAuto, WebSearchNative:
		return true
	}
	return false
}

// nativeWebExtraFor construit les parametres natifs de recherche web du
// provider (injectes dans "Extra" de la requete).
func nativeWebExtraFor(providerID string) map[string]any {
	if providerID == "deepseek" {
		return map[string]any{"enable_search": true}
	}
	return map[string]any{
		"plugins": []any{map[string]any{"id": "web", "max_results": 5}},
	}
}

// webEnabled : le globe (in.Web) est l'interrupteur principal de la
// recherche web ; le mode "off" la coupe egalement.
func (e *Engine) webEnabled(in TurnInput) bool {
	return in.Web && e.webSearchMode(in.User) != WebSearchOff
}

// searchDirective retourne la directive de recherche web du tour, en anglais.
// Imperative : le modele doit respecter l'etat du globe.
//
//	web=false        : aucune recherche web ce tour-ci.
//	web=true+natif   : recherche native du provider (declenchement auto).
//	web=true+outils  : recherche forcee via web_search/web_fetch.
func searchDirective(web, native bool) string {
	if !web {
		return "Web search is disabled for this turn. Answer from your own knowledge. " +
			"If you are unsure or the information might be outdated, say so explicitly instead of inventing facts."
	}
	if native {
		return "Web search is enabled (provider-native search). Use it whenever you are unsure, " +
			"the information may be recent or beyond your knowledge cutoff, you do not know the answer, " +
			"or the user explicitly asks for a web search. Cite your sources when you use search results."
	}
	return "Web search is enabled for this turn. Use the web_search and web_fetch tools whenever the answer " +
		"may depend on recent, external, or uncertain facts: when you are unsure, when you hit your knowledge " +
		"cutoff, when you do not know the answer, or when the user explicitly asks for a web search. " +
		"You MUST search the web rather than guessing in those cases. Cite your sources."
}

// deepWebDirective ordonne une recherche approfondie : l'agent ne se
// contente pas des extraits, il multiplie les angles et lit les sources
// en entier. Injectee uniquement quand WebDepth == "deep".
func deepWebDirective() string {
	return "DEEP WEB RESEARCH for this turn: do not stop at snippets. " +
		"Run several web_search calls with varied query angles (set max_results up to 10), " +
		"then web_fetch the most relevant pages and read their full content. " +
		"Cross-check important facts across at least two independent sources before answering, " +
		"and cite every source you relied on."
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

// replaceWebDirective remplace l'ancienne directive de recherche web par la
// nouvelle dans les messages systeme (fusionnes ou non). Utilise lors du
// repli natif -> outils : la directive ne doit plus parler de recherche
// native une fois le plugin desactive.
func replaceWebDirective(msgs []provider.Message, old, new string) []provider.Message {
	if old == "" || old == new {
		return msgs
	}
	for i, m := range msgs {
		if m.Role != "system" {
			continue
		}
		if s, ok := m.Content.(string); ok && strings.Contains(s, old) {
			msgs[i].Content = strings.Replace(s, old, new, 1)
			break
		}
	}
	return msgs
}

// webToolsFor indique si les outils web_search/web_fetch sont proposes au
// modele pour ce tour (chemin de repli).
func (e *Engine) webToolsFor(in TurnInput) bool {
	if !e.webEnabled(in) || in.User == "" || e.webTools() == nil {
		return false
	}
	// Recherche approfondie : le natif est desactive pour ce tour, les
	// outils sont donc requis quel que soit le mode global (sauf web
	// coupe, deja exclu par webEnabled).
	if in.WebDepth == "deep" {
		return true
	}
	if e.webSearchMode(in.User) == WebSearchNative {
		return false
	}
	return true
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
