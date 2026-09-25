package chat

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/store"
)

// emptyAfterToolsProvider simule le cas observe en prod : le modele emet
// d'abord un tool call (Read), puis renvoie un contenu vide a plusieurs
// reprises. Sans le dernier recours (ultime tentative outils coupes), le
// tour persistait un message assistant vide que l'utilisateur devait
// regenerer a la main.
type emptyAfterToolsProvider struct {
	mu         sync.Mutex
	calls      int
	emptyTimes int // reponses vides (apres le tool call) avant la reponse finale
	toolsSeen  []int
}

func (p *emptyAfterToolsProvider) ID() string { return "emptyprov" }

func (p *emptyAfterToolsProvider) Stream(ctx context.Context, req provider.Request, emit func(provider.Event) bool) (provider.Response, error) {
	p.mu.Lock()
	n := p.calls
	p.calls++
	p.toolsSeen = append(p.toolsSeen, len(req.Tools))
	p.mu.Unlock()
	if n == 0 {
		tc := provider.ToolCall{ID: "call_1", Type: "function"}
		tc.Function.Name = "Read"
		tc.Function.Arguments = `{"file_path":"secret.txt"}`
		return provider.Response{Content: "", ToolCalls: []provider.ToolCall{tc}}, nil
	}
	if n-1 < p.emptyTimes {
		return provider.Response{Content: "   "}, nil
	}
	return provider.Response{Content: "Voici le contenu lu : CONTENU-SECRET-42."}, nil
}

func newEmptyAfterToolsEngine(t *testing.T, emptyTimes int) (*Engine, *emptyAfterToolsProvider) {
	t.Helper()
	ws := t.TempDir()
	userDir := filepath.Join(ws, "u")
	if err := os.MkdirAll(userDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(userDir, "secret.txt"), []byte("CONTENU-SECRET-42"), 0o644); err != nil {
		t.Fatal(err)
	}
	ep := &emptyAfterToolsProvider{emptyTimes: emptyTimes}
	st, err := store.Open(filepath.Join(t.TempDir(), "chat.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	reg := provider.NewRegistry()
	reg.Set(ep)
	return NewEngine(reg, codeFamily(alias.Member{Provider: "emptyprov", Model: "m"}), st, nil, ws), ep
}

func lastAssistantText(c *Conversation) string {
	msgs := c.MessagesSnapshot()
	for i := len(msgs) - 1; i >= 0; i-- {
		if msgs[i].Role == "assistant" {
			if s, ok := msgs[i].Content.(string); ok {
				return s
			}
		}
	}
	return ""
}

// Apres epuisement des nudges "reponse vide", le moteur fait une ultime
// tentative outils coupes au lieu de persister un message vide.
func TestEmptyAfterToolsRecovers(t *testing.T) {
	e, ep := newEmptyAfterToolsEngine(t, 3)
	c := runTurn(t, e, "u", TurnInput{User: "u", Family: "code", Mode: "standard", Text: "lis secret.txt", Approve: false, AgentMode: true})
	ep.mu.Lock()
	calls := ep.calls
	toolsSeen := append([]int(nil), ep.toolsSeen...)
	ep.mu.Unlock()
	// 1 tool call + 2 nudges + 1 ultime tentative = 5 appels provider.
	if calls != 5 {
		t.Fatalf("appels provider = %d, attendu 5 (tool call, 2 nudges, ultime tentative, reponse)", calls)
	}
	if len(toolsSeen) != 5 || toolsSeen[4] != 0 {
		t.Fatalf("la derniere requete doit partir sans outils, outils vus par appel = %v", toolsSeen)
	}
	if got := lastAssistantText(c); !strings.Contains(got, "Voici le contenu lu") {
		t.Fatalf("message assistant final = %q, attendu non vide avec la reponse", got)
	}
}

// Si le modele reste vide meme sans outils, le tour se termine proprement
// (note honnete) sans boucle infinie : borne a 5 appels provider.
func TestEmptyAfterToolsBounded(t *testing.T) {
	e, ep := newEmptyAfterToolsEngine(t, 100)
	c := runTurn(t, e, "u", TurnInput{User: "u", Family: "code", Mode: "standard", Text: "lis secret.txt", Approve: false, AgentMode: true})
	ep.mu.Lock()
	calls := ep.calls
	ep.mu.Unlock()
	if calls != 5 {
		t.Fatalf("appels provider = %d, attendu 5 bornes (tool call, 2 nudges, ultime tentative, abandon)", calls)
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	sawNote := false
	for _, ev := range c.Log {
		if s, ok := ev.Delta["content"].(string); ok && strings.Contains(s, "n'a pas produit de reponse") {
			sawNote = true
		}
	}
	if !sawNote {
		t.Fatal("note honnete _(le modele n'a pas produit de reponse)_ absente du log")
	}
}
