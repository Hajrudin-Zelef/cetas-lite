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

// toolProvider simule un modele qui emet d'abord un tool call Read,
// puis repond avec le contenu lu.
type toolProvider struct {
	mu        sync.Mutex
	calls     int
	sawResult bool
}

func (p *toolProvider) ID() string { return "toolprov" }

func (p *toolProvider) Stream(ctx context.Context, req provider.Request, emit func(provider.Event) bool) (provider.Response, error) {
	p.mu.Lock()
	n := p.calls
	p.calls++
	p.mu.Unlock()
	// Verifier que le resultat d'outil precedent est bien present dans les messages.
	if n == 1 {
		found := false
		for _, m := range req.Messages {
			if m.Role == "tool" {
				if s, ok := m.Content.(string); ok && strings.Contains(s, "CONTENU-SECRET-42") {
					found = true
				}
			}
		}
		p.sawResult = found
		if !found {
			return provider.Response{Content: "ERREUR: resultat d'outil manquant"}, nil
		}
		return provider.Response{Content: "J'ai lu le fichier avec succes."}, nil
	}
	tc := provider.ToolCall{ID: "call_1", Type: "function"}
	tc.Function.Name = "Read"
	tc.Function.Arguments = `{"file_path":"secret.txt"}`
	return provider.Response{
		Content:   "",
		ToolCalls: []provider.ToolCall{tc},
	}, nil
}

func TestToolCallEndToEnd(t *testing.T) {
	ws := t.TempDir()
	// Le sandbox de l'agent est <workspace>/<user>/.
	userDir := filepath.Join(ws, "u")
	if err := os.MkdirAll(userDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(userDir, "secret.txt"), []byte("CONTENU-SECRET-42"), 0o644); err != nil {
		t.Fatal(err)
	}
	tp := &toolProvider{}
	st, err := store.Open(filepath.Join(t.TempDir(), "chat.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	reg := provider.NewRegistry()
	reg.Set(tp)
	e := NewEngine(reg, codeFamily(alias.Member{Provider: "toolprov", Model: "m"}), st, nil, ws)
	c := runTurn(t, e, "u", TurnInput{User: "u", Family: "code", Mode: "standard", Text: "lis secret.txt", Approve: false, AgentMode: true})
	tp.mu.Lock()
	sawResult := tp.sawResult
	calls := tp.calls
	tp.mu.Unlock()
	if calls != 2 {
		t.Fatalf("appels provider = %d, attendu 2 (tool call puis reponse finale)", calls)
	}
	if !sawResult {
		t.Fatal("le resultat de l'outil Read n'a pas ete reinjecte dans les messages du 2e appel")
	}
	// Verifier qu'un evenement tool start/end a ete emis.
	c.mu.Lock()
	defer c.mu.Unlock()
	sawTool := false
	for _, ev := range c.Log {
		if tm, ok := ev.Delta["tool"].(map[string]any); ok {
			if tm["name"] == "Read" {
				sawTool = true
			}
		}
	}
	if !sawTool {
		t.Fatal("aucun evenement d'outil Read dans le log")
	}
}
