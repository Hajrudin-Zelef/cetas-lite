package web

import (
	"net/http/httptest"
	"testing"
	"time"
)

const vaultTestPassword = "CoffreFort!2026Test"

func newVaultTestServer(t *testing.T) (string, string) {
	t.Helper()
	s := newTestServer(t, true)
	ts := httptest.NewServer(s.Handler())
	t.Cleanup(ts.Close)
	registerUser(t, ts.URL, "vaultuser", "secret1234")
	tok := loginUser(t, ts.URL, "vaultuser", "secret1234")
	return ts.URL, tok
}

func TestVaultLifecycle(t *testing.T) {
	base, tok := newVaultTestServer(t)

	// Pas de coffre au départ.
	if code, out := doAuthed(t, "GET", base+"/api/vault/status", tok, nil); code != 200 || out["exists"] != false {
		t.Fatalf("status initial = %d %v", code, out)
	}

	// Mot de passe faible refusé.
	if code, _ := doAuthed(t, "POST", base+"/api/vault/init", tok, map[string]string{"password": "court"}); code != 400 {
		t.Fatalf("init faible = %d, attendu 400", code)
	}

	// Création.
	if code, _ := doAuthed(t, "POST", base+"/api/vault/init", tok, map[string]string{"password": vaultTestPassword}); code != 201 {
		t.Fatalf("init = %d, attendu 201", code)
	}
	// Double création → 409.
	if code, _ := doAuthed(t, "POST", base+"/api/vault/init", tok, map[string]string{"password": vaultTestPassword}); code != 409 {
		t.Fatalf("init bis = %d, attendu 409", code)
	}

	// Mauvais mot de passe → 401.
	if code, _ := doAuthed(t, "POST", base+"/api/vault/unlock", tok, map[string]string{"password": "Mauvais!00000000"}); code != 401 {
		t.Fatalf("unlock faux = %d, attendu 401", code)
	}

	// Déverrouillage.
	if code, out := doAuthed(t, "POST", base+"/api/vault/unlock", tok, map[string]string{"password": vaultTestPassword}); code != 200 || out["unlocked"] != true {
		t.Fatalf("unlock = %d %v", code, out)
	}

	// Liste vide.
	if code, out := doAuthed(t, "GET", base+"/api/vault/entries", tok, nil); code != 200 {
		t.Fatalf("entries = %d", code)
	} else if keys, _ := out["keys"].([]any); len(keys) != 0 {
		t.Fatalf("entries non vide: %v", keys)
	}

	// Set / Get / Delete.
	if code, _ := doAuthed(t, "PUT", base+"/api/vault/entries/github", tok, map[string]any{"value": "ghp_secret"}); code != 200 {
		t.Fatalf("set = %d", code)
	}
	if code, out := doAuthed(t, "GET", base+"/api/vault/entries/github", tok, nil); code != 200 || out["value"] != "ghp_secret" {
		t.Fatalf("get = %d %v", code, out)
	}
	if code, _ := doAuthed(t, "GET", base+"/api/vault/entries/absent", tok, nil); code != 404 {
		t.Fatalf("get absent = %d, attendu 404", code)
	}
	if code, _ := doAuthed(t, "DELETE", base+"/api/vault/entries/github", tok, nil); code != 200 {
		t.Fatalf("delete = %d", code)
	}

	// Verrouillage.
	if code, _ := doAuthed(t, "POST", base+"/api/vault/lock", tok, nil); code != 200 {
		t.Fatalf("lock = %d", code)
	}
	if code, _ := doAuthed(t, "GET", base+"/api/vault/entries", tok, nil); code != 401 {
		t.Fatalf("entries verrouillé = %d, attendu 401", code)
	}
}

func TestVaultChangePassword(t *testing.T) {
	base, tok := newVaultTestServer(t)
	newPW := "NouveauCoffre!2042"

	doAuthed(t, "POST", base+"/api/vault/init", tok, map[string]string{"password": vaultTestPassword})
	doAuthed(t, "POST", base+"/api/vault/unlock", tok, map[string]string{"password": vaultTestPassword})
	doAuthed(t, "PUT", base+"/api/vault/entries/note", tok, map[string]any{"value": "secrete"})

	// Nouveau mot de passe faible refusé.
	if code, _ := doAuthed(t, "POST", base+"/api/vault/change-password", tok,
		map[string]string{"old_password": vaultTestPassword, "new_password": "faible"}); code != 400 {
		t.Fatalf("change faible = %d, attendu 400", code)
	}
	// Ancien mot de passe incorrect.
	if code, _ := doAuthed(t, "POST", base+"/api/vault/change-password", tok,
		map[string]string{"old_password": "Faux!0000000000", "new_password": newPW}); code != 401 {
		t.Fatalf("change faux ancien = %d, attendu 401", code)
	}
	// Changement OK.
	if code, _ := doAuthed(t, "POST", base+"/api/vault/change-password", tok,
		map[string]string{"old_password": vaultTestPassword, "new_password": newPW}); code != 200 {
		t.Fatalf("change = %d", code)
	}
	// Donnée toujours lisible avec le nouveau mot de passe.
	doAuthed(t, "POST", base+"/api/vault/lock", tok, nil)
	if code, _ := doAuthed(t, "POST", base+"/api/vault/unlock", tok, map[string]string{"password": vaultTestPassword}); code != 401 {
		t.Fatalf("ancien mdp après change = %d, attendu 401", code)
	}
	if code, _ := doAuthed(t, "POST", base+"/api/vault/unlock", tok, map[string]string{"password": newPW}); code != 200 {
		t.Fatalf("nouveau mdp = %d", code)
	}
	if code, out := doAuthed(t, "GET", base+"/api/vault/entries/note", tok, nil); code != 200 || out["value"] != "secrete" {
		t.Fatalf("get après change = %d %v", code, out)
	}
}

func TestVaultAutoLock(t *testing.T) {
	s := newTestServer(t, true)
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	registerUser(t, ts.URL, "vaultauto", "secret1234")
	tok := loginUser(t, ts.URL, "vaultauto", "secret1234")

	doAuthed(t, "POST", ts.URL+"/api/vault/init", tok, map[string]string{"password": vaultTestPassword})
	doAuthed(t, "POST", ts.URL+"/api/vault/unlock", tok, map[string]string{"password": vaultTestPassword})

	// Simule 16 minutes d'inactivité.
	s.vault.mu.Lock()
	s.vault.lastActivity = time.Now().Add(-16 * time.Minute)
	s.vault.mu.Unlock()

	if code, _ := doAuthed(t, "GET", ts.URL+"/api/vault/entries", tok, nil); code != 401 {
		t.Fatalf("après inactivité = %d, attendu 401", code)
	}
	// Le mot de passe a été effacé de la mémoire.
	s.vault.mu.Lock()
	wiped := s.vault.password == nil
	s.vault.mu.Unlock()
	if !wiped {
		t.Fatal("le mot de passe aurait dû être effacé après verrouillage auto")
	}
}

func TestVaultRequiresAuth(t *testing.T) {
	s := newTestServer(t, true)
	ts := httptest.NewServer(s.Handler())
	defer ts.Close()
	if code, _ := doAuthed(t, "GET", ts.URL+"/api/vault/status", "", nil); code != 401 {
		t.Fatalf("sans token = %d, attendu 401", code)
	}
}
