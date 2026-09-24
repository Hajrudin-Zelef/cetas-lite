package rag

import "sort"

// Lot 2 — fusion de classements par Reciprocal Rank Fusion.
//
// score(idx) = Σ 1 / (k + rang) sur chaque classement contenant idx.
// RRF ne demande aucune normalisation des scores (BM25 et cosinus vivent
// sur des echelles differentes) et reste stable quand une jambe est
// absente : le classement fusionne se reduit alors a l'autre jambe.

// rrfK : constante standard de la litterature.
const rrfK = 60.0

// Ranked : un candidat d'une jambe de recherche (BM25 ou semantique).
type Ranked struct {
	Idx   int // index du chunk dans l'index
	Rank  int // rang 1-based dans sa jambe
	Score float64
}

// FuseRRF : fusionne les classements, ordre decroissant de score fusionne.
// Deterministe : egalite => index de chunk croissant.
func FuseRRF(rankings ...[]Ranked) []Ranked {
	agg := map[int]*Ranked{}
	for _, r := range rankings {
		for _, c := range r {
			a, ok := agg[c.Idx]
			if !ok {
				a = &Ranked{Idx: c.Idx}
				agg[c.Idx] = a
			}
			a.Score += 1.0 / (rrfK + float64(c.Rank))
		}
	}
	out := make([]Ranked, 0, len(agg))
	for _, a := range agg {
		out = append(out, *a)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Score == out[j].Score {
			return out[i].Idx < out[j].Idx
		}
		return out[i].Score > out[j].Score
	})
	return out
}
