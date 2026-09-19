package chat

import (
	"log"

	"cetas-lite/internal/provider"
)

// Filet unique "appels d'outils en texte" (Phase 1).
//
// Quand le modele n'emet pas de function calls natifs, il ecrit parfois ses
// appels en texte : balisage DSML historique de l'app, ou pseudo-appel isole
// (etape 1b, ex : Cat "fichier"). parseFallbackToolCalls est le seul point
// d'entree de ce filet : il convertit ces formes en vrais provider.ToolCall,
// executes ensuite par le pipeline standard (approbation, blocs d'outils,
// suivi des modifications).
//
// Ordre d'essai :
//  1. DSML : plusieurs appels possibles, meme noyes dans du texte ;
//  2. pseudo-appels en texte : un par ligne, avec ou sans label d'annonce
//     ("Appel réel :", "Real call:", ...), garde-fous stricts (outil
//     annonce, chaque ligne entierement constituee de l'appel, parametre
//     non ambigu). Les outils a effets de bord sont convertis eux
//     aussi : l'approbation utilisateur du pipeline standard s'applique
//     avant toute execution.
//
// Retourne les appels convertis (nil si rien de convertible) et le texte a
// afficher : le balisage est retire dans tous les cas, et le pseudo-appel
// isole converti disparait de l'affichage (il est rendu comme bloc d'outil).
func parseFallbackToolCalls(content string, tools []provider.Tool) ([]provider.ToolCall, string) {
	if dsml := parseDSMLCalls(content); len(dsml) > 0 {
		for i := range dsml {
			dsml[i].Name = canonicalDSMLName(dsml[i].Name, tools)
		}
		tcs := dsmlToToolCalls(dsml)
		// Phase 3 : les appels vers des outils fusionnés (ex. Cat, via
		// DSML) sont réécrits vers l'outil canonique (Read) — le
		// pipeline (approbation, parallélisme, déduplication) ne voit
		// qu'un seul nom d'outil.
		for i := range tcs {
			tcs[i].Function.Name = canonicalToolName(tcs[i].Function.Name)
		}
		return tcs, stripDSMLFinal(content)
	}
	if tcs := parseTextToolCalls(content, tools); len(tcs) > 0 {
		return tcs, ""
	}
	if hasDSML(content) {
		return nil, stripDSMLFinal(content)
	}
	// Diagnostic : le contenu ressemble a une tentative d'appel annoncee
	// (nom d'outil connu + mot-cle d'annonce) mais le parsing strict l'a
	// refusee — nouvelle formulation d'un modele, argument ambigu, etc.
	// Loggue pour journalctl : c'est le signal direct pour la prochaine
	// fois, au lieu d'un angle mort.
	if looksLikeAnnouncedCall(content, tools) {
		preview := content
		if r := []rune(preview); len(r) > 200 {
			preview = string(r[:200]) + "…"
		}
		log.Printf("chat: parseFallbackToolCalls: tentative d'appel en texte non convertible (contenu: %q)", preview)
	}
	return nil, content
}
