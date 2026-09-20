package web

// Module Paramètres → Crédit API.
//
// Suivi du crédit restant par provider (OpenCode Zen/Go, OpenRouter,
// DeepSeek). L'utilisateur renseigne son crédit manuellement ; pour
// OpenRouter et DeepSeek, un bouton « actualiser » interroge l'API
// officielle (OpenCode n'expose pas d'endpoint public → manuel uniquement).
//
// Si le crédit d'un provider est épuisé (<= 0), le Sélecteur affiche un
// avertissement clair au moment de choisir un modèle — jamais une erreur
// silencieuse. Crédit non renseigné (nil) = affiché « non renseigné »,
// sans blocage.

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"cetas-lite/internal/vault"
)

const creditsSettingKey = "api_credits"

// creditProviderDef décrit un provider suivi.
type creditProviderDef struct {
	ID         string
	Label      string
	ManualOnly bool
}

var creditProviders = []creditProviderDef{
	{ID: "opencode-zen", Label: "OpenCode Zen", ManualOnly: true},
	{ID: "opencode-go", Label: "OpenCode Go", ManualOnly: true},
	{ID: "openrouter", Label: "OpenRouter", ManualOnly: false},
	{ID: "deepseek", Label: "DeepSeek", ManualOnly: false},
}

// creditEntry est la valeur stockée par provider.
type creditEntry struct {
	Credit    *float64 `json:"credit"`     // nil = non renseigné
	Source    string   `json:"source"`     // "manuel" ou "api"
	UpdatedAt string   `json:"updated_at"` // RFC3339, "" si jamais
}

// apiCredit est la forme exposée au frontend.
type apiCredit struct {
	Provider   string   `json:"provider"`
	Label      string   `json:"label"`
	ManualOnly bool     `json:"manual_only"`
	Credit     *float64 `json:"credit"`
	Source     string   `json:"source"`
	UpdatedAt  string   `json:"updated_at"`
	// Exhausted : crédit renseigné et <= 0.
	Exhausted bool `json:"exhausted"`
}

func (s *Server) loadCredits(user string) map[string]creditEntry {
	out := map[string]creditEntry{}
	if raw, ok := s.st.GetSetting(user, creditsSettingKey); ok {
		_ = json.Unmarshal(raw, &out)
	}
	return out
}

func (s *Server) saveCredits(user string, m map[string]creditEntry) error {
	raw, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return s.st.PutSetting(user, creditsSettingKey, raw)
}

func (s *Server) creditsList(user string) []apiCredit {
	stored := s.loadCredits(user)
	out := make([]apiCredit, 0, len(creditProviders))
	for _, def := range creditProviders {
		e := stored[def.ID]
		ac := apiCredit{
			Provider:   def.ID,
			Label:      def.Label,
			ManualOnly: def.ManualOnly,
			Credit:     e.Credit,
			Source:     e.Source,
			UpdatedAt:  e.UpdatedAt,
		}
		if e.Credit != nil && *e.Credit <= 0 {
			ac.Exhausted = true
		}
		out = append(out, ac)
	}
	return out
}

// GET /api/settings/credits — état des crédits par provider.
func (s *Server) handleCreditsGet(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	writeJSON(w, http.StatusOK, s.creditsList(claims.Username))
}

// POST /api/settings/credits — saisie manuelle {provider, credit}.
// credit null/omise = effacer (repasser à « non renseigné »).
func (s *Server) handleCreditsPost(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		Provider string   `json:"provider"`
		Credit   *float64 `json:"credit"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "corps invalide")
		return
	}
	known := false
	for _, def := range creditProviders {
		if def.ID == body.Provider {
			known = true
			break
		}
	}
	if !known {
		writeError(w, http.StatusBadRequest, "provider inconnu")
		return
	}
	if body.Credit != nil && *body.Credit < 0 {
		writeError(w, http.StatusBadRequest, "le credit ne peut pas etre negatif")
		return
	}
	stored := s.loadCredits(claims.Username)
	if body.Credit == nil {
		delete(stored, body.Provider)
	} else {
		stored[body.Provider] = creditEntry{
			Credit:    body.Credit,
			Source:    "manuel",
			UpdatedAt: time.Now().UTC().Format(time.RFC3339),
		}
	}
	if err := s.saveCredits(claims.Username, stored); err != nil {
		writeError(w, http.StatusInternalServerError, "sauvegarde impossible")
		return
	}
	writeJSON(w, http.StatusOK, s.creditsList(claims.Username))
}

// providerAPIKey lit la clé API d'un provider depuis le coffre chiffré.
func (s *Server) providerAPIKey(providerID string) (string, bool) {
	v, err := vault.Open(s.st)
	if err != nil {
		return "", false
	}
	ct, ok := s.st.GetSecret(providerID)
	if !ok {
		return "", false
	}
	pt, err := v.Decrypt(ct, []byte(providerID))
	if err != nil {
		return "", false
	}
	if len(pt) == 0 {
		return "", false
	}
	return string(pt), true
}

// fetchOpenRouterCredit interroge GET /api/v1/credits.
// Réponse : {"data": {"total_credits": X, "total_usage": Y}}.
func (s *Server) fetchOpenRouterCredit(ctx context.Context, key string) (float64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://openrouter.ai/api/v1/credits", nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Authorization", "Bearer "+key)
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("openrouter: HTTP %d", resp.StatusCode)
	}
	var body struct {
		Data struct {
			TotalCredits float64 `json:"total_credits"`
			TotalUsage   float64 `json:"total_usage"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return 0, err
	}
	return body.Data.TotalCredits - body.Data.TotalUsage, nil
}

// fetchDeepSeekCredit interroge GET /api.deepseek.com/user/balance.
// Réponse : {"is_available": true, "balance_infos": [{"currency":"USD","total_balance": X}]}.
func (s *Server) fetchDeepSeekCredit(ctx context.Context, key string) (float64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.deepseek.com/user/balance", nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Authorization", "Bearer "+key)
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("deepseek: HTTP %d", resp.StatusCode)
	}
	var body struct {
		IsAvailable  bool `json:"is_available"`
		BalanceInfos []struct {
			Currency     string  `json:"currency"`
			TotalBalance float64 `json:"total_balance"`
		} `json:"balance_infos"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return 0, err
	}
	if !body.IsAvailable {
		return 0, fmt.Errorf("deepseek: balance indisponible")
	}
	var total float64
	for _, b := range body.BalanceInfos {
		total += b.TotalBalance
	}
	return total, nil
}

// POST /api/settings/credits/refresh — actualise OpenRouter et DeepSeek
// via leurs APIs. Réponse : liste + erreurs éventuelles par provider.
func (s *Server) handleCreditsRefresh(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	stored := s.loadCredits(claims.Username)
	errs := map[string]string{}
	now := time.Now().UTC().Format(time.RFC3339)
	ctx, cancel := context.WithTimeout(r.Context(), 20*time.Second)
	defer cancel()

	refresh := func(providerID string, fetch func(context.Context, string) (float64, error)) {
		key, ok := s.providerAPIKey(providerID)
		if !ok {
			errs[providerID] = "cle API absente du coffre"
			return
		}
		credit, err := fetch(ctx, key)
		if err != nil {
			errs[providerID] = err.Error()
			return
		}
		c := credit
		stored[providerID] = creditEntry{Credit: &c, Source: "api", UpdatedAt: now}
	}
	refresh("openrouter", s.fetchOpenRouterCredit)
	refresh("deepseek", s.fetchDeepSeekCredit)

	if err := s.saveCredits(claims.Username, stored); err != nil {
		writeError(w, http.StatusInternalServerError, "sauvegarde impossible")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"credits": s.creditsList(claims.Username),
		"errors":  errs,
	})
}
