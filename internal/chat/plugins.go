package chat

import (
	"context"

	"cetas-lite/internal/plugins"
	"cetas-lite/internal/provider"
)

// PluginManager : abstraction du gestionnaire de plugins externes.
type PluginManager interface {
	Defs() []plugins.Def
	Execute(ctx context.Context, name, argsJSON string) (string, error)
	Plugins() []plugins.Info
	Errors() []string
	Load() error
}

// SetPlugins branche le gestionnaire de plugins externes sur le moteur.
func (e *Engine) SetPlugins(p PluginManager) {
	e.mu.Lock()
	e.pluginMgr = p
	e.mu.Unlock()
}

func (e *Engine) pluginManager() PluginManager {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.pluginMgr
}

// PluginStatus expose l'etat des plugins pour l'API web.
func (e *Engine) PluginStatus() (infos []plugins.Info, errs []string) {
	if pm := e.pluginManager(); pm != nil {
		return pm.Plugins(), pm.Errors()
	}
	return nil, nil
}

// PluginReload recharge les plugins depuis le dossier.
func (e *Engine) PluginReload() error {
	if pm := e.pluginManager(); pm != nil {
		return pm.Load()
	}
	return nil
}

func PluginToolSchemas(defs []plugins.Def) []provider.Tool {
	out := make([]provider.Tool, 0, len(defs))
	for _, d := range defs {
		params := any(d.Schema)
		if d.Schema == nil {
			params = map[string]any{"type": "object", "properties": map[string]any{}}
		}
		out = append(out, provider.Tool{Type: "function", Function: provider.ToolFunction{
			Name:        d.Name,
			Description: "[Plugin] " + d.Description,
			Parameters:  params,
		}})
	}
	return out
}

func (e *Engine) pluginExecute(ctx context.Context, name, argsJSON string) ToolResult {
	pm := e.pluginManager()
	if pm == nil {
		return ToolResult{Text: "[error] plugins unavailable"}
	}
	out, err := pm.Execute(ctx, name, argsJSON)
	if err != nil {
		return ToolResult{Text: "[erreur plugin] " + err.Error()}
	}
	return ToolResult{Text: out}
}
