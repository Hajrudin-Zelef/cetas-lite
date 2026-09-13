# cetas-lite

*Variante légère, stable et rapide de Cetas + Marexcode — un seul binaire Go.*

© Marexsoft Corporation — Fondateur : Kouassi Marius. Tous droits réservés.

## État

**Phases P0–P5 faites** : socle (bbolt, coffre AES-256-GCM, auth JWT), alias, moteur de chat stable
(journal rejouable, reconnexion `?from=`, stop/reset, heartbeat SSE, failover), outils agent + sandbox,
UI web (révélation, markdown, blocs outils), web search + fetch, mémoire Markdown (`mem_*`).
**P6** : durcissement (rate-limit auth, fuzz, tests de reconnexion) + packaging 6 binaires + CI.

- Alias à 2 niveaux : `SamAgent Nano` (free), `SamAgent N4` (flash/standard), `SamAgent N8` (flash/standard/elite), `Code` (flash/standard/elite, agent), `SamGen` (local : nano=llama.cpp, n4=Ollama, n8=LM Studio).
- Chat serveur : journal rejouable, reconnexion (`?from=`), stop/reset non bloquants, heartbeat SSE, failover de pool.
- Providers cloud : DeepSeek, OpenCode Zen, OpenCode Go, OpenRouter. Local : Ollama/LM Studio/llama.cpp (découverte auto).
- Outils agent : fichiers (Ls/Read/Write/Edit/Grep/Glob), Bash, RunScript (désactivé par défaut), TodoWrite, web (`web_search`/`web_fetch`), mémoire (`mem_*`).
- Mémoire : pages Markdown par user sous `$CETAS_LITE_HOME/memory/<user>/`, index `MEMORY.md` auto, recherche TF-IDF.

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
make test          # go test ./... -race
make ci            # gofmt check + vet + test -race
make smoke         # smoke UI Playwright (python3 + playwright + chromium)
```

## Build & distribution

```bash
make build         # binaire local -> bin/cetas-lite
make cross         # linux/windows/darwin x amd64/arm64 -> bin/
```

`CGO_ENABLED=0`, `-trimpath`, `-ldflags "-s -w -X main.version=..."`. CI GitHub Actions : tests
matrice (ubuntu/windows/macos) + artifact cross-build.

## Architecture

- `cmd/cetas-lite` — CLI (`serve`, `keys`, `version`)
- `internal/config` — configuration et répertoires
- `internal/store` — bbolt (users, settings, conversations, secrets)
- `internal/cryptovault` — AES-256-GCM + scrypt
- `internal/auth` — scrypt + JWT HS256
- `internal/chat` — conversation/journaux, agent, outils, sandbox
- `internal/provider` / `internal/local` — providers cloud + découverte locale
- `internal/search` — web search + fetch (garde SSRF)
- `internal/memory` — pages Markdown + index + TF-IDF
- `internal/web` — routeur HTTP + middleware
- `web/` — assets embarqués (`go:embed`)

## Licence

Propriétaire — tous droits réservés © Marexsoft Corporation.
