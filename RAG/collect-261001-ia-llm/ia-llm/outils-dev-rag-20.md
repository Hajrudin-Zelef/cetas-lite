---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-20
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Microsoft", "OpenAI"]
dates: ["2026-02-26", "2026-09-23", "2026-09-27"]
keywords: ["agent", "agentic", "agents", "aws", "claude", "copilot", "embedding", "embeddings", "mai", "mcp"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [3728, 3830]
sha256: f1ba2318aad30f13833893f108c17bc5f8c8852b5a247313b913d7bf8dc3962a
---

# Outils dev + ingénierie RAG (chunk & corpus)

- **VS Code 1.139 sorti le 23/09/2026** (stable) : sessions agents plus
  rapides, agents dans les Dev Containers distants (SSH/Tunnel/WSL),
  réglage `chat.agentHost.devContainer.enabled` (déploiement progressif).
  Insiders déjà à **1.140** → la prochaine stable suit le rythme
  bimensuel habituel (vérifié : ntcompatible, visualstudiomagazine, sept 2026).
- **Agents window en Stable (preview)** depuis mai 2026 (v1.120–1.123) :
  sessions distantes, BYOK (y compris environnements air-gapped), **Agent
  Host Protocol (AHP)** — protocole ouvert de synchronisation des sessions
  entre clients. La direction est claire : l'AHP et les agents distants
  sont l'architecture d'avenir, pas un gadget (vérifié : changelog Copilot
  mai 2026, formation msbart2 sept 2026).
- **Ce que ça change pour toi** : quand les agents distants sortiront de
  preview, ton workflow « collecte sur le serveur via session distante »
  deviendra natif. En attendant, le Remote-SSH classique (§1-10) reste la
  voie stable.

## 127. Supabase : ce qui bouge (annoncé, changelog sept 2026)

- **Supabase Pipelines — public alpha** (tous plans payants) : CDC managé
  qui réplique les changements Postgres vers BigQuery en quasi temps réel
  (vérifié : changelog supabase.com, sept 2026).
- **Passkeys pour Supabase Auth — Beta** (juin 2026) : connexion sans mot
  de passe (WebAuthn) — à surveiller pour sécuriser ton app RAG plus tard
  (vérifié : changelog, juin 2026).
- **Multigres 0.1 — alpha** (juin 2026) : Postgres sharded de Supabase
  (vérifié : changelog, juin 2026). Pertinent seulement si ton corpus
  dépasse un jour un seul nœud — pas ton cas.
- **Vector Buckets** (annoncés déc. 2025, partenariat AWS) : stocker de
  gros jeux de vecteurs sur S3 en les interrogeant depuis Postgres — utile
  quand `rag_chunks` dépassera la RAM confortable (vérifié : annonce
  Supabase/AWS, déc. 2025).
- **⚠ Rupture déjà effective : l'endpoint `logs.all` de la Management API a
  été supprimé le 23/09/2026** → remplacé par `analytics/endpoints/logs`
  (ClickHouse SQL, table unifiée). Si tes scripts l'appelaient, mets à
  jour ; le serveur MCP passe en **v0.10.0** (vérifié : byteiota, sept 2026).
- **Realtime** : filtres AND composés, opérateurs `like`/`ilike`/`is`/
  `match`/`imatch`/`isdistinct`, sélection des colonnes du payload —
  réduit les coûts sur abonnements volumineux (vérifié : sept 2026).
- **Limites Edge Functions relevées** : Pro 500 → 1 000 fonctions, Team
  1 000 → 2 000 (vérifié : sept 2026).

## 128. pgvector : mets à jour (annoncé, versions vérifiées)

Situation au 27/09/2026 (vérifié : changelog pgvector, analyses sept 2026) :

| Version | Date | Contenu |
|---|---|---|
| 0.8.0 | oct. 2024 | `iterative_scan` (filtered ANN), meilleure estimation du planificateur |
| 0.8.2 | 26/02/2026 | **CVE-2026-3172** : buffer overflow en build HNSW parallèle |
| 0.8.3 | juin 2026 | **corruption d'index HNSW pendant le vacuum** — critique |
| 0.8.4 | juil. 2026 | builds IVFFlat ne dépassent plus `maintenance_work_mem` |
| 0.8.6 | juil. 2026 | buffer overflow en construction IVFFlat (builds 32-bit) |

**Action pour toi** :
```sql
SELECT extversion FROM pg_extension WHERE extname = 'vector';
-- Si < 0.8.3 : mets à jour SANS TARDER (corruption HNSW au vacuum).
-- Les Postgres managés épinglent la version par majeure : vérifie ce que
-- ton hébergeur expose avant de supposer que tu es à jour.
```
Le 0.8.0 a aussi apporté `hnsw.max_scan_tuples`, `ivfflat.max_probes`,
`hnsw.scan_mem_multiplier` — les réglages fins de §65-66 restent valides.

## 129. Playwright : agents de test et MCP (annoncé)

- **Playwright 1.56** : **Playwright Test Agents** — trois définitions
  d'agents (planner / generator / healer) pour guider un LLM dans la
  création et la réparation de tests :
  `npx playwright init-agents --loop=vscode|claude|opencode`
  (vérifié : release notes playwright.dev, 2026).
- **Playwright ≥ 1.62** : le **serveur MCP Playwright et `playwrightcli`
  sont embarqués** — `npx playwright mcp` expose le navigateur à un agent
  IA. **Pertinent pour ton pipeline de collecte** : un agent pourra
  piloter le scraping via MCP au lieu de scripts ad hoc
  (vérifié : analyse THEXARIS Labs sur les release notes 1.62, août 2026).
- **Playwright 1.57** : bascule vers les builds **Chrome for Testing**
  (au lieu de Chromium) ; suppression de `page.accessibility` déprécié
  depuis 3 ans (vérifié : release notes Python).
- ⚠ Les numéros exacts varient selon les canaux au 27/09/2026 — vérifie
  avec `npx playwright --version` avant de t'appuyer sur une API précise.

## 130. Rumeurs, non-annoncés et méthode de veille

- **RUMEUR non confirmée — « Project Polaris »** : un modèle de code
  maison GitHub destiné à remplacer GPT-4 Turbo dans Copilot, évoqué par
  un seul média tech (juin 2026). Non confirmé par GitHub/Microsoft —
  **à ne pas citer comme un fait**.
- **GitHub — annoncé** : migrations d'images Actions à venir (impacte tes
  workflows CI, §20), **Agentic Workflows en public preview**
  (workflows markdown compilés en Actions durcis), facturation Copilot à
  l'usage (vérifié : digests GitHub 2026).
- **OpenAI embeddings — rien de nouveau officialisé au 27/09/2026** :
  `text-embedding-3-small` (janv. 2024) reste la génération courante ;
  aucun successeur annoncé. Côté génération : **GPT-5.4 mini / nano**
  annoncés (mars 2026) — options à évaluer pour `RAG_LLM` (§123).
- **tmux, ngrok, Tailscale, Cloudflare Tunnel — rien d'officialisé** dans
  mes recherches au 27/09/2026 : pas de roadmap publique trouvée.
- **Méthode de veille** (5 min/mois) : changelog supabase.com, release
  notes playwright.dev, `github.com/microsoft/vscode` releases, changelog
  pgvector, `SELECT extversion` sur ta base. Note la date de chaque
  vérification — comme ce guide le fait.

---
