package web

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/auth"
	"cetas-lite/internal/chat"
	"cetas-lite/internal/config"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/store"
)

func newTestServer(t *testing.T, registrationOpen bool) *Server {
	t.Helper()
	t.Setenv("CETAS_LITE_HOME", t.TempDir())
	t.Setenv("CETAS_LITE_REGISTRATION_OPEN", strconv.FormatBool(registrationOpen))
	cfg, err := config.Load()
	if err != nil {
		t.Fatal(err)
	}
	st, err := store.Open(cfg.DBPath)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })
	m, err := auth.New(st)
	if err != nil {
		t.Fatal(err)
	}
	engine := chat.NewEngine(provider.NewRegistry(), alias.Defaults(), st, nil, "")
	return New(cfg, st, m, engine, "test")
}

func doJSON(t *testing.T, h http.Handler, method, path, token string, body any) *httptest.ResponseRecorder {
	t.Helper()
	var reader io.Reader
	if body != nil {
		raw, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		reader = bytes.NewReader(raw)
	}
	req := httptest.NewRequest(method, path, reader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func decode(t *testing.T, rec *httptest.ResponseRecorder) map[string]any {
	t.Helper()
	var out map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("reponse JSON invalide (%d): %s", rec.Code, rec.Body.String())
	}
	return out
}

func TestHealth(t *testing.T) {
	s := newTestServer(t, true)
	rec := doJSON(t, s.Handler(), http.MethodGet, "/api/health", "", nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
	if got := decode(t, rec)["status"]; got != "ok" {
		t.Fatalf("status = %v", got)
	}
}

func TestConfig(t *testing.T) {
	s := newTestServer(t, true)
	rec := doJSON(t, s.Handler(), http.MethodGet, "/api/config", "", nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
	if got := decode(t, rec)["registration_open"]; got != true {
		t.Fatalf("registration_open = %v", got)
	}
}

func TestRegisterLoginMe(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()

	rec := doJSON(t, h, http.MethodPost, "/api/auth/register", "", map[string]string{
		"username": "sam", "password": "motdepasse",
	})
	if rec.Code != http.StatusCreated {
		t.Fatalf("register status = %d (%s)", rec.Code, rec.Body.String())
	}

	rec = doJSON(t, h, http.MethodPost, "/api/auth/login", "", map[string]string{
		"username": "sam", "password": "motdepasse",
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("login status = %d (%s)", rec.Code, rec.Body.String())
	}
	body := decode(t, rec)
	token, _ := body["token"].(string)
	if token == "" {
		t.Fatal("token manquant")
	}

	rec = doJSON(t, h, http.MethodGet, "/api/me", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("me status = %d (%s)", rec.Code, rec.Body.String())
	}
	if got := decode(t, rec)["username"]; got != "sam" {
		t.Fatalf("me username = %v", got)
	}
}

func TestMeUnauthorized(t *testing.T) {
	s := newTestServer(t, true)
	rec := doJSON(t, s.Handler(), http.MethodGet, "/api/me", "", nil)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d", rec.Code)
	}
	rec = doJSON(t, s.Handler(), http.MethodGet, "/api/me", "faux-token", nil)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("token invalide status = %d", rec.Code)
	}
}

func TestRegisterClosed(t *testing.T) {
	s := newTestServer(t, false)
	rec := doJSON(t, s.Handler(), http.MethodPost, "/api/auth/register", "", map[string]string{
		"username": "sam", "password": "motdepasse",
	})
	if rec.Code != http.StatusForbidden {
		t.Fatalf("status = %d", rec.Code)
	}
}

func TestRegisterDuplicateAndShort(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	_ = doJSON(t, h, http.MethodPost, "/api/auth/register", "", map[string]string{"username": "sam", "password": "motdepasse"})
	rec := doJSON(t, h, http.MethodPost, "/api/auth/register", "", map[string]string{"username": "sam", "password": "motdepasse"})
	if rec.Code != http.StatusConflict {
		t.Fatalf("doublon status = %d", rec.Code)
	}
	rec = doJSON(t, h, http.MethodPost, "/api/auth/register", "", map[string]string{"username": "ab", "password": "motdepasse"})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("username court status = %d", rec.Code)
	}
}

func TestLoginWrongPassword(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	_ = doJSON(t, h, http.MethodPost, "/api/auth/register", "", map[string]string{"username": "sam", "password": "motdepasse"})
	rec := doJSON(t, h, http.MethodPost, "/api/auth/login", "", map[string]string{"username": "sam", "password": "mauvais"})
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d", rec.Code)
	}
}

func TestStaticIndex(t *testing.T) {
	s := newTestServer(t, true)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	s.Handler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Cetas Lite") {
		t.Fatal("index.html non servi")
	}
}
