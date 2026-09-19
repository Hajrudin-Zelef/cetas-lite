package chat

// resolveEffort détermine l'effort de raisonnement à transmettre au
// provider. Quand le thinking est désactivé, l'effort est toujours vide :
// un effort non vide sans raisonnement activé serait incohérent (et
// certains providers le refusent).
import "strings"

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

// mechanicalVerbs : verbes de demande de lecture simple (FR/EN). Un
// message qui commence par l'un d'eux est probablement un tour 100 %
// mécanique (lire, lister, chercher) : le raisonnement obligatoire y est
// contre-productif.
var mechanicalVerbs = []string{
	"lis", "lire", "montre", "affiche", "liste", "lister",
	"cherche", "chercher", "trouve", "trouver", "ouvre", "ouvrir",
	"read", "list", "show", "display", "find", "search", "open", "grep",
}

// mechanicalActionStems : si le message contient l'un de ces radicaux,
// ce n'est PAS un tour purement mécanique (il y a aussi une action).
var mechanicalActionStems = []string{
	"corrig", "modifi", "écris", "ecris", "cré", "supprim",
	"ajoute", "change", "refactor", "fix", "update", "commit",
}

// mechanicalTurn détecte une demande simple de lecture (phase 4) : message
// court commençant par un verbe de lecture, sans verbe d'action.
func mechanicalTurn(text string) bool {
	t := strings.ToLower(strings.TrimSpace(text))
	if t == "" || len([]rune(t)) > 140 {
		return false
	}
	for _, s := range mechanicalActionStems {
		if strings.Contains(t, s) {
			return false
		}
	}
	first := t
	if i := strings.IndexAny(t, " \t\n.,!?;:"); i >= 0 {
		first = t[:i]
	}
	first = strings.Trim(first, "'\"«»")
	if i := strings.Index(first, "-"); i >= 0 {
		first = first[:i] // "montre-moi" -> "montre"
	}
	for _, v := range mechanicalVerbs {
		if first == v {
			return true
		}
	}
	return false
}

// mechanicalThinkDirective : directive pour les tours mécaniques (phase
// 4). Le raisonnement n'est pas obligatoire : le modèle appelle
// directement le(s) outil(s) de lecture puis répond brièvement. Réduit
// les tokens de thinking sur les opérations simples.
func mechanicalThinkDirective() string {
	return "No reasoning needed for this simple read-only request: " +
		"call the needed tool(s) directly, then answer briefly in the user's language."
}

// agentThinkDirective choisit la directive de raisonnement du tour : la
// directive mécanique (sans raisonnement obligatoire) quand le message
// est une simple demande de lecture et que l'utilisateur n'a pas imposé
// un effort explicite, la directive standard sinon.
func agentThinkDirective(think bool, text, requested string) string {
	effort := resolveEffort(think, text, requested)
	if think && effort == "low" && requested == "" && mechanicalTurn(text) {
		return mechanicalThinkDirective()
	}
	return thinkDirective(true, think, effort)
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
