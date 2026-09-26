package rag

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// fakeDesktop : serveur cetasrag factice : /embed renvoie des vecteurs de
// dims dims, /rerank renvoie l'ordre inverse des candidats.
func fakeDesktop(t *testing.T, dims int) (*httptest.Server, *int) {
	t.Helper()
	embedCalls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer cle-desktop" {
			t.Errorf("Authorization = %q", got)
		}
		switch r.URL.Path {
		case "/embed":
			embedCalls++
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("requete /embed illisible: %v", err)
			}
			inputs, _ := body["inputs"].([]any)
			out := make([][]float32, len(inputs))
			for i := range inputs {
				v := make([]float32, dims)
				v[i%dims] = 1
				out[i] = v
			}
			_ = json.NewEncoder(w).Encode(out)
		case "/rerank":
			var body rerankRequest
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("requete /rerank illisible: %v", err)
			}
			// Ordre inverse : le dernier candidat devient le premier.
			out := make([]rerankHit, 0, len(body.Texts))
			for i := len(body.Texts) - 1; i >= 0; i-- {
				out = append(out, rerankHit{Index: i, Score: float64(i) + 0.5})
			}
			_ = json.NewEncoder(w).Encode(out)
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(srv.Close)
	return srv, &embedCalls
}

func TestDesktopEmbedFormatAndDims(t *testing.T) {
	srv, calls := fakeDesktop(t, 1024)
	model := EmbedModel{Slug: "BAAI/bge-m3", Dims: 1024}
	emb := NewDesktopEmbedder(srv.URL, "cle-desktop", model, srv.Client())
	out, err := emb.Embed(context.Background(), []string{"a", "b", "c"}, false)
	if err != nil {
		t.Fatalf("Embed: %v", err)
	}
	if len(out) != 3 || len(out[0]) != 1024 {
		t.Fatalf("vecteurs = %d x %d, attendus 3 x 1024", len(out), len(out[0]))
	}
	if *calls != 1 {
		t.Fatalf("%d appels /embed, attendu 1", *calls)
	}
}

func TestDesktopEmbedDimMismatchRejected(t *testing.T) {
	srv, _ := fakeDesktop(t, 512)
	model := EmbedModel{Slug: "BAAI/bge-m3", Dims: 1024}
	emb := NewDesktopEmbedder(srv.URL, "cle-desktop", model, srv.Client())
	if _, err := emb.Embed(context.Background(), []string{"a"}, false); err == nil {
		t.Fatal("dim mismatch non detecte (512 vs 1024)")
	}
}

func TestDesktopEmbedNoKeyFailOpen(t *testing.T) {
	model := EmbedModel{Slug: "BAAI/bge-m3", Dims: 1024}
	emb := NewDesktopEmbedder("http://x", "", model, nil)
	if _, err := emb.Embed(context.Background(), []string{"a"}, false); err == nil {
		t.Fatal("cle vide doit rendre l'embedder indisponible")
	}
}

func TestDesktopEmbed401(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))
	t.Cleanup(srv.Close)
	model := EmbedModel{Slug: "BAAI/bge-m3", Dims: 1024}
	emb := NewDesktopEmbedder(srv.URL, "mauvaise", model, srv.Client())
	_, err := emb.Embed(context.Background(), []string{"a"}, false)
	if err == nil || !contains(err.Error(), "401") {
		t.Fatalf("401 attendu, obtenu: %v", err)
	}
}

func TestRerankerReorders(t *testing.T) {
	srv, _ := fakeDesktop(t, 1024)
	rrk := NewReranker(srv.URL, "cle-desktop", 20, 4000, srv.Client())
	hits, err := rrk.Rerank(context.Background(), "q", []string{"c0", "c1", "c2"})
	if err != nil {
		t.Fatalf("Rerank: %v", err)
	}
	if len(hits) != 3 || hits[0].Index != 2 || hits[2].Index != 0 {
		t.Fatalf("ordre rerank inattendu: %+v", hits)
	}
}

func TestRerankerFailOpenOnError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	t.Cleanup(srv.Close)
	rrk := NewReranker(srv.URL, "cle-desktop", 20, 4000, srv.Client())
	if _, err := rrk.Rerank(context.Background(), "q", []string{"a"}); err == nil {
		t.Fatal("5xx doit remonter une erreur (repli RRF cote appelant)")
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || indexOf(s, sub) >= 0)
}

func indexOf(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
