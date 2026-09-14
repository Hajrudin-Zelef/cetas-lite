package web

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthRateLimited(t *testing.T) {
	s := newTestServerWith(t, nil)
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	tokenFor(t, ts.URL)

	limited := false
	for i := 0; i < 8; i++ {
		resp, err := http.Post(ts.URL+"/api/auth/login", "application/json",
			strings.NewReader(`{"username":"sam","password":"mauvais"}`))
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()
		if resp.StatusCode == http.StatusTooManyRequests {
			limited = true
			break
		}
	}
	if !limited {
		t.Fatal("le rate-limit doit renvoyer 429 apres plusieurs tentatives")
	}
}

func TestClientIPTrustProxy(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.RemoteAddr = "10.0.0.1:1234"
	r.Header.Set("X-Forwarded-For", "203.0.113.7, 10.0.0.1")

	if got := clientIP(r, false); got != "10.0.0.1" {
		t.Fatalf("sans confiance = %q", got)
	}
	if got := clientIP(r, true); got != "10.0.0.1" {
		t.Fatalf("avec confiance (dernier hop) = %q", got)
	}

	r.Header.Set("X-Forwarded-For", "203.0.113.7")
	if got := clientIP(r, true); got != "203.0.113.7" {
		t.Fatalf("client reel = %q", got)
	}
	r.Header.Del("X-Forwarded-For")
	r.Header.Set("X-Real-IP", "198.51.100.9")
	if got := clientIP(r, true); got != "198.51.100.9" {
		t.Fatalf("X-Real-IP = %q", got)
	}
}
