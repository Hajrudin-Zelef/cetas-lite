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
	"cetas-lite/internal/skills"
)

func (e *Engine) runAgent(ctx context.Context, c *Conversation, epoch int, res resolution, base []provider.Message, in TurnInput) {
	sb, err := e.agentSandbox(in)
	if err != nil {
		c.appendDelta(epoch, map[string]any{"error": err.Error()})
		return
	}
	// Worktree d'isolation : l'agent travaille dans un checkout dedie du
	// depot au lieu du workspace partage. En cas d'echec, on continue sur
	// le workspace normal (erreur explicite, pas de silence).
	if in.Worktree {
		if repo := strings.TrimSpace(in.Repo); repo != "" {
			if wm := e.worktreeManager(); wm != nil {
				if wtPath, werr := wm.Ensure(repo, c.ID); werr == nil {
					if wsb, serr := NewSandbox(wtPath); serr == nil {
						wsb.AllowScript = sb.AllowScript
						wsb.Isolation = sb.Isolation
						wsb.GitHubToken = sb.GitHubToken
						sb = wsb
						e.noteWorktreeRepo(c.ID, repo)
						if run := e.agentForConv(c.ID); run != nil {
							run.setWorktreePath(wtPath)
						}
						c.appendDelta(epoch, map[string]any{"worktree": map[string]any{"path": wtPath, "repo": repo}})
					} else {
						c.appendDelta(epoch, map[string]any{"worktree_error": serr.Error()})
					}
				} else {
					c.appendDelta(epoch, map[string]any{"worktree_error": werr.Error()})
				}
			}
		} else {
			c.appendDelta(epoch, map[string]any{"worktree_error": "option worktree activee mais aucun depot indique"})
		}
	}
	reg := e.toolRegistry(in, sb)
	tools := reg.schemas(ctx)
	sys := []provider.Message{{Role: "system", Content: agentSystemPrompt()}}
	// Snapshot du workspace : le modèle voit la structure réelle et ne
	// devine jamais les chemins. Nom du projet si renseigné.
	projectName := ""
	if pid := strings.TrimSpace(in.ProjectID); pid != "" {
		if wm := e.workspaceManager(); wm != nil {
			if p, perr := wm.Get(pid); perr == nil {
				projectName = p.Name
			}
		}
	}
	sys = append(sys, workspaceSnapshotMessage(ctx, sb, projectName))
	// GitHub connecté : l'agent sait qu'il peut commit/diff/push.
	if login := githubLogin(e.st); login != "" {
		sys = append(sys, githubPromptMessage(login))
	}
	// Directive de raisonnement (imperative, en anglais) : le bouton
	// Thinking du composer pilote le raisonnement de l'agent.
	sys = append(sys, provider.Message{Role: "system", Content: thinkDirective(true, in.Think, in.Effort)})
	// Competences actives (Configuration -> Competences) : instructions
	// utilisateur injectees dans le prompt systeme de l'agent.
	if e.st != nil {
		if sk := skills.ActiveInstructions(e.st, in.User); sk != "" {
			sys = append(sys, provider.Message{Role: "system", Content: sk})
		}
	}
	// Recherche native du membre principal, calculee une seule fois (limiteur).
	// En recherche approfondie ("deep"), le natif est desactive : l'agent
	// passe par les outils web_search/web_fetch pour multiplier les
	// requetes et lire les pages en entier (le plugin natif, borne a
	// 5 resultats, ne le permet pas).
	deep := in.WebDepth == "deep"
	nativePrimary := len(res.members) > 0 && !deep && e.useNativeWebSearch(in, res.members[0].Provider, res.members[0].Model)
	// Directive de recherche web (imperative, en anglais) : le globe est
	// l'interrupteur principal. Le bouton Thinking pilote le
	// raisonnement (payload), avec effort selectionnable via in.Effort.
	sys = append(sys, provider.Message{Role: "system", Content: searchDirective(e.webEnabled(in), nativePrimary)})
	// Recherche approfondie : l'agent multiplie les requetes, elargit les
	// resultats et lit les pages en entier au lieu de se contenter des
	// extraits.
	if in.WebDepth == "deep" && e.webEnabled(in) {
		sys = append(sys, provider.Message{Role: "system", Content: deepWebDirective()})
	}
	if mm, ok, notice := e.marexForSession(c.ID); ok {
		sys = append(sys, mm)
		if notice {
			c.appendDelta(epoch, map[string]any{"system": "MAREX.md chargé"})
		}
	}
	msgs := normalizeSystemMessages(append(sys, base...))
	// Directive web courante : recalculee par membre, car un membre non natif
	// ne doit pas recevoir la directive "provider-native" du membre principal.
	webDir := searchDirective(e.webEnabled(in), nativePrimary)

	var lastErr error
	for i, m := range res.members {
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
		native := nativePrimary
		if i > 0 && !deep {
			native = e.useNativeWebSearch(in, m.Provider, m.Model)
		}
		if dir := searchDirective(e.webEnabled(in), native); dir != webDir {
			msgs = replaceWebDirective(msgs, webDir, dir)
			webDir = dir
		}
		content, err := e.agentMember(ctx, c, epoch, p, m, msgs, tools, reg, in.User, resolveEffort(in.Think, in.Text, in.Effort), &emitted, agentOpts{approve: in.Approve, plan: in.Plan, maxTokens: in.MaxTokens, nativeWeb: native, webFallback: e.webToolsFor(in), think: in.Think})
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
	nativeFallbackDone := false
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
		if ev.Reasoning != "" && opts.think {
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
		// Recherche native : les outils web deviennent redondants, on les
		// retire pour eviter une double recherche (natif + outils).
		if opts.nativeWeb {
			toolSet = dropWebTools(toolSet)
		}
		req := provider.Request{
			Model:           m.Model,
			Messages:        normalizeSystemMessages(msgs),
			Tools:           toolSet,
			Temperature:     0.7,
			MaxTokens:       opts.maxTokens,
			EnableReasoning: opts.think,
			ReasoningEffort: effort,
		}
		if opts.nativeWeb {
			req.Extra = nativeWebExtraFor(m.Provider)
			c.appendDelta(epoch, map[string]any{"search": map[string]any{"phase": "start", "native": true}})
		}
		resp, err := e.streamWithRetry(ctx, p, req, emit, emitted)
		if err == nil && opts.nativeWeb {
			c.appendDelta(epoch, map[string]any{"search": searchSourcesDelta(resp.Annotations, true)})
		}
		if err != nil {
			if ctx.Err() != nil {
				return last, nil
			}
			// Repli natif -> outils : la recherche native a echoue avant
			// toute emission. On retente une seule fois sans le plugin natif :
			// les outils web (retires pour le natif) sont reinjectes.
			if opts.nativeWeb && !*emitted && !nativeFallbackDone && opts.webFallback {
				nativeFallbackDone = true
				opts.nativeWeb = false
				c.appendDelta(epoch, map[string]any{"search": searchSourcesDelta(nil, true)})
				// Le natif est desactive pour la 2e tentative : la directive
				// ne doit plus parler de recherche native.
				msgs = replaceWebDirective(msgs, searchDirective(true, true), searchDirective(true, false))
				continue
			}
			var he *provider.HTTPError
			if errors.As(err, &he) && !disableTools && len(tools) > 0 {
				disableTools = true
				msgs = append(msgs, provider.Message{Role: "system", Content: "Stop calling tools. Answer directly now using only the information already gathered."})
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
				case opts.plan && !planApproved && !planToolAllowed(tc.Function.Name):
					out = ToolResult{Text: "[erreur] mode plan : tu es en phase d'exploration LECTURE SEULE. " +
						"Seuls Ls, Tree, Read, Cat, Grep, Glob et TodoWrite sont autorises tant que le plan n'est pas valide. " +
						"Construis ton plan avec TodoWrite puis presente-le."}
				case denied[key]:
					out = ToolResult{Text: "[refuse] l'utilisateur a deja refuse cet appel pendant ce tour."}
				case opts.approve && !alwaysApproved && needsApprovalFor(tc.Function.Name, args):
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
				toolDelta := map[string]any{
					"name": tc.Function.Name, "args": args, "phase": "end",
					"result": truncate(out.Text, toolMaxOutput), "diff": truncateDiff(out.Diff, 300),
				}
				// Sources structurees (recherche web) pour le panneau "Sources".
				if out.Meta != nil {
					if srcs, ok := out.Meta["sources"]; ok {
						toolDelta["sources"] = srcs
					}
					if sp, ok := out.Meta["search_provider"]; ok {
						toolDelta["search_provider"] = sp
					}
				}
				c.appendDelta(epoch, map[string]any{"tool": toolDelta})
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
			msgs = append(msgs, provider.Message{Role: "system", Content: "Plan approved by the user. Execute it now with all tools, " +
				"then VERIFY your work (compile/test) before concluding."})
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

// trackModification enregistre les fichiers/dossiers touches par les outils
// d'ecriture reussis (revue de fin de tour).
func trackModification(tool string, args map[string]any, modified map[string]bool) {
	var f string
	switch tool {
	case "Write", "Edit":
		f, _ = args["file_path"].(string)
	case "Mkdir":
		f, _ = args["path"].(string)
	case "Mv":
		f, _ = args["dst"].(string)
	case "Sed":
		if b, ok := args["in_place"].(bool); ok && b {
			f, _ = args["file"].(string)
		}
	default:
		return
	}
	if strings.TrimSpace(f) != "" {
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
	// maxTokens limite les tokens generes par reponse (0 = defaut).
	maxTokens int
	// nativeWeb : la recherche web passe par le plugin natif du provider
	// (OpenRouter / DeepSeek) au lieu des outils web_search/web_fetch.
	nativeWeb bool
	// webFallback : si la recherche native echoue avant toute emission,
	// retenter une fois avec les outils web_search/web_fetch (mode auto).
	webFallback bool
	// think pilote le raisonnement du modele (bouton Thinking du
	// composer) : false desactive le raisonnement cote provider et la
	// directive systeme ordonne une reponse directe.
	think bool
}

// needsApproval indique si un outil exige une validation utilisateur avant
// execution (ecriture, execution, outils externes).
func needsApproval(name string) bool {
	switch name {
	case "Write", "Edit", "Bash", "RunScript", "Mkdir", "Mv", "Curl",
		"GitHubRepoCreate", "GitHubIssueCreate", "GitHubIssueComment",
		"GitHubPRCreate", "GitHubPRMerge":
		return true
	}
	return strings.HasPrefix(name, "mcp_") || strings.HasPrefix(name, "custom_") || strings.HasPrefix(name, "plugin_")
}

// needsApprovalFor affine needsApproval en tenant compte des arguments :
// Sed et Awk n'exigent une approbation que pour leurs usages a effets de
// bord (sed -i, commandes d'execution/ecriture, system(), tubes shell...).
// En mode flux pur (lecture seule), ils s'executent sans friction.
func needsApprovalFor(name string, args map[string]any) bool {
	if needsApproval(name) {
		return true
	}
	switch name {
	case "Sed":
		return sedNeedsApproval(args)
	case "Awk":
		return awkNeedsApproval(args)
	}
	return false
}

// planAllowedTools est l'allowlist stricte du mode plan : lecture seule
// (exploration du workspace) + TodoWrite pour construire le plan.
// Tout autre outil est refuse a l'execution tant que le plan n'est pas
// valide, meme s'il n'exigerait pas d'approbation hors mode plan.
var planAllowedTools = map[string]bool{
	"Ls": true, "Tree": true, "Read": true, "Cat": true,
	"Grep": true, "Glob": true, "Echo": true, "TodoWrite": true,
	"GitHubRepos": true, "GitHubIssues": true, "GitHubIssueGet": true, "GitHubPRs": true,
}

// planToolAllowed indique si un outil peut s'executer en mode plan
// avant validation du plan.
func planToolAllowed(name string) bool { return planAllowedTools[name] }

// readOnlyTools ne garde que les outils autorises en mode plan.
func readOnlyTools(tools []provider.Tool) []provider.Tool {
	out := make([]provider.Tool, 0, len(tools))
	for _, t := range tools {
		if planAllowedTools[t.Function.Name] {
			out = append(out, t)
		}
	}
	return out
}

func planModePrompt() string {
	return "PLAN MODE: you are in the exploration phase. Use only read tools " +
		"(Ls, Tree, Read, Cat, Grep, Glob) and TodoWrite to build an action plan. " +
		"When exploration is done, present your plan clearly in your answer " +
		"(numbered steps) and wait for validation: do NOT call any write " +
		"or execution tool during this phase."
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
	if strings.HasPrefix(name, "mcp_") || strings.HasPrefix(name, "custom_") || strings.HasPrefix(name, "plugin_") || name == "ViewImage" {
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
