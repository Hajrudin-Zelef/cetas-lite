package search

import (
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

func TestNormalizeHitAliases(t *testing.T) {
	h, ok := normalizeHit(map[string]any{
		"headline": "Titre", "link": "https://exemple.fr/a", "snippet": "resume", "engine": "brave",
	})
	if !ok {
		t.Fatal("hit attendu")
	}
	if h.Title != "Titre" || h.URL != "https://exemple.fr/a" || h.Description != "resume" || h.Source != "brave" {
		t.Fatalf("normalisation inattendue: %+v", h)
	}
	if _, ok := normalizeHit(map[string]any{"foo": "bar"}); ok {
		t.Fatal("hit sans titre ni url doit etre rejete")
	}
}

func TestNormalizeHitStripsControl(t *testing.T) {
	h, ok := normalizeHit(map[string]any{"title": "  a\x00b\x07  ", "url": "https://x.fr"})
	if !ok {
		t.Fatal("hit attendu")
	}
	if h.Title != "ab" {
		t.Fatalf("caracteres de controle non retires: %q", h.Title)
	}
}

func TestSearchNoKeys(t *testing.T) {
	s := New(nil, nil)
	res := s.Search(context.Background(), "test", 5)
	if len(res.Hits) != 0 || res.Provider != "" || !strings.Contains(res.Error, "keys set") {
		t.Fatalf("erreur de configuration attendue: %+v", res)
	}
}

func TestSearchCascadePrefersEarlier(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/tavily":
			_ = json.NewEncoder(w).Encode(map[string]any{"results": []any{
				map[string]any{"title": "T", "url": "https://t.example", "content": "c"},
			}})
		case "/exa":
			_ = json.NewEncoder(w).Encode(map[string]any{"results": []any{
				map[string]any{"title": "E", "url": "https://e.example", "text": "x"},
			}})
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	s := New(map[string]string{"tavily": "k1", "exa": "k2"}, srv.Client())
	s.endpoints["tavily"] = srv.URL + "/tavily"
	s.endpoints["exa"] = srv.URL + "/exa"

	res := s.Search(context.Background(), "hello", 5)
	if res.Provider != "tavily" {
		t.Fatalf("provider prioritaire attendu: %+v", res)
	}
	if len(res.Hits) != 1 || res.Hits[0].Title != "T" {
		t.Fatalf("hits inattendus: %+v", res)
	}
}

func TestSearchFallbackWhenFirstEmpty(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/tavily":
			_ = json.NewEncoder(w).Encode(map[string]any{"results": []any{}})
		case "/exa":
			_ = json.NewEncoder(w).Encode(map[string]any{"results": []any{
				map[string]any{"title": "E", "url": "https://e.example", "text": "x"},
			}})
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	s := New(map[string]string{"tavily": "k1", "exa": "k2"}, srv.Client())
	s.endpoints["tavily"] = srv.URL + "/tavily"
	s.endpoints["exa"] = srv.URL + "/exa"

	res := s.Search(context.Background(), "hello", 5)
	if res.Provider != "exa" || len(res.Hits) != 1 {
		t.Fatalf("repli vers exa attendu: %+v", res)
	}
}

func TestSearchCache(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"results": []any{
			map[string]any{"title": "T", "url": "https://t.example"},
		}})
	}))
	defer srv.Close()

	s := New(map[string]string{"tavily": "k1"}, srv.Client())
	s.endpoints["tavily"] = srv.URL + "/search"

	_ = s.Search(context.Background(), "meme requete", 5)
	_ = s.Search(context.Background(), "meme requete", 5)
	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Fatalf("cache attendu (1 appel), obtenu %d", got)
	}
}

func TestIsPublicIP(t *testing.T) {
	cases := map[string]bool{
		"127.0.0.1":     false,
		"10.1.2.3":      false,
		"172.16.0.1":    false,
		"192.168.1.5":   false,
		"169.254.1.1":   false,
		"100.64.0.1":    false,
		"0.0.0.0":       false,
		"::1":           false,
		"fc00::1":       false,
		"fe80::1":       false,
		"8.8.8.8":       true,
		"93.184.216.34": true,
	}
	for raw, want := range cases {
		if got := isPublicIP(net.ParseIP(raw)); got != want {
			t.Fatalf("isPublicIP(%s) = %v, attendu %v", raw, got, want)
		}
	}
}

func TestValidateFetchURL(t *testing.T) {
	if _, err := validateFetchURL("file:///etc/passwd"); err == nil {
		t.Fatal("schema file refuse attendu")
	}
	if _, err := validateFetchURL("http://"); err == nil {
		t.Fatal("hote manquant refuse attendu")
	}
	if _, err := validateFetchURL("https://exemple.fr/a"); err != nil {
		t.Fatalf("URL valide refusee: %v", err)
	}
}

func TestFetchBlocksPrivate(t *testing.T) {
	s := New(nil, nil)
	for _, raw := range []string{"http://127.0.0.1:9/", "http://10.0.0.1/", "http://[::1]/"} {
		if _, _, err := s.Fetch(context.Background(), raw); err == nil {
			t.Fatalf("acces reseau prive autorise: %s", raw)
		}
	}
}

func TestFetchExtractsReadable(t *testing.T) {
	page := `<!doctype html><html><head><title>Le grand titre</title></head><body>
	<nav><a href="/">Accueil</a></nav>
	<article>
	<h1>Le grand titre</h1>
	<p>Premier paragraphe avec suffisamment de texte pour etre considere comme du contenu
	principal par l'algorithme de lisibilite. Il contient plusieurs phrases utiles.</p>
	<p>Deuxieme paragraphe qui developpe le sujet et ajoute du corps a l'article afin que
	l'extraction soit fiable et produise un resultat non vide.</p>
	<p>Troisieme paragraphe de conclusion, avec encore quelques mots pour la route et un
	peu de substance supplementaire.</p>
	</article>
	<footer>pied de page</footer>
	<script>var x = 1;</script>
	</body></html>`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = io.WriteString(w, page)
	}))
	defer srv.Close()

	s := New(nil, nil)
	s.fetchClient = srv.Client()
	title, md, err := s.Fetch(context.Background(), srv.URL)
	if err != nil {
		t.Fatalf("fetch: %v", err)
	}
	if !strings.Contains(title, "grand titre") {
		t.Fatalf("titre inattendu: %q", title)
	}
	if !strings.Contains(md, "Premier paragraphe") {
		t.Fatalf("contenu extrait inattendu: %q", md)
	}
	if !strings.HasPrefix(md, "#") && !strings.Contains(md, "**") {
		// le contenu doit rester lisible ; on tolere le repli texte
		t.Logf("markdown sans balise de titre detectee: %q", md)
	}
}
