package chat

import (
	"context"
	"errors"
	"strings"

	"cetas-lite/internal/vfs"
)

// Sandbox confine les outils fichiers de l'agent à un système de fichiers
// (local ou distant via SFTP). Tous les chemins manipulés sont relatifs.
type Sandbox struct {
	fs vfs.FS
	// GitHubToken, si non nil, fournit le token GitHub connecté pour
	// authentifier les opérations git (commit/diff/push) sans toucher
	// aux fichiers de configuration du dépôt.
	GitHubToken func() string

	AllowScript bool
	Isolation   string
}

var errOutsideSandbox = vfs.ErrOutsideRoot

// NewSandbox crée un sandbox local confiné à root (créé si absent).
func NewSandbox(root string) (*Sandbox, error) {
	lfs, err := vfs.NewLocal(root, "")
	if err != nil {
		return nil, err
	}
	return &Sandbox{fs: lfs}, nil
}

// NewSandboxFS crée un sandbox sur un FS déjà ouvert (local ou SFTP).
func NewSandboxFS(fsys vfs.FS) *Sandbox {
	return &Sandbox{fs: fsys}
}

// FS expose le système de fichiers sous-jacent.
func (s *Sandbox) FS() vfs.FS { return s.fs }

// Remote indique si le sandbox pointe vers un dossier distant.
func (s *Sandbox) Remote() bool { return s.fs.Remote() }

// Root retourne un libellé d'affichage de la racine.
func (s *Sandbox) Root() string { return s.fs.Name() }

// localRoot retourne le chemin local, ou "" si distant.
func (s *Sandbox) localRoot() string {
	if lfs, ok := s.fs.(*vfs.LocalFS); ok {
		return lfs.RootPath()
	}
	return ""
}

// Resolve vérifie le confinement et retourne le chemin relatif canonique.
// (Le chemin absolu n'est plus exposé car il n'a pas de sens en distant.)
func (s *Sandbox) Resolve(rel string) (string, error) {
	return s.fs.Resolve(rel)
}

// ReadFile lit un fichier via le FS (local ou distant).
func (s *Sandbox) ReadFile(ctx context.Context, rel string) ([]byte, error) {
	clean, err := s.fs.Resolve(rel)
	if err != nil {
		return nil, err
	}
	if clean == "" {
		return nil, errors.New("chemin vide")
	}
	return s.fs.ReadFile(ctx, clean)
}

// githubToken retourne le token GitHub connecté, ou "".
func (s *Sandbox) githubToken() string {
	if s.GitHubToken == nil {
		return ""
	}
	return strings.TrimSpace(s.GitHubToken())
}

// gitEnv construit l'environnement git pour authentifier les opérations
// GitHub via le token, sans modifier aucun fichier du dépôt.
// Le token transite par http.extraHeader (jamais dans l'URL) : un
// `git config --list` ne l'expose plus via la clé de configuration.
func gitEnv(token string) map[string]string {
	if token == "" {
		return nil
	}
	return map[string]string{
		"GIT_CONFIG_COUNT":    "1",
		"GIT_CONFIG_KEY_0":    "http.https://github.com/.extraheader",
		"GIT_CONFIG_VALUE_0":  "Authorization: Bearer " + token,
		"GIT_TERMINAL_PROMPT": "0",
	}
}
