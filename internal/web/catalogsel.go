package web

import (
	"encoding/json"
	"net/http"

	"cetas-lite/internal/provider"
)

// Selection du catalogue de modeles.
//
// GET /api/catalog/selection -> {"disabled": {"openrouter": ["id", ...]}}
// PUT /api/catalog/selection {"disabled": {...}} -> {"ok": true}
//
// Les cases cochees du panneau "API et Modeles" ("Selectionnez les modeles
// a utiliser") sont persistee ici, par fournisseur. Borne anti-abus :
// 5000 entrees max, 200 caracteres par identifiant.

const catalogSelectionMetaKey = "catalog_selection"
const catalogSelectionMaxEntries = 5000
const catalogSelectionMaxIDLen = 200

type catalogSelection struct {
	Disabled map[string][]string `json:"disabled"`
}

func (s *Server) handleCatalogSelectionGet(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	sel := catalogSelection{Disabled: map[string][]string{}}
	if raw, ok := s.st.GetMeta(catalogSelectionMetaKey); ok && len(raw) > 0 {
		var parsed catalogSelection
		if err := json.Unmarshal(raw, &parsed); err == nil && parsed.Disabled != nil {
			sel = parsed
		}
	}
	writeJSON(w, http.StatusOK, sel)
}

func (s *Server) handleCatalogSelectionPut(w http.ResponseWriter, r *http.Request) {
	if claimsFrom(r) == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var sel catalogSelection
	if err := decodeJSON(r, &sel); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	if sel.Disabled == nil {
		sel.Disabled = map[string][]string{}
	}
	total := 0
	for prov, ids := range sel.Disabled {
		if !provider.IsCloudProvider(prov) && !isLocalEngine(prov) {
			writeError(w, http.StatusBadRequest, "fournisseur inconnu : "+prov)
			return
		}
		seen := map[string]bool{}
		clean := make([]string, 0, len(ids))
		for _, id := range ids {
			if len(id) == 0 || len(id) > catalogSelectionMaxIDLen || seen[id] {
				continue
			}
			seen[id] = true
			clean = append(clean, id)
			total++
			if total > catalogSelectionMaxEntries {
				writeError(w, http.StatusBadRequest, "selection trop volumineuse")
				return
			}
		}
		sel.Disabled[prov] = clean
	}
	raw, err := json.Marshal(sel)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "encodage impossible")
		return
	}
	if err := s.st.PutMeta(catalogSelectionMetaKey, raw); err != nil {
		writeError(w, http.StatusInternalServerError, "enregistrement impossible")
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true})
}
