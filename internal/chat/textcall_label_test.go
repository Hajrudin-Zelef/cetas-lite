package chat

import (
	"encoding/json"
	"testing"
)

// Point A : un label d'annonce ("Appel réel :", "Real call:", ...) devant le
// pseudo-appel est toléré — c'est le pattern observé qui faisait figer le
// compteur d'outils (le parsing strict refusait le préfixe).
func TestParseTextToolCallLabel(t *testing.T) {
	tools := textCallTools()
	cases := []struct {
		in       string
		tool     string
		paramKey string
		paramVal string
	}{
		{`Appel réel : Read "docs/a.txt"`, "Read", "file_path", "docs/a.txt"},
		{`Real call: Ls`, "Ls", "", ""},
		{`Calling tool: Grep "motif"`, "Grep", "pattern", "motif"},
		{`J'appelle : Glob "**/*.go"`, "Glob", "pattern", "**/*.go"},
		{`Appel outil : Cat "dnsmasq.conf"`, "Read", "file_path", "dnsmasq.conf"},
		{`APPEL RÉEL : Ls`, "Ls", "", ""},
		// Forme stricte historique sans label : inchangée.
		{`Read "docs/a.txt"`, "Read", "file_path", "docs/a.txt"},
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

func TestParseTextToolCallLabelNoFalsePositive(t *testing.T) {
	tools := textCallTools()
	negatives := []string{
		`Voici le résultat : Read the docs`, // label neutre, pas d'annonce
		`Note : Ls`,                         // label neutre court
		`Rappel : Read "x"`,                 // "rappel" != mot-clé (borne)
		`Appel réel : Frobnicate "truc"`,    // outil inconnu
		`Appel réel : Write "a" "b"`,        // schéma ambigu : refusé
		`Appel réel : Ls src`,               // Ls sans paramètre : refusé
		"Pour lire un fichier, utilise :\nCat \"dnsmasq.conf\"\npour continuer.",
	}
	for _, in := range negatives {
		if tc := parseTextToolCall(in, tools); tc != nil {
			t.Errorf("%q: faux positif converti en %s", in, tc.Function.Name)
		}
	}
}

// Point A (multi-lignes) : le cas observé — deux appels annoncés sur deux
// lignes. Toutes les lignes doivent être convertibles, sinon rien
// (jamais d'exécution partielle).
func TestParseTextToolCallsMultiLine(t *testing.T) {
	tools := textCallTools()

	tcs := parseTextToolCalls("Appel réel : Ls\nReal call: Glob \"**/*.go\"", tools)
	if len(tcs) != 2 {
		t.Fatalf("attendu 2 appels, reçu %d", len(tcs))
	}
	if tcs[0].Function.Name != "Ls" || tcs[1].Function.Name != "Glob" {
		t.Fatalf("outils inattendus : %s, %s", tcs[0].Function.Name, tcs[1].Function.Name)
	}

	// Une seule ligne non convertible => rien n'est converti.
	if tcs := parseTextToolCalls("Appel réel : Ls\ndu texte", tools); len(tcs) != 0 {
		t.Errorf("exécution partielle : %d appels convertis", len(tcs))
	}

	// Sans label non plus : chaque ligne reste soumise aux garde-fous.
	tcs = parseTextToolCalls("Cat \"a\"\nGrep \"b\"", tools)
	if len(tcs) != 2 || tcs[0].Function.Name != "Read" || tcs[1].Function.Name != "Grep" {
		t.Fatalf("multi-lignes sans label non converti : %+v", tcs)
	}
}

// Points B/C : le détecteur ne se déclenche que sur une VRAIE tentative
// d'appel (mot-clé d'annonce + nom d'outil reconnu), jamais sur une
// réponse texte normale.
func TestLooksLikeAnnouncedCall(t *testing.T) {
	tools := textCallTools()
	positives := []string{
		`Appel réel : Write "a" "b"`, // refusé par le strict (ambigu) mais détecté
		`Je vais appeler l'outil Read pour lire le fichier`,
		"Appel réel : Ls src\nReal call: Glob", // Ls src non convertible mais détecté
		`Calling tool: Bash`,
	}
	for _, in := range positives {
		if !looksLikeAnnouncedCall(in, tools) {
			t.Errorf("%q: tentative d'appel non détectée", in)
		}
	}
	negatives := []string{
		"bonjour, comment ça va ?",
		"read the docs", // nom d'outil sans annonce : pas de nudge inutile
		"Voici le résultat : tout va bien",
		`Rappel : tout est prêt`,
		"",
	}
	for _, in := range negatives {
		if looksLikeAnnouncedCall(in, tools) {
			t.Errorf("%q: faux positif du détecteur", in)
		}
	}
}
