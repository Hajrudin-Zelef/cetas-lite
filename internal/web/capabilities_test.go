package web

import (
	"net/http"
	"testing"
)

func TestCapabilitiesRoundTrip(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "sam")

	rec := doJSON(t, h, http.MethodGet, "/api/capabilities", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("get status = %d", rec.Code)
	}
	if caps, _ := decode(t, rec)["caps"].(map[string]any); len(caps) != 0 {
		t.Fatalf("caps vides attendues: %v", caps)
	}

	rec = doJSON(t, h, http.MethodPut, "/api/capabilities", token, map[string]any{
		"caps": map[string]any{"openrouter/qwen-vl": map[string]any{"vision": true}},
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("put status = %d (%s)", rec.Code, rec.Body.String())
	}
	rec = doJSON(t, h, http.MethodGet, "/api/capabilities", token, nil)
	caps, _ := decode(t, rec)["caps"].(map[string]any)
	entry, _ := caps["openrouter/qwen-vl"].(map[string]any)
	if entry == nil || entry["vision"] != true {
		t.Fatalf("caps = %v", caps)
	}
	if !s.engine.Capabilities().Vision("openrouter", "qwen-vl") {
		t.Fatal("les capacites doivent etre appliquees a l'engine")
	}
}

func TestCapabilitiesUnauthorized(t *testing.T) {
	s := newTestServer(t, true)
	if rec := doJSON(t, s.Handler(), http.MethodGet, "/api/capabilities", "", nil); rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d", rec.Code)
	}
}
