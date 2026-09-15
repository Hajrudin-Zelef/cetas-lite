package web

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"cetas-lite/internal/vault"
)

// githubSecretKey est la clé du secret chiffré contenant le token GitHub.
const githubSecretKey = "github_token"

// GET /api/connectors — état des connecteurs configurés.
func (s *Server) handleConnectors(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	login := ""
	if b, ok := s.st.GetMeta("github_login"); ok {
		login = strings.TrimSpace(string(b))
	}
	_, connected := s.st.GetSecret(githubSecretKey)
	writeJSON(w, http.StatusOK, map[string]any{
		"github": map[string]any{"connected": connected && login != "", "login": login},
	})
}

// PUT /api/connectors/github — connecte un compte GitHub via token.
// Le token est validé contre l'API GitHub puis chiffré dans le coffre.
// L'agent l'utilise ensuite pour commit/diff/push via git, sans jamais
// modifier les fichiers de configuration du dépôt.
func (s *Server) handleGitHubPut(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		Token string `json:"token"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	token := strings.TrimSpace(body.Token)
	if token == "" {
		writeError(w, http.StatusBadRequest, "token vide")
		return
	}
	login, err := githubValidateToken(r.Context(), s.httpClient, token)
	if err != nil {
		writeError(w, http.StatusBadRequest, "token invalide: "+err.Error())
		return
	}
	v, err := vault.Open(s.st)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "coffre indisponible")
		return
	}
	ct, err := v.Encrypt([]byte(token), []byte(githubSecretKey))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "chiffrement impossible")
		return
	}
	if err := s.st.PutSecret(githubSecretKey, ct); err != nil {
		writeError(w, http.StatusInternalServerError, "enregistrement impossible")
		return
	}
	_ = s.st.PutMeta("github_login", []byte(login))
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "login": login})
}

// DELETE /api/connectors/github — déconnecte le compte GitHub.
func (s *Server) handleGitHubDelete(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	_ = s.st.DeleteSecret(githubSecretKey)
	_ = s.st.PutMeta("github_login", []byte(""))
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}

// githubValidateToken vérifie le token et retourne le login GitHub.
func githubValidateToken(ctx context.Context, client *http.Client, token string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.github.com/user", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github+json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", &httpError{status: resp.StatusCode}
	}
	var u struct {
		Login string `json:"login"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return "", err
	}
	if strings.TrimSpace(u.Login) == "" {
		return "", &httpError{status: resp.StatusCode}
	}
	return u.Login, nil
}

type httpError struct{ status int }

func (e *httpError) Error() string {
	if e.status == 401 || e.status == 403 {
		return "non autorise (verifiez le token)"
	}
	return "GitHub a repondu " + strings.TrimSpace(http.StatusText(e.status))
}
