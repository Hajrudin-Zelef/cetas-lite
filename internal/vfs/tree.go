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
	// Truncated indique que des enfants ont été réellement coupés
	// (plafond MaxEntries atteint). Atteindre la profondeur demandée
	// (MaxDepth) n'est PAS une troncature : c'est le fonctionnement
	// normal, notamment pour le chargement paresseux de l'arborescence
	// côté UI (un niveau par requête, suite au dépliage).
	Truncated bool `json:"truncated,omitempty"`
	// DepthLimited indique que la profondeur demandée a été atteinte :
	// des sous-dossiers peuvent exister plus bas sans avoir été explorés.
	DepthLimited bool `json:"depth_limited,omitempty"`
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
	depthLimited := false
	var rec func(n *Node, rp string, depth int) error
	rec = func(n *Node, rp string, depth int) error {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if depth >= opt.MaxDepth {
			// Palier de profondeur demandé : fonctionnement normal
			// (chargement paresseux), pas une troncature.
			depthLimited = true
			return nil
		}
		if count >= opt.MaxEntries {
			truncated = true
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
	root.DepthLimited = depthLimited
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
