package chat

import "cetas-lite/internal/alias"

// autoTierFor choisit le tier (mode) d'une famille cloud pour le mode
// "auto" (menu + : « Défaut ») en fonction de l'effort demandé.
//
// Les modes d'une famille sont ordonnés par coût croissant
// (flash < standard < elite). Règles :
//   - "high" explicite -> le tier le plus cher ;
//   - "medium" -> le tier médian (jamais le plus cher si > 2 tiers) ;
//   - "low" -> le tier le moins cher ;
//   - "default" (ou inconnu) -> reprend l'heuristique de resolveEffort
//     (texte long > 400 runes -> médian, sinon le moins cher).
//
// Seul un effort "high" explicite peut donc sélectionner le tier le plus
// cher : garde-fou coût. Avec un seul mode (ex. Nano), l'effort est sans
// effet et le mode unique est retourné.
func autoTierFor(modes []alias.Mode, requestedEffort, text string) string {
	n := len(modes)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return modes[0].ID
	}
	last := n - 1
	mid := (n - 1) / 2 // médian bas : jamais le plus cher pour n >= 2
	switch requestedEffort {
	case "high":
		return modes[last].ID
	case "medium":
		return modes[mid].ID
	case "low":
		return modes[0].ID
	default:
		if len([]rune(text)) > 400 {
			return modes[mid].ID
		}
		return modes[0].ID
	}
}
