package web

import (
	"encoding/json"
	"net/http"
	"strings"

	"cetas-lite/internal/search"
	"cetas-lite/internal/store"
	"cetas-lite/internal/vault"
)

// searchProviderMeta decrit un provider de recherche web pour l'UI.
type searchProviderMeta struct {
	ID       string
	Label    string
	KeyURL   string // page "obtenir une cle"
	KeyLabel string
	Keyless  bool // fonctionne sans cle (DuckDuckGo)
}

var searchProvidersMeta = []searchProviderMeta{
	{ID: "brave", Label: "Brave", KeyURL: "https://brave.com/search/api/", KeyLabel: "Clé API Brave"},
	{ID: "tavily", Label: "Tavily", KeyURL: "https://app.tavily.com", KeyLabel: "Clé API Tavily"},
	{ID: "jina", Label: "Jina", KeyURL: "https://jina.ai", KeyLabel: "Clé API Jina"},
	{ID: "exa", Label: "Exa", KeyURL: "https://dashboard.exa.ai", KeyLabel: "Clé API Exa"},
	{ID: "duckduckgo", Label: "DuckDuckGo", Keyless: true},
}

// searchSecretName : les secrets de recherche vivent sous "search_<id>".
func searchSecretName(id string) string { return "search_" + id }

// searchConfigMeta lit {enabled, mode} depuis le meta "search_config".
func searchConfigMeta(st *store.Store) (map[string]bool, string) {
	enabled := map[string]bool{}
	mode := "race"
	if raw, ok := st.GetMeta("search_config"); ok {
		var cfg struct {
			Enabled map[string]bool `json:"enabled"`
			Mode    string          `json:"mode"`
		}
		if json.Unmarshal(raw, &cfg) == nil {
			if cfg.Enabled != nil {
				enabled = cfg.Enabled
			}
			if cfg.Mode == "race" || cfg.Mode == "priority" {
				mode = cfg.Mode
			}
		}
	}
	return enabled, mode
}

// LoadSearchConfig construit la configuration du Searcher depuis le store :
// secrets "search_<id>" (repli anciens noms nus), interrupteurs et mode.
func LoadSearchConfig(st *store.Store, secrets map[string]string) search.Config {
	keys := map[string]string{}
	for _, m := range searchProvidersMeta {
		if m.Keyless {
			continue
		}
		if k := strings.TrimSpace(secrets[searchSecretName(m.ID)]); k != "" {
			keys[m.ID] = k
			continue
		}
		// Repli : anciens secrets stockes sous le nom nu du provider.
		if k := strings.TrimSpace(secrets[m.ID]); k != "" {
			keys[m.ID] = k
		}
	}
	enabled, mode := searchConfigMeta(st)
	return search.Config{Keys: keys, Enabled: enabled, Order: search.DefaultOrder, Mode: mode}
}

// refreshSearcher reconstruit le searcher du moteur a chaud.
// Applique a CETAS (chat) comme aux agents : le moteur est partage.
func (s *Server) refreshSearcher() {
	s.engine.SetSearcher(search.NewWithConfig(LoadSearchConfig(s.st, s.allSecrets()), s.httpClient))
}

func (s *Server) allSecrets() map[string]string {
	out := map[string]string{}
	names, err := s.st.ListSecrets()
	if err != nil {
		return out
	}
	v, err := vault.Open(s.st)
	if err != nil {
		return out
	}
	for _, name := range names {
		ct, ok := s.st.GetSecret(name)
		if !ok {
			continue
		}
		pt, err := v.Decrypt(ct, []byte(name))
		if err != nil {
			continue
		}
		out[name] = string(pt)
	}
	return out
}

// GET /api/search/settings — etat des providers de recherche web.
func (s *Server) handleSearchSettingsGet(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	secrets := s.allSecrets()
	enabled, mode := searchConfigMeta(s.st)
	providers := make([]map[string]any, 0, len(searchProvidersMeta))
	for _, m := range searchProvidersMeta {
		name := searchSecretName(m.ID)
		configured := strings.TrimSpace(secrets[name]) != "" || strings.TrimSpace(secrets[m.ID]) != ""
		en := true
		if v, ok := enabled[m.ID]; ok {
			en = v
		}
		providers = append(providers, map[string]any{
			"id":         m.ID,
			"label":      m.Label,
			"enabled":    en,
			"configured": configured || m.Keyless,
			"key_url":    m.KeyURL,
			"key_label":  m.KeyLabel,
			"keyless":    m.Keyless,
		})
	}
	writeJSON(w, http.StatusOK, map[string]any{"mode": mode, "providers": providers})
}

// PUT /api/search/settings — {mode?, providers: {id: {enabled?, key?}}}
// key present et vide => secret supprime ; key absent => inchange.
func (s *Server) handleSearchSettingsPut(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		Mode      *string `json:"mode"`
		Providers map[string]struct {
			Enabled *bool   `json:"enabled"`
			Key     *string `json:"key"`
		} `json:"providers"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	known := map[string]bool{}
	for _, m := range searchProvidersMeta {
		known[m.ID] = true
	}
	if body.Mode != nil {
		if *body.Mode != "race" && *body.Mode != "priority" {
			writeError(w, http.StatusBadRequest, "mode invalide")
			return
		}
	}
	for id := range body.Providers {
		if !known[id] {
			writeError(w, http.StatusBadRequest, "provider inconnu: "+id)
			return
		}
		if id == "duckduckgo" && body.Providers[id].Key != nil {
			writeError(w, http.StatusBadRequest, "duckduckgo ne prend pas de cle")
			return
		}
	}

	v, err := vault.Open(s.st)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "coffre indisponible")
		return
	}
	for id, p := range body.Providers {
		name := searchSecretName(id)
		if p.Key != nil {
			key := strings.TrimSpace(*p.Key)
			if key == "" {
				_ = s.st.DeleteSecret(name)
			} else {
				if len(key) > 4096 {
					writeError(w, http.StatusBadRequest, "cle trop longue: "+id)
					return
				}
				ct, err := v.Encrypt([]byte(key), []byte(name))
				if err != nil {
					writeError(w, http.StatusInternalServerError, "chiffrement impossible")
					return
				}
				if err := s.st.PutSecret(name, ct); err != nil {
					writeError(w, http.StatusInternalServerError, "enregistrement impossible")
					return
				}
			}
		}
	}
	enabled, mode := searchConfigMeta(s.st)
	if body.Mode != nil {
		mode = *body.Mode
	}
	for id, p := range body.Providers {
		if p.Enabled != nil {
			enabled[id] = *p.Enabled
		}
	}
	raw, _ := json.Marshal(map[string]any{"enabled": enabled, "mode": mode})
	if err := s.st.PutMeta("search_config", raw); err != nil {
		writeError(w, http.StatusInternalServerError, "enregistrement impossible")
		return
	}
	s.refreshSearcher()
	s.handleSearchSettingsGet(w, r)
}
