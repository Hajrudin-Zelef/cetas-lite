package chat

import (
	"errors"
	"testing"

	"cetas-lite/internal/alias"
)

func TestRegenerate(t *testing.T) {
	fp := &fakeProvider{id: "fake", content: map[string]string{"ok": "bonjour"}}
	e := newEngine(t, fp, codeFamily(alias.Member{Provider: "fake", Model: "ok"}))
	c := e.Conversation("sam")
	if err := c.StartTurn(TurnInput{User: "sam", Family: "code", Mode: "standard", Text: "salut"}); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return !c.IsGenerating() }, "tour")
	first := c.marshal()

	if err := e.Regenerate("sam"); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return !c.IsGenerating() }, "regeneration")

	msgs := c.MessagesSnapshot()
	users, assistants := 0, 0
	for _, m := range msgs {
		switch m.Role {
		case "user":
			users++
		case "assistant":
			assistants++
		}
	}
	if users != 1 || assistants != 1 {
		t.Fatalf("apres regeneration: users=%d assistants=%d msgs=%d", users, assistants, len(msgs))
	}
	if len(c.marshal().Log) == 0 {
		t.Fatal("log vide apres regeneration")
	}
	userEvents := 0
	for _, ev := range c.marshal().Log {
		if _, ok := ev.Delta["user"]; ok {
			userEvents++
		}
	}
	if userEvents != 1 {
		t.Fatalf("evenements user = %d, want 1", userEvents)
	}
	if first.Turn == nil {
		t.Fatal("le tour courant doit etre persiste pour regenerer")
	}
}

func TestRegenerateNoTurn(t *testing.T) {
	e := newEngine(t, &fakeProvider{id: "fake", content: map[string]string{"ok": "ok"}}, codeFamily(alias.Member{Provider: "fake", Model: "ok"}))
	if err := e.Regenerate("sam"); !errors.Is(err, ErrNoTurn) {
		t.Fatalf("err = %v, want ErrNoTurn", err)
	}
}
