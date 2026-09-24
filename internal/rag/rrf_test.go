package rag

import "testing"

func TestFuseRRFBasic(t *testing.T) {
	// BM25: c0 > c1 ; semantique: c1 > c0, c2 present.
	bm25 := []Ranked{{Idx: 0, Rank: 1}, {Idx: 1, Rank: 2}}
	sem := []Ranked{{Idx: 1, Rank: 1}, {Idx: 0, Rank: 2}, {Idx: 2, Rank: 3}}
	fused := FuseRRF(bm25, sem)
	if len(fused) != 3 {
		t.Fatalf("%d fusionnes", len(fused))
	}
	// c0 et c1 a egalite parfaite (1/61+1/62) : index croissant.
	if fused[0].Idx != 0 || fused[1].Idx != 1 || fused[2].Idx != 2 {
		t.Fatalf("ordre = %v", fused)
	}
	// Score attendu : 1/61 + 1/62.
	want := 1.0/61.0 + 1.0/62.0
	if diff := fused[0].Score - want; diff > 1e-12 || diff < -1e-12 {
		t.Fatalf("score = %f, attendu %f", fused[0].Score, want)
	}
}

func TestFuseRRFSingleLeg(t *testing.T) {
	// Une seule jambe : l'ordre est preserve.
	r := []Ranked{{Idx: 5, Rank: 1}, {Idx: 3, Rank: 2}, {Idx: 9, Rank: 3}}
	fused := FuseRRF(r)
	for i, f := range fused {
		if f.Idx != r[i].Idx {
			t.Fatalf("ordre modifie: %v", fused)
		}
	}
}

func TestFuseRRFDisjoint(t *testing.T) {
	// Classements disjoints : le #1 BM25 reste devant le #1 semantique
	// (1/61 > 1/62), puis les #2.
	bm25 := []Ranked{{Idx: 0, Rank: 1}, {Idx: 1, Rank: 2}}
	sem := []Ranked{{Idx: 2, Rank: 1}, {Idx: 3, Rank: 2}}
	fused := FuseRRF(bm25, sem)
	want := []int{0, 2, 1, 3}
	for i, f := range fused {
		if f.Idx != want[i] {
			t.Fatalf("ordre = %v, attendu %v", fused, want)
		}
	}
}

func TestFuseRRFEmpty(t *testing.T) {
	if len(FuseRRF()) != 0 {
		t.Fatal("fusion vide non vide")
	}
	if len(FuseRRF(nil, nil)) != 0 {
		t.Fatal("jambes vides non vides")
	}
}
