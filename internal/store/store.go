package store

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	bolt "go.etcd.io/bbolt"
)

var (
	bMeta          = []byte("meta")
	bUsers         = []byte("users")
	bSettings      = []byte("settings")
	bConversations = []byte("conversations")
	bArchives      = []byte("archives")
	bSecrets       = []byte("secrets")
)

type Store struct {
	db *bolt.DB
	mu sync.Mutex
}

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
	Created      string `json:"created"`
}

func Open(path string) (*Store, error) {
	if dir := filepath.Dir(path); dir != "" {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return nil, fmt.Errorf("creation du dossier de base: %w", err)
		}
	}
	db, err := bolt.Open(path, 0o600, &bolt.Options{Timeout: 5 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("ouverture bbolt: %w", err)
	}
	s := &Store{db: db}
	err = db.Update(func(tx *bolt.Tx) error {
		for _, b := range [][]byte{bMeta, bUsers, bSettings, bConversations, bSecrets} {
			if _, err := tx.CreateBucketIfNotExists(b); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("initialisation des buckets: %w", err)
	}
	return s, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) GetUser(username string) (*User, bool) {
	var (
		u  User
		ok bool
	)
	_ = s.db.View(func(tx *bolt.Tx) error {
		raw := tx.Bucket(bUsers).Get([]byte(username))
		if raw == nil {
			return nil
		}
		if err := json.Unmarshal(raw, &u); err != nil {
			return nil
		}
		ok = true
		return nil
	})
	if !ok {
		return nil, false
	}
	return &u, true
}

func (s *Store) PutUser(u *User) error {
	raw, err := json.Marshal(u)
	if err != nil {
		return fmt.Errorf("encodage utilisateur: %w", err)
	}
	return s.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bUsers).Put([]byte(u.Username), raw)
	})
}

func (s *Store) CountUsers() (int, error) {
	n := 0
	err := s.db.View(func(tx *bolt.Tx) error {
		n = tx.Bucket(bUsers).Stats().KeyN
		return nil
	})
	return n, err
}

func (s *Store) GetMeta(key string) ([]byte, bool) {
	var out []byte
	_ = s.db.View(func(tx *bolt.Tx) error {
		if v := tx.Bucket(bMeta).Get([]byte(key)); v != nil {
			out = append([]byte(nil), v...)
		}
		return nil
	})
	return out, out != nil
}

func (s *Store) PutMeta(key string, val []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bMeta).Put([]byte(key), val)
	})
}

func (s *Store) GetSecret(name string) ([]byte, bool) {
	var out []byte
	_ = s.db.View(func(tx *bolt.Tx) error {
		if v := tx.Bucket(bSecrets).Get([]byte(name)); v != nil {
			out = append([]byte(nil), v...)
		}
		return nil
	})
	return out, out != nil
}

func (s *Store) PutSecret(name string, val []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bSecrets).Put([]byte(name), val)
	})
}

func (s *Store) DeleteSecret(name string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bSecrets).Delete([]byte(name))
	})
}

func (s *Store) ListSecrets() ([]string, error) {
	var out []string
	err := s.db.View(func(tx *bolt.Tx) error {
		return tx.Bucket(bSecrets).ForEach(func(k, _ []byte) error {
			out = append(out, string(k))
			return nil
		})
	})
	return out, err
}

func (s *Store) GetSetting(user, key string) ([]byte, bool) {
	var out []byte
	_ = s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bSettings).Bucket([]byte(user))
		if b == nil {
			return nil
		}
		if v := b.Get([]byte(key)); v != nil {
			out = append([]byte(nil), v...)
		}
		return nil
	})
	return out, out != nil
}

func (s *Store) PutSetting(user, key string, val []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.Bucket(bSettings).CreateBucketIfNotExists([]byte(user))
		if err != nil {
			return err
		}
		return b.Put([]byte(key), val)
	})
}

func (s *Store) ListConversations(user string) ([]string, error) {
	var out []string
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bConversations).Bucket([]byte(user))
		if b == nil {
			return nil
		}
		return b.ForEach(func(k, _ []byte) error {
			out = append(out, string(k))
			return nil
		})
	})
	return out, err
}

func (s *Store) GetConversation(user, id string) ([]byte, bool) {
	var out []byte
	_ = s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bConversations).Bucket([]byte(user))
		if b == nil {
			return nil
		}
		if v := b.Get([]byte(id)); v != nil {
			out = append([]byte(nil), v...)
		}
		return nil
	})
	return out, out != nil
}

func (s *Store) PutConversation(user, id string, data []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.Bucket(bConversations).CreateBucketIfNotExists([]byte(user))
		if err != nil {
			return err
		}
		return b.Put([]byte(id), data)
	})
}

func (s *Store) DeleteConversation(user, id string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bConversations).Bucket([]byte(user))
		if b == nil {
			return nil
		}
		return b.Delete([]byte(id))
	})
}

func (s *Store) ArchiveConversation(user, id string, data []byte) error {
	ts := time.Now().Format("20060102_150405")
	archiveID := ts + "_" + id
	return s.db.Update(func(tx *bolt.Tx) error {
		top, err := tx.CreateBucketIfNotExists(bArchives)
		if err != nil {
			return err
		}
		b, err := top.CreateBucketIfNotExists([]byte(user))
		if err != nil {
			return err
		}
		return b.Put([]byte(archiveID), data)
	})
}

func (s *Store) ListArchives(user string) ([]string, error) {
	var out []string
	err := s.db.View(func(tx *bolt.Tx) error {
		top := tx.Bucket(bArchives)
		if top == nil {
			return nil
		}
		b := top.Bucket([]byte(user))
		if b == nil {
			return nil
		}
		return b.ForEach(func(k, _ []byte) error {
			out = append(out, string(k))
			return nil
		})
	})
	return out, err
}

func (s *Store) GetArchive(user, id string) ([]byte, bool) {
	var out []byte
	_ = s.db.View(func(tx *bolt.Tx) error {
		top := tx.Bucket(bArchives)
		if top == nil {
			return nil
		}
		b := top.Bucket([]byte(user))
		if b == nil {
			return nil
		}
		if v := b.Get([]byte(id)); v != nil {
			out = append([]byte(nil), v...)
		}
		return nil
	})
	return out, out != nil
}

func (s *Store) DeleteArchive(user, id string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		top := tx.Bucket(bArchives)
		if top == nil {
			return nil
		}
		b := top.Bucket([]byte(user))
		if b == nil {
			return nil
		}
		return b.Delete([]byte(id))
	})
}

func (s *Store) JWTSecret() ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var secret []byte
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bMeta)
		if v := b.Get([]byte("jwt_secret")); len(v) > 0 {
			secret = append([]byte(nil), v...)
			return nil
		}
		buf := make([]byte, 32)
		if _, err := io.ReadFull(rand.Reader, buf); err != nil {
			return fmt.Errorf("generation du secret JWT: %w", err)
		}
		secret = buf
		return b.Put([]byte("jwt_secret"), buf)
	})
	if err != nil {
		return nil, err
	}
	return secret, nil
}
