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

func (s *Server) handleSettingsGet(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	out := defaultUISettings()
	if raw, ok := s.st.GetSetting(claims.Username, "ui"); ok {
		if err := json.Unmarshal(raw, &out); err != nil {
			out = defaultUISettings()
		}
	}
	if !validTheme(out.Theme) {
		out.Theme = defaultUISettings().Theme
	}
	writeJSON(w, http.StatusOK, out)
}

func (s *Server) handleSettingsPut(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	body := defaultUISettings()
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	if !validTheme(body.Theme) {
		writeError(w, http.StatusBadRequest, "theme invalide")
		return
	}
	if body.Family != "" && body.Mode != "" {
		if _, ok := alias.Resolve(s.engine.Families(), body.Family, body.Mode); !ok {
			writeError(w, http.StatusBadRequest, "alias inconnu")
			return
		}
	}
	raw, err := json.Marshal(body)
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
