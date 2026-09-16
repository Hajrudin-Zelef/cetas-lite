package alias

import (
	"crypto/rand"
	"math/big"
	"sort"

	"cetas-lite/internal/catalog"
)

type Member struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
}

type Mode struct {
	ID     string   `json:"id"`
	Label  string   `json:"label"`
	Agent  bool     `json:"agent,omitempty"`
	Local  bool     `json:"local,omitempty"`
	Engine string   `json:"engine,omitempty"`
	Rule   string   `json:"rule,omitempty"`
	Pool   []Member `json:"pool,omitempty"`
}

type Family struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Local bool   `json:"local,omitempty"`
	Modes []Mode `json:"modes"`
}

type ResolvedMember struct {
	Provider    string  `json:"provider"`
	Model       string  `json:"model"`
	Label       string  `json:"label"`
	InputPer1M  float64 `json:"input_per_1m"`
	OutputPer1M float64 `json:"output_per_1m"`
}

type ResolvedMode struct {
	Family string           `json:"family"`
	Mode   string           `json:"mode"`
	Label  string           `json:"label"`
	Agent  bool             `json:"agent"`
	Local  bool             `json:"local"`
	Engine string           `json:"engine,omitempty"`
	Rule   string           `json:"rule,omitempty"`
	Pool   []ResolvedMember `json:"pool"`
}

func Defaults() []Family {
	return []Family{
		{
			ID: "samagent-nano", Label: "SamAgent Nano",
			Modes: []Mode{
				{ID: "free", Label: "Free", Rule: "gratuits : routeur openrouter + zen (tirage aleatoire)", Pool: nanoPool()},
			},
		},
		{
			ID: "samagent-n4", Label: "SamAgent N4",
			Modes: []Mode{
				{ID: "flash", Label: "Flash", Rule: "openrouter < $1.00 / 1M out (selection N4)", Pool: n4Pool(1.0)},
				{ID: "standard", Label: "Standard", Rule: "openrouter < $1.50 / 1M out (selection N4)", Pool: n4Pool(1.5)},
				{ID: "elite", Label: "Elite", Rule: "openrouter < $6.00 / 1M out (selection N4)", Pool: n4Pool(6.0)},
			},
		},
		{
			ID: "samagent-n8", Label: "SamAgent N8",
			Modes: []Mode{
				{ID: "flash", Label: "Flash", Rule: "deepseek v3.2", Pool: []Member{{"deepseek", "deepseek-chat"}}},
				{ID: "standard", Label: "Standard", Rule: "deepseek v4.1 flash + failover opencode-go < $1.20 / 1M out",
					Pool: append([]Member{{"deepseek", "deepseek-flash"}}, providerPool("opencode-go", 0, 1.2)...)},
				{ID: "elite", Label: "Elite", Rule: "deepseek v4 pro + failover opencode > $1.50 / 1M out",
					Pool: append([]Member{{"deepseek", "deepseek-v4-pro"}}, providerPool("opencode", 1.5, 0)...)},
			},
		},
		{
			ID: "code", Label: "Code",
			Modes: []Mode{
				{ID: "flash", Label: "Flash", Agent: true, Rule: "deepseek flash",
					Pool: []Member{{"deepseek", "deepseek-flash"}, {"openrouter", "deepseek/deepseek-v4-flash-0731"}}},
				{ID: "standard", Label: "Standard", Agent: true, Rule: "pool coding opencode",
					Pool: []Member{
						{"opencode", "qwen3.7-plus-zen"},
						{"opencode", "minimax-m3-zen"},
						{"opencode-go", "qwen3.7-plus-go"},
						{"opencode", "glm-5-zen"},
					}},
				{ID: "elite", Label: "Elite", Agent: true, Rule: "deepseek v4 pro",
					Pool: []Member{{"deepseek", "deepseek-v4-pro"}}},
			},
		},
		{
			ID: "samgen", Label: "SamGen", Local: true,
			Modes: []Mode{
				{ID: "nano", Label: "Nano (llama.cpp)", Local: true, Engine: "llamacpp"},
				{ID: "n4", Label: "N4 (Ollama)", Local: true, Engine: "ollama"},
				{ID: "n8", Label: "N8 (LM Studio)", Local: true, Engine: "lmstudio"},
			},
		},
	}
}

type ResolvedFamily struct {
	ID    string         `json:"id"`
	Label string         `json:"label"`
	Local bool           `json:"local,omitempty"`
	Modes []ResolvedMode `json:"modes"`
}

func ResolveAll(fams []Family) []ResolvedFamily {
	out := make([]ResolvedFamily, 0, len(fams))
	for _, f := range fams {
		rf := ResolvedFamily{ID: f.ID, Label: f.Label, Local: f.Local, Modes: make([]ResolvedMode, 0, len(f.Modes))}
		for _, m := range f.Modes {
			rf.Modes = append(rf.Modes, resolveMode(f, m))
		}
		out = append(out, rf)
	}
	return out
}

func Find(fams []Family, familyID string) (Family, bool) {
	for _, f := range fams {
		if f.ID == familyID {
			return f, true
		}
	}
	return Family{}, false
}

func Resolve(fams []Family, familyID, modeID string) (ResolvedMode, bool) {
	f, ok := Find(fams, familyID)
	if !ok {
		return ResolvedMode{}, false
	}
	for _, m := range f.Modes {
		if m.ID == modeID {
			return resolveMode(f, m), true
		}
	}
	return ResolvedMode{}, false
}

func resolveMode(f Family, m Mode) ResolvedMode {
	rm := ResolvedMode{
		Family: f.ID, Mode: m.ID, Label: m.Label,
		Agent: m.Agent, Local: m.Local, Engine: m.Engine, Rule: m.Rule,
		Pool: []ResolvedMember{},
	}
	for _, mem := range m.Pool {
		label := mem.Model
		var in, out float64
		if cm, ok := catalog.Lookup(mem.Provider, mem.Model); ok {
			label = cm.Label
			in = cm.InputPer1M
			out = cm.OutputPer1M
		}
		rm.Pool = append(rm.Pool, ResolvedMember{
			Provider: mem.Provider, Model: mem.Model, Label: label,
			InputPer1M: in, OutputPer1M: out,
		})
	}
	return rm
}

// n4Selection : les 24 modeles OpenRouter retenus pour N4 (IDs exacts
// verifies via l'API OpenRouter le 2026-09-16).
var n4Selection = []string{
	"qwen/qwen3.7-flash",
	"poolside/laguna-xs-2.1",
	"z-ai/glm-5.3-flash",
	"poolside/laguna-s-2.1",
	"xiaomi/mimo-v2.5",
	"mistralai/mistral-small-2603",
	"qwen/qwen3.6-flash",
	"qwen/qwen3.7-plus",
	"meituan/longcat-2.0",
	"xiaomi/mimo-v2.5-pro",
	"google/gemini-2.5-flash",
	"meta-llama/llama-4-maverick",
	"meta/muse-spark-1.3-contributor",
	"deepseek/deepseek-v4-flash-0731",
	"openai/gpt-5.6-luna",
	"aion-labs/aion-3.0-mini",
	"anthropic/claude-haiku-4.5",
	"deepseek/deepseek-v4-pro",
	"google/gemma-4-26b-a4b-it",
	"arcee-ai/trinity-large-thinking",
	"openai/gpt-5.4-nano",
	"deepseek/deepseek-v3.2",
	"deepseek/deepseek-r1-0528",
	"meta-llama/llama-4-scout",
}

// nanoPool : agregation des modeles gratuits — routeur OpenRouter
// "Free Models Router" + les 5 modeles Zen gratuits d'OpenCode.
func nanoPool() []Member {
	return []Member{
		{"openrouter", "openrouter/free"},
		{"opencode", "big-pickle-zen"},
		{"opencode", "ling-3.0-flash-fin-free-zen"},
		{"opencode", "mimo-v2.5-free-zen"},
		{"opencode", "nemotron-3-ultra-free-zen"},
		{"opencode", "nemotron-3.5-lightning-free-zen"},
	}
}

func byOutputAsc(ms []catalog.Model) {
	sort.Slice(ms, func(i, j int) bool {
		if ms[i].OutputPer1M != ms[j].OutputPer1M {
			return ms[i].OutputPer1M < ms[j].OutputPer1M
		}
		return ms[i].InputPer1M < ms[j].InputPer1M
	})
}

// n4Pool : selection N4 filtree par prix de sortie strictement
// inferieur a maxOutput, triee par cout croissant. Les paliers sont
// cumulatifs : standard inclut flash, elite inclut standard.
func n4Pool(maxOutput float64) []Member {
	allow := make(map[string]bool, len(n4Selection))
	for _, id := range n4Selection {
		allow[id] = true
	}
	var ms []catalog.Model
	for _, m := range catalog.ByProvider("openrouter") {
		if !allow[m.ID] {
			continue
		}
		if m.OutputPer1M <= 0 || m.OutputPer1M >= maxOutput {
			continue
		}
		ms = append(ms, m)
	}
	byOutputAsc(ms)
	out := make([]Member, 0, len(ms))
	for _, m := range ms {
		out = append(out, Member{Provider: "openrouter", Model: m.ID})
	}
	return out
}

// providerPool : modeles d'un provider filtres par prix de sortie —
// minOutput exclusif (0 = sans borne basse), maxOutput exclusif
// (0 = sans borne haute) — tries par cout croissant.
func providerPool(provider string, minOutput, maxOutput float64) []Member {
	var ms []catalog.Model
	for _, m := range catalog.ByProvider(provider) {
		if m.OutputPer1M <= 0 || m.OutputPer1M <= minOutput {
			continue
		}
		if maxOutput > 0 && m.OutputPer1M >= maxOutput {
			continue
		}
		ms = append(ms, m)
	}
	byOutputAsc(ms)
	out := make([]Member, 0, len(ms))
	for _, m := range ms {
		out = append(out, Member{Provider: provider, Model: m.ID})
	}
	return out
}

// ShufflePool retourne une copie du pool melangee (Fisher-Yates,
// crypto/rand). Utilise par Nano : tirage aleatoire du premier modele,
// puis fallback sequentiel sur le reste en cas d'echec.
func ShufflePool(pool []ResolvedMember) []ResolvedMember {
	out := append([]ResolvedMember(nil), pool...)
	for i := len(out) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return out
		}
		out[i], out[j.Int64()] = out[j.Int64()], out[i]
	}
	return out
}

type Overrides map[string]map[string][]Member

func Apply(fams []Family, ov Overrides) []Family {
	out := make([]Family, len(fams))
	for i, f := range fams {
		nf := Family{ID: f.ID, Label: f.Label, Local: f.Local, Modes: make([]Mode, len(f.Modes))}
		for j, m := range f.Modes {
			nm := m
			if byFamily, ok := ov[f.ID]; ok {
				if pool, ok := byFamily[m.ID]; ok {
					nm.Pool = append([]Member(nil), pool...)
				}
			}
			nf.Modes[j] = nm
		}
		out[i] = nf
	}
	return out
}
