package web

// Sessions : API REST des conversations.
//
//	GET    /api/sessions          -> {sessions: [{id,title,createdAt,updatedAt,messages}]}
//	POST   /api/sessions          -> {id}                      (nouvelle session, devient courante)
//	GET    /api/sessions/current  -> {id}                      (session courante)
//	POST   /api/sessions/{id}/open -> {id}                     (ouvre la session)
//	DELETE /api/sessions/{id}     -> {ok, currentId}            (suppression définitive)
//
// Chaque session est isolée : un utilisateur ne voit que ses sessions.
// La suppression est immédiate et totale (serveur) ; le client purge son
// cache local sur 200.

import (
	"net/http"

	"cetas-lite/internal/chat"
)

func (s *Server) handleSessionsList(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	sessions := s.engine.ListSessions(claims.Username)
	if sessions == nil {
		sessions = []chat.SessionInfo{}
	}
	writeJSON(w, http.StatusOK, map[string]any{"sessions": sessions})
}

func (s *Server) handleSessionCreate(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	id := s.engine.CreateSession(claims.Username)
	writeJSON(w, http.StatusCreated, map[string]any{"id": id})
}

func (s *Server) handleSessionCurrent(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"id": s.engine.CurrentSessionID(claims.Username)})
}

func (s *Server) handleSessionOpen(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	id := r.PathValue("id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "id requis")
		return
	}
	if !s.engine.OpenSession(claims.Username, id) {
		writeError(w, http.StatusNotFound, "session introuvable")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"id": id})
}

func (s *Server) handleSessionDelete(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	id := r.PathValue("id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "id requis")
		return
	}
	curID, ok := s.engine.DeleteSession(claims.Username, id)
	if !ok {
		writeError(w, http.StatusNotFound, "session introuvable")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "currentId": curID})
}
