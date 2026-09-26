package rag

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// hybridFixture : manager avec index synthetique + vecteurs factices
// (dims=3). L'embedder factice renvoie toujours e1 : la jambe semantique
// designe le chunk 0 (alpha.md).
func hybridFixture(t *testing.T, embedStatus int) *Manager {
	t.Helper()
	dir := miniCorpus(t)
	m := New(dir)
	if _, err := m.Reload(); err != nil {
		t.Fatal(err)
	}
	ix := m.current()
	if len(ix.chunks) != 3 {
		t.Fatalf("%d chunks, attendus 3", len(ix.chunks))
	}
	model := EmbedModel{Slug: "t", OpenRouterID: "t/m", Dims: 3}
	vs := &VectorStore{Model: model, Dims: 3, Count: 3, Hash: ix.CorpusHash()}
	vs.vecs = []float32{
		1, 0, 0, // alpha
		0, 1, 0, // beta
		0, 0, 1, // gamma
	}
	vs.normalize()
	m.vec.Store(vs)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(embedStatus)
		if embedStatus != http.StatusOK {
			return
		}
		_, _ = w.Write([]byte(`{"data":[{"embedding":[1,0,0],"index":0}]}`))
	}))
	t.Cleanup(srv.Close)
	old := openRouterEmbedURL
	openRouterEmbedURL = srv.URL
	t.Cleanup(func() { openRouterEmbedURL = old })
	m.emb.Store(NewOpenRouterEmbedder("cle-test", model, srv.Client()))
	return m
}

func TestSearchHybridFuses(t *testing.T) {
	m := hybridFixture(t, http.StatusOK)
	res := m.SearchHybrid(context.Background(), "pomme", nil, 3)
	if !res.Hybrid {
		t.Fatal("jambe semantique inactive")
	}
	if len(res.Hits) == 0 || res.Hits[0].Path != "alpha.md" {
		t.Fatalf("hits = %v", res.Hits)
	}
	if res.BM25Top <= 0 {
		t.Fatalf("BM25Top = %f", res.BM25Top)
	}
	if res.SemTop < 0.999 {
		t.Fatalf("SemTop = %f", res.SemTop)
	}
	// Hit.Score garde la semantique BM25.
	if res.Hits[0].Score != res.BM25Top {
		t.Fatalf("Score = %f, BM25Top = %f", res.Hits[0].Score, res.BM25Top)
	}
	// Determinisme.
	res2 := m.SearchHybrid(context.Background(), "pomme", nil, 3)
	for i := range res.Hits {
		if res.Hits[i].Path != res2.Hits[i].Path {
			t.Fatal("resultat non deterministe")
		}
	}
}

func TestSearchHybridSemanticOnly(t *testing.T) {
	m := hybridFixture(t, http.StatusOK)
	// Aucun terme BM25 connu : seule la semantique repond.
	res := m.SearchHybrid(context.Background(), "zzzqqq", nil, 3)
	if !res.Hybrid {
		t.Fatal("jambe semantique inactive")
	}
	if len(res.Hits) == 0 || res.Hits[0].Path != "alpha.md" {
		t.Fatalf("hits = %v", res.Hits)
	}
	if res.BM25Top != 0 {
		t.Fatalf("BM25Top = %f, attendu 0", res.BM25Top)
	}
	if res.SemTop < 0.999 {
		t.Fatalf("SemTop = %f", res.SemTop)
	}
}

func TestSearchHybridEmbedderDown(t *testing.T) {
	m := hybridFixture(t, http.StatusInternalServerError)
	res := m.SearchHybrid(context.Background(), "pomme", nil, 3)
	if res.Hybrid {
		t.Fatal("jambe semantique active malgre l'erreur")
	}
	ref := m.current().SearchBoosted(context.Background(), "pomme", nil, 3)
	if len(res.Hits) != len(ref.Hits) {
		t.Fatalf("%d hits vs %d", len(res.Hits), len(ref.Hits))
	}
	for i := range res.Hits {
		if res.Hits[i].Path != ref.Hits[i].Path || res.Hits[i].Score != ref.Hits[i].Score {
			t.Fatalf("divergence BM25: %+v vs %+v", res.Hits[i], ref.Hits[i])
		}
	}
}

func TestSearchHybridNoVectors(t *testing.T) {
	dir := miniCorpus(t)
	m := New(dir)
	if _, err := m.Reload(); err != nil {
		t.Fatal(err)
	}
	res := m.SearchHybrid(context.Background(), "pomme", nil, 3)
	if res.Hybrid {
		t.Fatal("hybride sans vecteurs")
	}
	ref := m.current().SearchBoosted(context.Background(), "pomme", nil, 3)
	if len(res.Hits) != len(ref.Hits) || res.BM25Top != ref.BM25Top {
		t.Fatalf("divergence BM25: %+v vs %+v", res, ref)
	}
}

func TestSearchHybridBoostKept(t *testing.T) {
	m := hybridFixture(t, http.StatusOK)
	// Le boost d'entite (iteration 6b) survit a la fusion.
	res := m.SearchHybrid(context.Background(), "pomme", map[string]float64{"pomme": 3.0}, 3)
	plain := m.SearchHybrid(context.Background(), "pomme", nil, 3)
	if res.BM25Top <= plain.BM25Top {
		t.Fatalf("boost sans effet: %f <= %f", res.BM25Top, plain.BM25Top)
	}
}

// captureSlog : remplace le logger slog par defaut par un buffer texte,
// pour verifier la ligne d'observabilite rag_hybrid.
func captureSlog(t *testing.T) *strings.Builder {
	t.Helper()
	var buf strings.Builder
	h := slog.NewTextHandler(&buf, nil)
	old := slog.Default()
	slog.SetDefault(slog.New(h))
	t.Cleanup(func() { slog.SetDefault(old) })
	return &buf
}

func TestSearchHybridLogsHybridMode(t *testing.T) {
	buf := captureSlog(t)
	m := hybridFixture(t, http.StatusOK)
	m.SearchHybrid(context.Background(), "pomme", nil, 3)
	out := buf.String()
	for _, want := range []string{"msg=rag_hybrid", "mode=hybrid", "model=t", "sem_top=1"} {
		if !strings.Contains(out, want) {
			t.Fatalf("log sans %q : %q", want, out)
		}
	}
}

func TestSearchHybridLogsFallbackReason(t *testing.T) {
	buf := captureSlog(t)
	m := hybridFixture(t, http.StatusInternalServerError)
	res := m.SearchHybrid(context.Background(), "pomme", nil, 3)
	if res.Hybrid {
		t.Fatal("jambe semantique active malgre l'erreur")
	}
	out := buf.String()
	for _, want := range []string{"msg=rag_hybrid", "mode=bm25_only", "reason=embed_error"} {
		if !strings.Contains(out, want) {
			t.Fatalf("log sans %q : %q", want, out)
		}
	}
}

func TestSearchHybridLogsNoVectors(t *testing.T) {
	buf := captureSlog(t)
	dir := miniCorpus(t)
	m := New(dir)
	if _, err := m.Reload(); err != nil {
		t.Fatal(err)
	}
	m.SearchHybrid(context.Background(), "pomme", nil, 3)
	out := buf.String()
	for _, want := range []string{"msg=rag_hybrid", "mode=bm25_only", "reason=no_vectors"} {
		if !strings.Contains(out, want) {
			t.Fatalf("log sans %q : %q", want, out)
		}
	}
}

func TestSemReasons(t *testing.T) {
	if got := semSkipReason(true, true); got != "no_vectors_no_embedder" {
		t.Fatalf("got %q", got)
	}
	if got := semSkipReason(true, false); got != "no_vectors" {
		t.Fatalf("got %q", got)
	}
	if got := semSkipReason(false, true); got != "no_embedder" {
		t.Fatalf("got %q", got)
	}
	if got := semErrReason(nil); got != "no_sem_hits" {
		t.Fatalf("got %q", got)
	}
	if got := semErrReason(context.DeadlineExceeded); got != "embed_timeout" {
		t.Fatalf("got %q", got)
	}
	if got := semErrReason(errors.New("boom")); got != "embed_error" {
		t.Fatalf("got %q", got)
	}
	if got := round3f(0.39421); got != 0.394 {
		t.Fatalf("got %f", got)
	}
}

// rerankFixture : manager avec le reranker desktop branche. Le faux serveur
// /rerank inverse l'ordre des candidats RRF.
func rerankFixture(t *testing.T) *Manager {
	t.Helper()
	m := hybridFixture(t, http.StatusOK)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rerank" {
			http.NotFound(w, r)
			return
		}
		var body rerankRequest
		_ = json.NewDecoder(r.Body).Decode(&body)
		out := make([]rerankHit, 0, len(body.Texts))
		for i := len(body.Texts) - 1; i >= 0; i-- {
			out = append(out, rerankHit{Index: i, Score: float64(i) + 0.5})
		}
		_ = json.NewEncoder(w).Encode(out)
	}))
	t.Cleanup(srv.Close)
	m.SetReranker(NewReranker(srv.URL, "cle-desktop", 20, 4000, srv.Client()))
	return m
}

func TestSearchHybridAppliesRerank(t *testing.T) {
	plain := hybridFixture(t, http.StatusOK)
	reranked := rerankFixture(t)

	a := plain.SearchHybrid(context.Background(), "pomme", nil, 3)
	b := reranked.SearchHybrid(context.Background(), "pomme", nil, 3)
	if len(a.Hits) < 2 || len(b.Hits) < 2 {
		t.Fatalf("hits insuffisants: %d / %d", len(a.Hits), len(b.Hits))
	}
	// Sans rerank, alpha.md (semantique) est premier ; avec, l'ordre est
	// inverse par le faux serveur.
	if a.Hits[0].Path != "alpha.md" {
		t.Fatalf("ordre RRF inattendu: %s", a.Hits[0].Path)
	}
	if b.Hits[0].Path == a.Hits[0].Path {
		t.Fatalf("rerank non applique: premier hit inchange (%s)", b.Hits[0].Path)
	}
}

func TestRerankFailOpenKeepsRRF(t *testing.T) {
	m := hybridFixture(t, http.StatusOK)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	t.Cleanup(srv.Close)
	m.SetReranker(NewReranker(srv.URL, "cle-desktop", 20, 4000, srv.Client()))

	plain := hybridFixture(t, http.StatusOK)
	got := m.SearchHybrid(context.Background(), "pomme", nil, 3)
	want := plain.SearchHybrid(context.Background(), "pomme", nil, 3)
	if len(got.Hits) != len(want.Hits) {
		t.Fatalf("hits = %d, attendus %d", len(got.Hits), len(want.Hits))
	}
	for i := range want.Hits {
		if got.Hits[i].Path != want.Hits[i].Path {
			t.Fatalf("repli RRF: %s != %s (position %d)", got.Hits[i].Path, want.Hits[i].Path, i)
		}
	}
}

func TestRerankErrReason(t *testing.T) {
	if rerankErrReason(context.DeadlineExceeded) != "rerank_timeout" {
		t.Fatal("timeout attendu")
	}
	if rerankErrReason(errors.New("boom")) != "rerank_error" {
		t.Fatal("erreur attendue")
	}
}
