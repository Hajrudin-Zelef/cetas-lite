package config

import (
	"os"
	"path/filepath"
	"testing"
)

func isDir(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.IsDir()
}

func TestLoadDefaults(t *testing.T) {
	home := t.TempDir()
	t.Setenv("CETAS_LITE_HOME", home)
	t.Setenv("CETAS_LITE_ADDR", "")
	t.Setenv("CETAS_LITE_REGISTRATION_OPEN", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Home != home {
		t.Errorf("Home = %q, want %q", cfg.Home, home)
	}
	if cfg.Addr != "127.0.0.1:8787" {
		t.Errorf("Addr = %q", cfg.Addr)
	}
	if cfg.DBPath != filepath.Join(home, "cetas-lite.db") {
		t.Errorf("DBPath = %q", cfg.DBPath)
	}
	if !cfg.RegistrationOpen {
		t.Error("RegistrationOpen devrait etre vrai par defaut")
	}
	for _, dir := range []string{cfg.Home, cfg.DataDir, cfg.WorkspaceDir, cfg.MemoryDir} {
		if !isDir(dir) {
			t.Errorf("repertoire absent: %s", dir)
		}
	}
}

func TestLoadOverrides(t *testing.T) {
	home := t.TempDir()
	t.Setenv("CETAS_LITE_HOME", home)
	t.Setenv("CETAS_LITE_ADDR", "0.0.0.0:9000")
	t.Setenv("CETAS_LITE_REGISTRATION_OPEN", "false")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Addr != "0.0.0.0:9000" {
		t.Errorf("Addr = %q", cfg.Addr)
	}
	if cfg.RegistrationOpen {
		t.Error("RegistrationOpen devrait etre faux")
	}
}

func TestLoadInvalidRegistrationFlag(t *testing.T) {
	t.Setenv("CETAS_LITE_HOME", t.TempDir())
	t.Setenv("CETAS_LITE_REGISTRATION_OPEN", "peut-etre")
	if _, err := Load(); err == nil {
		t.Fatal("Load devrait echouer sur un flag invalide")
	}
}

func TestAllowScriptFlag(t *testing.T) {
	t.Setenv("CETAS_LITE_HOME", t.TempDir())
	t.Setenv("CETAS_LITE_ALLOW_SCRIPT", "")
	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.AllowScript {
		t.Fatal("AllowScript doit etre faux par defaut")
	}
	t.Setenv("CETAS_LITE_ALLOW_SCRIPT", "true")
	cfg, err = Load()
	if err != nil {
		t.Fatal(err)
	}
	if !cfg.AllowScript {
		t.Fatal("AllowScript doit passer a vrai")
	}
	t.Setenv("CETAS_LITE_ALLOW_SCRIPT", "oups")
	if _, err := Load(); err == nil {
		t.Fatal("Load devrait echouer sur CETAS_LITE_ALLOW_SCRIPT invalide")
	}
}
