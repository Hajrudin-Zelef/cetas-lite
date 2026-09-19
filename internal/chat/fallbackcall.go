package chat

import (
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
//  2. pseudo-appel isole : un seul appel, garde-fous stricts (outil
//     annonce, reponse entierement constituee de l'appel, parametre
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
	if tc := parseTextToolCall(content, tools); tc != nil {
		return []provider.ToolCall{*tc}, ""
	}
	if hasDSML(content) {
		return nil, stripDSMLFinal(content)
	}
	return nil, content
}
