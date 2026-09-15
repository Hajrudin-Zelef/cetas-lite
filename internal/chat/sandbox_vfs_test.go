package chat

import (
	"context"
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

// TestGitEnv vérifie l'injection du token GitHub pour git.
func TestGitEnv(t *testing.T) {
	if gitEnv("") != nil {
		t.Fatal("token vide -> env attendu nil")
	}
	env := gitEnv("tok123")
	if !strings.Contains(env["GIT_CONFIG_KEY_0"], "tok123@github.com") {
		t.Fatalf("env git: %v", env)
	}
	if env["GIT_TERMINAL_PROMPT"] != "0" {
		t.Fatalf("prompt terminal non désactivé: %v", env)
	}
}
