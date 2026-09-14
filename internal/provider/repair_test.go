package provider

import (
	"encoding/json"
	"testing"
)

func TestRepairJSON(t *testing.T) {
	cases := []struct {
		name  string
		in    string
		valid bool
		check func(t *testing.T, out string)
	}{
		{name: "vide", in: "", valid: true, check: func(t *testing.T, out string) {
			if out != "{}" {
				t.Fatalf("attendu {}, got %q", out)
			}
		}},
		{name: "valide inchange", in: `{"a":1,"b":"x"}`, valid: true, check: func(t *testing.T, out string) {
			if out != `{"a":1,"b":"x"}` {
				t.Fatalf("got %q", out)
			}
		}},
		{name: "objet tronque", in: `{"file_path": "main.go", "content": "abc`, valid: true, check: func(t *testing.T, out string) {
			var m map[string]any
			if err := json.Unmarshal([]byte(out), &m); err != nil {
				t.Fatalf("resultat invalide %q: %v", out, err)
			}
			if m["file_path"] != "main.go" {
				t.Fatalf("file_path perdu: %q", out)
			}
		}},
		{name: "tableau tronque", in: `{"todos": [{"content": "a", "status": "pending"}, {"content": "b"`, valid: true, check: func(t *testing.T, out string) {
			if !json.Valid([]byte(out)) {
				t.Fatalf("resultat invalide: %q", out)
			}
		}},
		{name: "accolade manquante", in: `{"a": 1`, valid: true, check: func(t *testing.T, out string) {
			if out != `{"a": 1}` {
				t.Fatalf("got %q", out)
			}
		}},
		{name: "crochet ferme avant accolade", in: `{"a": [1, 2`, valid: true, check: func(t *testing.T, out string) {
			if !json.Valid([]byte(out)) {
				t.Fatalf("resultat invalide: %q", out)
			}
		}},
		{name: "crochet parasite final", in: `{"a": 1]`, valid: true, check: func(t *testing.T, out string) {
			if out != `{"a": 1}` {
				t.Fatalf("got %q", out)
			}
		}},
		{name: "que du bruit", in: `(((`, valid: false},
		{name: "chaine seule tronquee", in: `"abc`, valid: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out, ok := repairJSON(c.in)
			if ok != c.valid {
				t.Fatalf("repairJSON(%q) ok=%v, attendu %v (out=%q)", c.in, ok, c.valid, out)
			}
			if ok && c.check != nil {
				c.check(t, out)
			}
		})
	}
}
