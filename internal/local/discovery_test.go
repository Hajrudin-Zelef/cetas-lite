package local

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestDiscoveryKeepsAliveEngines(t *testing.T) {
	var hits int64
	up := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&hits, 1)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":[{"id":"qwen3:8b"},{"id":"llama3:8b"}]}`))
	}))
	defer up.Close()
	down := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer down.Close()

	d := New(DefaultEngines(map[string]string{"ollama": up.URL, "llamacpp": down.URL}, nil), up.Client())
	models := d.Models(context.Background(), false)
	if len(models) != 2 {
		t.Fatalf("modeles = %+v", models)
	}
	for _, m := range models {
		if m.Engine != "ollama" {
			t.Fatalf("moteur inattendu: %+v", m)
		}
	}
	// cache : pas de nouveau probe
	_ = d.Models(context.Background(), false)
	if atomic.LoadInt64(&hits) != 1 {
		t.Fatalf("le cache n'a pas evite un second probe (hits=%d)", hits)
	}
	// force : re-probe
	_ = d.Models(context.Background(), true)
	if atomic.LoadInt64(&hits) != 2 {
		t.Fatalf("force n'a pas re-sonde (hits=%d)", hits)
	}
}

func TestDiscoveryModelsShape(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"models":[{"name":"m1"},{"name":"m2"}]}`))
	}))
	defer srv.Close()
	d := New(DefaultEngines(map[string]string{"lmstudio": srv.URL}, nil), srv.Client())
	models := d.ModelsForEngine(context.Background(), "lmstudio")
	if len(models) != 2 || models[0].ID != "m1" {
		t.Fatalf("models = %+v", models)
	}
}

func TestDefaultEnginesIgnoresEmpty(t *testing.T) {
	engines := DefaultEngines(map[string]string{"ollama": " http://x "}, nil)
	if len(engines) != 1 || engines[0].URL != "http://x" {
		t.Fatalf("engines = %+v", engines)
	}
}

func TestTTLExpiry(t *testing.T) {
	var hits int64
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&hits, 1)
		_, _ = w.Write([]byte(`{"data":[{"id":"a"}]}`))
	}))
	defer srv.Close()
	d := New(DefaultEngines(map[string]string{"ollama": srv.URL}, nil), srv.Client())
	d.ttl = 10 * time.Millisecond
	_ = d.Models(context.Background(), false)
	time.Sleep(20 * time.Millisecond)
	_ = d.Models(context.Background(), false)
	if atomic.LoadInt64(&hits) != 2 {
		t.Fatalf("le cache devait expirer (hits=%d)", hits)
	}
}

func TestDefaultEnginesCarriesAPIKey(t *testing.T) {
	engines := DefaultEngines(
		map[string]string{"llamacpp": "https://neva.example/"},
		map[string]string{"llamacpp": "  secret-key  "},
	)
	if len(engines) != 1 {
		t.Fatalf("engines = %+v", engines)
	}
	if engines[0].ID != "llamacpp" || engines[0].URL != "https://neva.example" {
		t.Fatalf("engine inattendu: %+v", engines[0])
	}
	if engines[0].APIKey != "secret-key" {
		t.Fatalf("cle non reprise (ou non rognee) : %q", engines[0].APIKey)
	}
}

func TestDiscoverySendsBearerWhenKey(t *testing.T) {
	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		_, _ = w.Write([]byte(`{"data":[{"id":"/opt/llama.cpp/models/Qwen3.5-4B-Q4_K_M.gguf"}]}`))
	}))
	defer srv.Close()
	d := New(DefaultEngines(
		map[string]string{"llamacpp": srv.URL},
		map[string]string{"llamacpp": "secret-key"},
	), srv.Client())
	models := d.Models(context.Background(), false)
	if gotAuth != "Bearer secret-key" {
		t.Fatalf("en-tete d'auth = %q, attendu %q", gotAuth, "Bearer secret-key")
	}
	if len(models) != 1 || models[0].ID != "/opt/llama.cpp/models/Qwen3.5-4B-Q4_K_M.gguf" {
		t.Fatalf("models = %+v", models)
	}
}

func TestDiscoveryNoAuthWithoutKey(t *testing.T) {
	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		_, _ = w.Write([]byte(`{"data":[{"id":"a"}]}`))
	}))
	defer srv.Close()
	d := New(DefaultEngines(map[string]string{"llamacpp": srv.URL}, nil), srv.Client())
	_ = d.Models(context.Background(), false)
	if gotAuth != "" {
		t.Fatalf("aucun en-tete attendu sans cle, obtenu %q", gotAuth)
	}
}
