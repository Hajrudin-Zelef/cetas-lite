package skills

import "errors"

var (
	errTooMany        = errors.New("trop de competences (max 64)")
	errNoName         = errors.New("chaque competence doit avoir un nom")
	errNoInstructions = errors.New("chaque competence doit avoir des instructions")
	errNameTooLong    = errors.New("nom trop long (max 80 caracteres)")
	errDescTooLong    = errors.New("description trop longue (max 300 caracteres)")
	errInstrTooLong   = errors.New("instructions trop longues (max 8000 caracteres)")
	errTotalTooLong   = errors.New("total des instructions trop long (max 24000 caracteres)")
)
