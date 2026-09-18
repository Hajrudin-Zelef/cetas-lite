package web

import (
	"net/http"

	"cetas-lite/internal/catalog"
	"cetas-lite/internal/modelinfo"
)

// GET /api/model-info?provider=<id>&model=<id>
// Retourne la fenetre de contexte et les tarifs indicatifs d'un modele.
// La table statique (modelinfo) fournit contexte + tarifs ; si le modele
// n'y figure pas, on retombe sur le catalogue (tarifs seuls, contexte
// inconnu) pour ne pas repondre 404 sur un modele pourtant utilisable.
// Repond 404 uniquement si le modele est inconnu des deux sources.
func (s *Server) handleModelInfo(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	q := r.URL.Query()
	provider, model := q.Get("provider"), q.Get("model")
	if info, ok := modelinfo.Lookup(provider, model); ok {
		writeJSON(w, http.StatusOK, info)
		return
	}
	if m, ok := catalog.Lookup(provider, model); ok {
		writeJSON(w, http.StatusOK, modelinfo.Info{
			InputPer1M:  m.InputPer1M,
			OutputPer1M: m.OutputPer1M,
		})
		return
	}
	writeError(w, http.StatusNotFound, "modele inconnu")
}
