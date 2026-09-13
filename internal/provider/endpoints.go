package provider

import "net/http"

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
	return ep, true
}

func LocalEndpoint(engine, baseURL string) Endpoint {
	return Endpoint{Provider: engine, BaseURL: baseURL, Path: "/v1/chat/completions"}
}

type Registry struct {
	providers map[string]Provider
}

func NewRegistry() *Registry {
	return &Registry{providers: map[string]Provider{}}
}

func (r *Registry) Set(p Provider) {
	r.providers[p.ID()] = p
}

func (r *Registry) Get(id string) (Provider, bool) {
	p, ok := r.providers[id]
	return p, ok
}

func (r *Registry) IDs() []string {
	out := make([]string, 0, len(r.providers))
	for id := range r.providers {
		out = append(out, id)
	}
	return out
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
