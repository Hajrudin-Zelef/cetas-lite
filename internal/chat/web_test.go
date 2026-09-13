package chat

import (
	"context"
	"strings"
	"sync/atomic"
	"testing"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/search"
)

type stubWeb struct {
	result      search.Result
	fetchTitle  string
	fetchBody   string
	fetchErr    error
	searchCalls int32
}

func (s *stubWeb) Search(ctx context.Context, query string, maxResults int) search.Result {
	atomic.AddInt32(&s.searchCalls, 1)
	return s.result
}

func (s *stubWeb) Fetch(ctx context.Context, rawURL string) (string, string, error) {
	return s.fetchTitle, s.fetchBody, s.fetchErr
}

func hasToolName(tools []provider.Tool, name string) bool {
	for _, tl := range tools {
		if tl.Function.Name == name {
			return true
		}
	}
	return false
}

func TestAgentExposesWebToolsOnlyWhenEnabled(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.SetSearcher(&stubWeb{})

	runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "sans web"})
	reqs := sp.requests()
	if hasToolName(reqs[0].Tools, "web_search") {
		t.Fatal("les outils web ne doivent pas etre exposes sans le toggle")
	}

	sp2 := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e2 := newAgentEngine(t, sp2, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	e2.SetSearcher(&stubWeb{})
	runAgentTurn(t, e2, "sam", TurnInput{Family: "code", Mode: "standard", Text: "avec web", Web: true})
	reqs2 := sp2.requests()
	if !hasToolName(reqs2[0].Tools, "web_search") || !hasToolName(reqs2[0].Tools, "web_fetch") {
		t.Fatal("les outils web doivent etre exposes avec le toggle")
	}
}

func TestAgentWebSearchDispatch(t *testing.T) {
	stub := &stubWeb{result: search.Result{Provider: "tavily", Hits: []search.Hit{
		{Title: "Go", URL: "https://go.dev", Description: "langage"},
	}}}
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{
		{toolCalls: []provider.ToolCall{toolCall("c1", "web_search", `{"query":"golang"}`)}},
		{content: "voila"},
	}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.SetSearcher(stub)
	c := runAgentTurn(t, e, "sam", TurnInput{Family: "code", Mode: "standard", Text: "cherche", Web: true})

	if got := logText(c); !strings.Contains(got, "voila") {
		t.Fatalf("contenu = %q", got)
	}
	reqs := sp.requests()
	var toolResult string
	for _, m := range reqs[1].Messages {
		if m.Role == "tool" {
			toolResult, _ = m.Content.(string)
		}
	}
	if !strings.Contains(toolResult, "https://go.dev") {
		t.Fatalf("resultat d'outil web inattendu: %q", toolResult)
	}
	if atomic.LoadInt32(&stub.searchCalls) != 1 {
		t.Fatalf("appels de recherche = %d, want 1", stub.searchCalls)
	}
}

func TestNonAgentWebContextInjected(t *testing.T) {
	stub := &stubWeb{result: search.Result{Provider: "tavily", Hits: []search.Hit{
		{Title: "Go", URL: "https://go.dev", Description: "langage"},
	}}}
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	fams := []alias.Family{{ID: "plain", Label: "Plain", Modes: []alias.Mode{{
		ID: "standard", Pool: []alias.Member{{Provider: "fake", Model: "m"}},
	}}}}
	e := newAgentEngine(t, sp, fams)
	e.SetSearcher(stub)
	runAgentTurn(t, e, "sam", TurnInput{Family: "plain", Mode: "standard", Text: "golang ?", Web: true})

	reqs := sp.requests()
	found := false
	for _, m := range reqs[0].Messages {
		if m.Role == "system" {
			if s, _ := m.Content.(string); strings.Contains(s, "https://go.dev") {
				found = true
			}
		}
	}
	if !found {
		t.Fatal("le contexte web doit etre pre-injecte en mode non-agent")
	}
}

func TestWebRateLimit(t *testing.T) {
	sp := &scriptedProvider{id: "fake", steps: []scriptStep{{content: "ok"}}}
	e := newAgentEngine(t, sp, codeFamily(alias.Member{Provider: "fake", Model: "m"}))
	e.SetSearcher(&stubWeb{})
	for i := 0; i < 5; i++ {
		if !e.allowWeb("sam") {
			t.Fatalf("appel %d doit etre autorise", i+1)
		}
	}
	if e.allowWeb("sam") {
		t.Fatal("le 6e appel doit etre limite")
	}
}
