package chat

import (
	"context"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/rag"
	"cetas-lite/internal/search"
	"cetas-lite/internal/store"
)

type fakeRag struct {
	ready bool
	hits  []rag.Hit
	body  string
	err   error
}

func (f fakeRag) Ready() bool { return f.ready }

func (f fakeRag) Search(_ context.Context, _ string, limit int) rag.Result {
	hits := f.hits
	if limit > 0 && len(hits) > limit {
		hits = hits[:limit]
	}
	res := rag.Result{Hits: hits, Ready: f.ready, Total: len(f.hits)}
	// Le faux simule le producteur reel : BM25Top = meilleur score BM25.
	for _, h := range hits {
		if h.Score > res.BM25Top {
			res.BM25Top = h.Score
		}
	}
	return res
}

func (f fakeRag) SearchHybrid(_ context.Context, _ string, _ map[string]float64, limit int) rag.Result {
	return f.Search(context.Background(), "", limit)
}

func (f fakeRag) SearchCorpus(_ context.Context, _ string, _ string, limit int) rag.Result {
	return f.Search(context.Background(), "", limit)
}

func (f fakeRag) SearchBoosted(_ context.Context, _ string, _ map[string]float64, limit int) rag.Result {
	return f.Search(context.Background(), "", limit)
}

func (f fakeRag) Read(string, int, int) (string, error) { return f.body, f.err }

func (f fakeRag) Stats() rag.Stats { return rag.Stats{Ready: f.ready, Chunks: len(f.hits)} }

func ragEngine(t *testing.T, r RagTools) *Engine {
	t.Helper()
	st, err := store.Open(filepath.Join(t.TempDir(), "rag.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	e := NewEngine(provider.NewRegistry(), nil, st, nil, t.TempDir())
	e.SetRAG(r)
	return e
}

func TestRagToolSchemasAndHandles(t *testing.T) {
	names := map[string]bool{}
	for _, tool := range RagToolSchemas() {
		names[tool.Function.Name] = true
	}
	if !names["rag_search"] || !names["rag_read"] {
		t.Fatalf("schemas incomplets: %v", names)
	}
	fam := ragFamily{e: ragEngine(t, fakeRag{ready: true})}
	if !fam.handles("rag_search") || !fam.handles("rag_read") || fam.handles("mem_read") {
		t.Fatal("handles incorrect")
	}
}

func TestRagExecuteSearchReadAndErrors(t *testing.T) {
	e := ragEngine(t, fakeRag{
		ready: true,
		hits:  []rag.Hit{{Title: "Alpha", Path: "a/alpha.md", Snippet: "extrait alpha"}},
		body:  "1\tligne un",
	})
	ctx := context.Background()

	out := e.ragExecute(ctx, "rag_search", `{"query":"alpha"}`)
	if !strings.Contains(out.Text, "a/alpha.md") || !strings.Contains(out.Text, "extrait alpha") {
		t.Fatalf("recherche: %q", out.Text)
	}
	if out := e.ragExecute(ctx, "rag_search", `{"query":""}`); !strings.Contains(out.Text, "[error]") {
		t.Fatalf("requete vide: %q", out.Text)
	}
	if out := e.ragExecute(ctx, "rag_read", `{"path":"a/alpha.md"}`); !strings.Contains(out.Text, "ligne un") {
		t.Fatalf("lecture: %q", out.Text)
	}
	if out := e.ragExecute(ctx, "rag_read", `{"path":""}`); !strings.Contains(out.Text, "[error]") {
		t.Fatalf("chemin vide: %q", out.Text)
	}
	if out := e.ragExecute(ctx, "rag_bidon", `{}`); !strings.Contains(out.Text, "unknown tool") {
		t.Fatalf("outil inconnu: %q", out.Text)
	}
}

func TestRagExecuteUnavailable(t *testing.T) {
	ctx := context.Background()
	e := ragEngine(t, fakeRag{ready: false})
	if out := e.ragExecute(ctx, "rag_search", `{"query":"x"}`); !strings.Contains(out.Text, "unavailable") {
		t.Fatalf("non pret: %q", out.Text)
	}
	e2 := ragEngine(t, nil)
	if out := e2.ragExecute(ctx, "rag_search", `{"query":"x"}`); !strings.Contains(out.Text, "unavailable") {
		t.Fatalf("sans RAG: %q", out.Text)
	}
}

func TestRagContextFailOpenAndBudget(t *testing.T) {
	ctx := context.Background()
	if _, ok := ragContextFrom(ragEngine(t, nil).ragHits(ctx, "q", nil, false), ragContextBudget); ok {
		t.Fatal("sans RAG: ok attendu faux")
	}
	if _, ok := ragContextFrom(ragEngine(t, fakeRag{ready: false}).ragHits(ctx, "q", nil, false), ragContextBudget); ok {
		t.Fatal("non pret: ok attendu faux")
	}
	if _, ok := ragContextFrom(ragEngine(t, fakeRag{ready: true}).ragHits(ctx, "  ", nil, false), ragContextBudget); ok {
		t.Fatal("requete vide: ok attendu faux")
	}
	if _, ok := ragContextFrom(ragEngine(t, fakeRag{ready: true}).ragHits(ctx, "q", nil, false), ragContextBudget); ok {
		t.Fatal("aucun hit: ok attendu faux")
	}
	// Porte de score : des hits faibles ne sont pas injectes.
	weak := ragEngine(t, fakeRag{ready: true, hits: []rag.Hit{
		{Title: "W", Path: "w.md", Excerpt: "extrait faible", Score: ragStrongScore - 0.1},
	}})
	if _, ok := ragContextFrom(weak.ragHits(ctx, "q", nil, false), ragContextBudget); ok {
		t.Fatal("hits faibles: ok attendu faux")
	}

	long := strings.Repeat("mot ", 2000)
	e := ragEngine(t, fakeRag{ready: true, hits: []rag.Hit{
		{Title: "A", Path: "a.md", Excerpt: long, Score: ragStrongScore},
		{Title: "B", Path: "b.md", Excerpt: long, Score: ragStrongScore + 1},
		{Title: "C", Path: "c.md", Excerpt: long, Score: ragStrongScore + 2},
	}})
	txt, ok := ragContextFrom(e.ragHits(ctx, "q", nil, false), ragContextBudget)
	if !ok {
		t.Fatal("contexte attendu")
	}
	if len(txt) > ragContextBudget+600 {
		t.Fatalf("budget depasse: %d", len(txt))
	}
	if strings.Count(txt, "mot ") > 2000 {
		t.Fatal("plus d'un extrait complet injecte")
	}
}

type captureProvider struct {
	mu   sync.Mutex
	msgs []provider.Message
}

func msgText(m provider.Message) string {
	s, _ := m.Content.(string)
	return s
}

func (c *captureProvider) ID() string { return "cap" }

func (c *captureProvider) Stream(_ context.Context, req provider.Request, emit func(provider.Event) bool) (provider.Response, error) {
	c.mu.Lock()
	c.msgs = append([]provider.Message(nil), req.Messages...)
	c.mu.Unlock()
	emit(provider.Event{Content: "ok"})
	return provider.Response{Content: "ok"}, nil
}

func TestRagContextInjectedIntoRequest(t *testing.T) {
	st, err := store.Open(filepath.Join(t.TempDir(), "inj.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	reg := provider.NewRegistry()
	cp := &captureProvider{}
	reg.Set(cp)
	e := NewEngine(reg, nil, st, nil, t.TempDir())
	e.SetRAG(fakeRag{ready: true, hits: []rag.Hit{
		{Title: "Alpha", Path: "a/alpha.md", Excerpt: "EXTRAIT-RAG-UNIQUE", Score: ragStrongScore},
	}})
	e.SetFamilies([]alias.Family{{
		ID: "code", Label: "Code",
		Modes: []alias.Mode{{ID: "standard", Pool: []alias.Member{{Provider: "cap", Model: "m"}}}},
	}})
	c := e.Conversation("u")
	if err := c.StartTurn(TurnInput{User: "u", Family: "code", Mode: "standard", Text: "parle du noyau"}); err != nil {
		t.Fatalf("StartTurn: %v", err)
	}
	waitFor(t, func() bool { return !c.IsGenerating() }, "tour non termine")

	cp.mu.Lock()
	defer cp.mu.Unlock()
	for _, m := range cp.msgs {
		if m.Role == "system" && strings.Contains(msgText(m), "EXTRAIT-RAG-UNIQUE") {
			return
		}
	}
	t.Fatalf("contexte RAG non injecte (%d messages)", len(cp.msgs))
}

func TestRagNotInjectedWhenInactive(t *testing.T) {
	st, err := store.Open(filepath.Join(t.TempDir(), "inj2.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	reg := provider.NewRegistry()
	cp := &captureProvider{}
	reg.Set(cp)
	e := NewEngine(reg, nil, st, nil, t.TempDir())
	e.SetRAG(fakeRag{ready: false, hits: []rag.Hit{{Excerpt: "EXTRAIT-RAG-UNIQUE"}}})
	e.SetFamilies([]alias.Family{{
		ID: "code", Label: "Code",
		Modes: []alias.Mode{{ID: "standard", Pool: []alias.Member{{Provider: "cap", Model: "m"}}}},
	}})
	c := e.Conversation("u")
	if err := c.StartTurn(TurnInput{User: "u", Family: "code", Mode: "standard", Text: "parle du noyau"}); err != nil {
		t.Fatalf("StartTurn: %v", err)
	}
	waitFor(t, func() bool { return !c.IsGenerating() }, "tour non termine")

	cp.mu.Lock()
	defer cp.mu.Unlock()
	for _, m := range cp.msgs {
		if strings.Contains(msgText(m), "EXTRAIT-RAG-UNIQUE") {
			t.Fatal("contexte injecte alors que le RAG est inactif")
		}
	}
}

type fakeWeb struct {
	mu    sync.Mutex
	calls int
}

func (f *fakeWeb) Search(context.Context, string, int) search.Result {
	f.mu.Lock()
	f.calls++
	f.mu.Unlock()
	return search.Result{
		Hits:     []search.Hit{{Title: "T", URL: "https://example.test", Description: "d"}},
		Provider: "fake",
	}
}

func (f *fakeWeb) Fetch(context.Context, string) (string, string, error) { return "", "", nil }

func (f *fakeWeb) count() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.calls
}

func TestRagCoveredThreshold(t *testing.T) {
	if ragCovered(rag.Result{}) {
		t.Fatal("resultat vide ne couvre pas")
	}
	if ragCovered(rag.Result{Hits: []rag.Hit{{Score: ragStrongScore - 0.1}}, BM25Top: ragStrongScore - 0.1}) {
		t.Fatal("score sous le seuil ne couvre pas")
	}
	if !ragCovered(rag.Result{Hits: []rag.Hit{{Score: ragStrongScore}}, BM25Top: ragStrongScore}) {
		t.Fatal("score au seuil couvre")
	}
}

// TestRagCoveredHybrid : porte hybride du lot 2 — un cosinus semantique
// fort couvre meme sans BM25 ; deux signaux faibles => refuse.
func TestRagCoveredHybrid(t *testing.T) {
	sem := rag.Result{
		Hits:    []rag.Hit{{Score: 0}},
		BM25Top: 0,
		SemTop:  ragSemCoverScore,
		Hybrid:  true,
	}
	if !ragCovered(sem) {
		t.Fatal("cosinus au seuil semantique couvre sans BM25")
	}
	weak := rag.Result{
		Hits:    []rag.Hit{{Score: 1.0}},
		BM25Top: 1.0,
		SemTop:  ragSemCoverScore - 0.1,
		Hybrid:  true,
	}
	if ragCovered(weak) {
		t.Fatal("deux signaux faibles ne couvrent pas")
	}
	// Sans jambe semantique, la regle se reduit a l'ancienne porte BM25.
	bm25only := rag.Result{Hits: []rag.Hit{{Score: ragStrongScore}}, BM25Top: ragStrongScore}
	if !ragCovered(bm25only) {
		t.Fatal("BM25 seul au seuil couvre (compatibilite)")
	}
}

// TestRagCoveredSemGateDisabled : CETAS_LITE_SEM_COVER <= 0 desactive la
// jambe semantique dans la porte (retour strict au comportement BM25 seul).
func TestRagCoveredSemGateDisabled(t *testing.T) {
	old := ragSemCoverScore
	defer func() { ragSemCoverScore = old }()
	ragSemCoverScore = 0
	res := rag.Result{Hits: []rag.Hit{{Score: 0}}, BM25Top: 0, SemTop: 0.99, Hybrid: true}
	if ragCovered(res) {
		t.Fatal("porte semantique desactivee mais couverte")
	}
}

func TestSemCoverFromEnv(t *testing.T) {
	t.Setenv("CETAS_LITE_SEM_COVER", "0.7")
	if v := semCoverFromEnv(); v != 0.7 {
		t.Fatalf("v = %f", v)
	}
	t.Setenv("CETAS_LITE_SEM_COVER", "nimporte")
	if v := semCoverFromEnv(); v != 0.38 {
		t.Fatalf("defaut attendu (calibre 2026-09-25, entre 0.361 et 0.394), v = %f", v)
	}
	t.Setenv("CETAS_LITE_SEM_COVER", "")
	if v := semCoverFromEnv(); v != 0.38 {
		t.Fatalf("defaut sans variable, v = %f", v)
	}
}

// TestRagDirectiveNaturalNoCitations : l'iteration 2 impose le naturel
// total — plus aucune consigne de citation visible dans la directive
// injectee, et interdiction explicite d'inventer.
func TestRagDirectiveNaturalNoCitations(t *testing.T) {
	res := rag.Result{Hits: []rag.Hit{
		{Title: "T", Path: "t.md", Excerpt: "extrait", Score: ragStrongScore},
	}, BM25Top: ragStrongScore}
	txt, ok := ragContextFrom(res, ragContextBudget)
	if !ok {
		t.Fatal("contexte attendu pour des hits forts")
	}
	for _, banned := range []string{"cite them [1]", "Cite discrètement", "cite [1]"} {
		if strings.Contains(txt, banned) {
			t.Fatalf("consigne de citation encore presente: %q", banned)
		}
	}
	if !strings.Contains(txt, "treat it as things you know") {
		t.Fatal("directive 'treat it as things you know' absente")
	}
	if !strings.Contains(txt, "are banned") {
		t.Fatal("interdiction explicite (are banned) absente de la directive")
	}
	if !strings.Contains(txt, "without verbalizing this rule") {
		t.Fatal("interdiction de verbaliser la regle absente de la directive")
	}
	// Aucune chaîne FR dans la directive (règle globale : tout en anglais).
	for _, fr := range []string{"Extraits de la base", "Consignes de réponse", "aucune citation visible"} {
		if strings.Contains(txt, fr) {
			t.Fatalf("texte FR encore present dans la directive: %q", fr)
		}
	}
	if strings.Contains(localFirstDirective(), "[1]") {
		t.Fatal("localFirstDirective mentionne encore des citations")
	}
}

// TestRagDirectiveGroundingPlusReasoning : l'iteration 9 (directive
// assouplie) fait des extraits la source PRIORITAIRE (ancrage
// anti-hallucination conserve) mais plus la source EXCLUSIVE : le modele
// peut raisonner a travers eux et completer avec ses propres connaissances.
// L'ancienne formulation stricte a disparu.
func TestRagDirectiveGroundingPlusReasoning(t *testing.T) {
	res := rag.Result{Hits: []rag.Hit{
		{Title: "T", Path: "t.md", Excerpt: "extrait", Score: ragStrongScore},
	}, BM25Top: ragStrongScore}
	txt, ok := ragContextFrom(res, ragContextBudget)
	if !ok {
		t.Fatal("contexte attendu pour des hits forts")
	}
	// Ancrage conserve : les extraits sont la source prioritaire.
	if !strings.Contains(txt, "primary source") {
		t.Fatal("ancrage primary source absent de la directive")
	}
	if !strings.Contains(txt, "never present internal knowledge or guesses as coming from the local document base") {
		t.Fatal("interdiction de presenter des suppositions comme faits de la base absente")
	}
	// Autorisation : raisonnement croise + connaissances propres.
	for _, want := range []string{"Reason across the extracts", "your own knowledge", "reason further"} {
		if !strings.Contains(txt, want) {
			t.Fatalf("autorisation de raisonner absente: %q", want)
		}
	}
	// L'exclusivite stricte ne doit plus etre la.
	for _, banned := range []string{
		"only the facts present",
		"Use only the facts",
		"Prioritize the facts above",
	} {
		if strings.Contains(txt, banned) {
			t.Fatalf("exclusivite stricte encore presente: %q", banned)
		}
	}
	// Protocole it.2 preserve : couvert / non couvert (refus honnete).
	if _, ok := ragContextFrom(rag.Result{}, ragContextBudget); ok {
		t.Fatal("sans hit, aucun contexte attendu (fail-open)")
	}
	if !strings.Contains(ragNotCoveredNote(), "does not cover this request") {
		t.Fatal("refus honnete de non-couverture perdu")
	}
	// La regle permanente du prompt systeme ne bouge pas.
	if !strings.Contains(chatSystemPrompt(),
		"Never present internal knowledge or guesses as coming from the local document base.") {
		t.Fatal("regle anti-hallucination permanente modifiee")
	}
}

func TestChatSystemPromptAntiHallucination(t *testing.T) {
	p := chatSystemPrompt()
	if !strings.Contains(p, "Never present internal knowledge or guesses as coming from the local document base.") {
		t.Fatal("regle anti-hallucination RAG absente du prompt permanent")
	}
}

// TestRagNotCoveredNoteInjected : base active mais requete non couverte →
// la note « non couvert » est injectee comme message systeme ; couverte →
// absente.
func TestRagNotCoveredNoteInjected(t *testing.T) {
	check := func(hits []rag.Hit, wantNote bool) {
		t.Helper()
		st, err := store.Open(filepath.Join(t.TempDir(), "nc.db"))
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() { _ = st.Close() })
		reg := provider.NewRegistry()
		cp := &captureProvider{}
		reg.Set(cp)
		e := NewEngine(reg, nil, st, nil, t.TempDir())
		e.SetRAG(fakeRag{ready: true, hits: hits})
		e.SetFamilies([]alias.Family{{
			ID: "code", Label: "Code",
			Modes: []alias.Mode{{ID: "standard", Pool: []alias.Member{{Provider: "cap", Model: "m"}}}},
		}})
		c := e.Conversation("u")
		if err := c.StartTurn(TurnInput{User: "u", Family: "code", Mode: "standard", Text: "question hors base"}); err != nil {
			t.Fatalf("StartTurn: %v", err)
		}
		waitFor(t, func() bool { return !c.IsGenerating() }, "tour non termine")

		cp.mu.Lock()
		defer cp.mu.Unlock()
		found := false
		for _, m := range cp.msgs {
			if m.Role == "system" && strings.Contains(msgText(m), "does not cover this request") {
				found = true
			}
		}
		if found != wantNote {
			t.Fatalf("note 'non couvert' : trouvee=%v, attendue=%v", found, wantNote)
		}
	}
	check([]rag.Hit{{Title: "W", Path: "w.md", Excerpt: "faible", Score: 1.0}}, true)
	check([]rag.Hit{{Title: "S", Path: "s.md", Excerpt: "fort", Score: ragStrongScore}}, false)
}

// TestRagContextRefusedWhenWeak : la porte de score de l'iteration 1 —
// des hits sous le seuil ne produisent aucun contexte injecte, meme
// si des passages existent.
func TestRagContextRefusedWhenWeak(t *testing.T) {
	weak := rag.Result{Hits: []rag.Hit{
		{Title: "Hors sujet", Path: "x.md", Excerpt: "extrait", Score: 1.2},
		{Title: "Bruit", Path: "y.md", Excerpt: "extrait", Score: 0.8},
	}}
	if txt, ok := ragContextFrom(weak, ragContextBudget); ok || txt != "" {
		t.Fatalf("hits faibles injectes: ok=%v txt=%q", ok, txt)
	}
	strong := rag.Result{Hits: []rag.Hit{
		{Title: "Couvert", Path: "z.md", Excerpt: "extrait pertinent", Score: ragStrongScore},
	}, BM25Top: ragStrongScore}
	txt, ok := ragContextFrom(strong, ragContextBudget)
	if !ok || !strings.Contains(txt, "extrait pertinent") {
		t.Fatalf("hits forts non injectes: ok=%v", ok)
	}
}

func runWebTurn(t *testing.T, r RagTools) *fakeWeb {
	t.Helper()
	st, err := store.Open(filepath.Join(t.TempDir(), "w.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	reg := provider.NewRegistry()
	reg.Set(&captureProvider{})
	e := NewEngine(reg, nil, st, nil, t.TempDir())
	fw := &fakeWeb{}
	e.SetSearcher(fw)
	e.SetRAG(r)
	e.SetFamilies([]alias.Family{{
		ID: "code", Label: "Code",
		Modes: []alias.Mode{{ID: "standard", Pool: []alias.Member{{Provider: "cap", Model: "m"}}}},
	}})
	c := e.Conversation("u")
	if err := c.StartTurn(TurnInput{User: "u", Family: "code", Mode: "standard", Text: "sujet", Web: true}); err != nil {
		t.Fatalf("StartTurn: %v", err)
	}
	waitFor(t, func() bool { return !c.IsGenerating() }, "tour non termine")
	return fw
}

func TestRagSkipsWebPreSearchWhenCovered(t *testing.T) {
	fw := runWebTurn(t, fakeRag{ready: true, hits: []rag.Hit{
		{Title: "A", Path: "a.md", Excerpt: "extrait", Score: 9.0},
	}})
	if n := fw.count(); n != 0 {
		t.Fatalf("pre-recherche web non sautee: %d appel(s)", n)
	}
}

func TestRagKeepsWebPreSearchWhenWeak(t *testing.T) {
	fw := runWebTurn(t, fakeRag{ready: true, hits: []rag.Hit{
		{Title: "A", Path: "a.md", Excerpt: "extrait", Score: 0.4},
	}})
	if n := fw.count(); n == 0 {
		t.Fatal("pre-recherche web attendue (base locale non concluante)")
	}
}

// TestInsertBeforeLastUser : le contexte dynamique RAG doit etre insere
// juste avant le dernier message utilisateur (le tour courant), donc apres
// l'historique. Le prefixe [prompts stables + historique] reste ainsi
// identique d'un tour a l'autre, ce qui favorise le prompt caching.
func TestInsertBeforeLastUser(t *testing.T) {
	dyn := provider.Message{Role: "system", Content: "RAG"}
	msgs := []provider.Message{
		{Role: "system", Content: "sys"},
		{Role: "user", Content: "u1"},
		{Role: "assistant", Content: "a1"},
		{Role: "user", Content: "u2"},
	}
	got := insertBeforeLastUser(msgs, dyn)
	want := []string{"sys", "u1", "a1", "RAG", "u2"}
	if len(got) != len(want) {
		t.Fatalf("longueur = %d, attendu %d", len(got), len(want))
	}
	for i, w := range want {
		if msgText(got[i]) != w {
			t.Fatalf("position %d = %q, attendu %q", i, msgText(got[i]), w)
		}
	}

	// Sans dernier message utilisateur : ajout a la fin.
	msgs2 := []provider.Message{
		{Role: "system", Content: "sys"},
		{Role: "assistant", Content: "a"},
	}
	got2 := insertBeforeLastUser(msgs2, dyn)
	if msgText(got2[len(got2)-1]) != "RAG" {
		t.Fatalf("dernier message = %q, attendu RAG", msgText(got2[len(got2)-1]))
	}

	// Sans message dynamique : tranche inchangee.
	got3 := insertBeforeLastUser(msgs)
	if len(got3) != len(msgs) {
		t.Fatalf("longueur = %d, attendu %d", len(got3), len(msgs))
	}
}

// captureReadRag enregistre les bornes reellement transmises a Read.
type captureReadRag struct {
	fakeRag
	offset, limit int
}

func (c *captureReadRag) Read(_ string, offset, limit int) (string, error) {
	c.offset, c.limit = offset, limit
	return "ok", nil
}

// TestRagReadServerBounds : rag_read applique une limite par defaut et un
// plafond serveur strict avant d'appeler le lecteur.
func TestRagReadServerBounds(t *testing.T) {
	cases := []struct {
		name       string
		args       string
		wantOffset int
		wantLimit  int
	}{
		{"defaut", `{"path":"doc.md"}`, 0, ragReadDefaultLines},
		{"offset negatif normalise", `{"path":"doc.md","offset":-5}`, 0, ragReadDefaultLines},
		{"plafond serveur", `{"path":"doc.md","limit":10000}`, 0, ragReadMaxLines},
		{"valeurs explicites conservees", `{"path":"doc.md","offset":10,"limit":25}`, 10, 25},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := &captureReadRag{fakeRag: fakeRag{ready: true}}
			e := ragEngine(t, c)
			res := e.ragExecute(context.Background(), "rag_read", tc.args)
			if res.Text != "ok" {
				t.Fatalf("resultat = %q, attendu ok", res.Text)
			}
			if c.offset != tc.wantOffset || c.limit != tc.wantLimit {
				t.Fatalf("offset/limit = %d/%d, attendu %d/%d", c.offset, c.limit, tc.wantOffset, tc.wantLimit)
			}
		})
	}
}

// TestRagContextSynthesisDirective : l'iteration 6 demande une synthese
// croisee en prose continue, pas une section par extrait. Les garde-fous
// de l'iteration 2 (naturel, anti-hallucination) restent en place.
func TestRagContextSynthesisDirective(t *testing.T) {
	res := rag.Result{Hits: []rag.Hit{
		{Title: "T", Path: "t.md", Excerpt: "extrait", Score: ragStrongScore},
	}, BM25Top: ragStrongScore}
	txt, ok := ragContextFrom(res, ragContextBudget)
	if !ok {
		t.Fatal("contexte attendu pour des hits forts")
	}
	for _, want := range []string{
		"Weave the information together",
		"continuous natural prose",
		"no section per excerpt",
	} {
		if !strings.Contains(txt, want) {
			t.Fatalf("consigne de synthese absente: %q", want)
		}
	}
	for _, keep := range []string{
		"without verbalizing this rule",
		"say you don't know",
	} {
		if !strings.Contains(txt, keep) {
			t.Fatalf("garde-fou iteration 2 perdu: %q", keep)
		}
	}
}

// TestRagNotCoveredNoteNatural : la note « non couvert » garde le sens de
// l'iteration 2 (dire la non-couverture, interdire l'invention presentee
// comme issue du corpus) mais modelise un ton naturel, sans rapport formel.
func TestRagNotCoveredNoteNatural(t *testing.T) {
	n := ragNotCoveredNote()
	for _, want := range []string{
		"does not cover this request",
		"do not present anything as coming from it",
		"naturally",
		"without enumerating sources",
	} {
		if !strings.Contains(n, want) {
			t.Fatalf("note 'non couvert' incomplete, manque %q", want)
		}
	}
	// Aucune chaîne FR dans la note (règle globale : tout en anglais).
	for _, fr := range []string{"ne couvre pas", "présente rien", "énumérant"} {
		if strings.Contains(n, fr) {
			t.Fatalf("texte FR encore present dans la note: %q", fr)
		}
	}
}

// captureQueryRag : faux RagTools qui capture la requete envoyee a Search
// et le boost (iteration 6b).
type captureQueryRag struct {
	ready bool
	query string
	boost map[string]float64
}

func (f *captureQueryRag) Ready() bool { return f.ready }

func (f *captureQueryRag) Search(_ context.Context, q string, _ int) rag.Result {
	f.query = q
	f.boost = nil
	return rag.Result{Ready: f.ready}
}

func (f *captureQueryRag) SearchCorpus(_ context.Context, q, _ string, _ int) rag.Result {
	return f.Search(context.Background(), q, 0)
}

func (f *captureQueryRag) SearchBoosted(_ context.Context, q string, boost map[string]float64, _ int) rag.Result {
	return f.SearchHybrid(context.Background(), q, boost, 0)
}

func (f *captureQueryRag) SearchHybrid(_ context.Context, q string, boost map[string]float64, _ int) rag.Result {
	f.query = q
	f.boost = boost
	return rag.Result{Ready: f.ready}
}

func (f *captureQueryRag) Read(string, int, int) (string, error) { return "", nil }

func (f *captureQueryRag) Stats() rag.Stats { return rag.Stats{Ready: f.ready} }

// TestRagHitsEnrichesEllipticalQuery : le tour principal enrichit les
// requetes elliptiques avec le sujet de l'historique avant BM25
// (iteration 6) ; les requetes deja porteuses d'entites passent telles quelles.
func TestRagHitsEnrichesEllipticalQuery(t *testing.T) {
	f := &captureQueryRag{ready: true}
	e := ragEngine(t, f)
	hist := []provider.Message{
		{Role: "user", Content: "Qu'est-ce que Kimi K3 et que vaut-il ?"},
		{Role: "assistant", Content: "Kimi K3 est le modèle de Moonshot AI."},
	}
	e.ragHits(context.Background(), "on peut la faire tourner sur combien de cpu ?", hist, false)
	if !strings.Contains(f.query, "K3") {
		t.Fatalf("requete non enrichie: %q", f.query)
	}
	// Iteration 6b : l'enrichissement declenche SearchHybrid avec les
	// tokens d'entite et le facteur calibre. L'historique le plus recent
	// (assistant) domine : groupes « K3 » + « Moonshot AI » — « Kimi » en
	// debut de phrase n'est pas un marqueur d'entite (regle iter6 ; le
	// message utilisateur seul produirait « Kimi K3 », cf. 2e cas).
	if f.boost == nil {
		t.Fatal("boost absent : la requete enrichie doit passer par SearchHybrid")
	}
	for _, k := range []string{"k3", "moonshot", "ai"} {
		if f.boost[k] != entityBoostFactor {
			t.Fatalf("boost[%q] = %v, attendu %v (boost=%v)", k, f.boost[k], entityBoostFactor, f.boost)
		}
	}
	// Historique reduit au message utilisateur : le groupe complet est repris.
	e.ragHits(context.Background(), "on peut la faire tourner sur combien de cpu ?",
		[]provider.Message{{Role: "user", Content: "Qu'est-ce que Kimi K3 et que vaut-il ?"}}, false)
	if !strings.Contains(f.query, "Kimi K3") {
		t.Fatalf("sujet utilisateur non repris: %q", f.query)
	}
	e.ragHits(context.Background(), "parle-moi de Kimi K3", hist, false)
	if strings.Contains(f.query, "Moonshot") {
		t.Fatalf("requete porteuse d'entite enrichie a tort: %q", f.query)
	}
	if f.boost != nil {
		t.Fatalf("requête directe : boost inattendu %v (Search attendu)", f.boost)
	}
	// Historique vide : fail-open, requete inchangee, sans boost.
	e.ragHits(context.Background(), "on peut la faire tourner sur combien de cpu ?", nil, false)
	if strings.Contains(f.query, "Kimi") || f.boost != nil {
		t.Fatalf("enrichissement sans historique: %q boost=%v", f.query, f.boost)
	}
}

// recRag enregistre quelle methode de recherche est appelee (C2 : les
// moteurs locaux doivent passer par BM25 seul, jamais par la jambe
// semantique).
type recRag struct {
	RagTools
	searchCalls  int
	hybridCalls  int
	boostedCalls int
}

func (r *recRag) Search(ctx context.Context, q string, limit int) rag.Result {
	r.searchCalls++
	return r.RagTools.Search(ctx, q, limit)
}

func (r *recRag) SearchHybrid(ctx context.Context, q string, boost map[string]float64, limit int) rag.Result {
	r.hybridCalls++
	return r.RagTools.SearchHybrid(ctx, q, boost, limit)
}

func (r *recRag) SearchBoosted(ctx context.Context, q string, boost map[string]float64, limit int) rag.Result {
	r.boostedCalls++
	return r.RagTools.SearchBoosted(ctx, q, boost, limit)
}

func TestRagHitsLocalSkipsSemanticLeg(t *testing.T) {
	rec := &recRag{RagTools: fakeRag{ready: true, hits: []rag.Hit{{Title: "h"}}}}
	e := ragEngine(t, rec)
	e.ragHits(context.Background(), "une question anodine", nil, true)
	if rec.hybridCalls != 0 || rec.searchCalls != 1 {
		t.Fatalf("local => BM25 seul attendu (search=%d hybrid=%d)", rec.searchCalls, rec.hybridCalls)
	}
	e.ragHits(context.Background(), "une question anodine", nil, false)
	if rec.hybridCalls != 1 {
		t.Fatalf("cloud => hybride attendu (hybrid=%d)", rec.hybridCalls)
	}
}

// C1 : les extraits locaux tiennent dans un budget reduit (2 hits /
// 1200 caracteres) — le premier hit passe, le second depasse.
func TestBuildRagContextBudgetLocal(t *testing.T) {
	hits := []rag.Hit{
		{Title: "t1", Path: "p1", Excerpt: strings.Repeat("x", 900)},
		{Title: "t2", Path: "p2", Excerpt: strings.Repeat("y", 900)},
	}
	local := buildRagContext(rag.Result{Hits: hits}, ragContextBudgetLocal)
	if !strings.Contains(local, "t1") {
		t.Fatal("premier hit attendu dans le budget local")
	}
	if strings.Contains(local, "t2") {
		t.Fatal("second hit hors budget local : ne doit pas etre injecte")
	}
	full := buildRagContext(rag.Result{Hits: hits}, ragContextBudget)
	if !strings.Contains(full, "t2") {
		t.Fatal("budget cloud : les deux hits doivent tenir")
	}
}
