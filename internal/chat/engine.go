package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/local"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/store"

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

	mu         sync.Mutex
	families   []alias.Family
	convs      map[string]*Conversation
	webLimiter map[string]*rate.Limiter
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

func (e *Engine) Conversation(user string) *Conversation {
	e.mu.Lock()
	defer e.mu.Unlock()
	if c, ok := e.convs[user]; ok {
		return c
	}
	c := NewConversation(newID(), e, func(c *Conversation) { e.save(user, c) })
	if raw, ok := e.st.GetConversation(user, "active"); ok {
		var s snapshot
		if err := json.Unmarshal(raw, &s); err == nil && s.ID != "" {
			c.load(s)
		}
	}
	e.convs[user] = c
	return c
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
	return resolution{members: rm.Pool, agent: rm.Agent}
}

func (e *Engine) Run(ctx context.Context, c *Conversation, epoch int, in TurnInput) {
	start := time.Now()
	defer func() { c.finishTurn(epoch, time.Since(start)) }()

	res := e.resolve(ctx, in)
	if len(res.members) == 0 {
		c.appendDelta(epoch, map[string]any{"error": "aucun modele disponible pour cet alias"})
		return
	}
	msgs := c.MessagesSnapshot()

	if mi, ok := e.memoryIndexMessage(in.User); ok {
		msgs = append([]provider.Message{mi}, msgs...)
	}

	if res.agent && e.workspace != "" && in.User != "" {
		e.runAgent(ctx, c, epoch, res, msgs, in)
		return
	}

	if in.Web && in.User != "" {
		if wctx := e.webContext(ctx, in.User, in.Text); wctx != "" {
			msgs = append([]provider.Message{{Role: "system", Content: wctx}}, msgs...)
		}
	}

	var content strings.Builder
	emitted := false
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
		c.appendDelta(epoch, map[string]any{"route": map[string]any{
			"provider": m.Provider, "model": m.Model, "label": m.Label,
			"family": in.Family, "mode": in.Mode,
			"local": res.local, "fallback": res.fallback,
		}})

		resp, err := p.Stream(ctx, provider.Request{
			Model:       m.Model,
			Messages:    msgs,
			Temperature: 0.7,
		}, func(ev provider.Event) bool {
			if ev.Reasoning != "" {
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
	}
	if lastErr == nil {
		lastErr = fmt.Errorf("aucun modele disponible")
	}
	c.appendDelta(epoch, map[string]any{"error": lastErr.Error()})
}
