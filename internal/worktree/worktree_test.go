package worktree

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func initRepo(t *testing.T) string {
	t.Helper()
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git indisponible")
	}
	dir := t.TempDir()
	for _, args := range [][]string{
		{"init"},
		{"config", "user.email", "test@example.com"},
		{"config", "user.name", "test"},
	} {
		cmd := exec.Command("git", args...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v (%s)", args, err, out)
		}
	}
	if err := os.WriteFile(filepath.Join(dir, "f.txt"), []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{{"add", "."}, {"commit", "-m", "init"}} {
		cmd := exec.Command("git", args...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v (%s)", args, err, out)
		}
	}
	return dir
}

func TestEnsureAndRemove(t *testing.T) {
	repo := initRepo(t)
	m, err := NewManager(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	path, err := m.Ensure(repo, "agent-1")
	if err != nil {
		t.Fatalf("ensure: %v", err)
	}
	if _, err := os.Stat(filepath.Join(path, "f.txt")); err != nil {
		t.Fatalf("fichier absent du worktree: %v", err)
	}
	// Idempotence.
	path2, err := m.Ensure(repo, "agent-1")
	if err != nil || path2 != path {
		t.Fatalf("ensure idempotent: %v %q vs %q", err, path, path2)
	}
	if err := m.Remove(repo, "agent-1"); err != nil {
		t.Fatalf("remove: %v", err)
	}
	if _, err := os.Lstat(path); !os.IsNotExist(err) {
		t.Fatal("le worktree devrait etre supprime")
	}
}

func TestNotAGitRepo(t *testing.T) {
	m, err := NewManager(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := m.Ensure(t.TempDir(), "x"); err != ErrNotAGitRepo {
		t.Fatalf("attendu ErrNotAGitRepo, recu: %v", err)
	}
	if IsRepo(t.TempDir()) {
		t.Fatal("un dossier vide n'est pas un depot")
	}
}

func TestSanitize(t *testing.T) {
	m, err := NewManager(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	p, err := m.Path("/repo", "../../evil")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(p, "..") {
		t.Fatalf("traversee possible: %s", p)
	}
	if _, err := m.Path("/repo", ""); err != ErrBadName {
		t.Fatalf("nom vide: attendu ErrBadName, recu %v", err)
	}
}

func TestIsolationBetweenWorktrees(t *testing.T) {
	repo := initRepo(t)
	m, err := NewManager(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	p1, err := m.Ensure(repo, "a1")
	if err != nil {
		t.Fatal(err)
	}
	p2, err := m.Ensure(repo, "a2")
	if err != nil {
		t.Fatal(err)
	}
	if p1 == p2 {
		t.Fatal("deux worktrees doivent avoir des chemins differents")
	}
	if err := os.WriteFile(filepath.Join(p1, "w1.txt"), []byte("1"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(p2, "w1.txt")); !os.IsNotExist(err) {
		t.Fatal("les worktrees doivent etre isoles")
	}
	_ = m.Remove(repo, "a1")
	_ = m.Remove(repo, "a2")
}
