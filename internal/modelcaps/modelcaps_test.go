package modelcaps

import (
	"path/filepath"
	"testing"

	"cetas-lite/internal/store"
)

func TestSaveLoad(t *testing.T) {
	st, err := store.Open(filepath.Join(t.TempDir(), "caps.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer st.Close()

	if got := Load(st); len(got) != 0 {
		t.Fatalf("vide attendu: %v", got)
	}
	m := Map{Key("openrouter", "qwen-vl"): {Vision: true}}
	if err := Save(st, m); err != nil {
		t.Fatal(err)
	}
	got := Load(st)
	if !got.Vision("openrouter", "qwen-vl") {
		t.Fatalf("vision attendue: %v", got)
	}
	if got.Vision("deepseek", "deepseek-chat") {
		t.Fatal("deepseek ne doit pas etre vision")
	}
}
