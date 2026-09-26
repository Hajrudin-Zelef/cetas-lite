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
//	 magic "CVEC0002" (8 octets)
//	 slugLen u16 | slug | dims u32 | count u32
//	 chunkerVersionLen u8 | chunkerVersion
//	 corpusHash : SHA256 sur (path + "\x00" + body) de chaque chunk, en ordre
//	 vecteurs : count * dims float32
//
// Le hash garantit l'alignement 1:1 avec l'index : corpus modifie =>
// vecteurs declares obsoletes (fail-open), jamais melanges. La version du
// chunker couvre les changements de decoupage (meme contenu source, chunks
// differents) : mismatch => obsoletes.

const vectorFileMagic = "CVEC0002"

// ChunkerVersion : version du pipeline de decoupage. A incrementer quand
// build_rag.py change la maniere de produire les chunks (meme source, chunks
// differents). Le serveur desktop rejoue le meme build_rag.py : un mismatch
// declare les vecteurs obsoletes plutot que de les melanger a l'index courant.
const ChunkerVersion = "build_rag_v1"

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
	Model          EmbedModel
	Dims           int
	Count          int
	Hash           [32]byte
	ChunkerVersion string
	vecs           []float32 // Count * Dims, normalises
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
	chunker := string(r.bytes(int(r.u8())))
	var hash [32]byte
	copy(hash[:], r.bytes(32))
	if r.err != nil {
		return nil, fmt.Errorf("vecteurs: en-tete tronque: %w", r.err)
	}
	if chunker != ChunkerVersion {
		return nil, fmt.Errorf("vecteurs: version chunker %q != %q", chunker, ChunkerVersion)
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
	vs := &VectorStore{Model: model, Dims: dims, Count: count, Hash: hash, ChunkerVersion: chunker, vecs: vecs}
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

func (r *binReader) u8() uint8 {
	b := r.bytes(1)
	if b == nil {
		return 0
	}
	return b[0]
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
	cv := vs.ChunkerVersion
	if cv == "" {
		cv = ChunkerVersion
	}
	b = append(b, byte(len(cv)))
	b = append(b, cv...)
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
	// Plateforme desktop (cetasrag) : 8 textes maximum par requete /embed.
	desktopEmbedBatchSize   = 8
	desktopEmbedBatchTokens = 90000
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
	// La plateforme desktop (cetasrag) plafonne a 8 textes par requete :
	// le backend desktop impose donc sa propre borne.
	maxBatch := embedBatchMaxSize
	batchTokens := embedBatchTokens
	if emb.Desktop() {
		maxBatch = desktopEmbedBatchSize
		batchTokens = desktopEmbedBatchTokens
	}
	for start := 0; start < len(texts); {
		end := start
		budget := 0
		for end < len(texts) && end-start < maxBatch && budget < batchTokens {
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
	vs := &VectorStore{Model: model, Dims: model.Dims, Count: len(texts), Hash: hash, ChunkerVersion: ChunkerVersion, vecs: vecs}
	if err := writeVectorFile(VectorFilePath(root, model.Slug), vs); err != nil {
		return fmt.Errorf("vecteurs: ecriture: %w", err)
	}
	return nil
}
