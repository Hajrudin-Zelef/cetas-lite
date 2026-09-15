package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

	"cetas-lite/internal/securevault"
)

// vaultIdleTimeout — verrouillage automatique après inactivité.
const vaultIdleTimeout = 15 * time.Minute

// maxVaultValueBytes — taille maximale d'une valeur (JSON encodé).
const maxVaultValueBytes = 256 * 1024

// vaultSession garde le mot de passe maître en mémoire uniquement, tant que
// le coffre est déverrouillé. Chaque opération repasse par la dérivation
// Scrypt (comme l'implémentation Python d'origine : nouveau sel à chaque
// sauvegarde, donc jamais de réutilisation de nonce). Le mot de passe est
// effacé de la mémoire au verrouillage (manuel ou automatique).
type vaultSession struct {
	mu           sync.Mutex
	path         string
	password     []byte // nil = verrouillé
	lastActivity time.Time
}

func newVaultSession(path string) *vaultSession {
	return &vaultSession{path: path}
}

func (vs *vaultSession) vault() *securevault.Vault {
	return securevault.New(vs.path)
}

// wipe efface le mot de passe de la mémoire.
func (vs *vaultSession) wipe() {
	for i := range vs.password {
		vs.password[i] = 0
	}
	vs.password = nil
}

// passwordIfUnlocked retourne une copie du mot de passe si la session est
// active, en appliquant le verrouillage automatique après inactivité.
func (vs *vaultSession) passwordIfUnlocked() []byte {
	vs.mu.Lock()
	defer vs.mu.Unlock()
	if vs.password == nil {
		return nil
	}
	if time.Since(vs.lastActivity) > vaultIdleTimeout {
		vs.wipe()
		return nil
	}
	vs.lastActivity = time.Now()
	out := make([]byte, len(vs.password))
	copy(out, vs.password)
	return out
}

func (vs *vaultSession) status() (exists, unlocked bool) {
	exists = vs.vault().Exists()
	vs.mu.Lock()
	unlocked = vs.password != nil && time.Since(vs.lastActivity) <= vaultIdleTimeout
	if vs.password != nil && !unlocked {
		vs.wipe()
	}
	vs.mu.Unlock()
	return exists, unlocked
}

// unlock vérifie le mot de passe (avec migration V2/V3 si besoin) et
// conserve une copie en mémoire.
func (vs *vaultSession) unlock(password string) error {
	v := vs.vault()
	if !v.Exists() {
		return securevault.ErrNotFound
	}
	// Load gère V2/V3/V4 et migre vers V4 automatiquement.
	if _, err := v.Load(password); err != nil {
		return err
	}
	vs.mu.Lock()
	vs.wipe()
	vs.password = []byte(password)
	vs.lastActivity = time.Now()
	vs.mu.Unlock()
	// Efface la copie du mot de passe d'origine ? Non : l'appelant possède
	// sa chaîne ; on ne peut pas l'effacer (string immuable en Go).
	return nil
}

func (vs *vaultSession) lock() {
	vs.mu.Lock()
	vs.wipe()
	vs.mu.Unlock()
}

// jsonSize — taille JSON approximative (garde-fou volumétrie).
func jsonSize(v any) int {
	b, err := json.Marshal(v)
	if err != nil {
		return 0
	}
	return len(b)
}

// --- Handlers ---

func (s *Server) handleVaultStatus(w http.ResponseWriter, r *http.Request) {
	exists, unlocked := s.vault.status()
	writeJSON(w, http.StatusOK, map[string]any{"exists": exists, "unlocked": unlocked})
}

func (s *Server) handleVaultInit(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Password string `json:"password"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	v := s.vault.vault()
	if v.Exists() {
		writeError(w, http.StatusConflict, "un coffre existe déjà")
		return
	}
	if ok, msg := securevault.CheckPassword(body.Password); !ok {
		writeError(w, http.StatusBadRequest, "mot de passe trop faible : "+msg)
		return
	}
	if err := v.Save(body.Password, map[string]any{}); err != nil {
		writeError(w, http.StatusInternalServerError, "création du coffre impossible")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]any{"created": true})
}

func (s *Server) handleVaultUnlock(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Password string `json:"password"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	if body.Password == "" {
		writeError(w, http.StatusBadRequest, "mot de passe requis")
		return
	}
	switch err := s.vault.unlock(body.Password); {
	case err == nil:
		writeJSON(w, http.StatusOK, map[string]any{"unlocked": true})
	case errors.Is(err, securevault.ErrNotFound):
		writeError(w, http.StatusNotFound, "aucun coffre — créez-le d'abord")
	case errors.Is(err, securevault.ErrDecrypt):
		writeError(w, http.StatusUnauthorized, "mot de passe incorrect")
	default:
		writeError(w, http.StatusInternalServerError, "déverrouillage impossible")
	}
}

func (s *Server) handleVaultLock(w http.ResponseWriter, r *http.Request) {
	s.vault.lock()
	writeJSON(w, http.StatusOK, map[string]any{"unlocked": false})
}

// vaultPasswordOrLocked récupère le mot de passe de session ou répond 401/404.
func (s *Server) vaultPasswordOrLocked(w http.ResponseWriter, r *http.Request) string {
	pw := s.vault.passwordIfUnlocked()
	if pw == nil {
		if s.vault.vault().Exists() {
			writeError(w, http.StatusUnauthorized, "coffre verrouillé")
		} else {
			writeError(w, http.StatusNotFound, "aucun coffre")
		}
		return ""
	}
	return string(pw)
}

func (s *Server) handleVaultListKeys(w http.ResponseWriter, r *http.Request) {
	pw := s.vaultPasswordOrLocked(w, r)
	if pw == "" {
		return
	}
	keys, err := s.vault.vault().ListKeys(pw)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "lecture du coffre impossible")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"keys": keys})
}

func (s *Server) handleVaultGetEntry(w http.ResponseWriter, r *http.Request) {
	pw := s.vaultPasswordOrLocked(w, r)
	if pw == "" {
		return
	}
	name, err := securevault.SanitizeKey(r.PathValue("key"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "clé invalide")
		return
	}
	data, err := s.vault.vault().Load(pw)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "lecture du coffre impossible")
		return
	}
	val, ok := data[name]
	if !ok {
		writeError(w, http.StatusNotFound, "entrée introuvable")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"key": name, "value": val})
}

func (s *Server) handleVaultSetEntry(w http.ResponseWriter, r *http.Request) {
	pw := s.vaultPasswordOrLocked(w, r)
	if pw == "" {
		return
	}
	name, err := securevault.SanitizeKey(r.PathValue("key"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "clé invalide")
		return
	}
	var body struct {
		Value any `json:"value"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	v := s.vault.vault()
	data, err := v.Load(pw)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "lecture du coffre impossible")
		return
	}
	if data == nil {
		data = map[string]any{}
	}
	data[name] = body.Value
	// Garde-fou taille.
	if n := jsonSize(data); n > 4*1024*1024 {
		writeError(w, http.StatusBadRequest, "coffre trop volumineux (4 Mo max)")
		return
	}
	if err := v.Save(pw, data); err != nil {
		writeError(w, http.StatusInternalServerError, "écriture impossible")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"key": name})
}

func (s *Server) handleVaultDeleteEntry(w http.ResponseWriter, r *http.Request) {
	pw := s.vaultPasswordOrLocked(w, r)
	if pw == "" {
		return
	}
	name, err := securevault.SanitizeKey(r.PathValue("key"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "clé invalide")
		return
	}
	v := s.vault.vault()
	ok, err := v.DeleteKey(pw, name)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "lecture du coffre impossible")
		return
	}
	if !ok {
		writeError(w, http.StatusNotFound, "entrée introuvable")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"deleted": name})
}

func (s *Server) handleVaultChangePassword(w http.ResponseWriter, r *http.Request) {
	pw := s.vaultPasswordOrLocked(w, r)
	if pw == "" {
		return
	}
	var body struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	if body.OldPassword == "" || body.NewPassword == "" {
		writeError(w, http.StatusBadRequest, "anciens et nouveaux mots de passe requis")
		return
	}
	if body.OldPassword != pw {
		writeError(w, http.StatusUnauthorized, "ancien mot de passe incorrect")
		return
	}
	if ok, msg := securevault.CheckPassword(body.NewPassword); !ok {
		writeError(w, http.StatusBadRequest, "nouveau mot de passe trop faible : "+msg)
		return
	}
	v := s.vault.vault()
	data, err := v.Load(body.OldPassword)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "ancien mot de passe incorrect")
		return
	}
	// Rechiffre avec le nouveau (nouveau sel).
	if err := v.Save(body.NewPassword, data); err != nil {
		writeError(w, http.StatusInternalServerError, "changement impossible")
		return
	}
	// Met à jour la session avec le nouveau mot de passe.
	s.vault.mu.Lock()
	s.vault.wipe()
	s.vault.password = []byte(body.NewPassword)
	s.vault.lastActivity = time.Now()
	s.vault.mu.Unlock()
	writeJSON(w, http.StatusOK, map[string]any{"changed": true})
}
