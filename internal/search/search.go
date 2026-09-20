package search

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	searchTimeout = 8 * time.Second
	cacheTTL      = 300 * time.Second
	defaultMax    = 10
	// searchRetryDelay : delai avant l'unique retry sur echec total
	// (tous les backends en erreur). Pas de retry sur "aucun resultat".
	searchRetryDelay = time.Second
)

type Hit struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
	Source      string `json:"source,omitempty"`
}

type Result struct {
	Hits     []Hit  `json:"hits"`
	Provider string `json:"provider"`
	Error    string `json:"error,omitempty"`
}

type Searcher struct {
	client      *http.Client
	fetchClient *http.Client
	endpoints   map[string]string

	cfgMu   sync.RWMutex
	keys    map[string]string // id -> cle API
	enabled map[string]bool   // id -> interrupteur (defaut : utilisable)
	order   []string
	mode    string // "race" (le premier qui repond gagne) ou "priority"

	mu    sync.Mutex
	cache map[string]cacheEntry
}

// Config decrit la configuration du Searcher.
type Config struct {
	Keys    map[string]string
	Enabled map[string]bool
	Order   []string
	Mode    string // "race" (rapidite) ou "priority" (qualite)
}

// DefaultOrder : Brave -> Tavily -> Jina -> Exa -> DuckDuckGo.
var DefaultOrder = []string{"brave", "tavily", "jina", "exa", "duckduckgo"}

// keylessID indique si le provider fonctionne sans cle.
func keylessID(id string) bool { return id == "duckduckgo" }

type cacheEntry struct {
	at     time.Time
	result Result
}

func New(keys map[string]string, client *http.Client) *Searcher {
	return NewWithConfig(Config{Keys: keys, Mode: "priority"}, client)
}

func NewWithConfig(cfg Config, client *http.Client) *Searcher {
	if client == nil {
		client = &http.Client{Timeout: searchTimeout}
	}
	order := cfg.Order
	if len(order) == 0 {
		order = append([]string(nil), DefaultOrder...)
	}
	mode := cfg.Mode
	if mode != "race" && mode != "priority" {
		mode = "priority"
	}
	return &Searcher{
		keys:        cfg.Keys,
		enabled:     cfg.Enabled,
		order:       order,
		mode:        mode,
		client:      client,
		fetchClient: newFetchClient(),
		endpoints: map[string]string{
			"tavily": "https://api.tavily.com/search",
			"exa":    "https://api.exa.ai/search",
			"brave":  "https://api.search.brave.com/res/v1/web/search",
			"jina":   "https://s.jina.ai/",
			// jina_reader : lecteur Jina, repli de web_fetch quand le fetch
			// direct echoue (https://r.jina.ai/ + URL cible).
			"jina_reader": "https://r.jina.ai/",
		},
		cache: map[string]cacheEntry{},
	}
}

// SetConfig remplace la configuration a chaud (sans redemarrage).
func (s *Searcher) SetConfig(cfg Config) {
	s.cfgMu.Lock()
	defer s.cfgMu.Unlock()
	s.keys = cfg.Keys
	s.enabled = cfg.Enabled
	if len(cfg.Order) > 0 {
		s.order = append([]string(nil), cfg.Order...)
	}
	if cfg.Mode == "race" || cfg.Mode == "priority" {
		s.mode = cfg.Mode
	}
}

func (s *Searcher) key(id string) string {
	s.cfgMu.RLock()
	defer s.cfgMu.RUnlock()
	return s.keys[id]
}

func (s *Searcher) getMode() string {
	s.cfgMu.RLock()
	defer s.cfgMu.RUnlock()
	return s.mode
}

// usableLocked : le provider est active et utilisable (cle presente ou sans cle).
func (s *Searcher) usableLocked(id string) bool {
	if s.enabled != nil {
		if en, ok := s.enabled[id]; ok && !en {
			return false
		}
	}
	if keylessID(id) {
		return true
	}
	return strings.TrimSpace(s.keys[id]) != ""
}

func (s *Searcher) activeProviders() []string {
	s.cfgMu.RLock()
	defer s.cfgMu.RUnlock()
	var out []string
	for _, id := range s.order {
		if s.usableLocked(id) {
			out = append(out, id)
		}
	}
	return out
}

type providerFn func(ctx context.Context, query string, maxResults int) ([]Hit, error)

func (s *Searcher) provider(id string) providerFn {
	switch id {
	case "tavily":
		return s.searchTavily
	case "exa":
		return s.searchExa
	case "brave":
		return s.searchBrave
	case "jina":
		return s.searchJina
	case "duckduckgo":
		return s.searchDuckDuckGo
	}
	return nil
}

func (s *Searcher) Search(ctx context.Context, query string, maxResults int) Result {
	query = strings.TrimSpace(query)
	if query == "" {
		return Result{Error: "requete vide"}
	}
	if maxResults <= 0 || maxResults > defaultMax {
		maxResults = defaultMax
	}
	active := s.activeProviders()
	if len(active) == 0 {
		return Result{Provider: "none", Error: "recherche web non configuree (Configuration -> Recherche Web)"}
	}
	key := cacheKey(query, maxResults, active)
	if r, ok := s.cached(key); ok {
		return r
	}

	hits, provider, hadErr, errs := s.searchOnce(ctx, query, maxResults, active)
	// Retry unique sur echec total avec erreurs (micro-coupure reseau,
	// backend passagerement indisponible). Pas de retry quand les backends
	// ont repondu mais sans resultats : une seconde tentative ne servirait
	// a rien.
	if provider == "none" && hadErr {
		select {
		case <-ctx.Done():
		case <-time.After(searchRetryDelay):
			hits, provider, _, errs = s.searchOnce(ctx, query, maxResults, active)
		}
	}

	if provider != "none" {
		res := Result{Hits: hits, Provider: provider}
		s.store(key, res)
		return res
	}
	if len(errs) == 0 {
		errs = append(errs, "aucun resultat")
	}
	return Result{Hits: []Hit{}, Provider: "none", Error: strings.Join(errs, "; ")}
}

// searchOnce execute un tour de fan-out sur les providers actifs : tous sont
// interroges en parallele. En mode "race", le premier qui repond avec des
// resultats gagne (latence minimale). En mode "priority", on retourne le
// provider le plus prioritaire des qu'il a repondu, les autres sont annules.
// hadErr vaut true si au moins un provider a renvoye une erreur, par
// opposition a une reponse vide ("aucun resultat").
func (s *Searcher) searchOnce(ctx context.Context, query string, maxResults int, active []string) (hits []Hit, provider string, hadErr bool, errs []string) {
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()
	type outcome struct {
		idx  int
		hits []Hit
		err  error
	}
	ch := make(chan outcome, len(active))
	var wg sync.WaitGroup
	var launched []int
	for i, id := range active {
		fn := s.provider(id)
		if fn == nil {
			continue
		}
		launched = append(launched, i)
		wg.Add(1)
		go func(i int, fn providerFn) {
			defer wg.Done()
			pctx, pcancel := context.WithTimeout(cctx, searchTimeout)
			defer pcancel()
			hits, err := fn(pctx, query, maxResults)
			select {
			case ch <- outcome{idx: i, hits: hits, err: err}:
			case <-cctx.Done():
			}
		}(i, fn)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	pending := make(map[int]bool, len(launched))
	for _, i := range launched {
		pending[i] = true
	}
	best := -1
	var bestHits []Hit
	race := s.getMode() == "race"
	for o := range ch {
		delete(pending, o.idx)
		if len(o.hits) > 0 {
			// Mode "race" (rapidite) : le premier provider qui repond avec
			// des resultats gagne, les autres sont annules. La latence est
			// celle du provider le plus rapide, pas celle du plus prioritaire.
			// Mode "priority" (qualite) : on prefere le provider le plus
			// prioritaire parmi ceux qui ont repondu.
			if race || best == -1 || o.idx < best {
				best = o.idx
				bestHits = o.hits
			}
			if race {
				break
			}
		} else if o.err != nil {
			errs = append(errs, active[o.idx]+": "+o.err.Error())
		}
		if !race {
			// Inutile d'attendre les providers moins prioritaires que le
			// meilleur resultat deja obtenu.
			minPending := -1
			for idx := range pending {
				if minPending == -1 || idx < minPending {
					minPending = idx
				}
			}
			if best != -1 && (minPending == -1 || minPending > best) {
				break
			}
		}
	}
	cancel() // libere les providers encore en vol

	if best != -1 {
		return bestHits, active[best], false, nil
	}
	return nil, "none", len(errs) > 0, errs
}

func (s *Searcher) cached(key string) (Result, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.cache[key]
	if !ok || time.Since(e.at) > cacheTTL {
		return Result{}, false
	}
	return e.result, true
}

func (s *Searcher) store(key string, r Result) {
	s.mu.Lock()
	s.cache[key] = cacheEntry{at: time.Now(), result: r}
	s.mu.Unlock()
}

func cacheKey(query string, maxResults int, providers []string) string {
	raw := fmt.Sprintf("%s|%d|%s", query, maxResults, strings.Join(providers, ","))
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}

func truncateRunes(s string, max int) string {
	s = strings.TrimSpace(s)
	if max <= 0 {
		return s
	}
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	return string(r[:max]) + "\n[...]"
}

var (
	titleKeys  = []string{"title", "headline", "name", "heading"}
	urlKeys    = []string{"url", "link", "href", "uri", "permalink"}
	descKeys   = []string{"description", "snippet", "content", "preview", "summary", "text", "body"}
	sourceKeys = []string{"source", "domain", "displayLink", "displayed_link", "engine"}
)

func normalizeHit(raw map[string]any) (Hit, bool) {
	title := firstString(raw, titleKeys)
	link := firstString(raw, urlKeys)
	if title == "" && link == "" {
		return Hit{}, false
	}
	if title == "" {
		title = link
	}
	if link == "" {
		link = title
	}
	return Hit{
		Title:       sanitize(title),
		URL:         sanitize(link),
		Description: sanitize(firstString(raw, descKeys)),
		Source:      sanitize(firstString(raw, sourceKeys)),
	}, true
}

func hitsFrom(v any) []Hit {
	arr, ok := v.([]any)
	if !ok {
		return nil
	}
	out := make([]Hit, 0, len(arr))
	for _, item := range arr {
		m, ok := item.(map[string]any)
		if !ok {
			continue
		}
		if h, ok := normalizeHit(m); ok {
			out = append(out, h)
		}
	}
	return out
}

func firstString(raw map[string]any, keys []string) string {
	for _, k := range keys {
		if v, ok := raw[k].(string); ok && strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

func sanitize(s string) string {
	return strings.TrimSpace(strings.Map(func(r rune) rune {
		if r < 0x20 && r != '\t' && r != '\n' && r != '\r' {
			return -1
		}
		return r
	}, s))
}

func (s *Searcher) searchTavily(ctx context.Context, query string, maxResults int) ([]Hit, error) {
	req, err := s.jsonRequest(ctx, http.MethodPost, s.endpoints["tavily"], map[string]any{
		"query": query, "max_results": maxResults, "include_answer": false,
	})
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+s.key("tavily"))
	data, err := s.do(req)
	if err != nil {
		return nil, err
	}
	return hitsFrom(data["results"]), nil
}

func (s *Searcher) searchExa(ctx context.Context, query string, maxResults int) ([]Hit, error) {
	req, err := s.jsonRequest(ctx, http.MethodPost, s.endpoints["exa"], map[string]any{
		"query": query, "numResults": maxResults, "type": "auto",
	})
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", s.key("exa"))
	data, err := s.do(req)
	if err != nil {
		return nil, err
	}
	return hitsFrom(data["results"]), nil
}

func (s *Searcher) searchBrave(ctx context.Context, query string, maxResults int) ([]Hit, error) {
	endpoint := s.endpoints["brave"] + "?" + url.Values{
		"q":     {query},
		"count": {fmt.Sprint(maxResults)},
	}.Encode()
	req, err := s.jsonRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Subscription-Token", s.key("brave"))
	data, err := s.do(req)
	if err != nil {
		return nil, err
	}
	web, _ := data["web"].(map[string]any)
	return hitsFrom(web["results"]), nil
}

func (s *Searcher) searchJina(ctx context.Context, query string, maxResults int) ([]Hit, error) {
	endpoint := s.endpoints["jina"] + "?" + url.Values{"q": {query}}.Encode()
	req, err := s.jsonRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.key("jina"))
	data, err := s.do(req)
	if err != nil {
		return nil, err
	}
	hits := hitsFrom(data["data"])
	if len(hits) > maxResults {
		hits = hits[:maxResults]
	}
	return hits, nil
}

func (s *Searcher) jsonRequest(ctx context.Context, method, endpoint string, body any) (*http.Request, error) {
	var rdr io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		rdr = bytes.NewReader(b)
	}
	req, err := http.NewRequestWithContext(ctx, method, endpoint, rdr)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

func (s *Searcher) do(req *http.Request) (map[string]any, error) {
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 4096))
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	var out map[string]any
	if err := json.NewDecoder(io.LimitReader(resp.Body, 4<<20)).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}
