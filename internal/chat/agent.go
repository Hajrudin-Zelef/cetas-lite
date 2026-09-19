package chat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/skills"
)

// agentWorkspaceLabel retourne le libellé d'affichage de l'espace de
// travail lié au run : nom du projet si renseigné, sinon l'espace partagé
// historique. Utilisé pour l'annonce en tête de fil (événement
// "workspace") et le message de repli worktree.
func agentWorkspaceLabel(in TurnInput, sb *Sandbox, projectName string) (string, string) {
	if strings.TrimSpace(in.ProjectID) == "" {
		return "Espace partagé", "espace partagé"
	}
	label := strings.TrimSpace(projectName)
	if label == "" {
		label = sb.Root()
	}
	mode := "projet local"
	if sb.Remote() {
		mode = "projet distant (SFTP)"
	}
	return label, mode
}

func (e *Engine) runAgent(ctx context.Context, c *Conversation, epoch int, res resolution, base []provider.Message, in TurnInput) {
	sb, err := e.agentSandbox(in)
	if err != nil {
		c.appendDelta(epoch, map[string]any{"error": err.Error()})
		return
	}
	// Espace de travail lié au run : nom du projet si renseigné, sinon
	// l'espace partagé historique. Annoncé en tête de fil pour que
	// l'utilisateur voie toujours OÙ l'agent travaille (un agent lancé
	// sans projet ne travaille PAS dans le projet affiché ailleurs).
	projectName := ""
	if pid := strings.TrimSpace(in.ProjectID); pid != "" {
		if wm := e.workspaceManager(); wm != nil {
			if p, perr := wm.Get(pid); perr == nil {
				projectName = p.Name
			}
		}
	}
	wsLabel, wsMode := agentWorkspaceLabel(in, sb, projectName)
	c.appendDelta(epoch, map[string]any{"workspace": map[string]any{"name": wsLabel, "mode": wsMode}})
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
			c.appendDelta(epoch, map[string]any{"worktree_error": "option worktree activee mais aucun depot indique — repli sur « " + wsLabel + " »"})
		}
	}
	reg := e.toolRegistry(in, sb)
	tools := reg.schemas(ctx)
	sys := []provider.Message{{Role: "system", Content: agentSystemPrompt()}}
	// Snapshot du workspace : le modèle voit la structure réelle et ne
	// devine jamais les chemins. projectName calculé plus haut (annonce
	// de l'espace de travail en tête de fil).
	sys = append(sys, workspaceSnapshotMessage(ctx, sb, projectName))
	// GitHub connecté : l'agent sait qu'il peut commit/diff/push.
	if login := githubLogin(e.st); login != "" {
		sys = append(sys, githubPromptMessage(login))
	}
	// Directive de raisonnement (imperative, en anglais) : le bouton
	// Thinking du composer pilote le raisonnement de l'agent. Phase 4 :
	// les tours purement mécaniques (simple demande de lecture) n'ont
	// pas de raisonnement obligatoire — action directe.
	sys = append(sys, provider.Message{Role: "system", Content: agentThinkDirective(in.Think, in.Text, in.Effort)})
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
		content, reasoning, err := e.agentMember(ctx, c, epoch, p, m, msgs, tools, reg, in.User, resolveEffort(in.Think, in.Text, in.Effort), &emitted, agentOpts{approve: in.Approve, plan: in.Plan, maxTokens: in.MaxTokens, nativeWeb: native, webFallback: e.webToolsFor(in), think: in.Think})
		if err == nil {
			c.appendAssistant(epoch, content, reasoning)
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

// toolExecState regroupe l'état mutable d'exécution des outils pendant un
// tour : partagé entre la voie parallèle (runs) et la voie séquentielle.
// (Phase 2 : généralisation du parallélisme au-delà des blocs 100 % lecture.)
type toolExecState struct {
	done           map[string]string
	denied         map[string]bool
	repeats        map[string]int
	modified       map[string]bool
	verified       bool
	alwaysApproved bool
	planApproved   bool
}

// parallelToolOut est le résultat d'un appel exécuté en parallèle.
type parallelToolOut struct {
	id       string
	out      ToolResult
	followup *provider.Message
}

// parallelCheck : un appel peut-il rejoindre un run parallèle ?
// Mêmes règles que l'ancien tout-ou-rien, mais évaluées appel par appel
// pour permettre le partitionnement du bloc : parallélisable uniquement,
// sans approbation, sans déduplication, sans doublon dans le run.
// Retourne les arguments parsés et la clé de dédup, ou ok=false.
// Ne mute jamais st (lecture seule + seen local au run).
func parallelCheck(tc provider.ToolCall, opts agentOpts, st *toolExecState, seen map[string]bool) (args map[string]any, key string, ok bool) {
	name := tc.Function.Name
	if !parallelSafeTools[name] || tc.InvalidCall() {
		return nil, "", false
	}
	args = parseArgs(tc.Function.Arguments)
	key = name + "\x00" + tc.Function.Arguments
	if opts.plan && !st.planApproved && !planToolAllowed(name) {
		return nil, "", false
	}
	if st.denied[key] {
		return nil, "", false
	}
	if opts.approve && !st.alwaysApproved && needsApprovalFor(name, args) {
		return nil, "", false
	}
	if _, dup := st.done[key]; dup && dedupableTool(name) {
		return nil, "", false
	}
	if seen[key] {
		return nil, "", false
	}
	seen[key] = true
	return args, key, true
}

// toolSegment est un morceau du bloc d'appels : soit un run parallèle
// (appels éligibles consécutifs), soit un appel isolé en séquentiel.
type toolSegment struct {
	parallel bool
	calls    []provider.ToolCall
}

// partitionToolBlock découpe le bloc en segments en préservant l'ordre :
// suites maximales d'appels parallélisables -> runs, le reste -> séquentiel.
// Fonction pure (testable) : l'exécution est faite par executeToolBlock.
// L'ordre relatif entre un run et les appels qui l'entourent est conservé,
// donc une lecture placée après une écriture voit toujours son effet.
func partitionToolBlock(tcs []provider.ToolCall, opts agentOpts, st *toolExecState) []toolSegment {
	var segs []toolSegment
	i := 0
	for i < len(tcs) {
		seen := map[string]bool{}
		j := i
		for j < len(tcs) {
			if _, _, ok := parallelCheck(tcs[j], opts, st, seen); !ok {
				break
			}
			j++
		}
		if j > i {
			segs = append(segs, toolSegment{parallel: true, calls: tcs[i:j]})
			i = j
		} else {
			segs = append(segs, toolSegment{parallel: false, calls: tcs[i : i+1]})
			i++
		}
	}
	return segs
}

// execParallelRun exécute en parallèle un run d'appels quand c'est sans
// risque : au moins 2 appels, tous validés par parallelCheck. Les deltas
// "start" partent dans l'ordre d'origine, puis les "end" suivent le même
// ordre : l'UI reste strictement cohérente avec la voie séquentielle. Les
// messages d'outils sont retournés dans l'ordre pour ajout à l'historique
// par l'appelant.
// Retourne (nil, false) si le run doit passer par la voie séquentielle
// (l'état a pu changer entre le partitionnement et l'exécution).
func (e *Engine) execParallelRun(ctx context.Context, c *Conversation, epoch int, reg toolRegistry, tcs []provider.ToolCall, env toolEnv, opts agentOpts, st *toolExecState) ([]parallelToolOut, bool) {
	if len(tcs) < 2 {
		return nil, false
	}
	type pending struct {
		tc   provider.ToolCall
		args map[string]any
		key  string
	}
	items := make([]pending, 0, len(tcs))
	seen := make(map[string]bool, len(tcs))
	for _, tc := range tcs {
		args, key, ok := parallelCheck(tc, opts, st, seen)
		if !ok {
			return nil, false
		}
		items = append(items, pending{tc: tc, args: args, key: key})
	}
	// Deltas "start" dans l'ordre d'origine, comme en séquentiel.
	for _, it := range items {
		if ctx.Err() != nil {
			return nil, false
		}
		c.appendDelta(epoch, map[string]any{"tool": map[string]any{
			"name": it.tc.Function.Name, "args": it.args, "phase": "start",
		}})
	}
	// Exécution concurrente ; résultats collectés par index pour garder
	// l'ordre d'origine à la restitution.
	outs := make([]parallelToolOut, len(items))
	var wg sync.WaitGroup
	for i, it := range items {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out, followup := reg.execute(ctx, env, it.tc.Function.Name, it.tc.Function.Arguments)
			outs[i] = parallelToolOut{id: it.tc.ID, out: out, followup: followup}
		}()
	}
	wg.Wait()
	// Restitution : deltas "end" et déduplication dans l'ordre, exactement
	// comme la voie séquentielle. (L'annulation éventuelle est traitée en
	// tête de boucle appelante.)
	for i, it := range items {
		o := &outs[i]
		if !strings.HasPrefix(o.out.Text, "[erreur]") {
			st.done[it.key] = o.out.Text
		}
		toolDelta := map[string]any{
			"name": it.tc.Function.Name, "args": it.args, "phase": "end",
			"result": truncate(o.out.Text, toolMaxOutput), "diff": truncateDiff(o.out.Diff, 300),
		}
		if o.out.Meta != nil {
			if srcs, ok := o.out.Meta["sources"]; ok {
				toolDelta["sources"] = srcs
			}
			if sp, ok := o.out.Meta["search_provider"]; ok {
				toolDelta["search_provider"] = sp
			}
		}
		c.appendDelta(epoch, map[string]any{"tool": toolDelta})
	}
	return outs, true
}

// executeToolBlock exécute un bloc d'appels d'outils en partitionnant en
// runs parallélisables (Phase 2) : chaque run de ≥2 appels éligibles part
// en concurrence, tout le reste passe par la voie séquentielle — dans
// l'ordre d'émission du modèle. Retourne abort=true si le tour doit
// s'interrompre (approbation expirée).
func (e *Engine) executeToolBlock(ctx context.Context, c *Conversation, epoch int, reg toolRegistry, tcs []provider.ToolCall, env toolEnv, opts agentOpts, st *toolExecState, user string, m alias.ResolvedMember, msgs *[]provider.Message) (abort bool) {
	// Point C : un log par bloc d'outils execute — journalctl montre
	// enfin l'activite du chemin chat tour par tour.
	if len(tcs) > 0 {
		names := make([]string, 0, len(tcs))
		for _, tc := range tcs {
			names = append(names, tc.Function.Name)
		}
		log.Printf("chat: tour %d: execution de %d appel(s) d'outil: %s", epoch, len(tcs), strings.Join(names, ", "))
	}
	appendOut := func(id string, out ToolResult, followup *provider.Message) {
		*msgs = append(*msgs, provider.Message{Role: "tool", ToolCallID: id, Content: truncateToolForModel(out.Text)})
		if followup != nil {
			*msgs = append(*msgs, *followup)
		}
	}
	for _, seg := range partitionToolBlock(tcs, opts, st) {
		if ctx.Err() != nil {
			return false
		}
		if seg.parallel && len(seg.calls) >= 2 {
			if pouts, ok := e.execParallelRun(ctx, c, epoch, reg, seg.calls, env, opts, st); ok {
				for _, po := range pouts {
					appendOut(po.id, po.out, po.followup)
				}
				continue
			}
			// Repli : l'état a changé depuis le partitionnement, on
			// rejoue le run en séquentiel (la déduplication s'applique).
		}
		for _, tc := range seg.calls {
			if ctx.Err() != nil {
				return false
			}
			out, followup, ab := e.execSequentialCall(ctx, c, epoch, reg, tc, env, opts, st, user, m)
			if ab {
				return true
			}
			appendOut(tc.ID, out, followup)
		}
	}
	return false
}

// execSequentialCall exécute UN appel via la voie séquentielle : appel
// irrecevable, mode plan, refus mémorisé, approbation utilisateur,
// déduplication, puis exécution. Émet les deltas start/end comme avant.
// Retourne le résultat, un éventuel message de suivi, et abort=true si le
// tour doit s'interrompre (approbation expirée).
func (e *Engine) execSequentialCall(ctx context.Context, c *Conversation, epoch int, reg toolRegistry, tc provider.ToolCall, env toolEnv, opts agentOpts, st *toolExecState, user string, m alias.ResolvedMember) (ToolResult, *provider.Message, bool) {
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
	case opts.plan && !st.planApproved && !planToolAllowed(tc.Function.Name):
		out = ToolResult{Text: "[erreur] mode plan : tu es en phase d'exploration LECTURE SEULE. " +
			"Seuls Ls, Tree, Read, Cat, Grep, Glob et TodoWrite sont autorises tant que le plan n'est pas valide. " +
			"Construis ton plan avec TodoWrite puis presente-le."}
	case st.denied[key]:
		out = ToolResult{Text: "[refuse] l'utilisateur a deja refuse cet appel pendant ce tour."}
	case opts.approve && !st.alwaysApproved && needsApprovalFor(tc.Function.Name, args):
		d, aerr := c.RequestApproval(ctx, epoch, ApprovalRequest{
			Kind: "tool", Tool: tc.Function.Name, Args: args,
		})
		if aerr != nil {
			if ctx.Err() != nil {
				return ToolResult{}, nil, true
			}
			c.appendDelta(epoch, map[string]any{"content": "\n\n_Approbation expiree : tour interrompu._"})
			return ToolResult{}, nil, true
		}
		if d.always {
			st.alwaysApproved = true
		}
		if !d.approved {
			st.denied[key] = true
			out = ToolResult{Text: "[refuse] l'utilisateur a refuse l'execution de " + tc.Function.Name +
				". Propose une alternative ou demande des precisions au lieu de reessayer a l'identique."}
			break
		}
		out, followup = reg.execute(ctx, env, tc.Function.Name, tc.Function.Arguments)
		if !strings.HasPrefix(out.Text, "[erreur]") {
			st.done[key] = out.Text
			trackModification(tc.Function.Name, args, st.modified)
			if tc.Function.Name == "Bash" {
				if cmd, ok := args["command"].(string); ok && verifyCommandHeuristic(cmd) {
					st.verified = true
				}
			}
		}
	case func() bool { _, seen := st.done[key]; return seen && dedupableTool(tc.Function.Name) }():
		st.repeats[key]++
		out = ToolResult{Text: repeatedCallResult(st.done[key], st.repeats[key])}
	default:
		out, followup = reg.execute(ctx, env, tc.Function.Name, tc.Function.Arguments)
		if !strings.HasPrefix(out.Text, "[erreur]") {
			st.done[key] = out.Text
			trackModification(tc.Function.Name, args, st.modified)
			if tc.Function.Name == "Bash" {
				if cmd, ok := args["command"].(string); ok && verifyCommandHeuristic(cmd) {
					st.verified = true
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
	return out, followup, false
}

func (e *Engine) agentMember(ctx context.Context, c *Conversation, epoch int, p provider.Provider, m alias.ResolvedMember, base []provider.Message, tools []provider.Tool, reg toolRegistry, user, effort string, emitted *bool, opts agentOpts) (string, string, error) {
	const maxNudges = 2
	// Point B : le nudge "tentative d'appel en texte" est borne separement
	// (compteur dedie) : au-dela, le texte reste visible tel quel au lieu
	// de couter des allers-retours modele.
	const maxTextCallNudges = 2
	msgs := append([]provider.Message(nil), base...)
	// Phase 2 : l'état mutable d'exécution des outils est regroupé pour
	// être partagé entre la voie parallèle (runs) et la voie séquentielle.
	st := &toolExecState{
		done:     map[string]string{},
		denied:   map[string]bool{},
		repeats:  map[string]int{},
		modified: map[string]bool{},
	}
	nudges := 0
	textCallNudges := 0
	disableTools := false
	transientRetries := 0
	last := ""
	// lastReasoning suit last tour par tour : le raisonnement du dernier
	// tour est persiste avec le message assistant final (point 2 du fix
	// reasoning_content).
	lastReasoning := ""

	// Mode plan : phase 1 en lecture seule, la phase d'execution demarre
	// seulement apres validation du plan par l'utilisateur.
	fullTools := tools
	nativeFallbackDone := false
	if opts.plan {
		tools = readOnlyTools(tools)
		msgs = append(msgs, provider.Message{Role: "system", Content: planModePrompt()})
	}

	// Suivi du workflow plan -> code -> verification.
	verifyNudged := false

	// Filtre DSML du tour : recree a chaque tentative dans la boucle.
	var dsmlFilter *dsmlStreamFilter

	emit := func(ev provider.Event) bool {
		if ev.Reasoning != "" && opts.think {
			c.appendDelta(epoch, map[string]any{"reasoning_content": ev.Reasoning})
		}
		if ev.Content != "" {
			// Les blocs DSML (appels d'outils en texte) sont masques pendant
			// le streaming : ils seront convertis en vrais appels d'outils
			// une fois le tour termine, jamais affiches en brut.
			if visible := dsmlFilter.push(ev.Content); visible != "" {
				*emitted = true
				c.appendDelta(epoch, map[string]any{"content": visible})
			}
		}
		if ev.Usage != nil {
			c.appendDelta(epoch, map[string]any{"stats": map[string]any{
				"prompt_tokens":     ev.Usage.PromptTokens,
				"completion_tokens": ev.Usage.CompletionTokens,
			}})
		}
		return true
	}

	// flushDSML vide le reliquat de texte normal du filtre DSML vers le fil
	// (jamais de DSML), puis repart sur un filtre neuf pour la tentative
	// suivante. Chaque octet n'est emis qu'une fois : pas de duplication.
	flushDSML := func() {
		if dsmlFilter == nil {
			return
		}
		if tail := dsmlFilter.flush(); tail != "" {
			*emitted = true
			c.appendDelta(epoch, map[string]any{"content": tail})
		}
		dsmlFilter = &dsmlStreamFilter{}
	}

	for {
		if ctx.Err() != nil {
			return last, lastReasoning, nil
		}
		dsmlFilter = &dsmlStreamFilter{}
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
		flushDSML()
		if err == nil && opts.nativeWeb {
			c.appendDelta(epoch, map[string]any{"search": searchSourcesDelta(resp.Annotations, true)})
		}
		if err != nil {
			if ctx.Err() != nil {
				return last, lastReasoning, nil
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
			if errors.As(err, &he) && isTransientHTTP(he.Status) && !*emitted && transientRetries < 2 {
				// Erreur transitoire avant toute emission visible : on rejoue
				// la requete a l'identique, outils intacts. Aucune duplication
				// possible (rien n'a ete montre ni execute).
				transientRetries++
				continue
			}
			if errors.As(err, &he) && !disableTools && len(tools) > 0 {
				disableTools = true
				msgs = append(msgs, provider.Message{Role: "system", Content: "Stop calling tools. Answer directly now using only the information already gathered."})
				continue
			}
			return last, lastReasoning, err
		}

		// Le reasoning du tour est capture avec son texte : il sera
		// persiste avec le message assistant final (appendAssistant).
		// En mode thinking, DeepSeek exige reasoning_content sur tous
		// les messages assistant des qu'un seul en porte un — voir
		// withReasoningContentForced.
		lastReasoning = resp.Reasoning

		// Phase 1 : filet unique "appels en texte" (DSML ou pseudo-appel
		// isole) — un seul point d'entree, voir parseFallbackToolCalls.
		// Les appels convertis passent par le pipeline standard
		// (approbation, blocs d'outils, suivi des modifications) et le
		// balisage est retire du texte affiche dans tous les cas.
		if len(resp.ToolCalls) == 0 {
			tcalls, stripped := parseFallbackToolCalls(resp.Content, tools)
			resp.ToolCalls, resp.Content = tcalls, stripped
		} else if hasDSML(resp.Content) {
			resp.Content = stripDSMLFinal(resp.Content)
		}

		if len(resp.ToolCalls) > 0 {
			*emitted = true
			assistant := provider.Message{Role: "assistant", ToolCalls: resp.ToolCalls}
			if resp.Content != "" {
				assistant.Content = resp.Content
			}
			// Phase 0 (etendue) : en mode thinking, le raisonnement est
			// attache a chaque message assistant, meme vide — DeepSeek
			// exige le champ reasoning_content sur TOUS les messages
			// assistant des qu'un seul en porte un (HTTP 400 sinon), pas
			// seulement sur ceux avec tool_calls. La serialisation force
			// le champ sur le wire, voir withReasoningContentForced.
			assistant.ReasoningContent = resp.Reasoning
			msgs = append(msgs, assistant)
			// Phase 2 : le bloc est partitionné en runs parallélisables
			// (suites maximales d'appels indépendants) ; chaque run de ≥2
			// appels part en concurrence, le reste passe en séquentiel —
			// dans l'ordre d'émission, avec les mêmes deltas et le même
			// historique que la voie séquentielle. Voir executeToolBlock.
			if e.executeToolBlock(ctx, c, epoch, reg, resp.ToolCalls, toolEnv{user: user, member: m}, opts, st, user, m, &msgs) {
				return last, lastReasoning, nil
			}
			last = resp.Content
			continue
		}

		last = resp.Content
		// Phase 1 : le nudge "tentative d'appel en texte" est supprime...
		// Point B : ...puis restaure sous forme bornee (filet de securite).
		// Les pseudo-appels convertibles sont executes directement par le
		// filet unique (parseFallbackToolCalls) sans aller-retour ; quand
		// le contenu ressemble VRAIMENT a une tentative d'appel annoncee
		// (nom d'outil connu + mot-cle d'annonce, voir
		// looksLikeAnnouncedCall) mais que le parsing strict l'a refusee,
		// on demande au modele de reemettre de vrais appels — au plus
		// maxTextCallNudges fois par tour. Le declenchement exige un nom
		// d'outil reconnu : jamais de nudge sur une reponse texte normale.
		if strings.TrimSpace(resp.Content) != "" && !disableTools && len(tools) > 0 && looksLikeAnnouncedCall(resp.Content, tools) {
			if textCallNudges < maxTextCallNudges {
				textCallNudges++
				log.Printf("chat: tour %d: tentative d'appel en texte non convertible, nudge %d/%d", epoch, textCallNudges, maxTextCallNudges)
				msgs = append(msgs, provider.Message{Role: "user", Content: textCallNudgeText()})
				continue
			}
			// Point C : budget epuise — le tour se termine sans outil
			// alors que le texte ressemble a un appel : loggue pour
			// journalctl au lieu de rester silencieux.
			log.Printf("chat: tour %d: termine sans outil alors que le texte ressemble a un appel (budget nudge epuise)", epoch)
		}
		// Seul le nudge "reponse vide" subsiste (autre cas d'echec, sans rapport).
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
		if !verifyNudged && !st.verified && len(st.modified) > 0 && !disableTools && len(tools) > 0 {
			verifyNudged = true
			files := modifiedList(st.modified)
			c.appendDelta(epoch, map[string]any{"drop_reasoning": true})
			msgs = append(msgs, provider.Message{Role: "user", Content: "Tu as modifie ces fichiers sans les verifier : " + strings.Join(files, ", ") + ". " +
				"Avant de conclure, VERIFIE ton travail maintenant : compile et/ou lance les tests " +
				"avec l'outil Bash (ex : go build ./..., go test ./...). N'ecris pas de resume avant d'avoir verifie."})
			continue
		}
		// Mode plan : la phase d'exploration est terminee, on soumet le plan
		// a l'utilisateur avant de passer a l'execution.
		if opts.plan && !st.planApproved && !disableTools {
			planText := buildPlanText(msgs, last)
			d, aerr := c.RequestApproval(ctx, epoch, ApprovalRequest{Kind: "plan", Plan: planText})
			if aerr != nil {
				if ctx.Err() == nil {
					c.appendDelta(epoch, map[string]any{"content": "\n\n_Approbation du plan expiree : tour interrompu._"})
				}
				return last, lastReasoning, nil
			}
			if !d.approved {
				c.appendDelta(epoch, map[string]any{"content": "\n\n_Plan refuse par l'utilisateur._"})
				return last, lastReasoning, nil
			}
			st.planApproved = true
			tools = fullTools
			msgs = append(msgs, provider.Message{Role: "system", Content: "Plan approved by the user. Execute it now with all tools, " +
				"then VERIFY your work (compile/test) before concluding."})
			c.appendDelta(epoch, map[string]any{"content": "\n\n_Plan valide, execution en cours..._\n\n"})
			continue
		}
		return last, lastReasoning, nil
	}
}

// isTransientHTTP : erreurs HTTP transitoires qui meritent un reessai
// (rate-limit, timeout passerelle, passerelle indisponible...).
func isTransientHTTP(status int) bool {
	switch status {
	case 408, 429, 502, 503, 504:
		return true
	}
	return false
}

// streamWithRetry appelle le provider en reessayant les erreurs HTTP
// transitoires avec backoff borne, tant que rien n'a encore ete diffuse (pas
// de duplication). Les autres erreurs suivent le chemin degrade existant
// (desactivation des outils / failover vers le membre suivant du pool).
func (e *Engine) streamWithRetry(ctx context.Context, p provider.Provider, req provider.Request, emit func(provider.Event) bool, emitted *bool) (provider.Response, error) {
	var resp provider.Response
	var err error
	backoffs := []time.Duration{time.Second, 2 * time.Second, 4 * time.Second}
	for attempt := 0; ; attempt++ {
		resp, err = p.Stream(ctx, req, emit)
		if err == nil {
			return resp, nil
		}
		var he *provider.HTTPError
		retryable := errors.As(err, &he) && isTransientHTTP(he.Status)
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

// parallelSafeTools est l'allowlist explicite des outils parallélisables
// par les runs parallèles : lectures pures, sans aucun effet de bord.
// TodoWrite est volontairement exclu : il mute la liste de tâches
// (état UI + dérivation du plan) et ses appels ont une sémantique
// d'ordre ("étape 1 terminée, étape 2 en cours"). Echo est pur
// (simple renvoi de texte) et les outils GitHub listés sont des GET.
var parallelSafeTools = map[string]bool{
	"Read": true, "Ls": true, "Tree": true,
	"Grep": true, "Glob": true, "Echo": true,
	"GitHubRepos": true, "GitHubIssues": true,
	"GitHubIssueGet": true, "GitHubPRs": true,
	// Phase 3 : Cat est fusionné dans Read (alias) ; un appel résiduel
	// garde la sémantique lecture pure -> parallélisable.
	"Cat": true,
}

// parallelSafe indique si un outil peut s'exécuter en concurrence avec
// d'autres appels du même bloc.
func parallelSafe(name string) bool { return parallelSafeTools[name] }

// planAllowedTools est l'allowlist stricte du mode plan : lecture seule
// (exploration du workspace) + TodoWrite pour construire le plan.
// Tout autre outil est refuse a l'execution tant que le plan n'est pas
// valide, meme s'il n'exigerait pas d'approbation hors mode plan.
var planAllowedTools = map[string]bool{
	"Ls": true, "Tree": true, "Read": true,
	"Grep": true, "Glob": true, "Echo": true, "TodoWrite": true,
	"GitHubRepos": true, "GitHubIssues": true, "GitHubIssueGet": true, "GitHubPRs": true,
	// Phase 3 : alias de Read (fusion) — lecture pure autorisée en plan.
	"Cat": true,
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

// textCallNudgeText : nudge "tentative d'appel en texte" (point B), borne
// a maxTextCallNudges par tour. Demande au modele de reemettre de vrais
// appels au lieu d'annoncer en texte.
func textCallNudgeText() string {
	return "Your last message looked like a tool call written as text (e.g. \"Appel réel : ...\"), but no tool was actually called. In your NEXT message, emit the real tool call(s) directly as function calls, with no announcement text around them. Do not describe the call — make it."
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
