package rag

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write: %v", err)
	}
}

func richCorpus(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "corpusA", "01-topic", "alpha.md"),
		"# Alpha\n\nLe noyau Linux et les serveurs GPU.\n")
	writeFile(t, filepath.Join(root, "corpusA", "01-topic", "beta.md"),
		"# Beta\n\nQuantization et inference locale.\n")
	writeFile(t, filepath.Join(root, "corpusA", manifestFile), `{
  "corpus": "corpusA",
  "title": "Corpus A",
  "relationship": "base",
  "chunks": [
    {"path": "corpusA/01-topic/alpha.md", "title": "Alpha", "domain": "01-topic",
     "task": "analysis", "actors": ["Linux"], "dates": ["2026-01-02"],
     "keywords": ["kernel", "gpu"], "section": "Alpha"},
    {"path": "corpusA/01-topic/beta.md", "title": "Beta", "domain": "01-topic",
     "task": "analysis", "keywords": ["quantization"], "delta_of": "corpusA"}
  ]
}`)
	return root
}

func TestLoadMissingDirIsInactive(t *testing.T) {
	ix, err := Load(filepath.Join(t.TempDir(), "absent"))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if ix.Stats().Ready {
		t.Fatal("un dossier absent doit donner un index inactif")
	}
	if r := ix.Search(context.Background(), "linux", 5); r.Ready || len(r.Hits) != 0 {
		t.Fatalf("recherche sur index inactif: %+v", r)
	}
}

func TestLoadManifestRich(t *testing.T) {
	ix, err := Load(richCorpus(t))
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	st := ix.Stats()
	if !st.Ready || st.Corpora != 1 || st.Chunks != 2 {
		t.Fatalf("stats inattendues: %+v", st)
	}
	res := ix.Search(context.Background(), "quantization", 5)
	if len(res.Hits) != 1 {
		t.Fatalf("attendu 1 hit, obtenu %d", len(res.Hits))
	}
	h := res.Hits[0]
	if h.Path != "corpusA/01-topic/beta.md" || h.DeltaOf != "corpusA" {
		t.Fatalf("hit inattendu: %+v", h)
	}
	if res := ix.Search(context.Background(), "linux", 5); len(res.Hits) != 1 ||
		res.Hits[0].Actors[0] != "Linux" {
		t.Fatalf("facettes non indexees: %+v", res.Hits)
	}
}

func TestLoadRawFallback(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "perso", "notes.md"),
		"# Mes notes\n\nRecette de cafe filtre et maintenance du VPS.\n")
	ix, err := Load(root)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	st := ix.Stats()
	if !st.Ready || st.Corpora != 1 || st.Chunks != 1 {
		t.Fatalf("stats inattendues: %+v", st)
	}
	res := ix.Search(context.Background(), "cafe", 5)
	if len(res.Hits) != 1 {
		t.Fatalf("attendu 1 hit, obtenu %d", len(res.Hits))
	}
	if res.Hits[0].Title != "Mes notes" {
		t.Fatalf("titre brut inattendu: %q", res.Hits[0].Title)
	}
	if res.Hits[0].Domain != "perso" || res.Hits[0].Corpus != "perso" {
		t.Fatalf("metadonnees brutes inattendues: %+v", res.Hits[0])
	}
	if !strings.Contains(res.Hits[0].Snippet, "cafe") {
		t.Fatalf("snippet sans le terme: %q", res.Hits[0].Snippet)
	}
}

func TestReadLookupAndUnknown(t *testing.T) {
	ix, _ := Load(richCorpus(t))
	full, err := ix.Read("corpusA/01-topic/alpha.md", 1, 2)
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if !strings.HasPrefix(full, "1\t# Alpha") {
		t.Fatalf("lecture inattendue: %q", full)
	}
	if _, err := ix.Read("beta.md", 1, 5); err != nil {
		t.Fatalf("lecture par nom de base: %v", err)
	}
	if _, err := ix.Read("inconnu.md", 1, 5); err == nil {
		t.Fatal("chemin inconnu: erreur attendue")
	}
}

func TestSearchCanceledContext(t *testing.T) {
	ix, _ := Load(richCorpus(t))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	done := make(chan struct{})
	go func() {
		defer close(done)
		_ = ix.Search(ctx, "quantization linux", 5)
	}()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("recherche bloquee sur un contexte annule")
	}
}

func TestManagerBackgroundLoadAndConcurrency(t *testing.T) {
	m := New(richCorpus(t))
	if !m.Enabled() {
		t.Fatal("manager attendu actif")
	}
	m.Start()
	deadline := time.Now().Add(2 * time.Second)
	for !m.Ready() && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}
	if !m.Ready() {
		t.Fatal("index non pret apres Start")
	}
	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if r := m.Search(context.Background(), "quantization", 3); len(r.Hits) == 0 {
				t.Errorf("hit manquant")
			}
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if _, err := m.Reload(); err != nil {
			t.Errorf("Reload: %v", err)
		}
	}()
	wg.Wait()
}

func TestManagerEmptyRootDisabled(t *testing.T) {
	m := New("")
	if m.Enabled() {
		t.Fatal("root vide doit etre desactive")
	}
	if r := m.Search(context.Background(), "x", 5); r.Ready {
		t.Fatal("recherche sur manager desactive")
	}
}

func TestLoadRepoCorpus(t *testing.T) {
	root := filepath.Join("..", "..", "RAG")
	if _, err := os.Stat(filepath.Join(root, "briefing-ia-2026", manifestFile)); err != nil {
		t.Skip("corpus RAG absent")
	}
	ix, err := Load(root)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	st := ix.Stats()
	if st.Corpora != 4 {
		t.Fatalf("attendu 4 corpus, obtenu %d (%v)", st.Corpora, st.Errors)
	}
	if st.Chunks < 400 {
		t.Fatalf("trop peu de chunks: %d", st.Chunks)
	}
	res := ix.Search(context.Background(), "quantization", 5)
	if len(res.Hits) == 0 {
		t.Fatal("aucun hit pour 'quantization'")
	}
	if strings.HasPrefix(res.Hits[0].Snippet, "---") {
		t.Fatalf("front-matter non nettoye: %q", res.Hits[0].Snippet)
	}
}

func BenchmarkSearch(b *testing.B) {
	root := filepath.Join("..", "..", "RAG")
	if _, err := os.Stat(filepath.Join(root, "briefing-ia-2026", manifestFile)); err != nil {
		b.Skip("corpus RAG absent")
	}
	ix, err := Load(root)
	if err != nil || !ix.Stats().Ready {
		b.Skip("index indisponible")
	}
	ctx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if len(ix.Search(ctx, "inference quantization gpu memory", 8).Hits) == 0 {
			b.Fatal("aucun hit")
		}
	}
}
