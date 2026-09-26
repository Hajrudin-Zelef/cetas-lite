package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	Home             string
	Addr             string
	DataDir          string
	WorkspaceDir     string
	MemoryDir        string
	RagDir           string
	MCPPath          string
	ToolsPath        string
	PluginsDir       string
	DBPath           string
	RegistrationMode string
	AllowScript      bool
	TrustProxy       bool
	Sandbox          string
	// Jambe semantique : backend d'embedding (desktop|openrouter) et URL du
	// serveur desktop (bge-m3 + rerank). Vides => BM25 seul (fail-open).
	EmbedBackend    string
	EmbedURL        string
	EmbedModel      string
	RerankTopN      int
	RerankTimeoutMs int
}

func Load() (*Config, error) {
	home := strings.TrimSpace(os.Getenv("CETAS_LITE_HOME"))
	if home == "" {
		base, err := os.UserConfigDir()
		if err != nil {
			return nil, fmt.Errorf("resolution du repertoire de configuration: %w", err)
		}
		home = filepath.Join(base, "cetas-lite")
	}
	home, err := filepath.Abs(home)
	if err != nil {
		return nil, fmt.Errorf("chemin CETAS_LITE_HOME: %w", err)
	}

	addr := strings.TrimSpace(os.Getenv("CETAS_LITE_ADDR"))
	if addr == "" {
		addr = "127.0.0.1:8787"
	}

	ragDir := strings.TrimSpace(os.Getenv("CETAS_LITE_RAG_DIR"))
	if ragDir == "" {
		ragDir = filepath.Join(home, "rag")
	} else if abs, err := filepath.Abs(ragDir); err == nil {
		ragDir = abs
	}

	registrationMode := "bootstrap"
	if raw := strings.TrimSpace(os.Getenv("CETAS_LITE_REGISTRATION_OPEN")); raw != "" {
		switch strings.ToLower(raw) {
		case "true", "1", "yes":
			registrationMode = "open"
		case "false", "0", "no":
			registrationMode = "closed"
		default:
			return nil, fmt.Errorf("CETAS_LITE_REGISTRATION_OPEN invalide: %q", raw)
		}
	}

	allowScript := false
	if raw := strings.TrimSpace(os.Getenv("CETAS_LITE_ALLOW_SCRIPT")); raw != "" {
		v, err := strconv.ParseBool(raw)
		if err != nil {
			return nil, fmt.Errorf("CETAS_LITE_ALLOW_SCRIPT invalide: %q", raw)
		}
		allowScript = v
	}

	trustProxy := false
	if raw := strings.TrimSpace(os.Getenv("CETAS_LITE_TRUST_PROXY")); raw != "" {
		v, err := strconv.ParseBool(raw)
		if err != nil {
			return nil, fmt.Errorf("CETAS_LITE_TRUST_PROXY invalide: %q", raw)
		}
		trustProxy = v
	}

	sandbox := strings.ToLower(strings.TrimSpace(os.Getenv("CETAS_LITE_SANDBOX")))
	if sandbox == "" {
		// M4 : défaut "auto" (bwrap quand disponible, repli "none" sinon)
		// plutôt que "none" pur.
		sandbox = "auto"
	}
	switch sandbox {
	case "none", "auto", "bwrap":
	default:
		return nil, fmt.Errorf("CETAS_LITE_SANDBOX invalide: %q (none|auto|bwrap)", sandbox)
	}

	// Jambe semantique : backend d'embedding. "desktop" (defaut) pointe sur
	// la plateforme RAG bge-m3 ; "openrouter" conserve l'ancien chemin
	// (rollback manuel : rebuild vecteurs + flip de cette variable).
	embedBackend := strings.ToLower(strings.TrimSpace(os.Getenv("CETAS_LITE_EMBED_BACKEND")))
	if embedBackend == "" {
		embedBackend = "desktop"
	}
	switch embedBackend {
	case "desktop", "openrouter":
	default:
		return nil, fmt.Errorf("CETAS_LITE_EMBED_BACKEND invalide: %q (desktop|openrouter)", embedBackend)
	}
	embedURL := strings.TrimRight(strings.TrimSpace(os.Getenv("CETAS_LITE_EMBED_URL")), "/")
	embedModel := strings.TrimSpace(os.Getenv("CETAS_LITE_EMBED_MODEL"))
	if embedModel == "" {
		if embedBackend == "desktop" {
			embedModel = "BAAI/bge-m3"
		} else {
			embedModel = "openai-text-embedding-3-small"
		}
	}
	rerankTopN := 20
	if raw := strings.TrimSpace(os.Getenv("CETAS_LITE_RERANK_TOPN")); raw != "" {
		v, err := strconv.Atoi(raw)
		if err != nil || v < 1 {
			return nil, fmt.Errorf("CETAS_LITE_RERANK_TOPN invalide: %q", raw)
		}
		rerankTopN = v
	}
	rerankTimeoutMs := 4000
	if raw := strings.TrimSpace(os.Getenv("CETAS_LITE_RERANK_TIMEOUT_MS")); raw != "" {
		v, err := strconv.Atoi(raw)
		if err != nil || v < 1 {
			return nil, fmt.Errorf("CETAS_LITE_RERANK_TIMEOUT_MS invalide: %q", raw)
		}
		rerankTimeoutMs = v
	}

	cfg := &Config{
		Home:             home,
		Addr:             addr,
		DataDir:          filepath.Join(home, "data"),
		WorkspaceDir:     filepath.Join(home, "workspace"),
		MemoryDir:        filepath.Join(home, "memory"),
		RagDir:           ragDir,
		MCPPath:          filepath.Join(home, "mcp.json"),
		ToolsPath:        filepath.Join(home, "tools.json"),
		PluginsDir:       filepath.Join(home, "plugins"),
		DBPath:           filepath.Join(home, "cetas-lite.db"),
		RegistrationMode: registrationMode,
		AllowScript:      allowScript,
		TrustProxy:       trustProxy,
		Sandbox:          sandbox,
		EmbedBackend:     embedBackend,
		EmbedURL:         embedURL,
		EmbedModel:       embedModel,
		RerankTopN:       rerankTopN,
		RerankTimeoutMs:  rerankTimeoutMs,
	}

	for _, dir := range []string{cfg.Home, cfg.DataDir, cfg.WorkspaceDir, cfg.MemoryDir, cfg.PluginsDir} {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return nil, fmt.Errorf("creation de %s: %w", dir, err)
		}
	}
	// Dossier du corpus RAG : cree au mieux. Un chemin absent ou non
	// inscriptible laisse simplement le RAG inactif — jamais bloquant.
	_ = os.MkdirAll(cfg.RagDir, 0o700)

	return cfg, nil
}
