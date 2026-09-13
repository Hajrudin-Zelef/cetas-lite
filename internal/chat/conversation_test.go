package chat

import (
	"context"
	"sync"
	"testing"
	"time"
)

type fakeRunner struct {
	started chan TurnInput
	block   chan struct{}
}

func (f *fakeRunner) Run(ctx context.Context, c *Conversation, epoch int, in TurnInput) {
	if f.started != nil {
		f.started <- in
	}
	if f.block != nil {
		select {
		case <-f.block:
		case <-ctx.Done():
		}
	}
}

func waitFor(t *testing.T, cond func() bool, msg string) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if cond() {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("timeout: " + msg)
}

func TestStartTurnBusy(t *testing.T) {
	r := &fakeRunner{started: make(chan TurnInput, 1), block: make(chan struct{})}
	c := NewConversation("1", r, nil)
	if err := c.StartTurn(TurnInput{Family: "x", Mode: "y", Text: "salut"}); err != nil {
		t.Fatal(err)
	}
	<-r.started
	if err := c.StartTurn(TurnInput{Family: "x", Mode: "y", Text: "encore"}); err != ErrBusy {
		t.Fatalf("second tour = %v, want ErrBusy", err)
	}
	close(r.block)
}

func TestResetUnblocks(t *testing.T) {
	r := &fakeRunner{started: make(chan TurnInput, 1), block: make(chan struct{})}
	c := NewConversation("1", r, nil)
	_ = c.StartTurn(TurnInput{Text: "salut"})
	<-r.started
	if !c.IsGenerating() {
		t.Fatal("devrait generer")
	}
	c.Reset()
	if c.IsGenerating() {
		t.Fatal("Reset doit rendre la main")
	}
	r2 := &fakeRunner{started: make(chan TurnInput, 1)}
	c.setRunner(r2)
	if err := c.StartTurn(TurnInput{Text: "apres reset"}); err != nil {
		t.Fatalf("apres Reset, un tour doit pouvoir demarrer: %v", err)
	}
	<-r2.started
}

func TestStaleTurnDoesNotRelease(t *testing.T) {
	c := NewConversation("1", &fakeRunner{}, nil)
	oldEpoch := c.epoch

	c.Reset() // ancien epoch devient perime

	c.mu.Lock()
	c.Generating = true // un nouveau tour demarre
	c.mu.Unlock()

	c.finishTurn(oldEpoch, 0) // fin tardive du tour perime

	if !c.IsGenerating() {
		t.Fatal("un tour perime a libere le tour courant")
	}
}

func TestResetEmitsResetOnSubscribe(t *testing.T) {
	c := NewConversation("1", &fakeRunner{}, nil)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		mu     sync.Mutex
		events []map[string]any
	)
	go c.Subscribe(ctx, 0, func(ev map[string]any) bool {
		mu.Lock()
		events = append(events, ev)
		mu.Unlock()
		return true
	})

	waitFor(t, func() bool {
		mu.Lock()
		defer mu.Unlock()
		for _, ev := range events {
			if ev["caught_up"] == true {
				return true
			}
		}
		return false
	}, "caught_up non recu")

	c.Reset()

	waitFor(t, func() bool {
		mu.Lock()
		defer mu.Unlock()
		for _, ev := range events {
			if ev["reset"] == true {
				return true
			}
		}
		return false
	}, "reset non recu")
}

func TestReplayCoalescesAndFollows(t *testing.T) {
	c := NewConversation("1", &fakeRunner{}, nil)
	c.appendDelta(0, map[string]any{"content": "Bon"})
	c.appendDelta(0, map[string]any{"content": "jour"})
	c.appendDelta(0, map[string]any{"user": "salut"})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		mu     sync.Mutex
		events []map[string]any
	)
	go c.Subscribe(ctx, 0, func(ev map[string]any) bool {
		mu.Lock()
		events = append(events, ev)
		mu.Unlock()
		return true
	})

	waitFor(t, func() bool {
		mu.Lock()
		defer mu.Unlock()
		for _, ev := range events {
			if ev["caught_up"] == true {
				return true
			}
		}
		return false
	}, "caught_up non recu")

	mu.Lock()
	var coalesced string
	for _, ev := range events {
		if s, ok := ev["content"].(string); ok && ev["toks"] != nil {
			coalesced = s
		}
	}
	mu.Unlock()
	if coalesced != "Bonjour" {
		t.Fatalf("replay coalesce = %q, want Bonjour", coalesced)
	}

	c.appendDelta(0, map[string]any{"content": "!"})
	waitFor(t, func() bool {
		mu.Lock()
		defer mu.Unlock()
		for _, ev := range events {
			if ev["content"] == "!" {
				return true
			}
		}
		return false
	}, "direct non recu")
}
