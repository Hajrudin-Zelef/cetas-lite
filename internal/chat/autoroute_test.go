package chat

import (
	"strings"
	"testing"

	"cetas-lite/internal/alias"
)

// autoTierFor : 3 tiers (flash < standard < elite).
func TestAutoTierForThreeTiers(t *testing.T) {
	modes := []alias.Mode{{ID: "flash"}, {ID: "standard"}, {ID: "elite"}}
	cases := []struct {
		effort string
		text   string
		want   string
	}{
		{"high", "court", "elite"}, // seul high explicite -> le plus cher
		{"medium", "court", "standard"},
		{"low", "court", "flash"},
		{"default", "court", "flash"},                      // défaut + texte court -> le moins cher
		{"", "court", "flash"},                             // valeur inconnue -> comme default
		{"default", strings.Repeat("x ", 300), "standard"}, // défaut + long -> médian
		{"medium", strings.Repeat("x ", 300), "standard"},  // medium reste médian
		{"low", strings.Repeat("x ", 300), "flash"},        // low reste le moins cher
		// Garde-fou coût : même un très long texte ne mène jamais à elite
		// sans effort "high" explicite.
		{"default", strings.Repeat("x ", 2000), "standard"},
	}
	for _, c := range cases {
		if got := autoTierFor(modes, c.effort, c.text); got != c.want {
			t.Errorf("autoTierFor(effort=%q, len=%d) = %q, want %q",
				c.effort, len([]rune(c.text)), got, c.want)
		}
	}
}

// autoTierFor : 1 seul mode (Nano) -> l'effort est sans effet.
func TestAutoTierForSingleMode(t *testing.T) {
	modes := []alias.Mode{{ID: "free"}}
	for _, effort := range []string{"high", "medium", "low", "default", ""} {
		if got := autoTierFor(modes, effort, "court"); got != "free" {
			t.Errorf("autoTierFor(1 mode, effort=%q) = %q, want free", effort, got)
		}
	}
}

// autoTierFor : 2 tiers -> medium reste sur le moins cher (biais coût),
// high seul mène au plus cher.
func TestAutoTierForTwoTiers(t *testing.T) {
	modes := []alias.Mode{{ID: "flash"}, {ID: "standard"}}
	cases := []struct {
		effort string
		text   string
		want   string
	}{
		{"high", "court", "standard"},
		{"medium", "court", "flash"},
		{"low", "court", "flash"},
		{"default", strings.Repeat("x ", 300), "flash"},
	}
	for _, c := range cases {
		if got := autoTierFor(modes, c.effort, c.text); got != c.want {
			t.Errorf("autoTierFor(2 tiers, effort=%q) = %q, want %q", c.effort, got, c.want)
		}
	}
}

// autoTierFor : aucun mode -> vide.
func TestAutoTierForNoModes(t *testing.T) {
	if got := autoTierFor(nil, "high", "court"); got != "" {
		t.Fatalf("autoTierFor(nil) = %q, want \"\"", got)
	}
}
