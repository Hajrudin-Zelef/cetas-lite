package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestChatExportMarkdown(t *testing.T) {
	s := newTestServerWith(t, &fakeProvider{content: "bonjour"})
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)
	sendAndWait(t, ts.URL, tok, "salut")

	rec := doJSON(t, s.Handler(), http.MethodGet, "/api/chat/export?format=md", tok, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d (%s)", rec.Code, rec.Body.String())
	}
	body := rec.Body.String()
	if !strings.Contains(body, "## Vous") || !strings.Contains(body, "salut") || !strings.Contains(body, "## Assistant") {
		t.Fatalf("markdown inattendu: %q", body)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.Contains(ct, "text/markdown") {
		t.Fatalf("content-type = %q", ct)
	}
	if cd := rec.Header().Get("Content-Disposition"); !strings.Contains(cd, "attachment") {
		t.Fatalf("content-disposition = %q", cd)
	}
}

func TestChatExportJSON(t *testing.T) {
	s := newTestServerWith(t, &fakeProvider{content: "bonjour"})
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)
	sendAndWait(t, ts.URL, tok, "salut")

	rec := doJSON(t, s.Handler(), http.MethodGet, "/api/chat/export?format=json", tok, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), `"messages"`) {
		t.Fatalf("json inattendu: %q", rec.Body.String())
	}
}

func TestArchiveExport(t *testing.T) {
	s := newTestServerWith(t, &fakeProvider{content: "bonjour"})
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tok := tokenFor(t, ts.URL)
	sendAndWait(t, ts.URL, tok, "salut")

	rec := doJSON(t, s.Handler(), http.MethodPost, "/api/chat/reset", tok, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("reset status = %d", rec.Code)
	}
	list := doJSON(t, s.Handler(), http.MethodGet, "/api/conversations", tok, nil)
	body := decode(t, list)
	archives, _ := body["archives"].([]any)
	if len(archives) != 1 {
		t.Fatalf("archives = %v", archives)
	}
	id := archives[0].(map[string]any)["id"].(string)

	rec = doJSON(t, s.Handler(), http.MethodGet, "/api/conversations/"+id+"/export?format=md", tok, nil)
	if rec.Code != http.StatusOK {
		t.Fatalf("export status = %d (%s)", rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "salut") {
		t.Fatalf("archive export = %q", rec.Body.String())
	}

	rec = doJSON(t, s.Handler(), http.MethodGet, "/api/conversations/inconnu/export", tok, nil)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("export inconnu status = %d", rec.Code)
	}
}
