package web

import (
	"net/http"

	"cetas-lite/internal/plugins"
)

// handlePluginsStatus : GET /api/plugins -> plugins charges + erreurs.
func (s *Server) handlePluginsStatus(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	infos, errs := s.engine.PluginStatus()
	if infos == nil {
		infos = []plugins.Info{}
	}
	if errs == nil {
		errs = []string{}
	}
	writeJSON(w, http.StatusOK, map[string]any{"plugins": infos, "errors": errs})
}

// handlePluginsReload : POST /api/plugins/reload -> recharge les plugins.
func (s *Server) handlePluginsReload(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	if err := s.engine.PluginReload(); err != nil {
		writeError(w, http.StatusInternalServerError, "rechargement: "+err.Error())
		return
	}
	infos, errs := s.engine.PluginStatus()
	if infos == nil {
		infos = []plugins.Info{}
	}
	if errs == nil {
		errs = []string{}
	}
	writeJSON(w, http.StatusOK, map[string]any{"plugins": infos, "errors": errs})
}
