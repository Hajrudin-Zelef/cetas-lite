package chat

import (
	"context"
	"fmt"
	"testing"
	"time"

	"cetas-lite/internal/provider"
	"cetas-lite/internal/rag"
)

func TestRelPrefetchStore_PutGet(t *testing.T) {
	s := newRelPrefetchStore()
	want := rag.Result{Hits: []rag.Hit{{Title: "hit"}}}
	s.put("glm 5 3", want)
	got, ok := s.get("glm 5 3")
	if !ok {
		t.Fatal("entree stockee introuvable")
	}
	if len(got.Hits) != 1 || got.Hits[0].Title != "hit" {
		t.Fatalf("resultat degrade : %+v", got)
	}
	if _, ok := s.get("absent"); ok {
		t.Fatal("cle absente => echec attendu")
	}
}

func TestRelPrefetchStore_PutIdempotent(t *testing.T) {
	s := newRelPrefetchStore()
	s.put("a", rag.Result{Hits: []rag.Hit{{Title: "first"}}})
	s.put("a", rag.Result{Hits: []rag.Hit{{Title: "second"}}})
	got, ok := s.get("a")
	if !ok || got.Hits[0].Title != "first" {
		t.Fatalf("premier arrive doit gagner : %+v", got)
	}
	s.mu.Lock()
	n := len(s.order)
	s.mu.Unlock()
	if n != 1 {
		t.Fatalf("ordre corrompu par le doublon : %d", n)
	}
}

func TestRelPrefetchStore_TTLExpired(t *testing.T) {
	s := newRelPrefetchStore()
	s.put("a", rag.Result{Hits: []rag.Hit{{Title: "hit"}}})
	s.mu.Lock()
	s.entries["a"].createdAt = time.Now().Add(-time.Hour)
	s.mu.Unlock()
	if _, ok := s.get("a"); ok {
		t.Fatal("entree expiree => echec attendu")
	}
	s.mu.Lock()
	_, stillThere := s.entries["a"]
	s.mu.Unlock()
	if stillThere {
		t.Fatal("entree expiree non nettoyee")
	}
}

func TestRelPrefetchStore_EvictOldest(t *testing.T) {
	s := newRelPrefetchStore()
	for i := 0; i < relPrefetchMaxEntries; i++ {
		s.put(fmt.Sprintf("k%d", i), rag.Result{Hits: []rag.Hit{{Title: "hit"}}})
	}
	s.put("new", rag.Result{Hits: []rag.Hit{{Title: "hit"}}})
	if _, ok := s.get("k0"); ok {
		t.Fatal("k0 (plus ancien) aurait du etre evince")
	}
	if _, ok := s.get("new"); !ok {
		t.Fatal("nouvelle entree introuvable")
	}
	s.mu.Lock()
	n := len(s.entries)
	s.mu.Unlock()
	if n != relPrefetchMaxEntries {
		t.Fatalf("borne depassee : %d entrees", n)
	}
}

func TestRelPrefetchEnabledFromEnv(t *testing.T) {
	cases := []struct {
		val  string
		want bool
	}{
		{"0", false}, {"false", false}, {"FALSE", false},
		{"no", false}, {"off", false},
		{"", true}, {"1", true}, {"yes", true}, {"true", true},
	}
	for _, tc := range cases {
		t.Setenv("CETAS_LITE_RELATED_PREFETCH", tc.val)
		if got := relPrefetchEnabledFromEnv(); got != tc.want {
			t.Errorf("CETAS_LITE_RELATED_PREFETCH=%q => %v, attendu %v", tc.val, got, tc.want)
		}
	}
}

func TestRelPrefetchState_SingleFlight(t *testing.T) {
	var s relPrefetchState
	if !s.tryAcquire() {
		t.Fatal("premiere acquisition attendue")
	}
	if s.tryAcquire() {
		t.Fatal("seconde acquisition concurrente refusee attendue")
	}
	s.release()
	if !s.tryAcquire() {
		t.Fatal("re-acquisition apres release attendue")
	}
	s.release()
}

func relPrefetchLen(e *Engine) int {
	s := e.relPrefetch.getStore()
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.order)
}

func waitForPrefetch(t *testing.T, what string, cond func() bool) {
	t.Helper()
	waitFor(t, cond, what)
}

func TestMaybeRelatedPrefetch_StoresResults(t *testing.T) {
	old := relPrefetchEnabled
	relPrefetchEnabled = true
	defer func() { relPrefetchEnabled = old }()

	hits := []rag.Hit{{Title: "GLM-5.3 overview", Keywords: []string{"Moonshot AI"}}}
	e := ragEngine(t, fakeRag{ready: true, hits: hits})
	e.maybeRelatedPrefetch(
		TurnInput{Text: "je veux parler de GLM-5.3"},
		rag.Result{Hits: hits, Ready: true},
	)
	// Le job tourne en arriere-plan : attendre le depot.
	waitForPrefetch(t, "sujets pre-recuperes", func() bool { return relPrefetchLen(e) >= 2 })
	for _, key := range []string{"glm", "moonshot ai"} {
		if _, ok := e.relPrefetch.getStore().get(key); !ok {
			t.Errorf("sujet %q non pre-recupere", key)
		}
	}
	// Le job a libere le single-flight.
	if !e.relPrefetch.tryAcquire() {
		t.Error("single-flight non libere apres le job")
	} else {
		e.relPrefetch.release()
	}
}

func TestMaybeRelatedPrefetch_SkippedWhenRAGNotReady(t *testing.T) {
	old := relPrefetchEnabled
	relPrefetchEnabled = true
	defer func() { relPrefetchEnabled = old }()

	e := ragEngine(t, fakeRag{ready: false})
	e.maybeRelatedPrefetch(TurnInput{Text: "GLM-5.3"}, rag.Result{})
	time.Sleep(100 * time.Millisecond)
	if n := relPrefetchLen(e); n != 0 {
		t.Fatalf("RAG inactif => aucun pre-calcul attendu, %d entrees", n)
	}
}

func TestRagHits_ReusesPrefetched(t *testing.T) {
	old := relPrefetchEnabled
	relPrefetchEnabled = true
	defer func() { relPrefetchEnabled = old }()

	e := ragEngine(t, fakeRag{ready: true, hits: []rag.Hit{{Title: "live"}}})
	prefetched := rag.Result{Hits: []rag.Hit{{Title: "prefetched"}}, Ready: true}
	e.relPrefetch.getStore().put(rag.CleanQueryKey("GLM-5.3"), prefetched)

	res := e.ragHits(context.Background(), "GLM-5.3", nil, false)
	if len(res.Hits) != 1 || res.Hits[0].Title != "prefetched" {
		t.Fatalf("reutilisation attendue, obtenu %+v", res.Hits)
	}
}

func TestRagHits_NoReuseWhenBoosted(t *testing.T) {
	old := relPrefetchEnabled
	relPrefetchEnabled = true
	defer func() { relPrefetchEnabled = old }()

	e := ragEngine(t, fakeRag{ready: true, hits: []rag.Hit{{Title: "live"}}})
	history := []provider.Message{{Role: "user", Content: "Kimi K3 est un modele ouvert"}}
	// Requete elliptique : l'enrichissement ajoute des entites + boost.
	// Le depot contient meme la requete enrichie exacte, mais le boost
	// non nul interdit la reutilisation (pre-calcul sans boost).
	qe, ents := rag.EnrichQueryWithHistory("et le prix ?", []string{"Kimi K3 est un modele ouvert"})
	if len(ents) == 0 {
		t.Fatal("precondition : enrichissement attendu")
	}
	e.relPrefetch.getStore().put(rag.CleanQueryKey(qe),
		rag.Result{Hits: []rag.Hit{{Title: "prefetched"}}})

	res := e.ragHits(context.Background(), "et le prix ?", history, false)
	if len(res.Hits) != 1 || res.Hits[0].Title != "live" {
		t.Fatalf("chemin normal attendu (boost non nul), obtenu %+v", res.Hits)
	}
}

func TestRagHits_NoReuseWhenDisabled(t *testing.T) {
	old := relPrefetchEnabled
	relPrefetchEnabled = false
	defer func() { relPrefetchEnabled = old }()

	e := ragEngine(t, fakeRag{ready: true, hits: []rag.Hit{{Title: "live"}}})
	e.relPrefetch.getStore().put(rag.CleanQueryKey("GLM-5.3"),
		rag.Result{Hits: []rag.Hit{{Title: "prefetched"}}})

	res := e.ragHits(context.Background(), "GLM-5.3", nil, false)
	if len(res.Hits) != 1 || res.Hits[0].Title != "live" {
		t.Fatalf("mecanisme desactive => chemin normal, obtenu %+v", res.Hits)
	}
}
