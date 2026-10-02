package chat

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/attach"
	"cetas-lite/internal/provider"
)

func presentEngine(t *testing.T) (*Engine, *Sandbox, *attach.Store) {
	t.Helper()
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{toolCall("c1", "PresentFile", `{"file_path":"rapport.md"}`)}},
		{content: "voila"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	st := attach.New(t.TempDir(), 1<<20)
	e.SetAttachments(st)
	ws := t.TempDir()
	dir := filepath.Join(ws, "sam")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "rapport.md"), []byte("# Rapport\ncontenu"), 0o600); err != nil {
		t.Fatal(err)
	}
	sb, err := NewSandbox(dir)
	if err != nil {
		t.Fatal(err)
	}
	return e, sb, st
}

func TestPresentFileCreatesAttachment(t *testing.T) {
	e, sb, st := presentEngine(t)
	res := e.presentFile(t.Context(), "sam", sb, `{"file_path":"rapport.md"}`)
	if strings.HasPrefix(res.Text, "[error]") {
		t.Fatalf("presentFile: %s", res.Text)
	}
	info, ok := res.Meta["attachment"].(AttachmentInfo)
	if !ok || info.ID == "" {
		t.Fatalf("attachment attendu dans Meta, got %+v", res.Meta)
	}
	if info.Name != "rapport.md" {
		t.Fatalf("nom = %q", info.Name)
	}
	a, data, err := st.Get("sam", info.ID)
	if err != nil {
		t.Fatal(err)
	}
	if a.Name != "rapport.md" || string(data) != "# Rapport\ncontenu" {
		t.Fatalf("piece jointe = %+v %q", a, data)
	}
}

func TestPresentFileCustomName(t *testing.T) {
	e, sb, _ := presentEngine(t)
	res := e.presentFile(t.Context(), "sam", sb, `{"file_path":"rapport.md","name":"bilan.md"}`)
	info := res.Meta["attachment"].(AttachmentInfo)
	if info.Name != "bilan.md" {
		t.Fatalf("nom personnalise = %q", info.Name)
	}
}

func TestPresentFileMissing(t *testing.T) {
	e, sb, _ := presentEngine(t)
	res := e.presentFile(t.Context(), "sam", sb, `{"file_path":"absent.md"}`)
	if !strings.HasPrefix(res.Text, "[error]") {
		t.Fatalf("fichier absent: attendu [error], got %q", res.Text)
	}
}

func TestAttachmentInfos(t *testing.T) {
	e, _, st := presentEngine(t)
	a, err := st.Save("sam", "x.txt", []byte("x"))
	if err != nil {
		t.Fatal(err)
	}
	infos := e.AttachmentInfos("sam", []string{a.ID, "inconnu"})
	if len(infos) != 1 || infos[0].Name != "x.txt" || infos[0].Kind != "text" {
		t.Fatalf("infos = %+v", infos)
	}
}

func TestAttachmentDeltaCarriesInfos(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	st := attach.New(t.TempDir(), 1<<20)
	e.SetAttachments(st)
	doc, err := st.Save("sam", "c.md", []byte("Pythagore"))
	if err != nil {
		t.Fatal(err)
	}
	infos := e.AttachmentInfos("sam", []string{doc.ID})
	c := runTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "explique", Attachments: []string{doc.ID}, AttachmentInfos: infos})
	c.mu.Lock()
	found := false
	for _, ev := range c.Log {
		if _, ok := ev.Delta["user"]; ok {
			if atts, ok := ev.Delta["attachments"].([]AttachmentInfo); ok && len(atts) == 1 && atts[0].Name == "c.md" {
				found = true
			}
		}
	}
	c.mu.Unlock()
	if !found {
		t.Fatal("le delta user doit porter les metadonnees des pieces jointes")
	}
}
