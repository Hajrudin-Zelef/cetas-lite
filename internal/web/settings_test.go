package web

import (
	"net/http"
	"strings"
	"testing"
)

func registerAndLogin(t *testing.T, h http.Handler, username string) string {
	t.Helper()
	rec := doJSON(t, h, http.MethodPost, "/api/auth/register", "", map[string]string{
		"username": username, "password": "motdepasse",
	})
	if rec.Code != http.StatusCreated {
		t.Fatalf("register %s status = %d (%s)", username, rec.Code, rec.Body.String())
	}
	rec = doJSON(t, h, http.MethodPost, "/api/auth/login", "", map[string]string{
		"username": username, "password": "motdepasse",
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("login %s status = %d", username, rec.Code)
	}
	token, _ := decode(t, rec)["token"].(string)
	if token == "" {
		t.Fatal("token manquant")
	}
	return token
}

func TestSettingsDefaults(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	rec := doJSON(t, h, http.MethodGet, "/api/settings", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d (%s)", rec.Code, rec.Body.String())
	}
	body := decode(t, rec)
	if body["theme"] != "ocean" || body["family"] != "samagent-n4" || body["mode"] != "standard" {
		t.Fatalf("defauts = %v", body)
	}
}

func TestSettingsRoundTrip(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	rec := doJSON(t, h, http.MethodPut, "/api/settings", token, map[string]any{
		"theme": "clair", "family": "samagent-n8", "mode": "elite",
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("put status = %d (%s)", rec.Code, rec.Body.String())
	}
	rec = doJSON(t, h, http.MethodGet, "/api/settings", token, nil)
	body := decode(t, rec)
	if body["theme"] != "clair" || body["family"] != "samagent-n8" || body["mode"] != "elite" {
		t.Fatalf("round-trip = %v", body)
	}
}

func TestSettingsValidation(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	rec := doJSON(t, h, http.MethodPut, "/api/settings", token, map[string]any{
		"theme": "neon", "family": "samagent-n4", "mode": "standard",
	})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("theme inconnu status = %d", rec.Code)
	}

	rec = doJSON(t, h, http.MethodPut, "/api/settings", token, map[string]any{
		"theme": "ocean", "family": "inexistant", "mode": "standard",
	})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("famille inconnue status = %d", rec.Code)
	}
}

func TestSettingsIsolationPerUser(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	sam := registerAndLogin(t, h, "sam")
	alice := registerAndLogin(t, h, "alice")

	rec := doJSON(t, h, http.MethodPut, "/api/settings", sam, map[string]any{
		"theme": "sombre", "family": "samagent-n4", "mode": "flash",
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("put sam status = %d", rec.Code)
	}
	rec = doJSON(t, h, http.MethodGet, "/api/settings", alice, nil)
	body := decode(t, rec)
	if body["theme"] != "ocean" || body["mode"] != "standard" {
		t.Fatalf("alice ne doit pas voir les reglages de sam: %v", body)
	}
}

func TestSecurityHeaders(t *testing.T) {
	s := newTestServer(t, true)
	rec := doJSON(t, s.Handler(), http.MethodGet, "/", "", nil)
	csp := rec.Header().Get("Content-Security-Policy")
	if !strings.Contains(csp, "script-src 'self'") || !strings.Contains(csp, "style-src 'self' 'unsafe-inline'") {
		t.Fatalf("CSP inattendue: %q", csp)
	}
	if rec.Header().Get("Cache-Control") != "no-cache" {
		t.Fatalf("cache index = %q", rec.Header().Get("Cache-Control"))
	}
}

func TestAssetCacheHeader(t *testing.T) {
	s := newTestServer(t, true)
	req := doJSON(t, s.Handler(), http.MethodGet, "/js/app.js", "", nil)
	if req.Code != http.StatusOK {
		t.Fatalf("status = %d", req.Code)
	}
	if got := req.Header().Get("Cache-Control"); !strings.Contains(got, "max-age=3600") {
		t.Fatalf("cache asset = %q", got)
	}
}

func TestSettingsPartialMerge(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	rec := doJSON(t, h, http.MethodPut, "/api/settings", token, map[string]any{
		"theme": "ocean", "family": "samagent-n8", "mode": "elite",
		"web_default": true,
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("put plein status = %d (%s)", rec.Code, rec.Body.String())
	}

	rec = doJSON(t, h, http.MethodPut, "/api/settings", token, map[string]any{"theme": "clair"})
	if rec.Code != http.StatusOK {
		t.Fatalf("put partiel status = %d (%s)", rec.Code, rec.Body.String())
	}

	rec = doJSON(t, h, http.MethodGet, "/api/settings", token, nil)
	body := decode(t, rec)
	if body["theme"] != "clair" {
		t.Fatalf("theme = %v", body["theme"])
	}
	if body["web_default"] != true {
		t.Fatalf("merge partiel perdu: %v", body)
	}
	if body["family"] != "samagent-n8" || body["mode"] != "elite" {
		t.Fatalf("famille/mode perdus: %v", body)
	}
}

func TestSettingsThinking(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	rec := doJSON(t, h, http.MethodGet, "/api/settings", token, nil)
	body := decode(t, rec)
	if body["thinking_default"] != false || body["thinking_effort"] != "default" {
		t.Fatalf("defauts thinking = %v", body)
	}

	rec = doJSON(t, h, http.MethodPut, "/api/settings", token, map[string]any{
		"thinking_default": true, "thinking_effort": "high",
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("put status = %d (%s)", rec.Code, rec.Body.String())
	}
	rec = doJSON(t, h, http.MethodGet, "/api/settings", token, nil)
	body = decode(t, rec)
	if body["thinking_default"] != true || body["thinking_effort"] != "high" {
		t.Fatalf("round-trip thinking = %v", body)
	}

	rec = doJSON(t, h, http.MethodPut, "/api/settings", token, map[string]any{"thinking_effort": "ultra"})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("effort invalide status = %d", rec.Code)
	}
}

func TestSettingsMaxTokensRoundTrip(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "mtuser")

	rec := doJSON(t, h, http.MethodGet, "/api/settings", token, nil)
	if decode(t, rec)["max_tokens"].(float64) != 4096 {
		t.Fatalf("defaut max_tokens = %v", decode(t, rec)["max_tokens"])
	}
	rec = doJSON(t, h, http.MethodPut, "/api/settings", token, map[string]any{"max_tokens": 8000})
	if rec.Code != http.StatusOK {
		t.Fatalf("put status = %d (%s)", rec.Code, rec.Body.String())
	}
	rec = doJSON(t, h, http.MethodGet, "/api/settings", token, nil)
	if decode(t, rec)["max_tokens"].(float64) != 8000 {
		t.Fatalf("round-trip = %v", decode(t, rec)["max_tokens"])
	}
}

func TestSettingsMaxTokensValidation(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "mtuser2")

	for _, v := range []int{299, 32769, -5} {
		rec := doJSON(t, h, http.MethodPut, "/api/settings", token, map[string]any{"max_tokens": v})
		if rec.Code != http.StatusBadRequest {
			t.Fatalf("max_tokens=%d : status = %d, attendu 400", v, rec.Code)
		}
	}
	// borne haute exacte acceptée
	rec := doJSON(t, h, http.MethodPut, "/api/settings", token, map[string]any{"max_tokens": 32768})
	if rec.Code != http.StatusOK {
		t.Fatalf("max_tokens=32768 : status = %d (%s)", rec.Code, rec.Body.String())
	}
}

func TestSettingsWebSearchMode(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	tok := registerAndLogin(t, h, "sam")

	rec := doJSON(t, h, http.MethodGet, "/api/settings", tok, nil)
	if body := decode(t, rec); body["websearch_mode"] != "auto" {
		t.Fatalf("defaut websearch_mode = %v", body["websearch_mode"])
	}

	rec = doJSON(t, h, http.MethodPut, "/api/settings", tok, map[string]any{"websearch_mode": "natif"})
	if rec.Code != http.StatusOK {
		t.Fatalf("put natif status = %d (%s)", rec.Code, rec.Body.String())
	}
	rec = doJSON(t, h, http.MethodGet, "/api/settings", tok, nil)
	if body := decode(t, rec); body["websearch_mode"] != "natif" {
		t.Fatalf("round-trip websearch_mode = %v", body["websearch_mode"])
	}

	rec = doJSON(t, h, http.MethodPut, "/api/settings", tok, map[string]any{"websearch_mode": "partout"})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("mode invalide status = %d", rec.Code)
	}
}
