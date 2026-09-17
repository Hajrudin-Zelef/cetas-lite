package web

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"cetas-lite/internal/provider"
)

// DeepThink Global : traduction a la demande du raisonnement affiche dans
// le panneau lateral. Un bouton dans l'en-tete REASONING traduit le texte
// via un petit modele dedie, dans la langue choisie par l'utilisateur
// (module "DeepThink Global" des parametres). La traduction est effectuee
// cote serveur : les cles API ne sont jamais exposees au navigateur.

// Langues cibles proposees (code -> nom anglais utilise dans le prompt).
var deepThinkLangs = map[string]string{
	"es": "Spanish",
	"fr": "French",
	"it": "Italian",
	"de": "German",
	"en": "English",
}

// deepThinkModel : un modele de traduction (petit, bon marche, stable).
type deepThinkModel struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
	Label    string `json:"label"`
}

var deepThinkModels = []deepThinkModel{
	{Provider: "openrouter", Model: "openrouter/free", Label: "openrouter/free (OpenRouter)"},
	// Sur l'API DeepSeek, le modele de chat courant (V3.2) s'appelle deepseek-chat.
	{Provider: "deepseek", Model: "deepseek-chat", Label: "DeepSeek V3.2 (DeepSeek)"},
}

// DeepThinkSettings : reglages persistants du module.
type DeepThinkSettings struct {
	Lang     string `json:"lang"`
	Provider string `json:"provider"`
	Model    string `json:"model"`
}

func defaultDeepThinkSettings() DeepThinkSettings {
	return DeepThinkSettings{Lang: "fr", Provider: "openrouter", Model: "openrouter/free"}
}

func validDeepThinkModel(providerID, model string) bool {
	for _, m := range deepThinkModels {
		if m.Provider == providerID && m.Model == model {
			return true
		}
	}
	return false
}

// sanitizeDeepThinkSettings retombe sur les defauts pour tout champ invalide
// (ex. donnees corrompues ou modele retire du catalogue).
func sanitizeDeepThinkSettings(st DeepThinkSettings) DeepThinkSettings {
	def := defaultDeepThinkSettings()
	if _, ok := deepThinkLangs[st.Lang]; !ok {
		st.Lang = def.Lang
	}
	if !validDeepThinkModel(st.Provider, st.Model) {
		st.Provider, st.Model = def.Provider, def.Model
	}
	return st
}

func (s *Server) loadDeepThinkSettings() DeepThinkSettings {
	if raw, ok := s.st.GetMeta("deepthink"); ok {
		var st DeepThinkSettings
		if err := json.Unmarshal(raw, &st); err == nil {
			return sanitizeDeepThinkSettings(st)
		}
	}
	return defaultDeepThinkSettings()
}

// GET /api/deepthink -> reglages + choix disponibles.
func (s *Server) handleDeepThinkGet(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"settings": s.loadDeepThinkSettings(),
		"langs":    []string{"es", "fr", "it", "de", "en"},
		"models":   deepThinkModels,
	})
}

// PUT /api/deepthink {"lang","provider","model"} -> valide et persiste.
func (s *Server) handleDeepThinkPut(w http.ResponseWriter, r *http.Request) {
	var st DeepThinkSettings
	if err := decodeJSON(r, &st); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	if _, ok := deepThinkLangs[st.Lang]; !ok {
		writeError(w, http.StatusBadRequest, "langue inconnue : "+st.Lang)
		return
	}
	if !validDeepThinkModel(st.Provider, st.Model) {
		writeError(w, http.StatusBadRequest, "modele de traduction inconnu")
		return
	}
	if raw, err := json.Marshal(st); err == nil {
		_ = s.st.PutMeta("deepthink", raw)
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "settings": st})
}

// POST /api/deepthink/translate {"text": "..."} -> {"translation": "..."}.
// Utilise les reglages DeepThink (langue + modele). Les erreurs provider
// sont renvoyees en 502 avec un message lisible pour l'UI.
func (s *Server) handleDeepThinkTranslate(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Text string `json:"text"`
	}
	if err := decodeJSON(r, &in); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	text := strings.TrimSpace(in.Text)
	if text == "" {
		writeError(w, http.StatusBadRequest, "texte vide")
		return
	}
	// Garde-fou : un raisonnement fait rarement plus de quelques milliers
	// de caracteres ; au-dela on tronque pour rester rapide et stable.
	const maxInputChars = 12000
	if rs := []rune(text); len(rs) > maxInputChars {
		text = string(rs[:maxInputChars])
	}
	st := s.loadDeepThinkSettings()
	p, ok := s.registry.Get(st.Provider)
	if !ok {
		writeError(w, http.StatusBadRequest,
			fmt.Sprintf("provider %q non configuré : ajoutez votre clé API dans les paramètres", st.Provider))
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 90*time.Second)
	defer cancel()
	out, err := translateText(ctx, p, st, text)
	if err != nil {
		writeError(w, http.StatusBadGateway, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"translation": out, "lang": st.Lang})
}

// Cache des traductions : le meme raisonnement est souvent re-affiche
// (bouton "Raisonnement", instantanes) — la 2e fois c'est instantane.
var (
	dtCacheMu    sync.Mutex
	dtCache      = map[string]string{}
	dtCacheOrder []string
)

const dtCacheMax = 100

func dtCacheKey(st DeepThinkSettings, text string) string {
	h := sha256.New()
	h.Write([]byte(st.Provider + "\x00" + st.Model + "\x00" + st.Lang + "\x00" + text))
	return hex.EncodeToString(h.Sum(nil))
}

func dtCacheGet(key string) (string, bool) {
	dtCacheMu.Lock()
	defer dtCacheMu.Unlock()
	out, ok := dtCache[key]
	return out, ok
}

func dtCacheSet(key, val string) {
	dtCacheMu.Lock()
	defer dtCacheMu.Unlock()
	if _, ok := dtCache[key]; !ok {
		dtCacheOrder = append(dtCacheOrder, key)
		for len(dtCacheOrder) > dtCacheMax {
			old := dtCacheOrder[0]
			dtCacheOrder = dtCacheOrder[1:]
			delete(dtCache, old)
		}
	}
	dtCache[key] = val
}

// translateText appelle le modele de traduction. Le prompt est volontairement
// minimal (temperature basse) : la tache est deterministe. Factorisee pour
// les tests (provider mockable).
func translateText(ctx context.Context, p provider.Provider, st DeepThinkSettings, text string) (string, error) {
	key := dtCacheKey(st, text)
	if out, ok := dtCacheGet(key); ok {
		return out, nil
	}
	langName, ok := deepThinkLangs[st.Lang]
	if !ok {
		langName = deepThinkLangs[defaultDeepThinkSettings().Lang]
	}
	req := provider.Request{
		Model: st.Model,
		Messages: []provider.Message{
			{Role: "system", Content: "Translate the following text to " + langName + ". Return only the translation, without any explanation, preamble or quotes."},
			{Role: "user", Content: text},
		},
		Temperature:     0.2,
		EnableReasoning: false,
		MaxTokens:       4096,
	}
	resp, err := p.Stream(ctx, req, func(provider.Event) bool { return true })
	if err != nil {
		return "", fmt.Errorf("traduction impossible : %w", err)
	}
	out := strings.TrimSpace(resp.Content)
	if out == "" {
		return "", fmt.Errorf("traduction vide")
	}
	dtCacheSet(key, out)
	return out, nil
}
