package keysetup

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"cetas-lite/internal/securevault"
)

// Store — accès au coffre partagé avec l'application (même vault.enc).
// Le mot de passe maître n'est conservé qu'en mémoire le temps des opérations.
type Store struct {
	vault *securevault.Vault
}

// ResolveHome — même résolution que internal/config : CETAS_LITE_HOME,
// sinon <UserConfigDir>/cetas-lite.
func ResolveHome() (string, error) {
	home := strings.TrimSpace(os.Getenv("CETAS_LITE_HOME"))
	if home == "" {
		base, err := os.UserConfigDir()
		if err != nil {
			return "", fmt.Errorf("répertoire de configuration introuvable : %w", err)
		}
		home = filepath.Join(base, "cetas-lite")
	}
	abs, err := filepath.Abs(home)
	if err != nil {
		return "", fmt.Errorf("chemin CETAS_LITE_HOME : %w", err)
	}
	return abs, nil
}

// VaultPath — chemin du coffre (identique à celui du serveur).
func VaultPath(home string) string { return filepath.Join(home, "vault.enc") }

// New ouvre un Store sur le home donné (ne crée rien).
func New(home string) *Store {
	return &Store{vault: securevault.New(VaultPath(home))}
}

// Exists indique si un coffre existe déjà.
func (s *Store) Exists() bool { return s.vault.Exists() }

// Size taille du coffre en octets.
func (s *Store) Size() int64 { return s.vault.Size() }

// Path chemin du fichier coffre.
func (s *Store) Path() string { return s.vault.Path() }

// Load charge et déchiffre le coffre. Erreur si le mot de passe est faux
// ou si le fichier est corrompu.
func (s *Store) Load(password string) (map[string]any, error) {
	data, err := s.vault.Load(password)
	if err != nil {
		return nil, fmt.Errorf("ouverture du coffre impossible : %w", err)
	}
	if data == nil {
		data = map[string]any{}
	}
	return data, nil
}

// save chiffre et écrit le coffre de façon atomique (0600).
func (s *Store) save(password string, data map[string]any) error {
	if err := s.vault.Save(password, data); err != nil {
		return fmt.Errorf("sauvegarde du coffre impossible : %w", err)
	}
	return nil
}

// APIKeys extrait la map api_keys (jamais nil).
func APIKeys(data map[string]any) map[string]string {
	out := map[string]string{}
	raw, _ := data["api_keys"].(map[string]any)
	for k, v := range raw {
		if str, ok := v.(string); ok && strings.TrimSpace(str) != "" {
			out[k] = str
		}
	}
	return out
}

// SetAPIKey enregistre une clé (déjà validée par l'appelant).
func (s *Store) SetAPIKey(password, id, key string) error {
	if _, ok := ByID(id); !ok {
		return fmt.Errorf("provider inconnu : %s", id)
	}
	data, err := s.Load(password)
	if err != nil {
		return err
	}
	raw, _ := data["api_keys"].(map[string]any)
	if raw == nil {
		raw = map[string]any{}
	}
	raw[id] = key
	data["api_keys"] = raw
	return s.save(password, data)
}

// DeleteAPIKey supprime une clé. Retourne false si absente.
func (s *Store) DeleteAPIKey(password, id string) (bool, error) {
	data, err := s.Load(password)
	if err != nil {
		return false, err
	}
	raw, _ := data["api_keys"].(map[string]any)
	if raw == nil {
		return false, nil
	}
	if _, ok := raw[id]; !ok {
		return false, nil
	}
	delete(raw, id)
	data["api_keys"] = raw
	return true, s.save(password, data)
}

// ProxyKey retourne la clé de scellement du .env (32 octets),
// la créant si nécessaire. Stockée en hex dans le coffre.
func (s *Store) ProxyKey(password string) ([]byte, error) {
	data, err := s.Load(password)
	if err != nil {
		return nil, err
	}
	if hexKey, ok := data["proxy_key"].(string); ok {
		raw, err := hex.DecodeString(hexKey)
		if err == nil && len(raw) == 32 {
			return raw, nil
		}
		// proxy_key corrompue : on la régénère (le .env sera ré-exporté).
	}
	raw := make([]byte, 32)
	if _, err := rand.Read(raw); err != nil {
		return nil, fmt.Errorf("aléa indisponible : %w", err)
	}
	data["proxy_key"] = hex.EncodeToString(raw)
	if err := s.save(password, data); err != nil {
		return nil, err
	}
	return raw, nil
}

// Init crée un coffre vide avec le mot de passe maître donné.
func (s *Store) Init(password string) error {
	if s.Exists() {
		return fmt.Errorf("un coffre existe déjà : %s", s.Path())
	}
	if ok, msg := securevault.CheckPassword(password); !ok {
		return fmt.Errorf("mot de passe refusé : %s", msg)
	}
	if err := os.MkdirAll(filepath.Dir(s.Path()), 0o700); err != nil {
		return err
	}
	return s.save(password, map[string]any{"api_keys": map[string]any{}})
}

// ChangePassword recharge avec l'ancien mot de passe et ré-écrit avec le nouveau.
// Le sel est régénéré à l'écriture (jamais de réutilisation de sel/nonce).
func (s *Store) ChangePassword(oldPw, newPw string) error {
	if ok, msg := securevault.CheckPassword(newPw); !ok {
		return fmt.Errorf("nouveau mot de passe refusé : %s", msg)
	}
	data, err := s.Load(oldPw)
	if err != nil {
		return err
	}
	return s.save(newPw, data)
}

// PepperSet indique si CETAS_PEPPER est défini (recommandé en production).
func PepperSet() bool { return len(securevault.Pepper()) > 0 }

// Mask masque une clé pour l'affichage : "sk-or-...XXXX".
// Ne révèle jamais plus de 10 caractères.
func Mask(key string) string {
	if len(key) <= 10 {
		return "****"
	}
	return key[:6] + "..." + key[len(key)-4:]
}

func envOr(name, def string) string {
	if v := strings.TrimSpace(os.Getenv(name)); v != "" {
		return v
	}
	return def
}
