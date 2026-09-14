package web

import (
	"net/http"

	"cetas-lite/internal/auth"
	"cetas-lite/internal/chat"
	"cetas-lite/internal/config"
	"cetas-lite/internal/provider"
	"cetas-lite/internal/store"
	"cetas-lite/internal/terminal"
	webassets "cetas-lite/web"

	"golang.org/x/time/rate"
)

type Server struct {
	cfg         *config.Config
	st          *store.Store
	auth        *auth.Manager
	engine      *chat.Engine
	terminals   *terminal.Manager
	registry    *provider.Registry
	httpClient  *http.Client
	version     string
	authLimiter *ipLimiter
	handler     http.Handler
}

func New(cfg *config.Config, st *store.Store, authMgr *auth.Manager, engine *chat.Engine, termMgr *terminal.Manager, registry *provider.Registry, httpClient *http.Client, version string) *Server {
	s := &Server{
		cfg: cfg, st: st, auth: authMgr, engine: engine, terminals: termMgr,
		registry: registry, httpClient: httpClient, version: version,
		authLimiter: newIPLimiter(rate.Every(authLimitEvery), authLimitBurst),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", s.handleHealth)
	mux.HandleFunc("GET /api/config", s.handleConfig)
	mux.HandleFunc("POST /api/auth/register", s.withAuthRateLimit(s.handleRegister))
	mux.HandleFunc("POST /api/auth/login", s.withAuthRateLimit(s.handleLogin))
	mux.HandleFunc("GET /api/me", s.requireAuth(s.handleMe))
	mux.HandleFunc("GET /api/settings", s.requireAuth(s.handleSettingsGet))
	mux.HandleFunc("PUT /api/settings", s.requireAuth(s.handleSettingsPut))
	mux.HandleFunc("GET /api/mcp", s.requireAuth(s.handleMCPStatus))
	mux.HandleFunc("GET /api/capabilities", s.requireAuth(s.handleCapabilitiesGet))
	mux.HandleFunc("PUT /api/capabilities", s.requireAuth(s.handleCapabilitiesPut))
	mux.HandleFunc("GET /api/aliases", s.requireAuth(s.handleAliasesGet))
	mux.HandleFunc("PUT /api/aliases", s.requireAuth(s.handleAliasesPut))
	mux.HandleFunc("GET /api/catalog", s.requireAuth(s.handleCatalogGet))
	mux.HandleFunc("GET /api/providers", s.requireAuth(s.handleProvidersList))
	mux.HandleFunc("GET /api/model-info", s.requireAuth(s.handleModelInfo))
	mux.HandleFunc("GET /api/metrics", s.requireAuth(s.handleMetrics))
	mux.HandleFunc("PUT /api/providers/{id}", s.requireAuth(s.handleProviderPut))
	mux.HandleFunc("DELETE /api/providers/{id}", s.requireAuth(s.handleProviderDelete))
	mux.HandleFunc("POST /api/chat/send", s.requireAuth(s.handleChatSend))
	mux.HandleFunc("POST /api/agents", s.requireAuth(s.handleAgentsCreate))
	mux.HandleFunc("GET /api/agents", s.requireAuth(s.handleAgentsList))
	mux.HandleFunc("GET /api/agents/{id}/stream", s.requireAuth(s.handleAgentStream))
	mux.HandleFunc("GET /api/agents/{id}/state", s.requireAuth(s.handleAgentState))
	mux.HandleFunc("POST /api/agents/{id}/message", s.requireAuth(s.handleAgentMessage))
	mux.HandleFunc("POST /api/agents/{id}/stop", s.requireAuth(s.handleAgentStop))
	mux.HandleFunc("POST /api/agents/{id}/approve", s.requireAuth(s.handleAgentApprove))
	mux.HandleFunc("DELETE /api/agents/{id}", s.requireAuth(s.handleAgentDelete))
	mux.HandleFunc("POST /api/terminal", s.requireAuth(s.handleTerminalCreate))
	mux.HandleFunc("GET /api/terminal", s.requireAuth(s.handleTerminalList))
	mux.HandleFunc("GET /api/terminal/{id}/stream", s.requireAuth(s.handleTerminalStream))
	mux.HandleFunc("POST /api/terminal/{id}/input", s.requireAuth(s.handleTerminalInput))
	mux.HandleFunc("POST /api/terminal/{id}/resize", s.requireAuth(s.handleTerminalResize))
	mux.HandleFunc("DELETE /api/terminal/{id}", s.requireAuth(s.handleTerminalDelete))
	mux.HandleFunc("POST /api/chat/attach", s.requireAuth(s.handleAttachUpload))
	mux.HandleFunc("GET /api/chat/attach/{id}", s.requireAuth(s.handleAttachGet))
	mux.HandleFunc("DELETE /api/chat/attach/{id}", s.requireAuth(s.handleAttachDelete))
	mux.HandleFunc("GET /api/chat/stream", s.requireAuth(s.handleChatStream))
	mux.HandleFunc("POST /api/chat/stop", s.requireAuth(s.handleChatStop))
	mux.HandleFunc("POST /api/chat/approve", s.requireAuth(s.handleChatApprove))
	mux.HandleFunc("POST /api/chat/reset", s.requireAuth(s.handleChatReset))
	mux.HandleFunc("POST /api/chat/regenerate", s.requireAuth(s.handleChatRegenerate))
	mux.HandleFunc("GET /api/chat/state", s.requireAuth(s.handleChatState))
	mux.HandleFunc("GET /api/chat/export", s.requireAuth(s.handleChatExport))
	mux.HandleFunc("GET /api/conversations", s.requireAuth(s.handleConversationsList))
	mux.HandleFunc("POST /api/conversations/restore", s.requireAuth(s.handleConversationRestore))
	mux.HandleFunc("DELETE /api/conversations/{id}", s.requireAuth(s.handleConversationDelete))
	mux.HandleFunc("GET /api/conversations/{id}/export", s.requireAuth(s.handleArchiveExport))
	mux.Handle("GET /", http.FileServerFS(webassets.FS))

	s.handler = withRecovery(withSecurityHeaders(mux))
	return s
}

func (s *Server) Handler() http.Handler {
	return s.handler
}
