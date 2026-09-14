// Package vfs fournit une abstraction de système de fichiers utilisable
// par l'agent : disque local ou dossier distant via SFTP. Les outils de
// l'agent (Ls, Read, Write, Edit, Glob, Grep, Bash) travaillent contre
// cette interface, jamais contre os.* directement, de sorte que le mode
// distant se comporte exactement comme le mode local.
package vfs

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"
)

// Entry décrit un fichier ou dossier. Path est relatif à la racine,
// en slashs ; les dossiers se terminent par "/".
type Entry struct {
	Path    string    `json:"path"`
	IsDir   bool      `json:"is_dir"`
	Size    int64     `json:"size"`
	ModTime time.Time `json:"mod_time"`
}

var ErrOutsideRoot = errors.New("chemin hors racine")
var ErrNotFound = errors.New("introuvable")

// FS est un système de fichiers confiné à une racine. Tous les chemins
// sont relatifs (slashs), jamais absolus ; toute sortie de la racine
// est refusée avec ErrOutsideRoot.
type FS interface {
	// Name retourne un libellé d'affichage de la racine.
	Name() string
	// Remote indique si le FS est distant (SFTP) : latence, pas d'exec local.
	Remote() bool
	// Resolve vérifie le confinement et retourne le chemin canonique relatif.
	Resolve(rel string) (string, error)
	ReadFile(ctx context.Context, rel string) ([]byte, error)
	WriteFile(ctx context.Context, rel string, data []byte, perm os.FileMode) error
	MkdirAll(ctx context.Context, rel string) error
	Remove(ctx context.Context, rel string) error
	Stat(ctx context.Context, rel string) (Entry, error)
	ReadDir(ctx context.Context, rel string) ([]Entry, error)
	// Walk parcourt récursivement en profondeur, en ignorant les entrées
	// retournées par SkipEntry. S'arrête proprement sur ctx annulé.
	Walk(ctx context.Context, fn func(Entry) error) error
	// Exec exécute une commande (locale ou via SSH) avec timeout et
	// variables d'environnement additionnelles. Retourne la sortie combinée.
	Exec(ctx context.Context, name string, args []string, env map[string]string, timeout time.Duration) (string, error)
	Close() error
}

// SkipEntry indique si un nom de fichier/dossier doit être ignoré lors
// des parcours (dépendances, VCS, caches, fichiers cachés). Identique pour
// local et distant afin que l'agent voie toujours la même structure logique.
func SkipEntry(name string) bool {
	if name == ".runscript_tmp" || strings.HasPrefix(name, ".") {
		return true
	}
	switch name {
	case "node_modules", "__pycache__", ".venv",
		"venv", "dist", "build", "target",
		".next", ".nuxt", "coverage", ".pytest_cache", ".mypy_cache",
		".terraform", ".serverless":
		return true
	}
	return false
}

// CleanRel nettoie un chemin relatif fourni par l'utilisateur ou le modèle.
// Refuse les chemins absolus et les remontées hors racine.
func CleanRel(rel string) (string, error) {
	return cleanRel(rel)
}
