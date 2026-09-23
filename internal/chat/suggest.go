package chat

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Suggestion : une question proposée au « salut », taguée par corpus.
// Le clic pose la question avec focus corpus (RAG forcément sollicité).
type Suggestion struct {
	Question string `json:"question"`
	Corpus   string `json:"corpus"`
	Pinned   bool   `json:"pinned"`
}

// SuggestionsFile : nom du pool, lu à la racine du dossier RAG runtime
// (copié depuis RAG/ du dépôt ; ignoré par l'indexeur).
const SuggestionsFile = "suggested-questions.yaml"

// LoadSuggestions : parseur YAML minimal et strict pour le format
// documenté en tête de suggested-questions.yaml :
//
//   - question: "..."
//     corpus: <dossier>
//     pinned: true        # optionnel
//
// Pas de dépendance externe : le schéma est fixe. Les lignes « # » et
// vides sont ignorées. Une entrée sans question ou sans corpus est
// rejetée (erreur). Fail-open : fichier absent => (nil, nil).
func LoadSuggestions(path string) ([]Suggestion, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var out []Suggestion
	var cur *Suggestion
	lineNo := 0
	flush := func() error {
		if cur == nil {
			return nil
		}
		if strings.TrimSpace(cur.Question) == "" || strings.TrimSpace(cur.Corpus) == "" {
			return fmt.Errorf("ligne %d : question ou corpus manquant", lineNo)
		}
		out = append(out, *cur)
		cur = nil
		return nil
	}
	for _, line := range strings.Split(string(raw), "\n") {
		lineNo++
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		if strings.HasPrefix(trimmed, "- ") {
			if err := flush(); err != nil {
				return nil, err
			}
			cur = &Suggestion{}
			trimmed = strings.TrimSpace(strings.TrimPrefix(trimmed, "-"))
			if trimmed == "" {
				continue
			}
		}
		if cur == nil {
			return nil, fmt.Errorf("ligne %d : cle hors entree", lineNo)
		}
		key, val, ok := strings.Cut(trimmed, ":")
		if !ok {
			return nil, fmt.Errorf("ligne %d : syntaxe %q", lineNo, trimmed)
		}
		key = strings.TrimSpace(key)
		val = unquote(strings.TrimSpace(val))
		switch key {
		case "question":
			cur.Question = val
		case "corpus":
			cur.Corpus = val
		case "pinned":
			cur.Pinned = val == "true"
		default:
			return nil, fmt.Errorf("ligne %d : cle inconnue %q", lineNo, key)
		}
	}
	if err := flush(); err != nil {
		return nil, err
	}
	return out, nil
}

// unquote : retire les guillemets simples/doubles d'une valeur scalaire.
func unquote(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			r := strings.ReplaceAll(s[1:len(s)-1], `\"`, `"`)
			return strings.ReplaceAll(r, `\\`, `\`)
		}
		if s[0] == '\'' && s[len(s)-1] == '\'' {
			return strings.ReplaceAll(s[1:len(s)-1], `''`, `'`)
		}
	}
	return s
}

// SuggestionsPath : chemin du pool pour une racine RAG donnée.
func SuggestionsPath(ragRoot string) string {
	return filepath.Join(ragRoot, SuggestionsFile)
}

// greetings : salutations pures déclenchant la diversion d'accueil
// (3 questions suggérées en chips, côté navigateur).
var greetings = map[string]bool{
	"salut": true, "bonjour": true, "bonsoir": true,
	"hello": true, "hi": true, "hey": true,
	"coucou": true, "yo": true,
}

// IsGreeting : vrai si le message est une salutation pure (ponctuation
// finale tolérée). Conservateur : « salut, ça va ? » n'est pas une
// salutation pure et suit le flux normal.
func IsGreeting(text string) bool {
	t := strings.ToLower(strings.TrimSpace(text))
	t = strings.Trim(t, " !?.,…")
	t = strings.TrimSpace(t)
	return greetings[t]
}

// SetSuggestions remplace le pool de questions suggérées de l'engine.
func (e *Engine) SetSuggestions(s []Suggestion) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.suggestions = s
}

// Suggestions retourne une copie du pool (tirage côté navigateur).
func (e *Engine) Suggestions() []Suggestion {
	e.mu.Lock()
	defer e.mu.Unlock()
	out := make([]Suggestion, len(e.suggestions))
	copy(out, e.suggestions)
	return out
}
