package rag

import (
	"strings"
	"unicode"
)

// EnrichQueryWithHistory : enrichit une requete elliptique avec le sujet
// trouve dans l'historique de conversation (iteration 6).
//
// Cas vise : « on peut la faire tourner sur combien de cpu ? » apres un
// tour sur « Kimi K3 ». Le nettoyage de l'iteration 3 retire les pronoms
// (« la ») comme mots-outils, et BM25 recoit « peut faire tourner combien
// cpu » sans le sujet — d'ou des extraits hors sujet alors que la base
// couvre la question. Ici, si la requete ne porte aucune entite et
// ressemble a une ellipse (pronom ou tres courte), le sujet est repris
// des messages recents : « …cpu ? Kimi K3 ».
//
// history : contenus textuels des messages, ordre chronologique (le plus
// recent en dernier). Fail-open : historique vide, requete deja porteuse
// d'entites, requete non elliptique, ou aucun sujet trouve => requete
// inchangee et entites vides.
//
// Retourne aussi les tokens d'entite ajoutes (post-normalisation, meme
// forme que tokenize), cles du boost de scoring (iteration 6b).
func EnrichQueryWithHistory(query string, history []string) (string, []string) {
	if len(history) == 0 || hasStrongEntity(query) || !looksElliptical(query) {
		return query, nil
	}
	ents := entitiesFromHistory(history, query)
	if len(ents) == 0 {
		return query, nil
	}
	var keys []string
	for _, g := range ents {
		keys = append(keys, tokenize(g)...)
	}
	return query + " " + strings.Join(ents, " "), uniqueTerms(keys)
}

// hasStrongEntity : la requete porte deja un sujet identifiable —
// inutile de l'enrichir. Fort : token avec chiffre (k3, v3.2, 104b),
// passage entre guillemets, camelCase (vLLM), ou mot capitalise (sauf
// capitalisation de debut de phrase, ex. « Qu'est-ce… »).
func hasStrongEntity(q string) bool {
	toks := rawTokens(q)
	for _, t := range toks {
		if hasDigit(t) || isQuoted(t) || hasInteriorUpper(t) {
			return true
		}
	}
	for i, t := range toks {
		if isCapitalized(t) && (i > 0 || len(toks) == 1) {
			return true
		}
	}
	return false
}

// enrichPronouns : pronoms (formes normalisees) signalant une ellipse —
// le sujet est ailleurs, typiquement dans le tour precedent.
var enrichPronouns = map[string]bool{
	"la": true, "le": true, "les": true, "lui": true,
	"elle": true, "il": true, "ils": true, "elles": true,
	"on": true, "y": true, "en": true,
	"ca": true, "cela": true, "ceci": true,
	"celui": true, "celle": true, "ceux": true, "celles": true,
	"ce": true, "cet": true, "cette": true, "ces": true,
}

// looksElliptical : pronom present, ou requete reduite a presque rien
// apres nettoyage (ex. « et le prix ? » -> « prix »).
func looksElliptical(q string) bool {
	for _, t := range tokenize(q) {
		if enrichPronouns[t] {
			return true
		}
	}
	return len(tokenize(cleanQueryForSearch(q))) <= 2
}

// entitiesFromHistory : extrait jusqu'a 2 groupes d'entites du message
// le plus recent qui en contient (un groupe = suite de tokens « entite »
// consecutifs, ex. « Kimi K3 », « Moonshot AI »). Les groupes deja
// presents dans la requete sont exclus. Seuls les 4 derniers messages
// sont examines.
func entitiesFromHistory(history []string, query string) []string {
	nq := " " + normalize(query) + " "
	start := len(history) - 4
	if start < 0 {
		start = 0
	}
	for i := len(history) - 1; i >= start; i-- {
		var out []string
		for _, g := range entityGroups(history[i]) {
			if strings.Contains(nq, " "+normalize(g)+" ") {
				continue
			}
			out = append(out, g)
			if len(out) == 2 {
				break
			}
		}
		if len(out) > 0 {
			return out
		}
	}
	return nil
}

// entityGroups : decoupe les groupes de tokens « entite » consecutifs
// d'un message.
func entityGroups(s string) []string {
	toks := rawTokens(s)
	var groups []string
	var cur []string
	flush := func() {
		if len(cur) > 0 {
			groups = append(groups, strings.Join(cur, " "))
			cur = nil
		}
	}
	for i, t := range toks {
		if isEntityToken(t, i == 0) {
			cur = append(cur, t)
		} else {
			flush()
		}
	}
	flush()
	return groups
}

// isEntityToken : token « porteur de sujet ». En debut de message, une
// simple capitalisation ne suffit pas (debut de phrase) : il faut un
// chiffre, des guillemets, un sigle ou une majuscule interieure.
func isEntityToken(t string, first bool) bool {
	if hasDigit(t) || isQuoted(t) || hasInteriorUpper(t) || isAllUpper(t) {
		return true
	}
	return !first && isCapitalized(t)
}

// rawTokens : decoupe sur les blancs et nettoie la ponctuation de bord,
// en gardant la casse d'origine (les entites se reperent a la majuscule)
// et les traits d'union interieurs (GLM-5.3).
func rawTokens(s string) []string {
	var out []string
	for _, f := range strings.Fields(s) {
		if t := strings.Trim(f, "()[]{}:;,.!?…—–"); t != "" {
			out = append(out, t)
		}
	}
	return out
}

func hasDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func isQuoted(t string) bool {
	for _, q := range []string{`"`, `'`, "«", "“", "‘"} {
		if strings.HasPrefix(t, q) && len(t) > len(q) {
			return true
		}
	}
	return false
}

func isCapitalized(t string) bool {
	r := []rune(t)
	return len(r) > 2 && unicode.IsUpper(r[0])
}

func hasInteriorUpper(t string) bool {
	for _, r := range []rune(t)[1:] {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

func isAllUpper(t string) bool {
	letters := 0
	for _, r := range t {
		if unicode.IsLetter(r) {
			letters++
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	return letters >= 2
}
