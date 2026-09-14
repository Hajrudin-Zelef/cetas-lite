package backup

import (
	"archive/tar"
	"compress/gzip"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	home := t.TempDir()
	db := filepath.Join(home, DBName)
	mcp := filepath.Join(home, MCPName)
	if err := os.WriteFile(db, []byte("V1"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(mcp, []byte(`{"mcp":{}}`), 0o600); err != nil {
		t.Fatal(err)
	}
	bundle := filepath.Join(t.TempDir(), "backup.tar.gz")
	if err := Create(db, mcp, bundle); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(db, []byte("V2-corrompu"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := Restore(bundle, home); err != nil {
		t.Fatal(err)
	}
	got, _ := os.ReadFile(db)
	if string(got) != "V1" {
		t.Fatalf("db restauree = %q", got)
	}
	gotMCP, _ := os.ReadFile(mcp)
	if string(gotMCP) != `{"mcp":{}}` {
		t.Fatalf("mcp restaure = %q", gotMCP)
	}
	matches, _ := filepath.Glob(db + ".bak-*")
	if len(matches) != 1 {
		t.Fatalf("copie de surete attendue, got %v", matches)
	}
}

func TestCreateMissingDB(t *testing.T) {
	if err := Create(filepath.Join(t.TempDir(), "absent.db"), "", filepath.Join(t.TempDir(), "b.tar.gz")); err == nil {
		t.Fatal("erreur attendue si la base est absente")
	}
}

func TestRestoreRejectsTraversalAndNoDB(t *testing.T) {
	home := t.TempDir()
	outside := filepath.Join(filepath.Dir(home), "evil.txt")
	_ = os.Remove(outside)

	bundle := filepath.Join(t.TempDir(), "b.tar.gz")
	f, err := os.Create(bundle)
	if err != nil {
		t.Fatal(err)
	}
	gz := gzip.NewWriter(f)
	tw := tar.NewWriter(gz)
	write := func(name, body string) {
		if err := tw.WriteHeader(&tar.Header{Name: name, Mode: 0o600, Size: int64(len(body))}); err != nil {
			t.Fatal(err)
		}
		if _, err := tw.Write([]byte(body)); err != nil {
			t.Fatal(err)
		}
	}
	write("../evil.txt", "x")
	_ = tw.Close()
	_ = gz.Close()
	_ = f.Close()

	if err := Restore(bundle, home); err == nil || !strings.Contains(err.Error(), DBName) {
		t.Fatalf("erreur 'sans %s' attendue, got %v", DBName, err)
	}
	if _, err := os.Stat(outside); err == nil {
		t.Fatal("fichier hors home cree (traversee)")
	}
}
