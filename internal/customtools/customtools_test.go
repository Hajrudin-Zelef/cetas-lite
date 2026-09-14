package customtools

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeConfig(t *testing.T, tools map[string]Tool) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "tools.json")
	data, err := json.Marshal(Config{Tools: tools})
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestLoadConfig(t *testing.T) {
	if cfg, err := LoadConfig(filepath.Join(t.TempDir(), "absent.json")); err != nil || len(cfg.Tools) != 0 {
		t.Fatalf("absent: %+v %v", cfg, err)
	}
	path := writeConfig(t, map[string]Tool{"x": {URL: "http://x"}})
	cfg, err := LoadConfig(path)
	if err != nil || cfg.Tools["x"].URL != "http://x" {
		t.Fatalf("valide: %+v %v", cfg, err)
	}
	bad := filepath.Join(t.TempDir(), "bad.json")
	_ = os.WriteFile(bad, []byte("{"), 0o600)
	if _, err := LoadConfig(bad); err == nil {
		t.Fatal("JSON invalide doit echouer")
	}
}

func TestManagerDefsAndNames(t *testing.T) {
	path := writeConfig(t, map[string]Tool{
		"météo api": {Description: "Meteo", URL: "http://x", Parameters: map[string]any{"city": map[string]any{"type": "string"}}, Required: []string{"city"}},
		"sans-url":  {},
	})
	m, err := NewManager(path, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !m.Configured() {
		t.Fatal("doit etre configure")
	}
	defs := m.Defs()
	if len(defs) != 1 {
		t.Fatalf("defs = %+v", defs)
	}
	if defs[0].Name != "custom_m_t_o_api" {
		t.Fatalf("nom = %q", defs[0].Name)
	}
	if defs[0].Schema["type"] != "object" {
		t.Fatalf("schema = %+v", defs[0].Schema)
	}
	if _, ok := defs[0].Schema["required"]; !ok {
		t.Fatalf("required absent: %+v", defs[0].Schema)
	}
}

func TestExposedName(t *testing.T) {
	if got := ExposedName("my.tool/v2"); got != "custom_my_tool_v2" {
		t.Fatalf("got %q", got)
	}
	if got := ExposedName(strings.Repeat("a", 80)); len(got) != 64 {
		t.Fatalf("len = %d", len(got))
	}
}

func TestCallGETQuery(t *testing.T) {
	var gotQuery string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.RawQuery
		_, _ = w.Write([]byte("ok"))
	}))
	defer ts.Close()
	path := writeConfig(t, map[string]Tool{"s": {URL: ts.URL + "/search"}})
	m, _ := NewManager(path, nil)

	out, err := m.Call(context.Background(), "custom_s", `{"q":"hello world","n":3}`)
	if err != nil || out != "ok" {
		t.Fatalf("out=%q err=%v", out, err)
	}
	if !strings.Contains(gotQuery, "q=hello+world") || !strings.Contains(gotQuery, "n=3") {
		t.Fatalf("query = %q", gotQuery)
	}
}

func TestCallPathPlaceholder(t *testing.T) {
	var uri string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri = r.RequestURI
	}))
	defer ts.Close()
	path := writeConfig(t, map[string]Tool{"item": {URL: ts.URL + "/item/{id}"}})
	m, _ := NewManager(path, nil)
	if _, err := m.Call(context.Background(), "custom_item", `{"id":"a/b"}`); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(uri, "a%2Fb") {
		t.Fatalf("uri = %q", uri)
	}
}

func TestCallPOSTJSONAndBody(t *testing.T) {
	var body string
	var ctype string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		body = string(b)
		ctype = r.Header.Get("Content-Type")
	}))
	defer ts.Close()

	m1, _ := NewManager(writeConfig(t, map[string]Tool{"p": {Method: "POST", URL: ts.URL}}), nil)
	if _, err := m1.Call(context.Background(), "custom_p", `{"a":1}`); err != nil {
		t.Fatal(err)
	}
	if body != `{"a":1}` || !strings.Contains(ctype, "application/json") {
		t.Fatalf("body=%q ctype=%q", body, ctype)
	}

	m2, _ := NewManager(writeConfig(t, map[string]Tool{"t": {Method: "POST", URL: ts.URL, Body: `{"text":"{msg}"}`}}), nil)
	if _, err := m2.Call(context.Background(), "custom_t", `{"msg":"salut"}`); err != nil {
		t.Fatal(err)
	}
	if body != `{"text":"salut"}` {
		t.Fatalf("body template = %q", body)
	}
}

func TestCallHTTPErrorAndUnknown(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		_, _ = w.Write([]byte("nope"))
	}))
	defer ts.Close()
	m, _ := NewManager(writeConfig(t, map[string]Tool{"e": {URL: ts.URL}}), nil)

	out, err := m.Call(context.Background(), "custom_e", `{}`)
	if err != nil || !strings.HasPrefix(out, "[http 418]") {
		t.Fatalf("out=%q err=%v", out, err)
	}
	if _, err := m.Call(context.Background(), "custom_inconnu", `{}`); err == nil {
		t.Fatal("outil inconnu doit echouer")
	}
}
