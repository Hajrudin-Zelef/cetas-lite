package chat

import (
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/attach"
)

func TestAttachmentContextInjected(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	st := attach.New(t.TempDir(), 1<<20)
	e.SetAttachments(st)

	doc, err := st.Save("sam", "cours.md", []byte("Le theoreme de Pythagore."))
	if err != nil {
		t.Fatal(err)
	}
	img, err := st.Save("sam", "photo.png", []byte("\x89PNG"))
	if err != nil {
		t.Fatal(err)
	}

	runAgentTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "explique", Attachments: []string{doc.ID, img.ID}})

	reqs := sp.requests()
	if !hasSystemContaining(reqs, "Pythagore") {
		t.Fatal("le contenu du document doit etre injecte")
	}
	if hasSystemContaining(reqs, "photo.png") {
		t.Fatal("les images ne sont pas injectees au Lot A")
	}
}
