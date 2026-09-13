package web

import (
	"encoding/json"
	"net/http"

	"cetas-lite/internal/alias"
)

func (s *Server) handleAliasesGet(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{"families": alias.ResolveAll(s.engine.Families())})
}

func (s *Server) handleAliasesPut(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var ov alias.Overrides
	if err := decodeJSON(r, &ov); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	fams := alias.Apply(s.engine.Families(), ov)
	s.engine.SetFamilies(fams)
	if raw, err := json.Marshal(ov); err == nil {
		_ = s.st.PutMeta("aliases", raw)
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "families": alias.ResolveAll(fams)})
}
