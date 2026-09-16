package keysetup

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Seal scelle une clé API avec AES-256-GCM.
// Format byte-identique au setup.py d'origine :
//
//	<name>_key=<iv hex>:<ciphertext hex>   avec AAD = <name>
//
// La clé en clair n'apparaît JAMAIS dans le .env.
func Seal(proxyKey []byte, name, apiKey string) (string, error) {
	if len(proxyKey) != 32 {
		return "", fmt.Errorf("proxy_key invalide (%d octets, 32 attendus)", len(proxyKey))
	}
	block, err := aes.NewCipher(proxyKey)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	iv := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(iv); err != nil {
		return "", fmt.Errorf("aléa indisponible : %w", err)
	}
	ct := gcm.Seal(nil, iv, []byte(apiKey), []byte(name))
	return fmt.Sprintf("%s=%s:%s", EnvName(name), hex.EncodeToString(iv), hex.EncodeToString(ct)), nil
}

// Unseal déchiffre une entrée scellée (contrôle d'intégrité via l'AAD).
// Utilisée par les consommateurs du .env et par les tests.
func Unseal(proxyKey []byte, name, sealed string) (string, error) {
	if len(proxyKey) != 32 {
		return "", fmt.Errorf("proxy_key invalide (%d octets, 32 attendus)", len(proxyKey))
	}
	ivHex, ctHex, ok := strings.Cut(sealed, ":")
	if !ok {
		return "", fmt.Errorf("format scellé invalide (iv:ct attendu)")
	}
	iv, err := hex.DecodeString(strings.TrimSpace(ivHex))
	if err != nil {
		return "", fmt.Errorf("iv hex invalide : %w", err)
	}
	ct, err := hex.DecodeString(strings.TrimSpace(ctHex))
	if err != nil {
		return "", fmt.Errorf("ciphertext hex invalide : %w", err)
	}
	block, err := aes.NewCipher(proxyKey)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	plain, err := gcm.Open(nil, iv, ct, []byte(name))
	if err != nil {
		return "", fmt.Errorf("dé-scellement impossible (clé ou nom altéré) : %w", err)
	}
	return string(plain), nil
}

// BuildEnvLines construit les lignes du .env, triées, à partir des clés en clair.
// Les clés en clair ne sont présentes qu'en mémoire le temps de l'appel.
func BuildEnvLines(apiKeys map[string]string, proxyKey []byte) ([]string, error) {
	lines := make([]string, 0, len(apiKeys))
	for id, key := range apiKeys {
		if strings.TrimSpace(key) == "" {
			continue
		}
		line, err := Seal(proxyKey, id, key)
		if err != nil {
			return nil, fmt.Errorf("scellement %s : %w", id, err)
		}
		lines = append(lines, line)
	}
	sort.Strings(lines)
	return lines, nil
}

// ParseEnvLine découpe une ligne "<nom>_key=<iv>:<ct>" en (nom, valeur scellée).
func ParseEnvLine(line string) (name, sealed string, ok bool) {
	line = strings.TrimSpace(line)
	if line == "" || strings.HasPrefix(line, "#") {
		return "", "", false
	}
	k, v, found := strings.Cut(line, "=")
	if !found || !strings.HasSuffix(k, "_key") || v == "" {
		return "", "", false
	}
	return strings.TrimSuffix(k, "_key"), v, true
}

// WriteEnvFile écrit le .env de façon atomique (0600).
// Sans aucune clé, le fichier est supprimé (comme le setup.py d'origine).
// Retourne le nombre de clés exportées.
func WriteEnvFile(path string, apiKeys map[string]string, proxyKey []byte) (int, error) {
	if len(apiKeys) == 0 {
		if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
			return 0, fmt.Errorf("suppression du .env vide : %w", err)
		}
		return 0, nil
	}
	lines, err := BuildEnvLines(apiKeys, proxyKey)
	if err != nil {
		return 0, err
	}
	data := []byte(strings.Join(lines, "\n") + "\n")
	if err := atomicWriteFile(path, data, 0o600); err != nil {
		return 0, fmt.Errorf("écriture du .env : %w", err)
	}
	return len(lines), nil
}

// atomicWriteFile — écriture atomique : fichier temporaire + fsync + rename,
// permissions 0600 appliquées avant le rename (même sémantique que le coffre).
func atomicWriteFile(path string, data []byte, perm os.FileMode) (err error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(dir, ".tmp_env_")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer func() {
		if err != nil {
			os.Remove(tmpName)
		}
	}()
	if _, err = tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err = tmp.Sync(); err != nil {
		tmp.Close()
		return err
	}
	if err = tmp.Close(); err != nil {
		return err
	}
	if err = os.Chmod(tmpName, perm); err != nil {
		return err
	}
	return os.Rename(tmpName, path)
}
