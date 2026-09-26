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
	"cetas-lite/internal/store"
)

// multiCaptureProvider enregistre TOUTES les requetes (le captureProvider
// existant ne garde que la derniere).
type multiCaptureProvider struct {
	mu   sync.Mutex
	reqs []provider.Request
}

func (c *multiCaptureProvider) ID() string { return "cap" }

func (c *multiCaptureProvider) Stream(_ context.Context, req provider.Request, emit func(provider.Event) bool) (provider.Response, error) {
	c.mu.Lock()
	c.reqs = append(c.reqs, req)
	c.mu.Unlock()
	emit(provider.Event{Content: "ok"})
	return provider.Response{Content: "ok"}, nil
}

// commonPrefixLen : longueur du prefixe commun a deux chaines.
func commonPrefixLen(a, b string) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return i
		}
	}
	return n
}

// analyzePrefix rejoue 2 tours et renvoie les requetes envoyees au moteur
// (le cache de prefixe llama.cpp reutilise un prefixe token-a-token identique).
func analyzePrefix(t *testing.T, r RagTools) ([]provider.Message, []provider.Message) {
	t.Helper()
	st, err := store.Open(filepath.Join(t.TempDir(), "prefix.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	reg := provider.NewRegistry()
	cp := &multiCaptureProvider{}
	reg.Set(cp)
	e := NewEngine(reg, nil, st, nil, t.TempDir())
	if r != nil {
		e.SetRAG(r)
	}
	e.SetFamilies([]alias.Family{{
		ID: "code", Label: "Code",
		Modes: []alias.Mode{{ID: "standard", Pool: []alias.Member{{Provider: "cap", Model: "m"}}}},
	}})
	c := e.Conversation("u")
	for _, text := range []string{"premiere question sur le RAID", "deuxieme question sur le noyau"} {
		if err := c.StartTurn(TurnInput{User: "u", Family: "code", Mode: "standard", Text: text}); err != nil {
			t.Fatalf("StartTurn: %v", err)
		}
		waitFor(t, func() bool { return !c.IsGenerating() }, "tour non termine")
	}
	cp.mu.Lock()
	defer cp.mu.Unlock()
	if len(cp.reqs) != 2 {
		t.Fatalf("2 requetes attendues, %d", len(cp.reqs))
	}
	return cp.reqs[0].Messages, cp.reqs[1].Messages
}

// TestPromptSystemPrefixStable : sur deux tours, le message systeme de tete
// doit etre identique (prefixe stable) malgre l'historique qui grandit et
// le RAG dynamique. Condition du cache de prefixe llama.cpp.
func TestPromptSystemPrefixStable(t *testing.T) {
	r := fakeRag{ready: true, hits: []rag.Hit{
		{Title: "Alpha", Path: "a.md", Excerpt: "EXTRAIT-DYNAMIQUE-" + strings.Repeat("x", 40), Score: ragStrongScore},
	}}
	a, b := analyzePrefix(t, r)
	if len(a) == 0 || len(b) == 0 {
		t.Fatal("requetes vides")
	}
	if a[0].Role != "system" || b[0].Role != "system" {
		t.Fatalf("1er message non-systeme: %s / %s", a[0].Role, b[0].Role)
	}
	sa, sb := msgText(a[0]), msgText(b[0])
	if !strings.Contains(sa, "You are Cetas") {
		t.Fatalf("prompt systeme stable absent en tete: %.100q", sa)
	}
	common := commonPrefixLen(sa, sb)
	t.Logf("systeme t1=%d car, t2=%d car, prefixe commun=%d car", len(sa), len(sb), common)
	if common < len("You are Cetas, a senior teacher") {
		t.Fatalf("prefixe systeme casse entre deux tours (commun=%d)\n  t1=%.160q\n  t2=%.160q", common, sa, sb)
	}
}

// TestPromptHistoryPrefixGrows : le prefixe rejoue du tour 2 doit couvrir
// le tour 1 (systeme + historique + RAG du tour 1), sinon le cache est
// inutilisable a chaque nouveau message.
func TestPromptHistoryPrefixGrows(t *testing.T) {
	r := fakeRag{ready: true, hits: []rag.Hit{
		{Title: "Alpha", Path: "a.md", Excerpt: "EXTRAIT", Score: ragStrongScore},
	}}
	a, b := analyzePrefix(t, r)
	if len(b) <= len(a) {
		t.Fatalf("tour 2 (%d msgs) doit etre plus long que tour 1 (%d msgs)", len(b), len(a))
	}
	for i := 0; i < len(a); i++ {
		if a[i].Role != b[i].Role {
			t.Fatalf("role diverge au message %d: %s vs %s", i, a[i].Role, b[i].Role)
		}
		if a[i].Role != "system" && msgText(a[i]) != msgText(b[i]) {
			t.Fatalf("message %d (%s) different:\n  t1=%.120q\n  t2=%.120q", i, a[i].Role, msgText(a[i]), msgText(b[i]))
		}
	}
	t.Logf("prefixe rejoue: %d messages identiques", len(a))
}
