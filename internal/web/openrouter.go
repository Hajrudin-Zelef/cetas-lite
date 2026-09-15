package web

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Catalogue OpenRouter en direct.
//
// GET /api/openrouter/models?type=text|image&refresh=1
//
// L'endpoint public https://openrouter.ai/api/v1/models ne requiert aucune
// cle : le serveur le recupere et le met en cache 1h (memoire). Seuls les
// champs utiles au catalogue sont projetes. `refresh=1` force le
// rechargement (bouton Actualiser du panneau API et Modeles).

type orCatalogModel struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description,omitempty"`
	ContextLength   int     `json:"context_length,omitempty"`
	PromptPer1M     float64 `json:"prompt_per_1m"`
	CompletionPer1M float64 `json:"completion_per_1m"`
	IsImage         bool    `json:"is_image"`
}

type orModelsCache struct {
	mu     sync.Mutex
	at     time.Time
	models []orCatalogModel
}

var orCache = &orModelsCache{}

var orModelsURL = "https://openrouter.ai/api/v1/models"

const orCacheTTL = time.Hour

var errORBadStatus = errors.New("statut inattendu du catalogue OpenRouter")

func (s *Server) handleOpenRouterModels(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	wantImage := r.URL.Query().Get("type") == "image"
	refresh := r.URL.Query().Get("refresh") == "1"

	models, cachedAt, err := s.fetchOpenRouterModels(refresh)
	if err != nil {
		writeError(w, http.StatusBadGateway, "catalogue OpenRouter indisponible")
		return
	}
	out := make([]orCatalogModel, 0, len(models))
	for _, m := range models {
		if m.IsImage == wantImage {
			out = append(out, m)
		}
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"models":    out,
		"count":     len(out),
		"cached_at": cachedAt.UTC().Format(time.RFC3339),
	})
}

func (s *Server) fetchOpenRouterModels(refresh bool) ([]orCatalogModel, time.Time, error) {
	orCache.mu.Lock()
	defer orCache.mu.Unlock()
	if !refresh && time.Since(orCache.at) < orCacheTTL && orCache.models != nil {
		return orCache.models, orCache.at, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, orModelsURL, nil)
	if err != nil {
		return nil, time.Time{}, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, time.Time{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, time.Time{}, errORBadStatus
	}
	var payload struct {
		Data []struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			Description  string `json:"description"`
			Context      int    `json:"context_length"`
			Architecture struct {
				InputModalities  []string `json:"input_modalities"`
				OutputModalities []string `json:"output_modalities"`
			} `json:"architecture"`
			Pricing struct {
				Prompt     string `json:"prompt"`
				Completion string `json:"completion"`
			} `json:"pricing"`
		} `json:"data"`
	}
	dec := json.NewDecoder(io.LimitReader(resp.Body, 8<<20))
	if err := dec.Decode(&payload); err != nil {
		return nil, time.Time{}, err
	}
	models := make([]orCatalogModel, 0, len(payload.Data))
	for _, d := range payload.Data {
		if d.ID == "" || d.Name == "" {
			continue
		}
		isImage := false
		for _, m := range d.Architecture.OutputModalities {
			if m == "image" {
				isImage = true
				break
			}
		}
		models = append(models, orCatalogModel{
			ID:              d.ID,
			Name:            d.Name,
			Description:     d.Description,
			ContextLength:   d.Context,
			PromptPer1M:     orPricePer1M(d.Pricing.Prompt),
			CompletionPer1M: orPricePer1M(d.Pricing.Completion),
			IsImage:         isImage,
		})
	}
	orCache.models = models
	orCache.at = time.Now()
	return models, orCache.at, nil
}

func orPricePer1M(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil || f < 0 {
		return 0
	}
	return f * 1e6
}
