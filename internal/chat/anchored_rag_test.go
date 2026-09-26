package chat

import (
	"context"
	"path/filepath"
	"strings"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/rag"
	"cetas-lite/internal/store"
)

// captures the RAG context actually injected into a real turn, to prove the
// anchored fiche reaches the model end-to-end (not just via ragHits).
type anchorRag struct {
	fakeRag
}

func (a anchorRag) SearchHybrid(ctx context.Context, q string, _ map[string]float64, limit int) rag.Result {
	return a.Search(ctx, q, limit)
}

func TestAnchoredFicheReachesModel(t *testing.T) {
	st, err := store.Open(filepath.Join(t.TempDir(), "anchor.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	reg := provider.NewRegistry()
	cp := &multiCaptureProvider{}
	reg.Set(cp)
	e := NewEngine(reg, nil, st, nil, t.TempDir())

	fiche := "The xAI Colossus cluster counts 100,000 NVIDIA H100 GPUs and was built in 122 days."
	e.SetRAG(anchorRag{fakeRag: fakeRag{ready: true, hits: []rag.Hit{
		{Title: "inside-the-100k-gpu-xai-colossus", Path: "collect-250926-servers-hardware/colossus.md", Excerpt: fiche, Score: ragStrongScore},
	}}})
	e.SetFamilies([]alias.Family{{
		ID: "code", Label: "Code",
		Modes: []alias.Mode{{ID: "standard", Pool: []alias.Member{{Provider: "cap", Model: "m"}}}},
	}})

	c := e.Conversation("u")
	if err := c.StartTurn(TurnInput{
		User: "u", Family: "code", Mode: "standard",
		Text: "combien de GPU H100 compte Colossus ?",
	}); err != nil {
		t.Fatalf("StartTurn: %v", err)
	}
	waitFor(t, func() bool { return !c.IsGenerating() }, "tour non termine")

	cp.mu.Lock()
	defer cp.mu.Unlock()
	if len(cp.reqs) == 0 {
		t.Fatal("aucune requete capturee")
	}
	var joined strings.Builder
	for _, m := range cp.reqs[0].Messages {
		joined.WriteString(msgText(m))
		joined.WriteString("\n")
	}
	all := joined.String()
	if !strings.Contains(all, "100,000 NVIDIA H100") {
		t.Fatalf("extrait de la fiche absent du prompt modele:\n%.600q", all)
	}
	if !strings.Contains(all, "inside-the-100k-gpu-xai-colossus") {
		t.Fatalf("reference de la fiche absente du prompt modele")
	}
	t.Log("fiche ancree injectee dans le prompt (extrait + reference presents)")
}
