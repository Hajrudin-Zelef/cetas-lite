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
	meta, _ := json.Marshal(a)
	if err := os.WriteFile(filepath.Join(dir, id+".json"), meta, 0o600); err != nil {
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
	raw, err := os.ReadFile(filepath.Join(dir, id+".json"))
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
	_ = os.Remove(filepath.Join(dir, id+".json"))
	_ = os.Remove(filepath.Join(dir, id+"."+a.Ext))
	return nil
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
