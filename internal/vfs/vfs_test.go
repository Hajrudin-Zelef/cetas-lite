package vfs

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCleanRelConfinement(t *testing.T) {
	for _, bad := range []string{"", "/abs", "../x", "a/../../b", "..", "a/../.."} {
		if _, err := CleanRel(bad); err == nil {
			t.Fatalf("CleanRel(%q) devrait échouer", bad)
		}
	}
	for _, ok := range []string{"a/b.txt", "./a", "a\\b", ".", "  spaced /x "} {
		if _, err := CleanRel(ok); err != nil {
			t.Fatalf("CleanRel(%q) a échoué: %v", ok, err)
		}
	}
	if got, _ := CleanRel("a/./b"); got != "a/b" {
		t.Fatalf("nettoyage inattendu: %q", got)
	}
}

func testLocal(t *testing.T) *LocalFS {
	t.Helper()
	lfs, err := NewLocal(filepath.Join(t.TempDir(), "ws"), "test")
	if err != nil {
		t.Fatal(err)
	}
	return lfs
}

func TestLocalReadWriteWalk(t *testing.T) {
	ctx := context.Background()
	lfs := testLocal(t)
	if err := lfs.WriteFile(ctx, "dir/sub/f.txt", []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := lfs.WriteFile(ctx, ".git/config", []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := lfs.WriteFile(ctx, "node_modules/a.js", []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	b, err := lfs.ReadFile(ctx, "dir/sub/f.txt")
	if err != nil || string(b) != "hello" {
		t.Fatalf("lecture: %q %v", b, err)
	}
	if _, err := lfs.ReadFile(ctx, "nope.txt"); err == nil {
		t.Fatal("fichier absent devrait échouer")
	}
	// Walk ignore .git et node_modules.
	var got []string
	if err := lfs.Walk(ctx, func(e Entry) error { got = append(got, e.Path); return nil }); err != nil {
		t.Fatal(err)
	}
	for _, g := range got {
		if strings.Contains(g, ".git") || strings.Contains(g, "node_modules") {
			t.Fatalf("entrée ignorée présente: %q", g)
		}
	}
	found := false
	for _, g := range got {
		if g == "dir/sub/f.txt" {
			found = true
		}
	}
	if !found {
		t.Fatalf("f.txt introuvable dans %v", got)
	}
}

func TestLocalSymlinkEscape(t *testing.T) {
	ctx := context.Background()
	root := t.TempDir()
	lfs, err := NewLocal(filepath.Join(root, "ws"), "test")
	if err != nil {
		t.Fatal(err)
	}
	outside := filepath.Join(root, "outside")
	if err := os.MkdirAll(outside, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(outside, "secret.txt"), []byte("s"), 0o644); err != nil {
		t.Fatal(err)
	}
	// Lien symbolique vers l'extérieur : la lecture doit être refusée.
	_ = os.Symlink(outside, filepath.Join(root, "ws", "link"))
	if _, err := lfs.ReadFile(ctx, "link/secret.txt"); err == nil {
		t.Fatal("évasion par symlink acceptée")
	}
}

func TestTree(t *testing.T) {
	ctx := context.Background()
	lfs := testLocal(t)
	_ = lfs.WriteFile(ctx, "a.txt", []byte("a"), 0o644)
	_ = lfs.WriteFile(ctx, "sub/b.txt", []byte("b"), 0o644)
	_ = lfs.WriteFile(ctx, "sub/deep/c.txt", []byte("c"), 0o644)
	node, err := Tree(ctx, lfs, "", TreeOptions{MaxDepth: 1, MaxEntries: 100})
	if err != nil {
		t.Fatal(err)
	}
	if len(node.Children) != 2 { // a.txt + sub/
		t.Fatalf("enfants: %d", len(node.Children))
	}
	var sub *Node
	for _, c := range node.Children {
		if c.Name == "sub" {
			sub = c
		}
	}
	if sub == nil || !sub.IsDir || len(sub.Children) != 0 {
		t.Fatalf("profondeur non respectée: %+v", sub)
	}
}

func TestTreeTruncationSemantics(t *testing.T) {
	ctx := context.Background()
	lfs := testLocal(t)
	_ = lfs.WriteFile(ctx, "a.txt", []byte("a"), 0o644)
	_ = lfs.WriteFile(ctx, "sub/b.txt", []byte("b"), 0o644)

	// Palier de profondeur (cas du lazy loading UI, depth=1) : des enfants
	// sont présents, rien n'est coupé -> Truncated doit rester faux.
	node, err := Tree(ctx, lfs, "", TreeOptions{MaxDepth: 1, MaxEntries: 100})
	if err != nil {
		t.Fatal(err)
	}
	if node.Truncated {
		t.Fatal("Truncated ne doit pas être positionné au simple palier de profondeur")
	}
	if !node.DepthLimited {
		t.Fatal("DepthLimited attendu quand la profondeur demandée est atteinte")
	}
	if len(node.Children) != 2 { // a.txt + sub/
		t.Fatalf("enfants: %d", len(node.Children))
	}

	// Vrai tronquage : plafond d'entrées atteint -> Truncated vrai.
	node2, err := Tree(ctx, lfs, "", TreeOptions{MaxDepth: 4, MaxEntries: 2})
	if err != nil {
		t.Fatal(err)
	}
	if !node2.Truncated {
		t.Fatal("Truncated attendu quand le plafond d'entrées est atteint")
	}
}

func TestFlatList(t *testing.T) {
	ctx := context.Background()
	lfs := testLocal(t)
	_ = lfs.WriteFile(ctx, "b.txt", []byte("b"), 0o644)
	_ = lfs.WriteFile(ctx, "a.txt", []byte("a"), 0o644)
	list, trunc, err := FlatList(ctx, lfs, 10)
	if err != nil || trunc {
		t.Fatalf("flatlist: %v %v", trunc, err)
	}
	if len(list) != 2 || list[0] != "a.txt" || list[1] != "b.txt" {
		t.Fatalf("tri inattendu: %v", list)
	}
	_, trunc, err = FlatList(ctx, lfs, 1)
	if err != nil || !trunc {
		t.Fatalf("troncature attendue: %v %v", trunc, err)
	}
}

func TestSFTPCleanRel(t *testing.T) {
	s := NewSFTP(SFTPConfig{Host: "h", User: "u", RemotePath: "/srv/proj"}, "t")
	if _, err := s.Resolve("../evil"); err == nil {
		t.Fatal("confinement SFTP manquant")
	}
	rp, err := s.rpath("a/b.txt")
	if err != nil || rp != "/srv/proj/a/b.txt" {
		t.Fatalf("rpath: %q %v", rp, err)
	}
	if _, err := s.rpath("/abs"); err == nil {
		t.Fatal("chemin absolu accepté")
	}
}
