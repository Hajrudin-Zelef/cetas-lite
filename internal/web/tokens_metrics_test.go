package web

import (
	"bufio"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"cetas-lite/internal/provider"
)

// Provider capturant la requête pour vérifier la propagation de MaxTokens.
type capturingProvider struct {
	content string
	mu      sync.Mutex
	got     []provider.Request
}

func (c *capturingProvider) ID() string { return "fake" }

func (c *capturingProvider) Stream(ctx context.Context, req provider.Request, emit func(provider.Event) bool) (provider.Response, error) {
	c.mu.Lock()
	c.got = append(c.got, req)
	c.mu.Unlock()
	if !emit(provider.Event{Content: c.content}) {
		return provider.Response{Content: c.content}, nil
	}
	return provider.Response{Content: c.content}, nil
}

func (c *capturingProvider) requests() []provider.Request {
	c.mu.Lock()
	defer c.mu.Unlock()
	return append([]provider.Request(nil), c.got...)
}

func waitForCall(t *testing.T, cp *capturingProvider) []provider.Request {
	t.Helper()
	deadline := time.Now().Add(8 * time.Second)
	for time.Now().Before(deadline) {
		if got := cp.requests(); len(got) > 0 {
			return got
		}
		time.Sleep(50 * time.Millisecond)
	}
	t.Fatal("le provider n'a reçu aucun appel")
	return nil
}

func TestChatSendMaxTokensPropagated(t *testing.T) {
	cp := &capturingProvider{content: "ok"}
	s := newTestServerWith(t, cp)
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)

	body, _ := json.Marshal(map[string]any{"family": "code", "mode": "standard", "message": "salut", "max_tokens": 1234})
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/chat/send", strings.NewReader(string(body)))
	req.Header.Set("Authorization", "Bearer "+tok)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("send status = %d", resp.StatusCode)
	}

	streamReq, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/chat/stream?from=0", nil)
	streamReq.Header.Set("Authorization", "Bearer "+tok)
	streamResp, err := http.DefaultClient.Do(streamReq)
	if err != nil {
		t.Fatal(err)
	}
	defer streamResp.Body.Close()
	reader := bufio.NewReader(streamResp.Body)
	deadline := time.Now().Add(8 * time.Second)
	for time.Now().Before(deadline) {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(line, "turn_done") {
			break
		}
	}
	got := waitForCall(t, cp)
	for _, r := range got {
		if r.MaxTokens != 1234 {
			t.Fatalf("MaxTokens = %d, attendu 1234", r.MaxTokens)
		}
	}
}

func TestChatSendMaxTokensDefaultsToSetting(t *testing.T) {
	cp := &capturingProvider{content: "ok"}
	s := newTestServerWith(t, cp)
	h := s.Handler()
	token := registerAndLogin(t, h, "mtuser3")

	// Reglage utilisateur à 7000, sans max_tokens dans le body → le reglage gagne.
	// (le fixture de test ne connait que la famille "code", on la renseigne aussi)
	rec := doJSON(t, h, http.MethodPut, "/api/settings", token, map[string]any{
		"max_tokens": 7000, "family": "code", "mode": "standard",
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("put settings = %d", rec.Code)
	}
	rec = doJSON(t, h, http.MethodPost, "/api/chat/send", token, map[string]any{
		"family": "code", "mode": "standard", "message": "salut",
	})
	if rec.Code != http.StatusOK {
		t.Fatalf("send = %d (%s)", rec.Code, rec.Body.String())
	}
	// Le turn tourne en arrière-plan : attendre que le provider soit appelé.
	got := waitForCall(t, cp)
	if got[0].MaxTokens != 7000 {
		t.Fatalf("MaxTokens = %d, attendu 7000 (reglage)", got[0].MaxTokens)
	}
}

func TestModelInfoEndpoint(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "miuser")

	rec := doJSON(t, h, http.MethodGet, "/api/model-info?provider=deepseek&model=deepseek-chat", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d (%s)", rec.Code, rec.Body.String())
	}
	body := decode(t, rec)
	if body["context_window"].(float64) != 131072 {
		t.Fatalf("context_window = %v", body["context_window"])
	}

	rec = doJSON(t, h, http.MethodGet, "/api/model-info?provider=inconnu&model=xyz", token, nil)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, attendu 404", rec.Code)
	}

	// Repli catalogue : modele absent de la table statique mais present au
	// catalogue (ex. routeur gratuit OpenRouter) -> 200, contexte inconnu.
	rec = doJSON(t, h, http.MethodGet, "/api/model-info?provider=openrouter&model=openrouter/free", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("openrouter/free : status = %d (%s), attendu 200", rec.Code, rec.Body.String())
	}
	body = decode(t, rec)
	if body["context_window"].(float64) != 0 {
		t.Fatalf("context_window = %v, attendu 0 (inconnu)", body["context_window"])
	}
}

func TestMetricsEndpoint(t *testing.T) {
	s := newTestServer(t, true)
	h := s.Handler()
	token := registerAndLogin(t, h, "metuser")

	rec := doJSON(t, h, http.MethodGet, "/api/metrics", token, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d (%s)", rec.Code, rec.Body.String())
	}
	body := decode(t, rec)
	if _, ok := body["cpu"]; !ok {
		t.Error("cpu manquant")
	}
	for _, k := range []string{"ram", "disk", "net"} {
		if _, ok := body[k].(map[string]any); !ok {
			t.Errorf("%s manquant ou invalide", k)
		}
	}
}
