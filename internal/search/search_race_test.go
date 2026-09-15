package search

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const ddgFixture = `
<html><body>
<div class="result">
<a rel="nofollow" class="result__a" href="//duckduckgo.com/l/?uddg=https%3A%2F%2Fexemple.fr%2Farticle&amp;rut=abc">Titre <b>Article</b></a>
<a class="result__snippet" href="https://exemple.fr/article">Un r&eacute;sum&eacute; du contenu.</a>
</div>
<div class="result">
<a rel="nofollow" class="result__a" href="https://direct.fr/page">Page Directe</a>
</div>
</body></html>`

func TestParseDDGHTML(t *testing.T) {
	hits := parseDDGHTML(ddgFixture, 10)
	if len(hits) != 2 {
		t.Fatalf("attendu 2 hits, obtenu %d", len(hits))
	}
	if hits[0].URL != "https://exemple.fr/article" {
		t.Fatalf("redirecteur uddg non resolu: %q", hits[0].URL)
	}
	if hits[0].Title != "Titre Article" {
		t.Fatalf("titre mal nettoye: %q", hits[0].Title)
	}
	if hits[0].Description != "Un résumé du contenu." {
		t.Fatalf("snippet mal decode: %q", hits[0].Description)
	}
	if hits[0].Source != "duckduckgo" {
		t.Fatalf("source attendue duckduckgo: %q", hits[0].Source)
	}
	if hits[1].URL != "https://direct.fr/page" || hits[1].Title != "Page Directe" {
		t.Fatalf("hit direct inattendu: %+v", hits[1])
	}
	// Borne max.
	if got := parseDDGHTML(ddgFixture, 1); len(got) != 1 {
		t.Fatalf("borne max ignoree: %d", len(got))
	}
}

// mockJSONServer simule un provider JSON (tavily/exa/brave/jina) avec latence.
func mockJSONServer(t *testing.T, latency time.Duration, title string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(latency)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"results": []any{map[string]any{"title": title, "url": "https://x.fr/" + title}},
		})
	}))
}

func TestSearchRaceVsPriority(t *testing.T) {
	slow := mockJSONServer(t, 400*time.Millisecond, "lent")
	defer slow.Close()
	fast := mockJSONServer(t, 20*time.Millisecond, "rapide")
	defer fast.Close()

	newS := func(mode string) *Searcher {
		s := NewWithConfig(Config{
			Keys:    map[string]string{"tavily": "k", "exa": "k2"},
			Enabled: map[string]bool{"duckduckgo": false}, // hermetique : pas de reseau reel
			Mode:    mode,
		}, nil)
		s.endpoints["tavily"] = slow.URL
		s.endpoints["exa"] = fast.URL
		return s
	}

	// Mode race : le plus rapide gagne (exa), sans attendre tavily.
	start := time.Now()
	res := newS("race").Search(context.Background(), "test", 5)
	if res.Provider != "exa" {
		t.Fatalf("mode race: attendu exa, obtenu %q", res.Provider)
	}
	if d := time.Since(start); d > 350*time.Millisecond {
		t.Fatalf("mode race trop lent: %v (aurait du repondre des la reponse rapide)", d)
	}

	// Mode priority : tavily (plus prioritaire) gagne meme si plus lent.
	res = newS("priority").Search(context.Background(), "test", 5)
	if res.Provider != "tavily" {
		t.Fatalf("mode priority: attendu tavily, obtenu %q", res.Provider)
	}
	if len(res.Hits) != 1 || res.Hits[0].Title != "lent" {
		t.Fatalf("hits inattendus: %+v", res.Hits)
	}
}

func TestSearchEnabledFlag(t *testing.T) {
	srv := mockJSONServer(t, 0, "ok")
	defer srv.Close()
	s := NewWithConfig(Config{
		Keys:    map[string]string{"tavily": "k"},
		Enabled: map[string]bool{"tavily": false, "duckduckgo": false}, // hermetique : aucun reseau reel
		Mode:    "race",
	}, nil)
	s.endpoints["tavily"] = srv.URL
	res := s.Search(context.Background(), "test", 5)
	// Tavily desactive et DuckDuckGo coupe : aucun provider actif.
	if res.Provider == "tavily" {
		t.Fatal("tavily desactive mais utilise")
	}
	if res.Provider != "none" {
		t.Fatalf("aucun provider ne devrait repondre, obtenu %q", res.Provider)
	}
}

func TestSearchDuckDuckGoKeyless(t *testing.T) {
	s := NewWithConfig(Config{Keys: map[string]string{}, Mode: "race"}, nil)
	active := s.activeProviders()
	found := false
	for _, id := range active {
		if id == "duckduckgo" {
			found = true
		}
	}
	if !found {
		t.Fatalf("duckduckgo devrait etre actif sans cle, actifs=%v", active)
	}
}

func TestSetConfigHotReload(t *testing.T) {
	s := New(map[string]string{}, nil)
	if len(s.activeProviders()) != 1 { // duckduckgo seul
		t.Fatalf("actifs inattendus: %v", s.activeProviders())
	}
	s.SetConfig(Config{Keys: map[string]string{"brave": "k2"}, Mode: "race"})
	active := s.activeProviders()
	if len(active) != 2 || active[0] != "brave" || active[1] != "duckduckgo" {
		t.Fatalf("config chaude non appliquee: %v", active)
	}
	if s.getMode() != "race" {
		t.Fatalf("mode non applique: %q", s.getMode())
	}
}

func TestDefaultOrder(t *testing.T) {
	want := []string{"brave", "tavily", "jina", "exa", "duckduckgo"}
	if len(DefaultOrder) != len(want) {
		t.Fatalf("ordre par defaut inattendu: %v", DefaultOrder)
	}
	for i := range want {
		if DefaultOrder[i] != want[i] {
			t.Fatalf("ordre par defaut inattendu: %v (attendu %v)", DefaultOrder, want)
		}
	}
}
