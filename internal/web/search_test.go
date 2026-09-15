package web

import (
	"net/http"
	"strings"
	"testing"
)

func searchProvidersByID(t *testing.T, recBody map[string]any) map[string]map[string]any {
	t.Helper()
	list, _ := recBody["providers"].([]any)
	out := map[string]map[string]any{}
	for _, p := range list {
		m, _ := p.(map[string]any)
		id, _ := m["id"].(string)
		out[id] = m
	}
	return out
}

func TestSearchSettingsGet(t *testing.T) {
	s := newTestServer(t, true)
	token := registerAndLogin(t, s.Handler(), "sam")

	rec := doJSON(t, s.Handler(), http.MethodGet, "/api/search/settings", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("get status = %d", rec.Code)
	}
	body := decode(t, rec)
	if body["mode"] != "race" {
		t.Fatalf("mode par defaut attendu race, obtenu %v", body["mode"])
	}
	provs := searchProvidersByID(t, body)
	// Les 5 providers doivent etre listes dans l'ordre demande :
	// Brave -> Tavily -> Jina -> Exa -> DuckDuckGo.
	wantOrder := []string{"brave", "tavily", "jina", "exa", "duckduckgo"}
	list, _ := body["providers"].([]any)
	if len(list) != len(wantOrder) {
		t.Fatalf("attendu %d providers, obtenu %d", len(wantOrder), len(list))
	}
	for i, id := range wantOrder {
		m, _ := list[i].(map[string]any)
		if m["id"] != id {
			t.Fatalf("ordre inattendu a l'index %d: %v", i, m["id"])
		}
	}
	ddg := provs["duckduckgo"]
	if ddg["keyless"] != true {
		t.Fatal("duckduckgo devrait etre marque keyless")
	}
	if ddg["configured"] != true || ddg["enabled"] != true {
		t.Fatal("duckduckgo devrait etre actif sans cle")
	}
	if provs["tavily"]["configured"] != false {
		t.Fatal("tavily ne devrait pas etre configure")
	}
	// Aucune cle ne doit fuiter.
	if strings.Contains(rec.Body.String(), "sk-test") {
		t.Fatal("la reponse ne doit contenir aucune cle")
	}
}

func TestSearchSettingsPut(t *testing.T) {
	t.Setenv("CETAS_LITE_VAULT_PASSWORD", "mot-de-passe-test")
	s := newTestServer(t, true)
	token := registerAndLogin(t, s.Handler(), "sam")

	// Enregistre une cle Tavily + desactive Brave + passe en mode priority.
	rec := doJSON(t, s.Handler(), http.MethodPut, "/api/search/settings", token, map[string]any{
		"mode": "priority",
		"providers": map[string]any{
			"tavily": map[string]any{"key": "tvly-test-123", "enabled": true},
			"brave":  map[string]any{"enabled": false},
		},
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("put status = %d : %s", rec.Code, rec.Body.String())
	}
	provs := searchProvidersByID(t, decode(t, rec))
	if provs["tavily"]["configured"] != true {
		t.Fatal("tavily devrait etre configure apres PUT")
	}
	if provs["brave"]["enabled"] != false {
		t.Fatal("brave devrait etre desactive")
	}

	// La cle est chiffree dans le coffre, jamais en clair.
	rec2 := doJSON(t, s.Handler(), http.MethodGet, "/api/search/settings", token, nil)
	if strings.Contains(rec2.Body.String(), "tvly-test-123") {
		t.Fatal("la cle ne doit jamais ressortir en clair")
	}

	// Suppression de la cle (key vide) => deconfigure.
	rec3 := doJSON(t, s.Handler(), http.MethodPut, "/api/search/settings", token, map[string]any{
		"providers": map[string]any{"tavily": map[string]any{"key": ""}},
	})
	if rec3.Code != http.StatusOK {
		t.Fatalf("put delete status = %d", rec3.Code)
	}
	if searchProvidersByID(t, decode(t, rec3))["tavily"]["configured"] != false {
		t.Fatal("tavily devrait etre deconfigure apres suppression de la cle")
	}

	// Provider inconnu => 400.
	rec4 := doJSON(t, s.Handler(), http.MethodPut, "/api/search/settings", token, map[string]any{
		"providers": map[string]any{"nope": map[string]any{"enabled": true}},
	})
	if rec4.Code != http.StatusBadRequest {
		t.Fatalf("provider inconnu: attendu 400, obtenu %d", rec4.Code)
	}

	// Mode invalide => 400.
	rec5 := doJSON(t, s.Handler(), http.MethodPut, "/api/search/settings", token, map[string]any{
		"mode": "turbo",
	})
	if rec5.Code != http.StatusBadRequest {
		t.Fatalf("mode invalide: attendu 400, obtenu %d", rec5.Code)
	}
}

func TestSearchSettingsPutKeyPersists(t *testing.T) {
	t.Setenv("CETAS_LITE_VAULT_PASSWORD", "mot-de-passe-test")
	s := newTestServer(t, true)
	token := registerAndLogin(t, s.Handler(), "sam")

	rec := doJSON(t, s.Handler(), http.MethodPut, "/api/search/settings", token, map[string]any{
		"providers": map[string]any{"jina": map[string]any{"key": "jina-test-456", "enabled": true}},
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("put status = %d", rec.Code)
	}
	if searchProvidersByID(t, decode(t, rec))["jina"]["configured"] != true {
		t.Fatal("jina devrait etre configure avec une cle")
	}
	rec2 := doJSON(t, s.Handler(), http.MethodGet, "/api/search/settings", token, nil)
	if strings.Contains(rec2.Body.String(), "jina-test-456") {
		t.Fatal("la cle ne doit jamais ressortir en clair")
	}
}

func TestSettingsFeatures(t *testing.T) {
	s := newTestServer(t, true)
	token := registerAndLogin(t, s.Handler(), "sam")

	rec := doJSON(t, s.Handler(), http.MethodPut, "/api/settings", token, map[string]any{
		"tts":            "system",
		"transcription":  "none",
		"prompt_enhance": "openrouter",
		"summarizer":     "deepseek",
		"title_gen":      "conversation",
		"error_analysis": "none",
		"palette":        "violet",
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("put status = %d : %s", rec.Code, rec.Body.String())
	}
	rec2 := doJSON(t, s.Handler(), http.MethodGet, "/api/settings", token, nil)
	body := decode(t, rec2)
	for k, want := range map[string]string{
		"tts": "system", "prompt_enhance": "openrouter", "summarizer": "deepseek",
		"title_gen": "conversation", "palette": "violet",
	} {
		if body[k] != want {
			t.Fatalf("%s: attendu %q, obtenu %v", k, want, body[k])
		}
	}

	// Valeur invalide => 400.
	rec3 := doJSON(t, s.Handler(), http.MethodPut, "/api/settings", token, map[string]any{
		"tts": "INJECTE! DROP",
	})
	if rec3.Code != http.StatusBadRequest {
		t.Fatalf("valeur invalide: attendu 400, obtenu %d", rec3.Code)
	}
	rec4 := doJSON(t, s.Handler(), http.MethodPut, "/api/settings", token, map[string]any{
		"palette": "arc-en-ciel",
	})
	if rec4.Code != http.StatusBadRequest {
		t.Fatalf("palette invalide: attendu 400, obtenu %d", rec4.Code)
	}
}
