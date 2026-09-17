package chat

import (
	"testing"

	"cetas-lite/internal/alias"
)

func TestArchiveInfoAndDelete(t *testing.T) {
	fp := &fakeProvider{id: "fake", content: map[string]string{"ok": "bonjour"}}
	e := newEngine(t, fp, codeFamily(alias.Member{Provider: "fake", Model: "ok"}))
	c := e.Conversation("sam")
	if err := c.StartTurn(TurnInput{User: "sam", Family: "code", Mode: "standard", Text: "question initiale"}); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return !c.IsGenerating() }, "tour")

	archives := e.ListArchives("sam")
	if len(archives) != 0 {
		t.Fatalf("archives initiales = %v, want 0", archives)
	}

	if e.DeleteArchive("sam", "inconnu") {
		t.Fatal("delete d'une archive inconnue doit echouer")
	}
}
