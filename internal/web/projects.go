package web

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"
	"unicode/utf8"

	"cetas-lite/internal/vfs"
	"cetas-lite/internal/workspace"
)

func projectJSON(p *workspace.Project) map[string]any {
	m := map[string]any{"id": p.ID, "name": p.Name, "mode": p.Mode}
	if p.Mode == workspace.ModeSFTP {
		m["host"] = p.Host
		m["port"] = p.Port
		m["user"] = p.User
		m["remote_path"] = p.RemotePath
	}
	return m
}

func (s *Server) wsManager() *workspace.Manager { return s.projects }

// GET /api/projects — liste les projets.
func (s *Server) handleProjectsList(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	list, err := s.projects.List()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "lecture impossible")
		return
	}
	out := make([]map[string]any, 0, len(list))
	for _, p := range list {
		out = append(out, projectJSON(p))
	}
	active, _ := s.activeProject(claimsFrom(r).Username)
	writeJSON(w, http.StatusOK, map[string]any{"projects": out, "active": active})
}

// POST /api/projects — crée un projet local ou SFTP.
func (s *Server) handleProjectsCreate(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		Name          string `json:"name"`
		Mode          string `json:"mode"`
		Host          string `json:"host"`
		Port          int    `json:"port"`
		User          string `json:"user"`
		RemotePath    string `json:"remote_path"`
		Password      string `json:"password"`
		PrivateKey    string `json:"private_key"`
		KeyPassphrase string `json:"key_passphrase"`
		HostKey       string `json:"host_key"` // clé hôte validée (base64), TOFU
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	mode := strings.TrimSpace(body.Mode)
	if mode == "" {
		mode = workspace.ModeLocal
	}
	var p *workspace.Project
	var err error
	switch mode {
	case workspace.ModeLocal:
		p, err = s.projects.CreateLocal(body.Name)
	case workspace.ModeSFTP:
		hk, _ := b64decode(body.HostKey)
		creds := workspace.SFTPCredentials{Password: body.Password, PrivateKey: body.PrivateKey, KeyPassphrase: body.KeyPassphrase}
		p, err = s.projects.CreateSFTP(body.Name, body.Host, body.Port, body.User, body.RemotePath, creds, hk)
	default:
		writeError(w, http.StatusBadRequest, "mode inconnu")
		return
	}
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, projectJSON(p))
}

// POST /api/projects/sftp-test — teste la connexion SFTP (TOFU).
func (s *Server) handleProjectSFTPTest(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		Host          string `json:"host"`
		Port          int    `json:"port"`
		User          string `json:"user"`
		RemotePath    string `json:"remote_path"`
		Password      string `json:"password"`
		PrivateKey    string `json:"private_key"`
		KeyPassphrase string `json:"key_passphrase"`
		HostKey       string `json:"host_key"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	hk, _ := b64decode(body.HostKey)
	creds := workspace.SFTPCredentials{Password: body.Password, PrivateKey: body.PrivateKey, KeyPassphrase: body.KeyPassphrase}
	fp, err := s.projects.TestSFTP(r.Context(), body.Host, body.Port, body.User, body.RemotePath, creds, hk)
	var uh *vfs.ErrUnknownHostKey
	if errors.As(err, &uh) {
		writeJSON(w, http.StatusOK, map[string]any{
			"ok": false, "unknown_host_key": true,
			"fingerprint": uh.Fingerprint, "key": b64encode(uh.Key),
		})
		return
	}
	_ = fp
	if err != nil {
		writeError(w, http.StatusBadRequest, "connexion impossible: "+err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

// DELETE /api/projects/{id} — supprime un projet.
func (s *Server) handleProjectDelete(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	id := r.PathValue("id")
	if err := s.projects.Delete(id); err != nil {
		writeError(w, http.StatusNotFound, "projet introuvable")
		return
	}
	if active, _ := s.activeProject(claims.Username); active == id {
		_ = s.setActiveProject(claims.Username, "")
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

// GET /api/projects/{id}/tree — arborescence (profondeur bornée).
func (s *Server) handleProjectTree(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	fsys, err := s.projects.OpenFS(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "projet introuvable")
		return
	}
	rel := strings.TrimSpace(r.URL.Query().Get("path"))
	depth := 3
	if d, err := strconv.Atoi(r.URL.Query().Get("depth")); err == nil && d >= 0 && d <= 6 {
		depth = d
	}
	node, err := vfs.Tree(r.Context(), fsys, rel, vfs.TreeOptions{MaxDepth: depth, MaxEntries: 2000})
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, node)
}

// GET /api/projects/{id}/file — lit un fichier (texte inline, PDF binaire).
func (s *Server) handleProjectFileGet(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	fsys, err := s.projects.OpenFS(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "projet introuvable")
		return
	}
	rel := strings.TrimSpace(r.URL.Query().Get("path"))
	clean, err := fsys.Resolve(rel)
	if err != nil || clean == "" {
		writeError(w, http.StatusBadRequest, "chemin invalide")
		return
	}
	fi, err := fsys.Stat(r.Context(), clean)
	if err != nil || fi.IsDir {
		writeError(w, http.StatusNotFound, "fichier introuvable")
		return
	}
	if fi.Size > 8<<20 {
		writeError(w, http.StatusBadRequest, "fichier trop volumineux (8 Mo max)")
		return
	}
	b, err := fsys.ReadFile(r.Context(), clean)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "lecture impossible")
		return
	}
	lower := strings.ToLower(clean)
	if strings.HasSuffix(lower, ".pdf") {
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "inline; filename="+strconv.Quote(path.Base(clean)))
		_, _ = w.Write(b)
		return
	}
	if !utf8.Valid(b) {
		writeJSON(w, http.StatusOK, map[string]any{"binary": true, "size": len(b)})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"content": string(b), "size": len(b)})
}

// DELETE /api/projects/{id}/file — supprime un fichier du projet.
func (s *Server) handleProjectFileDelete(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	fsys, err := s.projects.OpenFS(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "projet introuvable")
		return
	}
	rel := strings.TrimSpace(r.URL.Query().Get("path"))
	clean, err := fsys.Resolve(rel)
	if err != nil || clean == "" {
		writeError(w, http.StatusBadRequest, "chemin invalide")
		return
	}
	if err := fsys.Remove(r.Context(), clean); err != nil {
		writeError(w, http.StatusInternalServerError, "suppression impossible")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

// POST /api/projects/{id}/upload — upload multipart de fichiers.
// Chaque fichier porte son chemin relatif via le nom du champ
// ("files") et le filename ; les dossiers sont créés à la volée.
func (s *Server) handleProjectUpload(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	p, err := s.projects.Get(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "projet introuvable")
		return
	}
	if p.Mode != workspace.ModeLocal {
		writeError(w, http.StatusBadRequest, "l'upload n'est possible que pour un projet local")
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 200<<20)
	if err := r.ParseMultipartForm(200 << 20); err != nil {
		writeError(w, http.StatusBadRequest, "fichiers trop volumineux (200 Mo max)")
		return
	}
	fsys, err := s.projects.OpenFS(p.ID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "projet inaccessible")
		return
	}
	// Chemins relatifs : soit via le champ "paths" (JSON), soit via les
	// filenames tels quels.
	var rels []string
	if js := r.FormValue("paths"); js != "" {
		_ = json.Unmarshal([]byte(js), &rels)
	}
	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		writeError(w, http.StatusBadRequest, "aucun fichier")
		return
	}
	count := 0
	for i, fh := range files {
		rel := fh.Filename
		if i < len(rels) && strings.TrimSpace(rels[i]) != "" {
			rel = rels[i]
		}
		clean, err := vfs.CleanRel(rel)
		if err != nil || clean == "" {
			continue
		}
		f, err := fh.Open()
		if err != nil {
			continue
		}
		b, err := io.ReadAll(io.LimitReader(f, 50<<20))
		f.Close()
		if err != nil {
			continue
		}
		if err := fsys.WriteFile(r.Context(), clean, b, 0o644); err != nil {
			continue
		}
		count++
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "count": count})
}

// GET /api/projects/active — projet actif de l'utilisateur.
func (s *Server) handleProjectActiveGet(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	active, _ := s.activeProject(claims.Username)
	writeJSON(w, http.StatusOK, map[string]any{"active": active})
}

// PUT /api/projects/active — définit le projet actif ({id} ou vide).
func (s *Server) handleProjectActivePut(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		ID string `json:"id"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	id := strings.TrimSpace(body.ID)
	if id != "" {
		if _, err := s.projects.Get(id); err != nil {
			writeError(w, http.StatusNotFound, "projet introuvable")
			return
		}
	}
	if err := s.setActiveProject(claims.Username, id); err != nil {
		writeError(w, http.StatusInternalServerError, "enregistrement impossible")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "active": id})
}

func (s *Server) activeProject(user string) (string, error) {
	b, ok := s.st.GetMeta("active_project:" + safeUser(user))
	if !ok {
		return "", nil
	}
	return strings.TrimSpace(string(b)), nil
}

func (s *Server) setActiveProject(user, id string) error {
	return s.st.PutMeta("active_project:"+safeUser(user), []byte(id))
}

func safeUser(user string) string {
	user = strings.TrimSpace(user)
	user = strings.ReplaceAll(user, "/", "_")
	user = strings.ReplaceAll(user, "\\", "_")
	if user == "" {
		user = "default"
	}
	return user
}

func b64encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func b64decode(s string) ([]byte, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}
	return base64.StdEncoding.DecodeString(s)
}
