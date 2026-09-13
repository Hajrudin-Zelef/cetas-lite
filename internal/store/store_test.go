package store

import (
	"bytes"
	"path/filepath"
	"testing"
)

func openTemp(t *testing.T) *Store {
	t.Helper()
	s, err := Open(filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	t.Cleanup(func() { _ = s.Close() })
	return s
}

func TestUserRoundTrip(t *testing.T) {
	s := openTemp(t)
	u := &User{Username: "sam", PasswordHash: "hash", Role: "user", Created: "2026-09-13"}
	if err := s.PutUser(u); err != nil {
		t.Fatal(err)
	}
	got, ok := s.GetUser("sam")
	if !ok {
		t.Fatal("utilisateur introuvable")
	}
	if got.Username != "sam" || got.PasswordHash != "hash" || got.Role != "user" {
		t.Fatalf("utilisateur = %+v", got)
	}
	if _, ok := s.GetUser("inconnu"); ok {
		t.Fatal("utilisateur inconnu ne devrait pas exister")
	}
}

func TestPersistenceAcrossReopen(t *testing.T) {
	path := filepath.Join(t.TempDir(), "persist.db")
	s1, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := s1.PutConversation("sam", "c1", []byte(`{"id":"c1"}`)); err != nil {
		t.Fatal(err)
	}
	if err := s1.Close(); err != nil {
		t.Fatal(err)
	}

	s2, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer s2.Close()
	got, ok := s2.GetConversation("sam", "c1")
	if !ok {
		t.Fatal("conversation perdue apres reopen")
	}
	if !bytes.Equal(got, []byte(`{"id":"c1"}`)) {
		t.Fatalf("conversation = %s", got)
	}
}

func TestConversationsArePerUser(t *testing.T) {
	s := openTemp(t)
	if err := s.PutConversation("sam", "c1", []byte("a")); err != nil {
		t.Fatal(err)
	}
	if err := s.PutConversation("bob", "c2", []byte("b")); err != nil {
		t.Fatal(err)
	}
	sam, _ := s.ListConversations("sam")
	if len(sam) != 1 || sam[0] != "c1" {
		t.Fatalf("conversations sam = %v", sam)
	}
	if _, ok := s.GetConversation("bob", "c1"); ok {
		t.Fatal("bob ne doit pas voir la conversation de sam")
	}
	if err := s.DeleteConversation("sam", "c1"); err != nil {
		t.Fatal(err)
	}
	if _, ok := s.GetConversation("sam", "c1"); ok {
		t.Fatal("conversation non supprimee")
	}
}

func TestJWTSecretStable(t *testing.T) {
	s := openTemp(t)
	a, err := s.JWTSecret()
	if err != nil {
		t.Fatal(err)
	}
	b, err := s.JWTSecret()
	if err != nil {
		t.Fatal(err)
	}
	if len(a) != 32 {
		t.Fatalf("secret de taille %d", len(a))
	}
	if !bytes.Equal(a, b) {
		t.Fatal("le secret JWT doit rester stable")
	}
}

func TestSecretsAndSettings(t *testing.T) {
	s := openTemp(t)
	if err := s.PutSecret("deepseek", []byte("enc")); err != nil {
		t.Fatal(err)
	}
	if v, ok := s.GetSecret("deepseek"); !ok || !bytes.Equal(v, []byte("enc")) {
		t.Fatalf("secret = %q ok=%v", v, ok)
	}
	if err := s.PutSetting("sam", "theme", []byte("ocean")); err != nil {
		t.Fatal(err)
	}
	if v, ok := s.GetSetting("sam", "theme"); !ok || !bytes.Equal(v, []byte("ocean")) {
		t.Fatalf("setting = %q ok=%v", v, ok)
	}
}
