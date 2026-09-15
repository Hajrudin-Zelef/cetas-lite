package skills

import (
	"strings"
	"sync"
	"testing"
)

type memStore struct {
	mu   sync.Mutex
	data map[string]map[string][]byte
}

func newMemStore() *memStore { return &memStore{data: map[string]map[string][]byte{}} }

func (m *memStore) GetSetting(user, key string) ([]byte, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.data[user][key]
	return v, ok
}

func (m *memStore) PutSetting(user, key string, val []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data[user] == nil {
		m.data[user] = map[string][]byte{}
	}
	m.data[user][key] = append([]byte(nil), val...)
	return nil
}

func TestSaveListRoundTrip(t *testing.T) {
	st := newMemStore()
	in := []Skill{
		{Name: "revue", Description: "revue de code", Instructions: "sois exigeant", Enabled: true},
		{Name: "doc", Instructions: "documente en francais", Enabled: false},
	}
	saved, err := Save(st, "sam", in)
	if err != nil {
		t.Fatalf("Save: %v", err)
	}
	if saved[0].ID == "" || saved[1].ID == "" {
		t.Fatal("chaque competence doit recevoir un id")
	}
	if saved[0].ID == saved[1].ID {
		t.Fatal("ids dupliques")
	}
	got := List(st, "sam")
	if len(got) != 2 || got[0].Name != "revue" || !got[0].Enabled {
		t.Fatalf("liste inattendue: %+v", got)
	}
}

func TestActiveInstructionsOnlyEnabled(t *testing.T) {
	st := newMemStore()
	if _, err := Save(st, "sam", []Skill{
		{Name: "a", Instructions: "fais A", Enabled: true},
		{Name: "b", Instructions: "fais B", Enabled: false},
	}); err != nil {
		t.Fatalf("Save: %v", err)
	}
	got := ActiveInstructions(st, "sam")
	if !strings.Contains(got, "fais A") {
		t.Fatalf("instruction active manquante: %q", got)
	}
	if strings.Contains(got, "fais B") {
		t.Fatalf("instruction desactivee ne doit pas apparaitre: %q", got)
	}
	if ActiveInstructions(st, "nobody") != "" {
		t.Fatal("aucune competence -> chaine vide")
	}
}

func TestSaveValidation(t *testing.T) {
	st := newMemStore()
	cases := []struct {
		name string
		in   []Skill
	}{
		{"sans nom", []Skill{{Instructions: "x"}}},
		{"sans instructions", []Skill{{Name: "x"}}},
		{"nom trop long", []Skill{{Name: strings.Repeat("n", 81), Instructions: "x"}}},
		{"instructions trop longues", []Skill{{Name: "x", Instructions: strings.Repeat("i", 8001)}}},
	}
	for _, c := range cases {
		if _, err := Save(st, "sam", c.in); err == nil {
			t.Fatalf("%s: erreur attendue", c.name)
		}
	}
	var many []Skill
	for i := 0; i < 65; i++ {
		many = append(many, Skill{Name: "s", Instructions: "x"})
	}
	if _, err := Save(st, "sam", many); err == nil {
		t.Fatal("trop de competences: erreur attendue")
	}
}
