package search

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// TestSearchRetryOnTotalFailure : si tous les backends echouent, Search
// retente une seule fois avant de remonter l'erreur.
func TestSearchRetryOnTotalFailure(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&calls, 1) == 1 {
			http.Error(w, "boom", http.StatusBadGateway)
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]any{
			"results": []any{map[string]any{"title": "ok", "url": "https://x.fr/ok"}},
		})
	}))
	defer srv.Close()

	s := NewWithConfig(Config{
		Keys:    map[string]string{"tavily": "k"},
		Enabled: map[string]bool{"duckduckgo": false}, // hermetique
		Mode:    "race",
	}, nil)
	s.endpoints["tavily"] = srv.URL

	start := time.Now()
	res := s.Search(context.Background(), "test retry", 5)
	if res.Provider != "tavily" {
		t.Fatalf("retry attendu vers tavily, obtenu provider=%q err=%q", res.Provider, res.Error)
	}
	if len(res.Hits) != 1 || res.Hits[0].Title != "ok" {
		t.Fatalf("hits inattendus: %+v", res.Hits)
	}
	if c := atomic.LoadInt32(&calls); c != 2 {
		t.Fatalf("attendu 2 appels (1 echec + 1 retry), obtenu %d", c)
	}
	if d := time.Since(start); d < searchRetryDelay {
		t.Fatalf("retry trop rapide (%v), le delai %v n'a pas ete respecte", d, searchRetryDelay)
	}
}

// TestSearchNoRetryOnEmptyResults : quand les backends repondent mais sans
// resultats, aucun retry n'est tente.
func TestSearchNoRetryOnEmptyResults(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		_ = json.NewEncoder(w).Encode(map[string]any{"results": []any{}})
	}))
	defer srv.Close()

	s := NewWithConfig(Config{
		Keys:    map[string]string{"tavily": "k"},
		Enabled: map[string]bool{"duckduckgo": false}, // hermetique
		Mode:    "race",
	}, nil)
	s.endpoints["tavily"] = srv.URL

	res := s.Search(context.Background(), "zzz introuvable", 5)
	if c := atomic.LoadInt32(&calls); c != 1 {
		t.Fatalf("aucun resultat: aucun retry attendu, obtenu %d appels", c)
	}
	if !strings.Contains(res.Error, "aucun resultat") {
		t.Fatalf("erreur attendue 'aucun resultat', obtenue %q", res.Error)
	}
}

// TestFetchJinaFallback : quand le fetch direct echoue, le lecteur Jina
// prend le relais (cle Jina configuree).
func TestFetchJinaFallback(t *testing.T) {
	var auth string
	jinaSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte("# Titre\n\nContenu via Jina."))
	}))
	defer jinaSrv.Close()
	// Cible directe sur loopback : le dialer du fetch direct refuse les IP
	// non publiques -> le repli Jina doit prendre le relais.
	directSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer directSrv.Close()

	s := NewWithConfig(Config{Keys: map[string]string{"jina": "k-jina"}}, nil)
	s.endpoints["jina_reader"] = jinaSrv.URL + "/"

	title, md, err := s.Fetch(context.Background(), directSrv.URL+"/page")
	if err != nil {
		t.Fatalf("repli jina attendu, erreur obtenue: %v", err)
	}
	if title != "" {
		t.Fatalf("titre attendu vide via le lecteur, obtenu %q", title)
	}
	if !strings.Contains(md, "Contenu via Jina") {
		t.Fatalf("contenu jina attendu, obtenu %q", md)
	}
	if auth != "Bearer k-jina" {
		t.Fatalf("Authorization jina attendu 'Bearer k-jina', obtenu %q", auth)
	}
}

// TestFetchNoJinaKeyNoFallback : sans cle Jina, un fetch direct en echec
// remonte simplement l'erreur d'origine.
func TestFetchNoJinaKeyNoFallback(t *testing.T) {
	directSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer directSrv.Close()

	s := NewWithConfig(Config{Keys: map[string]string{}}, nil)
	if _, _, err := s.Fetch(context.Background(), directSrv.URL+"/page"); err == nil {
		t.Fatal("sans cle jina, le fetch direct en echec doit remonter l'erreur")
	}
}
