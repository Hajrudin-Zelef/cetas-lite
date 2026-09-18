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

// isAssetPath : ressources embarquées dans le binaire (CSS/JS/images). Elles
// changent à chaque déploiement : elles ne doivent JAMAIS être servies depuis
// le cache navigateur sans revalidation (sinon l'UI reste figée jusqu'à une
// heure après un rebuild).
func isAssetPath(p string) bool {
	return strings.HasPrefix(p, "/js/") || strings.HasPrefix(p, "/css/") ||
		strings.HasPrefix(p, "/images/") ||
		strings.HasPrefix(p, "/fonts/") || strings.HasPrefix(p, "/favicon")
}

func withSecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("X-Content-Type-Options", "nosniff")
		h.Set("Referrer-Policy", "no-referrer")
		h.Set("X-Frame-Options", "DENY")
		h.Set("Content-Security-Policy",
			"default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline'; "+
				"img-src 'self' data:; connect-src 'self'; "+
				"frame-ancestors 'none'; base-uri 'self'; form-action 'self'")
		h.Set("Cache-Control", "no-cache")
		next.ServeHTTP(w, r)
	})
}

// withAssetRevalidation : pour les assets, « no-cache » (revalidation) + ETag
// dérivé de la version du binaire. Tant que le binaire ne change pas, le
// navigateur reçoit 304 ; après un rebuild, la version change et il recharge.
func (s *Server) withAssetRevalidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isAssetPath(r.URL.Path) {
			etag := `"` + s.version + `:` + r.URL.Path + `"`
			w.Header().Set("ETag", etag)
			if r.Header.Get("If-None-Match") == etag {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}
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
