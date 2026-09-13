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
	keys        map[string]string
	client      *http.Client
	fetchClient *http.Client
	order       []string
	endpoints   map[string]string

	mu    sync.Mutex
	cache map[string]cacheEntry
}

type cacheEntry struct {
	at     time.Time
	result Result
}

func New(keys map[string]string, client *http.Client) *Searcher {
	if client == nil {
		client = &http.Client{Timeout: searchTimeout}
	}
	return &Searcher{
		keys:        keys,
		client:      client,
		fetchClient: newFetchClient(),
		order:       []string{"tavily", "exa", "brave", "jina"},
		endpoints: map[string]string{
			"tavily": "https://api.tavily.com/search",
			"exa":    "https://api.exa.ai/search",
			"brave":  "https://api.search.brave.com/res/v1/web/search",
			"jina":   "https://s.jina.ai/",
		},
		cache: map[string]cacheEntry{},
	}
}

func (s *Searcher) activeProviders() []string {
	var out []string
	for _, id := range s.order {
		if strings.TrimSpace(s.keys[id]) != "" {
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
		return Result{Error: "recherche web non configuree (keys set tavily|exa|brave|jina)"}
	}
	key := cacheKey(query, maxResults, active)
	if r, ok := s.cached(key); ok {
		return r
	}

	type outcome struct {
		idx  int
		hits []Hit
		err  error
	}
	ch := make(chan outcome, len(active))
	var wg sync.WaitGroup
	for i, id := range active {
		fn := s.provider(id)
		if fn == nil {
			continue
		}
		wg.Add(1)
		go func(i int, fn providerFn) {
			defer wg.Done()
			cctx, cancel := context.WithTimeout(ctx, searchTimeout)
			defer cancel()
			hits, err := fn(cctx, query, maxResults)
			ch <- outcome{idx: i, hits: hits, err: err}
		}(i, fn)
	}
	wg.Wait()
	close(ch)

	var best *outcome
	var errs []string
	for o := range ch {
		if len(o.hits) > 0 {
			if best == nil || o.idx < best.idx {
				cp := o
				best = &cp
			}
			continue
		}
		if o.err != nil {
			errs = append(errs, active[o.idx]+": "+o.err.Error())
		}
	}
	if best != nil {
		res := Result{Hits: best.hits, Provider: active[best.idx]}
		s.store(key, res)
		return res
	}
	if len(errs) == 0 {
		errs = append(errs, "aucun resultat")
	}
	return Result{Hits: []Hit{}, Provider: "none", Error: strings.Join(errs, "; ")}
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
	req.Header.Set("Authorization", "Bearer "+s.keys["tavily"])
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
	req.Header.Set("x-api-key", s.keys["exa"])
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
	req.Header.Set("X-Subscription-Token", s.keys["brave"])
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
	req.Header.Set("Authorization", "Bearer "+s.keys["jina"])
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
