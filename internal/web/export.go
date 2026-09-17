package web

import (
	"net/http"
	"strings"
)

func exportFormat(r *http.Request) string {
	if strings.EqualFold(r.URL.Query().Get("format"), "json") {
		return "json"
	}
	return "md"
}

func (s *Server) handleChatExport(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	format := exportFormat(r)
	id, content, ok := s.engine.ExportActive(claims.Username, format)
	if !ok {
		writeError(w, http.StatusNotFound, "rien a exporter")
		return
	}
	writeDownload(w, format, "conversation-"+id, content)
}

func (s *Server) handleSessionExport(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	id := r.PathValue("id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "id requis")
		return
	}
	format := exportFormat(r)
	content, ok := s.engine.ExportSession(claims.Username, id, format)
	if !ok {
		writeError(w, http.StatusNotFound, "session introuvable")
		return
	}
	writeDownload(w, format, id, content)
}

func writeDownload(w http.ResponseWriter, format, name, content string) {
	ext := "md"
	ct := "text/markdown; charset=utf-8"
	if format == "json" {
		ext = "json"
		ct = "application/json; charset=utf-8"
	}
	w.Header().Set("Content-Type", ct)
	w.Header().Set("Content-Disposition", `attachment; filename="`+safeName(name)+`.`+ext+`"`)
	_, _ = w.Write([]byte(content))
}

func safeName(name string) string {
	var b strings.Builder
	for _, r := range name {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '.', r == '-', r == '_':
			b.WriteRune(r)
		default:
			b.WriteByte('-')
		}
	}
	if b.Len() == 0 {
		return "export"
	}
	return b.String()
}
