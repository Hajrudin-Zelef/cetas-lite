// Package cache : cache exact SQLite des réponses provider (itération 5a).
//
// Une entrée est adressée par le SHA256 d'une forme canonique JSON de la
// requête. Le cache n'est actif que pour les requêtes déterministes :
// température 0, sans outils, sans paramètres supplémentaires (ex. recherche
// web native du provider). Toute autre requête contourne le cache
// (ni lecture ni écriture).
package cache

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"
	"time"

	_ "modernc.org/sqlite"
)

// Usage : tokens de la réponse mise en cache (informatif : ces tokens
// n'ont pas été réellement consommés lors d'un hit).
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Entry : réponse complète rejouable sans appeler le provider.
type Entry struct {
	Content   string `json:"content"`
	Reasoning string `json:"reasoning"`
	Usage     Usage  `json:"usage"`
}

// Eligible reporte la porte d'activation : température 0 (déterministe),
// aucun outil, aucun paramètre supplémentaire.
func Eligible(temperature float64, numTools int, hasExtra bool) bool {
	return temperature == 0 && numTools == 0 && !hasExtra
}

// KeyOf rend le SHA256 hexadécimal de la forme canonique, ou "" si la
// requête n'est pas éligible (le cache est alors contourné).
func KeyOf(canonical []byte, temperature float64, numTools int, hasExtra bool) string {
	if !Eligible(temperature, numTools, hasExtra) {
		return ""
	}
	sum := sha256.Sum256(canonical)
	return hex.EncodeToString(sum[:])
}

// Cache : cache exact persistant (SQLite), compteurs hits/misses en mémoire.
type Cache struct {
	db         *sql.DB
	maxEntries int
	hits       atomic.Uint64
	misses     atomic.Uint64
}

// Open ouvre (ou crée) le cache SQLite à path. maxEntries <= 0 : sans borne.
func Open(path string, maxEntries int) (*Cache, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, fmt.Errorf("cache: mkdir: %w", err)
	}
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("cache: open: %w", err)
	}
	for _, pr := range []string{
		`PRAGMA journal_mode=WAL`,
		`PRAGMA synchronous=NORMAL`,
		`PRAGMA busy_timeout=5000`,
	} {
		if _, err := db.Exec(pr); err != nil {
			_ = db.Close()
			return nil, fmt.Errorf("cache: pragma: %w", err)
		}
	}
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS entries (
		"key" TEXT PRIMARY KEY,
		payload TEXT NOT NULL,
		created_at INTEGER NOT NULL
	)`); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("cache: schema: %w", err)
	}
	if _, err := db.Exec(`CREATE INDEX IF NOT EXISTS idx_entries_created ON entries(created_at)`); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("cache: index: %w", err)
	}
	db.SetMaxOpenConns(1)
	return &Cache{db: db, maxEntries: maxEntries}, nil
}

// Get rend l'entrée associée à key. Absente ou illisible => miss.
func (c *Cache) Get(key string) (Entry, bool) {
	var payload string
	if err := c.db.QueryRow(`SELECT payload FROM entries WHERE "key"=?`, key).Scan(&payload); err != nil {
		c.misses.Add(1)
		return Entry{}, false
	}
	var e Entry
	if err := json.Unmarshal([]byte(payload), &e); err != nil {
		c.misses.Add(1)
		return Entry{}, false
	}
	c.hits.Add(1)
	return e, true
}

// Set mémorise la réponse sous key (écrase l'existante), puis élague
// les plus anciennes au-delà de maxEntries.
func (c *Cache) Set(key string, e Entry) error {
	payload, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("cache: marshal: %w", err)
	}
	_, err = c.db.Exec(`INSERT INTO entries("key", payload, created_at) VALUES(?,?,?)
		ON CONFLICT("key") DO UPDATE SET payload=excluded.payload, created_at=excluded.created_at`,
		key, string(payload), time.Now().Unix())
	if err != nil {
		return fmt.Errorf("cache: insert: %w", err)
	}
	return c.trim()
}

func (c *Cache) trim() error {
	if c.maxEntries <= 0 {
		return nil
	}
	_, err := c.db.Exec(`DELETE FROM entries WHERE "key" NOT IN
		(SELECT "key" FROM entries ORDER BY created_at DESC, "key" DESC LIMIT ?)`, c.maxEntries)
	if err != nil {
		return fmt.Errorf("cache: trim: %w", err)
	}
	return nil
}

// Stats rend hits, misses (depuis l'ouverture) et le nombre d'entrées.
func (c *Cache) Stats() (hits, misses uint64, entries int) {
	_ = c.db.QueryRow(`SELECT COUNT(*) FROM entries`).Scan(&entries)
	return c.hits.Load(), c.misses.Load(), entries
}

// Close ferme la base.
func (c *Cache) Close() error { return c.db.Close() }
