package web

import (
	"encoding/json"
	"net/http"

	"cetas-lite/internal/alias"
)

type uiSettings struct {
	Theme      string `json:"theme"`
	Family     string `json:"family"`
	Mode       string `json:"mode"`
	AppMode    string `json:"app_mode"`
	WebDefault bool   `json:"web_default"`
	MCPDefault *bool  `json:"mcp_default,omitempty"`
	// WebSearchMode : "auto" (defaut), "natif", "outils" ou "off".
	WebSearchMode   string `json:"websearch_mode"`
	ThinkingDefault bool   `json:"thinking_default"`
	ThinkingEffort  string `json:"thinking_effort"`
	// MaxTokens : tokens max par reponse (300..32768, defaut 4096).
	MaxTokens int `json:"max_tokens"`
	// Palette d'accent : "bleu" (defaut), "violet", "vert", "vert_pur",
	// "bleu_ocean", "jaune_or", "rouge".
	Palette string `json:"palette"`
	// Fonctionnalites (module CETAS uniquement) : choix par fonctionnalite.
	// TTS / Transcription : "system" (navigateur) ou "none".
	TTS           string `json:"tts"`
	Transcription string `json:"transcription"`
	// Les autres : id provider, "none", et "conversation" pour TitleGen.
	PromptEnhance string `json:"prompt_enhance"`
	Summarizer    string `json:"summarizer"`
	TitleGen      string `json:"title_gen"`
	ErrorAnalysis string `json:"error_analysis"`
	// PregenSuggestions : pré-génération des questions suggérées (lot 5).
	// Pointeur pour distinguer « non défini » (défaut : activé) de
	// « désactivé ». nil = activé.
	PregenSuggestions *bool `json:"pregen_suggestions,omitempty"`
}

// pregenEnabled rend l'état du réglage de pré-génération : activé par
// défaut (nil), désactivé uniquement si explicitement false.
func (u uiSettings) pregenEnabled() bool {
	return u.PregenSuggestions == nil || *u.PregenSuggestions
}

func validEffort(e string) bool {
	switch e {
	case "", "default", "low", "medium", "high":
		return true
	}
	return false
}

func defaultUISettings() uiSettings {
	return uiSettings{
		Theme: "clair", Family: "samagent-n4", Mode: "standard", AppMode: "chat",
		WebSearchMode: "auto", MaxTokens: 4096,
		Palette: "bleu",
		TTS:     "system", Transcription: "system",
		PromptEnhance: "none", Summarizer: "none",
		TitleGen: "conversation", ErrorAnalysis: "none",
	}
}

func validAppMode(m string) bool {
	return m == "chat" || m == "agent"
}

// clampMaxTokensSetting borne le reglage "tokens max par reponse".
func clampMaxTokensSetting(n int) int {
	if n < 300 {
		return 300
	}
	if n > 32768 {
		return 32768
	}
	return n
}

func validTheme(t string) bool {
	switch t {
	case "ocean", "sombre", "clair", "hard_dark":
		return true
	}
	return false
}

// validPalette valide la palette d'accent.
func validPalette(p string) bool {
	switch p {
	case "", "bleu", "violet", "vert", "vert_pur", "bleu_ocean", "jaune_or", "rouge":
		return true
	}
	return false
}

// validFeatureValue valide un choix de fonctionnalite : "none", "system",
// "conversation" ou un identifiant de provider.
func validFeatureValue(v string) bool {
	switch v {
	case "", "none", "system", "conversation":
		return true
	}
	if len(v) > 64 {
		return false
	}
	for _, r := range v {
		if r != '_' && r != '-' && (r < 'a' || r > 'z') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}

func validWebSearchMode(m string) bool {
	switch m {
	case "", "auto", "natif", "outils", "off":
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
	// "ocean" est l'ancien nom du theme clair.
	if out.Theme == "ocean" {
		out.Theme = "clair"
	}
	if !validPalette(out.Palette) {
		out.Palette = defaultUISettings().Palette
	}
	for _, f := range []*string{&out.TTS, &out.Transcription, &out.PromptEnhance, &out.Summarizer, &out.TitleGen, &out.ErrorAnalysis} {
		if !validFeatureValue(*f) {
			*f = "none"
		}
	}
	if out.TTS == "" {
		out.TTS = "system"
	}
	if out.Transcription == "" {
		out.Transcription = "system"
	}
	if out.TitleGen == "" {
		out.TitleGen = "conversation"
	}
	if !validAppMode(out.AppMode) {
		out.AppMode = defaultUISettings().AppMode
	}
	if !validWebSearchMode(out.WebSearchMode) {
		out.WebSearchMode = "auto"
	}
	if out.ThinkingEffort == "" || !validEffort(out.ThinkingEffort) {
		out.ThinkingEffort = "default"
	}
	if out.MaxTokens <= 0 {
		out.MaxTokens = defaultUISettings().MaxTokens
	}
	out.MaxTokens = clampMaxTokensSetting(out.MaxTokens)
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
		Theme             *string `json:"theme"`
		Family            *string `json:"family"`
		Mode              *string `json:"mode"`
		AppMode           *string `json:"app_mode"`
		WebDefault        *bool   `json:"web_default"`
		MCPDefault        *bool   `json:"mcp_default"`
		WebSearchMode     *string `json:"websearch_mode"`
		ThinkingDefault   *bool   `json:"thinking_default"`
		ThinkingEffort    *string `json:"thinking_effort"`
		MaxTokens         *int    `json:"max_tokens"`
		Palette           *string `json:"palette"`
		TTS               *string `json:"tts"`
		Transcription     *string `json:"transcription"`
		PromptEnhance     *string `json:"prompt_enhance"`
		PregenSuggestions *bool   `json:"pregen_suggestions"`
		Summarizer        *string `json:"summarizer"`
		TitleGen          *string `json:"title_gen"`
		ErrorAnalysis     *string `json:"error_analysis"`
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
	if body.AppMode != nil {
		if !validAppMode(*body.AppMode) {
			writeError(w, http.StatusBadRequest, "mode applicatif invalide")
			return
		}
		cur.AppMode = *body.AppMode
	}
	if body.WebDefault != nil {
		cur.WebDefault = *body.WebDefault
	}
	if body.WebSearchMode != nil {
		if !validWebSearchMode(*body.WebSearchMode) {
			writeError(w, http.StatusBadRequest, "websearch_mode invalide (auto|natif|outils|off)")
			return
		}
		cur.WebSearchMode = *body.WebSearchMode
		if cur.WebSearchMode == "" {
			cur.WebSearchMode = "auto"
		}
	}
	if body.MCPDefault != nil {
		cur.MCPDefault = body.MCPDefault
	}
	if body.PregenSuggestions != nil {
		cur.PregenSuggestions = body.PregenSuggestions
	}
	if body.ThinkingDefault != nil {
		cur.ThinkingDefault = *body.ThinkingDefault
	}
	if body.ThinkingEffort != nil {
		if !validEffort(*body.ThinkingEffort) {
			writeError(w, http.StatusBadRequest, "effort invalide")
			return
		}
		cur.ThinkingEffort = *body.ThinkingEffort
	}
	if body.MaxTokens != nil {
		if *body.MaxTokens < 300 || *body.MaxTokens > 32768 {
			writeError(w, http.StatusBadRequest, "max_tokens invalide (300..32768)")
			return
		}
		cur.MaxTokens = *body.MaxTokens
	}
	if body.Palette != nil {
		if !validPalette(*body.Palette) {
			writeError(w, http.StatusBadRequest, "palette invalide")
			return
		}
		cur.Palette = *body.Palette
		if cur.Palette == "" {
			cur.Palette = defaultUISettings().Palette
		}
	}
	features := []struct {
		name string
		body **string
		cur  *string
	}{
		{"tts", &body.TTS, &cur.TTS},
		{"transcription", &body.Transcription, &cur.Transcription},
		{"prompt_enhance", &body.PromptEnhance, &cur.PromptEnhance},
		{"summarizer", &body.Summarizer, &cur.Summarizer},
		{"title_gen", &body.TitleGen, &cur.TitleGen},
		{"error_analysis", &body.ErrorAnalysis, &cur.ErrorAnalysis},
	}
	for _, f := range features {
		if *f.body == nil {
			continue
		}
		if !validFeatureValue(**f.body) {
			writeError(w, http.StatusBadRequest, "fonctionnalite invalide: "+f.name)
			return
		}
		*f.cur = **f.body
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
