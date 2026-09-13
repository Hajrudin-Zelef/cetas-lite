package memory

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func newStore(t *testing.T) *Store {
	t.Helper()
	return New(t.TempDir())
}

func TestAddReadList(t *testing.T) {
	s := newStore(t)
	if err := s.Add("alice", "Docker", "# Mes notes\ncontenu"); err != nil {
		t.Fatal(err)
	}
	pages := s.List("alice")
	if len(pages) != 1 || pages[0].Name != "Docker.md" || pages[0].Title != "Mes notes" {
		t.Fatalf("pages = %+v", pages)
	}
	out, err := s.Read("alice", "Docker.md", 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(out, "1\t# Mes notes") {
		t.Fatalf("read = %q", out)
	}
}

func TestAddRefusesOverwrite(t *testing.T) {
	s := newStore(t)
	if err := s.Add("alice", "a.md", "x"); err != nil {
		t.Fatal(err)
	}
	if err := s.Add("alice", "a.md", "y"); err == nil {
		t.Fatal("ecrasement doit etre refuse")
	}
}

func TestNameValidation(t *testing.T) {
	s := newStore(t)
	for _, bad := range []string{"", "   ", "../x", "a/b", "/etc/passwd", "a b", "a;b"} {
		if err := s.Add("alice", bad, "x"); err == nil {
			t.Errorf("nom %q doit etre refuse", bad)
		}
	}
	if err := s.Add("alice", "ok.name-1", "x"); err != nil {
		t.Errorf("nom valide refuse: %v", err)
	}
}

func TestUserIsolation(t *testing.T) {
	s := newStore(t)
	if err := s.Add("alice", "note.md", "secret alice"); err != nil {
		t.Fatal(err)
	}
	if got := s.List("bob"); len(got) != 0 {
		t.Fatalf("bob voit %d pages", len(got))
	}
	if got := s.Content("bob", "note.md"); got != "" {
		t.Fatalf("bob lit %q", got)
	}
}

func TestEditUniqueAndAlreadyApplied(t *testing.T) {
	s := newStore(t)
	if err := s.Add("alice", "n.md", "bonjour monde"); err != nil {
		t.Fatal(err)
	}
	if err := s.Edit("alice", "n.md", "monde", "terre"); err != nil {
		t.Fatal(err)
	}
	if got := s.Content("alice", "n.md"); !strings.Contains(got, "terre") {
		t.Fatalf("edit = %q", got)
	}
	if err := s.Edit("alice", "n.md", "monde", "terre"); !errors.Is(err, ErrAlreadyApplied) {
		t.Fatalf("deja applique attendu, got %v", err)
	}
	if err := s.Edit("alice", "n.md", "zzz", "y"); err == nil {
		t.Fatal("old introuvable doit echouer")
	}
}

func TestEditAmbiguous(t *testing.T) {
	s := newStore(t)
	if err := s.Add("alice", "n.md", "x x"); err != nil {
		t.Fatal(err)
	}
	if err := s.Edit("alice", "n.md", "x", "y"); err == nil {
		t.Fatal("ambigue doit echouer")
	}
}

func TestDeleteUpdatesIndex(t *testing.T) {
	s := newStore(t)
	if err := s.Add("alice", "a.md", "# Alpha"); err != nil {
		t.Fatal(err)
	}
	if err := s.Add("alice", "b.md", "# Beta"); err != nil {
		t.Fatal(err)
	}
	idx := s.Index("alice")
	if !strings.Contains(idx, "]("+"a.md"+")") || !strings.Contains(idx, "](b.md)") {
		t.Fatalf("index = %q", idx)
	}
	if err := s.Delete("alice", "a.md"); err != nil {
		t.Fatal(err)
	}
	idx = s.Index("alice")
	if strings.Contains(idx, "](a.md)") || !strings.Contains(idx, "](b.md)") {
		t.Fatalf("index apres delete = %q", idx)
	}
}

func TestIndexUpsertTitleOnEdit(t *testing.T) {
	s := newStore(t)
	if err := s.Add("alice", "a.md", "# Titre initial\ncorps"); err != nil {
		t.Fatal(err)
	}
	if err := s.Edit("alice", "a.md", "Titre initial", "Titre modifie"); err != nil {
		t.Fatal(err)
	}
	idx := s.Index("alice")
	if !strings.Contains(idx, "[Titre modifie](a.md)") {
		t.Fatalf("index titre = %q", idx)
	}
}

func TestIndexFileReserved(t *testing.T) {
	s := newStore(t)
	if err := s.Add("alice", "MEMORY.md", "x"); err == nil {
		t.Fatal("MEMORY.md doit etre reserve")
	}
	if err := s.Delete("alice", "MEMORY.md"); err == nil {
		t.Fatal("MEMORY.md ne doit pas etre supprimable")
	}
}

func TestReadPagination(t *testing.T) {
	s := newStore(t)
	body := strings.Join([]string{"l1", "l2", "l3", "l4"}, "\n")
	if err := s.Add("alice", "p.md", body); err != nil {
		t.Fatal(err)
	}
	out, err := s.Read("alice", "p.md", 2, 2)
	if err != nil {
		t.Fatal(err)
	}
	if out != "2\tl2\n3\tl3" {
		t.Fatalf("pagination = %q", out)
	}
}

func TestSearchCoverageThenIDF(t *testing.T) {
	s := newStore(t)
	// "nathan" est omnipresent ; "copine" est rare et discriminant.
	if err := s.Add("alice", "journal.md", "nathan nathan nathan nathan nathan nathan"); err != nil {
		t.Fatal(err)
	}
	if err := s.Add("alice", "relations.md", "copine Nathan vit ici"); err != nil {
		t.Fatal(err)
	}
	hits := s.Search("alice", "copine nathan", 8)
	if len(hits) == 0 || hits[0].File != "relations.md" {
		t.Fatalf("ranking = %+v", hits)
	}
}

func TestSearchNoMatchAndLimit(t *testing.T) {
	s := newStore(t)
	if err := s.Add("alice", "a.md", "alpha"); err != nil {
		t.Fatal(err)
	}
	if err := s.Add("alice", "b.md", "beta"); err != nil {
		t.Fatal(err)
	}
	if len(s.Search("alice", "gamma", 8)) != 0 {
		t.Fatal("aucun match attendu")
	}
	if got := s.Search("alice", "a e", 1); len(got) != 1 {
		t.Fatalf("limit = %+v", got)
	}
}

func TestSearchSnippet(t *testing.T) {
	s := newStore(t)
	if err := s.Add("alice", "a.md", "preambule\nmot cle ici\nsuite"); err != nil {
		t.Fatal(err)
	}
	hits := s.Search("alice", "cle", 8)
	if len(hits) != 1 || !strings.Contains(hits[0].Snippet, "cle") {
		t.Fatalf("snippet = %+v", hits)
	}
}

func TestStoreDisabled(t *testing.T) {
	s := New("")
	if s.List("alice") != nil {
		t.Fatal("store vide doit renvoyer nil")
	}
	if err := s.Add("alice", "a.md", "x"); err == nil {
		t.Fatal("store vide doit refuser l'ecriture")
	}
}

func TestFilesMode0600(t *testing.T) {
	s := newStore(t)
	if err := s.Add("alice", "a.md", "x"); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(filepath.Join(s.root, "alice", "a.md"))
	if err != nil {
		t.Fatal(err)
	}
	if perm := info.Mode().Perm(); perm != 0o600 {
		t.Fatalf("perm = %o", perm)
	}
}
