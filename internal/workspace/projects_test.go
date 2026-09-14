package workspace

import (
	"context"
	"testing"
)

// memSecrets est un SecretsStore en mémoire pour les tests.
type memSecrets struct{ m map[string][]byte }

func newMemSecrets() *memSecrets { return &memSecrets{m: map[string][]byte{}} }

func (s *memSecrets) PutSecret(name string, val []byte) error {
	s.m[name] = append([]byte(nil), val...)
	return nil
}
func (s *memSecrets) GetSecret(name string) ([]byte, bool) {
	v, ok := s.m[name]
	return v, ok
}
func (s *memSecrets) DeleteSecret(name string) error {
	delete(s.m, name)
	return nil
}

// xorEnc est un Encryptor réversible pour les tests (NON sécurisé).
type xorEnc struct{}

func (xorEnc) Encrypt(plain, aad []byte) ([]byte, error) {
	out := append([]byte(nil), plain...)
	for i := range out {
		out[i] ^= 0x5A
	}
	return out, nil
}
func (xorEnc) Decrypt(data, aad []byte) ([]byte, error) { return xorEnc{}.Encrypt(data, aad) }

func testManager(t *testing.T) *Manager {
	t.Helper()
	return New(t.TempDir(), newMemSecrets(), xorEnc{})
}

func TestCreateGetListDeleteLocal(t *testing.T) {
	m := testManager(t)
	p, err := m.CreateLocal("Mon projet")
	if err != nil {
		t.Fatal(err)
	}
	if p.Mode != ModeLocal || p.Name != "Mon projet" || p.ID == "" {
		t.Fatalf("projet: %+v", p)
	}
	if _, err := m.CreateLocal(""); err == nil {
		t.Fatal("nom vide accepté")
	}
	got, err := m.Get(p.ID)
	if err != nil || got.Name != p.Name {
		t.Fatalf("get: %+v %v", got, err)
	}
	list, err := m.List()
	if err != nil || len(list) != 1 {
		t.Fatalf("list: %v %v", list, err)
	}
	// FS local : écriture puis relecture.
	fsys, err := m.OpenFS(p.ID)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	if err := fsys.WriteFile(ctx, "hello.txt", []byte("salut"), 0o644); err != nil {
		t.Fatal(err)
	}
	b, err := fsys.ReadFile(ctx, "hello.txt")
	if err != nil || string(b) != "salut" {
		t.Fatalf("roundtrip: %q %v", b, err)
	}
	if err := m.Delete(p.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := m.Get(p.ID); err == nil {
		t.Fatal("projet supprimé toujours lisible")
	}
	if list, _ := m.List(); len(list) != 0 {
		t.Fatal("liste non vide après suppression")
	}
}

func TestCreateSFTPValidation(t *testing.T) {
	m := testManager(t)
	creds := SFTPCredentials{Password: "pw"}
	if _, err := m.CreateSFTP("srv", "", 22, "u", "/srv/p", creds, nil); err == nil {
		t.Fatal("hôte vide accepté")
	}
	if _, err := m.CreateSFTP("srv", "h", 22, "u", "rel/path", creds, nil); err == nil {
		t.Fatal("chemin relatif accepté")
	}
	p, err := m.CreateSFTP("srv", "example.com", 0, "deploy", "/srv/app", creds, []byte("hostkey"))
	if err != nil {
		t.Fatal(err)
	}
	if p.Port != 22 || p.Mode != ModeSFTP {
		t.Fatalf("projet: %+v", p)
	}
	// Les secrets sont bien stockés (chiffrés).
	stored, ok := m.secrets.GetSecret("sftp:" + p.ID)
	if !ok || len(stored) == 0 {
		t.Fatal("secrets SFTP non stockés")
	}
	gotCreds, err := m.loadSecrets(p.ID)
	if err != nil || gotCreds.Password != "pw" {
		t.Fatalf("secrets: %+v %v", gotCreds, err)
	}
	if err := m.Delete(p.ID); err != nil {
		t.Fatal(err)
	}
	if _, ok := m.secrets.GetSecret("sftp:" + p.ID); ok {
		t.Fatal("secrets non supprimés")
	}
}

func TestGetUnknown(t *testing.T) {
	m := testManager(t)
	if _, err := m.Get("nope"); err == nil {
		t.Fatal("projet inconnu accepté")
	}
	if _, err := m.Get("../evil"); err == nil {
		t.Fatal("traversal accepté")
	}
	if _, err := m.OpenFS("nope"); err == nil {
		t.Fatal("OpenFS inconnu accepté")
	}
}
