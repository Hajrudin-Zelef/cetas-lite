package chat

import (
	"context"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"cetas-lite/internal/provider"
	"cetas-lite/internal/vfs"
)

// TestStartTurnClientMsgID vérifie que l'identifiant d'écho optimiste est
// renvoyé dans le delta "user" quand il est fourni, et absent sinon.
func TestStartTurnClientMsgID(t *testing.T) {
	newConv := func() *Conversation {
		r := &fakeRunner{started: make(chan TurnInput, 1), block: make(chan struct{})}
		c := NewConversation("1", r, nil)
		return c
	}

	// Avec ClientMsgID : le delta user porte client_msg_id.
	c := newConv()
	in := TurnInput{Text: "bonjour", ClientMsgID: "abc-123"}
	if err := c.StartTurn(in); err != nil {
		t.Fatal(err)
	}
	found := false
	for _, ev := range c.Log {
		if u, ok := ev.Delta["user"]; ok {
			if u != "bonjour" {
				t.Fatalf("user = %v", u)
			}
			if id, ok := ev.Delta["client_msg_id"]; !ok || id != "abc-123" {
				t.Fatalf("client_msg_id = %v, want abc-123", ev.Delta["client_msg_id"])
			}
			found = true
		}
	}
	if !found {
		t.Fatal("delta user introuvable dans le log")
	}
	c.Reset()

	// Sans ClientMsgID : pas de clé client_msg_id (comportement historique).
	c2 := newConv()
	if err := c2.StartTurn(TurnInput{Text: "salut"}); err != nil {
		t.Fatal(err)
	}
	for _, ev := range c2.Log {
		if _, ok := ev.Delta["user"]; ok {
			if _, ok := ev.Delta["client_msg_id"]; ok {
				t.Fatal("client_msg_id ne devrait pas être présent sans ClientMsgID")
			}
		}
	}
	c2.Reset()
}

// TestWorkspaceSnapshotCache vérifie que le snapshot est servi depuis le
// cache pendant le TTL (pas de nouveau walk FS) puis régénéré après
// expiration.
func TestWorkspaceSnapshotCache(t *testing.T) {
	ctx := context.Background()
	lfs, err := vfs.NewLocal(filepath.Join(t.TempDir(), "ws"), "test")
	if err != nil {
		t.Fatal(err)
	}
	_ = lfs.WriteFile(ctx, "a.txt", []byte("a"), 0o644)
	sb := NewSandboxFS(lfs)
	root := sb.Root()
	key := root + "\x00local\x00Demo"

	snapshotCache.Lock()
	snapshotCache.entries = map[string]snapshotEntry{}
	snapshotCache.Unlock()

	// 1. Premier appel : walk réel, contient a.txt.
	m1 := workspaceSnapshotMessage(ctx, sb, "Demo")
	if !strings.Contains(m1.Content.(string), "a.txt") {
		t.Fatalf("snapshot sans a.txt:\n%s", m1.Content)
	}

	// 2. On modifie le FS, mais le cache sert l'ancien contenu.
	_ = lfs.WriteFile(ctx, "b.txt", []byte("b"), 0o644)
	m2 := workspaceSnapshotMessage(ctx, sb, "Demo")
	c2 := m2.Content.(string)
	if strings.Contains(c2, "b.txt") {
		t.Fatal("le cache aurait dû servir l'ancien snapshot (sans b.txt)")
	}
	if !strings.Contains(c2, "a.txt") {
		t.Fatal("le snapshot en cache devrait toujours contenir a.txt")
	}

	// 3. Expiration simulée : le walk est refait, b.txt apparaît.
	snapshotCache.Lock()
	e := snapshotCache.entries[key]
	e.at = time.Now().Add(-2 * snapshotCacheTTL)
	snapshotCache.entries[key] = e
	snapshotCache.Unlock()
	m3 := workspaceSnapshotMessage(ctx, sb, "Demo")
	if !strings.Contains(m3.Content.(string), "b.txt") {
		t.Fatalf("après expiration, le snapshot devrait contenir b.txt:\n%s", m3.Content)
	}

	// 4. Entrée pré-remplie : aucun walk (même un FS en erreur est ignoré).
	snapshotCache.Lock()
	snapshotCache.entries[key] = snapshotEntry{
		msg: provider.Message{Role: "system", Content: "CACHED"},
		at:  time.Now(),
	}
	snapshotCache.Unlock()
	m4 := workspaceSnapshotMessage(ctx, sb, "Demo")
	if m4.Content.(string) != "CACHED" {
		t.Fatalf("cache non utilisé: %q", m4.Content)
	}

	snapshotCache.Lock()
	delete(snapshotCache.entries, key)
	snapshotCache.Unlock()
}
