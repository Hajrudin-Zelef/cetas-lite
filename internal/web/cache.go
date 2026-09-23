package web

import "net/http"

// handleCacheStats expose les métriques du cache exact (itération 5a) :
// hits/misses depuis le démarrage, nombre d'entrées, état d'activation.
func (s *Server) handleCacheStats(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	enabled, hits, misses, entries := s.engine.CacheStats()
	writeJSON(w, http.StatusOK, map[string]any{
		"enabled": enabled,
		"hits":    hits,
		"misses":  misses,
		"entries": entries,
	})
}
