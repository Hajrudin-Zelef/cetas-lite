package chat

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Sandbox struct {
	root        string
	AllowScript bool
	Isolation   string
}

var errOutsideSandbox = errors.New("chemin hors sandbox")

func NewSandbox(root string) (*Sandbox, error) {
	root = strings.TrimSpace(root)
	if root == "" {
		return nil, errors.New("racine sandbox vide")
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(abs, 0o700); err != nil {
		return nil, fmt.Errorf("creation sandbox: %w", err)
	}
	if real, err := filepath.EvalSymlinks(abs); err == nil {
		abs = real
	}
	return &Sandbox{root: abs}, nil
}

func (s *Sandbox) Root() string { return s.root }

func (s *Sandbox) Resolve(rel string) (string, error) {
	rel = strings.TrimSpace(strings.ReplaceAll(rel, "\\", "/"))
	if rel == "" {
		return "", errors.New("chemin vide")
	}
	if filepath.IsAbs(rel) || strings.HasPrefix(rel, "/") {
		return "", errOutsideSandbox
	}
	cand := filepath.Join(s.root, filepath.FromSlash(rel))

	dir := cand
	for {
		if _, err := os.Lstat(dir); err == nil {
			break
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	realDir, err := filepath.EvalSymlinks(dir)
	if err != nil {
		return "", errOutsideSandbox
	}
	rest, err := filepath.Rel(dir, cand)
	if err != nil {
		return "", errOutsideSandbox
	}
	real := filepath.Join(realDir, rest)
	if real != s.root && !strings.HasPrefix(real, s.root+string(os.PathSeparator)) {
		return "", errOutsideSandbox
	}
	return real, nil
}
