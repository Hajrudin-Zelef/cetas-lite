package memory

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func FuzzMemoryName(f *testing.F) {
	root, err := os.MkdirTemp("", "memory-fuzz")
	if err != nil {
		f.Fatal(err)
	}
	defer os.RemoveAll(root)
	s := New(root)
	for _, seed := range []string{"a.md", "a", "../x", "/etc/passwd", "MEMORY.md", "a/b.md", "", ".", "..", "a b.md"} {
		f.Add("u", seed)
	}
	f.Fuzz(func(t *testing.T, user, name string) {
		dir, err := s.userDir(user)
		if err != nil {
			return
		}
		p, err := s.pagePath(user, name)
		if err != nil {
			return
		}
		if p != dir && !strings.HasPrefix(p, dir+string(os.PathSeparator)) {
			t.Fatalf("evasion memoire: user=%q name=%q -> %q", user, name, p)
		}
		if strings.ContainsAny(filepath.Base(p), `/\`) {
			t.Fatalf("nom avec separateur: %q", p)
		}
	})
}
