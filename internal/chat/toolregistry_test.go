package chat

import (
	"context"
	"strings"
	"testing"

	"cetas-lite/internal/provider"
)

type fakeFamily struct {
	prefix   string
	schemasN int
}

func (f fakeFamily) schemas(context.Context) []provider.Tool {
	out := make([]provider.Tool, f.schemasN)
	for i := range out {
		out[i] = provider.Tool{Function: provider.ToolFunction{Name: f.prefix + "tool"}}
	}
	return out
}

func (f fakeFamily) handles(name string) bool { return strings.HasPrefix(name, f.prefix) }

func (f fakeFamily) execute(_ context.Context, _ toolEnv, name, _ string) (ToolResult, *provider.Message) {
	return ToolResult{Text: "famille:" + name}, nil
}

type catchAllFamily struct{ tag string }

func (catchAllFamily) schemas(context.Context) []provider.Tool { return nil }
func (catchAllFamily) handles(string) bool                     { return true }
func (f catchAllFamily) execute(_ context.Context, _ toolEnv, _ string, _ string) (ToolResult, *provider.Message) {
	return ToolResult{Text: f.tag}, nil
}

func TestToolRegistryPrecedenceAndFallback(t *testing.T) {
	reg := toolRegistry{
		families: []toolFamily{fakeFamily{prefix: "mem_", schemasN: 2}, fakeFamily{prefix: "web_", schemasN: 1}},
		fallback: catchAllFamily{tag: "defaut"},
	}

	if got := len(reg.schemas(context.Background())); got != 3 {
		t.Fatalf("schemas = %d, veut 3", got)
	}

	out, _ := reg.execute(context.Background(), toolEnv{}, "mem_read", "{}")
	if out.Text != "famille:mem_read" {
		t.Fatalf("dispatch mem = %q", out.Text)
	}
	out, _ = reg.execute(context.Background(), toolEnv{}, "web_search", "{}")
	if out.Text != "famille:web_search" {
		t.Fatalf("dispatch web = %q", out.Text)
	}
	out, _ = reg.execute(context.Background(), toolEnv{}, "Read", "{}")
	if out.Text != "defaut" {
		t.Fatalf("repli = %q", out.Text)
	}
}

func TestToolRegistryBuiltinSchemas(t *testing.T) {
	reg := toolRegistry{
		families: []toolFamily{visionFamily{}},
		fallback: builtinFamily{allowScript: false},
	}
	names := map[string]bool{}
	for _, tl := range reg.schemas(context.Background()) {
		names[tl.Function.Name] = true
	}
	if names["RunScript"] {
		t.Fatal("RunScript ne doit pas etre expose quand desactive")
	}
	if !names["Read"] || !names["ViewImage"] {
		t.Fatalf("outils de base/vision attendus: %v", names)
	}
}

func TestToolRegistryNilFallback(t *testing.T) {
	reg := toolRegistry{}
	out, _ := reg.execute(context.Background(), toolEnv{}, "Inconnu", "{}")
	if !strings.HasPrefix(out.Text, "[erreur]") {
		t.Fatalf("outil inconnu sans repli = %q", out.Text)
	}
}
