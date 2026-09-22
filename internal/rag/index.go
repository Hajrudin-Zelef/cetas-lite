// Package rag : index de recherche local (mots-clés + facettes) sur un
// corpus decoupe pour la recuperation. 100% in-process : aucun appel reseau,
// aucune dependance externe, aucune ecriture. L'index est immuable une fois
// construit (lectures concurrentes sans verrou), et toute recherche est
// bornee dans le temps (fail-open : jamais d'erreur remontee).
package rag

import (
	"context"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

const (
	// DefaultLimit : nombre de resultats par defaut.
	DefaultLimit = 8
	// MaxLimit : nombre de resultats maximal.
	MaxLimit = 30
	// DefaultTimeout : budget de latence d'une recherche (fail-open au-dela).
	DefaultTimeout = 120 * time.Millisecond

	maxSnippet   = 240
	excerptChars = 900
	excerptTop   = 4
	maxReadLines = 400

	k1 = 1.2
	b  = 0.75

	minTermLen = 2
	maxTermLen = 40

	boostTitle   = 3
	boostKeyword = 2
	boostBody    = 1
)

// Hit : un chunk pertinent renvoye par Search.
type Hit struct {
	Path     string
	Corpus   string
	Title    string
	Domain   string
	Task     string
	Actors   []string
	Dates    []string
	Keywords []string
	Section  string
	DeltaOf  string
	Score    float64
	Snippet  string
	Excerpt  string
}

// Result : reponse d'une recherche.
type Result struct {
	Hits    []Hit
	Total   int
	Corpora int
	Ready   bool
}

// Stats : etat de l'index.
type Stats struct {
	Ready    bool
	Corpora  int
	Chunks   int
	Terms    int
	Bytes    int64
	Duration time.Duration
	Built    time.Time
	Errors   []string
}

type chunk struct {
	path     string
	corpus   string
	title    string
	domain   string
	task     string
	actors   []string
	dates    []string
	keywords []string
	section  string
	deltaOf  string
	body     string
	tf       map[string]int
	length   int
}

// Index : index immuable.
type Index struct {
	chunks   []chunk
	postings map[string][]int32
	avgLen   float64
	stats    Stats
}

// Search : classe les chunks par pertinence (BM25-lite, champs ponderes).
// Respecte ctx ; s'arrete proprement si le budget est depasse.
func (ix *Index) Search(ctx context.Context, query string, limit int) Result {
	limit = clampLimit(limit)
	if ix == nil {
		return Result{}
	}
	res := Result{Total: len(ix.chunks), Corpora: ix.stats.Corpora, Ready: ix.stats.Ready}
	terms := uniqueTerms(tokenize(query))
	if len(terms) == 0 || len(ix.chunks) == 0 {
		return res
	}

	n := float64(len(ix.chunks))
	scores := make([]float64, len(ix.chunks))
	matched := make([]int, len(ix.chunks))

	for _, t := range terms {
		if ctx != nil && ctx.Err() != nil {
			break
		}
		ids, ok := ix.postings[t]
		if !ok {
			continue
		}
		idf := math.Log(1 + (n-float64(len(ids))+0.5)/(float64(len(ids))+0.5))
		for _, id := range ids {
			c := &ix.chunks[int(id)]
			tf := float64(c.tf[t])
			dl := float64(c.length)
			denom := tf + k1*(1-b+b*dl/ix.avgLen)
			if denom <= 0 {
				continue
			}
			scores[id] += idf * (tf * (k1 + 1)) / denom
			matched[id]++
		}
	}

	order := make([]int, 0, len(ix.chunks))
	for i := range ix.chunks {
		if matched[i] > 0 {
			order = append(order, i)
		}
	}
	sort.SliceStable(order, func(i, j int) bool {
		a, b := order[i], order[j]
		if scores[a] != scores[b] {
			return scores[a] > scores[b]
		}
		return ix.chunks[a].path < ix.chunks[b].path
	})
	if len(order) > limit {
		order = order[:limit]
	}

	res.Hits = make([]Hit, 0, len(order))
	for rank, i := range order {
		c := &ix.chunks[i]
		// La plage de correspondance est calculee une seule fois (les
		// extraits 240 et 900 car. partagent le meme ancrage).
		mstart, mend, mok := firstMatchRange(c.body, terms)
		ex := excerptAround(c.body, mstart, mend, mok, maxSnippet)
		snip := ex
		if rank < excerptTop {
			ex = excerptAround(c.body, mstart, mend, mok, excerptChars)
			snip = snippetFrom(ex)
		}
		res.Hits = append(res.Hits, Hit{
			Path: c.path, Corpus: c.corpus, Title: c.title, Domain: c.domain,
			Task: c.task, Actors: c.actors, Dates: c.dates, Keywords: c.keywords,
			Section: c.section, DeltaOf: c.deltaOf,
			Score: scores[i], Snippet: snip, Excerpt: ex,
		})
	}
	return res
}

// Read : relit un chunk par son chemin (ou son nom de base), lignes numerotees.
func (ix *Index) Read(rel string, offset, limit int) (string, error) {
	if ix == nil || !ix.stats.Ready {
		return "", fmt.Errorf("index RAG indisponible")
	}
	rel = strings.TrimSpace(rel)
	if rel == "" {
		return "", fmt.Errorf("chemin vide")
	}
	idx := -1
	for i := range ix.chunks {
		if ix.chunks[i].path == rel {
			idx = i
			break
		}
	}
	if idx < 0 {
		base := pathBase(rel)
		for i := range ix.chunks {
			if pathBase(ix.chunks[i].path) == base {
				idx = i
				break
			}
		}
	}
	if idx < 0 {
		return "", fmt.Errorf("chunk '%s' introuvable", rel)
	}
	return numbered(ix.chunks[idx].body, offset, limit), nil
}

// Stats : copie de l'etat de l'index.
func (ix *Index) Stats() Stats {
	if ix == nil {
		return Stats{}
	}
	return ix.stats
}

func clampLimit(limit int) int {
	if limit <= 0 {
		return DefaultLimit
	}
	if limit > MaxLimit {
		return MaxLimit
	}
	return limit
}

func emptyIndex() *Index {
	return &Index{postings: map[string][]int32{}, avgLen: 1}
}

func buildIndex(chunks []chunk, corpora int, errs []string, start time.Time) *Index {
	ix := emptyIndex()
	ix.chunks = chunks
	var total, bytes int64
	for i := range chunks {
		c := &chunks[i]
		c.tf = map[string]int{}
		addTokens(c.tf, tokenize(c.body), boostBody)
		addTokens(c.tf, tokenize(c.title), boostTitle)
		addTokens(c.tf, tokenize(strings.Join(c.keywords, " ")), boostKeyword)
		addTokens(c.tf, tokenize(strings.Join(c.actors, " ")), boostKeyword)
		addTokens(c.tf, tokenize(c.domain), boostKeyword)
		addTokens(c.tf, tokenize(c.task), boostKeyword)
		c.length = 0
		for _, v := range c.tf {
			c.length += v
		}
		total += int64(c.length)
		bytes += int64(len(c.body))
	}
	for i := range chunks {
		for t := range chunks[i].tf {
			ix.postings[t] = append(ix.postings[t], int32(i))
		}
	}
	if len(chunks) > 0 {
		ix.avgLen = float64(total) / float64(len(chunks))
	}
	if ix.avgLen <= 0 {
		ix.avgLen = 1
	}
	ix.stats = Stats{
		Ready: len(chunks) > 0, Corpora: corpora, Chunks: len(chunks),
		Terms: len(ix.postings), Bytes: bytes,
		Duration: time.Since(start), Built: time.Now(), Errors: errs,
	}
	return ix
}

func addTokens(tf map[string]int, toks []string, weight int) {
	for _, t := range toks {
		tf[t] += weight
	}
}

func uniqueTerms(toks []string) []string {
	seen := make(map[string]bool, len(toks))
	out := make([]string, 0, len(toks))
	for _, t := range toks {
		if seen[t] {
			continue
		}
		seen[t] = true
		out = append(out, t)
	}
	return out
}

func normalize(s string) string {
	s = strings.ToLower(s)
	if isASCII(s) {
		return s
	}
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range norm.NFD.String(s) {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			return false
		}
	}
	return true
}

func tokenize(s string) []string {
	s = normalize(s)
	var out []string
	var b strings.Builder
	flush := func() {
		if n := b.Len(); n >= minTermLen && n <= maxTermLen {
			out = append(out, b.String())
		}
		b.Reset()
	}
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
		} else {
			flush()
		}
	}
	flush()
	return out
}

func snippetFrom(ex string) string {
	if len([]rune(ex)) <= maxSnippet {
		return ex
	}
	return truncateRunes(ex, maxSnippet) + "…"
}

// excerptOf : fenetre de texte centree sur le premier terme trouve (jusqu'a
// width runes), espaces normalises. Le texte retourne est toujours extrait
// du corps d'origine (accents et casse preserves). Sans correspondance :
// debut du document.
func excerptOf(body string, terms []string, width int) string {
	start, end, ok := firstMatchRange(body, terms)
	return excerptAround(body, start, end, ok, width)
}

// excerptAround : decoupe la fenetre autour d'une plage deja localisee
// (start/end : offsets octets dans body, sur des frontieres de runes).
func excerptAround(body string, start, end int, ok bool, width int) string {
	if body == "" || width <= 0 {
		return ""
	}
	if !ok {
		head := body
		if len(head) > width*2 {
			head = head[:runeBoundaryBack(body, width*2)]
		}
		return truncateRunes(flatten(head), width)
	}
	ws := start - width/3
	if ws < 0 {
		ws = 0
	}
	we := ws + width
	if we > len(body) {
		we = len(body)
	}
	ws = runeBoundaryFwd(body, ws)
	we = runeBoundaryBack(body, we)
	_ = end // la fin de la correspondance est incluse dans la fenetre ci-dessus
	return ellipsis(ws > 0) + flatten(body[ws:we]) + ellipsis(we < len(body))
}

// firstMatchRange : offsets (octets) dans body de la premiere occurrence
// d'un des termes. D'abord une recherche insensible a la casse (ASCII) sur
// le texte d'origine ; en repli, une recherche sur le texte replie
// (minuscules, sans accents) avec re-projection des offsets sur l'original,
// pour que l'extrait garde accents et casse. ok=false si rien n'est trouve.
func firstMatchRange(body string, terms []string) (start, end int, ok bool) {
	best := -1
	bestLen := 0
	for _, t := range terms {
		if i := indexFold(body, t); i >= 0 && (best < 0 || i < best) {
			best, bestLen = i, len(t)
		}
	}
	if best >= 0 {
		return best, best + bestLen, true
	}
	folded, fmap := foldWithMap(body)
	for _, t := range terms {
		if i := strings.Index(folded, t); i >= 0 && (best < 0 || i < best) {
			best, bestLen = i, len(t)
		}
	}
	if best < 0 {
		return 0, 0, false
	}
	return fmap[best], fmap[best+bestLen], true
}

// foldWithMap : normalisation (minuscules + suppression des accents, meme
// pipeline que normalize) avec table de correspondance : fmap[i] est
// l'offset octet dans s du debut de la rune d'origine dont est issu le
// i-eme octet replie ; fmap[len(folded)] vaut len(s) (sentinelle de fin).
// La normalisation est faite rune par rune (equivalente a normalize pour
// les textes latins ; peut differer sur des cas Unicode exotiques comme le
// sigma final grec, auquel cas la correspondance echoue proprement).
func foldWithMap(s string) (string, []int) {
	var b strings.Builder
	b.Grow(len(s))
	fmap := make([]int, 0, len(s)+1)
	for off := 0; off < len(s); {
		r, w := utf8.DecodeRuneInString(s[off:])
		for _, nr := range norm.NFD.String(string(unicode.ToLower(r))) {
			if unicode.Is(unicode.Mn, nr) {
				continue
			}
			var buf [utf8.UTFMax]byte
			n := utf8.EncodeRune(buf[:], nr)
			for k := 0; k < n; k++ {
				fmap = append(fmap, off)
			}
			b.Write(buf[:n])
		}
		off += w
	}
	fmap = append(fmap, len(s))
	return b.String(), fmap
}

func indexFold(hay, needle string) int {
	if needle == "" || len(needle) > len(hay) {
		return -1
	}
	last := len(hay) - len(needle)
	for i := 0; i <= last; i++ {
		if asciiLower(hay[i]) != needle[0] {
			continue
		}
		ok := true
		for j := 1; j < len(needle); j++ {
			if asciiLower(hay[i+j]) != needle[j] {
				ok = false
				break
			}
		}
		if ok {
			return i
		}
	}
	return -1
}

func asciiLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func flatten(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func ellipsis(v bool) string {
	if v {
		return "…"
	}
	return ""
}

func runeBoundaryBack(s string, i int) int {
	if i > len(s) {
		i = len(s)
	}
	for i > 0 && i < len(s) && !utf8.RuneStart(s[i]) {
		i--
	}
	return i
}

func runeBoundaryFwd(s string, i int) int {
	if i < 0 {
		i = 0
	}
	for i < len(s) && !utf8.RuneStart(s[i]) {
		i++
	}
	return i
}

func pathBase(p string) string {
	p = strings.TrimRight(p, "/")
	if i := strings.LastIndexByte(p, '/'); i >= 0 {
		return p[i+1:]
	}
	return p
}

func numbered(body string, offset, limit int) string {
	lines := strings.Split(body, "\n")
	if offset <= 0 {
		offset = 1
	}
	if offset > len(lines) {
		return ""
	}
	if limit <= 0 || limit > maxReadLines {
		limit = maxReadLines
	}
	var out strings.Builder
	for i := offset - 1; i < len(lines) && i < offset-1+limit; i++ {
		fmt.Fprintf(&out, "%d\t%s\n", i+1, lines[i])
	}
	return strings.TrimRight(out.String(), "\n")
}
