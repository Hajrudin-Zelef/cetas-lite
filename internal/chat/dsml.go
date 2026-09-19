package chat

// dsml.go — prise en charge du format texte d'appels d'outils "DSML".
//
// Certains modeles (ex : DeepSeek) emettent leurs appels d'outils en texte
// plutot qu'en function-calls natifs, surtout quand les outils natifs sont
// indisponibles ou apres une erreur provider :
//
//	< | | DSML | | calls>
//	< | | DSML | | invoke name="Read">
//	< | | DSML | | parameter name="file_path" string="true">docs/a.md</ | | DSML | | parameter>
//	</ | | DSML | | invoke>
//	</ | | DSML | | calls>
//
// Principe (inspire du harness DeepSeek) : ce format texte est CONVERTI en
// vrais appels d'outils — executes cote backend et rendus en blocs d'outils
// dans la vue Agents — au lieu d'etre affiche en brut. Le texte DSML est
// masque pendant le streaming (stripDSMLStreaming) puis retire du contenu
// final (stripDSMLFinal).

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"cetas-lite/internal/provider"
)

// dsmlCall est un appel d'outil extrait d'un bloc DSML.
type dsmlCall struct {
	Name string
	Args map[string]string
}

// dsmlTagRe reconnait une balise DSML, tolerante aux espaces, a la casse et
// aux guillemets. Groupes : 0=balise complete, 1="/", 2=kind
// (calls|invoke|parameter), 3=attributs.
var dsmlTagRe = regexp.MustCompile(`(?i)<\s*(/)?\s*\|\s*\|\s*DSML\s*\|\s*\|\s*(calls|invoke|parameter)\b([^<>]*)>`)

// dsmlAttrRe reconnait un attribut name="v", name='v' ou name=v.
var dsmlAttrRe = regexp.MustCompile(`([A-Za-z_][\w.-]*)\s*=\s*(?:"([^"]*)"|'([^']*)'|([^\s"'=<>/]+))`)

// normalizeDSML remplace les variantes pleine-chasse (|, /, <, >) par leur
// equivalent ASCII avant analyse.
func normalizeDSML(s string) string {
	if strings.ContainsAny(s, "｜＜＞∣／") {
		r := strings.NewReplacer("｜", "|", "∣", "|", "＜", "<", "＞", ">", "／", "/")
		return r.Replace(s)
	}
	return s
}

func hasDSML(s string) bool {
	return strings.Contains(strings.ToUpper(s), "DSML")
}

// dsmlAttr lit la valeur d'un attribut dans une chaine d'attributs.
func dsmlAttr(attrs, key string) string {
	for _, m := range dsmlAttrRe.FindAllStringSubmatch(attrs, -1) {
		if !strings.EqualFold(m[1], key) {
			continue
		}
		if m[2] != "" {
			return m[2]
		}
		if m[3] != "" {
			return m[3]
		}
		return m[4]
	}
	return ""
}

// parseDSMLCalls extrait les appels d'outils des blocs DSML du texte.
// Les invokes hors bloc <calls> sont aussi acceptes (bloc implicite).
func parseDSMLCalls(text string) []dsmlCall {
	norm := normalizeDSML(text)
	if !hasDSML(norm) {
		return nil
	}
	var calls []dsmlCall
	cur := -1 // index de l'invoke courant, -1 = aucun
	matches := dsmlTagRe.FindAllStringSubmatchIndex(norm, -1)
	for i := 0; i < len(matches); i++ {
		m := matches[i]
		closing := m[2] >= 0
		kind := strings.ToLower(norm[m[4]:m[5]])
		attrs := norm[m[6]:m[7]]
		tagEnd := m[1]
		switch {
		case kind == "calls" && closing:
			cur = -1
		case kind == "invoke" && !closing:
			name := dsmlAttr(attrs, "name")
			if name == "" {
				cur = -1
				continue
			}
			calls = append(calls, dsmlCall{Name: name, Args: map[string]string{}})
			cur = len(calls) - 1
		case kind == "invoke" && closing:
			cur = -1
		case kind == "parameter" && !closing && cur >= 0:
			pname := dsmlAttr(attrs, "name")
			// Valeur = texte brut jusqu'a la balise fermante </ | | DSML | | parameter >.
			j := i + 1
			for j < len(matches) {
				mj := matches[j]
				kj := strings.ToLower(norm[mj[4]:mj[5]])
				if kj == "parameter" && mj[2] >= 0 {
					break
				}
				j++
			}
			var val string
			if j < len(matches) {
				val = strings.TrimSpace(norm[tagEnd:matches[j][0]])
				i = j // consomme jusqu'a la fermante incluse
			} else {
				// Parametre non ferme (fin de flux) : prend le reste.
				val = strings.TrimSpace(norm[tagEnd:])
				i = len(matches)
			}
			if pname != "" {
				calls[cur].Args[pname] = val
			}
		}
	}
	return calls
}

// dsmlBlockSpans retourne les intervalles [debut, fin) du texte normalise a
// retirer de l'affichage : blocs <calls>...</calls> (fermes ou traînants) et
// invokes de premier niveau hors bloc calls.
func dsmlBlockSpans(norm string) [][2]int {
	var spans [][2]int
	matches := dsmlTagRe.FindAllStringSubmatchIndex(norm, -1)
	i := 0
	for i < len(matches) {
		m := matches[i]
		closing := m[2] >= 0
		kind := strings.ToLower(norm[m[4]:m[5]])
		if closing || (kind != "calls" && kind != "invoke") {
			i++
			continue
		}
		// Cherche la fermante correspondante de meme kind (profondeur 1).
		depth := 1
		j := i + 1
		for j < len(matches) {
			mj := matches[j]
			kj := strings.ToLower(norm[mj[4]:mj[5]])
			if kj == kind {
				if mj[2] >= 0 {
					depth--
				} else {
					depth++
				}
				if depth == 0 {
					break
				}
			}
			j++
		}
		end := len(norm)
		if j < len(matches) {
			end = matches[j][1]
		}
		spans = append(spans, [2]int{m[0], end})
		if j < len(matches) {
			i = j + 1
		} else {
			break
		}
	}
	return spans
}

// stripDSML retire les blocs DSML du texte affiche.
// En mode streaming, un bloc non ferme masque tout ce qui suit (pas de flash
// de syntaxe brute) ; en mode final, le bloc traînant est supprime aussi.
func stripDSML(text string, streaming bool) string {
	norm := normalizeDSML(text)
	if !hasDSML(norm) {
		return text
	}
	spans := dsmlBlockSpans(norm)
	if len(spans) == 0 {
		// Marqueur DSML present mais aucun bloc structure : en streaming on
		// masque a partir du premier tag (bloc en cours d'arrivee), en final
		// on laisse le texte tel quel.
		if streaming {
			if loc := dsmlTagRe.FindStringIndex(norm); loc != nil {
				return norm[:loc[0]]
			}
		}
		return norm
	}
	var sb strings.Builder
	sb.Grow(len(norm))
	pos := 0
	for _, sp := range spans {
		if sp[0] > pos {
			sb.WriteString(norm[pos:sp[0]])
		}
		if sp[1] > pos {
			pos = sp[1]
		}
	}
	sb.WriteString(norm[pos:])
	return sb.String()
}

// stripDSMLStreaming masque les blocs DSML pendant le streaming.
func stripDSMLStreaming(text string) string { return stripDSML(text, true) }

// stripDSMLFinal retire les blocs DSML du contenu final.
func stripDSMLFinal(text string) string { return stripDSML(text, false) }

// dsmlToToolCalls convertit les appels DSML en ToolCalls natifs (IDs
// synthetiques), directement executables par le pipeline d'outils existant.
func dsmlToToolCalls(calls []dsmlCall) []provider.ToolCall {
	out := make([]provider.ToolCall, 0, len(calls))
	for i, c := range calls {
		argsJSON, _ := json.Marshal(c.Args)
		out = append(out, provider.ToolCall{
			ID:   fmt.Sprintf("dsml-%d", i+1),
			Type: "function",
			Function: provider.Func{
				Name:      c.Name,
				Arguments: string(argsJSON),
			},
		})
	}
	return out
}

// canonicalDSMLName recale le nom d'outil DSML sur le nom canonique du schema
// (comparaison insensible a la casse : "read" -> "Read"). Si aucun outil
// disponible ne correspond, le nom est conserve tel quel et le pipeline
// renverra l'erreur standard "outil inconnu".
func canonicalDSMLName(name string, tools []provider.Tool) string {
	for _, t := range tools {
		if strings.EqualFold(t.Function.Name, name) {
			return t.Function.Name
		}
	}
	return name
}

// dsmlTagLoc decrit une balise DSML localisee dans le texte.
type dsmlTagLoc struct {
	start, end int
	closing    bool
	kind       string
}

// findDSMLTag retourne la premiere balise DSML (tous kinds) dans s.
func findDSMLTag(s string) *dsmlTagLoc {
	m := dsmlTagRe.FindStringSubmatchIndex(s)
	if m == nil {
		return nil
	}
	return &dsmlTagLoc{start: m[0], end: m[1], closing: m[2] >= 0, kind: strings.ToLower(s[m[4]:m[5]])}
}

// findDSMLTagKind retourne la premiere balise DSML du kind demande dans s.
func findDSMLTagKind(s, kind string) *dsmlTagLoc {
	for _, m := range dsmlTagRe.FindAllStringSubmatchIndex(s, -1) {
		if strings.EqualFold(s[m[4]:m[5]], kind) {
			return &dsmlTagLoc{start: m[0], end: m[1], closing: m[2] >= 0, kind: strings.ToLower(s[m[4]:m[5]])}
		}
	}
	return nil
}

// dsmlTagPrefixish dit si s pourrait etre le DEBUT d'une balise DSML coupee
// par une frontiere de chunk (ex : "<||DS", "<||DSML||inv", "<||DSML||invoke
// name=\"fi"). Les espaces sont ignores pour la comparaison.
func dsmlTagPrefixish(s string) bool {
	t := strings.Map(func(r rune) rune {
		switch r {
		case ' ', '\t', '\n', '\r':
			return -1
		}
		return r
	}, s)
	if !strings.HasPrefix(t, "<") {
		return false
	}
	rest := t[1:]
	if strings.HasPrefix(rest, "/") {
		rest = rest[1:]
	}
	const marker = "||DSML||"
	if strings.HasPrefix(marker, rest) {
		return true // encore dans le marqueur ("<", "<||D", ...)
	}
	if !strings.HasPrefix(rest, marker) {
		return false
	}
	kindPart := rest[len(marker):]
	for _, k := range []string{"calls", "invoke", "parameter"} {
		if strings.HasPrefix(k, kindPart) {
			return true // kind partiel : "<||DSML||ca"
		}
		if strings.HasPrefix(kindPart, k) {
			// Kind complet + debut d'attributs : pas de < ni > sinon la
			// balise serait complete (cas gere par la regex principale).
			if !strings.ContainsAny(kindPart[len(k):], "<>") {
				return true
			}
		}
	}
	return false
}

// dsmlStreamFilter masque les blocs DSML au fil de l'eau pour un flux
// append-only (le frontend ne fait qu'ajouter les deltas "content") : le
// texte DSML n'est jamais emis, le texte normal passe immediatement. Un
// eventuel debut de balise en fin de chunk est retenu jusqu'au chunk suivant
// pour ne jamais laisser fuiter de syntaxe brute.
type dsmlStreamFilter struct {
	buf     strings.Builder
	inBlock bool
	kind    string
	depth   int
}

// push absorbe un chunk et retourne le texte affichable (sans DSML).
func (f *dsmlStreamFilter) push(chunk string) string {
	if chunk == "" {
		return ""
	}
	f.buf.WriteString(normalizeDSML(chunk))
	var out strings.Builder
	for {
		s := f.buf.String()
		if s == "" {
			break
		}
		if f.inBlock {
			t := findDSMLTagKind(s, f.kind)
			if t == nil {
				// Contenu de bloc : jete, sauf un eventuel debut de balise
				// (fermante) coupe par la frontiere de chunk, retenu pour
				// le chunk suivant.
				keep := ""
				if i := strings.LastIndexByte(s, '<'); i >= 0 &&
					strings.IndexByte(s[i:], '>') < 0 && dsmlTagPrefixish(s[i:]) {
					keep = s[i:]
				}
				f.buf.Reset()
				f.buf.WriteString(keep)
				break
			}
			if t.closing {
				f.depth--
				if f.depth == 0 {
					f.inBlock = false
				}
			} else {
				f.depth++
			}
			f.buf.Reset()
			f.buf.WriteString(s[t.end:])
			continue
		}
		// Hors bloc : zone sure = tout sauf un eventuel debut de balise
		// retenu en fin de tampon.
		hold := len(s)
		if i := strings.LastIndexByte(s, '<'); i >= 0 &&
			strings.IndexByte(s[i:], '>') < 0 && dsmlTagPrefixish(s[i:]) {
			hold = i
		}
		t := findDSMLTag(s[:hold])
		if t == nil {
			out.WriteString(s[:hold])
			f.buf.Reset()
			f.buf.WriteString(s[hold:])
			break
		}
		out.WriteString(s[:t.start])
		f.buf.Reset()
		f.buf.WriteString(s[t.end:])
		if !t.closing {
			f.inBlock = true
			f.kind = t.kind
			f.depth = 1
		}
		// Balise fermante isolee hors bloc : jetee silencieusement.
	}
	return out.String()
}

// flush termine le flux : le reliquat de texte normal est emis ; le contenu
// d'un bloc DSML non ferme et un debut de balise inacheve sont jetes.
func (f *dsmlStreamFilter) flush() string {
	s := f.buf.String()
	inBlock := f.inBlock
	f.buf.Reset()
	f.inBlock = false
	if s == "" || inBlock {
		return ""
	}
	if i := strings.LastIndexByte(s, '<'); i >= 0 &&
		strings.IndexByte(s[i:], '>') < 0 && dsmlTagPrefixish(s[i:]) {
		return s[:i]
	}
	return s
}
