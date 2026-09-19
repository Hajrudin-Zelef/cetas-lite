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
//     n'est que pseudo-appel(s), un par ligne : jamais d'execution d'un
//     exemple noye dans du texte, d'une citation ou d'une explication ;
//     si une seule ligne echoue, rien n'est converti (pas d'execution
//     partielle) ;
//   - un label d'annonce en tete de ligne ("Appel réel :", "Real call:",
//     ...) est tolere, mais seulement s'il porte un mot-cle d'annonce
//     (appel/call/outil/tool/...) : un label neutre ("Voici le
//     resultat :") ne convertit jamais ;
//   - le nom doit correspondre (insensible a la casse, alias resolus) a
//     un outil reellement annonce dans ce tour : jamais d'outil invente ;
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

// textCallLabelRe reconnait un label d'annonce d'appel en tete de ligne :
// un prefixe court (<=48 caracteres, sans ":" ni saut de ligne) termine par
// ":". Exemples observes : "Appel réel : Read ...", "Real call: Ls ...".
//
// Le label n'est PAS une whitelist figee : chaque modele formule
// differemment ("Calling tool:", "J'appelle :", "Executing...", ...).
// La securite vient de la conjonction : label porteur d'un mot-cle
// d'annonce (callAnnounceRe) + pseudo-appel strict + nom d'outil
// reellement annonce dans ce tour. Un label neutre ("Voici le
// resultat :", "Note :") est refuse : c'est ce qui evite de convertir
// une phrase anodine contenant un mot-outil ("Read the docs").
var textCallLabelRe = regexp.MustCompile(`^([^:\n]{1,48}):\s*(.+)$`)

// callAnnounceRe : le label doit sentir l'annonce d'appel (insensible a la
// casse, accents toleres). Couvre "Appel réel :", "Real call:",
// "Calling tool:", "J'appelle :", "Exécution :". Le mot-cle ne doit pas
// etre noye dans un autre mot : "Rappel :" ou "recall:" sont refuses
// (borne : debut de chaine ou caractere non-lettre devant le mot-cle).
// Un label neutre ("Voici le résultat :", "Note :") est refuse : c'est ce
// qui evite de convertir une phrase anodine contenant un mot-outil
// ("Read the docs").
var callAnnounceRe = regexp.MustCompile(`(?i)(?:^|[^a-z])(?:appel|call|outil|tool|ex[ée]cut|invo[kq]u|lan[çc])`)

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
// convertible sans ambiguite. Le contenu multi-lignes est refuse ici
// (voir parseTextToolCalls) ; un label d'annonce ("Appel réel :",
// "Real call:", ...) est tolere devant l'appel.
func parseTextToolCall(content string, tools []provider.Tool) *provider.ToolCall {
	s := stripCodeFences(content)
	if s == "" || strings.Contains(s, "\n") {
		return nil
	}
	return parseTextToolCallLine(s, tools)
}

// parseTextToolCalls convertit un texte entierement constitue de
// pseudo-appels (un par ligne, avec ou sans label d'annonce) en vrais
// ToolCalls — ex. le cas observe ou le modele emet "Appel réel : Ls ..."
// puis "Appel réel : Glob ..." sur deux lignes. TOUTES les lignes non
// vides doivent etre des appels convertibles, sinon nil : le texte reste
// affiche tel quel, jamais d'execution partielle.
func parseTextToolCalls(content string, tools []provider.Tool) []provider.ToolCall {
	s := stripCodeFences(content)
	if s == "" {
		return nil
	}
	var out []provider.ToolCall
	for _, line := range strings.Split(s, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		tc := parseTextToolCallLine(line, tools)
		if tc == nil {
			return nil
		}
		out = append(out, *tc)
	}
	return out
}

// parseTextToolCallLine convertit UNE ligne (pseudo-appel isole, avec ou
// sans label d'annonce) en vrai ToolCall, ou retourne nil.
func parseTextToolCallLine(line string, tools []provider.Tool) *provider.ToolCall {
	s := strings.TrimSpace(line)
	if s == "" {
		return nil
	}
	// 1) Forme stricte historique : pas de label.
	if tc := matchTextCall(s, tools); tc != nil {
		return tc
	}
	// 2) Label d'annonce tolere : "<label>: <pseudo-appel>".
	// Le label doit porter un mot-cle d'annonce ; le reste doit etre un
	// pseudo-appel strict vers un outil connu.
	if m := textCallLabelRe.FindStringSubmatch(s); m != nil && callAnnounceRe.MatchString(m[1]) {
		return matchTextCall(strings.TrimSpace(m[2]), tools)
	}
	return nil
}

// looksLikeAnnouncedCall detecte une tentative d'appel d'outil annoncee en
// texte mais non convertible par le parsing strict : un nom d'outil connu
// (apres resolution des alias : "Cat" -> "Read") apparait dans une ligne
// portant un mot-cle d'annonce d'appel. Utilise par le nudge borne
// (agent.go) et le log de diagnostic (fallbackcall.go).
//
// Ne se declenche jamais sur une reponse texte normale : il faut a la fois
// un mot-cle d'annonce ET un nom d'outil reconnu. "Read the docs" seul ne
// matche pas (pas d'annonce) ; "Voici le resultat : ..." ne matche pas
// (pas de mot-cle d'annonce dans le label).
func looksLikeAnnouncedCall(content string, tools []provider.Tool) bool {
	s := stripCodeFences(content)
	if s == "" {
		return false
	}
	known := map[string]bool{}
	for _, t := range tools {
		if t.Function.Name != "" {
			known[strings.ToLower(t.Function.Name)] = true
		}
	}
	if len(known) == 0 {
		return false
	}
	mentionsTool := func(line string) bool {
		for _, w := range strings.Fields(line) {
			w = strings.Trim(w, "\"'(),:;")
			if w == "" {
				continue
			}
			if known[strings.ToLower(canonicalToolName(w))] {
				return true
			}
		}
		return false
	}
	for _, raw := range strings.Split(s, "\n") {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		if callAnnounceRe.MatchString(line) && mentionsTool(line) {
			return true
		}
	}
	return false
}

// matchTextCall applique textCallRe strict puis les garde-fous historiques
// (outil annonce dans ce tour, parametre "string" evident non ambigu).
// Retourne nil si non convertible.
func matchTextCall(s string, tools []provider.Tool) *provider.ToolCall {
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
