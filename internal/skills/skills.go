// Package skills : competences (skills) definies par l'utilisateur et
// injectees dans le prompt systeme de l'agent quand elles sont actives.
//
// Une competence = un nom, une description et des instructions (texte
// libre) que l'agent doit suivre. Elles sont stockees par utilisateur
// dans le store (cle "skills") et gerees depuis
// Configuration -> Competences.
package skills

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"strings"
	"sync"
)

const storeKey = "skills"

// Limites anti-abus : une competence reste un snippet de prompt, pas un
// roman.
const (
	maxSkills     = 64
	maxNameLen    = 80
	maxDescLen    = 300
	maxInstrLen   = 8000
	maxTotalInstr = 24000
)

// Skill est une competence configurable par l'utilisateur.
type Skill struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	Instructions string `json:"instructions"`
	Enabled      bool   `json:"enabled"`
}

// SettingsStore est le sous-ensemble du store utilise par ce package.
type SettingsStore interface {
	GetSetting(user, key string) ([]byte, bool)
	PutSetting(user, key string, val []byte) error
}

var mu sync.Mutex

func newID() string {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b[:])
}

// List renvoie les competences de l'utilisateur (liste vide si aucune).
func List(st SettingsStore, user string) []Skill {
	mu.Lock()
	defer mu.Unlock()
	raw, ok := st.GetSetting(user, storeKey)
	if !ok {
		return []Skill{}
	}
	var out []Skill
	if err := json.Unmarshal(raw, &out); err != nil {
		return []Skill{}
	}
	if out == nil {
		return []Skill{}
	}
	return out
}

// Save remplace la liste des competences apres validation.
func Save(st SettingsStore, user string, in []Skill) ([]Skill, error) {
	mu.Lock()
	defer mu.Unlock()
	out, err := normalize(in)
	if err != nil {
		return nil, err
	}
	raw, err := json.Marshal(out)
	if err != nil {
		return nil, err
	}
	if err := st.PutSetting(user, storeKey, raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ActiveInstructions renvoie les instructions des competences activees,
// dans l'ordre de la liste. Chaine vide si aucune.
func ActiveInstructions(st SettingsStore, user string) string {
	var sb strings.Builder
	for _, sk := range List(st, user) {
		if !sk.Enabled {
			continue
		}
		ins := strings.TrimSpace(sk.Instructions)
		if ins == "" {
			continue
		}
		if sb.Len() > 0 {
			sb.WriteString("\n\n")
		}
		name := strings.TrimSpace(sk.Name)
		if name == "" {
			name = "skill"
		}
		sb.WriteString("[Skill: " + name + "]\n" + ins)
	}
	return sb.String()
}

func normalize(in []Skill) ([]Skill, error) {
	if len(in) > maxSkills {
		return nil, errTooMany
	}
	out := make([]Skill, 0, len(in))
	total := 0
	seen := map[string]bool{}
	for _, sk := range in {
		name := strings.TrimSpace(sk.Name)
		ins := strings.TrimSpace(sk.Instructions)
		if name == "" {
			return nil, errNoName
		}
		if ins == "" {
			return nil, errNoInstructions
		}
		if len([]rune(name)) > maxNameLen {
			return nil, errNameTooLong
		}
		if len([]rune(sk.Description)) > maxDescLen {
			return nil, errDescTooLong
		}
		if len([]rune(ins)) > maxInstrLen {
			return nil, errInstrTooLong
		}
		total += len([]rune(ins))
		if total > maxTotalInstr {
			return nil, errTotalTooLong
		}
		id := strings.TrimSpace(sk.ID)
		if id == "" || seen[id] {
			id = newID()
		}
		seen[id] = true
		out = append(out, Skill{
			ID:           id,
			Name:         name,
			Description:  strings.TrimSpace(sk.Description),
			Instructions: ins,
			Enabled:      sk.Enabled,
		})
	}
	return out, nil
}
