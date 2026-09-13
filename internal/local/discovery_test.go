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

	d := New(DefaultEngines(map[string]string{"ollama": up.URL, "llamacpp": down.URL}), up.Client())
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
	d := New(DefaultEngines(map[string]string{"lmstudio": srv.URL}), srv.Client())
	models := d.ModelsForEngine(context.Background(), "lmstudio")
	if len(models) != 2 || models[0].ID != "m1" {
		t.Fatalf("models = %+v", models)
	}
}

func TestDefaultEnginesIgnoresEmpty(t *testing.T) {
	engines := DefaultEngines(map[string]string{"ollama": " http://x "})
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
	d := New(DefaultEngines(map[string]string{"ollama": srv.URL}), srv.Client())
	d.ttl = 10 * time.Millisecond
	_ = d.Models(context.Background(), false)
	time.Sleep(20 * time.Millisecond)
	_ = d.Models(context.Background(), false)
	if atomic.LoadInt64(&hits) != 2 {
		t.Fatalf("le cache devait expirer (hits=%d)", hits)
	}
}
