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
//   - conversion restreinte aux outils sans effet de bord (parallelSafe) :
//     une erreur d'interpretation ne peut au pire declencher qu'une
//     lecture. Les autres outils (Write, Bash, ...) passent par le nudge ;
//   - l'unique argument positionnel est mappe sur le parametre "string"
//     evident du schema (l'unique requis, sinon l'unique existant) :
//     jamais de devinette sur un schema ambigu.

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
	// Conversion restreinte aux lectures pures : une meprise ne doit
	// jamais declencher une ecriture ou une execution.
	if !parallelSafe(tool.Function.Name) {
		return nil
	}
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

// toolAttemptPhrases : formulations (FR/EN) par lesquelles un modele annonce
// un appel d'outil qu'il n'emet finalement pas en function call.
var toolAttemptPhrases = []string{
	"let me call", "i'll call", "i will call", "i'm calling",
	"need to call", "j'appelle l'outil",
	"je vais appeler", "je dois appeler",
}

// looksLikeToolAttempt detecte une reponse qui ressemble a une tentative
// d'appel d'outil non emise en function call. Retourne le nom canonique de
// l'outil vise, ou "" si la tentative est generique. Second retour : vrai
// si une tentative est detectee.
//
// La detection reste volontairement conservative : une ligne pseudo-appel
// isolee dans une reponse courte, ou une formulation d'intention explicite.
// Les longues explications contenant un exemple de code ne declenchent rien.
func looksLikeToolAttempt(content string, tools []provider.Tool) (string, bool) {
	lower := strings.ToLower(content)
	for _, p := range toolAttemptPhrases {
		if strings.Contains(lower, p) {
			return toolNameInText(content, tools), true
		}
	}
	s := stripCodeFences(content)
	if runeCount(s) > 300 {
		return "", false
	}
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		m := textCallRe.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		for i := range tools {
			if strings.EqualFold(tools[i].Function.Name, m[1]) {
				return tools[i].Function.Name, true
			}
		}
	}
	return "", false
}

// toolNameInText cherche un nom d'outil connu dans le texte (pour cibler
// le nudge). Chaine vide si aucun.
func toolNameInText(content string, tools []provider.Tool) string {
	lower := strings.ToLower(content)
	for i := range tools {
		if strings.Contains(lower, strings.ToLower(tools[i].Function.Name)) {
			return tools[i].Function.Name
		}
	}
	return ""
}

func runeCount(s string) int { return len([]rune(s)) }

// toolAttemptNudgeText construit le nudge envoye quand le modele a decrit
// un appel d'outil en texte au lieu de l'executer (inspire du pattern
// d'opencode : retour d'erreur explicite pour auto-correction).
func toolAttemptNudgeText(toolName string) string {
	target := "the right tool"
	if toolName != "" {
		target = "the " + toolName + " tool"
	}
	return "You just described a tool call in plain text instead of executing it — " +
		"writing it as text does NOTHING. In your NEXT message, call " + target +
		" NOW with a real function call (no confirmation question, no text version). " +
		"Text is for communicating with the user; tools are for acting."
}
