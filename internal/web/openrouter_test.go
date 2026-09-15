package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCatalogSelectionRoundTrip(t *testing.T) {
	s := newTestServer(t, true)
	token := registerAndLogin(t, s.Handler(), "sam")

	// Vide au depart.
	rec := doJSON(t, s.Handler(), http.MethodGet, "/api/catalog/selection", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("get status = %d", rec.Code)
	}
	body := decode(t, rec)
	if dis, _ := body["disabled"].(map[string]any); len(dis) != 0 {
		t.Fatalf("selection initiale non vide: %v", dis)
	}

	// Enregistrement puis relecture.
	put := map[string]any{"disabled": map[string]any{
		"openrouter": []string{"qwen/qwen3.7-flash", "poolside/laguna-xs-2.1"},
		"deepseek":   []string{"deepseek-chat"},
	}}
	rec = doJSON(t, s.Handler(), http.MethodPut, "/api/catalog/selection", token, put)
	if rec.Code != http.StatusOK {
		t.Fatalf("put status = %d", rec.Code)
	}
	rec = doJSON(t, s.Handler(), http.MethodGet, "/api/catalog/selection", token, nil)
	body = decode(t, rec)
	dis, _ := body["disabled"].(map[string]any)
	or, _ := dis["openrouter"].([]any)
	if len(or) != 2 {
		t.Fatalf("openrouter: 2 entrees attendues, vu %v", dis["openrouter"])
	}

	// Fournisseur inconnu -> 400.
	rec = doJSON(t, s.Handler(), http.MethodPut, "/api/catalog/selection", token,
		map[string]any{"disabled": map[string]any{"nope": []string{"x"}}})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("fournisseur inconnu: status = %d, attendu 400", rec.Code)
	}

	// Sans auth -> 401.
	rec = doJSON(t, s.Handler(), http.MethodGet, "/api/catalog/selection", "", nil)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("sans auth: status = %d, attendu 401", rec.Code)
	}
}

func TestLocalURLPut(t *testing.T) {
	s := newTestServer(t, true)
	token := registerAndLogin(t, s.Handler(), "sam")

	// URL valide : 200 + persistee + visible dans la liste.
	rec := doJSON(t, s.Handler(), http.MethodPut, "/api/providers/ollama", token,
		map[string]any{"url": "http://192.168.1.10:11434"})
	if rec.Code != http.StatusOK {
		t.Fatalf("put url status = %d", rec.Code)
	}
	rec = doJSON(t, s.Handler(), http.MethodGet, "/api/providers", token, nil)
	provs := providersByID(t, decode(t, rec))
	if provs["ollama"]["url"] != "http://192.168.1.10:11434" {
		t.Fatalf("url non exposee dans la liste: %v", provs["ollama"]["url"])
	}

	// URL invalide -> 400.
	for _, bad := range []string{"", "ftp://x", "notaurl"} {
		rec = doJSON(t, s.Handler(), http.MethodPut, "/api/providers/ollama", token,
			map[string]any{"url": bad})
		if rec.Code != http.StatusBadRequest {
			t.Fatalf("url %q: status = %d, attendu 400", bad, rec.Code)
		}
	}

	// Sur un provider cloud, {"url": ...} n'est pas une cle -> 400 (cle vide).
	rec = doJSON(t, s.Handler(), http.MethodPut, "/api/providers/deepseek", token,
		map[string]any{"url": "http://x"})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("url sur provider cloud: status = %d, attendu 400", rec.Code)
	}
}

func TestOpenRouterModelsCached(t *testing.T) {
	s := newTestServer(t, true)
	token := registerAndLogin(t, s.Handler(), "sam")

	// Serveur simule avec 2 modeles (1 texte, 1 image).
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[
			{"id":"qwen/qwen3-flash","name":"Qwen3 Flash","description":"Rapide.","context_length":100000,
			 "pricing":{"prompt":"0.0000003","completion":"0.0000012"},
			 "architecture":{"input_modalities":["text"],"output_modalities":["text"]}},
			{"id":"google/gemini-img","name":"Gemini Image","description":"Genere des images.","context_length":32000,
			 "pricing":{"prompt":"0.000001","completion":"0.000004"},
			 "architecture":{"input_modalities":["text"],"output_modalities":["image"]}}
		]}`))
	}))
	defer ts.Close()

	oldURL := orModelsURL
	orModelsURL = ts.URL
	defer func() {
		orModelsURL = oldURL
		orCache.mu.Lock()
		orCache.models = nil
		orCache.at = orCache.at.Add(-2 * orCacheTTL)
		orCache.mu.Unlock()
	}()
	// Vide le cache pour forcer le fetch.
	orCache.mu.Lock()
	orCache.models = nil
	orCache.at = orCache.at.Add(-2 * orCacheTTL)
	orCache.mu.Unlock()

	// type=text : seul le modele texte.
	rec := doJSON(t, s.Handler(), http.MethodGet, "/api/openrouter/models?type=text", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("get status = %d", rec.Code)
	}
	body := decode(t, rec)
	models, _ := body["models"].([]any)
	if len(models) != 1 {
		t.Fatalf("type=text: 1 modele attendu, vu %d", len(models))
	}
	m := models[0].(map[string]any)
	if m["id"] != "qwen/qwen3-flash" {
		t.Fatalf("id inattendu: %v", m["id"])
	}
	if m["prompt_per_1m"] != 0.3 {
		t.Fatalf("prix prompt/1M inattendu: %v", m["prompt_per_1m"])
	}

	// type=image : seul le modele image.
	rec = doJSON(t, s.Handler(), http.MethodGet, "/api/openrouter/models?type=image", token, nil)
	models, _ = decode(t, rec)["models"].([]any)
	if len(models) != 1 || models[0].(map[string]any)["id"] != "google/gemini-img" {
		t.Fatalf("type=image inattendu: %v", models)
	}

	// Sans auth -> 401.
	rec = doJSON(t, s.Handler(), http.MethodGet, "/api/openrouter/models", "", nil)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("sans auth: status = %d, attendu 401", rec.Code)
	}
}
