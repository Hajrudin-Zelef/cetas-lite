package chat

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func newTestSandbox(t *testing.T) *Sandbox {
	t.Helper()
	sb, err := NewSandbox(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	return sb
}

func TestToolSchemasCount(t *testing.T) {
	if got := len(ToolSchemas()); got != 9 {
		t.Fatalf("schemas = %d, want 9", got)
	}
}

func TestWriteThenReadPagination(t *testing.T) {
	sb := newTestSandbox(t)
	res := sb.Execute(context.Background(), "Write", `{"file_path":"a/b.txt","content":"l1\nl2\nl3\nl4"}`)
	if strings.HasPrefix(res.Text, "[erreur]") {
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
	if !strings.HasPrefix(got.Text, "[erreur]") {
		t.Fatalf("lecture hors sandbox acceptee: %q", got.Text)
	}
}

func TestEditCases(t *testing.T) {
	sb := newTestSandbox(t)
	sb.Execute(context.Background(), "Write", `{"file_path":"f.txt","content":"hello world"}`)

	ok := sb.Execute(context.Background(), "Edit", `{"file_path":"f.txt","old":"world","new":"terre"}`)
	if strings.HasPrefix(ok.Text, "[erreur]") {
		t.Fatalf("edit: %s", ok.Text)
	}
	again := sb.Execute(context.Background(), "Edit", `{"file_path":"f.txt","old":"world","new":"terre"}`)
	if !strings.Contains(again.Text, "deja") && !strings.Contains(again.Text, "déjà") {
		t.Fatalf("edit deja applique: %q", again.Text)
	}
	missing := sb.Execute(context.Background(), "Edit", `{"file_path":"f.txt","old":"zzz","new":"x"}`)
	if !strings.HasPrefix(missing.Text, "[erreur]") {
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
	if strings.HasPrefix(out.Text, "[erreur]") {
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
	if !strings.Contains(denied.Text, "non autorisee") && !strings.Contains(denied.Text, "non autorisée") {
		t.Fatalf("bash curl doit etre refuse: %q", denied.Text)
	}
	flag := sb.Execute(context.Background(), "Bash", `{"command":"echo -e x"}`)
	if !strings.Contains(flag.Text, "interdit") {
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
	out := sb.Execute(ctx, "Bash", `{"command":"tail -f /dev/null","timeout":60}`)
	if !strings.Contains(out.Text, "interrompue") {
		t.Fatalf("stop doit etre signale: %q", out.Text)
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
	if !strings.HasPrefix(out.Text, "[erreur]") {
		t.Fatalf("outil inconnu: %q", out.Text)
	}
}
