package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"cetas-lite/internal/auth"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if v != nil {
		_ = json.NewEncoder(w).Encode(v)
	}
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func decodeJSON(r *http.Request, dst any) error {
	dec := json.NewDecoder(http.MaxBytesReader(nil, r.Body, 1<<20))
	return dec.Decode(dst)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status":  "ok",
		"version": s.version,
	})
}

func (s *Server) handleConfig(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"registration_open": s.registrationAllowed(),
		"version":           s.version,
	})
}

func (s *Server) registrationAllowed() bool {
	switch s.cfg.RegistrationMode {
	case "open":
		return true
	case "closed":
		return false
	default:
		return s.st.UserCount() == 0
	}
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) handleRegister(w http.ResponseWriter, r *http.Request) {
	if !s.registrationAllowed() {
		writeError(w, http.StatusForbidden, "inscription fermee (definir CETAS_LITE_REGISTRATION_OPEN=true)")
		return
	}
	var body credentials
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	err := s.auth.Register(body.Username, body.Password)
	switch {
	case err == nil:
		writeJSON(w, http.StatusCreated, map[string]any{"ok": true})
	case errors.Is(err, auth.ErrUserExists):
		writeError(w, http.StatusConflict, err.Error())
	case errors.Is(err, auth.ErrInvalidUsername), errors.Is(err, auth.ErrInvalidPassword):
		writeError(w, http.StatusBadRequest, err.Error())
	default:
		writeError(w, http.StatusInternalServerError, "erreur serveur")
	}
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	var body credentials
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	u, err := s.auth.Authenticate(body.Username, body.Password)
	if err != nil {
		writeError(w, http.StatusUnauthorized, auth.ErrInvalidCredentials.Error())
		return
	}
	token, err := s.auth.Token(u)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "erreur serveur")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"token": token,
		"user":  map[string]string{"username": u.Username, "role": u.Role},
	})
}

func (s *Server) handleMe(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"username": claims.Username,
		"role":     claims.Role,
	})
}
