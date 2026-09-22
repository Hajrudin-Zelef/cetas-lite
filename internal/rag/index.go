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
	df       map[string]int
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
		ex := excerptOf(c.body, terms, maxSnippet)
		snip := ex
		if rank < excerptTop {
			ex = excerptOf(c.body, terms, excerptChars)
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
	return &Index{postings: map[string][]int32{}, df: map[string]int{}, avgLen: 1}
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
			ix.df[t]++
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
// width runes), espaces normalises. Sans correspondance : debut du document.
func excerptOf(body string, terms []string, width int) string {
	if body == "" || width <= 0 {
		return ""
	}
	idx, text := firstMatch(body, terms)
	if idx < 0 {
		head := body
		if len(head) > width*2 {
			head = head[:runeBoundaryBack(body, width*2)]
		}
		return truncateRunes(flatten(head), width)
	}
	start := idx - width/3
	if start < 0 {
		start = 0
	}
	end := start + width
	if end > len(text) {
		end = len(text)
	}
	start = runeBoundaryFwd(text, start)
	end = runeBoundaryBack(text, end)
	return ellipsis(start > 0) + flatten(text[start:end]) + ellipsis(end < len(text))
}

// firstMatch : position du premier terme (recherche insensible a la casse,
// sans allocation). Repli sur le texte replie (accents) si introuvable.
func firstMatch(body string, terms []string) (int, string) {
	best := -1
	for _, t := range terms {
		if i := indexFold(body, t); i >= 0 && (best < 0 || i < best) {
			best = i
		}
	}
	if best >= 0 {
		return best, body
	}
	folded := normalize(body)
	for _, t := range terms {
		if i := strings.Index(folded, t); i >= 0 && (best < 0 || i < best) {
			best = i
		}
	}
	return best, folded
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
