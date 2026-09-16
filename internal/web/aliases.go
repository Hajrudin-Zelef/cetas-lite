package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cetas-lite/internal/alias"
	"cetas-lite/internal/catalog"
	"cetas-lite/internal/provider"
)

func (s *Server) handleAliasesGet(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"families": alias.ResolveAll(s.engine.Families()),
		"defaults": alias.ResolveAll(alias.Defaults()),
	})
}

// validateOverrides refuse les couples provider/modele inconnus : un
// selecteur qui pointe vers un modele inexistant casserait les tours.
func validateOverrides(ov alias.Overrides) error {
	for fam, modes := range ov {
		f, ok := alias.Find(alias.Defaults(), fam)
		if !ok {
			return fmt.Errorf("famille inconnue : %s", fam)
		}
		if f.Local {
			for mode, members := range modes {
				for _, mb := range members {
					if !isLocalEngine(mb.Provider) {
						return fmt.Errorf("%s/%s : fournisseur local inconnu : %s", fam, mode, mb.Provider)
					}
					if mb.Model == "" {
						return fmt.Errorf("%s/%s : modele vide", fam, mode)
					}
				}
			}
			continue
		}
		for mode, members := range modes {
			if _, ok := alias.Resolve(alias.Defaults(), fam, mode); !ok {
				return fmt.Errorf("mode inconnu : %s/%s", fam, mode)
			}
			if len(members) == 0 {
				return fmt.Errorf("%s/%s : pool vide (omettez le mode pour revenir au defaut)", fam, mode)
			}
			for _, mb := range members {
				if !provider.IsCloudProvider(mb.Provider) {
					return fmt.Errorf("%s/%s : fournisseur inconnu : %s", fam, mode, mb.Provider)
				}
				if _, ok := catalog.Lookup(mb.Provider, mb.Model); !ok {
					return fmt.Errorf("%s/%s : modele inconnu : %s/%s", fam, mode, mb.Provider, mb.Model)
				}
			}
		}
	}
	return nil
}

func (s *Server) handleAliasesPut(w http.ResponseWriter, r *http.Request) {
	claims := claimsFrom(r)
	if claims == nil {
		writeError(w, http.StatusUnauthorized, "non authentifie")
		return
	}
	var ov alias.Overrides
	if err := decodeJSON(r, &ov); err != nil {
		writeError(w, http.StatusBadRequest, "corps JSON invalide")
		return
	}
	if ov == nil {
		ov = alias.Overrides{}
	}
	if err := validateOverrides(ov); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	fams := alias.Apply(s.engine.Families(), ov)
	s.engine.SetFamilies(fams)
	if raw, err := json.Marshal(ov); err == nil {
		_ = s.st.PutMeta("aliases", raw)
	}
	writeJSON(w, http.StatusOK, map[string]any{"ok": true, "families": alias.ResolveAll(fams)})
}
