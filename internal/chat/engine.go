package chat

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/attach"
	"cetas-lite/internal/cache"
	"cetas-lite/internal/local"
	"cetas-lite/internal/mcp"
	"cetas-lite/internal/modelcaps"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/rag"
	"cetas-lite/internal/store"
	"cetas-lite/internal/workspace"
	"cetas-lite/internal/worktree"

	"golang.org/x/time/rate"
)

type Engine struct {
	reg         *provider.Registry
	discover    LocalDiscoverer
	st          *store.Store
	workspace   string
	allowScript bool
	searcher    WebTools
	mem         MemoryTools
	rag         RagTools
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

	// Pool de questions suggérées (itération 4 : diversion d'accueil).
	// Chargé au démarrage depuis suggested-questions.yaml, servi tel quel
	// au navigateur qui effectue le tirage.
	suggestions []Suggestion

	// Cache exact des réponses provider (itération 5a) : actif uniquement
	// pour les requêtes déterministes (température 0, sans outils, sans
	// paramètres supplémentaires). Nil = désactivé.
	exactCache_ *cache.Cache

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

	// Pré-génération des questions suggérées (lot 5) : jobs en cours par
	// utilisateur et dépôt de réponses (valeur zéro utilisable).
	pregen pregenState
}

// LocalDiscoverer fournit les modeles decouverts sur les moteurs locaux.
// *local.Discoverer l'implemente ; les tests injectent un faux.
type LocalDiscoverer interface {
	ModelsForEngine(ctx context.Context, engine string) []local.Model
}

func NewEngine(reg *provider.Registry, families []alias.Family, st *store.Store, disc LocalDiscoverer, workspace string) *Engine {
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

// LocalModels retourne les modeles decouverts pour un moteur local
// (llamacpp, ollama, lmstudio). Utilise par le selecteur de modeles.
func (e *Engine) LocalModels(ctx context.Context, engine string) []string {
	if e.discover == nil {
		return nil
	}
	var out []string
	for _, m := range e.discover.ModelsForEngine(ctx, engine) {
		out = append(out, m.ID)
	}
	return out
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

func (e *Engine) SetRAG(r RagTools) {
	e.mu.Lock()
	e.rag = r
	e.mu.Unlock()
}

func (e *Engine) ragTools() RagTools {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.rag
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
	return e.conversationLocked(user)
}

// conversationLocked retourne la conversation de la session courante de
// l'utilisateur, en la restaurant depuis le store si besoin.
// e.mu doit être verrouillé par l'appelant.
func (e *Engine) conversationLocked(user string) *Conversation {
	if c, ok := e.convs[user]; ok {
		return c
	}
	e.ensureMigratedLocked(user)
	c := NewConversation(newID(), e, func(c *Conversation) { e.saveSession(user, c) })
	if id := e.currentIDLocked(user); id != "" {
		if rec, ok := e.getSessionLocked(user, id); ok {
			c.load(rec.Snapshot)
		} else {
			// Session courante sans enregistrement (ex. vierge) : on garde
			// son ID pour que _current reste stable.
			c.ID = id
		}
	} else {
		e.setCurrentLocked(user, c.ID)
	}
	e.convs[user] = c
	return c
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
	in := TurnInput{User: user, Family: last.Family, Mode: last.Mode, Text: last.Text, Web: last.Web, WebDepth: last.WebDepth, MCP: last.MCP, Think: last.Think, Effort: last.Effort, Approve: last.Approve, Plan: last.Plan, Worktree: last.Worktree, Repo: last.Repo, ProjectID: last.ProjectID, Attachments: last.Attachments, FocusCorpus: last.FocusCorpus}
	c.mu.Unlock()
	if c.persist != nil {
		c.persist(c)
	}
	return c.StartTurn(in)
}

type resolution struct {
	members  []alias.ResolvedMember
	local    bool
	fallback bool
	agent    bool
	tier     string // mode "auto" : tier choisi par le routage par effort
}

func (e *Engine) resolve(ctx context.Context, in TurnInput) resolution {
	fam, ok := alias.Find(e.Families(), in.Family)
	if !ok {
		return resolution{}
	}
	// Mode "auto" (menu +) : union des pools effectifs de tous les modes de
	// la famille, dedupliquee, tirage aleatoire a chaque requete. Choix
	// rapide "tout le fallback de l'alias", sans passer par le selecteur.
	if in.Mode == "auto" {
		return e.resolveAuto(ctx, fam, in)
	}
	if fam.Local {
		for _, m := range fam.Modes {
			if m.ID != in.Mode {
				continue
			}
			// Selection locale (selecteur de modeles) : si le mode a un pool
			// configure, on ne garde que les modeles decouverts selectionnes.
			// Pool vide = tous les modeles decouverts (comportement historique).
			var selected map[string]bool
			if len(m.Pool) > 0 {
				selected = make(map[string]bool, len(m.Pool))
				for _, p := range m.Pool {
					selected[p.Model] = true
				}
			}
			if e.discover != nil {
				if models := e.discover.ModelsForEngine(ctx, m.Engine); len(models) > 0 {
					var members []alias.ResolvedMember
					for _, mod := range models {
						if selected != nil && !selected[mod.ID] {
							continue
						}
						members = append(members, alias.ResolvedMember{Provider: m.Engine, Model: mod.ID, Label: mod.ID})
					}
					if len(members) > 0 {
						if len(members) > 1 {
							members = alias.ShufflePool(members)
						}
						return resolution{members: members, local: true}
					}
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
	if len(members) > 1 {
		// Fallback : tirage aleatoire du premier modele a chaque requete,
		// puis parcours sequentiel du reste en cas d'echec. Un pool d'un
		// seul membre = modele fixe, pas de fallback.
		members = alias.ShufflePool(members)
	}
	if len(rm.Fallback) > 0 {
		// Repli apres epuisement du pool primaire : ajoute dans l'ordre
		// declare, sans melange. Chemin agent uniquement (le chat general
		// n'utilise pas ce champ).
		combined := make([]alias.ResolvedMember, 0, len(members)+len(rm.Fallback))
		combined = append(combined, members...)
		combined = append(combined, rm.Fallback...)
		members = combined
	}
	return resolution{members: members, agent: rm.Agent && in.AgentMode}
}

// resolveAuto construit la resolution du mode "auto" (menu + : « Défaut »).
// Pour les familles cloud : routage par effort — le tier (mode) est choisi
// via autoTierFor (jamais le plus cher sans effort "high" explicite), puis
// le pool effectif du tier est mélangé (fallback intra-tier préservé :
// tirage aléatoire + repli séquentiel en cas d'échec). Si le tier choisi est
// vide (mauvaise configuration), repli sur l'union dédupliquée des pools
// effectifs (comportement historique). Pour les familles locales, l'union
// des modeles decouverts (selection du selecteur honoree).
func (e *Engine) resolveAuto(ctx context.Context, fam alias.Family, in TurnInput) resolution {
	if fam.Local {
		seen := make(map[string]bool)
		var members []alias.ResolvedMember
		for _, m := range fam.Modes {
			var selected map[string]bool
			if len(m.Pool) > 0 {
				selected = make(map[string]bool, len(m.Pool))
				for _, p := range m.Pool {
					selected[p.Model] = true
				}
			}
			if e.discover == nil {
				continue
			}
			for _, mod := range e.discover.ModelsForEngine(ctx, m.Engine) {
				if selected != nil && !selected[mod.ID] {
					continue
				}
				k := m.Engine + "/" + mod.ID
				if seen[k] {
					continue
				}
				seen[k] = true
				members = append(members, alias.ResolvedMember{Provider: m.Engine, Model: mod.ID, Label: mod.ID})
			}
		}
		if len(members) == 0 {
			fb, ok := alias.Resolve(e.Families(), "samagent-n4", "standard")
			if ok {
				return resolution{members: fb.Pool, fallback: true}
			}
			return resolution{}
		}
		if len(members) > 1 {
			members = alias.ShufflePool(members)
		}
		return resolution{members: members, local: true, fallback: len(members) > 1}
	}
	// Routage par effort : un seul tier, choisi selon l'effort demandé.
	if tier := autoTierFor(fam.Modes, in.Effort, in.Text); tier != "" {
		if rm, ok := alias.Resolve(e.Families(), fam.ID, tier); ok && len(rm.Pool) > 0 {
			members := rm.Pool
			if len(members) > 1 {
				members = alias.ShufflePool(members)
			}
			return resolution{members: members, fallback: len(members) > 1, tier: tier}
		}
	}
	// Repli : union dédupliquée des pools effectifs (tier vide ou inconnu).
	seen := make(map[string]bool)
	var members []alias.ResolvedMember
	for _, m := range fam.Modes {
		rm, ok := alias.Resolve(e.Families(), fam.ID, m.ID)
		if !ok {
			continue
		}
		for _, p := range rm.Pool {
			k := p.Provider + "/" + p.Model
			if seen[k] {
				continue
			}
			seen[k] = true
			members = append(members, p)
		}
	}
	if len(members) > 1 {
		members = alias.ShufflePool(members)
	}
	return resolution{members: members, fallback: len(members) > 1}
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

	// Historique conversationnel brut (iteration 6) : copie avant les
	// inserts systeme (memory index, pieces jointes) pour que
	// l'enrichissement de requete RAG travaille sur la conversation seule.
	ragHistory := append([]provider.Message(nil), msgs...)

	// Effort de raisonnement résolu une fois pour le tour : pilotage du
	// payload provider, de la directive system et de l'affichage (badge).
	reasoningEffort := resolveEffort(in.Think, in.Text, in.Effort)

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

	// Base documentaire locale : extraits pertinents injectes avant l'appel
	// modele (fail-open : absente ou vide, rien n'est ajoute). Le resultat
	// est calcule une seule fois et sert aussi a economiser la pre-recherche
	// web quand la base couvre la requete.
	// Placement cache-friendly : le contexte RAG (dynamique, propre a la
	// requete) est insere juste avant le message utilisateur courant, apres
	// l'historique. Le prefixe [prompts stables + historique] reste ainsi
	// identique d'un tour a l'autre, ce qui maximise les chances de prompt
	// caching (prefixe automatique cote provider ; aucun marquage de cache
	// explicite n'est emis par ce code).
	//
	// Focus corpus (clic sur une question suggérée) : recherche restreinte
	// au corpus, RAG forcément sollicité (sans porte de score, la question
	// étant curée). Fail-open : aucun hit => rien d'injecté, et la note
	// « non couvert » ne s'applique pas (choix explicite).
	focus := strings.TrimSpace(in.FocusCorpus) != ""
	var ragRes rag.Result
	if focus {
		ragRes = e.ragHitsCorpus(ctx, in.Text, in.FocusCorpus)
		if rc, ok := ragContextForced(ragRes); ok {
			msgs = insertBeforeLastUser(msgs, provider.Message{Role: "system", Content: rc})
		}
	} else {
		ragRes = e.ragHits(ctx, in.Text, ragHistory)
		if rc, ok := ragContextFrom(ragRes); ok {
			msgs = insertBeforeLastUser(msgs, provider.Message{Role: "system", Content: rc})
		}
		// Base locale active mais requete non couverte (iteration 2) : le dire
		// immediatement et interdire de presenter une invention comme issue du
		// corpus. Note stable (cache-friendly), inseree comme le contexte RAG
		// pour couvrir aussi le mode agent.
		if ragRes.Ready && !ragCovered(ragRes) {
			msgs = insertBeforeLastUser(msgs, provider.Message{Role: "system", Content: ragNotCoveredNote()})
		}
	}

	// Couverture pour la suite : en focus, « couvert » = au moins un hit
	// dans le corpus (pas de porte de score sur un choix explicite).
	ragCoveredNow := ragCovered(ragRes)
	if focus {
		ragCoveredNow = len(ragRes.Hits) > 0
	}

	if res.agent && e.workspace != "" && in.User != "" {
		e.runAgent(ctx, c, epoch, res, msgs, in)
		return
	}

	msgs = append([]provider.Message{{Role: "system", Content: chatSystemPrompt()}}, msgs...)
	// Directive de raisonnement (imperative, en anglais) : le bouton
	// Thinking du composer est l'interrupteur principal en mode chat.
	msgs = append([]provider.Message{{Role: "system", Content: thinkDirective(false, in.Think, reasoningEffort)}}, msgs...)

	// Choix du moteur de recherche AVANT toute pre-recherche : si le membre
	// principal utilise la recherche native du provider, la pre-recherche
	// locale est inutile (elle doublerait latence et cout). La decision est
	// calculee une seule fois pour ne pas consommer le limiteur deux fois.
	nativePrimary := len(res.members) > 0 && e.useNativeWebSearch(in, res.members[0].Provider, res.members[0].Model)

	// Directive de recherche web (imperative, en anglais) : le globe est
	// l'interrupteur principal, le mode "off" coupe aussi la recherche.
	webOn := e.webEnabled(in)
	msgs = append([]provider.Message{{Role: "system", Content: searchDirective(webOn, nativePrimary)}}, msgs...)
	// La base locale couvre la requete : on remplace la consigne "cherche
	// sur le web" par "reponds d'abord depuis la base locale".
	if webOn && ragCoveredNow {
		msgs = replaceWebDirective(msgs, searchDirective(true, nativePrimary), localFirstDirective())
	}

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
		// Base locale pertinente : on economise la pre-recherche web
		// (jusqu'a 8 s + des tokens) et on repond depuis les extraits.
		if !ragCoveredNow {
			if wctx := e.webContext(ctx, in.User, in.Text); wctx != "" {
				msgs = append([]provider.Message{{Role: "system", Content: wctx}}, msgs...)
			}
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
			"effort": reasoningEffort, "tier": res.tier,
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
			ReasoningEffort: reasoningEffort,
		}
		if native {
			req.Extra = nativeWebExtraFor(m.Provider)
			c.appendDelta(epoch, map[string]any{"search": map[string]any{"phase": "start", "native": true}})
		}
		// Cache exact (itération 5a) : requête déterministe (température 0,
		// sans outils, sans recherche native) => on rejoue la réponse en
		// cache au lieu d'appeler le provider. Clé vide = non éligible.
		cc := e.exactCache()
		cacheKey := ""
		if cc != nil {
			cacheKey = e.cacheKeyFor(m.Provider, req)
			if cacheKey != "" {
				if ce, ok := cc.Get(cacheKey); ok {
					e.replayCached(c, epoch, in, ce)
					return
				}
			}
		}
		// Pré-génération (lot 5) : si un fantôme a préparé la réponse de
		// ce tour exact, on la rejoue sans appel provider (~1 ms). Un
		// fantôme en cours est attendu (borné) plutôt que doublé.
		if e.pregenServe(ctx, c, epoch, in, m.Provider, req) {
			return
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
			e.cacheStore(cacheKey, resp)
			// Tour fantôme (lot 5) : la réponse réussie alimente le dépôt
			// de pré-génération (le clic la rejouera sans appel).
			e.pregenCapture(in, m.Provider, req, resp)
			c.appendAssistant(epoch, resp.Content, "")
			return
		}
		lastErr = err
		if ctx.Err() != nil {
			c.appendDelta(epoch, map[string]any{"content": "\n\n_Génération interrompue._"})
			return
		}
		if emitted {
			c.appendDelta(epoch, map[string]any{"error": clientSafeError(err, "La génération a été interrompue par une erreur — réessaie dans un instant.")})
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
	c.appendDelta(epoch, map[string]any{"error": clientSafeError(lastErr, "Échec de la génération — réessaie dans un instant.")})
}
