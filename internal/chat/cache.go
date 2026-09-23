package chat

import (
	"encoding/json"

	"cetas-lite/internal/cache"
	"cetas-lite/internal/provider"
)

// Branchement du cache exact (itération 5a) sur le tour principal.
//
// La clé couvre l'intégralité de la requête déterministe : provider, modèle,
// messages complets (y compris raisonnement forcé), température, maxTokens,
// raisonnement. La porte d'activation (température 0, sans outils, sans
// paramètres supplémentaires comme la recherche web native) est évaluée dans
// cache.KeyOf : toute requête non éligible contourne le cache (clé vide).

// cacheKeyInput : forme canonique sérialisée pour le SHA256.
type cacheKeyInput struct {
	Provider        string             `json:"provider"`
	Model           string             `json:"model"`
	Messages        []provider.Message `json:"messages"`
	Temperature     float64            `json:"temperature"`
	MaxTokens       int                `json:"max_tokens"`
	EnableReasoning bool               `json:"enable_reasoning"`
	ReasoningEffort string             `json:"reasoning_effort"`
}

// cacheKeyFor rend la clé de cache exacte pour une tentative membre,
// ou "" si la requête n'est pas éligible.
func (e *Engine) cacheKeyFor(providerID string, req provider.Request) string {
	raw, err := json.Marshal(cacheKeyInput{
		Provider:        providerID,
		Model:           req.Model,
		Messages:        req.Messages,
		Temperature:     req.Temperature,
		MaxTokens:       req.MaxTokens,
		EnableReasoning: req.EnableReasoning,
		ReasoningEffort: req.ReasoningEffort,
	})
	if err != nil {
		return ""
	}
	return cache.KeyOf(raw, req.Temperature, len(req.Tools), req.Extra != nil)
}

// replayCached rejoue une réponse en cache comme si elle venait du provider :
// mêmes deltas (raisonnement, contenu, stats marquées cached) puis message
// assistant, sans appel réseau.
func (e *Engine) replayCached(c *Conversation, epoch int, in TurnInput, ce cache.Entry) {
	if ce.Reasoning != "" && in.Think {
		c.appendDelta(epoch, map[string]any{"reasoning_content": ce.Reasoning})
	}
	if ce.Content != "" {
		c.appendDelta(epoch, map[string]any{"content": ce.Content})
	}
	c.appendDelta(epoch, map[string]any{"stats": map[string]any{
		"prompt_tokens":     ce.Usage.PromptTokens,
		"completion_tokens": ce.Usage.CompletionTokens,
		"cached":            true,
	}})
	c.appendAssistant(epoch, ce.Content, "")
}

// cacheStore mémorise une réponse réussie sous sa clé (clé vide : rien).
func (e *Engine) cacheStore(key string, resp provider.Response) {
	if key == "" || e.exactCache() == nil {
		return
	}
	_ = e.exactCache().Set(key, cache.Entry{
		Content:   resp.Content,
		Reasoning: resp.Reasoning,
		Usage: cache.Usage{
			PromptTokens:     resp.Usage.PromptTokens,
			CompletionTokens: resp.Usage.CompletionTokens,
			TotalTokens:      resp.Usage.TotalTokens,
		},
	})
}

// SetCache branche le cache exact sur le moteur (nil = désactivé).
func (e *Engine) SetCache(c *cache.Cache) {
	e.mu.Lock()
	e.exactCache_ = c
	e.mu.Unlock()
}

func (e *Engine) exactCache() *cache.Cache {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.exactCache_
}

// Cache rend le cache exact branché (nil si désactivé).
func (e *Engine) Cache() *cache.Cache { return e.exactCache() }

// CacheStats rend l'état du cache exact : actif, hits, misses, entrées.
func (e *Engine) CacheStats() (enabled bool, hits, misses uint64, entries int) {
	cc := e.exactCache()
	if cc == nil {
		return false, 0, 0, 0
	}
	h, m, n := cc.Stats()
	return true, h, m, n
}
