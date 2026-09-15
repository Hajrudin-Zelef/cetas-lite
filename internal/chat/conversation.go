package chat

import (
	"context"
	"encoding/json"
	"errors"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"cetas-lite/internal/provider"
)

const maxLogEvents = 200000

var (
	ErrBusy       = errors.New("generation en cours")
	ErrBadMessage = errors.New("message vide")
	ErrNoTurn     = errors.New("aucun tour a regenerer")
)

func newID() string { return strconv.FormatInt(time.Now().UnixNano(), 10) }

type TurnInput struct {
	User        string
	Family      string
	Mode        string
	Text        string
	Web         bool
	MCP         bool
	Think       bool
	Effort      string
	Attachments []string
	// Approve demande une validation utilisateur avant chaque outil
	// d'ecriture/execution (agent uniquement).
	Approve bool
	// AgentMode est le mode top-level choisi dans l'UI ("Chat" ou "Agent").
	// Les outils ne sont actifs que si AgentMode est vrai ET que la
	// famille/mode resolue supporte l'agent.
	AgentMode bool
	// Plan active le mode plan : l'agent explore puis propose un plan
	// a valider avant d'executer (agent uniquement).
	Plan bool
	// Worktree isole le tour dans un worktree git dedie (agent uniquement).
	// Repo est le chemin du depot a cloner en worktree.
	Worktree bool
	Repo     string
	// ProjectID est le projet (upload local ou dossier SFTP) sur lequel
	// l'agent travaille. Vide = espace partagé historique.
	ProjectID string
	// MaxTokens limite les tokens generes par reponse (0 = defaut).
	MaxTokens int
}

// Bornes du reglage "tokens max par reponse".
const (
	MinMaxTokens = 300
	MaxMaxTokens = 32768
)

// ClampMaxTokens ramene n dans [MinMaxTokens, MaxMaxTokens].
// Une valeur <= 0 retourne 0 (defaut du provider).
func ClampMaxTokens(n int) int {
	if n <= 0 {
		return 0
	}
	if n < MinMaxTokens {
		return MinMaxTokens
	}
	if n > MaxMaxTokens {
		return MaxMaxTokens
	}
	return n
}

type Runner interface {
	Run(ctx context.Context, c *Conversation, epoch int, in TurnInput)
}

type Conversation struct {
	mu   sync.Mutex
	cond *sync.Cond

	ID       string
	Messages []provider.Message
	Log      []LogEvent
	Seq      int

	Generating bool
	genStart   time.Time
	cancel     context.CancelFunc
	epoch      int
	lastTurn   *snapshotTurn

	// Approbations en attente, par id de demande.
	approvals map[string]chan approvalDecision

	runner  Runner
	persist func(c *Conversation)
}

func NewConversation(id string, runner Runner, persist func(c *Conversation)) *Conversation {
	c := &Conversation{ID: id, runner: runner, persist: persist}
	c.cond = sync.NewCond(&c.mu)
	return c
}

func (c *Conversation) setRunner(r Runner) {
	c.mu.Lock()
	c.runner = r
	c.mu.Unlock()
}

func (c *Conversation) StartTurn(in TurnInput) error {
	if strings.TrimSpace(in.Text) == "" {
		return ErrBadMessage
	}
	c.mu.Lock()
	if c.Generating {
		c.mu.Unlock()
		return ErrBusy
	}
	c.Generating = true
	c.genStart = time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	c.cancel = cancel
	c.Messages = append(c.Messages, provider.Message{Role: "user", Content: in.Text})
	c.lastTurn = &snapshotTurn{Family: in.Family, Mode: in.Mode, Text: in.Text, Web: in.Web, MCP: in.MCP, Think: in.Think, Effort: in.Effort, Approve: in.Approve, Plan: in.Plan, Worktree: in.Worktree, Repo: in.Repo, ProjectID: in.ProjectID, Attachments: in.Attachments}
	epoch := c.epoch
	runner := c.runner
	c.mu.Unlock()

	c.appendDelta(epoch, map[string]any{"user": in.Text})
	if c.persist != nil {
		c.persist(c)
	}
	if runner == nil {
		c.mu.Lock()
		c.Generating = false
		c.cancel = nil
		c.mu.Unlock()
		return errors.New("moteur indisponible")
	}
	go runner.Run(ctx, c, epoch, in)
	return nil
}

func (c *Conversation) appendDelta(epoch int, delta map[string]any) {
	c.mu.Lock()
	if c.epoch != epoch {
		c.mu.Unlock()
		return
	}
	c.Seq++
	c.Log = append(c.Log, LogEvent{Seq: c.Seq, TS: time.Now().UnixMilli(), Delta: delta})
	if len(c.Log) > maxLogEvents {
		c.Log = c.Log[len(c.Log)-maxLogEvents:]
	}
	c.cond.Broadcast()
	c.mu.Unlock()
}

func (c *Conversation) Stop() {
	c.mu.Lock()
	cancel := c.cancel
	c.mu.Unlock()
	if cancel != nil {
		cancel()
	}
}

func (c *Conversation) Reset() {
	c.Stop()
	c.failPendingApprovals()
	c.mu.Lock()
	c.Messages = nil
	c.Log = nil
	c.Seq = 0
	c.ID = newID()
	c.epoch++
	c.Generating = false
	c.cancel = nil
	c.lastTurn = nil
	c.cond.Broadcast()
	c.mu.Unlock()
	if c.persist != nil {
		c.persist(c)
	}
}

func (c *Conversation) IsGenerating() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Generating
}

func (c *Conversation) State() map[string]any {
	c.mu.Lock()
	defer c.mu.Unlock()
	var elapsed int64
	if c.Generating && !c.genStart.IsZero() {
		elapsed = time.Since(c.genStart).Milliseconds()
	}
	return map[string]any{
		"id":             c.ID,
		"seq":            c.Seq,
		"generating":     c.Generating,
		"gen_elapsed_ms": elapsed,
		"turns":          countTurns(c.Log),
	}
}

func countTurns(log []LogEvent) int {
	n := 0
	for _, ev := range log {
		if _, ok := ev.Delta["user"]; ok {
			n++
		}
	}
	return n
}

func (c *Conversation) MessagesSnapshot() []provider.Message {
	c.mu.Lock()
	defer c.mu.Unlock()
	return append([]provider.Message(nil), c.Messages...)
}

func (c *Conversation) finishTurn(epoch int, elapsed time.Duration) {
	c.mu.Lock()
	if c.epoch == epoch {
		c.cancel = nil
		c.Log = coalesceCompletedTurns(c.Log)
	}
	c.mu.Unlock()
	c.appendDelta(epoch, map[string]any{"turn_done": true, "elapsed_ms": elapsed.Milliseconds()})
	if c.persist != nil {
		c.persist(c)
	}
	c.mu.Lock()
	if c.epoch == epoch {
		c.Generating = false
	}
	c.cond.Broadcast()
	c.mu.Unlock()
}

func (c *Conversation) appendAssistant(epoch int, content string) {
	if strings.TrimSpace(content) == "" {
		return
	}
	c.mu.Lock()
	if c.epoch == epoch {
		c.Messages = append(c.Messages, provider.Message{Role: "assistant", Content: content})
	}
	c.mu.Unlock()
}

func (c *Conversation) marshal() snapshot {
	c.mu.Lock()
	defer c.mu.Unlock()
	return snapshot{
		ID:       c.ID,
		Messages: append([]provider.Message(nil), c.Messages...),
		Log:      append([]LogEvent(nil), c.Log...),
		Seq:      c.Seq,
		Turn:     c.lastTurn,
	}
}

func (c *Conversation) load(s snapshot) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ID = s.ID
	c.Messages = s.Messages
	c.Log = s.Log
	c.Seq = s.Seq
	c.lastTurn = s.Turn
	c.Generating = false
	c.cancel = nil
}

func (c *Conversation) isEmpty() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.Messages) == 0 && len(c.Log) == 0
}

func (c *Conversation) restore(s snapshot) {
	c.Stop()
	c.failPendingApprovals()
	c.mu.Lock()
	c.ID = s.ID
	c.Messages = s.Messages
	c.Log = s.Log
	c.Seq = s.Seq
	c.lastTurn = s.Turn
	c.Generating = false
	c.cancel = nil
	c.epoch++
	c.cond.Broadcast()
	c.mu.Unlock()
	if c.persist != nil {
		c.persist(c)
	}
}

func (c *Conversation) Subscribe(ctx context.Context, from int, emit func(map[string]any) bool) {
	go func() {
		<-ctx.Done()
		c.mu.Lock()
		c.cond.Broadcast()
		c.mu.Unlock()
	}()

	if !emit(map[string]any{"pad": strings.Repeat(".", 2048)}) {
		return
	}

	c.mu.Lock()
	snap := append([]LogEvent(nil), c.Log...)
	epoch := c.epoch
	c.mu.Unlock()

	if n := len(snap); n > 0 && from > snap[n-1].Seq {
		from = 0
	}
	last := from
	for _, ev := range coalesceReplay(snap, from) {
		if ctx.Err() != nil {
			return
		}
		if !emit(ev) {
			return
		}
		if s, ok := ev["seq"].(int); ok {
			last = s
		}
	}
	if !emit(map[string]any{"caught_up": true}) {
		return
	}
	if !emit(map[string]any{"pad": strings.Repeat(".", 16384)}) {
		return
	}

	c.mu.Lock()
	for {
		if ctx.Err() != nil {
			c.mu.Unlock()
			return
		}
		if c.epoch != epoch {
			epoch = c.epoch
			last = 0
			c.mu.Unlock()
			if !emit(map[string]any{"reset": true}) {
				return
			}
			c.mu.Lock()
			continue
		}
		i := sort.Search(len(c.Log), func(i int) bool { return c.Log[i].Seq > last })
		var pending []LogEvent
		if i < len(c.Log) {
			pending = append(pending, c.Log[i:]...)
		}
		if len(pending) == 0 {
			c.cond.Wait()
			continue
		}
		lastEmitted := last
		last = pending[len(pending)-1].Seq
		c.mu.Unlock()
		for _, ev := range pending {
			if !emit(decorateEvent(ev, lastEmitted)) {
				return
			}
		}
		c.mu.Lock()
	}
}

func coalesceCompletedTurns(log []LogEvent) []LogEvent {
	cut := 0
	for i := len(log) - 1; i >= 0; i-- {
		if _, ok := log[i].Delta["user"]; ok {
			cut = i
			break
		}
	}
	if cut <= 0 {
		return log
	}
	return append(coalesceLog(log[:cut]), log[cut:]...)
}

func cloneDelta(d map[string]any) map[string]any {
	out := make(map[string]any, len(d))
	for k, v := range d {
		out[k] = v
	}
	return out
}

func coalesceLog(log []LogEvent) []LogEvent {
	if len(log) == 0 {
		return log
	}
	out := make([]LogEvent, 0, len(log))
	for _, ev := range log {
		ev.Delta = cloneDelta(ev.Delta)
		k := textKey(ev.Delta)
		if k == "" || len(out) == 0 {
			out = append(out, ev)
			continue
		}
		prev := &out[len(out)-1]
		if textKey(prev.Delta) != k {
			out = append(out, ev)
			continue
		}
		prev.Delta[k] = prev.Delta[k].(string) + ev.Delta[k].(string)
		prev.Seq = ev.Seq
		prev.TS = ev.TS
		if ptok, ok := prev.Delta["toks"]; ok {
			if etok, ok := ev.Delta["toks"]; ok {
				prev.Delta["toks"] = ptok.(int) + etok.(int)
			}
		}
	}
	return out
}

func textKey(d map[string]any) string {
	if _, ok := d["content"].(string); ok {
		return "content"
	}
	if _, ok := d["reasoning_content"].(string); ok {
		return "reasoning_content"
	}
	return ""
}

func coalesceReplay(events []LogEvent, from int) []map[string]any {
	var out []map[string]any
	var buf strings.Builder
	key := ""
	var seq0, seq, toks int
	var ts0, ts int64
	flush := func() {
		if key == "" {
			return
		}
		m := map[string]any{key: buf.String(), "seq": seq, "ts": ts, "ts0": ts0, "seq0": seq0, "toks": toks}
		if seq0 <= from {
			m["replace"] = true
		}
		out = append(out, m)
		buf.Reset()
		key = ""
		toks = 0
	}
	for _, ev := range events {
		if ev.Seq <= from {
			continue
		}
		k := textKey(ev.Delta)
		if k != "" {
			if key != "" && key != k {
				flush()
			}
			if key == "" {
				key, ts0, seq0 = k, ev.TS, ev.Seq
			}
			buf.WriteString(ev.Delta[k].(string))
			seq, ts = ev.Seq, ev.TS
			toks++
			continue
		}
		flush()
		m := map[string]any{"seq": ev.Seq, "ts": ev.TS}
		for kk, vv := range ev.Delta {
			m[kk] = vv
		}
		out = append(out, m)
	}
	flush()
	return out
}

func decorateEvent(ev LogEvent, from int) map[string]any {
	m := map[string]any{"seq": ev.Seq, "ts": ev.TS}
	for k, v := range ev.Delta {
		m[k] = v
	}
	if textKey(ev.Delta) != "" {
		if seq0, ok := ev.Delta["seq0"]; ok {
			if s, ok := seq0.(int); ok && s <= from {
				m["replace"] = true
			}
		}
	}
	return m
}

func (c *Conversation) save() ([]byte, error) {
	return json.Marshal(c.marshal())
}

// ApprovalRequest decrit une demande de validation utilisateur : soit un
// outil sensible avant execution (kind "tool"), soit un plan a valider
// avant la phase d'execution du mode plan (kind "plan").
type ApprovalRequest struct {
	ID   string         `json:"id"`
	Kind string         `json:"kind"`
	Tool string         `json:"tool,omitempty"`
	Args map[string]any `json:"args,omitempty"`
	Plan string         `json:"plan,omitempty"`
}

type approvalDecision struct {
	approved bool
	always   bool
}

const approvalTimeout = 10 * time.Minute

// RequestApproval emet une demande d'approbation et bloque jusqu'a la decision
// de l'utilisateur, l'arret du tour ou l'expiration du delai.
func (c *Conversation) RequestApproval(ctx context.Context, epoch int, req ApprovalRequest) (approvalDecision, error) {
	if req.ID == "" {
		req.ID = newID()
	}
	ch := make(chan approvalDecision, 1)
	c.mu.Lock()
	if c.approvals == nil {
		c.approvals = map[string]chan approvalDecision{}
	}
	c.approvals[req.ID] = ch
	c.mu.Unlock()
	defer func() {
		c.mu.Lock()
		delete(c.approvals, req.ID)
		c.mu.Unlock()
	}()

	c.appendDelta(epoch, map[string]any{"approval": map[string]any{
		"id": req.ID, "kind": req.Kind, "tool": req.Tool, "args": req.Args,
		"plan": req.Plan, "phase": "request",
	}})

	timer := time.NewTimer(approvalTimeout)
	defer timer.Stop()
	select {
	case d := <-ch:
		c.appendDelta(epoch, map[string]any{"approval": map[string]any{
			"id": req.ID, "phase": "resolved", "approved": d.approved,
		}})
		return d, nil
	case <-ctx.Done():
		return approvalDecision{}, ctx.Err()
	case <-timer.C:
		c.appendDelta(epoch, map[string]any{"approval": map[string]any{
			"id": req.ID, "phase": "resolved", "approved": false, "timeout": true,
		}})
		return approvalDecision{}, errors.New("approbation expiree")
	}
}

// ResolveApproval transmet la decision de l'utilisateur a une demande en attente.
func (c *Conversation) ResolveApproval(id string, approved, always bool) bool {
	c.mu.Lock()
	ch, ok := c.approvals[id]
	c.mu.Unlock()
	if !ok {
		return false
	}
	select {
	case ch <- approvalDecision{approved: approved, always: always}:
		return true
	default:
		return false
	}
}

// failPendingApprovals refuse les demandes en attente (reset / restauration).
func (c *Conversation) failPendingApprovals() {
	c.mu.Lock()
	for id, ch := range c.approvals {
		select {
		case ch <- approvalDecision{}:
		default:
		}
		delete(c.approvals, id)
	}
	c.mu.Unlock()
}
