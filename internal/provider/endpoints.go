package provider

import (
	"net/http"
	"sort"
	"sync"
)

type endpointDef struct {
	base string
	path string
}

var cloudEndpoints = map[string]endpointDef{
	"deepseek":    {"https://api.deepseek.com", "/chat/completions"},
	"openrouter":  {"https://openrouter.ai/api/v1", "/chat/completions"},
	"opencode":    {"https://opencode.ai/zen/v1", "/chat/completions"},
	"opencode-go": {"https://opencode.ai/zen/go/v1", "/chat/completions"},
}

var LocalEngines = []string{"llamacpp", "ollama", "lmstudio"}

// CloudProviderIDs retourne les IDs des providers cloud connus, tries.
func CloudProviderIDs() []string {
	out := make([]string, 0, len(cloudEndpoints))
	for id := range cloudEndpoints {
		out = append(out, id)
	}
	sort.Strings(out)
	return out
}

// IsCloudProvider indique si l'ID correspond a un provider cloud connu.
func IsCloudProvider(id string) bool {
	_, ok := cloudEndpoints[id]
	return ok
}

func CloudEndpoint(providerID, apiKey string) (Endpoint, bool) {
	def, ok := cloudEndpoints[providerID]
	if !ok {
		return Endpoint{}, false
	}
	ep := Endpoint{Provider: providerID, BaseURL: def.base, Path: def.path, APIKey: apiKey}
	if providerID == "openrouter" {
		ep.Headers = map[string]string{
			"HTTP-Referer": "https://cetas.local",
			"X-Title":      "cetas-lite",
		}
	}
	if providerID == "opencode" || providerID == "opencode-go" {
		// Depuis 2026-09-05, la gateway OpenCode exige x-opencode-session
		// (ID stable par conversation : routage + réutilisation du cache
		// de prompt) ; sans lui : HTTP 400. Ici : UUID stable par instance
		// (les appelants n'ont pas d'ID de conversation sous la main) —
		// le per-conversation restera une optimisation ultérieure.
		ep.Headers = map[string]string{
			"x-opencode-session": "6ba7b811-9dad-11d1-80b4-00c04fd430c8",
		}
	}
	return ep, true
}

func LocalEndpoint(engine, baseURL string) Endpoint {
	return Endpoint{Provider: engine, BaseURL: baseURL, Path: "/v1/chat/completions"}
}

type Registry struct {
	mu        sync.RWMutex
	providers map[string]Provider
}

func NewRegistry() *Registry {
	return &Registry{providers: map[string]Provider{}}
}

func (r *Registry) Set(p Provider) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.providers[p.ID()] = p
}

func (r *Registry) Get(id string) (Provider, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.providers[id]
	return p, ok
}

func (r *Registry) IDs() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]string, 0, len(r.providers))
	for id := range r.providers {
		out = append(out, id)
	}
	return out
}

// Delete retire un provider du registre. Sans effet s'il est absent.
func (r *Registry) Delete(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.providers, id)
}

func Build(keys map[string]string, localURLs map[string]string, client *http.Client) *Registry {
	r := NewRegistry()
	for id := range cloudEndpoints {
		if key := keys[id]; key != "" {
			if ep, ok := CloudEndpoint(id, key); ok {
				r.Set(NewOpenAICompat(ep, client))
			}
		}
	}
	for engine, url := range localURLs {
		if url != "" {
			r.Set(NewOpenAICompat(LocalEndpoint(engine, url), client))
		}
	}
	return r
}
