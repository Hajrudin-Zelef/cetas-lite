package vfs

// Tests du cache de working-set à deux niveaux (phase 3) : L1 RAM, miroir
// disque L2 (FS distant), write-back, invalidation F6.4 et éviction LRU.

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"
)

// fakeFS : FS en mémoire, distant ou non, avec compteurs d'appels pour
// vérifier ce qui est réellement servi par le cache.
type fakeFile struct {
	data    []byte
	modTime time.Time
}

type fakeFS struct {
	mu     sync.Mutex
	remote bool
	files  map[string]*fakeFile
	reads  int
	stats  int
	writes int
	failW  bool
}

func newFakeFS(remote bool) *fakeFS {
	return &fakeFS{remote: remote, files: map[string]*fakeFile{}}
}

func (f *fakeFS) Name() string { return "fake" }
func (f *fakeFS) Remote() bool { return f.remote }
func (f *fakeFS) Resolve(rel string) (string, error) {
	return CleanRel(rel)
}
func (f *fakeFS) ReadFile(ctx context.Context, rel string) ([]byte, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	fl, ok := f.files[rel]
	if !ok {
		return nil, ErrNotFound
	}
	f.reads++
	out := make([]byte, len(fl.data))
	copy(out, fl.data)
	return out, nil
}
func (f *fakeFS) WriteFile(ctx context.Context, rel string, data []byte, perm os.FileMode) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.failW {
		return errors.New("écriture refusée")
	}
	f.writes++
	cp := make([]byte, len(data))
	copy(cp, data)
	f.files[rel] = &fakeFile{data: cp, modTime: time.Now()}
	return nil
}
func (f *fakeFS) MkdirAll(ctx context.Context, rel string) error { return nil }
func (f *fakeFS) Remove(ctx context.Context, rel string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	prefix := rel + "/"
	for k := range f.files {
		if k == rel || strings.HasPrefix(k, prefix) {
			delete(f.files, k)
		}
	}
	return nil
}
func (f *fakeFS) Rename(ctx context.Context, oldrel, newrel string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	prefix := oldrel + "/"
	for k, v := range f.files {
		switch {
		case k == oldrel:
			delete(f.files, k)
			f.files[newrel] = v
		case strings.HasPrefix(k, prefix):
			delete(f.files, k)
			f.files[newrel+"/"+strings.TrimPrefix(k, prefix)] = v
		}
	}
	return nil
}
func (f *fakeFS) Stat(ctx context.Context, rel string) (Entry, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	fl, ok := f.files[rel]
	if !ok {
		return Entry{}, ErrNotFound
	}
	f.stats++
	return Entry{Path: rel, Size: int64(len(fl.data)), ModTime: fl.modTime}, nil
}
func (f *fakeFS) ReadDir(ctx context.Context, rel string) ([]Entry, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	var out []Entry
	seen := map[string]bool{}
	prefix := ""
	if rel != "" && rel != "." {
		prefix = rel + "/"
	}
	for k, v := range f.files {
		if !strings.HasPrefix(k, prefix) {
			continue
		}
		rest := strings.TrimPrefix(k, prefix)
		name := rest
		isDir := false
		if i := strings.Index(rest, "/"); i >= 0 {
			name = rest[:i]
			isDir = true
		}
		if seen[name] {
			continue
		}
		seen[name] = true
		out = append(out, Entry{Path: prefix + name, IsDir: isDir, Size: int64(len(v.data)), ModTime: v.modTime})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Path < out[j].Path })
	return out, nil
}
func (f *fakeFS) Walk(ctx context.Context, fn func(Entry) error) error {
	var walk func(rel string) error
	walk = func(rel string) error {
		entries, _ := f.ReadDir(ctx, rel)
		for _, e := range entries {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			if err := fn(e); err != nil {
				return err
			}
			if e.IsDir {
				if err := walk(strings.TrimSuffix(e.Path, "/")); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return walk("")
}
func (f *fakeFS) Exec(ctx context.Context, name string, args []string, env map[string]string, timeout time.Duration) (string, error) {
	return "", errors.New("non supporté")
}
func (f *fakeFS) Close() error { return nil }

// put écrit directement dans le fake (simulation d'une modification externe,
// hors outils suivis : l'équivalent d'un `rm` via l'outil Bash).
func (f *fakeFS) put(key string, data []byte) {
	f.mu.Lock()
	defer f.mu.Unlock()
	cp := make([]byte, len(data))
	copy(cp, data)
	f.files[key] = &fakeFile{data: cp, modTime: time.Now().Add(time.Second)}
}

func (f *fakeFS) counters() (reads, stats, writes int) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.reads, f.stats, f.writes
}

func testCtx() context.Context { return context.Background() }

// withSmallCaps réduit temporairement les plafonds pour tester l'éviction.
func withSmallCaps(t *testing.T, l1, l2 int64) {
	t.Helper()
	oldL1, oldTotal := CacheL1MaxBytes, CacheTotalMaxBytes
	CacheL1MaxBytes, CacheTotalMaxBytes = l1, l1+l2 // L2 = total − L1
	t.Cleanup(func() { CacheL1MaxBytes, CacheTotalMaxBytes = oldL1, oldTotal })
}

func TestCachedFSL1Hit(t *testing.T) {
	f := newFakeFS(false)
	f.put("a.txt", []byte("hello"))
	c := NewCachedFS(f, CacheConfig{})

	b1, err := c.ReadFile(testCtx(), "a.txt")
	if err != nil || string(b1) != "hello" {
		t.Fatalf("lecture 1 : %v %q", err, b1)
	}
	reads, _, _ := f.counters()
	if reads != 1 {
		t.Fatalf("lecture 1 : %d lectures source, attendu 1", reads)
	}
	b2, err := c.ReadFile(testCtx(), "a.txt")
	if err != nil || string(b2) != "hello" {
		t.Fatalf("lecture 2 : %v %q", err, b2)
	}
	reads, _, _ = f.counters()
	if reads != 1 {
		t.Fatalf("lecture 2 : %d lectures source, attendu 1 (cache L1)", reads)
	}
	if l1, _ := c.Stats(); l1 != int64(len("hello")) {
		t.Fatalf("Stats L1 = %d, attendu %d", l1, len("hello"))
	}
}

func TestCachedFSF64ExternalChange(t *testing.T) {
	f := newFakeFS(false)
	f.put("a.txt", []byte("v1"))
	c := NewCachedFS(f, CacheConfig{})

	if _, err := c.ReadFile(testCtx(), "a.txt"); err != nil {
		t.Fatal(err)
	}
	// Modification externe (hors cache) : mtime + taille changent.
	f.put("a.txt", []byte("v2-plus-long"))
	b, err := c.ReadFile(testCtx(), "a.txt")
	if err != nil || string(b) != "v2-plus-long" {
		t.Fatalf("après modif externe : %v %q (F6.4)", err, b)
	}
	reads, _, _ := f.counters()
	if reads != 2 {
		t.Fatalf("re-fetch attendu après modif externe, lectures source = %d", reads)
	}
}

func TestCachedFSWriteBack(t *testing.T) {
	f := newFakeFS(false)
	c := NewCachedFS(f, CacheConfig{})

	if err := c.WriteFile(testCtx(), "n.txt", []byte("contenu"), 0o644); err != nil {
		t.Fatal(err)
	}
	// La source a reçu l'écriture...
	raw, err := f.ReadFile(testCtx(), "n.txt")
	if err != nil || string(raw) != "contenu" {
		t.Fatalf("source : %v %q", err, raw)
	}
	// ...et le cache sert la nouvelle valeur sans relire la source.
	readsBefore, _, _ := f.counters()
	b, err := c.ReadFile(testCtx(), "n.txt")
	if err != nil || string(b) != "contenu" {
		t.Fatalf("relecture : %v %q", err, b)
	}
	if reads, _, _ := f.counters(); reads != readsBefore {
		t.Fatalf("relecture après Write : %d lectures source (attendu %d, cache L1)",
			reads, readsBefore)
	}
}

func TestCachedFSWriteFailureKeepsCache(t *testing.T) {
	f := newFakeFS(false)
	f.put("a.txt", []byte("ok"))
	c := NewCachedFS(f, CacheConfig{})
	if _, err := c.ReadFile(testCtx(), "a.txt"); err != nil {
		t.Fatal(err)
	}
	f.failW = true
	if err := c.WriteFile(testCtx(), "a.txt", []byte("ko"), 0o644); err == nil {
		t.Fatal("l'écriture aurait dû échouer")
	}
	b, err := c.ReadFile(testCtx(), "a.txt")
	if err != nil || string(b) != "ok" {
		t.Fatalf("cache empoisonné après échec d'écriture : %v %q", err, b)
	}
}

func TestCachedFSRemoveRenameInvalidate(t *testing.T) {
	f := newFakeFS(false)
	f.put("a.txt", []byte("x"))
	f.put("d/b.txt", []byte("y"))
	c := NewCachedFS(f, CacheConfig{})
	if _, err := c.ReadFile(testCtx(), "a.txt"); err != nil {
		t.Fatal(err)
	}
	if _, err := c.ReadFile(testCtx(), "d/b.txt"); err != nil {
		t.Fatal(err)
	}
	if err := c.Remove(testCtx(), "d"); err != nil {
		t.Fatal(err)
	}
	if _, err := c.ReadFile(testCtx(), "d/b.txt"); err == nil {
		t.Fatal("d/b.txt devrait être introuvable après Remove(d)")
	}
	if l1, _ := c.Stats(); l1 != int64(len("x")) {
		t.Fatalf("Stats L1 = %d après invalidation du sous-arbre, attendu %d", l1, len("x"))
	}
	if err := c.Rename(testCtx(), "a.txt", "z.txt"); err != nil {
		t.Fatal(err)
	}
	if _, err := c.ReadFile(testCtx(), "a.txt"); err == nil {
		t.Fatal("a.txt devrait être introuvable après Rename")
	}
	b, err := c.ReadFile(testCtx(), "z.txt")
	if err != nil || string(b) != "x" {
		t.Fatalf("z.txt : %v %q", err, b)
	}
}

func TestCachedFSLRUEviction(t *testing.T) {
	withSmallCaps(t, 10, 10)
	f := newFakeFS(false)
	for i, name := range []string{"a", "b", "c"} {
		f.put(name, []byte("123456")) // 6 octets chacun
		_ = i
	}
	c := NewCachedFS(f, CacheConfig{})
	ctx := testCtx()
	for _, name := range []string{"a", "b", "c"} {
		if _, err := c.ReadFile(ctx, name); err != nil {
			t.Fatal(err)
		}
	}
	if l1, _ := c.Stats(); l1 > 10 {
		t.Fatalf("L1 = %d octets > plafond 10 (éviction LRU)", l1)
	}
	// "a" a été évincé (le plus ancien) : sa relecture repasse par la source.
	readsBefore, _, _ := f.counters()
	if _, err := c.ReadFile(ctx, "a"); err != nil {
		t.Fatal(err)
	}
	if reads, _, _ := f.counters(); reads == readsBefore {
		t.Fatal("a aurait dû être évincé (LRU) puis relu depuis la source")
	}
}

// withNoDiskGuard désactive temporairement le garde-fou d'espace disque.
func withNoDiskGuard(t *testing.T) {
	t.Helper()
	old := cacheMinDiskFree
	cacheMinDiskFree = 0
	t.Cleanup(func() { cacheMinDiskFree = old })
}

func TestCachedFSL2MirrorRemote(t *testing.T) {
	withNoDiskGuard(t)
	f := newFakeFS(true) // distant : L2 actif
	f.put("a.txt", []byte("distant"))
	l2dir := t.TempDir()
	c := NewCachedFS(f, CacheConfig{L2Dir: l2dir})

	b, err := c.ReadFile(testCtx(), "a.txt")
	if err != nil || string(b) != "distant" {
		t.Fatalf("lecture 1 : %v %q", err, b)
	}
	if _, l2 := c.Stats(); l2 == 0 {
		t.Fatal("le miroir disque L2 aurait dû être peuplé")
	}
	if _, err := os.Stat(filepath.Join(l2dir, "a.txt")); err != nil {
		t.Fatalf("fichier miroir absent : %v", err)
	}

	// Nouvelle instance = redémarrage serveur : le miroir persiste et évite
	// un re-téléchargement (seul le Stat de validation passe).
	readsBefore, _, _ := f.counters()
	c2 := NewCachedFS(f, CacheConfig{L2Dir: l2dir})
	b2, err := c2.ReadFile(testCtx(), "a.txt")
	if err != nil || string(b2) != "distant" {
		t.Fatalf("lecture via miroir : %v %q", err, b2)
	}
	if reads, _, _ := f.counters(); reads != readsBefore {
		t.Fatalf("re-téléchargement alors que le miroir était frais (lectures %d -> %d)",
			readsBefore, reads)
	}

	// Changement distant : le miroir périmé est détecté (F6.4) et re-fetché.
	f.put("a.txt", []byte("distant-v2"))
	b3, err := c2.ReadFile(testCtx(), "a.txt")
	if err != nil || string(b3) != "distant-v2" {
		t.Fatalf("après modif distante : %v %q", err, b3)
	}
}

func TestCachedFSL2Eviction(t *testing.T) {
	withNoDiskGuard(t)
	withSmallCaps(t, 1<<20, 10)
	f := newFakeFS(true)
	f.put("a", []byte("123456"))
	f.put("b", []byte("789012"))
	l2dir := t.TempDir()
	c := NewCachedFS(f, CacheConfig{L2Dir: l2dir})
	ctx := testCtx()
	if _, err := c.ReadFile(ctx, "a"); err != nil {
		t.Fatal(err)
	}
	if _, err := c.ReadFile(ctx, "b"); err != nil {
		t.Fatal(err)
	}
	if _, l2 := c.Stats(); l2 > 10 {
		t.Fatalf("L2 = %d > plafond 10 (éviction)", l2)
	}
}

func TestCachedFSLocalNoMirror(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "f.txt"), []byte("local"), 0o644); err != nil {
		t.Fatal(err)
	}
	lfs, err := NewLocal(dir, "test")
	if err != nil {
		t.Fatal(err)
	}
	c := NewCachedFS(lfs, CacheConfig{}) // pas de L2Dir : projets locaux
	b, err := c.ReadFile(testCtx(), "f.txt")
	if err != nil || string(b) != "local" {
		t.Fatalf("lecture locale : %v %q", err, b)
	}
	b2, err := c.ReadFile(testCtx(), "f.txt")
	if err != nil || string(b2) != "local" {
		t.Fatalf("relecture locale : %v %q", err, b2)
	}
	if _, l2 := c.Stats(); l2 != 0 {
		t.Fatalf("pas de miroir attendu en local, L2 = %d", l2)
	}
	// Modif externe directe sur le FS local : détectée par F6.4.
	if err := os.WriteFile(filepath.Join(dir, "f.txt"), []byte("local-v2!!"), 0o644); err != nil {
		t.Fatal(err)
	}
	b3, err := c.ReadFile(testCtx(), "f.txt")
	if err != nil || string(b3) != "local-v2!!" {
		t.Fatalf("après modif externe locale : %v %q", err, b3)
	}
}

func TestCachedFSWarmup(t *testing.T) {
	withSmallCaps(t, 1<<20, 1<<20)
	old := CachePreloadMaxBytes
	CachePreloadMaxBytes = 1 << 20
	defer func() { CachePreloadMaxBytes = old }()

	f := newFakeFS(false)
	f.put("a.txt", []byte("aa"))
	f.put("sub/b.txt", []byte("bb"))
	c := NewCachedFS(f, CacheConfig{})
	c.StartWarmup()
	c.StartWarmup() // 2e appel : sans effet (une seule fois)
	deadline := time.Now().Add(5 * time.Second)
	for {
		if l1, _ := c.Stats(); l1 == 4 {
			break
		}
		if time.Now().After(deadline) {
			l1, _ := c.Stats()
			t.Fatalf("préchargement incomplet après 5 s (L1 = %d)", l1)
		}
		time.Sleep(20 * time.Millisecond)
	}
	// Le dossier faisait moins que le seuil : tout est en RAM, sans lecture
	// source supplémentaire à la relecture.
	readsBefore, _, _ := f.counters()
	if _, err := c.ReadFile(testCtx(), "sub/b.txt"); err != nil {
		t.Fatal(err)
	}
	if reads, _, _ := f.counters(); reads != readsBefore {
		t.Fatal("relecture après warmup : aurait dû être servie par la RAM")
	}
}
