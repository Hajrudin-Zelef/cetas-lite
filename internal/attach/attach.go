package attach

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"cetas-lite/internal/docs"
)

const (
	KindImage = "image"
)

var imageExts = map[string]bool{".png": true, ".jpg": true, ".jpeg": true, ".gif": true, ".webp": true}

type Attachment struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Kind    string `json:"kind"`
	Ext     string `json:"ext"`
	Size    int64  `json:"size"`
	Created int64  `json:"created"`
}

type Store struct {
	root string
	max  int64
}

func New(root string, max int64) *Store {
	if max <= 0 {
		max = 20 << 20
	}
	return &Store{root: root, max: max}
}

func (s *Store) MaxBytes() int64 { return s.max }

func (s *Store) Save(user, name string, data []byte) (Attachment, error) {
	if s == nil {
		return Attachment{}, errors.New("stockage de pieces jointes indisponible")
	}
	if int64(len(data)) > s.max {
		return Attachment{}, fmt.Errorf("fichier trop volumineux (max %d Mo)", s.max>>20)
	}
	ext := strings.ToLower(filepath.Ext(name))
	id, err := newID()
	if err != nil {
		return Attachment{}, err
	}
	dir, err := s.userDir(user)
	if err != nil {
		return Attachment{}, err
	}
	a := Attachment{
		ID: id, Name: filepath.Base(name), Kind: KindFor(name), Ext: strings.TrimPrefix(ext, "."),
		Size: int64(len(data)), Created: time.Now().UnixMilli(),
	}
	if err := os.WriteFile(filepath.Join(dir, id+ext), data, 0o600); err != nil {
		return Attachment{}, err
	}
	metaPath, err := s.metaPath(user, id)
	if err != nil {
		_ = os.Remove(filepath.Join(dir, id+ext))
		return Attachment{}, err
	}
	meta, _ := json.Marshal(a)
	if err := os.WriteFile(metaPath, meta, 0o600); err != nil {
		_ = os.Remove(filepath.Join(dir, id+ext))
		return Attachment{}, err
	}
	return a, nil
}

func (s *Store) Get(user, id string) (Attachment, []byte, error) {
	if s == nil {
		return Attachment{}, nil, errors.New("stockage indisponible")
	}
	dir, err := s.userDir(user)
	if err != nil {
		return Attachment{}, nil, err
	}
	if !safeID(id) {
		return Attachment{}, nil, errors.New("identifiant invalide")
	}
	// Les metadonnees vivent dans .meta/ (hors du namespace des donnees) :
	// un fichier .json verrait sinon sa metadonnee id.json ecraser son
	// contenu. Repli sur l'ancien emplacement id.json (pieces existantes).
	raw, err := os.ReadFile(filepath.Join(dir, metaDir, id+".json"))
	if err != nil {
		raw, err = os.ReadFile(filepath.Join(dir, legacyMetaName(id)))
	}
	if err != nil {
		return Attachment{}, nil, errors.New("piece jointe introuvable")
	}
	var a Attachment
	if err := json.Unmarshal(raw, &a); err != nil {
		return Attachment{}, nil, errors.New("metadonnees illisibles")
	}
	data, err := os.ReadFile(filepath.Join(dir, id+"."+a.Ext))
	if err != nil {
		return Attachment{}, nil, errors.New("fichier introuvable")
	}
	return a, data, nil
}

func (s *Store) Delete(user, id string) error {
	if s == nil {
		return errors.New("stockage indisponible")
	}
	dir, err := s.userDir(user)
	if err != nil {
		return err
	}
	if !safeID(id) {
		return errors.New("identifiant invalide")
	}
	a, _, err := s.Get(user, id)
	if err != nil {
		return err
	}
	_ = os.Remove(filepath.Join(dir, metaDir, id+".json"))
	_ = os.Remove(filepath.Join(dir, legacyMetaName(id)))
	_ = os.Remove(filepath.Join(dir, id+"."+a.Ext))
	_ = os.Remove(filepath.Join(dir, id+textSidecarExt))
	_ = os.Remove(filepath.Join(dir, id+thumbSidecarExt))
	return nil
}

// metaDir : sous-dossier des metadonnees, separe des donnees pour qu'une
// piece jointe .json ne voie jamais sa metadonnee ecraser son contenu.
const metaDir = ".meta"

func legacyMetaName(id string) string { return id + ".json" }

// metaPath : chemin de la metadonnee d'une piece jointe (cree le dossier).
func (s *Store) metaPath(user, id string) (string, error) {
	dir, err := s.userDir(user)
	if err != nil {
		return "", err
	}
	mdir := filepath.Join(dir, metaDir)
	if err := os.MkdirAll(mdir, 0o700); err != nil {
		return "", err
	}
	return filepath.Join(mdir, id+".json"), nil
}

// textSidecarExt / thumbSidecarExt : suffixes des caches derives (texte
// extrait, vignette). Jamais servis comme pieces jointes (listes par .json).
const (
	textSidecarExt  = ".txt"
	thumbSidecarExt = ".thumb.jpg"
)

// ReadSidecar renvoie le contenu d'un sidecar derive (fail-open : absent ou
// illisible => ok=false, l'appelant re-genere).
func (s *Store) ReadSidecar(user, id, suffix string) ([]byte, bool) {
	if s == nil || !safeID(id) {
		return nil, false
	}
	dir, err := s.userDir(user)
	if err != nil {
		return nil, false
	}
	b, err := os.ReadFile(filepath.Join(dir, id+suffix))
	if err != nil {
		return nil, false
	}
	return b, true
}

// WriteSidecar persiste un cache derive. Best-effort : une erreur d'ecriture
// n'invalide pas la piece jointe (le prochain acces re-generera).
func (s *Store) WriteSidecar(user, id, suffix string, data []byte) error {
	if s == nil || !safeID(id) {
		return errors.New("identifiant invalide")
	}
	dir, err := s.userDir(user)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, id+suffix), data, 0o600)
}

// CleanOlderThan supprime les pièces jointes (métadonnées + fichier)
// dont la date de création dépasse maxAge. Appelé au démarrage du moteur :
// les fichiers ne sont plus supprimés à l'envoi (le tour agent les lit de
// façon asynchrone et la régénération peut les relire), ce nettoyage évite
// l'accumulation des orphelins.
func (s *Store) CleanOlderThan(maxAge time.Duration) int {
	if s == nil || maxAge <= 0 {
		return 0
	}
	cutoff := time.Now().Add(-maxAge).UnixMilli()
	removed := 0
	users, err := os.ReadDir(s.root)
	if err != nil {
		return 0
	}
	for _, u := range users {
		if !u.IsDir() {
			continue
		}
		dir := filepath.Join(s.root, u.Name())
		removed += cleanDir(dir, filepath.Join(dir, metaDir), cutoff)
		// Ancien emplacement (metadonnees id.json a la racine) : le garde
		// a.ID==base protege les fichiers .json utilisateur.
		removed += cleanDir(dir, dir, cutoff)
	}
	return removed
}

// cleanDir purge un dossier de metadonnees donne (le dossier .meta courant,
// ou l'ancien emplacement : les fichiers *.json a la racine du dossier user,
// uniquement s'ils portent un champ id coherent — un contenu .json
// utilisateur ne doit jamais etre pris pour une metadonnee).
func cleanDir(dir, metaDirPath string, cutoff int64) int {
	removed := 0
	entries, err := os.ReadDir(metaDirPath)
	if err != nil {
		return 0
	}
	for _, en := range entries {
		if en.IsDir() || !strings.HasSuffix(en.Name(), ".json") {
			continue
		}
		name := en.Name()
		raw, err := os.ReadFile(filepath.Join(metaDirPath, name))
		if err != nil {
			continue
		}
		var a Attachment
		base := strings.TrimSuffix(name, ".json")
		if err := json.Unmarshal(raw, &a); err != nil || a.Created >= cutoff {
			continue
		}
		// Metadonnee coherente : id du JSON == nom de fichier.
		if a.ID != base {
			continue
		}
		_ = os.Remove(filepath.Join(metaDirPath, name))
		_ = os.Remove(filepath.Join(dir, base+"."+a.Ext))
		_ = os.Remove(filepath.Join(dir, base+textSidecarExt))
		_ = os.Remove(filepath.Join(dir, base+thumbSidecarExt))
		removed++
	}
	return removed
}

func (s *Store) userDir(user string) (string, error) {
	safe := strings.NewReplacer("/", "_", "\\", "_").Replace(strings.TrimSpace(user))
	if safe == "" || safe == "." || safe == ".." {
		return "", errors.New("utilisateur invalide")
	}
	dir := filepath.Join(s.root, safe)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", err
	}
	return dir, nil
}

func KindFor(name string) string {
	ext := strings.ToLower(filepath.Ext(name))
	if imageExts[ext] {
		return KindImage
	}
	return string(docs.KindOf(name))
}

func IsImage(name string) bool {
	return imageExts[strings.ToLower(filepath.Ext(name))]
}

func newID() (string, error) {
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func safeID(id string) bool {
	if id == "" {
		return false
	}
	for _, r := range id {
		if !(r >= 'a' && r <= 'f' || r >= '0' && r <= '9') {
			return false
		}
	}
	return true
}
