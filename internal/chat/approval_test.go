package chat

import (
	"context"
	"strings"
	"testing"
	"time"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/provider"
)

func TestNeedsApproval(t *testing.T) {
	for _, name := range []string{"Write", "Edit", "Bash", "RunScript", "mcp_fetch", "custom_deploy"} {
		if !needsApproval(name) {
			t.Errorf("%s devrait exiger une approbation", name)
		}
	}
	for _, name := range []string{"Ls", "Read", "Grep", "Glob", "TodoWrite"} {
		if needsApproval(name) {
			t.Errorf("%s ne devrait pas exiger d'approbation", name)
		}
	}
}

func TestReadOnlyTools(t *testing.T) {
	tools := []provider.Tool{
		{Type: "function", Function: provider.ToolFunction{Name: "Ls"}},
		{Type: "function", Function: provider.ToolFunction{Name: "Write"}},
		{Type: "function", Function: provider.ToolFunction{Name: "Bash"}},
		{Type: "function", Function: provider.ToolFunction{Name: "TodoWrite"}},
	}
	got := readOnlyTools(tools)
	if len(got) != 2 || got[0].Function.Name != "Ls" || got[1].Function.Name != "TodoWrite" {
		t.Fatalf("filtrage incorrect: %v", got)
	}
}

func TestFormatTodos(t *testing.T) {
	args := map[string]any{"todos": []any{
		map[string]any{"content": "explorer", "status": "completed"},
		map[string]any{"content": "coder", "status": "in_progress"},
		map[string]any{"content": "tester", "status": "pending"},
	}}
	out := formatTodos(args)
	if !strings.Contains(out, "[x] explorer") || !strings.Contains(out, "[~] coder") || !strings.Contains(out, "[ ] tester") {
		t.Fatalf("format inattendu: %q", out)
	}
	if formatTodos(map[string]any{}) != "" {
		t.Fatal("args vides devraient donner une chaine vide")
	}
}

func TestApprovalRequestResolve(t *testing.T) {
	c := NewConversation("1", nil, nil)
	type res struct {
		d   approvalDecision
		err error
	}
	out := make(chan res, 1)
	go func() {
		d, err := c.RequestApproval(context.Background(), 0, ApprovalRequest{Kind: "tool", Tool: "Write", Args: map[string]any{"file_path": "a.go"}})
		out <- res{d, err}
	}()

	var id string
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		c.mu.Lock()
		for k := range c.approvals {
			id = k
		}
		c.mu.Unlock()
		if id != "" {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	if id == "" {
		t.Fatal("aucune demande d'approbation")
	}
	if c.ResolveApproval("inconnu", true, false) {
		t.Fatal("resolve d'un id inconnu devrait echouer")
	}
	if !c.ResolveApproval(id, true, true) {
		t.Fatal("resolve a echoue")
	}
	select {
	case r := <-out:
		if r.err != nil {
			t.Fatalf("err: %v", r.err)
		}
		if !r.d.approved || !r.d.always {
			t.Fatalf("decision = %+v", r.d)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timeout")
	}

	// Le log contient la demande et sa resolution.
	c.mu.Lock()
	defer c.mu.Unlock()
	var sawReq, sawRes bool
	for _, ev := range c.Log {
		a, ok := ev.Delta["approval"].(map[string]any)
		if !ok {
			continue
		}
		if a["phase"] == "request" && a["tool"] == "Write" {
			sawReq = true
		}
		if a["phase"] == "resolved" && a["approved"] == true {
			sawRes = true
		}
	}
	if !sawReq || !sawRes {
		t.Fatal("le log devrait contenir la demande et sa resolution")
	}
}

func TestApprovalResetFailsPending(t *testing.T) {
	c := NewConversation("1", nil, nil)
	out := make(chan error, 1)
	go func() {
		_, err := c.RequestApproval(context.Background(), 0, ApprovalRequest{Kind: "tool", Tool: "Bash"})
		out <- err
	}()
	waitFor(t, func() bool {
		c.mu.Lock()
		defer c.mu.Unlock()
		return len(c.approvals) > 0
	}, "demande en attente")
	c.Reset()
	select {
	case err := <-out:
		if err != nil {
			t.Fatalf("err: %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timeout : le reset devrait debloquer la demande")
	}
}

// autoResolveApprovals valide automatiquement toute demande en attente.
func autoResolveApprovals(t *testing.T, c *Conversation, approved bool, stop <-chan struct{}) {
	t.Helper()
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
			}
			c.mu.Lock()
			var ids []string
			for id := range c.approvals {
				ids = append(ids, id)
			}
			c.mu.Unlock()
			for _, id := range ids {
				c.ResolveApproval(id, approved, false)
			}
			time.Sleep(5 * time.Millisecond)
		}
	}()
}

func toolResults(c *Conversation) []string {
	c.mu.Lock()
	defer c.mu.Unlock()
	var out []string
	for _, ev := range c.Log {
		tm, ok := ev.Delta["tool"].(map[string]any)
		if !ok || tm["phase"] != "end" {
			continue
		}
		if r, ok := tm["result"].(string); ok {
			out = append(out, r)
		}
	}
	return out
}

func TestPlanModeExplorationThenApproval(t *testing.T) {
	todos := `{"todos": [{"content": "explorer le code", "status": "completed"}, {"content": "corriger le bug", "status": "pending"}]}`
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{toolCall("1", "TodoWrite", todos)}},
		{toolCalls: []provider.ToolCall{toolCall("2", "Write", `{"file_path": "x.go", "content": "y"}`)}},
		{content: "Voici mon plan : corriger le bug."},
		{content: "Execution terminee."},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	c := e.Conversation("sam")
	stop := make(chan struct{})
	defer close(stop)
	autoResolveApprovals(t, c, true, stop)

	if err := c.StartTurn(TurnInput{Family: "code", Mode: "standard", Text: "corrige le bug", Plan: true, User: "sam"}); err != nil {
		t.Fatalf("StartTurn: %v", err)
	}
	waitFor(t, func() bool { return !c.IsGenerating() }, "tour non termine")

	results := toolResults(c)
	var sawRefusal bool
	for _, r := range results {
		if strings.Contains(r, "phase d'exploration") {
			sawRefusal = true
		}
		if strings.Contains(r, "[ok]") && strings.Contains(strings.Join(results, ""), "x.go") {
			t.Fatal("Write aurait du etre refuse en phase d'exploration")
		}
	}
	if !sawRefusal {
		t.Fatalf("l'ecriture en phase d'exploration aurait du etre refusee : %v", results)
	}
	if got := logText(c); !strings.Contains(got, "Execution terminee") {
		t.Fatalf("la phase d'execution n'a pas eu lieu : %q", got)
	}
	// La phase 1 ne devait proposer que des outils de lecture.
	reqs := sp.requests()
	if len(reqs) == 0 || len(reqs[0].Tools) == 0 {
		t.Fatal("la phase d'exploration devrait proposer des outils")
	}
	for _, tl := range reqs[0].Tools {
		if needsApproval(tl.Function.Name) {
			t.Fatalf("outil sensible propose en phase d'exploration : %s", tl.Function.Name)
		}
	}
}

func TestPlanModeDenied(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{content: "Voici mon plan."},
		{content: "Ne devrait jamais arriver."},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	c := e.Conversation("sam")
	stop := make(chan struct{})
	defer close(stop)
	autoResolveApprovals(t, c, false, stop)

	if err := c.StartTurn(TurnInput{Family: "code", Mode: "standard", Text: "fais quelque chose", Plan: true, User: "sam"}); err != nil {
		t.Fatalf("StartTurn: %v", err)
	}
	waitFor(t, func() bool { return !c.IsGenerating() }, "tour non termine")

	if got := logText(c); strings.Contains(got, "Ne devrait jamais arriver") {
		t.Fatal("l'execution aurait du etre bloquee apres refus du plan")
	}
	reqs := sp.requests()
	if len(reqs) != 1 {
		t.Fatalf("requetes = %d, want 1 (pas de phase d'execution)", len(reqs))
	}
}

func TestToolApprovalDenied(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{toolCall("1", "Write", `{"file_path": "x.go", "content": "y"}`)}},
		{content: "J'ai compris, je propose autre chose."},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	c := e.Conversation("sam")
	stop := make(chan struct{})
	defer close(stop)
	autoResolveApprovals(t, c, false, stop)

	if err := c.StartTurn(TurnInput{Family: "code", Mode: "standard", Text: "ecris un fichier", Approve: true, User: "sam"}); err != nil {
		t.Fatalf("StartTurn: %v", err)
	}
	waitFor(t, func() bool { return !c.IsGenerating() }, "tour non termine")

	var sawRefusal bool
	for _, r := range toolResults(c) {
		if strings.Contains(r, "[refuse]") {
			sawRefusal = true
		}
	}
	if !sawRefusal {
		t.Fatal("le refus d'approbation aurait du etre transmis au modele")
	}
	if got := logText(c); !strings.Contains(got, "autre chose") {
		t.Fatalf("le modele aurait du continuer apres le refus : %q", got)
	}
}
