package chat

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

// ghMockServer monte un faux api.github.com et retourne l'URL de base,
// un compteur de requêtes et les en-têtes/payloads observés.
type ghMock struct {
	t        *testing.T
	hits     atomic.Int64
	lastAuth string
	lastBody []byte
	lastPath string
	lastMeth string
	handler  func(w http.ResponseWriter, r *http.Request)
}

func newGitHubMock(t *testing.T, handler func(w http.ResponseWriter, r *http.Request)) (*ghMock, *Sandbox, func()) {
	t.Helper()
	m := &ghMock{t: t, handler: handler}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.hits.Add(1)
		m.lastAuth = r.Header.Get("Authorization")
		m.lastPath = r.URL.Path
		m.lastMeth = r.Method
		b, _ := io.ReadAll(r.Body)
		m.lastBody = b
		m.handler(w, r)
	}))
	old := githubAPIBase
	githubAPIBase = srv.URL
	sb := newTestSandbox(t)
	sb.GitHubToken = func() string { return "test-token" }
	return m, sb, func() {
		githubAPIBase = old
		srv.Close()
	}
}

func ghJSONResp(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func TestGitHubRepos(t *testing.T) {
	m, sb, done := newGitHubMock(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/user/repos" {
			http.NotFound(w, r)
			return
		}
		ghJSONResp(w, 200, []map[string]any{{
			"name": "demo", "full_name": "octo/demo", "private": true,
			"description": "d", "html_url": "https://github.com/octo/demo",
			"default_branch": "main", "updated_at": "2026-09-15T00:00:00Z",
			"stargazers_count": 42, "secret_field": "doit-etre-filtre",
		}})
	})
	defer done()
	res := sb.Execute(context.Background(), "GitHubRepos", `{"limit":5}`)
	if strings.HasPrefix(res.Text, "[error]") {
		t.Fatalf("repos: %s", res.Text)
	}
	if m.lastAuth != "Bearer test-token" {
		t.Fatalf("auth = %q", m.lastAuth)
	}
	if !strings.Contains(res.Text, `"full_name":"octo/demo"`) {
		t.Fatalf("projection manquante: %s", res.Text)
	}
	if strings.Contains(res.Text, "secret_field") || strings.Contains(res.Text, "stargazers_count") {
		t.Fatalf("projection trop large: %s", res.Text)
	}
}

func TestGitHubIssueCreate(t *testing.T) {
	m, sb, done := newGitHubMock(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/repos/octo/demo/issues" {
			http.NotFound(w, r)
			return
		}
		ghJSONResp(w, 201, map[string]any{"number": 7, "html_url": "https://github.com/octo/demo/issues/7"})
	})
	defer done()
	res := sb.Execute(context.Background(), "GitHubIssueCreate",
		`{"owner":"octo","repo":"demo","title":"Bug","body":"desc","labels":["bug","urgent"]}`)
	if strings.HasPrefix(res.Text, "[error]") {
		t.Fatalf("create: %s", res.Text)
	}
	if !strings.Contains(res.Text, "#7") || !strings.Contains(res.Text, "issues/7") {
		t.Fatalf("confirmation inattendue: %s", res.Text)
	}
	var payload map[string]any
	if err := json.Unmarshal(m.lastBody, &payload); err != nil {
		t.Fatalf("payload illisible: %v", err)
	}
	if payload["title"] != "Bug" {
		t.Fatalf("title = %v", payload["title"])
	}
	labels, _ := payload["labels"].([]any)
	if len(labels) != 2 {
		t.Fatalf("labels = %v", payload["labels"])
	}
}

func TestGitHubPRMerge(t *testing.T) {
	m, sb, done := newGitHubMock(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" || r.URL.Path != "/repos/octo/demo/pulls/12/merge" {
			http.NotFound(w, r)
			return
		}
		ghJSONResp(w, 200, map[string]any{"merged": true, "message": "ok"})
	})
	defer done()
	res := sb.Execute(context.Background(), "GitHubPRMerge",
		`{"owner":"octo","repo":"demo","number":12,"merge_method":"squash"}`)
	if strings.HasPrefix(res.Text, "[error]") {
		t.Fatalf("merge: %s", res.Text)
	}
	if !strings.Contains(res.Text, "#12") || !strings.Contains(res.Text, "squash") {
		t.Fatalf("confirmation inattendue: %s", res.Text)
	}
	if !strings.Contains(string(m.lastBody), `"merge_method":"squash"`) {
		t.Fatalf("merge_method non transmis: %s", m.lastBody)
	}
}

func TestGitHubNoToken(t *testing.T) {
	sb := newTestSandbox(t) // pas de GitHubToken
	res := sb.Execute(context.Background(), "GitHubRepos", `{}`)
	if !strings.Contains(res.Text, "not connected") || !strings.Contains(res.Text, "Connectors") {
		t.Fatalf("message de guidage attendu: %s", res.Text)
	}
	res = sb.Execute(context.Background(), "GitHubIssueCreate", `{"owner":"o","repo":"r","title":"t"}`)
	if !strings.Contains(res.Text, "not connected") {
		t.Fatalf("l'écriture doit aussi guider: %s", res.Text)
	}
}

func TestGitHubValidation(t *testing.T) {
	m, sb, done := newGitHubMock(t, func(w http.ResponseWriter, r *http.Request) {
		ghJSONResp(w, 200, []any{})
	})
	defer done()
	// owner avec slash : rejeté avant tout appel HTTP.
	res := sb.Execute(context.Background(), "GitHubIssues", `{"owner":"octo/evil","repo":"demo"}`)
	if !strings.Contains(res.Text, "invalid") {
		t.Fatalf("validation owner attendue: %s", res.Text)
	}
	// number négatif.
	res = sb.Execute(context.Background(), "GitHubIssueGet", `{"owner":"octo","repo":"demo","number":-1}`)
	if !strings.Contains(res.Text, "positive") {
		t.Fatalf("validation number attendue: %s", res.Text)
	}
	// merge_method inconnu.
	res = sb.Execute(context.Background(), "GitHubPRMerge", `{"owner":"octo","repo":"demo","number":1,"merge_method":"yolo"}`)
	if !strings.Contains(res.Text, "merge_method") {
		t.Fatalf("validation merge_method attendue: %s", res.Text)
	}
	// state inconnu.
	res = sb.Execute(context.Background(), "GitHubPRs", `{"owner":"octo","repo":"demo","state":"bogus"}`)
	if !strings.Contains(res.Text, "state") {
		t.Fatalf("validation state attendue: %s", res.Text)
	}
	if m.hits.Load() != 0 {
		t.Fatalf("aucun appel HTTP ne devait partir, hits=%d", m.hits.Load())
	}
}

func TestGitHubAPIErrors(t *testing.T) {
	_, sb, done := newGitHubMock(t, func(w http.ResponseWriter, r *http.Request) {
		ghJSONResp(w, 404, map[string]any{"message": "Not Found"})
	})
	defer done()
	res := sb.Execute(context.Background(), "GitHubIssues", `{"owner":"octo","repo":"nope"}`)
	if !strings.Contains(res.Text, "not found") || !strings.Contains(res.Text, "404") {
		t.Fatalf("erreur 404 actionnable attendue: %s", res.Text)
	}
}

func TestGitHub401(t *testing.T) {
	_, sb, done := newGitHubMock(t, func(w http.ResponseWriter, r *http.Request) {
		ghJSONResp(w, 401, map[string]any{"message": "Bad credentials"})
	})
	defer done()
	res := sb.Execute(context.Background(), "GitHubRepos", `{}`)
	if !strings.Contains(res.Text, "401") || !strings.Contains(res.Text, "Reconnect") {
		t.Fatalf("erreur 401 actionnable attendue: %s", res.Text)
	}
}

func TestGitHubIssueGetWithComments(t *testing.T) {
	_, sb, done := newGitHubMock(t, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasSuffix(r.URL.Path, "/comments"):
			ghJSONResp(w, 200, []map[string]any{{
				"user": map[string]any{"login": "octo"}, "created_at": "2026-09-15T01:00:00Z",
				"body": "d'accord",
			}})
		default:
			ghJSONResp(w, 200, map[string]any{
				"number": 3, "title": "T", "state": "open",
				"html_url": "https://github.com/octo/demo/issues/3",
				"user":     map[string]any{"login": "octo"},
				"labels":   []any{}, "comments": 1, "created_at": "2026-09-14T00:00:00Z",
				"body": "corps",
			})
		}
	})
	defer done()
	res := sb.Execute(context.Background(), "GitHubIssueGet", `{"owner":"octo","repo":"demo","number":3}`)
	if strings.HasPrefix(res.Text, "[error]") {
		t.Fatalf("get: %s", res.Text)
	}
	if !strings.Contains(res.Text, "fetched_comments") || !strings.Contains(res.Text, "d'accord") {
		t.Fatalf("commentaires attendus: %s", res.Text)
	}
}

func TestGitHubApprovalPolicy(t *testing.T) {
	writes := []string{"GitHubRepoCreate", "GitHubIssueCreate", "GitHubIssueComment", "GitHubPRCreate", "GitHubPRMerge"}
	for _, n := range writes {
		if !needsApproval(n) {
			t.Fatalf("%s devrait exiger une approbation", n)
		}
		if planToolAllowed(n) {
			t.Fatalf("%s ne devrait pas être autorisé en mode plan", n)
		}
	}
	reads := []string{"GitHubRepos", "GitHubIssues", "GitHubIssueGet", "GitHubPRs"}
	for _, n := range reads {
		if needsApproval(n) {
			t.Fatalf("%s ne devrait pas exiger d'approbation", n)
		}
		if !planToolAllowed(n) {
			t.Fatalf("%s devrait être autorisé en mode plan", n)
		}
	}
}

func TestGitHubSchemasRegistered(t *testing.T) {
	names := map[string]bool{}
	for _, tl := range ToolSchemas() {
		names[tl.Function.Name] = true
	}
	for _, n := range []string{"GitHubRepos", "GitHubIssues", "GitHubIssueGet", "GitHubPRs",
		"GitHubRepoCreate", "GitHubIssueCreate", "GitHubIssueComment", "GitHubPRCreate", "GitHubPRMerge"} {
		if !names[n] {
			t.Fatalf("schéma manquant: %s", n)
		}
	}
}
