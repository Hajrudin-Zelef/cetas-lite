package web

import (
	"net/http"

	"cetas-lite/internal/catalog"
)

var providerLabels = map[string]string{
	"deepseek":    "DeepSeek",
	"openrouter":  "OpenRouter",
	"opencode":    "OpenCode Zen",
	"opencode-go": "OpenCode Go",
}

var providerOrder = []string{"deepseek", "openrouter", "opencode", "opencode-go"}

type catalogModel struct {
	ID          string  `json:"id"`
	Label       string  `json:"label"`
	InputPer1M  float64 `json:"input_per_1m"`
	OutputPer1M float64 `json:"output_per_1m"`
}

type catalogProvider struct {
	ID     string         `json:"id"`
	Label  string         `json:"label"`
	Models []catalogModel `json:"models"`
}

func (s *Server) handleCatalogGet(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	byProvider := map[string][]catalogModel{}
	for _, m := range catalog.Models() {
		byProvider[m.Editeur] = append(byProvider[m.Editeur], catalogModel{
			ID: m.ID, Label: m.Label, InputPer1M: m.InputPer1M, OutputPer1M: m.OutputPer1M,
		})
	}
	providers := make([]catalogProvider, 0, len(providerOrder))
	for _, id := range providerOrder {
		models := byProvider[id]
		if len(models) == 0 {
			continue
		}
		label := providerLabels[id]
		if label == "" {
			label = id
		}
		providers = append(providers, catalogProvider{ID: id, Label: label, Models: models})
	}
	writeJSON(w, http.StatusOK, map[string]any{"providers": providers})
}
