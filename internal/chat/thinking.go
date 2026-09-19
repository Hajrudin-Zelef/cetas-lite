package chat

// resolveEffort détermine l'effort de raisonnement à transmettre au
// provider. Quand le thinking est désactivé, l'effort est toujours vide :
// un effort non vide sans raisonnement activé serait incohérent (et
// certains providers le refusent).
func resolveEffort(think bool, text, requested string) string {
	if !think {
		return ""
	}
	switch requested {
	case "low", "medium", "high":
		return requested
	}
	if len([]rune(text)) > 400 {
		return "medium"
	}
	return "low"
}

// resolveAgentEffort : effort de raisonnement pour la vue Agents. Le
// "Défaut" du composer vaut "low" (et non "medium" selon la longueur) :
// la boucle agentique itère déjà, un raisonnement long par itération
// s'additionne et rend l'expérience lente. "low" = réfléchir brièvement
// puis agir vite. Choix explicite Faible/Moyen/Max toujours honoré.
func resolveAgentEffort(think bool, requested string) string {
	if !think {
		return ""
	}
	switch requested {
	case "low", "medium", "high":
		return requested
	}
	return "low"
}
