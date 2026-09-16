package keysetup

import (
	"crypto/rand"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func testProxyKey(t *testing.T) []byte {
	t.Helper()
	k := make([]byte, 32)
	if _, err := rand.Read(k); err != nil {
		t.Fatal(err)
	}
	return k
}

func TestSealRoundtrip(t *testing.T) {
	pk := testProxyKey(t)
	secret := "sk-or-vrai-secret-123"
	line, err := Seal(pk, "openrouter", secret)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(line, secret) {
		t.Fatal("la clé en clair fuit dans la ligne scellée")
	}
	if !strings.HasPrefix(line, "openrouter_key=") {
		t.Fatalf("préfixe inattendu : %s", line)
	}
	name, sealed, ok := ParseEnvLine(line)
	if !ok || name != "openrouter" {
		t.Fatalf("ParseEnvLine a échoué : %q %v", name, ok)
	}
	back, err := Unseal(pk, name, sealed)
	if err != nil {
		t.Fatal(err)
	}
	if back != secret {
		t.Fatal("roundtrip scellement corrompu")
	}
}

func TestSealTamperDetected(t *testing.T) {
	pk := testProxyKey(t)
	line, _ := Seal(pk, "deepseek", "sk-secret")
	_, sealed, _ := ParseEnvLine(line)
	// Altération d'un octet du ciphertext.
	tampered := sealed[:len(sealed)-2] + "ff"
	if _, err := Unseal(pk, "deepseek", tampered); err == nil {
		t.Fatal("ciphertext altéré accepté : échec d'intégrité")
	}
	// Mauvais AAD (nom).
	if _, err := Unseal(pk, "openrouter", sealed); err == nil {
		t.Fatal("AAD altéré accepté : échec d'intégrité")
	}
	// Mauvaise proxy_key.
	if _, err := Unseal(testProxyKey(t), "deepseek", sealed); err == nil {
		t.Fatal("mauvaise proxy_key acceptée")
	}
}

func TestSealRejectsBadProxyKey(t *testing.T) {
	if _, err := Seal([]byte("trop-court"), "x", "y"); err == nil {
		t.Fatal("proxy_key de 9 octets acceptée")
	}
}

func TestEnvName(t *testing.T) {
	cases := map[string]string{
		"openrouter":  "openrouter_key",
		"opencode":    "opencode_key",
		"opencode-go": "opencode-go_key",
		"freellmapi":  "freellmapi_key",
	}
	for id, want := range cases {
		if got := EnvName(id); got != want {
			t.Errorf("EnvName(%q) = %q, want %q", id, got, want)
		}
	}
}

func TestBuildEnvLinesSortedAndSealed(t *testing.T) {
	pk := testProxyKey(t)
	keys := map[string]string{"openrouter": "sk-or-1", "deepseek": "sk-ds-2"}
	lines, err := BuildEnvLines(keys, pk)
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 2 {
		t.Fatalf("2 lignes attendues, %d obtenues", len(lines))
	}
	if !strings.HasPrefix(lines[0], "deepseek_key=") || !strings.HasPrefix(lines[1], "openrouter_key=") {
		t.Fatalf("lignes non triées : %v", lines)
	}
	for _, l := range lines {
		if strings.Contains(l, "sk-or-1") || strings.Contains(l, "sk-ds-2") {
			t.Fatal("clé en clair dans le .env")
		}
	}
	// Roundtrip complet via Parse + Unseal.
	for _, l := range lines {
		name, sealed, ok := ParseEnvLine(l)
		if !ok {
			t.Fatalf("ligne illisible : %s", l)
		}
		back, err := Unseal(pk, name, sealed)
		if err != nil || back != keys[name] {
			t.Fatalf("roundtrip %s échoué", name)
		}
	}
}

func TestWriteEnvFileAtomicAndPerms(t *testing.T) {
	pk := testProxyKey(t)
	dir := t.TempDir()
	path := filepath.Join(dir, ".env")
	n, err := WriteEnvFile(path, map[string]string{"openrouter": "sk-or-xyz"}, pk)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("n=%d, want 1", n)
	}
	fi, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if fi.Mode().Perm() != 0o600 {
		t.Fatalf("permissions .env = %o, want 600", fi.Mode().Perm())
	}
	raw, _ := os.ReadFile(path)
	if strings.Contains(string(raw), "sk-or-xyz") {
		t.Fatal("clé en clair écrite dans le .env")
	}
	if !strings.HasSuffix(string(raw), "\n") {
		t.Fatal("newline final manquant")
	}
	// Sans clés : le fichier est supprimé.
	if _, err := WriteEnvFile(path, map[string]string{}, pk); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatal(".env aurait dû être supprimé")
	}
}

func TestParseEnvLineRejects(t *testing.T) {
	for _, l := range []string{"", "# commentaire", "FOO", "nokey=abc", "x_key="} {
		if _, _, ok := ParseEnvLine(l); ok {
			t.Errorf("ligne %q acceptée à tort", l)
		}
	}
}

func TestMask(t *testing.T) {
	if got := Mask("sk-or-abcdef123456"); got != "sk-or-...3456" {
		t.Errorf("Mask = %q", got)
	}
	if Mask("court") != "****" {
		t.Error("les clés courtes doivent être totalement masquées")
	}
}

func TestVaultStoreIntegration(t *testing.T) {
	home := t.TempDir()
	st := New(home)
	if st.Exists() {
		t.Fatal("coffre fantôme")
	}
	pw := "S3cur3-Master-Pass!"
	if err := st.Init(pw); err != nil {
		t.Fatal(err)
	}
	if err := st.Init(pw); err == nil {
		t.Fatal("double init acceptée")
	}
	data, err := st.Load(pw)
	if err != nil {
		t.Fatal(err)
	}
	if len(APIKeys(data)) != 0 {
		t.Fatal("coffre neuf non vide")
	}
	if _, err := st.Load("mauvais-mot-de-passe"); err == nil {
		t.Fatal("mauvais mot de passe accepté")
	}
	if err := st.SetAPIKey(pw, "openrouter", "sk-or-test"); err != nil {
		t.Fatal(err)
	}
	if err := st.SetAPIKey(pw, "inconnu", "x"); err == nil {
		t.Fatal("provider inconnu accepté")
	}
	data, _ = st.Load(pw)
	keys := APIKeys(data)
	if keys["openrouter"] != "sk-or-test" {
		t.Fatalf("clé non persistée : %v", keys)
	}
	pk, err := st.ProxyKey(pw)
	if err != nil || len(pk) != 32 {
		t.Fatal("proxy_key invalide")
	}
	pk2, _ := st.ProxyKey(pw)
	if string(pk) != string(pk2) {
		t.Fatal("proxy_key non stable")
	}
	n, err := WriteEnvFile(filepath.Join(home, ".env"), keys, pk)
	if err != nil || n != 1 {
		t.Fatal(err)
	}
	ok, err := st.DeleteAPIKey(pw, "openrouter")
	if err != nil || !ok {
		t.Fatal("suppression échouée")
	}
	ok, _ = st.DeleteAPIKey(pw, "openrouter")
	if ok {
		t.Fatal("double suppression rapportée comme effective")
	}
	if err := st.ChangePassword(pw, "N3w-S3cur3-Pass!"); err != nil {
		t.Fatal(err)
	}
	if _, err := st.Load(pw); err == nil {
		t.Fatal("ancien mot de passe encore accepté")
	}
	if _, err := st.Load("N3w-S3cur3-Pass!"); err != nil {
		t.Fatal("nouveau mot de passe refusé")
	}
	if err := st.ChangePassword("N3w-S3cur3-Pass!", "faible"); err == nil {
		t.Fatal("mot de passe faible accepté")
	}
}

func TestProviderCatalogComplete(t *testing.T) {
	seen := map[string]bool{}
	for _, p := range Providers {
		// freellmapi est volontairement désactivé sans CETAS_FREELLM_URL.
		if p.ID == "freellmapi" && p.TestURL == "" {
			continue
		}
		if p.ID == "" || p.Label == "" || p.TestURL == "" || p.TestModel == "" {
			t.Errorf("provider incomplet : %+v", p)
		}
		if seen[p.ID] {
			t.Errorf("ID dupliqué : %s", p.ID)
		}
		seen[p.ID] = true
	}
	// Les IDs synchronisables avec l'app.
	for _, id := range []string{"openrouter", "deepseek", "opencode", "opencode-go"} {
		p, ok := ByID(id)
		if !ok || p.AppID != id {
			t.Errorf("provider %s non synchronisable", id)
		}
	}
}

func TestAuthPassword(t *testing.T) {
	home := t.TempDir()
	if AuthExists(home) {
		t.Fatal("aucun mot de passe ne devrait exister")
	}
	if err := SetAuthPassword(home, "Yoro-test-99+"); err != nil {
		t.Fatal(err)
	}
	if !AuthExists(home) {
		t.Fatal("le mot de passe devrait exister")
	}
	if !VerifyAuthPassword(home, "Yoro-test-99+") {
		t.Fatal("le bon mot de passe est rejeté")
	}
	if VerifyAuthPassword(home, "mauvais") {
		t.Fatal("un mauvais mot de passe est accepté")
	}
	// Le mot de passe en clair ne doit apparaître nulle part dans le fichier.
	raw, err := os.ReadFile(filepath.Join(home, "setup.auth"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(raw), "Yoro-test-99+") {
		t.Fatal("le mot de passe en clair fuit dans setup.auth")
	}
	// Deux définitions → sels différents → empreintes différentes.
	home2 := t.TempDir()
	if err := SetAuthPassword(home2, "Yoro-test-99+"); err != nil {
		t.Fatal(err)
	}
	raw2, _ := os.ReadFile(filepath.Join(home2, "setup.auth"))
	if string(raw) == string(raw2) {
		t.Fatal("le sel devrait être aléatoire")
	}
	if err := SetAuthPassword(home, "abc"); err == nil {
		t.Fatal("un mot de passe trop court devrait être refusé")
	}
	if err := ResetAuth(home); err != nil {
		t.Fatal(err)
	}
	if AuthExists(home) {
		t.Fatal("la protection devrait être réinitialisée")
	}
}

// TestOpenCodeSessionHeader — depuis 2026-09-05 la gateway OpenCode exige
// x-opencode-session (sinon HTTP 400) ; la validation l'envoie donc.
func TestOpenCodeSessionHeader(t *testing.T) {
	for _, id := range []string{"opencode", "opencode-go"} {
		p, ok := ByID(id)
		if !ok {
			t.Fatalf("provider %s introuvable", id)
		}
		if v := p.Headers["x-opencode-session"]; v == "" {
			t.Errorf("provider %s : x-opencode-session manquant", id)
		}
	}
}
