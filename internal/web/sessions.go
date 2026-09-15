package web

import (
	"net/http"
	"sort"
)

// sessionItem : une session de conversation (chat CETAS ou agent).
type sessionItem struct {
	ID       string `json:"id"`
	Kind     string `json:"kind"` // "chat" | "agent"
	Title    string `json:"title"`
	Updated  int64  `json:"updated"`
	Messages int    `json:"messages,omitempty"`
	Status   string `json:"status,omitempty"` // agents : running | done | stopped
	Family   string `json:"family,omitempty"`
	Mode     string `json:"mode,omitempty"`
}

// GET /api/sessions — historique unifie des conversations : chat CETAS
// (archives) et agents. La suppression reuse les routes existantes
// (DELETE /api/conversations/{id}, DELETE /api/agents/{id}).
func (s *Server) handleSessionsList(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	chat := make([]sessionItem, 0)
	for _, a := range s.engine.ListArchives(claims.Username) {
		title := a.Title
		if title == "" {
			title = "Conversation sans titre"
		}
		chat = append(chat, sessionItem{
			ID: a.ID, Kind: "chat", Title: title,
			Updated: a.Updated, Messages: a.Messages,
		})
	}
	agents := make([]sessionItem, 0)
	for _, a := range s.engine.ListAgents(claims.Username) {
		title := a.Preview
		if title == "" {
			title = "Agent " + a.ID
		}
		agents = append(agents, sessionItem{
			ID: a.ID, Kind: "agent", Title: title,
			Updated: a.Created, Status: a.Status,
			Family: a.Family, Mode: a.Mode,
		})
	}
	sort.Slice(chat, func(i, j int) bool { return chat[i].Updated > chat[j].Updated })
	sort.Slice(agents, func(i, j int) bool { return agents[i].Updated > agents[j].Updated })
	writeJSON(w, http.StatusOK, map[string]any{"chat": chat, "agents": agents})
}
