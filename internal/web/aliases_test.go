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
	// Modele du pool par defaut : accepte (via le catalogue OU le pool).
	defOnly := alias.Overrides{
		"code": {"flash": {{Provider: "opencode", Model: "qwen3.6-plus-zen"}}},
	}
	if err := validateOverrides(defOnly); err != nil {
		t.Fatalf("modele du pool par defaut refuse : %v", err)
	}
	if err := validateOverrides(nil); err != nil {
		t.Fatalf("overrides vides refuses : %v", err)
	}
}

// inDefaultPool : exception a la validation catalogue (l'ID d'un modele du
// pool par defaut est transmis tel quel au provider, seule autorite reelle).
func TestInDefaultPool(t *testing.T) {
	if !inDefaultPool("code", "flash", alias.Member{Provider: "opencode", Model: "qwen3.6-plus-zen"}) {
		t.Fatal("modele du pool code/flash doit etre reconnu")
	}
	if inDefaultPool("code", "flash", alias.Member{Provider: "opencode", Model: "inconnu"}) {
		t.Fatal("modele absent du pool ne doit pas etre reconnu")
	}
	if inDefaultPool("code", "nope", alias.Member{Provider: "opencode", Model: "qwen3.6-plus-zen"}) {
		t.Fatal("mode inconnu ne doit pas etre reconnu")
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
