package web

import (
	"errors"
	"net/http"
	"strconv"

	"cetas-lite/internal/chat"
)

// POST /api/agents — cree un agent parallele et demarre son premier tour.
func (s *Server) handleAgentsCreate(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		Family      string   `json:"family"`
		Mode        string   `json:"mode"`
		Message     string   `json:"message"`
		Repo        string   `json:"repo"`
		ProjectID   string   `json:"project_id"`
		Worktree    bool     `json:"worktree"`
		Approve     bool     `json:"approve"`
		Plan        bool     `json:"plan"`
		Web         bool     `json:"web"`
		WebDepth    string   `json:"web_depth"`
		Think       *bool    `json:"think"`
		Effort      string   `json:"effort"`
		Attachments []string `json:"attachments"`
		// ClientMsgID : echo optimiste du message (voir chat.TurnInput).
		ClientMsgID string `json:"client_msg_id"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	if body.Family == "" || body.Mode == "" {
		writeError(w, http.StatusBadRequest, "family et mode requis")
		return
	}
	think := body.Think == nil || *body.Think
	run, err := s.engine.SpawnAgent(claims.Username, chat.TurnInput{
		User: claims.Username, Family: body.Family, Mode: body.Mode,
		Text: body.Message, Approve: body.Approve, Plan: body.Plan,
		Web: body.Web, WebDepth: body.WebDepth, Think: think, Effort: body.Effort,
		AgentMode: true, Attachments: body.Attachments,
		Worktree: body.Worktree, Repo: body.Repo, ProjectID: body.ProjectID,
		ClientMsgID: body.ClientMsgID,
		MaxTokens:   chat.ClampMaxTokens(s.storedSettings(claims.Username).MaxTokens),
	})
	if err != nil {
		switch {
		case errors.Is(err, chat.ErrBadMessage):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusBadRequest, err.Error())
		}
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"id": run.ID})
}

// GET /api/agents — liste les runs de l'utilisateur.
func (s *Server) handleAgentsList(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	list := s.engine.ListAgents(claims.Username)
	if list == nil {
		list = []chat.AgentSummary{}
	}
	writeJSON(w, http.StatusOK, map[string]any{"agents": list})
}

// GET /api/agents/{id}/stream — flux SSE du run.
func (s *Server) handleAgentStream(w http.ResponseWriter, r *http.Request) {
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
	if !s.engine.AgentSubscribe(r.Context(), claims.Username, r.PathValue("id"), from, emit) {
		// La reponse SSE a deja commence ; signaler l'erreur dans le flux.
		emit(map[string]any{"error": "agent introuvable"})
	}
}

// GET /api/agents/{id}/state — etat du run.
func (s *Server) handleAgentState(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	st := s.engine.AgentStateInfo(claims.Username, r.PathValue("id"))
	if st == nil {
		writeError(w, http.StatusNotFound, "agent introuvable")
		return
	}
	writeJSON(w, http.StatusOK, st)
}

// POST /api/agents/{id}/message — message de suivi.
func (s *Server) handleAgentMessage(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		Message     string   `json:"message"`
		Approve     bool     `json:"approve"`
		Plan        bool     `json:"plan"`
		Web         bool     `json:"web"`
		WebDepth    string   `json:"web_depth"`
		Think       *bool    `json:"think"`
		Effort      string   `json:"effort"`
		ProjectID   string   `json:"project_id"`
		Attachments []string `json:"attachments"`
		// ClientMsgID : echo optimiste du message (voir chat.TurnInput).
		ClientMsgID string `json:"client_msg_id"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	think := body.Think == nil || *body.Think
	err := s.engine.MessageAgent(claims.Username, r.PathValue("id"), chat.TurnInput{
		User: claims.Username, Text: body.Message, Approve: body.Approve, Plan: body.Plan,
		Web: body.Web, WebDepth: body.WebDepth, Think: think, Effort: body.Effort,
		AgentMode: true, ProjectID: body.ProjectID, Attachments: body.Attachments,
		ClientMsgID: body.ClientMsgID,
		MaxTokens:   chat.ClampMaxTokens(s.storedSettings(claims.Username).MaxTokens),
	})
	if err != nil {
		switch {
		case errors.Is(err, chat.ErrBusy):
			writeError(w, http.StatusConflict, "generation en cours")
		case errors.Is(err, chat.ErrBadMessage):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusBadRequest, err.Error())
		}
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

// POST /api/agents/{id}/stop — interrompt le tour en cours.
func (s *Server) handleAgentStop(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	if !s.engine.StopAgent(claims.Username, r.PathValue("id")) {
		writeError(w, http.StatusNotFound, "agent introuvable")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

// POST /api/agents/{id}/approve — decision d'approbation pour le run.
func (s *Server) handleAgentApprove(w http.ResponseWriter, r *http.Request) {
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
	if !s.engine.ResolveAgentApproval(claims.Username, r.PathValue("id"), body.ID, body.Approved, body.Always) {
		writeError(w, http.StatusGone, "demande d'approbation introuvable ou expiree")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

// DELETE /api/agents/{id} — supprime le run (arrete, nettoie le worktree).
func (s *Server) handleAgentDelete(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	if !s.engine.DeleteAgent(claims.Username, r.PathValue("id")) {
		writeError(w, http.StatusNotFound, "agent introuvable")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}
