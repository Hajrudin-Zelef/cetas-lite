package chat

import (
	"context"
	"strings"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/provider"
)

func TestMaybeCompactReduces(t *testing.T) {
	fp := &fakeProvider{id: "fake", content: map[string]string{
		"ok": "Resume dense et suffisamment long pour depasser le seuil minimal requis par la compaction.",
	}}
	fams := codeFamily(alias.Member{Provider: "fake", Model: "ok"})
	e := newEngine(t, fp, fams)
	c := e.Conversation("sam")

	big := strings.Repeat("x", 3000)
	c.mu.Lock()
	c.Messages = []provider.Message{{Role: "user", Content: "debut"}}
	for i := 0; i < 40; i++ {
		c.Messages = append(c.Messages, provider.Message{Role: "user", Content: big})
		c.Messages = append(c.Messages, provider.Message{Role: "assistant", Content: big})
	}
	before := len(c.Messages)
	epoch := c.epoch
	c.mu.Unlock()

	rm, ok := alias.Resolve(fams, "code", "standard")
	if !ok {
		t.Fatal("resolution alias")
	}
	if !e.maybeCompact(context.Background(), c, epoch, rm.Pool) {
		t.Fatal("maybeCompact doit signaler la compaction")
	}

	after := len(c.MessagesSnapshot())
	if after >= before {
		t.Fatalf("compaction attendue: %d -> %d", before, after)
	}
	found := false
	for _, m := range c.MessagesSnapshot() {
		if s, ok := m.Content.(string); ok && strings.Contains(s, "[CONTEXT COMPACTED]") {
			found = true
		}
	}
	if !found {
		t.Fatal("marqueur de compaction absent")
	}
}

func TestMaybeCompactNoopBelowThreshold(t *testing.T) {
	fp := &fakeProvider{id: "fake", content: map[string]string{"ok": "resume"}}
	fams := codeFamily(alias.Member{Provider: "fake", Model: "ok"})
	e := newEngine(t, fp, fams)
	c := e.Conversation("sam")

	c.mu.Lock()
	c.Messages = []provider.Message{{Role: "user", Content: "court"}}
	before := len(c.Messages)
	epoch := c.epoch
	c.mu.Unlock()

	rm, _ := alias.Resolve(fams, "code", "standard")
	if e.maybeCompact(context.Background(), c, epoch, rm.Pool) {
		t.Fatal("maybeCompact ne doit rien signaler sous le seuil")
	}

	if len(c.MessagesSnapshot()) != before {
		t.Fatal("aucune compaction ne doit avoir lieu sous le seuil")
	}
}

func TestCoalesceCompletedTurns(t *testing.T) {
	log := []LogEvent{
		{Seq: 1, Delta: map[string]any{"user": "q1"}},
		{Seq: 2, Delta: map[string]any{"content": "bon"}},
		{Seq: 3, Delta: map[string]any{"content": "jour"}},
		{Seq: 4, Delta: map[string]any{"turn_done": true}},
		{Seq: 5, Delta: map[string]any{"user": "q2"}},
		{Seq: 6, Delta: map[string]any{"content": "sa"}},
		{Seq: 7, Delta: map[string]any{"content": "lut"}},
		{Seq: 8, Delta: map[string]any{"turn_done": true}},
	}
	out := coalesceCompletedTurns(log)

	merged := false
	contentCount := 0
	for _, ev := range out {
		if s, ok := ev.Delta["content"].(string); ok {
			contentCount++
			if s == "bonjour" {
				merged = true
			}
		}
	}
	if !merged {
		t.Fatalf("le tour termine doit etre fusionne: %#v", out)
	}
	if contentCount != 3 {
		t.Fatalf("contenus = %d, want 3 (1 fusionne + 2 courants)", contentCount)
	}
	if s, _ := log[1].Delta["content"].(string); s != "bon" {
		t.Fatalf("le log d'origine ne doit pas etre mute: %q", s)
	}
}
