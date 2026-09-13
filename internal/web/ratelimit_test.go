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
