package chat

import (
	"strings"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/memory"
	"cetas-lite/internal/provider"
)

func newMemEngine(t *testing.T) *Engine {
	t.Helper()
	e := NewEngine(nil, nil, nil, nil, "")
	e.SetMemory(memory.New(t.TempDir()))
	return e
}

func TestMemoryToolSchemas(t *testing.T) {
	names := map[string]bool{}
	for _, tool := range MemoryToolSchemas() {
		names[tool.Function.Name] = true
	}
	for _, want := range []string{"mem_search", "mem_read", "mem_add", "mem_edit", "mem_delete"} {
		if !names[want] {
			t.Errorf("schema manquant: %s", want)
		}
	}
}

func TestMemExecuteCRUD(t *testing.T) {
	e := newMemEngine(t)
	user := "alice"

	if res := e.memExecute(user, "mem_add", `{"name":"notes.md","content":"# Projet\nle chat est stable"}`); !strings.HasPrefix(res.Text, "[ok]") {
		t.Fatalf("add = %q", res.Text)
	}
	if res := e.memExecute(user, "mem_search", `{"query":"stable"}`); !strings.Contains(res.Text, "notes.md") {
		t.Fatalf("search = %q", res.Text)
	}
	if res := e.memExecute(user, "mem_read", `{"name":"notes.md"}`); !strings.Contains(res.Text, "le chat est stable") {
		t.Fatalf("read = %q", res.Text)
	}
	if res := e.memExecute(user, "mem_edit", `{"name":"notes.md","old":"stable","new":"solide"}`); !strings.HasPrefix(res.Text, "[ok]") {
		t.Fatalf("edit = %q", res.Text)
	}
	if res := e.memExecute(user, "mem_edit", `{"name":"notes.md","old":"stable","new":"solide"}`); !strings.Contains(res.Text, "deja a jour") {
		t.Fatalf("edit idempotent = %q", res.Text)
	}
	if res := e.memExecute(user, "mem_delete", `{"name":"notes.md"}`); !strings.HasPrefix(res.Text, "[ok]") {
		t.Fatalf("delete = %q", res.Text)
	}
	if res := e.memExecute(user, "mem_search", `{"query":"solide"}`); !strings.Contains(res.Text, "aucun resultat") {
		t.Fatalf("search apres delete = %q", res.Text)
	}
}

func TestMemoryIndexMessage(t *testing.T) {
	e := newMemEngine(t)
	if res := e.memExecute("alice", "mem_add", `{"name":"a.md","content":"# Alpha"}`); !strings.HasPrefix(res.Text, "[ok]") {
		t.Fatalf("add = %q", res.Text)
	}
	msg, ok := e.memoryIndexMessage("alice")
	if !ok {
		t.Fatal("index attendu")
	}
	content, _ := msg.Content.(string)
	if !strings.Contains(content, "Alpha") || !strings.Contains(content, "a.md") || !strings.Contains(content, "mem_read") {
		t.Fatalf("index message = %q", content)
	}
	if _, ok := e.memoryIndexMessage("bob"); ok {
		t.Fatal("bob n'a pas d'index")
	}
}

func TestMemoryDisabled(t *testing.T) {
	e := NewEngine(nil, nil, nil, nil, "")
	if _, ok := e.memoryIndexMessage("alice"); ok {
		t.Fatal("pas d'index sans memoire")
	}
	if res := e.memExecute("alice", "mem_search", `{"query":"x"}`); !strings.Contains(res.Text, "indisponible") {
		t.Fatalf("memExecute = %q", res.Text)
	}
}

func TestAgentMemoryToolRoundTrip(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{toolCall("c1", "mem_add", `{"name":"pref.md","content":"# Pref\nj'aime le cafe"}`)}},
		{toolCalls: []provider.ToolCall{toolCall("c2", "mem_search", `{"query":"cafe"}`)}},
		{content: "note"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	mem := memory.New(t.TempDir())
	e.SetMemory(mem)
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "retiens"})

	if got := mem.Content("sam", "pref.md"); !strings.Contains(got, "cafe") {
		t.Fatalf("page non persistee: %q", got)
	}
	if !hasToolEvent(c, "start") || !hasToolEvent(c, "end") {
		t.Fatal("evenements tool attendus")
	}
	reqs := sp.requests()
	if len(reqs) != 3 {
		t.Fatalf("requetes = %d, want 3", len(reqs))
	}
	toolsOffered := false
	for _, tool := range reqs[0].Tools {
		if tool.Function.Name == "mem_add" {
			toolsOffered = true
		}
	}
	if !toolsOffered {
		t.Fatal("mem_add doit etre propose en mode agent")
	}
	searchResult := false
	for _, m := range reqs[2].Messages {
		if m.Role == "tool" {
			if s, _ := m.Content.(string); strings.Contains(s, "pref.md") {
				searchResult = true
			}
		}
	}
	if !searchResult {
		t.Fatal("le resultat de mem_search doit revenir au modele")
	}
}

func TestNonAgentGetsMemoryIndex(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "salut"}}}
	fams := []alias.Family{{ID: "plain", Label: "Plain", Modes: []alias.Mode{{ID: "standard", Pool: []alias.Member{{Provider: "fake", Model: "m"}}}}}}
	e := newAgentEngine(t, sp, fams)
	mem := memory.New(t.TempDir())
	e.SetMemory(mem)
	if err := mem.Add("sam", "a.md", "# Alpha"); err != nil {
		t.Fatal(err)
	}
	runAgentTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "salut"})

	reqs := sp.requests()
	if len(reqs[0].Tools) != 0 {
		t.Fatalf("non-agent sans outils, tools = %d", len(reqs[0].Tools))
	}
	found := false
	for _, m := range reqs[0].Messages {
		if m.Role == "system" {
			if s, _ := m.Content.(string); strings.Contains(s, "Alpha") && strings.Contains(s, "Index memoire") {
				found = true
			}
		}
	}
	if !found {
		t.Fatal("l'index memoire doit etre injecte en non-agent")
	}
}
