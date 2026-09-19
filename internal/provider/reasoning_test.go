package provider

import (
	"encoding/json"
	"testing"
)

func TestApplyReasoningDisabled(t *testing.T) {
	p := map[string]any{}
	applyReasoning(p, "openrouter", false, "")
	r, ok := p["reasoning"].(map[string]any)
	if !ok || r["enabled"] != false {
		t.Fatalf("openrouter off = %#v", p)
	}

	p = map[string]any{}
	applyReasoning(p, "deepseek", false, "")
	if p["chat_template_kwargs"] == nil {
		t.Fatalf("deepseek off = %#v", p)
	}

	p = map[string]any{}
	applyReasoning(p, "inconnu", false, "")
	if len(p) != 0 {
		t.Fatalf("provider inconnu ne doit rien ajouter: %#v", p)
	}
}

func TestApplyReasoningEffort(t *testing.T) {
	p := map[string]any{}
	applyReasoning(p, "openrouter", true, "low")
	r, _ := p["reasoning"].(map[string]any)
	if r["effort"] != "low" {
		t.Fatalf("openrouter effort = %#v", p)
	}

	p = map[string]any{}
	applyReasoning(p, "deepseek", true, "high")
	if p["reasoning_effort"] != "high" {
		t.Fatalf("deepseek effort = %#v", p)
	}

	p = map[string]any{}
	applyReasoning(p, "openrouter", true, "")
	if len(p) != 0 {
		t.Fatalf("effort vide ne doit rien ajouter: %#v", p)
	}
}

func TestWithReasoningContentForced(t *testing.T) {
	// Sans reasoning en jeu : payload strictement identique a avant.
	msgs := []Message{
		{Role: "user", Content: "hello"},
		{Role: "assistant", Content: "hi"},
	}
	a, _ := json.Marshal(withReasoningContentForced(msgs))
	b, _ := json.Marshal(msgs)
	if string(a) != string(b) {
		t.Fatalf("sans reasoning le payload doit etre identique:\n%s\n%s", a, b)
	}

	// Avec reasoning sur un seul message : tous les assistants portent
	// le champ sur le wire, meme vide.
	msgs2 := []Message{
		{Role: "user", Content: "q1"},
		{Role: "assistant", Content: "a1", ReasoningContent: "think1"},
		{Role: "assistant", Content: "a2", ToolCalls: []ToolCall{{ID: "1", Function: Func{Name: "Glob"}}}},
		{Role: "user", Content: "q2"},
		{Role: "assistant", Content: "a3"},
	}
	raw, err := json.Marshal(withReasoningContentForced(msgs2))
	if err != nil {
		t.Fatal(err)
	}
	var arr []map[string]any
	if err := json.Unmarshal(raw, &arr); err != nil {
		t.Fatal(err)
	}
	if len(arr) != 5 {
		t.Fatalf("attendu 5 messages, got %d", len(arr))
	}
	if arr[1]["reasoning_content"] != "think1" {
		t.Fatalf("reasoning du 1er assistant perdu: %v", arr[1]["reasoning_content"])
	}
	for _, i := range []int{2, 4} {
		rc, ok := arr[i]["reasoning_content"]
		if !ok {
			t.Fatalf("champ reasoning_content absent sur le message assistant %d", i)
		}
		if rc != "" {
			t.Fatalf("reasoning_content du message %d devrait etre vide, got %q", i, rc)
		}
	}
	// Les tool_calls survivent a la conversion forcee.
	tcs, ok := arr[2]["tool_calls"].([]any)
	if !ok || len(tcs) != 1 {
		t.Fatalf("tool_calls perdus dans le mode force: %v", arr[2]["tool_calls"])
	}
	// Les autres roles gardent leur serialisation d'origine.
	for _, i := range []int{0, 3} {
		if _, ok := arr[i]["reasoning_content"]; ok {
			t.Fatalf("un message %v ne doit pas porter reasoning_content", arr[i]["role"])
		}
	}
}
