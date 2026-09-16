package web

import (
	"context"
	"net/http"
	"time"
)

// Modeles decouverts sur les moteurs locaux.
//
// GET /api/local/models?engine=ollama -> {"engine": "ollama", "models": ["..."]}
//
// Utilise par le selecteur de modeles (famille SamGen) : l'utilisateur
// choisit 1 modele ou un ensemble en fallback parmi ceux decouverts.

func (s *Server) handleLocalModels(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	engine := r.URL.Query().Get("engine")
	if !isLocalEngine(engine) {
		writeError(w, http.StatusBadRequest, "moteur local inconnu")
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()
	models := s.engine.LocalModels(ctx, engine)
	if models == nil {
		models = []string{}
	}
	writeJSON(w, http.StatusOK, map[string]any{"engine": engine, "models": models})
}
