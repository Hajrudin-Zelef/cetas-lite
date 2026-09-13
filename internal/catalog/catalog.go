package catalog

type Model struct {
	ID          string
	Label       string
	Editeur     string
	InputPer1M  float64
	OutputPer1M float64
}

func Models() []Model {
	out := make([]Model, len(generated))
	copy(out, generated)
	return out
}

func Lookup(editeur, id string) (Model, bool) {
	for _, m := range generated {
		if m.Editeur == editeur && m.ID == id {
			return m, true
		}
	}
	return Model{}, false
}

func ByProvider(editeur string) []Model {
	var out []Model
	for _, m := range generated {
		if m.Editeur == editeur {
			out = append(out, m)
		}
	}
	return out
}
