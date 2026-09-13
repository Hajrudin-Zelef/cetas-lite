package web

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"cetas-lite/internal/auth"
)

type ctxKey int

const claimsKey ctxKey = iota

func claimsFrom(r *http.Request) *auth.Claims {
	if c, ok := r.Context().Value(claimsKey).(*auth.Claims); ok {
		return c
	}
	return nil
}

func (s *Server) requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			writeError(w, http.StatusUnauthorized, "non authentifie")
			return
		}
		claims, err := s.auth.Parse(strings.TrimSpace(strings.TrimPrefix(header, "Bearer ")))
		if err != nil {
			writeError(w, http.StatusUnauthorized, "token invalide ou expire")
			return
		}
		ctx := context.WithValue(r.Context(), claimsKey, claims)
		next(w, r.WithContext(ctx))
	}
}

func withSecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("X-Frame-Options", "DENY")
		next.ServeHTTP(w, r)
	})
}

func withRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				slog.Error("panic HTTP", "path", r.URL.Path, "recover", rec)
				writeError(w, http.StatusInternalServerError, "erreur serveur")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
