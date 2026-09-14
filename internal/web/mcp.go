package web

import (
	"context"
	"net/http"
	"time"

	"cetas-lite/internal/mcp"
)

func (s *Server) handleMCPStatus(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var servers []mcp.ServerStatus
	if r.URL.Query().Get("probe") == "1" {
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		servers = s.engine.MCPProbe(ctx)
	} else {
		servers = s.engine.MCPServers()
	}
	if servers == nil {
		servers = []mcp.ServerStatus{}
	}
	writeJSON(w, http.StatusOK, map[string]any{"servers": servers})
}
