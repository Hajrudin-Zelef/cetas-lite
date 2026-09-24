package rag

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
)

// Lot 2 — stockage des vecteurs de chunks.
//
// Format binaire (little-endian), namespacé par modele :
// .vectors/<slug>.bin sous la racine RAG (dossier ignore par le loader,
// qui saute tout ce qui commence par '.').
//
//	 magic "CVEC0001" (8 octets)
//	 slugLen u16 | slug | dims u32 | count u32
//	 corpusHash : SHA256 sur (path + "\x00" + body) de chaque chunk, en ordre
//	 vecteurs : count * dims float32
//
// Le hash garantit l'alignement 1:1 avec l'index : corpus modifie =>
// vecteurs declares obsoletes (fail-open), jamais melanges.

const vectorFileMagic = "CVEC0001"

// VectorsDir : sous-dossier des vecteurs dans la racine RAG.
const VectorsDir = ".vectors"

// VectorFileName : nom du fichier pour un modele.
func VectorFileName(slug string) string { return slug + ".bin" }

// VectorFilePath : chemin complet du fichier de vecteurs d'un modele.
func VectorFilePath(root, slug string) string {
	return filepath.Join(root, VectorsDir, VectorFileName(slug))
}

// VectorStore : vecteurs de documents, normalises (cosinus = produit scalaire).
type VectorStore struct {
	Model EmbedModel
	Dims  int
	Count int
	Hash  [32]byte
	vecs  []float32 // Count * Dims, normalises
}

// LoadVectorFile : lit et valide un fichier de vecteurs.
func LoadVectorFile(path string) (*VectorStore, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	r := &binReader{b: raw}
	if string(r.bytes(8)) != vectorFileMagic {
		return nil, fmt.Errorf("vecteurs: magic invalide")
	}
	slugLen := r.u16()
	slug := string(r.bytes(int(slugLen)))
	dims := int(r.u32())
	count := int(r.u32())
	var hash [32]byte
	copy(hash[:], r.bytes(32))
	if r.err != nil {
		return nil, fmt.Errorf("vecteurs: en-tete tronque: %w", r.err)
	}
	model, ok := LookupEmbedModel(slug)
	if !ok {
		return nil, fmt.Errorf("vecteurs: modele inconnu %q", slug)
	}
	if dims != model.Dims {
		return nil, fmt.Errorf("vecteurs: dims %d != %d (%s)", dims, model.Dims, slug)
	}
	need := count * dims
	if r.remaining() < need*4 {
		return nil, fmt.Errorf("vecteurs: fichier tronque (%d vecteurs attendus)", count)
	}
	vecs := make([]float32, need)
	for i := range vecs {
		vecs[i] = math.Float32frombits(r.u32())
	}
	vs := &VectorStore{Model: model, Dims: dims, Count: count, Hash: hash, vecs: vecs}
	vs.normalize()
	return vs, nil
}

// normalize : vecteurs unitaires pour un cosinus en produit scalaire.
func (vs *VectorStore) normalize() {
	for i := 0; i < vs.Count; i++ {
		v := vs.vecs[i*vs.Dims : (i+1)*vs.Dims]
		var n float64
		for _, x := range v {
			n += float64(x) * float64(x)
		}
		n = math.Sqrt(n)
		if n == 0 {
			continue
		}
		inv := 1 / n
		for j, x := range v {
			v[j] = float32(float64(x) * inv)
		}
	}
}

// SemHit : chunk candidat de la jambe semantique.
type SemHit struct {
	Idx    int     // index du chunk dans l'index
	Cosine float64 // similarite cosinus requete/chunk
}

// TopK : k chunks les plus proches de q (brute-force, 2487x1536 = ~4M
// multiplications : negligeable). Deterministe : egalite => index croissant.
func (vs *VectorStore) TopK(q []float32, k int) []SemHit {
	if vs == nil || vs.Count == 0 || len(q) != vs.Dims || k <= 0 {
		return nil
	}
	qn := normalizeVec(q)
	scores := make([]SemHit, vs.Count)
	for i := 0; i < vs.Count; i++ {
		v := vs.vecs[i*vs.Dims : (i+1)*vs.Dims]
		var dot float64
		for j, x := range v {
			dot += float64(x) * float64(qn[j])
		}
		scores[i] = SemHit{Idx: i, Cosine: dot}
	}
	sort.Slice(scores, func(a, b int) bool {
		if scores[a].Cosine == scores[b].Cosine {
			return scores[a].Idx < scores[b].Idx
		}
		return scores[a].Cosine > scores[b].Cosine
	})
	if k > len(scores) {
		k = len(scores)
	}
	return scores[:k]
}

func normalizeVec(v []float32) []float32 {
	out := make([]float32, len(v))
	var n float64
	for _, x := range v {
		n += float64(x) * float64(x)
	}
	if n == 0 {
		return out
	}
	inv := 1 / math.Sqrt(n)
	for i, x := range v {
		out[i] = float32(float64(x) * inv)
	}
	return out
}

// binReader : lecteur binaire minimaliste.
type binReader struct {
	b   []byte
	off int
	err error
}

func (r *binReader) bytes(n int) []byte {
	if r.err != nil {
		return nil
	}
	if r.off+n > len(r.b) {
		r.err = fmt.Errorf("fin inattendue")
		return nil
	}
	v := r.b[r.off : r.off+n]
	r.off += n
	return v
}

func (r *binReader) u16() uint16 {
	b := r.bytes(2)
	if b == nil {
		return 0
	}
	return binary.LittleEndian.Uint16(b)
}

func (r *binReader) u32() uint32 {
	b := r.bytes(4)
	if b == nil {
		return 0
	}
	return binary.LittleEndian.Uint32(b)
}

func (r *binReader) remaining() int { return len(r.b) - r.off }

// writeVectorFile : serialise le store.
func writeVectorFile(path string, vs *VectorStore) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	var b []byte
	b = append(b, vectorFileMagic...)
	sb := []byte(vs.Model.Slug)
	var tmp [4]byte
	binary.LittleEndian.PutUint16(tmp[:2], uint16(len(sb)))
	b = append(b, tmp[:2]...)
	b = append(b, sb...)
	binary.LittleEndian.PutUint32(tmp[:], uint32(vs.Dims))
	b = append(b, tmp[:]...)
	binary.LittleEndian.PutUint32(tmp[:], uint32(vs.Count))
	b = append(b, tmp[:]...)
	b = append(b, vs.Hash[:]...)
	for _, x := range vs.vecs {
		binary.LittleEndian.PutUint32(tmp[:], math.Float32bits(x))
		b = append(b, tmp[:]...)
	}
	return os.WriteFile(path, b, 0o600)
}

// Limites d'indexation : l'API d'embedding plafonne les entrees (~8191
// tokens pour les modeles OpenAI). On tronque et on lote par budget
// de tokens estimes (len/4), pas par nombre de textes.
const (
	embedMaxChars     = 20000
	embedBatchTokens  = 90000
	embedBatchMaxSize = 96
)

// BuildVectors : construit le fichier de vecteurs d'un modele pour le
// corpus courant. Operation explicite (CLI rag-vectors build) : cout API
// et cle requis. Les documents ne recoivent jamais le QueryPrefix.
func BuildVectors(root string, emb *Embedder, progress func(done, total int)) error {
	if emb == nil {
		return fmt.Errorf("vecteurs: embedder manquant")
	}
	ix, err := Load(root)
	if err != nil {
		return fmt.Errorf("vecteurs: chargement corpus: %w", err)
	}
	if !ix.Stats().Ready {
		return fmt.Errorf("vecteurs: aucun corpus dans %s", root)
	}
	model := emb.Model()
	texts := make([]string, 0, len(ix.chunks))
	for _, c := range ix.chunks {
		t := c.body
		if len(t) > embedMaxChars {
			t = t[:embedMaxChars]
		}
		texts = append(texts, t)
	}
	vecs := make([]float32, 0, len(texts)*model.Dims)
	// BuildVectors est appele depuis la CLI, sans requete parente.
	ctx := context.Background()
	for start := 0; start < len(texts); {
		end := start
		budget := 0
		for end < len(texts) && end-start < embedBatchMaxSize && budget < embedBatchTokens {
			budget += len(texts[end]) / 4
			end++
		}
		if end == start {
			end = start + 1
		}
		batch, err := emb.Embed(ctx, texts[start:end], false)
		if err != nil {
			return fmt.Errorf("vecteurs: lot %d-%d: %w", start, end, err)
		}
		for _, v := range batch {
			vecs = append(vecs, v...)
		}
		if progress != nil {
			progress(end, len(texts))
		}
		start = end
	}
	hash := ix.CorpusHash()
	vs := &VectorStore{Model: model, Dims: model.Dims, Count: len(texts), Hash: hash, vecs: vecs}
	if err := writeVectorFile(VectorFilePath(root, model.Slug), vs); err != nil {
		return fmt.Errorf("vecteurs: ecriture: %w", err)
	}
	return nil
}
