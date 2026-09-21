// Package workspace gère les projets de l'agent : un projet est soit un
// dossier local alimenté par upload, soit un dossier distant accédé
// directement en SFTP. Chaque projet expose un vfs.FS confiné que
// l'agent utilise pour lire et modifier les fichiers.
package workspace

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"cetas-lite/internal/vfs"
)

// Modes de projet.
const (
	ModeLocal = "local"
	ModeSFTP  = "sftp"
)

// SFTPCredentials : secrets SFTP (stockés chiffrés dans le store).
type SFTPCredentials struct {
	Password      string `json:"password,omitempty"`
	PrivateKey    string `json:"private_key,omitempty"`
	KeyPassphrase string `json:"key_passphrase,omitempty"`
}

// Project décrit un projet.
type Project struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Mode      string    `json:"mode"`
	CreatedAt time.Time `json:"created_at"`
	// SFTP (mode sftp uniquement ; aucun secret ici).
	Host       string `json:"host,omitempty"`
	Port       int    `json:"port,omitempty"`
	User       string `json:"user,omitempty"`
	RemotePath string `json:"remote_path,omitempty"`
	HostKey    string `json:"host_key,omitempty"` // clé hôte TOFU (base64)
}

// SecretsStore abstrait le stockage chiffré des secrets.
type SecretsStore interface {
	PutSecret(name string, val []byte) error
	GetSecret(name string) ([]byte, bool)
	DeleteSecret(name string) error
}

// Encryptor chiffre les secrets avant stockage.
type Encryptor interface {
	Encrypt(plain, aad []byte) ([]byte, error)
	Decrypt(data, aad []byte) ([]byte, error)
}

// Manager gère les projets sur disque.
type Manager struct {
	home    string
	secrets SecretsStore
	enc     Encryptor

	mu sync.Mutex
	// FS ouverts mis en cache (surtout SFTP : connexion persistante).
	open map[string]vfs.FS
}

// New crée un Manager. home = $CETAS_LITE_HOME.
func New(home string, secrets SecretsStore, enc Encryptor) *Manager {
	return &Manager{home: home, secrets: secrets, enc: enc, open: map[string]vfs.FS{}}
}

func (m *Manager) dir() string { return filepath.Join(m.home, "projects") }

func (m *Manager) projectDir(id string) string { return filepath.Join(m.dir(), id) }

func newID() string {
	var b [8]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}

// validateName vérifie le nom d'un projet.
func validateName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" || len(name) > 80 {
		return errors.New("nom de projet invalide")
	}
	return nil
}

// CreateLocal crée un projet local (dossier d'upload).
func (m *Manager) CreateLocal(name string) (*Project, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	p := &Project{ID: newID(), Name: strings.TrimSpace(name), Mode: ModeLocal, CreatedAt: time.Now()}
	if err := os.MkdirAll(filepath.Join(m.projectDir(p.ID), "files"), 0o700); err != nil {
		return nil, err
	}
	if err := m.saveLocked(p); err != nil {
		return nil, err
	}
	return p, nil
}

// CreateSFTP crée un projet distant. Les secrets sont chiffrés ; la clé
// hôte peut être vide au premier contact (TOFU via TestSFTP).
func (m *Manager) CreateSFTP(name, host string, port int, user, remotePath string, creds SFTPCredentials, hostKey []byte) (*Project, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}
	host = strings.TrimSpace(host)
	user = strings.TrimSpace(user)
	remotePath = strings.TrimSpace(remotePath)
	if host == "" || user == "" || remotePath == "" {
		return nil, errors.New("hote, utilisateur et dossier distant requis")
	}
	if port <= 0 {
		port = 22
	}
	if !strings.HasPrefix(remotePath, "/") {
		return nil, errors.New("le dossier distant doit etre un chemin absolu")
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	p := &Project{
		ID: newID(), Name: strings.TrimSpace(name), Mode: ModeSFTP,
		CreatedAt: time.Now(), Host: host, Port: port, User: user,
		RemotePath: remotePath,
	}
	if len(hostKey) > 0 {
		p.HostKey = b64(hostKey)
	}
	if err := os.MkdirAll(m.projectDir(p.ID), 0o700); err != nil {
		return nil, err
	}
	if err := m.saveSecrets(p.ID, creds); err != nil {
		return nil, err
	}
	if err := m.saveLocked(p); err != nil {
		return nil, err
	}
	return p, nil
}

// TestSFTP teste la connexion et renvoie l'empreinte de la clé hôte si
// elle est inconnue (premier contact TOFU).
func (m *Manager) TestSFTP(ctx context.Context, host string, port int, user, remotePath string, creds SFTPCredentials, knownHostKey []byte) (fingerprint string, err error) {
	cfg := vfs.SFTPConfig{
		Host: host, Port: port, User: user,
		Password: creds.Password, PrivateKey: []byte(creds.PrivateKey),
		KeyPassphrase: creds.KeyPassphrase, RemotePath: remotePath,
		HostKey: knownHostKey,
	}
	if port <= 0 {
		cfg.Port = 22
	}
	// dialSSH n'est pas exporté : on passe par un FS temporaire dont on
	// ne garde que la phase de connexion.
	tmp := vfs.NewSFTP(cfg, "test")
	defer tmp.Close()
	_, serr := tmp.Stat(ctx, "")
	var uh *vfs.ErrUnknownHostKey
	if errors.As(serr, &uh) {
		return uh.Fingerprint, uh
	}
	return "", serr
}

func (m *Manager) saveLocked(p *Project) error {
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(m.projectDir(p.ID), "project.json"), b, 0o600)
}

func (m *Manager) saveSecrets(id string, creds SFTPCredentials) error {
	if m.enc == nil || m.secrets == nil {
		return errors.New("coffre indisponible")
	}
	b, err := json.Marshal(creds)
	if err != nil {
		return err
	}
	ct, err := m.enc.Encrypt(b, []byte("sftp:"+id))
	if err != nil {
		return err
	}
	return m.secrets.PutSecret("sftp:"+id, ct)
}

func (m *Manager) loadSecrets(id string) (SFTPCredentials, error) {
	var creds SFTPCredentials
	if m.enc == nil || m.secrets == nil {
		return creds, errors.New("coffre indisponible")
	}
	ct, ok := m.secrets.GetSecret("sftp:" + id)
	if !ok {
		return creds, errors.New("secrets introuvables")
	}
	b, err := m.enc.Decrypt(ct, []byte("sftp:"+id))
	if err != nil {
		return creds, err
	}
	if err := json.Unmarshal(b, &creds); err != nil {
		return creds, err
	}
	return creds, nil
}

// Get retourne un projet par ID.
func (m *Manager) Get(id string) (*Project, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.getLocked(id)
}

func (m *Manager) getLocked(id string) (*Project, error) {
	if strings.Contains(id, "/") || strings.Contains(id, "\\") || id == "" {
		return nil, errors.New("projet introuvable")
	}
	b, err := os.ReadFile(filepath.Join(m.projectDir(id), "project.json"))
	if err != nil {
		return nil, errors.New("projet introuvable")
	}
	var p Project
	if err := json.Unmarshal(b, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// List retourne tous les projets (triés par nom).
func (m *Manager) List() ([]*Project, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	des, err := os.ReadDir(m.dir())
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var out []*Project
	for _, d := range des {
		if !d.IsDir() {
			continue
		}
		if p, err := m.getLocked(d.Name()); err == nil {
			// Ne jamais exposer les secrets : HostKey reste (non secret).
			out = append(out, p)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}

// Delete supprime un projet, ses fichiers locaux et ses secrets.
func (m *Manager) Delete(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	p, err := m.getLocked(id)
	if err != nil {
		return err
	}
	if fsys, ok := m.open[id]; ok {
		fsys.Close()
		delete(m.open, id)
	}
	if p.Mode == ModeSFTP && m.secrets != nil {
		_ = m.secrets.DeleteSecret("sftp:" + id)
	}
	return os.RemoveAll(m.projectDir(id))
}

// OpenFS retourne le FS confiné du projet (mis en cache).
func (m *Manager) OpenFS(id string) (vfs.FS, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if fsys, ok := m.open[id]; ok {
		return fsys, nil
	}
	p, err := m.getLocked(id)
	if err != nil {
		return nil, err
	}
	var fsys vfs.FS
	switch p.Mode {
	case ModeLocal:
		fsys, err = vfs.NewLocal(filepath.Join(m.projectDir(p.ID), "files"), p.Name)
	case ModeSFTP:
		creds, cerr := m.loadSecrets(p.ID)
		if cerr != nil {
			return nil, cerr
		}
		var hk []byte
		if p.HostKey != "" {
			hk, _ = unb64(p.HostKey)
		}
		fsys = vfs.NewSFTP(vfs.SFTPConfig{
			Host: p.Host, Port: p.Port, User: p.User,
			Password: creds.Password, PrivateKey: []byte(creds.PrivateKey),
			KeyPassphrase: creds.KeyPassphrase, RemotePath: p.RemotePath,
			HostKey: hk,
		}, p.Name+" ("+p.User+"@"+p.Host+")")
	default:
		return nil, errors.New("mode de projet inconnu")
	}
	if err != nil {
		return nil, err
	}
	// Phase 3 : cache de working-set à deux niveaux (RAM 150 Mo LRU + disque,
	// 512 Mo au total). Projets SFTP : miroir disque local sous projectDir
	// (supprimé avec le projet) ; projets locaux : le FS source fait office
	// de niveau disque, aucun miroir redondant. Décorateur transparent : les
	// outils de l'agent comme l'UI passent par OpenFS, aucune interface ne
	// change. Le préchargement (≤ 100 Mo → RAM) est asynchrone et borné.
	var l2dir string
	if p.Mode == ModeSFTP {
		l2dir = filepath.Join(m.projectDir(p.ID), "cache", "l2")
	}
	cached := vfs.NewCachedFS(fsys, vfs.CacheConfig{L2Dir: l2dir})
	cached.StartWarmup()
	m.open[id] = cached
	return cached, nil
}

// SetHostKey enregistre la clé hôte validée (TOFU) et invalide le FS caché.
func (m *Manager) SetHostKey(id string, hostKey []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	p, err := m.getLocked(id)
	if err != nil {
		return err
	}
	p.HostKey = b64(hostKey)
	if fsys, ok := m.open[id]; ok {
		fsys.Close()
		delete(m.open, id)
	}
	return m.saveLocked(p)
}

// CloseAll ferme les FS ouverts (arrêt serveur).
func (m *Manager) CloseAll() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for id, fsys := range m.open {
		fsys.Close()
		delete(m.open, id)
	}
}
