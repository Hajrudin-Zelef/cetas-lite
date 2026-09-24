package rag

import (
	"context"
	"net/http"
	"net/http/httptest"
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
