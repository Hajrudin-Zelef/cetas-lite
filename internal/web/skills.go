package web

import (
	"net/http"

	"cetas-lite/internal/skills"
)

// GET /api/skills — liste les competences de l'utilisateur.
func (s *Server) handleSkillsGet(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"skills": skills.List(s.st, claims.Username)})
}

// PUT /api/skills — remplace la liste des competences (validation cote
// serveur : noms, tailles, total borne).
func (s *Server) handleSkillsPut(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var body struct {
		Skills []skills.Skill `json:"skills"`
	}
	if err := decodeJSON(r, &body); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	if body.Skills == nil {
		body.Skills = []skills.Skill{}
	}
	saved, err := skills.Save(s.st, claims.Username, body.Skills)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"skills": saved})
}
