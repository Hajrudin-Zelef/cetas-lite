package web

import (
	"net/http"

	"cetas-lite/internal/mcp"
)

func (s *Server) handleMCPStatus(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	servers := s.engine.MCPServers()
	if servers == nil {
		servers = []mcp.ServerStatus{}
	}
	writeJSON(w, http.StatusOK, map[string]any{"servers": servers})
}
