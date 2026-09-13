package web

import (
	"net/http"

	"cetas-lite/internal/auth"
	"cetas-lite/internal/config"
	"cetas-lite/internal/store"
	webassets "cetas-lite/web"
)

type Server struct {
	cfg     *config.Config
	st      *store.Store
	auth    *auth.Manager
	version string
	handler http.Handler
}

func New(cfg *config.Config, st *store.Store, authMgr *auth.Manager, version string) *Server {
	s := &Server{cfg: cfg, st: st, auth: authMgr, version: version}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", s.handleHealth)
	mux.HandleFunc("GET /api/config", s.handleConfig)
	mux.HandleFunc("POST /api/auth/register", s.handleRegister)
	mux.HandleFunc("POST /api/auth/login", s.handleLogin)
	mux.HandleFunc("GET /api/me", s.requireAuth(s.handleMe))
	mux.Handle("GET /", http.FileServerFS(webassets.FS))

	s.handler = withRecovery(withSecurityHeaders(mux))
	return s
}

func (s *Server) Handler() http.Handler {
	return s.handler
}
