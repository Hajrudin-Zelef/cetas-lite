package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"cetas-lite/internal/config"
)

func TestAccessLogPreservesFlusherAndSetsRequestID(t *testing.T) {
	s := &Server{cfg: &config.Config{}}
	var flushed bool
	h := s.withAccessLog(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, ok := w.(http.Flusher)
		if !ok {
			t.Error("flusher manquant")
			return
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("x"))
		f.Flush()
		flushed = true
	}))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/chat/state", nil))
	if !flushed {
		t.Fatal("handler non execute")
	}
	if rec.Code != http.StatusCreated {
		t.Fatalf("status=%d, attendu 201", rec.Code)
	}
	if rec.Header().Get("X-Request-Id") == "" {
		t.Fatal("X-Request-Id absent")
	}
}

func TestAccessLogSkipsAssets(t *testing.T) {
	s := &Server{cfg: &config.Config{}}
	called := false
	h := s.withAccessLog(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	}))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/js/app.js", nil))
	if !called {
		t.Fatal("les assets doivent passer")
	}
	if rec.Header().Get("X-Request-Id") != "" {
		t.Fatal("X-Request-Id ne doit pas etre pose sur les assets")
	}
}
