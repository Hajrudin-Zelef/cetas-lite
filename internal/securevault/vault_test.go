package securevault

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

const testPepper = "test-pepper-cetas"
const testPassword = "TestMotDeP@sse!2024"

func withPepper(t *testing.T) {
	t.Helper()
	t.Setenv("CETAS_PEPPER", testPepper)
}

func testData() map[string]any {
	return map[string]any{
		"api_keys": map[string]any{"Groq": "gsk_abc123"},
		"note":     "héllo wörld — ünïcodé ✓",
	}
}

func TestRoundTrip(t *testing.T) {
	withPepper(t)
	v := New(filepath.Join(t.TempDir(), "vault.enc"))
	data := testData()
	if err := v.Save(testPassword, data); err != nil {
		t.Fatalf("save: %v", err)
	}
	if !v.Exists() {
		t.Fatal("le fichier devrait exister")
	}
	// Permissions 0600.
	fi, _ := os.Stat(v.Path())
	if fi.Mode().Perm() != 0o600 {
		t.Fatalf("permissions = %o, attendu 600", fi.Mode().Perm())
	}
	loaded, err := v.Load(testPassword)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if loaded["note"] != data["note"] {
		t.Fatalf("note = %v", loaded["note"])
	}
	keys, err := v.ListKeys(testPassword)
	if err != nil || len(keys) != 2 {
		t.Fatalf("listKeys = %v, %v", keys, err)
	}
}

func TestPythonV4Fixture(t *testing.T) {
	withPepper(t)
	// Coffre V4 généré par l'implémentation Python d'origine.
	v := New("testdata/v4.enc")
	loaded, err := v.Load(testPassword)
	if err != nil {
		t.Fatalf("load fixture python V4: %v", err)
	}
	if loaded["note"] != "héllo wörld — ünïcodé ✓" {
		t.Fatalf("note = %v", loaded["note"])
	}
	m, _ := loaded["api_keys"].(map[string]any)
	if m["Groq"] != "gsk_abc123" {
		t.Fatalf("api_keys = %v", loaded["api_keys"])
	}
}

func copyFixture(t *testing.T, name string) *Vault {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("testdata", name))
	if err != nil {
		t.Fatalf("lecture fixture: %v", err)
	}
	dst := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(dst, raw, 0o600); err != nil {
		t.Fatalf("copie fixture: %v", err)
	}
	return New(dst)
}

func TestMigrateV2(t *testing.T) {
	withPepper(t)
	v := copyFixture(t, "v2.enc")
	loaded, err := v.Load(testPassword)
	if err != nil {
		t.Fatalf("load V2: %v", err)
	}
	if loaded["note"] != "héllo wörld — ünïcodé ✓" {
		t.Fatalf("note = %v", loaded["note"])
	}
	// Migration : le fichier est désormais en V4.
	raw, _ := os.ReadFile(v.Path())
	if string(raw[:5]) != "VAULT" || raw[5] != 4 {
		t.Fatalf("fichier non migré vers V4 (version=%d)", raw[5])
	}
	if _, err := v.Load(testPassword); err != nil {
		t.Fatalf("relecture après migration: %v", err)
	}
}

func TestMigrateV3(t *testing.T) {
	withPepper(t)
	v := copyFixture(t, "v3.enc")
	loaded, err := v.Load(testPassword)
	if err != nil {
		t.Fatalf("load V3: %v", err)
	}
	if loaded["note"] != "héllo wörld — ünïcodé ✓" {
		t.Fatalf("note = %v", loaded["note"])
	}
	raw, _ := os.ReadFile(v.Path())
	if raw[5] != 4 {
		t.Fatalf("fichier non migré vers V4 (version=%d)", raw[5])
	}
}

func TestWrongPassword(t *testing.T) {
	withPepper(t)
	v := New(filepath.Join(t.TempDir(), "vault.enc"))
	if err := v.Save(testPassword, testData()); err != nil {
		t.Fatal(err)
	}
	if _, err := v.Load("mauvais-mot-de-passe-123"); !errors.Is(err, ErrDecrypt) {
		t.Fatalf("attendu ErrDecrypt, obtenu %v", err)
	}
}

func TestTampered(t *testing.T) {
	withPepper(t)
	v := New(filepath.Join(t.TempDir(), "vault.enc"))
	if err := v.Save(testPassword, testData()); err != nil {
		t.Fatal(err)
	}
	raw, _ := os.ReadFile(v.Path())
	raw[len(raw)-1] ^= 0xFF // altère le dernier octet du ciphertext
	if err := os.WriteFile(v.Path(), raw, 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := v.Load(testPassword); !errors.Is(err, ErrDecrypt) {
		t.Fatalf("attendu ErrDecrypt, obtenu %v", err)
	}
}

func TestCorrupt(t *testing.T) {
	withPepper(t)
	p := filepath.Join(t.TempDir(), "vault.enc")
	if err := os.WriteFile(p, []byte("pas un coffre"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := New(p).Load(testPassword); !errors.Is(err, ErrCorrupt) {
		t.Fatalf("attendu ErrCorrupt, obtenu %v", err)
	}
	if _, err := New(filepath.Join(t.TempDir(), "nope.enc")).Load(testPassword); !errors.Is(err, ErrNotFound) {
		t.Fatalf("attendu ErrNotFound, obtenu %v", err)
	}
}

func TestCheckPassword(t *testing.T) {
	if ok, _ := CheckPassword("court1!"); ok {
		t.Fatal("mot de passe court accepté")
	}
	if ok, _ := CheckPassword("que des minuscules sans chiffres"); ok {
		t.Fatal("mot de passe faible accepté")
	}
	if ok, msg := CheckPassword("TestMotDeP@sse!2024"); !ok {
		t.Fatalf("mot de passe fort refusé: %s", msg)
	}
}

func TestGeneratePassword(t *testing.T) {
	p1, err := GeneratePassword(32)
	if err != nil || len(p1) != 32 {
		t.Fatalf("generate: %q, %v", p1, err)
	}
	p2, _ := GeneratePassword(32)
	if p1 == p2 {
		t.Fatal("mots de passe identiques (aléa défaillant ?)")
	}
}

func TestCRUD(t *testing.T) {
	withPepper(t)
	v := New(filepath.Join(t.TempDir(), "vault.enc"))
	if err := v.Save(testPassword, map[string]any{}); err != nil {
		t.Fatal(err)
	}
	if err := v.Set(testPassword, "github", "ghp_secret"); err != nil {
		t.Fatal(err)
	}
	got, err := v.Get(testPassword, "github")
	if err != nil || got != "ghp_secret" {
		t.Fatalf("get = %v, %v", got, err)
	}
	if ok, err := v.DeleteKey(testPassword, "github"); err != nil || !ok {
		t.Fatalf("deleteKey = %v, %v", ok, err)
	}
	if ok, _ := v.DeleteKey(testPassword, "github"); ok {
		t.Fatal("double suppression devrait retourner false")
	}
	if err := v.ChangePassword(testPassword, "NouveauMotDeP@sse!99"); err != nil {
		t.Fatal(err)
	}
	if _, err := v.Load(testPassword); !errors.Is(err, ErrDecrypt) {
		t.Fatalf("ancien mot de passe devrait échouer: %v", err)
	}
	if _, err := v.Load("NouveauMotDeP@sse!99"); err != nil {
		t.Fatalf("nouveau mot de passe: %v", err)
	}
}

func TestBackupRestore(t *testing.T) {
	withPepper(t)
	dir := t.TempDir()
	v := New(filepath.Join(dir, "vault.enc"))
	if err := v.Save(testPassword, testData()); err != nil {
		t.Fatal(err)
	}
	bdir := filepath.Join(dir, "backups")
	dest, err := v.Backup(bdir)
	if err != nil {
		t.Fatalf("backup: %v", err)
	}
	if err := v.Set(testPassword, "extra", "1"); err != nil {
		t.Fatal(err)
	}
	if err := v.Restore(dest); err != nil {
		t.Fatalf("restore: %v", err)
	}
	if got, _ := v.Get(testPassword, "extra"); got != nil {
		t.Fatalf("restore n'a pas annulé: extra = %v", got)
	}
}

func TestDelete(t *testing.T) {
	withPepper(t)
	v := New(filepath.Join(t.TempDir(), "vault.enc"))
	if err := v.Save(testPassword, testData()); err != nil {
		t.Fatal(err)
	}
	if !v.Delete() {
		t.Fatal("delete devrait retourner true")
	}
	if v.Exists() {
		t.Fatal("le fichier devrait avoir disparu")
	}
	if v.Delete() {
		t.Fatal("delete sur absent devrait retourner false")
	}
}

func TestSanitizeKey(t *testing.T) {
	if _, err := SanitizeKey("  "); err == nil {
		t.Fatal("clé vide acceptée")
	}
	k, err := SanitizeKey("  github  ")
	if err != nil || k != "github" {
		t.Fatalf("sanitize = %q, %v", k, err)
	}
}

func TestSaltMatchesDerivation(t *testing.T) {
	withPepper(t)
	v := New(filepath.Join(t.TempDir(), "vault.enc"))
	if err := v.Save(testPassword, testData()); err != nil {
		t.Fatal(err)
	}
	salt, err := v.Salt()
	if err != nil {
		t.Fatalf("salt: %v", err)
	}
	key, err := DeriveKey(testPassword, salt)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := v.LoadWithKey(key); err != nil {
		t.Fatalf("LoadWithKey avec sel lu: %v", err)
	}
}
