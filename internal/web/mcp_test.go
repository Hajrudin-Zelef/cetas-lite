package web

import (
	"net/http"
	"testing"
)

func TestMCPStatusEmpty(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	rec := doJSON(t, h, http.MethodGet, "/api/mcp", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d (%s)", rec.Code, rec.Body.String())
	}
	body := decode(t, rec)
	servers, ok := body["servers"].([]any)
	if !ok {
		t.Fatalf("servers attendu en tableau: %T", body["servers"])
	}
	if len(servers) != 0 {
		t.Fatalf("serveurs attendus vides: %v", servers)
	}

	rec = doJSON(t, h, http.MethodGet, "/api/mcp?probe=1", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("probe status = %d (%s)", rec.Code, rec.Body.String())
	}
	if got, _ := decode(t, rec)["servers"].([]any); len(got) != 0 {
		t.Fatalf("probe sans serveur = %v", got)
	}
}

func TestMCPStatusUnauthorized(t *testing.T) {
	s := newTestServer(t, true)
	rec := doJSON(t, s.Handler(), http.MethodGet, "/api/mcp", "", nil)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d", rec.Code)
	}
}

func TestSettingsMCPDefault(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	rec := doJSON(t, h, http.MethodGet, "/api/settings", token, nil)
	if _, present := decode(t, rec)["mcp_default"]; present {
		t.Fatal("mcp_default ne doit pas etre present par defaut (auto = true)")
	}

	rec = doJSON(t, h, http.MethodPut, "/api/settings", token, map[string]any{"mcp_default": false})
	if rec.Code != http.StatusOK {
		t.Fatalf("put status = %d (%s)", rec.Code, rec.Body.String())
	}
	rec = doJSON(t, h, http.MethodGet, "/api/settings", token, nil)
	if decode(t, rec)["mcp_default"] != false {
		t.Fatalf("mcp_default = %v", decode(t, rec)["mcp_default"])
	}

	rec = doJSON(t, h, http.MethodPut, "/api/settings", token, map[string]any{"theme": "sombre"})
	if rec.Code != http.StatusOK {
		t.Fatalf("put partiel status = %d", rec.Code)
	}
	rec = doJSON(t, h, http.MethodGet, "/api/settings", token, nil)
	if decode(t, rec)["mcp_default"] != false {
		t.Fatalf("merge partiel a perdu mcp_default: %v", decode(t, rec))
	}
}
