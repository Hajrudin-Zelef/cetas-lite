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
	t.Setenv("CETAS_LITE_SANDBOX", "")
	t.Setenv("CETAS_LITE_RAG_DIR", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Home != home {
		t.Errorf("Home = %q, want %q", cfg.Home, home)
	}
	if cfg.RagDir != filepath.Join(home, "rag") {
		t.Errorf("RagDir = %q", cfg.RagDir)
	}
	if !isDir(cfg.RagDir) {
		t.Errorf("repertoire RAG absent: %s", cfg.RagDir)
	}
	if cfg.Addr != "127.0.0.1:8787" {
		t.Errorf("Addr = %q", cfg.Addr)
	}
	if cfg.DBPath != filepath.Join(home, "cetas-lite.db") {
		t.Errorf("DBPath = %q", cfg.DBPath)
	}
	if cfg.RegistrationMode != "bootstrap" {
		t.Errorf("RegistrationMode = %q, want bootstrap", cfg.RegistrationMode)
	}
	if cfg.TrustProxy {
		t.Error("TrustProxy doit etre faux par defaut")
	}
	if cfg.Sandbox != "auto" {
		t.Errorf("Sandbox = %q, want auto", cfg.Sandbox)
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
	t.Setenv("CETAS_LITE_TRUST_PROXY", "true")
	t.Setenv("CETAS_LITE_SANDBOX", "auto")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Addr != "0.0.0.0:9000" {
		t.Errorf("Addr = %q", cfg.Addr)
	}
	if cfg.RegistrationMode != "closed" {
		t.Errorf("RegistrationMode = %q, want closed", cfg.RegistrationMode)
	}
	if !cfg.TrustProxy {
		t.Error("TrustProxy devrait etre vrai")
	}
	if cfg.Sandbox != "auto" {
		t.Errorf("Sandbox = %q, want auto", cfg.Sandbox)
	}
}

func TestLoadInvalidRegistrationFlag(t *testing.T) {
	t.Setenv("CETAS_LITE_HOME", t.TempDir())
	t.Setenv("CETAS_LITE_REGISTRATION_OPEN", "peut-etre")
	if _, err := Load(); err == nil {
		t.Fatal("Load devrait echouer sur un flag invalide")
	}
}

func TestLoadInvalidSandbox(t *testing.T) {
	t.Setenv("CETAS_LITE_HOME", t.TempDir())
	t.Setenv("CETAS_LITE_SANDBOX", "containers")
	if _, err := Load(); err == nil {
		t.Fatal("Load devrait echouer sur un mode d'isolation invalide")
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

func TestRagDirOverride(t *testing.T) {
	home := t.TempDir()
	custom := filepath.Join(t.TempDir(), "corpus")
	t.Setenv("CETAS_LITE_HOME", home)
	t.Setenv("CETAS_LITE_RAG_DIR", custom)

	cfg, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.RagDir != custom {
		t.Fatalf("RagDir = %q, want %q", cfg.RagDir, custom)
	}
	if !isDir(custom) {
		t.Fatalf("repertoire RAG non cree: %s", custom)
	}
}
