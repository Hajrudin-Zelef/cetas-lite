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

func TestAgentThinkingOnEnablesReasoning(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "fais", Think: true})
	reqs := sp.requests()
	if len(reqs) == 0 || !reqs[0].EnableReasoning {
		t.Fatal("l'agent avec thinking actif doit activer le raisonnement")
	}
	if reqs[0].ReasoningEffort == "" {
		t.Fatal("effort attendu en mode agent")
	}
}

// TestAgentThinkingOffDisablesReasoning : le bouton Thinking eteint doit
// etre respecte de bout en bout — pas de raisonnement cote provider, pas
// de contenu de raisonnement diffuse, directive systeme "reponse directe".
func TestAgentThinkingOffDisablesReasoning(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok", reasoning: "reflexion privee"}}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "fais vite", Think: false})
	reqs := sp.requests()
	if len(reqs) == 0 {
		t.Fatal("aucune requete emise")
	}
	if reqs[0].EnableReasoning {
		t.Fatal("l'agent avec thinking eteint ne doit pas activer le raisonnement provider")
	}
	if logHasReasoning(c) {
		t.Fatal("aucun contenu de raisonnement ne doit etre diffuse quand le thinking est eteint")
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
	if got := resolveEffort(false, "court", "default"); got != "" {
		t.Fatalf("effort sans thinking = %q", got)
	}
	if got := resolveEffort(true, "court", "default"); got != "low" {
		t.Fatalf("auto court = %q", got)
	}
	long := make([]rune, 500)
	for i := range long {
		long[i] = 'x'
	}
	if got := resolveEffort(true, string(long), "default"); got != "medium" {
		t.Fatalf("auto long = %q", got)
	}
	if got := resolveEffort(true, "court", "high"); got != "high" {
		t.Fatalf("effort explicite = %q", got)
	}
	// Durcissement : un effort explicite ne survit pas à think=false.
	if got := resolveEffort(false, "court", "high"); got != "" {
		t.Fatalf("effort explicite + think=false = %q, attendu vide", got)
	}
}

func TestResolveAgentEffort(t *testing.T) {
	// Agent : le "Défaut" du composer vaut "low" (la boucle itère déjà,
	// un raisonnement long par itération ralentit l'expérience).
	if got := resolveAgentEffort(true, "default"); got != "low" {
		t.Fatalf("agent defaut = %q, attendu low", got)
	}
	for _, tc := range [][2]string{{"low", "low"}, {"medium", "medium"}, {"high", "high"}} {
		if got := resolveAgentEffort(true, tc[0]); got != tc[1] {
			t.Fatalf("agent effort %q = %q", tc[0], got)
		}
	}
	if got := resolveAgentEffort(false, "high"); got != "" {
		t.Fatalf("agent think=false = %q, attendu vide", got)
	}
}
