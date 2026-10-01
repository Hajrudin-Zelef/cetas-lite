---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-7
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft"]
dates: []
keywords: ["mcp", "agent", "agents", "claude", "memory"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [1077, 1229]
sha256: fa6952048e6dfca99e550a163ef9c24d76ca92fc14cb5272242ec33d21fcb093
---

# MCP, LSP, code-server, websearch — Guide pratique

| Symptôme | Cause probable | Remède |
|---|---|---|
| Le client ne reçoit rien / timeout | `print()` sur stdout | Tout log → `stderr` (`print(..., file=sys.stderr)`) |
| `tools/list` vide | décorateur oublié / serveur mal nommé | Vérifier `@mcp.tool()` + nom dans `MCPServer("...")` |
| Schéma d'entrée bizarre | Type non annoté ou `Optional` mal géré | Annoter tous les params ; défauts explicites |
| Le serveur ne quitte pas (stdio) | Boucle/thread non-daemon | À la fermeture de stdin, sortir proprement |
| `address already in use` (HTTP) | Ancien `mcp run` encore vivant | `pkill -f "mcp run"` / port différent |
| Gros documents → contexte saturé | Pas de borne | `max_chars`/`limit` systématiques (§29) |
| `structuredContent` ignoré | Vieux client | Toujours remplir aussi `content` (texte) |

## 44. Pense-bête développeur MCP

- Description du tool = 50 % de la qualité. Écris-la comme une doc,
  pas comme un commentaire.
- Un tool = une action. Si ton tool a un paramètre `mode` avec 6 valeurs,
  ce sont probablement 6 tools.
- `readOnlyHint: true` quand c'est vrai → moins de frictions (pas de
  confirmation), et c'est honnête.
- Borne tout : `limit`, `max_chars`, timeouts. Le modèle ne se bornera pas
  tout seul.
- `isError: true` + message actionnable > exception obscure.
- Teste au pipe (`echo ... | server.py`) : le test le plus rapide qui existe.
- Ne promets pas `subscriptions/listen` si tu ne détectes pas les changements.
- Versionne, changelog, jamais de secret en dur.

# PARTIE III — PANORAMA DES SERVEURS MCP EXISTANTS

> État fin septembre 2026, vérifié via registres et dépôts. Les serveurs
> « officiels » viennent d'Anthropic (`@modelcontextprotocol/*`) sauf mention.
> 📌 Deux serveurs officiels (postgres, google-drive) sont signalés
> **archivés** par certaines sources : préfère les alternatives actives
> indiquées.

## 45. Officiels : filesystem, git, fetch

- **`@modelcontextprotocol/server-filesystem`** — opérations fichiers dans
  des dossiers explicitement autorisés (passés en args). Le plus installé.
  ```bash
  claude mcp add fs-local -- npx -y @modelcontextprotocol/server-filesystem /home/zelef/workspace
  ```
- **`@modelcontextprotocol/server-git`** — lecture/manip git (log, diff,
  blame, search). Le seul « git MCP » dont tu as besoin.
- **`@modelcontextprotocol/server-fetch`** — récupère une URL et la rend en
  Markdown digeste pour le modèle. Parfait pour « résume cette doc en ligne ».

💡 Ces trois-là + ton `doc-rag` maison couvrent 80 % des besoins quotidiens.

## 46. Officiels : memory, sequential-thinking, sqlite

- **`@modelcontextprotocol/server-memory`** — graphe de connaissances
  persistant (entités/relations/observations) en local. Utile pour « souviens-
  toi que le client X a des onduleurs Easy UPS 3S ».
- **`@modelcontextprotocol/server-sequential-thinking`** — force le modèle à
  décomposer son raisonnement en étapes outillées. Gadget pour certains,
  utile sur les diagnostics complexes.
- **`@modelcontextprotocol/server-sqlite`** (+ forks communautaires) —
  requêtes SQLite locales avec inspection de schéma. Idéal pour un inventaire
  local (parc onduleurs, par exemple).

## 47. Bases de données : postgres, redis, supabase, neon

- **`@modelcontextprotocol/server-postgres`** — ⚠️ signalé archivé par
  plusieurs inventaires 2026. Lecture seule + inspection de schéma à l'origine.
  **Alternative active** : serveurs communautaires maintenus (ex : `dbhub`
  de Bytebase : `npx -y @bytebase/dbhub --dsn "postgresql://..."`, multi-BDD).
- **Redis** : serveurs officiels/communautaires (`redis` : get/set, pub/sub
  selon implémentation).
- **Supabase / Neon** : serveurs hébergés officiels (OAuth) — gèrent le
  Postgres managé : migrations, logs, branches.

🔒 Règle d'or BDD : expose **un rôle SQL en lecture seule** au serveur MCP,
sauf besoin d'écriture explicite. Et jamais les identifiants prod dans une
config versionnée.

## 48. Forges et productivité : GitHub, Slack, Linear, Notion, Jira

- **GitHub** (`@modelcontextprotocol/server-github`, token `GITHUB_TOKEN`) :
  issues, PR, code search, releases, actions. Le plus rentable au quotidien.
- **Slack** (communautaire/officiel selon période) : lire/poster messages —
  ⚠️ irréversible, scope minimal sur les channels utiles.
- **Linear, Notion, Asana** : serveurs hébergés officiels (OAuth via
  `mcp-remote` ou URL directe) — tickets, pages, projets.
- **Atlassian (Jira/Confluence)** : serveur distant `https://mcp.atlassian.com`
  via `mcp-remote`, OAuth 2.1, respecte tes permissions Atlassian.

💡 Pour Zelef : GitHub (tes scripts) + un serveur tickets, et tu as déjà un
cockpit agentique crédible.

## 49. Navigateur : Playwright MCP et chrome-devtools-mcp

- **`@playwright/mcp`** (Microsoft, officiel) — pilote un vrai Chromium via
  l'arbre d'accessibilité (pas besoin de modèle vision) : snapshot, clic,
  formulaire, navigation. LE serveur pour « va vérifier sur cette page ».
  ```bash
  claude mcp add playwright -- npx -y @playwright/mcp
  ```
- **`chrome-devtools-mcp`** (Google) — DevTools en MCP : performance, réseau,
  mémoire, screencast. Complémentaire de Playwright pour le debug.
- **`@anthropic/claude-in-chrome-mcp-server`** — pilote ton Chrome perso
  (avec tes sessions) — pratique, mais 🔒 périmètre énorme : à réserver à
  un profil navigateur dédié.

## 50. Documentation vivante : Context7, DeepWiki

- **`@upstash/context7-mcp`** — injecte la doc **à jour** des librairies
  dans le prompt. Fini les API hallucinées parce que le modèle connaît la
  version 2023. Outil anti-hallucination n°1 des codeurs.
- **DeepWiki** (communautaire) — génère une doc structurée de n'importe quel
  dépôt GitHub (architecture, API, relations).

💡 Combo gagnant pour apprendre un outil : Context7 (doc officielle à jour)
+ DeepWiki (repo réel).

## 51. Recherche web : Brave, Tavily, Exa, Firecrawl (MCP natifs)

Tous les fournisseurs du §99+ proposent un serveur MCP officiel :
- **Brave Search** (serveur MCP officiel) — recherche indépendante.
- **Tavily** (MCP officiel) — `tavily-search` + `tavily-extract`, pensé agents.
- **Exa** (MCP officiel, quota spécifique : ~150 appels/jour sur l'offre
  d'entrée) — recherche neurale + contenu de pages.
- **Firecrawl** (MCP officiel) — search + scrape en un appel.

💡 Si ton agent a besoin du web, **ne bricole pas de curl** : prends le MCP
officiel du fournisseur que tu as choisi (§109-112). Tu gagnes le formatage
« LLM-ready » et la gestion des quotas.

## 52. mcp-remote : le pont vers les serveurs OAuth distants

Beaucoup de serveurs hébergés parlent encore l'ancien HTTP+SSE ou exigent
OAuth. **`mcp-remote`** (proxy local, `npx -y mcp-remote <url>`) fait la
traduction : il expose du stdio à ton client et gère OAuth + transport
distant côté serveur.

```json
{"mcpServers": {"atlassian": {
  "command": "npx",
  "args": ["-y", "mcp-remote", "https://mcp.atlassian.com/v1/sse"]}}}
```

💡 À utiliser uniquement pour les serveurs de confiance qui n'offrent pas
de Streamable HTTP natif. C'est une rustine, pas une architecture.

## 53. mcp-language-server : le pont LSP → MCP

**`mcp-language-server`** (communautaire, isaacphi) : expose les capacités
LSP (voir Partie V) comme des tools MCP — `definition`, `references`,
`hover`, `diagnostics` — pour 30+ langages via les vrais serveurs de langage.

💡 C'est la jonction des deux mondes de ce guide : ton agent MCP obtient
l'intelligence sémantique du LSP sans réimplémenter le protocole.

## 54. Infra : docker, kubernetes, cloudflare, sentry

