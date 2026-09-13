package web

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/auth"
	"cetas-lite/internal/chat"
	"cetas-lite/internal/config"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/store"
)

type fakeProvider struct{ content string }

func (f *fakeProvider) ID() string { return "fake" }

func (f *fakeProvider) Stream(ctx context.Context, req provider.Request, emit func(provider.Event) bool) (provider.Response, error) {
	for _, r := range f.content {
		if !emit(provider.Event{Content: string(r)}) {
			break
		}
	}
	return provider.Response{Content: f.content}, nil
}

func newTestServerWith(t *testing.T, fp provider.Provider) *Server {
	t.Helper()
	t.Setenv("CETAS_LITE_HOME", t.TempDir())
	t.Setenv("CETAS_LITE_REGISTRATION_OPEN", "true")
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
	reg := provider.NewRegistry()
	if fp != nil {
		reg.Set(fp)
	}
	fams := []alias.Family{{
		ID: "code", Label: "Code",
		Modes: []alias.Mode{{ID: "standard", Agent: true, Pool: []alias.Member{{Provider: "fake", Model: "ok"}}}},
	}}
	engine := chat.NewEngine(reg, fams, st, nil, t.TempDir())
	return New(cfg, st, m, engine, "test")
}

func tokenFor(t *testing.T, base string) string {
	t.Helper()
	body, _ := json.Marshal(map[string]string{"username": "sam", "password": "motdepasse"})
	resp, err := http.Post(base+"/api/auth/register", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	resp, err = http.Post(base+"/api/auth/login", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	var out map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		t.Fatal(err)
	}
	tok, _ := out["token"].(string)
	if tok == "" {
		t.Fatal("token vide")
	}
	return tok
}

func TestAliasesEndpoint(t *testing.T) {
	s := newTestServerWith(t, &fakeProvider{content: "x"})
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/aliases", nil)
	req.Header.Set("Authorization", "Bearer "+tok)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d", resp.StatusCode)
	}
	var out struct {
		Families []alias.ResolvedFamily `json:"families"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		t.Fatal(err)
	}
	if len(out.Families) != 1 || out.Families[0].ID != "code" {
		t.Fatalf("families = %+v", out.Families)
	}
	if !out.Families[0].Modes[0].Agent {
		t.Fatal("code doit activer l'agent")
	}
}

func TestAliasesPutOverride(t *testing.T) {
	s := newTestServerWith(t, &fakeProvider{content: "x"})
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)

	body := `{"code":{"standard":[{"provider":"openrouter","model":"anthropic/claude-sonnet-4.5"}]}}`
	req, _ := http.NewRequest(http.MethodPut, ts.URL+"/api/aliases", strings.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+tok)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d", resp.StatusCode)
	}
	var out struct {
		Families []alias.ResolvedFamily `json:"families"`
	}
	_ = json.NewDecoder(resp.Body).Decode(&out)
	if out.Families[0].Modes[0].Pool[0].Model != "anthropic/claude-sonnet-4.5" {
		t.Fatalf("override non applique: %+v", out.Families[0].Modes[0].Pool)
	}
}

func TestChatSendAndStream(t *testing.T) {
	s := newTestServerWith(t, &fakeProvider{content: "bonjour"})
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)

	sendBody, _ := json.Marshal(map[string]string{"family": "code", "mode": "standard", "message": "salut"})
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/chat/send", bytes.NewReader(sendBody))
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
	gotContent, gotDone, gotRoute := false, false, false
	deadline := time.Now().Add(3 * time.Second)
	for !gotDone && time.Now().Before(deadline) {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(line, `"content":"bonjour"`) {
			gotContent = true
		}
		if strings.Contains(line, `"route"`) {
			gotRoute = true
		}
		if strings.Contains(line, "turn_done") {
			gotDone = true
		}
	}
	if !gotRoute {
		t.Error("evenement route manquant")
	}
	if !gotContent {
		t.Error("contenu stream manquant")
	}
	if !gotDone {
		t.Error("turn_done manquant")
	}
}

func TestChatBusyConflict(t *testing.T) {
	s := newTestServerWith(t, &fakeProvider{content: "lent"})
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)

	// Premier envoi puis reset pour ne pas laisser de tour bloque.
	sendBody, _ := json.Marshal(map[string]string{"family": "code", "mode": "standard", "message": "salut"})
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/chat/send", bytes.NewReader(sendBody))
	req.Header.Set("Authorization", "Bearer "+tok)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(req)
	resp.Body.Close()

	// attendre la fin du tour
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		stReq, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/chat/state", nil)
		stReq.Header.Set("Authorization", "Bearer "+tok)
		stResp, _ := http.DefaultClient.Do(stReq)
		var st map[string]any
		_ = json.NewDecoder(stResp.Body).Decode(&st)
		stResp.Body.Close()
		if st["generating"] == false {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
}
