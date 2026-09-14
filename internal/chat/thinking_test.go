package chat

import (
	"testing"

	"cetas-lite/internal/alias"
)

func logHasReasoning(c *Conversation) bool {
	for _, ev := range c.marshal().Log {
		if _, ok := ev.Delta["reasoning_content"]; ok {
			return true
		}
	}
	return false
}

func TestAgentForcesReasoning(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "fais"})
	reqs := sp.requests()
	if len(reqs) == 0 || !reqs[0].EnableReasoning {
		t.Fatal("l'agent doit forcer le raisonnement")
	}
	if reqs[0].ReasoningEffort == "" {
		t.Fatal("effort attendu en mode agent")
	}
}

func TestChatThinkingOffNoReasoning(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok", reasoning: "reflexion privee"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "bonjour"})
	if sp.requests()[0].EnableReasoning {
		t.Fatal("le chat par defaut ne doit pas activer le raisonnement")
	}
	if logHasReasoning(c) {
		t.Fatal("aucun raisonnement ne doit etre diffuse par defaut")
	}
}

func TestChatThinkingOnEmitsReasoning(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok", reasoning: "reflexion"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "bonjour", Think: true})
	if !sp.requests()[0].EnableReasoning {
		t.Fatal("think=true doit activer le raisonnement")
	}
	if !logHasReasoning(c) {
		t.Fatal("le raisonnement doit etre diffuse quand active")
	}
}

func TestResolveEffort(t *testing.T) {
	if got := resolveEffort(false, false, "court", "default"); got != "" {
		t.Fatalf("effort sans thinking = %q", got)
	}
	if got := resolveEffort(false, true, "court", "default"); got != "low" {
		t.Fatalf("auto court = %q", got)
	}
	long := make([]rune, 500)
	for i := range long {
		long[i] = 'x'
	}
	if got := resolveEffort(false, true, string(long), "default"); got != "medium" {
		t.Fatalf("auto long = %q", got)
	}
	if got := resolveEffort(true, true, "court", "high"); got != "high" {
		t.Fatalf("effort explicite = %q", got)
	}
}
