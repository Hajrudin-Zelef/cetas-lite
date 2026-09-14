package attach

import (
	"strings"
	"testing"
)

func TestSaveGetDelete(t *testing.T) {
	s := New(t.TempDir(), 1<<20)
	a, err := s.Save("sam", "notes.md", []byte("# Notes"))
	if err != nil {
		t.Fatal(err)
	}
	if a.Kind != "text" || a.Ext != "md" || a.Size != 7 {
		t.Fatalf("meta = %+v", a)
	}
	got, data, err := s.Get("sam", a.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Name != "notes.md" || string(data) != "# Notes" {
		t.Fatalf("get = %+v %q", got, data)
	}
	if err := s.Delete("sam", a.ID); err != nil {
		t.Fatal(err)
	}
	if _, _, err := s.Get("sam", a.ID); err == nil {
		t.Fatal("piece jointe supprimee doit etre introuvable")
	}
}

func TestKindDetection(t *testing.T) {
	cases := map[string]string{
		"a.md": "text", "a.pdf": "pdf", "a.html": "html",
		"a.png": KindImage, "a.JPEG": KindImage,
	}
	for name, want := range cases {
		if got := KindFor(name); got != want {
			t.Errorf("KindFor(%q) = %q, want %q", name, got, want)
		}
	}
}

func TestIsolationAndSafety(t *testing.T) {
	s := New(t.TempDir(), 1<<20)
	a, err := s.Save("sam", "secret.txt", []byte("x"))
	if err != nil {
		t.Fatal(err)
	}
	if _, _, err := s.Get("alice", a.ID); err == nil {
		t.Fatal("alice ne doit pas lire la piece jointe de sam")
	}
	if _, _, err := s.Get("sam", "../../etc/passwd"); err == nil {
		t.Fatal("identifiant invalide doit etre refuse")
	}
}

func TestSizeLimitAndName(t *testing.T) {
	s := New(t.TempDir(), 4)
	if _, err := s.Save("sam", "big.txt", []byte("trop long")); err == nil {
		t.Fatal("depassement de taille doit echouer")
	}
	a, err := s.Save("sam", "../evil.txt", []byte("ok"))
	if err != nil {
		t.Fatal(err)
	}
	if a.Name != "evil.txt" {
		t.Fatalf("nom doit etre un basename, got %q", a.Name)
	}
	if strings.Contains(a.Name, "/") {
		t.Fatal("pas de separateur dans le nom")
	}
}
