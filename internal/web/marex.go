package web

import (
	"net/http"
)

// GET /api/marex — lit le fichier MAREX.md de l'utilisateur.
// PUT /api/marex — enregistre le fichier MAREX.md (corps {"content": "..."}).
func (s *Server) handleMarex(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	switch r.Method {
	case http.MethodGet:
		content, err := s.engine.ReadMarex()
		if err != nil {
			writeError(w, http.StatusInternalServerError, "lecture impossible")
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{
			"path":    s.engine.MarexPath(),
			"content": content,
		})
	case http.MethodPut:
		var body struct {
			Content string `json:"content"`
		}
		if err := decodeJSON(r, &body); err != nil {
			writeError(w, http.StatusBadRequest, "corps JSON invalide")
			return
		}
		if err := s.engine.WriteMarex(body.Content); err != nil {
			writeError(w, http.StatusInternalServerError, "ecriture impossible")
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
	default:
		writeError(w, http.StatusMethodNotAllowed, "methode non supportee")
	}
}
