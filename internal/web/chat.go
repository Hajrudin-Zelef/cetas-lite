package web

import (
	"errors"
	"net/http"
	"strconv"

	"cetas-lite/internal/chat"
)

func (s *Server) handleChatSend(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		Family      string   `json:"family"`
		Mode        string   `json:"mode"`
		Message     string   `json:"message"`
		Web         bool     `json:"web"`
		MCP         bool     `json:"mcp"`
		Think       bool     `json:"think"`
		Effort      string   `json:"effort"`
		Approve     bool     `json:"approve"`
		Plan        bool     `json:"plan"`
		AgentMode   bool     `json:"agent_mode"`
		Worktree    bool     `json:"worktree"`
		Repo        string   `json:"repo"`
		Attachments []string `json:"attachments"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	if body.Family == "" || body.Mode == "" {
		writeError(w, http.StatusBadRequest, "family et mode requis")
		return
	}
	c := s.engine.Conversation(claims.Username)
	err := c.StartTurn(chat.TurnInput{User: claims.Username, Family: body.Family, Mode: body.Mode, Text: body.Message, Web: body.Web, MCP: body.MCP, Think: body.Think, Effort: body.Effort, Approve: body.Approve, Plan: body.Plan, AgentMode: body.AgentMode, Worktree: body.Worktree, Repo: body.Repo, Attachments: body.Attachments})
	if errors.Is(err, chat.ErrBusy) {
		writeError(w, http.StatusConflict, "generation en cours")
		return
	}
	if errors.Is(err, chat.ErrBadMessage) {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

func (s *Server) handleChatStream(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	from := 0
	if raw := r.URL.Query().Get("from"); raw != "" {
		if v, err := strconv.Atoi(raw); err == nil && v >= 0 {
			from = v
		}
	}
	mu, stop, ok := sseStart(w)
	if !ok {
		return
	}
	defer stop()
	emit := sseEmitter(w, mu)
	c := s.engine.Conversation(claims.Username)
	c.Subscribe(r.Context(), from, emit)
}

func (s *Server) handleChatStop(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	s.engine.Conversation(claims.Username).Stop()
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

// handleChatApprove transmet la decision de l'utilisateur pour une demande
// d'approbation en attente (outil sensible ou plan).
func (s *Server) handleChatApprove(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		ID       string `json:"id"`
		Approved bool   `json:"approved"`
		Always   bool   `json:"always"`
	}
	if err := decodeJSON(r, &body); err != nil || body.ID == "" {
		writeError(w, http.StatusBadRequest, "id requis")
		return
	}
	ok := s.engine.Conversation(claims.Username).ResolveApproval(body.ID, body.Approved, body.Always)
	if !ok {
		writeError(w, http.StatusGone, "demande d'approbation introuvable ou expiree")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

func (s *Server) handleChatReset(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	s.engine.ArchiveAndReset(claims.Username)
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

func (s *Server) handleChatRegenerate(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	err := s.engine.Regenerate(claims.Username)
	switch {
	case err == nil:
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
	case errors.Is(err, chat.ErrBusy):
		writeError(w, http.StatusConflict, "generation en cours")
	case errors.Is(err, chat.ErrNoTurn):
		writeError(w, http.StatusBadRequest, "aucun tour a regenerer")
	default:
		writeError(w, http.StatusInternalServerError, err.Error())
	}
}

func (s *Server) handleChatState(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	writeJSON(w, http.StatusOK, s.engine.Conversation(claims.Username).State())
}

func (s *Server) handleConversationDelete(w http.ResponseWriter, r *http.Request) {
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
	if !s.engine.DeleteArchive(claims.Username, id) {
		writeError(w, http.StatusNotFound, "archive introuvable")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

func (s *Server) handleConversationsList(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"archives": s.engine.ListArchives(claims.Username),
	})
}

func (s *Server) handleConversationRestore(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		ID string `json:"id"`
	}
	if err := decodeJSON(r, &body); err != nil || body.ID == "" {
		writeError(w, http.StatusBadRequest, "id requis")
		return
	}
	if !s.engine.RestoreArchive(claims.Username, body.ID) {
		writeError(w, http.StatusNotFound, "archive introuvable")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}
