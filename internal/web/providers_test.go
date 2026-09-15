package web

import (
	"net/http"
	"strings"
	"testing"
)

func providersByID(t *testing.T, recBody map[string]any) map[string]map[string]any {
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

func TestProvidersList(t *testing.T) {
	s := newTestServer(t, true)
	token := registerAndLogin(t, s.Handler(), "sam")

	rec := doJSON(t, s.Handler(), http.MethodGet, "/api/providers", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("get status = %d", rec.Code)
	}
	provs := providersByID(t, decode(t, rec))
	// Les providers cloud connus doivent tous etre listes, non configures.
	for _, id := range []string{"deepseek", "openrouter", "opencode", "opencode-go"} {
		p := provs[id]
		if p == nil {
			t.Fatalf("provider %s absent de la liste", id)
		}
		if p["label"] == "" || p["label"] == id {
			t.Fatalf("provider %s sans libelle lisible: %v", id, p["label"])
		}
		if p["configured"] != false {
			t.Fatalf("provider %s ne devrait pas etre configure", id)
		}
		if p["local"] != false {
			t.Fatalf("provider %s ne devrait pas etre local", id)
		}
	}
	// Les moteurs locaux sont listes avec local=true.
	for _, id := range []string{"llamacpp", "ollama", "lmstudio"} {
		p := provs[id]
		if p == nil {
			t.Fatalf("moteur local %s absent de la liste", id)
		}
		if p["local"] != true {
			t.Fatalf("moteur %s devrait etre marque local", id)
		}
	}
	// Aucune cle ne doit fuiter dans la reponse.
	if strings.Contains(rec.Body.String(), "sk-") {
		t.Fatal("la reponse ne doit contenir aucune cle")
	}
}

func TestProvidersPutGetDelete(t *testing.T) {
	t.Setenv("CETAS_LITE_VAULT_PASSWORD", "mot-de-passe-test")
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	// PUT avec une cle valide.
	rec := doJSON(t, h, http.MethodPut, "/api/providers/deepseek", token, map[string]any{"key": "sk-test-123"})
	if rec.Code != http.StatusOK {
		t.Fatalf("put status = %d (%s)", rec.Code, rec.Body.String())
	}
	// La cle ne doit pas apparaitre en clair dans la reponse.
	if strings.Contains(rec.Body.String(), "sk-test-123") {
		t.Fatal("la cle ne doit pas etre renvoyee en clair")
	}
	// Le registre live doit contenir le provider.
	if _, ok := s.registry.Get("deepseek"); !ok {
		t.Fatal("le registre live devrait contenir deepseek apres PUT")
	}
	// GET doit signaler configured=true, sans exposer la cle.
	rec = doJSON(t, h, http.MethodGet, "/api/providers", token, nil)
	provs := providersByID(t, decode(t, rec))
	if provs["deepseek"]["configured"] != true {
		t.Fatal("deepseek devrait etre marque configure")
	}
	if strings.Contains(rec.Body.String(), "sk-test-123") {
		t.Fatal("GET ne doit pas exposer la cle")
	}

	// DELETE supprime le secret et retire du registre.
	rec = doJSON(t, h, http.MethodDelete, "/api/providers/deepseek", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("delete status = %d (%s)", rec.Code, rec.Body.String())
	}
	if _, ok := s.registry.Get("deepseek"); ok {
		t.Fatal("le registre live ne devrait plus contenir deepseek apres DELETE")
	}
	rec = doJSON(t, h, http.MethodGet, "/api/providers", token, nil)
	if providersByID(t, decode(t, rec))["deepseek"]["configured"] != false {
		t.Fatal("deepseek ne devrait plus etre marque configure")
	}
}

func TestProvidersPutErrors(t *testing.T) {
	t.Setenv("CETAS_LITE_VAULT_PASSWORD", "mot-de-passe-test")
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	// Provider inconnu -> 404.
	if rec := doJSON(t, h, http.MethodPut, "/api/providers/nope", token, map[string]any{"key": "x"}); rec.Code != http.StatusNotFound {
		t.Fatalf("put inconnu: status = %d, attendu 404", rec.Code)
	}
	// Cle vide -> 400.
	if rec := doJSON(t, h, http.MethodPut, "/api/providers/deepseek", token, map[string]any{"key": "  "}); rec.Code != http.StatusBadRequest {
		t.Fatalf("put cle vide: status = %d, attendu 400", rec.Code)
	}
	// Corps invalide -> 400.
	if rec := doJSON(t, h, http.MethodPut, "/api/providers/deepseek", token, "pas-un-objet"); rec.Code != http.StatusBadRequest {
		t.Fatalf("put corps invalide: status = %d, attendu 400", rec.Code)
	}
	// DELETE sur provider inconnu -> 404.
	if rec := doJSON(t, h, http.MethodDelete, "/api/providers/nope", token, nil); rec.Code != http.StatusNotFound {
		t.Fatalf("delete inconnu: status = %d, attendu 404", rec.Code)
	}
	// Provider local : PUT cle refuse (pas de cle API ici, seule l'URL est gerable).
	if rec := doJSON(t, h, http.MethodPut, "/api/providers/ollama", token, map[string]any{"key": "x"}); rec.Code != http.StatusBadRequest {
		t.Fatalf("put local cle: status = %d, attendu 400", rec.Code)
	}
	// Provider local : PUT URL invalide -> 400.
	if rec := doJSON(t, h, http.MethodPut, "/api/providers/ollama", token, map[string]any{"url": "notaurl"}); rec.Code != http.StatusBadRequest {
		t.Fatalf("put local url invalide: status = %d, attendu 400", rec.Code)
	}
	// Provider local : PUT URL valide -> 200.
	if rec := doJSON(t, h, http.MethodPut, "/api/providers/ollama", token, map[string]any{"url": "http://192.168.1.10:11434"}); rec.Code != http.StatusOK {
		t.Fatalf("put local url: status = %d, attendu 200", rec.Code)
	}
	// Provider cloud : PUT URL refuse -> 400.
	if rec := doJSON(t, h, http.MethodPut, "/api/providers/deepseek", token, map[string]any{"url": "http://x"}); rec.Code != http.StatusBadRequest {
		t.Fatalf("put cloud url: status = %d, attendu 400", rec.Code)
	}
}

func TestProvidersUnauthorized(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	if rec := doJSON(t, h, http.MethodGet, "/api/providers", "", nil); rec.Code != http.StatusUnauthorized {
		t.Fatalf("get sans token: status = %d, attendu 401", rec.Code)
	}
	if rec := doJSON(t, h, http.MethodPut, "/api/providers/deepseek", "", map[string]any{"key": "x"}); rec.Code != http.StatusUnauthorized {
		t.Fatalf("put sans token: status = %d, attendu 401", rec.Code)
	}
	if rec := doJSON(t, h, http.MethodDelete, "/api/providers/deepseek", "", nil); rec.Code != http.StatusUnauthorized {
		t.Fatalf("delete sans token: status = %d, attendu 401", rec.Code)
	}
}

func TestProvidersPutWithoutVaultPassword(t *testing.T) {
	// Sans CETAS_LITE_VAULT_PASSWORD, le chiffrement est impossible -> 500.
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")
	rec := doJSON(t, h, http.MethodPut, "/api/providers/deepseek", token, map[string]any{"key": "sk-x"})
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("put sans mot de passe vault: status = %d, attendu 500", rec.Code)
	}
}
