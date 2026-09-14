// Package worktree : isolation des agents via des worktrees git.
//
// Chaque agent (ou tour agent avec l'option worktree) travaille dans son
// propre worktree : un checkout isole du depot, partageant les objets git.
// Si le dossier n'est pas un depot git, Ensure retourne une erreur et
// l'appelant retombe sur le workspace normal.
package worktree

import (
	"context"
	"errors"
	"fmt"
	"hash/fnv"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	ErrNotAGitRepo = errors.New("le dossier n'est pas un depot git")
	ErrBadName     = errors.New("nom de worktree invalide")
)

const gitTimeout = 60 * time.Second

// Manager cree et nettoie des worktrees sous root.
type Manager struct {
	mu   sync.Mutex
	root string
}

// NewManager prepare le repertoire racine des worktrees.
func NewManager(root string) (*Manager, error) {
	root = strings.TrimSpace(root)
	if root == "" {
		return nil, errors.New("racine worktree vide")
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(abs, 0o700); err != nil {
		return nil, fmt.Errorf("creation racine worktrees: %w", err)
	}
	if real, err := filepath.EvalSymlinks(abs); err == nil {
		abs = real
	}
	return &Manager{root: abs}, nil
}

// Root retourne la racine des worktrees.
func (m *Manager) Root() string { return m.root }

// runGit execute git avec un timeout.
func runGit(repo string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), gitTimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", args...)
	if repo != "" {
		cmd.Dir = repo
	}
	out, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		return "", errors.New("git: delai depasse")
	}
	if err != nil {
		return string(out), fmt.Errorf("git %s: %w (%s)", strings.Join(args, " "), err, strings.TrimSpace(string(out)))
	}
	return string(out), nil
}

// IsRepo indique si path est dans un depot git.
func IsRepo(path string) bool {
	if strings.TrimSpace(path) == "" {
		return false
	}
	out, err := runGit(path, "rev-parse", "--is-inside-work-tree")
	return err == nil && strings.TrimSpace(out) == "true"
}

func sanitizeName(name string) (string, error) {
	var b strings.Builder
	for _, r := range name {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '-', r == '_', r == '.':
			b.WriteRune(r)
		default:
			b.WriteRune('_')
		}
	}
	s := strings.Trim(b.String(), "._")
	if s == "" || s == "." || s == ".." {
		return "", ErrBadName
	}
	if len(s) > 64 {
		s = s[:64]
	}
	return s, nil
}

func repoSlug(repo string) string {
	abs, err := filepath.Abs(repo)
	if err != nil {
		abs = repo
	}
	h := fnv.New32a()
	_, _ = h.Write([]byte(abs))
	base, err := sanitizeName(filepath.Base(strings.TrimRight(abs, string(filepath.Separator))))
	if err != nil {
		base = "repo"
	}
	return fmt.Sprintf("%s-%08x", base, h.Sum32())
}

// Path retourne le chemin prevu du worktree (sans le creer).
func (m *Manager) Path(repo, name string) (string, error) {
	safe, err := sanitizeName(name)
	if err != nil {
		return "", err
	}
	return filepath.Join(m.root, repoSlug(repo), safe), nil
}

// Ensure cree le worktree s'il n'existe pas et retourne son chemin.
// L'operation est idempotente.
func (m *Manager) Ensure(repo, name string) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !IsRepo(repo) {
		return "", ErrNotAGitRepo
	}
	path, err := m.Path(repo, name)
	if err != nil {
		return "", err
	}
	if m.registeredLocked(repo, path) {
		return path, nil
	}
	// Nettoyer un repertoire residuel non enregistre.
	if _, err := os.Lstat(path); err == nil {
		_ = os.RemoveAll(path)
	}
	if _, err := runGit(repo, "worktree", "add", "--detach", path, "HEAD"); err != nil {
		return "", err
	}
	return path, nil
}

// registeredLocked verifie via git que path est un worktree du depot.
func (m *Manager) registeredLocked(repo, path string) bool {
	out, err := runGit(repo, "worktree", "list", "--porcelain")
	if err != nil {
		return false
	}
	real := path
	if r, err := filepath.EvalSymlinks(path); err == nil {
		real = r
	}
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimSpace(strings.TrimPrefix(line, "worktree "))
		if line == "" || strings.HasPrefix(line, "worktree ") {
			continue
		}
		if line == path || line == real {
			return true
		}
	}
	return false
}

// Remove supprime le worktree (force) puis elague la liste git.
func (m *Manager) Remove(repo, name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	path, err := m.Path(repo, name)
	if err != nil {
		return err
	}
	if IsRepo(repo) {
		_, _ = runGit(repo, "worktree", "remove", "--force", path)
		_, _ = runGit(repo, "worktree", "prune")
	}
	// Filet de securite : supprimer tout residu.
	if _, err := os.Lstat(path); err == nil {
		_ = os.RemoveAll(path)
	}
	// Nettoyer le dossier du depot s'il est vide.
	parent := filepath.Dir(path)
	if entries, err := os.ReadDir(parent); err == nil && len(entries) == 0 {
		_ = os.Remove(parent)
	}
	return nil
}

// Prune elague les worktrees obsoletes d'un depot.
func (m *Manager) Prune(repo string) error {
	if !IsRepo(repo) {
		return ErrNotAGitRepo
	}
	_, err := runGit(repo, "worktree", "prune")
	return err
}
