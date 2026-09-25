package chat

import (
	"context"
	"sync"
	"testing"
	"time"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/provider"
)

// --- Clés ---

func TestPregenInputKeyDeterministic(t *testing.T) {
	in := TurnInput{User: "u", Family: "f", Mode: "m", Text: "q?", MaxTokens: 100, FocusCorpus: "c"}
	h := []provider.Message{{Role: "user", Content: "salut"}}
	k1 := pregenInputKey(in, h)
	k2 := pregenInputKey(in, h)
	if k1 == "" || k1 != k2 {
		t.Fatalf("cle non deterministe : %q vs %q", k1, k2)
	}
	in2 := in
	in2.Text = "autre ?"
	if k3 := pregenInputKey(in2, h); k3 == k1 {
		t.Fatalf("cle identique pour une question differente")
	}
	// nil vs vide : même clé (pas de faux miss).
	if k4 := pregenInputKey(in, nil); k4 != pregenInputKey(in, []provider.Message{}) {
		t.Fatalf("nil != vide : %q vs %q", k4, pregenInputKey(in, []provider.Message{}))
	}
}

func TestPregenReqKeyDeterministic(t *testing.T) {
	req := provider.Request{Model: "m", Messages: []provider.Message{{Role: "user", Content: "q"}}, Temperature: 0.7, MaxTokens: 100}
	k1 := pregenReqKey("p", req)
	k2 := pregenReqKey("p", req)
	if k1 == "" || k1 != k2 {
		t.Fatalf("cle requete non deterministe")
	}
	req2 := req
	req2.Temperature = 0
	if pregenReqKey("p", req2) == k1 {
		t.Fatalf("cle identique pour une temperature differente")
	}
}

// --- Dépôt ---

func TestPregenStoreFillWait(t *testing.T) {
	s := newPregenStore()
	e := s.register("k")
	go func() {
		time.Sleep(50 * time.Millisecond)
		s.fill("k", "rk", provider.Response{Content: "bonjour"})
	}()
	select {
	case <-e.done:
	case <-time.After(5 * time.Second):
		t.Fatalf("attente jamais resolue")
	}
	if e.err != nil || e.content != "bonjour" || e.reqKey != "rk" {
		t.Fatalf("entree mal remplie : %+v", e)
	}
	// Double fill : premier gagne, pas de panic (double close).
	s.fill("k", "rk2", provider.Response{Content: "autre"})
	if e.content != "bonjour" {
		t.Fatalf("second fill a ecrase : %q", e.content)
	}
}

func TestPregenStoreAbort(t *testing.T) {
	s := newPregenStore()
	e := s.register("k")
	s.abort("k", context.Canceled)
	select {
	case <-e.done:
	default:
		t.Fatalf("done non ferme apres abort")
	}
	if e.err == nil {
		t.Fatalf("err attendu apres abort")
	}
}

func TestPregenAbortEntryIdentity(t *testing.T) {
	s := newPregenStore()
	pe1 := s.register("k")
	// Un fantôme d'un autre job (pointeur différent) ne doit pas
	// empoisonner l'entrée courante.
	other := &pregenEntry{done: make(chan struct{})}
	s.abortEntry("k", other, context.Canceled)
	select {
	case <-pe1.done:
		t.Fatalf("entree avortee par un fantome etranger")
	default:
	}
	// Le bon fantôme, si.
	s.abortEntry("k", pe1, context.Canceled)
	select {
	case <-pe1.done:
	default:
		t.Fatalf("done non ferme apres abort du bon fantome")
	}
}

func TestPregenStoreBound(t *testing.T) {
	s := newPregenStore()
	for i := 0; i < pregenMaxEntries+10; i++ {
		k := string(rune('a'+i%26)) + string(rune('0'+i/26))
		e := s.register(k)
		s.fill(k, "rk", provider.Response{Content: "x"})
		_ = e
	}
	s.mu.Lock()
	n := len(s.entries)
	s.mu.Unlock()
	if n > pregenMaxEntries {
		t.Fatalf("depot non borne : %d > %d", n, pregenMaxEntries)
	}
	// Un fantôme en cours n'est jamais évincé.
	s2 := newPregenStore()
	for i := 0; i < pregenMaxEntries+5; i++ {
		k := string(rune('a'+i%26)) + string(rune('0'+i/26))
		s2.register(k)
		if i < pregenMaxEntries {
			s2.fill(k, "rk", provider.Response{Content: "x"})
		}
	}
	inflight := 0
	s2.mu.Lock()
	for _, e := range s2.entries {
		if !e.completed {
			inflight++
		}
	}
	s2.mu.Unlock()
	if inflight != 5 {
		t.Fatalf("fantomes en cours evinces : %d restants", inflight)
	}
}

func TestPregenCancel(t *testing.T) {
	e := &Engine{}
	e.PregenCancel("u") // sans job : no-op, pas de panic
	e.pregen.mu.Lock()
	e.pregen.jobs = map[string]context.CancelFunc{"u": func() {}}
	e.pregen.mu.Unlock()
	e.PregenCancel("u")
	e.pregen.mu.Lock()
	_, ok := e.pregen.jobs["u"]
	e.pregen.mu.Unlock()
	if ok {
		t.Fatalf("job non supprime apres cancel")
	}
}

// --- Bout en bout : fantôme puis clic servi sans nouvel appel ---

type countingProvider struct {
	id    string
	mu    sync.Mutex
	calls int
	text  string
}

func (p *countingProvider) ID() string { return p.id }

func (p *countingProvider) Stream(ctx context.Context, req provider.Request, emit func(provider.Event) bool) (provider.Response, error) {
	p.mu.Lock()
	p.calls++
	p.mu.Unlock()
	emit(provider.Event{Content: p.text})
	return provider.Response{Content: p.text}, nil
}

func (p *countingProvider) n() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.calls
}

func pregenTestEngine(t *testing.T, cp *countingProvider) *Engine {
	t.Helper()
	fams := []alias.Family{{
		ID: "test", Label: "Test",
		Modes: []alias.Mode{{ID: "standard", Label: "Std", Pool: []alias.Member{{Provider: "cnt", Model: "m1"}}}},
	}}
	e := newEngine(t, &fakeProvider{id: "dummy"}, fams)
	reg := provider.NewRegistry()
	reg.Set(cp)
	e.reg = reg
	return e
}

func TestPregenShadowThenClickServed(t *testing.T) {
	cp := &countingProvider{id: "cnt", text: "réponse pré-générée"}
	e := pregenTestEngine(t, cp)
	base := TurnInput{User: "u", Family: "test", Mode: "standard", MaxTokens: 500}
	items := []PregenSuggestion{{Question: "qu'est-ce que vLLM ?", Corpus: "00-nvidia"}}
	e.PregenSuggestions("u", items, base, nil)

	// Attendre la fin du fantôme (1 seul appel provider).
	waitFor(t, func() bool { return cp.n() == 1 }, "fantome non termine")
	key := pregenInputKey(TurnInput{User: "u", Family: "test", Mode: "standard", Text: items[0].Question, MaxTokens: 500, FocusCorpus: "00-nvidia"}, nil)
	pe := e.pregen.getStore().get(key)
	if pe == nil {
		t.Fatalf("entree absente du depot")
	}
	select {
	case <-pe.done:
	case <-time.After(5 * time.Second):
		t.Fatalf("fantome non termine")
	}
	if pe.err != nil {
		t.Fatalf("fantome en erreur : %v", pe.err)
	}

	// Le clic : même question, PregenLookup => servi sans nouvel appel.
	in := TurnInput{User: "u", Family: "test", Mode: "standard", Text: items[0].Question, MaxTokens: 500, FocusCorpus: "00-nvidia", PregenLookup: true}
	c := runTurn(t, e, "u", in)
	if n := cp.n(); n != 1 {
		t.Fatalf("le clic a rappele le provider (%d appels), pregen non servie", n)
	}
	if txt := logText(c); txt != "réponse pré-générée" {
		t.Fatalf("contenu inattendu : %q", txt)
	}
	// Stats marquées cached.
	c.mu.Lock()
	cached := false
	for _, ev := range c.Log {
		if st, ok := ev.Delta["stats"].(map[string]any); ok {
			if v, ok := st["cached"].(bool); ok && v {
				cached = true
			}
		}
	}
	c.mu.Unlock()
	if !cached {
		t.Fatalf("stats non marquees cached")
	}
}

func TestPregenClickWithoutPregenNormalPath(t *testing.T) {
	cp := &countingProvider{id: "cnt", text: "réponse normale"}
	e := pregenTestEngine(t, cp)
	// Pas de pré-génération : le clic suit le chemin normal (1 appel).
	in := TurnInput{User: "u", Family: "test", Mode: "standard", Text: "question ?", MaxTokens: 500, FocusCorpus: "00-nvidia", PregenLookup: true}
	c := runTurn(t, e, "u", in)
	if n := cp.n(); n != 1 {
		t.Fatalf("appels provider inattendus : %d", n)
	}
	if txt := logText(c); txt != "réponse normale" {
		t.Fatalf("contenu inattendu : %q", txt)
	}
}

func TestPregenShadowSkipsAgent(t *testing.T) {
	cp := &countingProvider{id: "cnt", text: "x"}
	fams := []alias.Family{{
		ID: "test", Label: "Test",
		Modes: []alias.Mode{{ID: "agent", Label: "Ag", Agent: true, Pool: []alias.Member{{Provider: "cnt", Model: "m1"}}}},
	}}
	e := newEngine(t, &fakeProvider{id: "dummy"}, fams)
	reg := provider.NewRegistry()
	reg.Set(cp)
	e.reg = reg
	base := TurnInput{User: "u", Family: "test", Mode: "agent", MaxTokens: 500}
	e.PregenSuggestions("u", []PregenSuggestion{{Question: "q ?", Corpus: "c"}}, base, nil)
	// Le fantôme agent ne doit jamais partir : 0 appel, même après un délai.
	time.Sleep(300 * time.Millisecond)
	if n := cp.n(); n != 0 {
		t.Fatalf("fantome agent parti (%d appels) : execution speculative d'outils", n)
	}
}
