package chat

import (
	"encoding/json"
	"regexp"
	"strings"

	"cetas-lite/internal/provider"
)

// Parser de pseudo-appels d'outils ecrits en texte (etape 1b).
//
// Certains modeles (notamment des variantes "flash" aux function-calls
// natifs peu fiables) ecrivent l'appel au lieu de l'emettre, par exemple :
//
//	Cat "NEVA PVE/dnsmasq.conf"
//
// au lieu d'un vrai function call. Sans traitement, ce texte s'affiche
// tel quel et aucun outil ne tourne. Ce parser convertit ces pseudo-appels
// en vrais provider.ToolCall, executes par le pipeline standard
// (approbation, blocs d'outils, suivi des modifications).
//
// Garde-fous (anti faux positifs) :
//   - conversion uniquement si TOUTE la reponse (hors cloture markdown)
//     n'est qu'un seul pseudo-appel : jamais d'execution d'un exemple
//     noye dans du texte, d'une citation ou d'une explication ;
//   - le nom doit correspondre (insensible a la casse) a un outil
//     reellement annonce dans ce tour : jamais d'outil invente ;
//   - les outils a effets de bord (Bash, Mkdir...) sont convertis eux
//     aussi : ils passent par le pipeline standard et son approbation
//     utilisateur avant toute execution — une erreur d'interpretation
//     declenche au pire une carte d'approbation, jamais une execution
//     silencieuse ;
//   - l'unique argument positionnel est mappe sur le parametre "string"
//     evident du schema (l'unique requis, sinon l'unique existant) :
//     jamais de devinette sur un schema ambigu (Edit/Write, a plusieurs
//     parametres, ne sont pas convertibles par cette voie).

// textCallRe reconnait un pseudo-appel : Nom, optionnellement suivi d'un
// unique argument entre guillemets (ou nu), avec ou sans parentheses.
// Exemples : Cat "a/b", read('x'), Ls, Grep("motif").
var textCallRe = regexp.MustCompile(`^([A-Za-z][A-Za-z0-9_]{1,40})\s*(?:\(\s*)?["']?([^"'()\r\n]*?)["']?\s*\)?$`)

// stripCodeFences retire une cloture markdown entourant tout le contenu.
func stripCodeFences(s string) string {
	t := strings.TrimSpace(s)
	if strings.HasPrefix(t, "```") {
		if i := strings.Index(t, "\n"); i >= 0 {
			t = t[i+1:]
		}
		if j := strings.LastIndex(t, "```"); j >= 0 {
			t = t[:j]
		}
	}
	return strings.TrimSpace(t)
}

// textCallParam retourne le nom du parametre "string" evident du schema
// de l'outil : l'unique parametre requis de type string, ou a defaut
// l'unique parametre de type string. Chaine vide si ambigu.
func textCallParam(t provider.Tool) string {
	params, ok := t.Function.Parameters.(map[string]any)
	if !ok {
		if s, ok := t.Function.Parameters.(string); ok {
			var m map[string]any
			if json.Unmarshal([]byte(s), &m) == nil {
				params = m
			}
		}
	}
	if params == nil {
		return ""
	}
	props, _ := params["properties"].(map[string]any)
	if len(props) == 0 {
		return ""
	}
	isStr := func(name string) bool {
		p, _ := props[name].(map[string]any)
		typ, _ := p["type"].(string)
		return typ == "string"
	}
	var required []string
	switch r := params["required"].(type) {
	case []string:
		required = r
	case []any:
		for _, v := range r {
			if s, ok := v.(string); ok {
				required = append(required, s)
			}
		}
	}
	var reqStr []string
	for _, n := range required {
		if isStr(n) {
			reqStr = append(reqStr, n)
		}
	}
	if len(reqStr) == 1 {
		return reqStr[0]
	}
	var allStr []string
	for n := range props {
		if isStr(n) {
			allStr = append(allStr, n)
		}
	}
	if len(allStr) == 1 {
		return allStr[0]
	}
	return ""
}

// parseTextToolCall convertit un pseudo-appel ecrit en texte en vrai
// ToolCall, ou retourne nil si le contenu n'est pas un appel isole et
// convertible sans ambiguite.
func parseTextToolCall(content string, tools []provider.Tool) *provider.ToolCall {
	s := stripCodeFences(content)
	if s == "" || strings.Contains(s, "\n") {
		return nil
	}
	m := textCallRe.FindStringSubmatch(s)
	if m == nil {
		return nil
	}
	name, arg := m[1], strings.TrimSpace(m[2])
	// Phase 3 : un pseudo-appel vers un outil fusionné (ex. Cat "f")
	// est réécrit vers l'outil canonique (Read) avant le matching.
	name = canonicalToolName(name)
	var tool *provider.Tool
	for i := range tools {
		if strings.EqualFold(tools[i].Function.Name, name) {
			tool = &tools[i]
			break
		}
	}
	if tool == nil {
		return nil
	}
	// Les outils a effets de bord sont convertis eux aussi : le pipeline
	// standard exige l'approbation utilisateur avant execution
	// (needsApprovalFor). Une erreur d'interpretation declenche au pire
	// une carte d'approbation, jamais une execution silencieuse.
	args := map[string]any{}
	if arg != "" {
		param := textCallParam(*tool)
		if param == "" {
			return nil
		}
		args[param] = arg
	}
	argsJSON, _ := json.Marshal(args)
	return &provider.ToolCall{
		ID:   "text-" + newID(),
		Type: "function",
		Function: provider.Func{
			Name:      tool.Function.Name,
			Arguments: string(argsJSON),
		},
	}
}
