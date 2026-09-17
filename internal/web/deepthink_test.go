package web

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"cetas-lite/internal/provider"
)

// mockDeepThinkProvider : provider factice pour tester translateText.
type mockDeepThinkProvider struct {
	id      string
	content string
	err     error
	calls   int
	lastReq provider.Request
}

func (m *mockDeepThinkProvider) ID() string { return m.id }

func (m *mockDeepThinkProvider) Stream(ctx context.Context, req provider.Request, emit func(provider.Event) bool) (provider.Response, error) {
	m.calls++
	m.lastReq = req
	if m.err != nil {
		return provider.Response{}, m.err
	}
	return provider.Response{Content: m.content}, nil
}

func deepThinkToken(t *testing.T, h http.Handler) string {
	t.Helper()
	rec := doJSON(t, h, http.MethodPost, "/api/auth/register", "", map[string]string{
		"username": "sam", "password": "motdepasse123",
	})
	if rec.Code != http.StatusCreated {
		t.Fatalf("register status = %d (%s)", rec.Code, rec.Body.String())
	}
	rec = doJSON(t, h, http.MethodPost, "/api/auth/login", "", map[string]string{
		"username": "sam", "password": "motdepasse123",
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("login status = %d (%s)", rec.Code, rec.Body.String())
	}
	token, _ := decode(t, rec)["token"].(string)
	if token == "" {
		t.Fatal("token manquant")
	}
	return token
}

func TestDeepThinkDefaults(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := deepThinkToken(t, h)

	rec := doJSON(t, h, http.MethodGet, "/api/deepthink", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d (%s)", rec.Code, rec.Body.String())
	}
	body := decode(t, rec)
	st, _ := body["settings"].(map[string]any)
	if st["lang"] != "fr" {
		t.Fatalf("langue par defaut = %v, want fr", st["lang"])
	}
	if st["provider"] != "openrouter" || st["model"] != "openrouter/free" {
		t.Fatalf("modele par defaut = %v/%v", st["provider"], st["model"])
	}
	if langs, _ := body["langs"].([]any); len(langs) != 5 {
		t.Fatalf("langs = %v, want 5 langues", langs)
	}
	if models, _ := body["models"].([]any); len(models) != 2 {
		t.Fatalf("models = %v, want 2 modeles", models)
	}
}

func TestDeepThinkPutRoundTrip(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := deepThinkToken(t, h)

	rec := doJSON(t, h, http.MethodPut, "/api/deepthink", token, map[string]string{
		"lang": "es", "provider": "deepseek", "model": "deepseek-chat",
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("PUT status = %d (%s)", rec.Code, rec.Body.String())
	}
	rec = doJSON(t, h, http.MethodGet, "/api/deepthink", token, nil)
	st, _ := decode(t, rec)["settings"].(map[string]any)
	if st["lang"] != "es" || st["provider"] != "deepseek" || st["model"] != "deepseek-chat" {
		t.Fatalf("settings persistés = %v", st)
	}
}

func TestDeepThinkPutRejects(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := deepThinkToken(t, h)

	cases := []map[string]string{
		{"lang": "xx", "provider": "openrouter", "model": "openrouter/free"},
		{"lang": "fr", "provider": "openrouter", "model": "nope/nope"},
		{"lang": "fr", "provider": "nope", "model": "openrouter/free"},
		{"lang": "", "provider": "openrouter", "model": "openrouter/free"},
	}
	for i, body := range cases {
		rec := doJSON(t, h, http.MethodPut, "/api/deepthink", token, body)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("cas %d : status = %d, want 400", i, rec.Code)
		}
	}
}

func TestDeepThinkTranslateNoProvider(t *testing.T) {
	s := newTestServer(t, true) // registre vide : aucun provider configuré
	h := s.Handler()
	token := deepThinkToken(t, h)

	rec := doJSON(t, h, http.MethodPost, "/api/deepthink/translate", token, map[string]string{
		"text": "The user says hello.",
	})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
	if got := decode(t, rec)["error"]; got == nil || !containsStr(strOf(got), "clé API") {
		t.Fatalf("message d'erreur peu clair : %v", got)
	}
}

func TestDeepThinkTranslateEmpty(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := deepThinkToken(t, h)

	rec := doJSON(t, h, http.MethodPost, "/api/deepthink/translate", token, map[string]string{
		"text": "   ",
	})
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}

func TestTranslateTextCache(t *testing.T) {
	mock := &mockDeepThinkProvider{id: "openrouter", content: "L'utilisateur dit bonjour."}
	st := DeepThinkSettings{Lang: "fr", Provider: "openrouter", Model: "openrouter/free"}
	ctx := context.Background()

	out1, err := translateText(ctx, mock, st, "The user says hello.")
	if err != nil {
		t.Fatalf("traduction 1 : %v", err)
	}
	out2, err := translateText(ctx, mock, st, "The user says hello.")
	if err != nil {
		t.Fatalf("traduction 2 : %v", err)
	}
	if out1 != "L'utilisateur dit bonjour." || out2 != out1 {
		t.Fatalf("traductions = %q / %q", out1, out2)
	}
	if mock.calls != 1 {
		t.Fatalf("provider appelé %d fois, want 1 (cache)", mock.calls)
	}
	// Le prompt de traduction est minimal et deterministe.
	msgs := mock.lastReq.Messages
	if len(msgs) != 2 {
		t.Fatalf("messages = %d, want 2", len(msgs))
	}
	if sys, _ := msgs[0].Content.(string); !containsStr(sys, "French") || !containsStr(sys, "only the translation") {
		t.Fatalf("system prompt = %q", sys)
	}
	if mock.lastReq.Temperature != 0.2 {
		t.Fatalf("temperature = %v, want 0.2", mock.lastReq.Temperature)
	}
}

func TestTranslateTextEmpty(t *testing.T) {
	mock := &mockDeepThinkProvider{id: "openrouter", content: "   "}
	st := DeepThinkSettings{Lang: "fr", Provider: "openrouter", Model: "openrouter/free"}
	if _, err := translateText(context.Background(), mock, st, "Hello."); err == nil {
		t.Fatal("traduction vide : erreur attendue")
	}
}

func TestTranslateTextError(t *testing.T) {
	mock := &mockDeepThinkProvider{id: "openrouter", err: errors.New("boom")}
	st := DeepThinkSettings{Lang: "fr", Provider: "openrouter", Model: "openrouter/free"}
	if _, err := translateText(context.Background(), mock, st, "Hello."); err == nil {
		t.Fatal("erreur provider : erreur attendue")
	}
}

func TestSanitizeDeepThinkSettings(t *testing.T) {
	def := defaultDeepThinkSettings()
	if got := sanitizeDeepThinkSettings(DeepThinkSettings{}); got != def {
		t.Fatalf("vide -> %v, want %v", got, def)
	}
	st := sanitizeDeepThinkSettings(DeepThinkSettings{Lang: "it", Provider: "deepseek", Model: "deepseek-chat"})
	if st.Lang != "it" || st.Provider != "deepseek" {
		t.Fatalf("valide modifié : %v", st)
	}
	st = sanitizeDeepThinkSettings(DeepThinkSettings{Lang: "xx", Provider: "nope", Model: "x"})
	if st != def {
		t.Fatalf("invalide -> %v, want %v", st, def)
	}
}

func containsStr(s, sub string) bool {
	return len(s) >= len(sub) && (func() bool {
		for i := 0; i+len(sub) <= len(s); i++ {
			if s[i:i+len(sub)] == sub {
				return true
			}
		}
		return false
	})()
}

func strOf(v any) string {
	s, _ := v.(string)
	return s
}
