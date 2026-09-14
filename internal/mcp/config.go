package mcp

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type ServerConfig struct {
	Type    string            `json:"type,omitempty"`
	Command string            `json:"command,omitempty"`
	Args    []string          `json:"args,omitempty"`
	Env     map[string]string `json:"env,omitempty"`
	URL     string            `json:"url,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

type Config struct {
	MCP map[string]ServerConfig `json:"mcp"`
}

func LoadConfig(path string) (Config, error) {
	empty := Config{MCP: map[string]ServerConfig{}}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return empty, nil
		}
		return empty, fmt.Errorf("lecture %s: %w", path, err)
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return empty, fmt.Errorf("mcp.json invalide: %w", err)
	}
	if cfg.MCP == nil {
		cfg.MCP = map[string]ServerConfig{}
	}
	return cfg, nil
}

func (s ServerConfig) transport() (string, error) {
	switch {
	case s.Type == "stdio" || (s.Type == "" && strings.TrimSpace(s.Command) != ""):
		if strings.TrimSpace(s.Command) == "" {
			return "", fmt.Errorf("transport stdio: command requis")
		}
		return "stdio", nil
	case s.Type == "http" || (s.Type == "" && strings.TrimSpace(s.URL) != ""):
		if strings.TrimSpace(s.URL) == "" {
			return "", fmt.Errorf("transport http: url requise")
		}
		return "http", nil
	case s.Type == "sse":
		return "", fmt.Errorf("transport sse non supporte (utiliser http)")
	}
	return "", fmt.Errorf("transport indetermine: command ou url requis")
}

func (c Config) names() []string {
	out := make([]string, 0, len(c.MCP))
	for name := range c.MCP {
		out = append(out, name)
	}
	sort.Strings(out)
	return out
}
