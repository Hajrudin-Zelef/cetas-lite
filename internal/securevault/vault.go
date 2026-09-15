// Package securevault — port Go du SecureVault CETAS (Python).
//
// Coffre-fort chiffré dans un fichier unique, format V4 compatible avec
// l'implémentation Python d'origine :
//
//	Chiffrement : AES-256-GCM (AEAD)
//	KDF         : Scrypt N=2^16 r=8 p=1
//	AAD         : version + params KDF + salt (tout l'en-tête est authentifié)
//	Nonce       : dérivé via HKDF-SHA256(sel, info="cetas-nonce-v4") depuis le pepper
//	Pepper      : variable d'environnement CETAS_PEPPER (compatibilité)
//	Mot de passe: normalisé NFC, minimum 12 caractères
//	Écriture    : atomique (fichier temporaire + rename + fsync), mode 0600
//
// Format V4 sur disque :
//
//	MAGIC "VAULT" (5) | version u8 (1) | params KDF (10 : N u64 BE, r u8, p u8)
//	| salt (32) | nonce (12) | ciphertext (AES-256-GCM)
//
// Lecture des anciens formats V2/V3 avec migration automatique vers V4.
package securevault

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"golang.org/x/crypto/hkdf"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/text/unicode/norm"
)

const (
	// Version du format de fichier produit par Save.
	Version = 4

	magic        = "VAULT"
	magicLen     = 5
	saltLen      = 32
	nonceLen     = 12
	keyLen       = 32
	scryptN      = 1 << 16
	scryptR      = 8
	scryptP      = 1
	kdfParamsLen = 10 // N u64 BE (8) + r u8 (1) + p u8 (1)

	// PasswordMinLen — longueur minimale du mot de passe maître.
	PasswordMinLen = 12

	// pepperEnv — variable d'environnement du pepper (compatibilité Python).
	pepperEnv = "CETAS_PEPPER"

	nonceHKDFInfo = "cetas-nonce-v4"
)

var (
	// ErrDecrypt — mot de passe incorrect ou fichier altéré.
	ErrDecrypt = errors.New("securevault: déchiffrement impossible (mot de passe incorrect ou fichier altéré)")
	// ErrCorrupt — fichier corrompu ou format non reconnu.
	ErrCorrupt = errors.New("securevault: fichier corrompu ou non reconnu")
	// ErrNotFound — aucun coffre à ce chemin.
	ErrNotFound = errors.New("securevault: coffre introuvable")
	// ErrWeakPassword — mot de passe maître trop faible.
	ErrWeakPassword = errors.New("securevault: mot de passe trop faible (minimum 12 caractères)")
	// ErrExists — un coffre existe déjà à ce chemin.
	ErrExists = errors.New("securevault: un coffre existe déjà")
)

// Pepper retourne le pepper côté serveur (variable d'environnement).
// Sans lui, un attaquant disposant du .enc peut bruteforcer hors-ligne.
func Pepper() []byte {
	return []byte(os.Getenv(pepperEnv))
}

// normalizePassword — NFC + pepper préfixé (identique au Python).
func normalizePassword(password string) []byte {
	nfc := norm.NFC.Bytes([]byte(password))
	out := make([]byte, 0, len(Pepper())+len(nfc))
	out = append(out, Pepper()...)
	out = append(out, nfc...)
	return out
}

// kdfParamsBytes — 10 octets : N (u64 BE), r (u8), p (u8).
func kdfParamsBytes() []byte {
	b := make([]byte, kdfParamsLen)
	binary.BigEndian.PutUint64(b[0:8], scryptN)
	b[8] = scryptR
	b[9] = scryptP
	return b
}

// deriveKey — Scrypt avec les constantes du format (jamais celles du fichier,
// ce qui borne le coût même face à un fichier forgé).
func deriveKey(password string, salt []byte) ([]byte, error) {
	return scrypt.Key(normalizePassword(password), salt, scryptN, scryptR, scryptP, keyLen)
}

// deriveNonce — nonce déterministe via HKDF-SHA256(sel, info) depuis le pepper.
// Zéro risque de collision nonce/clé, pas de nonce à stocker au hasard.
func deriveNonce(salt, pepper []byte) ([]byte, error) {
	if len(pepper) == 0 {
		pepper = []byte("no-pepper")
	}
	r := hkdf.New(sha256.New, pepper, salt, []byte(nonceHKDFInfo))
	nonce := make([]byte, nonceLen)
	if _, err := io.ReadFull(r, nonce); err != nil {
		return nil, fmt.Errorf("securevault: dérivation du nonce: %w", err)
	}
	return nonce, nil
}

// CheckPassword — robustesse du mot de passe : 12 caractères minimum puis
// heuristique (au moins 3 catégories parmi majuscules/minuscules/chiffres/
// symboles), comme le repli du Python sans zxcvbn.
func CheckPassword(password string) (bool, string) {
	if len([]rune(password)) < PasswordMinLen {
		return false, fmt.Sprintf("minimum %d caractères requis", PasswordMinLen)
	}
	var upper, lower, digit, symbol bool
	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			upper = true
		case unicode.IsLower(r):
			lower = true
		case unicode.IsDigit(r):
			digit = true
		default:
			symbol = true
		}
	}
	cats := 0
	for _, c := range []bool{upper, lower, digit, symbol} {
		if c {
			cats++
		}
	}
	if cats < 3 {
		return false, "le mot de passe doit mélanger majuscules, minuscules, chiffres et/ou symboles"
	}
	return true, "mot de passe accepté"
}

// GeneratePassword — mot de passe aléatoire sans biais (même alphabet que le Python).
func GeneratePassword(length int) (string, error) {
	const alphabet = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789!@#$%&*"
	if length <= 0 {
		length = 32
	}
	out := make([]byte, length)
	max := big.NewInt(int64(len(alphabet)))
	for i := range out {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", fmt.Errorf("securevault: génération aléatoire: %w", err)
		}
		out[i] = alphabet[n.Int64()]
	}
	return string(out), nil
}

// Vault — coffre chiffré dans un fichier unique.
type Vault struct {
	path string
}

// New retourne un Vault lié à un chemin de fichier (le fichier peut ne pas exister).
func New(path string) *Vault {
	return &Vault{path: path}
}

// Path retourne le chemin du fichier coffre.
func (v *Vault) Path() string { return v.path }

// Exists indique si le fichier coffre existe.
func (v *Vault) Exists() bool {
	_, err := os.Stat(v.path)
	return err == nil
}

// Size retourne la taille du fichier (0 s'il n'existe pas).
func (v *Vault) Size() int64 {
	fi, err := os.Stat(v.path)
	if err != nil {
		return 0
	}
	return fi.Size()
}

// validateHeader vérifie magic + taille minimale AVANT tout Scrypt (coûteux).
func validateHeader(raw []byte) error {
	minLen := magicLen + 1 + kdfParamsLen + saltLen + nonceLen + 1
	if len(raw) < minLen {
		return ErrCorrupt
	}
	if string(raw[:magicLen]) != magic {
		return ErrCorrupt
	}
	return nil
}

// DeriveKey dérive la clé AES-256 depuis le mot de passe et le sel.
// Exporté pour les sessions : on dérive une fois au déverrouillage puis on
// réutilise la clé (plutôt que de refaire Scrypt à chaque requête).
func DeriveKey(password string, salt []byte) ([]byte, error) {
	return deriveKey(password, salt)
}

// Salt lit le sel depuis l'en-tête du fichier (sans mot de passe).
// V4 : MAGIC | version | params(10) | salt(32) | ...
// V2/V3 : MAGIC | version | salt(32) | ... (pas de params KDF)
func (v *Vault) Salt() ([]byte, error) {
	raw, err := os.ReadFile(v.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("securevault: lecture: %w", err)
	}
	if err := validateHeader(raw); err != nil {
		return nil, err
	}
	off := magicLen + 1
	if raw[magicLen] == Version {
		off += kdfParamsLen
	}
	salt := make([]byte, saltLen)
	copy(salt, raw[off:off+saltLen])
	return salt, nil
}

// LoadWithKey déchiffre avec une clé pré-dérivée (format V4 uniquement).
func (v *Vault) LoadWithKey(key []byte) (map[string]any, error) {
	raw, err := os.ReadFile(v.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("securevault: lecture: %w", err)
	}
	if err := validateHeader(raw); err != nil {
		return nil, err
	}
	if raw[magicLen] != Version {
		return nil, fmt.Errorf("%w: LoadWithKey exige le format V4", ErrCorrupt)
	}
	return decryptV4WithKey(key, raw[magicLen+1:])
}

// SaveWithKey chiffre avec une clé pré-dérivée (format V4).
func (v *Vault) SaveWithKey(key []byte, data map[string]any) error {
	if data == nil {
		return errors.New("securevault: les données doivent être un dictionnaire")
	}
	if len(key) != keyLen {
		return errors.New("securevault: clé invalide")
	}
	plain, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("securevault: sérialisation: %w", err)
	}
	salt := make([]byte, saltLen)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return fmt.Errorf("securevault: génération du sel: %w", err)
	}
	nonce, err := deriveNonce(salt, Pepper())
	if err != nil {
		return err
	}
	payload, err := sealV4(key, salt, nonce, plain)
	if err != nil {
		return err
	}
	return atomicWriteFile(v.path, payload, 0o600)
}

// sealV4 assemble un payload V4 depuis clé/sel/nonce/plain.
func sealV4(key, salt, nonce, plain []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("securevault: AES: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("securevault: GCM: %w", err)
	}
	params := kdfParamsBytes()
	aad := make([]byte, 0, 1+kdfParamsLen+saltLen)
	aad = append(aad, Version)
	aad = append(aad, params...)
	aad = append(aad, salt...)
	ciphertext := aead.Seal(nil, nonce, plain, aad)

	payload := make([]byte, 0, magicLen+1+kdfParamsLen+saltLen+nonceLen+len(ciphertext))
	payload = append(payload, magic...)
	payload = append(payload, Version)
	payload = append(payload, params...)
	payload = append(payload, salt...)
	payload = append(payload, nonce...)
	payload = append(payload, ciphertext...)
	return payload, nil
}

// decryptV4WithKey déchiffre un corps V4 avec une clé pré-dérivée.
func decryptV4WithKey(key, rest []byte) (map[string]any, error) {
	if len(key) != keyLen {
		return nil, ErrDecrypt
	}
	if len(rest) < kdfParamsLen+saltLen+nonceLen+1 {
		return nil, ErrCorrupt
	}
	params := rest[:kdfParamsLen]
	salt := rest[kdfParamsLen : kdfParamsLen+saltLen]
	nonce := rest[kdfParamsLen+saltLen : kdfParamsLen+saltLen+nonceLen]
	ciphertext := rest[kdfParamsLen+saltLen+nonceLen:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("securevault: AES: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("securevault: GCM: %w", err)
	}
	aad := make([]byte, 0, 1+kdfParamsLen+saltLen)
	aad = append(aad, Version)
	aad = append(aad, params...)
	aad = append(aad, salt...)
	plain, err := aead.Open(nil, nonce, ciphertext, aad)
	if err != nil {
		return nil, ErrDecrypt
	}
	return unmarshalData(plain)
}

// Save chiffre data (JSON) et l'écrit de façon atomique (mode 0600).
func (v *Vault) Save(password string, data map[string]any) error {
	if data == nil {
		return errors.New("securevault: les données doivent être un dictionnaire")
	}
	plain, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("securevault: sérialisation: %w", err)
	}
	salt := make([]byte, saltLen)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return fmt.Errorf("securevault: génération du sel: %w", err)
	}
	nonce, err := deriveNonce(salt, Pepper())
	if err != nil {
		return err
	}
	key, err := deriveKey(password, salt)
	if err != nil {
		return fmt.Errorf("securevault: dérivation de clé: %w", err)
	}
	payload, err := sealV4(key, salt, nonce, plain)
	if err != nil {
		return err
	}
	return atomicWriteFile(v.path, payload, 0o600)
}

// atomicWriteFile — écriture atomique : fichier temporaire dans le même
// répertoire + fsync + rename. Le temporaire est écrasé de zéros en cas d'erreur.
func atomicWriteFile(path string, data []byte, perm os.FileMode) (err error) {
	dir := filepath.Dir(path)
	if dir != "." {
		if mkErr := os.MkdirAll(dir, 0o700); mkErr != nil {
			return fmt.Errorf("securevault: création du répertoire: %w", mkErr)
		}
	}
	tmp, err := os.CreateTemp(dir, ".vault_tmp_*")
	if err != nil {
		return fmt.Errorf("securevault: fichier temporaire: %w", err)
	}
	tmpName := tmp.Name()
	defer func() {
		if err != nil {
			secureErase(tmpName)
		}
	}()
	if _, err = tmp.Write(data); err != nil {
		tmp.Close()
		return fmt.Errorf("securevault: écriture: %w", err)
	}
	if err = tmp.Sync(); err != nil {
		tmp.Close()
		return fmt.Errorf("securevault: fsync: %w", err)
	}
	if err = tmp.Close(); err != nil {
		return fmt.Errorf("securevault: fermeture: %w", err)
	}
	if err = os.Chmod(tmpName, perm); err != nil {
		return fmt.Errorf("securevault: permissions: %w", err)
	}
	if err = os.Rename(tmpName, path); err != nil {
		return fmt.Errorf("securevault: rename atomique: %w", err)
	}
	return nil
}

// secureErase — écrase un fichier de zéros (3 passes, best-effort) puis le supprime.
func secureErase(path string) {
	func() {
		fi, err := os.Stat(path)
		if err != nil {
			return
		}
		f, err := os.OpenFile(path, os.O_WRONLY, 0)
		if err != nil {
			return
		}
		defer f.Close()
		zeros := make([]byte, 64*1024)
		for p := 0; p < 3; p++ {
			if _, err := f.Seek(0, io.SeekStart); err != nil {
				return
			}
			remaining := fi.Size()
			for remaining > 0 {
				n := int64(len(zeros))
				if remaining < n {
					n = remaining
				}
				if _, err := f.Write(zeros[:n]); err != nil {
					return
				}
				remaining -= n
			}
			f.Sync()
		}
	}()
	os.Remove(path)
}

// Load déchiffre le coffre. Les formats V2/V3 sont migrés vers V4 automatiquement.
func (v *Vault) Load(password string) (map[string]any, error) {
	raw, err := os.ReadFile(v.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("securevault: lecture: %w", err)
	}
	if err := validateHeader(raw); err != nil {
		return nil, err
	}
	version := raw[magicLen]
	rest := raw[magicLen+1:]

	switch version {
	case 2:
		return v.loadV2(password, rest)
	case 3:
		return v.loadV3(password, rest)
	case Version:
		return v.loadV4(password, rest)
	default:
		return nil, fmt.Errorf("%w: version %d non supportée", ErrCorrupt, version)
	}
}

// loadV4 — MAGIC | 0x04 | params(10) | salt(32) | nonce(12) | ciphertext.
func (v *Vault) loadV4(password string, rest []byte) (map[string]any, error) {
	if len(rest) < kdfParamsLen+saltLen+nonceLen+1 {
		return nil, ErrCorrupt
	}
	salt := rest[kdfParamsLen : kdfParamsLen+saltLen]
	key, err := deriveKey(password, salt)
	if err != nil {
		return nil, fmt.Errorf("securevault: dérivation de clé: %w", err)
	}
	return decryptV4WithKey(key, rest)
}

// loadV2 — MAGIC | 0x02 | salt(32) | nonce(12) | mac(32) | ciphertext.
// Clé = Scrypt(sel, NFC(mot de passe)) sans pepper ; HMAC-SHA256 puis AES-GCM sans AAD.
func (v *Vault) loadV2(password string, rest []byte) (map[string]any, error) {
	if len(rest) < 32+nonceLen+32+1 {
		return nil, ErrCorrupt
	}
	salt := rest[:32]
	nonce := rest[32 : 32+nonceLen]
	macSum := rest[32+nonceLen : 32+nonceLen+32]
	ciphertext := rest[32+nonceLen+32:]

	nfc := norm.NFC.Bytes([]byte(password))
	key, err := scrypt.Key(nfc, salt, scryptN, scryptR, scryptP, keyLen)
	if err != nil {
		return nil, fmt.Errorf("securevault: dérivation de clé V2: %w", err)
	}
	hmacKey := sha256.Sum256(append(append([]byte{}, key...), "HMAC"...))
	expected := hmac.New(sha256.New, hmacKey[:])
	expected.Write(ciphertext)
	if !hmac.Equal(macSum, expected.Sum(nil)) {
		return nil, ErrDecrypt
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("securevault: AES: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("securevault: GCM: %w", err)
	}
	plain, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, ErrDecrypt
	}
	data, err := unmarshalData(plain)
	if err != nil {
		return nil, err
	}
	// Migration automatique vers V4.
	if err := v.Save(password, data); err != nil {
		return nil, fmt.Errorf("securevault: migration V2→V4: %w", err)
	}
	return data, nil
}

// loadV3 — MAGIC | 0x03 | salt(32) | nonce(12) | ciphertext.
// Clé = Scrypt(sel, pepper + NFC(mot de passe)) ; AAD = 0x03 + salt.
func (v *Vault) loadV3(password string, rest []byte) (map[string]any, error) {
	if len(rest) < saltLen+nonceLen+1 {
		return nil, ErrCorrupt
	}
	salt := rest[:saltLen]
	nonce := rest[saltLen : saltLen+nonceLen]
	ciphertext := rest[saltLen+nonceLen:]

	key, err := deriveKey(password, salt)
	if err != nil {
		return nil, fmt.Errorf("securevault: dérivation de clé V3: %w", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("securevault: AES: %w", err)
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("securevault: GCM: %w", err)
	}
	aad := append([]byte{3}, salt...)
	plain, err := aead.Open(nil, nonce, ciphertext, aad)
	if err != nil {
		return nil, ErrDecrypt
	}
	data, err := unmarshalData(plain)
	if err != nil {
		return nil, err
	}
	// Migration automatique vers V4.
	if err := v.Save(password, data); err != nil {
		return nil, fmt.Errorf("securevault: migration V3→V4: %w", err)
	}
	return data, nil
}

func unmarshalData(plain []byte) (map[string]any, error) {
	var data map[string]any
	dec := json.NewDecoder(bytes.NewReader(plain))
	if err := dec.Decode(&data); err != nil {
		return nil, fmt.Errorf("%w: données illisibles", ErrCorrupt)
	}
	if data == nil {
		data = map[string]any{}
	}
	return data, nil
}

// Get lit une entrée (nil si absente).
func (v *Vault) Get(password, key string) (any, error) {
	data, err := v.Load(password)
	if err != nil {
		return nil, err
	}
	return data[key], nil
}

// Set écrit une entrée.
func (v *Vault) Set(password, key string, value any) error {
	return v.Update(password, func(data map[string]any) map[string]any {
		data[key] = value
		return data
	})
}

// DeleteKey supprime une entrée (false si absente).
func (v *Vault) DeleteKey(password, key string) (bool, error) {
	data, err := v.Load(password)
	if err != nil {
		return false, err
	}
	if _, ok := data[key]; !ok {
		return false, nil
	}
	delete(data, key)
	if err := v.Save(password, data); err != nil {
		return false, err
	}
	return true, nil
}

// ListKeys retourne les noms des entrées.
func (v *Vault) ListKeys(password string) ([]string, error) {
	data, err := v.Load(password)
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	return keys, nil
}

// Update charge, applique updater, puis sauvegarde.
func (v *Vault) Update(password string, updater func(map[string]any) map[string]any) error {
	data, err := v.Load(password)
	if err != nil {
		return err
	}
	if data == nil {
		data = map[string]any{}
	}
	return v.Save(password, updater(data))
}

// ChangePassword recharge avec l'ancien mot de passe et rechiffre avec le nouveau.
func (v *Vault) ChangePassword(oldPassword, newPassword string) error {
	data, err := v.Load(oldPassword)
	if err != nil {
		return err
	}
	return v.Save(newPassword, data)
}

// Delete efface le fichier coffre (écrasé de zéros avant suppression).
func (v *Vault) Delete() bool {
	if !v.Exists() {
		return false
	}
	secureErase(v.path)
	return !v.Exists()
}

// Backup copie le coffre chiffré vers backupDir (horodaté, mode 0600).
func (v *Vault) Backup(backupDir string) (string, error) {
	if !v.Exists() {
		return "", ErrNotFound
	}
	raw, err := os.ReadFile(v.path)
	if err != nil {
		return "", fmt.Errorf("securevault: lecture: %w", err)
	}
	if len(raw) < magicLen || string(raw[:magicLen]) != magic {
		return "", fmt.Errorf("%w: source invalide, backup annulé", ErrCorrupt)
	}
	if err := os.MkdirAll(backupDir, 0o700); err != nil {
		return "", fmt.Errorf("securevault: répertoire de backup: %w", err)
	}
	dest := filepath.Join(backupDir, "vault_backup_"+time.Now().Format("20060102_150405")+".enc")
	if err := atomicWriteFile(dest, raw, 0o600); err != nil {
		return "", fmt.Errorf("securevault: backup: %w", err)
	}
	return dest, nil
}

// Restore remplace le coffre par un backup (vérifie la signature).
func (v *Vault) Restore(backupPath string) error {
	raw, err := os.ReadFile(backupPath)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrNotFound
		}
		return fmt.Errorf("securevault: lecture du backup: %w", err)
	}
	if len(raw) < magicLen || string(raw[:magicLen]) != magic {
		return fmt.Errorf("%w: backup invalide", ErrCorrupt)
	}
	return atomicWriteFile(v.path, raw, 0o600)
}

// SanitizeKey — clé d'entrée : non vide, sans espaces de tête/queue,
// caractères raisonnables.
func SanitizeKey(key string) (string, error) {
	k := strings.TrimSpace(key)
	if k == "" {
		return "", errors.New("securevault: clé vide")
	}
	if len(k) > 128 {
		return "", errors.New("securevault: clé trop longue (128 max)")
	}
	return k, nil
}
