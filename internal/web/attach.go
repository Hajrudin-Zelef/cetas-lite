package web

import (
	"io"
	"mime"
	"net/http"

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
	if attach.IsImage(name) {
		writeError(w, http.StatusBadRequest, "les images seront prises en charge au lot suivant")
		return
	}
	if !docs.IsSupported(name) {
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
	ct := mime.TypeByExtension("." + a.Ext)
	if ct == "" {
		ct = "application/octet-stream"
	}
	w.Header().Set("Content-Type", ct)
	w.Header().Set("Content-Disposition", `inline; filename="`+safeName(a.Name)+`"`)
	_, _ = w.Write(data)
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
