package provider

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func sseServer(t *testing.T, body string, status int, cut bool) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		if cut {
			hj, ok := w.(http.Hijacker)
			if !ok {
				t.Error("serveur non detournable")
				return
			}
			conn, _, err := hj.Hijack()
			if err != nil {
				t.Error(err)
				return
			}
			_ = conn.Close()
		}
	}))
	t.Cleanup(srv.Close)
	return srv
}

func providerFor(t *testing.T, srv *httptest.Server, path string) *OpenAICompat {
	t.Helper()
	return NewOpenAICompat(Endpoint{
		Provider: "test", BaseURL: srv.URL, Path: path, APIKey: "k",
	}, srv.Client())
}

func TestStreamContentUsageTools(t *testing.T) {
	body := strings.Join([]string{
		`data: {"choices":[{"delta":{"reasoning_content":"je reflechis"}}]}`,
		`data: {"choices":[{"delta":{"content":"Bon"}}]}`,
		`data: {"choices":[{"delta":{"content":"jour"}}]}`,
		`data: {"choices":[{"delta":{"tool_calls":[{"index":0,"id":"c1","type":"function","function":{"name":"Read","arguments":"{\"file_"}}]}}]}`,
		`data: {"choices":[{"delta":{"tool_calls":[{"index":0,"function":{"arguments":"path\":\"a.txt\"}"}}]}}]}`,
		`data: {"choices":[{"delta":{},"finish_reason":"tool_calls"}]}`,
		`data: {"choices":[],"usage":{"prompt_tokens":10,"completion_tokens":4,"total_tokens":14}}`,
		`data: [DONE]`,
		"",
	}, "\n\n")
	srv := sseServer(t, body, 200, false)
	p := providerFor(t, srv, "/chat/completions")

	var content, reasoning strings.Builder
	var usage *Usage
	resp, err := p.Stream(context.Background(), Request{Model: "m", Messages: []Message{{Role: "user", Content: "salut"}}}, func(ev Event) bool {
		content.WriteString(ev.Content)
		reasoning.WriteString(ev.Reasoning)
		if ev.Usage != nil {
			usage = ev.Usage
		}
		return true
	})
	if err != nil {
		t.Fatalf("Stream: %v", err)
	}
	if content.String() != "Bonjour" {
		t.Fatalf("content = %q", content.String())
	}
	if reasoning.String() != "je reflechis" {
		t.Fatalf("reasoning = %q", reasoning.String())
	}
	if resp.Content != "Bonjour" || resp.Reasoning != "je reflechis" {
		t.Fatalf("resp = %+v", resp)
	}
	if len(resp.ToolCalls) != 1 || resp.ToolCalls[0].Function.Name != "Read" ||
		resp.ToolCalls[0].Function.Arguments != `{"file_path":"a.txt"}` {
		t.Fatalf("tool calls = %+v", resp.ToolCalls)
	}
	if resp.FinishReason != "tool_calls" {
		t.Fatalf("finish = %q", resp.FinishReason)
	}
	if resp.Usage.PromptTokens != 10 || resp.Usage.TotalTokens != 14 {
		t.Fatalf("usage = %+v", resp.Usage)
	}
	if usage == nil {
		t.Fatal("usage non emis")
	}
}

func TestStreamNon200(t *testing.T) {
	srv := sseServer(t, "contexte depasse", http.StatusInternalServerError, false)
	p := providerFor(t, srv, "/chat/completions")
	_, err := p.Stream(context.Background(), Request{Model: "m"}, func(Event) bool { return true })
	if err == nil || !strings.Contains(err.Error(), "500") {
		t.Fatalf("erreur attendue avec le statut, got %v", err)
	}
}

func TestStreamCutRemonteUneErreur(t *testing.T) {
	srv := sseServer(t, `data: {"choices":[{"delta":{"content":"debut"}}]}`+"\n\n", 200, true)
	p := providerFor(t, srv, "/chat/completions")
	var got strings.Builder
	_, err := p.Stream(context.Background(), Request{Model: "m"}, func(ev Event) bool {
		got.WriteString(ev.Content)
		return true
	})
	if err == nil {
		t.Fatal("flux coupe: aucune erreur remontee")
	}
	if !errors.Is(err, ErrStreamCut) {
		t.Fatalf("erreur = %v, want ErrStreamCut", err)
	}
	if got.String() != "debut" {
		t.Fatalf("texte recu avant coupure perdu: %q", got.String())
	}
}

func TestStreamCancelSilencieux(t *testing.T) {
	srv := sseServer(t, `data: {"choices":[{"delta":{"content":"a"}}]}`+"\n\n", 200, true)
	p := providerFor(t, srv, "/chat/completions")
	ctx, cancel := context.WithCancel(context.Background())
	_, err := p.Stream(ctx, Request{Model: "m"}, func(ev Event) bool {
		if ev.Content != "" {
			cancel()
		}
		return true
	})
	if err == nil {
		t.Fatal("annulation attendue")
	}
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("stop volontaire doit etre context.Canceled, got %v", err)
	}
}

func TestCloudEndpoint(t *testing.T) {
	ep, ok := CloudEndpoint("deepseek", "sk")
	if !ok || ep.BaseURL != "https://api.deepseek.com" || ep.Path != "/chat/completions" {
		t.Fatalf("deepseek endpoint = %+v", ep)
	}
	or, _ := CloudEndpoint("openrouter", "sk")
	if or.Headers["X-Title"] == "" {
		t.Fatal("openrouter doit porter X-Title")
	}
	if _, ok := CloudEndpoint("inconnu", "sk"); ok {
		t.Fatal("provider inconnu ne doit pas resoudre")
	}
}

func TestBuildRegistry(t *testing.T) {
	r := Build(map[string]string{"deepseek": "sk", "openrouter": "sk2"}, map[string]string{"ollama": "http://127.0.0.1:11434"}, NewHTTPClient())
	for _, id := range []string{"deepseek", "openrouter", "ollama"} {
		if _, ok := r.Get(id); !ok {
			t.Errorf("provider %s absent", id)
		}
	}
	if _, ok := r.Get("opencode"); ok {
		t.Error("opencode sans cle ne doit pas etre enregistre")
	}
}
