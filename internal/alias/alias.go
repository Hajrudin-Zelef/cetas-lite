package alias

import (
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
				{ID: "free", Label: "Free", Rule: "offres gratuites", Pool: freePool()},
			},
		},
		{
			ID: "samagent-n4", Label: "SamAgent N4",
			Modes: []Mode{
				{ID: "flash", Label: "Flash", Rule: "openrouter <= $0.30 / 1M out", Pool: openrouterByOutput(0.30, true)},
				{ID: "standard", Label: "Standard", Rule: "openrouter < $1.20 / 1M out", Pool: openrouterByOutput(1.20, false)},
			},
		},
		{
			ID: "samagent-n8", Label: "SamAgent N8",
			Modes: []Mode{
				{ID: "flash", Label: "Flash", Rule: "deepseek v3.2", Pool: []Member{{"deepseek", "deepseek-chat"}}},
				{ID: "standard", Label: "Standard", Rule: "deepseek v4 flash + failover",
					Pool: []Member{{"deepseek", "deepseek-flash"}, {"openrouter", "deepseek/deepseek-v4-flash-0731"}}},
				{ID: "elite", Label: "Elite", Rule: "deepseek v4 pro", Pool: []Member{{"deepseek", "deepseek-v4-pro"}}},
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

func freePool() []Member {
	var out []Member
	for _, m := range catalog.Models() {
		if m.InputPer1M == 0 && m.OutputPer1M == 0 {
			out = append(out, Member{Provider: m.Editeur, Model: m.ID})
		}
	}
	return out
}

func openrouterByOutput(max float64, inclusive bool) []Member {
	var ms []catalog.Model
	for _, m := range catalog.ByProvider("openrouter") {
		if m.OutputPer1M <= 0 {
			continue
		}
		if (inclusive && m.OutputPer1M <= max) || (!inclusive && m.OutputPer1M < max) {
			ms = append(ms, m)
		}
	}
	sort.Slice(ms, func(i, j int) bool {
		if ms[i].OutputPer1M != ms[j].OutputPer1M {
			return ms[i].OutputPer1M < ms[j].OutputPer1M
		}
		return ms[i].InputPer1M < ms[j].InputPer1M
	})
	out := make([]Member, 0, len(ms))
	for _, m := range ms {
		out = append(out, Member{Provider: "openrouter", Model: m.ID})
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
