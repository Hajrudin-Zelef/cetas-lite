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

type scriptStep struct {
	content   string
	reasoning string
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
	if step.reasoning != "" {
		if !emit(provider.Event{Reasoning: step.reasoning}) {
			return provider.Response{Reasoning: step.reasoning}, nil
		}
	}
	for _, r := range step.content {
		if !emit(provider.Event{Content: string(r)}) {
			break
		}
	}
	return provider.Response{Content: step.content, Reasoning: step.reasoning, ToolCalls: step.toolCalls}, nil
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
	in.AgentMode = true // les tests agent simulent le mode "Agent" de l'UI
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

func TestAgentRetries429WithTools(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{err: &provider.HTTPError{Provider: "fake", Status: 429, Body: "rate limit"}},
		{content: "apres retry"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "?"})

	if got := logText(c); !strings.Contains(got, "apres retry") {
		t.Fatalf("contenu = %q", got)
	}
	reqs := sp.requests()
	if len(reqs) != 2 {
		t.Fatalf("requetes = %d, want 2", len(reqs))
	}
	for i, r := range reqs {
		if len(r.Tools) == 0 {
			t.Fatalf("requete %d : le retry 429 doit garder les outils", i)
		}
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
	if !hasSystemContaining(reqs, "senior teacher") {
		t.Fatal("mode non-agent doit porter le prompt systeme de chat")
	}
}

func TestChatModeDisablesAgentPath(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{toolCall("c1", "Read", `{"file_path":"a.txt"}`)}},
		{content: "Termine"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	// Famille compatible agent mais mode Chat (AgentMode=false) : aucun outil.
	c := e.Conversation("sam")
	in := TurnInput{User: "sam", Family: "code", Mode: "standard", Text: "lis a.txt", AgentMode: false}
	if err := c.StartTurn(in); err != nil {
		t.Fatalf("StartTurn: %v", err)
	}
	waitFor(t, func() bool { return !c.IsGenerating() }, "tour non termine")
	if hasToolEvent(c, "start") || hasToolEvent(c, "end") {
		t.Fatal("aucun outil ne devrait etre execute en mode Chat")
	}
	if got := len(sp.requests()); got != 1 {
		t.Fatalf("requetes = %d, attendu 1 (pas de second appel avec resultat d'outil)", got)
	}
}

func TestAgentParallelReads(t *testing.T) {
	ws := t.TempDir()
	// Le sandbox de l'agent est <workspace>/<user>.
	sandboxDir := filepath.Join(ws, "sam")
	if err := os.MkdirAll(sandboxDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sandboxDir, "a.txt"), []byte("contenu A"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sandboxDir, "b.txt"), []byte("contenu B"), 0o644); err != nil {
		t.Fatal(err)
	}
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{
			toolCall("c1", "Read", `{"file_path":"a.txt"}`),
			toolCall("c2", "Read", `{"file_path":"b.txt"}`),
			// Phase 3 : c3 utilise l'ancien nom Cat — le filet de
			// sécurité le réécrit en Read, en parallèle comme une
			// lecture pure.
			toolCall("c3", "Cat", `{"file_path":"a.txt"}`),
		}},
		{content: "Termine"},
	}}
	st, err := store.Open(filepath.Join(t.TempDir(), "agent.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	reg := provider.NewRegistry()
	reg.Set(sp)
	e := NewEngine(reg, codeFamily(alias.Member{Provider: "fake", Model: "m"}), st, nil, ws)
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "lis"})

	// Les 3 lectures du bloc ont été exécutées : résultats présents dans
	// les deltas "end".
	c.mu.Lock()
	var toolResults strings.Builder
	for _, ev := range c.Log {
		if tm, ok := ev.Delta["tool"].(map[string]any); ok && tm["phase"] == "end" {
			if r, ok := tm["result"].(string); ok {
				toolResults.WriteString(r)
			}
		}
	}
	c.mu.Unlock()
	for _, want := range []string{"contenu A", "contenu B"} {
		if !strings.Contains(toolResults.String(), want) {
			t.Fatalf("résultat manquant %q dans %q", want, toolResults.String())
		}
	}
	// Ordre des événements : 3 start puis 3 end, dans l'ordre d'origine.
	c.mu.Lock()
	var phases []string
	for _, ev := range c.Log {
		if tm, ok := ev.Delta["tool"].(map[string]any); ok {
			phases = append(phases, tm["name"].(string)+":"+tm["phase"].(string))
		}
	}
	c.mu.Unlock()
	want := []string{"Read:start", "Read:start", "Cat:start", "Read:end", "Read:end", "Cat:end"}
	if strings.Join(phases, ",") != strings.Join(want, ",") {
		t.Fatalf("ordre événements = %v, want %v", phases, want)
	}
	// Les messages d'outils suivent l'ordre des appels dans la requête 2.
	reqs := sp.requests()
	if len(reqs) != 2 {
		t.Fatalf("requêtes = %d, want 2", len(reqs))
	}
	var ids []string
	for _, m := range reqs[1].Messages {
		if m.Role == "tool" {
			ids = append(ids, m.ToolCallID)
		}
	}
	if strings.Join(ids, ",") != "c1,c2,c3" {
		t.Fatalf("ordre messages outils = %v, want c1,c2,c3", ids)
	}
}

func TestAgentParallelReadsFallsBackOnWrite(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{
			toolCall("c1", "Read", `{"file_path":"a.txt"}`),
			toolCall("c2", "Bash", `{"command":"echo hi"}`),
		}},
		{content: "Termine"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "mixte"})
	if got := logText(c); !strings.Contains(got, "Termine") {
		t.Fatalf("contenu final = %q", got)
	}
	// Le bloc mixte lecture/écriture passe en séquentiel : les deux
	// outils ont quand même été exécutés (start+end chacun).
	if n := countEvents(c, "tool"); n != 4 {
		t.Fatalf("événements tool = %d, want 4 (2 start + 2 end)", n)
	}
}

func TestAgentParallelReadsFallsBackOnTodoWrite(t *testing.T) {
	// TodoWrite n'est PAS parallélisable (état de la liste de tâches) :
	// un bloc Read + TodoWrite retombe sur la voie séquentielle.
	ws := t.TempDir()
	sandboxDir := filepath.Join(ws, "sam")
	if err := os.MkdirAll(sandboxDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sandboxDir, "a.txt"), []byte("contenu A"), 0o644); err != nil {
		t.Fatal(err)
	}
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{
			toolCall("c1", "Read", `{"file_path":"a.txt"}`),
			toolCall("c2", "TodoWrite", `{"todos":[{"content":"lire a.txt","status":"completed"}]}`),
		}},
		{content: "Termine"},
	}}
	st, err := store.Open(filepath.Join(t.TempDir(), "agent.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	reg := provider.NewRegistry()
	reg.Set(sp)
	e := NewEngine(reg, codeFamily(alias.Member{Provider: "fake", Model: "m"}), st, nil, ws)
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "lis et note"})
	if got := logText(c); !strings.Contains(got, "Termine") {
		t.Fatalf("contenu final = %q", got)
	}
	if n := countEvents(c, "tool"); n != 4 {
		t.Fatalf("événements tool = %d, want 4 (2 start + 2 end)", n)
	}
	// Les deux outils ont été exécutés malgré le repli séquentiel.
	if !hasToolEvent(c, "end") {
		t.Fatal("aucun événement tool end")
	}
}

func TestAgentParallelReadsErrorIsolated(t *testing.T) {
	// Une erreur sur un appel du bloc ne casse pas les autres :
	// chaque appel produit son événement "end", dans l'ordre.
	ws := t.TempDir()
	sandboxDir := filepath.Join(ws, "sam")
	if err := os.MkdirAll(sandboxDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sandboxDir, "ok.txt"), []byte("contenu OK"), 0o644); err != nil {
		t.Fatal(err)
	}
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{
			toolCall("c1", "Read", `{"file_path":"manquant.txt"}`),
			toolCall("c2", "Read", `{"file_path":"ok.txt"}`),
		}},
		{content: "Termine"},
	}}
	st, err := store.Open(filepath.Join(t.TempDir(), "agent.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	reg := provider.NewRegistry()
	reg.Set(sp)
	e := NewEngine(reg, codeFamily(alias.Member{Provider: "fake", Model: "m"}), st, nil, ws)
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "lis"})
	// Ordre des "end" : c1 (erreur) puis c2 (ok), comme l'ordre d'origine.
	c.mu.Lock()
	var ends []string
	for _, ev := range c.Log {
		if tm, ok := ev.Delta["tool"].(map[string]any); ok && tm["phase"] == "end" {
			ends = append(ends, tm["result"].(string))
		}
	}
	c.mu.Unlock()
	if len(ends) != 2 {
		t.Fatalf("événements end = %d, want 2", len(ends))
	}
	if !strings.Contains(ends[0], "[erreur]") {
		t.Fatalf("end[0] devrait être l'erreur, got %q", ends[0])
	}
	if !strings.Contains(ends[1], "contenu OK") {
		t.Fatalf("end[1] devrait contenir le fichier, got %q", ends[1])
	}
}

func TestAgentEchoesReasoningContentAfterToolCall(t *testing.T) {
	// Phase 0 : en mode thinking, le raisonnement emis avec un appel
	// d'outils doit etre renvoye dans le message assistant
	// (reasoning_content) au tour suivant, sinon DeepSeek refuse la
	// requete en HTTP 400.
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{reasoning: "je dois lire le fichier", toolCalls: []provider.ToolCall{toolCall("c1", "TodoWrite", `{"todos":[]}`)}},
		{content: "Termine"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "fais", Think: true})

	reqs := sp.requests()
	if len(reqs) != 2 {
		t.Fatalf("requetes = %d, want 2", len(reqs))
	}
	found := false
	for _, m := range reqs[1].Messages {
		if m.Role == "assistant" && len(m.ToolCalls) > 0 {
			found = true
			if m.ReasoningContent != "je dois lire le fichier" {
				t.Fatalf("reasoning_content = %q, want %q", m.ReasoningContent, "je dois lire le fichier")
			}
		}
	}
	if !found {
		t.Fatal("message assistant avec tool_calls attendu dans la requete 2")
	}
}

func TestAgentNoReasoningContentWhenNoReasoning(t *testing.T) {
	// Sans thinking, aucun reasoning_content ne doit etre ajoute.
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{toolCall("c1", "TodoWrite", `{"todos":[]}`)}},
		{content: "Termine"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "fais"})

	reqs := sp.requests()
	if len(reqs) != 2 {
		t.Fatalf("requetes = %d, want 2", len(reqs))
	}
	for _, m := range reqs[1].Messages {
		if m.Role == "assistant" && m.ReasoningContent != "" {
			t.Fatalf("reasoning_content inattendu = %q", m.ReasoningContent)
		}
	}
}
