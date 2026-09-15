package web

import (
	"net/http"

	"cetas-lite/internal/modelinfo"
)

// GET /api/model-info?provider=<id>&model=<id>
// Retourne la fenetre de contexte et les tarifs indicatifs d'un modele.
// Repond 404 si le modele est inconnu.
func (s *Server) handleModelInfo(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	q := r.URL.Query()
	info, ok := modelinfo.Lookup(q.Get("provider"), q.Get("model"))
	if !ok {
		writeError(w, http.StatusNotFound, "modele inconnu")
		return
	}
	writeJSON(w, http.StatusOK, info)
}
