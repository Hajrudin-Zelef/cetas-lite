package chat

import (
	"os"
	"strings"
	"unicode/utf8"

	"cetas-lite/internal/provider"
)

// MAREX.md : fichier de contexte long renseigne par l'utilisateur.
// CETAS (chat) et les Agents le lisent a chaque demarrage de session.
// Vide ou absent : ignore silencieusement. Rempli : le contenu est injecte
// en message systeme et "MAREX.md chargé" s'affiche discrètement dans le chat.

const marexMaxBytes = 64 << 10 // 64 Kio

// Nombre max de sessions mises en cache ; au-dela, le cache est purge
// (le contenu est simplement relu au prochain demarrage de session).
const marexMaxSessions = 200

type marexEntry struct {
	content string // "" si vide/absent
	shown   bool   // notice "MAREX.md chargé" deja emise pour cette session
}

// SetMarexPath definit le chemin du fichier MAREX.md de l'utilisateur.
func (e *Engine) SetMarexPath(path string) {
	e.marexMu.Lock()
	defer e.marexMu.Unlock()
	if e.marexPath != path {
		e.marexPath = path
		e.marexSessions = nil // changement de fichier : relecture aux prochaines sessions
	}
}

// MarexPath retourne le chemin configure du fichier MAREX.md.
func (e *Engine) MarexPath() string {
	e.marexMu.Lock()
	defer e.marexMu.Unlock()
	return e.marexPath
}

// ReadMarex lit le contenu brut du fichier MAREX.md ("", nil si absent).
func (e *Engine) ReadMarex() (string, error) {
	e.marexMu.Lock()
	path := e.marexPath
	e.marexMu.Unlock()
	if path == "" {
		return "", nil
	}
	b, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	return string(b), nil
}

// WriteMarex enregistre le contenu du fichier MAREX.md (borne a 1 Mio).
func (e *Engine) WriteMarex(content string) error {
	e.marexMu.Lock()
	path := e.marexPath
	e.marexMu.Unlock()
	if path == "" {
		return os.ErrNotExist
	}
	if len(content) > 1<<20 {
		content = content[:1<<20]
	}
	if !utf8.ValidString(content) {
		content = strings.ToValidUTF8(content, "")
	}
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		return err
	}
	// Le fichier a change : les prochaines sessions reliront le contenu.
	e.marexMu.Lock()
	e.marexSessions = nil
	e.marexMu.Unlock()
	return nil
}

// marexForSession retourne le message systeme MAREX.md pour la conversation
// convID. Le fichier est lu une seule fois, au demarrage de la session ;
// "notice" n'est vrai qu'au premier tour non vide (l'UI affiche alors le
// message discret "MAREX.md chargé"). Fichier vide ou absent : ignore
// silencieusement (ok=false).
func (e *Engine) marexForSession(convID string) (msg provider.Message, ok bool, notice bool) {
	e.marexMu.Lock()
	defer e.marexMu.Unlock()
	if e.marexSessions == nil {
		e.marexSessions = map[string]*marexEntry{}
	}
	if len(e.marexSessions) >= marexMaxSessions {
		e.marexSessions = map[string]*marexEntry{}
	}
	if ent, found := e.marexSessions[convID]; found {
		if ent.content == "" {
			return provider.Message{}, false, false
		}
		notice = !ent.shown
		ent.shown = true
		return provider.Message{Role: "system", Content: "[MAREX.md]\n\n" + ent.content}, true, notice
	}
	var content string
	if e.marexPath != "" {
		if b, err := os.ReadFile(e.marexPath); err == nil {
			if len(b) > marexMaxBytes {
				b = b[:marexMaxBytes]
			}
			content = strings.TrimSpace(string(b))
		}
	}
	ent := &marexEntry{content: content}
	e.marexSessions[convID] = ent
	if content == "" {
		return provider.Message{}, false, false
	}
	ent.shown = true
	return provider.Message{Role: "system", Content: "[MAREX.md]\n\n" + content}, true, true
}
