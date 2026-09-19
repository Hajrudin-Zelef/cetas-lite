package chat

import (
	"encoding/json"
	"testing"

	"cetas-lite/internal/provider"
)

func textCallTools() []provider.Tool { return ToolSchemas() }

func TestParseTextToolCallBasic(t *testing.T) {
	tc := parseTextToolCall(`Cat "NEVA PVE/dnsmasq.conf"`, textCallTools())
	if tc == nil {
		t.Fatal("pseudo-appel Cat non converti")
	}
	if tc.Function.Name != "Cat" {
		t.Fatalf("nom attendu Cat, recu %s", tc.Function.Name)
	}
	var args map[string]any
	if err := json.Unmarshal([]byte(tc.Function.Arguments), &args); err != nil {
		t.Fatal(err)
	}
	if args["file_path"] != "NEVA PVE/dnsmasq.conf" {
		t.Fatalf("file_path inattendu: %v", args)
	}
}

func TestParseTextToolCallVariants(t *testing.T) {
	tools := textCallTools()
	cases := []struct {
		in       string
		tool     string
		paramKey string
		paramVal string
	}{
		{`read 'docs/a.txt'`, "Read", "file_path", "docs/a.txt"},
		{`Read("docs/a.txt")`, "Read", "file_path", "docs/a.txt"},
		{"```\nCat \"dnsmasq.conf\"\n```", "Cat", "file_path", "dnsmasq.conf"},
		{`Grep "func main"`, "Grep", "pattern", "func main"},
		{`Glob "**/*.go"`, "Glob", "pattern", "**/*.go"},
		{`Echo "bonjour"`, "Echo", "text", "bonjour"},
		{`Ls`, "Ls", "", ""},
		{`Tree "docs"`, "Tree", "path", "docs"},
	}
	for _, c := range cases {
		tc := parseTextToolCall(c.in, tools)
		if tc == nil {
			t.Errorf("%q: non converti", c.in)
			continue
		}
		if tc.Function.Name != c.tool {
			t.Errorf("%q: outil %s, attendu %s", c.in, tc.Function.Name, c.tool)
		}
		if c.paramKey != "" {
			var args map[string]any
			_ = json.Unmarshal([]byte(tc.Function.Arguments), &args)
			if args[c.paramKey] != c.paramVal {
				t.Errorf("%q: %s=%v, attendu %q", c.in, c.paramKey, args[c.paramKey], c.paramVal)
			}
		}
	}
}

func TestParseTextToolCallNoFalsePositive(t *testing.T) {
	tools := textCallTools()
	// Ne jamais executer : exemples dans du texte, citations, prose,
	// outils inconnus, outils a effet de bord, appels ambigus.
	negatives := []string{
		"Pour lire un fichier, utilise :\nCat \"dnsmasq.conf\"\npuis analyse le contenu.",
		"```go\n// Exemple : Cat \"fichier.txt\" lit un fichier\n```",
		`Frobnicate "truc"`,
		`Write "a.txt"`,
		`Bash "rm -rf /"`,
		`Mkdir "nouveau"`,
		`Cat "a" "b"`,
		`voici le contenu`,
		`oui`,
		"",
		"```\nCat \"a\"\nCat \"b\"\n```",
	}
	for _, in := range negatives {
		if tc := parseTextToolCall(in, tools); tc != nil {
			t.Errorf("%q: faux positif converti en %s", in, tc.Function.Name)
		}
	}
}

func TestLooksLikeToolAttempt(t *testing.T) {
	tools := textCallTools()
	if name, ok := looksLikeToolAttempt("Je vais appeler l'outil Read pour lire le fichier.", tools); !ok || name != "Read" {
		t.Errorf("intention FR non detectee: %q %v", name, ok)
	}
	if _, ok := looksLikeToolAttempt("Let me call the Cat tool on this file.", tools); !ok {
		t.Error("intention EN non detectee")
	}
	if name, ok := looksLikeToolAttempt("Cat \"dnsmasq.conf\"", tools); !ok || name != "Cat" {
		t.Errorf("pseudo-appel isole non detecte: %q %v", name, ok)
	}
	// Longue explication avec exemple : pas de nudge.
	long := "Voici comment lire un fichier dans cet environnement de travail. " +
		"Tu peux utiliser l'outil Cat suivi du chemin entre guillemets, par exemple " +
		"Cat \"exemple.txt\" pour afficher le contenu. Pense aussi a verifier les permissions " +
		"avant toute ecriture, et a documenter chaque etape de ton raisonnement en detail."
	if _, ok := looksLikeToolAttempt(long, tools); ok {
		t.Error("faux positif sur longue explication")
	}
	if _, ok := looksLikeToolAttempt("Le fichier contient trois sections principales.", tools); ok {
		t.Error("faux positif sur texte normal")
	}
}

func TestToolAttemptNudgeText(t *testing.T) {
	s := toolAttemptNudgeText("Read")
	if s == "" {
		t.Fatal("nudge vide")
	}
	if s2 := toolAttemptNudgeText(""); s2 == "" {
		t.Fatal("nudge generique vide")
	}
}
