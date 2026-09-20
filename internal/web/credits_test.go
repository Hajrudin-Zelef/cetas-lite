package web

import (
	"encoding/json"
	"net/http"
	"testing"
)

func getCreditsList(t *testing.T, h http.Handler, token string) []map[string]any {
	t.Helper()
	rec := doJSON(t, h, http.MethodGet, "/api/settings/credits", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d (%s)", rec.Code, rec.Body.String())
	}
	var list []map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &list); err != nil {
		t.Fatalf("reponse JSON invalide: %s", rec.Body.String())
	}
	return list
}

func TestCreditsGetDefaults(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "credituser1")

	list := getCreditsList(t, h, token)
	if len(list) != 4 {
		t.Fatalf("credits = %d providers, want 4", len(list))
	}
	// Par défaut : non renseigné, pas d'épuisement.
	for _, m := range list {
		if m["credit"] != nil {
			t.Errorf("%s: credit = %v, want nil", m["provider"], m["credit"])
		}
		if m["exhausted"] != false {
			t.Errorf("%s: exhausted = true, want false", m["provider"])
		}
	}
}

func TestCreditsManualRoundTrip(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "credituser2")

	// Saisie manuelle.
	rec := doJSON(t, h, http.MethodPost, "/api/settings/credits", token, map[string]any{
		"provider": "opencode-go", "credit": 12.5,
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("post status = %d (%s)", rec.Code, rec.Body.String())
	}
	// Crédit à 0 → épuisé.
	rec = doJSON(t, h, http.MethodPost, "/api/settings/credits", token, map[string]any{
		"provider": "openrouter", "credit": 0,
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("post zero status = %d (%s)", rec.Code, rec.Body.String())
	}
	// Vérification.
	byID := map[string]map[string]any{}
	for _, m := range getCreditsList(t, h, token) {
		byID[m["provider"].(string)] = m
	}
	if byID["opencode-go"]["credit"] != 12.5 {
		t.Errorf("opencode-go credit = %v, want 12.5", byID["opencode-go"]["credit"])
	}
	if byID["opencode-go"]["source"] != "manuel" {
		t.Errorf("opencode-go source = %v, want manuel", byID["opencode-go"]["source"])
	}
	if byID["openrouter"]["exhausted"] != true {
		t.Error("openrouter exhausted = false, want true (credit 0)")
	}
	// Provider inconnu → 400.
	rec = doJSON(t, h, http.MethodPost, "/api/settings/credits", token, map[string]any{
		"provider": "nope", "credit": 1,
	})
	if rec.Code != http.StatusBadRequest {
		t.Errorf("provider inconnu status = %d, want 400", rec.Code)
	}
	// Crédit négatif → 400.
	rec = doJSON(t, h, http.MethodPost, "/api/settings/credits", token, map[string]any{
		"provider": "deepseek", "credit": -1,
	})
	if rec.Code != http.StatusBadRequest {
		t.Errorf("credit negatif status = %d, want 400", rec.Code)
	}
	// Effacer → retour à non renseigné.
	rec = doJSON(t, h, http.MethodPost, "/api/settings/credits", token, map[string]any{
		"provider": "openrouter", "credit": nil,
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("clear status = %d (%s)", rec.Code, rec.Body.String())
	}
	for _, m := range getCreditsList(t, h, token) {
		if m["provider"] == "openrouter" && m["credit"] != nil {
			t.Errorf("openrouter après effacement = %v, want nil", m["credit"])
		}
	}
}
