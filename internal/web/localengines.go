package web

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"cetas-lite/internal/provider"
	"cetas-lite/internal/store"
	"cetas-lite/internal/vault"
)

// Logique SamGen (IA locale) : volontairement separee de la logique des
// providers cloud. Trois moteurs seulement (llama.cpp, ollama, lmstudio),
// config simple : URL + cle API facultative. La cle est chiffree au coffre
// sous l'ID du moteur et n'est jamais exposee (reponse : has_key).
type localEngineInfo struct {
	ID     string `json:"id"`
	Label  string `json:"label"`
	URL    string `json:"url,omitempty"`
	HasKey bool   `json:"has_key"`
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

// localModelsURL construit l'URL de sonde /v1/models a partir de l'URL
// configuree, sans dupliquer /v1 (ex. "https://hote/v1" -> ".../v1/models",
// pas ".../v1/v1/models"). Ne touche pas au comportement du chat.
func localModelsURL(u string) string {
	base := strings.TrimSuffix(strings.TrimRight(u, "/"), "/v1")
	return base + "/v1/models"
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

func localEngineLabel(id string) string {
	if l, ok := providerLabels[id]; ok {
		return l
	}
	return id
}

// handleLocalEnginesList retourne les trois moteurs SamGen et leur etat.
// GET /api/local/engines
func (s *Server) handleLocalEnginesList(w http.ResponseWriter, r *http.Request) {
	withSecrets := map[string]bool{}
	if names, err := s.st.ListSecrets(); err == nil {
		for _, n := range names {
			withSecrets[n] = true
		}
	}
	out := make([]localEngineInfo, 0, len(provider.LocalEngines))
	for _, id := range provider.LocalEngines {
		out = append(out, localEngineInfo{
			ID:     id,
			Label:  localEngineLabel(id),
			URL:    localEngineURL(s.st, id),
			HasKey: withSecrets[id],
		})
	}
	writeJSON(w, http.StatusOK, map[string]any{"engines": out})
}

// handleLocalEnginePut met a jour un moteur SamGen : URL et/ou cle API.
// Accepte {"url": ...} et/ou {"key": ...} ; au moins un est requis.
// PUT /api/local/engines/{id}
func (s *Server) handleLocalEnginePut(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if !isLocalEngine(id) {
		writeError(w, http.StatusNotFound, "moteur local inconnu")
		return
	}
	var body struct {
		URL string `json:"url"`
		Key string `json:"key"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	u := strings.TrimSpace(body.URL)
	key := strings.TrimSpace(body.Key)
	if u == "" && key == "" {
		writeError(w, http.StatusBadRequest, "url ou cle requise")
		return
	}
	if u != "" {
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
	} else {
		u = localEngineURL(s.st, id)
	}
	if key != "" {
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
	}
	// Pas d'endpoint vide dans le registre : sans URL, le moteur reste
	// simplement indisponible jusqu'a sa configuration.
	if u != "" {
		// La cle du moteur (si rangee au coffre, ex. SamGen/llama.cpp distant)
		// doit survivre au changement d'URL : sans elle, la prise en compte a
		// chaud repartirait sans auth alors que le demarrage l'injecte.
		s.registry.Set(provider.NewOpenAICompat(provider.LocalEndpoint(id, u, s.localSecret(id)), s.httpClient))
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "id": id, "url": u, "has_key": s.localSecret(id) != ""})
}

// handleLocalEngineDelete supprime la cle d'un moteur SamGen (l'URL est
// conservee) ; le registre repart sans auth (fail-open).
// DELETE /api/local/engines/{id}
func (s *Server) handleLocalEngineDelete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if !isLocalEngine(id) {
		writeError(w, http.StatusNotFound, "moteur local inconnu")
		return
	}
	if err := s.st.DeleteSecret(id); err != nil {
		writeError(w, http.StatusInternalServerError, "suppression impossible")
		return
	}
	if u := localEngineURL(s.st, id); u != "" {
		s.registry.Set(provider.NewOpenAICompat(provider.LocalEndpoint(id, u, ""), s.httpClient))
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "id": id})
}

// handleLocalEngineTest sonde un moteur SamGen, inspire de `ajean test` :
// GET {url}/v1/models avec la cle du coffre si presente.
// Repond toujours HTTP 200 ; l'echec est porte par {"ok": false, ...}.
// POST /api/local/engines/{id}/test
func (s *Server) handleLocalEngineTest(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if !isLocalEngine(id) {
		writeError(w, http.StatusNotFound, "moteur local inconnu")
		return
	}
	u := localEngineURL(s.st, id)
	if u == "" {
		writeJSON(w, http.StatusOK, map[string]any{"ok": false, "error": "url non configuree"})
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, localModelsURL(u), nil)
	if err != nil {
		writeJSON(w, http.StatusOK, map[string]any{"ok": false, "error": "requete impossible"})
		return
	}
	if k := s.localSecret(id); k != "" {
		req.Header.Set("Authorization", "Bearer "+k)
	}
	t0 := time.Now()
	resp, err := s.httpClient.Do(req)
	latMs := time.Since(t0).Milliseconds()
	if err != nil {
		writeJSON(w, http.StatusOK, map[string]any{"ok": false, "error": "moteur injoignable"})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		writeJSON(w, http.StatusOK, map[string]any{"ok": false, "error": "HTTP " + strconv.Itoa(resp.StatusCode)})
		return
	}
	var payload struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
		Models []struct {
			Name string `json:"name"`
		} `json:"models"`
	}
	n := 0
	if err := json.NewDecoder(resp.Body).Decode(&payload); err == nil {
		seen := map[string]bool{}
		for _, m := range payload.Data {
			if m.ID != "" && !seen[m.ID] {
				seen[m.ID] = true
				n++
			}
		}
		for _, m := range payload.Models {
			if m.Name != "" && !seen[m.Name] {
				seen[m.Name] = true
				n++
			}
		}
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "latency_ms": latMs, "models": n})
}
