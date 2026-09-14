// Package modelinfo fournit des metadonnees statiques sur les modeles cloud :
// fenetre de contexte et tarifs indicatifs ($ par million de tokens).
// Les valeurs sont indicatives et peuvent evoluer ; quand un modele est
// inconnu, Lookup retourne false et l'UI affiche "n/a".
package modelinfo

import "strings"

// Info decrit un modele connu.
type Info struct {
	// ContextWindow : taille max du contexte en tokens.
	ContextWindow int `json:"context_window"`
	// InputPer1M / OutputPer1M : tarifs indicatifs en $ par million de tokens.
	InputPer1M  float64 `json:"input_per_1m"`
	OutputPer1M float64 `json:"output_per_1m"`
}

type entry struct {
	match []string
	info  Info
}

// table : modeles cloud courants. Les cles sont matchees en sous-chaine
// insensible a la casse sur l'identifiant du modele.
var table = []entry{
	// DeepSeek
	{[]string{"deepseek-reasoner", "deepseek-r1", "/r1"}, Info{131072, 0.55, 2.19}},
	{[]string{"deepseek-chat", "deepseek-v3"}, Info{131072, 0.27, 1.10}},
	{[]string{"deepseek"}, Info{131072, 0.27, 1.10}},
	// OpenAI
	{[]string{"gpt-5", "gpt5"}, Info{400000, 1.25, 10.00}},
	{[]string{"gpt-4.1-nano", "gpt-41-nano"}, Info{1047576, 0.10, 0.40}},
	{[]string{"gpt-4.1-mini", "gpt-41-mini"}, Info{1047576, 0.40, 1.60}},
	{[]string{"gpt-4.1", "gpt-41"}, Info{1047576, 2.00, 8.00}},
	{[]string{"gpt-4o-mini"}, Info{128000, 0.15, 0.60}},
	{[]string{"gpt-4o"}, Info{128000, 2.50, 10.00}},
	{[]string{"o3-mini"}, Info{200000, 1.10, 4.40}},
	{[]string{"o3"}, Info{200000, 2.00, 8.00}},
	{[]string{"o1-mini"}, Info{128000, 1.10, 4.40}},
	{[]string{"o1"}, Info{200000, 15.00, 60.00}},
	// Anthropic
	{[]string{"claude-opus-4"}, Info{200000, 15.00, 75.00}},
	{[]string{"claude-sonnet-4"}, Info{200000, 3.00, 15.00}},
	{[]string{"claude-3-7-sonnet", "claude-sonnet-3-7"}, Info{200000, 3.00, 15.00}},
	{[]string{"claude-3-5-sonnet", "claude-sonnet-3-5"}, Info{200000, 3.00, 15.00}},
	{[]string{"claude-3-5-haiku", "claude-haiku-3-5", "claude-haiku"}, Info{200000, 0.80, 4.00}},
	{[]string{"claude"}, Info{200000, 3.00, 15.00}},
	// Google
	{[]string{"gemini-2.5-pro"}, Info{1048576, 1.25, 10.00}},
	{[]string{"gemini-2.5-flash"}, Info{1048576, 0.30, 2.50}},
	{[]string{"gemini-2.0-flash"}, Info{1048576, 0.10, 0.40}},
	{[]string{"gemini"}, Info{1048576, 0.30, 2.50}},
	// Mistral
	{[]string{"mistral-large"}, Info{131072, 2.00, 6.00}},
	{[]string{"mistral-medium"}, Info{131072, 0.40, 2.00}},
	{[]string{"mistral-small"}, Info{131072, 0.10, 0.30}},
	{[]string{"codestral"}, Info{262144, 0.30, 0.90}},
	{[]string{"mistral"}, Info{131072, 0.40, 2.00}},
	// Qwen (Alibaba)
	{[]string{"qwen-max"}, Info{262144, 1.60, 6.40}},
	{[]string{"qwen-plus"}, Info{131072, 0.40, 1.20}},
	{[]string{"qwen-turbo"}, Info{1048576, 0.05, 0.20}},
	{[]string{"qwen"}, Info{131072, 0.40, 1.20}},
	// Meta Llama
	{[]string{"llama-4-maverick", "maverick"}, Info{1048576, 0.35, 1.05}},
	{[]string{"llama-4-scout", "scout"}, Info{10485760, 0.15, 0.60}},
	{[]string{"llama-3.3-70b", "llama-3-3-70b"}, Info{128000, 0.35, 0.40}},
	{[]string{"llama-3.1-405b", "llama-405b"}, Info{128000, 3.50, 3.50}},
	{[]string{"llama-3.1-70b", "llama-70b"}, Info{128000, 0.35, 0.40}},
	{[]string{"llama"}, Info{128000, 0.35, 0.40}},
	// xAI
	{[]string{"grok-3-mini", "grok-3mini"}, Info{131072, 0.30, 0.50}},
	{[]string{"grok-3"}, Info{131072, 3.00, 15.00}},
	{[]string{"grok"}, Info{131072, 3.00, 15.00}},
	// Moonshot
	{[]string{"kimi-k2", "moonshot-v1-128k", "kimi"}, Info{262144, 0.60, 2.50}},
	// Zhipu / autres
	{[]string{"glm-4"}, Info{131072, 0.50, 0.50}},
}

// normalize nettoie un identifiant de modele pour la recherche.
func normalize(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.ReplaceAll(s, "_", "-")
	return s
}

// Lookup retourne les infos d'un modele a partir de son identifiant
// (et optionnellement du provider). La premiere entree dont un motif
// apparait dans l'identifiant gagne ; les entrees specifiques sont
// placees avant les generiques dans la table.
func Lookup(provider, model string) (Info, bool) {
	id := normalize(provider + " " + model)
	for _, e := range table {
		for _, m := range e.match {
			if strings.Contains(id, m) {
				return e.info, true
			}
		}
	}
	return Info{}, false
}
