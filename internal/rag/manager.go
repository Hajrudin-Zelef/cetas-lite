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
	emb  atomic.Pointer[Embedder]
	vec  atomic.Pointer[VectorStore]
	rrk  atomic.Pointer[Reranker]
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

// SetEmbedder : branche le client d'embedding (lot 2). Appele une fois au
// demarrage, avant Start. Embedder nil => recherche BM25 seule.
func (m *Manager) SetEmbedder(emb *Embedder) {
	if m == nil {
		return
	}
	if emb == nil {
		return
	}
	m.emb.Store(emb)
}

func (m *Manager) embedder() *Embedder {
	if m == nil {
		return nil
	}
	return m.emb.Load()
}

// SetReranker : branche le client /rerank (plateforme desktop). Nil ou cle
// absente => pas de rerank, l'ordre RRF est conserve (fail-open).
func (m *Manager) SetReranker(r *Reranker) {
	if m == nil {
		return
	}
	if r == nil {
		return
	}
	m.rrk.Store(r)
}

func (m *Manager) reranker() *Reranker {
	if m == nil {
		return nil
	}
	return m.rrk.Load()
}

// RerankReady : un reranker est branche (observabilite / CLI status).
func (m *Manager) RerankReady() bool { return m.reranker() != nil }

func (m *Manager) vectors() *VectorStore {
	if m == nil {
		return nil
	}
	return m.vec.Load()
}

// VectorsReady : des vecteurs alignes avec l'index courant sont charges.
func (m *Manager) VectorsReady() bool { return m.vectors() != nil }

// EmbedModelSlug : modele d'embedding selectionne ("" si aucun).
func (m *Manager) EmbedModelSlug() string {
	if e := m.embedder(); e != nil {
		return e.Model().Slug
	}
	return ""
}

// loadVectors : charge les vecteurs du modele selectionne et verifie leur
// alignement avec l'index (hash du corpus). Absents, illisibles, obsoletes
// ou desalignes => vecteurs desactives, BM25 seul (fail-open, loggue).
func (m *Manager) loadVectors(ix *Index) {
	m.vec.Store(nil)
	emb := m.embedder()
	if emb == nil || ix == nil || !ix.Stats().Ready {
		return
	}
	slug := emb.Model().Slug
	path := VectorFilePath(m.root, slug)
	vs, err := LoadVectorFile(path)
	if err != nil {
		log.Printf("rag: vecteurs %s indisponibles (%v) : BM25 seul", slug, err)
		return
	}
	if want := ix.CorpusHash(); vs.Hash != want {
		log.Printf("rag: vecteurs %s obsoletes (corpus modifie) : BM25 seul, reconstruire via `rag-vectors build`", slug)
		return
	}
	m.vec.Store(vs)
	log.Printf("rag: vecteurs %s charges (%d chunks, dim %d)", slug, vs.Count, vs.Dims)
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
		log.Printf("rag: %d chunks, %d corpus, %d termes, %.1f Mo de texte en %s",
			st.Chunks, st.Corpora, st.Terms, float64(st.Bytes)/1024/1024,
			st.Duration.Round(time.Millisecond))
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
	m.loadVectors(ix)
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

// SearchCorpus : recherche restreinte aux chunks d'un corpus
// (focus au clic sur une question suggeree).
func (m *Manager) SearchCorpus(ctx context.Context, query, corpus string, limit int) Result {
	ix := m.current()
	if ix == nil || !ix.Stats().Ready {
		return Result{}
	}
	if ctx == nil {
		ctx = context.Background()
	}
	cctx, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()
	return ix.SearchCorpus(cctx, query, corpus, limit)
}

// SearchBoosted : recherche avec boost de termes (entites reprises de
// l'historique, iteration 6b). Meme fail-open que Search.
func (m *Manager) SearchBoosted(ctx context.Context, query string, boost map[string]float64, limit int) Result {
	ix := m.current()
	if ix == nil || !ix.Stats().Ready {
		return Result{}
	}
	if ctx == nil {
		ctx = context.Background()
	}
	cctx, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()
	return ix.SearchBoosted(cctx, query, boost, limit)
}

// Read : relit un chunk par chemin (ou nom de base), lignes numerotees.
func (m *Manager) Read(rel string, offset, limit int) (string, error) {
	return m.current().Read(rel, offset, limit)
}
