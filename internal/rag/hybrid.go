package rag

import (
	"context"
	"time"
)

// Lot 2 — recherche hybride : BM25 (mots exacts, entites) + semantique
// (embeddings, similarite cosinus), fusion RRF.
//
// La jambe semantique est strictement optionnelle. Vecteurs absents,
// obsoletes ou desalignes, embedder absent (pas de cle OpenRouter),
// erreur ou timeout d'embedding => repli pur BM25, comportement inchange
// (fail-open). Quand elle est active, la porte de couverture devient :
// BM25 fort OU cosinus fort (voir chat.ragCovered).

const (
	// hybridCandidates : candidats par jambe avant fusion RRF.
	hybridCandidates = 20
	// hybridEmbedTimeout : budget de l'embedding de la requete.
	// Au-dela, la jambe semantique est ignoree (fail-open).
	hybridEmbedTimeout = 2 * time.Second
)

// SearchHybrid : BM25 (+boost optionnel) et recherche semantique en
// parallele, fusion RRF, top `limit`. Sans jambe semantique, strictement
// equivalent a SearchBoosted (memes hits, memes scores BM25).
func (m *Manager) SearchHybrid(ctx context.Context, query string, boost map[string]float64, limit int) Result {
	ix := m.current()
	if ix == nil || !ix.Stats().Ready {
		return Result{}
	}
	if ctx == nil {
		ctx = context.Background()
	}
	vec := m.vectors()
	emb := m.embedder()
	if vec == nil || emb == nil {
		cctx, cancel := context.WithTimeout(ctx, DefaultTimeout)
		defer cancel()
		return ix.SearchBoosted(cctx, query, boost, limit)
	}
	terms := uniqueTerms(tokenize(cleanQueryForSearch(query)))
	base := Result{Total: len(ix.chunks), Corpora: ix.stats.Corpora, Ready: ix.stats.Ready}
	if len(terms) == 0 {
		return base
	}
	candN := limit * 2
	if candN < hybridCandidates {
		candN = hybridCandidates
	}
	if candN > 60 {
		candN = 60
	}

	// Les deux jambes tournent en parallele, chacune bornee.
	type bm25Out struct{ ranked []rankedChunk }
	bm25Ch := make(chan bm25Out, 1)
	go func() {
		cctx, cancel := context.WithTimeout(ctx, DefaultTimeout)
		defer cancel()
		bm25Ch <- bm25Out{ranked: ix.rankChunks(cctx, terms, "", boost)}
	}()
	type semOut struct {
		hits []SemHit
	}
	semCh := make(chan semOut, 1)
	go func() {
		cctx, cancel := context.WithTimeout(ctx, hybridEmbedTimeout)
		defer cancel()
		qv, err := emb.Embed(cctx, []string{query}, true)
		if err != nil || len(qv) == 0 {
			semCh <- semOut{}
			return
		}
		semCh <- semOut{hits: vec.TopK(qv[0], candN)}
	}()

	b := <-bm25Ch
	s := <-semCh

	bm25ByIdx := make(map[int]float64, len(b.ranked))
	bm25Ranked := make([]Ranked, 0, candN)
	for i, rc := range b.ranked {
		bm25ByIdx[rc.idx] = rc.score
		if i < candN {
			bm25Ranked = append(bm25Ranked, Ranked{Idx: rc.idx, Rank: i + 1, Score: rc.score})
		}
	}
	semRanked := make([]Ranked, 0, len(s.hits))
	for i, sh := range s.hits {
		semRanked = append(semRanked, Ranked{Idx: sh.Idx, Rank: i + 1, Score: sh.Cosine})
	}

	res := base
	res.Hybrid = len(semRanked) > 0
	if len(b.ranked) > 0 {
		res.BM25Top = b.ranked[0].score
	}
	if len(s.hits) > 0 {
		res.SemTop = s.hits[0].Cosine
	}
	fused := FuseRRF(bm25Ranked, semRanked)
	limit = clampLimit(limit)
	if len(fused) > limit {
		fused = fused[:limit]
	}
	res.Hits = make([]Hit, 0, len(fused))
	for rank, f := range fused {
		// Hit.Score garde sa semantique historique : le score BM25 du
		// chunk (0 si la jambe BM25 ne l'a pas vu). L'ordre vient de RRF.
		res.Hits = append(res.Hits, ix.hitFor(f.Idx, terms, bm25ByIdx[f.Idx], rank))
	}
	return res
}
