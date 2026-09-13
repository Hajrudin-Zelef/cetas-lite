package chat

import (
	"context"
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"cetas-lite/internal/alias"
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
	return NewEngine(reg, fams, st, nil)
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
	c := runTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "salut"})
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
	c := runTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "salut"})
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
	c := runTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "salut"})
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

	e1 := NewEngine(reg, fams, st, nil)
	c1 := runTurn(t, e1, "sam", TurnInput{Family: "code", Mode: "standard", Text: "salut"})
	if got := logText(c1); got != "memoire" {
		t.Fatalf("contenu = %q", got)
	}

	e2 := NewEngine(reg, fams, st, nil)
	c2 := e2.Conversation("sam")
	if got := logText(c2); got == "" {
		t.Fatal("conversation non rechargee depuis le store")
	}
}
