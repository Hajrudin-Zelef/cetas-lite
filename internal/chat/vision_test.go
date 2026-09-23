package chat

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/attach"
	"cetas-lite/internal/modelcaps"
	"cetas-lite/internal/provider"
)

func logHasError(c *Conversation, marker string) bool {
	for _, ev := range c.marshal().Log {
		if s, ok := ev.Delta["error"].(string); ok && strings.Contains(s, marker) {
			return true
		}
	}
	return false
}

func contentHasImageURL(m provider.Message) bool {
	parts, ok := m.Content.([]any)
	if !ok {
		return false
	}
	for _, p := range parts {
		if pm, ok := p.(map[string]any); ok && pm["type"] == "image_url" {
			return true
		}
	}
	return false
}

func TestChatImageMultimodal(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	st := attach.New(t.TempDir(), 1<<20)
	e.SetAttachments(st)
	e.SetCapabilities(modelcaps.Map{modelcaps.Key("fake", "m"): {Vision: true}})
	img, err := st.Save("sam", "pic.png", []byte("PNGDATA"))
	if err != nil {
		t.Fatal(err)
	}

	runAgentTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "decris", Attachments: []string{img.ID}})

	found := false
	for _, m := range sp.requests()[0].Messages {
		if m.Role == "user" && contentHasImageURL(m) {
			found = true
		}
	}
	if !found {
		t.Fatal("le message utilisateur doit porter une image (multimodal)")
	}
}

func TestChatImageNoVisionModel(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, plainFamily(alias.Member{Provider: "fake", Model: "m"}))
	st := attach.New(t.TempDir(), 1<<20)
	e.SetAttachments(st)
	img, err := st.Save("sam", "pic.png", []byte("PNGDATA"))
	if err != nil {
		t.Fatal(err)
	}

	c := runAgentTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "decris", Attachments: []string{img.ID}})
	if !logHasError(c, "modele vision") {
		t.Fatal("un message clair est attendu sans modele vision")
	}
	if len(sp.requests()) != 0 {
		t.Fatal("aucun appel provider ne doit avoir lieu sans modele vision")
	}
}

func TestAgentViewImage(t *testing.T) {
	tmp := t.TempDir()
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{toolCall("c1", "ViewImage", `{"file_path":"img.png"}`)}},
		{content: "vu"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.workspace = filepath.Join(tmp, "workspace")
	dir := filepath.Join(e.workspace, "sam")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "img.png"), []byte("PNGDATA"), 0o600); err != nil {
		t.Fatal(err)
	}
	e.SetCapabilities(modelcaps.Map{modelcaps.Key("fake", "m"): {Vision: true}})

	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "regarde"})

	reqs := sp.requests()
	if len(reqs) < 2 {
		t.Fatalf("requetes = %d", len(reqs))
	}
	found := false
	for _, m := range reqs[1].Messages {
		if m.Role == "user" && contentHasImageURL(m) {
			found = true
		}
	}
	if !found {
		t.Fatal("ViewImage doit injecter une image pour l'iteration suivante")
	}
}

func TestAgentViewImageWithoutVision(t *testing.T) {
	tmp := t.TempDir()
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{toolCall("c1", "ViewImage", `{"file_path":"img.png"}`)}},
		{content: "vu"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.workspace = filepath.Join(tmp, "workspace")
	dir := filepath.Join(e.workspace, "sam")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "img.png"), []byte("PNGDATA"), 0o600); err != nil {
		t.Fatal(err)
	}

	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "regarde"})
	var toolResult string
	for _, m := range sp.requests()[1].Messages {
		if m.Role == "tool" {
			toolResult, _ = m.Content.(string)
		}
	}
	if !strings.Contains(toolResult, "cannot read images") {
		t.Fatalf("message clair attendu, got %q", toolResult)
	}
}
