package chat

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"

	"cetas-lite/internal/provider"
	"cetas-lite/internal/vfs"
)

func newTestSandbox(t *testing.T) *Sandbox {
	t.Helper()
	sb, err := NewSandbox(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	sb.AllowScript = true
	return sb
}

func TestToolSchemasCount(t *testing.T) {
	// 9 outils de base + 7 outils encadrés (Tree, Echo, Mkdir, Mv,
	// Sed, Awk, Curl — Cat fusionné dans Read en phase 3) + 9 outils
	// GitHub (repos, issues, PRs).
	// Glob reste natif et n'est pas dupliqué.
	if got := len(ToolSchemas()); got != 25 {
		t.Fatalf("schemas = %d, want 25", got)
	}
}

func TestWriteThenReadPagination(t *testing.T) {
	sb := newTestSandbox(t)
	res := sb.Execute(context.Background(), "Write", `{"file_path":"a/b.txt","content":"l1\nl2\nl3\nl4"}`)
	if strings.HasPrefix(res.Text, "[error]") {
		t.Fatalf("write: %s", res.Text)
	}
	if len(res.Diff) != 4 {
		t.Fatalf("write diff = %d lignes, want 4", len(res.Diff))
	}
	got := sb.Execute(context.Background(), "Read", `{"file_path":"a/b.txt","offset":2,"limit":2}`)
	if !strings.Contains(got.Text, "l2\nl3") || strings.Contains(got.Text, "l1") || strings.Contains(got.Text, "l4") {
		t.Fatalf("read pagination = %q", got.Text)
	}
}

func TestReadRejectsEscape(t *testing.T) {
	sb := newTestSandbox(t)
	got := sb.Execute(context.Background(), "Read", `{"file_path":"../../../etc/passwd"}`)
	if !strings.HasPrefix(got.Text, "[error]") {
		t.Fatalf("lecture hors sandbox acceptee: %q", got.Text)
	}
}

func TestEditCases(t *testing.T) {
	sb := newTestSandbox(t)
	sb.Execute(context.Background(), "Write", `{"file_path":"f.txt","content":"hello world"}`)

	ok := sb.Execute(context.Background(), "Edit", `{"file_path":"f.txt","old":"world","new":"terre"}`)
	if strings.HasPrefix(ok.Text, "[error]") {
		t.Fatalf("edit: %s", ok.Text)
	}
	again := sb.Execute(context.Background(), "Edit", `{"file_path":"f.txt","old":"world","new":"terre"}`)
	if !strings.Contains(again.Text, "already up to date") {
		t.Fatalf("edit deja applique: %q", again.Text)
	}
	missing := sb.Execute(context.Background(), "Edit", `{"file_path":"f.txt","old":"zzz","new":"x"}`)
	if !strings.HasPrefix(missing.Text, "[error]") {
		t.Fatalf("edit introuvable doit echouer: %q", missing.Text)
	}
	sb.Execute(context.Background(), "Write", `{"file_path":"d.txt","content":"a\na"}`)
	dup := sb.Execute(context.Background(), "Edit", `{"file_path":"d.txt","old":"a","new":"b"}`)
	if !strings.Contains(dup.Text, "2") {
		t.Fatalf("edit ambigu doit signaler le compte: %q", dup.Text)
	}
}

func TestGrepAndGlob(t *testing.T) {
	sb := newTestSandbox(t)
	sb.Execute(context.Background(), "Write", `{"file_path":"src/x.go","content":"package x\nfunc Foo() {}"}`)
	sb.Execute(context.Background(), "Write", `{"file_path":"src/y.txt","content":"nothing"}`)

	g := sb.Execute(context.Background(), "Grep", `{"pattern":"^func ","path":"."}`)
	if !strings.Contains(g.Text, "src/x.go:2") {
		t.Fatalf("grep = %q", g.Text)
	}
	gl := sb.Execute(context.Background(), "Glob", `{"pattern":"**/*.go"}`)
	if !strings.Contains(gl.Text, "src/x.go") || strings.Contains(gl.Text, "y.txt") {
		t.Fatalf("glob = %q", gl.Text)
	}
}

func TestLsExcludesHiddenAndGit(t *testing.T) {
	sb := newTestSandbox(t)
	os.MkdirAll(filepath.Join(sb.Root(), ".git"), 0o700)
	os.WriteFile(filepath.Join(sb.Root(), ".git", "config"), []byte("x"), 0o600)
	os.WriteFile(filepath.Join(sb.Root(), ".hidden"), []byte("x"), 0o600)
	sb.Execute(context.Background(), "Write", `{"file_path":"visible.txt","content":"x"}`)

	out := sb.Execute(context.Background(), "Ls", `{}`)
	if !strings.Contains(out.Text, "visible.txt") {
		t.Fatalf("ls doit voir visible.txt: %q", out.Text)
	}
	if strings.Contains(out.Text, ".git") || strings.Contains(out.Text, ".hidden") {
		t.Fatalf("ls ne doit pas voir caches/.git: %q", out.Text)
	}
}

func TestTodoWrite(t *testing.T) {
	sb := newTestSandbox(t)
	out := sb.Execute(context.Background(), "TodoWrite", `{"todos":[{"content":"a","status":"pending"}]}`)
	if strings.HasPrefix(out.Text, "[error]") {
		t.Fatalf("todo: %q", out.Text)
	}
}

func TestBashAllowlist(t *testing.T) {
	sb := newTestSandbox(t)
	ok := sb.Execute(context.Background(), "Bash", `{"command":"echo bonjour"}`)
	if !strings.Contains(ok.Text, "bonjour") || !strings.Contains(ok.Text, "exit: 0") {
		t.Fatalf("bash echo = %q", ok.Text)
	}
	denied := sb.Execute(context.Background(), "Bash", `{"command":"curl http://x"}`)
	if !strings.Contains(denied.Text, "not allowed") {
		t.Fatalf("bash curl doit etre refuse: %q", denied.Text)
	}
	flag := sb.Execute(context.Background(), "Bash", `{"command":"echo -e x"}`)
	if !strings.Contains(flag.Text, "forbidden") {
		t.Fatalf("flag -e doit etre refuse: %q", flag.Text)
	}
}

func TestBashCwdIsSandbox(t *testing.T) {
	sb := newTestSandbox(t)
	out := sb.Execute(context.Background(), "Bash", `{"command":"pwd"}`)
	if !strings.Contains(out.Text, sb.Root()) {
		t.Fatalf("cwd = %q, want %q", out.Text, sb.Root())
	}
}

func TestBashInterrupt(t *testing.T) {
	sb := newTestSandbox(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	out := sb.Execute(ctx, "Bash", `{"command":"sleep 60","timeout":60}`)
	if !strings.Contains(out.Text, "interrompue") {
		t.Fatalf("stop doit etre signale: %q", out.Text)
	}
}

func TestBashRejectsAbsolutePath(t *testing.T) {
	sb := newTestSandbox(t)
	out := sb.Execute(context.Background(), "Bash", `{"command":"cat /etc/hostname"}`)
	if !strings.Contains(out.Text, "absolu") && !strings.Contains(out.Text, "hors sandbox") {
		t.Fatalf("chemin absolu doit etre refuse: %q", out.Text)
	}
}

func TestBashRejectsParentTraversal(t *testing.T) {
	sb := newTestSandbox(t)
	out := sb.Execute(context.Background(), "Bash", `{"command":"cat ../secret.txt"}`)
	if !strings.Contains(out.Text, "hors sandbox") {
		t.Fatalf("traversee doit etre refusee: %q", out.Text)
	}
	out = sb.Execute(context.Background(), "Bash", `{"command":"cat ../../etc/passwd"}`)
	if !strings.Contains(out.Text, "hors sandbox") && !strings.Contains(out.Text, "absolu") {
		t.Fatalf("traversee profonde doit etre refusee: %q", out.Text)
	}
}

func TestBashRejectsEmbeddedAbsoluteFlag(t *testing.T) {
	sb := newTestSandbox(t)
	out := sb.Execute(context.Background(), "Bash", `{"command":"git -C/etc status"}`)
	if !strings.Contains(out.Text, "hors sandbox") {
		t.Fatalf("flag avec chemin absolu doit etre refuse: %q", out.Text)
	}
	out = sb.Execute(context.Background(), "Bash", `{"command":"git --git-dir=/etc status"}`)
	if !strings.Contains(out.Text, "hors sandbox") {
		t.Fatalf("flag =/absolu doit etre refuse: %q", out.Text)
	}
}

func TestBashAllowsRelativeInside(t *testing.T) {
	sb := newTestSandbox(t)
	write := sb.Execute(context.Background(), "Write", `{"file_path":"sub/note.txt","content":"bonjour"}`)
	if strings.HasPrefix(write.Text, "[error]") {
		t.Fatalf("write: %q", write.Text)
	}
	out := sb.Execute(context.Background(), "Bash", `{"command":"cat sub/note.txt"}`)
	if !strings.Contains(out.Text, "bonjour") {
		t.Fatalf("chemin relatif interne doit etre autorise: %q", out.Text)
	}
}

func TestRunScriptDisabledByDefault(t *testing.T) {
	sb, err := NewSandbox(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	out := sb.Execute(context.Background(), "RunScript", `{"language":"python","code":"print(1)"}`)
	if !strings.Contains(out.Text, "disabled") {
		t.Fatalf("RunScript doit etre desactive par defaut: %q", out.Text)
	}
}

func TestRunScript(t *testing.T) {
	runner := ""
	for _, c := range []string{"python3", "node"} {
		if _, err := exec.LookPath(c); err == nil {
			runner = c
			break
		}
	}
	if runner == "" {
		t.Skip("ni python3 ni node disponible")
	}
	sb := newTestSandbox(t)
	lang := "python"
	code := `print("salut")`
	if runner == "node" {
		lang = "node"
		code = `console.log("salut")`
	}
	out := sb.Execute(context.Background(), "RunScript", `{"language":"`+lang+`","code":`+strconv.Quote(code)+`}`)
	if !strings.Contains(out.Text, "salut") {
		t.Fatalf("runscript = %q", out.Text)
	}
	bad := sb.Execute(context.Background(), "RunScript", `{"language":"ruby","code":"x"}`)
	if !strings.Contains(bad.Text, "Unsupported") && !strings.Contains(bad.Text, "support") {
		t.Fatalf("langage non supporte: %q", bad.Text)
	}
}

func TestUnknownTool(t *testing.T) {
	sb := newTestSandbox(t)
	out := sb.Execute(context.Background(), "Nope", `{}`)
	if !strings.HasPrefix(out.Text, "[error]") {
		t.Fatalf("outil inconnu: %q", out.Text)
	}
}

func TestTruncateToolForModel(t *testing.T) {
	short := "petit resultat"
	if got := truncateToolForModel(short); got != short {
		t.Fatalf("texte court modifié : %q", got)
	}
	long := strings.Repeat("x", toolModelMaxChars+500)
	got := truncateToolForModel(long)
	if len([]rune(got)) > toolModelMaxChars+200 {
		t.Fatalf("texte long non tronqué : %d runes", len([]rune(got)))
	}
	if !strings.Contains(got, "Read offset/limit") {
		t.Fatalf("indice de pagination manquant : %q", got[len(got)-120:])
	}
}

func TestReadDefaultLimit(t *testing.T) {
	sb := newTestSandbox(t)
	var b strings.Builder
	for i := 1; i <= 500; i++ {
		b.WriteString("ligne " + strconv.Itoa(i) + "\n")
	}
	if res := sb.Execute(context.Background(), "Write", `{"file_path":"gros.txt","content":`+strconv.Quote(b.String())+`}`); strings.HasPrefix(res.Text, "[error]") {
		t.Fatalf("write: %s", res.Text)
	}
	got := sb.Execute(context.Background(), "Read", `{"file_path":"gros.txt"}`)
	lines := strings.Split(strings.TrimSuffix(got.Text, "\n"), "\n")
	// 400 lignes + la ligne de compteur "… (500 lignes au total, …)".
	if len(lines) != defaultReadLines+1 {
		t.Fatalf("read sans limite = %d lignes, want %d", len(lines), defaultReadLines+1)
	}
	if !strings.Contains(got.Text, "lignes au total") {
		t.Fatalf("compteur de pagination manquant : %q", got.Text[len(got.Text)-80:])
	}
	// Une limite explicite reste honorée.
	got2 := sb.Execute(context.Background(), "Read", `{"file_path":"gros.txt","limit":10}`)
	if n := len(strings.Split(strings.TrimSpace(got2.Text), "\n")); n != 11 {
		t.Fatalf("read limit=10 = %d lignes, want 11 (10 + compteur)", n)
	}
}

// TestBannedEvalFlag vérifie la détection des flags d'évaluation collés
// (H1) : python3 -c<charge>, node --eval=<charge> — sans casser les usages
// légitimes des autres binaires (grep -c, sed -e...).
func TestBannedEvalFlag(t *testing.T) {
	banned := [][2]string{
		{"python3", "-c"}, {"python3", "-cprint(1)"}, {"python3", "-c print(1)"},
		{"python", "-c"}, {"node", "-e"}, {"node", "-e1+1"},
		{"node", "--eval"}, {"node", "--eval=1+1"},
	}
	for _, c := range banned {
		if !bannedEvalFlag(c[0], c[1]) {
			t.Errorf("bannedEvalFlag(%q, %q) = false, want true", c[0], c[1])
		}
	}
	allowed := [][2]string{
		{"grep", "-c"}, {"grep", "-e"}, {"grep", "--color"},
		{"sed", "-e"}, {"sed", "--expression=s/a/b/"},
		{"echo", "-e"}, {"ls", "-c"}, {"find", "-c"},
		{"python3", "--version"}, {"node", "--version"},
	}
	for _, c := range allowed {
		if bannedEvalFlag(c[0], c[1]) {
			t.Errorf("bannedEvalFlag(%q, %q) = true, want false (usage legitime)", c[0], c[1])
		}
	}
}

// TestBashRejectsGluedEvalFlags vérifie de bout en bout que les formes
// collées sont refusées avant toute exécution (H1).
func TestBashRejectsGluedEvalFlags(t *testing.T) {
	sb := newTestSandbox(t)
	for _, cmd := range []string{
		`python3 -cprint(1)`,
		`python3 -c print(1)`,
		`node --eval=1+1`,
		`node -e1+1`,
	} {
		out := sb.Execute(context.Background(), "Bash", `{"command":`+strconv.Quote(cmd)+`}`)
		if !strings.Contains(out.Text, "forbidden flag") {
			t.Errorf("commande %q aurait du etre refusee: %q", cmd, out.Text)
		}
	}
}

// TestBashRejectsFindExec vérifie que find -exec/-execdir est banni (H3).
func TestBashRejectsFindExec(t *testing.T) {
	sb := newTestSandbox(t)
	for _, cmd := range []string{
		`find . -exec echo {} \;`,
		`find . -execdir echo {} \;`,
	} {
		out := sb.Execute(context.Background(), "Bash", `{"command":`+strconv.Quote(cmd)+`}`)
		if !strings.Contains(out.Text, "forbidden flag") {
			t.Errorf("commande %q aurait du etre refusee: %q", cmd, out.Text)
		}
	}
	// find sans -exec reste utilisable.
	out := sb.Execute(context.Background(), "Bash", `{"command":"find . -maxdepth 1"}`)
	if strings.Contains(out.Text, "forbidden flag") {
		t.Errorf("find simple ne devrait pas etre bloque: %q", out.Text)
	}
}

// TestBashRejectsBinaryPath vérifie qu'un binaire passé par chemin
// (shadowing) est refusé (H6).
func TestBashRejectsBinaryPath(t *testing.T) {
	sb := newTestSandbox(t)
	for _, cmd := range []string{"outils/ls", "/bin/ls", "./ls", "../bin/ls"} {
		out := sb.Execute(context.Background(), "Bash", `{"command":`+strconv.Quote(cmd)+`}`)
		if !strings.Contains(out.Text, "binary path forbidden") {
			t.Errorf("commande %q aurait du etre refusee: %q", cmd, out.Text)
		}
	}
}

// TestBashEmbeddedProgram vérifie l'extraction du programme sed/awk
// embarqué dans une commande Bash, y compris avec des options à argument
// séparé (-v, -F) qui ne doivent pas masquer le vrai programme.
func TestBashEmbeddedProgram(t *testing.T) {
	cases := []struct {
		binary string
		args   []string
		want   string
	}{
		{"sed", []string{"-n", "1,5p", "f.txt"}, "1,5p"},
		{"awk", []string{"{print $1}", "f.txt"}, "{print $1}"},
		{"awk", []string{"-F,", "{print $2}", "f.txt"}, "{print $2}"},
		{"sed", []string{"-es/a/b/", "f.txt"}, "s/a/b/"},
		{"sed", []string{"--expression=s/a/b/", "f.txt"}, "s/a/b/"},
		{"awk", []string{"-v", "x=1", "{print $1}", "f.txt"}, "{print $1}"},
		{"awk", []string{"-F", ":", "BEGIN{print}", "f.txt"}, "BEGIN{print}"},
		{"awk", []string{"-v", "x=1", "BEGIN{system(\"id\")}", "f.txt"}, "BEGIN{system(\"id\")}"},
		{"sed", []string{"-e", "s/a/b/", "f.txt"}, "s/a/b/"},
	}
	for _, c := range cases {
		if got := bashEmbeddedProgram(c.binary, c.args); got != c.want {
			t.Errorf("bashEmbeddedProgram(%s, %v) = %q, want %q", c.binary, c.args, got, c.want)
		}
	}
}

// TestBashProgramRisk vérifie la détection des programmes sed/awk à effets
// de bord via Bash (H2, H5, H7) et l'absence de faux positifs courants.
func TestBashProgramRisk(t *testing.T) {
	risky := [][2]string{
		{"awk", `BEGIN{system("id")}`},
		{"gawk", `{print | "sort"}`},
		{"mawk", `{print "x" > "/tmp/f"}`},
		{"sed", "r /etc/passwd"},
		{"sed", "r/etc/passwd"},
		{"sed", "w/tmp/x"},
		{"sed", "s/a/b/w /tmp/x"},
		{"sed", "e whoami"},
	}
	for _, c := range risky {
		if msg := bashProgramRisk(c[0], c[1]); msg == "" {
			t.Errorf("bashProgramRisk(%q, %q) vide, risque attendu", c[0], c[1])
		}
	}
	safe := [][2]string{
		{"awk", "{print $1}"},
		{"awk", "NR>1 {sum+=$3} END {print sum}"},
		{"sed", "s/a/b/g"},
		{"sed", "10,20p"},
		{"sed", "/^#/d"},
		{"sed", "/root/d"},
		{"echo", "bonjour"},
	}
	for _, c := range safe {
		if msg := bashProgramRisk(c[0], c[1]); msg != "" {
			t.Errorf("bashProgramRisk(%q, %q) = %q, faux positif", c[0], c[1], msg)
		}
	}
}

// TestBashRejectsRiskyAwkProgram vérifie de bout en bout le refus d'un
// programme awk à effets de bord via Bash (H2).
func TestBashRejectsRiskyAwkProgram(t *testing.T) {
	sb := newTestSandbox(t)
	out := sb.Execute(context.Background(), "Bash", `{"command":"awk 'BEGIN{system(\"id\")}'"}`)
	if !strings.Contains(out.Text, "dedicated Awk tool") {
		t.Fatalf("awk system() via Bash aurait du etre refuse: %q", out.Text)
	}
	out = sb.Execute(context.Background(), "Bash", `{"command":"sed 'r /etc/passwd'"}`)
	if !strings.Contains(out.Text, "dedicated Sed tool") {
		t.Fatalf("sed r via Bash aurait du etre refuse: %q", out.Text)
	}
	// Une option à argument séparé (-v) ne doit pas masquer le programme.
	out = sb.Execute(context.Background(), "Bash", `{"command":"awk -v x=1 'BEGIN{system(\"id\")}'"}`)
	if !strings.Contains(out.Text, "dedicated Awk tool") {
		t.Fatalf("awk -v ... system() via Bash aurait du etre refuse: %q", out.Text)
	}
	// Un awk -v légitime reste utilisable.
	sb2 := newTestSandbox(t)
	if err := os.WriteFile(filepath.Join(sb2.localRoot(), "d.txt"), []byte("a b\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	out = sb2.Execute(context.Background(), "Bash", `{"command":"awk -v OFS=, '{print $1}' d.txt"}`)
	if strings.Contains(out.Text, "dedicated Awk tool") {
		t.Fatalf("awk -v legitime refuse a tort: %q", out.Text)
	}
}

// TestMustApproveFor (C1) : Bash/RunScript/Curl exigent une approbation
// quel que soit le mode ; les autres outils gardent leur comportement.
func TestMustApproveFor(t *testing.T) {
	for _, name := range []string{"Bash", "RunScript", "Curl"} {
		if !mustApproveFor(name, map[string]any{}) {
			t.Fatalf("%s devrait exiger une approbation meme en Espace Write", name)
		}
	}
	for _, name := range []string{"Write", "Edit", "Read", "Ls", "Echo"} {
		if mustApproveFor(name, map[string]any{}) {
			t.Fatalf("%s ne devrait pas etre force par mustApproveFor", name)
		}
	}
	// Réserve sed -i : sed in-place via Bash force l'approbation
	// (ici via l'appartenance de Bash à mustApproveTools ; la détection
	// fine est testée unitairement dans TestBashHasSedInplace).
	if !mustApproveFor("Bash", map[string]any{"command": "sed -i 's/a/b/' f.txt"}) {
		t.Fatal("sed -i via Bash devrait forcer l'approbation")
	}
}

// TestBashHasSedInplace : détection du flag in-place de sed via Bash.
func TestBashHasSedInplace(t *testing.T) {
	cases := []struct {
		cmd  string
		want bool
	}{
		{"sed -i 's/a/b/' f.txt", true},
		{"sed -i.bak 's/a/b/' f.txt", true},
		{`sed -i'' 's/a/b/' f.txt`, true},
		{`sed -i "" 's/a/b/' f.txt`, true},
		{"sed --in-place 's/a/b/' f.txt", true},
		{"sed --in-place=.bak 's/a/b/' f.txt", true},
		{"sed -in 's/a/b/' f.txt", true}, // -i -n combinés
		{"sed -e 's/a/b/' -i f.txt", true},
		{"sed -n 's/a/b/p' f.txt", false}, // flux pur
		{"sed 's/a/b/' f.txt", false},     // flux pur
		{"grep -i motif f.txt", false},    // -i de grep, pas sed
		{"echo 'sed -i'", false},          // pas un appel sed
		{"", false},
	}
	for _, c := range cases {
		if got := bashHasSedInplace(map[string]any{"command": c.cmd}); got != c.want {
			t.Fatalf("bashHasSedInplace(%q) = %v, want %v", c.cmd, got, c.want)
		}
	}
}

// TestToolEchoRuneTruncation (B1) : la troncature ne doit jamais couper
// un rune UTF-8 en deux.
func TestToolEchoRuneTruncation(t *testing.T) {
	sb := &Sandbox{}
	// "a" décale l'alignement : une troncature en octets couperait un "é".
	text := "a" + strings.Repeat("é", echoMaxLen)
	out := sb.toolEcho(context.Background(), map[string]any{"text": text})
	if !utf8.ValidString(out.Text) {
		t.Fatalf("sortie Echo invalide en UTF-8 : %q...", out.Text[:50])
	}
	if n := utf8.RuneCountInString(out.Text); n > echoMaxLen+20 {
		t.Fatalf("tronqué trop long : %d runes", n)
	}
}

// TestGlobStarBound (B2) : les motifs à trop de ** sont refusés avant
// le parcours.
func TestGlobStarBound(t *testing.T) {
	sb := &Sandbox{}
	r := sb.Execute(context.Background(), "Glob", `{"pattern":"**/**/**/**/**/**/**/**/**/x"}`)
	if !strings.Contains(r.Text, "too complex") {
		t.Fatalf("motif à 9 ** aurait du etre refuse : %q", r.Text)
	}
}

// walkErrFS : FS dont le Walk échoue (B3).
type walkErrFS struct {
	vfs.FS
	err error
}

func (f walkErrFS) Walk(ctx context.Context, fn func(vfs.Entry) error) error {
	return f.err
}

// TestWalkErrorWarnsModel (B3) : une erreur de parcours remonte un
// avertissement au modèle au lieu d'être silencieuse.
func TestWalkErrorWarnsModel(t *testing.T) {
	ctx := context.Background()
	lfs, err := vfs.NewLocal(filepath.Join(t.TempDir(), "ws"), "test")
	if err != nil {
		t.Fatal(err)
	}
	sb := NewSandboxFS(walkErrFS{FS: lfs, err: errors.New("boom")})
	cases := map[string]string{
		"Ls":   `{}`,
		"Glob": `{"pattern":"*.go"}`,
		"Grep": `{"pattern":"x"}`,
	}
	for name, args := range cases {
		r := sb.Execute(ctx, name, args)
		if !strings.Contains(r.Text, "[warning] incomplete walk") {
			t.Fatalf("%s : avertissement attendu, got %q", name, r.Text)
		}
	}
}

// TestCurlModelHeadersFiltersAuthorization (B4) : le modèle ne forge pas
// d'en-tête Authorization.
func TestCurlModelHeadersFiltersAuthorization(t *testing.T) {
	hdrs := curlModelHeaders(map[string]any{"headers": map[string]any{
		"Authorization": "Bearer secret-du-modele",
		"X-Custom":      "ok",
		"authorization": "Basic xyz",
		"":              "vide",
	}})
	if _, ok := hdrs["Authorization"]; ok {
		t.Fatal("Authorization aurait du etre filtre")
	}
	if _, ok := hdrs["authorization"]; ok {
		t.Fatal("authorization (casse differente) aurait du etre filtre")
	}
	if hdrs["X-Custom"] != "ok" {
		t.Fatal("les autres en-tetes doivent passer")
	}
}

// TestToolCurlRejectsUserinfo (B4) : userinfo dans l'URL refusé sans réseau.
func TestToolCurlRejectsUserinfo(t *testing.T) {
	sb := &Sandbox{}
	r := sb.toolCurl(context.Background(), map[string]any{"url": "http://user:pass@example.com/"})
	if !strings.Contains(r.Text, "userinfo") {
		t.Fatalf("userinfo aurait du etre refuse : %q", r.Text)
	}
}

// TestClientSafeError (F15) : le message client est générique, le détail
// ne fuit pas vers l'écran.
func TestClientSafeError(t *testing.T) {
	raw := errors.New("request req-abc-123 failed for model deepseek-x at /home/user/secret")
	got := clientSafeError(raw, "Échec de la génération — réessaie dans un instant.")
	if got != "Échec de la génération — réessaie dans un instant." {
		t.Fatalf("message inattendu : %q", got)
	}
	for _, leak := range []string{"req-abc-123", "deepseek-x", "/home/user/secret"} {
		if strings.Contains(got, leak) {
			t.Fatalf("fuite de detail serveur vers le client : %q", leak)
		}
	}
}

// TestToolDescriptionsAllEnglish : règle globale — aucun texte FR envoyé au
// modèle. Balaye les descriptions de tous les outils (y compris les
// descriptions de paramètres imbriquées) à la recherche de diacritiques
// français. C'est ce balayage qui aurait attrapé les 20 descriptions
// oubliées par la traduction initiale (memory, extra, github, web).
func TestToolDescriptionsAllEnglish(t *testing.T) {
	all := [][]provider.Tool{
		ToolSchemas(),
		RagToolSchemas(),
		MemoryToolSchemas(),
		extraToolSchemas(),
		githubToolSchemas(),
		WebToolSchemas(),
	}
	const frDiacritics = "éèêëàâäîïôöùûüç"
	var walk func(v any, where string)
	walk = func(v any, where string) {
		switch x := v.(type) {
		case string:
			for _, r := range x {
				if strings.ContainsRune(frDiacritics, r) {
					t.Errorf("texte FR dans %s : %q", where, x)
					return
				}
			}
		case map[string]any:
			for _, v2 := range x {
				walk(v2, where)
			}
		case []any:
			for _, v2 := range x {
				walk(v2, where)
			}
		}
	}
	for _, schemas := range all {
		for _, tool := range schemas {
			where := "outil " + tool.Function.Name
			walk(tool.Function.Description, where+" description")
			walk(tool.Function.Parameters, where+" paramètres")
		}
	}
}

// TestToolResultMarkersAllEnglish : règle globale — aucun résultat d'outil
// envoyé au modèle ne doit contenir les marqueurs FR [erreur], [refuse] ou
// [avertissement]. Exerce les chemins d'échec de chaque famille et vérifie
// que seuls [error], [denied], [warning], [info] et [ok] y figurent.
func TestToolResultMarkersAllEnglish(t *testing.T) {
	ctx := context.Background()
	sb := newTestSandbox(t)
	var texts []string
	collect := func(s string) { texts = append(texts, s) }

	for _, tc := range [][2]string{
		{"Read", `{"file_path":"../evil"}`},
		{"Read", `{"file_path":"nope.txt"}`},
		{"Write", `{"file_path":""}`},
		{"Edit", `{"file_path":"a","old":"","new":"b"}`},
		{"Grep", `{"pattern":""}`},
		{"Grep", `{"pattern":"(["}`},
		{"Glob", `{"pattern":"**/**/**/**/**/**/**/**/**/x"}`},
		{"Bash", `{"command":""}`},
		{"Bash", `{"command":"curl x"}`},
		{"Bash", `{"command":"rm -rf /"}`},
		{"Bash", `{"command":"/bin/ls"}`},
		{"RunScript", `{"language":"ruby","code":"1"}`},
		{"Tree", `{"path":"../evil"}`},
		{"Mkdir", `{"path":"../evil"}`},
		{"Mv", `{"src":"../x","dst":"y"}`},
		{"Sed", `{"expression":"","file":"f"}`},
		{"Awk", `{"program":""}`},
		{"Curl", `{"url":"http://u:p@h/"}`},
		{"Curl", `{"url":"ftp://h/"}`},
		{"GitHubRepos", `{}`},
		{"Nope", `{}`},
	} {
		collect(sb.Execute(ctx, tc[0], tc[1]).Text)
	}

	e := ragEngine(t, fakeRag{ready: false})
	collect(e.ragExecute(ctx, "rag_search", `{"query":""}`).Text)
	collect(e.ragExecute(ctx, "rag_bidon", `{}`).Text)
	collect(e.memExecute("u", "mem_search", `{"query":""}`).Text)
	collect(e.memExecute("u", "mem_bidon", `{}`).Text)

	collect(missingArgErr("Read", "file_path"))
	collect(uniformToolError("Read", ToolResult{Text: "[error] open /x: nope"}).Text)
	collect(uniformToolError("Read", ToolResult{Text: "contenu du fichier"}).Text)

	for _, txt := range texts {
		for _, fr := range []string{"[erreur]", "[refuse]", "[avertissement]"} {
			if strings.Contains(txt, fr) {
				t.Errorf("marqueur FR %q dans un resultat d'outil : %q", fr, txt)
			}
		}
	}
}

// TestAgentLoopMessagesAllEnglish : les messages de contrôle de la boucle
// agent (avertissement lectures en chaîne, déduplication) sont injectés au
// modèle — aucun texte FR ne doit y figurer.
func TestAgentLoopMessagesAllEnglish(t *testing.T) {
	const fr = "éèêëàâäîïôöùûüç"
	for _, s := range []string{
		readLoopWarnText,
		repeatedCallResult("prev", 1),
		repeatedCallResult("prev", 3),
		externalOutputMarker,
	} {
		for _, r := range s {
			if strings.ContainsRune(fr, r) {
				t.Errorf("texte FR dans un message de boucle agent : %q", s)
				break
			}
		}
	}
}
