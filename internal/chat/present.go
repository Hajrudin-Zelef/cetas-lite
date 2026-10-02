package chat

import (
	"context"
	"path/filepath"
	"strconv"
	"strings"

	"cetas-lite/internal/provider"
)

// PresentFile : outil agent qui LIVRE un fichier de l'espace de travail comme
// pièce jointe cliquable dans le fil (carte téléchargeable / ouvrable).
// Distinct de Write : l'agent écrit via Write, puis présente les fichiers
// destinés à l'utilisateur via PresentFile (pas de carte pour chaque fichier
// intermédiaire).

func presentFileSchema() provider.Tool {
	return provider.Tool{Type: "function", Function: provider.ToolFunction{
		Name:        "PresentFile",
		Description: "Present a workspace file to the user as a downloadable attachment (md, txt, py, json, csv, …). Write the file first, then call this with its path. Use for files the user should keep; not for intermediate files.",
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"file_path": map[string]any{"type": "string", "description": "Path of the file in the workspace."},
				"name":      map[string]any{"type": "string", "description": "Optional display/download name (defaults to the base name)."},
			},
			"required": []string{"file_path"},
		},
	}}
}

// presentFamily : famille d'outils de livraison de fichier (mode agent).
type presentFamily struct {
	e  *Engine
	sb *Sandbox
}

func (presentFamily) schemas(context.Context) []provider.Tool {
	return []provider.Tool{presentFileSchema()}
}

func (presentFamily) handles(name string) bool { return name == "PresentFile" }

func (f presentFamily) execute(ctx context.Context, env toolEnv, _ string, argsJSON string) (ToolResult, *provider.Message) {
	return f.e.presentFile(ctx, env.user, f.sb, argsJSON), nil
}

func (e *Engine) presentFile(ctx context.Context, user string, sb *Sandbox, argsJSON string) ToolResult {
	args := parseArgs(argsJSON)
	rel := strings.TrimSpace(strArg(args, "file_path"))
	if rel == "" {
		return ToolResult{Text: "[error] empty path"}
	}
	data, err := sb.ReadFile(ctx, rel)
	if err != nil {
		return ToolResult{Text: "[error] " + err.Error()}
	}
	display := strings.TrimSpace(strArg(args, "name"))
	if display == "" {
		display = filepath.Base(rel)
	}
	st := e.attachments()
	if st == nil {
		return ToolResult{Text: "[error] piece jointe indisponible"}
	}
	a, err := st.Save(user, display, data)
	if err != nil {
		return ToolResult{Text: "[error] " + err.Error()}
	}
	return ToolResult{
		Text: "[ok] " + a.Name + " présenté (" + strconv.Itoa(len(data)) + " octets)",
		Meta: map[string]any{"attachment": AttachmentInfo{ID: a.ID, Name: a.Name, Kind: a.Kind}},
	}
}
