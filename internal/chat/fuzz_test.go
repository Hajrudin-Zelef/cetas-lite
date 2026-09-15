package chat

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func FuzzSandboxResolve(f *testing.F) {
	root, err := os.MkdirTemp("", "sandbox-fuzz")
	if err != nil {
		f.Fatal(err)
	}
	defer os.RemoveAll(root)
	sb, err := NewSandbox(root)
	if err != nil {
		f.Fatal(err)
	}
	for _, s := range []string{"a.txt", "a/b.txt", "../x", "/etc/passwd", "~/x", "a/../../b", "", ".", "..", "a//b"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, rel string) {
		p, err := sb.Resolve(rel)
		if err != nil {
			return
		}
		// Le chemin résolu doit rester relatif et confiné : pas de
		// remontée, pas d'absolu, pas de segment "..".
		if p == "" {
			return
		}
		if strings.HasPrefix(p, "/") || p == ".." || strings.HasPrefix(p, "../") || strings.Contains(p, "/../") {
			t.Fatalf("evasion sandbox: %q -> %q", rel, p)
		}
	})
}

func FuzzSplitCommand(f *testing.F) {
	for _, s := range []string{
		"ls -la", `echo "a b"`, "cat 'x'", "git commit -m 'hi'", `a\ b`, "", "'", `"`, `a\\`, "a 'b c' d",
	} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		toks, err := splitCommand(s)
		if err != nil {
			return
		}
		for _, tok := range toks {
			_ = filepath.Base(tok)
			_ = strings.TrimSpace(tok)
		}
	})
}
