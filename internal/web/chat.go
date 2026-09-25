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
		MaxTokens   int      `json:"max_tokens"`
		// FocusCorpus restreint la recherche RAG à un corpus (clic sur
		// une question suggérée). Transmis tel quel au moteur.
		FocusCorpus string `json:"focus_corpus"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	if body.Family == "" || body.Mode == "" {
		writeError(w, http.StatusBadRequest, "family et mode requis")
		return
	}
	maxTokens := chat.ClampMaxTokens(body.MaxTokens)
	settings := s.storedSettings(claims.Username)
	if maxTokens == 0 {
		maxTokens = chat.ClampMaxTokens(settings.MaxTokens)
	}
	// Lot 5 (pré-génération) : un message tapé annule les fantômes en
	// cours (l'historique change, leurs réponses ne serviraient plus) ;
	// un clic sur une question suggérée (focus_corpus) les consomme.
	if body.FocusCorpus == "" {
		s.engine.PregenCancel(claims.Username)
	}
	c := s.engine.Conversation(claims.Username)
	err := c.StartTurn(chat.TurnInput{User: claims.Username, Family: body.Family, Mode: body.Mode, Text: body.Message, Web: body.Web, MCP: body.MCP, Think: body.Think, Effort: body.Effort, Approve: body.Approve, Plan: body.Plan, AgentMode: body.AgentMode, Worktree: body.Worktree, Repo: body.Repo, Attachments: body.Attachments, MaxTokens: maxTokens, FocusCorpus: body.FocusCorpus, PregenLookup: body.FocusCorpus != "" && settings.pregenEnabled()})
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
		Comment  string `json:"comment"`
	}
	if err := decodeJSON(r, &body); err != nil || body.ID == "" {
		writeError(w, http.StatusBadRequest, "id requis")
		return
	}
	ok := s.engine.Conversation(claims.Username).ResolveApproval(body.ID, body.Approved, body.Always, body.Comment)
	if !ok {
		writeError(w, http.StatusGone, "demande d'approbation introuvable ou expiree")
		return
	}
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

// handleChatSuggestions sert le pool de questions suggérées
// (suggested-questions.yaml). Le tirage (3 chips, anti-répétition,
// mixité, épinglés) est effectué côté navigateur.
func (s *Server) handleChatSuggestions(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	suggs := s.engine.Suggestions()
	if suggs == nil {
		suggs = []chat.Suggestion{}
	}
	writeJSON(w, http.StatusOK, map[string]any{"suggestions": suggs})
}

// handleChatPrefetch reçoit les questions suggérées affichées en chips
// et lance leur pré-génération en arrière-plan (lot 5). Le navigateur
// notifie après chaque tirage ; le serveur rejoue les 3 tours sur des
// conversations fantômes. 202 immédiat : la génération est asynchrone.
// Fail-open : réglage désactivé ou requête invalide => {"ok":true} sans
// effet, le clic suivra le chemin normal.
func (s *Server) handleChatPrefetch(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		Suggestions []struct {
			Question string `json:"question"`
			Corpus   string `json:"corpus"`
		} `json:"suggestions"`
		Family    string `json:"family"`
		Mode      string `json:"mode"`
		Web       bool   `json:"web"`
		MCP       bool   `json:"mcp"`
		Think     bool   `json:"think"`
		Effort    string `json:"effort"`
		MaxTokens int    `json:"max_tokens"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
		return
	}
	settings := s.storedSettings(claims.Username)
	if !settings.pregenEnabled() || body.Family == "" || body.Mode == "" || len(body.Suggestions) == 0 {
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
		return
	}
	maxTokens := chat.ClampMaxTokens(body.MaxTokens)
	if maxTokens == 0 {
		maxTokens = chat.ClampMaxTokens(settings.MaxTokens)
	}
	items := make([]chat.PregenSuggestion, 0, len(body.Suggestions))
	for _, sg := range body.Suggestions {
		if sg.Question == "" {
			continue
		}
		items = append(items, chat.PregenSuggestion{Question: sg.Question, Corpus: sg.Corpus})
	}
	if len(items) == 0 {
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
		return
	}
	history := s.engine.Conversation(claims.Username).MessagesSnapshot()
	s.engine.PregenSuggestions(claims.Username, items, chat.TurnInput{
		User: claims.Username, Family: body.Family, Mode: body.Mode,
		Web: body.Web, MCP: body.MCP, Think: body.Think, Effort: body.Effort,
		MaxTokens: maxTokens,
	}, history)
	writeJSON(w, http.StatusAccepted, map[string]any{"ok": true})
}
