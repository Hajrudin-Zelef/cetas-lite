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
	DBPath           string
	RegistrationOpen bool
	AllowScript      bool
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

	registrationOpen := true
	if raw := strings.TrimSpace(os.Getenv("CETAS_LITE_REGISTRATION_OPEN")); raw != "" {
		v, err := strconv.ParseBool(raw)
		if err != nil {
			return nil, fmt.Errorf("CETAS_LITE_REGISTRATION_OPEN invalide: %q", raw)
		}
		registrationOpen = v
	}

	allowScript := false
	if raw := strings.TrimSpace(os.Getenv("CETAS_LITE_ALLOW_SCRIPT")); raw != "" {
		v, err := strconv.ParseBool(raw)
		if err != nil {
			return nil, fmt.Errorf("CETAS_LITE_ALLOW_SCRIPT invalide: %q", raw)
		}
		allowScript = v
	}

	cfg := &Config{
		Home:             home,
		Addr:             addr,
		DataDir:          filepath.Join(home, "data"),
		WorkspaceDir:     filepath.Join(home, "workspace"),
		MemoryDir:        filepath.Join(home, "memory"),
		DBPath:           filepath.Join(home, "cetas-lite.db"),
		RegistrationOpen: registrationOpen,
		AllowScript:      allowScript,
	}

	for _, dir := range []string{cfg.Home, cfg.DataDir, cfg.WorkspaceDir, cfg.MemoryDir} {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return nil, fmt.Errorf("creation de %s: %w", dir, err)
		}
	}

	return cfg, nil
}
