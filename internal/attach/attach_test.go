package attach

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
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

// TestJSONContentNotClobbered : le contenu d'une piece jointe .json ne doit
// jamais etre ecrase par ses metadonnees (bug : id.json servait aux deux).
func TestJSONContentNotClobbered(t *testing.T) {
	s := New(t.TempDir(), 1<<20)
	content := []byte(`{"name":"manif","short_name":"m","icons":[192,512]}`)
	a, err := s.Save("sam", "manifest.json", content)
	if err != nil {
		t.Fatal(err)
	}
	got, data, err := s.Get("sam", a.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Name != "manifest.json" || got.Ext != "json" {
		t.Fatalf("meta = %+v", got)
	}
	if string(data) != string(content) {
		t.Fatalf("contenu .json ecrase par les metadonnees:\n got=%s", data)
	}
	// La metadonnee vit hors du namespace des donnees.
	dir, _ := s.userDir("sam")
	if _, err := os.Stat(filepath.Join(dir, ".meta", a.ID+".json")); err != nil {
		t.Fatalf("metadonnee .meta attendue: %v", err)
	}
}

// TestSidecarAndPurge : les sidecars derives (.txt, .thumb.jpg) sont ecrits,
// relus, et purges avec la piece jointe (Delete et CleanOlderThan).
func TestSidecarAndPurge(t *testing.T) {
	s := New(t.TempDir(), 1<<20)
	a, err := s.Save("sam", "doc.pdf", []byte("PDF"))
	if err != nil {
		t.Fatal(err)
	}
	if err := s.WriteSidecar("sam", a.ID, textSidecarExt, []byte("texte extrait")); err != nil {
		t.Fatal(err)
	}
	if b, ok := s.ReadSidecar("sam", a.ID, textSidecarExt); !ok || string(b) != "texte extrait" {
		t.Fatalf("sidecar texte = %q ok=%v", b, ok)
	}
	if err := s.WriteSidecar("sam", a.ID, thumbSidecarExt, []byte("JPEG")); err != nil {
		t.Fatal(err)
	}
	dir, _ := s.userDir("sam")
	if err := s.Delete("sam", a.ID); err != nil {
		t.Fatal(err)
	}
	for _, suf := range []string{".json", ".pdf", textSidecarExt, thumbSidecarExt} {
		if _, err := os.Stat(filepath.Join(dir, a.ID+suf)); err == nil {
			t.Fatalf("sidecar/fichier %q doit etre purge par Delete", suf)
		}
	}
}

func TestCleanOlderThanPurgesSidecars(t *testing.T) {
	s := New(t.TempDir(), 1<<20)
	a, err := s.Save("sam", "vieux.md", []byte("vieux"))
	if err != nil {
		t.Fatal(err)
	}
	if err := s.WriteSidecar("sam", a.ID, textSidecarExt, []byte("t")); err != nil {
		t.Fatal(err)
	}
	dir, _ := s.userDir("sam")
	rawOld := `{"id":"` + a.ID + `","name":"vieux.md","kind":"text","ext":"md","size":5,"created":1}`
	if err := os.WriteFile(filepath.Join(dir, a.ID+".json"), []byte(rawOld), 0o600); err != nil {
		t.Fatal(err)
	}
	if n := s.CleanOlderThan(7 * 24 * time.Hour); n != 1 {
		t.Fatalf("1 piece purgee attendue, obtenu %d", n)
	}
	if _, err := os.Stat(filepath.Join(dir, a.ID+textSidecarExt)); err == nil {
		t.Fatal("le sidecar doit etre purge par CleanOlderThan")
	}
}

// TestCleanOlderThan : seules les pièces jointes plus anciennes que maxAge
// sont purgées (métadonnées + fichier), les récentes sont conservées.
func TestCleanOlderThan(t *testing.T) {
	s := New(t.TempDir(), 1<<20)
	old, err := s.Save("sam", "vieux.md", []byte("vieux"))
	if err != nil {
		t.Fatal(err)
	}
	recent, err := s.Save("sam", "recent.md", []byte("recent"))
	if err != nil {
		t.Fatal(err)
	}
	dir, err := s.userDir("sam")
	if err != nil {
		t.Fatal(err)
	}
	// Vieillit artificiellement la première pièce jointe (created = 1 ms).
	rawOld := `{"id":"` + old.ID + `","name":"vieux.md","kind":"text","ext":"md","size":5,"created":1}`
	if err := os.WriteFile(filepath.Join(dir, old.ID+".json"), []byte(rawOld), 0o600); err != nil {
		t.Fatal(err)
	}

	if n := s.CleanOlderThan(7 * 24 * time.Hour); n != 1 {
		t.Fatalf("1 pièce purgée attendue, obtenu %d", n)
	}
	if _, _, err := s.Get("sam", old.ID); err == nil {
		t.Fatal("la vieille pièce jointe doit être purgée")
	}
	if _, _, err := s.Get("sam", recent.ID); err != nil {
		t.Fatalf("la pièce récente doit être conservée: %v", err)
	}
	if n := s.CleanOlderThan(0); n != 0 {
		t.Fatalf("maxAge <= 0 : rien à purger, obtenu %d", n)
	}
}
