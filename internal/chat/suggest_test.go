package chat

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"cetas-lite/internal/rag"
)

func TestLoadSuggestions(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "suggested-questions.yaml")
	content := `# commentaire
- question: "Qu'est-ce que Kimi K3 ?"
  corpus: ai-industry-kb-2026-wave6
  pinned: true
- question: 'vLLM ou SGLang ?'
  corpus: etape4-tracka-vllm-sglang
- question: Question sans guillemets
  corpus: tools-platforms-2026
`
	if err := os.WriteFile(p, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	suggs, err := LoadSuggestions(p)
	if err != nil {
		t.Fatalf("LoadSuggestions: %v", err)
	}
	if len(suggs) != 3 {
		t.Fatalf("%d suggestions, attendu 3", len(suggs))
	}
	if !suggs[0].Pinned || suggs[1].Pinned {
		t.Fatalf("pinned mal parsé : %+v", suggs)
	}
	if suggs[2].Question != "Question sans guillemets" {
		t.Fatalf("question nue mal parsée : %q", suggs[2].Question)
	}

	// Fichier absent => fail-open (nil, nil).
	if s, err := LoadSuggestions(filepath.Join(dir, "absent.yaml")); err != nil || s != nil {
		t.Fatalf("fichier absent : s=%v err=%v", s, err)
	}

	// Entrée invalide (corpus manquant) => erreur.
	bad := filepath.Join(dir, "bad.yaml")
	os.WriteFile(bad, []byte("- question: \"Q\"\n"), 0o600)
	if _, err := LoadSuggestions(bad); err == nil {
		t.Fatal("corpus manquant : erreur attendue")
	}

	// Clé inconnue => erreur.
	bad2 := filepath.Join(dir, "bad2.yaml")
	os.WriteFile(bad2, []byte("- question: \"Q\"\n  corpus: c\n  nope: 1\n"), 0o600)
	if _, err := LoadSuggestions(bad2); err == nil {
		t.Fatal("clé inconnue : erreur attendue")
	}
}

// TestShippedSuggestions : le YAML livré doit être valide et conséquent.
func TestShippedSuggestions(t *testing.T) {
	suggs, err := LoadSuggestions("../../RAG/suggested-questions.yaml")
	if err != nil {
		t.Fatalf("YAML livré illisible : %v", err)
	}
	if len(suggs) < 50 {
		t.Fatalf("%d questions, attendu 50+", len(suggs))
	}
	seen := map[string]bool{}
	corpora := map[string]bool{}
	pinned := 0
	for _, s := range suggs {
		if strings.TrimSpace(s.Question) == "" || strings.TrimSpace(s.Corpus) == "" {
			t.Fatalf("entrée incomplète : %+v", s)
		}
		k := strings.ToLower(strings.TrimSpace(s.Question))
		if seen[k] {
			t.Fatalf("doublon : %q", s.Question)
		}
		seen[k] = true
		corpora[s.Corpus] = true
		if s.Pinned {
			pinned++
		}
	}
	if pinned == 0 {
		t.Fatal("aucune question épinglée")
	}
	t.Logf("%d questions, %d corpus, %d épinglées", len(suggs), len(corpora), pinned)
}

func TestIsGreeting(t *testing.T) {
	ok := []string{"salut", "Salut !", "  bonjour ", "BONSOIR.", "hello", "Hi?", "coucou…", "yo !"}
	for _, s := range ok {
		if !IsGreeting(s) {
			t.Errorf("IsGreeting(%q) = false, attendu true", s)
		}
	}
	ko := []string{"", "salut, ça va ?", "bonjour à tous", "salu", "salutation", "ok", "merci"}
	for _, s := range ko {
		if IsGreeting(s) {
			t.Errorf("IsGreeting(%q) = true, attendu false", s)
		}
	}
}

func TestSuggestionsSetGet(t *testing.T) {
	e := ragEngine(t, fakeRag{})
	if got := e.Suggestions(); len(got) != 0 {
		t.Fatalf("pool initial non vide : %d", len(got))
	}
	in := []Suggestion{{Question: "Q1", Corpus: "c1", Pinned: true}}
	e.SetSuggestions(in)
	got := e.Suggestions()
	if len(got) != 1 || got[0].Question != "Q1" || !got[0].Pinned {
		t.Fatalf("pool inattendu : %+v", got)
	}
	// Copie : modifier le retour ne touche pas l'engine.
	got[0].Question = "MUT"
	if e.Suggestions()[0].Question != "Q1" {
		t.Fatal("Suggestions() ne retourne pas une copie")
	}
}

// corpusRag enregistre le corpus demandé à SearchCorpus.
type corpusRag struct {
	fakeRag
	gotCorpus string
}

func (f *corpusRag) SearchCorpus(_ context.Context, _ string, corpus string, limit int) rag.Result {
	f.gotCorpus = corpus
	return f.fakeRag.Search(context.Background(), "", limit)
}

func TestRagHitsCorpusForwardsCorpus(t *testing.T) {
	fr := &corpusRag{fakeRag: fakeRag{ready: true, hits: []rag.Hit{
		{Title: "T", Path: "p", Excerpt: "e", Score: 9.0},
	}}}
	e := ragEngine(t, fr)
	res := e.ragHitsCorpus(context.Background(), "kimi k3", "wave6", ragContextHits)
	if fr.gotCorpus != "wave6" {
		t.Fatalf("corpus transmis = %q, attendu wave6", fr.gotCorpus)
	}
	if len(res.Hits) != 1 {
		t.Fatalf("%d hits, attendu 1", len(res.Hits))
	}
	// RAG inactif => vide.
	e2 := ragEngine(t, &corpusRag{fakeRag: fakeRag{ready: false}})
	if res := e2.ragHitsCorpus(context.Background(), "q", "c", ragContextHits); len(res.Hits) != 0 {
		t.Fatal("RAG inactif : hits inattendus")
	}
}

func TestRagContextForcedIgnoresScoreGate(t *testing.T) {
	res := rag.Result{Ready: true, Hits: []rag.Hit{
		{Title: "T", Path: "p", Excerpt: "extrait pertinent", Score: 1.0},
	}}
	if _, ok := ragContextFrom(res, ragContextBudget); ok {
		t.Fatal("ragContextFrom aurait dû refuser (porte de score 2.5)")
	}
	rc, ok := ragContextForced(res, ragContextBudget)
	if !ok || !strings.Contains(rc, "extrait pertinent") {
		t.Fatal("ragContextForced aurait dû injecter les hits faibles")
	}
	if _, ok := ragContextForced(rag.Result{Ready: true}, ragContextBudget); ok {
		t.Fatal("ragContextForced sans hit : injection inattendue")
	}
}
