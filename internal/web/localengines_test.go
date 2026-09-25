package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func enginesByID(t *testing.T, recBody map[string]any) map[string]map[string]any {
	t.Helper()
	list, _ := recBody["engines"].([]any)
	out := map[string]map[string]any{}
	for _, e := range list {
		m, _ := e.(map[string]any)
		if m == nil {
			continue
		}
		id, _ := m["id"].(string)
		out[id] = m
	}
	return out
}

func TestLocalEnginesList(t *testing.T) {
	s := newTestServer(t, true)
	token := registerAndLogin(t, s.Handler(), "sam")

	rec := doJSON(t, s.Handler(), http.MethodGet, "/api/local/engines", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("get status = %d", rec.Code)
	}
	engs := enginesByID(t, decode(t, rec))
	for _, id := range []string{"llamacpp", "ollama", "lmstudio"} {
		e := engs[id]
		if e == nil {
			t.Fatalf("moteur %s absent de la liste", id)
		}
		if e["label"] == "" || e["label"] == id {
			t.Fatalf("moteur %s sans libelle lisible: %v", id, e["label"])
		}
		if e["has_key"] != false {
			t.Fatalf("moteur %s ne devrait pas avoir de cle", id)
		}
	}
	// Sans auth -> 401.
	if rec := doJSON(t, s.Handler(), http.MethodGet, "/api/local/engines", "", nil); rec.Code != http.StatusUnauthorized {
		t.Fatalf("get sans token: status = %d, attendu 401", rec.Code)
	}
}

func TestLocalEngineKeyRoundtrip(t *testing.T) {
	t.Setenv("CETAS_LITE_VAULT_PASSWORD", "mot-de-passe-test")
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	// PUT url + cle en une fois.
	rec := doJSON(t, h, http.MethodPut, "/api/local/engines/llamacpp", token,
		map[string]any{"url": "https://example.invalid/v1", "key": "gw-secret-1"})
	if rec.Code != http.StatusOK {
		t.Fatalf("put url+cle: status = %d (%s)", rec.Code, rec.Body.String())
	}
	if strings.Contains(rec.Body.String(), "gw-secret-1") {
		t.Fatal("la cle ne doit pas etre renvoyee en clair")
	}
	// GET : has_key=true, url persistee, pas de fuite.
	rec = doJSON(t, h, http.MethodGet, "/api/local/engines", token, nil)
	engs := enginesByID(t, decode(t, rec))
	e := engs["llamacpp"]
	if e == nil {
		t.Fatal("llamacpp absent de la liste")
	}
	if e["has_key"] != true {
		t.Fatal("llamacpp devrait avoir has_key=true")
	}
	if e["url"] != "https://example.invalid/v1" {
		t.Fatalf("url non persistee: %v", e["url"])
	}
	if strings.Contains(rec.Body.String(), "gw-secret-1") {
		t.Fatal("GET ne doit pas exposer la cle")
	}
	// Le registre live doit contenir le moteur.
	if _, ok := s.registry.Get("llamacpp"); !ok {
		t.Fatal("le registre live devrait contenir llamacpp apres PUT")
	}

	// PUT url seule : la cle survit (registre reconstruit avec).
	rec = doJSON(t, h, http.MethodPut, "/api/local/engines/llamacpp", token,
		map[string]any{"url": "https://example2.invalid/v1"})
	if rec.Code != http.StatusOK {
		t.Fatalf("put url seule: status = %d", rec.Code)
	}
	if decode(t, rec)["has_key"] != true {
		t.Fatal("la cle devrait survivre au changement d'URL")
	}

	// DELETE : supprime la cle, conserve l'URL.
	rec = doJSON(t, h, http.MethodDelete, "/api/local/engines/llamacpp", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("delete local: status = %d", rec.Code)
	}
	rec = doJSON(t, h, http.MethodGet, "/api/local/engines", token, nil)
	engs = enginesByID(t, decode(t, rec))
	if engs["llamacpp"]["has_key"] == true {
		t.Fatal("has_key devrait etre false apres DELETE")
	}
	if engs["llamacpp"]["url"] != "https://example2.invalid/v1" {
		t.Fatal("l'URL devrait etre conservee apres DELETE")
	}
	// DELETE sur moteur sans cle : idempotent, 200.
	if rec := doJSON(t, h, http.MethodDelete, "/api/local/engines/ollama", token, nil); rec.Code != http.StatusOK {
		t.Fatalf("delete local sans cle: status = %d, attendu 200", rec.Code)
	}
}

func TestLocalEnginePutErrors(t *testing.T) {
	t.Setenv("CETAS_LITE_VAULT_PASSWORD", "mot-de-passe-test")
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	// Moteur inconnu -> 404.
	if rec := doJSON(t, h, http.MethodPut, "/api/local/engines/nope", token, map[string]any{"key": "x"}); rec.Code != http.StatusNotFound {
		t.Fatalf("put moteur inconnu: status = %d, attendu 404", rec.Code)
	}
	// Corps vide -> 400.
	if rec := doJSON(t, h, http.MethodPut, "/api/local/engines/ollama", token, map[string]any{}); rec.Code != http.StatusBadRequest {
		t.Fatalf("put vide: status = %d, attendu 400", rec.Code)
	}
	// URL invalide -> 400.
	if rec := doJSON(t, h, http.MethodPut, "/api/local/engines/ollama", token, map[string]any{"url": "notaurl"}); rec.Code != http.StatusBadRequest {
		t.Fatalf("put url invalide: status = %d, attendu 400", rec.Code)
	}
	// Cle seule sans URL configuree : 200, mais aucun endpoint vide
	// n'entre dans le registre.
	rec := doJSON(t, h, http.MethodPut, "/api/local/engines/ollama", token, map[string]any{"key": "k-seule"})
	if rec.Code != http.StatusOK {
		t.Fatalf("put cle seule: status = %d", rec.Code)
	}
	if decode(t, rec)["url"] != "" {
		t.Fatal("put cle seule: url devrait rester vide")
	}
	if _, ok := s.registry.Get("ollama"); ok {
		t.Fatal("put cle seule sans URL: le registre ne devrait pas contenir d'endpoint vide")
	}
}

func TestLocalModelsURLNoDoubleV1(t *testing.T) {
	cases := map[string]string{
		"https://hote.example/v1":   "https://hote.example/v1/models",
		"https://hote.example/v1/":  "https://hote.example/v1/models",
		"http://192.168.1.10:11434": "http://192.168.1.10:11434/v1/models",
		"http://127.0.0.1:8080/":    "http://127.0.0.1:8080/v1/models",
	}
	for in, want := range cases {
		if got := localModelsURL(in); got != want {
			t.Fatalf("localModelsURL(%q) = %q, attendu %q", in, got, want)
		}
	}
}

func TestLocalEngineTestEndpoint(t *testing.T) {
	t.Setenv("CETAS_LITE_VAULT_PASSWORD", "mot-de-passe-test")
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	// Sans URL : ok=false.
	rec := doJSON(t, h, http.MethodPost, "/api/local/engines/ollama/test", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("test sans url: status HTTP = %d", rec.Code)
	}
	if body := decode(t, rec); body["ok"] != false {
		t.Fatalf("test sans url: ok devrait etre false, recu %v", body)
	}

	// Faux moteur : verifie l'auth Bearer, le chemin /v1/models (pas de
	// double /v1) et le comptage des modeles (2 formes, dedupliques).
	var gotAuth, gotPath string
	fake := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":[{"id":"m1"},{"id":"m2"}],"models":[{"name":"m2"},{"name":"m3"}]}`))
	}))
	t.Cleanup(fake.Close)
	rec = doJSON(t, h, http.MethodPut, "/api/local/engines/ollama", token,
		map[string]any{"url": fake.URL + "/v1", "key": "k-test"})
	if rec.Code != http.StatusOK {
		t.Fatalf("put url+cle: status = %d", rec.Code)
	}
	rec = doJSON(t, h, http.MethodPost, "/api/local/engines/ollama/test", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("test: status HTTP = %d", rec.Code)
	}
	body := decode(t, rec)
	if body["ok"] != true {
		t.Fatalf("test: ok devrait etre true, recu %v", body)
	}
	if body["models"] != float64(3) {
		t.Fatalf("test: models = %v, attendu 3 (dedupliques)", body["models"])
	}
	if _, ok := body["latency_ms"]; !ok {
		t.Fatal("test: latency_ms manquante")
	}
	if gotAuth != "Bearer k-test" {
		t.Fatalf("test: Authorization = %q, attendu Bearer avec la cle du coffre", gotAuth)
	}
	if gotPath != "/v1/models" {
		t.Fatalf("test: chemin sonde = %q, attendu /v1/models (pas de double /v1)", gotPath)
	}

	// Moteur inconnu -> 404 ; sans auth -> 401.
	if rec := doJSON(t, h, http.MethodPost, "/api/local/engines/nope/test", token, nil); rec.Code != http.StatusNotFound {
		t.Fatalf("test moteur inconnu: status = %d, attendu 404", rec.Code)
	}
	if rec := doJSON(t, h, http.MethodPost, "/api/local/engines/ollama/test", "", nil); rec.Code != http.StatusUnauthorized {
		t.Fatalf("test sans token: status = %d, attendu 401", rec.Code)
	}
}
