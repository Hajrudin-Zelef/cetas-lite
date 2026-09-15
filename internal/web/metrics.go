package web

import (
	"net/http"

	"cetas-lite/internal/metrics"
)

// GET /api/metrics : metriques systeme (CPU %, RAM, disque, reseau).
func (s *Server) handleMetrics(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	writeJSON(w, http.StatusOK, metrics.Sample())
}
