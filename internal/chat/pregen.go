// Pré-génération des questions suggérées (lot 5).
//
// Au « salut » (diversion d'accueil, itération 4), le navigateur affiche 3
// chips et notifie le serveur (POST /api/chat/prefetch). Le serveur rejoue
// alors, en arrière-plan, le tour exact qu'un clic produirait — sur des
// conversations fantômes — et met les réponses en cache. Au clic, le tour
// normal retrouve la réponse prête et la rejoue en ~1 ms au lieu de
// 1-10 s.
//
// Garanties (devise : stabilité d'abord) :
//   - Zéro duplication de logique : le fantôme exécute le vrai Run ; la
//     clé de clé (requête provider) n'est servie que si le tour du clic
//     aurait construit la requête à l'octet près (vérification reqKey).
//   - Fail-open total : échec, annulation, mode agent, divergence de
//     requête => le clic suit le chemin normal, sans erreur visible.
//   - Borné : 3 fantômes max par pré-génération, timeout par fantôme,
//     annulation si l'utilisateur envoie un message avant la fin.
//   - Pas de double facturation : un clic pendant la pré-génération
//     attend le fantôme en cours au lieu de relancer un appel.
package chat

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sync"
	"time"

	"cetas-lite/internal/cache"
	"cetas-lite/internal/provider"
)

const (
	// pregenMaxItems : au plus 3 réponses pré-générées (les 3 chips).
	pregenMaxItems = 3
	// pregenShadowTimeout : borne d'un tour fantôme.
	pregenShadowTimeout = 90 * time.Second
	// pregenWaitTimeout : attente max d'un fantôme en cours au clic
	// (couvre le timeout fantôme + marge, pour ne jamais doubler l'appel).
	pregenWaitTimeout = 100 * time.Second
	// pregenEntryTTL : durée de vie d'une entrée servable.
	pregenEntryTTL = 15 * time.Minute
	// pregenMaxEntries : borne du dépôt (au-delà, les plus anciennes
	// terminées sont évincées ; les fantômes en cours ne le sont jamais).
	pregenMaxEntries = 32
)

// PregenSuggestion : une question affichée en chip, à pré-générer.
type PregenSuggestion struct {
	Question string
	Corpus   string
}

// pregenEntry : une réponse pré-générée (ou en cours).
// Les écritures ont lieu avant close(done) ; les lectures après <-done.
type pregenEntry struct {
	done      chan struct{}
	createdAt time.Time
	completed bool
	content   string
	reasoning string
	usage     cache.Usage
	reqKey    string
	err       error
}

// pregenStore : dépôt borné d'entrées, adressé par clé d'entrée.
type pregenStore struct {
	mu      sync.Mutex
	entries map[string]*pregenEntry
	order   []string // insertion, pour l'éviction des plus anciennes
}

func newPregenStore() *pregenStore {
	return &pregenStore{entries: map[string]*pregenEntry{}}
}

// register crée une entrée en cours (ou rend l'existante non terminée).
func (s *pregenStore) register(key string) *pregenEntry {
	s.mu.Lock()
	defer s.mu.Unlock()
	if e, ok := s.entries[key]; ok && !e.completed {
		return e
	}
	s.evictLocked()
	e := &pregenEntry{done: make(chan struct{}), createdAt: time.Now()}
	s.entries[key] = e
	s.order = append(s.order, key)
	return e
}

// fill remplit une entrée en cours (premier remplissage gagne) et la termine.
func (s *pregenStore) fill(key, reqKey string, resp provider.Response) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.entries[key]
	if !ok || e.completed {
		return
	}
	e.content = resp.Content
	e.reasoning = resp.Reasoning
	e.usage = cache.Usage{
		PromptTokens:     resp.Usage.PromptTokens,
		CompletionTokens: resp.Usage.CompletionTokens,
		TotalTokens:      resp.Usage.TotalTokens,
	}
	e.reqKey = reqKey
	s.finishLocked(e)
}

// abort termine une entrée en cours en échec (annulation, timeout, erreur).
// abortEntry ne termine que si l'entrée est toujours celle du fantôme
// appelant : un fantôme lent d'un job annulé ne doit pas empoisonner
// l'entrée recréée par le job suivant (même clé).
func (s *pregenStore) abort(key string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.abortEntryLocked(key, nil, err)
}

func (s *pregenStore) abortEntry(key string, pe *pregenEntry, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.abortEntryLocked(key, pe, err)
}

func (s *pregenStore) abortEntryLocked(key string, pe *pregenEntry, err error) {
	e, ok := s.entries[key]
	if !ok || e.completed {
		return
	}
	if pe != nil && e != pe {
		return
	}
	e.err = err
	s.finishLocked(e)
}

func (s *pregenStore) finishLocked(e *pregenEntry) {
	e.completed = true
	close(e.done)
}

// get rend l'entrée (en cours ou terminée), ou nil.
func (s *pregenStore) get(key string) *pregenEntry {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.entries[key]
}

// evictLocked supprime les entrées terminées périmées, puis les plus
// anciennes terminées au-delà de la borne. Jamais un fantôme en cours.
func (s *pregenStore) evictLocked() {
	now := time.Now()
	keep := s.order[:0]
	for _, k := range s.order {
		e, ok := s.entries[k]
		if !ok {
			continue
		}
		if e.completed && now.Sub(e.createdAt) > pregenEntryTTL {
			delete(s.entries, k)
			continue
		}
		keep = append(keep, k)
	}
	s.order = keep
	// Borne : fait de la place AVANT l'ajout (>=, pas >). Ne retire
	// jamais un fantôme en cours (un clic peut l'attendre) : on évince
	// le plus ancien *terminé*, même s'il n'est pas en tête.
	for len(s.order) >= pregenMaxEntries {
		idx := -1
		for i, k := range s.order {
			if e, ok := s.entries[k]; ok && e.completed {
				idx = i
				break
			}
		}
		if idx == -1 {
			break // que des fantômes en cours : on garde tout
		}
		delete(s.entries, s.order[idx])
		s.order = append(s.order[:idx], s.order[idx+1:]...)
	}
}

// pregenState : jobs et dépôt par moteur. Valeur zéro utilisable.
type pregenState struct {
	mu    sync.Mutex
	jobs  map[string]context.CancelFunc
	store *pregenStore
}

func (s *pregenState) getStore() *pregenStore {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.store == nil {
		s.store = newPregenStore()
	}
	return s.store
}

// cancelJobLocked annule le job en cours d'un utilisateur.
func (s *pregenState) cancelJobLocked(user string) {
	if s.jobs == nil {
		return
	}
	if cancel, ok := s.jobs[user]; ok {
		cancel()
		delete(s.jobs, user)
	}
}

// pregenInput : forme canonique d'un tour pour la clé d'entrée.
// Couvre tout ce qui détermine la requête provider, plus l'historique.
type pregenInput struct {
	User        string             `json:"user"`
	Family      string             `json:"family"`
	Mode        string             `json:"mode"`
	Text        string             `json:"text"`
	Web         bool               `json:"web"`
	MCP         bool               `json:"mcp"`
	Think       bool               `json:"think"`
	Effort      string             `json:"effort"`
	MaxTokens   int                `json:"max_tokens"`
	FocusCorpus string             `json:"focus_corpus"`
	AgentMode   bool               `json:"agent_mode"`
	History     []provider.Message `json:"history"`
}

// pregenInputKey rend la clé d'entrée d'un tour (déterministe).
// history = messages AVANT le message utilisateur du tour.
func pregenInputKey(in TurnInput, history []provider.Message) string {
	if history == nil {
		history = []provider.Message{}
	}
	raw, err := json.Marshal(pregenInput{
		User: in.User, Family: in.Family, Mode: in.Mode, Text: in.Text,
		Web: in.Web, MCP: in.MCP, Think: in.Think, Effort: in.Effort,
		MaxTokens: in.MaxTokens, FocusCorpus: in.FocusCorpus,
		AgentMode: in.AgentMode, History: history,
	})
	if err != nil {
		return ""
	}
	sum := sha256.Sum256(raw)
	return "pregen:" + hex.EncodeToString(sum[:])
}

// pregenReqKey rend la clé de la requête provider (même forme canonique
// que le cache exact, sans la porte d'éligibilité température 0).
func pregenReqKey(providerID string, req provider.Request) string {
	raw, err := json.Marshal(cacheKeyInput{
		Provider:        providerID,
		Model:           req.Model,
		Messages:        req.Messages,
		Temperature:     req.Temperature,
		MaxTokens:       req.MaxTokens,
		EnableReasoning: req.EnableReasoning,
		ReasoningEffort: req.ReasoningEffort,
	})
	if err != nil {
		return ""
	}
	sum := sha256.Sum256(raw)
	return "pregenreq:" + hex.EncodeToString(sum[:])
}

// pregenHistory rend l'historique d'un snapshot de messages : tout sauf
// le message utilisateur final du tour (ajouté par StartTurn).
func pregenHistory(msgs []provider.Message) []provider.Message {
	if n := len(msgs); n > 0 && msgs[n-1].Role == "user" {
		return msgs[:n-1]
	}
	return msgs
}

// PregenSuggestions lance la pré-génération des réponses aux questions
// suggérées affichées. history = messages courants de la conversation
// réelle (snapshot). Remplace tout job précédent de l'utilisateur.
// Ne bloque jamais ; les échecs sont silencieux (fail-open).
func (e *Engine) PregenSuggestions(user string, items []PregenSuggestion, base TurnInput, history []provider.Message) {
	if len(items) > pregenMaxItems {
		items = items[:pregenMaxItems]
	}
	// Copie défensive de l'historique (le snapshot peut être réutilisé).
	hist := make([]provider.Message, len(history))
	copy(hist, history)

	e.pregen.mu.Lock()
	e.pregen.cancelJobLocked(user)
	ctx, cancel := context.WithCancel(context.Background())
	if e.pregen.jobs == nil {
		e.pregen.jobs = map[string]context.CancelFunc{}
	}
	e.pregen.jobs[user] = cancel
	e.pregen.mu.Unlock()

	store := e.pregen.getStore()
	for _, it := range items {
		if it.Question == "" {
			continue
		}
		turnIn := base
		turnIn.User = user
		turnIn.Text = it.Question
		turnIn.FocusCorpus = it.Corpus
		turnIn.IsPregenShadow = true
		turnIn.PregenLookup = false
		turnIn.AgentMode = false
		turnIn.Plan = false
		turnIn.Worktree = false
		turnIn.Attachments = nil
		key := pregenInputKey(turnIn, hist)
		if key == "" {
			continue
		}
		turnIn.PregenKey = key
		pe := store.register(key)
		go e.pregenShadow(ctx, pe, turnIn, hist)
	}
}

// pregenShadow exécute un tour fantôme : le vrai Run sur une conversation
// jetable initialisée avec l'historique snapshoté. Le hook de capture
// (dans Run) remplit l'entrée en cas de succès ; sinon l'entrée est
// avortée (le clic suivra le chemin normal).
func (e *Engine) pregenShadow(ctx context.Context, pe *pregenEntry, in TurnInput, history []provider.Message) {
	store := e.pregen.getStore()
	defer func() {
		store.abortEntry(in.PregenKey, pe, context.Canceled)
	}()
	// Pas de pré-génération spéculative en mode agent (outils) : le
	// fantôme ne rejoue que le chemin simple. La sonde rejoue la
	// résolution avec AgentMode=true car le fantôme force false (sans
	// outils) alors que le vrai tour pourrait être agent.
	probe := in
	probe.AgentMode = true
	if res := e.resolve(ctx, probe); len(res.members) == 0 || res.agent {
		return
	}
	sctx, cancel := context.WithTimeout(ctx, pregenShadowTimeout)
	defer cancel()
	sh := NewConversation(newID(), e, nil)
	sh.Messages = make([]provider.Message, 0, len(history)+1)
	sh.Messages = append(sh.Messages, history...)
	sh.Messages = append(sh.Messages, provider.Message{Role: "user", Content: in.Text})
	e.Run(sctx, sh, 0, in)
}

// pregenCapture mémorise la réponse d'un tour fantôme (hook dans Run,
// chemin simple, après un appel provider réussi). La clé a été calculée
// une fois côté PregenSuggestions (in.PregenKey) : on ne la recalcule
// pas depuis msgs (prompts système déjà injectés).
func (e *Engine) pregenCapture(in TurnInput, providerID string, req provider.Request, resp provider.Response) {
	if !in.IsPregenShadow || in.PregenKey == "" {
		return
	}
	e.pregen.getStore().fill(in.PregenKey, pregenReqKey(providerID, req), resp)
}

// pregenServe cherche une réponse pré-générée pour le tour courant.
// Retourne true si la réponse a été rejouée (tour terminé). En cours =>
// attend le fantôme (borné) au lieu de relancer un appel.
// La clé est recalculée depuis le snapshot conversationnel (historique
// brut : les prompts système ne vivent que dans le msgs éphémère de
// Run) en retirant le message utilisateur du tour courant, déjà présent
// (StartTurn l'a enregistré avant Run).
func (e *Engine) pregenServe(ctx context.Context, c *Conversation, epoch int, in TurnInput, providerID string, req provider.Request) bool {
	if !in.PregenLookup {
		return false
	}
	key := pregenInputKey(in, pregenHistory(c.MessagesSnapshot()))
	if key == "" {
		return false
	}
	pe := e.pregen.getStore().get(key)
	if pe == nil {
		return false
	}
	select {
	case <-pe.done:
	case <-ctx.Done():
		return false
	case <-time.After(pregenWaitTimeout):
		return false
	}
	if pe.err != nil || pe.reqKey == "" {
		return false
	}
	// Stabilité : on ne sert que la requête à l'octet près.
	if pe.reqKey != pregenReqKey(providerID, req) {
		return false
	}
	e.replayCached(c, epoch, in, cache.Entry{
		Content:   pe.content,
		Reasoning: pe.reasoning,
		Usage:     pe.usage,
	})
	return true
}

// PregenCancel annule le job de pré-génération d'un utilisateur
// (nouveau message tapé avant la fin : l'historique change, les
// fantômes ne serviraient plus).
func (e *Engine) PregenCancel(user string) {
	e.pregen.mu.Lock()
	defer e.pregen.mu.Unlock()
	e.pregen.cancelJobLocked(user)
}
