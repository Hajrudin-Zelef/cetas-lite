package web

import (
	"net/http"

	"cetas-lite/internal/modelcaps"
)

func (s *Server) handleCapabilitiesGet(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"caps": s.engine.Capabilities()})
}

func (s *Server) handleCapabilitiesPut(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		Caps modelcaps.Map `json:"caps"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	if body.Caps == nil {
		body.Caps = modelcaps.Map{}
	}
	if err := modelcaps.Save(s.st, body.Caps); err != nil {
		writeError(w, http.StatusInternalServerError, "erreur serveur")
		return
	}
	s.engine.SetCapabilities(body.Caps)
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}
