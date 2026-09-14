package web

import (
	"net/http"
	"testing"
)

func TestCatalogGroupedByProvider(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	rec := doJSON(t, h, http.MethodGet, "/api/catalog", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
	body := decode(t, rec)
	providers, _ := body["providers"].([]any)
	if len(providers) == 0 {
		t.Fatalf("providers vides: %v", body)
	}
	found := false
	for _, raw := range providers {
		p, _ := raw.(map[string]any)
		if p["id"] != "deepseek" {
			continue
		}
		found = true
		models, _ := p["models"].([]any)
		if len(models) == 0 {
			t.Fatalf("deepseek sans modeles: %v", p)
		}
		m, _ := models[0].(map[string]any)
		if m["id"] == "" || m["label"] == "" {
			t.Fatalf("modele mal forme: %v", m)
		}
	}
	if !found {
		t.Fatal("fournisseur deepseek absent du catalogue")
	}
}

func TestCatalogUnauthorized(t *testing.T) {
	s := newTestServer(t, true)
	rec := doJSON(t, s.Handler(), http.MethodGet, "/api/catalog", "", nil)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d", rec.Code)
	}
}
