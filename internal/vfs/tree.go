package vfs

import (
	"context"
	"sort"
	"strings"
)

// Node est un nœud d'arborescence JSON (dossier avec enfants, ou fichier).
type Node struct {
	Name     string  `json:"name"`
	Path     string  `json:"path"`
	IsDir    bool    `json:"is_dir"`
	Size     int64   `json:"size,omitempty"`
	Children []*Node `json:"children,omitempty"`
	// Truncated indique que des enfants ont été coupés (limite atteinte).
	Truncated bool `json:"truncated,omitempty"`
}

// TreeOptions borne l'arborescence pour rester économe.
type TreeOptions struct {
	// MaxDepth : profondeur max sous la racine demandée (0 = racine seule).
	MaxDepth int
	// MaxEntries : nombre total max de nœuds.
	MaxEntries int
}

func DefaultTreeOptions() TreeOptions { return TreeOptions{MaxDepth: 4, MaxEntries: 2000} }

// Tree construit l'arborescence sous rel ("" = racine).
func Tree(ctx context.Context, fsys FS, rel string, opt TreeOptions) (*Node, error) {
	clean := ""
	if strings.TrimSpace(rel) != "" {
		var err error
		clean, err = fsys.Resolve(rel)
		if err != nil {
			return nil, err
		}
	}
	root := &Node{Name: baseName(clean), Path: clean, IsDir: true}
	count := 1
	truncated := false
	var rec func(n *Node, rp string, depth int) error
	rec = func(n *Node, rp string, depth int) error {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if depth >= opt.MaxDepth || count >= opt.MaxEntries {
			if depth >= opt.MaxDepth || count >= opt.MaxEntries {
				truncated = true
			}
			return nil
		}
		ents, err := fsys.ReadDir(ctx, rp)
		if err != nil {
			return nil
		}
		for _, e := range ents {
			if count >= opt.MaxEntries {
				truncated = true
				break
			}
			child := &Node{
				Name:  baseName(strings.TrimSuffix(e.Path, "/")),
				Path:  strings.TrimSuffix(e.Path, "/"),
				IsDir: e.IsDir,
				Size:  e.Size,
			}
			count++
			n.Children = append(n.Children, child)
			if e.IsDir {
				if err := rec(child, strings.TrimSuffix(e.Path, "/"), depth+1); err != nil {
					return err
				}
			}
		}
		return nil
	}
	if err := rec(root, clean, 0); err != nil {
		return nil, err
	}
	root.Truncated = truncated
	return root, nil
}

// FlatList retourne une liste triée "chemin/" pour les dossiers, "chemin"
// pour les fichiers — format compact idéal pour le prompt de l'agent.
func FlatList(ctx context.Context, fsys FS, maxEntries int) ([]string, bool, error) {
	var out []string
	truncated := false
	err := fsys.Walk(ctx, func(e Entry) error {
		if len(out) >= maxEntries {
			truncated = true
			return errWalkStop
		}
		out = append(out, e.Path)
		return nil
	})
	if err != nil && err != errWalkStop {
		return nil, false, err
	}
	sort.Strings(out)
	return out, truncated, nil
}

var errWalkStop = errStop{}

type errStop struct{}

func (errStop) Error() string { return "stop" }

func baseName(p string) string {
	p = strings.TrimSuffix(p, "/")
	if i := strings.LastIndex(p, "/"); i >= 0 {
		return p[i+1:]
	}
	if p == "" {
		return "/"
	}
	return p
}
