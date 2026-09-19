package chat

// Tests Phase 2 : partition en runs parallèles, erreurs uniformes,
// troncature intelligente.

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"cetas-lite/internal/provider"
)

func phase2Call(id, name, args string) provider.ToolCall {
	tc := provider.ToolCall{ID: id, Type: "function"}
	tc.Function.Name = name
	tc.Function.Arguments = args
	return tc
}

func phase2State() *toolExecState {
	return &toolExecState{
		done:     map[string]string{},
		denied:   map[string]bool{},
		repeats:  map[string]int{},
		modified: map[string]bool{},
	}
}

// TestPartitionToolBlock vérifie le découpage en runs : les suites maximales
// d'appels indépendants partent en parallèle, le reste en séquentiel,
// dans l'ordre d'émission.
func TestPartitionToolBlock(t *testing.T) {
	readA := phase2Call("1", "Read", `{"file_path":"a.txt"}`)
	readB := phase2Call("2", "Read", `{"file_path":"b.txt"}`)
	readC := phase2Call("3", "Read", `{"file_path":"c.txt"}`)
	write := phase2Call("4", "Write", `{"file_path":"w.txt","content":"x"}`)
	unknown := phase2Call("5", "", `{}`)

	tests := []struct {
		name string
		tcs  []provider.ToolCall
		opts agentOpts
		// segments attendus : "P" = run parallèle (avec son nombre
		// d'appels), "S" = appel séquentiel isolé.
		want []string
	}{
		{
			name: "tout parallele",
			tcs:  []provider.ToolCall{readA, readB},
			want: []string{"P2"},
		},
		{
			name: "mixte lecture ecriture lecture",
			tcs:  []provider.ToolCall{readA, write, readB, readC},
			want: []string{"P1", "S", "P2"},
		},
		{
			name: "que des ecritures",
			tcs:  []provider.ToolCall{write, write},
			want: []string{"S", "S"},
		},
		{
			name: "outil inconnu isole",
			tcs:  []provider.ToolCall{readA, unknown, readB},
			want: []string{"P1", "S", "P1"},
		},
		{
			name: "doublon dans le run",
			tcs:  []provider.ToolCall{readA, readB, readA},
			want: []string{"P2", "P1"}, // le doublon sort du run ; run de 1 -> sequentiel
		},
		{
			name: "run de trois",
			tcs:  []provider.ToolCall{readA, readB, readC},
			want: []string{"P3"},
		},
		{
			name: "vide",
			tcs:  nil,
			want: nil,
		},
		{
			name: "approbation : write sequentiel, reads paralleles",
			tcs:  []provider.ToolCall{write, readA, readB},
			opts: agentOpts{approve: true},
			want: []string{"S", "P2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			segs := partitionToolBlock(tt.tcs, tt.opts, phase2State())
			var got []string
			for _, s := range segs {
				if s.parallel {
					got = append(got, fmt.Sprintf("P%d", len(s.calls)))
				} else {
					if len(s.calls) != 1 {
						t.Fatalf("segment sequentiel avec %d appels", len(s.calls))
					}
					got = append(got, "S")
				}
			}
			if len(got) != len(tt.want) {
				t.Fatalf("segments = %v, attendu %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("segments = %v, attendu %v", got, tt.want)
				}
			}
			// L'ordre des appels est préservé à travers les segments.
			var ids []string
			for _, s := range segs {
				for _, c := range s.calls {
					ids = append(ids, c.ID)
				}
			}
			for i, tc := range tt.tcs {
				if ids[i] != tc.ID {
					t.Fatalf("ordre non preserve : %v", ids)
				}
			}
		})
	}
}

// TestPartitionToolBlockPlan vérifie qu'en mode plan non validé, les outils
// hors allowlist restent séquentiels (ils seront refusés un par un).
func TestPartitionToolBlockPlan(t *testing.T) {
	read := phase2Call("1", "Read", `{"file_path":"a.txt"}`)
	write := phase2Call("2", "Write", `{"file_path":"w.txt","content":"x"}`)
	segs := partitionToolBlock([]provider.ToolCall{read, write}, agentOpts{plan: true}, phase2State())
	if len(segs) != 2 || !segs[0].parallel || segs[1].parallel {
		t.Fatalf("plan non valide : attendu [run(Read), seq(Write)], obtenu %v", segs)
	}
}

// sleepFamily simule des outils lents en suivant la concurrence maximale.
type sleepFamily struct {
	mu       sync.Mutex
	cur, max int
}

func (f *sleepFamily) schemas(context.Context) []provider.Tool { return nil }
func (f *sleepFamily) handles(name string) bool                { return name == "Read" }
func (f *sleepFamily) execute(context.Context, toolEnv, string, string) (ToolResult, *provider.Message) {
	f.mu.Lock()
	f.cur++
	if f.cur > f.max {
		f.max = f.cur
	}
	f.mu.Unlock()
	time.Sleep(80 * time.Millisecond)
	f.mu.Lock()
	f.cur--
	f.mu.Unlock()
	return ToolResult{Text: "ok"}, nil
}

func (f *sleepFamily) maxSeen() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.max
}

func testConversation() *Conversation {
	c := &Conversation{}
	c.cond = sync.NewCond(&c.mu)
	return c
}

// TestExecParallelRunConcurrency vérifie que les appels d'un run
// s'exécutent vraiment en concurrence, dans l'ordre pour les deltas.
func TestExecParallelRunConcurrency(t *testing.T) {
	fam := &sleepFamily{}
	reg := toolRegistry{families: []toolFamily{fam}}
	e := &Engine{}
	c := testConversation()
	tcs := []provider.ToolCall{
		phase2Call("call_1", "Read", `{"file_path":"a.txt"}`),
		phase2Call("call_2", "Read", `{"file_path":"b.txt"}`),
		phase2Call("call_3", "Read", `{"file_path":"c.txt"}`),
		phase2Call("call_4", "Read", `{"file_path":"d.txt"}`),
	}
	start := time.Now()
	outs, ok := e.execParallelRun(context.Background(), c, 0, reg, tcs, toolEnv{}, agentOpts{}, phase2State())
	el := time.Since(start)
	if !ok {
		t.Fatal("run 100 % lecture refuse, attendu ok=true")
	}
	if len(outs) != 4 {
		t.Fatalf("resultats = %d, attendu 4", len(outs))
	}
	// 4 x 80 ms en séquentiel = 320 ms ; en parallèle ~80 ms.
	if el > 250*time.Millisecond {
		t.Fatalf("execution trop lente (%v) : pas de parallelisme ?", el)
	}
	if m := fam.maxSeen(); m < 2 {
		t.Fatalf("concurrence max = %d, attendu >= 2", m)
	}
	// Ordre préservé : résultats et deltas dans l'ordre d'émission.
	for i, o := range outs {
		if o.id != tcs[i].ID {
			t.Fatalf("resultat %d hors ordre : %s", i, o.id)
		}
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	var phases []string
	for _, ev := range c.Log {
		if tm, ok := ev.Delta["tool"].(map[string]any); ok {
			phases = append(phases, tm["phase"].(string))
		}
	}
	if len(phases) != 8 {
		t.Fatalf("deltas outil = %d, attendu 8 (4 start + 4 end)", len(phases))
	}
	for i := 0; i < 4; i++ {
		if phases[i] != "start" || phases[i+4] != "end" {
			t.Fatalf("deltas hors ordre : %v", phases)
		}
	}
}

// TestExecParallelRunRefused vérifie le repli séquentiel quand un appel du
// run n'est plus éligible (ici : déjà exécuté = déduplication).
func TestExecParallelRunRefused(t *testing.T) {
	fam := &sleepFamily{}
	reg := toolRegistry{families: []toolFamily{fam}}
	e := &Engine{}
	c := testConversation()
	st := phase2State()
	st.done["Read\x00"+`{"file_path":"a.txt"}`] = "deja lu"
	tcs := []provider.ToolCall{
		phase2Call("call_1", "Read", `{"file_path":"a.txt"}`),
		phase2Call("call_2", "Read", `{"file_path":"b.txt"}`),
	}
	if _, ok := e.execParallelRun(context.Background(), c, 0, reg, tcs, toolEnv{}, agentOpts{}, st); ok {
		t.Fatal("run accepte malgre un appel deja execute, attendu ok=false")
	}
}

// TestUniformToolError vérifie l'uniformisation des erreurs : toute erreur
// brute devient "[erreur] <outil> : <cause> — <consigne>".
func TestUniformToolError(t *testing.T) {
	tests := []struct {
		name string
		tool string
		in   string
		want string // préfixe attendu ; "" = texte inchangé
		unch bool
	}{
		{
			name: "erreur systeme nue",
			tool: "Read",
			in:   "[erreur] open /x.txt: no such file or directory",
			want: "[erreur] Read : open /x.txt: no such file or directory — corrige les arguments et renvoie l'appel.",
		},
		{
			name: "deja actionnable : champ manquant",
			tool: "Read",
			in:   `[erreur] Read : champ requis manquant ou vide : "file_path" — renvoie l'appel avec ce champ renseigne.`,
			unch: true,
		},
		{
			name: "deja actionnable : consigne utilise",
			tool: "Read",
			in:   "[erreur] fichier introuvable: x. Utilise Glob pour trouver le bon chemin.",
			unch: true,
		},
		{
			name: "succes intact",
			tool: "Read",
			in:   "contenu du fichier",
			unch: true,
		},
		{
			name: "refus intact",
			tool: "Write",
			in:   "[refuse] l'utilisateur a refuse.",
			unch: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniformToolError(tt.tool, ToolResult{Text: tt.in}).Text
			if tt.unch {
				if got != tt.in {
					t.Fatalf("texte modifie : %q", got)
				}
				return
			}
			if got != tt.want {
				t.Fatalf("obtenu %q, attendu %q", got, tt.want)
			}
		})
	}
}

// TestTruncateToolForModelSmart vérifie la troncature intelligente :
// tête + queue conservées, marqueur d'omission, budget respecté.
func TestTruncateToolForModelSmart(t *testing.T) {
	var b strings.Builder
	for i := 1; i <= 500; i++ {
		b.WriteString("ligne " + strings.Repeat("x", 20) + "\n")
	}
	got := truncateToolForModel(b.String())
	if !strings.Contains(got, "lignes omises") {
		t.Fatalf("marqueur d'omission manquant : %q", got[len(got)-200:])
	}
	if !strings.HasPrefix(got, "ligne ") {
		t.Fatal("debut du resultat perdu")
	}
	if !strings.Contains(got, "Read offset/limit") {
		t.Fatal("indice de pagination manquant")
	}
	if n := len([]rune(got)); n > toolModelMaxChars+200 {
		t.Fatalf("budget depasse : %d runes", n)
	}
	// Ligne unique très longue : repli caractère tête+queue.
	long := strings.Repeat("y", toolModelMaxChars+500)
	got2 := truncateToolForModel(long)
	if !strings.Contains(got2, "caracteres omis") {
		t.Fatalf("repli caractere manquant : %q", got2[len(got2)-200:])
	}
	if !strings.HasPrefix(got2, "yyy") || !strings.HasSuffix(got2, "yyy") {
		t.Fatal("tete ou queue perdue en repli caractere")
	}
}
