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
)

func newID() string { return strconv.FormatInt(time.Now().UnixNano(), 10) }

type TurnInput struct {
	Family string
	Mode   string
	Text   string
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
	c.mu.Lock()
	c.Messages = nil
	c.Log = nil
	c.Seq = 0
	c.ID = newID()
	c.epoch++
	c.Generating = false
	c.cancel = nil
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
		c.Generating = false
		c.cancel = nil
	}
	c.mu.Unlock()
	c.appendDelta(epoch, map[string]any{"turn_done": true, "elapsed_ms": elapsed.Milliseconds()})
	if c.persist != nil {
		c.persist(c)
	}
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
	}
}

func (c *Conversation) load(s snapshot) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ID = s.ID
	c.Messages = s.Messages
	c.Log = s.Log
	c.Seq = s.Seq
	c.Generating = false
	c.cancel = nil
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
