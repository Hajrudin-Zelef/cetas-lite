# cetas-lite

*Variante légère, stable et rapide de Cetas + Marexcode — un seul binaire Go.*

© Marexsoft Corporation — Fondateur : Kouassi Marius. Tous droits réservés.

## État

**Phase 1 (chat)** : socle (bbolt, coffre AES-256-GCM, auth JWT) + **alias** + **moteur de chat stable**.
- Alias à 2 niveaux : `SamAgent Nano` (free), `SamAgent N4` (flash/standard), `SamAgent N8` (flash/standard/elite), `Code` (flash/standard/elite, agent), `SamGen` (local : nano=llama.cpp, n4=Ollama, n8=LM Studio).
- Chat serveur : journal rejouable, reconnexion (`?from=`), stop/reset non bloquants, heartbeat SSE, failover de pool.
- Providers cloud : DeepSeek, OpenCode Zen, OpenCode Go, OpenRouter. Local : Ollama/LM Studio/llama.cpp (découverte auto).
- UI web et outils agent : phases suivantes.

## Démarrage

```bash
make build
CETAS_LITE_HOME=~/.cetas-lite ./bin/cetas-lite serve
# http://127.0.0.1:8787/api/health
```

Alias et chat (après login) :

```bash
curl -s localhost:8787/api/aliases -H "Authorization: Bearer $TOKEN"
curl -s -X POST localhost:8787/api/chat/send -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' -d '{"family":"code","mode":"standard","message":"salut"}'
curl -sN "localhost:8787/api/chat/stream?from=0" -H "Authorization: Bearer $TOKEN"
```

Moteurs locaux (optionnel) : `CETAS_LITE_OLLAMA_URL`, `CETAS_LITE_LMSTUDIO_URL`, `CETAS_LITE_LLAMACPP_URL`.

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
