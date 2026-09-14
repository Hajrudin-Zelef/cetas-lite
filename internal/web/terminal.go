package web

import (
	"encoding/base64"
	"net/http"
	"os"
	"path/filepath"

	"cetas-lite/internal/terminal"
)

type terminalInfo struct {
	ID      string `json:"id"`
	Cwd     string `json:"cwd"`
	Created int64  `json:"created"`
}

func toTerminalInfo(s *terminal.Session) terminalInfo {
	return terminalInfo{ID: s.ID, Cwd: s.Cwd, Created: s.Created.Unix()}
}

// POST /api/terminal — cree une session shell. Corps optionnel : {"cwd": "..."}.
func (s *Server) handleTerminalCreate(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		Cwd string `json:"cwd"`
	}
	_ = decodeJSON(r, &body)
	cwd := body.Cwd
	if cwd == "" {
		cwd = filepath.Join(s.cfg.WorkspaceDir, claims.Username)
	}
	if err := os.MkdirAll(cwd, 0o700); err != nil {
		writeError(w, http.StatusBadRequest, "repertoire inaccessible")
		return
	}
	sess, err := s.terminals.Create(claims.Username, cwd)
	if err != nil {
		switch err {
		case terminal.ErrBadCwd:
			writeError(w, http.StatusForbidden, "repertoire de travail non autorise")
		case terminal.ErrTooMany:
			writeError(w, http.StatusConflict, "trop de terminaux ouverts (max 6)")
		default:
			writeError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	writeJSON(w, http.StatusOK, toTerminalInfo(sess))
}

// GET /api/terminal — liste les sessions de l'utilisateur.
func (s *Server) handleTerminalList(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	list := s.terminals.List(claims.Username)
	out := make([]terminalInfo, 0, len(list))
	for _, sess := range list {
		out = append(out, toTerminalInfo(sess))
	}
	writeJSON(w, http.StatusOK, map[string]any{"sessions": out})
}

// GET /api/terminal/{id}/stream — flux SSE de la sortie (base64).
func (s *Server) handleTerminalStream(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	sess, ok := s.terminals.Get(claims.Username, r.PathValue("id"))
	if !ok {
		writeError(w, http.StatusNotFound, "session introuvable")
		return
	}
	mu, stop, ok := sseStart(w)
	if !ok {
		return
	}
	defer stop()
	emit := sseEmitter(w, mu)

	ch, backlog, exited := sess.Subscribe()
	defer sess.Unsubscribe(ch)

	if len(backlog) > 0 {
		if !emit(map[string]any{"output": base64.StdEncoding.EncodeToString(backlog)}) {
			return
		}
	}
	for {
		select {
		case <-r.Context().Done():
			return
		case <-exited:
			// Vider ce qui reste avant d'annoncer la fin.
			for {
				select {
				case p, ok := <-ch:
					if !ok {
						emit(map[string]any{"exit": true})
						return
					}
					if len(p) > 0 && !emit(map[string]any{"output": base64.StdEncoding.EncodeToString(p)}) {
						return
					}
				default:
					emit(map[string]any{"exit": true})
					return
				}
			}
		case p, ok := <-ch:
			if !ok {
				emit(map[string]any{"exit": true})
				return
			}
			if len(p) > 0 && !emit(map[string]any{"output": base64.StdEncoding.EncodeToString(p)}) {
				return
			}
		}
	}
}

// POST /api/terminal/{id}/input — ecrit sur l'entree du shell.
func (s *Server) handleTerminalInput(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	sess, ok := s.terminals.Get(claims.Username, r.PathValue("id"))
	if !ok {
		writeError(w, http.StatusNotFound, "session introuvable")
		return
	}
	var body struct {
		Data string `json:"data"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	if err := sess.WriteInput([]byte(body.Data)); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

// POST /api/terminal/{id}/resize — redimensionne le PTY.
func (s *Server) handleTerminalResize(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	sess, ok := s.terminals.Get(claims.Username, r.PathValue("id"))
	if !ok {
		writeError(w, http.StatusNotFound, "session introuvable")
		return
	}
	var body struct {
		Cols int `json:"cols"`
		Rows int `json:"rows"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	_ = sess.Resize(body.Cols, body.Rows)
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

// DELETE /api/terminal/{id} — tue la session.
func (s *Server) handleTerminalDelete(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	if !s.terminals.Delete(claims.Username, r.PathValue("id")) {
		writeError(w, http.StatusNotFound, "session introuvable")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}
