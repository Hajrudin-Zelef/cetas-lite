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
	MCPPath          string
	ToolsPath        string
	PluginsDir       string
	DBPath           string
	RegistrationMode string
	AllowScript      bool
	TrustProxy       bool
	Sandbox          string
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
		sandbox = "none"
	}
	switch sandbox {
	case "none", "auto", "bwrap":
	default:
		return nil, fmt.Errorf("CETAS_LITE_SANDBOX invalide: %q (none|auto|bwrap)", sandbox)
	}

	cfg := &Config{
		Home:             home,
		Addr:             addr,
		DataDir:          filepath.Join(home, "data"),
		WorkspaceDir:     filepath.Join(home, "workspace"),
		MemoryDir:        filepath.Join(home, "memory"),
		MCPPath:          filepath.Join(home, "mcp.json"),
		ToolsPath:        filepath.Join(home, "tools.json"),
		PluginsDir:       filepath.Join(home, "plugins"),
		DBPath:           filepath.Join(home, "cetas-lite.db"),
		RegistrationMode: registrationMode,
		AllowScript:      allowScript,
		TrustProxy:       trustProxy,
		Sandbox:          sandbox,
	}

	for _, dir := range []string{cfg.Home, cfg.DataDir, cfg.WorkspaceDir, cfg.MemoryDir, cfg.PluginsDir} {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return nil, fmt.Errorf("creation de %s: %w", dir, err)
		}
	}

	return cfg, nil
}
