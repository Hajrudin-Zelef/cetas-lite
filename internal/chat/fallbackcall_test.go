package chat

import (
	"strings"
	"testing"

	"cetas-lite/internal/provider"
)

// Tests du filet unique Phase 1 : un seul point d'entree pour les appels
// emis en texte (DSML ou pseudo-appel isole).

func fallbackTools() []provider.Tool { return ToolSchemas() }

func TestFallbackPrefersDSML(t *testing.T) {
	text := "Voici :\n<||DSML||calls>\n<||DSML||invoke name=\"read\">\n<||DSML||parameter name=\"file_path\">docs/a.md</||DSML||parameter>\n</||DSML||invoke>\n</||DSML||calls>\nFin."
	tcalls, stripped := parseFallbackToolCalls(text, fallbackTools())
	if len(tcalls) != 1 {
		t.Fatalf("attendu 1 appel, obtenu %d", len(tcalls))
	}
	// Nom recale sur le canonique du schema.
	if tcalls[0].Function.Name != "Read" {
		t.Errorf("nom = %q, attendu Read", tcalls[0].Function.Name)
	}
	// Le balisage est retire du texte affiche, le texte normal reste.
	if strings.Contains(stripped, "DSML") {
		t.Errorf("balisage non retire : %q", stripped)
	}
	if !strings.Contains(stripped, "Voici") || !strings.Contains(stripped, "Fin.") {
		t.Errorf("texte normal perdu : %q", stripped)
	}
}

func TestFallbackTextCall(t *testing.T) {
	tcalls, stripped := parseFallbackToolCalls(`Cat "NEVA PVE/dnsmasq.conf"`, fallbackTools())
	if len(tcalls) != 1 {
		t.Fatalf("attendu 1 appel, obtenu %d", len(tcalls))
	}
	// Phase 3 : Cat est fusionné dans Read ; le pseudo-appel est réécrit.
	if tcalls[0].Function.Name != "Read" {
		t.Errorf("nom = %q, attendu Read", tcalls[0].Function.Name)
	}
	// Le pseudo-appel converti disparait de l'affichage (rendu en bloc).
	if stripped != "" {
		t.Errorf("texte affiche = %q, attendu vide", stripped)
	}
}

func TestFallbackNothingToConvert(t *testing.T) {
	text := "Le fichier contient trois sections principales."
	tcalls, stripped := parseFallbackToolCalls(text, fallbackTools())
	if len(tcalls) != 0 {
		t.Fatalf("attendu 0 appel, obtenu %d", len(tcalls))
	}
	if stripped != text {
		t.Errorf("texte modifie : %q", stripped)
	}
}

func TestFallbackStripsUnparseableDSML(t *testing.T) {
	// Balisage present mais aucun appel parsable : pas d'appel, mais le
	// balisage ne doit pas fuiter dans l'affichage (comportement historique).
	text := "regarde <||DSML||calls> ceci"
	tcalls, stripped := parseFallbackToolCalls(text, fallbackTools())
	if len(tcalls) != 0 {
		t.Fatalf("attendu 0 appel, obtenu %d", len(tcalls))
	}
	if strings.Contains(stripped, "DSML") {
		t.Errorf("balisage non retire : %q", stripped)
	}
}

func TestFallbackRefusesUnsafeTextCall(t *testing.T) {
	// Pseudo-appel vers un outil a effet de bord : refuse (garde-fou),
	// texte conserve tel quel pour affichage.
	tcalls, stripped := parseFallbackToolCalls(`Bash "ls -la"`, fallbackTools())
	if len(tcalls) != 0 {
		t.Fatalf("attendu 0 appel (outil a effet de bord), obtenu %d", len(tcalls))
	}
	if stripped != `Bash "ls -la"` {
		t.Errorf("texte modifie : %q", stripped)
	}
}

// TestFallbackCatAliasPhase3 vérifie la fusion Cat -> Read (phase 3) :
// Cat n'est plus annoncé dans les schémas, mais un appel résiduel
// (DSML ou pseudo-texte) est réécrit en Read.
func TestFallbackCatAliasPhase3(t *testing.T) {
	for _, s := range ToolSchemas() {
		if s.Function.Name == "Cat" {
			t.Fatal("Cat est encore annoncé dans les schémas")
		}
	}
	for name, want := range map[string]string{
		"Cat": "Read", "cat": "Read", "CAT": "Read",
		"Read": "Read", "Grep": "Grep",
	} {
		if got := canonicalToolName(name); got != want {
			t.Errorf("canonicalToolName(%q) = %q, attendu %q", name, got, want)
		}
	}
	// DSML avec l'ancien nom -> Read.
	text := "<||DSML||calls>\n<||DSML||invoke name=\"cat\">\n<||DSML||parameter name=\"file_path\">docs/a.md</||DSML||parameter>\n</||DSML||invoke>\n</||DSML||calls>"
	tcalls, _ := parseFallbackToolCalls(text, fallbackTools())
	if len(tcalls) != 1 {
		t.Fatalf("attendu 1 appel DSML, obtenu %d", len(tcalls))
	}
	if tcalls[0].Function.Name != "Read" {
		t.Errorf("DSML cat -> %q, attendu Read", tcalls[0].Function.Name)
	}
	// Pseudo-texte avec l'ancien nom -> Read.
	tcalls, _ = parseFallbackToolCalls(`Cat "dnsmasq.conf"`, fallbackTools())
	if len(tcalls) != 1 || tcalls[0].Function.Name != "Read" {
		t.Fatalf("pseudo-appel Cat non réécrit en Read : %+v", tcalls)
	}
}
