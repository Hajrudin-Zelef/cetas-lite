// Package vault ouvre le coffre chiffre partage qui protege les cles API
// des providers. Le mot de passe vient de CETAS_LITE_VAULT_PASSWORD et le
// sel est persiste dans les metas du store. Utilise par la CLI et le
// serveur web : une seule implementation, un seul comportement.
package vault

import (
	"errors"
	"os"

	"cetas-lite/internal/cryptovault"
	"cetas-lite/internal/store"
)

// Open ouvre le coffre chiffre. Cree le sel s'il n'existe pas encore.
func Open(st *store.Store) (*cryptovault.Vault, error) {
	password := os.Getenv("CETAS_LITE_VAULT_PASSWORD")
	if password == "" {
		return nil, errors.New("CETAS_LITE_VAULT_PASSWORD requis")
	}
	salt, ok := st.GetMeta("vault_salt")
	if !ok {
		var err error
		salt, err = cryptovault.NewSalt()
		if err != nil {
			return nil, err
		}
		if err := st.PutMeta("vault_salt", salt); err != nil {
			return nil, err
		}
	}
	return cryptovault.New(password, salt)
}
