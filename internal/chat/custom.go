package chat

import (
	"context"

	"cetas-lite/internal/customtools"
	"cetas-lite/internal/provider"
)

type CustomTools interface {
	Defs() []customtools.Def
	Call(ctx context.Context, name, argsJSON string) (string, error)
}

func CustomToolSchemas(defs []customtools.Def) []provider.Tool {
	out := make([]provider.Tool, 0, len(defs))
	for _, d := range defs {
		params := any(d.Schema)
		if d.Schema == nil {
			params = map[string]any{"type": "object", "properties": map[string]any{}}
		}
		out = append(out, provider.Tool{Type: "function", Function: provider.ToolFunction{
			Name:        d.Name,
			Description: "[Outil] " + d.Description,
			Parameters:  params,
		}})
	}
	return out
}

func (e *Engine) customExecute(ctx context.Context, name, argsJSON string) ToolResult {
	ct := e.customTools()
	if ct == nil {
		return ToolResult{Text: "[error] custom tools unavailable"}
	}
	out, err := ct.Call(ctx, name, argsJSON)
	if err != nil {
		return ToolResult{Text: "[error] " + err.Error()}
	}
	return ToolResult{Text: truncate(out, toolMaxOutput)}
}
