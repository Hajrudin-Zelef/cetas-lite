package web

import (
	"net/http"

	"cetas-lite/internal/auth"
	"cetas-lite/internal/chat"
	"cetas-lite/internal/config"
	"cetas-lite/internal/store"
	webassets "cetas-lite/web"
)

type Server struct {
	cfg     *config.Config
	st      *store.Store
	auth    *auth.Manager
	engine  *chat.Engine
	version string
	handler http.Handler
}

func New(cfg *config.Config, st *store.Store, authMgr *auth.Manager, engine *chat.Engine, version string) *Server {
	s := &Server{cfg: cfg, st: st, auth: authMgr, engine: engine, version: version}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", s.handleHealth)
	mux.HandleFunc("GET /api/config", s.handleConfig)
	mux.HandleFunc("POST /api/auth/register", s.handleRegister)
	mux.HandleFunc("POST /api/auth/login", s.handleLogin)
	mux.HandleFunc("GET /api/me", s.requireAuth(s.handleMe))
	mux.HandleFunc("GET /api/aliases", s.requireAuth(s.handleAliasesGet))
	mux.HandleFunc("PUT /api/aliases", s.requireAuth(s.handleAliasesPut))
	mux.HandleFunc("POST /api/chat/send", s.requireAuth(s.handleChatSend))
	mux.HandleFunc("GET /api/chat/stream", s.requireAuth(s.handleChatStream))
	mux.HandleFunc("POST /api/chat/stop", s.requireAuth(s.handleChatStop))
	mux.HandleFunc("POST /api/chat/reset", s.requireAuth(s.handleChatReset))
	mux.HandleFunc("GET /api/chat/state", s.requireAuth(s.handleChatState))
	mux.Handle("GET /", http.FileServerFS(webassets.FS))

	s.handler = withRecovery(withSecurityHeaders(mux))
	return s
}

func (s *Server) Handler() http.Handler {
	return s.handler
}
