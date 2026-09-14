package web

import (
	"net/http"
	"strings"

	"cetas-lite/internal/provider"
	"cetas-lite/internal/vault"
)

// providerLabel retourne le libelle lisible d'un provider connu
// (voir providerLabels dans catalog.go).
type providerInfo struct {
	ID         string `json:"id"`
	Label      string `json:"label"`
	Configured bool   `json:"configured"`
	Local      bool   `json:"local"`
}

func providerLabel(id string) string {
	if l, ok := providerLabels[id]; ok {
		return l
	}
	return id
}

// handleProvidersList retourne les providers connus avec leur etat de
// configuration. Les cles ne sont jamais exposees, meme chiffrees.
func (s *Server) handleProvidersList(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	withSecrets := map[string]bool{}
	if names, err := s.st.ListSecrets(); err == nil {
		for _, n := range names {
			withSecrets[n] = true
		}
	}
	out := make([]providerInfo, 0, len(provider.CloudProviderIDs())+len(provider.LocalEngines))
	for _, id := range provider.CloudProviderIDs() {
		out = append(out, providerInfo{
			ID:         id,
			Label:      providerLabel(id),
			Configured: withSecrets[id],
			Local:      false,
		})
	}
	for _, id := range provider.LocalEngines {
		_, inRegistry := s.registry.Get(id)
		out = append(out, providerInfo{
			ID:         id,
			Label:      providerLabel(id),
			Configured: inRegistry,
			Local:      true,
		})
	}
	writeJSON(w, http.StatusOK, map[string]any{"providers": out})
}

// handleProviderPut enregistre (chiffree) la cle API d'un provider cloud et
// met a jour le registre live pour une prise en compte immediate.
func (s *Server) handleProviderPut(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	id := r.PathValue("id")
	if !provider.IsCloudProvider(id) {
		writeError(w, http.StatusNotFound, "provider inconnu")
		return
	}
	var body struct {
		Key string `json:"key"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	key := strings.TrimSpace(body.Key)
	if key == "" {
		writeError(w, http.StatusBadRequest, "cle vide")
		return
	}
	v, err := vault.Open(s.st)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "coffre indisponible")
		return
	}
	ct, err := v.Encrypt([]byte(key), []byte(id))
	if err != nil {
		writeError(w, http.StatusInternalServerError, "chiffrement impossible")
		return
	}
	if err := s.st.PutSecret(id, ct); err != nil {
		writeError(w, http.StatusInternalServerError, "enregistrement impossible")
		return
	}
	ep, ok := provider.CloudEndpoint(id, key)
	if !ok {
		writeError(w, http.StatusInternalServerError, "endpoint inconnu")
		return
	}
	s.registry.Set(provider.NewOpenAICompat(ep, s.httpClient))
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "id": id})
}

// handleProviderDelete supprime la cle d'un provider cloud et le retire du
// registre live.
func (s *Server) handleProviderDelete(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	id := r.PathValue("id")
	if !provider.IsCloudProvider(id) {
		writeError(w, http.StatusNotFound, "provider inconnu")
		return
	}
	if err := s.st.DeleteSecret(id); err != nil {
		writeError(w, http.StatusInternalServerError, "suppression impossible")
		return
	}
	s.registry.Delete(id)
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "id": id})
}
