package chat

import (
	"context"
	"encoding/json"

	"cetas-lite/internal/mcp"
	"cetas-lite/internal/provider"
)

type MCPTools interface {
	Tools(ctx context.Context) []mcp.Tool
	Call(ctx context.Context, name string, args map[string]any) (string, error)
	Servers() []mcp.ServerStatus
}

func MCPToolSchemas(defs []mcp.Tool) []provider.Tool {
	out := make([]provider.Tool, 0, len(defs))
	for _, d := range defs {
		params := any(map[string]any{"type": "object", "properties": map[string]any{}})
		if len(d.InputSchema) > 0 {
			var v any
			if json.Unmarshal(d.InputSchema, &v) == nil {
				params = v
			}
		}
		desc := d.Description
		if desc == "" {
			desc = d.Exposed
		}
		out = append(out, provider.Tool{Type: "function", Function: provider.ToolFunction{
			Name:        d.Exposed,
			Description: "[MCP:" + d.Server + "] " + desc,
			Parameters:  params,
		}})
	}
	return out
}

func (e *Engine) mcpSchemas(ctx context.Context) []provider.Tool {
	mt := e.mcpTools()
	if mt == nil {
		return nil
	}
	defs := mt.Tools(ctx)
	if len(defs) == 0 {
		return nil
	}
	return MCPToolSchemas(defs)
}

func (e *Engine) mcpExecute(ctx context.Context, name, argsJSON string) ToolResult {
	mt := e.mcpTools()
	if mt == nil {
		return ToolResult{Text: "[error] MCP unavailable"}
	}
	out, err := mt.Call(ctx, name, parseArgs(argsJSON))
	if err != nil {
		return ToolResult{Text: "[error] MCP: " + err.Error()}
	}
	return ToolResult{Text: truncate(out, toolMaxOutput)}
}
