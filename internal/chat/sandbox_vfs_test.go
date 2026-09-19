package chat

import (
	"context"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"cetas-lite/internal/vfs"
)

// TestSandboxVFSTools vérifie que les outils fichiers fonctionnent sur
// l'abstraction VFS (locale ici) : Ls/Read/Write/Edit/Glob/Grep.
func TestSandboxVFSTools(t *testing.T) {
	ctx := context.Background()
	lfs, err := vfs.NewLocal(filepath.Join(t.TempDir(), "ws"), "test")
	if err != nil {
		t.Fatal(err)
	}
	sb := NewSandboxFS(lfs)

	if r := sb.Execute(ctx, "Write", `{"file_path":"src/main.go","content":"package main"}`); !strings.HasPrefix(r.Text, "[ok]") {
		t.Fatalf("write: %s", r.Text)
	}
	if r := sb.Execute(ctx, "Read", `{"file_path":"src/main.go"}`); !strings.Contains(r.Text, "package main") {
		t.Fatalf("read: %s", r.Text)
	}
	if r := sb.Execute(ctx, "Edit", `{"file_path":"src/main.go","old":"package main","new":"package app"}`); !strings.Contains(r.Text, "[ok]") {
		t.Fatalf("edit: %s", r.Text)
	}
	if r := sb.Execute(ctx, "Ls", `{}`); !strings.Contains(r.Text, "src/main.go") {
		t.Fatalf("ls: %s", r.Text)
	}
	if r := sb.Execute(ctx, "Glob", `{"pattern":"**/*.go"}`); !strings.Contains(r.Text, "src/main.go") {
		t.Fatalf("glob: %s", r.Text)
	}
	if r := sb.Execute(ctx, "Grep", `{"pattern":"package app"}`); !strings.Contains(r.Text, "src/main.go:1") {
		t.Fatalf("grep: %s", r.Text)
	}
	// Confinement : sortie de racine refusée.
	if r := sb.Execute(ctx, "Read", `{"file_path":"../evil"}`); !strings.Contains(r.Text, "[erreur]") {
		t.Fatalf("confinement: %s", r.Text)
	}
	if r := sb.Execute(ctx, "Read", `{"file_path":"/abs"}`); !strings.Contains(r.Text, "[erreur]") {
		t.Fatalf("absolu: %s", r.Text)
	}
}

// TestWorkspaceSnapshotMessage vérifie le snapshot injecté dans le prompt.
func TestWorkspaceSnapshotMessage(t *testing.T) {
	ctx := context.Background()
	lfs, err := vfs.NewLocal(filepath.Join(t.TempDir(), "ws"), "test")
	if err != nil {
		t.Fatal(err)
	}
	_ = lfs.WriteFile(ctx, "README.md", []byte("# hi"), 0o644)
	sb := NewSandboxFS(lfs)
	msg := workspaceSnapshotMessage(ctx, sb, "Demo")
	if msg.Role != "system" {
		t.Fatal("rôle inattendu")
	}
	c := msg.Content.(string)
	for _, want := range []string{"WORKSPACE:", "Demo", "README.md", "NEVER guess", "relative"} {
		if !strings.Contains(c, want) {
			t.Fatalf("snapshot sans %q:\n%s", want, c)
		}
	}
	if !strings.Contains(c, "[local]") {
		t.Fatalf("mode local attendu:\n%s", c)
	}
}

// TestGitEnv vérifie l'injection du token GitHub pour git : le token
// transite par http.extraHeader, jamais dans la clé de configuration
// (un `git config --list` ne l'exposait plus autrement).
func TestGitEnv(t *testing.T) {
	if gitEnv("") != nil {
		t.Fatal("token vide -> env attendu nil")
	}
	env := gitEnv("tok123")
	if env["GIT_CONFIG_KEY_0"] != "http.https://github.com/.extraheader" {
		t.Fatalf("cle extraheader attendue, got: %v", env["GIT_CONFIG_KEY_0"])
	}
	if strings.Contains(env["GIT_CONFIG_KEY_0"], "tok123") {
		t.Fatalf("le token ne doit pas figurer dans la cle: %v", env["GIT_CONFIG_KEY_0"])
	}
	if !strings.Contains(env["GIT_CONFIG_VALUE_0"], "tok123") {
		t.Fatalf("token absent de la valeur extraheader: %v", env)
	}
	if env["GIT_TERMINAL_PROMPT"] != "0" {
		t.Fatalf("prompt terminal non désactivé: %v", env)
	}
}

// TestGitShowsConfig vérifie la détection des commandes git qui affichent
// la configuration (sortie expurgée du token).
func TestGitShowsConfig(t *testing.T) {
	yes := [][]string{
		{"git", "config", "--list"},
		{"git", "config", "--show-origin", "--list"},
		{"git", "config", "-l"},
		{"git", "var", "-l"},
		{"git", "-c", "user.name=x", "config", "--list"},
		{"git", "--no-pager", "config", "-l"},
	}
	for _, toks := range yes {
		if !gitShowsConfig("git", toks) {
			t.Errorf("gitShowsConfig(%v) = false, want true", toks)
		}
	}
	no := [][]string{
		{"git", "status"},
		{"git", "config", "--get", "user.name"},
		{"git", "log", "--oneline"},
		{"git", "-c", "user.name=x", "status"},
	}
	for _, toks := range no {
		if gitShowsConfig("git", toks) {
			t.Errorf("gitShowsConfig(%v) = true, want false", toks)
		}
	}
	if gitShowsConfig("ls", []string{"ls", "config", "--list"}) {
		t.Error("gitShowsConfig avec binaire ls = true, want false")
	}
}

// TestGitConfigListRedactsToken vérifie de bout en bout que `git config
// --list` ne fait jamais remonter le token au modèle.
func TestGitConfigListRedactsToken(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git indisponible")
	}
	sb := newTestSandbox(t)
	sb.GitHubToken = func() string { return "SECRET123TOKEN" }
	out := sb.Execute(context.Background(), "Bash", `{"command":"git config --list"}`)
	if strings.Contains(out.Text, "SECRET123TOKEN") {
		t.Fatalf("token en clair dans la sortie: %q", out.Text)
	}
	// Si git a pris en compte l'extraHeader, la redaction doit apparaître.
	if strings.Contains(out.Text, "extraheader") && !strings.Contains(out.Text, "***") {
		t.Fatalf("redaction attendue (***): %q", out.Text)
	}
}
