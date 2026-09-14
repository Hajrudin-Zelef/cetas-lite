package chat

import (
	"context"
	"strings"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/customtools"
	"cetas-lite/internal/provider"
)

type stubCustom struct {
	defs  []customtools.Def
	out   string
	err   error
	calls []string
}

func (s *stubCustom) Defs() []customtools.Def { return s.defs }
func (s *stubCustom) Call(ctx context.Context, name, argsJSON string) (string, error) {
	s.calls = append(s.calls, name)
	return s.out, s.err
}

func demoCustom() *stubCustom {
	return &stubCustom{defs: []customtools.Def{{
		Name:        "custom_demo",
		Description: "Demo",
		Schema:      map[string]any{"type": "object", "properties": map[string]any{"q": map[string]any{"type": "string"}}},
	}}}
}

func TestAgentExposesCustomTools(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.SetCustom(demoCustom())
	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "salut"})
	if !hasToolName(sp.requests()[0].Tools, "custom_demo") {
		t.Fatal("l'outil personnalise doit etre expose en mode agent")
	}
}

func TestAgentCustomDispatch(t *testing.T) {
	stub := demoCustom()
	stub.out = "resultat custom"
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{toolCall("c1", "custom_demo", `{"q":"x"}`)}},
		{content: "fini"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.SetCustom(stub)
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "demo"})

	if got := logText(c); !strings.Contains(got, "fini") {
		t.Fatalf("contenu = %q", got)
	}
	if len(stub.calls) != 1 || stub.calls[0] != "custom_demo" {
		t.Fatalf("appels = %+v", stub.calls)
	}
	var toolResult string
	for _, m := range sp.requests()[1].Messages {
		if m.Role == "tool" {
			toolResult, _ = m.Content.(string)
		}
	}
	if toolResult != "resultat custom" {
		t.Fatalf("resultat = %q", toolResult)
	}
}

func TestCustomToolSchemas(t *testing.T) {
	defs := []customtools.Def{
		{Name: "custom_demo", Description: "Demo", Schema: map[string]any{"type": "object", "properties": map[string]any{}}},
		{Name: "custom_bare"},
	}
	schemas := CustomToolSchemas(defs)
	if len(schemas) != 2 {
		t.Fatalf("schemas = %d", len(schemas))
	}
	if schemas[0].Function.Name != "custom_demo" || !strings.Contains(schemas[0].Function.Description, "[Outil]") {
		t.Fatalf("schema 0 = %+v", schemas[0].Function)
	}
	if _, ok := schemas[1].Function.Parameters.(map[string]any); !ok {
		t.Fatalf("schema de repli attendu: %T", schemas[1].Function.Parameters)
	}
}

func TestCustomDedupExcluded(t *testing.T) {
	if dedupableTool("custom_demo") {
		t.Fatal("les outils personnalises ne doivent pas etre dedupliques")
	}
}
