package web

import (
	"encoding/json"
	"net/http"

	"cetas-lite/internal/alias"
)

type uiSettings struct {
	Theme        string `json:"theme"`
	Family       string `json:"family"`
	Mode         string `json:"mode"`
	AgentDefault bool   `json:"agent_default"`
	WebDefault   bool   `json:"web_default"`
}

func defaultUISettings() uiSettings {
	return uiSettings{Theme: "ocean", Family: "samagent-n4", Mode: "standard"}
}

func validTheme(t string) bool {
	switch t {
	case "ocean", "sombre", "clair":
		return true
	}
	return false
}

func (s *Server) storedSettings(user string) uiSettings {
	out := defaultUISettings()
	if raw, ok := s.st.GetSetting(user, "ui"); ok {
		if err := json.Unmarshal(raw, &out); err != nil {
			out = defaultUISettings()
		}
	}
	if !validTheme(out.Theme) {
		out.Theme = defaultUISettings().Theme
	}
	return out
}

func (s *Server) handleSettingsGet(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	writeJSON(w, http.StatusOK, s.storedSettings(claims.Username))
}

func (s *Server) handleSettingsPut(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		Theme        *string `json:"theme"`
		Family       *string `json:"family"`
		Mode         *string `json:"mode"`
		AgentDefault *bool   `json:"agent_default"`
		WebDefault   *bool   `json:"web_default"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	cur := s.storedSettings(claims.Username)
	if body.Theme != nil {
		if !validTheme(*body.Theme) {
			writeError(w, http.StatusBadRequest, "theme invalide")
			return
		}
		cur.Theme = *body.Theme
	}
	if body.Family != nil {
		cur.Family = *body.Family
	}
	if body.Mode != nil {
		cur.Mode = *body.Mode
	}
	if body.AgentDefault != nil {
		cur.AgentDefault = *body.AgentDefault
	}
	if body.WebDefault != nil {
		cur.WebDefault = *body.WebDefault
	}
	if cur.Family != "" && cur.Mode != "" {
		if _, ok := alias.Resolve(s.engine.Families(), cur.Family, cur.Mode); !ok {
			writeError(w, http.StatusBadRequest, "alias inconnu")
			return
		}
	}
	raw, err := json.Marshal(cur)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "erreur serveur")
		return
	}
	if err := s.st.PutSetting(claims.Username, "ui", raw); err != nil {
		writeError(w, http.StatusInternalServerError, "erreur serveur")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}
