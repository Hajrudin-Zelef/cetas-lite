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
	e.ArchiveAndReset("sam")

	archives := e.ListArchives("sam")
	if len(archives) != 1 {
		t.Fatalf("archives = %v", archives)
	}
	a := archives[0]
	if a.Title != "question initiale" {
		t.Fatalf("titre = %q", a.Title)
	}
	if a.Updated == 0 {
		t.Fatal("updated attendu")
	}
	if a.Messages != 2 {
		t.Fatalf("messages = %d", a.Messages)
	}

	if e.DeleteArchive("sam", "inconnu") {
		t.Fatal("delete d'une archive inconnue doit echouer")
	}
	if !e.DeleteArchive("sam", a.ID) {
		t.Fatal("delete attendu")
	}
	if got := e.ListArchives("sam"); len(got) != 0 {
		t.Fatalf("archives restantes = %v", got)
	}
}

func TestArchiveIsolation(t *testing.T) {
	fp := &fakeProvider{id: "fake", content: map[string]string{"ok": "bonjour"}}
	e := newEngine(t, fp, codeFamily(alias.Member{Provider: "fake", Model: "ok"}))
	c := e.Conversation("sam")
	if err := c.StartTurn(TurnInput{User: "sam", Family: "code", Mode: "standard", Text: "secret"}); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return !c.IsGenerating() }, "tour")
	e.ArchiveAndReset("sam")

	if got := e.ListArchives("alice"); len(got) != 0 {
		t.Fatalf("alice ne doit rien voir: %v", got)
	}
	if e.DeleteArchive("alice", e.ListArchives("sam")[0].ID) {
		t.Fatal("alice ne doit pas pouvoir supprimer l'archive de sam")
	}
}
