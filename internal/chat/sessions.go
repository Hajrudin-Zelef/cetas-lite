package chat

// Sessions : chaque conversation est une session isolée, persistée sous la
// clé "sess:<id>" dans le bucket conversations/<user> (bbolt).
//
// Fini la dualité "conversation active" / "archives" :
//   - une session = un enregistrement JSON (id, titre, dates, snapshot) ;
//   - ouvrir une session charge son snapshot dans la conversation courante ;
//   - supprimer une session efface son enregistrement, partout, sans retour.
//
// Isolation : une session n'écrit que sous sa propre clé (persist vérifie
// l'ID de la conversation). Deux sessions ne partagent jamais de données.
// La suppression pose une pierre tombale (tombstone) en mémoire : si un tour
// était encore en vol sur la session supprimée, sa persistance différée est
// ignorée — la session ne peut pas "revenir".

import (
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"cetas-lite/internal/provider"
)

const (
	// sessionKeyPrefix préfixe des clés de sessions dans le bucket
	// conversations/<user>. Les autres clés ("_current", anciennes clés)
	// sont ignorées par ListSessions.
	sessionKeyPrefix = "sess:"
	// currentKey stocke l'ID de la session courante de l'utilisateur.
	currentKey = "_current"
	// migrateFlagPrefix marque la migration archives -> sessions comme faite.
	migrateFlagPrefix = "sessions_migrated_v1:"
)

// SessionInfo décrit une session pour la liste (sidebar).
type SessionInfo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	Messages  int    `json:"messages"`
}

// sessionRecord est la "mémoire" d'une session : métadonnées + snapshot
// complet (messages + journal d'événements pour rejouer la vue).
type sessionRecord struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	CreatedAt int64    `json:"createdAt"`
	UpdatedAt int64    `json:"updatedAt"`
	Snapshot  snapshot `json:"snapshot"`
}

// tombstones empêche la résurrection d'une session supprimée dont un tour
// était encore en vol : sa persistance différée est ignorée.
// Gardé par sessMu (jamais pris sous e.mu pour éviter tout cycle).
var (
	sessMu     sync.Mutex
	tombstones = map[string]map[string]bool{}
)

func tombstoneLocked(user, id string) {
	m, ok := tombstones[user]
	if !ok {
		m = map[string]bool{}
		tombstones[user] = m
	}
	m[id] = true
}

func isTombstoned(user, id string) bool {
	sessMu.Lock()
	defer sessMu.Unlock()
	return tombstones[user][id]
}

// ListSessions retourne les sessions de l'utilisateur, triées par
// mise à jour décroissante. Effectue la migration unique depuis l'ancien
// format (archives + conversation "active").
func (e *Engine) ListSessions(user string) []SessionInfo {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.ensureMigratedLocked(user)
	if e.st == nil {
		return nil
	}
	ids, _ := e.st.ListConversations(user)
	out := make([]SessionInfo, 0, len(ids))
	for _, key := range ids {
		if !strings.HasPrefix(key, sessionKeyPrefix) {
			continue
		}
		id := strings.TrimPrefix(key, sessionKeyPrefix)
		if isTombstoned(user, id) {
			continue
		}
		raw, ok := e.st.GetConversation(user, key)
		if !ok {
			continue
		}
		var rec sessionRecord
		if json.Unmarshal(raw, &rec) != nil || rec.ID == "" {
			continue
		}
		out = append(out, SessionInfo{
			ID:        rec.ID,
			Title:     rec.Title,
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt,
			Messages:  countChatMessages(rec.Snapshot.Messages),
		})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].UpdatedAt != out[j].UpdatedAt {
			return out[i].UpdatedAt > out[j].UpdatedAt
		}
		return out[i].ID > out[j].ID
	})
	return out
}

// CreateSession démarre une nouvelle session vide et en fait la session
// courante. L'enregistrement est écrit aussitôt : la session apparaît
// dans la liste même avant le premier message (parité avec les grands
// providers cloud).
func (e *Engine) CreateSession(user string) string {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.ensureMigratedLocked(user)
	c := e.conversationLocked(user)
	id := newID()
	c.restore(snapshot{ID: id})
	e.setCurrentLocked(user, id)
	if e.st != nil {
		if data, err := json.Marshal(e.recordFor(c)); err == nil {
			_ = e.st.PutConversation(user, sessionKeyPrefix+id, data)
		}
	}
	return id
}

// OpenSession charge la session id comme session courante. Retourne false
// si la session n'existe pas (ou a été supprimée).
func (e *Engine) OpenSession(user, id string) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.ensureMigratedLocked(user)
	if e.st == nil || isTombstoned(user, id) {
		return false
	}
	rec, ok := e.getSessionLocked(user, id)
	if !ok {
		return false
	}
	c := e.conversationLocked(user)
	if c.ID != id {
		// restore() stoppe un éventuel tour en cours, recharge le contenu
		// et notifie les abonnés SSE (epoch -> {reset:true} -> rejouent).
		c.restore(rec.Snapshot)
	}
	e.setCurrentLocked(user, id)
	return true
}

// DeleteSession supprime définitivement une session : enregistrement effacé,
// pierre tombale posée (aucune persistance différée ne peut la ressusciter).
// Si c'était la session courante, une session vierge la remplace aussitôt.
// Retourne l'ID de la session courante après suppression.
func (e *Engine) DeleteSession(user, id string) (curID string, ok bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.ensureMigratedLocked(user)
	if e.st == nil {
		return "", false
	}
	if _, found := e.getSessionLocked(user, id); !found {
		return "", false
	}
	sessMu.Lock()
	tombstoneLocked(user, id)
	sessMu.Unlock()
	_ = e.st.DeleteConversation(user, sessionKeyPrefix+id)
	c := e.conversationLocked(user)
	if c.ID == id {
		nid := newID()
		c.restore(snapshot{ID: nid})
		e.setCurrentLocked(user, nid)
		// Pas d'enregistrement : la remplaçante apparaîtra au 1er message.
		return nid, true
	}
	return e.currentIDLocked(user), true
}

// CurrentSessionID retourne l'ID de la session courante de l'utilisateur.
func (e *Engine) CurrentSessionID(user string) string {
	e.mu.Lock()
	defer e.mu.Unlock()
	if c, ok := e.convs[user]; ok {
		return c.ID
	}
	e.ensureMigratedLocked(user)
	return e.currentIDLocked(user)
}

// saveSession est le callback de persistance des conversations : chaque
// session écrit sous SA propre clé. Les sessions vides ne sont jamais
// persistées implicitement ; les sessions supprimées (tombstone) sont
// ignorées — c'est ce qui rend la suppression définitive.
func (e *Engine) saveSession(user string, c *Conversation) {
	if e.st == nil || c.isEmpty() || isTombstoned(user, c.ID) {
		return
	}
	data, err := json.Marshal(e.recordFor(c))
	if err != nil {
		return
	}
	_ = e.st.PutConversation(user, sessionKeyPrefix+c.ID, data)
}

// recordFor construit l'enregistrement JSON d'une session ("mémoire" de la
// session) à partir de l'état courant de la conversation.
func (e *Engine) recordFor(c *Conversation) sessionRecord {
	snap := c.marshal()
	title := sessionTitle(snap)
	updated := time.Now().UnixMilli()
	if len(snap.Log) > 0 {
		updated = snap.Log[len(snap.Log)-1].TS
	}
	return sessionRecord{
		ID:        snap.ID,
		Title:     title,
		CreatedAt: sessionCreatedAt(snap.ID),
		UpdatedAt: updated,
		Snapshot:  snap,
	}
}

func (e *Engine) getSessionLocked(user, id string) (sessionRecord, bool) {
	var rec sessionRecord
	raw, ok := e.st.GetConversation(user, sessionKeyPrefix+id)
	if !ok {
		return rec, false
	}
	if json.Unmarshal(raw, &rec) != nil || rec.ID == "" {
		return rec, false
	}
	return rec, true
}

func (e *Engine) currentIDLocked(user string) string {
	if e.st == nil {
		return ""
	}
	raw, ok := e.st.GetConversation(user, currentKey)
	if !ok {
		return ""
	}
	return string(raw)
}

func (e *Engine) setCurrentLocked(user, id string) {
	if e.st == nil {
		return
	}
	_ = e.st.PutConversation(user, currentKey, []byte(id))
}

// sessionCreatedAt déduit la date de création de l'ID : les IDs générés par
// newID() sont des UnixNano, les IDs migrés portent un préfixe horodaté
// "20060102_150405".
func sessionCreatedAt(id string) int64 {
	if len(id) >= 15 {
		if t, err := time.ParseInLocation("20060102_150405", id[:15], time.Local); err == nil {
			return t.UnixMilli()
		}
	}
	if nanos, err := strconv.ParseInt(id, 10, 64); err == nil {
		return nanos / 1e6
	}
	return time.Now().UnixMilli()
}

func sessionTitle(snap snapshot) string {
	for _, m := range snap.Messages {
		if m.Role != "user" {
			continue
		}
		if txt, ok := m.Content.(string); ok && strings.TrimSpace(txt) != "" {
			return condense(txt, 80)
		}
	}
	return "(sans titre)"
}

func countChatMessages(msgs []provider.Message) int {
	n := 0
	for _, m := range msgs {
		if m.Role == "user" || m.Role == "assistant" {
			n++
		}
	}
	return n
}

// ensureMigratedLocked importe une fois pour toutes l'ancien format
// (bucket archives/* + clé "active") vers les sessions. Idempotent,
// protégé par un flag en meta.
func (e *Engine) ensureMigratedLocked(user string) {
	if e.st == nil {
		return
	}
	flag := migrateFlagPrefix + user
	if _, ok := e.st.GetMeta(flag); ok {
		return
	}
	// 1. Archives -> sessions (on conserve l'ID d'archive : favoris et
	// catégories du client restent valides).
	if ids, _ := e.st.ListArchives(user); ids != nil {
		for _, aid := range ids {
			raw, ok := e.st.GetArchive(user, aid)
			if !ok {
				continue
			}
			var s snapshot
			if json.Unmarshal(raw, &s) != nil || s.ID == "" {
				continue
			}
			s.ID = aid
			rec := sessionRecord{
				ID:        aid,
				Title:     sessionTitle(s),
				CreatedAt: sessionCreatedAt(aid),
				UpdatedAt: sessionUpdatedAt(s),
				Snapshot:  s,
			}
			if data, err := json.Marshal(rec); err == nil {
				_ = e.st.PutConversation(user, sessionKeyPrefix+aid, data)
			}
		}
	}
	// 2. Ancienne conversation "active" -> session (conserve son ID).
	if raw, ok := e.st.GetConversation(user, "active"); ok {
		var s snapshot
		if json.Unmarshal(raw, &s) == nil && s.ID != "" && len(s.Messages) > 0 {
			rec := sessionRecord{
				ID:        s.ID,
				Title:     sessionTitle(s),
				CreatedAt: sessionCreatedAt(s.ID),
				UpdatedAt: sessionUpdatedAt(s),
				Snapshot:  s,
			}
			if data, err := json.Marshal(rec); err == nil {
				_ = e.st.PutConversation(user, sessionKeyPrefix+s.ID, data)
				e.setCurrentLocked(user, s.ID)
			}
		}
		_ = e.st.DeleteConversation(user, "active")
	}
	_ = e.st.PutMeta(flag, []byte("1"))
}

func sessionUpdatedAt(s snapshot) int64 {
	if len(s.Log) > 0 {
		return s.Log[len(s.Log)-1].TS
	}
	return sessionCreatedAt(s.ID)
}
