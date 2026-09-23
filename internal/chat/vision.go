package chat

import (
	"context"
	"encoding/base64"
	"path/filepath"
	"strings"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/attach"
	"cetas-lite/internal/modelcaps"
	"cetas-lite/internal/provider"
)

func viewImageSchema() provider.Tool {
	return provider.Tool{Type: "function", Function: provider.ToolFunction{
		Name:        "ViewImage",
		Description: "View an image file from the workspace (vision models only).",
		Parameters: map[string]any{
			"type":       "object",
			"properties": map[string]any{"file_path": map[string]any{"type": "string"}},
			"required":   []string{"file_path"},
		},
	}}
}

func (e *Engine) viewImage(ctx context.Context, m alias.ResolvedMember, sb *Sandbox, argsJSON string) (ToolResult, *provider.Message) {
	if !e.capabilities().Vision(m.Provider, m.Model) {
		return ToolResult{Text: "[error] this model cannot read images; choose a vision model (Settings -> Model capabilities)."}, nil
	}
	rel := strings.TrimSpace(strArg(parseArgs(argsJSON), "file_path"))
	data, err := sb.ReadFile(ctx, rel)
	if err != nil {
		return ToolResult{Text: "[error] " + err.Error()}, nil
	}
	ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(rel)), ".")
	switch ext {
	case "png", "jpg", "jpeg", "gif", "webp":
	default:
		return ToolResult{Text: "[error] image file expected: " + rel}, nil
	}
	img := map[string]any{
		"type":      "image_url",
		"image_url": map[string]any{"url": "data:" + mimeFor(ext) + ";base64," + base64.StdEncoding.EncodeToString(data)},
	}
	msg := provider.Message{Role: "user", Content: []any{
		map[string]any{"type": "text", "text": "Image " + rel + " :"},
		img,
	}}
	return ToolResult{Text: "[image] " + rel}, &msg
}

func (e *Engine) SetCapabilities(m modelcaps.Map) {
	e.mu.Lock()
	e.caps = m
	e.mu.Unlock()
}

func (e *Engine) capabilities() modelcaps.Map {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.caps
}

func (e *Engine) Capabilities() modelcaps.Map {
	m := e.capabilities()
	if m == nil {
		return modelcaps.Map{}
	}
	out := make(modelcaps.Map, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}

func (e *Engine) hasImageAttachment(user string, ids []string) bool {
	st := e.attachments()
	if st == nil {
		return false
	}
	for _, id := range ids {
		a, _, err := st.Get(user, id)
		if err == nil && a.Kind == attach.KindImage {
			return true
		}
	}
	return false
}

func (e *Engine) filterVision(members []alias.ResolvedMember) []alias.ResolvedMember {
	caps := e.capabilities()
	out := make([]alias.ResolvedMember, 0, len(members))
	for _, m := range members {
		if caps.Vision(m.Provider, m.Model) {
			out = append(out, m)
		}
	}
	return out
}

func (e *Engine) expandImageParts(msgs []provider.Message, user string, ids []string) []provider.Message {
	st := e.attachments()
	if st == nil {
		return msgs
	}
	var images []any
	for _, id := range ids {
		a, data, err := st.Get(user, id)
		if err != nil || a.Kind != attach.KindImage {
			continue
		}
		images = append(images, map[string]any{
			"type":      "image_url",
			"image_url": map[string]any{"url": "data:" + mimeFor(a.Ext) + ";base64," + base64.StdEncoding.EncodeToString(data)},
		})
	}
	if len(images) == 0 {
		return msgs
	}
	idx := -1
	for i := len(msgs) - 1; i >= 0; i-- {
		if msgs[i].Role == "user" {
			idx = i
			break
		}
	}
	if idx < 0 {
		return msgs
	}
	text, _ := msgs[idx].Content.(string)
	parts := make([]any, 0, len(images)+1)
	parts = append(parts, map[string]any{"type": "text", "text": text})
	parts = append(parts, images...)

	out := append([]provider.Message(nil), msgs...)
	out[idx].Content = parts
	return out
}

func mimeFor(ext string) string {
	switch ext {
	case "png":
		return "image/png"
	case "jpg", "jpeg":
		return "image/jpeg"
	case "gif":
		return "image/gif"
	case "webp":
		return "image/webp"
	}
	return "application/octet-stream"
}
