// Package rag : index de recherche local (mots-clés + facettes) sur un
// corpus decoupe pour la recuperation. 100% in-process : aucun appel reseau,
// aucune dependance externe, aucune ecriture. L'index est immuable une fois
// construit (lectures concurrentes sans verrou), et toute recherche est
// bornee dans le temps (fail-open : jamais d'erreur remontee).
package rag

import (
	"context"
	"crypto/sha256"
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
	// Lot 2 (hybride) : signaux de couverture, independants du classement.
	// BM25Top : meilleur score BM25 brut de la requete (0 si aucun hit).
	// SemTop : meilleure similarite cosinus semantique (0 si la jambe
	// semantique n'a pas contribue).
	// Hybrid : la jambe semantique a contribue au resultat.
	BM25Top float64
	SemTop  float64
	Hybrid  bool
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
	return ix.search(ctx, query, limit, "", nil)
}

// SearchCorpus : comme Search mais restreint aux chunks d'un corpus
// (focus au clic sur une question suggeree). Corpus vide ou inconnu :
// aucun chunk ne matche, resultat vide (fail-open).
func (ix *Index) SearchCorpus(ctx context.Context, query, corpus string, limit int) Result {
	return ix.search(ctx, query, limit, corpus, nil)
}

// SearchBoosted : comme Search mais avec un facteur de boost par terme de
// requete (iteration 6b : boost des entites reprises de l'historique,
// dont l'idf faible — df eleve — les ferait passer apres des termes
// conversationnels rares). boost nil ou terme absent => facteur 1.0
// (strictement identique a Search).
func (ix *Index) SearchBoosted(ctx context.Context, query string, boost map[string]float64, limit int) Result {
	return ix.search(ctx, query, limit, "", boost)
}

func (ix *Index) search(ctx context.Context, query string, limit int, corpus string, boost map[string]float64) Result {
	limit = clampLimit(limit)
	if ix == nil {
		return Result{}
	}
	res := Result{Total: len(ix.chunks), Corpora: ix.stats.Corpora, Ready: ix.stats.Ready}
	terms := uniqueTerms(tokenize(cleanQueryForSearch(query)))
	if len(terms) == 0 || len(ix.chunks) == 0 {
		return res
	}
	ranked := ix.rankChunks(ctx, terms, corpus, boost)
	if len(ranked) > limit {
		ranked = ranked[:limit]
	}
	if len(ranked) > 0 {
		res.BM25Top = ranked[0].score
	}
	res.Hits = make([]Hit, 0, len(ranked))
	for rank, rc := range ranked {
		res.Hits = append(res.Hits, ix.hitFor(rc.idx, terms, rc.score, rank))
	}
	return res
}

// rankedChunk : chunk candidat avec son score BM25 (fusion hybride, lot 2).
type rankedChunk struct {
	idx   int
	score float64
}

// rankChunks : scoring BM25 et tri decroissant (egalite => chemin croissant,
// deterministe). Retourne les candidats avant troncature a limit.
func (ix *Index) rankChunks(ctx context.Context, terms []string, corpus string, boost map[string]float64) []rankedChunk {
	inScope := func(i int) bool {
		return corpus == "" || ix.chunks[i].corpus == corpus
	}

	n := 0
	for i := range ix.chunks {
		if inScope(i) {
			n++
		}
	}
	if n == 0 {
		return nil
	}

	nf := float64(n)
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
		df := 0
		for _, id := range ids {
			if inScope(int(id)) {
				df++
			}
		}
		if df == 0 {
			continue
		}
		idf := math.Log(1 + (nf-float64(df)+0.5)/(float64(df)+0.5))
		bt := 1.0
		if boost != nil {
			if v, ok := boost[t]; ok {
				bt = v
			}
		}
		for _, id := range ids {
			if !inScope(int(id)) {
				continue
			}
			c := &ix.chunks[int(id)]
			tf := float64(c.tf[t])
			dl := float64(c.length)
			denom := tf + k1*(1-b+b*dl/ix.avgLen)
			if denom <= 0 {
				continue
			}
			scores[id] += bt * idf * (tf * (k1 + 1)) / denom
			matched[id]++
		}
	}

	order := make([]int, 0, len(ix.chunks))
	for i := range ix.chunks {
		if matched[i] > 0 && inScope(i) {
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
	ranked := make([]rankedChunk, 0, len(order))
	for _, i := range order {
		ranked = append(ranked, rankedChunk{idx: i, score: scores[i]})
	}
	return ranked
}

// hitFor : construit le Hit d'un chunk (extrait ancre sur terms).
// rerankText : texte d'un chunk envoye au reranker. Le reranker juge la
// pertinence sur le fond du document, pas sur l'extrait requete : on
// envoie le debut du corps (borne, le serveur a ses propres limites).
func (ix *Index) rerankText(i int) string {
	if ix == nil || i < 0 || i >= len(ix.chunks) {
		return ""
	}
	b := ix.chunks[i].body
	if len(b) > rerankTextMaxChars {
		b = b[:rerankTextMaxChars]
	}
	return b
}

// rerankTextMaxChars : borne du texte transmis au reranker (par candidat).
const rerankTextMaxChars = 4000

func (ix *Index) hitFor(i int, terms []string, score float64, rank int) Hit {
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
	return Hit{
		Path: c.path, Corpus: c.corpus, Title: c.title, Domain: c.domain,
		Task: c.task, Actors: c.actors, Dates: c.dates, Keywords: c.keywords,
		Section: c.section, DeltaOf: c.deltaOf,
		Score: score, Snippet: snip, Excerpt: ex,
	}
}

// ChunkRef : reference lisible d'un chunk (diagnostic : `rag-vectors probe`).
func (ix *Index) ChunkRef(i int) string {
	if ix == nil || i < 0 || i >= len(ix.chunks) {
		return ""
	}
	c := &ix.chunks[i]
	if c.section != "" {
		return c.path + " :: " + c.section
	}
	return c.path
}

// CorpusHash : empreinte du corpus (path + "\x00" + body par chunk, en
// ordre). Les vecteurs d'embedding sont alignes 1:1 avec l'index via ce
// hash : corpus modifie => vecteurs obsoletes (fail-open).
func (ix *Index) CorpusHash() [32]byte {
	h := sha256.New()
	if ix != nil {
		for _, c := range ix.chunks {
			h.Write([]byte(c.path))
			h.Write([]byte{0})
			h.Write([]byte(c.body))
			h.Write([]byte{0})
		}
	}
	var out [32]byte
	copy(out[:], h.Sum(nil))
	return out
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

// ragQueryFillerPhrases : formulations conversationnelles retirees de la
// requete avant BM25 (iteration 3). Un mot-outil rare dans le corpus
// (ex. « nouveau ») recevait un IDF eleve + boost titre et ejectait les
// vrais sujets : « De nouveau sur kimi k3? » classait un article contenant
// « nouveau » dans son titre devant la section Kimi dediee. Formes deja
// normalisees (minuscules, sans accents, separateurs -> espaces) ;
// comparaison avec frontieres de mots. Les identifiants significatifs
// (kimi, k3, glm-5.3…) ne sont jamais touches : seules les formulations
// vides de sens documentaire sont retirees, jamais un mot isole qui
// pourrait etre porteur.
//
// Extension (correctif 2026-09-24, incohérence inter-sessions GLM) :
// « Il y a du nouveau chez glm » (sens « des nouveautés », pas la
// reformulation « de nouveau ») laissait passer « nouveau » — terme rare
// donc IDF élevé — qui éjectait les vrais sujets (« glm », df élevé donc
// IDF faible). Mesuré sur le corpus réel (2487 chunks) : sans « du
// nouveau », un article Kimi contenant « nouveau » dans son titre
// arrivait #1 (16.32) devant tout extrait GLM.
var ragQueryFillerPhrases = []string{
	// Francais.
	"de nouveau", "du nouveau", "parle moi de",
	"qu est ce que", "qu est ce qu",
	"s il te plait", "s il vous plait",
	"dis moi", "explique moi",
	"c est quoi", "stp", "svp",
	// Anglais.
	"tell me about", "what is", "what s",
	"can you", "could you", "please",
}

// ragQueryStopwords : mots-outils FR/EN retires de la requete (niveau
// token). Un mot-outil rare dans le corpus (ex. « sur », df=13) recevait
// un IDF eleve et faussait le classement meme apres retrait des
// formulations (« De nouveau sur kimi k3? » -> « sur kimi k3 »). Formes
// normalisees. Jamais de mot pouvant etre un identifiant.
//
// Extension (correctif 2026-09-24) : résidu conversationnel mesuré sur
// les requêtes GLM fautives — « tu », « je », « peut », « dire »,
// « sorti », « dernier », « pense », « men » (sms), « na » (sms),
// « aucun », « info », « chez », « ya » (sms), « nouveau »… survivaient
// au nettoyage ; rares dans le corpus (IDF élevé), ils écrasaient
// l'entité « glm » (fréquente, IDF faible). Ex. « GLM-5.3 est sorti tu
// peut men dire plus? » -> « glm sorti tu peut men dire » classait un
// article « PC gaming titane » #1 (16.00) sans aucun extrait GLM dans
// le top-3. Les pronoms « je »/« tu » manquaient à la liste initiale.
var ragQueryStopwords = map[string]bool{
	// Francais.
	"le": true, "la": true, "les": true, "de": true, "des": true, "du": true,
	"un": true, "une": true, "et": true, "est": true, "sont": true,
	"dans": true, "pour": true, "avec": true, "sur": true, "par": true,
	"au": true, "aux": true, "ce": true, "cet": true, "cette": true,
	"ces": true, "il": true, "elle": true, "ils": true, "elles": true,
	"qui": true, "que": true, "quoi": true, "quand": true, "ou": true,
	"comment": true, "ne": true, "pas": true, "plus": true, "moins": true,
	"se": true, "son": true, "sa": true, "ses": true,
	"notre": true, "nos": true, "votre": true, "vos": true,
	"leur": true, "leurs": true, "mais": true, "donc": true, "or": true,
	"ni": true, "car": true, "comme": true, "tout": true, "tous": true,
	"toute": true, "toutes": true, "aussi": true, "tres": true,
	"bien": true, "encore": true, "deja": true, "alors": true, "si": true,
	"etre": true, "avoir": true, "faire": true,
	// Correctif 2026-09-24 : residu conversationnel (mesure ci-dessus).
	"je": true, "tu": true, "nous": true, "vous": true,
	"me": true, "te": true, "moi": true, "toi": true, "en": true,
	"y": true, "cela": true, "ceci": true, "ca": true,
	"na": true, "men": true, "ya": true,
	"peut": true, "peux": true, "peuvent": true, "pouvoir": true,
	"dire": true, "dis": true, "dit": true, "dites": true,
	"sais": true, "sait": true, "savoir": true,
	"connais": true, "connait": true,
	"sorti": true, "sortie": true, "sortis": true, "sorties": true,
	"nouveau": true, "nouvelle": true, "nouveaux": true, "nouvelles": true,
	"dernier": true, "derniere": true, "derniers": true, "dernieres": true,
	"chez": true, "aucun": true, "aucune": true,
	"info": true, "infos": true,
	"modele": true, "modeles": true,
	"pense": true, "penses": true, "penser": true,
	// Anglais.
	"the": true, "an": true, "of": true, "in": true, "on": true,
	"and": true, "is": true, "are": true, "was": true, "were": true,
	"be": true, "been": true, "to": true, "for": true, "with": true,
	"as": true, "at": true, "by": true, "from": true, "that": true,
	"this": true, "it": true, "its": true, "not": true, "you": true,
	"your": true, "we": true, "they": true, "their": true,
	"has": true, "have": true, "had": true, "will": true, "would": true,
	"do": true, "does": true, "did": true, "if": true, "then": true,
	"than": true, "so": true, "no": true,
	// Correctif 2026-09-24 : equivalents EN du residu ci-dessus.
	"know": true, "mean": true, "means": true,
	"tell": true, "say": true, "says": true,
	"latest": true, "new": true,
}

// cleanQueryForSearch : nettoie la requete des formulations
// conversationnelles et des mots-outils avant tokenisation BM25.
// Les separateurs deviennent des espaces pour que « plaît, » matche
// comme « plait ». Si le nettoyage vide la requete, les termes d'origine
// sont conserves (fail-open).
func cleanQueryForSearch(q string) string {
	var b strings.Builder
	b.Grow(len(q))
	for _, r := range normalize(q) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' {
			b.WriteRune(r)
		} else {
			b.WriteRune(' ')
		}
	}
	n := " " + strings.Join(strings.Fields(b.String()), " ") + " "
	for _, p := range ragQueryFillerPhrases {
		n = strings.ReplaceAll(n, " "+p+" ", " ")
	}
	var kept []string
	for _, t := range tokenize(n) {
		if !ragQueryStopwords[t] {
			kept = append(kept, t)
		}
	}
	if len(kept) == 0 {
		return q
	}
	return strings.Join(kept, " ")
}

func snippetFrom(ex string) string {
	if len([]rune(ex)) <= maxSnippet {
		return ex
	}
	return truncateRunes(ex, maxSnippet) + "…"
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
