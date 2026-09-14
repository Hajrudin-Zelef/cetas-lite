package mcp

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

type ServerStatus struct {
	Name      string `json:"name"`
	Transport string `json:"transport"`
	Connected bool   `json:"connected"`
	Tools     int    `json:"tools"`
	Error     string `json:"error,omitempty"`
}

type Manager struct {
	path    string
	version string
	ttl     time.Duration
	timeout time.Duration
	// callTimeout borne chaque appel d'outil : un serveur MCP bloque ne doit
	// jamais figer un tour de conversation.
	callTimeout time.Duration

	mu       sync.Mutex
	config   Config
	clients  map[string]*Client
	lastErr  map[string]string
	counts   map[string]int
	cache    []Tool
	byName   map[string]Tool
	cachedAt time.Time

	refreshMu sync.Mutex
}

func NewManager(path, version string) (*Manager, error) {
	cfg, err := LoadConfig(path)
	if err != nil {
		return nil, err
	}
	return &Manager{
		path:        path,
		version:     version,
		ttl:         60 * time.Second,
		timeout:     10 * time.Second,
		callTimeout: 120 * time.Second,
		config:      cfg,
		clients:     map[string]*Client{},
		lastErr:     map[string]string{},
		counts:      map[string]int{},
		byName:      map[string]Tool{},
	}, nil
}

func (m *Manager) Configured() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.config.MCP) > 0
}

func (m *Manager) getClient(name string, cfg ServerConfig) (*Client, error) {
	m.mu.Lock()
	if c, ok := m.clients[name]; ok {
		m.mu.Unlock()
		return c, nil
	}
	m.mu.Unlock()
	c, err := NewClient(name, m.version, cfg)
	if err != nil {
		return nil, err
	}
	m.mu.Lock()
	if existing, ok := m.clients[name]; ok {
		m.mu.Unlock()
		_ = c.Close()
		return existing, nil
	}
	m.clients[name] = c
	m.mu.Unlock()
	return c, nil
}

func (m *Manager) dropClient(name string) {
	m.mu.Lock()
	c := m.clients[name]
	delete(m.clients, name)
	m.mu.Unlock()
	if c != nil {
		_ = c.Close()
	}
}

func (m *Manager) serverConfig(name string) *ServerConfig {
	m.mu.Lock()
	defer m.mu.Unlock()
	sc, ok := m.config.MCP[name]
	if !ok {
		return nil
	}
	return &sc
}

func (m *Manager) refresh(ctx context.Context) {
	m.mu.Lock()
	cfg := m.config
	m.mu.Unlock()

	// Rafraichissement parallele : un serveur lent ou bloque ne retarde pas
	// les autres (chaque serveur garde son propre timeout).
	type srvRes struct {
		name  string
		tools []Tool
		err   string
	}
	names := cfg.names()
	results := make(chan srvRes, len(names))
	var wg sync.WaitGroup
	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			sc := cfg.MCP[name]
			c, err := m.getClient(name, sc)
			if err != nil {
				m.dropClient(name)
				results <- srvRes{name: name, err: err.Error()}
				return
			}
			cctx, cancel := context.WithTimeout(ctx, m.timeout)
			tools, err := c.ListTools(cctx)
			cancel()
			if err != nil {
				m.dropClient(name)
				results <- srvRes{name: name, err: err.Error()}
				return
			}
			results <- srvRes{name: name, tools: tools}
		}(name)
	}
	wg.Wait()
	close(results)

	var all []Tool
	byName := map[string]Tool{}
	errs := map[string]string{}
	counts := map[string]int{}
	for r := range results {
		if r.err != "" {
			errs[r.name] = r.err
			continue
		}
		counts[r.name] = len(r.tools)
		for _, t := range r.tools {
			exposed := exposeName(r.name, t.Name)
			for {
				if _, exists := byName[exposed]; !exists {
					break
				}
				exposed = bumpName(exposed)
			}
			t.Exposed = exposed
			all = append(all, t)
			byName[exposed] = t
		}
	}

	m.mu.Lock()
	m.cache = all
	m.byName = byName
	m.lastErr = errs
	m.counts = counts
	m.cachedAt = time.Now()
	m.mu.Unlock()
}

func (m *Manager) Tools(ctx context.Context) []Tool {
	m.mu.Lock()
	if len(m.config.MCP) == 0 {
		m.mu.Unlock()
		return nil
	}
	if !m.cachedAt.IsZero() && time.Since(m.cachedAt) < m.ttl {
		out := append([]Tool(nil), m.cache...)
		m.mu.Unlock()
		return out
	}
	m.mu.Unlock()

	m.refreshMu.Lock()
	defer m.refreshMu.Unlock()
	m.mu.Lock()
	if !m.cachedAt.IsZero() && time.Since(m.cachedAt) < m.ttl {
		out := append([]Tool(nil), m.cache...)
		m.mu.Unlock()
		return out
	}
	m.mu.Unlock()

	m.refresh(ctx)
	m.mu.Lock()
	out := append([]Tool(nil), m.cache...)
	m.mu.Unlock()
	return out
}

func (m *Manager) Call(ctx context.Context, exposed string, args map[string]any) (string, error) {
	m.mu.Lock()
	t, ok := m.byName[exposed]
	m.mu.Unlock()
	if !ok {
		m.Tools(ctx)
		m.mu.Lock()
		t, ok = m.byName[exposed]
		m.mu.Unlock()
		if !ok {
			return "", fmt.Errorf("outil MCP inconnu: %s", exposed)
		}
	}
	sc := m.serverConfig(t.Server)
	if sc == nil {
		return "", fmt.Errorf("serveur MCP inconnu: %s", t.Server)
	}
	c, err := m.getClient(t.Server, *sc)
	if err != nil {
		return "", err
	}
	cctx, cancel := context.WithTimeout(ctx, m.callTimeout)
	defer cancel()
	out, err := c.CallTool(cctx, t.Name, args)
	if err != nil {
		m.dropClient(t.Server)
		return "", err
	}
	return out, nil
}

func (m *Manager) Servers() []ServerStatus {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make([]ServerStatus, 0, len(m.config.MCP))
	for _, name := range m.config.names() {
		sc := m.config.MCP[name]
		kind, err := sc.transport()
		st := ServerStatus{Name: name, Transport: kind, Tools: m.counts[name], Error: m.lastErr[name]}
		if err != nil {
			st.Transport = sc.Type
			st.Error = err.Error()
		}
		if c, ok := m.clients[name]; ok {
			st.Connected = c.Initialized()
		}
		out = append(out, st)
	}
	return out
}

func (m *Manager) Close() {
	m.mu.Lock()
	clients := m.clients
	m.clients = map[string]*Client{}
	m.mu.Unlock()
	for _, c := range clients {
		_ = c.Close()
	}
}

func sanitize(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '_', r == '-':
			b.WriteRune(r)
		default:
			b.WriteByte('_')
		}
	}
	return b.String()
}

func exposeName(server, tool string) string {
	name := "mcp_" + sanitize(server) + "_" + sanitize(tool)
	if len(name) > 64 {
		name = name[:64]
	}
	return name
}

func bumpName(name string) string {
	if len(name) >= 64 {
		name = name[:61]
	}
	return name + "_x"
}
