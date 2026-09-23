package chat

import (
	"path/filepath"
	"testing"

	"cetas-lite/internal/cache"
	"cetas-lite/internal/provider"
)

// engineReq imite la requête construite par le tour principal.
func engineReq() provider.Request {
	return provider.Request{
		Model:       "spark",
		Messages:    []provider.Message{{Role: "user", Content: "bonjour"}},
		Temperature: 0.7,
	}
}

func TestCacheKeyForGate(t *testing.T) {
	e := newEngine(t, &fakeProvider{id: "p"}, nil)
	// Température du tour principal (0.7) : non éligible => clé vide.
	if k := e.cacheKeyFor("p", engineReq()); k != "" {
		t.Fatalf("température 0.7 : clé = %q, attendu vide", k)
	}
	// Température 0 : éligible => clé non vide.
	r0 := engineReq()
	r0.Temperature = 0
	if k := e.cacheKeyFor("p", r0); k == "" {
		t.Fatal("température 0 : clé vide inattendue")
	}
	// Outils ou extra : non éligible même à température 0.
	r0.Tools = []provider.Tool{{Type: "function"}}
	if k := e.cacheKeyFor("p", r0); k != "" {
		t.Fatalf("avec outils : clé = %q, attendu vide", k)
	}
	r0.Tools = nil
	r0.Extra = map[string]any{"web": true}
	if k := e.cacheKeyFor("p", r0); k != "" {
		t.Fatalf("avec extra : clé = %q, attendu vide", k)
	}
}

func TestCacheKeyForDeterminism(t *testing.T) {
	e := newEngine(t, &fakeProvider{id: "p"}, nil)
	r := engineReq()
	r.Temperature = 0
	k1 := e.cacheKeyFor("p", r)
	k2 := e.cacheKeyFor("p", r)
	if k1 == "" || k1 != k2 {
		t.Fatalf("clé non déterministe : %q vs %q", k1, k2)
	}
	r.Model = "autre"
	if k := e.cacheKeyFor("p", r); k == k1 {
		t.Fatal("modèle différent : même clé")
	}
}

func testCache(t *testing.T) *cache.Cache {
	t.Helper()
	c, err := cache.Open(filepath.Join(t.TempDir(), "cache.db"), 100)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = c.Close() })
	return c
}

func TestCacheStoreAndLookup(t *testing.T) {
	cc := testCache(t)
	e := newEngine(t, &fakeProvider{id: "p"}, nil)
	e.SetCache(cc)
	r := engineReq()
	r.Temperature = 0
	key := e.cacheKeyFor("p", r)
	e.cacheStore(key, provider.Response{
		Content: "réponse", Reasoning: "r",
		Usage: provider.Usage{PromptTokens: 5, CompletionTokens: 2, TotalTokens: 7},
	})
	ce, ok := cc.Get(key)
	if !ok {
		t.Fatal("entrée non retrouvée après cacheStore")
	}
	if ce.Content != "réponse" || ce.Usage.CompletionTokens != 2 {
		t.Fatalf("entrée inattendue : %+v", ce)
	}
	// Clé vide ou cache absent : no-op, pas de panique.
	e.cacheStore("", provider.Response{Content: "x"})
	e2 := newEngine(t, &fakeProvider{id: "p"}, nil)
	e2.cacheStore(key, provider.Response{Content: "x"})
}

func TestCacheStats(t *testing.T) {
	e := newEngine(t, &fakeProvider{id: "p"}, nil)
	if enabled, _, _, _ := e.CacheStats(); enabled {
		t.Fatal("sans cache : enabled inattendu")
	}
	cc := testCache(t)
	e.SetCache(cc)
	cc.Get("absente")
	enabled, hits, misses, entries := e.CacheStats()
	if !enabled || hits != 0 || misses != 1 || entries != 0 {
		t.Fatalf("stats = %v/%d/%d/%d", enabled, hits, misses, entries)
	}
	if e.Cache() != cc {
		t.Fatal("Cache() ne rend pas le cache branché")
	}
}
