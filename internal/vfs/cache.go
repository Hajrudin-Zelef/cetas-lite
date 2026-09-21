// ============================================================================
// Cache de working-set à deux niveaux (phase 3, volet workspaces).
//
// CachedFS décore un vfs.FS sans changer son interface : transparent pour les
// outils de l'agent comme pour l'UI (tous passent par OpenFS).
//
//   - Niveau 1 (chaud) : RAM, 150 Mo max, éviction LRU. Les fichiers lus ou
//     modifiés y sont chargés ; comptabilité exacte en octets.
//   - Niveau 2 (tiède) : miroir sur le disque serveur, réservé aux FS
//     distants (SFTP) — le contenu distant est rapatrié une fois puis servi
//     depuis RAM/disque. Projets locaux : le FS lui-même est le niveau
//     disque, aucun miroir redondant. Plafond global RAM+disque : 512 Mo.
//   - Écriture : write-back synchronisé à chaque fin de modification
//     (WriteFile met à jour la source, puis L1 et L2).
//   - F6.4 : validation mtime+taille via Stat à chaque lecture — capte les
//     modifications faites hors outils suivis (outil Bash, édition distante).
//   - F6.5 : garde-fou disque — le miroir n'est peuplé que s'il reste assez
//     d'espace libre ; éviction LRU au-delà du plafond.
//
// Le préchargement (dossier ≤ 100 Mo → RAM à l'ouverture) est asynchrone et
// borné (30 s) : il ne ralentit jamais l'ouverture du projet.
// ============================================================================
package vfs

import (
	"container/list"
	"context"
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Plafonds du cache. Déclarés en variables (non constantes) pour que les
// tests du package puissent les réduire temporairement ; en production ce
// sont les valeurs ci-dessous qui s'appliquent.
var (
	// CacheL1MaxBytes : plafond RAM du niveau 1 (150 Mo).
	CacheL1MaxBytes int64 = 150 << 20
	// CacheTotalMaxBytes : plafond global RAM + disque (512 Mo).
	CacheTotalMaxBytes int64 = 512 << 20
	// CachePreloadMaxBytes : seuil de préchargement (dossier ≤ 100 Mo → RAM).
	CachePreloadMaxBytes int64 = 100 << 20
)

// CacheL2MaxBytes : plafond du miroir disque = total − RAM. Fonction (et non
// valeur figée à l'init) pour que les tests qui réduisent les plafonds
// restent cohérents.
func CacheL2MaxBytes() int64 { return CacheTotalMaxBytes - CacheL1MaxBytes }

const (
	// cacheWarmTimeout : budget max de l'estimation + préchargement.
	cacheWarmTimeout = 30 * time.Second
)

// cacheMinDiskFree : garde-fou F6.5 — on ne peuple le miroir disque que
// s'il reste au moins 1 Go libre sur le volume. Variable pour les tests.
var cacheMinDiskFree int64 = 1 << 30

// CacheConfig configure le décorateur.
type CacheConfig struct {
	// L2Dir : répertoire du miroir disque (projets distants). Vide = pas de
	// niveau disque (projets locaux : le FS source fait office de L2).
	L2Dir string
}

// cacheEntry : une entrée du niveau RAM.
type cacheEntry struct {
	key     string // chemin canonique relatif (slashes), sans ".."
	data    []byte // jamais muté en place : remplacé à chaque écriture
	size    int64
	modTime time.Time
}

// l2Entry : suivi LRU du miroir disque (le contenu est dans le fichier).
type l2Entry struct {
	key  string
	size int64
}

// CachedFS décore un FS avec le cache à deux niveaux. Thread-safe.
type CachedFS struct {
	inner FS
	cfg   CacheConfig

	mu      sync.Mutex
	l1      map[string]*list.Element // key -> élément (Value: *cacheEntry)
	lru     *list.List               // avant = le plus récemment utilisé
	l1Bytes int64

	l2      map[string]*list.Element // key -> élément (Value: *l2Entry)
	l2lru   *list.List
	l2Bytes int64

	warmed int32 // préchargement lancé (une seule fois)
}

// NewCachedFS enveloppe inner. Le niveau disque est désactivé proprement
// (avec log) si le répertoire miroir est inutilisable — le cache ne fait
// jamais échouer l'ouverture du projet.
func NewCachedFS(inner FS, cfg CacheConfig) *CachedFS {
	c := &CachedFS{
		inner: inner,
		cfg:   cfg,
		l1:    make(map[string]*list.Element),
		lru:   list.New(),
		l2:    make(map[string]*list.Element),
		l2lru: list.New(),
	}
	if cfg.L2Dir != "" {
		if err := os.MkdirAll(cfg.L2Dir, 0o700); err != nil {
			log.Printf("cache: miroir disque désactivé (%s): %v", cfg.L2Dir, err)
			c.cfg.L2Dir = ""
		} else {
			c.scanL2()
		}
	}
	return c
}

// ---------------------------------------------------------------------------
// vfs.FS : délégation pure (aucun état caché ne dépend de ces appels).
// ---------------------------------------------------------------------------

func (c *CachedFS) Name() string                       { return c.inner.Name() }
func (c *CachedFS) Remote() bool                       { return c.inner.Remote() }
func (c *CachedFS) Close() error                       { return c.inner.Close() }
func (c *CachedFS) Resolve(rel string) (string, error) { return c.inner.Resolve(rel) }
func (c *CachedFS) MkdirAll(ctx context.Context, rel string) error {
	return c.inner.MkdirAll(ctx, rel)
}
func (c *CachedFS) ReadDir(ctx context.Context, rel string) ([]Entry, error) {
	return c.inner.ReadDir(ctx, rel)
}
func (c *CachedFS) Stat(ctx context.Context, rel string) (Entry, error) {
	return c.inner.Stat(ctx, rel)
}
func (c *CachedFS) Walk(ctx context.Context, fn func(Entry) error) error {
	return c.inner.Walk(ctx, fn)
}
func (c *CachedFS) Exec(ctx context.Context, name string, args []string, env map[string]string, timeout time.Duration) (string, error) {
	return c.inner.Exec(ctx, name, args, env, timeout)
}

// ---------------------------------------------------------------------------
// Lecture : L1 -> L2 -> source, avec validation F6.4.
// ---------------------------------------------------------------------------

// ReadFile lit un fichier via le cache. La donnée retournée est une copie.
func (c *CachedFS) ReadFile(ctx context.Context, rel string) ([]byte, error) {
	clean, err := c.inner.Resolve(rel)
	if err != nil {
		return nil, err
	}
	if clean == "" {
		return nil, errors.New("chemin vide")
	}
	if s := c.l1Get(clean); s != nil {
		// F6.4 : le Stat valide que la source n'a pas changé hors cache
		// (outil Bash, édition distante, upload UI direct...).
		if c.fresh(ctx, s) {
			c.l1Touch(clean)
			return cloneBytes(s.data), nil
		}
	}
	return c.fetch(ctx, clean)
}

// snapshot d'une entrée L1 (copie des métadonnées ; data partagé en lecture
// seule — les écritures remplacent l'entrée, jamais en place).
type l1snap struct {
	key     string
	data    []byte
	size    int64
	modTime time.Time
}

func (c *CachedFS) l1Get(key string) *l1snap {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.l1[key]; ok {
		e := el.Value.(*cacheEntry)
		return &l1snap{key: key, data: e.data, size: e.size, modTime: e.modTime}
	}
	return nil
}

func (c *CachedFS) l1Touch(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.l1[key]; ok {
		c.lru.MoveToFront(el)
	}
}

// fresh : l'entrée est-elle toujours identique à la source ?
func (c *CachedFS) fresh(ctx context.Context, s *l1snap) bool {
	st, err := c.inner.Stat(ctx, s.key)
	if err != nil {
		return false
	}
	return st.Size == s.size && st.ModTime.Equal(s.modTime)
}

// fetch remplit le cache depuis L2 (miroir distant) ou la source.
func (c *CachedFS) fetch(ctx context.Context, key string) ([]byte, error) {
	var rst *Entry
	if c.cfg.L2Dir != "" {
		if data, st, ok := c.l2Serve(ctx, key); ok {
			c.l1Set(key, data, *st)
			return cloneBytes(data), nil
		} else if st != nil {
			rst = st // Stat distant déjà payé par l2Serve : on le réutilise
		}
	}
	data, err := c.inner.ReadFile(ctx, key)
	if err != nil {
		return nil, err
	}
	st := rst
	if st == nil {
		if s, serr := c.inner.Stat(ctx, key); serr == nil {
			st = &s
		} else {
			st = &Entry{Path: key, Size: int64(len(data)), ModTime: time.Now()}
		}
	}
	c.l1Set(key, data, *st)
	c.l2Store(key, data, *st)
	return cloneBytes(data), nil
}

// l1Set insère en L1 (copie défensive), avec éviction LRU au-delà du plafond.
func (c *CachedFS) l1Set(key string, data []byte, st Entry) {
	size := int64(len(data))
	if size > CacheL1MaxBytes {
		return // un seul fichier plus gros que toute la RAM : pas de L1
	}
	cp := cloneBytes(data)
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.l1[key]; ok {
		c.l1Bytes -= el.Value.(*cacheEntry).size
		c.lru.Remove(el)
		delete(c.l1, key)
	}
	for c.l1Bytes+size > CacheL1MaxBytes && c.lru.Len() > 0 {
		back := c.lru.Back()
		e := back.Value.(*cacheEntry)
		c.l1Bytes -= e.size
		delete(c.l1, e.key)
		c.lru.Remove(back)
	}
	e := &cacheEntry{key: key, data: cp, size: size, modTime: st.ModTime}
	c.l1[key] = c.lru.PushFront(e)
	c.l1Bytes += size
}

// ---------------------------------------------------------------------------
// Écriture : write-back synchronisé à chaque fin de modification.
// ---------------------------------------------------------------------------

// WriteFile écrit la source puis synchronise L1 et L2. Si l'écriture source
// échoue, le cache n'est pas touché.
func (c *CachedFS) WriteFile(ctx context.Context, rel string, data []byte, perm os.FileMode) error {
	clean, err := c.inner.Resolve(rel)
	if err != nil {
		return err
	}
	if clean == "" {
		return errors.New("chemin vide")
	}
	if err := c.inner.WriteFile(ctx, clean, data, perm); err != nil {
		return err
	}
	var st Entry
	if s, serr := c.inner.Stat(ctx, clean); serr == nil {
		st = s
	} else {
		st = Entry{Path: clean, Size: int64(len(data)), ModTime: time.Now()}
	}
	c.l1Set(clean, data, st)
	c.l2Store(clean, data, st)
	return nil
}

// Remove supprime la source puis invalide le cache (chemin + sous-arbre).
func (c *CachedFS) Remove(ctx context.Context, rel string) error {
	clean, err := c.inner.Resolve(rel)
	if err != nil {
		return err
	}
	if clean == "" {
		return errors.New("chemin vide")
	}
	if err := c.inner.Remove(ctx, clean); err != nil {
		return err
	}
	c.invalidatePrefix(clean)
	return nil
}

// Rename déplace la source puis invalide les deux côtés (simple et sûr).
func (c *CachedFS) Rename(ctx context.Context, oldrel, newrel string) error {
	oldclean, err := c.inner.Resolve(oldrel)
	if err != nil {
		return err
	}
	newclean, err := c.inner.Resolve(newrel)
	if err != nil {
		return err
	}
	if err := c.inner.Rename(ctx, oldclean, newclean); err != nil {
		return err
	}
	c.invalidatePrefix(oldclean)
	c.invalidatePrefix(newclean)
	return nil
}

// invalidatePrefix retire du cache un chemin et tout son sous-arbre.
func (c *CachedFS) invalidatePrefix(prefix string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, el := range c.l1 {
		if key == prefix || strings.HasPrefix(key, prefix+"/") {
			c.l1Bytes -= el.Value.(*cacheEntry).size
			c.lru.Remove(el)
			delete(c.l1, key)
		}
	}
	for key, el := range c.l2 {
		if key == prefix || strings.HasPrefix(key, prefix+"/") {
			c.l2Bytes -= el.Value.(*l2Entry).size
			c.l2lru.Remove(el)
			delete(c.l2, key)
		}
	}
	if c.cfg.L2Dir != "" {
		if p := c.l2Path(prefix); p != "" {
			_ = os.Remove(p)    // fichier miroir
			_ = os.RemoveAll(p) // ou sous-arbre miroir
		}
	}
}

// ---------------------------------------------------------------------------
// Niveau 2 : miroir disque (projets distants uniquement).
// ---------------------------------------------------------------------------

// l2Path traduit une clé en chemin miroir, en restant confiné au L2Dir.
func (c *CachedFS) l2Path(key string) string {
	if c.cfg.L2Dir == "" {
		return ""
	}
	p := filepath.Join(c.cfg.L2Dir, filepath.FromSlash(key))
	rel, err := filepath.Rel(c.cfg.L2Dir, p)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		return ""
	}
	return p
}

// l2Serve tente de servir depuis le miroir après validation F6.4 (le mtime
// du fichier miroir = le mtime distant au moment du rapatriement).
// Retourne aussi le Stat distant quand il a été obtenu (évite un 2e appel).
func (c *CachedFS) l2Serve(ctx context.Context, key string) (data []byte, st *Entry, ok bool) {
	p := c.l2Path(key)
	if p == "" {
		return nil, nil, false
	}
	fi, err := os.Stat(p)
	if err != nil || fi.IsDir() {
		return nil, nil, false
	}
	rst, err := c.inner.Stat(ctx, key)
	if err != nil {
		return nil, nil, false
	}
	if fi.Size() != rst.Size || !fi.ModTime().Equal(rst.ModTime) {
		return nil, &rst, false // miroir périmé : la source a changé
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return nil, &rst, false
	}
	c.l2Touch(key)
	return b, &rst, true
}

func (c *CachedFS) l2Touch(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.l2[key]; ok {
		c.l2lru.MoveToFront(el)
	}
}

// l2Store écrit le miroir (avec le mtime distant) en respectant le plafond
// et le garde-fou disque F6.5. Silencieux en cas d'échec : le cache reste
// un accélérateur, jamais un point de défaillance.
func (c *CachedFS) l2Store(key string, data []byte, st Entry) {
	if c.cfg.L2Dir == "" {
		return
	}
	p := c.l2Path(key)
	if p == "" {
		return
	}
	size := int64(len(data))
	if size > CacheL2MaxBytes() {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.l2[key]; ok { // remplacement : on retire l'ancien d'abord
		c.l2Bytes -= el.Value.(*l2Entry).size
		c.l2lru.Remove(el)
		delete(c.l2, key)
	}
	for c.l2Bytes+size > CacheL2MaxBytes() && c.l2lru.Len() > 0 {
		c.l2EvictLocked()
	}
	if !c.diskOK() {
		log.Printf("cache: miroir disque en pause (espace libre < 1 Go)")
		return
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
		return
	}
	if err := os.WriteFile(p, data, 0o600); err != nil {
		return
	}
	// Le mtime du miroir = le mtime distant : la validation F6.4 compare
	// ensuite directement les deux sans side-car.
	_ = os.Chtimes(p, st.ModTime, st.ModTime)
	c.l2[key] = c.l2lru.PushFront(&l2Entry{key: key, size: size})
	c.l2Bytes += size
}

// l2EvictLocked retire l'entrée la moins récemment utilisée (verrou acquis).
func (c *CachedFS) l2EvictLocked() {
	back := c.l2lru.Back()
	if back == nil {
		return
	}
	e := back.Value.(*l2Entry)
	c.l2Bytes -= e.size
	delete(c.l2, e.key)
	c.l2lru.Remove(back)
	if p := c.l2Path(e.key); p != "" {
		_ = os.Remove(p)
	}
}

// diskOK : garde-fou F6.5 — assez d'espace libre pour peupler le miroir ?
// diskFreeBytes est specifique a la plateforme (Statfs absent sous Windows).
func (c *CachedFS) diskOK() bool {
	free, ok := diskFreeBytes(c.cfg.L2Dir)
	if !ok {
		return true // doute : on ne bloque pas le cache
	}
	return free >= cacheMinDiskFree
}

// scanL2 reconstruit l'index du miroir existant (persiste entre redémarrages).
func (c *CachedFS) scanL2() {
	type found struct {
		key string
		fi  os.FileInfo
	}
	var files []found
	_ = filepath.WalkDir(c.cfg.L2Dir, func(p string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		rel, rerr := filepath.Rel(c.cfg.L2Dir, p)
		if rerr != nil {
			return nil
		}
		fi, ferr := d.Info()
		if ferr != nil {
			return nil
		}
		files = append(files, found{key: filepath.ToSlash(rel), fi: fi})
		return nil
	})
	// Ordre LRU approximatif : le plus récent devant.
	for i := len(files) - 1; i >= 0; i-- {
		f := files[i]
		if c.l2Bytes+f.fi.Size() > CacheL2MaxBytes() {
			_ = os.Remove(filepath.Join(c.cfg.L2Dir, filepath.FromSlash(f.key)))
			continue
		}
		c.l2[f.key] = c.l2lru.PushFront(&l2Entry{key: f.key, size: f.fi.Size()})
		c.l2Bytes += f.fi.Size()
	}
}

// ---------------------------------------------------------------------------
// Préchargement : dossier ≤ 100 Mo → RAM, asynchrone et borné.
// ---------------------------------------------------------------------------

// StartWarmup lance le préchargement une seule fois (goroutine dédiée).
func (c *CachedFS) StartWarmup() {
	if !atomic.CompareAndSwapInt32(&c.warmed, 0, 1) {
		return
	}
	go c.warmup()
}

func (c *CachedFS) warmup() {
	ctx, cancel := context.WithTimeout(context.Background(), cacheWarmTimeout)
	defer cancel()
	// 1) Estimation bornée : on arrête de compter passé 100 Mo.
	var total int64
	_ = c.inner.Walk(ctx, func(e Entry) error {
		if e.IsDir || SkipEntry(path.Base(e.Path)) {
			return nil
		}
		total += e.Size
		if total > CachePreloadMaxBytes {
			cancel() // Walk s'arrête proprement sur ctx annulé
		}
		return nil
	})
	if total == 0 || total > CachePreloadMaxBytes {
		return // vide, trop gros ou timeout : peuplement paresseux
	}
	// 2) Chargement : chaque ReadFile peuple L1 (et L2 si distant).
	_ = c.inner.Walk(ctx, func(e Entry) error {
		if e.IsDir || SkipEntry(path.Base(e.Path)) {
			return nil
		}
		_, _ = c.ReadFile(ctx, e.Path)
		return nil
	})
}

// ---------------------------------------------------------------------------
// Introspection (tests, debug).
// ---------------------------------------------------------------------------

// Stats retourne l'occupation actuelle des deux niveaux (octets).
func (c *CachedFS) Stats() (l1Bytes, l2Bytes int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.l1Bytes, c.l2Bytes
}

func cloneBytes(b []byte) []byte {
	out := make([]byte, len(b))
	copy(out, b)
	return out
}
