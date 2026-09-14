package modelcaps

import (
	"encoding/json"

	"cetas-lite/internal/store"
)

const metaKey = "model_caps"

type Caps struct {
	Vision bool `json:"vision,omitempty"`
	TTS    bool `json:"tts,omitempty"`
	STT    bool `json:"stt,omitempty"`
}

type Map map[string]Caps

func Key(provider, model string) string { return provider + "/" + model }

func Load(st *store.Store) Map {
	out := Map{}
	if st == nil {
		return out
	}
	raw, ok := st.GetMeta(metaKey)
	if !ok {
		return out
	}
	_ = json.Unmarshal(raw, &out)
	return out
}

func Save(st *store.Store, m Map) error {
	if st == nil {
		return nil
	}
	raw, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return st.PutMeta(metaKey, raw)
}

func (m Map) Get(provider, model string) Caps {
	if m == nil {
		return Caps{}
	}
	return m[Key(provider, model)]
}

func (m Map) Vision(provider, model string) bool {
	return m.Get(provider, model).Vision
}
