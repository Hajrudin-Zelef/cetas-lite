package terminal

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestCreateAndIO(t *testing.T) {
	dir := t.TempDir()
	m := NewManager(dir)
	s, err := m.Create("alice", dir)
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	ch, backlog, exited := s.Subscribe()
	defer s.Unsubscribe(ch)
	_ = backlog

	if err := s.WriteInput([]byte("echo hello-terminal\n")); err != nil {
		t.Fatalf("write: %v", err)
	}
	deadline := time.After(5 * time.Second)
	got := ""
	for {
		select {
		case p, ok := <-ch:
			if !ok {
				t.Fatalf("canal ferme trop tot, recu: %q", got)
			}
			got += string(p)
			if strings.Contains(got, "hello-terminal") {
				goto done
			}
		case <-deadline:
			t.Fatalf("timeout, recu: %q", got)
		case <-exited:
			t.Fatalf("shell termine prematurement, recu: %q", got)
		}
	}
done:
	if !m.Delete("alice", s.ID) {
		t.Fatal("delete a echoue")
	}
	select {
	case <-exited:
	case <-time.After(5 * time.Second):
		t.Fatal("la session ne s'est pas terminee")
	}
}

func TestCwdValidation(t *testing.T) {
	allowed := t.TempDir()
	outside := t.TempDir()
	m := NewManager(allowed)
	if _, err := m.Create("alice", outside); err != ErrBadCwd {
		t.Fatalf("attendu ErrBadCwd, recu: %v", err)
	}
	// Sous-repertoire autorise.
	sub := filepath.Join(allowed, "sub")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	s, err := m.Create("alice", sub)
	if err != nil {
		t.Fatalf("create sub: %v", err)
	}
	m.Delete("alice", s.ID)
	// Traversee de chemin refusee.
	if _, err := m.Create("alice", filepath.Join(allowed, "..", "x")); err != ErrBadCwd {
		t.Fatalf("traversee: attendu ErrBadCwd, recu: %v", err)
	}
}

func TestIsolationBetweenUsers(t *testing.T) {
	dir := t.TempDir()
	m := NewManager(dir)
	s, err := m.Create("alice", dir)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := m.Get("bob", s.ID); ok {
		t.Fatal("bob ne devrait pas voir la session d'alice")
	}
	if m.Delete("bob", s.ID) {
		t.Fatal("bob ne devrait pas pouvoir supprimer la session d'alice")
	}
	if !m.Delete("alice", s.ID) {
		t.Fatal("alice devrait pouvoir supprimer sa session")
	}
}

func TestMaxSessions(t *testing.T) {
	dir := t.TempDir()
	m := NewManager(dir)
	var ids []string
	for i := 0; i < MaxSessionsPerUser; i++ {
		s, err := m.Create("alice", dir)
		if err != nil {
			t.Fatalf("create %d: %v", i, err)
		}
		ids = append(ids, s.ID)
	}
	if _, err := m.Create("alice", dir); err != ErrTooMany {
		t.Fatalf("attendu ErrTooMany, recu: %v", err)
	}
	for _, id := range ids {
		m.Delete("alice", id)
	}
	// Apres nettoyage, la creation refonctionne.
	s, err := m.Create("alice", dir)
	if err != nil {
		t.Fatalf("create apres nettoyage: %v", err)
	}
	m.Delete("alice", s.ID)
}

func TestMaxSessionsConcurrent(t *testing.T) {
	dir := t.TempDir()
	m := NewManager(dir)
	const n = 3 * MaxSessionsPerUser
	var wg sync.WaitGroup
	var okCount atomic.Int32
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := m.Create("alice", dir); err == nil {
				okCount.Add(1)
			}
		}()
	}
	wg.Wait()
	if got := okCount.Load(); got != int32(MaxSessionsPerUser) {
		t.Fatalf("sessions creees = %d, attendu %d", got, MaxSessionsPerUser)
	}
	if got := len(m.List("alice")); got != MaxSessionsPerUser {
		t.Fatalf("sessions listees = %d, attendu %d", got, MaxSessionsPerUser)
	}
	m.Close()
}
