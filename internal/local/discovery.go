package local

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Engine struct {
	ID  string
	URL string
}

type Model struct {
	Engine string
	ID     string
}

var EngineIDs = []string{"llamacpp", "ollama", "lmstudio"}

func DefaultEngines(urls map[string]string) []Engine {
	var out []Engine
	for _, id := range EngineIDs {
		if u := strings.TrimSpace(urls[id]); u != "" {
			out = append(out, Engine{ID: id, URL: strings.TrimRight(u, "/")})
		}
	}
	return out
}

type Discoverer struct {
	engines  []Engine
	client   *http.Client
	ttl      time.Duration
	mu       sync.Mutex
	cache    []Model
	cachedAt time.Time
}

func New(engines []Engine, client *http.Client) *Discoverer {
	if client == nil {
		client = &http.Client{Timeout: 2500 * time.Millisecond}
	}
	return &Discoverer{engines: engines, client: client, ttl: 30 * time.Second}
}

func (d *Discoverer) Models(ctx context.Context, force bool) []Model {
	d.mu.Lock()
	defer d.mu.Unlock()
	if !force && !d.cachedAt.IsZero() && time.Since(d.cachedAt) < d.ttl {
		return append([]Model(nil), d.cache...)
	}
	var out []Model
	for _, e := range d.engines {
		for _, id := range d.probe(ctx, e) {
			out = append(out, Model{Engine: e.ID, ID: id})
		}
	}
	d.cache = out
	d.cachedAt = time.Now()
	return append([]Model(nil), out...)
}

func (d *Discoverer) ModelsForEngine(ctx context.Context, engine string) []Model {
	all := d.Models(ctx, false)
	var out []Model
	for _, m := range all {
		if m.Engine == engine {
			out = append(out, m)
		}
	}
	return out
}

func (d *Discoverer) probe(ctx context.Context, e Engine) []string {
	probeCtx, cancel := context.WithTimeout(ctx, 2500*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(probeCtx, http.MethodGet, e.URL+"/v1/models", nil)
	if err != nil {
		return nil
	}
	resp, err := d.client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil
	}
	var payload struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
		Models []struct {
			Name string `json:"name"`
		} `json:"models"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil
	}
	var ids []string
	for _, m := range payload.Data {
		if m.ID != "" {
			ids = append(ids, m.ID)
		}
	}
	for _, m := range payload.Models {
		if m.Name != "" {
			ids = append(ids, m.Name)
		}
	}
	return ids
}
