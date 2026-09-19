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
	// Phase 3 : Cat est un alias de Read (fusion).
	if tc.Function.Name != "Read" {
		t.Fatalf("nom attendu Read, recu %s", tc.Function.Name)
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
		{"```\nCat \"dnsmasq.conf\"\n```", "Read", "file_path", "dnsmasq.conf"},
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
	// outils inconnus, appels ambigus ou a schema non convertible
	// (Write : plusieurs parametres). Les outils a effets de bord a
	// parametre unique (Bash, Mkdir) sont convertis mais passent par
	// l'approbation (voir TestParseTextToolCallDangerous).
	negatives := []string{
		"Pour lire un fichier, utilise :\nCat \"dnsmasq.conf\"\npuis analyse le contenu.",
		"```go\n// Exemple : Cat \"fichier.txt\" lit un fichier\n```",
		`Frobnicate "truc"`,
		`Write "a.txt"`,
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

// TestParseTextToolCallDangerous : un pseudo-appel dangereux a parametre
// unique (Bash, Mkdir) est converti en vrai ToolCall — mais le pipeline
// standard exigera l'approbation utilisateur avant execution.
func TestParseTextToolCallDangerous(t *testing.T) {
	tools := textCallTools()
	for _, in := range []string{`Bash "echo hello"`, `Mkdir "nouveau"`} {
		tc := parseTextToolCall(in, tools)
		if tc == nil {
			t.Fatalf("%q: pseudo-appel dangereux non converti", in)
		}
		var args map[string]any
		if err := json.Unmarshal([]byte(tc.Function.Arguments), &args); err != nil {
			t.Fatal(err)
		}
		if !needsApprovalFor(tc.Function.Name, args) {
			t.Fatalf("%q: %s devrait exiger une approbation", in, tc.Function.Name)
		}
	}
}
