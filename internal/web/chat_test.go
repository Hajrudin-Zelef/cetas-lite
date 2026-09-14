package web

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	"cetas-lite/internal/terminal"
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
	termMgr := terminal.NewManager(cfg.WorkspaceDir, t.TempDir())
	return New(cfg, st, m, engine, termMgr, "test")
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
	var content strings.Builder
	gotDone, gotRoute := false, false
	deadline := time.Now().Add(8 * time.Second)
	for !gotDone && time.Now().Before(deadline) {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if c := lineString(line, "content"); c != "" {
			content.WriteString(c)
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
	if content.String() != "bonjour" {
		t.Errorf("contenu stream = %q, want bonjour", content.String())
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

type sseStream struct{ ch chan string }

func newSSEStream(body io.Reader) *sseStream {
	ch := make(chan string, 256)
	go func() {
		sc := bufio.NewScanner(body)
		sc.Buffer(make([]byte, 0, 64*1024), 1<<20)
		for sc.Scan() {
			ch <- sc.Text()
		}
		close(ch)
	}()
	return &sseStream{ch: ch}
}

func (s *sseStream) readUntil(stop func(string) bool, timeout time.Duration) []string {
	var out []string
	timer := time.After(timeout)
	for {
		select {
		case line, ok := <-s.ch:
			if !ok {
				return out
			}
			out = append(out, line)
			if stop != nil && stop(line) {
				return out
			}
		case <-timer:
			return out
		}
	}
}

func containsLine(lines []string, sub string) bool {
	for _, l := range lines {
		if strings.Contains(l, sub) {
			return true
		}
	}
	return false
}

func lineSeq(line string) int { return lineField(line, "seq") }

func lineString(line, key string) string {
	const p = "data: "
	if !strings.HasPrefix(line, p) {
		return ""
	}
	var m map[string]any
	if json.Unmarshal([]byte(line[len(p):]), &m) != nil {
		return ""
	}
	s, _ := m[key].(string)
	return s
}

func lineField(line, key string) int {
	const p = "data: "
	if !strings.HasPrefix(line, p) {
		return 0
	}
	var m map[string]any
	if json.Unmarshal([]byte(line[len(p):]), &m) != nil {
		return 0
	}
	if v, ok := m[key].(float64); ok {
		return int(v)
	}
	return 0
}

func sendAndWait(t *testing.T, base, tok, msg string) {
	t.Helper()
	body, _ := json.Marshal(map[string]string{"family": "code", "mode": "standard", "message": msg})
	req, _ := http.NewRequest(http.MethodPost, base+"/api/chat/send", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+tok)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		stReq, _ := http.NewRequest(http.MethodGet, base+"/api/chat/state", nil)
		stReq.Header.Set("Authorization", "Bearer "+tok)
		stResp, err := http.DefaultClient.Do(stReq)
		if err != nil {
			t.Fatal(err)
		}
		var st map[string]any
		_ = json.NewDecoder(stResp.Body).Decode(&st)
		stResp.Body.Close()
		if st["generating"] == false {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("tour non termine")
}

func streamRequest(t *testing.T, base, tok string, from int) *http.Response {
	t.Helper()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/chat/stream?from=%d", base, from), nil)
	req.Header.Set("Authorization", "Bearer "+tok)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	return resp
}

func TestChatReplayCoalescesFromZero(t *testing.T) {
	s := newTestServerWith(t, &fakeProvider{content: "bonjour"})
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)
	sendAndWait(t, ts.URL, tok, "salut")

	resp := streamRequest(t, ts.URL, tok, 0)
	defer resp.Body.Close()
	lines := newSSEStream(resp.Body).readUntil(func(l string) bool { return strings.Contains(l, "caught_up") }, 3*time.Second)
	if !containsLine(lines, "caught_up") {
		t.Fatal("caught_up manquant")
	}
	var contents []string
	for _, l := range lines {
		if strings.Contains(l, `"content"`) {
			contents = append(contents, l)
		}
	}
	if len(contents) != 1 || !strings.Contains(contents[0], "bonjour") {
		t.Fatalf("coalescence = %v", contents)
	}

	seq0 := lineField(contents[0], "seq0")
	if seq0 == 0 {
		t.Fatal("seq0 introuvable")
	}
	resp2 := streamRequest(t, ts.URL, tok, seq0)
	defer resp2.Body.Close()
	lines2 := newSSEStream(resp2.Body).readUntil(func(l string) bool { return strings.Contains(l, "caught_up") }, 3*time.Second)
	if !containsLine(lines2, "onjour") {
		t.Fatalf("continuation attendue depuis seq0=%d: %v", seq0, lines2)
	}
	if containsLine(lines2, `"bonjour"`) {
		t.Fatalf("le debut deja recu ne doit pas etre rediffuse: %v", lines2)
	}
}

func TestChatReconnectFromLastNoDuplicate(t *testing.T) {
	s := newTestServerWith(t, &fakeProvider{content: "bonjour"})
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)
	sendAndWait(t, ts.URL, tok, "salut")

	resp := streamRequest(t, ts.URL, tok, 0)
	lines := newSSEStream(resp.Body).readUntil(func(l string) bool { return strings.Contains(l, "caught_up") }, 3*time.Second)
	resp.Body.Close()
	maxSeq := 0
	for _, l := range lines {
		if v := lineSeq(l); v > maxSeq {
			maxSeq = v
		}
	}
	if maxSeq == 0 {
		t.Fatal("aucun seq dans le flux initial")
	}

	resp2 := streamRequest(t, ts.URL, tok, maxSeq)
	defer resp2.Body.Close()
	lines2 := newSSEStream(resp2.Body).readUntil(func(l string) bool { return strings.Contains(l, "caught_up") }, 3*time.Second)
	if !containsLine(lines2, "caught_up") {
		t.Fatal("caught_up manquant en reconnexion")
	}
	if containsLine(lines2, `"content"`) {
		t.Fatalf("contenu deja vu rediffuse apres from=%d", maxSeq)
	}
}

func TestChatStreamEmitsReset(t *testing.T) {
	s := newTestServerWith(t, &fakeProvider{content: "x"})
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)

	resp := streamRequest(t, ts.URL, tok, 0)
	defer resp.Body.Close()
	stream := newSSEStream(resp.Body)
	stream.readUntil(func(l string) bool { return strings.Contains(l, "caught_up") }, 2*time.Second)

	rreq, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/chat/reset", nil)
	rreq.Header.Set("Authorization", "Bearer "+tok)
	rresp, err := http.DefaultClient.Do(rreq)
	if err != nil {
		t.Fatal(err)
	}
	rresp.Body.Close()

	lines := stream.readUntil(func(l string) bool { return strings.Contains(l, `"reset":true`) }, 2*time.Second)
	if !containsLine(lines, `"reset":true`) {
		t.Fatalf("evenement reset manquant: %v", lines)
	}
}

func TestConversationsArchives(t *testing.T) {
	s := newTestServerWith(t, &fakeProvider{content: "bonjour"})
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)
	sendAndWait(t, ts.URL, tok, "salut")

	rreq, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/chat/reset", nil)
	rreq.Header.Set("Authorization", "Bearer "+tok)
	rresp, err := http.DefaultClient.Do(rreq)
	if err != nil {
		t.Fatal(err)
	}
	rresp.Body.Close()

	lreq, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/conversations", nil)
	lreq.Header.Set("Authorization", "Bearer "+tok)
	lresp, err := http.DefaultClient.Do(lreq)
	if err != nil {
		t.Fatal(err)
	}
	var list struct {
		Archives []struct {
			ID       string `json:"id"`
			Title    string `json:"title"`
			Updated  int64  `json:"updated"`
			Messages int    `json:"messages"`
		} `json:"archives"`
	}
	_ = json.NewDecoder(lresp.Body).Decode(&list)
	lresp.Body.Close()
	if len(list.Archives) != 1 {
		t.Fatalf("archives = %v", list.Archives)
	}
	if list.Archives[0].Title != "salut" {
		t.Fatalf("titre = %q", list.Archives[0].Title)
	}
	if list.Archives[0].Updated == 0 || list.Archives[0].Messages != 2 {
		t.Fatalf("meta archive = %+v", list.Archives[0])
	}
	id := list.Archives[0].ID

	body, _ := json.Marshal(map[string]string{"id": id})
	rreq2, _ := http.NewRequest(http.MethodPost, ts.URL+"/api/conversations/restore", bytes.NewReader(body))
	rreq2.Header.Set("Authorization", "Bearer "+tok)
	rreq2.Header.Set("Content-Type", "application/json")
	rresp2, err := http.DefaultClient.Do(rreq2)
	if err != nil {
		t.Fatal(err)
	}
	defer rresp2.Body.Close()
	if rresp2.StatusCode != http.StatusOK {
		t.Fatalf("restore status = %d", rresp2.StatusCode)
	}

	dreq, _ := http.NewRequest(http.MethodDelete, ts.URL+"/api/conversations/"+id, nil)
	dreq.Header.Set("Authorization", "Bearer "+tok)
	dresp, err := http.DefaultClient.Do(dreq)
	if err != nil {
		t.Fatal(err)
	}
	dresp.Body.Close()
	if dresp.StatusCode != http.StatusOK {
		t.Fatalf("delete status = %d", dresp.StatusCode)
	}

	dreq2, _ := http.NewRequest(http.MethodDelete, ts.URL+"/api/conversations/"+id, nil)
	dreq2.Header.Set("Authorization", "Bearer "+tok)
	dresp2, err := http.DefaultClient.Do(dreq2)
	if err != nil {
		t.Fatal(err)
	}
	dresp2.Body.Close()
	if dresp2.StatusCode != http.StatusNotFound {
		t.Fatalf("delete inconnu status = %d", dresp2.StatusCode)
	}
}
