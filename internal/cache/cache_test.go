package cache

import (
	"path/filepath"
	"strings"
	"testing"
)

func openTest(t *testing.T, maxEntries int) *Cache {
	t.Helper()
	c, err := Open(filepath.Join(t.TempDir(), "cache.db"), maxEntries)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	t.Cleanup(func() { _ = c.Close() })
	return c
}

func TestEligible(t *testing.T) {
	if !Eligible(0, 0, false) {
		t.Fatal("température 0, sans outils, sans extra : éligible attendu")
	}
	for _, tc := range []struct {
		name  string
		temp  float64
		tools int
		extra bool
	}{
		{"température 0.7", 0.7, 0, false},
		{"température 0.2", 0.2, 0, false},
		{"avec outils", 0, 2, false},
		{"avec extra", 0, 0, true},
	} {
		if Eligible(tc.temp, tc.tools, tc.extra) {
			t.Errorf("%s : inéligible attendu", tc.name)
		}
	}
}

func TestKeyOf(t *testing.T) {
	raw := []byte(`{"model":"m","messages":[]}`)
	k := KeyOf(raw, 0, 0, false)
	if len(k) != 64 {
		t.Fatalf("clé = %q, attendu 64 caractères hex", k)
	}
	if KeyOf(raw, 0, 0, false) != k {
		t.Fatal("clé non déterministe")
	}
	if k2 := KeyOf([]byte(`{"model":"autre"}`), 0, 0, false); k2 == k {
		t.Fatal("requêtes différentes : clés identiques")
	}
	// Porte : toute requête non éligible => clé vide.
	for _, tc := range [][3]any{{0.7, 0, false}, {0.0, 1, false}, {0.0, 0, true}} {
		if k := KeyOf(raw, tc[0].(float64), tc[1].(int), tc[2].(bool)); k != "" {
			t.Errorf("KeyOf(%v) = %q, attendu vide", tc, k)
		}
	}
}

func TestSetGetRoundtrip(t *testing.T) {
	c := openTest(t, 100)
	in := Entry{Content: "bonjour", Reasoning: "raisonnement", Usage: Usage{PromptTokens: 10, CompletionTokens: 3, TotalTokens: 13}}
	if err := c.Set("k1", in); err != nil {
		t.Fatalf("Set: %v", err)
	}
	got, ok := c.Get("k1")
	if !ok {
		t.Fatal("Get: entrée absente")
	}
	if got != in {
		t.Fatalf("Get = %+v, attendu %+v", got, in)
	}
	if _, ok := c.Get("absente"); ok {
		t.Fatal("Get sur clé absente : hit inattendu")
	}
	// Écrasement.
	in2 := Entry{Content: "v2"}
	if err := c.Set("k1", in2); err != nil {
		t.Fatalf("Set overwrite: %v", err)
	}
	if got, _ := c.Get("k1"); got.Content != "v2" {
		t.Fatalf("écrasement : contenu = %q", got.Content)
	}
}

func TestStatsHitsMisses(t *testing.T) {
	c := openTest(t, 100)
	c.Get("nope") // miss
	_ = c.Set("k", Entry{Content: "x"})
	c.Get("k") // hit
	c.Get("k") // hit
	h, m, n := c.Stats()
	if h != 2 || m != 1 || n != 1 {
		t.Fatalf("hits=%d misses=%d entries=%d, attendu 2/1/1", h, m, n)
	}
}

func TestTrim(t *testing.T) {
	c := openTest(t, 2)
	keys := []string{"ka", "kb", "kc"}
	for _, k := range keys {
		if err := c.Set(k, Entry{Content: k}); err != nil {
			t.Fatalf("Set %s: %v", k, err)
		}
	}
	_, _, n := c.Stats()
	if n != 2 {
		t.Fatalf("entries=%d, attendu 2 après élagage", n)
	}
	// created_at à la seconde près : en cas d'égalité, la plus petite clé
	// part (ORDER BY created_at DESC, key DESC).
	if _, ok := c.Get("ka"); ok {
		t.Fatal("ka aurait dû être élaguée")
	}
}

func TestOpenCreatesParent(t *testing.T) {
	p := filepath.Join(t.TempDir(), "sous", "rep", "cache.db")
	c, err := Open(p, 10)
	if err != nil {
		t.Fatalf("Open avec parents manquants: %v", err)
	}
	_ = c.Close()
}

func TestKeyIsHex(t *testing.T) {
	k := KeyOf([]byte("x"), 0, 0, false)
	if strings.Trim(k, "0123456789abcdef") != "" {
		t.Fatalf("clé non hexadécimale : %q", k)
	}
}
