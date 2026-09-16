package chat

import (
	"context"
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/local"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/store"
)

type fakeProvider struct {
	id      string
	fail    map[string]bool
	content map[string]string
}

func (f *fakeProvider) ID() string { return f.id }

func (f *fakeProvider) Stream(ctx context.Context, req provider.Request, emit func(provider.Event) bool) (provider.Response, error) {
	if f.fail[req.Model] {
		return provider.Response{}, errors.New("panne simulee " + req.Model)
	}
	txt := f.content[req.Model]
	for _, r := range txt {
		if !emit(provider.Event{Content: string(r)}) {
			break
		}
	}
	return provider.Response{Content: txt}, nil
}

func newEngine(t *testing.T, fp *fakeProvider, fams []alias.Family) *Engine {
	t.Helper()
	st, err := store.Open(filepath.Join(t.TempDir(), "chat.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	reg := provider.NewRegistry()
	reg.Set(fp)
	return NewEngine(reg, fams, st, nil, t.TempDir())
}

func codeFamily(pool ...alias.Member) []alias.Family {
	return []alias.Family{{
		ID: "code", Label: "Code",
		Modes: []alias.Mode{{ID: "standard", Agent: true, Pool: pool}},
	}}
}

func runTurn(t *testing.T, e *Engine, user string, in TurnInput) *Conversation {
	t.Helper()
	c := e.Conversation(user)
	if err := c.StartTurn(in); err != nil {
		t.Fatalf("StartTurn: %v", err)
	}
	waitFor(t, func() bool { return !c.IsGenerating() }, "tour non termine")
	return c
}

func logText(c *Conversation) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	var b strings.Builder
	for _, ev := range c.Log {
		if s, ok := ev.Delta["content"].(string); ok {
			b.WriteString(s)
		}
	}
	return b.String()
}

func lastError(c *Conversation) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i := len(c.Log) - 1; i >= 0; i-- {
		if s, ok := c.Log[i].Delta["error"].(string); ok {
			return s
		}
	}
	return ""
}

func TestEngineStreamsContent(t *testing.T) {
	fp := &fakeProvider{id: "fake", content: map[string]string{"ok": "Bonjour le monde"}}
	e := newEngine(t, fp, codeFamily(alias.Member{Provider: "fake", Model: "ok"}))
	c := runTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "salut", AgentMode: true})
	if got := logText(c); got != "Bonjour le monde" {
		t.Fatalf("contenu = %q", got)
	}
	if err := lastError(c); err != "" {
		t.Fatalf("erreur inattendue: %s", err)
	}
}

func TestEngineFailover(t *testing.T) {
	fp := &fakeProvider{
		id:      "fake",
		fail:    map[string]bool{"bad": true},
		content: map[string]string{"ok": "secours"},
	}
	e := newEngine(t, fp, codeFamily(
		alias.Member{Provider: "fake", Model: "bad"},
		alias.Member{Provider: "fake", Model: "ok"},
	))
	c := runTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "salut", AgentMode: true})
	if got := logText(c); got != "secours" {
		t.Fatalf("failover contenu = %q", got)
	}
	if err := lastError(c); err != "" {
		t.Fatalf("failover ne doit pas laisser d'erreur: %s", err)
	}
}

func TestEngineAllFail(t *testing.T) {
	fp := &fakeProvider{id: "fake", fail: map[string]bool{"bad": true, "worse": true}}
	e := newEngine(t, fp, codeFamily(
		alias.Member{Provider: "fake", Model: "bad"},
		alias.Member{Provider: "fake", Model: "worse"},
	))
	c := runTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "salut", AgentMode: true})
	if err := lastError(c); err == "" {
		t.Fatal("une erreur doit etre remontee")
	}
}

func TestResolveLocalFallbackWhenEngineDown(t *testing.T) {
	fp := &fakeProvider{id: "fake"}
	e := newEngine(t, fp, alias.Defaults())
	res := e.resolve(context.Background(), TurnInput{Family: "samgen", Mode: "n4"})
	if len(res.members) == 0 {
		t.Fatal("le repli cloud doit fournir un pool")
	}
	if !res.fallback {
		t.Fatal("le repli doit etre marque fallback")
	}
	if res.members[0].Provider != "openrouter" {
		t.Fatalf("repli attendu openrouter, got %q", res.members[0].Provider)
	}
}

func TestResolveNanoShufflesPool(t *testing.T) {
	fp := &fakeProvider{id: "fake"}
	e := newEngine(t, fp, alias.Defaults())
	in := TurnInput{Family: "samagent-nano", Mode: "free"}
	ref, ok := alias.Resolve(alias.Defaults(), "samagent-nano", "free")
	if !ok {
		t.Fatal("nano free introuvable")
	}
	firsts := map[string]bool{}
	for i := 0; i < 30; i++ {
		res := e.resolve(context.Background(), in)
		if len(res.members) != len(ref.Pool) {
			t.Fatalf("resolve %d: %d membres, want %d", i, len(res.members), len(ref.Pool))
		}
		// permutation : meme ensemble, sans perte ni doublon
		seen := map[string]bool{}
		for _, m := range res.members {
			k := m.Provider + "/" + m.Model
			if seen[k] {
				t.Fatalf("resolve %d: doublon %s", i, k)
			}
			seen[k] = true
		}
		for _, m := range ref.Pool {
			if !seen[m.Provider+"/"+m.Model] {
				t.Fatalf("resolve %d: %s/%s perdu", i, m.Provider, m.Model)
			}
		}
		firsts[res.members[0].Provider+"/"+res.members[0].Model] = true
	}
	if len(firsts) < 2 {
		t.Fatal("le pool nano doit etre melange : 30 tirages, une seule tete")
	}
}

func TestResolveNonNanoKeepsOrder(t *testing.T) {
	// Regle du selecteur de modeles : tout pool de >1 membre est melange
	// a chaque requete (fallback aleatoire) ; un pool d'un seul membre
	// reste fixe, sans fallback.
	fp := &fakeProvider{id: "fake"}
	e := newEngine(t, fp, alias.Defaults())

	// N8 flash = 1 membre : ordre et contenu inchanges.
	ref1, _ := alias.Resolve(alias.Defaults(), "samagent-n8", "flash")
	res1 := e.resolve(context.Background(), TurnInput{Family: "samagent-n8", Mode: "flash"})
	if len(res1.members) != 1 || res1.members[0].Model != ref1.Pool[0].Model {
		t.Fatalf("pool singleton modifie : %+v", res1.members)
	}

	// N4 flash = 14 membres : melange (permutation, tete variable).
	refN, _ := alias.Resolve(alias.Defaults(), "samagent-n4", "flash")
	firsts := map[string]bool{}
	for i := 0; i < 30; i++ {
		res := e.resolve(context.Background(), TurnInput{Family: "samagent-n4", Mode: "flash"})
		if len(res.members) != len(refN.Pool) {
			t.Fatalf("resolve %d: %d membres, want %d", i, len(res.members), len(refN.Pool))
		}
		seen := map[string]bool{}
		for _, m := range res.members {
			k := m.Provider + "/" + m.Model
			if seen[k] {
				t.Fatalf("resolve %d: doublon %s", i, k)
			}
			seen[k] = true
		}
		for _, m := range refN.Pool {
			if !seen[m.Provider+"/"+m.Model] {
				t.Fatalf("resolve %d: %s/%s perdu", i, m.Provider, m.Model)
			}
		}
		firsts[res.members[0].Provider+"/"+res.members[0].Model] = true
	}
	if len(firsts) < 2 {
		t.Fatal("le pool n4 flash doit etre melange : 30 tirages, une seule tete")
	}
}

// fakeDiscoverer simule la decouverte de modeles locaux.
type fakeDiscoverer struct {
	models map[string][]string
}

func (f *fakeDiscoverer) ModelsForEngine(ctx context.Context, engine string) []local.Model {
	var out []local.Model
	for _, id := range f.models[engine] {
		out = append(out, local.Model{Engine: engine, ID: id})
	}
	return out
}

func newEngineWithDiscoverer(t *testing.T, fams []alias.Family, disc LocalDiscoverer) *Engine {
	t.Helper()
	st, err := store.Open(filepath.Join(t.TempDir(), "chat.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	reg := provider.NewRegistry()
	return NewEngine(reg, fams, st, disc, t.TempDir())
}

func TestResolveLocalModelSelection(t *testing.T) {
	disc := &fakeDiscoverer{models: map[string][]string{
		"ollama": {"llama3.1:8b", "qwen2.5:14b", "mistral:7b"},
	}}
	samgen := []alias.Family{{
		ID: "samgen", Label: "SamGen", Local: true,
		Modes: []alias.Mode{{ID: "n4", Label: "N4 (Ollama)", Local: true, Engine: "ollama"}},
	}}

	// Sans selection : tous les modeles decouverts, melanges (fallback).
	e := newEngineWithDiscoverer(t, samgen, disc)
	res := e.resolve(context.Background(), TurnInput{Family: "samgen", Mode: "n4"})
	if len(res.members) != 3 || !res.local {
		t.Fatalf("sans selection: 3 modeles locaux attendus, got %+v", res.members)
	}

	// Selection "1 modele" : un seul membre, fixe.
	sel := []alias.Family{{
		ID: "samgen", Label: "SamGen", Local: true,
		Modes: []alias.Mode{{ID: "n4", Label: "N4 (Ollama)", Local: true, Engine: "ollama",
			Pool: []alias.Member{{Provider: "ollama", Model: "qwen2.5:14b"}}}},
	}}
	e2 := newEngineWithDiscoverer(t, sel, disc)
	res2 := e2.resolve(context.Background(), TurnInput{Family: "samgen", Mode: "n4"})
	if len(res2.members) != 1 || res2.members[0].Model != "qwen2.5:14b" {
		t.Fatalf("selection 1 modele: got %+v", res2.members)
	}

	// Selection fallback : sous-ensemble, melange.
	sel2 := []alias.Family{{
		ID: "samgen", Label: "SamGen", Local: true,
		Modes: []alias.Mode{{ID: "n4", Label: "N4 (Ollama)", Local: true, Engine: "ollama",
			Pool: []alias.Member{
				{Provider: "ollama", Model: "llama3.1:8b"},
				{Provider: "ollama", Model: "mistral:7b"},
			}}},
	}}
	e3 := newEngineWithDiscoverer(t, sel2, disc)
	firsts := map[string]bool{}
	for i := 0; i < 20; i++ {
		r := e3.resolve(context.Background(), TurnInput{Family: "samgen", Mode: "n4"})
		if len(r.members) != 2 {
			t.Fatalf("fallback local: 2 membres attendus, got %d", len(r.members))
		}
		for _, m := range r.members {
			if m.Model == "qwen2.5:14b" {
				t.Fatalf("fallback local: modele non selectionne present : %s", m.Model)
			}
		}
		firsts[r.members[0].Model] = true
	}
	if len(firsts) < 2 {
		t.Fatal("le fallback local doit etre melange")
	}
}

func TestConversationPersistence(t *testing.T) {
	fp := &fakeProvider{id: "fake", content: map[string]string{"ok": "memoire"}}
	st, err := store.Open(filepath.Join(t.TempDir(), "persist.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer st.Close()
	reg := provider.NewRegistry()
	reg.Set(fp)
	fams := codeFamily(alias.Member{Provider: "fake", Model: "ok"})

	e1 := NewEngine(reg, fams, st, nil, t.TempDir())
	c1 := runTurn(t, e1, "sam", TurnInput{Family: "code", Mode: "standard", Text: "salut", AgentMode: true})
	if got := logText(c1); got != "memoire" {
		t.Fatalf("contenu = %q", got)
	}

	e2 := NewEngine(reg, fams, st, nil, t.TempDir())
	c2 := e2.Conversation("sam")
	if got := logText(c2); got == "" {
		t.Fatal("conversation non rechargee depuis le store")
	}
}

// TestResolveAutoUnionCloud : le mode "auto" unit les pools effectifs de
// tous les modes de la famille, dedupliques, melanges a chaque requete.
func TestResolveAutoUnionCloud(t *testing.T) {
	fams := []alias.Family{{
		ID: "samagent-n4", Label: "N4",
		Modes: []alias.Mode{
			{ID: "flash", Label: "Flash", Pool: []alias.Member{
				{Provider: "deepseek", Model: "deepseek-flash"},
				{Provider: "openrouter", Model: "mimo-v2.5-flash"},
			}},
			{ID: "standard", Label: "Standard", Pool: []alias.Member{
				{Provider: "openrouter", Model: "mimo-v2.5-flash"}, // doublon volontaire
				{Provider: "openrouter", Model: "glm-5-flash"},
			}},
		},
	}}
	e := newEngineWithDiscoverer(t, fams, nil)
	in := TurnInput{Family: "samagent-n4", Mode: "auto"}
	seen := map[string]int{}
	for i := 0; i < 30; i++ {
		res := e.resolve(context.Background(), in)
		if len(res.members) != 3 {
			t.Fatalf("auto doit unir les pools dedupliques : %d membres", len(res.members))
		}
		if !res.fallback {
			t.Fatal("auto multi-membres doit etre signale comme fallback")
		}
		keys := map[string]bool{}
		for _, m := range res.members {
			k := m.Provider + "/" + m.Model
			if keys[k] {
				t.Fatalf("doublon dans l'union auto : %s", k)
			}
			keys[k] = true
			seen[m.Model]++
		}
	}
	for _, want := range []string{"deepseek-flash", "mimo-v2.5-flash", "glm-5-flash"} {
		if seen[want] == 0 {
			t.Fatalf("modele manquant dans l'union auto : %s", want)
		}
	}
}

// TestResolveAutoSingleMember : un seul membre au total = pas de fallback.
func TestResolveAutoSingleMember(t *testing.T) {
	fams := []alias.Family{{
		ID: "samagent-nano", Label: "Nano",
		Modes: []alias.Mode{
			{ID: "free", Label: "Free", Pool: []alias.Member{{Provider: "openrouter", Model: "openrouter/free"}}},
		},
	}}
	e := newEngineWithDiscoverer(t, fams, nil)
	res := e.resolve(context.Background(), TurnInput{Family: "samagent-nano", Mode: "auto"})
	if len(res.members) != 1 || res.fallback {
		t.Fatalf("auto singleton : 1 membre, pas de fallback (got %d, fallback=%v)", len(res.members), res.fallback)
	}
	if res.members[0].Model != "openrouter/free" {
		t.Fatalf("membre inattendu : %s", res.members[0].Model)
	}
}

// TestResolveAutoUnknownFamily : famille inconnue -> resolution vide.
func TestResolveAutoUnknownFamily(t *testing.T) {
	e := newEngineWithDiscoverer(t, nil, nil)
	res := e.resolve(context.Background(), TurnInput{Family: "nope", Mode: "auto"})
	if len(res.members) != 0 {
		t.Fatalf("famille inconnue : resolution vide attendue, got %d", len(res.members))
	}
}

// TestResolveAutoLocal : auto sur famille locale = union des modeles
// decouverts sur tous les moteurs, selection du selecteur honoree.
func TestResolveAutoLocal(t *testing.T) {
	disc := &fakeDiscoverer{models: map[string][]string{
		"ollama":   {"llama3.1:8b", "qwen2.5:14b"},
		"llamacpp": {"qwen2.5:14b", "mistral:7b"}, // qwen2.5:14b en doublon inter-moteurs
	}}
	samgen := []alias.Family{{
		ID: "samgen", Label: "SamGen", Local: true,
		Modes: []alias.Mode{
			{ID: "n4", Label: "N4 (Ollama)", Local: true, Engine: "ollama",
				Pool: []alias.Member{{Provider: "ollama", Model: "llama3.1:8b"}}},
			{ID: "nano", Label: "Nano (llama.cpp)", Local: true, Engine: "llamacpp"},
		},
	}}
	e := newEngineWithDiscoverer(t, samgen, disc)
	res := e.resolve(context.Background(), TurnInput{Family: "samgen", Mode: "auto"})
	if !res.local {
		t.Fatal("auto local doit etre marque local")
	}
	// Attendu : ollama/llama3.1:8b (selection) + llamacpp/qwen2.5:14b + llamacpp/mistral:7b.
	// "ollama/qwen2.5:14b" est exclu par la selection du mode n4.
	if len(res.members) != 3 {
		t.Fatalf("union locale attendue : 3 membres, got %d", len(res.members))
	}
	keys := map[string]bool{}
	for _, m := range res.members {
		keys[m.Provider+"/"+m.Model] = true
	}
	for _, want := range []string{"ollama/llama3.1:8b", "llamacpp/qwen2.5:14b", "llamacpp/mistral:7b"} {
		if !keys[want] {
			t.Fatalf("membre manquant dans l'union locale : %s", want)
		}
	}
	if keys["ollama/qwen2.5:14b"] {
		t.Fatal("la selection du selecteur doit etre honoree en mode auto local")
	}
}
