package chat

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestNewSandboxEmptyRoot(t *testing.T) {
	if _, err := NewSandbox("  "); err == nil {
		t.Fatal("racine vide doit echouer")
	}
}

func TestSandboxResolveNestedMissing(t *testing.T) {
	root := t.TempDir()
	sb, err := NewSandbox(root)
	if err != nil {
		t.Fatal(err)
	}
	got, err := sb.Resolve("a/b/c.txt")
	if err != nil {
		t.Fatalf("chemin imbrique doit etre accepte: %v", err)
	}
	want := filepath.Join(root, "a", "b", "c.txt")
	if got != want {
		t.Fatalf("resolve = %q, want %q", got, want)
	}
}

func TestSandboxResolveRejectsTraversal(t *testing.T) {
	sb, err := NewSandbox(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	for _, rel := range []string{"../x", "../../etc/passwd", "a/../../x"} {
		if got, err := sb.Resolve(rel); err == nil {
			t.Fatalf("traversee %q acceptee (%q)", rel, got)
		}
	}
}

func TestSandboxResolveRejectsAbsolute(t *testing.T) {
	sb, err := NewSandbox(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	if got, err := sb.Resolve("/etc/passwd"); err == nil {
		t.Fatalf("chemin absolu accepte (%q)", got)
	}
}

func TestSandboxResolveRejectsSymlinkEscape(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("symlinks non fiables sous windows")
	}
	root := t.TempDir()
	outside := t.TempDir()
	if err := os.Symlink(outside, filepath.Join(root, "link")); err != nil {
		t.Skipf("symlink indisponible: %v", err)
	}
	sb, err := NewSandbox(root)
	if err != nil {
		t.Fatal(err)
	}
	if got, err := sb.Resolve("link/secret.txt"); err == nil {
		t.Fatalf("symlink sortant accepte (%q)", got)
	}
}

func TestUserWorkspacePathSanitizes(t *testing.T) {
	base := t.TempDir()
	got, err := userWorkspacePath(base, "../../etc")
	if err != nil {
		t.Fatalf("chemin utilisateur doit etre assaini: %v", err)
	}
	if !strings.HasPrefix(got, base+string(os.PathSeparator)) {
		t.Fatalf("workspace utilisateur hors base: %q", got)
	}
	for _, bad := range []string{"", "..", ".", "   "} {
		if _, err := userWorkspacePath(base, bad); err == nil {
			t.Fatalf("utilisateur %q doit etre refuse", bad)
		}
	}
	if _, err := userWorkspacePath("", "sam"); err == nil {
		t.Fatal("base vide doit etre refusee")
	}
}

func TestSandboxResolveAllowsInsideSymlink(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("symlinks non fiables sous windows")
	}
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "real"), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Join(root, "real"), filepath.Join(root, "link")); err != nil {
		t.Skipf("symlink indisponible: %v", err)
	}
	sb, err := NewSandbox(root)
	if err != nil {
		t.Fatal(err)
	}
	got, err := sb.Resolve("link/f.txt")
	if err != nil {
		t.Fatalf("symlink interne refuse: %v", err)
	}
	if !strings.HasPrefix(got, root+string(os.PathSeparator)) {
		t.Fatalf("chemin resolu hors racine: %q", got)
	}
}
