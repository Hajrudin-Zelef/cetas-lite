package provider

import "testing"

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
