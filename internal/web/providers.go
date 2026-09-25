package web

import (
	"net/http"
	"strings"

	"cetas-lite/internal/provider"
	"cetas-lite/internal/store"
	"cetas-lite/internal/vault"
)

// providerLabel retourne le libelle lisible d'un provider connu
// (voir providerLabels dans catalog.go).
type providerInfo struct {
	ID         string `json:"id"`
	Label      string `json:"label"`
	Configured bool   `json:"configured"`
	Local      bool   `json:"local"`
	URL        string `json:"url,omitempty"`
}

func isLocalEngine(id string) bool {
	for _, e := range provider.LocalEngines {
		if e == id {
			return true
		}
	}
	return false
}

// localEngineURL retourne l'URL effective d'un moteur local : valeur definie
// depuis l'interface (meta "local_url_<id>") en priorite, sinon "".
func localEngineURL(st *store.Store, id string) string {
	if raw, ok := st.GetMeta("local_url_" + id); ok && len(raw) > 0 {
		return string(raw)
	}
	return ""
}

// handleLocalURLPut met a jour l'URL d'un moteur local : application
// immediate au registre + persistance.
func (s *Server) handleLocalURLPut(w http.ResponseWriter, r *http.Request, id string) {
	var body struct {
		URL string `json:"url"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	u := strings.TrimSpace(body.URL)
	if u == "" {
		writeError(w, http.StatusBadRequest, "url vide")
		return
	}
	if !strings.HasPrefix(u, "http://") && !strings.HasPrefix(u, "https://") {
		writeError(w, http.StatusBadRequest, "url invalide (http:// ou https:// requis)")
		return
	}
	if len(u) > 200 {
		writeError(w, http.StatusBadRequest, "url trop longue")
		return
	}
	if err := s.st.PutMeta("local_url_"+id, []byte(u)); err != nil {
		writeError(w, http.StatusInternalServerError, "enregistrement impossible")
		return
	}
	// La cle du moteur (si rangee au coffre, ex. SamGen/llama.cpp distant)
	// doit survivre au changement d'URL : sans elle, la prise en compte a
	// chaud repartirait sans auth alors que le demarrage l'injecte.
	s.registry.Set(provider.NewOpenAICompat(provider.LocalEndpoint(id, u, s.localSecret(id)), s.httpClient))
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "id": id, "url": u})
}

// localSecret rend la cle dechiffree d'un moteur local, ou "" si absente ou
// illisible (fail-open : le moteur reste sans auth, comportement historique).
func (s *Server) localSecret(id string) string {
	ct, ok := s.st.GetSecret(id)
	if !ok {
		return ""
	}
	v, err := vault.Open(s.st)
	if err != nil {
		return ""
	}
	pt, err := v.Decrypt(ct, []byte(id))
	if err != nil {
		return ""
	}
	return string(pt)
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
			URL:        localEngineURL(s.st, id),
		})
	}
	writeJSON(w, http.StatusOK, map[string]any{"providers": out})
}

// handleProviderPut enregistre (chiffree) la cle API d'un provider cloud et
// met a jour le registre live pour une prise en compte immediate.
// Pour un moteur local (llamacpp, ollama, lmstudio), accepte {"url": ...} :
// l'URL est appliquee immediatement au registre et persistee (meta
// "local_url_<id>") pour survivre aux redemarrages.
func (s *Server) handleProviderPut(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	id := r.PathValue("id")
	if isLocalEngine(id) {
		s.handleLocalURLPut(w, r, id)
		return
	}
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
