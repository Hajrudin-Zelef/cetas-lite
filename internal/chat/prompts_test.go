package chat

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/provider"
)

func plainFamily(pool ...alias.Member) []alias.Family {
	return []alias.Family{{
		ID: "plain", Label: "Plain",
		Modes: []alias.Mode{{ID: "standard", Pool: pool}},
	}}
}

func hasSystemContaining(reqs []provider.Request, marker string) bool {
	for _, r := range reqs {
		for _, m := range r.Messages {
			if m.Role != "system" {
				continue
			}
			if s, ok := m.Content.(string); ok && strings.Contains(s, marker) {
				return true
			}
		}
	}
	return false
}

func TestChatSystemPromptInjected(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	runTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "hello"})
	if !hasSystemContaining(sp.requests(), "senior teacher") {
		t.Fatal("le prompt systeme de chat doit etre injecte")
	}
}

func TestChatSystemPromptTokenBudget(t *testing.T) {
	p := chatSystemPrompt()
	// ~4 caracteres par token : 700 caracteres ~= 175 tokens,
	// tres largement sous la limite d'1k tokens en entree pour un "salut".
	if len(p) > 700 {
		t.Fatalf("prompt systeme trop long : %d caracteres (budget 700)", len(p))
	}
	for _, must := range []string{"senior teacher", "premium pedagogy", "do not produce code", "never invent", "user's language"} {
		if !strings.Contains(p, must) {
			t.Fatalf("le prompt systeme devrait contenir %q", must)
		}
	}
}

func TestSearchDirective(t *testing.T) {
	off := searchDirective(false, false)
	if !strings.Contains(off, "Web search is disabled") {
		t.Fatalf("globe desactive : directive inattendue %q", off)
	}
	nat := searchDirective(true, true)
	if !strings.Contains(nat, "provider-native") || !strings.Contains(nat, "knowledge cutoff") {
		t.Fatalf("natif : directive inattendue %q", nat)
	}
	tools := searchDirective(true, false)
	if !strings.Contains(tools, "web_search and web_fetch") || !strings.Contains(tools, "MUST search") {
		t.Fatalf("outils : directive inattendue %q", tools)
	}
}

func TestNativeWebExtraFor(t *testing.T) {
	ds := nativeWebExtraFor("deepseek")
	if ds["enable_search"] != true {
		t.Fatalf("deepseek : enable_search attendu, got %v", ds)
	}
	or := nativeWebExtraFor("openrouter")
	plugs, ok := or["plugins"].([]any)
	if !ok || len(plugs) != 1 {
		t.Fatalf("openrouter : plugin web attendu, got %v", or)
	}
	plug, ok := plugs[0].(map[string]any)
	if !ok || plug["id"] != "web" {
		t.Fatalf("openrouter : plugin web attendu, got %v", or)
	}
}

func TestMarexInjectedForAgent(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "MAREX.md")
	if err := os.WriteFile(path, []byte("PROJECT RULE: use tabs"), 0o600); err != nil {
		t.Fatal(err)
	}
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.SetMarexPath(path)
	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "fais"})

	reqs := sp.requests()
	if !hasSystemContaining(reqs, "[MAREX.md]") {
		t.Fatal("MAREX.md doit etre injecte en mode agent")
	}
	if !hasSystemContaining(reqs, "PROJECT RULE") {
		t.Fatal("contenu MAREX.md attendu")
	}
}

func TestMarexNoticeOncePerSession(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "MAREX.md")
	if err := os.WriteFile(path, []byte("  RULE: be terse  "), 0o600); err != nil {
		t.Fatal(err)
	}
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.SetMarexPath(path)

	msg, ok, notice := e.marexForSession("conv1")
	if !ok || !notice {
		t.Fatal("premier tour : injection + notice attendues")
	}
	if s, _ := msg.Content.(string); !strings.Contains(s, "RULE: be terse") {
		t.Fatalf("contenu rogne attendu, got %q", s)
	}
	if _, ok, notice := e.marexForSession("conv1"); !ok || notice {
		t.Fatal("second tour : injection sans notice")
	}
	// Autre session : relecture + nouvelle notice.
	if _, ok, notice := e.marexForSession("conv2"); !ok || !notice {
		t.Fatal("nouvelle session : injection + notice attendues")
	}
}

func TestMarexEmptyOrAbsentIgnored(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))

	// Fichier absent.
	e.SetMarexPath(filepath.Join(t.TempDir(), "absent.md"))
	if _, ok, _ := e.marexForSession("c1"); ok {
		t.Fatal("fichier absent : ignore silencieusement")
	}
	// Fichier vide / blanc.
	path := filepath.Join(t.TempDir(), "MAREX.md")
	if err := os.WriteFile(path, []byte("   \n  "), 0o600); err != nil {
		t.Fatal(err)
	}
	e.SetMarexPath(path)
	if _, ok, _ := e.marexForSession("c2"); ok {
		t.Fatal("fichier vide : ignore silencieusement")
	}
}

func TestMarexTruncated(t *testing.T) {
	path := filepath.Join(t.TempDir(), "MAREX.md")
	if err := os.WriteFile(path, []byte(strings.Repeat("x", marexMaxBytes*2)), 0o600); err != nil {
		t.Fatal(err)
	}
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.SetMarexPath(path)
	msg, ok, _ := e.marexForSession("c1")
	if !ok {
		t.Fatal("injection attendue")
	}
	s, _ := msg.Content.(string)
	if got := len(s) - len("[MAREX.md]\n\n"); got != marexMaxBytes {
		t.Fatalf("taille = %d, want %d", got, marexMaxBytes)
	}
}

func TestThinkDirective(t *testing.T) {
	if d := thinkDirective(false, false, ""); !strings.Contains(d, "Answer directly") {
		t.Fatalf("chat sans thinking : reponse directe attendue, obtenu %q", d)
	}
	if d := thinkDirective(false, true, ""); !strings.Contains(d, "think before answering") {
		t.Fatalf("chat avec thinking : raisonnement attendu, obtenu %q", d)
	}
	if d := thinkDirective(true, true, ""); !strings.Contains(d, "mandatory") {
		t.Fatalf("agent avec thinking : raisonnement obligatoire attendu, obtenu %q", d)
	}
	if d := thinkDirective(true, true, "high"); !strings.Contains(d, "think carefully") {
		t.Fatalf("agent avec thinking high : raisonnement approfondi attendu, obtenu %q", d)
	}
	if d := thinkDirective(true, true, "low"); !strings.Contains(d, "brief") {
		t.Fatalf("agent avec thinking low : raisonnement bref attendu, obtenu %q", d)
	}
	if d := thinkDirective(true, true, "low"); !strings.Contains(d, "trivial") {
		t.Fatalf("agent avec thinking low : consigne messages triviaux attendue, obtenu %q", d)
	}
	if d := thinkDirective(true, false, ""); !strings.Contains(d, "Answer directly") {
		t.Fatalf("agent sans thinking : reponse directe attendue, obtenu %q", d)
	}
}

// La directive thinking est injectee dans les requetes chat et agent.
func TestThinkDirectiveInjected(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "x", Effort: "high", Think: true})
	reqs := sp.requests()
	found := false
	for _, m := range reqs[0].Messages {
		if m.Role == "system" {
			if s, _ := m.Content.(string); strings.Contains(s, "Reasoning is mandatory") && strings.Contains(s, "think carefully") {
				found = true
			}
		}
	}
	if !found {
		t.Fatal("directive thinking agent (effort high) attendue dans la requete")
	}

	// Agent avec thinking eteint : la directive ordonne une reponse directe.
	sp0 := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e0 := newAgentEngine(t, sp0, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	runAgentTurn(t, e0, "sam", TurnInput{Family: "code", Mode: "standard", Text: "x", Think: false})
	found = false
	for _, m := range sp0.requests()[0].Messages {
		if m.Role == "system" {
			if s, _ := m.Content.(string); strings.Contains(s, "Answer directly") {
				found = true
			}
		}
	}
	if !found {
		t.Fatal("directive 'Answer directly' attendue quand le thinking agent est eteint")
	}

	sp2 := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e2 := newChatEngine(t, []*scriptedProvider{sp2}, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	runTurn(t, e2, "sam", TurnInput{User: "sam", Family: "plain", Mode: "standard", Text: "x", Think: false})
	reqs2 := sp2.requests()
	found = false
	for _, m := range reqs2[0].Messages {
		if m.Role == "system" {
			if s, _ := m.Content.(string); strings.Contains(s, "Answer directly") {
				found = true
			}
		}
	}
	if !found {
		t.Fatal("directive thinking desactivee attendue dans la requete chat")
	}
}
