package chat

import (
	"context"
	"log"
	"strings"
	"time"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/customtools"
	"cetas-lite/internal/provider"
)

type toolEnv struct {
	user   string
	member alias.ResolvedMember
}

type toolFamily interface {
	schemas(ctx context.Context) []provider.Tool
	handles(name string) bool
	execute(ctx context.Context, env toolEnv, name, argsJSON string) (ToolResult, *provider.Message)
}

type toolRegistry struct {
	families []toolFamily
	fallback toolFamily
	// sb sert uniquement a la mesure F6.1 (label du backend local/sftp).
	// Peut etre nil (tests) : backendKind est protege.
	sb *Sandbox
}

func (r toolRegistry) schemas(ctx context.Context) []provider.Tool {
	out := append([]provider.Tool(nil), r.fallback.schemas(ctx)...)
	for _, f := range r.families {
		out = append(out, f.schemas(ctx)...)
	}
	return out
}

func (r toolRegistry) execute(ctx context.Context, env toolEnv, name, argsJSON string) (ToolResult, *provider.Message) {
	for _, f := range r.families {
		if f.handles(name) {
			// Phase 2 : les erreurs sont uniformisées au point de passage
			// unique ([erreur] <outil> : <cause> — <consigne>), quelle que
			// soit la famille d'outils.
			out, followup := timedExecute(r.sb, name, func() (ToolResult, *provider.Message) {
				return f.execute(ctx, env, name, argsJSON)
			})
			return uniformToolError(name, out), followup
		}
	}
	if r.fallback != nil {
		out, followup := timedExecute(r.sb, name, func() (ToolResult, *provider.Message) {
			return r.fallback.execute(ctx, env, name, argsJSON)
		})
		return uniformToolError(name, out), followup
	}
	return ToolResult{Text: "[erreur] outil inconnu: " + name}, nil
}

// timedExecute mesure la duree d'une execution d'outil reelle (F6.1, phase 1).
// Point de passage unique : voies sequentielle ET parallele. L'attente
// d'approbation est exclue (reg.execute n'est appele qu'apres decision).
// Format journalctl : chat: outil Read [sftp]: 342 ms (marqueur LENT >= 2 s).
func timedExecute(sb *Sandbox, name string, fn func() (ToolResult, *provider.Message)) (ToolResult, *provider.Message) {
	start := time.Now()
	out, followup := fn()
	elapsed := time.Since(start)
	slow := ""
	if elapsed >= 2*time.Second {
		slow = " LENT"
	}
	log.Printf("chat: outil %s [%s]: %d ms%s", name, sb.backendKind(), elapsed.Milliseconds(), slow)
	return out, followup
}

func (e *Engine) toolRegistry(in TurnInput, sb *Sandbox) toolRegistry {
	families := []toolFamily{visionFamily{e: e, sb: sb}}
	if e.memoryTools() != nil {
		families = append(families, memoryFamily{e: e})
	}
	if rt := e.ragTools(); rt != nil && rt.Ready() {
		families = append(families, ragFamily{e: e})
	}
	if ct := e.customTools(); ct != nil {
		families = append(families, customFamily{e: e, defs: ct.Defs()})
	}
	if e.pluginManager() != nil {
		families = append(families, pluginFamily{e: e})
	}
	if in.MCP {
		families = append(families, mcpFamily{e: e})
	}
	if e.webToolsFor(in) {
		families = append(families, webFamily{e: e})
	}
	return toolRegistry{families: families, fallback: builtinFamily{sb: sb, allowScript: e.scriptAllowed()}, sb: sb}
}

type builtinFamily struct {
	sb          *Sandbox
	allowScript bool
}

func (f builtinFamily) schemas(context.Context) []provider.Tool {
	ts := ToolSchemas()
	if !f.allowScript {
		ts = filterTools(ts, "RunScript")
	}
	return ts
}

func (builtinFamily) handles(string) bool { return true }

func (f builtinFamily) execute(ctx context.Context, env toolEnv, name, argsJSON string) (ToolResult, *provider.Message) {
	return f.sb.Execute(ctx, name, argsJSON), nil
}

type visionFamily struct {
	e  *Engine
	sb *Sandbox
}

func (visionFamily) schemas(context.Context) []provider.Tool {
	return []provider.Tool{viewImageSchema()}
}

func (visionFamily) handles(name string) bool { return name == "ViewImage" }

func (f visionFamily) execute(ctx context.Context, env toolEnv, _ string, argsJSON string) (ToolResult, *provider.Message) {
	return f.e.viewImage(ctx, env.member, f.sb, argsJSON)
}

type memoryFamily struct{ e *Engine }

func (memoryFamily) schemas(context.Context) []provider.Tool { return MemoryToolSchemas() }

func (memoryFamily) handles(name string) bool { return strings.HasPrefix(name, "mem_") }

func (f memoryFamily) execute(_ context.Context, env toolEnv, name, argsJSON string) (ToolResult, *provider.Message) {
	return f.e.memExecute(env.user, name, argsJSON), nil
}

type ragFamily struct{ e *Engine }

func (ragFamily) schemas(context.Context) []provider.Tool { return RagToolSchemas() }

func (ragFamily) handles(name string) bool { return strings.HasPrefix(name, "rag_") }

func (f ragFamily) execute(ctx context.Context, _ toolEnv, name, argsJSON string) (ToolResult, *provider.Message) {
	return f.e.ragExecute(ctx, name, argsJSON), nil
}

type mcpFamily struct{ e *Engine }

func (f mcpFamily) schemas(ctx context.Context) []provider.Tool { return f.e.mcpSchemas(ctx) }

func (mcpFamily) handles(name string) bool { return strings.HasPrefix(name, "mcp_") }

func (f mcpFamily) execute(ctx context.Context, _ toolEnv, name, argsJSON string) (ToolResult, *provider.Message) {
	return f.e.mcpExecute(ctx, name, argsJSON), nil
}

type customFamily struct {
	e    *Engine
	defs []customtools.Def
}

func (f customFamily) schemas(context.Context) []provider.Tool { return CustomToolSchemas(f.defs) }

func (customFamily) handles(name string) bool { return strings.HasPrefix(name, "custom_") }

func (f customFamily) execute(ctx context.Context, _ toolEnv, name, argsJSON string) (ToolResult, *provider.Message) {
	return f.e.customExecute(ctx, name, argsJSON), nil
}

type pluginFamily struct{ e *Engine }

func (f pluginFamily) schemas(context.Context) []provider.Tool {
	if pm := f.e.pluginManager(); pm != nil {
		return PluginToolSchemas(pm.Defs())
	}
	return nil
}

func (pluginFamily) handles(name string) bool { return strings.HasPrefix(name, "plugin_") }

func (f pluginFamily) execute(ctx context.Context, _ toolEnv, name, argsJSON string) (ToolResult, *provider.Message) {
	return f.e.pluginExecute(ctx, name, argsJSON), nil
}

type webFamily struct{ e *Engine }

func (webFamily) schemas(context.Context) []provider.Tool { return WebToolSchemas() }

func (webFamily) handles(name string) bool {
	return name == "web_search" || name == "web_fetch"
}

func (f webFamily) execute(ctx context.Context, env toolEnv, name, argsJSON string) (ToolResult, *provider.Message) {
	return f.e.webExecute(ctx, env.user, name, argsJSON), nil
}
