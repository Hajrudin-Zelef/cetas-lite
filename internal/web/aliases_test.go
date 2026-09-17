package web

import (
	"testing"

	"cetas-lite/internal/alias"
)

func TestValidateOverridesOK(t *testing.T) {
	ov := alias.Overrides{
		"samagent-nano": {"free": {{Provider: "openrouter", Model: "openrouter/free"}}},
		"samagent-n4":   {"flash": {{Provider: "openrouter", Model: "deepseek/deepseek-v3.2"}}},
		"samagent-n8":   {"elite": {{Provider: "deepseek", Model: "deepseek-v4-pro"}}},
		"samgen":        {"n4": {{Provider: "ollama", Model: "llama3.1:8b"}}},
		"code":          {"flash": {{Provider: "deepseek", Model: "deepseek-flash"}}},
	}
	if err := validateOverrides(ov); err != nil {
		t.Fatalf("overrides valides refuses : %v", err)
	}
	// Modele du pool par defaut mais absent du catalogue statique : accepte.
	defOnly := alias.Overrides{
		"code": {"flash": {{Provider: "opencode", Model: "qwen/qwen3.8-flash"}}},
	}
	if err := validateOverrides(defOnly); err != nil {
		t.Fatalf("modele du pool par defaut refuse : %v", err)
	}
	if err := validateOverrides(nil); err != nil {
		t.Fatalf("overrides vides refuses : %v", err)
	}
}

func TestValidateOverridesRejects(t *testing.T) {
	cases := map[string]alias.Overrides{
		"famille inconnue": {
			"nope": {"flash": {{Provider: "deepseek", Model: "deepseek-flash"}}},
		},
		"mode inconnu": {
			"samagent-n4": {"nope": {{Provider: "openrouter", Model: "deepseek/deepseek-v3.2"}}},
		},
		"pool vide": {
			"samagent-n4": {"flash": {}},
		},
		"fournisseur inconnu": {
			"samagent-n4": {"flash": {{Provider: "nope", Model: "x"}}},
		},
		"modele inconnu": {
			"samagent-n4": {"flash": {{Provider: "openrouter", Model: "nope/nope"}}},
		},
		"fournisseur local inconnu": {
			"samgen": {"n4": {{Provider: "nope", Model: "llama3.1:8b"}}},
		},
		"modele local vide": {
			"samgen": {"n4": {{Provider: "ollama", Model: ""}}},
		},
	}
	for name, ov := range cases {
		if err := validateOverrides(ov); err == nil {
			t.Errorf("%s : attendu une erreur, got nil", name)
		}
	}
}
