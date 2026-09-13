package web

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

const (
	authLimitEvery = 6 * time.Second
	authLimitBurst = 5
	maxLimiters    = 10000
)

type ipLimiter struct {
	mu       sync.Mutex
	limiters map[string]*rate.Limiter
	every    rate.Limit
	burst    int
}

func newIPLimiter(every rate.Limit, burst int) *ipLimiter {
	return &ipLimiter{limiters: map[string]*rate.Limiter{}, every: every, burst: burst}
}

func (l *ipLimiter) allow(key string) bool {
	l.mu.Lock()
	lim, ok := l.limiters[key]
	if !ok {
		if len(l.limiters) >= maxLimiters {
			l.limiters = map[string]*rate.Limiter{}
		}
		lim = rate.NewLimiter(l.every, l.burst)
		l.limiters[key] = lim
	}
	l.mu.Unlock()
	return lim.Allow()
}

func clientIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

func (s *Server) withAuthRateLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !s.authLimiter.allow(clientIP(r)) {
			writeError(w, http.StatusTooManyRequests, "trop de tentatives, reessaie dans une minute")
			return
		}
		next(w, r)
	}
}
