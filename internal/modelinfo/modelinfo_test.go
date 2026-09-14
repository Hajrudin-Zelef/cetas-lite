package modelinfo

import "testing"

func TestLookupKnown(t *testing.T) {
	cases := []struct {
		provider, model string
		ctx             int
	}{
		{"deepseek", "deepseek-chat", 131072},
		{"openrouter", "deepseek/deepseek-chat-v3-0324", 131072},
		{"openai", "gpt-4o-mini", 128000},
		{"anthropic", "claude-sonnet-4-20250514", 200000},
		{"google", "gemini-2.5-flash", 1048576},
		{"mistral", "mistral-medium-2505", 131072},
		{"openrouter", "qwen/qwen-plus", 131072},
		{"groq", "llama-3.3-70b-versatile", 128000},
		{"xai", "grok-3-mini-beta", 131072},
	}
	for _, c := range cases {
		info, ok := Lookup(c.provider, c.model)
		if !ok {
			t.Fatalf("Lookup(%q,%q) = inconnu", c.provider, c.model)
		}
		if info.ContextWindow != c.ctx {
			t.Fatalf("Lookup(%q,%q) ctx = %d, want %d", c.provider, c.model, info.ContextWindow, c.ctx)
		}
		if info.InputPer1M <= 0 || info.OutputPer1M <= 0 {
			t.Fatalf("Lookup(%q,%q) tarifs manquants: %+v", c.provider, c.model, info)
		}
	}
}

func TestLookupUnknown(t *testing.T) {
	if _, ok := Lookup("monfournisseur", "modele-totalement-inconnu-xyz"); ok {
		t.Fatal("un modele inconnu ne doit pas matcher")
	}
	if _, ok := Lookup("", ""); ok {
		t.Fatal("un modele vide ne doit pas matcher")
	}
}
