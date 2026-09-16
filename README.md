# cetas-lite

*Variante légère, stable et rapide de Cetas + Marexcode — un seul binaire Go.*

© Marexsoft Corporation — Fondateur : Kouassi Marius. Tous droits réservés.

## État

**Phases P0–P11 faites** : socle (bbolt, coffre AES-256-GCM, auth JWT), alias, moteur de chat stable
(journal rejouable, reconnexion `?from=`, stop/reset, heartbeat SSE, failover), outils agent + sandbox,
UI web (révélation, markdown, blocs outils), web search + fetch, mémoire Markdown (`mem_*`).
**P6** : durcissement (rate-limit auth, fuzz, tests de reconnexion) + packaging 6 binaires + CI.
**P7** : compaction du contexte + archives multi-sessions. **P8.1** : client MCP (`mcp_*`).
**Tier 1–4** : archives UI, export/régénérer, inscription bootstrap + proxy + bwrap + favicon,
custom tools HTTP (`custom_*`). **Lots 0/A/C/B/D** : prompts pédagogiques + `MAREX.md`, lecture de
documents (PDF/texte/HTML), thinking+effort, vision, voix navigateur. **P9** : registre d'outils + docs.
**P10** : skills, WebDepth, GitHub tools, pièces jointes TTL, menus composer portalés.
**P11** : coffre **V4** (`internal/securevault`) + UI Coffre, CLI **`cetas-keys`**, panneaux
API Modèles / Configuration / centre d'aide, sidebar Agents plate, thème modale.
**UI** : le front reprend le design de Cetas (thèmes clair/ocean/sombre, glassmorphism, sidebar +
`input-area` + `plus-menu`, messages `.message-wrapper`) adapté au backend cetas-lite.

Reste : LSP (faible valeur) et gestionnaire de moteur local **niveau B** (llama.cpp/GPU — différé).

- Alias à 2 niveaux : `SamAgent Nano` (free), `SamAgent N4` (flash/standard), `SamAgent N8` (flash/standard/elite), `Code` (flash/standard/elite, agent), `SamGen` (local : nano=llama.cpp, n4=Ollama, n8=LM Studio).
- Chat serveur : journal rejouable, reconnexion (`?from=`), stop/reset non bloquants, heartbeat SSE, failover de pool.
- Providers cloud : DeepSeek, OpenCode Zen, OpenCode Go, OpenRouter (en-tête `x-opencode-session` requis par la gateway). Local : Ollama/LM Studio/llama.cpp (découverte auto). Panneau **API Modèles** (catalogue, tarifs, logos, toggle textes/images).
- **Coffre chiffré** : `internal/securevault`, format **V4** compatible CETAS/Python (`$CETAS_LITE_HOME/vault.enc`, AES-256-GCM + Scrypt N=2^16 + nonce HKDF + pepper `CETAS_PEPPER`), migration V2/V3, verrouillage auto 15 min, UI onglet « Coffre ».
- **`cetas-keys`** : CLI séparée (`cmd/cetas-keys`) pour gérer le coffre et un `.env` **scellé** (`init/add/list/test/delete/export/passwd/status/sync`), validation **live** des clés, bannière TUI, 16 providers.
- **Panneaux UI** : API Modèles, Configuration (onglets Fonctionnalités / Recherche Web / Apparence / Remote SFTP / Compétences, sauvegarde par onglet), centre d'aide intégré (recherche plein-texte), sidebar Agents à 2 sections plates.
- Outils agent : fichiers (Ls/Read/Write/Edit/Grep/Glob), Bash, RunScript (désactivé par défaut), TodoWrite, web (`web_search`/`web_fetch`), mémoire (`mem_*`), MCP (`mcp_*`), GitHub (repos/issues/PRs).
- Mémoire : pages Markdown par user sous `$CETAS_LITE_HOME/memory/<user>/`, index `MEMORY.md` auto, recherche TF-IDF.
- MCP : serveurs déclarés dans `$CETAS_LITE_HOME/mcp.json` (`stdio` ou `http`), outils exposés à l'agent sous `mcp_<serveur>_<outil>` (diagnostic : `./bin/cetas-lite mcp`).
- Custom tools : outils HTTP définis dans `$CETAS_LITE_HOME/tools.json`, exposés sous `custom_<outil>` (diagnostic : `./bin/cetas-lite tools`).
- Plugins externes : dossiers `$CETAS_LITE_HOME/plugins/<nom>/plugin.json`, outils `exec` (tout langage, JSON sur stdin/stdout) ou `http`, exposés sous `plugin_<plugin>_<outil>` (exemple : `examples/plugins/horloge/`, doc : `docs/plugins.md`).
- Rôle : **cetas-lite = tuteur/professeur senior** (chat général, pédagogique) ; l'**agent = code pur** (outils).
- Thinking : **off par défaut** en chat (toggle `Think`), **forcé** en agent ; effort `défaut/faible/moyen/max`.
- Lecture : **images** (vision, modèles déclarés dans Settings → Capacités) et **documents** (PDF/texte/HTML) en pièces jointes (bouton + drag & drop). Pas d'Office (convertir en PDF).
- Voix : TTS (« Lire ») et STT (micro) via le **navigateur** ; backend non implémenté.
- Archives : liste/restauration/suppression/export dans la sidebar ; export Markdown ou JSON de l'active.
- Confort : copier/régénérer un message, jetons affichés, compaction visible, citations cliquables, notice de session expirée.
- Sauvegarde : `cetas-lite backup <fichier.tar.gz>` / `restore <fichier>` (serveur arrêté).
- **Terminal intégré** : tiroir bas avec onglets, xterm.js embarqué (fonctionne hors-ligne),
  vrai PTY natif (Unix) / **ConPTY** (Windows, via l'API pseudo-console : redimensionnement
  et programmes interactifs supportés), jusqu'à 6 sessions par utilisateur, `cwd` confiné au
  workspace, sortie diffusée en SSE (base64), redimensionnement dynamique.
- **Worktrees d'isolation** : chaque run agent peut s'exécuter dans un `git worktree --detach`
  dédié sous `$CETAS_LITE_HOME/worktrees` (toggle dans le modal), nettoyé à la suppression.
- **Multi-agents en parallèle** : panneau « Agents » (création via modal : mission, famille/mode
  agent, approbations, plan, worktree), chaque agent a son fil SSE, ses approbations et son
  worktree ; suivi/stop/suppression, état `running/done/stopped/error`, persistance après
  redémarrage (reprise à l'arrêt).

## Démarrage

```bash
make build
CETAS_LITE_HOME=~/.cetas-lite ./bin/cetas-lite serve
# http://127.0.0.1:8787/api/health
```

## Application bureau (Windows)

```bash
make desktop
# -> bin/cetas-lite-desktop-windows-amd64.exe : double-clic, aucune console.
```

Le mode bureau demarre le serveur en local sur `127.0.0.1` (port ephemere, `CETAS_LITE_ADDR`
ignore) puis ouvre une fenetre native **WebView2** (moteur Edge, inclus dans Windows 10/11) ;
la fermeture de la fenetre arrete proprement le serveur. Les logs vont aussi dans
`$CETAS_LITE_HOME/desktop.log` (pas de console en mode GUI). `CETAS_LITE_DEBUG=true`
active les outils de dev du WebView.

```bash
cetas-lite desktop   # depuis un terminal, ou sans argument sous Windows
```

Alias et chat (après login) :

```bash
curl -s localhost:8787/api/aliases -H "Authorization: Bearer $TOKEN"
curl -s -X POST localhost:8787/api/chat/send -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' -d '{"family":"code","mode":"standard","message":"salut"}'
curl -sN "localhost:8787/api/chat/stream?from=0" -H "Authorization: Bearer $TOKEN"
```

Moteurs locaux (optionnel) : `CETAS_LITE_OLLAMA_URL`, `CETAS_LITE_LMSTUDIO_URL`, `CETAS_LITE_LLAMACPP_URL`.

## Sécurité (variables d'environnement)

- `CETAS_LITE_REGISTRATION_OPEN` : non défini = **bootstrap** (le 1er compte est créé, puis l'inscription
  se ferme) ; `true` = ouvert ; `false` = fermé.
- `CETAS_LITE_TRUST_PROXY=true` : derrière un reverse proxy (nginx), utiliser le dernier hop
  `X-Forwarded-For`/`X-Real-IP` pour le rate-limit par IP.
- `CETAS_LITE_SANDBOX=none|auto|bwrap` (défaut `none`) : isole `Bash`/`RunScript` dans **bubblewrap**
  (système en lecture seule, bind du seul workspace) ; sonde au démarrage, repli sûr.
- `CETAS_LITE_ALLOW_SCRIPT` (défaut `false`) : active l'outil `RunScript`.
- `CETAS_PEPPER` : pepper du coffre V4 (`internal/securevault`), préfixé au mot de passe maître avant Scrypt.
  **Le perdre rend le coffre `vault.enc` indéchiffrable.**

## Clés providers (chiffrées)

Deux coffres coexistent :

```bash
# clés providers côté serveur (bbolt `secrets`, cryptovault, serveur arrêté)
export CETAS_LITE_VAULT_PASSWORD='...'
./bin/cetas-lite keys set deepseek 'sk-...'
./bin/cetas-lite keys list

# coffre V4 + .env scellé (CLI séparée, serveur peut tourner)
go build -o bin/cetas-keys ./cmd/cetas-keys/
export CETAS_PEPPER='...'
./bin/cetas-keys            # assistant interactif
./bin/cetas-keys list
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
make desktop       # app bureau Windows (amd64 + arm64, subsystem GUI -> bin/*.exe)
```

`CGO_ENABLED=0`, `-trimpath`, `-ldflags "-s -w -X main.version=..."`. CI GitHub Actions : tests
matrice (ubuntu/windows/macos) + artifact cross-build.

## Architecture

- `cmd/cetas-lite` — CLI (`serve`, `keys`, `mcp`, `tools`, `backup`, `restore`, `version`)
- `cmd/cetas-keys` — CLI du coffre V4 (assistant + `init/add/list/test/delete/export/passwd/status/sync`)
- `internal/config` — configuration et répertoires
- `internal/store` — bbolt (users, settings, conversations, secrets)
- `internal/cryptovault` — AES-256-GCM + scrypt (clés `keys`, N=32768)
- `internal/securevault` — coffre V4 `vault.enc` (Scrypt N=2^16, nonce HKDF, pepper `CETAS_PEPPER`)
- `internal/keysetup` — logique du CLI `cetas-keys` (seal, validation live, `.env` scellé)
- `internal/auth` — scrypt + JWT HS256
- `internal/chat` — conversation/journaux, agent, registre d'outils, sandbox/isolation
- `internal/provider` / `internal/local` — providers cloud + découverte locale
- `internal/search` — web search + fetch (garde SSRF)
- `internal/memory` — pages Markdown + index + TF-IDF
- `internal/mcp` — client MCP (stdio + HTTP, outils `mcp_*`)
- `internal/customtools` — outils HTTP d'opérateur (`custom_*`)
- `internal/plugins` — plugins externes (`plugin_*` : manifeste + exec/HTTP)
- `internal/docs` / `internal/attach` — extraction documents + pièces jointes
- `internal/modelcaps` — capacités provider/model (vision/tts/stt)
- `internal/backup` — bundle tar.gz (base + mcp.json), restauration anti-traversée
- `internal/web` — routeur HTTP + middleware
- `web/` — assets embarqués (`go:embed`)

## Licence

Propriétaire — tous droits réservés © Marexsoft Corporation.
