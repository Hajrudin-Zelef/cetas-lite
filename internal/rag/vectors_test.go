package rag

import (
	"os"
	"path/filepath"
	"testing"
)

// miniCorpus : 3 fichiers markdown, vocabulaires disjoints.
func miniCorpus(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	files := map[string]string{
		"alpha.md": "# Alpha\n\npomme banane cerise fruit rouge",
		"beta.md":  "# Beta\n\nvoiture moteur roue vitesse circuit",
		"gamma.md": "# Gamma\n\npython golang rust langage compilation",
	}
	for name, body := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func TestVectorFileRoundtrip(t *testing.T) {
	dir := miniCorpus(t)
	ix, err := Load(dir)
	if err != nil {
		t.Fatal(err)
	}
	model, _ := LookupEmbedModel("openai-text-embedding-3-small")
	dims := model.Dims
	n := len(ix.chunks)
	vecs := make([]float32, n*dims)
	for i := range vecs {
		vecs[i] = float32(i%7) * 0.13
	}
	vs := &VectorStore{Model: model, Dims: dims, Count: n, Hash: ix.CorpusHash(), vecs: vecs}
	path := filepath.Join(t.TempDir(), "vec.bin")
	if err := writeVectorFile(path, vs); err != nil {
		t.Fatal(err)
	}
	back, err := LoadVectorFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if back.Dims != dims || back.Count != n || back.Hash != vs.Hash {
		t.Fatalf("en-tete corrompu: %+v", back)
	}
	// Les vecteurs sont normalises au chargement : on compare les
	// directions, pas les normes.
	for i := 0; i < n; i++ {
		a := vecs[i*dims : (i+1)*dims]
		b := back.vecs[i*dims : (i+1)*dims]
		var dot, na, nb float64
		for j := range a {
			dot += float64(a[j]) * float64(b[j])
			na += float64(a[j]) * float64(a[j])
			nb += float64(b[j]) * float64(b[j])
		}
		if na == 0 || nb == 0 {
			continue
		}
		cos := dot / (sqrt(na) * sqrt(nb))
		if cos < 0.99999 {
			t.Fatalf("chunk %d: direction alteree (cos=%f)", i, cos)
		}
	}
}

func sqrt(x float64) float64 {
	// Newton, 20 iterations : largement assez pour un test.
	z := x
	for i := 0; i < 20; i++ {
		z = (z + x/z) / 2
	}
	return z
}

func TestVectorFileCorrupt(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bad.bin")
	if err := os.WriteFile(path, []byte("pas un fichier de vecteurs...................."), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadVectorFile(path); err == nil {
		t.Fatal("fichier corrompu accepte")
	}
	if _, err := LoadVectorFile(filepath.Join(t.TempDir(), "absent.bin")); err == nil {
		t.Fatal("fichier absent accepte")
	}
}

func TestCorpusHashStaleness(t *testing.T) {
	dir := miniCorpus(t)
	ix1, err := Load(dir)
	if err != nil {
		t.Fatal(err)
	}
	h1 := ix1.CorpusHash()
	// Modification d'un chunk => hash different.
	f, err := os.OpenFile(filepath.Join(dir, "alpha.md"), os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteString("\npoire"); err != nil {
		t.Fatal(err)
	}
	_ = f.Close()
	ix2, err := Load(dir)
	if err != nil {
		t.Fatal(err)
	}
	if ix2.CorpusHash() == h1 {
		t.Fatal("hash inchange apres modification du corpus")
	}
	// Corpus identique => hash stable.
	ix3, err := Load(dir)
	if err != nil {
		t.Fatal(err)
	}
	if ix3.CorpusHash() != ix2.CorpusHash() {
		t.Fatal("hash instable sur corpus identique")
	}
}

func TestTopKCorrectness(t *testing.T) {
	model := EmbedModel{Slug: "t", OpenRouterID: "t/m", Dims: 3}
	// e1, e2, et un vecteur a 45 degres entre les deux.
	vs := &VectorStore{Model: model, Dims: 3, Count: 3}
	vs.vecs = []float32{
		1, 0, 0,
		0, 1, 0,
		0.7071, 0.7071, 0,
	}
	vs.normalize()
	top := vs.TopK([]float32{1, 0, 0}, 3)
	if len(top) != 3 {
		t.Fatalf("%d resultats", len(top))
	}
	if top[0].Idx != 0 || top[0].Cosine < 0.999 {
		t.Fatalf("top1 = %+v", top[0])
	}
	if top[1].Idx != 2 {
		t.Fatalf("top2 = %+v, attendu idx 2", top[1])
	}
	if top[2].Idx != 1 {
		t.Fatalf("top3 = %+v, attendu idx 1", top[2])
	}
	// Determinisme sur egalite : deux vecteurs identiques.
	vs2 := &VectorStore{Model: model, Dims: 2, Count: 2}
	vs2.vecs = []float32{1, 0, 1, 0}
	vs2.normalize()
	a := vs2.TopK([]float32{0, 1}, 2)
	b := vs2.TopK([]float32{0, 1}, 2)
	if a[0].Idx != b[0].Idx || a[1].Idx != b[1].Idx || a[0].Idx != 0 {
		t.Fatalf("non deterministe: %v vs %v", a, b)
	}
	// Cas limites.
	if vs.TopK([]float32{1, 0}, 2) != nil {
		t.Fatal("dimension incoherente acceptee")
	}
	if vs.TopK([]float32{1, 0, 0}, 0) != nil {
		t.Fatal("k=0 accepte")
	}
	if vs.TopK([]float32{1, 0, 0}, 99)[2].Idx != 1 {
		t.Fatal("k > count mal gere")
	}
}
