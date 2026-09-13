# cetas-lite

*Variante légère, stable et rapide de Cetas + Marexcode — un seul binaire Go.*

© Marexsoft Corporation — Fondateur : Kouassi Marius. Tous droits réservés.

## État

**Phase 0 (socle)** : serveur HTTP, bbolt, coffre chiffré AES-256-GCM, auth JWT multi-utilisateurs, health, assets embarqués. Les phases chat/agent/web/mémoire suivent (voir `docs/plans/`).

## Démarrage

```bash
make build
CETAS_LITE_HOME=~/.cetas-lite ./bin/cetas-lite serve
# http://127.0.0.1:8787/api/health
```

## Clés providers (chiffrées)

```bash
export CETAS_LITE_VAULT_PASSWORD='...'
./bin/cetas-lite keys set deepseek 'sk-...'
./bin/cetas-lite keys list
```

## Tests

```bash
make test
```

## Architecture

- `cmd/cetas-lite` — CLI (`serve`, `keys`, `version`)
- `internal/config` — configuration et répertoires
- `internal/store` — bbolt (users, settings, conversations, secrets)
- `internal/cryptovault` — AES-256-GCM + scrypt
- `internal/auth` — scrypt + JWT HS256
- `internal/web` — routeur HTTP + middleware
- `web/` — assets embarqués (`go:embed`)

## Licence

Propriétaire — tous droits réservés © Marexsoft Corporation.
