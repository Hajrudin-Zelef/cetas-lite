// Pre-recuperation des sujets lies (prefetch continu, version maigre).
//
// Apres chaque tour reussi, le moteur pre-recupere en arriere-plan la
// recherche RAG des sujets lies probables (voir rag.RelatedTopicQueries)
// et les met en depot borne. Au tour suivant, si la requete nettoyee
// correspond exactement a un sujet pre-recupere, la recuperation est
// servie depuis le depot au lieu d'etre rejouee (latence economisee :
// typiquement la jambe semantique, 0,5 a 1,8 s en production).
//
// Garanties (devise : stabilite d'abord) :
//   - Seule la recuperation est pre-calculee, jamais de generation ;
//     aucun texte n'est envoye au modele par ce mecanisme.
//   - Reutilisation exacte uniquement : la requete du tour suivant doit
//     nettoyer a la meme cle, et le boost d'entites doit etre nul des
//     deux cotes (le pre-calcul tourne sans boost). Sinon, chemin normal.
//   - Fail-open total : RAG inactif, aucune anticipation, job deja en
//     cours, erreur, timeout => le tour suit le chemin normal, sans
//     erreur visible.
//   - Borne : 3 sujets max par tour, depot de 16 entrees, TTL 10 min,
//     un seul job de fond a la fois (single-flight), timeout global.
//   - Non bloquant : le declenchement ne retarde jamais la reponse.
package chat

import (
	"context"
	"log/slog"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"cetas-lite/internal/rag"
)

const (
	// relPrefetchMaxQueries : sujets lies pre-recuperes par tour reussi.
	// Borne aussi le cout d'embeddings speculatifs (un appel par sujet).
	relPrefetchMaxQueries = 3
	// relPrefetchMaxEntries : borne du depot (au-dela, les plus anciennes
	// sont evincees).
	relPrefetchMaxEntries = 16
	// relPrefetchTTL : duree de vie d'une entree servable.
	relPrefetchTTL = 10 * time.Minute
	// relPrefetchTimeout : borne totale du job de fond.
	relPrefetchTimeout = 30 * time.Second
)

// relPrefetchEnabled : interrupteur (defaut actif). Desactive par
// CETAS_LITE_RELATED_PREFETCH=0|false|no|off.
var relPrefetchEnabled = relPrefetchEnabledFromEnv()

func relPrefetchEnabledFromEnv() bool {
	if raw := strings.TrimSpace(os.Getenv("CETAS_LITE_RELATED_PREFETCH")); raw != "" {
		switch strings.ToLower(raw) {
		case "0", "false", "no", "off":
			return false
		}
	}
	return true
}

// relPrefetchEntry : une recuperation pre-calculee.
type relPrefetchEntry struct {
	res       rag.Result
	createdAt time.Time
}

// relPrefetchStore : depot borne d'entrees, adresse par cle de requete
// nettoyee (rag.CleanQueryKey).
type relPrefetchStore struct {
	mu      sync.Mutex
	entries map[string]*relPrefetchEntry
	order   []string // insertion, pour l'eviction des plus anciennes
}

func newRelPrefetchStore() *relPrefetchStore {
	return &relPrefetchStore{entries: map[string]*relPrefetchEntry{}}
}

// put enregistre une recuperation (premier arrive gagne, TTL a la lecture).
func (s *relPrefetchStore) put(key string, res rag.Result) {
	if key == "" {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.entries[key]; ok {
		return
	}
	for len(s.order) >= relPrefetchMaxEntries {
		oldest := s.order[0]
		s.order = s.order[1:]
		delete(s.entries, oldest)
	}
	s.entries[key] = &relPrefetchEntry{res: res, createdAt: time.Now()}
	s.order = append(s.order, key)
}

// get rend une entree servable : presente et non expiree. Une entree
// expiree est supprimee et compte comme un echec (chemin normal).
func (s *relPrefetchStore) get(key string) (rag.Result, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.entries[key]
	if !ok {
		return rag.Result{}, false
	}
	if time.Since(e.createdAt) > relPrefetchTTL {
		delete(s.entries, key)
		for i, k := range s.order {
			if k == key {
				s.order = append(s.order[:i], s.order[i+1:]...)
				break
			}
		}
		return rag.Result{}, false
	}
	return e.res, true
}

// relPrefetchState : job de fond (single-flight) et depot par moteur.
// Valeur zero utilisable.
type relPrefetchState struct {
	busy  atomic.Int32 // 1 = un job de fond est en cours
	mu    sync.Mutex
	store *relPrefetchStore
}

func (s *relPrefetchState) getStore() *relPrefetchStore {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.store == nil {
		s.store = newRelPrefetchStore()
	}
	return s.store
}

func (s *relPrefetchState) tryAcquire() bool {
	return s.busy.CompareAndSwap(0, 1)
}

func (s *relPrefetchState) release() {
	s.busy.Store(0)
}

// maybeRelatedPrefetch declenche en arriere-plan la pre-recuperation des
// sujets lies probables, apres un tour reussi. Non bloquant, fail-open :
// ne fait rien si le mecanisme est desactive, le RAG inactif, la requete
// vide, un job deja en cours, ou aucune anticipation.
func (e *Engine) maybeRelatedPrefetch(in TurnInput, ragRes rag.Result) {
	if !relPrefetchEnabled {
		return
	}
	rt := e.ragTools()
	if rt == nil || !rt.Ready() {
		return
	}
	if strings.TrimSpace(in.Text) == "" {
		return
	}
	if !e.relPrefetch.tryAcquire() {
		slog.Info("rag_prefetch", "event", "skipped_busy")
		return
	}
	queries := rag.RelatedTopicQueries(in.Text, ragRes, relPrefetchMaxQueries)
	if len(queries) == 0 {
		e.relPrefetch.release()
		return
	}
	go func() {
		defer e.relPrefetch.release()
		ctx, cancel := context.WithTimeout(context.Background(), relPrefetchTimeout)
		defer cancel()
		store := e.relPrefetch.getStore()
		stored := 0
		start := time.Now()
		for _, q := range queries {
			if ctx.Err() != nil {
				break
			}
			// Meme appel que le chemin normal sans boost (la
			// reutilisation exige un boost nul des deux cotes) : le
			// resultat servi est identique a une recuperation live.
			res := rt.SearchHybrid(ctx, q, nil, ragContextHits)
			if len(res.Hits) > 0 {
				store.put(q, res)
				stored++
			}
		}
		// Observabilite (mesure avant toute augmentation, directive
		// RAG) : sujets anticipes, entrees stockees, duree. Le taux de
		// reutilisation se lit avec les lignes event=reuse.
		slog.Info("rag_prefetch", "event", "prefetch_done",
			"queries", queries, "stored", stored,
			"ms", time.Since(start).Milliseconds())
	}()
}

// relPrefetchLookup tente de servir une recuperation pre-calculee pour la
// requete (deja enrichie) du tour courant. N'est appele que quand le boost
// d'entites est nul, comme au pre-calcul : le resultat servi est alors
// identique a celui du chemin normal.
func (e *Engine) relPrefetchLookup(q string) (rag.Result, bool) {
	if !relPrefetchEnabled {
		return rag.Result{}, false
	}
	key := rag.CleanQueryKey(q)
	if key == "" {
		return rag.Result{}, false
	}
	res, ok := e.relPrefetch.getStore().get(key)
	if ok {
		slog.Info("rag_prefetch", "event", "reuse", "key", key, "hits", len(res.Hits))
	}
	return res, ok
}
