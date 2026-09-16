package alias

import "testing"

func TestDefaultsStructure(t *testing.T) {
	fams := Defaults()
	wantOrder := []string{"samagent-nano", "samagent-n4", "samagent-n8", "code", "samgen"}
	if len(fams) != len(wantOrder) {
		t.Fatalf("familles = %d, want %d", len(fams), len(wantOrder))
	}
	for i, id := range wantOrder {
		if fams[i].ID != id {
			t.Errorf("famille %d = %q, want %q", i, fams[i].ID, id)
		}
	}
}

func TestN4PoolsRespectPriceRules(t *testing.T) {
	fams := Defaults()
	// paliers cumulatifs, seuils stricts (<)
	cases := []struct {
		mode  string
		max   float64
		count int
	}{
		{"flash", 1.0, 14},
		{"standard", 1.5, 20},
		{"elite", 6.0, 24},
	}
	for _, c := range cases {
		rm, ok := Resolve(fams, "samagent-n4", c.mode)
		if !ok {
			t.Fatalf("N4 %s introuvable", c.mode)
		}
		if len(rm.Pool) != c.count {
			t.Fatalf("N4 %s = %d modeles, want %d", c.mode, len(rm.Pool), c.count)
		}
		seen := map[string]bool{}
		for i, m := range rm.Pool {
			if m.Provider != "openrouter" {
				t.Errorf("N4 %s membre %d provider = %q", c.mode, i, m.Provider)
			}
			if m.OutputPer1M <= 0 || m.OutputPer1M >= c.max {
				t.Errorf("N4 %s membre %d (%s) output = %v, want < %v", c.mode, i, m.Model, m.OutputPer1M, c.max)
			}
			if seen[m.Model] {
				t.Errorf("N4 %s doublon: %s", c.mode, m.Model)
			}
			seen[m.Model] = true
			if i > 0 && rm.Pool[i-1].OutputPer1M > m.OutputPer1M {
				t.Errorf("N4 %s non trie par prix: %v > %v", c.mode, rm.Pool[i-1].OutputPer1M, m.OutputPer1M)
			}
		}
	}
	// composition exacte : la selection des 24, rien d'autre
	elite, _ := Resolve(fams, "samagent-n4", "elite")
	if len(n4Selection) != 24 {
		t.Fatalf("n4Selection = %d, want 24", len(n4Selection))
	}
	for _, id := range n4Selection {
		found := false
		for _, m := range elite.Pool {
			if m.Model == id {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("N4 elite: %s manquant", id)
		}
	}
}

func TestNanoPool(t *testing.T) {
	fams := Defaults()
	rm, ok := Resolve(fams, "samagent-nano", "free")
	if !ok {
		t.Fatal("nano free introuvable")
	}
	want := []Member{
		{"openrouter", "openrouter/free"},
		{"opencode", "big-pickle-zen"},
		{"opencode", "ling-3.0-flash-fin-free-zen"},
		{"opencode", "mimo-v2.5-free-zen"},
		{"opencode", "nemotron-3-ultra-free-zen"},
		{"opencode", "nemotron-3.5-lightning-free-zen"},
	}
	if len(rm.Pool) != len(want) {
		t.Fatalf("nano = %d modeles, want %d", len(rm.Pool), len(want))
	}
	for i, w := range want {
		if rm.Pool[i].Provider != w.Provider || rm.Pool[i].Model != w.Model {
			t.Errorf("nano[%d] = %v, want %v", i, rm.Pool[i], w)
		}
	}
}

func TestShufflePool(t *testing.T) {
	fams := Defaults()
	rm, _ := Resolve(fams, "samagent-nano", "free")
	orig := append([]ResolvedMember(nil), rm.Pool...)
	got := ShufflePool(rm.Pool)
	if len(got) != len(orig) {
		t.Fatalf("shuffle = %d, want %d", len(got), len(orig))
	}
	// ni perte ni doublon
	seen := map[string]int{}
	for _, m := range got {
		seen[m.Provider+"/"+m.Model]++
	}
	for _, m := range orig {
		k := m.Provider + "/" + m.Model
		if seen[k] != 1 {
			t.Errorf("shuffle: %s present %d fois", k, seen[k])
		}
	}
	// l'original ne doit pas avoir bouge
	for i := range orig {
		if rm.Pool[i] != orig[i] {
			t.Error("ShufflePool a mute le pool d'origine")
			break
		}
	}
}

func TestN8Mappings(t *testing.T) {
	fams := Defaults()
	cases := map[string]string{"flash": "deepseek-chat", "standard": "deepseek-flash", "elite": "deepseek-v4-pro"}
	for mode, model := range cases {
		rm, ok := Resolve(fams, "samagent-n8", mode)
		if !ok {
			t.Fatalf("N8 %s introuvable", mode)
		}
		if len(rm.Pool) == 0 || rm.Pool[0].Model != model {
			t.Fatalf("N8 %s tete de pool = %+v, want %s", mode, rm.Pool, model)
		}
		if rm.Agent {
			t.Errorf("N8 %s ne doit pas activer l'agent", mode)
		}
	}
	standard, _ := Resolve(fams, "samagent-n8", "standard")
	if len(standard.Pool) < 2 || standard.Pool[1].Provider != "opencode-go" {
		t.Fatalf("N8 standard doit avoir un failover opencode-go: %+v", standard.Pool)
	}
	for _, m := range standard.Pool[1:] {
		if m.OutputPer1M <= 0 || m.OutputPer1M >= 1.2 {
			t.Errorf("N8 standard failover %s output = %v, want < 1.2", m.Model, m.OutputPer1M)
		}
	}
	elite, _ := Resolve(fams, "samagent-n8", "elite")
	if len(elite.Pool) < 2 || elite.Pool[1].Provider != "opencode" {
		t.Fatalf("N8 elite doit avoir un failover opencode: %+v", elite.Pool)
	}
	for _, m := range elite.Pool[1:] {
		if m.OutputPer1M <= 1.5 {
			t.Errorf("N8 elite failover %s output = %v, want > 1.5", m.Model, m.OutputPer1M)
		}
	}
}

func TestCodeEnablesAgent(t *testing.T) {
	fams := Defaults()
	for _, mode := range []string{"flash", "standard", "elite"} {
		rm, ok := Resolve(fams, "code", mode)
		if !ok {
			t.Fatalf("code %s introuvable", mode)
		}
		if !rm.Agent {
			t.Errorf("code %s doit activer l'agent", mode)
		}
	}
}

func TestSamGenLocal(t *testing.T) {
	fams := Defaults()
	engines := map[string]string{"nano": "llamacpp", "n4": "ollama", "n8": "lmstudio"}
	for mode, engine := range engines {
		rm, ok := Resolve(fams, "samgen", mode)
		if !ok {
			t.Fatalf("samgen %s introuvable", mode)
		}
		if !rm.Local || rm.Engine != engine {
			t.Fatalf("samgen %s = local=%v engine=%q", mode, rm.Local, rm.Engine)
		}
		if rm.Agent {
			t.Errorf("samgen %s ne doit pas activer l'agent", mode)
		}
	}
}

func TestResolveLabelsFromCatalog(t *testing.T) {
	fams := Defaults()
	rm, _ := Resolve(fams, "samagent-n8", "elite")
	if rm.Pool[0].Label != "DeepSeek V4 Pro" || rm.Pool[0].OutputPer1M != 1.98 {
		t.Fatalf("label/prix = %+v", rm.Pool[0])
	}
}

func TestApplyOverrides(t *testing.T) {
	fams := Defaults()
	ov := Overrides{"samagent-n8": {"elite": {{"openrouter", "anthropic/claude-sonnet-4.5"}}}}
	got := Apply(fams, ov)
	rm, _ := Resolve(got, "samagent-n8", "elite")
	if len(rm.Pool) != 1 || rm.Pool[0].Model != "anthropic/claude-sonnet-4.5" {
		t.Fatalf("override non applique: %+v", rm.Pool)
	}
	// l'original ne doit pas avoir bouge
	orig, _ := Resolve(fams, "samagent-n8", "elite")
	if orig.Pool[0].Model != "deepseek-v4-pro" {
		t.Fatalf("defaults modifies par Apply: %+v", orig.Pool)
	}
}
