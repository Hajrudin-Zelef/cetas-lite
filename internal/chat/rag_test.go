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
	return rag.Result{Hits: hits, Ready: f.ready, Total: len(f.hits)}
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
	if out := e.ragExecute(ctx, "rag_search", `{"query":""}`); !strings.Contains(out.Text, "[erreur]") {
		t.Fatalf("requete vide: %q", out.Text)
	}
	if out := e.ragExecute(ctx, "rag_read", `{"path":"a/alpha.md"}`); !strings.Contains(out.Text, "ligne un") {
		t.Fatalf("lecture: %q", out.Text)
	}
	if out := e.ragExecute(ctx, "rag_read", `{"path":""}`); !strings.Contains(out.Text, "[erreur]") {
		t.Fatalf("chemin vide: %q", out.Text)
	}
	if out := e.ragExecute(ctx, "rag_bidon", `{}`); !strings.Contains(out.Text, "outil inconnu") {
		t.Fatalf("outil inconnu: %q", out.Text)
	}
}

func TestRagExecuteUnavailable(t *testing.T) {
	ctx := context.Background()
	e := ragEngine(t, fakeRag{ready: false})
	if out := e.ragExecute(ctx, "rag_search", `{"query":"x"}`); !strings.Contains(out.Text, "indisponible") {
		t.Fatalf("non pret: %q", out.Text)
	}
	e2 := ragEngine(t, nil)
	if out := e2.ragExecute(ctx, "rag_search", `{"query":"x"}`); !strings.Contains(out.Text, "indisponible") {
		t.Fatalf("sans RAG: %q", out.Text)
	}
}

func TestRagContextFailOpenAndBudget(t *testing.T) {
	ctx := context.Background()
	if _, ok := ragContextFrom(ragEngine(t, nil).ragHits(ctx, "q")); ok {
		t.Fatal("sans RAG: ok attendu faux")
	}
	if _, ok := ragContextFrom(ragEngine(t, fakeRag{ready: false}).ragHits(ctx, "q")); ok {
		t.Fatal("non pret: ok attendu faux")
	}
	if _, ok := ragContextFrom(ragEngine(t, fakeRag{ready: true}).ragHits(ctx, "  ")); ok {
		t.Fatal("requete vide: ok attendu faux")
	}
	if _, ok := ragContextFrom(ragEngine(t, fakeRag{ready: true}).ragHits(ctx, "q")); ok {
		t.Fatal("aucun hit: ok attendu faux")
	}
	// Porte de score : des hits faibles ne sont pas injectes.
	weak := ragEngine(t, fakeRag{ready: true, hits: []rag.Hit{
		{Title: "W", Path: "w.md", Excerpt: "extrait faible", Score: ragStrongScore - 0.1},
	}})
	if _, ok := ragContextFrom(weak.ragHits(ctx, "q")); ok {
		t.Fatal("hits faibles: ok attendu faux")
	}

	long := strings.Repeat("mot ", 2000)
	e := ragEngine(t, fakeRag{ready: true, hits: []rag.Hit{
		{Title: "A", Path: "a.md", Excerpt: long, Score: ragStrongScore},
		{Title: "B", Path: "b.md", Excerpt: long, Score: ragStrongScore + 1},
		{Title: "C", Path: "c.md", Excerpt: long, Score: ragStrongScore + 2},
	}})
	txt, ok := ragContextFrom(e.ragHits(ctx, "q"))
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
	if ragCovered(rag.Result{Hits: []rag.Hit{{Score: ragStrongScore - 0.1}}}) {
		t.Fatal("score sous le seuil ne couvre pas")
	}
	if !ragCovered(rag.Result{Hits: []rag.Hit{{Score: ragStrongScore}}}) {
		t.Fatal("score au seuil couvre")
	}
}

// TestRagDirectiveNaturalNoCitations : l'iteration 2 impose le naturel
// total — plus aucune consigne de citation visible dans la directive
// injectee, et interdiction explicite d'inventer.
func TestRagDirectiveNaturalNoCitations(t *testing.T) {
	res := rag.Result{Hits: []rag.Hit{
		{Title: "T", Path: "t.md", Excerpt: "extrait", Score: ragStrongScore},
	}}
	txt, ok := ragContextFrom(res)
	if !ok {
		t.Fatal("contexte attendu pour des hits forts")
	}
	for _, banned := range []string{"cite them [1]", "Cite discrètement", "cite [1]"} {
		if strings.Contains(txt, banned) {
			t.Fatalf("consigne de citation encore presente: %q", banned)
		}
	}
	if !strings.Contains(txt, "aucune citation visible") {
		t.Fatal("directive 'naturel total' absente")
	}
	if !strings.Contains(txt, "dis-le franchement au lieu de l'inventer") {
		t.Fatal("interdiction d'inventer absente de la directive")
	}
	if strings.Contains(localFirstDirective(), "[1]") {
		t.Fatal("localFirstDirective mentionne encore des citations")
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
			if m.Role == "system" && strings.Contains(msgText(m), "ne couvre pas cette demande") {
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
	if txt, ok := ragContextFrom(weak); ok || txt != "" {
		t.Fatalf("hits faibles injectes: ok=%v txt=%q", ok, txt)
	}
	strong := rag.Result{Hits: []rag.Hit{
		{Title: "Couvert", Path: "z.md", Excerpt: "extrait pertinent", Score: ragStrongScore},
	}}
	txt, ok := ragContextFrom(strong)
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
