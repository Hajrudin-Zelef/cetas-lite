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
	flash, ok := Resolve(fams, "samagent-n4", "flash")
	if !ok {
		t.Fatal("N4 flash introuvable")
	}
	if len(flash.Pool) == 0 {
		t.Fatal("N4 flash vide")
	}
	for i, m := range flash.Pool {
		if m.Provider != "openrouter" {
			t.Errorf("N4 flash membre %d provider = %q", i, m.Provider)
		}
		if m.OutputPer1M <= 0 || m.OutputPer1M > 0.30 {
			t.Errorf("N4 flash membre %d output = %v", i, m.OutputPer1M)
		}
		if i > 0 && flash.Pool[i-1].OutputPer1M > m.OutputPer1M {
			t.Errorf("N4 flash non trie par prix: %v > %v", flash.Pool[i-1].OutputPer1M, m.OutputPer1M)
		}
	}

	std, ok := Resolve(fams, "samagent-n4", "standard")
	if !ok {
		t.Fatal("N4 standard introuvable")
	}
	for i, m := range std.Pool {
		if m.OutputPer1M <= 0 || m.OutputPer1M >= 1.20 {
			t.Errorf("N4 standard membre %d output = %v", i, m.OutputPer1M)
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
	if len(standard.Pool) < 2 || standard.Pool[1].Provider != "openrouter" {
		t.Fatalf("N8 standard doit avoir un failover openrouter: %+v", standard.Pool)
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
