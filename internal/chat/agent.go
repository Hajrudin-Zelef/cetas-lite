package chat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/provider"
)

func (e *Engine) runAgent(ctx context.Context, c *Conversation, epoch int, res resolution, base []provider.Message, in TurnInput) {
	root, err := userWorkspacePath(e.workspace, in.User)
	if err != nil {
		c.appendDelta(epoch, map[string]any{"error": err.Error()})
		return
	}
	sb, err := NewSandbox(root)
	if err != nil {
		c.appendDelta(epoch, map[string]any{"error": err.Error()})
		return
	}
	tools := ToolSchemas()
	if !e.scriptAllowed() {
		tools = filterTools(tools, "RunScript")
	}
	tools = append(tools, viewImageSchema())
	if e.memoryTools() != nil {
		tools = append(tools, MemoryToolSchemas()...)
	}
	if ct := e.customTools(); ct != nil {
		tools = append(tools, CustomToolSchemas(ct.Defs())...)
	}
	if in.MCP {
		tools = append(tools, e.mcpSchemas(ctx)...)
	}
	if in.Web && e.webTools() != nil {
		tools = append(tools, WebToolSchemas()...)
	}
	sb.AllowScript = e.scriptAllowed()
	sb.Isolation = e.isolation()
	sys := []provider.Message{{Role: "system", Content: agentSystemPrompt()}}
	if mm, ok := e.marexMessage(in.User); ok {
		sys = append(sys, mm)
	}
	msgs := normalizeSystemMessages(append(sys, base...))

	var lastErr error
	for _, m := range res.members {
		if ctx.Err() != nil {
			c.appendDelta(epoch, map[string]any{"content": "\n\n_Génération interrompue._"})
			return
		}
		p, ok := e.reg.Get(m.Provider)
		if !ok {
			lastErr = fmt.Errorf("provider %s non configure", m.Provider)
			continue
		}
		c.appendDelta(epoch, routeDelta(m, in.Family, in.Mode, false, res.fallback))

		emitted := false
		content, err := e.agentMember(ctx, c, epoch, p, m, msgs, tools, sb, in.User, resolveEffort(true, true, in.Text, in.Effort), &emitted)
		if err == nil {
			c.appendAssistant(epoch, content)
			return
		}
		lastErr = err
		if ctx.Err() != nil {
			c.appendDelta(epoch, map[string]any{"content": "\n\n_Génération interrompue._"})
			return
		}
		if emitted {
			c.appendDelta(epoch, map[string]any{"error": err.Error()})
			return
		}
	}
	if lastErr == nil {
		lastErr = fmt.Errorf("aucun modele disponible")
	}
	c.appendDelta(epoch, map[string]any{"error": lastErr.Error()})
}

func (e *Engine) agentMember(ctx context.Context, c *Conversation, epoch int, p provider.Provider, m alias.ResolvedMember, base []provider.Message, tools []provider.Tool, sb *Sandbox, user, effort string, emitted *bool) (string, error) {
	const maxNudges = 2
	msgs := append([]provider.Message(nil), base...)
	done := map[string]string{}
	repeats := map[string]int{}
	nudges := 0
	disableTools := false
	last := ""

	for {
		if ctx.Err() != nil {
			return last, nil
		}
		var toolSet []provider.Tool
		if !disableTools {
			toolSet = tools
		}
		resp, err := p.Stream(ctx, provider.Request{
			Model:           m.Model,
			Messages:        normalizeSystemMessages(msgs),
			Tools:           toolSet,
			Temperature:     0.7,
			EnableReasoning: true,
			ReasoningEffort: effort,
		}, func(ev provider.Event) bool {
			if ev.Reasoning != "" {
				c.appendDelta(epoch, map[string]any{"reasoning_content": ev.Reasoning})
			}
			if ev.Content != "" {
				*emitted = true
				c.appendDelta(epoch, map[string]any{"content": ev.Content})
			}
			if ev.Usage != nil {
				c.appendDelta(epoch, map[string]any{"stats": map[string]any{
					"prompt_tokens":     ev.Usage.PromptTokens,
					"completion_tokens": ev.Usage.CompletionTokens,
				}})
			}
			return true
		})
		if err != nil {
			if ctx.Err() != nil {
				return last, nil
			}
			var he *provider.HTTPError
			if errors.As(err, &he) && !disableTools && len(tools) > 0 {
				disableTools = true
				msgs = append(msgs, provider.Message{Role: "system", Content: "N'appelle plus d'outil. Reponds maintenant directement a partir des informations deja obtenues."})
				continue
			}
			return last, err
		}

		if len(resp.ToolCalls) > 0 {
			*emitted = true
			assistant := provider.Message{Role: "assistant", ToolCalls: resp.ToolCalls}
			if resp.Content != "" {
				assistant.Content = resp.Content
			}
			msgs = append(msgs, assistant)
			for _, tc := range resp.ToolCalls {
				if ctx.Err() != nil {
					return last, nil
				}
				args := parseArgs(tc.Function.Arguments)
				c.appendDelta(epoch, map[string]any{"tool": map[string]any{
					"name": tc.Function.Name, "args": args, "phase": "start",
				}})
				key := tc.Function.Name + "\x00" + tc.Function.Arguments
				var out ToolResult
				var followup *provider.Message
				if prev, seen := done[key]; seen && dedupableTool(tc.Function.Name) {
					repeats[key]++
					out.Text = repeatedCallResult(prev, repeats[key])
				} else {
					if tc.Function.Name == "web_search" || tc.Function.Name == "web_fetch" {
						out = e.webExecute(ctx, user, tc.Function.Name, tc.Function.Arguments)
					} else if tc.Function.Name == "ViewImage" {
						out, followup = e.viewImage(m, sb, tc.Function.Arguments)
					} else if strings.HasPrefix(tc.Function.Name, "mem_") {
						out = e.memExecute(user, tc.Function.Name, tc.Function.Arguments)
					} else if strings.HasPrefix(tc.Function.Name, "mcp_") {
						out = e.mcpExecute(ctx, tc.Function.Name, tc.Function.Arguments)
					} else if strings.HasPrefix(tc.Function.Name, "custom_") {
						out = e.customExecute(ctx, tc.Function.Name, tc.Function.Arguments)
					} else {
						out = sb.Execute(ctx, tc.Function.Name, tc.Function.Arguments)
					}
					if !strings.HasPrefix(out.Text, "[erreur]") {
						done[key] = out.Text
					}
				}
				c.appendDelta(epoch, map[string]any{"tool": map[string]any{
					"name": tc.Function.Name, "args": args, "phase": "end",
					"result": truncate(out.Text, toolMaxOutput), "diff": out.Diff,
				}})
				msgs = append(msgs, provider.Message{Role: "tool", ToolCallID: tc.ID, Content: out.Text})
				if followup != nil {
					msgs = append(msgs, *followup)
				}
			}
			last = resp.Content
			continue
		}

		last = resp.Content
		if strings.TrimSpace(resp.Content) == "" {
			if !disableTools && nudges < maxNudges && len(tools) > 0 {
				nudges++
				c.appendDelta(epoch, map[string]any{"drop_reasoning": true})
				msgs = append(msgs, provider.Message{Role: "user", Content: nudgeText(nudges)})
				continue
			}
			c.appendDelta(epoch, map[string]any{"content": "_(le modele n'a pas produit de reponse)_"})
		}
		return last, nil
	}
}

func userWorkspacePath(base, user string) (string, error) {
	base = strings.TrimSpace(base)
	if base == "" {
		return "", errors.New("workspace non configure")
	}
	safe := strings.ReplaceAll(user, "/", "_")
	safe = strings.ReplaceAll(safe, "\\", "_")
	safe = strings.TrimSpace(safe)
	if safe == "" || safe == "." || safe == ".." {
		return "", errors.New("utilisateur invalide")
	}
	return filepath.Join(base, safe), nil
}

func routeDelta(m alias.ResolvedMember, family, mode string, local, fallback bool) map[string]any {
	return map[string]any{"route": map[string]any{
		"provider": m.Provider, "model": m.Model, "label": m.Label,
		"family": family, "mode": mode,
		"local": local, "fallback": fallback,
	}}
}

func dedupableTool(name string) bool {
	if strings.HasPrefix(name, "mcp_") || strings.HasPrefix(name, "custom_") || name == "ViewImage" {
		return false
	}
	return name != "Bash" && name != "RunScript"
}

func repeatedCallResult(prev string, repeats int) string {
	if repeats >= 2 {
		return "[deja fait] Cet appel exact a deja ete execute " + strconv.Itoa(repeats) +
			" fois dans ce tour ; son resultat est plus haut dans la conversation. " +
			"Ne le redemande plus : reponds avec ce que tu as, ou change d'approche."
	}
	return "[deja fait] Appel identique deja execute dans ce tour — non rejoue. " +
		"Voici a nouveau son resultat ; ne le redemande pas une troisieme fois.\n\n" + prev
}

func nudgeText(n int) string {
	if n > 1 {
		return "You are stuck re-describing the same plan without executing it. Stop reasoning. In your NEXT message, either call ONE tool right now, or write your final answer in plain text using only what you already know — no more planning, no more thinking, act or answer this instant."
	}
	return "You reasoned but did not call a tool or answer. Act NOW: call the appropriate tool directly, or give your final answer if you already have the info. Don't explain, act."
}

func normalizeSystemMessages(msgs []provider.Message) []provider.Message {
	var sys []string
	out := make([]provider.Message, 0, len(msgs))
	for _, m := range msgs {
		if m.Role == "system" {
			if s, ok := m.Content.(string); ok && s != "" {
				sys = append(sys, s)
			}
			continue
		}
		out = append(out, m)
	}
	if len(sys) > 0 {
		out = append([]provider.Message{{Role: "system", Content: strings.Join(sys, "\n\n")}}, out...)
	}
	return out
}

func parseArgs(raw string) map[string]any {
	args := map[string]any{}
	_ = json.Unmarshal([]byte(raw), &args)
	return args
}
