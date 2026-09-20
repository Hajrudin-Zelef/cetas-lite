package catalog

import "testing"

func TestCatalogCounts(t *testing.T) {
	models := Models()
	if len(models) != 66 {
		t.Fatalf("catalogue = %d modeles, want 66", len(models))
	}
	counts := map[string]int{}
	for _, m := range models {
		counts[m.Editeur]++
	}
	want := map[string]int{"deepseek": 3, "opencode": 14, "opencode-go": 8, "openrouter": 41}
	for ed, n := range want {
		if counts[ed] != n {
			t.Errorf("%s = %d, want %d", ed, counts[ed], n)
		}
	}
}

func TestLookup(t *testing.T) {
	m, ok := Lookup("deepseek", "deepseek-v4-pro")
	if !ok {
		t.Fatal("modele introuvable")
	}
	if m.Label != "DeepSeek V4 Pro" || m.InputPer1M != 0.43 || m.OutputPer1M != 1.98 {
		t.Fatalf("modele = %+v", m)
	}
	if _, ok := Lookup("deepseek", "inconnu"); ok {
		t.Fatal("modele inconnu ne doit pas etre trouve")
	}
}

func TestByProvider(t *testing.T) {
	if n := len(ByProvider("deepseek")); n != 3 {
		t.Fatalf("deepseek = %d", n)
	}
	if n := len(ByProvider("inexistant")); n != 0 {
		t.Fatalf("inexistant = %d", n)
	}
}
