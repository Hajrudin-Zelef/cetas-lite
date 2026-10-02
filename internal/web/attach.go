package web

import (
	"io"
	"mime"
	"net/http"
	"strings"

	"cetas-lite/internal/attach"
	"cetas-lite/internal/docs"
)

const attachMaxBytes = 20 << 20

func (s *Server) handleAttachUpload(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, attachMaxBytes+(1<<20))
	if err := r.ParseMultipartForm(attachMaxBytes); err != nil {
		writeError(w, http.StatusRequestEntityTooLarge, "fichier trop volumineux (max 20 Mo)")
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		writeError(w, http.StatusBadRequest, "fichier manquant")
		return
	}
	defer file.Close()
	name := header.Filename
	if !attach.IsImage(name) && !docs.IsSupported(name) {
		writeError(w, http.StatusBadRequest, "format non supporte: "+name)
		return
	}
	data, err := io.ReadAll(io.LimitReader(file, attachMaxBytes+1))
	if err != nil {
		writeError(w, http.StatusBadRequest, "lecture impossible")
		return
	}
	if int64(len(data)) > attachMaxBytes {
		writeError(w, http.StatusRequestEntityTooLarge, "fichier trop volumineux (max 20 Mo)")
		return
	}
	a, err := s.engine.SaveAttachment(claims.Username, name, data)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, a)
}

func (s *Server) handleAttachGet(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	id := r.PathValue("id")
	a, data, err := s.engine.OpenAttachment(claims.Username, id)
	if err != nil {
		writeError(w, http.StatusNotFound, "piece jointe introuvable")
		return
	}
	if r.URL.Query().Get("snippet") == "1" {
		s.serveAttachSnippet(w, claims.Username, a)
		return
	}
	if r.URL.Query().Get("text") == "1" {
		s.serveAttachText(w, claims.Username, a)
		return
	}
	// P0-D : vignette legerer, generee a la demande et mise en sidecar.
	if r.URL.Query().Get("thumb") == "1" && a.Kind == attach.KindImage {
		if thumb, ok := s.engine.AttachmentThumb(claims.Username, a, data); ok {
			w.Header().Set("Content-Type", "image/jpeg")
			w.Header().Set("Content-Disposition", `inline; filename="`+safeName(a.Name)+`.thumb.jpg"`)
			_, _ = w.Write(thumb)
			return
		}
	}
	ct := mime.TypeByExtension("." + a.Ext)
	if ct == "" {
		ct = "application/octet-stream"
	}
	w.Header().Set("Content-Type", ct)
	if r.URL.Query().Get("download") == "1" {
		w.Header().Set("Content-Disposition", `attachment; filename="`+safeName(a.Name)+`"`)
		_, _ = w.Write(data)
		return
	}
	if a.Kind == attach.KindImage {
		// Image pleine : inline, mais jamais rendue dans l'origine (le
		// blob est recupere par fetch et affiche via object URL).
		w.Header().Set("Content-Disposition", `inline; filename="`+safeName(a.Name)+`"`)
	} else {
		// Tout document non-image est un telechargement : jamais rendu
		// inline dans le navigateur (mitige HTML/SVG/PDF actif).
		w.Header().Set("Content-Disposition", `attachment; filename="`+safeName(a.Name)+`"`)
		w.Header().Set("Content-Security-Policy", "sandbox")
	}
	_, _ = w.Write(data)
}

// attachTextMaxBytes : borne du texte servi au modal (alignee sur docs.maxText).
const attachTextMaxBytes = 512 << 10

// serveAttachSnippet renvoie le debut du texte extrait d'une piece jointe
// non-image (apercu 1 ligne dans le chip). Les images utilisent ?thumb=1.
func (s *Server) serveAttachSnippet(w http.ResponseWriter, user string, a attach.Attachment) {
	if attach.IsImage(a.Name) {
		writeError(w, http.StatusBadRequest, "utiliser ?thumb=1 pour les images")
		return
	}
	text, xerr := s.engine.AttachmentText(user, a)
	if xerr != "" || strings.TrimSpace(text) == "" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte(truncateRunesWeb(text, 600)))
}

// serveAttachText renvoie le texte extrait complet (borne) pour le modal de
// contenu. Cache memoire + sidecar .txt cote moteur. Les images/PDF binaires
// n'ont pas de texte : le client affiche l'image ou l'iframe selon le kind.
func (s *Server) serveAttachText(w http.ResponseWriter, user string, a attach.Attachment) {
	if attach.IsImage(a.Name) {
		writeError(w, http.StatusBadRequest, "utiliser l'image pour les images")
		return
	}
	text, xerr := s.engine.AttachmentText(user, a)
	if xerr != "" || strings.TrimSpace(text) == "" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte(capAttachTextBytes(text, attachTextMaxBytes)))
}

// capAttachTextBytes borne s a max octets (UTF-8-safe), marqueur si tronque.
func capAttachTextBytes(s string, max int) string {
	if max <= 0 || len(s) <= max {
		return s
	}
	end := max
	for end > 0 && s[end]&0xC0 == 0x80 {
		end--
	}
	return s[:end] + "\n… [tronqué]"
}

// truncateRunesWeb coupe s a n runes (UTF-8-safe), avec ellipse si tronque.
func truncateRunesWeb(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n]) + "…"
}

func (s *Server) handleAttachDelete(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	if err := s.engine.DeleteAttachment(claims.Username, r.PathValue("id")); err != nil {
		writeError(w, http.StatusNotFound, "piece jointe introuvable")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}
