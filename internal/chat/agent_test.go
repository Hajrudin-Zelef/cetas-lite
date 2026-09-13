package chat

import (
	"context"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/store"
)

type scriptStep struct {
	content   string
	toolCalls []provider.ToolCall
	err       error
}

type scriptedProvider struct {
	id    string
	mu    sync.Mutex
	steps []scriptStep
	idx   int
	reqs  []provider.Request
}

func (s *scriptedProvider) ID() string { return s.id }

func (s *scriptedProvider) Stream(ctx context.Context, req provider.Request, emit func(provider.Event) bool) (provider.Response, error) {
	s.mu.Lock()
	s.reqs = append(s.reqs, req)
	i := s.idx
	if i >= len(s.steps) {
		i = len(s.steps) - 1
	}
	s.idx++
	step := s.steps[i]
	s.mu.Unlock()
	if step.err != nil {
		return provider.Response{}, step.err
	}
	for _, r := range step.content {
		if !emit(provider.Event{Content: string(r)}) {
			break
		}
	}
	return provider.Response{Content: step.content, ToolCalls: step.toolCalls}, nil
}

func (s *scriptedProvider) requests() []provider.Request {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]provider.Request(nil), s.reqs...)
}

func newAgentEngine(t *testing.T, sp *scriptedProvider, fams []alias.Family) *Engine {
	t.Helper()
	st, err := store.Open(filepath.Join(t.TempDir(), "agent.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	reg := provider.NewRegistry()
	reg.Set(sp)
	return NewEngine(reg, fams, st, nil, t.TempDir())
}

func runAgentTurn(t *testing.T, e *Engine, user string, in TurnInput) *Conversation {
	t.Helper()
	in.User = user
	c := e.Conversation(user)
	if err := c.StartTurn(in); err != nil {
		t.Fatalf("StartTurn: %v", err)
	}
	waitFor(t, func() bool { return !c.IsGenerating() }, "tour non termine")
	return c
}

func toolCall(id, name, args string) provider.ToolCall {
	return provider.ToolCall{ID: id, Type: "function", Function: provider.Func{Name: name, Arguments: args}}
}

func hasToolEvent(c *Conversation, phase string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, ev := range c.Log {
		t, ok := ev.Delta["tool"].(map[string]any)
		if ok && t["phase"] == phase {
			return true
		}
	}
	return false
}

func countEvents(c *Conversation, key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	n := 0
	for _, ev := range c.Log {
		if _, ok := ev.Delta[key]; ok {
			n++
		}
	}
	return n
}

func TestAgentExecutesToolThenAnswers(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{toolCall("c1", "TodoWrite", `{"todos":[{"content":"a","status":"pending"}]}`)}},
		{content: "Termine"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "fais"})

	if got := logText(c); !strings.Contains(got, "Termine") {
		t.Fatalf("contenu final = %q", got)
	}
	if !hasToolEvent(c, "start") || !hasToolEvent(c, "end") {
		t.Fatal("evenements tool start/end attendus")
	}
	reqs := sp.requests()
	if len(reqs) != 2 {
		t.Fatalf("requetes = %d, want 2", len(reqs))
	}
	foundTool := false
	for _, m := range reqs[1].Messages {
		if m.Role == "tool" {
			foundTool = true
		}
	}
	if !foundTool {
		t.Fatal("le resultat d'outil doit etre renvoye au modele")
	}
}

func TestAgentDedupNonBash(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{
			toolCall("c1", "Write", `{"file_path":"a.txt","content":"x"}`),
			toolCall("c2", "Write", `{"file_path":"a.txt","content":"x"}`),
		}},
		{content: "ok"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "ecris"})

	reqs := sp.requests()
	found := false
	for _, m := range reqs[1].Messages {
		if m.Role == "tool" {
			if s, _ := m.Content.(string); strings.HasPrefix(s, "[deja fait]") {
				found = true
			}
		}
	}
	if !found {
		t.Fatal("le second appel identique doit etre deduplique")
	}
}

func TestAgentBashNotDeduped(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{
			toolCall("c1", "Bash", `{"command":"echo x"}`),
			toolCall("c2", "Bash", `{"command":"echo x"}`),
		}},
		{content: "ok"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "echo"})

	reqs := sp.requests()
	for _, m := range reqs[1].Messages {
		if m.Role == "tool" {
			if s, _ := m.Content.(string); strings.HasPrefix(s, "[deja fait]") {
				t.Fatal("Bash ne doit jamais etre deduplique")
			}
		}
	}
}

func TestAgentNudgeThenAnswer(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{},
		{},
		{content: "enfin"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "?"})

	if got := logText(c); !strings.Contains(got, "enfin") {
		t.Fatalf("contenu = %q", got)
	}
	if n := countEvents(c, "drop_reasoning"); n != 2 {
		t.Fatalf("drop_reasoning = %d, want 2", n)
	}
	if len(sp.requests()) != 3 {
		t.Fatalf("requetes = %d, want 3", len(sp.requests()))
	}
}

func TestAgentRetryWithoutToolsOnHTTPError(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{err: &provider.HTTPError{Provider: "fake", Status: 500, Body: "boom"}},
		{content: "sans outils"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "?"})

	if got := logText(c); !strings.Contains(got, "sans outils") {
		t.Fatalf("contenu = %q", got)
	}
	reqs := sp.requests()
	if len(reqs) != 2 {
		t.Fatalf("requetes = %d, want 2", len(reqs))
	}
	if len(reqs[1].Tools) != 0 {
		t.Fatalf("le retry doit desactiver les outils, tools = %d", len(reqs[1].Tools))
	}
}

func TestNonAgentModeHasNoTools(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "salut"}}}
	fams := []alias.Family{{ID: "plain", Label: "Plain", Modes: []alias.Mode{{ID: "standard", Pool: []alias.Member{{Provider: "fake", Model: "m"}}}}}}
	e := newAgentEngine(t, sp, fams)
	runAgentTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "salut"})

	reqs := sp.requests()
	if len(reqs) != 1 {
		t.Fatalf("requetes = %d, want 1", len(reqs))
	}
	if len(reqs[0].Tools) != 0 {
		t.Fatalf("mode non-agent ne doit pas porter d'outils, tools = %d", len(reqs[0].Tools))
	}
	for _, m := range reqs[0].Messages {
		if m.Role == "system" {
			t.Fatalf("mode non-agent ne doit pas avoir de prompt systeme: %v", m.Content)
		}
	}
}
