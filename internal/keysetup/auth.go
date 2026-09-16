package keysetup

// Mot de passe de protection de l'outil (l'équivalent du "setup password"
// du setup.py d'origine) : il verrouille cetas-keys lui-même, avant même le
// mot de passe maître du coffre. Seule une empreinte salée (scrypt) est
// stockée — jamais le mot de passe en clair, ni ici ni ailleurs.
//
// Fichier : <home>/setup.auth (JSON, 0600, écriture atomique).
// En cas d'oubli : supprimer setup.auth réinitialise la protection
// (nécessite l'accès au répertoire home, comme le coffre lui-même).

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/crypto/scrypt"
)

const (
	authFileName = "setup.auth"
	authVersion  = 1

	// Paramètres scrypt de l'empreinte (volontairement plus légers que ceux
	// du coffre : c'est une porte d'entrée, pas le chiffrement des secrets).
	authScryptN = 32768
	authScryptR = 8
	authScryptP = 1
	authKeyLen  = 32
	authSaltLen = 32
)

type authFile struct {
	Version int    `json:"v"`
	Salt    string `json:"salt"`
	Hash    string `json:"hash"`
}

func authPath(home string) string {
	return filepath.Join(home, authFileName)
}

// AuthExists indique si un mot de passe de protection est déjà défini.
func AuthExists(home string) bool {
	_, err := os.Stat(authPath(home))
	return err == nil
}

// SetAuthPassword définit le mot de passe de protection : ne stocke que
// sel + empreinte scrypt (irréversible), jamais le mot de passe en clair.
func SetAuthPassword(home, password string) error {
	if len(password) < 4 {
		return fmt.Errorf("mot de passe de protection trop court (4 caractères minimum)")
	}
	salt := make([]byte, authSaltLen)
	if _, err := rand.Read(salt); err != nil {
		return fmt.Errorf("génération du sel : %w", err)
	}
	hash, err := scrypt.Key([]byte(password), salt, authScryptN, authScryptR, authScryptP, authKeyLen)
	if err != nil {
		return fmt.Errorf("dérivation : %w", err)
	}
	af := authFile{
		Version: authVersion,
		Salt:    hex.EncodeToString(salt),
		Hash:    hex.EncodeToString(hash),
	}
	raw, err := json.Marshal(af)
	if err != nil {
		return err
	}
	return writeFileAtomic(authPath(home), raw, 0o600)
}

// writeFileAtomic — écriture atomique avec permissions 0600 appliquées
// avant le rename (même sémantique que le coffre et le .env).
func writeFileAtomic(path string, data []byte, perm os.FileMode) error {
	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, ".tmp_auth_")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Chmod(perm); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmpName, path)
}

// VerifyAuthPassword vérifie le mot de passe contre l'empreinte stockée
// (comparaison à temps constant).
func VerifyAuthPassword(home, password string) bool {
	raw, err := os.ReadFile(authPath(home))
	if err != nil {
		return false
	}
	var af authFile
	if err := json.Unmarshal(raw, &af); err != nil {
		return false
	}
	if af.Version != authVersion {
		return false
	}
	salt, err := hex.DecodeString(af.Salt)
	if err != nil {
		return false
	}
	want, err := hex.DecodeString(af.Hash)
	if err != nil {
		return false
	}
	got, err := scrypt.Key([]byte(password), salt, authScryptN, authScryptR, authScryptP, authKeyLen)
	if err != nil {
		return false
	}
	return subtle.ConstantTimeCompare(got, want) == 1
}

// ResetAuth supprime la protection (en cas d'oubli du mot de passe).
func ResetAuth(home string) error {
	err := os.Remove(authPath(home))
	if os.IsNotExist(err) {
		return nil
	}
	return err
}
