package chat

import (
	"strings"

	"cetas-lite/internal/attach"
	"cetas-lite/internal/docs"
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
	return e.attachments().Delete(user, id)
}

func (e *Engine) attachmentContext(user string, ids []string) string {
	st := e.attachments()
	if st == nil || len(ids) == 0 {
		return ""
	}
	var parts []string
	for _, id := range ids {
		a, data, err := st.Get(user, id)
		if err != nil {
			continue
		}
		if a.Kind == attach.KindImage {
			continue
		}
		text, _, err := docs.Extract(a.Name, data)
		if err != nil {
			parts = append(parts, "Piece jointe \""+a.Name+"\": "+err.Error())
			continue
		}
		parts = append(parts, "Piece jointe \""+a.Name+"\":\n\n"+text)
	}
	if len(parts) == 0 {
		return ""
	}
	return "Documents joints par l'utilisateur (a utiliser comme contexte):\n\n" + strings.Join(parts, "\n\n---\n\n")
}
