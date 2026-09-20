package chat

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"sync"
	"time"

	"cetas-lite/internal/provider"
	"cetas-lite/internal/worktree"
)

// AgentRun represente un agent qui tourne en parallele : sa propre
// conversation (donc son propre tour, ses approbations, son SSE) et,
// si demande, son propre worktree git isole.
type AgentRun struct {
	ID        string
	User      string
	Family    string
	Mode      string
	Repo      string
	ProjectID string
	Created   time.Time
	conv      *Conversation
	convID    string
	mu        sync.Mutex
	wtPath    string
	restored  bool
}

// WorktreePath retourne le chemin du worktree associe ("" si aucun).
func (r *AgentRun) WorktreePath() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.wtPath
}

func (r *AgentRun) setWorktreePath(p string) {
	r.mu.Lock()
	r.wtPath = p
	r.mu.Unlock()
}

// AgentSummary est la vue liste d'un run pour l'UI.
type AgentSummary struct {
	ID       string `json:"id"`
	Status   string `json:"status"` // running | done | stopped
	Family   string `json:"family"`
	Mode     string `json:"mode"`
	Repo     string `json:"repo,omitempty"`
	Worktree string `json:"worktree,omitempty"`
	Preview  string `json:"preview,omitempty"`
	Created  int64  `json:"created"`
	Restored bool   `json:"restored,omitempty"`
}

func agentStoreKey(runID string) string { return "agent:" + runID }

func isAgentStoreKey(k string) bool { return strings.HasPrefix(k, "agent:") }

// SetWorktreeManager branche le gestionnaire de worktrees sur le moteur.
func (e *Engine) SetWorktreeManager(wm *worktree.Manager) {
	e.mu.Lock()
	e.worktrees = wm
	e.mu.Unlock()
}

func (e *Engine) worktreeManager() *worktree.Manager {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.worktrees
}

// SpawnAgent cree un agent parallele et demarre son premier tour.
// La famille doit etre une famille agent.
func (e *Engine) SpawnAgent(user string, in TurnInput) (*AgentRun, error) {
	user = strings.TrimSpace(user)
	if user == "" {
		return nil, errors.New("utilisateur requis")
	}
	if strings.TrimSpace(in.Text) == "" {
		return nil, ErrBadMessage
	}
	in.AgentMode = true // un agent parallele tourne toujours en mode agent
	res := e.resolve(context.Background(), in)
	if len(res.members) == 0 {
		return nil, errors.New("aucun modele disponible pour cet alias")
	}
	if !res.agent {
		return nil, errors.New("la famille choisie n'est pas une famille agent")
	}
	in.User = user
	run := &AgentRun{
		ID:        newID(),
		User:      user,
		Family:    in.Family,
		Mode:      in.Mode,
		Repo:      strings.TrimSpace(in.Repo),
		ProjectID: strings.TrimSpace(in.ProjectID),
		Created:   time.Now(),
	}
	run.convID = newID()
	run.conv = NewConversation(run.convID, e, func(c *Conversation) { e.saveAgent(run, c) })

	e.agentsMu.Lock()
	if e.agents == nil {
		e.agents = map[string]*AgentRun{}
	}
	e.agents[run.ID] = run
	e.agentsMu.Unlock()

	if err := run.conv.StartTurn(in); err != nil {
		e.agentsMu.Lock()
		delete(e.agents, run.ID)
		e.agentsMu.Unlock()
		return nil, err
	}
	return run, nil
}

// agentForConv retrouve le run associe a une conversation (pour y
// enregistrer le chemin du worktree cree dans runAgent).
func (e *Engine) agentForConv(convID string) *AgentRun {
	e.agentsMu.Lock()
	defer e.agentsMu.Unlock()
	for _, r := range e.agents {
		if r.convID == convID {
			return r
		}
	}
	return nil
}

// GetAgent retourne le run, en le restaurant depuis le store si besoin.
func (e *Engine) GetAgent(user, id string) *AgentRun {
	e.agentsMu.Lock()
	if r, ok := e.agents[id]; ok && r.User == user {
		e.agentsMu.Unlock()
		return r
	}
	e.agentsMu.Unlock()
	return e.restoreAgent(user, id)
}

// restoreAgent recharge un run persiste (apres redemarrage) sans relancer
// de tour : statut "stopped", reprise possible par message.
func (e *Engine) restoreAgent(user, id string) *AgentRun {
	if e.st == nil {
		log.Printf("chat: restoreAgent %s: pas de store", id)
		return nil
	}
	if agentDeleted(user, id) {
		return nil
	}
	raw, ok := e.st.GetConversation(user, agentStoreKey(id))
	if !ok {
		log.Printf("chat: restoreAgent %s: cle absente du store", id)
		return nil
	}
	var s snapshot
	if err := json.Unmarshal(raw, &s); err != nil || s.ID == "" {
		log.Printf("chat: restoreAgent %s: snapshot illisible (err=%v)", id, err)
		return nil
	}
	run := &AgentRun{ID: id, User: user, Created: time.Now(), restored: true}
	if s.Turn != nil {
		run.Family = s.Turn.Family
		run.Mode = s.Turn.Mode
		// Conserver le depot : le worktree associe (derive du convID) doit
		// pouvoir etre nettoye meme apres un redemarrage.
		run.Repo = strings.TrimSpace(s.Turn.Repo)
	}
	run.convID = s.ID
	run.conv = NewConversation(s.ID, e, func(c *Conversation) { e.saveAgent(run, c) })
	run.conv.restore(s)
	e.agentsMu.Lock()
	if e.agents == nil {
		e.agents = map[string]*AgentRun{}
	}
	// Ne pas ecraser un run recemment cree en memoire.
	if existing, ok := e.agents[id]; ok {
		e.agentsMu.Unlock()
		if existing.User == user {
			return existing
		}
		return nil
	}
	e.agents[id] = run
	e.agentsMu.Unlock()
	return run
}

func (e *Engine) saveAgent(run *AgentRun, c *Conversation) {
	if e.st == nil {
		return
	}
	// Un agent supprime ne doit jamais etre ressuscite par la persistance
	// differee d'un tour encore en vol (meme pierre tombale que les sessions).
	if agentDeleted(run.User, run.ID) {
		return
	}
	data, err := c.save()
	if err != nil {
		log.Printf("chat: saveAgent %s: marshal impossible: %v", run.ID, err)
		return
	}
	if err := e.st.PutConversation(run.User, agentStoreKey(run.ID), data); err != nil {
		log.Printf("chat: saveAgent %s: ecriture store impossible: %v", run.ID, err)
	}
}

// Pierre tombale d'agent : cle prefixee pour ne jamais entrer en collision
// avec les ids de session dans la map partagee.
func agentDeleted(user, id string) bool { return isTombstoned(user, "agent:"+id) }

func tombstoneAgent(user, id string) {
	sessMu.Lock()
	tombstoneLocked(user, "agent:"+id)
	sessMu.Unlock()
}

// ListAgents retourne les runs de l'utilisateur : en memoire d'abord,
// puis ceux connus uniquement via le store (statut "stopped").
func (e *Engine) ListAgents(user string) []AgentSummary {
	e.agentsMu.Lock()
	seen := map[string]bool{}
	var out []AgentSummary
	for _, r := range e.agents {
		if r.User != user {
			continue
		}
		seen[r.ID] = true
		out = append(out, r.summary())
	}
	e.agentsMu.Unlock()

	if e.st != nil {
		if ids, err := e.st.ListConversations(user); err == nil {
			for _, k := range ids {
				if !isAgentStoreKey(k) {
					continue
				}
				id := strings.TrimPrefix(k, "agent:")
				if seen[id] || agentDeleted(user, id) {
					continue
				}
				if s := e.storedAgentSummary(user, k, id); s != nil {
					out = append(out, *s)
				}
			}
		}
	}
	return out
}

func (e *Engine) storedAgentSummary(user, key, id string) *AgentSummary {
	raw, ok := e.st.GetConversation(user, key)
	if !ok {
		return nil
	}
	var s snapshot
	if err := json.Unmarshal(raw, &s); err != nil || s.ID == "" {
		return nil
	}
	sum := &AgentSummary{ID: id, Status: "stopped", Restored: true}
	if s.Turn != nil {
		sum.Family = s.Turn.Family
		sum.Mode = s.Turn.Mode
		sum.Repo = strings.TrimSpace(s.Turn.Repo)
	}
	sum.Preview = lastAssistantPreview(s.Messages)
	return sum
}

func (r *AgentRun) summary() AgentSummary {
	status := "done"
	if r.conv.IsGenerating() {
		status = "running"
	} else if r.restored {
		status = "stopped"
	}
	return AgentSummary{
		ID:       r.ID,
		Status:   status,
		Family:   r.Family,
		Mode:     r.Mode,
		Repo:     r.Repo,
		Worktree: r.WorktreePath(),
		Preview:  r.preview(),
		Created:  r.Created.Unix(),
		Restored: r.restored,
	}
}

func (r *AgentRun) preview() string {
	msgs := r.conv.MessagesSnapshot()
	return lastAssistantPreview(msgs)
}

func lastAssistantPreview(msgs []provider.Message) string {
	for i := len(msgs) - 1; i >= 0; i-- {
		if msgs[i].Role != "assistant" {
			continue
		}
		s, ok := msgs[i].Content.(string)
		if !ok {
			continue
		}
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		return truncateRunes(s, 160)
	}
	return ""
}

func truncateRunes(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n]) + "…"
}

// StopAgent interrompt le tour en cours du run.
func (e *Engine) StopAgent(user, id string) bool {
	r := e.GetAgent(user, id)
	if r == nil {
		return false
	}
	r.conv.Stop()
	return true
}

// MessageAgent envoie un message de suivi au run (refuse si un tour tourne).
func (e *Engine) MessageAgent(user, id string, in TurnInput) error {
	r := e.GetAgent(user, id)
	if r == nil {
		return errors.New("agent introuvable")
	}
	in.User = user
	in.AgentMode = true // un agent parallele tourne toujours en mode agent
	// Conserver la configuration d'origine si non precisee.
	if in.Family == "" {
		in.Family = r.Family
	}
	if in.Mode == "" {
		in.Mode = r.Mode
	}
	if in.Repo == "" {
		in.Repo = r.Repo
	}
	if in.ProjectID == "" {
		in.ProjectID = r.ProjectID
	}
	return r.conv.StartTurn(in)
}

// ResolveAgentApproval transmet une decision d'approbation au run.
func (e *Engine) ResolveAgentApproval(user, agentID, approvalID string, approved, always bool, comment string) bool {
	r := e.GetAgent(user, agentID)
	if r == nil {
		return false
	}
	return r.conv.ResolveApproval(approvalID, approved, always, comment)
}

// DeleteAgent arrete le run, supprime son worktree et ses donnees.
func (e *Engine) DeleteAgent(user, id string) bool {
	e.agentsMu.Lock()
	r, ok := e.agents[id]
	if ok && r.User != user {
		e.agentsMu.Unlock()
		return false
	}
	if ok {
		delete(e.agents, id)
	}
	e.agentsMu.Unlock()
	// Pierre tombale AVANT l'arret/effacement : si un tour est encore en vol,
	// sa persistance differee (saveAgent) sera ignoree.
	tombstoneAgent(user, id)
	if !ok {
		// Run connu uniquement via le store : supprimer la cle si elle existe.
		if e.st != nil {
			if _, found := e.st.GetConversation(user, agentStoreKey(id)); found {
				_ = e.st.DeleteConversation(user, agentStoreKey(id))
				return true
			}
		}
		return false
	}
	r.conv.Stop()
	r.conv.failPendingApprovals()
	if wm := e.worktreeManager(); wm != nil && r.Repo != "" {
		_ = wm.Remove(r.Repo, r.convID)
	}
	if e.st != nil {
		_ = e.st.DeleteConversation(user, agentStoreKey(id))
	}
	return true
}

// releaseConversationWorktree supprime le worktree associe a une
// conversation (chat principal ou agent).
func (e *Engine) releaseConversationWorktree(convID, repo string) {
	if repo == "" {
		return
	}
	if wm := e.worktreeManager(); wm != nil {
		_ = wm.Remove(repo, convID)
	}
}

// convWorktreeRepo memorise le depot utilise pour le worktree d'une
// conversation, afin de le nettoyer a la reinitialisation/suppression.
func (e *Engine) noteWorktreeRepo(convID, repo string) {
	e.agentsMu.Lock()
	if e.wtRepos == nil {
		e.wtRepos = map[string]string{}
	}
	e.wtRepos[convID] = repo
	e.agentsMu.Unlock()
}

func (e *Engine) takeWorktreeRepo(convID string) string {
	e.agentsMu.Lock()
	repo := e.wtRepos[convID]
	delete(e.wtRepos, convID)
	e.agentsMu.Unlock()
	return repo
}

// AgentStateInfo retourne l'etat du run pour l'UI (nil si introuvable).
func (e *Engine) AgentStateInfo(user, id string) map[string]any {
	r := e.GetAgent(user, id)
	if r == nil {
		return nil
	}
	st := r.conv.State()
	sum := r.summary()
	st["agent"] = map[string]any{
		"id":       sum.ID,
		"status":   sum.Status,
		"family":   sum.Family,
		"mode":     sum.Mode,
		"repo":     sum.Repo,
		"worktree": sum.Worktree,
		"created":  sum.Created,
	}
	return st
}

// AgentSubscribe abonne un emetteur SSE au flux du run.
func (e *Engine) AgentSubscribe(ctx context.Context, user, id string, from int, emit func(map[string]any) bool) bool {
	r := e.GetAgent(user, id)
	if r == nil {
		return false
	}
	r.conv.Subscribe(ctx, from, emit)
	return true
}
