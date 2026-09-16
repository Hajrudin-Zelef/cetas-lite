package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/attach"
	"cetas-lite/internal/local"
	"cetas-lite/internal/mcp"
	"cetas-lite/internal/modelcaps"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/store"
	"cetas-lite/internal/workspace"
	"cetas-lite/internal/worktree"

	"golang.org/x/time/rate"
)

type Engine struct {
	reg         *provider.Registry
	discover    *local.Discoverer
	st          *store.Store
	workspace   string
	allowScript bool
	searcher    WebTools
	mem         MemoryTools
	ext         MCPTools
	custom      CustomTools
	pluginMgr   PluginManager
	attach      *attach.Store
	caps        modelcaps.Map
	sandboxMode string

	mu         sync.Mutex
	families   []alias.Family
	convs      map[string]*Conversation
	webLimiter map[string]*rate.Limiter
	worktrees  *worktree.Manager

	// Runs d'agents paralleles et depot associe aux worktrees par
	// conversation (chat principal inclus).
	agentsMu sync.Mutex
	agents   map[string]*AgentRun
	wtRepos  map[string]string

	// Projets (upload local / dossier SFTP) : workspace de l'agent.
	wsProjects *workspace.Manager

	// MAREX.md : chemin du fichier renseigne par l'utilisateur, et cache
	// du contenu lu au demarrage de chaque session (par ID de conversation).
	marexMu       sync.Mutex
	marexPath     string
	marexSessions map[string]*marexEntry
}

func NewEngine(reg *provider.Registry, families []alias.Family, st *store.Store, disc *local.Discoverer, workspace string) *Engine {
	if families == nil {
		families = alias.Defaults()
	}
	return &Engine{reg: reg, discover: disc, st: st, workspace: workspace, families: families, convs: map[string]*Conversation{}}
}

func (e *Engine) Families() []alias.Family {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.families
}

func (e *Engine) SetFamilies(f []alias.Family) {
	e.mu.Lock()
	e.families = f
	e.mu.Unlock()
}

func (e *Engine) SetAllowScript(v bool) {
	e.mu.Lock()
	e.allowScript = v
	e.mu.Unlock()
}

func (e *Engine) SetSearcher(w WebTools) {
	e.mu.Lock()
	e.searcher = w
	if w != nil && e.webLimiter == nil {
		e.webLimiter = map[string]*rate.Limiter{}
	}
	e.mu.Unlock()
}

func (e *Engine) webTools() WebTools {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.searcher
}

func (e *Engine) SetMemory(m MemoryTools) {
	e.mu.Lock()
	e.mem = m
	e.mu.Unlock()
}

func (e *Engine) memoryTools() MemoryTools {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.mem
}

func (e *Engine) SetMCP(m MCPTools) {
	e.mu.Lock()
	e.ext = m
	e.mu.Unlock()
}

func (e *Engine) mcpTools() MCPTools {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.ext
}

func (e *Engine) MCPServers() []mcp.ServerStatus {
	if m := e.mcpTools(); m != nil {
		return m.Servers()
	}
	return nil
}

func (e *Engine) MCPProbe(ctx context.Context) []mcp.ServerStatus {
	m := e.mcpTools()
	if m == nil {
		return nil
	}
	m.Tools(ctx)
	return m.Servers()
}

func (e *Engine) SetCustom(c CustomTools) {
	e.mu.Lock()
	e.custom = c
	e.mu.Unlock()
}

func (e *Engine) customTools() CustomTools {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.custom
}

func (e *Engine) allowWeb(user string) bool {
	e.mu.Lock()
	if e.searcher == nil {
		e.mu.Unlock()
		return false
	}
	if e.webLimiter == nil {
		e.webLimiter = map[string]*rate.Limiter{}
	}
	l, ok := e.webLimiter[user]
	if !ok {
		l = rate.NewLimiter(rate.Every(3*time.Second), 5)
		e.webLimiter[user] = l
	}
	e.mu.Unlock()
	return l.Allow()
}

func (e *Engine) scriptAllowed() bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.allowScript
}

func (e *Engine) SetIsolation(mode string) {
	e.mu.Lock()
	e.sandboxMode = mode
	e.mu.Unlock()
}

func (e *Engine) isolation() string {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.sandboxMode
}

func (e *Engine) Conversation(user string) *Conversation {
	e.mu.Lock()
	defer e.mu.Unlock()
	if c, ok := e.convs[user]; ok {
		return c
	}
	c := NewConversation(newID(), e, func(c *Conversation) { e.save(user, c) })
	if e.st != nil {
		if raw, ok := e.st.GetConversation(user, "active"); ok {
			var s snapshot
			if err := json.Unmarshal(raw, &s); err == nil && s.ID != "" {
				c.load(s)
			}
		}
	}
	e.convs[user] = c
	return c
}

func (e *Engine) ArchiveAndReset(user string) {
	c := e.Conversation(user)
	e.archiveCurrent(user, c)
	if repo := e.takeWorktreeRepo(c.ID); repo != "" {
		e.releaseConversationWorktree(c.ID, repo)
	}
	c.Reset()
}

func (e *Engine) Regenerate(user string) error {
	c := e.Conversation(user)
	c.mu.Lock()
	if c.Generating {
		c.mu.Unlock()
		return ErrBusy
	}
	last := c.lastTurn
	if last == nil {
		c.mu.Unlock()
		return ErrNoTurn
	}
	idx := -1
	for i := len(c.Messages) - 1; i >= 0; i-- {
		if c.Messages[i].Role == "user" {
			idx = i
			break
		}
	}
	if idx < 0 {
		c.mu.Unlock()
		return ErrNoTurn
	}
	c.Messages = append([]provider.Message(nil), c.Messages[:idx]...)
	cut := 0
	for i := len(c.Log) - 1; i >= 0; i-- {
		if _, ok := c.Log[i].Delta["user"]; ok {
			cut = i
			break
		}
	}
	c.Log = append([]LogEvent(nil), c.Log[:cut]...)
	c.epoch++
	c.cond.Broadcast()
	in := TurnInput{User: user, Family: last.Family, Mode: last.Mode, Text: last.Text, Web: last.Web, WebDepth: last.WebDepth, MCP: last.MCP, Think: last.Think, Effort: last.Effort, Approve: last.Approve, Plan: last.Plan, Worktree: last.Worktree, Repo: last.Repo, ProjectID: last.ProjectID, Attachments: last.Attachments}
	c.mu.Unlock()
	if c.persist != nil {
		c.persist(c)
	}
	return c.StartTurn(in)
}

func (e *Engine) archiveCurrent(user string, c *Conversation) {
	if e.st == nil || c.isEmpty() {
		return
	}
	data, err := c.save()
	if err == nil {
		_ = e.st.ArchiveConversation(user, c.ID, data)
	}
}

func (e *Engine) RestoreArchive(user, archiveID string) bool {
	if e.st == nil {
		return false
	}
	raw, ok := e.st.GetArchive(user, archiveID)
	if !ok {
		return false
	}
	var s snapshot
	if err := json.Unmarshal(raw, &s); err != nil || s.ID == "" {
		return false
	}
	c := e.Conversation(user)
	e.archiveCurrent(user, c)
	c.restore(s)
	return true
}

func (e *Engine) save(user string, c *Conversation) {
	data, err := json.Marshal(c.marshal())
	if err != nil {
		return
	}
	_ = e.st.PutConversation(user, "active", data)
}

type resolution struct {
	members  []alias.ResolvedMember
	local    bool
	fallback bool
	agent    bool
}

func (e *Engine) resolve(ctx context.Context, in TurnInput) resolution {
	fam, ok := alias.Find(e.Families(), in.Family)
	if !ok {
		return resolution{}
	}
	if fam.Local {
		for _, m := range fam.Modes {
			if m.ID != in.Mode {
				continue
			}
			if e.discover != nil {
				if models := e.discover.ModelsForEngine(ctx, m.Engine); len(models) > 0 {
					var members []alias.ResolvedMember
					for _, mod := range models {
						members = append(members, alias.ResolvedMember{Provider: m.Engine, Model: mod.ID, Label: mod.ID})
					}
					return resolution{members: members, local: true}
				}
			}
			fb, ok := alias.Resolve(e.Families(), "samagent-n4", "standard")
			if ok {
				return resolution{members: fb.Pool, fallback: true}
			}
			return resolution{}
		}
		return resolution{}
	}
	rm, ok := alias.Resolve(e.Families(), in.Family, in.Mode)
	if !ok {
		return resolution{}
	}
	members := rm.Pool
	if in.Family == "samagent-nano" {
		// Nano : tirage aleatoire du premier modele gratuit, puis
		// fallback sequentiel sur le reste en cas d'echec.
		members = alias.ShufflePool(members)
	}
	return resolution{members: members, agent: rm.Agent && in.AgentMode}
}

func (e *Engine) Run(ctx context.Context, c *Conversation, epoch int, in TurnInput) {
	start := time.Now()
	res := e.resolve(ctx, in)
	defer func() {
		if ctx.Err() == nil {
			ct, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			if e.maybeCompact(ct, c, epoch, res.members) {
				c.appendDelta(epoch, map[string]any{"compact": true})
			}
			cancel()
		}
		c.finishTurn(epoch, time.Since(start))
	}()

	if len(res.members) == 0 {
		c.appendDelta(epoch, map[string]any{"error": "aucun modele disponible pour cet alias"})
		return
	}
	msgs := c.MessagesSnapshot()

	needsVision := e.hasImageAttachment(in.User, in.Attachments)
	if needsVision {
		res.members = e.filterVision(res.members)
		if len(res.members) == 0 {
			c.appendDelta(epoch, map[string]any{"error": "Aucun modele vision selectionne : declare un modele vision dans Settings -> Capacites des modeles."})
			return
		}
		msgs = e.expandImageParts(msgs, in.User, in.Attachments)
	}

	if mi, ok := e.memoryIndexMessage(in.User); ok {
		msgs = append([]provider.Message{mi}, msgs...)
	}

	if actx := e.attachmentContext(in.User, in.Attachments); actx != "" {
		msgs = append([]provider.Message{{Role: "system", Content: actx}}, msgs...)
	}

	if res.agent && e.workspace != "" && in.User != "" {
		e.runAgent(ctx, c, epoch, res, msgs, in)
		return
	}

	msgs = append([]provider.Message{{Role: "system", Content: chatSystemPrompt()}}, msgs...)
	// Directive de raisonnement (imperative, en anglais) : le bouton
	// Thinking du composer est l'interrupteur principal en mode chat.
	msgs = append([]provider.Message{{Role: "system", Content: thinkDirective(false, in.Think, in.Effort)}}, msgs...)

	// Choix du moteur de recherche AVANT toute pre-recherche : si le membre
	// principal utilise la recherche native du provider, la pre-recherche
	// locale est inutile (elle doublerait latence et cout). La decision est
	// calculee une seule fois pour ne pas consommer le limiteur deux fois.
	nativePrimary := len(res.members) > 0 && e.useNativeWebSearch(in, res.members[0].Provider, res.members[0].Model)

	// Directive de recherche web (imperative, en anglais) : le globe est
	// l'interrupteur principal, le mode "off" coupe aussi la recherche.
	msgs = append([]provider.Message{{Role: "system", Content: searchDirective(e.webEnabled(in), nativePrimary)}}, msgs...)

	// MAREX.md : lu une seule fois au demarrage de la session ; s'il est
	// rempli, un message discret "MAREX.md chargé" s'affiche dans le chat.
	if mm, ok, notice := e.marexForSession(c.ID); ok {
		msgs = append([]provider.Message{mm}, msgs...)
		if notice {
			c.appendDelta(epoch, map[string]any{"system": "MAREX.md chargé"})
		}
	}

	webDone := false
	nativeOff := false
	if e.webToolsFor(in) && !nativePrimary {
		if wctx := e.webContext(ctx, in.User, in.Text); wctx != "" {
			msgs = append([]provider.Message{{Role: "system", Content: wctx}}, msgs...)
		}
		webDone = true
	}

	var content strings.Builder
	emitted := false
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
		c.appendDelta(epoch, map[string]any{"route": map[string]any{
			"provider": m.Provider, "model": m.Model, "label": m.Label,
			"family": in.Family, "mode": in.Mode,
			"local": res.local, "fallback": res.fallback,
		}})

		// Recherche web native du provider (OpenRouter, DeepSeek) : le natif
		// cherche pendant la generation ; les outils web_search/web_fetch
		// deviennent redondants pour ce membre.
		native := false
		if !nativeOff {
			if i == 0 {
				native = nativePrimary
			} else {
				native = e.useNativeWebSearch(in, m.Provider, m.Model)
			}
		}
		req := provider.Request{
			Model:           m.Model,
			Messages:        msgs,
			Temperature:     0.7,
			MaxTokens:       in.MaxTokens,
			EnableReasoning: in.Think,
			ReasoningEffort: resolveEffort(in.Think, in.Text, in.Effort),
		}
		if native {
			req.Extra = nativeWebExtraFor(m.Provider)
			c.appendDelta(epoch, map[string]any{"search": map[string]any{"phase": "start", "native": true}})
		}
		resp, err := p.Stream(ctx, req, func(ev provider.Event) bool {
			if ev.Reasoning != "" && in.Think {
				c.appendDelta(epoch, map[string]any{"reasoning_content": ev.Reasoning})
			}
			if ev.Content != "" {
				emitted = true
				content.WriteString(ev.Content)
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
		if err == nil {
			if native {
				c.appendDelta(epoch, map[string]any{"search": searchSourcesDelta(resp.Annotations, true)})
			}
			c.appendAssistant(epoch, resp.Content)
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
		// Repli natif -> outils : la recherche native a echoue avant toute
		// emission et le mode "auto" autorise les outils. On ferme l'indicateur
		// natif, on fait une seule pre-recherche locale, puis les membres
		// suivants tournent sans natif (pas de double recherche).
		if native && !webDone && e.webToolsFor(in) {
			c.appendDelta(epoch, map[string]any{"search": searchSourcesDelta(nil, true)})
			if wctx := e.webContext(ctx, in.User, in.Text); wctx != "" {
				msgs = append([]provider.Message{{Role: "system", Content: wctx}}, msgs...)
			}
			webDone = true
			nativeOff = true
			// Le natif est desactive : la directive ne doit plus parler de
			// recherche native pour les membres suivants.
			msgs = replaceWebDirective(msgs, searchDirective(true, true), searchDirective(true, false))
		}
	}
	if lastErr == nil {
		lastErr = fmt.Errorf("aucun modele disponible")
	}
	c.appendDelta(epoch, map[string]any{"error": lastErr.Error()})
}
