package rag

import (
	"context"
	"log"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Manager detient l'index courant et le (re)charge en tache de fond. Toutes
// les lectures passent par un pointeur atomique : aucun verrou sur le chemin
// chaud, aucune attente au demarrage.
type Manager struct {
	root string
	idx  atomic.Pointer[Index]
	mu   sync.Mutex
}

// New : root vide => manager desactive (Enabled=false), cout nul.
func New(root string) *Manager {
	root = strings.TrimSpace(root)
	if root != "" {
		if abs, err := filepath.Abs(root); err == nil {
			root = abs
		}
	}
	return &Manager{root: root}
}

// Enabled : vrai si un dossier racine est configure.
func (m *Manager) Enabled() bool { return m != nil && m.root != "" }

// Root : dossier racine configure.
func (m *Manager) Root() string {
	if m == nil {
		return ""
	}
	return m.root
}

// Start lance la construction de l'index en arriere-plan. Ne bloque jamais.
func (m *Manager) Start() {
	if !m.Enabled() {
		return
	}
	go func() {
		ix, err := m.Reload()
		if err != nil {
			log.Printf("rag: chargement impossible: %v", err)
			return
		}
		st := ix.Stats()
		if !st.Ready {
			log.Printf("rag: aucun corpus dans %s (inactif)", m.root)
			return
		}
		log.Printf("rag: %d chunks, %d corpus, %d termes en %s",
			st.Chunks, st.Corpora, st.Terms, st.Duration.Round(time.Millisecond))
		for _, e := range st.Errors {
			log.Printf("rag: corpus ignore: %s", e)
		}
	}()
}

// Reload reconstruit l'index (bloquant) et le publie atomiquement.
func (m *Manager) Reload() (*Index, error) {
	if !m.Enabled() {
		ix := emptyIndex()
		m.idx.Store(ix)
		return ix, nil
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	ix, err := Load(m.root)
	if err != nil {
		return nil, err
	}
	m.idx.Store(ix)
	return ix, nil
}

func (m *Manager) current() *Index {
	if m == nil {
		return nil
	}
	return m.idx.Load()
}

// Ready : un index non vide est publie.
func (m *Manager) Ready() bool {
	ix := m.current()
	return ix != nil && ix.Stats().Ready
}

// Stats : etat courant.
func (m *Manager) Stats() Stats {
	if ix := m.current(); ix != nil {
		return ix.Stats()
	}
	return Stats{}
}

// Search : recherche bornee dans le temps (fail-open). Ne bloque jamais.
func (m *Manager) Search(ctx context.Context, query string, limit int) Result {
	ix := m.current()
	if ix == nil || !ix.Stats().Ready {
		return Result{}
	}
	if ctx == nil {
		ctx = context.Background()
	}
	cctx, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()
	return ix.Search(cctx, query, limit)
}

// Read : relit un chunk par chemin (ou nom de base), lignes numerotees.
func (m *Manager) Read(rel string, offset, limit int) (string, error) {
	return m.current().Read(rel, offset, limit)
}
