package chat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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
	sb.AllowScript = e.scriptAllowed()
	sb.Isolation = e.isolation()
	reg := e.toolRegistry(in, sb)
	tools := reg.schemas(ctx)
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
		content, err := e.agentMember(ctx, c, epoch, p, m, msgs, tools, reg, in.User, resolveEffort(true, true, in.Text, in.Effort), &emitted, agentOpts{approve: in.Approve, plan: in.Plan})
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

func (e *Engine) agentMember(ctx context.Context, c *Conversation, epoch int, p provider.Provider, m alias.ResolvedMember, base []provider.Message, tools []provider.Tool, reg toolRegistry, user, effort string, emitted *bool, opts agentOpts) (string, error) {
	const maxNudges = 2
	msgs := append([]provider.Message(nil), base...)
	done := map[string]string{}
	denied := map[string]bool{}
	repeats := map[string]int{}
	nudges := 0
	disableTools := false
	last := ""

	// Mode plan : phase 1 en lecture seule, la phase d'execution demarre
	// seulement apres validation du plan par l'utilisateur.
	fullTools := tools
	planApproved := false
	if opts.plan {
		tools = readOnlyTools(tools)
		msgs = append(msgs, provider.Message{Role: "system", Content: planModePrompt()})
	}

	// Suivi du workflow plan -> code -> verification.
	modified := map[string]bool{}
	verified := false
	verifyNudged := false
	alwaysApproved := false

	emit := func(ev provider.Event) bool {
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
	}

	for {
		if ctx.Err() != nil {
			return last, nil
		}
		var toolSet []provider.Tool
		if !disableTools {
			toolSet = tools
		}
		resp, err := e.streamWithRetry(ctx, p, provider.Request{
			Model:           m.Model,
			Messages:        normalizeSystemMessages(msgs),
			Tools:           toolSet,
			Temperature:     0.7,
			EnableReasoning: true,
			ReasoningEffort: effort,
		}, emit, emitted)
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
				switch {
				case tc.InvalidCall():
					out = ToolResult{Text: "[erreur] appel d'outil irrecevable (nom vide ou arguments JSON incomplets). " +
						"Renvoie exactement le meme appel avec un nom d'outil valide et des arguments JSON complets."}
				case opts.plan && !planApproved && needsApproval(tc.Function.Name):
					out = ToolResult{Text: "[erreur] mode plan : tu es en phase d'exploration. " +
						"Les outils d'ecriture et d'execution sont interdits tant que le plan n'est pas valide. " +
						"Construis ton plan avec TodoWrite puis presente-le."}
				case denied[key]:
					out = ToolResult{Text: "[refuse] l'utilisateur a deja refuse cet appel pendant ce tour."}
				case opts.approve && !alwaysApproved && needsApproval(tc.Function.Name):
					d, aerr := c.RequestApproval(ctx, epoch, ApprovalRequest{
						Kind: "tool", Tool: tc.Function.Name, Args: args,
					})
					if aerr != nil {
						if ctx.Err() != nil {
							return last, nil
						}
						c.appendDelta(epoch, map[string]any{"content": "\n\n_Approbation expiree : tour interrompu._"})
						return last, nil
					}
					if d.always {
						alwaysApproved = true
					}
					if !d.approved {
						denied[key] = true
						out = ToolResult{Text: "[refuse] l'utilisateur a refuse l'execution de " + tc.Function.Name +
							". Propose une alternative ou demande des precisions au lieu de reessayer a l'identique."}
						break
					}
					out, followup = reg.execute(ctx, toolEnv{user: user, member: m}, tc.Function.Name, tc.Function.Arguments)
					if !strings.HasPrefix(out.Text, "[erreur]") {
						done[key] = out.Text
						trackModification(tc.Function.Name, args, modified)
						if tc.Function.Name == "Bash" {
							if cmd, ok := args["command"].(string); ok && verifyCommandHeuristic(cmd) {
								verified = true
							}
						}
					}
				case func() bool { _, seen := done[key]; return seen && dedupableTool(tc.Function.Name) }():
					repeats[key]++
					out = ToolResult{Text: repeatedCallResult(done[key], repeats[key])}
				default:
					out, followup = reg.execute(ctx, toolEnv{user: user, member: m}, tc.Function.Name, tc.Function.Arguments)
					if !strings.HasPrefix(out.Text, "[erreur]") {
						done[key] = out.Text
						trackModification(tc.Function.Name, args, modified)
						if tc.Function.Name == "Bash" {
							if cmd, ok := args["command"].(string); ok && verifyCommandHeuristic(cmd) {
								verified = true
							}
						}
					}
				}
				c.appendDelta(epoch, map[string]any{"tool": map[string]any{
					"name": tc.Function.Name, "args": args, "phase": "end",
					"result": truncate(out.Text, toolMaxOutput), "diff": truncateDiff(out.Diff, 300),
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
		// Phase VERIFY du workflow : si du code a ete ecrit mais jamais verifie,
		// exiger une verification avant de conclure (une seule fois).
		if !verifyNudged && !verified && len(modified) > 0 && !disableTools && len(tools) > 0 {
			verifyNudged = true
			files := modifiedList(modified)
			c.appendDelta(epoch, map[string]any{"drop_reasoning": true})
			msgs = append(msgs, provider.Message{Role: "user", Content: "Tu as modifie ces fichiers sans les verifier : " + strings.Join(files, ", ") + ". " +
				"Avant de conclure, VERIFIE ton travail maintenant : compile et/ou lance les tests " +
				"avec l'outil Bash (ex : go build ./..., go test ./...). N'ecris pas de resume avant d'avoir verifie."})
			continue
		}
		// Mode plan : la phase d'exploration est terminee, on soumet le plan
		// a l'utilisateur avant de passer a l'execution.
		if opts.plan && !planApproved && !disableTools {
			planText := buildPlanText(msgs, last)
			d, aerr := c.RequestApproval(ctx, epoch, ApprovalRequest{Kind: "plan", Plan: planText})
			if aerr != nil {
				if ctx.Err() == nil {
					c.appendDelta(epoch, map[string]any{"content": "\n\n_Approbation du plan expiree : tour interrompu._"})
				}
				return last, nil
			}
			if !d.approved {
				c.appendDelta(epoch, map[string]any{"content": "\n\n_Plan refuse par l'utilisateur._"})
				return last, nil
			}
			planApproved = true
			tools = fullTools
			msgs = append(msgs, provider.Message{Role: "system", Content: "Plan valide par l'utilisateur. Execute-le maintenant avec tous les outils, " +
				"puis VERIFIE ton travail (compile/teste) avant de conclure."})
			c.appendDelta(epoch, map[string]any{"content": "\n\n_Plan valide, execution en cours..._\n\n"})
			continue
		}
		return last, nil
	}
}

// streamWithRetry appelle le provider en reessayant les 429 (rate-limit
// transitoire) avec backoff, tant que rien n'a encore ete diffuse (pas de
// duplication). Les autres erreurs HTTP suivent le chemin degrade existant
// (desactivation des outils / failover vers le membre suivant du pool).
func (e *Engine) streamWithRetry(ctx context.Context, p provider.Provider, req provider.Request, emit func(provider.Event) bool, emitted *bool) (provider.Response, error) {
	var resp provider.Response
	var err error
	backoffs := []time.Duration{time.Second, 2 * time.Second}
	for attempt := 0; ; attempt++ {
		resp, err = p.Stream(ctx, req, emit)
		if err == nil {
			return resp, nil
		}
		var he *provider.HTTPError
		retryable := errors.As(err, &he) && he.Status == 429
		if !retryable || *emitted || attempt >= len(backoffs) {
			return resp, err
		}
		select {
		case <-ctx.Done():
			return resp, ctx.Err()
		case <-time.After(backoffs[attempt]):
		}
	}
}

// trackModification enregistre les fichiers touches par Write/Edit reussis.
func trackModification(tool string, args map[string]any, modified map[string]bool) {
	if tool != "Write" && tool != "Edit" {
		return
	}
	if f, ok := args["file_path"].(string); ok && strings.TrimSpace(f) != "" {
		modified[f] = true
	}
}

func modifiedList(modified map[string]bool) []string {
	out := make([]string, 0, len(modified))
	for f := range modified {
		out = append(out, f)
	}
	return out
}

// truncateDiff borne la taille des diffs envoyes au frontend.
func truncateDiff(d []DiffLine, max int) []DiffLine {
	if len(d) <= max {
		return d
	}
	out := append([]DiffLine(nil), d[:max]...)
	return append(out, DiffLine{Kind: "…", Text: fmt.Sprintf("(%d lignes masquees)", len(d)-max)})
}

// agentOpts regroupe les options de controle du tour agent.
type agentOpts struct {
	// approve demande une validation utilisateur avant chaque outil
	// d'ecriture ou d'execution.
	approve bool
	// plan active le mode plan : exploration en lecture seule, puis
	// validation du plan avant execution.
	plan bool
}

// needsApproval indique si un outil exige une validation utilisateur avant
// execution (ecriture, execution, outils externes).
func needsApproval(name string) bool {
	switch name {
	case "Write", "Edit", "Bash", "RunScript":
		return true
	}
	return strings.HasPrefix(name, "mcp_") || strings.HasPrefix(name, "custom_")
}

// readOnlyTools ne garde que les outils de lecture et de planification.
func readOnlyTools(tools []provider.Tool) []provider.Tool {
	keep := map[string]bool{"Ls": true, "Read": true, "Grep": true, "Glob": true, "TodoWrite": true}
	out := make([]provider.Tool, 0, len(tools))
	for _, t := range tools {
		if keep[t.Function.Name] {
			out = append(out, t)
		}
	}
	return out
}

func planModePrompt() string {
	return "MODE PLAN : tu es en phase d'exploration. Utilise uniquement les outils " +
		"de lecture (Ls, Read, Grep, Glob) et TodoWrite pour construire un plan d'action. " +
		"Quand ton exploration est terminee, presente ton plan clairement dans ta reponse " +
		"(etapes numerotees) et attends la validation : n'appelle AUCUN outil d'ecriture " +
		"ou d'execution pendant cette phase."
}

// buildPlanText reconstitue le plan a soumettre : derniers todos + conclusion.
func buildPlanText(msgs []provider.Message, content string) string {
	var todos string
	for i := len(msgs) - 1; i >= 0; i-- {
		for _, tc := range msgs[i].ToolCalls {
			if tc.Function.Name == "TodoWrite" {
				todos = formatTodos(parseArgs(tc.Function.Arguments))
				break
			}
		}
		if todos != "" {
			break
		}
	}
	var b strings.Builder
	if todos != "" {
		b.WriteString(todos)
		b.WriteString("\n\n")
	}
	b.WriteString(strings.TrimSpace(content))
	return strings.TrimSpace(b.String())
}

func formatTodos(args map[string]any) string {
	raw, ok := args["todos"].([]any)
	if !ok || len(raw) == 0 {
		return ""
	}
	var b strings.Builder
	for i, item := range raw {
		m, ok := item.(map[string]any)
		if !ok {
			continue
		}
		content, _ := m["content"].(string)
		status, _ := m["status"].(string)
		mark := "[ ]"
		switch status {
		case "in_progress":
			mark = "[~]"
		case "completed":
			mark = "[x]"
		}
		fmt.Fprintf(&b, "%d. %s %s\n", i+1, mark, strings.TrimSpace(content))
	}
	return strings.TrimSpace(b.String())
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
