package web

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"net/http"
	"strings"
	"time"

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
				"img-src 'self' data: blob:; connect-src 'self'; "+
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

// statusWriter : capture statut et volume pour le journal d'acces. Il doit
// preserver http.Flusher (le SSE en depend) et Unwrap (http.ResponseController).
type statusWriter struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (w *statusWriter) WriteHeader(code int) {
	if w.status == 0 {
		w.status = code
	}
	w.ResponseWriter.WriteHeader(code)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	n, err := w.ResponseWriter.Write(b)
	w.bytes += n
	return n, err
}

func (w *statusWriter) Flush() {
	if f, ok := w.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

func (w *statusWriter) Unwrap() http.ResponseWriter { return w.ResponseWriter }

// newRequestID : identifiant court par requete, renvoye en X-Request-Id pour
// correler les logs serveur avec le navigateur.
func newRequestID() string {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "-"
	}
	return hex.EncodeToString(b[:])
}

// withAccessLog : une ligne par requete (hors assets). Sert a prouver, entre
// autres, si un tour a recu une ou deux requetes HTTP identiques.
func (s *Server) withAccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isAssetPath(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}
		id := newRequestID()
		w.Header().Set("X-Request-Id", id)
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w}
		next.ServeHTTP(sw, r)
		status := sw.status
		if status == 0 {
			status = http.StatusOK
		}
		slog.Info("http",
			"id", id,
			"method", r.Method,
			"path", r.URL.Path,
			"status", status,
			"ms", time.Since(start).Milliseconds(),
			"bytes", sw.bytes,
			"ip", clientIP(r, s.cfg.TrustProxy),
		)
	})
}
