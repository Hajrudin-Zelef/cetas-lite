package chat

import (
	"strings"
	"sync"
	"unicode/utf8"

	"cetas-lite/internal/attach"
	"cetas-lite/internal/docs"
)

// P0-A — plafond du contexte PJ : 8 Kio par fichier, 16 Kio au total.
const (
	attachPerFileBudget = 8 << 10
	attachContextBudget = 16 << 10
)

// P0-B — cache memoire LRU des textes extraits (cle = id de piece jointe).
// Borne volontairement basse : le sidecar disque reste la source durable.
const attachTextCacheMax = 64

type attachTextEntry struct {
	text string
	err  string
}

type attachTextCache struct {
	mu    sync.Mutex
	items map[string]attachTextEntry
	order []string
}

func (c *attachTextCache) get(id string) (attachTextEntry, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.items[id]
	return e, ok
}

func (c *attachTextCache) put(id string, e attachTextEntry) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.items == nil {
		c.items = map[string]attachTextEntry{}
	}
	if _, ok := c.items[id]; ok {
		c.touchLocked(id)
	}
	c.items[id] = e
	c.order = append(c.order, id)
	for len(c.order) > attachTextCacheMax {
		old := c.order[0]
		c.order = c.order[1:]
		delete(c.items, old)
	}
}

func (c *attachTextCache) touchLocked(id string) {
	for i, x := range c.order {
		if x == id {
			c.order = append(c.order[:i], c.order[i+1:]...)
			return
		}
	}
}

func (c *attachTextCache) remove(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, id)
	c.touchLocked(id)
}

func (e *Engine) attachText(id string) (attachTextEntry, bool) {
	return e.attachCache.get(id)
}

func (e *Engine) cacheAttachText(id string, v attachTextEntry) {
	e.attachCache.put(id, v)
}

// attachTextSidecarExt / attachThumbSidecarExt : suffixes des caches derives
// (texte extrait, vignette), alignes sur le store.
const (
	attachTextSidecarExt  = ".txt"
	attachThumbSidecarExt = ".thumb.jpg"
)

func (e *Engine) SetAttachments(s *attach.Store) {
	e.mu.Lock()
	e.attach = s
	e.mu.Unlock()
}

func (e *Engine) attachments() *attach.Store {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.attach
}

func (e *Engine) SaveAttachment(user, name string, data []byte) (attach.Attachment, error) {
	return e.attachments().Save(user, name, data)
}

func (e *Engine) OpenAttachment(user, id string) (attach.Attachment, []byte, error) {
	return e.attachments().Get(user, id)
}

func (e *Engine) DeleteAttachment(user, id string) error {
	e.attachCache.remove(id)
	return e.attachments().Delete(user, id)
}

// AttachmentText : texte extrait d'une piece jointe (cache memoire puis
// sidecar .txt puis extraction). Expose pour le modal de contenu (web).
// Renvoie (texte, "" ) ou ("", message d'erreur).
func (e *Engine) AttachmentText(user string, a attach.Attachment) (string, string) {
	st := e.attachments()
	if st == nil {
		return "", "stockage indisponible"
	}
	return e.extractAttachmentText(st, user, a)
}

// extractAttachmentText : texte d'une piece jointe non-image, memoise (memoire
// puis sidecar .txt puis extraction). Fail-open : une erreur est memorisee et
// presentee a l'utilisateur, sans bloquer le tour.
func (e *Engine) extractAttachmentText(st *attach.Store, user string, a attach.Attachment) (string, string) {
	if v, ok := e.attachText(a.ID); ok {
		return v.text, v.err
	}
	if raw, ok := st.ReadSidecar(user, a.ID, attachTextSidecarExt); ok {
		v := attachTextEntry{text: string(raw)}
		e.cacheAttachText(a.ID, v)
		return v.text, v.err
	}
	_, data, err := st.Get(user, a.ID)
	if err != nil {
		v := attachTextEntry{err: err.Error()}
		e.cacheAttachText(a.ID, v)
		return v.text, v.err
	}
	text, _, err := docs.Extract(a.Name, data)
	if err != nil {
		v := attachTextEntry{err: err.Error()}
		e.cacheAttachText(a.ID, v)
		return v.text, v.err
	}
	e.cacheAttachText(a.ID, attachTextEntry{text: text})
	_ = st.WriteSidecar(user, a.ID, attachTextSidecarExt, []byte(text))
	return text, ""
}

// AttachmentContext construit le message systeme des documents joints. Le
// contenu est borne (P0-A) : 8 Kio par fichier, 16 Kio au total, troncature
// signalee. L'appelant le passe a StartTurn (CtxMessage) pour qu'il soit
// persiste dans l'historique (P0-C).
func (e *Engine) AttachmentContext(user string, ids []string) string {
	st := e.attachments()
	if st == nil || len(ids) == 0 {
		return ""
	}
	var parts []string
	total := 0
	for _, id := range ids {
		a, _, err := st.Get(user, id)
		if err != nil {
			continue
		}
		if a.Kind == attach.KindImage {
			continue
		}
		if total >= attachContextBudget {
			parts = append(parts, "Attachment \""+a.Name+"\": … [budget PJ atteint]")
			break
		}
		text, xerr := e.extractAttachmentText(st, user, a)
		if xerr != "" {
			parts = append(parts, "Attachment \""+a.Name+"\": "+xerr)
			continue
		}
		text, perFile := capAttachmentText(text, attachPerFileBudget)
		remaining := attachContextBudget - total
		var totalTrunc bool
		if len(text) > remaining {
			text, totalTrunc = capAttachmentText(text, remaining)
		}
		total += len(text)
		body := text
		if perFile {
			body += "\n… [tronqué à 8 Kio]"
		}
		if totalTrunc {
			body += "\n… [budget PJ atteint]"
		}
		parts = append(parts, "Attachment \""+a.Name+"\":\n\n"+body)
	}
	if len(parts) == 0 {
		return ""
	}
	return "Documents attached by the user (use as context):\n\n" + strings.Join(parts, "\n\n---\n\n")
}

// AttachmentThumb renvoie la vignette JPEG d'une image jointe (P0-D) :
// sidecar .thumb.jpg s'il existe, sinon generation (max 256 px, q75) puis
// persistance. Fail-open : false => l'appelant sert l'image pleine.
func (e *Engine) AttachmentThumb(user string, a attach.Attachment, data []byte) ([]byte, bool) {
	st := e.attachments()
	if st == nil {
		return nil, false
	}
	if raw, ok := st.ReadSidecar(user, a.ID, attachThumbSidecarExt); ok {
		return raw, true
	}
	thumb, err := resizeJPEG(data, attachThumbMaxPx, attachThumbQuality)
	if err != nil {
		return nil, false
	}
	_ = st.WriteSidecar(user, a.ID, attachThumbSidecarExt, thumb)
	return thumb, true
}

// capAttachmentText tronque s a max octets sur une frontiere de caractere.
func capAttachmentText(s string, max int) (string, bool) {
	if max <= 0 {
		return "", len(s) > 0
	}
	if len(s) <= max {
		return s, false
	}
	end := max
	for end > 0 && !utf8.RuneStart(s[end]) {
		end--
	}
	return s[:end], true
}
