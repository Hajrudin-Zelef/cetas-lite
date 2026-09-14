package mcp

import (
	"context"
	"os"
	"testing"
	"time"
)

func hangConfig() ServerConfig {
	return ServerConfig{Command: os.Args[0], Env: map[string]string{
		"CETAS_MCP_HELPER": "1",
		"CETAS_MCP_HANG":   "1",
	}}
}

// Un appel annule ne doit laisser aucun canal orphelin dans pending.
func TestStdioCallCancelledCleansPending(t *testing.T) {
	tr, err := newStdioTransport(hangConfig())
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = tr.Close() }()
	st := tr

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	_, err = tr.Call(ctx, "tools/list", map[string]any{})
	if err == nil {
		t.Fatal("erreur attendue (serveur bloque)")
	}
	st.mu.Lock()
	n := len(st.pending)
	st.mu.Unlock()
	if n != 0 {
		t.Fatalf("pending=%d apres annulation, attendu 0", n)
	}
}

// Le rafraichissement est parallele : deux serveurs bloques + un sain ne
// doivent pas couter la somme des timeouts.
func TestRefreshParallelWithSlowServers(t *testing.T) {
	path := writeConfig(t, t.TempDir(), map[string]ServerConfig{
		"lent1": hangConfig(),
		"lent2": hangConfig(),
		"sain":  helperConfig(),
	})
	m, err := NewManager(path, "test")
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()
	m.timeout = 800 * time.Millisecond

	start := time.Now()
	tools := m.Tools(context.Background())
	elapsed := time.Since(start)

	// Serie : 2 x 800 ms = 1,6 s minimum. Parallele : ~800 ms.
	if elapsed > 1500*time.Millisecond {
		t.Fatalf("rafraichissement trop lent (%v) : les serveurs semblent traites en serie", elapsed)
	}
	found := false
	for _, tl := range tools {
		if tl.Server == "sain" {
			found = true
		}
	}
	if !found {
		t.Fatalf("outils du serveur sain absents: %+v", tools)
	}
	for _, s := range m.Servers() {
		if (s.Name == "lent1" || s.Name == "lent2") && s.Error == "" {
			t.Fatalf("serveur %s : erreur attendue, statut=%+v", s.Name, s)
		}
	}
}

// Un appel d'outil est borne par callTimeout : un serveur bloque ne fige
// jamais le tour.
func TestManagerCallTimeout(t *testing.T) {
	path := writeConfig(t, t.TempDir(), map[string]ServerConfig{
		"lent": hangConfig(),
	})
	m, err := NewManager(path, "test")
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()
	m.timeout = 300 * time.Millisecond
	m.callTimeout = 500 * time.Millisecond

	// Force la decouverte (le serveur bloque echoue vite)...
	_ = m.Tools(context.Background())
	// ...puis reinjecte l'outil pour tester l'appel borne.
	m.mu.Lock()
	m.byName["mcp_lent_echo"] = Tool{Exposed: "mcp_lent_echo", Server: "lent", Name: "echo"}
	m.mu.Unlock()

	start := time.Now()
	_, err = m.Call(context.Background(), "mcp_lent_echo", map[string]any{"text": "x"})
	elapsed := time.Since(start)
	if err == nil {
		t.Fatal("erreur de timeout attendue")
	}
	if elapsed > 2*time.Second {
		t.Fatalf("appel non borne (%v)", elapsed)
	}
}

// Le stderr capture est borne : pas de gonflement memoire sur la duree.
func TestLimitedSyncBuffer(t *testing.T) {
	b := &limitedSyncBuffer{max: 1024}
	big := make([]byte, 100*1024)
	if _, err := b.Write(big); err != nil {
		t.Fatal(err)
	}
	if got := len(b.String()); got != 1024 {
		t.Fatalf("stderr borne a 1024, obtenu %d", got)
	}
	if _, err := b.Write([]byte("x")); err != nil {
		t.Fatal(err)
	}
	if got := len(b.String()); got != 1024 {
		t.Fatalf("stderr doit rester borne, obtenu %d", got)
	}
}
