package vfs

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// cleanRel nettoie un chemin relatif et vérifie le confinement.
func cleanRel(rel string) (string, error) {
	rel = strings.TrimSpace(strings.ReplaceAll(rel, "\\", "/"))
	if rel == "" {
		return "", errors.New("chemin vide")
	}
	if strings.HasPrefix(rel, "/") || filepath.IsAbs(rel) {
		return "", ErrOutsideRoot
	}
	clean := filepath.ToSlash(filepath.Clean(rel))
	if clean == "." {
		return "", nil
	}
	if clean == ".." || strings.HasPrefix(clean, "../") || strings.Contains(clean, "/../") {
		return "", ErrOutsideRoot
	}
	return clean, nil
}

// LocalFS est l'implémentation disque local de FS.
type LocalFS struct {
	root string
	name string
}

// NewLocal crée un FS local confiné à root (créé si absent).
func NewLocal(root, name string) (*LocalFS, error) {
	root = strings.TrimSpace(root)
	if root == "" {
		return nil, errors.New("racine vide")
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(abs, 0o700); err != nil {
		return nil, err
	}
	if real, err := filepath.EvalSymlinks(abs); err == nil {
		abs = real
	}
	if name == "" {
		name = abs
	}
	return &LocalFS{root: abs, name: name}, nil
}

func (l *LocalFS) Name() string { return l.name }
func (l *LocalFS) Remote() bool { return false }

// Resolve vérifie le confinement (avec résolution des liens symboliques)
// et retourne le chemin relatif canonique. "" désigne la racine.
func (l *LocalFS) Resolve(rel string) (string, error) {
	if strings.TrimSpace(rel) == "" {
		return "", nil
	}
	clean, err := cleanRel(rel)
	if err != nil {
		return "", err
	}
	if clean == "" {
		return "", nil
	}
	cand := filepath.Join(l.root, filepath.FromSlash(clean))
	// Trouve le plus proche parent existant pour évaluer les symlinks.
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
		return "", ErrOutsideRoot
	}
	rest, err := filepath.Rel(dir, cand)
	if err != nil {
		return "", ErrOutsideRoot
	}
	real := filepath.Join(realDir, rest)
	if real != l.root && !strings.HasPrefix(real, l.root+string(os.PathSeparator)) {
		return "", ErrOutsideRoot
	}
	return clean, nil
}

func (l *LocalFS) abs(rel string) (string, error) {
	clean, err := l.Resolve(rel)
	if err != nil {
		return "", err
	}
	if clean == "" {
		return l.root, nil
	}
	return filepath.Join(l.root, filepath.FromSlash(clean)), nil
}

func (l *LocalFS) ReadFile(ctx context.Context, rel string) ([]byte, error) {
	p, err := l.abs(rel)
	if err != nil {
		return nil, err
	}
	b, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return b, nil
}

func (l *LocalFS) WriteFile(ctx context.Context, rel string, data []byte, perm os.FileMode) error {
	p, err := l.abs(rel)
	if err != nil {
		return err
	}
	if dir := filepath.Dir(p); dir != l.root {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return err
		}
	}
	if fi, err := os.Stat(p); err == nil {
		perm = fi.Mode()
	} else if perm == 0 {
		perm = 0o644
	}
	return os.WriteFile(p, data, perm)
}

func (l *LocalFS) MkdirAll(ctx context.Context, rel string) error {
	p, err := l.abs(rel)
	if err != nil {
		return err
	}
	return os.MkdirAll(p, 0o700)
}

func (l *LocalFS) Remove(ctx context.Context, rel string) error {
	p, err := l.abs(rel)
	if err != nil {
		return err
	}
	if err := os.RemoveAll(p); err != nil {
		return err
	}
	return nil
}

func toEntry(rel string, fi os.FileInfo) Entry {
	rel = filepath.ToSlash(rel)
	e := Entry{Path: rel, IsDir: fi.IsDir(), Size: fi.Size(), ModTime: fi.ModTime()}
	if e.IsDir && !strings.HasSuffix(e.Path, "/") {
		e.Path += "/"
	}
	return e
}

func (l *LocalFS) Stat(ctx context.Context, rel string) (Entry, error) {
	p, err := l.abs(rel)
	if err != nil {
		return Entry{}, err
	}
	fi, err := os.Stat(p)
	if err != nil {
		if os.IsNotExist(err) {
			return Entry{}, ErrNotFound
		}
		return Entry{}, err
	}
	clean, _ := cleanRel(rel)
	if clean == "" {
		return Entry{Path: "", IsDir: true}, nil
	}
	return toEntry(clean, fi), nil
}

func (l *LocalFS) ReadDir(ctx context.Context, rel string) ([]Entry, error) {
	p, err := l.abs(rel)
	if err != nil {
		return nil, err
	}
	des, err := os.ReadDir(p)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	clean, _ := cleanRel(rel)
	out := make([]Entry, 0, len(des))
	for _, d := range des {
		if SkipEntry(d.Name()) {
			continue
		}
		fi, err := d.Info()
		if err != nil {
			continue
		}
		rp := d.Name()
		if clean != "" {
			rp = clean + "/" + d.Name()
		}
		out = append(out, toEntry(rp, fi))
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].IsDir != out[j].IsDir {
			return out[i].IsDir
		}
		return out[i].Path < out[j].Path
	})
	return out, nil
}

func (l *LocalFS) Walk(ctx context.Context, fn func(Entry) error) error {
	return filepath.WalkDir(l.root, func(p string, d fs.DirEntry, err error) error {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if err != nil {
			return nil
		}
		if p == l.root {
			return nil
		}
		if d.IsDir() && SkipEntry(d.Name()) {
			return filepath.SkipDir
		}
		if SkipEntry(d.Name()) {
			return nil
		}
		rel, rerr := filepath.Rel(l.root, p)
		if rerr != nil {
			return nil
		}
		fi, ferr := d.Info()
		if ferr != nil {
			return nil
		}
		return fn(toEntry(rel, fi))
	})
}

// Exec exécute une commande localement avec timeout et env additionnel.
func (l *LocalFS) Exec(ctx context.Context, name string, args []string, env map[string]string, timeout time.Duration) (string, error) {
	if timeout <= 0 {
		timeout = 60 * time.Second
	}
	cctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	cmd := exec.CommandContext(cctx, name, args...)
	cmd.Dir = l.root
	if len(env) > 0 {
		cmd.Env = os.Environ()
		for k, v := range env {
			cmd.Env = append(cmd.Env, k+"="+v)
		}
	}
	out, err := cmd.CombinedOutput()
	if cctx.Err() == context.DeadlineExceeded {
		return string(out), errors.New("delai depasse")
	}
	return string(out), err
}

func (l *LocalFS) Close() error { return nil }

// RootPath expose le chemin local (utile pour git local, terminal).
func (l *LocalFS) RootPath() string { return l.root }
