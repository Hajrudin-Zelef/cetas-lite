package chat

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/mcp"
	"cetas-lite/internal/provider"
)

type stubMCP struct {
	defs  []mcp.Tool
	out   string
	err   error
	calls []string
}

func (s *stubMCP) Tools(ctx context.Context) []mcp.Tool { return s.defs }
func (s *stubMCP) Call(ctx context.Context, name string, args map[string]any) (string, error) {
	s.calls = append(s.calls, name)
	return s.out, s.err
}
func (s *stubMCP) Servers() []mcp.ServerStatus { return nil }

func echoTool() stubMCP {
	return stubMCP{defs: []mcp.Tool{{
		Exposed: "mcp_demo_echo", Server: "demo", Name: "echo", Description: "Echo",
		InputSchema: json.RawMessage(`{"type":"object","properties":{"text":{"type":"string"}},"required":["text"]}`),
	}}}
}

func TestAgentExposesMCPOnlyWhenEnabled(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	echo := echoTool()
	e.SetMCP(&echo)

	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "sans mcp"})
	if hasToolName(sp.requests()[0].Tools, "mcp_demo_echo") {
		t.Fatal("les outils MCP ne doivent pas etre exposes sans le toggle")
	}

	sp2 := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e2 := newAgentEngine(t, sp2, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	echo2 := echoTool()
	e2.SetMCP(&echo2)
	runAgentTurn(t, e2, "sam", TurnInput{Family: "code", Mode: "standard", Text: "avec mcp", MCP: true})
	if !hasToolName(sp2.requests()[0].Tools, "mcp_demo_echo") {
		t.Fatal("les outils MCP doivent etre exposes avec le toggle")
	}
}

func TestAgentMCPDispatch(t *testing.T) {
	stub := echoTool()
	stub.out = "resultat mcp"
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{toolCall("c1", "mcp_demo_echo", `{"text":"hi"}`)}},
		{content: "fini"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.SetMCP(&stub)
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "echo", MCP: true})

	if got := logText(c); !strings.Contains(got, "fini") {
		t.Fatalf("contenu = %q", got)
	}
	if len(stub.calls) != 1 || stub.calls[0] != "mcp_demo_echo" {
		t.Fatalf("appels = %+v", stub.calls)
	}
	var toolResult string
	for _, m := range sp.requests()[1].Messages {
		if m.Role == "tool" {
			toolResult, _ = m.Content.(string)
		}
	}
	if toolResult != "resultat mcp" {
		t.Fatalf("resultat outil = %q", toolResult)
	}
}

func TestMCPToolSchemas(t *testing.T) {
	defs := []mcp.Tool{{
		Exposed: "mcp_demo_echo", Server: "demo", Name: "echo", Description: "Echo",
		InputSchema: json.RawMessage(`{"type":"object","properties":{"text":{"type":"string"}}}`),
	}, {
		Exposed: "mcp_x_y", Server: "x", Name: "y",
	}}
	schemas := MCPToolSchemas(defs)
	if len(schemas) != 2 {
		t.Fatalf("schemas=%d", len(schemas))
	}
	if schemas[0].Function.Name != "mcp_demo_echo" || !strings.Contains(schemas[0].Function.Description, "[MCP:demo]") {
		t.Fatalf("schema 0 = %+v", schemas[0].Function)
	}
	if _, ok := schemas[0].Function.Parameters.(map[string]any); !ok {
		t.Fatalf("parametres attendus en objet: %T", schemas[0].Function.Parameters)
	}
	if schemas[1].Function.Description == "" {
		t.Fatal("description de repli attendue")
	}
}

func TestMCPDedupExcluded(t *testing.T) {
	if dedupableTool("mcp_demo_echo") {
		t.Fatal("les outils mcp ne doivent pas etre dedupliques")
	}
	if !dedupableTool("Read") {
		t.Fatal("Read doit rester deduplique")
	}
}
