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
	runTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "bonjour"})
	if !hasSystemContaining(sp.requests(), "senior teacher") {
		t.Fatal("le prompt systeme de chat doit etre injecte")
	}
}

func TestMarexInjectedForAgent(t *testing.T) {
	tmp := t.TempDir()
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.workspace = filepath.Join(tmp, "workspace")
	if err := os.MkdirAll(filepath.Join(e.workspace, "sam"), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(e.workspace, "sam", "MAREX.md"), []byte("PROJECT RULE: use tabs"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tmp, "MAREX.md"), []byte("GLOBAL RULE: be terse"), 0o600); err != nil {
		t.Fatal(err)
	}
	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "fais"})

	reqs := sp.requests()
	if !hasSystemContaining(reqs, "[MAREX.md]") {
		t.Fatal("MAREX.md doit etre injecte en mode agent")
	}
	if !hasSystemContaining(reqs, "PROJECT RULE") || !hasSystemContaining(reqs, "GLOBAL RULE") {
		t.Fatal("contenu global + projet attendu")
	}
}

func TestMarexAbsentNoInjection(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.workspace = filepath.Join(t.TempDir(), "workspace")
	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "fais"})
	if hasSystemContaining(sp.requests(), "[MAREX.md]") {
		t.Fatal("aucune injection sans MAREX.md")
	}
}

func TestReadCapped(t *testing.T) {
	path := filepath.Join(t.TempDir(), "MAREX.md")
	if err := os.WriteFile(path, []byte(strings.Repeat("x", marexMaxBytes*2)), 0o600); err != nil {
		t.Fatal(err)
	}
	if got := len(readCapped(path)); got != marexMaxBytes {
		t.Fatalf("taille = %d, want %d", got, marexMaxBytes)
	}
	if readCapped(filepath.Join(t.TempDir(), "absent.md")) != "" {
		t.Fatal("fichier absent doit donner une chaine vide")
	}
}
